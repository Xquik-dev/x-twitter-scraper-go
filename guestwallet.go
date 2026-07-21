// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package xtwitterscraper

import (
	"context"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/Xquik-dev/x-twitter-scraper-go/internal/apijson"
	"github.com/Xquik-dev/x-twitter-scraper-go/internal/requestconfig"
	"github.com/Xquik-dev/x-twitter-scraper-go/option"
	"github.com/Xquik-dev/x-twitter-scraper-go/packages/param"
	"github.com/Xquik-dev/x-twitter-scraper-go/packages/respjson"
	"github.com/Xquik-dev/x-twitter-scraper-go/shared/constant"
)

// Accountless prepaid access for paid read endpoints
//
// GuestWalletService contains methods and other services that help with
// interacting with the x-twitter-scraper API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewGuestWalletService] method instead.
type GuestWalletService struct {
	options []option.RequestOption
}

// NewGuestWalletService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewGuestWalletService(opts ...option.RequestOption) (r GuestWalletService) {
	r = GuestWalletService{}
	r.options = opts
	return
}

// Create a one-use Stripe-hosted checkout after the user explicitly confirms a
// $10-$250 USD amount. This request creates no charge by itself. The user opens
// checkout_url on Stripe. This endpoint returns the paid-read API key without
// requiring an Xquik account, email, dashboard, or Xquik web page. An idempotent
// replay returns the same key.
func (r *GuestWalletService) New(ctx context.Context, params GuestWalletNewParams, opts ...option.RequestOption) (res *GuestWalletNewResponse, err error) {
	if !param.IsOmitted(params.IdempotencyKey) {
		opts = append(opts, option.WithHeader("Idempotency-Key", fmt.Sprintf("%v", params.IdempotencyKey)))
	}
	var preClientOpts = []option.RequestOption{requestconfig.WithSecurity(requestconfig.Security{})}
	opts = slices.Concat(preClientOpts, r.options, opts)
	path := "guest-wallets"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Poll after Stripe payment. Use usable to decide whether paid reads can run. An
// active wallet can remain usable while a top-up is pending. A new wallet becomes
// usable only after verified webhook fulfillment. Send the guest key as
// Authorization: Bearer.
func (r *GuestWalletService) GetStatus(ctx context.Context, opts ...option.RequestOption) (res *GuestWalletGetStatusResponse, err error) {
	var preClientOpts = []option.RequestOption{requestconfig.WithAPIKeySecurity()}
	opts = slices.Concat(preClientOpts, r.options, opts)
	path := "guest-wallets/status"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Create a one-use Stripe-hosted checkout for an existing paid-read guest key
// after the user explicitly confirms a $10-$250 USD amount. The key remains the
// same. This request creates no charge by itself and never redirects through an
// Xquik web page.
func (r *GuestWalletService) Topup(ctx context.Context, params GuestWalletTopupParams, opts ...option.RequestOption) (res *GuestWalletTopupResponse, err error) {
	if !param.IsOmitted(params.IdempotencyKey) {
		opts = append(opts, option.WithHeader("Idempotency-Key", fmt.Sprintf("%v", params.IdempotencyKey)))
	}
	var preClientOpts = []option.RequestOption{requestconfig.WithAPIKeySecurity()}
	opts = slices.Concat(preClientOpts, r.options, opts)
	path := "guest-wallets/topups"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Confirmed USD amount for a guest wallet purchase.
type GuestWalletAmount struct {
	// USD amount in cents. Accepted range is $10-$250.
	AmountMinor int64        `json:"amount_minor" api:"required"`
	Currency    constant.Usd `json:"currency" default:"usd"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AmountMinor respjson.Field
		Currency    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GuestWalletAmount) RawJSON() string { return r.JSON.raw }
func (r *GuestWalletAmount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Initial guest wallet response containing the one-time key.
type GuestWalletNewResponse struct {
	AccountRequired bool `json:"account_required" api:"required"`
	// Confirmed USD amount for a guest wallet purchase.
	Amount GuestWalletAmount `json:"amount" api:"required"`
	// Paid-read bearer credential returned only by initial creation. Store it as a
	// secret. Never place it in a URL or log.
	APIKey        string                              `json:"api_key" api:"required" format:"password"`
	Authorization GuestWalletNewResponseAuthorization `json:"authorization" api:"required"`
	// Raw Stripe-hosted checkout URL for user interaction.
	CheckoutURL      string                                                                                             `json:"checkout_url" api:"required" format:"uri"`
	CredentialNotice constant.StoreAPIKeyAndTheIdempotencyKeySecurelyBeforeSharingCheckoutURLNoEmailRecoveryIsAvailable `json:"credential_notice" default:"Store api_key and the Idempotency-Key securely before sharing checkout_url. No email recovery is available."`
	// Credits granted after verified payment.
	Credits string `json:"credits" api:"required"`
	// Time when the pending checkout expires.
	ExpiresAt    time.Time                                                                                                                                                                         `json:"expires_at" api:"required" format:"date-time"`
	Instructions constant.GiveCheckoutURLToTheUserTheyMustCompletePaymentOnStripeNeverSubmitPaymentForThemAfterPaymentPollStatusURLEveryPollAfterSecondsUntilLatestPurchaseStatusIsNoLongerPending `json:"instructions" default:"Give checkout_url to the user. They must complete payment on Stripe. Never submit payment for them. After payment, poll status_url every poll_after_seconds until latest_purchase.status is no longer pending."`
	// Wait at least this long before polling status_url.
	PollAfterSeconds        int64  `json:"poll_after_seconds" api:"required"`
	PurchaseID              string `json:"purchase_id" api:"required"`
	RequiresUserInteraction bool   `json:"requires_user_interaction" api:"required"`
	// Any of "creating", "pending", "paid", "expired", "failed", "refunded",
	// "disputed".
	Status    GuestWalletNewResponseStatus                  `json:"status" api:"required"`
	StatusURL constant.HTTPSXquikComAPIV1GuestWalletsStatus `json:"status_url" default:"https://xquik.com/api/v1/guest-wallets/status"`
	WalletID  string                                        `json:"wallet_id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccountRequired         respjson.Field
		Amount                  respjson.Field
		APIKey                  respjson.Field
		Authorization           respjson.Field
		CheckoutURL             respjson.Field
		CredentialNotice        respjson.Field
		Credits                 respjson.Field
		ExpiresAt               respjson.Field
		Instructions            respjson.Field
		PollAfterSeconds        respjson.Field
		PurchaseID              respjson.Field
		RequiresUserInteraction respjson.Field
		Status                  respjson.Field
		StatusURL               respjson.Field
		WalletID                respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GuestWalletNewResponse) RawJSON() string { return r.JSON.raw }
func (r *GuestWalletNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GuestWalletNewResponseAuthorization struct {
	// Any of "Authorization".
	Header string `json:"header" api:"required"`
	// Any of "Bearer".
	Scheme string `json:"scheme" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Header      respjson.Field
		Scheme      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GuestWalletNewResponseAuthorization) RawJSON() string { return r.JSON.raw }
func (r *GuestWalletNewResponseAuthorization) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GuestWalletNewResponseStatus string

const (
	GuestWalletNewResponseStatusCreating GuestWalletNewResponseStatus = "creating"
	GuestWalletNewResponseStatusPending  GuestWalletNewResponseStatus = "pending"
	GuestWalletNewResponseStatusPaid     GuestWalletNewResponseStatus = "paid"
	GuestWalletNewResponseStatusExpired  GuestWalletNewResponseStatus = "expired"
	GuestWalletNewResponseStatusFailed   GuestWalletNewResponseStatus = "failed"
	GuestWalletNewResponseStatusRefunded GuestWalletNewResponseStatus = "refunded"
	GuestWalletNewResponseStatusDisputed GuestWalletNewResponseStatus = "disputed"
)

// Current balance, usability, and latest guest purchase state.
type GuestWalletGetStatusResponse struct {
	Balance string `json:"balance" api:"required"`
	// Latest guest wallet purchase fulfillment state.
	LatestPurchase GuestWalletGetStatusResponseLatestPurchase `json:"latest_purchase" api:"required"`
	// Polling delay while payment is pending. Null means stop.
	//
	// Any of 2.
	PollAfterSeconds int64              `json:"poll_after_seconds" api:"required"`
	Scope            constant.PaidReads `json:"scope" default:"paid_reads"`
	// Combined wallet and pending-checkout state. A pending top-up can coexist with
	// usable true. Terminal expired or failed states require a new guest wallet.
	//
	// Any of "active", "pending", "expired", "failed", "frozen", "closed".
	Status GuestWalletGetStatusResponseStatus `json:"status" api:"required"`
	// Top-up action when usable and no checkout is pending.
	TopUp GuestWalletGetStatusResponseTopUp `json:"top_up" api:"required"`
	// Authoritative paid-read readiness. Use instead of status.
	Usable   bool   `json:"usable" api:"required"`
	WalletID string `json:"wallet_id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Balance          respjson.Field
		LatestPurchase   respjson.Field
		PollAfterSeconds respjson.Field
		Scope            respjson.Field
		Status           respjson.Field
		TopUp            respjson.Field
		Usable           respjson.Field
		WalletID         respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GuestWalletGetStatusResponse) RawJSON() string { return r.JSON.raw }
func (r *GuestWalletGetStatusResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Latest guest wallet purchase fulfillment state.
type GuestWalletGetStatusResponseLatestPurchase struct {
	// Confirmed USD amount for a guest wallet purchase.
	Amount GuestWalletAmount `json:"amount" api:"required"`
	// Present only while the purchase is pending.
	CheckoutURL string    `json:"checkout_url" api:"required" format:"uri"`
	Credits     string    `json:"credits" api:"required"`
	ExpiresAt   time.Time `json:"expires_at" api:"required" format:"date-time"`
	PurchaseID  string    `json:"purchase_id" api:"required"`
	// Any of "creating", "pending", "paid", "expired", "failed", "refunded",
	// "disputed".
	Status string `json:"status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount      respjson.Field
		CheckoutURL respjson.Field
		Credits     respjson.Field
		ExpiresAt   respjson.Field
		PurchaseID  respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GuestWalletGetStatusResponseLatestPurchase) RawJSON() string { return r.JSON.raw }
func (r *GuestWalletGetStatusResponseLatestPurchase) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Combined wallet and pending-checkout state. A pending top-up can coexist with
// usable true. Terminal expired or failed states require a new guest wallet.
type GuestWalletGetStatusResponseStatus string

const (
	GuestWalletGetStatusResponseStatusActive  GuestWalletGetStatusResponseStatus = "active"
	GuestWalletGetStatusResponseStatusPending GuestWalletGetStatusResponseStatus = "pending"
	GuestWalletGetStatusResponseStatusExpired GuestWalletGetStatusResponseStatus = "expired"
	GuestWalletGetStatusResponseStatusFailed  GuestWalletGetStatusResponseStatus = "failed"
	GuestWalletGetStatusResponseStatusFrozen  GuestWalletGetStatusResponseStatus = "frozen"
	GuestWalletGetStatusResponseStatusClosed  GuestWalletGetStatusResponseStatus = "closed"
)

// Top-up action when usable and no checkout is pending.
type GuestWalletGetStatusResponseTopUp struct {
	Method constant.Post                    `json:"method" default:"POST"`
	Path   constant.APIV1GuestWalletsTopups `json:"path" default:"/api/v1/guest-wallets/topups"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Method      respjson.Field
		Path        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GuestWalletGetStatusResponseTopUp) RawJSON() string { return r.JSON.raw }
func (r *GuestWalletGetStatusResponseTopUp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Pending Stripe checkout and guest wallet purchase details.
type GuestWalletTopupResponse struct {
	AccountRequired bool `json:"account_required" api:"required"`
	// Confirmed USD amount for a guest wallet purchase.
	Amount GuestWalletAmount `json:"amount" api:"required"`
	// Raw Stripe-hosted checkout URL for user interaction.
	CheckoutURL string `json:"checkout_url" api:"required" format:"uri"`
	// Credits granted after verified payment.
	Credits string `json:"credits" api:"required"`
	// Time when the pending checkout expires.
	ExpiresAt    time.Time                                                                                                                                                                         `json:"expires_at" api:"required" format:"date-time"`
	Instructions constant.GiveCheckoutURLToTheUserTheyMustCompletePaymentOnStripeNeverSubmitPaymentForThemAfterPaymentPollStatusURLEveryPollAfterSecondsUntilLatestPurchaseStatusIsNoLongerPending `json:"instructions" default:"Give checkout_url to the user. They must complete payment on Stripe. Never submit payment for them. After payment, poll status_url every poll_after_seconds until latest_purchase.status is no longer pending."`
	// Wait at least this long before polling status_url.
	PollAfterSeconds        int64  `json:"poll_after_seconds" api:"required"`
	PurchaseID              string `json:"purchase_id" api:"required"`
	RequiresUserInteraction bool   `json:"requires_user_interaction" api:"required"`
	// Any of "creating", "pending", "paid", "expired", "failed", "refunded",
	// "disputed".
	Status    GuestWalletTopupResponseStatus                `json:"status" api:"required"`
	StatusURL constant.HTTPSXquikComAPIV1GuestWalletsStatus `json:"status_url" default:"https://xquik.com/api/v1/guest-wallets/status"`
	WalletID  string                                        `json:"wallet_id" api:"required"`
	// Paid-read bearer credential returned only by initial creation. Store it as a
	// secret. Never place it in a URL or log.
	APIKey        string                                `json:"api_key" format:"password"`
	Authorization GuestWalletTopupResponseAuthorization `json:"authorization"`
	// Any of "Store api_key and the Idempotency-Key securely before sharing
	// checkout_url. No email recovery is available.".
	CredentialNotice GuestWalletTopupResponseCredentialNotice `json:"credential_notice"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccountRequired         respjson.Field
		Amount                  respjson.Field
		CheckoutURL             respjson.Field
		Credits                 respjson.Field
		ExpiresAt               respjson.Field
		Instructions            respjson.Field
		PollAfterSeconds        respjson.Field
		PurchaseID              respjson.Field
		RequiresUserInteraction respjson.Field
		Status                  respjson.Field
		StatusURL               respjson.Field
		WalletID                respjson.Field
		APIKey                  respjson.Field
		Authorization           respjson.Field
		CredentialNotice        respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GuestWalletTopupResponse) RawJSON() string { return r.JSON.raw }
func (r *GuestWalletTopupResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GuestWalletTopupResponseStatus string

const (
	GuestWalletTopupResponseStatusCreating GuestWalletTopupResponseStatus = "creating"
	GuestWalletTopupResponseStatusPending  GuestWalletTopupResponseStatus = "pending"
	GuestWalletTopupResponseStatusPaid     GuestWalletTopupResponseStatus = "paid"
	GuestWalletTopupResponseStatusExpired  GuestWalletTopupResponseStatus = "expired"
	GuestWalletTopupResponseStatusFailed   GuestWalletTopupResponseStatus = "failed"
	GuestWalletTopupResponseStatusRefunded GuestWalletTopupResponseStatus = "refunded"
	GuestWalletTopupResponseStatusDisputed GuestWalletTopupResponseStatus = "disputed"
)

type GuestWalletTopupResponseAuthorization struct {
	// Any of "Authorization".
	Header string `json:"header" api:"required"`
	// Any of "Bearer".
	Scheme string `json:"scheme" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Header      respjson.Field
		Scheme      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GuestWalletTopupResponseAuthorization) RawJSON() string { return r.JSON.raw }
func (r *GuestWalletTopupResponseAuthorization) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GuestWalletTopupResponseCredentialNotice string

const (
	GuestWalletTopupResponseCredentialNoticeStoreAPIKeyAndTheIdempotencyKeySecurelyBeforeSharingCheckoutURLNoEmailRecoveryIsAvailable GuestWalletTopupResponseCredentialNotice = "Store api_key and the Idempotency-Key securely before sharing checkout_url. No email recovery is available."
)

type GuestWalletNewParams struct {
	// Confirmed USD amount in cents.
	AmountMinor    int64  `json:"amount_minor" api:"required"`
	IdempotencyKey string `header:"Idempotency-Key" api:"required" json:"-"`
	// This field can be elided, and will marshal its zero value as "usd".
	Currency constant.Usd `json:"currency" default:"usd"`
	paramObj
}

func (r GuestWalletNewParams) MarshalJSON() (data []byte, err error) {
	type shadow GuestWalletNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GuestWalletNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GuestWalletTopupParams struct {
	// Confirmed USD amount in cents.
	AmountMinor    int64  `json:"amount_minor" api:"required"`
	IdempotencyKey string `header:"Idempotency-Key" api:"required" json:"-"`
	// This field can be elided, and will marshal its zero value as "usd".
	Currency constant.Usd `json:"currency" default:"usd"`
	paramObj
}

func (r GuestWalletTopupParams) MarshalJSON() (data []byte, err error) {
	type shadow GuestWalletTopupParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GuestWalletTopupParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
