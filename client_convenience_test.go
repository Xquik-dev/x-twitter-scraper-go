// Copyright the Xquik contributors.
// SPDX-License-Identifier: Apache-2.0

package xtwitterscraper_test

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Xquik-dev/x-twitter-scraper-go"
	"github.com/Xquik-dev/x-twitter-scraper-go/option"
)

func TestClientConvenienceMethods(t *testing.T) {
	var methods []string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		methods = append(methods, req.Method)
		w.Header().Set("Content-Type", "application/json")
		_, _ = io.WriteString(w, "{}")
	}))
	defer server.Close()

	client := xtwitterscraper.NewClient(
		option.WithBaseURL(server.URL),
		option.WithMaxRetries(0),
	)
	ctx := context.Background()
	var dst map[string]any
	calls := []func() error{
		func() error { return client.Get(ctx, "/get", nil, &dst) },
		func() error { return client.Post(ctx, "/post", strings.NewReader("{}"), &dst) },
		func() error { return client.Put(ctx, "/put", []byte("{}"), &dst) },
		func() error { return client.Patch(ctx, "/patch", map[string]bool{"ok": true}, &dst) },
		func() error { return client.Delete(ctx, "/delete", nil, &dst) },
		func() error { return client.Execute(ctx, http.MethodHead, "/head", nil, nil) },
	}
	for _, call := range calls {
		if err := call(); err != nil {
			t.Fatal(err)
		}
	}
	want := []string{
		http.MethodGet,
		http.MethodPost,
		http.MethodPut,
		http.MethodPatch,
		http.MethodDelete,
		http.MethodHead,
	}
	if strings.Join(methods, ",") != strings.Join(want, ",") {
		t.Fatalf("methods = %v", methods)
	}
}

func TestDefaultClientOptionsFromEnvironment(t *testing.T) {
	t.Setenv("X_TWITTER_SCRAPER_BASE_URL", "https://example.com/v1/")
	t.Setenv("X_TWITTER_SCRAPER_API_KEY", "api-key")
	t.Setenv("X_TWITTER_SCRAPER_BEARER_TOKEN", "bearer-token")
	t.Setenv("X_TWITTER_SCRAPER_CUSTOM_HEADERS", "X-Trace: trace-id\ninvalid\nX-Agent: customer")

	client := xtwitterscraper.NewClient()
	if err := client.Get(context.Background(), "/resource", nil, nil, option.WithHTTPClient(
		testHTTPDoer(func(req *http.Request) (*http.Response, error) {
			if req.Header.Get("X-Trace") != "trace-id" ||
				req.Header.Get("X-Agent") != "customer" ||
				req.Header.Get("X-API-Key") != "api-key" ||
				req.Header.Get("Authorization") != "Bearer bearer-token" {
				t.Fatalf("environment options were not applied: %v", req.Header)
			}
			return &http.Response{
				StatusCode: http.StatusNoContent,
				Header:     http.Header{},
				Body:       http.NoBody,
			}, nil
		}),
	)); err != nil {
		t.Fatal(err)
	}
}

type testHTTPDoer func(*http.Request) (*http.Response, error)

func (do testHTTPDoer) Do(req *http.Request) (*http.Response, error) {
	return do(req)
}
