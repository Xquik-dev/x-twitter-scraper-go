// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package xtwitterscraper

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/Xquik-dev/x-twitter-scraper-go/internal/apijson"
	"github.com/Xquik-dev/x-twitter-scraper-go/internal/requestconfig"
	"github.com/Xquik-dev/x-twitter-scraper-go/option"
	"github.com/Xquik-dev/x-twitter-scraper-go/packages/param"
	"github.com/Xquik-dev/x-twitter-scraper-go/packages/respjson"
)

// X write actions (tweets, likes, follows, DMs)
//
// XCommunityJoinService contains methods and other services that help with
// interacting with the x-twitter-scraper API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewXCommunityJoinService] method instead.
type XCommunityJoinService struct {
	options []option.RequestOption
}

// NewXCommunityJoinService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewXCommunityJoinService(opts ...option.RequestOption) (r XCommunityJoinService) {
	r = XCommunityJoinService{}
	r.options = opts
	return
}

// Join community
func (r *XCommunityJoinService) New(ctx context.Context, id string, body XCommunityJoinNewParams, opts ...option.RequestOption) (res *XCommunityJoinNewResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("x/communities/%s/join", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Leave community
func (r *XCommunityJoinService) DeleteAll(ctx context.Context, id string, body XCommunityJoinDeleteAllParams, opts ...option.RequestOption) (res *XCommunityJoinDeleteAllResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("x/communities/%s/join", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return res, err
}

// Result of a community join or leave action.
type XCommunityJoinNewResponse struct {
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
func (r XCommunityJoinNewResponse) RawJSON() string { return r.JSON.raw }
func (r *XCommunityJoinNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Result of a community join or leave action.
type XCommunityJoinDeleteAllResponse struct {
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
func (r XCommunityJoinDeleteAllResponse) RawJSON() string { return r.JSON.raw }
func (r *XCommunityJoinDeleteAllResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type XCommunityJoinNewParams struct {
	// X account identifier (@username or account ID)
	Account string `json:"account" api:"required"`
	paramObj
}

func (r XCommunityJoinNewParams) MarshalJSON() (data []byte, err error) {
	type shadow XCommunityJoinNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *XCommunityJoinNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type XCommunityJoinDeleteAllParams struct {
	// X account identifier (@username or account ID)
	Account string `json:"account" api:"required"`
	paramObj
}

func (r XCommunityJoinDeleteAllParams) MarshalJSON() (data []byte, err error) {
	type shadow XCommunityJoinDeleteAllParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *XCommunityJoinDeleteAllParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
