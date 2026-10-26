// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

package apijson

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/tidwall/gjson"
)

func P[T any](v T) *T { return &v }

type Primitives struct {
	A bool    `json:"a"`
	B int     `json:"b"`
	C uint    `json:"c"`
	D float64 `json:"d"`
	E float32 `json:"e"`
	F []int   `json:"f"`
}

type PrimitivePointers struct {
	A *bool    `json:"a"`
	B *int     `json:"b"`
	C *uint    `json:"c"`
	D *float64 `json:"d"`
	E *float32 `json:"e"`
	F *[]int   `json:"f"`
}

type Slices struct {
	Slice []Primitives `json:"slices"`
}

type DateTime struct {
	Date     time.Time `json:"date" format:"date"`
	DateTime time.Time `json:"date-time" format:"date-time"`
}

type AdditionalProperties struct {
	A           bool                       `json:"a"`
	ExtraFields map[string]json.RawMessage `json:"-" api:"extrafields"`
}

type TypedAdditionalProperties struct {
	A           bool           `json:"a"`
	ExtraFields map[string]int `json:"-" api:"extrafields"`
}

type EmbeddedStruct struct {
	A bool   `json:"a"`
	B string `json:"b"`

	JSON EmbeddedStructJSON
}

type EmbeddedStructJSON struct {
	A           Field
	B           Field
	ExtraFields map[string]Field
	raw         string
}

type EmbeddedStructs struct {
	EmbeddedStruct
	A           *int                       `json:"a"`
	ExtraFields map[string]json.RawMessage `json:"-" api:"extrafields"`

	JSON EmbeddedStructsJSON
}

type EmbeddedStructsJSON struct {
	A           Field
	ExtraFields map[string]Field
	raw         string
}

type Recursive struct {
	Name  string     `json:"name"`
	Child *Recursive `json:"child"`
}

type JSONFieldStruct struct {
	A           bool                `json:"a"`
	B           int64               `json:"b"`
	C           string              `json:"c"`
	D           string              `json:"d"`
	ExtraFields map[string]int64    `json:"" api:"extrafields"`
	JSON        JSONFieldStructJSON `json:"-" api:"metadata"`
}

type JSONFieldStructJSON struct {
	A           Field
	B           Field
	C           Field
	D           Field
	ExtraFields map[string]Field
	raw         string
}

type UnknownStruct struct {
	Unknown json.RawMessage `json:"unknown"`
}

type UnionStruct struct {
	Union Union `json:"union" format:"date"`
}

type Union interface {
	union()
}

type Inline struct {
	InlineField Primitives `json:",inline"`
	JSON        InlineJSON `json:"-" api:"metadata"`
}

type InlineArray struct {
	InlineField []string   `json:",inline"`
	JSON        InlineJSON `json:"-" api:"metadata"`
}

type InlineJSON struct {
	InlineField Field
	raw         string
}

type UnionInteger int64

func (UnionInteger) union() {}

type UnionStructA struct {
	Type string `json:"type"`
	A    string `json:"a"`
	B    string `json:"b"`
}

func (UnionStructA) union() {}

type UnionStructB struct {
	Type string `json:"type"`
	A    string `json:"a"`
}

func (UnionStructB) union() {}

type UnionTime time.Time

func (UnionTime) union() {}

func init() {
	RegisterUnion[Union]("type",
		UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(UnionTime{}),
		},
		UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(UnionInteger(0)),
		},
		UnionVariant{
			TypeFilter:         gjson.JSON,
			DiscriminatorValue: "typeA",
			Type:               reflect.TypeOf(UnionStructA{}),
		},
		UnionVariant{
			TypeFilter:         gjson.JSON,
			DiscriminatorValue: "typeB",
			Type:               reflect.TypeOf(UnionStructB{}),
		},
	)
}

type ComplexUnionStruct struct {
	Union ComplexUnion `json:"union"`
}

type ComplexUnion interface {
	complexUnion()
}

