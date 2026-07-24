// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

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
	"github.com/Xquik-dev/x-twitter-scraper-go/shared/constant"
)

// X write actions (tweets, likes, follows, DMs)
//
// XCommunityJoinService contains methods and other services that help with
// interacting with the x-twitter-scraper API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewXCommunityJoinService] method instead.
type XCommunityJoinService struct {
	options []option.RequestOption
}

// NewXCommunityJoinService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewXCommunityJoinService(opts ...option.RequestOption) (r XCommunityJoinService) {
	r = XCommunityJoinService{}
	r.options = opts
	return
}

// Join community
func (r *XCommunityJoinService) New(ctx context.Context, id string, params XCommunityJoinNewParams, opts ...option.RequestOption) (res *XCommunityJoinNewResponse, err error) {
	if !param.IsOmitted(params.IdempotencyKey) {
		opts = append(opts, option.WithHeader("Idempotency-Key", fmt.Sprintf("%v", params.IdempotencyKey)))
	}
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("x/communities/%s/join", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Leave community
func (r *XCommunityJoinService) DeleteAll(ctx context.Context, id string, params XCommunityJoinDeleteAllParams, opts ...option.RequestOption) (res *XCommunityJoinDeleteAllResponse, err error) {
	if !param.IsOmitted(params.IdempotencyKey) {
		opts = append(opts, option.WithHeader("Idempotency-Key", fmt.Sprintf("%v", params.IdempotencyKey)))
	}
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("x/communities/%s/join", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, params, &res, opts...)
	return res, err
}

// Durable write lifecycle record. Poll statusUrl until terminal is true. Reusing
// the original Idempotency-Key returns this same record. Submit a new write only
// when safeToRetry is true, using a new key.
type XCommunityJoinNewResponse struct {
	ID string `json:"id" api:"required"`
	// Connected account selected for the write.
	Account XCommunityJoinNewResponseAccount `json:"account" api:"required"`
	// Any of "create_tweet", "delete_tweet", "like", "unlike", "retweet", "unretweet",
	// "follow", "unfollow", "remove_follower", "send_dm", "upload_media",
	// "update_profile", "update_avatar", "update_banner", "create_community",
	// "delete_community", "join_community", "leave_community".
	Action XCommunityJoinNewResponseAction `json:"action" api:"required"`
	// plannedCredits is the approved maximum. chargedCredits comes from the settled
	// credit ledger. Pending or failed writes are not charged.
	Billing        XCommunityJoinNewResponseBilling `json:"billing" api:"required"`
	Charged        bool                             `json:"charged" api:"required"`
	ChargedCredits string                           `json:"chargedCredits" api:"required"`
	// Exact follow-up an API client or agent should perform.
	NextAction  XCommunityJoinNewResponseNextAction `json:"nextAction" api:"required"`
	Object      constant.XWriteAction               `json:"object" default:"x_write_action"`
	PollAfterMs int64                               `json:"pollAfterMs" api:"required"`
	// Stable fingerprint and sanitized payload for replay checks.
	Request XCommunityJoinNewResponseRequest `json:"request" api:"required"`
	// Confirmed result produced by the write, when available.
	Result XCommunityJoinNewResponseResult `json:"result" api:"required"`
	// True only when a new attempt can reasonably succeed.
	Retryable bool `json:"retryable" api:"required"`
	// True only when no write was dispatched and a new idempotency key may be used.
	SafeToRetry    bool `json:"safeToRetry" api:"required"`
	SendDispatched bool `json:"sendDispatched" api:"required"`
	// Any of "accepted", "dispatching", "pending_confirmation", "success", "failed",
	// "expired".
	Status    XCommunityJoinNewResponseStatus `json:"status" api:"required"`
	StatusURL string                          `json:"statusUrl" api:"required"`
	Success   bool                            `json:"success" api:"required"`
	// Existing X resource targeted by the write, when applicable.
	Target        XCommunityJoinNewResponseTarget `json:"target" api:"required"`
	TargetID      string                          `json:"targetId" api:"required"`
	Terminal      bool                            `json:"terminal" api:"required"`
	WriteActionID string                          `json:"writeActionId" api:"required"`
	// Compatibility field for a confirmed community ID.
	CommunityID string `json:"communityId"`
	// Confirmed community name when available.
	CommunityName         string    `json:"communityName"`
	CompletedAt           time.Time `json:"completedAt" format:"date-time"`
	ConfirmationAttempts  int64     `json:"confirmationAttempts"`
	ConfirmationCheckedAt time.Time `json:"confirmationCheckedAt" format:"date-time"`
	ConfirmedAt           time.Time `json:"confirmedAt" format:"date-time"`
	CreatedAt             time.Time `json:"createdAt" format:"date-time"`
	// Structured recovery context for a failed write.
	Details map[string]any `json:"details"`
	Error   string         `json:"error"`
	// Deadline for resolving a non-terminal write. This is not the Idempotency-Key
	// retention deadline.
	ExpiresAt  time.Time `json:"expiresAt" format:"date-time"`
	Idempotent bool      `json:"idempotent"`
	// Media count, kind, size, and billing details when used.
	Media map[string]any `json:"media"`
	// Compatibility field for a confirmed media upload ID.
	MediaID string `json:"mediaId"`
	// Public media URL when the upload creates one.
	MediaURL string `json:"mediaUrl" format:"uri"`
	Message  string `json:"message"`
	// Compatibility field for a confirmed direct message ID.
	MessageID   string `json:"messageId"`
	RequestHash string `json:"requestHash"`
	RequestID   string `json:"requestId"`
	// Compatibility result ID for other write actions.
	ResultID string `json:"resultId"`
	// Dispatch timestamp when the write reached execution.
	SendDispatchedAt time.Time `json:"sendDispatchedAt" format:"date-time"`
	// Compatibility field for a confirmed tweet result ID.
	TweetID   string    `json:"tweetId"`
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                    respjson.Field
		Account               respjson.Field
		Action                respjson.Field
		Billing               respjson.Field
		Charged               respjson.Field
		ChargedCredits        respjson.Field
		NextAction            respjson.Field
		Object                respjson.Field
		PollAfterMs           respjson.Field
		Request               respjson.Field
		Result                respjson.Field
		Retryable             respjson.Field
		SafeToRetry           respjson.Field
		SendDispatched        respjson.Field
		Status                respjson.Field
		StatusURL             respjson.Field
		Success               respjson.Field
		Target                respjson.Field
		TargetID              respjson.Field
		Terminal              respjson.Field
		WriteActionID         respjson.Field
		CommunityID           respjson.Field
		CommunityName         respjson.Field
		CompletedAt           respjson.Field
		ConfirmationAttempts  respjson.Field
		ConfirmationCheckedAt respjson.Field
		ConfirmedAt           respjson.Field
		CreatedAt             respjson.Field
		Details               respjson.Field
		Error                 respjson.Field
		ExpiresAt             respjson.Field
		Idempotent            respjson.Field
		Media                 respjson.Field
		MediaID               respjson.Field
		MediaURL              respjson.Field
		Message               respjson.Field
		MessageID             respjson.Field
		RequestHash           respjson.Field
		RequestID             respjson.Field
		ResultID              respjson.Field
		SendDispatchedAt      respjson.Field
		TweetID               respjson.Field
		UpdatedAt             respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XCommunityJoinNewResponse) RawJSON() string { return r.JSON.raw }
func (r *XCommunityJoinNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Connected account selected for the write.
type XCommunityJoinNewResponseAccount struct {
	ID       string `json:"id" api:"required"`
	Username string `json:"username" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Username    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XCommunityJoinNewResponseAccount) RawJSON() string { return r.JSON.raw }
func (r *XCommunityJoinNewResponseAccount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type XCommunityJoinNewResponseAction string

const (
	XCommunityJoinNewResponseActionCreateTweet     XCommunityJoinNewResponseAction = "create_tweet"
	XCommunityJoinNewResponseActionDeleteTweet     XCommunityJoinNewResponseAction = "delete_tweet"
	XCommunityJoinNewResponseActionLike            XCommunityJoinNewResponseAction = "like"
	XCommunityJoinNewResponseActionUnlike          XCommunityJoinNewResponseAction = "unlike"
	XCommunityJoinNewResponseActionRetweet         XCommunityJoinNewResponseAction = "retweet"
	XCommunityJoinNewResponseActionUnretweet       XCommunityJoinNewResponseAction = "unretweet"
	XCommunityJoinNewResponseActionFollow          XCommunityJoinNewResponseAction = "follow"
	XCommunityJoinNewResponseActionUnfollow        XCommunityJoinNewResponseAction = "unfollow"
	XCommunityJoinNewResponseActionRemoveFollower  XCommunityJoinNewResponseAction = "remove_follower"
	XCommunityJoinNewResponseActionSendDm          XCommunityJoinNewResponseAction = "send_dm"
	XCommunityJoinNewResponseActionUploadMedia     XCommunityJoinNewResponseAction = "upload_media"
	XCommunityJoinNewResponseActionUpdateProfile   XCommunityJoinNewResponseAction = "update_profile"
	XCommunityJoinNewResponseActionUpdateAvatar    XCommunityJoinNewResponseAction = "update_avatar"
	XCommunityJoinNewResponseActionUpdateBanner    XCommunityJoinNewResponseAction = "update_banner"
	XCommunityJoinNewResponseActionCreateCommunity XCommunityJoinNewResponseAction = "create_community"
	XCommunityJoinNewResponseActionDeleteCommunity XCommunityJoinNewResponseAction = "delete_community"
	XCommunityJoinNewResponseActionJoinCommunity   XCommunityJoinNewResponseAction = "join_community"
	XCommunityJoinNewResponseActionLeaveCommunity  XCommunityJoinNewResponseAction = "leave_community"
)

// plannedCredits is the approved maximum. chargedCredits comes from the settled
// credit ledger. Pending or failed writes are not charged.
type XCommunityJoinNewResponseBilling struct {
	Charged        bool   `json:"charged" api:"required"`
	ChargedCredits string `json:"chargedCredits" api:"required"`
	PlannedCredits string `json:"plannedCredits" api:"required"`
	// Any of "not_charged", "pending", "charged", "charge_failed", "refunded".
	Status string `json:"status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Charged        respjson.Field
		ChargedCredits respjson.Field
		PlannedCredits respjson.Field
		Status         respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XCommunityJoinNewResponseBilling) RawJSON() string { return r.JSON.raw }
func (r *XCommunityJoinNewResponseBilling) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Exact follow-up an API client or agent should perform.
type XCommunityJoinNewResponseNextAction struct {
	// Any of "poll", "retry", "verify_result", "fix_request".
	Type                      string `json:"type" api:"required"`
	AfterMs                   int64  `json:"afterMs"`
	RequiresNewIdempotencyKey bool   `json:"requiresNewIdempotencyKey"`
	URL                       string `json:"url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type                      respjson.Field
		AfterMs                   respjson.Field
		RequiresNewIdempotencyKey respjson.Field
		URL                       respjson.Field
		ExtraFields               map[string]respjson.Field
		raw                       string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XCommunityJoinNewResponseNextAction) RawJSON() string { return r.JSON.raw }
func (r *XCommunityJoinNewResponseNextAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Stable fingerprint and sanitized payload for replay checks.
type XCommunityJoinNewResponseRequest struct {
	// Stable hash of account, action, target, and payload.
	Hash string `json:"hash" api:"required"`
	// Exact sanitized payload dispatched for this action.
	Payload map[string]any `json:"payload" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Hash        respjson.Field
		Payload     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XCommunityJoinNewResponseRequest) RawJSON() string { return r.JSON.raw }
func (r *XCommunityJoinNewResponseRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Confirmed result produced by the write, when available.
type XCommunityJoinNewResponseResult struct {
	ID    string `json:"id"`
	State string `json:"state"`
	// Any of "tweet", "direct_message", "media", "community", "state_change".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		State       respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XCommunityJoinNewResponseResult) RawJSON() string { return r.JSON.raw }
func (r *XCommunityJoinNewResponseResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type XCommunityJoinNewResponseStatus string

const (
	XCommunityJoinNewResponseStatusAccepted            XCommunityJoinNewResponseStatus = "accepted"
	XCommunityJoinNewResponseStatusDispatching         XCommunityJoinNewResponseStatus = "dispatching"
	XCommunityJoinNewResponseStatusPendingConfirmation XCommunityJoinNewResponseStatus = "pending_confirmation"
	XCommunityJoinNewResponseStatusSuccess             XCommunityJoinNewResponseStatus = "success"
	XCommunityJoinNewResponseStatusFailed              XCommunityJoinNewResponseStatus = "failed"
	XCommunityJoinNewResponseStatusExpired             XCommunityJoinNewResponseStatus = "expired"
)

// Existing X resource targeted by the write, when applicable.
type XCommunityJoinNewResponseTarget struct {
	ID string `json:"id" api:"required"`
	// Any of "tweet", "user", "community".
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XCommunityJoinNewResponseTarget) RawJSON() string { return r.JSON.raw }
func (r *XCommunityJoinNewResponseTarget) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Durable write lifecycle record. Poll statusUrl until terminal is true. Reusing
// the original Idempotency-Key returns this same record. Submit a new write only
// when safeToRetry is true, using a new key.
type XCommunityJoinDeleteAllResponse struct {
	ID string `json:"id" api:"required"`
	// Connected account selected for the write.
	Account XCommunityJoinDeleteAllResponseAccount `json:"account" api:"required"`
	// Any of "create_tweet", "delete_tweet", "like", "unlike", "retweet", "unretweet",
	// "follow", "unfollow", "remove_follower", "send_dm", "upload_media",
	// "update_profile", "update_avatar", "update_banner", "create_community",
	// "delete_community", "join_community", "leave_community".
	Action XCommunityJoinDeleteAllResponseAction `json:"action" api:"required"`
	// plannedCredits is the approved maximum. chargedCredits comes from the settled
	// credit ledger. Pending or failed writes are not charged.
	Billing        XCommunityJoinDeleteAllResponseBilling `json:"billing" api:"required"`
	Charged        bool                                   `json:"charged" api:"required"`
	ChargedCredits string                                 `json:"chargedCredits" api:"required"`
	// Exact follow-up an API client or agent should perform.
	NextAction  XCommunityJoinDeleteAllResponseNextAction `json:"nextAction" api:"required"`
	Object      constant.XWriteAction                     `json:"object" default:"x_write_action"`
	PollAfterMs int64                                     `json:"pollAfterMs" api:"required"`
	// Stable fingerprint and sanitized payload for replay checks.
	Request XCommunityJoinDeleteAllResponseRequest `json:"request" api:"required"`
	// Confirmed result produced by the write, when available.
	Result XCommunityJoinDeleteAllResponseResult `json:"result" api:"required"`
	// True only when a new attempt can reasonably succeed.
	Retryable bool `json:"retryable" api:"required"`
	// True only when no write was dispatched and a new idempotency key may be used.
	SafeToRetry    bool `json:"safeToRetry" api:"required"`
	SendDispatched bool `json:"sendDispatched" api:"required"`
	// Any of "accepted", "dispatching", "pending_confirmation", "success", "failed",
	// "expired".
	Status    XCommunityJoinDeleteAllResponseStatus `json:"status" api:"required"`
	StatusURL string                                `json:"statusUrl" api:"required"`
	Success   bool                                  `json:"success" api:"required"`
	// Existing X resource targeted by the write, when applicable.
	Target        XCommunityJoinDeleteAllResponseTarget `json:"target" api:"required"`
	TargetID      string                                `json:"targetId" api:"required"`
	Terminal      bool                                  `json:"terminal" api:"required"`
	WriteActionID string                                `json:"writeActionId" api:"required"`
	// Compatibility field for a confirmed community ID.
	CommunityID string `json:"communityId"`
	// Confirmed community name when available.
	CommunityName         string    `json:"communityName"`
	CompletedAt           time.Time `json:"completedAt" format:"date-time"`
	ConfirmationAttempts  int64     `json:"confirmationAttempts"`
	ConfirmationCheckedAt time.Time `json:"confirmationCheckedAt" format:"date-time"`
	ConfirmedAt           time.Time `json:"confirmedAt" format:"date-time"`
	CreatedAt             time.Time `json:"createdAt" format:"date-time"`
	// Structured recovery context for a failed write.
	Details map[string]any `json:"details"`
	Error   string         `json:"error"`
	// Deadline for resolving a non-terminal write. This is not the Idempotency-Key
	// retention deadline.
	ExpiresAt  time.Time `json:"expiresAt" format:"date-time"`
	Idempotent bool      `json:"idempotent"`
	// Media count, kind, size, and billing details when used.
	Media map[string]any `json:"media"`
	// Compatibility field for a confirmed media upload ID.
	MediaID string `json:"mediaId"`
	// Public media URL when the upload creates one.
	MediaURL string `json:"mediaUrl" format:"uri"`
	Message  string `json:"message"`
	// Compatibility field for a confirmed direct message ID.
	MessageID   string `json:"messageId"`
	RequestHash string `json:"requestHash"`
	RequestID   string `json:"requestId"`
	// Compatibility result ID for other write actions.
	ResultID string `json:"resultId"`
	// Dispatch timestamp when the write reached execution.
	SendDispatchedAt time.Time `json:"sendDispatchedAt" format:"date-time"`
	// Compatibility field for a confirmed tweet result ID.
	TweetID   string    `json:"tweetId"`
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                    respjson.Field
		Account               respjson.Field
		Action                respjson.Field
		Billing               respjson.Field
		Charged               respjson.Field
		ChargedCredits        respjson.Field
		NextAction            respjson.Field
		Object                respjson.Field
		PollAfterMs           respjson.Field
		Request               respjson.Field
		Result                respjson.Field
		Retryable             respjson.Field
		SafeToRetry           respjson.Field
		SendDispatched        respjson.Field
		Status                respjson.Field
		StatusURL             respjson.Field
		Success               respjson.Field
		Target                respjson.Field
		TargetID              respjson.Field
		Terminal              respjson.Field
		WriteActionID         respjson.Field
		CommunityID           respjson.Field
		CommunityName         respjson.Field
		CompletedAt           respjson.Field
		ConfirmationAttempts  respjson.Field
		ConfirmationCheckedAt respjson.Field
		ConfirmedAt           respjson.Field
		CreatedAt             respjson.Field
		Details               respjson.Field
		Error                 respjson.Field
		ExpiresAt             respjson.Field
		Idempotent            respjson.Field
		Media                 respjson.Field
		MediaID               respjson.Field
		MediaURL              respjson.Field
		Message               respjson.Field
		MessageID             respjson.Field
		RequestHash           respjson.Field
		RequestID             respjson.Field
		ResultID              respjson.Field
		SendDispatchedAt      respjson.Field
		TweetID               respjson.Field
		UpdatedAt             respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XCommunityJoinDeleteAllResponse) RawJSON() string { return r.JSON.raw }
func (r *XCommunityJoinDeleteAllResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Connected account selected for the write.
type XCommunityJoinDeleteAllResponseAccount struct {
	ID       string `json:"id" api:"required"`
	Username string `json:"username" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Username    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XCommunityJoinDeleteAllResponseAccount) RawJSON() string { return r.JSON.raw }
func (r *XCommunityJoinDeleteAllResponseAccount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type XCommunityJoinDeleteAllResponseAction string

const (
	XCommunityJoinDeleteAllResponseActionCreateTweet     XCommunityJoinDeleteAllResponseAction = "create_tweet"
	XCommunityJoinDeleteAllResponseActionDeleteTweet     XCommunityJoinDeleteAllResponseAction = "delete_tweet"
	XCommunityJoinDeleteAllResponseActionLike            XCommunityJoinDeleteAllResponseAction = "like"
	XCommunityJoinDeleteAllResponseActionUnlike          XCommunityJoinDeleteAllResponseAction = "unlike"
	XCommunityJoinDeleteAllResponseActionRetweet         XCommunityJoinDeleteAllResponseAction = "retweet"
	XCommunityJoinDeleteAllResponseActionUnretweet       XCommunityJoinDeleteAllResponseAction = "unretweet"
	XCommunityJoinDeleteAllResponseActionFollow          XCommunityJoinDeleteAllResponseAction = "follow"
	XCommunityJoinDeleteAllResponseActionUnfollow        XCommunityJoinDeleteAllResponseAction = "unfollow"
	XCommunityJoinDeleteAllResponseActionRemoveFollower  XCommunityJoinDeleteAllResponseAction = "remove_follower"
	XCommunityJoinDeleteAllResponseActionSendDm          XCommunityJoinDeleteAllResponseAction = "send_dm"
	XCommunityJoinDeleteAllResponseActionUploadMedia     XCommunityJoinDeleteAllResponseAction = "upload_media"
	XCommunityJoinDeleteAllResponseActionUpdateProfile   XCommunityJoinDeleteAllResponseAction = "update_profile"
	XCommunityJoinDeleteAllResponseActionUpdateAvatar    XCommunityJoinDeleteAllResponseAction = "update_avatar"
	XCommunityJoinDeleteAllResponseActionUpdateBanner    XCommunityJoinDeleteAllResponseAction = "update_banner"
	XCommunityJoinDeleteAllResponseActionCreateCommunity XCommunityJoinDeleteAllResponseAction = "create_community"
	XCommunityJoinDeleteAllResponseActionDeleteCommunity XCommunityJoinDeleteAllResponseAction = "delete_community"
	XCommunityJoinDeleteAllResponseActionJoinCommunity   XCommunityJoinDeleteAllResponseAction = "join_community"
	XCommunityJoinDeleteAllResponseActionLeaveCommunity  XCommunityJoinDeleteAllResponseAction = "leave_community"
)

// plannedCredits is the approved maximum. chargedCredits comes from the settled
// credit ledger. Pending or failed writes are not charged.
type XCommunityJoinDeleteAllResponseBilling struct {
	Charged        bool   `json:"charged" api:"required"`
	ChargedCredits string `json:"chargedCredits" api:"required"`
	PlannedCredits string `json:"plannedCredits" api:"required"`
	// Any of "not_charged", "pending", "charged", "charge_failed", "refunded".
	Status string `json:"status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Charged        respjson.Field
		ChargedCredits respjson.Field
		PlannedCredits respjson.Field
		Status         respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XCommunityJoinDeleteAllResponseBilling) RawJSON() string { return r.JSON.raw }
func (r *XCommunityJoinDeleteAllResponseBilling) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Exact follow-up an API client or agent should perform.
type XCommunityJoinDeleteAllResponseNextAction struct {
	// Any of "poll", "retry", "verify_result", "fix_request".
	Type                      string `json:"type" api:"required"`
	AfterMs                   int64  `json:"afterMs"`
	RequiresNewIdempotencyKey bool   `json:"requiresNewIdempotencyKey"`
	URL                       string `json:"url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type                      respjson.Field
		AfterMs                   respjson.Field
		RequiresNewIdempotencyKey respjson.Field
		URL                       respjson.Field
		ExtraFields               map[string]respjson.Field
		raw                       string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XCommunityJoinDeleteAllResponseNextAction) RawJSON() string { return r.JSON.raw }
func (r *XCommunityJoinDeleteAllResponseNextAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Stable fingerprint and sanitized payload for replay checks.
type XCommunityJoinDeleteAllResponseRequest struct {
	// Stable hash of account, action, target, and payload.
	Hash string `json:"hash" api:"required"`
	// Exact sanitized payload dispatched for this action.
	Payload map[string]any `json:"payload" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Hash        respjson.Field
		Payload     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XCommunityJoinDeleteAllResponseRequest) RawJSON() string { return r.JSON.raw }
func (r *XCommunityJoinDeleteAllResponseRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Confirmed result produced by the write, when available.
type XCommunityJoinDeleteAllResponseResult struct {
	ID    string `json:"id"`
	State string `json:"state"`
	// Any of "tweet", "direct_message", "media", "community", "state_change".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		State       respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XCommunityJoinDeleteAllResponseResult) RawJSON() string { return r.JSON.raw }
func (r *XCommunityJoinDeleteAllResponseResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type XCommunityJoinDeleteAllResponseStatus string

const (
	XCommunityJoinDeleteAllResponseStatusAccepted            XCommunityJoinDeleteAllResponseStatus = "accepted"
	XCommunityJoinDeleteAllResponseStatusDispatching         XCommunityJoinDeleteAllResponseStatus = "dispatching"
	XCommunityJoinDeleteAllResponseStatusPendingConfirmation XCommunityJoinDeleteAllResponseStatus = "pending_confirmation"
	XCommunityJoinDeleteAllResponseStatusSuccess             XCommunityJoinDeleteAllResponseStatus = "success"
	XCommunityJoinDeleteAllResponseStatusFailed              XCommunityJoinDeleteAllResponseStatus = "failed"
	XCommunityJoinDeleteAllResponseStatusExpired             XCommunityJoinDeleteAllResponseStatus = "expired"
)

// Existing X resource targeted by the write, when applicable.
type XCommunityJoinDeleteAllResponseTarget struct {
	ID string `json:"id" api:"required"`
	// Any of "tweet", "user", "community".
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XCommunityJoinDeleteAllResponseTarget) RawJSON() string { return r.JSON.raw }
func (r *XCommunityJoinDeleteAllResponseTarget) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type XCommunityJoinNewParams struct {
	// X account identifier (@username or account ID)
	Account        string `json:"account" api:"required"`
	IdempotencyKey string `header:"Idempotency-Key" api:"required" json:"-"`
	paramObj
}

func (r XCommunityJoinNewParams) MarshalJSON() (data []byte, err error) {
	type shadow XCommunityJoinNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *XCommunityJoinNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type XCommunityJoinDeleteAllParams struct {
	// X account identifier (@username or account ID)
	Account        string `json:"account" api:"required"`
	IdempotencyKey string `header:"Idempotency-Key" api:"required" json:"-"`
	paramObj
}

func (r XCommunityJoinDeleteAllParams) MarshalJSON() (data []byte, err error) {
	type shadow XCommunityJoinDeleteAllParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *XCommunityJoinDeleteAllParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
