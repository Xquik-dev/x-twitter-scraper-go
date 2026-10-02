// Copyright the Xquik contributors.
// SPDX-License-Identifier: Apache-2.0

package coveragegate_test

import (
	"testing"

	"github.com/Xquik-dev/x-twitter-scraper-go/internal/coveragegate"
)

func TestComparePreservesEachMetricRatioOrUncoveredCount(t *testing.T) {
	for _, test := range []struct {
		name    string
		current coveragegate.Metric
		accept  bool
	}{
		{"same ratio with more uncovered items", coveragegate.Metric{Total: 2000, Covered: 1900}, true},
		{"same uncovered count", coveragegate.Metric{Total: 900, Covered: 850}, true},
		{"fewer uncovered items", coveragegate.Metric{Total: 900, Covered: 851}, true},
		{"added uncovered item", coveragegate.Metric{Total: 900, Covered: 849}, false},
	} {
		for _, metric := range []string{"statements", "branches", "functions", "lines"} {
			t.Run(test.name+"/"+metric, func(t *testing.T) {
				parent, current := map[string]coveragegate.Metric{}, map[string]coveragegate.Metric{}
				for _, name := range []string{"statements", "branches", "functions", "lines"} {
					parent[name] = coveragegate.Metric{Total: 1000, Covered: 950}
					current[name] = coveragegate.Metric{Total: 1000, Covered: 950}
				}
				current[metric] = test.current
				err := coveragegate.Compare(parent, current)
				if (err == nil) != test.accept {
					t.Fatalf("accept = %v, want %v; error: %v", err == nil, test.accept, err)
				}
			})
		}
	}
}
