// Copyright the Xquik contributors.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"fmt"
	"os/exec"
)

func checkLOC(parent string) error {
	untracked, err := git(exec.Command("git", "ls-files", "--others", "--exclude-standard", "--", ":/"))
	if err != nil {
		return err
	}
	if untracked != "" {
		return fmt.Errorf("LOC check requires new files in the index; review them and use git add --intent-to-add")
	}
	head, err := git(exec.Command("git", "rev-parse", "--verify", "HEAD"))
	if err != nil {
		return err
	}
	command := exec.Command("git", "diff", "--shortstat", "--no-ext-diff", "--no-textconv", "--find-renames", "HEAD", "--", ":/",
		":(top,exclude)aliases.go", ":(top,exclude)shared/shared.go", ":(top,exclude)xtweet.go", ":(top,exclude).stats.yml", ":(top,exclude)go.mod", ":(top,exclude)go.sum", ":(top,exclude)*.md", ":(top,exclude)*.txt", ":(top,exclude)LICENSES/**")
	if parent != head && parent != "HEAD" {
		previous, err := git(exec.Command("git", "rev-parse", "--verify", "HEAD^"))
		if err != nil {
			return err
		}
		if parent != previous && parent != "HEAD^" {
			return fmt.Errorf("LOC comparison requires the current revision or its parent")
		}
		command = exec.Command("git", "diff", "--shortstat", "--no-ext-diff", "--no-textconv", "--find-renames", "HEAD^", "--", ":/",
			":(top,exclude)aliases.go", ":(top,exclude)shared/shared.go", ":(top,exclude)xtweet.go", ":(top,exclude).stats.yml", ":(top,exclude)go.mod", ":(top,exclude)go.sum", ":(top,exclude)*.md", ":(top,exclude)*.txt", ":(top,exclude)LICENSES/**")
	}
	stats, err := git(command)
	if err != nil {
		return err
	}
	fmt.Printf("Git LOC statistics against %s: %s\n", parent, stats)
	return nil
}
