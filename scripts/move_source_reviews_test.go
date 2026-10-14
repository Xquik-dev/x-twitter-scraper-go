//go:build sourcereviews

// Copyright the Xquik contributors.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
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

func repositoryWith(records ...record) *fakeRepository {
	var markdown strings.Builder
	markdown.WriteString("# Approved source findings\n\n")
	var index []string
	for _, r := range records {
		markdown.WriteString(r.text)
		index = append(index, r.fingerprint())
	}
	data, _ := json.MarshalIndent(index, "", "  ")
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
	repo := repositoryWith(first, second, gone)
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
	repo := repositoryWith(r)
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
	_ = json.Unmarshal([]byte(repo.written[indexPath]), &index)
	if len(index) != 1 || index[0] != r.fingerprint() {
		t.Fatalf("index %v", index)
	}
}

func TestRunRefusesARecordThatTheIndexLacks(t *testing.T) {
	r := recordFor("11", "10: func b() {\n11: \tdo()\n12: \treturn\n")
	repo := repositoryWith(r)
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
	data, _ := os.ReadFile("../" + indexPath)
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
