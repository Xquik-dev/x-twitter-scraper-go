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

const fakeCodeQL = `#!/bin/sh
case "$1 $2" in
  "version --format=terse") echo "${FAKE_CODEQL_VERSION:-2.27.2}" ;;
  "pack install") ;;
  "database create") ;;
  "database analyze")
    for argument; do
      case "$argument" in --output=*) cp "$FAKE_REPORT" "${argument#--output=}" ;; esac
    done ;;
  *) echo "unexpected codeql call $*" >&2; exit 97 ;;
esac
`

const (
	extractionDiagnostic = `{"id":"go/diagnostics/extraction-errors"}`
	extractedFile        = `{"locations":[{"physicalLocation":{"artifactLocation":{"uri":"client.go","uriBaseId":"%SRCROOT%"}}}],"message":{"text":""},"level":"none","descriptor":{"id":"go/diagnostics/successfully-extracted-files"}}`
	unreadFile           = `{"message":{"text":"Extraction failed in client.go with error undefined: missingName"},"level":"error","descriptor":{"id":"go/diagnostics/extraction-errors"}}`
)

func codeqlReport(diagnostics string, notifications ...string) string {
	return `{"version":"2.1.0","runs":[{"tool":{"driver":{"name":"CodeQL","notifications":[` + diagnostics +
		`]}},"invocations":[{"toolExecutionNotifications":[` + strings.Join(notifications, ",") +
		`],"executionSuccessful":true}],"results":[]}]}`
}

func runCodeQLSecurity(t *testing.T, report string, environment ...string) auditResult {
	t.Helper()
	script, err := filepath.Abs("../../scripts/codeql-security")
	if err != nil {
		t.Fatal(err)
	}
	root := t.TempDir()
	codeql := filepath.Join(root, "codeql")
	reportPath := filepath.Join(root, "report.sarif")
	if err := os.WriteFile(codeql, []byte(fakeCodeQL), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(reportPath, []byte(report), 0o600); err != nil {
		t.Fatal(err)
	}
	command := exec.Command(script)
	command.Env = append(os.Environ(), "XQUIK_CODEQL_BIN="+codeql, "FAKE_REPORT="+reportPath, "TMPDIR="+root)
	command.Env = append(command.Env, environment...)
	var stdout, stderr bytes.Buffer
	command.Stdout, command.Stderr = &stdout, &stderr
	result := auditResult{}
	if err := command.Run(); err != nil {
		var exit *exec.ExitError
		if !errors.As(err, &exit) {
			t.Fatal(err)
		}
		result.code = exit.ExitCode()
	}
	result.stdout, result.stderr = stdout.String(), stderr.String()
	return result
}

func TestCodeQLSecurity(t *testing.T) {
	for _, test := range []struct {
		name, report string
		code         int
		stdout       string
		stderr       string
	}{
		{
			name:   "a scan that read every file passes",
			report: codeqlReport(extractionDiagnostic, extractedFile),
			stdout: "CodeQL completed without findings and extracted 1 files.\n",
		},
		{
			name:   "a file CodeQL could not read fails the scan and is named",
			report: codeqlReport(extractionDiagnostic, unreadFile, extractedFile),
			code:   1,
			stderr: `CodeQL could not read all Go code, so the scan is incomplete. It said "Extraction failed in client.go with error undefined: missingName". Fix that code so Go compiles it, then rerun bin/check-local.` + "\n",
		},
		{
			name:   "a scan without the extraction diagnostic fails",
			report: codeqlReport(`{"id":"go/diagnostics/successfully-extracted-files"}`, extractedFile),
			code:   1,
			stderr: "The CodeQL scan lacks go/diagnostics/extraction-errors, so it cannot show that it read every file. Make the suite in scripts/codeql-security run that diagnostic query.\n",
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			result := runCodeQLSecurity(t, test.report)
			if result.code != test.code {
				t.Fatalf("exit code %d, want %d\nstderr: %s", result.code, test.code, result.stderr)
			}
			if want := test.report + test.stdout; result.stdout != want {
				t.Fatalf("stdout %q, want %q", result.stdout, want)
			}
			if !strings.Contains(result.stderr, test.stderr) {
				t.Fatalf("stderr %q does not contain %q", result.stderr, test.stderr)
			}
		})
	}
}

func TestCodeQLSecurityVersionFloor(t *testing.T) {
	report := codeqlReport(extractionDiagnostic, extractedFile)
	for _, test := range []struct {
		name, version string
		code          int
		stderr        string
	}{
		{name: "the minimum CodeQL runs the scan", version: "2.27.2"},
		{name: "a newer CodeQL runs the scan", version: "2.27.10"},
		{
			name:    "an older CodeQL is refused before the scan",
			version: "2.27.1",
			code:    1,
			stderr:  "CodeQL 2.27.1 is older than 2.27.2. Install CodeQL 2.27.2 or newer as CONTRIBUTING.md says, or set XQUIK_CODEQL_BIN.\n",
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			result := runCodeQLSecurity(t, report, "FAKE_CODEQL_VERSION="+test.version)
			if result.code != test.code {
				t.Fatalf("exit code %d, want %d\nstderr: %s", result.code, test.code, result.stderr)
			}
			if test.code == 0 && !strings.Contains(result.stdout, "CodeQL completed without findings") {
				t.Fatalf("stdout %q lacks the scan result", result.stdout)
			}
			if test.stderr != "" && result.stderr != test.stderr {
				t.Fatalf("stderr %q, want %q", result.stderr, test.stderr)
			}
		})
	}
}
