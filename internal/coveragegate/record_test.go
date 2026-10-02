// Copyright the Xquik contributors.
// SPDX-License-Identifier: Apache-2.0

package coveragegate_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	"github.com/Xquik-dev/x-twitter-scraper-go/internal/coveragegate"
)

func fixtureRoot(t *testing.T) *os.Root {
	t.Helper()
	root, err := os.OpenRoot(t.TempDir())
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if err := root.Close(); err != nil {
			t.Fatal(err)
		}
	})
	return root
}

func fixtureCommand(t *testing.T, root *os.Root, command *exec.Cmd) []byte {
	t.Helper()
	command.Dir = root.Name()
	output, err := command.CombinedOutput()
	if err != nil {
		t.Fatalf("%v: %v\n%s", command.Args, err, output)
	}
	return output
}

func commitFixture(t *testing.T, root *os.Root) {
	t.Helper()
	fixtureCommand(t, root, exec.Command("git", "add", "--", "."))
	fixtureCommand(t, root, exec.Command("git", "-c", "user.name=Coverage test", "-c", "user.email=coverage@example.test", "-c", "commit.gpgsign=false", "commit", "-qm", "coverage fixture"))
}

func baselineFixture(t *testing.T) (*os.Root, string) {
	t.Helper()
	directory := fixtureRoot(t)
	write := func(name, contents string) {
		t.Helper()
		if err := directory.WriteFile(name, []byte(contents), 0o600); err != nil {
			t.Fatal(err)
		}
	}
	write("go.mod", "module example.test/coveragefixture\n\ngo 1.27.1\n")
	write("value.go", "package fixture\nfunc Value(v int) int {\n if v > 0 { return 1 }\n return -1\n}\n")
	write("value_test.go", "package fixture\nimport \"testing\"\nfunc TestValue(t *testing.T) {\n for _, v := range []int{1, 0} {\n  want := -1; if v > 0 { want = 1 }; if got := Value(v); got != want { t.Fatalf(\"got %d, want %d\", got, want) }\n }\n}\n")
	fixtureCommand(t, directory, exec.Command("git", "init", "-q", "--initial-branch=main"))
	commitFixture(t, directory)
	output := fixtureCommand(t, directory, exec.Command("git", "rev-parse", "HEAD"))
	return directory, strings.TrimSpace(string(output))
}

func TestRecordBaselineCollectsRevisionBoundMetrics(t *testing.T) {
	directory, revision := baselineFixture(t)
	var output bytes.Buffer
	measured, err := coveragegate.Record(directory.Name(), &output)
	if err != nil {
		t.Fatalf("record real native coverage: %v\n%s", err, output.String())
	}
	data, err := directory.ReadFile(filepath.Join(".git", "coverage-baselines", "go-packages-v3", revision, "report.json"))
	if err != nil {
		t.Fatal(err)
	}
	var report coveragegate.Report
	if err := json.Unmarshal(data, &report); err != nil {
		t.Fatal(err)
	}
	if report.Revision != revision || report.Scope != runtime.Version()+"/"+runtime.GOOS+"/"+runtime.GOARCH+"/go-packages-v3" {
		t.Fatalf("source binding = %s, %s", report.Revision, report.Scope)
	}
	if measured.Revision != report.Revision || measured.Scope != report.Scope || len(measured.Total) != 4 {
		t.Fatalf("returned source binding = %+v", measured)
	}
	for _, name := range []string{"statements", "branches", "functions", "lines"} {
		metric, exists := report.Total[name]
		if !exists || metric.Total <= 0 || metric.Covered != metric.Total || metric.Skipped != 0 || measured.Total[name] != metric {
			t.Fatalf("%s metric = %+v, exists = %v", name, metric, exists)
		}
	}
	if report.Total["statements"].Total != 3 || report.Total["functions"].Total != 1 || report.Total["branches"].Total != 1 || report.Total["lines"].Total != 2 {
		t.Fatalf("fixture counts = %+v", report.Total)
	}
	for _, name := range []string{"coverage.out", "coverage.xml", "branches.json", "branches.cover"} {
		data, err := directory.ReadFile(filepath.Join(".git", "coverage-baselines", "go-packages-v3", revision, name))
		if err != nil || len(data) == 0 {
			t.Fatalf("retained %s = %d bytes, %v", name, len(data), err)
		}
	}
}

