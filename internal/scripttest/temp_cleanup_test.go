// Copyright the Xquik contributors.
// SPDX-License-Identifier: Apache-2.0

package scripttest

import (
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"syscall"
	"testing"
	"time"
)

const fakeTool = `#!/bin/sh
case "$*" in
  "version --format=terse") echo 2.27.2; exit 0 ;;
  "list -f "*|"list -deps "*) echo github.com/Xquik-dev/x-twitter-scraper-go; exit 0 ;;
esac
if [ -n "$GOMODCACHE" ]; then
  mkdir -p "$GOMODCACHE/cache/download" && : >"$GOMODCACHE/cache/download/list" && chmod -R a-w "$GOMODCACHE"
fi
if [ -n "$FAKE_STARTED" ]; then
  : >"$FAKE_STARTED"
  exec sleep 30
fi
echo "fake $(basename "$0") failed" >&2
exit 3
`

type cleanupCase struct {
	name    string
	command []string
	tools   []string
	setup   func(t *testing.T, root string) []string
}

func repositoryRoot(t *testing.T) string {
	t.Helper()
	root, err := filepath.Abs("../..")
	if err != nil {
		t.Fatal(err)
	}
	return root
}

func cleanupCases(t *testing.T) []cleanupCase {
	repo := repositoryRoot(t)
	script := func(path string) string { return filepath.Join(repo, path) }
	return []cleanupCase{
		{name: "check-local", command: []string{script("bin/check-local"), "--no-status"}, tools: []string{"go", "jq", "uv"}},
		{name: "collect-coverage", command: []string{script("scripts/collect-coverage")}, tools: []string{"go"}},
		{name: "collect-branches", command: []string{script("scripts/collect-branches")}, tools: []string{"go"}},
		{name: "coverage-baseline", command: []string{script("scripts/coverage-baseline"), "HEAD"}, tools: []string{"go"}},
		{name: "codeql-security", command: []string{script("scripts/codeql-security")}, tools: []string{"codeql"}, setup: func(t *testing.T, root string) []string {
			return []string{"XQUIK_CODEQL_BIN=" + filepath.Join(root, "bin", "codeql")}
		}},
		{name: "source-security", command: []string{script("scripts/source-security")}, tools: []string{"go"}},
		{name: "signing-audit", command: []string{"bash", script("scripts/signing-audit")}, tools: []string{"uv"}},
		{name: "dependency-audit", command: []string{script("scripts/dependency-audit")}, tools: []string{"go"}, setup: func(t *testing.T, root string) []string {
			project := filepath.Join(root, "project")
			if err := os.MkdirAll(project, 0o750); err != nil {
				t.Fatal(err)
			}
			return []string{"FAKE_ARGUMENT=" + project}
		}},
		{name: "reproducible-build", command: []string{script("scripts/reproducible-build")}, tools: []string{"go"}},
		{name: "verify-release", command: []string{"bash", script("scripts/verify-release"), "v1.2.3", "0123456789abcdef0123456789abcdef01234567"}, tools: []string{"go"}},
	}
}

func startCleanupCase(t *testing.T, test cleanupCase, started string) (*exec.Cmd, string) {
	t.Helper()
	root := t.TempDir()
	private := filepath.Join(root, "tmp")
	bin := filepath.Join(root, "bin")
	for _, directory := range []string{private, bin} {
		if err := os.MkdirAll(directory, 0o750); err != nil {
			t.Fatal(err)
		}
	}
	for _, tool := range test.tools {
		if err := os.WriteFile(filepath.Join(bin, tool), []byte(fakeTool), 0o755); err != nil {
			t.Fatal(err)
		}
	}
	env := []string{"PATH=" + bin + string(os.PathListSeparator) + os.Getenv("PATH"), "TMPDIR=" + private, "FUZZ_TIME=1s"}
	if started != "" {
		env = append(env, "FAKE_STARTED="+started)
	}
	command := append([]string(nil), test.command...)
	if test.setup != nil {
		for _, variable := range test.setup(t, root) {
			if argument, ok := strings.CutPrefix(variable, "FAKE_ARGUMENT="); ok {
				command = append(command, argument)
				continue
			}
			env = append(env, variable)
		}
	}
	cmd := exec.Command(command[0], command[1:]...)
	cmd.Env = append(os.Environ(), env...)
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	if err := cmd.Start(); err != nil {
		t.Fatal(err)
	}
	return cmd, private
}

func exitCode(t *testing.T, err error) int {
	t.Helper()
	var exit *exec.ExitError
	if err != nil && !errors.As(err, &exit) {
		t.Fatal(err)
	}
	if exit == nil {
		return 0
	}
	return exit.ExitCode()
}

