// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package json implements encoding of JSON as defined in
// RFC 7159. The mapping between Go values and JSON is described
// in the documentation for the Marshal function.
package json

import (
	"bytes"
	"cmp"
	"encoding"
	"encoding/base64"
	"fmt"
	"github.com/Xquik-dev/x-twitter-scraper-go/internal/encoding/json/sentinel"
	"github.com/Xquik-dev/x-twitter-scraper-go/internal/reflectvalue"
	"math"
	"reflect"
	"slices"
	"strconv"
	"strings"
	"sync"
	"unicode"
	"unicode/utf8"
)

func Marshal[T any](v T, opts ...Option) ([]byte, error) {
	e := &encodeState{ptrSeen: map[cycleKey]struct{}{}}
	if err := e.marshal(reflect.ValueOf(v), encOpts{}.apply(opts...)); err != nil {
		return nil, err
	}
	return e.Bytes(), nil
}

type Marshaler interface {
	MarshalJSON() ([]byte, error)
}

type UnsupportedTypeError struct {
	Type reflect.Type
}

func (e *UnsupportedTypeError) Error() string {
	return "json: unsupported type: " + e.Type.String()
}

type UnsupportedValueError struct {
	Value reflect.Value
	Str   string
}

func (e *UnsupportedValueError) Error() string {
	return "json: unsupported value: " + e.Str
}

type MarshalerError struct {
	Type       reflect.Type
	Err        error
	sourceFunc string
}

func (e *MarshalerError) Error() string {
	srcFunc := e.sourceFunc
	if srcFunc == "" {
		srcFunc = "MarshalJSON"
	}
	return "json: error calling " + srcFunc +
		" for type " + e.Type.String() +
		": " + e.Err.Error()
}

func (e *MarshalerError) Unwrap() error { return e.Err }

const hex = "0123456789abcdef"

type encodeState struct {
	bytes.Buffer

	ptrLevel uint
	ptrSeen  map[cycleKey]struct{}

	err error
}

type cycleKey struct {
	typ reflect.Type
	ptr uintptr
	len int
}

const startDetectingCyclesAfter = 1000

func (e *encodeState) marshal(v reflect.Value, opts encOpts) (err error) {
	defer func() {
		if e.err != nil {
			recover()
			err = e.err
		}
	}()
	e.reflectValue(v, opts)
	return nil
}

func (e *encodeState) error(err error) {
	e.err = err
	panic(err)
}

func isEmptyValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.String:
		return v.Len() == 0
	case reflect.Array, reflect.Map, reflect.Slice:
		return v.Len() == 0
	case reflect.Bool,
		reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr,
		reflect.Float32, reflect.Float64,
		reflect.Interface, reflect.Pointer:
		return v.IsZero()
	}
	return false
}

func (e *encodeState) reflectValue(v reflect.Value, opts encOpts) {
	valueEncoder(v)(e, v, opts)
}

type encOpts struct {
	quoted         bool
	timefmt        string
	skipCompaction bool
}

type encoderFunc func(e *encodeState, v reflect.Value, opts encOpts)

var encoderCache = syncMap[reflect.Type, encoderFunc]{m: map[reflect.Type]encoderFunc{}}

func valueEncoder(v reflect.Value) encoderFunc {
	if !v.IsValid() {
		return invalidValueEncoder
	}
	return typeEncoder(v.Type())
}

func typeEncoder(t reflect.Type) encoderFunc {
	if f, ok := encoderCache.load(t); ok {
		return f
	}

	var (
		wg sync.WaitGroup
		f  encoderFunc
	)
	wg.Add(1)
	fi, loaded := encoderCache.loadOrStore(t, func(e *encodeState, v reflect.Value, opts encOpts) {
		wg.Wait()
		f(e, v, opts)
	})
	if loaded {
		return fi
	}

	f = newTypeEncoder(t, true)
	wg.Done()
	encoderCache.store(t, f)
	return f
}

