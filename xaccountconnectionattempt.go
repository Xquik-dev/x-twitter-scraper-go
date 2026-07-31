// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package xtwitterscraper

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/Xquik-dev/x-twitter-scraper-go/internal/apijson"
	"github.com/Xquik-dev/x-twitter-scraper-go/internal/requestconfig"
	"github.com/Xquik-dev/x-twitter-scraper-go/option"
	"github.com/Xquik-dev/x-twitter-scraper-go/packages/respjson"
	"github.com/Xquik-dev/x-twitter-scraper-go/shared/constant"
)

// Connected X account management
//
// XAccountConnectionAttemptService contains methods and other services that help
// with interacting with the x-twitter-scraper API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewXAccountConnectionAttemptService] method instead.
type XAccountConnectionAttemptService struct {
	options []option.RequestOption
}

// NewXAccountConnectionAttemptService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewXAccountConnectionAttemptService(opts ...option.RequestOption) (r XAccountConnectionAttemptService) {
	r = XAccountConnectionAttemptService{}
	r.options = opts
	return
}

// Get X account connection status
func (r *XAccountConnectionAttemptService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *XAccountConnectionAttemptGetResponseUnion, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("x/account-connection-attempts/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// XAccountConnectionAttemptGetResponseUnion contains all possible properties and
// values from [XAccountConnectionAttemptGetResponsePending],
// [XAccountConnectionAttemptGetResponseSuccess],
// [XAccountConnectionAttemptGetResponseFailed],
// [XAccountConnectionAttemptGetResponseRequiresEmailCode].
//
// Use the [XAccountConnectionAttemptGetResponseUnion.AsAny] method to switch on
// the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type XAccountConnectionAttemptGetResponseUnion struct {
	ID     string `json:"id"`
	Object string `json:"object"`
	// This field is from variant [XAccountConnectionAttemptGetResponsePending].
	PollAfterMs int64 `json:"pollAfterMs"`
	// Any of "pending", "success", "failed", "requires_email_code".
	Status string `json:"status"`
	// This field is from variant [XAccountConnectionAttemptGetResponseFailed].
	Error string `json:"error"`
	// This field is from variant [XAccountConnectionAttemptGetResponseFailed].
	Retryable bool `json:"retryable"`
	// This field is from variant [XAccountConnectionAttemptGetResponseFailed].
	Reason string `json:"reason"`
	// This field is from variant
	// [XAccountConnectionAttemptGetResponseRequiresEmailCode].
	ExpiresAt time.Time `json:"expiresAt"`
	// This field is from variant
	// [XAccountConnectionAttemptGetResponseRequiresEmailCode].
	Message string `json:"message"`
	// This field is from variant
	// [XAccountConnectionAttemptGetResponseRequiresEmailCode].
	Username string `json:"username"`
	JSON     struct {
		ID          respjson.Field
		Object      respjson.Field
		PollAfterMs respjson.Field
		Status      respjson.Field
		Error       respjson.Field
		Retryable   respjson.Field
		Reason      respjson.Field
		ExpiresAt   respjson.Field
		Message     respjson.Field
		Username    respjson.Field
		raw         string
	} `json:"-"`
}

// anyXAccountConnectionAttemptGetResponse is implemented by each variant of
// [XAccountConnectionAttemptGetResponseUnion] to add type safety for the return
// type of [XAccountConnectionAttemptGetResponseUnion.AsAny]
type anyXAccountConnectionAttemptGetResponse interface {
	implXAccountConnectionAttemptGetResponseUnion()
}

func (XAccountConnectionAttemptGetResponsePending) implXAccountConnectionAttemptGetResponseUnion() {}
func (XAccountConnectionAttemptGetResponseSuccess) implXAccountConnectionAttemptGetResponseUnion() {}
func (XAccountConnectionAttemptGetResponseFailed) implXAccountConnectionAttemptGetResponseUnion()  {}
func (XAccountConnectionAttemptGetResponseRequiresEmailCode) implXAccountConnectionAttemptGetResponseUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := XAccountConnectionAttemptGetResponseUnion.AsAny().(type) {
//	case xtwitterscraper.XAccountConnectionAttemptGetResponsePending:
//	case xtwitterscraper.XAccountConnectionAttemptGetResponseSuccess:
//	case xtwitterscraper.XAccountConnectionAttemptGetResponseFailed:
//	case xtwitterscraper.XAccountConnectionAttemptGetResponseRequiresEmailCode:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u XAccountConnectionAttemptGetResponseUnion) AsAny() anyXAccountConnectionAttemptGetResponse {
	switch u.Status {
	case "pending":
		return u.AsPending()
	case "success":
		return u.AsSuccess()
	case "failed":
		return u.AsFailed()
	case "requires_email_code":
		return u.AsRequiresEmailCode()
	}
	return nil
}

func (u XAccountConnectionAttemptGetResponseUnion) AsPending() (v XAccountConnectionAttemptGetResponsePending) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u XAccountConnectionAttemptGetResponseUnion) AsSuccess() (v XAccountConnectionAttemptGetResponseSuccess) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u XAccountConnectionAttemptGetResponseUnion) AsFailed() (v XAccountConnectionAttemptGetResponseFailed) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u XAccountConnectionAttemptGetResponseUnion) AsRequiresEmailCode() (v XAccountConnectionAttemptGetResponseRequiresEmailCode) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u XAccountConnectionAttemptGetResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *XAccountConnectionAttemptGetResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The connection is still in progress.
type XAccountConnectionAttemptGetResponsePending struct {
	ID          string                             `json:"id" api:"required"`
	Object      constant.XAccountConnectionAttempt `json:"object" default:"x_account_connection_attempt"`
	PollAfterMs int64                              `json:"pollAfterMs" api:"required"`
	Status      constant.Pending                   `json:"status" default:"pending"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Object      respjson.Field
		PollAfterMs respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XAccountConnectionAttemptGetResponsePending) RawJSON() string { return r.JSON.raw }
func (r *XAccountConnectionAttemptGetResponsePending) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The account connected successfully.
type XAccountConnectionAttemptGetResponseSuccess struct {
	ID     string                             `json:"id" api:"required"`
	Object constant.XAccountConnectionAttempt `json:"object" default:"x_account_connection_attempt"`
	Status constant.Success                   `json:"status" default:"success"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Object      respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XAccountConnectionAttemptGetResponseSuccess) RawJSON() string { return r.JSON.raw }
