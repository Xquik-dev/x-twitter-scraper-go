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

func TestRetryHelpers(t *testing.T) {
	req, _ := http.NewRequest(http.MethodPost, "https://example.com", strings.NewReader("body"))
	req.GetBody = nil
	if shouldRetry(req, nil) {
		t.Fatal("request with an unrecoverable body was retried")
	}
	req.GetBody = func() (io.ReadCloser, error) {
		return io.NopCloser(strings.NewReader("body")), nil
	}

	tests := []struct {
		name string
		res  *http.Response
		want bool
	}{
		{name: "connection error", res: nil, want: true},
		{name: "explicit yes", res: &http.Response{StatusCode: 200, Header: http.Header{"X-Should-Retry": {"true"}}}, want: true},
		{name: "explicit no", res: &http.Response{StatusCode: 500, Header: http.Header{"X-Should-Retry": {"false"}}}, want: false},
		{name: "timeout", res: &http.Response{StatusCode: 408, Header: http.Header{}}, want: true},
		{name: "conflict", res: &http.Response{StatusCode: 409, Header: http.Header{}}, want: true},
		{name: "rate limit", res: &http.Response{StatusCode: 429, Header: http.Header{}}, want: true},
		{name: "server error", res: &http.Response{StatusCode: 503, Header: http.Header{}}, want: true},
		{name: "client error", res: &http.Response{StatusCode: 400, Header: http.Header{}}, want: false},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := shouldRetry(req, test.res); got != test.want {
				t.Fatalf("shouldRetry() = %t", got)
			}
		})
	}

	if _, ok := parseRetryAfterHeader(nil); ok {
		t.Fatal("nil response returned retry delay")
	}
	headers := []struct {
		value string
		want  time.Duration
	}{
		{value: "250", want: 250 * time.Millisecond},
		{value: "0.5", want: 500 * time.Microsecond},
	}
	for _, test := range headers {
		res := &http.Response{Header: http.Header{"Retry-After-Ms": {test.value}}}
		if got, ok := parseRetryAfterHeader(res); !ok || got != test.want {
			t.Fatalf("parseRetryAfterHeader() = %s, %t", got, ok)
		}
	}
	date := time.Now().Add(2 * time.Second).UTC().Format(time.RFC1123)
	if got, ok := parseRetryAfterHeader(&http.Response{Header: http.Header{"Retry-After": {date}}}); !ok ||
		got <= 0 || got > 3*time.Second {
		t.Fatalf("date retry delay = %s, %t", got, ok)
	}
	if _, ok := parseRetryAfterHeader(&http.Response{Header: http.Header{"Retry-After": {"invalid"}}}); ok {
		t.Fatal("invalid retry delay was accepted")
	}
	if got := retryDelay(&http.Response{Header: http.Header{"Retry-After-Ms": {"0"}}}, 0); got != 0 {
		t.Fatalf("retryDelay() = %s", got)
	}
	if got := retryDelay(nil, 20); got < 6*time.Second || got > 8*time.Second {
		t.Fatalf("capped retryDelay() = %s", got)
	}
}

