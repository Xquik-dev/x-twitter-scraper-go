// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package xtwitterscraper

import (
	"context"
	"net/http"
	"slices"

	"github.com/stainless-sdks/x-twitter-scraper-go/internal/apijson"
	"github.com/stainless-sdks/x-twitter-scraper-go/internal/requestconfig"
	"github.com/stainless-sdks/x-twitter-scraper-go/option"
	"github.com/stainless-sdks/x-twitter-scraper-go/packages/param"
	"github.com/stainless-sdks/x-twitter-scraper-go/packages/respjson"
)

// Telegram bot service endpoints
//
// BotService contains methods and other services that help with interacting with
// the x-twitter-scraper API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBotService] method instead.
type BotService struct {
	options []option.RequestOption
	// Telegram bot service endpoints
	PlatformLinks BotPlatformLinkService
}

// NewBotService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewBotService(opts ...option.RequestOption) (r BotService) {
	r = BotService{}
	r.options = opts
	r.PlatformLinks = NewBotPlatformLinkService(opts...)
	return
}

// Track bot token usage
func (r *BotService) TrackUsage(ctx context.Context, body BotTrackUsageParams, opts ...option.RequestOption) (res *BotTrackUsageResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "bot/usage"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type BotTrackUsageResponse struct {
	Success bool `json:"success"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BotTrackUsageResponse) RawJSON() string { return r.JSON.raw }
func (r *BotTrackUsageResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BotTrackUsageParams struct {
	InputTokens    int64  `json:"inputTokens" api:"required"`
	OutputTokens   int64  `json:"outputTokens" api:"required"`
	PlatformUserID string `json:"platformUserId" api:"required"`
	paramObj
}

func (r BotTrackUsageParams) MarshalJSON() (data []byte, err error) {
	type shadow BotTrackUsageParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BotTrackUsageParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
