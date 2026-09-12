// Copyright the Xquik contributors.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"archive/zip"
	"bytes"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

func TestArchiveWorkingTree(t *testing.T) {
	t.Chdir(t.TempDir())
	t.Setenv("TZ", "Pacific/Honolulu")
	git := func(args ...string) []byte {
		t.Helper()
		data, err := exec.Command("git", args...).CombinedOutput()
		if err != nil {
			t.Fatalf("git %v: %v\n%s", args, err, data)
		}
		return data
	}
	write := func(name, value string) {
		t.Helper()
		if err := os.WriteFile(name, []byte(value), 0o644); err != nil {
			t.Fatal(err)
		}
	}
	git("init", "-q")
	for name, value := range map[string]string{".gitignore": "ignored.dat\n", "plain": "initial", "executable": "executable", "obsolete": "remove"} {
		write(name, value)
	}
	if err := os.Chmod("executable", 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.Symlink("plain", "link"); err != nil {
		t.Fatal(err)
	}
	git("add", ".")
	git("-c", "user.name=Test", "-c", "user.email=test@example.invalid", "-c", "commit.gpgsign=false", "commit", "-qm", "fixture")
	write("plain", "staged")
	git("add", "plain")
	write("plain", "current")
	write("new", "untracked")
	write("ignored.dat", "excluded")
	if err := os.Remove("obsolete"); err != nil {
		t.Fatal(err)
	}
	index := git("diff", "--cached", "--raw")
	first, count, err := buildArchive()
	if err != nil || count != 5 {
		t.Fatalf("buildArchive: %d files, %v", count, err)
	}
	second, _, err := buildArchive()
	if err != nil || !bytes.Equal(first, second) || !bytes.Equal(index, git("diff", "--cached", "--raw")) {
		t.Fatalf("archive differs or index changed: %v", err)
	}
	reader, err := zip.NewReader(bytes.NewReader(first), int64(len(first)))
	if err != nil {
		t.Fatal(err)
	}
	want := map[string]string{".gitignore": "ignored.dat\n", "plain": "current", "executable": "executable", "link": "plain", "new": "untracked"}
	for _, file := range reader.File {
		if file.FileInfo().IsDir() {
			continue
		}
		name := filepath.Base(file.Name)
		stream, err := file.Open()
		if err != nil {
			t.Fatal(err)
		}
		data, readErr := io.ReadAll(stream)
		closeErr := stream.Close()
		expected, exists := want[name]
		if !exists || file.Name != "x-twitter-scraper-go/"+name || string(data) != expected || readErr != nil || closeErr != nil || file.Modified.UTC().Format("2006-01-02T15:04:05") != "1980-01-01T00:00:00" {
			t.Fatalf("entry %s: %q, %v, %v, %v", file.Name, data, file.Modified, readErr, closeErr)
		}
		if name == "executable" && file.Mode().Perm() != 0o755 || name == "link" && file.Mode()&os.ModeSymlink == 0 {
			t.Fatalf("entry %s mode: %v", name, file.Mode())
		}
		delete(want, name)
	}
	if len(want) != 0 {
		t.Fatalf("missing entries: %v", want)
	}
}
