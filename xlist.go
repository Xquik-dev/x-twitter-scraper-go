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

// X data lookups (subscription required)
//
// XListService contains methods and other services that help with interacting with
// the x-twitter-scraper API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewXListService] method instead.
type XListService struct {
	options []option.RequestOption
}

// NewXListService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewXListService(opts ...option.RequestOption) (r XListService) {
	r = XListService{}
	r.options = opts
	return
}

// Get list followers
func (r *XListService) GetFollowers(ctx context.Context, id string, query XListGetFollowersParams, opts ...option.RequestOption) (res *shared.PaginatedUsers, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("x/lists/%s/followers", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Get list members
func (r *XListService) GetMembers(ctx context.Context, id string, query XListGetMembersParams, opts ...option.RequestOption) (res *shared.PaginatedUsers, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("x/lists/%s/members", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Get list tweets
func (r *XListService) GetTweets(ctx context.Context, id string, query XListGetTweetsParams, opts ...option.RequestOption) (res *shared.PaginatedTweets, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("x/lists/%s/tweets", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type XListGetFollowersParams struct {
	// Pagination cursor for list followers
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [XListGetFollowersParams]'s query parameters as
// `url.Values`.
func (r XListGetFollowersParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type XListGetMembersParams struct {
	// Pagination cursor for list members
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [XListGetMembersParams]'s query parameters as `url.Values`.
func (r XListGetMembersParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type XListGetTweetsParams struct {
	// Pagination cursor for list tweets
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Include replies (default false)
	IncludeReplies param.Opt[bool] `query:"includeReplies,omitzero" json:"-"`
	// Unix timestamp - filter after
	SinceTime param.Opt[string] `query:"sinceTime,omitzero" json:"-"`
	// Unix timestamp - filter before
	UntilTime param.Opt[string] `query:"untilTime,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [XListGetTweetsParams]'s query parameters as `url.Values`.
func (r XListGetTweetsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
