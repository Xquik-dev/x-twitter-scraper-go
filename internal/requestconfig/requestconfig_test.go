// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

package requestconfig

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/Xquik-dev/x-twitter-scraper-go/packages/param"
)

type jsonMarshalerBody struct {
	param.APIObject
	data []byte
	err  error
}

func (body jsonMarshalerBody) MarshalJSON() ([]byte, error) {
	return body.data, body.err
}

type multipartBody struct {
	param.APIObject
	data        []byte
	contentType string
	err         error
}

func (body multipartBody) MarshalMultipart() ([]byte, string, error) {
	return body.data, body.contentType, body.err
}

type jsonMultipartBody struct {
	param.APIObject
	jsonMarshalerBody
	multipartBody
}

type optionalFilesBody struct {
	jsonMultipartBody
	files bool
}

func (body optionalFilesBody) HasFiles() bool {
	return body.files
}

type queryBody struct {
	param.APIObject
	values url.Values
	err    error
}

func (body queryBody) URLQuery() (url.Values, error) {
	return body.values, body.err
}

type testHTTPDoer func(*http.Request) (*http.Response, error)

func (do testHTTPDoer) Do(req *http.Request) (*http.Response, error) {
	return do(req)
}

type errorReadCloser struct {
	err error
}

func (reader errorReadCloser) Read([]byte) (int, error) {
	return 0, reader.err
}

func (errorReadCloser) Close() error {
	return nil
}

type fallbackBody struct {
	param.APIObject
	Query string `json:"query"`
}

type unsupportedBody struct {
	param.APIObject
	Unsupported chan int `json:"unsupported"`
}

func testConfig[D any](t *testing.T, method string, path string, body param.ParamStruct, dst D) *RequestConfig {
	t.Helper()
	cfg, err := NewRequestConfig(context.Background(), method, path, body, dst)
	if err != nil {
		t.Fatal(err)
	}
	return cfg
}

func executeWithServer[D any](
	t *testing.T,
	handler http.HandlerFunc,
	method string,
	body param.ParamStruct,
	dst D,
) (*RequestConfig, error) {
	t.Helper()
	server := httptest.NewServer(handler)
	t.Cleanup(server.Close)
	cfg := testConfig(t, method, "/resource", body, dst)
	baseURL, err := url.Parse(server.URL + "/")
	if err != nil {
		t.Fatal(err)
	}
	cfg.BaseURL = baseURL
	cfg.MaxRetries = 0
	return cfg, cfg.Execute()
}

