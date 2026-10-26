// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"encoding/json"
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/Xquik-dev/x-twitter-scraper-go/internal/apicheck"
)

var (
	compositeKeys = []string{"$ref", "allOf", "oneOf", "anyOf", "description", "example", "examples", "title", "deprecated", "readOnly", "writeOnly"}
	bodyTypes     = []string{"application/json", "multipart/form-data", "application/x-www-form-urlencoded"}
)

type schema struct {
	object               bool
	extra                bool
	Ref                  *string            `json:"$ref"`
	AllOf                []*schema          `json:"allOf"`
	OneOf                []*schema          `json:"oneOf"`
	AnyOf                []*schema          `json:"anyOf"`
	Type                 json.RawMessage    `json:"type"`
	Nullable             bool               `json:"nullable"`
	Required             []string           `json:"required"`
	Properties           map[string]*schema `json:"properties"`
	AdditionalProperties json.RawMessage    `json:"additionalProperties"`
	Enum                 []json.RawMessage  `json:"enum"`
	Const                json.RawMessage    `json:"const"`
	Items                *schema            `json:"items"`
	PrefixItems          json.RawMessage    `json:"prefixItems"`
}

func (s *schema) UnmarshalJSON(data []byte) error {
	keys, isObject := apicheck.Decode[map[string]json.RawMessage](data)
	if !isObject {
		return nil
	}
	type fields schema
	decoded, _ := apicheck.Decode[fields](data)
	*s = schema(decoded)
	s.object = true
	for key := range keys {
		if !slices.Contains(compositeKeys, key) && key != "type" && key != "nullable" {
			s.extra = true
		}
	}
	return nil
}

type reference struct {
	Ref *string `json:"$ref"`
}

type parameter struct {
	Name     string  `json:"name"`
	In       string  `json:"in"`
	Required bool    `json:"required"`
	Schema   *schema `json:"schema"`
}

type withContent struct {
	Content json.RawMessage `json:"content"`
}

type mediaType struct {
	Schema json.RawMessage `json:"schema"`
}

type specReader struct {
	document json.RawMessage
	objects  map[string]map[string]json.RawMessage
	schemas  map[string]*schema
	problems map[string]bool
}

func newSpecReader(document json.RawMessage) specReader {
	return specReader{document: document, objects: map[string]map[string]json.RawMessage{}, schemas: map[string]*schema{}, problems: map[string]bool{}}
}

type operationShapes struct {
	query    *shape
	body     *shape
	response *shape
}

func (r specReader) resolve(pointer string) json.RawMessage {
	current, prefix := r.document, ""
	for _, part := range strings.Split(strings.TrimPrefix(pointer, "#/"), "/") {
		container, seen := r.objects[prefix]
		if !seen {
			container, _ = apicheck.Decode[map[string]json.RawMessage](current)
			r.objects[prefix] = container
		}
		prefix += "/" + part
		current = container[strings.ReplaceAll(strings.ReplaceAll(part, "~1", "/"), "~0", "~")]
	}
	if len(current) == 0 || string(current) == "null" {
		r.problems[fmt.Sprintf("The field check cannot resolve %s in the API spec. Fix the reference in the spec.", pointer)] = true
	}
	return current
}

func (r specReader) schemaAt(pointer string) *schema {
	if cached, ok := r.schemas[pointer]; ok {
		return cached
	}
	found, _ := apicheck.Decode[*schema](r.resolve(pointer))
	r.schemas[pointer] = found
	return found
}

func resolved[T any](r specReader, raw json.RawMessage) T {
	if target, _ := apicheck.Decode[reference](raw); target.Ref != nil {
		raw = r.resolve(*target.Ref)
	}
	value, _ := apicheck.Decode[T](raw)
	return value
}

func wireValue(value json.RawMessage) string {
	var text string
	var number float64
	var truth bool
	switch {
	case string(value) == "null":
	case json.Unmarshal(value, &text) == nil:
		return text
	case json.Unmarshal(value, &number) == nil:
		return strconv.FormatFloat(number, 'f', -1, 64)
	case json.Unmarshal(value, &truth) == nil:
		return strconv.FormatBool(truth)
	}
	return "null"
}

func (s *schema) typeNames() ([]string, bool) {
	if name, ok := apicheck.Decode[string](s.Type); ok && string(s.Type) != "null" {
		return []string{name}, name == "null"
	}
	list, ok := apicheck.Decode[[]json.RawMessage](s.Type)
	if !ok || list == nil {
		return nil, false
	}
	var names []string
	nullable := false
	for _, value := range list {
		name, _ := apicheck.Decode[string](value)
		if name == "null" {
			nullable = true
		} else {
			names = append(names, name)
		}
	}
	return names, nullable
}

func (r specReader) shape(current *schema, stack []string) (*shape, bool) {
	if current == nil || !current.object {
		return opaque(unknown), false
	}
	parts, nullable, composite := r.composites(current, stack)
	next, null := r.plain(current, stack)
	if !composite {
		return next, null
	}
	if current.extra {
		parts = append(parts, next)
	}
	return intersect(parts), nullable || null
}

