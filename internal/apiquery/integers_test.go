// Copyright the Xquik contributors.
// SPDX-License-Identifier: Apache-2.0

package apiquery

import "testing"

func TestEightBitIntegers(t *testing.T) {
	values, err := Marshal(map[string]any{"max": uint8(255), "min": int8(-128), "zero": int8(0)})
	if err != nil {
		t.Fatal(err)
	}
	if got := values.Encode(); got != "max=255&min=-128&zero=0" {
		t.Fatalf("eight-bit query = %q", got)
	}
}
