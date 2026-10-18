// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

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
// You should not instantiate this service directly, and instead use the
// [NewXCommunityService] method instead.
type XCommunityService struct {
	options []option.RequestOption
	// X write actions (tweets, likes, follows, DMs)
	Join XCommunityJoinService
	// X Community info, members, and tweets
	Tweets XCommunityTweetService
	// Find X Communities & read their info, members, posts & media
	Media XCommunityMediaService
}

// NewXCommunityService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewXCommunityService(opts ...option.RequestOption) (r XCommunityService) {
	r = XCommunityService{}
	r.options = opts
	r.Join = NewXCommunityJoinService(opts...)
	r.Tweets = NewXCommunityTweetService(opts...)
	r.Media = NewXCommunityMediaService(opts...)
	return
}

// Creates an X Community through a connected account.
func (r *XCommunityService) New(ctx context.Context, params XCommunityNewParams, opts ...option.RequestOption) (res *XCommunityNewResponse, err error) {
	if !param.IsOmitted(params.IdempotencyKey) {
		opts = append(opts, option.WithHeader("Idempotency-Key", fmt.Sprintf("%v", params.IdempotencyKey)))
	}
	opts = slices.Concat(r.options, opts)
	path := "x/communities"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Deletes an owned X Community through a connected account.
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

// Returns the communities X finds for a keyword, in X's order. Each community
// costs 1 credit. To search the tweets in Communities, use /x/communities/search.
func (r *XCommunityService) Find(ctx context.Context, query XCommunityFindParams, opts ...option.RequestOption) (res *shared.CommunityPage, err error) {
	opts = slices.Concat(r.options, opts)
	path := "x/communities/find"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Returns community activity for X's latest period. Includes likes, new members,
// posts, replies & unique posters. Each metric has a count and a changePercent.
// Each request costs 1 credit.
func (r *XCommunityService) GetAnalytics(ctx context.Context, id string, query XCommunityGetAnalyticsParams, opts ...option.RequestOption) (res *shared.CommunityAnalytics, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("x/communities/%s/analytics", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Returns public identity and membership counts for one community.
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

// Returns public member profiles for one community.
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

// Returns public moderator profiles for one community.
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

// Returns one tweet page, not community records.
func (r *XCommunityService) GetSearch(ctx context.Context, query XCommunityGetSearchParams, opts ...option.RequestOption) (res *shared.PaginatedTweets, err error) {
	opts = slices.Concat(r.options, opts)
	path := "x/communities/search"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Returns the communities X ranks highest under a topic or subtopic. Name it by
// topicId or by topic name from /x/communities/topics. Each community costs 1
// credit. A missing or unknown topic returns invalid_community_topic & lists the
// topic names.
func (r *XCommunityService) ListPopular(ctx context.Context, query XCommunityListPopularParams, opts ...option.RequestOption) (res *shared.CommunityPage, err error) {
	opts = slices.Concat(r.options, opts)
	path := "x/communities/popular"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Returns the communities X suggests, with their full public fields. X picks them
// per request, so 2 calls can differ. If X suggests none, Xquik asks up to 2 more
// times. An empty list is free. Each community costs 1 credit.
func (r *XCommunityService) ListSuggested(ctx context.Context, query XCommunityListSuggestedParams, opts ...option.RequestOption) (res *XCommunityListSuggestedResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "x/communities/suggested"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Returns the community topics with their subtopics. Pass an ID as topicId, or a
// name as topic, to /x/communities/popular. Each top-level topic costs 1 credit,
// its subtopics included.
func (r *XCommunityService) ListTopics(ctx context.Context, query XCommunityListTopicsParams, opts ...option.RequestOption) (res *XCommunityListTopicsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "x/communities/topics"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Finds a community's members by name or username, as the search box on X's
// members page does. X checks up to 10 accounts that match, & the answer holds the
// members among them, each with its role. Each member costs 1 credit. An answer
// without members is free.
func (r *XCommunityService) SearchMembers(ctx context.Context, id string, query XCommunitySearchMembersParams, opts ...option.RequestOption) (res *shared.PaginatedUsers, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("x/communities/%s/members/search", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Durable write record. Poll statusUrl until terminal is true. Reusing its
// Idempotency-Key returns this record. Create another action only when safeToRetry
// is true.
type XCommunityNewResponse struct {
	ID string `json:"id" api:"required"`
	// Connected account selected for the write.
	Account XCommunityNewResponseAccount `json:"account" api:"required"`
	// Any of "create_tweet", "delete_tweet", "like", "unlike", "retweet", "unretweet",
	// "follow", "unfollow", "remove_follower", "send_dm", "delete_dm", "upload_media",
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
	XCommunityNewResponseActionDeleteDm        XCommunityNewResponseAction = "delete_dm"
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

// Durable write record. Poll statusUrl until terminal is true. Reusing its
// Idempotency-Key returns this record. Create another action only when safeToRetry
// is true.
type XCommunityDeleteResponse struct {
	ID string `json:"id" api:"required"`
	// Connected account selected for the write.
	Account XCommunityDeleteResponseAccount `json:"account" api:"required"`
	// Any of "create_tweet", "delete_tweet", "like", "unlike", "retweet", "unretweet",
	// "follow", "unfollow", "remove_follower", "send_dm", "delete_dm", "upload_media",
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
	XCommunityDeleteResponseActionDeleteDm        XCommunityDeleteResponseAction = "delete_dm"
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
	// Public X profile.
	Admin shared.UserProfile `json:"admin"`
	// The banner X shows for a community, with its size & colors.
	Banner shared.CommunityBanner `json:"banner"`
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
	IsNsfw bool `json:"is_nsfw" api:"required"`
	// Join policy (open or restricted)
	JoinPolicy string `json:"join_policy"`
	// Total member count
	MemberCount int64 `json:"member_count" api:"required"`
	// The members X previews beside the member count. A keyword or topic result gives
	// each member's ID & picture only.
	MembersFacepile []shared.CommunityMembersFacepileUnion `json:"membersFacepile"`
	// Total moderator count
	ModeratorCount int64 `json:"moderator_count"`
	// Display name of the community
	Name string `json:"name" api:"required"`
	// Primary topic
	PrimaryTopic XCommunityGetInfoResponseCommunityPrimaryTopic `json:"primary_topic"`
	// What the community asks a user who joins
	Question string `json:"question"`
	// Community rules
	Rules []XCommunityGetInfoResponseCommunityRule `json:"rules"`
	// Tags X finds the community by
	SearchTags []string `json:"searchTags"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		Admin           respjson.Field
		Banner          respjson.Field
		BannerURL       respjson.Field
		CreatedAt       respjson.Field
		Creator         respjson.Field
		Description     respjson.Field
		InvitesPolicy   respjson.Field
		IsNsfw          respjson.Field
		JoinPolicy      respjson.Field
		MemberCount     respjson.Field
		MembersFacepile respjson.Field
		ModeratorCount  respjson.Field
		Name            respjson.Field
		PrimaryTopic    respjson.Field
		Question        respjson.Field
		Rules           respjson.Field
		SearchTags      respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
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
	// Null when X shows no label.
	AccountBasedIn shared.UserProfileAccountBasedIn `json:"accountBasedIn" api:"nullable"`
	// False when X notes accountBasedIn may be inaccurate.
	AccountBasedInAccurate bool `json:"accountBasedInAccurate"`
	// True when X withheld accountBasedIn from this read. Retry later.
	AccountBasedInUnavailable bool `json:"accountBasedInUnavailable"`
	// X Ads services the account uses, such as `analytics`, `media_studio` or `smb`.
	// Empty for an account without ads.
	AdvertiserAccountServiceLevels []string `json:"advertiserAccountServiceLevels"`
	// The account's X Ads type. `none` for an account without ads, or
	// `promotable_user` for one that may run them.
	AdvertiserAccountType string `json:"advertiserAccountType"`
	// Organization affiliation label shown on an X profile.
	AffiliatesHighlightedLabel shared.UserProfileAffiliatesHighlightedLabel `json:"affiliatesHighlightedLabel"`
	// Handle of the organization X shows the account is affiliated with, from its
	// About page.
	AffiliateUsername string `json:"affiliateUsername"`
	AutomatedBy       string `json:"automatedBy"`
	// The parts of the birthday the account made public.
	Birthdate shared.UserProfileBirthdate `json:"birthdate"`
	// Whole number from 0.
	BusinessAccountAffiliatesCount int64 `json:"businessAccountAffiliatesCount"`
	// Whether X treats the profile as a business profile. `none` or `enabled`.
	BusinessProfileState string `json:"businessProfileState"`
	// Community role when returned by community member reads
	CommunityRole string `json:"communityRole"`
	// How the account first connected to X.
	ConnectedVia string `json:"connectedVia"`
	// False when X notes that country may be inaccurate.
	ConnectedViaAccurate bool `json:"connectedViaAccurate"`
	// Emails, phone numbers & links the account publishes on its own profile. Omitted
	// when it publishes none.
	ContactInfo  shared.UserProfileContactInfo `json:"contactInfo"`
	CoverPicture string                        `json:"coverPicture"`
	// Account creation time.
	CreatedAt string `json:"createdAt"`
	// Whole number from 0.
	CreatorSubscriptionsCount int64 `json:"creatorSubscriptionsCount"`
	// Whether the profile keeps X's default theme.
	DefaultProfile bool `json:"defaultProfile"`
	// Whether the profile keeps X's default avatar.
	DefaultProfileImage bool   `json:"defaultProfileImage"`
	Description         string `json:"description"`
	// First email the account publishes on its profile, lowercased. Omitted when it
	// publishes none.
	Email string `json:"email"`
	// Followers X counts as fast followers.
	FastFollowersCount int64 `json:"fastFollowersCount"`
	FavouritesCount    int64 `json:"favouritesCount"`
	// Omitted when X sends no count, as for the account that first posted reused
	// media.
	Followers int64 `json:"followers"`
	Following int64 `json:"following"`
	// Whether the profile may add a location to posts.
	GeoEnabled bool `json:"geoEnabled"`
	// Public profile bio translation returned by X
	GrokTranslatedBio  map[string]any `json:"grokTranslatedBio"`
	HasCustomTimelines bool           `json:"hasCustomTimelines"`
	// Whether the profile filled X's extended fields, such as a birth date. True or
	// false.
	HasExtendedProfile              bool `json:"hasExtendedProfile"`
	HasGraduatedAccess              bool `json:"hasGraduatedAccess"`
	HasHiddenSubscriptionsOnProfile bool `json:"hasHiddenSubscriptionsOnProfile"`
	// Profile highlight availability and count metadata.
	HighlightsInfo shared.UserProfileHighlightsInfo `json:"highlightsInfo"`
	// Identity verification X shows on the profile & its About page.
	IdentityVerification shared.UserProfileIdentityVerification `json:"identityVerification"`
	IsAutomated          bool                                   `json:"isAutomated"`
	// Whether X shows a blue verification badge
	IsBlueVerified        bool `json:"isBlueVerified"`
	IsProfileTranslatable bool `json:"isProfileTranslatable"`
	// Whether the account turned on X's translation of its posts. True or false.
	IsTranslationEnabled bool `json:"isTranslationEnabled"`
	IsTranslator         bool `json:"isTranslator"`
	// Whether X marks the profile as verified
	IsVerified bool `json:"isVerified"`
	// Public lists that include the profile.
	ListedCount int64 `json:"listedCount"`
	// Account owner's public profile location text
	Location   string `json:"location"`
	MediaCount int64  `json:"mediaCount"`
	// Followers minus fast followers.
	NormalFollowersCount     int64    `json:"normalFollowersCount"`
	ParodyCommentaryFanLabel string   `json:"parodyCommentaryFanLabel"`
	PinnedTweetIDs           []string `json:"pinnedTweetIds"`
	PossiblySensitive        bool     `json:"possiblySensitive"`
	// Professional metadata with category display settings
	Professional map[string]any `json:"professional"`
	// Structured profile bio with entity annotations
	ProfileBio map[string]any `json:"profile_bio"`
	// Original X profile banner field when available
	ProfileBannerURL           string `json:"profileBannerUrl"`
	ProfileDescriptionLanguage string `json:"profileDescriptionLanguage"`
	ProfileImageShape          string `json:"profileImageShape"`
	ProfileInterstitialType    string `json:"profileInterstitialType"`
	// Profile picture at 48 px, as X sends it.
	ProfilePicture string `json:"profilePicture"`
	// The same profile picture at full size, as the account uploaded it. Download this
	// one for the image.
	ProfilePictureOriginal string `json:"profilePictureOriginal"`
	ProfileSortEnabled     bool   `json:"profileSortEnabled"`
	ProfileTranslatorType  string `json:"profileTranslatorType"`
	// Whether the profile protects its posts
	Protected bool `json:"protected"`
	// Handle x.com sends this profile's visitors to, such as the new handle of a
	// renamed account whose old handle now holds a placeholder. Missing when x.com
	// shows the profile itself.
	RedirectsToUsername string `json:"redirectsToUsername"`
	// UTC repost time with includeRetweetTimestamp. Null if the newest profile page
	// has no match or lookup failed. Otherwise omitted.
	RetweetedAt   time.Time `json:"retweetedAt" api:"nullable" format:"date-time"`
	StatusesCount int64     `json:"statusesCount"`
	// Accounts subscribed to the profile's paid Subscriptions. 0 when it has none.
	SubscribersCount    int64 `json:"subscribersCount"`
	SuperFollowEligible bool  `json:"superFollowEligible"`
	// Whether X marks the subscription profile as active.
	SuperFollowsUserProfileActive bool `json:"superFollowsUserProfileActive"`
	// Public payment and creator-support handles shown on X.
	TipJar            shared.UserProfileTipJar `json:"tipJar"`
	Unavailable       bool                     `json:"unavailable"`
	UnavailableReason string                   `json:"unavailableReason"`
	URL               string                   `json:"url"`
	// How often the account changed its username.
	UsernameChanges shared.UserProfileUsernameChanges `json:"usernameChanges"`
	// X's text on why the profile is verified, with its links.
	VerificationReason  shared.UserProfileVerificationReason `json:"verificationReason"`
	Verified            bool                                 `json:"verified"`
	Name                string                               `json:"name"`
	VerifiedType        string                               `json:"verifiedType"`
	WithheldInCountries []string                             `json:"withheldInCountries"`
	// Whether X withholds a post or user
	WithheldScope string `json:"withheldScope"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                              respjson.Field
		Username                        respjson.Field
		AccountBasedIn                  respjson.Field
		AccountBasedInAccurate          respjson.Field
		AccountBasedInUnavailable       respjson.Field
		AdvertiserAccountServiceLevels  respjson.Field
		AdvertiserAccountType           respjson.Field
		AffiliatesHighlightedLabel      respjson.Field
		AffiliateUsername               respjson.Field
		AutomatedBy                     respjson.Field
		Birthdate                       respjson.Field
		BusinessAccountAffiliatesCount  respjson.Field
		BusinessProfileState            respjson.Field
		CommunityRole                   respjson.Field
		ConnectedVia                    respjson.Field
		ConnectedViaAccurate            respjson.Field
		ContactInfo                     respjson.Field
		CoverPicture                    respjson.Field
		CreatedAt                       respjson.Field
		CreatorSubscriptionsCount       respjson.Field
		DefaultProfile                  respjson.Field
		DefaultProfileImage             respjson.Field
		Description                     respjson.Field
		Email                           respjson.Field
		FastFollowersCount              respjson.Field
		FavouritesCount                 respjson.Field
		Followers                       respjson.Field
		Following                       respjson.Field
		GeoEnabled                      respjson.Field
		GrokTranslatedBio               respjson.Field
		HasCustomTimelines              respjson.Field
		HasExtendedProfile              respjson.Field
		HasGraduatedAccess              respjson.Field
		HasHiddenSubscriptionsOnProfile respjson.Field
		HighlightsInfo                  respjson.Field
		IdentityVerification            respjson.Field
		IsAutomated                     respjson.Field
		IsBlueVerified                  respjson.Field
		IsProfileTranslatable           respjson.Field
		IsTranslationEnabled            respjson.Field
		IsTranslator                    respjson.Field
		IsVerified                      respjson.Field
		ListedCount                     respjson.Field
		Location                        respjson.Field
		MediaCount                      respjson.Field
		NormalFollowersCount            respjson.Field
		ParodyCommentaryFanLabel        respjson.Field
		PinnedTweetIDs                  respjson.Field
		PossiblySensitive               respjson.Field
		Professional                    respjson.Field
		ProfileBio                      respjson.Field
		ProfileBannerURL                respjson.Field
		ProfileDescriptionLanguage      respjson.Field
		ProfileImageShape               respjson.Field
		ProfileInterstitialType         respjson.Field
		ProfilePicture                  respjson.Field
		ProfilePictureOriginal          respjson.Field
		ProfileSortEnabled              respjson.Field
		ProfileTranslatorType           respjson.Field
		Protected                       respjson.Field
		RedirectsToUsername             respjson.Field
		RetweetedAt                     respjson.Field
		StatusesCount                   respjson.Field
		SubscribersCount                respjson.Field
		SuperFollowEligible             respjson.Field
		SuperFollowsUserProfileActive   respjson.Field
		TipJar                          respjson.Field
		Unavailable                     respjson.Field
		UnavailableReason               respjson.Field
		URL                             respjson.Field
		UsernameChanges                 respjson.Field
		VerificationReason              respjson.Field
		Verified                        respjson.Field
		Name                            respjson.Field
		VerifiedType                    respjson.Field
		WithheldInCountries             respjson.Field
		WithheldScope                   respjson.Field
		ExtraFields                     map[string]respjson.Field
		raw                             string
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
	Name string `json:"name" api:"required"`
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

type XCommunityListSuggestedResponse struct {
	// Suggested communities in X's order
	Communities []shared.Community `json:"communities" api:"required"`
	// Communities returned & charged
	Count int64 `json:"count" api:"required"`
	// What to do next. Present only when X suggests none.
	Message string `json:"message"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Communities respjson.Field
		Count       respjson.Field
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XCommunityListSuggestedResponse) RawJSON() string { return r.JSON.raw }
func (r *XCommunityListSuggestedResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type XCommunityListTopicsResponse struct {
	// Top-level topics returned & charged
	Count int64 `json:"count" api:"required"`
	// Topics in X's order
	Topics []shared.CommunityTopic `json:"topics" api:"required"`
	// What to do next. Present only when X lists no topic.
	Message string `json:"message"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Topics      respjson.Field
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XCommunityListTopicsResponse) RawJSON() string { return r.JSON.raw }
func (r *XCommunityListTopicsResponse) UnmarshalJSON(data []byte) error {
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

type XCommunityGetAnalyticsParams struct {
	// Fields to keep in each row, comma-separated. Use dots for nested fields, such as
	// author.username. A list keeps the field in each item. Case & underscores do not
	// matter. A row that lacks a field leaves it out. Cursors, counts & diagnostics
	// always stay. An unknown name answers 400 unknown_field with the valid names,
	// before any read. The price stays the same. SDK row types then hold only the
	// named fields.
	Fields param.Opt[string] `query:"fields,omitzero" json:"-"`
	// Fields to drop from each row, comma-separated. It takes the same names as
	// fields. With fields, it drops from the kept fields.
	Omit param.Opt[string] `query:"omit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [XCommunityGetAnalyticsParams]'s query parameters as
// `url.Values`.
func (r XCommunityGetAnalyticsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type XCommunityGetMembersParams struct {
	// Match any comma-separated or line-separated bio term, ignoring case.
	BioContains param.Opt[string] `query:"bioContains,omitzero" json:"-"`
	// Only return profiles X marks as blue verified. X marks many business &
	// government accounts blue too. Use verifiedType for 1 badge type.
	BlueVerifiedOnly param.Opt[bool] `query:"blueVerifiedOnly,omitzero" json:"-"`
	// Only return profiles that publish an email outside free webmail, such as Gmail.
	BusinessEmailOnly param.Opt[bool] `query:"businessEmailOnly,omitzero" json:"-"`
	// Pagination cursor for collection results.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Only return profiles with an email at one of these domains. Separate domains
	// with commas. Subdomains match.
	EmailDomains param.Opt[string] `query:"emailDomains,omitzero" json:"-"`
	// Fields to keep in each row, comma-separated. Use dots for nested fields, such as
	// author.username. A list keeps the field in each item. Case & underscores do not
	// matter. A row that lacks a field leaves it out. Cursors, counts & diagnostics
	// always stay. An unknown name answers 400 unknown_field with the valid names,
	// before any read. The price stays the same. SDK row types then hold only the
	// named fields.
	Fields param.Opt[string] `query:"fields,omitzero" json:"-"`
	// Only return profiles that publish an email.
	HasEmail param.Opt[bool] `query:"hasEmail,omitzero" json:"-"`
	// Only return profiles with a location.
	HasLocation param.Opt[bool] `query:"hasLocation,omitzero" json:"-"`
	// Only return profiles with a website.
	HasWebsite param.Opt[bool] `query:"hasWebsite,omitzero" json:"-"`
	// Match a location substring, ignoring case.
	LocationContains param.Opt[string] `query:"locationContains,omitzero" json:"-"`
	// Maximum follower count. Missing counts pass this maximum.
	MaxFollowers param.Opt[int64] `query:"maxFollowers,omitzero" json:"-"`
	// Profiles may follow at most this many accounts.
	MaxFollowing param.Opt[int64] `query:"maxFollowing,omitzero" json:"-"`
	// Maximum post count. maxPosts is also accepted.
	MaxStatuses param.Opt[int64] `query:"maxStatuses,omitzero" json:"-"`
	// Minimum account age in whole days.
	MinAccountAgeDays param.Opt[int64] `query:"minAccountAgeDays,omitzero" json:"-"`
	// Minimum follower count. Filtering happens before billing.
	MinFollowers param.Opt[int64] `query:"minFollowers,omitzero" json:"-"`
	// Profiles must follow at least this many accounts.
	MinFollowing param.Opt[int64] `query:"minFollowing,omitzero" json:"-"`
	// Minimum post count. minPosts is also accepted.
	MinStatuses param.Opt[int64] `query:"minStatuses,omitzero" json:"-"`
	// Fields to drop from each row, comma-separated. It takes the same names as
	// fields. With fields, it drops from the kept fields.
	Omit param.Opt[string] `query:"omit,omitzero" json:"-"`
	// Maximum user profiles per page (1-200, default 20).
	PageSize param.Opt[int64] `query:"pageSize,omitzero" json:"-"`
	// Match a username substring, ignoring case.
	UsernameContains param.Opt[string] `query:"usernameContains,omitzero" json:"-"`
	// Only return verified profiles.
	VerifiedOnly param.Opt[bool] `query:"verifiedOnly,omitzero" json:"-"`
	// Match the verification type exactly, ignoring case.
	VerifiedType param.Opt[string] `query:"verifiedType,omitzero" json:"-"`
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
	// Match any comma-separated or line-separated bio term, ignoring case.
	BioContains param.Opt[string] `query:"bioContains,omitzero" json:"-"`
	// Only return profiles X marks as blue verified. X marks many business &
	// government accounts blue too. Use verifiedType for 1 badge type.
	BlueVerifiedOnly param.Opt[bool] `query:"blueVerifiedOnly,omitzero" json:"-"`
	// Only return profiles that publish an email outside free webmail, such as Gmail.
	BusinessEmailOnly param.Opt[bool] `query:"businessEmailOnly,omitzero" json:"-"`
	// Pagination cursor for collection results.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Only return profiles with an email at one of these domains. Separate domains
	// with commas. Subdomains match.
	EmailDomains param.Opt[string] `query:"emailDomains,omitzero" json:"-"`
	// Fields to keep in each row, comma-separated. Use dots for nested fields, such as
	// author.username. A list keeps the field in each item. Case & underscores do not
	// matter. A row that lacks a field leaves it out. Cursors, counts & diagnostics
	// always stay. An unknown name answers 400 unknown_field with the valid names,
	// before any read. The price stays the same. SDK row types then hold only the
	// named fields.
	Fields param.Opt[string] `query:"fields,omitzero" json:"-"`
	// Only return profiles that publish an email.
	HasEmail param.Opt[bool] `query:"hasEmail,omitzero" json:"-"`
	// Only return profiles with a location.
	HasLocation param.Opt[bool] `query:"hasLocation,omitzero" json:"-"`
	// Only return profiles with a website.
	HasWebsite param.Opt[bool] `query:"hasWebsite,omitzero" json:"-"`
	// Match a location substring, ignoring case.
	LocationContains param.Opt[string] `query:"locationContains,omitzero" json:"-"`
	// Maximum follower count. Missing counts pass this maximum.
	MaxFollowers param.Opt[int64] `query:"maxFollowers,omitzero" json:"-"`
	// Profiles may follow at most this many accounts.
	MaxFollowing param.Opt[int64] `query:"maxFollowing,omitzero" json:"-"`
	// Maximum post count. maxPosts is also accepted.
	MaxStatuses param.Opt[int64] `query:"maxStatuses,omitzero" json:"-"`
	// Minimum account age in whole days.
	MinAccountAgeDays param.Opt[int64] `query:"minAccountAgeDays,omitzero" json:"-"`
	// Minimum follower count. Filtering happens before billing.
	MinFollowers param.Opt[int64] `query:"minFollowers,omitzero" json:"-"`
	// Profiles must follow at least this many accounts.
	MinFollowing param.Opt[int64] `query:"minFollowing,omitzero" json:"-"`
	// Minimum post count. minPosts is also accepted.
	MinStatuses param.Opt[int64] `query:"minStatuses,omitzero" json:"-"`
	// Fields to drop from each row, comma-separated. It takes the same names as
	// fields. With fields, it drops from the kept fields.
	Omit param.Opt[string] `query:"omit,omitzero" json:"-"`
	// Match a username substring, ignoring case.
	UsernameContains param.Opt[string] `query:"usernameContains,omitzero" json:"-"`
	// Only return verified profiles.
	VerifiedOnly param.Opt[bool] `query:"verifiedOnly,omitzero" json:"-"`
	// Match the verification type exactly, ignoring case.
	VerifiedType param.Opt[string] `query:"verifiedType,omitzero" json:"-"`
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
	// Words or quoted phrases where any one can match. Separate with spaces, commas,
	// or lines.
	AnyWords param.Opt[string] `query:"anyWords,omitzero" json:"-"`
	// Only return tweets from Blue-verified authors.
	BlueVerifiedOnly param.Opt[bool] `query:"blueVerifiedOnly,omitzero" json:"-"`
	// Cashtags separated by spaces, commas, or lines.
	Cashtags param.Opt[string] `query:"cashtags,omitzero" json:"-"`
	// Conversation ID filter.
	ConversationID param.Opt[string] `query:"conversationId,omitzero" json:"-"`
	// Pagination cursor for community search
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Match this literal phrase, including any hyphens.
	ExactPhrase param.Opt[string] `query:"exactPhrase,omitzero" json:"-"`
	// Words or quoted phrases to exclude. Separate with spaces, commas, or lines.
	ExcludeWords param.Opt[string] `query:"excludeWords,omitzero" json:"-"`
	// Fields to keep in each row, comma-separated. Use dots for nested fields, such as
	// author.username. A list keeps the field in each item. Case & underscores do not
	// matter. A row that lacks a field leaves it out. Cursors, counts & diagnostics
	// always stay. An unknown name answers 400 unknown_field with the valid names,
	// before any read. The price stays the same. SDK row types then hold only the
	// named fields.
	Fields param.Opt[string] `query:"fields,omitzero" json:"-"`
	// Filter by author username.
	FromUser param.Opt[string] `query:"fromUser,omitzero" json:"-"`
	// Hashtags separated by spaces, commas, or lines.
	Hashtags param.Opt[string] `query:"hashtags,omitzero" json:"-"`
	// Only replies to this tweet ID.
	InReplyToTweetID param.Opt[string] `query:"inReplyToTweetId,omitzero" json:"-"`
	// Filter by language. Alias `lang` is accepted.
	Language param.Opt[string] `query:"language,omitzero" json:"-"`
	// Maximum likes threshold. maxLikes & max_likes also work.
	MaxFaves param.Opt[int64] `query:"maxFaves,omitzero" json:"-"`
	// Return Tweets at or below this ID.
	MaxID param.Opt[string] `query:"maxId,omitzero" json:"-"`
	// Maximum quotes threshold.
	MaxQuotes param.Opt[int64] `query:"maxQuotes,omitzero" json:"-"`
	// Maximum replies threshold.
	MaxReplies param.Opt[int64] `query:"maxReplies,omitzero" json:"-"`
	// Maximum retweets threshold.
	MaxRetweets param.Opt[int64] `query:"maxRetweets,omitzero" json:"-"`
	// Filter tweets mentioning a username.
	Mentioning param.Opt[string] `query:"mentioning,omitzero" json:"-"`
	// Minimum bookmark count threshold.
	MinBookmarks param.Opt[int64] `query:"minBookmarks,omitzero" json:"-"`
	// Minimum likes. Aliases: minFaves, min_likes, min_faves.
	MinLikes param.Opt[int64] `query:"minLikes,omitzero" json:"-"`
	// Minimum quote count threshold.
	MinQuotes param.Opt[int64] `query:"minQuotes,omitzero" json:"-"`
	// Minimum replies threshold.
	MinReplies param.Opt[int64] `query:"minReplies,omitzero" json:"-"`
	// Minimum retweets threshold.
	MinRetweets param.Opt[int64] `query:"minRetweets,omitzero" json:"-"`
	// Minimum view count threshold.
	MinViews param.Opt[int64] `query:"minViews,omitzero" json:"-"`
	// Only return native reposts.
	NativeRetweets param.Opt[bool] `query:"nativeRetweets,omitzero" json:"-"`
	// Fields to drop from each row, comma-separated. It takes the same names as
	// fields. With fields, it drops from the kept fields.
	Omit param.Opt[string] `query:"omit,omitzero" json:"-"`
	// Maximum page items (1-100, default 20). Source, filters, or credits can reduce
	// results. Follow next_cursor while the response reports more pages. Deprecated
	// limit and count aliases remain accepted.
	PageSize param.Opt[int64] `query:"pageSize,omitzero" json:"-"`
	// Only quotes of this tweet ID.
	QuotesOfTweetID param.Opt[string] `query:"quotesOfTweetId,omitzero" json:"-"`
	// Only retweets of this tweet ID.
	RetweetsOfTweetID param.Opt[string] `query:"retweetsOfTweetId,omitzero" json:"-"`
	// Start date in YYYY-MM-DD format.
	SinceDate param.Opt[time.Time] `query:"sinceDate,omitzero" format:"date" json:"-"`
	// Return Tweets whose IDs exceed this ID.
	SinceID param.Opt[string] `query:"sinceId,omitzero" json:"-"`
	// Filter replies sent to a username.
	ToUser param.Opt[string] `query:"toUser,omitzero" json:"-"`
	// End date in YYYY-MM-DD format.
	UntilDate param.Opt[time.Time] `query:"untilDate,omitzero" format:"date" json:"-"`
	// URL substring or domain filter.
	URL param.Opt[string] `query:"url,omitzero" json:"-"`
	// Only return tweets from verified authors.
	VerifiedOnly param.Opt[bool] `query:"verifiedOnly,omitzero" json:"-"`
	// Keeps Tweets from a recent window, such as 90m or 7d. Use a whole number & s, m,
	// h or d.
	WithinTime param.Opt[string] `query:"withinTime,omitzero" json:"-"`
	// Filter media. Aliases: has_video, has_media.
	//
	// Any of "images", "videos", "gifs", "media", "links", "none".
	MediaType XCommunityGetSearchParamsMediaType `query:"mediaType,omitzero" json:"-"`
	// Sort order (Latest or Top)
	//
	// Any of "Latest", "Top".
	QueryType XCommunityGetSearchParamsQueryType `query:"queryType,omitzero" json:"-"`
	// Only when the caller requests a quote mode.
	//
	// Any of "include", "exclude", "only".
	Quotes XCommunityGetSearchParamsQuotes `query:"quotes,omitzero" json:"-"`
	// Reply mode. Exclude keeps own-thread replies.
	//
	// Any of "include", "exclude", "only".
	Replies XCommunityGetSearchParamsReplies `query:"replies,omitzero" json:"-"`
	// Only when the caller requests a repost mode.
	//
	// Any of "include", "exclude", "only".
	Retweets XCommunityGetSearchParamsRetweets `query:"retweets,omitzero" json:"-"`
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

// Filter media. Aliases: has_video, has_media.
type XCommunityGetSearchParamsMediaType string

const (
	XCommunityGetSearchParamsMediaTypeImages XCommunityGetSearchParamsMediaType = "images"
	XCommunityGetSearchParamsMediaTypeVideos XCommunityGetSearchParamsMediaType = "videos"
	XCommunityGetSearchParamsMediaTypeGifs   XCommunityGetSearchParamsMediaType = "gifs"
	XCommunityGetSearchParamsMediaTypeMedia  XCommunityGetSearchParamsMediaType = "media"
	XCommunityGetSearchParamsMediaTypeLinks  XCommunityGetSearchParamsMediaType = "links"
	XCommunityGetSearchParamsMediaTypeNone   XCommunityGetSearchParamsMediaType = "none"
)

// Sort order (Latest or Top)
type XCommunityGetSearchParamsQueryType string

const (
	XCommunityGetSearchParamsQueryTypeLatest XCommunityGetSearchParamsQueryType = "Latest"
	XCommunityGetSearchParamsQueryTypeTop    XCommunityGetSearchParamsQueryType = "Top"
)

type XCommunityFindParams struct {
	// Keyword to find communities by, such as ai. Alias query.
	Q string `query:"q" api:"required" json:"-"`
	// Pagination cursor for collection results. Also read as next_cursor, next_token,
	// pagination_token or continuation_token.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Fields to keep in each row, comma-separated. Use dots for nested fields, such as
	// author.username. A list keeps the field in each item. Case & underscores do not
	// matter. A row that lacks a field leaves it out. Cursors, counts & diagnostics
	// always stay. An unknown name answers 400 unknown_field with the valid names,
	// before any read. The price stays the same. SDK row types then hold only the
	// named fields.
	Fields param.Opt[string] `query:"fields,omitzero" json:"-"`
	// Fields to drop from each row, comma-separated. It takes the same names as
	// fields. With fields, it drops from the kept fields.
	Omit param.Opt[string] `query:"omit,omitzero" json:"-"`
	// Items per page (1-100, default 20). Source, filters or credits can return fewer.
	// Follow next_cursor for more pages. Its aliases are limit, count, max_results,
	// maxItems, max_items & per_page.
	PageSize param.Opt[int64] `query:"pageSize,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [XCommunityFindParams]'s query parameters as `url.Values`.
func (r XCommunityFindParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type XCommunityListPopularParams struct {
	// Pagination cursor for collection results. Also read as next_cursor, next_token,
	// pagination_token or continuation_token.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Fields to keep in each row, comma-separated. Use dots for nested fields, such as
	// author.username. A list keeps the field in each item. Case & underscores do not
	// matter. A row that lacks a field leaves it out. Cursors, counts & diagnostics
	// always stay. An unknown name answers 400 unknown_field with the valid names,
	// before any read. The price stays the same. SDK row types then hold only the
	// named fields.
	Fields param.Opt[string] `query:"fields,omitzero" json:"-"`
	// Fields to drop from each row, comma-separated. It takes the same names as
	// fields. With fields, it drops from the kept fields.
	Omit param.Opt[string] `query:"omit,omitzero" json:"-"`
	// Items per page (1-100, default 20). Source, filters or credits can return fewer.
	// Follow next_cursor for more pages. Its aliases are limit, count, max_results,
	// maxItems, max_items & per_page.
	PageSize param.Opt[int64] `query:"pageSize,omitzero" json:"-"`
	// Topic or subtopic name in any case, such as technology. Send this or topicId.
	Topic param.Opt[string] `query:"topic,omitzero" json:"-"`
	// Topic or subtopic ID from /x/communities/topics, such as 301. Wins over topic.
	TopicID param.Opt[string] `query:"topicId,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [XCommunityListPopularParams]'s query parameters as `url.Values`.
func (r XCommunityListPopularParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type XCommunityListSuggestedParams struct {
	// Fields to keep in each row, comma-separated. Use dots for nested fields, such as
	// author.username. A list keeps the field in each item. Case & underscores do not
	// matter. A row that lacks a field leaves it out. Cursors, counts & diagnostics
	// always stay. An unknown name answers 400 unknown_field with the valid names,
	// before any read. The price stays the same. SDK row types then hold only the
	// named fields.
	Fields param.Opt[string] `query:"fields,omitzero" json:"-"`
	// Maximum communities to return. Omit it for every suggestion. A low credit
	// balance can return fewer. Aliases pageSize, count, max_results, maxItems,
	// max_items & per_page.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Fields to drop from each row, comma-separated. It takes the same names as
	// fields. With fields, it drops from the kept fields.
	Omit param.Opt[string] `query:"omit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [XCommunityListSuggestedParams]'s query parameters as `url.Values`.
func (r XCommunityListSuggestedParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type XCommunityListTopicsParams struct {
	// Fields to keep in each row, comma-separated. Use dots for nested fields, such as
	// author.username. A list keeps the field in each item. Case & underscores do not
	// matter. A row that lacks a field leaves it out. Cursors, counts & diagnostics
	// always stay. An unknown name answers 400 unknown_field with the valid names,
	// before any read. The price stays the same. SDK row types then hold only the
	// named fields.
	Fields param.Opt[string] `query:"fields,omitzero" json:"-"`
	// Maximum topics to return. Omit it for every topic. A low credit balance can
	// return fewer. Aliases pageSize, count, max_results, maxItems, max_items &
	// per_page.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Fields to drop from each row, comma-separated. It takes the same names as
	// fields. With fields, it drops from the kept fields.
	Omit param.Opt[string] `query:"omit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [XCommunityListTopicsParams]'s query parameters as `url.Values`.
func (r XCommunityListTopicsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type XCommunitySearchMembersParams struct {
	// Name or username to find members by, such as marckohlbrugge. A leading @ is
	// dropped. Alias query.
	Q string `query:"q" api:"required" json:"-"`
	// Match any comma-separated or line-separated bio term, ignoring case.
	BioContains param.Opt[string] `query:"bioContains,omitzero" json:"-"`
	// Only return profiles X marks as blue verified. X marks many business &
	// government accounts blue too. Use verifiedType for 1 badge type.
	BlueVerifiedOnly param.Opt[bool] `query:"blueVerifiedOnly,omitzero" json:"-"`
	// Only return profiles that publish an email outside free webmail, such as Gmail.
	BusinessEmailOnly param.Opt[bool] `query:"businessEmailOnly,omitzero" json:"-"`
	// Only return profiles with an email at one of these domains. Separate domains
	// with commas. Subdomains match.
	EmailDomains param.Opt[string] `query:"emailDomains,omitzero" json:"-"`
	// Fields to keep in each row, comma-separated. Use dots for nested fields, such as
	// author.username. A list keeps the field in each item. Case & underscores do not
	// matter. A row that lacks a field leaves it out. Cursors, counts & diagnostics
	// always stay. An unknown name answers 400 unknown_field with the valid names,
	// before any read. The price stays the same. SDK row types then hold only the
	// named fields.
	Fields param.Opt[string] `query:"fields,omitzero" json:"-"`
	// Only return profiles that publish an email.
	HasEmail param.Opt[bool] `query:"hasEmail,omitzero" json:"-"`
	// Only return profiles with a location.
	HasLocation param.Opt[bool] `query:"hasLocation,omitzero" json:"-"`
	// Only return profiles with a website.
	HasWebsite param.Opt[bool] `query:"hasWebsite,omitzero" json:"-"`
	// Match a location substring, ignoring case.
	LocationContains param.Opt[string] `query:"locationContains,omitzero" json:"-"`
	// Maximum follower count. A profile without a count passes.
	MaxFollowers param.Opt[int64] `query:"maxFollowers,omitzero" json:"-"`
	// Profiles may follow at most this many accounts.
	MaxFollowing param.Opt[int64] `query:"maxFollowing,omitzero" json:"-"`
	// Maximum post count. maxPosts is also accepted.
	MaxStatuses param.Opt[int64] `query:"maxStatuses,omitzero" json:"-"`
	// Minimum account age in whole days.
	MinAccountAgeDays param.Opt[int64] `query:"minAccountAgeDays,omitzero" json:"-"`
	// Minimum follower count. Profiles the filter drops cost nothing.
	MinFollowers param.Opt[int64] `query:"minFollowers,omitzero" json:"-"`
	// Profiles must follow at least this many accounts.
	MinFollowing param.Opt[int64] `query:"minFollowing,omitzero" json:"-"`
	// Minimum post count. minPosts is also accepted.
	MinStatuses param.Opt[int64] `query:"minStatuses,omitzero" json:"-"`
	// Fields to drop from each row, comma-separated. It takes the same names as
	// fields. With fields, it drops from the kept fields.
	Omit param.Opt[string] `query:"omit,omitzero" json:"-"`
	// Match a username substring, ignoring case.
	UsernameContains param.Opt[string] `query:"usernameContains,omitzero" json:"-"`
	// Only return verified profiles.
	VerifiedOnly param.Opt[bool] `query:"verifiedOnly,omitzero" json:"-"`
	// Match the verification type exactly, ignoring case.
	VerifiedType param.Opt[string] `query:"verifiedType,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [XCommunitySearchMembersParams]'s query parameters as `url.Values`.
func (r XCommunitySearchMembersParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Reply mode. Exclude keeps own-thread replies.
type XCommunityGetSearchParamsReplies string

const (
	XCommunityGetSearchParamsRepliesInclude XCommunityGetSearchParamsReplies = "include"
	XCommunityGetSearchParamsRepliesExclude XCommunityGetSearchParamsReplies = "exclude"
	XCommunityGetSearchParamsRepliesOnly    XCommunityGetSearchParamsReplies = "only"
)

// Only when the caller requests a repost mode.
type XCommunityGetSearchParamsRetweets string

const (
	XCommunityGetSearchParamsRetweetsInclude XCommunityGetSearchParamsRetweets = "include"
	XCommunityGetSearchParamsRetweetsExclude XCommunityGetSearchParamsRetweets = "exclude"
	XCommunityGetSearchParamsRetweetsOnly    XCommunityGetSearchParamsRetweets = "only"
)

// Only when the caller requests a quote mode.
type XCommunityGetSearchParamsQuotes string

const (
	XCommunityGetSearchParamsQuotesInclude XCommunityGetSearchParamsQuotes = "include"
	XCommunityGetSearchParamsQuotesExclude XCommunityGetSearchParamsQuotes = "exclude"
	XCommunityGetSearchParamsQuotesOnly    XCommunityGetSearchParamsQuotes = "only"
)
