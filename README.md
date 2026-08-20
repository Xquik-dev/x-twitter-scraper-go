# Xquik Go SDK: Twitter Search, Followers & X Automation

[![OpenSSF Best Practices](https://www.bestpractices.dev/projects/13734/badge)](https://www.bestpractices.dev/projects/13734)
[![Ask DeepWiki](https://deepwiki.com/badge.svg?url=https%3A%2F%2Fgithub.com%2FXquik-dev%2Fx-twitter-scraper-go)](https://deepwiki.com/Xquik-dev/x-twitter-scraper-go)
[![Skills.sh x-twitter-scraper Skill](https://skills.sh/b/xquik-dev/x-twitter-scraper)](https://skills.sh/xquik-dev/x-twitter-scraper)

<!-- x-release-please-start-version -->

<a href="https://pkg.go.dev/github.com/Xquik-dev/x-twitter-scraper-go"><img src="https://pkg.go.dev/badge/github.com/Xquik-dev/x-twitter-scraper-go.svg" alt="Go Reference"></a>

<!-- x-release-please-end -->

Use the Xquik Go SDK for Twitter search, profiles, followers & media.
Manage webhooks, bulk extractions & X automation with typed Go methods.
The module calls the documented Xquik REST API.

## Twitter API Alternative

This package does not call or emulate the official X API.
Use it when the official X API does not fit your workflow.

[Go Reference](https://pkg.go.dev/github.com/Xquik-dev/x-twitter-scraper-go) | [REST API Docs](https://docs.xquik.com/api-reference/overview) | [OpenAPI Spec](https://xquik.com/openapi.json) | [Context7](https://context7.com/xquik-dev/x-twitter-scraper-go) | [Webhooks](https://docs.xquik.com/api-reference/webhooks/create) | [OAuth-First MCP Guide](https://docs.xquik.com/mcp/overview)

## Common Twitter & X Tasks

| Task | REST Route | Workflow Note |
| --- | --- | --- |
| Search tweets without the X API | `GET /x/tweets/search` | Use keyword or advanced operator queries. |
| Read an X profile timeline | `GET /x/users/{id}/tweets` | Paginate bounded results. |
| Scrape Twitter followers | `GET /x/users/{id}/followers` | Use an extraction for complete datasets. |
| Scrape following accounts | `GET /x/users/{id}/following` | Use an extraction for complete datasets. |
| Read a home timeline | `GET /x/timeline` | Approve this private read. |
| Export large X datasets | `POST /extractions` | Poll status, then download results. |
| Download or upload media | `/x/media/*` | Use typed file helpers. |
| Monitor an account | `POST /monitors` | Deliver events through HMAC webhooks. |
| Post or reply | `POST /x/tweets` | Confirm the account and payload. |

## Installation

<!-- x-release-please-start-version -->

```go
import (
	"github.com/Xquik-dev/x-twitter-scraper-go" // imported as xtwitterscraper
)
```

<!-- x-release-please-end -->

Or to pin the version:

<!-- x-release-please-start-version -->

```sh
go get -u 'github.com/Xquik-dev/x-twitter-scraper-go@v0.18.1'
```

<!-- x-release-please-end -->

## Verify a Release

Verify the matching source archive before upgrading:

```sh
release_tag=vVERSION
archive="x-twitter-scraper-go-$release_tag.zip"

gh release download "$release_tag" \
  --repo Xquik-dev/x-twitter-scraper-go \
  --pattern "$archive"

gh attestation verify "$archive" \
  --repo Xquik-dev/x-twitter-scraper-go \
  --signer-workflow Xquik-dev/x-twitter-scraper-go/.github/workflows/release-provenance.yml \
  --source-ref "refs/tags/$release_tag" \
  --deny-self-hosted-runners
```

Require the Xquik-dev repository and release workflow.
GitHub verifies the digest, signer identity, and transparency proof.

## Requirements

Use Go 1.26.6 or newer.

## Usage

See [api.md](api.md) for the complete API.

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Xquik-dev/x-twitter-scraper-go"
	"github.com/Xquik-dev/x-twitter-scraper-go/option"
)

func main() {
	client := xtwitterscraper.NewClient(
		option.WithAPIKey(os.Getenv("X_TWITTER_SCRAPER_API_KEY")),
	)
	response, err := client.X.Tweets.Search(context.TODO(), xtwitterscraper.XTweetSearchParams{
		Q:     "from:elonmusk",
		Limit: xtwitterscraper.Int(10),
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", response)
}

```

### Request Fields

Request structs follow Go's [`omitzero`](https://tip.golang.org/doc/go1.24#encodingjsonpkgencodingjson) semantics.
Fields tagged <code>\`api:"required"\`</code> serialize even when their value is zero.
Optional primitives use `param.Opt[T]` and constructors such as `xtwitterscraper.String`.
Maps, slices, structs & enums tagged <code>\`json:"...,omitzero"\`</code> omit zero values.
Use `param.IsOmitted(any)` to test any `omitzero` field.

```go
p := xtwitterscraper.ExampleParams{
	ID:   "id_xxx",                      // Required.
	Name: xtwitterscraper.String("..."), // Optional.

	Point: xtwitterscraper.Point{
		X: 0,                      // Serializes as 0.
		Y: xtwitterscraper.Int(1), // Serializes as 1.
	},

	Origin: xtwitterscraper.Origin{}, // Omitted because it is zero.
}
```

To send `null` instead of a `param.Opt[T]`, use `param.Null[T]()`.
To send `null` instead of a struct `T`, use `param.NullStruct[T]()`.

```go
p.Name = param.Null[string]()       // 'null' instead of string
p.Point = param.NullStruct[Point]() // 'null' instead of struct

param.IsNull(p.Name)  // true
param.IsNull(p.Point) // true
```

Use `.SetExtraFields(map[string]any)` only with trusted data.
Extra fields can overwrite matching struct fields.

To send a custom value instead of a struct, use `param.Override[T](value)`.

```go
// Override the documented integer type.
p.SetExtraFields(map[string]any{
	"x": 0.01,
})

// Send a number instead of an object.
custom := param.Override[xtwitterscraper.FooParams](12)
```

### Request Unions

Each union variant uses an `Of`-prefixed field. Set only one field.
The nonzero field is serialized. Getter methods return mutable pointers when present.

```go
// Set only one variant.
type AnimalUnionParam struct {
	OfCat *Cat `json:",omitzero,inline`
	OfDog *Dog `json:",omitzero,inline`
}

animal := AnimalUnionParam{
	OfCat: &Cat{
		Name: "Whiskers",
		Owner: PersonParam{
			Address: AddressParam{Street: "3333 Coyote Hill Rd", Zip: 0},
		},
	},
}

// Mutate a present field.
if address := animal.GetOwner().GetAddress(); address != nil {
	address.ZipCode = 94304
}
```

### Response Objects

Response fields use value types. The `JSON` field records decoding metadata.

```go
type Animal struct {
	Name   string `json:"name,nullable"`
	Owners int    `json:"owners"`
	Age    int    `json:"age"`
	JSON   struct {
		Name        respjson.Field
		Owner       respjson.Field
		Age         respjson.Field
		ExtraFields map[string]respjson.Field
	} `json:"-"`
}
```

`.Valid()` reports whether a field was present, non-null & decoded successfully.
Invalid or omitted fields keep their zero value.

```go
raw := `{"owners": 1, "name": null}`

var res Animal
json.Unmarshal([]byte(raw), &res)

// Values.
res.Owners // 1
res.Name   // ""
res.Age    // 0

// Presence checks.
res.JSON.Owners.Valid() // true
res.JSON.Name.Valid()   // false
res.JSON.Age.Valid()    // false

// Raw JSON.
res.JSON.Owners.Raw()                  // "1"
res.JSON.Name.Raw() == "null"          // true
res.JSON.Name.Raw() == respjson.Null   // true
res.JSON.Age.Raw() == ""               // true
res.JSON.Age.Raw() == respjson.Omitted // true
```

`JSON.ExtraFields` retains response properties absent from the generated struct.

```go
body := res.JSON.ExtraFields["my_unexpected_field"].Raw()
```

### Response Unions

Response unions flatten fields from every object variant.
Use `.AsFooVariant()` or `.AsAny()` to select a variant.
Primitive variants use `Of`-prefixed fields tagged `json:"...,inline"`.

```go
type AnimalUnion struct {
	// From variants [Dog], [Cat]
	Owner Person `json:"owner"`
	// From variant [Dog]
	DogBreed string `json:"dog_breed"`
	// From variant [Cat]
	CatBreed string `json:"cat_breed"`
	// ...

	JSON struct {
		Owner respjson.Field
		// ...
	} `json:"-"`
}

// Validate a shared field.
if animal.Owner.Address.ZipCode == "" {
	panic("missing zip code")
}

// Select the variant.
switch variant := animal.AsAny().(type) {
case Dog:
case Cat:
default:
	panic("unexpected type")
}
```

### RequestOptions

The `option` package returns `RequestOption` closures that update `RequestConfig`.
Apply them to the client or an individual request.

```go
client := xtwitterscraper.NewClient(
	// Add a header to every request.
	option.WithHeader("X-Some-Header", "custom_header_info"),
)

client.X.Tweets.Search(context.TODO(), ...,
	// Override the client header.
	option.WithHeader("X-Some-Header", "some_other_custom_header_info"),
	// Add a request field with sjson syntax.
	option.WithJSONSet("some.json.path", map[string]string{"my": "object"}),
)
```

`option.WithDebugLog(nil)` writes request and response content to the default logger.
Use it only with safe local data.

See the [full list of request options](https://pkg.go.dev/github.com/Xquik-dev/x-twitter-scraper-go/option).

### Pagination

`.ListAutoPaging()` iterates across every page.
`.List()` fetches one page; call `.GetNextPage()` for the next.

### Errors

Non-2xx responses return `*xtwitterscraper.Error`.
It contains the status, request, response & decoded error body.
Use `errors.As` to inspect those details:

```go
_, err := client.X.Tweets.Search(context.TODO(), xtwitterscraper.XTweetSearchParams{
	Q:     "from:elonmusk",
	Limit: xtwitterscraper.Int(10),
})
if err != nil {
	var apierr *xtwitterscraper.Error
	if errors.As(err, &apierr) {
		fmt.Printf("Request failed with HTTP %d\n", apierr.StatusCode)
	}
	panic(err.Error()) // GET "/x/tweets/search": 400 Bad Request { ... }
}
```

Other errors remain unwrapped. Transport failures may return `*url.Error` wrapping `*net.OpError`.

### Timeouts

Requests have no default timeout. Use a context for the full request lifecycle.
Retries share that context deadline. Use `option.WithRequestTimeout()` for each attempt.

```go
// Set one deadline across all attempts.
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
defer cancel()
client.X.Tweets.Search(
	ctx,
	xtwitterscraper.XTweetSearchParams{
		Q:     "from:elonmusk",
		Limit: xtwitterscraper.Int(10),
	},
	// Limit each attempt.
	option.WithRequestTimeout(20*time.Second),
)
```

### File Uploads

Multipart file parameters accept `io.Reader`.
The default filename is `anonymous_file`; the default media type is `application/octet-stream`.
Implement `Name() string` or `ContentType() string` to override either value.
`os.File` already supplies its disk filename through `Name()`.
Use `xtwitterscraper.File` to wrap any reader with explicit metadata.

```go
// Read from disk.
file, err := os.Open("/path/to/file")
xtwitterscraper.XMediaUploadParams{
	Account: "@elonmusk",
	File:    file,
}

// Read from a string.
xtwitterscraper.XMediaUploadParams{
	Account: "@elonmusk",
	File:    strings.NewReader("my file contents"),
}

// Set a filename and media type.
xtwitterscraper.XMediaUploadParams{
	Account: "@elonmusk",
	File:    xtwitterscraper.File(strings.NewReader(`{"hello": "foo"}`), "file.go", "application/json"),
}
```

### Retries

The SDK retries connection errors & HTTP 408, 409, 429, and 5xx responses.
It uses exponential backoff with 2 retries by default.
Use `WithMaxRetries` to change or disable retries:

```go
// Set the client default.
client := xtwitterscraper.NewClient(
	option.WithMaxRetries(0), // default is 2
)

// Override one request.
client.X.Tweets.Search(
	context.TODO(),
	xtwitterscraper.XTweetSearchParams{
		Q:     "from:elonmusk",
		Limit: xtwitterscraper.Int(10),
	},
	option.WithMaxRetries(5),
)
```

### Raw Response Data

Use `option.WithResponseInto()` to inspect response headers and status codes.

```go
var response *http.Response
paginatedTweets, err := client.X.Tweets.Search(
	context.TODO(),
	xtwitterscraper.XTweetSearchParams{
		Q:     "from:elonmusk",
		Limit: xtwitterscraper.Int(10),
	},
	option.WithResponseInto(&response),
)
if err != nil {
	panic(err)
}
fmt.Printf("%+v\n", paginatedTweets)

fmt.Printf("Status Code: %d\n", response.StatusCode)
fmt.Printf("Headers: %+#v\n", response.Header)
```

### Custom Requests

Use `client.Get`, `client.Post`, or another verb for undocumented endpoints.
These methods retain client options such as retries.

```go
var (
	// Accepts io.Reader, []byte, JSON-compatible values, or SDK params.
	params map[string]any

	// Accepts []byte, *http.Response, JSON-compatible values, or SDK models.
	result *http.Response
)
err := client.Post(context.Background(), "/unspecified", params, &result)
if err != nil {
	panic(err)
}
```

Use `option.WithQuerySet()` or `option.WithJSONSet()` for undocumented parameters.

```go
params := FooNewParams{
	ID: "id_xxxx",
	Data: FooNewParamsData{
		FirstName: xtwitterscraper.String("John"),
	},
}
client.Foo.New(context.Background(), params, option.WithJSONSet("data.last_name", "Doe"))
```

Use `result.JSON.RawJSON()` for the full body and `.Foo.Raw()` for one field.
`result.JSON.ExtraFields()` returns fields absent from the generated response struct.

### Middleware

Use `option.WithMiddleware` to wrap requests.

```go
func Logger(req *http.Request, next option.MiddlewareNext) (res *http.Response, err error) {
	// Before the request.
	start := time.Now()
	LogReq(req)

	// Run the next handler.
	res, err = next(req)

	// After the request.
	LogRes(res, err, time.Since(start))

	return res, err
}

client := xtwitterscraper.NewClient(
	option.WithMiddleware(Logger),
)
```

Middleware runs left to right. Client middleware runs before request middleware.
Use `option.WithHTTPClient(client)` to replace the current HTTP client.
The replacement receives requests after middleware runs.

## Semantic Versioning

This package follows [SemVer](https://semver.org/spec/v2.0.0.html).
Before v1.0, minor releases may change undocumented internals.
Open an [issue](https://github.com/Xquik-dev/x-twitter-scraper-go/issues) before depending on them.

## Contributing

Read the [contribution guide](CONTRIBUTING.md).

## Security & Project Health

- [Security policy](SECURITY.md)
- [Governance](GOVERNANCE.md)
- [OpenSSF evidence](OPENSSF.md)
- [Changelog](CHANGELOG.md)

Xquik is an independent third-party service. Not affiliated with X Corp. "Twitter" and "X" are trademarks of X Corp.
