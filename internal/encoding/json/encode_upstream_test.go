// SPDX-License-Identifier: BSD-3-Clause

// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package json

import (
	"bytes"
	"encoding"
	stdjson "encoding/json"
	"errors"
	"fmt"
	"image"
	"log"
	"maps"
	"math"
	"path"
	"reflect"
	"regexp"
	"runtime"
	"slices"
	"strconv"
	"strings"
	"testing"
	"time"
)

type OptionalsEmpty struct {
	Sr string `json:"sr"`
	So string `json:"so,omitempty"`
	Sw string `json:"-"`

	Ir int `json:"omitempty"` // actually named omitempty, not an option
	Io int `json:"io,omitempty"`

	Slr []string `json:"slr"`
	Slo []string `json:"slo,omitempty"`

	Mr map[string]int `json:"mr"`
	Mo map[string]int `json:",omitempty"`

	Fr float64 `json:"fr"`
	Fo float64 `json:"fo,omitempty"`

	Br bool `json:"br"`
	Bo bool `json:"bo,omitempty"`

	Ur uint `json:"ur"`
	Uo uint `json:"uo,omitempty"`

	Str struct{} `json:"str"`
	Sto struct{} `json:"sto,omitempty"`
}

func TestOmitEmpty(t *testing.T) {
	const want = `{
 "sr": "",
 "omitempty": 0,
 "slr": null,
 "mr": {},
 "fr": 0,
 "br": false,
 "ur": 0,
 "str": {},
 "sto": {}
}`
	var o OptionalsEmpty
	o.Sw = "something"
	o.Mr = map[string]int{}
	o.Mo = map[string]int{}

	got, err := marshalIndent(&o, "", " ")
	if err != nil {
		t.Fatalf("MarshalIndent error: %v", err)
	}
	if got := string(got); got != want {
		t.Errorf("MarshalIndent:\n\tgot:  %s\n\twant: %s\n", indentNewlines(got), indentNewlines(want))
	}
}

type NonZeroStruct struct{}

func (nzs NonZeroStruct) IsZero() bool {
	return false
}

type NoPanicStruct struct {
	Int int `json:"int,omitzero"`
}

func (nps *NoPanicStruct) IsZero() bool {
	return nps.Int != 0
}

type OptionalsZero struct {
	Sr string `json:"sr"`
	So string `json:"so,omitzero"`
	Sw string `json:"-"`

	Ir int `json:"omitzero"` // actually named omitzero, not an option
	Io int `json:"io,omitzero"`

	Slr       []string `json:"slr"`
	Slo       []string `json:"slo,omitzero"`
	SloNonNil []string `json:"slononnil,omitzero"`

	Mr  map[string]int `json:"mr"`
	Mo  map[string]int `json:",omitzero"`
	Moo map[string]int `json:"moo,omitzero"`

	Fr   float64    `json:"fr"`
	Fo   float64    `json:"fo,omitzero"`
	Foo  float64    `json:"foo,omitzero"`
	Foo2 [2]float64 `json:"foo2,omitzero"`

	Br bool `json:"br"`
	Bo bool `json:"bo,omitzero"`

	Ur uint `json:"ur"`
	Uo uint `json:"uo,omitzero"`

	Str struct{} `json:"str"`
	Sto struct{} `json:"sto,omitzero"`

	Time      time.Time     `json:"time,omitzero"`
	TimeLocal time.Time     `json:"timelocal,omitzero"`
	Nzs       NonZeroStruct `json:"nzs,omitzero"`

	NilIsZeroer    isZeroer       `json:"niliszeroer,omitzero"`    // nil interface
	NonNilIsZeroer isZeroer       `json:"nonniliszeroer,omitzero"` // non-nil interface
	NoPanicStruct0 isZeroer       `json:"nps0,omitzero"`           // non-nil interface with nil pointer
	NoPanicStruct1 isZeroer       `json:"nps1,omitzero"`           // non-nil interface with non-nil pointer
	NoPanicStruct2 *NoPanicStruct `json:"nps2,omitzero"`           // nil pointer
	NoPanicStruct3 *NoPanicStruct `json:"nps3,omitzero"`           // non-nil pointer
	NoPanicStruct4 NoPanicStruct  `json:"nps4,omitzero"`           // concrete type
}

func TestOmitZero(t *testing.T) {
	const want = `{
 "sr": "",
 "omitzero": 0,
 "slr": null,
 "slononnil": [],
 "mr": {},
 "Mo": {},
 "fr": 0,
 "br": false,
 "ur": 0,
 "str": {},
 "nzs": {},
 "nps1": {},
 "nps3": {},
 "nps4": {}
}`
	var o OptionalsZero
	o.Sw = "something"
	o.SloNonNil = make([]string, 0)
	o.Mr = map[string]int{}
	o.Mo = map[string]int{}

	o.Foo = -0
	o.Foo2 = [2]float64{+0, -0}

	o.TimeLocal = time.Time{}.Local()

	o.NonNilIsZeroer = time.Time{}
	o.NoPanicStruct0 = (*NoPanicStruct)(nil)
	o.NoPanicStruct1 = &NoPanicStruct{}
	o.NoPanicStruct3 = &NoPanicStruct{}

	got, err := marshalIndent(&o, "", " ")
	if err != nil {
		t.Fatalf("MarshalIndent error: %v", err)
	}
	if got := string(got); got != want {
		t.Errorf("MarshalIndent:\n\tgot:  %s\n\twant: %s\n", indentNewlines(got), indentNewlines(want))
	}
}

func TestOmitZeroMap(t *testing.T) {
	const want = `{
 "foo": {
  "sr": "",
  "omitzero": 0,
  "slr": null,
  "mr": null,
  "fr": 0,
  "br": false,
  "ur": 0,
  "str": {},
  "nzs": {},
  "nps4": {}
 }
}`
	m := map[string]OptionalsZero{"foo": {}}
	got, err := marshalIndent(m, "", " ")
	if err != nil {
		t.Fatalf("MarshalIndent error: %v", err)
	}
	if got := string(got); got != want {
		fmt.Println(got)
		t.Errorf("MarshalIndent:\n\tgot:  %s\n\twant: %s\n", indentNewlines(got), indentNewlines(want))
	}
}

type OptionalsEmptyZero struct {
	Sr string `json:"sr"`
	So string `json:"so,omitempty,omitzero"`
	Sw string `json:"-"`

	Io int `json:"io,omitempty,omitzero"`

	Slr       []string `json:"slr"`
	Slo       []string `json:"slo,omitempty,omitzero"`
	SloNonNil []string `json:"slononnil,omitempty,omitzero"`

	Mr map[string]int `json:"mr"`
	Mo map[string]int `json:",omitempty,omitzero"`

	Fr float64 `json:"fr"`
	Fo float64 `json:"fo,omitempty,omitzero"`

	Br bool `json:"br"`
	Bo bool `json:"bo,omitempty,omitzero"`

	Ur uint `json:"ur"`
	Uo uint `json:"uo,omitempty,omitzero"`

	Str struct{} `json:"str"`
	Sto struct{} `json:"sto,omitempty,omitzero"`

	Time time.Time     `json:"time,omitempty,omitzero"`
	Nzs  NonZeroStruct `json:"nzs,omitempty,omitzero"`
}

func TestOmitEmptyZero(t *testing.T) {
	const want = `{
 "sr": "",
 "slr": null,
 "mr": {},
 "fr": 0,
 "br": false,
 "ur": 0,
 "str": {},
 "nzs": {}
}`
	var o OptionalsEmptyZero
	o.Sw = "something"
	o.SloNonNil = make([]string, 0)
	o.Mr = map[string]int{}
	o.Mo = map[string]int{}

	got, err := marshalIndent(&o, "", " ")
	if err != nil {
		t.Fatalf("MarshalIndent error: %v", err)
	}
	if got := string(got); got != want {
		t.Errorf("MarshalIndent:\n\tgot:  %s\n\twant: %s\n", indentNewlines(got), indentNewlines(want))
	}
}

