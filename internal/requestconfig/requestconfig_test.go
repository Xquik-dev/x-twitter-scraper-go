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
	if _, err := NewRequestConfig(
		context.Background(),
		http.MethodPost,
		"/tweets",
		map[string]any{"unsupported": make(chan int)},
		nil,
	); err == nil {
		t.Fatal("unsupported JSON body was accepted")
	}

	if _, err := NewRequestConfig(
		context.Background(),
		"bad\nmethod",
		"/tweets",
		nil,
		nil,
	); err == nil {
		t.Fatal("invalid method was accepted")
	}

	wantErr := errors.New("option failed")
	failingOption := RequestOptionFunc(func(*RequestConfig) error { return wantErr })
	if _, err := NewRequestConfig(
		context.Background(),
		http.MethodGet,
		"/tweets",
		nil,
		nil,
		failingOption,
	); !errors.Is(err, wantErr) {
		t.Fatalf("option error = %v", err)
	}
}

func TestNewRequestConfigAppliesTimeoutAndSecurity(t *testing.T) {
	cfg, err := NewRequestConfig(
		context.Background(),
		http.MethodGet,
		"/tweets",
		nil,
		nil,
		RequestOptionFunc(func(cfg *RequestConfig) error {
			cfg.RequestTimeout = 3 * time.Second
			cfg.APIKey = "api-key"
			cfg.BearerToken = "bearer-token"
			return nil
		}),
	)
	if err != nil {
		t.Fatal(err)
	}
	if got := cfg.Request.Header.Get("X-Stainless-Timeout"); got != "3" {
		t.Fatalf("X-Stainless-Timeout = %q", got)
	}
	if cfg.Request.Header.Get("X-API-Key") != "api-key" ||
		cfg.Request.Header.Get("Authorization") != "Bearer bearer-token" {
		t.Fatal("security headers were not applied")
	}

	override, err := NewRequestConfig(
		context.Background(),
		http.MethodGet,
		"/tweets",
		nil,
		nil,
		RequestOptionFunc(func(cfg *RequestConfig) error {
			cfg.Request.Header.Set("X-Stainless-Timeout", "custom")
			cfg.Request.Header.Set("X-API-Key", "caller")
			cfg.Request.Header.Set("Authorization", "caller")
			return nil
		}),
	)
	if err != nil {
		t.Fatal(err)
	}
	if override.Request.Header.Get("X-Stainless-Timeout") != "custom" ||
		override.Request.Header.Get("X-API-Key") != "caller" ||
		override.Request.Header.Get("Authorization") != "caller" {
		t.Fatal("caller headers were overwritten")
	}
}

func TestExecuteDecodesResponses(t *testing.T) {
	t.Run("JSON object", func(t *testing.T) {
		var dst struct {
			ID string `json:"id"`
		}
		_, err := executeWithServer(t, func(w http.ResponseWriter, _ *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			_, _ = io.WriteString(w, `{"id":"tweet-1"}`)
		}, http.MethodGet, nil, &dst)
		if err != nil {
			t.Fatal(err)
		}
		if dst.ID != "tweet-1" {
			t.Fatalf("ID = %q", dst.ID)
		}
	})

	t.Run("JSON bytes", func(t *testing.T) {
		var dst []byte
		_, err := executeWithServer(t, func(w http.ResponseWriter, _ *http.Request) {
			w.Header().Set("Content-Type", "application/problem+json")
			_, _ = io.WriteString(w, `{"ok":true}`)
		}, http.MethodGet, nil, &dst)
		if err != nil {
			t.Fatal(err)
		}
		if got := string(dst); got != `{"ok":true}` {
			t.Fatalf("body = %q", got)
		}
	})

	t.Run("plain string", func(t *testing.T) {
		var dst string
		_, err := executeWithServer(t, func(w http.ResponseWriter, _ *http.Request) {
			w.Header().Set("Content-Type", "text/plain; charset=utf-8")
			_, _ = io.WriteString(w, "timeline")
		}, http.MethodGet, nil, &dst)
		if err != nil {
			t.Fatal(err)
		}
		if dst != "timeline" {
			t.Fatalf("body = %q", dst)
		}
	})

	t.Run("plain string pointer", func(t *testing.T) {
		var dst *string
		_, err := executeWithServer(t, func(w http.ResponseWriter, _ *http.Request) {
			w.Header().Set("Content-Type", "text/plain")
			_, _ = io.WriteString(w, "followers")
		}, http.MethodGet, nil, &dst)
		if err != nil {
			t.Fatal(err)
		}
		if dst == nil || *dst != "followers" {
			t.Fatalf("body = %v", dst)
		}
	})

	t.Run("plain bytes", func(t *testing.T) {
		var dst []byte
		_, err := executeWithServer(t, func(w http.ResponseWriter, _ *http.Request) {
			w.Header().Set("Content-Type", "application/octet-stream")
			_, _ = io.WriteString(w, "media")
		}, http.MethodGet, nil, &dst)
		if err != nil || string(dst) != "media" {
			t.Fatalf("body = %q, error = %v", dst, err)
		}
	})

	t.Run("unsupported plain destination", func(t *testing.T) {
		var dst int
		_, err := executeWithServer(t, func(w http.ResponseWriter, _ *http.Request) {
			w.Header().Set("Content-Type", "text/plain")
			_, _ = io.WriteString(w, "42")
		}, http.MethodGet, nil, &dst)
		if err == nil {
			t.Fatal("unsupported destination was accepted")
		}
	})

	t.Run("invalid JSON", func(t *testing.T) {
		var dst map[string]any
		_, err := executeWithServer(t, func(w http.ResponseWriter, _ *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			_, _ = io.WriteString(w, "{")
		}, http.MethodGet, nil, &dst)
		if err == nil || !strings.Contains(err.Error(), "error parsing response json") {
			t.Fatalf("error = %v", err)
		}
	})
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
			t.Fatalf("request body = %q", content)
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
	cfg := testConfig(t, http.MethodPost, "/resource", bytes.NewBufferString("body"), &dst)
	cfg.BaseURL, _ = url.Parse(server.URL + "/")
	cfg.MaxRetries = 1
	if err := cfg.Execute(); err != nil {
		t.Fatal(err)
	}
	if attempts != 2 {
		t.Fatalf("attempts = %d", attempts)
	}

	readerCfg := testConfig(t, http.MethodPost, "/resource", bytes.NewReader([]byte("body")), nil)
	readerCfg.BaseURL, _ = url.Parse(server.URL + "/")
	readerCfg.MaxRetries = 0
	if err := readerCfg.Execute(); err != nil {
		t.Fatal(err)
	}

	closerCfg := testConfig(t, http.MethodPost, "/resource", io.NopCloser(strings.NewReader("body")), nil)
	closerCfg.BaseURL, _ = url.Parse(server.URL + "/")
	closerCfg.MaxRetries = 0
	if err := closerCfg.Execute(); err != nil {
		t.Fatal(err)
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
