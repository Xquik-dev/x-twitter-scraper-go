// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
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

type PublicMetadata = shared.PublicMetadata

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
	Extra  shared.PublicMetadata        ` + "`" + `json:"extra"` + "`" + `
	Raw    json.RawMessage              ` + "`" + `json:"raw"` + "`" + `
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

type PublicMetadata map[string]json.RawMessage

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

type node interface{ json.Marshaler }

type (
	obj     map[string]node
	arr     []node
	text    string
	num     float64
	boolean bool
	null    struct{}
)

func (o obj) MarshalJSON() ([]byte, error)     { return json.Marshal(map[string]node(o)) }
func (a arr) MarshalJSON() ([]byte, error)     { return json.Marshal([]node(a)) }
func (s text) MarshalJSON() ([]byte, error)    { return json.Marshal(string(s)) }
func (n num) MarshalJSON() ([]byte, error)     { return json.Marshal(float64(n)) }
func (b boolean) MarshalJSON() ([]byte, error) { return json.Marshal(bool(b)) }
func (null) MarshalJSON() ([]byte, error)      { return []byte("null"), nil }

func asObj(value node) obj {
	object, ok := value.(obj)
	if !ok {
		panic(fmt.Sprintf("the test spec holds %T where it needs an object", value))
	}
	return object
}

func asArr(value node) arr {
	list, ok := value.(arr)
	if !ok {
		panic(fmt.Sprintf("the test spec holds %T where it needs a list", value))
	}
	return list
}

func objectSchema(required []string, properties obj) obj {
	schema := obj{"type": text("object"), "properties": properties}
	if required != nil {
		names := arr{}
		for _, name := range required {
			names = append(names, text(name))
		}
		schema["required"] = names
	}
	return schema
}

func jsonContent(schema node) obj {
	return obj{"application/json": obj{"schema": schema}}
}

func str() obj { return obj{"type": text("string")} }

func baseSpec() obj {
	tweet := objectSchema([]string{"id"}, obj{
		"id":     str(),
		"text":   str(),
		"note":   obj{"type": arr{text("string"), text("null")}},
		"type":   obj{"type": text("string"), "const": text("Point")},
		"quoted": obj{"$ref": text("#/components/schemas/Tweet")},
		"tags":   obj{"type": text("array"), "items": str()},
	})
	profile := objectSchema([]string{"id"}, obj{"id": str(), "name": str()})
	response := objectSchema([]string{"tweet"}, obj{
		"tweet":  obj{"$ref": text("#/components/schemas/Tweet"), "description": text("The tweet.")},
		"author": obj{"$ref": text("#/components/schemas/UserProfile")},
		"media":  obj{"type": text("array"), "items": objectSchema([]string{"url"}, obj{"url": str()})},
		"extra":  obj{"type": text("object"), "additionalProperties": boolean(true)},
		"raw":    objectSchema(nil, obj{"anything": str()}),
		"answer": obj{"oneOf": arr{obj{"type": text("number")}, str(), obj{"type": text("null")}}},
		"result": obj{"oneOf": arr{
			objectSchema([]string{"kind", "score"}, obj{"kind": str(), "score": obj{"type": text("number")}}),
			objectSchema([]string{"kind"}, obj{"kind": str()}),
		}},
	})
	return obj{
		"openapi": text("3.1.0"),
		"servers": arr{obj{"url": text("https://xquik.com")}},
		"components": obj{
			"schemas":    obj{"Tweet": tweet, "UserProfile": profile},
			"parameters": obj{"Cursor": obj{"name": text("cursor"), "in": text("query"), "schema": str()}},
		},
		"paths": obj{
			"/api/v1/x/tweets/{tweetId}": obj{
				"parameters": arr{obj{"name": text("tweetId"), "in": text("path"), "required": boolean(true), "schema": str()}},
				"get": obj{
					"operationId": text("getTweet"),
					"parameters": arr{
						obj{"name": text("q"), "in": text("query"), "required": boolean(true), "schema": str()},
						obj{"$ref": text("#/components/parameters/Cursor")},
						obj{"name": text("mode"), "in": text("query"), "schema": obj{"type": text("string"), "enum": arr{text("latest"), text("top")}}},
						obj{"name": text("Accept-Language"), "in": text("header"), "schema": str()},
					},
					"responses": obj{
						"200":     obj{"description": text("OK"), "content": jsonContent(response)},
						"404":     obj{"description": text("Missing"), "content": jsonContent(objectSchema(nil, obj{"error": str()}))},
						"default": obj{"description": text("Error")},
					},
				},
			},
			"/api/v1/x/tweets": obj{
				"post": obj{
					"requestBody": obj{"content": obj{
						"application/json":    obj{"schema": objectSchema([]string{"text"}, obj{"text": str(), "media": obj{"type": text("array"), "items": str()}})},
						"multipart/form-data": obj{"schema": objectSchema([]string{"text"}, obj{"text": str(), "media": obj{"type": text("array"), "items": str()}})},
					}},
					"responses": obj{"201": obj{"description": text("Created"), "content": jsonContent(objectSchema([]string{"id"}, obj{"id": str(), "status": obj{"type": text("string"), "enum": arr{text("pending")}}}))}},
				},
			},
			"/api/v1/x/tweets/{id}/export": obj{
				"get": obj{"responses": obj{"200": obj{"description": text("CSV"), "content": obj{"text/csv": obj{"schema": str()}}}}},
			},
			"/api/v1/x/media": obj{
				"post": obj{
					"requestBody": obj{"content": obj{"multipart/form-data": obj{"schema": objectSchema([]string{"file"}, obj{"file": obj{"type": text("string"), "format": text("binary")}})}}},
					"responses":   obj{"200": obj{"description": text("OK"), "content": jsonContent(obj{"$ref": text("#/components/schemas/UserProfile")})}},
				},
			},
			"/api/v1/events/stream": obj{
				"get": obj{"responses": obj{"200": obj{"description": text("Stream"), "content": obj{"text/event-stream": obj{"schema": str()}}}}},
			},
		},
	}
}

