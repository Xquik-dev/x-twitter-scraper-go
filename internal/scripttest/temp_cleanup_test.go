// Copyright the Xquik contributors.
// SPDX-License-Identifier: Apache-2.0

package scripttest

import (
	"bytes"
	"errors"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"syscall"
	"testing"
	"time"
)

type cleanupCase struct {
	name    string
	command func(root string) *exec.Cmd
	tools   []string
	setup   func(t *testing.T, root string) []string
}

func cleanupCases() []cleanupCase {
	return []cleanupCase{
		{name: "check-local", command: func(string) *exec.Cmd { return exec.Command("../../bin/check-local", "--no-status") }, tools: []string{"go", "jq", "uv"}},
		{name: "test", command: func(string) *exec.Cmd { return exec.Command("../../scripts/test") }, tools: []string{"go"}},
		{name: "fuzz", command: func(string) *exec.Cmd { return exec.Command("../../scripts/fuzz") }, tools: []string{"go"}},
		{name: "lint", command: func(string) *exec.Cmd { return exec.Command("../../scripts/lint") }, tools: []string{"go", "gofmt"}},
		{name: "collect-coverage", command: func(string) *exec.Cmd { return exec.Command("../../scripts/collect-coverage") }, tools: []string{"go"}},
		{name: "collect-branches", command: func(string) *exec.Cmd { return exec.Command("../../scripts/collect-branches") }, tools: []string{"go"}},
		{name: "coverage-baseline", command: func(string) *exec.Cmd { return exec.Command("../../scripts/coverage-baseline", "HEAD") }, tools: []string{"go"}},
		{name: "codeql-security", command: func(string) *exec.Cmd { return exec.Command("../../scripts/codeql-security") }, tools: []string{"codeql"}, setup: func(t *testing.T, root string) []string {
			return []string{"XQUIK_CODEQL_BIN=" + filepath.Join(root, "bin", "codeql")}
		}},
		{name: "source-security", command: func(string) *exec.Cmd { return exec.Command("../../scripts/source-security") }, tools: []string{"go"}},
		{name: "signing-audit", command: func(string) *exec.Cmd { return exec.Command("bash", "../../scripts/signing-audit") }, tools: []string{"uv"}},
		{name: "dependency-audit", command: func(root string) *exec.Cmd {
			command := exec.Command("scripts/dependency-audit", ".")
			command.Dir = filepath.Join(root, "project")
			return command
		}, tools: []string{"go"}, setup: func(t *testing.T, root string) []string {
			project := filepath.Join(root, "project")
			if err := os.MkdirAll(project, 0o750); err != nil {
				t.Fatal(err)
			}
			linkAuditScript(t, project)
			return nil
		}},
		{name: "reproducible-build", command: func(string) *exec.Cmd { return exec.Command("../../scripts/reproducible-build") }, tools: []string{"go"}},
		{name: "verify-release", command: func(string) *exec.Cmd {
			return exec.Command("bash", "../../scripts/verify-release", "v1.2.3", "0123456789abcdef0123456789abcdef01234567")
		}, tools: []string{"go"}},
	}
}

func startCleanupCase(t *testing.T, test cleanupCase, started string) (*exec.Cmd, string) {
	t.Helper()
	fakeTool, err := filepath.Abs("testdata/fake-tool")
	if err != nil {
		t.Fatal(err)
	}
	root := t.TempDir()
	private := filepath.Join(root, "tmp")
	bin := filepath.Join(root, "bin")
	moduleCache := filepath.Join(root, "modules")
	for _, directory := range []string{private, bin} {
		if err := os.MkdirAll(directory, 0o750); err != nil {
			t.Fatal(err)
		}
	}
	for _, tool := range test.tools {
		if err := os.Symlink(fakeTool, filepath.Join(bin, tool)); err != nil {
			t.Fatal(err)
		}
	}
	env := []string{"PATH=" + bin + string(os.PathListSeparator) + os.Getenv("PATH"), "TMPDIR=" + private, "GOTMPDIR=", "GOMODCACHE=" + moduleCache, "FUZZ_TIME=1s"}
	if started != "" {
		env = append(env, "FAKE_STARTED="+started)
	}
	if test.setup != nil {
		env = append(env, test.setup(t, root)...)
	}
	cmd := test.command(root)
	cmd.Env = append(os.Environ(), env...)
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	var stdout, stderr bytes.Buffer
	cmd.Stdout, cmd.Stderr = &stdout, &stderr
	t.Cleanup(func() {
		cache, err := os.OpenRoot(moduleCache)
		if errors.Is(err, os.ErrNotExist) {
			return
		}
		if err != nil {
			t.Errorf("%v\n%s\n%s", err, cmd.Stdout, cmd.Stderr)
			return
		}
		defer cache.Close()
		for _, directory := range []string{".", "cache", filepath.Join("cache", "download")} {
			if err := cache.Chmod(directory, 0o750); err != nil && !errors.Is(err, os.ErrNotExist) {
				t.Errorf("%v\n%s\n%s", err, cmd.Stdout, cmd.Stderr)
			}
		}
	})
	if err := cmd.Start(); err != nil {
		t.Fatalf("%v\n%s\n%s", err, cmd.Stdout, cmd.Stderr)
	}
	return cmd, private
}