var (
	marshalerType     = reflect.TypeFor[Marshaler]()
	textMarshalerType = reflect.TypeFor[encoding.TextMarshaler]()
)

func newTypeEncoder(t reflect.Type, allowAddr bool) encoderFunc {
	if t == timeType {
		return timeEncoder
	}

	if t.Kind() != reflect.Pointer && allowAddr && reflect.PointerTo(t).Implements(marshalerType) {
		return newCondAddrEncoder(addrMarshalerEncoder, newTypeEncoder(t, false))
	}

	if t.Implements(marshalerType) {
		return marshalerEncoder
	}
	if t.Kind() != reflect.Pointer && allowAddr && reflect.PointerTo(t).Implements(textMarshalerType) {
		return newCondAddrEncoder(addrTextMarshalerEncoder, newTypeEncoder(t, false))
	}
	if t.Implements(textMarshalerType) {
		return textMarshalerEncoder
	}

	switch t.Kind() {
	case reflect.Bool:
		return boolEncoder
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return intEncoder
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return uintEncoder
	case reflect.Float32:
		return float32Encoder
	case reflect.Float64:
		return float64Encoder
	case reflect.String:
		return stringEncoder
	case reflect.Interface:
		return interfaceEncoder
	case reflect.Struct:
		return newStructEncoder(t)
	case reflect.Map:
		return newMapEncoder(t)
	case reflect.Slice:
		return newSliceEncoder(t)
	case reflect.Array:
		return newArrayEncoder(t)
	case reflect.Pointer:
		return newPtrEncoder(t)
	default:
		return unsupportedTypeEncoder
	}
}

func invalidValueEncoder(e *encodeState, v reflect.Value, _ encOpts) {
	e.Buffer.WriteString("null")
}

func marshalerEncoder(e *encodeState, v reflect.Value, opts encOpts) {
	encodeMarshaler(e, v, v, opts)
}

func addrMarshalerEncoder(e *encodeState, v reflect.Value, opts encOpts) {
	encodeMarshaler(e, v, v.Addr(), opts)
}

func encodeMarshaler(e *encodeState, v, m reflect.Value, opts encOpts) {
	if m.Kind() == reflect.Pointer && m.IsNil() {
		e.Buffer.WriteString("null")
		return
	}
	marshaler, ok := reflect.TypeAssert[Marshaler](m)
	if !ok {
		e.Buffer.WriteString("null")
		return
	}

	b, err := marshalWithTimeLayout(marshaler, opts.timefmt)
	if err == nil {
		e.Grow(len(b))
		out := e.AvailableBuffer()
		out, err = appendCompact(out, b, opts)
		e.Buffer.Write(out)
	}
	if err != nil {
		e.error(&MarshalerError{v.Type(), err, "MarshalJSON"})
	}
}

func textMarshalerEncoder(e *encodeState, v reflect.Value, opts encOpts) {
	encodeTextMarshaler(e, v, v, opts)
}

func addrTextMarshalerEncoder(e *encodeState, v reflect.Value, opts encOpts) {
	encodeTextMarshaler(e, v, v.Addr(), opts)
}

func encodeTextMarshaler(e *encodeState, v, m reflect.Value, opts encOpts) {
	if m.Kind() == reflect.Pointer && m.IsNil() {
		e.Buffer.WriteString("null")
		return
	}
	marshaler, ok := reflect.TypeAssert[encoding.TextMarshaler](m)
	if !ok {
		e.Buffer.WriteString("null")
		return
	}
	b, err := marshaler.MarshalText()
	if err != nil {
		e.error(&MarshalerError{v.Type(), err, "MarshalText"})
	}
	e.Buffer.Write(appendString(e.AvailableBuffer(), b, true))
}

func boolEncoder(e *encodeState, v reflect.Value, opts encOpts) {
	b := e.AvailableBuffer()
	b = mayAppendQuote(b, opts.quoted)
	b = strconv.AppendBool(b, v.Bool())
	b = mayAppendQuote(b, opts.quoted)
	e.Buffer.Write(b)
}

