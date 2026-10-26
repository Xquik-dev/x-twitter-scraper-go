// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

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

// SubscribeService covers subscription, billing, and credits.
//
// SubscribeService contains methods and other services that help with interacting
// with the x-twitter-scraper API.
//
// You should not instantiate this service directly, and instead use the
// [NewSubscribeService] method instead.
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

// New creates a subscription checkout or billing-management URL only after the user
// confirms. The request never completes payment by itself.
func (r *SubscribeService) New(ctx context.Context, body SubscribeNewParams, opts ...option.RequestOption) (res *SubscribeNewResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "subscribe"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type SubscribeNewResponse struct {
	// Currency the checkout charges. Only a checkout_created answer has it.
	//
	// Any of "USD", "EUR", "GBP", "TRY", "JPY", "CNY", "KRW", "CAD", "AUD",
	// "MXN".
	ChargedCurrency Currency `json:"chargedCurrency"`
	Message         string   `json:"message" api:"required"`
	// Any of "checkout_created", "already_subscribed", "payment_issue".
	Status SubscribeNewResponseStatus `json:"status" api:"required"`
	URL    string                     `json:"url" api:"required" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChargedCurrency respjson.Field
		Message         respjson.Field
		Status          respjson.Field
		URL             respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// RawJSON returns the unmodified JSON received from the API.
func (r SubscribeNewResponse) RawJSON() string { return r.JSON.raw }
func (r *SubscribeNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type SubscribeNewResponseStatus string

const (
	SubscribeNewResponseStatusCheckoutCreated   SubscribeNewResponseStatus = "checkout_created"
	SubscribeNewResponseStatusAlreadySubscribed SubscribeNewResponseStatus = "already_subscribed"
	SubscribeNewResponseStatusPaymentIssue      SubscribeNewResponseStatus = "payment_issue"
)

type SubscribeNewParams struct {
	// Currency to pay in. Lowercase codes work too. Xquik converts the USD plan
	// price at the latest daily ECB rate. Each renewal charges the same amount
	// before tax. An account that already pays in USD gets a USD checkout. With
	// rates over 3 ECB working days old, the checkout charges USD.
	// `chargedCurrency` shows the currency charged.
	//
	// Any of "USD", "EUR", "GBP", "TRY", "JPY", "CNY", "KRW", "CAD", "AUD",
	// "MXN".
	Currency Currency `json:"currency,omitzero"`
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
	return apijson.UnmarshalRoot(data, r, nil)
}

// SubscribeNewParamsTier is a subscription tier to pre-select.
type SubscribeNewParamsTier string

const (
	SubscribeNewParamsTierStarter  SubscribeNewParamsTier = "starter"
	SubscribeNewParamsTierPro      SubscribeNewParamsTier = "pro"
	SubscribeNewParamsTierBusiness SubscribeNewParamsTier = "business"
)
