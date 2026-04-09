// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package xtwitterscraper

import (
	"context"
	"errors"
	"fmt"
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
	"github.com/stainless-sdks/x-twitter-scraper-go/shared"
)

// Activity events from monitored accounts
//
// EventService contains methods and other services that help with interacting with
// the x-twitter-scraper API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEventService] method instead.
type EventService struct {
	options []option.RequestOption
}

// NewEventService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewEventService(opts ...option.RequestOption) (r EventService) {
	r = EventService{}
	r.options = opts
	return
}

// Get event
func (r *EventService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *EventDetail, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("events/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// List events
func (r *EventService) List(ctx context.Context, query EventListParams, opts ...option.RequestOption) (res *EventListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "events"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Monitor event summary with type, username, and occurrence time.
type Event struct {
	ID         string         `json:"id" api:"required"`
	Data       map[string]any `json:"data" api:"required"`
	MonitorID  string         `json:"monitorId" api:"required"`
	OccurredAt time.Time      `json:"occurredAt" api:"required" format:"date-time"`
	// Type of monitor event fired when account activity occurs.
	//
	// Any of "tweet.new", "tweet.reply", "tweet.retweet", "tweet.quote",
	// "follower.gained", "follower.lost".
	Type     shared.EventType `json:"type" api:"required"`
	Username string           `json:"username" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Data        respjson.Field
		MonitorID   respjson.Field
		OccurredAt  respjson.Field
		Type        respjson.Field
		Username    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Event) RawJSON() string { return r.JSON.raw }
func (r *Event) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Full monitor event including payload data and optional X event ID.
type EventDetail struct {
	ID string `json:"id" api:"required"`
	// Event payload — shape varies by event type (JSON)
	Data       map[string]any `json:"data" api:"required"`
	MonitorID  string         `json:"monitorId" api:"required"`
	OccurredAt time.Time      `json:"occurredAt" api:"required" format:"date-time"`
	// Type of monitor event fired when account activity occurs.
	//
	// Any of "tweet.new", "tweet.reply", "tweet.retweet", "tweet.quote",
	// "follower.gained", "follower.lost".
	Type     shared.EventType `json:"type" api:"required"`
	Username string           `json:"username" api:"required"`
	XEventID string           `json:"xEventId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Data        respjson.Field
		MonitorID   respjson.Field
		OccurredAt  respjson.Field
		Type        respjson.Field
		Username    respjson.Field
		XEventID    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EventDetail) RawJSON() string { return r.JSON.raw }
func (r *EventDetail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EventListResponse struct {
	Events     []Event `json:"events" api:"required"`
	HasMore    bool    `json:"hasMore" api:"required"`
	NextCursor string  `json:"nextCursor"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Events      respjson.Field
		HasMore     respjson.Field
		NextCursor  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EventListResponse) RawJSON() string { return r.JSON.raw }
func (r *EventListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EventListParams struct {
	// Cursor for keyset pagination
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// Maximum number of items to return (1-100, default 50)
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Filter events by monitor ID
	MonitorID param.Opt[string] `query:"monitorId,omitzero" json:"-"`
	// Filter events by type
	//
	// Any of "tweet.new", "tweet.reply", "tweet.retweet", "tweet.quote",
	// "follower.gained", "follower.lost".
	EventType shared.EventType `query:"eventType,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [EventListParams]'s query parameters as `url.Values`.
func (r EventListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
