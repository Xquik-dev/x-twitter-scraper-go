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

// XCommunityMediaService lets you find X Communities & read their info, members, posts & media.
//
// XCommunityMediaService contains methods and other services that help with
// interacting with the x-twitter-scraper API.
//
// You should not instantiate this service directly, and instead use the
// [NewXCommunityMediaService] method instead.
type XCommunityMediaService struct {
	options []option.RequestOption
}

// NewXCommunityMediaService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewXCommunityMediaService(opts ...option.RequestOption) (r XCommunityMediaService) {
	r = XCommunityMediaService{}
	r.options = opts
	return
}

// List returns the tweets with photos or videos on 1 community's Media tab, newest
// first. Each tweet costs 1 credit.
func (r *XCommunityMediaService) List(ctx context.Context, id string, query XCommunityMediaListParams, opts ...option.RequestOption) (res *shared.PaginatedTweets, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("x/communities/%s/media", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type XCommunityMediaListParams struct {
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
	MediaType XCommunityMediaListParamsMediaType `query:"mediaType,omitzero" json:"-"`
	// Only when the caller requests a quote mode.
	//
	// Any of "include", "exclude", "only".
	Quotes XCommunityMediaListParamsQuotes `query:"quotes,omitzero" json:"-"`
	// Reply mode. Exclude keeps own-thread replies.
	//
	// Any of "include", "exclude", "only".
	Replies XCommunityMediaListParamsReplies `query:"replies,omitzero" json:"-"`
	// Only when the caller requests a repost mode.
	//
	// Any of "include", "exclude", "only".
	Retweets XCommunityMediaListParamsRetweets `query:"retweets,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [XCommunityMediaListParams]'s query parameters as `url.Values`.
func (r XCommunityMediaListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// XCommunityMediaListParamsMediaType is the type of [XCommunityMediaListParams.MediaType]. Filter media. Its aliases are has_video & has_media.
type XCommunityMediaListParamsMediaType string

const (
	XCommunityMediaListParamsMediaTypeImages XCommunityMediaListParamsMediaType = "images"
	XCommunityMediaListParamsMediaTypeVideos XCommunityMediaListParamsMediaType = "videos"
	XCommunityMediaListParamsMediaTypeGifs   XCommunityMediaListParamsMediaType = "gifs"
	XCommunityMediaListParamsMediaTypeMedia  XCommunityMediaListParamsMediaType = "media"
	XCommunityMediaListParamsMediaTypeLinks  XCommunityMediaListParamsMediaType = "links"
	XCommunityMediaListParamsMediaTypeNone   XCommunityMediaListParamsMediaType = "none"
)

// XCommunityMediaListParamsQuotes applies only when the caller requests a quote mode.
type XCommunityMediaListParamsQuotes string

const (
	XCommunityMediaListParamsQuotesInclude XCommunityMediaListParamsQuotes = "include"
	XCommunityMediaListParamsQuotesExclude XCommunityMediaListParamsQuotes = "exclude"
	XCommunityMediaListParamsQuotesOnly    XCommunityMediaListParamsQuotes = "only"
)

// XCommunityMediaListParamsReplies is a reply mode. Exclude keeps own-thread replies.
type XCommunityMediaListParamsReplies string

const (
	XCommunityMediaListParamsRepliesInclude XCommunityMediaListParamsReplies = "include"
	XCommunityMediaListParamsRepliesExclude XCommunityMediaListParamsReplies = "exclude"
	XCommunityMediaListParamsRepliesOnly    XCommunityMediaListParamsReplies = "only"
)

// XCommunityMediaListParamsRetweets applies only when the caller requests a repost mode.
type XCommunityMediaListParamsRetweets string

const (
	XCommunityMediaListParamsRetweetsInclude XCommunityMediaListParamsRetweets = "include"
	XCommunityMediaListParamsRetweetsExclude XCommunityMediaListParamsRetweets = "exclude"
	XCommunityMediaListParamsRetweetsOnly    XCommunityMediaListParamsRetweets = "only"
)
