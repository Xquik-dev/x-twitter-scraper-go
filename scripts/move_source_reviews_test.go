// Copyright the Xquik contributors.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"testing/iotest"
)

func sum(data string) string { return fmt.Sprintf("%x", sha256.Sum256([]byte(data))) }

const baseFile = "// header\n\npackage x\n\nfunc a() {\n\tdo()\n\treturn\n}\n\nfunc b() {\n\tdo()\n\treturn\n}\n"

func recordFor(line string, code string) record {
	r := record{Rule: "G104", File: "x.go", Line: line, Code: code, Hash: sum(baseFile), Severity: "LOW", Confidence: "HIGH", Details: "Errors unhandled"}
	r.text = fmt.Sprintf("## G104 in x.go:%s\n\nReason.\n\nSource SHA-256: `%s`\nSeverity: LOW. Confidence: HIGH.\nFinding: Errors unhandled\n\n```go\n%s```\n\n", line, r.Hash, shownCode(code))
	return r
}

func TestMapLine(t *testing.T) {
	hunks := parseHunks("@@ -1,2 +0,0 @@\n-a\n-b\n@@ -6,0 +5 @@\n+c\n@@ -9 +8,2 @@\n-d\n+e\n+f\n")
	cases := []struct {
		old, new int
		ok       bool
	}{{1, 0, false}, {2, 0, false}, {3, 1, true}, {6, 4, true}, {7, 6, true}, {9, 0, false}, {10, 10, true}}
	for _, c := range cases {
		got, ok := mapLine(hunks, c.old)
		if got != c.new || ok != c.ok {
			t.Errorf("mapLine(%d) = %d, %v, want %d, %v", c.old, got, ok, c.new, c.ok)
		}
	}
}

func TestAnOutOfRangeHunkNumberReadsAsZero(t *testing.T) {
	if hunks := parseHunks("@@ -99999999999999999999 +1 @@\n"); len(hunks) != 1 || hunks[0] != (hunk{oldStart: 0, oldCount: 1, newCount: 1}) {
		t.Fatalf("hunks %+v", hunks)
	}
}

func TestAFindingBelowRemovedLinesMoves(t *testing.T) {
	current := strings.Replace(baseFile, "// header\n\n", "", 1)
	r := recordFor("11", "10: func b() {\n11: \tdo()\n12: \treturn\n")
	diff := "@@ -1,2 +0,0 @@\n"

	moved, result := move(r, []byte(baseFile), []byte(current), diff)

	if !result.moved || moved.Line != "9" || moved.Code != "8: func b() {\n9: \tdo()\n10: \treturn\n" || moved.Hash != sum(current) {
		t.Fatalf("got %+v %+v", result, moved)
	}
	if !strings.Contains(moved.text, "## G104 in x.go:9\n") || !strings.Contains(moved.text, "Source SHA-256: `"+sum(current)+"`") || !strings.Contains(moved.text, "```go\n8: func b() {\n9: \tdo()\n10: \treturn\n```") {
		t.Fatalf("text %q", moved.text)
	}
}

func TestAnEmptyExcerptLineKeepsTheShortFormInTheRecords(t *testing.T) {
	current := strings.Replace(baseFile, "// header\n\n", "", 1)
	r := recordFor("10", "9: \n10: func b() {\n11: \tdo()\n")

	moved, result := move(r, []byte(baseFile), []byte(current), "@@ -1,2 +0,0 @@\n")

	if !result.moved || moved.Code != "7: \n8: func b() {\n9: \tdo()\n" || !strings.Contains(moved.text, "```go\n7:\n8: func b() {\n9: \tdo()\n```") {
		t.Fatalf("got %+v %q", result, moved.text)
	}
}

func TestAFindingWhoseSurroundingLineChangedKeepsItsPin(t *testing.T) {
	current := strings.Replace(baseFile, "func b() {\n\tdo()\n\treturn\n", "func b() {\n\tdo()\n\treturn // changed\n", 1)
	r := recordFor("11", "10: func b() {\n11: \tdo()\n12: \treturn\n")

	moved, result := move(r, []byte(baseFile), []byte(current), "@@ -12 +12 @@\n-\treturn\n+\treturn // changed\n")

	if result.moved || result.reason != "line 12 changed" || moved.fingerprint() != r.fingerprint() {
		t.Fatalf("got %+v", result)
	}
}

