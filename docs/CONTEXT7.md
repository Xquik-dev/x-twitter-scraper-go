# Context7 Guide

Install, authenticate, and send your first Xquik request from Go.

## Install

```sh
go get github.com/Xquik-dev/x-twitter-scraper-go@latest
```

Use Go 1.26.6 or newer.

Import the SDK as `xtwitterscraper`:

```go
import (
	xtwitterscraper "github.com/Xquik-dev/x-twitter-scraper-go"
	"github.com/Xquik-dev/x-twitter-scraper-go/option"
)
```

## Authenticate

Set an API key in the process environment:

```sh
export X_TWITTER_SCRAPER_API_KEY="your-api-key"
```

Create a client:

```go
client := xtwitterscraper.NewClient(
	option.WithAPIKey(os.Getenv("X_TWITTER_SCRAPER_API_KEY")),
)
```

Never place API keys, webhook signing values, or user credentials in source
files, logs, examples, issues, or commits.

## First Request

Search recent tweets with X query operators:

```go
paginatedTweets, err := client.X.Tweets.Search(context.TODO(), xtwitterscraper.XTweetSearchParams{
	Q:     "from:elonmusk",
	Limit: xtwitterscraper.Int(10),
})
if err != nil {
	panic(err)
}

for _, tweet := range paginatedTweets.Tweets {
	fmt.Printf("@%s: %s\n", tweet.Author.Username, tweet.Text)
}
```

## Common Workflows

| Workflow | SDK entry point |
| --- | --- |
| Tweet search | `client.X.Tweets.Search` |
| Tweet lookup | `client.X.Tweets.Get` |
| User lookup | `client.X.Users.Get` |
| User search | `client.X.Users.GetSearch` |
| Follower export | `client.X.Users.GetFollowers` |
| Following export | `client.X.Users.GetFollowing` |
| Media upload | `client.X.Media.Upload` |
| Media download | `client.X.Media.Download` |
| Account monitoring | `client.Monitors.New` |
| Monitor events | `client.Events.List` |
| HMAC webhooks | `client.Webhooks.New` |
| Giveaway draws | `client.Draws.Run` |
| Bulk extractions | `client.Extractions.Run` |
| Regional trends | `client.Trends.List` |

Use `api.md` for methods and [pkg.go.dev](https://pkg.go.dev/github.com/Xquik-dev/x-twitter-scraper-go) for type signatures.

## Public Sources

[GitHub](https://github.com/Xquik-dev/x-twitter-scraper-go) | [Go Reference](https://pkg.go.dev/github.com/Xquik-dev/x-twitter-scraper-go) | [SDK Guide](https://docs.xquik.com/sdks/go) | [REST API](https://docs.xquik.com/api-reference/overview) | [OpenAPI](https://xquik.com/openapi.json) | [Context7](https://context7.com/xquik-dev/x-twitter-scraper-go) | [DeepWiki](https://deepwiki.com/Xquik-dev/x-twitter-scraper-go)

Xquik is an independent third-party service. Not affiliated with X Corp. "Twitter" and "X" are trademarks of X Corp.
