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
	"github.com/Xquik-dev/x-twitter-scraper-go/shared/constant"
)

// XDmService contains methods and other services that help with interacting with
// the x-twitter-scraper API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewXDmService] method instead.
type XDmService struct {
	options []option.RequestOption
}

// NewXDmService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewXDmService(opts ...option.RequestOption) (r XDmService) {
	r = XDmService{}
	r.options = opts
	return
}

// Get DM conversation history
func (r *XDmService) GetHistory(ctx context.Context, userID string, query XDmGetHistoryParams, opts ...option.RequestOption) (res *XDmGetHistoryResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if userID == "" {
		err = errors.New("missing required userId parameter")
		return nil, err
	}
	path := fmt.Sprintf("x/dm/%s/history", url.PathEscape(userID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Send direct message
func (r *XDmService) Send(ctx context.Context, userID string, params XDmSendParams, opts ...option.RequestOption) (res *XDmSendResponse, err error) {
	if !param.IsOmitted(params.IdempotencyKey) {
		opts = append(opts, option.WithHeader("Idempotency-Key", fmt.Sprintf("%v", params.IdempotencyKey)))
	}
	opts = slices.Concat(r.options, opts)
	if userID == "" {
		err = errors.New("missing required userId parameter")
		return nil, err
	}
	path := fmt.Sprintf("x/dm/%s", url.PathEscape(userID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

type XDmGetHistoryResponse struct {
	HasNextPage bool                           `json:"has_next_page" api:"required"`
	Messages    []XDmGetHistoryResponseMessage `json:"messages" api:"required"`
	NextCursor  string                         `json:"next_cursor" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		HasNextPage respjson.Field
		Messages    respjson.Field
		NextCursor  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XDmGetHistoryResponse) RawJSON() string { return r.JSON.raw }
func (r *XDmGetHistoryResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type XDmGetHistoryResponseMessage struct {
	ID         string `json:"id" api:"required"`
	ReceiverID string `json:"receiverId" api:"required"`
	SenderID   string `json:"senderId" api:"required"`
	CreatedAt  string `json:"createdAt"`
	// URL of attached media (image, GIF, or video). Omitted when the message has no
	// media attachment.
	MediaURL string `json:"mediaUrl"`
	Text     string `json:"text"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ReceiverID  respjson.Field
		SenderID    respjson.Field
		CreatedAt   respjson.Field
		MediaURL    respjson.Field
		Text        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XDmGetHistoryResponseMessage) RawJSON() string { return r.JSON.raw }
func (r *XDmGetHistoryResponseMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Durable write lifecycle record. Poll statusUrl until terminal is true. Reusing
// the original Idempotency-Key returns this same record. Submit a new write only
// when safeToRetry is true, using a new key.
type XDmSendResponse struct {
	ID string `json:"id" api:"required"`
	// Connected account selected for the write.
	Account XDmSendResponseAccount `json:"account" api:"required"`
	// Any of "create_tweet", "delete_tweet", "like", "unlike", "retweet", "unretweet",
	// "follow", "unfollow", "remove_follower", "send_dm", "upload_media",
	// "update_profile", "update_avatar", "update_banner", "create_community",
	// "delete_community", "join_community", "leave_community".
	Action XDmSendResponseAction `json:"action" api:"required"`
	// plannedCredits is the approved maximum. chargedCredits comes from the settled
	// credit ledger. Pending or failed writes are not charged.
	Billing        XDmSendResponseBilling `json:"billing" api:"required"`
	Charged        bool                   `json:"charged" api:"required"`
	ChargedCredits string                 `json:"chargedCredits" api:"required"`
	// Exact follow-up an API client or agent should perform.
	NextAction  XDmSendResponseNextAction `json:"nextAction" api:"required"`
	Object      constant.XWriteAction     `json:"object" default:"x_write_action"`
	PollAfterMs int64                     `json:"pollAfterMs" api:"required"`
	// Stable fingerprint and sanitized payload for replay checks.
	Request XDmSendResponseRequest `json:"request" api:"required"`
	// Confirmed result produced by the write, when available.
	Result XDmSendResponseResult `json:"result" api:"required"`
	// True only when a new attempt can reasonably succeed.
	Retryable bool `json:"retryable" api:"required"`
	// True only when no write was dispatched and a new idempotency key may be used.
	SafeToRetry    bool `json:"safeToRetry" api:"required"`
	SendDispatched bool `json:"sendDispatched" api:"required"`
	// Any of "accepted", "dispatching", "pending_confirmation", "success", "failed",
	// "expired".
	Status    XDmSendResponseStatus `json:"status" api:"required"`
	StatusURL string                `json:"statusUrl" api:"required"`
	Success   bool                  `json:"success" api:"required"`
	// Existing X resource targeted by the write, when applicable.
	Target        XDmSendResponseTarget `json:"target" api:"required"`
	TargetID      string                `json:"targetId" api:"required"`
	Terminal      bool                  `json:"terminal" api:"required"`
	WriteActionID string                `json:"writeActionId" api:"required"`
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
func (r XDmSendResponse) RawJSON() string { return r.JSON.raw }
func (r *XDmSendResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Connected account selected for the write.
type XDmSendResponseAccount struct {
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
func (r XDmSendResponseAccount) RawJSON() string { return r.JSON.raw }
func (r *XDmSendResponseAccount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type XDmSendResponseAction string

const (
	XDmSendResponseActionCreateTweet     XDmSendResponseAction = "create_tweet"
	XDmSendResponseActionDeleteTweet     XDmSendResponseAction = "delete_tweet"
	XDmSendResponseActionLike            XDmSendResponseAction = "like"
	XDmSendResponseActionUnlike          XDmSendResponseAction = "unlike"
	XDmSendResponseActionRetweet         XDmSendResponseAction = "retweet"
	XDmSendResponseActionUnretweet       XDmSendResponseAction = "unretweet"
	XDmSendResponseActionFollow          XDmSendResponseAction = "follow"
	XDmSendResponseActionUnfollow        XDmSendResponseAction = "unfollow"
	XDmSendResponseActionRemoveFollower  XDmSendResponseAction = "remove_follower"
	XDmSendResponseActionSendDm          XDmSendResponseAction = "send_dm"
	XDmSendResponseActionUploadMedia     XDmSendResponseAction = "upload_media"
	XDmSendResponseActionUpdateProfile   XDmSendResponseAction = "update_profile"
	XDmSendResponseActionUpdateAvatar    XDmSendResponseAction = "update_avatar"
	XDmSendResponseActionUpdateBanner    XDmSendResponseAction = "update_banner"
	XDmSendResponseActionCreateCommunity XDmSendResponseAction = "create_community"
	XDmSendResponseActionDeleteCommunity XDmSendResponseAction = "delete_community"
	XDmSendResponseActionJoinCommunity   XDmSendResponseAction = "join_community"
	XDmSendResponseActionLeaveCommunity  XDmSendResponseAction = "leave_community"
)

// plannedCredits is the approved maximum. chargedCredits comes from the settled
// credit ledger. Pending or failed writes are not charged.
type XDmSendResponseBilling struct {
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
func (r XDmSendResponseBilling) RawJSON() string { return r.JSON.raw }
func (r *XDmSendResponseBilling) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Exact follow-up an API client or agent should perform.
type XDmSendResponseNextAction struct {
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
func (r XDmSendResponseNextAction) RawJSON() string { return r.JSON.raw }
func (r *XDmSendResponseNextAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Stable fingerprint and sanitized payload for replay checks.
type XDmSendResponseRequest struct {
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
func (r XDmSendResponseRequest) RawJSON() string { return r.JSON.raw }
func (r *XDmSendResponseRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Confirmed result produced by the write, when available.
type XDmSendResponseResult struct {
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
func (r XDmSendResponseResult) RawJSON() string { return r.JSON.raw }
func (r *XDmSendResponseResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type XDmSendResponseStatus string

const (
	XDmSendResponseStatusAccepted            XDmSendResponseStatus = "accepted"
	XDmSendResponseStatusDispatching         XDmSendResponseStatus = "dispatching"
	XDmSendResponseStatusPendingConfirmation XDmSendResponseStatus = "pending_confirmation"
	XDmSendResponseStatusSuccess             XDmSendResponseStatus = "success"
	XDmSendResponseStatusFailed              XDmSendResponseStatus = "failed"
	XDmSendResponseStatusExpired             XDmSendResponseStatus = "expired"
)

// Existing X resource targeted by the write, when applicable.
type XDmSendResponseTarget struct {
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
func (r XDmSendResponseTarget) RawJSON() string { return r.JSON.raw }
func (r *XDmSendResponseTarget) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type XDmGetHistoryParams struct {
	// X handle (without the `@` prefix) of the connected X account used to read the
	// conversation. The account must be a participant in the conversation.
	Account string `query:"account" api:"required" json:"-"`
	// Pagination cursor for DM history
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Legacy pagination cursor (backward compat)
	MaxID param.Opt[string] `query:"maxId,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [XDmGetHistoryParams]'s query parameters as `url.Values`.
func (r XDmGetHistoryParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type XDmSendParams struct {
	// X account (@username or ID) sending the DM
	Account        string `json:"account" api:"required"`
	Text           string `json:"text" api:"required"`
	IdempotencyKey string `header:"Idempotency-Key" api:"required" json:"-"`
	// Optional array containing exactly 1 uploaded media ID.
	MediaIDs []string `json:"media_ids,omitzero"`
	paramObj
}

func (r XDmSendParams) MarshalJSON() (data []byte, err error) {
	type shadow XDmSendParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *XDmSendParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
