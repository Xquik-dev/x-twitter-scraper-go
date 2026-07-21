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
// XUserFollowService contains methods and other services that help with
// interacting with the x-twitter-scraper API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewXUserFollowService] method instead.
type XUserFollowService struct {
	options []option.RequestOption
}

// NewXUserFollowService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewXUserFollowService(opts ...option.RequestOption) (r XUserFollowService) {
	r = XUserFollowService{}
	r.options = opts
	return
}

// Follow user
func (r *XUserFollowService) New(ctx context.Context, id string, params XUserFollowNewParams, opts ...option.RequestOption) (res *XUserFollowNewResponse, err error) {
	if !param.IsOmitted(params.IdempotencyKey) {
		opts = append(opts, option.WithHeader("Idempotency-Key", fmt.Sprintf("%v", params.IdempotencyKey)))
	}
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("x/users/%s/follow", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Unfollow user
func (r *XUserFollowService) DeleteAll(ctx context.Context, id string, params XUserFollowDeleteAllParams, opts ...option.RequestOption) (res *XUserFollowDeleteAllResponse, err error) {
	if !param.IsOmitted(params.IdempotencyKey) {
		opts = append(opts, option.WithHeader("Idempotency-Key", fmt.Sprintf("%v", params.IdempotencyKey)))
	}
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("x/users/%s/follow", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, params, &res, opts...)
	return res, err
}

// Durable write lifecycle record. Poll statusUrl until terminal is true. Reusing
// the original Idempotency-Key returns this same record. Submit a new write only
// when safeToRetry is true, using a new key.
type XUserFollowNewResponse struct {
	ID string `json:"id" api:"required"`
	// Connected account selected for the write.
	Account XUserFollowNewResponseAccount `json:"account" api:"required"`
	// Any of "create_tweet", "delete_tweet", "like", "unlike", "retweet", "unretweet",
	// "follow", "unfollow", "remove_follower", "send_dm", "upload_media",
	// "update_profile", "update_avatar", "update_banner", "create_community",
	// "delete_community", "join_community", "leave_community".
	Action XUserFollowNewResponseAction `json:"action" api:"required"`
	// plannedCredits is the approved maximum. chargedCredits comes from the settled
	// credit ledger. Pending or failed writes are not charged.
	Billing        XUserFollowNewResponseBilling `json:"billing" api:"required"`
	Charged        bool                          `json:"charged" api:"required"`
	ChargedCredits string                        `json:"chargedCredits" api:"required"`
	// Exact follow-up an API client or agent should perform.
	NextAction  XUserFollowNewResponseNextAction `json:"nextAction" api:"required"`
	Object      constant.XWriteAction            `json:"object" default:"x_write_action"`
	PollAfterMs int64                            `json:"pollAfterMs" api:"required"`
	// Stable fingerprint and sanitized payload for replay checks.
	Request XUserFollowNewResponseRequest `json:"request" api:"required"`
	// Confirmed result produced by the write, when available.
	Result XUserFollowNewResponseResult `json:"result" api:"required"`
	// True only when a new attempt can reasonably succeed.
	Retryable bool `json:"retryable" api:"required"`
	// True only when no write was dispatched and a new idempotency key may be used.
	SafeToRetry    bool `json:"safeToRetry" api:"required"`
	SendDispatched bool `json:"sendDispatched" api:"required"`
	// Any of "accepted", "dispatching", "pending_confirmation", "success", "failed",
	// "expired".
	Status    XUserFollowNewResponseStatus `json:"status" api:"required"`
	StatusURL string                       `json:"statusUrl" api:"required"`
	Success   bool                         `json:"success" api:"required"`
	// Existing X resource targeted by the write, when applicable.
	Target        XUserFollowNewResponseTarget `json:"target" api:"required"`
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
func (r XUserFollowNewResponse) RawJSON() string { return r.JSON.raw }
func (r *XUserFollowNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Connected account selected for the write.
type XUserFollowNewResponseAccount struct {
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
func (r XUserFollowNewResponseAccount) RawJSON() string { return r.JSON.raw }
func (r *XUserFollowNewResponseAccount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type XUserFollowNewResponseAction string

const (
	XUserFollowNewResponseActionCreateTweet     XUserFollowNewResponseAction = "create_tweet"
	XUserFollowNewResponseActionDeleteTweet     XUserFollowNewResponseAction = "delete_tweet"
	XUserFollowNewResponseActionLike            XUserFollowNewResponseAction = "like"
	XUserFollowNewResponseActionUnlike          XUserFollowNewResponseAction = "unlike"
	XUserFollowNewResponseActionRetweet         XUserFollowNewResponseAction = "retweet"
	XUserFollowNewResponseActionUnretweet       XUserFollowNewResponseAction = "unretweet"
	XUserFollowNewResponseActionFollow          XUserFollowNewResponseAction = "follow"
	XUserFollowNewResponseActionUnfollow        XUserFollowNewResponseAction = "unfollow"
	XUserFollowNewResponseActionRemoveFollower  XUserFollowNewResponseAction = "remove_follower"
	XUserFollowNewResponseActionSendDm          XUserFollowNewResponseAction = "send_dm"
	XUserFollowNewResponseActionUploadMedia     XUserFollowNewResponseAction = "upload_media"
	XUserFollowNewResponseActionUpdateProfile   XUserFollowNewResponseAction = "update_profile"
	XUserFollowNewResponseActionUpdateAvatar    XUserFollowNewResponseAction = "update_avatar"
	XUserFollowNewResponseActionUpdateBanner    XUserFollowNewResponseAction = "update_banner"
	XUserFollowNewResponseActionCreateCommunity XUserFollowNewResponseAction = "create_community"
	XUserFollowNewResponseActionDeleteCommunity XUserFollowNewResponseAction = "delete_community"
	XUserFollowNewResponseActionJoinCommunity   XUserFollowNewResponseAction = "join_community"
	XUserFollowNewResponseActionLeaveCommunity  XUserFollowNewResponseAction = "leave_community"
)

// plannedCredits is the approved maximum. chargedCredits comes from the settled
// credit ledger. Pending or failed writes are not charged.
type XUserFollowNewResponseBilling struct {
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
func (r XUserFollowNewResponseBilling) RawJSON() string { return r.JSON.raw }
func (r *XUserFollowNewResponseBilling) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Exact follow-up an API client or agent should perform.
type XUserFollowNewResponseNextAction struct {
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
func (r XUserFollowNewResponseNextAction) RawJSON() string { return r.JSON.raw }
func (r *XUserFollowNewResponseNextAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Stable fingerprint and sanitized payload for replay checks.
type XUserFollowNewResponseRequest struct {
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
func (r XUserFollowNewResponseRequest) RawJSON() string { return r.JSON.raw }
func (r *XUserFollowNewResponseRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Confirmed result produced by the write, when available.
type XUserFollowNewResponseResult struct {
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
func (r XUserFollowNewResponseResult) RawJSON() string { return r.JSON.raw }
func (r *XUserFollowNewResponseResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type XUserFollowNewResponseStatus string

const (
	XUserFollowNewResponseStatusAccepted            XUserFollowNewResponseStatus = "accepted"
	XUserFollowNewResponseStatusDispatching         XUserFollowNewResponseStatus = "dispatching"
	XUserFollowNewResponseStatusPendingConfirmation XUserFollowNewResponseStatus = "pending_confirmation"
	XUserFollowNewResponseStatusSuccess             XUserFollowNewResponseStatus = "success"
	XUserFollowNewResponseStatusFailed              XUserFollowNewResponseStatus = "failed"
	XUserFollowNewResponseStatusExpired             XUserFollowNewResponseStatus = "expired"
)

// Existing X resource targeted by the write, when applicable.
type XUserFollowNewResponseTarget struct {
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
func (r XUserFollowNewResponseTarget) RawJSON() string { return r.JSON.raw }
func (r *XUserFollowNewResponseTarget) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Durable write lifecycle record. Poll statusUrl until terminal is true. Reusing
// the original Idempotency-Key returns this same record. Submit a new write only
// when safeToRetry is true, using a new key.
type XUserFollowDeleteAllResponse struct {
	ID string `json:"id" api:"required"`
	// Connected account selected for the write.
	Account XUserFollowDeleteAllResponseAccount `json:"account" api:"required"`
	// Any of "create_tweet", "delete_tweet", "like", "unlike", "retweet", "unretweet",
	// "follow", "unfollow", "remove_follower", "send_dm", "upload_media",
	// "update_profile", "update_avatar", "update_banner", "create_community",
	// "delete_community", "join_community", "leave_community".
	Action XUserFollowDeleteAllResponseAction `json:"action" api:"required"`
	// plannedCredits is the approved maximum. chargedCredits comes from the settled
	// credit ledger. Pending or failed writes are not charged.
	Billing        XUserFollowDeleteAllResponseBilling `json:"billing" api:"required"`
	Charged        bool                                `json:"charged" api:"required"`
	ChargedCredits string                              `json:"chargedCredits" api:"required"`
	// Exact follow-up an API client or agent should perform.
	NextAction  XUserFollowDeleteAllResponseNextAction `json:"nextAction" api:"required"`
	Object      constant.XWriteAction                  `json:"object" default:"x_write_action"`
	PollAfterMs int64                                  `json:"pollAfterMs" api:"required"`
	// Stable fingerprint and sanitized payload for replay checks.
	Request XUserFollowDeleteAllResponseRequest `json:"request" api:"required"`
	// Confirmed result produced by the write, when available.
	Result XUserFollowDeleteAllResponseResult `json:"result" api:"required"`
	// True only when a new attempt can reasonably succeed.
	Retryable bool `json:"retryable" api:"required"`
	// True only when no write was dispatched and a new idempotency key may be used.
	SafeToRetry    bool `json:"safeToRetry" api:"required"`
	SendDispatched bool `json:"sendDispatched" api:"required"`
	// Any of "accepted", "dispatching", "pending_confirmation", "success", "failed",
	// "expired".
	Status    XUserFollowDeleteAllResponseStatus `json:"status" api:"required"`
	StatusURL string                             `json:"statusUrl" api:"required"`
	Success   bool                               `json:"success" api:"required"`
	// Existing X resource targeted by the write, when applicable.
	Target        XUserFollowDeleteAllResponseTarget `json:"target" api:"required"`
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
func (r XUserFollowDeleteAllResponse) RawJSON() string { return r.JSON.raw }
func (r *XUserFollowDeleteAllResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Connected account selected for the write.
type XUserFollowDeleteAllResponseAccount struct {
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
func (r XUserFollowDeleteAllResponseAccount) RawJSON() string { return r.JSON.raw }
func (r *XUserFollowDeleteAllResponseAccount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type XUserFollowDeleteAllResponseAction string

const (
	XUserFollowDeleteAllResponseActionCreateTweet     XUserFollowDeleteAllResponseAction = "create_tweet"
	XUserFollowDeleteAllResponseActionDeleteTweet     XUserFollowDeleteAllResponseAction = "delete_tweet"
	XUserFollowDeleteAllResponseActionLike            XUserFollowDeleteAllResponseAction = "like"
	XUserFollowDeleteAllResponseActionUnlike          XUserFollowDeleteAllResponseAction = "unlike"
	XUserFollowDeleteAllResponseActionRetweet         XUserFollowDeleteAllResponseAction = "retweet"
	XUserFollowDeleteAllResponseActionUnretweet       XUserFollowDeleteAllResponseAction = "unretweet"
	XUserFollowDeleteAllResponseActionFollow          XUserFollowDeleteAllResponseAction = "follow"
	XUserFollowDeleteAllResponseActionUnfollow        XUserFollowDeleteAllResponseAction = "unfollow"
	XUserFollowDeleteAllResponseActionRemoveFollower  XUserFollowDeleteAllResponseAction = "remove_follower"
	XUserFollowDeleteAllResponseActionSendDm          XUserFollowDeleteAllResponseAction = "send_dm"
	XUserFollowDeleteAllResponseActionUploadMedia     XUserFollowDeleteAllResponseAction = "upload_media"
	XUserFollowDeleteAllResponseActionUpdateProfile   XUserFollowDeleteAllResponseAction = "update_profile"
	XUserFollowDeleteAllResponseActionUpdateAvatar    XUserFollowDeleteAllResponseAction = "update_avatar"
	XUserFollowDeleteAllResponseActionUpdateBanner    XUserFollowDeleteAllResponseAction = "update_banner"
	XUserFollowDeleteAllResponseActionCreateCommunity XUserFollowDeleteAllResponseAction = "create_community"
	XUserFollowDeleteAllResponseActionDeleteCommunity XUserFollowDeleteAllResponseAction = "delete_community"
	XUserFollowDeleteAllResponseActionJoinCommunity   XUserFollowDeleteAllResponseAction = "join_community"
	XUserFollowDeleteAllResponseActionLeaveCommunity  XUserFollowDeleteAllResponseAction = "leave_community"
)

// plannedCredits is the approved maximum. chargedCredits comes from the settled
// credit ledger. Pending or failed writes are not charged.
type XUserFollowDeleteAllResponseBilling struct {
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
func (r XUserFollowDeleteAllResponseBilling) RawJSON() string { return r.JSON.raw }
func (r *XUserFollowDeleteAllResponseBilling) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Exact follow-up an API client or agent should perform.
type XUserFollowDeleteAllResponseNextAction struct {
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
func (r XUserFollowDeleteAllResponseNextAction) RawJSON() string { return r.JSON.raw }
func (r *XUserFollowDeleteAllResponseNextAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Stable fingerprint and sanitized payload for replay checks.
type XUserFollowDeleteAllResponseRequest struct {
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
func (r XUserFollowDeleteAllResponseRequest) RawJSON() string { return r.JSON.raw }
func (r *XUserFollowDeleteAllResponseRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Confirmed result produced by the write, when available.
type XUserFollowDeleteAllResponseResult struct {
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
func (r XUserFollowDeleteAllResponseResult) RawJSON() string { return r.JSON.raw }
func (r *XUserFollowDeleteAllResponseResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type XUserFollowDeleteAllResponseStatus string

const (
	XUserFollowDeleteAllResponseStatusAccepted            XUserFollowDeleteAllResponseStatus = "accepted"
	XUserFollowDeleteAllResponseStatusDispatching         XUserFollowDeleteAllResponseStatus = "dispatching"
	XUserFollowDeleteAllResponseStatusPendingConfirmation XUserFollowDeleteAllResponseStatus = "pending_confirmation"
	XUserFollowDeleteAllResponseStatusSuccess             XUserFollowDeleteAllResponseStatus = "success"
	XUserFollowDeleteAllResponseStatusFailed              XUserFollowDeleteAllResponseStatus = "failed"
	XUserFollowDeleteAllResponseStatusExpired             XUserFollowDeleteAllResponseStatus = "expired"
)

// Existing X resource targeted by the write, when applicable.
type XUserFollowDeleteAllResponseTarget struct {
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
func (r XUserFollowDeleteAllResponseTarget) RawJSON() string { return r.JSON.raw }
func (r *XUserFollowDeleteAllResponseTarget) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type XUserFollowNewParams struct {
	// X account identifier (@username or account ID)
	Account        string `json:"account" api:"required"`
	IdempotencyKey string `header:"Idempotency-Key" api:"required" json:"-"`
	paramObj
}

func (r XUserFollowNewParams) MarshalJSON() (data []byte, err error) {
	type shadow XUserFollowNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *XUserFollowNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type XUserFollowDeleteAllParams struct {
	// X account identifier (@username or account ID)
	Account        string `json:"account" api:"required"`
	IdempotencyKey string `header:"Idempotency-Key" api:"required" json:"-"`
	paramObj
}

func (r XUserFollowDeleteAllParams) MarshalJSON() (data []byte, err error) {
	type shadow XUserFollowDeleteAllParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *XUserFollowDeleteAllParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
