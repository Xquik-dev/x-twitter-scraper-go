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
)

// Real-time X account monitoring
//
// MonitorKeywordService contains methods and other services that help with
// interacting with the x-twitter-scraper API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMonitorKeywordService] method instead.
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

// Creates an instant keyword monitor. Keyword monitors are unlimited. Active
// monitors check every 1 second and cost 21 credits per hour. Events and webhook
// deliveries are included. Creation requires available credits for the first
// hourly charge.
func (r *MonitorKeywordService) New(ctx context.Context, body MonitorKeywordNewParams, opts ...option.RequestOption) (res *MonitorKeywordNewResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "monitors/keywords"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Get keyword monitor
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

// Update keyword monitor
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

// List keyword monitors
func (r *MonitorKeywordService) List(ctx context.Context, opts ...option.RequestOption) (res *MonitorKeywordListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "monitors/keywords"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Delete keyword monitor
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

// Keyword monitor that tracks matching public X activity.
type MonitorKeywordNewResponse struct {
	ID        string    `json:"id" api:"required"`
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Array of event types to subscribe to.
	EventTypes []shared.EventType `json:"eventTypes" api:"required"`
	IsActive   bool               `json:"isActive" api:"required"`
	// Next hourly credit charge time for this keyword query monitor.
	NextBillingAt time.Time `json:"nextBillingAt" api:"required" format:"date-time"`
	Query         string    `json:"query" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		CreatedAt     respjson.Field
		EventTypes    respjson.Field
		IsActive      respjson.Field
		NextBillingAt respjson.Field
		Query         respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorKeywordNewResponse) RawJSON() string { return r.JSON.raw }
func (r *MonitorKeywordNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Keyword monitor that tracks matching public X activity.
type MonitorKeywordGetResponse struct {
	ID        string    `json:"id" api:"required"`
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Array of event types to subscribe to.
	EventTypes []shared.EventType `json:"eventTypes" api:"required"`
	IsActive   bool               `json:"isActive" api:"required"`
	// Next hourly credit charge time for this keyword query monitor.
	NextBillingAt time.Time `json:"nextBillingAt" api:"required" format:"date-time"`
	Query         string    `json:"query" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		CreatedAt     respjson.Field
		EventTypes    respjson.Field
		IsActive      respjson.Field
		NextBillingAt respjson.Field
		Query         respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorKeywordGetResponse) RawJSON() string { return r.JSON.raw }
func (r *MonitorKeywordGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Keyword monitor that tracks matching public X activity.
type MonitorKeywordUpdateResponse struct {
	ID        string    `json:"id" api:"required"`
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Array of event types to subscribe to.
	EventTypes []shared.EventType `json:"eventTypes" api:"required"`
	IsActive   bool               `json:"isActive" api:"required"`
	// Next hourly credit charge time for this keyword query monitor.
	NextBillingAt time.Time `json:"nextBillingAt" api:"required" format:"date-time"`
	Query         string    `json:"query" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		CreatedAt     respjson.Field
		EventTypes    respjson.Field
		IsActive      respjson.Field
		NextBillingAt respjson.Field
		Query         respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorKeywordUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *MonitorKeywordUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MonitorKeywordListResponse struct {
	Monitors []MonitorKeywordListResponseMonitor `json:"monitors" api:"required"`
	Total    int64                               `json:"total" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Monitors    respjson.Field
		Total       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorKeywordListResponse) RawJSON() string { return r.JSON.raw }
func (r *MonitorKeywordListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Keyword monitor that tracks matching public X activity.
type MonitorKeywordListResponseMonitor struct {
	ID        string    `json:"id" api:"required"`
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Array of event types to subscribe to.
	EventTypes []shared.EventType `json:"eventTypes" api:"required"`
	IsActive   bool               `json:"isActive" api:"required"`
	// Next hourly credit charge time for this keyword query monitor.
	NextBillingAt time.Time `json:"nextBillingAt" api:"required" format:"date-time"`
	Query         string    `json:"query" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		CreatedAt     respjson.Field
		EventTypes    respjson.Field
		IsActive      respjson.Field
		NextBillingAt respjson.Field
		Query         respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorKeywordListResponseMonitor) RawJSON() string { return r.JSON.raw }
func (r *MonitorKeywordListResponseMonitor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MonitorKeywordDeactivateResponse struct {
	Success bool `json:"success" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorKeywordDeactivateResponse) RawJSON() string { return r.JSON.raw }
func (r *MonitorKeywordDeactivateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MonitorKeywordNewParams struct {
	// Array of event types to subscribe to.
	EventTypes []shared.EventType `json:"eventTypes,omitzero" api:"required"`
	// X search query to monitor. Whitespace is normalized.
	Query string `json:"query" api:"required"`
	paramObj
}

func (r MonitorKeywordNewParams) MarshalJSON() (data []byte, err error) {
	type shadow MonitorKeywordNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MonitorKeywordNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MonitorKeywordUpdateParams struct {
	IsActive param.Opt[bool] `json:"isActive,omitzero"`
	// Array of event types to subscribe to.
	EventTypes []shared.EventType `json:"eventTypes,omitzero"`
	paramObj
}

func (r MonitorKeywordUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow MonitorKeywordUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MonitorKeywordUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
