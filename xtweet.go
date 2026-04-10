// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package xtwitterscraper

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/stainless-sdks/x-twitter-scraper-go/internal/apijson"
	"github.com/stainless-sdks/x-twitter-scraper-go/internal/apiquery"
	"github.com/stainless-sdks/x-twitter-scraper-go/internal/requestconfig"
	"github.com/stainless-sdks/x-twitter-scraper-go/option"
	"github.com/stainless-sdks/x-twitter-scraper-go/packages/param"
	"github.com/stainless-sdks/x-twitter-scraper-go/packages/respjson"
	"github.com/stainless-sdks/x-twitter-scraper-go/shared"
)

// XTweetService contains methods and other services that help with interacting
// with the x-twitter-scraper API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewXTweetService] method instead.
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

// Create tweet
func (r *XTweetService) New(ctx context.Context, body XTweetNewParams, opts ...option.RequestOption) (res *XTweetNewResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "x/tweets"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Get tweet with full text, author, metrics & media
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

// Get multiple tweets by IDs
func (r *XTweetService) List(ctx context.Context, query XTweetListParams, opts ...option.RequestOption) (res *shared.PaginatedTweets, err error) {
	opts = slices.Concat(r.options, opts)
	path := "x/tweets"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Delete tweet
func (r *XTweetService) Delete(ctx context.Context, id string, body XTweetDeleteParams, opts ...option.RequestOption) (res *XTweetDeleteResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("x/tweets/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return res, err
}

// List users who liked a tweet
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

// List quote tweets of a tweet
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

// List replies to a tweet
func (r *XTweetService) GetReplies(ctx context.Context, id string, query XTweetGetRepliesParams, opts ...option.RequestOption) (res *shared.PaginatedTweets, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("x/tweets/%s/replies", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// List users who retweeted a tweet
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

// Get full conversation thread for a tweet
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

// Search tweets with X query operators & pagination
func (r *XTweetService) Search(ctx context.Context, query XTweetSearchParams, opts ...option.RequestOption) (res *shared.PaginatedTweets, err error) {
	opts = slices.Concat(r.options, opts)
	path := "x/tweets/search"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Author of a tweet with follower count and verification status.
type TweetAuthor struct {
	ID             string `json:"id" api:"required"`
	Followers      int64  `json:"followers" api:"required"`
	Username       string `json:"username" api:"required"`
	Verified       bool   `json:"verified" api:"required"`
	ProfilePicture string `json:"profilePicture"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		Followers      respjson.Field
		Username       respjson.Field
		Verified       respjson.Field
		ProfilePicture respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TweetAuthor) RawJSON() string { return r.JSON.raw }
func (r *TweetAuthor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Full tweet with text, engagement metrics, media, and metadata.
type TweetDetail struct {
	ID            string `json:"id" api:"required"`
	BookmarkCount int64  `json:"bookmarkCount" api:"required"`
	LikeCount     int64  `json:"likeCount" api:"required"`
	QuoteCount    int64  `json:"quoteCount" api:"required"`
	ReplyCount    int64  `json:"replyCount" api:"required"`
	RetweetCount  int64  `json:"retweetCount" api:"required"`
	Text          string `json:"text" api:"required"`
	ViewCount     int64  `json:"viewCount" api:"required"`
	// ID of the root tweet in the conversation thread
	ConversationID string `json:"conversationId"`
	CreatedAt      string `json:"createdAt"`
	// Parsed entities from the tweet text (URLs, mentions, hashtags, media)
	Entities map[string]any `json:"entities"`
	// Whether this is a Note Tweet (long-form post, up to 25,000 characters)
	IsNoteTweet bool `json:"isNoteTweet"`
	// Whether this tweet quotes another tweet
	IsQuoteStatus bool `json:"isQuoteStatus"`
	// Whether this tweet is a reply to another tweet
	IsReply bool `json:"isReply"`
	// Attached media items, omitted when the tweet has no media
	Media []TweetDetailMedia `json:"media"`
	// The quoted tweet object, present when isQuoteStatus is true
	QuotedTweet map[string]any `json:"quoted_tweet"`
	// Client application used to post this tweet
	Source string `json:"source"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		BookmarkCount  respjson.Field
		LikeCount      respjson.Field
		QuoteCount     respjson.Field
		ReplyCount     respjson.Field
		RetweetCount   respjson.Field
		Text           respjson.Field
		ViewCount      respjson.Field
		ConversationID respjson.Field
		CreatedAt      respjson.Field
		Entities       respjson.Field
		IsNoteTweet    respjson.Field
		IsQuoteStatus  respjson.Field
		IsReply        respjson.Field
		Media          respjson.Field
		QuotedTweet    respjson.Field
		Source         respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TweetDetail) RawJSON() string { return r.JSON.raw }
func (r *TweetDetail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TweetDetailMedia struct {
	MediaURL string `json:"mediaUrl"`
	// Any of "photo", "video", "animated_gif".
	Type string `json:"type"`
	URL  string `json:"url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MediaURL    respjson.Field
		Type        respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TweetDetailMedia) RawJSON() string { return r.JSON.raw }
func (r *TweetDetailMedia) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type XTweetNewResponse struct {
	Success bool   `json:"success" api:"required"`
	TweetID string `json:"tweetId" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Success     respjson.Field
		TweetID     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XTweetNewResponse) RawJSON() string { return r.JSON.raw }
func (r *XTweetNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type XTweetGetResponse struct {
	// Full tweet with text, engagement metrics, media, and metadata.
	Tweet TweetDetail `json:"tweet" api:"required"`
	// Author of a tweet with follower count and verification status.
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
	return apijson.UnmarshalRoot(data, r)
}

type XTweetDeleteResponse struct {
	Success bool `json:"success" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XTweetDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *XTweetDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type XTweetNewParams struct {
	// X account (@username or account ID)
	Account        string            `json:"account" api:"required"`
	Text           string            `json:"text" api:"required"`
	AttachmentURL  param.Opt[string] `json:"attachment_url,omitzero"`
	CommunityID    param.Opt[string] `json:"community_id,omitzero"`
	IsNoteTweet    param.Opt[bool]   `json:"is_note_tweet,omitzero"`
	ReplyToTweetID param.Opt[string] `json:"reply_to_tweet_id,omitzero"`
	MediaIDs       []string          `json:"media_ids,omitzero"`
	paramObj
}

func (r XTweetNewParams) MarshalJSON() (data []byte, err error) {
	type shadow XTweetNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *XTweetNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type XTweetListParams struct {
	// Comma-separated tweet IDs (max 100)
	IDs string `query:"ids" api:"required" json:"-"`
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
	Account string `json:"account" api:"required"`
	paramObj
}

func (r XTweetDeleteParams) MarshalJSON() (data []byte, err error) {
	type shadow XTweetDeleteParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *XTweetDeleteParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type XTweetGetFavoritersParams struct {
	// Pagination cursor for favoriters
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
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
	// Pagination cursor for quote tweets
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Include reply quotes (default false)
	IncludeReplies param.Opt[bool] `query:"includeReplies,omitzero" json:"-"`
	// Unix timestamp - return quotes posted after this time
	SinceTime param.Opt[string] `query:"sinceTime,omitzero" json:"-"`
	// Unix timestamp - return quotes posted before this time
	UntilTime param.Opt[string] `query:"untilTime,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [XTweetGetQuotesParams]'s query parameters as `url.Values`.
func (r XTweetGetQuotesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type XTweetGetRepliesParams struct {
	// Pagination cursor for tweet replies
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Unix timestamp - return replies posted after this time
	SinceTime param.Opt[string] `query:"sinceTime,omitzero" json:"-"`
	// Unix timestamp - return replies posted before this time
	UntilTime param.Opt[string] `query:"untilTime,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [XTweetGetRepliesParams]'s query parameters as `url.Values`.
func (r XTweetGetRepliesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type XTweetGetRetweetersParams struct {
	// Pagination cursor for retweeters
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
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
	// Pagination cursor for thread tweets
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [XTweetGetThreadParams]'s query parameters as `url.Values`.
func (r XTweetGetThreadParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type XTweetSearchParams struct {
	// Search query (keywords,
	Q string `query:"q" api:"required" json:"-"`
	// Pagination cursor from previous response
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Max tweets to return (server paginates internally). Omit for single page (~20).
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// ISO 8601 timestamp — only return tweets after this time
	SinceTime param.Opt[string] `query:"sinceTime,omitzero" json:"-"`
	// ISO 8601 timestamp — only return tweets before this time
	UntilTime param.Opt[string] `query:"untilTime,omitzero" json:"-"`
	// Sort order — Latest (chronological) or Top (engagement-ranked)
	//
	// Any of "Latest", "Top".
	QueryType XTweetSearchParamsQueryType `query:"queryType,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [XTweetSearchParams]'s query parameters as `url.Values`.
func (r XTweetSearchParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Sort order — Latest (chronological) or Top (engagement-ranked)
type XTweetSearchParamsQueryType string

const (
	XTweetSearchParamsQueryTypeLatest XTweetSearchParamsQueryType = "Latest"
	XTweetSearchParamsQueryTypeTop    XTweetSearchParamsQueryType = "Top"
)
