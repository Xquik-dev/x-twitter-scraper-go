// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

package param_test

import (
	"bytes"
	"encoding/json"
	"reflect"
	"testing"
	"time"

	shimjson "github.com/Xquik-dev/x-twitter-scraper-go/internal/encoding/json"
	"github.com/Xquik-dev/x-twitter-scraper-go/packages/param"
)

type Struct struct {
	A string `json:"a"`
	B int64  `json:"b"`
	param.APIObject
}

func (r Struct) MarshalJSON() (data []byte, err error) {
	type shadow Struct
	return param.MarshalObject(r, (*shadow)(&r))
}

// Note that the order of fields affects the JSON
// key order. Changing the order of the fields in this struct
// will fail tests unnecessarily.
type FieldStruct struct {
	A param.Opt[string]    `json:"a,omitzero"`
	B param.Opt[int64]     `json:"b,omitzero"`
	C Struct               `json:"c,omitzero"`
	D time.Time            `json:"d,omitzero" format:"date"`
	E time.Time            `json:"e,omitzero"`
	F param.Opt[time.Time] `json:"f,omitzero" format:"date"`
	G param.Opt[time.Time] `json:"g,omitzero"`
	H param.Opt[time.Time] `json:"h,omitzero" format:"date-time"`
	param.APIObject
}

func (r FieldStruct) MarshalJSON() (data []byte, err error) {
	type shadow FieldStruct
	return param.MarshalObject(r, (*shadow)(&r))
}

type StructWithAdditionalProperties struct {
	First       string         `json:"first"`
	Second      int            `json:"second"`
	ExtraFields map[string]any `json:"-"`
	param.APIObject
}

func (s StructWithAdditionalProperties) MarshalJSON() ([]byte, error) {
	type shadow StructWithAdditionalProperties
	return param.MarshalWithExtras(s, (*shadow)(&s), s.ExtraFields)
}

func assertJSON(t *testing.T, value any, expected string) {
	t.Helper()
	data, err := json.Marshal(value)
	if err != nil {
		t.Fatalf("marshal %T: %v; expected %s", value, err, expected)
	}
	if string(data) != expected {
		t.Fatalf("marshal %T: expected %s, received %s", value, expected, data)
	}
}

func TestIsNullish(t *testing.T) {
	for name, test := range map[string]struct {
		value param.ParamNullable
		null  bool
	}{
		"null_string": {param.Null[string](), true},
		"null_int64":  {param.Null[int64](), true},
		"null_time":   {param.Null[time.Time](), true},
		"null_struct": {param.NullStruct[Struct](), true},
		"omit_string": {param.Opt[string]{}, false},
		"omit_int64":  {param.Opt[int64]{}, false},
		"omit_time":   {param.Opt[time.Time]{}, false},
		"omit_struct": {Struct{}, false},
	} {
		t.Run(name, func(t *testing.T) {
			if got := param.IsNull(test.value); got != test.null {
				t.Fatalf("%s: IsNull = %t, want %t", name, got, test.null)
			}
			if got := param.IsOmitted(test.value); got == test.null {
				t.Fatalf("%s: IsOmitted = %t, want %t", name, got, !test.null)
			}
		})
	}
}

