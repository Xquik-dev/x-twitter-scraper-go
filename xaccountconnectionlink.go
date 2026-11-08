// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

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
	"github.com/Xquik-dev/x-twitter-scraper-go/shared/constant"
)

// Connected X account management
//
// XAccountConnectionLinkService contains methods and other services that help with
// interacting with the x-twitter-scraper API.
//
// You should not instantiate this service directly, and instead use the
// [NewXAccountConnectionLinkService] method instead.
type XAccountConnectionLinkService struct {
	options []option.RequestOption
}

// NewXAccountConnectionLinkService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewXAccountConnectionLinkService(opts ...option.RequestOption) (r XAccountConnectionLinkService) {
	r = XAccountConnectionLinkService{}
	r.options = opts
	return
}

// Creates a link to a page Xquik hosts, where a person connects an X account with
// no Xquik signup. Your backend never receives their password or 2FA secret, and
// the account belongs to you. Send AccountID to reconnect 1 of your accounts
// instead. Connecting is free. Takes a Full access API key or a dashboard session.
// The link works for 30 minutes, and the response is the only place its token
// appears.
func (r *XAccountConnectionLinkService) New(ctx context.Context, body XAccountConnectionLinkNewParams, opts ...option.RequestOption) (res *XAccountConnectionLinkNewResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "x/account-connection-links"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Reads what happened on a link. Status reads connected once the person finishes,
// with XUserID, XUsername and Health. It never holds the token. Takes a Full access
// API key or a dashboard session.
func (r *XAccountConnectionLinkService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *XAccountConnectionLink, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("x/account-connection-links/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Stops a link working. An account it already connected stays connected. Takes a
// Full access API key or a dashboard session.
func (r *XAccountConnectionLinkService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *XAccountConnectionLinkDeleteResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("x/account-connection-links/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// A link where a person connects an X account on Xquik's page.
type XAccountConnectionLink struct {
	ID          string    `json:"id" api:"required"`
	AccountID   string    `json:"accountId" api:"required,nullable"`
	CompletedAt time.Time `json:"completedAt" api:"required,nullable" format:"date-time"`
	CreatedAt   time.Time `json:"createdAt" api:"required" format:"date-time"`
	ExpiresAt   time.Time `json:"expiresAt" api:"required" format:"date-time"`
	// Why the last login didn't finish.
	Failure string                          `json:"failure" api:"required,nullable"`
	Health  string                          `json:"health" api:"required,nullable"`
	Object  constant.XAccountConnectionLink `json:"object" default:"x_account_connection_link"`
	// Any of "connect", "reconnect".
	Purpose   XAccountConnectionLinkPurpose `json:"purpose" api:"required"`
	Reference string                        `json:"reference" api:"required,nullable"`
	ReturnURL string                        `json:"returnUrl" api:"required,nullable" format:"uri"`
	// failed means 5 logins were refused, so make a new link.
	//
	// Any of "pending", "connecting", "requires_email_code", "connected", "failed",
	// "expired", "revoked".
	Status    XAccountConnectionLinkStatus `json:"status" api:"required"`
	XUserID   string                       `json:"xUserId" api:"required,nullable"`
	XUsername string                       `json:"xUsername" api:"required,nullable"`
	// Present only in the response that creates the link.
	URL string `json:"url" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		AccountID   respjson.Field
		CompletedAt respjson.Field
		CreatedAt   respjson.Field
		ExpiresAt   respjson.Field
		Failure     respjson.Field
		Health      respjson.Field
		Object      respjson.Field
		Purpose     respjson.Field
		Reference   respjson.Field
		ReturnURL   respjson.Field
		Status      respjson.Field
		XUserID     respjson.Field
		XUsername   respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XAccountConnectionLink) RawJSON() string { return r.JSON.raw }
func (r *XAccountConnectionLink) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// What a link does.
type XAccountConnectionLinkPurpose string

const (
	XAccountConnectionLinkPurposeConnect   XAccountConnectionLinkPurpose = "connect"
	XAccountConnectionLinkPurposeReconnect XAccountConnectionLinkPurpose = "reconnect"
)

// failed means 5 logins were refused, so make a new link.
type XAccountConnectionLinkStatus string

const (
	XAccountConnectionLinkStatusPending           XAccountConnectionLinkStatus = "pending"
	XAccountConnectionLinkStatusConnecting        XAccountConnectionLinkStatus = "connecting"
	XAccountConnectionLinkStatusRequiresEmailCode XAccountConnectionLinkStatus = "requires_email_code"
	XAccountConnectionLinkStatusConnected         XAccountConnectionLinkStatus = "connected"
	XAccountConnectionLinkStatusFailed            XAccountConnectionLinkStatus = "failed"
	XAccountConnectionLinkStatusExpired           XAccountConnectionLinkStatus = "expired"
	XAccountConnectionLinkStatusRevoked           XAccountConnectionLinkStatus = "revoked"
)

// The new link. It holds the only copy of URL.
type XAccountConnectionLinkNewResponse struct {
	ID          string    `json:"id" api:"required"`
	AccountID   string    `json:"accountId" api:"required,nullable"`
	CompletedAt time.Time `json:"completedAt" api:"required,nullable" format:"date-time"`
	CreatedAt   time.Time `json:"createdAt" api:"required" format:"date-time"`
	ExpiresAt   time.Time `json:"expiresAt" api:"required" format:"date-time"`
	// Why the last login didn't finish.
	Failure string                          `json:"failure" api:"required,nullable"`
	Health  string                          `json:"health" api:"required,nullable"`
	Object  constant.XAccountConnectionLink `json:"object" default:"x_account_connection_link"`
	// Any of "connect", "reconnect".
	Purpose   XAccountConnectionLinkPurpose `json:"purpose" api:"required"`
	Reference string                        `json:"reference" api:"required,nullable"`
	ReturnURL string                        `json:"returnUrl" api:"required,nullable" format:"uri"`
	// failed means 5 logins were refused, so make a new link.
	//
	// Any of "pending", "connecting", "requires_email_code", "connected", "failed",
	// "expired", "revoked".
	Status XAccountConnectionLinkStatus `json:"status" api:"required"`
	// The page to open. Its fragment holds the token, so treat it as a secret and keep
	// it out of logs.
	URL       string `json:"url" api:"required" format:"uri"`
	XUserID   string `json:"xUserId" api:"required,nullable"`
	XUsername string `json:"xUsername" api:"required,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		AccountID   respjson.Field
		CompletedAt respjson.Field
		CreatedAt   respjson.Field
		ExpiresAt   respjson.Field
		Failure     respjson.Field
		Health      respjson.Field
		Object      respjson.Field
		Purpose     respjson.Field
		Reference   respjson.Field
		ReturnURL   respjson.Field
		Status      respjson.Field
		URL         respjson.Field
		XUserID     respjson.Field
		XUsername   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XAccountConnectionLinkNewResponse) RawJSON() string { return r.JSON.raw }
func (r *XAccountConnectionLinkNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The answer to revoking a link.
type XAccountConnectionLinkDeleteResponse struct {
	Success bool `json:"success" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XAccountConnectionLinkDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *XAccountConnectionLinkDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type XAccountConnectionLinkNewParams struct {
	// The ID of 1 of your X accounts. The link then reconnects that account.
	AccountID param.Opt[string] `json:"accountId,omitzero"`
	// Your own ID for this person. It comes back when you read the link.
	Reference param.Opt[string] `json:"reference,omitzero"`
	// A public https URL the page offers as a button once the account connects. Xquik
	// never redirects by itself, and never adds the token.
	ReturnURL param.Opt[string] `json:"returnUrl,omitzero" format:"uri"`
	paramObj
}

func (r XAccountConnectionLinkNewParams) MarshalJSON() (data []byte, err error) {
	type shadow XAccountConnectionLinkNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *XAccountConnectionLinkNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
