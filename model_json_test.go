// Copyright the Xquik contributors.
// SPDX-License-Identifier: Apache-2.0

package xtwitterscraper

import (
	"go/ast"
	"go/parser"
	"go/token"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
)

type rawJSONModel interface {
	UnmarshalJSON([]byte) error
	RawJSON() string
}

func rawJSONModelTypes() []reflect.Type {
	module := reflect.TypeFor[Client]().PkgPath()
	visited := make(map[reflect.Type]bool)
	var models []reflect.Type
	var inspect func(reflect.Type)
	inspect = func(modelType reflect.Type) {
		owner := modelType
		if owner.Kind() == reflect.Pointer {
			owner = owner.Elem()
		}
		packagePath := owner.PkgPath()
		if visited[modelType] || packagePath != "" && packagePath != module && !strings.HasPrefix(packagePath, module+"/") {
			return
		}
		visited[modelType] = true
		for method := range modelType.Methods() {
			inspect(method.Type)
		}
		switch modelType.Kind() {
		case reflect.Struct:
			pointer := reflect.PointerTo(modelType)
			if ast.IsExported(modelType.Name()) && pointer.Implements(reflect.TypeFor[rawJSONModel]()) {
				models = append(models, modelType)
			}
			inspect(pointer)
			for field := range modelType.Fields() {
				inspect(field.Type)
			}
		case reflect.Array, reflect.Chan, reflect.Pointer, reflect.Slice:
			inspect(modelType.Elem())
		case reflect.Map:
			inspect(modelType.Key())
			inspect(modelType.Elem())
		case reflect.Func:
			for argument := range modelType.Ins() {
				inspect(argument)
			}
			for result := range modelType.Outs() {
				inspect(result)
			}
		}
	}
	inspect(reflect.TypeFor[Client]())
	inspect(reflect.TypeFor[Error]())
	return models
}

func TestGeneratedModelsPreserveRawJSON(t *testing.T) {
	raw := []byte(`{"_contract_probe":true}`)
	models := rawJSONModelTypes()
	registered := make(map[string]struct{}, len(models))
	for _, modelType := range models {
		if _, duplicate := registered[modelType.Name()]; duplicate {
			t.Fatalf("duplicate model registration: %s", modelType.Name())
		}
		registered[modelType.Name()] = struct{}{}
		t.Run(modelType.Name(), func(t *testing.T) {
			model := reflect.New(modelType).Interface().(rawJSONModel)
			if err := model.UnmarshalJSON(raw); err != nil {
				t.Fatalf("UnmarshalJSON() error = %v", err)
			}
			if got := model.RawJSON(); got != string(raw) {
				t.Fatalf("RawJSON() = %q, want %q", got, raw)
			}
		})
	}

	discovered := rawJSONReceiverNames(t)
	for name := range discovered {
		if _, ok := registered[name]; !ok {
			t.Errorf("generated model is not registered: %s", name)
		}
	}
	for name := range registered {
		if _, ok := discovered[name]; !ok {
			t.Errorf("registered model has no RawJSON method: %s", name)
		}
	}
	if len(registered) != len(discovered) {
		t.Fatalf(
			"registered %d generated models, discovered %d",
			len(registered),
			len(discovered),
		)
	}
}

func TestScorerWeightPreservesNull(t *testing.T) {
	raw := []byte(`{"context":"Production value unknown","signal":"favorite","weight":null}`)
	var model ComposeNewResponseComposePrepareResultScorerWeight
	if err := model.UnmarshalJSON(raw); err != nil {
		t.Fatal(err)
	}
	if model.Weight != nil || model.JSON.Weight.Raw() != "null" {
		t.Fatalf("weight = %#v, metadata = %#v; want explicit null", model.Weight, model.JSON.Weight)
	}
	if model.RawJSON() != string(raw) {
		t.Fatalf("RawJSON() = %q, want %q", model.RawJSON(), raw)
	}
}

func rawJSONReceiverNames(t *testing.T) map[string]struct{} {
	t.Helper()

	names := rawJSONReceiverNamesInDirectory(t, ".")
	aliasTargets := map[string]map[string]struct{}{
		"apierror": rawJSONReceiverNamesInDirectory(t, "internal/apierror"),
		"shared":   rawJSONReceiverNamesInDirectory(t, "shared"),
	}

	aliasFile, err := parser.ParseFile(token.NewFileSet(), "aliases.go", nil, 0)
	if err != nil {
		t.Fatal(err)
	}
	for _, declaration := range aliasFile.Decls {
		general, ok := declaration.(*ast.GenDecl)
		if !ok || general.Tok != token.TYPE {
			continue
		}
		for _, specification := range general.Specs {
			alias, ok := specification.(*ast.TypeSpec)
			if !ok || !alias.Assign.IsValid() {
				continue
			}
			target, ok := alias.Type.(*ast.SelectorExpr)
			if !ok {
				continue
			}
			packageName, ok := target.X.(*ast.Ident)
			if !ok {
				continue
			}
			targetNames := aliasTargets[packageName.Name]
			if _, ok := targetNames[target.Sel.Name]; ok {
				names[alias.Name.Name] = struct{}{}
			}
		}
	}
	return names
}

func rawJSONReceiverNamesInDirectory(
	t *testing.T,
	directory string,
) map[string]struct{} {
	t.Helper()

	paths, err := filepath.Glob(filepath.Join(directory, "*.go"))
	if err != nil {
		t.Fatal(err)
	}
	names := make(map[string]struct{})
	for _, path := range paths {
		if strings.HasSuffix(path, "_test.go") {
			continue
		}
		file, err := parser.ParseFile(token.NewFileSet(), path, nil, 0)
		if err != nil {
			t.Fatalf("parse %s: %v", path, err)
		}
		for _, declaration := range file.Decls {
			function, ok := declaration.(*ast.FuncDecl)
			if !ok || function.Name.Name != "RawJSON" || function.Recv == nil {
				continue
			}
			receiver := function.Recv.List[0].Type
			if pointer, ok := receiver.(*ast.StarExpr); ok {
				receiver = pointer.X
			}
			identifier, ok := receiver.(*ast.Ident)
			if ok && ast.IsExported(identifier.Name) {
				names[identifier.Name] = struct{}{}
			}
		}
	}
	return names
}