type StringTag struct {
	BoolStr    bool    `json:",string"`
	IntStr     int64   `json:",string"`
	UintptrStr uintptr `json:",string"`
	StrStr     string  `json:",string"`
}

func TestRoundtripStringTag(t *testing.T) {
	tests := []struct {
		CaseName
		in   StringTag
		want string // empty to just test that we roundtrip
	}{{
		CaseName: Name("AllTypes"),
		in: StringTag{
			BoolStr:    true,
			IntStr:     42,
			UintptrStr: 44,
			StrStr:     "xzbit",
		},
		want: `{
	"BoolStr": "true",
	"IntStr": "42",
	"UintptrStr": "44",
	"StrStr": "\"xzbit\""
}`,
	}, {
		// See golang.org/issues/38173.
		CaseName: Name("StringDoubleEscapes"),
		in: StringTag{
			StrStr: "\b\f\n\r\t\"\\",
		},
		want: `{
	"BoolStr": "false",
	"IntStr": "0",
	"UintptrStr": "0",
	"StrStr": "\"\\b\\f\\n\\r\\t\\\"\\\\\""
}`,
	}}
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			got, err := marshalIndent(&tt.in, "", "\t")
			if err != nil {
				t.Fatalf("%s: MarshalIndent error: %v", tt.Where, err)
			}
			if got := string(got); got != tt.want {
				t.Fatalf("%s: MarshalIndent:\n\tgot:  %s\n\twant: %s", tt.Where, stripWhitespace(got), stripWhitespace(tt.want))
			}

			// Verify that it round-trips.
			var s2 StringTag
			if err := stdjson.Unmarshal(got, &s2); err != nil {
				t.Fatalf("%s: Decode error: %v", tt.Where, err)
			}
			if !reflect.DeepEqual(s2, tt.in) {
				t.Fatalf("%s: Decode:\n\tinput: %s\n\tgot:  %#v\n\twant: %#v", tt.Where, indentNewlines(string(got)), s2, tt.in)
			}
		})
	}
}

// byte slices are special even if they're renamed types.
type renamedByte byte
type renamedByteSlice []byte
type renamedRenamedByteSlice []renamedByte

func TestEncodeRenamedByteSlice(t *testing.T) {
	s := renamedByteSlice("abc")
	got, err := Marshal(s)
	if err != nil {
		t.Fatalf("Marshal error: %v", err)
	}
	want := `"YWJj"`
	if string(got) != want {
		t.Errorf("Marshal:\n\tgot:  %s\n\twant: %s", got, want)
	}
	r := renamedRenamedByteSlice("abc")
	got, err = Marshal(r)
	if err != nil {
		t.Fatalf("Marshal error: %v", err)
	}
	if string(got) != want {
		t.Errorf("Marshal:\n\tgot:  %s\n\twant: %s", got, want)
	}
}

type SamePointerNoCycle struct {
	Ptr1, Ptr2 *SamePointerNoCycle
}

var samePointerNoCycle = &SamePointerNoCycle{}

type PointerCycle struct {
	Ptr *PointerCycle
}

var pointerCycle = &PointerCycle{}

type PointerCycleIndirect struct {
	Ptrs []element
}

func (*PointerCycleIndirect) element() {}

type RecursiveSlice []RecursiveSlice

type RecursiveMap map[string]RecursiveMap

var (
	pointerCycleIndirect = &PointerCycleIndirect{}
	mapCycle             = RecursiveMap{}
	sliceCycle           = RecursiveSlice{nil}
	sliceNoCycle         = RecursiveSlice{nil, nil}
	recursiveSliceCycle  = []RecursiveSlice{nil}
)

func init() {
	ptr := &SamePointerNoCycle{}
	samePointerNoCycle.Ptr1 = ptr
	samePointerNoCycle.Ptr2 = ptr

	pointerCycle.Ptr = pointerCycle
	pointerCycleIndirect.Ptrs = []element{pointerCycleIndirect}

	mapCycle["x"] = mapCycle
	sliceCycle[0] = sliceCycle
	sliceNoCycle[1] = sliceNoCycle[:1]
	for i := startDetectingCyclesAfter; i > 0; i-- {
		sliceNoCycle = RecursiveSlice{sliceNoCycle}
	}
	recursiveSliceCycle[0] = recursiveSliceCycle
}

func TestSamePointerNoCycle(t *testing.T) {
	if _, err := Marshal(samePointerNoCycle); err != nil {
		t.Fatalf("Marshal error: %v", err)
	}
}

func TestSliceNoCycle(t *testing.T) {
	if _, err := Marshal(sliceNoCycle); err != nil {
		t.Fatalf("Marshal error: %v", err)
	}
}

func TestUnsupportedValues(t *testing.T) {
	tests := []struct {
		CaseName
		in func() ([]byte, error)
	}{
		{Name(""), marshalOf(math.NaN())},
		{Name(""), marshalOf(math.Inf(-1))},
		{Name(""), marshalOf(math.Inf(1))},
		{Name(""), marshalOf(pointerCycle)},
		{Name(""), marshalOf(pointerCycleIndirect)},
		{Name(""), marshalOf(mapCycle)},
		{Name(""), marshalOf(sliceCycle)},
		{Name(""), marshalOf(recursiveSliceCycle)},
	}
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			if _, err := tt.in(); err != nil {
				if _, ok := err.(*UnsupportedValueError); !ok {
					t.Errorf("%s: Marshal error:\n\tgot:  %T\n\twant: %T", tt.Where, err, new(UnsupportedValueError))
				}
			} else {
				t.Errorf("%s: Marshal error: got nil, want non-nil", tt.Where)
			}
		})
	}
}

// Issue 43207
func TestMarshalTextFloatMap(t *testing.T) {
	m := map[textfloat]string{
		textfloat(math.NaN()): "1",
		textfloat(math.NaN()): "1",
	}
	got, err := Marshal(m)
	if err != nil {
		t.Errorf("Marshal error: %v", err)
	}
	want := `{"TF:NaN":"1","TF:NaN":"1"}`
	if string(got) != want {
		t.Errorf("Marshal:\n\tgot:  %s\n\twant: %s", got, want)
	}
}

// Ref has Marshaler and Unmarshaler methods with pointer receiver.
type Ref int

func (*Ref) MarshalJSON() ([]byte, error) {
	return []byte(`"ref"`), nil
}

func (r *Ref) UnmarshalJSON([]byte) error {
	*r = 12
	return nil
}

// Val has Marshaler methods with value receiver.
type Val int

func (Val) MarshalJSON() ([]byte, error) {
	return []byte(`"val"`), nil
}

// RefText has Marshaler and Unmarshaler methods with pointer receiver.
type RefText int

func (*RefText) MarshalText() ([]byte, error) {
	return []byte(`"ref"`), nil
}

func (r *RefText) UnmarshalText([]byte) error {
	*r = 13
	return nil
}

// ValText has Marshaler methods with value receiver.
type ValText int

func (ValText) MarshalText() ([]byte, error) {
	return []byte(`"val"`), nil
}

func TestRefValMarshal(t *testing.T) {
	var s = struct {
		R0 Ref
		R1 *Ref
		R2 RefText
		R3 *RefText
		V0 Val
		V1 *Val
		V2 ValText
		V3 *ValText
	}{
		R0: 12,
		R1: new(Ref),
		R2: 14,
		R3: new(RefText),
		V0: 13,
		V1: new(Val),
		V2: 15,
		V3: new(ValText),
	}
	const want = `{"R0":"ref","R1":"ref","R2":"\"ref\"","R3":"\"ref\"","V0":"val","V1":"val","V2":"\"val\"","V3":"\"val\""}`
	b, err := Marshal(&s)
	if err != nil {
		t.Fatalf("Marshal error: %v", err)
	}
	if got := string(b); got != want {
		t.Errorf("Marshal:\n\tgot:  %s\n\twant: %s", got, want)
	}
}

