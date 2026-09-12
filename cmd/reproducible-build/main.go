// Copyright the Xquik contributors.
// SPDX-License-Identifier: Apache-2.0

// reproducible-build verifies the source archive used by this Go library.
package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func buildArchive() ([]byte, int, error) {
	directory, err := os.MkdirTemp("", "xquik-archive-")
	if err != nil {
		return nil, 0, err
	}
	defer os.RemoveAll(directory)
	git := func(args ...string) ([]byte, error) {
		command := exec.Command("git", args...)
		command.Env = append(os.Environ(), "GIT_INDEX_FILE="+filepath.Join(directory, "index"), "TZ=UTC")
		var diagnostic bytes.Buffer
		command.Stderr = &diagnostic
		output, err := command.Output()
		if err != nil {
			return nil, fmt.Errorf("git %s: %w: %s", args[0], err, diagnostic.String())
		}
		return output, nil
	}
	for _, args := range [][]string{{"read-tree", "HEAD"}, {"add", "--all", "--", "."}} {
		if _, err := git(args...); err != nil {
			return nil, 0, err
		}
	}
	tree, err := git("write-tree")
	if err != nil {
		return nil, 0, err
	}
	revision := strings.TrimSpace(string(tree))
	paths, err := git("ls-tree", "-rz", "--name-only", revision)
	if err != nil {
		return nil, 0, err
	}
	archive, err := git("archive", "--format=zip", "--prefix=x-twitter-scraper-go/", "--mtime=1980-01-01T00:00:00Z", "-0", revision)
	return archive, bytes.Count(paths, []byte{0}), err
}

func run(args []string) error {
	if len(args) > 1 {
		return fmt.Errorf("usage: ./scripts/reproducible-build [output.zip]")
	}
	first, count, err := buildArchive()
	if err != nil {
		return err
	}
	second, _, err := buildArchive()
	if err != nil {
		return err
	}
	firstHash, secondHash := sha256.Sum256(first), sha256.Sum256(second)
	if firstHash != secondHash {
		return fmt.Errorf("module archives differ: %x != %x", firstHash, secondHash)
	}
	if len(args) == 1 {
		if err := os.MkdirAll(filepath.Dir(args[0]), 0o755); err != nil {
			return err
		}
		if err := os.WriteFile(args[0], first, 0o644); err != nil {
			return err
		}
	}
	fmt.Printf("reproducible module archive: %d files, sha256:%x\n", count, firstHash)
	return nil
}

func main() {
	if err := run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
