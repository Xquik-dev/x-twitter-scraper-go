// Copyright the Xquik contributors.
// SPDX-License-Identifier: Apache-2.0

package coveragegate_test

import (
	"bytes"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"github.com/Xquik-dev/x-twitter-scraper-go/internal/coveragegate"
)

func failingBaselineFixture(t *testing.T) (*os.Root, string) {
	t.Helper()
	directory, _ := baselineFixture(t)
	const source = "package fixture\nimport \"testing\"\nfunc TestNativeFailure(t *testing.T) { t.Fatal(\"native fixture failure\") }\n"
	if err := directory.WriteFile("failure_test.go", []byte(source), 0o600); err != nil {
		t.Fatal(err)
	}
	commitFixture(t, directory)
	revision := fixtureCommand(t, directory, exec.Command("git", "rev-parse", "HEAD"))
	return directory, strings.TrimSpace(string(revision))
}

func TestRecordBaselineRetainsFailedNativeEvidence(t *testing.T) {
	directory, revision := failingBaselineFixture(t)
	var output bytes.Buffer
	_, err := coveragegate.Record(directory.Name(), &output)
	if err == nil || !strings.Contains(err.Error(), "collect statements") || !strings.Contains(output.String(), "native fixture failure") {
		t.Fatalf("failed collection: %v\n%s", err, output.String())
	}
	baseline := filepath.Join(".git", "coverage-baselines", "go-packages-v3", revision)
	if data, err := directory.ReadFile(filepath.Join(baseline, "unit-coverage.out")); err != nil || len(data) == 0 {
		t.Fatalf("failed collection lost native evidence: %d bytes, %v", len(data), err)
	}
	if _, err := directory.Stat(filepath.Join(baseline, "report.json")); !os.IsNotExist(err) {
		t.Fatalf("failed collection published a baseline: %v", err)
	}
}

func TestRecordBaselinePreservesOccupiedPublicationPaths(t *testing.T) {
	for _, filename := range []string{"report.json", ".baseline-pending.json"} {
		t.Run(filename, func(t *testing.T) {
			directory, _ := baselineFixture(t)
			source := fmt.Sprintf("package fixture\nimport (\"os\"; \"path/filepath\"; \"testing\")\nfunc TestPublicationCollision(t *testing.T) { if err := os.WriteFile(filepath.Join(os.Getenv(\"XQUIK_COVERAGE_DIR\"), %q), []byte(\"retained occupied evidence\\n\"), 0600); err != nil { t.Fatal(err) } }\n", filename)
			if err := directory.WriteFile("publication_test.go", []byte(source), 0o600); err != nil {
				t.Fatal(err)
			}
			commitFixture(t, directory)
			revision := strings.TrimSpace(string(fixtureCommand(t, directory, exec.Command("git", "rev-parse", "HEAD"))))
			var output bytes.Buffer
			if _, err := coveragegate.Record(directory.Name(), &output); err == nil || !strings.Contains(err.Error(), filename) {
				t.Fatalf("occupied publication path: %v\n%s", err, output.String())
			}
			baseline := filepath.Join(".git", "coverage-baselines", "go-packages-v3", revision)
			if data, err := directory.ReadFile(filepath.Join(baseline, filename)); err != nil || string(data) != "retained occupied evidence\n" {
				t.Fatalf("occupied evidence changed: %q, %v", data, err)
			}
			if filename == ".baseline-pending.json" {
				if _, err := directory.Stat(filepath.Join(baseline, "report.json")); !os.IsNotExist(err) {
					t.Fatalf("occupied pending evidence published a baseline: %v", err)
				}
			}
		})
	}
}

