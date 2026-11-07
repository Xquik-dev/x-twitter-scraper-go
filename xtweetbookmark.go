// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

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
)

// X write actions (tweets, likes, follows, DMs)
//
// XTweetBookmarkService contains methods and other services that help with
// interacting with the x-twitter-scraper API.
//
// You should not instantiate this service directly, and instead use the
// [NewXTweetBookmarkService] method instead.
type XTweetBookmarkService struct {
	options []option.RequestOption
}

// NewXTweetBookmarkService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewXTweetBookmarkService(opts ...option.RequestOption) (r XTweetBookmarkService) {
	r = XTweetBookmarkService{}
	r.options = opts
	return
}

// Bookmarks a post through a connected X account. Only that account sees its
// bookmarks, and GET /x/bookmarks lists them.
func (r *XTweetBookmarkService) New(ctx context.Context, id string, params XTweetBookmarkNewParams, opts ...option.RequestOption) (res *XWriteActionGetResponse, err error) {
	if !param.IsOmitted(params.IdempotencyKey) {
		opts = append(opts, option.WithHeader("Idempotency-Key", fmt.Sprintf("%v", params.IdempotencyKey)))
	}
	if !param.IsOmitted(params.Prefer) {
		opts = append(opts, option.WithHeader("Prefer", fmt.Sprintf("%v", params.Prefer.Value)))
	}
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("x/tweets/%s/bookmark", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Removes a bookmark through a connected X account.
func (r *XTweetBookmarkService) Delete(ctx context.Context, id string, params XTweetBookmarkDeleteParams, opts ...option.RequestOption) (res *XWriteActionGetResponse, err error) {
	if !param.IsOmitted(params.IdempotencyKey) {
		opts = append(opts, option.WithHeader("Idempotency-Key", fmt.Sprintf("%v", params.IdempotencyKey)))
	}
	if !param.IsOmitted(params.Prefer) {
		opts = append(opts, option.WithHeader("Prefer", fmt.Sprintf("%v", params.Prefer.Value)))
	}
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("x/tweets/%s/bookmark", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, params, &res, opts...)
	return res, err
}

type XTweetBookmarkNewParams struct {
	// X account identifier (@username or account ID)
	Account        string `json:"account" api:"required"`
	IdempotencyKey string `header:"Idempotency-Key" api:"required" json:"-"`
	// A long write answers 202 with a status link, then lands on X. wait=<seconds>
	// asks for the answer within that many seconds, up to 85. If the write is still
	// running then, the API answers 202 with Location. Poll Location until terminal
	// is true, and don't send the write again.
	Prefer param.Opt[string] `header:"Prefer,omitzero" json:"-"`
	paramObj
}

func (r XTweetBookmarkNewParams) MarshalJSON() (data []byte, err error) {
	type shadow XTweetBookmarkNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *XTweetBookmarkNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type XTweetBookmarkDeleteParams struct {
	// X account identifier (@username or account ID)
	Account        string `json:"account" api:"required"`
	IdempotencyKey string `header:"Idempotency-Key" api:"required" json:"-"`
	// A long write answers 202 with a status link, then lands on X. wait=<seconds>
	// asks for the answer within that many seconds, up to 85. If the write is still
	// running then, the API answers 202 with Location. Poll Location until terminal
	// is true, and don't send the write again.
	Prefer param.Opt[string] `header:"Prefer,omitzero" json:"-"`
	paramObj
}

func (r XTweetBookmarkDeleteParams) MarshalJSON() (data []byte, err error) {
	type shadow XTweetBookmarkDeleteParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *XTweetBookmarkDeleteParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
