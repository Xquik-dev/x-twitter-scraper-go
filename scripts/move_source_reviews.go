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
	"regexp"
	"sort"
	"strconv"
	"strings"
)

const (
	recordsPath = "security/source-reviews.md"
	indexPath   = "security/source-reviews.json"
)

type record struct {
	Rule, File, Line, Code, Hash, Severity, Confidence, Details string
	text                                                        string
}

var (
	headingPattern  = regexp.MustCompile(`(?m)^## (\S+) in (\S+):(\d+(?:-\d+)?)$`)
	hashPattern     = regexp.MustCompile("(?m)^Source SHA-256: `([0-9a-f]{64})`$")
	severityPattern = regexp.MustCompile(`(?m)^Severity: (\S+)\. Confidence: (\S+)\.$`)
	findingPattern  = regexp.MustCompile(`(?m)^Finding: (.*)$`)
	codePattern     = regexp.MustCompile("(?s)\n```go\n(.*?\n)```")
	hunkPattern     = regexp.MustCompile(`(?m)^@@ -(\d+)(?:,(\d+))? \+(\d+)(?:,(\d+))? @@`)
	excerptPattern  = regexp.MustCompile(`^(\d+): (.*)$`)
	emptyPattern    = regexp.MustCompile(`(?m)^(\d+):$`)
	blankPattern    = regexp.MustCompile(`(?m)^(\d+): $`)
)

// The records file shows an empty source line as "12:", while gosec writes
// "12: " with a trailing space. The fingerprint uses gosec's form.
// https://github.com/securego/gosec/blob/v2.29.0/issue/issue.go#L204
func scannerCode(shown string) string { return emptyPattern.ReplaceAllString(shown, "$1: ") }

func shownCode(scanner string) string { return blankPattern.ReplaceAllString(scanner, "$1:") }

func (r record) fingerprint() string {
	item := [8]string{r.Rule, r.File, r.Line, r.Code, r.Hash, r.Severity, r.Confidence, r.Details}
	return fmt.Sprintf("%x", sha256.Sum256([]byte(fmt.Sprintf("%q", item))))
}

func parseRecords(markdown string) (string, []record, error) {
	starts := headingPattern.FindAllStringIndex(markdown, -1)
	if len(starts) == 0 {
		return "", nil, errors.New("the records file has no record")
	}
	records := make([]record, 0, len(starts))
	for i, start := range starts {
		end := len(markdown)
		if i+1 < len(starts) {
			end = starts[i+1][0]
		}
		text := markdown[start[0]:end]
		heading := headingPattern.FindStringSubmatch(text)
		hash := hashPattern.FindStringSubmatch(text)
		severity := severityPattern.FindStringSubmatch(text)
		finding := findingPattern.FindStringSubmatch(text)
		code := codePattern.FindStringSubmatch(text)
		if hash == nil || severity == nil || finding == nil || code == nil {
			return "", nil, fmt.Errorf("record %q lacks its hash, severity, finding or code", heading[0])
		}
		records = append(records, record{
			Rule: heading[1], File: heading[2], Line: heading[3], Code: scannerCode(code[1]), Hash: hash[1],
			Severity: severity[1], Confidence: severity[2], Details: finding[1], text: text,
		})
	}
	return markdown[:starts[0][0]], records, nil
}

type hunk struct{ oldStart, oldCount, newCount int }

func number(digits string) int {
	value, err := strconv.Atoi(digits)
	if err != nil {
		return 0
	}
	return value
}

func lineCount(digits string) int {
	if digits == "" {
		return 1
	}
	return number(digits)
}

func parseHunks(diff string) []hunk {
	var hunks []hunk
	for _, match := range hunkPattern.FindAllStringSubmatch(diff, -1) {
		hunks = append(hunks, hunk{oldStart: number(match[1]), oldCount: lineCount(match[2]), newCount: lineCount(match[4])})
	}
	return hunks
}

func mapLine(hunks []hunk, line int) (int, bool) {
	shift := 0
	for _, h := range hunks {
		lastAffected := h.oldStart + h.oldCount - 1
		if h.oldCount == 0 {
			lastAffected = h.oldStart
		} else if line >= h.oldStart && line <= lastAffected {
			return 0, false
		}
		if line > lastAffected {
			shift += h.newCount - h.oldCount
		}
	}
	return line + shift, true
}

func mapLineField(hunks []hunk, field string) string {
	parts := strings.Split(field, "-")
	for i, part := range parts {
		line := number(part)
		updated, _ := mapLine(hunks, line)
		parts[i] = strconv.Itoa(updated)
	}
	return strings.Join(parts, "-")
}

