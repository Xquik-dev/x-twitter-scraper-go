// Copyright the Xquik contributors.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"encoding/json"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	"github.com/Xquik-dev/x-twitter-scraper-go/internal/coveragegate"
)

func testRoot(t *testing.T) *os.Root {
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

func TestRunAgainstGitParent(t *testing.T) {
	fixture := testRoot(t)
	t.Chdir(fixture.Name())
	directory := testRoot(t)
	if err := run(directory.Name()); err == nil {
		t.Fatal("non-repository passed")
	}
	gitCommand := func(command *exec.Cmd) {
		t.Helper()
		if data, err := command.CombinedOutput(); err != nil {
			t.Fatalf("git %v: %v\n%s", command.Args, err, data)
		}
	}
	write := func(path, data string) {
		t.Helper()
		if err := fixture.WriteFile(path, []byte(data), 0o600); err != nil {
			t.Fatal(err)
		}
	}
	if err := checkLOC("HEAD"); err == nil {
		t.Fatal("LOC check outside Git passed")
	}
	gitCommand(exec.Command("git", "init", "-q", "--initial-branch=main"))
	gitCommand(exec.Command("git", "config", "user.name", "Coverage test"))
	gitCommand(exec.Command("git", "config", "user.email", "coverage@example.test"))
	gitCommand(exec.Command("git", "config", "commit.gpgsign", "false"))
	if err := run(directory.Name()); err == nil {
		t.Fatal("repository without parent passed")
	}
	const original = "package fixture\nconst one = 1\nconst two = 2\n"
	write("fixture.go", original)
	gitCommand(exec.Command("git", "add", "fixture.go"))
	gitCommand(exec.Command("git", "commit", "-qm", "parent"))
	write("fixture.go", "package fixture\n")
	if err := run(directory.Name()); err == nil || !strings.Contains(err.Error(), "read verified coverage") {
		t.Fatalf("missing baseline error = %v", err)
	}
	parent, err := git(exec.Command("git", "rev-parse", "HEAD"))
	if err != nil {
		t.Fatal(err)
	}
	baseline := coveragegate.Report{Revision: parent, Scope: runtime.Version() + "/" + runtime.GOOS + "/" + runtime.GOARCH + "/go-packages-v3", Total: map[string]coveragegate.Metric{}}
	for _, name := range []string{"statements", "branches", "functions", "lines"} {
		baseline.Total[name] = coveragegate.Metric{Total: 1, Covered: 1}
	}
	path := filepath.Join(".git", "coverage-baselines", "go-packages-v3", parent)
	if err := fixture.MkdirAll(path, 0o700); err != nil {
		t.Fatal(err)
	}
	data, err := json.Marshal(baseline)
	if err != nil {
		t.Fatal(err)
	}
	write(filepath.Join(path, "report.json"), string(data))
	if err := run(directory.Name()); err == nil || !strings.Contains(err.Error(), "coverage.out") {
		t.Fatalf("missing current report error = %v", err)
	}
	const xml = `<coverage lines-covered="1" lines-valid="1"><packages><package name="example"><classes><class filename="a.go"><methods><method line-rate="1"/></methods></class></classes></package></packages></coverage>`
	for name, contents := range map[string]string{
		"coverage.out":   "mode: set\nexample/a.go:1.1,2.1 1 1\n",
		"coverage.xml":   xml,
		"branches.json":  `{"packages":[{"import_path":"example","files":[{"name":"a.go","blocks":[{"idx":1,"line":1,"col":1,"end_line":2,"end_col":1}],"branches":[{"block_idx":1}]}]}]}`,
		"branches.cover": "mode: set\nexample/a.go:1.1,2.1 1 1\n",
	} {
		if err := directory.WriteFile(name, []byte(contents), 0o600); err != nil {
			t.Fatal(err)
		}
	}
	if err := run(directory.Name()); err != nil {
		t.Fatal(err)
	}
	report, err := directory.ReadFile("summary.json")
	if err != nil || !strings.Contains(string(report), `"Covered": 1`) {
		t.Fatalf("summary = %s, %v", report, err)
	}
	t.Run("confines artifact writes", func(t *testing.T) {
		outside := testRoot(t)
		const original = "retained outside report\n"
		if err := outside.WriteFile("outside.json", []byte(original), 0o600); err != nil {
			t.Fatal(err)
		}
		if err := directory.Remove("summary.json"); err != nil {
			t.Fatal(err)
		}
		if err := directory.Symlink(filepath.Join(outside.Name(), "outside.json"), "summary.json"); err != nil {
			t.Fatal(err)
		}
		t.Cleanup(func() {
			if err := directory.Remove("summary.json"); err != nil {
				t.Fatal(err)
			}
		})
		if err := run(directory.Name()); err == nil {
			t.Fatal("accepted an artifact path outside the report directory")
		}
		data, err := outside.ReadFile("outside.json")
		if err != nil || string(data) != original {
			t.Fatalf("outside evidence changed: %s, %v", data, err)
		}
	})
	for _, test := range []struct {
		name, source, extra, failure string
		untracked                    bool
	}{
		{name: "deletion", source: "package fixture\n"},
		{name: "empty file", source: ""},
		{name: "addition", source: original + "const three = 3\n"},
		{name: "unchanged", source: original},
		{name: "replacement", source: strings.ReplaceAll(original, "one", "first")},
		{name: "generated addition", source: "package fixture\n", extra: "xtweet.go"},
		{name: "documentation addition", source: "package fixture\n", extra: "README.md"},
		{name: "new source", source: "package fixture\n", extra: "extra.go"},
		{name: "untracked source", source: "package fixture\n", extra: "extra.go", untracked: true, failure: "requires new files"},
	} {
		t.Run(test.name, func(t *testing.T) {
			write("fixture.go", test.source)
			if test.extra != "" {
				write(test.extra, original)
				if !test.untracked {
					gitCommand(exec.Command("git", "add", "--intent-to-add", "--", "."))
				}
				t.Cleanup(func() {
					gitCommand(exec.Command("git", "reset", "-q", "--", "."))
					if err := fixture.Remove(test.extra); err != nil {
						t.Fatal(err)
					}
				})
			}
			err := checkLOC(parent)
			if test.failure == "" && err != nil || test.failure != "" && (err == nil || !strings.Contains(err.Error(), test.failure)) {
				t.Fatalf("LOC error = %v, want %q", err, test.failure)
			}
		})
	}
	if err := checkLOC("missing-parent"); err == nil {
		t.Fatal("missing parent passed")
	}
	write("fixture.go", "package fixture\n")
	if err := directory.WriteFile("coverage.xml", []byte(strings.Replace(xml, `lines-covered="1"`, `lines-covered="0"`, 1)), 0o600); err != nil {
		t.Fatal(err)
	}
	if err := run(directory.Name()); err == nil || !strings.Contains(err.Error(), "lines coverage") {
		t.Fatalf("coverage regression error = %v", err)
	}
	if err := directory.WriteFile("coverage.xml", []byte(xml), 0o600); err != nil {
		t.Fatal(err)
	}
	gitCommand(exec.Command("git", "add", "fixture.go"))
	gitCommand(exec.Command("git", "commit", "-qm", "candidate"))
	if err := run(directory.Name()); err != nil {
		t.Fatalf("clean candidate must compare its parent: %v", err)
	}
}
