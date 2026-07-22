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

// XTweetService contains methods and other services that help with interacting
// with the x-twitter-scraper API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewXTweetService] method instead.
type XTweetService struct {
	options []option.RequestOption
	// X write actions (tweets, likes, follows, DMs)
	Like XTweetLikeService
	// X write actions (tweets, likes, follows, DMs)
	Retweet XTweetRetweetService
}

// NewXTweetService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewXTweetService(opts ...option.RequestOption) (r XTweetService) {
	r = XTweetService{}
	r.options = opts
	r.Like = NewXTweetLikeService(opts...)
	r.Retweet = NewXTweetRetweetService(opts...)
	return
}

// Create tweet
func (r *XTweetService) New(ctx context.Context, params XTweetNewParams, opts ...option.RequestOption) (res *XTweetNewResponse, err error) {
	if !param.IsOmitted(params.IdempotencyKey) {
		opts = append(opts, option.WithHeader("Idempotency-Key", fmt.Sprintf("%v", params.IdempotencyKey)))
	}
	opts = slices.Concat(r.options, opts)
	path := "x/tweets"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Get tweet with full text, author, metrics and media
func (r *XTweetService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *XTweetGetResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("x/tweets/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Get multiple tweets by IDs
func (r *XTweetService) List(ctx context.Context, query XTweetListParams, opts ...option.RequestOption) (res *shared.PaginatedTweets, err error) {
	opts = slices.Concat(r.options, opts)
	path := "x/tweets"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Delete tweet
func (r *XTweetService) Delete(ctx context.Context, id string, params XTweetDeleteParams, opts ...option.RequestOption) (res *XTweetDeleteResponse, err error) {
	if !param.IsOmitted(params.IdempotencyKey) {
		opts = append(opts, option.WithHeader("Idempotency-Key", fmt.Sprintf("%v", params.IdempotencyKey)))
	}
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("x/tweets/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, params, &res, opts...)
	return res, err
}

// Returns liker profiles that X makes visible for the post. X can withhold liker
// identities even when the post reports likes. In that case this endpoint returns
// 424 `favoriters_unavailable` instead of a misleading empty success.
func (r *XTweetService) GetFavoriters(ctx context.Context, id string, query XTweetGetFavoritersParams, opts ...option.RequestOption) (res *shared.PaginatedUsers, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("x/tweets/%s/favoriters", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// List quote tweets of a tweet
func (r *XTweetService) GetQuotes(ctx context.Context, id string, query XTweetGetQuotesParams, opts ...option.RequestOption) (res *shared.PaginatedTweets, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("x/tweets/%s/quotes", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Returns visible replies. For an unfiltered first page, Xquik compares a terminal
// page with the post's reported reply count. If the page is visibly incomplete,
// the endpoint returns 424 `replies_incomplete` instead of presenting partial
// coverage as complete. Use tweet search with a `conversation_id:{id}` query as
// the broader fallback.
func (r *XTweetService) GetReplies(ctx context.Context, id string, query XTweetGetRepliesParams, opts ...option.RequestOption) (res *shared.PaginatedTweets, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("x/tweets/%s/replies", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// List users who retweeted a tweet
func (r *XTweetService) GetRetweeters(ctx context.Context, id string, query XTweetGetRetweetersParams, opts ...option.RequestOption) (res *shared.PaginatedUsers, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("x/tweets/%s/retweeters", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Get full conversation thread for a tweet
func (r *XTweetService) GetThread(ctx context.Context, id string, query XTweetGetThreadParams, opts ...option.RequestOption) (res *shared.PaginatedTweets, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("x/tweets/%s/thread", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Search tweets by query, Tweet ID, X status URL, or account date window
func (r *XTweetService) Search(ctx context.Context, query XTweetSearchParams, opts ...option.RequestOption) (res *shared.PaginatedTweets, err error) {
	opts = slices.Concat(r.options, opts)
	path := "x/tweets/search"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Tweet author profile. The lookup route always includes follower count and
// verification state. Other profile fields appear when available.
type TweetAuthor struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.UserProfile
}

// Returns the unmodified JSON received from the API
func (r TweetAuthor) RawJSON() string { return r.JSON.raw }
func (r *TweetAuthor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Full tweet with text, engagement metrics, media, and metadata. A zero metric can
// mean X did not report the count.
type TweetDetail struct {
	ID            string `json:"id" api:"required"`
	BookmarkCount int64  `json:"bookmarkCount" api:"required"`
	LikeCount     int64  `json:"likeCount" api:"required"`
	QuoteCount    int64  `json:"quoteCount" api:"required"`
	ReplyCount    int64  `json:"replyCount" api:"required"`
	RetweetCount  int64  `json:"retweetCount" api:"required"`
	Text          string `json:"text" api:"required"`
	ViewCount     int64  `json:"viewCount" api:"required"`
	// Tweet author profile. The lookup route always includes follower count and
	// verification state. Other profile fields appear when available.
	Author TweetAuthor `json:"author"`
	// Content disclosure metadata shown by X when a tweet is labeled as paid
	// partnership content or AI-generated media.
	ContentDisclosure shared.ContentDisclosure `json:"contentDisclosure"`
	// ID of the root tweet in the conversation thread
	ConversationID string `json:"conversationId"`
	CreatedAt      string `json:"createdAt"`
	// Start and end offsets for rendered tweet text
	DisplayTextRange []int64 `json:"displayTextRange"`
	// Parsed entities from the tweet text (URLs, mentions, hashtags, media)
	Entities map[string]any `json:"entities"`
	// Tweet ID being replied to
	InReplyToID string `json:"inReplyToId"`
	// User ID being replied to
	InReplyToUserID string `json:"inReplyToUserId"`
	// Username being replied to
	InReplyToUsername string `json:"inReplyToUsername"`
	// Whether replies are limited for this tweet
	IsLimitedReply bool `json:"isLimitedReply"`
	// Whether this is a Note Tweet (long-form post, up to 25,000 characters)
	IsNoteTweet bool `json:"isNoteTweet"`
	// Whether this tweet quotes another tweet
	IsQuoteStatus bool `json:"isQuoteStatus"`
	// Whether this tweet is a reply to another tweet
	IsReply bool `json:"isReply"`
	// Tweet language code
	Lang string `json:"lang"`
	// Attached media items, omitted when the tweet has no media
	Media []shared.TweetMedia `json:"media"`
	// Quoted or retweeted tweet context. Every object includes id, text, and
	// engagement metrics. A zero metric can mean X did not report the count. Author,
	// media, and conversation fields appear when available.
	QuotedTweet shared.EmbeddedTweet `json:"quoted_tweet"`
	// Quoted or retweeted tweet context. Every object includes id, text, and
	// engagement metrics. A zero metric can mean X did not report the count. Author,
	// media, and conversation fields appear when available.
	RetweetedTweet shared.EmbeddedTweet `json:"retweeted_tweet"`
	// Client application used to post this tweet
	Source string `json:"source"`
	// Tweet result type
	Type string `json:"type"`
	// Tweet permalink URL
	URL string `json:"url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		BookmarkCount     respjson.Field
		LikeCount         respjson.Field
		QuoteCount        respjson.Field
		ReplyCount        respjson.Field
		RetweetCount      respjson.Field
		Text              respjson.Field
		ViewCount         respjson.Field
		Author            respjson.Field
		ContentDisclosure respjson.Field
		ConversationID    respjson.Field
		CreatedAt         respjson.Field
		DisplayTextRange  respjson.Field
		Entities          respjson.Field
		InReplyToID       respjson.Field
		InReplyToUserID   respjson.Field
		InReplyToUsername respjson.Field
		IsLimitedReply    respjson.Field
		IsNoteTweet       respjson.Field
		IsQuoteStatus     respjson.Field
		IsReply           respjson.Field
		Lang              respjson.Field
		Media             respjson.Field
		QuotedTweet       respjson.Field
		RetweetedTweet    respjson.Field
		Source            respjson.Field
		Type              respjson.Field
		URL               respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TweetDetail) RawJSON() string { return r.JSON.raw }
func (r *TweetDetail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Durable write lifecycle record. Poll statusUrl until terminal is true. Reusing
// the original Idempotency-Key returns this same record. Submit a new write only
// when safeToRetry is true, using a new key.
type XTweetNewResponse struct {
	ID string `json:"id" api:"required"`
	// Connected account selected for the write.
	Account XTweetNewResponseAccount `json:"account" api:"required"`
	// Any of "create_tweet", "delete_tweet", "like", "unlike", "retweet", "unretweet",
	// "follow", "unfollow", "remove_follower", "send_dm", "upload_media",
	// "update_profile", "update_avatar", "update_banner", "create_community",
	// "delete_community", "join_community", "leave_community".
	Action XTweetNewResponseAction `json:"action" api:"required"`
	// plannedCredits is the approved maximum. chargedCredits comes from the settled
	// credit ledger. Pending or failed writes are not charged.
	Billing        XTweetNewResponseBilling `json:"billing" api:"required"`
	Charged        bool                     `json:"charged" api:"required"`
	ChargedCredits string                   `json:"chargedCredits" api:"required"`
	// Exact follow-up an API client or agent should perform.
	NextAction  XTweetNewResponseNextAction `json:"nextAction" api:"required"`
	Object      constant.XWriteAction       `json:"object" default:"x_write_action"`
	PollAfterMs int64                       `json:"pollAfterMs" api:"required"`
	// Stable fingerprint and sanitized payload for replay checks.
	Request XTweetNewResponseRequest `json:"request" api:"required"`
	// Confirmed result produced by the write, when available.
	Result XTweetNewResponseResult `json:"result" api:"required"`
	// True only when a new attempt can reasonably succeed.
	Retryable bool `json:"retryable" api:"required"`
	// True only when no write was dispatched and a new idempotency key may be used.
	SafeToRetry    bool `json:"safeToRetry" api:"required"`
	SendDispatched bool `json:"sendDispatched" api:"required"`
	// Any of "accepted", "dispatching", "pending_confirmation", "success", "failed",
	// "expired".
	Status    XTweetNewResponseStatus `json:"status" api:"required"`
	StatusURL string                  `json:"statusUrl" api:"required"`
	Success   bool                    `json:"success" api:"required"`
	// Existing X resource targeted by the write, when applicable.
	Target        XTweetNewResponseTarget `json:"target" api:"required"`
	TargetID      string                  `json:"targetId" api:"required"`
	Terminal      bool                    `json:"terminal" api:"required"`
	WriteActionID string                  `json:"writeActionId" api:"required"`
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
func (r XTweetNewResponse) RawJSON() string { return r.JSON.raw }
func (r *XTweetNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Connected account selected for the write.
type XTweetNewResponseAccount struct {
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
func (r XTweetNewResponseAccount) RawJSON() string { return r.JSON.raw }
func (r *XTweetNewResponseAccount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type XTweetNewResponseAction string

const (
	XTweetNewResponseActionCreateTweet     XTweetNewResponseAction = "create_tweet"
	XTweetNewResponseActionDeleteTweet     XTweetNewResponseAction = "delete_tweet"
	XTweetNewResponseActionLike            XTweetNewResponseAction = "like"
	XTweetNewResponseActionUnlike          XTweetNewResponseAction = "unlike"
	XTweetNewResponseActionRetweet         XTweetNewResponseAction = "retweet"
	XTweetNewResponseActionUnretweet       XTweetNewResponseAction = "unretweet"
	XTweetNewResponseActionFollow          XTweetNewResponseAction = "follow"
	XTweetNewResponseActionUnfollow        XTweetNewResponseAction = "unfollow"
	XTweetNewResponseActionRemoveFollower  XTweetNewResponseAction = "remove_follower"
	XTweetNewResponseActionSendDm          XTweetNewResponseAction = "send_dm"
	XTweetNewResponseActionUploadMedia     XTweetNewResponseAction = "upload_media"
	XTweetNewResponseActionUpdateProfile   XTweetNewResponseAction = "update_profile"
	XTweetNewResponseActionUpdateAvatar    XTweetNewResponseAction = "update_avatar"
	XTweetNewResponseActionUpdateBanner    XTweetNewResponseAction = "update_banner"
	XTweetNewResponseActionCreateCommunity XTweetNewResponseAction = "create_community"
	XTweetNewResponseActionDeleteCommunity XTweetNewResponseAction = "delete_community"
	XTweetNewResponseActionJoinCommunity   XTweetNewResponseAction = "join_community"
	XTweetNewResponseActionLeaveCommunity  XTweetNewResponseAction = "leave_community"
)

// plannedCredits is the approved maximum. chargedCredits comes from the settled
// credit ledger. Pending or failed writes are not charged.
type XTweetNewResponseBilling struct {
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
func (r XTweetNewResponseBilling) RawJSON() string { return r.JSON.raw }
func (r *XTweetNewResponseBilling) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Exact follow-up an API client or agent should perform.
type XTweetNewResponseNextAction struct {
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
func (r XTweetNewResponseNextAction) RawJSON() string { return r.JSON.raw }
func (r *XTweetNewResponseNextAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Stable fingerprint and sanitized payload for replay checks.
type XTweetNewResponseRequest struct {
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
func (r XTweetNewResponseRequest) RawJSON() string { return r.JSON.raw }
func (r *XTweetNewResponseRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Confirmed result produced by the write, when available.
type XTweetNewResponseResult struct {
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
func (r XTweetNewResponseResult) RawJSON() string { return r.JSON.raw }
func (r *XTweetNewResponseResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type XTweetNewResponseStatus string

const (
	XTweetNewResponseStatusAccepted            XTweetNewResponseStatus = "accepted"
	XTweetNewResponseStatusDispatching         XTweetNewResponseStatus = "dispatching"
	XTweetNewResponseStatusPendingConfirmation XTweetNewResponseStatus = "pending_confirmation"
	XTweetNewResponseStatusSuccess             XTweetNewResponseStatus = "success"
	XTweetNewResponseStatusFailed              XTweetNewResponseStatus = "failed"
	XTweetNewResponseStatusExpired             XTweetNewResponseStatus = "expired"
)

// Existing X resource targeted by the write, when applicable.
type XTweetNewResponseTarget struct {
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
func (r XTweetNewResponseTarget) RawJSON() string { return r.JSON.raw }
func (r *XTweetNewResponseTarget) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type XTweetGetResponse struct {
	// Full tweet with text, engagement metrics, media, and metadata. A zero metric can
	// mean X did not report the count.
	Tweet TweetDetail `json:"tweet" api:"required"`
	// Tweet author profile. The lookup route always includes follower count and
	// verification state. Other profile fields appear when available.
	Author TweetAuthor `json:"author"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Tweet       respjson.Field
		Author      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XTweetGetResponse) RawJSON() string { return r.JSON.raw }
func (r *XTweetGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Durable write lifecycle record. Poll statusUrl until terminal is true. Reusing
// the original Idempotency-Key returns this same record. Submit a new write only
// when safeToRetry is true, using a new key.
type XTweetDeleteResponse struct {
	ID string `json:"id" api:"required"`
	// Connected account selected for the write.
	Account XTweetDeleteResponseAccount `json:"account" api:"required"`
	// Any of "create_tweet", "delete_tweet", "like", "unlike", "retweet", "unretweet",
	// "follow", "unfollow", "remove_follower", "send_dm", "upload_media",
	// "update_profile", "update_avatar", "update_banner", "create_community",
	// "delete_community", "join_community", "leave_community".
	Action XTweetDeleteResponseAction `json:"action" api:"required"`
	// plannedCredits is the approved maximum. chargedCredits comes from the settled
	// credit ledger. Pending or failed writes are not charged.
	Billing        XTweetDeleteResponseBilling `json:"billing" api:"required"`
	Charged        bool                        `json:"charged" api:"required"`
	ChargedCredits string                      `json:"chargedCredits" api:"required"`
	// Exact follow-up an API client or agent should perform.
	NextAction  XTweetDeleteResponseNextAction `json:"nextAction" api:"required"`
	Object      constant.XWriteAction          `json:"object" default:"x_write_action"`
	PollAfterMs int64                          `json:"pollAfterMs" api:"required"`
	// Stable fingerprint and sanitized payload for replay checks.
	Request XTweetDeleteResponseRequest `json:"request" api:"required"`
	// Confirmed result produced by the write, when available.
	Result XTweetDeleteResponseResult `json:"result" api:"required"`
	// True only when a new attempt can reasonably succeed.
	Retryable bool `json:"retryable" api:"required"`
	// True only when no write was dispatched and a new idempotency key may be used.
	SafeToRetry    bool `json:"safeToRetry" api:"required"`
	SendDispatched bool `json:"sendDispatched" api:"required"`
	// Any of "accepted", "dispatching", "pending_confirmation", "success", "failed",
	// "expired".
	Status    XTweetDeleteResponseStatus `json:"status" api:"required"`
	StatusURL string                     `json:"statusUrl" api:"required"`
	Success   bool                       `json:"success" api:"required"`
	// Existing X resource targeted by the write, when applicable.
	Target        XTweetDeleteResponseTarget `json:"target" api:"required"`
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
func (r XTweetDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *XTweetDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Connected account selected for the write.
type XTweetDeleteResponseAccount struct {
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
func (r XTweetDeleteResponseAccount) RawJSON() string { return r.JSON.raw }
func (r *XTweetDeleteResponseAccount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type XTweetDeleteResponseAction string

const (
	XTweetDeleteResponseActionCreateTweet     XTweetDeleteResponseAction = "create_tweet"
	XTweetDeleteResponseActionDeleteTweet     XTweetDeleteResponseAction = "delete_tweet"
	XTweetDeleteResponseActionLike            XTweetDeleteResponseAction = "like"
	XTweetDeleteResponseActionUnlike          XTweetDeleteResponseAction = "unlike"
	XTweetDeleteResponseActionRetweet         XTweetDeleteResponseAction = "retweet"
	XTweetDeleteResponseActionUnretweet       XTweetDeleteResponseAction = "unretweet"
	XTweetDeleteResponseActionFollow          XTweetDeleteResponseAction = "follow"
	XTweetDeleteResponseActionUnfollow        XTweetDeleteResponseAction = "unfollow"
	XTweetDeleteResponseActionRemoveFollower  XTweetDeleteResponseAction = "remove_follower"
	XTweetDeleteResponseActionSendDm          XTweetDeleteResponseAction = "send_dm"
	XTweetDeleteResponseActionUploadMedia     XTweetDeleteResponseAction = "upload_media"
	XTweetDeleteResponseActionUpdateProfile   XTweetDeleteResponseAction = "update_profile"
	XTweetDeleteResponseActionUpdateAvatar    XTweetDeleteResponseAction = "update_avatar"
	XTweetDeleteResponseActionUpdateBanner    XTweetDeleteResponseAction = "update_banner"
	XTweetDeleteResponseActionCreateCommunity XTweetDeleteResponseAction = "create_community"
	XTweetDeleteResponseActionDeleteCommunity XTweetDeleteResponseAction = "delete_community"
	XTweetDeleteResponseActionJoinCommunity   XTweetDeleteResponseAction = "join_community"
	XTweetDeleteResponseActionLeaveCommunity  XTweetDeleteResponseAction = "leave_community"
)

// plannedCredits is the approved maximum. chargedCredits comes from the settled
// credit ledger. Pending or failed writes are not charged.
type XTweetDeleteResponseBilling struct {
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
func (r XTweetDeleteResponseBilling) RawJSON() string { return r.JSON.raw }
func (r *XTweetDeleteResponseBilling) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Exact follow-up an API client or agent should perform.
type XTweetDeleteResponseNextAction struct {
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
func (r XTweetDeleteResponseNextAction) RawJSON() string { return r.JSON.raw }
func (r *XTweetDeleteResponseNextAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Stable fingerprint and sanitized payload for replay checks.
type XTweetDeleteResponseRequest struct {
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
func (r XTweetDeleteResponseRequest) RawJSON() string { return r.JSON.raw }
func (r *XTweetDeleteResponseRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Confirmed result produced by the write, when available.
type XTweetDeleteResponseResult struct {
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
func (r XTweetDeleteResponseResult) RawJSON() string { return r.JSON.raw }
func (r *XTweetDeleteResponseResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type XTweetDeleteResponseStatus string

const (
	XTweetDeleteResponseStatusAccepted            XTweetDeleteResponseStatus = "accepted"
	XTweetDeleteResponseStatusDispatching         XTweetDeleteResponseStatus = "dispatching"
	XTweetDeleteResponseStatusPendingConfirmation XTweetDeleteResponseStatus = "pending_confirmation"
	XTweetDeleteResponseStatusSuccess             XTweetDeleteResponseStatus = "success"
	XTweetDeleteResponseStatusFailed              XTweetDeleteResponseStatus = "failed"
	XTweetDeleteResponseStatusExpired             XTweetDeleteResponseStatus = "expired"
)

// Existing X resource targeted by the write, when applicable.
type XTweetDeleteResponseTarget struct {
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
func (r XTweetDeleteResponseTarget) RawJSON() string { return r.JSON.raw }
func (r *XTweetDeleteResponseTarget) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type XTweetNewParams struct {
	// X account (@username or account ID)
	Account        string            `json:"account" api:"required"`
	IdempotencyKey string            `header:"Idempotency-Key" api:"required" json:"-"`
	CommunityID    param.Opt[string] `json:"community_id,omitzero"`
	IsNoteTweet    param.Opt[bool]   `json:"is_note_tweet,omitzero"`
	ReplyToTweetID param.Opt[string] `json:"reply_to_tweet_id,omitzero"`
	// Tweet text (optional when media is provided)
	Text param.Opt[string] `json:"text,omitzero"`
	// Array of public media URLs to attach. Supports up to 4 images or exactly 1 MP4
	// video up to 100 MB. Each URL must be publicly reachable. Attached media adds 2
	// credits per started MB across all files.
	Media []string `json:"media,omitzero"`
	paramObj
}

func (r XTweetNewParams) MarshalJSON() (data []byte, err error) {
	type shadow XTweetNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *XTweetNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type XTweetListParams struct {
	// Comma-separated tweet IDs (max 100)
	IDs string `query:"ids" api:"required" json:"-"`
	paramObj
}

// URLQuery serializes [XTweetListParams]'s query parameters as `url.Values`.
func (r XTweetListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type XTweetDeleteParams struct {
	// X account identifier (@username or account ID)
	Account        string `json:"account" api:"required"`
	IdempotencyKey string `header:"Idempotency-Key" api:"required" json:"-"`
	paramObj
}

func (r XTweetDeleteParams) MarshalJSON() (data []byte, err error) {
	type shadow XTweetDeleteParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *XTweetDeleteParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type XTweetGetFavoritersParams struct {
	// Pagination cursor for favoriters
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Maximum user profiles requested from this page (20-200, default 200). The
	// response can contain fewer profiles because the source returned fewer or
	// remaining credits cover fewer results. Keep requesting next_cursor while
	// has_next_page is true. The deprecated limit and count aliases remain accepted.
	PageSize param.Opt[int64] `query:"pageSize,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [XTweetGetFavoritersParams]'s query parameters as
// `url.Values`.
func (r XTweetGetFavoritersParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type XTweetGetQuotesParams struct {
	// Words or quoted phrases where any one can match. Separate with spaces, commas,
	// or lines.
	AnyWords param.Opt[string] `query:"anyWords,omitzero" json:"-"`
	// Cashtags separated by spaces, commas, or lines.
	Cashtags param.Opt[string] `query:"cashtags,omitzero" json:"-"`
	// Conversation ID filter.
	ConversationID param.Opt[string] `query:"conversationId,omitzero" json:"-"`
	// Pagination cursor for quote tweets
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Exact phrase to match.
	ExactPhrase param.Opt[string] `query:"exactPhrase,omitzero" json:"-"`
	// Words or quoted phrases to exclude. Separate with spaces, commas, or lines.
	ExcludeWords param.Opt[string] `query:"excludeWords,omitzero" json:"-"`
	// Filter by author username.
	FromUser param.Opt[string] `query:"fromUser,omitzero" json:"-"`
	// Hashtags separated by spaces, commas, or lines.
	Hashtags param.Opt[string] `query:"hashtags,omitzero" json:"-"`
	// Include reply quotes (default false)
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
	// Maximum items requested from this page (1-100, default 20). The response can
	// contain fewer items because the source returned fewer, filters removed items, or
	// remaining credits cover fewer results. Keep requesting next_cursor while
	// has_next_page is true, even when a page is empty. The deprecated limit and count
	// aliases remain accepted.
	PageSize param.Opt[int64] `query:"pageSize,omitzero" json:"-"`
	// Only quotes of this tweet ID.
	QuotesOfTweetID param.Opt[string] `query:"quotesOfTweetId,omitzero" json:"-"`
	// Only retweets of this tweet ID.
	RetweetsOfTweetID param.Opt[string] `query:"retweetsOfTweetId,omitzero" json:"-"`
	// Start date in YYYY-MM-DD format.
	SinceDate param.Opt[time.Time] `query:"sinceDate,omitzero" format:"date" json:"-"`
	// Unix timestamp - return quotes posted after this time
	SinceTime param.Opt[string] `query:"sinceTime,omitzero" json:"-"`
	// Filter replies sent to a username.
	ToUser param.Opt[string] `query:"toUser,omitzero" json:"-"`
	// End date in YYYY-MM-DD format.
	UntilDate param.Opt[time.Time] `query:"untilDate,omitzero" format:"date" json:"-"`
	// Unix timestamp - return quotes posted before this time
	UntilTime param.Opt[string] `query:"untilTime,omitzero" json:"-"`
	// URL substring or domain filter.
	URL param.Opt[string] `query:"url,omitzero" json:"-"`
	// Only return tweets from verified authors.
	VerifiedOnly param.Opt[bool] `query:"verifiedOnly,omitzero" json:"-"`
	// Filter by media type.
	//
	// Any of "images", "videos", "gifs", "media", "links", "none".
	MediaType XTweetGetQuotesParamsMediaType `query:"mediaType,omitzero" json:"-"`
	// Quote mode.
	//
	// Any of "include", "exclude", "only".
	Quotes XTweetGetQuotesParamsQuotes `query:"quotes,omitzero" json:"-"`
	// Reply mode.
	//
	// Any of "include", "exclude", "only".
	Replies XTweetGetQuotesParamsReplies `query:"replies,omitzero" json:"-"`
	// Retweet mode.
	//
	// Any of "include", "exclude", "only".
	Retweets XTweetGetQuotesParamsRetweets `query:"retweets,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [XTweetGetQuotesParams]'s query parameters as `url.Values`.
func (r XTweetGetQuotesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by media type.
type XTweetGetQuotesParamsMediaType string

const (
	XTweetGetQuotesParamsMediaTypeImages XTweetGetQuotesParamsMediaType = "images"
	XTweetGetQuotesParamsMediaTypeVideos XTweetGetQuotesParamsMediaType = "videos"
	XTweetGetQuotesParamsMediaTypeGifs   XTweetGetQuotesParamsMediaType = "gifs"
	XTweetGetQuotesParamsMediaTypeMedia  XTweetGetQuotesParamsMediaType = "media"
	XTweetGetQuotesParamsMediaTypeLinks  XTweetGetQuotesParamsMediaType = "links"
	XTweetGetQuotesParamsMediaTypeNone   XTweetGetQuotesParamsMediaType = "none"
)

// Quote mode.
type XTweetGetQuotesParamsQuotes string

const (
	XTweetGetQuotesParamsQuotesInclude XTweetGetQuotesParamsQuotes = "include"
	XTweetGetQuotesParamsQuotesExclude XTweetGetQuotesParamsQuotes = "exclude"
	XTweetGetQuotesParamsQuotesOnly    XTweetGetQuotesParamsQuotes = "only"
)

// Reply mode.
type XTweetGetQuotesParamsReplies string

const (
	XTweetGetQuotesParamsRepliesInclude XTweetGetQuotesParamsReplies = "include"
	XTweetGetQuotesParamsRepliesExclude XTweetGetQuotesParamsReplies = "exclude"
	XTweetGetQuotesParamsRepliesOnly    XTweetGetQuotesParamsReplies = "only"
)

// Retweet mode.
type XTweetGetQuotesParamsRetweets string

const (
	XTweetGetQuotesParamsRetweetsInclude XTweetGetQuotesParamsRetweets = "include"
	XTweetGetQuotesParamsRetweetsExclude XTweetGetQuotesParamsRetweets = "exclude"
	XTweetGetQuotesParamsRetweetsOnly    XTweetGetQuotesParamsRetweets = "only"
)

type XTweetGetRepliesParams struct {
	// Words or quoted phrases where any one can match. Separate with spaces, commas,
	// or lines.
	AnyWords param.Opt[string] `query:"anyWords,omitzero" json:"-"`
	// Cashtags separated by spaces, commas, or lines.
	Cashtags param.Opt[string] `query:"cashtags,omitzero" json:"-"`
	// Conversation ID filter.
	ConversationID param.Opt[string] `query:"conversationId,omitzero" json:"-"`
	// Pagination cursor for tweet replies
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
	// Maximum items requested from this page (1-100, default 20). The response can
	// contain fewer items because the source returned fewer, filters removed items, or
	// remaining credits cover fewer results. Keep requesting next_cursor while
	// has_next_page is true, even when a page is empty. The deprecated limit and count
	// aliases remain accepted.
	PageSize param.Opt[int64] `query:"pageSize,omitzero" json:"-"`
	// Only quotes of this tweet ID.
	QuotesOfTweetID param.Opt[string] `query:"quotesOfTweetId,omitzero" json:"-"`
	// Only retweets of this tweet ID.
	RetweetsOfTweetID param.Opt[string] `query:"retweetsOfTweetId,omitzero" json:"-"`
	// Start date in YYYY-MM-DD format.
	SinceDate param.Opt[time.Time] `query:"sinceDate,omitzero" format:"date" json:"-"`
	// Unix timestamp - return replies posted after this time
	SinceTime param.Opt[string] `query:"sinceTime,omitzero" json:"-"`
	// Filter replies sent to a username.
	ToUser param.Opt[string] `query:"toUser,omitzero" json:"-"`
	// End date in YYYY-MM-DD format.
	UntilDate param.Opt[time.Time] `query:"untilDate,omitzero" format:"date" json:"-"`
	// Unix timestamp - return replies posted before this time
	UntilTime param.Opt[string] `query:"untilTime,omitzero" json:"-"`
	// URL substring or domain filter.
	URL param.Opt[string] `query:"url,omitzero" json:"-"`
	// Only return tweets from verified authors.
	VerifiedOnly param.Opt[bool] `query:"verifiedOnly,omitzero" json:"-"`
	// Filter by media type.
	//
	// Any of "images", "videos", "gifs", "media", "links", "none".
	MediaType XTweetGetRepliesParamsMediaType `query:"mediaType,omitzero" json:"-"`
	// Quote mode.
	//
	// Any of "include", "exclude", "only".
	Quotes XTweetGetRepliesParamsQuotes `query:"quotes,omitzero" json:"-"`
	// Reply mode.
	//
	// Any of "include", "exclude", "only".
	Replies XTweetGetRepliesParamsReplies `query:"replies,omitzero" json:"-"`
	// Retweet mode.
	//
	// Any of "include", "exclude", "only".
	Retweets XTweetGetRepliesParamsRetweets `query:"retweets,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [XTweetGetRepliesParams]'s query parameters as `url.Values`.
func (r XTweetGetRepliesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by media type.
type XTweetGetRepliesParamsMediaType string

const (
	XTweetGetRepliesParamsMediaTypeImages XTweetGetRepliesParamsMediaType = "images"
	XTweetGetRepliesParamsMediaTypeVideos XTweetGetRepliesParamsMediaType = "videos"
	XTweetGetRepliesParamsMediaTypeGifs   XTweetGetRepliesParamsMediaType = "gifs"
	XTweetGetRepliesParamsMediaTypeMedia  XTweetGetRepliesParamsMediaType = "media"
	XTweetGetRepliesParamsMediaTypeLinks  XTweetGetRepliesParamsMediaType = "links"
	XTweetGetRepliesParamsMediaTypeNone   XTweetGetRepliesParamsMediaType = "none"
)

// Quote mode.
type XTweetGetRepliesParamsQuotes string

const (
	XTweetGetRepliesParamsQuotesInclude XTweetGetRepliesParamsQuotes = "include"
	XTweetGetRepliesParamsQuotesExclude XTweetGetRepliesParamsQuotes = "exclude"
	XTweetGetRepliesParamsQuotesOnly    XTweetGetRepliesParamsQuotes = "only"
)

// Reply mode.
type XTweetGetRepliesParamsReplies string

const (
	XTweetGetRepliesParamsRepliesInclude XTweetGetRepliesParamsReplies = "include"
	XTweetGetRepliesParamsRepliesExclude XTweetGetRepliesParamsReplies = "exclude"
	XTweetGetRepliesParamsRepliesOnly    XTweetGetRepliesParamsReplies = "only"
)

// Retweet mode.
type XTweetGetRepliesParamsRetweets string

const (
	XTweetGetRepliesParamsRetweetsInclude XTweetGetRepliesParamsRetweets = "include"
	XTweetGetRepliesParamsRetweetsExclude XTweetGetRepliesParamsRetweets = "exclude"
	XTweetGetRepliesParamsRetweetsOnly    XTweetGetRepliesParamsRetweets = "only"
)

type XTweetGetRetweetersParams struct {
	// Pagination cursor for retweeters
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Maximum user profiles requested from this page (20-200, default 200). The
	// response can contain fewer profiles because the source returned fewer or
	// remaining credits cover fewer results. Keep requesting next_cursor while
	// has_next_page is true. The deprecated limit and count aliases remain accepted.
	PageSize param.Opt[int64] `query:"pageSize,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [XTweetGetRetweetersParams]'s query parameters as
// `url.Values`.
func (r XTweetGetRetweetersParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type XTweetGetThreadParams struct {
	// Pagination cursor for thread tweets
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Maximum items requested from this page (1-100, default 20). The response can
	// contain fewer items because the source returned fewer, filters removed items, or
	// remaining credits cover fewer results. Keep requesting next_cursor while
	// has_next_page is true, even when a page is empty. The deprecated limit and count
	// aliases remain accepted.
	PageSize param.Opt[int64] `query:"pageSize,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [XTweetGetThreadParams]'s query parameters as `url.Values`.
func (r XTweetGetThreadParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type XTweetSearchParams struct {
	// Search query (keywords,
	Q string `query:"q" api:"required" json:"-"`
	// Raw advanced search query appended as-is.
	AdvancedQuery param.Opt[string] `query:"advancedQuery,omitzero" json:"-"`
	// Words or quoted phrases where any one can match. Separate with spaces, commas,
	// or lines.
	AnyWords param.Opt[string] `query:"anyWords,omitzero" json:"-"`
	// Geo bounding box, e.g. -74.1 40.6 -73.9 40.8.
	BoundingBox param.Opt[string] `query:"boundingBox,omitzero" json:"-"`
	// Cashtags separated by spaces, commas, or lines.
	Cashtags param.Opt[string] `query:"cashtags,omitzero" json:"-"`
	// Conversation ID filter.
	ConversationID param.Opt[string] `query:"conversationId,omitzero" json:"-"`
	// Pagination cursor from previous response
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
	// Max tweets to return (server paginates internally). Omit for single page (~20).
	// This is an upper bound for paid authenticated calls: remaining credits can
	// reduce the returned page size, and zero affordable results returns 402
	// insufficient_credits.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Search within a list ID.
	ListID param.Opt[string] `query:"listId,omitzero" json:"-"`
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
	// Search within a place ID.
	Place param.Opt[string] `query:"place,omitzero" json:"-"`
	// Search within a country code.
	PlaceCountry param.Opt[string] `query:"placeCountry,omitzero" json:"-"`
	// Geo point radius, e.g. -73.99 40.73 25mi.
	PointRadius param.Opt[string] `query:"pointRadius,omitzero" json:"-"`
	// Only quotes of this tweet ID.
	QuotesOfTweetID param.Opt[string] `query:"quotesOfTweetId,omitzero" json:"-"`
	// Only retweets of this tweet ID.
	RetweetsOfTweetID param.Opt[string] `query:"retweetsOfTweetId,omitzero" json:"-"`
	// Start date in YYYY-MM-DD format.
	SinceDate param.Opt[time.Time] `query:"sinceDate,omitzero" format:"date" json:"-"`
	// ISO 8601 timestamp - only return tweets after this time
	SinceTime param.Opt[string] `query:"sinceTime,omitzero" json:"-"`
	// Filter replies sent to a username.
	ToUser param.Opt[string] `query:"toUser,omitzero" json:"-"`
	// End date in YYYY-MM-DD format.
	UntilDate param.Opt[time.Time] `query:"untilDate,omitzero" format:"date" json:"-"`
	// ISO 8601 timestamp - only return tweets before this time
	UntilTime param.Opt[string] `query:"untilTime,omitzero" json:"-"`
	// URL substring or domain filter.
	URL param.Opt[string] `query:"url,omitzero" json:"-"`
	// Only return tweets from verified authors.
	VerifiedOnly param.Opt[bool] `query:"verifiedOnly,omitzero" json:"-"`
	// Filter by media type.
	//
	// Any of "images", "videos", "gifs", "media", "links", "none".
	MediaType XTweetSearchParamsMediaType `query:"mediaType,omitzero" json:"-"`
	// Sort order - Latest (chronological) or Top (engagement-ranked)
	//
	// Any of "Latest", "Top".
	QueryType XTweetSearchParamsQueryType `query:"queryType,omitzero" json:"-"`
	// Quote mode.
	//
	// Any of "include", "exclude", "only".
	Quotes XTweetSearchParamsQuotes `query:"quotes,omitzero" json:"-"`
	// Reply mode.
	//
	// Any of "include", "exclude", "only".
	Replies XTweetSearchParamsReplies `query:"replies,omitzero" json:"-"`
	// Retweet mode.
	//
	// Any of "include", "exclude", "only".
	Retweets XTweetSearchParamsRetweets `query:"retweets,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [XTweetSearchParams]'s query parameters as `url.Values`.
func (r XTweetSearchParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by media type.
type XTweetSearchParamsMediaType string

const (
	XTweetSearchParamsMediaTypeImages XTweetSearchParamsMediaType = "images"
	XTweetSearchParamsMediaTypeVideos XTweetSearchParamsMediaType = "videos"
	XTweetSearchParamsMediaTypeGifs   XTweetSearchParamsMediaType = "gifs"
	XTweetSearchParamsMediaTypeMedia  XTweetSearchParamsMediaType = "media"
	XTweetSearchParamsMediaTypeLinks  XTweetSearchParamsMediaType = "links"
	XTweetSearchParamsMediaTypeNone   XTweetSearchParamsMediaType = "none"
)

// Sort order - Latest (chronological) or Top (engagement-ranked)
type XTweetSearchParamsQueryType string

const (
	XTweetSearchParamsQueryTypeLatest XTweetSearchParamsQueryType = "Latest"
	XTweetSearchParamsQueryTypeTop    XTweetSearchParamsQueryType = "Top"
)

// Quote mode.
type XTweetSearchParamsQuotes string

const (
	XTweetSearchParamsQuotesInclude XTweetSearchParamsQuotes = "include"
	XTweetSearchParamsQuotesExclude XTweetSearchParamsQuotes = "exclude"
	XTweetSearchParamsQuotesOnly    XTweetSearchParamsQuotes = "only"
)

// Reply mode.
type XTweetSearchParamsReplies string

const (
	XTweetSearchParamsRepliesInclude XTweetSearchParamsReplies = "include"
	XTweetSearchParamsRepliesExclude XTweetSearchParamsReplies = "exclude"
	XTweetSearchParamsRepliesOnly    XTweetSearchParamsReplies = "only"
)

// Retweet mode.
type XTweetSearchParamsRetweets string

const (
	XTweetSearchParamsRetweetsInclude XTweetSearchParamsRetweets = "include"
	XTweetSearchParamsRetweetsExclude XTweetSearchParamsRetweets = "exclude"
	XTweetSearchParamsRetweetsOnly    XTweetSearchParamsRetweets = "only"
)
