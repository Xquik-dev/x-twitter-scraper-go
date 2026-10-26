// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

package xtwitterscraper

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/Xquik-dev/x-twitter-scraper-go/internal/apijson"
	"github.com/Xquik-dev/x-twitter-scraper-go/internal/apiquery"
	"github.com/Xquik-dev/x-twitter-scraper-go/internal/requestconfig"
	"github.com/Xquik-dev/x-twitter-scraper-go/option"
	"github.com/Xquik-dev/x-twitter-scraper-go/packages/param"
	"github.com/Xquik-dev/x-twitter-scraper-go/packages/respjson"
	"github.com/Xquik-dev/x-twitter-scraper-go/shared"
	"github.com/Xquik-dev/x-twitter-scraper-go/shared/constant"
)

// XTweetService contains methods and other services that help with interacting
// with the x-twitter-scraper API.
//
// You should not instantiate this service directly, and instead use the
// [NewXTweetService] method instead.
type XTweetService struct {
	options []option.RequestOption
	// X write actions (tweets, likes, follows, DMs)
	Like XTweetLikeService
	// X write actions (tweets, likes, follows, DMs)
	Retweet XTweetRetweetService
}

// NewXTweetService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewXTweetService(opts ...option.RequestOption) (r XTweetService) {
	r = XTweetService{}
	r.options = opts
	r.Like = NewXTweetLikeService(opts...)
	r.Retweet = NewXTweetRetweetService(opts...)
	return
}

