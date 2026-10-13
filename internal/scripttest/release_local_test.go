// Copyright the Xquik contributors.
// SPDX-License-Identifier: Apache-2.0

package scripttest

import (
	"bytes"
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

// The fake GitHub CLI is signed in, knows no release, and records each call.
// It saves the notes of a created release so a test can read them.
const fakeGH = `#!/bin/sh
echo "gh $*" >>"$FAKE_CALLS"
case "$1 $2" in
  "auth status") exit 0 ;;
  "release create")
    while [ $# -gt 0 ]; do
      if [ "$1" = "--notes-file" ]; then cat "$2" >"$FAKE_NOTES"; fi
      shift
    done
    exit 0 ;;
  *) exit 1 ;;
esac
`

// The fake curl answers as the module proxy does before and after a release.
const fakeCurl = `#!/bin/sh
for url; do :; done
echo "curl $url" >>"$FAKE_CALLS"
case "$url" in
  */@v/list) echo v1.2.2 ;;
  */@v/v1.2.3.info) echo '{"Version":"v1.2.3"}' ;;
  *) exit 22 ;;
esac
`

const fixtureChangelog = `# Changelog

## 1.2.3 (2026-01-02)

### Features

* add retweet timestamps

## 1.2.2 (2026-01-01)

* older notes
`

type releaseFixture struct {
	t                   *testing.T
	repo, origin, calls string
	notes               string
	env                 []string
}

// newReleaseFixture builds a repository whose main matches a local origin. It
// holds a copy of bin/release-local, stand-ins for the 2 check scripts, and
// fake gh and curl commands. Git and Go are real, and nothing leaves the folder.
func newReleaseFixture(t *testing.T) *releaseFixture {
	t.Helper()
	script, err := os.ReadFile("../../bin/release-local")
	if err != nil {
		t.Fatal(err)
	}
	root := t.TempDir()
	fixture := &releaseFixture{
		t:      t,
		repo:   filepath.Join(root, "repo"),
		origin: filepath.Join(root, "origin.git"),
		calls:  filepath.Join(root, "calls.log"),
		notes:  filepath.Join(root, "notes.md"),
	}
	tools := filepath.Join(root, "tools")
	settings := filepath.Join(root, "gh-settings")
	recorder := func(name string) string {
		return "#!/bin/sh\necho " + name + " >>\"$FAKE_CALLS\"\n"
	}
	for name, file := range map[string]struct {
		data string
		mode os.FileMode
	}{
		"repo/go.mod":                        {"module github.com/Xquik-dev/x-twitter-scraper-go\n\ngo 1.24\n", 0o644},
		"repo/lib.go":                        {"package xtwitterscraper\n", 0o644},
		"repo/internal/version.go":           {"package internal\n\nconst PackageVersion = \"1.2.3\" // x-release-please-version\n", 0o644},
		"repo/.release-please-manifest.json": {"{\n  \".\": \"1.2.3\"\n}\n", 0o644},
		"repo/CHANGELOG.md":                  {fixtureChangelog, 0o644},
		"repo/bin/release-local":             {string(script), 0o755},
		"repo/bin/check-local":               {recorder("check-local"), 0o755},
		"repo/scripts/dependency-audit":      {recorder("dependency-audit"), 0o755},
		"tools/gh":                           {fakeGH, 0o755},
		"tools/curl":                         {fakeCurl, 0o755},
		"gh-settings/.keep":                  {"", 0o644},
	} {
		path := filepath.Join(root, name)
		if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(path, []byte(file.data), file.mode); err != nil {
			t.Fatal(err)
		}
	}
	for _, variable := range os.Environ() {
		name, _, _ := strings.Cut(variable, "=")
		switch name {
		case "PATH", "GH_TOKEN", "GITHUB_TOKEN", "GH_ENTERPRISE_TOKEN", "GITHUB_ENTERPRISE_TOKEN":
		default:
			fixture.env = append(fixture.env, variable)
		}
	}
	// The fixture ignores this machine's Git settings and GitHub login.
	fixture.env = append(fixture.env,
		"PATH="+tools+string(os.PathListSeparator)+os.Getenv("PATH"),
		"FAKE_CALLS="+fixture.calls,
		"FAKE_NOTES="+fixture.notes,
		"GH_CONFIG_DIR="+settings,
		"GIT_CONFIG_GLOBAL="+os.DevNull,
		"GIT_CONFIG_NOSYSTEM=1",
		"GIT_AUTHOR_NAME=Release test",
		"GIT_AUTHOR_EMAIL=release@example.test",
		"GIT_COMMITTER_NAME=Release test",
		"GIT_COMMITTER_EMAIL=release@example.test",
	)
	fixture.git(root, "init", "--quiet", "--bare", "--initial-branch=main", fixture.origin)
	fixture.git(fixture.repo, "init", "--quiet", "--initial-branch=main")
	fixture.git(fixture.repo, "add", "--all")
	fixture.git(fixture.repo, "commit", "--quiet", "--message", "fixture")
	fixture.git(fixture.repo, "remote", "add", "origin", fixture.origin)
	fixture.git(fixture.repo, "push", "--quiet", "origin", "main")
	return fixture
}

func (fixture *releaseFixture) git(directory string, args ...string) string {
	fixture.t.Helper()
	command := exec.Command("git", args...)
	command.Dir, command.Env = directory, fixture.env
	output, err := command.CombinedOutput()
	if err != nil {
		fixture.t.Fatalf("git %v: %v\n%s", args, err, output)
	}
	return strings.TrimSpace(string(output))
}

func (fixture *releaseFixture) release(args ...string) (code int, stdout, stderr string) {
	fixture.t.Helper()
	command := exec.Command(filepath.Join(fixture.repo, "bin", "release-local"), args...)
	command.Dir, command.Env = fixture.t.TempDir(), fixture.env
	var out, failure bytes.Buffer
	command.Stdout, command.Stderr = &out, &failure
	if err := command.Run(); err != nil {
		var exit *exec.ExitError
		if !errors.As(err, &exit) {
			fixture.t.Fatal(err)
		}
		code = exit.ExitCode()
	}
	return code, out.String(), failure.String()
}

func (fixture *releaseFixture) recorded() string {
	fixture.t.Helper()
	data, err := os.ReadFile(fixture.calls)
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		fixture.t.Fatal(err)
	}
	return string(data)
}