func intEncoder(e *encodeState, v reflect.Value, opts encOpts) {
	b := e.AvailableBuffer()
	b = mayAppendQuote(b, opts.quoted)
	b = strconv.AppendInt(b, v.Int(), 10)
	b = mayAppendQuote(b, opts.quoted)
	e.Buffer.Write(b)
}

func uintEncoder(e *encodeState, v reflect.Value, opts encOpts) {
	b := e.AvailableBuffer()
	b = mayAppendQuote(b, opts.quoted)
	b = strconv.AppendUint(b, v.Uint(), 10)
	b = mayAppendQuote(b, opts.quoted)
	e.Buffer.Write(b)
}

type floatEncoder int

func (bits floatEncoder) encode(e *encodeState, v reflect.Value, opts encOpts) {
	f := v.Float()
	if math.IsInf(f, 0) || math.IsNaN(f) {
		e.error(&UnsupportedValueError{v, strconv.FormatFloat(f, 'g', -1, int(bits))})
	}

	b := e.AvailableBuffer()
	b = mayAppendQuote(b, opts.quoted)
	abs := math.Abs(f)
	fmt := byte('f')
	if abs != 0 {
		if bits == 64 && (abs < 1e-6 || abs >= 1e21) || bits == 32 && (float32(abs) < 1e-6 || float32(abs) >= 1e21) {
			fmt = 'e'
		}
	}
	b = strconv.AppendFloat(b, f, fmt, -1, int(bits))
	if fmt == 'e' {
		n := len(b)
		if n >= 4 && b[n-4] == 'e' && b[n-3] == '-' && b[n-2] == '0' {
			b[n-2] = b[n-1]
			b = b[:n-1]
		}
	}
	b = mayAppendQuote(b, opts.quoted)
	e.Buffer.Write(b)
}

var (
	float32Encoder = (floatEncoder(32)).encode
	float64Encoder = (floatEncoder(64)).encode
)

func stringEncoder(e *encodeState, v reflect.Value, opts encOpts) {
	if opts.quoted {
		b := appendString(nil, v.String(), true)
		e.Buffer.Write(appendString(e.AvailableBuffer(), b, false))
	} else {
		e.Buffer.Write(appendString(e.AvailableBuffer(), v.String(), true))
	}
}

func interfaceEncoder(e *encodeState, v reflect.Value, opts encOpts) {
	if v.IsNil() {
		e.Buffer.WriteString("null")
		return
	}
	e.reflectValue(v.Elem(), opts)
}

func unsupportedTypeEncoder(e *encodeState, v reflect.Value, _ encOpts) {
	e.error(&UnsupportedTypeError{v.Type()})
}

type structEncoder struct {
	fields structFields
}

type structFields struct {
	list []field
}

func (se structEncoder) encode(e *encodeState, v reflect.Value, opts encOpts) {
	next := byte('{')
FieldLoop:
	for i := range se.fields.list {
		f := &se.fields.list[i]

		fv := v
		for _, i := range f.index {
			if fv.Kind() == reflect.Pointer {
				if fv.IsNil() {
					continue FieldLoop
				}
				fv = fv.Elem()
			}
			fv = fv.Field(i)
		}

		if (f.omitEmpty && isEmptyValue(fv)) ||
			(f.omitZero && (f.isZero == nil && fv.IsZero() || (f.isZero != nil && f.isZero(fv)))) {
			continue
		}
		e.Buffer.WriteByte(next)
		next = ','
		e.Buffer.WriteString(f.nameEscHTML)
		opts.quoted = f.quoted
		f.encoder(e, fv, opts)
	}
	if next == '{' {
		e.Buffer.WriteString("{}")
	} else {
		e.Buffer.WriteByte('}')
	}
}