// C implements Marshaler and returns unescaped JSON.
type C int

func (C) MarshalJSON() ([]byte, error) {
	return []byte(`"<&>"`), nil
}

// CText implements Marshaler and returns unescaped text.
type CText int

func (CText) MarshalText() ([]byte, error) {
	return []byte(`"<&>"`), nil
}

func TestMarshalerEscaping(t *testing.T) {
	var c C
	want := `"\u003c\u0026\u003e"`
	b, err := Marshal(c)
	if err != nil {
		t.Fatalf("Marshal error: %v", err)
	}
	if got := string(b); got != want {
		t.Errorf("Marshal:\n\tgot:  %s\n\twant: %s", got, want)
	}

	var ct CText
	want = `"\"\u003c\u0026\u003e\""`
	b, err = Marshal(ct)
	if err != nil {
		t.Fatalf("Marshal error: %v", err)
	}
	if got := string(b); got != want {
		t.Errorf("Marshal:\n\tgot:  %s\n\twant: %s", got, want)
	}
}

func TestAnonymousFields(t *testing.T) {
	tests := []struct {
		CaseName
		makeInput func() ([]byte, error) // Function to marshal the input value
		want      string                 // Expected JSON output
	}{{
		// Both S1 and S2 have a field named X. From the perspective of S,
		// it is ambiguous which one X refers to.
		// This should not serialize either field.
		CaseName: Name("AmbiguousField"),
		makeInput: func() ([]byte, error) {
			type (
				S1 struct{ x, X int }
				S2 struct{ x, X int }
				S  struct {
					S1
					S2
				}
			)
			return Marshal(S{S1{1, 2}, S2{3, 4}})
		},
		want: `{}`,
	}, {
		CaseName: Name("DominantField"),
		// Both S1 and S2 have a field named X, but since S has an X field as
		// well, it takes precedence over S1.X and S2.X.
		makeInput: func() ([]byte, error) {
			type (
				S1 struct{ x, X int }
				S2 struct{ x, X int }
				S  struct {
					S1
					S2
					x, X int
				}
			)
			return Marshal(S{S1{1, 2}, S2{3, 4}, 5, 6})
		},
		want: `{"X":6}`,
	}, {
		// Unexported embedded field of non-struct type should not be serialized.
		CaseName: Name("UnexportedEmbeddedInt"),
		makeInput: func() ([]byte, error) {
			type (
				myInt int
				S     struct{ myInt }
			)
			return Marshal(S{5})
		},
		want: `{}`,
	}, {
		// Exported embedded field of non-struct type should be serialized.
		CaseName: Name("ExportedEmbeddedInt"),
		makeInput: func() ([]byte, error) {
			type (
				MyInt int
				S     struct{ MyInt }
			)
			return Marshal(S{5})
		},
		want: `{"MyInt":5}`,
	}, {
		// Unexported embedded field of pointer to non-struct type
		// should not be serialized.
		CaseName: Name("UnexportedEmbeddedIntPointer"),
		makeInput: func() ([]byte, error) {
			type (
				myInt int
				S     struct{ *myInt }
			)
			s := S{new(myInt)}
			*s.myInt = 5
			return Marshal(s)
		},
		want: `{}`,
	}, {
		// Exported embedded field of pointer to non-struct type
		// should be serialized.
		CaseName: Name("ExportedEmbeddedIntPointer"),
		makeInput: func() ([]byte, error) {
			type (
				MyInt int
				S     struct{ *MyInt }
			)
			s := S{new(MyInt)}
			*s.MyInt = 5
			return Marshal(s)
		},
		want: `{"MyInt":5}`,
	}, {
		// Exported fields of embedded structs should have their
		// exported fields be serialized regardless of whether the struct types
		// themselves are exported.
		CaseName: Name("EmbeddedStruct"),
		makeInput: func() ([]byte, error) {
			type (
				s1 struct{ x, X int }
				S2 struct{ y, Y int }
				S  struct {
					s1
					S2
				}
			)
			return Marshal(S{s1{1, 2}, S2{3, 4}})
		},
		want: `{"X":2,"Y":4}`,
	}, {
		// Exported fields of pointers to embedded structs should have their
		// exported fields be serialized regardless of whether the struct types
		// themselves are exported.
		CaseName: Name("EmbeddedStructPointer"),
		makeInput: func() ([]byte, error) {
			type (
				s1 struct{ x, X int }
				S2 struct{ y, Y int }
				S  struct {
					*s1
					*S2
				}
			)
			return Marshal(S{&s1{1, 2}, &S2{3, 4}})
		},
		want: `{"X":2,"Y":4}`,
	}, {
		// Exported fields on embedded unexported structs at multiple levels
		// of nesting should still be serialized.
		CaseName: Name("NestedStructAndInts"),
		makeInput: func() ([]byte, error) {
			type (
				MyInt1 int
				MyInt2 int
				myInt  int
				s2     struct {
					MyInt2
					myInt
				}
				s1 struct {
					MyInt1
					myInt
					s2
				}
				S struct {
					s1
					myInt
				}
			)
			return Marshal(S{s1{1, 2, s2{3, 4}}, 6})
		},
		want: `{"MyInt1":1,"MyInt2":3}`,
	}, {
		// If an anonymous struct pointer field is nil, we should ignore
		// the embedded fields behind it. Not properly doing so may
		// result in the wrong output or reflect panics.
		CaseName: Name("EmbeddedFieldBehindNilPointer"),
		makeInput: func() ([]byte, error) {
			type (
				S2 struct{ Field string }
				S  struct{ *S2 }
			)
			return Marshal(S{})
		},
		want: `{}`,
	}}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			b, err := tt.makeInput()
			if err != nil {
				t.Fatalf("%s: Marshal error: %v", tt.Where, err)
			}
			if string(b) != tt.want {
				t.Fatalf("%s: Marshal:\n\tgot:  %s\n\twant: %s", tt.Where, b, tt.want)
			}
		})
	}
}

type BugA struct {
	S string
}

type BugB struct {
	BugA
	S string
}

type BugC struct {
	S string
}

// Legal Go: We never use the repeated embedded field (S).
type BugX struct {
	A int
	BugA
	BugB
}

// golang.org/issue/16042.
// Even if a nil interface value is passed in, as long as
// it implements Marshaler, it should be marshaled.
type nilJSONMarshaler string

func (*nilJSONMarshaler) element() {}

func (nm *nilJSONMarshaler) MarshalJSON() ([]byte, error) {
	if nm == nil {
		return Marshal("0zenil0")
	}
	return Marshal("zenil:" + string(*nm))
}

// golang.org/issue/34235.
// Even if a nil interface value is passed in, as long as
// it implements encoding.TextMarshaler, it should be marshaled.
type nilTextMarshaler string

func (*nilTextMarshaler) element() {}

func (nm *nilTextMarshaler) MarshalText() ([]byte, error) {
	if nm == nil {
		return []byte("0zenil0"), nil
	}
	return []byte("zenil:" + string(*nm)), nil
}