func exitCode(t *testing.T, cmd *exec.Cmd, err error) int {
	t.Helper()
	var exit *exec.ExitError
	if err != nil && !errors.As(err, &exit) {
		t.Fatalf("%v\n%s\n%s", err, cmd.Stdout, cmd.Stderr)
	}
	if exit == nil {
		return 0
	}
	return exit.ExitCode()
}

func expectEmpty(t *testing.T, cmd *exec.Cmd, private string) {
	t.Helper()
	entries, err := os.ReadDir(private)
	if err != nil {
		t.Fatalf("%v\n%s\n%s", err, cmd.Stdout, cmd.Stderr)
	}
	var left []string
	for _, entry := range entries {
		left = append(left, entry.Name())
	}
	if len(left) > 0 {
		t.Fatalf("the run left %s in its TMPDIR\n%s\n%s", strings.Join(left, ", "), cmd.Stdout, cmd.Stderr)
	}
}

func terminateRun(t *testing.T, cmd *exec.Cmd, started string, signal syscall.Signal) {
	t.Helper()
	deadline := time.Now().Add(20 * time.Second)
	for {
		if _, err := os.Stat(started); err == nil {
			break
		}
		if time.Now().After(deadline) {
			killErr := syscall.Kill(-cmd.Process.Pid, syscall.SIGKILL)
			waitErr := cmd.Wait()
			if killErr != nil {
				t.Errorf("%v\n%s\n%s", killErr, cmd.Stdout, cmd.Stderr)
			}
			if waitErr != nil {
				t.Errorf("%v\n%s\n%s", waitErr, cmd.Stdout, cmd.Stderr)
			}
			t.Fatalf("the run never reached its fake tool\n%s\n%s", cmd.Stdout, cmd.Stderr)
		}
		time.Sleep(20 * time.Millisecond)
	}
	if err := syscall.Kill(-cmd.Process.Pid, signal); err != nil {
		waitErr := cmd.Wait()
		t.Fatalf("terminate %v, wait %v\n%s\n%s", err, waitErr, cmd.Stdout, cmd.Stderr)
	}
	code := exitCode(t, cmd, cmd.Wait())
	if signal == syscall.SIGINT {
		status, ok := cmd.ProcessState.Sys().(syscall.WaitStatus)
		if code != -1 || !ok || !status.Signaled() || status.Signal() != syscall.SIGINT {
			t.Fatalf("process state %v, want SIGINT\n%s\n%s", cmd.ProcessState, cmd.Stdout, cmd.Stderr)
		}
		return
	}
	if status, ok := cmd.ProcessState.Sys().(syscall.WaitStatus); ok && status.Signaled() && status.Signal() == signal {
		return
	}
	if want := 128 + int(signal); code != want {
		t.Fatalf("exit code %d, want %d\n%s\n%s", code, want, cmd.Stdout, cmd.Stderr)
	}
}

func TestChecksRemoveTheirTempDirsWhenAStepFails(t *testing.T) {
	for _, test := range cleanupCases() {
		t.Run(test.name, func(t *testing.T) {
			cmd, private := startCleanupCase(t, test, "")
			if code := exitCode(t, cmd, cmd.Wait()); code == 0 {
				t.Fatalf("the failing step did not fail the run\n%s\n%s", cmd.Stdout, cmd.Stderr)
			}
			expectEmpty(t, cmd, private)
		})
	}
}

func TestChecksRemoveTheirTempDirsWhenTerminated(t *testing.T) {
	for _, test := range cleanupCases() {
		t.Run(test.name, func(t *testing.T) {
			for _, signal := range []syscall.Signal{syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM} {
				t.Run(signal.String(), func(t *testing.T) {
					started := filepath.Join(t.TempDir(), "started")
					cmd, private := startCleanupCase(t, test, started)
					terminateRun(t, cmd, started, signal)
					expectEmpty(t, cmd, private)
				})
			}
		})
	}
}

func TestInterruptStopsBashParent(t *testing.T) {
	root := t.TempDir()
	started := filepath.Join(root, "started")
	continued := filepath.Join(root, "continued")
	test := cleanupCase{
		command: func(string) *exec.Cmd {
			return exec.Command("bash", "-c", "\"$1\"\n: >\"$FAKE_CONTINUED\"\n", "bash", "../../scripts/test")
		},
		tools: []string{"go"},
		setup: func(*testing.T, string) []string { return []string{"FAKE_CONTINUED=" + continued} },
	}
	cmd, private := startCleanupCase(t, test, started)
	terminateRun(t, cmd, started, syscall.SIGINT)
	expectEmpty(t, cmd, private)
	if _, err := os.Stat(continued); !errors.Is(err, os.ErrNotExist) {
		t.Fatalf("the bash parent ran its second command, stat %v\n%s\n%s", err, cmd.Stdout, cmd.Stderr)
	}
}

