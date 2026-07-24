// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package xtwitterscraper

import (
	"context"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/Xquik-dev/x-twitter-scraper-go/internal/apijson"
	"github.com/Xquik-dev/x-twitter-scraper-go/internal/requestconfig"
	"github.com/Xquik-dev/x-twitter-scraper-go/option"
	"github.com/Xquik-dev/x-twitter-scraper-go/packages/param"
	"github.com/Xquik-dev/x-twitter-scraper-go/packages/respjson"
	"github.com/Xquik-dev/x-twitter-scraper-go/shared/constant"
)

// XMediaService contains methods and other services that help with interacting
// with the x-twitter-scraper API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewXMediaService] method instead.
type XMediaService struct {
	options []option.RequestOption
}

// NewXMediaService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewXMediaService(opts ...option.RequestOption) (r XMediaService) {
	r = XMediaService{}
	r.options = opts
	return
}

// Download images and videos from tweets
func (r *XMediaService) Download(ctx context.Context, body XMediaDownloadParams, opts ...option.RequestOption) (res *XMediaDownloadResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "x/media/download"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Upload media
func (r *XMediaService) Upload(ctx context.Context, params XMediaUploadParams, opts ...option.RequestOption) (res *XMediaUploadResponse, err error) {
	if !param.IsOmitted(params.IdempotencyKey) {
		opts = append(opts, option.WithHeader("Idempotency-Key", fmt.Sprintf("%v", params.IdempotencyKey)))
	}
	opts = slices.Concat(r.options, opts)
	path := "x/media"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

type XMediaDownloadResponse struct {
	CacheHit    bool   `json:"cacheHit"`
	GalleryURL  string `json:"galleryUrl"`
	TotalMedia  int64  `json:"totalMedia"`
	TotalTweets int64  `json:"totalTweets"`
	TweetID     string `json:"tweetId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CacheHit    respjson.Field
		GalleryURL  respjson.Field
		TotalMedia  respjson.Field
		TotalTweets respjson.Field
		TweetID     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XMediaDownloadResponse) RawJSON() string { return r.JSON.raw }
func (r *XMediaDownloadResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Durable write lifecycle record. Poll statusUrl until terminal is true. Reusing
// the original Idempotency-Key returns this same record. Submit a new write only
// when safeToRetry is true, using a new key.
type XMediaUploadResponse struct {
	ID string `json:"id" api:"required"`
	// Connected account selected for the write.
	Account XMediaUploadResponseAccount `json:"account" api:"required"`
	// Any of "create_tweet", "delete_tweet", "like", "unlike", "retweet", "unretweet",
	// "follow", "unfollow", "remove_follower", "send_dm", "upload_media",
	// "update_profile", "update_avatar", "update_banner", "create_community",
	// "delete_community", "join_community", "leave_community".
	Action XMediaUploadResponseAction `json:"action" api:"required"`
	// plannedCredits is the approved maximum. chargedCredits comes from the settled
	// credit ledger. Pending or failed writes are not charged.
	Billing        XMediaUploadResponseBilling `json:"billing" api:"required"`
	Charged        bool                        `json:"charged" api:"required"`
	ChargedCredits string                      `json:"chargedCredits" api:"required"`
	// Exact follow-up an API client or agent should perform.
	NextAction  XMediaUploadResponseNextAction `json:"nextAction" api:"required"`
	Object      constant.XWriteAction          `json:"object" default:"x_write_action"`
	PollAfterMs int64                          `json:"pollAfterMs" api:"required"`
	// Stable fingerprint and sanitized payload for replay checks.
	Request XMediaUploadResponseRequest `json:"request" api:"required"`
	// Confirmed result produced by the write, when available.
	Result XMediaUploadResponseResult `json:"result" api:"required"`
	// True only when a new attempt can reasonably succeed.
	Retryable bool `json:"retryable" api:"required"`
	// True only when no write was dispatched and a new idempotency key may be used.
	SafeToRetry    bool `json:"safeToRetry" api:"required"`
	SendDispatched bool `json:"sendDispatched" api:"required"`
	// Any of "accepted", "dispatching", "pending_confirmation", "success", "failed",
	// "expired".
	Status    XMediaUploadResponseStatus `json:"status" api:"required"`
	StatusURL string                     `json:"statusUrl" api:"required"`
	Success   bool                       `json:"success" api:"required"`
	// Existing X resource targeted by the write, when applicable.
	Target        XMediaUploadResponseTarget `json:"target" api:"required"`
	TargetID      string                     `json:"targetId" api:"required"`
	Terminal      bool                       `json:"terminal" api:"required"`
	WriteActionID string                     `json:"writeActionId" api:"required"`
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
func (r XMediaUploadResponse) RawJSON() string { return r.JSON.raw }
func (r *XMediaUploadResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Connected account selected for the write.
type XMediaUploadResponseAccount struct {
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
func (r XMediaUploadResponseAccount) RawJSON() string { return r.JSON.raw }
func (r *XMediaUploadResponseAccount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type XMediaUploadResponseAction string

const (
	XMediaUploadResponseActionCreateTweet     XMediaUploadResponseAction = "create_tweet"
	XMediaUploadResponseActionDeleteTweet     XMediaUploadResponseAction = "delete_tweet"
	XMediaUploadResponseActionLike            XMediaUploadResponseAction = "like"
	XMediaUploadResponseActionUnlike          XMediaUploadResponseAction = "unlike"
	XMediaUploadResponseActionRetweet         XMediaUploadResponseAction = "retweet"
	XMediaUploadResponseActionUnretweet       XMediaUploadResponseAction = "unretweet"
	XMediaUploadResponseActionFollow          XMediaUploadResponseAction = "follow"
	XMediaUploadResponseActionUnfollow        XMediaUploadResponseAction = "unfollow"
	XMediaUploadResponseActionRemoveFollower  XMediaUploadResponseAction = "remove_follower"
	XMediaUploadResponseActionSendDm          XMediaUploadResponseAction = "send_dm"
	XMediaUploadResponseActionUploadMedia     XMediaUploadResponseAction = "upload_media"
	XMediaUploadResponseActionUpdateProfile   XMediaUploadResponseAction = "update_profile"
	XMediaUploadResponseActionUpdateAvatar    XMediaUploadResponseAction = "update_avatar"
	XMediaUploadResponseActionUpdateBanner    XMediaUploadResponseAction = "update_banner"
	XMediaUploadResponseActionCreateCommunity XMediaUploadResponseAction = "create_community"
	XMediaUploadResponseActionDeleteCommunity XMediaUploadResponseAction = "delete_community"
	XMediaUploadResponseActionJoinCommunity   XMediaUploadResponseAction = "join_community"
	XMediaUploadResponseActionLeaveCommunity  XMediaUploadResponseAction = "leave_community"
)

// plannedCredits is the approved maximum. chargedCredits comes from the settled
// credit ledger. Pending or failed writes are not charged.
type XMediaUploadResponseBilling struct {
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
func (r XMediaUploadResponseBilling) RawJSON() string { return r.JSON.raw }
func (r *XMediaUploadResponseBilling) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Exact follow-up an API client or agent should perform.
type XMediaUploadResponseNextAction struct {
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
func (r XMediaUploadResponseNextAction) RawJSON() string { return r.JSON.raw }
func (r *XMediaUploadResponseNextAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Stable fingerprint and sanitized payload for replay checks.
type XMediaUploadResponseRequest struct {
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
func (r XMediaUploadResponseRequest) RawJSON() string { return r.JSON.raw }
func (r *XMediaUploadResponseRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Confirmed result produced by the write, when available.
type XMediaUploadResponseResult struct {
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
func (r XMediaUploadResponseResult) RawJSON() string { return r.JSON.raw }
func (r *XMediaUploadResponseResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type XMediaUploadResponseStatus string

const (
	XMediaUploadResponseStatusAccepted            XMediaUploadResponseStatus = "accepted"
	XMediaUploadResponseStatusDispatching         XMediaUploadResponseStatus = "dispatching"
	XMediaUploadResponseStatusPendingConfirmation XMediaUploadResponseStatus = "pending_confirmation"
	XMediaUploadResponseStatusSuccess             XMediaUploadResponseStatus = "success"
	XMediaUploadResponseStatusFailed              XMediaUploadResponseStatus = "failed"
	XMediaUploadResponseStatusExpired             XMediaUploadResponseStatus = "expired"
)

// Existing X resource targeted by the write, when applicable.
type XMediaUploadResponseTarget struct {
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
func (r XMediaUploadResponseTarget) RawJSON() string { return r.JSON.raw }
func (r *XMediaUploadResponseTarget) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type XMediaDownloadParams struct {
	// Numeric tweet ID alias for tweetInput
	TweetID param.Opt[string] `json:"tweetId,omitzero"`
	// Tweet URL or ID (single tweet)
	TweetInput param.Opt[string] `json:"tweetInput,omitzero"`
	// Tweet URL alias for tweetInput
	TweetURL param.Opt[string] `json:"tweetUrl,omitzero"`
	// Array of tweet URLs or IDs (bulk, max 50 string items)
	TweetIDs []string `json:"tweetIds,omitzero"`
	paramObj
}

func (r XMediaDownloadParams) MarshalJSON() (data []byte, err error) {
	type shadow XMediaDownloadParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *XMediaDownloadParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type XMediaUploadParams struct {
	// X account (@username or ID) uploading media from URL
	Account string `json:"account" api:"required"`
	// HTTPS URL to download and upload as media
	URL            string `json:"url" api:"required" format:"uri"`
	IdempotencyKey string `header:"Idempotency-Key" api:"required" json:"-"`
	paramObj
}

func (r XMediaUploadParams) MarshalJSON() (data []byte, err error) {
	type shadow XMediaUploadParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *XMediaUploadParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
