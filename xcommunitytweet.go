// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package xtwitterscraper

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/stainless-sdks/x-twitter-scraper-go/internal/apiquery"
	"github.com/stainless-sdks/x-twitter-scraper-go/internal/requestconfig"
	"github.com/stainless-sdks/x-twitter-scraper-go/option"
	"github.com/stainless-sdks/x-twitter-scraper-go/packages/pagination"
	"github.com/stainless-sdks/x-twitter-scraper-go/packages/param"
	"github.com/stainless-sdks/x-twitter-scraper-go/shared"
)

// X data lookups (subscription required)
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

// Search tweets across all communities
func (r *XCommunityTweetService) List(ctx context.Context, query XCommunityTweetListParams, opts ...option.RequestOption) (res *pagination.CursorPage[shared.PaginatedTweets], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "x/communities/tweets"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Search tweets across all communities
func (r *XCommunityTweetService) ListAutoPaging(ctx context.Context, query XCommunityTweetListParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[shared.PaginatedTweets] {
	return pagination.NewCursorPageAutoPager(r.List(ctx, query, opts...))
}

// Get community tweets
func (r *XCommunityTweetService) ListByCommunity(ctx context.Context, id string, query XCommunityTweetListByCommunityParams, opts ...option.RequestOption) (res *pagination.CursorPage[shared.PaginatedTweets], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("x/communities/%s/tweets", url.PathEscape(id))
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Get community tweets
func (r *XCommunityTweetService) ListByCommunityAutoPaging(ctx context.Context, id string, query XCommunityTweetListByCommunityParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[shared.PaginatedTweets] {
	return pagination.NewCursorPageAutoPager(r.ListByCommunity(ctx, id, query, opts...))
}

type XCommunityTweetListParams struct {
	// Search query for cross-community tweets
	Q string `query:"q" api:"required" json:"-"`
	// Pagination cursor for cross-community results
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Sort order for cross-community results (Latest or Top)
	QueryType param.Opt[string] `query:"queryType,omitzero" json:"-"`
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

type XCommunityTweetListByCommunityParams struct {
	// Pagination cursor for community tweets
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
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
