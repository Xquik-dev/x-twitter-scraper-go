// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

package xtwitterscraper

import (
	"context"
	"encoding/json"
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
// You should not instantiate this service directly, and instead use the
// [NewXUserService] method instead.
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

// Returns one public profile with counts and verification details.
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

// Returns the accounts an organization lists as its affiliates, in X's order. An
// account without affiliates returns an empty list.
func (r *XUserService) GetAffiliates(ctx context.Context, id string, query XUserGetAffiliatesParams, opts ...option.RequestOption) (res *shared.PaginatedUsers, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("x/users/%s/affiliates", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Returns the posts on the user's Articles tab, newest first. Each post's
// `article` holds the Article's ID, title, preview, cover & dates. Read the full
// text with GET /x/articles/{tweetId}.
func (r *XUserService) GetArticles(ctx context.Context, id string, query XUserGetArticlesParams, opts ...option.RequestOption) (res *shared.PaginatedTweets, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("x/users/%s/articles", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Returns the newest posts of up to 20 users, named by user ID, username or
// profile link, in the order named. Send ids or usernames, not both. Each post
// costs 1 credit. A user in unavailable_ids or unprocessed_ids costs nothing.
func (r *XUserService) GetBatchTweets(ctx context.Context, query XUserGetBatchTweetsParams, opts ...option.RequestOption) (res *XUserGetBatchTweetsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "x/users/batch/tweets"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Returns the user IDs of the selected user's followers, newest first, up to 5,000
// a page. Use it when you need IDs only. Each ID costs 1 credit.
func (r *XUserService) GetFollowerIDs(ctx context.Context, id string, query XUserGetFollowerIDsParams, opts ...option.RequestOption) (res *XUserGetFollowerIDsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("x/users/%s/follower-ids", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Returns the user IDs of the accounts the selected user follows, newest first, up
// to 5,000 a page. Use it when you need IDs only. Each ID costs 1 credit.
func (r *XUserService) GetFollowingIDs(ctx context.Context, id string, query XUserGetFollowingIDsParams, opts ...option.RequestOption) (res *XUserGetFollowingIDsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("x/users/%s/following-ids", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Returns the posts on the user's Highlights tab, in X's order. An account without
// highlights returns an empty list.
func (r *XUserService) GetHighlights(ctx context.Context, id string, query XUserGetHighlightsParams, opts ...option.RequestOption) (res *shared.PaginatedTweets, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("x/users/%s/highlights", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Returns engagement insights for 1 profile. They are engagement rate, averages
// per post and posting rate, from its counts and its last 20 posts. Reposts and
// the pinned post stay out of the averages. Each report costs 1 credit, whatever
// it reads.
func (r *XUserService) GetInsights(ctx context.Context, id string, query XUserGetInsightsParams, opts ...option.RequestOption) (res *XUserGetInsightsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("x/users/%s/insights", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Returns the public X Lists that include the selected user, as the profile's List
// memberships page shows them. Each List costs 1 credit. A protected account
// answers 403.
func (r *XUserService) GetListMemberships(ctx context.Context, id string, query XUserGetListMembershipsParams, opts ...option.RequestOption) (res *XUserGetListMembershipsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("x/users/%s/list-memberships", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Returns the public X Lists the selected user created, in X's order, as the
// profile's Lists tab shows them. Each List costs 1 credit. A protected account
// answers 403.
func (r *XUserService) GetLists(ctx context.Context, id string, query XUserGetListsParams, opts ...option.RequestOption) (res *XUserGetListsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("x/users/%s/lists", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Says whether 1 user is broadcasting live video on X right now, with their latest
// broadcast, live or ended. Each check costs 1 credit.
func (r *XUserService) GetLive(ctx context.Context, id string, query XUserGetLiveParams, opts ...option.RequestOption) (res *XUserGetLiveResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("x/users/%s/live", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Returns the public accounts X lists as similar to the selected user, in X's
// order. X can list slightly different accounts on each call. Each account costs 1
// credit.
func (r *XUserService) GetSimilar(ctx context.Context, id string, query XUserGetSimilarParams, opts ...option.RequestOption) (res *shared.PaginatedUsers, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("x/users/%s/similar", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Returns the creators the selected user subscribes to, in X's order. An account
// that hides its subscriptions returns an empty list.
func (r *XUserService) GetSubscriptions(ctx context.Context, id string, query XUserGetSubscriptionsParams, opts ...option.RequestOption) (res *shared.PaginatedUsers, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("x/users/%s/subscriptions", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Removes a follower from a connected X account.
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

// Returns public profiles for the requested user IDs.
func (r *XUserService) GetBatch(ctx context.Context, query XUserGetBatchParams, opts ...option.RequestOption) (res *XUserGetBatchResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "x/users/batch"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Returns public follower profiles for the selected user.
func (r *XUserService) GetFollowers(ctx context.Context, id string, query XUserGetFollowersParams, opts ...option.RequestOption) (res *XUserGetFollowersResponseUnion, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("x/users/%s/followers", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Returns mutual profiles visible to the connected X account.
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

// Returns public profiles followed by the selected user.
func (r *XUserService) GetFollowing(ctx context.Context, id string, query XUserGetFollowingParams, opts ...option.RequestOption) (res *XUserGetFollowingResponseUnion, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("x/users/%s/following", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Returns visible tweets liked by the selected user.
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

// Returns public media posts from the selected user.
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

// Returns public posts mentioning the selected user.
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

// Returns target-authored posts and replies. Omit mode for automatic maximum
// coverage. Pass next_cursor unchanged. Unprefixed cursors stay legacy. Excludes
// other-author context.
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

// Returns public profiles matching a name or username.
func (r *XUserService) GetSearch(ctx context.Context, query XUserGetSearchParams, opts ...option.RequestOption) (res *shared.PaginatedUsers, err error) {
	opts = slices.Concat(r.options, opts)
	path := "x/users/search"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// For multiple users, use createExtraction with tweet_search_extractor and
// targetUsernames. Omit mode for automatic maximum coverage. Pass next_cursor
// unchanged. Unprefixed cursors use legacy pagination.
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

// Returns verified public profiles following the selected user.
func (r *XUserService) GetVerifiedFollowers(ctx context.Context, id string, query XUserGetVerifiedFollowersParams, opts ...option.RequestOption) (res *XUserGetVerifiedFollowersResponseUnion, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("x/users/%s/verified-followers", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Returns X's own translation of 1 profile's bio into the language you pick, with
// the links, mentions, and tags X places in the translated text. A profile without
// a bio returns translation null & a message, free of charge.
func (r *XUserService) TranslateBio(ctx context.Context, id string, query XUserTranslateBioParams, opts ...option.RequestOption) (res *XUserTranslateBioResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("x/users/%s/translation", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Durable write record. Poll statusUrl until terminal is true. Reusing its
// Idempotency-Key returns this record. Create another action only when safeToRetry
// is true.
type XUserRemoveFollowerResponse struct {
	ID string `json:"id" api:"required"`
	// Connected account selected for the write.
	Account XUserRemoveFollowerResponseAccount `json:"account" api:"required"`
	// Any of "create_tweet", "delete_tweet", "like", "unlike", "retweet", "unretweet",
	// "follow", "unfollow", "remove_follower", "send_dm", "delete_dm", "upload_media",
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
	XUserRemoveFollowerResponseActionDeleteDm        XUserRemoveFollowerResponseAction = "delete_dm"
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

// Batch lookup preserves first-seen ID order and ignores duplicate IDs.
// unavailable_ids were processed without profiles. unprocessed_ids were
// credit-limited.
type XUserGetBatchResponse struct {
	// Processed IDs that failed to load, in first-seen request order. They cost
	// nothing. Retry them.
	FailedIDs []string `json:"failed_ids" api:"required"`
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
	// Profiles rejected by the requested filters.
	FilteredCount int64 `json:"filtered_count"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FailedIDs      respjson.Field
		HasNextPage    respjson.Field
		NextCursor     respjson.Field
		ProcessedCount respjson.Field
		RequestedCount respjson.Field
		ReturnedCount  respjson.Field
		UnavailableIDs respjson.Field
		UnprocessedIDs respjson.Field
		Users          respjson.Field
		FilteredCount  respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XUserGetBatchResponse) RawJSON() string { return r.JSON.raw }
func (r *XUserGetBatchResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// XUserGetFollowersResponseUnion contains all possible properties and values from
// [shared.PaginatedUsers], [XUserGetFollowersResponseUserListCoverageResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type XUserGetFollowersResponseUnion struct {
	HasNextPage bool   `json:"has_next_page"`
	NextCursor  string `json:"next_cursor"`
	// Present on the last page of a filtered automatic read that sent no maxScan, once
	// the list scanned 100,000 profiles across its pages. That page is the list's
	// last.
	ScanLimit XUserGetFollowersResponseScanLimit `json:"scanLimit"`
	// This field is from variant [shared.PaginatedUsers],
	// [XUserGetFollowersResponseUserListCoverageResponse].
	Users []shared.UserProfile `json:"users"`
	// This field is from variant [shared.PaginatedUsers],
	// [XUserGetFollowersResponseUserListCoverageResponse].
	FilteredCount int64 `json:"filtered_count"`
	// This field is from variant [XUserGetFollowersResponseUserListCoverageResponse].
	Diagnostic XUserGetFollowersResponseUserListCoverageResponseDiagnostic `json:"diagnostic"`
	JSON       struct {
		HasNextPage   respjson.Field
		NextCursor    respjson.Field
		ScanLimit     respjson.Field
		Users         respjson.Field
		FilteredCount respjson.Field
		Diagnostic    respjson.Field
		raw           string
	} `json:"-"`
}

func (u XUserGetFollowersResponseUnion) AsPaginatedUsers() (v shared.PaginatedUsers) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u XUserGetFollowersResponseUnion) AsXUserGetFollowersResponseUserListCoverageResponse() (v XUserGetFollowersResponseUserListCoverageResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u XUserGetFollowersResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *XUserGetFollowersResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Present on the last page of a filtered automatic read that sent no maxScan, once
// the list scanned 100,000 profiles across its pages. That page is the list's
// last.
type XUserGetFollowersResponseScanLimit struct {
	// The profiles a filtered read scans of a list without maxScan.
	//
	// Any of 100000.
	Limit int64 `json:"limit" api:"required"`
	// What stopped the list & how to read further.
	Message string `json:"message" api:"required"`
	// Profiles the list read across its pages, kept or dropped by the filters. The
	// last page may read past the limit.
	RowsRead int64 `json:"rowsRead" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Limit       respjson.Field
		Message     respjson.Field
		RowsRead    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XUserGetFollowersResponseScanLimit) RawJSON() string { return r.JSON.raw }
func (r *XUserGetFollowersResponseScanLimit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Terminal relationship coverage response with diagnostics.
type XUserGetFollowersResponseUserListCoverageResponse struct {
	// Coverage evidence across parallel relationship strategies.
	Diagnostic XUserGetFollowersResponseUserListCoverageResponseDiagnostic `json:"diagnostic" api:"required"`
	// Any of false.
	HasNextPage bool `json:"has_next_page"`
	// Any of "".
	NextCursor string `json:"next_cursor"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Diagnostic  respjson.Field
		HasNextPage respjson.Field
		NextCursor  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.PaginatedUsers
}

// Returns the unmodified JSON received from the API
func (r XUserGetFollowersResponseUserListCoverageResponse) RawJSON() string { return r.JSON.raw }
func (r *XUserGetFollowersResponseUserListCoverageResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Coverage evidence across parallel relationship strategies.
type XUserGetFollowersResponseUserListCoverageResponseDiagnostic struct {
	// True after all active strategies exhaust their sources.
	Complete            bool  `json:"complete" api:"required"`
	CursorFailureCount  int64 `json:"cursorFailureCount" api:"required"`
	DeadlineReached     bool  `json:"deadlineReached" api:"required"`
	DuplicateCount      int64 `json:"duplicateCount" api:"required"`
	FailedStrategyCount int64 `json:"failedStrategyCount" api:"required"`
	MalformedCount      int64 `json:"malformedCount" api:"required"`
	PagesFetched        int64 `json:"pagesFetched" api:"required"`
	// True when credits or the requested limit reduce output.
	ResponseTruncated    bool                                                                  `json:"responseTruncated" api:"required"`
	ResultLimitReached   bool                                                                  `json:"resultLimitReached" api:"required"`
	ReturnedUsers        int64                                                                 `json:"returnedUsers" api:"required"`
	StalledStrategyCount int64                                                                 `json:"stalledStrategyCount" api:"required"`
	Strategies           []XUserGetFollowersResponseUserListCoverageResponseDiagnosticStrategy `json:"strategies" api:"required"`
	StrategyCount        int64                                                                 `json:"strategyCount" api:"required"`
	UniqueUsers          int64                                                                 `json:"uniqueUsers" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Complete             respjson.Field
		CursorFailureCount   respjson.Field
		DeadlineReached      respjson.Field
		DuplicateCount       respjson.Field
		FailedStrategyCount  respjson.Field
		MalformedCount       respjson.Field
		PagesFetched         respjson.Field
		ResponseTruncated    respjson.Field
		ResultLimitReached   respjson.Field
		ReturnedUsers        respjson.Field
		StalledStrategyCount respjson.Field
		Strategies           respjson.Field
		StrategyCount        respjson.Field
		UniqueUsers          respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XUserGetFollowersResponseUserListCoverageResponseDiagnostic) RawJSON() string {
	return r.JSON.raw
}
func (r *XUserGetFollowersResponseUserListCoverageResponseDiagnostic) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Result counts and stop reason for one relationship strategy.
type XUserGetFollowersResponseUserListCoverageResponseDiagnosticStrategy struct {
	DuplicateCount int64 `json:"duplicateCount" api:"required"`
	PagesFetched   int64 `json:"pagesFetched" api:"required"`
	// Reason a coverage strategy stopped.
	//
	// Any of "cursor_failure", "deadline", "exhausted", "failed", "page_limit",
	// "result_limit", "stalled".
	StopReason  string `json:"stopReason" api:"required"`
	Strategy    int64  `json:"strategy" api:"required"`
	UniqueAdded int64  `json:"uniqueAdded" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DuplicateCount respjson.Field
		PagesFetched   respjson.Field
		StopReason     respjson.Field
		Strategy       respjson.Field
		UniqueAdded    respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XUserGetFollowersResponseUserListCoverageResponseDiagnosticStrategy) RawJSON() string {
	return r.JSON.raw
}
func (r *XUserGetFollowersResponseUserListCoverageResponseDiagnosticStrategy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// XUserGetFollowingResponseUnion contains all possible properties and values from
// [shared.PaginatedUsers], [XUserGetFollowingResponseUserListCoverageResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type XUserGetFollowingResponseUnion struct {
	HasNextPage bool   `json:"has_next_page"`
	NextCursor  string `json:"next_cursor"`
	// Present on the last page of a filtered automatic read that sent no maxScan, once
	// the list scanned 100,000 profiles across its pages. That page is the list's
	// last.
	ScanLimit XUserGetFollowingResponseScanLimit `json:"scanLimit"`
	// This field is from variant [shared.PaginatedUsers],
	// [XUserGetFollowingResponseUserListCoverageResponse].
	Users []shared.UserProfile `json:"users"`
	// This field is from variant [shared.PaginatedUsers],
	// [XUserGetFollowingResponseUserListCoverageResponse].
	FilteredCount int64 `json:"filtered_count"`
	// This field is from variant [XUserGetFollowingResponseUserListCoverageResponse].
	Diagnostic XUserGetFollowingResponseUserListCoverageResponseDiagnostic `json:"diagnostic"`
	JSON       struct {
		HasNextPage   respjson.Field
		NextCursor    respjson.Field
		ScanLimit     respjson.Field
		Users         respjson.Field
		FilteredCount respjson.Field
		Diagnostic    respjson.Field
		raw           string
	} `json:"-"`
}

func (u XUserGetFollowingResponseUnion) AsPaginatedUsers() (v shared.PaginatedUsers) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u XUserGetFollowingResponseUnion) AsXUserGetFollowingResponseUserListCoverageResponse() (v XUserGetFollowingResponseUserListCoverageResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u XUserGetFollowingResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *XUserGetFollowingResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Present on the last page of a filtered automatic read that sent no maxScan, once
// the list scanned 100,000 profiles across its pages. That page is the list's
// last.
type XUserGetFollowingResponseScanLimit struct {
	// The profiles a filtered read scans of a list without maxScan.
	//
	// Any of 100000.
	Limit int64 `json:"limit" api:"required"`
	// What stopped the list & how to read further.
	Message string `json:"message" api:"required"`
	// Profiles the list read across its pages, kept or dropped by the filters. The
	// last page may read past the limit.
	RowsRead int64 `json:"rowsRead" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Limit       respjson.Field
		Message     respjson.Field
		RowsRead    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XUserGetFollowingResponseScanLimit) RawJSON() string { return r.JSON.raw }
func (r *XUserGetFollowingResponseScanLimit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Terminal relationship coverage response with diagnostics.
type XUserGetFollowingResponseUserListCoverageResponse struct {
	// Coverage evidence across parallel relationship strategies.
	Diagnostic XUserGetFollowingResponseUserListCoverageResponseDiagnostic `json:"diagnostic" api:"required"`
	// Any of false.
	HasNextPage bool `json:"has_next_page"`
	// Any of "".
	NextCursor string `json:"next_cursor"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Diagnostic  respjson.Field
		HasNextPage respjson.Field
		NextCursor  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.PaginatedUsers
}

// Returns the unmodified JSON received from the API
func (r XUserGetFollowingResponseUserListCoverageResponse) RawJSON() string { return r.JSON.raw }
func (r *XUserGetFollowingResponseUserListCoverageResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Coverage evidence across parallel relationship strategies.
type XUserGetFollowingResponseUserListCoverageResponseDiagnostic struct {
	// True after all active strategies exhaust their sources.
	Complete            bool  `json:"complete" api:"required"`
	CursorFailureCount  int64 `json:"cursorFailureCount" api:"required"`
	DeadlineReached     bool  `json:"deadlineReached" api:"required"`
	DuplicateCount      int64 `json:"duplicateCount" api:"required"`
	FailedStrategyCount int64 `json:"failedStrategyCount" api:"required"`
	MalformedCount      int64 `json:"malformedCount" api:"required"`
	PagesFetched        int64 `json:"pagesFetched" api:"required"`
	// True when credits or the requested limit reduce output.
	ResponseTruncated    bool                                                                  `json:"responseTruncated" api:"required"`
	ResultLimitReached   bool                                                                  `json:"resultLimitReached" api:"required"`
	ReturnedUsers        int64                                                                 `json:"returnedUsers" api:"required"`
	StalledStrategyCount int64                                                                 `json:"stalledStrategyCount" api:"required"`
	Strategies           []XUserGetFollowingResponseUserListCoverageResponseDiagnosticStrategy `json:"strategies" api:"required"`
	StrategyCount        int64                                                                 `json:"strategyCount" api:"required"`
	UniqueUsers          int64                                                                 `json:"uniqueUsers" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Complete             respjson.Field
		CursorFailureCount   respjson.Field
		DeadlineReached      respjson.Field
		DuplicateCount       respjson.Field
		FailedStrategyCount  respjson.Field
		MalformedCount       respjson.Field
		PagesFetched         respjson.Field
		ResponseTruncated    respjson.Field
		ResultLimitReached   respjson.Field
		ReturnedUsers        respjson.Field
		StalledStrategyCount respjson.Field
		Strategies           respjson.Field
		StrategyCount        respjson.Field
		UniqueUsers          respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XUserGetFollowingResponseUserListCoverageResponseDiagnostic) RawJSON() string {
	return r.JSON.raw
}
func (r *XUserGetFollowingResponseUserListCoverageResponseDiagnostic) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Result counts and stop reason for one relationship strategy.
type XUserGetFollowingResponseUserListCoverageResponseDiagnosticStrategy struct {
	DuplicateCount int64 `json:"duplicateCount" api:"required"`
	PagesFetched   int64 `json:"pagesFetched" api:"required"`
	// Reason a coverage strategy stopped.
	//
	// Any of "cursor_failure", "deadline", "exhausted", "failed", "page_limit",
	// "result_limit", "stalled".
	StopReason  string `json:"stopReason" api:"required"`
	Strategy    int64  `json:"strategy" api:"required"`
	UniqueAdded int64  `json:"uniqueAdded" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DuplicateCount respjson.Field
		PagesFetched   respjson.Field
		StopReason     respjson.Field
		Strategy       respjson.Field
		UniqueAdded    respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XUserGetFollowingResponseUserListCoverageResponseDiagnosticStrategy) RawJSON() string {
	return r.JSON.raw
}
func (r *XUserGetFollowingResponseUserListCoverageResponseDiagnosticStrategy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// XUserGetVerifiedFollowersResponseUnion contains all possible properties and
// values from [shared.PaginatedUsers],
// [XUserGetVerifiedFollowersResponseUserListCoverageResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type XUserGetVerifiedFollowersResponseUnion struct {
	HasNextPage bool   `json:"has_next_page"`
	NextCursor  string `json:"next_cursor"`
	// Present on the last page of a filtered automatic read that sent no maxScan, once
	// the list scanned 100,000 profiles across its pages. That page is the list's
	// last.
	ScanLimit XUserGetVerifiedFollowersResponseScanLimit `json:"scanLimit"`
	// This field is from variant [shared.PaginatedUsers],
	// [XUserGetVerifiedFollowersResponseUserListCoverageResponse].
	Users []shared.UserProfile `json:"users"`
	// This field is from variant [shared.PaginatedUsers],
	// [XUserGetVerifiedFollowersResponseUserListCoverageResponse].
	FilteredCount int64 `json:"filtered_count"`
	// This field is from variant
	// [XUserGetVerifiedFollowersResponseUserListCoverageResponse].
	Diagnostic XUserGetVerifiedFollowersResponseUserListCoverageResponseDiagnostic `json:"diagnostic"`
	JSON       struct {
		HasNextPage   respjson.Field
		NextCursor    respjson.Field
		ScanLimit     respjson.Field
		Users         respjson.Field
		FilteredCount respjson.Field
		Diagnostic    respjson.Field
		raw           string
	} `json:"-"`
}

func (u XUserGetVerifiedFollowersResponseUnion) AsPaginatedUsers() (v shared.PaginatedUsers) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u XUserGetVerifiedFollowersResponseUnion) AsXUserGetVerifiedFollowersResponseUserListCoverageResponse() (v XUserGetVerifiedFollowersResponseUserListCoverageResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u XUserGetVerifiedFollowersResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *XUserGetVerifiedFollowersResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Present on the last page of a filtered automatic read that sent no maxScan, once
// the list scanned 100,000 profiles across its pages. That page is the list's
// last.
type XUserGetVerifiedFollowersResponseScanLimit struct {
	// The profiles a filtered read scans of a list without maxScan.
	//
	// Any of 100000.
	Limit int64 `json:"limit" api:"required"`
	// What stopped the list & how to read further.
	Message string `json:"message" api:"required"`
	// Profiles the list read across its pages, kept or dropped by the filters. The
	// last page may read past the limit.
	RowsRead int64 `json:"rowsRead" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Limit       respjson.Field
		Message     respjson.Field
		RowsRead    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XUserGetVerifiedFollowersResponseScanLimit) RawJSON() string { return r.JSON.raw }
func (r *XUserGetVerifiedFollowersResponseScanLimit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Terminal relationship coverage response with diagnostics.
type XUserGetVerifiedFollowersResponseUserListCoverageResponse struct {
	// Coverage evidence across parallel relationship strategies.
	Diagnostic XUserGetVerifiedFollowersResponseUserListCoverageResponseDiagnostic `json:"diagnostic" api:"required"`
	// Any of false.
	HasNextPage bool `json:"has_next_page"`
	// Any of "".
	NextCursor string `json:"next_cursor"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Diagnostic  respjson.Field
		HasNextPage respjson.Field
		NextCursor  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.PaginatedUsers
}

// Returns the unmodified JSON received from the API
func (r XUserGetVerifiedFollowersResponseUserListCoverageResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *XUserGetVerifiedFollowersResponseUserListCoverageResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Coverage evidence across parallel relationship strategies.
type XUserGetVerifiedFollowersResponseUserListCoverageResponseDiagnostic struct {
	// True after all active strategies exhaust their sources.
	Complete            bool  `json:"complete" api:"required"`
	CursorFailureCount  int64 `json:"cursorFailureCount" api:"required"`
	DeadlineReached     bool  `json:"deadlineReached" api:"required"`
	DuplicateCount      int64 `json:"duplicateCount" api:"required"`
	FailedStrategyCount int64 `json:"failedStrategyCount" api:"required"`
	MalformedCount      int64 `json:"malformedCount" api:"required"`
	PagesFetched        int64 `json:"pagesFetched" api:"required"`
	// True when credits or the requested limit reduce output.
	ResponseTruncated    bool                                                                          `json:"responseTruncated" api:"required"`
	ResultLimitReached   bool                                                                          `json:"resultLimitReached" api:"required"`
	ReturnedUsers        int64                                                                         `json:"returnedUsers" api:"required"`
	StalledStrategyCount int64                                                                         `json:"stalledStrategyCount" api:"required"`
	Strategies           []XUserGetVerifiedFollowersResponseUserListCoverageResponseDiagnosticStrategy `json:"strategies" api:"required"`
	StrategyCount        int64                                                                         `json:"strategyCount" api:"required"`
	UniqueUsers          int64                                                                         `json:"uniqueUsers" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Complete             respjson.Field
		CursorFailureCount   respjson.Field
		DeadlineReached      respjson.Field
		DuplicateCount       respjson.Field
		FailedStrategyCount  respjson.Field
		MalformedCount       respjson.Field
		PagesFetched         respjson.Field
		ResponseTruncated    respjson.Field
		ResultLimitReached   respjson.Field
		ReturnedUsers        respjson.Field
		StalledStrategyCount respjson.Field
		Strategies           respjson.Field
		StrategyCount        respjson.Field
		UniqueUsers          respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XUserGetVerifiedFollowersResponseUserListCoverageResponseDiagnostic) RawJSON() string {
	return r.JSON.raw
}
func (r *XUserGetVerifiedFollowersResponseUserListCoverageResponseDiagnostic) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Result counts and stop reason for one relationship strategy.
type XUserGetVerifiedFollowersResponseUserListCoverageResponseDiagnosticStrategy struct {
	DuplicateCount int64 `json:"duplicateCount" api:"required"`
	PagesFetched   int64 `json:"pagesFetched" api:"required"`
	// Reason a coverage strategy stopped.
	//
	// Any of "cursor_failure", "deadline", "exhausted", "failed", "page_limit",
	// "result_limit", "stalled".
	StopReason  string `json:"stopReason" api:"required"`
	Strategy    int64  `json:"strategy" api:"required"`
	UniqueAdded int64  `json:"uniqueAdded" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DuplicateCount respjson.Field
		PagesFetched   respjson.Field
		StopReason     respjson.Field
		Strategy       respjson.Field
		UniqueAdded    respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XUserGetVerifiedFollowersResponseUserListCoverageResponseDiagnosticStrategy) RawJSON() string {
	return r.JSON.raw
}
func (r *XUserGetVerifiedFollowersResponseUserListCoverageResponseDiagnosticStrategy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Each user's newest posts, in the order the users were named, with the pinned
// post first. The ID lists name the requested IDs or usernames. They name a
// profile link by its username.
type XUserGetBatchTweetsResponse struct {
	// Always false. The answer holds every post in 1 page.
	//
	// Any of false.
	HasNextPage bool `json:"has_next_page" api:"required"`
	// Always empty. The answer holds every post in 1 page.
	NextCursor string               `json:"next_cursor" api:"required"`
	Tweets     []shared.SearchTweet `json:"tweets" api:"required"`
	// Users with no account on X or with protected posts, in first-seen request order.
	// They cost nothing.
	UnavailableIDs []string `json:"unavailable_ids" api:"required"`
	// Users the call did not read, in first-seen request order. Their read failed, the
	// call ran out of time or credits ran out. They cost nothing. Send them again.
	UnprocessedIDs []string `json:"unprocessed_ids" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		HasNextPage    respjson.Field
		NextCursor     respjson.Field
		Tweets         respjson.Field
		UnavailableIDs respjson.Field
		UnprocessedIDs respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XUserGetBatchTweetsResponse) RawJSON() string { return r.JSON.raw }
func (r *XUserGetBatchTweetsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A page of user IDs without profiles. Follow next_cursor while the response
// reports more pages.
type XUserGetFollowerIDsResponse struct {
	HasNextPage bool `json:"has_next_page" api:"required"`
	// X user IDs, newest relationship first.
	IDs        []string `json:"ids" api:"required"`
	NextCursor string   `json:"next_cursor" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		HasNextPage respjson.Field
		IDs         respjson.Field
		NextCursor  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XUserGetFollowerIDsResponse) RawJSON() string { return r.JSON.raw }
func (r *XUserGetFollowerIDsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A page of user IDs without profiles. Follow next_cursor while the response
// reports more pages.
type XUserGetFollowingIDsResponse struct {
	HasNextPage bool `json:"has_next_page" api:"required"`
	// X user IDs, newest relationship first.
	IDs        []string `json:"ids" api:"required"`
	NextCursor string   `json:"next_cursor" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		HasNextPage respjson.Field
		IDs         respjson.Field
		NextCursor  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XUserGetFollowingIDsResponse) RawJSON() string { return r.JSON.raw }
func (r *XUserGetFollowingIDsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type XUserGetInsightsResponse struct {
	// Engagement insights of 1 profile. Engagements are likes, replies, reposts and
	// quotes. A rate is null when nothing divides it, such as a profile without posts
	// of its own.
	Insights XUserGetInsightsResponseInsights `json:"insights" api:"required"`
	// Says why the report has counts only.
	Message string `json:"message"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Insights    respjson.Field
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XUserGetInsightsResponse) RawJSON() string { return r.JSON.raw }
func (r *XUserGetInsightsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Engagement insights of 1 profile. Engagements are likes, replies, reposts and
// quotes. A rate is null when nothing divides it, such as a profile without posts
// of its own.
type XUserGetInsightsResponseInsights struct {
	// Whole days since the account joined X.
	AccountAgeDays float64 `json:"accountAgeDays" api:"required"`
	// Each count per post of the sample. Null without posts.
	Averages XUserGetInsightsResponseInsightsAverages `json:"averages" api:"required"`
	// Engagements as a percent of views across the sample.
	EngagementPerView float64 `json:"engagementPerView" api:"required"`
	// Average engagements per post, as a percent of followers.
	EngagementRate float64 `json:"engagementRate" api:"required"`
	// Followers.
	Followers float64 `json:"followers" api:"required"`
	// Accounts the user follows.
	Following float64 `json:"following" api:"required"`
	// Posts the user has made.
	Posts float64 `json:"posts" api:"required"`
	// Sampled posts per day, from the oldest one until now.
	PostsPerDay float64 `json:"postsPerDay" api:"required"`
	// True when only the account's followers see its posts. The report then has its
	// counts only.
	PostsPrivate bool `json:"postsPrivate" api:"required"`
	// The recent posts the averages read.
	Sample XUserGetInsightsResponseInsightsSample `json:"sample" api:"required"`
	// The sampled post with the most engagements.
	TopPostID string `json:"topPostId" api:"required"`
	// ID of the profile the insights describe.
	UserID string `json:"userId" api:"required"`
	// Username of that profile.
	UserName string `json:"userName" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccountAgeDays    respjson.Field
		Averages          respjson.Field
		EngagementPerView respjson.Field
		EngagementRate    respjson.Field
		Followers         respjson.Field
		Following         respjson.Field
		Posts             respjson.Field
		PostsPerDay       respjson.Field
		PostsPrivate      respjson.Field
		Sample            respjson.Field
		TopPostID         respjson.Field
		UserID            respjson.Field
		UserName          respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XUserGetInsightsResponseInsights) RawJSON() string { return r.JSON.raw }
func (r *XUserGetInsightsResponseInsights) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Each count per post of the sample. Null without posts.
type XUserGetInsightsResponseInsightsAverages struct {
	// Bookmarks.
	Bookmarks float64 `json:"bookmarks" api:"required"`
	// Likes, replies, reposts and quotes per post.
	Engagements float64 `json:"engagements" api:"required"`
	// Likes.
	Likes float64 `json:"likes" api:"required"`
	// Quotes.
	Quotes float64 `json:"quotes" api:"required"`
	// Replies.
	Replies float64 `json:"replies" api:"required"`
	// Reposts.
	Reposts float64 `json:"reposts" api:"required"`
	// Views.
	Views float64 `json:"views" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Bookmarks   respjson.Field
		Engagements respjson.Field
		Likes       respjson.Field
		Quotes      respjson.Field
		Replies     respjson.Field
		Reposts     respjson.Field
		Views       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XUserGetInsightsResponseInsightsAverages) RawJSON() string { return r.JSON.raw }
func (r *XUserGetInsightsResponseInsightsAverages) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The recent posts the averages read.
type XUserGetInsightsResponseInsightsSample struct {
	// When the newest of them went up.
	NewestAt time.Time `json:"newestAt" api:"required" format:"date-time"`
	// When the oldest of them went up.
	OldestAt time.Time `json:"oldestAt" api:"required" format:"date-time"`
	// Posts read.
	Posts float64 `json:"posts" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		NewestAt    respjson.Field
		OldestAt    respjson.Field
		Posts       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XUserGetInsightsResponseInsightsSample) RawJSON() string { return r.JSON.raw }
func (r *XUserGetInsightsResponseInsightsSample) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// 1 page of public X Lists. Follow next_cursor while the response reports more
// pages. Each returned List costs 1 credit.
type XUserGetListMembershipsResponse struct {
	HasNextPage bool                                  `json:"has_next_page" api:"required"`
	Lists       []XUserGetListMembershipsResponseList `json:"lists" api:"required"`
	NextCursor  string                                `json:"next_cursor" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		HasNextPage respjson.Field
		Lists       respjson.Field
		NextCursor  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XUserGetListMembershipsResponse) RawJSON() string { return r.JSON.raw }
func (r *XUserGetListMembershipsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Public X List with its owner's public profile. Leaves out what X tells the
// reading account, such as whether it follows the List.
type XUserGetListMembershipsResponseList struct {
	// Numeric ID of the List.
	ID string `json:"id" api:"required"`
	// Name the owner gave the List.
	Name string `json:"name" api:"required"`
	// URL of the banner the owner uploaded. Absent when the owner uploaded none.
	BannerURL string `json:"bannerUrl"`
	// When the owner made the List.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Banner the owner uploaded.
	CustomBanner XUserGetListMembershipsResponseListCustomBanner `json:"customBanner"`
	// Banner X shows when the owner uploaded none. X picks it per reading account, so
	// 2 calls can return different ones.
	DefaultBanner XUserGetListMembershipsResponseListDefaultBanner `json:"defaultBanner"`
	// About text of the List. Empty when the owner wrote none.
	Description string `json:"description"`
	// Avatar URLs of a few List members, which X may pick per request.
	FacepileURLs []string `json:"facepileUrls"`
	// Accounts on the List.
	MemberCount int64 `json:"memberCount"`
	// Member count as X words it, such as 35 members.
	MembersContext string `json:"membersContext"`
	// Public profile of the account that owns the List.
	Owner shared.UserProfile `json:"owner"`
	// Accounts that follow the List, which X calls followers.
	SubscriberCount int64 `json:"subscriberCount"`
	// List page on x.com.
	URL string `json:"url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		Name            respjson.Field
		BannerURL       respjson.Field
		CreatedAt       respjson.Field
		CustomBanner    respjson.Field
		DefaultBanner   respjson.Field
		Description     respjson.Field
		FacepileURLs    respjson.Field
		MemberCount     respjson.Field
		MembersContext  respjson.Field
		Owner           respjson.Field
		SubscriberCount respjson.Field
		URL             respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XUserGetListMembershipsResponseList) RawJSON() string { return r.JSON.raw }
func (r *XUserGetListMembershipsResponseList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Banner the owner uploaded.
type XUserGetListMembershipsResponseListCustomBanner struct {
	// Banner image URL.
	URL string `json:"url" api:"required"`
	// Banner height in pixels.
	Height int64 `json:"height"`
	// Media ID of an uploaded banner.
	MediaID string `json:"mediaId"`
	// Stable media key of an uploaded banner.
	MediaKey string `json:"mediaKey"`
	// Dominant colors of an uploaded banner and their shares.
	Palette []XUserGetListMembershipsResponseListCustomBannerPalette `json:"palette"`
	// Area X marks as the subject of the banner, in pixels.
	SalientRect XUserGetListMembershipsResponseListCustomBannerSalientRect `json:"salientRect"`
	// Banner width in pixels.
	Width int64 `json:"width"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		URL         respjson.Field
		Height      respjson.Field
		MediaID     respjson.Field
		MediaKey    respjson.Field
		Palette     respjson.Field
		SalientRect respjson.Field
		Width       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XUserGetListMembershipsResponseListCustomBanner) RawJSON() string { return r.JSON.raw }
func (r *XUserGetListMembershipsResponseListCustomBanner) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type XUserGetListMembershipsResponseListCustomBannerPalette struct {
	// Share of the banner
	Percentage float64 `json:"percentage"`
	// The color's red, green and blue values.
	Rgb XUserGetListMembershipsResponseListCustomBannerPaletteRgb `json:"rgb"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Percentage  respjson.Field
		Rgb         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XUserGetListMembershipsResponseListCustomBannerPalette) RawJSON() string { return r.JSON.raw }
func (r *XUserGetListMembershipsResponseListCustomBannerPalette) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The color's red, green and blue values.
type XUserGetListMembershipsResponseListCustomBannerPaletteRgb struct {
	Blue  int64 `json:"blue"`
	Green int64 `json:"green"`
	Red   int64 `json:"red"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Blue        respjson.Field
		Green       respjson.Field
		Red         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XUserGetListMembershipsResponseListCustomBannerPaletteRgb) RawJSON() string {
	return r.JSON.raw
}
func (r *XUserGetListMembershipsResponseListCustomBannerPaletteRgb) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Area X marks as the subject of the banner, in pixels.
type XUserGetListMembershipsResponseListCustomBannerSalientRect struct {
	// Height of the marked area.
	Height int64 `json:"height"`
	// Pixels from the banner's left.
	Left int64 `json:"left"`
	// Pixels from the banner's top.
	Top int64 `json:"top"`
	// Width of the marked area.
	Width int64 `json:"width"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Height      respjson.Field
		Left        respjson.Field
		Top         respjson.Field
		Width       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XUserGetListMembershipsResponseListCustomBannerSalientRect) RawJSON() string {
	return r.JSON.raw
}
func (r *XUserGetListMembershipsResponseListCustomBannerSalientRect) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Banner X shows when the owner uploaded none. X picks it per reading account, so
// 2 calls can return different ones.
type XUserGetListMembershipsResponseListDefaultBanner struct {
	// Banner image URL.
	URL string `json:"url" api:"required"`
	// Banner height in pixels.
	Height int64 `json:"height"`
	// Media ID of an uploaded banner.
	MediaID string `json:"mediaId"`
	// Stable media key of an uploaded banner.
	MediaKey string `json:"mediaKey"`
	// Dominant colors of an uploaded banner and their shares.
	Palette []XUserGetListMembershipsResponseListDefaultBannerPalette `json:"palette"`
	// Area X marks as the subject of the banner, in pixels.
	SalientRect XUserGetListMembershipsResponseListDefaultBannerSalientRect `json:"salientRect"`
	// Banner width in pixels.
	Width int64 `json:"width"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		URL         respjson.Field
		Height      respjson.Field
		MediaID     respjson.Field
		MediaKey    respjson.Field
		Palette     respjson.Field
		SalientRect respjson.Field
		Width       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XUserGetListMembershipsResponseListDefaultBanner) RawJSON() string { return r.JSON.raw }
func (r *XUserGetListMembershipsResponseListDefaultBanner) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type XUserGetListMembershipsResponseListDefaultBannerPalette struct {
	// Share of the banner
	Percentage float64 `json:"percentage"`
	// The color's red, green and blue values.
	Rgb XUserGetListMembershipsResponseListDefaultBannerPaletteRgb `json:"rgb"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Percentage  respjson.Field
		Rgb         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XUserGetListMembershipsResponseListDefaultBannerPalette) RawJSON() string { return r.JSON.raw }
func (r *XUserGetListMembershipsResponseListDefaultBannerPalette) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The color's red, green and blue values.
type XUserGetListMembershipsResponseListDefaultBannerPaletteRgb struct {
	Blue  int64 `json:"blue"`
	Green int64 `json:"green"`
	Red   int64 `json:"red"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Blue        respjson.Field
		Green       respjson.Field
		Red         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XUserGetListMembershipsResponseListDefaultBannerPaletteRgb) RawJSON() string {
	return r.JSON.raw
}
func (r *XUserGetListMembershipsResponseListDefaultBannerPaletteRgb) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Area X marks as the subject of the banner, in pixels.
type XUserGetListMembershipsResponseListDefaultBannerSalientRect struct {
	// Height of the marked area.
	Height int64 `json:"height"`
	// Pixels from the banner's left.
	Left int64 `json:"left"`
	// Pixels from the banner's top.
	Top int64 `json:"top"`
	// Width of the marked area.
	Width int64 `json:"width"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Height      respjson.Field
		Left        respjson.Field
		Top         respjson.Field
		Width       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XUserGetListMembershipsResponseListDefaultBannerSalientRect) RawJSON() string {
	return r.JSON.raw
}
func (r *XUserGetListMembershipsResponseListDefaultBannerSalientRect) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// 1 page of public X Lists. Follow next_cursor while the response reports more
// pages. Each returned List costs 1 credit.
type XUserGetListsResponse struct {
	HasNextPage bool                        `json:"has_next_page" api:"required"`
	Lists       []XUserGetListsResponseList `json:"lists" api:"required"`
	NextCursor  string                      `json:"next_cursor" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		HasNextPage respjson.Field
		Lists       respjson.Field
		NextCursor  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XUserGetListsResponse) RawJSON() string { return r.JSON.raw }
func (r *XUserGetListsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Public X List with its owner's public profile. Leaves out what X tells the
// reading account, such as whether it follows the List.
type XUserGetListsResponseList struct {
	// Numeric ID of the List.
	ID string `json:"id" api:"required"`
	// Name the owner gave the List.
	Name string `json:"name" api:"required"`
	// URL of the banner the owner uploaded. Absent when the owner uploaded none.
	BannerURL string `json:"bannerUrl"`
	// When the owner made the List.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Banner the owner uploaded.
	CustomBanner XUserGetListsResponseListCustomBanner `json:"customBanner"`
	// Banner X shows when the owner uploaded none. X picks it per reading account, so
	// 2 calls can return different ones.
	DefaultBanner XUserGetListsResponseListDefaultBanner `json:"defaultBanner"`
	// About text of the List. Empty when the owner wrote none.
	Description string `json:"description"`
	// Avatar URLs of a few List members, which X may pick per request.
	FacepileURLs []string `json:"facepileUrls"`
	// Accounts on the List.
	MemberCount int64 `json:"memberCount"`
	// Member count as X words it, such as 35 members.
	MembersContext string `json:"membersContext"`
	// Public profile of the account that owns the List.
	Owner shared.UserProfile `json:"owner"`
	// Accounts that follow the List, which X calls followers.
	SubscriberCount int64 `json:"subscriberCount"`
	// List page on x.com.
	URL string `json:"url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		Name            respjson.Field
		BannerURL       respjson.Field
		CreatedAt       respjson.Field
		CustomBanner    respjson.Field
		DefaultBanner   respjson.Field
		Description     respjson.Field
		FacepileURLs    respjson.Field
		MemberCount     respjson.Field
		MembersContext  respjson.Field
		Owner           respjson.Field
		SubscriberCount respjson.Field
		URL             respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XUserGetListsResponseList) RawJSON() string { return r.JSON.raw }
func (r *XUserGetListsResponseList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Banner the owner uploaded.
type XUserGetListsResponseListCustomBanner struct {
	// Banner image URL.
	URL string `json:"url" api:"required"`
	// Banner height in pixels.
	Height int64 `json:"height"`
	// Media ID of an uploaded banner.
	MediaID string `json:"mediaId"`
	// Stable media key of an uploaded banner.
	MediaKey string `json:"mediaKey"`
	// Dominant colors of an uploaded banner and their shares.
	Palette []XUserGetListsResponseListCustomBannerPalette `json:"palette"`
	// Area X marks as the subject of the banner, in pixels.
	SalientRect XUserGetListsResponseListCustomBannerSalientRect `json:"salientRect"`
	// Banner width in pixels.
	Width int64 `json:"width"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		URL         respjson.Field
		Height      respjson.Field
		MediaID     respjson.Field
		MediaKey    respjson.Field
		Palette     respjson.Field
		SalientRect respjson.Field
		Width       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XUserGetListsResponseListCustomBanner) RawJSON() string { return r.JSON.raw }
func (r *XUserGetListsResponseListCustomBanner) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type XUserGetListsResponseListCustomBannerPalette struct {
	// Share of the banner
	Percentage float64 `json:"percentage"`
	// The color's red, green and blue values.
	Rgb XUserGetListsResponseListCustomBannerPaletteRgb `json:"rgb"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Percentage  respjson.Field
		Rgb         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XUserGetListsResponseListCustomBannerPalette) RawJSON() string { return r.JSON.raw }
func (r *XUserGetListsResponseListCustomBannerPalette) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The color's red, green and blue values.
type XUserGetListsResponseListCustomBannerPaletteRgb struct {
	Blue  int64 `json:"blue"`
	Green int64 `json:"green"`
	Red   int64 `json:"red"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Blue        respjson.Field
		Green       respjson.Field
		Red         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XUserGetListsResponseListCustomBannerPaletteRgb) RawJSON() string { return r.JSON.raw }
func (r *XUserGetListsResponseListCustomBannerPaletteRgb) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Area X marks as the subject of the banner, in pixels.
type XUserGetListsResponseListCustomBannerSalientRect struct {
	// Height of the marked area.
	Height int64 `json:"height"`
	// Pixels from the banner's left.
	Left int64 `json:"left"`
	// Pixels from the banner's top.
	Top int64 `json:"top"`
	// Width of the marked area.
	Width int64 `json:"width"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Height      respjson.Field
		Left        respjson.Field
		Top         respjson.Field
		Width       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XUserGetListsResponseListCustomBannerSalientRect) RawJSON() string { return r.JSON.raw }
func (r *XUserGetListsResponseListCustomBannerSalientRect) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Banner X shows when the owner uploaded none. X picks it per reading account, so
// 2 calls can return different ones.
type XUserGetListsResponseListDefaultBanner struct {
	// Banner image URL.
	URL string `json:"url" api:"required"`
	// Banner height in pixels.
	Height int64 `json:"height"`
	// Media ID of an uploaded banner.
	MediaID string `json:"mediaId"`
	// Stable media key of an uploaded banner.
	MediaKey string `json:"mediaKey"`
	// Dominant colors of an uploaded banner and their shares.
	Palette []XUserGetListsResponseListDefaultBannerPalette `json:"palette"`
	// Area X marks as the subject of the banner, in pixels.
	SalientRect XUserGetListsResponseListDefaultBannerSalientRect `json:"salientRect"`
	// Banner width in pixels.
	Width int64 `json:"width"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		URL         respjson.Field
		Height      respjson.Field
		MediaID     respjson.Field
		MediaKey    respjson.Field
		Palette     respjson.Field
		SalientRect respjson.Field
		Width       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XUserGetListsResponseListDefaultBanner) RawJSON() string { return r.JSON.raw }
func (r *XUserGetListsResponseListDefaultBanner) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type XUserGetListsResponseListDefaultBannerPalette struct {
	// Share of the banner
	Percentage float64 `json:"percentage"`
	// The color's red, green and blue values.
	Rgb XUserGetListsResponseListDefaultBannerPaletteRgb `json:"rgb"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Percentage  respjson.Field
		Rgb         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XUserGetListsResponseListDefaultBannerPalette) RawJSON() string { return r.JSON.raw }
func (r *XUserGetListsResponseListDefaultBannerPalette) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The color's red, green and blue values.
type XUserGetListsResponseListDefaultBannerPaletteRgb struct {
	Blue  int64 `json:"blue"`
	Green int64 `json:"green"`
	Red   int64 `json:"red"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Blue        respjson.Field
		Green       respjson.Field
		Red         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XUserGetListsResponseListDefaultBannerPaletteRgb) RawJSON() string { return r.JSON.raw }
func (r *XUserGetListsResponseListDefaultBannerPaletteRgb) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Area X marks as the subject of the banner, in pixels.
type XUserGetListsResponseListDefaultBannerSalientRect struct {
	// Height of the marked area.
	Height int64 `json:"height"`
	// Pixels from the banner's left.
	Left int64 `json:"left"`
	// Pixels from the banner's top.
	Top int64 `json:"top"`
	// Width of the marked area.
	Width int64 `json:"width"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Height      respjson.Field
		Left        respjson.Field
		Top         respjson.Field
		Width       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XUserGetListsResponseListDefaultBannerSalientRect) RawJSON() string { return r.JSON.raw }
func (r *XUserGetListsResponseListDefaultBannerSalientRect) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Says whether a user broadcasts live video on X now. `broadcast` is their latest
// broadcast, live or ended, and null when X shows none.
type XUserGetLiveResponse struct {
	// Latest broadcast, live or ended. Null when X shows none.
	Broadcast XUserGetLiveResponseBroadcast `json:"broadcast" api:"required"`
	// True while the user's latest broadcast is running.
	Live bool `json:"live" api:"required"`
	// ID of the user.
	UserID string `json:"userId" api:"required"`
	// Username of the user.
	UserName string `json:"userName" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Broadcast   respjson.Field
		Live        respjson.Field
		UserID      respjson.Field
		UserName    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XUserGetLiveResponse) RawJSON() string { return r.JSON.raw }
func (r *XUserGetLiveResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Latest broadcast, live or ended. Null when X shows none.
type XUserGetLiveResponseBroadcast struct {
	// Broadcast ID
	ID string `json:"id" api:"required"`
	// Camera rotation in degrees
	CameraRotation int64 `json:"cameraRotation"`
	// X's code for who may chat
	ChatOption int64 `json:"chatOption"`
	// Public X profile.
	Creator shared.UserProfile `json:"creator"`
	// When the broadcast ended
	EndedAt time.Time `json:"endedAt" format:"date-time"`
	// Video height in pixels
	Height int64 `json:"height"`
	// Whether X offers a recording of the broadcast
	IsAvailableForReplay bool `json:"isAvailableForReplay"`
	// Whether the video stream runs in high-latency mode
	IsHighLatency bool `json:"isHighLatency"`
	// Whether X marks the chat private
	IsPrivateChat bool `json:"isPrivateChat"`
	// Whether the broadcaster edited the title of the replay
	IsTitleEdited bool `json:"isTitleEdited"`
	// When the video stream last reached X
	LastPingAt time.Time `json:"lastPingAt" format:"date-time"`
	// X's media key for the video
	MediaKey string `json:"mediaKey"`
	// Broadcaster, as X's live video service names the account
	PeriscopeUser XUserGetLiveResponseBroadcastPeriscopeUser `json:"periscopeUser"`
	// Image X shows before the broadcast starts
	PreLiveSlateURL string `json:"preLiveSlateUrl" format:"uri"`
	// When the broadcast was set to start
	ScheduledStart time.Time `json:"scheduledStart" format:"date-time"`
	// What produced the video stream, as X names it
	Source string `json:"source"`
	// When the broadcast went live
	StartedAt time.Time `json:"startedAt" format:"date-time"`
	// State, such as NotStarted, Running or Ended, as X names it
	State string `json:"state"`
	// Broadcast title
	Title string `json:"title"`
	// Accounts that watched, live or after
	TotalWatched int64 `json:"totalWatched"`
	// Accounts watching, as X counts them
	TotalWatching int64 `json:"totalWatching"`
	// Search tweet with author. Zero metrics may mean X omitted counts.
	Tweet shared.SearchTweet `json:"tweet"`
	// ID of the post that announces the broadcast
	TweetID string `json:"tweetId"`
	// X's revision number of the broadcast
	Version int64 `json:"version"`
	// Video width in pixels
	Width int64 `json:"width"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                   respjson.Field
		CameraRotation       respjson.Field
		ChatOption           respjson.Field
		Creator              respjson.Field
		EndedAt              respjson.Field
		Height               respjson.Field
		IsAvailableForReplay respjson.Field
		IsHighLatency        respjson.Field
		IsPrivateChat        respjson.Field
		IsTitleEdited        respjson.Field
		LastPingAt           respjson.Field
		MediaKey             respjson.Field
		PeriscopeUser        respjson.Field
		PreLiveSlateURL      respjson.Field
		ScheduledStart       respjson.Field
		Source               respjson.Field
		StartedAt            respjson.Field
		State                respjson.Field
		Title                respjson.Field
		TotalWatched         respjson.Field
		TotalWatching        respjson.Field
		Tweet                respjson.Field
		TweetID              respjson.Field
		Version              respjson.Field
		Width                respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XUserGetLiveResponseBroadcast) RawJSON() string { return r.JSON.raw }
func (r *XUserGetLiveResponseBroadcast) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Broadcaster, as X's live video service names the account
type XUserGetLiveResponseBroadcastPeriscopeUser struct {
	// Live video account ID
	ID string `json:"id"`
	// Live video display name
	DisplayName string `json:"displayName"`
	// Live video avatar URL
	ProfilePicture string `json:"profilePicture" format:"uri"`
	// The broadcaster's avatar at full size
	ProfilePictureOriginal string `json:"profilePictureOriginal" format:"uri"`
	// Live video username
	Username string `json:"username"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                     respjson.Field
		DisplayName            respjson.Field
		ProfilePicture         respjson.Field
		ProfilePictureOriginal respjson.Field
		Username               respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XUserGetLiveResponseBroadcastPeriscopeUser) RawJSON() string { return r.JSON.raw }
func (r *XUserGetLiveResponseBroadcastPeriscopeUser) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type XUserTranslateBioResponse struct {
	// Null when the profile has no bio.
	Translation XUserTranslateBioResponseTranslation `json:"translation" api:"required"`
	// Says why translation is null.
	Message string `json:"message"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Translation respjson.Field
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XUserTranslateBioResponse) RawJSON() string { return r.JSON.raw }
func (r *XUserTranslateBioResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Null when the profile has no bio.
type XUserTranslateBioResponseTranslation struct {
	// Lists hashtags, symbols, links, and mentions from tweet text.
	Entities XUserTranslateBioResponseTranslationEntities `json:"entities" api:"required"`
	// Language code of the translated bio, in lower case.
	Language string `json:"language" api:"required"`
	// Translated bio.
	Text string `json:"text" api:"required"`
	// ID of the account whose bio X translated.
	UserID string `json:"userId" api:"required"`
	// Username of that account.
	UserName string `json:"userName" api:"required"`
	// Language code X detects in the bio, when it names one.
	SourceLanguage string `json:"sourceLanguage"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Entities       respjson.Field
		Language       respjson.Field
		Text           respjson.Field
		UserID         respjson.Field
		UserName       respjson.Field
		SourceLanguage respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XUserTranslateBioResponseTranslation) RawJSON() string { return r.JSON.raw }
func (r *XUserTranslateBioResponseTranslation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Lists hashtags, symbols, links, and mentions from tweet text.
type XUserTranslateBioResponseTranslationEntities struct {
	Hashtags     []XUserTranslateBioResponseTranslationEntitiesHashtag     `json:"hashtags"`
	Smarttags    []XUserTranslateBioResponseTranslationEntitiesSmarttag    `json:"smarttags"`
	Symbols      []XUserTranslateBioResponseTranslationEntitiesSymbol      `json:"symbols"`
	Timestamps   []XUserTranslateBioResponseTranslationEntitiesTimestamp   `json:"timestamps"`
	URLs         []XUserTranslateBioResponseTranslationEntitiesURL         `json:"urls"`
	UserMentions []XUserTranslateBioResponseTranslationEntitiesUserMention `json:"user_mentions"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Hashtags     respjson.Field
		Smarttags    respjson.Field
		Symbols      respjson.Field
		Timestamps   respjson.Field
		URLs         respjson.Field
		UserMentions respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XUserTranslateBioResponseTranslationEntities) RawJSON() string { return r.JSON.raw }
func (r *XUserTranslateBioResponseTranslationEntities) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Hashtag text and its source offsets in a tweet.
type XUserTranslateBioResponseTranslationEntitiesHashtag struct {
	Text    string  `json:"text" api:"required"`
	Indices []int64 `json:"indices"`
	// X's SmartTag for a cashtag, smarttag or hashtag. Xquik leaves out X's trading
	// platforms, which differ by reader.
	Tag XUserTranslateBioResponseTranslationEntitiesHashtagTag `json:"tag"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Text        respjson.Field
		Indices     respjson.Field
		Tag         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XUserTranslateBioResponseTranslationEntitiesHashtag) RawJSON() string { return r.JSON.raw }
func (r *XUserTranslateBioResponseTranslationEntitiesHashtag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// X's SmartTag for a cashtag, smarttag or hashtag. Xquik leaves out X's trading
// platforms, which differ by reader.
type XUserTranslateBioResponseTranslationEntitiesHashtagTag struct {
	// X's ID for the SmartTag.
	ID   string                                                     `json:"id"`
	Info XUserTranslateBioResponseTranslationEntitiesHashtagTagInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XUserTranslateBioResponseTranslationEntitiesHashtagTag) RawJSON() string { return r.JSON.raw }
func (r *XUserTranslateBioResponseTranslationEntitiesHashtagTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type XUserTranslateBioResponseTranslationEntitiesHashtagTagInfo struct {
	Info XUserTranslateBioResponseTranslationEntitiesHashtagTagInfoInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XUserTranslateBioResponseTranslationEntitiesHashtagTagInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *XUserTranslateBioResponseTranslationEntitiesHashtagTagInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type XUserTranslateBioResponseTranslationEntitiesHashtagTagInfoInfo struct {
	// Whether X marks the SmartTag active.
	Active bool   `json:"active"`
	Name   string `json:"name"`
	Ticker string `json:"ticker"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Active      respjson.Field
		Name        respjson.Field
		Ticker      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XUserTranslateBioResponseTranslationEntitiesHashtagTagInfoInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *XUserTranslateBioResponseTranslationEntitiesHashtagTagInfoInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indexed smart-tag, cashtag, or video timestamp metadata.
type XUserTranslateBioResponseTranslationEntitiesSmarttag struct {
	Indices []int64 `json:"indices"`
	Seconds float64 `json:"seconds"`
	// X's SmartTag for a cashtag, smarttag or hashtag. Xquik leaves out X's trading
	// platforms, which differ by reader.
	Tag  XUserTranslateBioResponseTranslationEntitiesSmarttagTag `json:"tag"`
	Text string                                                  `json:"text"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Indices     respjson.Field
		Seconds     respjson.Field
		Tag         respjson.Field
		Text        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XUserTranslateBioResponseTranslationEntitiesSmarttag) RawJSON() string { return r.JSON.raw }
func (r *XUserTranslateBioResponseTranslationEntitiesSmarttag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// X's SmartTag for a cashtag, smarttag or hashtag. Xquik leaves out X's trading
// platforms, which differ by reader.
type XUserTranslateBioResponseTranslationEntitiesSmarttagTag struct {
	// X's ID for the SmartTag.
	ID   string                                                      `json:"id"`
	Info XUserTranslateBioResponseTranslationEntitiesSmarttagTagInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XUserTranslateBioResponseTranslationEntitiesSmarttagTag) RawJSON() string { return r.JSON.raw }
func (r *XUserTranslateBioResponseTranslationEntitiesSmarttagTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type XUserTranslateBioResponseTranslationEntitiesSmarttagTagInfo struct {
	Info XUserTranslateBioResponseTranslationEntitiesSmarttagTagInfoInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XUserTranslateBioResponseTranslationEntitiesSmarttagTagInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *XUserTranslateBioResponseTranslationEntitiesSmarttagTagInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type XUserTranslateBioResponseTranslationEntitiesSmarttagTagInfoInfo struct {
	// Whether X marks the SmartTag active.
	Active bool   `json:"active"`
	Name   string `json:"name"`
	Ticker string `json:"ticker"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Active      respjson.Field
		Name        respjson.Field
		Ticker      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XUserTranslateBioResponseTranslationEntitiesSmarttagTagInfoInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *XUserTranslateBioResponseTranslationEntitiesSmarttagTagInfoInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indexed smart-tag, cashtag, or video timestamp metadata.
type XUserTranslateBioResponseTranslationEntitiesSymbol struct {
	Indices []int64 `json:"indices"`
	Seconds float64 `json:"seconds"`
	// X's SmartTag for a cashtag, smarttag or hashtag. Xquik leaves out X's trading
	// platforms, which differ by reader.
	Tag  XUserTranslateBioResponseTranslationEntitiesSymbolTag `json:"tag"`
	Text string                                                `json:"text"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Indices     respjson.Field
		Seconds     respjson.Field
		Tag         respjson.Field
		Text        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XUserTranslateBioResponseTranslationEntitiesSymbol) RawJSON() string { return r.JSON.raw }
func (r *XUserTranslateBioResponseTranslationEntitiesSymbol) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// X's SmartTag for a cashtag, smarttag or hashtag. Xquik leaves out X's trading
// platforms, which differ by reader.
type XUserTranslateBioResponseTranslationEntitiesSymbolTag struct {
	// X's ID for the SmartTag.
	ID   string                                                    `json:"id"`
	Info XUserTranslateBioResponseTranslationEntitiesSymbolTagInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XUserTranslateBioResponseTranslationEntitiesSymbolTag) RawJSON() string { return r.JSON.raw }
func (r *XUserTranslateBioResponseTranslationEntitiesSymbolTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type XUserTranslateBioResponseTranslationEntitiesSymbolTagInfo struct {
	Info XUserTranslateBioResponseTranslationEntitiesSymbolTagInfoInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XUserTranslateBioResponseTranslationEntitiesSymbolTagInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *XUserTranslateBioResponseTranslationEntitiesSymbolTagInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type XUserTranslateBioResponseTranslationEntitiesSymbolTagInfoInfo struct {
	// Whether X marks the SmartTag active.
	Active bool   `json:"active"`
	Name   string `json:"name"`
	Ticker string `json:"ticker"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Active      respjson.Field
		Name        respjson.Field
		Ticker      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XUserTranslateBioResponseTranslationEntitiesSymbolTagInfoInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *XUserTranslateBioResponseTranslationEntitiesSymbolTagInfoInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indexed smart-tag, cashtag, or video timestamp metadata.
type XUserTranslateBioResponseTranslationEntitiesTimestamp struct {
	Indices []int64 `json:"indices"`
	Seconds float64 `json:"seconds"`
	// X's SmartTag for a cashtag, smarttag or hashtag. Xquik leaves out X's trading
	// platforms, which differ by reader.
	Tag  XUserTranslateBioResponseTranslationEntitiesTimestampTag `json:"tag"`
	Text string                                                   `json:"text"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Indices     respjson.Field
		Seconds     respjson.Field
		Tag         respjson.Field
		Text        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XUserTranslateBioResponseTranslationEntitiesTimestamp) RawJSON() string { return r.JSON.raw }
func (r *XUserTranslateBioResponseTranslationEntitiesTimestamp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// X's SmartTag for a cashtag, smarttag or hashtag. Xquik leaves out X's trading
// platforms, which differ by reader.
type XUserTranslateBioResponseTranslationEntitiesTimestampTag struct {
	// X's ID for the SmartTag.
	ID   string                                                       `json:"id"`
	Info XUserTranslateBioResponseTranslationEntitiesTimestampTagInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XUserTranslateBioResponseTranslationEntitiesTimestampTag) RawJSON() string { return r.JSON.raw }
func (r *XUserTranslateBioResponseTranslationEntitiesTimestampTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type XUserTranslateBioResponseTranslationEntitiesTimestampTagInfo struct {
	Info XUserTranslateBioResponseTranslationEntitiesTimestampTagInfoInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XUserTranslateBioResponseTranslationEntitiesTimestampTagInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *XUserTranslateBioResponseTranslationEntitiesTimestampTagInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type XUserTranslateBioResponseTranslationEntitiesTimestampTagInfoInfo struct {
	// Whether X marks the SmartTag active.
	Active bool   `json:"active"`
	Name   string `json:"name"`
	Ticker string `json:"ticker"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Active      respjson.Field
		Name        respjson.Field
		Ticker      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XUserTranslateBioResponseTranslationEntitiesTimestampTagInfoInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *XUserTranslateBioResponseTranslationEntitiesTimestampTagInfoInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Shortened, display, and expanded URLs from tweet text.
type XUserTranslateBioResponseTranslationEntitiesURL struct {
	DisplayURL  string  `json:"display_url"`
	ExpandedURL string  `json:"expanded_url"`
	Indices     []int64 `json:"indices"`
	URL         string  `json:"url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DisplayURL  respjson.Field
		ExpandedURL respjson.Field
		Indices     respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XUserTranslateBioResponseTranslationEntitiesURL) RawJSON() string { return r.JSON.raw }
func (r *XUserTranslateBioResponseTranslationEntitiesURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Profile identity and source offsets of a mention.
type XUserTranslateBioResponseTranslationEntitiesUserMention struct {
	ScreenName string  `json:"screen_name" api:"required"`
	IDStr      string  `json:"id_str"`
	Indices    []int64 `json:"indices"`
	Name       string  `json:"name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ScreenName  respjson.Field
		IDStr       respjson.Field
		Indices     respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XUserTranslateBioResponseTranslationEntitiesUserMention) RawJSON() string { return r.JSON.raw }
func (r *XUserTranslateBioResponseTranslationEntitiesUserMention) UnmarshalJSON(data []byte) error {
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
	// Comma-separated X usernames, with or without @, or profile links such as
	// x.com/nasa (1-100 values), in place of ids. A name in another case is the same
	// account.
	Usernames param.Opt[string] `query:"usernames,omitzero" json:"-"`
	// Only return verified profiles.
	VerifiedOnly param.Opt[bool] `query:"verifiedOnly,omitzero" json:"-"`
	// Match the verification type exactly, ignoring case.
	VerifiedType param.Opt[string] `query:"verifiedType,omitzero" json:"-"`
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
	// Match any comma-separated or line-separated bio term, ignoring case.
	BioContains param.Opt[string] `query:"bioContains,omitzero" json:"-"`
	// Only return profiles X marks as blue verified. X marks many business &
	// government accounts blue too. Use verifiedType for 1 badge type.
	BlueVerifiedOnly param.Opt[bool] `query:"blueVerifiedOnly,omitzero" json:"-"`
	// Only return profiles that publish an email outside free webmail, such as Gmail.
	BusinessEmailOnly param.Opt[bool] `query:"businessEmailOnly,omitzero" json:"-"`
	// Cursor from the previous response. Xquik cursors resume automatic coverage.
	// Existing unprefixed cursors keep legacy standard behavior.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Only return profiles with an email at one of these domains. Separate domains
	// with commas. Subdomains match.
	EmailDomains param.Opt[string] `query:"emailDomains,omitzero" json:"-"`
	// Add each row's full profile, as GET /x/users/{id} returns it. Pages take several
	// seconds longer.
	EnrichProfiles param.Opt[bool] `query:"enrichProfiles,omitzero" json:"-"`
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
	// Legacy follower limit. Coverage accepts 1-10000. Prefer pageSize.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Match a location substring, ignoring case.
	LocationContains param.Opt[string] `query:"locationContains,omitzero" json:"-"`
	// Maximum follower count. Missing counts pass this maximum.
	MaxFollowers param.Opt[int64] `query:"maxFollowers,omitzero" json:"-"`
	// Profiles may follow at most this many accounts.
	MaxFollowing param.Opt[int64] `query:"maxFollowing,omitzero" json:"-"`
	// Most profiles 1 request reads, kept or dropped by the filters. The request then
	// answers the profiles it kept, the count it dropped & the cursor to read on. Each
	// request of a cursor may send its own value. Without it, a filtered read scans at
	// most 100,000 profiles of the list across all its pages, & the page that reaches
	// them ends the list with `scanLimit`. An unfiltered read has no scan limit.
	// Automatic mode only.
	MaxScan param.Opt[int64] `query:"maxScan,omitzero" json:"-"`
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
	// Maximum user profiles: automatic 300; standard 200. Sources return fewer
	// profiles. Follow next_cursor while the response reports more pages.
	PageSize param.Opt[int64] `query:"pageSize,omitzero" json:"-"`
	// Match a username substring, ignoring case.
	UsernameContains param.Opt[string] `query:"usernameContains,omitzero" json:"-"`
	// Only return verified profiles.
	VerifiedOnly param.Opt[bool] `query:"verifiedOnly,omitzero" json:"-"`
	// Match the verification type exactly, ignoring case.
	VerifiedType param.Opt[string] `query:"verifiedType,omitzero" json:"-"`
	// Omit mode for resumable maximum coverage. Standard keeps legacy pagination.
	// Coverage returns diagnostics once and rejects cursors.
	//
	// Any of "standard", "coverage".
	Mode XUserGetFollowersParamsMode `query:"mode,omitzero" json:"-"`
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

// Omit mode for resumable maximum coverage. Standard keeps legacy pagination.
// Coverage returns diagnostics once and rejects cursors.
type XUserGetFollowersParamsMode string

const (
	XUserGetFollowersParamsModeStandard XUserGetFollowersParamsMode = "standard"
	XUserGetFollowersParamsModeCoverage XUserGetFollowersParamsMode = "coverage"
)

type XUserGetFollowersYouKnowParams struct {
	// Match any comma-separated or line-separated bio term, ignoring case.
	BioContains param.Opt[string] `query:"bioContains,omitzero" json:"-"`
	// Only return profiles X marks as blue verified. X marks many business &
	// government accounts blue too. Use verifiedType for 1 badge type.
	BlueVerifiedOnly param.Opt[bool] `query:"blueVerifiedOnly,omitzero" json:"-"`
	// Only return profiles that publish an email outside free webmail, such as Gmail.
	BusinessEmailOnly param.Opt[bool] `query:"businessEmailOnly,omitzero" json:"-"`
	// Pagination cursor for followers-you-know
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
	// Maximum user profiles requested from this page (1-200, default 200). Source,
	// filters, or credits can return fewer profiles. Follow next_cursor while the
	// response reports more pages. Deprecated aliases remain accepted.
	PageSize param.Opt[int64] `query:"pageSize,omitzero" json:"-"`
	// Match a username substring, ignoring case.
	UsernameContains param.Opt[string] `query:"usernameContains,omitzero" json:"-"`
	// Only return verified profiles.
	VerifiedOnly param.Opt[bool] `query:"verifiedOnly,omitzero" json:"-"`
	// Match the verification type exactly, ignoring case.
	VerifiedType param.Opt[string] `query:"verifiedType,omitzero" json:"-"`
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
	// Match any comma-separated or line-separated bio term, ignoring case.
	BioContains param.Opt[string] `query:"bioContains,omitzero" json:"-"`
	// Only return profiles X marks as blue verified. X marks many business &
	// government accounts blue too. Use verifiedType for 1 badge type.
	BlueVerifiedOnly param.Opt[bool] `query:"blueVerifiedOnly,omitzero" json:"-"`
	// Only return profiles that publish an email outside free webmail, such as Gmail.
	BusinessEmailOnly param.Opt[bool] `query:"businessEmailOnly,omitzero" json:"-"`
	// Cursor from the previous response. Xquik cursors resume automatic coverage.
	// Existing unprefixed cursors keep legacy standard behavior.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Only return profiles with an email at one of these domains. Separate domains
	// with commas. Subdomains match.
	EmailDomains param.Opt[string] `query:"emailDomains,omitzero" json:"-"`
	// Add each row's full profile, as GET /x/users/{id} returns it. Pages take several
	// seconds longer.
	EnrichProfiles param.Opt[bool] `query:"enrichProfiles,omitzero" json:"-"`
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
	// Legacy page-size alias outside explicit coverage mode. Coverage accepts 1-10000.
	// Prefer pageSize.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Match a location substring, ignoring case.
	LocationContains param.Opt[string] `query:"locationContains,omitzero" json:"-"`
	// Maximum follower count. Missing counts pass this maximum.
	MaxFollowers param.Opt[int64] `query:"maxFollowers,omitzero" json:"-"`
	// Profiles may follow at most this many accounts.
	MaxFollowing param.Opt[int64] `query:"maxFollowing,omitzero" json:"-"`
	// Most profiles 1 request reads, kept or dropped by the filters. The request then
	// answers the profiles it kept, the count it dropped & the cursor to read on. Each
	// request of a cursor may send its own value. Without it, a filtered read scans at
	// most 100,000 profiles of the list across all its pages, & the page that reaches
	// them ends the list with `scanLimit`. An unfiltered read has no scan limit.
	// Automatic mode only.
	MaxScan param.Opt[int64] `query:"maxScan,omitzero" json:"-"`
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
	// Maximum user profiles: automatic 300; standard 200. Sources return fewer
	// profiles. Follow next_cursor while the response reports more pages.
	PageSize param.Opt[int64] `query:"pageSize,omitzero" json:"-"`
	// Match a username substring, ignoring case.
	UsernameContains param.Opt[string] `query:"usernameContains,omitzero" json:"-"`
	// Only return verified profiles.
	VerifiedOnly param.Opt[bool] `query:"verifiedOnly,omitzero" json:"-"`
	// Match the verification type exactly, ignoring case.
	VerifiedType param.Opt[string] `query:"verifiedType,omitzero" json:"-"`
	// Omit mode for resumable maximum coverage. Standard keeps legacy pagination.
	// Coverage returns diagnostics once and rejects cursors.
	//
	// Any of "standard", "coverage".
	Mode XUserGetFollowingParamsMode `query:"mode,omitzero" json:"-"`
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

// Omit mode for resumable maximum coverage. Standard keeps legacy pagination.
// Coverage returns diagnostics once and rejects cursors.
type XUserGetFollowingParamsMode string

const (
	XUserGetFollowingParamsModeStandard XUserGetFollowingParamsMode = "standard"
	XUserGetFollowingParamsModeCoverage XUserGetFollowingParamsMode = "coverage"
)

type XUserGetLikesParams struct {
	// Words or quoted phrases where any one can match. Separate with spaces, commas,
	// or lines.
	AnyWords param.Opt[string] `query:"anyWords,omitzero" json:"-"`
	// Only return tweets from Blue-verified authors.
	BlueVerifiedOnly param.Opt[bool] `query:"blueVerifiedOnly,omitzero" json:"-"`
	// Match the Tweet card name.
	CardName param.Opt[string] `query:"cardName,omitzero" json:"-"`
	// Cashtags separated by spaces, commas, or lines.
	Cashtags param.Opt[string] `query:"cashtags,omitzero" json:"-"`
	// Conversation ID filter.
	ConversationID param.Opt[string] `query:"conversationId,omitzero" json:"-"`
	// Pagination cursor for liked tweets
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Match this literal phrase, including any hyphens.
	ExactPhrase param.Opt[string] `query:"exactPhrase,omitzero" json:"-"`
	// Exclude a source application.
	ExcludeSource param.Opt[string] `query:"excludeSource,omitzero" json:"-"`
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
	// Match latitude, longitude, and radius.
	Geocode param.Opt[string] `query:"geocode,omitzero" json:"-"`
	// Hashtags separated by spaces, commas, or lines.
	Hashtags param.Opt[string] `query:"hashtags,omitzero" json:"-"`
	// Only replies to this tweet ID.
	InReplyToTweetID param.Opt[string] `query:"inReplyToTweetId,omitzero" json:"-"`
	// Filter by language. Alias `lang` is accepted.
	Language param.Opt[string] `query:"language,omitzero" json:"-"`
	// Maximum likes threshold. maxLikes is also accepted.
	MaxFaves param.Opt[int64] `query:"maxFaves,omitzero" json:"-"`
	// Return Tweets older than this Tweet ID.
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
	// Deprecated: Use MinLikes. MinFaves sends the minFaves query key, which the
	// API reads the same way as minLikes.
	MinFaves param.Opt[int64] `query:"minFaves,omitzero" json:"-"`
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
	// Match a place name.
	Near param.Opt[string] `query:"near,omitzero" json:"-"`
	// Only return news results.
	News param.Opt[bool] `query:"news,omitzero" json:"-"`
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
	// Enable the safe-search filter.
	Safe param.Opt[bool] `query:"safe,omitzero" json:"-"`
	// Start date in YYYY-MM-DD format.
	SinceDate param.Opt[time.Time] `query:"sinceDate,omitzero" format:"date" json:"-"`
	// Return Tweets newer than this Tweet ID.
	SinceID param.Opt[string] `query:"sinceId,omitzero" json:"-"`
	// Match the source application.
	Source param.Opt[string] `query:"source,omitzero" json:"-"`
	// Filter replies sent to a username.
	ToUser param.Opt[string] `query:"toUser,omitzero" json:"-"`
	// End date in YYYY-MM-DD format.
	UntilDate param.Opt[time.Time] `query:"untilDate,omitzero" format:"date" json:"-"`
	// URL substring or domain filter.
	URL param.Opt[string] `query:"url,omitzero" json:"-"`
	// Only return tweets from verified authors.
	VerifiedOnly param.Opt[bool] `query:"verifiedOnly,omitzero" json:"-"`
	// Set the radius for the near filter.
	Within param.Opt[string] `query:"within,omitzero" json:"-"`
	// Match Tweets inside a recent time window.
	WithinTime param.Opt[string] `query:"withinTime,omitzero" json:"-"`
	// Filter media. Aliases: has_video, has_media.
	//
	// Any of "images", "videos", "gifs", "media", "links", "none".
	MediaType XUserGetLikesParamsMediaType `query:"mediaType,omitzero" json:"-"`
	// Only when the caller requests a quote mode.
	//
	// Any of "include", "exclude", "only".
	Quotes XUserGetLikesParamsQuotes `query:"quotes,omitzero" json:"-"`
	// Only when the caller requests a reply mode.
	//
	// Any of "include", "exclude", "only".
	Replies XUserGetLikesParamsReplies `query:"replies,omitzero" json:"-"`
	// Only when the caller requests a repost mode.
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

// Filter media. Aliases: has_video, has_media.
type XUserGetLikesParamsMediaType string

const (
	XUserGetLikesParamsMediaTypeImages XUserGetLikesParamsMediaType = "images"
	XUserGetLikesParamsMediaTypeVideos XUserGetLikesParamsMediaType = "videos"
	XUserGetLikesParamsMediaTypeGifs   XUserGetLikesParamsMediaType = "gifs"
	XUserGetLikesParamsMediaTypeMedia  XUserGetLikesParamsMediaType = "media"
	XUserGetLikesParamsMediaTypeLinks  XUserGetLikesParamsMediaType = "links"
	XUserGetLikesParamsMediaTypeNone   XUserGetLikesParamsMediaType = "none"
)

// Only when the caller requests a quote mode.
type XUserGetLikesParamsQuotes string

const (
	XUserGetLikesParamsQuotesInclude XUserGetLikesParamsQuotes = "include"
	XUserGetLikesParamsQuotesExclude XUserGetLikesParamsQuotes = "exclude"
	XUserGetLikesParamsQuotesOnly    XUserGetLikesParamsQuotes = "only"
)

// Only when the caller requests a reply mode.
type XUserGetLikesParamsReplies string

const (
	XUserGetLikesParamsRepliesInclude XUserGetLikesParamsReplies = "include"
	XUserGetLikesParamsRepliesExclude XUserGetLikesParamsReplies = "exclude"
	XUserGetLikesParamsRepliesOnly    XUserGetLikesParamsReplies = "only"
)

// Only when the caller requests a repost mode.
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
	// Only return tweets from Blue-verified authors.
	BlueVerifiedOnly param.Opt[bool] `query:"blueVerifiedOnly,omitzero" json:"-"`
	// Match the Tweet card name.
	CardName param.Opt[string] `query:"cardName,omitzero" json:"-"`
	// Cashtags separated by spaces, commas, or lines.
	Cashtags param.Opt[string] `query:"cashtags,omitzero" json:"-"`
	// Conversation ID filter.
	ConversationID param.Opt[string] `query:"conversationId,omitzero" json:"-"`
	// Pagination cursor for media tweets
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Match this literal phrase, including any hyphens.
	ExactPhrase param.Opt[string] `query:"exactPhrase,omitzero" json:"-"`
	// Exclude a source application.
	ExcludeSource param.Opt[string] `query:"excludeSource,omitzero" json:"-"`
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
	// Match latitude, longitude, and radius.
	Geocode param.Opt[string] `query:"geocode,omitzero" json:"-"`
	// Hashtags separated by spaces, commas, or lines.
	Hashtags param.Opt[string] `query:"hashtags,omitzero" json:"-"`
	// Only replies to this tweet ID.
	InReplyToTweetID param.Opt[string] `query:"inReplyToTweetId,omitzero" json:"-"`
	// Filter by language. Alias `lang` is accepted.
	Language param.Opt[string] `query:"language,omitzero" json:"-"`
	// Maximum likes threshold. maxLikes is also accepted.
	MaxFaves param.Opt[int64] `query:"maxFaves,omitzero" json:"-"`
	// Return Tweets older than this Tweet ID.
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
	// Deprecated: Use MinLikes. MinFaves sends the minFaves query key, which the
	// API reads the same way as minLikes.
	MinFaves param.Opt[int64] `query:"minFaves,omitzero" json:"-"`
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
	// Match a place name.
	Near param.Opt[string] `query:"near,omitzero" json:"-"`
	// Only return news results.
	News param.Opt[bool] `query:"news,omitzero" json:"-"`
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
	// Enable the safe-search filter.
	Safe param.Opt[bool] `query:"safe,omitzero" json:"-"`
	// Start date in YYYY-MM-DD format.
	SinceDate param.Opt[time.Time] `query:"sinceDate,omitzero" format:"date" json:"-"`
	// Return Tweets newer than this Tweet ID.
	SinceID param.Opt[string] `query:"sinceId,omitzero" json:"-"`
	// Match the source application.
	Source param.Opt[string] `query:"source,omitzero" json:"-"`
	// Filter replies sent to a username.
	ToUser param.Opt[string] `query:"toUser,omitzero" json:"-"`
	// End date in YYYY-MM-DD format.
	UntilDate param.Opt[time.Time] `query:"untilDate,omitzero" format:"date" json:"-"`
	// URL substring or domain filter.
	URL param.Opt[string] `query:"url,omitzero" json:"-"`
	// Only return tweets from verified authors.
	VerifiedOnly param.Opt[bool] `query:"verifiedOnly,omitzero" json:"-"`
	// Set the radius for the near filter.
	Within param.Opt[string] `query:"within,omitzero" json:"-"`
	// Match Tweets inside a recent time window.
	WithinTime param.Opt[string] `query:"withinTime,omitzero" json:"-"`
	// Filter media. Aliases: has_video, has_media.
	//
	// Any of "images", "videos", "gifs", "media", "links", "none".
	MediaType XUserGetMediaParamsMediaType `query:"mediaType,omitzero" json:"-"`
	// Only when the caller requests a quote mode.
	//
	// Any of "include", "exclude", "only".
	Quotes XUserGetMediaParamsQuotes `query:"quotes,omitzero" json:"-"`
	// Only when the caller requests a reply mode.
	//
	// Any of "include", "exclude", "only".
	Replies XUserGetMediaParamsReplies `query:"replies,omitzero" json:"-"`
	// Only when the caller requests a repost mode.
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

// Filter media. Aliases: has_video, has_media.
type XUserGetMediaParamsMediaType string

const (
	XUserGetMediaParamsMediaTypeImages XUserGetMediaParamsMediaType = "images"
	XUserGetMediaParamsMediaTypeVideos XUserGetMediaParamsMediaType = "videos"
	XUserGetMediaParamsMediaTypeGifs   XUserGetMediaParamsMediaType = "gifs"
	XUserGetMediaParamsMediaTypeMedia  XUserGetMediaParamsMediaType = "media"
	XUserGetMediaParamsMediaTypeLinks  XUserGetMediaParamsMediaType = "links"
	XUserGetMediaParamsMediaTypeNone   XUserGetMediaParamsMediaType = "none"
)

// Only when the caller requests a quote mode.
type XUserGetMediaParamsQuotes string

const (
	XUserGetMediaParamsQuotesInclude XUserGetMediaParamsQuotes = "include"
	XUserGetMediaParamsQuotesExclude XUserGetMediaParamsQuotes = "exclude"
	XUserGetMediaParamsQuotesOnly    XUserGetMediaParamsQuotes = "only"
)

// Only when the caller requests a reply mode.
type XUserGetMediaParamsReplies string

const (
	XUserGetMediaParamsRepliesInclude XUserGetMediaParamsReplies = "include"
	XUserGetMediaParamsRepliesExclude XUserGetMediaParamsReplies = "exclude"
	XUserGetMediaParamsRepliesOnly    XUserGetMediaParamsReplies = "only"
)

// Only when the caller requests a repost mode.
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
	// Only return tweets from Blue-verified authors.
	BlueVerifiedOnly param.Opt[bool] `query:"blueVerifiedOnly,omitzero" json:"-"`
	// Match the Tweet card name.
	CardName param.Opt[string] `query:"cardName,omitzero" json:"-"`
	// Cashtags separated by spaces, commas, or lines.
	Cashtags param.Opt[string] `query:"cashtags,omitzero" json:"-"`
	// Conversation ID filter.
	ConversationID param.Opt[string] `query:"conversationId,omitzero" json:"-"`
	// Pagination cursor for mentions
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Match this literal phrase, including any hyphens.
	ExactPhrase param.Opt[string] `query:"exactPhrase,omitzero" json:"-"`
	// Exclude a source application.
	ExcludeSource param.Opt[string] `query:"excludeSource,omitzero" json:"-"`
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
	// Match latitude, longitude, and radius.
	Geocode param.Opt[string] `query:"geocode,omitzero" json:"-"`
	// Hashtags separated by spaces, commas, or lines.
	Hashtags param.Opt[string] `query:"hashtags,omitzero" json:"-"`
	// Only replies to this tweet ID.
	InReplyToTweetID param.Opt[string] `query:"inReplyToTweetId,omitzero" json:"-"`
	// Filter by language. Alias `lang` is accepted.
	Language param.Opt[string] `query:"language,omitzero" json:"-"`
	// Maximum likes threshold. maxLikes is also accepted.
	MaxFaves param.Opt[int64] `query:"maxFaves,omitzero" json:"-"`
	// Return Tweets older than this Tweet ID.
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
	// Deprecated: Use MinLikes. MinFaves sends the minFaves query key, which the
	// API reads the same way as minLikes.
	MinFaves param.Opt[int64] `query:"minFaves,omitzero" json:"-"`
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
	// Match a place name.
	Near param.Opt[string] `query:"near,omitzero" json:"-"`
	// Only return news results.
	News param.Opt[bool] `query:"news,omitzero" json:"-"`
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
	// Enable the safe-search filter.
	Safe param.Opt[bool] `query:"safe,omitzero" json:"-"`
	// Start date in YYYY-MM-DD format.
	SinceDate param.Opt[time.Time] `query:"sinceDate,omitzero" format:"date" json:"-"`
	// Return Tweets newer than this Tweet ID.
	SinceID param.Opt[string] `query:"sinceId,omitzero" json:"-"`
	// Inclusive ISO bound for Tweet creation time.
	SinceTime param.Opt[string] `query:"sinceTime,omitzero" json:"-"`
	// Match the source application.
	Source param.Opt[string] `query:"source,omitzero" json:"-"`
	// Filter replies sent to a username.
	ToUser param.Opt[string] `query:"toUser,omitzero" json:"-"`
	// End date in YYYY-MM-DD format.
	UntilDate param.Opt[time.Time] `query:"untilDate,omitzero" format:"date" json:"-"`
	// Exclusive ISO bound for Tweet creation time.
	UntilTime param.Opt[string] `query:"untilTime,omitzero" json:"-"`
	// URL substring or domain filter.
	URL param.Opt[string] `query:"url,omitzero" json:"-"`
	// Only return tweets from verified authors.
	VerifiedOnly param.Opt[bool] `query:"verifiedOnly,omitzero" json:"-"`
	// Set the radius for the near filter.
	Within param.Opt[string] `query:"within,omitzero" json:"-"`
	// Match Tweets inside a recent time window.
	WithinTime param.Opt[string] `query:"withinTime,omitzero" json:"-"`
	// Filter media. Aliases: has_video, has_media.
	//
	// Any of "images", "videos", "gifs", "media", "links", "none".
	MediaType XUserGetMentionsParamsMediaType `query:"mediaType,omitzero" json:"-"`
	// Only when the caller requests a quote mode.
	//
	// Any of "include", "exclude", "only".
	Quotes XUserGetMentionsParamsQuotes `query:"quotes,omitzero" json:"-"`
	// Only when the caller requests a reply mode.
	//
	// Any of "include", "exclude", "only".
	Replies XUserGetMentionsParamsReplies `query:"replies,omitzero" json:"-"`
	// Only when the caller requests a repost mode.
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

// Filter media. Aliases: has_video, has_media.
type XUserGetMentionsParamsMediaType string

const (
	XUserGetMentionsParamsMediaTypeImages XUserGetMentionsParamsMediaType = "images"
	XUserGetMentionsParamsMediaTypeVideos XUserGetMentionsParamsMediaType = "videos"
	XUserGetMentionsParamsMediaTypeGifs   XUserGetMentionsParamsMediaType = "gifs"
	XUserGetMentionsParamsMediaTypeMedia  XUserGetMentionsParamsMediaType = "media"
	XUserGetMentionsParamsMediaTypeLinks  XUserGetMentionsParamsMediaType = "links"
	XUserGetMentionsParamsMediaTypeNone   XUserGetMentionsParamsMediaType = "none"
)

// Only when the caller requests a quote mode.
type XUserGetMentionsParamsQuotes string

const (
	XUserGetMentionsParamsQuotesInclude XUserGetMentionsParamsQuotes = "include"
	XUserGetMentionsParamsQuotesExclude XUserGetMentionsParamsQuotes = "exclude"
	XUserGetMentionsParamsQuotesOnly    XUserGetMentionsParamsQuotes = "only"
)

// Only when the caller requests a reply mode.
type XUserGetMentionsParamsReplies string

const (
	XUserGetMentionsParamsRepliesInclude XUserGetMentionsParamsReplies = "include"
	XUserGetMentionsParamsRepliesExclude XUserGetMentionsParamsReplies = "exclude"
	XUserGetMentionsParamsRepliesOnly    XUserGetMentionsParamsReplies = "only"
)

// Only when the caller requests a repost mode.
type XUserGetMentionsParamsRetweets string

const (
	XUserGetMentionsParamsRetweetsInclude XUserGetMentionsParamsRetweets = "include"
	XUserGetMentionsParamsRetweetsExclude XUserGetMentionsParamsRetweets = "exclude"
	XUserGetMentionsParamsRetweetsOnly    XUserGetMentionsParamsRetweets = "only"
)

type XUserGetRepliesParams struct {
	// Raw advanced search query appended as-is.
	AdvancedQuery param.Opt[string] `query:"advancedQuery,omitzero" json:"-"`
	// Words or quoted phrases where any one can match. Separate with spaces, commas,
	// or lines.
	AnyWords param.Opt[string] `query:"anyWords,omitzero" json:"-"`
	// Only return tweets from Blue-verified authors.
	BlueVerifiedOnly param.Opt[bool] `query:"blueVerifiedOnly,omitzero" json:"-"`
	// Geo bounding box, such as -74.1 40.6 -73.9 40.8.
	BoundingBox param.Opt[string] `query:"boundingBox,omitzero" json:"-"`
	// Match the Tweet card name.
	CardName param.Opt[string] `query:"cardName,omitzero" json:"-"`
	// Cashtags separated by spaces, commas, or lines.
	Cashtags param.Opt[string] `query:"cashtags,omitzero" json:"-"`
	// Conversation ID filter.
	ConversationID param.Opt[string] `query:"conversationId,omitzero" json:"-"`
	// Cursor from the previous response. Xquik cursors resume automatic coverage.
	// Existing unprefixed cursors keep legacy standard behavior.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Match this literal phrase, including any hyphens.
	ExactPhrase param.Opt[string] `query:"exactPhrase,omitzero" json:"-"`
	// Exclude a source application.
	ExcludeSource param.Opt[string] `query:"excludeSource,omitzero" json:"-"`
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
	// Match latitude, longitude, and radius.
	Geocode param.Opt[string] `query:"geocode,omitzero" json:"-"`
	// Hashtags separated by spaces, commas, or lines.
	Hashtags param.Opt[string] `query:"hashtags,omitzero" json:"-"`
	// Include each reply's parent tweet.
	IncludeParentTweet param.Opt[bool] `query:"includeParentTweet,omitzero" json:"-"`
	// Only replies to this tweet ID.
	InReplyToTweetID param.Opt[string] `query:"inReplyToTweetId,omitzero" json:"-"`
	// Words the Tweets must match, in X search syntax.
	Keywords param.Opt[string] `query:"keywords,omitzero" json:"-"`
	// Filter by language. Alias `lang` is accepted.
	Language param.Opt[string] `query:"language,omitzero" json:"-"`
	// Search within a list ID.
	ListID param.Opt[string] `query:"listId,omitzero" json:"-"`
	// Maximum likes threshold. maxLikes is also accepted.
	MaxFaves param.Opt[int64] `query:"maxFaves,omitzero" json:"-"`
	// Return Tweets older than this Tweet ID.
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
	// Deprecated: Use MinLikes. MinFaves sends the minFaves query key, which the
	// API reads the same way as minLikes.
	MinFaves param.Opt[int64] `query:"minFaves,omitzero" json:"-"`
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
	// Match a place name.
	Near param.Opt[string] `query:"near,omitzero" json:"-"`
	// Only return news results.
	News param.Opt[bool] `query:"news,omitzero" json:"-"`
	// Fields to drop from each row, comma-separated. It takes the same names as
	// fields. With fields, it drops from the kept fields.
	Omit param.Opt[string] `query:"omit,omitzero" json:"-"`
	// Automatic pages accept 1-300 Tweets. Standard pages keep 1-100. Default 20.
	// Follow next_cursor while the response reports more pages. Deprecated aliases
	// remain accepted.
	PageSize param.Opt[int64] `query:"pageSize,omitzero" json:"-"`
	// Search within a place ID.
	Place param.Opt[string] `query:"place,omitzero" json:"-"`
	// Search within a country code.
	PlaceCountry param.Opt[string] `query:"placeCountry,omitzero" json:"-"`
	// Geo point radius, such as -73.99 40.73 25mi.
	PointRadius param.Opt[string] `query:"pointRadius,omitzero" json:"-"`
	// Only quotes of this tweet ID.
	QuotesOfTweetID param.Opt[string] `query:"quotesOfTweetId,omitzero" json:"-"`
	// Only retweets of this tweet ID.
	RetweetsOfTweetID param.Opt[string] `query:"retweetsOfTweetId,omitzero" json:"-"`
	// Enable the safe-search filter.
	Safe param.Opt[bool] `query:"safe,omitzero" json:"-"`
	// Start date in YYYY-MM-DD format.
	SinceDate param.Opt[time.Time] `query:"sinceDate,omitzero" format:"date" json:"-"`
	// Return Tweets newer than this Tweet ID.
	SinceID param.Opt[string] `query:"sinceId,omitzero" json:"-"`
	// Inclusive ISO bound for Tweet creation time, such as 2026-09-25T19:20:41Z. Unix
	// seconds also work. A time without an offset is UTC.
	SinceTime param.Opt[string] `query:"sinceTime,omitzero" json:"-"`
	// Match the source application.
	Source param.Opt[string] `query:"source,omitzero" json:"-"`
	// Filter replies sent to a username.
	ToUser param.Opt[string] `query:"toUser,omitzero" json:"-"`
	// End date in YYYY-MM-DD format.
	UntilDate param.Opt[time.Time] `query:"untilDate,omitzero" format:"date" json:"-"`
	// Exclusive ISO bound for Tweet creation time. Unix seconds also work. A time
	// without an offset is UTC.
	UntilTime param.Opt[string] `query:"untilTime,omitzero" json:"-"`
	// URL substring or domain filter.
	URL param.Opt[string] `query:"url,omitzero" json:"-"`
	// Only return tweets from verified authors.
	VerifiedOnly param.Opt[bool] `query:"verifiedOnly,omitzero" json:"-"`
	// Set the radius for the near filter.
	Within param.Opt[string] `query:"within,omitzero" json:"-"`
	// Match Tweets inside a recent time window.
	WithinTime param.Opt[string] `query:"withinTime,omitzero" json:"-"`
	// Filter media. Aliases: has_video, has_media.
	//
	// Any of "images", "videos", "gifs", "media", "links", "none".
	MediaType XUserGetRepliesParamsMediaType `query:"mediaType,omitzero" json:"-"`
	// Optional legacy pagination override. Search-only filters such as keywords apply
	// only without it.
	//
	// Any of "standard".
	Mode XUserGetRepliesParamsMode `query:"mode,omitzero" json:"-"`
	// Only when the caller requests a quote mode.
	//
	// Any of "include", "exclude", "only".
	Quotes XUserGetRepliesParamsQuotes `query:"quotes,omitzero" json:"-"`
	// Only when the caller requests a reply mode.
	//
	// Any of "include", "exclude", "only".
	Replies XUserGetRepliesParamsReplies `query:"replies,omitzero" json:"-"`
	// Only when the caller requests a repost mode.
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

// Filter media. Aliases: has_video, has_media.
type XUserGetRepliesParamsMediaType string

const (
	XUserGetRepliesParamsMediaTypeImages XUserGetRepliesParamsMediaType = "images"
	XUserGetRepliesParamsMediaTypeVideos XUserGetRepliesParamsMediaType = "videos"
	XUserGetRepliesParamsMediaTypeGifs   XUserGetRepliesParamsMediaType = "gifs"
	XUserGetRepliesParamsMediaTypeMedia  XUserGetRepliesParamsMediaType = "media"
	XUserGetRepliesParamsMediaTypeLinks  XUserGetRepliesParamsMediaType = "links"
	XUserGetRepliesParamsMediaTypeNone   XUserGetRepliesParamsMediaType = "none"
)

// Only when the caller requests a quote mode.
type XUserGetRepliesParamsQuotes string

const (
	XUserGetRepliesParamsQuotesInclude XUserGetRepliesParamsQuotes = "include"
	XUserGetRepliesParamsQuotesExclude XUserGetRepliesParamsQuotes = "exclude"
	XUserGetRepliesParamsQuotesOnly    XUserGetRepliesParamsQuotes = "only"
)

// Only when the caller requests a reply mode.
type XUserGetRepliesParamsReplies string

const (
	XUserGetRepliesParamsRepliesInclude XUserGetRepliesParamsReplies = "include"
	XUserGetRepliesParamsRepliesExclude XUserGetRepliesParamsReplies = "exclude"
	XUserGetRepliesParamsRepliesOnly    XUserGetRepliesParamsReplies = "only"
)

// Only when the caller requests a repost mode.
type XUserGetRepliesParamsRetweets string

const (
	XUserGetRepliesParamsRetweetsInclude XUserGetRepliesParamsRetweets = "include"
	XUserGetRepliesParamsRetweetsExclude XUserGetRepliesParamsRetweets = "exclude"
	XUserGetRepliesParamsRetweetsOnly    XUserGetRepliesParamsRetweets = "only"
)

type XUserGetSearchParams struct {
	// User search query
	Q string `query:"q" api:"required" json:"-"`
	// Match any comma-separated or line-separated bio term, ignoring case.
	BioContains param.Opt[string] `query:"bioContains,omitzero" json:"-"`
	// Only return profiles X marks as blue verified. X marks many business &
	// government accounts blue too. Use verifiedType for 1 badge type.
	BlueVerifiedOnly param.Opt[bool] `query:"blueVerifiedOnly,omitzero" json:"-"`
	// Only return profiles that publish an email outside free webmail, such as Gmail.
	BusinessEmailOnly param.Opt[bool] `query:"businessEmailOnly,omitzero" json:"-"`
	// Pagination cursor for user search
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Only return profiles with an email at one of these domains. Separate domains
	// with commas. Subdomains match.
	EmailDomains param.Opt[string] `query:"emailDomains,omitzero" json:"-"`
	// Add each row's full profile, as GET /x/users/{id} returns it. Pages take several
	// seconds longer.
	EnrichProfiles param.Opt[bool] `query:"enrichProfiles,omitzero" json:"-"`
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
	// Maximum page items (1-100, default 20). Source, filters, or credits can reduce
	// results. Follow next_cursor while the response reports more pages. Deprecated
	// limit and count aliases remain accepted.
	PageSize param.Opt[int64] `query:"pageSize,omitzero" json:"-"`
	// Match a username substring, ignoring case.
	UsernameContains param.Opt[string] `query:"usernameContains,omitzero" json:"-"`
	// Only return verified profiles.
	VerifiedOnly param.Opt[bool] `query:"verifiedOnly,omitzero" json:"-"`
	// Match the verification type exactly, ignoring case.
	VerifiedType param.Opt[string] `query:"verifiedType,omitzero" json:"-"`
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
	// Raw advanced search query appended as-is.
	AdvancedQuery param.Opt[string] `query:"advancedQuery,omitzero" json:"-"`
	// Words or quoted phrases where any one can match. Separate with spaces, commas,
	// or lines.
	AnyWords param.Opt[string] `query:"anyWords,omitzero" json:"-"`
	// Only return tweets from Blue-verified authors.
	BlueVerifiedOnly param.Opt[bool] `query:"blueVerifiedOnly,omitzero" json:"-"`
	// Geo bounding box, such as -74.1 40.6 -73.9 40.8.
	BoundingBox param.Opt[string] `query:"boundingBox,omitzero" json:"-"`
	// Match the Tweet card name.
	CardName param.Opt[string] `query:"cardName,omitzero" json:"-"`
	// Cashtags separated by spaces, commas, or lines.
	Cashtags param.Opt[string] `query:"cashtags,omitzero" json:"-"`
	// Conversation ID filter.
	ConversationID param.Opt[string] `query:"conversationId,omitzero" json:"-"`
	// Cursor from the previous response. Xquik cursors resume automatic coverage.
	// Existing unprefixed cursors keep legacy standard behavior.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Match this literal phrase, including any hyphens.
	ExactPhrase param.Opt[string] `query:"exactPhrase,omitzero" json:"-"`
	// Exclude a source application.
	ExcludeSource param.Opt[string] `query:"excludeSource,omitzero" json:"-"`
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
	// Match latitude, longitude, and radius.
	Geocode param.Opt[string] `query:"geocode,omitzero" json:"-"`
	// Hashtags separated by spaces, commas, or lines.
	Hashtags param.Opt[string] `query:"hashtags,omitzero" json:"-"`
	// Include parent tweet for replies
	IncludeParentTweet param.Opt[bool] `query:"includeParentTweet,omitzero" json:"-"`
	// Include reply tweets unless replies specifies another mode.
	IncludeReplies param.Opt[bool] `query:"includeReplies,omitzero" json:"-"`
	// Only replies to this tweet ID.
	InReplyToTweetID param.Opt[string] `query:"inReplyToTweetId,omitzero" json:"-"`
	// Words the Tweets must match, in X search syntax.
	Keywords param.Opt[string] `query:"keywords,omitzero" json:"-"`
	// Filter by language. Alias `lang` is accepted.
	Language param.Opt[string] `query:"language,omitzero" json:"-"`
	// Search within a list ID.
	ListID param.Opt[string] `query:"listId,omitzero" json:"-"`
	// Maximum likes threshold. maxLikes is also accepted.
	MaxFaves param.Opt[int64] `query:"maxFaves,omitzero" json:"-"`
	// Return Tweets older than this Tweet ID.
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
	// Deprecated: Use MinLikes. MinFaves sends the minFaves query key, which the
	// API reads the same way as minLikes.
	MinFaves param.Opt[int64] `query:"minFaves,omitzero" json:"-"`
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
	// Match a place name.
	Near param.Opt[string] `query:"near,omitzero" json:"-"`
	// Only return news results.
	News param.Opt[bool] `query:"news,omitzero" json:"-"`
	// Fields to drop from each row, comma-separated. It takes the same names as
	// fields. With fields, it drops from the kept fields.
	Omit param.Opt[string] `query:"omit,omitzero" json:"-"`
	// Automatic pages accept 1-300 Tweets. Standard pages keep 1-100. Default 20.
	// Follow next_cursor while the response reports more pages. Deprecated aliases
	// remain accepted.
	PageSize param.Opt[int64] `query:"pageSize,omitzero" json:"-"`
	// Search within a place ID.
	Place param.Opt[string] `query:"place,omitzero" json:"-"`
	// Search within a country code.
	PlaceCountry param.Opt[string] `query:"placeCountry,omitzero" json:"-"`
	// Geo point radius, such as -73.99 40.73 25mi.
	PointRadius param.Opt[string] `query:"pointRadius,omitzero" json:"-"`
	// Only quotes of this tweet ID.
	QuotesOfTweetID param.Opt[string] `query:"quotesOfTweetId,omitzero" json:"-"`
	// Only retweets of this tweet ID.
	RetweetsOfTweetID param.Opt[string] `query:"retweetsOfTweetId,omitzero" json:"-"`
	// Enable the safe-search filter.
	Safe param.Opt[bool] `query:"safe,omitzero" json:"-"`
	// Start date in YYYY-MM-DD format.
	SinceDate param.Opt[time.Time] `query:"sinceDate,omitzero" format:"date" json:"-"`
	// Return Tweets newer than this Tweet ID.
	SinceID param.Opt[string] `query:"sinceId,omitzero" json:"-"`
	// Inclusive ISO bound for Tweet creation time, such as 2026-09-25T19:20:41Z. Unix
	// seconds also work. A time without an offset is UTC.
	SinceTime param.Opt[string] `query:"sinceTime,omitzero" json:"-"`
	// Match the source application.
	Source param.Opt[string] `query:"source,omitzero" json:"-"`
	// Filter replies sent to a username.
	ToUser param.Opt[string] `query:"toUser,omitzero" json:"-"`
	// End date in YYYY-MM-DD format.
	UntilDate param.Opt[time.Time] `query:"untilDate,omitzero" format:"date" json:"-"`
	// Exclusive ISO bound for Tweet creation time. Unix seconds also work. A time
	// without an offset is UTC.
	UntilTime param.Opt[string] `query:"untilTime,omitzero" json:"-"`
	// URL substring or domain filter.
	URL param.Opt[string] `query:"url,omitzero" json:"-"`
	// Only return tweets from verified authors.
	VerifiedOnly param.Opt[bool] `query:"verifiedOnly,omitzero" json:"-"`
	// Set the radius for the near filter.
	Within param.Opt[string] `query:"within,omitzero" json:"-"`
	// Match Tweets inside a recent time window.
	WithinTime param.Opt[string] `query:"withinTime,omitzero" json:"-"`
	// Filter media. Aliases: has_video, has_media.
	//
	// Any of "images", "videos", "gifs", "media", "links", "none".
	MediaType XUserGetTweetsParamsMediaType `query:"mediaType,omitzero" json:"-"`
	// Optional legacy pagination override. Search-only filters such as keywords apply
	// only without it.
	//
	// Any of "standard".
	Mode XUserGetTweetsParamsMode `query:"mode,omitzero" json:"-"`
	// Only when the caller requests a quote mode.
	//
	// Any of "include", "exclude", "only".
	Quotes XUserGetTweetsParamsQuotes `query:"quotes,omitzero" json:"-"`
	// Only when the caller requests a reply mode.
	//
	// Any of "include", "exclude", "only".
	Replies XUserGetTweetsParamsReplies `query:"replies,omitzero" json:"-"`
	// Only when the caller requests a repost mode.
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

// Filter media. Aliases: has_video, has_media.
type XUserGetTweetsParamsMediaType string

const (
	XUserGetTweetsParamsMediaTypeImages XUserGetTweetsParamsMediaType = "images"
	XUserGetTweetsParamsMediaTypeVideos XUserGetTweetsParamsMediaType = "videos"
	XUserGetTweetsParamsMediaTypeGifs   XUserGetTweetsParamsMediaType = "gifs"
	XUserGetTweetsParamsMediaTypeMedia  XUserGetTweetsParamsMediaType = "media"
	XUserGetTweetsParamsMediaTypeLinks  XUserGetTweetsParamsMediaType = "links"
	XUserGetTweetsParamsMediaTypeNone   XUserGetTweetsParamsMediaType = "none"
)

// Only when the caller requests a quote mode.
type XUserGetTweetsParamsQuotes string

const (
	XUserGetTweetsParamsQuotesInclude XUserGetTweetsParamsQuotes = "include"
	XUserGetTweetsParamsQuotesExclude XUserGetTweetsParamsQuotes = "exclude"
	XUserGetTweetsParamsQuotesOnly    XUserGetTweetsParamsQuotes = "only"
)

// Only when the caller requests a reply mode.
type XUserGetTweetsParamsReplies string

const (
	XUserGetTweetsParamsRepliesInclude XUserGetTweetsParamsReplies = "include"
	XUserGetTweetsParamsRepliesExclude XUserGetTweetsParamsReplies = "exclude"
	XUserGetTweetsParamsRepliesOnly    XUserGetTweetsParamsReplies = "only"
)

// Only when the caller requests a repost mode.
type XUserGetTweetsParamsRetweets string

const (
	XUserGetTweetsParamsRetweetsInclude XUserGetTweetsParamsRetweets = "include"
	XUserGetTweetsParamsRetweetsExclude XUserGetTweetsParamsRetweets = "exclude"
	XUserGetTweetsParamsRetweetsOnly    XUserGetTweetsParamsRetweets = "only"
)

type XUserGetVerifiedFollowersParams struct {
	// Legacy verified-follower cursor alias. Prefer cursor.
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// Match any comma-separated or line-separated bio term, ignoring case.
	BioContains param.Opt[string] `query:"bioContains,omitzero" json:"-"`
	// Only return profiles X marks as blue verified. X marks many business &
	// government accounts blue too. Use verifiedType for 1 badge type.
	BlueVerifiedOnly param.Opt[bool] `query:"blueVerifiedOnly,omitzero" json:"-"`
	// Only return profiles that publish an email outside free webmail, such as Gmail.
	BusinessEmailOnly param.Opt[bool] `query:"businessEmailOnly,omitzero" json:"-"`
	// Cursor from the previous response. Xquik cursors resume automatic coverage.
	// Existing unprefixed cursors keep legacy standard behavior.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Only return profiles with an email at one of these domains. Separate domains
	// with commas. Subdomains match.
	EmailDomains param.Opt[string] `query:"emailDomains,omitzero" json:"-"`
	// Add each row's full profile, as GET /x/users/{id} returns it. Pages take several
	// seconds longer.
	EnrichProfiles param.Opt[bool] `query:"enrichProfiles,omitzero" json:"-"`
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
	// Legacy verified-follower page-size alias. Coverage accepts 1-10000. Prefer
	// pageSize.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Match a location substring, ignoring case.
	LocationContains param.Opt[string] `query:"locationContains,omitzero" json:"-"`
	// Maximum follower count. Missing counts pass this maximum.
	MaxFollowers param.Opt[int64] `query:"maxFollowers,omitzero" json:"-"`
	// Profiles may follow at most this many accounts.
	MaxFollowing param.Opt[int64] `query:"maxFollowing,omitzero" json:"-"`
	// Most profiles 1 request reads, kept or dropped by the filters. The request then
	// answers the profiles it kept, the count it dropped & the cursor to read on. Each
	// request of a cursor may send its own value. Without it, a filtered read scans at
	// most 100,000 profiles of the list across all its pages, & the page that reaches
	// them ends the list with `scanLimit`. An unfiltered read has no scan limit.
	// Automatic mode only.
	MaxScan param.Opt[int64] `query:"maxScan,omitzero" json:"-"`
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
	// Maximum user profiles: automatic 300; standard 200. Sources return fewer
	// profiles. Follow next_cursor while the response reports more pages.
	PageSize param.Opt[int64] `query:"pageSize,omitzero" json:"-"`
	// Match a username substring, ignoring case.
	UsernameContains param.Opt[string] `query:"usernameContains,omitzero" json:"-"`
	// Only return verified profiles.
	VerifiedOnly param.Opt[bool] `query:"verifiedOnly,omitzero" json:"-"`
	// Match the verification type exactly, ignoring case.
	VerifiedType param.Opt[string] `query:"verifiedType,omitzero" json:"-"`
	// Omit mode for resumable maximum coverage. Standard keeps legacy pagination.
	// Coverage returns diagnostics once and rejects cursors.
	//
	// Any of "standard", "coverage".
	Mode XUserGetVerifiedFollowersParamsMode `query:"mode,omitzero" json:"-"`
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

// Omit mode for resumable maximum coverage. Standard keeps legacy pagination.
// Coverage returns diagnostics once and rejects cursors.
type XUserGetVerifiedFollowersParamsMode string

const (
	XUserGetVerifiedFollowersParamsModeStandard XUserGetVerifiedFollowersParamsMode = "standard"
	XUserGetVerifiedFollowersParamsModeCoverage XUserGetVerifiedFollowersParamsMode = "coverage"
)

type XUserGetAffiliatesParams struct {
	// Match any comma-separated or line-separated bio term, ignoring case.
	BioContains param.Opt[string] `query:"bioContains,omitzero" json:"-"`
	// Only return profiles X marks as blue verified. X marks many business &
	// government accounts blue too. Use verifiedType for 1 badge type.
	BlueVerifiedOnly param.Opt[bool] `query:"blueVerifiedOnly,omitzero" json:"-"`
	// Only return profiles that publish an email outside free webmail, such as Gmail.
	BusinessEmailOnly param.Opt[bool] `query:"businessEmailOnly,omitzero" json:"-"`
	// Pagination cursor for collection results. Also read as next_cursor, next_token,
	// pagination_token or continuation_token.
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
	// Profiles to request from this page (1-200, default 200). Source, filters or
	// credits can return fewer. Follow next_cursor for more pages. Its aliases are
	// limit, count, max_results, maxItems, max_items & per_page.
	PageSize param.Opt[int64] `query:"pageSize,omitzero" json:"-"`
	// Match a username substring, ignoring case.
	UsernameContains param.Opt[string] `query:"usernameContains,omitzero" json:"-"`
	// Only return verified profiles.
	VerifiedOnly param.Opt[bool] `query:"verifiedOnly,omitzero" json:"-"`
	// Match the verification type exactly, ignoring case.
	VerifiedType param.Opt[string] `query:"verifiedType,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [XUserGetAffiliatesParams]'s query parameters as `url.Values`.
func (r XUserGetAffiliatesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type XUserGetArticlesParams struct {
	// Words or quoted phrases where any one can match. Separate with spaces, commas,
	// or lines.
	AnyWords param.Opt[string] `query:"anyWords,omitzero" json:"-"`
	// Only return tweets from Blue-verified authors.
	BlueVerifiedOnly param.Opt[bool] `query:"blueVerifiedOnly,omitzero" json:"-"`
	// Cashtags separated by spaces, commas, or lines.
	Cashtags param.Opt[string] `query:"cashtags,omitzero" json:"-"`
	// Conversation ID filter.
	ConversationID param.Opt[string] `query:"conversationId,omitzero" json:"-"`
	// Pagination cursor for collection results. Also read as next_cursor, next_token,
	// pagination_token or continuation_token.
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
	// Filter by language. Its alias is `lang`.
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
	// Minimum likes. Its aliases are minFaves, min_likes & min_faves.
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
	// Items per page (1-100, default 20). Source, filters or credits can return fewer.
	// Follow next_cursor for more pages. Its aliases are limit, count, max_results,
	// maxItems, max_items & per_page.
	PageSize param.Opt[int64] `query:"pageSize,omitzero" json:"-"`
	// Only quotes of this tweet ID.
	QuotesOfTweetID param.Opt[string] `query:"quotesOfTweetId,omitzero" json:"-"`
	// Only retweets of this tweet ID.
	RetweetsOfTweetID param.Opt[string] `query:"retweetsOfTweetId,omitzero" json:"-"`
	// Start date in YYYY-MM-DD format, as a UTC day.
	SinceDate param.Opt[time.Time] `query:"sinceDate,omitzero" format:"date" json:"-"`
	// Return Tweets whose IDs exceed this ID.
	SinceID param.Opt[string] `query:"sinceId,omitzero" json:"-"`
	// Filter replies sent to a username.
	ToUser param.Opt[string] `query:"toUser,omitzero" json:"-"`
	// End date in YYYY-MM-DD format, as a UTC day, inclusive, so its own posts count.
	UntilDate param.Opt[time.Time] `query:"untilDate,omitzero" format:"date" json:"-"`
	// URL substring or domain filter.
	URL param.Opt[string] `query:"url,omitzero" json:"-"`
	// Only return tweets from verified authors.
	VerifiedOnly param.Opt[bool] `query:"verifiedOnly,omitzero" json:"-"`
	// Keeps Tweets from a recent window, such as 90m or 7d. Use a whole number & s, m,
	// h or d.
	WithinTime param.Opt[string] `query:"withinTime,omitzero" json:"-"`
	// Filter media. Its aliases are has_video & has_media.
	//
	// Any of "images", "videos", "gifs", "media", "links", "none".
	MediaType XUserGetArticlesParamsMediaType `query:"mediaType,omitzero" json:"-"`
	// Only when the caller requests a quote mode.
	//
	// Any of "include", "exclude", "only".
	Quotes XUserGetArticlesParamsQuotes `query:"quotes,omitzero" json:"-"`
	// Reply mode. Exclude keeps own-thread replies.
	//
	// Any of "include", "exclude", "only".
	Replies XUserGetArticlesParamsReplies `query:"replies,omitzero" json:"-"`
	// Only when the caller requests a repost mode.
	//
	// Any of "include", "exclude", "only".
	Retweets XUserGetArticlesParamsRetweets `query:"retweets,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [XUserGetArticlesParams]'s query parameters as `url.Values`.
func (r XUserGetArticlesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter media. Its aliases are has_video & has_media.
type XUserGetArticlesParamsMediaType string

const (
	XUserGetArticlesParamsMediaTypeImages XUserGetArticlesParamsMediaType = "images"
	XUserGetArticlesParamsMediaTypeVideos XUserGetArticlesParamsMediaType = "videos"
	XUserGetArticlesParamsMediaTypeGifs   XUserGetArticlesParamsMediaType = "gifs"
	XUserGetArticlesParamsMediaTypeMedia  XUserGetArticlesParamsMediaType = "media"
	XUserGetArticlesParamsMediaTypeLinks  XUserGetArticlesParamsMediaType = "links"
	XUserGetArticlesParamsMediaTypeNone   XUserGetArticlesParamsMediaType = "none"
)

// Only when the caller requests a quote mode.
type XUserGetArticlesParamsQuotes string

const (
	XUserGetArticlesParamsQuotesInclude XUserGetArticlesParamsQuotes = "include"
	XUserGetArticlesParamsQuotesExclude XUserGetArticlesParamsQuotes = "exclude"
	XUserGetArticlesParamsQuotesOnly    XUserGetArticlesParamsQuotes = "only"
)

// Reply mode. Exclude keeps own-thread replies.
type XUserGetArticlesParamsReplies string

const (
	XUserGetArticlesParamsRepliesInclude XUserGetArticlesParamsReplies = "include"
	XUserGetArticlesParamsRepliesExclude XUserGetArticlesParamsReplies = "exclude"
	XUserGetArticlesParamsRepliesOnly    XUserGetArticlesParamsReplies = "only"
)

// Only when the caller requests a repost mode.
type XUserGetArticlesParamsRetweets string

const (
	XUserGetArticlesParamsRetweetsInclude XUserGetArticlesParamsRetweets = "include"
	XUserGetArticlesParamsRetweetsExclude XUserGetArticlesParamsRetweets = "exclude"
	XUserGetArticlesParamsRetweetsOnly    XUserGetArticlesParamsRetweets = "only"
)

type XUserGetBatchTweetsParams struct {
	// Fields to keep in each row, comma-separated. Use dots for nested fields, such as
	// author.username. A list keeps the field in each item. Case & underscores do not
	// matter. A row that lacks a field leaves it out. Cursors, counts & diagnostics
	// always stay. An unknown name answers 400 unknown_field with the valid names,
	// before any read. The price stays the same. SDK row types then hold only the
	// named fields.
	Fields param.Opt[string] `query:"fields,omitzero" json:"-"`
	// Comma-separated numeric user IDs (1-20 values). Xquik ignores duplicate IDs and
	// keeps the first-seen order.
	IDs param.Opt[string] `query:"ids,omitzero" json:"-"`
	// Fields to drop from each row, comma-separated. It takes the same names as
	// fields. With fields, it drops from the kept fields.
	Omit param.Opt[string] `query:"omit,omitzero" json:"-"`
	// Posts per user (1-20, default 20). A low credit balance can return fewer. Its
	// aliases are limit, count, max_results, maxItems, max_items & per_page.
	PageSize param.Opt[int64] `query:"pageSize,omitzero" json:"-"`
	// Comma-separated X usernames, with or without @, or profile links such as
	// x.com/nasa (1-20 values), in place of ids. A name in another case is the same
	// account. A number is a user ID.
	Usernames param.Opt[string] `query:"usernames,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [XUserGetBatchTweetsParams]'s query parameters as `url.Values`.
func (r XUserGetBatchTweetsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type XUserGetFollowerIDsParams struct {
	// Pagination cursor for collection results. Also read as next_cursor, next_token,
	// pagination_token or continuation_token.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// User IDs to request from this page (1-5000, default 5000). A low credit balance
	// can return fewer. Follow next_cursor for more pages. Its aliases are limit,
	// count, max_results, maxItems, max_items & per_page.
	PageSize param.Opt[int64] `query:"pageSize,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [XUserGetFollowerIDsParams]'s query parameters as `url.Values`.
func (r XUserGetFollowerIDsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type XUserGetFollowingIDsParams struct {
	// Pagination cursor for collection results. Also read as next_cursor, next_token,
	// pagination_token or continuation_token.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// User IDs to request from this page (1-5000, default 5000). A low credit balance
	// can return fewer. Follow next_cursor for more pages. Its aliases are limit,
	// count, max_results, maxItems, max_items & per_page.
	PageSize param.Opt[int64] `query:"pageSize,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [XUserGetFollowingIDsParams]'s query parameters as `url.Values`.
func (r XUserGetFollowingIDsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type XUserGetHighlightsParams struct {
	// Words or quoted phrases where any one can match. Separate with spaces, commas,
	// or lines.
	AnyWords param.Opt[string] `query:"anyWords,omitzero" json:"-"`
	// Only return tweets from Blue-verified authors.
	BlueVerifiedOnly param.Opt[bool] `query:"blueVerifiedOnly,omitzero" json:"-"`
	// Cashtags separated by spaces, commas, or lines.
	Cashtags param.Opt[string] `query:"cashtags,omitzero" json:"-"`
	// Conversation ID filter.
	ConversationID param.Opt[string] `query:"conversationId,omitzero" json:"-"`
	// Pagination cursor for collection results. Also read as next_cursor, next_token,
	// pagination_token or continuation_token.
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
	// Filter by language. Its alias is `lang`.
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
	// Minimum likes. Its aliases are minFaves, min_likes & min_faves.
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
	// Items per page (1-100, default 20). Source, filters or credits can return fewer.
	// Follow next_cursor for more pages. Its aliases are limit, count, max_results,
	// maxItems, max_items & per_page.
	PageSize param.Opt[int64] `query:"pageSize,omitzero" json:"-"`
	// Only quotes of this tweet ID.
	QuotesOfTweetID param.Opt[string] `query:"quotesOfTweetId,omitzero" json:"-"`
	// Only retweets of this tweet ID.
	RetweetsOfTweetID param.Opt[string] `query:"retweetsOfTweetId,omitzero" json:"-"`
	// Start date in YYYY-MM-DD format, as a UTC day.
	SinceDate param.Opt[time.Time] `query:"sinceDate,omitzero" format:"date" json:"-"`
	// Return Tweets whose IDs exceed this ID.
	SinceID param.Opt[string] `query:"sinceId,omitzero" json:"-"`
	// Filter replies sent to a username.
	ToUser param.Opt[string] `query:"toUser,omitzero" json:"-"`
	// End date in YYYY-MM-DD format, as a UTC day, inclusive, so its own posts count.
	UntilDate param.Opt[time.Time] `query:"untilDate,omitzero" format:"date" json:"-"`
	// URL substring or domain filter.
	URL param.Opt[string] `query:"url,omitzero" json:"-"`
	// Only return tweets from verified authors.
	VerifiedOnly param.Opt[bool] `query:"verifiedOnly,omitzero" json:"-"`
	// Keeps Tweets from a recent window, such as 90m or 7d. Use a whole number & s, m,
	// h or d.
	WithinTime param.Opt[string] `query:"withinTime,omitzero" json:"-"`
	// Filter media. Its aliases are has_video & has_media.
	//
	// Any of "images", "videos", "gifs", "media", "links", "none".
	MediaType XUserGetHighlightsParamsMediaType `query:"mediaType,omitzero" json:"-"`
	// Only when the caller requests a quote mode.
	//
	// Any of "include", "exclude", "only".
	Quotes XUserGetHighlightsParamsQuotes `query:"quotes,omitzero" json:"-"`
	// Reply mode. Exclude keeps own-thread replies.
	//
	// Any of "include", "exclude", "only".
	Replies XUserGetHighlightsParamsReplies `query:"replies,omitzero" json:"-"`
	// Only when the caller requests a repost mode.
	//
	// Any of "include", "exclude", "only".
	Retweets XUserGetHighlightsParamsRetweets `query:"retweets,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [XUserGetHighlightsParams]'s query parameters as `url.Values`.
func (r XUserGetHighlightsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter media. Its aliases are has_video & has_media.
type XUserGetHighlightsParamsMediaType string

const (
	XUserGetHighlightsParamsMediaTypeImages XUserGetHighlightsParamsMediaType = "images"
	XUserGetHighlightsParamsMediaTypeVideos XUserGetHighlightsParamsMediaType = "videos"
	XUserGetHighlightsParamsMediaTypeGifs   XUserGetHighlightsParamsMediaType = "gifs"
	XUserGetHighlightsParamsMediaTypeMedia  XUserGetHighlightsParamsMediaType = "media"
	XUserGetHighlightsParamsMediaTypeLinks  XUserGetHighlightsParamsMediaType = "links"
	XUserGetHighlightsParamsMediaTypeNone   XUserGetHighlightsParamsMediaType = "none"
)

// Only when the caller requests a quote mode.
type XUserGetHighlightsParamsQuotes string

const (
	XUserGetHighlightsParamsQuotesInclude XUserGetHighlightsParamsQuotes = "include"
	XUserGetHighlightsParamsQuotesExclude XUserGetHighlightsParamsQuotes = "exclude"
	XUserGetHighlightsParamsQuotesOnly    XUserGetHighlightsParamsQuotes = "only"
)

// Reply mode. Exclude keeps own-thread replies.
type XUserGetHighlightsParamsReplies string

const (
	XUserGetHighlightsParamsRepliesInclude XUserGetHighlightsParamsReplies = "include"
	XUserGetHighlightsParamsRepliesExclude XUserGetHighlightsParamsReplies = "exclude"
	XUserGetHighlightsParamsRepliesOnly    XUserGetHighlightsParamsReplies = "only"
)

// Only when the caller requests a repost mode.
type XUserGetHighlightsParamsRetweets string

const (
	XUserGetHighlightsParamsRetweetsInclude XUserGetHighlightsParamsRetweets = "include"
	XUserGetHighlightsParamsRetweetsExclude XUserGetHighlightsParamsRetweets = "exclude"
	XUserGetHighlightsParamsRetweetsOnly    XUserGetHighlightsParamsRetweets = "only"
)

type XUserGetInsightsParams struct {
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

// URLQuery serializes [XUserGetInsightsParams]'s query parameters as `url.Values`.
func (r XUserGetInsightsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type XUserGetListMembershipsParams struct {
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

// URLQuery serializes [XUserGetListMembershipsParams]'s query parameters as `url.Values`.
func (r XUserGetListMembershipsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type XUserGetListsParams struct {
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

// URLQuery serializes [XUserGetListsParams]'s query parameters as `url.Values`.
func (r XUserGetListsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type XUserGetLiveParams struct {
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

// URLQuery serializes [XUserGetLiveParams]'s query parameters as `url.Values`.
func (r XUserGetLiveParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type XUserGetSimilarParams struct {
	// Match any comma-separated or line-separated bio term, ignoring case.
	BioContains param.Opt[string] `query:"bioContains,omitzero" json:"-"`
	// Only return profiles X marks as blue verified. X marks many business &
	// government accounts blue too. Use verifiedType for 1 badge type.
	BlueVerifiedOnly param.Opt[bool] `query:"blueVerifiedOnly,omitzero" json:"-"`
	// Only return profiles that publish an email outside free webmail, such as Gmail.
	BusinessEmailOnly param.Opt[bool] `query:"businessEmailOnly,omitzero" json:"-"`
	// Pagination cursor for collection results. Also read as next_cursor, next_token,
	// pagination_token or continuation_token.
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
	// Profiles to request from this page (1-200, default 200). Source, filters or
	// credits can return fewer. Follow next_cursor for more pages. Its aliases are
	// limit, count, max_results, maxItems, max_items & per_page.
	PageSize param.Opt[int64] `query:"pageSize,omitzero" json:"-"`
	// Match a username substring, ignoring case.
	UsernameContains param.Opt[string] `query:"usernameContains,omitzero" json:"-"`
	// Only return verified profiles.
	VerifiedOnly param.Opt[bool] `query:"verifiedOnly,omitzero" json:"-"`
	// Match the verification type exactly, ignoring case.
	VerifiedType param.Opt[string] `query:"verifiedType,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [XUserGetSimilarParams]'s query parameters as `url.Values`.
func (r XUserGetSimilarParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type XUserGetSubscriptionsParams struct {
	// Match any comma-separated or line-separated bio term, ignoring case.
	BioContains param.Opt[string] `query:"bioContains,omitzero" json:"-"`
	// Only return profiles X marks as blue verified. X marks many business &
	// government accounts blue too. Use verifiedType for 1 badge type.
	BlueVerifiedOnly param.Opt[bool] `query:"blueVerifiedOnly,omitzero" json:"-"`
	// Only return profiles that publish an email outside free webmail, such as Gmail.
	BusinessEmailOnly param.Opt[bool] `query:"businessEmailOnly,omitzero" json:"-"`
	// Pagination cursor for collection results. Also read as next_cursor, next_token,
	// pagination_token or continuation_token.
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
	// Profiles to request from this page (1-200, default 200). Source, filters or
	// credits can return fewer. Follow next_cursor for more pages. Its aliases are
	// limit, count, max_results, maxItems, max_items & per_page.
	PageSize param.Opt[int64] `query:"pageSize,omitzero" json:"-"`
	// Match a username substring, ignoring case.
	UsernameContains param.Opt[string] `query:"usernameContains,omitzero" json:"-"`
	// Only return verified profiles.
	VerifiedOnly param.Opt[bool] `query:"verifiedOnly,omitzero" json:"-"`
	// Match the verification type exactly, ignoring case.
	VerifiedType param.Opt[string] `query:"verifiedType,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [XUserGetSubscriptionsParams]'s query parameters as `url.Values`.
func (r XUserGetSubscriptionsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type XUserTranslateBioParams struct {
	// Language code to translate into, such as `es`, `tr`, `pt-br`, or `zh-tw`. Upper
	// and lower case both work.
	Language string `query:"language" api:"required" json:"-"`
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

// URLQuery serializes [XUserTranslateBioParams]'s query parameters as `url.Values`.
func (r XUserTranslateBioParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Optional legacy pagination override. Search-only filters such as keywords apply
// only without it.
type XUserGetRepliesParamsMode string

const (
	XUserGetRepliesParamsModeStandard XUserGetRepliesParamsMode = "standard"
)

// Optional legacy pagination override. Search-only filters such as keywords apply
// only without it.
type XUserGetTweetsParamsMode string

const (
	XUserGetTweetsParamsModeStandard XUserGetTweetsParamsMode = "standard"
)
