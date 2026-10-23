// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"bytes"
	"encoding/json"
	"io"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"testing/fstest"

	"github.com/Xquik-dev/x-twitter-scraper-go/internal/requestconfig"
	"github.com/Xquik-dev/x-twitter-scraper-go/option"
)

const serviceSource = `package sdk

func (r *XTweetService) Get(ctx context.Context, id string, query XTweetGetParams, opts ...option.RequestOption) (res *XTweetGetResponse, err error) {
	path := fmt.Sprintf("x/tweets/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

func (r *XTweetService) New(ctx context.Context, body XTweetNewParams, opts ...option.RequestOption) (res *shared.Action, err error) {
	path := "x/tweets"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

func (r *XTweetService) Export(ctx context.Context, id string, opts ...option.RequestOption) (res *http.Response, err error) {
	path := fmt.Sprintf("x/tweets/%s/export", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

func (r *XMediaService) Upload(ctx context.Context, body XMediaUploadParams, opts ...option.RequestOption) (res *UserProfile, err error) {
	path := "x/media"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

func (r *EventService) StreamStreaming(ctx context.Context, opts ...option.RequestOption) (stream *ssestream.Stream[Event]) {
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, "events/stream", nil, &raw, opts...)
	return
}

func (r *XLegacyService) Get(ctx context.Context, opts ...option.RequestOption) (res *UserProfile, err error) {
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, "x/legacy", nil, &res, opts...)
	return
}
`

const modelSource = `package sdk

type paramObj = param.APIObject

type UserProfile = shared.UserProfile

type XTweetGetParams struct {
	Q      string              ` + "`" + `query:"q" api:"required" json:"-"` + "`" + `
	Cursor param.Opt[string]   ` + "`" + `query:"cursor,omitzero" json:"-"` + "`" + `
	Mode   XTweetGetParamsMode ` + "`" + `query:"mode,omitzero" json:"-"` + "`" + `
	Lang   param.Opt[string]   ` + "`" + `header:"Accept-Language,omitzero" json:"-"` + "`" + `
	paramObj
}

type XTweetGetParamsMode string

const (
	XTweetGetParamsModeLatest XTweetGetParamsMode = "latest"
	XTweetGetParamsModeTop    XTweetGetParamsMode = "top"
)

type XTweetNewParams struct {
	Text  string   ` + "`" + `json:"text" api:"required"` + "`" + `
	Media []string ` + "`" + `json:"media,omitzero"` + "`" + `
	paramObj
}

type XMediaUploadParams struct {
	File io.Reader ` + "`" + `json:"file" api:"required" format:"binary"` + "`" + `
	paramObj
}

type XTweetGetResponse struct {
	Tweet  Tweet                        ` + "`" + `json:"tweet" api:"required"` + "`" + `
	Author UserProfile                  ` + "`" + `json:"author"` + "`" + `
	Media  []TweetMedia                 ` + "`" + `json:"media"` + "`" + `
	Extra  map[string]any               ` + "`" + `json:"extra"` + "`" + `
	Raw    any                          ` + "`" + `json:"raw"` + "`" + `
	Answer XTweetGetResponseAnswerUnion ` + "`" + `json:"answer"` + "`" + `
	Result XTweetGetResponseResultUnion ` + "`" + `json:"result"` + "`" + `
	JSON   struct {
		Tweet       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} ` + "`" + `json:"-"` + "`" + `
}

type Tweet struct {
	ID     string         ` + "`" + `json:"id" api:"required"` + "`" + `
	Text   string         ` + "`" + `json:"text"` + "`" + `
	Note   string         ` + "`" + `json:"note" api:"nullable"` + "`" + `
	Type   constant.Point ` + "`" + `json:"type" default:"Point"` + "`" + `
	Quoted *Tweet         ` + "`" + `json:"quoted"` + "`" + `
	Tags   []string       ` + "`" + `json:"tags"` + "`" + `
}

type TweetMedia struct {
	URL string ` + "`" + `json:"url" api:"required"` + "`" + `
}

type XTweetGetResponseAnswerUnion struct {
	OfFloat  float64 ` + "`" + `json:",inline"` + "`" + `
	OfString string  ` + "`" + `json:",inline"` + "`" + `
}

type XTweetGetResponseResultUnion struct {
	Kind  string  ` + "`" + `json:"kind"` + "`" + `
	Score float64 ` + "`" + `json:"score"` + "`" + `
}
`

const sharedSource = `package shared

type UserProfile struct {
	ID   string ` + "`" + `json:"id" api:"required"` + "`" + `
	Name string ` + "`" + `json:"name"` + "`" + `
	JSON struct {
		ID respjson.Field
	} ` + "`" + `json:"-"` + "`" + `
}

type Action struct {
	ID     string           ` + "`" + `json:"id" api:"required"` + "`" + `
	Status constant.Pending ` + "`" + `json:"status" default:"pending"` + "`" + `
}
`

const constantSource = `package constant

type Point string
type Pending string

func (c Point) Default() Point     { return "Point" }
func (c Pending) Default() Pending { return "pending" }
`

var productionBaseURL = requestconfig.WithDefaultBaseURL("https://xquik.com/api/v1/")

func sdkTree(accepted string, overrides map[string]string) fstest.MapFS {
	files := map[string]string{
		"service.go":                    serviceSource,
		"model.go":                      modelSource,
		"model_test.go":                 "package sdk\n\nfunc (r *XTestService) Get() { requestconfig.ExecuteNewRequest(ctx, method, path) }\n",
		"shared/shared.go":              sharedSource,
		"shared/constant/constants.go":  constantSource,
		"shared/constant/README.md":     "# Constants\n",
		acceptedFile:                    accepted,
		"internal/apijson/decode.go":    "package apijson\n",
		"examples/tweets/main.go":       "package main\n",
		"shared/constant/doc_test.go":   "package constant\n",
		"shared/constant/extra/skip.go": "package extra\n",
	}
	for name, content := range overrides {
		files[name] = content
	}
	tree := fstest.MapFS{}
	for name, content := range files {
		tree[name] = &fstest.MapFile{Data: []byte(content)}
	}
	return tree
}

