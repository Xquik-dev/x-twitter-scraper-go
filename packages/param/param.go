// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

// Package param holds the types that build request params, such as [Opt] for an
// optional field and [Null] for an explicit null.
package param

import (
	"encoding/json"
	"github.com/Xquik-dev/x-twitter-scraper-go/internal/encoding/json/sentinel"
	"reflect"
)

// NullStruct is used to set a struct to the JSON value null.
// Check for null structs with [IsNull].
//
// Only the first type parameter should be provided,
// the type PtrT will be inferred.
//
//	json.Marshal(param.NullStruct[MyStruct]()) -> 'null'
//
// To send null to an [Opt] field use [Null].
func NullStruct[T ParamStruct, PtrT InferPtr[T]]() T {
	var t T
	pt := PtrT(&t)
	pt.setMetadata(nil)
	return *pt
}

// Override replaces the value of a struct with raw JSON.
//
// Only the first type parameter should be provided,
// the type PtrT will be inferred.
//
//	param.Override[MyStruct](json.RawMessage(`{"foo": "bar"}`))
//
// The public fields of the returned struct T will be unset.
//
// To override a specific field in a struct, use its [SetExtraFields] method.
func Override[T ParamStruct, PtrT InferPtr[T]](v json.RawMessage) T {
	var t T
	pt := PtrT(&t)
	pt.setMetadata(v)
	return *pt
}

// SetJSON configures a param struct to serialize with the provided raw JSON data.
// Use this when you have existing JSON that you want to send as request parameters.
//
//	var req example.NewUserParams
//	var rawJSON = []byte(`{"name": "...", "age": 40}`)
//	param.SetJSON(rawJSON, &req)
//	res, err := client.Users.New(ctx, req)
//
// Note: The struct's existing fields will be ignored; only the provided JSON is serialized.
func SetJSON(rawJSON []byte, ptr anyParamStruct) {
	ptr.setMetadata(json.RawMessage(rawJSON))
}

// IsOmitted returns true if v is the zero value of its type.
//
// If IsOmitted is true, and the field uses a `json:"...,omitzero"` tag,
// the field will be omitted from the request.
//
// If v is set explicitly to the JSON value "null", IsOmitted returns false.
func IsOmitted[T any](v T) bool {
	value := reflect.ValueOf(v)
	if !value.IsValid() {
		return false
	}
	if o, ok := reflect.TypeAssert[Optional](value); ok {
		return o.isZero()
	}
	return value.IsZero()
}

// IsNull returns true if v was set to the JSON value null.
//
// To set a param to null use [NullStruct], [Null], [NullMap], or [NullSlice]
// depending on the type of v.
//
// IsNull returns false if the value is omitted.
func IsNull[T any](v T) bool {
	if nullable, ok := reflect.TypeAssert[ParamNullable](reflect.ValueOf(&v).Elem()); ok {
		return nullable.null()
	}

	switch reflect.TypeOf(v).Kind() {
	case reflect.Slice, reflect.Map:
		return sentinel.IsNull(v)
	}

	return false
}

// ParamNullable encapsulates all structs in parameters,
// and all [Opt] types in parameters.
type ParamNullable interface {
	null() bool
}

// ParamStruct represents the set of all structs that are
// used in API parameters, by convention these usually end in
// "Params" or "Param".
type ParamStruct interface {
	Overrides() (json.RawMessage, bool)
	null() bool
	extraFields() map[string]json.RawMessage
}

type anyParamStruct interface {
	setMetadata(json.RawMessage)
}

// InferPtr is an implementation detail and should never be explicitly set.
type InferPtr[T ParamStruct] interface {
	setMetadata(json.RawMessage)
	*T
}

// APIObject should be embedded in api object fields, preferably using an alias to make private
type APIObject struct{ metadata }

// APIUnion should be embedded in all api unions fields, preferably using an alias to make private
type APIUnion struct{ metadata }

// Overrides returns the value of the struct when it is created with
// [Override], the second argument helps differentiate an explicit null.
func (m metadata) Overrides() (json.RawMessage, bool) {
	switch value := m.value.(type) {
	case metadataNull:
		return json.RawMessage("null"), true
	case metadataJSON:
		return json.RawMessage(value), true
	}
	return nil, false
}

// ExtraFields returns the extra fields added to the JSON object.
func (m metadata) ExtraFields() map[string]json.RawMessage {
	if extras, ok := m.value.(metadataExtraFields); ok {
		return extras
	}
	return nil
}

// Omit can be used with [metadata.SetExtraFields] to ensure that a
// required field is omitted. It is the nil raw message.
var Omit json.RawMessage

// SetExtraFields adds extra fields to the JSON object.
//
// SetExtraFields will override any existing fields with the same key.
// For security reasons, ensure this is only used with trusted input data.
//
// To intentionally omit a required field, use [Omit].
//
//	foo.SetExtraFields(map[string]json.RawMessage{"bar": param.Omit})
//
// If the struct already contains the field ExtraFields, then this
// method will have no effect.
func (m *metadata) SetExtraFields(extraFields map[string]json.RawMessage) {
	m.value = metadataExtraFields(extraFields)
}

func (m metadata) extraFields() map[string]json.RawMessage { return m.ExtraFields() }

func (m metadata) null() bool {
	switch value := m.value.(type) {
	case metadataNull:
		return true
	case metadataJSON:
		return string(value) == "null"
	}
	return false
}

type metadata struct{ value metadataValue }

type metadataValue interface{ isMetadata() }

type metadataNull struct{}
type metadataExtraFields map[string]json.RawMessage
type metadataJSON json.RawMessage

func (metadataNull) isMetadata()        {}
func (metadataExtraFields) isMetadata() {}
func (metadataJSON) isMetadata()        {}

func (m *metadata) setMetadata(override json.RawMessage) {
	if override == nil {
		m.value = metadataNull{}
		return
	}
	m.value = metadataJSON(override)
}
