// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

package json

import (
	"errors"
	"reflect"
	"testing"
)

func TestMarshalNumericMapKeys(t *testing.T) {
	for _, tt := range []struct {
		CaseName
		in   func() ([]byte, error)
		want string
	}{
		{Name(""), marshalOf(map[int]string{10: "b", -2: "a"}), `{"-2":"a","10":"b"}`},
		{Name(""), marshalOf(map[int8]bool{-1: true}), `{"-1":true}`},
		{Name(""), marshalOf(map[uint16]int{7: 1, 30: 2}), `{"30":2,"7":1}`},
		{Name(""), marshalOf(map[uintptr]string{0: "zero"}), `{"0":"zero"}`},
	} {
		t.Run(tt.Name, func(t *testing.T) {
			got, err := tt.in()
			if string(got) != tt.want || err != nil {
				t.Errorf("%s: Marshal = %s, %v, want %s", tt.Where, got, err, tt.want)
			}
		})
	}
}

func TestUnsupportedTypes(t *testing.T) {
	for _, tt := range []struct {
		CaseName
		in   func() ([]byte, error)
		want string
	}{
		{Name(""), marshalOf(make(chan int)), "json: unsupported type: chan int"},
		{Name(""), marshalOf(complex(1, 2)), "json: unsupported type: complex128"},
		{Name(""), marshalOf(map[bool]int{true: 1}), "json: unsupported type: map[bool]int"},
		{Name(""), marshalOf(struct{ F func() }{}), "json: unsupported type: func()"},
	} {
		t.Run(tt.Name, func(t *testing.T) {
			got, err := tt.in()
			var unsupported *UnsupportedTypeError
			if got != nil || !errors.As(err, &unsupported) || err.Error() != tt.want {
				t.Errorf("%s: Marshal = %s, %v, want error %q", tt.Where, got, err, tt.want)
			}
		})
	}
}

func TestCallIsZeroReportsOtherTypesAsNonZero(t *testing.T) {
	if callIsZero(reflect.ValueOf(1)) {
		t.Fatal("an int without IsZero was reported as zero")
	}
}
