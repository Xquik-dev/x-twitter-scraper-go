// Copyright the Xquik contributors.
// SPDX-License-Identifier: Apache-2.0

package coveragebranch

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"testing"
	"testing/iotest"
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

func TestReadProfileRejectsOverflow(t *testing.T) {
	overflow := strings.Repeat("9", 100)
	for index, name := range []string{"start line", "start column", "end line", "end column", "count"} {
		t.Run(name, func(t *testing.T) {
			fields := []string{"1", "2", "3", "4", "1"}
			fields[index] = overflow
			profile := fmt.Sprintf("client.go:%s.%s,%s.%s 1 %s\n", fields[0], fields[1], fields[2], fields[3], fields[4])
			entries, err := readProfile(strings.NewReader(profile))
			if entries != nil || !errors.Is(err, strconv.ErrRange) || !strings.Contains(err.Error(), overflow) {
				t.Fatalf("readProfile() = %v, %v; want wrapped overflow with value", entries, err)
			}
		})
	}
}

func TestReadProfileSkipsUncoveredCoordinates(t *testing.T) {
	profile := "client.go:" + strings.Repeat("9", 100) + ".1,2.3 1 0\n"
	entries, err := readProfile(strings.NewReader(profile))
	if err != nil || len(entries) != 0 {
		t.Fatalf("readProfile() = %v, %v; want empty profile", entries, err)
	}
}

func TestReadProfilePreservesReadError(t *testing.T) {
	failure := errors.New("profile unavailable")
	entries, err := readProfile(iotest.ErrReader(failure))
	if entries != nil || !errors.Is(err, failure) || err.Error() != "read profile: profile unavailable" {
		t.Fatalf("readProfile() = %v, %v; want wrapped read error", entries, err)
	}
}
