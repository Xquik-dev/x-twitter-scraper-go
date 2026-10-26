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

// XAccountConnectionChallengeService covers connected X account management.
//
// XAccountConnectionChallengeService contains methods and other services that help
// with interacting with the x-twitter-scraper API.
//
// You should not instantiate this service directly, and instead use the
// [NewXAccountConnectionChallengeService] method instead.
type XAccountConnectionChallengeService struct {
	options []option.RequestOption
}

// NewXAccountConnectionChallengeService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewXAccountConnectionChallengeService(opts ...option.RequestOption) (r XAccountConnectionChallengeService) {
	r = XAccountConnectionChallengeService{}
	r.options = opts
	return
}

// Submit submits the email code requested during X account connection.
func (r *XAccountConnectionChallengeService) Submit(ctx context.Context, id string, body XAccountConnectionChallengeSubmitParams, opts ...option.RequestOption) (res *XAccountConnectionChallengeSubmitResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("x/account-connection-challenges/%s/submit", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// XAccountConnectionChallengeSubmitResponse is a sanitized X account summary returned by connect and reauth.
type XAccountConnectionChallengeSubmitResponse struct {
	ID        string    `json:"id" api:"required"`
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	ExpiresAt time.Time `json:"expiresAt" format:"date-time"`
	// Any of "healthy", "locked", "needsReauth", "recovering", "suspended",
	// "temporaryIssue".
	Health    XAccountConnectionChallengeSubmitResponseHealth `json:"health"`
	Status    constant.Active                                 `json:"status" default:"active"`
	Object    constant.XAccountConnectionChallenge            `json:"object" default:"x_account_connection_challenge"`
	Message   string                                          `json:"message"`
	Username  string                                          `json:"username"`
	XUserID   string                                          `json:"xUserId"`
	XUsername string                                          `json:"xUsername"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		ExpiresAt   respjson.Field
		Health      respjson.Field
		Status      respjson.Field
		Object      respjson.Field
		Message     respjson.Field
		Username    respjson.Field
		XUserID     respjson.Field
		XUsername   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// RawJSON returns the unmodified JSON received from the API.
func (r XAccountConnectionChallengeSubmitResponse) RawJSON() string { return r.JSON.raw }
func (r *XAccountConnectionChallengeSubmitResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type XAccountConnectionChallengeSubmitResponseHealth string

const (
	XAccountConnectionChallengeSubmitResponseHealthHealthy        XAccountConnectionChallengeSubmitResponseHealth = "healthy"
	XAccountConnectionChallengeSubmitResponseHealthLocked         XAccountConnectionChallengeSubmitResponseHealth = "locked"
	XAccountConnectionChallengeSubmitResponseHealthNeedsReauth    XAccountConnectionChallengeSubmitResponseHealth = "needsReauth"
	XAccountConnectionChallengeSubmitResponseHealthRecovering     XAccountConnectionChallengeSubmitResponseHealth = "recovering"
	XAccountConnectionChallengeSubmitResponseHealthSuspended      XAccountConnectionChallengeSubmitResponseHealth = "suspended"
	XAccountConnectionChallengeSubmitResponseHealthTemporaryIssue XAccountConnectionChallengeSubmitResponseHealth = "temporaryIssue"
)

type XAccountConnectionChallengeSubmitParams struct {
	// Code sent to the account email.
	EmailCode string `json:"email_code" api:"required"`
	paramObj
}

func (r XAccountConnectionChallengeSubmitParams) MarshalJSON() (data []byte, err error) {
	type shadow XAccountConnectionChallengeSubmitParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *XAccountConnectionChallengeSubmitParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, nil)
}