func childBaselineFixture(t *testing.T) *os.Root {
	t.Helper()
	directory, _ := baselineFixture(t)
	for _, name := range []string{"cmd/value", "untested"} {
		if err := directory.MkdirAll(name, 0o700); err != nil {
			t.Fatal(err)
		}
	}
	for name, source := range map[string]string{
		"value.go":           "package fixture\nfunc Value(v int) int {\n if v > 0 { return 1 }\n return -1\n}\nfunc UnitOnly() int { return 7 }\n",
		"cmd/value/main.go":  "package main\nimport (\"fmt\"; fixture \"example.test/coveragefixture\")\nfunc main() { fmt.Println(fixture.Value(-1)) }\n",
		"untested/unused.go": "package untested\nfunc Unused(v int) int {\n if v > 0 { return 9 }\n return 0\n}\n",
		"value_test.go": `package fixture
import ("os"; "os/exec"; "testing")
func TestUnitBehavior(t *testing.T) {
 if got := Value(1); got != 1 { t.Fatalf("positive value = %d", got) }
 if got := UnitOnly(); got != 7 { t.Fatalf("unit value = %d", got) }
}
func TestCommandBehavior(t *testing.T) {
 build := exec.Command("go", "build", "-cover", "-covermode=set", "-coverpkg=./...", "-o", ".git/value", "./cmd/value")
 if output, err := build.CombinedOutput(); err != nil { t.Fatalf("build command: %v\n%s", err, output) }
 counters := os.Getenv("XQUIK_GO_COVERAGE_DIR")
 if counters == "" { counters = ".git/child-native-coverage" }
 if err := os.MkdirAll(counters, 0700); err != nil { t.Fatal(err) }
 command := exec.Command("./.git/value")
 command.Env = append(os.Environ(), "GOCOVERDIR="+counters)
 output, err := command.CombinedOutput()
 if err != nil || string(output) != "-1\n" { t.Fatalf("command result: %q, %v", output, err) }
}
`,
	} {
		if err := directory.WriteFile(name, []byte(source), 0o600); err != nil {
			t.Fatal(err)
		}
	}
	commitFixture(t, directory)
	return directory
}

func TestRecordBaselineIncludesUnitChildAndUnexecutedSource(t *testing.T) {
	directory := childBaselineFixture(t)
	var output bytes.Buffer
	report, err := coveragegate.Record(directory.Name(), &output)
	if err != nil {
		t.Fatalf("record complete native execution: %v\n%s", err, output.String())
	}
	for name, want := range map[string]coveragegate.Metric{
		"statements": {Covered: 5, Total: 8},
		"branches":   {Covered: 1, Total: 2},
		"functions":  {Covered: 3, Total: 4},
		"lines":      {Covered: 4, Total: 6},
	} {
		if got := report.Total[name]; got != want {
			t.Errorf("%s = %+v, want %+v", name, got, want)
		}
	}
}

func TestRecordBaselineRejectsDifferentNativeToolchain(t *testing.T) {
	directory, revision := baselineFixture(t)
	otherOS := "freebsd"
	if runtime.GOOS == otherOS {
		otherOS = "linux"
	}
	t.Setenv("GOOS", otherOS)
	var output bytes.Buffer
	_, err := coveragegate.Record(directory.Name(), &output)
	if err == nil || !strings.Contains(err.Error(), "native toolchain does not match recorder") {
		t.Fatalf("different native target: %v\n%s", err, output.String())
	}
	if _, err := directory.Stat(filepath.Join(".git", "coverage-baselines", "go-packages-v3", revision)); !os.IsNotExist(err) {
		t.Fatalf("different native target reserved evidence: %v", err)
	}
}