func TestRequestConfigUtilities(t *testing.T) {
	if !isBeforeContextDeadline(time.Now().Add(time.Hour), context.Background()) {
		t.Fatal("context without deadline rejected time")
	}
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Hour))
	defer cancel()
	if isBeforeContextDeadline(time.Now().Add(2*time.Hour), ctx) {
		t.Fatal("time after deadline was accepted")
	}

	stopped := false
	body := &bodyWithTimeout{
		stop: func() { stopped = true },
		rc:   io.NopCloser(strings.NewReader("body")),
	}
	content, err := io.ReadAll(body)
	if err != nil || string(content) != "body" {
		t.Fatalf("body = %q, error = %v", content, err)
	}
	if err := body.Close(); err != nil || !stopped {
		t.Fatalf("Close() error = %v, stopped = %t", err, stopped)
	}

	wantErr := errors.New("read failed")
	failingBody := &bodyWithTimeout{stop: func() {}, rc: errorReadCloser{err: wantErr}}
	if _, err := failingBody.Read(make([]byte, 1)); !errors.Is(err, wantErr) {
		t.Fatalf("Read() error = %v", err)
	}

	if clone := (*RequestConfig)(nil).Clone(context.Background()); clone != nil {
		t.Fatal("nil config clone was non-nil")
	}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	}))
	defer server.Close()
	cfg := testConfig(t, http.MethodPost, "/resource", bytes.NewBufferString("body"), nil)
	cfg.BaseURL, _ = url.Parse(server.URL + "/")
	if err := cfg.Execute(); err != nil {
		t.Fatal(err)
	}
	clone := cfg.Clone(context.Background())
	if clone == nil || clone.Request == cfg.Request {
		t.Fatal("request config was not cloned")
	}

	wantApplyErr := errors.New("apply failed")
	if err := cfg.Apply(
		RequestOptionFunc(func(*RequestConfig) error { return nil }),
		RequestOptionFunc(func(*RequestConfig) error { return wantApplyErr }),
	); !errors.Is(err, wantApplyErr) {
		t.Fatalf("Apply() error = %v", err)
	}

	pre, err := PreRequestOptions(
		RequestOptionFunc(func(cfg *RequestConfig) error {
			cfg.APIKey = "ignored"
			return nil
		}),
		PreRequestOptionFunc(func(cfg *RequestConfig) error {
			cfg.APIKey = "applied"
			return nil
		}),
	)
	if err != nil || pre.APIKey != "applied" {
		t.Fatalf("PreRequestOptions() = %+v, %v", pre, err)
	}
	if _, err := PreRequestOptions(PreRequestOptionFunc(func(*RequestConfig) error {
		return wantApplyErr
	})); !errors.Is(err, wantApplyErr) {
		t.Fatalf("PreRequestOptions() error = %v", err)
	}
}

func TestSecurityOptions(t *testing.T) {
	req, _ := http.NewRequest(http.MethodGet, "https://example.com", nil)
	cfg := &RequestConfig{Request: req, APIKey: "key", BearerToken: "token"}

	options := []RequestOption{
		WithSecurity(Security{APIKey: true, OAuthBearer: true}),
		WithAPIKeySecurity(),
		WithOAuthBearerSecurity(),
	}
	for _, option := range options {
		if err := option.Apply(cfg); err != nil {
			t.Fatal(err)
		}
	}

	cfg.Security = Security{APIKey: true, OAuthBearer: true}
	ApplySecurity(*cfg)
	if req.Header.Get("X-API-Key") != "key" ||
		req.Header.Get("Authorization") != "Bearer token" {
		t.Fatal("security headers were not applied")
	}

	defaultURL := WithDefaultBaseURL("https://example.com/v1/")
	if err := defaultURL.Apply(cfg); err != nil || cfg.DefaultBaseURL == nil {
		t.Fatalf("WithDefaultBaseURL() = %v", err)
	}
	if err := WithDefaultBaseURL("://invalid").Apply(cfg); err == nil {
		t.Fatal("invalid default base URL was accepted")
	}
}

func TestPlatformNormalization(t *testing.T) {
	operatingSystems := map[string]string{
		"ios":     "iOS",
		"android": "Android",
		"darwin":  "MacOS",
		"window":  "Windows",
		"freebsd": "FreeBSD",
		"openbsd": "OpenBSD",
		"linux":   "Linux",
		"plan9":   "Other:plan9",
	}
	for input, want := range operatingSystems {
		if got := normalizeOS(input); got != want {
			t.Fatalf("normalizeOS(%q) = %q, want %q", input, got, want)
		}
	}
	architectures := map[string]string{
		"386":     "x32",
		"amd64":   "x64",
		"arm":     "arm",
		"arm64":   "arm64",
		"riscv64": "other:riscv64",
	}
	for input, want := range architectures {
		if got := normalizeArchitecture(input); got != want {
			t.Fatalf("normalizeArchitecture(%q) = %q, want %q", input, got, want)
		}
	}
}

func TestMiddlewareAndExecuteNewRequest(t *testing.T) {
	req, _ := http.NewRequest(http.MethodGet, "https://example.com", nil)
	called := false
	handler := applyMiddleware(
		func(req *http.Request, next middlewareNext) (*http.Response, error) {
			called = true
			return next(req)
		},
		func(*http.Request) (*http.Response, error) {
			return &http.Response{StatusCode: http.StatusNoContent}, nil
		},
	)
	if _, err := handler(req); err != nil || !called {
		t.Fatalf("middleware result = %v, called = %t", err, called)
	}

	if err := ExecuteNewRequest(
		context.Background(),
		"bad\nmethod",
		"/resource",
		nil,
		nil,
	); err == nil {
		t.Fatal("ExecuteNewRequest() accepted an invalid method")
	}
}
