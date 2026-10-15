// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

package apicheck

import (
	"go/token"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"testing"
	"testing/fstest"

	"github.com/Xquik-dev/x-twitter-scraper-go/internal/requestconfig"
)

const services = `package sdk

func (r *XTweetService) Get(ctx context.Context, id string) (res *Tweet, err error) {
	path := fmt.Sprintf("x/tweets/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res)
	return
}

func (r *XTweetService) Search(ctx context.Context) (res *Page, err error) {
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, "x/tweets/search", nil, &res)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, "x/tweets/search", nil, &res)
	return
}

func (r XDmService) Delete(ctx context.Context, id string, params Params) (res *Action, err error) {
	path := fmt.Sprintf("x/dm/%s/messages/%s", url.PathEscape(params.UserID), url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, params, &res)
	return
}

func (r *Client) Execute(ctx context.Context, method string, path string) error {
	return requestconfig.ExecuteNewRequest(ctx, method, path, nil, nil)
}

func (Store) Load() {}

func helper() { requestconfig.Other() }
`

func writeFile(t *testing.T, content string) string {
	t.Helper()
	file := filepath.Join(t.TempDir(), "openapi.json")
	if err := os.WriteFile(file, []byte(content), 0o600); err != nil {
		t.Fatal(err)
	}
	return file
}

func TestLoadSpecReadsOperations(t *testing.T) {
	spec := `{"servers": [{"url": "https://xquik.com/"}], "paths": {
		"/api/v1/x/tweets/{id}": {"get": {"operationId": "getTweet"}, "parameters": []},
		"/ask": {"servers": [{"url": "https://xquik.com"}], "post": {"servers": []}},
		"/x/media": {"post": {"servers": [{"url": "https://upload.xquik.com/api/v1"}]}}
	}}`
	loaded, err := LoadSpec(writeFile(t, spec), http.DefaultClient)
	if err != nil {
		t.Fatal(err)
	}
	var labels []string
	for _, operation := range loaded.Operations {
		labels = append(labels, operation.Label())
	}
	slices.Sort(labels)
	want := []string{"GET /api/v1/x/tweets/{id} (getTweet)", "POST /api/v1/x/media", "POST /ask"}
	if !slices.Equal(labels, want) {
		t.Fatalf("labels = %q, want %q", labels, want)
	}
	if _, ok := loaded.Operations["GET /api/v1/x/tweets/{}"]; !ok {
		t.Fatalf("operations = %v", loaded.Operations)
	}
}

func TestLoadSpecFetchesURLs(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/missing" {
			http.NotFound(w, r)
			return
		}
		if _, err := w.Write([]byte(`{"paths": {"/x": {"get": {}}}}`)); err != nil {
			t.Error(err)
		}
	}))
	defer server.Close()
	if loaded, err := LoadSpec(server.URL+"/openapi.json", server.Client()); err != nil || len(loaded.Operations) != 1 {
		t.Fatalf("loaded = %v, err = %v", loaded, err)
	}
	for _, source := range []string{server.URL + "/missing", "http://[::1]:namedport", "http://127.0.0.1:1/openapi.json"} {
		want := "Could not read the API spec from " + source + ". Check network access, or set XQUIK_OPENAPI_SPEC to a local copy of https://xquik.com/openapi.json."
		if _, err := LoadSpec(source, server.Client()); err == nil || err.Error() != want {
			t.Fatalf("source %s err = %v", source, err)
		}
	}
}

func TestLoadSpecRejectsUnreadableSpecs(t *testing.T) {
	for name, content := range map[string]string{
		"not JSON":             "{",
		"no paths":             `{"paths": {}}`,
		"path item not object": `{"paths": {"/x": []}}`,
		"method not object":    `{"paths": {"/x": {"get": "GET"}}}`,
		"servers not a list":   `{"servers": "https://xquik.com", "paths": {"/x": {"get": {}}}}`,
		"bad server URL":       `{"servers": [{"url": "http://[::1"}], "paths": {"/x": {"get": {}}}}`,
	} {
		t.Run(name, func(t *testing.T) {
			file := writeFile(t, content)
			if _, err := LoadSpec(file, http.DefaultClient); err == nil || !strings.HasPrefix(err.Error(), "Could not read the API spec from "+file+".") {
				t.Fatalf("err = %v", err)
			}
		})
	}
}

