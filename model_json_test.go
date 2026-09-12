// Copyright the Xquik contributors.
// SPDX-License-Identifier: Apache-2.0

package xtwitterscraper

import (
	"go/ast"
	"go/types"
	"maps"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/Xquik-dev/x-twitter-scraper-go/packages/respjson"
	"golang.org/x/tools/go/packages"
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
	if !maps.Equal(registered, discovered) {
		t.Fatalf("generated model inventories differ: registered=%v, discovered=%v", registered, discovered)
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

func TestRepostTimestampStates(t *testing.T) {
	matched := 0
	for _, modelType := range rawJSONModelTypes() {
		timestampField, exists := modelType.FieldByName("RetweetedAt")
		if !exists {
			continue
		}
		matched++
		for _, raw := range []string{`{}`, `{"retweetedAt":null}`, `{"retweetedAt":"2026-09-11T04:35:22Z"}`} {
			t.Run(modelType.Name()+"/"+raw, func(t *testing.T) {
				value := reflect.New(modelType)
				model := value.Interface().(rawJSONModel)
				if err := model.UnmarshalJSON([]byte(raw)); err != nil {
					t.Fatal(err)
				}
				got := value.Elem().FieldByName("RetweetedAt").Interface().(time.Time)
				owner := value.Elem()
				for _, index := range timestampField.Index[:len(timestampField.Index)-1] {
					owner = owner.Field(index)
				}
				field := owner.FieldByName("JSON").FieldByName("RetweetedAt").Interface().(respjson.Field)
				want := strings.TrimSuffix(strings.TrimPrefix(raw, `{"retweetedAt":`), "}")
				if raw == `{}` {
					want = ""
				}
				if field.Raw() != want || model.RawJSON() != raw {
					t.Fatalf("timestamp metadata = %q, payload = %q; want %q, %q", field.Raw(), model.RawJSON(), want, raw)
				}
				present := strings.HasPrefix(want, `"`)
				if field.Valid() != present || got.IsZero() == present {
					t.Fatalf("timestamp = %v, valid = %v; expected timestamp present = %v", got, field.Valid(), present)
				}
				if present && got.Format(time.RFC3339) != strings.Trim(want, `"`) {
					t.Fatalf("decoded timestamp = %v; want %s", got, want)
				}
			})
		}
	}
	if matched == 0 {
		t.Fatal("no timestamp models discovered")
	}
}

func TestNumericOrStringResponseValues(t *testing.T) {
	matched := 0
	for _, modelType := range rawJSONModelTypes() {
		model, ok := reflect.New(modelType).Interface().(interface {
			rawJSONModel
			AsFloat() float64
			AsString() string
		})
		if !ok {
			continue
		}
		matched++
		for _, test := range []struct {
			raw, text string
			number    float64
		}{
			{raw: "1789187722000", number: 1789187722000, text: "1789187722000"},
			{raw: `"1789187722000"`, number: 1789187722000, text: "1789187722000"},
			{raw: `"scheduled"`, text: "scheduled"},
			{raw: "null"},
		} {
			t.Run(modelType.Name()+"/"+test.raw, func(t *testing.T) {
				if err := model.UnmarshalJSON([]byte(test.raw)); err != nil {
					t.Fatal(err)
				}
				if model.AsFloat() != test.number || model.AsString() != test.text || model.RawJSON() != test.raw {
					t.Fatalf("decoded number = %v, text = %q, raw = %q; want %v, %q, %q", model.AsFloat(), model.AsString(), model.RawJSON(), test.number, test.text, test.raw)
				}
			})
		}
	}
	if matched == 0 {
		t.Fatal("no numeric-or-string response models discovered")
	}
}

func TestConnectionAttemptResponseVariants(t *testing.T) {
	for status, want := range map[string]any{
		"pending":             XAccountConnectionAttemptGetResponsePending{},
		"success":             XAccountConnectionAttemptGetResponseSuccess{},
		"failed":              XAccountConnectionAttemptGetResponseFailed{},
		"requires_email_code": XAccountConnectionAttemptGetResponseRequiresEmailCode{},
		"unknown":             nil,
	} {
		t.Run(status, func(t *testing.T) {
			raw := `{"status":"` + status + `","attemptId":"probe"}`
			var model XAccountConnectionAttemptGetResponseUnion
			if err := model.UnmarshalJSON([]byte(raw)); err != nil {
				t.Fatal(err)
			}
			variant := model.AsAny()
			if reflect.TypeOf(variant) != reflect.TypeOf(want) {
				t.Fatalf("response variant = %T, want %T", variant, want)
			}
			if variant != nil && variant.(interface{ RawJSON() string }).RawJSON() != raw {
				t.Fatal("response variant lost its original payload")
			}
		})
	}
}

func rawJSONReceiverNames(t *testing.T) map[string]struct{} {
	t.Helper()
	loaded, err := packages.Load(&packages.Config{Mode: packages.NeedName | packages.NeedTypes}, ".", "./shared", "./internal/apierror")
	if err != nil {
		t.Fatal(err)
	}
	if len(loaded) != 3 || packages.PrintErrors(loaded) != 0 {
		t.Fatal("incomplete model package inventory")
	}
	names := make(map[string]struct{})
	for _, pkg := range loaded {
		for _, name := range pkg.Types.Scope().Names() {
			object, ok := pkg.Types.Scope().Lookup(name).(*types.TypeName)
			if !ok || !object.Exported() {
				continue
			}
			if method := types.NewMethodSet(types.NewPointer(object.Type())).Lookup(pkg.Types, "RawJSON"); method != nil {
				names[name] = struct{}{}
			}
		}
	}
	return names
}
