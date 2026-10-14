// Copyright the Xquik contributors.
// SPDX-License-Identifier: Apache-2.0

package httpdebug

import "net/http"

func RedactHeaders(headers http.Header) http.Header {
	var redacted http.Header
	for name, values := range headers {
		switch http.CanonicalHeaderKey(name) {
		case "Authorization", "Proxy-Authorization", "Api-Key", "X-Api-Key", "Cookie", "Set-Cookie":
		default:
			continue
		}
		if len(values) == 0 {
			continue
		}
		if redacted == nil {
			redacted = headers.Clone()
		}
		for index := range values {
			redacted[name][index] = "***"
		}
	}
	if redacted == nil {
		return headers
	}
	return redacted
}