func TestFieldMarshal(t *testing.T) {
	tests := map[string]struct {
		value    any
		expected string
	}{
		"null_string": {param.Null[string](), "null"},
		"null_int64":  {param.Null[int64](), "null"},
		"null_time":   {param.Null[time.Time](), "null"},
		"null_struct": {param.NullStruct[Struct](), "null"},

		"float_zero":  {param.NewOpt(float64(0.0)), "0"},
		"string_zero": {param.NewOpt(""), `""`},
		"time_zero":   {param.NewOpt(time.Time{}), `"0001-01-01T00:00:00Z"`},

		"string": {param.Opt[string]{Value: "string"}, `"string"`},
		"int":    {param.Opt[int64]{Value: 123}, "123"},
		"int64":  {param.Opt[int64]{Value: int64(123456789123456789)}, "123456789123456789"},
		"struct": {Struct{A: "yo", B: 123}, `{"a":"yo","b":123}`},
		"datetime": {
			param.Opt[time.Time]{Value: time.Date(2023, time.March, 18, 14, 47, 38, 0, time.UTC)},
			`"2023-03-18T14:47:38Z"`,
		},
		"optional_date": {
			FieldStruct{
				F: param.Opt[time.Time]{Value: time.Date(2023, time.March, 18, 14, 47, 38, 0, time.UTC)},
			},
			`{"f":"2023-03-18"}`,
		},
		"optional_time": {
			FieldStruct{
				G: param.Opt[time.Time]{Value: time.Date(2023, time.March, 18, 14, 47, 38, 0, time.UTC)},
			},
			`{"g":"2023-03-18T14:47:38Z"}`,
		},
		"optional_datetime_explicit_format": {
			FieldStruct{
				H: param.Opt[time.Time]{Value: time.Date(2023, time.March, 18, 14, 47, 38, 0, time.UTC)},
			},
			`{"h":"2023-03-18T14:47:38Z"}`,
		},
		"param_struct": {
			FieldStruct{
				A: param.Opt[string]{Value: "hello"},
				B: param.Opt[int64]{Value: int64(12)},
				D: time.Date(2023, time.March, 18, 14, 47, 38, 0, time.UTC),
				E: time.Date(2023, time.March, 18, 14, 47, 38, 0, time.UTC),
				F: param.Opt[time.Time]{Value: time.Date(2023, time.March, 18, 14, 47, 38, 0, time.UTC)},
				G: param.Opt[time.Time]{Value: time.Date(2023, time.March, 18, 14, 47, 38, 0, time.UTC)},
				H: param.Opt[time.Time]{Value: time.Date(2023, time.March, 18, 14, 47, 38, 0, time.UTC)},
			},
			`{"a":"hello","b":12,"d":"2023-03-18","e":"2023-03-18T14:47:38Z","f":"2023-03-18","g":"2023-03-18T14:47:38Z","h":"2023-03-18T14:47:38Z"}`,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			assertJSON(t, test.value, test.expected)
		})
	}
}

func TestAdditionalProperties(t *testing.T) {
	s := StructWithAdditionalProperties{
		First:  "hello",
		Second: 14,
		ExtraFields: map[string]any{
			"hi": "there",
		},
	}
	exp := `{"first":"hello","second":14,"hi":"there"}`

	assertJSON(t, s, exp)
}

func TestExtraFields(t *testing.T) {
	v := Struct{
		A: "hello",
		B: 123,
	}
	v.SetExtraFields(map[string]any{
		"extra": Struct{A: "recursive"},
		"b":     nil,
	})
	assertJSON(t, v, `{"a":"hello","b":null,"extra":{"a":"recursive","b":0}}`)
	if v.B != 123 {
		t.Fatalf("marshal modified field B: got %v", v.B)
	}
}

func TestExtraFieldsForceOmitted(t *testing.T) {
	v := Struct{
		// Testing with the zero value.
		// A: "",
		// B: 0,
	}
	v.SetExtraFields(map[string]any{
		"b": param.Omit,
	})
	assertJSON(t, v, `{"a":""}`)
}

type UnionWithDates struct {
	OfDate param.Opt[time.Time]
	OfTime param.Opt[time.Time]
	param.APIUnion
}

func (r UnionWithDates) MarshalJSON() (data []byte, err error) {
	return param.MarshalUnion(r, param.EncodedAsDate(r.OfDate), r.OfTime)
}

