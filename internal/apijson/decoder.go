// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

// Package apijson decodes API responses into SDK types and encodes params as
// JSON. It records whether each field was present, null or invalid.
package apijson

import (
	"encoding/json"
	"fmt"
	"github.com/Xquik-dev/x-twitter-scraper-go/internal/paramutil"
	"github.com/Xquik-dev/x-twitter-scraper-go/packages/param"
	"reflect"
	"strconv"
	"sync"
	"time"

	"github.com/tidwall/gjson"
)

var decoders paramutil.SyncMap[decoderEntry, decoderFunc]

func Unmarshal[T any](raw []byte, to *T) error {
	d := &decoderBuilder{dateFormat: time.RFC3339}
	return d.unmarshal(gjson.ParseBytes(raw), reflect.ValueOf(to).Elem())
}

func UnmarshalRoot[T any](raw []byte, to *T, rawJSON *string) error {
	d := &decoderBuilder{dateFormat: time.RFC3339, root: true}
	result := gjson.ParseBytes(raw)
	if rawJSON != nil {
		*rawJSON = result.Raw
	}
	return d.unmarshal(result, reflect.ValueOf(to).Elem())
}

type decoderBuilder struct {
	root       bool
	dateFormat string
}

type decoderState struct {
	strict    bool
	exactness exactness
	validator *validationEntry
}

type exactness int8

const (
	loose exactness = iota
	extras
	exact
)

type decoderFunc func(node gjson.Result, value reflect.Value, state *decoderState) error

type decoderField struct {
	tag    parsedStructTag
	fn     decoderFunc
	idx    []int
	goname string
}

type decoderEntry struct {
	typ        reflect.Type
	dateFormat string
	root       bool
}

func (d *decoderBuilder) unmarshal(result gjson.Result, value reflect.Value) error {
	if !value.IsValid() {
		return fmt.Errorf("apijson: cannot marshal into invalid value")
	}
	return d.typeDecoder(value.Type())(result, value, &decoderState{strict: false, exactness: exact})
}

func (d *decoderBuilder) unmarshalWithExactness(raw []byte, value reflect.Value) (exactness, error) {
	result := gjson.ParseBytes(raw)
	if !value.IsValid() {
		return 0, fmt.Errorf("apijson: cannot marshal into invalid value")
	}
	state := decoderState{strict: false, exactness: exact}
	err := d.typeDecoder(value.Type())(result, value, &state)
	return state.exactness, err
}

func (d *decoderBuilder) typeDecoder(t reflect.Type) decoderFunc {
	entry := decoderEntry{
		typ:        t,
		dateFormat: d.dateFormat,
		root:       d.root,
	}

	if fi, ok := decoders.Load(entry); ok {
		return fi
	}

	var (
		wg sync.WaitGroup
		f  decoderFunc
	)
	wg.Add(1)
	fi, loaded := decoders.LoadOrStore(entry, decoderFunc(func(node gjson.Result, v reflect.Value, state *decoderState) error {
		wg.Wait()
		return f(node, v, state)
	}))
	if loaded {
		return fi
	}

	f = d.newTypeDecoder(t)
	wg.Done()
	decoders.Store(entry, f)
	return f
}

func (d *decoderBuilder) validatedTypeDecoder(t reflect.Type, entry *validationEntry) decoderFunc {
	dec := d.typeDecoder(t)
	if entry == nil {
		return dec
	}

	return func(node gjson.Result, v reflect.Value, state *decoderState) error {
		state.validator = entry
		err := dec(node, v, state)
		state.validator = nil
		return err
	}
}

func indirectUnmarshalerDecoder(n gjson.Result, v reflect.Value, state *decoderState) error {
	return unmarshalRaw(v.Addr(), n.Raw)
}

func unmarshalerDecoder(n gjson.Result, v reflect.Value, state *decoderState) error {
	if v.Kind() == reflect.Pointer && v.CanSet() {
		v.Set(reflect.New(v.Type().Elem()))
	}
	return unmarshalRaw(v, n.Raw)
}

func unmarshalRaw(v reflect.Value, raw string) error {
	unmarshaler, ok := reflect.TypeAssert[json.Unmarshaler](v)
	if !ok {
		return fmt.Errorf("apijson: %s does not implement json.Unmarshaler", v.Type())
	}
	return unmarshaler.UnmarshalJSON([]byte(raw))
}

