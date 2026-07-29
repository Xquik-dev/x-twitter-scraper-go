// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package xtwitterscraper

import (
	"context"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/Xquik-dev/x-twitter-scraper-go/internal/apijson"
	"github.com/Xquik-dev/x-twitter-scraper-go/internal/apiquery"
	"github.com/Xquik-dev/x-twitter-scraper-go/internal/requestconfig"
	"github.com/Xquik-dev/x-twitter-scraper-go/option"
	"github.com/Xquik-dev/x-twitter-scraper-go/packages/param"
	"github.com/Xquik-dev/x-twitter-scraper-go/packages/respjson"
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
	// Radar item identifier.
	ID string `json:"id" api:"required"`
	// Any of "general", "tech", "dev", "science", "culture", "politics", "business",
	// "entertainment".
	Category  RadarItemCategory `json:"category" api:"required"`
	CreatedAt time.Time         `json:"createdAt" api:"required" format:"date-time"`
	// BCP-47 language code. und means the source did not identify a language.
	Language string `json:"language" api:"required"`
	// Source-specific fields. Shape varies per source:
	//
	//   - reddit: { author, authorId?, subreddit, subredditId?, subredditSubscribers?,
	//     sourceFormat, score?, upvoteRatio?, estimatedUpvotes?, estimatedDownvotes?,
	//     numberComments?, numberCrossposts?, selftext?, contentUrl?, domain?,
	//     postHint?, linkFlairText?, distinguished?, totalAwardsReceived?, viewCount?,
	//     editedAt?, galleryImageUrls?, redditVideo?, archived?, contestMode?,
	//     isCrosspostable?, isMeta?, isNsfw?, isOriginalContent?, isRobotIndexable?,
	//     isSelf?, isSpoiler?, isVideo?, locked?, stickied? }. `score` is Reddit's
	//     public net score. Exact public upvote and downvote counts are not available.
	//     Estimated counts derive from the public score and upvote ratio, which Reddit
	//     may fuzz. Comment bodies are not included. Current items combine public
	//     listing discovery with server-rendered post data and use `sourceFormat: html`;
	//     `json` and `rss` remain for legacy rows.
	//   - github: { starsToday: number }
	//   - hacker_news: { points: number, numberComments: number }
	//   - google_trends: { approxTraffic: number }
	//   - polymarket: { volume24hr: number }
	//   - wikipedia: { views: number }
	//   - trustmrr: { mrr, growthPercent, last30Days, total, customers,
	//     activeSubscriptions, onSale, xHandle?, category?, askingPrice?, country?,
	//     foundedDate?, googleSearchImpressionsLast30Days?, growthMrrPercent?,
	//     multiple?, paymentProvider?, profitMarginLast30Days?, rank?,
	//     revenuePerVisitor?, targetAudience?, visitorsLast30Days? } For the startup
	//     growth source, xHandle is the founder's X username without @. The rank field
	//     is the source's revenue rank. Result order represents reported 30-day
	//     revenue-growth rank.
	Metadata    RadarItemMetadata `json:"metadata" api:"required"`
	PublishedAt time.Time         `json:"publishedAt" api:"required" format:"date-time"`
	Region      string            `json:"region" api:"required"`
	Score       float64           `json:"score" api:"required"`
	// Any of "github", "google_trends", "hacker_news", "polymarket", "reddit",
	// "trustmrr", "wikipedia".
	Source RadarItemSource `json:"source" api:"required"`
	// Source-specific identifier used for deduplication.
	SourceID    string `json:"sourceId" api:"required"`
	Title       string `json:"title" api:"required"`
	Description string `json:"description"`
	// Source image. Startup growth items return the logo here.
	ImageURL string `json:"imageUrl"`
	URL      string `json:"url"`
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

// Source-specific fields. Shape varies per source:
//
//   - reddit: { author, authorId?, subreddit, subredditId?, subredditSubscribers?,
//     sourceFormat, score?, upvoteRatio?, estimatedUpvotes?, estimatedDownvotes?,
//     numberComments?, numberCrossposts?, selftext?, contentUrl?, domain?,
//     postHint?, linkFlairText?, distinguished?, totalAwardsReceived?, viewCount?,
//     editedAt?, galleryImageUrls?, redditVideo?, archived?, contestMode?,
//     isCrosspostable?, isMeta?, isNsfw?, isOriginalContent?, isRobotIndexable?,
//     isSelf?, isSpoiler?, isVideo?, locked?, stickied? }. `score` is Reddit's
//     public net score. Exact public upvote and downvote counts are not available.
//     Estimated counts derive from the public score and upvote ratio, which Reddit
//     may fuzz. Comment bodies are not included. Current items combine public
//     listing discovery with server-rendered post data and use `sourceFormat: html`;
//     `json` and `rss` remain for legacy rows.
//   - github: { starsToday: number }
//   - hacker_news: { points: number, numberComments: number }
//   - google_trends: { approxTraffic: number }
//   - polymarket: { volume24hr: number }
//   - wikipedia: { views: number }
//   - trustmrr: { mrr, growthPercent, last30Days, total, customers,
//     activeSubscriptions, onSale, xHandle?, category?, askingPrice?, country?,
//     foundedDate?, googleSearchImpressionsLast30Days?, growthMrrPercent?,
//     multiple?, paymentProvider?, profitMarginLast30Days?, rank?,
//     revenuePerVisitor?, targetAudience?, visitorsLast30Days? } For the startup
//     growth source, xHandle is the founder's X username without @. The rank field
//     is the source's revenue rank. Result order represents reported 30-day
//     revenue-growth rank.
type RadarItemMetadata struct {
	Author             string `json:"author"`
	ContentURL         string `json:"contentUrl" format:"uri"`
	EstimatedDownvotes int64  `json:"estimatedDownvotes"`
	EstimatedUpvotes   int64  `json:"estimatedUpvotes"`
	NumberComments     int64  `json:"numberComments"`
	Score              int64  `json:"score"`
	Selftext           string `json:"selftext"`
	// Current items use html. json and rss are retained for legacy rows.
	//
	// Any of "html", "json", "rss".
	SourceFormat string         `json:"sourceFormat"`
	Subreddit    string         `json:"subreddit"`
	UpvoteRatio  float64        `json:"upvoteRatio"`
	ExtraFields  map[string]any `json:"" api:"extrafields"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Author             respjson.Field
		ContentURL         respjson.Field
		EstimatedDownvotes respjson.Field
		EstimatedUpvotes   respjson.Field
		NumberComments     respjson.Field
		Score              respjson.Field
		Selftext           respjson.Field
		SourceFormat       respjson.Field
		Subreddit          respjson.Field
		UpvoteRatio        respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RadarItemMetadata) RawJSON() string { return r.JSON.raw }
func (r *RadarItemMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

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
	// Lookback window in hours (1-72, default 6).
	Hours param.Opt[int64] `query:"hours,omitzero" json:"-"`
	// Number of items to return (1-100, default 50).
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Region filter. Use `global` or a region code such as `US`, `GB`, `TR`, or `ES`.
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
