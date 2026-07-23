// Copyright the Xquik contributors.
// SPDX-License-Identifier: Apache-2.0

package coveragebranch

import (
	"strings"
	"testing"
)

func TestMeasureMatchesPositiveOverlaps(t *testing.T) {
	report := []byte(`{
		"packages": [{
			"import_path": "example.test/sdk",
			"files": [{
				"name": "client.go",
				"blocks": [
					{"idx": 1, "line": 10, "col": 2, "end_line": 12, "end_col": 3},
					{"idx": 2, "line": 20, "col": 2, "end_line": 22, "end_col": 3},
					{"idx": 3, "line": 30, "col": 2, "end_line": 32, "end_col": 3}
				],
				"branches": [
					{"block_idx": 1},
					{"block_idx": 2},
					{"block_idx": 3}
				]
			}]
		}]
	}`)
	profile := strings.NewReader(
		"mode: set\n" +
			"example.test/sdk/client.go:10.2,10.3 1 1\n" +
			"example.test/sdk/client.go:20.1,20.2 1 1\n" +
			"example.test/sdk/client.go:30.2,31.1 1 0\n",
	)

	covered, total, err := Measure(report, profile)
	if err != nil {
		t.Fatal(err)
	}
	if covered != 1 || total != 3 {
		t.Fatalf("Measure() = %d/%d, want 1/3", covered, total)
	}
}

func TestMeasureRejectsInvalidInputs(t *testing.T) {
	tests := map[string]struct {
		report  string
		profile string
	}{
		"invalid report JSON": {
			report:  `{"packages":[]`,
			profile: "mode: set\n",
		},
		"invalid profile line": {
			report:  `{"packages":[]}`,
			profile: "not a coverage profile\n",
		},
		"missing branch block": {
			report: `{
				"packages": [{
					"import_path": "example.test/sdk",
					"files": [{
						"name": "client.go",
						"blocks": [],
						"branches": [{"block_idx": 9}]
					}]
				}]
			}`,
			profile: "mode: set\n",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			_, _, err := Measure(
				[]byte(test.report),
				strings.NewReader(test.profile),
			)
			if err == nil {
				t.Fatal("Measure() accepted invalid input")
			}
		})
	}
}
