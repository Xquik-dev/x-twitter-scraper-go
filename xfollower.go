// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package xtwitterscraper

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/Xquik-dev/x-twitter-scraper-go/internal/apijson"
	"github.com/Xquik-dev/x-twitter-scraper-go/internal/apiquery"
	"github.com/Xquik-dev/x-twitter-scraper-go/internal/requestconfig"
	"github.com/Xquik-dev/x-twitter-scraper-go/option"
	"github.com/Xquik-dev/x-twitter-scraper-go/packages/respjson"
)

// Look up, search, and explore user profiles and relationships
//
// XFollowerService contains methods and other services that help with interacting
// with the x-twitter-scraper API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewXFollowerService] method instead.
type XFollowerService struct {
	options []option.RequestOption
}

// NewXFollowerService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewXFollowerService(opts ...option.RequestOption) (r XFollowerService) {
	r = XFollowerService{}
	r.options = opts
	return
}

// Check if one user follows another
func (r *XFollowerService) Check(ctx context.Context, query XFollowerCheckParams, opts ...option.RequestOption) (res *XFollowerCheckResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "x/followers/check"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type XFollowerCheckResponse struct {
	IsFollowedBy   bool   `json:"isFollowedBy" api:"required"`
	IsFollowing    bool   `json:"isFollowing" api:"required"`
	SourceUsername string `json:"sourceUsername" api:"required"`
	TargetUsername string `json:"targetUsername" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IsFollowedBy   respjson.Field
		IsFollowing    respjson.Field
		SourceUsername respjson.Field
		TargetUsername respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XFollowerCheckResponse) RawJSON() string { return r.JSON.raw }
func (r *XFollowerCheckResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type XFollowerCheckParams struct {
	// Username to check (without @)
	Source string `query:"source" api:"required" json:"-"`
	// Target username (without @)
	Target string `query:"target" api:"required" json:"-"`
	paramObj
}

// URLQuery serializes [XFollowerCheckParams]'s query parameters as `url.Values`.
func (r XFollowerCheckParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
