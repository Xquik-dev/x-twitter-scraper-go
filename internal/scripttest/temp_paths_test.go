// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

package scripttest

import (
	"errors"
	"os"
	"os/exec"
	"strings"
	"testing"
)

func TestTempPathTrimsTheTrailingSlashOfTMPDIR(t *testing.T) {
	for name, suffix := range map[string]string{"with a trailing slash": "/", "without one": ""} {
		t.Run(name, func(t *testing.T) {
			root := t.TempDir()
			cmd := exec.Command("bash", "-c", `source scripts/temp-paths && temp_path made probe && printf '%s\n%s\n' "$made" "$GOTMPDIR"`)
			cmd.Dir = repositoryRoot(t)
			cmd.Env = append(os.Environ(), "TMPDIR="+root+suffix)
			out, err := cmd.CombinedOutput()
			if err != nil {
				t.Fatalf("%v\n%s", err, out)
			}
			made, goTmpDir, _ := strings.Cut(strings.TrimSuffix(string(out), "\n"), "\n")
			if !strings.HasPrefix(made, root+"/xquik-probe.") || strings.Contains(strings.TrimPrefix(made, root), "//") || goTmpDir != made {
				t.Fatalf("temp_path made %q and GOTMPDIR %q under TMPDIR %q", made, goTmpDir, root+suffix)
			}
			if _, err := os.Stat(made); !errors.Is(err, os.ErrNotExist) {
				t.Fatalf("the exit trap left %s, stat error %v", made, err)
			}
		})
	}
}
