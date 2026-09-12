// Copyright the Xquik contributors.
// SPDX-License-Identifier: Apache-2.0

package coveragegate

import (
	"strings"
	"testing"
)

func TestCompare(t *testing.T) {
	for _, test := range []struct {
		name            string
		parent, current Metric
		missing         bool
		failure         string
	}{
		{"exact gain below old floors", Metric{1000, 500, 0}, Metric{1000, 501, 0}, false, ""},
		{"gain below 0.1 passes", Metric{10000, 5000, 0}, Metric{10000, 5009, 0}, false, ""},
		{"below old branch floor", Metric{1000, 0, 0}, Metric{1000, 799, 0}, false, ""},
		{"at old branch floor", Metric{1000, 0, 0}, Metric{1000, 800, 0}, false, ""},
		{"below old statement floor", Metric{1000, 0, 0}, Metric{1000, 899, 0}, false, ""},
		{"at old statement floor", Metric{1000, 0, 0}, Metric{1000, 900, 0}, false, ""},
		{"unchanged passes", Metric{1000, 950, 0}, Metric{1000, 950, 0}, false, ""},
		{"finish coverage", Metric{10000, 9999, 0}, Metric{10000, 10000, 0}, false, ""},
		{"keep complete", Metric{1000, 1000, 0}, Metric{1000, 1000, 0}, false, ""},
		{"partial regression", Metric{10000, 5000, 0}, Metric{10000, 4999, 0}, false, "must reach"},
		{"complete regression", Metric{10000, 10000, 0}, Metric{10000, 9999, 0}, false, "must reach"},
		{"missing metric", Metric{1000, 950, 0}, Metric{1000, 951, 0}, true, "invalid"},
		{"empty denominator", Metric{1000, 950, 0}, Metric{0, 0, 0}, false, "invalid"},
		{"negative denominator", Metric{1000, 950, 0}, Metric{-1, 0, 0}, false, "invalid"},
		{"invalid covered", Metric{1000, 950, 0}, Metric{1000, 1001, 0}, false, "invalid"},
		{"negative covered", Metric{1000, 950, 0}, Metric{1000, -1, 0}, false, "invalid"},
		{"skipped code", Metric{1000, 950, 0}, Metric{1000, 951, 1}, false, "invalid"},
		{"invalid parent", Metric{0, 0, 0}, Metric{1000, 951, 0}, false, "invalid"},
	} {
		for _, name := range []string{"statements", "branches", "functions", "lines"} {
			t.Run(test.name+"/"+name, func(t *testing.T) {
				parent, current := map[string]Metric{}, map[string]Metric{}
				for _, key := range []string{"statements", "branches", "functions", "lines"} {
					parent[key], current[key] = Metric{1000, 950, 0}, Metric{1000, 951, 0}
				}
				parent[name], current[name] = test.parent, test.current
				if test.missing {
					delete(current, name)
				}
				err := Compare(parent, current)
				if test.failure == "" {
					if err != nil {
						t.Fatal(err)
					}
				} else if err == nil || !strings.Contains(err.Error(), test.failure) || !strings.Contains(err.Error(), name) {
					t.Fatalf("expected %s failure for %s, got %v", test.failure, name, err)
				}
			})
		}
	}
}
