// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package json

import (
	stdjson "encoding/json"
	"reflect"
	"testing"
)

func TestStructTagObjectKey(t *testing.T) {
	tests := []struct {
		CaseName
		tag   reflect.StructTag
		value string
		key   string
	}{
		{Name(""), `json:"$%-/"`, "2x", "$%-/"},
		{Name(""), `json:"0123456789"`, "3x", "0123456789"},
		{Name(""), `json:"ABCDEFGHIJKLMO"`, "4x", "ABCDEFGHIJKLMO"},
		{Name(""), `json:"PQRSTUVWXYZ_"`, "5x", "PQRSTUVWXYZ_"},
		{Name(""), `json:"abcdefghijklmno"`, "6x", "abcdefghijklmno"},
		{Name(""), `json:"pqrstuvwxyz"`, "7x", "pqrstuvwxyz"},
		{Name(""), `json:"色は匂へど"`, "いろはにほへと", "色は匂へど"},
		{Name(""), `json:"-,"`, "foo", "-"},
		{Name(""), ``, "Pour Moi", "V"},
		{Name(""), `jsom:"Misnamed"`, "Animal Kingdom", "V"},
		{Name(""), `:"BadFormat"`, "Orfevre", "V"},
		{Name(""), `json:" !\"#&'()*+,."`, "Reliable Man", "V"},
		{Name(""), `json:"text/html%"`, "brut", "text/html%"},                                           // https://golang.org/issue/2718
		{Name(""), `json:"!#$%&()*+-./:;<=>?@[]^_{|}~ "`, "Union Rags", "!#$%&()*+-./:;<=>?@[]^_{|}~ "}, // https://golang.org/issue/3546
		{Name(""), `json:"With space"`, "Perreddu", "With space"},
		{Name(""), `json:"Ελλάδα"`, "Loukanikos", "Ελλάδα"},
		{Name(""), `json:"random,random"`, "Unknown Option", "random"},
	}
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			field := reflect.StructField{Name: "V", Type: reflect.TypeFor[string](), Tag: tt.tag}
			v := reflect.New(reflect.StructOf([]reflect.StructField{field})).Elem()
			v.Field(0).SetString(tt.value)
			b, err := Marshal(v.Interface())
			if err != nil {
				t.Fatalf("%s: Marshal error: %v", tt.Where, err)
			}
			var got map[string]string
			if err := stdjson.Unmarshal(b, &got); err != nil {
				t.Fatalf("%s: Unmarshal error: %v", tt.Where, err)
			}
			if len(got) != 1 || got[tt.key] != tt.value {
				t.Fatalf("%s: Marshal = %#q, want key %q with value %q", tt.Where, b, tt.key, tt.value)
			}
		})
	}
}