func TestRecordBaselinePinsCollectorToolchain(t *testing.T) {
	directory, _ := baselineFixture(t)
	const source = "package fixture\nimport (\"os\"; \"runtime\"; \"testing\")\nfunc TestCollectorToolchain(t *testing.T) { if got := os.Getenv(\"GOTOOLCHAIN\"); got != runtime.Version() { t.Fatalf(\"collector toolchain = %q, want %q\", got, runtime.Version()) } }\n"
	if err := directory.WriteFile("toolchain_test.go", []byte(source), 0o600); err != nil {
		t.Fatal(err)
	}
	commitFixture(t, directory)
	t.Setenv("GOTOOLCHAIN", "auto")
	var output bytes.Buffer
	if _, err := coveragegate.Record(directory.Name(), &output); err != nil {
		t.Fatalf("collect with a pinned native toolchain: %v\n%s", err, output.String())
	}
}

func TestRecordBaselineRejectsSourceChangesDuringCollection(t *testing.T) {
	for _, test := range []struct {
		name, source string
	}{
		{"tracked source", "package fixture\nimport (\"os\"; \"testing\")\nfunc TestChange(t *testing.T) { f, err := os.OpenFile(\"value.go\", os.O_WRONLY|os.O_APPEND, 0600); if err != nil { t.Fatal(err) }; if _, err := f.WriteString(\"// changed during collection\\n\"); err != nil { t.Fatal(err) }; if err := f.Close(); err != nil { t.Fatal(err) } }\n"},
		{"clean new revision", "package fixture\nimport (\"os/exec\"; \"testing\")\nfunc TestChange(t *testing.T) { c := exec.Command(\"git\", \"-c\", \"user.name=Coverage test\", \"-c\", \"user.email=coverage@example.test\", \"-c\", \"commit.gpgsign=false\", \"commit\", \"--allow-empty\", \"-qm\", \"changed while collecting\"); if data, err := c.CombinedOutput(); err != nil { t.Fatalf(\"%v: %s\", err, data) } }\n"},
	} {
		t.Run(test.name, func(t *testing.T) {
			directory, _ := baselineFixture(t)
			if err := directory.WriteFile("mutation_test.go", []byte(test.source), 0o600); err != nil {
				t.Fatal(err)
			}
			commitFixture(t, directory)
			revision := fixtureCommand(t, directory, exec.Command("git", "rev-parse", "HEAD"))

			var output bytes.Buffer
			_, err := coveragegate.Record(directory.Name(), &output)
			if err == nil || !strings.Contains(err.Error(), "source changed during collection") {
				t.Fatalf("changed source: %v\n%s", err, output.String())
			}
			if _, err := directory.Stat(filepath.Join(".git", "coverage-baselines", "go-packages-v3", strings.TrimSpace(string(revision)), "report.json")); !os.IsNotExist(err) {
				t.Fatalf("changed source baseline exists: %v", err)
			}
		})
	}
}

func TestRecordBaselineRejectsUncommittedSource(t *testing.T) {
	directory, revision := baselineFixture(t)
	if err := directory.WriteFile("untracked.go", []byte("package fixture\n"), 0o600); err != nil {
		t.Fatal(err)
	}
	var output bytes.Buffer
	_, err := coveragegate.Record(directory.Name(), &output)
	if err == nil || !strings.Contains(err.Error(), "requires a clean checkout") {
		t.Fatalf("dirty source: %v\n%s", err, output.String())
	}
	if _, err := directory.Stat(filepath.Join(".git", "coverage-baselines", "go-packages-v3", revision, "report.json")); !os.IsNotExist(err) {
		t.Fatalf("dirty source baseline exists: %v", err)
	}
}

func TestRecordBaselinePreservesExistingEvidence(t *testing.T) {
	directory, revision := baselineFixture(t)
	baseline := filepath.Join(".git", "coverage-baselines", "go-packages-v3", revision)
	if err := directory.MkdirAll(baseline, 0o700); err != nil {
		t.Fatal(err)
	}
	filename := filepath.Join(baseline, "report.json")
	const original = "retained baseline evidence\n"
	if err := directory.WriteFile(filename, []byte(original), 0o600); err != nil {
		t.Fatal(err)
	}
	var output bytes.Buffer
	_, err := coveragegate.Record(directory.Name(), &output)
	if err == nil || !strings.Contains(err.Error(), "baseline already exists") {
		t.Fatalf("existing evidence: %v\n%s", err, output.String())
	}
	data, err := directory.ReadFile(filename)
	if err != nil || string(data) != original {
		t.Fatalf("evidence changed: %s, %v", data, err)
	}
}

