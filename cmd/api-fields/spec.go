// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/Xquik-dev/x-twitter-scraper-go/internal/apicheck"
)

var (
	compositeKeys = []string{"$ref", "allOf", "oneOf", "anyOf", "description", "example", "examples", "title", "deprecated", "readOnly", "writeOnly"}
	formTypes     = []string{"multipart/form-data", "application/x-www-form-urlencoded"}
)

type specReader struct {
	document map[string]any
	problems map[string]bool
}

type operationShapes struct {
	query    *shape
	body     *shape
	response *shape
}

func (r specReader) resolve(reference string) any {
	var current any = r.document
	for _, part := range strings.Split(strings.TrimPrefix(reference, "#/"), "/") {
		container, _ := current.(map[string]any)
		current = container[strings.ReplaceAll(strings.ReplaceAll(part, "~1", "/"), "~0", "~")]
	}
	if current == nil {
		r.problems[fmt.Sprintf("The field check cannot resolve %s in the API spec. Fix the reference in the spec.", reference)] = true
	}
	return current
}

func requiredNames(schema map[string]any) []string {
	var names []string
	list, _ := schema["required"].([]any)
	for _, name := range list {
		if text, ok := name.(string); ok {
			names = append(names, text)
		}
	}
	return names
}

func wireValue(value any) string {
	switch typed := value.(type) {
	case string:
		return typed
	case float64:
		return strconv.FormatFloat(typed, 'f', -1, 64)
	case bool:
		return strconv.FormatBool(typed)
	}
	return "null"
}

func typeNames(schema map[string]any) ([]string, bool) {
	switch typed := schema["type"].(type) {
	case string:
		return []string{typed}, typed == "null"
	case []any:
		var names []string
		nullable := false
		for _, value := range typed {
			name, _ := value.(string)
			if name == "null" {
				nullable = true
			} else {
				names = append(names, name)
			}
		}
		return names, nullable
	}
	return nil, false
}

func (r specReader) shape(raw any, stack []string) (*shape, bool) {
	schema, ok := raw.(map[string]any)
	if !ok {
		return opaque(unknown), false
	}
	parts, nullable, composite := r.composites(schema, stack)
	if !composite {
		return r.plain(schema, stack)
	}
	rest := without(schema, compositeKeys...)
	next, null := r.plain(rest, stack)
	if len(without(rest, "type", "nullable")) > 0 {
		parts = append(parts, next)
	}
	return intersect(parts), nullable || null
}

func (r specReader) composites(schema map[string]any, stack []string) ([]*shape, bool, bool) {
	var parts []*shape
	nullable, composite := false, false
	if reference, ok := schema["$ref"].(string); ok {
		composite = true
		if slices.Contains(stack, reference) {
			parts = append(parts, opaque(cycle))
		} else {
			next, null := r.shape(r.resolve(reference), append(slices.Clone(stack), reference))
			parts = append(parts, next)
			nullable = null
		}
	}
	if list, ok := schema["allOf"].([]any); ok {
		composite = true
		for _, part := range list {
			next, _ := r.shape(part, stack)
			parts = append(parts, next)
		}
	}
	for _, key := range []string{"oneOf", "anyOf"} {
		if list, ok := schema[key].([]any); ok {
			composite = true
			next, null := r.choices(list, stack)
			parts = append(parts, next)
			nullable = nullable || null
		}
	}
	return parts, nullable, composite
}

func (r specReader) plain(schema map[string]any, stack []string) (*shape, bool) {
	names, nullable := typeNames(schema)
	if value, ok := schema["nullable"].(bool); ok && value {
		nullable = true
	}
	if len(names) == 1 && names[0] == "null" {
		return opaque(unknown), true
	}
	_, hasProperties := schema["properties"]
	_, hasEnum := schema["enum"]
	_, hasConst := schema["const"]
	if requires := requiredNames(schema); len(requires) > 0 && !hasProperties && !hasEnum && !hasConst && (len(names) == 0 || slices.Equal(names, []string{"object"})) {
		return &shape{kind: unknown, requires: requires}, nullable
	}
	if value, ok := schema["const"]; ok {
		return scalarOf([]string{wireValue(value)}), nullable
	}
	if values, ok := schema["enum"].([]any); ok {
		enum := []string{}
		for _, value := range values {
			if value == nil {
				nullable = true
			} else {
				enum = append(enum, wireValue(value))
			}
		}
		return scalarOf(enum), nullable
	}
	properties, _ := schema["properties"].(map[string]any)
	if slices.Contains(names, "object") || hasProperties {
		if len(properties) == 0 {
			return opaque(free), nullable
		}
		return r.object(schema, properties, stack), nullable
	}
	if slices.Contains(names, "array") {
		_, tuple := schema["prefixItems"]
		items, ok := schema["items"]
		if !ok || tuple {
			return &shape{kind: array, items: opaque(unknown)}, nullable
		}
		next, _ := r.shape(items, stack)
		return &shape{kind: array, items: next}, nullable
	}
	return scalarOf(nil), nullable
}

