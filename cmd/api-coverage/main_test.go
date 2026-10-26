// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"testing/fstest"

	"github.com/Xquik-dev/x-twitter-scraper-go/internal/apicheck"
	"github.com/Xquik-dev/x-twitter-scraper-go/internal/requestconfig"
	"github.com/Xquik-dev/x-twitter-scraper-go/option"
)

const tweetSource = `package sdk

func (r *XTweetService) Get(ctx context.Context, id string) (res *Tweet, err error) {
	path := fmt.Sprintf("x/tweets/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res)
	return res, err
}

func (r *XTweetService) Search(ctx context.Context) (res *Page, err error) {
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, "x/tweets/search", nil, &res)
	return res, err
}

func (r XDmService) DeleteMessage(ctx context.Context, messageID string, params Params) (res *Action, err error) {
	path := fmt.Sprintf("x/dm/%s/messages/%s", url.PathEscape(params.UserID), url.PathEscape(messageID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, params, &res)
	return res, err
}

func (r *Client) Execute(ctx context.Context, method string, path string) error {
	return requestconfig.ExecuteNewRequest(ctx, method, path, nil, nil)
}

func (r *Store[T]) Load() {}

func helper() {}
`

const coveredSpec = `{
  "servers": [{"url": "https://xquik.com"}],
  "paths": {
    "/api/v1/x/tweets/{tweetId}": {"get": {"operationId": "getTweet"}},
    "/api/v1/x/tweets/search": {"get": {"operationId": "searchTweets"}},
    "/api/v1/x/dm/{userId}/messages/{messageId}": {"delete": {"operationId": "deleteDm"}, "parameters": []},
    "/ask": {"servers": [{"url": "https://xquik.com/"}], "get": {"operationId": "askDocs"}, "post": {"servers": [{"url": "https://docs.xquik.com"}]}}
  }
}`

const docsExclusion = `[{"reason": "Docs search lives outside the SDK base URL.", "operations": ["GET /ask", "POST /ask"]}]`

var productionBaseURL = requestconfig.WithDefaultBaseURL("https://xquik.com/api/v1/")

func sdkFiles(exclusions string) fstest.MapFS {
	return fstest.MapFS{
		"tweet.go":         {Data: []byte(tweetSource)},
		"tweet_test.go":    {Data: []byte("package sdk\n\nfunc (r *XTestService) Get() { requestconfig.ExecuteNewRequest(ctx, method, path) }\n")},
		"README.md":        {Data: []byte("# SDK\n")},
		"internal/skip.go": {Data: []byte("package internal\n")},
		exclusionFile:      {Data: []byte(exclusions)},
	}
}

func writeSpec(t *testing.T, content string) string {
	t.Helper()
	file := filepath.Join(t.TempDir(), "openapi.json")
	if err := os.WriteFile(file, []byte(content), 0o600); err != nil {
		t.Fatal(err)
	}
	return file
}

func runWith(t *testing.T, root fs.FS, baseURL option.RequestOption, source string) (int, string, string) {
	t.Helper()
	var out, failures bytes.Buffer
	code := run(source, root, baseURL, http.DefaultClient, &out, &failures)
	return code, out.String(), failures.String()
}

func runCheck(t *testing.T, root fs.FS, source string) (int, string, string) {
	t.Helper()
	return runWith(t, root, productionBaseURL, source)
}

func TestCoveredSpecPasses(t *testing.T) {
	for _, spec := range []string{
		coveredSpec,
		strings.Replace(coveredSpec, `"servers": [{"url": "https://xquik.com"}],`, "", 1),
	} {
		code, out, failures := runCheck(t, sdkFiles(docsExclusion), writeSpec(t, spec))
		want := "The SDK covers 3 of 5 API operations, and excluded-operations.json lists the other 2.\n"
		if code != 0 || out != want || failures != "" {
			t.Fatalf("code = %d, out = %q, failures = %q", code, out, failures)
		}
	}
}