func objectSchema(required []string, properties map[string]any) map[string]any {
	schema := map[string]any{"type": "object", "properties": properties}
	if required != nil {
		schema["required"] = required
	}
	return schema
}

func jsonContent(schema any) map[string]any {
	return map[string]any{"application/json": map[string]any{"schema": schema}}
}

func str() map[string]any { return map[string]any{"type": "string"} }

func baseSpec() map[string]any {
	tweet := objectSchema([]string{"id"}, map[string]any{
		"id":     str(),
		"text":   str(),
		"note":   map[string]any{"type": []any{"string", "null"}},
		"type":   map[string]any{"type": "string", "const": "Point"},
		"quoted": map[string]any{"$ref": "#/components/schemas/Tweet"},
		"tags":   map[string]any{"type": "array", "items": str()},
	})
	profile := objectSchema([]string{"id"}, map[string]any{"id": str(), "name": str()})
	response := objectSchema([]string{"tweet"}, map[string]any{
		"tweet":  map[string]any{"$ref": "#/components/schemas/Tweet", "description": "The tweet."},
		"author": map[string]any{"$ref": "#/components/schemas/UserProfile"},
		"media":  map[string]any{"type": "array", "items": objectSchema([]string{"url"}, map[string]any{"url": str()})},
		"extra":  map[string]any{"type": "object", "additionalProperties": true},
		"raw":    objectSchema(nil, map[string]any{"anything": str()}),
		"answer": map[string]any{"oneOf": []any{map[string]any{"type": "number"}, str(), map[string]any{"type": "null"}}},
		"result": map[string]any{"oneOf": []any{
			objectSchema([]string{"kind", "score"}, map[string]any{"kind": str(), "score": map[string]any{"type": "number"}}),
			objectSchema([]string{"kind"}, map[string]any{"kind": str()}),
		}},
	})
	return map[string]any{
		"openapi": "3.1.0",
		"servers": []any{map[string]any{"url": "https://xquik.com"}},
		"components": map[string]any{
			"schemas":    map[string]any{"Tweet": tweet, "UserProfile": profile},
			"parameters": map[string]any{"Cursor": map[string]any{"name": "cursor", "in": "query", "schema": str()}},
		},
		"paths": map[string]any{
			"/api/v1/x/tweets/{tweetId}": map[string]any{
				"parameters": []any{map[string]any{"name": "tweetId", "in": "path", "required": true, "schema": str()}},
				"get": map[string]any{
					"operationId": "getTweet",
					"parameters": []any{
						map[string]any{"name": "q", "in": "query", "required": true, "schema": str()},
						map[string]any{"$ref": "#/components/parameters/Cursor"},
						map[string]any{"name": "mode", "in": "query", "schema": map[string]any{"type": "string", "enum": []any{"latest", "top"}}},
						map[string]any{"name": "Accept-Language", "in": "header", "schema": str()},
					},
					"responses": map[string]any{
						"200":     map[string]any{"description": "OK", "content": jsonContent(response)},
						"404":     map[string]any{"description": "Missing", "content": jsonContent(objectSchema(nil, map[string]any{"error": str()}))},
						"default": map[string]any{"description": "Error"},
					},
				},
			},
			"/api/v1/x/tweets": map[string]any{
				"post": map[string]any{
					"requestBody": map[string]any{"content": map[string]any{
						"application/json":    map[string]any{"schema": objectSchema([]string{"text"}, map[string]any{"text": str(), "media": map[string]any{"type": "array", "items": str()}})},
						"multipart/form-data": map[string]any{"schema": objectSchema([]string{"text"}, map[string]any{"text": str(), "media": map[string]any{"type": "array", "items": str()}})},
					}},
					"responses": map[string]any{"201": map[string]any{"description": "Created", "content": jsonContent(objectSchema([]string{"id"}, map[string]any{"id": str(), "status": map[string]any{"type": "string", "enum": []any{"pending"}}}))}},
				},
			},
			"/api/v1/x/tweets/{id}/export": map[string]any{
				"get": map[string]any{"responses": map[string]any{"200": map[string]any{"description": "CSV", "content": map[string]any{"text/csv": map[string]any{"schema": str()}}}}},
			},
			"/api/v1/x/media": map[string]any{
				"post": map[string]any{
					"requestBody": map[string]any{"content": map[string]any{"multipart/form-data": map[string]any{"schema": objectSchema([]string{"file"}, map[string]any{"file": map[string]any{"type": "string", "format": "binary"}})}}},
					"responses":   map[string]any{"200": map[string]any{"description": "OK", "content": jsonContent(map[string]any{"$ref": "#/components/schemas/UserProfile"})}},
				},
			},
			"/api/v1/events/stream": map[string]any{
				"get": map[string]any{"responses": map[string]any{"200": map[string]any{"description": "Stream", "content": map[string]any{"text/event-stream": map[string]any{"schema": str()}}}}},
			},
		},
	}
}

func path(spec map[string]any, keys ...string) map[string]any {
	current := spec
	for _, key := range keys {
		current = current[key].(map[string]any)
	}
	return current
}

func lookupResponse(spec map[string]any) map[string]any {
	return path(spec, "paths", "/api/v1/x/tweets/{tweetId}", "get", "responses", "200", "content", "application/json", "schema", "properties")
}

func lookupParameters(spec map[string]any) []any {
	return path(spec, "paths", "/api/v1/x/tweets/{tweetId}", "get")["parameters"].([]any)
}

func tweetSchema(spec map[string]any) map[string]any {
	return path(spec, "components", "schemas", "Tweet")
}

