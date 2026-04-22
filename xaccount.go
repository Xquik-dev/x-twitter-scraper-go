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

// Connected X account management
//
// XAccountService contains methods and other services that help with interacting
// with the x-twitter-scraper API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewXAccountService] method instead.
type XAccountService struct {
	options []option.RequestOption
}

// NewXAccountService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewXAccountService(opts ...option.RequestOption) (r XAccountService) {
	r = XAccountService{}
	r.options = opts
	return
}

// Connect X account
func (r *XAccountService) New(ctx context.Context, body XAccountNewParams, opts ...option.RequestOption) (res *XAccountNewResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "x/accounts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Get X account details
func (r *XAccountService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *XAccountDetail, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("x/accounts/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// List connected X accounts
func (r *XAccountService) List(ctx context.Context, opts ...option.RequestOption) (res *XAccountListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "x/accounts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Disconnect X account
func (r *XAccountService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *XAccountDeleteResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("x/accounts/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Clears loginFailedAt and loginFailureReason for all accounts with transient or
// automated failure reasons, making them eligible for retry on next use.
func (r *XAccountService) BulkRetry(ctx context.Context, opts ...option.RequestOption) (res *XAccountBulkRetryResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "x/accounts/bulk-retry"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Re-authenticate X account
func (r *XAccountService) Reauth(ctx context.Context, id string, body XAccountReauthParams, opts ...option.RequestOption) (res *XAccountReauthResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("x/accounts/%s/reauth", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Linked X account summary with username and connection status.
type XAccount struct {
	ID        string    `json:"id" api:"required"`
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Derived login/cookie health. `healthy` = cookies valid. `needsReauth` = user
	// must submit fresh credentials. `locked` = X locked the account; unlock on x.com
	// first. `suspended` = X banned the account. `recovering` = past cooldown, will
	// auto-retry on next use. `temporaryIssue` = transient backend problem; retry
	// shortly.
	//
	// Any of "healthy", "locked", "needsReauth", "recovering", "suspended",
	// "temporaryIssue".
	Health    XAccountHealth `json:"health" api:"required"`
	Status    string         `json:"status" api:"required"`
	XUserID   string         `json:"xUserId" api:"required"`
	XUsername string         `json:"xUsername" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Health      respjson.Field
		Status      respjson.Field
		XUserID     respjson.Field
		XUsername   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XAccount) RawJSON() string { return r.JSON.raw }
func (r *XAccount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Derived login/cookie health. `healthy` = cookies valid. `needsReauth` = user
// must submit fresh credentials. `locked` = X locked the account; unlock on x.com
// first. `suspended` = X banned the account. `recovering` = past cooldown, will
// auto-retry on next use. `temporaryIssue` = transient backend problem; retry
// shortly.
type XAccountHealth string

const (
	XAccountHealthHealthy        XAccountHealth = "healthy"
	XAccountHealthLocked         XAccountHealth = "locked"
	XAccountHealthNeedsReauth    XAccountHealth = "needsReauth"
	XAccountHealthRecovering     XAccountHealth = "recovering"
	XAccountHealthSuspended      XAccountHealth = "suspended"
	XAccountHealthTemporaryIssue XAccountHealth = "temporaryIssue"
)

// Full X account details including proxy, cookies, and update timestamp.
type XAccountDetail struct {
	ID        string    `json:"id" api:"required"`
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Any of "healthy", "locked", "needsReauth", "recovering", "suspended",
	// "temporaryIssue".
	Health            XAccountDetailHealth `json:"health" api:"required"`
	Status            string               `json:"status" api:"required"`
	XUserID           string               `json:"xUserId" api:"required"`
	XUsername         string               `json:"xUsername" api:"required"`
	CookiesObtainedAt time.Time            `json:"cookiesObtainedAt" format:"date-time"`
	ProxyCountry      string               `json:"proxyCountry"`
	UpdatedAt         time.Time            `json:"updatedAt" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		CreatedAt         respjson.Field
		Health            respjson.Field
		Status            respjson.Field
		XUserID           respjson.Field
		XUsername         respjson.Field
		CookiesObtainedAt respjson.Field
		ProxyCountry      respjson.Field
		UpdatedAt         respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XAccountDetail) RawJSON() string { return r.JSON.raw }
func (r *XAccountDetail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type XAccountDetailHealth string

const (
	XAccountDetailHealthHealthy        XAccountDetailHealth = "healthy"
	XAccountDetailHealthLocked         XAccountDetailHealth = "locked"
	XAccountDetailHealthNeedsReauth    XAccountDetailHealth = "needsReauth"
	XAccountDetailHealthRecovering     XAccountDetailHealth = "recovering"
	XAccountDetailHealthSuspended      XAccountDetailHealth = "suspended"
	XAccountDetailHealthTemporaryIssue XAccountDetailHealth = "temporaryIssue"
)

// Sanitized X account summary returned by connect and reauth. Includes an optional
// `loginCountry` field surfaced only when the declared proxy region had no Driver
// capacity and the login fell back to a single US consumer device for this
// one-time action. Future activity continues to use the selected `proxy_country`;
// the field is omitted on normal logins.
type XAccountNewResponse struct {
	ID        string    `json:"id" api:"required"`
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Any of "healthy", "locked", "needsReauth", "recovering", "suspended",
	// "temporaryIssue".
	Health    XAccountNewResponseHealth `json:"health" api:"required"`
	Status    string                    `json:"status" api:"required"`
	XUserID   string                    `json:"xUserId" api:"required"`
	XUsername string                    `json:"xUsername" api:"required"`
	// ISO-3166-1 alpha-2 country code of the Driver consumer device used for this
	// login. Present only when the US fallback was triggered because Driver had no
	// capacity in the declared region. Omitted otherwise.
	LoginCountry string `json:"loginCountry"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		CreatedAt    respjson.Field
		Health       respjson.Field
		Status       respjson.Field
		XUserID      respjson.Field
		XUsername    respjson.Field
		LoginCountry respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XAccountNewResponse) RawJSON() string { return r.JSON.raw }
func (r *XAccountNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type XAccountNewResponseHealth string

const (
	XAccountNewResponseHealthHealthy        XAccountNewResponseHealth = "healthy"
	XAccountNewResponseHealthLocked         XAccountNewResponseHealth = "locked"
	XAccountNewResponseHealthNeedsReauth    XAccountNewResponseHealth = "needsReauth"
	XAccountNewResponseHealthRecovering     XAccountNewResponseHealth = "recovering"
	XAccountNewResponseHealthSuspended      XAccountNewResponseHealth = "suspended"
	XAccountNewResponseHealthTemporaryIssue XAccountNewResponseHealth = "temporaryIssue"
)

type XAccountListResponse struct {
	Accounts []XAccount `json:"accounts" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Accounts    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XAccountListResponse) RawJSON() string { return r.JSON.raw }
func (r *XAccountListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type XAccountDeleteResponse struct {
	Success bool `json:"success" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XAccountDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *XAccountDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type XAccountBulkRetryResponse struct {
	// Number of accounts cleared
	Cleared int64 `json:"cleared" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Cleared     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XAccountBulkRetryResponse) RawJSON() string { return r.JSON.raw }
func (r *XAccountBulkRetryResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Sanitized X account summary returned by connect and reauth. Includes an optional
// `loginCountry` field surfaced only when the declared proxy region had no Driver
// capacity and the login fell back to a single US consumer device for this
// one-time action. Future activity continues to use the selected `proxy_country`;
// the field is omitted on normal logins.
type XAccountReauthResponse struct {
	ID        string    `json:"id" api:"required"`
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Any of "healthy", "locked", "needsReauth", "recovering", "suspended",
	// "temporaryIssue".
	Health    XAccountReauthResponseHealth `json:"health" api:"required"`
	Status    string                       `json:"status" api:"required"`
	XUserID   string                       `json:"xUserId" api:"required"`
	XUsername string                       `json:"xUsername" api:"required"`
	// ISO-3166-1 alpha-2 country code of the Driver consumer device used for this
	// login. Present only when the US fallback was triggered because Driver had no
	// capacity in the declared region. Omitted otherwise.
	LoginCountry string `json:"loginCountry"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		CreatedAt    respjson.Field
		Health       respjson.Field
		Status       respjson.Field
		XUserID      respjson.Field
		XUsername    respjson.Field
		LoginCountry respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XAccountReauthResponse) RawJSON() string { return r.JSON.raw }
func (r *XAccountReauthResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type XAccountReauthResponseHealth string

const (
	XAccountReauthResponseHealthHealthy        XAccountReauthResponseHealth = "healthy"
	XAccountReauthResponseHealthLocked         XAccountReauthResponseHealth = "locked"
	XAccountReauthResponseHealthNeedsReauth    XAccountReauthResponseHealth = "needsReauth"
	XAccountReauthResponseHealthRecovering     XAccountReauthResponseHealth = "recovering"
	XAccountReauthResponseHealthSuspended      XAccountReauthResponseHealth = "suspended"
	XAccountReauthResponseHealthTemporaryIssue XAccountReauthResponseHealth = "temporaryIssue"
)

type XAccountNewParams struct {
	// Account email
	Email string `json:"email" api:"required"`
	// Account password
	Password string `json:"password" api:"required"`
	// X username
	Username string `json:"username" api:"required"`
	// Proxy country code
	ProxyCountry param.Opt[string] `json:"proxy_country,omitzero"`
	// TOTP secret for 2FA
	TotpSecret param.Opt[string] `json:"totp_secret,omitzero"`
	paramObj
}

func (r XAccountNewParams) MarshalJSON() (data []byte, err error) {
	type shadow XAccountNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *XAccountNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type XAccountReauthParams struct {
	// Updated account password
	Password string `json:"password" api:"required"`
	// Email for the X account (updates stored email)
	Email param.Opt[string] `json:"email,omitzero"`
	// Two-letter country code for login proxy region
	ProxyCountry param.Opt[string] `json:"proxy_country,omitzero"`
	// TOTP secret for 2FA re-authentication
	TotpSecret param.Opt[string] `json:"totp_secret,omitzero"`
	paramObj
}

func (r XAccountReauthParams) MarshalJSON() (data []byte, err error) {
	type shadow XAccountReauthParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *XAccountReauthParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
