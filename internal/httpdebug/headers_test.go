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
	headers := http.Header{"X-Public": {"visible"}}
	if got := RedactHeaders(headers); !reflect.DeepEqual(got, headers) {
		t.Fatalf("public headers changed: %v", got)
	}
}