func TestUnionDateMarshal(t *testing.T) {
	tests := map[string]struct {
		value    UnionWithDates
		expected string
	}{
		"date_only": {
			UnionWithDates{
				OfDate: param.Opt[time.Time]{Value: time.Date(2023, time.March, 18, 0, 0, 0, 0, time.UTC)},
			},
			`"2023-03-18"`,
		},
		"datetime_only": {
			UnionWithDates{
				OfTime: param.Opt[time.Time]{Value: time.Date(2023, time.March, 18, 14, 47, 38, 0, time.UTC)},
			},
			`"2023-03-18T14:47:38Z"`,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			assertJSON(t, test.value, test.expected)
		})
	}
}

func TestOverride(t *testing.T) {
	tests := map[string]struct {
		value    param.ParamStruct
		expected string
	}{
		"param_struct": {
			param.Override[FieldStruct](map[string]any{
				"a": "hello",
				"b": 12,
				"c": nil,
			}),
			`{"a":"hello","b":12,"c":null}`,
		},
		"param_struct_primitive": {
			param.Override[FieldStruct](12),
			`12`,
		},
		"param_struct_null": {
			param.Override[FieldStruct](nil),
			`null`,
		},
	}

	f := FieldStruct{}

	f.SetExtraFields(map[string]any{
		"z": "ok",
	})

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			assertJSON(t, test.value, test.expected)
			if _, ok := test.value.Overrides(); !ok {
				t.Fatalf("expected to be overridden")
			}
		})
	}
}

// Despite implementing the interface, this struct is not an param.Optional
// since it was defined in a different package.
type almostOpt struct{}

func (almostOpt) Valid() bool  { return true }
func (almostOpt) Null() bool   { return false }
func (almostOpt) isZero() bool { return false }

func (almostOpt) implOpt() {}

func TestOptionalInterfaceAssignability(t *testing.T) {
	optInt := param.Opt[int]{}
	if _, ok := any(optInt).(param.Optional); !ok {
		t.Fatalf("failed to assign")
	}

	notOpt := almostOpt{}
	if _, ok := any(notOpt).(param.Optional); ok {
		t.Fatalf("unexpected successful assignment")
	}

	notOpt.implOpt() // silence the warning
}

type PrimitiveUnion struct {
	OfString param.Opt[string]
	OfInt    param.Opt[int]
	param.APIUnion
}

func (p PrimitiveUnion) MarshalJSON() (data []byte, err error) {
	return param.MarshalUnion(p, p.OfString, p.OfInt)
}

func TestOverriddenUnion(t *testing.T) {
	tests := map[string]struct {
		value    PrimitiveUnion
		expected string
	}{
		"string": {
			param.Override[PrimitiveUnion](json.RawMessage(`"hello"`)),
			`"hello"`,
		},
		"int": {
			param.Override[PrimitiveUnion](json.RawMessage(`42`)),
			`42`,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			assertJSON(t, test.value, test.expected)
		})
	}
}

func TestNullStructUnion(t *testing.T) {
	nullUnion := param.NullStruct[PrimitiveUnion]()

	assertJSON(t, nullUnion, "null")
}

//
// Compaction optimization
//

type NonCompactedDoubleParent struct {
	Prop   string             `json:"prop"`
	Parent NonCompactedParent `json:"parent"`

	param.APIObject
}

type NonCompactedParent struct {
	BadChild NonCompacted `json:"bad_child"`

	param.APIObject
}

type NonCompacted struct {
	Raw string

	param.APIObject
}

func (a NonCompactedDoubleParent) MarshalJSON() ([]byte, error) {
	type shadow NonCompactedDoubleParent
	return param.MarshalObject(a, (*shadow)(&a))
}

func (a NonCompactedParent) MarshalJSON() ([]byte, error) {
	type shadow NonCompactedParent
	return param.MarshalObject(a, (*shadow)(&a))
}

func (a NonCompacted) MarshalJSON() ([]byte, error) {
	if a.Raw == "" {
		a.Raw = nonCompactedRaw
	}
	return []byte(a.Raw), nil
}

var nonCompactedRaw string = ` { "foo": "bar" } `

