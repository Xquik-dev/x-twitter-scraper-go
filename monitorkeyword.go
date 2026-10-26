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
	"github.com/Xquik-dev/x-twitter-scraper-go/internal/requestconfig"
	"github.com/Xquik-dev/x-twitter-scraper-go/option"
	"github.com/Xquik-dev/x-twitter-scraper-go/packages/param"
	"github.com/Xquik-dev/x-twitter-scraper-go/packages/respjson"
	"github.com/Xquik-dev/x-twitter-scraper-go/shared"
	"github.com/Xquik-dev/x-twitter-scraper-go/shared/constant"
)

// MonitorKeywordService covers X account monitoring with 1-second checks.
//
// MonitorKeywordService contains methods and other services that help with
// interacting with the x-twitter-scraper API.
//
// You should not instantiate this service directly, and instead use the
// [NewMonitorKeywordService] method instead.
type MonitorKeywordService struct {
	options []option.RequestOption
}

// NewMonitorKeywordService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewMonitorKeywordService(opts ...option.RequestOption) (r MonitorKeywordService) {
	r = MonitorKeywordService{}
	r.options = opts
	return
}

// New creates a keyword monitor. Keyword monitors are unlimited. Active monitors check
// every 1 second and cost 30 credits per hour. Events and webhook deliveries are
// included. Creation requires available credits for the first hourly charge.
func (r *MonitorKeywordService) New(ctx context.Context, body MonitorKeywordNewParams, opts ...option.RequestOption) (res *MonitorKeywordNewResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "monitors/keywords"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Get returns configuration and status for one keyword monitor.
func (r *MonitorKeywordService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *MonitorKeywordGetResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("monitors/keywords/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Update updates mutable settings for an existing keyword monitor.
func (r *MonitorKeywordService) Update(ctx context.Context, id string, body MonitorKeywordUpdateParams, opts ...option.RequestOption) (res *MonitorKeywordUpdateResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("monitors/keywords/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// List returns keyword monitors with their current operating states.
func (r *MonitorKeywordService) List(ctx context.Context, opts ...option.RequestOption) (res *MonitorKeywordListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "monitors/keywords"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Deactivate stops one keyword monitor, then removes its stored events in bounded batches.
// Poll statusUrl until the monitor returns 404.
func (r *MonitorKeywordService) Deactivate(ctx context.Context, id string, opts ...option.RequestOption) (res *MonitorKeywordDeactivateResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("monitors/keywords/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// MonitorKeywordNewResponse is a keyword monitor that tracks matching public X activity.
type MonitorKeywordNewResponse struct {
	ID        string    `json:"id" api:"required"`
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Array of event types to subscribe to.
	EventTypes []shared.EventType `json:"eventTypes" api:"required"`
	IsActive   bool               `json:"isActive" api:"required"`
	// Next hourly credit charge time for this keyword query monitor.
	NextBillingAt time.Time `json:"nextBillingAt" api:"required" format:"date-time"`
	Query         string    `json:"query" api:"required"`
	// Customer-set label of the keyword monitor, when set.
	Name string `json:"name"`
	// When Xquik paused this keyword monitor.
	PausedAt time.Time `json:"pausedAt" format:"date-time"`
	// Why Xquik paused this keyword monitor. The balance could not pay an hour, so add
	// credits & resume. Absent when the monitor is active or you paused it yourself.
	PausedReason constant.InsufficientCredits `json:"pausedReason" default:"insufficient_credits"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		CreatedAt     respjson.Field
		EventTypes    respjson.Field
		IsActive      respjson.Field
		NextBillingAt respjson.Field
		Query         respjson.Field
		Name          respjson.Field
		PausedAt      respjson.Field
		PausedReason  respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// RawJSON returns the unmodified JSON received from the API.
func (r MonitorKeywordNewResponse) RawJSON() string { return r.JSON.raw }
func (r *MonitorKeywordNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// MonitorKeywordGetResponse is a keyword monitor that tracks matching public X activity.
type MonitorKeywordGetResponse struct {
	ID        string    `json:"id" api:"required"`
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Array of event types to subscribe to.
	EventTypes []shared.EventType `json:"eventTypes" api:"required"`
	IsActive   bool               `json:"isActive" api:"required"`
	// Next hourly credit charge time for this keyword query monitor.
	NextBillingAt time.Time `json:"nextBillingAt" api:"required" format:"date-time"`
	Query         string    `json:"query" api:"required"`
	// Customer-set label of the keyword monitor, when set.
	Name string `json:"name"`
	// When Xquik paused this keyword monitor.
	PausedAt time.Time `json:"pausedAt" format:"date-time"`
	// Why Xquik paused this keyword monitor. The balance could not pay an hour, so add
	// credits & resume. Absent when the monitor is active or you paused it yourself.
	PausedReason constant.InsufficientCredits `json:"pausedReason" default:"insufficient_credits"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		CreatedAt     respjson.Field
		EventTypes    respjson.Field
		IsActive      respjson.Field
		NextBillingAt respjson.Field
		Query         respjson.Field
		Name          respjson.Field
		PausedAt      respjson.Field
		PausedReason  respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// RawJSON returns the unmodified JSON received from the API.
func (r MonitorKeywordGetResponse) RawJSON() string { return r.JSON.raw }
func (r *MonitorKeywordGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// MonitorKeywordUpdateResponse is a keyword monitor that tracks matching public X activity.
type MonitorKeywordUpdateResponse struct {
	ID        string    `json:"id" api:"required"`
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Array of event types to subscribe to.
	EventTypes []shared.EventType `json:"eventTypes" api:"required"`
	IsActive   bool               `json:"isActive" api:"required"`
	// Next hourly credit charge time for this keyword query monitor.
	NextBillingAt time.Time `json:"nextBillingAt" api:"required" format:"date-time"`
	Query         string    `json:"query" api:"required"`
	// Customer-set label of the keyword monitor, when set.
	Name string `json:"name"`
	// When Xquik paused this keyword monitor.
	PausedAt time.Time `json:"pausedAt" format:"date-time"`
	// Why Xquik paused this keyword monitor. The balance could not pay an hour, so add
	// credits & resume. Absent when the monitor is active or you paused it yourself.
	PausedReason constant.InsufficientCredits `json:"pausedReason" default:"insufficient_credits"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		CreatedAt     respjson.Field
		EventTypes    respjson.Field
		IsActive      respjson.Field
		NextBillingAt respjson.Field
		Query         respjson.Field
		Name          respjson.Field
		PausedAt      respjson.Field
		PausedReason  respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// RawJSON returns the unmodified JSON received from the API.
func (r MonitorKeywordUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *MonitorKeywordUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type MonitorKeywordListResponse struct {
	HasMore    bool                                `json:"hasMore" api:"required"`
	Monitors   []MonitorKeywordListResponseMonitor `json:"monitors" api:"required"`
	Total      int64                               `json:"total" api:"required"`
	NextCursor string                              `json:"nextCursor"`
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

// RawJSON returns the unmodified JSON received from the API.
func (r MonitorKeywordListResponse) RawJSON() string { return r.JSON.raw }
func (r *MonitorKeywordListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// MonitorKeywordListResponseMonitor is a keyword monitor that tracks matching public X activity.
type MonitorKeywordListResponseMonitor struct {
	ID        string    `json:"id" api:"required"`
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Array of event types to subscribe to.
	EventTypes []shared.EventType `json:"eventTypes" api:"required"`
	IsActive   bool               `json:"isActive" api:"required"`
	// Next hourly credit charge time for this keyword query monitor.
	NextBillingAt time.Time `json:"nextBillingAt" api:"required" format:"date-time"`
	Query         string    `json:"query" api:"required"`
	// Customer-set label of the keyword monitor, when set.
	Name string `json:"name"`
	// When Xquik paused this keyword monitor.
	PausedAt time.Time `json:"pausedAt" format:"date-time"`
	// Why Xquik paused this keyword monitor. The balance could not pay an hour, so add
	// credits & resume. Absent when the monitor is active or you paused it yourself.
	PausedReason constant.InsufficientCredits `json:"pausedReason" default:"insufficient_credits"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		CreatedAt     respjson.Field
		EventTypes    respjson.Field
		IsActive      respjson.Field
		NextBillingAt respjson.Field
		Query         respjson.Field
		Name          respjson.Field
		PausedAt      respjson.Field
		PausedReason  respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// RawJSON returns the unmodified JSON received from the API.
func (r MonitorKeywordListResponseMonitor) RawJSON() string { return r.JSON.raw }
func (r *MonitorKeywordListResponseMonitor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type MonitorKeywordDeactivateResponse struct {
	DeletionStatus constant.Deleting `json:"deletionStatus" default:"deleting"`
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

// RawJSON returns the unmodified JSON received from the API.
func (r MonitorKeywordDeactivateResponse) RawJSON() string { return r.JSON.raw }
func (r *MonitorKeywordDeactivateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type MonitorKeywordNewParams struct {
	// Array of event types to subscribe to.
	EventTypes []shared.EventType `json:"eventTypes,omitzero" api:"required"`
	// X search query to monitor. Whitespace is normalized.
	Query string `json:"query" api:"required"`
	// Your own label for the keyword monitor. Its events carry it as monitorName.
	Name param.Opt[string] `json:"name,omitzero"`
	paramObj
}

func (r MonitorKeywordNewParams) MarshalJSON() (data []byte, err error) {
	type shadow MonitorKeywordNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MonitorKeywordNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, nil)
}

type MonitorKeywordUpdateParams struct {
	IsActive param.Opt[bool] `json:"isActive,omitzero"`
	// New label for the keyword monitor. Send null to remove it.
	Name param.Opt[string] `json:"name,omitzero"`
	// Array of event types to subscribe to.
	EventTypes []shared.EventType `json:"eventTypes,omitzero"`
	paramObj
}

func (r MonitorKeywordUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow MonitorKeywordUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MonitorKeywordUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, nil)
}
