// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

// Package reflectvalue reads a reflect.Value as a static Go type.
package reflectvalue

import (
	"fmt"
	"reflect"
)

// As returns the value v holds, converted to T, or an error when v cannot
// become a T.
func As[T any](v reflect.Value) (T, error) {
	if target := reflect.TypeFor[T](); v.CanConvert(target) {
		v = v.Convert(target)
	}
	x, ok := reflect.TypeAssert[T](v)
	if !ok {
		return x, fmt.Errorf("reflectvalue: %s does not convert to %s", v.Type(), reflect.TypeFor[T]())
	}
	return x, nil
}
