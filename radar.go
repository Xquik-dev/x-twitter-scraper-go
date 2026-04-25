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

// Trending topic with score, category, source, region, language, and
// source-specific metadata.
type RadarItem struct {
	// Internal numeric identifier (stringified bigint).
	ID string `json:"id" api:"required"`
	// Any of "general", "tech", "dev", "science", "culture", "politics", "business",
	// "entertainment".
	Category  RadarItemCategory `json:"category" api:"required"`
	CreatedAt time.Time         `json:"createdAt" api:"required" format:"date-time"`
	Language  string            `json:"language" api:"required"`
	// Source-specific fields. Shape varies per source:
	//
	//   - reddit: { subreddit: string, author: string }
	//   - github: { starsToday: number }
	//   - hacker_news: { points: number, numberComments: number }
	//   - google_trends: { approxTraffic: number }
	//   - polymarket: { volume24hr: number }
	//   - wikipedia: { views: number }
	//   - trustmrr: { mrr, growthPercent, last30Days, total, customers,
	//     activeSubscriptions, onSale, xHandle?, category?, askingPrice?, country?,
	//     growthMrrPercent?, multiple?, paymentProvider?, rank? }
	Metadata    map[string]any `json:"metadata" api:"required"`
	PublishedAt time.Time      `json:"publishedAt" api:"required" format:"date-time"`
	Region      string         `json:"region" api:"required"`
	Score       float64        `json:"score" api:"required"`
	// Any of "github", "google_trends", "hacker_news", "polymarket", "reddit",
	// "trustmrr", "wikipedia".
	Source RadarItemSource `json:"source" api:"required"`
	// Source-specific identifier used for deduplication.
	SourceID    string `json:"sourceId" api:"required"`
	Title       string `json:"title" api:"required"`
	Description string `json:"description"`
	ImageURL    string `json:"imageUrl"`
	URL         string `json:"url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Category    respjson.Field
		CreatedAt   respjson.Field
		Language    respjson.Field
		Metadata    respjson.Field
		PublishedAt respjson.Field
		Region      respjson.Field
		Score       respjson.Field
		Source      respjson.Field
		SourceID    respjson.Field
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

type RadarItemCategory string

const (
	RadarItemCategoryGeneral       RadarItemCategory = "general"
	RadarItemCategoryTech          RadarItemCategory = "tech"
	RadarItemCategoryDev           RadarItemCategory = "dev"
	RadarItemCategoryScience       RadarItemCategory = "science"
	RadarItemCategoryCulture       RadarItemCategory = "culture"
	RadarItemCategoryPolitics      RadarItemCategory = "politics"
	RadarItemCategoryBusiness      RadarItemCategory = "business"
	RadarItemCategoryEntertainment RadarItemCategory = "entertainment"
)

type RadarItemSource string

const (
	RadarItemSourceGitHub       RadarItemSource = "github"
	RadarItemSourceGoogleTrends RadarItemSource = "google_trends"
	RadarItemSourceHackerNews   RadarItemSource = "hacker_news"
	RadarItemSourcePolymarket   RadarItemSource = "polymarket"
	RadarItemSourceReddit       RadarItemSource = "reddit"
	RadarItemSourceTrustmrr     RadarItemSource = "trustmrr"
	RadarItemSourceWikipedia    RadarItemSource = "wikipedia"
)

type RadarGetTrendingTopicsResponse struct {
	HasMore bool        `json:"hasMore" api:"required"`
	Items   []RadarItem `json:"items" api:"required"`
	// Opaque cursor for the next page (present only when hasMore is true).
	NextCursor string `json:"nextCursor"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		HasMore     respjson.Field
		Items       respjson.Field
		NextCursor  respjson.Field
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
	// Cursor for pagination (from prior response nextCursor).
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// Lookback window in hours (1-168, default 24).
	Hours param.Opt[int64] `query:"hours,omitzero" json:"-"`
	// Number of items to return (1-100, default 50).
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Region filter (us, global, etc.)
	Region param.Opt[string] `query:"region,omitzero" json:"-"`
	// Filter by category.
	//
	// Any of "general", "tech", "dev", "science", "culture", "politics", "business",
	// "entertainment".
	Category RadarGetTrendingTopicsParamsCategory `query:"category,omitzero" json:"-"`
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

// Filter by category.
type RadarGetTrendingTopicsParamsCategory string

const (
	RadarGetTrendingTopicsParamsCategoryGeneral       RadarGetTrendingTopicsParamsCategory = "general"
	RadarGetTrendingTopicsParamsCategoryTech          RadarGetTrendingTopicsParamsCategory = "tech"
	RadarGetTrendingTopicsParamsCategoryDev           RadarGetTrendingTopicsParamsCategory = "dev"
	RadarGetTrendingTopicsParamsCategoryScience       RadarGetTrendingTopicsParamsCategory = "science"
	RadarGetTrendingTopicsParamsCategoryCulture       RadarGetTrendingTopicsParamsCategory = "culture"
	RadarGetTrendingTopicsParamsCategoryPolitics      RadarGetTrendingTopicsParamsCategory = "politics"
	RadarGetTrendingTopicsParamsCategoryBusiness      RadarGetTrendingTopicsParamsCategory = "business"
	RadarGetTrendingTopicsParamsCategoryEntertainment RadarGetTrendingTopicsParamsCategory = "entertainment"
)

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