func TestALineAddedInsideTheExcerptKeepsItsPin(t *testing.T) {
	current := strings.Replace(baseFile, "func b() {\n\tdo()\n", "func b() {\n\tnew()\n\tdo()\n", 1)
	r := recordFor("11", "10: func b() {\n11: \tdo()\n12: \treturn\n")

	_, result := move(r, []byte(baseFile), []byte(current), "@@ -10,0 +11 @@\n+\tnew()\n")

	if result.moved || result.reason != "lines were added before line 11" {
		t.Fatalf("got %+v", result)
	}
}

func TestADiffThatMissesAChangeIsCaughtByTheText(t *testing.T) {
	current := strings.Replace(baseFile, "func b() {\n\tdo()\n", "func b() {\n\tredo()\n", 1)
	r := recordFor("11", "10: func b() {\n11: \tdo()\n12: \treturn\n")

	_, result := move(r, []byte(baseFile), []byte(current), "")

	if result.moved || result.reason != "line 11 reads differently at line 11" {
		t.Fatalf("got %+v", result)
	}
}

func TestARecordWhoseBaseFileHasAnotherHashKeepsItsPin(t *testing.T) {
	r := recordFor("11", "10: func b() {\n11: \tdo()\n12: \treturn\n")

	_, result := move(r, []byte(baseFile+"\n"), []byte(baseFile), "")

	if result.moved || result.reason != "its file at the base commit has another hash" {
		t.Fatalf("got %+v", result)
	}
}

func TestAnUnchangedFileKeepsItsPin(t *testing.T) {
	r := recordFor("11", "10: func b() {\n11: \tdo()\n12: \treturn\n")

	moved, result := move(r, []byte(baseFile), []byte(baseFile), "")

	if !result.moved || moved.fingerprint() != r.fingerprint() {
		t.Fatalf("got %+v", result)
	}
}

func TestALineRangeMoves(t *testing.T) {
	current := strings.Replace(baseFile, "// header\n\n", "", 1)
	r := recordFor("10-12", "10: func b() {\n11: \tdo()\n12: \treturn\n")

	moved, result := move(r, []byte(baseFile), []byte(current), "@@ -1,2 +0,0 @@\n")

	if !result.moved || moved.Line != "8-10" || !strings.Contains(moved.text, "## G104 in x.go:8-10\n") {
		t.Fatalf("got %+v %q", result, moved.text)
	}
}

type fakeRepository struct {
	base, current, diffs map[string]string
	written              map[string]string
}

func (f *fakeRepository) baseFile(path string) ([]byte, error) {
	if data, ok := f.base[path]; ok {
		return []byte(data), nil
	}
	return nil, os.ErrNotExist
}

func (f *fakeRepository) currentFile(path string) ([]byte, error) {
	if data, ok := f.current[path]; ok {
		return []byte(data), nil
	}
	return nil, os.ErrNotExist
}

func (f *fakeRepository) diff(path string) (string, error) { return f.diffs[path], nil }

func (f *fakeRepository) write(path string, data []byte) error {
	f.written[path] = string(data)
	return nil
}

func repositoryWith(t *testing.T, records ...record) *fakeRepository {
	t.Helper()
	var markdown strings.Builder
	markdown.WriteString("# Approved source findings\n\n")
	var index []string
	for _, r := range records {
		markdown.WriteString(r.text)
		index = append(index, r.fingerprint())
	}
	data, err := json.MarshalIndent(index, "", "  ")
	if err != nil {
		t.Fatal(err)
	}
	current := strings.Replace(baseFile, "// header\n\n", "", 1)
	return &fakeRepository{
		base:    map[string]string{"x.go": baseFile},
		current: map[string]string{recordsPath: markdown.String(), indexPath: string(data) + "\n", "x.go": current},
		diffs:   map[string]string{"x.go": "@@ -1,2 +0,0 @@\n"},
		written: map[string]string{},
	}
}

