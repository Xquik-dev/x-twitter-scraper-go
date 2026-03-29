// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package xtwitterscraper

import (
	"context"
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

// Telegram bot service endpoints
//
// BotPlatformLinkService contains methods and other services that help with
// interacting with the x-twitter-scraper API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBotPlatformLinkService] method instead.
type BotPlatformLinkService struct {
	options []option.RequestOption
}

// NewBotPlatformLinkService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewBotPlatformLinkService(opts ...option.RequestOption) (r BotPlatformLinkService) {
	r = BotPlatformLinkService{}
	r.options = opts
	return
}

// Link a platform user to an Xquik account
func (r *BotPlatformLinkService) New(ctx context.Context, body BotPlatformLinkNewParams, opts ...option.RequestOption) (res *BotPlatformLinkNewResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "bot/platform-links"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Unlink a platform user from an Xquik account
func (r *BotPlatformLinkService) Delete(ctx context.Context, body BotPlatformLinkDeleteParams, opts ...option.RequestOption) (res *BotPlatformLinkDeleteResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "bot/platform-links"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return res, err
}

// Look up an Xquik user by platform identity
func (r *BotPlatformLinkService) Lookup(ctx context.Context, query BotPlatformLinkLookupParams, opts ...option.RequestOption) (res *BotPlatformLinkLookupResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "bot/platform-links/lookup"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type BotPlatformLinkNewResponse struct {
	ID             string    `json:"id" api:"required"`
	CreatedAt      time.Time `json:"createdAt" api:"required" format:"date-time"`
	Platform       string    `json:"platform" api:"required"`
	PlatformUserID string    `json:"platformUserId" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		CreatedAt      respjson.Field
		Platform       respjson.Field
		PlatformUserID respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BotPlatformLinkNewResponse) RawJSON() string { return r.JSON.raw }
func (r *BotPlatformLinkNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BotPlatformLinkDeleteResponse struct {
	Success bool `json:"success"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BotPlatformLinkDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *BotPlatformLinkDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BotPlatformLinkLookupResponse struct {
	UserID string `json:"userId" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		UserID      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BotPlatformLinkLookupResponse) RawJSON() string { return r.JSON.raw }
func (r *BotPlatformLinkLookupResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BotPlatformLinkNewParams struct {
	// Any of "telegram".
	Platform       BotPlatformLinkNewParamsPlatform `json:"platform,omitzero" api:"required"`
	PlatformUserID string                           `json:"platformUserId" api:"required"`
	paramObj
}

func (r BotPlatformLinkNewParams) MarshalJSON() (data []byte, err error) {
	type shadow BotPlatformLinkNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BotPlatformLinkNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BotPlatformLinkNewParamsPlatform string

const (
	BotPlatformLinkNewParamsPlatformTelegram BotPlatformLinkNewParamsPlatform = "telegram"
)

type BotPlatformLinkDeleteParams struct {
	// Any of "telegram".
	Platform       BotPlatformLinkDeleteParamsPlatform `json:"platform,omitzero" api:"required"`
	PlatformUserID string                              `json:"platformUserId" api:"required"`
	paramObj
}

func (r BotPlatformLinkDeleteParams) MarshalJSON() (data []byte, err error) {
	type shadow BotPlatformLinkDeleteParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BotPlatformLinkDeleteParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BotPlatformLinkDeleteParamsPlatform string

const (
	BotPlatformLinkDeleteParamsPlatformTelegram BotPlatformLinkDeleteParamsPlatform = "telegram"
)

type BotPlatformLinkLookupParams struct {
	Platform       string `query:"platform" api:"required" json:"-"`
	PlatformUserID string `query:"platformUserId" api:"required" json:"-"`
	paramObj
}

// URLQuery serializes [BotPlatformLinkLookupParams]'s query parameters as
// `url.Values`.
func (r BotPlatformLinkLookupParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