func without(schema map[string]any, keys ...string) map[string]any {
	rest := map[string]any{}
	for key, value := range schema {
		if !slices.Contains(keys, key) {
			rest[key] = value
		}
	}
	return rest
}

func (r specReader) choices(choices []any, stack []string) (*shape, bool) {
	var variants []*shape
	nullable := false
	for _, choice := range choices {
		if option, ok := choice.(map[string]any); ok {
			enum, _ := option["enum"].([]any)
			if option["type"] == "null" || (len(enum) == 1 && enum[0] == nil) {
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

func (r specReader) object(schema, properties map[string]any, stack []string) *shape {
	required := map[string]bool{}
	for _, name := range requiredNames(schema) {
		required[name] = true
	}
	result := objectOf(map[string]field{})
	for name, property := range properties {
		next, nullable := r.shape(property, stack)
		result.fields[name] = field{required: flagOf(required[name]), nullable: flagOf(nullable), shape: next}
	}
	if extra, present := schema["additionalProperties"]; present && extra != false && extra != nil {
		result.open = true
	}
	return result
}

func (r specReader) mapOf(value any) map[string]any {
	object, _ := value.(map[string]any)
	if reference, ok := object["$ref"].(string); ok {
		object, _ = r.resolve(reference).(map[string]any)
	}
	return object
}

func (r specReader) media(content any, accept func(string) bool) []*shape {
	types := r.mapOf(content)
	names := make([]string, 0, len(types))
	for name := range types {
		names = append(names, name)
	}
	slices.Sort(names)
	var shapes []*shape
	for _, name := range names {
		media := r.mapOf(types[name])
		base, _, _ := strings.Cut(name, ";")
		if schema, ok := media["schema"]; ok && accept(strings.TrimSpace(base)) {
			next, _ := r.shape(schema, nil)
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
	parameters := map[string]map[string]any{}
	for _, owner := range []map[string]any{operation.Item, operation.Operation} {
		list, _ := owner["parameters"].([]any)
		for _, raw := range list {
			parameter := r.mapOf(raw)
			where, _ := parameter["in"].(string)
			name, _ := parameter["name"].(string)
			parameters[where+" "+name] = parameter
		}
	}
	query := objectOf(map[string]field{})
	for key, parameter := range parameters {
		where, name, _ := strings.Cut(key, " ")
		if where != "query" {
			continue
		}
		next, _ := r.shape(parameter["schema"], nil)
		required, _ := parameter["required"].(bool)
		query.fields[name] = field{required: flagOf(required), shape: next}
	}
	content := r.mapOf(r.mapOf(operation.Operation["requestBody"])["content"])
	bodies := r.media(content, func(name string) bool { return name == "application/json" })
	if len(bodies) == 0 {
		bodies = r.media(content, func(name string) bool { return slices.Contains(formTypes, name) })
	}
	shapes := operationShapes{query: query, body: optional(bodies)}
	responses := r.mapOf(operation.Operation["responses"])
	codes := make([]string, 0, len(responses))
	for code := range responses {
		codes = append(codes, code)
	}
	slices.Sort(codes)
	var results []*shape
	for _, code := range codes {
		if strings.HasPrefix(code, "2") {
			results = append(results, r.media(r.mapOf(responses[code])["content"], func(name string) bool {
				return name == "application/json" || strings.HasSuffix(name, "+json")
			})...)
		}
	}
	shapes.response = optional(results)
	return shapes
}
