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
	"time"

	"github.com/Xquik-dev/x-twitter-scraper-go/internal/apijson"
	"github.com/Xquik-dev/x-twitter-scraper-go/internal/apiquery"
	"github.com/Xquik-dev/x-twitter-scraper-go/internal/requestconfig"
	"github.com/Xquik-dev/x-twitter-scraper-go/option"
	"github.com/Xquik-dev/x-twitter-scraper-go/packages/param"
	"github.com/Xquik-dev/x-twitter-scraper-go/packages/respjson"
	"github.com/Xquik-dev/x-twitter-scraper-go/shared"
)

// X live video details, watch counts, and broadcaster
//
// XBroadcastService contains methods and other services that help with interacting
// with the x-twitter-scraper API.
//
// You should not instantiate this service directly, and instead use the
// [NewXBroadcastService] method instead.
type XBroadcastService struct {
	options []option.RequestOption
}

// NewXBroadcastService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewXBroadcastService(opts ...option.RequestOption) (r XBroadcastService) {
	r = XBroadcastService{}
	r.options = opts
	return
}

// Returns 1 X broadcast, the live video X hosts at x.com/i/broadcasts, live or
// ended. Accepts the broadcast ID or the URL-encoded broadcast URL. Reading a
// broadcast adds no view to it.
func (r *XBroadcastService) Get(ctx context.Context, id string, query XBroadcastGetParams, opts ...option.RequestOption) (res *XBroadcastGetResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("x/broadcasts/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type XBroadcastGetResponse struct {
	// X broadcast, a live video
	Broadcast XBroadcastGetResponseBroadcast `json:"broadcast" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Broadcast   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XBroadcastGetResponse) RawJSON() string { return r.JSON.raw }
func (r *XBroadcastGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// X broadcast, a live video
type XBroadcastGetResponseBroadcast struct {
	// Broadcast ID
	ID string `json:"id" api:"required"`
	// Camera rotation in degrees
	CameraRotation int64 `json:"cameraRotation"`
	// X's code for who may chat
	ChatOption int64 `json:"chatOption"`
	// Public X profile.
	Creator shared.UserProfile `json:"creator"`
	// When the broadcast ended
	EndedAt time.Time `json:"endedAt" format:"date-time"`
	// Video height in pixels
	Height int64 `json:"height"`
	// Whether X offers a recording of the broadcast
	IsAvailableForReplay bool `json:"isAvailableForReplay"`
	// Whether the video stream runs in high-latency mode
	IsHighLatency bool `json:"isHighLatency"`
	// Whether X marks the chat private
	IsPrivateChat bool `json:"isPrivateChat"`
	// Whether the broadcaster edited the title of the replay
	IsTitleEdited bool `json:"isTitleEdited"`
	// When the video stream last reached X
	LastPingAt time.Time `json:"lastPingAt" format:"date-time"`
	// X's media key for the video
	MediaKey string `json:"mediaKey"`
	// Broadcaster, as X's live video service names the account
	PeriscopeUser XBroadcastGetResponseBroadcastPeriscopeUser `json:"periscopeUser"`
	// Image X shows before the broadcast starts
	PreLiveSlateURL string `json:"preLiveSlateUrl" format:"uri"`
	// When the broadcast was set to start
	ScheduledStart time.Time `json:"scheduledStart" format:"date-time"`
	// What produced the video stream, as X names it
	Source string `json:"source"`
	// When the broadcast went live
	StartedAt time.Time `json:"startedAt" format:"date-time"`
	// State, such as NotStarted, Running or Ended, as X names it
	State string `json:"state"`
	// Broadcast title
	Title string `json:"title"`
	// Accounts that watched, live or after
	TotalWatched int64 `json:"totalWatched"`
	// Accounts watching, as X counts them
	TotalWatching int64 `json:"totalWatching"`
	// Search tweet with author. Zero metrics may mean X omitted counts.
	Tweet shared.SearchTweet `json:"tweet"`
	// ID of the post that announces the broadcast
	TweetID string `json:"tweetId"`
	// X's revision number of the broadcast
	Version int64 `json:"version"`
	// Video width in pixels
	Width int64 `json:"width"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                   respjson.Field
		CameraRotation       respjson.Field
		ChatOption           respjson.Field
		Creator              respjson.Field
		EndedAt              respjson.Field
		Height               respjson.Field
		IsAvailableForReplay respjson.Field
		IsHighLatency        respjson.Field
		IsPrivateChat        respjson.Field
		IsTitleEdited        respjson.Field
		LastPingAt           respjson.Field
		MediaKey             respjson.Field
		PeriscopeUser        respjson.Field
		PreLiveSlateURL      respjson.Field
		ScheduledStart       respjson.Field
		Source               respjson.Field
		StartedAt            respjson.Field
		State                respjson.Field
		Title                respjson.Field
		TotalWatched         respjson.Field
		TotalWatching        respjson.Field
		Tweet                respjson.Field
		TweetID              respjson.Field
		Version              respjson.Field
		Width                respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XBroadcastGetResponseBroadcast) RawJSON() string { return r.JSON.raw }
func (r *XBroadcastGetResponseBroadcast) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Broadcaster, as X's live video service names the account
type XBroadcastGetResponseBroadcastPeriscopeUser struct {
	// Live video account ID
	ID string `json:"id"`
	// Live video display name
	DisplayName string `json:"displayName"`
	// Live video avatar URL
	ProfilePicture string `json:"profilePicture" format:"uri"`
	// The broadcaster's avatar at full size
	ProfilePictureOriginal string `json:"profilePictureOriginal" format:"uri"`
	// Live video username
	Username string `json:"username"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                     respjson.Field
		DisplayName            respjson.Field
		ProfilePicture         respjson.Field
		ProfilePictureOriginal respjson.Field
		Username               respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XBroadcastGetResponseBroadcastPeriscopeUser) RawJSON() string { return r.JSON.raw }
func (r *XBroadcastGetResponseBroadcastPeriscopeUser) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type XBroadcastGetParams struct {
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

// URLQuery serializes [XBroadcastGetParams]'s query parameters as `url.Values`.
func (r XBroadcastGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