func TestSourceDefaultsToThePublicSpec(t *testing.T) {
	if Source("") != "https://xquik.com/openapi.json" || Source("local.json") != "local.json" {
		t.Fatalf("Source() = %q, %q", Source(""), Source("local.json"))
	}
	if !KnownMethod("PATCH") || KnownMethod("FETCH") {
		t.Fatal("KnownMethod accepted the wrong methods")
	}
}

func TestBasePathReadsTheDefaultBaseURL(t *testing.T) {
	if base, err := Fields.BasePath(requestconfig.WithDefaultBaseURL("https://xquik.com/api/v1/")); err != nil || base != "/api/v1" {
		t.Fatalf("base = %q, err = %v", base, err)
	}
	if _, err := Coverage.BasePath(requestconfig.WithDefaultBaseURL("://bad")); err == nil || !strings.HasPrefix(err.Error(), "The coverage check could not read") {
		t.Fatalf("err = %v", err)
	}
}

func TestSDKOperationsReadServiceRequests(t *testing.T) {
	root := fstest.MapFS{
		"tweet.go":      {Data: []byte(services)},
		"tweet_test.go": {Data: []byte("package sdk\n")},
		"README.md":     {Data: []byte("# SDK\n")},
		"shared/a.go":   {Data: []byte("package shared\n")},
	}
	found, unreadable, err := Fields.SDKOperations(root, "/api/v1")
	if err != nil || len(unreadable) > 0 {
		t.Fatalf("err = %v, unreadable = %q", err, unreadable)
	}
	var keys []string
	for key, operation := range found {
		keys = append(keys, key)
		if operation.Function == nil || operation.Call == nil || operation.Position.Filename != "tweet.go" {
			t.Fatalf("operation %s = %+v", key, operation)
		}
	}
	slices.Sort(keys)
	want := []string{"DELETE /api/v1/x/dm/{}/messages/{}", "GET /api/v1/x/tweets/search", "GET /api/v1/x/tweets/{}"}
	if !slices.Equal(keys, want) {
		t.Fatalf("keys = %q, want %q", keys, want)
	}
	if found["GET /api/v1/x/tweets/search"].Position.Line != 10 {
		t.Fatalf("search reads line %d, want the first call", found["GET /api/v1/x/tweets/search"].Position.Line)
	}
}

func TestSDKOperationsNameUnreadableRequests(t *testing.T) {
	source := `package sdk

func (r *AService) A(ctx context.Context) error {
	return requestconfig.ExecuteNewRequest(ctx)
}

func (r *BService) B(ctx context.Context, method string) error {
	return requestconfig.ExecuteNewRequest(ctx, method, "x")
}

func (r *CService) C(ctx context.Context) error {
	path := "x"
	path = "y"
	return requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path)
}

func (r *DService) D(ctx context.Context, id string) error {
	return requestconfig.ExecuteNewRequest(ctx, http.MethodGet, fmt.Sprintf("x/%s/%s", id))
}

func (r *EService) E(ctx context.Context, id string) error {
	return requestconfig.ExecuteNewRequest(ctx, http.MethodGet, strings.Join(nil, id))
}

func (r *FService) F(ctx context.Context, id string) error {
	return requestconfig.ExecuteNewRequest(ctx, http.MethodGet, fmt.Sprintf(format, id))
}

func (r *GService) G(ctx context.Context, id string) error {
	return requestconfig.ExecuteNewRequest(ctx, http.MethodGet, build())
}

func (r *HService) H(ctx context.Context, id string) error {
	return requestconfig.ExecuteNewRequest(ctx, http.MethodGet, 42)
}

func (r *IService) I(ctx context.Context, id string) error {
	return requestconfig.ExecuteNewRequest(ctx, http.Method, "x")
}
`
	_, unreadable, err := Coverage.SDKOperations(fstest.MapFS{"a.go": {Data: []byte(source)}}, "")
	if err != nil {
		t.Fatal(err)
	}
	cannot := " sends a request whose method or path the coverage check cannot read. Use an http.Method constant and a path string or fmt.Sprintf format."
	want := []string{
		"a.go:4 calls ExecuteNewRequest without a method and a path. Pass both as arguments.",
		"a.go:8" + cannot, "a.go:14" + cannot, "a.go:18" + cannot, "a.go:22" + cannot,
		"a.go:26" + cannot, "a.go:30" + cannot, "a.go:34" + cannot, "a.go:38" + cannot,
	}
	if !slices.Equal(unreadable, want) {
		t.Fatalf("unreadable = %q\nwant %q", unreadable, want)
	}
}

