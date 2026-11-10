# X (Twitter) API Golang SDK & Scraper by Xquik

[![Framer builds a tweet ticker with Xquik](https://img.youtube.com/vi/4UOSpoOoC3Y/maxresdefault.jpg)](https://youtu.be/4UOSpoOoC3Y?t=367)\
[Watch Framer build a tweet ticker with Xquik, from 6:07](https://youtu.be/4UOSpoOoC3Y?t=367).
This Go SDK calls the Xquik API.

Use this Twitter scraper Go SDK by Xquik to read tweets into Go structs.
Every call takes a `context.Context`.
Public reads need no X (Twitter) developer account.

Xquik is the world's fastest and cheapest X (Twitter) scraper service with the most complete X (Twitter) data.
Most other Apify Actors charge before filtering or deduplication.
Xquik charges only for delivered, unique, filter-matching results and never pads rows.
Xquik is independent and not affiliated with X Corp.

<!-- trust:win:start -->

The Xquik REST API read 45 tweets per second on 2026-10-15. That cost $0.15 per 1,000. The fastest of 4 rival Actors read 12.5 per second. The cheapest cost $0.173 per 1,000. [Compare each run](https://xquik.com/benchmarks#benchmark-x-tweet-scraper).

<!-- trust:win:end -->

## Install the Go X (Twitter) API & scraper SDK

The SDK requires Go 1.27.3 or later.

```sh
go get 'github.com/Xquik-dev/x-twitter-scraper-go@v0.23.0'
```

## Scrape tweets with the Golang X (Twitter) API

Set `X_TWITTER_SCRAPER_API_KEY`, then run this search.
Search takes the same operators as X (Twitter) search & returns typed tweets.

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Xquik-dev/x-twitter-scraper-go" // imported as xtwitterscraper
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

Next, read the [Go SDK guide](https://docs.xquik.com/sdks/go), [Go reference](https://pkg.go.dev/github.com/Xquik-dev/x-twitter-scraper-go) or [REST API reference](https://docs.xquik.com/api-reference/overview).

<!-- trust:start -->

## Benchmarks & usage

[![X (Twitter) API package downloads](https://img.shields.io/endpoint?url=https%3A%2F%2Fxquik.com%2Fbadges%2Fdownloads.json)](https://xquik.com/benchmarks#downloads)
[![X (Twitter) API p99 server read time](https://img.shields.io/endpoint?url=https%3A%2F%2Fxquik.com%2Fbadges%2Frest-p99.json)](https://xquik.com/benchmarks#speed)
[![OpenSSF Best Practices](https://www.bestpractices.dev/projects/13734/badge)](https://www.bestpractices.dev/projects/13734)

- 63,000+ downloads of Xquik's X (Twitter) API packages across 6 registries ([source](https://xquik.com/benchmarks#downloads)) <!-- trust:downloads 63829 -->
- 5.8 s p99 REST server read time over 7 days, failures included ([source](https://xquik.com/benchmarks#speed)) <!-- trust:rest-p99 5793 -->
- This repo meets the OpenSSF Best Practices passing criteria ([source](https://www.bestpractices.dev/projects/13734)) <!-- trust:openssf-x-twitter-scraper-go 1 -->

Checked on 2026-11-07. See each number with its date & source on [xquik.com/benchmarks](https://xquik.com/benchmarks) or as [JSON](https://xquik.com/trust.json).

<!-- trust:end -->

## X (Twitter) API alternative

This package does not call or emulate the official X API.
Build on the [OpenAPI spec](https://xquik.com/openapi.json), [webhooks](https://docs.xquik.com/api-reference/webhooks/create) or [MCP guide](https://docs.xquik.com/mcp/overview).
AI coding tools can read this SDK through [Context7](https://context7.com/xquik-dev/x-twitter-scraper-go).

## Common Twitter & X tasks

| Task | REST route | Workflow note |
| --- | --- | --- |
| Search tweets without the X API | `GET /x/tweets/search` | Use keyword or advanced operator queries. |
| Read an X profile timeline | `GET /x/users/{id}/tweets` | Pass the cursor to read the next page. |
| Scrape Twitter followers | `GET /x/users/{id}/followers` | Use an extraction for complete datasets. |
| Scrape following accounts | `GET /x/users/{id}/following` | Use an extraction for complete datasets. |
| Read a home timeline | `GET /x/timeline` | Connect an X account first. |
| Export large X datasets | `POST /extractions` | Poll status, then download results. |
| Download or upload media | `/x/media/*` | Use typed file helpers. |
| Monitor an account | `POST /monitors` | Deliver events through HMAC webhooks. |
| Post or reply | `POST /x/tweets` | Confirm the account and payload. |

## Verify a release

Go checks every module download against the public checksum database.
Run this command to see the commit behind a release.

```sh
go mod download -json github.com/Xquik-dev/x-twitter-scraper-go@v0.23.0
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

## Usage

See [api.md](api.md) for the complete API.

### Request fields

Request structs follow Go's [`omitzero`](https://tip.golang.org/doc/go1.24#encodingjsonpkgencodingjson) semantics.
Fields tagged <code>\`api:"required"\`</code> serialize even when their value is zero.
Optional primitives use `param.Opt[T]` and constructors such as `xtwitterscraper.String`.
Maps, slices, structs & enums tagged <code>\`json:"...,omitzero"\`</code> omit zero values.
Use `param.IsOmitted(any)` to test any `omitzero` field.

```go
params := xtwitterscraper.XTweetSearchParams{
	Q:                "from:NASA",                 // Required.
	Limit:            xtwitterscraper.Int(10),     // Optional.
	BlueVerifiedOnly: xtwitterscraper.Bool(false), // Optional, and sends false.
}

fmt.Println(param.IsOmitted(params.AdvancedQuery)) // true
```

To send `null` instead of a `param.Opt[T]`, use `param.Null[T]()`.
To send `null` instead of a struct `T`, use `param.NullStruct[T]()`.

```go
params := xtwitterscraper.XTweetSearchParams{Q: "from:NASA"}
params.Limit = param.Null[int64]()

fmt.Println(param.IsNull(params.Limit)) // true
```

Use `option.WithQuerySet()` or `option.WithJSONSet()` to send a field the struct lacks.
To send a custom value instead of a struct, use `param.Override[T](value)`.

### Response objects

Response fields use value types. Each struct's `JSON` field records decoding metadata.
`.Valid()` reports whether a field was present, non-null & decoded without error.
Invalid or omitted fields keep their zero value.
`JSON.ExtraFields` retains response properties absent from the response struct.

```go
raw := `{"id": "1", "text": "Launch day.", "likeCount": 1200, "communityId": null, "source": "web"}`

var tweet xtwitterscraper.SearchTweet
if err := json.Unmarshal([]byte(raw), &tweet); err != nil {
	panic(err)
}

// Values.
fmt.Println(tweet.LikeCount)   // 1200
fmt.Println(tweet.CommunityID) // ""
fmt.Println(tweet.ViewCount)   // 0

// Presence checks.
fmt.Println(tweet.JSON.LikeCount.Valid())   // true
fmt.Println(tweet.JSON.CommunityID.Valid()) // false
fmt.Println(tweet.JSON.ViewCount.Valid())   // false

// Raw JSON.
fmt.Println(tweet.JSON.LikeCount.Raw())                       // 1200
fmt.Println(tweet.JSON.CommunityID.Raw() == respjson.Null)    // true
fmt.Println(tweet.JSON.ViewCount.Raw() == respjson.Omitted)   // true
fmt.Println(tweet.JSON.ExtraFields["source"].Raw())           // "web"
```

### Response unions

Response unions flatten fields from every object variant.
Use `.AsFooVariant()` or `.AsAny()` to select a variant.

```go
raw := `{"id": "xatt_1", "object": "x_account_connection_attempt", "status": "failed", "error": "Wrong password."}`

var attempt xtwitterscraper.XAccountConnectionAttemptGetResponseUnion
if err := json.Unmarshal([]byte(raw), &attempt); err != nil {
	panic(err)
}

// Read a shared field.
fmt.Println(attempt.Status) // failed

// Select the variant.
switch variant := attempt.AsAny().(type) {
case xtwitterscraper.XAccountConnectionAttemptGetResponseFailed:
	fmt.Println(variant.Error) // Wrong password.
case xtwitterscraper.XAccountConnectionAttemptGetResponseSuccess:
	fmt.Println(variant.ID)
default:
	fmt.Println("still pending")
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

response, err := client.X.Tweets.Search(
	context.TODO(),
	xtwitterscraper.XTweetSearchParams{Q: "from:NASA", Limit: xtwitterscraper.Int(10)},
	// Override the client header.
	option.WithHeader("X-Some-Header", "some_other_custom_header_info"),
	// Add a query field the params struct lacks.
	option.WithQuerySet("queryType", "Latest"),
)
if err != nil {
	panic(err)
}
fmt.Println(len(response.Tweets))
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
lastID := "" // The last event ID you saw, if any.

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

The default client waits up to 10 minutes for response headers.
It sets no other timeout. Use a context for the full request lifecycle.
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
download, err := http.Get("https://xquik.com/logo.png")
if err != nil {
	panic(err)
}
image, err := io.ReadAll(download.Body)
download.Body.Close()
if err != nil {
	panic(err)
}
if err := os.WriteFile("logo.png", image, 0o600); err != nil {
	panic(err)
}

// Read from disk.
file, err := os.Open("logo.png")
if err != nil {
	panic(err)
}
defer file.Close()

uploads := []xtwitterscraper.XMediaUploadParams{
	{Account: "@elonmusk", File: file, IdempotencyKey: "upload-disk"},
	// Read from memory.
	{Account: "@elonmusk", File: bytes.NewReader(image), IdempotencyKey: "upload-memory"},
	// Set a filename and media type.
	{Account: "@elonmusk", File: xtwitterscraper.File(bytes.NewReader(image), "logo.png", "image/png"), IdempotencyKey: "upload-named"},
}
for _, params := range uploads {
	if _, err := client.X.Media.Upload(context.TODO(), params); err != nil {
		panic(err)
	}
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

```go
var (
	// Accepts io.Reader, []byte, JSON-compatible values, or SDK params.
	params map[string]any

	// Accepts []byte, *http.Response, JSON-compatible values, or SDK models.
	result *http.Response
)
err := client.Get(context.Background(), "x/tweets/search?q=from:NASA&limit=1", params, &result)
if err != nil {
	panic(err)
}
fmt.Println(result.StatusCode)
```

Use `option.WithQuerySet()` or `option.WithJSONSet()` for undocumented parameters.

```go
response, err := client.X.Tweets.Search(
	context.Background(),
	xtwitterscraper.XTweetSearchParams{Q: "from:NASA", Limit: xtwitterscraper.Int(1)},
	option.WithQuerySet("queryType", "Top"),
)
if err != nil {
	panic(err)
}
fmt.Println(len(response.JSON.RawJSON()) > 0)
```

Use `result.JSON.RawJSON()` for the full body and `.Foo.Raw()` for one field.
`result.JSON.ExtraFields` returns fields absent from the response struct.

### Middleware

Use `option.WithMiddleware` to wrap requests.

```go
logger := func(req *http.Request, next option.MiddlewareNext) (*http.Response, error) {
	// Before the request.
	start := time.Now()

	// Run the next handler.
	res, err := next(req)

	// After the request.
	fmt.Println(req.Method, req.URL.Path, time.Since(start))

	return res, err
}

client := xtwitterscraper.NewClient(
	option.WithMiddleware(logger),
)

if _, err := client.X.Tweets.Search(context.TODO(), xtwitterscraper.XTweetSearchParams{Q: "from:NASA", Limit: xtwitterscraper.Int(10)}); err != nil {
	panic(err)
}
```

Middleware runs left to right. Client middleware runs before request middleware.
Use `option.WithHTTPClient(client)` to replace the current HTTP client.
The replacement receives requests after middleware runs.

## Semantic versioning

This package follows [SemVer](https://semver.org/spec/v2.0.0.html).
Before v1.0, minor releases may change undocumented internals.
Open an [issue](https://github.com/Xquik-dev/x-twitter-scraper-go/issues) before depending on them.

## Questions

### Do I need an X (Twitter) developer account?

No. Public reads need only an Xquik API key.
Private reads & X account actions need a connected X account.

### Is this the official X API?

No. Xquik runs its own REST API. This SDK calls that API, not the official X API.

## More Xquik tools

- [X (Twitter) scraper API & X (Twitter) MCP server](https://github.com/Xquik-dev/x-twitter-scraper) with REST, MCP & Skill setup.
- X (Twitter) API SDKs for [TypeScript](https://github.com/Xquik-dev/x-twitter-scraper-typescript), [Python](https://github.com/Xquik-dev/x-twitter-scraper-python), [Ruby](https://github.com/Xquik-dev/x-twitter-scraper-ruby) & [Java](https://github.com/Xquik-dev/x-twitter-scraper-java).
- More X (Twitter) API SDKs for [Kotlin](https://github.com/Xquik-dev/x-twitter-scraper-kotlin), [C#](https://github.com/Xquik-dev/x-twitter-scraper-csharp) & [PHP](https://github.com/Xquik-dev/x-twitter-scraper-php).
- [Twitter CLI & X CLI](https://github.com/Xquik-dev/x-twitter-scraper-cli) for terminal workflows.
- [Twitter advanced search operators](https://github.com/Xquik-dev/x-twitter-advanced-search) with tested examples.
- Twitter agent plugins for [OpenClaw](https://github.com/Xquik-dev/tweetclaw) & [Hermes Agent](https://github.com/Xquik-dev/hermes-tweet).
- Twitter integrations for [n8n](https://github.com/Xquik-dev/n8n-nodes-xquik), [Prefect](https://github.com/Xquik-dev/prefect-xquik), [Haystack](https://github.com/Xquik-dev/xquik-haystack) & [Paperclip](https://github.com/Xquik-dev/paperclip-plugin-xquik).
- More Twitter integrations for [Terraform](https://github.com/Xquik-dev/terraform-provider-x-twitter-scraper), [Convex](https://github.com/Xquik-dev/xquik-convex), [Supabase](https://github.com/Xquik-dev/xquik-supabase) & [Strapi](https://github.com/Xquik-dev/strapi-plugin-xquik).
- X (Twitter) MCP for [AI agents](https://github.com/Xquik-dev/xquik-mcp), [VS Code](https://github.com/Xquik-dev/xquik-vscode) & [Raycast](https://github.com/Xquik-dev/xquik-raycast).
- [Xquik X API & Twitter scraper](https://xquik.com), [X (Twitter) API docs](https://docs.xquik.com/api-reference/overview) & [docs source](https://github.com/Xquik-dev/xquik-docs).

## Help & contributing

- Report vulnerabilities through the [security policy](SECURITY.md).
- Ask usage questions through the [organization support policy](https://github.com/Xquik-dev/.github/blob/main/SUPPORT.md).
- Read the [contribution guide](CONTRIBUTING.md) & [governance policy](GOVERNANCE.md).
- Read the [OpenSSF evidence](OPENSSF.md), [changelog](CHANGELOG.md) & [citation metadata](CITATION.cff).

Xquik is an independent third-party service. Not affiliated with X Corp. "Twitter" and "X" are trademarks of X Corp.
