// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

package apiquery

import (
	"encoding/json"
	"fmt"
	"github.com/Xquik-dev/x-twitter-scraper-go/internal/paramutil"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/Xquik-dev/x-twitter-scraper-go/packages/param"
)

var encoders paramutil.SyncMap[encoderEntry, encoderFunc]

type encoder struct {
	dateFormat string
	root       bool
	settings   QuerySettings
}

type encoderFunc func(key string, value reflect.Value) ([]Pair, error)

type encoderField struct {
	tag parsedStructTag
	fn  encoderFunc
	idx []int
}

type encoderEntry struct {
	typ        reflect.Type
	dateFormat string
	root       bool
	settings   QuerySettings
}

type Pair struct {
	key   string
	value string
}

func (e *encoder) typeEncoder(t reflect.Type) encoderFunc {
	entry := encoderEntry{
		typ:        t,
		dateFormat: e.dateFormat,
		root:       e.root,
		settings:   e.settings,
	}

	if fi, ok := encoders.Load(entry); ok {
		return fi
	}

	var (
		wg sync.WaitGroup
		f  encoderFunc
	)
	wg.Add(1)
	fi, loaded := encoders.LoadOrStore(entry, encoderFunc(func(key string, v reflect.Value) ([]Pair, error) {
		wg.Wait()
		return f(key, v)
	}))
	if loaded {
		return fi
	}

	f = e.newTypeEncoder(t)
	wg.Done()
	encoders.Store(entry, f)
	return f
}

func marshalerEncoder(key string, value reflect.Value) ([]Pair, error) {
	marshaler, ok := reflect.TypeAssert[json.Marshaler](value)
	if !ok {
		return nil, fmt.Errorf("apiquery: %s does not implement json.Marshaler", value.Type())
	}
	s, err := marshaler.MarshalJSON()
	if err != nil {
		return nil, fmt.Errorf("apiquery: json fallback marshal error %s", err)
	}
	return []Pair{{key, string(s)}}, nil
}

func (e *encoder) newTypeEncoder(t reflect.Type) encoderFunc {
	if t.ConvertibleTo(reflect.TypeOf(time.Time{})) {
		return e.newTimeTypeEncoder(t)
	}

	if t.Implements(reflect.TypeOf((*param.Optional)(nil)).Elem()) {
		return e.newRichFieldTypeEncoder(t)
	}

	if !e.root && t.Implements(reflect.TypeOf((*json.Marshaler)(nil)).Elem()) {
		return marshalerEncoder
	}

	e.root = false
	switch t.Kind() {
	case reflect.Pointer:
		encoder := e.typeEncoder(t.Elem())
		return func(key string, value reflect.Value) (pairs []Pair, err error) {
			if !value.IsValid() || value.IsNil() {
				return pairs, err
			}
			return encoder(key, value.Elem())
		}
	case reflect.Struct:
		return e.newStructTypeEncoder(t)
	case reflect.Array:
		fallthrough
	case reflect.Slice:
		return e.newArrayTypeEncoder(t)
	case reflect.Map:
		return e.newMapEncoder(t)
	case reflect.Interface:
		return e.newInterfaceEncoder()
	default:
		return e.newPrimitiveTypeEncoder(t)
	}
}