func TestServerPathPrefixesSpecPaths(t *testing.T) {
	spec := `{"servers": [{"url": "https://xquik.com/api/v1/"}], "paths": {"/x/tweets/search": {"get": {}}}}`
	code, out, failures := runCheck(t, sdkFiles("[]"), writeSpec(t, spec))
	if code != 0 || out != "The SDK covers 1 of 1 API operations, and excluded-operations.json lists the other 0.\n" || failures != "" {
		t.Fatalf("code = %d, out = %q, failures = %q", code, out, failures)
	}
}

func TestMissingOperationFails(t *testing.T) {
	spec := strings.Replace(coveredSpec, `"paths": {`, `"paths": {
    "/api/v1/x/trends/history": {"get": {"operationId": "getXTrendsHistory"}},
    "/api/v1/x/trends": {"get": {}},`, 1)
	code, out, failures := runCheck(t, sdkFiles(docsExclusion), writeSpec(t, spec))
	want := "GET /api/v1/x/trends is in the API but not in the SDK.\n" +
		"GET /api/v1/x/trends/history (getXTrendsHistory) is in the API but not in the SDK.\n" +
		"Add each missing operation above to the SDK, or list it in excluded-operations.json with a reason.\n"
	if code != 1 || out != "" || failures != want {
		t.Fatalf("code = %d, out = %q, failures = %q", code, out, failures)
	}
}

func TestStaleExclusionsFail(t *testing.T) {
	exclusions := `[
  {"reason": "Docs search lives outside the SDK base URL.", "operations": ["GET /ask", "POST /ask"]},
  {"reason": "Not ported yet.", "operations": ["GET /api/v1/old", "GET /api/v1/x/tweets/{id}"]}
]`
	code, out, failures := runCheck(t, sdkFiles(exclusions), writeSpec(t, coveredSpec))
	want := "excluded-operations.json lists GET /api/v1/old, which the API no longer has. Remove it.\n" +
		"excluded-operations.json lists GET /api/v1/x/tweets/{id}, which the SDK already implements. Remove it.\n"
	if code != 1 || out != "" || failures != want {
		t.Fatalf("code = %d, out = %q, failures = %q", code, out, failures)
	}
}

func TestMissingAndStaleFailTogether(t *testing.T) {
	exclusions := `[{"reason": "Docs search lives outside the SDK base URL.", "operations": ["GET /ask", "POST /ask", "GET /api/v1/old"]}]`
	spec := strings.Replace(coveredSpec, `"paths": {`, `"paths": {"/api/v1/x/trends": {"get": {}},`, 1)
	code, out, failures := runCheck(t, sdkFiles(exclusions), writeSpec(t, spec))
	want := "GET /api/v1/x/trends is in the API but not in the SDK.\n" +
		"Add each missing operation above to the SDK, or list it in excluded-operations.json with a reason.\n" +
		"excluded-operations.json lists GET /api/v1/old, which the API no longer has. Remove it.\n"
	if code != 1 || out != "" || failures != want {
		t.Fatalf("code = %d, out = %q, failures = %q", code, out, failures)
	}
}

func TestPathParameterNamesDoNotMatter(t *testing.T) {
	spec := `{"servers": [{"url": "https://xquik.com"}], "paths": {
    "/api/v1/x/tweets/{id}": {"get": {}},
    "/api/v1/x/tweets/search": {"get": {}},
    "/api/v1/x/dm/{a}/messages/{b}": {"delete": {}},
    "/api/v1/x/users/{userId}": {"get": {}}
  }}`
	exclusions := `[{"reason": "Not ported yet.", "operations": ["GET /api/v1/x/users/{id}"]}]`
	code, out, failures := runCheck(t, sdkFiles(exclusions), writeSpec(t, spec))
	want := "The SDK covers 3 of 4 API operations, and excluded-operations.json lists the other 1.\n"
	if code != 0 || out != want || failures != "" {
		t.Fatalf("code = %d, out = %q, failures = %q", code, out, failures)
	}
}

