// Copyright the Xquik contributors.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"bytes"
	"os"
	"os/exec"
	"strings"
	"testing"
)

func codeqlCommandFixture(t *testing.T) *os.Root {
	t.Helper()
	directory, err := os.OpenRoot(t.TempDir())
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if err := directory.Close(); err != nil {
			t.Fatal(err)
		}
	})
	build := exec.Command("go", "install", "-cover", "-covermode=set", "-coverpkg=./...", ".")
	build.Env = append(os.Environ(), "GOBIN="+directory.Name())
	if output, err := build.CombinedOutput(); err != nil {
		t.Fatalf("build validator: %v\n%s", err, output)
	}
	counters := os.Getenv("XQUIK_GO_COVERAGE_DIR")
	if counters == "" {
		counters = t.TempDir()
	}
	t.Setenv("GOCOVERDIR", counters)
	for name, source := range map[string]string{
		"value.go":      "package fixture\nfunc Value() int { return 1 }\n",
		"value_test.go": "package fixture\nimport \"testing\"\nfunc TestValue(t *testing.T) { if Value() != 1 { t.Fatal(\"wrong value\") } }\n",
	} {
		if err := directory.WriteFile(name, []byte(source), 0o600); err != nil {
			t.Fatal(err)
		}
	}
	for _, command := range []*exec.Cmd{exec.Command("git", "init", "-q"), exec.Command("git", "add", "value.go", "value_test.go")} {
		command.Dir = directory.Name()
		if output, err := command.CombinedOutput(); err != nil {
			t.Fatalf("prepare source inventory: %v\n%s", err, output)
		}
	}
	return directory
}

const completeCodeqlFixture = `{"version":"2.1.0","runs":[{"results":[],"invocations":[{"executionSuccessful":true,"toolExecutionNotifications":[{"level":"none","descriptor":{"id":"go/diagnostics/successfully-extracted-files"},"locations":[{"physicalLocation":{"artifactLocation":{"uri":"value.go","uriBaseId":"%SRCROOT%"}}},{"physicalLocation":{"artifactLocation":{"uri":"value_test.go","uriBaseId":"%SRCROOT%"}}}]}]}]}]}`

func TestCommandRejectsIncompleteSourceExtraction(t *testing.T) {
	for _, scenario := range []struct{ name, report string }{
		{"missing test source", strings.Replace(completeCodeqlFixture, `,{"physicalLocation":{"artifactLocation":{"uri":"value_test.go","uriBaseId":"%SRCROOT%"}}}`, "", 1)},
		{"wrong source", strings.Replace(completeCodeqlFixture, "value_test.go", "other_test.go", 1)},
		{"wrong source root", strings.ReplaceAll(completeCodeqlFixture, "%SRCROOT%", "OTHER")},
		{"no locations", `{"version":"2.1.0","runs":[{"results":[],"invocations":[{"executionSuccessful":true,"toolExecutionNotifications":[{"level":"none","descriptor":{"id":"go/diagnostics/successfully-extracted-files"}}]}]}]}`},
	} {
		t.Run(scenario.name, func(t *testing.T) {
			directory := codeqlCommandFixture(t)
			command := exec.Command("./codeql-security")
			command.Dir = directory.Name()
			command.Stdin = strings.NewReader(scenario.report)
			var stdout, stderr bytes.Buffer
			command.Stdout, command.Stderr = &stdout, &stderr
			err := command.Run()
			if exit, ok := err.(*exec.ExitError); !ok || exit.ExitCode() != 1 {
				t.Fatalf("incomplete source exit: %v\n%s\n%s", err, stdout.String(), stderr.String())
			}
			if stdout.Len() != 0 || !strings.Contains(stderr.String(), "source inventory") {
				t.Fatalf("incomplete source result: stdout=%q stderr=%q", stdout.String(), stderr.String())
			}
		})
	}
}

func TestCommandAcceptsCompleteSourceExtraction(t *testing.T) {
	directory := codeqlCommandFixture(t)
	command := exec.Command("./codeql-security")
	command.Dir = directory.Name()
	command.Stdin = strings.NewReader(completeCodeqlFixture)
	output, err := command.CombinedOutput()
	if err != nil || string(output) != "CodeQL completed without findings; extracted Go files: 2\n" {
		t.Fatalf("complete source result: %q, %v", output, err)
	}
}
