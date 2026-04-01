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

	"github.com/stainless-sdks/x-twitter-scraper-go/internal/apijson"
	"github.com/stainless-sdks/x-twitter-scraper-go/internal/apiquery"
	"github.com/stainless-sdks/x-twitter-scraper-go/internal/requestconfig"
	"github.com/stainless-sdks/x-twitter-scraper-go/option"
	"github.com/stainless-sdks/x-twitter-scraper-go/packages/param"
	"github.com/stainless-sdks/x-twitter-scraper-go/packages/respjson"
)

// Tweet composition, drafts, writing styles & radar
//
// DraftService contains methods and other services that help with interacting with
// the x-twitter-scraper API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDraftService] method instead.
type DraftService struct {
	options []option.RequestOption
}

// NewDraftService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewDraftService(opts ...option.RequestOption) (r DraftService) {
	r = DraftService{}
	r.options = opts
	return
}

// Save a tweet draft
func (r *DraftService) New(ctx context.Context, body DraftNewParams, opts ...option.RequestOption) (res *DraftNewResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "drafts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Get draft by ID
func (r *DraftService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *DraftGetResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("drafts/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// List saved drafts
func (r *DraftService) List(ctx context.Context, query DraftListParams, opts ...option.RequestOption) (res *DraftListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "drafts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Delete a draft
func (r *DraftService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("drafts/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

type DraftNewResponse struct {
	ID        string    `json:"id" api:"required"`
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	Text      string    `json:"text" api:"required"`
	UpdatedAt time.Time `json:"updatedAt" api:"required" format:"date-time"`
	Goal      string    `json:"goal"`
	Topic     string    `json:"topic"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Text        respjson.Field
		UpdatedAt   respjson.Field
		Goal        respjson.Field
		Topic       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DraftNewResponse) RawJSON() string { return r.JSON.raw }
func (r *DraftNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DraftGetResponse struct {
	ID        string    `json:"id" api:"required"`
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	Text      string    `json:"text" api:"required"`
	UpdatedAt time.Time `json:"updatedAt" api:"required" format:"date-time"`
	Goal      string    `json:"goal"`
	Topic     string    `json:"topic"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Text        respjson.Field
		UpdatedAt   respjson.Field
		Goal        respjson.Field
		Topic       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DraftGetResponse) RawJSON() string { return r.JSON.raw }
func (r *DraftGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DraftListResponse struct {
	Drafts     []DraftListResponseDraft `json:"drafts" api:"required"`
	HasMore    bool                     `json:"hasMore" api:"required"`
	NextCursor string                   `json:"nextCursor"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Drafts      respjson.Field
		HasMore     respjson.Field
		NextCursor  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DraftListResponse) RawJSON() string { return r.JSON.raw }
func (r *DraftListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DraftListResponseDraft struct {
	ID        string    `json:"id" api:"required"`
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	Text      string    `json:"text" api:"required"`
	Goal      string    `json:"goal"`
	Topic     string    `json:"topic"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Text        respjson.Field
		Goal        respjson.Field
		Topic       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DraftListResponseDraft) RawJSON() string { return r.JSON.raw }
func (r *DraftListResponseDraft) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DraftNewParams struct {
	Text  string            `json:"text" api:"required"`
	Topic param.Opt[string] `json:"topic,omitzero"`
	// Any of "engagement", "followers", "authority", "conversation".
	Goal DraftNewParamsGoal `json:"goal,omitzero"`
	paramObj
}

func (r DraftNewParams) MarshalJSON() (data []byte, err error) {
	type shadow DraftNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DraftNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DraftNewParamsGoal string

const (
	DraftNewParamsGoalEngagement   DraftNewParamsGoal = "engagement"
	DraftNewParamsGoalFollowers    DraftNewParamsGoal = "followers"
	DraftNewParamsGoalAuthority    DraftNewParamsGoal = "authority"
	DraftNewParamsGoalConversation DraftNewParamsGoal = "conversation"
)

type DraftListParams struct {
	// Cursor for pagination
	AfterCursor param.Opt[string] `query:"afterCursor,omitzero" json:"-"`
	Limit       param.Opt[int64]  `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [DraftListParams]'s query parameters as `url.Values`.
func (r DraftListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
