// Copyright the Xquik contributors.
// SPDX-License-Identifier: Apache-2.0

package coveragegate

import (
	"bytes"
	"fmt"

	"golang.org/x/tools/cover"
)

type nativeBlock struct {
	file  string
	block cover.ProfileBlock
}

func parseSetCoverage(data []byte) ([]*cover.Profile, error) {
	if !bytes.HasPrefix(data, []byte("mode: set\n")) {
		return nil, fmt.Errorf("native coverage requires set instrumentation")
	}
	profiles, err := cover.ParseProfilesFromReader(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	for _, profile := range profiles {
		for _, block := range profile.Blocks {
			if block.Count != 0 && block.Count != 1 {
				return nil, fmt.Errorf("nonbinary set counter in %s", profile.FileName)
			}
		}
	}
	return profiles, nil
}

func blockIdentity(file string, block cover.ProfileBlock) nativeBlock {
	block.Count = 0
	return nativeBlock{file: file, block: block}
}

// MergeNativeCoverage combines native execution without shrinking the unit catalog.
func MergeNativeCoverage(unit, child []byte) ([]byte, error) {
	profiles, err := parseSetCoverage(unit)
	if err != nil {
		return nil, err
	}
	catalog := make(map[nativeBlock]*cover.ProfileBlock)
	for _, profile := range profiles {
		for index := range profile.Blocks {
			block := &profile.Blocks[index]
			catalog[blockIdentity(profile.FileName, *block)] = block
		}
	}
	if len(catalog) == 0 {
		return nil, fmt.Errorf("native unit coverage has no source blocks")
	}
	if len(child) != 0 {
		children, err := parseSetCoverage(child)
		if err != nil {
			return nil, err
		}
		for _, profile := range children {
			for _, block := range profile.Blocks {
				unitBlock, exists := catalog[blockIdentity(profile.FileName, block)]
				if !exists {
					return nil, fmt.Errorf("child instrumentation differs from unit catalog in %s", profile.FileName)
				}
				unitBlock.Count |= block.Count
			}
		}
	}
	var result bytes.Buffer
	result.WriteString("mode: set\n")
	for _, profile := range profiles {
		for _, block := range profile.Blocks {
			fmt.Fprintf(&result, "%s:%d.%d,%d.%d %d %d\n", profile.FileName, block.StartLine, block.StartCol, block.EndLine, block.EndCol, block.NumStmt, block.Count)
		}
	}
	return result.Bytes(), nil
}
