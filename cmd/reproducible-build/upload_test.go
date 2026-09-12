// Copyright the Xquik contributors.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func TestUploadDoesNotExposeCredentials(t *testing.T) {
	script, err := filepath.Abs("../../scripts/utils/upload-artifact.sh")
	if err != nil {
		t.Fatal(err)
	}
	for _, status := range []int{http.StatusOK, http.StatusForbidden} {
		t.Run(fmt.Sprint(status), func(t *testing.T) {
			directory := t.TempDir()
			if err := os.WriteFile(filepath.Join(directory, "fixture.go"), []byte("package fixture\n"), 0o644); err != nil {
				t.Fatal(err)
			}
			uploads := make(chan []byte, 1)
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.Method == http.MethodPost {
					if r.Header.Get("Authorization") != "Bearer credential-fixture" {
						t.Error("missing authorization")
					}
					fmt.Fprintf(w, `{"url":"http://%s/upload?signed-fixture"}`, r.Host)
					return
				}
				uploaded, readErr := io.ReadAll(r.Body)
				uploads <- uploaded
				if readErr != nil || r.URL.RawQuery != "signed-fixture" {
					t.Errorf("upload: %v, query %q", readErr, r.URL.RawQuery)
				}
				w.WriteHeader(status)
			}))
			defer server.Close()
			command := exec.Command("bash", "-x", script)
			command.Dir = directory
			command.Env = append(os.Environ(), "URL="+server.URL, "AUTH=credential-fixture", "SHA=fixture")
			output, err := command.CombinedOutput()
			var uploaded []byte
			select {
			case uploaded = <-uploads:
			default:
			}
			if (err == nil) != (status == http.StatusOK) || !bytes.HasPrefix(uploaded, []byte("PK")) {
				t.Fatalf("upload: %v, archive bytes %d\n%s", err, len(uploaded), output)
			}
			if strings.Contains(string(output), "credential-fixture") || strings.Contains(string(output), "signed-fixture") {
				t.Fatal("upload output exposed credentials")
			}
		})
	}
}
