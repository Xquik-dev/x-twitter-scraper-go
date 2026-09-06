// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

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
	"github.com/Xquik-dev/x-twitter-scraper-go/internal/requestconfig"
	"github.com/Xquik-dev/x-twitter-scraper-go/option"
	"github.com/Xquik-dev/x-twitter-scraper-go/packages/param"
	"github.com/Xquik-dev/x-twitter-scraper-go/packages/respjson"
	"github.com/Xquik-dev/x-twitter-scraper-go/shared"
	"github.com/Xquik-dev/x-twitter-scraper-go/shared/constant"
)

// X account monitoring with 1-second checks
//
// MonitorService contains methods and other services that help with interacting
// with the x-twitter-scraper API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMonitorService] method instead.
type MonitorService struct {
	options []option.RequestOption
	// X account monitoring with 1-second checks
	Keywords MonitorKeywordService
}

// NewMonitorService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewMonitorService(opts ...option.RequestOption) (r MonitorService) {
	r = MonitorService{}
	r.options = opts
	r.Keywords = NewMonitorKeywordService(opts...)
	return
}

// Creates an account monitor. Monitors are unlimited. Active monitors check every
// 1 second and cost 21 credits per hour. Events and webhook deliveries are
// included. Creation requires available credits for the first hourly charge and
// username lookup.
func (r *MonitorService) New(ctx context.Context, body MonitorNewParams, opts ...option.RequestOption) (res *MonitorNewResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "monitors"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Returns configuration and status for one account monitor.
func (r *MonitorService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *Monitor, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("monitors/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Updates mutable settings for an existing account monitor.
func (r *MonitorService) Update(ctx context.Context, id string, body MonitorUpdateParams, opts ...option.RequestOption) (res *Monitor, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("monitors/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Returns account monitors with their current operating states.
func (r *MonitorService) List(ctx context.Context, opts ...option.RequestOption) (res *MonitorListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "monitors"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Stops one account monitor, then removes its stored events in bounded batches.
// Poll statusUrl until the monitor returns 404.
func (r *MonitorService) Deactivate(ctx context.Context, id string, opts ...option.RequestOption) (res *MonitorDeactivateResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("monitors/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Account monitor that tracks activity for a given X user.
type Monitor struct {
	ID        string    `json:"id" api:"required"`
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Array of event types to subscribe to.
	EventTypes []shared.EventType `json:"eventTypes" api:"required"`
	IsActive   bool               `json:"isActive" api:"required"`
	// Next hourly credit charge time for this account monitor.
	NextBillingAt time.Time `json:"nextBillingAt" api:"required" format:"date-time"`
	Username      string    `json:"username" api:"required"`
	XUserID       string    `json:"xUserId" api:"required"`
	// When Xquik automatically paused this monitor.
	PausedAt time.Time `json:"pausedAt" format:"date-time"`
	// Why Xquik automatically paused this monitor.
	//
	// Any of "x_user_not_found".
	PausedReason MonitorPausedReason `json:"pausedReason"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		CreatedAt     respjson.Field
		EventTypes    respjson.Field
		IsActive      respjson.Field
		NextBillingAt respjson.Field
		Username      respjson.Field
		XUserID       respjson.Field
		PausedAt      respjson.Field
		PausedReason  respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Monitor) RawJSON() string { return r.JSON.raw }
func (r *Monitor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Why Xquik automatically paused this monitor.
type MonitorPausedReason string

const (
	MonitorPausedReasonXUserNotFound MonitorPausedReason = "x_user_not_found"
)

type MonitorNewResponse struct {
	ID        string    `json:"id" api:"required"`
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Array of event types to subscribe to.
	EventTypes []shared.EventType `json:"eventTypes" api:"required"`
	IsActive   bool               `json:"isActive" api:"required"`
	// Next hourly credit charge time. New active monitors are due immediately.
	NextBillingAt time.Time `json:"nextBillingAt" api:"required" format:"date-time"`
	Username      string    `json:"username" api:"required"`
	XUserID       string    `json:"xUserId" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		CreatedAt     respjson.Field
		EventTypes    respjson.Field
		IsActive      respjson.Field
		NextBillingAt respjson.Field
		Username      respjson.Field
		XUserID       respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorNewResponse) RawJSON() string { return r.JSON.raw }
func (r *MonitorNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MonitorListResponse struct {
	Monitors []Monitor `json:"monitors" api:"required"`
	Total    int64     `json:"total" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Monitors    respjson.Field
		Total       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorListResponse) RawJSON() string { return r.JSON.raw }
func (r *MonitorListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MonitorDeactivateResponse struct {
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

// Returns the unmodified JSON received from the API
func (r MonitorDeactivateResponse) RawJSON() string { return r.JSON.raw }
func (r *MonitorDeactivateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MonitorNewParams struct {
	// Array of event types to subscribe to.
	EventTypes []shared.EventType `json:"eventTypes,omitzero" api:"required"`
	// X username (without @)
	Username string `json:"username" api:"required"`
	paramObj
}

func (r MonitorNewParams) MarshalJSON() (data []byte, err error) {
	type shadow MonitorNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MonitorNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MonitorUpdateParams struct {
	IsActive param.Opt[bool] `json:"isActive,omitzero"`
	// Array of event types to subscribe to.
	EventTypes []shared.EventType `json:"eventTypes,omitzero"`
	paramObj
}

func (r MonitorUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow MonitorUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MonitorUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
