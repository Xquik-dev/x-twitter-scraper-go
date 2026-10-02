// Copyright the Xquik contributors.
// SPDX-License-Identifier: Apache-2.0

package coveragegate_test

import (
	"bytes"
	"io"
	"os"
	"os/exec"
	"testing"

	"github.com/Xquik-dev/x-twitter-scraper-go/internal/coveragegate"
)

func TestCollectorCommandReportsNativeFixtureCoverage(t *testing.T) {
	for _, test := range []struct {
		name        string
		command     *exec.Cmd
		counterpart func(string, string, io.Writer) error
	}{
		{"statements", exec.Command("./.git/coverage-gain", "--collect-statements"), coveragegate.CollectBranches},
		{"branches", exec.Command("./.git/coverage-gain", "--collect-branches"), coveragegate.CollectStatements},
	} {
		t.Run(test.name, func(t *testing.T) {
			directory, _ := baselineFixture(t)
			buildCollectorCommand(t, directory)
			reports := fixtureRoot(t)
			var output bytes.Buffer
			if err := test.counterpart(directory.Name(), reports.Name(), &output); err != nil {
				t.Fatalf("collect counterpart: %v\n%s", err, output.String())
			}
			test.command.Env = append(os.Environ(), "XQUIK_COVERAGE_DIR="+reports.Name(), "GOCOVERDIR="+os.Getenv("XQUIK_GO_COVERAGE_DIR"))
			fixtureCommand(t, directory, test.command)
			metrics, err := coveragegate.Measure(reports.FS())
			if err != nil {
				t.Fatal(err)
			}
			for name, want := range map[string]coveragegate.Metric{
				"statements": {Covered: 3, Total: 3},
				"branches":   {Covered: 1, Total: 1},
				"functions":  {Covered: 1, Total: 1},
				"lines":      {Covered: 2, Total: 2},
			} {
				if got := metrics[name]; got != want {
					t.Errorf("%s = %+v, want %+v", name, got, want)
				}
			}
		})
	}
}

func TestRecordBaselineIncludesBranchesExecutedOnlyByChildCommands(t *testing.T) {
	directory := childBaselineFixture(t)
	const source = "package fixture\nfunc Value(v int) int {\n if v < 0 { return -1 }\n return 1\n}\nfunc UnitOnly() int { return 7 }\n"
	if err := directory.WriteFile("value.go", []byte(source), 0o600); err != nil {
		t.Fatal(err)
	}
	commitFixture(t, directory)
	var output bytes.Buffer
	report, err := coveragegate.Record(directory.Name(), &output)
	if err != nil {
		t.Fatalf("record child branch coverage: %v\n%s", err, output.String())
	}
	want := coveragegate.Metric{Covered: 1, Total: 2}
	if got := report.Total["branches"]; got != want {
		t.Fatalf("child-only branch = %+v, want %+v", got, want)
	}
}

func TestRecordBaselineIncludesBranchesExecutedByOtherPackages(t *testing.T) {
	directory, _ := baselineFixture(t)
	if err := directory.Remove("value_test.go"); err != nil {
		t.Fatal(err)
	}
	if err := directory.Mkdir("caller", 0o700); err != nil {
		t.Fatal(err)
	}
	if err := directory.WriteFile("caller/value.go", []byte("package caller\nimport fixture \"example.test/coveragefixture\"\nfunc Value(v int) int { return fixture.Value(v) }\n"), 0o600); err != nil {
		t.Fatal(err)
	}
	const source = `package caller
import "testing"
func TestValue(t *testing.T) {
 if got := Value(1); got != 1 { t.Fatalf("positive value = %d", got) }
 if got := Value(0); got != -1 { t.Fatalf("zero value = %d", got) }
}
`
	if err := directory.WriteFile("caller/value_test.go", []byte(source), 0o600); err != nil {
		t.Fatal(err)
	}
	commitFixture(t, directory)
	var output bytes.Buffer
	report, err := coveragegate.Record(directory.Name(), &output)
	if err != nil {
		t.Fatalf("record cross-package branch coverage: %v\n%s", err, output.String())
	}
	want := coveragegate.Metric{Covered: 1, Total: 1}
	if got := report.Total["branches"]; got != want {
		t.Fatalf("cross-package branch = %+v, want %+v", got, want)
	}
}

func TestRecordBaselineIncludesInternalAndExternalTestExecution(t *testing.T) {
	directory, _ := baselineFixture(t)
	const value = "package fixture\nfunc Value(v int) int {\n if v < 0 { return -1 }\n return 1\n}\n"
	const internalTest = `package fixture
import "testing"
func TestPositiveValue(t *testing.T) {
 if got := Value(1); got != 1 { t.Fatalf("positive value = %d", got) }
}
`
	const externalTest = `package fixture_test
import (
 fixture "example.test/coveragefixture"
 "testing"
)
func TestNegativeValue(t *testing.T) {
 if got := fixture.Value(-1); got != -1 { t.Fatalf("negative value = %d", got) }
}
`
	for name, contents := range map[string]string{
		"value.go":         value,
		"value_test.go":    internalTest,
		"external_test.go": externalTest,
	} {
		if err := directory.WriteFile(name, []byte(contents), 0o600); err != nil {
			t.Fatal(err)
		}
	}
	commitFixture(t, directory)
	var output bytes.Buffer
	report, err := coveragegate.Record(directory.Name(), &output)
	if err != nil {
		t.Fatalf("record internal and external execution: %v\n%s", err, output.String())
	}
	for name, want := range map[string]coveragegate.Metric{
		"statements": {Covered: 3, Total: 3},
		"branches":   {Covered: 1, Total: 1},
		"functions":  {Covered: 1, Total: 1},
		"lines":      {Covered: 2, Total: 2},
	} {
		if got := report.Total[name]; got != want {
			t.Errorf("%s = %+v, want %+v", name, got, want)
		}
	}
}
