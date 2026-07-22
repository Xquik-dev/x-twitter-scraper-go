// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package xtwitterscraper

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/Xquik-dev/x-twitter-scraper-go/internal/apijson"
	"github.com/Xquik-dev/x-twitter-scraper-go/internal/apiquery"
	"github.com/Xquik-dev/x-twitter-scraper-go/internal/requestconfig"
	"github.com/Xquik-dev/x-twitter-scraper-go/option"
	"github.com/Xquik-dev/x-twitter-scraper-go/packages/param"
	"github.com/Xquik-dev/x-twitter-scraper-go/packages/respjson"
)

// Subscription, billing, and credits
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

// Redirect to an active top-up payment page
func (r *CreditService) RedirectTopupCheckout(ctx context.Context, query CreditRedirectTopupCheckoutParams, opts ...option.RequestOption) (err error) {
	var preClientOpts = []option.RequestOption{requestconfig.WithSecurity(requestconfig.Security{})}
	opts = slices.Concat(preClientOpts, r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "credits/topup/redirect"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, nil, opts...)
	return err
}

// Get credits balance
func (r *CreditService) GetBalance(ctx context.Context, opts ...option.RequestOption) (res *CreditGetBalanceResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "credits"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Get top-up billing status
func (r *CreditService) GetTopupStatus(ctx context.Context, query CreditGetTopupStatusParams, opts ...option.RequestOption) (res *CreditGetTopupStatusResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "credits/topup/status"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Create a Stripe Checkout session only after the user confirms. The request never
// completes payment or adds credits by itself.
func (r *CreditService) TopupBalance(ctx context.Context, body CreditTopupBalanceParams, opts ...option.RequestOption) (res *CreditTopupBalanceResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "credits/topup"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type CreditGetBalanceResponse struct {
	// Configured dollar amount for each automatic top-up.
	AutoTopupAmountDollars float64 `json:"auto_topup_amount_dollars" api:"required"`
	AutoTopupEnabled       bool    `json:"auto_topup_enabled" api:"required"`
	// Credit balance threshold that triggers automatic top-up when enabled,
	// represented as a bigint string.
	AutoTopupThreshold string `json:"auto_topup_threshold" api:"required"`
	// Current credit balance as a bigint string to preserve precision above
	// Number.MAX_SAFE_INTEGER.
	Balance string `json:"balance" api:"required"`
	// Lifetime purchased credits as a bigint string.
	LifetimePurchased string `json:"lifetime_purchased" api:"required"`
	// Lifetime consumed credits as a bigint string.
	LifetimeUsed string `json:"lifetime_used" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AutoTopupAmountDollars respjson.Field
		AutoTopupEnabled       respjson.Field
		AutoTopupThreshold     respjson.Field
		Balance                respjson.Field
		LifetimePurchased      respjson.Field
		LifetimeUsed           respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CreditGetBalanceResponse) RawJSON() string { return r.JSON.raw }
func (r *CreditGetBalanceResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CreditGetTopupStatusResponse struct {
	// Any of "paid", "processing", "failed", "expired".
	Status CreditGetTopupStatusResponseStatus `json:"status" api:"required"`
	// Dollar amount requested for the top-up.
	AmountDollars int64 `json:"amount_dollars" api:"nullable"`
	// Bigint string credit amount granted or pending.
	Credits string `json:"credits"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Status        respjson.Field
		AmountDollars respjson.Field
		Credits       respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CreditGetTopupStatusResponse) RawJSON() string { return r.JSON.raw }
func (r *CreditGetTopupStatusResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CreditGetTopupStatusResponseStatus string

const (
	CreditGetTopupStatusResponseStatusPaid       CreditGetTopupStatusResponseStatus = "paid"
	CreditGetTopupStatusResponseStatusProcessing CreditGetTopupStatusResponseStatus = "processing"
	CreditGetTopupStatusResponseStatusFailed     CreditGetTopupStatusResponseStatus = "failed"
	CreditGetTopupStatusResponseStatusExpired    CreditGetTopupStatusResponseStatus = "expired"
)

type CreditTopupBalanceResponse struct {
	// Stable first-party Xquik redirect URL for the active Stripe Checkout session.
	RedirectURL string `json:"redirect_url" api:"required" format:"uri"`
	// Same stable first-party Xquik redirect URL as redirect_url. The response never
	// exposes a raw Stripe Checkout URL.
	URL string `json:"url" api:"required" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		RedirectURL respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CreditTopupBalanceResponse) RawJSON() string { return r.JSON.raw }
func (r *CreditTopupBalanceResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CreditRedirectTopupCheckoutParams struct {
	// Billing session ID returned by the top-up billing flow.
	SessionID string `query:"session_id" api:"required" json:"-"`
	paramObj
}

// URLQuery serializes [CreditRedirectTopupCheckoutParams]'s query parameters as
// `url.Values`.
func (r CreditRedirectTopupCheckoutParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CreditGetTopupStatusParams struct {
	// Top-up session ID to inspect.
	SessionID string `query:"session_id" api:"required" json:"-"`
	paramObj
}

// URLQuery serializes [CreditGetTopupStatusParams]'s query parameters as
// `url.Values`.
func (r CreditGetTopupStatusParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CreditTopupBalanceParams struct {
	// Amount to top up in US dollars. Minimum 10.
	Dollars int64 `json:"dollars" api:"required"`
	// Optional checkout locale. Defaults to en.
	Locale param.Opt[string] `json:"locale,omitzero"`
	paramObj
}

func (r CreditTopupBalanceParams) MarshalJSON() (data []byte, err error) {
	type shadow CreditTopupBalanceParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreditTopupBalanceParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
