// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

package apijson

import (
	"errors"
	"github.com/Xquik-dev/x-twitter-scraper-go/packages/param"
	"reflect"

	"github.com/tidwall/gjson"
)

var apiUnionType = reflect.TypeOf(param.APIUnion{})

func isStructUnion(t reflect.Type) bool {
	if t.Kind() != reflect.Struct {
		return false
	}
	for i := 0; i < t.NumField(); i++ {
		if t.Field(i).Type == apiUnionType && t.Field(i).Anonymous {
			return true
		}
	}
	return false
}

func RegisterDiscriminatedUnion[T any](key string, mappings map[string]reflect.Type) {
	var t T
	entry := unionEntry{
		discriminatorKey: key,
		variants:         []UnionVariant{},
	}
	for k, typ := range mappings {
		entry.variants = append(entry.variants, UnionVariant{
			DiscriminatorValue: k,
			Type:               typ,
		})
	}
	unionRegistry[reflect.TypeOf(t)] = entry
}

func (d *decoderBuilder) newStructUnionDecoder(t reflect.Type) decoderFunc {
	type variantDecoder struct {
		decoder decoderFunc
		field   reflect.StructField
	}
	decoders := []variantDecoder{}
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		if field.Anonymous && field.Type == apiUnionType {
			continue
		}

		decoder := d.typeDecoder(field.Type)
		decoders = append(decoders, variantDecoder{
			decoder: decoder,
			field:   field,
		})
	}

	type discriminatedDecoder struct {
		variantDecoder
		discriminator string
	}
	discriminatedDecoders := []discriminatedDecoder{}
	unionEntry, discriminated := unionRegistry[t]
	for _, variant := range unionEntry.variants {
		for _, decoder := range decoders {
			if decoder.field.Type.Elem() == variant.Type {
				discriminatedDecoders = append(discriminatedDecoders, discriminatedDecoder{
					decoder,
					variant.DiscriminatorValue,
				})
				break
			}
		}
	}

	return func(n gjson.Result, v reflect.Value, state *decoderState) error {
		if discriminated && n.Type == gjson.JSON && len(unionEntry.discriminatorKey) != 0 {
			discriminator := n.Get(EscapeSJSONKey(unionEntry.discriminatorKey))
			for _, decoder := range discriminatedDecoders {
				if discriminator.Type == gjson.String && discriminator.Str == decoder.discriminator {
					inner := v.FieldByIndex(decoder.field.Index)
					return decoder.decoder(n, inner, state)
				}
			}
			return errors.New("apijson: was not able to find discriminated union variant")
		}

		bestExactness := loose - 1
		bestVariant := -1
		for i, decoder := range decoders {
			if n.Type != gjson.JSON && decoder.field.Type.Kind() == reflect.Ptr {
				continue
			}

			sub := decoderState{strict: state.strict, exactness: exact}
			inner := v.FieldByIndex(decoder.field.Index)
			err := decoder.decoder(n, inner, &sub)
			if err != nil {
				continue
			}
			if sub.exactness == exact {
				bestExactness = exact
				bestVariant = i
				break
			}
			if sub.exactness > bestExactness {
				bestExactness = sub.exactness
				bestVariant = i
			}
		}

		if bestExactness < loose {
			return errors.New("apijson: was not able to coerce type as union")
		}

		if guardStrict(state, bestExactness != exact) {
			return errors.New("apijson: was not able to coerce type as union strictly")
		}

		for i := 0; i < len(decoders); i++ {
			if i == bestVariant {
				continue
			}
			v.FieldByIndex(decoders[i].field.Index).SetZero()
		}

		return nil
	}
}

func (d *decoderBuilder) newUnionDecoder(t reflect.Type) decoderFunc {
	unionEntry, ok := unionRegistry[t]
	if !ok {
		panic("apijson: couldn't find union of type " + t.String() + " in union registry")
	}
	decoders := []decoderFunc{}
	for _, variant := range unionEntry.variants {
		decoder := d.typeDecoder(variant.Type)
		decoders = append(decoders, decoder)
	}
	return func(n gjson.Result, v reflect.Value, state *decoderState) error {
		for idx, variant := range unionEntry.variants {
			decoder := decoders[idx]
			if variant.TypeFilter != n.Type {
				continue
			}

			if len(unionEntry.discriminatorKey) != 0 {
				discriminatorValue := n.Get(EscapeSJSONKey(unionEntry.discriminatorKey))
				if discriminatorValue.Type == gjson.String && discriminatorValue.Str == variant.DiscriminatorValue {
					inner := reflect.New(variant.Type).Elem()
					err := decoder(n, inner, state)
					v.Set(inner)
					return err
				}
			}
		}

		bestExactness := loose - 1
		for idx, variant := range unionEntry.variants {
			decoder := decoders[idx]
			if variant.TypeFilter != n.Type {
				continue
			}
			sub := decoderState{strict: state.strict, exactness: exact}
			inner := reflect.New(variant.Type).Elem()
			err := decoder(n, inner, &sub)
			if err != nil {
				continue
			}
			if sub.exactness == exact {
				v.Set(inner)
				return nil
			}
			if sub.exactness > bestExactness {
				v.Set(inner)
				bestExactness = sub.exactness
			}
		}

		if bestExactness < loose {
			return errors.New("apijson: was not able to coerce type as union")
		}

		if guardStrict(state, bestExactness != exact) {
			return errors.New("apijson: was not able to coerce type as union strictly")
		}

		return nil
	}
}
