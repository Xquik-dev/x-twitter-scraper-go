// Copyright the Xquik contributors.
// SPDX-License-Identifier: Apache-2.0

package httpdebug

import (
	"net/http"
	"reflect"
	"testing"
)

func TestRedactHeadersPreservesInput(t *testing.T) {
	headers := http.Header{
		"Authorization": {"Bearer secret"},
		"Cookie":        {"one", "two"},
		"X-Public":      {"visible"},
	}
	original := headers.Clone()

	redacted := RedactHeaders(headers)
	if !reflect.DeepEqual(headers, original) {
		t.Fatalf("input headers changed: %v", headers)
	}
	if got := redacted.Values("Cookie"); !reflect.DeepEqual(got, []string{"***", "***"}) {
		t.Fatalf("redacted cookies = %v", got)
	}
	if redacted.Get("Authorization") != "***" || redacted.Get("X-Public") != "visible" {
		t.Fatalf("redacted headers = %v", redacted)
	}
}

func TestRedactHeadersReturnsPublicHeaders(t *testing.T) {
	for _, headers := range []http.Header{nil, {"X-Public": {"visible"}, "Authorization": nil, "Cookie": {}}} {
		if got := RedactHeaders(headers); !reflect.DeepEqual(got, headers) {
			t.Fatalf("public or empty headers changed: %v", got)
		}
	}
}

func TestRedactHeadersMatchesEveryKeyCase(t *testing.T) {
	for _, name := range []string{"authorization", "AUTHORIZATION", "proxy-authorization", "PROXY-AUTHORIZATION", "api-key", "API-KEY", "x-api-key", "X-API-KEY", "cookie", "COOKIE", "set-cookie", "SET-COOKIE"} {
		t.Run(name, func(t *testing.T) {
			headers := http.Header{name: {"synthetic-one", "synthetic-two"}, "X-Public": {"visible"}}
			original := headers.Clone()
			redacted := RedactHeaders(headers)
			if !reflect.DeepEqual(redacted[name], []string{"***", "***"}) {
				t.Fatalf("sensitive header %s was not fully redacted", name)
			}
			if !reflect.DeepEqual(headers, original) || redacted.Get("X-Public") != "visible" {
				t.Fatal("redaction changed input or public metadata")
			}
		})
	}
}
