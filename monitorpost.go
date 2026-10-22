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
)

// Post monitoring with replies, quotes, edits, deletion & 6 metrics
//
// MonitorPostService contains methods and other services that help with
// interacting with the x-twitter-scraper API.
//
// You should not instantiate this service directly, and instead use the
// [NewMonitorPostService] method instead.
type MonitorPostService struct {
	options []option.RequestOption
}

// NewMonitorPostService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewMonitorPostService(opts ...option.RequestOption) (r MonitorPostService) {
	r = MonitorPostService{}
	r.options = opts
	return
}

// Creates a monitor for 1 X post. It tracks replies, quotes, edits, deletion & 6
// metrics, for 30 credits an hour with events & webhook deliveries included. That
// hourly price is a monitor's only charge. Creation needs the credits of the first
// hour & reads the post once.
func (r *MonitorPostService) New(ctx context.Context, body MonitorPostNewParams, opts ...option.RequestOption) (res *PostMonitor, err error) {
	opts = slices.Concat(r.options, opts)
	path := "monitors/posts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Returns 1 post monitor with its post & the 6 metrics, each with its current
// value & gain since tracking began.
func (r *MonitorPostService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *PostMonitor, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("monitors/posts/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Updates a post monitor's label, event types or active state. Turning a paused
// monitor on reads the post first, & answers 409 while X hides it.
func (r *MonitorPostService) Update(ctx context.Context, id string, body MonitorPostUpdateParams, opts ...option.RequestOption) (res *PostMonitor, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("monitors/posts/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Returns post monitors with each post's 6 metrics, their current values & gains
// since tracking began.
func (r *MonitorPostService) List(ctx context.Context, query MonitorPostListParams, opts ...option.RequestOption) (res *MonitorPostListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "monitors/posts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Stops 1 post monitor, then removes its stored events & metric history in bounded
// batches. Poll statusUrl until the monitor returns 404.
func (r *MonitorPostService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *MonitorPostDeleteResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("monitors/posts/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Returns 1 bucket per hour with each metric's value at the hour's end & its gain
// over that hour. The last bucket is the current hour. Hours before tracking began
// carry the starting values with a gain of 0.
func (r *MonitorPostService) GetMetrics(ctx context.Context, id string, query MonitorPostGetMetricsParams, opts ...option.RequestOption) (res *PostMonitorMetrics, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("monitors/posts/%s/metrics", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// A metric's current value & its gain since tracking began. Both are null when
// X shows no count, as for views on old posts.
type PostMetricReading struct {
	Current int64 `json:"current" api:"required"`
	Gain    int64 `json:"gain" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Current     respjson.Field
		Gain        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PostMetricReading) RawJSON() string { return r.JSON.raw }
func (r *PostMetricReading) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The 6 metrics of a post.
type PostMetrics struct {
	// A metric's current value & its gain since tracking began. Both are null when
	// X shows no count, as for views on old posts.
	Views PostMetricReading `json:"views" api:"required"`
	// A metric's current value & its gain since tracking began. Both are null when
	// X shows no count, as for views on old posts.
	Likes PostMetricReading `json:"likes" api:"required"`
	// A metric's current value & its gain since tracking began. Both are null when
	// X shows no count, as for views on old posts.
	Retweets PostMetricReading `json:"retweets" api:"required"`
	// A metric's current value & its gain since tracking began. Both are null when
	// X shows no count, as for views on old posts.
	Replies PostMetricReading `json:"replies" api:"required"`
	// A metric's current value & its gain since tracking began. Both are null when
	// X shows no count, as for views on old posts.
	Quotes PostMetricReading `json:"quotes" api:"required"`
	// A metric's current value & its gain since tracking began. Both are null when
	// X shows no count, as for views on old posts.
	Bookmarks PostMetricReading `json:"bookmarks" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Views       respjson.Field
		Likes       respjson.Field
		Retweets    respjson.Field
		Replies     respjson.Field
		Quotes      respjson.Field
		Bookmarks   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PostMetrics) RawJSON() string { return r.JSON.raw }
func (r *PostMetrics) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Post monitor that tracks 1 X post's replies, quotes, edits & metrics.
type PostMonitor struct {
	ID string `json:"id" api:"required"`
	// Customer-set label of the post monitor, when set.
	Name string `json:"name"`
	// ID of the post you gave.
	PostID  string `json:"postId" api:"required"`
	PostURL string `json:"postUrl" api:"required"`
	// Permanent X user ID of the post's author.
	AuthorID       string `json:"authorId" api:"required"`
	AuthorUsername string `json:"authorUsername" api:"required"`
	AuthorName     string `json:"authorName" api:"required"`
	// Text of the newest version Xquik read.
	Text     string    `json:"text" api:"required"`
	PostedAt time.Time `json:"postedAt" api:"required" format:"date-time"`
	// Post event types to send. Without it, a monitor sends new replies, new
	// quotes, like & retweet changes, edits & deletion.
	EventTypes []string  `json:"eventTypes" api:"required"`
	IsActive   bool      `json:"isActive" api:"required"`
	CreatedAt  time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Next hourly credit charge time for this post monitor.
	NextBillingAt time.Time `json:"nextBillingAt" api:"required" format:"date-time"`
	// Why Xquik paused this post monitor. insufficient_credits means the balance
	// could not pay an hour, so add credits & resume. The other 3 mean X stopped
	// showing the post. Absent when the monitor is active or you paused it
	// yourself.
	PausedReason string `json:"pausedReason"`
	// When Xquik paused this post monitor.
	PausedAt time.Time `json:"pausedAt" format:"date-time"`
	// The 6 metrics of a post.
	Metrics PostMetrics `json:"metrics" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		Name           respjson.Field
		PostID         respjson.Field
		PostURL        respjson.Field
		AuthorID       respjson.Field
		AuthorUsername respjson.Field
		AuthorName     respjson.Field
		Text           respjson.Field
		PostedAt       respjson.Field
		EventTypes     respjson.Field
		IsActive       respjson.Field
		CreatedAt      respjson.Field
		NextBillingAt  respjson.Field
		PausedReason   respjson.Field
		PausedAt       respjson.Field
		Metrics        respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PostMonitor) RawJSON() string { return r.JSON.raw }
func (r *PostMonitor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A metric's value at the hour's end & its gain over that hour. Both are null
// when X shows no count.
type PostMetricPoint struct {
	Value int64 `json:"value" api:"required"`
	Gain  int64 `json:"gain" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Value       respjson.Field
		Gain        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PostMetricPoint) RawJSON() string { return r.JSON.raw }
func (r *PostMetricPoint) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PostMonitorMetricsBucketMetrics struct {
	// A metric's value at the hour's end & its gain over that hour. Both are null
	// when X shows no count.
	Views PostMetricPoint `json:"views" api:"required"`
	// A metric's value at the hour's end & its gain over that hour. Both are null
	// when X shows no count.
	Likes PostMetricPoint `json:"likes" api:"required"`
	// A metric's value at the hour's end & its gain over that hour. Both are null
	// when X shows no count.
	Retweets PostMetricPoint `json:"retweets" api:"required"`
	// A metric's value at the hour's end & its gain over that hour. Both are null
	// when X shows no count.
	Replies PostMetricPoint `json:"replies" api:"required"`
	// A metric's value at the hour's end & its gain over that hour. Both are null
	// when X shows no count.
	Quotes PostMetricPoint `json:"quotes" api:"required"`
	// A metric's value at the hour's end & its gain over that hour. Both are null
	// when X shows no count.
	Bookmarks PostMetricPoint `json:"bookmarks" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Views       respjson.Field
		Likes       respjson.Field
		Retweets    respjson.Field
		Replies     respjson.Field
		Quotes      respjson.Field
		Bookmarks   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PostMonitorMetricsBucketMetrics) RawJSON() string { return r.JSON.raw }
func (r *PostMonitorMetricsBucketMetrics) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PostMonitorMetricsBucket struct {
	// Start of the hour.
	At      time.Time                       `json:"at" api:"required" format:"date-time"`
	Metrics PostMonitorMetricsBucketMetrics `json:"metrics" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		At          respjson.Field
		Metrics     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PostMonitorMetricsBucket) RawJSON() string { return r.JSON.raw }
func (r *PostMonitorMetricsBucket) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Hourly metric buckets of 1 post monitor, oldest first.
type PostMonitorMetrics struct {
	Hours   int64                      `json:"hours" api:"required"`
	Buckets []PostMonitorMetricsBucket `json:"buckets" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Hours       respjson.Field
		Buckets     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PostMonitorMetrics) RawJSON() string { return r.JSON.raw }
func (r *PostMonitorMetrics) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MonitorPostListResponse struct {
	HasMore  bool          `json:"hasMore" api:"required"`
	Monitors []PostMonitor `json:"monitors" api:"required"`
	// Post monitors on this page.
	Total      int64  `json:"total" api:"required"`
	NextCursor string `json:"nextCursor"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		HasMore     respjson.Field
		Monitors    respjson.Field
		Total       respjson.Field
		NextCursor  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorPostListResponse) RawJSON() string { return r.JSON.raw }
func (r *MonitorPostListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MonitorPostDeleteResponse struct {
	DeletionStatus string `json:"deletionStatus" api:"required"`
	// Poll this monitor URL until it returns 404.
	StatusURL string `json:"statusUrl" api:"required"`
	Success   bool   `json:"success" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DeletionStatus respjson.Field
		StatusURL      respjson.Field
		Success        respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorPostDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *MonitorPostDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MonitorPostNewParams struct {
	// Post link or ID. x.com, twitter.com, mobile & /i/web/status/ links work with
	// any query string.
	Post string `json:"post" api:"required"`
	// Your own label for the post monitor. Its events carry it as monitorName.
	Name param.Opt[string] `json:"name,omitzero"`
	// Post event types to send. Without it, a monitor sends new replies, new quotes,
	// like & retweet changes, edits & deletion.
	EventTypes []string `json:"eventTypes,omitzero"`
	paramObj
}

func (r MonitorPostNewParams) MarshalJSON() (data []byte, err error) {
	type shadow MonitorPostNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MonitorPostNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MonitorPostUpdateParams struct {
	IsActive param.Opt[bool] `json:"isActive,omitzero"`
	// New label for the post monitor. Send null to remove it.
	Name param.Opt[string] `json:"name,omitzero"`
	// Post event types to send. Without it, a monitor sends new replies, new quotes,
	// like & retweet changes, edits & deletion.
	EventTypes []string `json:"eventTypes,omitzero"`
	paramObj
}

func (r MonitorPostUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow MonitorPostUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MonitorPostUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MonitorPostListParams struct {
	// Previous nextCursor. The after alias also works. Offset pagination is not
	// supported.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Maximum items per page, from 1 to 200. Default 200. Pass nextCursor as cursor
	// for the next page.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MonitorPostListParams]'s query parameters as `url.Values`.
func (r MonitorPostListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MonitorPostGetMetricsParams struct {
	// Hours to return, from 1 to 720.
	Hours param.Opt[int64] `query:"hours,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MonitorPostGetMetricsParams]'s query parameters as
// `url.Values`.
func (r MonitorPostGetMetricsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
