// Copyright the Xquik contributors.
// SPDX-License-Identifier: Apache-2.0

package paramutil

import (
	"reflect"
	"testing"

	"github.com/Xquik-dev/x-twitter-scraper-go/packages/param"
	"github.com/Xquik-dev/x-twitter-scraper-go/packages/respjson"
)

type testUnion struct {
	param.APIUnion
	Text   *string
	Number *int
}

type testObject struct {
	param.APIObject
}

func TestVariantFromUnion(t *testing.T) {
	text := "tweet"
	number := 42
	tests := map[string]struct {
		value   any
		want    any
		wantErr bool
	}{
		"one variant": {
			value: testUnion{Text: &text},
			want:  &text,
		},
		"pointer": {
			value: &testUnion{Number: &number},
			want:  &number,
		},
		"non-struct": {
			value:   "not a union",
			wantErr: true,
		},
		"non-union": {
			value:   struct{ Text *string }{Text: &text},
			wantErr: true,
		},
		"multiple variants": {
			value:   testUnion{Text: &text, Number: &number},
			wantErr: true,
		},
		"no variants": {
			value:   testUnion{},
			wantErr: true,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := VariantFromUnion(reflect.ValueOf(test.value))
			if test.wantErr {
				if err == nil {
					t.Fatalf("VariantFromUnion() = %v", got)
				}
				return
			}
			if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(got, test.want) {
				t.Fatalf("VariantFromUnion() = %v, want %v", got, test.want)
			}
		})
	}
}

func TestFieldUtilities(t *testing.T) {
	value := param.NewOpt("tweet")
	if got := AddrIfPresent(value); got == nil || *got != "tweet" {
		t.Fatalf("AddrIfPresent() = %v", got)
	}
	if got := AddrIfPresent(param.Opt[string]{}); got != nil {
		t.Fatalf("AddrIfPresent(omitted) = %v", got)
	}

	valid := ToOpt("timeline", respjson.NewField(`"timeline"`))
	if !valid.Valid() || valid.Value != "timeline" {
		t.Fatalf("ToOpt(valid) = %+v", valid)
	}
	null := ToOpt("timeline", respjson.NewField(respjson.Null))
	if null.Valid() {
		t.Fatalf("ToOpt(null) = %+v", null)
	}
	omitted := ToOpt("timeline", respjson.Field{})
	if omitted.Valid() {
		t.Fatalf("ToOpt(omitted) = %+v", omitted)
	}

	if Valid(testObject{}) {
		t.Fatal("zero object was valid")
	}
	if Valid(param.NullStruct[testObject]()) {
		t.Fatal("null object was valid")
	}
	if !Valid(param.Override[testObject](map[string]any{"query": "tweet"})) {
		t.Fatal("overridden object was invalid")
	}
}