func expectEmpty(t *testing.T, private string) {
	t.Helper()
	entries, err := os.ReadDir(private)
	if err != nil {
		t.Fatal(err)
	}
	var left []string
	for _, entry := range entries {
		left = append(left, entry.Name())
	}
	if len(left) > 0 {
		t.Fatalf("the run left %s in its TMPDIR", strings.Join(left, ", "))
	}
}

func TestChecksRemoveTheirTempDirsWhenAStepFails(t *testing.T) {
	for _, test := range cleanupCases(t) {
		t.Run(test.name, func(t *testing.T) {
			cmd, private := startCleanupCase(t, test, "")
			if code := exitCode(t, cmd.Wait()); code == 0 {
				t.Fatal("the failing step did not fail the run")
			}
			expectEmpty(t, private)
		})
	}
}

func TestChecksRemoveTheirTempDirsWhenTerminated(t *testing.T) {
	for _, test := range cleanupCases(t) {
		t.Run(test.name, func(t *testing.T) {
			started := filepath.Join(t.TempDir(), "started")
			cmd, private := startCleanupCase(t, test, started)
			deadline := time.Now().Add(20 * time.Second)
			for {
				if _, err := os.Stat(started); err == nil {
					break
				}
				if time.Now().After(deadline) {
					syscall.Kill(-cmd.Process.Pid, syscall.SIGKILL)
					cmd.Wait()
					t.Fatal("the run never reached its fake tool")
				}
				time.Sleep(20 * time.Millisecond)
			}
			if err := syscall.Kill(-cmd.Process.Pid, syscall.SIGTERM); err != nil {
				t.Fatal(err)
			}
			if code := exitCode(t, cmd.Wait()); code != 143 {
				t.Fatalf("exit code %d, want 143", code)
			}
			expectEmpty(t, private)
		})
	}
}

func startRelease(t *testing.T, fixture *releaseFixture, checkLocal string, args ...string) (*exec.Cmd, string) {
	t.Helper()
	if checkLocal != "" {
		if err := os.WriteFile(filepath.Join(fixture.repo, "bin", "check-local"), []byte(checkLocal), 0o755); err != nil {
			t.Fatal(err)
		}
		fixture.git(fixture.repo, "commit", "--quiet", "--all", "--message", "check-local")
		fixture.git(fixture.repo, "push", "--quiet", "origin", "main")
	}
	private := filepath.Join(t.TempDir(), "tmp")
	if err := os.MkdirAll(private, 0o750); err != nil {
		t.Fatal(err)
	}
	cmd := exec.Command(filepath.Join(fixture.repo, "bin", "release-local"), args...)
	cmd.Dir = t.TempDir()
	cmd.Env = append(append([]string(nil), fixture.env...), "TMPDIR="+private)
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	if err := cmd.Start(); err != nil {
		t.Fatal(err)
	}
	return cmd, private
}

func TestReleaseLocalRemovesItsTempDir(t *testing.T) {
	t.Run("after a dry run with a read-only module cache", func(t *testing.T) {
		cmd, private := startRelease(t, newReleaseFixture(t), "", "--dry-run")
		if code := exitCode(t, cmd.Wait()); code != 0 {
			t.Fatalf("dry run exit %d", code)
		}
		expectEmpty(t, private)
	})
	t.Run("when a step fails", func(t *testing.T) {
		cmd, private := startRelease(t, newReleaseFixture(t), "#!/bin/sh\nexit 3\n", "--dry-run")
		if code := exitCode(t, cmd.Wait()); code != 3 {
			t.Fatalf("exit %d, want 3", code)
		}
		expectEmpty(t, private)
	})
	t.Run("when terminated", func(t *testing.T) {
		started := filepath.Join(t.TempDir(), "started")
		cmd, private := startRelease(t, newReleaseFixture(t), "#!/bin/sh\n: >\""+started+"\"\nexec sleep 30\n", "--dry-run")
		deadline := time.Now().Add(20 * time.Second)
		for {
			if _, err := os.Stat(started); err == nil {
				break
			}
			if time.Now().After(deadline) {
				syscall.Kill(-cmd.Process.Pid, syscall.SIGKILL)
				cmd.Wait()
				t.Fatal("the release never reached bin/check-local")
			}
			time.Sleep(20 * time.Millisecond)
		}
		if err := syscall.Kill(-cmd.Process.Pid, syscall.SIGTERM); err != nil {
			t.Fatal(err)
		}
		if code := exitCode(t, cmd.Wait()); code != 143 {
			t.Fatalf("exit code %d, want 143", code)
		}
		expectEmpty(t, private)
	})
}