func newStructEncoder(t reflect.Type) encoderFunc {
	se := structEncoder{fields: cachedTypeFields(t)}
	return se.encode
}

type mapEncoder struct {
	elemEnc encoderFunc
}

func (me mapEncoder) encode(e *encodeState, v reflect.Value, opts encOpts) {
	if v.IsNil() || sentinel.IsValueNull(v) {
		e.Buffer.WriteString("null")
		return
	}
	if e.ptrLevel++; e.ptrLevel > startDetectingCyclesAfter {
		ptr := cycleKey{v.Type(), v.Pointer(), 0}
		if _, ok := e.ptrSeen[ptr]; ok {
			e.error(&UnsupportedValueError{v, fmt.Sprintf("encountered a cycle via %s", v.Type())})
		}
		e.ptrSeen[ptr] = struct{}{}
		defer delete(e.ptrSeen, ptr)
	}
	e.Buffer.WriteByte('{')

	var (
		sv  = make([]reflectWithString, v.Len())
		mi  = v.MapRange()
		err error
	)
	for i := 0; mi.Next(); i++ {
		if sv[i].ks, err = resolveKeyName(mi.Key()); err != nil {
			e.error(fmt.Errorf("json: encoding error for type %q: %q", v.Type().String(), err.Error()))
		}
		sv[i].v = mi.Value()
	}
	slices.SortFunc(sv, func(i, j reflectWithString) int {
		return strings.Compare(i.ks, j.ks)
	})

	for i, kv := range sv {
		if i > 0 {
			e.Buffer.WriteByte(',')
		}
		e.Buffer.Write(appendString(e.AvailableBuffer(), kv.ks, true))
		e.Buffer.WriteByte(':')
		me.elemEnc(e, kv.v, opts)
	}
	e.Buffer.WriteByte('}')
	e.ptrLevel--
}

func newMapEncoder(t reflect.Type) encoderFunc {
	switch t.Key().Kind() {
	case reflect.String,
		reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
	default:
		if !t.Key().Implements(textMarshalerType) {
			return unsupportedTypeEncoder
		}
	}
	me := mapEncoder{typeEncoder(t.Elem())}
	return me.encode
}

func encodeByteSlice(e *encodeState, v reflect.Value, _ encOpts) {
	if v.IsNil() {
		e.Buffer.WriteString("null")
		return
	}

	s := v.Bytes()
	b := e.AvailableBuffer()
	b = append(b, '"')
	b = base64.StdEncoding.AppendEncode(b, s)
	b = append(b, '"')
	e.Buffer.Write(b)
}

type sliceEncoder struct {
	arrayEnc encoderFunc
}

func (se sliceEncoder) encode(e *encodeState, v reflect.Value, opts encOpts) {
	if v.IsNil() || sentinel.IsValueNull(v) {
		e.Buffer.WriteString("null")
		return
	}
	if e.ptrLevel++; e.ptrLevel > startDetectingCyclesAfter {
		ptr := cycleKey{v.Type(), v.Pointer(), v.Len()}
		if _, ok := e.ptrSeen[ptr]; ok {
			e.error(&UnsupportedValueError{v, fmt.Sprintf("encountered a cycle via %s", v.Type())})
		}
		e.ptrSeen[ptr] = struct{}{}
		defer delete(e.ptrSeen, ptr)
	}
	se.arrayEnc(e, v, opts)
	e.ptrLevel--
}

func newSliceEncoder(t reflect.Type) encoderFunc {
	if t.Elem().Kind() == reflect.Uint8 {
		p := reflect.PointerTo(t.Elem())
		if !p.Implements(marshalerType) && !p.Implements(textMarshalerType) {
			return encodeByteSlice
		}
	}
	enc := sliceEncoder{newArrayEncoder(t)}
	return enc.encode
}

type arrayEncoder struct {
	elemEnc encoderFunc
}