func (e *encoder) newStructTypeEncoder(t reflect.Type) encoderFunc {
	if t.Implements(reflect.TypeOf((*param.Optional)(nil)).Elem()) {
		return e.newRichFieldTypeEncoder(t)
	}

	for i := 0; i < t.NumField(); i++ {
		if t.Field(i).Type == paramUnionType && t.Field(i).Anonymous {
			return e.newStructUnionTypeEncoder(t)
		}
	}

	encoderFields := []encoderField{}

	var collectEncoderFields func(r reflect.Type, index []int)
	collectEncoderFields = func(r reflect.Type, index []int) {
		for i := 0; i < r.NumField(); i++ {
			idx := append(index, i)
			field := t.FieldByIndex(idx)
			if !field.IsExported() {
				continue
			}
			if field.Anonymous {
				collectEncoderFields(field.Type, idx)
				continue
			}
			ptag, ok := parseQueryStructTag(field)
			if !ok {
				continue
			}

			if (ptag.name == "-" || ptag.name == "") && !ptag.inline {
				continue
			}

			dateFormat, ok := parseFormatStructTag(field)
			oldFormat := e.dateFormat
			if ok {
				switch dateFormat {
				case "date-time":
					e.dateFormat = time.RFC3339
				case "date":
					e.dateFormat = "2006-01-02"
				}
			}
			var encoderFn encoderFunc
			if ptag.omitzero {
				typeEncoderFn := e.typeEncoder(field.Type)
				encoderFn = func(key string, value reflect.Value) ([]Pair, error) {
					if value.IsZero() {
						return nil, nil
					}
					return typeEncoderFn(key, value)
				}
			} else {
				encoderFn = e.typeEncoder(field.Type)
			}
			encoderFields = append(encoderFields, encoderField{ptag, encoderFn, idx})
			e.dateFormat = oldFormat
		}
	}
	collectEncoderFields(t, []int{})

	return func(key string, value reflect.Value) (pairs []Pair, err error) {
		for _, ef := range encoderFields {
			subkey := e.renderKeyPath(key, ef.tag.name)
			if ef.tag.inline {
				subkey = key
			}

			field := value.FieldByIndex(ef.idx)
			subpairs, suberr := ef.fn(subkey, field)
			if suberr != nil {
				err = suberr
			}
			pairs = append(pairs, subpairs...)
		}
		return pairs, err
	}
}

var paramUnionType = reflect.TypeOf((*param.APIUnion)(nil)).Elem()

func (e *encoder) newStructUnionTypeEncoder(t reflect.Type) encoderFunc {
	var fieldEncoders []encoderFunc
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Type == paramUnionType && field.Anonymous {
			fieldEncoders = append(fieldEncoders, nil)
			continue
		}
		fieldEncoders = append(fieldEncoders, e.typeEncoder(field.Type))
	}

	return func(key string, value reflect.Value) (pairs []Pair, err error) {
		for i := 0; i < t.NumField(); i++ {
			if value.Field(i).Type() == paramUnionType {
				continue
			}
			if !value.Field(i).IsZero() {
				return fieldEncoders[i](key, value.Field(i))
			}
		}
		return nil, fmt.Errorf("apiquery: union %s has no field set", t.String())
	}
}

func (e *encoder) newMapEncoder(t reflect.Type) encoderFunc {
	keyEncoder := e.typeEncoder(t.Key())
	elementEncoder := e.typeEncoder(t.Elem())
	return func(key string, value reflect.Value) (pairs []Pair, err error) {
		iter := value.MapRange()
		for iter.Next() {
			encodedKey, err := keyEncoder("", iter.Key())
			if err != nil {
				return nil, err
			}
			if len(encodedKey) != 1 {
				return nil, fmt.Errorf("apiquery: unexpected number of parts for encoded map key, map may contain non-primitive")
			}
			subkey := encodedKey[0].value
			keyPath := e.renderKeyPath(key, subkey)
			subpairs, suberr := elementEncoder(keyPath, iter.Value())
			if suberr != nil {
				err = suberr
			}
			pairs = append(pairs, subpairs...)
		}
		return pairs, err
	}
}

func (e *encoder) renderKeyPath(key string, subkey string) string {
	if len(key) == 0 {
		return subkey
	}
	if e.settings.NestedFormat == NestedQueryFormatDots {
		return fmt.Sprintf("%s.%s", key, subkey)
	}
	return fmt.Sprintf("%s[%s]", key, subkey)
}

