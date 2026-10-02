// Copyright the Xquik contributors.
// SPDX-License-Identifier: Apache-2.0

package coveragegate_test

import (
	"bytes"
	"testing"

	"github.com/Xquik-dev/x-twitter-scraper-go/internal/coveragegate"
)

func TestMergeNativeCoverageIncludesUnitChildAndUnexecutedBlocks(t *testing.T) {
	unit := []byte("mode: set\nfixture/a.go:1.1,2.1 1 1\nfixture/b.go:1.1,2.1 1 0\nfixture/unused.go:1.1,2.1 1 0\n")
	child := []byte("mode: set\nfixture/a.go:1.1,2.1 1 0\nfixture/b.go:1.1,2.1 1 1\n")
	got, err := coveragegate.MergeNativeCoverage(unit, child)
	want := []byte("mode: set\nfixture/a.go:1.1,2.1 1 1\nfixture/b.go:1.1,2.1 1 1\nfixture/unused.go:1.1,2.1 1 0\n")
	if err != nil || !bytes.Equal(got, want) {
		t.Fatalf("merged native coverage: %q, %v; want %q", got, err, want)
	}
}

func TestMergeNativeCoveragePreservesUnitsWithoutChildren(t *testing.T) {
	unit := []byte("mode: set\nfixture/a.go:1.1,2.1 1 1\nfixture/unused.go:1.1,2.1 1 0\n")
	got, err := coveragegate.MergeNativeCoverage(unit, nil)
	if err != nil || !bytes.Equal(got, unit) {
		t.Fatalf("unit-only native coverage: %q, %v; want %q", got, err, unit)
	}
}

func TestMergeNativeCoverageRefusesDifferentInstrumentation(t *testing.T) {
	const unit = "mode: set\nfixture/a.go:1.1,2.1 1 0\n"
	for _, test := range []struct {
		name, unit, child string
	}{
		{"empty unit catalog", "mode: set\n", ""},
		{"unit mode", "mode: count\nfixture/a.go:1.1,2.1 1 0\n", ""},
		{"child mode", unit, "mode: atomic\nfixture/a.go:1.1,2.1 1 1\n"},
		{"malformed unit", "mode: set\ninvalid\n", ""},
		{"malformed child", unit, "mode: set\ninvalid\n"},
		{"foreign source", unit, "mode: set\nother/a.go:1.1,2.1 1 1\n"},
		{"foreign range", unit, "mode: set\nfixture/a.go:1.1,3.1 1 1\n"},
		{"changed statement count", unit, "mode: set\nfixture/a.go:1.1,2.1 2 1\n"},
		{"nonbinary unit counter", "mode: set\nfixture/a.go:1.1,2.1 1 2\n", ""},
		{"nonbinary child counter", unit, "mode: set\nfixture/a.go:1.1,2.1 1 2\n"},
	} {
		t.Run(test.name, func(t *testing.T) {
			if result, err := coveragegate.MergeNativeCoverage([]byte(test.unit), []byte(test.child)); err == nil || len(result) != 0 {
				t.Fatalf("accepted different instrumentation: %q, %v", result, err)
			}
		})
	}
}
