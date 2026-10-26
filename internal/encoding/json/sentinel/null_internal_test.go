// SPDX-FileCopyrightText: 2026 Xquik contributors
//
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

func TestNullSentinelReplacesACachedValueOfAnotherType(t *testing.T) {
	type seeded []byte
	nullCache.Lock()
	nullCache.m[reflect.TypeFor[seeded]()] = reflect.ValueOf(1)
	nullCache.Unlock()
	null := NewNullSentinel(func() seeded { return make(seeded, 0, 1) })
	if null == nil || !IsNull(null) || !IsNull(NewNullSentinel(func() seeded { return nil })) {
		t.Fatalf("sentinel = %#v was not recognized as null", null)
	}
}