func TestRecordBaselineConfinesEvidenceDirectory(t *testing.T) {
	for _, kind := range []string{"file", "outward symlink"} {
		t.Run(kind, func(t *testing.T) {
			directory, _ := baselineFixture(t)
			outside := fixtureRoot(t)
			if kind == "file" {
				if err := directory.WriteFile(".git/coverage-baselines", []byte("retained directory obstacle\n"), 0o600); err != nil {
					t.Fatal(err)
				}
			} else if err := directory.Symlink(outside.Name(), ".git/coverage-baselines"); err != nil {
				t.Fatal(err)
			}
			var output bytes.Buffer
			if _, err := coveragegate.Record(directory.Name(), &output); err == nil || output.Len() != 0 {
				t.Fatalf("invalid evidence directory reached collection: %v\n%s", err, output.String())
			}
			if entries, err := fs.ReadDir(outside.FS(), "."); err != nil || len(entries) != 0 {
				t.Fatalf("collection wrote outside Git: %v, %v", entries, err)
			}
			if kind == "file" {
				if data, err := directory.ReadFile(".git/coverage-baselines"); err != nil || string(data) != "retained directory obstacle\n" {
					t.Fatalf("directory obstacle changed: %q, %v", data, err)
				}
			}
		})
	}
}

func TestCollectStatementsPreservesOccupiedXML(t *testing.T) {
	directory, _ := baselineFixture(t)
	reports := fixtureRoot(t)
	if err := reports.WriteFile("coverage.xml", []byte("retained XML evidence\n"), 0o600); err != nil {
		t.Fatal(err)
	}
	var output bytes.Buffer
	if err := coveragegate.CollectStatements(directory.Name(), reports.Name(), &output); err == nil {
		t.Fatalf("accepted occupied XML: %s", output.String())
	}
	if data, err := reports.ReadFile("coverage.xml"); err != nil || string(data) != "retained XML evidence\n" {
		t.Fatalf("occupied XML changed: %q, %v", data, err)
	}
}

func TestCollectStatementsRejectsUnreadableChildEvidence(t *testing.T) {
	directory, _ := baselineFixture(t)
	reports := fixtureRoot(t)
	if err := reports.Mkdir("child-coverage.out", 0o700); err != nil {
		t.Fatal(err)
	}
	var output bytes.Buffer
	if err := coveragegate.CollectStatements(directory.Name(), reports.Name(), &output); err == nil {
		t.Fatalf("accepted a directory as child evidence: %s", output.String())
	}
	if _, err := reports.Stat("coverage.xml"); !os.IsNotExist(err) {
		t.Fatalf("unreadable child evidence published XML: %v", err)
	}
	if info, err := reports.Stat("child-coverage.out"); err != nil || !info.IsDir() {
		t.Fatalf("child evidence obstacle changed: %v, %v", info, err)
	}
}

func TestCollectorCommandRejectsNativeTestFailure(t *testing.T) {
	for _, test := range []struct {
		name    string
		command *exec.Cmd
	}{
		{"statements", exec.Command("./.git/coverage-gain", "--collect-statements")},
		{"branches", exec.Command("./.git/coverage-gain", "--collect-branches")},
		{"record", exec.Command("./.git/coverage-gain", "--record")},
	} {
		t.Run(test.name, func(t *testing.T) {
			directory, _ := failingBaselineFixture(t)
			buildCollectorCommand(t, directory)
			reports := fixtureRoot(t)
			command := test.command
			command.Dir = directory.Name()
			command.Env = append(os.Environ(), "XQUIK_COVERAGE_DIR="+reports.Name(), "GOCOVERDIR="+os.Getenv("XQUIK_GO_COVERAGE_DIR"))
			var stdout, stderr bytes.Buffer
			command.Stdout, command.Stderr = &stdout, &stderr
			err := command.Run()
			if exit, ok := err.(*exec.ExitError); !ok || exit.ExitCode() != 1 || stdout.Len() != 0 || !strings.Contains(stderr.String(), "native fixture failure") {
				t.Fatalf("failed collector result: %v, stdout=%q, stderr=%q", err, stdout.String(), stderr.String())
			}
		})
	}
}

func TestRecordBaselineRejectsUnavailableGitSource(t *testing.T) {
	directory := fixtureRoot(t)
	var output bytes.Buffer
	if _, err := coveragegate.Record(directory.Name(), &output); err == nil || output.Len() != 0 {
		t.Fatalf("non-Git source reached collection: %v\n%s", err, output.String())
	}
	if entries, err := fs.ReadDir(directory.FS(), "."); err != nil || len(entries) != 0 {
		t.Fatalf("non-Git source changed: %v, %v", entries, err)
	}
}