func TestNewRequestConfigSerializesSupportedBodies(t *testing.T) {
	wantErr := errors.New("serialization failed")
	tests := map[string]struct {
		body        param.ParamStruct
		path        string
		contentType string
		wantBody    string
		wantQuery   string
		wantErr     error
	}{
		"JSON marshaler": {
			body:        jsonMarshalerBody{data: []byte(`{"tweet":"hello"}`)},
			path:        "/tweets",
			contentType: "application/json",
			wantBody:    `{"tweet":"hello"}`,
		},
		"JSON marshaler error": {
			body:    jsonMarshalerBody{err: wantErr},
			path:    "/tweets",
			wantErr: wantErr,
		},
		"multipart marshaler": {
			body:        multipartBody{data: []byte("multipart"), contentType: "multipart/form-data; boundary=test"},
			path:        "/media",
			contentType: "multipart/form-data; boundary=test",
			wantBody:    "multipart",
		},
		"multipart marshaler error": {
			body:    multipartBody{err: wantErr},
			path:    "/media",
			wantErr: wantErr,
		},
		"multipart marshaler takes precedence over JSON": {
			body: jsonMultipartBody{
				jsonMarshalerBody: jsonMarshalerBody{err: wantErr},
				multipartBody:     multipartBody{data: []byte("multipart"), contentType: "multipart/form-data; boundary=test"},
			},
			path:        "/media",
			contentType: "multipart/form-data; boundary=test",
			wantBody:    "multipart",
		},
		"multipart marshaler with files": {
			body: optionalFilesBody{
				jsonMultipartBody: jsonMultipartBody{
					jsonMarshalerBody: jsonMarshalerBody{data: []byte(`{"body":"Please help."}`)},
					multipartBody:     multipartBody{data: []byte("multipart"), contentType: "multipart/form-data; boundary=test"},
				},
				files: true,
			},
			path:        "/support/tickets",
			contentType: "multipart/form-data; boundary=test",
			wantBody:    "multipart",
		},
		"multipart marshaler without files": {
			body: optionalFilesBody{
				jsonMultipartBody: jsonMultipartBody{
					jsonMarshalerBody: jsonMarshalerBody{data: []byte(`{"body":"Please help."}`)},
					multipartBody:     multipartBody{err: errors.New("multipart was marshaled")},
				},
			},
			path:        "/support/tickets",
			contentType: "application/json",
			wantBody:    `{"body":"Please help."}`,
		},
		"query": {
			body:      queryBody{values: url.Values{"q": {"tweet search"}}},
			path:      "/tweets?existing=1",
			wantQuery: "existing=1&q=tweet+search",
		},
		"empty query": {
			body: queryBody{values: url.Values{}},
			path: "/tweets",
		},
		"query error": {
			body:    queryBody{err: wantErr},
			path:    "/tweets",
			wantErr: wantErr,
		},
		"fallback JSON": {
			body:        fallbackBody{Query: "timeline"},
			path:        "/tweets",
			contentType: "application/json",
			wantBody:    "{\"query\":\"timeline\"}\n",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			cfg, err := NewRequestConfig(
				context.Background(),
				http.MethodPost,
				test.path,
				test.body,
				ResponseBody{},
			)
			if test.wantErr != nil {
				if !errors.Is(err, test.wantErr) {
					t.Fatalf("NewRequestConfig() error = %v", err)
				}
				return
			}
			if err != nil {
				t.Fatal(err)
			}
			if got := cfg.Request.Header.Get("Content-Type"); got != test.contentType {
				t.Fatalf("Content-Type = %q", got)
			}
			if got := cfg.Request.URL.RawQuery; got != test.wantQuery {
				t.Fatalf("RawQuery = %q", got)
			}
			if test.wantBody != "" {
				content, readErr := io.ReadAll(cfg.Body)
				if readErr != nil {
					t.Fatal(readErr)
				}
				if got := string(content); got != test.wantBody {
					t.Fatalf("Body = %q", got)
				}
			}
		})
	}
}

func TestNewRequestConfigRejectsInvalidInput(t *testing.T) {
	wantErr := errors.New("option failed")
	for _, test := range []struct {
		name    string
		method  string
		path    string
		body    param.ParamStruct
		option  RequestOption
		wantErr error
	}{
		{"unsupported JSON body", http.MethodPost, "/tweets", unsupportedBody{Unsupported: make(chan int)}, nil, nil},
		{"invalid method", "bad\nmethod", "/tweets", nil, nil, nil},
		{"invalid query URL", http.MethodGet, "/%zz", queryBody{values: url.Values{"limit": {"1"}}}, nil, url.EscapeError("%zz")},
		{"option failure", http.MethodGet, "/tweets", nil, RequestOptionFunc(func(*RequestConfig) error { return wantErr }), wantErr},
	} {
		t.Run(test.name, func(t *testing.T) {
			var options []RequestOption
			if test.option != nil {
				options = append(options, test.option)
			}
			_, err := NewRequestConfig(context.Background(), test.method, test.path, test.body, ResponseBody{}, options...)
			if err == nil || test.wantErr != nil && !errors.Is(err, test.wantErr) {
				t.Fatalf("error = %v, expected rejection with cause %v", err, test.wantErr)
			}
		})
	}
}