func TestAppendCompactBroken(t *testing.T) {
	tests := map[string]struct {
		value json.Marshaler
	}{
		"red/illegal-json": {
			NonCompacted{Raw: `{ "broken": "json" `},
		},
		"red/nested-with-illegal-json": {
			NonCompactedParent{BadChild: NonCompacted{
				Raw: `{ "broken": "json" `,
			}},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			v, err := json.Marshal(test.value)
			if err == nil {
				t.Fatal("expected error got", v)
			}
		})
	}
}

// TestAppendCompact validates an optimization for internal SDK types to
// avoid O(keys^2) iteration over each JSON object.
//
// It's possible to intentionally trigger this behavior as both a user and
// SDK developer. However, the edge case is quite pathological and requires
// calling [json.Marshaler.MarshalJSON] rather than [json.Marshal].
func TestAppendCompact(t *testing.T) {

	tests := map[string]struct {
		value    json.Marshaler
		expected string
	}{
		//
		// Non-compacted cases
		//
		// Note this is how to exploit the compacter to fail, you must call [json.Marshaler.MarshalJSON] rather than [json.Marshal].
		// The type must also embed [param.APIObject] and return non-compacted JSON.
		//

		"no-compact/fails-compaction": {
			NonCompacted{Raw: nonCompactedRaw},
			nonCompactedRaw,
		},
		"no-compact/nested-with-bad-child": {
			NonCompactedParent{BadChild: NonCompacted{
				Raw: nonCompactedRaw,
			}},
			`{"bad_child":` + nonCompactedRaw + `}`,
		},
		"no-compact/double-nested-with-bad-child": {
			NonCompactedDoubleParent{Prop: "1", Parent: NonCompactedParent{BadChild: NonCompacted{
				Raw: nonCompactedRaw,
			}}},
			`{"prop":"1","parent":{"bad_child":` + nonCompactedRaw + `}}`,
		},

		//
		// Compacted cases
		//

		"override/spaces-within": {
			param.Override[NonCompactedDoubleParent](json.RawMessage(`{"com": "pact"}`)),
			`{"com":"pact"}`,
		},
		"override/spaces-after": {
			param.Override[NonCompactedDoubleParent](json.RawMessage(`{"com":"pact"}  `)),
			`{"com":"pact"}`,
		},
		"override/spaces-before": {
			param.Override[NonCompactedDoubleParent](json.RawMessage(`  {"com":"pact"}`)),
			`{"com":"pact"}`,
		},
		"override/spaces-around": {
			param.Override[NonCompactedDoubleParent](json.RawMessage(` { "com": "pact"   }`)),
			`{"com":"pact"}`,
		},
		"override/override-with-nested": {
			param.Override[NonCompactedDoubleParent](NonCompactedParent{}),
			`{"bad_child":{"foo":"bar"}}`,
		},
		"override/override-with-non-compacted": {
			param.Override[NonCompactedDoubleParent](NonCompacted{}),
			`{"foo":"bar"}`,
		},
	}

	for name, test := range tests {
		var compacted bytes.Buffer
		if err := json.Compact(&compacted, []byte(test.expected)); err != nil {
			t.Fatalf("invalid expected JSON for %s: %v", name, err)
		}
		for method, check := range map[string]struct {
			marshal  func() ([]byte, error)
			expected string
		}{
			"marshal-json":     {test.value.MarshalJSON, test.expected},
			"json-marshal":     {func() ([]byte, error) { return json.Marshal(test.value) }, compacted.String()},
			"shimjson-marshal": {func() ([]byte, error) { return shimjson.Marshal(test.value) }, compacted.String()},
		} {
			t.Run(name+"/"+method, func(t *testing.T) {
				b, err := check.marshal()
				if err != nil {
					t.Fatalf("didn't expect error %v, expected %s", err, check.expected)
				}
				if string(b) != check.expected {
					t.Fatalf("expected %s (%s), received %s", check.expected, reflect.TypeOf(test.value), b)
				}
			})
		}
	}
}
