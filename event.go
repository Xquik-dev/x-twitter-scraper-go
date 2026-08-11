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
	"github.com/Xquik-dev/x-twitter-scraper-go/shared"
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

// Monitor event summary with source metadata and occurrence time.
type Event struct {
	ID   string         `json:"id" api:"required"`
	Data map[string]any `json:"data" api:"required"`
	// Account monitor ID for account events, or keyword monitor ID for keyword events.
	MonitorID string `json:"monitorId" api:"required"`
	// Source monitor type.
	//
	// Any of "account", "keyword".
	MonitorType EventMonitorType `json:"monitorType" api:"required"`
	OccurredAt  time.Time        `json:"occurredAt" api:"required" format:"date-time"`
	// Type of monitor event fired when account activity occurs.
	//
	// Any of "tweet.new", "tweet.reply", "tweet.retweet", "tweet.quote",
	// "tweet.media", "tweet.link", "tweet.poll", "tweet.mention", "tweet.hashtag",
	// "tweet.longform", "profile.avatar.changed", "profile.banner.changed",
	// "profile.name.changed", "profile.username.changed", "profile.bio.changed",
	// "profile.location.changed", "profile.url.changed", "profile.verified.changed",
	// "profile.protected.changed", "profile.pinned_tweet.changed",
	// "profile.unavailable.changed".
	Type shared.EventType `json:"type" api:"required"`
	// Keyword monitor ID, present for keyword monitor events.
	KeywordMonitorID string `json:"keywordMonitorId"`
	// Keyword query, present for keyword monitor events.
	Query string `json:"query"`
	// Account username, present for account monitor events.
	Username string `json:"username"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		Data             respjson.Field
		MonitorID        respjson.Field
		MonitorType      respjson.Field
		OccurredAt       respjson.Field
		Type             respjson.Field
		KeywordMonitorID respjson.Field
		Query            respjson.Field
		Username         respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Event) RawJSON() string { return r.JSON.raw }
func (r *Event) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Source monitor type.
type EventMonitorType string

const (
	EventMonitorTypeAccount EventMonitorType = "account"
	EventMonitorTypeKeyword EventMonitorType = "keyword"
)

// Full monitor event including payload data and optional X event ID.
type EventDetail struct {
	ID string `json:"id" api:"required"`
	// Event payload - shape varies by event type (JSON)
	Data map[string]any `json:"data" api:"required"`
	// Monitor ID associated with this detailed event payload.
	MonitorID string `json:"monitorId" api:"required"`
	// Source monitor type for this detailed event.
	//
	// Any of "account", "keyword".
	MonitorType EventDetailMonitorType `json:"monitorType" api:"required"`
	OccurredAt  time.Time              `json:"occurredAt" api:"required" format:"date-time"`
	// Type of monitor event fired when account activity occurs.
	//
	// Any of "tweet.new", "tweet.reply", "tweet.retweet", "tweet.quote",
	// "tweet.media", "tweet.link", "tweet.poll", "tweet.mention", "tweet.hashtag",
	// "tweet.longform", "profile.avatar.changed", "profile.banner.changed",
	// "profile.name.changed", "profile.username.changed", "profile.bio.changed",
	// "profile.location.changed", "profile.url.changed", "profile.verified.changed",
	// "profile.protected.changed", "profile.pinned_tweet.changed",
	// "profile.unavailable.changed".
	Type shared.EventType `json:"type" api:"required"`
	// Keyword monitor ID included on detailed keyword events.
	KeywordMonitorID string `json:"keywordMonitorId"`
	// Keyword query for this detailed monitor event.
	Query string `json:"query"`
	// Account username for this detailed monitor event.
	Username string `json:"username"`
	XEventID string `json:"xEventId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		Data             respjson.Field
		MonitorID        respjson.Field
		MonitorType      respjson.Field
		OccurredAt       respjson.Field
		Type             respjson.Field
		KeywordMonitorID respjson.Field
		Query            respjson.Field
		Username         respjson.Field
		XEventID         respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EventDetail) RawJSON() string { return r.JSON.raw }
func (r *EventDetail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Source monitor type for this detailed event.
type EventDetailMonitorType string

const (
	EventDetailMonitorTypeAccount EventDetailMonitorType = "account"
	EventDetailMonitorTypeKeyword EventDetailMonitorType = "keyword"
)

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
	// Previous nextCursor.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Keyword monitor ID.
	KeywordMonitorID param.Opt[string] `query:"keywordMonitorId,omitzero" json:"-"`
	// Maximum number of items to return (1-100, default 50). For paid per-result
	// endpoints, the returned count may be lower when remaining credits cannot cover
	// the requested page. If zero paid results are affordable, the endpoint returns
	// 402 insufficient_credits.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Account monitor ID.
	MonitorID param.Opt[string] `query:"monitorId,omitzero" json:"-"`
	// Filter events by type
	//
	// Any of "tweet.new", "tweet.reply", "tweet.retweet", "tweet.quote",
	// "tweet.media", "tweet.link", "tweet.poll", "tweet.mention", "tweet.hashtag",
	// "tweet.longform", "profile.avatar.changed", "profile.banner.changed",
	// "profile.name.changed", "profile.username.changed", "profile.bio.changed",
	// "profile.location.changed", "profile.url.changed", "profile.verified.changed",
	// "profile.protected.changed", "profile.pinned_tweet.changed",
	// "profile.unavailable.changed".
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