func writeSpec(t *testing.T, spec any) string {
	t.Helper()
	data, err := json.Marshal(spec)
	if err != nil {
		t.Fatal(err)
	}
	file := filepath.Join(t.TempDir(), "openapi.json")
	if err := os.WriteFile(file, data, 0o600); err != nil {
		t.Fatal(err)
	}
	return file
}

func runWith(t *testing.T, root fs.FS, baseURL option.RequestOption, source string) (int, string, string) {
	t.Helper()
	var out, failures bytes.Buffer
	code := run(source, root, baseURL, http.DefaultClient, &out, &failures)
	return code, out.String(), failures.String()
}

func expectFailures(t *testing.T, root fs.FS, spec any, want ...string) {
	t.Helper()
	code, out, failures := runWith(t, root, productionBaseURL, writeSpec(t, spec))
	if expected := strings.Join(want, "\n") + "\n"; code != 1 || out != "" || failures != expected {
		t.Fatalf("code = %d, out = %q\nfailures:\n%s\nwant:\n%s", code, out, failures, expected)
	}
}

func expectPass(t *testing.T, root fs.FS, spec any, want string) {
	t.Helper()
	code, out, failures := runWith(t, root, productionBaseURL, writeSpec(t, spec))
	if code != 0 || out != want+"\n" || failures != "" {
		t.Fatalf("code = %d, out = %q, failures = %q", code, out, failures)
	}
}

func TestMatchingFieldsPass(t *testing.T) {
	expectPass(t, sdkTree("[]", nil), baseSpec(), "The SDK's fields match the API for 5 operations, and accepted-field-differences.json lists 0 accepted differences.")
}

func TestRequestBodyMediaTypes(t *testing.T) {
	const service = `package sdk

func (r *SupportTicketService) New(ctx context.Context, body SupportTicketNewParams, opts ...option.RequestOption) (res *http.Response, err error) {
	path := "support/tickets"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}
`
	const jsonModel = `package sdk

type SupportTicketNewParams struct {
	Body string ` + "`" + `json:"body" api:"required"` + "`" + `
	paramObj
}

type paramObj = param.APIObject
`
	const mergedModel = `package sdk

type SupportTicketNewParams struct {
	Body        param.Opt[string] ` + "`" + `json:"body,omitzero"` + "`" + `
	Attachments []io.Reader       ` + "`" + `json:"attachments,omitzero"` + "`" + `
	paramObj
}

type paramObj = param.APIObject
`
	for _, mediaType := range []string{"multipart/form-data", "application/x-www-form-urlencoded"} {
		t.Run(mediaType, func(t *testing.T) {
			spec := map[string]any{
				"openapi": "3.1.0",
				"paths": map[string]any{
					"/api/v1/support/tickets": map[string]any{
						"post": map[string]any{
							"requestBody": map[string]any{"content": map[string]any{
								"application/json": map[string]any{"schema": objectSchema([]string{"body"}, map[string]any{"body": str()})},
								mediaType: map[string]any{"schema": objectSchema(nil, map[string]any{
									"body":        str(),
									"attachments": map[string]any{"type": "array", "items": map[string]any{"type": "string", "format": "binary"}},
								})},
							}},
						},
					},
				},
			}
			t.Run("JSON-only SDK", func(t *testing.T) {
				expectFailures(t, sdkTree("[]", map[string]string{"service.go": service, "model.go": jsonModel}), spec,
					"POST /api/v1/support/tickets request missing attachments",
					"POST /api/v1/support/tickets request required body",
					"Fix each field above in the SDK, or list it in accepted-field-differences.json with a reason.")
			})
			t.Run("matching SDK", func(t *testing.T) {
				expectPass(t, sdkTree("[]", map[string]string{"service.go": service, "model.go": mergedModel}), spec,
					"The SDK's fields match the API for 1 operation, and accepted-field-differences.json lists 0 accepted differences.")
			})
		})
	}
}

