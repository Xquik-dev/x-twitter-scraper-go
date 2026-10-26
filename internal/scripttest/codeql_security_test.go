// SPDX-FileCopyrightText: 2026 Xquik contributors
//
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
	"testing"
)

const (
	extractionDiagnostic = `{"id":"go/diagnostics/extraction-errors"}`
	extractedFile        = `{"locations":[{"physicalLocation":{"artifactLocation":{"uri":"client.go","uriBaseId":"%SRCROOT%"}}}],"message":{"text":""},"level":"none","descriptor":{"id":"go/diagnostics/successfully-extracted-files"}}`
	unreadFile           = `{"message":{"text":"Extraction failed in client.go with error undefined: missingName"},"level":"error","descriptor":{"id":"go/diagnostics/extraction-errors"}}`
	codeqlCalls          = `version --format=terse
pack install codeql/no-code-comments
pack install $REPORT/query-tests/no-code-comments-test --additional-packs=$REPORT/query-tests
test run $REPORT/query-tests/no-code-comments-test --additional-packs=$REPORT/query-tests --threads=4
test run saw extract tests true and 0 fixture files
database create $REPORT/database --language=go --build-mode=autobuild --source-root=$REPO --threads=4
database analyze $REPORT/database codeql/go-queries@1.6.12:codeql-suites/go-security-extended.qls codeql/no-code-comments --format=sarifv2.1.0 --output=$REPORT/report.sarif --threads=4
`
)

func codeqlReport(diagnostics string, notifications ...string) string {
	return `{"version":"2.1.0","runs":[{"tool":{"driver":{"name":"CodeQL","notifications":[` + diagnostics +
		`]}},"invocations":[{"toolExecutionNotifications":[` + strings.Join(notifications, ",") +
		`],"executionSuccessful":true}],"results":[]}]}`
}

func runCodeQLSecurity(t *testing.T, report string) (auditResult, string) {
	t.Helper()
	root := t.TempDir()
	codeql, err := filepath.Abs("testdata/codeql")
	if err != nil {
		t.Fatal(err)
	}
	reportPath := filepath.Join(root, "report.sarif")
	if err := os.WriteFile(reportPath, []byte(report), 0o600); err != nil {
		t.Fatal(err)
	}
	callsPath := filepath.Join(root, "calls.log")
	command := exec.Command("../../scripts/codeql-security")
	command.Env = append(os.Environ(), "XQUIK_CODEQL_BIN="+codeql, "FAKE_REPORT="+reportPath, "FAKE_CALLS="+callsPath, "TMPDIR="+root)
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
	calls, err := fs.ReadFile(os.DirFS(root), "calls.log")
	if err != nil {
		t.Fatal(err)
	}
	return result, string(calls)
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
			result, calls := runCodeQLSecurity(t, test.report)
			if result.code != test.code {
				t.Fatalf("exit code %d, want %d\nstderr: %s", result.code, test.code, result.stderr)
			}
			if calls != codeqlCalls {
				t.Fatalf("codeql calls\n%s\nwant\n%s", calls, codeqlCalls)
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
