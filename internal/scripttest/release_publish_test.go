// Copyright the Xquik contributors.
// SPDX-License-Identifier: Apache-2.0

package scripttest

import (
	"bytes"
	"errors"
	"os"
	"os/exec"
	"strings"
	"testing"
)

func TestReleasePublishRefusesBeforeRunningChecks(t *testing.T) {
	for _, test := range []struct {
		name, want string
		prepare    func(fixture *releaseFixture)
	}{
		{
			name: "a commit missing from main",
			want: "Merge the release into main first.\n",
			prepare: func(fixture *releaseFixture) {
				fixture.git(fixture.repo, "commit", "--quiet", "--allow-empty", "--message", "unmerged")
				fixture.git(fixture.repo, "tag", "v1.2.3")
			},
		},
		{
			name:    "a missing tag",
			want:    "Create the matching immutable release tag first.\n",
			prepare: func(*releaseFixture) {},
		},
		{
			name: "a missing signing identity",
			want: "XQUIK_RELEASE_IDENTITY: Set the authorized signing identity.\n",
			prepare: func(fixture *releaseFixture) {
				fixture.git(fixture.repo, "tag", "v1.2.3")
			},
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			fixture := newReleaseFixture(t)
			script, err := os.ReadFile("../../scripts/release")
			if err != nil {
				t.Fatal(err)
			}
			repository, err := os.OpenRoot(fixture.repo)
			if err != nil {
				t.Fatal(err)
			}
			defer repository.Close()
			if err := repository.WriteFile("scripts/release", script, 0o600); err != nil {
				t.Fatal(err)
			}
			fixture.git(fixture.repo, "add", "scripts/release")
			fixture.git(fixture.repo, "commit", "--quiet", "--message", "release script")
			fixture.git(fixture.repo, "push", "--quiet", "origin", "main")
			test.prepare(fixture)

			command := exec.Command("bash", "scripts/release", "publish")
			command.Dir = fixture.repo
			command.Env = append(fixture.env, "XQUIK_RELEASE_IDENTITY=", "XQUIK_RELEASE_ISSUER=")
			var stdout, stderr bytes.Buffer
			command.Stdout, command.Stderr = &stdout, &stderr
			code := 0
			if err := command.Run(); err != nil {
				var exit *exec.ExitError
				if !errors.As(err, &exit) {
					t.Fatal(err)
				}
				code = exit.ExitCode()
			}
			if code != 1 || stdout.String() != "{\"status\":\"failed\",\"exitCode\":1}\n" || !strings.HasSuffix(stderr.String(), test.want) {
				t.Fatalf("exit %d, stdout %q, stderr %q, want stderr ending in %q", code, stdout.String(), stderr.String(), test.want)
			}
			if calls := fixture.recorded(); calls != "" {
				t.Fatalf("publish ran steps before it refused\n%s", calls)
			}
		})
	}
}
