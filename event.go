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

	"github.com/Xquik-dev/x-twitter-scraper-go/internal/apijson"
	"github.com/Xquik-dev/x-twitter-scraper-go/internal/apiquery"
	"github.com/Xquik-dev/x-twitter-scraper-go/internal/requestconfig"
	"github.com/Xquik-dev/x-twitter-scraper-go/option"
	"github.com/Xquik-dev/x-twitter-scraper-go/packages/param"
	"github.com/Xquik-dev/x-twitter-scraper-go/packages/respjson"
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
func (r *EventService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *EventGetResponse, err error) {
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
	var preClientOpts = []option.RequestOption{requestconfig.WithAPIKeySecurity()}
	opts = slices.Concat(preClientOpts, r.options, opts)
	path := "events"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type EventGetResponse struct {
	ID         string         `json:"id" api:"required"`
	Data       map[string]any `json:"data" api:"required"`
	MonitorID  string         `json:"monitorId" api:"required"`
	OccurredAt time.Time      `json:"occurredAt" api:"required" format:"date-time"`
	// Any of "tweet.new", "tweet.reply", "tweet.retweet", "tweet.quote",
	// "follower.gained", "follower.lost".
	Type     EventGetResponseType `json:"type" api:"required"`
	Username string               `json:"username" api:"required"`
	XEventID string               `json:"xEventId"`
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
func (r EventGetResponse) RawJSON() string { return r.JSON.raw }
func (r *EventGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EventGetResponseType string

const (
	EventGetResponseTypeTweetNew       EventGetResponseType = "tweet.new"
	EventGetResponseTypeTweetReply     EventGetResponseType = "tweet.reply"
	EventGetResponseTypeTweetRetweet   EventGetResponseType = "tweet.retweet"
	EventGetResponseTypeTweetQuote     EventGetResponseType = "tweet.quote"
	EventGetResponseTypeFollowerGained EventGetResponseType = "follower.gained"
	EventGetResponseTypeFollowerLost   EventGetResponseType = "follower.lost"
)

type EventListResponse struct {
	Events     []EventListResponseEvent `json:"events" api:"required"`
	HasMore    bool                     `json:"hasMore" api:"required"`
	NextCursor string                   `json:"nextCursor"`
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

type EventListResponseEvent struct {
	ID         string         `json:"id" api:"required"`
	Data       map[string]any `json:"data" api:"required"`
	MonitorID  string         `json:"monitorId" api:"required"`
	OccurredAt time.Time      `json:"occurredAt" api:"required" format:"date-time"`
	// Any of "tweet.new", "tweet.reply", "tweet.retweet", "tweet.quote",
	// "follower.gained", "follower.lost".
	Type     string `json:"type" api:"required"`
	Username string `json:"username" api:"required"`
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
func (r EventListResponseEvent) RawJSON() string { return r.JSON.raw }
func (r *EventListResponseEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EventListParams struct {
	// Cursor for pagination
	After     param.Opt[string] `query:"after,omitzero" json:"-"`
	Limit     param.Opt[int64]  `query:"limit,omitzero" json:"-"`
	MonitorID param.Opt[string] `query:"monitorId,omitzero" json:"-"`
	// Any of "tweet.new", "tweet.reply", "tweet.retweet", "tweet.quote",
	// "follower.gained", "follower.lost".
	EventType EventListParamsEventType `query:"eventType,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [EventListParams]'s query parameters as `url.Values`.
func (r EventListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type EventListParamsEventType string

const (
	EventListParamsEventTypeTweetNew       EventListParamsEventType = "tweet.new"
	EventListParamsEventTypeTweetReply     EventListParamsEventType = "tweet.reply"
	EventListParamsEventTypeTweetRetweet   EventListParamsEventType = "tweet.retweet"
	EventListParamsEventTypeTweetQuote     EventListParamsEventType = "tweet.quote"
	EventListParamsEventTypeFollowerGained EventListParamsEventType = "follower.gained"
	EventListParamsEventTypeFollowerLost   EventListParamsEventType = "follower.lost"
)
