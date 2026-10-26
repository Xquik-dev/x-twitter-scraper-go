// Copyright the Xquik contributors.
// SPDX-License-Identifier: Apache-2.0

// Package coveragegate measures statement, branch, function and line coverage,
// and fails when any of them falls below the parent revision's.
package coveragegate

import (
	"fmt"
)

type Metric struct {
	Total, Covered, Skipped int64
}

func Compare(parent, current map[string]Metric) error {
	for _, name := range []string{"statements", "branches", "functions", "lines"} {
		percentages := [2]float64{}
		for i, report := range []map[string]Metric{parent, current} {
			value, exists := report[name]
			if !exists || value.Total <= 0 || value.Covered < 0 || value.Covered > value.Total || value.Skipped != 0 {
				return fmt.Errorf("invalid %s counts in report %d", name, i)
			}
			percentages[i] = 100 * float64(value.Covered) / float64(value.Total)
		}
		required := percentages[0]
		if percentages[1]+1e-10 < required {
			return fmt.Errorf("%s coverage %.8f%% must reach %.8f%%", name, percentages[1], required)
		}
	}
	return nil
}
