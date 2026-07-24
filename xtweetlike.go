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
// XTweetLikeService contains methods and other services that help with interacting
// with the x-twitter-scraper API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewXTweetLikeService] method instead.
type XTweetLikeService struct {
	options []option.RequestOption
}

// NewXTweetLikeService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewXTweetLikeService(opts ...option.RequestOption) (r XTweetLikeService) {
	r = XTweetLikeService{}
	r.options = opts
	return
}

// Like tweet
func (r *XTweetLikeService) New(ctx context.Context, id string, params XTweetLikeNewParams, opts ...option.RequestOption) (res *XTweetLikeNewResponse, err error) {
	if !param.IsOmitted(params.IdempotencyKey) {
		opts = append(opts, option.WithHeader("Idempotency-Key", fmt.Sprintf("%v", params.IdempotencyKey)))
	}
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("x/tweets/%s/like", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Unlike tweet
func (r *XTweetLikeService) Delete(ctx context.Context, id string, params XTweetLikeDeleteParams, opts ...option.RequestOption) (res *XTweetLikeDeleteResponse, err error) {
	if !param.IsOmitted(params.IdempotencyKey) {
		opts = append(opts, option.WithHeader("Idempotency-Key", fmt.Sprintf("%v", params.IdempotencyKey)))
	}
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("x/tweets/%s/like", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, params, &res, opts...)
	return res, err
}

// Durable write lifecycle record. Poll statusUrl until terminal is true. Reusing
// the original Idempotency-Key returns this same record. Submit a new write only
// when safeToRetry is true, using a new key.
type XTweetLikeNewResponse struct {
	ID string `json:"id" api:"required"`
	// Connected account selected for the write.
	Account XTweetLikeNewResponseAccount `json:"account" api:"required"`
	// Any of "create_tweet", "delete_tweet", "like", "unlike", "retweet", "unretweet",
	// "follow", "unfollow", "remove_follower", "send_dm", "upload_media",
	// "update_profile", "update_avatar", "update_banner", "create_community",
	// "delete_community", "join_community", "leave_community".
	Action XTweetLikeNewResponseAction `json:"action" api:"required"`
	// plannedCredits is the approved maximum. chargedCredits comes from the settled
	// credit ledger. Pending or failed writes are not charged.
	Billing        XTweetLikeNewResponseBilling `json:"billing" api:"required"`
	Charged        bool                         `json:"charged" api:"required"`
	ChargedCredits string                       `json:"chargedCredits" api:"required"`
	// Exact follow-up an API client or agent should perform.
	NextAction  XTweetLikeNewResponseNextAction `json:"nextAction" api:"required"`
	Object      constant.XWriteAction           `json:"object" default:"x_write_action"`
	PollAfterMs int64                           `json:"pollAfterMs" api:"required"`
	// Stable fingerprint and sanitized payload for replay checks.
	Request XTweetLikeNewResponseRequest `json:"request" api:"required"`
	// Confirmed result produced by the write, when available.
	Result XTweetLikeNewResponseResult `json:"result" api:"required"`
	// True only when a new attempt can reasonably succeed.
	Retryable bool `json:"retryable" api:"required"`
	// True only when no write was dispatched and a new idempotency key may be used.
	SafeToRetry    bool `json:"safeToRetry" api:"required"`
	SendDispatched bool `json:"sendDispatched" api:"required"`
	// Any of "accepted", "dispatching", "pending_confirmation", "success", "failed",
	// "expired".
	Status    XTweetLikeNewResponseStatus `json:"status" api:"required"`
	StatusURL string                      `json:"statusUrl" api:"required"`
	Success   bool                        `json:"success" api:"required"`
	// Existing X resource targeted by the write, when applicable.
	Target        XTweetLikeNewResponseTarget `json:"target" api:"required"`
	TargetID      string                      `json:"targetId" api:"required"`
	Terminal      bool                        `json:"terminal" api:"required"`
	WriteActionID string                      `json:"writeActionId" api:"required"`
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
func (r XTweetLikeNewResponse) RawJSON() string { return r.JSON.raw }
func (r *XTweetLikeNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Connected account selected for the write.
type XTweetLikeNewResponseAccount struct {
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
func (r XTweetLikeNewResponseAccount) RawJSON() string { return r.JSON.raw }
func (r *XTweetLikeNewResponseAccount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type XTweetLikeNewResponseAction string

const (
	XTweetLikeNewResponseActionCreateTweet     XTweetLikeNewResponseAction = "create_tweet"
	XTweetLikeNewResponseActionDeleteTweet     XTweetLikeNewResponseAction = "delete_tweet"
	XTweetLikeNewResponseActionLike            XTweetLikeNewResponseAction = "like"
	XTweetLikeNewResponseActionUnlike          XTweetLikeNewResponseAction = "unlike"
	XTweetLikeNewResponseActionRetweet         XTweetLikeNewResponseAction = "retweet"
	XTweetLikeNewResponseActionUnretweet       XTweetLikeNewResponseAction = "unretweet"
	XTweetLikeNewResponseActionFollow          XTweetLikeNewResponseAction = "follow"
	XTweetLikeNewResponseActionUnfollow        XTweetLikeNewResponseAction = "unfollow"
	XTweetLikeNewResponseActionRemoveFollower  XTweetLikeNewResponseAction = "remove_follower"
	XTweetLikeNewResponseActionSendDm          XTweetLikeNewResponseAction = "send_dm"
	XTweetLikeNewResponseActionUploadMedia     XTweetLikeNewResponseAction = "upload_media"
	XTweetLikeNewResponseActionUpdateProfile   XTweetLikeNewResponseAction = "update_profile"
	XTweetLikeNewResponseActionUpdateAvatar    XTweetLikeNewResponseAction = "update_avatar"
	XTweetLikeNewResponseActionUpdateBanner    XTweetLikeNewResponseAction = "update_banner"
	XTweetLikeNewResponseActionCreateCommunity XTweetLikeNewResponseAction = "create_community"
	XTweetLikeNewResponseActionDeleteCommunity XTweetLikeNewResponseAction = "delete_community"
	XTweetLikeNewResponseActionJoinCommunity   XTweetLikeNewResponseAction = "join_community"
	XTweetLikeNewResponseActionLeaveCommunity  XTweetLikeNewResponseAction = "leave_community"
)

// plannedCredits is the approved maximum. chargedCredits comes from the settled
// credit ledger. Pending or failed writes are not charged.
type XTweetLikeNewResponseBilling struct {
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
func (r XTweetLikeNewResponseBilling) RawJSON() string { return r.JSON.raw }
func (r *XTweetLikeNewResponseBilling) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Exact follow-up an API client or agent should perform.
type XTweetLikeNewResponseNextAction struct {
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
func (r XTweetLikeNewResponseNextAction) RawJSON() string { return r.JSON.raw }
func (r *XTweetLikeNewResponseNextAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Stable fingerprint and sanitized payload for replay checks.
type XTweetLikeNewResponseRequest struct {
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
func (r XTweetLikeNewResponseRequest) RawJSON() string { return r.JSON.raw }
func (r *XTweetLikeNewResponseRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Confirmed result produced by the write, when available.
type XTweetLikeNewResponseResult struct {
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
func (r XTweetLikeNewResponseResult) RawJSON() string { return r.JSON.raw }
func (r *XTweetLikeNewResponseResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type XTweetLikeNewResponseStatus string

const (
	XTweetLikeNewResponseStatusAccepted            XTweetLikeNewResponseStatus = "accepted"
	XTweetLikeNewResponseStatusDispatching         XTweetLikeNewResponseStatus = "dispatching"
	XTweetLikeNewResponseStatusPendingConfirmation XTweetLikeNewResponseStatus = "pending_confirmation"
	XTweetLikeNewResponseStatusSuccess             XTweetLikeNewResponseStatus = "success"
	XTweetLikeNewResponseStatusFailed              XTweetLikeNewResponseStatus = "failed"
	XTweetLikeNewResponseStatusExpired             XTweetLikeNewResponseStatus = "expired"
)

// Existing X resource targeted by the write, when applicable.
type XTweetLikeNewResponseTarget struct {
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
func (r XTweetLikeNewResponseTarget) RawJSON() string { return r.JSON.raw }
func (r *XTweetLikeNewResponseTarget) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Durable write lifecycle record. Poll statusUrl until terminal is true. Reusing
// the original Idempotency-Key returns this same record. Submit a new write only
// when safeToRetry is true, using a new key.
type XTweetLikeDeleteResponse struct {
	ID string `json:"id" api:"required"`
	// Connected account selected for the write.
	Account XTweetLikeDeleteResponseAccount `json:"account" api:"required"`
	// Any of "create_tweet", "delete_tweet", "like", "unlike", "retweet", "unretweet",
	// "follow", "unfollow", "remove_follower", "send_dm", "upload_media",
	// "update_profile", "update_avatar", "update_banner", "create_community",
	// "delete_community", "join_community", "leave_community".
	Action XTweetLikeDeleteResponseAction `json:"action" api:"required"`
	// plannedCredits is the approved maximum. chargedCredits comes from the settled
	// credit ledger. Pending or failed writes are not charged.
	Billing        XTweetLikeDeleteResponseBilling `json:"billing" api:"required"`
	Charged        bool                            `json:"charged" api:"required"`
	ChargedCredits string                          `json:"chargedCredits" api:"required"`
	// Exact follow-up an API client or agent should perform.
	NextAction  XTweetLikeDeleteResponseNextAction `json:"nextAction" api:"required"`
	Object      constant.XWriteAction              `json:"object" default:"x_write_action"`
	PollAfterMs int64                              `json:"pollAfterMs" api:"required"`
	// Stable fingerprint and sanitized payload for replay checks.
	Request XTweetLikeDeleteResponseRequest `json:"request" api:"required"`
	// Confirmed result produced by the write, when available.
	Result XTweetLikeDeleteResponseResult `json:"result" api:"required"`
	// True only when a new attempt can reasonably succeed.
	Retryable bool `json:"retryable" api:"required"`
	// True only when no write was dispatched and a new idempotency key may be used.
	SafeToRetry    bool `json:"safeToRetry" api:"required"`
	SendDispatched bool `json:"sendDispatched" api:"required"`
	// Any of "accepted", "dispatching", "pending_confirmation", "success", "failed",
	// "expired".
	Status    XTweetLikeDeleteResponseStatus `json:"status" api:"required"`
	StatusURL string                         `json:"statusUrl" api:"required"`
	Success   bool                           `json:"success" api:"required"`
	// Existing X resource targeted by the write, when applicable.
	Target        XTweetLikeDeleteResponseTarget `json:"target" api:"required"`
	TargetID      string                         `json:"targetId" api:"required"`
	Terminal      bool                           `json:"terminal" api:"required"`
	WriteActionID string                         `json:"writeActionId" api:"required"`
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
func (r XTweetLikeDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *XTweetLikeDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Connected account selected for the write.
type XTweetLikeDeleteResponseAccount struct {
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
func (r XTweetLikeDeleteResponseAccount) RawJSON() string { return r.JSON.raw }
func (r *XTweetLikeDeleteResponseAccount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type XTweetLikeDeleteResponseAction string

const (
	XTweetLikeDeleteResponseActionCreateTweet     XTweetLikeDeleteResponseAction = "create_tweet"
	XTweetLikeDeleteResponseActionDeleteTweet     XTweetLikeDeleteResponseAction = "delete_tweet"
	XTweetLikeDeleteResponseActionLike            XTweetLikeDeleteResponseAction = "like"
	XTweetLikeDeleteResponseActionUnlike          XTweetLikeDeleteResponseAction = "unlike"
	XTweetLikeDeleteResponseActionRetweet         XTweetLikeDeleteResponseAction = "retweet"
	XTweetLikeDeleteResponseActionUnretweet       XTweetLikeDeleteResponseAction = "unretweet"
	XTweetLikeDeleteResponseActionFollow          XTweetLikeDeleteResponseAction = "follow"
	XTweetLikeDeleteResponseActionUnfollow        XTweetLikeDeleteResponseAction = "unfollow"
	XTweetLikeDeleteResponseActionRemoveFollower  XTweetLikeDeleteResponseAction = "remove_follower"
	XTweetLikeDeleteResponseActionSendDm          XTweetLikeDeleteResponseAction = "send_dm"
	XTweetLikeDeleteResponseActionUploadMedia     XTweetLikeDeleteResponseAction = "upload_media"
	XTweetLikeDeleteResponseActionUpdateProfile   XTweetLikeDeleteResponseAction = "update_profile"
	XTweetLikeDeleteResponseActionUpdateAvatar    XTweetLikeDeleteResponseAction = "update_avatar"
	XTweetLikeDeleteResponseActionUpdateBanner    XTweetLikeDeleteResponseAction = "update_banner"
	XTweetLikeDeleteResponseActionCreateCommunity XTweetLikeDeleteResponseAction = "create_community"
	XTweetLikeDeleteResponseActionDeleteCommunity XTweetLikeDeleteResponseAction = "delete_community"
	XTweetLikeDeleteResponseActionJoinCommunity   XTweetLikeDeleteResponseAction = "join_community"
	XTweetLikeDeleteResponseActionLeaveCommunity  XTweetLikeDeleteResponseAction = "leave_community"
)

// plannedCredits is the approved maximum. chargedCredits comes from the settled
// credit ledger. Pending or failed writes are not charged.
type XTweetLikeDeleteResponseBilling struct {
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
func (r XTweetLikeDeleteResponseBilling) RawJSON() string { return r.JSON.raw }
func (r *XTweetLikeDeleteResponseBilling) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Exact follow-up an API client or agent should perform.
type XTweetLikeDeleteResponseNextAction struct {
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
func (r XTweetLikeDeleteResponseNextAction) RawJSON() string { return r.JSON.raw }
func (r *XTweetLikeDeleteResponseNextAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Stable fingerprint and sanitized payload for replay checks.
type XTweetLikeDeleteResponseRequest struct {
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
func (r XTweetLikeDeleteResponseRequest) RawJSON() string { return r.JSON.raw }
func (r *XTweetLikeDeleteResponseRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Confirmed result produced by the write, when available.
type XTweetLikeDeleteResponseResult struct {
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
func (r XTweetLikeDeleteResponseResult) RawJSON() string { return r.JSON.raw }
func (r *XTweetLikeDeleteResponseResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type XTweetLikeDeleteResponseStatus string

const (
	XTweetLikeDeleteResponseStatusAccepted            XTweetLikeDeleteResponseStatus = "accepted"
	XTweetLikeDeleteResponseStatusDispatching         XTweetLikeDeleteResponseStatus = "dispatching"
	XTweetLikeDeleteResponseStatusPendingConfirmation XTweetLikeDeleteResponseStatus = "pending_confirmation"
	XTweetLikeDeleteResponseStatusSuccess             XTweetLikeDeleteResponseStatus = "success"
	XTweetLikeDeleteResponseStatusFailed              XTweetLikeDeleteResponseStatus = "failed"
	XTweetLikeDeleteResponseStatusExpired             XTweetLikeDeleteResponseStatus = "expired"
)

// Existing X resource targeted by the write, when applicable.
type XTweetLikeDeleteResponseTarget struct {
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
func (r XTweetLikeDeleteResponseTarget) RawJSON() string { return r.JSON.raw }
func (r *XTweetLikeDeleteResponseTarget) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type XTweetLikeNewParams struct {
	// X account identifier (@username or account ID)
	Account        string `json:"account" api:"required"`
	IdempotencyKey string `header:"Idempotency-Key" api:"required" json:"-"`
	paramObj
}

func (r XTweetLikeNewParams) MarshalJSON() (data []byte, err error) {
	type shadow XTweetLikeNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *XTweetLikeNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type XTweetLikeDeleteParams struct {
	// X account identifier (@username or account ID)
	Account        string `json:"account" api:"required"`
	IdempotencyKey string `header:"Idempotency-Key" api:"required" json:"-"`
	paramObj
}

func (r XTweetLikeDeleteParams) MarshalJSON() (data []byte, err error) {
	type shadow XTweetLikeDeleteParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *XTweetLikeDeleteParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