func TestInvalidExclusionFilesFail(t *testing.T) {
	shape := "It must be a JSON list of groups, each with only a reason string and a list of strings named operations."
	notOperation := ", which is not an upper case HTTP method, a space and a path starting with /."
	for _, test := range []struct{ name, content, want string }{
		{"not JSON", "[}", "It is not valid JSON near byte 2."},
		{"cut short", "[{", shape},
		{"not a list", `{"reason": "x"}`, shape},
		{"unknown field", `[{"reason": "x", "operations": ["GET /x"], "note": "y"}]`, shape},
		{"two values", `[] []`, "It holds more than 1 JSON value."},
		{"trailing bracket", `[] ]`, "It holds more than 1 JSON value."},
		{"missing reason", `[{"operations": ["GET /x"]}]`, "Group 1 needs a non-empty reason."},
		{"blank reason", `[{"reason": " ", "operations": ["GET /x"]}]`, "Group 1 needs a non-empty reason."},
		{"missing operations", `[{"reason": "x"}]`, "Group 1 needs at least 1 operation."},
		{"empty operations", `[{"reason": "x", "operations": ["GET /x"]}, {"reason": "y", "operations": []}]`, "Group 2 needs at least 1 operation."},
		{"lower case method", `[{"reason": "x", "operations": ["get /x"]}]`, `Group 1 lists "get /x"` + notOperation},
		{"unknown method", `[{"reason": "x", "operations": ["FETCH /x"]}]`, `Group 1 lists "FETCH /x"` + notOperation},
		{"relative path", `[{"reason": "x", "operations": ["GET x"]}]`, `Group 1 lists "GET x"` + notOperation},
		{"duplicate", `[{"reason": "x", "operations": ["GET /x/{id}"]}, {"reason": "y", "operations": ["GET /x/{alias}"]}]`, "It lists GET /x/{id} and GET /x/{alias}, which name the same operation. List it once."},
	} {
		t.Run(test.name, func(t *testing.T) {
			code, out, failures := runCheck(t, sdkFiles(test.content), writeSpec(t, coveredSpec))
			if code != 1 || out != "" || failures != "excluded-operations.json is not valid. "+test.want+"\n" {
				t.Fatalf("code = %d, out = %q, failures = %q", code, out, failures)
			}
		})
	}
}

func TestMissingExclusionFileFails(t *testing.T) {
	files := sdkFiles("[]")
	delete(files, exclusionFile)
	code, _, failures := runCheck(t, files, writeSpec(t, coveredSpec))
	if code != 1 || failures != "excluded-operations.json is not valid. The repository root has no readable copy. Create it, even as an empty list [].\n" {
		t.Fatalf("code = %d, failures = %q", code, failures)
	}
}

func TestUnreadableSpecFails(t *testing.T) {
	failing := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer failing.Close()
	for _, source := range []string{
		filepath.Join(t.TempDir(), "missing.json"),
		writeSpec(t, "openapi: 3.1.0\n"),
		writeSpec(t, `{"paths": {}}`),
		writeSpec(t, `{"paths": {"/x": []}}`),
		writeSpec(t, `{"paths": {"/x": {"get": {"servers": [{"url": "https://[::1"}]}}}}`),
		failing.URL,
		"http://127.0.0.1:0/openapi.json",
		"https://exa mple.com/openapi.json",
	} {
		t.Run(source, func(t *testing.T) {
			code, out, failures := runCheck(t, sdkFiles(docsExclusion), source)
			want := "Could not read the API spec from " + source + ". Check network access, or set XQUIK_OPENAPI_SPEC to a local copy of https://xquik.com/openapi.json.\n"
			if code != 1 || out != "" || failures != want {
				t.Fatalf("code = %d, out = %q, failures = %q", code, out, failures)
			}
		})
	}
}

func TestSpecFromURL(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		if _, err := w.Write([]byte(coveredSpec)); err != nil {
			t.Error(err)
		}
	}))
	defer server.Close()
	code, out, failures := runCheck(t, sdkFiles(docsExclusion), server.URL)
	if code != 0 || out != "The SDK covers 3 of 5 API operations, and excluded-operations.json lists the other 2.\n" || failures != "" {
		t.Fatalf("code = %d, out = %q, failures = %q", code, out, failures)
	}
}

type roundTripper func(*http.Request) (*http.Response, error)

func (f roundTripper) RoundTrip(request *http.Request) (*http.Response, error) { return f(request) }

