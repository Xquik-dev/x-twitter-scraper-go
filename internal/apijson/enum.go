// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

package apijson

import (
	"fmt"
	"reflect"
	"slices"

	"github.com/tidwall/gjson"
)

type validationEntry struct {
	field       reflect.StructField
	legalValues struct {
		strings []string
		bools   int
		ints    []int64
	}
}

var validationRegistry = map[reflect.Type][]validationEntry{}

func RegisterFieldValidator[T any, V string | bool | int | float64](fieldName string, values ...V) {
	var t T
	parentType := reflect.TypeOf(t)

	if _, ok := validationRegistry[parentType]; !ok {
		validationRegistry[parentType] = []validationEntry{}
	}

	if parentType.Kind() != reflect.Struct {
		panic(fmt.Sprintf("apijson: cannot initialize validator for non-struct %s", parentType.String()))
	}

	var field reflect.StructField
	found := false
	for i := 0; i < parentType.NumField(); i++ {
		ptag, ok := parseJSONStructTag(parentType.Field(i))
		if ok && ptag.name == fieldName {
			field = parentType.Field(i)
			found = true
			break
		}
	}

	if !found {
		panic(fmt.Sprintf("apijson: cannot find field %s in struct %s", fieldName, parentType.String()))
	}

	newEntry := validationEntry{field: field}
	newEntry.legalValues.bools = -1

	list := reflect.ValueOf(values)
	if values, ok := reflect.TypeAssert[[]string](list); ok {
		newEntry.legalValues.strings = values
	} else if values, ok := reflect.TypeAssert[[]int](list); ok {
		newEntry.legalValues.ints = make([]int64, len(values))
		for i, value := range values {
			newEntry.legalValues.ints[i] = int64(value)
		}
	} else if values, ok := reflect.TypeAssert[[]bool](list); ok {
		for i, value := range values {
			var next int
			if value {
				next = 1
			}
			if i > 0 && newEntry.legalValues.bools != next {
				newEntry.legalValues.bools = -1
				break
			}
			newEntry.legalValues.bools = next
		}
	}

	validationRegistry[parentType] = append(validationRegistry[parentType], newEntry)
}

func (state *decoderState) validateString(v reflect.Value) {
	if state.validator == nil {
		return
	}
	if !slices.Contains(state.validator.legalValues.strings, v.String()) {
		state.exactness = loose
	}
}

func (state *decoderState) validateInt(v reflect.Value) {
	if state.validator == nil {
		return
	}
	if !slices.Contains(state.validator.legalValues.ints, v.Int()) {
		state.exactness = loose
	}
}

func (state *decoderState) validateBool(v reflect.Value) {
	if state.validator == nil {
		return
	}
	b := v.Bool()
	if state.validator.legalValues.bools == 1 && !b {
		state.exactness = loose
	} else if state.validator.legalValues.bools == 0 && b {
		state.exactness = loose
	}
}

func (state *decoderState) validateOptKind(node gjson.Result, t reflect.Type) {
	switch node.Type {
	case gjson.JSON:
		state.exactness = loose
	case gjson.Null:
		return
	case gjson.False, gjson.True:
		if t.Kind() != reflect.Bool {
			state.exactness = loose
		}
	case gjson.Number:
		switch t.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
			reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
			reflect.Float32, reflect.Float64:
			return
		default:
			state.exactness = loose
		}
	case gjson.String:
		if t.Kind() != reflect.String {
			state.exactness = loose
		}
	}
}