func path(spec obj, keys ...string) obj {
	current := spec
	for _, key := range keys {
		current = asObj(current[key])
	}
	return current
}

func lookupResponse(spec obj) obj {
	return path(spec, "paths", "/api/v1/x/tweets/{tweetId}", "get", "responses", "200", "content", "application/json", "schema", "properties")
}

func lookupParameters(spec obj) arr {
	return asArr(path(spec, "paths", "/api/v1/x/tweets/{tweetId}", "get")["parameters"])
}

func tweetSchema(spec obj) obj {
	return path(spec, "components", "schemas", "Tweet")
}

func writeSpec(t *testing.T, spec node) string {
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

func expectFailures(t *testing.T, root fs.FS, spec node, want ...string) {
	t.Helper()
	code, out, failures := runWith(t, root, productionBaseURL, writeSpec(t, spec))
	if expected := strings.Join(want, "\n") + "\n"; code != 1 || out != "" || failures != expected {
		t.Fatalf("code = %d, out = %q\nfailures:\n%s\nwant:\n%s", code, out, failures, expected)
	}
}

func expectPass(t *testing.T, root fs.FS, spec node, want string) {
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
			spec := obj{
				"openapi": text("3.1.0"),
				"paths": obj{
					"/api/v1/support/tickets": obj{
						"post": obj{
							"requestBody": obj{"content": obj{
								"application/json": obj{"schema": objectSchema([]string{"body"}, obj{"body": str()})},
								mediaType: obj{"schema": objectSchema(nil, obj{
									"body":        str(),
									"attachments": obj{"type": text("array"), "items": obj{"type": text("string"), "format": text("binary")}},
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
		change func(obj)
		want   []string
	}{
		{"missing query param", func(spec obj) {
			path(spec, "paths", "/api/v1/x/tweets/{tweetId}", "get")["parameters"] = append(lookupParameters(spec), obj{"name": text("limit"), "in": text("query"), "schema": obj{"type": text("integer")}})
		}, []string{"GET /api/v1/x/tweets/{tweetId} request missing limit"}},
		{"extra query param", func(spec obj) {
			path(spec, "paths", "/api/v1/x/tweets/{tweetId}", "get")["parameters"] = lookupParameters(spec)[:1]
		}, []string{"GET /api/v1/x/tweets/{tweetId} request extra cursor", "GET /api/v1/x/tweets/{tweetId} request extra mode"}},
		{"query enum", func(spec obj) {
			asObj(lookupParameters(spec)[2])["schema"] = obj{"type": text("string"), "enum": arr{text("latest"), text("top"), text("live")}}
		}, []string{"GET /api/v1/x/tweets/{tweetId} request enum mode"}},
		{"query required", func(spec obj) {
			asObj(lookupParameters(spec)[0])["required"] = boolean(false)
		}, []string{"GET /api/v1/x/tweets/{tweetId} request required q"}},
		{"nullable", func(spec obj) {
			path(tweetSchema(spec), "properties")["note"] = str()
		}, []string{"GET /api/v1/x/tweets/{tweetId} response nullable tweet.note"}},
		{"response required", func(spec obj) {
			tweetSchema(spec)["required"] = arr{}
		}, []string{"GET /api/v1/x/tweets/{tweetId} response required tweet.id"}},
		{"missing nested object counts once", func(spec obj) {
			path(tweetSchema(spec), "properties")["entities"] = objectSchema(nil, obj{"urls": obj{"type": text("array"), "items": objectSchema(nil, obj{"url": str()})}})
		}, []string{"GET /api/v1/x/tweets/{tweetId} response missing tweet.entities"}},
		{"missing array item field", func(spec obj) {
			path(lookupResponse(spec), "media", "items", "properties")["alt"] = str()
		}, []string{"GET /api/v1/x/tweets/{tweetId} response missing media[].alt"}},
		{"extra response field", func(spec obj) {
			delete(path(tweetSchema(spec), "properties"), "tags")
		}, []string{"GET /api/v1/x/tweets/{tweetId} response extra tweet.tags"}},
		{"another kind counts as missing", func(spec obj) {
			path(tweetSchema(spec), "properties")["text"] = objectSchema(nil, obj{"value": str()})
			path(tweetSchema(spec), "properties")["tags"] = str()
		}, []string{"GET /api/v1/x/tweets/{tweetId} response missing tweet.tags", "GET /api/v1/x/tweets/{tweetId} response missing tweet.text"}},
		{"response enum", func(spec obj) {
			path(spec, "paths", "/api/v1/x/tweets", "post", "responses", "201", "content", "application/json", "schema", "properties")["status"] = obj{"type": text("string"), "enum": arr{text("pending"), text("sent")}}
		}, []string{"POST /api/v1/x/tweets response enum status"}},
		{"several 2xx bodies merge", func(spec obj) {
			path(spec, "paths", "/api/v1/x/tweets", "post", "responses")["202"] = obj{"description": text("Queued"), "content": jsonContent(objectSchema([]string{"jobId"}, obj{"jobId": str()}))}
		}, []string{"POST /api/v1/x/tweets response missing jobId", "POST /api/v1/x/tweets response required id"}},
		{"untyped JSON response", func(spec obj) {
			path(spec, "paths", "/api/v1/x/tweets/{id}/export", "get", "responses", "200")["content"] = jsonContent(objectSchema(nil, obj{"rows": str(), "total": str()}))
		}, []string{"GET /api/v1/x/tweets/{id}/export response missing rows", "GET /api/v1/x/tweets/{id}/export response missing total"}},
		{"untyped scalar response", func(spec obj) {
			path(spec, "paths", "/api/v1/x/tweets/{id}/export", "get", "responses", "200")["content"] = jsonContent(str())
		}, []string{"GET /api/v1/x/tweets/{id}/export response missing (body)"}},
		{"response root of another kind", func(spec obj) {
			path(spec, "paths", "/api/v1/x/media", "post", "responses", "200")["content"] = jsonContent(obj{"type": text("array"), "items": str()})
		}, []string{"POST /api/v1/x/media response missing (body)"}},
		{"missing body field", func(spec obj) {
			path(spec, "paths", "/api/v1/x/tweets", "post", "requestBody", "content", "application/json", "schema", "properties")["replyTo"] = str()
		}, []string{"POST /api/v1/x/tweets request missing replyTo"}},
		{"body the SDK does not send", func(spec obj) {
			path(spec, "paths", "/api/v1/x/tweets/{id}/export", "get")["requestBody"] = obj{"content": jsonContent(objectSchema(nil, obj{"format": str()}))}
		}, []string{"GET /api/v1/x/tweets/{id}/export request missing format"}},
		{"body the spec does not take", func(spec obj) {
			delete(path(spec, "paths", "/api/v1/x/tweets", "post"), "requestBody")
		}, []string{"POST /api/v1/x/tweets request extra media", "POST /api/v1/x/tweets request extra text"}},
		{"array body the SDK does not send", func(spec obj) {
			path(spec, "paths", "/api/v1/x/tweets/{id}/export", "get")["requestBody"] = obj{"content": jsonContent(obj{"type": text("array"), "items": str()})}
		}, []string{"GET /api/v1/x/tweets/{id}/export request missing (body)"}},
		{"multipart body when there is no JSON body", func(spec obj) {
			path(spec, "paths", "/api/v1/x/media", "post", "requestBody", "content", "multipart/form-data", "schema", "properties")["alt"] = str()
		}, []string{"POST /api/v1/x/media request missing alt"}},
		{"closed object rejects extra SDK fields", func(spec obj) {
			delete(path(tweetSchema(spec), "properties"), "tags")
			tweetSchema(spec)["additionalProperties"] = boolean(false)
		}, []string{"GET /api/v1/x/tweets/{tweetId} response extra tweet.tags"}},
		{"typed response the spec does not document", func(spec obj) {
			path(spec, "paths", "/api/v1/x/media", "post", "responses", "200")["content"] = obj{"text/csv": obj{"schema": str()}}
		}, []string{"POST /api/v1/x/media response extra id", "POST /api/v1/x/media response extra name"}},
		{"vendor JSON response", func(spec obj) {
			responses := path(spec, "paths", "/api/v1/x/media", "post", "responses", "200")
			schema := objectSchema([]string{"id"}, obj{"id": str(), "name": str(), "bio": str()})
			responses["content"] = obj{"application/vnd.xquik+json; charset=utf-8": obj{"schema": schema}}
		}, []string{"POST /api/v1/x/media response missing bio"}},
		{"typed object against a free-form SDK field", func(spec obj) {
			lookupResponse(spec)["extra"] = objectSchema(nil, obj{"value": str()})
		}, []string{"GET /api/v1/x/tweets/{tweetId} response missing extra"}},
		{"free-form object against an SDK array", func(spec obj) {
			lookupResponse(spec)["media"] = obj{"type": text("object"), "additionalProperties": boolean(true)}
		}, []string{"GET /api/v1/x/tweets/{tweetId} response missing media"}},
		{"free-form union against an SDK scalar union", func(spec obj) {
			lookupResponse(spec)["answer"] = obj{"anyOf": arr{obj{"type": text("object")}, str()}}
		}, []string{"GET /api/v1/x/tweets/{tweetId} response missing answer"}},
		{"union field required only when every variant requires it", func(spec obj) {
			path(lookupResponse(spec), "result")["oneOf"] = arr{objectSchema([]string{"kind"}, obj{"kind": str(), "label": str()})}
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

func TestSchemaRulesKeepMatchingFieldsQuiet(t *testing.T) {
	const summary = "The SDK's fields match the API for 5 operations, and accepted-field-differences.json lists 0 accepted differences."
	tests := []struct {
		name   string
		change func(obj)
	}{
		{"open object allows extra SDK fields", func(spec obj) {
			delete(path(tweetSchema(spec), "properties"), "tags")
			tweetSchema(spec)["additionalProperties"] = obj{"type": text("string")}
		}},
		{"allOf merges parts and a required list", func(spec obj) {
			schemas := path(spec, "components", "schemas")
			profile := asObj(schemas["UserProfile"])
			delete(profile, "required")
			schemas["UserProfile"] = obj{"allOf": arr{profile, obj{"required": arr{text("id")}}}}
		}},
		{"required list beside allOf", func(spec obj) {
			schemas := path(spec, "components", "schemas")
			profile := asObj(schemas["UserProfile"])
			delete(profile, "required")
			schemas["UserProfile"] = obj{"allOf": arr{profile}, "required": arr{text("id")}}
		}},
		{"allOf of arrays intersects items", func(spec obj) {
			media := asObj(lookupResponse(spec)["media"])
			lookupResponse(spec)["media"] = obj{"allOf": arr{media, obj{"type": text("array"), "items": obj{"required": arr{text("url")}}}}}
			path(media, "items")["required"] = arr{}
		}},
		{"ref with siblings", func(spec obj) {
			lookupResponse(spec)["author"] = obj{"$ref": text("#/components/schemas/UserProfile"), "type": arr{text("object"), text("null")}, "description": text("Author.")}
		}},
		{"ref with schema siblings", func(spec obj) {
			lookupResponse(spec)["author"] = obj{"$ref": text("#/components/schemas/UserProfile"), "properties": obj{"name": str()}}
		}},
		{"oneOf with shared properties", func(spec obj) {
			lookupResponse(spec)["result"] = obj{
				"properties": obj{"kind": str()},
				"required":   arr{text("kind")},
				"anyOf":      arr{objectSchema(nil, obj{"score": obj{"type": text("number")}}), obj{"enum": arr{null{}}}, objectSchema(nil, obj{})},
			}
		}},
		{"oneOf of 1 object", func(spec obj) {
			lookupResponse(spec)["author"] = obj{"oneOf": arr{obj{"$ref": text("#/components/schemas/UserProfile")}, obj{"type": text("null")}}}
		}},
		{"oneOf of only null", func(spec obj) {
			lookupResponse(spec)["raw"] = obj{"oneOf": arr{obj{"type": text("null")}}}
		}},
		{"null type and nullable flag", func(spec obj) {
			lookupResponse(spec)["raw"] = obj{"type": text("null")}
			path(tweetSchema(spec), "properties")["note"] = obj{"type": text("string"), "nullable": boolean(true)}
		}},
		{"enum with null", func(spec obj) {
			path(tweetSchema(spec), "properties")["note"] = obj{"type": text("string"), "enum": arr{text("a"), null{}}}
		}},
		{"free-form on both sides and tuple shapes stop the walk", func(spec obj) {
			lookupResponse(spec)["extra"] = obj{"anyOf": arr{obj{"type": text("object")}, str()}}
			path(tweetSchema(spec), "properties")["tags"] = obj{"type": text("array"), "prefixItems": arr{str()}}
		}},
		{"non-object schemas stay unknown", func(spec obj) {
			lookupResponse(spec)["raw"] = boolean(true)
			path(tweetSchema(spec), "properties")["tags"] = obj{"oneOf": arr{obj{"type": text("array"), "items": str()}, obj{"type": text("array"), "items": obj{"type": text("integer")}}}}
		}},
		{"allOf requires a field any part requires", func(spec obj) {
			schemas := path(spec, "components", "schemas")
			schemas["UserProfile"] = obj{"allOf": arr{
				objectSchema([]string{"id"}, obj{"id": str(), "name": str()}),
				objectSchema(nil, obj{"id": str()}),
			}}
		}},
		{"allOf of 2 unions distributes over both", func(spec obj) {
			lookupResponse(spec)["result"] = obj{"allOf": arr{
				obj{"oneOf": arr{objectSchema([]string{"kind"}, obj{"kind": str()}), objectSchema(nil, obj{"kind": str()})}},
				obj{"oneOf": arr{objectSchema(nil, obj{"score": obj{"type": text("number")}}), objectSchema(nil, obj{})}},
			}}
		}},
		{"oneOf beside allOf", func(spec obj) {
			lookupResponse(spec)["result"] = obj{
				"allOf": arr{objectSchema([]string{"kind"}, obj{"kind": str()})},
				"oneOf": arr{objectSchema(nil, obj{"score": obj{"type": text("number")}}), objectSchema(nil, obj{})},
			}
		}},
		{"null type beside a ref with other siblings", func(spec obj) {
			path(spec, "components", "schemas")["Note"] = str()
			path(tweetSchema(spec), "properties")["note"] = obj{"$ref": text("#/components/schemas/Note"), "type": arr{text("string"), text("null")}, "format": text("text")}
		}},
		{"operation params override path item params", func(spec obj) {
			item := path(spec, "paths", "/api/v1/x/tweets/{tweetId}")
			item["parameters"] = append(asArr(item["parameters"]), obj{"name": text("q"), "in": text("query"), "required": boolean(false), "schema": str()})
		}},
		{"numeric and boolean enums", func(spec obj) {
			lookupResponse(spec)["answer"] = obj{"oneOf": arr{obj{"enum": arr{num(1.5), boolean(true)}}, obj{"enum": arr{obj{}}}, obj{"const": null{}}}}
		}},
		{"free-form object accepts a typed SDK struct", func(spec obj) {
			lookupResponse(spec)["author"] = obj{"type": text("object"), "additionalProperties": boolean(true)}
		}},
		{"a keyword of another JSON type reads as absent", func(spec obj) {
			path(tweetSchema(spec), "properties")["text"] = obj{"type": text("string"), "required": text("text")}
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
	lookupResponse(spec)["author"] = obj{"$ref": text("#/components/schemas/Missing")}
	expectFailures(t, sdkTree("[]", nil), spec, "The field check cannot resolve #/components/schemas/Missing in the API spec. Fix the reference in the spec.")

	broken := strings.Replace(modelSource, "Media  []TweetMedia ", "Media  []TweetMedium", 1)
	broken = strings.Replace(broken, "Extra  shared.PublicMetadata ", "Extra  chan int ", 1)
	broken += "\ntype Other struct {\n\tValue *param.Field[string] `json:\"value\"`\n\tOdd (string) `json:\"odd\"`\n\tPicked pkg.Generic[string] `json:\"picked\"`\n}\n"
	broken = strings.Replace(broken, "Result XTweetGetResponseResultUnion ", "Result Other                        ", 1)
	expectFailures(t, sdkTree("[]", map[string]string{"model.go": broken}), baseSpec(),
		"The field check cannot resolve the Go type (string). Use a type from the SDK, shared, shared/constant, time or io packages.",
		"The field check cannot resolve the Go type TweetMedium. Use a type from the SDK, shared, shared/constant, time or io packages.",
		"The field check cannot resolve the Go type chan int. Use a type from the SDK, shared, shared/constant, time or io packages.",
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

func TestFreeFormFields(t *testing.T) {
	const reject = "The field check rejects the Go type %s. Use shared.PublicMetadata for a free-form object and json.RawMessage for a free-form value."
	for _, test := range []struct {
		goType   string
		failures []string
	}{
		{"shared.PublicMetadata", nil},
		{"PublicMetadata", nil},
		{"map[string]json.RawMessage", nil},
		{"json.RawMessage", nil},
		{"map[string]string", nil},
		{"[]json.RawMessage", []string{"GET /api/v1/x/tweets/{tweetId} response missing extra", nextStep}},
		{"map[string]any", []string{fmt.Sprintf(reject, "map[string]any")}},
		{"map[string]interface{}", []string{fmt.Sprintf(reject, "map[string]interface{}")}},
		{"any", []string{fmt.Sprintf(reject, "any")}},
		{"interface{}", []string{fmt.Sprintf(reject, "interface{}")}},
		{"[]any", []string{fmt.Sprintf(reject, "any")}},
	} {
		t.Run(test.goType, func(t *testing.T) {
			model := strings.Replace(modelSource, "Extra  shared.PublicMetadata ", "Extra  "+test.goType+" ", 1)
			if !strings.Contains(model, "Extra  "+test.goType+" ") {
				t.Fatal("the model has no Extra field to retype")
			}
			root := sdkTree("[]", map[string]string{"model.go": model})
			if test.failures == nil {
				expectPass(t, root, baseSpec(), "The SDK's fields match the API for 5 operations, and accepted-field-differences.json lists 0 accepted differences.")
				return
			}
			expectFailures(t, root, baseSpec(), test.failures...)
		})
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
	Raw     interface{ implRaw() } ` + "`" + `json:"raw"` + "`" + `
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

func edgeSpec() obj {
	response := objectSchema(nil, obj{
		"raw":    objectSchema(nil, obj{"z": str()}),
		"inline": objectSchema(nil, obj{"a": str()}),
		"variant": obj{"allOf": arr{
			objectSchema([]string{"kind"}, obj{"kind": str()}),
			obj{"oneOf": arr{objectSchema(nil, obj{"x": str()}), objectSchema(nil, obj{"y": str()})}},
		}},
		"blank":   obj{"oneOf": arr{boolean(true), boolean(true)}},
		"req":     obj{"allOf": arr{obj{"required": arr{text("a")}}}},
		"shared":  str(),
		"skipped": objectSchema(nil, obj{"hidden": str()}),
	})
	point := obj{"type": text("string"), "const": text("Point")}
	body := obj{"oneOf": arr{
		objectSchema([]string{"kind"}, obj{"kind": point, "note": obj{"type": arr{text("string"), text("null")}}}),
		objectSchema(nil, obj{"kind": point, "note": str()}),
	}}
	ok := obj{"200": obj{"description": text("OK"), "content": jsonContent(response)}}
	return obj{
		"servers": arr{obj{"url": text("https://xquik.com")}},
		"paths": obj{
			"/api/v1/x/edge": obj{
				"servers": arr{},
				"get":     obj{"servers": arr{}, "responses": ok},
				"post":    obj{"requestBody": obj{"content": jsonContent(body)}, "responses": ok},
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
	for name, change := range map[string]func(obj){
		"servers not a list": func(spec obj) { spec["servers"] = text("https://xquik.com") },
		"method not an object": func(spec obj) {
			path(spec, "paths", "/api/v1/x/edge")["get"] = text("GET")
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
