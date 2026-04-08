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
)

// Webhook endpoint management & delivery
//
// WebhookService contains methods and other services that help with interacting
// with the x-twitter-scraper API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWebhookService] method instead.
type WebhookService struct {
	options []option.RequestOption
}

// NewWebhookService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewWebhookService(opts ...option.RequestOption) (r WebhookService) {
	r = WebhookService{}
	r.options = opts
	return
}

// Create webhook
func (r *WebhookService) New(ctx context.Context, body WebhookNewParams, opts ...option.RequestOption) (res *WebhookNewResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "webhooks"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Update webhook
func (r *WebhookService) Update(ctx context.Context, id string, body WebhookUpdateParams, opts ...option.RequestOption) (res *WebhookUpdateResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("webhooks/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// List webhooks
func (r *WebhookService) List(ctx context.Context, opts ...option.RequestOption) (res *WebhookListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "webhooks"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Deactivate webhook
func (r *WebhookService) Deactivate(ctx context.Context, id string, opts ...option.RequestOption) (res *WebhookDeactivateResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("webhooks/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// List webhook deliveries
func (r *WebhookService) ListDeliveries(ctx context.Context, id string, opts ...option.RequestOption) (res *WebhookListDeliveriesResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("webhooks/%s/deliveries", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Test webhook endpoint
func (r *WebhookService) Test(ctx context.Context, id string, opts ...option.RequestOption) (res *WebhookTestResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("webhooks/%s/test", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

type WebhookNewResponse struct {
	ID        string    `json:"id" api:"required"`
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Array of event types to subscribe to.
	//
	// Any of "tweet.new", "tweet.reply", "tweet.retweet", "tweet.quote",
	// "follower.gained", "follower.lost".
	EventTypes []string `json:"eventTypes" api:"required"`
	Secret     string   `json:"secret" api:"required"`
	URL        string   `json:"url" api:"required" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		EventTypes  respjson.Field
		Secret      respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookNewResponse) RawJSON() string { return r.JSON.raw }
func (r *WebhookNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Webhook endpoint registered to receive event deliveries.
type WebhookUpdateResponse struct {
	ID        string    `json:"id" api:"required"`
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Array of event types to subscribe to.
	//
	// Any of "tweet.new", "tweet.reply", "tweet.retweet", "tweet.quote",
	// "follower.gained", "follower.lost".
	EventTypes []string `json:"eventTypes" api:"required"`
	IsActive   bool     `json:"isActive" api:"required"`
	URL        string   `json:"url" api:"required" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		EventTypes  respjson.Field
		IsActive    respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *WebhookUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookListResponse struct {
	Webhooks []WebhookListResponseWebhook `json:"webhooks" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Webhooks    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookListResponse) RawJSON() string { return r.JSON.raw }
func (r *WebhookListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Webhook endpoint registered to receive event deliveries.
type WebhookListResponseWebhook struct {
	ID        string    `json:"id" api:"required"`
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Array of event types to subscribe to.
	//
	// Any of "tweet.new", "tweet.reply", "tweet.retweet", "tweet.quote",
	// "follower.gained", "follower.lost".
	EventTypes []string `json:"eventTypes" api:"required"`
	IsActive   bool     `json:"isActive" api:"required"`
	URL        string   `json:"url" api:"required" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		EventTypes  respjson.Field
		IsActive    respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookListResponseWebhook) RawJSON() string { return r.JSON.raw }
func (r *WebhookListResponseWebhook) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookDeactivateResponse struct {
	Success bool `json:"success" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookDeactivateResponse) RawJSON() string { return r.JSON.raw }
func (r *WebhookDeactivateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookListDeliveriesResponse struct {
	Deliveries []WebhookListDeliveriesResponseDelivery `json:"deliveries" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Deliveries  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookListDeliveriesResponse) RawJSON() string { return r.JSON.raw }
func (r *WebhookListDeliveriesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Webhook delivery attempt record with status and retry count.
type WebhookListDeliveriesResponseDelivery struct {
	ID             string    `json:"id" api:"required"`
	Attempts       int64     `json:"attempts" api:"required"`
	CreatedAt      time.Time `json:"createdAt" api:"required" format:"date-time"`
	Status         string    `json:"status" api:"required"`
	StreamEventID  string    `json:"streamEventId" api:"required"`
	DeliveredAt    time.Time `json:"deliveredAt" format:"date-time"`
	LastError      string    `json:"lastError"`
	LastStatusCode int64     `json:"lastStatusCode"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		Attempts       respjson.Field
		CreatedAt      respjson.Field
		Status         respjson.Field
		StreamEventID  respjson.Field
		DeliveredAt    respjson.Field
		LastError      respjson.Field
		LastStatusCode respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookListDeliveriesResponseDelivery) RawJSON() string { return r.JSON.raw }
func (r *WebhookListDeliveriesResponseDelivery) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookTestResponse struct {
	StatusCode int64  `json:"statusCode" api:"required"`
	Success    bool   `json:"success" api:"required"`
	Error      string `json:"error"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		StatusCode  respjson.Field
		Success     respjson.Field
		Error       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookTestResponse) RawJSON() string { return r.JSON.raw }
func (r *WebhookTestResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookNewParams struct {
	// Array of event types to subscribe to.
	//
	// Any of "tweet.new", "tweet.reply", "tweet.retweet", "tweet.quote",
	// "follower.gained", "follower.lost".
	EventTypes []string `json:"eventTypes,omitzero" api:"required"`
	// HTTPS URL
	URL string `json:"url" api:"required" format:"uri"`
	paramObj
}

func (r WebhookNewParams) MarshalJSON() (data []byte, err error) {
	type shadow WebhookNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WebhookNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookUpdateParams struct {
	IsActive param.Opt[bool]   `json:"isActive,omitzero"`
	URL      param.Opt[string] `json:"url,omitzero" format:"uri"`
	// Array of event types to subscribe to.
	//
	// Any of "tweet.new", "tweet.reply", "tweet.retweet", "tweet.quote",
	// "follower.gained", "follower.lost".
	EventTypes []string `json:"eventTypes,omitzero"`
	paramObj
}

func (r WebhookUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow WebhookUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WebhookUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