func (e *encoder) newArrayTypeEncoder(t reflect.Type) encoderFunc {
	switch e.settings.ArrayFormat {
	case ArrayQueryFormatComma:
		innerEncoder := e.typeEncoder(t.Elem())
		return func(key string, v reflect.Value) ([]Pair, error) {
			elements := []string{}
			for i := 0; i < v.Len(); i++ {
				innerPairs, err := innerEncoder("", v.Index(i))
				if err != nil {
					return nil, err
				}
				for _, pair := range innerPairs {
					elements = append(elements, pair.value)
				}
			}
			if len(elements) == 0 {
				return []Pair{}, nil
			}
			return []Pair{{key, strings.Join(elements, ",")}}, nil
		}
	case ArrayQueryFormatRepeat:
		innerEncoder := e.typeEncoder(t.Elem())
		return func(key string, value reflect.Value) (pairs []Pair, err error) {
			for i := 0; i < value.Len(); i++ {
				subpairs, suberr := innerEncoder(key, value.Index(i))
				if suberr != nil {
					err = suberr
				}
				pairs = append(pairs, subpairs...)
			}
			return pairs, err
		}
	case ArrayQueryFormatIndices:
		panic("The array indices format is not supported yet")
	case ArrayQueryFormatBrackets:
		innerEncoder := e.typeEncoder(t.Elem())
		return func(key string, value reflect.Value) (pairs []Pair, err error) {
			pairs = []Pair{}
			for i := 0; i < value.Len(); i++ {
				subpairs, suberr := innerEncoder(key+"[]", value.Index(i))
				if suberr != nil {
					err = suberr
				}
				pairs = append(pairs, subpairs...)
			}
			return pairs, err
		}
	default:
		panic(fmt.Sprintf("Unknown ArrayFormat value: %d", e.settings.ArrayFormat))
	}
}

func (e *encoder) newPrimitiveTypeEncoder(t reflect.Type) encoderFunc {
	switch t.Kind() {
	case reflect.Pointer:
		inner := t.Elem()

		innerEncoder := e.newPrimitiveTypeEncoder(inner)
		return func(key string, v reflect.Value) ([]Pair, error) {
			if !v.IsValid() || v.IsNil() {
				return nil, nil
			}
			return innerEncoder(key, v.Elem())
		}
	case reflect.String:
		return func(key string, v reflect.Value) ([]Pair, error) {
			return []Pair{{key, v.String()}}, nil
		}
	case reflect.Bool:
		return func(key string, v reflect.Value) ([]Pair, error) {
			return []Pair{{key, strconv.FormatBool(v.Bool())}}, nil
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return func(key string, v reflect.Value) ([]Pair, error) {
			return []Pair{{key, strconv.FormatInt(v.Int(), 10)}}, nil
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return func(key string, v reflect.Value) ([]Pair, error) {
			return []Pair{{key, strconv.FormatUint(v.Uint(), 10)}}, nil
		}
	case reflect.Float32, reflect.Float64:
		return func(key string, v reflect.Value) ([]Pair, error) {
			return []Pair{{key, strconv.FormatFloat(v.Float(), 'f', -1, 64)}}, nil
		}
	case reflect.Complex64, reflect.Complex128:
		bitSize := t.Bits()
		return func(key string, v reflect.Value) ([]Pair, error) {
			return []Pair{{key, strconv.FormatComplex(v.Complex(), 'f', -1, bitSize)}}, nil
		}
	default:
		return func(key string, v reflect.Value) ([]Pair, error) {
			return nil, nil
		}
	}
}

func (e *encoder) newTimeTypeEncoder(_ reflect.Type) encoderFunc {
	format := e.dateFormat
	return func(key string, value reflect.Value) ([]Pair, error) {
		return []Pair{{
			key,
			formatTime(value, format),
		}}, nil
	}
}

func formatTime(value reflect.Value, format string) string {
	var t time.Time
	reflect.ValueOf(&t).Elem().Set(value.Convert(reflect.TypeFor[time.Time]()))
	return t.Format(format)
}

func (e encoder) newInterfaceEncoder() encoderFunc {
	return func(key string, value reflect.Value) ([]Pair, error) {
		value = value.Elem()
		if !value.IsValid() {
			return nil, nil
		}
		return e.typeEncoder(value.Type())(key, value)
	}

}