func (ae arrayEncoder) encode(e *encodeState, v reflect.Value, opts encOpts) {
	e.Buffer.WriteByte('[')
	n := v.Len()
	for i := 0; i < n; i++ {
		if i > 0 {
			e.Buffer.WriteByte(',')
		}
		ae.elemEnc(e, v.Index(i), opts)
	}
	e.Buffer.WriteByte(']')
}

func newArrayEncoder(t reflect.Type) encoderFunc {
	enc := arrayEncoder{typeEncoder(t.Elem())}
	return enc.encode
}

type ptrEncoder struct {
	elemEnc encoderFunc
}

func (pe ptrEncoder) encode(e *encodeState, v reflect.Value, opts encOpts) {
	if v.IsNil() {
		e.Buffer.WriteString("null")
		return
	}

	if e.ptrLevel++; e.ptrLevel > startDetectingCyclesAfter {
		ptr := cycleKey{v.Type(), v.Pointer(), 0}
		if _, ok := e.ptrSeen[ptr]; ok {
			e.error(&UnsupportedValueError{v, fmt.Sprintf("encountered a cycle via %s", v.Type())})
		}
		e.ptrSeen[ptr] = struct{}{}
		defer delete(e.ptrSeen, ptr)
	}
	pe.elemEnc(e, v.Elem(), opts)
	e.ptrLevel--
}

func newPtrEncoder(t reflect.Type) encoderFunc {
	enc := ptrEncoder{typeEncoder(t.Elem())}
	return enc.encode
}

type condAddrEncoder struct {
	canAddrEnc, elseEnc encoderFunc
}

func (ce condAddrEncoder) encode(e *encodeState, v reflect.Value, opts encOpts) {
	if v.CanAddr() {
		ce.canAddrEnc(e, v, opts)
	} else {
		ce.elseEnc(e, v, opts)
	}
}

func newCondAddrEncoder(canAddrEnc, elseEnc encoderFunc) encoderFunc {
	enc := condAddrEncoder{canAddrEnc: canAddrEnc, elseEnc: elseEnc}
	return enc.encode
}

func isValidTag(s string) bool {
	if s == "" {
		return false
	}
	for _, c := range s {
		switch {
		case strings.ContainsRune("!#$%&()*+-./:;<=>?@[]^_{|}~ ", c):
		case !unicode.IsLetter(c) && !unicode.IsDigit(c):
			return false
		}
	}
	return true
}

func typeByIndex(t reflect.Type, index []int) reflect.Type {
	for _, i := range index {
		if t.Kind() == reflect.Pointer {
			t = t.Elem()
		}
		t = t.Field(i).Type
	}
	return t
}

type reflectWithString struct {
	v  reflect.Value
	ks string
}

func resolveKeyName(k reflect.Value) (string, error) {
	if k.Kind() == reflect.String {
		return k.String(), nil
	}
	if tm, ok := k.Interface().(encoding.TextMarshaler); ok {
		if k.Kind() == reflect.Pointer && k.IsNil() {
			return "", nil
		}
		buf, err := tm.MarshalText()
		return string(buf), err
	}
	switch k.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(k.Int(), 10), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return strconv.FormatUint(k.Uint(), 10), nil
	}
	panic("unexpected map key type")
}