func (d *decoderBuilder) newTypeDecoder(t reflect.Type) decoderFunc {
	if t.ConvertibleTo(reflect.TypeOf(time.Time{})) {
		return d.newTimeTypeDecoder(t)
	}

	if t.Implements(reflect.TypeOf((*param.Optional)(nil)).Elem()) {
		return d.newOptTypeDecoder(t)
	}

	if !d.root && t.Implements(reflect.TypeOf((*json.Unmarshaler)(nil)).Elem()) {
		return unmarshalerDecoder
	}
	if !d.root && reflect.PointerTo(t).Implements(reflect.TypeOf((*json.Unmarshaler)(nil)).Elem()) {
		if _, ok := unionVariants[t]; !ok {
			return indirectUnmarshalerDecoder
		}
	}
	d.root = false

	if _, ok := unionRegistry[t]; ok {
		if isStructUnion(t) {
			return d.newStructUnionDecoder(t)
		}
		return d.newUnionDecoder(t)
	}

	switch t.Kind() {
	case reflect.Pointer:
		inner := t.Elem()
		innerDecoder := d.typeDecoder(inner)

		return func(n gjson.Result, v reflect.Value, state *decoderState) error {
			if !v.IsValid() {
				return fmt.Errorf("apijson: unexpected invalid reflection value %+#v", v)
			}

			newValue := reflect.New(inner).Elem()
			err := innerDecoder(n, newValue, state)
			if err != nil {
				return err
			}

			v.Set(newValue.Addr())
			return nil
		}
	case reflect.Struct:
		if isStructUnion(t) {
			return d.newStructUnionDecoder(t)
		}
		return d.newStructTypeDecoder(t)
	case reflect.Array:
		fallthrough
	case reflect.Slice:
		return d.newArrayTypeDecoder(t)
	case reflect.Map:
		return d.newMapDecoder(t)
	case reflect.Interface:
		return func(node gjson.Result, value reflect.Value, state *decoderState) error {
			if !value.IsValid() {
				return fmt.Errorf("apijson: unexpected invalid value %+#v", value)
			}
			if node.Value() != nil && value.CanSet() {
				value.Set(reflect.ValueOf(node.Value()))
			}
			return nil
		}
	default:
		return d.newPrimitiveTypeDecoder(t)
	}
}

func (d *decoderBuilder) newMapDecoder(t reflect.Type) decoderFunc {
	keyType := t.Key()
	itemType := t.Elem()
	itemDecoder := d.typeDecoder(itemType)

	return func(node gjson.Result, value reflect.Value, state *decoderState) (err error) {
		mapValue := reflect.MakeMapWithSize(t, len(node.Map()))

		node.ForEach(func(key, value gjson.Result) bool {
			keyValue := reflect.ValueOf(key.Value())
			if !keyValue.IsValid() {
				if err == nil {
					err = fmt.Errorf("apijson: received invalid key type %v", keyValue.String())
				}
				return false
			}
			if keyValue.Type() != keyType {
				if err == nil {
					err = fmt.Errorf("apijson: expected key type %v but got %v", keyType, keyValue.Type())
				}
				return false
			}

			itemValue := reflect.New(itemType).Elem()
			itemerr := itemDecoder(value, itemValue, state)
			if itemerr != nil {
				if err == nil {
					err = itemerr
				}
				return false
			}

			mapValue.SetMapIndex(keyValue, itemValue)
			return true
		})

		if err != nil {
			return err
		}
		value.Set(mapValue)
		return nil
	}
}

func (d *decoderBuilder) newArrayTypeDecoder(t reflect.Type) decoderFunc {
	itemDecoder := d.typeDecoder(t.Elem())

	return func(node gjson.Result, value reflect.Value, state *decoderState) (err error) {
		if !node.IsArray() {
			return fmt.Errorf("apijson: could not deserialize to an array")
		}

		arrayNode := node.Array()

		arrayValue := reflect.MakeSlice(reflect.SliceOf(t.Elem()), len(arrayNode), len(arrayNode))
		for i, itemNode := range arrayNode {
			err = itemDecoder(itemNode, arrayValue.Index(i), state)
			if err != nil {
				return err
			}
		}

		value.Set(arrayValue)
		return nil
	}
}

