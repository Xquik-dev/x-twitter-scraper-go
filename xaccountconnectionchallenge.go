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
// XAccountConnectionChallengeService contains methods and other services that help
// with interacting with the x-twitter-scraper API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewXAccountConnectionChallengeService] method instead.
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

// Submit X account email verification code
func (r *XAccountConnectionChallengeService) Submit(ctx context.Context, id string, body XAccountConnectionChallengeSubmitParams, opts ...option.RequestOption) (res *XAccountConnectionChallengeSubmitResponse, err error) {
	var preClientOpts = []option.RequestOption{requestconfig.WithSecurity(requestconfig.Security{
		APIKey:      true,
		OAuthBearer: true,
	})}
	opts = slices.Concat(preClientOpts, r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("x/account-connection-challenges/%s/submit", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Sanitized X account summary returned by connect and reauth.
type XAccountConnectionChallengeSubmitResponse struct {
	ID        string    `json:"id" api:"required"`
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Any of "healthy", "locked", "needsReauth", "recovering", "suspended",
	// "temporaryIssue".
	Health    XAccountConnectionChallengeSubmitResponseHealth `json:"health" api:"required"`
	Status    string                                          `json:"status" api:"required"`
	XUserID   string                                          `json:"xUserId" api:"required"`
	XUsername string                                          `json:"xUsername" api:"required"`
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
func (r XAccountConnectionChallengeSubmitResponse) RawJSON() string { return r.JSON.raw }
func (r *XAccountConnectionChallengeSubmitResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
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
	return apijson.UnmarshalRoot(data, r)
}
