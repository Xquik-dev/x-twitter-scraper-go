// Copyright the Xquik contributors.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
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
	scope := runtime.Version() + "/" + runtime.GOOS + "/" + runtime.GOARCH + "/go-packages-v1"
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
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "usage: coverage-gain REPORT_DIRECTORY")
		os.Exit(1)
	}
	if err := run(os.Args[1]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println("All four Go coverage metrics preserve parent coverage. SDK LOC changes are advisory.")
}