func TestRunMovesEveryRecordAndRewritesBothFiles(t *testing.T) {
	first := recordFor("6", "5: func a() {\n6: \tdo()\n7: \treturn\n")
	second := recordFor("11", "10: func b() {\n11: \tdo()\n12: \treturn\n")
	gone := recordFor("3", "2: \n3: package x\n4: \n")
	gone.File = "gone.go"
	gone.text = strings.Replace(gone.text, "x.go:3", "gone.go:3", 1)
	repo := repositoryWith(t, first, second, gone)
	var out strings.Builder

	err := run(repo, &out)

	if err != nil {
		t.Fatal(err)
	}
	current := repo.current["x.go"]
	wantOut := "moved G104 x.go:6 to line 4\nmoved G104 x.go:11 to line 9\nfile x.go " + sum(baseFile) + " -> " + sum(current) + "\norphaned G104 gone.go:3 because its file is not in the base commit\n2 findings moved in 1 files. 0 findings kept their pins. 1 records are orphaned. 0 findings could not move.\n"
	if out.String() != wantOut {
		t.Fatalf("output %q", out.String())
	}
	_, records, err := parseRecords(repo.written[recordsPath])
	if err != nil || len(records) != 3 || records[0].Line != "4" || records[1].Line != "9" || records[2].File != "gone.go" {
		t.Fatalf("records %v %+v", err, records)
	}
	var index []string
	if err := json.Unmarshal([]byte(repo.written[indexPath]), &index); err != nil {
		t.Fatal(err)
	}
	if index[0] != records[0].fingerprint() || index[1] != records[1].fingerprint() || index[2] != gone.fingerprint() {
		t.Fatalf("index %v", index)
	}
}

func TestRunFailsWhenAFindingCannotMove(t *testing.T) {
	r := recordFor("11", "10: func b() {\n11: \tdo()\n12: \treturn\n")
	repo := repositoryWith(t, r)
	repo.current["x.go"] = strings.Replace(repo.current["x.go"], "\treturn\n}\n\nfunc b", "\treturn\n}\n\nfunc bb", 1)
	repo.diffs["x.go"] = "@@ -1,2 +0,0 @@\n@@ -10 +8 @@\n"
	var out strings.Builder

	err := run(repo, &out)

	if err == nil || err.Error() != "1 findings keep their old pins, so cmd/source-security rejects them until they are reviewed again" {
		t.Fatalf("err %v", err)
	}
	if !strings.Contains(out.String(), "refused G104 x.go:11 because line 10 changed\n") {
		t.Fatalf("output %q", out.String())
	}
	var index []string
	if err := json.Unmarshal([]byte(repo.written[indexPath]), &index); err != nil {
		t.Fatal(err)
	}
	if len(index) != 1 || index[0] != r.fingerprint() {
		t.Fatalf("index %v", index)
	}
}

func TestRunRefusesARecordThatTheIndexLacks(t *testing.T) {
	r := recordFor("11", "10: func b() {\n11: \tdo()\n12: \treturn\n")
	repo := repositoryWith(t, r)
	repo.current[indexPath] = "[\"0000\"]\n"

	err := run(repo, io.Discard)

	if err == nil || err.Error() != "record G104 in x.go:11 matches no fingerprint in the index" {
		t.Fatalf("err %v", err)
	}
	if len(repo.written) != 0 {
		t.Fatalf("wrote %v", repo.written)
	}
}

func TestTheRealRecordsMatchTheRealIndex(t *testing.T) {
	markdown, err := os.ReadFile("../" + recordsPath)
	if err != nil {
		t.Fatal(err)
	}
	data, err := os.ReadFile("../" + indexPath)
	if err != nil {
		t.Fatal(err)
	}
	var index []string
	if err := json.Unmarshal(data, &index); err != nil {
		t.Fatal(err)
	}
	_, records, err := parseRecords(string(markdown))
	if err != nil {
		t.Fatal(err)
	}
	approved := map[string]bool{}
	for _, f := range index {
		approved[f] = true
	}
	for _, r := range records {
		if !approved[r.fingerprint()] {
			t.Errorf("%s %s:%s", r.Rule, r.File, r.Line)
		}
	}
	if len(records) != len(index) {
		t.Errorf("%d records, %d fingerprints", len(records), len(index))
	}
}

func treeWith(t *testing.T, files map[string]string, records ...record) string {
	t.Helper()
	var markdown strings.Builder
	markdown.WriteString("# Approved source findings\n\n")
	var index []string
	for _, r := range records {
		markdown.WriteString(r.text)
		index = append(index, r.fingerprint())
	}
	data, err := json.MarshalIndent(index, "", "  ")
	if err != nil {
		t.Fatal(err)
	}
	files[recordsPath] = markdown.String()
	files[indexPath] = string(data) + "\n"
	root := t.TempDir()
	for name, content := range files {
		target := filepath.Join(root, name)
		if err := os.MkdirAll(filepath.Dir(target), 0o700); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(target, []byte(content), 0o600); err != nil {
			t.Fatal(err)
		}
	}
	return root
}