func appendString[Bytes []byte | string](dst []byte, src Bytes, escapeHTML bool) []byte {
	dst = append(dst, '"')
	start := 0
	for i := 0; i < len(src); {
		if b := src[i]; b < utf8.RuneSelf {
			if htmlSafeSet[b] || (!escapeHTML && safeSet[b]) {
				i++
				continue
			}
			dst = append(dst, src[start:i]...)
			switch b {
			case '\\', '"':
				dst = append(dst, '\\', b)
			case '\b':
				dst = append(dst, '\\', 'b')
			case '\f':
				dst = append(dst, '\\', 'f')
			case '\n':
				dst = append(dst, '\\', 'n')
			case '\r':
				dst = append(dst, '\\', 'r')
			case '\t':
				dst = append(dst, '\\', 't')
			default:
				dst = append(dst, '\\', 'u', '0', '0', hex[b>>4], hex[b&0xF])
			}
			i++
			start = i
			continue
		}
		// TODO(https://go.dev/issue/56948): Use generic utf8 functionality.
		n := len(src) - i
		if n > utf8.UTFMax {
			n = utf8.UTFMax
		}
		c, size := utf8.DecodeRuneInString(string(src[i : i+n]))
		if c == utf8.RuneError && size == 1 {
			dst = append(dst, src[start:i]...)
			dst = append(dst, `\ufffd`...)
			i += size
			start = i
			continue
		}
		// U+2028 is LINE SEPARATOR.
		// U+2029 is PARAGRAPH SEPARATOR.
		// They are both technically valid characters in JSON strings,
		// but don't work in JSONP, which has to be evaluated as JavaScript,
		// and can lead to security holes there. It is valid JSON to
		// escape them, so we do so unconditionally.
		// See https://en.wikipedia.org/wiki/JSON#Safety.
		if c == '\u2028' || c == '\u2029' {
			dst = append(dst, src[start:i]...)
			dst = append(dst, '\\', 'u', '2', '0', '2', hex[c&0xF])
			i += size
			start = i
			continue
		}
		i += size
	}
	dst = append(dst, src[start:]...)
	dst = append(dst, '"')
	return dst
}

type field struct {
	name string

	nameEscHTML string

	tag       bool
	index     []int
	typ       reflect.Type
	omitEmpty bool
	omitZero  bool
	isZero    func(reflect.Value) bool
	quoted    bool

	encoder encoderFunc

	timefmt string
}

type isZeroer interface {
	IsZero() bool
}

var isZeroerType = reflect.TypeFor[isZeroer]()

func callIsZero(v reflect.Value) bool {
	z, err := reflectvalue.As[isZeroer](v)
	return err == nil && z.IsZero()
}

