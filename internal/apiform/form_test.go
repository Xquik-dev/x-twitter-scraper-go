// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

package apiform

import (
	"bytes"
	"github.com/Xquik-dev/x-twitter-scraper-go/packages/param"
	"io"
	"mime/multipart"
	"strings"
	"testing"
	"time"
)

func P[T any](v T) *T { return &v }

type Primitives struct {
	A bool    `form:"a"`
	B int     `form:"b"`
	C uint    `form:"c"`
	D float64 `form:"d"`
	E float32 `form:"e"`
	F []int   `form:"f"`
}

// These aliases are necessary to bypass the cache.
// This only relevant during testing.
type int_ int
type PrimitivesBrackets struct {
	F []int_ `form:"f"`
}

type PrimitivePointers struct {
	A *bool    `form:"a"`
	B *int     `form:"b"`
	C *uint    `form:"c"`
	D *float64 `form:"d"`
	E *float32 `form:"e"`
	F *[]int   `form:"f"`
}

type Slices struct {
	Slice []Primitives `form:"slices"`
}

type DateTime struct {
	Date     time.Time `form:"date" format:"date"`
	DateTime time.Time `form:"date-time" format:"date-time"`
}

type AdditionalProperties struct {
	A      bool           `form:"a"`
	Extras map[string]any `form:"-" api:"extrafields"`
}

type TypedAdditionalProperties struct {
	A      bool           `form:"a"`
	Extras map[string]int `form:"-" api:"extrafields"`
}

type EmbeddedStructs struct {
	AdditionalProperties
	A      *int           `form:"number2"`
	Extras map[string]any `form:"-" api:"extrafields"`
}

type Recursive struct {
	Name  string     `form:"name"`
	Child *Recursive `form:"child"`
}

type UnknownStruct struct {
	Unknown any `form:"unknown"`
}

type UnionStruct struct {
	Union Union `form:"union" format:"date"`
}

type Union interface {
	union()
}

type UnionInteger int64

func (UnionInteger) union() {}

type UnionStructA struct {
	Type string `form:"type"`
	A    string `form:"a"`
	B    string `form:"b"`
}

func (UnionStructA) union() {}

type UnionStructB struct {
	Type string `form:"type"`
	A    string `form:"a"`
}

func (UnionStructB) union() {}

type UnionTime time.Time

func (UnionTime) union() {}

type ReaderStruct struct {
	File io.Reader `form:"file"`
}

type NamedEnum string

const NamedEnumFoo NamedEnum = "foo"

type StructUnionWrapper struct {
	Union StructUnion `form:"union"`
}

type StructUnion struct {
	OfInt    param.Opt[int64]     `form:",omitzero,inline"`
	OfString param.Opt[string]    `form:",omitzero,inline"`
	OfEnum   param.Opt[NamedEnum] `form:",omitzero,inline"`
	OfA      UnionStructA         `form:",omitzero,inline"`
	OfB      UnionStructB         `form:",omitzero,inline"`
	param.APIUnion
}

type ConstantStruct struct {
	Anchor  string `form:"anchor" default:"created_at"`
	Seconds int    `form:"seconds"`
}

type MultipartMarshalerParent struct {
	Middle MultipartMarshalerMiddleNext `form:"middle"`
}

type MultipartMarshalerMiddleNext struct {
	MiddleNext MultipartMarshalerMiddle `form:"middleNext"`
}

type MultipartMarshalerMiddle struct {
	Child int `form:"child"`
}

// expectedForm supplies literal MIME framing without using the encoder under test.
func expectedForm(fields ...[2]string) string {
	var body strings.Builder
	for _, field := range fields {
		body.WriteString("--xxx\nContent-Disposition: form-data; name=\"" + field[0] + "\"\n\n" + field[1] + "\n")
	}
	body.WriteString("--xxx--\n")
	return body.String()
}

