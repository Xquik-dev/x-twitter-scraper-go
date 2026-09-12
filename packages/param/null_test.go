// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

package param_test

import (
	"github.com/Xquik-dev/x-twitter-scraper-go/packages/param"
	"testing"
)

type Nullables struct {
	Slice []int          `json:"slice,omitzero"`
	Map   map[string]int `json:"map,omitzero"`
	param.APIObject
}

func (n Nullables) MarshalJSON() ([]byte, error) {
	type shadow Nullables
	return param.MarshalObject(n, (*shadow)(&n))
}

func TestNullMarshal(t *testing.T) {
	assertJSON(t, Nullables{}, `{}`)

	obj := Nullables{
		Slice: param.NullSlice[[]int](),
		Map:   param.NullMap[map[string]int](),
	}
	assertJSON(t, obj, `{"slice":null,"map":null}`)

	if !param.IsNull(obj.Slice) {
		t.Fatal("failed null check")
	}
	if !param.IsNull(obj.Map) {
		t.Fatal("failed null check")
	}

}
