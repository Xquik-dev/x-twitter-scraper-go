// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

package xtwitterscraper

import (
	"context"
	"net/http"
	"slices"
	"time"

	"github.com/Xquik-dev/x-twitter-scraper-go/internal/apijson"
	"github.com/Xquik-dev/x-twitter-scraper-go/internal/requestconfig"
	"github.com/Xquik-dev/x-twitter-scraper-go/option"
	"github.com/Xquik-dev/x-twitter-scraper-go/packages/param"
	"github.com/Xquik-dev/x-twitter-scraper-go/packages/respjson"
)

// Account info and settings
//
// AccountService contains methods and other services that help with interacting
// with the x-twitter-scraper API.
//
// You should not instantiate this service directly, and instead use the
// [NewAccountService] method instead.
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

// Returns plan, credit, monitoring, and account settings for the authenticated
// user.
func (r *AccountService) Get(ctx context.Context, opts ...option.RequestOption) (res *AccountGetResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "account"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Links an X username to the authenticated Xquik account.
func (r *AccountService) SetXUsername(ctx context.Context, body AccountSetXUsernameParams, opts ...option.RequestOption) (res *AccountSetXUsernameResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "account/x-identity"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

// Changes the locale used for account responses and notifications.
func (r *AccountService) UpdateLocale(ctx context.Context, body AccountUpdateLocaleParams, opts ...option.RequestOption) (res *AccountUpdateLocaleResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "account"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

type AccountGetResponse struct {
	MonitorBilling AccountGetResponseMonitorBilling `json:"monitorBilling" api:"required"`
	// Deprecated. Monitor slots are unlimited, so this is always
	// Number.MAX_SAFE_INTEGER.
	//
	// Deprecated: Monitor slots are unlimited. Use monitorBilling.unlimitedSlots
	// instead.
	MonitorsAllowed int64 `json:"monitorsAllowed" api:"required"`
	MonitorsUsed    int64 `json:"monitorsUsed" api:"required"`
	// Any of "active", "inactive".
	Plan AccountGetResponsePlan `json:"plan" api:"required"`
	// Subscription state without billing provider identifiers.
	Subscription AccountGetResponseSubscription `json:"subscription" api:"required"`
	CreditInfo   AccountGetResponseCreditInfo   `json:"creditInfo"`
	// Linked X username, omitted when no X account is connected.
	XUsername string `json:"xUsername"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MonitorBilling  respjson.Field
		MonitorsAllowed respjson.Field
		MonitorsUsed    respjson.Field
		Plan            respjson.Field
		Subscription    respjson.Field
		CreditInfo      respjson.Field
		XUsername       respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccountGetResponse) RawJSON() string { return r.JSON.raw }
func (r *AccountGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Subscription state without billing provider identifiers.
type AccountGetResponseSubscription struct {
	// Any of "none", "active", "past_due", "canceling", "incomplete".
	Status AccountGetResponseSubscriptionStatus `json:"status" api:"required"`
	// Renewal date, or the end date when canceling.
	CurrentPeriodEnd   time.Time `json:"currentPeriodEnd" format:"date-time"`
	CurrentPeriodStart time.Time `json:"currentPeriodStart" format:"date-time"`
	// Any of "starter", "pro", "business".
	Tier AccountGetResponseSubscriptionTier `json:"tier"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Status             respjson.Field
		CurrentPeriodEnd   respjson.Field
		CurrentPeriodStart respjson.Field
		Tier               respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccountGetResponseSubscription) RawJSON() string { return r.JSON.raw }
func (r *AccountGetResponseSubscription) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountGetResponseSubscriptionStatus string

const (
	AccountGetResponseSubscriptionStatusNone       AccountGetResponseSubscriptionStatus = "none"
	AccountGetResponseSubscriptionStatusActive     AccountGetResponseSubscriptionStatus = "active"
	AccountGetResponseSubscriptionStatusPastDue    AccountGetResponseSubscriptionStatus = "past_due"
	AccountGetResponseSubscriptionStatusCanceling  AccountGetResponseSubscriptionStatus = "canceling"
	AccountGetResponseSubscriptionStatusIncomplete AccountGetResponseSubscriptionStatus = "incomplete"
)

type AccountGetResponseSubscriptionTier string

const (
	AccountGetResponseSubscriptionTierStarter  AccountGetResponseSubscriptionTier = "starter"
	AccountGetResponseSubscriptionTierPro      AccountGetResponseSubscriptionTier = "pro"
	AccountGetResponseSubscriptionTierBusiness AccountGetResponseSubscriptionTier = "business"
)

type AccountGetResponseMonitorBilling struct {
	// Estimated daily credits for currently active monitors.
	ActiveDailyEstimate string `json:"activeDailyEstimate" api:"required"`
	// Credits charged each hour for currently active monitors.
	ActiveHourlyBurn string `json:"activeHourlyBurn" api:"required"`
	// Rounded daily estimate for 1 active monitor.
	CreditsPerActiveMonitorDay string `json:"creditsPerActiveMonitorDay" api:"required"`
	// Hourly credits charged for 1 active monitor.
	CreditsPerActiveMonitorHour string `json:"creditsPerActiveMonitorHour" api:"required"`
	// Webhook and event deliveries are included in monitor billing.
	EventsIncluded bool `json:"eventsIncluded" api:"required"`
	// Active monitors check every 1 second.
	InstantCheckIntervalSeconds int64 `json:"instantCheckIntervalSeconds" api:"required"`
	// Monitor slot count is unlimited.
	UnlimitedSlots bool `json:"unlimitedSlots" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ActiveDailyEstimate         respjson.Field
		ActiveHourlyBurn            respjson.Field
		CreditsPerActiveMonitorDay  respjson.Field
		CreditsPerActiveMonitorHour respjson.Field
		EventsIncluded              respjson.Field
		InstantCheckIntervalSeconds respjson.Field
		UnlimitedSlots              respjson.Field
		ExtraFields                 map[string]respjson.Field
		raw                         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccountGetResponseMonitorBilling) RawJSON() string { return r.JSON.raw }
func (r *AccountGetResponseMonitorBilling) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountGetResponsePlan string

const (
	AccountGetResponsePlanActive   AccountGetResponsePlan = "active"
	AccountGetResponsePlanInactive AccountGetResponsePlan = "inactive"
)

type AccountGetResponseCreditInfo struct {
	// Dollar amount charged when automatic top-up runs.
	AutoTopupAmountDollars float64 `json:"autoTopupAmountDollars" api:"required"`
	AutoTopupEnabled       bool    `json:"autoTopupEnabled" api:"required"`
	// True when 3 declined charges in a row on 1 card stopped automatic top-up. The
	// setting stays enabled. A paid top-up, a paid plan renewal, or the next automatic
	// top-up on a new card clears it.
	AutoTopupStopped bool `json:"autoTopupStopped" api:"required"`
	// Bigint string threshold that triggers automatic top-up when enabled.
	AutoTopupThreshold string `json:"autoTopupThreshold" api:"required"`
	// Bigint string to preserve precision above Number.MAX_SAFE_INTEGER.
	Balance string `json:"balance" api:"required"`
	// Total purchased credits as a bigint string.
	LifetimePurchased string `json:"lifetimePurchased" api:"required"`
	// Total consumed credits as a bigint string.
	LifetimeUsed string `json:"lifetimeUsed" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AutoTopupAmountDollars respjson.Field
		AutoTopupEnabled       respjson.Field
		AutoTopupStopped       respjson.Field
		AutoTopupThreshold     respjson.Field
		Balance                respjson.Field
		LifetimePurchased      respjson.Field
		LifetimeUsed           respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
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
	// Fields the request sent that this route does not read, in the order sent. They
	// changed nothing. The answer leaves this out when the route read every field.
	IgnoredFields []string `json:"ignoredFields"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Success       respjson.Field
		IgnoredFields respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
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
