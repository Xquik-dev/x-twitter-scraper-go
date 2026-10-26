// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"slices"
)

type kind int8

const (
	unknown kind = iota
	cycle
	free
	object
	array
	scalar
	union
)

type flag int8

const (
	unencoded flag = iota
	off
	on
)

func flagOf(value bool) flag {
	if value {
		return on
	}
	return off
}

type shape struct {
	kind     kind
	fields   map[string]field
	open     bool
	items    *shape
	enum     []string
	variants []*shape
	requires []string
}

type field struct {
	required flag
	nullable flag
	shape    *shape
}

func opaque(kind kind) *shape {
	return &shape{kind: kind}
}

func objectOf(fields map[string]field) *shape {
	return &shape{kind: object, fields: fields}
}

func scalarOf(values []string) *shape {
	if values == nil {
		return &shape{kind: scalar}
	}
	enum := []string{}
	for _, value := range values {
		if !slices.Contains(enum, value) {
			enum = append(enum, value)
		}
	}
	slices.Sort(enum)
	return &shape{kind: scalar, enum: enum}
}

func unionOf(variants []*shape) *shape {
	if len(variants) == 1 {
		return variants[0]
	}
	return &shape{kind: union, variants: variants}
}

func (s *shape) unread() bool {
	return s.kind == unknown || s.kind == cycle
}

func (s *shape) opaque() bool {
	return s.unread() || s.kind == free
}

func intersect(parts []*shape) *shape {
	var known, unions, rest, objects, arrays []*shape
	var requires []string
	for _, part := range parts {
		requires = append(requires, part.requires...)
		if part.kind != unknown {
			known = append(known, part)
		}
	}
	if len(known) == 0 {
		return &shape{kind: unknown, requires: requires}
	}
	for _, part := range known {
		switch part.kind {
		case union:
			unions = append(unions, part)
			continue
		case object:
			objects = append(objects, part)
		case array:
			arrays = append(arrays, part)
		}
		rest = append(rest, part)
	}
	if len(unions) > 0 {
		rest = append(append(rest, unions[1:]...), &shape{kind: unknown, requires: requires})
		variants := make([]*shape, len(unions[0].variants))
		for index, variant := range unions[0].variants {
			variants[index] = intersect(append(slices.Clone(rest), variant))
		}
		return &shape{kind: union, variants: variants}
	}
	if len(arrays) == len(known) {
		items := make([]*shape, len(arrays))
		for index, part := range arrays {
			items[index] = part.items
		}
		return &shape{kind: array, items: intersect(items)}
	}
	if len(objects) == 0 {
		return known[0]
	}
	merged := objectOf(map[string]field{})
	for _, part := range objects {
		merged.open = merged.open || part.open
		for name, next := range part.fields {
			earlier, seen := merged.fields[name]
			if !seen {
				merged.fields[name] = next
				continue
			}
			merged.fields[name] = field{
				required: flagOf(earlier.required == on || next.required == on),
				nullable: flagOf(earlier.nullable == on && next.nullable == on),
				shape:    intersect([]*shape{earlier.shape, next.shape}),
			}
		}
	}
	for _, name := range requires {
		if current, ok := merged.fields[name]; ok {
			current.required = on
			merged.fields[name] = current
		}
	}
	return merged
}

func canonical(s *shape) *shape {
	switch s.kind {
	case object:
		fields := make(map[string]field, len(s.fields))
		for name, current := range s.fields {
			current.shape = canonical(current.shape)
			fields[name] = current
		}
		return &shape{kind: object, fields: fields, open: s.open}
	case array:
		return &shape{kind: array, items: canonical(s.items)}
	case union:
		variants := make([]*shape, len(s.variants))
		for index, variant := range s.variants {
			variants[index] = canonical(variant)
		}
		return merge(variants)
	}
	return s
}

func merge(variants []*shape) *shape {
	var typed []*shape
	for _, variant := range variants {
		if variant.kind != cycle && variant.kind != unknown {
			typed = append(typed, variant)
		}
	}
	if len(typed) == 0 {
		typed = variants
	}
	if len(typed) == 1 {
		return typed[0]
	}
	groups := map[kind][]*shape{}
	for _, variant := range typed {
		groups[variant.kind] = append(groups[variant.kind], variant)
	}
	switch {
	case len(groups[free]) > 0:
		return opaque(free)
	case len(groups[object]) > 0:
		return mergeObjects(groups[object])
	case len(groups[array]) > 0:
		var items []*shape
		for _, variant := range groups[array] {
			items = append(items, variant.items)
		}
		return &shape{kind: array, items: merge(items)}
	case len(groups[scalar]) > 0:
		var values []string
		for _, variant := range groups[scalar] {
			if variant.enum == nil {
				return scalarOf(nil)
			}
			values = append(values, variant.enum...)
		}
		return scalarOf(values)
	}
	return opaque(unknown)
}

func mergeObjects(objects []*shape) *shape {
	merged := objectOf(map[string]field{})
	for _, current := range objects {
		merged.open = merged.open || current.open
	}
	names := map[string]bool{}
	for _, current := range objects {
		for name := range current.fields {
			names[name] = true
		}
	}
	for name := range names {
		var present []field
		for _, current := range objects {
			if value, ok := current.fields[name]; ok {
				present = append(present, value)
			}
		}
		required, nullable := on, off
		var shapes []*shape
		for _, value := range present {
			shapes = append(shapes, value.shape)
			switch {
			case value.required == unencoded:
				required = unencoded
			case value.required == off && required == on:
				required = off
			}
			switch {
			case value.nullable == unencoded:
				nullable = unencoded
			case value.nullable == on && nullable == off:
				nullable = on
			}
		}
		if len(present) < len(objects) && required != unencoded {
			required = off
		}
		merged.fields[name] = field{required: required, nullable: nullable, shape: merge(shapes)}
	}
	return merged
}
