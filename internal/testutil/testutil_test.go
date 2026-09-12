// Copyright the Xquik contributors.
// SPDX-License-Identifier: Apache-2.0

package testutil

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestCheckTestServerClosesResponse(t *testing.T) {
	closed := make(chan struct{})
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.(http.Flusher).Flush()
		<-r.Context().Done()
		close(closed)
	}))
	defer server.Close()
	defer server.CloseClientConnections()
	if !CheckTestServer(t, server.URL) {
		t.Fatal("running server rejected")
	}
	select {
	case <-closed:
	case <-time.After(time.Second):
		t.Fatal("probe left response body open")
	}
}
