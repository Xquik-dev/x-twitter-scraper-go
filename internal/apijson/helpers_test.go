// Copyright the Xquik contributors.
// SPDX-License-Identifier: Apache-2.0

package apijson

import (
	"reflect"
	"testing"
)

type registryValue struct {
	Kind string `json:"kind"`
}

type registryUnion interface{ registry() }

func TestFieldAndRegistryHelpers(t *testing.T) {
	field := Field{raw: `"tweet"`, status: valid}
	if field.IsNull() || field.IsMissing() || field.IsInvalid() || field.Raw() != `"tweet"` {
		t.Fatalf("field = %+v", field)
	}
	if !(Field{status: null}).IsNull() || !(Field{}).IsMissing() ||
		!(Field{status: invalid}).IsInvalid() {
		t.Fatal("field status helpers returned the wrong result")
	}

	variant := Discriminator[registryValue]("tweet")
	if variant.DiscriminatorValue != "tweet" || variant.Type != reflect.TypeOf(registryValue{}) {
		t.Fatalf("variant = %+v", variant)
	}
	RegisterUnion[registryUnion]("kind", variant)

	var wrapper UnionUnmarshaler[registryValue]
	if err := wrapper.UnmarshalJSON([]byte(`{"kind":"timeline"}`)); err != nil {
		t.Fatal(err)
	}
	if wrapper.Value.Kind != "timeline" {
		t.Fatalf("union wrapper = %+v", wrapper.Value)
	}
}

func TestMarshalerHelpersRejectOtherTypes(t *testing.T) {
	if err := unmarshalRaw(reflect.ValueOf(1), "1"); err == nil {
		t.Fatal("unmarshalRaw() accepted a type without UnmarshalJSON")
	}
	if _, err := marshalerEncoder(reflect.ValueOf(1)); err == nil {
		t.Fatal("marshalerEncoder() accepted a type without MarshalJSON")
	}
	if encoded, err := (&encoder{}).newTimeTypeEncoder()(reflect.ValueOf(1)); err == nil || encoded != nil {
		t.Fatalf("time encoder of an int = %q, %v", encoded, err)
	}
}