var tests = map[string]struct {
	buf string
	val any
}{
	"file": {
		buf: `--xxx
Content-Disposition: form-data; name="file"; filename="anonymous_file"
Content-Type: application/octet-stream

some file contents...
--xxx--
`,
		val: ReaderStruct{
			File: io.Reader(bytes.NewBuffer([]byte("some file contents..."))),
		},
	},
	"eight_bit_integers": {expectedForm([2]string{"max", "255"}, [2]string{"min", "-128"}, [2]string{"zero", "0"}), map[string]any{"max": uint8(255), "min": int8(-128), "zero": int8(0)}},
	"map_string": {
		expectedForm(
			[2]string{"foo", "bar"},
		),
		map[string]string{"foo": "bar"},
	},

	"map_interface": {
		expectedForm(
			[2]string{"a", "1"},
			[2]string{"b", "str"},
			[2]string{"c", "false"},
		),
		map[string]any{"a": float64(1), "b": "str", "c": false},
	},

	"primitive_struct": {
		expectedForm(
			[2]string{"a", "false"},
			[2]string{"b", "237628372683"},
			[2]string{"c", "654"},
			[2]string{"d", "9999.43"},
			[2]string{"e", "43.76"},
			[2]string{"f.0", "1"},
			[2]string{"f.1", "2"},
			[2]string{"f.2", "3"},
			[2]string{"f.3", "4"},
		),
		Primitives{A: false, B: 237628372683, C: uint(654), D: 9999.43, E: 43.76, F: []int{1, 2, 3, 4}},
	},
	"primitive_struct,brackets": {
		expectedForm(
			[2]string{"f[]", "1"},
			[2]string{"f[]", "2"},
			[2]string{"f[]", "3"},
			[2]string{"f[]", "4"},
		),
		PrimitivesBrackets{F: []int_{1, 2, 3, 4}},
	},

	"slices": {
		expectedForm(
			[2]string{"slices.0.a", "false"},
			[2]string{"slices.0.b", "237628372683"},
			[2]string{"slices.0.c", "654"},
			[2]string{"slices.0.d", "9999.43"},
			[2]string{"slices.0.e", "43.76"},
			[2]string{"slices.0.f.0", "1"},
			[2]string{"slices.0.f.1", "2"},
			[2]string{"slices.0.f.2", "3"},
			[2]string{"slices.0.f.3", "4"},
		),
		Slices{
			Slice: []Primitives{{A: false, B: 237628372683, C: uint(654), D: 9999.43, E: 43.76, F: []int{1, 2, 3, 4}}},
		},
	},
	"primitive_pointer_struct": {
		expectedForm(
			[2]string{"a", "false"},
			[2]string{"b", "237628372683"},
			[2]string{"c", "654"},
			[2]string{"d", "9999.43"},
			[2]string{"e", "43.76"},
			[2]string{"f.0", "1"},
			[2]string{"f.1", "2"},
			[2]string{"f.2", "3"},
			[2]string{"f.3", "4"},
			[2]string{"f.4", "5"},
		),
		PrimitivePointers{
			A: P(false),
			B: P(237628372683),
			C: P(uint(654)),
			D: P(9999.43),
			E: P(float32(43.76)),
			F: &[]int{1, 2, 3, 4, 5},
		},
	},

	"datetime_struct": {
		expectedForm(
			[2]string{"date", "2006-01-02"},
			[2]string{"date-time", "2006-01-02T15:04:05Z"},
		),
		DateTime{
			Date:     time.Date(2006, time.January, 2, 0, 0, 0, 0, time.UTC),
			DateTime: time.Date(2006, time.January, 2, 15, 4, 5, 0, time.UTC),
		},
	},

	"additional_properties": {
		expectedForm(
			[2]string{"a", "true"},
			[2]string{"bar", "value"},
			[2]string{"foo", "true"},
		),
		AdditionalProperties{
			A: true,
			Extras: map[string]any{
				"bar": "value",
				"foo": true,
			},
		},
	},
	"recursive_struct,brackets": {
		expectedForm(
			[2]string{"child[name]", "Alex"},
			[2]string{"name", "Robert"},
		),
		Recursive{Name: "Robert", Child: &Recursive{Name: "Alex"}},
	},

	"recursive_struct": {
		expectedForm(
			[2]string{"child.name", "Alex"},
			[2]string{"name", "Robert"},
		),
		Recursive{Name: "Robert", Child: &Recursive{Name: "Alex"}},
	},

	"unknown_struct_number": {
		expectedForm(
			[2]string{"unknown", "12"},
		),
		UnknownStruct{
			Unknown: 12.,
		},
	},

	"unknown_struct_map": {
		expectedForm(
			[2]string{"unknown.foo", "bar"},
		),
		UnknownStruct{
			Unknown: map[string]any{
				"foo": "bar",
			},
		},
	},

	"struct_union_integer": {
		expectedForm(
			[2]string{"union", "12"},
		),
		StructUnionWrapper{
			Union: StructUnion{OfInt: param.NewOpt[int64](12)},
		},
	},

	"union_integer": {
		expectedForm(
			[2]string{"union", "12"},
		),
		UnionStruct{
			Union: UnionInteger(12),
		},
	},

	"struct_union_struct_discriminated_a": {
		expectedForm(
			[2]string{"union.a", "foo"},
			[2]string{"union.b", "bar"},
			[2]string{"union.type", "typeA"},
		),
		StructUnionWrapper{
			Union: StructUnion{OfA: UnionStructA{
				Type: "typeA",
				A:    "foo",
				B:    "bar",
			}},
		},
	},

	"union_struct_discriminated_a": {
		expectedForm(
			[2]string{"union.a", "foo"},
			[2]string{"union.b", "bar"},
			[2]string{"union.type", "typeA"},
		),

		UnionStruct{
			Union: UnionStructA{
				Type: "typeA",
				A:    "foo",
				B:    "bar",
			},
		},
	},

	"struct_union_struct_discriminated_b": {
		expectedForm(
			[2]string{"union.a", "foo"},
			[2]string{"union.type", "typeB"},
		),
		StructUnionWrapper{
			Union: StructUnion{OfB: UnionStructB{
				Type: "typeB",
				A:    "foo",
			}},
		},
	},

	"union_struct_discriminated_b": {
		expectedForm(
			[2]string{"union.a", "foo"},
			[2]string{"union.type", "typeB"},
		),
		UnionStruct{
			Union: UnionStructB{
				Type: "typeB",
				A:    "foo",
			},
		},
	},

	"union_struct_time": {
		expectedForm(
			[2]string{"union", "2010-05-23"},
		),
		UnionStruct{
			Union: UnionTime(time.Date(2010, 05, 23, 0, 0, 0, 0, time.UTC)),
		},
	},
	"constant_zero_value": {
		expectedForm(
			[2]string{"anchor", "created_at"},
			[2]string{"seconds", "3600"},
		),
		ConstantStruct{
			Seconds: 3600,
		},
	},
	"constant_explicit_value": {
		expectedForm(
			[2]string{"anchor", "created_at_override"},
			[2]string{"seconds", "3600"},
		),
		ConstantStruct{
			Anchor:  "created_at_override",
			Seconds: 3600,
		},
	},
	"deeply-nested-struct,brackets": {
		expectedForm(
			[2]string{"middle[middleNext][child]", "10"},
		),
		MultipartMarshalerParent{
			Middle: MultipartMarshalerMiddleNext{
				MiddleNext: MultipartMarshalerMiddle{
					Child: 10,
				},
			},
		},
	},
	"deeply-nested-map,brackets": {
		expectedForm(
			[2]string{"middle[middleNext][child]", "10"},
		),
		map[string]any{"middle": map[string]any{"middleNext": map[string]any{"child": 10}}},
	},
}