func TestNewRequestConfigAppliesTimeoutAndSecurity(t *testing.T) {
	for _, override := range []bool{false, true} {
		name := "configured headers"
		want := map[string]string{
			"X-Stainless-Timeout": "3",
			"X-API-Key":           "api-key",
			"Authorization":       "Bearer bearer-token",
		}
		if override {
			name = "caller headers"
			want = map[string]string{"X-Stainless-Timeout": "custom", "X-API-Key": "caller", "Authorization": "caller"}
		}
		t.Run(name, func(t *testing.T) {
			cfg, err := NewRequestConfig(context.Background(), http.MethodGet, "/tweets", nil, ResponseBody{},
				RequestOptionFunc(func(cfg *RequestConfig) error {
					cfg.RequestTimeout = 3 * time.Second
					cfg.APIKey = "api-key"
					cfg.BearerToken = "bearer-token"
					if override {
						for key, value := range want {
							cfg.Request.Header.Set(key, value)
						}
					}
					return nil
				}))
			if err != nil {
				t.Fatal(err)
			}
			for key, value := range want {
				if got := cfg.Request.Header.Get(key); got != value {
					t.Errorf("%s = %q, want %q", key, got, value)
				}
			}
		})
	}
}

func expectDecoded[D any](t *testing.T, contentType string, body string, want D, wantErr bool, errorText string) {
	t.Helper()
	dst := new(D)
	_, err := executeWithServer(t, func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", contentType)
		if _, err := io.WriteString(w, body); err != nil {
			t.Error(err)
		}
	}, http.MethodGet, nil, dst)
	if (err != nil) != wantErr {
		t.Fatalf("error = %v, want error = %t", err, wantErr)
	}
	if err != nil {
		if !strings.Contains(err.Error(), errorText) {
			t.Fatalf("error = %v, want containing %q", err, errorText)
		}
		return
	}
	if !reflect.DeepEqual(*dst, want) {
		t.Fatalf("destination = %#v, want %#v", *dst, want)
	}
}

func TestExecuteDecodesResponses(t *testing.T) {
	type tweet struct {
		ID string `json:"id"`
	}
	followers := "followers"
	t.Run("JSON object", func(t *testing.T) {
		expectDecoded(t, "application/json", `{"id":"tweet-1"}`, tweet{ID: "tweet-1"}, false, "")
	})
	t.Run("JSON bytes", func(t *testing.T) {
		expectDecoded(t, "application/problem+json", `{"ok":true}`, []byte(`{"ok":true}`), false, "")
	})
	t.Run("plain string", func(t *testing.T) {
		expectDecoded(t, "text/plain; charset=utf-8", "timeline", "timeline", false, "")
	})
	t.Run("plain string pointer", func(t *testing.T) {
		expectDecoded(t, "text/plain", "followers", &followers, false, "")
	})
	t.Run("plain bytes", func(t *testing.T) {
		expectDecoded(t, "application/octet-stream", "media", []byte("media"), false, "")
	})
	t.Run("unsupported plain destination", func(t *testing.T) {
		expectDecoded(t, "text/plain", "42", 0, true, "")
	})
	t.Run("invalid JSON", func(t *testing.T) {
		expectDecoded(t, "application/json", "{", json.RawMessage(nil), true, "error parsing response json")
	})
}

func TestIntoLeavesNilInterfacesUnread(t *testing.T) {
	if body := Into(io.Reader(nil)); body.raw != nil || body.decode != nil {
		t.Fatal("a nil destination reads the response body")
	}
}

