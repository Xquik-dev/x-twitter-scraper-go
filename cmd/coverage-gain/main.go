// Copyright the Xquik contributors.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/Xquik-dev/x-twitter-scraper-go/internal/coveragegate"
)

func git(command *exec.Cmd) (string, error) {
	command.Env = append(os.Environ(), "LC_ALL=C")
	data, err := command.Output()
	return strings.TrimSpace(string(data)), err
}

func publishSummary(reports *os.Root, report coveragegate.Report) error {
	if _, err := reports.Stat("summary.json"); err != nil && !os.IsNotExist(err) {
		return err
	}
	pending, err := reports.OpenFile(".summary-pending.json", os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0o600)
	if err != nil {
		return err
	}
	encoder := json.NewEncoder(pending)
	encoder.SetIndent("", "  ")
	if err := errors.Join(encoder.Encode(report), pending.Close()); err != nil {
		return err
	}
	return reports.Rename(".summary-pending.json", "summary.json")
}

func run(directory string) error {
	status, err := git(exec.Command("git", "status", "--porcelain", "--untracked-files=all"))
	if err != nil {
		return err
	}
	parentCommand := exec.Command("git", "rev-parse", "--verify", "HEAD^")
	if status != "" {
		parentCommand = exec.Command("git", "rev-parse", "--verify", "HEAD")
	}
	parent, err := git(parentCommand)
	if err != nil {
		return err
	}
	common, err := git(exec.Command("git", "rev-parse", "--git-common-dir"))
	if err != nil {
		return err
	}
	scope := coveragegate.Scope()
	baseline, err := coveragegate.Baseline(os.DirFS(filepath.Join(common, "coverage-baselines", coveragegate.InstrumentationScope, parent)), parent, scope)
	if err != nil {
		return fmt.Errorf("read verified coverage for parent %s: %w", parent, err)
	}
	reports, err := os.OpenRoot(directory)
	if err != nil {
		return err
	}
	defer reports.Close()
	metrics, err := coveragegate.Measure(reports.FS())
	if err != nil {
		return err
	}
	if err := publishSummary(reports, coveragegate.Report{Scope: scope, Total: metrics}); err != nil {
		return err
	}
	if err := coveragegate.Compare(baseline.Total, metrics); err != nil {
		return err
	}
	return checkLOC(parent)
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "usage: coverage-gain REPORT_DIRECTORY | --record | --collect-statements | --collect-branches")
		os.Exit(1)
	}
	if os.Args[1] == "--collect-statements" || os.Args[1] == "--collect-branches" {
		if err := coveragegate.Collect(os.Args[1], os.Stderr); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		return
	}
	if os.Args[1] == "--record" {
		report, err := coveragegate.Record(".", os.Stderr)
		if err == nil {
			err = json.NewEncoder(os.Stdout).Encode(report)
		}
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		return
	}
	if err := run(os.Args[1]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println("All four Go coverage metrics preserve parent coverage. SDK LOC changes are advisory.")
}