func TestEncode(t *testing.T) {
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			buf := bytes.NewBuffer(nil)
			writer := multipart.NewWriter(buf)
			err := writer.SetBoundary("xxx")
			if err != nil {
				t.Errorf("setting boundary for %v failed with error %v", test.val, err)
			}

			arrayFmt := "indices:dots"
			if tags := strings.Split(name, ","); len(tags) > 1 {
				arrayFmt = tags[1]
			}

			err = MarshalWithSettings(test.val, writer, arrayFmt)
			if err != nil {
				t.Errorf("serialization of %v failed with error %v", test.val, err)
			}
			err = writer.Close()
			if err != nil {
				t.Errorf("serialization of %v failed with error %v", test.val, err)
			}
			raw := buf.Bytes()
			if string(raw) != strings.ReplaceAll(test.buf, "\n", "\r\n") {
				t.Errorf("expected %+#v to serialize to '%s' but got '%s' (with format %s)", test.val, test.buf, string(raw), arrayFmt)
			}
		})
	}
}

func FuzzMarshalRoundTrip(f *testing.F) {
	f.Add("")
	f.Add("plain text")
	f.Add("quotes \" and slashes \\")
	f.Add("line one\r\nline two\x00")

	f.Fuzz(func(t *testing.T, value string) {
		if len(value) > 64*1024 {
			t.Skip()
		}

		buf := bytes.NewBuffer(nil)
		writer := multipart.NewWriter(buf)
		boundary := writer.Boundary()
		if err := MarshalWithSettings(map[string]string{"value": value}, writer, "indices:dots"); err != nil {
			t.Fatalf("marshal value: %v", err)
		}
		if err := writer.Close(); err != nil {
			t.Fatalf("close multipart writer: %v", err)
		}

		form, err := multipart.NewReader(buf, boundary).ReadForm(128 * 1024)
		if err != nil {
			t.Fatalf("read multipart form: %v", err)
		}
		defer form.RemoveAll()

		values := form.Value["value"]
		if len(values) != 1 || values[0] != value {
			t.Fatalf("round trip changed value: got %q, want %q", values, value)
		}
	})
}