func (r specReader) composites(current *schema, stack []string) ([]*shape, bool, bool) {
	var parts []*shape
	nullable, composite := false, false
	if current.Ref != nil {
		composite = true
		pointer := *current.Ref
		if slices.Contains(stack, pointer) {
			parts = append(parts, opaque(cycle))
		} else {
			next, null := r.shape(r.schemaAt(pointer), append(slices.Clone(stack), pointer))
			parts = append(parts, next)
			nullable = null
		}
	}
	if current.AllOf != nil {
		composite = true
		for _, part := range current.AllOf {
			next, _ := r.shape(part, stack)
			parts = append(parts, next)
		}
	}
	for _, list := range [][]*schema{current.OneOf, current.AnyOf} {
		if list != nil {
			composite = true
			next, null := r.choices(list, stack)
			parts = append(parts, next)
			nullable = nullable || null
		}
	}
	return parts, nullable, composite
}

func (r specReader) plain(current *schema, stack []string) (*shape, bool) {
	names, nullable := current.typeNames()
	if current.Nullable {
		nullable = true
	}
	if len(names) == 1 && names[0] == "null" {
		return opaque(unknown), true
	}
	hasProperties, hasEnum, hasConst := current.Properties != nil, current.Enum != nil, len(current.Const) > 0
	if len(current.Required) > 0 && !hasProperties && !hasEnum && !hasConst && (len(names) == 0 || slices.Equal(names, []string{"object"})) {
		return &shape{kind: unknown, requires: current.Required}, nullable
	}
	if hasConst {
		return scalarOf([]string{wireValue(current.Const)}), nullable
	}
	if hasEnum {
		enum := []string{}
		for _, value := range current.Enum {
			if string(value) == "null" {
				nullable = true
			} else {
				enum = append(enum, wireValue(value))
			}
		}
		return scalarOf(enum), nullable
	}
	if slices.Contains(names, "object") || hasProperties {
		if len(current.Properties) == 0 {
			return opaque(free), nullable
		}
		return r.object(current, stack), nullable
	}
	if slices.Contains(names, "array") {
		if current.Items == nil || len(current.PrefixItems) > 0 {
			return &shape{kind: array, items: opaque(unknown)}, nullable
		}
		next, _ := r.shape(current.Items, stack)
		return &shape{kind: array, items: next}, nullable
	}
	return scalarOf(nil), nullable
}

func (r specReader) choices(choices []*schema, stack []string) (*shape, bool) {
	var variants []*shape
	nullable := false
	for _, choice := range choices {
		if choice != nil && choice.object {
			names, _ := choice.typeNames()
			if slices.Equal(names, []string{"null"}) || (len(choice.Enum) == 1 && string(choice.Enum[0]) == "null") {
				nullable = true
				continue
			}
		}
		variant, null := r.shape(choice, stack)
		nullable = nullable || null
		variants = append(variants, variant)
	}
	if len(variants) == 0 {
		return opaque(unknown), nullable
	}
	return unionOf(variants), nullable
}

func (r specReader) object(current *schema, stack []string) *shape {
	required := map[string]bool{}
	for _, name := range current.Required {
		required[name] = true
	}
	result := objectOf(map[string]field{})
	for name, property := range current.Properties {
		next, nullable := r.shape(property, stack)
		result.fields[name] = field{required: flagOf(required[name]), nullable: flagOf(nullable), shape: next}
	}
	if extra := string(current.AdditionalProperties); extra != "" && extra != "false" && extra != "null" {
		result.open = true
	}
	return result
}

func (r specReader) media(content json.RawMessage, accept func(string) bool) []*shape {
	types := resolved[map[string]json.RawMessage](r, content)
	names := make([]string, 0, len(types))
	for name := range types {
		names = append(names, name)
	}
	slices.Sort(names)
	var shapes []*shape
	for _, name := range names {
		media := resolved[mediaType](r, types[name])
		base, _, _ := strings.Cut(name, ";")
		if len(media.Schema) > 0 && accept(strings.TrimSpace(base)) {
			found, _ := apicheck.Decode[*schema](media.Schema)
			next, _ := r.shape(found, nil)
			shapes = append(shapes, next)
		}
	}
	return shapes
}

func optional(shapes []*shape) *shape {
	if len(shapes) == 0 {
		return nil
	}
	return unionOf(shapes)
}

func (r specReader) operation(operation apicheck.SpecOperation) operationShapes {
	parameters := map[string]parameter{}
	for _, raw := range slices.Concat(operation.Item.Parameters, operation.Operation.Parameters) {
		found := resolved[parameter](r, raw)
		parameters[found.In+" "+found.Name] = found
	}
	query := objectOf(map[string]field{})
	for _, found := range parameters {
		if found.In != "query" {
			continue
		}
		next, _ := r.shape(found.Schema, nil)
		query.fields[found.Name] = field{required: flagOf(found.Required), shape: next}
	}
	body := resolved[withContent](r, operation.Operation.RequestBody)
	bodies := r.media(body.Content, func(name string) bool { return slices.Contains(bodyTypes, name) })
	shapes := operationShapes{query: query, body: optional(bodies)}
	responses := resolved[map[string]json.RawMessage](r, operation.Operation.Responses)
	codes := make([]string, 0, len(responses))
	for code := range responses {
		codes = append(codes, code)
	}
	slices.Sort(codes)
	var results []*shape
	for _, code := range codes {
		if strings.HasPrefix(code, "2") {
			results = append(results, r.media(resolved[withContent](r, responses[code]).Content, func(name string) bool {
				return name == "application/json" || strings.HasSuffix(name, "+json")
			})...)
		}
	}
	shapes.response = optional(results)
	return shapes
}
