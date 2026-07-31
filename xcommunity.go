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
	"github.com/Xquik-dev/x-twitter-scraper-go/internal/apiquery"
	"github.com/Xquik-dev/x-twitter-scraper-go/internal/requestconfig"
	"github.com/Xquik-dev/x-twitter-scraper-go/option"
	"github.com/Xquik-dev/x-twitter-scraper-go/packages/param"
	"github.com/Xquik-dev/x-twitter-scraper-go/packages/respjson"
	"github.com/Xquik-dev/x-twitter-scraper-go/shared"
	"github.com/Xquik-dev/x-twitter-scraper-go/shared/constant"
)

// XCommunityService contains methods and other services that help with interacting
// with the x-twitter-scraper API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewXCommunityService] method instead.
type XCommunityService struct {
	options []option.RequestOption
	// X write actions (tweets, likes, follows, DMs)
	Join XCommunityJoinService
	// X Community info, members, and tweets
	Tweets XCommunityTweetService
}

// NewXCommunityService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewXCommunityService(opts ...option.RequestOption) (r XCommunityService) {
	r = XCommunityService{}
	r.options = opts
	r.Join = NewXCommunityJoinService(opts...)
	r.Tweets = NewXCommunityTweetService(opts...)
	return
}

// Create community
func (r *XCommunityService) New(ctx context.Context, params XCommunityNewParams, opts ...option.RequestOption) (res *XCommunityNewResponse, err error) {
	if !param.IsOmitted(params.IdempotencyKey) {
		opts = append(opts, option.WithHeader("Idempotency-Key", fmt.Sprintf("%v", params.IdempotencyKey)))
	}
	opts = slices.Concat(r.options, opts)
	path := "x/communities"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Delete community
func (r *XCommunityService) Delete(ctx context.Context, id string, params XCommunityDeleteParams, opts ...option.RequestOption) (res *XCommunityDeleteResponse, err error) {
	if !param.IsOmitted(params.IdempotencyKey) {
		opts = append(opts, option.WithHeader("Idempotency-Key", fmt.Sprintf("%v", params.IdempotencyKey)))
	}
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("x/communities/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, params, &res, opts...)
	return res, err
}

// Get community name, description and member count
func (r *XCommunityService) GetInfo(ctx context.Context, id string, opts ...option.RequestOption) (res *XCommunityGetInfoResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("x/communities/%s/info", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// List members of a community
func (r *XCommunityService) GetMembers(ctx context.Context, id string, query XCommunityGetMembersParams, opts ...option.RequestOption) (res *shared.PaginatedUsers, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("x/communities/%s/members", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// List moderators of a community
func (r *XCommunityService) GetModerators(ctx context.Context, id string, query XCommunityGetModeratorsParams, opts ...option.RequestOption) (res *shared.PaginatedUsers, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("x/communities/%s/moderators", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Returns tweets, not community records. Requires a Community ID.
func (r *XCommunityService) GetSearch(ctx context.Context, query XCommunityGetSearchParams, opts ...option.RequestOption) (res *shared.PaginatedTweets, err error) {
	opts = slices.Concat(r.options, opts)
	path := "x/communities/search"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Durable write lifecycle record. Poll statusUrl until terminal is true. Reusing
// the original Idempotency-Key returns this same record. Submit a new write only
// when safeToRetry is true, using a new key.
type XCommunityNewResponse struct {
	ID string `json:"id" api:"required"`
	// Connected account selected for the write.
	Account XCommunityNewResponseAccount `json:"account" api:"required"`
	// Any of "create_tweet", "delete_tweet", "like", "unlike", "retweet", "unretweet",
	// "follow", "unfollow", "remove_follower", "send_dm", "upload_media",
	// "update_profile", "update_avatar", "update_banner", "create_community",
	// "delete_community", "join_community", "leave_community".
	Action XCommunityNewResponseAction `json:"action" api:"required"`
	// plannedCredits is the approved maximum. chargedCredits comes from the settled
	// credit ledger. Pending or failed writes are not charged.
	Billing        XCommunityNewResponseBilling `json:"billing" api:"required"`
	Charged        bool                         `json:"charged" api:"required"`
	ChargedCredits string                       `json:"chargedCredits" api:"required"`
	// Exact follow-up an API client or agent should perform.
	NextAction  XCommunityNewResponseNextAction `json:"nextAction" api:"required"`
	Object      constant.XWriteAction           `json:"object" default:"x_write_action"`
	PollAfterMs int64                           `json:"pollAfterMs" api:"required"`
	// Stable fingerprint and sanitized payload for replay checks.
	Request XCommunityNewResponseRequest `json:"request" api:"required"`
	// Confirmed result produced by the write, when available.
	Result XCommunityNewResponseResult `json:"result" api:"required"`
	// True only when a new attempt can reasonably succeed.
	Retryable bool `json:"retryable" api:"required"`
	// True only when no write was dispatched and a new idempotency key may be used.
	SafeToRetry    bool `json:"safeToRetry" api:"required"`
	SendDispatched bool `json:"sendDispatched" api:"required"`
	// Any of "accepted", "dispatching", "pending_confirmation", "success", "failed",
	// "expired".
	Status    XCommunityNewResponseStatus `json:"status" api:"required"`
	StatusURL string                      `json:"statusUrl" api:"required"`
	Success   bool                        `json:"success" api:"required"`
	// Existing X resource targeted by the write, when applicable.
	Target        XCommunityNewResponseTarget `json:"target" api:"required"`
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
func (r XCommunityNewResponse) RawJSON() string { return r.JSON.raw }
func (r *XCommunityNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Connected account selected for the write.
type XCommunityNewResponseAccount struct {
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
func (r XCommunityNewResponseAccount) RawJSON() string { return r.JSON.raw }
func (r *XCommunityNewResponseAccount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type XCommunityNewResponseAction string

const (
	XCommunityNewResponseActionCreateTweet     XCommunityNewResponseAction = "create_tweet"
	XCommunityNewResponseActionDeleteTweet     XCommunityNewResponseAction = "delete_tweet"
	XCommunityNewResponseActionLike            XCommunityNewResponseAction = "like"
	XCommunityNewResponseActionUnlike          XCommunityNewResponseAction = "unlike"
	XCommunityNewResponseActionRetweet         XCommunityNewResponseAction = "retweet"
	XCommunityNewResponseActionUnretweet       XCommunityNewResponseAction = "unretweet"
	XCommunityNewResponseActionFollow          XCommunityNewResponseAction = "follow"
	XCommunityNewResponseActionUnfollow        XCommunityNewResponseAction = "unfollow"
	XCommunityNewResponseActionRemoveFollower  XCommunityNewResponseAction = "remove_follower"
	XCommunityNewResponseActionSendDm          XCommunityNewResponseAction = "send_dm"
	XCommunityNewResponseActionUploadMedia     XCommunityNewResponseAction = "upload_media"
	XCommunityNewResponseActionUpdateProfile   XCommunityNewResponseAction = "update_profile"
	XCommunityNewResponseActionUpdateAvatar    XCommunityNewResponseAction = "update_avatar"
	XCommunityNewResponseActionUpdateBanner    XCommunityNewResponseAction = "update_banner"
	XCommunityNewResponseActionCreateCommunity XCommunityNewResponseAction = "create_community"
	XCommunityNewResponseActionDeleteCommunity XCommunityNewResponseAction = "delete_community"
	XCommunityNewResponseActionJoinCommunity   XCommunityNewResponseAction = "join_community"
	XCommunityNewResponseActionLeaveCommunity  XCommunityNewResponseAction = "leave_community"
)

// plannedCredits is the approved maximum. chargedCredits comes from the settled
// credit ledger. Pending or failed writes are not charged.
type XCommunityNewResponseBilling struct {
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
func (r XCommunityNewResponseBilling) RawJSON() string { return r.JSON.raw }
func (r *XCommunityNewResponseBilling) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Exact follow-up an API client or agent should perform.
type XCommunityNewResponseNextAction struct {
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
func (r XCommunityNewResponseNextAction) RawJSON() string { return r.JSON.raw }
func (r *XCommunityNewResponseNextAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Stable fingerprint and sanitized payload for replay checks.
type XCommunityNewResponseRequest struct {
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
func (r XCommunityNewResponseRequest) RawJSON() string { return r.JSON.raw }
func (r *XCommunityNewResponseRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Confirmed result produced by the write, when available.
type XCommunityNewResponseResult struct {
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
func (r XCommunityNewResponseResult) RawJSON() string { return r.JSON.raw }
func (r *XCommunityNewResponseResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type XCommunityNewResponseStatus string

const (
	XCommunityNewResponseStatusAccepted            XCommunityNewResponseStatus = "accepted"
	XCommunityNewResponseStatusDispatching         XCommunityNewResponseStatus = "dispatching"
	XCommunityNewResponseStatusPendingConfirmation XCommunityNewResponseStatus = "pending_confirmation"
	XCommunityNewResponseStatusSuccess             XCommunityNewResponseStatus = "success"
	XCommunityNewResponseStatusFailed              XCommunityNewResponseStatus = "failed"
	XCommunityNewResponseStatusExpired             XCommunityNewResponseStatus = "expired"
)

// Existing X resource targeted by the write, when applicable.
type XCommunityNewResponseTarget struct {
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
func (r XCommunityNewResponseTarget) RawJSON() string { return r.JSON.raw }
func (r *XCommunityNewResponseTarget) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Durable write lifecycle record. Poll statusUrl until terminal is true. Reusing
// the original Idempotency-Key returns this same record. Submit a new write only
// when safeToRetry is true, using a new key.
type XCommunityDeleteResponse struct {
	ID string `json:"id" api:"required"`
	// Connected account selected for the write.
	Account XCommunityDeleteResponseAccount `json:"account" api:"required"`
	// Any of "create_tweet", "delete_tweet", "like", "unlike", "retweet", "unretweet",
	// "follow", "unfollow", "remove_follower", "send_dm", "upload_media",
	// "update_profile", "update_avatar", "update_banner", "create_community",
	// "delete_community", "join_community", "leave_community".
	Action XCommunityDeleteResponseAction `json:"action" api:"required"`
	// plannedCredits is the approved maximum. chargedCredits comes from the settled
	// credit ledger. Pending or failed writes are not charged.
	Billing        XCommunityDeleteResponseBilling `json:"billing" api:"required"`
	Charged        bool                            `json:"charged" api:"required"`
	ChargedCredits string                          `json:"chargedCredits" api:"required"`
	// Exact follow-up an API client or agent should perform.
	NextAction  XCommunityDeleteResponseNextAction `json:"nextAction" api:"required"`
	Object      constant.XWriteAction              `json:"object" default:"x_write_action"`
	PollAfterMs int64                              `json:"pollAfterMs" api:"required"`
	// Stable fingerprint and sanitized payload for replay checks.
	Request XCommunityDeleteResponseRequest `json:"request" api:"required"`
	// Confirmed result produced by the write, when available.
	Result XCommunityDeleteResponseResult `json:"result" api:"required"`
	// True only when a new attempt can reasonably succeed.
	Retryable bool `json:"retryable" api:"required"`
	// True only when no write was dispatched and a new idempotency key may be used.
	SafeToRetry    bool `json:"safeToRetry" api:"required"`
	SendDispatched bool `json:"sendDispatched" api:"required"`
	// Any of "accepted", "dispatching", "pending_confirmation", "success", "failed",
	// "expired".
	Status    XCommunityDeleteResponseStatus `json:"status" api:"required"`
	StatusURL string                         `json:"statusUrl" api:"required"`
	Success   bool                           `json:"success" api:"required"`
	// Existing X resource targeted by the write, when applicable.
	Target        XCommunityDeleteResponseTarget `json:"target" api:"required"`
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
func (r XCommunityDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *XCommunityDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Connected account selected for the write.
type XCommunityDeleteResponseAccount struct {
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
func (r XCommunityDeleteResponseAccount) RawJSON() string { return r.JSON.raw }
func (r *XCommunityDeleteResponseAccount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type XCommunityDeleteResponseAction string

const (
	XCommunityDeleteResponseActionCreateTweet     XCommunityDeleteResponseAction = "create_tweet"
	XCommunityDeleteResponseActionDeleteTweet     XCommunityDeleteResponseAction = "delete_tweet"
	XCommunityDeleteResponseActionLike            XCommunityDeleteResponseAction = "like"
	XCommunityDeleteResponseActionUnlike          XCommunityDeleteResponseAction = "unlike"
	XCommunityDeleteResponseActionRetweet         XCommunityDeleteResponseAction = "retweet"
	XCommunityDeleteResponseActionUnretweet       XCommunityDeleteResponseAction = "unretweet"
	XCommunityDeleteResponseActionFollow          XCommunityDeleteResponseAction = "follow"
	XCommunityDeleteResponseActionUnfollow        XCommunityDeleteResponseAction = "unfollow"
	XCommunityDeleteResponseActionRemoveFollower  XCommunityDeleteResponseAction = "remove_follower"
	XCommunityDeleteResponseActionSendDm          XCommunityDeleteResponseAction = "send_dm"
	XCommunityDeleteResponseActionUploadMedia     XCommunityDeleteResponseAction = "upload_media"
	XCommunityDeleteResponseActionUpdateProfile   XCommunityDeleteResponseAction = "update_profile"
	XCommunityDeleteResponseActionUpdateAvatar    XCommunityDeleteResponseAction = "update_avatar"
	XCommunityDeleteResponseActionUpdateBanner    XCommunityDeleteResponseAction = "update_banner"
	XCommunityDeleteResponseActionCreateCommunity XCommunityDeleteResponseAction = "create_community"
	XCommunityDeleteResponseActionDeleteCommunity XCommunityDeleteResponseAction = "delete_community"
	XCommunityDeleteResponseActionJoinCommunity   XCommunityDeleteResponseAction = "join_community"
	XCommunityDeleteResponseActionLeaveCommunity  XCommunityDeleteResponseAction = "leave_community"
)

// plannedCredits is the approved maximum. chargedCredits comes from the settled
// credit ledger. Pending or failed writes are not charged.
type XCommunityDeleteResponseBilling struct {
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
func (r XCommunityDeleteResponseBilling) RawJSON() string { return r.JSON.raw }
func (r *XCommunityDeleteResponseBilling) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Exact follow-up an API client or agent should perform.
type XCommunityDeleteResponseNextAction struct {
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
func (r XCommunityDeleteResponseNextAction) RawJSON() string { return r.JSON.raw }
func (r *XCommunityDeleteResponseNextAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Stable fingerprint and sanitized payload for replay checks.
type XCommunityDeleteResponseRequest struct {
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
func (r XCommunityDeleteResponseRequest) RawJSON() string { return r.JSON.raw }
func (r *XCommunityDeleteResponseRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Confirmed result produced by the write, when available.
type XCommunityDeleteResponseResult struct {
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
func (r XCommunityDeleteResponseResult) RawJSON() string { return r.JSON.raw }
func (r *XCommunityDeleteResponseResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type XCommunityDeleteResponseStatus string

const (
	XCommunityDeleteResponseStatusAccepted            XCommunityDeleteResponseStatus = "accepted"
	XCommunityDeleteResponseStatusDispatching         XCommunityDeleteResponseStatus = "dispatching"
	XCommunityDeleteResponseStatusPendingConfirmation XCommunityDeleteResponseStatus = "pending_confirmation"
	XCommunityDeleteResponseStatusSuccess             XCommunityDeleteResponseStatus = "success"
	XCommunityDeleteResponseStatusFailed              XCommunityDeleteResponseStatus = "failed"
	XCommunityDeleteResponseStatusExpired             XCommunityDeleteResponseStatus = "expired"
)

// Existing X resource targeted by the write, when applicable.
type XCommunityDeleteResponseTarget struct {
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
func (r XCommunityDeleteResponseTarget) RawJSON() string { return r.JSON.raw }
func (r *XCommunityDeleteResponseTarget) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type XCommunityGetInfoResponse struct {
	// Community info object
	Community XCommunityGetInfoResponseCommunity `json:"community" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Community   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XCommunityGetInfoResponse) RawJSON() string { return r.JSON.raw }
func (r *XCommunityGetInfoResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Community info object
type XCommunityGetInfoResponseCommunity struct {
	// Unique community identifier
	ID string `json:"id" api:"required"`
	// Community banner image URL
	BannerURL string `json:"banner_url"`
	// Community creation timestamp
	CreatedAt string                                    `json:"created_at"`
	Creator   XCommunityGetInfoResponseCommunityCreator `json:"creator"`
	// About text for the community
	Description string `json:"description"`
	// Invitation policy
	InvitesPolicy string `json:"invites_policy"`
	// Whether the community is marked sensitive
	IsNsfw bool `json:"is_nsfw"`
	// Join policy (open or restricted)
	JoinPolicy string `json:"join_policy"`
	// Total member count
	MemberCount int64 `json:"member_count"`
	// Total moderator count
	ModeratorCount int64 `json:"moderator_count"`
	// Display name of the community
	Name string `json:"name"`
	// Primary topic
	PrimaryTopic XCommunityGetInfoResponseCommunityPrimaryTopic `json:"primary_topic"`
	// Community rules
	Rules []XCommunityGetInfoResponseCommunityRule `json:"rules"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		BannerURL      respjson.Field
		CreatedAt      respjson.Field
		Creator        respjson.Field
		Description    respjson.Field
		InvitesPolicy  respjson.Field
		IsNsfw         respjson.Field
		JoinPolicy     respjson.Field
		MemberCount    respjson.Field
		ModeratorCount respjson.Field
		Name           respjson.Field
		PrimaryTopic   respjson.Field
		Rules          respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XCommunityGetInfoResponseCommunity) RawJSON() string { return r.JSON.raw }
func (r *XCommunityGetInfoResponseCommunity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type XCommunityGetInfoResponseCommunityCreator struct {
	ID       string `json:"id" api:"required"`
	Username string `json:"username" api:"required"`
	Verified bool   `json:"verified" api:"required"`
	Name     string `json:"name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Username    respjson.Field
		Verified    respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XCommunityGetInfoResponseCommunityCreator) RawJSON() string { return r.JSON.raw }
func (r *XCommunityGetInfoResponseCommunityCreator) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Primary topic
type XCommunityGetInfoResponseCommunityPrimaryTopic struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XCommunityGetInfoResponseCommunityPrimaryTopic) RawJSON() string { return r.JSON.raw }
func (r *XCommunityGetInfoResponseCommunityPrimaryTopic) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type XCommunityGetInfoResponseCommunityRule struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	Name        string `json:"name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Description respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XCommunityGetInfoResponseCommunityRule) RawJSON() string { return r.JSON.raw }
func (r *XCommunityGetInfoResponseCommunityRule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type XCommunityNewParams struct {
	// X account (@username or ID) creating the community
	Account string `json:"account" api:"required"`
	// Community name
	Name           string `json:"name" api:"required"`
	IdempotencyKey string `header:"Idempotency-Key" api:"required" json:"-"`
	// Community description
	Description param.Opt[string] `json:"description,omitzero"`
	paramObj
}

func (r XCommunityNewParams) MarshalJSON() (data []byte, err error) {
	type shadow XCommunityNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *XCommunityNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type XCommunityDeleteParams struct {
	// X account (@username or ID) deleting the community
	Account string `json:"account" api:"required"`
	// Community name for confirmation
	CommunityName  string `json:"community_name" api:"required"`
	IdempotencyKey string `header:"Idempotency-Key" api:"required" json:"-"`
	paramObj
}

func (r XCommunityDeleteParams) MarshalJSON() (data []byte, err error) {
	type shadow XCommunityDeleteParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *XCommunityDeleteParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type XCommunityGetMembersParams struct {
	// Pagination cursor
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Items per page (20-200, default 20). This is an upper bound for paid
	// authenticated calls: remaining credits can reduce the returned page size, and
	// zero affordable results returns 402 insufficient_credits.
	PageSize param.Opt[int64] `query:"pageSize,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [XCommunityGetMembersParams]'s query parameters as
// `url.Values`.
func (r XCommunityGetMembersParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type XCommunityGetModeratorsParams struct {
	// Pagination cursor for community moderators
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [XCommunityGetModeratorsParams]'s query parameters as
// `url.Values`.
func (r XCommunityGetModeratorsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type XCommunityGetSearchParams struct {
	// Numeric ID of the community whose posts to search
	CommunityID string `query:"communityId" api:"required" json:"-"`
	// Search query
	Q string `query:"q" api:"required" json:"-"`
	// Pagination cursor for community search
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Maximum page items (1-100, default 20). Source, filters, or credits can reduce
	// results. Continue while has_next_page is true. Deprecated limit and count
	// aliases remain accepted.
	PageSize param.Opt[int64] `query:"pageSize,omitzero" json:"-"`
	// Sort order (Latest or Top)
	//
	// Any of "Latest", "Top".
	QueryType XCommunityGetSearchParamsQueryType `query:"queryType,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [XCommunityGetSearchParams]'s query parameters as
// `url.Values`.
func (r XCommunityGetSearchParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Sort order (Latest or Top)
type XCommunityGetSearchParamsQueryType string

const (
	XCommunityGetSearchParamsQueryTypeLatest XCommunityGetSearchParamsQueryType = "Latest"
	XCommunityGetSearchParamsQueryTypeTop    XCommunityGetSearchParamsQueryType = "Top"
)
