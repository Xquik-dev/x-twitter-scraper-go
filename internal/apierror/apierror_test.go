// Copyright the Xquik contributors.
// SPDX-License-Identifier: Apache-2.0

package apierror

import (
	"io"
	"net/http"
	"strings"
	"sync"
	"testing"
)

func TestErrorDebuggingHelpers(t *testing.T) {
	req, err := http.NewRequest(http.MethodPost, "https://example.com/tweets", strings.NewReader(`{"text":"hello"}`))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Authorization", "Bearer secret")
	req.Header.Set("X-API-Key", "api-secret")
	req.Header.Set("Cookie", "session=secret")
	resp := &http.Response{
		StatusCode: http.StatusBadRequest,
		Header: http.Header{
			"Content-Type": {"application/json"},
			"Set-Cookie":   {"session=secret"},
		},
		Body: io.NopCloser(strings.NewReader(`{"message":"bad request"}`)),
	}
	apiErr := &Error{Request: req, Response: resp, StatusCode: resp.StatusCode}
	if err := apiErr.UnmarshalJSON([]byte(`{"message":"bad request"}`)); err != nil {
		t.Fatal(err)
	}
	if got := apiErr.RawJSON(); got != `{"message":"bad request"}` {
		t.Fatalf("RawJSON() = %q", got)
	}
	if got := apiErr.Error(); !strings.Contains(got, "400 Bad Request") {
		t.Fatalf("Error() = %q", got)
	}
	if dump := apiErr.DumpRequest(true); !strings.Contains(string(dump), `"text":"hello"`) {
		t.Fatalf("DumpRequest() = %s", dump)
	} else if strings.Contains(string(dump), "secret") {
		t.Fatalf("DumpRequest() exposed a sensitive header: %s", dump)
	}
	if dump := apiErr.DumpResponse(true); !strings.Contains(string(dump), "bad request") {
		t.Fatalf("DumpResponse() = %s", dump)
	} else if strings.Contains(string(dump), "session=secret") {
		t.Fatalf("DumpResponse() exposed a cookie: %s", dump)
	}
	if req.Header.Get("Authorization") != "Bearer secret" ||
		resp.Header.Get("Set-Cookie") != "session=secret" {
		t.Fatal("diagnostic dumps modified live headers")
	}
}

func TestErrorDebuggingHelpersAreConcurrencySafe(t *testing.T) {
	const body = `{"message":"bad request"}`
	req, err := http.NewRequest(
		http.MethodPost,
		"https://example.com/tweets",
		strings.NewReader(body),
	)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Authorization", "Bearer secret")
	resp := &http.Response{
		StatusCode: http.StatusBadRequest,
		Header:     http.Header{"Set-Cookie": {"session=secret"}},
		Body:       io.NopCloser(strings.NewReader(body)),
	}
	apiErr := &Error{Request: req, Response: resp, StatusCode: resp.StatusCode}
	if err := apiErr.UnmarshalJSON([]byte(body)); err != nil {
		t.Fatal(err)
	}

	var waitGroup sync.WaitGroup
	for range 8 {
		waitGroup.Add(1)
		go func() {
			defer waitGroup.Done()
			for range 100 {
				requestDump := string(apiErr.DumpRequest(true))
				if !strings.Contains(requestDump, body) ||
					strings.Contains(requestDump, "Bearer secret") {
					t.Errorf("unsafe request dump: %s", requestDump)
					return
				}
				responseDump := string(apiErr.DumpResponse(true))
				if !strings.Contains(responseDump, body) ||
					strings.Contains(responseDump, "session=secret") {
					t.Errorf("unsafe response dump: %s", responseDump)
					return
				}
			}
		}()
	}
	waitGroup.Wait()

	if req.Header.Get("Authorization") != "Bearer secret" ||
		resp.Header.Get("Set-Cookie") != "session=secret" {
		t.Fatal("concurrent dumps modified live headers")
	}
	requestBody, err := io.ReadAll(req.Body)
	if err != nil {
		t.Fatal(err)
	}
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	if string(requestBody) != body || string(responseBody) != body {
		t.Fatalf("concurrent dumps modified live bodies: request=%q response=%q", requestBody, responseBody)
	}
}