func TestDifferencesFail(t *testing.T) {
	const next = "Fix each field above in the SDK, or list it in accepted-field-differences.json with a reason."
	tests := []struct {
		name   string
		change func(map[string]any)
		want   []string
	}{
		{"missing query param", func(spec map[string]any) {
			path(spec, "paths", "/api/v1/x/tweets/{tweetId}", "get")["parameters"] = append(lookupParameters(spec), map[string]any{"name": "limit", "in": "query", "schema": map[string]any{"type": "integer"}})
		}, []string{"GET /api/v1/x/tweets/{tweetId} request missing limit"}},
		{"extra query param", func(spec map[string]any) {
			path(spec, "paths", "/api/v1/x/tweets/{tweetId}", "get")["parameters"] = lookupParameters(spec)[:1]
		}, []string{"GET /api/v1/x/tweets/{tweetId} request extra cursor", "GET /api/v1/x/tweets/{tweetId} request extra mode"}},
		{"query enum", func(spec map[string]any) {
			lookupParameters(spec)[2].(map[string]any)["schema"] = map[string]any{"type": "string", "enum": []any{"latest", "top", "live"}}
		}, []string{"GET /api/v1/x/tweets/{tweetId} request enum mode"}},
		{"query required", func(spec map[string]any) {
			lookupParameters(spec)[0].(map[string]any)["required"] = false
		}, []string{"GET /api/v1/x/tweets/{tweetId} request required q"}},
		{"nullable", func(spec map[string]any) {
			path(tweetSchema(spec), "properties")["note"] = str()
		}, []string{"GET /api/v1/x/tweets/{tweetId} response nullable tweet.note"}},
		{"response required", func(spec map[string]any) {
			tweetSchema(spec)["required"] = []any{}
		}, []string{"GET /api/v1/x/tweets/{tweetId} response required tweet.id"}},
		{"missing nested object counts once", func(spec map[string]any) {
			path(tweetSchema(spec), "properties")["entities"] = objectSchema(nil, map[string]any{"urls": map[string]any{"type": "array", "items": objectSchema(nil, map[string]any{"url": str()})}})
		}, []string{"GET /api/v1/x/tweets/{tweetId} response missing tweet.entities"}},
		{"missing array item field", func(spec map[string]any) {
			path(lookupResponse(spec), "media", "items", "properties")["alt"] = str()
		}, []string{"GET /api/v1/x/tweets/{tweetId} response missing media[].alt"}},
		{"extra response field", func(spec map[string]any) {
			delete(path(tweetSchema(spec), "properties"), "tags")
		}, []string{"GET /api/v1/x/tweets/{tweetId} response extra tweet.tags"}},
		{"another kind counts as missing", func(spec map[string]any) {
			path(tweetSchema(spec), "properties")["text"] = objectSchema(nil, map[string]any{"value": str()})
			path(tweetSchema(spec), "properties")["tags"] = str()
		}, []string{"GET /api/v1/x/tweets/{tweetId} response missing tweet.tags", "GET /api/v1/x/tweets/{tweetId} response missing tweet.text"}},
		{"response enum", func(spec map[string]any) {
			path(spec, "paths", "/api/v1/x/tweets", "post", "responses", "201", "content", "application/json", "schema", "properties")["status"] = map[string]any{"type": "string", "enum": []any{"pending", "sent"}}
		}, []string{"POST /api/v1/x/tweets response enum status"}},
		{"several 2xx bodies merge", func(spec map[string]any) {
			path(spec, "paths", "/api/v1/x/tweets", "post", "responses")["202"] = map[string]any{"description": "Queued", "content": jsonContent(objectSchema([]string{"jobId"}, map[string]any{"jobId": str()}))}
		}, []string{"POST /api/v1/x/tweets response missing jobId", "POST /api/v1/x/tweets response required id"}},
		{"untyped JSON response", func(spec map[string]any) {
			path(spec, "paths", "/api/v1/x/tweets/{id}/export", "get", "responses", "200")["content"] = jsonContent(objectSchema(nil, map[string]any{"rows": str(), "total": str()}))
		}, []string{"GET /api/v1/x/tweets/{id}/export response missing rows", "GET /api/v1/x/tweets/{id}/export response missing total"}},
		{"untyped scalar response", func(spec map[string]any) {
			path(spec, "paths", "/api/v1/x/tweets/{id}/export", "get", "responses", "200")["content"] = jsonContent(str())
		}, []string{"GET /api/v1/x/tweets/{id}/export response missing (body)"}},
		{"response root of another kind", func(spec map[string]any) {
			path(spec, "paths", "/api/v1/x/media", "post", "responses", "200")["content"] = jsonContent(map[string]any{"type": "array", "items": str()})
		}, []string{"POST /api/v1/x/media response missing (body)"}},
		{"missing body field", func(spec map[string]any) {
			path(spec, "paths", "/api/v1/x/tweets", "post", "requestBody", "content", "application/json", "schema", "properties")["replyTo"] = str()
		}, []string{"POST /api/v1/x/tweets request missing replyTo"}},
		{"body the SDK does not send", func(spec map[string]any) {
			path(spec, "paths", "/api/v1/x/tweets/{id}/export", "get")["requestBody"] = map[string]any{"content": jsonContent(objectSchema(nil, map[string]any{"format": str()}))}
		}, []string{"GET /api/v1/x/tweets/{id}/export request missing format"}},
		{"body the spec does not take", func(spec map[string]any) {
			delete(path(spec, "paths", "/api/v1/x/tweets", "post"), "requestBody")
		}, []string{"POST /api/v1/x/tweets request extra media", "POST /api/v1/x/tweets request extra text"}},
		{"array body the SDK does not send", func(spec map[string]any) {
			path(spec, "paths", "/api/v1/x/tweets/{id}/export", "get")["requestBody"] = map[string]any{"content": jsonContent(map[string]any{"type": "array", "items": str()})}
		}, []string{"GET /api/v1/x/tweets/{id}/export request missing (body)"}},
		{"multipart body when there is no JSON body", func(spec map[string]any) {
			path(spec, "paths", "/api/v1/x/media", "post", "requestBody", "content", "multipart/form-data", "schema", "properties")["alt"] = str()
		}, []string{"POST /api/v1/x/media request missing alt"}},
		{"closed object rejects extra SDK fields", func(spec map[string]any) {
			delete(path(tweetSchema(spec), "properties"), "tags")
			tweetSchema(spec)["additionalProperties"] = false
		}, []string{"GET /api/v1/x/tweets/{tweetId} response extra tweet.tags"}},
		{"typed response the spec does not document", func(spec map[string]any) {
			path(spec, "paths", "/api/v1/x/media", "post", "responses", "200")["content"] = map[string]any{"text/csv": map[string]any{"schema": str()}}
		}, []string{"POST /api/v1/x/media response extra id", "POST /api/v1/x/media response extra name"}},
		{"vendor JSON response", func(spec map[string]any) {
			responses := path(spec, "paths", "/api/v1/x/media", "post", "responses", "200")
			schema := objectSchema([]string{"id"}, map[string]any{"id": str(), "name": str(), "bio": str()})
			responses["content"] = map[string]any{"application/vnd.xquik+json; charset=utf-8": map[string]any{"schema": schema}}
		}, []string{"POST /api/v1/x/media response missing bio"}},
		{"union field required only when every variant requires it", func(spec map[string]any) {
			path(lookupResponse(spec), "result")["oneOf"] = []any{objectSchema([]string{"kind"}, map[string]any{"kind": str(), "label": str()})}
		}, []string{"GET /api/v1/x/tweets/{tweetId} response extra result.score", "GET /api/v1/x/tweets/{tweetId} response missing result.label"}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			spec := baseSpec()
			test.change(spec)
			expectFailures(t, sdkTree("[]", nil), spec, append(test.want, next)...)
		})
	}
}

