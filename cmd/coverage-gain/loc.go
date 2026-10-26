// Copyright the Xquik contributors.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"fmt"
	"os/exec"
)

func checkLOC(parent string, dirty bool) error {
	untracked, err := git(exec.Command("git", "ls-files", "--others", "--exclude-standard", "--", ":/"))
	if err != nil {
		return err
	}
	if untracked != "" {
		return fmt.Errorf("LOC check requires new files in the index; review them and use git add --intent-to-add")
	}
	against := "HEAD^"
	if dirty {
		against = "HEAD"
	}
	stats, err := git(exec.Command("git", "diff", "--shortstat", "--no-ext-diff", "--no-textconv", "--find-renames", against, "--", ":/",
		":(top,exclude)aliases.go", ":(top,exclude)shared/shared.go", ":(top,exclude)xtweet.go", ":(top,exclude)go.mod", ":(top,exclude)go.sum", ":(top,exclude)*.md", ":(top,exclude)*.txt", ":(top,exclude)LICENSES/**"))
	if err != nil {
		return err
	}
	fmt.Printf("Git LOC statistics against %s: %s\n", parent, stats)
	return nil
}
