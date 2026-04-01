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
	"github.com/stainless-sdks/x-twitter-scraper-go/internal/requestconfig"
	"github.com/stainless-sdks/x-twitter-scraper-go/option"
	"github.com/stainless-sdks/x-twitter-scraper-go/packages/param"
	"github.com/stainless-sdks/x-twitter-scraper-go/packages/respjson"
)

// Real-time X account monitoring
//
// MonitorService contains methods and other services that help with interacting
// with the x-twitter-scraper API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMonitorService] method instead.
type MonitorService struct {
	options []option.RequestOption
}

// NewMonitorService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewMonitorService(opts ...option.RequestOption) (r MonitorService) {
	r = MonitorService{}
	r.options = opts
	return
}

// Create monitor
func (r *MonitorService) New(ctx context.Context, body MonitorNewParams, opts ...option.RequestOption) (res *MonitorNewResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "monitors"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Get monitor
func (r *MonitorService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *MonitorGetResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("monitors/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Update monitor
func (r *MonitorService) Update(ctx context.Context, id string, body MonitorUpdateParams, opts ...option.RequestOption) (res *MonitorUpdateResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("monitors/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// List monitors
func (r *MonitorService) List(ctx context.Context, opts ...option.RequestOption) (res *MonitorListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "monitors"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Deactivate monitor
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

type MonitorNewResponse struct {
	ID        string    `json:"id" api:"required"`
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Any of "tweet.new", "tweet.reply", "tweet.retweet", "tweet.quote",
	// "follower.gained", "follower.lost".
	EventTypes []string `json:"eventTypes" api:"required"`
	Username   string   `json:"username" api:"required"`
	XUserID    string   `json:"xUserId" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		EventTypes  respjson.Field
		Username    respjson.Field
		XUserID     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorNewResponse) RawJSON() string { return r.JSON.raw }
func (r *MonitorNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MonitorGetResponse struct {
	ID        string    `json:"id" api:"required"`
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Any of "tweet.new", "tweet.reply", "tweet.retweet", "tweet.quote",
	// "follower.gained", "follower.lost".
	EventTypes []string `json:"eventTypes" api:"required"`
	IsActive   bool     `json:"isActive" api:"required"`
	Username   string   `json:"username" api:"required"`
	XUserID    string   `json:"xUserId" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		EventTypes  respjson.Field
		IsActive    respjson.Field
		Username    respjson.Field
		XUserID     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorGetResponse) RawJSON() string { return r.JSON.raw }
func (r *MonitorGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MonitorUpdateResponse struct {
	ID        string    `json:"id" api:"required"`
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Any of "tweet.new", "tweet.reply", "tweet.retweet", "tweet.quote",
	// "follower.gained", "follower.lost".
	EventTypes []string `json:"eventTypes" api:"required"`
	IsActive   bool     `json:"isActive" api:"required"`
	Username   string   `json:"username" api:"required"`
	XUserID    string   `json:"xUserId" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		EventTypes  respjson.Field
		IsActive    respjson.Field
		Username    respjson.Field
		XUserID     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *MonitorUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MonitorListResponse struct {
	Monitors []MonitorListResponseMonitor `json:"monitors" api:"required"`
	Total    int64                        `json:"total" api:"required"`
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

type MonitorListResponseMonitor struct {
	ID        string    `json:"id" api:"required"`
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Any of "tweet.new", "tweet.reply", "tweet.retweet", "tweet.quote",
	// "follower.gained", "follower.lost".
	EventTypes []string `json:"eventTypes" api:"required"`
	IsActive   bool     `json:"isActive" api:"required"`
	Username   string   `json:"username" api:"required"`
	XUserID    string   `json:"xUserId" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		EventTypes  respjson.Field
		IsActive    respjson.Field
		Username    respjson.Field
		XUserID     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorListResponseMonitor) RawJSON() string { return r.JSON.raw }
func (r *MonitorListResponseMonitor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MonitorDeactivateResponse struct {
	Success bool `json:"success" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorDeactivateResponse) RawJSON() string { return r.JSON.raw }
func (r *MonitorDeactivateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MonitorNewParams struct {
	// Any of "tweet.new", "tweet.reply", "tweet.retweet", "tweet.quote",
	// "follower.gained", "follower.lost".
	EventTypes []string `json:"eventTypes,omitzero" api:"required"`
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
	// Any of "tweet.new", "tweet.reply", "tweet.retweet", "tweet.quote",
	// "follower.gained", "follower.lost".
	EventTypes []string `json:"eventTypes,omitzero"`
	paramObj
}

func (r MonitorUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow MonitorUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MonitorUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
