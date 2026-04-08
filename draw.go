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

// Giveaway draws from tweet replies
//
// DrawService contains methods and other services that help with interacting with
// the x-twitter-scraper API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDrawService] method instead.
type DrawService struct {
	options []option.RequestOption
}

// NewDrawService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewDrawService(opts ...option.RequestOption) (r DrawService) {
	r = DrawService{}
	r.options = opts
	return
}

// Get draw details
func (r *DrawService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *DrawGetResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("draws/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// List draws
func (r *DrawService) List(ctx context.Context, query DrawListParams, opts ...option.RequestOption) (res *DrawListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "draws"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Export draw data
func (r *DrawService) Export(ctx context.Context, id string, query DrawExportParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/octet-stream")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("draws/%s/export", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Run giveaway draw
func (r *DrawService) Run(ctx context.Context, body DrawRunParams, opts ...option.RequestOption) (res *DrawRunResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "draws"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Full giveaway draw with tweet metrics, entries, and timing.
type DrawDetail struct {
	ID                  string    `json:"id" api:"required"`
	CreatedAt           time.Time `json:"createdAt" api:"required" format:"date-time"`
	Status              string    `json:"status" api:"required"`
	TotalEntries        int64     `json:"totalEntries" api:"required"`
	TweetAuthorUsername string    `json:"tweetAuthorUsername" api:"required"`
	TweetID             string    `json:"tweetId" api:"required"`
	TweetLikeCount      int64     `json:"tweetLikeCount" api:"required"`
	TweetQuoteCount     int64     `json:"tweetQuoteCount" api:"required"`
	TweetReplyCount     int64     `json:"tweetReplyCount" api:"required"`
	TweetRetweetCount   int64     `json:"tweetRetweetCount" api:"required"`
	TweetText           string    `json:"tweetText" api:"required"`
	TweetURL            string    `json:"tweetUrl" api:"required" format:"uri"`
	ValidEntries        int64     `json:"validEntries" api:"required"`
	DrawnAt             time.Time `json:"drawnAt" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                  respjson.Field
		CreatedAt           respjson.Field
		Status              respjson.Field
		TotalEntries        respjson.Field
		TweetAuthorUsername respjson.Field
		TweetID             respjson.Field
		TweetLikeCount      respjson.Field
		TweetQuoteCount     respjson.Field
		TweetReplyCount     respjson.Field
		TweetRetweetCount   respjson.Field
		TweetText           respjson.Field
		TweetURL            respjson.Field
		ValidEntries        respjson.Field
		DrawnAt             respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DrawDetail) RawJSON() string { return r.JSON.raw }
func (r *DrawDetail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Giveaway draw summary with entry counts and status.
type DrawListItem struct {
	ID           string    `json:"id" api:"required"`
	CreatedAt    time.Time `json:"createdAt" api:"required" format:"date-time"`
	Status       string    `json:"status" api:"required"`
	TotalEntries int64     `json:"totalEntries" api:"required"`
	TweetURL     string    `json:"tweetUrl" api:"required" format:"uri"`
	ValidEntries int64     `json:"validEntries" api:"required"`
	DrawnAt      time.Time `json:"drawnAt" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		CreatedAt    respjson.Field
		Status       respjson.Field
		TotalEntries respjson.Field
		TweetURL     respjson.Field
		ValidEntries respjson.Field
		DrawnAt      respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DrawListItem) RawJSON() string { return r.JSON.raw }
func (r *DrawListItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Giveaway draw winner with position and backup flag.
type Winner struct {
	AuthorUsername string `json:"authorUsername" api:"required"`
	IsBackup       bool   `json:"isBackup" api:"required"`
	Position       int64  `json:"position" api:"required"`
	TweetID        string `json:"tweetId" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AuthorUsername respjson.Field
		IsBackup       respjson.Field
		Position       respjson.Field
		TweetID        respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Winner) RawJSON() string { return r.JSON.raw }
func (r *Winner) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DrawGetResponse struct {
	// Full giveaway draw with tweet metrics, entries, and timing.
	Draw    DrawDetail `json:"draw" api:"required"`
	Winners []Winner   `json:"winners" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Draw        respjson.Field
		Winners     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DrawGetResponse) RawJSON() string { return r.JSON.raw }
func (r *DrawGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DrawListResponse struct {
	Draws      []DrawListItem `json:"draws" api:"required"`
	HasMore    bool           `json:"hasMore" api:"required"`
	NextCursor string         `json:"nextCursor"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Draws       respjson.Field
		HasMore     respjson.Field
		NextCursor  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DrawListResponse) RawJSON() string { return r.JSON.raw }
func (r *DrawListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DrawRunResponse struct {
	ID           string   `json:"id" api:"required"`
	TotalEntries int64    `json:"totalEntries" api:"required"`
	TweetID      string   `json:"tweetId" api:"required"`
	ValidEntries int64    `json:"validEntries" api:"required"`
	Winners      []Winner `json:"winners" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		TotalEntries respjson.Field
		TweetID      respjson.Field
		ValidEntries respjson.Field
		Winners      respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DrawRunResponse) RawJSON() string { return r.JSON.raw }
func (r *DrawRunResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DrawListParams struct {
	// Cursor for keyset pagination
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// Maximum number of items to return (1-100, default 50)
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [DrawListParams]'s query parameters as `url.Values`.
func (r DrawListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type DrawExportParams struct {
	// Export output format
	//
	// Any of "csv", "json", "md", "md-document", "pdf", "txt", "xlsx".
	Format DrawExportParamsFormat `query:"format,omitzero" json:"-"`
	// Export winners or all entries
	//
	// Any of "winners", "entries".
	Type DrawExportParamsType `query:"type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [DrawExportParams]'s query parameters as `url.Values`.
func (r DrawExportParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Export output format
type DrawExportParamsFormat string

const (
	DrawExportParamsFormatCsv        DrawExportParamsFormat = "csv"
	DrawExportParamsFormatJson       DrawExportParamsFormat = "json"
	DrawExportParamsFormatMd         DrawExportParamsFormat = "md"
	DrawExportParamsFormatMdDocument DrawExportParamsFormat = "md-document"
	DrawExportParamsFormatPdf        DrawExportParamsFormat = "pdf"
	DrawExportParamsFormatTxt        DrawExportParamsFormat = "txt"
	DrawExportParamsFormatXlsx       DrawExportParamsFormat = "xlsx"
)

// Export winners or all entries
type DrawExportParamsType string

const (
	DrawExportParamsTypeWinners DrawExportParamsType = "winners"
	DrawExportParamsTypeEntries DrawExportParamsType = "entries"
)

type DrawRunParams struct {
	TweetURL             string            `json:"tweetUrl" api:"required" format:"uri"`
	BackupCount          param.Opt[int64]  `json:"backupCount,omitzero"`
	FilterAccountAgeDays param.Opt[int64]  `json:"filterAccountAgeDays,omitzero"`
	FilterLanguage       param.Opt[string] `json:"filterLanguage,omitzero"`
	FilterMinFollowers   param.Opt[int64]  `json:"filterMinFollowers,omitzero"`
	MustFollowUsername   param.Opt[string] `json:"mustFollowUsername,omitzero"`
	MustRetweet          param.Opt[bool]   `json:"mustRetweet,omitzero"`
	UniqueAuthorsOnly    param.Opt[bool]   `json:"uniqueAuthorsOnly,omitzero"`
	WinnerCount          param.Opt[int64]  `json:"winnerCount,omitzero"`
	RequiredHashtags     []string          `json:"requiredHashtags,omitzero"`
	RequiredKeywords     []string          `json:"requiredKeywords,omitzero"`
	RequiredMentions     []string          `json:"requiredMentions,omitzero"`
	paramObj
}

func (r DrawRunParams) MarshalJSON() (data []byte, err error) {
	type shadow DrawRunParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DrawRunParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
