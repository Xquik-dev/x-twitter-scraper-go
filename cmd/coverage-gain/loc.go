// Copyright the Xquik contributors.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"fmt"
)

func checkLOC(parent string) error {
	untracked, err := git("ls-files", "--others", "--exclude-standard", "--", ":/")
	if err != nil {
		return err
	}
	if untracked != "" {
		return fmt.Errorf("LOC check requires new files in the index; review them and use git add --intent-to-add")
	}
	stats, err := git("diff", "--shortstat", "--no-ext-diff", "--no-textconv", "--find-renames", parent, "--", ":/",
		":(top,exclude)aliases.go", ":(top,exclude)shared/shared.go", ":(top,exclude)xtweet.go", ":(top,exclude).stats.yml", ":(top,exclude)go.mod", ":(top,exclude)go.sum", ":(top,exclude)*.md", ":(top,exclude)*.txt", ":(top,exclude)LICENSES/**")
	if err != nil {
		return err
	}
	fmt.Printf("Git LOC statistics against %s: %s\n", parent, stats)
	return nil
}
