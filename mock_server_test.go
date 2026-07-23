// Copyright the Xquik contributors.
// SPDX-License-Identifier: Apache-2.0

package xtwitterscraper_test

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, _ *http.Request) {
		writer.Header().Set("Content-Type", "application/json")
		_, _ = writer.Write([]byte("{}"))
	}))
	defer server.Close()
	if err := os.Setenv("TEST_API_BASE_URL", server.URL); err != nil {
		panic(err)
	}
	// Return normally so instrumented test wrappers can flush observations.
	m.Run()
}
