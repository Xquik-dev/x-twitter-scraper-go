// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

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

// X Community info, members, and tweets
//
// XCommunityTweetService contains methods and other services that help with
// interacting with the x-twitter-scraper API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewXCommunityTweetService] method instead.
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

// One resumable page. Requires a Community ID and query.
func (r *XCommunityTweetService) List(ctx context.Context, query XCommunityTweetListParams, opts ...option.RequestOption) (res *shared.PaginatedTweets, err error) {
	opts = slices.Concat(r.options, opts)
	path := "x/communities/tweets"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Returns public tweets posted within one community.
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
	// Pagination cursor for community search
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Filter by language. Alias `lang` is accepted.
	Language param.Opt[string] `query:"language,omitzero" json:"-"`
	// Minimum likes. Aliases: minFaves, min_likes, min_faves.
	MinLikes param.Opt[int64] `query:"minLikes,omitzero" json:"-"`
	// Minimum replies threshold.
	MinReplies param.Opt[int64] `query:"minReplies,omitzero" json:"-"`
	// Minimum retweets threshold.
	MinRetweets param.Opt[int64] `query:"minRetweets,omitzero" json:"-"`
	// Minimum view count threshold.
	MinViews param.Opt[int64] `query:"minViews,omitzero" json:"-"`
	// Maximum page items (1-100, default 20). Source, filters, or credits can reduce
	// results. Follow next_cursor while the response reports more pages. Deprecated
	// limit and count aliases remain accepted.
	PageSize param.Opt[int64] `query:"pageSize,omitzero" json:"-"`
	// Start date in YYYY-MM-DD format.
	SinceDate param.Opt[time.Time] `query:"sinceDate,omitzero" format:"date" json:"-"`
	// End date in YYYY-MM-DD format.
	UntilDate param.Opt[time.Time] `query:"untilDate,omitzero" format:"date" json:"-"`
	// Only return tweets from verified authors.
	VerifiedOnly param.Opt[bool] `query:"verifiedOnly,omitzero" json:"-"`
	// Filter media. Aliases: has_video, has_media.
	//
	// Any of "images", "videos", "gifs", "media", "links", "none".
	MediaType XCommunityTweetListParamsMediaType `query:"mediaType,omitzero" json:"-"`
	// Sort order (Latest or Top)
	//
	// Any of "Latest", "Top".
	QueryType XCommunityTweetListParamsQueryType `query:"queryType,omitzero" json:"-"`
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

// Filter media. Aliases: has_video, has_media.
type XCommunityTweetListParamsMediaType string

const (
	XCommunityTweetListParamsMediaTypeImages XCommunityTweetListParamsMediaType = "images"
	XCommunityTweetListParamsMediaTypeVideos XCommunityTweetListParamsMediaType = "videos"
	XCommunityTweetListParamsMediaTypeGifs   XCommunityTweetListParamsMediaType = "gifs"
	XCommunityTweetListParamsMediaTypeMedia  XCommunityTweetListParamsMediaType = "media"
	XCommunityTweetListParamsMediaTypeLinks  XCommunityTweetListParamsMediaType = "links"
	XCommunityTweetListParamsMediaTypeNone   XCommunityTweetListParamsMediaType = "none"
)

// Sort order (Latest or Top)
type XCommunityTweetListParamsQueryType string

const (
	XCommunityTweetListParamsQueryTypeLatest XCommunityTweetListParamsQueryType = "Latest"
	XCommunityTweetListParamsQueryTypeTop    XCommunityTweetListParamsQueryType = "Top"
)

type XCommunityTweetListByCommunityParams struct {
	// Pagination cursor for collection results.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Filter by language. Alias `lang` is accepted.
	Language param.Opt[string] `query:"language,omitzero" json:"-"`
	// Minimum likes. Aliases: minFaves, min_likes, min_faves.
	MinLikes param.Opt[int64] `query:"minLikes,omitzero" json:"-"`
	// Minimum replies threshold.
	MinReplies param.Opt[int64] `query:"minReplies,omitzero" json:"-"`
	// Minimum retweets threshold.
	MinRetweets param.Opt[int64] `query:"minRetweets,omitzero" json:"-"`
	// Minimum view count threshold.
	MinViews param.Opt[int64] `query:"minViews,omitzero" json:"-"`
	// Maximum page items (1-100, default 20). Source, filters, or credits can reduce
	// results. Follow next_cursor while the response reports more pages. Deprecated
	// limit and count aliases remain accepted.
	PageSize param.Opt[int64] `query:"pageSize,omitzero" json:"-"`
	// Start date in YYYY-MM-DD format.
	SinceDate param.Opt[time.Time] `query:"sinceDate,omitzero" format:"date" json:"-"`
	// End date in YYYY-MM-DD format.
	UntilDate param.Opt[time.Time] `query:"untilDate,omitzero" format:"date" json:"-"`
	// Only return tweets from verified authors.
	VerifiedOnly param.Opt[bool] `query:"verifiedOnly,omitzero" json:"-"`
	// Filter media. Aliases: has_video, has_media.
	//
	// Any of "images", "videos", "gifs", "media", "links", "none".
	MediaType XCommunityTweetListByCommunityParamsMediaType `query:"mediaType,omitzero" json:"-"`
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

// Filter media. Aliases: has_video, has_media.
type XCommunityTweetListByCommunityParamsMediaType string

const (
	XCommunityTweetListByCommunityParamsMediaTypeImages XCommunityTweetListByCommunityParamsMediaType = "images"
	XCommunityTweetListByCommunityParamsMediaTypeVideos XCommunityTweetListByCommunityParamsMediaType = "videos"
	XCommunityTweetListByCommunityParamsMediaTypeGifs   XCommunityTweetListByCommunityParamsMediaType = "gifs"
	XCommunityTweetListByCommunityParamsMediaTypeMedia  XCommunityTweetListByCommunityParamsMediaType = "media"
	XCommunityTweetListByCommunityParamsMediaTypeLinks  XCommunityTweetListByCommunityParamsMediaType = "links"
	XCommunityTweetListByCommunityParamsMediaTypeNone   XCommunityTweetListByCommunityParamsMediaType = "none"
)
