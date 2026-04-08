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

// Subscription & billing
//
// CreditService contains methods and other services that help with interacting
// with the x-twitter-scraper API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCreditService] method instead.
type CreditService struct {
	options []option.RequestOption
}

// NewCreditService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewCreditService(opts ...option.RequestOption) (r CreditService) {
	r = CreditService{}
	r.options = opts
	return
}

// Get credits balance
func (r *CreditService) GetBalance(ctx context.Context, opts ...option.RequestOption) (res *CreditGetBalanceResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "credits"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Top up credits balance
func (r *CreditService) TopupBalance(ctx context.Context, body CreditTopupBalanceParams, opts ...option.RequestOption) (res *CreditTopupBalanceResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "credits/topup"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type CreditGetBalanceResponse struct {
	AutoTopupEnabled  bool  `json:"auto_topup_enabled" api:"required"`
	Balance           int64 `json:"balance" api:"required"`
	LifetimePurchased int64 `json:"lifetime_purchased" api:"required"`
	LifetimeUsed      int64 `json:"lifetime_used" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AutoTopupEnabled  respjson.Field
		Balance           respjson.Field
		LifetimePurchased respjson.Field
		LifetimeUsed      respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CreditGetBalanceResponse) RawJSON() string { return r.JSON.raw }
func (r *CreditGetBalanceResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CreditTopupBalanceResponse struct {
	Success bool `json:"success" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CreditTopupBalanceResponse) RawJSON() string { return r.JSON.raw }
func (r *CreditTopupBalanceResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CreditTopupBalanceParams struct {
	// Amount to top up in credits
	Amount int64 `json:"amount" api:"required"`
	paramObj
}

func (r CreditTopupBalanceParams) MarshalJSON() (data []byte, err error) {
	type shadow CreditTopupBalanceParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreditTopupBalanceParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