func TestRecordBaselineRejectsInvalidNativeRuntimeConfiguration(t *testing.T) {
	directory, _ := baselineFixture(t)
	t.Setenv("GOTOOLCHAIN", "invalid")
	var output bytes.Buffer
	_, err := coveragegate.Record(directory.Name(), &output)
	if err == nil || !strings.Contains(err.Error(), "read native Go toolchain") || output.Len() != 0 {
		t.Fatalf("invalid runtime configuration reached collection: %v\n%s", err, output.String())
	}
	if _, err := directory.Stat(".git/coverage-baselines"); !os.IsNotExist(err) {
		t.Fatalf("invalid runtime configuration reserved evidence: %v", err)
	}
}

func TestRecordBaselinePreservesUnwritableRevisionDirectory(t *testing.T) {
	directory, revision := baselineFixture(t)
	const path = ".git/coverage-baselines/go-packages-v3"
	if err := directory.MkdirAll(path, 0o700); err != nil {
		t.Fatal(err)
	}
	if err := directory.Chmod(path, 0o500); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if err := directory.Chmod(path, 0o700); err != nil {
			t.Fatal(err)
		}
	})
	var output bytes.Buffer
	if _, err := coveragegate.Record(directory.Name(), &output); err == nil || output.Len() != 0 {
		t.Fatalf("unwritable revision directory reached collection: %v\n%s", err, output.String())
	}
	if _, err := directory.Stat(filepath.Join(path, revision)); !os.IsNotExist(err) {
		t.Fatalf("unwritable evidence directory changed: %v", err)
	}
}

func TestCollectRejectsUnknownCollector(t *testing.T) {
	for _, kind := range []string{"explicit report directory", "temporary report directory"} {
		t.Run(kind, func(t *testing.T) {
			directory := fixtureRoot(t)
			t.Setenv("TMPDIR", directory.Name())
			t.Setenv("XQUIK_COVERAGE_DIR", "")
			if kind == "explicit report directory" {
				t.Setenv("XQUIK_COVERAGE_DIR", directory.Name())
			}
			var output bytes.Buffer
			err := coveragegate.Collect("--unknown", &output)
			if err == nil || err.Error() != "unknown native collector: --unknown" {
				t.Fatalf("unknown collector result: %v\n%s", err, output.String())
			}
			reported := strings.TrimSuffix(strings.TrimPrefix(output.String(), "Native coverage reports: "), "\n")
			if reported != directory.Name() && filepath.Dir(reported) != directory.Name() {
				t.Fatalf("unknown collector escaped its report root: %q", reported)
			}
			entries, err := os.ReadDir(reported)
			if err != nil || len(entries) != 0 {
				t.Fatalf("unknown collector executed native work: %v, %v", entries, err)
			}
		})
	}
}

func TestCollectPreservesInvalidTemporaryRoot(t *testing.T) {
	directory := fixtureRoot(t)
	if err := directory.WriteFile("temp-root", []byte("retained temporary root obstacle\n"), 0o600); err != nil {
		t.Fatal(err)
	}
	t.Setenv("TMPDIR", filepath.Join(directory.Name(), "temp-root"))
	t.Setenv("XQUIK_COVERAGE_DIR", "")
	var output bytes.Buffer
	if err := coveragegate.Collect("--collect-statements", &output); err == nil || output.Len() != 0 {
		t.Fatalf("invalid temporary root reached collection: %v\n%s", err, output.String())
	}
	if data, err := directory.ReadFile("temp-root"); err != nil || string(data) != "retained temporary root obstacle\n" {
		t.Fatalf("temporary root obstacle changed: %q, %v", data, err)
	}
}

func TestCollectStatementsPreservesMalformedChildEvidence(t *testing.T) {
	directory, _ := baselineFixture(t)
	reports := fixtureRoot(t)
	const child = "mode: count\nexample.test/coveragefixture/value.go:2.23,3.11 1 1\n"
	if err := reports.WriteFile("child-coverage.out", []byte(child), 0o600); err != nil {
		t.Fatal(err)
	}
	var output bytes.Buffer
	if err := coveragegate.CollectStatements(directory.Name(), reports.Name(), &output); err == nil || !strings.Contains(err.Error(), "set instrumentation") {
		t.Fatalf("invalid child instrumentation passed: %v\n%s", err, output.String())
	}
	if data, err := reports.ReadFile("child-coverage.out"); err != nil || string(data) != child {
		t.Fatalf("malformed child evidence changed: %q, %v", data, err)
	}
	if data, err := reports.ReadFile("unit-coverage.out"); err != nil || len(data) == 0 {
		t.Fatalf("native unit evidence lost: %d bytes, %v", len(data), err)
	}
	if _, err := reports.Stat("coverage.xml"); !os.IsNotExist(err) {
		t.Fatalf("malformed child evidence published XML: %v", err)
	}
}

