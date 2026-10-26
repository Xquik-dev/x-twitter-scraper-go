// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

package apiquery

import (
	"fmt"
	"github.com/Xquik-dev/x-twitter-scraper-go/packages/param"
	"net/url"
	"testing"
	"time"
)

func P[T any](v T) *T { return &v }

type Primitives struct {
	A bool    `query:"a"`
	B int     `query:"b"`
	C uint    `query:"c"`
	D float64 `query:"d"`
	E float32 `query:"e"`
	F []int   `query:"f"`
}

type PrimitivePointers struct {
	A *bool    `query:"a"`
	B *int     `query:"b"`
	C *uint    `query:"c"`
	D *float64 `query:"d"`
	E *float32 `query:"e"`
	F *[]int   `query:"f"`
}

type Slices struct {
	Slice []Primitives `query:"slices"`
	Mixed []queryValue `query:"mixed"`
}

type DateTime struct {
	Date     time.Time `query:"date" format:"date"`
	DateTime time.Time `query:"date-time" format:"date-time"`
}

type AdditionalProperties struct {
	A      bool                  `query:"a"`
	Extras map[string]queryValue `query:"-,inline"`
}

type Recursive struct {
	Name  string     `query:"name"`
	Child *Recursive `query:"child"`
}

type UnknownStruct struct {
	Unknown queryValue `query:"unknown"`
}

type UnionStruct struct {
	Union Union `query:"union" format:"date"`
}

type Union interface {
	union()
}

type UnionInteger int64

func (UnionInteger) union() {}

type UnionString string

func (UnionString) union() {}

type UnionStructA struct {
	Type string `query:"type"`
	A    string `query:"a"`
	B    string `query:"b"`
}

func (UnionStructA) union() {}

type UnionStructB struct {
	Type string `query:"type"`
	A    string `query:"a"`
}

func (UnionStructB) union() {}

type UnionTime time.Time

func (UnionTime) union() {}

type DeeplyNested struct {
	A DeeplyNested1 `query:"a"`
}

type DeeplyNested1 struct {
	B DeeplyNested2 `query:"b"`
}

type DeeplyNested2 struct {
	C DeeplyNested3 `query:"c"`
}

type DeeplyNested3 struct {
	D *string `query:"d"`
}

func nestedQuery(value *string) DeeplyNested {
	return DeeplyNested{
		A: DeeplyNested1{
			B: DeeplyNested2{
				C: DeeplyNested3{D: value},
			},
		},
	}
}

type RichPrimitives struct {
	A param.Opt[string] `query:"a"`
}

type QueryOmitTest struct {
	A param.Opt[string] `query:"a,omitzero"`
	B string            `query:"b,omitzero"`
}

type NamedEnum string

const NamedEnumFoo NamedEnum = "foo"

type StructUnionWrapper struct {
	Union StructUnion `query:"union"`
}

type StructUnion struct {
	OfInt    param.Opt[int64]     `query:",omitzero,inline"`
	OfString param.Opt[string]    `query:",omitzero,inline"`
	OfEnum   param.Opt[NamedEnum] `query:",omitzero,inline"`
	OfA      UnionStructA         `query:",omitzero,inline"`
	OfB      UnionStructB         `query:",omitzero,inline"`
	param.APIUnion
}

type queryValue interface{ queryValue() }

type (
	queryBool   bool
	queryInt    int
	queryInt8   int8
	queryUint8  uint8
	queryFloat  float64
	queryString string
	queryMap    map[string]queryValue
)

func (queryBool) queryValue()   {}
func (queryInt) queryValue()    {}
func (queryInt8) queryValue()   {}
func (queryUint8) queryValue()  {}
func (queryFloat) queryValue()  {}
func (queryString) queryValue() {}
func (queryMap) queryValue()    {}

type queryTest struct {
	enc     string
	val     string
	marshal func() (url.Values, error)
}

func queryCase[T any](enc string, val T, settings QuerySettings) queryTest {
	return queryTest{enc, fmt.Sprintf("%+#v", val), func() (url.Values, error) {
		return MarshalWithSettings(val, settings)
	}}
}

