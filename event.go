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
	"github.com/Xquik-dev/x-twitter-scraper-go/packages/ssestream"
	"github.com/Xquik-dev/x-twitter-scraper-go/shared"
)

// EventService covers activity events from monitored accounts.
//
// EventService contains methods and other services that help with interacting with
// the x-twitter-scraper API.
//
// You should not instantiate this service directly, and instead use the
// [NewEventService] method instead.
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

// Get returns one monitor event owned by the account.
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

// List returns monitor events with cursor pagination and filters.
func (r *EventService) List(ctx context.Context, query EventListParams, opts ...option.RequestOption) (res *EventListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "events"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// StreamStreaming streams the account's monitor events as server-sent events, each named by its
// event type and numbered by its event ID. A reconnect that sends Last-Event-ID
// first gets every event after that one from the last 30 days, then each new
// event. A connection lasts up to 30 minutes and sends a comment every 15 seconds
// while idle. Keep up to 5 streams open per account. Events and streams cost no
// credits.
func (r *EventService) StreamStreaming(ctx context.Context, params EventStreamParams, opts ...option.RequestOption) (stream *ssestream.Stream[Event]) {
	var (
		raw *http.Response
		err error
	)
	if !param.IsOmitted(params.LastEventID) {
		opts = append(opts, option.WithHeader("Last-Event-ID", fmt.Sprintf("%v", params.LastEventID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/event-stream")}, opts...)
	path := "events/stream"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &raw, opts...)
	return ssestream.NewStream[Event](ssestream.NewDecoder(raw), err)
}

// Event is a monitor event summary with source metadata and occurrence time.
type Event struct {
	ID string `json:"id" api:"required"`
	// Post or profile-change payload. Read its variant with
	// [MonitorEventDataUnion.AsAny], passing the event's Type.
	Data MonitorEventDataUnion `json:"data" api:"required"`
	// Account monitor ID for account events, or keyword monitor ID for keyword events.
	MonitorID string `json:"monitorId" api:"required"`
	// Source monitor type.
	//
	// Any of "account", "keyword", "post".
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
	// "profile.unavailable.changed", "tweet.deleted", "broadcast.started",
	// "space.started", "following.added", "following.removed", "follower.added",
	// "follower.removed", "profile.followers.changed", "profile.following.changed".
	Type shared.EventType `json:"type" api:"required"`
	// Permanent X user ID of the monitored post's author, present for post events.
	AuthorID string `json:"authorId"`
	// Username of the monitored post's author, present for post events.
	AuthorUsername string `json:"authorUsername"`
	// Keyword monitor ID, present for keyword monitor events.
	KeywordMonitorID string `json:"keywordMonitorId"`
	// Your monitor label when Xquik created the event, present when set. A rename
	// leaves the event's label unchanged.
	MonitorName string `json:"monitorName"`
	// ID of the monitored post, present for post events.
	PostID string `json:"postId"`
	// Link to the monitored post, present for post events.
	PostURL string `json:"postUrl" format:"uri"`
	// Keyword query, present for keyword monitor events.
	Query string `json:"query"`
	// Account username, present for account monitor events.
	Username string `json:"username"`
	// Permanent X user ID of the monitored account, present for account monitor
	// events. It stays the same after a username change.
	XUserID string `json:"xUserId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		Data             respjson.Field
		MonitorID        respjson.Field
		MonitorType      respjson.Field
		OccurredAt       respjson.Field
		Type             respjson.Field
		AuthorID         respjson.Field
		AuthorUsername   respjson.Field
		KeywordMonitorID respjson.Field
		MonitorName      respjson.Field
		PostID           respjson.Field
		PostURL          respjson.Field
		Query            respjson.Field
		Username         respjson.Field
		XUserID          respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// RawJSON returns the unmodified JSON received from the API.
func (r Event) RawJSON() string { return r.JSON.raw }
func (r *Event) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// EventMonitorType is a source monitor type.
type EventMonitorType string

const (
	EventMonitorTypeAccount EventMonitorType = "account"
	EventMonitorTypeKeyword EventMonitorType = "keyword"
	EventMonitorTypePost    EventMonitorType = "post"
)

// EventDetail is a full monitor event including payload data and optional X event ID.
type EventDetail struct {
	ID string `json:"id" api:"required"`
	// Post or profile-change payload for the event type. Read its variant with
	// [MonitorEventDataUnion.AsAny], passing the event's Type.
	Data MonitorEventDataUnion `json:"data" api:"required"`
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
	// "profile.unavailable.changed", "tweet.deleted", "broadcast.started",
	// "space.started", "following.added", "following.removed", "follower.added",
	// "follower.removed", "profile.followers.changed", "profile.following.changed".
	Type shared.EventType `json:"type" api:"required"`
	// Keyword monitor ID included on detailed keyword events.
	KeywordMonitorID string `json:"keywordMonitorId"`
	// Your monitor label when Xquik created this event, present when set. A rename
	// leaves the event's label unchanged.
	MonitorName string `json:"monitorName"`
	// Keyword query for this detailed monitor event.
	Query string `json:"query"`
	// Account username for this detailed monitor event.
	Username string `json:"username"`
	XEventID string `json:"xEventId"`
	// Permanent X user ID of the monitored account on account events. It stays the
	// same after a username change.
	XUserID string `json:"xUserId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		Data             respjson.Field
		MonitorID        respjson.Field
		MonitorType      respjson.Field
		OccurredAt       respjson.Field
		Type             respjson.Field
		KeywordMonitorID respjson.Field
		MonitorName      respjson.Field
		Query            respjson.Field
		Username         respjson.Field
		XEventID         respjson.Field
		XUserID          respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// RawJSON returns the unmodified JSON received from the API.
func (r EventDetail) RawJSON() string { return r.JSON.raw }
func (r *EventDetail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// EventDetailMonitorType is a source monitor type for this detailed event.
type EventDetailMonitorType string

const (
	EventDetailMonitorTypeAccount EventDetailMonitorType = "account"
	EventDetailMonitorTypeKeyword EventDetailMonitorType = "keyword"
	EventDetailMonitorTypePost    EventDetailMonitorType = "post"
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

// RawJSON returns the unmodified JSON received from the API.
func (r EventListResponse) RawJSON() string { return r.JSON.raw }
func (r *EventListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type EventListParams struct {
	// Previous nextCursor. Offset pagination is not supported.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Keyword monitor ID.
	KeywordMonitorID param.Opt[string] `query:"keywordMonitorId,omitzero" json:"-"`
	// Maximum items per page: 1 to 100, default 50. Credits can reduce paid results.
	// The endpoint returns 402 insufficient_credits when none are affordable.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Account monitor ID.
	MonitorID param.Opt[string] `query:"monitorId,omitzero" json:"-"`
	// Post monitor ID. The post_monitor_id alias also works.
	PostMonitorID param.Opt[string] `query:"postMonitorId,omitzero" json:"-"`
	// Filter events by type
	//
	// Any of "tweet.new", "tweet.reply", "tweet.retweet", "tweet.quote",
	// "tweet.media", "tweet.link", "tweet.poll", "tweet.mention", "tweet.hashtag",
	// "tweet.longform", "profile.avatar.changed", "profile.banner.changed",
	// "profile.name.changed", "profile.username.changed", "profile.bio.changed",
	// "profile.location.changed", "profile.url.changed", "profile.verified.changed",
	// "profile.protected.changed", "profile.pinned_tweet.changed",
	// "profile.unavailable.changed", "tweet.deleted", "broadcast.started",
	// "space.started", "following.added", "following.removed", "follower.added",
	// "follower.removed", "profile.followers.changed", "profile.following.changed".
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

type EventStreamParams struct {
	// The last event ID the client got. The stream first sends every event after it.
	LastEventID param.Opt[string] `header:"Last-Event-ID,omitzero" json:"-"`
	// Set to post_metrics to also get each post monitor's saved numbers. Any
	// other value is refused.
	//
	// Any of "post_metrics".
	Include EventStreamParamsInclude `query:"include,omitzero" json:"-"`
	// Set to true to resume 30 seconds before Last-Event-ID. Events saved late
	// then still arrive. Events you have can arrive again, so skip IDs you've
	// seen. An ID the stream can't find starts 5 minutes back, not 30 days.
	LookBack param.Opt[bool] `query:"lookBack,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [EventStreamParams]'s query parameters as `url.Values`.
func (r EventStreamParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// EventStreamParamsInclude is the extra frame kind the stream sends.
type EventStreamParamsInclude string

const (
	EventStreamParamsIncludePostMetrics EventStreamParamsInclude = "post_metrics"
)
