// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

// Package testutil checks for the mock API server that the SDK's tests run
// against.
package testutil

import (
	"net/http"
	"os"
	"strconv"
	"testing"
)

func CheckTestServer(t *testing.T, url string) bool {
	request, err := http.NewRequestWithContext(t.Context(), http.MethodGet, url, nil)
	if err != nil {
		t.Fatal(err)
	}
	if response, err := http.DefaultClient.Do(request); err != nil {
		const skipMockTests = "SKIP_MOCK_TESTS"
		if str, ok := os.LookupEnv(skipMockTests); ok {
			skip, err := strconv.ParseBool(str)
			if err != nil {
				t.Fatalf("strconv.ParseBool(os.LookupEnv(%s)) failed: %s", skipMockTests, err)
			}
			if skip {
				t.Skip("The test will not run without a mock server running against your OpenAPI spec")
				return false
			}
			t.Errorf("The test will not run without a mock server running against your OpenAPI spec. You can set the environment variable %s to true to skip running any tests that require the mock server", skipMockTests)
			return false
		}
	} else if err := response.Body.Close(); err != nil {
		t.Errorf("Close test server response: %v", err)
		return false
	}
	return true
}