func (d *decoderBuilder) newStructTypeDecoder(t reflect.Type) decoderFunc {
	decoderFields := map[string]decoderField{}
	anonymousDecoders := []decoderField{}
	extraDecoder := (*decoderField)(nil)
	var inlineDecoders []decoderField

	validationEntries := validationRegistry[t]

	for i := 0; i < t.NumField(); i++ {
		idx := []int{i}
		field := t.FieldByIndex(idx)
		if !field.IsExported() {
			continue
		}

		var validator *validationEntry
		for _, entry := range validationEntries {
			if entry.field.Offset == field.Offset {
				validator = &entry
				break
			}
		}

		if field.Anonymous {
			anonymousDecoders = append(anonymousDecoders, decoderField{
				fn:  d.typeDecoder(field.Type),
				idx: idx[:],
			})
			continue
		}
		ptag, ok := parseJSONStructTag(field)
		if !ok {
			continue
		}
		if ptag.extras {
			extraDecoder = &decoderField{ptag, d.typeDecoder(field.Type.Elem()), idx, field.Name}
			continue
		}
		if ptag.inline {
			df := decoderField{ptag, d.typeDecoder(field.Type), idx, field.Name}
			inlineDecoders = append(inlineDecoders, df)
			continue
		}
		if ptag.metadata {
			continue
		}

		oldFormat := d.dateFormat
		dateFormat, ok := parseFormatStructTag(field)
		if ok {
			switch dateFormat {
			case "date-time":
				d.dateFormat = time.RFC3339
			case "date":
				d.dateFormat = "2006-01-02"
			}
		}

		decoderFields[ptag.name] = decoderField{
			ptag,
			d.validatedTypeDecoder(field.Type, validator),
			idx, field.Name,
		}

		d.dateFormat = oldFormat
	}

	return func(node gjson.Result, value reflect.Value, state *decoderState) (err error) {
		for _, decoder := range anonymousDecoders {
			_ = decoder.fn(node, value.FieldByIndex(decoder.idx), state)
		}

		for _, inlineDecoder := range inlineDecoders {
			var meta Field
			dest := value.FieldByIndex(inlineDecoder.idx)
			isValid := false
			if dest.IsValid() && node.Type != gjson.Null {
				inlineState := decoderState{exactness: state.exactness, strict: true}
				err = inlineDecoder.fn(node, dest, &inlineState)
				if err == nil {
					isValid = true
				}
			}

			if node.Type == gjson.Null {
				meta = Field{
					raw:    node.Raw,
					status: null,
				}
			} else if !isValid {
				if dest.IsValid() {
					dest.SetZero()
				}
				continue
			} else if isValid {
				meta = Field{
					raw:    node.Raw,
					status: valid,
				}
			}
			setMetadataSubField(value, inlineDecoder.idx, inlineDecoder.goname, meta)
		}

		typedExtraType := reflect.Type(nil)
		typedExtraFields := reflect.Value{}
		if extraDecoder != nil {
			typedExtraType = value.FieldByIndex(extraDecoder.idx).Type()
			typedExtraFields = reflect.MakeMap(typedExtraType)
		}
		untypedExtraFields := map[string]Field{}

		for fieldName, itemNode := range node.Map() {
			df, explicit := decoderFields[fieldName]
			var (
				dest reflect.Value
				fn   decoderFunc
				meta Field
			)
			if explicit {
				fn = df.fn
				dest = value.FieldByIndex(df.idx)
			}
			if !explicit && extraDecoder != nil {
				dest = reflect.New(typedExtraType.Elem()).Elem()
				fn = extraDecoder.fn
			}

			isValid := false
			if dest.IsValid() && itemNode.Type != gjson.Null {
				err = fn(itemNode, dest, state)
				if err == nil {
					isValid = true
				}
			}

			if itemNode.Type == gjson.Null && dest.IsValid() && dest.Type().Implements(reflect.TypeOf((*param.Optional)(nil)).Elem()) {
				_ = unmarshalRaw(dest.Addr(), itemNode.Raw)
				continue
			}

			if itemNode.Type == gjson.Null {
				meta = Field{
					raw:    itemNode.Raw,
					status: null,
				}
			} else if !isValid {
				meta = Field{
					raw:    itemNode.Raw,
					status: invalid,
				}
			} else if isValid {
				meta = Field{
					raw:    itemNode.Raw,
					status: valid,
				}
			}

			if explicit {
				setMetadataSubField(value, df.idx, df.goname, meta)
			}
			if !explicit {
				untypedExtraFields[fieldName] = meta
			}
			if !explicit && extraDecoder != nil {
				typedExtraFields.SetMapIndex(reflect.ValueOf(fieldName), dest)
			}
		}

		if extraDecoder != nil && typedExtraFields.Len() > 0 {
			value.FieldByIndex(extraDecoder.idx).Set(typedExtraFields)
		}

		if len(untypedExtraFields) > 0 && state.exactness > extras {
			state.exactness = extras
		}

		if len(untypedExtraFields) > 0 {
			setMetadataExtraFields(value, []int{-1}, "ExtraFields", untypedExtraFields)
		}
		return nil
	}
}