func typeFields(t reflect.Type) structFields {
	current := []field{}
	next := []field{{typ: t}}

	var count, nextCount map[reflect.Type]int

	visited := map[reflect.Type]bool{}

	var fields []field

	var nameEscBuf []byte

	for len(next) > 0 {
		current, next = next, current[:0]
		count, nextCount = nextCount, map[reflect.Type]int{}

		for _, f := range current {
			if visited[f.typ] {
				continue
			}
			visited[f.typ] = true

			for i := 0; i < f.typ.NumField(); i++ {
				sf := f.typ.Field(i)
				if sf.Anonymous {
					t := sf.Type
					if t.Kind() == reflect.Pointer {
						t = t.Elem()
					}
					if !sf.IsExported() && t.Kind() != reflect.Struct {
						continue
					}
				} else if !sf.IsExported() {
					continue
				}
				tag := sf.Tag.Get("json")
				if tag == "-" {
					continue
				}
				name, opts := parseTag(tag)
				if !isValidTag(name) {
					name = ""
				}
				index := make([]int, len(f.index)+1)
				copy(index, f.index)
				index[len(f.index)] = i

				ft := sf.Type
				if ft.Name() == "" && ft.Kind() == reflect.Pointer {
					ft = ft.Elem()
				}

				quoted := false
				if opts.Contains("string") {
					switch ft.Kind() {
					case reflect.Bool,
						reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
						reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr,
						reflect.Float32, reflect.Float64,
						reflect.String:
						quoted = true
					}
				}

				if name != "" || !sf.Anonymous || ft.Kind() != reflect.Struct {
					tagged := name != ""
					if name == "" {
						name = sf.Name
					}
					field := field{
						name:      name,
						tag:       tagged,
						index:     index,
						typ:       ft,
						omitEmpty: opts.Contains("omitempty"),
						omitZero:  opts.Contains("omitzero"),
						quoted:    quoted,
						timefmt:   sf.Tag.Get("format"),
					}
					nameEscBuf = appendHTMLEscape(nameEscBuf[:0], []byte(field.name))
					field.nameEscHTML = `"` + string(nameEscBuf) + `":`

					if field.omitZero {
						t := sf.Type
						switch {
						case t.Kind() == reflect.Interface && t.Implements(isZeroerType):
							field.isZero = func(v reflect.Value) bool {
								return v.IsNil() ||
									(v.Elem().Kind() == reflect.Pointer && v.Elem().IsNil()) ||
									callIsZero(v)
							}
						case t.Kind() == reflect.Pointer && t.Implements(isZeroerType):
							field.isZero = func(v reflect.Value) bool {
								return v.IsNil() || callIsZero(v)
							}
						case t.Implements(isZeroerType):
							field.isZero = func(v reflect.Value) bool {
								return callIsZero(v)
							}
						case reflect.PointerTo(t).Implements(isZeroerType):
							field.isZero = func(v reflect.Value) bool {
								if !v.CanAddr() {
									v2 := reflect.New(v.Type()).Elem()
									v2.Set(v)
									v = v2
								}
								return callIsZero(v.Addr())
							}
						}
					}

					fields = append(fields, field)
					if count[f.typ] > 1 {
						fields = append(fields, fields[len(fields)-1])
					}
					continue
				}

				nextCount[ft]++
				if nextCount[ft] == 1 {
					next = append(next, field{name: ft.Name(), index: index, typ: ft})
				}
			}
		}
	}

	slices.SortFunc(fields, func(a, b field) int {
		if c := strings.Compare(a.name, b.name); c != 0 {
			return c
		}
		if c := cmp.Compare(len(a.index), len(b.index)); c != 0 {
			return c
		}
		if a.tag != b.tag {
			if a.tag {
				return -1
			}
			return +1
		}
		return slices.Compare(a.index, b.index)
	})

	out := fields[:0]
	for advance, i := 0, 0; i < len(fields); i += advance {
		fi := fields[i]
		name := fi.name
		for advance = 1; i+advance < len(fields); advance++ {
			fj := fields[i+advance]
			if fj.name != name {
				break
			}
		}
		if advance == 1 {
			out = append(out, fi)
			continue
		}
		dominant, ok := dominantField(fields[i : i+advance])
		if ok {
			out = append(out, dominant)
		}
	}

	fields = out
	slices.SortFunc(fields, func(i, j field) int {
		return slices.Compare(i.index, j.index)
	})

	for i := range fields {
		f := &fields[i]
		f.encoder = typeEncoder(typeByIndex(t, f.index))

		if f.timefmt != "" {
			f.encoder = continueWithTimeFmt(f.timefmt, f.encoder)
		}
	}
	return structFields{fields}
}

func dominantField(fields []field) (field, bool) {
	if len(fields) > 1 && len(fields[0].index) == len(fields[1].index) && fields[0].tag == fields[1].tag {
		return field{}, false
	}
	return fields[0], true
}

var fieldCache = syncMap[reflect.Type, structFields]{m: map[reflect.Type]structFields{}}

func cachedTypeFields(t reflect.Type) structFields {
	if f, ok := fieldCache.load(t); ok {
		return f
	}
	f, _ := fieldCache.loadOrStore(t, typeFields(t))
	return f
}

type syncMap[K comparable, V any] struct {
	mu sync.RWMutex
	m  map[K]V
}

func (s *syncMap[K, V]) load(k K) (V, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	v, ok := s.m[k]
	return v, ok
}

func (s *syncMap[K, V]) loadOrStore(k K, v V) (V, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if actual, ok := s.m[k]; ok {
		return actual, true
	}
	s.m[k] = v
	return v, false
}

func (s *syncMap[K, V]) store(k K, v V) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.m[k] = v
}

func mayAppendQuote(b []byte, quoted bool) []byte {
	if quoted {
		b = append(b, '"')
	}
	return b
}
