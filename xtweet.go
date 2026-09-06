// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

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

// Publishes a post through a connected X account.
func (r *XTweetService) New(ctx context.Context, params XTweetNewParams, opts ...option.RequestOption) (res *XTweetNewResponse, err error) {
	if !param.IsOmitted(params.IdempotencyKey) {
		opts = append(opts, option.WithHeader("Idempotency-Key", fmt.Sprintf("%v", params.IdempotencyKey)))
	}
	opts = slices.Concat(r.options, opts)
	path := "x/tweets"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Returns one public tweet with author, metrics, and media.
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

// Returns public tweet records for the requested IDs.
func (r *XTweetService) List(ctx context.Context, query XTweetListParams, opts ...option.RequestOption) (res *shared.PaginatedTweets, err error) {
	opts = slices.Concat(r.options, opts)
	path := "x/tweets"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Deletes an authored post through a connected X account.
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

// Returns public posts quoting the selected tweet.
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

// Returns direct replies with automatic maximum coverage and pagination. Complete
// mode adds nested replies, diagnostics, and a 424 below 80% coverage.
func (r *XTweetService) GetReplies(ctx context.Context, id string, query XTweetGetRepliesParams, opts ...option.RequestOption) (res *XTweetGetRepliesResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("x/tweets/%s/replies", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Returns public profiles that reposted the selected tweet.
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

// Returns visible posts from the selected conversation thread.
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

// Returns normalized tweets with author, like count, media, URL, and cursors. Set
// q, limit, queryType, and minLikes. Omit mode for maximum coverage. Reuse
// next_cursor without changing filters.
func (r *XTweetService) Search(ctx context.Context, query XTweetSearchParams, opts ...option.RequestOption) (res *XTweetSearchResponseUnion, err error) {
	opts = slices.Concat(r.options, opts)
	path := "x/tweets/search"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Tweet author profile. The lookup route always includes follower count and
// verification state. Other profile fields appear when available.
type TweetAuthor struct {
	Followers int64 `json:"followers" api:"required"`
	Verified  bool  `json:"verified" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Followers   respjson.Field
		Verified    respjson.Field
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
	// Describes an X Article preview and its lifecycle metadata.
	Article TweetDetailArticle `json:"article"`
	// Tweet author profile. The lookup route always includes follower count and
	// verification state. Other profile fields appear when available.
	Author TweetAuthor `json:"author"`
	// Describes a public card and its referenced profiles.
	Card TweetDetailCard `json:"card"`
	// Community ID.
	CommunityID string `json:"communityId"`
	// Community Note presentation metadata returned by X.
	CommunityNote TweetDetailCommunityNote `json:"communityNote"`
	// Content disclosure metadata shown by X when a tweet is labeled as paid
	// partnership content or AI-generated media.
	ContentDisclosure shared.ContentDisclosure `json:"contentDisclosure"`
	// Public reply policy and conversation owner.
	ConversationControl TweetDetailConversationControl `json:"conversationControl"`
	ConversationID      string                         `json:"conversationId"`
	CreatedAt           string                         `json:"createdAt"`
	DisplayTextRange    []int64                        `json:"displayTextRange"`
	// Lists edit-chain identifiers and the remaining edit window.
	Edit TweetDetailEdit `json:"edit"`
	// Lists hashtags, symbols, links, and mentions from tweet text.
	Entities TweetDetailEntities `json:"entities"`
	// Public metadata whose fields are defined by X.
	GrokShareAttachment map[string]any `json:"grokShareAttachment"`
	InReplyToID         string         `json:"inReplyToId"`
	InReplyToUserID     string         `json:"inReplyToUserId"`
	InReplyToUsername   string         `json:"inReplyToUsername"`
	IsLimitedReply      bool           `json:"isLimitedReply"`
	IsNoteTweet         bool           `json:"isNoteTweet"`
	IsQuoteStatus       bool           `json:"isQuoteStatus"`
	IsReply             bool           `json:"isReply"`
	IsTranslatable      bool           `json:"isTranslatable"`
	// Public metadata whose fields are defined by X.
	JetfuelAttachment map[string]any `json:"jetfuelAttachment"`
	Lang              string         `json:"lang"`
	// Public interaction restrictions and user-facing prompts.
	LimitedActions []TweetDetailLimitedAction `json:"limitedActions"`
	// Attached media items, omitted when unavailable.
	Media []shared.TweetMedia `json:"media"`
	// Complete Note Tweet content and rich-text metadata.
	NoteTweet TweetDetailNoteTweet `json:"noteTweet"`
	// Describes public place metadata on a geotagged tweet.
	Place             TweetDetailPlace `json:"place"`
	PossiblySensitive bool             `json:"possiblySensitive"`
	// Public metadata whose fields are defined by X.
	PostCta map[string]any `json:"postCta"`
	// Engagement counts retained from a prior tweet edit.
	PreviousCounts TweetDetailPreviousCounts `json:"previousCounts"`
	// Quoted or retweeted tweet context.
	QuotedTweet shared.EmbeddedTweet `json:"quoted_tweet"`
	// Quoted tweet ID.
	QuotedTweetID string `json:"quotedTweetId"`
	// Public post and user referenced by this reaction.
	ReactionContext TweetDetailReactionContext `json:"reactionContext"`
	// Quoted or retweeted tweet context.
	RetweetedTweet shared.EmbeddedTweet `json:"retweeted_tweet"`
	// Public metadata whose fields are defined by X.
	Scopes map[string]any `json:"scopes"`
	Source string         `json:"source"`
	// Public visibility notice attached to an available tweet.
	Tombstone TweetDetailTombstone `json:"tombstone"`
	Type      string               `json:"type"`
	// User IDs that left this conversation.
	UnmentionedUserIDs []string `json:"unmentionedUserIds"`
	URL                string   `json:"url"`
	ViewState          string   `json:"viewState"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                  respjson.Field
		BookmarkCount       respjson.Field
		LikeCount           respjson.Field
		QuoteCount          respjson.Field
		ReplyCount          respjson.Field
		RetweetCount        respjson.Field
		Text                respjson.Field
		ViewCount           respjson.Field
		Article             respjson.Field
		Author              respjson.Field
		Card                respjson.Field
		CommunityID         respjson.Field
		CommunityNote       respjson.Field
		ContentDisclosure   respjson.Field
		ConversationControl respjson.Field
		ConversationID      respjson.Field
		CreatedAt           respjson.Field
		DisplayTextRange    respjson.Field
		Edit                respjson.Field
		Entities            respjson.Field
		GrokShareAttachment respjson.Field
		InReplyToID         respjson.Field
		InReplyToUserID     respjson.Field
		InReplyToUsername   respjson.Field
		IsLimitedReply      respjson.Field
		IsNoteTweet         respjson.Field
		IsQuoteStatus       respjson.Field
		IsReply             respjson.Field
		IsTranslatable      respjson.Field
		JetfuelAttachment   respjson.Field
		Lang                respjson.Field
		LimitedActions      respjson.Field
		Media               respjson.Field
		NoteTweet           respjson.Field
		Place               respjson.Field
		PossiblySensitive   respjson.Field
		PostCta             respjson.Field
		PreviousCounts      respjson.Field
		QuotedTweet         respjson.Field
		QuotedTweetID       respjson.Field
		ReactionContext     respjson.Field
		RetweetedTweet      respjson.Field
		Scopes              respjson.Field
		Source              respjson.Field
		Tombstone           respjson.Field
		Type                respjson.Field
		UnmentionedUserIDs  respjson.Field
		URL                 respjson.Field
		ViewState           respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TweetDetail) RawJSON() string { return r.JSON.raw }
func (r *TweetDetail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Describes an X Article preview and its lifecycle metadata.
type TweetDetailArticle struct {
	ID string `json:"id"`
	// Public metadata whose fields are defined by X.
	CoverMedia    map[string]any `json:"coverMedia"`
	CoverMediaURL string         `json:"coverMediaUrl"`
	// Public metadata whose fields are defined by X.
	LifecycleState map[string]any `json:"lifecycleState"`
	// Public metadata whose fields are defined by X.
	Metadata    map[string]any `json:"metadata"`
	PreviewText string         `json:"previewText"`
	Title       string         `json:"title"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		CoverMedia     respjson.Field
		CoverMediaURL  respjson.Field
		LifecycleState respjson.Field
		Metadata       respjson.Field
		PreviewText    respjson.Field
		Title          respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TweetDetailArticle) RawJSON() string { return r.JSON.raw }
func (r *TweetDetailArticle) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Describes a public card and its referenced profiles.
type TweetDetailCard struct {
	ID string `json:"id"`
	// Public metadata whose fields are defined by X.
	BindingValues map[string]any `json:"bindingValues"`
	Name          string         `json:"name"`
	// Public metadata whose fields are defined by X.
	Platform map[string]any `json:"platform"`
	URL      string         `json:"url"`
	// Unresolved card user references.
	UserReferenceErrors []TweetDetailCardUserReferenceError `json:"userReferenceErrors"`
	UserReferences      []shared.UserProfile                `json:"userReferences"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                  respjson.Field
		BindingValues       respjson.Field
		Name                respjson.Field
		Platform            respjson.Field
		URL                 respjson.Field
		UserReferenceErrors respjson.Field
		UserReferences      respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TweetDetailCard) RawJSON() string { return r.JSON.raw }
func (r *TweetDetailCard) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TweetDetailCardUserReferenceError struct {
	Message string `json:"message"`
	Reason  string `json:"reason"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		Reason      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TweetDetailCardUserReferenceError) RawJSON() string { return r.JSON.raw }
func (r *TweetDetailCardUserReferenceError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Community Note presentation metadata returned by X.
type TweetDetailCommunityNote struct {
	ID             string `json:"id"`
	DestinationURL string `json:"destinationUrl"`
	Footer         string `json:"footer"`
	FooterIconType string `json:"footerIconType"`
	IconType       string `json:"iconType"`
	// Public metadata whose fields are defined by X.
	Metadata    map[string]any `json:"metadata"`
	ShortTitle  string         `json:"shortTitle"`
	Subtitle    string         `json:"subtitle"`
	Title       string         `json:"title"`
	VisualStyle string         `json:"visualStyle"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		DestinationURL respjson.Field
		Footer         respjson.Field
		FooterIconType respjson.Field
		IconType       respjson.Field
		Metadata       respjson.Field
		ShortTitle     respjson.Field
		Subtitle       respjson.Field
		Title          respjson.Field
		VisualStyle    respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TweetDetailCommunityNote) RawJSON() string { return r.JSON.raw }
func (r *TweetDetailCommunityNote) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Public reply policy and conversation owner.
type TweetDetailConversationControl struct {
	InviteViaMention bool   `json:"inviteViaMention"`
	OwnerUsername    string `json:"ownerUsername"`
	Policy           string `json:"policy"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		InviteViaMention respjson.Field
		OwnerUsername    respjson.Field
		Policy           respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TweetDetailConversationControl) RawJSON() string { return r.JSON.raw }
func (r *TweetDetailConversationControl) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Lists edit-chain identifiers and the remaining edit window.
type TweetDetailEdit struct {
	EditableUntilMsecs string   `json:"editableUntilMsecs"`
	EditTweetIDs       []string `json:"editTweetIds"`
	InitialTweetID     string   `json:"initialTweetId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EditableUntilMsecs respjson.Field
		EditTweetIDs       respjson.Field
		InitialTweetID     respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TweetDetailEdit) RawJSON() string { return r.JSON.raw }
func (r *TweetDetailEdit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Lists hashtags, symbols, links, and mentions from tweet text.
type TweetDetailEntities struct {
	Hashtags     []TweetDetailEntitiesHashtag     `json:"hashtags"`
	Smarttags    []TweetDetailEntitiesSmarttag    `json:"smarttags"`
	Symbols      []TweetDetailEntitiesSymbol      `json:"symbols"`
	Timestamps   []TweetDetailEntitiesTimestamp   `json:"timestamps"`
	URLs         []TweetDetailEntitiesURL         `json:"urls"`
	UserMentions []TweetDetailEntitiesUserMention `json:"user_mentions"`
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
func (r TweetDetailEntities) RawJSON() string { return r.JSON.raw }
func (r *TweetDetailEntities) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Provides hashtag text and source offsets within a tweet.
type TweetDetailEntitiesHashtag struct {
	Text    string  `json:"text" api:"required"`
	Indices []int64 `json:"indices"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Text        respjson.Field
		Indices     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TweetDetailEntitiesHashtag) RawJSON() string { return r.JSON.raw }
func (r *TweetDetailEntitiesHashtag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indexed smart-tag, cashtag, or video timestamp metadata.
type TweetDetailEntitiesSmarttag struct {
	Indices []int64                        `json:"indices"`
	Seconds float64                        `json:"seconds"`
	Tag     TweetDetailEntitiesSmarttagTag `json:"tag"`
	Text    string                         `json:"text"`
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
func (r TweetDetailEntitiesSmarttag) RawJSON() string { return r.JSON.raw }
func (r *TweetDetailEntitiesSmarttag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TweetDetailEntitiesSmarttagTag struct {
	Info TweetDetailEntitiesSmarttagTagInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TweetDetailEntitiesSmarttagTag) RawJSON() string { return r.JSON.raw }
func (r *TweetDetailEntitiesSmarttagTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TweetDetailEntitiesSmarttagTagInfo struct {
	Info TweetDetailEntitiesSmarttagTagInfoInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TweetDetailEntitiesSmarttagTagInfo) RawJSON() string { return r.JSON.raw }
func (r *TweetDetailEntitiesSmarttagTagInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TweetDetailEntitiesSmarttagTagInfoInfo struct {
	Name   string `json:"name"`
	Ticker string `json:"ticker"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Ticker      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TweetDetailEntitiesSmarttagTagInfoInfo) RawJSON() string { return r.JSON.raw }
func (r *TweetDetailEntitiesSmarttagTagInfoInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indexed smart-tag, cashtag, or video timestamp metadata.
type TweetDetailEntitiesSymbol struct {
	Indices []int64                      `json:"indices"`
	Seconds float64                      `json:"seconds"`
	Tag     TweetDetailEntitiesSymbolTag `json:"tag"`
	Text    string                       `json:"text"`
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
func (r TweetDetailEntitiesSymbol) RawJSON() string { return r.JSON.raw }
func (r *TweetDetailEntitiesSymbol) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TweetDetailEntitiesSymbolTag struct {
	Info TweetDetailEntitiesSymbolTagInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TweetDetailEntitiesSymbolTag) RawJSON() string { return r.JSON.raw }
func (r *TweetDetailEntitiesSymbolTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TweetDetailEntitiesSymbolTagInfo struct {
	Info TweetDetailEntitiesSymbolTagInfoInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TweetDetailEntitiesSymbolTagInfo) RawJSON() string { return r.JSON.raw }
func (r *TweetDetailEntitiesSymbolTagInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TweetDetailEntitiesSymbolTagInfoInfo struct {
	Name   string `json:"name"`
	Ticker string `json:"ticker"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Ticker      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TweetDetailEntitiesSymbolTagInfoInfo) RawJSON() string { return r.JSON.raw }
func (r *TweetDetailEntitiesSymbolTagInfoInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indexed smart-tag, cashtag, or video timestamp metadata.
type TweetDetailEntitiesTimestamp struct {
	Indices []int64                         `json:"indices"`
	Seconds float64                         `json:"seconds"`
	Tag     TweetDetailEntitiesTimestampTag `json:"tag"`
	Text    string                          `json:"text"`
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
func (r TweetDetailEntitiesTimestamp) RawJSON() string { return r.JSON.raw }
func (r *TweetDetailEntitiesTimestamp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TweetDetailEntitiesTimestampTag struct {
	Info TweetDetailEntitiesTimestampTagInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TweetDetailEntitiesTimestampTag) RawJSON() string { return r.JSON.raw }
func (r *TweetDetailEntitiesTimestampTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TweetDetailEntitiesTimestampTagInfo struct {
	Info TweetDetailEntitiesTimestampTagInfoInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TweetDetailEntitiesTimestampTagInfo) RawJSON() string { return r.JSON.raw }
func (r *TweetDetailEntitiesTimestampTagInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TweetDetailEntitiesTimestampTagInfoInfo struct {
	Name   string `json:"name"`
	Ticker string `json:"ticker"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Ticker      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TweetDetailEntitiesTimestampTagInfoInfo) RawJSON() string { return r.JSON.raw }
func (r *TweetDetailEntitiesTimestampTagInfoInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Provides shortened, display, and expanded URLs from tweet text.
type TweetDetailEntitiesURL struct {
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
func (r TweetDetailEntitiesURL) RawJSON() string { return r.JSON.raw }
func (r *TweetDetailEntitiesURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Provides profile identity and source offsets for a mention.
type TweetDetailEntitiesUserMention struct {
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
func (r TweetDetailEntitiesUserMention) RawJSON() string { return r.JSON.raw }
func (r *TweetDetailEntitiesUserMention) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TweetDetailLimitedAction struct {
	Action string                         `json:"action"`
	Prompt TweetDetailLimitedActionPrompt `json:"prompt"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Action      respjson.Field
		Prompt      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TweetDetailLimitedAction) RawJSON() string { return r.JSON.raw }
func (r *TweetDetailLimitedAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TweetDetailLimitedActionPrompt struct {
	CtaType  string `json:"ctaType"`
	Headline string `json:"headline"`
	// Public metadata whose fields are defined by X.
	Metadata map[string]any `json:"metadata"`
	Subtext  string         `json:"subtext"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CtaType     respjson.Field
		Headline    respjson.Field
		Metadata    respjson.Field
		Subtext     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TweetDetailLimitedActionPrompt) RawJSON() string { return r.JSON.raw }
func (r *TweetDetailLimitedActionPrompt) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Complete Note Tweet content and rich-text metadata.
type TweetDetailNoteTweet struct {
	Text string `json:"text" api:"required"`
	ID   string `json:"id"`
	// Lists hashtags, symbols, links, and mentions from tweet text.
	Entities TweetDetailNoteTweetEntities `json:"entities"`
	// Inline media positions in the Note Tweet text.
	InlineMedia  []TweetDetailNoteTweetInlineMedia `json:"inlineMedia"`
	IsExpandable bool                              `json:"isExpandable"`
	RichtextTags []TweetDetailNoteTweetRichtextTag `json:"richtextTags"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Text         respjson.Field
		ID           respjson.Field
		Entities     respjson.Field
		InlineMedia  respjson.Field
		IsExpandable respjson.Field
		RichtextTags respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TweetDetailNoteTweet) RawJSON() string { return r.JSON.raw }
func (r *TweetDetailNoteTweet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Lists hashtags, symbols, links, and mentions from tweet text.
type TweetDetailNoteTweetEntities struct {
	Hashtags     []TweetDetailNoteTweetEntitiesHashtag     `json:"hashtags"`
	Smarttags    []TweetDetailNoteTweetEntitiesSmarttag    `json:"smarttags"`
	Symbols      []TweetDetailNoteTweetEntitiesSymbol      `json:"symbols"`
	Timestamps   []TweetDetailNoteTweetEntitiesTimestamp   `json:"timestamps"`
	URLs         []TweetDetailNoteTweetEntitiesURL         `json:"urls"`
	UserMentions []TweetDetailNoteTweetEntitiesUserMention `json:"user_mentions"`
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
func (r TweetDetailNoteTweetEntities) RawJSON() string { return r.JSON.raw }
func (r *TweetDetailNoteTweetEntities) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Provides hashtag text and source offsets within a tweet.
type TweetDetailNoteTweetEntitiesHashtag struct {
	Text    string  `json:"text" api:"required"`
	Indices []int64 `json:"indices"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Text        respjson.Field
		Indices     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TweetDetailNoteTweetEntitiesHashtag) RawJSON() string { return r.JSON.raw }
func (r *TweetDetailNoteTweetEntitiesHashtag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indexed smart-tag, cashtag, or video timestamp metadata.
type TweetDetailNoteTweetEntitiesSmarttag struct {
	Indices []int64                                 `json:"indices"`
	Seconds float64                                 `json:"seconds"`
	Tag     TweetDetailNoteTweetEntitiesSmarttagTag `json:"tag"`
	Text    string                                  `json:"text"`
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
func (r TweetDetailNoteTweetEntitiesSmarttag) RawJSON() string { return r.JSON.raw }
func (r *TweetDetailNoteTweetEntitiesSmarttag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TweetDetailNoteTweetEntitiesSmarttagTag struct {
	Info TweetDetailNoteTweetEntitiesSmarttagTagInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TweetDetailNoteTweetEntitiesSmarttagTag) RawJSON() string { return r.JSON.raw }
func (r *TweetDetailNoteTweetEntitiesSmarttagTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TweetDetailNoteTweetEntitiesSmarttagTagInfo struct {
	Info TweetDetailNoteTweetEntitiesSmarttagTagInfoInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TweetDetailNoteTweetEntitiesSmarttagTagInfo) RawJSON() string { return r.JSON.raw }
func (r *TweetDetailNoteTweetEntitiesSmarttagTagInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TweetDetailNoteTweetEntitiesSmarttagTagInfoInfo struct {
	Name   string `json:"name"`
	Ticker string `json:"ticker"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Ticker      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TweetDetailNoteTweetEntitiesSmarttagTagInfoInfo) RawJSON() string { return r.JSON.raw }
func (r *TweetDetailNoteTweetEntitiesSmarttagTagInfoInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indexed smart-tag, cashtag, or video timestamp metadata.
type TweetDetailNoteTweetEntitiesSymbol struct {
	Indices []int64                               `json:"indices"`
	Seconds float64                               `json:"seconds"`
	Tag     TweetDetailNoteTweetEntitiesSymbolTag `json:"tag"`
	Text    string                                `json:"text"`
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
func (r TweetDetailNoteTweetEntitiesSymbol) RawJSON() string { return r.JSON.raw }
func (r *TweetDetailNoteTweetEntitiesSymbol) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TweetDetailNoteTweetEntitiesSymbolTag struct {
	Info TweetDetailNoteTweetEntitiesSymbolTagInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TweetDetailNoteTweetEntitiesSymbolTag) RawJSON() string { return r.JSON.raw }
func (r *TweetDetailNoteTweetEntitiesSymbolTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TweetDetailNoteTweetEntitiesSymbolTagInfo struct {
	Info TweetDetailNoteTweetEntitiesSymbolTagInfoInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TweetDetailNoteTweetEntitiesSymbolTagInfo) RawJSON() string { return r.JSON.raw }
func (r *TweetDetailNoteTweetEntitiesSymbolTagInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TweetDetailNoteTweetEntitiesSymbolTagInfoInfo struct {
	Name   string `json:"name"`
	Ticker string `json:"ticker"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Ticker      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TweetDetailNoteTweetEntitiesSymbolTagInfoInfo) RawJSON() string { return r.JSON.raw }
func (r *TweetDetailNoteTweetEntitiesSymbolTagInfoInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indexed smart-tag, cashtag, or video timestamp metadata.
type TweetDetailNoteTweetEntitiesTimestamp struct {
	Indices []int64                                  `json:"indices"`
	Seconds float64                                  `json:"seconds"`
	Tag     TweetDetailNoteTweetEntitiesTimestampTag `json:"tag"`
	Text    string                                   `json:"text"`
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
func (r TweetDetailNoteTweetEntitiesTimestamp) RawJSON() string { return r.JSON.raw }
func (r *TweetDetailNoteTweetEntitiesTimestamp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TweetDetailNoteTweetEntitiesTimestampTag struct {
	Info TweetDetailNoteTweetEntitiesTimestampTagInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TweetDetailNoteTweetEntitiesTimestampTag) RawJSON() string { return r.JSON.raw }
func (r *TweetDetailNoteTweetEntitiesTimestampTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TweetDetailNoteTweetEntitiesTimestampTagInfo struct {
	Info TweetDetailNoteTweetEntitiesTimestampTagInfoInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TweetDetailNoteTweetEntitiesTimestampTagInfo) RawJSON() string { return r.JSON.raw }
func (r *TweetDetailNoteTweetEntitiesTimestampTagInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TweetDetailNoteTweetEntitiesTimestampTagInfoInfo struct {
	Name   string `json:"name"`
	Ticker string `json:"ticker"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Ticker      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TweetDetailNoteTweetEntitiesTimestampTagInfoInfo) RawJSON() string { return r.JSON.raw }
func (r *TweetDetailNoteTweetEntitiesTimestampTagInfoInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Provides shortened, display, and expanded URLs from tweet text.
type TweetDetailNoteTweetEntitiesURL struct {
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
func (r TweetDetailNoteTweetEntitiesURL) RawJSON() string { return r.JSON.raw }
func (r *TweetDetailNoteTweetEntitiesURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Provides profile identity and source offsets for a mention.
type TweetDetailNoteTweetEntitiesUserMention struct {
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
func (r TweetDetailNoteTweetEntitiesUserMention) RawJSON() string { return r.JSON.raw }
func (r *TweetDetailNoteTweetEntitiesUserMention) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TweetDetailNoteTweetInlineMedia struct {
	Index   int64  `json:"index" api:"required"`
	MediaID string `json:"mediaId" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Index       respjson.Field
		MediaID     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TweetDetailNoteTweetInlineMedia) RawJSON() string { return r.JSON.raw }
func (r *TweetDetailNoteTweetInlineMedia) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TweetDetailNoteTweetRichtextTag struct {
	FromIndex int64    `json:"fromIndex" api:"required"`
	ToIndex   int64    `json:"toIndex" api:"required"`
	Types     []string `json:"types" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FromIndex   respjson.Field
		ToIndex     respjson.Field
		Types       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TweetDetailNoteTweetRichtextTag) RawJSON() string { return r.JSON.raw }
func (r *TweetDetailNoteTweetRichtextTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Describes public place metadata on a geotagged tweet.
type TweetDetailPlace struct {
	ID string `json:"id"`
	// Public metadata whose fields are defined by X.
	BoundingBox map[string]any `json:"boundingBox"`
	Country     string         `json:"country"`
	CountryCode string         `json:"countryCode"`
	FullName    string         `json:"fullName"`
	Name        string         `json:"name"`
	PlaceType   string         `json:"placeType"`
	URL         string         `json:"url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		BoundingBox respjson.Field
		Country     respjson.Field
		CountryCode respjson.Field
		FullName    respjson.Field
		Name        respjson.Field
		PlaceType   respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TweetDetailPlace) RawJSON() string { return r.JSON.raw }
func (r *TweetDetailPlace) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Engagement counts retained from a prior tweet edit.
type TweetDetailPreviousCounts struct {
	BookmarkCount int64 `json:"bookmarkCount"`
	LikeCount     int64 `json:"likeCount"`
	QuoteCount    int64 `json:"quoteCount"`
	ReplyCount    int64 `json:"replyCount"`
	RetweetCount  int64 `json:"retweetCount"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BookmarkCount respjson.Field
		LikeCount     respjson.Field
		QuoteCount    respjson.Field
		ReplyCount    respjson.Field
		RetweetCount  respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TweetDetailPreviousCounts) RawJSON() string { return r.JSON.raw }
func (r *TweetDetailPreviousCounts) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Public post and user referenced by this reaction.
type TweetDetailReactionContext struct {
	// Referenced post ID.
	ReactedToPostID string `json:"reactedToPostId"`
	// X user profile with bio, follower counts, and verification status.
	ReactedToUser shared.UserProfile `json:"reactedToUser"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ReactedToPostID respjson.Field
		ReactedToUser   respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TweetDetailReactionContext) RawJSON() string { return r.JSON.raw }
func (r *TweetDetailReactionContext) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Public visibility notice attached to an available tweet.
type TweetDetailTombstone struct {
	Text TweetDetailTombstoneText `json:"text"`
	// Visibility notice type.
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Text        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TweetDetailTombstone) RawJSON() string { return r.JSON.raw }
func (r *TweetDetailTombstone) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TweetDetailTombstoneText struct {
	Entities []TweetDetailTombstoneTextEntity `json:"entities"`
	// Right-to-left text direction.
	Rtl bool `json:"rtl"`
	// Human-readable notice text.
	Text string `json:"text"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Entities    respjson.Field
		Rtl         respjson.Field
		Text        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TweetDetailTombstoneText) RawJSON() string { return r.JSON.raw }
func (r *TweetDetailTombstoneText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TweetDetailTombstoneTextEntity struct {
	FromIndex int64                             `json:"fromIndex"`
	Ref       TweetDetailTombstoneTextEntityRef `json:"ref"`
	ToIndex   int64                             `json:"toIndex"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FromIndex   respjson.Field
		Ref         respjson.Field
		ToIndex     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TweetDetailTombstoneTextEntity) RawJSON() string { return r.JSON.raw }
func (r *TweetDetailTombstoneTextEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TweetDetailTombstoneTextEntityRef struct {
	Type    string `json:"type"`
	URL     string `json:"url"`
	URLType string `json:"urlType"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		URL         respjson.Field
		URLType     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TweetDetailTombstoneTextEntityRef) RawJSON() string { return r.JSON.raw }
func (r *TweetDetailTombstoneTextEntityRef) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Durable write record. Poll statusUrl until terminal is true. Reusing its
// Idempotency-Key returns this record. Create another action only when safeToRetry
// is true.
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

// Durable write record. Poll statusUrl until terminal is true. Reusing its
// Idempotency-Key returns this record. Create another action only when safeToRetry
// is true.
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

// Direct reply rows. No-mode requests use resumable automatic coverage. Complete
// mode also returns nested replies and coverage diagnostics. Keep nested replies
// separate from direct coverage.
type XTweetGetRepliesResponse struct {
	// Evidence for direct-reply coverage and collector behavior.
	Diagnostic XTweetGetRepliesResponseDiagnostic `json:"diagnostic"`
	// Nested replies. Excluded from direct coverage.
	NestedReplies []shared.SearchTweet `json:"nested_replies"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Diagnostic    respjson.Field
		NestedReplies respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
	shared.PaginatedTweets
}

// Returns the unmodified JSON received from the API
func (r XTweetGetRepliesResponse) RawJSON() string { return r.JSON.raw }
func (r *XTweetGetRepliesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Evidence for direct-reply coverage and collector behavior.
type XTweetGetRepliesResponseDiagnostic struct {
	// Whether coverage met the target without truncation.
	Complete bool `json:"complete" api:"required"`
	// Unique direct replies as a percentage of the reported count.
	CoveragePercentage float64 `json:"coveragePercentage" api:"required"`
	// Cursor requests that failed.
	CursorFailures int64 `json:"cursorFailures" api:"required"`
	// Duplicate tweet IDs removed across pages and strategies.
	DuplicateCount int64 `json:"duplicateCount" api:"required"`
	// Empty pages rejected because they did not make progress.
	EmptyFalseProgressPages int64 `json:"emptyFalseProgressPages" api:"required"`
	// Malformed response items rejected.
	MalformedCount int64 `json:"malformedCount" api:"required"`
	// Expected response modules or fields missing from X.
	MissingResponseModulesOrFields []string `json:"missingResponseModulesOrFields" api:"required"`
	// Unique nested replies kept outside direct coverage.
	NestedReplyCount int64 `json:"nestedReplyCount" api:"required"`
	// Total pages attempted across all strategies.
	PagesAttempted int64 `json:"pagesAttempted" api:"required"`
	// Recommended next action when coverage is incomplete.
	RecommendedFallback string `json:"recommendedFallback" api:"required"`
	// Repeated cursors rejected to prevent loops.
	RepeatedCursorCount int64 `json:"repeatedCursorCount" api:"required"`
	// Reply count reported on the source post.
	ReportedReplyCount int64 `json:"reportedReplyCount" api:"required"`
	// Whether the requested row limit truncated safe results.
	ResponseTruncated bool `json:"responseTruncated" api:"required"`
	// Field-presence counts across the collected direct replies.
	Richness XTweetGetRepliesResponseDiagnosticRichness `json:"richness" api:"required"`
	// Per-strategy pagination and contribution evidence.
	StrategiesAttempted []XTweetGetRepliesResponseDiagnosticStrategiesAttempted `json:"strategiesAttempted" api:"required"`
	// Minimum direct replies required for the coverage target.
	TargetDirectReplies int64 `json:"targetDirectReplies" api:"required"`
	// Unique replies whose parent ID equals the source post ID.
	UniqueDirectReplies int64 `json:"uniqueDirectReplies" api:"required"`
	// Tweets rejected because they belonged elsewhere.
	UnrelatedCount int64 `json:"unrelatedCount" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Complete                       respjson.Field
		CoveragePercentage             respjson.Field
		CursorFailures                 respjson.Field
		DuplicateCount                 respjson.Field
		EmptyFalseProgressPages        respjson.Field
		MalformedCount                 respjson.Field
		MissingResponseModulesOrFields respjson.Field
		NestedReplyCount               respjson.Field
		PagesAttempted                 respjson.Field
		RecommendedFallback            respjson.Field
		RepeatedCursorCount            respjson.Field
		ReportedReplyCount             respjson.Field
		ResponseTruncated              respjson.Field
		Richness                       respjson.Field
		StrategiesAttempted            respjson.Field
		TargetDirectReplies            respjson.Field
		UniqueDirectReplies            respjson.Field
		UnrelatedCount                 respjson.Field
		ExtraFields                    map[string]respjson.Field
		raw                            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XTweetGetRepliesResponseDiagnostic) RawJSON() string { return r.JSON.raw }
func (r *XTweetGetRepliesResponseDiagnostic) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Field-presence counts across the collected direct replies.
type XTweetGetRepliesResponseDiagnosticRichness struct {
	// Replies with article content.
	Article int64 `json:"article" api:"required"`
	// Replies with author details.
	Author int64 `json:"author" api:"required"`
	// Replies with card metadata.
	Card int64 `json:"card" api:"required"`
	// Replies with community-note data.
	CommunityNote int64 `json:"communityNote" api:"required"`
	// Replies with a creation timestamp.
	CreatedAt int64 `json:"createdAt" api:"required"`
	// Replies with engagement counts.
	EngagementCounts int64 `json:"engagementCounts" api:"required"`
	// Replies with entity metadata.
	Entities int64 `json:"entities" api:"required"`
	// Replies with a language value.
	Language int64 `json:"language" api:"required"`
	// Replies with media metadata.
	Media int64 `json:"media" api:"required"`
	// Replies with quoted or reposted tweet data.
	QuotedOrRepostedTweet int64 `json:"quotedOrRepostedTweet" api:"required"`
	// Replies with text.
	Text int64 `json:"text" api:"required"`
	// Total unique direct replies evaluated for richness.
	TotalReplies int64 `json:"totalReplies" api:"required"`
	// Replies with a canonical URL.
	URL int64 `json:"url" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Article               respjson.Field
		Author                respjson.Field
		Card                  respjson.Field
		CommunityNote         respjson.Field
		CreatedAt             respjson.Field
		EngagementCounts      respjson.Field
		Entities              respjson.Field
		Language              respjson.Field
		Media                 respjson.Field
		QuotedOrRepostedTweet respjson.Field
		Text                  respjson.Field
		TotalReplies          respjson.Field
		URL                   respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XTweetGetRepliesResponseDiagnosticRichness) RawJSON() string { return r.JSON.raw }
func (r *XTweetGetRepliesResponseDiagnosticRichness) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Result counts and stop reason for one reply strategy.
type XTweetGetRepliesResponseDiagnosticStrategiesAttempted struct {
	Name             string `json:"name" api:"required"`
	NewDirectReplies int64  `json:"newDirectReplies" api:"required"`
	NewNestedReplies int64  `json:"newNestedReplies" api:"required"`
	PagesAttempted   int64  `json:"pagesAttempted" api:"required"`
	// Any of "deadline", "empty_pages", "error", "missing_cursor", "no_next_page",
	// "page_cap", "repeated_cursor".
	StopReason string `json:"stopReason" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name             respjson.Field
		NewDirectReplies respjson.Field
		NewNestedReplies respjson.Field
		PagesAttempted   respjson.Field
		StopReason       respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XTweetGetRepliesResponseDiagnosticStrategiesAttempted) RawJSON() string { return r.JSON.raw }
func (r *XTweetGetRepliesResponseDiagnosticStrategiesAttempted) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// XTweetSearchResponseUnion contains all possible properties and values from
// [shared.PaginatedTweets], [XTweetSearchResponseTweetSearchCoverageResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type XTweetSearchResponseUnion struct {
	HasNextPage bool   `json:"has_next_page"`
	NextCursor  string `json:"next_cursor"`
	// This field is from variant [shared.PaginatedTweets],
	// [XTweetSearchResponseTweetSearchCoverageResponse].
	Tweets []shared.SearchTweet `json:"tweets"`
	// This field is from variant [shared.PaginatedTweets],
	// [XTweetSearchResponseTweetSearchCoverageResponse].
	FilteredCount int64 `json:"filtered_count"`
	// This field is from variant [XTweetSearchResponseTweetSearchCoverageResponse].
	Diagnostic XTweetSearchResponseTweetSearchCoverageResponseDiagnostic `json:"diagnostic"`
	JSON       struct {
		HasNextPage   respjson.Field
		NextCursor    respjson.Field
		Tweets        respjson.Field
		FilteredCount respjson.Field
		Diagnostic    respjson.Field
		raw           string
	} `json:"-"`
}

func (u XTweetSearchResponseUnion) AsPaginatedTweets() (v shared.PaginatedTweets) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u XTweetSearchResponseUnion) AsXTweetSearchResponseTweetSearchCoverageResponse() (v XTweetSearchResponseTweetSearchCoverageResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u XTweetSearchResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *XTweetSearchResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Terminal Tweet search coverage response with diagnostics.
type XTweetSearchResponseTweetSearchCoverageResponse struct {
	// Coverage evidence across parallel search strategies.
	Diagnostic XTweetSearchResponseTweetSearchCoverageResponseDiagnostic `json:"diagnostic" api:"required"`
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
	shared.PaginatedTweets
}

// Returns the unmodified JSON received from the API
func (r XTweetSearchResponseTweetSearchCoverageResponse) RawJSON() string { return r.JSON.raw }
func (r *XTweetSearchResponseTweetSearchCoverageResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Coverage evidence across parallel search strategies.
type XTweetSearchResponseTweetSearchCoverageResponseDiagnostic struct {
	// True after all active strategies exhaust their sources.
	Complete            bool  `json:"complete" api:"required"`
	CursorFailureCount  int64 `json:"cursorFailureCount" api:"required"`
	DeadlineReached     bool  `json:"deadlineReached" api:"required"`
	DuplicateCount      int64 `json:"duplicateCount" api:"required"`
	FailedStrategyCount int64 `json:"failedStrategyCount" api:"required"`
	MalformedCount      int64 `json:"malformedCount" api:"required"`
	PagesFetched        int64 `json:"pagesFetched" api:"required"`
	// Whether bounded time windows ran in parallel.
	Partitioned bool `json:"partitioned" api:"required"`
	// True when credits or the requested limit reduce output.
	ResponseTruncated    bool                                                                `json:"responseTruncated" api:"required"`
	ResultLimitReached   bool                                                                `json:"resultLimitReached" api:"required"`
	ReturnedTweets       int64                                                               `json:"returnedTweets" api:"required"`
	StalledStrategyCount int64                                                               `json:"stalledStrategyCount" api:"required"`
	Strategies           []XTweetSearchResponseTweetSearchCoverageResponseDiagnosticStrategy `json:"strategies" api:"required"`
	StrategyCount        int64                                                               `json:"strategyCount" api:"required"`
	UniqueTweets         int64                                                               `json:"uniqueTweets" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Complete             respjson.Field
		CursorFailureCount   respjson.Field
		DeadlineReached      respjson.Field
		DuplicateCount       respjson.Field
		FailedStrategyCount  respjson.Field
		MalformedCount       respjson.Field
		PagesFetched         respjson.Field
		Partitioned          respjson.Field
		ResponseTruncated    respjson.Field
		ResultLimitReached   respjson.Field
		ReturnedTweets       respjson.Field
		StalledStrategyCount respjson.Field
		Strategies           respjson.Field
		StrategyCount        respjson.Field
		UniqueTweets         respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XTweetSearchResponseTweetSearchCoverageResponseDiagnostic) RawJSON() string {
	return r.JSON.raw
}
func (r *XTweetSearchResponseTweetSearchCoverageResponseDiagnostic) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Result counts and stop reason for one Tweet search strategy.
type XTweetSearchResponseTweetSearchCoverageResponseDiagnosticStrategy struct {
	DuplicateCount int64 `json:"duplicateCount" api:"required"`
	PagesFetched   int64 `json:"pagesFetched" api:"required"`
	// Any of "Latest", "Top".
	QueryType string `json:"queryType" api:"required"`
	// Reason a coverage strategy stopped.
	//
	// Any of "cursor_failure", "deadline", "exhausted", "failed", "page_limit",
	// "result_limit", "stalled".
	StopReason  string `json:"stopReason" api:"required"`
	Strategy    int64  `json:"strategy" api:"required"`
	UniqueAdded int64  `json:"uniqueAdded" api:"required"`
	// Non-overlapping time partition used by one strategy.
	Window XTweetSearchResponseTweetSearchCoverageResponseDiagnosticStrategyWindow `json:"window"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DuplicateCount respjson.Field
		PagesFetched   respjson.Field
		QueryType      respjson.Field
		StopReason     respjson.Field
		Strategy       respjson.Field
		UniqueAdded    respjson.Field
		Window         respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XTweetSearchResponseTweetSearchCoverageResponseDiagnosticStrategy) RawJSON() string {
	return r.JSON.raw
}
func (r *XTweetSearchResponseTweetSearchCoverageResponseDiagnosticStrategy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Non-overlapping time partition used by one strategy.
type XTweetSearchResponseTweetSearchCoverageResponseDiagnosticStrategyWindow struct {
	SinceTime time.Time `json:"sinceTime" api:"required" format:"date-time"`
	UntilTime time.Time `json:"untilTime" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		SinceTime   respjson.Field
		UntilTime   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XTweetSearchResponseTweetSearchCoverageResponseDiagnosticStrategyWindow) RawJSON() string {
	return r.JSON.raw
}
func (r *XTweetSearchResponseTweetSearchCoverageResponseDiagnosticStrategyWindow) UnmarshalJSON(data []byte) error {
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
	// Match any comma-separated or line-separated bio term, ignoring case.
	BioContains param.Opt[string] `query:"bioContains,omitzero" json:"-"`
	// Pagination cursor for favoriters
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
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
	// Only return tweets from Blue-verified authors.
	BlueVerifiedOnly param.Opt[bool] `query:"blueVerifiedOnly,omitzero" json:"-"`
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
	// Filter by author username.
	FromUser param.Opt[string] `query:"fromUser,omitzero" json:"-"`
	// Match latitude, longitude, and radius.
	Geocode param.Opt[string] `query:"geocode,omitzero" json:"-"`
	// Hashtags separated by spaces, commas, or lines.
	Hashtags param.Opt[string] `query:"hashtags,omitzero" json:"-"`
	// Include reply tweets unless replies specifies another mode.
	IncludeReplies param.Opt[bool] `query:"includeReplies,omitzero" json:"-"`
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
	// Automatic pages accept 1-300 Tweets. Standard pages keep 1-100. Default 20.
	// Follow next_cursor while the response reports more pages. Deprecated aliases
	// remain accepted.
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
	MediaType XTweetGetQuotesParamsMediaType `query:"mediaType,omitzero" json:"-"`
	// Optional legacy pagination override.
	//
	// Any of "standard".
	Mode XTweetGetQuotesParamsMode `query:"mode,omitzero" json:"-"`
	// Only when the caller requests a quote mode.
	//
	// Any of "include", "exclude", "only".
	Quotes XTweetGetQuotesParamsQuotes `query:"quotes,omitzero" json:"-"`
	// Only when the caller requests a reply mode.
	//
	// Any of "include", "exclude", "only".
	Replies XTweetGetQuotesParamsReplies `query:"replies,omitzero" json:"-"`
	// Only when the caller requests a repost mode.
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

// Filter media. Aliases: has_video, has_media.
type XTweetGetQuotesParamsMediaType string

const (
	XTweetGetQuotesParamsMediaTypeImages XTweetGetQuotesParamsMediaType = "images"
	XTweetGetQuotesParamsMediaTypeVideos XTweetGetQuotesParamsMediaType = "videos"
	XTweetGetQuotesParamsMediaTypeGifs   XTweetGetQuotesParamsMediaType = "gifs"
	XTweetGetQuotesParamsMediaTypeMedia  XTweetGetQuotesParamsMediaType = "media"
	XTweetGetQuotesParamsMediaTypeLinks  XTweetGetQuotesParamsMediaType = "links"
	XTweetGetQuotesParamsMediaTypeNone   XTweetGetQuotesParamsMediaType = "none"
)

// Optional legacy pagination override.
type XTweetGetQuotesParamsMode string

const (
	XTweetGetQuotesParamsModeStandard XTweetGetQuotesParamsMode = "standard"
)

// Only when the caller requests a quote mode.
type XTweetGetQuotesParamsQuotes string

const (
	XTweetGetQuotesParamsQuotesInclude XTweetGetQuotesParamsQuotes = "include"
	XTweetGetQuotesParamsQuotesExclude XTweetGetQuotesParamsQuotes = "exclude"
	XTweetGetQuotesParamsQuotesOnly    XTweetGetQuotesParamsQuotes = "only"
)

// Only when the caller requests a reply mode.
type XTweetGetQuotesParamsReplies string

const (
	XTweetGetQuotesParamsRepliesInclude XTweetGetQuotesParamsReplies = "include"
	XTweetGetQuotesParamsRepliesExclude XTweetGetQuotesParamsReplies = "exclude"
	XTweetGetQuotesParamsRepliesOnly    XTweetGetQuotesParamsReplies = "only"
)

// Only when the caller requests a repost mode.
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
	// Only return tweets from Blue-verified authors.
	BlueVerifiedOnly param.Opt[bool] `query:"blueVerifiedOnly,omitzero" json:"-"`
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
	// Exclude replies written by the source-post author.
	ExcludeOriginalAuthor param.Opt[bool] `query:"excludeOriginalAuthor,omitzero" json:"-"`
	// Exclude a source application.
	ExcludeSource param.Opt[string] `query:"excludeSource,omitzero" json:"-"`
	// Words or quoted phrases to exclude. Separate with spaces, commas, or lines.
	ExcludeWords param.Opt[string] `query:"excludeWords,omitzero" json:"-"`
	// Filter by author username.
	FromUser param.Opt[string] `query:"fromUser,omitzero" json:"-"`
	// Match latitude, longitude, and radius.
	Geocode param.Opt[string] `query:"geocode,omitzero" json:"-"`
	// Hashtags separated by spaces, commas, or lines.
	Hashtags param.Opt[string] `query:"hashtags,omitzero" json:"-"`
	// Only return replies containing media.
	HasMediaOnly param.Opt[bool] `query:"hasMediaOnly,omitzero" json:"-"`
	// Include the source post and count it toward limit.
	IncludeOriginalPost param.Opt[bool] `query:"includeOriginalPost,omitzero" json:"-"`
	// Only replies to this tweet ID.
	InReplyToTweetID param.Opt[string] `query:"inReplyToTweetId,omitzero" json:"-"`
	// Filter by language. Alias `lang` is accepted.
	Language param.Opt[string] `query:"language,omitzero" json:"-"`
	// Complete mode caps combined direct and nested replies at 25,000. Automatic pages
	// accept 1-300. Standard pages accept 1-100.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Maximum reply depth from the source post.
	MaxDepth param.Opt[int64] `query:"maxDepth,omitzero" json:"-"`
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
	// Automatic pages accept 1-300 Tweets. Standard pages keep 1-100. Default 20.
	// Follow next_cursor while the response reports more pages. Deprecated aliases
	// remain accepted.
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
	MediaType XTweetGetRepliesParamsMediaType `query:"mediaType,omitzero" json:"-"`
	// Override automatic coverage. Standard uses legacy pagination. Complete adds
	// nested replies, diagnostics, scope, depth, sorting, and original-post controls.
	//
	// Any of "standard", "complete".
	Mode XTweetGetRepliesParamsMode `query:"mode,omitzero" json:"-"`
	// Only when the caller requests a quote mode.
	//
	// Any of "include", "exclude", "only".
	Quotes XTweetGetRepliesParamsQuotes `query:"quotes,omitzero" json:"-"`
	// Only when the caller requests a reply mode.
	//
	// Any of "include", "exclude", "only".
	Replies XTweetGetRepliesParamsReplies `query:"replies,omitzero" json:"-"`
	// Only when the caller requests a repost mode.
	//
	// Any of "include", "exclude", "only".
	Retweets XTweetGetRepliesParamsRetweets `query:"retweets,omitzero" json:"-"`
	// Select all replies, direct replies, or nested replies.
	//
	// Any of "all", "direct", "nested".
	Scope XTweetGetRepliesParamsScope `query:"scope,omitzero" json:"-"`
	// Sort the selected replies before applying limit.
	//
	// Any of "relevance", "latest", "oldest", "likes".
	Sort XTweetGetRepliesParamsSort `query:"sort,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [XTweetGetRepliesParams]'s query parameters as `url.Values`.
func (r XTweetGetRepliesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter media. Aliases: has_video, has_media.
type XTweetGetRepliesParamsMediaType string

const (
	XTweetGetRepliesParamsMediaTypeImages XTweetGetRepliesParamsMediaType = "images"
	XTweetGetRepliesParamsMediaTypeVideos XTweetGetRepliesParamsMediaType = "videos"
	XTweetGetRepliesParamsMediaTypeGifs   XTweetGetRepliesParamsMediaType = "gifs"
	XTweetGetRepliesParamsMediaTypeMedia  XTweetGetRepliesParamsMediaType = "media"
	XTweetGetRepliesParamsMediaTypeLinks  XTweetGetRepliesParamsMediaType = "links"
	XTweetGetRepliesParamsMediaTypeNone   XTweetGetRepliesParamsMediaType = "none"
)

// Override automatic coverage. Standard uses legacy pagination. Complete adds
// nested replies, diagnostics, scope, depth, sorting, and original-post controls.
type XTweetGetRepliesParamsMode string

const (
	XTweetGetRepliesParamsModeStandard XTweetGetRepliesParamsMode = "standard"
	XTweetGetRepliesParamsModeComplete XTweetGetRepliesParamsMode = "complete"
)

// Only when the caller requests a quote mode.
type XTweetGetRepliesParamsQuotes string

const (
	XTweetGetRepliesParamsQuotesInclude XTweetGetRepliesParamsQuotes = "include"
	XTweetGetRepliesParamsQuotesExclude XTweetGetRepliesParamsQuotes = "exclude"
	XTweetGetRepliesParamsQuotesOnly    XTweetGetRepliesParamsQuotes = "only"
)

// Only when the caller requests a reply mode.
type XTweetGetRepliesParamsReplies string

const (
	XTweetGetRepliesParamsRepliesInclude XTweetGetRepliesParamsReplies = "include"
	XTweetGetRepliesParamsRepliesExclude XTweetGetRepliesParamsReplies = "exclude"
	XTweetGetRepliesParamsRepliesOnly    XTweetGetRepliesParamsReplies = "only"
)

// Only when the caller requests a repost mode.
type XTweetGetRepliesParamsRetweets string

const (
	XTweetGetRepliesParamsRetweetsInclude XTweetGetRepliesParamsRetweets = "include"
	XTweetGetRepliesParamsRetweetsExclude XTweetGetRepliesParamsRetweets = "exclude"
	XTweetGetRepliesParamsRetweetsOnly    XTweetGetRepliesParamsRetweets = "only"
)

// Select all replies, direct replies, or nested replies.
type XTweetGetRepliesParamsScope string

const (
	XTweetGetRepliesParamsScopeAll    XTweetGetRepliesParamsScope = "all"
	XTweetGetRepliesParamsScopeDirect XTweetGetRepliesParamsScope = "direct"
	XTweetGetRepliesParamsScopeNested XTweetGetRepliesParamsScope = "nested"
)

// Sort the selected replies before applying limit.
type XTweetGetRepliesParamsSort string

const (
	XTweetGetRepliesParamsSortRelevance XTweetGetRepliesParamsSort = "relevance"
	XTweetGetRepliesParamsSortLatest    XTweetGetRepliesParamsSort = "latest"
	XTweetGetRepliesParamsSortOldest    XTweetGetRepliesParamsSort = "oldest"
	XTweetGetRepliesParamsSortLikes     XTweetGetRepliesParamsSort = "likes"
)

type XTweetGetRetweetersParams struct {
	// Match any comma-separated or line-separated bio term, ignoring case.
	BioContains param.Opt[string] `query:"bioContains,omitzero" json:"-"`
	// Pagination cursor for retweeters
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
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

// URLQuery serializes [XTweetGetRetweetersParams]'s query parameters as
// `url.Values`.
func (r XTweetGetRetweetersParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type XTweetGetThreadParams struct {
	// Words or quoted phrases where any one can match. Separate with spaces, commas,
	// or lines.
	AnyWords param.Opt[string] `query:"anyWords,omitzero" json:"-"`
	// Only return tweets from Blue-verified authors.
	BlueVerifiedOnly param.Opt[bool] `query:"blueVerifiedOnly,omitzero" json:"-"`
	// Cashtags separated by spaces, commas, or lines.
	Cashtags param.Opt[string] `query:"cashtags,omitzero" json:"-"`
	// Conversation ID filter.
	ConversationID param.Opt[string] `query:"conversationId,omitzero" json:"-"`
	// Pagination cursor for thread tweets
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Match this literal phrase, including any hyphens.
	ExactPhrase param.Opt[string] `query:"exactPhrase,omitzero" json:"-"`
	// Words or quoted phrases to exclude. Separate with spaces, commas, or lines.
	ExcludeWords param.Opt[string] `query:"excludeWords,omitzero" json:"-"`
	// Filter by author username.
	FromUser param.Opt[string] `query:"fromUser,omitzero" json:"-"`
	// Hashtags separated by spaces, commas, or lines.
	Hashtags param.Opt[string] `query:"hashtags,omitzero" json:"-"`
	// Only replies to this tweet ID.
	InReplyToTweetID param.Opt[string] `query:"inReplyToTweetId,omitzero" json:"-"`
	// Filter by language. Alias `lang` is accepted.
	Language param.Opt[string] `query:"language,omitzero" json:"-"`
	// Maximum likes threshold. maxLikes is also accepted.
	MaxFaves param.Opt[int64] `query:"maxFaves,omitzero" json:"-"`
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
	// Filter replies sent to a username.
	ToUser param.Opt[string] `query:"toUser,omitzero" json:"-"`
	// End date in YYYY-MM-DD format.
	UntilDate param.Opt[time.Time] `query:"untilDate,omitzero" format:"date" json:"-"`
	// URL substring or domain filter.
	URL param.Opt[string] `query:"url,omitzero" json:"-"`
	// Only return tweets from verified authors.
	VerifiedOnly param.Opt[bool] `query:"verifiedOnly,omitzero" json:"-"`
	// Filter media. Aliases: has_video, has_media.
	//
	// Any of "images", "videos", "gifs", "media", "links", "none".
	MediaType XTweetGetThreadParamsMediaType `query:"mediaType,omitzero" json:"-"`
	// Only when the caller requests a quote mode.
	//
	// Any of "include", "exclude", "only".
	Quotes XTweetGetThreadParamsQuotes `query:"quotes,omitzero" json:"-"`
	// Only when the caller requests a reply mode.
	//
	// Any of "include", "exclude", "only".
	Replies XTweetGetThreadParamsReplies `query:"replies,omitzero" json:"-"`
	// Only when the caller requests a repost mode.
	//
	// Any of "include", "exclude", "only".
	Retweets XTweetGetThreadParamsRetweets `query:"retweets,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [XTweetGetThreadParams]'s query parameters as `url.Values`.
func (r XTweetGetThreadParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter media. Aliases: has_video, has_media.
type XTweetGetThreadParamsMediaType string

const (
	XTweetGetThreadParamsMediaTypeImages XTweetGetThreadParamsMediaType = "images"
	XTweetGetThreadParamsMediaTypeVideos XTweetGetThreadParamsMediaType = "videos"
	XTweetGetThreadParamsMediaTypeGifs   XTweetGetThreadParamsMediaType = "gifs"
	XTweetGetThreadParamsMediaTypeMedia  XTweetGetThreadParamsMediaType = "media"
	XTweetGetThreadParamsMediaTypeLinks  XTweetGetThreadParamsMediaType = "links"
	XTweetGetThreadParamsMediaTypeNone   XTweetGetThreadParamsMediaType = "none"
)

// Only when the caller requests a quote mode.
type XTweetGetThreadParamsQuotes string

const (
	XTweetGetThreadParamsQuotesInclude XTweetGetThreadParamsQuotes = "include"
	XTweetGetThreadParamsQuotesExclude XTweetGetThreadParamsQuotes = "exclude"
	XTweetGetThreadParamsQuotesOnly    XTweetGetThreadParamsQuotes = "only"
)

// Only when the caller requests a reply mode.
type XTweetGetThreadParamsReplies string

const (
	XTweetGetThreadParamsRepliesInclude XTweetGetThreadParamsReplies = "include"
	XTweetGetThreadParamsRepliesExclude XTweetGetThreadParamsReplies = "exclude"
	XTweetGetThreadParamsRepliesOnly    XTweetGetThreadParamsReplies = "only"
)

// Only when the caller requests a repost mode.
type XTweetGetThreadParamsRetweets string

const (
	XTweetGetThreadParamsRetweetsInclude XTweetGetThreadParamsRetweets = "include"
	XTweetGetThreadParamsRetweetsExclude XTweetGetThreadParamsRetweets = "exclude"
	XTweetGetThreadParamsRetweetsOnly    XTweetGetThreadParamsRetweets = "only"
)

type XTweetSearchParams struct {
	// Query, Tweet ID, or URL. Hyphens negate terms. Use exactPhrase for literals.
	// Valid bounds apply per page.
	Q string `query:"q" api:"required" json:"-"`
	// Raw advanced search query appended as-is.
	AdvancedQuery param.Opt[string] `query:"advancedQuery,omitzero" json:"-"`
	// Words or quoted phrases where any one can match. Separate with spaces, commas,
	// or lines.
	AnyWords param.Opt[string] `query:"anyWords,omitzero" json:"-"`
	// Only return tweets from Blue-verified authors.
	BlueVerifiedOnly param.Opt[bool] `query:"blueVerifiedOnly,omitzero" json:"-"`
	// Geo bounding box, e.g. -74.1 40.6 -73.9 40.8.
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
	// Unique matching result upper bound after filtering. Default 20. Explicit
	// coverage defaults to 2000. It returns retained rows and deadline diagnostics
	// when time expires. Only returned rows are billed. Credits may reduce results;
	// zero affordable rows returns 402. Aliases: pageSize, count, max_results.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
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
	MediaType XTweetSearchParamsMediaType `query:"mediaType,omitzero" json:"-"`
	// Omit mode for resumable maximum coverage. Standard keeps legacy pagination.
	// Coverage returns diagnostics once and rejects cursors.
	//
	// Any of "standard", "coverage".
	Mode XTweetSearchParamsMode `query:"mode,omitzero" json:"-"`
	// Latest is chronological; Top ranks engagement. Aliases: result_type, sort_order.
	//
	// Any of "Latest", "Top".
	QueryType XTweetSearchParamsQueryType `query:"queryType,omitzero" json:"-"`
	// Only when the caller requests a quote mode.
	//
	// Any of "include", "exclude", "only".
	Quotes XTweetSearchParamsQuotes `query:"quotes,omitzero" json:"-"`
	// Only when the caller requests a reply mode.
	//
	// Any of "include", "exclude", "only".
	Replies XTweetSearchParamsReplies `query:"replies,omitzero" json:"-"`
	// Only when the caller requests a repost mode.
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

// Filter media. Aliases: has_video, has_media.
type XTweetSearchParamsMediaType string

const (
	XTweetSearchParamsMediaTypeImages XTweetSearchParamsMediaType = "images"
	XTweetSearchParamsMediaTypeVideos XTweetSearchParamsMediaType = "videos"
	XTweetSearchParamsMediaTypeGifs   XTweetSearchParamsMediaType = "gifs"
	XTweetSearchParamsMediaTypeMedia  XTweetSearchParamsMediaType = "media"
	XTweetSearchParamsMediaTypeLinks  XTweetSearchParamsMediaType = "links"
	XTweetSearchParamsMediaTypeNone   XTweetSearchParamsMediaType = "none"
)

// Omit mode for resumable maximum coverage. Standard keeps legacy pagination.
// Coverage returns diagnostics once and rejects cursors.
type XTweetSearchParamsMode string

const (
	XTweetSearchParamsModeStandard XTweetSearchParamsMode = "standard"
	XTweetSearchParamsModeCoverage XTweetSearchParamsMode = "coverage"
)

// Latest is chronological; Top ranks engagement. Aliases: result_type, sort_order.
type XTweetSearchParamsQueryType string

const (
	XTweetSearchParamsQueryTypeLatest XTweetSearchParamsQueryType = "Latest"
	XTweetSearchParamsQueryTypeTop    XTweetSearchParamsQueryType = "Top"
)

// Only when the caller requests a quote mode.
type XTweetSearchParamsQuotes string

const (
	XTweetSearchParamsQuotesInclude XTweetSearchParamsQuotes = "include"
	XTweetSearchParamsQuotesExclude XTweetSearchParamsQuotes = "exclude"
	XTweetSearchParamsQuotesOnly    XTweetSearchParamsQuotes = "only"
)

// Only when the caller requests a reply mode.
type XTweetSearchParamsReplies string

const (
	XTweetSearchParamsRepliesInclude XTweetSearchParamsReplies = "include"
	XTweetSearchParamsRepliesExclude XTweetSearchParamsReplies = "exclude"
	XTweetSearchParamsRepliesOnly    XTweetSearchParamsReplies = "only"
)

// Only when the caller requests a repost mode.
type XTweetSearchParamsRetweets string

const (
	XTweetSearchParamsRetweetsInclude XTweetSearchParamsRetweets = "include"
	XTweetSearchParamsRetweetsExclude XTweetSearchParamsRetweets = "exclude"
	XTweetSearchParamsRetweetsOnly    XTweetSearchParamsRetweets = "only"
)
