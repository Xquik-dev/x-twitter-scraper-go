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
func (r *ExtractionService) Run(ctx context.Context, params ExtractionRunParams, opts ...option.RequestOption) (res *ExtractionRunResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "extractions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
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
	// Any of "followers", "following", "collection", "paginationCap", "posts",
	// "quoteCount", "replyCount", "resultsLimit", "retweetCount", "unknown".
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
	ExtractionEstimateCostResponseSourceCollection    ExtractionEstimateCostResponseSource = "collection"
	ExtractionEstimateCostResponseSourcePaginationCap ExtractionEstimateCostResponseSource = "paginationCap"
	ExtractionEstimateCostResponseSourcePosts         ExtractionEstimateCostResponseSource = "posts"
	ExtractionEstimateCostResponseSourceQuoteCount    ExtractionEstimateCostResponseSource = "quoteCount"
	ExtractionEstimateCostResponseSourceReplyCount    ExtractionEstimateCostResponseSource = "replyCount"
	ExtractionEstimateCostResponseSourceResultsLimit  ExtractionEstimateCostResponseSource = "resultsLimit"
	ExtractionEstimateCostResponseSourceRetweetCount  ExtractionEstimateCostResponseSource = "retweetCount"
	ExtractionEstimateCostResponseSourceUnknown       ExtractionEstimateCostResponseSource = "unknown"
)

