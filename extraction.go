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

// Bulk data extraction (20 tool types)
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

type ExtractionJob struct {
	ID        string    `json:"id" api:"required"`
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Any of "running", "completed", "failed".
	Status ExtractionJobStatus `json:"status" api:"required"`
	// Any of "article_extractor", "community_extractor",
	// "community_moderator_explorer", "community_post_extractor", "community_search",
	// "follower_explorer", "following_explorer", "list_follower_explorer",
	// "list_member_extractor", "list_post_extractor", "mention_extractor",
	// "people_search", "post_extractor", "quote_extractor", "reply_extractor",
	// "repost_extractor", "space_explorer", "thread_extractor",
	// "tweet_search_extractor", "verified_follower_explorer".
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

type ExtractionJobToolType string

const (
	ExtractionJobToolTypeArticleExtractor           ExtractionJobToolType = "article_extractor"
	ExtractionJobToolTypeCommunityExtractor         ExtractionJobToolType = "community_extractor"
	ExtractionJobToolTypeCommunityModeratorExplorer ExtractionJobToolType = "community_moderator_explorer"
	ExtractionJobToolTypeCommunityPostExtractor     ExtractionJobToolType = "community_post_extractor"
	ExtractionJobToolTypeCommunitySearch            ExtractionJobToolType = "community_search"
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
	ExtractionJobToolTypeVerifiedFollowerExplorer   ExtractionJobToolType = "verified_follower_explorer"
)

type ExtractionGetResponse struct {
	HasMore    bool             `json:"hasMore" api:"required"`
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
	Allowed          bool    `json:"allowed" api:"required"`
	EstimatedResults int64   `json:"estimatedResults" api:"required"`
	ProjectedPercent float64 `json:"projectedPercent" api:"required"`
	Source           string  `json:"source" api:"required"`
	UsagePercent     float64 `json:"usagePercent" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Allowed          respjson.Field
		EstimatedResults respjson.Field
		ProjectedPercent respjson.Field
		Source           respjson.Field
		UsagePercent     respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExtractionEstimateCostResponse) RawJSON() string { return r.JSON.raw }
func (r *ExtractionEstimateCostResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExtractionRunResponse struct {
	ID string `json:"id" api:"required"`
	// Any of "running".
	Status ExtractionRunResponseStatus `json:"status" api:"required"`
	// Any of "article_extractor", "community_extractor",
	// "community_moderator_explorer", "community_post_extractor", "community_search",
	// "follower_explorer", "following_explorer", "list_follower_explorer",
	// "list_member_extractor", "list_post_extractor", "mention_extractor",
	// "people_search", "post_extractor", "quote_extractor", "reply_extractor",
	// "repost_extractor", "space_explorer", "thread_extractor",
	// "tweet_search_extractor", "verified_follower_explorer".
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

type ExtractionRunResponseStatus string

const (
	ExtractionRunResponseStatusRunning ExtractionRunResponseStatus = "running"
)

type ExtractionRunResponseToolType string

const (
	ExtractionRunResponseToolTypeArticleExtractor           ExtractionRunResponseToolType = "article_extractor"
	ExtractionRunResponseToolTypeCommunityExtractor         ExtractionRunResponseToolType = "community_extractor"
	ExtractionRunResponseToolTypeCommunityModeratorExplorer ExtractionRunResponseToolType = "community_moderator_explorer"
	ExtractionRunResponseToolTypeCommunityPostExtractor     ExtractionRunResponseToolType = "community_post_extractor"
	ExtractionRunResponseToolTypeCommunitySearch            ExtractionRunResponseToolType = "community_search"
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
	ExtractionRunResponseToolTypeVerifiedFollowerExplorer   ExtractionRunResponseToolType = "verified_follower_explorer"
)

type ExtractionGetParams struct {
	// Cursor for pagination
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	Limit param.Opt[int64]  `query:"limit,omitzero" json:"-"`
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
	// Cursor for pagination
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	Limit param.Opt[int64]  `query:"limit,omitzero" json:"-"`
	// Any of "running", "completed", "failed".
	Status ExtractionListParamsStatus `query:"status,omitzero" json:"-"`
	// Any of "article_extractor", "community_extractor",
	// "community_moderator_explorer", "community_post_extractor", "community_search",
	// "follower_explorer", "following_explorer", "list_follower_explorer",
	// "list_member_extractor", "list_post_extractor", "mention_extractor",
	// "people_search", "post_extractor", "quote_extractor", "reply_extractor",
	// "repost_extractor", "space_explorer", "thread_extractor",
	// "tweet_search_extractor", "verified_follower_explorer".
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

type ExtractionListParamsStatus string

const (
	ExtractionListParamsStatusRunning   ExtractionListParamsStatus = "running"
	ExtractionListParamsStatusCompleted ExtractionListParamsStatus = "completed"
	ExtractionListParamsStatusFailed    ExtractionListParamsStatus = "failed"
)

type ExtractionListParamsToolType string

const (
	ExtractionListParamsToolTypeArticleExtractor           ExtractionListParamsToolType = "article_extractor"
	ExtractionListParamsToolTypeCommunityExtractor         ExtractionListParamsToolType = "community_extractor"
	ExtractionListParamsToolTypeCommunityModeratorExplorer ExtractionListParamsToolType = "community_moderator_explorer"
	ExtractionListParamsToolTypeCommunityPostExtractor     ExtractionListParamsToolType = "community_post_extractor"
	ExtractionListParamsToolTypeCommunitySearch            ExtractionListParamsToolType = "community_search"
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
	ExtractionListParamsToolTypeVerifiedFollowerExplorer   ExtractionListParamsToolType = "verified_follower_explorer"
)

type ExtractionEstimateCostParams struct {
	// Any of "article_extractor", "community_extractor",
	// "community_moderator_explorer", "community_post_extractor", "community_search",
	// "follower_explorer", "following_explorer", "list_follower_explorer",
	// "list_member_extractor", "list_post_extractor", "mention_extractor",
	// "people_search", "post_extractor", "quote_extractor", "reply_extractor",
	// "repost_extractor", "space_explorer", "thread_extractor",
	// "tweet_search_extractor", "verified_follower_explorer".
	ToolType ExtractionEstimateCostParamsToolType `json:"toolType,omitzero" api:"required"`
	// Raw advanced search query appended as-is (tweet_search_extractor)
	AdvancedQuery param.Opt[string] `json:"advancedQuery,omitzero"`
	// Exact phrase to match (tweet_search_extractor)
	ExactPhrase param.Opt[string] `json:"exactPhrase,omitzero"`
	// Words to exclude from results (tweet_search_extractor)
	ExcludeWords      param.Opt[string] `json:"excludeWords,omitzero"`
	SearchQuery       param.Opt[string] `json:"searchQuery,omitzero"`
	TargetCommunityID param.Opt[string] `json:"targetCommunityId,omitzero"`
	TargetListID      param.Opt[string] `json:"targetListId,omitzero"`
	TargetSpaceID     param.Opt[string] `json:"targetSpaceId,omitzero"`
	TargetTweetID     param.Opt[string] `json:"targetTweetId,omitzero"`
	TargetUsername    param.Opt[string] `json:"targetUsername,omitzero"`
	paramObj
}

func (r ExtractionEstimateCostParams) MarshalJSON() (data []byte, err error) {
	type shadow ExtractionEstimateCostParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractionEstimateCostParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExtractionEstimateCostParamsToolType string

const (
	ExtractionEstimateCostParamsToolTypeArticleExtractor           ExtractionEstimateCostParamsToolType = "article_extractor"
	ExtractionEstimateCostParamsToolTypeCommunityExtractor         ExtractionEstimateCostParamsToolType = "community_extractor"
	ExtractionEstimateCostParamsToolTypeCommunityModeratorExplorer ExtractionEstimateCostParamsToolType = "community_moderator_explorer"
	ExtractionEstimateCostParamsToolTypeCommunityPostExtractor     ExtractionEstimateCostParamsToolType = "community_post_extractor"
	ExtractionEstimateCostParamsToolTypeCommunitySearch            ExtractionEstimateCostParamsToolType = "community_search"
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
	ExtractionEstimateCostParamsToolTypeVerifiedFollowerExplorer   ExtractionEstimateCostParamsToolType = "verified_follower_explorer"
)

type ExtractionExportResultsParams struct {
	// Any of "csv", "json", "md", "md-document", "pdf", "txt", "xlsx".
	Format ExtractionExportResultsParamsFormat `query:"format,omitzero" json:"-"`
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
	// Any of "article_extractor", "community_extractor",
	// "community_moderator_explorer", "community_post_extractor", "community_search",
	// "follower_explorer", "following_explorer", "list_follower_explorer",
	// "list_member_extractor", "list_post_extractor", "mention_extractor",
	// "people_search", "post_extractor", "quote_extractor", "reply_extractor",
	// "repost_extractor", "space_explorer", "thread_extractor",
	// "tweet_search_extractor", "verified_follower_explorer".
	ToolType ExtractionRunParamsToolType `json:"toolType,omitzero" api:"required"`
	// Raw advanced search query appended as-is (tweet_search_extractor)
	AdvancedQuery param.Opt[string] `json:"advancedQuery,omitzero"`
	// Exact phrase to match (tweet_search_extractor)
	ExactPhrase param.Opt[string] `json:"exactPhrase,omitzero"`
	// Words to exclude from results (tweet_search_extractor)
	ExcludeWords      param.Opt[string] `json:"excludeWords,omitzero"`
	SearchQuery       param.Opt[string] `json:"searchQuery,omitzero"`
	TargetCommunityID param.Opt[string] `json:"targetCommunityId,omitzero"`
	TargetListID      param.Opt[string] `json:"targetListId,omitzero"`
	TargetSpaceID     param.Opt[string] `json:"targetSpaceId,omitzero"`
	TargetTweetID     param.Opt[string] `json:"targetTweetId,omitzero"`
	TargetUsername    param.Opt[string] `json:"targetUsername,omitzero"`
	paramObj
}

func (r ExtractionRunParams) MarshalJSON() (data []byte, err error) {
	type shadow ExtractionRunParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractionRunParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExtractionRunParamsToolType string

const (
	ExtractionRunParamsToolTypeArticleExtractor           ExtractionRunParamsToolType = "article_extractor"
	ExtractionRunParamsToolTypeCommunityExtractor         ExtractionRunParamsToolType = "community_extractor"
	ExtractionRunParamsToolTypeCommunityModeratorExplorer ExtractionRunParamsToolType = "community_moderator_explorer"
	ExtractionRunParamsToolTypeCommunityPostExtractor     ExtractionRunParamsToolType = "community_post_extractor"
	ExtractionRunParamsToolTypeCommunitySearch            ExtractionRunParamsToolType = "community_search"
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
	ExtractionRunParamsToolTypeVerifiedFollowerExplorer   ExtractionRunParamsToolType = "verified_follower_explorer"
)
