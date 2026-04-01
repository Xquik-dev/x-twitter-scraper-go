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
// XUserFollowService contains methods and other services that help with
// interacting with the x-twitter-scraper API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewXUserFollowService] method instead.
type XUserFollowService struct {
	options []option.RequestOption
}

// NewXUserFollowService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewXUserFollowService(opts ...option.RequestOption) (r XUserFollowService) {
	r = XUserFollowService{}
	r.options = opts
	return
}

// Follow user
func (r *XUserFollowService) New(ctx context.Context, userID string, body XUserFollowNewParams, opts ...option.RequestOption) (res *XUserFollowNewResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if userID == "" {
		err = errors.New("missing required userId parameter")
		return nil, err
	}
	path := fmt.Sprintf("x/users/%s/follow", url.PathEscape(userID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Unfollow user
func (r *XUserFollowService) DeleteAll(ctx context.Context, userID string, body XUserFollowDeleteAllParams, opts ...option.RequestOption) (res *XUserFollowDeleteAllResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if userID == "" {
		err = errors.New("missing required userId parameter")
		return nil, err
	}
	path := fmt.Sprintf("x/users/%s/follow", url.PathEscape(userID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return res, err
}

type XUserFollowNewResponse struct {
	Success bool `json:"success" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XUserFollowNewResponse) RawJSON() string { return r.JSON.raw }
func (r *XUserFollowNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type XUserFollowDeleteAllResponse struct {
	Success bool `json:"success" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XUserFollowDeleteAllResponse) RawJSON() string { return r.JSON.raw }
func (r *XUserFollowDeleteAllResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type XUserFollowNewParams struct {
	// X account (@username or account ID)
	Account string `json:"account" api:"required"`
	paramObj
}

func (r XUserFollowNewParams) MarshalJSON() (data []byte, err error) {
	type shadow XUserFollowNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *XUserFollowNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type XUserFollowDeleteAllParams struct {
	// X account (@username or account ID)
	Account string `json:"account" api:"required"`
	paramObj
}

func (r XUserFollowDeleteAllParams) MarshalJSON() (data []byte, err error) {
	type shadow XUserFollowDeleteAllParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *XUserFollowDeleteAllParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