// See golang.org/issue/16042 and golang.org/issue/34235.
func TestNilMarshal(t *testing.T) {
	tests := []struct {
		CaseName
		in   func() ([]byte, error)
		want string
	}{
		{Name(""), marshalOf[element](nil), `null`},
		{Name(""), marshalOf(new(float64)), `0`},
		{Name(""), marshalOf([]element(nil)), `null`},
		{Name(""), marshalOf([]string(nil)), `null`},
		{Name(""), marshalOf(map[string]string(nil)), `null`},
		{Name(""), marshalOf([]byte(nil)), `null`},
		{Name(""), marshalOf(struct{ M string }{"gopher"}), `{"M":"gopher"}`},
		{Name(""), marshalOf(struct{ M Marshaler }{}), `{"M":null}`},
		{Name(""), marshalOf(struct{ M Marshaler }{(*nilJSONMarshaler)(nil)}), `{"M":"0zenil0"}`},
		{Name(""), marshalOf(struct{ M element }{(*nilJSONMarshaler)(nil)}), `{"M":null}`},
		{Name(""), marshalOf(struct{ M encoding.TextMarshaler }{}), `{"M":null}`},
		{Name(""), marshalOf(struct{ M encoding.TextMarshaler }{(*nilTextMarshaler)(nil)}), `{"M":"0zenil0"}`},
		{Name(""), marshalOf(struct{ M element }{(*nilTextMarshaler)(nil)}), `{"M":null}`},
	}
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			switch got, err := tt.in(); {
			case err != nil:
				t.Fatalf("%s: Marshal error: %v", tt.Where, err)
			case string(got) != tt.want:
				t.Fatalf("%s: Marshal:\n\tgot:  %s\n\twant: %s", tt.Where, got, tt.want)
			}
		})
	}
}

// Issue 5245.
func TestEmbeddedBug(t *testing.T) {
	v := BugB{
		BugA{"A"},
		"B",
	}
	b, err := Marshal(v)
	if err != nil {
		t.Fatal("Marshal error:", err)
	}
	want := `{"S":"B"}`
	got := string(b)
	if got != want {
		t.Fatalf("Marshal:\n\tgot:  %s\n\twant: %s", got, want)
	}
	// Now check that the duplicate field, S, does not appear.
	x := BugX{
		A: 23,
	}
	b, err = Marshal(x)
	if err != nil {
		t.Fatal("Marshal error:", err)
	}
	want = `{"A":23}`
	got = string(b)
	if got != want {
		t.Fatalf("Marshal:\n\tgot:  %s\n\twant: %s", got, want)
	}
}

type BugD struct { // Same as BugA after tagging.
	XXX string `json:"S"`
}

// BugD's tagged S field should dominate BugA's.
type BugY struct {
	BugA
	BugD
}

// Test that a field with a tag dominates untagged fields.
func TestTaggedFieldDominates(t *testing.T) {
	v := BugY{
		BugA{"BugA"},
		BugD{"BugD"},
	}
	b, err := Marshal(v)
	if err != nil {
		t.Fatal("Marshal error:", err)
	}
	want := `{"S":"BugD"}`
	got := string(b)
	if got != want {
		t.Fatalf("Marshal:\n\tgot:  %s\n\twant: %s", got, want)
	}
}

// There are no tags here, so S should not appear.
type BugZ struct {
	BugA
	BugC
	BugY // Contains a tagged S field through BugD; should not dominate.
}

func TestDuplicatedFieldDisappears(t *testing.T) {
	v := BugZ{
		BugA{"BugA"},
		BugC{"BugC"},
		BugY{
			BugA{"nested BugA"},
			BugD{"nested BugD"},
		},
	}
	b, err := Marshal(v)
	if err != nil {
		t.Fatal("Marshal error:", err)
	}
	want := `{}`
	got := string(b)
	if got != want {
		t.Fatalf("Marshal:\n\tgot:  %s\n\twant: %s", got, want)
	}
}

func TestMarshalErrorAndReuseEncodeState(t *testing.T) {
	// Trigger an error in Marshal with cyclic data.
	type Dummy struct {
		Name string
		Next *Dummy
	}
	dummy := Dummy{Name: "Dummy"}
	dummy.Next = &dummy
	if _, err := Marshal(dummy); err == nil {
		t.Errorf("Marshal error: got nil, want non-nil")
	}

	type Data struct {
		A string
		I int
	}
	want := Data{A: "a", I: 1}
	b, err := Marshal(want)
	if err != nil {
		t.Errorf("Marshal error: %v", err)
	}

	var got Data
	if err := stdjson.Unmarshal(b, &got); err != nil {
		t.Errorf("Unmarshal error: %v", err)
	}
	if got != want {
		t.Errorf("Unmarshal:\n\tgot:  %v\n\twant: %v", got, want)
	}
}

// golang.org/issue/8582
func TestEncodePointerString(t *testing.T) {
	type stringPointer struct {
		N *int64 `json:"n,string"`
	}
	var n int64 = 42
	b, err := Marshal(stringPointer{N: &n})
	if err != nil {
		t.Fatalf("Marshal error: %v", err)
	}
	if got, want := string(b), `{"n":"42"}`; got != want {
		t.Fatalf("Marshal:\n\tgot:  %s\n\twant: %s", got, want)
	}
	var back stringPointer
	switch err = stdjson.Unmarshal(b, &back); {
	case err != nil:
		t.Fatalf("Unmarshal error: %v", err)
	case back.N == nil:
		t.Fatalf("Unmarshal: back.N = nil, want non-nil")
	case *back.N != 42:
		t.Fatalf("Unmarshal: *back.N = %d, want 42", *back.N)
	}
}

var encodeStringTests = []struct {
	in  string
	out string
}{
	{"\x00", `"\u0000"`},
	{"\x01", `"\u0001"`},
	{"\x02", `"\u0002"`},
	{"\x03", `"\u0003"`},
	{"\x04", `"\u0004"`},
	{"\x05", `"\u0005"`},
	{"\x06", `"\u0006"`},
	{"\x07", `"\u0007"`},
	{"\x08", `"\b"`},
	{"\x09", `"\t"`},
	{"\x0a", `"\n"`},
	{"\x0b", `"\u000b"`},
	{"\x0c", `"\f"`},
	{"\x0d", `"\r"`},
	{"\x0e", `"\u000e"`},
	{"\x0f", `"\u000f"`},
	{"\x10", `"\u0010"`},
	{"\x11", `"\u0011"`},
	{"\x12", `"\u0012"`},
	{"\x13", `"\u0013"`},
	{"\x14", `"\u0014"`},
	{"\x15", `"\u0015"`},
	{"\x16", `"\u0016"`},
	{"\x17", `"\u0017"`},
	{"\x18", `"\u0018"`},
	{"\x19", `"\u0019"`},
	{"\x1a", `"\u001a"`},
	{"\x1b", `"\u001b"`},
	{"\x1c", `"\u001c"`},
	{"\x1d", `"\u001d"`},
	{"\x1e", `"\u001e"`},
	{"\x1f", `"\u001f"`},
}

func TestEncodeString(t *testing.T) {
	for _, tt := range encodeStringTests {
		b, err := Marshal(tt.in)
		if err != nil {
			t.Errorf("Marshal(%q) error: %v", tt.in, err)
			continue
		}
		out := string(b)
		if out != tt.out {
			t.Errorf("Marshal(%q) = %#q, want %#q", tt.in, out, tt.out)
		}
	}
}

type jsonbyte byte

func (b jsonbyte) MarshalJSON() ([]byte, error) { return fmt.Appendf(nil, `{"JB":%d}`, b), nil }

type textbyte byte

func (b textbyte) MarshalText() ([]byte, error) { return fmt.Appendf(nil, `TB:%d`, b), nil }

type jsonint int

func (i jsonint) MarshalJSON() ([]byte, error) { return fmt.Appendf(nil, `{"JI":%d}`, i), nil }

type textint int

func (i textint) MarshalText() ([]byte, error) { return fmt.Appendf(nil, `TI:%d`, i), nil }

type textfloat float64

func (f textfloat) MarshalText() ([]byte, error) { return fmt.Appendf(nil, `TF:%0.2f`, f), nil }