func (r *XAccountConnectionAttemptGetResponseSuccess) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The connection reached a final failure.
type XAccountConnectionAttemptGetResponseFailed struct {
	ID        string                             `json:"id" api:"required"`
	Error     string                             `json:"error" api:"required"`
	Object    constant.XAccountConnectionAttempt `json:"object" default:"x_account_connection_attempt"`
	Retryable bool                               `json:"retryable" api:"required"`
	Status    constant.Failed                    `json:"status" default:"failed"`
	Reason    string                             `json:"reason"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Error       respjson.Field
		Object      respjson.Field
		Retryable   respjson.Field
		Status      respjson.Field
		Reason      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XAccountConnectionAttemptGetResponseFailed) RawJSON() string { return r.JSON.raw }
func (r *XAccountConnectionAttemptGetResponseFailed) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resumable account connection challenge. Submit the email code to finish the same
// connection attempt.
type XAccountConnectionAttemptGetResponseRequiresEmailCode struct {
	ID        string                               `json:"id" api:"required"`
	ExpiresAt time.Time                            `json:"expiresAt" api:"required" format:"date-time"`
	Message   string                               `json:"message" api:"required"`
	Object    constant.XAccountConnectionChallenge `json:"object" default:"x_account_connection_challenge"`
	Status    constant.RequiresEmailCode           `json:"status" default:"requires_email_code"`
	Username  string                               `json:"username" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ExpiresAt   respjson.Field
		Message     respjson.Field
		Object      respjson.Field
		Status      respjson.Field
		Username    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XAccountConnectionAttemptGetResponseRequiresEmailCode) RawJSON() string { return r.JSON.raw }
func (r *XAccountConnectionAttemptGetResponseRequiresEmailCode) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
