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
	"github.com/stainless-sdks/x-twitter-scraper-go/internal/requestconfig"
	"github.com/stainless-sdks/x-twitter-scraper-go/option"
	"github.com/stainless-sdks/x-twitter-scraper-go/packages/param"
	"github.com/stainless-sdks/x-twitter-scraper-go/packages/respjson"
)

// X write actions (tweets, likes, follows, DMs)
//
// XTweetRetweetService contains methods and other services that help with
// interacting with the x-twitter-scraper API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewXTweetRetweetService] method instead.
type XTweetRetweetService struct {
	options []option.RequestOption
}

// NewXTweetRetweetService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewXTweetRetweetService(opts ...option.RequestOption) (r XTweetRetweetService) {
	r = XTweetRetweetService{}
	r.options = opts
	return
}

// Retweet
func (r *XTweetRetweetService) New(ctx context.Context, tweetID string, body XTweetRetweetNewParams, opts ...option.RequestOption) (res *XTweetRetweetNewResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if tweetID == "" {
		err = errors.New("missing required tweetId parameter")
		return nil, err
	}
	path := fmt.Sprintf("x/tweets/%s/retweet", url.PathEscape(tweetID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Unretweet
func (r *XTweetRetweetService) Delete(ctx context.Context, tweetID string, body XTweetRetweetDeleteParams, opts ...option.RequestOption) (res *XTweetRetweetDeleteResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if tweetID == "" {
		err = errors.New("missing required tweetId parameter")
		return nil, err
	}
	path := fmt.Sprintf("x/tweets/%s/retweet", url.PathEscape(tweetID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return res, err
}

type XTweetRetweetNewResponse struct {
	Success bool `json:"success" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XTweetRetweetNewResponse) RawJSON() string { return r.JSON.raw }
func (r *XTweetRetweetNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type XTweetRetweetDeleteResponse struct {
	Success bool `json:"success" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XTweetRetweetDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *XTweetRetweetDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type XTweetRetweetNewParams struct {
	// X account (@username or account ID)
	Account string `json:"account" api:"required"`
	paramObj
}

func (r XTweetRetweetNewParams) MarshalJSON() (data []byte, err error) {
	type shadow XTweetRetweetNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *XTweetRetweetNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type XTweetRetweetDeleteParams struct {
	// X account (@username or account ID)
	Account string `json:"account" api:"required"`
	paramObj
}

func (r XTweetRetweetDeleteParams) MarshalJSON() (data []byte, err error) {
	type shadow XTweetRetweetDeleteParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *XTweetRetweetDeleteParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
