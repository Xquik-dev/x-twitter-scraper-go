// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

package sentinel

import (
	"reflect"
	"sync"
)

// nullCache holds the sentinel of each map or slice type, keyed by that type.
var nullCache = struct {
	sync.RWMutex
	m map[reflect.Type]reflect.Value
}{m: map[reflect.Type]reflect.Value{}}

func NewNullSentinel[T any](mk func() T) T {
	t := reflect.TypeFor[T]()
	nullCache.Lock()
	defer nullCache.Unlock()
	null, ok := nullCache.m[t]
	if !ok {
		null = reflect.ValueOf(mk())
		nullCache.m[t] = null
	}
	x, _ := reflect.TypeAssert[T](null)
	return x
}

func IsValueNull(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Map, reflect.Slice:
		return isNull(v.Type(), v)
	}
	return false
}

func IsNull[T any](v T) bool {
	t := reflect.TypeFor[T]()
	switch t.Kind() {
	case reflect.Map, reflect.Slice:
		return isNull(t, reflect.ValueOf(v))
	}
	return false
}

// isNull reports whether v, a map or slice of type t, is the sentinel of t.
func isNull(t reflect.Type, v reflect.Value) bool {
	nullCache.RLock()
	defer nullCache.RUnlock()
	null, ok := nullCache.m[t]
	return ok && v.Pointer() == null.Pointer()
}
