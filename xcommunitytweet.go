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

// Requires a Community ID and keyword query.
func (r *XCommunityTweetService) List(ctx context.Context, query XCommunityTweetListParams, opts ...option.RequestOption) (res *shared.PaginatedTweets, err error) {
	opts = slices.Concat(r.options, opts)
	path := "x/communities/tweets"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// List tweets posted in a community
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
	// Numeric ID of the community to search
	CommunityID string `query:"communityId" api:"required" json:"-"`
	// Keyword query within the selected community
	Q string `query:"q" api:"required" json:"-"`
	// Pagination cursor for community results
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Maximum items requested from this page (1-100, default 20). The response can
	// contain fewer items because the source returned fewer, filters removed items, or
	// remaining credits cover fewer results. Keep requesting next_cursor while
	// has_next_page is true, even when a page is empty. The deprecated limit and count
	// aliases remain accepted.
	PageSize param.Opt[int64] `query:"pageSize,omitzero" json:"-"`
	// Sort order for community results (Latest or Top)
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

// Sort order for community results (Latest or Top)
type XCommunityTweetListParamsQueryType string

const (
	XCommunityTweetListParamsQueryTypeLatest XCommunityTweetListParamsQueryType = "Latest"
	XCommunityTweetListParamsQueryTypeTop    XCommunityTweetListParamsQueryType = "Top"
)

type XCommunityTweetListByCommunityParams struct {
	// Pagination cursor for community tweets
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Maximum items requested from this page (1-100, default 20). The response can
	// contain fewer items because the source returned fewer, filters removed items, or
	// remaining credits cover fewer results. Keep requesting next_cursor while
	// has_next_page is true, even when a page is empty. The deprecated limit and count
	// aliases remain accepted.
	PageSize param.Opt[int64] `query:"pageSize,omitzero" json:"-"`
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
