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

const recordingRunner = `#!/bin/sh
printf '%s|' "$@"
exit 3
`

type checkLocalRun struct {
	code           int
	stdout, stderr string
}

func runCheckLocal(t *testing.T, env ...string) checkLocalRun {
	t.Helper()
	script, err := os.ReadFile("../../bin/check-local")
	if err != nil {
		t.Fatal(err)
	}
	root := t.TempDir()
	checkout := filepath.Join(root, "checkout")
	if err := os.MkdirAll(filepath.Join(checkout, "bin"), 0o750); err != nil {
		t.Fatal(err)
	}
	for name, data := range map[string]string{"checkout/bin/check-local": string(script), "runner": recordingRunner} {
		if err := os.WriteFile(filepath.Join(root, name), []byte(data), 0o600); err != nil {
			t.Fatal(err)
		}
	}
	command := exec.Command("bash", "-c", `chmod 700 "$RUNNER" && exec bash "$CHECKOUT/bin/check-local" "$@"`, "check-local-test", "--bogus", "two words")
	command.Dir = t.TempDir()
	command.Env = append(os.Environ(),
		"CHECKOUT="+checkout,
		"RUNNER="+filepath.Join(root, "runner"),
		"CHECK_LOCAL_HERE=",
		"GIT_CONFIG_COUNT=1",
		"GIT_CONFIG_KEY_0=checklocal.runner",
		"GIT_CONFIG_VALUE_0="+filepath.Join(root, "runner"),
	)
	command.Env = append(command.Env, env...)
	var out, failure bytes.Buffer
	command.Stdout, command.Stderr = &out, &failure
	run := checkLocalRun{}
	if err := command.Run(); err != nil {
		var exit *exec.ExitError
		if !errors.As(err, &exit) {
			t.Fatal(err)
		}
		run.code = exit.ExitCode()
	}
	resolved, err := filepath.EvalSymlinks(checkout)
	if err != nil {
		t.Fatal(err)
	}
	run.stdout = strings.ReplaceAll(out.String(), resolved, "<checkout>")
	run.stderr = failure.String()
	return run
}

func TestCheckLocalHandsItsRunToTheConfiguredRunner(t *testing.T) {
	got := runCheckLocal(t)
	want := checkLocalRun{code: 3, stdout: "<checkout>|--bogus|two words|"}
	if got != want {
		t.Fatalf("got %+v, want %+v", got, want)
	}
}

func TestCheckLocalRunsHereWithoutARunnerOrWithCheckLocalHere(t *testing.T) {
	for name, env := range map[string][]string{
		"no runner":          {"GIT_CONFIG_COUNT=0"},
		"CHECK_LOCAL_HERE=1": {"CHECK_LOCAL_HERE=1"},
	} {
		t.Run(name, func(t *testing.T) {
			got := runCheckLocal(t, env...)
			want := checkLocalRun{
				code:   2,
				stderr: "Unknown option --bogus. Run bin/check-local or bin/check-local --no-status.\n",
			}
			if got != want {
				t.Fatalf("got %+v, want %+v", got, want)
			}
		})
	}
}
