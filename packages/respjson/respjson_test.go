// Copyright the Xquik contributors.
// SPDX-License-Identifier: Apache-2.0

package respjson

import "testing"

func TestFieldStates(t *testing.T) {
	tests := map[string]struct {
		field Field
		valid bool
		raw   string
	}{
		"omitted": {field: Field{}, valid: false, raw: Omitted},
		"null":    {field: NewField(Null), valid: false, raw: Null},
		"invalid": {field: NewInvalidField("{"), valid: false, raw: "{"},
		"valid":   {field: NewField(`"tweet"`), valid: true, raw: `"tweet"`},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			if got := test.field.Valid(); got != test.valid {
				t.Fatalf("Valid() = %t, want %t", got, test.valid)
			}
			if got := test.field.Raw(); got != test.raw {
				t.Fatalf("Raw() = %q, want %q", got, test.raw)
			}
		})
	}
}
