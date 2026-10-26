// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

package xtwitterscraper

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/Xquik-dev/x-twitter-scraper-go/internal/apijson"
	"github.com/Xquik-dev/x-twitter-scraper-go/internal/apiquery"
	"github.com/Xquik-dev/x-twitter-scraper-go/internal/requestconfig"
	"github.com/Xquik-dev/x-twitter-scraper-go/option"
	"github.com/Xquik-dev/x-twitter-scraper-go/packages/param"
	"github.com/Xquik-dev/x-twitter-scraper-go/packages/respjson"
)

// XFollowerService lets you look up, search, and explore user profiles and relationships.
//
// XFollowerService contains methods and other services that help with interacting
// with the x-twitter-scraper API.
//
// You should not instantiate this service directly, and instead use the
// [NewXFollowerService] method instead.
type XFollowerService struct {
	options []option.RequestOption
}

// NewXFollowerService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewXFollowerService(opts ...option.RequestOption) (r XFollowerService) {
	r = XFollowerService{}
	r.options = opts
	return
}

// Check returns whether one public account follows another.
func (r *XFollowerService) Check(ctx context.Context, query XFollowerCheckParams, opts ...option.RequestOption) (res *XFollowerCheckResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "x/followers/check"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type XFollowerCheckResponse struct {
	IsFollowedBy   bool   `json:"isFollowedBy" api:"required"`
	IsFollowing    bool   `json:"isFollowing" api:"required"`
	SourceUsername string `json:"sourceUsername" api:"required"`
	TargetUsername string `json:"targetUsername" api:"required"`
	// The source account's user ID, as X gives it.
	SourceID string `json:"sourceId"`
	// The target account's user ID, as X gives it.
	TargetID string `json:"targetId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IsFollowedBy   respjson.Field
		IsFollowing    respjson.Field
		SourceUsername respjson.Field
		TargetUsername respjson.Field
		SourceID       respjson.Field
		TargetID       respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// RawJSON returns the unmodified JSON received from the API.
func (r XFollowerCheckResponse) RawJSON() string { return r.JSON.raw }
func (r *XFollowerCheckResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type XFollowerCheckParams struct {
	// Source username, @username, or X or Twitter profile URL
	Source string `query:"source" api:"required" json:"-"`
	// Target username, @username, or X or Twitter profile URL
	Target string `query:"target" api:"required" json:"-"`
	// Fields to keep in each row, comma-separated. Use dots for nested fields, such as
	// author.username. A list keeps the field in each item. Case & underscores do not
	// matter. A row that lacks a field leaves it out. Cursors, counts & diagnostics
	// always stay. An unknown name answers 400 unknown_field with the valid names,
	// before any read. The price stays the same. SDK row types then hold only the
	// named fields.
	Fields param.Opt[string] `query:"fields,omitzero" json:"-"`
	// Fields to drop from each row, comma-separated. It takes the same names as
	// fields. With fields, it drops from the kept fields.
	Omit param.Opt[string] `query:"omit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [XFollowerCheckParams]'s query parameters as `url.Values`.
func (r XFollowerCheckParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
