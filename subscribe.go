// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package xtwitterscraper

import (
	"context"
	"net/http"
	"slices"

	"github.com/Xquik-dev/x-twitter-scraper-go/internal/apijson"
	"github.com/Xquik-dev/x-twitter-scraper-go/internal/requestconfig"
	"github.com/Xquik-dev/x-twitter-scraper-go/option"
	"github.com/Xquik-dev/x-twitter-scraper-go/packages/param"
	"github.com/Xquik-dev/x-twitter-scraper-go/packages/respjson"
)

// Subscription, billing, and credits
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

// Create a subscription checkout or billing-management URL only after the user
// confirms. The request never completes payment by itself.
func (r *SubscribeService) New(ctx context.Context, body SubscribeNewParams, opts ...option.RequestOption) (res *SubscribeNewResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "subscribe"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type SubscribeNewResponse struct {
	Message string `json:"message" api:"required"`
	// Any of "checkout_created", "already_subscribed", "payment_issue".
	Status SubscribeNewResponseStatus `json:"status" api:"required"`
	URL    string                     `json:"url" api:"required" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		Status      respjson.Field
		URL         respjson.Field
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

type SubscribeNewParams struct {
	// Subscription tier to pre-select.
	//
	// Any of "starter", "pro", "business".
	Tier SubscribeNewParamsTier `json:"tier,omitzero"`
	paramObj
}

func (r SubscribeNewParams) MarshalJSON() (data []byte, err error) {
	type shadow SubscribeNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SubscribeNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Subscription tier to pre-select.
type SubscribeNewParamsTier string

const (
	SubscribeNewParamsTierStarter  SubscribeNewParamsTier = "starter"
	SubscribeNewParamsTierPro      SubscribeNewParamsTier = "pro"
	SubscribeNewParamsTierBusiness SubscribeNewParamsTier = "business"
)
