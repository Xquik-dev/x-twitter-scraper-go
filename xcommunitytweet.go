// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package xtwitterscraper

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/Xquik-dev/x-twitter-scraper-go/internal/apiquery"
	"github.com/Xquik-dev/x-twitter-scraper-go/internal/requestconfig"
	"github.com/Xquik-dev/x-twitter-scraper-go/option"
	"github.com/Xquik-dev/x-twitter-scraper-go/packages/param"
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
func (r *XCommunityTweetService) List(ctx context.Context, query XCommunityTweetListParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "x/communities/tweets"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, nil, opts...)
	return err
}

type XCommunityTweetListParams struct {
	// Search query
	Q string `query:"q" api:"required" json:"-"`
	// Pagination cursor
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Sort order (Latest or Top)
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
