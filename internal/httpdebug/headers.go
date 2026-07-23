// Copyright the Xquik contributors.
// SPDX-License-Identifier: Apache-2.0

// Package httpdebug provides safe helpers for HTTP diagnostic output.
package httpdebug

import "net/http"

var sensitiveHeaders = []string{
	"authorization",
	"api-key",
	"x-api-key",
	"cookie",
	"set-cookie",
}

// RedactHeaders replaces sensitive values without changing the input.
func RedactHeaders(headers http.Header) http.Header {
	var redacted http.Header
	for _, name := range sensitiveHeaders {
		values := headers.Values(name)
		if len(values) == 0 {
			continue
		}
		if redacted == nil {
			redacted = headers.Clone()
		}
		redacted.Del(name)
		for range values {
			redacted.Add(name, "***")
		}
	}
	if redacted == nil {
		return headers
	}
	return redacted
}
