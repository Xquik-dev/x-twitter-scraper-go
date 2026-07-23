// Copyright the Xquik contributors.
// SPDX-License-Identifier: Apache-2.0

package sentinel

import (
	"reflect"
	"testing"
)

func TestNullSentinelCacheAndUnsupportedKinds(t *testing.T) {
	created := 0
	first := NewNullSentinel(func() []int {
		created++
		return make([]int, 0, 1)
	})
	second := NewNullSentinel(func() []int {
		created++
		return make([]int, 0, 1)
	})

	if created != 1 {
		t.Fatalf("sentinel factory ran %d times, want 1", created)
	}
	if !IsNull(first) || !IsNull(second) {
		t.Fatal("cached sentinels were not recognized")
	}
	if IsNull(1) {
		t.Fatal("integer was recognized as a null sentinel")
	}
	if IsValueNull(reflect.ValueOf(1)) {
		t.Fatal("integer value was recognized as a null sentinel")
	}
}
