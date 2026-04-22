// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package xtwitterscraper

import (
	"context"
	"net/http"
	"slices"

	"github.com/stainless-sdks/x-twitter-scraper-go/internal/apijson"
	"github.com/stainless-sdks/x-twitter-scraper-go/internal/requestconfig"
	"github.com/stainless-sdks/x-twitter-scraper-go/option"
	"github.com/stainless-sdks/x-twitter-scraper-go/packages/param"
	"github.com/stainless-sdks/x-twitter-scraper-go/packages/respjson"
)

// Account info and settings
//
// AccountService contains methods and other services that help with interacting
// with the x-twitter-scraper API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountService] method instead.
type AccountService struct {
	options []option.RequestOption
}

// NewAccountService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAccountService(opts ...option.RequestOption) (r AccountService) {
	r = AccountService{}
	r.options = opts
	return
}

// Get account info
func (r *AccountService) Get(ctx context.Context, opts ...option.RequestOption) (res *AccountGetResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "account"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Set linked X username
func (r *AccountService) SetXUsername(ctx context.Context, body AccountSetXUsernameParams, opts ...option.RequestOption) (res *AccountSetXUsernameResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "account/x-identity"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

// Update account locale
func (r *AccountService) UpdateLocale(ctx context.Context, body AccountUpdateLocaleParams, opts ...option.RequestOption) (res *AccountUpdateLocaleResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "account"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

type AccountGetResponse struct {
	MonitorsAllowed int64 `json:"monitorsAllowed" api:"required"`
	MonitorsUsed    int64 `json:"monitorsUsed" api:"required"`
	// Any of "active", "inactive".
	Plan       AccountGetResponsePlan       `json:"plan" api:"required"`
	CreditInfo AccountGetResponseCreditInfo `json:"creditInfo"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MonitorsAllowed respjson.Field
		MonitorsUsed    respjson.Field
		Plan            respjson.Field
		CreditInfo      respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccountGetResponse) RawJSON() string { return r.JSON.raw }
func (r *AccountGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountGetResponsePlan string

const (
	AccountGetResponsePlanActive   AccountGetResponsePlan = "active"
	AccountGetResponsePlanInactive AccountGetResponsePlan = "inactive"
)

type AccountGetResponseCreditInfo struct {
	AutoTopupEnabled  bool  `json:"autoTopupEnabled" api:"required"`
	Balance           int64 `json:"balance" api:"required"`
	LifetimePurchased int64 `json:"lifetimePurchased" api:"required"`
	LifetimeUsed      int64 `json:"lifetimeUsed" api:"required"`
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
func (r AccountGetResponseCreditInfo) RawJSON() string { return r.JSON.raw }
func (r *AccountGetResponseCreditInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountSetXUsernameResponse struct {
	Success   bool   `json:"success" api:"required"`
	XUsername string `json:"xUsername" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Success     respjson.Field
		XUsername   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccountSetXUsernameResponse) RawJSON() string { return r.JSON.raw }
func (r *AccountSetXUsernameResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountUpdateLocaleResponse struct {
	Success bool `json:"success" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccountUpdateLocaleResponse) RawJSON() string { return r.JSON.raw }
func (r *AccountUpdateLocaleResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountSetXUsernameParams struct {
	// X username without @
	Username string `json:"username" api:"required"`
	paramObj
}

func (r AccountSetXUsernameParams) MarshalJSON() (data []byte, err error) {
	type shadow AccountSetXUsernameParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AccountSetXUsernameParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountUpdateLocaleParams struct {
	// Any of "en", "tr", "es".
	Locale AccountUpdateLocaleParamsLocale `json:"locale,omitzero" api:"required"`
	paramObj
}

func (r AccountUpdateLocaleParams) MarshalJSON() (data []byte, err error) {
	type shadow AccountUpdateLocaleParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AccountUpdateLocaleParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountUpdateLocaleParamsLocale string

const (
	AccountUpdateLocaleParamsLocaleEn AccountUpdateLocaleParamsLocale = "en"
	AccountUpdateLocaleParamsLocaleTr AccountUpdateLocaleParamsLocale = "tr"
	AccountUpdateLocaleParamsLocaleEs AccountUpdateLocaleParamsLocale = "es"
)