type ComplexUnionA struct {
	Boo string `json:"boo"`
	Foo bool   `json:"foo"`
}

func (ComplexUnionA) complexUnion() {}

type ComplexUnionB struct {
	Boo bool   `json:"boo"`
	Foo string `json:"foo"`
}

func (ComplexUnionB) complexUnion() {}

type ComplexUnionC struct {
	Boo int64 `json:"boo"`
}

func (ComplexUnionC) complexUnion() {}

type ComplexUnionTypeA struct {
	Baz  int64 `json:"baz"`
	Type TypeA `json:"type"`
}

func (ComplexUnionTypeA) complexUnion() {}

type TypeA string

func (t TypeA) IsKnown() bool {
	return t == "a"
}

type ComplexUnionTypeB struct {
	Baz  int64 `json:"baz"`
	Type TypeB `json:"type"`
}

type TypeB string

func (t TypeB) IsKnown() bool {
	return t == "b"
}

type UnmarshalStruct struct {
	Foo  string `json:"foo"`
	prop bool   `json:"-"`
}

func (r *UnmarshalStruct) UnmarshalJSON(json []byte) error {
	r.prop = true
	return UnmarshalRoot(json, r, nil)
}

func (ComplexUnionTypeB) complexUnion() {}

func init() {
	RegisterUnion[ComplexUnion]("",
		UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ComplexUnionA{}),
		},
		UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ComplexUnionB{}),
		},
		UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ComplexUnionC{}),
		},
		UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ComplexUnionTypeA{}),
		},
		UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ComplexUnionTypeB{}),
		},
	)
}

type MarshallingUnionStruct struct {
	Union MarshallingUnion
}

func (r *MarshallingUnionStruct) UnmarshalJSON(data []byte) (err error) {
	*r = MarshallingUnionStruct{}
	err = UnmarshalRoot(data, &r.Union, nil)
	return err
}

func (r MarshallingUnionStruct) MarshalJSON() (data []byte, err error) {
	return MarshalRoot(r.Union)
}

type MarshallingUnion interface {
	marshallingUnion()
}

type MarshallingUnionA struct {
	Boo string `json:"boo"`
}

func (MarshallingUnionA) marshallingUnion() {}

func (r *MarshallingUnionA) UnmarshalJSON(data []byte) (err error) {
	return UnmarshalRoot(data, r, nil)
}

type MarshallingUnionB struct {
	Foo string `json:"foo"`
}

func (MarshallingUnionB) marshallingUnion() {}

func (r *MarshallingUnionB) UnmarshalJSON(data []byte) (err error) {
	return UnmarshalRoot(data, r, nil)
}

func init() {
	RegisterUnion[MarshallingUnion](
		"",
		UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(MarshallingUnionA{}),
		},
		UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(MarshallingUnionB{}),
		},
	)
}

type jsonTest struct {
	buf    string
	val    string
	decode func() error
	encode func() ([]byte, error)
}

func jsonCase[T any](buf string, val T) jsonTest {
	return jsonTest{
		buf: buf,
		val: fmt.Sprintf("%#v", val),
		decode: func() error {
			var result T
			if err := Unmarshal([]byte(buf), &result); err != nil {
				return fmt.Errorf("deserialization failed with error %w", err)
			}
			if !reflect.DeepEqual(result, val) {
				return fmt.Errorf("got\n%#v", result)
			}
			return nil
		},
		encode: func() ([]byte, error) { return Marshal(val) },
	}
}

