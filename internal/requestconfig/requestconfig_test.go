// Copyright the Xquik contributors.
// SPDX-License-Identifier: Apache-2.0

package requestconfig

import (
	"bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"strings"
	"testing"
	"time"
)

type jsonMarshalerBody struct {
	data []byte
	err  error
}

func (body jsonMarshalerBody) MarshalJSON() ([]byte, error) {
	return body.data, body.err
}

type multipartBody struct {
	data        []byte
	contentType string
	err         error
}

func (body multipartBody) MarshalMultipart() ([]byte, string, error) {
	return body.data, body.contentType, body.err
}

type queryBody struct {
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

func testConfig(t *testing.T, method string, path string, body any, dst any) *RequestConfig {
	t.Helper()
	cfg, err := NewRequestConfig(context.Background(), method, path, body, dst)
	if err != nil {
		t.Fatal(err)
	}
	return cfg
}

func executeWithServer(
	t *testing.T,
	handler http.HandlerFunc,
	method string,
	body any,
	dst any,
) (*RequestConfig, error) {
	t.Helper()
	server := httptest.NewServer(handler)
	t.Cleanup(server.Close)
	cfg := testConfig(t, method, "/resource", body, dst)
	cfg.BaseURL, _ = url.Parse(server.URL + "/")
	cfg.MaxRetries = 0
	return cfg, cfg.Execute()
}

func TestNewRequestConfigSerializesSupportedBodies(t *testing.T) {
	wantErr := errors.New("serialization failed")
	tests := map[string]struct {
		body        any
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
		"bytes": {
			body:        []byte("bytes"),
			path:        "/tweets",
			contentType: "application/json",
			wantBody:    "bytes",
		},
		"reader": {
			body:        strings.NewReader("reader"),
			path:        "/tweets",
			contentType: "application/json",
			wantBody:    "reader",
		},
		"fallback JSON": {
			body:        map[string]string{"query": "timeline"},
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
				nil,
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
		body    any
		option  RequestOption
		wantErr error
	}{
		{"unsupported JSON body", http.MethodPost, "/tweets", map[string]any{"unsupported": make(chan int)}, nil, nil},
		{"invalid method", "bad\nmethod", "/tweets", nil, nil, nil},
		{"invalid query URL", http.MethodGet, "/%zz", queryBody{values: url.Values{"limit": {"1"}}}, nil, url.EscapeError("%zz")},
		{"option failure", http.MethodGet, "/tweets", nil, RequestOptionFunc(func(*RequestConfig) error { return wantErr }), wantErr},
	} {
		t.Run(test.name, func(t *testing.T) {
			var options []RequestOption
			if test.option != nil {
				options = append(options, test.option)
			}
			_, err := NewRequestConfig(context.Background(), test.method, test.path, test.body, nil, options...)
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
			cfg, err := NewRequestConfig(context.Background(), http.MethodGet, "/tweets", nil, nil,
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

func TestExecuteDecodesResponses(t *testing.T) {
	type tweet struct {
		ID string `json:"id"`
	}
	followers := "followers"
	for _, test := range []struct {
		name        string
		contentType string
		body        string
		dst         any
		want        any
		wantErr     bool
		errorText   string
	}{
		{"JSON object", "application/json", `{"id":"tweet-1"}`, new(tweet), tweet{ID: "tweet-1"}, false, ""},
		{"JSON bytes", "application/problem+json", `{"ok":true}`, new([]byte), []byte(`{"ok":true}`), false, ""},
		{"plain string", "text/plain; charset=utf-8", "timeline", new(string), "timeline", false, ""},
		{"plain string pointer", "text/plain", "followers", new(*string), &followers, false, ""},
		{"plain bytes", "application/octet-stream", "media", new([]byte), []byte("media"), false, ""},
		{"unsupported plain destination", "text/plain", "42", new(int), nil, true, ""},
		{"invalid JSON", "application/json", "{", new(map[string]any), nil, true, "error parsing response json"},
	} {
		t.Run(test.name, func(t *testing.T) {
			_, err := executeWithServer(t, func(w http.ResponseWriter, _ *http.Request) {
				w.Header().Set("Content-Type", test.contentType)
				_, _ = io.WriteString(w, test.body)
			}, http.MethodGet, nil, test.dst)
			if (err != nil) != test.wantErr {
				t.Fatalf("error = %v, want error = %t", err, test.wantErr)
			}
			if err != nil {
				if !strings.Contains(err.Error(), test.errorText) {
					t.Fatalf("error = %v, want containing %q", err, test.errorText)
				}
				return
			}
			if got := reflect.ValueOf(test.dst).Elem().Interface(); !reflect.DeepEqual(got, test.want) {
				t.Fatalf("destination = %#v, want %#v", got, test.want)
			}
		})
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
		cfg.BaseURL, _ = url.Parse(server.URL + "/")
		cfg.MaxRetries = 0
		cfg.ResponseInto = &response
		err := cfg.Execute()
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
			_, _ = io.WriteString(w, `{"message":"bad request"}`)
		}, http.MethodGet, nil, nil)
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
			_, _ = io.WriteString(w, "{")
		}, http.MethodGet, nil, nil)
		if err == nil {
			t.Fatal("invalid API error JSON was accepted")
		}
	})

	t.Run("read error", func(t *testing.T) {
		var dst map[string]any
		cfg := testConfig(t, http.MethodGet, "/resource", nil, &dst)
		cfg.BaseURL, _ = url.Parse("https://example.com/")
		cfg.MaxRetries = 0
		cfg.CustomHTTPDoer = testHTTPDoer(func(*http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: http.StatusOK,
				Header:     http.Header{"Content-Type": {"application/json"}},
				Body:       errorReadCloser{err: errors.New("read failed")},
			}, nil
		})
		err := cfg.Execute()
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
		_, _ = io.WriteString(w, "{}")
	}))
	defer server.Close()

	var dst map[string]any
	reader := bytes.NewReader([]byte("skipbody"))
	if _, err := reader.Seek(4, io.SeekStart); err != nil {
		t.Fatal(err)
	}
	cfg := testConfig(t, http.MethodPost, "/resource", reader, &dst)
	cfg.BaseURL, _ = url.Parse(server.URL + "/")
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
			readerCfg := testConfig(t, http.MethodPost, "/resource", body, nil)
			readerCfg.BaseURL = cfg.BaseURL
			readerCfg.MaxRetries = 0
			if err := readerCfg.Execute(); err != nil {
				t.Fatal(err)
			}
		})
	}
}

func TestExecuteErrorsWithoutBaseURLOrAfterTransportFailure(t *testing.T) {
	cfg := testConfig(t, http.MethodGet, "/resource", nil, nil)
	if err := cfg.Execute(); err == nil || !strings.Contains(err.Error(), "base url is not set") {
		t.Fatalf("error = %v", err)
	}

	cfg.DefaultBaseURL, _ = url.Parse("https://example.com/")
	wantErr := errors.New("transport failed")
	cfg.CustomHTTPDoer = testHTTPDoer(func(*http.Request) (*http.Response, error) {
		return nil, wantErr
	})
	cfg.MaxRetries = 0
	if err := cfg.Execute(); !errors.Is(err, wantErr) {
		t.Fatalf("error = %v", err)
	}
}
