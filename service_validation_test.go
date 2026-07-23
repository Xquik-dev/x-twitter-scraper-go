// Copyright the Xquik contributors.
// SPDX-License-Identifier: Apache-2.0

package xtwitterscraper_test

import (
	"context"
	"reflect"
	"strings"
	"testing"

	"github.com/Xquik-dev/x-twitter-scraper-go"
)

var contextType = reflect.TypeOf((*context.Context)(nil)).Elem()

func TestEveryServiceRejectsEmptyPathParameters(t *testing.T) {
	client := xtwitterscraper.NewClient()
	checked := validateServiceMethods(t, reflect.ValueOf(&client).Elem(), "Client")
	if checked < 60 {
		t.Fatalf("validated %d methods, want at least 60", checked)
	}
}

func validateServiceMethods(t *testing.T, service reflect.Value, path string) int {
	t.Helper()

	checked := 0
	for fieldIndex := 0; fieldIndex < service.NumField(); fieldIndex++ {
		field := service.Field(fieldIndex)
		fieldInfo := service.Type().Field(fieldIndex)
		if !fieldInfo.IsExported() || !strings.HasSuffix(field.Type().Name(), "Service") {
			continue
		}

		methods := field.Addr()
		for methodIndex := 0; methodIndex < methods.NumMethod(); methodIndex++ {
			method := methods.Method(methodIndex)
			methodInfo := methods.Type().Method(methodIndex)
			for argumentIndex := 0; argumentIndex < method.Type().NumIn(); argumentIndex++ {
				if method.Type().In(argumentIndex).Kind() != reflect.String {
					continue
				}
				checked++
				t.Run(path+"."+fieldInfo.Name+"."+methodInfo.Name, func(t *testing.T) {
					results := method.CallSlice(serviceMethodArguments(method.Type(), argumentIndex))
					err, ok := results[len(results)-1].Interface().(error)
					if !ok || err == nil || !strings.Contains(err.Error(), "missing required") {
						t.Fatalf("empty path parameter returned %v", results)
					}
				})
			}
		}

		checked += validateServiceMethods(t, field, path+"."+fieldInfo.Name)
	}
	return checked
}

func serviceMethodArguments(method reflect.Type, emptyStringIndex int) []reflect.Value {
	arguments := make([]reflect.Value, method.NumIn())
	for index := range arguments {
		argumentType := method.In(index)
		switch {
		case argumentType.Implements(contextType):
			arguments[index] = reflect.ValueOf(context.Background())
		case argumentType.Kind() == reflect.String && index != emptyStringIndex:
			arguments[index] = reflect.ValueOf("id").Convert(argumentType)
		default:
			arguments[index] = reflect.Zero(argumentType)
		}
	}
	return arguments
}
