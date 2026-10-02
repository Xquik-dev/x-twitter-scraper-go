// Copyright the Xquik contributors.
// SPDX-License-Identifier: Apache-2.0

package json

import (
	"errors"
	"reflect"
	"strconv"
	"testing"
)

func TestUnmarshalRejectsTextKeysOutsideByteRange(t *testing.T) {
	for _, input := range []string{`{"u-1":7}`, `{"u256":7}`} {
		t.Run(input, func(t *testing.T) {
			got := map[u8marshal]int{9: 8}
			err := Unmarshal([]byte(input), &got)
			if !errors.Is(err, strconv.ErrRange) {
				t.Fatalf("out-of-range text key error: %v", err)
			}
			if want := (map[u8marshal]int{9: 8}); !reflect.DeepEqual(got, want) {
				t.Fatalf("out-of-range text key changed the map: got %v, want %v", got, want)
			}
		})
	}
}

func TestUnmarshalPreservesBoundedTextKeySyntax(t *testing.T) {
	for _, test := range []struct {
		input string
		want  map[u8marshal]int
	}{
		{`{"u0":1}`, map[u8marshal]int{0: 1}},
		{`{"u255":2}`, map[u8marshal]int{255: 2}},
		{`{"u+1":3}`, map[u8marshal]int{1: 3}},
		{`{"u-0":4}`, map[u8marshal]int{0: 4}},
	} {
		t.Run(test.input, func(t *testing.T) {
			var got map[u8marshal]int
			if err := Unmarshal([]byte(test.input), &got); err != nil {
				t.Fatalf("bounded text key error: %v", err)
			}
			if !reflect.DeepEqual(got, test.want) {
				t.Fatalf("bounded text keys: got %v, want %v", got, test.want)
			}
		})
	}
}
