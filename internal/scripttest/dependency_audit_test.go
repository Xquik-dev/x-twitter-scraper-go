// Copyright the Xquik contributors.
// SPDX-License-Identifier: Apache-2.0

package scripttest

import (
	"bytes"
	"encoding/json"
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

const fakeGo = `#!/bin/sh
case "$1 $2" in
  "list -deps") cat "$FAKE_SHIPPED_MODULES" ;;
  "list -f") cat "$FAKE_PACKAGES" ;;
  "run golang.org/x/vuln/cmd/govulncheck@v1.9.0") cat "$FAKE_REPORT" ;;
  *) echo "unexpected go call $*" >&2; exit 97 ;;
esac
`

const (
	shippedModule     = "example.com/shipped"
	developmentModule = "example.com/tooling"
	failed1           = "Dependency audit failed with 1 blocking finding(s). For call traces, run go run golang.org/x/vuln/cmd/govulncheck@v1.9.0 -test ./...\n"
)

type frame struct {
	Module   string `json:"module"`
	Version  string `json:"version,omitempty"`
	Package  string `json:"package,omitempty"`
	Function string `json:"function,omitempty"`
}

type finding struct {
	OSV          string  `json:"osv"`
	FixedVersion string  `json:"fixed_version,omitempty"`
	Trace        []frame `json:"trace"`
}

type exception struct {
	ID      string `json:"id"`
	Module  string `json:"module"`
	Reason  string `json:"reason"`
	Recheck string `json:"recheck"`
}

func day(offset int) string {
	return time.Now().UTC().AddDate(0, 0, offset).Format("2006-01-02")
}

func required(id, module, fixed string) finding {
	return finding{OSV: id, FixedVersion: fixed, Trace: []frame{{Module: module, Version: "v1.0.0"}}}
}

func called(id, module, fixed string) finding {
	return finding{OSV: id, FixedVersion: fixed, Trace: []frame{
		{Module: module, Version: "v1.0.0", Package: module + "/parse", Function: "Parse"},
		{Module: "github.com/Xquik-dev/x-twitter-scraper-go", Package: "github.com/Xquik-dev/x-twitter-scraper-go/cmd/tool", Function: "main"},
	}}
}

func listed(recheck string) exception {
	return exception{ID: "GO-2026-0001", Module: developmentModule, Reason: "build-time only", Recheck: recheck}
}

type auditResult struct {
	code           int
	stdout, stderr string
}

func runAudit(t *testing.T, goCommand string, findings []finding, exceptions any) auditResult {
	t.Helper()
	script, err := filepath.Abs("../../scripts/dependency-audit")
	if err != nil {
		t.Fatal(err)
	}
	root := t.TempDir()
	write := func(name string, data []byte, mode os.FileMode) string {
		t.Helper()
		path := filepath.Join(root, name)
		if err := os.MkdirAll(filepath.Dir(path), 0o750); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(path, data, mode); err != nil {
			t.Fatal(err)
		}
		return path
	}
	encode := func(value any) []byte {
		t.Helper()
		data, err := json.Marshal(value)
		if err != nil {
			t.Fatal(err)
		}
		return data
	}
	report := []byte(`{"config":{"scanner_name":"govulncheck","scan_level":"symbol"}}` + "\n" + `{"progress":{"message":"Scanning your code"}}` + "\n")
	for _, item := range findings {
		report = append(report, encode(map[string]finding{"finding": item})...)
		report = append(report, '\n')
	}
	write("bin/go", []byte(goCommand), 0o755)
	write("audit-exceptions.json", encode(exceptions), 0o644)
	command := exec.Command(script, root)
	command.Env = append(os.Environ(),
		"PATH="+filepath.Join(root, "bin")+string(os.PathListSeparator)+os.Getenv("PATH"),
		"FAKE_REPORT="+write("report.json", report, 0o644),
		"FAKE_PACKAGES="+write("packages.txt", []byte("github.com/Xquik-dev/x-twitter-scraper-go\n\ngithub.com/Xquik-dev/x-twitter-scraper-go/internal/tooling\n"), 0o644),
		"FAKE_SHIPPED_MODULES="+write("modules.txt", []byte("\ngithub.com/Xquik-dev/x-twitter-scraper-go\n"+shippedModule+"\n"), 0o644),
	)
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

func TestDependencyAudit(t *testing.T) {
	for _, test := range []struct {
		name       string
		findings   []finding
		exceptions any
		want       auditResult
	}{
		{
			name:       "passes when govulncheck reports nothing",
			exceptions: []exception{},
			want:       auditResult{stdout: "Dependency audit passed.\n"},
		},
		{
			name:       "blocks a shipped finding that no code calls",
			findings:   []finding{required("GO-2026-0001", shippedModule, "v1.0.1")},
			exceptions: []exception{},
			want: auditResult{
				code:   1,
				stdout: "Blocked. GO-2026-0001 is in example.com/shipped, which ships to callers. Upgrade to v1.0.1.\n",
				stderr: failed1,
			},
		},
		{
			name:       "blocks a shipped finding without a fix, even when it is listed",
			findings:   []finding{required("GO-2026-0001", shippedModule, "")},
			exceptions: []exception{{ID: "GO-2026-0001", Module: shippedModule, Reason: "build-time only", Recheck: day(10)}},
			want: auditResult{
				code:   1,
				stdout: "Blocked. GO-2026-0001 is in example.com/shipped, which ships to callers. It has no fix yet, so replace the module. Remove its audit-exceptions.json entry.\n",
				stderr: failed1,
			},
		},
		{
			name:       "blocks a standard library finding",
			findings:   []finding{required("GO-2026-0001", "stdlib", "v1.27.3")},
			exceptions: []exception{},
			want: auditResult{
				code:   1,
				stdout: "Blocked. GO-2026-0001 is in stdlib, which ships to callers. Upgrade to v1.27.3.\n",
				stderr: failed1,
			},
		},
		{
			name:       "blocks a development finding that has a fix, even when it is listed",
			findings:   []finding{called("GO-2026-0001", developmentModule, "v1.0.1")},
			exceptions: []exception{listed(day(10))},
			want: auditResult{
				code: 1,
				stdout: "Blocked. GO-2026-0001 in example.com/tooling is fixed in v1.0.1. Upgrade it.\n" +
					"audit-exceptions.json lists GO-2026-0001 for example.com/tooling, which no longer applies. Remove it.\n",
				stderr: failed1,
			},
		},
		{
			name:       "blocks an unlisted development finding without a fix",
			findings:   []finding{required("GO-2026-0001", developmentModule, "")},
			exceptions: []exception{},
			want: auditResult{
				code:   1,
				stdout: "Blocked. GO-2026-0001 in example.com/tooling has no fix. Review it and list it in audit-exceptions.json.\n",
				stderr: failed1,
			},
		},
		{
			name: "accepts a listed development finding without a fix, once for all its traces",
			findings: []finding{
				required("GO-2026-0001", developmentModule, ""),
				called("GO-2026-0001", developmentModule, ""),
			},
			exceptions: []exception{listed(day(30))},
			want: auditResult{
				stdout: "Accepted the reviewed exception for GO-2026-0001 in example.com/tooling.\n" +
					"Dependency audit passed with 1 reviewed exception.\n",
			},
		},
		{
			name:       "blocks a listed finding on its recheck date",
			findings:   []finding{required("GO-2026-0001", developmentModule, "")},
			exceptions: []exception{listed(day(0))},
			want: auditResult{
				code:   1,
				stdout: "Blocked. The exception for GO-2026-0001 in example.com/tooling expired on " + day(0) + ". Recheck the finding.\n",
				stderr: failed1,
			},
		},
		{
			name:       "blocks a recheck date more than 30 days out",
			findings:   []finding{required("GO-2026-0001", developmentModule, "")},
			exceptions: []exception{listed(day(31))},
			want: auditResult{
				code:   1,
				stdout: "Blocked. The exception for GO-2026-0001 in example.com/tooling rechecks on " + day(31) + ". Use " + day(30) + " or earlier.\n",
				stderr: failed1,
			},
		},
		{
			name:       "blocks an exception with an unknown reason",
			findings:   []finding{required("GO-2026-0001", developmentModule, "")},
			exceptions: []exception{{ID: "GO-2026-0001", Module: developmentModule, Reason: "low severity", Recheck: day(10)}},
			want: auditResult{
				code:   1,
				stdout: "Blocked. The exception for GO-2026-0001 in example.com/tooling needs a \"reason\" of \"build-time only\" or \"peer dependency provided by the host\".\n",
				stderr: failed1,
			},
		},
		{
			name:       "blocks an exceptions file that is not a list of entries",
			exceptions: map[string]any{"exceptions": []exception{}},
			want: auditResult{
				code:   1,
				stdout: "Blocked. audit-exceptions.json must hold an array of entries with \"id\", \"module\", \"reason\" and \"recheck\".\n",
				stderr: failed1,
			},
		},
		{
			name:       "passes and names an exception that matches no finding",
			exceptions: []exception{listed(day(10))},
			want: auditResult{
				stdout: "audit-exceptions.json lists GO-2026-0001 for example.com/tooling, which no longer applies. Remove it.\n" +
					"Dependency audit passed.\n",
			},
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			got := runAudit(t, fakeGo, test.findings, test.exceptions)
			if got != test.want {
				t.Fatalf("audit result differs\ngot  code %d\nstdout %s\nstderr %s\nwant code %d\nstdout %s\nstderr %s",
					got.code, got.stdout, got.stderr, test.want.code, test.want.stdout, test.want.stderr)
			}
		})
	}
}

func TestDependencyAuditFailsWhenGoFails(t *testing.T) {
	got := runAudit(t, "#!/bin/sh\nexit 1\n", nil, []exception{})
	if got.code == 0 || strings.Contains(got.stdout, "Dependency audit passed") {
		t.Fatalf("audit passed although go failed\ncode %d\nstdout %s\nstderr %s", got.code, got.stdout, got.stderr)
	}
}