var tests = map[string]jsonTest{
	"true":               jsonCase("true", true),
	"false":              jsonCase("false", false),
	"int":                jsonCase("1", 1),
	"int_bigger":         jsonCase("12324", 12324),
	"int_string_coerce":  jsonCase(`"65"`, 65),
	"int_boolean_coerce": jsonCase("true", 1),
	"int64":              jsonCase("1", int64(1)),
	"int64_huge":         jsonCase("123456789123456789", int64(123456789123456789)),
	"uint":               jsonCase("1", uint(1)),
	"uint_bigger":        jsonCase("12324", uint(12324)),
	"uint_coerce":        jsonCase(`"65"`, uint(65)),
	"float_1.54":         jsonCase("1.54", float32(1.54)),
	"float_1.89":         jsonCase("1.89", float64(1.89)),
	"string":             jsonCase(`"str"`, "str"),
	"string_int_coerce":  jsonCase(`12`, "12"),
	"array_string":       jsonCase(`["foo","bar"]`, []string{"foo", "bar"}),
	"array_int":          jsonCase(`[1,2]`, []int{1, 2}),
	"array_int_coerce":   jsonCase(`["1",2]`, []int{1, 2}),

	"ptr_true":               jsonCase("true", P(true)),
	"ptr_false":              jsonCase("false", P(false)),
	"ptr_int":                jsonCase("1", P(1)),
	"ptr_int_bigger":         jsonCase("12324", P(12324)),
	"ptr_int_string_coerce":  jsonCase(`"65"`, P(65)),
	"ptr_int_boolean_coerce": jsonCase("true", P(1)),
	"ptr_int64":              jsonCase("1", P(int64(1))),
	"ptr_int64_huge":         jsonCase("123456789123456789", P(int64(123456789123456789))),
	"ptr_uint":               jsonCase("1", P(uint(1))),
	"ptr_uint_bigger":        jsonCase("12324", P(uint(12324))),
	"ptr_uint_coerce":        jsonCase(`"65"`, P(uint(65))),
	"ptr_float_1.54":         jsonCase("1.54", P(float32(1.54))),
	"ptr_float_1.89":         jsonCase("1.89", P(float64(1.89))),

	"date_time":             jsonCase(`"2007-03-01T13:00:00Z"`, time.Date(2007, time.March, 1, 13, 0, 0, 0, time.UTC)),
	"date_time_nano_coerce": jsonCase(`"2007-03-01T13:03:05.123456789Z"`, time.Date(2007, time.March, 1, 13, 3, 5, 123456789, time.UTC)),

	"date_time_missing_t_coerce":              jsonCase(`"2007-03-01 13:03:05Z"`, time.Date(2007, time.March, 1, 13, 3, 5, 0, time.UTC)),
	"date_time_missing_timezone_coerce":       jsonCase(`"2007-03-01T13:03:05"`, time.Date(2007, time.March, 1, 13, 3, 5, 0, time.UTC)),
	"date_time_missing_timezone_colon_coerce": jsonCase(`"2007-03-01T13:03:05-1200"`, time.Date(2007, time.March, 1, 13, 3, 5, 0, time.FixedZone("", -12*60*60))),
	"date_time_nano_missing_t_coerce":         jsonCase(`"2007-03-01 13:03:05.123456789Z"`, time.Date(2007, time.March, 1, 13, 3, 5, 123456789, time.UTC)),

	"map_string":                       jsonCase(`{"foo":"bar"}`, map[string]string{"foo": "bar"}),
	"map_string_with_sjson_path_chars": jsonCase(`{":a.b.c*:d*-1e.f":"bar"}`, map[string]string{":a.b.c*:d*-1e.f": "bar"}),
	"map_interface":                    jsonCase(`{"a":1,"b":"str","c":false}`, map[string]json.RawMessage{"a": json.RawMessage("1"), "b": json.RawMessage(`"str"`), "c": json.RawMessage("false")}),

	"primitive_struct": jsonCase(`{"a":false,"b":237628372683,"c":654,"d":9999.43,"e":43.76,"f":[1,2,3,4]}`, Primitives{A: false, B: 237628372683, C: uint(654), D: 9999.43, E: 43.76, F: []int{1, 2, 3, 4}}),

	"slices": jsonCase(`{"slices":[{"a":false,"b":237628372683,"c":654,"d":9999.43,"e":43.76,"f":[1,2,3,4]}]}`, Slices{
		Slice: []Primitives{{A: false, B: 237628372683, C: uint(654), D: 9999.43, E: 43.76, F: []int{1, 2, 3, 4}}},
	}),

	"primitive_pointer_struct": jsonCase(`{"a":false,"b":237628372683,"c":654,"d":9999.43,"e":43.76,"f":[1,2,3,4,5]}`, PrimitivePointers{
		A: P(false),
		B: P(237628372683),
		C: P(uint(654)),
		D: P(9999.43),
		E: P(float32(43.76)),
		F: &[]int{1, 2, 3, 4, 5},
	}),

	"datetime_struct": jsonCase(`{"date":"2006-01-02","date-time":"2006-01-02T15:04:05Z"}`, DateTime{
		Date:     time.Date(2006, time.January, 2, 0, 0, 0, 0, time.UTC),
		DateTime: time.Date(2006, time.January, 2, 15, 4, 5, 0, time.UTC),
	}),

	"additional_properties": jsonCase(`{"a":true,"bar":"value","foo":true}`, AdditionalProperties{
		A: true,
		ExtraFields: map[string]json.RawMessage{
			"bar": json.RawMessage(`"value"`),
			"foo": json.RawMessage("true"),
		},
	}),

	"embedded_struct": jsonCase(`{"a":1,"b":"bar"}`, EmbeddedStructs{
		EmbeddedStruct: EmbeddedStruct{
			A: true,
			B: "bar",
			JSON: EmbeddedStructJSON{
				A: Field{raw: `1`, status: valid},
				B: Field{raw: `"bar"`, status: valid},
			},
		},
		A:           P(1),
		ExtraFields: map[string]json.RawMessage{"b": json.RawMessage(`"bar"`)},
		JSON: EmbeddedStructsJSON{
			A: Field{raw: `1`, status: valid},
			ExtraFields: map[string]Field{
				"b": {raw: `"bar"`, status: valid},
			},
		},
	}),

	"recursive_struct": jsonCase(`{"child":{"name":"Alex"},"name":"Robert"}`, Recursive{Name: "Robert", Child: &Recursive{Name: "Alex"}}),

	"metadata_coerce": jsonCase(`{"a":"12","b":"12","c":null,"extra_typed":12,"extra_untyped":{"foo":"bar"}}`, JSONFieldStruct{
		A: false,
		B: 12,
		C: "",
		JSON: JSONFieldStructJSON{
			raw: `{"a":"12","b":"12","c":null,"extra_typed":12,"extra_untyped":{"foo":"bar"}}`,
			A:   Field{raw: `"12"`, status: invalid},
			B:   Field{raw: `"12"`, status: valid},
			C:   Field{raw: "null", status: null},
			D:   Field{raw: "", status: missing},
			ExtraFields: map[string]Field{
				"extra_typed": {
					raw:    "12",
					status: valid,
				},
				"extra_untyped": {
					raw:    `{"foo":"bar"}`,
					status: invalid,
				},
			},
		},
		ExtraFields: map[string]int64{
			"extra_typed":   12,
			"extra_untyped": 0,
		},
	}),

	"unknown_struct_number": jsonCase(`{"unknown":12}`, UnknownStruct{
		Unknown: json.RawMessage("12"),
	}),

	"unknown_struct_map": jsonCase(`{"unknown":{"foo":"bar"}}`, UnknownStruct{
		Unknown: json.RawMessage(`{"foo":"bar"}`),
	}),

	"union_integer": jsonCase(`{"union":12}`, UnionStruct{
		Union: UnionInteger(12),
	}),

	"union_struct_discriminated_a": jsonCase(`{"union":{"a":"foo","b":"bar","type":"typeA"}}`, UnionStruct{
		Union: UnionStructA{
			Type: "typeA",
			A:    "foo",
			B:    "bar",
		},
	}),

	"union_struct_discriminated_b": jsonCase(`{"union":{"a":"foo","type":"typeB"}}`, UnionStruct{
		Union: UnionStructB{
			Type: "typeB",
			A:    "foo",
		},
	}),

	"union_struct_time": jsonCase(`{"union":"2010-05-23"}`, UnionStruct{
		Union: UnionTime(time.Date(2010, 05, 23, 0, 0, 0, 0, time.UTC)),
	}),

	"complex_union_a": jsonCase(`{"union":{"boo":"12","foo":true}}`, ComplexUnionStruct{Union: ComplexUnionA{Boo: "12", Foo: true}}),

	"complex_union_b": jsonCase(`{"union":{"boo":true,"foo":"12"}}`, ComplexUnionStruct{Union: ComplexUnionB{Boo: true, Foo: "12"}}),

	"complex_union_c": jsonCase(`{"union":{"boo":12}}`, ComplexUnionStruct{Union: ComplexUnionC{Boo: 12}}),

	"complex_union_type_a": jsonCase(`{"union":{"baz":12,"type":"a"}}`, ComplexUnionStruct{Union: ComplexUnionTypeA{Baz: 12, Type: TypeA("a")}}),

	"complex_union_type_b": jsonCase(`{"union":{"baz":12,"type":"b"}}`, ComplexUnionStruct{Union: ComplexUnionTypeB{Baz: 12, Type: TypeB("b")}}),

	"marshalling_union_a": jsonCase(`{"boo":"hello"}`, MarshallingUnionStruct{Union: MarshallingUnionA{Boo: "hello"}}),
	"marshalling_union_b": jsonCase(`{"foo":"hi"}`, MarshallingUnionStruct{Union: MarshallingUnionB{Foo: "hi"}}),

	"unmarshal": jsonCase(`{"foo":"hello"}`, &UnmarshalStruct{Foo: "hello", prop: true}),

	"array_of_unmarshal": jsonCase(`[{"foo":"hello"}]`, []UnmarshalStruct{{Foo: "hello", prop: true}}),

	"inline_coerce": jsonCase(`{"a":false,"b":237628372683,"c":654,"d":9999.43,"e":43.76,"f":[1,2,3,4]}`, Inline{
		InlineField: Primitives{A: false, B: 237628372683, C: 0x28e, D: 9999.43, E: 43.76, F: []int{1, 2, 3, 4}},
		JSON: InlineJSON{
			InlineField: Field{raw: "{\"a\":false,\"b\":237628372683,\"c\":654,\"d\":9999.43,\"e\":43.76,\"f\":[1,2,3,4]}", status: 3},
			raw:         "{\"a\":false,\"b\":237628372683,\"c\":654,\"d\":9999.43,\"e\":43.76,\"f\":[1,2,3,4]}",
		},
	}),

	"inline_array_coerce": jsonCase(`["Hello","foo","bar"]`, InlineArray{
		InlineField: []string{"Hello", "foo", "bar"},
		JSON: InlineJSON{
			InlineField: Field{raw: `["Hello","foo","bar"]`, status: 3},
			raw:         `["Hello","foo","bar"]`,
		},
	}),
}

func TestDecode(t *testing.T) {
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			if err := test.decode(); err != nil {
				t.Fatalf("expected '%s' to deserialize to \n%s\nbut %v", test.buf, test.val, err)
			}
		})
	}
}

func TestEncode(t *testing.T) {
	for name, test := range tests {
		if strings.HasSuffix(name, "_coerce") {
			continue
		}
		t.Run(name, func(t *testing.T) {
			raw, err := test.encode()
			if err != nil {
				t.Fatalf("serialization of %s failed with error %v", test.val, err)
			}
			if string(raw) != test.buf {
				t.Fatalf("expected %s to serialize to %s but got %s", test.val, test.buf, string(raw))
			}
		})
	}
}

type StructWithDefault struct {
	Type string `json:"type" default:"foo"`
}

func TestDefault(t *testing.T) {
	value := StructWithDefault{}
	expected := `{"type":"foo"}`

	raw, err := Marshal(value)
	if err != nil {
		t.Fatalf("serialization of %v failed with error %v", value, err)
	}
	if string(raw) != expected {
		t.Fatalf("expected %+#v to serialize to %s but got %s", value, expected, string(raw))
	}
}
