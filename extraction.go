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

	"github.com/Xquik-dev/x-twitter-scraper-go/internal/apijson"
	"github.com/Xquik-dev/x-twitter-scraper-go/internal/apiquery"
	"github.com/Xquik-dev/x-twitter-scraper-go/internal/requestconfig"
	"github.com/Xquik-dev/x-twitter-scraper-go/option"
	"github.com/Xquik-dev/x-twitter-scraper-go/packages/param"
	"github.com/Xquik-dev/x-twitter-scraper-go/packages/respjson"
	"github.com/Xquik-dev/x-twitter-scraper-go/shared/constant"
)

// Bulk data extraction (23 tool types)
//
// ExtractionService contains methods and other services that help with interacting
// with the x-twitter-scraper API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewExtractionService] method instead.
type ExtractionService struct {
	options []option.RequestOption
}

// NewExtractionService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewExtractionService(opts ...option.RequestOption) (r ExtractionService) {
	r = ExtractionService{}
	r.options = opts
	return
}

// Get extraction results
func (r *ExtractionService) Get(ctx context.Context, id string, query ExtractionGetParams, opts ...option.RequestOption) (res *ExtractionGetResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("extractions/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// List extraction jobs
func (r *ExtractionService) List(ctx context.Context, query ExtractionListParams, opts ...option.RequestOption) (res *ExtractionListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "extractions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Estimate extraction cost
func (r *ExtractionService) EstimateCost(ctx context.Context, body ExtractionEstimateCostParams, opts ...option.RequestOption) (res *ExtractionEstimateCostResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "extractions/estimate"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Export extraction results
func (r *ExtractionService) ExportResults(ctx context.Context, id string, query ExtractionExportResultsParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/octet-stream")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("extractions/%s/export", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Run extraction
func (r *ExtractionService) Run(ctx context.Context, body ExtractionRunParams, opts ...option.RequestOption) (res *ExtractionRunResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "extractions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Extraction job tracking status, tool type, and result count.
type ExtractionJob struct {
	ID        string    `json:"id" api:"required"`
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Any of "running", "completed", "failed".
	Status ExtractionJobStatus `json:"status" api:"required"`
	// Identifier for the extraction tool used to run a job.
	//
	// Any of "article_extractor", "community_extractor",
	// "community_moderator_explorer", "community_post_extractor", "community_search",
	// "favoriters", "follower_explorer", "following_explorer",
	// "list_follower_explorer", "list_member_extractor", "list_post_extractor",
	// "mention_extractor", "people_search", "post_extractor", "quote_extractor",
	// "reply_extractor", "repost_extractor", "space_explorer", "thread_extractor",
	// "tweet_search_extractor", "user_likes", "user_media",
	// "verified_follower_explorer".
	ToolType     ExtractionJobToolType `json:"toolType" api:"required"`
	TotalResults int64                 `json:"totalResults" api:"required"`
	CompletedAt  time.Time             `json:"completedAt" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		CreatedAt    respjson.Field
		Status       respjson.Field
		ToolType     respjson.Field
		TotalResults respjson.Field
		CompletedAt  respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExtractionJob) RawJSON() string { return r.JSON.raw }
func (r *ExtractionJob) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExtractionJobStatus string

const (
	ExtractionJobStatusRunning   ExtractionJobStatus = "running"
	ExtractionJobStatusCompleted ExtractionJobStatus = "completed"
	ExtractionJobStatusFailed    ExtractionJobStatus = "failed"
)

// Identifier for the extraction tool used to run a job.
type ExtractionJobToolType string

const (
	ExtractionJobToolTypeArticleExtractor           ExtractionJobToolType = "article_extractor"
	ExtractionJobToolTypeCommunityExtractor         ExtractionJobToolType = "community_extractor"
	ExtractionJobToolTypeCommunityModeratorExplorer ExtractionJobToolType = "community_moderator_explorer"
	ExtractionJobToolTypeCommunityPostExtractor     ExtractionJobToolType = "community_post_extractor"
	ExtractionJobToolTypeCommunitySearch            ExtractionJobToolType = "community_search"
	ExtractionJobToolTypeFavoriters                 ExtractionJobToolType = "favoriters"
	ExtractionJobToolTypeFollowerExplorer           ExtractionJobToolType = "follower_explorer"
	ExtractionJobToolTypeFollowingExplorer          ExtractionJobToolType = "following_explorer"
	ExtractionJobToolTypeListFollowerExplorer       ExtractionJobToolType = "list_follower_explorer"
	ExtractionJobToolTypeListMemberExtractor        ExtractionJobToolType = "list_member_extractor"
	ExtractionJobToolTypeListPostExtractor          ExtractionJobToolType = "list_post_extractor"
	ExtractionJobToolTypeMentionExtractor           ExtractionJobToolType = "mention_extractor"
	ExtractionJobToolTypePeopleSearch               ExtractionJobToolType = "people_search"
	ExtractionJobToolTypePostExtractor              ExtractionJobToolType = "post_extractor"
	ExtractionJobToolTypeQuoteExtractor             ExtractionJobToolType = "quote_extractor"
	ExtractionJobToolTypeReplyExtractor             ExtractionJobToolType = "reply_extractor"
	ExtractionJobToolTypeRepostExtractor            ExtractionJobToolType = "repost_extractor"
	ExtractionJobToolTypeSpaceExplorer              ExtractionJobToolType = "space_explorer"
	ExtractionJobToolTypeThreadExtractor            ExtractionJobToolType = "thread_extractor"
	ExtractionJobToolTypeTweetSearchExtractor       ExtractionJobToolType = "tweet_search_extractor"
	ExtractionJobToolTypeUserLikes                  ExtractionJobToolType = "user_likes"
	ExtractionJobToolTypeUserMedia                  ExtractionJobToolType = "user_media"
	ExtractionJobToolTypeVerifiedFollowerExplorer   ExtractionJobToolType = "verified_follower_explorer"
)

type ExtractionGetResponse struct {
	HasMore bool `json:"hasMore" api:"required"`
	// Extraction job metadata - shape varies by tool type (JSON)
	Job        map[string]any   `json:"job" api:"required"`
	Results    []map[string]any `json:"results" api:"required"`
	NextCursor string           `json:"nextCursor"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		HasMore     respjson.Field
		Job         respjson.Field
		Results     respjson.Field
		NextCursor  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExtractionGetResponse) RawJSON() string { return r.JSON.raw }
func (r *ExtractionGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExtractionListResponse struct {
	Extractions []ExtractionJob `json:"extractions" api:"required"`
	HasMore     bool            `json:"hasMore" api:"required"`
	NextCursor  string          `json:"nextCursor"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Extractions respjson.Field
		HasMore     respjson.Field
		NextCursor  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExtractionListResponse) RawJSON() string { return r.JSON.raw }
func (r *ExtractionListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExtractionEstimateCostResponse struct {
	Allowed          bool   `json:"allowed" api:"required"`
	CreditsAvailable string `json:"creditsAvailable" api:"required"`
	CreditsRequired  string `json:"creditsRequired" api:"required"`
	EstimatedResults int64  `json:"estimatedResults" api:"required"`
	// Any of "followers", "following", "paginationCap", "posts", "quoteCount",
	// "replyCount", "resultsLimit", "retweetCount", "unknown".
	Source          ExtractionEstimateCostResponseSource `json:"source" api:"required"`
	ResolvedXUserID string                               `json:"resolvedXUserId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Allowed          respjson.Field
		CreditsAvailable respjson.Field
		CreditsRequired  respjson.Field
		EstimatedResults respjson.Field
		Source           respjson.Field
		ResolvedXUserID  respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExtractionEstimateCostResponse) RawJSON() string { return r.JSON.raw }
func (r *ExtractionEstimateCostResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExtractionEstimateCostResponseSource string

const (
	ExtractionEstimateCostResponseSourceFollowers     ExtractionEstimateCostResponseSource = "followers"
	ExtractionEstimateCostResponseSourceFollowing     ExtractionEstimateCostResponseSource = "following"
	ExtractionEstimateCostResponseSourcePaginationCap ExtractionEstimateCostResponseSource = "paginationCap"
	ExtractionEstimateCostResponseSourcePosts         ExtractionEstimateCostResponseSource = "posts"
	ExtractionEstimateCostResponseSourceQuoteCount    ExtractionEstimateCostResponseSource = "quoteCount"
	ExtractionEstimateCostResponseSourceReplyCount    ExtractionEstimateCostResponseSource = "replyCount"
	ExtractionEstimateCostResponseSourceResultsLimit  ExtractionEstimateCostResponseSource = "resultsLimit"
	ExtractionEstimateCostResponseSourceRetweetCount  ExtractionEstimateCostResponseSource = "retweetCount"
	ExtractionEstimateCostResponseSourceUnknown       ExtractionEstimateCostResponseSource = "unknown"
)

type ExtractionRunResponse struct {
	ID     string           `json:"id" api:"required"`
	Status constant.Running `json:"status" default:"running"`
	// Identifier for the extraction tool used to run a job.
	//
	// Any of "article_extractor", "community_extractor",
	// "community_moderator_explorer", "community_post_extractor", "community_search",
	// "favoriters", "follower_explorer", "following_explorer",
	// "list_follower_explorer", "list_member_extractor", "list_post_extractor",
	// "mention_extractor", "people_search", "post_extractor", "quote_extractor",
	// "reply_extractor", "repost_extractor", "space_explorer", "thread_extractor",
	// "tweet_search_extractor", "user_likes", "user_media",
	// "verified_follower_explorer".
	ToolType ExtractionRunResponseToolType `json:"toolType" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Status      respjson.Field
		ToolType    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExtractionRunResponse) RawJSON() string { return r.JSON.raw }
func (r *ExtractionRunResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Identifier for the extraction tool used to run a job.
type ExtractionRunResponseToolType string

const (
	ExtractionRunResponseToolTypeArticleExtractor           ExtractionRunResponseToolType = "article_extractor"
	ExtractionRunResponseToolTypeCommunityExtractor         ExtractionRunResponseToolType = "community_extractor"
	ExtractionRunResponseToolTypeCommunityModeratorExplorer ExtractionRunResponseToolType = "community_moderator_explorer"
	ExtractionRunResponseToolTypeCommunityPostExtractor     ExtractionRunResponseToolType = "community_post_extractor"
	ExtractionRunResponseToolTypeCommunitySearch            ExtractionRunResponseToolType = "community_search"
	ExtractionRunResponseToolTypeFavoriters                 ExtractionRunResponseToolType = "favoriters"
	ExtractionRunResponseToolTypeFollowerExplorer           ExtractionRunResponseToolType = "follower_explorer"
	ExtractionRunResponseToolTypeFollowingExplorer          ExtractionRunResponseToolType = "following_explorer"
	ExtractionRunResponseToolTypeListFollowerExplorer       ExtractionRunResponseToolType = "list_follower_explorer"
	ExtractionRunResponseToolTypeListMemberExtractor        ExtractionRunResponseToolType = "list_member_extractor"
	ExtractionRunResponseToolTypeListPostExtractor          ExtractionRunResponseToolType = "list_post_extractor"
	ExtractionRunResponseToolTypeMentionExtractor           ExtractionRunResponseToolType = "mention_extractor"
	ExtractionRunResponseToolTypePeopleSearch               ExtractionRunResponseToolType = "people_search"
	ExtractionRunResponseToolTypePostExtractor              ExtractionRunResponseToolType = "post_extractor"
	ExtractionRunResponseToolTypeQuoteExtractor             ExtractionRunResponseToolType = "quote_extractor"
	ExtractionRunResponseToolTypeReplyExtractor             ExtractionRunResponseToolType = "reply_extractor"
	ExtractionRunResponseToolTypeRepostExtractor            ExtractionRunResponseToolType = "repost_extractor"
	ExtractionRunResponseToolTypeSpaceExplorer              ExtractionRunResponseToolType = "space_explorer"
	ExtractionRunResponseToolTypeThreadExtractor            ExtractionRunResponseToolType = "thread_extractor"
	ExtractionRunResponseToolTypeTweetSearchExtractor       ExtractionRunResponseToolType = "tweet_search_extractor"
	ExtractionRunResponseToolTypeUserLikes                  ExtractionRunResponseToolType = "user_likes"
	ExtractionRunResponseToolTypeUserMedia                  ExtractionRunResponseToolType = "user_media"
	ExtractionRunResponseToolTypeVerifiedFollowerExplorer   ExtractionRunResponseToolType = "verified_follower_explorer"
)

type ExtractionGetParams struct {
	// Cursor for keyset pagination from prior response next_cursor
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Maximum number of results to return (1-1000, default 100)
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ExtractionGetParams]'s query parameters as `url.Values`.
func (r ExtractionGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ExtractionListParams struct {
	// Cursor for keyset pagination from prior response next_cursor
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Maximum number of items to return (1-100, default 50). For paid per-result
	// endpoints, the returned count may be lower when remaining credits cannot cover
	// the requested page. If zero paid results are affordable, the endpoint returns
	// 402 insufficient_credits.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Filter by job status
	//
	// Any of "running", "completed", "failed".
	Status ExtractionListParamsStatus `query:"status,omitzero" json:"-"`
	// Filter by extraction tool type
	//
	// Any of "article_extractor", "community_extractor",
	// "community_moderator_explorer", "community_post_extractor", "community_search",
	// "favoriters", "follower_explorer", "following_explorer",
	// "list_follower_explorer", "list_member_extractor", "list_post_extractor",
	// "mention_extractor", "people_search", "post_extractor", "quote_extractor",
	// "reply_extractor", "repost_extractor", "space_explorer", "thread_extractor",
	// "tweet_search_extractor", "user_likes", "user_media",
	// "verified_follower_explorer".
	ToolType ExtractionListParamsToolType `query:"toolType,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ExtractionListParams]'s query parameters as `url.Values`.
func (r ExtractionListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by job status
type ExtractionListParamsStatus string

const (
	ExtractionListParamsStatusRunning   ExtractionListParamsStatus = "running"
	ExtractionListParamsStatusCompleted ExtractionListParamsStatus = "completed"
	ExtractionListParamsStatusFailed    ExtractionListParamsStatus = "failed"
)

// Filter by extraction tool type
type ExtractionListParamsToolType string

const (
	ExtractionListParamsToolTypeArticleExtractor           ExtractionListParamsToolType = "article_extractor"
	ExtractionListParamsToolTypeCommunityExtractor         ExtractionListParamsToolType = "community_extractor"
	ExtractionListParamsToolTypeCommunityModeratorExplorer ExtractionListParamsToolType = "community_moderator_explorer"
	ExtractionListParamsToolTypeCommunityPostExtractor     ExtractionListParamsToolType = "community_post_extractor"
	ExtractionListParamsToolTypeCommunitySearch            ExtractionListParamsToolType = "community_search"
	ExtractionListParamsToolTypeFavoriters                 ExtractionListParamsToolType = "favoriters"
	ExtractionListParamsToolTypeFollowerExplorer           ExtractionListParamsToolType = "follower_explorer"
	ExtractionListParamsToolTypeFollowingExplorer          ExtractionListParamsToolType = "following_explorer"
	ExtractionListParamsToolTypeListFollowerExplorer       ExtractionListParamsToolType = "list_follower_explorer"
	ExtractionListParamsToolTypeListMemberExtractor        ExtractionListParamsToolType = "list_member_extractor"
	ExtractionListParamsToolTypeListPostExtractor          ExtractionListParamsToolType = "list_post_extractor"
	ExtractionListParamsToolTypeMentionExtractor           ExtractionListParamsToolType = "mention_extractor"
	ExtractionListParamsToolTypePeopleSearch               ExtractionListParamsToolType = "people_search"
	ExtractionListParamsToolTypePostExtractor              ExtractionListParamsToolType = "post_extractor"
	ExtractionListParamsToolTypeQuoteExtractor             ExtractionListParamsToolType = "quote_extractor"
	ExtractionListParamsToolTypeReplyExtractor             ExtractionListParamsToolType = "reply_extractor"
	ExtractionListParamsToolTypeRepostExtractor            ExtractionListParamsToolType = "repost_extractor"
	ExtractionListParamsToolTypeSpaceExplorer              ExtractionListParamsToolType = "space_explorer"
	ExtractionListParamsToolTypeThreadExtractor            ExtractionListParamsToolType = "thread_extractor"
	ExtractionListParamsToolTypeTweetSearchExtractor       ExtractionListParamsToolType = "tweet_search_extractor"
	ExtractionListParamsToolTypeUserLikes                  ExtractionListParamsToolType = "user_likes"
	ExtractionListParamsToolTypeUserMedia                  ExtractionListParamsToolType = "user_media"
	ExtractionListParamsToolTypeVerifiedFollowerExplorer   ExtractionListParamsToolType = "verified_follower_explorer"
)

type ExtractionEstimateCostParams struct {
	// Identifier for the extraction tool used to run a job.
	//
	// Any of "article_extractor", "community_extractor",
	// "community_moderator_explorer", "community_post_extractor", "community_search",
	// "favoriters", "follower_explorer", "following_explorer",
	// "list_follower_explorer", "list_member_extractor", "list_post_extractor",
	// "mention_extractor", "people_search", "post_extractor", "quote_extractor",
	// "reply_extractor", "repost_extractor", "space_explorer", "thread_extractor",
	// "tweet_search_extractor", "user_likes", "user_media",
	// "verified_follower_explorer".
	ToolType ExtractionEstimateCostParamsToolType `json:"toolType,omitzero" api:"required"`
	// Raw advanced query string appended to the estimate (tweet_search_extractor)
	AdvancedQuery param.Opt[string] `json:"advancedQuery,omitzero"`
	// Alternative words or quoted phrases for estimated results. Separate with spaces,
	// commas, or lines.
	AnyWords param.Opt[string] `json:"anyWords,omitzero"`
	// Geo bounding box used for estimation, e.g. -74.1 40.6 -73.9 40.8
	// (tweet_search_extractor)
	BoundingBox param.Opt[string] `json:"boundingBox,omitzero"`
	// Cashtags applied to the estimate, separated by spaces, commas, or lines.
	Cashtags param.Opt[string] `json:"cashtags,omitzero"`
	// Conversation ID filter used for estimation (tweet_search_extractor)
	ConversationID param.Opt[string] `json:"conversationId,omitzero"`
	// Exact phrase filter for search estimation
	ExactPhrase param.Opt[string] `json:"exactPhrase,omitzero"`
	// Words or quoted phrases excluded from estimated results. Separate with spaces,
	// commas, or lines.
	ExcludeWords param.Opt[string] `json:"excludeWords,omitzero"`
	// Estimate only tweets from this author username (tweet_search_extractor)
	FromUser param.Opt[string] `json:"fromUser,omitzero"`
	// Hashtags applied to the estimate, separated by spaces, commas, or lines.
	Hashtags param.Opt[string] `json:"hashtags,omitzero"`
	// Estimate only replies to this tweet ID (tweet_search_extractor)
	InReplyToTweetID param.Opt[string] `json:"inReplyToTweetId,omitzero"`
	// Language code used for estimate filtering (tweet_search_extractor)
	Language param.Opt[string] `json:"language,omitzero"`
	// Estimate search results within this list ID (tweet_search_extractor)
	ListID param.Opt[string] `json:"listId,omitzero"`
	// Estimate tweets mentioning this username (tweet_search_extractor)
	Mentioning param.Opt[string] `json:"mentioning,omitzero"`
	// Minimum likes threshold for estimated results (tweet_search_extractor)
	MinFaves param.Opt[int64] `json:"minFaves,omitzero"`
	// Minimum quote count threshold for estimated results (tweet_search_extractor)
	MinQuotes param.Opt[int64] `json:"minQuotes,omitzero"`
	// Minimum replies threshold for estimated results (tweet_search_extractor)
	MinReplies param.Opt[int64] `json:"minReplies,omitzero"`
	// Minimum retweets threshold for estimated results (tweet_search_extractor)
	MinRetweets param.Opt[int64] `json:"minRetweets,omitzero"`
	// Estimate search results within this place ID (tweet_search_extractor)
	Place param.Opt[string] `json:"place,omitzero"`
	// Estimate search results within this country code (tweet_search_extractor)
	PlaceCountry param.Opt[string] `json:"placeCountry,omitzero"`
	// Geo point radius used for estimation, e.g. -73.99 40.73 25mi
	// (tweet_search_extractor)
	PointRadius param.Opt[string] `json:"pointRadius,omitzero"`
	// Estimate only quotes of this tweet ID (tweet_search_extractor)
	QuotesOfTweetID param.Opt[string] `json:"quotesOfTweetId,omitzero"`
	// Maximum number of results to estimate. When set, the estimate caps projected
	// results to this value.
	ResultsLimit param.Opt[int64] `json:"resultsLimit,omitzero"`
	// Estimate only retweets of this tweet ID (tweet_search_extractor)
	RetweetsOfTweetID param.Opt[string] `json:"retweetsOfTweetId,omitzero"`
	// Required for tweet_search_extractor & community_search.
	SearchQuery param.Opt[string] `json:"searchQuery,omitzero"`
	// Estimate start date in YYYY-MM-DD format (tweet_search_extractor)
	SinceDate param.Opt[time.Time] `json:"sinceDate,omitzero" format:"date"`
	// Required for community_post_extractor & community_search.
	TargetCommunityID param.Opt[string] `json:"targetCommunityId,omitzero"`
	// Required for list_follower_explorer, list_member_extractor &
	// list_post_extractor.
	TargetListID param.Opt[string] `json:"targetListId,omitzero"`
	// Required for space_explorer.
	TargetSpaceID  param.Opt[string] `json:"targetSpaceId,omitzero"`
	TargetTweetID  param.Opt[string] `json:"targetTweetId,omitzero"`
	TargetUsername param.Opt[string] `json:"targetUsername,omitzero"`
	// Estimate replies sent to this username (tweet_search_extractor)
	ToUser param.Opt[string] `json:"toUser,omitzero"`
	// Estimate end date in YYYY-MM-DD format (tweet_search_extractor)
	UntilDate param.Opt[time.Time] `json:"untilDate,omitzero" format:"date"`
	// URL substring or domain filter used for estimation (tweet_search_extractor)
	URL param.Opt[string] `json:"url,omitzero"`
	// Estimate only verified authors (tweet_search_extractor)
	VerifiedOnly param.Opt[bool] `json:"verifiedOnly,omitzero"`
	// Media type used for estimate filtering (tweet_search_extractor)
	//
	// Any of "images", "videos", "gifs", "media", "links", "none".
	MediaType ExtractionEstimateCostParamsMediaType `json:"mediaType,omitzero"`
	// Quote mode used for estimation (tweet_search_extractor)
	//
	// Any of "include", "exclude", "only".
	Quotes ExtractionEstimateCostParamsQuotes `json:"quotes,omitzero"`
	// Reply mode used for estimation (tweet_search_extractor)
	//
	// Any of "include", "exclude", "only".
	Replies ExtractionEstimateCostParamsReplies `json:"replies,omitzero"`
	// Retweet mode used for estimation (tweet_search_extractor)
	//
	// Any of "include", "exclude", "only".
	Retweets ExtractionEstimateCostParamsRetweets `json:"retweets,omitzero"`
	paramObj
}

func (r ExtractionEstimateCostParams) MarshalJSON() (data []byte, err error) {
	type shadow ExtractionEstimateCostParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractionEstimateCostParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Identifier for the extraction tool used to run a job.
type ExtractionEstimateCostParamsToolType string

const (
	ExtractionEstimateCostParamsToolTypeArticleExtractor           ExtractionEstimateCostParamsToolType = "article_extractor"
	ExtractionEstimateCostParamsToolTypeCommunityExtractor         ExtractionEstimateCostParamsToolType = "community_extractor"
	ExtractionEstimateCostParamsToolTypeCommunityModeratorExplorer ExtractionEstimateCostParamsToolType = "community_moderator_explorer"
	ExtractionEstimateCostParamsToolTypeCommunityPostExtractor     ExtractionEstimateCostParamsToolType = "community_post_extractor"
	ExtractionEstimateCostParamsToolTypeCommunitySearch            ExtractionEstimateCostParamsToolType = "community_search"
	ExtractionEstimateCostParamsToolTypeFavoriters                 ExtractionEstimateCostParamsToolType = "favoriters"
	ExtractionEstimateCostParamsToolTypeFollowerExplorer           ExtractionEstimateCostParamsToolType = "follower_explorer"
	ExtractionEstimateCostParamsToolTypeFollowingExplorer          ExtractionEstimateCostParamsToolType = "following_explorer"
	ExtractionEstimateCostParamsToolTypeListFollowerExplorer       ExtractionEstimateCostParamsToolType = "list_follower_explorer"
	ExtractionEstimateCostParamsToolTypeListMemberExtractor        ExtractionEstimateCostParamsToolType = "list_member_extractor"
	ExtractionEstimateCostParamsToolTypeListPostExtractor          ExtractionEstimateCostParamsToolType = "list_post_extractor"
	ExtractionEstimateCostParamsToolTypeMentionExtractor           ExtractionEstimateCostParamsToolType = "mention_extractor"
	ExtractionEstimateCostParamsToolTypePeopleSearch               ExtractionEstimateCostParamsToolType = "people_search"
	ExtractionEstimateCostParamsToolTypePostExtractor              ExtractionEstimateCostParamsToolType = "post_extractor"
	ExtractionEstimateCostParamsToolTypeQuoteExtractor             ExtractionEstimateCostParamsToolType = "quote_extractor"
	ExtractionEstimateCostParamsToolTypeReplyExtractor             ExtractionEstimateCostParamsToolType = "reply_extractor"
	ExtractionEstimateCostParamsToolTypeRepostExtractor            ExtractionEstimateCostParamsToolType = "repost_extractor"
	ExtractionEstimateCostParamsToolTypeSpaceExplorer              ExtractionEstimateCostParamsToolType = "space_explorer"
	ExtractionEstimateCostParamsToolTypeThreadExtractor            ExtractionEstimateCostParamsToolType = "thread_extractor"
	ExtractionEstimateCostParamsToolTypeTweetSearchExtractor       ExtractionEstimateCostParamsToolType = "tweet_search_extractor"
	ExtractionEstimateCostParamsToolTypeUserLikes                  ExtractionEstimateCostParamsToolType = "user_likes"
	ExtractionEstimateCostParamsToolTypeUserMedia                  ExtractionEstimateCostParamsToolType = "user_media"
	ExtractionEstimateCostParamsToolTypeVerifiedFollowerExplorer   ExtractionEstimateCostParamsToolType = "verified_follower_explorer"
)

// Media type used for estimate filtering (tweet_search_extractor)
type ExtractionEstimateCostParamsMediaType string

const (
	ExtractionEstimateCostParamsMediaTypeImages ExtractionEstimateCostParamsMediaType = "images"
	ExtractionEstimateCostParamsMediaTypeVideos ExtractionEstimateCostParamsMediaType = "videos"
	ExtractionEstimateCostParamsMediaTypeGifs   ExtractionEstimateCostParamsMediaType = "gifs"
	ExtractionEstimateCostParamsMediaTypeMedia  ExtractionEstimateCostParamsMediaType = "media"
	ExtractionEstimateCostParamsMediaTypeLinks  ExtractionEstimateCostParamsMediaType = "links"
	ExtractionEstimateCostParamsMediaTypeNone   ExtractionEstimateCostParamsMediaType = "none"
)

// Quote mode used for estimation (tweet_search_extractor)
type ExtractionEstimateCostParamsQuotes string

const (
	ExtractionEstimateCostParamsQuotesInclude ExtractionEstimateCostParamsQuotes = "include"
	ExtractionEstimateCostParamsQuotesExclude ExtractionEstimateCostParamsQuotes = "exclude"
	ExtractionEstimateCostParamsQuotesOnly    ExtractionEstimateCostParamsQuotes = "only"
)

// Reply mode used for estimation (tweet_search_extractor)
type ExtractionEstimateCostParamsReplies string

const (
	ExtractionEstimateCostParamsRepliesInclude ExtractionEstimateCostParamsReplies = "include"
	ExtractionEstimateCostParamsRepliesExclude ExtractionEstimateCostParamsReplies = "exclude"
	ExtractionEstimateCostParamsRepliesOnly    ExtractionEstimateCostParamsReplies = "only"
)

// Retweet mode used for estimation (tweet_search_extractor)
type ExtractionEstimateCostParamsRetweets string

const (
	ExtractionEstimateCostParamsRetweetsInclude ExtractionEstimateCostParamsRetweets = "include"
	ExtractionEstimateCostParamsRetweetsExclude ExtractionEstimateCostParamsRetweets = "exclude"
	ExtractionEstimateCostParamsRetweetsOnly    ExtractionEstimateCostParamsRetweets = "only"
)

type ExtractionExportResultsParams struct {
	// Export file format
	//
	// Any of "csv", "json", "md", "md-document", "pdf", "txt", "xlsx".
	Format ExtractionExportResultsParamsFormat `query:"format,omitzero" api:"required" json:"-"`
	paramObj
}

// URLQuery serializes [ExtractionExportResultsParams]'s query parameters as
// `url.Values`.
func (r ExtractionExportResultsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Export file format
type ExtractionExportResultsParamsFormat string

const (
	ExtractionExportResultsParamsFormatCsv        ExtractionExportResultsParamsFormat = "csv"
	ExtractionExportResultsParamsFormatJson       ExtractionExportResultsParamsFormat = "json"
	ExtractionExportResultsParamsFormatMd         ExtractionExportResultsParamsFormat = "md"
	ExtractionExportResultsParamsFormatMdDocument ExtractionExportResultsParamsFormat = "md-document"
	ExtractionExportResultsParamsFormatPdf        ExtractionExportResultsParamsFormat = "pdf"
	ExtractionExportResultsParamsFormatTxt        ExtractionExportResultsParamsFormat = "txt"
	ExtractionExportResultsParamsFormatXlsx       ExtractionExportResultsParamsFormat = "xlsx"
)

type ExtractionRunParams struct {
	// Identifier for the extraction tool used to run a job.
	//
	// Any of "article_extractor", "community_extractor",
	// "community_moderator_explorer", "community_post_extractor", "community_search",
	// "favoriters", "follower_explorer", "following_explorer",
	// "list_follower_explorer", "list_member_extractor", "list_post_extractor",
	// "mention_extractor", "people_search", "post_extractor", "quote_extractor",
	// "reply_extractor", "repost_extractor", "space_explorer", "thread_extractor",
	// "tweet_search_extractor", "user_likes", "user_media",
	// "verified_follower_explorer".
	ToolType ExtractionRunParamsToolType `json:"toolType,omitzero" api:"required"`
	// Raw advanced search query appended as-is (tweet_search_extractor)
	AdvancedQuery param.Opt[string] `json:"advancedQuery,omitzero"`
	// Words or quoted phrases where any one can match. Separate with spaces, commas,
	// or lines. (tweet_search_extractor)
	AnyWords param.Opt[string] `json:"anyWords,omitzero"`
	// Geo bounding box, e.g. -74.1 40.6 -73.9 40.8 (tweet_search_extractor)
	BoundingBox param.Opt[string] `json:"boundingBox,omitzero"`
	// Cashtags separated by spaces, commas, or lines. (tweet_search_extractor)
	Cashtags param.Opt[string] `json:"cashtags,omitzero"`
	// Conversation ID filter (tweet_search_extractor)
	ConversationID param.Opt[string] `json:"conversationId,omitzero"`
	// Exact phrase to match (tweet_search_extractor)
	ExactPhrase param.Opt[string] `json:"exactPhrase,omitzero"`
	// Words or quoted phrases to exclude. Separate with spaces, commas, or lines.
	// (tweet_search_extractor)
	ExcludeWords param.Opt[string] `json:"excludeWords,omitzero"`
	// Filter by author username (tweet_search_extractor)
	FromUser param.Opt[string] `json:"fromUser,omitzero"`
	// Hashtags separated by spaces, commas, or lines. (tweet_search_extractor)
	Hashtags param.Opt[string] `json:"hashtags,omitzero"`
	// Only replies to this tweet ID (tweet_search_extractor)
	InReplyToTweetID param.Opt[string] `json:"inReplyToTweetId,omitzero"`
	// Language code filter (tweet_search_extractor)
	Language param.Opt[string] `json:"language,omitzero"`
	// Search within a list ID (tweet_search_extractor)
	ListID param.Opt[string] `json:"listId,omitzero"`
	// Filter tweets mentioning a username (tweet_search_extractor)
	Mentioning param.Opt[string] `json:"mentioning,omitzero"`
	// Minimum likes threshold (tweet_search_extractor)
	MinFaves param.Opt[int64] `json:"minFaves,omitzero"`
	// Minimum quote count threshold (tweet_search_extractor)
	MinQuotes param.Opt[int64] `json:"minQuotes,omitzero"`
	// Minimum replies threshold (tweet_search_extractor)
	MinReplies param.Opt[int64] `json:"minReplies,omitzero"`
	// Minimum retweets threshold (tweet_search_extractor)
	MinRetweets param.Opt[int64] `json:"minRetweets,omitzero"`
	// Search within a place ID (tweet_search_extractor)
	Place param.Opt[string] `json:"place,omitzero"`
	// Search within a country code (tweet_search_extractor)
	PlaceCountry param.Opt[string] `json:"placeCountry,omitzero"`
	// Geo point radius, e.g. -73.99 40.73 25mi (tweet_search_extractor)
	PointRadius param.Opt[string] `json:"pointRadius,omitzero"`
	// Only quotes of this tweet ID (tweet_search_extractor)
	QuotesOfTweetID param.Opt[string] `json:"quotesOfTweetId,omitzero"`
	// Maximum number of results to extract. When set, the extraction stops after
	// reaching this limit.
	ResultsLimit param.Opt[int64] `json:"resultsLimit,omitzero"`
	// Only retweets of this tweet ID (tweet_search_extractor)
	RetweetsOfTweetID param.Opt[string] `json:"retweetsOfTweetId,omitzero"`
	// Required for tweet_search_extractor & community_search.
	SearchQuery param.Opt[string] `json:"searchQuery,omitzero"`
	// Start date YYYY-MM-DD (tweet_search_extractor)
	SinceDate param.Opt[time.Time] `json:"sinceDate,omitzero" format:"date"`
	// Required for community_post_extractor & community_search.
	TargetCommunityID param.Opt[string] `json:"targetCommunityId,omitzero"`
	// Required for list_follower_explorer, list_member_extractor &
	// list_post_extractor.
	TargetListID param.Opt[string] `json:"targetListId,omitzero"`
	// Required for space_explorer.
	TargetSpaceID  param.Opt[string] `json:"targetSpaceId,omitzero"`
	TargetTweetID  param.Opt[string] `json:"targetTweetId,omitzero"`
	TargetUsername param.Opt[string] `json:"targetUsername,omitzero"`
	// Filter replies sent to a username (tweet_search_extractor)
	ToUser param.Opt[string] `json:"toUser,omitzero"`
	// End date YYYY-MM-DD (tweet_search_extractor)
	UntilDate param.Opt[time.Time] `json:"untilDate,omitzero" format:"date"`
	// URL substring or domain filter (tweet_search_extractor)
	URL param.Opt[string] `json:"url,omitzero"`
	// Only verified authors (tweet_search_extractor)
	VerifiedOnly param.Opt[bool] `json:"verifiedOnly,omitzero"`
	// Media type filter (tweet_search_extractor)
	//
	// Any of "images", "videos", "gifs", "media", "links", "none".
	MediaType ExtractionRunParamsMediaType `json:"mediaType,omitzero"`
	// Quote mode (tweet_search_extractor)
	//
	// Any of "include", "exclude", "only".
	Quotes ExtractionRunParamsQuotes `json:"quotes,omitzero"`
	// Reply mode (tweet_search_extractor)
	//
	// Any of "include", "exclude", "only".
	Replies ExtractionRunParamsReplies `json:"replies,omitzero"`
	// Retweet mode (tweet_search_extractor)
	//
	// Any of "include", "exclude", "only".
	Retweets ExtractionRunParamsRetweets `json:"retweets,omitzero"`
	paramObj
}

func (r ExtractionRunParams) MarshalJSON() (data []byte, err error) {
	type shadow ExtractionRunParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractionRunParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Identifier for the extraction tool used to run a job.
type ExtractionRunParamsToolType string

const (
	ExtractionRunParamsToolTypeArticleExtractor           ExtractionRunParamsToolType = "article_extractor"
	ExtractionRunParamsToolTypeCommunityExtractor         ExtractionRunParamsToolType = "community_extractor"
	ExtractionRunParamsToolTypeCommunityModeratorExplorer ExtractionRunParamsToolType = "community_moderator_explorer"
	ExtractionRunParamsToolTypeCommunityPostExtractor     ExtractionRunParamsToolType = "community_post_extractor"
	ExtractionRunParamsToolTypeCommunitySearch            ExtractionRunParamsToolType = "community_search"
	ExtractionRunParamsToolTypeFavoriters                 ExtractionRunParamsToolType = "favoriters"
	ExtractionRunParamsToolTypeFollowerExplorer           ExtractionRunParamsToolType = "follower_explorer"
	ExtractionRunParamsToolTypeFollowingExplorer          ExtractionRunParamsToolType = "following_explorer"
	ExtractionRunParamsToolTypeListFollowerExplorer       ExtractionRunParamsToolType = "list_follower_explorer"
	ExtractionRunParamsToolTypeListMemberExtractor        ExtractionRunParamsToolType = "list_member_extractor"
	ExtractionRunParamsToolTypeListPostExtractor          ExtractionRunParamsToolType = "list_post_extractor"
	ExtractionRunParamsToolTypeMentionExtractor           ExtractionRunParamsToolType = "mention_extractor"
	ExtractionRunParamsToolTypePeopleSearch               ExtractionRunParamsToolType = "people_search"
	ExtractionRunParamsToolTypePostExtractor              ExtractionRunParamsToolType = "post_extractor"
	ExtractionRunParamsToolTypeQuoteExtractor             ExtractionRunParamsToolType = "quote_extractor"
	ExtractionRunParamsToolTypeReplyExtractor             ExtractionRunParamsToolType = "reply_extractor"
	ExtractionRunParamsToolTypeRepostExtractor            ExtractionRunParamsToolType = "repost_extractor"
	ExtractionRunParamsToolTypeSpaceExplorer              ExtractionRunParamsToolType = "space_explorer"
	ExtractionRunParamsToolTypeThreadExtractor            ExtractionRunParamsToolType = "thread_extractor"
	ExtractionRunParamsToolTypeTweetSearchExtractor       ExtractionRunParamsToolType = "tweet_search_extractor"
	ExtractionRunParamsToolTypeUserLikes                  ExtractionRunParamsToolType = "user_likes"
	ExtractionRunParamsToolTypeUserMedia                  ExtractionRunParamsToolType = "user_media"
	ExtractionRunParamsToolTypeVerifiedFollowerExplorer   ExtractionRunParamsToolType = "verified_follower_explorer"
)

// Media type filter (tweet_search_extractor)
type ExtractionRunParamsMediaType string

const (
	ExtractionRunParamsMediaTypeImages ExtractionRunParamsMediaType = "images"
	ExtractionRunParamsMediaTypeVideos ExtractionRunParamsMediaType = "videos"
	ExtractionRunParamsMediaTypeGifs   ExtractionRunParamsMediaType = "gifs"
	ExtractionRunParamsMediaTypeMedia  ExtractionRunParamsMediaType = "media"
	ExtractionRunParamsMediaTypeLinks  ExtractionRunParamsMediaType = "links"
	ExtractionRunParamsMediaTypeNone   ExtractionRunParamsMediaType = "none"
)

// Quote mode (tweet_search_extractor)
type ExtractionRunParamsQuotes string

const (
	ExtractionRunParamsQuotesInclude ExtractionRunParamsQuotes = "include"
	ExtractionRunParamsQuotesExclude ExtractionRunParamsQuotes = "exclude"
	ExtractionRunParamsQuotesOnly    ExtractionRunParamsQuotes = "only"
)

// Reply mode (tweet_search_extractor)
type ExtractionRunParamsReplies string

const (
	ExtractionRunParamsRepliesInclude ExtractionRunParamsReplies = "include"
	ExtractionRunParamsRepliesExclude ExtractionRunParamsReplies = "exclude"
	ExtractionRunParamsRepliesOnly    ExtractionRunParamsReplies = "only"
)

// Retweet mode (tweet_search_extractor)
type ExtractionRunParamsRetweets string

const (
	ExtractionRunParamsRetweetsInclude ExtractionRunParamsRetweets = "include"
	ExtractionRunParamsRetweetsExclude ExtractionRunParamsRetweets = "exclude"
	ExtractionRunParamsRetweetsOnly    ExtractionRunParamsRetweets = "only"
)
