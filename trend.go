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
	"github.com/Xquik-dev/x-twitter-scraper-go/packages/param"
	"github.com/Xquik-dev/x-twitter-scraper-go/packages/respjson"
)

// Trending topics and hashtags by region
//
// TrendService contains methods and other services that help with interacting with
// the x-twitter-scraper API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTrendService] method instead.
type TrendService struct {
	options []option.RequestOption
}

// NewTrendService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewTrendService(opts ...option.RequestOption) (r TrendService) {
	r = TrendService{}
	r.options = opts
	return
}

// Get trending hashtags and topics by region (alias)
func (r *TrendService) List(ctx context.Context, query TrendListParams, opts ...option.RequestOption) (res *TrendListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "trends"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type TrendListResponse struct {
	Total  int64                    `json:"total" api:"required"`
	Trends []TrendListResponseTrend `json:"trends" api:"required"`
	Woeid  int64                    `json:"woeid" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Total       respjson.Field
		Trends      respjson.Field
		Woeid       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TrendListResponse) RawJSON() string { return r.JSON.raw }
func (r *TrendListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TrendListResponseTrend struct {
	Name        string `json:"name" api:"required"`
	Description string `json:"description"`
	Query       string `json:"query"`
	Rank        int64  `json:"rank"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Description respjson.Field
		Query       respjson.Field
		Rank        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TrendListResponseTrend) RawJSON() string { return r.JSON.raw }
func (r *TrendListResponseTrend) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TrendListParams struct {
	// Number of trending topics to return (1-50, default 30)
	Count param.Opt[int64] `query:"count,omitzero" json:"-"`
	// Region WOEID (1=Worldwide, 23424977=US, 23424975=UK, 23424969=Turkey)
	Woeid param.Opt[int64] `query:"woeid,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [TrendListParams]'s query parameters as `url.Values`.
func (r TrendListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