func TestEmptySourceUsesPublicSpec(t *testing.T) {
	var requested string
	client := &http.Client{Transport: roundTripper(func(request *http.Request) (*http.Response, error) {
		requested = request.URL.String()
		return nil, errors.New("offline")
	})}
	var out, failures bytes.Buffer
	code := run("", sdkFiles(docsExclusion), productionBaseURL, client, &out, &failures)
	if code != 1 || requested != "https://xquik.com/openapi.json" || !strings.HasPrefix(failures.String(), "Could not read the API spec from https://xquik.com/openapi.json.") {
		t.Fatalf("code = %d, requested = %q, failures = %q", code, requested, failures.String())
	}
}

func TestUnreadableRequestsFail(t *testing.T) {
	unreadable := " sends a request whose method or path the coverage check cannot read. Use an http.Method constant and a path string or fmt.Sprintf format.\n"
	for _, test := range []struct{ name, body, line string }{
		{"variable method", `method := http.MethodGet
	requestconfig.ExecuteNewRequest(ctx, method, "x", nil, nil)`, "tweet.go:5"},
		{"method name", `requestconfig.ExecuteNewRequest(ctx, "GET", "x", nil, nil)`, "tweet.go:4"},
		{"unknown method constant", `requestconfig.ExecuteNewRequest(ctx, http.MethodFetch, "x", nil, nil)`, "tweet.go:4"},
		{"other package constant", `requestconfig.ExecuteNewRequest(ctx, web.MethodGet, "x", nil, nil)`, "tweet.go:4"},
		{"path assigned twice", `path := "x"
	path = "y"
	requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, nil)`, "tweet.go:6"},
		{"path from a variable", `path, other := build()
	requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, other)`, "tweet.go:5"},
		{"path from another call", `requestconfig.ExecuteNewRequest(ctx, http.MethodGet, strings.Join(parts, "/"), nil, nil)`, "tweet.go:4"},
		{"path from an empty call", `requestconfig.ExecuteNewRequest(ctx, http.MethodGet, build(), nil, nil)`, "tweet.go:4"},
		{"format from a variable", `requestconfig.ExecuteNewRequest(ctx, http.MethodGet, fmt.Sprintf(format, id), nil, nil)`, "tweet.go:4"},
		{"format with too few values", `requestconfig.ExecuteNewRequest(ctx, http.MethodGet, fmt.Sprintf("x/%s/%s", id), nil, nil)`, "tweet.go:4"},
		{"format with an escaped percent", `requestconfig.ExecuteNewRequest(ctx, http.MethodGet, fmt.Sprintf("x/%s%%", id), nil, nil)`, "tweet.go:4"},
		{"path not a string", `requestconfig.ExecuteNewRequest(ctx, http.MethodGet, 42, nil, nil)`, "tweet.go:4"},
	} {
		t.Run(test.name, func(t *testing.T) {
			files := sdkFiles(docsExclusion)
			files["tweet.go"] = &fstest.MapFile{Data: []byte("package sdk\n\nfunc (r *XTweetService) Get(ctx context.Context) {\n\t" + test.body + "\n}\n")}
			code, out, failures := runCheck(t, files, writeSpec(t, coveredSpec))
			if code != 1 || out != "" || failures != test.line+unreadable {
				t.Fatalf("code = %d, out = %q, failures = %q", code, out, failures)
			}
		})
	}
}

func TestEveryUnreadableRequestIsReported(t *testing.T) {
	files := sdkFiles(docsExclusion)
	files["tweet.go"] = &fstest.MapFile{Data: []byte(`package sdk

func (r *XTweetService) Get(ctx context.Context) {
	requestconfig.ExecuteNewRequest(ctx)
	requestconfig.ExecuteNewRequest(ctx, method, "x", nil, nil)
}
`)}
	code, _, failures := runCheck(t, files, writeSpec(t, coveredSpec))
	want := "tweet.go:4 calls ExecuteNewRequest without a method and a path. Pass both as arguments.\n" +
		"tweet.go:5 sends a request whose method or path the coverage check cannot read. Use an http.Method constant and a path string or fmt.Sprintf format.\n"
	if code != 1 || failures != want {
		t.Fatalf("code = %d, failures = %q", code, failures)
	}
}

type brokenFS struct {
	files  fs.FS
	broken string
}

func (b brokenFS) Open(name string) (fs.File, error) {
	if name == b.broken {
		return nil, fs.ErrPermission
	}
	return b.files.Open(name)
}