func TestReleaseLocalDryRunPublishesNothing(t *testing.T) {
	fixture := newReleaseFixture(t)
	revision := fixture.git(fixture.repo, "rev-parse", "HEAD")

	code, stdout, stderr := fixture.release("--dry-run")
	if code != 0 {
		t.Fatalf("dry run exit %d\n%s\n%s", code, stdout, stderr)
	}
	listing := "Module path  github.com/Xquik-dev/x-twitter-scraper-go\n" +
		"Version      v1.2.3\n" +
		"Commit       " + revision + "\n" +
		"Checksum     h1:"
	files := "Files        8\n" +
		".release-please-manifest.json\n" +
		"CHANGELOG.md\n" +
		"bin/check-local\n" +
		"bin/release-local\n" +
		"go.mod\n" +
		"internal/version.go\n" +
		"lib.go\n" +
		"scripts/dependency-audit\n" +
		"==> Checking whether proxy.golang.org already serves v1.2.3\n" +
		"Dry run finished. Nothing was published, tagged or released.\n"
	if !strings.Contains(stdout, listing) || !strings.HasSuffix(stdout, files) {
		t.Fatalf("dry run output lacks the module listing\n%s", stdout)
	}
	if tags := fixture.git(fixture.repo, "ls-remote", "--tags", "origin"); tags != "" {
		t.Fatalf("dry run pushed a tag\n%s", tags)
	}
	if tags := fixture.git(fixture.repo, "tag", "--list"); tags != "" {
		t.Fatalf("dry run created a local tag\n%s", tags)
	}
	calls := fixture.recorded()
	for _, call := range []string{"check-local\n", "dependency-audit\n", "/@v/list\n"} {
		if !strings.Contains(calls, call) {
			t.Fatalf("dry run skipped %q\n%s", call, calls)
		}
	}
	if strings.Contains(calls, "gh release") {
		t.Fatalf("dry run touched a GitHub release\n%s", calls)
	}
}

func TestReleaseLocalTagsAndReleases(t *testing.T) {
	fixture := newReleaseFixture(t)
	revision := fixture.git(fixture.repo, "rev-parse", "HEAD")

	code, stdout, stderr := fixture.release()
	if code != 0 || !strings.HasSuffix(stdout, "Released v1.2.3.\n") {
		t.Fatalf("release exit %d\n%s\n%s", code, stdout, stderr)
	}
	if kind := fixture.git(fixture.origin, "cat-file", "-t", "refs/tags/v1.2.3"); kind != "tag" {
		t.Fatalf("origin tag v1.2.3 is a %s, want an annotated tag", kind)
	}
	if tagged := fixture.git(fixture.origin, "rev-parse", "refs/tags/v1.2.3^{commit}"); tagged != revision {
		t.Fatalf("origin tag v1.2.3 points at %s, want %s", tagged, revision)
	}
	calls := fixture.recorded()
	if !strings.Contains(calls, "gh release create v1.2.3 --repo Xquik-dev/x-twitter-scraper-go --verify-tag --title v1.2.3 --notes-file ") {
		t.Fatalf("release did not create the GitHub release\n%s", calls)
	}
	notes, err := os.ReadFile(fixture.notes)
	if err != nil {
		t.Fatal(err)
	}
	if want := "\n### Features\n\n* add retweet timestamps\n\n"; string(notes) != want {
		t.Fatalf("release notes = %q, want %q", notes, want)
	}
}

func TestReleaseLocalRejectsUnknownOption(t *testing.T) {
	fixture := newReleaseFixture(t)

	code, stdout, stderr := fixture.release("--dryrun")
	want := "Unknown option --dryrun. Run bin/release-local or bin/release-local --dry-run.\n"
	if code != 2 || stdout != "" || stderr != want {
		t.Fatalf("exit %d, stdout %q, stderr %q", code, stdout, stderr)
	}
	if calls := fixture.recorded(); calls != "" {
		t.Fatalf("an unknown option still ran steps\n%s", calls)
	}
	if tags := fixture.git(fixture.repo, "ls-remote", "--tags", "origin"); tags != "" {
		t.Fatalf("an unknown option pushed a tag\n%s", tags)
	}
}
