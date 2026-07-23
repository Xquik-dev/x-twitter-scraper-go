// Copyright the Xquik contributors.
// SPDX-License-Identifier: Apache-2.0

package option

import (
	"bytes"
	"errors"
	"io"
	"log"
	"net/http"
	"strings"
	"testing"

	"github.com/Xquik-dev/x-twitter-scraper-go/internal/requestconfig"
)

func TestDebugLoggingRedactsSensitiveHeaders(t *testing.T) {
	req, err := http.NewRequest(http.MethodPost, "https://example.com/tweets", strings.NewReader("body"))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Authorization", "Bearer secret")
	req.Header.Add("Cookie", "one")
	req.Header.Add("Cookie", "two")
	req.Header.Set("X-Public", "visible")

	requestDump, err := dumpRedactedRequest(req)
	if err != nil {
		t.Fatal(err)
	}
	if bytes.Contains(requestDump, []byte("secret")) || bytes.Contains(requestDump, []byte("Cookie: one")) {
		t.Fatalf("request dump exposed a sensitive value: %s", requestDump)
	}
	if !bytes.Contains(requestDump, []byte("Authorization: ***")) {
		t.Fatalf("request dump did not redact authorization: %s", requestDump)
	}
	if req.Header.Get("Authorization") != "Bearer secret" {
		t.Fatal("request headers were modified")
	}

	resp := &http.Response{
		StatusCode: http.StatusOK,
		Header: http.Header{
			"Set-Cookie": {"session=secret"},
			"X-Public":   {"visible"},
		},
		Body: io.NopCloser(strings.NewReader("response")),
	}
	responseDump, err := dumpRedactedResponse(resp)
	if err != nil {
		t.Fatal(err)
	}
	if bytes.Contains(responseDump, []byte("session=secret")) {
		t.Fatalf("response dump exposed a cookie: %s", responseDump)
	}
	if resp.Header.Get("Set-Cookie") != "session=secret" {
		t.Fatal("response headers were modified")
	}

	public := http.Header{"X-Public": {"visible"}}
	if got := redactDebugHeaders(public); got.Get("X-Public") != "visible" {
		t.Fatalf("public headers changed: %v", got)
	}
}

func TestWithDebugLog(t *testing.T) {
	var output bytes.Buffer
	logger := log.New(&output, "", 0)
	req, err := http.NewRequest(http.MethodGet, "https://example.com/tweets", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("X-API-Key", "secret")
	cfg := &requestconfig.RequestConfig{Request: req}
	if err := WithDebugLog(logger).Apply(cfg); err != nil {
		t.Fatal(err)
	}
	if len(cfg.Middlewares) != 1 {
		t.Fatalf("Middlewares = %d", len(cfg.Middlewares))
	}

	resp, err := cfg.Middlewares[0](req, func(*http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: http.StatusOK,
			Header:     http.Header{"Content-Type": {"application/json"}},
			Body:       io.NopCloser(strings.NewReader("{}")),
		}, nil
	})
	if err != nil || resp.StatusCode != http.StatusOK {
		t.Fatalf("middleware response = %v, %v", resp, err)
	}
	if strings.Contains(output.String(), "secret") {
		t.Fatalf("debug log exposed a credential: %s", output.String())
	}
	if !strings.Contains(output.String(), "Request Content:") ||
		!strings.Contains(output.String(), "Response Content:") {
		t.Fatalf("debug log is incomplete: %s", output.String())
	}

	wantErr := errors.New("transport failed")
	_, err = cfg.Middlewares[0](req, func(*http.Request) (*http.Response, error) {
		return nil, wantErr
	})
	if !errors.Is(err, wantErr) {
		t.Fatalf("middleware error = %v", err)
	}
}
