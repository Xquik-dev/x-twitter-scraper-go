// Copyright the Xquik contributors.
// SPDX-License-Identifier: Apache-2.0

package json

import (
	"errors"
	"testing"
)

type customTimeJSON []byte

func (v customTimeJSON) MarshalJSON() ([]byte, error)            { return []byte(`"fallback"`), nil }
func (v customTimeJSON) MarshalJSONWithTimeLayout(string) []byte { return v }

func TestCustomTimeJSON(t *testing.T) {
	if got, err := Marshal(customTimeJSON(nil)); string(got) != `"fallback"` || err != nil {
		t.Fatalf("nil custom output = %q, %v", got, err)
	}
	for _, test := range []struct{ input, want string }{
		{` "2026-09-12" `, `"2026-09-12"`},
		{`{`, ""},
		{`null trailing`, ""},
	} {
		got, err := Marshal(customTimeJSON(test.input))
		var cause *SyntaxError
		if string(got) != test.want || (err != nil) != (test.want == "") || (test.want == "") != errors.As(err, &cause) {
			t.Errorf("Marshal(%q) = %q, %v", test.input, got, err)
		}
	}
}