func renamed(r record, file string) record {
	r.text = strings.Replace(r.text, "x.go:"+r.Line, file+":"+r.Line, 1)
	r.File = file
	return r
}

func TestCommandMovesPinsFromAPipedDiff(t *testing.T) {
	current := strings.Replace(strings.Replace(baseFile, "// header\n\n", "", 1), "\tdo()", "\tdo2()", 1)
	moving := recordFor("11", "10: func b() {\n11: \tdo()\n12: \treturn\n")
	kept := renamed(recordFor("6", "5: func a() {\n6: \tdo()\n7: \treturn\n"), "y.go")
	added := renamed(recordFor("3", "2: \n3: package x\n4: \n"), "new.go")
	root := treeWith(t, map[string]string{"x.go": current, "y.go": baseFile, "new.go": baseFile}, moving, kept, added)
	diff := "diff --git a/x.go b/x.go\nindex 1111111..2222222 100644\n--- a/x.go\n+++ b/x.go\n" +
		"@@ -1,2 +0,0 @@\n-// header\n-\n@@ -6 +4 @@\n-\tdo()\n+\tdo2()\n" +
		"diff --git a/new.go b/new.go\nnew file mode 100644\nindex 0000000..3333333\n--- /dev/null\n+++ b/new.go\n@@ -0,0 +1,2 @@\n+// header\n+\n"
	var out strings.Builder

	err := command([]string{"move"}, strings.NewReader(diff), &out, root)

	if err != nil {
		t.Fatal(err)
	}
	want := "moved G104 x.go:11 to line 9\nfile x.go " + sum(baseFile) + " -> " + sum(current) + "\norphaned G104 new.go:3 because its file is not in the base commit\n1 findings moved in 1 files. 1 findings kept their pins. 1 records are orphaned. 0 findings could not move.\n"
	if out.String() != want {
		t.Fatalf("output %q", out.String())
	}
	markdown, err := fs.ReadFile(os.DirFS(root), recordsPath)
	if err != nil {
		t.Fatal(err)
	}
	_, records, err := parseRecords(string(markdown))
	if err != nil || len(records) != 3 || records[0].Line != "9" || records[0].Hash != sum(current) || records[1] != kept || records[2] != added {
		t.Fatalf("records %v %+v", err, records)
	}
}

func TestCommandRejectsArgumentsAndUnreadableInput(t *testing.T) {
	if err := command([]string{"move", "main"}, strings.NewReader(""), io.Discard, t.TempDir()); err == nil || err.Error() != usage {
		t.Fatalf("err %v", err)
	}
	if err := command([]string{"move"}, iotest.ErrReader(errors.New("closed")), io.Discard, t.TempDir()); err == nil || err.Error() != "closed" {
		t.Fatalf("err %v", err)
	}
}

func TestMainKeepsUnchangedPins(t *testing.T) {
	root := treeWith(t, map[string]string{"y.go": baseFile}, renamed(recordFor("6", "5: func a() {\n6: \tdo()\n7: \treturn\n"), "y.go"))
	stdin, err := os.CreateTemp(t.TempDir(), "diff")
	if err != nil {
		t.Fatal(err)
	}
	defer stdin.Close()
	stdout, err := os.CreateTemp(t.TempDir(), "stdout")
	if err != nil {
		t.Fatal(err)
	}
	defer stdout.Close()
	args, input, output := os.Args, os.Stdin, os.Stdout
	defer func() { os.Args, os.Stdin, os.Stdout = args, input, output }()
	os.Args, os.Stdin, os.Stdout = []string{"move"}, stdin, stdout
	t.Chdir(root)

	main()

	if _, err := stdout.Seek(0, io.SeekStart); err != nil {
		t.Fatal(err)
	}
	got, err := io.ReadAll(stdout)
	if err != nil {
		t.Fatal(err)
	}
	if want := "0 findings moved in 0 files. 1 findings kept their pins. 0 records are orphaned. 0 findings could not move.\n"; string(got) != want {
		t.Fatalf("main printed %q, want %q", got, want)
	}
}