func TestMapValueFields(t *testing.T) {
	const next = "Fix each field above in the SDK, or list it in accepted-field-differences.json with a reason."
	typed := sdkTree("[]", map[string]string{"model.go": strings.Replace(modelSource, "Extra  map[string]any ", "Extra  map[string]TweetMedia ", 1)})
	values := func(properties map[string]any) map[string]any {
		spec := baseSpec()
		lookupResponse(spec)["extra"] = map[string]any{"type": "object", "additionalProperties": objectSchema([]string{"url"}, properties)}
		return spec
	}
	t.Run("missing value field", func(t *testing.T) {
		expectFailures(t, typed, values(map[string]any{"url": str(), "alt": str()}), "GET /api/v1/x/tweets/{tweetId} response missing extra{}.alt", next)
	})
	t.Run("extra value field", func(t *testing.T) {
		expectFailures(t, typed, values(map[string]any{}), "GET /api/v1/x/tweets/{tweetId} response extra extra{}.url", next)
	})
	for keyword, combine := range map[string]string{"allOf": "intersected", "oneOf": "merged"} {
		t.Run(combine+" value fields", func(t *testing.T) {
			spec := baseSpec()
			mapOf := func(required []string, properties map[string]any) map[string]any {
				return map[string]any{"type": "object", "additionalProperties": objectSchema(required, properties)}
			}
			lookupResponse(spec)["extra"] = map[string]any{keyword: []any{mapOf([]string{"url"}, map[string]any{"url": str()}), mapOf(nil, map[string]any{"alt": str()})}}
			expectFailures(t, typed, spec, "GET /api/v1/x/tweets/{tweetId} response missing extra{}.alt", next)
		})
	}
	t.Run("matching value fields", func(t *testing.T) {
		code, _, failures := runWith(t, typed, productionBaseURL, writeSpec(t, values(map[string]any{"url": str()})))
		if code != 0 || failures != "" {
			t.Fatalf("code = %d, failures = %q", code, failures)
		}
	})
}

func TestSchemaRulesKeepMatchingFieldsQuiet(t *testing.T) {
	const summary = "The SDK's fields match the API for 5 operations, and accepted-field-differences.json lists 0 accepted differences."
	tests := []struct {
		name   string
		change func(map[string]any)
	}{
		{"open object allows extra SDK fields", func(spec map[string]any) {
			delete(path(tweetSchema(spec), "properties"), "tags")
			tweetSchema(spec)["additionalProperties"] = map[string]any{"type": "string"}
		}},
		{"allOf merges parts and a required list", func(spec map[string]any) {
			schemas := path(spec, "components", "schemas")
			profile := schemas["UserProfile"].(map[string]any)
			delete(profile, "required")
			schemas["UserProfile"] = map[string]any{"allOf": []any{profile, map[string]any{"required": []any{"id"}}}}
		}},
		{"required list beside allOf", func(spec map[string]any) {
			schemas := path(spec, "components", "schemas")
			profile := schemas["UserProfile"].(map[string]any)
			delete(profile, "required")
			schemas["UserProfile"] = map[string]any{"allOf": []any{profile}, "required": []any{"id"}}
		}},
		{"allOf of arrays intersects items", func(spec map[string]any) {
			media := lookupResponse(spec)["media"].(map[string]any)
			lookupResponse(spec)["media"] = map[string]any{"allOf": []any{media, map[string]any{"type": "array", "items": map[string]any{"required": []any{"url"}}}}}
			path(media, "items")["required"] = []any{}
		}},
		{"ref with siblings", func(spec map[string]any) {
			lookupResponse(spec)["author"] = map[string]any{"$ref": "#/components/schemas/UserProfile", "type": []any{"object", "null"}, "description": "Author."}
		}},
		{"ref with schema siblings", func(spec map[string]any) {
			lookupResponse(spec)["author"] = map[string]any{"$ref": "#/components/schemas/UserProfile", "properties": map[string]any{"name": str()}}
		}},
		{"oneOf with shared properties", func(spec map[string]any) {
			lookupResponse(spec)["result"] = map[string]any{
				"properties": map[string]any{"kind": str()},
				"required":   []any{"kind"},
				"anyOf":      []any{objectSchema(nil, map[string]any{"score": map[string]any{"type": "number"}}), map[string]any{"enum": []any{nil}}, objectSchema(nil, map[string]any{})},
			}
		}},
		{"oneOf of 1 object", func(spec map[string]any) {
			lookupResponse(spec)["author"] = map[string]any{"oneOf": []any{map[string]any{"$ref": "#/components/schemas/UserProfile"}, map[string]any{"type": "null"}}}
		}},
		{"oneOf of only null", func(spec map[string]any) {
			lookupResponse(spec)["raw"] = map[string]any{"oneOf": []any{map[string]any{"type": "null"}}}
		}},
		{"null type and nullable flag", func(spec map[string]any) {
			lookupResponse(spec)["raw"] = map[string]any{"type": "null"}
			path(tweetSchema(spec), "properties")["note"] = map[string]any{"type": "string", "nullable": true}
		}},
		{"enum with null", func(spec map[string]any) {
			path(tweetSchema(spec), "properties")["note"] = map[string]any{"type": "string", "enum": []any{"a", nil}}
		}},
		{"free-form and tuple shapes stop the walk", func(spec map[string]any) {
			lookupResponse(spec)["media"] = map[string]any{"type": "object"}
			path(tweetSchema(spec), "properties")["tags"] = map[string]any{"type": "array", "prefixItems": []any{str()}}
			lookupResponse(spec)["answer"] = map[string]any{"anyOf": []any{map[string]any{"type": "object"}, str()}}
		}},
		{"non-object schemas stay unknown", func(spec map[string]any) {
			lookupResponse(spec)["raw"] = true
			path(tweetSchema(spec), "properties")["tags"] = map[string]any{"oneOf": []any{map[string]any{"type": "array", "items": str()}, map[string]any{"type": "array", "items": map[string]any{"type": "integer"}}}}
		}},
		{"allOf requires a field any part requires", func(spec map[string]any) {
			schemas := path(spec, "components", "schemas")
			schemas["UserProfile"] = map[string]any{"allOf": []any{
				objectSchema([]string{"id"}, map[string]any{"id": str(), "name": str()}),
				objectSchema(nil, map[string]any{"id": str()}),
			}}
		}},
		{"allOf of 2 unions distributes over both", func(spec map[string]any) {
			lookupResponse(spec)["result"] = map[string]any{"allOf": []any{
				map[string]any{"oneOf": []any{objectSchema([]string{"kind"}, map[string]any{"kind": str()}), objectSchema(nil, map[string]any{"kind": str()})}},
				map[string]any{"oneOf": []any{objectSchema(nil, map[string]any{"score": map[string]any{"type": "number"}}), objectSchema(nil, map[string]any{})}},
			}}
		}},
		{"oneOf beside allOf", func(spec map[string]any) {
			lookupResponse(spec)["result"] = map[string]any{
				"allOf": []any{objectSchema([]string{"kind"}, map[string]any{"kind": str()})},
				"oneOf": []any{objectSchema(nil, map[string]any{"score": map[string]any{"type": "number"}}), objectSchema(nil, map[string]any{})},
			}
		}},
		{"null type beside a ref with other siblings", func(spec map[string]any) {
			path(spec, "components", "schemas")["Note"] = str()
			path(tweetSchema(spec), "properties")["note"] = map[string]any{"$ref": "#/components/schemas/Note", "type": []any{"string", "null"}, "format": "text"}
		}},
		{"operation params override path item params", func(spec map[string]any) {
			item := path(spec, "paths", "/api/v1/x/tweets/{tweetId}")
			item["parameters"] = append(item["parameters"].([]any), map[string]any{"name": "q", "in": "query", "required": false, "schema": str()})
		}},
		{"numeric and boolean enums", func(spec map[string]any) {
			lookupResponse(spec)["answer"] = map[string]any{"oneOf": []any{map[string]any{"enum": []any{1.5, true}}, map[string]any{"enum": []any{map[string]any{}}}}}
		}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			spec := baseSpec()
			test.change(spec)
			expectPass(t, sdkTree("[]", nil), spec, summary)
		})
	}
}

