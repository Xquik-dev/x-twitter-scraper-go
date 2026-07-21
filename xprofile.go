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

// X write actions (tweets, likes, follows, DMs)
//
// XProfileService contains methods and other services that help with interacting
// with the x-twitter-scraper API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewXProfileService] method instead.
type XProfileService struct {
	options []option.RequestOption
}

// NewXProfileService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewXProfileService(opts ...option.RequestOption) (r XProfileService) {
	r = XProfileService{}
	r.options = opts
	return
}

// Update X profile
func (r *XProfileService) Update(ctx context.Context, params XProfileUpdateParams, opts ...option.RequestOption) (res *XProfileUpdateResponse, err error) {
	if !param.IsOmitted(params.IdempotencyKey) {
		opts = append(opts, option.WithHeader("Idempotency-Key", fmt.Sprintf("%v", params.IdempotencyKey)))
	}
	opts = slices.Concat(r.options, opts)
	path := "x/profile"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Update profile avatar
func (r *XProfileService) UpdateAvatar(ctx context.Context, params XProfileUpdateAvatarParams, opts ...option.RequestOption) (res *XProfileUpdateAvatarResponse, err error) {
	if !param.IsOmitted(params.IdempotencyKey) {
		opts = append(opts, option.WithHeader("Idempotency-Key", fmt.Sprintf("%v", params.IdempotencyKey)))
	}
	opts = slices.Concat(r.options, opts)
	path := "x/profile/avatar"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Update profile banner
func (r *XProfileService) UpdateBanner(ctx context.Context, params XProfileUpdateBannerParams, opts ...option.RequestOption) (res *XProfileUpdateBannerResponse, err error) {
	if !param.IsOmitted(params.IdempotencyKey) {
		opts = append(opts, option.WithHeader("Idempotency-Key", fmt.Sprintf("%v", params.IdempotencyKey)))
	}
	opts = slices.Concat(r.options, opts)
	path := "x/profile/banner"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Durable write lifecycle record. Poll statusUrl until terminal is true. Reusing
// the original Idempotency-Key returns this same record. Submit a new write only
// when safeToRetry is true, using a new key.
type XProfileUpdateResponse struct {
	ID string `json:"id" api:"required"`
	// Connected account selected for the write.
	Account XProfileUpdateResponseAccount `json:"account" api:"required"`
	// Any of "create_tweet", "delete_tweet", "like", "unlike", "retweet", "unretweet",
	// "follow", "unfollow", "remove_follower", "send_dm", "upload_media",
	// "update_profile", "update_avatar", "update_banner", "create_community",
	// "delete_community", "join_community", "leave_community".
	Action XProfileUpdateResponseAction `json:"action" api:"required"`
	// plannedCredits is the approved maximum. chargedCredits comes from the settled
	// credit ledger. Pending or failed writes are not charged.
	Billing        XProfileUpdateResponseBilling `json:"billing" api:"required"`
	Charged        bool                          `json:"charged" api:"required"`
	ChargedCredits string                        `json:"chargedCredits" api:"required"`
	// Exact follow-up an API client or agent should perform.
	NextAction  XProfileUpdateResponseNextAction `json:"nextAction" api:"required"`
	Object      constant.XWriteAction            `json:"object" default:"x_write_action"`
	PollAfterMs int64                            `json:"pollAfterMs" api:"required"`
	// Stable fingerprint and sanitized payload for replay checks.
	Request XProfileUpdateResponseRequest `json:"request" api:"required"`
	// Confirmed result produced by the write, when available.
	Result XProfileUpdateResponseResult `json:"result" api:"required"`
	// True only when a new attempt can reasonably succeed.
	Retryable bool `json:"retryable" api:"required"`
	// True only when no write was dispatched and a new idempotency key may be used.
	SafeToRetry    bool `json:"safeToRetry" api:"required"`
	SendDispatched bool `json:"sendDispatched" api:"required"`
	// Any of "accepted", "dispatching", "pending_confirmation", "success", "failed",
	// "expired".
	Status    XProfileUpdateResponseStatus `json:"status" api:"required"`
	StatusURL string                       `json:"statusUrl" api:"required"`
	Success   bool                         `json:"success" api:"required"`
	// Existing X resource targeted by the write, when applicable.
	Target        XProfileUpdateResponseTarget `json:"target" api:"required"`
	TargetID      string                       `json:"targetId" api:"required"`
	Terminal      bool                         `json:"terminal" api:"required"`
	WriteActionID string                       `json:"writeActionId" api:"required"`
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
func (r XProfileUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *XProfileUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Connected account selected for the write.
type XProfileUpdateResponseAccount struct {
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
func (r XProfileUpdateResponseAccount) RawJSON() string { return r.JSON.raw }
func (r *XProfileUpdateResponseAccount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type XProfileUpdateResponseAction string

const (
	XProfileUpdateResponseActionCreateTweet     XProfileUpdateResponseAction = "create_tweet"
	XProfileUpdateResponseActionDeleteTweet     XProfileUpdateResponseAction = "delete_tweet"
	XProfileUpdateResponseActionLike            XProfileUpdateResponseAction = "like"
	XProfileUpdateResponseActionUnlike          XProfileUpdateResponseAction = "unlike"
	XProfileUpdateResponseActionRetweet         XProfileUpdateResponseAction = "retweet"
	XProfileUpdateResponseActionUnretweet       XProfileUpdateResponseAction = "unretweet"
	XProfileUpdateResponseActionFollow          XProfileUpdateResponseAction = "follow"
	XProfileUpdateResponseActionUnfollow        XProfileUpdateResponseAction = "unfollow"
	XProfileUpdateResponseActionRemoveFollower  XProfileUpdateResponseAction = "remove_follower"
	XProfileUpdateResponseActionSendDm          XProfileUpdateResponseAction = "send_dm"
	XProfileUpdateResponseActionUploadMedia     XProfileUpdateResponseAction = "upload_media"
	XProfileUpdateResponseActionUpdateProfile   XProfileUpdateResponseAction = "update_profile"
	XProfileUpdateResponseActionUpdateAvatar    XProfileUpdateResponseAction = "update_avatar"
	XProfileUpdateResponseActionUpdateBanner    XProfileUpdateResponseAction = "update_banner"
	XProfileUpdateResponseActionCreateCommunity XProfileUpdateResponseAction = "create_community"
	XProfileUpdateResponseActionDeleteCommunity XProfileUpdateResponseAction = "delete_community"
	XProfileUpdateResponseActionJoinCommunity   XProfileUpdateResponseAction = "join_community"
	XProfileUpdateResponseActionLeaveCommunity  XProfileUpdateResponseAction = "leave_community"
)

// plannedCredits is the approved maximum. chargedCredits comes from the settled
// credit ledger. Pending or failed writes are not charged.
type XProfileUpdateResponseBilling struct {
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
func (r XProfileUpdateResponseBilling) RawJSON() string { return r.JSON.raw }
func (r *XProfileUpdateResponseBilling) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Exact follow-up an API client or agent should perform.
type XProfileUpdateResponseNextAction struct {
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
func (r XProfileUpdateResponseNextAction) RawJSON() string { return r.JSON.raw }
func (r *XProfileUpdateResponseNextAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Stable fingerprint and sanitized payload for replay checks.
type XProfileUpdateResponseRequest struct {
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
func (r XProfileUpdateResponseRequest) RawJSON() string { return r.JSON.raw }
func (r *XProfileUpdateResponseRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Confirmed result produced by the write, when available.
type XProfileUpdateResponseResult struct {
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
func (r XProfileUpdateResponseResult) RawJSON() string { return r.JSON.raw }
func (r *XProfileUpdateResponseResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type XProfileUpdateResponseStatus string

const (
	XProfileUpdateResponseStatusAccepted            XProfileUpdateResponseStatus = "accepted"
	XProfileUpdateResponseStatusDispatching         XProfileUpdateResponseStatus = "dispatching"
	XProfileUpdateResponseStatusPendingConfirmation XProfileUpdateResponseStatus = "pending_confirmation"
	XProfileUpdateResponseStatusSuccess             XProfileUpdateResponseStatus = "success"
	XProfileUpdateResponseStatusFailed              XProfileUpdateResponseStatus = "failed"
	XProfileUpdateResponseStatusExpired             XProfileUpdateResponseStatus = "expired"
)

// Existing X resource targeted by the write, when applicable.
type XProfileUpdateResponseTarget struct {
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
func (r XProfileUpdateResponseTarget) RawJSON() string { return r.JSON.raw }
func (r *XProfileUpdateResponseTarget) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Durable write lifecycle record. Poll statusUrl until terminal is true. Reusing
// the original Idempotency-Key returns this same record. Submit a new write only
// when safeToRetry is true, using a new key.
type XProfileUpdateAvatarResponse struct {
	ID string `json:"id" api:"required"`
	// Connected account selected for the write.
	Account XProfileUpdateAvatarResponseAccount `json:"account" api:"required"`
	// Any of "create_tweet", "delete_tweet", "like", "unlike", "retweet", "unretweet",
	// "follow", "unfollow", "remove_follower", "send_dm", "upload_media",
	// "update_profile", "update_avatar", "update_banner", "create_community",
	// "delete_community", "join_community", "leave_community".
	Action XProfileUpdateAvatarResponseAction `json:"action" api:"required"`
	// plannedCredits is the approved maximum. chargedCredits comes from the settled
	// credit ledger. Pending or failed writes are not charged.
	Billing        XProfileUpdateAvatarResponseBilling `json:"billing" api:"required"`
	Charged        bool                                `json:"charged" api:"required"`
	ChargedCredits string                              `json:"chargedCredits" api:"required"`
	// Exact follow-up an API client or agent should perform.
	NextAction  XProfileUpdateAvatarResponseNextAction `json:"nextAction" api:"required"`
	Object      constant.XWriteAction                  `json:"object" default:"x_write_action"`
	PollAfterMs int64                                  `json:"pollAfterMs" api:"required"`
	// Stable fingerprint and sanitized payload for replay checks.
	Request XProfileUpdateAvatarResponseRequest `json:"request" api:"required"`
	// Confirmed result produced by the write, when available.
	Result XProfileUpdateAvatarResponseResult `json:"result" api:"required"`
	// True only when a new attempt can reasonably succeed.
	Retryable bool `json:"retryable" api:"required"`
	// True only when no write was dispatched and a new idempotency key may be used.
	SafeToRetry    bool `json:"safeToRetry" api:"required"`
	SendDispatched bool `json:"sendDispatched" api:"required"`
	// Any of "accepted", "dispatching", "pending_confirmation", "success", "failed",
	// "expired".
	Status    XProfileUpdateAvatarResponseStatus `json:"status" api:"required"`
	StatusURL string                             `json:"statusUrl" api:"required"`
	Success   bool                               `json:"success" api:"required"`
	// Existing X resource targeted by the write, when applicable.
	Target        XProfileUpdateAvatarResponseTarget `json:"target" api:"required"`
	TargetID      string                             `json:"targetId" api:"required"`
	Terminal      bool                               `json:"terminal" api:"required"`
	WriteActionID string                             `json:"writeActionId" api:"required"`
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
func (r XProfileUpdateAvatarResponse) RawJSON() string { return r.JSON.raw }
func (r *XProfileUpdateAvatarResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Connected account selected for the write.
type XProfileUpdateAvatarResponseAccount struct {
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
func (r XProfileUpdateAvatarResponseAccount) RawJSON() string { return r.JSON.raw }
func (r *XProfileUpdateAvatarResponseAccount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type XProfileUpdateAvatarResponseAction string

const (
	XProfileUpdateAvatarResponseActionCreateTweet     XProfileUpdateAvatarResponseAction = "create_tweet"
	XProfileUpdateAvatarResponseActionDeleteTweet     XProfileUpdateAvatarResponseAction = "delete_tweet"
	XProfileUpdateAvatarResponseActionLike            XProfileUpdateAvatarResponseAction = "like"
	XProfileUpdateAvatarResponseActionUnlike          XProfileUpdateAvatarResponseAction = "unlike"
	XProfileUpdateAvatarResponseActionRetweet         XProfileUpdateAvatarResponseAction = "retweet"
	XProfileUpdateAvatarResponseActionUnretweet       XProfileUpdateAvatarResponseAction = "unretweet"
	XProfileUpdateAvatarResponseActionFollow          XProfileUpdateAvatarResponseAction = "follow"
	XProfileUpdateAvatarResponseActionUnfollow        XProfileUpdateAvatarResponseAction = "unfollow"
	XProfileUpdateAvatarResponseActionRemoveFollower  XProfileUpdateAvatarResponseAction = "remove_follower"
	XProfileUpdateAvatarResponseActionSendDm          XProfileUpdateAvatarResponseAction = "send_dm"
	XProfileUpdateAvatarResponseActionUploadMedia     XProfileUpdateAvatarResponseAction = "upload_media"
	XProfileUpdateAvatarResponseActionUpdateProfile   XProfileUpdateAvatarResponseAction = "update_profile"
	XProfileUpdateAvatarResponseActionUpdateAvatar    XProfileUpdateAvatarResponseAction = "update_avatar"
	XProfileUpdateAvatarResponseActionUpdateBanner    XProfileUpdateAvatarResponseAction = "update_banner"
	XProfileUpdateAvatarResponseActionCreateCommunity XProfileUpdateAvatarResponseAction = "create_community"
	XProfileUpdateAvatarResponseActionDeleteCommunity XProfileUpdateAvatarResponseAction = "delete_community"
	XProfileUpdateAvatarResponseActionJoinCommunity   XProfileUpdateAvatarResponseAction = "join_community"
	XProfileUpdateAvatarResponseActionLeaveCommunity  XProfileUpdateAvatarResponseAction = "leave_community"
)

// plannedCredits is the approved maximum. chargedCredits comes from the settled
// credit ledger. Pending or failed writes are not charged.
type XProfileUpdateAvatarResponseBilling struct {
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
func (r XProfileUpdateAvatarResponseBilling) RawJSON() string { return r.JSON.raw }
func (r *XProfileUpdateAvatarResponseBilling) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Exact follow-up an API client or agent should perform.
type XProfileUpdateAvatarResponseNextAction struct {
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
func (r XProfileUpdateAvatarResponseNextAction) RawJSON() string { return r.JSON.raw }
func (r *XProfileUpdateAvatarResponseNextAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Stable fingerprint and sanitized payload for replay checks.
type XProfileUpdateAvatarResponseRequest struct {
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
func (r XProfileUpdateAvatarResponseRequest) RawJSON() string { return r.JSON.raw }
func (r *XProfileUpdateAvatarResponseRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Confirmed result produced by the write, when available.
type XProfileUpdateAvatarResponseResult struct {
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
func (r XProfileUpdateAvatarResponseResult) RawJSON() string { return r.JSON.raw }
func (r *XProfileUpdateAvatarResponseResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type XProfileUpdateAvatarResponseStatus string

const (
	XProfileUpdateAvatarResponseStatusAccepted            XProfileUpdateAvatarResponseStatus = "accepted"
	XProfileUpdateAvatarResponseStatusDispatching         XProfileUpdateAvatarResponseStatus = "dispatching"
	XProfileUpdateAvatarResponseStatusPendingConfirmation XProfileUpdateAvatarResponseStatus = "pending_confirmation"
	XProfileUpdateAvatarResponseStatusSuccess             XProfileUpdateAvatarResponseStatus = "success"
	XProfileUpdateAvatarResponseStatusFailed              XProfileUpdateAvatarResponseStatus = "failed"
	XProfileUpdateAvatarResponseStatusExpired             XProfileUpdateAvatarResponseStatus = "expired"
)

// Existing X resource targeted by the write, when applicable.
type XProfileUpdateAvatarResponseTarget struct {
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
func (r XProfileUpdateAvatarResponseTarget) RawJSON() string { return r.JSON.raw }
func (r *XProfileUpdateAvatarResponseTarget) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Durable write lifecycle record. Poll statusUrl until terminal is true. Reusing
// the original Idempotency-Key returns this same record. Submit a new write only
// when safeToRetry is true, using a new key.
type XProfileUpdateBannerResponse struct {
	ID string `json:"id" api:"required"`
	// Connected account selected for the write.
	Account XProfileUpdateBannerResponseAccount `json:"account" api:"required"`
	// Any of "create_tweet", "delete_tweet", "like", "unlike", "retweet", "unretweet",
	// "follow", "unfollow", "remove_follower", "send_dm", "upload_media",
	// "update_profile", "update_avatar", "update_banner", "create_community",
	// "delete_community", "join_community", "leave_community".
	Action XProfileUpdateBannerResponseAction `json:"action" api:"required"`
	// plannedCredits is the approved maximum. chargedCredits comes from the settled
	// credit ledger. Pending or failed writes are not charged.
	Billing        XProfileUpdateBannerResponseBilling `json:"billing" api:"required"`
	Charged        bool                                `json:"charged" api:"required"`
	ChargedCredits string                              `json:"chargedCredits" api:"required"`
	// Exact follow-up an API client or agent should perform.
	NextAction  XProfileUpdateBannerResponseNextAction `json:"nextAction" api:"required"`
	Object      constant.XWriteAction                  `json:"object" default:"x_write_action"`
	PollAfterMs int64                                  `json:"pollAfterMs" api:"required"`
	// Stable fingerprint and sanitized payload for replay checks.
	Request XProfileUpdateBannerResponseRequest `json:"request" api:"required"`
	// Confirmed result produced by the write, when available.
	Result XProfileUpdateBannerResponseResult `json:"result" api:"required"`
	// True only when a new attempt can reasonably succeed.
	Retryable bool `json:"retryable" api:"required"`
	// True only when no write was dispatched and a new idempotency key may be used.
	SafeToRetry    bool `json:"safeToRetry" api:"required"`
	SendDispatched bool `json:"sendDispatched" api:"required"`
	// Any of "accepted", "dispatching", "pending_confirmation", "success", "failed",
	// "expired".
	Status    XProfileUpdateBannerResponseStatus `json:"status" api:"required"`
	StatusURL string                             `json:"statusUrl" api:"required"`
	Success   bool                               `json:"success" api:"required"`
	// Existing X resource targeted by the write, when applicable.
	Target        XProfileUpdateBannerResponseTarget `json:"target" api:"required"`
	TargetID      string                             `json:"targetId" api:"required"`
	Terminal      bool                               `json:"terminal" api:"required"`
	WriteActionID string                             `json:"writeActionId" api:"required"`
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
func (r XProfileUpdateBannerResponse) RawJSON() string { return r.JSON.raw }
func (r *XProfileUpdateBannerResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Connected account selected for the write.
type XProfileUpdateBannerResponseAccount struct {
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
func (r XProfileUpdateBannerResponseAccount) RawJSON() string { return r.JSON.raw }
func (r *XProfileUpdateBannerResponseAccount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type XProfileUpdateBannerResponseAction string

const (
	XProfileUpdateBannerResponseActionCreateTweet     XProfileUpdateBannerResponseAction = "create_tweet"
	XProfileUpdateBannerResponseActionDeleteTweet     XProfileUpdateBannerResponseAction = "delete_tweet"
	XProfileUpdateBannerResponseActionLike            XProfileUpdateBannerResponseAction = "like"
	XProfileUpdateBannerResponseActionUnlike          XProfileUpdateBannerResponseAction = "unlike"
	XProfileUpdateBannerResponseActionRetweet         XProfileUpdateBannerResponseAction = "retweet"
	XProfileUpdateBannerResponseActionUnretweet       XProfileUpdateBannerResponseAction = "unretweet"
	XProfileUpdateBannerResponseActionFollow          XProfileUpdateBannerResponseAction = "follow"
	XProfileUpdateBannerResponseActionUnfollow        XProfileUpdateBannerResponseAction = "unfollow"
	XProfileUpdateBannerResponseActionRemoveFollower  XProfileUpdateBannerResponseAction = "remove_follower"
	XProfileUpdateBannerResponseActionSendDm          XProfileUpdateBannerResponseAction = "send_dm"
	XProfileUpdateBannerResponseActionUploadMedia     XProfileUpdateBannerResponseAction = "upload_media"
	XProfileUpdateBannerResponseActionUpdateProfile   XProfileUpdateBannerResponseAction = "update_profile"
	XProfileUpdateBannerResponseActionUpdateAvatar    XProfileUpdateBannerResponseAction = "update_avatar"
	XProfileUpdateBannerResponseActionUpdateBanner    XProfileUpdateBannerResponseAction = "update_banner"
	XProfileUpdateBannerResponseActionCreateCommunity XProfileUpdateBannerResponseAction = "create_community"
	XProfileUpdateBannerResponseActionDeleteCommunity XProfileUpdateBannerResponseAction = "delete_community"
	XProfileUpdateBannerResponseActionJoinCommunity   XProfileUpdateBannerResponseAction = "join_community"
	XProfileUpdateBannerResponseActionLeaveCommunity  XProfileUpdateBannerResponseAction = "leave_community"
)

// plannedCredits is the approved maximum. chargedCredits comes from the settled
// credit ledger. Pending or failed writes are not charged.
type XProfileUpdateBannerResponseBilling struct {
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
func (r XProfileUpdateBannerResponseBilling) RawJSON() string { return r.JSON.raw }
func (r *XProfileUpdateBannerResponseBilling) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Exact follow-up an API client or agent should perform.
type XProfileUpdateBannerResponseNextAction struct {
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
func (r XProfileUpdateBannerResponseNextAction) RawJSON() string { return r.JSON.raw }
func (r *XProfileUpdateBannerResponseNextAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Stable fingerprint and sanitized payload for replay checks.
type XProfileUpdateBannerResponseRequest struct {
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
func (r XProfileUpdateBannerResponseRequest) RawJSON() string { return r.JSON.raw }
func (r *XProfileUpdateBannerResponseRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Confirmed result produced by the write, when available.
type XProfileUpdateBannerResponseResult struct {
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
func (r XProfileUpdateBannerResponseResult) RawJSON() string { return r.JSON.raw }
func (r *XProfileUpdateBannerResponseResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type XProfileUpdateBannerResponseStatus string

const (
	XProfileUpdateBannerResponseStatusAccepted            XProfileUpdateBannerResponseStatus = "accepted"
	XProfileUpdateBannerResponseStatusDispatching         XProfileUpdateBannerResponseStatus = "dispatching"
	XProfileUpdateBannerResponseStatusPendingConfirmation XProfileUpdateBannerResponseStatus = "pending_confirmation"
	XProfileUpdateBannerResponseStatusSuccess             XProfileUpdateBannerResponseStatus = "success"
	XProfileUpdateBannerResponseStatusFailed              XProfileUpdateBannerResponseStatus = "failed"
	XProfileUpdateBannerResponseStatusExpired             XProfileUpdateBannerResponseStatus = "expired"
)

// Existing X resource targeted by the write, when applicable.
type XProfileUpdateBannerResponseTarget struct {
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
func (r XProfileUpdateBannerResponseTarget) RawJSON() string { return r.JSON.raw }
func (r *XProfileUpdateBannerResponseTarget) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type XProfileUpdateParams struct {
	// X account (@username or ID) to update profile
	Account        string `json:"account" api:"required"`
	IdempotencyKey string `header:"Idempotency-Key" api:"required" json:"-"`
	// Bio description
	Description param.Opt[string] `json:"description,omitzero"`
	Location    param.Opt[string] `json:"location,omitzero"`
	// Display name
	Name param.Opt[string] `json:"name,omitzero"`
	// Website URL
	URL param.Opt[string] `json:"url,omitzero"`
	paramObj
}

func (r XProfileUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow XProfileUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *XProfileUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type XProfileUpdateAvatarParams struct {
	// X account (@username or ID) receiving avatar from URL
	Account string `json:"account" api:"required"`
	// HTTPS URL to the avatar image to download
	URL            string `json:"url" api:"required" format:"uri"`
	IdempotencyKey string `header:"Idempotency-Key" api:"required" json:"-"`
	paramObj
}

func (r XProfileUpdateAvatarParams) MarshalJSON() (data []byte, err error) {
	type shadow XProfileUpdateAvatarParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *XProfileUpdateAvatarParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type XProfileUpdateBannerParams struct {
	// X account (@username or ID) receiving banner from URL
	Account string `json:"account" api:"required"`
	// HTTPS URL to the banner image to download
	URL            string `json:"url" api:"required" format:"uri"`
	IdempotencyKey string `header:"Idempotency-Key" api:"required" json:"-"`
	paramObj
}

func (r XProfileUpdateBannerParams) MarshalJSON() (data []byte, err error) {
	type shadow XProfileUpdateBannerParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *XProfileUpdateBannerParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
