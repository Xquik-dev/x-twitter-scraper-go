// Copyright the Xquik contributors.
// SPDX-License-Identifier: Apache-2.0

package coveragegate_test

import (
	"bytes"
	"encoding/json"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	"github.com/Xquik-dev/x-twitter-scraper-go/internal/coveragegate"
)

func comparisonCommandFixture(t *testing.T) (*os.Root, *os.Root) {
	t.Helper()
	directory, revision := baselineFixture(t)
	buildCollectorCommand(t, directory)
	baseline := coveragegate.Report{
		Revision: revision,
		Scope:    runtime.Version() + "/" + runtime.GOOS + "/" + runtime.GOARCH + "/go-packages-v3",
		Total: map[string]coveragegate.Metric{
			"statements": {Covered: 1, Total: 1},
			"branches":   {Covered: 1, Total: 1},
			"functions":  {Covered: 1, Total: 1},
			"lines":      {Covered: 1, Total: 1},
		},
	}
	path := filepath.Join(".git", "coverage-baselines", "go-packages-v3", revision)
	if err := directory.MkdirAll(path, 0o700); err != nil {
		t.Fatal(err)
	}
	data, err := json.Marshal(baseline)
	if err != nil {
		t.Fatal(err)
	}
	if err := directory.WriteFile(filepath.Join(path, "report.json"), data, 0o600); err != nil {
		t.Fatal(err)
	}
	if err := directory.WriteFile("README.md", []byte("Coverage fixture\n"), 0o600); err != nil {
		t.Fatal(err)
	}
	commitFixture(t, directory)
	if err := directory.Mkdir(".git/reports", 0o700); err != nil {
		t.Fatal(err)
	}
	reports, err := directory.OpenRoot(".git/reports")
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if err := reports.Close(); err != nil {
			t.Fatal(err)
		}
	})
	for name, contents := range map[string]string{
		"coverage.out":   "mode: set\nexample/a.go:1.1,2.1 1 1\n",
		"coverage.xml":   `<coverage lines-covered="1" lines-valid="1"><packages><package name="example"><classes><class filename="a.go"><methods><method line-rate="1"/></methods></class></classes></package></packages></coverage>`,
		"branches.json":  `{"packages":[{"import_path":"example","files":[{"name":"a.go","blocks":[{"idx":1,"line":1,"col":1,"end_line":2,"end_col":1}],"branches":[{"block_idx":1}]}]}]}`,
		"branches.cover": "mode: set\nexample/a.go:1.1,2.1 1 1\n",
	} {
		if err := reports.WriteFile(name, []byte(contents), 0o600); err != nil {
			t.Fatal(err)
		}
	}
	return directory, reports
}

func runCoverageCommand(t *testing.T, directory *os.Root, command *exec.Cmd) (string, string, error) {
	t.Helper()
	command.Dir = directory.Name()
	command.Env = append(os.Environ(), "GOCOVERDIR="+os.Getenv("XQUIK_GO_COVERAGE_DIR"))
	var stdout, stderr bytes.Buffer
	command.Stdout, command.Stderr = &stdout, &stderr
	err := command.Run()
	return stdout.String(), stderr.String(), err
}

func requireCommandFailure(t *testing.T, stdout, stderr string, err error, diagnostic string) {
	t.Helper()
	if exit, ok := err.(*exec.ExitError); !ok || exit.ExitCode() != 1 || stdout != "" || !strings.Contains(stderr, diagnostic) {
		t.Fatalf("command failure: %v, stdout=%q, stderr=%q; need %q", err, stdout, stderr, diagnostic)
	}
}

func TestCoverageCommandRejectsMissingArguments(t *testing.T) {
	directory, _ := baselineFixture(t)
	buildCollectorCommand(t, directory)
	stdout, stderr, err := runCoverageCommand(t, directory, exec.Command("./.git/coverage-gain"))
	requireCommandFailure(t, stdout, stderr, err, "usage: coverage-gain")
}

func TestCoverageCommandRejectsInvalidReportDirectory(t *testing.T) {
	directory, reports := comparisonCommandFixture(t)
	if err := reports.WriteFile("not-a-directory", []byte("retained report obstacle\n"), 0o600); err != nil {
		t.Fatal(err)
	}
	stdout, stderr, err := runCoverageCommand(t, directory, exec.Command("./.git/coverage-gain", ".git/reports/not-a-directory"))
	requireCommandFailure(t, stdout, stderr, err, "not a directory")
	if data, err := reports.ReadFile("not-a-directory"); err != nil || string(data) != "retained report obstacle\n" {
		t.Fatalf("report obstacle changed: %q, %v", data, err)
	}
}