type outcome struct {
	record  record
	moved   bool
	newLine string
	reason  string
}

func move(r record, base, current []byte, diff string) (record, outcome) {
	if fmt.Sprintf("%x", sha256.Sum256(base)) != r.Hash {
		return r, outcome{record: r, reason: "its file at the base commit has another hash"}
	}
	currentHash := fmt.Sprintf("%x", sha256.Sum256(current))
	if currentHash == r.Hash {
		return r, outcome{record: r, moved: true, newLine: r.Line}
	}
	hunks := parseHunks(diff)
	lines := strings.Split(string(current), "\n")
	var code strings.Builder
	previous := 0
	for _, excerpt := range strings.Split(strings.TrimSuffix(r.Code, "\n"), "\n") {
		parts := excerptPattern.FindStringSubmatch(excerpt)
		if parts == nil {
			return r, outcome{record: r, reason: fmt.Sprintf("its excerpt line %q has no line number", excerpt)}
		}
		old := number(parts[1])
		updated, unchanged := mapLine(hunks, old)
		switch {
		case !unchanged:
			return r, outcome{record: r, reason: fmt.Sprintf("line %d changed", old)}
		case previous != 0 && updated != previous+1:
			return r, outcome{record: r, reason: fmt.Sprintf("lines were added before line %d", old)}
		case updated > len(lines) || lines[updated-1] != parts[2]:
			return r, outcome{record: r, reason: fmt.Sprintf("line %d reads differently at line %d", old, updated)}
		}
		previous = updated
		fmt.Fprintf(&code, "%d: %s\n", updated, parts[2])
	}
	moved := r
	moved.Line = mapLineField(hunks, r.Line)
	moved.Code = code.String()
	moved.Hash = currentHash
	moved.text = strings.Replace(r.text, fmt.Sprintf("## %s in %s:%s\n", r.Rule, r.File, r.Line), fmt.Sprintf("## %s in %s:%s\n", r.Rule, r.File, moved.Line), 1)
	moved.text = strings.Replace(moved.text, "Source SHA-256: `"+r.Hash+"`", "Source SHA-256: `"+currentHash+"`", 1)
	moved.text = strings.Replace(moved.text, "```go\n"+shownCode(r.Code)+"```", "```go\n"+shownCode(moved.Code)+"```", 1)
	return moved, outcome{record: moved, moved: true, newLine: moved.Line}
}

type repository interface {
	baseFile(path string) ([]byte, error)
	currentFile(path string) ([]byte, error)
	diff(path string) (string, error)
	write(path string, data []byte) error
}

func run(repo repository, out io.Writer) error {
	markdown, err := repo.currentFile(recordsPath)
	if err != nil {
		return err
	}
	indexData, err := repo.currentFile(indexPath)
	if err != nil {
		return err
	}
	var index []string
	if err := json.Unmarshal(indexData, &index); err != nil {
		return err
	}
	preamble, records, err := parseRecords(string(markdown))
	if err != nil {
		return err
	}
	if len(records) != len(index) {
		return fmt.Errorf("the records file holds %d records and the index %d fingerprints", len(records), len(index))
	}
	position := map[string]int{}
	for i, fingerprint := range index {
		position[fingerprint] = i
	}
	type fileHashes struct{ old, new string }
	hashes := map[string]fileHashes{}
	var refused, orphaned []string
	moved, kept := 0, 0
	var text strings.Builder
	text.WriteString(preamble)
	for _, r := range records {
		slot, ok := position[r.fingerprint()]
		if !ok {
			return fmt.Errorf("record %s in %s:%s matches no fingerprint in the index", r.Rule, r.File, r.Line)
		}
		updated, result := r, outcome{record: r}
		base, baseErr := repo.baseFile(r.File)
		current, currentErr := repo.currentFile(r.File)
		switch {
		case baseErr != nil:
			orphaned = append(orphaned, fmt.Sprintf("orphaned %s %s:%s because its file is not in the base commit", r.Rule, r.File, r.Line))
			index[slot] = r.fingerprint()
			text.WriteString(r.text)
			continue
		case currentErr != nil:
			result.reason = "its file is gone"
		default:
			diff, err := repo.diff(r.File)
			if err != nil {
				return fmt.Errorf("git diff of %s failed. %w", r.File, err)
			}
			updated, result = move(r, base, current, diff)
		}
		switch {
		case !result.moved:
			refused = append(refused, fmt.Sprintf("refused %s %s:%s because %s", r.Rule, r.File, r.Line, result.reason))
		case updated.Hash == r.Hash:
			kept++
		default:
			moved++
			hashes[r.File] = fileHashes{old: r.Hash, new: updated.Hash}
			fmt.Fprintf(out, "moved %s %s:%s to line %s\n", r.Rule, r.File, r.Line, updated.Line)
		}
		index[slot] = updated.fingerprint()
		text.WriteString(updated.text)
	}
	files := make([]string, 0, len(hashes))
	for file := range hashes {
		files = append(files, file)
	}
	sort.Strings(files)
	for _, file := range files {
		fmt.Fprintf(out, "file %s %s -> %s\n", file, hashes[file].old, hashes[file].new)
	}
	for _, line := range append(orphaned, refused...) {
		fmt.Fprintln(out, line)
	}
	newIndex, err := json.MarshalIndent(index, "", "  ")
	if err != nil {
		return err
	}
	if err := repo.write(indexPath, append(newIndex, '\n')); err != nil {
		return err
	}
	if err := repo.write(recordsPath, []byte(text.String())); err != nil {
		return err
	}
	fmt.Fprintf(out, "%d findings moved in %d files. %d findings kept their pins. %d records are orphaned. %d findings could not move.\n", moved, len(files), kept, len(orphaned), len(refused))
	if len(refused) > 0 {
		return fmt.Errorf("%d findings keep their old pins, so cmd/source-security rejects them until they are reviewed again", len(refused))
	}
	return nil
}

