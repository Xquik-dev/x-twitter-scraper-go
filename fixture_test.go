// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

package xtwitterscraper_test

import (
	"bytes"
	"embed"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"maps"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"slices"
	"strconv"
	"strings"
	"testing"

	"github.com/Xquik-dev/x-twitter-scraper-go"
	"github.com/Xquik-dev/x-twitter-scraper-go/option"
	"github.com/Xquik-dev/x-twitter-scraper-go/packages/respjson"
)

//go:embed testdata/fixtures
var fixtures embed.FS

type capturedRequest struct {
	method string
	path   string
	query  url.Values
	header http.Header
	body   []byte
}

func fixtureClient(t *testing.T, fixture string) (xtwitterscraper.Client, *capturedRequest) {
	t.Helper()
	payload, err := fixtures.ReadFile("testdata/fixtures/" + fixture)
	if err != nil {
		t.Fatal(err)
	}
	captured := &capturedRequest{}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			t.Error(err)
		}
		*captured = capturedRequest{method: r.Method, path: r.URL.Path, query: r.URL.Query(), header: r.Header.Clone(), body: body}
		w.Header().Set("Content-Type", "application/json")
		if _, err := w.Write(payload); err != nil {
			t.Error(err)
		}
	}))
	t.Cleanup(server.Close)
	client := xtwitterscraper.NewClient(
		option.WithBaseURL(server.URL+"/api/v1/"),
		option.WithAPIKey("My API Key"),
		option.WithMaxRetries(0),
	)
	return client, captured
}

func (r *capturedRequest) expect(t *testing.T, method, path, query string) {
	t.Helper()
	if r.method != method || r.path != path || r.query.Encode() != query {
		t.Fatalf("request = %s %s?%s, want %s %s?%s", r.method, r.path, r.query.Encode(), method, path, query)
	}
}

func (r *capturedRequest) expectBody(t *testing.T, want string) {
	t.Helper()
	got, err := canonicalJSON(r.body)
	if err != nil {
		t.Fatalf("request body %q is not JSON: %v", r.body, err)
	}
	expected, err := canonicalJSON([]byte(want))
	if err != nil {
		t.Fatal(err)
	}
	if got != expected {
		t.Fatalf("request body = %s, want %s", r.body, want)
	}
}

// canonicalJSON writes a JSON value with sorted object keys, no spaces and
// each number and string in one form, so 2 values that decode to the same data
// give the same text.
func canonicalJSON(raw []byte) (string, error) {
	raw = bytes.TrimSpace(raw)
	if len(raw) == 0 {
		return "", errors.New("empty JSON value")
	}
	switch first := raw[0]; {
	case first == '{':
		var object map[string]json.RawMessage
		if err := json.Unmarshal(raw, &object); err != nil {
			return "", err
		}
		parts := make([]string, 0, len(object))
		for _, key := range slices.Sorted(maps.Keys(object)) {
			name, err := json.Marshal(key)
			if err != nil {
				return "", err
			}
			value, err := canonicalJSON(object[key])
			if err != nil {
				return "", err
			}
			parts = append(parts, string(name)+":"+value)
		}
		return "{" + strings.Join(parts, ",") + "}", nil
	case first == '[':
		var items []json.RawMessage
		if err := json.Unmarshal(raw, &items); err != nil {
			return "", err
		}
		parts := make([]string, 0, len(items))
		for _, item := range items {
			value, err := canonicalJSON(item)
			if err != nil {
				return "", err
			}
			parts = append(parts, value)
		}
		return "[" + strings.Join(parts, ",") + "]", nil
	case first == '"':
		var text string
		if err := json.Unmarshal(raw, &text); err != nil {
			return "", err
		}
		encoded, err := json.Marshal(text)
		return string(encoded), err
	case first == '-' || first >= '0' && first <= '9':
		var number float64
		if err := json.Unmarshal(raw, &number); err != nil {
			return "", err
		}
		return strconv.FormatFloat(number, 'g', -1, 64), nil
	default:
		var literal bool
		if string(raw) == "null" {
			return "null", nil
		}
		if err := json.Unmarshal(raw, &literal); err != nil {
			return "", err
		}
		return strconv.FormatBool(literal), nil
	}
}

func expectDecoded[T any](t *testing.T, response T) {
	t.Helper()
	var problems []string
	walkDecoded(reflect.ValueOf(response), "res", nil, &problems)
	if len(problems) > 0 {
		t.Fatalf("the fixture did not decode into the response type\n%s", strings.Join(problems, "\n"))
	}
}

func jsonKeys(kind reflect.Type, keys map[string]bool) {
	for index := range kind.NumField() {
		field := kind.Field(index)
		if field.Anonymous && field.Type.Kind() == reflect.Struct {
			jsonKeys(field.Type, keys)
			continue
		}
		if name, _, _ := strings.Cut(field.Tag.Get("json"), ","); name != "" && name != "-" {
			keys[name] = true
		}
	}
}

func walkDecoded(value reflect.Value, path string, family map[string]bool, problems *[]string) {
	switch value.Kind() {
	case reflect.Pointer, reflect.Interface:
		if !value.IsNil() {
			walkDecoded(value.Elem(), path, nil, problems)
		}
	case reflect.Slice:
		for index := range value.Len() {
			walkDecoded(value.Index(index), fmt.Sprintf("%s[%d]", path, index), nil, problems)
		}
	case reflect.Map:
		for _, key := range value.MapKeys() {
			walkDecoded(value.MapIndex(key), fmt.Sprintf("%s[%v]", path, key), nil, problems)
		}
	case reflect.Struct:
		kind := value.Type()
		if strings.HasSuffix(kind.Name(), "Union") {
			return
		}
		keys := map[string]bool{}
		for key := range family {
			keys[key] = true
		}
		jsonKeys(kind, keys)
		if metadata := value.FieldByName("JSON"); metadata.IsValid() {
			checkMetadata(metadata, path, keys, problems)
		}
		for index := range kind.NumField() {
			field := kind.Field(index)
			switch {
			case field.Anonymous && field.Type.Kind() == reflect.Struct:
				walkDecoded(value.Field(index), path, keys, problems)
			case field.IsExported() && field.Name != "JSON":
				walkDecoded(value.Field(index), path+"."+field.Name, nil, problems)
			}
		}
	}
}

func checkMetadata(metadata reflect.Value, path string, keys map[string]bool, problems *[]string) {
	for index := range metadata.NumField() {
		field := metadata.Type().Field(index)
		switch {
		case !field.IsExported():
		case field.Name == "ExtraFields":
			for _, key := range metadata.Field(index).MapKeys() {
				if !keys[key.String()] {
					*problems = append(*problems, fmt.Sprintf("%s has the unknown key %v", path, key))
				}
			}
		default:
			status, ok := reflect.TypeAssert[respjson.Field](metadata.Field(index))
			if !ok {
				*problems = append(*problems, fmt.Sprintf("%s.JSON.%s is a %s", path, field.Name, field.Type))
				continue
			}
			if raw := status.Raw(); raw != "" && raw != respjson.Null && !status.Valid() {
				*problems = append(*problems, fmt.Sprintf("%s.%s cannot hold %s", path, field.Name, raw))
			}
		}
	}
}
