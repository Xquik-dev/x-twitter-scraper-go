// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

package respjson_test

import (
	"encoding/json"
	"fmt"
	"github.com/Xquik-dev/x-twitter-scraper-go/internal/apijson"
	rj "github.com/Xquik-dev/x-twitter-scraper-go/packages/respjson"
	"reflect"
	"testing"
)

type UnionOfStringIntOrObject struct {
	OfString string    `json:",inline"`
	OfInt    int       `json:",inline"`
	Type     string    `json:"type"`
	Function SubFields `json:"function"`
	JSON     struct {
		OfString rj.Field
		OfInt    rj.Field
		Type     rj.Field
		Function rj.Field
		raw      string
	} `json:"-"`
}

func (u UnionOfStringIntOrObject) RawJSON() string { return u.JSON.raw }
func (r *UnionOfStringIntOrObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SubFields struct {
	OfBool bool   `json:",inline"`
	Name   string `json:"name" api:"required"`
	JSON   struct {
		OfBool      rj.Field
		Name        rj.Field
		ExtraFields map[string]rj.Field
		raw         string
	} `json:"-"`
}

func (r SubFields) RawJSON() string { return r.JSON.raw }
func (r *SubFields) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func TestUnmarshalUnion(t *testing.T) {
	for _, test := range []struct {
		name     string
		raw      string
		want     UnionOfStringIntOrObject
		metadata map[string]string
	}{
		{"string", `"123"`, UnionOfStringIntOrObject{OfString: "123"}, map[string]string{"string": `"123"`}},
		{"int", `123`, UnionOfStringIntOrObject{OfInt: 123}, map[string]string{"int": "123"}},
		{"zero", `0`, UnionOfStringIntOrObject{}, map[string]string{"int": "0"}},
		{"object", `{"type":"auto","function":{"name":"test_fn"}}`,
			UnionOfStringIntOrObject{Type: "auto", Function: SubFields{Name: "test_fn"}},
			map[string]string{"type": `"auto"`, "function": `{"name":"test_fn"}`, "function.name": `"test_fn"`}},
		{"inline", `{"type":"auto","function":true}`,
			UnionOfStringIntOrObject{Type: "auto", Function: SubFields{OfBool: true}},
			map[string]string{"type": `"auto"`, "function": "true", "function.bool": "true"}},
	} {
		t.Run(test.name, func(t *testing.T) {
			var res UnionOfStringIntOrObject
			if err := json.Unmarshal([]byte(test.raw), &res); err != nil {
				t.Fatal(err)
			}
			for label, err := range map[string]error{
				"rawJSON":       checkEqual(res.RawJSON(), test.raw),
				"string":        checkEqual(res.OfString, test.want.OfString),
				"int":           checkEqual(res.OfInt, test.want.OfInt),
				"type":          checkEqual(res.Type, test.want.Type),
				"function.name": checkEqual(res.Function.Name, test.want.Function.Name),
				"function.bool": checkEqual(res.Function.OfBool, test.want.Function.OfBool),
			} {
				if err != nil {
					t.Errorf("%s: %v", label, err)
				}
			}
			for label, field := range map[string]rj.Field{
				"string":        res.JSON.OfString,
				"int":           res.JSON.OfInt,
				"type":          res.JSON.Type,
				"function":      res.JSON.Function,
				"function.name": res.Function.JSON.Name,
				"function.bool": res.Function.JSON.OfBool,
			} {
				raw, present := test.metadata[label]
				status := shouldBeNullish
				if present {
					status = shouldBePresent
				}
				if err := checkMeta(field, raw, status); err != nil {
					t.Errorf("%s metadata: %v", label, err)
				}
			}
		})
	}
}

func checkEqual[T any](got, expected T) error {
	if reflect.DeepEqual(got, expected) {
		return nil
	}
	return fmt.Errorf("not equal: got %v, expected %v", got, expected)
}

type metaStatus int

const (
	shouldBePresent metaStatus = iota
	shouldBeNullish
	shouldBeInvalid
)

func checkMeta(got rj.Field, raw string, stat metaStatus) error {
	switch stat {
	case shouldBePresent:
		if !got.Valid() {
			return fmt.Errorf("expected field to be present, but got nullish")
		}
		if got.Raw() != raw {
			return fmt.Errorf("expected field to be present with raw value %v, but got %v", raw, got.Raw())
		}
	case shouldBeNullish:
		if got.Valid() || (got.Raw() != rj.Omitted && got.Raw() != rj.Null) {
			return fmt.Errorf("expected field to be nullish, but got %v", got.Raw())
		}
	case shouldBeInvalid:
		if !got.Valid() || got.Raw() == "" {
			return fmt.Errorf("expected field to be invalid, but got valid value %v", got.Raw())
		}
		if got.Raw() != raw {
			return fmt.Errorf("expected field to be invalid, but got valid value %v", got.Raw())
		}
	default:
		return fmt.Errorf("unknown metaStatus: %v", stat)
	}
	return nil
}