type change struct {
	hunks string
	added bool
}

func parseChanges(diff string) map[string]change {
	changes := map[string]change{}
	var path string
	var current change
	inHunks := false
	flush := func() {
		if path != "" {
			changes[path] = current
		}
		path, current, inHunks = "", change{}, false
	}
	for _, line := range strings.SplitAfter(diff, "\n") {
		switch {
		case strings.HasPrefix(line, "diff --git "):
			flush()
		case inHunks || strings.HasPrefix(line, "@@ "):
			inHunks = true
			current.hunks += line
		case line == "--- /dev/null\n":
			current.added = true
		case strings.HasPrefix(line, "--- a/") && path == "":
			path = strings.TrimSuffix(strings.TrimPrefix(line, "--- a/"), "\n")
		case strings.HasPrefix(line, "+++ b/") && path == "":
			path = strings.TrimSuffix(strings.TrimPrefix(line, "+++ b/"), "\n")
		}
	}
	flush()
	return changes
}

func reverse(current []byte, hunks string) []byte {
	lines := strings.SplitAfter(string(current), "\n")
	var base strings.Builder
	next := 0
	for _, line := range strings.SplitAfter(hunks, "\n") {
		if header := hunkPattern.FindStringSubmatch(line); header != nil {
			start, count := number(header[3]), lineCount(header[4])
			if count > 0 {
				start--
			}
			start = min(max(start, next), len(lines))
			base.WriteString(strings.Join(lines[next:start], ""))
			next = min(start+count, len(lines))
		} else if strings.HasPrefix(line, "-") {
			base.WriteString(line[1:])
		}
	}
	base.WriteString(strings.Join(lines[next:], ""))
	return []byte(base.String())
}

type workingTree struct {
	root    string
	changes map[string]change
}

func (w workingTree) currentFile(path string) ([]byte, error) {
	return fs.ReadFile(os.DirFS(w.root), path)
}

func (w workingTree) baseFile(path string) ([]byte, error) {
	changed, ok := w.changes[path]
	if changed.added {
		return nil, fs.ErrNotExist
	}
	current, err := w.currentFile(path)
	if !ok {
		return current, err
	}
	return reverse(current, changed.hunks), nil
}

func (w workingTree) diff(path string) (string, error) { return w.changes[path].hunks, nil }

func (w workingTree) write(path string, data []byte) error {
	return os.WriteFile(filepath.Join(w.root, path), data, 0o600)
}

const usage = "Run git diff --no-color --no-ext-diff --no-renames -U0 <base commit> | go run ./scripts from the repository root."

type failure string

func (f failure) Error() string { return string(f) }

func command(args []string, stdin io.Reader, out io.Writer, root string) error {
	if len(args) != 1 {
		return failure(usage)
	}
	diff, err := io.ReadAll(stdin)
	if err != nil {
		return err
	}
	return run(workingTree{root: root, changes: parseChanges(string(diff))}, out)
}

func main() {
	if err := command(os.Args, os.Stdin, os.Stdout, "."); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
