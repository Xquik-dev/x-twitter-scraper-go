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

func TestRunAgainstGitParent(t *testing.T) {
	t.Chdir(t.TempDir())
	directory := t.TempDir()
	if err := run(directory); err == nil {
		t.Fatal("non-repository passed")
	}
	gitCommand := func(args ...string) {
		t.Helper()
		command := exec.Command("git", append([]string{"-c", "user.name=Coverage test", "-c", "user.email=coverage@example.test", "-c", "commit.gpgsign=false"}, args...)...)
		if data, err := command.CombinedOutput(); err != nil {
			t.Fatalf("git %v: %v\n%s", args, err, data)
		}
	}
	write := func(path, data string) {
		t.Helper()
		if err := os.WriteFile(path, []byte(data), 0o600); err != nil {
			t.Fatal(err)
		}
	}
	if err := checkLOC("HEAD"); err == nil {
		t.Fatal("LOC check outside Git passed")
	}
	gitCommand("init", "-q", "--initial-branch=main")
	if err := run(directory); err == nil {
		t.Fatal("repository without parent passed")
	}
	const original = "package fixture\nconst one = 1\nconst two = 2\n"
	write("fixture.go", original)
	gitCommand("add", "fixture.go")
	gitCommand("commit", "-qm", "parent")
	write("fixture.go", "package fixture\n")
	if err := run(directory); err == nil || !strings.Contains(err.Error(), "read verified coverage") {
		t.Fatalf("missing baseline error = %v", err)
	}
	parent, err := git("rev-parse", "HEAD")
	if err != nil {
		t.Fatal(err)
	}
	baseline := coveragegate.Report{Revision: parent, Scope: runtime.Version() + "/" + runtime.GOOS + "/" + runtime.GOARCH + "/go-packages-v1", Total: map[string]coveragegate.Metric{}}
	for _, name := range []string{"statements", "branches", "functions", "lines"} {
		baseline.Total[name] = coveragegate.Metric{Total: 1, Covered: 1}
	}
	path := filepath.Join(".git", "coverage-baselines", parent)
	if err := os.MkdirAll(path, 0o700); err != nil {
		t.Fatal(err)
	}
	data, err := json.Marshal(baseline)
	if err != nil {
		t.Fatal(err)
	}
	write(filepath.Join(path, "report.json"), string(data))
	if err := run(directory); err == nil || !strings.Contains(err.Error(), "coverage.out") {
		t.Fatalf("missing current report error = %v", err)
	}
	const xml = `<coverage lines-covered="1" lines-valid="1"><packages><package name="example"><classes><class filename="a.go"><methods><method line-rate="1"/></methods></class></classes></package></packages></coverage>`
	for name, contents := range map[string]string{
		"coverage.out":   "mode: set\nexample/a.go:1.1,2.1 1 1\n",
		"coverage.xml":   xml,
		"branches.json":  `{"packages":[{"import_path":"example","files":[{"name":"a.go","blocks":[{"idx":1,"line":1,"col":1,"end_line":2,"end_col":1}],"branches":[{"block_idx":1}]}]}]}`,
		"branches.cover": "mode: set\nexample/a.go:1.1,2.1 1 1\n",
	} {
		write(filepath.Join(directory, name), contents)
	}
	if err := run(directory); err != nil {
		t.Fatal(err)
	}
	report, err := os.ReadFile(filepath.Join(directory, "summary.json"))
	if err != nil || !strings.Contains(string(report), `"Covered": 1`) {
		t.Fatalf("summary = %s, %v", report, err)
	}
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
					gitCommand("add", "--intent-to-add", test.extra)
				}
				t.Cleanup(func() {
					gitCommand("reset", "-q", "--", test.extra)
					if err := os.Remove(test.extra); err != nil {
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
	write(filepath.Join(directory, "coverage.xml"), strings.Replace(xml, `lines-covered="1"`, `lines-covered="0"`, 1))
	if err := run(directory); err == nil || !strings.Contains(err.Error(), "lines coverage") {
		t.Fatalf("coverage regression error = %v", err)
	}
	write(filepath.Join(directory, "coverage.xml"), xml)
	gitCommand("add", "fixture.go")
	gitCommand("commit", "-qm", "candidate")
	if err := run(directory); err != nil {
		t.Fatalf("clean candidate must compare its parent: %v", err)
	}
}
