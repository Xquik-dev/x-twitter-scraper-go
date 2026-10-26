// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

package reflectvalue

import (
	"reflect"
	"testing"
	"time"
)

type stamp time.Time

func TestAsConvertsToTheStaticType(t *testing.T) {
	want := time.Date(2025, time.January, 15, 12, 0, 0, 0, time.UTC)
	got, err := As[time.Time](reflect.ValueOf(stamp(want)))
	if err != nil || !got.Equal(want) {
		t.Fatalf("As = %v, %v, want %v", got, err, want)
	}
}

func TestAsReportsAValueThatCannotConvert(t *testing.T) {
	got, err := As[time.Time](reflect.ValueOf(42))
	if err == nil || !got.IsZero() {
		t.Fatalf("As = %v, %v, want the zero time and an error", got, err)
	}
}
