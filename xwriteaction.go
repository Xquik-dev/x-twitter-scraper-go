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
	"github.com/Xquik-dev/x-twitter-scraper-go/packages/respjson"
	"github.com/Xquik-dev/x-twitter-scraper-go/shared/constant"
)

// X write actions (tweets, likes, follows, DMs)
//
// XWriteActionService contains methods and other services that help with
// interacting with the x-twitter-scraper API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewXWriteActionService] method instead.
type XWriteActionService struct {
	options []option.RequestOption
}

// NewXWriteActionService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewXWriteActionService(opts ...option.RequestOption) (r XWriteActionService) {
	r = XWriteActionService{}
	r.options = opts
	return
}

// Get write action status
func (r *XWriteActionService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *XWriteActionGetResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("x/write-actions/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Durable write lifecycle record. Poll statusUrl until terminal is true. Reusing
// the original Idempotency-Key returns this same record. Submit a new write only
// when safeToRetry is true, using a new key.
type XWriteActionGetResponse struct {
	ID string `json:"id" api:"required"`
	// Connected account selected for the write.
	Account XWriteActionGetResponseAccount `json:"account" api:"required"`
	// Any of "create_tweet", "delete_tweet", "like", "unlike", "retweet", "unretweet",
	// "follow", "unfollow", "remove_follower", "send_dm", "upload_media",
	// "update_profile", "update_avatar", "update_banner", "create_community",
	// "delete_community", "join_community", "leave_community".
	Action XWriteActionGetResponseAction `json:"action" api:"required"`
	// plannedCredits is the approved maximum. chargedCredits comes from the settled
	// credit ledger. Pending or failed writes are not charged.
	Billing        XWriteActionGetResponseBilling `json:"billing" api:"required"`
	Charged        bool                           `json:"charged" api:"required"`
	ChargedCredits string                         `json:"chargedCredits" api:"required"`
	// Exact follow-up an API client or agent should perform.
	NextAction  XWriteActionGetResponseNextAction `json:"nextAction" api:"required"`
	Object      constant.XWriteAction             `json:"object" default:"x_write_action"`
	PollAfterMs int64                             `json:"pollAfterMs" api:"required"`
	// Stable fingerprint and sanitized payload for replay checks.
	Request XWriteActionGetResponseRequest `json:"request" api:"required"`
	// Confirmed result produced by the write, when available.
	Result XWriteActionGetResponseResult `json:"result" api:"required"`
	// True only when a new attempt can reasonably succeed.
	Retryable bool `json:"retryable" api:"required"`
	// True only when no write was dispatched and a new idempotency key may be used.
	SafeToRetry    bool `json:"safeToRetry" api:"required"`
	SendDispatched bool `json:"sendDispatched" api:"required"`
	// Any of "accepted", "dispatching", "pending_confirmation", "success", "failed",
	// "expired".
	Status    XWriteActionGetResponseStatus `json:"status" api:"required"`
	StatusURL string                        `json:"statusUrl" api:"required"`
	Success   bool                          `json:"success" api:"required"`
	// Existing X resource targeted by the write, when applicable.
	Target        XWriteActionGetResponseTarget `json:"target" api:"required"`
	TargetID      string                        `json:"targetId" api:"required"`
	Terminal      bool                          `json:"terminal" api:"required"`
	WriteActionID string                        `json:"writeActionId" api:"required"`
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
func (r XWriteActionGetResponse) RawJSON() string { return r.JSON.raw }
func (r *XWriteActionGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Connected account selected for the write.
type XWriteActionGetResponseAccount struct {
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
func (r XWriteActionGetResponseAccount) RawJSON() string { return r.JSON.raw }
func (r *XWriteActionGetResponseAccount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type XWriteActionGetResponseAction string

const (
	XWriteActionGetResponseActionCreateTweet     XWriteActionGetResponseAction = "create_tweet"
	XWriteActionGetResponseActionDeleteTweet     XWriteActionGetResponseAction = "delete_tweet"
	XWriteActionGetResponseActionLike            XWriteActionGetResponseAction = "like"
	XWriteActionGetResponseActionUnlike          XWriteActionGetResponseAction = "unlike"
	XWriteActionGetResponseActionRetweet         XWriteActionGetResponseAction = "retweet"
	XWriteActionGetResponseActionUnretweet       XWriteActionGetResponseAction = "unretweet"
	XWriteActionGetResponseActionFollow          XWriteActionGetResponseAction = "follow"
	XWriteActionGetResponseActionUnfollow        XWriteActionGetResponseAction = "unfollow"
	XWriteActionGetResponseActionRemoveFollower  XWriteActionGetResponseAction = "remove_follower"
	XWriteActionGetResponseActionSendDm          XWriteActionGetResponseAction = "send_dm"
	XWriteActionGetResponseActionUploadMedia     XWriteActionGetResponseAction = "upload_media"
	XWriteActionGetResponseActionUpdateProfile   XWriteActionGetResponseAction = "update_profile"
	XWriteActionGetResponseActionUpdateAvatar    XWriteActionGetResponseAction = "update_avatar"
	XWriteActionGetResponseActionUpdateBanner    XWriteActionGetResponseAction = "update_banner"
	XWriteActionGetResponseActionCreateCommunity XWriteActionGetResponseAction = "create_community"
	XWriteActionGetResponseActionDeleteCommunity XWriteActionGetResponseAction = "delete_community"
	XWriteActionGetResponseActionJoinCommunity   XWriteActionGetResponseAction = "join_community"
	XWriteActionGetResponseActionLeaveCommunity  XWriteActionGetResponseAction = "leave_community"
)

// plannedCredits is the approved maximum. chargedCredits comes from the settled
// credit ledger. Pending or failed writes are not charged.
type XWriteActionGetResponseBilling struct {
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
func (r XWriteActionGetResponseBilling) RawJSON() string { return r.JSON.raw }
func (r *XWriteActionGetResponseBilling) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Exact follow-up an API client or agent should perform.
type XWriteActionGetResponseNextAction struct {
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
func (r XWriteActionGetResponseNextAction) RawJSON() string { return r.JSON.raw }
func (r *XWriteActionGetResponseNextAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Stable fingerprint and sanitized payload for replay checks.
type XWriteActionGetResponseRequest struct {
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
func (r XWriteActionGetResponseRequest) RawJSON() string { return r.JSON.raw }
func (r *XWriteActionGetResponseRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Confirmed result produced by the write, when available.
type XWriteActionGetResponseResult struct {
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
func (r XWriteActionGetResponseResult) RawJSON() string { return r.JSON.raw }
func (r *XWriteActionGetResponseResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type XWriteActionGetResponseStatus string

const (
	XWriteActionGetResponseStatusAccepted            XWriteActionGetResponseStatus = "accepted"
	XWriteActionGetResponseStatusDispatching         XWriteActionGetResponseStatus = "dispatching"
	XWriteActionGetResponseStatusPendingConfirmation XWriteActionGetResponseStatus = "pending_confirmation"
	XWriteActionGetResponseStatusSuccess             XWriteActionGetResponseStatus = "success"
	XWriteActionGetResponseStatusFailed              XWriteActionGetResponseStatus = "failed"
	XWriteActionGetResponseStatusExpired             XWriteActionGetResponseStatus = "expired"
)

// Existing X resource targeted by the write, when applicable.
type XWriteActionGetResponseTarget struct {
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
func (r XWriteActionGetResponseTarget) RawJSON() string { return r.JSON.raw }
func (r *XWriteActionGetResponseTarget) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