// Issue 13783
func TestEncodeBytekind(t *testing.T) {
	tests := []struct {
		CaseName
		in   func() ([]byte, error)
		want string
	}{
		{Name(""), marshalOf(byte(7)), "7"},
		{Name(""), marshalOf(jsonbyte(7)), `{"JB":7}`},
		{Name(""), marshalOf(textbyte(4)), `"TB:4"`},
		{Name(""), marshalOf(jsonint(5)), `{"JI":5}`},
		{Name(""), marshalOf(textint(1)), `"TI:1"`},
		{Name(""), marshalOf([]byte{0, 1}), `"AAE="`},
		{Name(""), marshalOf([]jsonbyte{0, 1}), `[{"JB":0},{"JB":1}]`},
		{Name(""), marshalOf([][]jsonbyte{{0, 1}, {3}}), `[[{"JB":0},{"JB":1}],[{"JB":3}]]`},
		{Name(""), marshalOf([]textbyte{2, 3}), `["TB:2","TB:3"]`},
		{Name(""), marshalOf([]jsonint{5, 4}), `[{"JI":5},{"JI":4}]`},
		{Name(""), marshalOf([]textint{9, 3}), `["TI:9","TI:3"]`},
		{Name(""), marshalOf([]int{9, 3}), `[9,3]`},
		{Name(""), marshalOf([]textfloat{12, 3}), `["TF:12.00","TF:3.00"]`},
	}
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			b, err := tt.in()
			if err != nil {
				t.Errorf("%s: Marshal error: %v", tt.Where, err)
			}
			got, want := string(b), tt.want
			if got != want {
				t.Errorf("%s: Marshal:\n\tgot:  %s\n\twant: %s", tt.Where, got, want)
			}
		})
	}
}

func TestTextMarshalerMapKeysAreSorted(t *testing.T) {
	got, err := Marshal(map[unmarshalerText]int{
		{"x", "y"}: 1,
		{"y", "x"}: 2,
		{"a", "z"}: 3,
		{"z", "a"}: 4,
	})
	if err != nil {
		t.Fatalf("Marshal error: %v", err)
	}
	const want = `{"a:z":3,"x:y":1,"y:x":2,"z:a":4}`
	if string(got) != want {
		t.Errorf("Marshal:\n\tgot:  %s\n\twant: %s", got, want)
	}
}

// https://golang.org/issue/33675
func TestNilMarshalerTextMapKey(t *testing.T) {
	got, err := Marshal(map[*unmarshalerText]int{
		(*unmarshalerText)(nil): 1,
		{"A", "B"}:              2,
	})
	if err != nil {
		t.Fatalf("Marshal error: %v", err)
	}
	const want = `{"":1,"A:B":2}`
	if string(got) != want {
		t.Errorf("Marshal:\n\tgot:  %s\n\twant: %s", got, want)
	}
}

var re = regexp.MustCompile

// syntactic checks on form of marshaled floating point numbers.
var badFloatREs = []*regexp.Regexp{
	re(`p`),                     // no binary exponential notation
	re(`^\+`),                   // no leading + sign
	re(`^-?0[^.]`),              // no unnecessary leading zeros
	re(`^-?\.`),                 // leading zero required before decimal point
	re(`\.(e|$)`),               // no trailing decimal
	re(`\.[0-9]+0(e|$)`),        // no trailing zero in fraction
	re(`^-?(0|[0-9]{2,})\..*e`), // exponential notation must have normalized mantissa
	re(`e[0-9]`),                // positive exponent must be signed
	re(`e[+-]0`),                // exponent must not have leading zeros
	re(`e-[1-6]$`),              // not tiny enough for exponential notation
	re(`e+(.|1.|20)$`),          // not big enough for exponential notation
	re(`^-?0\.0000000`),         // too tiny, should use exponential notation
	re(`^-?[0-9]{22}`),          // too big, should use exponential notation
	re(`[1-9][0-9]{16}[1-9]`),   // too many significant digits in integer
	re(`[1-9][0-9.]{17}[1-9]`),  // too many significant digits in decimal
	// below here for float32 only
	re(`[1-9][0-9]{8}[1-9]`),  // too many significant digits in integer
	re(`[1-9][0-9.]{9}[1-9]`), // too many significant digits in decimal
}

func TestMarshalFloat(t *testing.T) {
	t.Parallel()
	nfail := 0
	test := func(f float64, bits int) {
		var bout []byte
		var err error
		if bits == 32 {
			f = float64(float32(f)) // round
			bout, err = Marshal(float32(f))
		} else {
			bout, err = Marshal(f)
		}
		if err != nil {
			t.Errorf("Marshal(float%d(%g)) error: %v", bits, f, err)
			nfail++
			return
		}
		out := string(bout)

		// result must convert back to the same float
		g, err := strconv.ParseFloat(out, bits)
		if err != nil {
			t.Errorf("ParseFloat(%q) error: %v", out, err)
			nfail++
			return
		}
		if f != g || fmt.Sprint(f) != fmt.Sprint(g) { // fmt.Sprint handles ±0
			t.Errorf("ParseFloat(%q):\n\tgot:  %g\n\twant: %g", out, float32(g), f)
			nfail++
			return
		}

		bad := badFloatREs
		if bits == 64 {
			bad = bad[:len(bad)-2]
		}
		for _, re := range bad {
			if re.MatchString(out) {
				t.Errorf("Marshal(float%d(%g)) = %q; must not match /%s/", bits, f, out, re)
				nfail++
				return
			}
		}
	}

	var (
		bigger  = math.Inf(+1)
		smaller = math.Inf(-1)
	)

	var digits = "1.2345678901234567890123"
	for i := len(digits); i >= 2; i-- {
		if testing.Short() && i < len(digits)-4 {
			break
		}
		for exp := -30; exp <= 30; exp++ {
			for _, sign := range "+-" {
				for bits := 32; bits <= 64; bits += 32 {
					s := fmt.Sprintf("%c%se%d", sign, digits[:i], exp)
					f, err := strconv.ParseFloat(s, bits)
					if err != nil {
						log.Fatal(err)
					}
					next := math.Nextafter
					if bits == 32 {
						next = func(g, h float64) float64 {
							return float64(math.Nextafter32(float32(g), float32(h)))
						}
					}
					test(f, bits)
					test(next(f, bigger), bits)
					test(next(f, smaller), bits)
					if nfail > 50 {
						t.Fatalf("stopping test early")
					}
				}
			}
		}
	}
	test(0, 64)
	test(math.Copysign(0, -1), 64)
	test(0, 32)
	test(math.Copysign(0, -1), 32)
}

