// Copyright the Xquik contributors.
// SPDX-License-Identifier: Apache-2.0

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
	git := func(command *exec.Cmd, env ...string) ([]byte, error) {
		command.Env = append(append(os.Environ(), "GIT_INDEX_FILE="+filepath.Join(directory, "index"), "TZ=UTC"), env...)
		var diagnostic bytes.Buffer
		command.Stderr = &diagnostic
		output, err := command.Output()
		if err != nil {
			return nil, fmt.Errorf("git %s: %w: %s", command.Args[1], err, diagnostic.String())
		}
		return output, nil
	}
	var outputs [4][]byte
	for i, command := range []*exec.Cmd{
		exec.Command("git", "read-tree", "HEAD"),
		exec.Command("git", "add", "--all", "--", "."),
		exec.Command("git", "write-tree"),
		exec.Command("git", "rev-parse", "--path-format=absolute", "--git-common-dir"),
	} {
		if outputs[i], err = git(command); err != nil {
			return nil, 0, err
		}
	}
	if err := os.WriteFile(filepath.Join(directory, "HEAD"), outputs[2], 0o600); err != nil {
		return nil, 0, err
	}
	snapshot := []string{"GIT_DIR=" + directory, "GIT_COMMON_DIR=" + strings.TrimSpace(string(outputs[3]))}
	paths, err := git(exec.Command("git", "ls-tree", "-rz", "--name-only", "HEAD"), snapshot...)
	if err != nil {
		return nil, 0, err
	}
	archive, err := git(exec.Command("git", "archive", "--format=zip", "--prefix=x-twitter-scraper-go/", "--mtime=1980-01-01T00:00:00Z", "-0", "HEAD"), snapshot...)
	return archive, bytes.Count(paths, []byte{0}), err
}

func run(args []string) error {
	if len(args) > 1 || len(args) == 1 && args[0] != "--archive" {
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
		if _, err := os.Stdout.Write(first); err != nil {
			return err
		}
	}
	fmt.Fprintf(os.Stderr, "reproducible module archive: %d files, sha256:%x\n", count, firstHash)
	return nil
}

func main() {
	if err := run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