func (d *decoderBuilder) newPrimitiveTypeDecoder(t reflect.Type) decoderFunc {
	switch t.Kind() {
	case reflect.String:
		return func(n gjson.Result, v reflect.Value, state *decoderState) error {
			v.SetString(n.String())
			if guardStrict(state, n.Type != gjson.String) {
				return fmt.Errorf("apijson: failed to parse string strictly")
			}
			if n.Type == gjson.JSON {
				return fmt.Errorf("apijson: failed to parse string")
			}

			state.validateString(v)

			if guardUnknown(state, v) {
				return fmt.Errorf("apijson: failed string enum validation")
			}
			return nil
		}
	case reflect.Bool:
		return func(n gjson.Result, v reflect.Value, state *decoderState) error {
			v.SetBool(n.Bool())
			if guardStrict(state, n.Type != gjson.True && n.Type != gjson.False) {
				return fmt.Errorf("apijson: failed to parse bool strictly")
			}
			if n.Type == gjson.String && (n.Raw != "true" && n.Raw != "false") || n.Type == gjson.JSON {
				return fmt.Errorf("apijson: failed to parse bool")
			}

			state.validateBool(v)

			if guardUnknown(state, v) {
				return fmt.Errorf("apijson: failed bool enum validation")
			}
			return nil
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return func(n gjson.Result, v reflect.Value, state *decoderState) error {
			v.SetInt(n.Int())
			if guardStrict(state, n.Type != gjson.Number || n.Num != float64(int(n.Num))) {
				return fmt.Errorf("apijson: failed to parse int strictly")
			}
			if n.Type == gjson.JSON || (n.Type == gjson.String && !canParseAsNumber(n.Str)) {
				return fmt.Errorf("apijson: failed to parse int")
			}

			state.validateInt(v)

			if guardUnknown(state, v) {
				return fmt.Errorf("apijson: failed int enum validation")
			}
			return nil
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return func(n gjson.Result, v reflect.Value, state *decoderState) error {
			v.SetUint(n.Uint())
			if guardStrict(state, n.Type != gjson.Number || n.Num != float64(int(n.Num)) || n.Num < 0) {
				return fmt.Errorf("apijson: failed to parse uint strictly")
			}
			if n.Type == gjson.JSON || (n.Type == gjson.String && !canParseAsNumber(n.Str)) {
				return fmt.Errorf("apijson: failed to parse uint")
			}
			if guardUnknown(state, v) {
				return fmt.Errorf("apijson: failed uint enum validation")
			}
			return nil
		}
	case reflect.Float32, reflect.Float64:
		return func(n gjson.Result, v reflect.Value, state *decoderState) error {
			v.SetFloat(n.Float())
			if guardStrict(state, n.Type != gjson.Number) {
				return fmt.Errorf("apijson: failed to parse float strictly")
			}
			if n.Type == gjson.JSON || (n.Type == gjson.String && !canParseAsNumber(n.Str)) {
				return fmt.Errorf("apijson: failed to parse float")
			}
			if guardUnknown(state, v) {
				return fmt.Errorf("apijson: failed float enum validation")
			}
			return nil
		}
	default:
		return func(node gjson.Result, v reflect.Value, state *decoderState) error {
			return fmt.Errorf("unknown type received at primitive decoder: %s", t.String())
		}
	}
}

func (d *decoderBuilder) newOptTypeDecoder(t reflect.Type) decoderFunc {
	for t.Kind() == reflect.Pointer {
		t = t.Elem()
	}
	valueField, _ := t.FieldByName("Value")
	return func(n gjson.Result, v reflect.Value, state *decoderState) error {
		state.validateOptKind(n, valueField.Type)
		return unmarshalRaw(v.Addr(), n.Raw)
	}
}

func (d *decoderBuilder) newTimeTypeDecoder(t reflect.Type) decoderFunc {
	format := d.dateFormat
	return func(n gjson.Result, v reflect.Value, state *decoderState) error {
		parsed, err := time.Parse(format, n.Str)
		if err == nil {
			v.Set(reflect.ValueOf(parsed).Convert(t))
			return nil
		}

		if guardStrict(state, true) {
			return err
		}

		layouts := []string{
			"2006-01-02",
			"2006-01-02T15:04:05Z07:00",
			"2006-01-02T15:04:05Z0700",
			"2006-01-02T15:04:05",
			"2006-01-02 15:04:05Z07:00",
			"2006-01-02 15:04:05Z0700",
			"2006-01-02 15:04:05",
		}

		for _, layout := range layouts {
			parsed, err := time.Parse(layout, n.Str)
			if err == nil {
				v.Set(reflect.ValueOf(parsed).Convert(t))
				return nil
			}
		}

		return fmt.Errorf("unable to leniently parse date-time string: %s", n.Str)
	}
}

func guardStrict(state *decoderState, cond bool) bool {
	if cond && !state.strict {
		state.exactness = loose
	}
	return cond && state.strict
}

func canParseAsNumber(str string) bool {
	_, err := strconv.ParseFloat(str, 64)
	return err == nil
}

var stringType = reflect.TypeOf(string(""))

func guardUnknown(state *decoderState, v reflect.Value) bool {
	if have, ok := v.Interface().(interface{ IsKnown() bool }); guardStrict(state, ok && !have.IsKnown()) {
		return true
	}

	constantString, ok := v.Interface().(interface{ Default() string })
	named := v.Type() != stringType
	return guardStrict(state, ok && named && v.Equal(reflect.ValueOf(constantString.Default())))
}
