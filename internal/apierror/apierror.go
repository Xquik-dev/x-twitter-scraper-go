// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package apierror

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	"strings"

	"github.com/Xquik-dev/x-twitter-scraper-go/internal/apijson"
	"github.com/Xquik-dev/x-twitter-scraper-go/internal/httpdebug"
	"github.com/Xquik-dev/x-twitter-scraper-go/packages/respjson"
)

// Error represents an error that originates from the API, i.e. when a request is
// made and the API returns a response with a HTTP status code. Other errors are
// not wrapped by this SDK.
type Error struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	StatusCode int
	Request    *http.Request
	Response   *http.Response
}

// Returns the unmodified JSON received from the API
func (r Error) RawJSON() string { return r.JSON.raw }
func (r *Error) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func (r *Error) Error() string {
	// Attempt to re-populate the response body
	return fmt.Sprintf("%s %q: %d %s %s", r.Request.Method, r.Request.URL, r.Response.StatusCode, http.StatusText(r.Response.StatusCode), r.JSON.raw)
}

func (r *Error) DumpRequest(body bool) []byte {
	requestCopy := *r.Request
	requestCopy.Header = httpdebug.RedactHeaders(r.Request.Header).Clone()
	if body && r.Request.GetBody != nil {
		if bodyCopy, err := r.Request.GetBody(); err == nil {
			requestCopy.Body = bodyCopy
		} else {
			body = false
		}
	} else if body {
		body = false
	}
	out, _ := httputil.DumpRequestOut(&requestCopy, body)
	return out
}

func (r *Error) DumpResponse(body bool) []byte {
	responseCopy := *r.Response
	responseCopy.Header = httpdebug.RedactHeaders(r.Response.Header).Clone()
	if body && r.JSON.raw != "" {
		responseCopy.Body = io.NopCloser(strings.NewReader(r.JSON.raw))
		responseCopy.ContentLength = int64(len(r.JSON.raw))
	} else if body {
		body = false
	}
	out, _ := httputil.DumpResponse(&responseCopy, body)
	return out
}
