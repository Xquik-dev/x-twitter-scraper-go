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

// Push notification integrations (Telegram)
//
// IntegrationService contains methods and other services that help with
// interacting with the x-twitter-scraper API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewIntegrationService] method instead.
type IntegrationService struct {
	options []option.RequestOption
}

// NewIntegrationService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewIntegrationService(opts ...option.RequestOption) (r IntegrationService) {
	r = IntegrationService{}
	r.options = opts
	return
}

// Create integration
func (r *IntegrationService) New(ctx context.Context, body IntegrationNewParams, opts ...option.RequestOption) (res *IntegrationNewResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "integrations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Get integration details
func (r *IntegrationService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *IntegrationGetResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("integrations/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Update integration
func (r *IntegrationService) Update(ctx context.Context, id string, body IntegrationUpdateParams, opts ...option.RequestOption) (res *IntegrationUpdateResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("integrations/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// List integrations
func (r *IntegrationService) List(ctx context.Context, opts ...option.RequestOption) (res *IntegrationListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "integrations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Delete integration
func (r *IntegrationService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *IntegrationDeleteResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("integrations/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// List integration delivery history
func (r *IntegrationService) ListDeliveries(ctx context.Context, id string, query IntegrationListDeliveriesParams, opts ...option.RequestOption) (res *IntegrationListDeliveriesResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("integrations/%s/deliveries", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Send test delivery
func (r *IntegrationService) SendTest(ctx context.Context, id string, opts ...option.RequestOption) (res *IntegrationSendTestResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("integrations/%s/test", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Third-party integration (e.g. Telegram) subscribed to monitor events.
type IntegrationNewResponse struct {
	ID string `json:"id" api:"required"`
	// Integration config — shape varies by type (JSON)
	Config    map[string]any `json:"config" api:"required"`
	CreatedAt time.Time      `json:"createdAt" api:"required" format:"date-time"`
	// Array of event types to subscribe to.
	//
	// Any of "tweet.new", "tweet.reply", "tweet.retweet", "tweet.quote",
	// "follower.gained", "follower.lost".
	EventTypes []string `json:"eventTypes" api:"required"`
	IsActive   bool     `json:"isActive" api:"required"`
	Name       string   `json:"name" api:"required"`
	// Any of "telegram".
	Type IntegrationNewResponseType `json:"type" api:"required"`
	// Event filter rules (JSON)
	Filters          map[string]any `json:"filters"`
	MessageTemplate  string         `json:"messageTemplate"`
	ScopeAllMonitors bool           `json:"scopeAllMonitors"`
	SilentPush       bool           `json:"silentPush"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		Config           respjson.Field
		CreatedAt        respjson.Field
		EventTypes       respjson.Field
		IsActive         respjson.Field
		Name             respjson.Field
		Type             respjson.Field
		Filters          respjson.Field
		MessageTemplate  respjson.Field
		ScopeAllMonitors respjson.Field
		SilentPush       respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IntegrationNewResponse) RawJSON() string { return r.JSON.raw }
func (r *IntegrationNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IntegrationNewResponseType string

const (
	IntegrationNewResponseTypeTelegram IntegrationNewResponseType = "telegram"
)

// Third-party integration (e.g. Telegram) subscribed to monitor events.
type IntegrationGetResponse struct {
	ID string `json:"id" api:"required"`
	// Integration config — shape varies by type (JSON)
	Config    map[string]any `json:"config" api:"required"`
	CreatedAt time.Time      `json:"createdAt" api:"required" format:"date-time"`
	// Array of event types to subscribe to.
	//
	// Any of "tweet.new", "tweet.reply", "tweet.retweet", "tweet.quote",
	// "follower.gained", "follower.lost".
	EventTypes []string `json:"eventTypes" api:"required"`
	IsActive   bool     `json:"isActive" api:"required"`
	Name       string   `json:"name" api:"required"`
	// Any of "telegram".
	Type IntegrationGetResponseType `json:"type" api:"required"`
	// Event filter rules (JSON)
	Filters          map[string]any `json:"filters"`
	MessageTemplate  string         `json:"messageTemplate"`
	ScopeAllMonitors bool           `json:"scopeAllMonitors"`
	SilentPush       bool           `json:"silentPush"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		Config           respjson.Field
		CreatedAt        respjson.Field
		EventTypes       respjson.Field
		IsActive         respjson.Field
		Name             respjson.Field
		Type             respjson.Field
		Filters          respjson.Field
		MessageTemplate  respjson.Field
		ScopeAllMonitors respjson.Field
		SilentPush       respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IntegrationGetResponse) RawJSON() string { return r.JSON.raw }
func (r *IntegrationGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IntegrationGetResponseType string

const (
	IntegrationGetResponseTypeTelegram IntegrationGetResponseType = "telegram"
)

// Third-party integration (e.g. Telegram) subscribed to monitor events.
type IntegrationUpdateResponse struct {
	ID string `json:"id" api:"required"`
	// Integration config — shape varies by type (JSON)
	Config    map[string]any `json:"config" api:"required"`
	CreatedAt time.Time      `json:"createdAt" api:"required" format:"date-time"`
	// Array of event types to subscribe to.
	//
	// Any of "tweet.new", "tweet.reply", "tweet.retweet", "tweet.quote",
	// "follower.gained", "follower.lost".
	EventTypes []string `json:"eventTypes" api:"required"`
	IsActive   bool     `json:"isActive" api:"required"`
	Name       string   `json:"name" api:"required"`
	// Any of "telegram".
	Type IntegrationUpdateResponseType `json:"type" api:"required"`
	// Event filter rules (JSON)
	Filters          map[string]any `json:"filters"`
	MessageTemplate  string         `json:"messageTemplate"`
	ScopeAllMonitors bool           `json:"scopeAllMonitors"`
	SilentPush       bool           `json:"silentPush"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		Config           respjson.Field
		CreatedAt        respjson.Field
		EventTypes       respjson.Field
		IsActive         respjson.Field
		Name             respjson.Field
		Type             respjson.Field
		Filters          respjson.Field
		MessageTemplate  respjson.Field
		ScopeAllMonitors respjson.Field
		SilentPush       respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IntegrationUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *IntegrationUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IntegrationUpdateResponseType string

const (
	IntegrationUpdateResponseTypeTelegram IntegrationUpdateResponseType = "telegram"
)

type IntegrationListResponse struct {
	Integrations []IntegrationListResponseIntegration `json:"integrations" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Integrations respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IntegrationListResponse) RawJSON() string { return r.JSON.raw }
func (r *IntegrationListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Third-party integration (e.g. Telegram) subscribed to monitor events.
type IntegrationListResponseIntegration struct {
	ID string `json:"id" api:"required"`
	// Integration config — shape varies by type (JSON)
	Config    map[string]any `json:"config" api:"required"`
	CreatedAt time.Time      `json:"createdAt" api:"required" format:"date-time"`
	// Array of event types to subscribe to.
	//
	// Any of "tweet.new", "tweet.reply", "tweet.retweet", "tweet.quote",
	// "follower.gained", "follower.lost".
	EventTypes []string `json:"eventTypes" api:"required"`
	IsActive   bool     `json:"isActive" api:"required"`
	Name       string   `json:"name" api:"required"`
	// Any of "telegram".
	Type string `json:"type" api:"required"`
	// Event filter rules (JSON)
	Filters          map[string]any `json:"filters"`
	MessageTemplate  string         `json:"messageTemplate"`
	ScopeAllMonitors bool           `json:"scopeAllMonitors"`
	SilentPush       bool           `json:"silentPush"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		Config           respjson.Field
		CreatedAt        respjson.Field
		EventTypes       respjson.Field
		IsActive         respjson.Field
		Name             respjson.Field
		Type             respjson.Field
		Filters          respjson.Field
		MessageTemplate  respjson.Field
		ScopeAllMonitors respjson.Field
		SilentPush       respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IntegrationListResponseIntegration) RawJSON() string { return r.JSON.raw }
func (r *IntegrationListResponseIntegration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IntegrationDeleteResponse struct {
	Success bool `json:"success" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IntegrationDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *IntegrationDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IntegrationListDeliveriesResponse struct {
	Deliveries []IntegrationListDeliveriesResponseDelivery `json:"deliveries" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Deliveries  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IntegrationListDeliveriesResponse) RawJSON() string { return r.JSON.raw }
func (r *IntegrationListDeliveriesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Integration delivery attempt record with status and retry count.
type IntegrationListDeliveriesResponseDelivery struct {
	ID             string    `json:"id" api:"required"`
	Attempts       int64     `json:"attempts" api:"required"`
	CreatedAt      time.Time `json:"createdAt" api:"required" format:"date-time"`
	EventType      string    `json:"eventType" api:"required"`
	Status         string    `json:"status" api:"required"`
	DeliveredAt    time.Time `json:"deliveredAt" format:"date-time"`
	LastError      string    `json:"lastError"`
	LastStatusCode int64     `json:"lastStatusCode"`
	SourceID       string    `json:"sourceId"`
	SourceType     string    `json:"sourceType"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		Attempts       respjson.Field
		CreatedAt      respjson.Field
		EventType      respjson.Field
		Status         respjson.Field
		DeliveredAt    respjson.Field
		LastError      respjson.Field
		LastStatusCode respjson.Field
		SourceID       respjson.Field
		SourceType     respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IntegrationListDeliveriesResponseDelivery) RawJSON() string { return r.JSON.raw }
func (r *IntegrationListDeliveriesResponseDelivery) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IntegrationSendTestResponse struct {
	Success bool `json:"success" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IntegrationSendTestResponse) RawJSON() string { return r.JSON.raw }
func (r *IntegrationSendTestResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IntegrationNewParams struct {
	// Integration config (e.g. Telegram chatId)
	Config IntegrationNewParamsConfig `json:"config,omitzero" api:"required"`
	// Array of event types to subscribe to.
	//
	// Any of "tweet.new", "tweet.reply", "tweet.retweet", "tweet.quote",
	// "follower.gained", "follower.lost".
	EventTypes []string `json:"eventTypes,omitzero" api:"required"`
	Name       string   `json:"name" api:"required"`
	// Any of "telegram".
	Type IntegrationNewParamsType `json:"type,omitzero" api:"required"`
	paramObj
}

func (r IntegrationNewParams) MarshalJSON() (data []byte, err error) {
	type shadow IntegrationNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *IntegrationNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Integration config (e.g. Telegram chatId)
//
// The property ChatID is required.
type IntegrationNewParamsConfig struct {
	ChatID string `json:"chatId" api:"required"`
	paramObj
}

func (r IntegrationNewParamsConfig) MarshalJSON() (data []byte, err error) {
	type shadow IntegrationNewParamsConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *IntegrationNewParamsConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IntegrationNewParamsType string

const (
	IntegrationNewParamsTypeTelegram IntegrationNewParamsType = "telegram"
)

type IntegrationUpdateParams struct {
	IsActive         param.Opt[bool]   `json:"isActive,omitzero"`
	Name             param.Opt[string] `json:"name,omitzero"`
	ScopeAllMonitors param.Opt[bool]   `json:"scopeAllMonitors,omitzero"`
	SilentPush       param.Opt[bool]   `json:"silentPush,omitzero"`
	// Array of event types to subscribe to.
	//
	// Any of "tweet.new", "tweet.reply", "tweet.retweet", "tweet.quote",
	// "follower.gained", "follower.lost".
	EventTypes []string `json:"eventTypes,omitzero"`
	// Event filter rules (JSON)
	Filters map[string]any `json:"filters,omitzero"`
	// Custom message template (JSON)
	MessageTemplate map[string]any `json:"messageTemplate,omitzero"`
	paramObj
}

func (r IntegrationUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow IntegrationUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *IntegrationUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IntegrationListDeliveriesParams struct {
	// Maximum number of items to return (1-100, default 50)
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [IntegrationListDeliveriesParams]'s query parameters as
// `url.Values`.
func (r IntegrationListDeliveriesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