type ExtractionRunResponse struct {
	Allowed          bool   `json:"allowed" api:"required"`
	CreditsAvailable string `json:"creditsAvailable" api:"required"`
	CreditsRequired  string `json:"creditsRequired" api:"required"`
	EstimatedResults int64  `json:"estimatedResults" api:"required"`
	Source           string `json:"source" api:"required"`
	ResolvedXUserID  string `json:"resolvedXUserId"`
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
func (r ExtractionRunResponse) RawJSON() string { return r.JSON.raw }
func (r *ExtractionRunResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExtractionGetParams struct {
	// Previous nextCursor.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Use outputMode=raw instead.
	IncludeRaw param.Opt[bool] `query:"includeRaw,omitzero" json:"-"`
	// Maximum number of results to return (1-1000, default 100)
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Preserve source keys or convert result field names.
	//
	// Any of "source", "camelCase", "snake_case".
	FieldStyle ExtractionGetParamsFieldStyle `query:"fieldStyle,omitzero" json:"-"`
	// Select compact, full, or raw-compatible result fields.
	//
	// Any of "compact", "full", "raw".
	OutputMode ExtractionGetParamsOutputMode `query:"outputMode,omitzero" json:"-"`
	// Keep enrichment nested or merge it into each result.
	//
	// Any of "nested", "flat".
	OutputPreset ExtractionGetParamsOutputPreset `query:"outputPreset,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ExtractionGetParams]'s query parameters as `url.Values`.
func (r ExtractionGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Preserve source keys or convert result field names.
type ExtractionGetParamsFieldStyle string

const (
	ExtractionGetParamsFieldStyleSource    ExtractionGetParamsFieldStyle = "source"
	ExtractionGetParamsFieldStyleCamelCase ExtractionGetParamsFieldStyle = "camelCase"
	ExtractionGetParamsFieldStyleSnakeCase ExtractionGetParamsFieldStyle = "snake_case"
)

// Select compact, full, or raw-compatible result fields.
type ExtractionGetParamsOutputMode string

const (
	ExtractionGetParamsOutputModeCompact ExtractionGetParamsOutputMode = "compact"
	ExtractionGetParamsOutputModeFull    ExtractionGetParamsOutputMode = "full"
	ExtractionGetParamsOutputModeRaw     ExtractionGetParamsOutputMode = "raw"
)

// Keep enrichment nested or merge it into each result.
type ExtractionGetParamsOutputPreset string

const (
	ExtractionGetParamsOutputPresetNested ExtractionGetParamsOutputPreset = "nested"
	ExtractionGetParamsOutputPresetFlat   ExtractionGetParamsOutputPreset = "flat"
)

type ExtractionListParams struct {
	// Previous nextCursor.
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
	// Raw advanced search query appended as-is (tweet_search_extractor)
	AdvancedQuery param.Opt[string] `json:"advancedQuery,omitzero"`
	// Words or quoted phrases where any one can match. Separate with spaces, commas,
	// or lines. (tweet_search_extractor)
	AnyWords param.Opt[string] `json:"anyWords,omitzero"`
	// Bio terms separated by commas or lines.
	BioContains param.Opt[string] `json:"bioContains,omitzero"`
	// Return only Blue-verified Tweet authors.
	BlueVerifiedOnly param.Opt[bool] `json:"blueVerifiedOnly,omitzero"`
	// Geo bounding box, e.g. -74.1 40.6 -73.9 40.8 (tweet_search_extractor)
	BoundingBox param.Opt[string] `json:"boundingBox,omitzero"`
	// Match the Tweet card name.
	CardName param.Opt[string] `json:"cardName,omitzero"`
	// Cashtags separated by spaces, commas, or lines. (tweet_search_extractor)
	Cashtags param.Opt[string] `json:"cashtags,omitzero"`
	// Conversation ID filter (tweet_search_extractor)
	ConversationID param.Opt[string] `json:"conversationId,omitzero"`
	// Merge duplicate results across collection targets.
	DedupeAcrossTargets param.Opt[bool] `json:"dedupeAcrossTargets,omitzero"`
	// Exact phrase to match (tweet_search_extractor)
	ExactPhrase param.Opt[string] `json:"exactPhrase,omitzero"`
	// Exclude replies from the source author.
	ExcludeOriginalAuthor param.Opt[bool] `json:"excludeOriginalAuthor,omitzero"`
	// Exclude a source application.
	ExcludeSource param.Opt[string] `json:"excludeSource,omitzero"`
	// Words or quoted phrases to exclude. Separate with spaces, commas, or lines.
	// (tweet_search_extractor)
	ExcludeWords param.Opt[string] `json:"excludeWords,omitzero"`
	// Filter by author username (tweet_search_extractor)
	FromUser param.Opt[string] `json:"fromUser,omitzero"`
	// Match latitude, longitude, and radius.
	Geocode param.Opt[string] `json:"geocode,omitzero"`
	// Hashtags separated by spaces, commas, or lines. (tweet_search_extractor)
	Hashtags param.Opt[string] `json:"hashtags,omitzero"`
	// Require a profile location.
	HasLocation param.Opt[bool] `json:"hasLocation,omitzero"`
	// Return only replies with media.
	HasMediaOnly param.Opt[bool] `json:"hasMediaOnly,omitzero"`
	// Require a profile website.
	HasWebsite param.Opt[bool] `json:"hasWebsite,omitzero"`
	// Include the source post in reply results.
	IncludeOriginalPost param.Opt[bool] `json:"includeOriginalPost,omitzero"`
	// Add matching search terms to collection metadata.
	IncludeSearchTerms param.Opt[bool] `json:"includeSearchTerms,omitzero"`
	// Add source target metadata to each result.
	IncludeTargetMetadata param.Opt[bool] `json:"includeTargetMetadata,omitzero"`
	// Only replies to this tweet ID (tweet_search_extractor)
	InReplyToTweetID param.Opt[string] `json:"inReplyToTweetId,omitzero"`
	// Language code filter (tweet_search_extractor)
	Language param.Opt[string] `json:"language,omitzero"`
	// Search within a list ID (tweet_search_extractor)
	ListID param.Opt[string] `json:"listId,omitzero"`
	// Required profile location text.
	LocationContains param.Opt[string] `json:"locationContains,omitzero"`
	// Maximum nested reply depth.
	MaxDepth param.Opt[int64] `json:"maxDepth,omitzero"`
	// Maximum follower count for profile results.
	MaxFollowers param.Opt[int64] `json:"maxFollowers,omitzero"`
	// Maximum following count for profile results.
	MaxFollowing param.Opt[int64] `json:"maxFollowing,omitzero"`
	// Return Tweets older than this Tweet ID.
	MaxID param.Opt[string] `json:"maxId,omitzero"`
	// Maximum results collected for each target.
	MaxItemsPerTarget param.Opt[int64] `json:"maxItemsPerTarget,omitzero"`
	// Maximum Tweet like count.
	MaxLikes param.Opt[int64] `json:"maxLikes,omitzero"`
	// Reply pages collected for each target.
	MaxPagesPerTarget param.Opt[int64] `json:"maxPagesPerTarget,omitzero"`
	// Maximum post count for profile results.
	MaxPosts param.Opt[int64] `json:"maxPosts,omitzero"`
	// Maximum Tweet quote count.
	MaxQuotes param.Opt[int64] `json:"maxQuotes,omitzero"`
	// Maximum Tweet reply count.
	MaxReplies param.Opt[int64] `json:"maxReplies,omitzero"`
	// Maximum Tweet repost count.
	MaxRetweets param.Opt[int64] `json:"maxRetweets,omitzero"`
	// Filter tweets mentioning a username (tweet_search_extractor)
	Mentioning param.Opt[string] `json:"mentioning,omitzero"`
	// Minimum profile age in days.
	MinAccountAgeDays param.Opt[int64] `json:"minAccountAgeDays,omitzero"`
	// Minimum Tweet bookmark count.
	MinBookmarks param.Opt[int64] `json:"minBookmarks,omitzero"`
	// Minimum likes threshold (tweet_search_extractor)
	MinFaves param.Opt[int64] `json:"minFaves,omitzero"`
	// Minimum follower count for profile results.
	MinFollowers param.Opt[int64] `json:"minFollowers,omitzero"`
	// Minimum following count for profile results.
	MinFollowing param.Opt[int64] `json:"minFollowing,omitzero"`
	// Minimum post count for profile results.
	MinPosts param.Opt[int64] `json:"minPosts,omitzero"`
	// Minimum quote count threshold (tweet_search_extractor)
	MinQuotes param.Opt[int64] `json:"minQuotes,omitzero"`
	// Minimum replies threshold (tweet_search_extractor)
	MinReplies param.Opt[int64] `json:"minReplies,omitzero"`
	// Minimum retweets threshold (tweet_search_extractor)
	MinRetweets param.Opt[int64] `json:"minRetweets,omitzero"`
	// Minimum Tweet view count.
	MinViews param.Opt[int64] `json:"minViews,omitzero"`
	// Only return native reposts.
	NativeRetweets param.Opt[bool] `json:"nativeRetweets,omitzero"`
	// Match a place name.
	Near param.Opt[string] `json:"near,omitzero"`
	// Only return news results.
	News param.Opt[bool] `json:"news,omitzero"`
	// Shortcut for dedupeMode=merge.
	OverlapMode param.Opt[bool] `json:"overlapMode,omitzero"`
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
	// Enable the safe-search filter.
	Safe param.Opt[bool] `json:"safe,omitzero"`
	// Required for tweet_search_extractor & community_search.
	SearchQuery param.Opt[string] `json:"searchQuery,omitzero"`
	// Start date YYYY-MM-DD (tweet_search_extractor)
	SinceDate param.Opt[time.Time] `json:"sinceDate,omitzero" format:"date"`
	// Return Tweets newer than this Tweet ID.
	SinceID param.Opt[string] `json:"sinceId,omitzero"`
	// Match the source application.
	Source param.Opt[string] `json:"source,omitzero"`
	// Resume one reply target from this cursor.
	StartCursor param.Opt[string] `json:"startCursor,omitzero"`
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
	// Required username text.
	UsernameContains param.Opt[string] `json:"usernameContains,omitzero"`
	// Only verified authors (tweet_search_extractor)
	VerifiedOnly param.Opt[bool] `json:"verifiedOnly,omitzero"`
	// Exact profile verification type.
	VerifiedType param.Opt[string] `json:"verifiedType,omitzero"`
	// Set the radius for the near filter.
	Within param.Opt[string] `json:"within,omitzero"`
	// Match Tweets inside a recent time window.
	WithinTime param.Opt[string] `json:"withinTime,omitzero"`
	// Reply collection strategy.
	//
	// Any of "auto", "complete", "direct", "search", "thread".
	CollectionStrategy ExtractionEstimateCostParamsCollectionStrategy `json:"collectionStrategy,omitzero"`
	// Keep target duplicates, first rows, or merged overlap.
	//
	// Any of "none", "first", "merge".
	DedupeMode ExtractionEstimateCostParamsDedupeMode `json:"dedupeMode,omitzero"`
	// Media type filter (tweet_search_extractor)
	//
	// Any of "images", "videos", "gifs", "media", "links", "none".
	MediaType ExtractionEstimateCostParamsMediaType `json:"mediaType,omitzero"`
	// Search ranking applied to every query.
	//
	// Any of "Latest", "Top", "Both".
	QueryType ExtractionEstimateCostParamsQueryType `json:"queryType,omitzero"`
	// Quote mode (tweet_search_extractor)
	//
	// Any of "include", "exclude", "only".
	Quotes ExtractionEstimateCostParamsQuotes `json:"quotes,omitzero"`
	// Profile relations processed within one job.
	RelationTargets []ExtractionEstimateCostParamsRelationTarget `json:"relationTargets,omitzero"`
	// Reply mode (tweet_search_extractor)
	//
	// Any of "include", "exclude", "only".
	Replies ExtractionEstimateCostParamsReplies `json:"replies,omitzero"`
	// Retweet mode (tweet_search_extractor)
	//
	// Any of "include", "exclude", "only".
	Retweets ExtractionEstimateCostParamsRetweets `json:"retweets,omitzero"`
	// Reply depth scope.
	//
	// Any of "all", "direct", "nested".
	Scope ExtractionEstimateCostParamsScope `json:"scope,omitzero"`
	// Search queries processed as one collection job.
	SearchQueries []string `json:"searchQueries,omitzero"`
	// Reply start time as ISO 8601 or Unix seconds.
	SinceTime ExtractionEstimateCostParamsSinceTimeUnion `json:"sinceTime,omitzero" format:"date-time"`
	// Reply result order.
	//
	// Any of "relevance", "latest", "oldest", "likes".
	Sort ExtractionEstimateCostParamsSort `json:"sort,omitzero"`
	// Community IDs processed as one collection job.
	TargetCommunityIDs []string `json:"targetCommunityIds,omitzero"`
	// List IDs processed as one collection job.
	TargetListIDs []string `json:"targetListIds,omitzero"`
	// Mixed targets auto-routed within one job.
	Targets []ExtractionEstimateCostParamsTargetUnion `json:"targets,omitzero"`
	// Tweet IDs processed as one collection job.
	TargetTweetIDs []string `json:"targetTweetIds,omitzero"`
	// Usernames processed as one collection job.
	TargetUsernames []string `json:"targetUsernames,omitzero"`
	// Reply end time as ISO 8601 or Unix seconds.
	UntilTime ExtractionEstimateCostParamsUntilTimeUnion `json:"untilTime,omitzero" format:"date-time"`
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

// Reply collection strategy.
type ExtractionEstimateCostParamsCollectionStrategy string

const (
	ExtractionEstimateCostParamsCollectionStrategyAuto     ExtractionEstimateCostParamsCollectionStrategy = "auto"
	ExtractionEstimateCostParamsCollectionStrategyComplete ExtractionEstimateCostParamsCollectionStrategy = "complete"
	ExtractionEstimateCostParamsCollectionStrategyDirect   ExtractionEstimateCostParamsCollectionStrategy = "direct"
	ExtractionEstimateCostParamsCollectionStrategySearch   ExtractionEstimateCostParamsCollectionStrategy = "search"
	ExtractionEstimateCostParamsCollectionStrategyThread   ExtractionEstimateCostParamsCollectionStrategy = "thread"
)

// Keep target duplicates, first rows, or merged overlap.
type ExtractionEstimateCostParamsDedupeMode string

const (
	ExtractionEstimateCostParamsDedupeModeNone  ExtractionEstimateCostParamsDedupeMode = "none"
	ExtractionEstimateCostParamsDedupeModeFirst ExtractionEstimateCostParamsDedupeMode = "first"
	ExtractionEstimateCostParamsDedupeModeMerge ExtractionEstimateCostParamsDedupeMode = "merge"
)

// Media type filter (tweet_search_extractor)
type ExtractionEstimateCostParamsMediaType string

const (
	ExtractionEstimateCostParamsMediaTypeImages ExtractionEstimateCostParamsMediaType = "images"
	ExtractionEstimateCostParamsMediaTypeVideos ExtractionEstimateCostParamsMediaType = "videos"
	ExtractionEstimateCostParamsMediaTypeGifs   ExtractionEstimateCostParamsMediaType = "gifs"
	ExtractionEstimateCostParamsMediaTypeMedia  ExtractionEstimateCostParamsMediaType = "media"
	ExtractionEstimateCostParamsMediaTypeLinks  ExtractionEstimateCostParamsMediaType = "links"
	ExtractionEstimateCostParamsMediaTypeNone   ExtractionEstimateCostParamsMediaType = "none"
)

// Search ranking applied to every query.
type ExtractionEstimateCostParamsQueryType string

const (
	ExtractionEstimateCostParamsQueryTypeLatest ExtractionEstimateCostParamsQueryType = "Latest"
	ExtractionEstimateCostParamsQueryTypeTop    ExtractionEstimateCostParamsQueryType = "Top"
	ExtractionEstimateCostParamsQueryTypeBoth   ExtractionEstimateCostParamsQueryType = "Both"
)

// Quote mode (tweet_search_extractor)
type ExtractionEstimateCostParamsQuotes string

const (
	ExtractionEstimateCostParamsQuotesInclude ExtractionEstimateCostParamsQuotes = "include"
	ExtractionEstimateCostParamsQuotesExclude ExtractionEstimateCostParamsQuotes = "exclude"
	ExtractionEstimateCostParamsQuotesOnly    ExtractionEstimateCostParamsQuotes = "only"
)

// One target and relation in a mixed profile collection.
//
// The properties Relation, Value are required.
type ExtractionEstimateCostParamsRelationTarget struct {
	// Any of "community_members", "followers", "following", "list_followers",
	// "list_members", "verified_followers".
	Relation string `json:"relation,omitzero" api:"required"`
	Value    string `json:"value" api:"required"`
	paramObj
}

func (r ExtractionEstimateCostParamsRelationTarget) MarshalJSON() (data []byte, err error) {
	type shadow ExtractionEstimateCostParamsRelationTarget
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractionEstimateCostParamsRelationTarget) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ExtractionEstimateCostParamsRelationTarget](
		"relation", "community_members", "followers", "following", "list_followers", "list_members", "verified_followers",
	)
}

// Reply mode (tweet_search_extractor)
type ExtractionEstimateCostParamsReplies string

const (
	ExtractionEstimateCostParamsRepliesInclude ExtractionEstimateCostParamsReplies = "include"
	ExtractionEstimateCostParamsRepliesExclude ExtractionEstimateCostParamsReplies = "exclude"
	ExtractionEstimateCostParamsRepliesOnly    ExtractionEstimateCostParamsReplies = "only"
)

// Retweet mode (tweet_search_extractor)
type ExtractionEstimateCostParamsRetweets string

const (
	ExtractionEstimateCostParamsRetweetsInclude ExtractionEstimateCostParamsRetweets = "include"
	ExtractionEstimateCostParamsRetweetsExclude ExtractionEstimateCostParamsRetweets = "exclude"
	ExtractionEstimateCostParamsRetweetsOnly    ExtractionEstimateCostParamsRetweets = "only"
)

// Reply depth scope.
type ExtractionEstimateCostParamsScope string

const (
	ExtractionEstimateCostParamsScopeAll    ExtractionEstimateCostParamsScope = "all"
	ExtractionEstimateCostParamsScopeDirect ExtractionEstimateCostParamsScope = "direct"
	ExtractionEstimateCostParamsScopeNested ExtractionEstimateCostParamsScope = "nested"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractionEstimateCostParamsSinceTimeUnion struct {
	OfTime param.Opt[time.Time] `json:",omitzero,inline"`
	OfInt  param.Opt[int64]     `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractionEstimateCostParamsSinceTimeUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfTime, u.OfInt)
}
func (u *ExtractionEstimateCostParamsSinceTimeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Reply result order.
type ExtractionEstimateCostParamsSort string

const (
	ExtractionEstimateCostParamsSortRelevance ExtractionEstimateCostParamsSort = "relevance"
	ExtractionEstimateCostParamsSortLatest    ExtractionEstimateCostParamsSort = "latest"
	ExtractionEstimateCostParamsSortOldest    ExtractionEstimateCostParamsSort = "oldest"
	ExtractionEstimateCostParamsSortLikes     ExtractionEstimateCostParamsSort = "likes"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractionEstimateCostParamsTargetUnion struct {
	OfString                              param.Opt[string]                         `json:",omitzero,inline"`
	OfExtractionEstimateCostsTargetObject *ExtractionEstimateCostParamsTargetObject `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractionEstimateCostParamsTargetUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfExtractionEstimateCostsTargetObject)
}
func (u *ExtractionEstimateCostParamsTargetUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The properties Kind, Value are required.
type ExtractionEstimateCostParamsTargetObject struct {
	// Any of "favoriters", "list", "profile", "profile_likes", "profile_media",
	// "profile_replies", "quotes", "replies", "retweeters", "search", "thread",
	// "tweet".
	Kind  string `json:"kind,omitzero" api:"required"`
	Value string `json:"value" api:"required"`
	paramObj
}

func (r ExtractionEstimateCostParamsTargetObject) MarshalJSON() (data []byte, err error) {
	type shadow ExtractionEstimateCostParamsTargetObject
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractionEstimateCostParamsTargetObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ExtractionEstimateCostParamsTargetObject](
		"kind", "favoriters", "list", "profile", "profile_likes", "profile_media", "profile_replies", "quotes", "replies", "retweeters", "search", "thread", "tweet",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractionEstimateCostParamsUntilTimeUnion struct {
	OfTime param.Opt[time.Time] `json:",omitzero,inline"`
	OfInt  param.Opt[int64]     `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractionEstimateCostParamsUntilTimeUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfTime, u.OfInt)
}
func (u *ExtractionEstimateCostParamsUntilTimeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

type ExtractionExportResultsParams struct {
	// Export file format
	//
	// Any of "csv", "json", "md", "md-document", "pdf", "txt", "xlsx".
	Format ExtractionExportResultsParamsFormat `query:"format,omitzero" api:"required" json:"-"`
	// Require a non-empty description.
	HasDescription param.Opt[bool] `query:"hasDescription,omitzero" json:"-"`
	// Require a non-empty location.
	HasLocation param.Opt[bool] `query:"hasLocation,omitzero" json:"-"`
	// Require media.
	HasMedia param.Opt[bool] `query:"hasMedia,omitzero" json:"-"`
	// Filter by language code.
	Lang param.Opt[string] `query:"lang,omitzero" json:"-"`
	// Maximum follower count.
	MaxFollowers param.Opt[int64] `query:"maxFollowers,omitzero" json:"-"`
	// Maximum following count.
	MaxFollowing param.Opt[int64] `query:"maxFollowing,omitzero" json:"-"`
	// Maximum post count.
	MaxPosts param.Opt[int64] `query:"maxPosts,omitzero" json:"-"`
	// Minimum follower count.
	MinFollowers param.Opt[int64] `query:"minFollowers,omitzero" json:"-"`
	// Minimum following count.
	MinFollowing param.Opt[int64] `query:"minFollowing,omitzero" json:"-"`
	// Minimum like count.
	MinLikes param.Opt[int64] `query:"minLikes,omitzero" json:"-"`
	// Minimum post count.
	MinPosts param.Opt[int64] `query:"minPosts,omitzero" json:"-"`
	// Minimum reply count.
	MinReplies param.Opt[int64] `query:"minReplies,omitzero" json:"-"`
	// Minimum repost count.
	MinRetweets param.Opt[int64] `query:"minRetweets,omitzero" json:"-"`
	// Minimum view count.
	MinViews param.Opt[int64] `query:"minViews,omitzero" json:"-"`
	// Search exported result text.
	Search param.Opt[string] `query:"search,omitzero" json:"-"`
	// Include results on or after this date.
	SinceDate param.Opt[time.Time] `query:"sinceDate,omitzero" format:"date" json:"-"`
	// Include results on or before this date.
	UntilDate param.Opt[time.Time] `query:"untilDate,omitzero" format:"date" json:"-"`
	// Filter by verified status.
	Verified param.Opt[bool] `query:"verified,omitzero" json:"-"`
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
	// Estimate cost without creating an extraction.
	DryRun param.Opt[bool] `query:"dry_run,omitzero" json:"-"`
	// Raw advanced search query appended as-is (tweet_search_extractor)
	AdvancedQuery param.Opt[string] `json:"advancedQuery,omitzero"`
	// Words or quoted phrases where any one can match. Separate with spaces, commas,
	// or lines. (tweet_search_extractor)
	AnyWords param.Opt[string] `json:"anyWords,omitzero"`
	// Bio terms separated by commas or lines.
	BioContains param.Opt[string] `json:"bioContains,omitzero"`
	// Return only Blue-verified Tweet authors.
	BlueVerifiedOnly param.Opt[bool] `json:"blueVerifiedOnly,omitzero"`
	// Geo bounding box, e.g. -74.1 40.6 -73.9 40.8 (tweet_search_extractor)
	BoundingBox param.Opt[string] `json:"boundingBox,omitzero"`
	// Match the Tweet card name.
	CardName param.Opt[string] `json:"cardName,omitzero"`
	// Cashtags separated by spaces, commas, or lines. (tweet_search_extractor)
	Cashtags param.Opt[string] `json:"cashtags,omitzero"`
	// Conversation ID filter (tweet_search_extractor)
	ConversationID param.Opt[string] `json:"conversationId,omitzero"`
	// Merge duplicate results across collection targets.
	DedupeAcrossTargets param.Opt[bool] `json:"dedupeAcrossTargets,omitzero"`
	// Exact phrase to match (tweet_search_extractor)
	ExactPhrase param.Opt[string] `json:"exactPhrase,omitzero"`
	// Exclude replies from the source author.
	ExcludeOriginalAuthor param.Opt[bool] `json:"excludeOriginalAuthor,omitzero"`
	// Exclude a source application.
	ExcludeSource param.Opt[string] `json:"excludeSource,omitzero"`
	// Words or quoted phrases to exclude. Separate with spaces, commas, or lines.
	// (tweet_search_extractor)
	ExcludeWords param.Opt[string] `json:"excludeWords,omitzero"`
	// Filter by author username (tweet_search_extractor)
	FromUser param.Opt[string] `json:"fromUser,omitzero"`
	// Match latitude, longitude, and radius.
	Geocode param.Opt[string] `json:"geocode,omitzero"`
	// Hashtags separated by spaces, commas, or lines. (tweet_search_extractor)
	Hashtags param.Opt[string] `json:"hashtags,omitzero"`
	// Require a profile location.
	HasLocation param.Opt[bool] `json:"hasLocation,omitzero"`
	// Return only replies with media.
	HasMediaOnly param.Opt[bool] `json:"hasMediaOnly,omitzero"`
	// Require a profile website.
	HasWebsite param.Opt[bool] `json:"hasWebsite,omitzero"`
	// Include the source post in reply results.
	IncludeOriginalPost param.Opt[bool] `json:"includeOriginalPost,omitzero"`
	// Add matching search terms to collection metadata.
	IncludeSearchTerms param.Opt[bool] `json:"includeSearchTerms,omitzero"`
	// Add source target metadata to each result.
	IncludeTargetMetadata param.Opt[bool] `json:"includeTargetMetadata,omitzero"`
	// Only replies to this tweet ID (tweet_search_extractor)
	InReplyToTweetID param.Opt[string] `json:"inReplyToTweetId,omitzero"`
	// Language code filter (tweet_search_extractor)
	Language param.Opt[string] `json:"language,omitzero"`
	// Search within a list ID (tweet_search_extractor)
	ListID param.Opt[string] `json:"listId,omitzero"`
	// Required profile location text.
	LocationContains param.Opt[string] `json:"locationContains,omitzero"`
	// Maximum nested reply depth.
	MaxDepth param.Opt[int64] `json:"maxDepth,omitzero"`
	// Maximum follower count for profile results.
	MaxFollowers param.Opt[int64] `json:"maxFollowers,omitzero"`
	// Maximum following count for profile results.
	MaxFollowing param.Opt[int64] `json:"maxFollowing,omitzero"`
	// Return Tweets older than this Tweet ID.
	MaxID param.Opt[string] `json:"maxId,omitzero"`
	// Maximum results collected for each target.
	MaxItemsPerTarget param.Opt[int64] `json:"maxItemsPerTarget,omitzero"`
	// Maximum Tweet like count.
	MaxLikes param.Opt[int64] `json:"maxLikes,omitzero"`
	// Reply pages collected for each target.
	MaxPagesPerTarget param.Opt[int64] `json:"maxPagesPerTarget,omitzero"`
	// Maximum post count for profile results.
	MaxPosts param.Opt[int64] `json:"maxPosts,omitzero"`
	// Maximum Tweet quote count.
	MaxQuotes param.Opt[int64] `json:"maxQuotes,omitzero"`
	// Maximum Tweet reply count.
	MaxReplies param.Opt[int64] `json:"maxReplies,omitzero"`
	// Maximum Tweet repost count.
	MaxRetweets param.Opt[int64] `json:"maxRetweets,omitzero"`
	// Filter tweets mentioning a username (tweet_search_extractor)
	Mentioning param.Opt[string] `json:"mentioning,omitzero"`
	// Minimum profile age in days.
	MinAccountAgeDays param.Opt[int64] `json:"minAccountAgeDays,omitzero"`
	// Minimum Tweet bookmark count.
	MinBookmarks param.Opt[int64] `json:"minBookmarks,omitzero"`
	// Minimum likes threshold (tweet_search_extractor)
	MinFaves param.Opt[int64] `json:"minFaves,omitzero"`
	// Minimum follower count for profile results.
	MinFollowers param.Opt[int64] `json:"minFollowers,omitzero"`
	// Minimum following count for profile results.
	MinFollowing param.Opt[int64] `json:"minFollowing,omitzero"`
	// Minimum post count for profile results.
	MinPosts param.Opt[int64] `json:"minPosts,omitzero"`
	// Minimum quote count threshold (tweet_search_extractor)
	MinQuotes param.Opt[int64] `json:"minQuotes,omitzero"`
	// Minimum replies threshold (tweet_search_extractor)
	MinReplies param.Opt[int64] `json:"minReplies,omitzero"`
	// Minimum retweets threshold (tweet_search_extractor)
	MinRetweets param.Opt[int64] `json:"minRetweets,omitzero"`
	// Minimum Tweet view count.
	MinViews param.Opt[int64] `json:"minViews,omitzero"`
	// Only return native reposts.
	NativeRetweets param.Opt[bool] `json:"nativeRetweets,omitzero"`
	// Match a place name.
	Near param.Opt[string] `json:"near,omitzero"`
	// Only return news results.
	News param.Opt[bool] `json:"news,omitzero"`
	// Shortcut for dedupeMode=merge.
	OverlapMode param.Opt[bool] `json:"overlapMode,omitzero"`
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
	// Enable the safe-search filter.
	Safe param.Opt[bool] `json:"safe,omitzero"`
	// Required for tweet_search_extractor & community_search.
	SearchQuery param.Opt[string] `json:"searchQuery,omitzero"`
	// Start date YYYY-MM-DD (tweet_search_extractor)
	SinceDate param.Opt[time.Time] `json:"sinceDate,omitzero" format:"date"`
	// Return Tweets newer than this Tweet ID.
	SinceID param.Opt[string] `json:"sinceId,omitzero"`
	// Match the source application.
	Source param.Opt[string] `json:"source,omitzero"`
	// Resume one reply target from this cursor.
	StartCursor param.Opt[string] `json:"startCursor,omitzero"`
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
	// Required username text.
	UsernameContains param.Opt[string] `json:"usernameContains,omitzero"`
	// Only verified authors (tweet_search_extractor)
	VerifiedOnly param.Opt[bool] `json:"verifiedOnly,omitzero"`
	// Exact profile verification type.
	VerifiedType param.Opt[string] `json:"verifiedType,omitzero"`
	// Set the radius for the near filter.
	Within param.Opt[string] `json:"within,omitzero"`
	// Match Tweets inside a recent time window.
	WithinTime param.Opt[string] `json:"withinTime,omitzero"`
	// Reply collection strategy.
	//
	// Any of "auto", "complete", "direct", "search", "thread".
	CollectionStrategy ExtractionRunParamsCollectionStrategy `json:"collectionStrategy,omitzero"`
	// Keep target duplicates, first rows, or merged overlap.
	//
	// Any of "none", "first", "merge".
	DedupeMode ExtractionRunParamsDedupeMode `json:"dedupeMode,omitzero"`
	// Media type filter (tweet_search_extractor)
	//
	// Any of "images", "videos", "gifs", "media", "links", "none".
	MediaType ExtractionRunParamsMediaType `json:"mediaType,omitzero"`
	// Search ranking applied to every query.
	//
	// Any of "Latest", "Top", "Both".
	QueryType ExtractionRunParamsQueryType `json:"queryType,omitzero"`
	// Quote mode (tweet_search_extractor)
	//
	// Any of "include", "exclude", "only".
	Quotes ExtractionRunParamsQuotes `json:"quotes,omitzero"`
	// Profile relations processed within one job.
	RelationTargets []ExtractionRunParamsRelationTarget `json:"relationTargets,omitzero"`
	// Reply mode (tweet_search_extractor)
	//
	// Any of "include", "exclude", "only".
	Replies ExtractionRunParamsReplies `json:"replies,omitzero"`
	// Retweet mode (tweet_search_extractor)
	//
	// Any of "include", "exclude", "only".
	Retweets ExtractionRunParamsRetweets `json:"retweets,omitzero"`
	// Reply depth scope.
	//
	// Any of "all", "direct", "nested".
	Scope ExtractionRunParamsScope `json:"scope,omitzero"`
	// Search queries processed as one collection job.
	SearchQueries []string `json:"searchQueries,omitzero"`
	// Reply start time as ISO 8601 or Unix seconds.
	SinceTime ExtractionRunParamsSinceTimeUnion `json:"sinceTime,omitzero" format:"date-time"`
	// Reply result order.
	//
	// Any of "relevance", "latest", "oldest", "likes".
	Sort ExtractionRunParamsSort `json:"sort,omitzero"`
	// Community IDs processed as one collection job.
	TargetCommunityIDs []string `json:"targetCommunityIds,omitzero"`
	// List IDs processed as one collection job.
	TargetListIDs []string `json:"targetListIds,omitzero"`
	// Mixed targets auto-routed within one job.
	Targets []ExtractionRunParamsTargetUnion `json:"targets,omitzero"`
	// Tweet IDs processed as one collection job.
	TargetTweetIDs []string `json:"targetTweetIds,omitzero"`
	// Usernames processed as one collection job.
	TargetUsernames []string `json:"targetUsernames,omitzero"`
	// Reply end time as ISO 8601 or Unix seconds.
	UntilTime ExtractionRunParamsUntilTimeUnion `json:"untilTime,omitzero" format:"date-time"`
	paramObj
}

func (r ExtractionRunParams) MarshalJSON() (data []byte, err error) {
	type shadow ExtractionRunParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractionRunParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [ExtractionRunParams]'s query parameters as `url.Values`.
func (r ExtractionRunParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
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

// Reply collection strategy.
type ExtractionRunParamsCollectionStrategy string

const (
	ExtractionRunParamsCollectionStrategyAuto     ExtractionRunParamsCollectionStrategy = "auto"
	ExtractionRunParamsCollectionStrategyComplete ExtractionRunParamsCollectionStrategy = "complete"
	ExtractionRunParamsCollectionStrategyDirect   ExtractionRunParamsCollectionStrategy = "direct"
	ExtractionRunParamsCollectionStrategySearch   ExtractionRunParamsCollectionStrategy = "search"
	ExtractionRunParamsCollectionStrategyThread   ExtractionRunParamsCollectionStrategy = "thread"
)

// Keep target duplicates, first rows, or merged overlap.
type ExtractionRunParamsDedupeMode string

const (
	ExtractionRunParamsDedupeModeNone  ExtractionRunParamsDedupeMode = "none"
	ExtractionRunParamsDedupeModeFirst ExtractionRunParamsDedupeMode = "first"
	ExtractionRunParamsDedupeModeMerge ExtractionRunParamsDedupeMode = "merge"
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

// Search ranking applied to every query.
type ExtractionRunParamsQueryType string

const (
	ExtractionRunParamsQueryTypeLatest ExtractionRunParamsQueryType = "Latest"
	ExtractionRunParamsQueryTypeTop    ExtractionRunParamsQueryType = "Top"
	ExtractionRunParamsQueryTypeBoth   ExtractionRunParamsQueryType = "Both"
)

// Quote mode (tweet_search_extractor)
type ExtractionRunParamsQuotes string

const (
	ExtractionRunParamsQuotesInclude ExtractionRunParamsQuotes = "include"
	ExtractionRunParamsQuotesExclude ExtractionRunParamsQuotes = "exclude"
	ExtractionRunParamsQuotesOnly    ExtractionRunParamsQuotes = "only"
)

// One target and relation in a mixed profile collection.
//
// The properties Relation, Value are required.
type ExtractionRunParamsRelationTarget struct {
	// Any of "community_members", "followers", "following", "list_followers",
	// "list_members", "verified_followers".
	Relation string `json:"relation,omitzero" api:"required"`
	Value    string `json:"value" api:"required"`
	paramObj
}

func (r ExtractionRunParamsRelationTarget) MarshalJSON() (data []byte, err error) {
	type shadow ExtractionRunParamsRelationTarget
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractionRunParamsRelationTarget) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ExtractionRunParamsRelationTarget](
		"relation", "community_members", "followers", "following", "list_followers", "list_members", "verified_followers",
	)
}

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

// Reply depth scope.
type ExtractionRunParamsScope string

const (
	ExtractionRunParamsScopeAll    ExtractionRunParamsScope = "all"
	ExtractionRunParamsScopeDirect ExtractionRunParamsScope = "direct"
	ExtractionRunParamsScopeNested ExtractionRunParamsScope = "nested"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractionRunParamsSinceTimeUnion struct {
	OfTime param.Opt[time.Time] `json:",omitzero,inline"`
	OfInt  param.Opt[int64]     `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractionRunParamsSinceTimeUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfTime, u.OfInt)
}
func (u *ExtractionRunParamsSinceTimeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Reply result order.
type ExtractionRunParamsSort string

const (
	ExtractionRunParamsSortRelevance ExtractionRunParamsSort = "relevance"
	ExtractionRunParamsSortLatest    ExtractionRunParamsSort = "latest"
	ExtractionRunParamsSortOldest    ExtractionRunParamsSort = "oldest"
	ExtractionRunParamsSortLikes     ExtractionRunParamsSort = "likes"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractionRunParamsTargetUnion struct {
	OfString                     param.Opt[string]                `json:",omitzero,inline"`
	OfExtractionRunsTargetObject *ExtractionRunParamsTargetObject `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractionRunParamsTargetUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfExtractionRunsTargetObject)
}
func (u *ExtractionRunParamsTargetUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The properties Kind, Value are required.
type ExtractionRunParamsTargetObject struct {
	// Any of "favoriters", "list", "profile", "profile_likes", "profile_media",
	// "profile_replies", "quotes", "replies", "retweeters", "search", "thread",
	// "tweet".
	Kind  string `json:"kind,omitzero" api:"required"`
	Value string `json:"value" api:"required"`
	paramObj
}

func (r ExtractionRunParamsTargetObject) MarshalJSON() (data []byte, err error) {
	type shadow ExtractionRunParamsTargetObject
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractionRunParamsTargetObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ExtractionRunParamsTargetObject](
		"kind", "favoriters", "list", "profile", "profile_likes", "profile_media", "profile_replies", "quotes", "replies", "retweeters", "search", "thread", "tweet",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractionRunParamsUntilTimeUnion struct {
	OfTime param.Opt[time.Time] `json:",omitzero,inline"`
	OfInt  param.Opt[int64]     `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractionRunParamsUntilTimeUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfTime, u.OfInt)
}
func (u *ExtractionRunParamsUntilTimeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}