func TestAcceptedDifferences(t *testing.T) {
	spec := baseSpec()
	path(spec, "paths", "/api/v1/x/tweets/{tweetId}", "get")["parameters"] = lookupParameters(spec)[:2]
	accepted := `[{"reason": "Kept so existing calls still compile.", "differences": ["GET /api/v1/x/tweets/{id} request extra mode"]}]`
	expectPass(t, sdkTree(accepted, nil), spec, "The SDK's fields match the API for 5 operations, and accepted-field-differences.json lists 1 accepted difference.")

	stale := `[{"reason": "Kept so existing calls still compile.", "differences": ["GET /api/v1/x/tweets/{id} request extra mode", "GET /api/v1/x/tweets/{id} response missing tweet.gone"]}]`
	expectFailures(t, sdkTree(stale, nil), spec,
		"accepted-field-differences.json lists GET /api/v1/x/tweets/{id} response missing tweet.gone, which the check no longer finds. Remove it.")

	both := `[{"reason": "Kept.", "differences": ["POST /api/v1/x/tweets request extra text"]}]`
	expectFailures(t, sdkTree(both, nil), spec,
		"GET /api/v1/x/tweets/{tweetId} request extra mode",
		"Fix each field above in the SDK, or list it in accepted-field-differences.json with a reason.",
		"accepted-field-differences.json lists POST /api/v1/x/tweets request extra text, which the check no longer finds. Remove it.")
}

func TestInvalidAcceptedFilesFail(t *testing.T) {
	shape := "an upper case HTTP method, a path starting with /, request or response, missing, extra, enum, required or nullable, and a wire path, separated by spaces"
	tests := map[string]struct {
		content string
		want    string
	}{
		"not JSON":       {"[}", "It is not valid JSON near byte 2."},
		"wrong key":      {`[{"reason": "x", "operations": ["GET /x request extra y"]}]`, "It must be a JSON list of groups, each with only a reason string and a list of strings named differences."},
		"wrong kind":     {`[{"reason": "x", "differences": ["GET /x request typo y"]}]`, `Group 1 lists "GET /x request typo y", which is not ` + shape + "."},
		"unknown method": {`[{"reason": "x", "differences": ["FETCH /x request extra y"]}]`, `Group 1 lists "FETCH /x request extra y", which is not ` + shape + "."},
		"duplicate":      {`[{"reason": "x", "differences": ["GET /x/{a} request extra y", "GET /x/{b} request extra y"]}]`, "It lists GET /x/{a} request extra y and GET /x/{b} request extra y, which name the same difference. List it once."},
		"empty group":    {`[{"reason": "x", "differences": []}]`, "Group 1 needs at least 1 difference."},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			expectFailures(t, sdkTree(test.content, nil), baseSpec(), "accepted-field-differences.json is not valid. "+test.want)
		})
	}
	root := sdkTree("[]", nil)
	delete(root, acceptedFile)
	expectFailures(t, root, baseSpec(), "accepted-field-differences.json is not valid. The repository root has no readable copy. Create it, even as an empty list [].")
}