func TestRecordBaselineRejectsReportsChangedDuringCollection(t *testing.T) {
	directory, _ := baselineFixture(t)
	const source = `package fixture
import ("os"; "path/filepath"; "testing")
func TestChangeCollectedReport(t *testing.T) {
 path := filepath.Join(os.Getenv("XQUIK_COVERAGE_DIR"), "coverage.xml")
 if _, err := os.Stat(path); os.IsNotExist(err) { return } else if err != nil { t.Fatal(err) }
 if err := os.WriteFile(path, []byte("malformed native line report\n"), 0600); err != nil { t.Fatal(err) }
}
`
	if err := directory.WriteFile("changed_report_test.go", []byte(source), 0o600); err != nil {
		t.Fatal(err)
	}
	commitFixture(t, directory)
	revision := strings.TrimSpace(string(fixtureCommand(t, directory, exec.Command("git", "rev-parse", "HEAD"))))
	var output bytes.Buffer
	_, err := coveragegate.Record(directory.Name(), &output)
	if err == nil || !strings.Contains(err.Error(), "read line and function coverage") {
		t.Fatalf("changed native report accepted: %v\n%s", err, output.String())
	}
	path := filepath.Join(".git", "coverage-baselines", "go-packages-v3", revision)
	if data, err := directory.ReadFile(filepath.Join(path, "coverage.xml")); err != nil || string(data) != "malformed native line report\n" {
		t.Fatalf("changed native evidence was lost: %q, %v", data, err)
	}
	for _, filename := range []string{"coverage.out", "unit-branch-coverage.out"} {
		if data, err := directory.ReadFile(filepath.Join(path, filename)); err != nil || len(data) == 0 {
			t.Fatalf("native evidence was lost: %s, %d bytes, %v", filename, len(data), err)
		}
	}
	if _, err := directory.Stat(filepath.Join(path, "report.json")); !os.IsNotExist(err) {
		t.Fatalf("changed native report published a baseline: %v", err)
	}
}

func TestRecordBaselineRejectsMissingNativeBranchMetric(t *testing.T) {
	directory, _ := baselineFixture(t)
	for name, source := range map[string]string{
		"value.go":      "package fixture\nfunc Value(v int) int { return v }\n",
		"value_test.go": "package fixture\nimport \"testing\"\nfunc TestValue(t *testing.T) { if Value(1) != 1 { t.Fatal(\"wrong native value\") } }\n",
	} {
		if err := directory.WriteFile(name, []byte(source), 0o600); err != nil {
			t.Fatal(err)
		}
	}
	commitFixture(t, directory)
	revision := strings.TrimSpace(string(fixtureCommand(t, directory, exec.Command("git", "rev-parse", "HEAD"))))
	var output bytes.Buffer
	_, err := coveragegate.Record(directory.Name(), &output)
	if err == nil || !strings.Contains(err.Error(), "invalid branches counts in report 0") {
		t.Fatalf("missing branch metric accepted: %v\n%s", err, output.String())
	}
	path := filepath.Join(".git", "coverage-baselines", "go-packages-v3", revision)
	reports, err := directory.OpenRoot(path)
	if err != nil {
		t.Fatal(err)
	}
	defer reports.Close()
	metrics, err := coveragegate.Measure(reports.FS())
	if err != nil || metrics["branches"] != (coveragegate.Metric{Covered: 0, Total: 0}) {
		t.Fatalf("native branch metric = %v, %v", metrics["branches"], err)
	}
	for _, filename := range []string{"report.json", ".baseline-pending.json"} {
		if _, err := reports.Stat(filename); !os.IsNotExist(err) {
			t.Fatalf("missing branch metric published evidence: %s, %v", filename, err)
		}
	}
}
