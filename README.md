# X API & Twitter API Go SDK by Xquik

[![OpenSSF Best Practices](https://www.bestpractices.dev/projects/13734/badge)](https://www.bestpractices.dev/projects/13734)
[![Ask DeepWiki](https://deepwiki.com/badge.svg?url=https%3A%2F%2Fgithub.com%2FXquik-dev%2Fx-twitter-scraper-go)](https://deepwiki.com/Xquik-dev/x-twitter-scraper-go)

<a href="https://pkg.go.dev/github.com/Xquik-dev/x-twitter-scraper-go"><img src="https://pkg.go.dev/badge/github.com/Xquik-dev/x-twitter-scraper-go.svg" alt="Go Reference"></a>

Use this X API & Twitter API SDK to work with X from Go.
It works as a Twitter scraper & tweet scraper for search, profiles & followers.
Scrape Twitter timelines, lists, communities & trends into typed Go structs.
Manage media, webhooks, monitors & X automation with typed Go methods.
It provides a Twitter API alternative through documented Xquik REST routes.

Xquik is the world's fastest and cheapest X Twitter scraper service with the most complete X data.
Xquik charges only for delivered, unique, filter-matching results and never pads rows.

> Xquik is an independent third-party service. Not affiliated with X Corp.
> "Twitter" and "X" are trademarks of X Corp.

<!-- trust:start -->

## Trust & usage

[![Twitter API package downloads](https://img.shields.io/endpoint?url=https%3A%2F%2Fxquik.com%2Fbadges%2Fdownloads.json)](https://xquik.com/trust#downloads)
[![Twitter scraper users on Apify](https://img.shields.io/endpoint?url=https%3A%2F%2Fxquik.com%2Fbadges%2Fapify-users.json)](https://apify.com/xquik)
[![Twitter scraper runs on Apify](https://img.shields.io/endpoint?url=https%3A%2F%2Fxquik.com%2Fbadges%2Fapify-runs.json)](https://apify.com/xquik)
[![Tweet scraper rating on Apify](https://img.shields.io/endpoint?url=https%3A%2F%2Fxquik.com%2Fbadges%2Fapify-rating.json)](https://apify.com/xquik/x-tweet-scraper/reviews)
[![Xquik Twitter API accounts](https://img.shields.io/endpoint?url=https%3A%2F%2Fxquik.com%2Fbadges%2Faccounts.json)](https://xquik.com/trust#accounts)
[![Twitter Skill installs](https://img.shields.io/endpoint?url=https%3A%2F%2Fxquik.com%2Fbadges%2Fskill-installs.json)](https://xquik.com/trust#skill-installs)
[![Xquik verified revenue on TrustMRR](https://img.shields.io/endpoint?url=https%3A%2F%2Fxquik.com%2Fbadges%2Frevenue.json)](https://trustmrr.com/startup/xquik)
[![Twitter API median response time](https://img.shields.io/endpoint?url=https%3A%2F%2Fxquik.com%2Fbadges%2Frest-p50.json)](https://xquik.com/trust#speed)
[![Xquik GitHub stars](https://img.shields.io/endpoint?url=https%3A%2F%2Fxquik.com%2Fbadges%2Fgithub-stars.json)](https://github.com/Xquik-dev)
[![OpenSSF Best Practices](https://www.bestpractices.dev/projects/13734/badge)](https://www.bestpractices.dev/projects/13734)

- 51,000+ downloads of Xquik's Twitter API packages on npm, PyPI, RubyGems, NuGet, Packagist & Terraform ([source](https://xquik.com/trust#downloads)) <!-- trust:downloads 51020 -->
- 7,800+ users of Xquik's Twitter scraper Actors on Apify ([source](https://apify.com/xquik)) <!-- trust:apify-users 7800 -->
- 5,700,000+ runs of Xquik's Twitter scraper Actors on Apify ([source](https://apify.com/xquik)) <!-- trust:apify-runs 5704701 -->
- 4.66 of 5 for the X Tweet Scraper from 16 Apify reviews ([source](https://apify.com/xquik/x-tweet-scraper/reviews)) <!-- trust:apify-rating 4.66 16 -->
- 3,900+ accounts on xquik.com ([source](https://xquik.com/trust#accounts)) <!-- trust:accounts 3924 -->
- 1,600+ installs of Xquik's Twitter Skills on skills.sh ([source](https://xquik.com/trust#skill-installs)) <!-- trust:skill-installs 1612 -->
- $38,000+ in revenue verified by TrustMRR ([source](https://trustmrr.com/startup/xquik)) <!-- trust:revenue 38115 -->
- 501 ms median REST API read time over the last 7 days ([source](https://xquik.com/trust#speed)) <!-- trust:rest-p50 501 -->
- 370+ GitHub stars across Xquik's repos ([source](https://github.com/Xquik-dev)) <!-- trust:github-stars 374 -->
- This repo meets the OpenSSF Best Practices passing criteria ([source](https://www.bestpractices.dev/projects/13734)) <!-- trust:openssf-x-twitter-scraper-go 1 -->

Checked on 2026-10-15. See each number with its date & source on [xquik.com/trust](https://xquik.com/trust) or as [JSON](https://xquik.com/trust.json).

<!-- trust:end -->

## Twitter API alternative

This package does not call or emulate the official X API.
Use it when the official X API does not fit your workflow.

[Go reference](https://pkg.go.dev/github.com/Xquik-dev/x-twitter-scraper-go) | [REST API docs](https://docs.xquik.com/api-reference/overview) | [OpenAPI spec](https://xquik.com/openapi.json) | [Context7](https://context7.com/xquik-dev/x-twitter-scraper-go) | [Webhooks](https://docs.xquik.com/api-reference/webhooks/create) | [OAuth-first MCP guide](https://docs.xquik.com/mcp/overview)

## Common Twitter & X tasks

| Task | REST route | Workflow note |
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

```go
import (
	"github.com/Xquik-dev/x-twitter-scraper-go" // imported as xtwitterscraper
)
```

Or pin the version.

```sh
go get -u 'github.com/Xquik-dev/x-twitter-scraper-go@v0.22.0'
```

## Verify a release

Go checks every module download against the public checksum database.
Run this command to see the commit behind a release.

```sh
go mod download -json github.com/Xquik-dev/x-twitter-scraper-go@vVERSION
```

`Origin.Hash` names the tagged commit.

Releases through v0.18.1 also attach a source archive signed by the former release workflow.
Verify one of those archives before upgrading.

```sh
release_tag=v0.18.1
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

GitHub verifies the digest, signer identity, and transparency proof.

## Requirements

Use Go 1.27.3 or newer.

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

### Request fields

Request structs follow Go's [`omitzero`](https://tip.golang.org/doc/go1.24#encodingjsonpkgencodingjson) semantics.
Fields tagged <code>\`api:"required"\`</code> serialize even when their value is zero.
Optional primitives use `param.Opt[T]` and constructors such as `xtwitterscraper.String`.
Maps, slices, structs & enums tagged <code>\`json:"...,omitzero"\`</code> omit zero values.
Use `param.IsOmitted(field)` to test any `omitzero` field.

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

Use `.SetExtraFields(map[string]json.RawMessage)` only with trusted data.
Extra fields can overwrite matching struct fields.

To send a custom value instead of a struct, use `param.Override[T](value)`.

```go
// Override the documented integer type.
p.SetExtraFields(map[string]json.RawMessage{
	"x": json.RawMessage(`0.01`),
})

// Send a number instead of an object.
custom := param.Override[xtwitterscraper.FooParams](json.RawMessage(`12`))
```

### Request unions

Each union variant uses an `Of`-prefixed field. Set only one field.
The SDK serializes the nonzero field. Getter methods return mutable pointers when the field is present.

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

### Response objects

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

`JSON.ExtraFields` retains response properties absent from the response struct.

```go
body := res.JSON.ExtraFields["my_unexpected_field"].Raw()
```

### Response unions

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
variant, err := animal.AsAny()
if err != nil {
	panic(err)
}
switch variant := variant.(type) {
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

Each list method returns 1 page.
Pass the cursor from the last response to read the next page.

```go
params := xtwitterscraper.XAccountListParams{Limit: xtwitterscraper.Int(100)}
for {
	page, err := client.X.Accounts.List(context.TODO(), params)
	if err != nil {
		panic(err)
	}
	for _, account := range page.Accounts {
		fmt.Println(account.XUsername)
	}
	if !page.HasMore || page.NextCursor == "" {
		break
	}
	params.Cursor = xtwitterscraper.String(page.NextCursor)
}
```

Tweet & user lists report `HasNextPage` & `NextCursor` instead of `HasMore`.

### Server-sent events

`Events.StreamStreaming` streams monitor events as they happen.
Call `Next` until it returns false, then check `Err`.
Close the stream when you finish.

```go
params := xtwitterscraper.EventStreamParams{}
if lastID != "" {
	params.LastEventID = xtwitterscraper.String(lastID)
}
stream := client.Events.StreamStreaming(context.TODO(), params)
defer stream.Close()
for stream.Next() {
	event := stream.Current()
	lastID = event.ID
	fmt.Println(event.Type, event.ID)
}
if err := stream.Err(); err != nil {
	panic(err.Error())
}
```

The API ends each connection after 30 minutes at most.
Then start a new stream with `LastEventID` set to the last event ID you read.
Leave `LastEventID` unset on the first connection.
The stream skips the keep-alive comments the API sends while idle.
A client-level `option.WithRequestTimeout` also ends the stream.
Pass `option.WithRequestTimeout(0)` to the stream call to lift it.

### Errors

Non-2xx responses return `*xtwitterscraper.Error`.
It contains the status, request, response & decoded error body.
Use `errors.As` to inspect those details.

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
Retries share that context deadline. Use `option.WithRequestTimeout()` to limit each attempt.

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

### File uploads

Multipart file parameters accept `io.Reader`.
Media, avatar and banner uploads read the file type from the first bytes.
They know PNG, JPEG, GIF, WebP, AVIF and MP4.
An unnamed upload gets a name like `upload.png`, or `upload` for other bytes.
Other upload bytes go out as `application/octet-stream`.
Ticket attachments default to `anonymous_file` and `application/octet-stream`.
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
	File:    xtwitterscraper.File(strings.NewReader("my video bytes"), "clip.mp4", "video/mp4"),
}
```

### Retries

The SDK retries connection errors & HTTP 408, 409, 429, and 5xx responses.
It uses exponential backoff with 2 retries by default.
Use `WithMaxRetries` to change or disable retries.

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

### Raw response data

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

### Custom requests

Use `client.Get`, `client.Post`, or another verb for undocumented endpoints.
These methods retain client options such as retries.
They take SDK params or nil, and decode the response into a pointer to an SDK model or `*json.RawMessage`.
`option.WithRequestBody` sends an `io.Reader`, such as `bytes.NewReader(data)`, as the body instead.

```go
var result json.RawMessage
err := client.Get(context.Background(), "/unspecified", nil, &result)
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
`result.JSON.ExtraFields()` returns fields absent from the response struct.

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

## Semantic versioning

This package follows [SemVer](https://semver.org/spec/v2.0.0.html).
Before v1.0, minor releases may change undocumented internals.
Open an [issue](https://github.com/Xquik-dev/x-twitter-scraper-go/issues) before depending on them.

## Contributing

Read the [contribution guide](CONTRIBUTING.md).

## More Xquik tools

- [X (Twitter) scraper API & Twitter MCP server](https://github.com/Xquik-dev/x-twitter-scraper) with REST, MCP & Skill setup.
- Twitter API SDKs for [TypeScript](https://github.com/Xquik-dev/x-twitter-scraper-typescript), [Python](https://github.com/Xquik-dev/x-twitter-scraper-python), [Ruby](https://github.com/Xquik-dev/x-twitter-scraper-ruby), [Java](https://github.com/Xquik-dev/x-twitter-scraper-java), [Kotlin](https://github.com/Xquik-dev/x-twitter-scraper-kotlin), [C#](https://github.com/Xquik-dev/x-twitter-scraper-csharp) & [PHP](https://github.com/Xquik-dev/x-twitter-scraper-php).
- [Twitter CLI & X CLI](https://github.com/Xquik-dev/x-twitter-scraper-cli) for terminal workflows.
- [Twitter advanced search operators](https://github.com/Xquik-dev/x-twitter-advanced-search) with tested examples.
- [Xquik X API & Twitter scraper](https://xquik.com) and the [Twitter API documentation](https://docs.xquik.com/api-reference/overview).

## Security & project health

- [Security policy](SECURITY.md)
- [Governance](GOVERNANCE.md)
- [OpenSSF evidence](OPENSSF.md)
- [Changelog](CHANGELOG.md)

Xquik is an independent third-party service. Not affiliated with X Corp. "Twitter" and "X" are trademarks of X Corp.