func TestUnreadableInputsFail(t *testing.T) {
	missing := filepath.Join(t.TempDir(), "missing.json")
	code, out, failures := runWith(t, sdkTree("[]", nil), productionBaseURL, missing)
	want := "Could not read the API spec from " + missing + ". Check network access, or set XQUIK_OPENAPI_SPEC to a local copy of https://xquik.com/openapi.json.\n"
	if code != 1 || out != "" || failures != want {
		t.Fatalf("code = %d, out = %q, failures = %q", code, out, failures)
	}

	spec := baseSpec()
	lookupResponse(spec)["author"] = map[string]any{"$ref": "#/components/schemas/Missing"}
	expectFailures(t, sdkTree("[]", nil), spec, "The field check cannot resolve #/components/schemas/Missing in the API spec. Fix the reference in the spec.")

	broken := strings.Replace(modelSource, "Media  []TweetMedia ", "Media  []TweetMedium", 1)
	broken = strings.Replace(broken, "Extra  map[string]any ", "Extra  chan int       ", 1)
	broken = strings.Replace(broken, "Raw    any  ", "Raw    json.RawMessage", 1)
	broken += "\ntype Other struct {\n\tValue *param.Field[string] `json:\"value\"`\n\tOdd (string) `json:\"odd\"`\n\tPicked pkg.Generic[string] `json:\"picked\"`\n}\n"
	broken = strings.Replace(broken, "Result XTweetGetResponseResultUnion ", "Result Other                        ", 1)
	expectFailures(t, sdkTree("[]", map[string]string{"model.go": broken}), baseSpec(),
		"The field check cannot resolve the Go type (string). Use a type from the SDK, shared, shared/constant, time or io packages.",
		"The field check cannot resolve the Go type TweetMedium. Use a type from the SDK, shared, shared/constant, time or io packages.",
		"The field check cannot resolve the Go type chan int. Use a type from the SDK, shared, shared/constant, time or io packages.",
		"The field check cannot resolve the Go type json.RawMessage. Use a type from the SDK, shared, shared/constant, time or io packages.",
		"The field check cannot resolve the Go type param.Field[string]. Use a type from the SDK, shared, shared/constant, time or io packages.",
		"The field check cannot resolve the Go type pkg.Generic[string]. Use a type from the SDK, shared, shared/constant, time or io packages.")

	params := strings.Replace(serviceSource, "path, query, &res", "path, query.Clean(), &res", 1)
	params = strings.Replace(params, "path, body, &res, opts...)\n\treturn\n}\n\nfunc (r *XTweetService) Export", "path, other, &res, opts...)\n\treturn\n}\n\nfunc (r *XTweetService) Export", 1)
	expectFailures(t, sdkTree("[]", map[string]string{"service.go": params}), baseSpec(),
		"service.go:11 passes request params whose type the field check cannot read. Pass the method's own params argument.",
		"service.go:5 passes request params whose type the field check cannot read. Pass the method's own params argument.")

	unreadable := strings.Replace(serviceSource, `"x/legacy"`, "legacyPath()", 1)
	expectFailures(t, sdkTree("[]", map[string]string{"service.go": unreadable}), baseSpec(),
		"service.go:33 sends a request whose method or path the field check cannot read. Use an http.Method constant and a path string or fmt.Sprintf format.")

	for name, root := range map[string]fs.FS{
		"syntax error":   sdkTree("[]", map[string]string{"shared/shared.go": "package shared\n\ntype {"}),
		"no shared":      fstest.MapFS{"service.go": {Data: []byte(serviceSource)}, acceptedFile: {Data: []byte("[]")}},
		"no root listed": fstest.MapFS{acceptedFile: {Data: []byte("[]")}, "nested/file.go": {Data: []byte("package nested\n")}},
	} {
		t.Run(name, func(t *testing.T) {
			code, out, failures := runWith(t, root, productionBaseURL, writeSpec(t, baseSpec()))
			if code != 1 || out != "" || !strings.HasPrefix(failures, "The field check could not") {
				t.Fatalf("code = %d, out = %q, failures = %q", code, out, failures)
			}
		})
	}

	code, out, failures = runWith(t, sdkTree("[]", nil), requestconfig.WithDefaultBaseURL("://bad"), writeSpec(t, baseSpec()))
	if code != 1 || out != "" || failures != "The field check could not read the default base URL from option.WithEnvironmentProduction. Fix that option.\n" {
		t.Fatalf("code = %d, out = %q, failures = %q", code, out, failures)
	}
}

func TestMainChecksTheWorkingDirectory(t *testing.T) {
	root := t.TempDir()
	for name, file := range sdkTree("[]", nil) {
		target := filepath.Join(root, filepath.FromSlash(name))
		if err := os.MkdirAll(filepath.Dir(target), 0o700); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(target, file.Data, 0o600); err != nil {
			t.Fatal(err)
		}
	}
	t.Setenv("XQUIK_OPENAPI_SPEC", writeSpec(t, baseSpec()))
	t.Chdir(root)
	output, err := os.CreateTemp(t.TempDir(), "stdout")
	if err != nil {
		t.Fatal(err)
	}
	defer output.Close()
	stdout := os.Stdout
	os.Stdout = output
	main()
	os.Stdout = stdout
	if _, err := output.Seek(0, io.SeekStart); err != nil {
		t.Fatal(err)
	}
	got, err := io.ReadAll(output)
	if err != nil {
		t.Fatal(err)
	}
	if want := "The SDK's fields match the API for 5 operations, and accepted-field-differences.json lists 0 accepted differences.\n"; string(got) != want {
		t.Fatalf("main printed %q, want %q", got, want)
	}
}

const edgeService = `package sdk

func (r *XEdgeService) Get(ctx context.Context, opts ...option.RequestOption) (res *EdgeResponse, err error) {
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, "x/edge", nil, &res, opts...)
	return
}

func (r *XEdgeService) New(ctx context.Context, body EdgeNewParams, opts ...option.RequestOption) (res *EdgeResponse, err error) {
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, "x/edge", body, &res, opts...)
	return
}
`