func startRelease(t *testing.T, fixture *releaseFixture, checkLocal string) (*exec.Cmd, string) {
	t.Helper()
	checkLocal = "#!/bin/sh\nmktemp \"${GOTMPDIR:-$TMPDIR}/go-build.XXXXXX\" >/dev/null || exit 1\n" + checkLocal
	if err := os.WriteFile(filepath.Join(fixture.repo, "bin", "check-local"), []byte(checkLocal), 0o600); err != nil {
		t.Fatal(err)
	}
	fixture.git(fixture.repo, exec.Command("git", "commit", "--quiet", "--all", "--message", "check-local"))
	fixture.git(fixture.repo, exec.Command("git", "push", "--quiet", "origin", "main"))
	root := t.TempDir()
	private := filepath.Join(root, "tmp")
	if err := os.MkdirAll(private, 0o750); err != nil {
		t.Fatal(err)
	}
	cmd := exec.Command(releaseLocal, "--dry-run")
	cmd.Dir = fixture.outside
	cmd.Env = append(append([]string(nil), fixture.env...), "TMPDIR="+private, "GOTMPDIR=", "GOMODCACHE="+filepath.Join(root, "modules"))
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	var stdout, stderr bytes.Buffer
	cmd.Stdout, cmd.Stderr = &stdout, &stderr
	if err := cmd.Start(); err != nil {
		t.Fatalf("%v\n%s\n%s", err, cmd.Stdout, cmd.Stderr)
	}
	return cmd, private
}

func TestReleaseLocalRemovesItsTempDir(t *testing.T) {
	t.Run("after a dry run", func(t *testing.T) {
		cmd, private := startRelease(t, newReleaseFixture(t), "")
		if code := exitCode(t, cmd, cmd.Wait()); code != 0 {
			t.Fatalf("dry run exit %d\n%s\n%s", code, cmd.Stdout, cmd.Stderr)
		}
		expectEmpty(t, cmd, private)
	})
	t.Run("when a step fails", func(t *testing.T) {
		cmd, private := startRelease(t, newReleaseFixture(t), "exit 3\n")
		if code := exitCode(t, cmd, cmd.Wait()); code != 3 {
			t.Fatalf("exit %d, want 3\n%s\n%s", code, cmd.Stdout, cmd.Stderr)
		}
		expectEmpty(t, cmd, private)
	})
	for _, signal := range []syscall.Signal{syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM} {
		t.Run("when terminated by "+signal.String(), func(t *testing.T) {
			started := filepath.Join(t.TempDir(), "started")
			cmd, private := startRelease(t, newReleaseFixture(t), ": >\""+started+"\"\nexec sleep 30\n")
			terminateRun(t, cmd, started, signal)
			expectEmpty(t, cmd, private)
		})
	}
}

func TestCoverageReportsOnlyRetainedProfiles(t *testing.T) {
	for _, retain := range []bool{false, true} {
		name := "without an output directory"
		if retain {
			name = "with an output directory"
		}
		t.Run(name, func(t *testing.T) {
			output := ""
			if retain {
				output = t.TempDir()
			}
			test := cleanupCase{
				command: func(string) *exec.Cmd { return exec.Command("../../scripts/collect-coverage") },
				tools:   []string{"go"},
				setup: func(t *testing.T, root string) []string {
					return []string{"FAKE_COVERAGE=true", "XQUIK_COVERAGE_DIR=" + output}
				},
			}
			cmd, private := startCleanupCase(t, test, "")
			if code := exitCode(t, cmd, cmd.Wait()); code != 0 {
				t.Fatalf("exit %d, want 0\n%s\n%s", code, cmd.Stdout, cmd.Stderr)
			}
			want := "==> Collecting statement, function, and line coverage\n"
			if retain {
				want += "Coverage profile: " + filepath.Join(output, "coverage.out") + "\n"
				for _, name := range []string{"coverage.out", "coverage.xml"} {
					data, err := fs.ReadFile(os.DirFS(output), name)
					if err != nil || string(data) != "mode: set\n" {
						t.Fatalf("retained %s = %q, error %v", name, data, err)
					}
				}
			}
			if stdout, ok := cmd.Stdout.(*bytes.Buffer); !ok || stdout.String() != want {
				t.Fatalf("stdout %q, want %q", stdout.String(), want)
			}
			expectEmpty(t, cmd, private)
		})
	}
}