func TestMarshalRawMessageValue(t *testing.T) {
	type (
		T1 struct {
			M RawMessage `json:",omitempty"`
		}
		T2 struct {
			M *RawMessage `json:",omitempty"`
		}
	)

	var (
		rawNil   = RawMessage(nil)
		rawEmpty = RawMessage([]byte{})
		rawText  = RawMessage([]byte(`"foo"`))
	)

	tests := []struct {
		CaseName
		in   func() ([]byte, error)
		want string
		ok   bool
	}{
		// Test with nil RawMessage.
		{Name(""), marshalOf(rawNil), "null", true},
		{Name(""), marshalOf(&rawNil), "null", true},
		{Name(""), marshalOf([]Marshaler{rawNil}), "[null]", true},
		{Name(""), marshalOf(&[]Marshaler{rawNil}), "[null]", true},
		{Name(""), marshalOf([]Marshaler{&rawNil}), "[null]", true},
		{Name(""), marshalOf(&[]Marshaler{&rawNil}), "[null]", true},
		{Name(""), marshalOf(struct{ M RawMessage }{rawNil}), `{"M":null}`, true},
		{Name(""), marshalOf(&struct{ M RawMessage }{rawNil}), `{"M":null}`, true},
		{Name(""), marshalOf(struct{ M *RawMessage }{&rawNil}), `{"M":null}`, true},
		{Name(""), marshalOf(&struct{ M *RawMessage }{&rawNil}), `{"M":null}`, true},
		{Name(""), marshalOf(map[string]Marshaler{"M": rawNil}), `{"M":null}`, true},
		{Name(""), marshalOf(&map[string]Marshaler{"M": rawNil}), `{"M":null}`, true},
		{Name(""), marshalOf(map[string]Marshaler{"M": &rawNil}), `{"M":null}`, true},
		{Name(""), marshalOf(&map[string]Marshaler{"M": &rawNil}), `{"M":null}`, true},
		{Name(""), marshalOf(T1{rawNil}), "{}", true},
		{Name(""), marshalOf(T2{&rawNil}), `{"M":null}`, true},
		{Name(""), marshalOf(&T1{rawNil}), "{}", true},
		{Name(""), marshalOf(&T2{&rawNil}), `{"M":null}`, true},

		// Test with empty, but non-nil, RawMessage.
		{Name(""), marshalOf(rawEmpty), "", false},
		{Name(""), marshalOf(&rawEmpty), "", false},
		{Name(""), marshalOf([]Marshaler{rawEmpty}), "", false},
		{Name(""), marshalOf(&[]Marshaler{rawEmpty}), "", false},
		{Name(""), marshalOf([]Marshaler{&rawEmpty}), "", false},
		{Name(""), marshalOf(&[]Marshaler{&rawEmpty}), "", false},
		{Name(""), marshalOf(struct{ X RawMessage }{rawEmpty}), "", false},
		{Name(""), marshalOf(&struct{ X RawMessage }{rawEmpty}), "", false},
		{Name(""), marshalOf(struct{ X *RawMessage }{&rawEmpty}), "", false},
		{Name(""), marshalOf(&struct{ X *RawMessage }{&rawEmpty}), "", false},
		{Name(""), marshalOf(map[string]Marshaler{"nil": rawEmpty}), "", false},
		{Name(""), marshalOf(&map[string]Marshaler{"nil": rawEmpty}), "", false},
		{Name(""), marshalOf(map[string]Marshaler{"nil": &rawEmpty}), "", false},
		{Name(""), marshalOf(&map[string]Marshaler{"nil": &rawEmpty}), "", false},
		{Name(""), marshalOf(T1{rawEmpty}), "{}", true},
		{Name(""), marshalOf(T2{&rawEmpty}), "", false},
		{Name(""), marshalOf(&T1{rawEmpty}), "{}", true},
		{Name(""), marshalOf(&T2{&rawEmpty}), "", false},

		// Test with RawMessage with some text.
		//
		// The tests below marked with Issue6458 used to generate "ImZvbyI=" instead "foo".
		// This behavior was intentionally changed in Go 1.8.
		// See https://golang.org/issues/14493#issuecomment-255857318
		{Name(""), marshalOf(rawText), `"foo"`, true}, // Issue6458
		{Name(""), marshalOf(&rawText), `"foo"`, true},
		{Name(""), marshalOf([]Marshaler{rawText}), `["foo"]`, true},  // Issue6458
		{Name(""), marshalOf(&[]Marshaler{rawText}), `["foo"]`, true}, // Issue6458
		{Name(""), marshalOf([]Marshaler{&rawText}), `["foo"]`, true},
		{Name(""), marshalOf(&[]Marshaler{&rawText}), `["foo"]`, true},
		{Name(""), marshalOf(struct{ M RawMessage }{rawText}), `{"M":"foo"}`, true}, // Issue6458
		{Name(""), marshalOf(&struct{ M RawMessage }{rawText}), `{"M":"foo"}`, true},
		{Name(""), marshalOf(struct{ M *RawMessage }{&rawText}), `{"M":"foo"}`, true},
		{Name(""), marshalOf(&struct{ M *RawMessage }{&rawText}), `{"M":"foo"}`, true},
		{Name(""), marshalOf(map[string]Marshaler{"M": rawText}), `{"M":"foo"}`, true},  // Issue6458
		{Name(""), marshalOf(&map[string]Marshaler{"M": rawText}), `{"M":"foo"}`, true}, // Issue6458
		{Name(""), marshalOf(map[string]Marshaler{"M": &rawText}), `{"M":"foo"}`, true},
		{Name(""), marshalOf(&map[string]Marshaler{"M": &rawText}), `{"M":"foo"}`, true},
		{Name(""), marshalOf(T1{rawText}), `{"M":"foo"}`, true}, // Issue6458
		{Name(""), marshalOf(T2{&rawText}), `{"M":"foo"}`, true},
		{Name(""), marshalOf(&T1{rawText}), `{"M":"foo"}`, true},
		{Name(""), marshalOf(&T2{&rawText}), `{"M":"foo"}`, true},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			b, err := tt.in()
			if ok := (err == nil); ok != tt.ok {
				if err != nil {
					t.Errorf("%s: Marshal error: %v", tt.Where, err)
				} else {
					t.Errorf("%s: Marshal error: got nil, want non-nil", tt.Where)
				}
			}
			if got := string(b); got != tt.want {
				t.Errorf("%s: Marshal:\n\tgot:  %s\n\twant: %s", tt.Where, got, tt.want)
			}
		})
	}
}

type marshalPanic struct{}

func (marshalPanic) MarshalJSON() ([]byte, error) { panic(0xdead) }

func TestMarshalPanic(t *testing.T) {
	defer func() {
		if got := fmt.Sprintf("(%[1]T)(%[1]v)", recover()); got != "(int)(57005)" {
			t.Errorf("panic() = %s, want (int)(57005)", got)
		}
	}()
	_, err := Marshal(&marshalPanic{})
	t.Errorf("Marshal returned %v, want a panic", err)
}

func TestMarshalUncommonFieldNames(t *testing.T) {
	v := struct {
		A0, À, Aβ int
	}{}
	b, err := Marshal(v)
	if err != nil {
		t.Fatal("Marshal error:", err)
	}
	want := `{"A0":0,"À":0,"Aβ":0}`
	got := string(b)
	if got != want {
		t.Fatalf("Marshal:\n\tgot:  %s\n\twant: %s", got, want)
	}
}

func TestMarshalerError(t *testing.T) {
	s := "test variable"
	st := reflect.TypeOf(s)
	const errText = "json: test error"

	tests := []struct {
		CaseName
		err  *MarshalerError
		want string
	}{{
		Name(""),
		&MarshalerError{st, errors.New(errText), ""},
		"json: error calling MarshalJSON for type " + st.String() + ": " + errText,
	}, {
		Name(""),
		&MarshalerError{st, errors.New(errText), "TestMarshalerError"},
		"json: error calling TestMarshalerError for type " + st.String() + ": " + errText,
	}}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			got := tt.err.Error()
			if got != tt.want {
				t.Errorf("%s: Error:\n\tgot:  %s\n\twant: %s", tt.Where, got, tt.want)
			}
		})
	}
}

type marshaledValue string

func (v marshaledValue) MarshalJSON() ([]byte, error) {
	return []byte(v), nil
}

func TestIssue63379(t *testing.T) {
	for _, v := range []string{
		"[]<",
		"[]>",
		"[]&",
		"[]\u2028",
		"[]\u2029",
		"{}<",
		"{}>",
		"{}&",
		"{}\u2028",
		"{}\u2029",
	} {
		_, err := Marshal(marshaledValue(v))
		if err == nil {
			t.Errorf("expected error for %q", v)
		}
	}
}

// element lets a table or a field hold values of different types
// through an interface that is not empty.
type element interface{ element() }

type elementFloat float64

func (elementFloat) element() {}

// marshalOf returns a function that marshals v, so that a table can
// hold inputs of different types.
func marshalOf[T any](v T) func() ([]byte, error) {
	return func() ([]byte, error) { return Marshal(v) }
}

// marshalIndent marshals v and indents the result with the standard library.
func marshalIndent[T any](v T, prefix, indent string) ([]byte, error) {
	b, err := Marshal(v)
	if err != nil {
		return nil, err
	}
	var out bytes.Buffer
	err = stdjson.Indent(&out, b, prefix, indent)
	return out.Bytes(), err
}

// TODO(https://go.dev/issue/52751): Replace with native testing support.

// CaseName is a case name annotated with a file and line.
type CaseName struct {
	Name  string
	Where CasePos
}

// Name annotates a case name with the file and line of the caller.
func Name(s string) (c CaseName) {
	c.Name = s
	runtime.Callers(2, c.Where.pc[:])
	return c
}

