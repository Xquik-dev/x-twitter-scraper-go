# Changelog

## 0.21.0 (2026-10-18)

[Full changelog from v0.20.0 to v0.21.0](https://github.com/Xquik-dev/x-twitter-scraper-go/compare/v0.20.0...v0.21.0)

The SDK catches up with the API. Compatibility lists the 5 changes that break callers.

- Support tickets and replies take up to 4 `Attachments`. `Body` is optional with attachments.
- `X.Media.Upload`, `X.Profile.UpdateAvatar` and `X.Profile.UpdateBanner` take a `File`, sent as multipart. `X.Media.Upload` also takes `IsLongVideo`. The SDK reads a file's type from its first bytes, unless the reader names one.
- Requests without files still send JSON.
- `X.Communities.GetAnalytics` reads a Community's activity for X's latest period.
- Bulk retry can target specific accounts with `AccountIDs`.
- Embedded parent, quoted and retweeted posts carry every tweet field and the author.
- Posts and embedded posts name the counts X did not send in `UnknownCounts`. Each of those counts reads 0, so treat it as unknown.
- `shared.UserProfileIdentityVerification` gains the optional `OverrideVerifiedYear` field for profiles and users in tweets.
- The API fields check compares JSON, multipart and form request bodies.
- `bin/check-local`, `bin/release-local` and their scripts remove their temp files on every exit, including a failed step, Ctrl+C and a terminate signal. `scripts/verify-release` no longer prints the path of its deleted module cache.

### Compatibility

- `X.Accounts.BulkRetry` takes a required `XAccountBulkRetryParams` argument, so `BulkRetry(ctx)` stops compiling. Pass `x.XAccountBulkRetryParams{}` to retry every failed account as before.
- `Body` on `SupportTicketNewParams` and `SupportTicketReplyParams` now uses `param.Opt[string]`. Wrap body strings with `xtwitterscraper.String`.
- Support tickets and replies need a non-empty `Body` or at least 1 attachment. Otherwise the SDK returns `ErrTicketNeedsBodyOrAttachments` before any request.
- `URL` on `XMediaUploadParams`, `XProfileUpdateAvatarParams` and `XProfileUpdateBannerParams` now uses `param.Opt[string]`. Wrap URL strings with `xtwitterscraper.String`.
- An upload with both `URL` and `File` returns `ErrMediaHasURLAndFile` before any request. An upload with neither returns `ErrMediaNeedsURLOrFile`. `IsLongVideo` without a file returns `ErrLongVideoNeedsFile`.

## 0.20.0 (2026-10-16)

[Full changelog from v0.19.2 to v0.20.0](https://github.com/Xquik-dev/x-twitter-scraper-go/compare/v0.19.2...v0.20.0)

The SDK catches up with the API. It adds the 55 operations it lacked and the request and response fields it dropped. It now covers 180 of the 186 API operations, and `excluded-operations.json` says why the other 6 stay out. Every change is additive except the 1 that Compatibility lists.

- The new `X.Analysis` service has `Brand`, `Classify`, `MarketSignals`, `News`, `Sentiment` and `ViralScore`.
- `Events.StreamStreaming` streams monitor events as server-sent events. The new `packages/ssestream` package reads them and skips keep-alive comments.
- `Webhooks.NewTelegramLink` returns a link that connects a Telegram chat to monitor events.
- `constant.Succeeded` joins the shared constants.
- `X` gains 4 services. `X.Jobs` has `Get`, `ListLocations` and `Search`. `X.Spaces` has `Get`, `GetReplay` and `Search`. `X.Broadcasts` has `Get`. `X.CommunityNotes` has `GetContributor`, `ListContributorNotes` and `ListContributorDeletedNotes`.
- `shared.UserProfile` gains the 24 fields the API returns and the SDK dropped, such as `ListedCount`, `Email`, `ContactInfo` and `VerificationReason`.
- `shared.SearchTweet` and `TweetDetail` gain the 17 tweet fields the SDK dropped, such as `IsPinned`, `IsRetweet`, `ParentTweet`, `Interstitial` and `Geo`. Their entities, conversation control, note tweets, sports context and tombstones gain their missing fields too.
- `shared.PaginatedTweets` gains `ReadBackTo`, `SearchContext`, `SearchedQuery` and `WindowComplete`, and `shared.TweetMedia` gains `GrokImageAnnotation` and `SourceUser`.
- The params of 28 read methods gain `Fields` and `Omit`. `X.GetArticle`, `X.Communities.GetInfo`, `X.Tweets.Get` and `X.Users.Get` keep their signatures and take both through `option.WithQuery`.
- Enum types gain the values the API added, such as `delete_dm` for write actions, 2 monitor pause reasons and the `likeCount` estimate source.
- `X.Communities` gains `Find`, `ListPopular`, `ListSuggested`, `ListTopics` and `SearchMembers`, and the new `X.Communities.Media` service lists a Community's media posts.
- `X.Lists` gains `Get` and `GetSearch`, and `X.Dm` gains `DeleteMessage`.
- `X` gains `AutocompleteSearch`, `GetHashflags`, `GetNewsStoryVersions`, `GetTrendLocations`, `GetTrendsHistory` and `SearchPlaces`.
- `X.GetTrends` takes `At`, `Country`, `Location`, `Fields` and `Omit`, and its response gains `CapturedAt`, `RequestedAt` and `Location`. `At` reads the trends of a past time.
- The shared package gains `Community`, `CommunityPage` and `CommunityTopic`.
- `X.Users` gains `GetAffiliates`, `GetArticles`, `GetBatchTweets`, `GetFollowerIDs`, `GetFollowingIDs`, `GetHighlights`, `GetInsights`, `GetListMemberships`, `GetLists`, `GetLive`, `GetSimilar`, `GetSubscriptions` and `TranslateBio`.
- `X.Tweets` gains `CheckRepost`, `Embed`, `GetHiddenReplies`, `GetInsights`, `GetRetweeterIDs`, `GetSubtitles`, `ListTop`, `ResolveLinks` and `Translate`.
- `bin/check-local` starts with an API coverage check. It fails when the API has an operation that the SDK neither implements nor lists in `excluded-operations.json`. CONTRIBUTING.md explains the check and the file.
- `bin/check-local` runs an API fields check right after the coverage check. It compares the query params, request bodies and JSON responses with the spec at every depth. `accepted-field-differences.json` lists the 73 differences the SDK keeps on purpose, each with its reason. CONTRIBUTING.md explains the check and the file.
- The params of 29 methods gain the query params and body fields the API added. They include the email filters `HasEmail`, `BusinessEmailOnly` and `EmailDomains`, `BlueVerifiedOnly`, `MaxScan` and `EnrichProfiles` on profile reads, the search operators on Community and List tweets, and `Name` on monitors.
- Responses gain the fields the API returns and the SDK dropped. They include the Community creator profile, banner and members facepile, the scan limit of filtered follower and member reads, article links and counts, reply diagnostics, the started job from `Extractions.Run` and the email code challenge from `X.AccountConnectionChallenges.Submit`. New fields reuse the shared models where the spec does.
- `shared.EventType` gains `tweet.deleted`, `broadcast.started`, `space.started`, the 4 follower and following events and 2 profile count events.
- The `api:"required"` tags of 14 response fields now match the spec.

### Compatibility

- `XAccount` and `XAccountDetail` drop a timestamp field the API stopped returning. It always read empty.

## 0.19.2 (2026-10-14)

[Full changelog from v0.19.1 to v0.19.2](https://github.com/Xquik-dev/x-twitter-scraper-go/compare/v0.19.1...v0.19.2)

The API is the same as in 0.19.1.

- A program on Windows sends `Windows` in the `X-Stainless-OS` header. It sent `Other:windows` before.
- The README links the other Xquik repositories by name.
- The source has no code comments, and CodeQL fails on a new one.

## 0.19.1 (2026-10-14)

[Full changelog from v0.19.0 to v0.19.1](https://github.com/Xquik-dev/x-twitter-scraper-go/compare/v0.19.0...v0.19.1)

The API and its behavior are the same as in 0.19.0.

- The README uses sentence case headings, active voice and whole sentences. Its facts, links and legal lines stay the same.

## 0.19.0 (2026-10-13)

[Full changelog from v0.18.1 to v0.19.0](https://github.com/Xquik-dev/x-twitter-scraper-go/compare/v0.18.1...v0.19.0)

- **Breaking.** `X.Accounts.List` returns 1 page, with 50 accounts by default. A caller with more accounts gets fewer rows and no error. Read every page with `HasMore` & `NextCursor`.
- **Breaking.** Code written for 0.18 can stop compiling. "Upgrading from 0.18" shows each change with old & new code.
- `MinLikes` is the new name of the like filter on 9 params types. `MinFaves` still works and is deprecated.
- Add optional retweet timestamp requests and nullable `retweetedAt` response fields.
- Redact sensitive headers regardless of capitalization without changing the input.
- Use Go 1.27.2 and native license checks.
- Preserve JSON behavior while replacing obsolete compatibility helpers.
- Doc comments use the current price. An active monitor costs 30 credits per hour.

### Upgrading from 0.18

`X.Accounts.List` returns 1 page, with 50 accounts by default.
A caller with more than 50 accounts gets fewer rows and no error.
Read every page with `HasMore` & `NextCursor`, as the first section shows.

Go's `apidiff` tool reports 70 incompatible changes from v0.18.1 to v0.19.0.
The sections below cover all 70, and each heading gives its count.
The Go 1.27.2 minimum comes on top, because `apidiff` compares only Go declarations.

#### `X.Accounts.List` takes params & returns 1 page (1 change)

The method has a new `XAccountListParams` argument.
A page holds 50 accounts by default, and `Limit` accepts 1 to 100.
The response has 2 new fields, `HasMore` & `NextCursor`.

This code compiled with 0.18.

```go
response, err := client.X.Accounts.List(ctx)
if err != nil {
	return err
}
accounts := response.Accounts
```

Write this with 0.19. It reads every page.

```go
params := xtwitterscraper.XAccountListParams{Limit: xtwitterscraper.Int(100)}
var accounts []xtwitterscraper.XAccount
for {
	page, err := client.X.Accounts.List(ctx, params)
	if err != nil {
		return err
	}
	accounts = append(accounts, page.Accounts...)
	if !page.HasMore || page.NextCursor == "" {
		break
	}
	params.Cursor = xtwitterscraper.String(page.NextCursor)
}
```

#### `X.Lists.GetFollowers` & `X.Lists.GetMembers` return union types (2 changes)

Each method returned `*shared.PaginatedUsers`.
`GetFollowers` now returns `*XListGetFollowersResponseUnion`, and `GetMembers` returns `*XListGetMembersResponseUnion`.
`HasNextPage`, `NextCursor` & `Users` keep their names, so only code that names the type changes.

This code compiled with 0.18.

```go
var followers *shared.PaginatedUsers
followers, err := client.X.Lists.GetFollowers(ctx, listID, xtwitterscraper.XListGetFollowersParams{})
if err != nil {
	return err
}
var members *shared.PaginatedUsers
members, err = client.X.Lists.GetMembers(ctx, listID, xtwitterscraper.XListGetMembersParams{})
if err != nil {
	return err
}
fmt.Println(len(followers.Users), len(members.Users))
```

Write this with 0.19.

```go
var followers *xtwitterscraper.XListGetFollowersResponseUnion
followers, err := client.X.Lists.GetFollowers(ctx, listID, xtwitterscraper.XListGetFollowersParams{})
if err != nil {
	return err
}
var members *xtwitterscraper.XListGetMembersResponseUnion
members, err = client.X.Lists.GetMembers(ctx, listID, xtwitterscraper.XListGetMembersParams{})
if err != nil {
	return err
}
fmt.Println(len(followers.Users), len(members.Users))

// AsPaginatedUsers returns the type that 0.18 returned.
var page shared.PaginatedUsers = followers.AsPaginatedUsers()
fmt.Println(page.HasNextPage)
```

#### `Entities` is a struct on 6 tweet types (6 changes)

`Entities` was `map[string]any`.
It is now a struct with `Hashtags`, `Smarttags`, `Symbols`, `Timestamps`, `URLs` & `UserMentions`.
The 6 types are `TweetDetail`, `TweetDetailNoteTweet`, `shared.EmbeddedTweet`, `shared.EmbeddedTweetNoteTweet`, `shared.SearchTweet` & `shared.SearchTweetNoteTweet`.

This code compiled with 0.18.

```go
page, err := client.X.Tweets.Search(ctx, xtwitterscraper.XTweetSearchParams{Q: "from:elonmusk"})
if err != nil {
	return err
}
for _, tweet := range page.Tweets {
	urls, _ := tweet.Entities["urls"].([]any)
	for _, item := range urls {
		link, _ := item.(map[string]any)
		fmt.Println(link["expanded_url"])
	}
}
```

Write this with 0.19.

```go
page, err := client.X.Tweets.Search(ctx, xtwitterscraper.XTweetSearchParams{Q: "from:elonmusk"})
if err != nil {
	return err
}
for _, tweet := range page.Tweets {
	for _, link := range tweet.Entities.URLs {
		fmt.Println(link.ExpandedURL)
	}
}
```

#### Quoted & retweeted tweets are structs on `shared.EmbeddedTweet` (2 changes)

`QuotedTweet` & `RetweetedTweet` were `*shared.EmbeddedTweet`.
They are now the structs `shared.EmbeddedTweetQuotedTweet` & `shared.EmbeddedTweetRetweetedTweet`.
A nil check no longer compiles, so check presence with the `JSON` metadata.

This code compiled with 0.18.

```go
embedded := tweet.QuotedTweet
if embedded.QuotedTweet != nil {
	fmt.Println(embedded.QuotedTweet.Text)
}
if embedded.RetweetedTweet != nil {
	fmt.Println(embedded.RetweetedTweet.Text)
}
```

Write this with 0.19.

```go
embedded := tweet.QuotedTweet
if embedded.JSON.QuotedTweet.Valid() {
	fmt.Println(embedded.QuotedTweet.Text)
}
if embedded.JSON.RetweetedTweet.Valid() {
	fmt.Println(embedded.RetweetedTweet.Text)
}
```

#### Extraction jobs, results & sources are typed (3 changes)

`ExtractionGetResponse.Job` was `map[string]any` and is now `ExtractionJob`.
`ExtractionGetResponse.Results` was `[]map[string]any` and is now `[]ExtractionGetResponseResult`.
`ExtractionRunResponse.Source` was `string` and is now `ExtractionRunResponseSource`.

This code compiled with 0.18.

```go
response, err := client.Extractions.Get(ctx, extractionID, xtwitterscraper.ExtractionGetParams{})
if err != nil {
	return err
}
status, _ := response.Job["status"].(string)
fmt.Println(status)
for _, row := range response.Results {
	fmt.Println(row["tweetText"])
}

estimate, err := client.Extractions.Run(ctx, params)
if err != nil {
	return err
}
var source string = estimate.Source
fmt.Println(source)
```

Write this with 0.19.

```go
response, err := client.Extractions.Get(ctx, extractionID, xtwitterscraper.ExtractionGetParams{})
if err != nil {
	return err
}
status := string(response.Job.Status)
fmt.Println(status)
for _, row := range response.Results {
	fmt.Println(row.TweetText)
}

estimate, err := client.Extractions.Run(ctx, params)
if err != nil {
	return err
}
var source string = string(estimate.Source)
fmt.Println(source)
```

#### Go 1.27.2 is the minimum (not among the 79)

The module required Go 1.26.6 in 0.18 and requires Go 1.27.2 in 0.19.
`go get` raises the `go` line in your `go.mod` to match.

Your `go.mod` held this line or an older one with 0.18.

```
go 1.26.6
```

It holds this line with 0.19.

```
go 1.27.2
```

Go downloads the 1.27.2 toolchain when yours is older.
It stops with an error instead when `GOTOOLCHAIN=local` is set.
The module also requires `golang.org/x/tools` v0.50.0, `golang.org/x/mod` v0.41.0 & `golang.org/x/sync` v0.23.0 for its tests & commands.
A build that uses an older version of 1 of them moves up to that version.

#### `MinFaves` is deprecated on 9 params types

`MinLikes` is the new name of the same filter, and it sends the `minLikes` query key.
`MinFaves` still compiles and still sends `minFaves`, which the API reads the same way.
The 9 types are `XTweetGetQuotesParams`, `XTweetGetRepliesParams`, `XTweetGetThreadParams`, `XTweetSearchParams`, `XUserGetLikesParams`, `XUserGetMediaParams`, `XUserGetMentionsParams`, `XUserGetRepliesParams` & `XUserGetTweetsParams`.
Code written for 0.18 keeps working. Move to `MinLikes` when you next edit it.

This code compiled with 0.18 and still compiles with 0.19.

```go
params := xtwitterscraper.XTweetSearchParams{
	Q:        "from:elonmusk",
	MinFaves: xtwitterscraper.Int(100),
}
```

Prefer this with 0.19.

```go
params := xtwitterscraper.XTweetSearchParams{
	Q:        "from:elonmusk",
	MinLikes: xtwitterscraper.Int(100),
}
```

#### 3 detail types embed their summary type (19 changes)

`DraftDetail` embeds `Draft`, `DrawDetail` embeds `DrawListItem`, and `StyleProfile` embeds `StyleProfileSummary`.
`apidiff` counts 16 moved fields & 3 changed `JSON` structs.
Field reads such as `draft.ID` and struct literals still compile.
A presence check on a moved field now goes through the embedded type.

- `DraftDetail` moved `ID`, `CreatedAt`, `Text`, `Goal` & `Topic`.
- `DrawDetail` moved `ID`, `CreatedAt`, `DrawnAt`, `Status`, `TotalEntries`, `TweetURL` & `ValidEntries`.
- `StyleProfile` moved `FetchedAt`, `IsOwnAccount`, `TweetCount` & `XUsername`.

This code compiled with 0.18.

```go
draft, err := client.Drafts.Get(ctx, draftID)
if err != nil {
	return err
}
if draft.JSON.Goal.Valid() {
	fmt.Println(draft.Goal)
}
```

Write this with 0.19.

```go
draft, err := client.Drafts.Get(ctx, draftID)
if err != nil {
	return err
}
if draft.Draft.JSON.Goal.Valid() {
	fmt.Println(draft.Goal)
}
```

`draw.JSON.Status` becomes `draw.DrawListItem.JSON.Status` the same way.
`profile.JSON.TweetCount` becomes `profile.StyleProfileSummary.JSON.TweetCount`.

#### 1 constant & 1 constant type are gone (3 changes)

`ExtractionEstimateCostResponseSourcePosts` is gone.
`constant.ProductionWeightNotPublishedByX` is gone.
`ComposeNewResponseComposePrepareResultEngagementMultiplier.Multiplier` now has the type `constant.SourceDefaultProductionValueCanVary`.

This code compiled with 0.18.

```go
if estimate.Source == xtwitterscraper.ExtractionEstimateCostResponseSourcePosts {
	fmt.Println("The estimate counts posts.")
}
var note constant.ProductionWeightNotPublishedByX = multiplier.Multiplier
fmt.Println(note)
```

Write this with 0.19.

```go
if estimate.Source == "posts" {
	fmt.Println("The estimate counts posts.")
}
var note constant.SourceDefaultProductionValueCanVary = multiplier.Multiplier
fmt.Println(note)
```

#### 34 `JSON` metadata structs gained fields (34 changes)

`apidiff` flags a type when its `JSON` struct gains a field.
No caller code changes for these 34.
Each `JSON` struct has an unexported field, so code outside the SDK cannot name its type.

- `ExtractionGetResponse`
- `ExtractionJob`
- `MonitorDeactivateResponse`
- `MonitorKeywordDeactivateResponse`
- `TweetDetail`
- `TweetDetailArticle`
- `TweetDetailCard`
- `TweetDetailCommunityNote`
- `TweetDetailEdit`
- `TweetDetailNoteTweet`
- `XAccountListResponse`
- `XBookmarkGetFoldersResponseFolder`
- `XTweetSearchResponseUnion`
- `XUserGetBatchResponse`
- `XUserGetFollowersResponseUnion`
- `XUserGetFollowingResponseUnion`
- `XUserGetVerifiedFollowersResponseUnion`
- `shared.EmbeddedTweet`
- `shared.EmbeddedTweetArticle`
- `shared.EmbeddedTweetCard`
- `shared.EmbeddedTweetCommunityNote`
- `shared.EmbeddedTweetEdit`
- `shared.EmbeddedTweetNoteTweet`
- `shared.PaginatedTweets`
- `shared.PaginatedUsers`
- `shared.SearchTweet`
- `shared.SearchTweetArticle`
- `shared.SearchTweetCard`
- `shared.SearchTweetCommunityNote`
- `shared.SearchTweetEdit`
- `shared.SearchTweetNoteTweet`
- `shared.TweetMedia`
- `shared.UserProfile`
- `shared.UserProfileAffiliatesHighlightedLabel`

## 0.18.1 (2026-08-21)

Full Changelog: [v0.18.0...v0.18.1](https://github.com/Xquik-dev/x-twitter-scraper-go/compare/v0.18.0...v0.18.1)

### Documentation

* tighten SDK guidance and examples

### Bug Fixes

* correct the middleware duration example

## 0.18.0 (2026-08-18)

Full Changelog: [v0.17.0...v0.18.0](https://github.com/Xquik-dev/x-twitter-scraper-go/compare/v0.17.0...v0.18.0)

### Features

* **api:** api update ([65b7e46](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/65b7e46a5529edeaae963564862575fed80d0ffa))
* **api:** api update ([f09db18](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/f09db18dbb10b37dd8f8d70cf21410b7f5508279))
* **api:** api update ([7222208](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/72222089240bb6903f794d2e0b0e7373116a271d))
* **api:** api update ([f8084e8](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/f8084e83e19e250a9b46a526c4439fbe57c94619))
* **api:** api update ([ae71ffd](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/ae71ffd41a4d4f11b449f4224ed901df3609ed8d))
* **api:** publish provider-neutral checkout contracts ([95db019](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/95db019c2824c35ce1c7d7c625cdeb5f0ea81565))
* enforce OpenSSF Gold quality gates ([696a2c1](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/696a2c1104239667f9182916a5e6474bf810532b))


### Bug Fixes

* make diagnostic dumps concurrency-safe ([6c35deb](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/6c35deb74fab127bcde9947c0a915e90dbad067b))
* make draft release lookup reliable ([577b8c9](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/577b8c9585ce146f61c8f2b20a99696dd28ab680))
* publish assets before immutable release ([acc5ac8](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/acc5ac866388710e5050b27e08c4ad54fb2fc9d1))


### Chores

* add FUNDING.yml pointing to xquik.com ([97cc8a9](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/97cc8a9cb1d7e268721debb5b177fdcdbaf2fad2))
* add per-file license evidence ([ab79e56](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/ab79e56db99acb4aa4069f48fcf469de17dca319))
* avoid failing Context7 refresh without secret ([7401976](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/740197678fc873102a3c8eb3cfdee13f1aebfac0))
* **deps:** bump actions/attest from 4.2.1 to 4.2.2 ([#39](https://github.com/Xquik-dev/x-twitter-scraper-go/issues/39)) ([51c218a](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/51c218a4944bb6c68a92cef70f684ceef4873e26))
* **deps:** bump actions/setup-go from 6.5.0 to 7.0.0 ([#21](https://github.com/Xquik-dev/x-twitter-scraper-go/issues/21)) ([f5bc310](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/f5bc310215ca309302332ffb8cd60b438c8e9d9c))
* **deps:** bump github/codeql-action/upload-sarif from 4.37.1 to 4.37.6 ([#32](https://github.com/Xquik-dev/x-twitter-scraper-go/issues/32)) ([5b4a515](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/5b4a5154c313ecc613fe5c4fa468625092bb5c0b))
* **deps:** bump ossf/scorecard-action from 2.4.3 to 2.4.4 ([#34](https://github.com/Xquik-dev/x-twitter-scraper-go/issues/34)) ([f52356b](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/f52356bad0ff705081d48e4191eb884afd081af0))
* prepare v0.5.2 ([e49dced](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/e49dced219d42d7ab36990012e459061ac9f51b9))
* refresh Go SDK dependencies ([9f148cc](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/9f148cc763746a8d68a886ed8a23763ac043ee8a))
* sync repo ([9d47120](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/9d471208569e4262e693f18990eb915e6accb67a))
* update Go toolchain ([d0ee365](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/d0ee365dc02f842f6fb5594f1743681d0b956638))


### Documentation

* add Contributor Covenant 2.1 Code of Conduct ([80c6f70](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/80c6f70b5cf1ace74f29e9e867671d9856bed365))
* add DeepWiki badge ([884f895](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/884f895bd2883295b94a71f0162dab7d13e43df2))
* clarify repository discovery ([26c6fa3](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/26c6fa3b735576ffb942812f0d7b0ddb590058a9))
* complete v0.12.0 release notes ([7d335f2](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/7d335f2552a800b804fe5095d28ff4555c07e95a))
* map common X data tasks ([39cfa7e](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/39cfa7efd4ee63bf301de922d90367682130fab7))

## 0.17.0 (2026-08-18)

Full Changelog: [v0.16.1...v0.17.0](https://github.com/Xquik-dev/x-twitter-scraper-go/compare/v0.16.1...v0.17.0)

### Features

* **api:** api update ([65b7e46](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/65b7e46a5529edeaae963564862575fed80d0ffa))
* **api:** api update ([f09db18](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/f09db18dbb10b37dd8f8d70cf21410b7f5508279))
* **api:** api update ([7222208](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/72222089240bb6903f794d2e0b0e7373116a271d))
* **api:** api update ([f8084e8](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/f8084e83e19e250a9b46a526c4439fbe57c94619))
* **api:** api update ([ae71ffd](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/ae71ffd41a4d4f11b449f4224ed901df3609ed8d))
* **api:** publish provider-neutral checkout contracts ([95db019](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/95db019c2824c35ce1c7d7c625cdeb5f0ea81565))
* enforce OpenSSF Gold quality gates ([696a2c1](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/696a2c1104239667f9182916a5e6474bf810532b))


### Bug Fixes

* make diagnostic dumps concurrency-safe ([6c35deb](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/6c35deb74fab127bcde9947c0a915e90dbad067b))
* make draft release lookup reliable ([577b8c9](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/577b8c9585ce146f61c8f2b20a99696dd28ab680))
* publish assets before immutable release ([acc5ac8](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/acc5ac866388710e5050b27e08c4ad54fb2fc9d1))


### Chores

* add FUNDING.yml pointing to xquik.com ([97cc8a9](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/97cc8a9cb1d7e268721debb5b177fdcdbaf2fad2))
* add per-file license evidence ([ab79e56](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/ab79e56db99acb4aa4069f48fcf469de17dca319))
* avoid failing Context7 refresh without secret ([7401976](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/740197678fc873102a3c8eb3cfdee13f1aebfac0))
* **deps:** bump actions/attest from 4.2.1 to 4.2.2 ([#39](https://github.com/Xquik-dev/x-twitter-scraper-go/issues/39)) ([51c218a](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/51c218a4944bb6c68a92cef70f684ceef4873e26))
* **deps:** bump actions/setup-go from 6.5.0 to 7.0.0 ([#21](https://github.com/Xquik-dev/x-twitter-scraper-go/issues/21)) ([f5bc310](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/f5bc310215ca309302332ffb8cd60b438c8e9d9c))
* **deps:** bump github/codeql-action/upload-sarif from 4.37.1 to 4.37.6 ([#32](https://github.com/Xquik-dev/x-twitter-scraper-go/issues/32)) ([5b4a515](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/5b4a5154c313ecc613fe5c4fa468625092bb5c0b))
* **deps:** bump ossf/scorecard-action from 2.4.3 to 2.4.4 ([#34](https://github.com/Xquik-dev/x-twitter-scraper-go/issues/34)) ([f52356b](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/f52356bad0ff705081d48e4191eb884afd081af0))
* prepare v0.5.2 ([e49dced](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/e49dced219d42d7ab36990012e459061ac9f51b9))
* refresh Go SDK dependencies ([9f148cc](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/9f148cc763746a8d68a886ed8a23763ac043ee8a))
* sync repo ([9d47120](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/9d471208569e4262e693f18990eb915e6accb67a))
* update Go toolchain ([d0ee365](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/d0ee365dc02f842f6fb5594f1743681d0b956638))


### Documentation

* add Contributor Covenant 2.1 Code of Conduct ([80c6f70](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/80c6f70b5cf1ace74f29e9e867671d9856bed365))
* add DeepWiki badge ([884f895](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/884f895bd2883295b94a71f0162dab7d13e43df2))
* clarify repository discovery ([26c6fa3](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/26c6fa3b735576ffb942812f0d7b0ddb590058a9))
* complete v0.12.0 release notes ([7d335f2](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/7d335f2552a800b804fe5095d28ff4555c07e95a))
* map common X data tasks ([39cfa7e](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/39cfa7efd4ee63bf301de922d90367682130fab7))

## 0.16.1 (2026-08-18)

Full Changelog: [v0.16.0...v0.16.1](https://github.com/Xquik-dev/x-twitter-scraper-go/compare/v0.16.0...v0.16.1)

### Chores

* update Go toolchain ([d0ee365](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/d0ee365dc02f842f6fb5594f1743681d0b956638))

## 0.16.0 (2026-08-12)

Full Changelog: [v0.15.1...v0.16.0](https://github.com/Xquik-dev/x-twitter-scraper-go/compare/v0.15.1...v0.16.0)

### Features

* **api:** sync public contracts ([57581ab](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/57581ab0c65ed9fdcc516d70a46165af67205122))

## 0.15.1 (2026-08-11)

Full Changelog: [v0.15.0...v0.15.1](https://github.com/Xquik-dev/x-twitter-scraper-go/compare/v0.15.0...v0.15.1)

### Bug Fixes

* **api:** clarify idempotent guest-wallet replays ([ae71ffd](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/ae71ffd41a4d4f11b449f4224ed901df3609ed8d))

## 0.15.0 (2026-08-11)

Full Changelog: [v0.14.0...v0.15.0](https://github.com/Xquik-dev/x-twitter-scraper-go/compare/v0.14.0...v0.15.0)

### Features

* **api:** publish provider-neutral checkout contracts ([95db019](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/95db019c2824c35ce1c7d7c625cdeb5f0ea81565))

## 0.14.0 (2026-08-11)

Full Changelog: [v0.13.2...v0.14.0](https://github.com/Xquik-dev/x-twitter-scraper-go/compare/v0.13.2...v0.14.0)

### Features

* **api:** api update ([e634040](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/e634040ab9a8d504a72cde1d33af279a21f641bb))

## 0.13.2 (2026-08-11)

Full Changelog: [v0.13.1...v0.13.2](https://github.com/Xquik-dev/x-twitter-scraper-go/compare/v0.13.1...v0.13.2)

### Chores

* **deps:** bump actions/attest from 4.2.1 to 4.2.2 ([#39](https://github.com/Xquik-dev/x-twitter-scraper-go/issues/39)) ([51c218a](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/51c218a4944bb6c68a92cef70f684ceef4873e26))

## 0.13.1 (2026-08-11)

Full Changelog: [v0.13.0...v0.13.1](https://github.com/Xquik-dev/x-twitter-scraper-go/compare/v0.13.0...v0.13.1)

### Chores

* **deps:** bump github/codeql-action/upload-sarif from 4.37.1 to 4.37.6 ([#32](https://github.com/Xquik-dev/x-twitter-scraper-go/issues/32)) ([5b4a515](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/5b4a5154c313ecc613fe5c4fa468625092bb5c0b))
* **deps:** bump ossf/scorecard-action from 2.4.3 to 2.4.4 ([#34](https://github.com/Xquik-dev/x-twitter-scraper-go/issues/34)) ([f52356b](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/f52356bad0ff705081d48e4191eb884afd081af0))

## 0.13.0 (2026-07-31)

Full Changelog: [v0.12.0...v0.13.0](https://github.com/Xquik-dev/x-twitter-scraper-go/compare/v0.12.0...v0.13.0)

### Features

* **api:** api update ([65b7e46](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/65b7e46a5529edeaae963564862575fed80d0ffa))
* **api:** api update ([f09db18](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/f09db18dbb10b37dd8f8d70cf21410b7f5508279))
* **api:** api update ([7222208](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/72222089240bb6903f794d2e0b0e7373116a271d))
* **api:** api update ([f8084e8](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/f8084e83e19e250a9b46a526c4439fbe57c94619))
* enforce OpenSSF Gold quality gates ([696a2c1](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/696a2c1104239667f9182916a5e6474bf810532b))


### Bug Fixes

* make diagnostic dumps concurrency-safe ([6c35deb](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/6c35deb74fab127bcde9947c0a915e90dbad067b))
* make draft release lookup reliable ([577b8c9](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/577b8c9585ce146f61c8f2b20a99696dd28ab680))
* publish assets before immutable release ([acc5ac8](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/acc5ac866388710e5050b27e08c4ad54fb2fc9d1))


### Chores

* add FUNDING.yml pointing to xquik.com ([97cc8a9](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/97cc8a9cb1d7e268721debb5b177fdcdbaf2fad2))
* add per-file license evidence ([ab79e56](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/ab79e56db99acb4aa4069f48fcf469de17dca319))
* avoid failing Context7 refresh without secret ([7401976](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/740197678fc873102a3c8eb3cfdee13f1aebfac0))
* **deps:** bump actions/setup-go from 6.5.0 to 7.0.0 ([#21](https://github.com/Xquik-dev/x-twitter-scraper-go/issues/21)) ([f5bc310](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/f5bc310215ca309302332ffb8cd60b438c8e9d9c))
* prepare v0.5.2 ([e49dced](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/e49dced219d42d7ab36990012e459061ac9f51b9))
* refresh Go SDK dependencies ([9f148cc](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/9f148cc763746a8d68a886ed8a23763ac043ee8a))
* sync repo ([9d47120](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/9d471208569e4262e693f18990eb915e6accb67a))


### Documentation

* add Contributor Covenant 2.1 Code of Conduct ([80c6f70](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/80c6f70b5cf1ace74f29e9e867671d9856bed365))
* add DeepWiki badge ([884f895](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/884f895bd2883295b94a71f0162dab7d13e43df2))
* clarify repository discovery ([26c6fa3](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/26c6fa3b735576ffb942812f0d7b0ddb590058a9))
* map common X data tasks ([39cfa7e](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/39cfa7efd4ee63bf301de922d90367682130fab7))

## 0.12.0 (2026-07-31)

Full Changelog: [v0.11.0...v0.12.0](https://github.com/Xquik-dev/x-twitter-scraper-go/compare/v0.11.0...v0.12.0)

### Features

* **api:** api update ([65b7e46](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/65b7e46a5529edeaae963564862575fed80d0ffa))
* **api:** api update ([f09db18](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/f09db18dbb10b37dd8f8d70cf21410b7f5508279))
* **api:** api update ([7222208](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/72222089240bb6903f794d2e0b0e7373116a271d))
* **api:** api update ([f8084e8](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/f8084e83e19e250a9b46a526c4439fbe57c94619))
* enforce OpenSSF Gold quality gates ([696a2c1](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/696a2c1104239667f9182916a5e6474bf810532b))


### Bug Fixes

* make diagnostic dumps concurrency-safe ([6c35deb](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/6c35deb74fab127bcde9947c0a915e90dbad067b))
* make draft release lookup reliable ([577b8c9](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/577b8c9585ce146f61c8f2b20a99696dd28ab680))
* publish assets before immutable release ([acc5ac8](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/acc5ac866388710e5050b27e08c4ad54fb2fc9d1))


### Chores

* add FUNDING.yml pointing to xquik.com ([97cc8a9](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/97cc8a9cb1d7e268721debb5b177fdcdbaf2fad2))
* add per-file license evidence ([ab79e56](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/ab79e56db99acb4aa4069f48fcf469de17dca319))
* avoid failing Context7 refresh without secret ([7401976](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/740197678fc873102a3c8eb3cfdee13f1aebfac0))
* **deps:** bump actions/setup-go from 6.5.0 to 7.0.0 ([#21](https://github.com/Xquik-dev/x-twitter-scraper-go/issues/21)) ([f5bc310](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/f5bc310215ca309302332ffb8cd60b438c8e9d9c))
* prepare v0.5.2 ([e49dced](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/e49dced219d42d7ab36990012e459061ac9f51b9))
* refresh Go SDK dependencies ([9f148cc](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/9f148cc763746a8d68a886ed8a23763ac043ee8a))
* sync repo ([9d47120](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/9d471208569e4262e693f18990eb915e6accb67a))


### Documentation

* add Contributor Covenant 2.1 Code of Conduct ([80c6f70](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/80c6f70b5cf1ace74f29e9e867671d9856bed365))
* add DeepWiki badge ([884f895](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/884f895bd2883295b94a71f0162dab7d13e43df2))
* clarify repository discovery ([26c6fa3](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/26c6fa3b735576ffb942812f0d7b0ddb590058a9))
* map common X data tasks ([39cfa7e](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/39cfa7efd4ee63bf301de922d90367682130fab7))

## 0.11.0 (2026-07-30)

Full Changelog: [v0.10.0...v0.11.0](https://github.com/Xquik-dev/x-twitter-scraper-go/compare/v0.10.0...v0.11.0)

### Features

* **api:** api update ([65b7e46](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/65b7e46a5529edeaae963564862575fed80d0ffa))
* **api:** api update ([f09db18](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/f09db18dbb10b37dd8f8d70cf21410b7f5508279))
* **api:** api update ([7222208](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/72222089240bb6903f794d2e0b0e7373116a271d))
* **api:** api update ([f8084e8](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/f8084e83e19e250a9b46a526c4439fbe57c94619))
* enforce OpenSSF Gold quality gates ([696a2c1](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/696a2c1104239667f9182916a5e6474bf810532b))


### Bug Fixes

* make diagnostic dumps concurrency-safe ([6c35deb](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/6c35deb74fab127bcde9947c0a915e90dbad067b))
* publish assets before immutable release ([acc5ac8](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/acc5ac866388710e5050b27e08c4ad54fb2fc9d1))


### Chores

* add FUNDING.yml pointing to xquik.com ([97cc8a9](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/97cc8a9cb1d7e268721debb5b177fdcdbaf2fad2))
* add per-file license evidence ([ab79e56](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/ab79e56db99acb4aa4069f48fcf469de17dca319))
* avoid failing Context7 refresh without secret ([7401976](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/740197678fc873102a3c8eb3cfdee13f1aebfac0))
* **deps:** bump actions/setup-go from 6.5.0 to 7.0.0 ([#21](https://github.com/Xquik-dev/x-twitter-scraper-go/issues/21)) ([f5bc310](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/f5bc310215ca309302332ffb8cd60b438c8e9d9c))
* prepare v0.5.2 ([e49dced](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/e49dced219d42d7ab36990012e459061ac9f51b9))
* refresh Go SDK dependencies ([9f148cc](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/9f148cc763746a8d68a886ed8a23763ac043ee8a))
* sync repo ([9d47120](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/9d471208569e4262e693f18990eb915e6accb67a))


### Documentation

* add Contributor Covenant 2.1 Code of Conduct ([80c6f70](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/80c6f70b5cf1ace74f29e9e867671d9856bed365))
* add DeepWiki badge ([884f895](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/884f895bd2883295b94a71f0162dab7d13e43df2))
* clarify repository discovery ([26c6fa3](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/26c6fa3b735576ffb942812f0d7b0ddb590058a9))
* map common X data tasks ([39cfa7e](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/39cfa7efd4ee63bf301de922d90367682130fab7))

## 0.10.0 (2026-07-29)

Full Changelog: [v0.9.0...v0.10.0](https://github.com/Xquik-dev/x-twitter-scraper-go/compare/v0.9.0...v0.10.0)

### Features

* **api:** api update ([e834e64](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/e834e64e11218d8f05e9e6e92a25ff768898baa9))
* **api:** api update ([e54c3dd](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/e54c3dda9a3cd09dd92044e1067cf1f0ab2d508d))
* **api:** api update ([65b7e46](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/65b7e46a5529edeaae963564862575fed80d0ffa))
* **api:** api update ([f09db18](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/f09db18dbb10b37dd8f8d70cf21410b7f5508279))
* **api:** api update ([7222208](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/72222089240bb6903f794d2e0b0e7373116a271d))
* **api:** api update ([f8084e8](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/f8084e83e19e250a9b46a526c4439fbe57c94619))
* enforce OpenSSF Gold quality gates ([696a2c1](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/696a2c1104239667f9182916a5e6474bf810532b))


### Bug Fixes

* make diagnostic dumps concurrency-safe ([6c35deb](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/6c35deb74fab127bcde9947c0a915e90dbad067b))
* publish assets before immutable release ([acc5ac8](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/acc5ac866388710e5050b27e08c4ad54fb2fc9d1))


### Chores

* add FUNDING.yml pointing to xquik.com ([97cc8a9](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/97cc8a9cb1d7e268721debb5b177fdcdbaf2fad2))
* add per-file license evidence ([ab79e56](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/ab79e56db99acb4aa4069f48fcf469de17dca319))
* avoid failing Context7 refresh without secret ([7401976](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/740197678fc873102a3c8eb3cfdee13f1aebfac0))
* **deps:** bump actions/setup-go from 6.5.0 to 7.0.0 ([#21](https://github.com/Xquik-dev/x-twitter-scraper-go/issues/21)) ([f5bc310](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/f5bc310215ca309302332ffb8cd60b438c8e9d9c))
* prepare v0.5.2 ([e49dced](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/e49dced219d42d7ab36990012e459061ac9f51b9))
* refresh Go SDK dependencies ([9f148cc](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/9f148cc763746a8d68a886ed8a23763ac043ee8a))
* sync repo ([9d47120](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/9d471208569e4262e693f18990eb915e6accb67a))


### Documentation

* add Contributor Covenant 2.1 Code of Conduct ([80c6f70](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/80c6f70b5cf1ace74f29e9e867671d9856bed365))
* add DeepWiki badge ([884f895](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/884f895bd2883295b94a71f0162dab7d13e43df2))
* clarify repository discovery ([26c6fa3](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/26c6fa3b735576ffb942812f0d7b0ddb590058a9))
* map common X data tasks ([39cfa7e](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/39cfa7efd4ee63bf301de922d90367682130fab7))

## 0.9.0 (2026-07-28)

Full Changelog: [v0.8.0...v0.9.0](https://github.com/Xquik-dev/x-twitter-scraper-go/compare/v0.8.0...v0.9.0)

### Features

* **api:** api update ([65b7e46](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/65b7e46a5529edeaae963564862575fed80d0ffa))
* **api:** api update ([f09db18](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/f09db18dbb10b37dd8f8d70cf21410b7f5508279))
* **api:** api update ([7222208](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/72222089240bb6903f794d2e0b0e7373116a271d))
* **api:** api update ([f8084e8](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/f8084e83e19e250a9b46a526c4439fbe57c94619))
* enforce OpenSSF Gold quality gates ([696a2c1](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/696a2c1104239667f9182916a5e6474bf810532b))


### Bug Fixes

* make diagnostic dumps concurrency-safe ([6c35deb](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/6c35deb74fab127bcde9947c0a915e90dbad067b))
* publish assets before immutable release ([acc5ac8](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/acc5ac866388710e5050b27e08c4ad54fb2fc9d1))


### Chores

* add FUNDING.yml pointing to xquik.com ([97cc8a9](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/97cc8a9cb1d7e268721debb5b177fdcdbaf2fad2))
* add per-file license evidence ([ab79e56](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/ab79e56db99acb4aa4069f48fcf469de17dca319))
* avoid failing Context7 refresh without secret ([7401976](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/740197678fc873102a3c8eb3cfdee13f1aebfac0))
* **deps:** bump actions/setup-go from 6.5.0 to 7.0.0 ([#21](https://github.com/Xquik-dev/x-twitter-scraper-go/issues/21)) ([f5bc310](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/f5bc310215ca309302332ffb8cd60b438c8e9d9c))
* prepare v0.5.2 ([e49dced](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/e49dced219d42d7ab36990012e459061ac9f51b9))
* refresh Go SDK dependencies ([9f148cc](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/9f148cc763746a8d68a886ed8a23763ac043ee8a))
* sync repo ([9d47120](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/9d471208569e4262e693f18990eb915e6accb67a))


### Documentation

* add Contributor Covenant 2.1 Code of Conduct ([80c6f70](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/80c6f70b5cf1ace74f29e9e867671d9856bed365))
* add DeepWiki badge ([884f895](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/884f895bd2883295b94a71f0162dab7d13e43df2))
* clarify repository discovery ([26c6fa3](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/26c6fa3b735576ffb942812f0d7b0ddb590058a9))
* map common X data tasks ([39cfa7e](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/39cfa7efd4ee63bf301de922d90367682130fab7))

## 0.8.0 (2026-07-25)

Full Changelog: [v0.7.1...v0.8.0](https://github.com/Xquik-dev/x-twitter-scraper-go/compare/v0.7.1...v0.8.0)

### Features

* **api:** api update ([65b7e46](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/65b7e46a5529edeaae963564862575fed80d0ffa))
* **api:** api update ([f09db18](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/f09db18dbb10b37dd8f8d70cf21410b7f5508279))
* **api:** api update ([7222208](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/72222089240bb6903f794d2e0b0e7373116a271d))
* **api:** api update ([f8084e8](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/f8084e83e19e250a9b46a526c4439fbe57c94619))
* enforce OpenSSF Gold quality gates ([696a2c1](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/696a2c1104239667f9182916a5e6474bf810532b))


### Bug Fixes

* make diagnostic dumps concurrency-safe ([6c35deb](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/6c35deb74fab127bcde9947c0a915e90dbad067b))
* publish assets before immutable release ([acc5ac8](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/acc5ac866388710e5050b27e08c4ad54fb2fc9d1))


### Chores

* add FUNDING.yml pointing to xquik.com ([97cc8a9](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/97cc8a9cb1d7e268721debb5b177fdcdbaf2fad2))
* add per-file license evidence ([ab79e56](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/ab79e56db99acb4aa4069f48fcf469de17dca319))
* avoid failing Context7 refresh without secret ([7401976](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/740197678fc873102a3c8eb3cfdee13f1aebfac0))
* prepare v0.5.2 ([e49dced](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/e49dced219d42d7ab36990012e459061ac9f51b9))
* refresh Go SDK dependencies ([9f148cc](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/9f148cc763746a8d68a886ed8a23763ac043ee8a))
* sync repo ([9d47120](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/9d471208569e4262e693f18990eb915e6accb67a))


### Documentation

* add Contributor Covenant 2.1 Code of Conduct ([80c6f70](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/80c6f70b5cf1ace74f29e9e867671d9856bed365))
* add DeepWiki badge ([884f895](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/884f895bd2883295b94a71f0162dab7d13e43df2))
* clarify repository discovery ([26c6fa3](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/26c6fa3b735576ffb942812f0d7b0ddb590058a9))
* map common X data tasks ([39cfa7e](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/39cfa7efd4ee63bf301de922d90367682130fab7))

## 0.7.1 (2026-07-24)

Full Changelog: [v0.7.0...v0.7.1](https://github.com/Xquik-dev/x-twitter-scraper-go/compare/v0.7.0...v0.7.1)

### Bug Fixes

* publish assets before immutable release ([acc5ac8](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/acc5ac866388710e5050b27e08c4ad54fb2fc9d1))

## 0.7.0 (2026-07-24)

Full Changelog: [v0.6.2...v0.7.0](https://github.com/Xquik-dev/x-twitter-scraper-go/compare/v0.6.2...v0.7.0)

### Features

* enforce OpenSSF Gold quality gates ([696a2c1](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/696a2c1104239667f9182916a5e6474bf810532b))


### Bug Fixes

* make diagnostic dumps concurrency-safe ([6c35deb](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/6c35deb74fab127bcde9947c0a915e90dbad067b))


### Chores

* add per-file license evidence ([ab79e56](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/ab79e56db99acb4aa4069f48fcf469de17dca319))


### Documentation

* clarify repository discovery ([26c6fa3](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/26c6fa3b735576ffb942812f0d7b0ddb590058a9))

## 0.6.2 (2026-07-23)

Full Changelog: [v0.6.1...v0.6.2](https://github.com/Xquik-dev/x-twitter-scraper-go/compare/v0.6.1...v0.6.2)

### Documentation

* map common X data tasks ([39cfa7e](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/39cfa7efd4ee63bf301de922d90367682130fab7))

## 0.6.1 (2026-07-22)

Full Changelog: [v0.6.0...v0.6.1](https://github.com/Xquik-dev/x-twitter-scraper-go/compare/v0.6.0...v0.6.1)

### Chores

* sync generated SDK contracts ([bd021dd](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/bd021dd68f93c08f042cd97dd345cc3abf7ae53b))
* verify final public contract ([20613fa](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/20613fa74d190dc880a717893507dcce0a6c63b5))

## 0.6.0 (2026-07-22)

Full Changelog: [v0.5.2...v0.6.0](https://github.com/Xquik-dev/x-twitter-scraper-go/compare/v0.5.2...v0.6.0)

### Features

* **api:** propagate durable write actions ([64520c8](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/64520c82ca6684096f078bcaff6c39c1e8b28797))
* **api:** propagate production write-action contract ([56e4783](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/56e47832ac93e1d7057b7d4fe8b8452798acd66f))

## 0.5.1 (2026-07-14)

Full Changelog: [v0.5.0...v0.5.1](https://github.com/Xquik-dev/x-twitter-scraper-go/compare/v0.5.0...v0.5.1)

## 0.5.0 (2026-07-14)

Full Changelog: [v0.4.2...v0.5.0](https://github.com/Xquik-dev/x-twitter-scraper-go/compare/v0.4.2...v0.5.0)

### Features

* add accountless guest wallet checkout, status, and top-up operations
* add credit top-up redirect and status operations
* add keyword monitor and account connection challenge resources
* expand typed read, write, extraction, event, and webhook contracts

### Chores

* refresh generated request encoding and transport helpers
* pin GitHub Actions used by CI

## 0.4.2 (2026-07-13)

Full Changelog: [v0.4.1...v0.4.2](https://github.com/Xquik-dev/x-twitter-scraper-go/compare/v0.4.1...v0.4.2)

### Documentation

* add the Xquik non-affiliation notice
* improve SDK discovery and Context7 guidance

### Chores

* refresh Go SDK dependencies

## 0.4.1 (2026-04-25)

Full Changelog: [v0.4.0...v0.4.1](https://github.com/Xquik-dev/x-twitter-scraper-go/compare/v0.4.0...v0.4.1)

### Bug Fixes

* escape ampersand in OpenAPI summaries for C# XML docs ([66d1655](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/66d165598acb070e5fb94b2df46204618ebc7d85))


### Chores

* **internal:** more robust bootstrap script ([bc2d3bb](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/bc2d3bb3585703357f63126e9a3ead9496699f3c))


### Documentation

* add Contributor Covenant 2.1 Code of Conduct ([80c6f70](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/80c6f70b5cf1ace74f29e9e867671d9856bed365))

## 0.4.0 (2026-04-22)

Full Changelog: [v0.3.0...v0.4.0](https://github.com/Xquik-dev/x-twitter-scraper-go/compare/v0.3.0...v0.4.0)

### Features

* **api:** api update ([984db84](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/984db84082cbeab503df44e06376dfb3badb2a49))
* **api:** api update ([4274efa](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/4274efae4c2859961dafffe05944efbc3aad509f))
* **api:** api update ([f907b7d](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/f907b7db089914b772f78fd4ca928cb4cc916815))
* **api:** api update ([ff4b0bf](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/ff4b0bfccc8307b6300dbf9a3ecd26f8ae234067))
* **api:** api update ([1c60929](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/1c60929a9d485a5a4adad98e34402632a13a22bc))


### Bug Fixes

* move SearchTweet & UserProfile to shared models to resolve Go import cycle ([4188d03](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/4188d03924c0b7c9072ec4095efe54bf9056d094))


### Chores

* sync OpenAPI spec ([2fba029](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/2fba029d3dfa17b72647265fecdb80f0f1c9623f))
* wire production_repo for all targets ([bea11a5](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/bea11a51e898b5663d45845f9b2f54da1ab98a1a))

## 0.3.0 (2026-04-08)

Full Changelog: [v0.2.0...v0.3.0](https://github.com/Xquik-dev/x-twitter-scraper-go/compare/v0.2.0...v0.3.0)

### Features

* **api:** api update ([bc5f66c](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/bc5f66c0ea5d0757a20dbf7a53db2fabeb867735))
* **api:** api update ([4d74a83](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/4d74a83c6125f9f97113fc2a1f9ec85147ff6c79))
* **api:** api update ([b9d3ddc](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/b9d3ddca43cad5e4372046c0c0d45369369fcea6))
* **api:** api update ([e404a12](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/e404a12b5a488c542ca38f5dae167d53793d0196))
* **api:** api update ([4950584](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/4950584aeaeb08cb95b0ebdc082e892852b9db09))
* **api:** api update ([5a54946](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/5a549469eaef079514a4a3560f65925c04fea137))
* **api:** api update ([7e5b55c](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/7e5b55c9b5d823c5e76e1068d0b3805abc3791f9))
* **api:** api update ([6e635f2](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/6e635f2ee3944937be7685ab54bb4f54e4c5c975))


### Chores

* update SDK settings ([4111b6c](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/4111b6c98cedb4d318ba84109fe3c9213cf5661b))

## 0.2.0 (2026-04-01)

Full Changelog: [v0.1.0...v0.2.0](https://github.com/Xquik-dev/x-twitter-scraper-go/compare/v0.1.0...v0.2.0)

### Features

* **api:** api update ([be87d38](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/be87d381ae77f7a66686045a62280f2768b2f5ed))
* **api:** api update ([a612b9e](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/a612b9e0034f7fac69e5c0cf6f03d7d8311c7570))


### Chores

* update SDK settings ([3ea8aff](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/3ea8aff9f6b696c52f37339a8321df6ed2bb5f4e))

## 0.1.0 (2026-03-30)

Full Changelog: [v0.0.1...v0.1.0](https://github.com/Xquik-dev/x-twitter-scraper-go/compare/v0.0.1...v0.1.0)

### Features

* **api:** api update ([f09db18](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/f09db18dbb10b37dd8f8d70cf21410b7f5508279))
* **api:** api update ([7222208](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/72222089240bb6903f794d2e0b0e7373116a271d))
* **api:** api update ([f8084e8](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/f8084e83e19e250a9b46a526c4439fbe57c94619))


### Chores

* update SDK settings ([b056ca2](https://github.com/Xquik-dev/x-twitter-scraper-go/commit/b056ca25f53816416c8173777c77e35204544970))
