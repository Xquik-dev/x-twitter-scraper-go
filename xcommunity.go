// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package xtwitterscraper

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/stainless-sdks/x-twitter-scraper-go/internal/apijson"
	"github.com/stainless-sdks/x-twitter-scraper-go/internal/apiquery"
	"github.com/stainless-sdks/x-twitter-scraper-go/internal/requestconfig"
	"github.com/stainless-sdks/x-twitter-scraper-go/option"
	"github.com/stainless-sdks/x-twitter-scraper-go/packages/param"
	"github.com/stainless-sdks/x-twitter-scraper-go/packages/respjson"
	"github.com/stainless-sdks/x-twitter-scraper-go/shared"
)

// XCommunityService contains methods and other services that help with interacting
// with the x-twitter-scraper API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewXCommunityService] method instead.
type XCommunityService struct {
	options []option.RequestOption
	// X write actions (tweets, likes, follows, DMs)
	Join XCommunityJoinService
	// X Community info, members, and tweets
	Tweets XCommunityTweetService
}

// NewXCommunityService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewXCommunityService(opts ...option.RequestOption) (r XCommunityService) {
	r = XCommunityService{}
	r.options = opts
	r.Join = NewXCommunityJoinService(opts...)
	r.Tweets = NewXCommunityTweetService(opts...)
	return
}

// Create community
func (r *XCommunityService) New(ctx context.Context, body XCommunityNewParams, opts ...option.RequestOption) (res *XCommunityNewResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "x/communities"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Delete community
func (r *XCommunityService) Delete(ctx context.Context, id string, body XCommunityDeleteParams, opts ...option.RequestOption) (res *XCommunityDeleteResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("x/communities/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return res, err
}

// Get community name, description & member count
func (r *XCommunityService) GetInfo(ctx context.Context, id string, opts ...option.RequestOption) (res *XCommunityGetInfoResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("x/communities/%s/info", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// List members of a community
func (r *XCommunityService) GetMembers(ctx context.Context, id string, query XCommunityGetMembersParams, opts ...option.RequestOption) (res *shared.PaginatedUsers, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("x/communities/%s/members", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// List moderators of a community
func (r *XCommunityService) GetModerators(ctx context.Context, id string, query XCommunityGetModeratorsParams, opts ...option.RequestOption) (res *shared.PaginatedUsers, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("x/communities/%s/moderators", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Search for communities by keyword
func (r *XCommunityService) GetSearch(ctx context.Context, query XCommunityGetSearchParams, opts ...option.RequestOption) (res *shared.PaginatedTweets, err error) {
	opts = slices.Concat(r.options, opts)
	path := "x/communities/search"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Result of a community join or leave action.
type CommunityActionResult struct {
	CommunityID   string `json:"communityId" api:"required"`
	CommunityName string `json:"communityName" api:"required"`
	Success       bool   `json:"success" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CommunityID   respjson.Field
		CommunityName respjson.Field
		Success       respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CommunityActionResult) RawJSON() string { return r.JSON.raw }
func (r *CommunityActionResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type XCommunityNewResponse struct {
	CommunityID   string `json:"communityId" api:"required"`
	Success       bool   `json:"success" api:"required"`
	CommunityName string `json:"communityName"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CommunityID   respjson.Field
		Success       respjson.Field
		CommunityName respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XCommunityNewResponse) RawJSON() string { return r.JSON.raw }
func (r *XCommunityNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type XCommunityDeleteResponse struct {
	Success bool `json:"success" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XCommunityDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *XCommunityDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type XCommunityGetInfoResponse struct {
	// Community info object
	Community XCommunityGetInfoResponseCommunity `json:"community" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Community   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XCommunityGetInfoResponse) RawJSON() string { return r.JSON.raw }
func (r *XCommunityGetInfoResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Community info object
type XCommunityGetInfoResponseCommunity struct {
	// Unique community identifier
	ID string `json:"id" api:"required"`
	// Community banner image URL
	BannerURL string `json:"banner_url"`
	// Community creation timestamp
	CreatedAt string `json:"created_at"`
	// About text for the community
	Description string `json:"description"`
	// Join policy (open or restricted)
	JoinPolicy string `json:"join_policy"`
	// Total member count
	MemberCount int64 `json:"member_count"`
	// Total moderator count
	ModeratorCount int64 `json:"moderator_count"`
	// Display name of the community
	Name string `json:"name"`
	// Primary topic
	PrimaryTopic XCommunityGetInfoResponseCommunityPrimaryTopic `json:"primary_topic"`
	// Community rules
	Rules []XCommunityGetInfoResponseCommunityRule `json:"rules"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		BannerURL      respjson.Field
		CreatedAt      respjson.Field
		Description    respjson.Field
		JoinPolicy     respjson.Field
		MemberCount    respjson.Field
		ModeratorCount respjson.Field
		Name           respjson.Field
		PrimaryTopic   respjson.Field
		Rules          respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XCommunityGetInfoResponseCommunity) RawJSON() string { return r.JSON.raw }
func (r *XCommunityGetInfoResponseCommunity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Primary topic
type XCommunityGetInfoResponseCommunityPrimaryTopic struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XCommunityGetInfoResponseCommunityPrimaryTopic) RawJSON() string { return r.JSON.raw }
func (r *XCommunityGetInfoResponseCommunityPrimaryTopic) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type XCommunityGetInfoResponseCommunityRule struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	Name        string `json:"name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Description respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XCommunityGetInfoResponseCommunityRule) RawJSON() string { return r.JSON.raw }
func (r *XCommunityGetInfoResponseCommunityRule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type XCommunityNewParams struct {
	// X account (@username or ID) creating the community
	Account string `json:"account" api:"required"`
	// Community name
	Name string `json:"name" api:"required"`
	// Community description
	Description param.Opt[string] `json:"description,omitzero"`
	paramObj
}

func (r XCommunityNewParams) MarshalJSON() (data []byte, err error) {
	type shadow XCommunityNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *XCommunityNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type XCommunityDeleteParams struct {
	// X account (@username or ID) deleting the community
	Account string `json:"account" api:"required"`
	// Community name for confirmation
	CommunityName string `json:"community_name" api:"required"`
	paramObj
}

func (r XCommunityDeleteParams) MarshalJSON() (data []byte, err error) {
	type shadow XCommunityDeleteParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *XCommunityDeleteParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type XCommunityGetMembersParams struct {
	// Pagination cursor
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [XCommunityGetMembersParams]'s query parameters as
// `url.Values`.
func (r XCommunityGetMembersParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type XCommunityGetModeratorsParams struct {
	// Pagination cursor for community moderators
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [XCommunityGetModeratorsParams]'s query parameters as
// `url.Values`.
func (r XCommunityGetModeratorsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type XCommunityGetSearchParams struct {
	// Search query
	Q string `query:"q" api:"required" json:"-"`
	// Pagination cursor for community search
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Sort order (Latest or Top)
	QueryType param.Opt[string] `query:"queryType,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [XCommunityGetSearchParams]'s query parameters as
// `url.Values`.
func (r XCommunityGetSearchParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