func TestUnreadableSourcesFail(t *testing.T) {
	baseURLProblem := "The coverage check could not read the default base URL from option.WithEnvironmentProduction. Fix that option."
	for _, test := range []struct {
		name    string
		files   fs.FS
		baseURL option.RequestOption
		want    string
	}{
		{"broken base URL", sdkFiles(docsExclusion), requestconfig.WithDefaultBaseURL("://xquik.com"), baseURLProblem},
		{"no base URL", sdkFiles(docsExclusion), option.WithMaxRetries(0), baseURLProblem},
		{"broken service file", fstest.MapFS{"tweet.go": {Data: []byte("package sdk\nfunc (")}, exclusionFile: {Data: []byte("[]")}}, productionBaseURL,
			"The coverage check could not parse tweet.go. Run go vet ./... and fix the syntax error."},
		{"unreadable service file", brokenFS{files: sdkFiles(docsExclusion), broken: "tweet.go"}, productionBaseURL,
			"The coverage check could not read tweet.go. Run it from the repository root."},
		{"unlistable root", brokenFS{files: sdkFiles(docsExclusion), broken: "."}, productionBaseURL,
			"The coverage check could not list the SDK files. Run it from the repository root."},
	} {
		t.Run(test.name, func(t *testing.T) {
			code, out, failures := runWith(t, test.files, test.baseURL, writeSpec(t, coveredSpec))
			if code != 1 || out != "" || failures != test.want+"\n" {
				t.Fatalf("code = %d, out = %q, failures = %q", code, out, failures)
			}
		})
	}
}

func TestRepositorySourcesParse(t *testing.T) {
	operations, unreadable, err := apicheck.Coverage.SDKOperations(os.DirFS(filepath.Join("..", "..")), "/api/v1")
	if err != nil || len(unreadable) > 0 {
		t.Fatalf("err = %v, unreadable = %q", err, unreadable)
	}
	for _, want := range []string{
		"GET /api/v1/x/tweets/search",
		"POST /api/v1/x/tweets",
		"DELETE /api/v1/x/tweets/{}",
		"GET /api/v1/credits/topup/redirect",
	} {
		if _, ok := operations[want]; !ok {
			t.Errorf("the SDK sources lack %s", want)
		}
	}
}

func TestMainPassesOnTheRepository(t *testing.T) {
	root, err := filepath.Abs(filepath.Join("..", ".."))
	if err != nil {
		t.Fatal(err)
	}
	implemented, _, err := apicheck.Coverage.SDKOperations(os.DirFS(root), "/api/v1")
	if err != nil {
		t.Fatal(err)
	}
	excluded, err := exclusions.Read(os.DirFS(root))
	if err != nil {
		t.Fatal(err)
	}
	paths := map[string]map[string]json.RawMessage{}
	var keys []string
	for key := range implemented {
		keys = append(keys, key)
	}
	for key := range excluded {
		keys = append(keys, key)
	}
	for _, key := range keys {
		method, operationPath, _ := strings.Cut(key, " ")
		if paths[operationPath] == nil {
			paths[operationPath] = map[string]json.RawMessage{}
		}
		paths[operationPath][strings.ToLower(method)] = json.RawMessage("{}")
	}
	document, err := json.Marshal(struct {
		Paths map[string]map[string]json.RawMessage `json:"paths"`
	}{paths})
	if err != nil {
		t.Fatal(err)
	}
	t.Setenv(apicheck.SpecVariable, writeSpec(t, string(document)))
	file, err := os.CreateTemp(t.TempDir(), "stdout")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()
	stdout := os.Stdout
	os.Stdout = file
	t.Chdir(root)
	main()
	os.Stdout = stdout
	if _, err := file.Seek(0, io.SeekStart); err != nil {
		t.Fatal(err)
	}
	got, err := io.ReadAll(file)
	if err != nil {
		t.Fatal(err)
	}
	want := fmt.Sprintf("The SDK covers %d of %d API operations, and excluded-operations.json lists the other %d.\n",
		len(implemented), len(implemented)+len(excluded), len(excluded))
	if string(got) != want {
		t.Fatalf("main printed %q, want %q", got, want)
	}
}