func TestParseGoFilesFailsWithNextSteps(t *testing.T) {
	tests := map[string]struct {
		root fstest.MapFS
		dir  string
		want string
	}{
		"missing dir":  {fstest.MapFS{"a.go": {Data: []byte("package sdk\n")}}, "shared", "The field check could not list the SDK files. Run it from the repository root."},
		"syntax error": {fstest.MapFS{"shared/a.go": {Data: []byte("package shared\n\nfunc {")}}, "shared", "The field check could not parse shared/a.go. Run go vet ./... and fix the syntax error."},
		"unreadable":   {fstest.MapFS{"shared/a.go": {Data: []byte("package shared\n"), Mode: os.ModeSymlink}}, "shared", "The field check could not read shared/a.go. Run it from the repository root."},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			if _, err := Fields.ParseGoFiles(test.root, test.dir, token.NewFileSet()); err == nil || err.Error() != test.want {
				t.Fatalf("err = %v, want %s", err, test.want)
			}
		})
	}
	if _, _, err := Coverage.SDKOperations(fstest.MapFS{}, ""); err != nil {
		t.Fatalf("an empty tree reads no operations, err = %v", err)
	}
	broken := fstest.MapFS{"a.go": {Data: []byte("package sdk\n\nfunc {")}}
	if _, _, err := Coverage.SDKOperations(broken, ""); err == nil || err.Error() != "The coverage check could not parse a.go. Run go vet ./... and fix the syntax error." {
		t.Fatalf("err = %v", err)
	}
}

func TestGroupFileValidatesGroups(t *testing.T) {
	file := GroupFile{Name: "groups.json", ItemsKey: "operations", Item: "operation", ItemShape: "an operation", Parse: func(text string) (string, bool) {
		return strings.ToLower(text), strings.HasPrefix(text, "GET ")
	}}
	tests := map[string]string{
		"[}":                                  "It is not valid JSON near byte 2",
		`{}`:                                  "It must be a JSON list of groups, each with only a reason string and a list of strings named operations",
		`[] []`:                               "It holds more than 1 JSON value",
		`[{"reason": "x", "other": []}]`:      "It must be a JSON list of groups, each with only a reason string and a list of strings named operations",
		`[{"reason": 1, "operations": []}]`:   "It must be a JSON list of groups, each with only a reason string and a list of strings named operations",
		`[{"reason": " ", "operations": []}]`: "Group 1 needs a non-empty reason",
		`[{"reason": "x", "operations": []}]`: "Group 1 needs at least 1 operation",
		`[{"reason": "x", "operations": ["POST /x"]}]`:          `Group 1 lists "POST /x", which is not an operation`,
		`[{"reason": "x", "operations": ["GET /x", "GET /X"]}]`: "It lists GET /x and GET /X, which name the same operation. List it once",
	}
	for content, want := range tests {
		t.Run(content, func(t *testing.T) {
			_, err := file.Read(fstest.MapFS{"groups.json": {Data: []byte(content)}})
			if err == nil || err.Error() != "groups.json is not valid. "+want+"." {
				t.Fatalf("err = %v, want %s", err, want)
			}
		})
	}
	if _, err := file.Read(fstest.MapFS{}); err == nil || err.Error() != "groups.json is not valid. The repository root has no readable copy. Create it, even as an empty list []." {
		t.Fatalf("err = %v", err)
	}
	found, err := (GroupFile{Name: "groups.json", ItemsKey: "differences", Item: "difference", ItemShape: "a difference", Parse: file.Parse}).Read(fstest.MapFS{"groups.json": {Data: []byte(`[{"reason": "x", "differences": ["GET /A"]}]`)}})
	if err != nil || found["get /a"] != "GET /A" {
		t.Fatalf("found = %v, err = %v", found, err)
	}
}
