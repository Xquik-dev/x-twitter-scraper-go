// Copyright the Xquik contributors.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"

	"github.com/Xquik-dev/x-twitter-scraper-go/internal/coveragegate"
)

func git(args ...string) (string, error) {
	command := exec.Command("git", args...)
	command.Env = append(os.Environ(), "LC_ALL=C")
	data, err := command.Output()
	return strings.TrimSpace(string(data)), err
}

func run(directory string) error {
	status, err := git("status", "--porcelain", "--untracked-files=all")
	if err != nil {
		return err
	}
	ref := "HEAD^"
	if status != "" {
		ref = "HEAD"
	}
	parent, err := git("rev-parse", "--verify", ref)
	if err != nil {
		return err
	}
	common, err := git("rev-parse", "--git-common-dir")
	if err != nil {
		return err
	}
	scope := instrumentationScope()
	baseline, err := coveragegate.Baseline(os.DirFS(filepath.Join(common, "coverage-baselines", parent)), parent, scope)
	if err != nil {
		return fmt.Errorf("read verified coverage for parent %s: %w", parent, err)
	}
	metrics, err := coveragegate.Measure(os.DirFS(directory))
	if err != nil {
		return err
	}
	data, err := json.MarshalIndent(coveragegate.Report{Scope: scope, Total: metrics}, "", "  ")
	if err != nil {
		return err
	}
	if err := os.WriteFile(filepath.Join(directory, "summary.json"), data, 0o600); err != nil {
		return err
	}
	if err := coveragegate.Compare(baseline.Total, metrics); err != nil {
		return err
	}
	return checkLOC(parent)
}

func main() {
	switch {
	case len(os.Args) == 4 && os.Args[1] == "-record":
		if err := record(os.Args[2], os.Args[3]); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		fmt.Printf("Recorded the verified coverage of %s.\n", os.Args[2])
	case len(os.Args) == 2:
		if err := run(os.Args[1]); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		fmt.Println("All four Go coverage metrics preserve parent coverage. SDK LOC changes are advisory.")
	default:
		fmt.Fprintln(os.Stderr, "usage: coverage-gain REPORT_DIRECTORY, or coverage-gain -record REVISION REPORT_DIRECTORY")
		os.Exit(1)
	}
}

func instrumentationScope() string {
	return runtime.Version() + "/" + runtime.GOOS + "/" + runtime.GOARCH + "/go-packages-v1"
}

// A full SHA-1 or SHA-256 commit id, so a revision cannot name a path.
var commitID = regexp.MustCompile(`^[0-9a-f]{40}(?:[0-9a-f]{24})?$`)

// record measures the reports that the collectors wrote for revision and
// keeps them as that revision's verified baseline, where run reads it.
func record(revision, directory string) error {
	if !commitID.MatchString(revision) {
		return fmt.Errorf("revision %q is not a full commit id", revision)
	}
	common, err := git("rev-parse", "--path-format=absolute", "--git-common-dir")
	if err != nil {
		return err
	}
	metrics, err := coveragegate.Measure(os.DirFS(directory))
	if err != nil {
		return err
	}
	data, err := json.MarshalIndent(coveragegate.Report{Revision: revision, Scope: instrumentationScope(), Total: metrics}, "", "  ")
	if err != nil {
		return err
	}
	target := filepath.Join(common, "coverage-baselines", revision)
	if err := os.MkdirAll(target, 0o750); err != nil {
		return err
	}
	return os.WriteFile(filepath.Join(target, "report.json"), data, 0o600)
}
