// Copyright the Xquik contributors.
// SPDX-License-Identifier: Apache-2.0

package apiquery

import (
	"errors"
	"reflect"
	"strconv"
	"strings"
	"testing"

	"github.com/Xquik-dev/x-twitter-scraper-go/packages/param"
)

type JSONQueryValue struct {
	value string
	err   error
}

type QueryCoverageEmbedded struct {
	Embedded string `query:"embedded"`
}

type queryCoverage struct {
	QueryCoverageEmbedded
	hidden string `query:"hidden"`
	NoTag  string
	Skip   string            `query:"-"`
	Empty  string            `query:""`
	Omit   string            `query:"omit,omitzero"`
	Rich   param.Opt[string] `query:"rich"`
	Array  [2]string         `query:"array"`
}

type queryCoverageTags struct {
	All    string `query:"all,omitzero,omitempty,inline"`
	NoTag  string
	Format string `query:"format" format:"date-time"`
}

func (value JSONQueryValue) MarshalJSON() ([]byte, error) {
	if value.err != nil {
		return nil, value.err
	}
	return []byte(`"` + value.value + `"`), nil
}

func TestMarshalEntryPointAndJSONFallback(t *testing.T) {
	values, err := Marshal(struct {
		Query JSONQueryValue `query:"query"`
	}{
		Query: JSONQueryValue{value: "tweet search"},
	})
	if err != nil {
		t.Fatal(err)
	}
	if got := values.Get("query"); got != `"tweet search"` {
		t.Fatalf("query = %q", got)
	}

	wantErr := errors.New("marshal failed")
	_, err = Marshal(struct {
		Query JSONQueryValue `query:"query"`
	}{
		Query: JSONQueryValue{err: wantErr},
	})
	if err == nil || !strings.Contains(err.Error(), wantErr.Error()) {
		t.Fatalf("marshal error = %v", err)
	}

	values, err = Marshal(nil)
	if err != nil || values != nil {
		t.Fatalf("Marshal(nil) = %v, %v", values, err)
	}
}

func TestMarshalAdditionalKindsAndInvalidSettings(t *testing.T) {
	values, err := Marshal(struct {
		Value complex128 `query:"value"`
	}{Value: complex(1, 2)})
	if err != nil {
		t.Fatal(err)
	}
	if got := values.Get("value"); got != "(1+2i)" {
		t.Fatalf("complex value = %q", got)
	}

	key := struct {
		First  string `query:"first"`
		Second string `query:"second"`
	}{First: "a", Second: "b"}
	if _, err := Marshal(map[struct {
		First  string `query:"first"`
		Second string `query:"second"`
	}]string{key: "value"}); err == nil {
		t.Fatal("compound map key was accepted")
	}

	for _, format := range []ArrayQueryFormat{ArrayQueryFormatIndices, ArrayQueryFormat(99)} {
		t.Run(strconv.Itoa(int(format)), func(t *testing.T) {
			defer func() {
				if recover() == nil {
					t.Fatal("invalid array format did not panic")
				}
			}()
			_, _ = MarshalWithSettings(
				[]string{"tweet"},
				QuerySettings{ArrayFormat: format},
			)
		})
	}
}

func TestEncoderBranchContracts(t *testing.T) {
	values, err := MarshalWithSettings(queryCoverage{
		QueryCoverageEmbedded: QueryCoverageEmbedded{Embedded: "tweet"},
		Rich:                  param.Null[string](),
		Array:                 [2]string{"search", "timeline"},
	}, QuerySettings{ArrayFormat: ArrayQueryFormatRepeat})
	if err != nil {
		t.Fatal(err)
	}
	if values.Get("rich") != "null" || len(values["array"]) != 2 {
		t.Fatalf("values = %v", values)
	}

	wantErr := errors.New("element failed")
	for _, format := range []ArrayQueryFormat{
		ArrayQueryFormatComma,
		ArrayQueryFormatRepeat,
		ArrayQueryFormatBrackets,
	} {
		t.Run(strconv.Itoa(int(format)), func(t *testing.T) {
			_, err := MarshalWithSettings(
				[]JSONQueryValue{{err: wantErr}},
				QuerySettings{ArrayFormat: format},
			)
			if err == nil || !strings.Contains(err.Error(), wantErr.Error()) {
				t.Fatalf("array error = %v", err)
			}
		})
	}

	primitive := (&encoder{}).newPrimitiveTypeEncoder(reflect.TypeOf((*string)(nil)))
	if pairs, err := primitive("value", reflect.ValueOf((*string)(nil))); err != nil || pairs != nil {
		t.Fatalf("nil pointer = %v, %v", pairs, err)
	}
	unsupported := (&encoder{}).newPrimitiveTypeEncoder(reflect.TypeOf((chan int)(nil)))
	if pairs, err := unsupported("value", reflect.ValueOf((chan int)(nil))); err != nil || pairs != nil {
		t.Fatalf("unsupported primitive = %v, %v", pairs, err)
	}

	var value any
	encodeInterface := (encoder{}).newInterfaceEncoder()
	if pairs, err := encodeInterface("value", reflect.ValueOf(&value).Elem()); err != nil || pairs != nil {
		t.Fatalf("nil interface = %v, %v", pairs, err)
	}
}

func TestQueryTagBranches(t *testing.T) {
	typ := reflect.TypeOf(queryCoverageTags{})
	all, ok := parseQueryStructTag(typ.Field(0))
	if !ok || all.name != "all" || !all.omitzero || !all.omitempty || !all.inline {
		t.Fatalf("all tag = %+v, %t", all, ok)
	}
	if _, ok := parseQueryStructTag(typ.Field(1)); ok {
		t.Fatal("untagged field was accepted")
	}
	if format, ok := parseFormatStructTag(typ.Field(2)); !ok || format != "date-time" {
		t.Fatalf("format = %q, %t", format, ok)
	}
}