// CasePos represents a file and line number.
type CasePos struct{ pc [1]uintptr }

func (pos CasePos) String() string {
	frames := runtime.CallersFrames(pos.pc[:])
	frame, _ := frames.Next()
	return fmt.Sprintf("%s:%d", path.Base(frame.File), frame.Line)
}

func indentNewlines(s string) string {
	return strings.Join(strings.Split(s, "\n"), "\n\t")
}

func stripWhitespace(s string) string {
	return strings.Map(func(r rune) rune {
		if r == ' ' || r == '\n' || r == '\r' || r == '\t' {
			return -1
		}
		return r
	}, s)
}

func diff(t *testing.T, a, b []byte) {
	t.Helper()
	for i := 0; ; i++ {
		if i >= len(a) || i >= len(b) || a[i] != b[i] {
			j := i - 10
			if j < 0 {
				j = 0
			}
			t.Errorf("diverge at %d: «%s» vs «%s»", i, trim(a[j:]), trim(b[j:]))
			return
		}
	}
}

func trim(b []byte) []byte {
	return b[:min(len(b), 20)]
}

type unmarshalerText struct {
	A, B string
}

// needed for re-marshaling tests
func (u unmarshalerText) MarshalText() ([]byte, error) {
	return []byte(u.A + ":" + u.B), nil
}

// Test data structures for anonymous fields.

type Point struct {
	Z int
}

type Top struct {
	Level0 int
	Embed0
	*Embed0a
	*Embed0b `json:"e,omitempty"` // treated as named
	Embed0c  `json:"-"`           // ignored
	Loop
	Embed0p // has Point with X, Y, used
	Embed0q // has Point with Z, used
	embed   // contains exported field
}

type Embed0 struct {
	Level1a int // overridden by Embed0a's Level1a with json tag
	Level1b int // used because Embed0a's Level1b is renamed
	Level1c int // used because Embed0a's Level1c is ignored
	Level1d int // annihilated by Embed0a's Level1d
	Level1e int `json:"x"` // annihilated by Embed0a.Level1e
}

type Embed0a struct {
	Level1a int `json:"Level1a,omitempty"`
	Level1b int `json:"LEVEL1B,omitempty"`
	Level1c int `json:"-"`
	Level1d int // annihilated by Embed0's Level1d
	Level1f int `json:"x"` // annihilated by Embed0's Level1e
}

type Embed0b Embed0

type Embed0c Embed0

type Embed0p struct {
	image.Point
}

type Embed0q struct {
	Point
}

type embed struct {
	Q int
}

type Loop struct {
	Loop1 int `json:",omitempty"`
	Loop2 int `json:",omitempty"`
	*Loop
}

func TestMarshal(t *testing.T) {
	b, err := Marshal(allValue)
	if err != nil {
		t.Fatalf("Marshal error: %v", err)
	}
	if string(b) != allValueCompact {
		t.Errorf("Marshal:")
		diff(t, b, []byte(allValueCompact))
		return
	}

	b, err = Marshal(pallValue)
	if err != nil {
		t.Fatalf("Marshal error: %v", err)
	}
	if string(b) != pallValueCompact {
		t.Errorf("Marshal:")
		diff(t, b, []byte(pallValueCompact))
		return
	}
}

func TestMarshalInvalidUTF8(t *testing.T) {
	tests := []struct {
		CaseName
		in   string
		want string
	}{
		{Name(""), "hello\xffworld", `"hello\ufffdworld"`},
		{Name(""), "", `""`},
		{Name(""), "\xff", `"\ufffd"`},
		{Name(""), "\xff\xff", `"\ufffd\ufffd"`},
		{Name(""), "a\xffb", `"a\ufffdb"`},
		{Name(""), "\xe6\x97\xa5\xe6\x9c\xac\xff\xaa\x9e", `"日本\ufffd\ufffd\ufffd"`},
	}
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			got, err := Marshal(tt.in)
			if string(got) != tt.want || err != nil {
				t.Errorf("%s: Marshal(%q):\n\tgot:  (%q, %v)\n\twant: (%q, nil)", tt.Where, tt.in, got, err, tt.want)
			}
		})
	}
}

func TestMarshalEmbeds(t *testing.T) {
	top := &Top{
		Level0: 1,
		Embed0: Embed0{
			Level1b: 2,
			Level1c: 3,
		},
		Embed0a: &Embed0a{
			Level1a: 5,
			Level1b: 6,
		},
		Embed0b: &Embed0b{
			Level1a: 8,
			Level1b: 9,
			Level1c: 10,
			Level1d: 11,
			Level1e: 12,
		},
		Loop: Loop{
			Loop1: 13,
			Loop2: 14,
		},
		Embed0p: Embed0p{
			Point: image.Point{X: 15, Y: 16},
		},
		Embed0q: Embed0q{
			Point: Point{Z: 17},
		},
		embed: embed{
			Q: 18,
		},
	}
	got, err := Marshal(top)
	if err != nil {
		t.Fatalf("Marshal error: %v", err)
	}
	want := "{\"Level0\":1,\"Level1b\":2,\"Level1c\":3,\"Level1a\":5,\"LEVEL1B\":6,\"e\":{\"Level1a\":8,\"Level1b\":9,\"Level1c\":10,\"Level1d\":11,\"x\":12},\"Loop1\":13,\"Loop2\":14,\"X\":15,\"Y\":16,\"Z\":17,\"Q\":18}"
	if string(got) != want {
		t.Errorf("Marshal:\n\tgot:  %s\n\twant: %s", got, want)
	}
}

func TestLargeByteSlice(t *testing.T) {
	s0 := make([]byte, 2000)
	for i := range s0 {
		s0[i] = byte(i)
	}
	b, err := Marshal(s0)
	if err != nil {
		t.Fatalf("Marshal error: %v", err)
	}
	var s1 []byte
	if err := stdjson.Unmarshal(b, &s1); err != nil {
		t.Fatalf("Unmarshal error: %v", err)
	}
	if !bytes.Equal(s0, s1) {
		t.Errorf("Marshal:")
		diff(t, s0, s1)
	}
}

func TestEscape(t *testing.T) {
	const input = `"foobar"<html>` + " [\u2028 \u2029]"
	const want = `"\"foobar\"\u003chtml\u003e [\u2028 \u2029]"`
	got, err := Marshal(input)
	if err != nil {
		t.Fatalf("Marshal error: %v", err)
	}
	if string(got) != want {
		t.Errorf("Marshal(%#q):\n\tgot:  %s\n\twant: %s", input, got, want)
	}
}

type All struct {
	Bool    bool
	Int     int
	Int8    int8
	Int16   int16
	Int32   int32
	Int64   int64
	Uint    uint
	Uint8   uint8
	Uint16  uint16
	Uint32  uint32
	Uint64  uint64
	Uintptr uintptr
	Float32 float32
	Float64 float64

	Foo  string `json:"bar"`
	Foo2 string `json:"bar2"`

	IntStr     int64   `json:",string"`
	UintptrStr uintptr `json:",string"`

	PBool    *bool
	PInt     *int
	PInt8    *int8
	PInt16   *int16
	PInt32   *int32
	PInt64   *int64
	PUint    *uint
	PUint8   *uint8
	PUint16  *uint16
	PUint32  *uint32
	PUint64  *uint64
	PUintptr *uintptr
	PFloat32 *float32
	PFloat64 *float64

	String  string
	PString *string

	Map   map[string]Small
	MapP  map[string]*Small
	PMap  *map[string]Small
	PMapP *map[string]*Small

	EmptyMap map[string]Small
	NilMap   map[string]Small

	Slice   []Small
	SliceP  []*Small
	PSlice  *[]Small
	PSliceP *[]*Small

	EmptySlice []Small
	NilSlice   []Small

	StringSlice []string
	ByteSlice   []byte

	Small   Small
	PSmall  *Small
	PPSmall **Small

	Interface  element
	PInterface *element

	unexported int
}

