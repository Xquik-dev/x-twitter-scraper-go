// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

package param

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"time"

	shimjson "github.com/Xquik-dev/x-twitter-scraper-go/internal/encoding/json"

	"github.com/tidwall/sjson"
)

// EncodedAsDate is not be stable and shouldn't be relied upon
type EncodedAsDate Opt[time.Time]

// EscapeSJSONKey escapes the characters sjson reads as path syntax. If we want to set a literal key value into JSON using sjson, we need to make sure it doesn't have
// special characters that sjson interprets as a path.
var EscapeSJSONKey = strings.NewReplacer("\\", "\\\\", "|", "\\|", "#", "\\#", "@", "\\@", "*", "\\*", ".", "\\.", ":", "\\:", "?", "\\?").Replace

func (m EncodedAsDate) MarshalJSON() ([]byte, error) {
	underlying := Opt[time.Time](m)
	bytes := underlying.MarshalJSONWithTimeLayout("2006-01-02")
	if len(bytes) > 0 {
		return bytes, nil
	}
	return underlying.MarshalJSON()
}

// MarshalObject comes with no guarantee of API stability.
func MarshalObject[T ParamStruct, U any](f T, underlying U) ([]byte, error) {
	return MarshalWithExtras(f, underlying, f.extraFields())
}

// MarshalWithExtras is used to marshal a struct with additional properties.
//
// Stability for the API of MarshalWithExtras is not guaranteed.
func MarshalWithExtras[T ParamStruct, U any](f T, underlying U, extras map[string]json.RawMessage) ([]byte, error) {
	if f.null() {
		return []byte("null"), nil
	} else if len(extras) > 0 {
		bytes, err := shimjson.Marshal(underlying)
		if err != nil {
			return nil, err
		}
		for k, v := range extras {
			if string(v) == string(Omit) {
				if b, e := sjson.DeleteBytes(bytes, k); e == nil {
					bytes = b
				}
				continue
			}
			bytes, err = sjson.SetBytes(bytes, EscapeSJSONKey(k), v)
			if err != nil {
				return nil, err
			}
		}
		return bytes, nil
	} else if ovr, ok := f.Overrides(); ok {
		return shimjson.Marshal(ovr)
	} else {
		return shimjson.Marshal(underlying, shimjson.WithSkipCompaction(true))
	}
}

// MarshalUnion comes with no guarantee of API stability.
func MarshalUnion[T ParamStruct](metadata T, variants ...json.Marshaler) ([]byte, error) {
	nPresent := 0
	presentIdx := -1
	for i, variant := range variants {
		if !IsOmitted(variant) {
			nPresent++
			presentIdx = i
		}
	}
	if nPresent == 0 || presentIdx == -1 {
		if metadata.null() {
			return []byte("null"), nil
		}
		if ovr, ok := metadata.Overrides(); ok {
			return shimjson.Marshal(ovr)
		}
		return []byte(`null`), nil
	} else if nPresent > 1 {
		return nil, &json.MarshalerError{
			Type: typeFor[T](),
			Err:  fmt.Errorf("expected union to have only one present variant, got %d", nPresent),
		}
	}
	return shimjson.Marshal(variants[presentIdx], shimjson.WithSkipCompaction(true))
}

func typeFor[T any]() reflect.Type {
	var v T
	if t := reflect.TypeOf(v); t != nil {
		return t
	}
	return reflect.TypeOf((*T)(nil)).Elem()
}
