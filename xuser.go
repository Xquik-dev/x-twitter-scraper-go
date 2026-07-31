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

// XUserService contains methods and other services that help with interacting with
// the x-twitter-scraper API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewXUserService] method instead.
type XUserService struct {
	options []option.RequestOption
	// X write actions (tweets, likes, follows, DMs)
	Follow XUserFollowService
}

// NewXUserService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewXUserService(opts ...option.RequestOption) (r XUserService) {
	r = XUserService{}
	r.options = opts
	r.Follow = NewXUserFollowService(opts...)
	return
}

// Get user profile with follower counts and verification
func (r *XUserService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *shared.UserProfile, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("x/users/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Remove follower
func (r *XUserService) RemoveFollower(ctx context.Context, id string, params XUserRemoveFollowerParams, opts ...option.RequestOption) (res *XUserRemoveFollowerResponse, err error) {
	if !param.IsOmitted(params.IdempotencyKey) {
		opts = append(opts, option.WithHeader("Idempotency-Key", fmt.Sprintf("%v", params.IdempotencyKey)))
	}
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("x/users/%s/remove-follower", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Look up multiple users by IDs in one call
func (r *XUserService) GetBatch(ctx context.Context, query XUserGetBatchParams, opts ...option.RequestOption) (res *XUserGetBatchResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "x/users/batch"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// List followers of a user
func (r *XUserService) GetFollowers(ctx context.Context, id string, query XUserGetFollowersParams, opts ...option.RequestOption) (res *shared.PaginatedUsers, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("x/users/%s/followers", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// List mutual followers between you and a user
func (r *XUserService) GetFollowersYouKnow(ctx context.Context, id string, query XUserGetFollowersYouKnowParams, opts ...option.RequestOption) (res *shared.PaginatedUsers, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("x/users/%s/followers-you-know", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// List accounts a user follows
func (r *XUserService) GetFollowing(ctx context.Context, id string, query XUserGetFollowingParams, opts ...option.RequestOption) (res *shared.PaginatedUsers, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("x/users/%s/following", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// List tweets liked by a user
func (r *XUserService) GetLikes(ctx context.Context, id string, query XUserGetLikesParams, opts ...option.RequestOption) (res *shared.PaginatedTweets, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("x/users/%s/likes", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// List media tweets posted by a user
func (r *XUserService) GetMedia(ctx context.Context, id string, query XUserGetMediaParams, opts ...option.RequestOption) (res *shared.PaginatedTweets, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("x/users/%s/media", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// List tweets mentioning a user
func (r *XUserService) GetMentions(ctx context.Context, id string, query XUserGetMentionsParams, opts ...option.RequestOption) (res *shared.PaginatedTweets, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("x/users/%s/mentions", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Returns the user's timeline with replies included by default.
func (r *XUserService) GetReplies(ctx context.Context, id string, query XUserGetRepliesParams, opts ...option.RequestOption) (res *shared.PaginatedTweets, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("x/users/%s/replies", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Search users by name or username
func (r *XUserService) GetSearch(ctx context.Context, query XUserGetSearchParams, opts ...option.RequestOption) (res *shared.PaginatedUsers, err error) {
	opts = slices.Concat(r.options, opts)
	path := "x/users/search"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// List recent tweets posted by a user
func (r *XUserService) GetTweets(ctx context.Context, id string, query XUserGetTweetsParams, opts ...option.RequestOption) (res *shared.PaginatedTweets, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("x/users/%s/tweets", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// List verified followers of a user
func (r *XUserService) GetVerifiedFollowers(ctx context.Context, id string, query XUserGetVerifiedFollowersParams, opts ...option.RequestOption) (res *shared.PaginatedUsers, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("x/users/%s/verified-followers", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Durable write lifecycle record. Poll statusUrl until terminal is true. Reusing
// the original Idempotency-Key returns this same record. Submit a new write only
// when safeToRetry is true, using a new key.
type XUserRemoveFollowerResponse struct {
	ID string `json:"id" api:"required"`
	// Connected account selected for the write.
	Account XUserRemoveFollowerResponseAccount `json:"account" api:"required"`
	// Any of "create_tweet", "delete_tweet", "like", "unlike", "retweet", "unretweet",
	// "follow", "unfollow", "remove_follower", "send_dm", "upload_media",
	// "update_profile", "update_avatar", "update_banner", "create_community",
	// "delete_community", "join_community", "leave_community".
	Action XUserRemoveFollowerResponseAction `json:"action" api:"required"`
	// plannedCredits is the approved maximum. chargedCredits comes from the settled
	// credit ledger. Pending or failed writes are not charged.
	Billing        XUserRemoveFollowerResponseBilling `json:"billing" api:"required"`
	Charged        bool                               `json:"charged" api:"required"`
	ChargedCredits string                             `json:"chargedCredits" api:"required"`
	// Exact follow-up an API client or agent should perform.
	NextAction  XUserRemoveFollowerResponseNextAction `json:"nextAction" api:"required"`
	Object      constant.XWriteAction                 `json:"object" default:"x_write_action"`
	PollAfterMs int64                                 `json:"pollAfterMs" api:"required"`
	// Stable fingerprint and sanitized payload for replay checks.
	Request XUserRemoveFollowerResponseRequest `json:"request" api:"required"`
	// Confirmed result produced by the write, when available.
	Result XUserRemoveFollowerResponseResult `json:"result" api:"required"`
	// True only when a new attempt can reasonably succeed.
	Retryable bool `json:"retryable" api:"required"`
	// True only when no write was dispatched and a new idempotency key may be used.
	SafeToRetry    bool `json:"safeToRetry" api:"required"`
	SendDispatched bool `json:"sendDispatched" api:"required"`
	// Any of "accepted", "dispatching", "pending_confirmation", "success", "failed",
	// "expired".
	Status    XUserRemoveFollowerResponseStatus `json:"status" api:"required"`
	StatusURL string                            `json:"statusUrl" api:"required"`
	Success   bool                              `json:"success" api:"required"`
	// Existing X resource targeted by the write, when applicable.
	Target        XUserRemoveFollowerResponseTarget `json:"target" api:"required"`
	TargetID      string                            `json:"targetId" api:"required"`
	Terminal      bool                              `json:"terminal" api:"required"`
	WriteActionID string                            `json:"writeActionId" api:"required"`
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
func (r XUserRemoveFollowerResponse) RawJSON() string { return r.JSON.raw }
func (r *XUserRemoveFollowerResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Connected account selected for the write.
type XUserRemoveFollowerResponseAccount struct {
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
func (r XUserRemoveFollowerResponseAccount) RawJSON() string { return r.JSON.raw }
func (r *XUserRemoveFollowerResponseAccount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type XUserRemoveFollowerResponseAction string

const (
	XUserRemoveFollowerResponseActionCreateTweet     XUserRemoveFollowerResponseAction = "create_tweet"
	XUserRemoveFollowerResponseActionDeleteTweet     XUserRemoveFollowerResponseAction = "delete_tweet"
	XUserRemoveFollowerResponseActionLike            XUserRemoveFollowerResponseAction = "like"
	XUserRemoveFollowerResponseActionUnlike          XUserRemoveFollowerResponseAction = "unlike"
	XUserRemoveFollowerResponseActionRetweet         XUserRemoveFollowerResponseAction = "retweet"
	XUserRemoveFollowerResponseActionUnretweet       XUserRemoveFollowerResponseAction = "unretweet"
	XUserRemoveFollowerResponseActionFollow          XUserRemoveFollowerResponseAction = "follow"
	XUserRemoveFollowerResponseActionUnfollow        XUserRemoveFollowerResponseAction = "unfollow"
	XUserRemoveFollowerResponseActionRemoveFollower  XUserRemoveFollowerResponseAction = "remove_follower"
	XUserRemoveFollowerResponseActionSendDm          XUserRemoveFollowerResponseAction = "send_dm"
	XUserRemoveFollowerResponseActionUploadMedia     XUserRemoveFollowerResponseAction = "upload_media"
	XUserRemoveFollowerResponseActionUpdateProfile   XUserRemoveFollowerResponseAction = "update_profile"
	XUserRemoveFollowerResponseActionUpdateAvatar    XUserRemoveFollowerResponseAction = "update_avatar"
	XUserRemoveFollowerResponseActionUpdateBanner    XUserRemoveFollowerResponseAction = "update_banner"
	XUserRemoveFollowerResponseActionCreateCommunity XUserRemoveFollowerResponseAction = "create_community"
	XUserRemoveFollowerResponseActionDeleteCommunity XUserRemoveFollowerResponseAction = "delete_community"
	XUserRemoveFollowerResponseActionJoinCommunity   XUserRemoveFollowerResponseAction = "join_community"
	XUserRemoveFollowerResponseActionLeaveCommunity  XUserRemoveFollowerResponseAction = "leave_community"
)

// plannedCredits is the approved maximum. chargedCredits comes from the settled
// credit ledger. Pending or failed writes are not charged.
type XUserRemoveFollowerResponseBilling struct {
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
func (r XUserRemoveFollowerResponseBilling) RawJSON() string { return r.JSON.raw }
func (r *XUserRemoveFollowerResponseBilling) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Exact follow-up an API client or agent should perform.
type XUserRemoveFollowerResponseNextAction struct {
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
func (r XUserRemoveFollowerResponseNextAction) RawJSON() string { return r.JSON.raw }
func (r *XUserRemoveFollowerResponseNextAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Stable fingerprint and sanitized payload for replay checks.
type XUserRemoveFollowerResponseRequest struct {
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
func (r XUserRemoveFollowerResponseRequest) RawJSON() string { return r.JSON.raw }
func (r *XUserRemoveFollowerResponseRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Confirmed result produced by the write, when available.
type XUserRemoveFollowerResponseResult struct {
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
func (r XUserRemoveFollowerResponseResult) RawJSON() string { return r.JSON.raw }
func (r *XUserRemoveFollowerResponseResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type XUserRemoveFollowerResponseStatus string

const (
	XUserRemoveFollowerResponseStatusAccepted            XUserRemoveFollowerResponseStatus = "accepted"
	XUserRemoveFollowerResponseStatusDispatching         XUserRemoveFollowerResponseStatus = "dispatching"
	XUserRemoveFollowerResponseStatusPendingConfirmation XUserRemoveFollowerResponseStatus = "pending_confirmation"
	XUserRemoveFollowerResponseStatusSuccess             XUserRemoveFollowerResponseStatus = "success"
	XUserRemoveFollowerResponseStatusFailed              XUserRemoveFollowerResponseStatus = "failed"
	XUserRemoveFollowerResponseStatusExpired             XUserRemoveFollowerResponseStatus = "expired"
)

// Existing X resource targeted by the write, when applicable.
type XUserRemoveFollowerResponseTarget struct {
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
func (r XUserRemoveFollowerResponseTarget) RawJSON() string { return r.JSON.raw }
func (r *XUserRemoveFollowerResponseTarget) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Batch user lookup results. Duplicate requested IDs are ignored while preserving
// first-seen order. unavailable_ids identifies processed IDs with no returned
// profile. unprocessed_ids identifies IDs skipped when available credits limit
// processing.
type XUserGetBatchResponse struct {
	// Batch lookups never paginate.
	HasNextPage bool `json:"has_next_page" api:"required"`
	// Empty because batch lookups never paginate.
	NextCursor string `json:"next_cursor" api:"required"`
	// Number of requested IDs included in the lookup.
	ProcessedCount int64 `json:"processed_count" api:"required"`
	// Number of unique IDs requested.
	RequestedCount int64 `json:"requested_count" api:"required"`
	// Number of user profiles returned and charged.
	ReturnedCount int64 `json:"returned_count" api:"required"`
	// Processed IDs with no returned profile, in first-seen request order.
	UnavailableIDs []string `json:"unavailable_ids" api:"required"`
	// Requested IDs skipped because available credits limited processing. Retry these
	// IDs after adding credits.
	UnprocessedIDs []string             `json:"unprocessed_ids" api:"required"`
	Users          []shared.UserProfile `json:"users" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		HasNextPage    respjson.Field
		NextCursor     respjson.Field
		ProcessedCount respjson.Field
		RequestedCount respjson.Field
		ReturnedCount  respjson.Field
		UnavailableIDs respjson.Field
		UnprocessedIDs respjson.Field
		Users          respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XUserGetBatchResponse) RawJSON() string { return r.JSON.raw }
func (r *XUserGetBatchResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type XUserRemoveFollowerParams struct {
	// X account identifier (@username or account ID)
	Account        string `json:"account" api:"required"`
	IdempotencyKey string `header:"Idempotency-Key" api:"required" json:"-"`
	paramObj
}

func (r XUserRemoveFollowerParams) MarshalJSON() (data []byte, err error) {
	type shadow XUserRemoveFollowerParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *XUserRemoveFollowerParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type XUserGetBatchParams struct {
	// Comma-separated numeric user IDs (1-100 values). Duplicate IDs are ignored while
	// preserving first-seen order.
	IDs string `query:"ids" api:"required" json:"-"`
	paramObj
}

// URLQuery serializes [XUserGetBatchParams]'s query parameters as `url.Values`.
func (r XUserGetBatchParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type XUserGetFollowersParams struct {
	// Legacy cursor alias. Prefer cursor.
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// Pagination cursor for followers list
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Legacy integer page size alias for following lists. Prefer pageSize.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Maximum user profiles requested from this page (20-200, default 200). The
	// response can contain fewer profiles because the source returned fewer or
	// remaining credits cover fewer results. Keep requesting next_cursor while
	// has_next_page is true. The deprecated limit and count aliases remain accepted.
	PageSize param.Opt[int64] `query:"pageSize,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [XUserGetFollowersParams]'s query parameters as
// `url.Values`.
func (r XUserGetFollowersParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type XUserGetFollowersYouKnowParams struct {
	// Pagination cursor for followers-you-know
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Maximum user profiles requested from this page (20-200, default 200). The
	// response can contain fewer profiles because the source returned fewer or
	// remaining credits cover fewer results. Keep requesting next_cursor while
	// has_next_page is true. The deprecated limit and count aliases remain accepted.
	PageSize param.Opt[int64] `query:"pageSize,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [XUserGetFollowersYouKnowParams]'s query parameters as
// `url.Values`.
func (r XUserGetFollowersYouKnowParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type XUserGetFollowingParams struct {
	// Deprecated following cursor alias. Prefer cursor.
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// Pagination cursor for following list
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Legacy page size alias. Prefer pageSize.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Maximum user profiles requested from this page (20-200, default 200). The
	// response can contain fewer profiles because the source returned fewer or
	// remaining credits cover fewer results. Keep requesting next_cursor while
	// has_next_page is true. The deprecated limit and count aliases remain accepted.
	PageSize param.Opt[int64] `query:"pageSize,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [XUserGetFollowingParams]'s query parameters as
// `url.Values`.
func (r XUserGetFollowingParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type XUserGetLikesParams struct {
	// Words or quoted phrases where any one can match. Separate with spaces, commas,
	// or lines.
	AnyWords param.Opt[string] `query:"anyWords,omitzero" json:"-"`
	// Cashtags separated by spaces, commas, or lines.
	Cashtags param.Opt[string] `query:"cashtags,omitzero" json:"-"`
	// Conversation ID filter.
	ConversationID param.Opt[string] `query:"conversationId,omitzero" json:"-"`
	// Pagination cursor for liked tweets
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Exact phrase to match.
	ExactPhrase param.Opt[string] `query:"exactPhrase,omitzero" json:"-"`
	// Words or quoted phrases to exclude. Separate with spaces, commas, or lines.
	ExcludeWords param.Opt[string] `query:"excludeWords,omitzero" json:"-"`
	// Filter by author username.
	FromUser param.Opt[string] `query:"fromUser,omitzero" json:"-"`
	// Hashtags separated by spaces, commas, or lines.
	Hashtags param.Opt[string] `query:"hashtags,omitzero" json:"-"`
	// Only replies to this tweet ID.
	InReplyToTweetID param.Opt[string] `query:"inReplyToTweetId,omitzero" json:"-"`
	// Language code filter, e.g. en or tr.
	Language param.Opt[string] `query:"language,omitzero" json:"-"`
	// Filter tweets mentioning a username.
	Mentioning param.Opt[string] `query:"mentioning,omitzero" json:"-"`
	// Minimum likes threshold.
	MinFaves param.Opt[int64] `query:"minFaves,omitzero" json:"-"`
	// Minimum quote count threshold.
	MinQuotes param.Opt[int64] `query:"minQuotes,omitzero" json:"-"`
	// Minimum replies threshold.
	MinReplies param.Opt[int64] `query:"minReplies,omitzero" json:"-"`
	// Minimum retweets threshold.
	MinRetweets param.Opt[int64] `query:"minRetweets,omitzero" json:"-"`
	// Maximum page items (1-100, default 20). Source, filters, or credits can reduce
	// results. Continue while has_next_page is true. Deprecated limit and count
	// aliases remain accepted.
	PageSize param.Opt[int64] `query:"pageSize,omitzero" json:"-"`
	// Only quotes of this tweet ID.
	QuotesOfTweetID param.Opt[string] `query:"quotesOfTweetId,omitzero" json:"-"`
	// Only retweets of this tweet ID.
	RetweetsOfTweetID param.Opt[string] `query:"retweetsOfTweetId,omitzero" json:"-"`
	// Start date in YYYY-MM-DD format.
	SinceDate param.Opt[time.Time] `query:"sinceDate,omitzero" format:"date" json:"-"`
	// Filter replies sent to a username.
	ToUser param.Opt[string] `query:"toUser,omitzero" json:"-"`
	// End date in YYYY-MM-DD format.
	UntilDate param.Opt[time.Time] `query:"untilDate,omitzero" format:"date" json:"-"`
	// URL substring or domain filter.
	URL param.Opt[string] `query:"url,omitzero" json:"-"`
	// Only return tweets from verified authors.
	VerifiedOnly param.Opt[bool] `query:"verifiedOnly,omitzero" json:"-"`
	// Filter by media type.
	//
	// Any of "images", "videos", "gifs", "media", "links", "none".
	MediaType XUserGetLikesParamsMediaType `query:"mediaType,omitzero" json:"-"`
	// Quote mode.
	//
	// Any of "include", "exclude", "only".
	Quotes XUserGetLikesParamsQuotes `query:"quotes,omitzero" json:"-"`
	// Reply mode.
	//
	// Any of "include", "exclude", "only".
	Replies XUserGetLikesParamsReplies `query:"replies,omitzero" json:"-"`
	// Retweet mode.
	//
	// Any of "include", "exclude", "only".
	Retweets XUserGetLikesParamsRetweets `query:"retweets,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [XUserGetLikesParams]'s query parameters as `url.Values`.
func (r XUserGetLikesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by media type.
type XUserGetLikesParamsMediaType string

const (
	XUserGetLikesParamsMediaTypeImages XUserGetLikesParamsMediaType = "images"
	XUserGetLikesParamsMediaTypeVideos XUserGetLikesParamsMediaType = "videos"
	XUserGetLikesParamsMediaTypeGifs   XUserGetLikesParamsMediaType = "gifs"
	XUserGetLikesParamsMediaTypeMedia  XUserGetLikesParamsMediaType = "media"
	XUserGetLikesParamsMediaTypeLinks  XUserGetLikesParamsMediaType = "links"
	XUserGetLikesParamsMediaTypeNone   XUserGetLikesParamsMediaType = "none"
)

// Quote mode.
type XUserGetLikesParamsQuotes string

const (
	XUserGetLikesParamsQuotesInclude XUserGetLikesParamsQuotes = "include"
	XUserGetLikesParamsQuotesExclude XUserGetLikesParamsQuotes = "exclude"
	XUserGetLikesParamsQuotesOnly    XUserGetLikesParamsQuotes = "only"
)

// Reply mode.
type XUserGetLikesParamsReplies string

const (
	XUserGetLikesParamsRepliesInclude XUserGetLikesParamsReplies = "include"
	XUserGetLikesParamsRepliesExclude XUserGetLikesParamsReplies = "exclude"
	XUserGetLikesParamsRepliesOnly    XUserGetLikesParamsReplies = "only"
)

// Retweet mode.
type XUserGetLikesParamsRetweets string

const (
	XUserGetLikesParamsRetweetsInclude XUserGetLikesParamsRetweets = "include"
	XUserGetLikesParamsRetweetsExclude XUserGetLikesParamsRetweets = "exclude"
	XUserGetLikesParamsRetweetsOnly    XUserGetLikesParamsRetweets = "only"
)

type XUserGetMediaParams struct {
	// Words or quoted phrases where any one can match. Separate with spaces, commas,
	// or lines.
	AnyWords param.Opt[string] `query:"anyWords,omitzero" json:"-"`
	// Cashtags separated by spaces, commas, or lines.
	Cashtags param.Opt[string] `query:"cashtags,omitzero" json:"-"`
	// Conversation ID filter.
	ConversationID param.Opt[string] `query:"conversationId,omitzero" json:"-"`
	// Pagination cursor for media tweets
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Exact phrase to match.
	ExactPhrase param.Opt[string] `query:"exactPhrase,omitzero" json:"-"`
	// Words or quoted phrases to exclude. Separate with spaces, commas, or lines.
	ExcludeWords param.Opt[string] `query:"excludeWords,omitzero" json:"-"`
	// Filter by author username.
	FromUser param.Opt[string] `query:"fromUser,omitzero" json:"-"`
	// Hashtags separated by spaces, commas, or lines.
	Hashtags param.Opt[string] `query:"hashtags,omitzero" json:"-"`
	// Only replies to this tweet ID.
	InReplyToTweetID param.Opt[string] `query:"inReplyToTweetId,omitzero" json:"-"`
	// Language code filter, e.g. en or tr.
	Language param.Opt[string] `query:"language,omitzero" json:"-"`
	// Filter tweets mentioning a username.
	Mentioning param.Opt[string] `query:"mentioning,omitzero" json:"-"`
	// Minimum likes threshold.
	MinFaves param.Opt[int64] `query:"minFaves,omitzero" json:"-"`
	// Minimum quote count threshold.
	MinQuotes param.Opt[int64] `query:"minQuotes,omitzero" json:"-"`
	// Minimum replies threshold.
	MinReplies param.Opt[int64] `query:"minReplies,omitzero" json:"-"`
	// Minimum retweets threshold.
	MinRetweets param.Opt[int64] `query:"minRetweets,omitzero" json:"-"`
	// Maximum page items (1-100, default 20). Source, filters, or credits can reduce
	// results. Continue while has_next_page is true. Deprecated limit and count
	// aliases remain accepted.
	PageSize param.Opt[int64] `query:"pageSize,omitzero" json:"-"`
	// Only quotes of this tweet ID.
	QuotesOfTweetID param.Opt[string] `query:"quotesOfTweetId,omitzero" json:"-"`
	// Only retweets of this tweet ID.
	RetweetsOfTweetID param.Opt[string] `query:"retweetsOfTweetId,omitzero" json:"-"`
	// Start date in YYYY-MM-DD format.
	SinceDate param.Opt[time.Time] `query:"sinceDate,omitzero" format:"date" json:"-"`
	// Filter replies sent to a username.
	ToUser param.Opt[string] `query:"toUser,omitzero" json:"-"`
	// End date in YYYY-MM-DD format.
	UntilDate param.Opt[time.Time] `query:"untilDate,omitzero" format:"date" json:"-"`
	// URL substring or domain filter.
	URL param.Opt[string] `query:"url,omitzero" json:"-"`
	// Only return tweets from verified authors.
	VerifiedOnly param.Opt[bool] `query:"verifiedOnly,omitzero" json:"-"`
	// Filter by media type.
	//
	// Any of "images", "videos", "gifs", "media", "links", "none".
	MediaType XUserGetMediaParamsMediaType `query:"mediaType,omitzero" json:"-"`
	// Quote mode.
	//
	// Any of "include", "exclude", "only".
	Quotes XUserGetMediaParamsQuotes `query:"quotes,omitzero" json:"-"`
	// Reply mode.
	//
	// Any of "include", "exclude", "only".
	Replies XUserGetMediaParamsReplies `query:"replies,omitzero" json:"-"`
	// Retweet mode.
	//
	// Any of "include", "exclude", "only".
	Retweets XUserGetMediaParamsRetweets `query:"retweets,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [XUserGetMediaParams]'s query parameters as `url.Values`.
func (r XUserGetMediaParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by media type.
type XUserGetMediaParamsMediaType string

const (
	XUserGetMediaParamsMediaTypeImages XUserGetMediaParamsMediaType = "images"
	XUserGetMediaParamsMediaTypeVideos XUserGetMediaParamsMediaType = "videos"
	XUserGetMediaParamsMediaTypeGifs   XUserGetMediaParamsMediaType = "gifs"
	XUserGetMediaParamsMediaTypeMedia  XUserGetMediaParamsMediaType = "media"
	XUserGetMediaParamsMediaTypeLinks  XUserGetMediaParamsMediaType = "links"
	XUserGetMediaParamsMediaTypeNone   XUserGetMediaParamsMediaType = "none"
)

// Quote mode.
type XUserGetMediaParamsQuotes string

const (
	XUserGetMediaParamsQuotesInclude XUserGetMediaParamsQuotes = "include"
	XUserGetMediaParamsQuotesExclude XUserGetMediaParamsQuotes = "exclude"
	XUserGetMediaParamsQuotesOnly    XUserGetMediaParamsQuotes = "only"
)

// Reply mode.
type XUserGetMediaParamsReplies string

const (
	XUserGetMediaParamsRepliesInclude XUserGetMediaParamsReplies = "include"
	XUserGetMediaParamsRepliesExclude XUserGetMediaParamsReplies = "exclude"
	XUserGetMediaParamsRepliesOnly    XUserGetMediaParamsReplies = "only"
)

// Retweet mode.
type XUserGetMediaParamsRetweets string

const (
	XUserGetMediaParamsRetweetsInclude XUserGetMediaParamsRetweets = "include"
	XUserGetMediaParamsRetweetsExclude XUserGetMediaParamsRetweets = "exclude"
	XUserGetMediaParamsRetweetsOnly    XUserGetMediaParamsRetweets = "only"
)

type XUserGetMentionsParams struct {
	// Words or quoted phrases where any one can match. Separate with spaces, commas,
	// or lines.
	AnyWords param.Opt[string] `query:"anyWords,omitzero" json:"-"`
	// Cashtags separated by spaces, commas, or lines.
	Cashtags param.Opt[string] `query:"cashtags,omitzero" json:"-"`
	// Conversation ID filter.
	ConversationID param.Opt[string] `query:"conversationId,omitzero" json:"-"`
	// Pagination cursor for mentions
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Exact phrase to match.
	ExactPhrase param.Opt[string] `query:"exactPhrase,omitzero" json:"-"`
	// Words or quoted phrases to exclude. Separate with spaces, commas, or lines.
	ExcludeWords param.Opt[string] `query:"excludeWords,omitzero" json:"-"`
	// Filter by author username.
	FromUser param.Opt[string] `query:"fromUser,omitzero" json:"-"`
	// Hashtags separated by spaces, commas, or lines.
	Hashtags param.Opt[string] `query:"hashtags,omitzero" json:"-"`
	// Only replies to this tweet ID.
	InReplyToTweetID param.Opt[string] `query:"inReplyToTweetId,omitzero" json:"-"`
	// Language code filter, e.g. en or tr.
	Language param.Opt[string] `query:"language,omitzero" json:"-"`
	// Filter tweets mentioning a username.
	Mentioning param.Opt[string] `query:"mentioning,omitzero" json:"-"`
	// Minimum likes threshold.
	MinFaves param.Opt[int64] `query:"minFaves,omitzero" json:"-"`
	// Minimum quote count threshold.
	MinQuotes param.Opt[int64] `query:"minQuotes,omitzero" json:"-"`
	// Minimum replies threshold.
	MinReplies param.Opt[int64] `query:"minReplies,omitzero" json:"-"`
	// Minimum retweets threshold.
	MinRetweets param.Opt[int64] `query:"minRetweets,omitzero" json:"-"`
	// Maximum page items (1-100, default 20). Source, filters, or credits can reduce
	// results. Continue while has_next_page is true. Deprecated limit and count
	// aliases remain accepted.
	PageSize param.Opt[int64] `query:"pageSize,omitzero" json:"-"`
	// Only quotes of this tweet ID.
	QuotesOfTweetID param.Opt[string] `query:"quotesOfTweetId,omitzero" json:"-"`
	// Only retweets of this tweet ID.
	RetweetsOfTweetID param.Opt[string] `query:"retweetsOfTweetId,omitzero" json:"-"`
	// Start date in YYYY-MM-DD format.
	SinceDate param.Opt[time.Time] `query:"sinceDate,omitzero" format:"date" json:"-"`
	// Unix timestamp - return mentions after this time
	SinceTime param.Opt[string] `query:"sinceTime,omitzero" json:"-"`
	// Filter replies sent to a username.
	ToUser param.Opt[string] `query:"toUser,omitzero" json:"-"`
	// End date in YYYY-MM-DD format.
	UntilDate param.Opt[time.Time] `query:"untilDate,omitzero" format:"date" json:"-"`
	// Unix timestamp - return mentions before this time
	UntilTime param.Opt[string] `query:"untilTime,omitzero" json:"-"`
	// URL substring or domain filter.
	URL param.Opt[string] `query:"url,omitzero" json:"-"`
	// Only return tweets from verified authors.
	VerifiedOnly param.Opt[bool] `query:"verifiedOnly,omitzero" json:"-"`
	// Filter by media type.
	//
	// Any of "images", "videos", "gifs", "media", "links", "none".
	MediaType XUserGetMentionsParamsMediaType `query:"mediaType,omitzero" json:"-"`
	// Quote mode.
	//
	// Any of "include", "exclude", "only".
	Quotes XUserGetMentionsParamsQuotes `query:"quotes,omitzero" json:"-"`
	// Reply mode.
	//
	// Any of "include", "exclude", "only".
	Replies XUserGetMentionsParamsReplies `query:"replies,omitzero" json:"-"`
	// Retweet mode.
	//
	// Any of "include", "exclude", "only".
	Retweets XUserGetMentionsParamsRetweets `query:"retweets,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [XUserGetMentionsParams]'s query parameters as `url.Values`.
func (r XUserGetMentionsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by media type.
type XUserGetMentionsParamsMediaType string

const (
	XUserGetMentionsParamsMediaTypeImages XUserGetMentionsParamsMediaType = "images"
	XUserGetMentionsParamsMediaTypeVideos XUserGetMentionsParamsMediaType = "videos"
	XUserGetMentionsParamsMediaTypeGifs   XUserGetMentionsParamsMediaType = "gifs"
	XUserGetMentionsParamsMediaTypeMedia  XUserGetMentionsParamsMediaType = "media"
	XUserGetMentionsParamsMediaTypeLinks  XUserGetMentionsParamsMediaType = "links"
	XUserGetMentionsParamsMediaTypeNone   XUserGetMentionsParamsMediaType = "none"
)

// Quote mode.
type XUserGetMentionsParamsQuotes string

const (
	XUserGetMentionsParamsQuotesInclude XUserGetMentionsParamsQuotes = "include"
	XUserGetMentionsParamsQuotesExclude XUserGetMentionsParamsQuotes = "exclude"
	XUserGetMentionsParamsQuotesOnly    XUserGetMentionsParamsQuotes = "only"
)

// Reply mode.
type XUserGetMentionsParamsReplies string

const (
	XUserGetMentionsParamsRepliesInclude XUserGetMentionsParamsReplies = "include"
	XUserGetMentionsParamsRepliesExclude XUserGetMentionsParamsReplies = "exclude"
	XUserGetMentionsParamsRepliesOnly    XUserGetMentionsParamsReplies = "only"
)

// Retweet mode.
type XUserGetMentionsParamsRetweets string

const (
	XUserGetMentionsParamsRetweetsInclude XUserGetMentionsParamsRetweets = "include"
	XUserGetMentionsParamsRetweetsExclude XUserGetMentionsParamsRetweets = "exclude"
	XUserGetMentionsParamsRetweetsOnly    XUserGetMentionsParamsRetweets = "only"
)

type XUserGetRepliesParams struct {
	// Words or quoted phrases where any one can match. Separate with spaces, commas,
	// or lines.
	AnyWords param.Opt[string] `query:"anyWords,omitzero" json:"-"`
	// Cashtags separated by spaces, commas, or lines.
	Cashtags param.Opt[string] `query:"cashtags,omitzero" json:"-"`
	// Conversation ID filter.
	ConversationID param.Opt[string] `query:"conversationId,omitzero" json:"-"`
	// Pagination cursor for user replies
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Exact phrase to match.
	ExactPhrase param.Opt[string] `query:"exactPhrase,omitzero" json:"-"`
	// Words or quoted phrases to exclude. Separate with spaces, commas, or lines.
	ExcludeWords param.Opt[string] `query:"excludeWords,omitzero" json:"-"`
	// Filter by author username.
	FromUser param.Opt[string] `query:"fromUser,omitzero" json:"-"`
	// Hashtags separated by spaces, commas, or lines.
	Hashtags param.Opt[string] `query:"hashtags,omitzero" json:"-"`
	// Include each reply's parent tweet.
	IncludeParentTweet param.Opt[bool] `query:"includeParentTweet,omitzero" json:"-"`
	// Only replies to this tweet ID.
	InReplyToTweetID param.Opt[string] `query:"inReplyToTweetId,omitzero" json:"-"`
	// Language code filter, e.g. en or tr.
	Language param.Opt[string] `query:"language,omitzero" json:"-"`
	// Filter tweets mentioning a username.
	Mentioning param.Opt[string] `query:"mentioning,omitzero" json:"-"`
	// Minimum likes threshold.
	MinFaves param.Opt[int64] `query:"minFaves,omitzero" json:"-"`
	// Minimum quote count threshold.
	MinQuotes param.Opt[int64] `query:"minQuotes,omitzero" json:"-"`
	// Minimum replies threshold.
	MinReplies param.Opt[int64] `query:"minReplies,omitzero" json:"-"`
	// Minimum retweets threshold.
	MinRetweets param.Opt[int64] `query:"minRetweets,omitzero" json:"-"`
	// Maximum page items (1-100, default 20). Source, filters, or credits can reduce
	// results. Continue while has_next_page is true. Deprecated limit and count
	// aliases remain accepted.
	PageSize param.Opt[int64] `query:"pageSize,omitzero" json:"-"`
	// Only quotes of this tweet ID.
	QuotesOfTweetID param.Opt[string] `query:"quotesOfTweetId,omitzero" json:"-"`
	// Only retweets of this tweet ID.
	RetweetsOfTweetID param.Opt[string] `query:"retweetsOfTweetId,omitzero" json:"-"`
	// Start date in YYYY-MM-DD format.
	SinceDate param.Opt[time.Time] `query:"sinceDate,omitzero" format:"date" json:"-"`
	// Filter replies sent to a username.
	ToUser param.Opt[string] `query:"toUser,omitzero" json:"-"`
	// End date in YYYY-MM-DD format.
	UntilDate param.Opt[time.Time] `query:"untilDate,omitzero" format:"date" json:"-"`
	// URL substring or domain filter.
	URL param.Opt[string] `query:"url,omitzero" json:"-"`
	// Only return tweets from verified authors.
	VerifiedOnly param.Opt[bool] `query:"verifiedOnly,omitzero" json:"-"`
	// Filter by media type.
	//
	// Any of "images", "videos", "gifs", "media", "links", "none".
	MediaType XUserGetRepliesParamsMediaType `query:"mediaType,omitzero" json:"-"`
	// Quote mode.
	//
	// Any of "include", "exclude", "only".
	Quotes XUserGetRepliesParamsQuotes `query:"quotes,omitzero" json:"-"`
	// Reply mode.
	//
	// Any of "include", "exclude", "only".
	Replies XUserGetRepliesParamsReplies `query:"replies,omitzero" json:"-"`
	// Retweet mode.
	//
	// Any of "include", "exclude", "only".
	Retweets XUserGetRepliesParamsRetweets `query:"retweets,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [XUserGetRepliesParams]'s query parameters as `url.Values`.
func (r XUserGetRepliesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by media type.
type XUserGetRepliesParamsMediaType string

const (
	XUserGetRepliesParamsMediaTypeImages XUserGetRepliesParamsMediaType = "images"
	XUserGetRepliesParamsMediaTypeVideos XUserGetRepliesParamsMediaType = "videos"
	XUserGetRepliesParamsMediaTypeGifs   XUserGetRepliesParamsMediaType = "gifs"
	XUserGetRepliesParamsMediaTypeMedia  XUserGetRepliesParamsMediaType = "media"
	XUserGetRepliesParamsMediaTypeLinks  XUserGetRepliesParamsMediaType = "links"
	XUserGetRepliesParamsMediaTypeNone   XUserGetRepliesParamsMediaType = "none"
)

// Quote mode.
type XUserGetRepliesParamsQuotes string

const (
	XUserGetRepliesParamsQuotesInclude XUserGetRepliesParamsQuotes = "include"
	XUserGetRepliesParamsQuotesExclude XUserGetRepliesParamsQuotes = "exclude"
	XUserGetRepliesParamsQuotesOnly    XUserGetRepliesParamsQuotes = "only"
)

// Reply mode.
type XUserGetRepliesParamsReplies string

const (
	XUserGetRepliesParamsRepliesInclude XUserGetRepliesParamsReplies = "include"
	XUserGetRepliesParamsRepliesExclude XUserGetRepliesParamsReplies = "exclude"
	XUserGetRepliesParamsRepliesOnly    XUserGetRepliesParamsReplies = "only"
)

// Retweet mode.
type XUserGetRepliesParamsRetweets string

const (
	XUserGetRepliesParamsRetweetsInclude XUserGetRepliesParamsRetweets = "include"
	XUserGetRepliesParamsRetweetsExclude XUserGetRepliesParamsRetweets = "exclude"
	XUserGetRepliesParamsRetweetsOnly    XUserGetRepliesParamsRetweets = "only"
)

type XUserGetSearchParams struct {
	// User search query
	Q string `query:"q" api:"required" json:"-"`
	// Pagination cursor for user search
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [XUserGetSearchParams]'s query parameters as `url.Values`.
func (r XUserGetSearchParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type XUserGetTweetsParams struct {
	// Words or quoted phrases where any one can match. Separate with spaces, commas,
	// or lines.
	AnyWords param.Opt[string] `query:"anyWords,omitzero" json:"-"`
	// Cashtags separated by spaces, commas, or lines.
	Cashtags param.Opt[string] `query:"cashtags,omitzero" json:"-"`
	// Conversation ID filter.
	ConversationID param.Opt[string] `query:"conversationId,omitzero" json:"-"`
	// Pagination cursor for user tweets
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Exact phrase to match.
	ExactPhrase param.Opt[string] `query:"exactPhrase,omitzero" json:"-"`
	// Words or quoted phrases to exclude. Separate with spaces, commas, or lines.
	ExcludeWords param.Opt[string] `query:"excludeWords,omitzero" json:"-"`
	// Filter by author username.
	FromUser param.Opt[string] `query:"fromUser,omitzero" json:"-"`
	// Hashtags separated by spaces, commas, or lines.
	Hashtags param.Opt[string] `query:"hashtags,omitzero" json:"-"`
	// Include parent tweet for replies
	IncludeParentTweet param.Opt[bool] `query:"includeParentTweet,omitzero" json:"-"`
	// Include reply tweets
	IncludeReplies param.Opt[bool] `query:"includeReplies,omitzero" json:"-"`
	// Only replies to this tweet ID.
	InReplyToTweetID param.Opt[string] `query:"inReplyToTweetId,omitzero" json:"-"`
	// Language code filter, e.g. en or tr.
	Language param.Opt[string] `query:"language,omitzero" json:"-"`
	// Filter tweets mentioning a username.
	Mentioning param.Opt[string] `query:"mentioning,omitzero" json:"-"`
	// Minimum likes threshold.
	MinFaves param.Opt[int64] `query:"minFaves,omitzero" json:"-"`
	// Minimum quote count threshold.
	MinQuotes param.Opt[int64] `query:"minQuotes,omitzero" json:"-"`
	// Minimum replies threshold.
	MinReplies param.Opt[int64] `query:"minReplies,omitzero" json:"-"`
	// Minimum retweets threshold.
	MinRetweets param.Opt[int64] `query:"minRetweets,omitzero" json:"-"`
	// Maximum page items (1-100, default 20). Source, filters, or credits can reduce
	// results. Continue while has_next_page is true. Deprecated limit and count
	// aliases remain accepted.
	PageSize param.Opt[int64] `query:"pageSize,omitzero" json:"-"`
	// Only quotes of this tweet ID.
	QuotesOfTweetID param.Opt[string] `query:"quotesOfTweetId,omitzero" json:"-"`
	// Only retweets of this tweet ID.
	RetweetsOfTweetID param.Opt[string] `query:"retweetsOfTweetId,omitzero" json:"-"`
	// Start date in YYYY-MM-DD format.
	SinceDate param.Opt[time.Time] `query:"sinceDate,omitzero" format:"date" json:"-"`
	// Filter replies sent to a username.
	ToUser param.Opt[string] `query:"toUser,omitzero" json:"-"`
	// End date in YYYY-MM-DD format.
	UntilDate param.Opt[time.Time] `query:"untilDate,omitzero" format:"date" json:"-"`
	// URL substring or domain filter.
	URL param.Opt[string] `query:"url,omitzero" json:"-"`
	// Only return tweets from verified authors.
	VerifiedOnly param.Opt[bool] `query:"verifiedOnly,omitzero" json:"-"`
	// Filter by media type.
	//
	// Any of "images", "videos", "gifs", "media", "links", "none".
	MediaType XUserGetTweetsParamsMediaType `query:"mediaType,omitzero" json:"-"`
	// Quote mode.
	//
	// Any of "include", "exclude", "only".
	Quotes XUserGetTweetsParamsQuotes `query:"quotes,omitzero" json:"-"`
	// Reply mode.
	//
	// Any of "include", "exclude", "only".
	Replies XUserGetTweetsParamsReplies `query:"replies,omitzero" json:"-"`
	// Retweet mode.
	//
	// Any of "include", "exclude", "only".
	Retweets XUserGetTweetsParamsRetweets `query:"retweets,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [XUserGetTweetsParams]'s query parameters as `url.Values`.
func (r XUserGetTweetsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by media type.
type XUserGetTweetsParamsMediaType string

const (
	XUserGetTweetsParamsMediaTypeImages XUserGetTweetsParamsMediaType = "images"
	XUserGetTweetsParamsMediaTypeVideos XUserGetTweetsParamsMediaType = "videos"
	XUserGetTweetsParamsMediaTypeGifs   XUserGetTweetsParamsMediaType = "gifs"
	XUserGetTweetsParamsMediaTypeMedia  XUserGetTweetsParamsMediaType = "media"
	XUserGetTweetsParamsMediaTypeLinks  XUserGetTweetsParamsMediaType = "links"
	XUserGetTweetsParamsMediaTypeNone   XUserGetTweetsParamsMediaType = "none"
)

// Quote mode.
type XUserGetTweetsParamsQuotes string

const (
	XUserGetTweetsParamsQuotesInclude XUserGetTweetsParamsQuotes = "include"
	XUserGetTweetsParamsQuotesExclude XUserGetTweetsParamsQuotes = "exclude"
	XUserGetTweetsParamsQuotesOnly    XUserGetTweetsParamsQuotes = "only"
)

// Reply mode.
type XUserGetTweetsParamsReplies string

const (
	XUserGetTweetsParamsRepliesInclude XUserGetTweetsParamsReplies = "include"
	XUserGetTweetsParamsRepliesExclude XUserGetTweetsParamsReplies = "exclude"
	XUserGetTweetsParamsRepliesOnly    XUserGetTweetsParamsReplies = "only"
)

// Retweet mode.
type XUserGetTweetsParamsRetweets string

const (
	XUserGetTweetsParamsRetweetsInclude XUserGetTweetsParamsRetweets = "include"
	XUserGetTweetsParamsRetweetsExclude XUserGetTweetsParamsRetweets = "exclude"
	XUserGetTweetsParamsRetweetsOnly    XUserGetTweetsParamsRetweets = "only"
)

type XUserGetVerifiedFollowersParams struct {
	// Pagination cursor for verified followers
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Maximum user profiles requested from this page (20-200, default 200). The
	// response can contain fewer profiles because the source returned fewer or
	// remaining credits cover fewer results. Keep requesting next_cursor while
	// has_next_page is true. The deprecated limit and count aliases remain accepted.
	PageSize param.Opt[int64] `query:"pageSize,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [XUserGetVerifiedFollowersParams]'s query parameters as
// `url.Values`.
func (r XUserGetVerifiedFollowersParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