func buildCollectorCommand(t *testing.T, directory *os.Root) {
	t.Helper()
	source, err := filepath.Abs("../..")
	if err != nil {
		t.Fatal(err)
	}
	workspace := fixtureRoot(t)
	contents := fmt.Sprintf("go 1.27.1\nuse (\n%q\n%q\n)\n", source, directory.Name())
	if err := workspace.WriteFile("go.work", []byte(contents), 0o600); err != nil {
		t.Fatal(err)
	}
	build := exec.Command("go", "build", "-cover", "-covermode=set", "-coverpkg=github.com/Xquik-dev/x-twitter-scraper-go/...", "-o", ".git/coverage-gain", "github.com/Xquik-dev/x-twitter-scraper-go/cmd/coverage-gain")
	build.Env = append(os.Environ(), "GOWORK="+filepath.Join(workspace.Name(), "go.work"))
	fixtureCommand(t, directory, build)
}

func TestRecordBaselineSurvivesComparison(t *testing.T) {
	directory, revision := baselineFixture(t)
	buildCollectorCommand(t, directory)
	command := exec.Command("./.git/coverage-gain", "--record")
	command.Dir = directory.Name()
	command.Env = append(os.Environ(), "GOCOVERDIR="+os.Getenv("XQUIK_GO_COVERAGE_DIR"))
	var stderr bytes.Buffer
	command.Stderr = &stderr
	output, err := command.Output()
	if err != nil {
		t.Fatalf("record command: %v\n%s", err, stderr.String())
	}

	var published coveragegate.Report
	if err := json.Unmarshal(output, &published); err != nil || published.Revision != revision {
		t.Fatalf("record response: %s, %v", output, err)
	}
	baseline := filepath.Join(".git", "coverage-baselines", "go-packages-v3", revision)
	filename := filepath.Join(baseline, "report.json")
	original, err := directory.ReadFile(filename)
	if err != nil {
		t.Fatal(err)
	}
	if err := directory.WriteFile("README.md", []byte("Coverage fixture\n"), 0o600); err != nil {
		t.Fatal(err)
	}
	commitFixture(t, directory)
	if err := directory.Symlink(filepath.Join("coverage-baselines", "go-packages-v3", revision), ".git/baseline"); err != nil {
		t.Fatal(err)
	}
	for _, alias := range []string{"regular summary", "relative symlink", "hard link"} {
		t.Run(alias, func(t *testing.T) {
			summary := filepath.Join(baseline, "summary.json")
			if err := directory.Remove(summary); err != nil && !os.IsNotExist(err) {
				t.Fatal(err)
			}
			if err := directory.WriteFile(filename, original, 0o600); err != nil {
				t.Fatal(err)
			}
			switch alias {
			case "relative symlink":
				if err := directory.Symlink("report.json", summary); err != nil {
					t.Fatal(err)
				}
			case "hard link":
				if err := directory.Link(filename, summary); err != nil {
					t.Fatal(err)
				}
			}
			comparison := exec.Command("./.git/coverage-gain", ".git/baseline")
			comparison.Env = append(os.Environ(), "GOCOVERDIR="+os.Getenv("XQUIK_GO_COVERAGE_DIR"))
			fixtureCommand(t, directory, comparison)
			retained, err := directory.ReadFile(filename)
			if err != nil || !bytes.Equal(retained, original) {
				t.Fatalf("comparison changed verified baseline: %s, %v", retained, err)
			}
		})
	}
}

func TestRecordBaselineConfinesNestedChildCounters(t *testing.T) {
	outer := fixtureRoot(t)
	t.Setenv("XQUIK_GO_COVERAGE_DIR", outer.Name())
	directory := childBaselineFixture(t)
	var output bytes.Buffer
	if _, err := coveragegate.Record(directory.Name(), &output); err != nil {
		t.Fatalf("record nested native coverage: %v\n%s", err, output.String())
	}
	entries, err := fs.ReadDir(outer.FS(), ".")
	if err != nil || len(entries) != 0 {
		t.Fatalf("nested source counters escaped into the caller's collection: %d files, %v", len(entries), err)
	}
}