// Publishes a post through a connected X account.
func (r *XTweetService) New(ctx context.Context, params XTweetNewParams, opts ...option.RequestOption) (res *XTweetNewResponse, err error) {
	if !param.IsOmitted(params.IdempotencyKey) {
		opts = append(opts, option.WithHeader("Idempotency-Key", fmt.Sprintf("%v", params.IdempotencyKey)))
	}
	opts = slices.Concat(r.options, opts)
	path := "x/tweets"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Returns one public tweet with author, metrics, and media.
func (r *XTweetService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *XTweetGetResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("x/tweets/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Returns public tweet records for the requested IDs.
func (r *XTweetService) List(ctx context.Context, query XTweetListParams, opts ...option.RequestOption) (res *shared.PaginatedTweets, err error) {
	opts = slices.Concat(r.options, opts)
	path := "x/tweets"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Deletes an authored post through a connected X account.
func (r *XTweetService) Delete(ctx context.Context, id string, params XTweetDeleteParams, opts ...option.RequestOption) (res *XTweetDeleteResponse, err error) {
	if !param.IsOmitted(params.IdempotencyKey) {
		opts = append(opts, option.WithHeader("Idempotency-Key", fmt.Sprintf("%v", params.IdempotencyKey)))
	}
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("x/tweets/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, params, &res, opts...)
	return res, err
}

// Returns whether 1 account reposted the post, read from its newest posts. A
// repost older than the posts read returns false with complete false. Each check
// costs 1 credit.
func (r *XTweetService) CheckRepost(ctx context.Context, id string, query XTweetCheckRepostParams, opts ...option.RequestOption) (res *XTweetCheckRepostResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("x/tweets/%s/retweeters/check", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Returns the HTML that X gives websites to embed 1 public tweet, with the tweet's
// link and author.
func (r *XTweetService) Embed(ctx context.Context, id string, query XTweetEmbedParams, opts ...option.RequestOption) (res *XTweetEmbedResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("x/tweets/%s/embed", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Returns liker profiles that X makes visible for the post. X can withhold liker
// identities even when the post reports likes. In that case this endpoint returns
// 424 `favoriters_unavailable` instead of a misleading empty success.
func (r *XTweetService) GetFavoriters(ctx context.Context, id string, query XTweetGetFavoritersParams, opts ...option.RequestOption) (res *shared.PaginatedUsers, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("x/tweets/%s/favoriters", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Returns the replies the post's author hid, oldest first, with the replies to
// them. A post without hidden replies returns an empty list.
func (r *XTweetService) GetHiddenReplies(ctx context.Context, id string, query XTweetGetHiddenRepliesParams, opts ...option.RequestOption) (res *shared.PaginatedTweets, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("x/tweets/%s/hidden-replies", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Returns engagement insights for 1 tweet. They are engagements per hour since it
// went up, per view and per follower of its author, and how far it spread. Each
// report costs 1 credit, whatever it reads.
func (r *XTweetService) GetInsights(ctx context.Context, id string, query XTweetGetInsightsParams, opts ...option.RequestOption) (res *XTweetGetInsightsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("x/tweets/%s/insights", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Returns public posts quoting the selected tweet.
func (r *XTweetService) GetQuotes(ctx context.Context, id string, query XTweetGetQuotesParams, opts ...option.RequestOption) (res *shared.PaginatedTweets, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("x/tweets/%s/quotes", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Returns direct replies with automatic maximum coverage and pagination. Complete
// mode adds nested replies, diagnostics, and a 424 below 80% coverage.
func (r *XTweetService) GetReplies(ctx context.Context, id string, query XTweetGetRepliesParams, opts ...option.RequestOption) (res *XTweetGetRepliesResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("x/tweets/%s/replies", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Returns the user IDs of a post's reposters, in the order & pages of
// /x/tweets/{id}/retweeters. Use it when you need IDs only. Each ID costs 1
// credit. A post nobody reposted returns an empty list for free.
func (r *XTweetService) GetRetweeterIDs(ctx context.Context, id string, query XTweetGetRetweeterIDsParams, opts ...option.RequestOption) (res *XTweetGetRetweeterIDsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("x/tweets/%s/retweeter-ids", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Lists reposting profiles. Optional timestamps come from the newest profile page.
// Extra reads add latency. Missing or failed matches leave null timestamps.
func (r *XTweetService) GetRetweeters(ctx context.Context, id string, query XTweetGetRetweetersParams, opts ...option.RequestOption) (res *shared.PaginatedUsers, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("x/tweets/%s/retweeters", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Returns the caption tracks of 1 public tweet's videos, as X's player reads them.
// Each track has its language, name, timed cues, and full text. A tweet without
// video, or a video without captions, returns no track and costs nothing.
func (r *XTweetService) GetSubtitles(ctx context.Context, id string, query XTweetGetSubtitlesParams, opts ...option.RequestOption) (res *XTweetGetSubtitlesResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("x/tweets/%s/subtitles", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Returns visible posts from the selected conversation thread.
func (r *XTweetService) GetThread(ctx context.Context, id string, query XTweetGetThreadParams, opts ...option.RequestOption) (res *shared.PaginatedTweets, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("x/tweets/%s/thread", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Returns the tweets X ranks highest by 1 metric over the last day, week or month,
// as x.com's Inspiration page shows them. Covers all of X unless you pick a
// country or a language. Pages hold up to 20 tweets, and the list ends at 100.
// Each tweet costs 1 credit.
func (r *XTweetService) ListTop(ctx context.Context, query XTweetListTopParams, opts ...option.RequestOption) (res *shared.PaginatedTweets, err error) {
	opts = slices.Concat(r.options, opts)
	path := "x/tweets/top"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Returns where each t.co link leads, as X's own redirect says. Xquik reads the
// redirect and never visits the target site. A link t.co does not have comes back
// without `resolvedUrl` and costs nothing. A link X does not answer for in time
// comes back in `unprocessedUrls`, free. Each resolved link costs 1 credit.
func (r *XTweetService) ResolveLinks(ctx context.Context, query XTweetResolveLinksParams, opts ...option.RequestOption) (res *XTweetResolveLinksResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "x/links/resolve"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Returns normalized tweets with author, like count, media, URL, and cursors. Set
// q, limit, queryType, and minLikes. Omit mode for maximum coverage. Reuse
// next_cursor without changing filters.
func (r *XTweetService) Search(ctx context.Context, query XTweetSearchParams, opts ...option.RequestOption) (res *XTweetSearchResponseUnion, err error) {
	opts = slices.Concat(r.options, opts)
	path := "x/tweets/search"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Returns X's own translation of 1 public tweet into the language you pick, with
// the links, mentions, and tags X places in the translated text.
func (r *XTweetService) Translate(ctx context.Context, id string, query XTweetTranslateParams, opts ...option.RequestOption) (res *XTweetTranslateResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("x/tweets/%s/translation", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Tweet author profile. The lookup route always includes follower count and
// verification state. Other profile fields appear when available.
type TweetAuthor struct {
	Followers int64 `json:"followers" api:"required"`
	Verified  bool  `json:"verified" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Followers   respjson.Field
		Verified    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.UserProfile
}

// Returns the unmodified JSON received from the API
func (r TweetAuthor) RawJSON() string { return r.JSON.raw }
func (r *TweetAuthor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// Full tweet with text, engagement metrics, media, and metadata. A zero metric can
// mean X did not report the count.
type TweetDetail struct {
	ID            string `json:"id" api:"required"`
	BookmarkCount int64  `json:"bookmarkCount" api:"required"`
	LikeCount     int64  `json:"likeCount" api:"required"`
	QuoteCount    int64  `json:"quoteCount" api:"required"`
	ReplyCount    int64  `json:"replyCount" api:"required"`
	RetweetCount  int64  `json:"retweetCount" api:"required"`
	Text          string `json:"text" api:"required"`
	ViewCount     int64  `json:"viewCount" api:"required"`
	// Describes an X Article preview and its lifecycle metadata.
	Article TweetDetailArticle `json:"article"`
	// Tweet author profile. The lookup route always includes follower count and
	// verification state. Other profile fields appear when available.
	Author TweetAuthor `json:"author"`
	// Why the tweet has no author. X serves a suspended or deleted account as an
	// unavailable user with a reason such as Suspended or NotFound. Omitted when the
	// author is available.
	AuthorUnavailable TweetDetailAuthorUnavailable `json:"authorUnavailable"`
	// Describes a public card and its referenced profiles.
	Card TweetDetailCard `json:"card"`
	// Price charts the author attached under the post, 1 per cashtag. Absent when the
	// post has none. X may send a cashtag without a chart window.
	CashtagAttachments []TweetDetailCashtagAttachment `json:"cashtagAttachments"`
	// Community ID.
	CommunityID string `json:"communityId"`
	// Community name. Present when X sends it with the post.
	CommunityName string `json:"communityName"`
	// Community Note presentation metadata returned by X.
	CommunityNote TweetDetailCommunityNote `json:"communityNote"`
	// Content disclosure metadata shown by X when a tweet is labeled as paid
	// partnership content or AI-generated media.
	ContentDisclosure shared.ContentDisclosure `json:"contentDisclosure"`
	// Public reply policy and conversation owner.
	ConversationControl TweetDetailConversationControl `json:"conversationControl"`
	ConversationID      string                         `json:"conversationId"`
	// Author-published point, with longitude before latitude.
	Coordinates      TweetDetailCoordinates `json:"coordinates"`
	CreatedAt        string                 `json:"createdAt"`
	DisplayTextRange []int64                `json:"displayTextRange"`
	// Lists edit-chain identifiers and the remaining edit window.
	Edit TweetDetailEdit `json:"edit"`
	// Lists hashtags, symbols, links, and mentions from tweet text.
	Entities TweetDetailEntities `json:"entities"`
	// Marks a subscriber-only post with the creator whose subscription unlocks it.
	// Non-subscribers get the preview text & the subscribe link. Omitted on public
	// posts.
	ExclusiveContent TweetDetailExclusiveContent `json:"exclusiveContent"`
	// Author-published point, with latitude before longitude.
	Geo TweetDetailGeo `json:"geo"`
	// Public metadata whose fields are defined by X.
	GrokShareAttachment PublicMetadata `json:"grokShareAttachment"`
	// Grok's translation of the post, as X shows it.
	GrokTranslatedPost PublicMetadata `json:"grokTranslatedPost"`
	// Whether X holds Community Notes on the post, shown or not.
	HasCommunityNotes bool   `json:"hasCommunityNotes"`
	InReplyToID       string `json:"inReplyToId"`
	InReplyToUserID   string `json:"inReplyToUserId"`
	InReplyToUsername string `json:"inReplyToUsername"`
	// The notice x.com covers the post with until the reader views it, such as X's
	// public interest notice on a post that broke its rules but stays up.
	Interstitial   TweetDetailInterstitial `json:"interstitial"`
	IsLimitedReply bool                    `json:"isLimitedReply"`
	IsNoteTweet    bool                    `json:"isNoteTweet"`
	IsPinned       bool                    `json:"isPinned"`
	IsQuoteStatus  bool                    `json:"isQuoteStatus"`
	IsReply        bool                    `json:"isReply"`
	IsRetweet      bool                    `json:"isRetweet"`
	IsTranslatable bool                    `json:"isTranslatable"`
	// Public metadata whose fields are defined by X.
	JetfuelAttachment PublicMetadata `json:"jetfuelAttachment"`
	Lang              string         `json:"lang"`
	// Public interaction restrictions and user-facing prompts.
	LimitedActions []TweetDetailLimitedAction `json:"limitedActions"`
	// Attached media items, omitted when unavailable.
	Media []shared.TweetMedia `json:"media"`
	// Complete Note Tweet content and rich-text metadata.
	NoteTweet TweetDetailNoteTweet `json:"noteTweet"`
	// A parent, quoted or retweeted tweet. It has the TweetPublicData fields and an
	// author. Nesting stops after 4 levels.
	ParentTweet shared.EmbeddedTweet `json:"parentTweet"`
	// Describes public place metadata on a geotagged tweet.
	Place             TweetDetailPlace `json:"place"`
	PossiblySensitive bool             `json:"possiblySensitive"`
	// Public metadata whose fields are defined by X.
	PostCta PublicMetadata `json:"postCta"`
	// Engagement counts retained from a prior tweet edit.
	PreviousCounts TweetDetailPreviousCounts `json:"previousCounts"`
	// Quoted or retweeted tweet context.
	QuotedTweet shared.EmbeddedTweet `json:"quoted_tweet"`
	// Quoted tweet ID.
	QuotedTweetID string `json:"quotedTweetId"`
	// The quoted post's link, kept when X withholds that post.
	QuotedTweetPermalink TweetDetailQuotedTweetPermalink `json:"quotedTweetPermalink"`
	// Why X shows no quoted post. Omitted when X names no reason, as its legacy
	// answers never do.
	QuotedTweetUnavailable TweetDetailQuotedTweetUnavailable `json:"quotedTweetUnavailable"`
	// Public post and user referenced by this reaction.
	ReactionContext TweetDetailReactionContext `json:"reactionContext"`
	// Quoted or retweeted tweet context.
	RetweetedTweet shared.EmbeddedTweet `json:"retweeted_tweet"`
	// Repost event time in UTC. Null when unavailable; omitted for other posts.
	RetweetedAt time.Time `json:"retweetedAt" api:"nullable" format:"date-time"`
	// Public metadata whose fields are defined by X.
	Scopes PublicMetadata `json:"scopes"`
	// The label X shows under a post whose reach it limits. It links to the page
	// that explains it.
	SoftIntervention shared.SoftIntervention `json:"softIntervention"`
	Source           string                  `json:"source"`
	// Sports game context attached to the post, when available.
	SportsContext TweetDetailSportsContext `json:"sportsContext"`
	// The Subscriber badge X shows on a reply. The reply's author subscribes to the
	// creator who started the conversation. Omitted on a post without the badge.
	SubscriberBadge TweetDetailSubscriberBadge `json:"subscriberBadge"`
	// Where `text` matches the search, on each row a search read. Offsets count code
	// points of `text`, as `displayTextRange` does. Empty when X marks nothing in the
	// post. X marks matches only in a post's first 280 characters.
	TextHighlights []TweetDetailTextHighlight `json:"textHighlights"`
	// Public visibility notice attached to an available tweet.
	Tombstone TweetDetailTombstone `json:"tombstone"`
	Type      string               `json:"type"`
	// The author's ID when X withholds the author's profile.
	UnavailableAuthorID string `json:"unavailableAuthorId"`
	// X did not send these counts for this post, so treat each 0 as unknown.
	UnknownCounts []shared.TweetUnknownCount `json:"unknownCounts"`
	// User IDs that left this conversation.
	UnmentionedUserIDs []string `json:"unmentionedUserIds"`
	URL                string   `json:"url"`
	ViewState          string   `json:"viewState"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                     respjson.Field
		BookmarkCount          respjson.Field
		LikeCount              respjson.Field
		QuoteCount             respjson.Field
		ReplyCount             respjson.Field
		RetweetCount           respjson.Field
		Text                   respjson.Field
		ViewCount              respjson.Field
		Article                respjson.Field
		Author                 respjson.Field
		AuthorUnavailable      respjson.Field
		Card                   respjson.Field
		CashtagAttachments     respjson.Field
		CommunityID            respjson.Field
		CommunityName          respjson.Field
		CommunityNote          respjson.Field
		ContentDisclosure      respjson.Field
		ConversationControl    respjson.Field
		ConversationID         respjson.Field
		Coordinates            respjson.Field
		CreatedAt              respjson.Field
		DisplayTextRange       respjson.Field
		Edit                   respjson.Field
		Entities               respjson.Field
		ExclusiveContent       respjson.Field
		Geo                    respjson.Field
		GrokShareAttachment    respjson.Field
		GrokTranslatedPost     respjson.Field
		HasCommunityNotes      respjson.Field
		InReplyToID            respjson.Field
		InReplyToUserID        respjson.Field
		InReplyToUsername      respjson.Field
		Interstitial           respjson.Field
		IsLimitedReply         respjson.Field
		IsNoteTweet            respjson.Field
		IsPinned               respjson.Field
		IsQuoteStatus          respjson.Field
		IsReply                respjson.Field
		IsRetweet              respjson.Field
		IsTranslatable         respjson.Field
		JetfuelAttachment      respjson.Field
		Lang                   respjson.Field
		LimitedActions         respjson.Field
		Media                  respjson.Field
		NoteTweet              respjson.Field
		ParentTweet            respjson.Field
		Place                  respjson.Field
		PossiblySensitive      respjson.Field
		PostCta                respjson.Field
		PreviousCounts         respjson.Field
		QuotedTweet            respjson.Field
		QuotedTweetID          respjson.Field
		QuotedTweetPermalink   respjson.Field
		QuotedTweetUnavailable respjson.Field
		ReactionContext        respjson.Field
		RetweetedTweet         respjson.Field
		RetweetedAt            respjson.Field
		Scopes                 respjson.Field
		SoftIntervention       respjson.Field
		Source                 respjson.Field
		SportsContext          respjson.Field
		SubscriberBadge        respjson.Field
		TextHighlights         respjson.Field
		Tombstone              respjson.Field
		Type                   respjson.Field
		UnavailableAuthorID    respjson.Field
		UnknownCounts          respjson.Field
		UnmentionedUserIDs     respjson.Field
		URL                    respjson.Field
		ViewState              respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TweetDetail) RawJSON() string { return r.JSON.raw }
func (r *TweetDetail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// The quoted post's link, kept when X withholds that post.
type TweetDetailQuotedTweetPermalink struct {
	DisplayURL  string `json:"displayUrl"`
	ExpandedURL string `json:"expandedUrl"`
	URL         string `json:"url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DisplayURL  respjson.Field
		ExpandedURL respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TweetDetailQuotedTweetPermalink) RawJSON() string { return r.JSON.raw }
func (r *TweetDetailQuotedTweetPermalink) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// Why X shows no quoted post. Omitted when X names no reason, as its legacy
// answers never do.
type TweetDetailQuotedTweetUnavailable struct {
	// X's reason, such as `Suspended` for a suspended author or `Protected` for an
	// author who protects their posts.
	Reason string `json:"reason" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Reason      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TweetDetailQuotedTweetUnavailable) RawJSON() string { return r.JSON.raw }
func (r *TweetDetailQuotedTweetUnavailable) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type TweetDetailTextHighlight struct {
	// The first code point after the match.
	EndIndex int64 `json:"endIndex" api:"required"`
	// The first code point of the match.
	StartIndex int64 `json:"startIndex" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EndIndex    respjson.Field
		StartIndex  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TweetDetailTextHighlight) RawJSON() string { return r.JSON.raw }
func (r *TweetDetailTextHighlight) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// Why the tweet has no author. X serves a suspended or deleted account as an
// unavailable user with a reason such as Suspended or NotFound. Omitted when the
// author is available.
type TweetDetailAuthorUnavailable struct {
	// X's reason code.
	Reason string `json:"reason" api:"required"`
	// X's notice text for the reason.
	Message string `json:"message"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Reason      respjson.Field
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TweetDetailAuthorUnavailable) RawJSON() string { return r.JSON.raw }
func (r *TweetDetailAuthorUnavailable) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type TweetDetailCashtagAttachment struct {
	// The cashtag with its `$`.
	Cashtag string `json:"cashtag" api:"required"`
	// Span of each candle.
	CandleSize string `json:"candleSize"`
	// Time of the chart's last price.
	EndTime time.Time `json:"endTime" format:"date-time"`
	// X's size for the chart.
	SizeVariant int64 `json:"sizeVariant"`
	// Time of the chart's first candle.
	StartTime time.Time `json:"startTime" format:"date-time"`
	// Span the chart covers.
	Timeframe string `json:"timeframe"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Cashtag     respjson.Field
		CandleSize  respjson.Field
		EndTime     respjson.Field
		SizeVariant respjson.Field
		StartTime   respjson.Field
		Timeframe   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TweetDetailCashtagAttachment) RawJSON() string { return r.JSON.raw }
func (r *TweetDetailCashtagAttachment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// Marks a subscriber-only post with the creator whose subscription unlocks it.
// Non-subscribers get the preview text & the subscribe link. Omitted on public
// posts.
type TweetDetailExclusiveContent struct {
	// Username of the creator whose subscription unlocks the post.
	CreatorUserName string `json:"creatorUserName" api:"required"`
	// Subscribe link X shows non-subscribers. Omitted when X serves the full post.
	SubscribeURL string `json:"subscribeUrl"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatorUserName respjson.Field
		SubscribeURL    respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TweetDetailExclusiveContent) RawJSON() string { return r.JSON.raw }
func (r *TweetDetailExclusiveContent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// Author-published point, with longitude before latitude.
type TweetDetailCoordinates struct {
	Coordinates []float64      `json:"coordinates" api:"required"`
	Type        constant.Point `json:"type" default:"Point"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Coordinates respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TweetDetailCoordinates) RawJSON() string { return r.JSON.raw }
func (r *TweetDetailCoordinates) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// Author-published point, with latitude before longitude.
type TweetDetailGeo struct {
	Coordinates []float64      `json:"coordinates" api:"required"`
	Type        constant.Point `json:"type" default:"Point"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Coordinates respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TweetDetailGeo) RawJSON() string { return r.JSON.raw }
func (r *TweetDetailGeo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// The Subscriber badge X shows on a reply. The reply's author subscribes to the
// creator who started the conversation. Omitted on a post without the badge.
type TweetDetailSubscriberBadge struct {
	// Account ID of the creator the author subscribes to.
	CreatorID string `json:"creatorId" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatorID   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TweetDetailSubscriberBadge) RawJSON() string { return r.JSON.raw }
func (r *TweetDetailSubscriberBadge) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// The notice x.com covers the post with until the reader views it, such as X's
// public interest notice on a post that broke its rules but stays up.
type TweetDetailInterstitial struct {
	// How x.com shows the notice. `NonCompliant` marks a post that broke X's rules.
	DisplayType string                      `json:"displayType"`
	Text        TweetDetailInterstitialText `json:"text"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DisplayType respjson.Field
		Text        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TweetDetailInterstitial) RawJSON() string { return r.JSON.raw }
func (r *TweetDetailInterstitial) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type TweetDetailInterstitialText struct {
	// Links inside the notice, by character range.
	Entities []TweetDetailInterstitialTextEntity `json:"entities"`
	// Whether X writes it right to left.
	Rtl bool `json:"rtl"`
	// The notice's text.
	Text string `json:"text"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Entities    respjson.Field
		Rtl         respjson.Field
		Text        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TweetDetailInterstitialText) RawJSON() string { return r.JSON.raw }
func (r *TweetDetailInterstitialText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type TweetDetailInterstitialTextEntity struct {
	// Whole number from 0.
	FromIndex int64                                `json:"fromIndex"`
	Ref       TweetDetailInterstitialTextEntityRef `json:"ref"`
	// Whole number from 0.
	ToIndex int64 `json:"toIndex"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FromIndex   respjson.Field
		Ref         respjson.Field
		ToIndex     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TweetDetailInterstitialTextEntity) RawJSON() string { return r.JSON.raw }
func (r *TweetDetailInterstitialTextEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type TweetDetailInterstitialTextEntityRef struct {
	ScreenName string `json:"screenName"`
	Text       string `json:"text"`
	Type       string `json:"type"`
	URL        string `json:"url"`
	URLType    string `json:"urlType"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ScreenName  respjson.Field
		Text        respjson.Field
		Type        respjson.Field
		URL         respjson.Field
		URLType     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TweetDetailInterstitialTextEntityRef) RawJSON() string { return r.JSON.raw }
func (r *TweetDetailInterstitialTextEntityRef) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// Describes an X Article preview and its lifecycle metadata.
type TweetDetailArticle struct {
	ID string `json:"id"`
	// Public metadata whose fields are defined by X.
	CoverMedia    PublicMetadata `json:"coverMedia"`
	CoverMediaURL string         `json:"coverMediaUrl"`
	// Public metadata whose fields are defined by X.
	LifecycleState PublicMetadata `json:"lifecycleState"`
	// Public metadata whose fields are defined by X.
	Metadata    PublicMetadata `json:"metadata"`
	PreviewText string         `json:"previewText"`
	Title       string         `json:"title"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		CoverMedia     respjson.Field
		CoverMediaURL  respjson.Field
		LifecycleState respjson.Field
		Metadata       respjson.Field
		PreviewText    respjson.Field
		Title          respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TweetDetailArticle) RawJSON() string { return r.JSON.raw }
func (r *TweetDetailArticle) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// Describes a public card and its referenced profiles.
type TweetDetailCard struct {
	ID string `json:"id"`
	// Public metadata whose fields are defined by X.
	BindingValues PublicMetadata `json:"bindingValues"`
	Name          string         `json:"name"`
	// Public metadata whose fields are defined by X.
	Platform PublicMetadata `json:"platform"`
	URL      string         `json:"url"`
	// Unresolved card user references.
	UserReferenceErrors []TweetDetailCardUserReferenceError `json:"userReferenceErrors"`
	UserReferences      []shared.UserProfile                `json:"userReferences"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                  respjson.Field
		BindingValues       respjson.Field
		Name                respjson.Field
		Platform            respjson.Field
		URL                 respjson.Field
		UserReferenceErrors respjson.Field
		UserReferences      respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TweetDetailCard) RawJSON() string { return r.JSON.raw }
func (r *TweetDetailCard) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type TweetDetailCardUserReferenceError struct {
	Message string `json:"message"`
	Reason  string `json:"reason"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		Reason      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TweetDetailCardUserReferenceError) RawJSON() string { return r.JSON.raw }
func (r *TweetDetailCardUserReferenceError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// Community Note presentation metadata returned by X.
type TweetDetailCommunityNote struct {
	ID             string `json:"id"`
	DestinationURL string `json:"destinationUrl"`
	Footer         string `json:"footer"`
	FooterIconType string `json:"footerIconType"`
	IconType       string `json:"iconType"`
	// Public metadata whose fields are defined by X.
	Metadata    PublicMetadata `json:"metadata"`
	ShortTitle  string         `json:"shortTitle"`
	Subtitle    string         `json:"subtitle"`
	Title       string         `json:"title"`
	VisualStyle string         `json:"visualStyle"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		DestinationURL respjson.Field
		Footer         respjson.Field
		FooterIconType respjson.Field
		IconType       respjson.Field
		Metadata       respjson.Field
		ShortTitle     respjson.Field
		Subtitle       respjson.Field
		Title          respjson.Field
		VisualStyle    respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TweetDetailCommunityNote) RawJSON() string { return r.JSON.raw }
func (r *TweetDetailCommunityNote) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// Public reply policy and conversation owner.
type TweetDetailConversationControl struct {
	// Regions whose accounts may reply, as country codes. Sent with the
	// `country_of_origin` mode.
	AllowedCountryCodes []string `json:"allowedCountryCodes"`
	InviteViaMention    bool     `json:"inviteViaMention"`
	// Reply policy from X's newer field, such as `my_network`, `premium` or
	// `country_of_origin`. Omitted when X sends none.
	Mode string `json:"mode"`
	// The conversation owner's account ID.
	OwnerID       string `json:"ownerId"`
	OwnerUsername string `json:"ownerUsername"`
	Policy        string `json:"policy"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AllowedCountryCodes respjson.Field
		InviteViaMention    respjson.Field
		Mode                respjson.Field
		OwnerID             respjson.Field
		OwnerUsername       respjson.Field
		Policy              respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TweetDetailConversationControl) RawJSON() string { return r.JSON.raw }
func (r *TweetDetailConversationControl) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// Lists edit-chain identifiers and the remaining edit window.
type TweetDetailEdit struct {
	EditableUntilMsecs string   `json:"editableUntilMsecs"`
	EditTweetIDs       []string `json:"editTweetIds"`
	InitialTweetID     string   `json:"initialTweetId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EditableUntilMsecs respjson.Field
		EditTweetIDs       respjson.Field
		InitialTweetID     respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TweetDetailEdit) RawJSON() string { return r.JSON.raw }
func (r *TweetDetailEdit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// Lists hashtags, symbols, links, and mentions from tweet text.
type TweetDetailEntities struct {
	Hashtags     []TweetDetailEntitiesHashtag     `json:"hashtags"`
	Smarttags    []TweetDetailEntitiesSmarttag    `json:"smarttags"`
	Symbols      []TweetDetailEntitiesSymbol      `json:"symbols"`
	Timestamps   []TweetDetailEntitiesTimestamp   `json:"timestamps"`
	URLs         []TweetDetailEntitiesURL         `json:"urls"`
	UserMentions []TweetDetailEntitiesUserMention `json:"user_mentions"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Hashtags     respjson.Field
		Smarttags    respjson.Field
		Symbols      respjson.Field
		Timestamps   respjson.Field
		URLs         respjson.Field
		UserMentions respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TweetDetailEntities) RawJSON() string { return r.JSON.raw }
func (r *TweetDetailEntities) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// Provides hashtag text and source offsets within a tweet.
type TweetDetailEntitiesHashtag struct {
	Text    string  `json:"text" api:"required"`
	Indices []int64 `json:"indices"`
	// X's SmartTag for a cashtag, smarttag or hashtag. Xquik leaves out X's trading
	// platforms, which differ by reader.
	Tag TweetDetailEntitiesHashtagTag `json:"tag"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Text        respjson.Field
		Indices     respjson.Field
		Tag         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TweetDetailEntitiesHashtag) RawJSON() string { return r.JSON.raw }
func (r *TweetDetailEntitiesHashtag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// X's SmartTag for a cashtag, smarttag or hashtag. Xquik leaves out X's trading
// platforms, which differ by reader.
type TweetDetailEntitiesHashtagTag struct {
	// X's ID for the SmartTag.
	ID   string                            `json:"id"`
	Info TweetDetailEntitiesHashtagTagInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TweetDetailEntitiesHashtagTag) RawJSON() string { return r.JSON.raw }
func (r *TweetDetailEntitiesHashtagTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type TweetDetailEntitiesHashtagTagInfo struct {
	Info TweetDetailEntitiesHashtagTagInfoInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TweetDetailEntitiesHashtagTagInfo) RawJSON() string { return r.JSON.raw }
func (r *TweetDetailEntitiesHashtagTagInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type TweetDetailEntitiesHashtagTagInfoInfo struct {
	// Whether X marks the SmartTag active.
	Active bool   `json:"active"`
	Name   string `json:"name"`
	Ticker string `json:"ticker"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Active      respjson.Field
		Name        respjson.Field
		Ticker      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TweetDetailEntitiesHashtagTagInfoInfo) RawJSON() string { return r.JSON.raw }
func (r *TweetDetailEntitiesHashtagTagInfoInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// Indexed smart-tag, cashtag, or video timestamp metadata.
type TweetDetailEntitiesSmarttag struct {
	Indices []int64                        `json:"indices"`
	Seconds float64                        `json:"seconds"`
	Tag     TweetDetailEntitiesSmarttagTag `json:"tag"`
	Text    string                         `json:"text"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Indices     respjson.Field
		Seconds     respjson.Field
		Tag         respjson.Field
		Text        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TweetDetailEntitiesSmarttag) RawJSON() string { return r.JSON.raw }
func (r *TweetDetailEntitiesSmarttag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type TweetDetailEntitiesSmarttagTag struct {
	// X's ID for the SmartTag.
	ID   string                             `json:"id"`
	Info TweetDetailEntitiesSmarttagTagInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TweetDetailEntitiesSmarttagTag) RawJSON() string { return r.JSON.raw }
func (r *TweetDetailEntitiesSmarttagTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type TweetDetailEntitiesSmarttagTagInfo struct {
	Info TweetDetailEntitiesSmarttagTagInfoInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TweetDetailEntitiesSmarttagTagInfo) RawJSON() string { return r.JSON.raw }
func (r *TweetDetailEntitiesSmarttagTagInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type TweetDetailEntitiesSmarttagTagInfoInfo struct {
	// Whether X marks the SmartTag active.
	Active bool   `json:"active"`
	Name   string `json:"name"`
	Ticker string `json:"ticker"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Active      respjson.Field
		Name        respjson.Field
		Ticker      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TweetDetailEntitiesSmarttagTagInfoInfo) RawJSON() string { return r.JSON.raw }
func (r *TweetDetailEntitiesSmarttagTagInfoInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// Indexed smart-tag, cashtag, or video timestamp metadata.
type TweetDetailEntitiesSymbol struct {
	Indices []int64                      `json:"indices"`
	Seconds float64                      `json:"seconds"`
	Tag     TweetDetailEntitiesSymbolTag `json:"tag"`
	Text    string                       `json:"text"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Indices     respjson.Field
		Seconds     respjson.Field
		Tag         respjson.Field
		Text        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TweetDetailEntitiesSymbol) RawJSON() string { return r.JSON.raw }
func (r *TweetDetailEntitiesSymbol) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type TweetDetailEntitiesSymbolTag struct {
	// X's ID for the SmartTag.
	ID   string                           `json:"id"`
	Info TweetDetailEntitiesSymbolTagInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TweetDetailEntitiesSymbolTag) RawJSON() string { return r.JSON.raw }
func (r *TweetDetailEntitiesSymbolTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type TweetDetailEntitiesSymbolTagInfo struct {
	Info TweetDetailEntitiesSymbolTagInfoInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TweetDetailEntitiesSymbolTagInfo) RawJSON() string { return r.JSON.raw }
func (r *TweetDetailEntitiesSymbolTagInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type TweetDetailEntitiesSymbolTagInfoInfo struct {
	// Whether X marks the SmartTag active.
	Active bool   `json:"active"`
	Name   string `json:"name"`
	Ticker string `json:"ticker"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Active      respjson.Field
		Name        respjson.Field
		Ticker      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TweetDetailEntitiesSymbolTagInfoInfo) RawJSON() string { return r.JSON.raw }
func (r *TweetDetailEntitiesSymbolTagInfoInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// Indexed smart-tag, cashtag, or video timestamp metadata.
type TweetDetailEntitiesTimestamp struct {
	Indices []int64                         `json:"indices"`
	Seconds float64                         `json:"seconds"`
	Tag     TweetDetailEntitiesTimestampTag `json:"tag"`
	Text    string                          `json:"text"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Indices     respjson.Field
		Seconds     respjson.Field
		Tag         respjson.Field
		Text        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TweetDetailEntitiesTimestamp) RawJSON() string { return r.JSON.raw }
func (r *TweetDetailEntitiesTimestamp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type TweetDetailEntitiesTimestampTag struct {
	// X's ID for the SmartTag.
	ID   string                              `json:"id"`
	Info TweetDetailEntitiesTimestampTagInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TweetDetailEntitiesTimestampTag) RawJSON() string { return r.JSON.raw }
func (r *TweetDetailEntitiesTimestampTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type TweetDetailEntitiesTimestampTagInfo struct {
	Info TweetDetailEntitiesTimestampTagInfoInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TweetDetailEntitiesTimestampTagInfo) RawJSON() string { return r.JSON.raw }
func (r *TweetDetailEntitiesTimestampTagInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type TweetDetailEntitiesTimestampTagInfoInfo struct {
	// Whether X marks the SmartTag active.
	Active bool   `json:"active"`
	Name   string `json:"name"`
	Ticker string `json:"ticker"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Active      respjson.Field
		Name        respjson.Field
		Ticker      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TweetDetailEntitiesTimestampTagInfoInfo) RawJSON() string { return r.JSON.raw }
func (r *TweetDetailEntitiesTimestampTagInfoInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// Provides shortened, display, and expanded URLs from tweet text.
type TweetDetailEntitiesURL struct {
	DisplayURL  string  `json:"display_url"`
	ExpandedURL string  `json:"expanded_url"`
	Indices     []int64 `json:"indices"`
	URL         string  `json:"url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DisplayURL  respjson.Field
		ExpandedURL respjson.Field
		Indices     respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TweetDetailEntitiesURL) RawJSON() string { return r.JSON.raw }
func (r *TweetDetailEntitiesURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// Provides profile identity and source offsets for a mention.
type TweetDetailEntitiesUserMention struct {
	ScreenName string  `json:"screen_name" api:"required"`
	IDStr      string  `json:"id_str"`
	Indices    []int64 `json:"indices"`
	Name       string  `json:"name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ScreenName  respjson.Field
		IDStr       respjson.Field
		Indices     respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TweetDetailEntitiesUserMention) RawJSON() string { return r.JSON.raw }
func (r *TweetDetailEntitiesUserMention) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type TweetDetailLimitedAction struct {
	Action string                         `json:"action"`
	Prompt TweetDetailLimitedActionPrompt `json:"prompt"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Action      respjson.Field
		Prompt      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TweetDetailLimitedAction) RawJSON() string { return r.JSON.raw }
func (r *TweetDetailLimitedAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type TweetDetailLimitedActionPrompt struct {
	CtaType  string `json:"ctaType"`
	Headline string `json:"headline"`
	// Public metadata whose fields are defined by X.
	Metadata PublicMetadata `json:"metadata"`
	Subtext  string         `json:"subtext"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CtaType     respjson.Field
		Headline    respjson.Field
		Metadata    respjson.Field
		Subtext     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TweetDetailLimitedActionPrompt) RawJSON() string { return r.JSON.raw }
func (r *TweetDetailLimitedActionPrompt) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// Complete Note Tweet content and rich-text metadata.
type TweetDetailNoteTweet struct {
	Text string `json:"text" api:"required"`
	ID   string `json:"id"`
	// Lists hashtags, symbols, links, and mentions from tweet text.
	Entities TweetDetailNoteTweetEntities `json:"entities"`
	// Inline media positions in the Note Tweet text.
	InlineMedia  []TweetDetailNoteTweetInlineMedia `json:"inlineMedia"`
	IsExpandable bool                              `json:"isExpandable"`
	// The Note's own ID, apart from the post's.
	NoteID       string                            `json:"noteId"`
	RichtextTags []TweetDetailNoteTweetRichtextTag `json:"richtextTags"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Text         respjson.Field
		ID           respjson.Field
		Entities     respjson.Field
		InlineMedia  respjson.Field
		IsExpandable respjson.Field
		NoteID       respjson.Field
		RichtextTags respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TweetDetailNoteTweet) RawJSON() string { return r.JSON.raw }
func (r *TweetDetailNoteTweet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// Lists hashtags, symbols, links, and mentions from tweet text.
type TweetDetailNoteTweetEntities struct {
	Hashtags     []TweetDetailNoteTweetEntitiesHashtag     `json:"hashtags"`
	Smarttags    []TweetDetailNoteTweetEntitiesSmarttag    `json:"smarttags"`
	Symbols      []TweetDetailNoteTweetEntitiesSymbol      `json:"symbols"`
	Timestamps   []TweetDetailNoteTweetEntitiesTimestamp   `json:"timestamps"`
	URLs         []TweetDetailNoteTweetEntitiesURL         `json:"urls"`
	UserMentions []TweetDetailNoteTweetEntitiesUserMention `json:"user_mentions"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Hashtags     respjson.Field
		Smarttags    respjson.Field
		Symbols      respjson.Field
		Timestamps   respjson.Field
		URLs         respjson.Field
		UserMentions respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TweetDetailNoteTweetEntities) RawJSON() string { return r.JSON.raw }
func (r *TweetDetailNoteTweetEntities) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// Provides hashtag text and source offsets within a tweet.
type TweetDetailNoteTweetEntitiesHashtag struct {
	Text    string  `json:"text" api:"required"`
	Indices []int64 `json:"indices"`
	// X's SmartTag for a cashtag, smarttag or hashtag. Xquik leaves out X's trading
	// platforms, which differ by reader.
	Tag TweetDetailNoteTweetEntitiesHashtagTag `json:"tag"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Text        respjson.Field
		Indices     respjson.Field
		Tag         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TweetDetailNoteTweetEntitiesHashtag) RawJSON() string { return r.JSON.raw }
func (r *TweetDetailNoteTweetEntitiesHashtag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// X's SmartTag for a cashtag, smarttag or hashtag. Xquik leaves out X's trading
// platforms, which differ by reader.
type TweetDetailNoteTweetEntitiesHashtagTag struct {
	// X's ID for the SmartTag.
	ID   string                                     `json:"id"`
	Info TweetDetailNoteTweetEntitiesHashtagTagInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TweetDetailNoteTweetEntitiesHashtagTag) RawJSON() string { return r.JSON.raw }
func (r *TweetDetailNoteTweetEntitiesHashtagTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type TweetDetailNoteTweetEntitiesHashtagTagInfo struct {
	Info TweetDetailNoteTweetEntitiesHashtagTagInfoInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TweetDetailNoteTweetEntitiesHashtagTagInfo) RawJSON() string { return r.JSON.raw }
func (r *TweetDetailNoteTweetEntitiesHashtagTagInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type TweetDetailNoteTweetEntitiesHashtagTagInfoInfo struct {
	// Whether X marks the SmartTag active.
	Active bool   `json:"active"`
	Name   string `json:"name"`
	Ticker string `json:"ticker"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Active      respjson.Field
		Name        respjson.Field
		Ticker      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TweetDetailNoteTweetEntitiesHashtagTagInfoInfo) RawJSON() string { return r.JSON.raw }
func (r *TweetDetailNoteTweetEntitiesHashtagTagInfoInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// Indexed smart-tag, cashtag, or video timestamp metadata.
type TweetDetailNoteTweetEntitiesSmarttag struct {
	Indices []int64                                 `json:"indices"`
	Seconds float64                                 `json:"seconds"`
	Tag     TweetDetailNoteTweetEntitiesSmarttagTag `json:"tag"`
	Text    string                                  `json:"text"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Indices     respjson.Field
		Seconds     respjson.Field
		Tag         respjson.Field
		Text        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TweetDetailNoteTweetEntitiesSmarttag) RawJSON() string { return r.JSON.raw }
func (r *TweetDetailNoteTweetEntitiesSmarttag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type TweetDetailNoteTweetEntitiesSmarttagTag struct {
	// X's ID for the SmartTag.
	ID   string                                      `json:"id"`
	Info TweetDetailNoteTweetEntitiesSmarttagTagInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TweetDetailNoteTweetEntitiesSmarttagTag) RawJSON() string { return r.JSON.raw }
func (r *TweetDetailNoteTweetEntitiesSmarttagTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type TweetDetailNoteTweetEntitiesSmarttagTagInfo struct {
	Info TweetDetailNoteTweetEntitiesSmarttagTagInfoInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TweetDetailNoteTweetEntitiesSmarttagTagInfo) RawJSON() string { return r.JSON.raw }
func (r *TweetDetailNoteTweetEntitiesSmarttagTagInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type TweetDetailNoteTweetEntitiesSmarttagTagInfoInfo struct {
	// Whether X marks the SmartTag active.
	Active bool   `json:"active"`
	Name   string `json:"name"`
	Ticker string `json:"ticker"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Active      respjson.Field
		Name        respjson.Field
		Ticker      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TweetDetailNoteTweetEntitiesSmarttagTagInfoInfo) RawJSON() string { return r.JSON.raw }
func (r *TweetDetailNoteTweetEntitiesSmarttagTagInfoInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// Indexed smart-tag, cashtag, or video timestamp metadata.
type TweetDetailNoteTweetEntitiesSymbol struct {
	Indices []int64                               `json:"indices"`
	Seconds float64                               `json:"seconds"`
	Tag     TweetDetailNoteTweetEntitiesSymbolTag `json:"tag"`
	Text    string                                `json:"text"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Indices     respjson.Field
		Seconds     respjson.Field
		Tag         respjson.Field
		Text        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TweetDetailNoteTweetEntitiesSymbol) RawJSON() string { return r.JSON.raw }
func (r *TweetDetailNoteTweetEntitiesSymbol) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type TweetDetailNoteTweetEntitiesSymbolTag struct {
	// X's ID for the SmartTag.
	ID   string                                    `json:"id"`
	Info TweetDetailNoteTweetEntitiesSymbolTagInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TweetDetailNoteTweetEntitiesSymbolTag) RawJSON() string { return r.JSON.raw }
func (r *TweetDetailNoteTweetEntitiesSymbolTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type TweetDetailNoteTweetEntitiesSymbolTagInfo struct {
	Info TweetDetailNoteTweetEntitiesSymbolTagInfoInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TweetDetailNoteTweetEntitiesSymbolTagInfo) RawJSON() string { return r.JSON.raw }
func (r *TweetDetailNoteTweetEntitiesSymbolTagInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type TweetDetailNoteTweetEntitiesSymbolTagInfoInfo struct {
	// Whether X marks the SmartTag active.
	Active bool   `json:"active"`
	Name   string `json:"name"`
	Ticker string `json:"ticker"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Active      respjson.Field
		Name        respjson.Field
		Ticker      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TweetDetailNoteTweetEntitiesSymbolTagInfoInfo) RawJSON() string { return r.JSON.raw }
func (r *TweetDetailNoteTweetEntitiesSymbolTagInfoInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// Indexed smart-tag, cashtag, or video timestamp metadata.
type TweetDetailNoteTweetEntitiesTimestamp struct {
	Indices []int64                                  `json:"indices"`
	Seconds float64                                  `json:"seconds"`
	Tag     TweetDetailNoteTweetEntitiesTimestampTag `json:"tag"`
	Text    string                                   `json:"text"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Indices     respjson.Field
		Seconds     respjson.Field
		Tag         respjson.Field
		Text        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TweetDetailNoteTweetEntitiesTimestamp) RawJSON() string { return r.JSON.raw }
func (r *TweetDetailNoteTweetEntitiesTimestamp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type TweetDetailNoteTweetEntitiesTimestampTag struct {
	// X's ID for the SmartTag.
	ID   string                                       `json:"id"`
	Info TweetDetailNoteTweetEntitiesTimestampTagInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TweetDetailNoteTweetEntitiesTimestampTag) RawJSON() string { return r.JSON.raw }
func (r *TweetDetailNoteTweetEntitiesTimestampTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type TweetDetailNoteTweetEntitiesTimestampTagInfo struct {
	Info TweetDetailNoteTweetEntitiesTimestampTagInfoInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TweetDetailNoteTweetEntitiesTimestampTagInfo) RawJSON() string { return r.JSON.raw }
func (r *TweetDetailNoteTweetEntitiesTimestampTagInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type TweetDetailNoteTweetEntitiesTimestampTagInfoInfo struct {
	// Whether X marks the SmartTag active.
	Active bool   `json:"active"`
	Name   string `json:"name"`
	Ticker string `json:"ticker"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Active      respjson.Field
		Name        respjson.Field
		Ticker      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TweetDetailNoteTweetEntitiesTimestampTagInfoInfo) RawJSON() string { return r.JSON.raw }
func (r *TweetDetailNoteTweetEntitiesTimestampTagInfoInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// Provides shortened, display, and expanded URLs from tweet text.
type TweetDetailNoteTweetEntitiesURL struct {
	DisplayURL  string  `json:"display_url"`
	ExpandedURL string  `json:"expanded_url"`
	Indices     []int64 `json:"indices"`
	URL         string  `json:"url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DisplayURL  respjson.Field
		ExpandedURL respjson.Field
		Indices     respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TweetDetailNoteTweetEntitiesURL) RawJSON() string { return r.JSON.raw }
func (r *TweetDetailNoteTweetEntitiesURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// Provides profile identity and source offsets for a mention.
type TweetDetailNoteTweetEntitiesUserMention struct {
	ScreenName string  `json:"screen_name" api:"required"`
	IDStr      string  `json:"id_str"`
	Indices    []int64 `json:"indices"`
	Name       string  `json:"name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ScreenName  respjson.Field
		IDStr       respjson.Field
		Indices     respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TweetDetailNoteTweetEntitiesUserMention) RawJSON() string { return r.JSON.raw }
func (r *TweetDetailNoteTweetEntitiesUserMention) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type TweetDetailNoteTweetInlineMedia struct {
	Index   int64  `json:"index" api:"required"`
	MediaID string `json:"mediaId" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Index       respjson.Field
		MediaID     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TweetDetailNoteTweetInlineMedia) RawJSON() string { return r.JSON.raw }
func (r *TweetDetailNoteTweetInlineMedia) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type TweetDetailNoteTweetRichtextTag struct {
	FromIndex int64    `json:"fromIndex" api:"required"`
	ToIndex   int64    `json:"toIndex" api:"required"`
	Types     []string `json:"types" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FromIndex   respjson.Field
		ToIndex     respjson.Field
		Types       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TweetDetailNoteTweetRichtextTag) RawJSON() string { return r.JSON.raw }
func (r *TweetDetailNoteTweetRichtextTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// Describes public place metadata on a geotagged tweet.
type TweetDetailPlace struct {
	ID string `json:"id"`
	// Public metadata whose fields are defined by X.
	BoundingBox PublicMetadata `json:"boundingBox"`
	Country     string         `json:"country"`
	CountryCode string         `json:"countryCode"`
	FullName    string         `json:"fullName"`
	Name        string         `json:"name"`
	PlaceType   string         `json:"placeType"`
	URL         string         `json:"url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		BoundingBox respjson.Field
		Country     respjson.Field
		CountryCode respjson.Field
		FullName    respjson.Field
		Name        respjson.Field
		PlaceType   respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TweetDetailPlace) RawJSON() string { return r.JSON.raw }
func (r *TweetDetailPlace) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// Engagement counts retained from a prior tweet edit.
type TweetDetailPreviousCounts struct {
	BookmarkCount int64 `json:"bookmarkCount"`
	LikeCount     int64 `json:"likeCount"`
	QuoteCount    int64 `json:"quoteCount"`
	ReplyCount    int64 `json:"replyCount"`
	RetweetCount  int64 `json:"retweetCount"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BookmarkCount respjson.Field
		LikeCount     respjson.Field
		QuoteCount    respjson.Field
		ReplyCount    respjson.Field
		RetweetCount  respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TweetDetailPreviousCounts) RawJSON() string { return r.JSON.raw }
func (r *TweetDetailPreviousCounts) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// Public post and user referenced by this reaction.
type TweetDetailReactionContext struct {
	// Referenced post ID.
	ReactedToPostID string `json:"reactedToPostId"`
	// Public X profile.
	ReactedToUser shared.UserProfile `json:"reactedToUser"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ReactedToPostID respjson.Field
		ReactedToUser   respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TweetDetailReactionContext) RawJSON() string { return r.JSON.raw }
func (r *TweetDetailReactionContext) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// Sports game context attached to the post, when available.
type TweetDetailSportsContext struct {
	// Teams and their reported scores.
	Competitors []TweetDetailSportsContextCompetitor `json:"competitors"`
	// Game ID.
	GameID string `json:"gameId"`
	// League ID.
	League string `json:"league"`
	// Scheduled Unix time in milliseconds, preserving the source representation.
	ScheduledAtMs TweetDetailSportsContextScheduledAtMsUnion `json:"scheduledAtMs"`
	// Game state reported by the source.
	State string `json:"state"`
	// Display text for game progress.
	StatusText string `json:"statusText"`
	// Game title.
	Title string `json:"title"`
	// Game link.
	URL string `json:"url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Competitors   respjson.Field
		GameID        respjson.Field
		League        respjson.Field
		ScheduledAtMs respjson.Field
		State         respjson.Field
		StatusText    respjson.Field
		Title         respjson.Field
		URL           respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TweetDetailSportsContext) RawJSON() string { return r.JSON.raw }
func (r *TweetDetailSportsContext) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type TweetDetailSportsContextCompetitor struct {
	// Team abbreviation.
	Abbreviation string `json:"abbreviation"`
	// Team display color.
	Color string `json:"color"`
	// Team logo for dark backgrounds.
	DarkLogoURL string `json:"darkLogoUrl"`
	// Active competitor.
	IsActive bool `json:"isActive"`
	// Winning competitor.
	IsWinner bool `json:"isWinner"`
	// Team logo for light backgrounds.
	LogoURL string `json:"logoUrl"`
	// Team name.
	Name string `json:"name"`
	// Score display text.
	Score string `json:"score"`
	// Team ID.
	TeamID string `json:"teamId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Abbreviation respjson.Field
		Color        respjson.Field
		DarkLogoURL  respjson.Field
		IsActive     respjson.Field
		IsWinner     respjson.Field
		LogoURL      respjson.Field
		Name         respjson.Field
		Score        respjson.Field
		TeamID       respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TweetDetailSportsContextCompetitor) RawJSON() string { return r.JSON.raw }
func (r *TweetDetailSportsContextCompetitor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// TweetDetailSportsContextScheduledAtMsUnion contains all possible properties and
// values from [float64], [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfFloat OfString]
type TweetDetailSportsContextScheduledAtMsUnion struct {
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	JSON     struct {
		OfFloat  respjson.Field
		OfString respjson.Field
		raw      string
	} `json:"-"`
}

func (u TweetDetailSportsContextScheduledAtMsUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v, nil)
	return
}

func (u TweetDetailSportsContextScheduledAtMsUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v, nil)
	return
}

// Returns the unmodified JSON received from the API
func (u TweetDetailSportsContextScheduledAtMsUnion) RawJSON() string { return u.JSON.raw }

func (u *TweetDetailSportsContextScheduledAtMsUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u, &u.JSON.raw)
}

// Public visibility notice attached to an available tweet.
type TweetDetailTombstone struct {
	Text TweetDetailTombstoneText `json:"text"`
	// Visibility notice type.
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Text        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TweetDetailTombstone) RawJSON() string { return r.JSON.raw }
func (r *TweetDetailTombstone) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type TweetDetailTombstoneText struct {
	Entities []TweetDetailTombstoneTextEntity `json:"entities"`
	// Right-to-left text direction.
	Rtl bool `json:"rtl"`
	// Human-readable notice text.
	Text string `json:"text"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Entities    respjson.Field
		Rtl         respjson.Field
		Text        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TweetDetailTombstoneText) RawJSON() string { return r.JSON.raw }
func (r *TweetDetailTombstoneText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type TweetDetailTombstoneTextEntity struct {
	FromIndex int64                             `json:"fromIndex"`
	Ref       TweetDetailTombstoneTextEntityRef `json:"ref"`
	ToIndex   int64                             `json:"toIndex"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FromIndex   respjson.Field
		Ref         respjson.Field
		ToIndex     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TweetDetailTombstoneTextEntity) RawJSON() string { return r.JSON.raw }
func (r *TweetDetailTombstoneTextEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type TweetDetailTombstoneTextEntityRef struct {
	ScreenName string `json:"screenName"`
	Text       string `json:"text"`
	Type       string `json:"type"`
	URL        string `json:"url"`
	URLType    string `json:"urlType"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ScreenName  respjson.Field
		Text        respjson.Field
		Type        respjson.Field
		URL         respjson.Field
		URLType     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TweetDetailTombstoneTextEntityRef) RawJSON() string { return r.JSON.raw }
func (r *TweetDetailTombstoneTextEntityRef) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// Durable write record. Poll statusUrl until terminal is true. Reusing its
// Idempotency-Key returns this record. Create another action only when safeToRetry
// is true.
type XTweetNewResponse struct {
	ID string `json:"id" api:"required"`
	// Connected account selected for the write.
	Account XTweetNewResponseAccount `json:"account" api:"required"`
	// Any of "create_tweet", "delete_tweet", "like", "unlike", "retweet", "unretweet",
	// "follow", "unfollow", "remove_follower", "send_dm", "delete_dm", "upload_media",
	// "update_profile", "update_avatar", "update_banner", "create_community",
	// "delete_community", "join_community", "leave_community".
	Action XTweetNewResponseAction `json:"action" api:"required"`
	// plannedCredits is the approved maximum. chargedCredits comes from the settled
	// credit ledger. Pending or failed writes are not charged.
	Billing        XTweetNewResponseBilling `json:"billing" api:"required"`
	Charged        bool                     `json:"charged" api:"required"`
	ChargedCredits string                   `json:"chargedCredits" api:"required"`
	// Exact follow-up an API client or agent should perform.
	NextAction  XTweetNewResponseNextAction `json:"nextAction" api:"required"`
	Object      constant.XWriteAction       `json:"object" default:"x_write_action"`
	PollAfterMs int64                       `json:"pollAfterMs" api:"required"`
	// Stable fingerprint and sanitized payload for replay checks.
	Request XTweetNewResponseRequest `json:"request" api:"required"`
	// Confirmed result produced by the write, when available.
	Result XTweetNewResponseResult `json:"result" api:"required"`
	// True only when a new attempt can reasonably succeed.
	Retryable bool `json:"retryable" api:"required"`
	// True only when no write was dispatched and a new idempotency key may be used.
	SafeToRetry    bool `json:"safeToRetry" api:"required"`
	SendDispatched bool `json:"sendDispatched" api:"required"`
	// Any of "accepted", "dispatching", "pending_confirmation", "success", "failed",
	// "expired".
	Status    XTweetNewResponseStatus `json:"status" api:"required"`
	StatusURL string                  `json:"statusUrl" api:"required"`
	Success   bool                    `json:"success" api:"required"`
	// Existing X resource targeted by the write, when applicable.
	Target        XTweetNewResponseTarget `json:"target" api:"required"`
	TargetID      string                  `json:"targetId" api:"required"`
	Terminal      bool                    `json:"terminal" api:"required"`
	WriteActionID string                  `json:"writeActionId" api:"required"`
	// Compatibility field for a confirmed community ID.
	CommunityID string `json:"communityId"`
	// Confirmed community name when available.
	CommunityName         string    `json:"communityName"`
	CompletedAt           time.Time `json:"completedAt" format:"date-time"`
	ConfirmationAttempts  int64     `json:"confirmationAttempts"`
	ConfirmationCheckedAt time.Time `json:"confirmationCheckedAt" format:"date-time"`
	ConfirmedAt           time.Time `json:"confirmedAt" format:"date-time"`
	CreatedAt             time.Time `json:"createdAt" format:"date-time"`
	// Structured recovery context for a failed write.
	Details PublicMetadata `json:"details"`
	Error   string         `json:"error"`
	// Deadline for resolving a non-terminal write. This is not the Idempotency-Key
	// retention deadline.
	ExpiresAt  time.Time `json:"expiresAt" format:"date-time"`
	Idempotent bool      `json:"idempotent"`
	// Media count, kind, size, and billing details when used.
	Media PublicMetadata `json:"media"`
	// Compatibility field for a confirmed media upload ID.
	MediaID string `json:"mediaId"`
	// Public media URL when the upload creates one.
	MediaURL string `json:"mediaUrl" format:"uri"`
	Message  string `json:"message"`
	// Compatibility field for a confirmed direct message ID.
	MessageID   string `json:"messageId"`
	RequestHash string `json:"requestHash"`
	RequestID   string `json:"requestId"`
	// Compatibility result ID for other write actions.
	ResultID string `json:"resultId"`
	// Dispatch timestamp when the write reached execution.
	SendDispatchedAt time.Time `json:"sendDispatchedAt" format:"date-time"`
	// Compatibility field for a confirmed tweet result ID.
	TweetID   string    `json:"tweetId"`
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                    respjson.Field
		Account               respjson.Field
		Action                respjson.Field
		Billing               respjson.Field
		Charged               respjson.Field
		ChargedCredits        respjson.Field
		NextAction            respjson.Field
		Object                respjson.Field
		PollAfterMs           respjson.Field
		Request               respjson.Field
		Result                respjson.Field
		Retryable             respjson.Field
		SafeToRetry           respjson.Field
		SendDispatched        respjson.Field
		Status                respjson.Field
		StatusURL             respjson.Field
		Success               respjson.Field
		Target                respjson.Field
		TargetID              respjson.Field
		Terminal              respjson.Field
		WriteActionID         respjson.Field
		CommunityID           respjson.Field
		CommunityName         respjson.Field
		CompletedAt           respjson.Field
		ConfirmationAttempts  respjson.Field
		ConfirmationCheckedAt respjson.Field
		ConfirmedAt           respjson.Field
		CreatedAt             respjson.Field
		Details               respjson.Field
		Error                 respjson.Field
		ExpiresAt             respjson.Field
		Idempotent            respjson.Field
		Media                 respjson.Field
		MediaID               respjson.Field
		MediaURL              respjson.Field
		Message               respjson.Field
		MessageID             respjson.Field
		RequestHash           respjson.Field
		RequestID             respjson.Field
		ResultID              respjson.Field
		SendDispatchedAt      respjson.Field
		TweetID               respjson.Field
		UpdatedAt             respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XTweetNewResponse) RawJSON() string { return r.JSON.raw }
func (r *XTweetNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// Connected account selected for the write.
type XTweetNewResponseAccount struct {
	ID       string `json:"id" api:"required"`
	Username string `json:"username" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Username    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XTweetNewResponseAccount) RawJSON() string { return r.JSON.raw }
func (r *XTweetNewResponseAccount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type XTweetNewResponseAction string

const (
	XTweetNewResponseActionCreateTweet     XTweetNewResponseAction = "create_tweet"
	XTweetNewResponseActionDeleteTweet     XTweetNewResponseAction = "delete_tweet"
	XTweetNewResponseActionLike            XTweetNewResponseAction = "like"
	XTweetNewResponseActionUnlike          XTweetNewResponseAction = "unlike"
	XTweetNewResponseActionRetweet         XTweetNewResponseAction = "retweet"
	XTweetNewResponseActionUnretweet       XTweetNewResponseAction = "unretweet"
	XTweetNewResponseActionFollow          XTweetNewResponseAction = "follow"
	XTweetNewResponseActionUnfollow        XTweetNewResponseAction = "unfollow"
	XTweetNewResponseActionRemoveFollower  XTweetNewResponseAction = "remove_follower"
	XTweetNewResponseActionSendDm          XTweetNewResponseAction = "send_dm"
	XTweetNewResponseActionDeleteDm        XTweetNewResponseAction = "delete_dm"
	XTweetNewResponseActionUploadMedia     XTweetNewResponseAction = "upload_media"
	XTweetNewResponseActionUpdateProfile   XTweetNewResponseAction = "update_profile"
	XTweetNewResponseActionUpdateAvatar    XTweetNewResponseAction = "update_avatar"
	XTweetNewResponseActionUpdateBanner    XTweetNewResponseAction = "update_banner"
	XTweetNewResponseActionCreateCommunity XTweetNewResponseAction = "create_community"
	XTweetNewResponseActionDeleteCommunity XTweetNewResponseAction = "delete_community"
	XTweetNewResponseActionJoinCommunity   XTweetNewResponseAction = "join_community"
	XTweetNewResponseActionLeaveCommunity  XTweetNewResponseAction = "leave_community"
)

// plannedCredits is the approved maximum. chargedCredits comes from the settled
// credit ledger. Pending or failed writes are not charged.
type XTweetNewResponseBilling struct {
	Charged        bool   `json:"charged" api:"required"`
	ChargedCredits string `json:"chargedCredits" api:"required"`
	PlannedCredits string `json:"plannedCredits" api:"required"`
	// Any of "not_charged", "pending", "charged", "charge_failed", "refunded".
	Status string `json:"status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Charged        respjson.Field
		ChargedCredits respjson.Field
		PlannedCredits respjson.Field
		Status         respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XTweetNewResponseBilling) RawJSON() string { return r.JSON.raw }
func (r *XTweetNewResponseBilling) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// Exact follow-up an API client or agent should perform.
type XTweetNewResponseNextAction struct {
	// Any of "poll", "retry", "verify_result", "fix_request".
	Type                      string `json:"type" api:"required"`
	AfterMs                   int64  `json:"afterMs"`
	RequiresNewIdempotencyKey bool   `json:"requiresNewIdempotencyKey"`
	URL                       string `json:"url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type                      respjson.Field
		AfterMs                   respjson.Field
		RequiresNewIdempotencyKey respjson.Field
		URL                       respjson.Field
		ExtraFields               map[string]respjson.Field
		raw                       string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XTweetNewResponseNextAction) RawJSON() string { return r.JSON.raw }
func (r *XTweetNewResponseNextAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// Stable fingerprint and sanitized payload for replay checks.
type XTweetNewResponseRequest struct {
	// Stable hash of account, action, target, and payload.
	Hash string `json:"hash" api:"required"`
	// Exact sanitized payload dispatched for this action.
	Payload PublicMetadata `json:"payload" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Hash        respjson.Field
		Payload     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XTweetNewResponseRequest) RawJSON() string { return r.JSON.raw }
func (r *XTweetNewResponseRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// Confirmed result produced by the write, when available.
type XTweetNewResponseResult struct {
	ID    string `json:"id"`
	State string `json:"state"`
	// Any of "tweet", "direct_message", "media", "community", "state_change".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		State       respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XTweetNewResponseResult) RawJSON() string { return r.JSON.raw }
func (r *XTweetNewResponseResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type XTweetNewResponseStatus string

const (
	XTweetNewResponseStatusAccepted            XTweetNewResponseStatus = "accepted"
	XTweetNewResponseStatusDispatching         XTweetNewResponseStatus = "dispatching"
	XTweetNewResponseStatusPendingConfirmation XTweetNewResponseStatus = "pending_confirmation"
	XTweetNewResponseStatusSuccess             XTweetNewResponseStatus = "success"
	XTweetNewResponseStatusFailed              XTweetNewResponseStatus = "failed"
	XTweetNewResponseStatusExpired             XTweetNewResponseStatus = "expired"
)

// Existing X resource targeted by the write, when applicable.
type XTweetNewResponseTarget struct {
	ID string `json:"id" api:"required"`
	// Any of "tweet", "user", "community".
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XTweetNewResponseTarget) RawJSON() string { return r.JSON.raw }
func (r *XTweetNewResponseTarget) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type XTweetGetResponse struct {
	// Full tweet with text, engagement metrics, media, and metadata. A zero metric can
	// mean X did not report the count.
	Tweet TweetDetail `json:"tweet" api:"required"`
	// Tweet author profile. The lookup route always includes follower count and
	// verification state. Other profile fields appear when available.
	Author TweetAuthor `json:"author"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Tweet       respjson.Field
		Author      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XTweetGetResponse) RawJSON() string { return r.JSON.raw }
func (r *XTweetGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// Durable write record. Poll statusUrl until terminal is true. Reusing its
// Idempotency-Key returns this record. Create another action only when safeToRetry
// is true.
type XTweetDeleteResponse struct {
	ID string `json:"id" api:"required"`
	// Connected account selected for the write.
	Account XTweetDeleteResponseAccount `json:"account" api:"required"`
	// Any of "create_tweet", "delete_tweet", "like", "unlike", "retweet", "unretweet",
	// "follow", "unfollow", "remove_follower", "send_dm", "delete_dm", "upload_media",
	// "update_profile", "update_avatar", "update_banner", "create_community",
	// "delete_community", "join_community", "leave_community".
	Action XTweetDeleteResponseAction `json:"action" api:"required"`
	// plannedCredits is the approved maximum. chargedCredits comes from the settled
	// credit ledger. Pending or failed writes are not charged.
	Billing        XTweetDeleteResponseBilling `json:"billing" api:"required"`
	Charged        bool                        `json:"charged" api:"required"`
	ChargedCredits string                      `json:"chargedCredits" api:"required"`
	// Exact follow-up an API client or agent should perform.
	NextAction  XTweetDeleteResponseNextAction `json:"nextAction" api:"required"`
	Object      constant.XWriteAction          `json:"object" default:"x_write_action"`
	PollAfterMs int64                          `json:"pollAfterMs" api:"required"`
	// Stable fingerprint and sanitized payload for replay checks.
	Request XTweetDeleteResponseRequest `json:"request" api:"required"`
	// Confirmed result produced by the write, when available.
	Result XTweetDeleteResponseResult `json:"result" api:"required"`
	// True only when a new attempt can reasonably succeed.
	Retryable bool `json:"retryable" api:"required"`
	// True only when no write was dispatched and a new idempotency key may be used.
	SafeToRetry    bool `json:"safeToRetry" api:"required"`
	SendDispatched bool `json:"sendDispatched" api:"required"`
	// Any of "accepted", "dispatching", "pending_confirmation", "success", "failed",
	// "expired".
	Status    XTweetDeleteResponseStatus `json:"status" api:"required"`
	StatusURL string                     `json:"statusUrl" api:"required"`
	Success   bool                       `json:"success" api:"required"`
	// Existing X resource targeted by the write, when applicable.
	Target        XTweetDeleteResponseTarget `json:"target" api:"required"`
	TargetID      string                     `json:"targetId" api:"required"`
	Terminal      bool                       `json:"terminal" api:"required"`
	WriteActionID string                     `json:"writeActionId" api:"required"`
	// Compatibility field for a confirmed community ID.
	CommunityID string `json:"communityId"`
	// Confirmed community name when available.
	CommunityName         string    `json:"communityName"`
	CompletedAt           time.Time `json:"completedAt" format:"date-time"`
	ConfirmationAttempts  int64     `json:"confirmationAttempts"`
	ConfirmationCheckedAt time.Time `json:"confirmationCheckedAt" format:"date-time"`
	ConfirmedAt           time.Time `json:"confirmedAt" format:"date-time"`
	CreatedAt             time.Time `json:"createdAt" format:"date-time"`
	// Structured recovery context for a failed write.
	Details PublicMetadata `json:"details"`
	Error   string         `json:"error"`
	// Deadline for resolving a non-terminal write. This is not the Idempotency-Key
	// retention deadline.
	ExpiresAt  time.Time `json:"expiresAt" format:"date-time"`
	Idempotent bool      `json:"idempotent"`
	// Media count, kind, size, and billing details when used.
	Media PublicMetadata `json:"media"`
	// Compatibility field for a confirmed media upload ID.
	MediaID string `json:"mediaId"`
	// Public media URL when the upload creates one.
	MediaURL string `json:"mediaUrl" format:"uri"`
	Message  string `json:"message"`
	// Compatibility field for a confirmed direct message ID.
	MessageID   string `json:"messageId"`
	RequestHash string `json:"requestHash"`
	RequestID   string `json:"requestId"`
	// Compatibility result ID for other write actions.
	ResultID string `json:"resultId"`
	// Dispatch timestamp when the write reached execution.
	SendDispatchedAt time.Time `json:"sendDispatchedAt" format:"date-time"`
	// Compatibility field for a confirmed tweet result ID.
	TweetID   string    `json:"tweetId"`
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                    respjson.Field
		Account               respjson.Field
		Action                respjson.Field
		Billing               respjson.Field
		Charged               respjson.Field
		ChargedCredits        respjson.Field
		NextAction            respjson.Field
		Object                respjson.Field
		PollAfterMs           respjson.Field
		Request               respjson.Field
		Result                respjson.Field
		Retryable             respjson.Field
		SafeToRetry           respjson.Field
		SendDispatched        respjson.Field
		Status                respjson.Field
		StatusURL             respjson.Field
		Success               respjson.Field
		Target                respjson.Field
		TargetID              respjson.Field
		Terminal              respjson.Field
		WriteActionID         respjson.Field
		CommunityID           respjson.Field
		CommunityName         respjson.Field
		CompletedAt           respjson.Field
		ConfirmationAttempts  respjson.Field
		ConfirmationCheckedAt respjson.Field
		ConfirmedAt           respjson.Field
		CreatedAt             respjson.Field
		Details               respjson.Field
		Error                 respjson.Field
		ExpiresAt             respjson.Field
		Idempotent            respjson.Field
		Media                 respjson.Field
		MediaID               respjson.Field
		MediaURL              respjson.Field
		Message               respjson.Field
		MessageID             respjson.Field
		RequestHash           respjson.Field
		RequestID             respjson.Field
		ResultID              respjson.Field
		SendDispatchedAt      respjson.Field
		TweetID               respjson.Field
		UpdatedAt             respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XTweetDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *XTweetDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// Connected account selected for the write.
type XTweetDeleteResponseAccount struct {
	ID       string `json:"id" api:"required"`
	Username string `json:"username" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Username    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XTweetDeleteResponseAccount) RawJSON() string { return r.JSON.raw }
func (r *XTweetDeleteResponseAccount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type XTweetDeleteResponseAction string

const (
	XTweetDeleteResponseActionCreateTweet     XTweetDeleteResponseAction = "create_tweet"
	XTweetDeleteResponseActionDeleteTweet     XTweetDeleteResponseAction = "delete_tweet"
	XTweetDeleteResponseActionLike            XTweetDeleteResponseAction = "like"
	XTweetDeleteResponseActionUnlike          XTweetDeleteResponseAction = "unlike"
	XTweetDeleteResponseActionRetweet         XTweetDeleteResponseAction = "retweet"
	XTweetDeleteResponseActionUnretweet       XTweetDeleteResponseAction = "unretweet"
	XTweetDeleteResponseActionFollow          XTweetDeleteResponseAction = "follow"
	XTweetDeleteResponseActionUnfollow        XTweetDeleteResponseAction = "unfollow"
	XTweetDeleteResponseActionRemoveFollower  XTweetDeleteResponseAction = "remove_follower"
	XTweetDeleteResponseActionSendDm          XTweetDeleteResponseAction = "send_dm"
	XTweetDeleteResponseActionDeleteDm        XTweetDeleteResponseAction = "delete_dm"
	XTweetDeleteResponseActionUploadMedia     XTweetDeleteResponseAction = "upload_media"
	XTweetDeleteResponseActionUpdateProfile   XTweetDeleteResponseAction = "update_profile"
	XTweetDeleteResponseActionUpdateAvatar    XTweetDeleteResponseAction = "update_avatar"
	XTweetDeleteResponseActionUpdateBanner    XTweetDeleteResponseAction = "update_banner"
	XTweetDeleteResponseActionCreateCommunity XTweetDeleteResponseAction = "create_community"
	XTweetDeleteResponseActionDeleteCommunity XTweetDeleteResponseAction = "delete_community"
	XTweetDeleteResponseActionJoinCommunity   XTweetDeleteResponseAction = "join_community"
	XTweetDeleteResponseActionLeaveCommunity  XTweetDeleteResponseAction = "leave_community"
)

// plannedCredits is the approved maximum. chargedCredits comes from the settled
// credit ledger. Pending or failed writes are not charged.
type XTweetDeleteResponseBilling struct {
	Charged        bool   `json:"charged" api:"required"`
	ChargedCredits string `json:"chargedCredits" api:"required"`
	PlannedCredits string `json:"plannedCredits" api:"required"`
	// Any of "not_charged", "pending", "charged", "charge_failed", "refunded".
	Status string `json:"status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Charged        respjson.Field
		ChargedCredits respjson.Field
		PlannedCredits respjson.Field
		Status         respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XTweetDeleteResponseBilling) RawJSON() string { return r.JSON.raw }
func (r *XTweetDeleteResponseBilling) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// Exact follow-up an API client or agent should perform.
type XTweetDeleteResponseNextAction struct {
	// Any of "poll", "retry", "verify_result", "fix_request".
	Type                      string `json:"type" api:"required"`
	AfterMs                   int64  `json:"afterMs"`
	RequiresNewIdempotencyKey bool   `json:"requiresNewIdempotencyKey"`
	URL                       string `json:"url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type                      respjson.Field
		AfterMs                   respjson.Field
		RequiresNewIdempotencyKey respjson.Field
		URL                       respjson.Field
		ExtraFields               map[string]respjson.Field
		raw                       string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XTweetDeleteResponseNextAction) RawJSON() string { return r.JSON.raw }
func (r *XTweetDeleteResponseNextAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// Stable fingerprint and sanitized payload for replay checks.
type XTweetDeleteResponseRequest struct {
	// Stable hash of account, action, target, and payload.
	Hash string `json:"hash" api:"required"`
	// Exact sanitized payload dispatched for this action.
	Payload PublicMetadata `json:"payload" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Hash        respjson.Field
		Payload     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XTweetDeleteResponseRequest) RawJSON() string { return r.JSON.raw }
func (r *XTweetDeleteResponseRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// Confirmed result produced by the write, when available.
type XTweetDeleteResponseResult struct {
	ID    string `json:"id"`
	State string `json:"state"`
	// Any of "tweet", "direct_message", "media", "community", "state_change".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		State       respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XTweetDeleteResponseResult) RawJSON() string { return r.JSON.raw }
func (r *XTweetDeleteResponseResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type XTweetDeleteResponseStatus string

const (
	XTweetDeleteResponseStatusAccepted            XTweetDeleteResponseStatus = "accepted"
	XTweetDeleteResponseStatusDispatching         XTweetDeleteResponseStatus = "dispatching"
	XTweetDeleteResponseStatusPendingConfirmation XTweetDeleteResponseStatus = "pending_confirmation"
	XTweetDeleteResponseStatusSuccess             XTweetDeleteResponseStatus = "success"
	XTweetDeleteResponseStatusFailed              XTweetDeleteResponseStatus = "failed"
	XTweetDeleteResponseStatusExpired             XTweetDeleteResponseStatus = "expired"
)

// Existing X resource targeted by the write, when applicable.
type XTweetDeleteResponseTarget struct {
	ID string `json:"id" api:"required"`
	// Any of "tweet", "user", "community".
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XTweetDeleteResponseTarget) RawJSON() string { return r.JSON.raw }
func (r *XTweetDeleteResponseTarget) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// Direct reply rows. No-mode requests use resumable automatic coverage. Complete
// mode also returns nested replies and coverage diagnostics. Keep nested replies
// separate from direct coverage.
type XTweetGetRepliesResponse struct {
	// Evidence for direct-reply coverage and collector behavior.
	Diagnostic XTweetGetRepliesResponseDiagnostic `json:"diagnostic"`
	// Present on the last automatic page, and on a complete-mode answer without rows.
	// X counts this many more replies but shows them to nobody. They may be hidden,
	// protected or deleted. A retry returns the same replies.
	HiddenReplies int64 `json:"hiddenReplies"`
	// Present when your limit cut a complete-mode answer. The post has more replies.
	// Raise limit to get them all.
	//
	// Any of true.
	LimitReached bool `json:"limit_reached"`
	// Explains an answer cut by your limit.
	Message string `json:"message"`
	// Nested replies. Excluded from direct coverage.
	NestedReplies []shared.SearchTweet `json:"nested_replies"`
	// Present with latest or oldest sort. The read stopped before it held every reply
	// X shows. Rows stay sorted by date, but replies may be missing between them.
	// Retry later for a gap-free order. Complete mode and automatic pages set it
	// alike.
	//
	// Any of true.
	OrderPartial bool `json:"orderPartial"`
	// Present on an answer that lacks replies of a large thread posted before this
	// time. X shows only part of the replies older than 7 days. A retry returns about
	// the same replies.
	PartialBefore time.Time `json:"partialBefore" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Diagnostic    respjson.Field
		HiddenReplies respjson.Field
		LimitReached  respjson.Field
		Message       respjson.Field
		NestedReplies respjson.Field
		OrderPartial  respjson.Field
		PartialBefore respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
	shared.PaginatedTweets
}

// Returns the unmodified JSON received from the API
func (r XTweetGetRepliesResponse) RawJSON() string { return r.JSON.raw }
func (r *XTweetGetRepliesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// Evidence for direct-reply coverage and collector behavior.
type XTweetGetRepliesResponseDiagnostic struct {
	// Whether coverage met the target without truncation.
	Complete bool `json:"complete" api:"required"`
	// Unique direct replies as a percentage of the reported count.
	CoveragePercentage float64 `json:"coveragePercentage" api:"required"`
	// Cursor requests that failed.
	CursorFailures int64 `json:"cursorFailures" api:"required"`
	// Duplicate tweet IDs removed across pages and strategies.
	DuplicateCount int64 `json:"duplicateCount" api:"required"`
	// Empty pages rejected because they did not make progress.
	EmptyFalseProgressPages int64 `json:"emptyFalseProgressPages" api:"required"`
	// Malformed response items rejected.
	MalformedCount int64 `json:"malformedCount" api:"required"`
	// Expected response modules or fields missing from X.
	MissingResponseModulesOrFields []string `json:"missingResponseModulesOrFields" api:"required"`
	// Unique nested replies kept outside direct coverage.
	NestedReplyCount int64 `json:"nestedReplyCount" api:"required"`
	// Total pages attempted across all strategies.
	PagesAttempted int64 `json:"pagesAttempted" api:"required"`
	// Recommended next action when coverage is incomplete.
	RecommendedFallback string `json:"recommendedFallback" api:"required"`
	// Repeated cursors rejected to prevent loops.
	RepeatedCursorCount int64 `json:"repeatedCursorCount" api:"required"`
	// Reply count reported on the source post.
	ReportedReplyCount int64 `json:"reportedReplyCount" api:"required"`
	// Whether the requested row limit truncated safe results.
	ResponseTruncated bool `json:"responseTruncated" api:"required"`
	// Field-presence counts across the collected direct replies.
	Richness XTweetGetRepliesResponseDiagnosticRichness `json:"richness" api:"required"`
	// Every source ended, or no reply view shows a reply. A retry returns the same
	// rows.
	SourcesEnded bool `json:"sourcesEnded" api:"required"`
	// Per-strategy pagination and contribution evidence.
	StrategiesAttempted []XTweetGetRepliesResponseDiagnosticStrategiesAttempted `json:"strategiesAttempted" api:"required"`
	// Minimum direct replies required for the coverage target.
	TargetDirectReplies int64 `json:"targetDirectReplies" api:"required"`
	// Unique replies whose parent ID equals the source post ID.
	UniqueDirectReplies int64 `json:"uniqueDirectReplies" api:"required"`
	// Tweets rejected because they belonged elsewhere.
	UnrelatedCount int64 `json:"unrelatedCount" api:"required"`
	// Replies are accounted for through the requested scope & depth. Source errors
	// remain separate.
	ReplyScopeComplete bool `json:"replyScopeComplete"`
	// All child counts are known & match collected replies.
	ReplyTreeComplete bool `json:"replyTreeComplete"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Complete                       respjson.Field
		CoveragePercentage             respjson.Field
		CursorFailures                 respjson.Field
		DuplicateCount                 respjson.Field
		EmptyFalseProgressPages        respjson.Field
		MalformedCount                 respjson.Field
		MissingResponseModulesOrFields respjson.Field
		NestedReplyCount               respjson.Field
		PagesAttempted                 respjson.Field
		RecommendedFallback            respjson.Field
		RepeatedCursorCount            respjson.Field
		ReportedReplyCount             respjson.Field
		ResponseTruncated              respjson.Field
		Richness                       respjson.Field
		SourcesEnded                   respjson.Field
		StrategiesAttempted            respjson.Field
		TargetDirectReplies            respjson.Field
		UniqueDirectReplies            respjson.Field
		UnrelatedCount                 respjson.Field
		ReplyScopeComplete             respjson.Field
		ReplyTreeComplete              respjson.Field
		ExtraFields                    map[string]respjson.Field
		raw                            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XTweetGetRepliesResponseDiagnostic) RawJSON() string { return r.JSON.raw }
func (r *XTweetGetRepliesResponseDiagnostic) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// Field-presence counts across the collected direct replies.
type XTweetGetRepliesResponseDiagnosticRichness struct {
	// Replies with article content.
	Article int64 `json:"article" api:"required"`
	// Replies with author details.
	Author int64 `json:"author" api:"required"`
	// Replies with card metadata.
	Card int64 `json:"card" api:"required"`
	// Replies with community-note data.
	CommunityNote int64 `json:"communityNote" api:"required"`
	// Replies with a creation timestamp.
	CreatedAt int64 `json:"createdAt" api:"required"`
	// Replies with engagement counts.
	EngagementCounts int64 `json:"engagementCounts" api:"required"`
	// Replies with entity metadata.
	Entities int64 `json:"entities" api:"required"`
	// Replies with a language value.
	Language int64 `json:"language" api:"required"`
	// Replies with media metadata.
	Media int64 `json:"media" api:"required"`
	// Replies with quoted or reposted tweet data.
	QuotedOrRepostedTweet int64 `json:"quotedOrRepostedTweet" api:"required"`
	// Replies with text.
	Text int64 `json:"text" api:"required"`
	// Total unique direct replies evaluated for richness.
	TotalReplies int64 `json:"totalReplies" api:"required"`
	// Replies with a canonical URL.
	URL int64 `json:"url" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Article               respjson.Field
		Author                respjson.Field
		Card                  respjson.Field
		CommunityNote         respjson.Field
		CreatedAt             respjson.Field
		EngagementCounts      respjson.Field
		Entities              respjson.Field
		Language              respjson.Field
		Media                 respjson.Field
		QuotedOrRepostedTweet respjson.Field
		Text                  respjson.Field
		TotalReplies          respjson.Field
		URL                   respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XTweetGetRepliesResponseDiagnosticRichness) RawJSON() string { return r.JSON.raw }
func (r *XTweetGetRepliesResponseDiagnosticRichness) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// Result counts and stop reason for one reply strategy.
type XTweetGetRepliesResponseDiagnosticStrategiesAttempted struct {
	Name             string `json:"name" api:"required"`
	NewDirectReplies int64  `json:"newDirectReplies" api:"required"`
	NewNestedReplies int64  `json:"newNestedReplies" api:"required"`
	PagesAttempted   int64  `json:"pagesAttempted" api:"required"`
	// Any of "deadline", "empty_pages", "error", "missing_cursor", "no_next_page",
	// "page_cap", "repeated_cursor".
	StopReason string `json:"stopReason" api:"required"`
	// For a search of the post's conversation, the page that listed each new direct
	// reply, from 1.
	NewDirectReplyPages []int64 `json:"newDirectReplyPages"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name                respjson.Field
		NewDirectReplies    respjson.Field
		NewNestedReplies    respjson.Field
		PagesAttempted      respjson.Field
		StopReason          respjson.Field
		NewDirectReplyPages respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XTweetGetRepliesResponseDiagnosticStrategiesAttempted) RawJSON() string { return r.JSON.raw }
func (r *XTweetGetRepliesResponseDiagnosticStrategiesAttempted) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// XTweetSearchResponseUnion contains all possible properties and values from
// [shared.PaginatedTweets], [XTweetSearchResponseTweetSearchCoverageResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type XTweetSearchResponseUnion struct {
	HasNextPage bool   `json:"has_next_page"`
	NextCursor  string `json:"next_cursor"`
	// This field is from variant [shared.PaginatedTweets],
	// [XTweetSearchResponseTweetSearchCoverageResponse].
	Tweets []shared.SearchTweet `json:"tweets"`
	// This field is from variant [shared.PaginatedTweets],
	// [XTweetSearchResponseTweetSearchCoverageResponse].
	FilteredCount int64 `json:"filtered_count"`
	// This field is from variant [XTweetSearchResponseTweetSearchCoverageResponse].
	Diagnostic XTweetSearchResponseTweetSearchCoverageResponseDiagnostic `json:"diagnostic"`
	// This field is from variant [shared.PaginatedTweets],
	// [XTweetSearchResponseTweetSearchCoverageResponse].
	ReadBackTo time.Time `json:"readBackTo"`
	// This field is from variant [shared.PaginatedTweets],
	// [XTweetSearchResponseTweetSearchCoverageResponse].
	SearchContext shared.PaginatedTweetsSearchContext `json:"searchContext"`
	// This field is from variant [shared.PaginatedTweets],
	// [XTweetSearchResponseTweetSearchCoverageResponse].
	SearchedQuery string `json:"searchedQuery"`
	// This field is from variant [shared.PaginatedTweets],
	// [XTweetSearchResponseTweetSearchCoverageResponse].
	WindowComplete bool `json:"windowComplete"`
	JSON           struct {
		HasNextPage    respjson.Field
		NextCursor     respjson.Field
		Tweets         respjson.Field
		FilteredCount  respjson.Field
		Diagnostic     respjson.Field
		ReadBackTo     respjson.Field
		SearchContext  respjson.Field
		SearchedQuery  respjson.Field
		WindowComplete respjson.Field
		raw            string
	} `json:"-"`
}

func (u XTweetSearchResponseUnion) AsPaginatedTweets() (v shared.PaginatedTweets) {
	_ = v.UnmarshalJSON(json.RawMessage(u.JSON.raw))
	return
}

func (u XTweetSearchResponseUnion) AsXTweetSearchResponseTweetSearchCoverageResponse() (v XTweetSearchResponseTweetSearchCoverageResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v, &v.JSON.raw)
	return
}

// Returns the unmodified JSON received from the API
func (u XTweetSearchResponseUnion) RawJSON() string { return u.JSON.raw }

func (u *XTweetSearchResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u, &u.JSON.raw)
}

// Terminal Tweet search coverage response with diagnostics.
type XTweetSearchResponseTweetSearchCoverageResponse struct {
	// Coverage evidence across parallel search strategies.
	Diagnostic XTweetSearchResponseTweetSearchCoverageResponseDiagnostic `json:"diagnostic" api:"required"`
	// Any of false.
	HasNextPage bool `json:"has_next_page"`
	// Any of "".
	NextCursor string `json:"next_cursor"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Diagnostic  respjson.Field
		HasNextPage respjson.Field
		NextCursor  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.PaginatedTweets
}

// Returns the unmodified JSON received from the API
func (r XTweetSearchResponseTweetSearchCoverageResponse) RawJSON() string { return r.JSON.raw }
func (r *XTweetSearchResponseTweetSearchCoverageResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// Coverage evidence across parallel search strategies.
type XTweetSearchResponseTweetSearchCoverageResponseDiagnostic struct {
	// True after all active strategies exhaust their sources.
	Complete            bool  `json:"complete" api:"required"`
	CursorFailureCount  int64 `json:"cursorFailureCount" api:"required"`
	DeadlineReached     bool  `json:"deadlineReached" api:"required"`
	DuplicateCount      int64 `json:"duplicateCount" api:"required"`
	FailedStrategyCount int64 `json:"failedStrategyCount" api:"required"`
	MalformedCount      int64 `json:"malformedCount" api:"required"`
	PagesFetched        int64 `json:"pagesFetched" api:"required"`
	// Whether bounded time windows ran in parallel.
	Partitioned bool `json:"partitioned" api:"required"`
	// True when credits or the requested limit reduce output.
	ResponseTruncated    bool                                                                `json:"responseTruncated" api:"required"`
	ResultLimitReached   bool                                                                `json:"resultLimitReached" api:"required"`
	ReturnedTweets       int64                                                               `json:"returnedTweets" api:"required"`
	StalledStrategyCount int64                                                               `json:"stalledStrategyCount" api:"required"`
	Strategies           []XTweetSearchResponseTweetSearchCoverageResponseDiagnosticStrategy `json:"strategies" api:"required"`
	StrategyCount        int64                                                               `json:"strategyCount" api:"required"`
	UniqueTweets         int64                                                               `json:"uniqueTweets" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Complete             respjson.Field
		CursorFailureCount   respjson.Field
		DeadlineReached      respjson.Field
		DuplicateCount       respjson.Field
		FailedStrategyCount  respjson.Field
		MalformedCount       respjson.Field
		PagesFetched         respjson.Field
		Partitioned          respjson.Field
		ResponseTruncated    respjson.Field
		ResultLimitReached   respjson.Field
		ReturnedTweets       respjson.Field
		StalledStrategyCount respjson.Field
		Strategies           respjson.Field
		StrategyCount        respjson.Field
		UniqueTweets         respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XTweetSearchResponseTweetSearchCoverageResponseDiagnostic) RawJSON() string {
	return r.JSON.raw
}
func (r *XTweetSearchResponseTweetSearchCoverageResponseDiagnostic) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// Result counts and stop reason for one Tweet search strategy.
type XTweetSearchResponseTweetSearchCoverageResponseDiagnosticStrategy struct {
	DuplicateCount int64 `json:"duplicateCount" api:"required"`
	PagesFetched   int64 `json:"pagesFetched" api:"required"`
	// Any of "Latest", "Top".
	QueryType string `json:"queryType" api:"required"`
	// Reason a coverage strategy stopped.
	//
	// Any of "cursor_failure", "deadline", "exhausted", "failed", "page_limit",
	// "result_limit", "stalled".
	StopReason  string `json:"stopReason" api:"required"`
	Strategy    int64  `json:"strategy" api:"required"`
	UniqueAdded int64  `json:"uniqueAdded" api:"required"`
	// Non-overlapping time partition used by one strategy.
	Window XTweetSearchResponseTweetSearchCoverageResponseDiagnosticStrategyWindow `json:"window"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DuplicateCount respjson.Field
		PagesFetched   respjson.Field
		QueryType      respjson.Field
		StopReason     respjson.Field
		Strategy       respjson.Field
		UniqueAdded    respjson.Field
		Window         respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XTweetSearchResponseTweetSearchCoverageResponseDiagnosticStrategy) RawJSON() string {
	return r.JSON.raw
}
func (r *XTweetSearchResponseTweetSearchCoverageResponseDiagnosticStrategy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// Non-overlapping time partition used by one strategy.
type XTweetSearchResponseTweetSearchCoverageResponseDiagnosticStrategyWindow struct {
	SinceTime time.Time `json:"sinceTime" api:"required" format:"date-time"`
	UntilTime time.Time `json:"untilTime" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		SinceTime   respjson.Field
		UntilTime   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XTweetSearchResponseTweetSearchCoverageResponseDiagnosticStrategyWindow) RawJSON() string {
	return r.JSON.raw
}
func (r *XTweetSearchResponseTweetSearchCoverageResponseDiagnosticStrategyWindow) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type XTweetCheckRepostResponse struct {
	// How many of the account's newest posts the check read.
	CheckedPosts int64 `json:"checkedPosts" api:"required"`
	// Whether the check read the account's whole timeline.
	Complete bool `json:"complete" api:"required"`
	// Whether the account reposted the post.
	IsRetweeted bool `json:"isRetweeted" api:"required"`
	// When the account reposted it, else null.
	RetweetedAt time.Time `json:"retweetedAt" api:"required" format:"date-time"`
	// The ID of the account's repost, else null.
	RetweetID string `json:"retweetId" api:"required"`
	TweetID   string `json:"tweetId" api:"required"`
	UserID    string `json:"userId" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CheckedPosts respjson.Field
		Complete     respjson.Field
		IsRetweeted  respjson.Field
		RetweetedAt  respjson.Field
		RetweetID    respjson.Field
		TweetID      respjson.Field
		UserID       respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XTweetCheckRepostResponse) RawJSON() string { return r.JSON.raw }
func (r *XTweetCheckRepostResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type XTweetEmbedResponse struct {
	// Holds the HTML that embeds 1 tweet on a website.
	Embed XTweetEmbedResponseEmbed `json:"embed" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Embed       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XTweetEmbedResponse) RawJSON() string { return r.JSON.raw }
func (r *XTweetEmbedResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// Holds the HTML that embeds 1 tweet on a website.
type XTweetEmbedResponseEmbed struct {
	// Display name of the tweet's author.
	AuthorName string `json:"authorName" api:"required"`
	// Link to the author's profile.
	AuthorURL string `json:"authorUrl" api:"required"`
	// Markup to place on a page. It loads X's widget script unless `omitScript` is
	// `true`.
	HTML string `json:"html" api:"required"`
	// ID of the embedded tweet.
	TweetID string `json:"tweetId" api:"required"`
	// Link to the tweet with its author's username.
	URL string `json:"url" api:"required"`
	// Height in pixels, when X fixes it.
	Height int64 `json:"height"`
	// Width of the embed in pixels.
	Width int64 `json:"width"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AuthorName  respjson.Field
		AuthorURL   respjson.Field
		HTML        respjson.Field
		TweetID     respjson.Field
		URL         respjson.Field
		Height      respjson.Field
		Width       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XTweetEmbedResponseEmbed) RawJSON() string { return r.JSON.raw }
func (r *XTweetEmbedResponseEmbed) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type XTweetGetInsightsResponse struct {
	// Engagement insights of 1 tweet. Engagements are likes, replies, reposts and
	// quotes. A rate is null when nothing divides it, such as a tweet without views.
	Insights XTweetGetInsightsResponseInsights `json:"insights" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Insights    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XTweetGetInsightsResponse) RawJSON() string { return r.JSON.raw }
func (r *XTweetGetInsightsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// Engagement insights of 1 tweet. Engagements are likes, replies, reposts and
// quotes. A rate is null when nothing divides it, such as a tweet without views.
type XTweetGetInsightsResponseInsights struct {
	// Hours since the tweet went up.
	AgeHours float64 `json:"ageHours" api:"required"`
	// Reposts and quotes per like.
	Amplification float64 `json:"amplification" api:"required"`
	// ID of the author.
	AuthorID string `json:"authorId" api:"required"`
	// Username of the author.
	AuthorUserName string `json:"authorUserName" api:"required"`
	// Replies per like.
	Conversation float64 `json:"conversation" api:"required"`
	// When the tweet went up.
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Engagements as a percent of views.
	EngagementPerView float64 `json:"engagementPerView" api:"required"`
	// Engagements as a percent of the author's followers.
	EngagementRate float64 `json:"engagementRate" api:"required"`
	// Likes, replies, reposts and quotes.
	Engagements float64 `json:"engagements" api:"required"`
	// Engagements per hour since the tweet went up.
	EngagementsPerHour float64 `json:"engagementsPerHour" api:"required"`
	// The public counts of 1 tweet.
	Metrics XTweetGetInsightsResponseInsightsMetrics `json:"metrics" api:"required"`
	// ID of the tweet the insights describe.
	TweetID string `json:"tweetId" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AgeHours           respjson.Field
		Amplification      respjson.Field
		AuthorID           respjson.Field
		AuthorUserName     respjson.Field
		Conversation       respjson.Field
		CreatedAt          respjson.Field
		EngagementPerView  respjson.Field
		EngagementRate     respjson.Field
		Engagements        respjson.Field
		EngagementsPerHour respjson.Field
		Metrics            respjson.Field
		TweetID            respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XTweetGetInsightsResponseInsights) RawJSON() string { return r.JSON.raw }
func (r *XTweetGetInsightsResponseInsights) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// The public counts of 1 tweet.
type XTweetGetInsightsResponseInsightsMetrics struct {
	// Bookmarks.
	Bookmarks float64 `json:"bookmarks" api:"required"`
	// Likes.
	Likes float64 `json:"likes" api:"required"`
	// Quotes.
	Quotes float64 `json:"quotes" api:"required"`
	// Replies.
	Replies float64 `json:"replies" api:"required"`
	// Reposts.
	Reposts float64 `json:"reposts" api:"required"`
	// Views.
	Views float64 `json:"views" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Bookmarks   respjson.Field
		Likes       respjson.Field
		Quotes      respjson.Field
		Replies     respjson.Field
		Reposts     respjson.Field
		Views       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XTweetGetInsightsResponseInsightsMetrics) RawJSON() string { return r.JSON.raw }
func (r *XTweetGetInsightsResponseInsightsMetrics) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// A page of user IDs without profiles. Follow next_cursor while the response
// reports more pages.
type XTweetGetRetweeterIDsResponse struct {
	HasNextPage bool `json:"has_next_page" api:"required"`
	// X user IDs, newest relationship first.
	IDs        []string `json:"ids" api:"required"`
	NextCursor string   `json:"next_cursor" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		HasNextPage respjson.Field
		IDs         respjson.Field
		NextCursor  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XTweetGetRetweeterIDsResponse) RawJSON() string { return r.JSON.raw }
func (r *XTweetGetRetweeterIDsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type XTweetGetSubtitlesResponse struct {
	// Holds the caption tracks of 1 tweet's videos.
	Subtitles XTweetGetSubtitlesResponseSubtitles `json:"subtitles" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Subtitles   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XTweetGetSubtitlesResponse) RawJSON() string { return r.JSON.raw }
func (r *XTweetGetSubtitlesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// Holds the caption tracks of 1 tweet's videos.
type XTweetGetSubtitlesResponseSubtitles struct {
	// Caption tracks, empty when no video has captions.
	Tracks []XTweetGetSubtitlesResponseSubtitlesTrack `json:"tracks" api:"required"`
	// ID of the tweet.
	TweetID string `json:"tweetId" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Tracks      respjson.Field
		TweetID     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XTweetGetSubtitlesResponseSubtitles) RawJSON() string { return r.JSON.raw }
func (r *XTweetGetSubtitlesResponseSubtitles) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// Holds 1 caption track of a video.
type XTweetGetSubtitlesResponseSubtitlesTrack struct {
	// True when X transcribed the audio itself.
	AutoGenerated bool `json:"autoGenerated" api:"required"`
	// Timed captions in order.
	Cues []XTweetGetSubtitlesResponseSubtitlesTrackCue `json:"cues" api:"required"`
	// Every cue's text in order, as 1 line.
	Text string `json:"text" api:"required"`
	// True when the track holds more than 50 caption files. The cues then cover the
	// first 50.
	Truncated bool `json:"truncated" api:"required"`
	// Link to the track's first WebVTT file.
	URL string `json:"url" api:"required"`
	// Language code of the track, in lower case.
	Language string `json:"language"`
	// Media key of the video, as media rows carry it.
	MediaKey string `json:"mediaKey"`
	// Track name the uploader or X gave it.
	Name string `json:"name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AutoGenerated respjson.Field
		Cues          respjson.Field
		Text          respjson.Field
		Truncated     respjson.Field
		URL           respjson.Field
		Language      respjson.Field
		MediaKey      respjson.Field
		Name          respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XTweetGetSubtitlesResponseSubtitlesTrack) RawJSON() string { return r.JSON.raw }
func (r *XTweetGetSubtitlesResponseSubtitlesTrack) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type XTweetGetSubtitlesResponseSubtitlesTrackCue struct {
	// End time in milliseconds.
	EndMs int64 `json:"endMs" api:"required"`
	// Start time in milliseconds.
	StartMs int64 `json:"startMs" api:"required"`
	// Caption text, with a line break between lines.
	Text string `json:"text" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EndMs       respjson.Field
		StartMs     respjson.Field
		Text        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XTweetGetSubtitlesResponseSubtitlesTrackCue) RawJSON() string { return r.JSON.raw }
func (r *XTweetGetSubtitlesResponseSubtitlesTrackCue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type XTweetResolveLinksResponse struct {
	Links []XTweetResolveLinksResponseLink `json:"links" api:"required"`
	// Links not read, in the order sent. X did not answer for them in time, or
	// available credits ran out. Send them again.
	UnprocessedURLs []string `json:"unprocessedUrls" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Links           respjson.Field
		UnprocessedURLs respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XTweetResolveLinksResponse) RawJSON() string { return r.JSON.raw }
func (r *XTweetResolveLinksResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// Pairs 1 t.co link with the URL it leads to.
type XTweetResolveLinksResponseLink struct {
	// The t.co link, as `https://t.co/<code>`.
	URL string `json:"url" api:"required"`
	// Target URL, as the link's author set it. Xquik does not check it. Absent for a
	// link t.co does not have.
	ResolvedURL string `json:"resolvedUrl"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		URL         respjson.Field
		ResolvedURL respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XTweetResolveLinksResponseLink) RawJSON() string { return r.JSON.raw }
func (r *XTweetResolveLinksResponseLink) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type XTweetTranslateResponse struct {
	// Holds X's translation of 1 tweet into the requested language. `entities` holds
	// the links, mentions, and tags with their offsets in the translated text, and is
	// empty when the text has none.
	Translation XTweetTranslateResponseTranslation `json:"translation" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Translation respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XTweetTranslateResponse) RawJSON() string { return r.JSON.raw }
func (r *XTweetTranslateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// Holds X's translation of 1 tweet into the requested language. `entities` holds
// the links, mentions, and tags with their offsets in the translated text, and is
// empty when the text has none.
type XTweetTranslateResponseTranslation struct {
	// Lists hashtags, symbols, links, and mentions from tweet text.
	Entities XTweetTranslateResponseTranslationEntities `json:"entities" api:"required"`
	// Language code of the translated text, in lower case.
	Language string `json:"language" api:"required"`
	// Translated tweet text.
	Text string `json:"text" api:"required"`
	// ID of the translated tweet.
	TweetID string `json:"tweetId" api:"required"`
	// Language code X detects in the tweet, when it names one.
	SourceLanguage string `json:"sourceLanguage"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Entities       respjson.Field
		Language       respjson.Field
		Text           respjson.Field
		TweetID        respjson.Field
		SourceLanguage respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XTweetTranslateResponseTranslation) RawJSON() string { return r.JSON.raw }
func (r *XTweetTranslateResponseTranslation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// Lists hashtags, symbols, links, and mentions from tweet text.
type XTweetTranslateResponseTranslationEntities struct {
	Hashtags     []XTweetTranslateResponseTranslationEntitiesHashtag     `json:"hashtags"`
	Smarttags    []XTweetTranslateResponseTranslationEntitiesSmarttag    `json:"smarttags"`
	Symbols      []XTweetTranslateResponseTranslationEntitiesSymbol      `json:"symbols"`
	Timestamps   []XTweetTranslateResponseTranslationEntitiesTimestamp   `json:"timestamps"`
	URLs         []XTweetTranslateResponseTranslationEntitiesURL         `json:"urls"`
	UserMentions []XTweetTranslateResponseTranslationEntitiesUserMention `json:"user_mentions"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Hashtags     respjson.Field
		Smarttags    respjson.Field
		Symbols      respjson.Field
		Timestamps   respjson.Field
		URLs         respjson.Field
		UserMentions respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XTweetTranslateResponseTranslationEntities) RawJSON() string { return r.JSON.raw }
func (r *XTweetTranslateResponseTranslationEntities) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// Hashtag text and its source offsets in a tweet.
type XTweetTranslateResponseTranslationEntitiesHashtag struct {
	Text    string  `json:"text" api:"required"`
	Indices []int64 `json:"indices"`
	// X's SmartTag for a cashtag, smarttag or hashtag. Xquik leaves out X's trading
	// platforms, which differ by reader.
	Tag XTweetTranslateResponseTranslationEntitiesHashtagTag `json:"tag"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Text        respjson.Field
		Indices     respjson.Field
		Tag         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XTweetTranslateResponseTranslationEntitiesHashtag) RawJSON() string { return r.JSON.raw }
func (r *XTweetTranslateResponseTranslationEntitiesHashtag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// X's SmartTag for a cashtag, smarttag or hashtag. Xquik leaves out X's trading
// platforms, which differ by reader.
type XTweetTranslateResponseTranslationEntitiesHashtagTag struct {
	// X's ID for the SmartTag.
	ID   string                                                   `json:"id"`
	Info XTweetTranslateResponseTranslationEntitiesHashtagTagInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XTweetTranslateResponseTranslationEntitiesHashtagTag) RawJSON() string { return r.JSON.raw }
func (r *XTweetTranslateResponseTranslationEntitiesHashtagTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type XTweetTranslateResponseTranslationEntitiesHashtagTagInfo struct {
	Info XTweetTranslateResponseTranslationEntitiesHashtagTagInfoInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XTweetTranslateResponseTranslationEntitiesHashtagTagInfo) RawJSON() string { return r.JSON.raw }
func (r *XTweetTranslateResponseTranslationEntitiesHashtagTagInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type XTweetTranslateResponseTranslationEntitiesHashtagTagInfoInfo struct {
	// Whether X marks the SmartTag active.
	Active bool   `json:"active"`
	Name   string `json:"name"`
	Ticker string `json:"ticker"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Active      respjson.Field
		Name        respjson.Field
		Ticker      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XTweetTranslateResponseTranslationEntitiesHashtagTagInfoInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *XTweetTranslateResponseTranslationEntitiesHashtagTagInfoInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// Indexed smart-tag, cashtag, or video timestamp metadata.
type XTweetTranslateResponseTranslationEntitiesSmarttag struct {
	Indices []int64 `json:"indices"`
	Seconds float64 `json:"seconds"`
	// X's SmartTag for a cashtag, smarttag or hashtag. Xquik leaves out X's trading
	// platforms, which differ by reader.
	Tag  XTweetTranslateResponseTranslationEntitiesSmarttagTag `json:"tag"`
	Text string                                                `json:"text"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Indices     respjson.Field
		Seconds     respjson.Field
		Tag         respjson.Field
		Text        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XTweetTranslateResponseTranslationEntitiesSmarttag) RawJSON() string { return r.JSON.raw }
func (r *XTweetTranslateResponseTranslationEntitiesSmarttag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// X's SmartTag for a cashtag, smarttag or hashtag. Xquik leaves out X's trading
// platforms, which differ by reader.
type XTweetTranslateResponseTranslationEntitiesSmarttagTag struct {
	// X's ID for the SmartTag.
	ID   string                                                    `json:"id"`
	Info XTweetTranslateResponseTranslationEntitiesSmarttagTagInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XTweetTranslateResponseTranslationEntitiesSmarttagTag) RawJSON() string { return r.JSON.raw }
func (r *XTweetTranslateResponseTranslationEntitiesSmarttagTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type XTweetTranslateResponseTranslationEntitiesSmarttagTagInfo struct {
	Info XTweetTranslateResponseTranslationEntitiesSmarttagTagInfoInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XTweetTranslateResponseTranslationEntitiesSmarttagTagInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *XTweetTranslateResponseTranslationEntitiesSmarttagTagInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type XTweetTranslateResponseTranslationEntitiesSmarttagTagInfoInfo struct {
	// Whether X marks the SmartTag active.
	Active bool   `json:"active"`
	Name   string `json:"name"`
	Ticker string `json:"ticker"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Active      respjson.Field
		Name        respjson.Field
		Ticker      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XTweetTranslateResponseTranslationEntitiesSmarttagTagInfoInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *XTweetTranslateResponseTranslationEntitiesSmarttagTagInfoInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// Indexed smart-tag, cashtag, or video timestamp metadata.
type XTweetTranslateResponseTranslationEntitiesSymbol struct {
	Indices []int64 `json:"indices"`
	Seconds float64 `json:"seconds"`
	// X's SmartTag for a cashtag, smarttag or hashtag. Xquik leaves out X's trading
	// platforms, which differ by reader.
	Tag  XTweetTranslateResponseTranslationEntitiesSymbolTag `json:"tag"`
	Text string                                              `json:"text"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Indices     respjson.Field
		Seconds     respjson.Field
		Tag         respjson.Field
		Text        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XTweetTranslateResponseTranslationEntitiesSymbol) RawJSON() string { return r.JSON.raw }
func (r *XTweetTranslateResponseTranslationEntitiesSymbol) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// X's SmartTag for a cashtag, smarttag or hashtag. Xquik leaves out X's trading
// platforms, which differ by reader.
type XTweetTranslateResponseTranslationEntitiesSymbolTag struct {
	// X's ID for the SmartTag.
	ID   string                                                  `json:"id"`
	Info XTweetTranslateResponseTranslationEntitiesSymbolTagInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XTweetTranslateResponseTranslationEntitiesSymbolTag) RawJSON() string { return r.JSON.raw }
func (r *XTweetTranslateResponseTranslationEntitiesSymbolTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type XTweetTranslateResponseTranslationEntitiesSymbolTagInfo struct {
	Info XTweetTranslateResponseTranslationEntitiesSymbolTagInfoInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XTweetTranslateResponseTranslationEntitiesSymbolTagInfo) RawJSON() string { return r.JSON.raw }
func (r *XTweetTranslateResponseTranslationEntitiesSymbolTagInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type XTweetTranslateResponseTranslationEntitiesSymbolTagInfoInfo struct {
	// Whether X marks the SmartTag active.
	Active bool   `json:"active"`
	Name   string `json:"name"`
	Ticker string `json:"ticker"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Active      respjson.Field
		Name        respjson.Field
		Ticker      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XTweetTranslateResponseTranslationEntitiesSymbolTagInfoInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *XTweetTranslateResponseTranslationEntitiesSymbolTagInfoInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// Indexed smart-tag, cashtag, or video timestamp metadata.
type XTweetTranslateResponseTranslationEntitiesTimestamp struct {
	Indices []int64 `json:"indices"`
	Seconds float64 `json:"seconds"`
	// X's SmartTag for a cashtag, smarttag or hashtag. Xquik leaves out X's trading
	// platforms, which differ by reader.
	Tag  XTweetTranslateResponseTranslationEntitiesTimestampTag `json:"tag"`
	Text string                                                 `json:"text"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Indices     respjson.Field
		Seconds     respjson.Field
		Tag         respjson.Field
		Text        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XTweetTranslateResponseTranslationEntitiesTimestamp) RawJSON() string { return r.JSON.raw }
func (r *XTweetTranslateResponseTranslationEntitiesTimestamp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// X's SmartTag for a cashtag, smarttag or hashtag. Xquik leaves out X's trading
// platforms, which differ by reader.
type XTweetTranslateResponseTranslationEntitiesTimestampTag struct {
	// X's ID for the SmartTag.
	ID   string                                                     `json:"id"`
	Info XTweetTranslateResponseTranslationEntitiesTimestampTagInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XTweetTranslateResponseTranslationEntitiesTimestampTag) RawJSON() string { return r.JSON.raw }
func (r *XTweetTranslateResponseTranslationEntitiesTimestampTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type XTweetTranslateResponseTranslationEntitiesTimestampTagInfo struct {
	Info XTweetTranslateResponseTranslationEntitiesTimestampTagInfoInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XTweetTranslateResponseTranslationEntitiesTimestampTagInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *XTweetTranslateResponseTranslationEntitiesTimestampTagInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type XTweetTranslateResponseTranslationEntitiesTimestampTagInfoInfo struct {
	// Whether X marks the SmartTag active.
	Active bool   `json:"active"`
	Name   string `json:"name"`
	Ticker string `json:"ticker"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Active      respjson.Field
		Name        respjson.Field
		Ticker      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XTweetTranslateResponseTranslationEntitiesTimestampTagInfoInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *XTweetTranslateResponseTranslationEntitiesTimestampTagInfoInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// Shortened, display, and expanded URLs from tweet text.
type XTweetTranslateResponseTranslationEntitiesURL struct {
	DisplayURL  string  `json:"display_url"`
	ExpandedURL string  `json:"expanded_url"`
	Indices     []int64 `json:"indices"`
	URL         string  `json:"url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DisplayURL  respjson.Field
		ExpandedURL respjson.Field
		Indices     respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XTweetTranslateResponseTranslationEntitiesURL) RawJSON() string { return r.JSON.raw }
func (r *XTweetTranslateResponseTranslationEntitiesURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// Profile identity and source offsets of a mention.
type XTweetTranslateResponseTranslationEntitiesUserMention struct {
	ScreenName string  `json:"screen_name" api:"required"`
	IDStr      string  `json:"id_str"`
	Indices    []int64 `json:"indices"`
	Name       string  `json:"name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ScreenName  respjson.Field
		IDStr       respjson.Field
		Indices     respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XTweetTranslateResponseTranslationEntitiesUserMention) RawJSON() string { return r.JSON.raw }
func (r *XTweetTranslateResponseTranslationEntitiesUserMention) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type XTweetNewParams struct {
	// X account (@username or account ID)
	Account        string            `json:"account" api:"required"`
	IdempotencyKey string            `header:"Idempotency-Key" api:"required" json:"-"`
	CommunityID    param.Opt[string] `json:"community_id,omitzero"`
	IsNoteTweet    param.Opt[bool]   `json:"is_note_tweet,omitzero"`
	ReplyToTweetID param.Opt[string] `json:"reply_to_tweet_id,omitzero"`
	// Tweet text (optional when media is provided)
	Text param.Opt[string] `json:"text,omitzero"`
	// Array of public media URLs to attach. Supports up to 4 images or exactly 1 MP4
	// video up to 100 MB. Each URL must be publicly reachable. Attached media adds 2
	// credits per started MB across all files.
	Media []string `json:"media,omitzero"`
	paramObj
}

func (r XTweetNewParams) MarshalJSON() (data []byte, err error) {
	type shadow XTweetNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *XTweetNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, nil)
}

type XTweetListParams struct {
	// Comma-separated tweet IDs (max 100)
	IDs string `query:"ids" api:"required" json:"-"`
	// Fields to keep in each row, comma-separated. Use dots for nested fields, such as
	// author.username. A list keeps the field in each item. Case & underscores do not
	// matter. A row that lacks a field leaves it out. Cursors, counts & diagnostics
	// always stay. An unknown name answers 400 unknown_field with the valid names,
	// before any read. The price stays the same. SDK row types then hold only the
	// named fields.
	Fields param.Opt[string] `query:"fields,omitzero" json:"-"`
	// Fields to drop from each row, comma-separated. It takes the same names as
	// fields. With fields, it drops from the kept fields.
	Omit param.Opt[string] `query:"omit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [XTweetListParams]'s query parameters as `url.Values`.
func (r XTweetListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type XTweetDeleteParams struct {
	// X account identifier (@username or account ID)
	Account        string `json:"account" api:"required"`
	IdempotencyKey string `header:"Idempotency-Key" api:"required" json:"-"`
	paramObj
}

func (r XTweetDeleteParams) MarshalJSON() (data []byte, err error) {
	type shadow XTweetDeleteParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *XTweetDeleteParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, nil)
}

type XTweetGetFavoritersParams struct {
	// Match any comma-separated or line-separated bio term, ignoring case.
	BioContains param.Opt[string] `query:"bioContains,omitzero" json:"-"`
	// Only return profiles X marks as blue verified. X marks many business &
	// government accounts blue too. Use verifiedType for 1 badge type.
	BlueVerifiedOnly param.Opt[bool] `query:"blueVerifiedOnly,omitzero" json:"-"`
	// Only return profiles that publish an email outside free webmail, such as Gmail.
	BusinessEmailOnly param.Opt[bool] `query:"businessEmailOnly,omitzero" json:"-"`
	// Pagination cursor for favoriters
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Only return profiles with an email at one of these domains. Separate domains
	// with commas. Subdomains match.
	EmailDomains param.Opt[string] `query:"emailDomains,omitzero" json:"-"`
	// Fields to keep in each row, comma-separated. Use dots for nested fields, such as
	// author.username. A list keeps the field in each item. Case & underscores do not
	// matter. A row that lacks a field leaves it out. Cursors, counts & diagnostics
	// always stay. An unknown name answers 400 unknown_field with the valid names,
	// before any read. The price stays the same. SDK row types then hold only the
	// named fields.
	Fields param.Opt[string] `query:"fields,omitzero" json:"-"`
	// Only return profiles that publish an email.
	HasEmail param.Opt[bool] `query:"hasEmail,omitzero" json:"-"`
	// Only return profiles with a location.
	HasLocation param.Opt[bool] `query:"hasLocation,omitzero" json:"-"`
	// Only return profiles with a website.
	HasWebsite param.Opt[bool] `query:"hasWebsite,omitzero" json:"-"`
	// Match a location substring, ignoring case.
	LocationContains param.Opt[string] `query:"locationContains,omitzero" json:"-"`
	// Maximum follower count. Missing counts pass this maximum.
	MaxFollowers param.Opt[int64] `query:"maxFollowers,omitzero" json:"-"`
	// Profiles may follow at most this many accounts.
	MaxFollowing param.Opt[int64] `query:"maxFollowing,omitzero" json:"-"`
	// Maximum post count. maxPosts is also accepted.
	MaxStatuses param.Opt[int64] `query:"maxStatuses,omitzero" json:"-"`
	// Minimum account age in whole days.
	MinAccountAgeDays param.Opt[int64] `query:"minAccountAgeDays,omitzero" json:"-"`
	// Minimum follower count. Filtering happens before billing.
	MinFollowers param.Opt[int64] `query:"minFollowers,omitzero" json:"-"`
	// Profiles must follow at least this many accounts.
	MinFollowing param.Opt[int64] `query:"minFollowing,omitzero" json:"-"`
	// Minimum post count. minPosts is also accepted.
	MinStatuses param.Opt[int64] `query:"minStatuses,omitzero" json:"-"`
	// Fields to drop from each row, comma-separated. It takes the same names as
	// fields. With fields, it drops from the kept fields.
	Omit param.Opt[string] `query:"omit,omitzero" json:"-"`
	// Maximum user profiles requested from this page (1-200, default 200). Source,
	// filters, or credits can return fewer profiles. Follow next_cursor while the
	// response reports more pages. Deprecated aliases remain accepted.
	PageSize param.Opt[int64] `query:"pageSize,omitzero" json:"-"`
	// Match a username substring, ignoring case.
	UsernameContains param.Opt[string] `query:"usernameContains,omitzero" json:"-"`
	// Only return verified profiles.
	VerifiedOnly param.Opt[bool] `query:"verifiedOnly,omitzero" json:"-"`
	// Match the verification type exactly, ignoring case.
	VerifiedType param.Opt[string] `query:"verifiedType,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [XTweetGetFavoritersParams]'s query parameters as
// `url.Values`.
func (r XTweetGetFavoritersParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type XTweetGetQuotesParams struct {
	// Raw advanced search query appended as-is.
	AdvancedQuery param.Opt[string] `query:"advancedQuery,omitzero" json:"-"`
	// Words or quoted phrases where any one can match. Separate with spaces, commas,
	// or lines.
	AnyWords param.Opt[string] `query:"anyWords,omitzero" json:"-"`
	// Only return tweets from Blue-verified authors.
	BlueVerifiedOnly param.Opt[bool] `query:"blueVerifiedOnly,omitzero" json:"-"`
	// Geo bounding box, such as -74.1 40.6 -73.9 40.8.
	BoundingBox param.Opt[string] `query:"boundingBox,omitzero" json:"-"`
	// Match the Tweet card name.
	CardName param.Opt[string] `query:"cardName,omitzero" json:"-"`
	// Cashtags separated by spaces, commas, or lines.
	Cashtags param.Opt[string] `query:"cashtags,omitzero" json:"-"`
	// Conversation ID filter.
	ConversationID param.Opt[string] `query:"conversationId,omitzero" json:"-"`
	// Cursor from the previous response. Xquik cursors resume automatic coverage.
	// Existing unprefixed cursors keep legacy standard behavior.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Match this literal phrase, including any hyphens.
	ExactPhrase param.Opt[string] `query:"exactPhrase,omitzero" json:"-"`
	// Exclude a source application.
	ExcludeSource param.Opt[string] `query:"excludeSource,omitzero" json:"-"`
	// Words or quoted phrases to exclude. Separate with spaces, commas, or lines.
	ExcludeWords param.Opt[string] `query:"excludeWords,omitzero" json:"-"`
	// Fields to keep in each row, comma-separated. Use dots for nested fields, such as
	// author.username. A list keeps the field in each item. Case & underscores do not
	// matter. A row that lacks a field leaves it out. Cursors, counts & diagnostics
	// always stay. An unknown name answers 400 unknown_field with the valid names,
	// before any read. The price stays the same. SDK row types then hold only the
	// named fields.
	Fields param.Opt[string] `query:"fields,omitzero" json:"-"`
	// Filter by author username.
	FromUser param.Opt[string] `query:"fromUser,omitzero" json:"-"`
	// Match latitude, longitude, and radius.
	Geocode param.Opt[string] `query:"geocode,omitzero" json:"-"`
	// Hashtags separated by spaces, commas, or lines.
	Hashtags param.Opt[string] `query:"hashtags,omitzero" json:"-"`
	// Include reply tweets unless replies specifies another mode.
	IncludeReplies param.Opt[bool] `query:"includeReplies,omitzero" json:"-"`
	// Only replies to this tweet ID.
	InReplyToTweetID param.Opt[string] `query:"inReplyToTweetId,omitzero" json:"-"`
	// Words the Tweets must match, in X search syntax.
	Keywords param.Opt[string] `query:"keywords,omitzero" json:"-"`
	// Filter by language. Alias `lang` is accepted.
	Language param.Opt[string] `query:"language,omitzero" json:"-"`
	// Search within a list ID.
	ListID param.Opt[string] `query:"listId,omitzero" json:"-"`
	// Maximum likes threshold. maxLikes is also accepted.
	MaxFaves param.Opt[int64] `query:"maxFaves,omitzero" json:"-"`
	// Return Tweets older than this Tweet ID.
	MaxID param.Opt[string] `query:"maxId,omitzero" json:"-"`
	// Maximum quotes threshold.
	MaxQuotes param.Opt[int64] `query:"maxQuotes,omitzero" json:"-"`
	// Maximum replies threshold.
	MaxReplies param.Opt[int64] `query:"maxReplies,omitzero" json:"-"`
	// Maximum retweets threshold.
	MaxRetweets param.Opt[int64] `query:"maxRetweets,omitzero" json:"-"`
	// Filter tweets mentioning a username.
	Mentioning param.Opt[string] `query:"mentioning,omitzero" json:"-"`
	// Minimum bookmark count threshold.
	MinBookmarks param.Opt[int64] `query:"minBookmarks,omitzero" json:"-"`
	// Deprecated: Use MinLikes. MinFaves sends the minFaves query key, which the
	// API reads the same way as minLikes.
	MinFaves param.Opt[int64] `query:"minFaves,omitzero" json:"-"`
	// Minimum likes. Aliases: minFaves, min_likes, min_faves.
	MinLikes param.Opt[int64] `query:"minLikes,omitzero" json:"-"`
	// Minimum quote count threshold.
	MinQuotes param.Opt[int64] `query:"minQuotes,omitzero" json:"-"`
	// Minimum replies threshold.
	MinReplies param.Opt[int64] `query:"minReplies,omitzero" json:"-"`
	// Minimum retweets threshold.
	MinRetweets param.Opt[int64] `query:"minRetweets,omitzero" json:"-"`
	// Minimum view count threshold.
	MinViews param.Opt[int64] `query:"minViews,omitzero" json:"-"`
	// Only return native reposts.
	NativeRetweets param.Opt[bool] `query:"nativeRetweets,omitzero" json:"-"`
	// Match a place name.
	Near param.Opt[string] `query:"near,omitzero" json:"-"`
	// Only return news results.
	News param.Opt[bool] `query:"news,omitzero" json:"-"`
	// Fields to drop from each row, comma-separated. It takes the same names as
	// fields. With fields, it drops from the kept fields.
	Omit param.Opt[string] `query:"omit,omitzero" json:"-"`
	// Automatic pages accept 1-300 Tweets. Standard pages keep 1-100. Default 20.
	// Follow next_cursor while the response reports more pages. Deprecated aliases
	// remain accepted.
	PageSize param.Opt[int64] `query:"pageSize,omitzero" json:"-"`
	// Search within a place ID.
	Place param.Opt[string] `query:"place,omitzero" json:"-"`
	// Search within a country code.
	PlaceCountry param.Opt[string] `query:"placeCountry,omitzero" json:"-"`
	// Geo point radius, such as -73.99 40.73 25mi.
	PointRadius param.Opt[string] `query:"pointRadius,omitzero" json:"-"`
	// Only quotes of this tweet ID.
	QuotesOfTweetID param.Opt[string] `query:"quotesOfTweetId,omitzero" json:"-"`
	// Only retweets of this tweet ID.
	RetweetsOfTweetID param.Opt[string] `query:"retweetsOfTweetId,omitzero" json:"-"`
	// Enable the safe-search filter.
	Safe param.Opt[bool] `query:"safe,omitzero" json:"-"`
	// Start date in YYYY-MM-DD format.
	SinceDate param.Opt[time.Time] `query:"sinceDate,omitzero" format:"date" json:"-"`
	// Return Tweets newer than this Tweet ID.
	SinceID param.Opt[string] `query:"sinceId,omitzero" json:"-"`
	// Inclusive ISO bound for Tweet creation time.
	SinceTime param.Opt[string] `query:"sinceTime,omitzero" json:"-"`
	// Match the source application.
	Source param.Opt[string] `query:"source,omitzero" json:"-"`
	// Filter replies sent to a username.
	ToUser param.Opt[string] `query:"toUser,omitzero" json:"-"`
	// End date in YYYY-MM-DD format.
	UntilDate param.Opt[time.Time] `query:"untilDate,omitzero" format:"date" json:"-"`
	// Exclusive ISO bound for Tweet creation time.
	UntilTime param.Opt[string] `query:"untilTime,omitzero" json:"-"`
	// URL substring or domain filter.
	URL param.Opt[string] `query:"url,omitzero" json:"-"`
	// Only return tweets from verified authors.
	VerifiedOnly param.Opt[bool] `query:"verifiedOnly,omitzero" json:"-"`
	// Set the radius for the near filter.
	Within param.Opt[string] `query:"within,omitzero" json:"-"`
	// Match Tweets inside a recent time window.
	WithinTime param.Opt[string] `query:"withinTime,omitzero" json:"-"`
	// Filter media. Aliases: has_video, has_media.
	//
	// Any of "images", "videos", "gifs", "media", "links", "none".
	MediaType XTweetGetQuotesParamsMediaType `query:"mediaType,omitzero" json:"-"`
	// Optional legacy pagination override.
	//
	// Any of "standard".
	Mode XTweetGetQuotesParamsMode `query:"mode,omitzero" json:"-"`
	// Latest for newest first, or Top for the quotes X ranks highest. Top is X's
	// ranking, not a count order, & a later page can hold a quote X ranks higher. Any
	// case works. Alias sort, which also takes Recency & Relevance. Send the same
	// value with each cursor. Standard pages list Latest.
	//
	// Any of "Latest", "Top".
	QueryType XTweetGetQuotesParamsQueryType `query:"queryType,omitzero" json:"-"`
	// Only when the caller requests a quote mode.
	//
	// Any of "include", "exclude", "only".
	Quotes XTweetGetQuotesParamsQuotes `query:"quotes,omitzero" json:"-"`
	// Only when the caller requests a reply mode.
	//
	// Any of "include", "exclude", "only".
	Replies XTweetGetQuotesParamsReplies `query:"replies,omitzero" json:"-"`
	// Only when the caller requests a repost mode.
	//
	// Any of "include", "exclude", "only".
	Retweets XTweetGetQuotesParamsRetweets `query:"retweets,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [XTweetGetQuotesParams]'s query parameters as `url.Values`.
func (r XTweetGetQuotesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter media. Aliases: has_video, has_media.
type XTweetGetQuotesParamsMediaType string

const (
	XTweetGetQuotesParamsMediaTypeImages XTweetGetQuotesParamsMediaType = "images"
	XTweetGetQuotesParamsMediaTypeVideos XTweetGetQuotesParamsMediaType = "videos"
	XTweetGetQuotesParamsMediaTypeGifs   XTweetGetQuotesParamsMediaType = "gifs"
	XTweetGetQuotesParamsMediaTypeMedia  XTweetGetQuotesParamsMediaType = "media"
	XTweetGetQuotesParamsMediaTypeLinks  XTweetGetQuotesParamsMediaType = "links"
	XTweetGetQuotesParamsMediaTypeNone   XTweetGetQuotesParamsMediaType = "none"
)

// Optional legacy pagination override.
type XTweetGetQuotesParamsMode string

const (
	XTweetGetQuotesParamsModeStandard XTweetGetQuotesParamsMode = "standard"
)

// Only when the caller requests a quote mode.
type XTweetGetQuotesParamsQuotes string

const (
	XTweetGetQuotesParamsQuotesInclude XTweetGetQuotesParamsQuotes = "include"
	XTweetGetQuotesParamsQuotesExclude XTweetGetQuotesParamsQuotes = "exclude"
	XTweetGetQuotesParamsQuotesOnly    XTweetGetQuotesParamsQuotes = "only"
)

// Only when the caller requests a reply mode.
type XTweetGetQuotesParamsReplies string

const (
	XTweetGetQuotesParamsRepliesInclude XTweetGetQuotesParamsReplies = "include"
	XTweetGetQuotesParamsRepliesExclude XTweetGetQuotesParamsReplies = "exclude"
	XTweetGetQuotesParamsRepliesOnly    XTweetGetQuotesParamsReplies = "only"
)

// Only when the caller requests a repost mode.
type XTweetGetQuotesParamsRetweets string

const (
	XTweetGetQuotesParamsRetweetsInclude XTweetGetQuotesParamsRetweets = "include"
	XTweetGetQuotesParamsRetweetsExclude XTweetGetQuotesParamsRetweets = "exclude"
	XTweetGetQuotesParamsRetweetsOnly    XTweetGetQuotesParamsRetweets = "only"
)

type XTweetGetRepliesParams struct {
	// Legacy cursor alias. Prefer cursor.
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// Words or quoted phrases where any one can match. Separate with spaces, commas,
	// or lines.
	AnyWords param.Opt[string] `query:"anyWords,omitzero" json:"-"`
	// Only return tweets from Blue-verified authors.
	BlueVerifiedOnly param.Opt[bool] `query:"blueVerifiedOnly,omitzero" json:"-"`
	// Match the Tweet card name.
	CardName param.Opt[string] `query:"cardName,omitzero" json:"-"`
	// Cashtags separated by spaces, commas, or lines.
	Cashtags param.Opt[string] `query:"cashtags,omitzero" json:"-"`
	// Conversation ID filter.
	ConversationID param.Opt[string] `query:"conversationId,omitzero" json:"-"`
	// Cursor from the previous response. Xquik cursors resume automatic coverage.
	// Existing unprefixed cursors keep legacy standard behavior.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Match this literal phrase, including any hyphens.
	ExactPhrase param.Opt[string] `query:"exactPhrase,omitzero" json:"-"`
	// Exclude replies written by the source-post author.
	ExcludeOriginalAuthor param.Opt[bool] `query:"excludeOriginalAuthor,omitzero" json:"-"`
	// Exclude a source application.
	ExcludeSource param.Opt[string] `query:"excludeSource,omitzero" json:"-"`
	// Words or quoted phrases to exclude. Separate with spaces, commas, or lines.
	ExcludeWords param.Opt[string] `query:"excludeWords,omitzero" json:"-"`
	// Fields to keep in each row, comma-separated. Use dots for nested fields, such as
	// author.username. A list keeps the field in each item. Case & underscores do not
	// matter. A row that lacks a field leaves it out. Cursors, counts & diagnostics
	// always stay. An unknown name answers 400 unknown_field with the valid names,
	// before any read. The price stays the same. SDK row types then hold only the
	// named fields.
	Fields param.Opt[string] `query:"fields,omitzero" json:"-"`
	// Filter by author username.
	FromUser param.Opt[string] `query:"fromUser,omitzero" json:"-"`
	// Match latitude, longitude, and radius.
	Geocode param.Opt[string] `query:"geocode,omitzero" json:"-"`
	// Hashtags separated by spaces, commas, or lines.
	Hashtags param.Opt[string] `query:"hashtags,omitzero" json:"-"`
	// Only return replies containing media.
	HasMediaOnly param.Opt[bool] `query:"hasMediaOnly,omitzero" json:"-"`
	// Complete mode only. Adds the replies the post's author hid, each with
	// hiddenByAuthor true. They count toward limit like any reply. A paged read
	// answers 400. Get its hidden replies from GET /x/tweets/{id}/hidden-replies.
	IncludeHidden param.Opt[bool] `query:"includeHidden,omitzero" json:"-"`
	// Include the source post and count it toward limit.
	IncludeOriginalPost param.Opt[bool] `query:"includeOriginalPost,omitzero" json:"-"`
	// Only replies to this tweet ID.
	InReplyToTweetID param.Opt[string] `query:"inReplyToTweetId,omitzero" json:"-"`
	// Filter by language. Alias `lang` is accepted.
	Language param.Opt[string] `query:"language,omitzero" json:"-"`
	// Complete mode defaults to 25,000 combined direct and nested replies. Set a
	// smaller or larger total with limit. Automatic pages accept 1-300. Standard pages
	// accept 1-100.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Maximum reply depth from the source post.
	MaxDepth param.Opt[int64] `query:"maxDepth,omitzero" json:"-"`
	// Maximum likes threshold. maxLikes is also accepted.
	MaxFaves param.Opt[int64] `query:"maxFaves,omitzero" json:"-"`
	// Return Tweets older than this Tweet ID.
	MaxID param.Opt[string] `query:"maxId,omitzero" json:"-"`
	// Maximum quotes threshold.
	MaxQuotes param.Opt[int64] `query:"maxQuotes,omitzero" json:"-"`
	// Maximum replies threshold.
	MaxReplies param.Opt[int64] `query:"maxReplies,omitzero" json:"-"`
	// Maximum retweets threshold.
	MaxRetweets param.Opt[int64] `query:"maxRetweets,omitzero" json:"-"`
	// Filter tweets mentioning a username.
	Mentioning param.Opt[string] `query:"mentioning,omitzero" json:"-"`
	// Minimum bookmark count threshold.
	MinBookmarks param.Opt[int64] `query:"minBookmarks,omitzero" json:"-"`
	// Deprecated: Use MinLikes. MinFaves sends the minFaves query key, which the
	// API reads the same way as minLikes.
	MinFaves param.Opt[int64] `query:"minFaves,omitzero" json:"-"`
	// Minimum likes. Aliases: minFaves, min_likes, min_faves.
	MinLikes param.Opt[int64] `query:"minLikes,omitzero" json:"-"`
	// Minimum quote count threshold.
	MinQuotes param.Opt[int64] `query:"minQuotes,omitzero" json:"-"`
	// Minimum replies threshold.
	MinReplies param.Opt[int64] `query:"minReplies,omitzero" json:"-"`
	// Minimum retweets threshold.
	MinRetweets param.Opt[int64] `query:"minRetweets,omitzero" json:"-"`
	// Minimum view count threshold.
	MinViews param.Opt[int64] `query:"minViews,omitzero" json:"-"`
	// Only return native reposts.
	NativeRetweets param.Opt[bool] `query:"nativeRetweets,omitzero" json:"-"`
	// Match a place name.
	Near param.Opt[string] `query:"near,omitzero" json:"-"`
	// Only return news results.
	News param.Opt[bool] `query:"news,omitzero" json:"-"`
	// Fields to drop from each row, comma-separated. It takes the same names as
	// fields. With fields, it drops from the kept fields.
	Omit param.Opt[string] `query:"omit,omitzero" json:"-"`
	// Automatic pages accept 1-300 Tweets. Standard pages keep 1-100. Default 20.
	// Follow next_cursor while the response reports more pages. Deprecated aliases
	// remain accepted.
	PageSize param.Opt[int64] `query:"pageSize,omitzero" json:"-"`
	// Only quotes of this tweet ID.
	QuotesOfTweetID param.Opt[string] `query:"quotesOfTweetId,omitzero" json:"-"`
	// Only retweets of this tweet ID.
	RetweetsOfTweetID param.Opt[string] `query:"retweetsOfTweetId,omitzero" json:"-"`
	// Enable the safe-search filter.
	Safe param.Opt[bool] `query:"safe,omitzero" json:"-"`
	// Start date in YYYY-MM-DD format.
	SinceDate param.Opt[time.Time] `query:"sinceDate,omitzero" format:"date" json:"-"`
	// Return Tweets newer than this Tweet ID.
	SinceID param.Opt[string] `query:"sinceId,omitzero" json:"-"`
	// Inclusive ISO bound for Tweet creation time.
	SinceTime param.Opt[string] `query:"sinceTime,omitzero" json:"-"`
	// Match the source application.
	Source param.Opt[string] `query:"source,omitzero" json:"-"`
	// Filter replies sent to a username.
	ToUser param.Opt[string] `query:"toUser,omitzero" json:"-"`
	// End date in YYYY-MM-DD format.
	UntilDate param.Opt[time.Time] `query:"untilDate,omitzero" format:"date" json:"-"`
	// Exclusive ISO bound for Tweet creation time.
	UntilTime param.Opt[string] `query:"untilTime,omitzero" json:"-"`
	// URL substring or domain filter.
	URL param.Opt[string] `query:"url,omitzero" json:"-"`
	// Only return tweets from verified authors.
	VerifiedOnly param.Opt[bool] `query:"verifiedOnly,omitzero" json:"-"`
	// Set the radius for the near filter.
	Within param.Opt[string] `query:"within,omitzero" json:"-"`
	// Match Tweets inside a recent time window.
	WithinTime param.Opt[string] `query:"withinTime,omitzero" json:"-"`
	// Filter media. Aliases: has_video, has_media.
	//
	// Any of "images", "videos", "gifs", "media", "links", "none".
	MediaType XTweetGetRepliesParamsMediaType `query:"mediaType,omitzero" json:"-"`
	// Override automatic coverage. Standard uses legacy pagination. Complete adds
	// nested replies, diagnostics, scope, depth, sorting, and original-post controls.
	//
	// Any of "standard", "complete".
	Mode XTweetGetRepliesParamsMode `query:"mode,omitzero" json:"-"`
	// Only when the caller requests a quote mode.
	//
	// Any of "include", "exclude", "only".
	Quotes XTweetGetRepliesParamsQuotes `query:"quotes,omitzero" json:"-"`
	// Only when the caller requests a reply mode.
	//
	// Any of "include", "exclude", "only".
	Replies XTweetGetRepliesParamsReplies `query:"replies,omitzero" json:"-"`
	// Only when the caller requests a repost mode.
	//
	// Any of "include", "exclude", "only".
	Retweets XTweetGetRepliesParamsRetweets `query:"retweets,omitzero" json:"-"`
	// Select all replies, direct replies, or nested replies.
	//
	// Any of "all", "direct", "nested".
	Scope XTweetGetRepliesParamsScope `query:"scope,omitzero" json:"-"`
	// Sort the selected replies before applying limit.
	//
	// Any of "relevance", "latest", "oldest", "likes".
	Sort XTweetGetRepliesParamsSort `query:"sort,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [XTweetGetRepliesParams]'s query parameters as `url.Values`.
func (r XTweetGetRepliesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter media. Aliases: has_video, has_media.
type XTweetGetRepliesParamsMediaType string

const (
	XTweetGetRepliesParamsMediaTypeImages XTweetGetRepliesParamsMediaType = "images"
	XTweetGetRepliesParamsMediaTypeVideos XTweetGetRepliesParamsMediaType = "videos"
	XTweetGetRepliesParamsMediaTypeGifs   XTweetGetRepliesParamsMediaType = "gifs"
	XTweetGetRepliesParamsMediaTypeMedia  XTweetGetRepliesParamsMediaType = "media"
	XTweetGetRepliesParamsMediaTypeLinks  XTweetGetRepliesParamsMediaType = "links"
	XTweetGetRepliesParamsMediaTypeNone   XTweetGetRepliesParamsMediaType = "none"
)

// Override automatic coverage. Standard uses legacy pagination. Complete adds
// nested replies, diagnostics, scope, depth, sorting, and original-post controls.
type XTweetGetRepliesParamsMode string

const (
	XTweetGetRepliesParamsModeStandard XTweetGetRepliesParamsMode = "standard"
	XTweetGetRepliesParamsModeComplete XTweetGetRepliesParamsMode = "complete"
)

// Only when the caller requests a quote mode.
type XTweetGetRepliesParamsQuotes string

const (
	XTweetGetRepliesParamsQuotesInclude XTweetGetRepliesParamsQuotes = "include"
	XTweetGetRepliesParamsQuotesExclude XTweetGetRepliesParamsQuotes = "exclude"
	XTweetGetRepliesParamsQuotesOnly    XTweetGetRepliesParamsQuotes = "only"
)

// Only when the caller requests a reply mode.
type XTweetGetRepliesParamsReplies string

const (
	XTweetGetRepliesParamsRepliesInclude XTweetGetRepliesParamsReplies = "include"
	XTweetGetRepliesParamsRepliesExclude XTweetGetRepliesParamsReplies = "exclude"
	XTweetGetRepliesParamsRepliesOnly    XTweetGetRepliesParamsReplies = "only"
)

// Only when the caller requests a repost mode.
type XTweetGetRepliesParamsRetweets string

const (
	XTweetGetRepliesParamsRetweetsInclude XTweetGetRepliesParamsRetweets = "include"
	XTweetGetRepliesParamsRetweetsExclude XTweetGetRepliesParamsRetweets = "exclude"
	XTweetGetRepliesParamsRetweetsOnly    XTweetGetRepliesParamsRetweets = "only"
)

// Select all replies, direct replies, or nested replies.
type XTweetGetRepliesParamsScope string

const (
	XTweetGetRepliesParamsScopeAll    XTweetGetRepliesParamsScope = "all"
	XTweetGetRepliesParamsScopeDirect XTweetGetRepliesParamsScope = "direct"
	XTweetGetRepliesParamsScopeNested XTweetGetRepliesParamsScope = "nested"
)

// Sort the selected replies before applying limit.
type XTweetGetRepliesParamsSort string

const (
	XTweetGetRepliesParamsSortRelevance XTweetGetRepliesParamsSort = "relevance"
	XTweetGetRepliesParamsSortLatest    XTweetGetRepliesParamsSort = "latest"
	XTweetGetRepliesParamsSortOldest    XTweetGetRepliesParamsSort = "oldest"
	XTweetGetRepliesParamsSortLikes     XTweetGetRepliesParamsSort = "likes"
)

type XTweetGetRetweetersParams struct {
	// Match any comma-separated or line-separated bio term, ignoring case.
	BioContains param.Opt[string] `query:"bioContains,omitzero" json:"-"`
	// Only return profiles X marks as blue verified. X marks many business &
	// government accounts blue too. Use verifiedType for 1 badge type.
	BlueVerifiedOnly param.Opt[bool] `query:"blueVerifiedOnly,omitzero" json:"-"`
	// Only return profiles that publish an email outside free webmail, such as Gmail.
	BusinessEmailOnly param.Opt[bool] `query:"businessEmailOnly,omitzero" json:"-"`
	// Pagination cursor for retweeters
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Only return profiles with an email at one of these domains. Separate domains
	// with commas. Subdomains match.
	EmailDomains param.Opt[string] `query:"emailDomains,omitzero" json:"-"`
	// Add each row's full profile, as GET /x/users/{id} returns it. Pages take several
	// seconds longer.
	EnrichProfiles param.Opt[bool] `query:"enrichProfiles,omitzero" json:"-"`
	// Fields to keep in each row, comma-separated. Use dots for nested fields, such as
	// author.username. A list keeps the field in each item. Case & underscores do not
	// matter. A row that lacks a field leaves it out. Cursors, counts & diagnostics
	// always stay. An unknown name answers 400 unknown_field with the valid names,
	// before any read. The price stays the same. SDK row types then hold only the
	// named fields.
	Fields param.Opt[string] `query:"fields,omitzero" json:"-"`
	// Only return profiles that publish an email.
	HasEmail param.Opt[bool] `query:"hasEmail,omitzero" json:"-"`
	// Only return profiles with a location.
	HasLocation param.Opt[bool] `query:"hasLocation,omitzero" json:"-"`
	// Only return profiles with a website.
	HasWebsite param.Opt[bool] `query:"hasWebsite,omitzero" json:"-"`
	// Look up repost event times.
	IncludeRetweetTimestamp param.Opt[bool] `query:"includeRetweetTimestamp,omitzero" json:"-"`
	// Match a location substring, ignoring case.
	LocationContains param.Opt[string] `query:"locationContains,omitzero" json:"-"`
	// Maximum follower count. Missing counts pass this maximum.
	MaxFollowers param.Opt[int64] `query:"maxFollowers,omitzero" json:"-"`
	// Profiles may follow at most this many accounts.
	MaxFollowing param.Opt[int64] `query:"maxFollowing,omitzero" json:"-"`
	// Maximum post count. maxPosts is also accepted.
	MaxStatuses param.Opt[int64] `query:"maxStatuses,omitzero" json:"-"`
	// Minimum account age in whole days.
	MinAccountAgeDays param.Opt[int64] `query:"minAccountAgeDays,omitzero" json:"-"`
	// Minimum follower count. Filtering happens before billing.
	MinFollowers param.Opt[int64] `query:"minFollowers,omitzero" json:"-"`
	// Profiles must follow at least this many accounts.
	MinFollowing param.Opt[int64] `query:"minFollowing,omitzero" json:"-"`
	// Minimum post count. minPosts is also accepted.
	MinStatuses param.Opt[int64] `query:"minStatuses,omitzero" json:"-"`
	// Fields to drop from each row, comma-separated. It takes the same names as
	// fields. With fields, it drops from the kept fields.
	Omit param.Opt[string] `query:"omit,omitzero" json:"-"`
	// Maximum user profiles requested from this page (1-200, default 200). Source,
	// filters, or credits can return fewer profiles. Follow next_cursor while the
	// response reports more pages. Deprecated aliases remain accepted.
	PageSize param.Opt[int64] `query:"pageSize,omitzero" json:"-"`
	// Match a username substring, ignoring case.
	UsernameContains param.Opt[string] `query:"usernameContains,omitzero" json:"-"`
	// Only return verified profiles.
	VerifiedOnly param.Opt[bool] `query:"verifiedOnly,omitzero" json:"-"`
	// Match the verification type exactly, ignoring case.
	VerifiedType param.Opt[string] `query:"verifiedType,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [XTweetGetRetweetersParams]'s query parameters as
// `url.Values`.
func (r XTweetGetRetweetersParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type XTweetGetThreadParams struct {
	// Words or quoted phrases where any one can match. Separate with spaces, commas,
	// or lines.
	AnyWords param.Opt[string] `query:"anyWords,omitzero" json:"-"`
	// Only return tweets from Blue-verified authors.
	BlueVerifiedOnly param.Opt[bool] `query:"blueVerifiedOnly,omitzero" json:"-"`
	// Cashtags separated by spaces, commas, or lines.
	Cashtags param.Opt[string] `query:"cashtags,omitzero" json:"-"`
	// Conversation ID filter.
	ConversationID param.Opt[string] `query:"conversationId,omitzero" json:"-"`
	// Pagination cursor for thread tweets
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Match this literal phrase, including any hyphens.
	ExactPhrase param.Opt[string] `query:"exactPhrase,omitzero" json:"-"`
	// Words or quoted phrases to exclude. Separate with spaces, commas, or lines.
	ExcludeWords param.Opt[string] `query:"excludeWords,omitzero" json:"-"`
	// Fields to keep in each row, comma-separated. Use dots for nested fields, such as
	// author.username. A list keeps the field in each item. Case & underscores do not
	// matter. A row that lacks a field leaves it out. Cursors, counts & diagnostics
	// always stay. An unknown name answers 400 unknown_field with the valid names,
	// before any read. The price stays the same. SDK row types then hold only the
	// named fields.
	Fields param.Opt[string] `query:"fields,omitzero" json:"-"`
	// Filter by author username.
	FromUser param.Opt[string] `query:"fromUser,omitzero" json:"-"`
	// Hashtags separated by spaces, commas, or lines.
	Hashtags param.Opt[string] `query:"hashtags,omitzero" json:"-"`
	// Only replies to this tweet ID.
	InReplyToTweetID param.Opt[string] `query:"inReplyToTweetId,omitzero" json:"-"`
	// Filter by language. Alias `lang` is accepted.
	Language param.Opt[string] `query:"language,omitzero" json:"-"`
	// Maximum likes threshold. maxLikes is also accepted.
	MaxFaves param.Opt[int64] `query:"maxFaves,omitzero" json:"-"`
	// Return Tweets at or below this ID.
	MaxID param.Opt[string] `query:"maxId,omitzero" json:"-"`
	// Maximum quotes threshold.
	MaxQuotes param.Opt[int64] `query:"maxQuotes,omitzero" json:"-"`
	// Maximum replies threshold.
	MaxReplies param.Opt[int64] `query:"maxReplies,omitzero" json:"-"`
	// Maximum retweets threshold.
	MaxRetweets param.Opt[int64] `query:"maxRetweets,omitzero" json:"-"`
	// Filter tweets mentioning a username.
	Mentioning param.Opt[string] `query:"mentioning,omitzero" json:"-"`
	// Minimum bookmark count threshold.
	MinBookmarks param.Opt[int64] `query:"minBookmarks,omitzero" json:"-"`
	// Deprecated: Use MinLikes. MinFaves sends the minFaves query key, which the
	// API reads the same way as minLikes.
	MinFaves param.Opt[int64] `query:"minFaves,omitzero" json:"-"`
	// Minimum likes. Aliases: minFaves, min_likes, min_faves.
	MinLikes param.Opt[int64] `query:"minLikes,omitzero" json:"-"`
	// Minimum quote count threshold.
	MinQuotes param.Opt[int64] `query:"minQuotes,omitzero" json:"-"`
	// Minimum replies threshold.
	MinReplies param.Opt[int64] `query:"minReplies,omitzero" json:"-"`
	// Minimum retweets threshold.
	MinRetweets param.Opt[int64] `query:"minRetweets,omitzero" json:"-"`
	// Minimum view count threshold.
	MinViews param.Opt[int64] `query:"minViews,omitzero" json:"-"`
	// Only return native reposts.
	NativeRetweets param.Opt[bool] `query:"nativeRetweets,omitzero" json:"-"`
	// Fields to drop from each row, comma-separated. It takes the same names as
	// fields. With fields, it drops from the kept fields.
	Omit param.Opt[string] `query:"omit,omitzero" json:"-"`
	// Maximum page items (1-100, default 20). Source, filters, or credits can reduce
	// results. Follow next_cursor while the response reports more pages. Deprecated
	// limit and count aliases remain accepted.
	PageSize param.Opt[int64] `query:"pageSize,omitzero" json:"-"`
	// Only quotes of this tweet ID.
	QuotesOfTweetID param.Opt[string] `query:"quotesOfTweetId,omitzero" json:"-"`
	// Only retweets of this tweet ID.
	RetweetsOfTweetID param.Opt[string] `query:"retweetsOfTweetId,omitzero" json:"-"`
	// Start date in YYYY-MM-DD format.
	SinceDate param.Opt[time.Time] `query:"sinceDate,omitzero" format:"date" json:"-"`
	// Return Tweets whose IDs exceed this ID.
	SinceID param.Opt[string] `query:"sinceId,omitzero" json:"-"`
	// Filter replies sent to a username.
	ToUser param.Opt[string] `query:"toUser,omitzero" json:"-"`
	// End date in YYYY-MM-DD format.
	UntilDate param.Opt[time.Time] `query:"untilDate,omitzero" format:"date" json:"-"`
	// URL substring or domain filter.
	URL param.Opt[string] `query:"url,omitzero" json:"-"`
	// Only return tweets from verified authors.
	VerifiedOnly param.Opt[bool] `query:"verifiedOnly,omitzero" json:"-"`
	// Keeps Tweets from a recent window, such as 90m or 7d. Use a whole number & s, m,
	// h or d.
	WithinTime param.Opt[string] `query:"withinTime,omitzero" json:"-"`
	// Filter media. Aliases: has_video, has_media.
	//
	// Any of "images", "videos", "gifs", "media", "links", "none".
	MediaType XTweetGetThreadParamsMediaType `query:"mediaType,omitzero" json:"-"`
	// Only when the caller requests a quote mode.
	//
	// Any of "include", "exclude", "only".
	Quotes XTweetGetThreadParamsQuotes `query:"quotes,omitzero" json:"-"`
	// Only when the caller requests a reply mode.
	//
	// Any of "include", "exclude", "only".
	Replies XTweetGetThreadParamsReplies `query:"replies,omitzero" json:"-"`
	// Only when the caller requests a repost mode.
	//
	// Any of "include", "exclude", "only".
	Retweets XTweetGetThreadParamsRetweets `query:"retweets,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [XTweetGetThreadParams]'s query parameters as `url.Values`.
func (r XTweetGetThreadParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter media. Aliases: has_video, has_media.
type XTweetGetThreadParamsMediaType string

const (
	XTweetGetThreadParamsMediaTypeImages XTweetGetThreadParamsMediaType = "images"
	XTweetGetThreadParamsMediaTypeVideos XTweetGetThreadParamsMediaType = "videos"
	XTweetGetThreadParamsMediaTypeGifs   XTweetGetThreadParamsMediaType = "gifs"
	XTweetGetThreadParamsMediaTypeMedia  XTweetGetThreadParamsMediaType = "media"
	XTweetGetThreadParamsMediaTypeLinks  XTweetGetThreadParamsMediaType = "links"
	XTweetGetThreadParamsMediaTypeNone   XTweetGetThreadParamsMediaType = "none"
)

// Only when the caller requests a quote mode.
type XTweetGetThreadParamsQuotes string

const (
	XTweetGetThreadParamsQuotesInclude XTweetGetThreadParamsQuotes = "include"
	XTweetGetThreadParamsQuotesExclude XTweetGetThreadParamsQuotes = "exclude"
	XTweetGetThreadParamsQuotesOnly    XTweetGetThreadParamsQuotes = "only"
)

// Only when the caller requests a reply mode.
type XTweetGetThreadParamsReplies string

const (
	XTweetGetThreadParamsRepliesInclude XTweetGetThreadParamsReplies = "include"
	XTweetGetThreadParamsRepliesExclude XTweetGetThreadParamsReplies = "exclude"
	XTweetGetThreadParamsRepliesOnly    XTweetGetThreadParamsReplies = "only"
)

// Only when the caller requests a repost mode.
type XTweetGetThreadParamsRetweets string

const (
	XTweetGetThreadParamsRetweetsInclude XTweetGetThreadParamsRetweets = "include"
	XTweetGetThreadParamsRetweetsExclude XTweetGetThreadParamsRetweets = "exclude"
	XTweetGetThreadParamsRetweetsOnly    XTweetGetThreadParamsRetweets = "only"
)

type XTweetSearchParams struct {
	// Query, Tweet ID, or URL. Hyphens negate terms. Use exactPhrase for literals.
	// Valid bounds apply per page.
	Q string `query:"q" api:"required" json:"-"`
	// Raw advanced search query appended as-is.
	AdvancedQuery param.Opt[string] `query:"advancedQuery,omitzero" json:"-"`
	// Words or quoted phrases where any one can match. Separate with spaces, commas,
	// or lines.
	AnyWords param.Opt[string] `query:"anyWords,omitzero" json:"-"`
	// Only return tweets from Blue-verified authors.
	BlueVerifiedOnly param.Opt[bool] `query:"blueVerifiedOnly,omitzero" json:"-"`
	// Geo bounding box, e.g. -74.1 40.6 -73.9 40.8.
	BoundingBox param.Opt[string] `query:"boundingBox,omitzero" json:"-"`
	// Match the Tweet card name.
	CardName param.Opt[string] `query:"cardName,omitzero" json:"-"`
	// Cashtags separated by spaces, commas, or lines.
	Cashtags param.Opt[string] `query:"cashtags,omitzero" json:"-"`
	// Conversation ID filter.
	ConversationID param.Opt[string] `query:"conversationId,omitzero" json:"-"`
	// Cursor from the previous response. Xquik cursors resume automatic coverage.
	// Existing unprefixed cursors keep legacy standard behavior.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Match this literal phrase, including any hyphens.
	ExactPhrase param.Opt[string] `query:"exactPhrase,omitzero" json:"-"`
	// Exclude a source application.
	ExcludeSource param.Opt[string] `query:"excludeSource,omitzero" json:"-"`
	// Words or quoted phrases to exclude. Separate with spaces, commas, or lines.
	ExcludeWords param.Opt[string] `query:"excludeWords,omitzero" json:"-"`
	// Fields to keep in each row, comma-separated. Use dots for nested fields, such as
	// author.username. A list keeps the field in each item. Case & underscores do not
	// matter. A row that lacks a field leaves it out. Cursors, counts & diagnostics
	// always stay. An unknown name answers 400 unknown_field with the valid names,
	// before any read. The price stays the same. SDK row types then hold only the
	// named fields.
	Fields param.Opt[string] `query:"fields,omitzero" json:"-"`
	// Filter by author username.
	FromUser param.Opt[string] `query:"fromUser,omitzero" json:"-"`
	// Match latitude, longitude, and radius.
	Geocode param.Opt[string] `query:"geocode,omitzero" json:"-"`
	// Hashtags separated by spaces, commas, or lines.
	Hashtags param.Opt[string] `query:"hashtags,omitzero" json:"-"`
	// Only replies to this tweet ID.
	InReplyToTweetID param.Opt[string] `query:"inReplyToTweetId,omitzero" json:"-"`
	// Words the Tweets must match, in X search syntax.
	Keywords param.Opt[string] `query:"keywords,omitzero" json:"-"`
	// Filter by language. Alias `lang` is accepted.
	Language param.Opt[string] `query:"language,omitzero" json:"-"`
	// Unique matching result upper bound after filtering. Default 20. Explicit
	// coverage defaults to 2000. It returns retained rows and deadline diagnostics
	// when time expires. Only returned rows are billed. Credits may reduce results;
	// zero affordable rows returns 402. Aliases: pageSize, count, max_results.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Search within a list ID.
	ListID param.Opt[string] `query:"listId,omitzero" json:"-"`
	// Maximum likes threshold. maxLikes is also accepted.
	MaxFaves param.Opt[int64] `query:"maxFaves,omitzero" json:"-"`
	// Return Tweets older than this Tweet ID.
	MaxID param.Opt[string] `query:"maxId,omitzero" json:"-"`
	// Maximum quotes threshold.
	MaxQuotes param.Opt[int64] `query:"maxQuotes,omitzero" json:"-"`
	// Maximum replies threshold.
	MaxReplies param.Opt[int64] `query:"maxReplies,omitzero" json:"-"`
	// Maximum retweets threshold.
	MaxRetweets param.Opt[int64] `query:"maxRetweets,omitzero" json:"-"`
	// Filter tweets mentioning a username.
	Mentioning param.Opt[string] `query:"mentioning,omitzero" json:"-"`
	// Minimum bookmark count threshold.
	MinBookmarks param.Opt[int64] `query:"minBookmarks,omitzero" json:"-"`
	// Deprecated: Use MinLikes. MinFaves sends the minFaves query key, which the
	// API reads the same way as minLikes.
	MinFaves param.Opt[int64] `query:"minFaves,omitzero" json:"-"`
	// Minimum likes. Aliases: minFaves, min_likes, min_faves.
	MinLikes param.Opt[int64] `query:"minLikes,omitzero" json:"-"`
	// Minimum quote count threshold.
	MinQuotes param.Opt[int64] `query:"minQuotes,omitzero" json:"-"`
	// Minimum replies threshold.
	MinReplies param.Opt[int64] `query:"minReplies,omitzero" json:"-"`
	// Minimum retweets threshold.
	MinRetweets param.Opt[int64] `query:"minRetweets,omitzero" json:"-"`
	// Minimum view count threshold.
	MinViews param.Opt[int64] `query:"minViews,omitzero" json:"-"`
	// Only return native reposts.
	NativeRetweets param.Opt[bool] `query:"nativeRetweets,omitzero" json:"-"`
	// Match a place name.
	Near param.Opt[string] `query:"near,omitzero" json:"-"`
	// Only return news results.
	News param.Opt[bool] `query:"news,omitzero" json:"-"`
	// Fields to drop from each row, comma-separated. It takes the same names as
	// fields. With fields, it drops from the kept fields.
	Omit param.Opt[string] `query:"omit,omitzero" json:"-"`
	// Search within a place ID.
	Place param.Opt[string] `query:"place,omitzero" json:"-"`
	// Search within a country code.
	PlaceCountry param.Opt[string] `query:"placeCountry,omitzero" json:"-"`
	// Geo point radius, e.g. -73.99 40.73 25mi.
	PointRadius param.Opt[string] `query:"pointRadius,omitzero" json:"-"`
	// Only quotes of this tweet ID.
	QuotesOfTweetID param.Opt[string] `query:"quotesOfTweetId,omitzero" json:"-"`
	// Only retweets of this tweet ID.
	RetweetsOfTweetID param.Opt[string] `query:"retweetsOfTweetId,omitzero" json:"-"`
	// Enable the safe-search filter.
	Safe param.Opt[bool] `query:"safe,omitzero" json:"-"`
	// Start date in YYYY-MM-DD format.
	SinceDate param.Opt[time.Time] `query:"sinceDate,omitzero" format:"date" json:"-"`
	// Return Tweets newer than this Tweet ID.
	SinceID param.Opt[string] `query:"sinceId,omitzero" json:"-"`
	// Inclusive ISO bound for Tweet creation time.
	SinceTime param.Opt[string] `query:"sinceTime,omitzero" json:"-"`
	// Match the source application.
	Source param.Opt[string] `query:"source,omitzero" json:"-"`
	// Filter replies sent to a username.
	ToUser param.Opt[string] `query:"toUser,omitzero" json:"-"`
	// End date in YYYY-MM-DD format.
	UntilDate param.Opt[time.Time] `query:"untilDate,omitzero" format:"date" json:"-"`
	// Exclusive ISO bound for Tweet creation time.
	UntilTime param.Opt[string] `query:"untilTime,omitzero" json:"-"`
	// URL substring or domain filter.
	URL param.Opt[string] `query:"url,omitzero" json:"-"`
	// Only return tweets from verified authors.
	VerifiedOnly param.Opt[bool] `query:"verifiedOnly,omitzero" json:"-"`
	// Set the radius for the near filter.
	Within param.Opt[string] `query:"within,omitzero" json:"-"`
	// Match Tweets inside a recent time window.
	WithinTime param.Opt[string] `query:"withinTime,omitzero" json:"-"`
	// Filter media. Aliases: has_video, has_media.
	//
	// Any of "images", "videos", "gifs", "media", "links", "none".
	MediaType XTweetSearchParamsMediaType `query:"mediaType,omitzero" json:"-"`
	// Omit mode for resumable maximum coverage. Standard keeps legacy pagination.
	// Coverage returns diagnostics once and rejects cursors.
	//
	// Any of "standard", "coverage".
	Mode XTweetSearchParamsMode `query:"mode,omitzero" json:"-"`
	// Latest is chronological; Top ranks engagement. Aliases: result_type, sort_order.
	//
	// Any of "Latest", "Top".
	QueryType XTweetSearchParamsQueryType `query:"queryType,omitzero" json:"-"`
	// Only when the caller requests a quote mode.
	//
	// Any of "include", "exclude", "only".
	Quotes XTweetSearchParamsQuotes `query:"quotes,omitzero" json:"-"`
	// Only when the caller requests a reply mode.
	//
	// Any of "include", "exclude", "only".
	Replies XTweetSearchParamsReplies `query:"replies,omitzero" json:"-"`
	// Only when the caller requests a repost mode.
	//
	// Any of "include", "exclude", "only".
	Retweets XTweetSearchParamsRetweets `query:"retweets,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [XTweetSearchParams]'s query parameters as `url.Values`.
func (r XTweetSearchParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter media. Aliases: has_video, has_media.
type XTweetSearchParamsMediaType string

const (
	XTweetSearchParamsMediaTypeImages XTweetSearchParamsMediaType = "images"
	XTweetSearchParamsMediaTypeVideos XTweetSearchParamsMediaType = "videos"
	XTweetSearchParamsMediaTypeGifs   XTweetSearchParamsMediaType = "gifs"
	XTweetSearchParamsMediaTypeMedia  XTweetSearchParamsMediaType = "media"
	XTweetSearchParamsMediaTypeLinks  XTweetSearchParamsMediaType = "links"
	XTweetSearchParamsMediaTypeNone   XTweetSearchParamsMediaType = "none"
)

// Omit mode for resumable maximum coverage. Standard keeps legacy pagination.
// Coverage returns diagnostics once and rejects cursors.
type XTweetSearchParamsMode string

const (
	XTweetSearchParamsModeStandard XTweetSearchParamsMode = "standard"
	XTweetSearchParamsModeCoverage XTweetSearchParamsMode = "coverage"
)

// Latest is chronological; Top ranks engagement. Aliases: result_type, sort_order.
type XTweetSearchParamsQueryType string

const (
	XTweetSearchParamsQueryTypeLatest XTweetSearchParamsQueryType = "Latest"
	XTweetSearchParamsQueryTypeTop    XTweetSearchParamsQueryType = "Top"
)

// Only when the caller requests a quote mode.
type XTweetSearchParamsQuotes string

const (
	XTweetSearchParamsQuotesInclude XTweetSearchParamsQuotes = "include"
	XTweetSearchParamsQuotesExclude XTweetSearchParamsQuotes = "exclude"
	XTweetSearchParamsQuotesOnly    XTweetSearchParamsQuotes = "only"
)

// Only when the caller requests a reply mode.
type XTweetSearchParamsReplies string

const (
	XTweetSearchParamsRepliesInclude XTweetSearchParamsReplies = "include"
	XTweetSearchParamsRepliesExclude XTweetSearchParamsReplies = "exclude"
	XTweetSearchParamsRepliesOnly    XTweetSearchParamsReplies = "only"
)

// Only when the caller requests a repost mode.
type XTweetSearchParamsRetweets string

const (
	XTweetSearchParamsRetweetsInclude XTweetSearchParamsRetweets = "include"
	XTweetSearchParamsRetweetsExclude XTweetSearchParamsRetweets = "exclude"
	XTweetSearchParamsRetweetsOnly    XTweetSearchParamsRetweets = "only"
)

type XTweetCheckRepostParams struct {
	// User ID, username with or without @, or profile URL.
	User string `query:"user" api:"required" json:"-"`
	// Fields to keep in each row, comma-separated. Use dots for nested fields, such as
	// author.username. A list keeps the field in each item. Case & underscores do not
	// matter. A row that lacks a field leaves it out. Cursors, counts & diagnostics
	// always stay. An unknown name answers 400 unknown_field with the valid names,
	// before any read. The price stays the same. SDK row types then hold only the
	// named fields.
	Fields param.Opt[string] `query:"fields,omitzero" json:"-"`
	// Fields to drop from each row, comma-separated. It takes the same names as
	// fields. With fields, it drops from the kept fields.
	Omit param.Opt[string] `query:"omit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [XTweetCheckRepostParams]'s query parameters as `url.Values`.
func (r XTweetCheckRepostParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type XTweetEmbedParams struct {
	// Fields to keep in each row, comma-separated. Use dots for nested fields, such as
	// author.username. A list keeps the field in each item. Case & underscores do not
	// matter. A row that lacks a field leaves it out. Cursors, counts & diagnostics
	// always stay. An unknown name answers 400 unknown_field with the valid names,
	// before any read. The price stays the same. SDK row types then hold only the
	// named fields.
	Fields param.Opt[string] `query:"fields,omitzero" json:"-"`
	// Set `true` to leave photos, videos, and link cards out.
	HideMedia param.Opt[bool] `query:"hideMedia,omitzero" json:"-"`
	// Set `true` to leave the tweet this one replies to out.
	HideThread param.Opt[bool] `query:"hideThread,omitzero" json:"-"`
	// Fields to drop from each row, comma-separated. It takes the same names as
	// fields. With fields, it drops from the kept fields.
	Omit param.Opt[string] `query:"omit,omitzero" json:"-"`
	// Set `true` to leave the widget script tag out, for pages that load it once
	// themselves.
	OmitScript param.Opt[bool] `query:"omitScript,omitzero" json:"-"`
	// Color theme of the embed. X uses light when left out.
	//
	// Any of "light", "dark".
	Theme XTweetEmbedParamsTheme `query:"theme,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [XTweetEmbedParams]'s query parameters as `url.Values`.
func (r XTweetEmbedParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Color theme of the embed. X uses light when left out.
type XTweetEmbedParamsTheme string

const (
	XTweetEmbedParamsThemeLight XTweetEmbedParamsTheme = "light"
	XTweetEmbedParamsThemeDark  XTweetEmbedParamsTheme = "dark"
)

type XTweetGetHiddenRepliesParams struct {
	// Words or quoted phrases where any one can match. Separate with spaces, commas,
	// or lines.
	AnyWords param.Opt[string] `query:"anyWords,omitzero" json:"-"`
	// Only return tweets from Blue-verified authors.
	BlueVerifiedOnly param.Opt[bool] `query:"blueVerifiedOnly,omitzero" json:"-"`
	// Cashtags separated by spaces, commas, or lines.
	Cashtags param.Opt[string] `query:"cashtags,omitzero" json:"-"`
	// Conversation ID filter.
	ConversationID param.Opt[string] `query:"conversationId,omitzero" json:"-"`
	// Pagination cursor for collection results. Also read as next_cursor, next_token,
	// pagination_token or continuation_token.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Match this literal phrase, including any hyphens.
	ExactPhrase param.Opt[string] `query:"exactPhrase,omitzero" json:"-"`
	// Words or quoted phrases to exclude. Separate with spaces, commas, or lines.
	ExcludeWords param.Opt[string] `query:"excludeWords,omitzero" json:"-"`
	// Fields to keep in each row, comma-separated. Use dots for nested fields, such as
	// author.username. A list keeps the field in each item. Case & underscores do not
	// matter. A row that lacks a field leaves it out. Cursors, counts & diagnostics
	// always stay. An unknown name answers 400 unknown_field with the valid names,
	// before any read. The price stays the same. SDK row types then hold only the
	// named fields.
	Fields param.Opt[string] `query:"fields,omitzero" json:"-"`
	// Filter by author username.
	FromUser param.Opt[string] `query:"fromUser,omitzero" json:"-"`
	// Hashtags separated by spaces, commas, or lines.
	Hashtags param.Opt[string] `query:"hashtags,omitzero" json:"-"`
	// Only replies to this tweet ID.
	InReplyToTweetID param.Opt[string] `query:"inReplyToTweetId,omitzero" json:"-"`
	// Filter by language. Its alias is `lang`.
	Language param.Opt[string] `query:"language,omitzero" json:"-"`
	// Maximum likes threshold. maxLikes & max_likes also work.
	MaxFaves param.Opt[int64] `query:"maxFaves,omitzero" json:"-"`
	// Return Tweets at or below this ID.
	MaxID param.Opt[string] `query:"maxId,omitzero" json:"-"`
	// Maximum quotes threshold.
	MaxQuotes param.Opt[int64] `query:"maxQuotes,omitzero" json:"-"`
	// Maximum replies threshold.
	MaxReplies param.Opt[int64] `query:"maxReplies,omitzero" json:"-"`
	// Maximum retweets threshold.
	MaxRetweets param.Opt[int64] `query:"maxRetweets,omitzero" json:"-"`
	// Filter tweets mentioning a username.
	Mentioning param.Opt[string] `query:"mentioning,omitzero" json:"-"`
	// Minimum bookmark count threshold.
	MinBookmarks param.Opt[int64] `query:"minBookmarks,omitzero" json:"-"`
	// Minimum likes. Its aliases are minFaves, min_likes & min_faves.
	MinLikes param.Opt[int64] `query:"minLikes,omitzero" json:"-"`
	// Minimum quote count threshold.
	MinQuotes param.Opt[int64] `query:"minQuotes,omitzero" json:"-"`
	// Minimum replies threshold.
	MinReplies param.Opt[int64] `query:"minReplies,omitzero" json:"-"`
	// Minimum retweets threshold.
	MinRetweets param.Opt[int64] `query:"minRetweets,omitzero" json:"-"`
	// Minimum view count threshold.
	MinViews param.Opt[int64] `query:"minViews,omitzero" json:"-"`
	// Only return native reposts.
	NativeRetweets param.Opt[bool] `query:"nativeRetweets,omitzero" json:"-"`
	// Fields to drop from each row, comma-separated. It takes the same names as
	// fields. With fields, it drops from the kept fields.
	Omit param.Opt[string] `query:"omit,omitzero" json:"-"`
	// Items per page (1-100, default 20). Source, filters or credits can return fewer.
	// Follow next_cursor for more pages. Its aliases are limit, count, max_results,
	// maxItems, max_items & per_page.
	PageSize param.Opt[int64] `query:"pageSize,omitzero" json:"-"`
	// Only quotes of this tweet ID.
	QuotesOfTweetID param.Opt[string] `query:"quotesOfTweetId,omitzero" json:"-"`
	// Only retweets of this tweet ID.
	RetweetsOfTweetID param.Opt[string] `query:"retweetsOfTweetId,omitzero" json:"-"`
	// Start date in YYYY-MM-DD format, as a UTC day.
	SinceDate param.Opt[time.Time] `query:"sinceDate,omitzero" format:"date" json:"-"`
	// Return Tweets whose IDs exceed this ID.
	SinceID param.Opt[string] `query:"sinceId,omitzero" json:"-"`
	// Filter replies sent to a username.
	ToUser param.Opt[string] `query:"toUser,omitzero" json:"-"`
	// End date in YYYY-MM-DD format, as a UTC day, inclusive, so its own posts count.
	UntilDate param.Opt[time.Time] `query:"untilDate,omitzero" format:"date" json:"-"`
	// URL substring or domain filter.
	URL param.Opt[string] `query:"url,omitzero" json:"-"`
	// Only return tweets from verified authors.
	VerifiedOnly param.Opt[bool] `query:"verifiedOnly,omitzero" json:"-"`
	// Keeps Tweets from a recent window, such as 90m or 7d. Use a whole number & s, m,
	// h or d.
	WithinTime param.Opt[string] `query:"withinTime,omitzero" json:"-"`
	// Filter media. Its aliases are has_video & has_media.
	//
	// Any of "images", "videos", "gifs", "media", "links", "none".
	MediaType XTweetGetHiddenRepliesParamsMediaType `query:"mediaType,omitzero" json:"-"`
	// Only when the caller requests a quote mode.
	//
	// Any of "include", "exclude", "only".
	Quotes XTweetGetHiddenRepliesParamsQuotes `query:"quotes,omitzero" json:"-"`
	// Reply mode. Exclude keeps own-thread replies.
	//
	// Any of "include", "exclude", "only".
	Replies XTweetGetHiddenRepliesParamsReplies `query:"replies,omitzero" json:"-"`
	// Only when the caller requests a repost mode.
	//
	// Any of "include", "exclude", "only".
	Retweets XTweetGetHiddenRepliesParamsRetweets `query:"retweets,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [XTweetGetHiddenRepliesParams]'s query parameters as `url.Values`.
func (r XTweetGetHiddenRepliesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter media. Its aliases are has_video & has_media.
type XTweetGetHiddenRepliesParamsMediaType string

const (
	XTweetGetHiddenRepliesParamsMediaTypeImages XTweetGetHiddenRepliesParamsMediaType = "images"
	XTweetGetHiddenRepliesParamsMediaTypeVideos XTweetGetHiddenRepliesParamsMediaType = "videos"
	XTweetGetHiddenRepliesParamsMediaTypeGifs   XTweetGetHiddenRepliesParamsMediaType = "gifs"
	XTweetGetHiddenRepliesParamsMediaTypeMedia  XTweetGetHiddenRepliesParamsMediaType = "media"
	XTweetGetHiddenRepliesParamsMediaTypeLinks  XTweetGetHiddenRepliesParamsMediaType = "links"
	XTweetGetHiddenRepliesParamsMediaTypeNone   XTweetGetHiddenRepliesParamsMediaType = "none"
)

// Only when the caller requests a quote mode.
type XTweetGetHiddenRepliesParamsQuotes string

const (
	XTweetGetHiddenRepliesParamsQuotesInclude XTweetGetHiddenRepliesParamsQuotes = "include"
	XTweetGetHiddenRepliesParamsQuotesExclude XTweetGetHiddenRepliesParamsQuotes = "exclude"
	XTweetGetHiddenRepliesParamsQuotesOnly    XTweetGetHiddenRepliesParamsQuotes = "only"
)

// Reply mode. Exclude keeps own-thread replies.
type XTweetGetHiddenRepliesParamsReplies string

const (
	XTweetGetHiddenRepliesParamsRepliesInclude XTweetGetHiddenRepliesParamsReplies = "include"
	XTweetGetHiddenRepliesParamsRepliesExclude XTweetGetHiddenRepliesParamsReplies = "exclude"
	XTweetGetHiddenRepliesParamsRepliesOnly    XTweetGetHiddenRepliesParamsReplies = "only"
)

// Only when the caller requests a repost mode.
type XTweetGetHiddenRepliesParamsRetweets string

const (
	XTweetGetHiddenRepliesParamsRetweetsInclude XTweetGetHiddenRepliesParamsRetweets = "include"
	XTweetGetHiddenRepliesParamsRetweetsExclude XTweetGetHiddenRepliesParamsRetweets = "exclude"
	XTweetGetHiddenRepliesParamsRetweetsOnly    XTweetGetHiddenRepliesParamsRetweets = "only"
)

type XTweetGetInsightsParams struct {
	// Fields to keep in each row, comma-separated. Use dots for nested fields, such as
	// author.username. A list keeps the field in each item. Case & underscores do not
	// matter. A row that lacks a field leaves it out. Cursors, counts & diagnostics
	// always stay. An unknown name answers 400 unknown_field with the valid names,
	// before any read. The price stays the same. SDK row types then hold only the
	// named fields.
	Fields param.Opt[string] `query:"fields,omitzero" json:"-"`
	// Fields to drop from each row, comma-separated. It takes the same names as
	// fields. With fields, it drops from the kept fields.
	Omit param.Opt[string] `query:"omit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [XTweetGetInsightsParams]'s query parameters as `url.Values`.
func (r XTweetGetInsightsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type XTweetGetRetweeterIDsParams struct {
	// Pagination cursor for collection results. Also read as next_cursor, next_token,
	// pagination_token or continuation_token.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Profiles to request from this page (1-200, default 200). Source, filters or
	// credits can return fewer. Follow next_cursor for more pages. Its aliases are
	// limit, count, max_results, maxItems, max_items & per_page.
	PageSize param.Opt[int64] `query:"pageSize,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [XTweetGetRetweeterIDsParams]'s query parameters as `url.Values`.
func (r XTweetGetRetweeterIDsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type XTweetGetSubtitlesParams struct {
	// Fields to keep in each row, comma-separated. Use dots for nested fields, such as
	// author.username. A list keeps the field in each item. Case & underscores do not
	// matter. A row that lacks a field leaves it out. Cursors, counts & diagnostics
	// always stay. An unknown name answers 400 unknown_field with the valid names,
	// before any read. The price stays the same. SDK row types then hold only the
	// named fields.
	Fields param.Opt[string] `query:"fields,omitzero" json:"-"`
	// Fields to drop from each row, comma-separated. It takes the same names as
	// fields. With fields, it drops from the kept fields.
	Omit param.Opt[string] `query:"omit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [XTweetGetSubtitlesParams]'s query parameters as `url.Values`.
func (r XTweetGetSubtitlesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type XTweetListTopParams struct {
	// Only tweets from 1 country, as a 2- or 3-letter ISO 3166 code, such as US or
	// USA. Send country or language, not both.
	Country param.Opt[string] `query:"country,omitzero" json:"-"`
	// Pagination cursor for collection results. Also read as next_cursor, next_token,
	// pagination_token or continuation_token.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Fields to keep in each row, comma-separated. Use dots for nested fields, such as
	// author.username. A list keeps the field in each item. Case & underscores do not
	// matter. A row that lacks a field leaves it out. Cursors, counts & diagnostics
	// always stay. An unknown name answers 400 unknown_field with the valid names,
	// before any read. The price stays the same. SDK row types then hold only the
	// named fields.
	Fields param.Opt[string] `query:"fields,omitzero" json:"-"`
	// Only tweets in 1 language, as a 2-letter code, such as en. Send country or
	// language, not both.
	Language param.Opt[string] `query:"language,omitzero" json:"-"`
	// Fields to drop from each row, comma-separated. It takes the same names as
	// fields. With fields, it drops from the kept fields.
	Omit param.Opt[string] `query:"omit,omitzero" json:"-"`
	// Items per page (1-100, default 20). Source, filters or credits can return fewer.
	// Follow next_cursor for more pages. Its aliases are limit, count, max_results,
	// maxItems, max_items & per_page.
	PageSize param.Opt[int64] `query:"pageSize,omitzero" json:"-"`
	// What to rank tweets by. Defaults to likes. Singular names work too, such as
	// like.
	//
	// Any of "likes", "replies", "quotes", "bookmarks", "shares", "views".
	Metric XTweetListTopParamsMetric `query:"metric,omitzero" json:"-"`
	// How far back to look. Defaults to day. Also takes 24h, 7d, 30d, daily, weekly
	// and monthly.
	//
	// Any of "day", "week", "month".
	Window XTweetListTopParamsWindow `query:"window,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [XTweetListTopParams]'s query parameters as `url.Values`.
func (r XTweetListTopParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// What to rank tweets by. Defaults to likes. Singular names work too, such as
// like.
type XTweetListTopParamsMetric string

const (
	XTweetListTopParamsMetricLikes     XTweetListTopParamsMetric = "likes"
	XTweetListTopParamsMetricReplies   XTweetListTopParamsMetric = "replies"
	XTweetListTopParamsMetricQuotes    XTweetListTopParamsMetric = "quotes"
	XTweetListTopParamsMetricBookmarks XTweetListTopParamsMetric = "bookmarks"
	XTweetListTopParamsMetricShares    XTweetListTopParamsMetric = "shares"
	XTweetListTopParamsMetricViews     XTweetListTopParamsMetric = "views"
)

// How far back to look. Defaults to day. Also takes 24h, 7d, 30d, daily, weekly
// and monthly.
type XTweetListTopParamsWindow string

const (
	XTweetListTopParamsWindowDay   XTweetListTopParamsWindow = "day"
	XTweetListTopParamsWindowWeek  XTweetListTopParamsWindow = "week"
	XTweetListTopParamsWindowMonth XTweetListTopParamsWindow = "month"
)

type XTweetResolveLinksParams struct {
	// 1 to 100 links of the form `https://t.co/Xe1nm2pkSX`, separated by commas. Any
	// other link is refused.
	URLs string `query:"urls" api:"required" json:"-"`
	// Fields to keep in each row, comma-separated. Use dots for nested fields, such as
	// author.username. A list keeps the field in each item. Case & underscores do not
	// matter. A row that lacks a field leaves it out. Cursors, counts & diagnostics
	// always stay. An unknown name answers 400 unknown_field with the valid names,
	// before any read. The price stays the same. SDK row types then hold only the
	// named fields.
	Fields param.Opt[string] `query:"fields,omitzero" json:"-"`
	// Fields to drop from each row, comma-separated. It takes the same names as
	// fields. With fields, it drops from the kept fields.
	Omit param.Opt[string] `query:"omit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [XTweetResolveLinksParams]'s query parameters as `url.Values`.
func (r XTweetResolveLinksParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type XTweetTranslateParams struct {
	// Language code to translate into, such as `en`, `tr`, `pt-br`, or `zh-tw`. Upper
	// and lower case both work.
	Language string `query:"language" api:"required" json:"-"`
	// Fields to keep in each row, comma-separated. Use dots for nested fields, such as
	// author.username. A list keeps the field in each item. Case & underscores do not
	// matter. A row that lacks a field leaves it out. Cursors, counts & diagnostics
	// always stay. An unknown name answers 400 unknown_field with the valid names,
	// before any read. The price stays the same. SDK row types then hold only the
	// named fields.
	Fields param.Opt[string] `query:"fields,omitzero" json:"-"`
	// Fields to drop from each row, comma-separated. It takes the same names as
	// fields. With fields, it drops from the kept fields.
	Omit param.Opt[string] `query:"omit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [XTweetTranslateParams]'s query parameters as `url.Values`.
func (r XTweetTranslateParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Latest for newest first, or Top for the quotes X ranks highest. Top is X's
// ranking, not a count order, & a later page can hold a quote X ranks higher. Any
// case works. Alias sort, which also takes Recency & Relevance. Send the same
// value with each cursor. Standard pages list Latest.
type XTweetGetQuotesParamsQueryType string

const (
	XTweetGetQuotesParamsQueryTypeLatest XTweetGetQuotesParamsQueryType = "Latest"
	XTweetGetQuotesParamsQueryTypeTop    XTweetGetQuotesParamsQueryType = "Top"
)
