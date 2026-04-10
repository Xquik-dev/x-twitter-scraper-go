// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package xtwitterscraper

import (
	"context"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/stainless-sdks/x-twitter-scraper-go/internal/apijson"
	"github.com/stainless-sdks/x-twitter-scraper-go/internal/apiquery"
	"github.com/stainless-sdks/x-twitter-scraper-go/internal/requestconfig"
	"github.com/stainless-sdks/x-twitter-scraper-go/option"
	"github.com/stainless-sdks/x-twitter-scraper-go/packages/param"
	"github.com/stainless-sdks/x-twitter-scraper-go/packages/respjson"
)

// AI tweet composition, drafts, writing styles, and radar
//
// RadarService contains methods and other services that help with interacting with
// the x-twitter-scraper API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRadarService] method instead.
type RadarService struct {
	options []option.RequestOption
}

// NewRadarService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewRadarService(opts ...option.RequestOption) (r RadarService) {
	r = RadarService{}
	r.options = opts
	return
}

// Get trending topics from curated sources
func (r *RadarService) GetTrendingTopics(ctx context.Context, query RadarGetTrendingTopicsParams, opts ...option.RequestOption) (res *RadarGetTrendingTopicsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "radar"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Trending topic with score, category, source, and region.
type RadarItem struct {
	Category    string    `json:"category" api:"required"`
	PublishedAt time.Time `json:"publishedAt" api:"required" format:"date-time"`
	Region      string    `json:"region" api:"required"`
	Score       float64   `json:"score" api:"required"`
	Source      string    `json:"source" api:"required"`
	Title       string    `json:"title" api:"required"`
	Description string    `json:"description"`
	ImageURL    string    `json:"imageUrl"`
	URL         string    `json:"url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Category    respjson.Field
		PublishedAt respjson.Field
		Region      respjson.Field
		Score       respjson.Field
		Source      respjson.Field
		Title       respjson.Field
		Description respjson.Field
		ImageURL    respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RadarItem) RawJSON() string { return r.JSON.raw }
func (r *RadarItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RadarGetTrendingTopicsResponse struct {
	Items []RadarItem `json:"items" api:"required"`
	Total int64       `json:"total" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		Total       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RadarGetTrendingTopicsResponse) RawJSON() string { return r.JSON.raw }
func (r *RadarGetTrendingTopicsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RadarGetTrendingTopicsParams struct {
	// Filter by category (general, tech, dev, etc.)
	Category param.Opt[string] `query:"category,omitzero" json:"-"`
	// Number of items to return
	Count param.Opt[int64] `query:"count,omitzero" json:"-"`
	// Lookback window in hours
	Hours param.Opt[int64] `query:"hours,omitzero" json:"-"`
	// Region filter (us, global, etc.)
	Region param.Opt[string] `query:"region,omitzero" json:"-"`
	// Source filter. One of: github, google_trends, hacker_news, polymarket, reddit,
	// trustmrr, wikipedia
	//
	// Any of "github", "google_trends", "hacker_news", "polymarket", "reddit",
	// "trustmrr", "wikipedia".
	Source RadarGetTrendingTopicsParamsSource `query:"source,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [RadarGetTrendingTopicsParams]'s query parameters as
// `url.Values`.
func (r RadarGetTrendingTopicsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Source filter. One of: github, google_trends, hacker_news, polymarket, reddit,
// trustmrr, wikipedia
type RadarGetTrendingTopicsParamsSource string

const (
	RadarGetTrendingTopicsParamsSourceGitHub       RadarGetTrendingTopicsParamsSource = "github"
	RadarGetTrendingTopicsParamsSourceGoogleTrends RadarGetTrendingTopicsParamsSource = "google_trends"
	RadarGetTrendingTopicsParamsSourceHackerNews   RadarGetTrendingTopicsParamsSource = "hacker_news"
	RadarGetTrendingTopicsParamsSourcePolymarket   RadarGetTrendingTopicsParamsSource = "polymarket"
	RadarGetTrendingTopicsParamsSourceReddit       RadarGetTrendingTopicsParamsSource = "reddit"
	RadarGetTrendingTopicsParamsSourceTrustmrr     RadarGetTrendingTopicsParamsSource = "trustmrr"
	RadarGetTrendingTopicsParamsSourceWikipedia    RadarGetTrendingTopicsParamsSource = "wikipedia"
)
