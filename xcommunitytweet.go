// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

package xtwitterscraper

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/Xquik-dev/x-twitter-scraper-go/internal/apiquery"
	"github.com/Xquik-dev/x-twitter-scraper-go/internal/requestconfig"
	"github.com/Xquik-dev/x-twitter-scraper-go/option"
	"github.com/Xquik-dev/x-twitter-scraper-go/packages/param"
	"github.com/Xquik-dev/x-twitter-scraper-go/shared"
)

// XCommunityTweetService covers X Community info, members, and tweets.
//
// XCommunityTweetService contains methods and other services that help with
// interacting with the x-twitter-scraper API.
//
// You should not instantiate this service directly, and instead use the
// [NewXCommunityTweetService] method instead.
type XCommunityTweetService struct {
	options []option.RequestOption
}

// NewXCommunityTweetService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewXCommunityTweetService(opts ...option.RequestOption) (r XCommunityTweetService) {
	r = XCommunityTweetService{}
	r.options = opts
	return
}

// List returns one resumable page. Requires a Community ID and query.
func (r *XCommunityTweetService) List(ctx context.Context, query XCommunityTweetListParams, opts ...option.RequestOption) (res *shared.PaginatedTweets, err error) {
	opts = slices.Concat(r.options, opts)
	path := "x/communities/tweets"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// ListByCommunity returns public tweets posted within one community.
func (r *XCommunityTweetService) ListByCommunity(ctx context.Context, id string, query XCommunityTweetListByCommunityParams, opts ...option.RequestOption) (res *shared.PaginatedTweets, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("x/communities/%s/tweets", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type XCommunityTweetListParams struct {
	// Numeric ID of the community whose posts to search
	CommunityID string `query:"communityId" api:"required" json:"-"`
	// Search query
	Q string `query:"q" api:"required" json:"-"`
	// Words or quoted phrases where any one can match. Separate with spaces, commas,
	// or lines.
	AnyWords param.Opt[string] `query:"anyWords,omitzero" json:"-"`
	// Only return tweets from Blue-verified authors.
	BlueVerifiedOnly param.Opt[bool] `query:"blueVerifiedOnly,omitzero" json:"-"`
	// Cashtags separated by spaces, commas, or lines.
	Cashtags param.Opt[string] `query:"cashtags,omitzero" json:"-"`
	// Conversation ID filter.
	ConversationID param.Opt[string] `query:"conversationId,omitzero" json:"-"`
	// Pagination cursor for community search
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
	MediaType XCommunityTweetListParamsMediaType `query:"mediaType,omitzero" json:"-"`
	// Sort order (Latest or Top)
	//
	// Any of "Latest", "Top".
	QueryType XCommunityTweetListParamsQueryType `query:"queryType,omitzero" json:"-"`
	// Only when the caller requests a quote mode.
	//
	// Any of "include", "exclude", "only".
	Quotes XCommunityTweetListParamsQuotes `query:"quotes,omitzero" json:"-"`
	// Reply mode. Exclude keeps own-thread replies.
	//
	// Any of "include", "exclude", "only".
	Replies XCommunityTweetListParamsReplies `query:"replies,omitzero" json:"-"`
	// Only when the caller requests a repost mode.
	//
	// Any of "include", "exclude", "only".
	Retweets XCommunityTweetListParamsRetweets `query:"retweets,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [XCommunityTweetListParams]'s query parameters as
// `url.Values`.
func (r XCommunityTweetListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// XCommunityTweetListParamsMediaType is the type of [XCommunityTweetListParams.MediaType]. Filter media. Aliases: has_video, has_media.
type XCommunityTweetListParamsMediaType string

const (
	XCommunityTweetListParamsMediaTypeImages XCommunityTweetListParamsMediaType = "images"
	XCommunityTweetListParamsMediaTypeVideos XCommunityTweetListParamsMediaType = "videos"
	XCommunityTweetListParamsMediaTypeGifs   XCommunityTweetListParamsMediaType = "gifs"
	XCommunityTweetListParamsMediaTypeMedia  XCommunityTweetListParamsMediaType = "media"
	XCommunityTweetListParamsMediaTypeLinks  XCommunityTweetListParamsMediaType = "links"
	XCommunityTweetListParamsMediaTypeNone   XCommunityTweetListParamsMediaType = "none"
)

// XCommunityTweetListParamsQueryType is a sort order (Latest or Top).
type XCommunityTweetListParamsQueryType string

const (
	XCommunityTweetListParamsQueryTypeLatest XCommunityTweetListParamsQueryType = "Latest"
	XCommunityTweetListParamsQueryTypeTop    XCommunityTweetListParamsQueryType = "Top"
)

type XCommunityTweetListByCommunityParams struct {
	// Words or quoted phrases where any one can match. Separate with spaces, commas,
	// or lines.
	AnyWords param.Opt[string] `query:"anyWords,omitzero" json:"-"`
	// Only return tweets from Blue-verified authors.
	BlueVerifiedOnly param.Opt[bool] `query:"blueVerifiedOnly,omitzero" json:"-"`
	// Cashtags separated by spaces, commas, or lines.
	Cashtags param.Opt[string] `query:"cashtags,omitzero" json:"-"`
	// Conversation ID filter.
	ConversationID param.Opt[string] `query:"conversationId,omitzero" json:"-"`
	// Pagination cursor for collection results.
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
	MediaType XCommunityTweetListByCommunityParamsMediaType `query:"mediaType,omitzero" json:"-"`
	// Latest for newest first, or Top for the posts X ranks highest. Any case works.
	// Aliases ranking & sort, which also take Recency & Relevance. Send the same value
	// with each cursor.
	//
	// Any of "Latest", "Top".
	QueryType XCommunityTweetListByCommunityParamsQueryType `query:"queryType,omitzero" json:"-"`
	// Only when the caller requests a quote mode.
	//
	// Any of "include", "exclude", "only".
	Quotes XCommunityTweetListByCommunityParamsQuotes `query:"quotes,omitzero" json:"-"`
	// Reply mode. Exclude keeps own-thread replies.
	//
	// Any of "include", "exclude", "only".
	Replies XCommunityTweetListByCommunityParamsReplies `query:"replies,omitzero" json:"-"`
	// Only when the caller requests a repost mode.
	//
	// Any of "include", "exclude", "only".
	Retweets XCommunityTweetListByCommunityParamsRetweets `query:"retweets,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [XCommunityTweetListByCommunityParams]'s query parameters as
// `url.Values`.
func (r XCommunityTweetListByCommunityParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// XCommunityTweetListByCommunityParamsMediaType is the type of [XCommunityTweetListByCommunityParams.MediaType]. Filter media. Aliases: has_video, has_media.
type XCommunityTweetListByCommunityParamsMediaType string

const (
	XCommunityTweetListByCommunityParamsMediaTypeImages XCommunityTweetListByCommunityParamsMediaType = "images"
	XCommunityTweetListByCommunityParamsMediaTypeVideos XCommunityTweetListByCommunityParamsMediaType = "videos"
	XCommunityTweetListByCommunityParamsMediaTypeGifs   XCommunityTweetListByCommunityParamsMediaType = "gifs"
	XCommunityTweetListByCommunityParamsMediaTypeMedia  XCommunityTweetListByCommunityParamsMediaType = "media"
	XCommunityTweetListByCommunityParamsMediaTypeLinks  XCommunityTweetListByCommunityParamsMediaType = "links"
	XCommunityTweetListByCommunityParamsMediaTypeNone   XCommunityTweetListByCommunityParamsMediaType = "none"
)

// XCommunityTweetListParamsReplies is a reply mode. Exclude keeps own-thread replies.
type XCommunityTweetListParamsReplies string

const (
	XCommunityTweetListParamsRepliesInclude XCommunityTweetListParamsReplies = "include"
	XCommunityTweetListParamsRepliesExclude XCommunityTweetListParamsReplies = "exclude"
	XCommunityTweetListParamsRepliesOnly    XCommunityTweetListParamsReplies = "only"
)

// XCommunityTweetListParamsRetweets applies only when the caller requests a repost mode.
type XCommunityTweetListParamsRetweets string

const (
	XCommunityTweetListParamsRetweetsInclude XCommunityTweetListParamsRetweets = "include"
	XCommunityTweetListParamsRetweetsExclude XCommunityTweetListParamsRetweets = "exclude"
	XCommunityTweetListParamsRetweetsOnly    XCommunityTweetListParamsRetweets = "only"
)

// XCommunityTweetListParamsQuotes applies only when the caller requests a quote mode.
type XCommunityTweetListParamsQuotes string

const (
	XCommunityTweetListParamsQuotesInclude XCommunityTweetListParamsQuotes = "include"
	XCommunityTweetListParamsQuotesExclude XCommunityTweetListParamsQuotes = "exclude"
	XCommunityTweetListParamsQuotesOnly    XCommunityTweetListParamsQuotes = "only"
)

// XCommunityTweetListByCommunityParamsQueryType is the type of [XCommunityTweetListByCommunityParams.QueryType]. Latest for newest first, or Top for the posts X ranks highest. Any case works.
// Aliases ranking & sort, which also take Recency & Relevance. Send the same value
// with each cursor.
type XCommunityTweetListByCommunityParamsQueryType string

const (
	XCommunityTweetListByCommunityParamsQueryTypeLatest XCommunityTweetListByCommunityParamsQueryType = "Latest"
	XCommunityTweetListByCommunityParamsQueryTypeTop    XCommunityTweetListByCommunityParamsQueryType = "Top"
)

// XCommunityTweetListByCommunityParamsReplies is a reply mode. Exclude keeps own-thread replies.
type XCommunityTweetListByCommunityParamsReplies string

const (
	XCommunityTweetListByCommunityParamsRepliesInclude XCommunityTweetListByCommunityParamsReplies = "include"
	XCommunityTweetListByCommunityParamsRepliesExclude XCommunityTweetListByCommunityParamsReplies = "exclude"
	XCommunityTweetListByCommunityParamsRepliesOnly    XCommunityTweetListByCommunityParamsReplies = "only"
)

// XCommunityTweetListByCommunityParamsRetweets applies only when the caller requests a repost mode.
type XCommunityTweetListByCommunityParamsRetweets string

const (
	XCommunityTweetListByCommunityParamsRetweetsInclude XCommunityTweetListByCommunityParamsRetweets = "include"
	XCommunityTweetListByCommunityParamsRetweetsExclude XCommunityTweetListByCommunityParamsRetweets = "exclude"
	XCommunityTweetListByCommunityParamsRetweetsOnly    XCommunityTweetListByCommunityParamsRetweets = "only"
)

// XCommunityTweetListByCommunityParamsQuotes applies only when the caller requests a quote mode.
type XCommunityTweetListByCommunityParamsQuotes string

const (
	XCommunityTweetListByCommunityParamsQuotesInclude XCommunityTweetListByCommunityParamsQuotes = "include"
	XCommunityTweetListByCommunityParamsQuotesExclude XCommunityTweetListByCommunityParamsQuotes = "exclude"
	XCommunityTweetListByCommunityParamsQuotesOnly    XCommunityTweetListByCommunityParamsQuotes = "only"
)
