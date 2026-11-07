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
// XUserBlockService contains methods and other services that help with
// interacting with the x-twitter-scraper API.
//
// You should not instantiate this service directly, and instead use the
// [NewXUserBlockService] method instead.
type XUserBlockService struct {
	options []option.RequestOption
}

// NewXUserBlockService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewXUserBlockService(opts ...option.RequestOption) (r XUserBlockService) {
	r = XUserBlockService{}
	r.options = opts
	return
}

// Blocks a user through a connected X account. X removes follows both ways, and
// unblocking doesn't restore them.
func (r *XUserBlockService) New(ctx context.Context, id string, params XUserBlockNewParams, opts ...option.RequestOption) (res *XWriteActionGetResponse, err error) {
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
	path := fmt.Sprintf("x/users/%s/block", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Unblocks a user through a connected X account.
func (r *XUserBlockService) Delete(ctx context.Context, id string, params XUserBlockDeleteParams, opts ...option.RequestOption) (res *XWriteActionGetResponse, err error) {
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
	path := fmt.Sprintf("x/users/%s/block", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, params, &res, opts...)
	return res, err
}

type XUserBlockNewParams struct {
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

func (r XUserBlockNewParams) MarshalJSON() (data []byte, err error) {
	type shadow XUserBlockNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *XUserBlockNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type XUserBlockDeleteParams struct {
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

func (r XUserBlockDeleteParams) MarshalJSON() (data []byte, err error) {
	type shadow XUserBlockDeleteParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *XUserBlockDeleteParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