func TestCoverageCommandPreservesOccupiedPendingSummary(t *testing.T) {
	directory, reports := comparisonCommandFixture(t)
	if err := reports.WriteFile(".summary-pending.json", []byte("retained summary evidence\n"), 0o600); err != nil {
		t.Fatal(err)
	}
	stdout, stderr, err := runCoverageCommand(t, directory, exec.Command("./.git/coverage-gain", ".git/reports"))
	requireCommandFailure(t, stdout, stderr, err, ".summary-pending.json")
	if data, err := reports.ReadFile(".summary-pending.json"); err != nil || string(data) != "retained summary evidence\n" {
		t.Fatalf("pending summary changed: %q, %v", data, err)
	}
	if _, err := reports.Stat("summary.json"); !os.IsNotExist(err) {
		t.Fatalf("occupied pending summary published a replacement: %v", err)
	}
}

func TestCoverageCommandReportsCoverageRegression(t *testing.T) {
	directory, reports := comparisonCommandFixture(t)
	data, err := reports.ReadFile("coverage.xml")
	if err != nil {
		t.Fatal(err)
	}
	data = bytes.Replace(data, []byte(`lines-covered="1"`), []byte(`lines-covered="0"`), 1)
	if err := reports.WriteFile("coverage.xml", data, 0o600); err != nil {
		t.Fatal(err)
	}
	stdout, stderr, err := runCoverageCommand(t, directory, exec.Command("./.git/coverage-gain", ".git/reports"))
	requireCommandFailure(t, stdout, stderr, err, "lines coverage")
	data, err = reports.ReadFile("summary.json")
	if err != nil {
		t.Fatal(err)
	}
	var summary coveragegate.Report
	if err := json.Unmarshal(data, &summary); err != nil || summary.Total["lines"] != (coveragegate.Metric{Covered: 0, Total: 1}) {
		t.Fatalf("regression summary: %s, %v", data, err)
	}
}

func TestCoverageCommandRejectsUnavailableGitSource(t *testing.T) {
	source, _ := baselineFixture(t)
	buildCollectorCommand(t, source)
	directory := fixtureRoot(t)
	if err := directory.Mkdir(".git", 0o700); err != nil {
		t.Fatal(err)
	}
	binary, err := source.ReadFile(".git/coverage-gain")
	if err != nil {
		t.Fatal(err)
	}
	if err := directory.WriteFile(".git/coverage-gain", binary, 0o700); err != nil {
		t.Fatal(err)
	}
	stdout, stderr, err := runCoverageCommand(t, directory, exec.Command("./.git/coverage-gain", ".git/reports"))
	requireCommandFailure(t, stdout, stderr, err, "exit status 128")
	if _, err := directory.Stat(".git/reports"); !os.IsNotExist(err) {
		t.Fatalf("non-Git source published reports: %v", err)
	}
}

func TestCoverageCommandRejectsUnavailableParentRevision(t *testing.T) {
	directory, _ := baselineFixture(t)
	buildCollectorCommand(t, directory)
	stdout, stderr, err := runCoverageCommand(t, directory, exec.Command("./.git/coverage-gain", ".git/reports"))
	requireCommandFailure(t, stdout, stderr, err, "exit status 128")
	if _, err := directory.Stat(".git/reports"); !os.IsNotExist(err) {
		t.Fatalf("missing parent revision published reports: %v", err)
	}
	if _, err := directory.Stat(".git/coverage-baselines"); !os.IsNotExist(err) {
		t.Fatalf("missing parent revision reserved a baseline: %v", err)
	}
}

func TestCoverageCommandPreservesOutwardSummarySymlink(t *testing.T) {
	directory, reports := comparisonCommandFixture(t)
	outside := fixtureRoot(t)
	const evidence = "retained external summary evidence\n"
	if err := outside.WriteFile("summary.json", []byte(evidence), 0o600); err != nil {
		t.Fatal(err)
	}
	target := filepath.Join(outside.Name(), "summary.json")
	if err := reports.Symlink(target, "summary.json"); err != nil {
		t.Fatal(err)
	}
	stdout, stderr, err := runCoverageCommand(t, directory, exec.Command("./.git/coverage-gain", ".git/reports"))
	requireCommandFailure(t, stdout, stderr, err, "escapes")
	if data, err := outside.ReadFile("summary.json"); err != nil || string(data) != evidence {
		t.Fatalf("external summary evidence changed: %q, %v", data, err)
	}
	if retained, err := reports.Readlink("summary.json"); err != nil || retained != target {
		t.Fatalf("summary symlink changed: %q, %v", retained, err)
	}
	if _, err := reports.Stat(".summary-pending.json"); !os.IsNotExist(err) {
		t.Fatalf("outward summary symlink reserved a replacement: %v", err)
	}
}