var tests = map[string]queryTest{
	"primitives": queryCase("a=false&b=237628372683&c=654&d=9999.43&e=43.7599983215332&f=1,2,3,4", Primitives{A: false, B: 237628372683, C: uint(654), D: 9999.43, E: 43.76, F: []int{1, 2, 3, 4}}, QuerySettings{}),

	"slices_brackets": queryCase(`mixed[]=1&mixed[]=2.3&mixed[]=hello&slices[][a]=false&slices[][a]=false&slices[][b]=237628372683&slices[][b]=237628372683&slices[][c]=654&slices[][c]=654&slices[][d]=9999.43&slices[][d]=9999.43&slices[][e]=43.7599983215332&slices[][e]=43.7599983215332&slices[][f][]=1&slices[][f][]=2&slices[][f][]=3&slices[][f][]=4&slices[][f][]=1&slices[][f][]=2&slices[][f][]=3&slices[][f][]=4`, Slices{
		Slice: []Primitives{
			{A: false, B: 237628372683, C: uint(654), D: 9999.43, E: 43.76, F: []int{1, 2, 3, 4}},
			{A: false, B: 237628372683, C: uint(654), D: 9999.43, E: 43.76, F: []int{1, 2, 3, 4}},
		},
		Mixed: []queryValue{queryInt(1), queryFloat(2.3), queryString("hello")},
	}, QuerySettings{ArrayFormat: ArrayQueryFormatBrackets}),

	"slices_comma": queryCase(`mixed=1,2.3,hello`, Slices{
		Mixed: []queryValue{queryInt(1), queryFloat(2.3), queryString("hello")},
	}, QuerySettings{ArrayFormat: ArrayQueryFormatComma}),

	"slices_repeat": queryCase(`mixed=1&mixed=2.3&mixed=hello&slices[a]=false&slices[a]=false&slices[b]=237628372683&slices[b]=237628372683&slices[c]=654&slices[c]=654&slices[d]=9999.43&slices[d]=9999.43&slices[e]=43.7599983215332&slices[e]=43.7599983215332&slices[f]=1&slices[f]=2&slices[f]=3&slices[f]=4&slices[f]=1&slices[f]=2&slices[f]=3&slices[f]=4`, Slices{
		Slice: []Primitives{
			{A: false, B: 237628372683, C: uint(654), D: 9999.43, E: 43.76, F: []int{1, 2, 3, 4}},
			{A: false, B: 237628372683, C: uint(654), D: 9999.43, E: 43.76, F: []int{1, 2, 3, 4}},
		},
		Mixed: []queryValue{queryInt(1), queryFloat(2.3), queryString("hello")},
	}, QuerySettings{ArrayFormat: ArrayQueryFormatRepeat}),

	"primitive_pointer_struct": queryCase("a=false&b=237628372683&c=654&d=9999.43&e=43.7599983215332&f=1,2,3,4,5", PrimitivePointers{
		A: P(false),
		B: P(237628372683),
		C: P(uint(654)),
		D: P(9999.43),
		E: P(float32(43.76)),
		F: &[]int{1, 2, 3, 4, 5},
	}, QuerySettings{}),

	"datetime_struct": queryCase(`date=2006-01-02&date-time=2006-01-02T15:04:05Z`, DateTime{
		Date:     time.Date(2006, time.January, 2, 0, 0, 0, 0, time.UTC),
		DateTime: time.Date(2006, time.January, 2, 15, 4, 5, 0, time.UTC),
	}, QuerySettings{}),

	"additional_properties": queryCase(`a=true&bar=value&foo=true`, AdditionalProperties{
		A: true,
		Extras: map[string]queryValue{
			"bar": queryString("value"),
			"foo": queryBool(true),
		},
	}, QuerySettings{}),

	"recursive_struct_brackets": queryCase(`child[name]=Alex&name=Robert`, Recursive{Name: "Robert", Child: &Recursive{Name: "Alex"}}, QuerySettings{NestedFormat: NestedQueryFormatBrackets}),

	"recursive_struct_dots": queryCase(`child.name=Alex&name=Robert`, Recursive{Name: "Robert", Child: &Recursive{Name: "Alex"}}, QuerySettings{NestedFormat: NestedQueryFormatDots}),

	"unknown_struct_number": queryCase(`unknown=12`, UnknownStruct{
		Unknown: queryFloat(12.),
	}, QuerySettings{}),

	"unknown_struct_map_brackets": queryCase(`unknown[foo]=bar`, UnknownStruct{
		Unknown: queryMap{
			"foo": queryString("bar"),
		},
	}, QuerySettings{NestedFormat: NestedQueryFormatBrackets}),

	"unknown_struct_map_dots": queryCase(`unknown.foo=bar`, UnknownStruct{
		Unknown: queryMap{
			"foo": queryString("bar"),
		},
	}, QuerySettings{NestedFormat: NestedQueryFormatDots}),

	"struct_union_string": queryCase(`union=hello`, StructUnionWrapper{
		Union: StructUnion{OfString: param.NewOpt("hello")},
	}, QuerySettings{}),

	"union_string": queryCase(`union=hello`, UnionStruct{
		Union: UnionString("hello"),
	}, QuerySettings{}),

	"struct_union_integer": queryCase(`union=12`, StructUnionWrapper{
		Union: StructUnion{OfInt: param.NewOpt[int64](12)},
	}, QuerySettings{}),

	"union_integer": queryCase(`union=12`, UnionStruct{
		Union: UnionInteger(12),
	}, QuerySettings{}),

	"struct_union_enum": queryCase(`union=foo`, StructUnionWrapper{
		Union: StructUnion{OfEnum: param.NewOpt[NamedEnum](NamedEnumFoo)},
	}, QuerySettings{}),

	"struct_union_struct_discriminated_a": queryCase(`union[a]=foo&union[b]=bar&union[type]=typeA`, StructUnionWrapper{
		Union: StructUnion{OfA: UnionStructA{
			Type: "typeA",
			A:    "foo",
			B:    "bar",
		}},
	}, QuerySettings{}),

	"union_struct_discriminated_a": queryCase(`union[a]=foo&union[b]=bar&union[type]=typeA`, UnionStruct{
		Union: UnionStructA{
			Type: "typeA",
			A:    "foo",
			B:    "bar",
		},
	}, QuerySettings{}),

	"struct_union_struct_discriminated_b": queryCase(`union[a]=foo&union[type]=typeB`, StructUnionWrapper{
		Union: StructUnion{OfB: UnionStructB{
			Type: "typeB",
			A:    "foo",
		}},
	}, QuerySettings{}),

	"union_struct_discriminated_b": queryCase(`union[a]=foo&union[type]=typeB`, UnionStruct{
		Union: UnionStructB{
			Type: "typeB",
			A:    "foo",
		},
	}, QuerySettings{}),

	"union_struct_time": queryCase(`union=2010-05-23`, UnionStruct{
		Union: UnionTime(time.Date(2010, 05, 23, 0, 0, 0, 0, time.UTC)),
	}, QuerySettings{}),

	"deeply_nested_brackets": queryCase(`a[b][c][d]=hello`, nestedQuery(P("hello")), QuerySettings{NestedFormat: NestedQueryFormatBrackets}),

	"deeply_nested_dots": queryCase(`a.b.c.d=hello`, nestedQuery(P("hello")), QuerySettings{NestedFormat: NestedQueryFormatDots}),

	"deeply_nested_brackets_empty": queryCase(``, nestedQuery(nil), QuerySettings{NestedFormat: NestedQueryFormatBrackets}),

	"deeply_nested_dots_empty": queryCase(``, nestedQuery(nil), QuerySettings{NestedFormat: NestedQueryFormatDots}),

	"rich_primitives": queryCase(`a=hello`, RichPrimitives{
		A: param.Opt[string]{Value: "hello"},
	}, QuerySettings{}),

	"rich_primitives_omit": queryCase(``, QueryOmitTest{
		A: param.Opt[string]{},
	}, QuerySettings{}),
	"query_omit": queryCase(`a=hello`, QueryOmitTest{
		A: param.Opt[string]{Value: "hello"},
	}, QuerySettings{}),
}

func TestEncode(t *testing.T) {
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			values, err := test.marshal()
			if err != nil {
				t.Fatalf("failed to marshal url %s", err)
			}
			str, err := url.QueryUnescape(values.Encode())
			if err != nil {
				t.Fatal(err)
			}
			if str != test.enc {
				t.Fatalf("expected %s to serialize to %s but got %s", test.val, test.enc, str)
			}
		})
	}
}