type Small struct {
	Tag string
}

var allValue = All{
	Bool:       true,
	Int:        2,
	Int8:       3,
	Int16:      4,
	Int32:      5,
	Int64:      6,
	Uint:       7,
	Uint8:      8,
	Uint16:     9,
	Uint32:     10,
	Uint64:     11,
	Uintptr:    12,
	Float32:    14.1,
	Float64:    15.1,
	Foo:        "foo",
	Foo2:       "foo2",
	IntStr:     42,
	UintptrStr: 44,
	String:     "16",
	Map: map[string]Small{
		"17": {Tag: "tag17"},
		"18": {Tag: "tag18"},
	},
	MapP: map[string]*Small{
		"19": {Tag: "tag19"},
		"20": nil,
	},
	EmptyMap:    map[string]Small{},
	Slice:       []Small{{Tag: "tag20"}, {Tag: "tag21"}},
	SliceP:      []*Small{{Tag: "tag22"}, nil, {Tag: "tag23"}},
	EmptySlice:  []Small{},
	StringSlice: []string{"str24", "str25", "str26"},
	ByteSlice:   []byte{27, 28, 29},
	Small:       Small{Tag: "tag30"},
	PSmall:      &Small{Tag: "tag31"},
	Interface:   elementFloat(5.2),
	unexported:  99,
}

var pallValue = All{
	PBool:      &allValue.Bool,
	PInt:       &allValue.Int,
	PInt8:      &allValue.Int8,
	PInt16:     &allValue.Int16,
	PInt32:     &allValue.Int32,
	PInt64:     &allValue.Int64,
	PUint:      &allValue.Uint,
	PUint8:     &allValue.Uint8,
	PUint16:    &allValue.Uint16,
	PUint32:    &allValue.Uint32,
	PUint64:    &allValue.Uint64,
	PUintptr:   &allValue.Uintptr,
	PFloat32:   &allValue.Float32,
	PFloat64:   &allValue.Float64,
	PString:    &allValue.String,
	PMap:       &allValue.Map,
	PMapP:      &allValue.MapP,
	PSlice:     &allValue.Slice,
	PSliceP:    &allValue.SliceP,
	PPSmall:    &allValue.PSmall,
	PInterface: &allValue.Interface,
}

var allValueIndent = `{
	"Bool": true,
	"Int": 2,
	"Int8": 3,
	"Int16": 4,
	"Int32": 5,
	"Int64": 6,
	"Uint": 7,
	"Uint8": 8,
	"Uint16": 9,
	"Uint32": 10,
	"Uint64": 11,
	"Uintptr": 12,
	"Float32": 14.1,
	"Float64": 15.1,
	"bar": "foo",
	"bar2": "foo2",
	"IntStr": "42",
	"UintptrStr": "44",
	"PBool": null,
	"PInt": null,
	"PInt8": null,
	"PInt16": null,
	"PInt32": null,
	"PInt64": null,
	"PUint": null,
	"PUint8": null,
	"PUint16": null,
	"PUint32": null,
	"PUint64": null,
	"PUintptr": null,
	"PFloat32": null,
	"PFloat64": null,
	"String": "16",
	"PString": null,
	"Map": {
		"17": {
			"Tag": "tag17"
		},
		"18": {
			"Tag": "tag18"
		}
	},
	"MapP": {
		"19": {
			"Tag": "tag19"
		},
		"20": null
	},
	"PMap": null,
	"PMapP": null,
	"EmptyMap": {},
	"NilMap": null,
	"Slice": [
		{
			"Tag": "tag20"
		},
		{
			"Tag": "tag21"
		}
	],
	"SliceP": [
		{
			"Tag": "tag22"
		},
		null,
		{
			"Tag": "tag23"
		}
	],
	"PSlice": null,
	"PSliceP": null,
	"EmptySlice": [],
	"NilSlice": null,
	"StringSlice": [
		"str24",
		"str25",
		"str26"
	],
	"ByteSlice": "Gxwd",
	"Small": {
		"Tag": "tag30"
	},
	"PSmall": {
		"Tag": "tag31"
	},
	"PPSmall": null,
	"Interface": 5.2,
	"PInterface": null
}`

var allValueCompact = stripWhitespace(allValueIndent)

var pallValueIndent = `{
	"Bool": false,
	"Int": 0,
	"Int8": 0,
	"Int16": 0,
	"Int32": 0,
	"Int64": 0,
	"Uint": 0,
	"Uint8": 0,
	"Uint16": 0,
	"Uint32": 0,
	"Uint64": 0,
	"Uintptr": 0,
	"Float32": 0,
	"Float64": 0,
	"bar": "",
	"bar2": "",
        "IntStr": "0",
	"UintptrStr": "0",
	"PBool": true,
	"PInt": 2,
	"PInt8": 3,
	"PInt16": 4,
	"PInt32": 5,
	"PInt64": 6,
	"PUint": 7,
	"PUint8": 8,
	"PUint16": 9,
	"PUint32": 10,
	"PUint64": 11,
	"PUintptr": 12,
	"PFloat32": 14.1,
	"PFloat64": 15.1,
	"String": "",
	"PString": "16",
	"Map": null,
	"MapP": null,
	"PMap": {
		"17": {
			"Tag": "tag17"
		},
		"18": {
			"Tag": "tag18"
		}
	},
	"PMapP": {
		"19": {
			"Tag": "tag19"
		},
		"20": null
	},
	"EmptyMap": null,
	"NilMap": null,
	"Slice": null,
	"SliceP": null,
	"PSlice": [
		{
			"Tag": "tag20"
		},
		{
			"Tag": "tag21"
		}
	],
	"PSliceP": [
		{
			"Tag": "tag22"
		},
		null,
		{
			"Tag": "tag23"
		}
	],
	"EmptySlice": null,
	"NilSlice": null,
	"StringSlice": null,
	"ByteSlice": null,
	"Small": {
		"Tag": ""
	},
	"PSmall": null,
	"PPSmall": {
		"Tag": "tag31"
	},
	"Interface": null,
	"PInterface": 5.2
}`

var pallValueCompact = stripWhitespace(pallValueIndent)

func TestStringKind(t *testing.T) {
	type stringKind string
	want := map[stringKind]int{"foo": 42}
	data, err := Marshal(want)
	if err != nil {
		t.Fatalf("Marshal error: %v", err)
	}
	var got map[stringKind]int
	err = stdjson.Unmarshal(data, &got)
	if err != nil {
		t.Fatalf("Unmarshal error: %v", err)
	}
	if !maps.Equal(got, want) {
		t.Fatalf("Marshal/Unmarshal mismatch:\n\tgot:  %v\n\twant: %v", got, want)
	}
}

// Custom types with []byte as underlying type could not be marshaled
// and then unmarshaled.
// Issue 8962.
func TestByteKind(t *testing.T) {
	type byteKind []byte
	want := byteKind("hello")
	data, err := Marshal(want)
	if err != nil {
		t.Fatalf("Marshal error: %v", err)
	}
	var got byteKind
	err = stdjson.Unmarshal(data, &got)
	if err != nil {
		t.Fatalf("Unmarshal error: %v", err)
	}
	if !slices.Equal(got, want) {
		t.Fatalf("Marshal/Unmarshal mismatch:\n\tgot:  %v\n\twant: %v", got, want)
	}
}

// The fix for issue 8962 introduced a regression.
// Issue 12921.
func TestSliceOfCustomByte(t *testing.T) {
	type Uint8 uint8
	want := []Uint8("hello")
	data, err := Marshal(want)
	if err != nil {
		t.Fatalf("Marshal error: %v", err)
	}
	var got []Uint8
	err = stdjson.Unmarshal(data, &got)
	if err != nil {
		t.Fatalf("Unmarshal error: %v", err)
	}
	if !slices.Equal(got, want) {
		t.Fatalf("Marshal/Unmarshal mismatch:\n\tgot:  %v\n\twant: %v", got, want)
	}
}
