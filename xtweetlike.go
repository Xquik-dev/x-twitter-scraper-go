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
// XTweetLikeService contains methods and other services that help with interacting
// with the x-twitter-scraper API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewXTweetLikeService] method instead.
type XTweetLikeService struct {
	options []option.RequestOption
}

// NewXTweetLikeService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewXTweetLikeService(opts ...option.RequestOption) (r XTweetLikeService) {
	r = XTweetLikeService{}
	r.options = opts
	return
}

// Like tweet
func (r *XTweetLikeService) New(ctx context.Context, tweetID string, body XTweetLikeNewParams, opts ...option.RequestOption) (res *XTweetLikeNewResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if tweetID == "" {
		err = errors.New("missing required tweetId parameter")
		return nil, err
	}
	path := fmt.Sprintf("x/tweets/%s/like", url.PathEscape(tweetID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Unlike tweet
func (r *XTweetLikeService) Delete(ctx context.Context, tweetID string, body XTweetLikeDeleteParams, opts ...option.RequestOption) (res *XTweetLikeDeleteResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if tweetID == "" {
		err = errors.New("missing required tweetId parameter")
		return nil, err
	}
	path := fmt.Sprintf("x/tweets/%s/like", url.PathEscape(tweetID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return res, err
}

type XTweetLikeNewResponse struct {
	Success bool `json:"success" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XTweetLikeNewResponse) RawJSON() string { return r.JSON.raw }
func (r *XTweetLikeNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type XTweetLikeDeleteResponse struct {
	Success bool `json:"success" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XTweetLikeDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *XTweetLikeDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type XTweetLikeNewParams struct {
	// X account (@username or account ID)
	Account string `json:"account" api:"required"`
	paramObj
}

func (r XTweetLikeNewParams) MarshalJSON() (data []byte, err error) {
	type shadow XTweetLikeNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *XTweetLikeNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type XTweetLikeDeleteParams struct {
	// X account (@username or account ID)
	Account string `json:"account" api:"required"`
	paramObj
}

func (r XTweetLikeDeleteParams) MarshalJSON() (data []byte, err error) {
	type shadow XTweetLikeDeleteParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *XTweetLikeDeleteParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
