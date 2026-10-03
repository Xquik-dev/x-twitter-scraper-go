// Copyright the Xquik contributors.
// SPDX-License-Identifier: Apache-2.0

package xtwitterscraper_test

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func runMockTests(m *testing.M) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, _ *http.Request) {
		writer.Header().Set("Content-Type", "application/json")
		if _, err := writer.Write([]byte("{}")); err != nil {
			panic(err)
		}
	}))
	defer server.Close()
	if err := os.Setenv("TEST_API_BASE_URL", server.URL); err != nil {
		panic(err)
	}
	m.Run()
}
