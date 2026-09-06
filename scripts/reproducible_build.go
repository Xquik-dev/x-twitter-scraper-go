//go:build ignore

// Copyright the Xquik contributors.
// SPDX-License-Identifier: Apache-2.0

// reproducible_build verifies the source archive used by this Go library.
package main

import (
	"archive/zip"
	"bytes"
	"crypto/sha256"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

var archiveTime = time.Date(1980, time.January, 1, 0, 0, 0, 0, time.UTC)

func main() {
	if len(os.Args) > 2 {
		exitf("usage: ./scripts/reproducible-build [output.zip]")
	}
	paths := repositoryFiles()
	first := buildArchive(paths)
	firstHash := sha256.Sum256(first)
	secondHash := sha256.Sum256(buildArchive(paths))
	if firstHash != secondHash {
		exitf("module archives differ: %x != %x", firstHash, secondHash)
	}
	fmt.Printf(
		"reproducible module archive: %d files, sha256:%x\n",
		len(paths),
		firstHash,
	)
	if len(os.Args) == 2 {
		writeArchive(os.Args[1], first)
	}
}

func writeArchive(outputPath string, contents []byte) {
	if err := os.MkdirAll(filepath.Dir(outputPath), 0o755); err != nil {
		exitf("create archive directory: %v", err)
	}
	if err := os.WriteFile(outputPath, contents, 0o644); err != nil {
		exitf("write module archive: %v", err)
	}
}

func repositoryFiles() []string {
	command := exec.Command(
		"git",
		"ls-files",
		"-z",
		"--cached",
		"--others",
		"--exclude-standard",
	)
	output, err := command.Output()
	if err != nil {
		exitf("list repository files: %v", err)
	}
	paths := strings.FieldsFunc(filepath.ToSlash(string(output)), func(char rune) bool { return char == 0 })
	sort.Strings(paths)
	return paths
}

func buildArchive(paths []string) []byte {
	var output bytes.Buffer
	archive := zip.NewWriter(&output)
	for _, path := range paths {
		addFile(archive, path)
	}
	if err := archive.Close(); err != nil {
		exitf("close module archive: %v", err)
	}
	return output.Bytes()
}

func addFile(archive *zip.Writer, path string) {
	info, err := os.Lstat(filepath.FromSlash(path))
	if err != nil {
		exitf("inspect %s: %v", path, err)
	}
	if !info.Mode().IsRegular() && info.Mode()&os.ModeSymlink == 0 {
		exitf("unsupported repository entry %s with mode %s", path, info.Mode())
	}

	header := &zip.FileHeader{
		Name:     "x-twitter-scraper-go/" + path,
		Method:   zip.Store,
		Modified: archiveTime,
	}
	mode := os.FileMode(0o644)
	if info.Mode()&os.ModeSymlink != 0 {
		mode = os.ModeSymlink | 0o777
	} else if info.Mode()&0o111 != 0 {
		mode = 0o755
	}
	header.SetMode(mode)

	writer, err := archive.CreateHeader(header)
	if err != nil {
		exitf("create archive entry %s: %v", path, err)
	}
	if info.Mode()&os.ModeSymlink != 0 {
		target, err := os.Readlink(filepath.FromSlash(path))
		if err != nil {
			exitf("read symlink %s: %v", path, err)
		}
		if _, err := io.Copy(writer, strings.NewReader(target)); err != nil {
			exitf("archive symlink %s: %v", path, err)
		}
		return
	}

	file, err := os.Open(filepath.FromSlash(path))
	if err != nil {
		exitf("open %s: %v", path, err)
	}
	if _, err := io.Copy(writer, file); err != nil {
		_ = file.Close()
		exitf("archive %s: %v", path, err)
	}
	if err := file.Close(); err != nil {
		exitf("close %s: %v", path, err)
	}
}

func exitf(format string, args ...any) {
	fmt.Fprintf(os.Stderr, format+"\n", args...)
	os.Exit(1)
}