const edgeModel = `package sdk

type paramObj = param.APIObject

type EdgeKind int

const plainLimit = 1

const (
	EdgeKindFirst EdgeKind = iota
	EdgeKindLimit EdgeKind = 10
)

var EdgeKindDefault EdgeKind = 1

func (c *EdgeKind) Default() EdgeKind { return 1 }
func (c EdgeKind) Default() EdgeKind  { return c }

type EdgeResponse struct {
	Raw     interface{} ` + "`" + `json:"raw"` + "`" + `
	Inline  struct {
		A string ` + "`" + `json:"a"` + "`" + `
	} ` + "`" + `json:"inline"` + "`" + `
	Variant EdgeVariantUnion ` + "`" + `json:"variant"` + "`" + `
	Blank   string           ` + "`" + `json:"blank"` + "`" + `
	Req     string           ` + "`" + `json:"req"` + "`" + `
	EdgeEmbedded
	EdgeSkipped ` + "`" + `json:"skipped"` + "`" + `
	EdgeChoiceUnion
}

type EdgeEmbedded struct {
	Shared string ` + "`" + `json:"shared"` + "`" + `
	Blank  string ` + "`" + `json:"blank"` + "`" + `
}

type EdgeSkipped struct {
	Hidden string ` + "`" + `json:"hidden"` + "`" + `
}

type EdgeChoiceUnion struct {
	OfString string  ` + "`" + `json:",inline"` + "`" + `
	OfFloat  float64 ` + "`" + `json:",inline"` + "`" + `
}

type EdgeVariantUnion struct {
	Kind string ` + "`" + `json:"kind"` + "`" + `
	X    string ` + "`" + `json:"x"` + "`" + `
	Y    string ` + "`" + `json:"y"` + "`" + `
}

type EdgeNewParams struct {
	OfFirst  *EdgeNewParamsFirst  ` + "`" + `json:",inline"` + "`" + `
	OfSecond *EdgeNewParamsSecond ` + "`" + `json:",inline"` + "`" + `
	paramObj
}

type EdgeNewParamsFirst struct {
	Kind constant.Point ` + "`" + `json:"kind" default:"Point"` + "`" + `
	Note string         ` + "`" + `json:"note" api:"nullable"` + "`" + `
	paramObj
}

type EdgeNewParamsSecond struct {
	Kind constant.Point ` + "`" + `json:"kind" default:"Point"` + "`" + `
	Note string         ` + "`" + `json:"note"` + "`" + `
	paramObj
}
`

func edgeSpec() map[string]any {
	response := objectSchema(nil, map[string]any{
		"raw":    objectSchema(nil, map[string]any{"z": str()}),
		"inline": objectSchema(nil, map[string]any{"a": str()}),
		"variant": map[string]any{"allOf": []any{
			objectSchema([]string{"kind"}, map[string]any{"kind": str()}),
			map[string]any{"oneOf": []any{objectSchema(nil, map[string]any{"x": str()}), objectSchema(nil, map[string]any{"y": str()})}},
		}},
		"blank":   map[string]any{"oneOf": []any{true, true}},
		"req":     map[string]any{"allOf": []any{map[string]any{"required": []any{"a"}}}},
		"shared":  str(),
		"skipped": objectSchema(nil, map[string]any{"hidden": str()}),
	})
	point := map[string]any{"type": "string", "const": "Point"}
	body := map[string]any{"oneOf": []any{
		objectSchema([]string{"kind"}, map[string]any{"kind": point, "note": map[string]any{"type": []any{"string", "null"}}}),
		objectSchema(nil, map[string]any{"kind": point, "note": str()}),
	}}
	ok := map[string]any{"200": map[string]any{"description": "OK", "content": jsonContent(response)}}
	return map[string]any{
		"servers": []any{map[string]any{"url": "https://xquik.com"}},
		"paths": map[string]any{
			"/api/v1/x/edge": map[string]any{
				"servers": []any{},
				"get":     map[string]any{"servers": []any{}, "responses": ok},
				"post":    map[string]any{"requestBody": map[string]any{"content": jsonContent(body)}, "responses": ok},
			},
		},
	}
}

func edgeTree(accepted string) fstest.MapFS {
	tree := sdkTree(accepted, map[string]string{"service.go": edgeService, "model.go": edgeModel})
	return tree
}

func TestReaderEdgeShapes(t *testing.T) {
	expectPass(t, edgeTree("[]"), edgeSpec(), "The SDK's fields match the API for 2 operations, and accepted-field-differences.json lists 0 accepted differences.")

	spec := edgeSpec()
	path(spec, "paths", "/api/v1/x/edge", "get", "responses", "200", "content", "application/json", "schema", "properties")["hidden"] = str()
	expectFailures(t, edgeTree("[]"), spec,
		"GET /api/v1/x/edge response missing hidden",
		"POST /api/v1/x/edge response missing hidden",
		"Fix each field above in the SDK, or list it in accepted-field-differences.json with a reason.")

	missing := strings.Replace(edgeModel, "Blank   string           `", "Blank   shared.Missing   `", 1)
	expectFailures(t, sdkTree("[]", map[string]string{"service.go": edgeService, "model.go": missing}), edgeSpec(),
		"The field check cannot resolve the Go type shared.Missing. Use a type from the SDK, shared, shared/constant, time or io packages.")
}

func TestUnreadableSpecShapesFail(t *testing.T) {
	for name, change := range map[string]func(map[string]any){
		"servers not a list": func(spec map[string]any) { spec["servers"] = "https://xquik.com" },
		"method not an object": func(spec map[string]any) {
			path(spec, "paths", "/api/v1/x/edge")["get"] = "GET"
		},
	} {
		t.Run(name, func(t *testing.T) {
			spec := edgeSpec()
			change(spec)
			file := writeSpec(t, spec)
			code, out, failures := runWith(t, edgeTree("[]"), productionBaseURL, file)
			want := "Could not read the API spec from " + file + ". Check network access, or set XQUIK_OPENAPI_SPEC to a local copy of https://xquik.com/openapi.json.\n"
			if code != 1 || out != "" || failures != want {
				t.Fatalf("code = %d, out = %q, failures = %q", code, out, failures)
			}
		})
	}
	expectFailures(t, edgeTree(`[{"reason": 1, "differences": ["GET /x request extra y"]}]`), edgeSpec(),
		"accepted-field-differences.json is not valid. It must be a JSON list of groups, each with only a reason string and a list of strings named differences.")
}
