// Copyright the Xquik contributors.
// SPDX-License-Identifier: Apache-2.0

package xtwitterscraper_test

import (
	"encoding/json"
	"io/fs"
	"os"
	"os/exec"
	"strings"
	"testing"
)

func releaseGit(t *testing.T, repository *os.Root, command *exec.Cmd) {
	t.Helper()
	command.Dir = repository.Name()
	if output, err := command.CombinedOutput(); err != nil {
		t.Fatalf("%v: %v\n%s", command.Args, err, output)
	}
}

func releaseRoot(t *testing.T, directory string) *os.Root {
	t.Helper()
	root, err := os.OpenRoot(directory)
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

func releaseFixture(t *testing.T) *os.Root {
	t.Helper()
	repository := releaseRoot(t, t.TempDir())
	source := releaseRoot(t, ".")
	if err := repository.Mkdir("scripts", 0o700); err != nil {
		t.Fatal(err)
	}
	for _, name := range []string{"scripts/release", "scripts/check-all", "package.json"} {
		data, err := source.ReadFile(name)
		if err != nil {
			t.Fatal(err)
		}
		if err := repository.WriteFile(name, data, 0o700); err != nil {
			t.Fatal(err)
		}
	}
	if err := repository.WriteFile(".release-please-manifest.json", []byte("{\".\":\"0.19.0\"}\n"), 0o600); err != nil {
		t.Fatal(err)
	}
	releaseGit(t, repository, exec.Command("git", "init", "-q", "--initial-branch=main"))
	releaseGit(t, repository, exec.Command("git", "config", "user.name", "Release test"))
	releaseGit(t, repository, exec.Command("git", "config", "user.email", "release@example.test"))
	releaseGit(t, repository, exec.Command("git", "config", "commit.gpgsign", "false"))
	releaseGit(t, repository, exec.Command("git", "add", "."))
	releaseGit(t, repository, exec.Command("git", "commit", "-qm", "release fixture"))
	return repository
}

func TestPublishRefusesInvalidReleaseSourceBeforePreparingArtifacts(t *testing.T) {
	for _, scenario := range []struct {
		name string
		want string
	}{
		{"missing tag", "Create the matching immutable release tag first."},
		{"tag on another commit", "Create the matching immutable release tag first."},
		{"unmerged commit", "Merge the release into main first."},
	} {
		t.Run(scenario.name, func(t *testing.T) {
			repository := releaseFixture(t)
			if scenario.name == "tag on another commit" {
				releaseGit(t, repository, exec.Command("git", "tag", "v0.19.0"))
				releaseGit(t, repository, exec.Command("git", "commit", "--allow-empty", "-qm", "candidate"))
			}
			if scenario.name == "unmerged commit" {
				releaseGit(t, repository, exec.Command("git", "init", "-q", "--bare", ".git/origin.git"))
				releaseGit(t, repository, exec.Command("git", "remote", "add", "origin", ".git/origin.git"))
				releaseGit(t, repository, exec.Command("git", "push", "-q", "origin", "HEAD:refs/heads/main"))
				releaseGit(t, repository, exec.Command("git", "commit", "--allow-empty", "-qm", "unmerged candidate"))
				releaseGit(t, repository, exec.Command("git", "tag", "v0.19.0"))
				releaseGit(t, repository, exec.Command("git", "push", "-q", "origin", "refs/tags/v0.19.0"))
			}
			workspace := releaseRoot(t, t.TempDir())
			command := exec.Command("bash", "scripts/release", "publish")
			command.Dir = repository.Name()
			command.Env = append(os.Environ(), "TMPDIR="+workspace.Name())
			var stdout, stderr strings.Builder
			command.Stdout, command.Stderr = &stdout, &stderr
			err := command.Run()
			if exit, ok := err.(*exec.ExitError); !ok || exit.ExitCode() != 1 {
				t.Fatalf("publish exit: %v\n%s\n%s", err, stdout.String(), stderr.String())
			}
			var result struct {
				Status   string `json:"status"`
				ExitCode int    `json:"exitCode"`
			}
			if err := json.Unmarshal([]byte(stdout.String()), &result); err != nil || result.Status != "failed" || result.ExitCode != 1 {
				t.Fatalf("publish result: %q, %v", stdout.String(), err)
			}
			if !strings.Contains(stderr.String(), scenario.want) {
				t.Fatalf("publish diagnostic: %q; want %q", stderr.String(), scenario.want)
			}
			if _, err := repository.Stat("dist"); !os.IsNotExist(err) {
				t.Fatalf("publish created artifacts before validating its source: %v", err)
			}
			if entries, err := fs.ReadDir(workspace.FS(), "."); err != nil || len(entries) != 0 {
				t.Fatalf("publish ran native preparation before validating its source: %v, %v", entries, err)
			}
		})
	}
}
