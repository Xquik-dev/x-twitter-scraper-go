// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package xtwitterscraper

import (
	"context"
	"net/http"
	"slices"

	"github.com/stainless-sdks/x-twitter-scraper-go/internal/apijson"
	"github.com/stainless-sdks/x-twitter-scraper-go/internal/requestconfig"
	"github.com/stainless-sdks/x-twitter-scraper-go/option"
	"github.com/stainless-sdks/x-twitter-scraper-go/packages/respjson"
)

// Subscription & billing
//
// SubscribeService contains methods and other services that help with interacting
// with the x-twitter-scraper API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSubscribeService] method instead.
type SubscribeService struct {
	options []option.RequestOption
}

// NewSubscribeService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSubscribeService(opts ...option.RequestOption) (r SubscribeService) {
	r = SubscribeService{}
	r.options = opts
	return
}

// Get checkout or billing URL
func (r *SubscribeService) New(ctx context.Context, opts ...option.RequestOption) (res *SubscribeNewResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "subscribe"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

type SubscribeNewResponse struct {
	URL     string `json:"url" api:"required" format:"uri"`
	Message string `json:"message"`
	// Any of "checkout_created", "already_subscribed", "payment_issue".
	Status SubscribeNewResponseStatus `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		URL         respjson.Field
		Message     respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SubscribeNewResponse) RawJSON() string { return r.JSON.raw }
func (r *SubscribeNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SubscribeNewResponseStatus string

const (
	SubscribeNewResponseStatusCheckoutCreated   SubscribeNewResponseStatus = "checkout_created"
	SubscribeNewResponseStatusAlreadySubscribed SubscribeNewResponseStatus = "already_subscribed"
	SubscribeNewResponseStatusPaymentIssue      SubscribeNewResponseStatus = "payment_issue"
)
