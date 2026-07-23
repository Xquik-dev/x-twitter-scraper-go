// Copyright the Xquik contributors.
// SPDX-License-Identifier: Apache-2.0

package option

import (
	"bytes"
	"io"
	"net/http"
	"net/url"
	"strings"
	"testing"
	"time"

	"github.com/Xquik-dev/x-twitter-scraper-go/internal/requestconfig"
)

type testHTTPDoer func(*http.Request) (*http.Response, error)

func (do testHTTPDoer) Do(req *http.Request) (*http.Response, error) {
	return do(req)
}

func newOptionTestConfig(t *testing.T) *requestconfig.RequestConfig {
	t.Helper()
	req, err := http.NewRequest(http.MethodGet, "https://example.com/items?keep=1", nil)
	if err != nil {
		t.Fatal(err)
	}
	return &requestconfig.RequestConfig{Request: req, HTTPClient: http.DefaultClient}
}

func applyOption(t *testing.T, cfg *requestconfig.RequestConfig, opt RequestOption) error {
	t.Helper()
	return opt.Apply(cfg)
}

func TestRequestOptions(t *testing.T) {
	t.Run("base URL", func(t *testing.T) {
		cfg := newOptionTestConfig(t)
		if err := applyOption(t, cfg, WithBaseURL("https://api.example.com/v1")); err != nil {
			t.Fatal(err)
		}
		if got := cfg.BaseURL.String(); got != "https://api.example.com/v1/" {
			t.Fatalf("BaseURL = %q", got)
		}
		if err := applyOption(t, cfg, WithBaseURL("://invalid")); err == nil {
			t.Fatal("WithBaseURL() accepted an invalid URL")
		}
	})

	t.Run("HTTP clients", func(t *testing.T) {
		cfg := newOptionTestConfig(t)
		client := &http.Client{}
		if err := applyOption(t, cfg, WithHTTPClient(client)); err != nil {
			t.Fatal(err)
		}
		if cfg.HTTPClient != client || cfg.CustomHTTPDoer != nil {
			t.Fatal("native HTTP client was not selected")
		}

		custom := testHTTPDoer(func(*http.Request) (*http.Response, error) {
			return nil, nil
		})
		if err := applyOption(t, cfg, WithHTTPClient(custom)); err != nil {
			t.Fatal(err)
		}
		if cfg.CustomHTTPDoer == nil {
			t.Fatal("custom HTTP doer was not selected")
		}
		if err := applyOption(t, cfg, WithHTTPClient(nil)); err == nil {
			t.Fatal("WithHTTPClient() accepted nil")
		}
	})

	t.Run("middleware and retry count", func(t *testing.T) {
		cfg := newOptionTestConfig(t)
		middleware := func(req *http.Request, next MiddlewareNext) (*http.Response, error) {
			return next(req)
		}
		if err := applyOption(t, cfg, WithMiddleware(middleware)); err != nil {
			t.Fatal(err)
		}
		if len(cfg.Middlewares) != 1 {
			t.Fatalf("Middlewares = %d", len(cfg.Middlewares))
		}
		if err := applyOption(t, cfg, WithMaxRetries(4)); err != nil {
			t.Fatal(err)
		}
		if cfg.MaxRetries != 4 {
			t.Fatalf("MaxRetries = %d", cfg.MaxRetries)
		}
		defer func() {
			if recover() == nil {
				t.Fatal("WithMaxRetries() did not panic for a negative value")
			}
		}()
		_ = WithMaxRetries(-1)
	})

	t.Run("headers and query", func(t *testing.T) {
		cfg := newOptionTestConfig(t)
		options := []RequestOption{
			WithHeader("X-Test", "one"),
			WithHeaderAdd("X-Test", "two"),
			WithQuery("q", "tweet search"),
			WithQueryAdd("q", "timeline"),
		}
		for _, opt := range options {
			if err := applyOption(t, cfg, opt); err != nil {
				t.Fatal(err)
			}
		}
		if got := cfg.Request.Header.Values("X-Test"); len(got) != 2 {
			t.Fatalf("X-Test values = %v", got)
		}
		if got := cfg.Request.URL.Query()["q"]; len(got) != 2 {
			t.Fatalf("q values = %v", got)
		}
		if err := applyOption(t, cfg, WithHeaderDel("X-Test")); err != nil {
			t.Fatal(err)
		}
		if err := applyOption(t, cfg, WithQueryDel("q")); err != nil {
			t.Fatal(err)
		}
		if cfg.Request.Header.Get("X-Test") != "" || cfg.Request.URL.Query().Has("q") {
			t.Fatal("header or query deletion failed")
		}
	})

	t.Run("JSON mutation", func(t *testing.T) {
		cfg := newOptionTestConfig(t)
		if err := applyOption(t, cfg, WithJSONSet("tweet.text", "hello")); err != nil {
			t.Fatal(err)
		}
		if err := applyOption(t, cfg, WithJSONSet("tweet.id", "123")); err != nil {
			t.Fatal(err)
		}
		if err := applyOption(t, cfg, WithJSONDel("tweet.id")); err != nil {
			t.Fatal(err)
		}
		body, err := io.ReadAll(cfg.Body)
		if err != nil {
			t.Fatal(err)
		}
		if got := string(body); got != `{"tweet":{"text":"hello"}}` {
			t.Fatalf("Body = %q", got)
		}

		cfg.Body = strings.NewReader("{}")
		if err := applyOption(t, cfg, WithJSONSet("x", true)); err == nil {
			t.Fatal("WithJSONSet() accepted a non-buffer body")
		}
		if err := applyOption(t, cfg, WithJSONDel("x")); err == nil {
			t.Fatal("WithJSONDel() accepted a non-buffer body")
		}
	})

	t.Run("response destinations and request bodies", func(t *testing.T) {
		cfg := newOptionTestConfig(t)
		var response *http.Response
		var bodyDestination string
		if err := applyOption(t, cfg, WithResponseBodyInto(&bodyDestination)); err != nil {
			t.Fatal(err)
		}
		if err := applyOption(t, cfg, WithResponseInto(&response)); err != nil {
			t.Fatal(err)
		}
		if cfg.ResponseBodyInto != &bodyDestination || cfg.ResponseInto != &response {
			t.Fatal("response destination was not retained")
		}

		if err := applyOption(t, cfg, WithRequestBody("text/plain", strings.NewReader("reader"))); err != nil {
			t.Fatal(err)
		}
		if cfg.Request.Header.Get("Content-Type") != "text/plain" {
			t.Fatal("reader content type was not set")
		}
		if err := applyOption(t, cfg, WithRequestBody("application/octet-stream", []byte("bytes"))); err != nil {
			t.Fatal(err)
		}
		if _, ok := cfg.Body.(*bytes.Buffer); !ok {
			t.Fatalf("Body type = %T", cfg.Body)
		}
		if err := applyOption(t, cfg, WithRequestBody("text/plain", 42)); err == nil {
			t.Fatal("WithRequestBody() accepted an unsupported body")
		}
	})

	t.Run("timeout environment and credentials", func(t *testing.T) {
		cfg := newOptionTestConfig(t)
		options := []RequestOption{
			WithRequestTimeout(3 * time.Second),
			WithEnvironmentProduction(),
			WithAPIKey("api-key"),
			WithBearerToken("bearer-token"),
		}
		for _, opt := range options {
			if err := applyOption(t, cfg, opt); err != nil {
				t.Fatal(err)
			}
		}
		if cfg.RequestTimeout != 3*time.Second {
			t.Fatalf("RequestTimeout = %s", cfg.RequestTimeout)
		}
		wantURL, _ := url.Parse("https://xquik.com/api/v1/")
		if cfg.DefaultBaseURL.String() != wantURL.String() {
			t.Fatalf("DefaultBaseURL = %s", cfg.DefaultBaseURL)
		}
		if cfg.APIKey != "api-key" || cfg.BearerToken != "bearer-token" {
			t.Fatal("credentials were not retained")
		}
	})
}