func TestExecuteReturnsResponsesAndAPIErrors(t *testing.T) {
	t.Run("response destinations", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			w.WriteHeader(http.StatusNoContent)
		}))
		defer server.Close()
		var response *http.Response
		var responseBody *http.Response
		cfg := testConfig(t, http.MethodGet, "/resource", nil, &responseBody)
		baseURL, err := url.Parse(server.URL + "/")
		if err != nil {
			t.Fatal(err)
		}
		cfg.BaseURL = baseURL
		cfg.MaxRetries = 0
		cfg.ResponseInto = &response
		err = cfg.Execute()
		if err != nil {
			t.Fatal(err)
		}
		if response == nil || responseBody == nil {
			t.Fatal("raw response destination was not populated")
		}
	})

	t.Run("API error", func(t *testing.T) {
		cfg, err := executeWithServer(t, func(w http.ResponseWriter, _ *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			if _, err := io.WriteString(w, `{"message":"bad request"}`); err != nil {
				t.Error(err)
			}
		}, http.MethodGet, nil, ResponseBody{})
		if err == nil || !strings.Contains(err.Error(), "400 Bad Request") {
			t.Fatalf("error = %v", err)
		}
		if cfg.Request == nil {
			t.Fatal("request was not retained")
		}
	})

	t.Run("invalid API error", func(t *testing.T) {
		_, err := executeWithServer(t, func(w http.ResponseWriter, _ *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			if _, err := io.WriteString(w, "{"); err != nil {
				t.Error(err)
			}
		}, http.MethodGet, nil, ResponseBody{})
		if err == nil {
			t.Fatal("invalid API error JSON was accepted")
		}
	})

	t.Run("read error", func(t *testing.T) {
		var dst json.RawMessage
		cfg := testConfig(t, http.MethodGet, "/resource", nil, &dst)
		baseURL, err := url.Parse("https://example.com/")
		if err != nil {
			t.Fatal(err)
		}
		cfg.BaseURL = baseURL
		cfg.MaxRetries = 0
		cfg.CustomHTTPDoer = testHTTPDoer(func(*http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: http.StatusOK,
				Header:     http.Header{"Content-Type": {"application/json"}},
				Body:       errorReadCloser{err: errors.New("read failed")},
			}, nil
		})
		err = cfg.Execute()
		if err == nil || !strings.Contains(err.Error(), "error reading response body") {
			t.Fatalf("error = %v", err)
		}
	})
}

func TestExecuteRetriesAndUsesRequestBodies(t *testing.T) {
	var attempts int
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		attempts++
		content, err := io.ReadAll(req.Body)
		if err != nil {
			t.Fatal(err)
		}
		if string(content) != "body" {
			t.Errorf("request body = %q", content)
			return
		}
		if attempts == 1 {
			w.Header().Set("Retry-After-Ms", "0")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		if _, err := io.WriteString(w, "{}"); err != nil {
			t.Error(err)
		}
	}))
	defer server.Close()

	var dst json.RawMessage
	reader := bytes.NewReader([]byte("skipbody"))
	if _, err := reader.Seek(4, io.SeekStart); err != nil {
		t.Fatal(err)
	}
	cfg := testConfig(t, http.MethodPost, "/resource", nil, &dst)
	cfg.Body = reader
	baseURL, err := url.Parse(server.URL + "/")
	if err != nil {
		t.Fatal(err)
	}
	cfg.BaseURL = baseURL
	cfg.MaxRetries = 1
	if err := cfg.Execute(); err != nil {
		t.Fatal(err)
	}
	if attempts != 2 {
		t.Fatalf("attempts = %d", attempts)
	}

	for name, body := range map[string]io.Reader{
		"byte reader":   bytes.NewReader([]byte("body")),
		"byte buffer":   bytes.NewBufferString("body"),
		"string reader": strings.NewReader("body"),
		"read closer":   io.NopCloser(strings.NewReader("body")),
	} {
		t.Run(name, func(t *testing.T) {
			readerCfg := testConfig(t, http.MethodPost, "/resource", nil, ResponseBody{})
			readerCfg.Body = body
			readerCfg.BaseURL = cfg.BaseURL
			readerCfg.MaxRetries = 0
			if err := readerCfg.Execute(); err != nil {
				t.Fatal(err)
			}
		})
	}
}

func TestExecuteErrorsWithoutBaseURLOrAfterTransportFailure(t *testing.T) {
	cfg := testConfig(t, http.MethodGet, "/resource", nil, ResponseBody{})
	if err := cfg.Execute(); err == nil || !strings.Contains(err.Error(), "base url is not set") {
		t.Fatalf("error = %v", err)
	}

	baseURL, err := url.Parse("https://example.com/")
	if err != nil {
		t.Fatal(err)
	}
	cfg.DefaultBaseURL = baseURL
	wantErr := errors.New("transport failed")
	cfg.CustomHTTPDoer = testHTTPDoer(func(*http.Request) (*http.Response, error) {
		return nil, wantErr
	})
	cfg.MaxRetries = 0
	if err := cfg.Execute(); !errors.Is(err, wantErr) {
		t.Fatalf("error = %v", err)
	}
}
