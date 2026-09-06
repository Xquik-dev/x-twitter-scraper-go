// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

import (
	"time"

	"github.com/Xquik-dev/x-twitter-scraper-go/internal/apijson"
	"github.com/Xquik-dev/x-twitter-scraper-go/packages/param"
	"github.com/Xquik-dev/x-twitter-scraper-go/packages/respjson"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

// Content disclosure metadata shown by X when a tweet is labeled as paid
// partnership content or AI-generated media.
type ContentDisclosure struct {
	Advertising ContentDisclosureAdvertising `json:"advertising"`
	AIGenerated ContentDisclosureAIGenerated `json:"aiGenerated"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Advertising respjson.Field
		AIGenerated respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContentDisclosure) RawJSON() string { return r.JSON.raw }
func (r *ContentDisclosure) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContentDisclosureAdvertising struct {
	// True when X labels the tweet as paid promotion content.
	IsPaidPromotion bool `json:"isPaidPromotion"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IsPaidPromotion respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContentDisclosureAdvertising) RawJSON() string { return r.JSON.raw }
func (r *ContentDisclosureAdvertising) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContentDisclosureAIGenerated struct {
	// Source of the AI-generated media disclosure.
	DetectionSource string `json:"detectionSource"`
	// True when X labels the tweet as containing AI-generated media.
	HasAIGeneratedMedia bool `json:"hasAiGeneratedMedia"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DetectionSource     respjson.Field
		HasAIGeneratedMedia respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContentDisclosureAIGenerated) RawJSON() string { return r.JSON.raw }
func (r *ContentDisclosureAIGenerated) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Quoted or retweeted tweet context.
type EmbeddedTweet struct {
	ID            string `json:"id" api:"required"`
	BookmarkCount int64  `json:"bookmarkCount" api:"required"`
	LikeCount     int64  `json:"likeCount" api:"required"`
	QuoteCount    int64  `json:"quoteCount" api:"required"`
	ReplyCount    int64  `json:"replyCount" api:"required"`
	RetweetCount  int64  `json:"retweetCount" api:"required"`
	Text          string `json:"text" api:"required"`
	ViewCount     int64  `json:"viewCount" api:"required"`
	// Describes an X Article preview and its lifecycle metadata.
	Article EmbeddedTweetArticle `json:"article"`
	// X user profile with bio, follower counts, and verification status.
	Author UserProfile `json:"author"`
	// Describes a public card and its referenced profiles.
	Card EmbeddedTweetCard `json:"card"`
	// Community ID.
	CommunityID string `json:"communityId"`
	// Community Note presentation metadata returned by X.
	CommunityNote EmbeddedTweetCommunityNote `json:"communityNote"`
	// Content disclosure metadata shown by X when a tweet is labeled as paid
	// partnership content or AI-generated media.
	ContentDisclosure ContentDisclosure `json:"contentDisclosure"`
	// Public reply policy and conversation owner.
	ConversationControl EmbeddedTweetConversationControl `json:"conversationControl"`
	ConversationID      string                           `json:"conversationId"`
	CreatedAt           string                           `json:"createdAt"`
	DisplayTextRange    []int64                          `json:"displayTextRange"`
	// Lists edit-chain identifiers and the remaining edit window.
	Edit EmbeddedTweetEdit `json:"edit"`
	// Lists hashtags, symbols, links, and mentions from tweet text.
	Entities EmbeddedTweetEntities `json:"entities"`
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
	LimitedActions []EmbeddedTweetLimitedAction `json:"limitedActions"`
	// Attached media items, omitted when unavailable.
	Media []TweetMedia `json:"media"`
	// Complete Note Tweet content and rich-text metadata.
	NoteTweet EmbeddedTweetNoteTweet `json:"noteTweet"`
	// Describes public place metadata on a geotagged tweet.
	Place             EmbeddedTweetPlace `json:"place"`
	PossiblySensitive bool               `json:"possiblySensitive"`
	// Public metadata whose fields are defined by X.
	PostCta map[string]any `json:"postCta"`
	// Engagement counts retained from a prior tweet edit.
	PreviousCounts EmbeddedTweetPreviousCounts `json:"previousCounts"`
	// Nested tweet context at depth 2.
	QuotedTweet EmbeddedTweetQuotedTweet `json:"quoted_tweet"`
	// Quoted tweet ID.
	QuotedTweetID string `json:"quotedTweetId"`
	// Public post and user referenced by this reaction.
	ReactionContext EmbeddedTweetReactionContext `json:"reactionContext"`
	// Nested tweet context at depth 2.
	RetweetedTweet EmbeddedTweetRetweetedTweet `json:"retweeted_tweet"`
	// Public metadata whose fields are defined by X.
	Scopes map[string]any `json:"scopes"`
	Source string         `json:"source"`
	// Public visibility notice attached to an available tweet.
	Tombstone EmbeddedTweetTombstone `json:"tombstone"`
	Type      string                 `json:"type"`
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
func (r EmbeddedTweet) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Describes an X Article preview and its lifecycle metadata.
type EmbeddedTweetArticle struct {
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
func (r EmbeddedTweetArticle) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetArticle) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Describes a public card and its referenced profiles.
type EmbeddedTweetCard struct {
	ID string `json:"id"`
	// Public metadata whose fields are defined by X.
	BindingValues map[string]any `json:"bindingValues"`
	Name          string         `json:"name"`
	// Public metadata whose fields are defined by X.
	Platform map[string]any `json:"platform"`
	URL      string         `json:"url"`
	// Unresolved card user references.
	UserReferenceErrors []EmbeddedTweetCardUserReferenceError `json:"userReferenceErrors"`
	UserReferences      []UserProfile                         `json:"userReferences"`
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
func (r EmbeddedTweetCard) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetCard) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetCardUserReferenceError struct {
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
func (r EmbeddedTweetCardUserReferenceError) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetCardUserReferenceError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Community Note presentation metadata returned by X.
type EmbeddedTweetCommunityNote struct {
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
func (r EmbeddedTweetCommunityNote) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetCommunityNote) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Public reply policy and conversation owner.
type EmbeddedTweetConversationControl struct {
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
func (r EmbeddedTweetConversationControl) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetConversationControl) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Lists edit-chain identifiers and the remaining edit window.
type EmbeddedTweetEdit struct {
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
func (r EmbeddedTweetEdit) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetEdit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Lists hashtags, symbols, links, and mentions from tweet text.
type EmbeddedTweetEntities struct {
	Hashtags     []EmbeddedTweetEntitiesHashtag     `json:"hashtags"`
	Smarttags    []EmbeddedTweetEntitiesSmarttag    `json:"smarttags"`
	Symbols      []EmbeddedTweetEntitiesSymbol      `json:"symbols"`
	Timestamps   []EmbeddedTweetEntitiesTimestamp   `json:"timestamps"`
	URLs         []EmbeddedTweetEntitiesURL         `json:"urls"`
	UserMentions []EmbeddedTweetEntitiesUserMention `json:"user_mentions"`
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
func (r EmbeddedTweetEntities) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetEntities) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Provides hashtag text and source offsets within a tweet.
type EmbeddedTweetEntitiesHashtag struct {
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
func (r EmbeddedTweetEntitiesHashtag) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetEntitiesHashtag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indexed smart-tag, cashtag, or video timestamp metadata.
type EmbeddedTweetEntitiesSmarttag struct {
	Indices []int64                          `json:"indices"`
	Seconds float64                          `json:"seconds"`
	Tag     EmbeddedTweetEntitiesSmarttagTag `json:"tag"`
	Text    string                           `json:"text"`
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
func (r EmbeddedTweetEntitiesSmarttag) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetEntitiesSmarttag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetEntitiesSmarttagTag struct {
	Info EmbeddedTweetEntitiesSmarttagTagInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetEntitiesSmarttagTag) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetEntitiesSmarttagTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetEntitiesSmarttagTagInfo struct {
	Info EmbeddedTweetEntitiesSmarttagTagInfoInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetEntitiesSmarttagTagInfo) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetEntitiesSmarttagTagInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetEntitiesSmarttagTagInfoInfo struct {
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
func (r EmbeddedTweetEntitiesSmarttagTagInfoInfo) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetEntitiesSmarttagTagInfoInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indexed smart-tag, cashtag, or video timestamp metadata.
type EmbeddedTweetEntitiesSymbol struct {
	Indices []int64                        `json:"indices"`
	Seconds float64                        `json:"seconds"`
	Tag     EmbeddedTweetEntitiesSymbolTag `json:"tag"`
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
func (r EmbeddedTweetEntitiesSymbol) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetEntitiesSymbol) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetEntitiesSymbolTag struct {
	Info EmbeddedTweetEntitiesSymbolTagInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetEntitiesSymbolTag) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetEntitiesSymbolTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetEntitiesSymbolTagInfo struct {
	Info EmbeddedTweetEntitiesSymbolTagInfoInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetEntitiesSymbolTagInfo) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetEntitiesSymbolTagInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetEntitiesSymbolTagInfoInfo struct {
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
func (r EmbeddedTweetEntitiesSymbolTagInfoInfo) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetEntitiesSymbolTagInfoInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indexed smart-tag, cashtag, or video timestamp metadata.
type EmbeddedTweetEntitiesTimestamp struct {
	Indices []int64                           `json:"indices"`
	Seconds float64                           `json:"seconds"`
	Tag     EmbeddedTweetEntitiesTimestampTag `json:"tag"`
	Text    string                            `json:"text"`
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
func (r EmbeddedTweetEntitiesTimestamp) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetEntitiesTimestamp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetEntitiesTimestampTag struct {
	Info EmbeddedTweetEntitiesTimestampTagInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetEntitiesTimestampTag) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetEntitiesTimestampTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetEntitiesTimestampTagInfo struct {
	Info EmbeddedTweetEntitiesTimestampTagInfoInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetEntitiesTimestampTagInfo) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetEntitiesTimestampTagInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetEntitiesTimestampTagInfoInfo struct {
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
func (r EmbeddedTweetEntitiesTimestampTagInfoInfo) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetEntitiesTimestampTagInfoInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Provides shortened, display, and expanded URLs from tweet text.
type EmbeddedTweetEntitiesURL struct {
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
func (r EmbeddedTweetEntitiesURL) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetEntitiesURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Provides profile identity and source offsets for a mention.
type EmbeddedTweetEntitiesUserMention struct {
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
func (r EmbeddedTweetEntitiesUserMention) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetEntitiesUserMention) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetLimitedAction struct {
	Action string                           `json:"action"`
	Prompt EmbeddedTweetLimitedActionPrompt `json:"prompt"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Action      respjson.Field
		Prompt      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetLimitedAction) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetLimitedAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetLimitedActionPrompt struct {
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
func (r EmbeddedTweetLimitedActionPrompt) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetLimitedActionPrompt) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Complete Note Tweet content and rich-text metadata.
type EmbeddedTweetNoteTweet struct {
	Text string `json:"text" api:"required"`
	ID   string `json:"id"`
	// Lists hashtags, symbols, links, and mentions from tweet text.
	Entities EmbeddedTweetNoteTweetEntities `json:"entities"`
	// Inline media positions in the Note Tweet text.
	InlineMedia  []EmbeddedTweetNoteTweetInlineMedia `json:"inlineMedia"`
	IsExpandable bool                                `json:"isExpandable"`
	RichtextTags []EmbeddedTweetNoteTweetRichtextTag `json:"richtextTags"`
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
func (r EmbeddedTweetNoteTweet) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetNoteTweet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Lists hashtags, symbols, links, and mentions from tweet text.
type EmbeddedTweetNoteTweetEntities struct {
	Hashtags     []EmbeddedTweetNoteTweetEntitiesHashtag     `json:"hashtags"`
	Smarttags    []EmbeddedTweetNoteTweetEntitiesSmarttag    `json:"smarttags"`
	Symbols      []EmbeddedTweetNoteTweetEntitiesSymbol      `json:"symbols"`
	Timestamps   []EmbeddedTweetNoteTweetEntitiesTimestamp   `json:"timestamps"`
	URLs         []EmbeddedTweetNoteTweetEntitiesURL         `json:"urls"`
	UserMentions []EmbeddedTweetNoteTweetEntitiesUserMention `json:"user_mentions"`
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
func (r EmbeddedTweetNoteTweetEntities) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetNoteTweetEntities) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Provides hashtag text and source offsets within a tweet.
type EmbeddedTweetNoteTweetEntitiesHashtag struct {
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
func (r EmbeddedTweetNoteTweetEntitiesHashtag) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetNoteTweetEntitiesHashtag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indexed smart-tag, cashtag, or video timestamp metadata.
type EmbeddedTweetNoteTweetEntitiesSmarttag struct {
	Indices []int64                                   `json:"indices"`
	Seconds float64                                   `json:"seconds"`
	Tag     EmbeddedTweetNoteTweetEntitiesSmarttagTag `json:"tag"`
	Text    string                                    `json:"text"`
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
func (r EmbeddedTweetNoteTweetEntitiesSmarttag) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetNoteTweetEntitiesSmarttag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetNoteTweetEntitiesSmarttagTag struct {
	Info EmbeddedTweetNoteTweetEntitiesSmarttagTagInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetNoteTweetEntitiesSmarttagTag) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetNoteTweetEntitiesSmarttagTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetNoteTweetEntitiesSmarttagTagInfo struct {
	Info EmbeddedTweetNoteTweetEntitiesSmarttagTagInfoInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetNoteTweetEntitiesSmarttagTagInfo) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetNoteTweetEntitiesSmarttagTagInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetNoteTweetEntitiesSmarttagTagInfoInfo struct {
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
func (r EmbeddedTweetNoteTweetEntitiesSmarttagTagInfoInfo) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetNoteTweetEntitiesSmarttagTagInfoInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indexed smart-tag, cashtag, or video timestamp metadata.
type EmbeddedTweetNoteTweetEntitiesSymbol struct {
	Indices []int64                                 `json:"indices"`
	Seconds float64                                 `json:"seconds"`
	Tag     EmbeddedTweetNoteTweetEntitiesSymbolTag `json:"tag"`
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
func (r EmbeddedTweetNoteTweetEntitiesSymbol) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetNoteTweetEntitiesSymbol) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetNoteTweetEntitiesSymbolTag struct {
	Info EmbeddedTweetNoteTweetEntitiesSymbolTagInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetNoteTweetEntitiesSymbolTag) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetNoteTweetEntitiesSymbolTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetNoteTweetEntitiesSymbolTagInfo struct {
	Info EmbeddedTweetNoteTweetEntitiesSymbolTagInfoInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetNoteTweetEntitiesSymbolTagInfo) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetNoteTweetEntitiesSymbolTagInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetNoteTweetEntitiesSymbolTagInfoInfo struct {
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
func (r EmbeddedTweetNoteTweetEntitiesSymbolTagInfoInfo) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetNoteTweetEntitiesSymbolTagInfoInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indexed smart-tag, cashtag, or video timestamp metadata.
type EmbeddedTweetNoteTweetEntitiesTimestamp struct {
	Indices []int64                                    `json:"indices"`
	Seconds float64                                    `json:"seconds"`
	Tag     EmbeddedTweetNoteTweetEntitiesTimestampTag `json:"tag"`
	Text    string                                     `json:"text"`
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
func (r EmbeddedTweetNoteTweetEntitiesTimestamp) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetNoteTweetEntitiesTimestamp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetNoteTweetEntitiesTimestampTag struct {
	Info EmbeddedTweetNoteTweetEntitiesTimestampTagInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetNoteTweetEntitiesTimestampTag) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetNoteTweetEntitiesTimestampTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetNoteTweetEntitiesTimestampTagInfo struct {
	Info EmbeddedTweetNoteTweetEntitiesTimestampTagInfoInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetNoteTweetEntitiesTimestampTagInfo) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetNoteTweetEntitiesTimestampTagInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetNoteTweetEntitiesTimestampTagInfoInfo struct {
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
func (r EmbeddedTweetNoteTweetEntitiesTimestampTagInfoInfo) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetNoteTweetEntitiesTimestampTagInfoInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Provides shortened, display, and expanded URLs from tweet text.
type EmbeddedTweetNoteTweetEntitiesURL struct {
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
func (r EmbeddedTweetNoteTweetEntitiesURL) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetNoteTweetEntitiesURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Provides profile identity and source offsets for a mention.
type EmbeddedTweetNoteTweetEntitiesUserMention struct {
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
func (r EmbeddedTweetNoteTweetEntitiesUserMention) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetNoteTweetEntitiesUserMention) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetNoteTweetInlineMedia struct {
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
func (r EmbeddedTweetNoteTweetInlineMedia) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetNoteTweetInlineMedia) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetNoteTweetRichtextTag struct {
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
func (r EmbeddedTweetNoteTweetRichtextTag) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetNoteTweetRichtextTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Describes public place metadata on a geotagged tweet.
type EmbeddedTweetPlace struct {
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
func (r EmbeddedTweetPlace) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetPlace) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Engagement counts retained from a prior tweet edit.
type EmbeddedTweetPreviousCounts struct {
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
func (r EmbeddedTweetPreviousCounts) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetPreviousCounts) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Nested tweet context at depth 2.
type EmbeddedTweetQuotedTweet struct {
	ID            string `json:"id" api:"required"`
	BookmarkCount int64  `json:"bookmarkCount" api:"required"`
	LikeCount     int64  `json:"likeCount" api:"required"`
	QuoteCount    int64  `json:"quoteCount" api:"required"`
	ReplyCount    int64  `json:"replyCount" api:"required"`
	RetweetCount  int64  `json:"retweetCount" api:"required"`
	Text          string `json:"text" api:"required"`
	ViewCount     int64  `json:"viewCount" api:"required"`
	// Describes an X Article preview and its lifecycle metadata.
	Article EmbeddedTweetQuotedTweetArticle `json:"article"`
	// X user profile with bio, follower counts, and verification status.
	Author UserProfile `json:"author"`
	// Describes a public card and its referenced profiles.
	Card EmbeddedTweetQuotedTweetCard `json:"card"`
	// Community ID.
	CommunityID string `json:"communityId"`
	// Community Note presentation metadata returned by X.
	CommunityNote EmbeddedTweetQuotedTweetCommunityNote `json:"communityNote"`
	// Content disclosure metadata shown by X when a tweet is labeled as paid
	// partnership content or AI-generated media.
	ContentDisclosure ContentDisclosure `json:"contentDisclosure"`
	// Public reply policy and conversation owner.
	ConversationControl EmbeddedTweetQuotedTweetConversationControl `json:"conversationControl"`
	ConversationID      string                                      `json:"conversationId"`
	CreatedAt           string                                      `json:"createdAt"`
	DisplayTextRange    []int64                                     `json:"displayTextRange"`
	// Lists edit-chain identifiers and the remaining edit window.
	Edit EmbeddedTweetQuotedTweetEdit `json:"edit"`
	// Lists hashtags, symbols, links, and mentions from tweet text.
	Entities EmbeddedTweetQuotedTweetEntities `json:"entities"`
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
	LimitedActions []EmbeddedTweetQuotedTweetLimitedAction `json:"limitedActions"`
	// Attached media items, omitted when unavailable.
	Media []TweetMedia `json:"media"`
	// Complete Note Tweet content and rich-text metadata.
	NoteTweet EmbeddedTweetQuotedTweetNoteTweet `json:"noteTweet"`
	// Describes public place metadata on a geotagged tweet.
	Place             EmbeddedTweetQuotedTweetPlace `json:"place"`
	PossiblySensitive bool                          `json:"possiblySensitive"`
	// Public metadata whose fields are defined by X.
	PostCta map[string]any `json:"postCta"`
	// Engagement counts retained from a prior tweet edit.
	PreviousCounts EmbeddedTweetQuotedTweetPreviousCounts `json:"previousCounts"`
	// Nested tweet context at depth 3.
	QuotedTweet EmbeddedTweetQuotedTweetQuotedTweet `json:"quoted_tweet"`
	// Quoted tweet ID.
	QuotedTweetID string `json:"quotedTweetId"`
	// Public post and user referenced by this reaction.
	ReactionContext EmbeddedTweetQuotedTweetReactionContext `json:"reactionContext"`
	// Nested tweet context at depth 3.
	RetweetedTweet EmbeddedTweetQuotedTweetRetweetedTweet `json:"retweeted_tweet"`
	// Public metadata whose fields are defined by X.
	Scopes map[string]any `json:"scopes"`
	Source string         `json:"source"`
	// Public visibility notice attached to an available tweet.
	Tombstone EmbeddedTweetQuotedTweetTombstone `json:"tombstone"`
	Type      string                            `json:"type"`
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
func (r EmbeddedTweetQuotedTweet) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetQuotedTweet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Describes an X Article preview and its lifecycle metadata.
type EmbeddedTweetQuotedTweetArticle struct {
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
func (r EmbeddedTweetQuotedTweetArticle) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetQuotedTweetArticle) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Describes a public card and its referenced profiles.
type EmbeddedTweetQuotedTweetCard struct {
	ID string `json:"id"`
	// Public metadata whose fields are defined by X.
	BindingValues map[string]any `json:"bindingValues"`
	Name          string         `json:"name"`
	// Public metadata whose fields are defined by X.
	Platform map[string]any `json:"platform"`
	URL      string         `json:"url"`
	// Unresolved card user references.
	UserReferenceErrors []EmbeddedTweetQuotedTweetCardUserReferenceError `json:"userReferenceErrors"`
	UserReferences      []UserProfile                                    `json:"userReferences"`
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
func (r EmbeddedTweetQuotedTweetCard) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetQuotedTweetCard) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetCardUserReferenceError struct {
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
func (r EmbeddedTweetQuotedTweetCardUserReferenceError) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetQuotedTweetCardUserReferenceError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Community Note presentation metadata returned by X.
type EmbeddedTweetQuotedTweetCommunityNote struct {
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
func (r EmbeddedTweetQuotedTweetCommunityNote) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetQuotedTweetCommunityNote) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Public reply policy and conversation owner.
type EmbeddedTweetQuotedTweetConversationControl struct {
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
func (r EmbeddedTweetQuotedTweetConversationControl) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetQuotedTweetConversationControl) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Lists edit-chain identifiers and the remaining edit window.
type EmbeddedTweetQuotedTweetEdit struct {
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
func (r EmbeddedTweetQuotedTweetEdit) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetQuotedTweetEdit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Lists hashtags, symbols, links, and mentions from tweet text.
type EmbeddedTweetQuotedTweetEntities struct {
	Hashtags     []EmbeddedTweetQuotedTweetEntitiesHashtag     `json:"hashtags"`
	Smarttags    []EmbeddedTweetQuotedTweetEntitiesSmarttag    `json:"smarttags"`
	Symbols      []EmbeddedTweetQuotedTweetEntitiesSymbol      `json:"symbols"`
	Timestamps   []EmbeddedTweetQuotedTweetEntitiesTimestamp   `json:"timestamps"`
	URLs         []EmbeddedTweetQuotedTweetEntitiesURL         `json:"urls"`
	UserMentions []EmbeddedTweetQuotedTweetEntitiesUserMention `json:"user_mentions"`
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
func (r EmbeddedTweetQuotedTweetEntities) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetQuotedTweetEntities) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Provides hashtag text and source offsets within a tweet.
type EmbeddedTweetQuotedTweetEntitiesHashtag struct {
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
func (r EmbeddedTweetQuotedTweetEntitiesHashtag) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetQuotedTweetEntitiesHashtag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indexed smart-tag, cashtag, or video timestamp metadata.
type EmbeddedTweetQuotedTweetEntitiesSmarttag struct {
	Indices []int64                                     `json:"indices"`
	Seconds float64                                     `json:"seconds"`
	Tag     EmbeddedTweetQuotedTweetEntitiesSmarttagTag `json:"tag"`
	Text    string                                      `json:"text"`
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
func (r EmbeddedTweetQuotedTweetEntitiesSmarttag) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetQuotedTweetEntitiesSmarttag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetEntitiesSmarttagTag struct {
	Info EmbeddedTweetQuotedTweetEntitiesSmarttagTagInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetQuotedTweetEntitiesSmarttagTag) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetQuotedTweetEntitiesSmarttagTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetEntitiesSmarttagTagInfo struct {
	Info EmbeddedTweetQuotedTweetEntitiesSmarttagTagInfoInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetQuotedTweetEntitiesSmarttagTagInfo) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetQuotedTweetEntitiesSmarttagTagInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetEntitiesSmarttagTagInfoInfo struct {
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
func (r EmbeddedTweetQuotedTweetEntitiesSmarttagTagInfoInfo) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetQuotedTweetEntitiesSmarttagTagInfoInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indexed smart-tag, cashtag, or video timestamp metadata.
type EmbeddedTweetQuotedTweetEntitiesSymbol struct {
	Indices []int64                                   `json:"indices"`
	Seconds float64                                   `json:"seconds"`
	Tag     EmbeddedTweetQuotedTweetEntitiesSymbolTag `json:"tag"`
	Text    string                                    `json:"text"`
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
func (r EmbeddedTweetQuotedTweetEntitiesSymbol) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetQuotedTweetEntitiesSymbol) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetEntitiesSymbolTag struct {
	Info EmbeddedTweetQuotedTweetEntitiesSymbolTagInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetQuotedTweetEntitiesSymbolTag) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetQuotedTweetEntitiesSymbolTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetEntitiesSymbolTagInfo struct {
	Info EmbeddedTweetQuotedTweetEntitiesSymbolTagInfoInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetQuotedTweetEntitiesSymbolTagInfo) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetQuotedTweetEntitiesSymbolTagInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetEntitiesSymbolTagInfoInfo struct {
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
func (r EmbeddedTweetQuotedTweetEntitiesSymbolTagInfoInfo) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetQuotedTweetEntitiesSymbolTagInfoInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indexed smart-tag, cashtag, or video timestamp metadata.
type EmbeddedTweetQuotedTweetEntitiesTimestamp struct {
	Indices []int64                                      `json:"indices"`
	Seconds float64                                      `json:"seconds"`
	Tag     EmbeddedTweetQuotedTweetEntitiesTimestampTag `json:"tag"`
	Text    string                                       `json:"text"`
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
func (r EmbeddedTweetQuotedTweetEntitiesTimestamp) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetQuotedTweetEntitiesTimestamp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetEntitiesTimestampTag struct {
	Info EmbeddedTweetQuotedTweetEntitiesTimestampTagInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetQuotedTweetEntitiesTimestampTag) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetQuotedTweetEntitiesTimestampTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetEntitiesTimestampTagInfo struct {
	Info EmbeddedTweetQuotedTweetEntitiesTimestampTagInfoInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetQuotedTweetEntitiesTimestampTagInfo) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetQuotedTweetEntitiesTimestampTagInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetEntitiesTimestampTagInfoInfo struct {
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
func (r EmbeddedTweetQuotedTweetEntitiesTimestampTagInfoInfo) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetQuotedTweetEntitiesTimestampTagInfoInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Provides shortened, display, and expanded URLs from tweet text.
type EmbeddedTweetQuotedTweetEntitiesURL struct {
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
func (r EmbeddedTweetQuotedTweetEntitiesURL) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetQuotedTweetEntitiesURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Provides profile identity and source offsets for a mention.
type EmbeddedTweetQuotedTweetEntitiesUserMention struct {
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
func (r EmbeddedTweetQuotedTweetEntitiesUserMention) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetQuotedTweetEntitiesUserMention) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetLimitedAction struct {
	Action string                                      `json:"action"`
	Prompt EmbeddedTweetQuotedTweetLimitedActionPrompt `json:"prompt"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Action      respjson.Field
		Prompt      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetQuotedTweetLimitedAction) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetQuotedTweetLimitedAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetLimitedActionPrompt struct {
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
func (r EmbeddedTweetQuotedTweetLimitedActionPrompt) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetQuotedTweetLimitedActionPrompt) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Complete Note Tweet content and rich-text metadata.
type EmbeddedTweetQuotedTweetNoteTweet struct {
	Text string `json:"text" api:"required"`
	ID   string `json:"id"`
	// Lists hashtags, symbols, links, and mentions from tweet text.
	Entities EmbeddedTweetQuotedTweetNoteTweetEntities `json:"entities"`
	// Inline media positions in the Note Tweet text.
	InlineMedia  []EmbeddedTweetQuotedTweetNoteTweetInlineMedia `json:"inlineMedia"`
	IsExpandable bool                                           `json:"isExpandable"`
	RichtextTags []EmbeddedTweetQuotedTweetNoteTweetRichtextTag `json:"richtextTags"`
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
func (r EmbeddedTweetQuotedTweetNoteTweet) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetQuotedTweetNoteTweet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Lists hashtags, symbols, links, and mentions from tweet text.
type EmbeddedTweetQuotedTweetNoteTweetEntities struct {
	Hashtags     []EmbeddedTweetQuotedTweetNoteTweetEntitiesHashtag     `json:"hashtags"`
	Smarttags    []EmbeddedTweetQuotedTweetNoteTweetEntitiesSmarttag    `json:"smarttags"`
	Symbols      []EmbeddedTweetQuotedTweetNoteTweetEntitiesSymbol      `json:"symbols"`
	Timestamps   []EmbeddedTweetQuotedTweetNoteTweetEntitiesTimestamp   `json:"timestamps"`
	URLs         []EmbeddedTweetQuotedTweetNoteTweetEntitiesURL         `json:"urls"`
	UserMentions []EmbeddedTweetQuotedTweetNoteTweetEntitiesUserMention `json:"user_mentions"`
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
func (r EmbeddedTweetQuotedTweetNoteTweetEntities) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetQuotedTweetNoteTweetEntities) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Provides hashtag text and source offsets within a tweet.
type EmbeddedTweetQuotedTweetNoteTweetEntitiesHashtag struct {
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
func (r EmbeddedTweetQuotedTweetNoteTweetEntitiesHashtag) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetQuotedTweetNoteTweetEntitiesHashtag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indexed smart-tag, cashtag, or video timestamp metadata.
type EmbeddedTweetQuotedTweetNoteTweetEntitiesSmarttag struct {
	Indices []int64                                              `json:"indices"`
	Seconds float64                                              `json:"seconds"`
	Tag     EmbeddedTweetQuotedTweetNoteTweetEntitiesSmarttagTag `json:"tag"`
	Text    string                                               `json:"text"`
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
func (r EmbeddedTweetQuotedTweetNoteTweetEntitiesSmarttag) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetQuotedTweetNoteTweetEntitiesSmarttag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetNoteTweetEntitiesSmarttagTag struct {
	Info EmbeddedTweetQuotedTweetNoteTweetEntitiesSmarttagTagInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetQuotedTweetNoteTweetEntitiesSmarttagTag) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetQuotedTweetNoteTweetEntitiesSmarttagTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetNoteTweetEntitiesSmarttagTagInfo struct {
	Info EmbeddedTweetQuotedTweetNoteTweetEntitiesSmarttagTagInfoInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetQuotedTweetNoteTweetEntitiesSmarttagTagInfo) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetQuotedTweetNoteTweetEntitiesSmarttagTagInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetNoteTweetEntitiesSmarttagTagInfoInfo struct {
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
func (r EmbeddedTweetQuotedTweetNoteTweetEntitiesSmarttagTagInfoInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetNoteTweetEntitiesSmarttagTagInfoInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indexed smart-tag, cashtag, or video timestamp metadata.
type EmbeddedTweetQuotedTweetNoteTweetEntitiesSymbol struct {
	Indices []int64                                            `json:"indices"`
	Seconds float64                                            `json:"seconds"`
	Tag     EmbeddedTweetQuotedTweetNoteTweetEntitiesSymbolTag `json:"tag"`
	Text    string                                             `json:"text"`
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
func (r EmbeddedTweetQuotedTweetNoteTweetEntitiesSymbol) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetQuotedTweetNoteTweetEntitiesSymbol) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetNoteTweetEntitiesSymbolTag struct {
	Info EmbeddedTweetQuotedTweetNoteTweetEntitiesSymbolTagInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetQuotedTweetNoteTweetEntitiesSymbolTag) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetQuotedTweetNoteTweetEntitiesSymbolTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetNoteTweetEntitiesSymbolTagInfo struct {
	Info EmbeddedTweetQuotedTweetNoteTweetEntitiesSymbolTagInfoInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetQuotedTweetNoteTweetEntitiesSymbolTagInfo) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetQuotedTweetNoteTweetEntitiesSymbolTagInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetNoteTweetEntitiesSymbolTagInfoInfo struct {
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
func (r EmbeddedTweetQuotedTweetNoteTweetEntitiesSymbolTagInfoInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetNoteTweetEntitiesSymbolTagInfoInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indexed smart-tag, cashtag, or video timestamp metadata.
type EmbeddedTweetQuotedTweetNoteTweetEntitiesTimestamp struct {
	Indices []int64                                               `json:"indices"`
	Seconds float64                                               `json:"seconds"`
	Tag     EmbeddedTweetQuotedTweetNoteTweetEntitiesTimestampTag `json:"tag"`
	Text    string                                                `json:"text"`
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
func (r EmbeddedTweetQuotedTweetNoteTweetEntitiesTimestamp) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetQuotedTweetNoteTweetEntitiesTimestamp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetNoteTweetEntitiesTimestampTag struct {
	Info EmbeddedTweetQuotedTweetNoteTweetEntitiesTimestampTagInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetQuotedTweetNoteTweetEntitiesTimestampTag) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetQuotedTweetNoteTweetEntitiesTimestampTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetNoteTweetEntitiesTimestampTagInfo struct {
	Info EmbeddedTweetQuotedTweetNoteTweetEntitiesTimestampTagInfoInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetQuotedTweetNoteTweetEntitiesTimestampTagInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetNoteTweetEntitiesTimestampTagInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetNoteTweetEntitiesTimestampTagInfoInfo struct {
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
func (r EmbeddedTweetQuotedTweetNoteTweetEntitiesTimestampTagInfoInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetNoteTweetEntitiesTimestampTagInfoInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Provides shortened, display, and expanded URLs from tweet text.
type EmbeddedTweetQuotedTweetNoteTweetEntitiesURL struct {
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
func (r EmbeddedTweetQuotedTweetNoteTweetEntitiesURL) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetQuotedTweetNoteTweetEntitiesURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Provides profile identity and source offsets for a mention.
type EmbeddedTweetQuotedTweetNoteTweetEntitiesUserMention struct {
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
func (r EmbeddedTweetQuotedTweetNoteTweetEntitiesUserMention) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetQuotedTweetNoteTweetEntitiesUserMention) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetNoteTweetInlineMedia struct {
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
func (r EmbeddedTweetQuotedTweetNoteTweetInlineMedia) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetQuotedTweetNoteTweetInlineMedia) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetNoteTweetRichtextTag struct {
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
func (r EmbeddedTweetQuotedTweetNoteTweetRichtextTag) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetQuotedTweetNoteTweetRichtextTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Describes public place metadata on a geotagged tweet.
type EmbeddedTweetQuotedTweetPlace struct {
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
func (r EmbeddedTweetQuotedTweetPlace) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetQuotedTweetPlace) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Engagement counts retained from a prior tweet edit.
type EmbeddedTweetQuotedTweetPreviousCounts struct {
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
func (r EmbeddedTweetQuotedTweetPreviousCounts) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetQuotedTweetPreviousCounts) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Nested tweet context at depth 3.
type EmbeddedTweetQuotedTweetQuotedTweet struct {
	ID            string `json:"id" api:"required"`
	BookmarkCount int64  `json:"bookmarkCount" api:"required"`
	LikeCount     int64  `json:"likeCount" api:"required"`
	QuoteCount    int64  `json:"quoteCount" api:"required"`
	ReplyCount    int64  `json:"replyCount" api:"required"`
	RetweetCount  int64  `json:"retweetCount" api:"required"`
	Text          string `json:"text" api:"required"`
	ViewCount     int64  `json:"viewCount" api:"required"`
	// Describes an X Article preview and its lifecycle metadata.
	Article EmbeddedTweetQuotedTweetQuotedTweetArticle `json:"article"`
	// X user profile with bio, follower counts, and verification status.
	Author UserProfile `json:"author"`
	// Describes a public card and its referenced profiles.
	Card EmbeddedTweetQuotedTweetQuotedTweetCard `json:"card"`
	// Community ID.
	CommunityID string `json:"communityId"`
	// Community Note presentation metadata returned by X.
	CommunityNote EmbeddedTweetQuotedTweetQuotedTweetCommunityNote `json:"communityNote"`
	// Content disclosure metadata shown by X when a tweet is labeled as paid
	// partnership content or AI-generated media.
	ContentDisclosure ContentDisclosure `json:"contentDisclosure"`
	// Public reply policy and conversation owner.
	ConversationControl EmbeddedTweetQuotedTweetQuotedTweetConversationControl `json:"conversationControl"`
	ConversationID      string                                                 `json:"conversationId"`
	CreatedAt           string                                                 `json:"createdAt"`
	DisplayTextRange    []int64                                                `json:"displayTextRange"`
	// Lists edit-chain identifiers and the remaining edit window.
	Edit EmbeddedTweetQuotedTweetQuotedTweetEdit `json:"edit"`
	// Lists hashtags, symbols, links, and mentions from tweet text.
	Entities EmbeddedTweetQuotedTweetQuotedTweetEntities `json:"entities"`
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
	LimitedActions []EmbeddedTweetQuotedTweetQuotedTweetLimitedAction `json:"limitedActions"`
	// Attached media items, omitted when unavailable.
	Media []TweetMedia `json:"media"`
	// Complete Note Tweet content and rich-text metadata.
	NoteTweet EmbeddedTweetQuotedTweetQuotedTweetNoteTweet `json:"noteTweet"`
	// Describes public place metadata on a geotagged tweet.
	Place             EmbeddedTweetQuotedTweetQuotedTweetPlace `json:"place"`
	PossiblySensitive bool                                     `json:"possiblySensitive"`
	// Public metadata whose fields are defined by X.
	PostCta map[string]any `json:"postCta"`
	// Engagement counts retained from a prior tweet edit.
	PreviousCounts EmbeddedTweetQuotedTweetQuotedTweetPreviousCounts `json:"previousCounts"`
	// Final nested tweet context at depth 4.
	QuotedTweet EmbeddedTweetQuotedTweetQuotedTweetQuotedTweet `json:"quoted_tweet"`
	// Quoted tweet ID.
	QuotedTweetID string `json:"quotedTweetId"`
	// Public post and user referenced by this reaction.
	ReactionContext EmbeddedTweetQuotedTweetQuotedTweetReactionContext `json:"reactionContext"`
	// Final nested tweet context at depth 4.
	RetweetedTweet EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweet `json:"retweeted_tweet"`
	// Public metadata whose fields are defined by X.
	Scopes map[string]any `json:"scopes"`
	Source string         `json:"source"`
	// Public visibility notice attached to an available tweet.
	Tombstone EmbeddedTweetQuotedTweetQuotedTweetTombstone `json:"tombstone"`
	Type      string                                       `json:"type"`
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
func (r EmbeddedTweetQuotedTweetQuotedTweet) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetQuotedTweetQuotedTweet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Describes an X Article preview and its lifecycle metadata.
type EmbeddedTweetQuotedTweetQuotedTweetArticle struct {
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
func (r EmbeddedTweetQuotedTweetQuotedTweetArticle) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetQuotedTweetQuotedTweetArticle) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Describes a public card and its referenced profiles.
type EmbeddedTweetQuotedTweetQuotedTweetCard struct {
	ID string `json:"id"`
	// Public metadata whose fields are defined by X.
	BindingValues map[string]any `json:"bindingValues"`
	Name          string         `json:"name"`
	// Public metadata whose fields are defined by X.
	Platform map[string]any `json:"platform"`
	URL      string         `json:"url"`
	// Unresolved card user references.
	UserReferenceErrors []EmbeddedTweetQuotedTweetQuotedTweetCardUserReferenceError `json:"userReferenceErrors"`
	UserReferences      []UserProfile                                               `json:"userReferences"`
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
func (r EmbeddedTweetQuotedTweetQuotedTweetCard) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetQuotedTweetQuotedTweetCard) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetQuotedTweetCardUserReferenceError struct {
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
func (r EmbeddedTweetQuotedTweetQuotedTweetCardUserReferenceError) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetQuotedTweetCardUserReferenceError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Community Note presentation metadata returned by X.
type EmbeddedTweetQuotedTweetQuotedTweetCommunityNote struct {
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
func (r EmbeddedTweetQuotedTweetQuotedTweetCommunityNote) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetQuotedTweetQuotedTweetCommunityNote) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Public reply policy and conversation owner.
type EmbeddedTweetQuotedTweetQuotedTweetConversationControl struct {
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
func (r EmbeddedTweetQuotedTweetQuotedTweetConversationControl) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetQuotedTweetQuotedTweetConversationControl) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Lists edit-chain identifiers and the remaining edit window.
type EmbeddedTweetQuotedTweetQuotedTweetEdit struct {
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
func (r EmbeddedTweetQuotedTweetQuotedTweetEdit) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetQuotedTweetQuotedTweetEdit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Lists hashtags, symbols, links, and mentions from tweet text.
type EmbeddedTweetQuotedTweetQuotedTweetEntities struct {
	Hashtags     []EmbeddedTweetQuotedTweetQuotedTweetEntitiesHashtag     `json:"hashtags"`
	Smarttags    []EmbeddedTweetQuotedTweetQuotedTweetEntitiesSmarttag    `json:"smarttags"`
	Symbols      []EmbeddedTweetQuotedTweetQuotedTweetEntitiesSymbol      `json:"symbols"`
	Timestamps   []EmbeddedTweetQuotedTweetQuotedTweetEntitiesTimestamp   `json:"timestamps"`
	URLs         []EmbeddedTweetQuotedTweetQuotedTweetEntitiesURL         `json:"urls"`
	UserMentions []EmbeddedTweetQuotedTweetQuotedTweetEntitiesUserMention `json:"user_mentions"`
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
func (r EmbeddedTweetQuotedTweetQuotedTweetEntities) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetQuotedTweetQuotedTweetEntities) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Provides hashtag text and source offsets within a tweet.
type EmbeddedTweetQuotedTweetQuotedTweetEntitiesHashtag struct {
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
func (r EmbeddedTweetQuotedTweetQuotedTweetEntitiesHashtag) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetQuotedTweetQuotedTweetEntitiesHashtag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indexed smart-tag, cashtag, or video timestamp metadata.
type EmbeddedTweetQuotedTweetQuotedTweetEntitiesSmarttag struct {
	Indices []int64                                                `json:"indices"`
	Seconds float64                                                `json:"seconds"`
	Tag     EmbeddedTweetQuotedTweetQuotedTweetEntitiesSmarttagTag `json:"tag"`
	Text    string                                                 `json:"text"`
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
func (r EmbeddedTweetQuotedTweetQuotedTweetEntitiesSmarttag) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetQuotedTweetQuotedTweetEntitiesSmarttag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetQuotedTweetEntitiesSmarttagTag struct {
	Info EmbeddedTweetQuotedTweetQuotedTweetEntitiesSmarttagTagInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetQuotedTweetQuotedTweetEntitiesSmarttagTag) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetQuotedTweetQuotedTweetEntitiesSmarttagTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetQuotedTweetEntitiesSmarttagTagInfo struct {
	Info EmbeddedTweetQuotedTweetQuotedTweetEntitiesSmarttagTagInfoInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetQuotedTweetQuotedTweetEntitiesSmarttagTagInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetQuotedTweetEntitiesSmarttagTagInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetQuotedTweetEntitiesSmarttagTagInfoInfo struct {
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
func (r EmbeddedTweetQuotedTweetQuotedTweetEntitiesSmarttagTagInfoInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetQuotedTweetEntitiesSmarttagTagInfoInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indexed smart-tag, cashtag, or video timestamp metadata.
type EmbeddedTweetQuotedTweetQuotedTweetEntitiesSymbol struct {
	Indices []int64                                              `json:"indices"`
	Seconds float64                                              `json:"seconds"`
	Tag     EmbeddedTweetQuotedTweetQuotedTweetEntitiesSymbolTag `json:"tag"`
	Text    string                                               `json:"text"`
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
func (r EmbeddedTweetQuotedTweetQuotedTweetEntitiesSymbol) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetQuotedTweetQuotedTweetEntitiesSymbol) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetQuotedTweetEntitiesSymbolTag struct {
	Info EmbeddedTweetQuotedTweetQuotedTweetEntitiesSymbolTagInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetQuotedTweetQuotedTweetEntitiesSymbolTag) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetQuotedTweetQuotedTweetEntitiesSymbolTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetQuotedTweetEntitiesSymbolTagInfo struct {
	Info EmbeddedTweetQuotedTweetQuotedTweetEntitiesSymbolTagInfoInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetQuotedTweetQuotedTweetEntitiesSymbolTagInfo) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetQuotedTweetQuotedTweetEntitiesSymbolTagInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetQuotedTweetEntitiesSymbolTagInfoInfo struct {
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
func (r EmbeddedTweetQuotedTweetQuotedTweetEntitiesSymbolTagInfoInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetQuotedTweetEntitiesSymbolTagInfoInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indexed smart-tag, cashtag, or video timestamp metadata.
type EmbeddedTweetQuotedTweetQuotedTweetEntitiesTimestamp struct {
	Indices []int64                                                 `json:"indices"`
	Seconds float64                                                 `json:"seconds"`
	Tag     EmbeddedTweetQuotedTweetQuotedTweetEntitiesTimestampTag `json:"tag"`
	Text    string                                                  `json:"text"`
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
func (r EmbeddedTweetQuotedTweetQuotedTweetEntitiesTimestamp) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetQuotedTweetQuotedTweetEntitiesTimestamp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetQuotedTweetEntitiesTimestampTag struct {
	Info EmbeddedTweetQuotedTweetQuotedTweetEntitiesTimestampTagInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetQuotedTweetQuotedTweetEntitiesTimestampTag) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetQuotedTweetQuotedTweetEntitiesTimestampTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetQuotedTweetEntitiesTimestampTagInfo struct {
	Info EmbeddedTweetQuotedTweetQuotedTweetEntitiesTimestampTagInfoInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetQuotedTweetQuotedTweetEntitiesTimestampTagInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetQuotedTweetEntitiesTimestampTagInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetQuotedTweetEntitiesTimestampTagInfoInfo struct {
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
func (r EmbeddedTweetQuotedTweetQuotedTweetEntitiesTimestampTagInfoInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetQuotedTweetEntitiesTimestampTagInfoInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Provides shortened, display, and expanded URLs from tweet text.
type EmbeddedTweetQuotedTweetQuotedTweetEntitiesURL struct {
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
func (r EmbeddedTweetQuotedTweetQuotedTweetEntitiesURL) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetQuotedTweetQuotedTweetEntitiesURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Provides profile identity and source offsets for a mention.
type EmbeddedTweetQuotedTweetQuotedTweetEntitiesUserMention struct {
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
func (r EmbeddedTweetQuotedTweetQuotedTweetEntitiesUserMention) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetQuotedTweetQuotedTweetEntitiesUserMention) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetQuotedTweetLimitedAction struct {
	Action string                                                 `json:"action"`
	Prompt EmbeddedTweetQuotedTweetQuotedTweetLimitedActionPrompt `json:"prompt"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Action      respjson.Field
		Prompt      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetQuotedTweetQuotedTweetLimitedAction) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetQuotedTweetQuotedTweetLimitedAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetQuotedTweetLimitedActionPrompt struct {
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
func (r EmbeddedTweetQuotedTweetQuotedTweetLimitedActionPrompt) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetQuotedTweetQuotedTweetLimitedActionPrompt) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Complete Note Tweet content and rich-text metadata.
type EmbeddedTweetQuotedTweetQuotedTweetNoteTweet struct {
	Text string `json:"text" api:"required"`
	ID   string `json:"id"`
	// Lists hashtags, symbols, links, and mentions from tweet text.
	Entities EmbeddedTweetQuotedTweetQuotedTweetNoteTweetEntities `json:"entities"`
	// Inline media positions in the Note Tweet text.
	InlineMedia  []EmbeddedTweetQuotedTweetQuotedTweetNoteTweetInlineMedia `json:"inlineMedia"`
	IsExpandable bool                                                      `json:"isExpandable"`
	RichtextTags []EmbeddedTweetQuotedTweetQuotedTweetNoteTweetRichtextTag `json:"richtextTags"`
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
func (r EmbeddedTweetQuotedTweetQuotedTweetNoteTweet) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetQuotedTweetQuotedTweetNoteTweet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Lists hashtags, symbols, links, and mentions from tweet text.
type EmbeddedTweetQuotedTweetQuotedTweetNoteTweetEntities struct {
	Hashtags     []EmbeddedTweetQuotedTweetQuotedTweetNoteTweetEntitiesHashtag     `json:"hashtags"`
	Smarttags    []EmbeddedTweetQuotedTweetQuotedTweetNoteTweetEntitiesSmarttag    `json:"smarttags"`
	Symbols      []EmbeddedTweetQuotedTweetQuotedTweetNoteTweetEntitiesSymbol      `json:"symbols"`
	Timestamps   []EmbeddedTweetQuotedTweetQuotedTweetNoteTweetEntitiesTimestamp   `json:"timestamps"`
	URLs         []EmbeddedTweetQuotedTweetQuotedTweetNoteTweetEntitiesURL         `json:"urls"`
	UserMentions []EmbeddedTweetQuotedTweetQuotedTweetNoteTweetEntitiesUserMention `json:"user_mentions"`
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
func (r EmbeddedTweetQuotedTweetQuotedTweetNoteTweetEntities) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetQuotedTweetQuotedTweetNoteTweetEntities) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Provides hashtag text and source offsets within a tweet.
type EmbeddedTweetQuotedTweetQuotedTweetNoteTweetEntitiesHashtag struct {
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
func (r EmbeddedTweetQuotedTweetQuotedTweetNoteTweetEntitiesHashtag) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetQuotedTweetNoteTweetEntitiesHashtag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indexed smart-tag, cashtag, or video timestamp metadata.
type EmbeddedTweetQuotedTweetQuotedTweetNoteTweetEntitiesSmarttag struct {
	Indices []int64                                                         `json:"indices"`
	Seconds float64                                                         `json:"seconds"`
	Tag     EmbeddedTweetQuotedTweetQuotedTweetNoteTweetEntitiesSmarttagTag `json:"tag"`
	Text    string                                                          `json:"text"`
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
func (r EmbeddedTweetQuotedTweetQuotedTweetNoteTweetEntitiesSmarttag) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetQuotedTweetNoteTweetEntitiesSmarttag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetQuotedTweetNoteTweetEntitiesSmarttagTag struct {
	Info EmbeddedTweetQuotedTweetQuotedTweetNoteTweetEntitiesSmarttagTagInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetQuotedTweetQuotedTweetNoteTweetEntitiesSmarttagTag) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetQuotedTweetNoteTweetEntitiesSmarttagTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetQuotedTweetNoteTweetEntitiesSmarttagTagInfo struct {
	Info EmbeddedTweetQuotedTweetQuotedTweetNoteTweetEntitiesSmarttagTagInfoInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetQuotedTweetQuotedTweetNoteTweetEntitiesSmarttagTagInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetQuotedTweetNoteTweetEntitiesSmarttagTagInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetQuotedTweetNoteTweetEntitiesSmarttagTagInfoInfo struct {
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
func (r EmbeddedTweetQuotedTweetQuotedTweetNoteTweetEntitiesSmarttagTagInfoInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetQuotedTweetNoteTweetEntitiesSmarttagTagInfoInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indexed smart-tag, cashtag, or video timestamp metadata.
type EmbeddedTweetQuotedTweetQuotedTweetNoteTweetEntitiesSymbol struct {
	Indices []int64                                                       `json:"indices"`
	Seconds float64                                                       `json:"seconds"`
	Tag     EmbeddedTweetQuotedTweetQuotedTweetNoteTweetEntitiesSymbolTag `json:"tag"`
	Text    string                                                        `json:"text"`
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
func (r EmbeddedTweetQuotedTweetQuotedTweetNoteTweetEntitiesSymbol) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetQuotedTweetNoteTweetEntitiesSymbol) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetQuotedTweetNoteTweetEntitiesSymbolTag struct {
	Info EmbeddedTweetQuotedTweetQuotedTweetNoteTweetEntitiesSymbolTagInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetQuotedTweetQuotedTweetNoteTweetEntitiesSymbolTag) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetQuotedTweetNoteTweetEntitiesSymbolTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetQuotedTweetNoteTweetEntitiesSymbolTagInfo struct {
	Info EmbeddedTweetQuotedTweetQuotedTweetNoteTweetEntitiesSymbolTagInfoInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetQuotedTweetQuotedTweetNoteTweetEntitiesSymbolTagInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetQuotedTweetNoteTweetEntitiesSymbolTagInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetQuotedTweetNoteTweetEntitiesSymbolTagInfoInfo struct {
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
func (r EmbeddedTweetQuotedTweetQuotedTweetNoteTweetEntitiesSymbolTagInfoInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetQuotedTweetNoteTweetEntitiesSymbolTagInfoInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indexed smart-tag, cashtag, or video timestamp metadata.
type EmbeddedTweetQuotedTweetQuotedTweetNoteTweetEntitiesTimestamp struct {
	Indices []int64                                                          `json:"indices"`
	Seconds float64                                                          `json:"seconds"`
	Tag     EmbeddedTweetQuotedTweetQuotedTweetNoteTweetEntitiesTimestampTag `json:"tag"`
	Text    string                                                           `json:"text"`
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
func (r EmbeddedTweetQuotedTweetQuotedTweetNoteTweetEntitiesTimestamp) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetQuotedTweetNoteTweetEntitiesTimestamp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetQuotedTweetNoteTweetEntitiesTimestampTag struct {
	Info EmbeddedTweetQuotedTweetQuotedTweetNoteTweetEntitiesTimestampTagInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetQuotedTweetQuotedTweetNoteTweetEntitiesTimestampTag) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetQuotedTweetNoteTweetEntitiesTimestampTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetQuotedTweetNoteTweetEntitiesTimestampTagInfo struct {
	Info EmbeddedTweetQuotedTweetQuotedTweetNoteTweetEntitiesTimestampTagInfoInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetQuotedTweetQuotedTweetNoteTweetEntitiesTimestampTagInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetQuotedTweetNoteTweetEntitiesTimestampTagInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetQuotedTweetNoteTweetEntitiesTimestampTagInfoInfo struct {
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
func (r EmbeddedTweetQuotedTweetQuotedTweetNoteTweetEntitiesTimestampTagInfoInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetQuotedTweetNoteTweetEntitiesTimestampTagInfoInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Provides shortened, display, and expanded URLs from tweet text.
type EmbeddedTweetQuotedTweetQuotedTweetNoteTweetEntitiesURL struct {
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
func (r EmbeddedTweetQuotedTweetQuotedTweetNoteTweetEntitiesURL) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetQuotedTweetQuotedTweetNoteTweetEntitiesURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Provides profile identity and source offsets for a mention.
type EmbeddedTweetQuotedTweetQuotedTweetNoteTweetEntitiesUserMention struct {
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
func (r EmbeddedTweetQuotedTweetQuotedTweetNoteTweetEntitiesUserMention) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetQuotedTweetNoteTweetEntitiesUserMention) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetQuotedTweetNoteTweetInlineMedia struct {
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
func (r EmbeddedTweetQuotedTweetQuotedTweetNoteTweetInlineMedia) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetQuotedTweetQuotedTweetNoteTweetInlineMedia) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetQuotedTweetNoteTweetRichtextTag struct {
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
func (r EmbeddedTweetQuotedTweetQuotedTweetNoteTweetRichtextTag) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetQuotedTweetQuotedTweetNoteTweetRichtextTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Describes public place metadata on a geotagged tweet.
type EmbeddedTweetQuotedTweetQuotedTweetPlace struct {
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
func (r EmbeddedTweetQuotedTweetQuotedTweetPlace) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetQuotedTweetQuotedTweetPlace) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Engagement counts retained from a prior tweet edit.
type EmbeddedTweetQuotedTweetQuotedTweetPreviousCounts struct {
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
func (r EmbeddedTweetQuotedTweetQuotedTweetPreviousCounts) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetQuotedTweetQuotedTweetPreviousCounts) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Final nested tweet context at depth 4.
type EmbeddedTweetQuotedTweetQuotedTweetQuotedTweet struct {
	ID            string `json:"id" api:"required"`
	BookmarkCount int64  `json:"bookmarkCount" api:"required"`
	LikeCount     int64  `json:"likeCount" api:"required"`
	QuoteCount    int64  `json:"quoteCount" api:"required"`
	ReplyCount    int64  `json:"replyCount" api:"required"`
	RetweetCount  int64  `json:"retweetCount" api:"required"`
	Text          string `json:"text" api:"required"`
	ViewCount     int64  `json:"viewCount" api:"required"`
	// Describes an X Article preview and its lifecycle metadata.
	Article EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetArticle `json:"article"`
	// X user profile with bio, follower counts, and verification status.
	Author UserProfile `json:"author"`
	// Describes a public card and its referenced profiles.
	Card EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetCard `json:"card"`
	// Community ID.
	CommunityID string `json:"communityId"`
	// Community Note presentation metadata returned by X.
	CommunityNote EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetCommunityNote `json:"communityNote"`
	// Content disclosure metadata shown by X when a tweet is labeled as paid
	// partnership content or AI-generated media.
	ContentDisclosure ContentDisclosure `json:"contentDisclosure"`
	// Public reply policy and conversation owner.
	ConversationControl EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetConversationControl `json:"conversationControl"`
	ConversationID      string                                                            `json:"conversationId"`
	CreatedAt           string                                                            `json:"createdAt"`
	DisplayTextRange    []int64                                                           `json:"displayTextRange"`
	// Lists edit-chain identifiers and the remaining edit window.
	Edit EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetEdit `json:"edit"`
	// Lists hashtags, symbols, links, and mentions from tweet text.
	Entities EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetEntities `json:"entities"`
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
	LimitedActions []EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetLimitedAction `json:"limitedActions"`
	// Attached media items, omitted when unavailable.
	Media []TweetMedia `json:"media"`
	// Complete Note Tweet content and rich-text metadata.
	NoteTweet EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetNoteTweet `json:"noteTweet"`
	// Describes public place metadata on a geotagged tweet.
	Place             EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetPlace `json:"place"`
	PossiblySensitive bool                                                `json:"possiblySensitive"`
	// Public metadata whose fields are defined by X.
	PostCta map[string]any `json:"postCta"`
	// Engagement counts retained from a prior tweet edit.
	PreviousCounts EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetPreviousCounts `json:"previousCounts"`
	// Quoted tweet ID.
	QuotedTweetID string `json:"quotedTweetId"`
	// Public post and user referenced by this reaction.
	ReactionContext EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetReactionContext `json:"reactionContext"`
	// Public metadata whose fields are defined by X.
	Scopes map[string]any `json:"scopes"`
	Source string         `json:"source"`
	// Public visibility notice attached to an available tweet.
	Tombstone EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetTombstone `json:"tombstone"`
	Type      string                                                  `json:"type"`
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
		QuotedTweetID       respjson.Field
		ReactionContext     respjson.Field
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
func (r EmbeddedTweetQuotedTweetQuotedTweetQuotedTweet) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetQuotedTweetQuotedTweetQuotedTweet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Describes an X Article preview and its lifecycle metadata.
type EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetArticle struct {
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
func (r EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetArticle) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetArticle) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Describes a public card and its referenced profiles.
type EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetCard struct {
	ID string `json:"id"`
	// Public metadata whose fields are defined by X.
	BindingValues map[string]any `json:"bindingValues"`
	Name          string         `json:"name"`
	// Public metadata whose fields are defined by X.
	Platform map[string]any `json:"platform"`
	URL      string         `json:"url"`
	// Unresolved card user references.
	UserReferenceErrors []EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetCardUserReferenceError `json:"userReferenceErrors"`
	UserReferences      []UserProfile                                                          `json:"userReferences"`
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
func (r EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetCard) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetCard) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetCardUserReferenceError struct {
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
func (r EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetCardUserReferenceError) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetCardUserReferenceError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Community Note presentation metadata returned by X.
type EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetCommunityNote struct {
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
func (r EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetCommunityNote) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetCommunityNote) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Public reply policy and conversation owner.
type EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetConversationControl struct {
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
func (r EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetConversationControl) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetConversationControl) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Lists edit-chain identifiers and the remaining edit window.
type EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetEdit struct {
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
func (r EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetEdit) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetEdit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Lists hashtags, symbols, links, and mentions from tweet text.
type EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetEntities struct {
	Hashtags     []EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetEntitiesHashtag     `json:"hashtags"`
	Smarttags    []EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetEntitiesSmarttag    `json:"smarttags"`
	Symbols      []EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetEntitiesSymbol      `json:"symbols"`
	Timestamps   []EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetEntitiesTimestamp   `json:"timestamps"`
	URLs         []EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetEntitiesURL         `json:"urls"`
	UserMentions []EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetEntitiesUserMention `json:"user_mentions"`
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
func (r EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetEntities) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetEntities) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Provides hashtag text and source offsets within a tweet.
type EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetEntitiesHashtag struct {
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
func (r EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetEntitiesHashtag) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetEntitiesHashtag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indexed smart-tag, cashtag, or video timestamp metadata.
type EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetEntitiesSmarttag struct {
	Indices []int64                                                           `json:"indices"`
	Seconds float64                                                           `json:"seconds"`
	Tag     EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetEntitiesSmarttagTag `json:"tag"`
	Text    string                                                            `json:"text"`
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
func (r EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetEntitiesSmarttag) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetEntitiesSmarttag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetEntitiesSmarttagTag struct {
	Info EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetEntitiesSmarttagTagInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetEntitiesSmarttagTag) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetEntitiesSmarttagTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetEntitiesSmarttagTagInfo struct {
	Info EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetEntitiesSmarttagTagInfoInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetEntitiesSmarttagTagInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetEntitiesSmarttagTagInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetEntitiesSmarttagTagInfoInfo struct {
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
func (r EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetEntitiesSmarttagTagInfoInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetEntitiesSmarttagTagInfoInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indexed smart-tag, cashtag, or video timestamp metadata.
type EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetEntitiesSymbol struct {
	Indices []int64                                                         `json:"indices"`
	Seconds float64                                                         `json:"seconds"`
	Tag     EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetEntitiesSymbolTag `json:"tag"`
	Text    string                                                          `json:"text"`
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
func (r EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetEntitiesSymbol) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetEntitiesSymbol) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetEntitiesSymbolTag struct {
	Info EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetEntitiesSymbolTagInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetEntitiesSymbolTag) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetEntitiesSymbolTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetEntitiesSymbolTagInfo struct {
	Info EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetEntitiesSymbolTagInfoInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetEntitiesSymbolTagInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetEntitiesSymbolTagInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetEntitiesSymbolTagInfoInfo struct {
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
func (r EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetEntitiesSymbolTagInfoInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetEntitiesSymbolTagInfoInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indexed smart-tag, cashtag, or video timestamp metadata.
type EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetEntitiesTimestamp struct {
	Indices []int64                                                            `json:"indices"`
	Seconds float64                                                            `json:"seconds"`
	Tag     EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetEntitiesTimestampTag `json:"tag"`
	Text    string                                                             `json:"text"`
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
func (r EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetEntitiesTimestamp) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetEntitiesTimestamp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetEntitiesTimestampTag struct {
	Info EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetEntitiesTimestampTagInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetEntitiesTimestampTag) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetEntitiesTimestampTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetEntitiesTimestampTagInfo struct {
	Info EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetEntitiesTimestampTagInfoInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetEntitiesTimestampTagInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetEntitiesTimestampTagInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetEntitiesTimestampTagInfoInfo struct {
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
func (r EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetEntitiesTimestampTagInfoInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetEntitiesTimestampTagInfoInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Provides shortened, display, and expanded URLs from tweet text.
type EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetEntitiesURL struct {
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
func (r EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetEntitiesURL) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetEntitiesURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Provides profile identity and source offsets for a mention.
type EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetEntitiesUserMention struct {
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
func (r EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetEntitiesUserMention) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetEntitiesUserMention) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetLimitedAction struct {
	Action string                                                            `json:"action"`
	Prompt EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetLimitedActionPrompt `json:"prompt"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Action      respjson.Field
		Prompt      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetLimitedAction) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetLimitedAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetLimitedActionPrompt struct {
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
func (r EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetLimitedActionPrompt) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetLimitedActionPrompt) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Complete Note Tweet content and rich-text metadata.
type EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetNoteTweet struct {
	Text string `json:"text" api:"required"`
	ID   string `json:"id"`
	// Lists hashtags, symbols, links, and mentions from tweet text.
	Entities EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetNoteTweetEntities `json:"entities"`
	// Inline media positions in the Note Tweet text.
	InlineMedia  []EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetNoteTweetInlineMedia `json:"inlineMedia"`
	IsExpandable bool                                                                 `json:"isExpandable"`
	RichtextTags []EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetNoteTweetRichtextTag `json:"richtextTags"`
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
func (r EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetNoteTweet) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetNoteTweet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Lists hashtags, symbols, links, and mentions from tweet text.
type EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetNoteTweetEntities struct {
	Hashtags     []EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetNoteTweetEntitiesHashtag     `json:"hashtags"`
	Smarttags    []EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetNoteTweetEntitiesSmarttag    `json:"smarttags"`
	Symbols      []EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetNoteTweetEntitiesSymbol      `json:"symbols"`
	Timestamps   []EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetNoteTweetEntitiesTimestamp   `json:"timestamps"`
	URLs         []EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetNoteTweetEntitiesURL         `json:"urls"`
	UserMentions []EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetNoteTweetEntitiesUserMention `json:"user_mentions"`
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
func (r EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetNoteTweetEntities) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetNoteTweetEntities) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Provides hashtag text and source offsets within a tweet.
type EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetNoteTweetEntitiesHashtag struct {
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
func (r EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetNoteTweetEntitiesHashtag) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetNoteTweetEntitiesHashtag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indexed smart-tag, cashtag, or video timestamp metadata.
type EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetNoteTweetEntitiesSmarttag struct {
	Indices []int64                                                                    `json:"indices"`
	Seconds float64                                                                    `json:"seconds"`
	Tag     EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetNoteTweetEntitiesSmarttagTag `json:"tag"`
	Text    string                                                                     `json:"text"`
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
func (r EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetNoteTweetEntitiesSmarttag) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetNoteTweetEntitiesSmarttag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetNoteTweetEntitiesSmarttagTag struct {
	Info EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetNoteTweetEntitiesSmarttagTagInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetNoteTweetEntitiesSmarttagTag) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetNoteTweetEntitiesSmarttagTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetNoteTweetEntitiesSmarttagTagInfo struct {
	Info EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetNoteTweetEntitiesSmarttagTagInfoInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetNoteTweetEntitiesSmarttagTagInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetNoteTweetEntitiesSmarttagTagInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetNoteTweetEntitiesSmarttagTagInfoInfo struct {
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
func (r EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetNoteTweetEntitiesSmarttagTagInfoInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetNoteTweetEntitiesSmarttagTagInfoInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indexed smart-tag, cashtag, or video timestamp metadata.
type EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetNoteTweetEntitiesSymbol struct {
	Indices []int64                                                                  `json:"indices"`
	Seconds float64                                                                  `json:"seconds"`
	Tag     EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetNoteTweetEntitiesSymbolTag `json:"tag"`
	Text    string                                                                   `json:"text"`
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
func (r EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetNoteTweetEntitiesSymbol) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetNoteTweetEntitiesSymbol) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetNoteTweetEntitiesSymbolTag struct {
	Info EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetNoteTweetEntitiesSymbolTagInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetNoteTweetEntitiesSymbolTag) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetNoteTweetEntitiesSymbolTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetNoteTweetEntitiesSymbolTagInfo struct {
	Info EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetNoteTweetEntitiesSymbolTagInfoInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetNoteTweetEntitiesSymbolTagInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetNoteTweetEntitiesSymbolTagInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetNoteTweetEntitiesSymbolTagInfoInfo struct {
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
func (r EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetNoteTweetEntitiesSymbolTagInfoInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetNoteTweetEntitiesSymbolTagInfoInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indexed smart-tag, cashtag, or video timestamp metadata.
type EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetNoteTweetEntitiesTimestamp struct {
	Indices []int64                                                                     `json:"indices"`
	Seconds float64                                                                     `json:"seconds"`
	Tag     EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetNoteTweetEntitiesTimestampTag `json:"tag"`
	Text    string                                                                      `json:"text"`
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
func (r EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetNoteTweetEntitiesTimestamp) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetNoteTweetEntitiesTimestamp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetNoteTweetEntitiesTimestampTag struct {
	Info EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetNoteTweetEntitiesTimestampTagInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetNoteTweetEntitiesTimestampTag) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetNoteTweetEntitiesTimestampTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetNoteTweetEntitiesTimestampTagInfo struct {
	Info EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetNoteTweetEntitiesTimestampTagInfoInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetNoteTweetEntitiesTimestampTagInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetNoteTweetEntitiesTimestampTagInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetNoteTweetEntitiesTimestampTagInfoInfo struct {
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
func (r EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetNoteTweetEntitiesTimestampTagInfoInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetNoteTweetEntitiesTimestampTagInfoInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Provides shortened, display, and expanded URLs from tweet text.
type EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetNoteTweetEntitiesURL struct {
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
func (r EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetNoteTweetEntitiesURL) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetNoteTweetEntitiesURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Provides profile identity and source offsets for a mention.
type EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetNoteTweetEntitiesUserMention struct {
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
func (r EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetNoteTweetEntitiesUserMention) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetNoteTweetEntitiesUserMention) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetNoteTweetInlineMedia struct {
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
func (r EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetNoteTweetInlineMedia) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetNoteTweetInlineMedia) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetNoteTweetRichtextTag struct {
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
func (r EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetNoteTweetRichtextTag) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetNoteTweetRichtextTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Describes public place metadata on a geotagged tweet.
type EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetPlace struct {
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
func (r EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetPlace) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetPlace) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Engagement counts retained from a prior tweet edit.
type EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetPreviousCounts struct {
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
func (r EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetPreviousCounts) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetPreviousCounts) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Public post and user referenced by this reaction.
type EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetReactionContext struct {
	// Referenced post ID.
	ReactedToPostID string `json:"reactedToPostId"`
	// X user profile with bio, follower counts, and verification status.
	ReactedToUser UserProfile `json:"reactedToUser"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ReactedToPostID respjson.Field
		ReactedToUser   respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetReactionContext) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetReactionContext) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Public visibility notice attached to an available tweet.
type EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetTombstone struct {
	Text EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetTombstoneText `json:"text"`
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
func (r EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetTombstone) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetTombstone) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetTombstoneText struct {
	Entities []EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetTombstoneTextEntity `json:"entities"`
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
func (r EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetTombstoneText) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetTombstoneText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetTombstoneTextEntity struct {
	FromIndex int64                                                                `json:"fromIndex"`
	Ref       EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetTombstoneTextEntityRef `json:"ref"`
	ToIndex   int64                                                                `json:"toIndex"`
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
func (r EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetTombstoneTextEntity) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetTombstoneTextEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetTombstoneTextEntityRef struct {
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
func (r EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetTombstoneTextEntityRef) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetTombstoneTextEntityRef) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Public post and user referenced by this reaction.
type EmbeddedTweetQuotedTweetQuotedTweetReactionContext struct {
	// Referenced post ID.
	ReactedToPostID string `json:"reactedToPostId"`
	// X user profile with bio, follower counts, and verification status.
	ReactedToUser UserProfile `json:"reactedToUser"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ReactedToPostID respjson.Field
		ReactedToUser   respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetQuotedTweetQuotedTweetReactionContext) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetQuotedTweetQuotedTweetReactionContext) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Final nested tweet context at depth 4.
type EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweet struct {
	ID            string `json:"id" api:"required"`
	BookmarkCount int64  `json:"bookmarkCount" api:"required"`
	LikeCount     int64  `json:"likeCount" api:"required"`
	QuoteCount    int64  `json:"quoteCount" api:"required"`
	ReplyCount    int64  `json:"replyCount" api:"required"`
	RetweetCount  int64  `json:"retweetCount" api:"required"`
	Text          string `json:"text" api:"required"`
	ViewCount     int64  `json:"viewCount" api:"required"`
	// Describes an X Article preview and its lifecycle metadata.
	Article EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetArticle `json:"article"`
	// X user profile with bio, follower counts, and verification status.
	Author UserProfile `json:"author"`
	// Describes a public card and its referenced profiles.
	Card EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetCard `json:"card"`
	// Community ID.
	CommunityID string `json:"communityId"`
	// Community Note presentation metadata returned by X.
	CommunityNote EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetCommunityNote `json:"communityNote"`
	// Content disclosure metadata shown by X when a tweet is labeled as paid
	// partnership content or AI-generated media.
	ContentDisclosure ContentDisclosure `json:"contentDisclosure"`
	// Public reply policy and conversation owner.
	ConversationControl EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetConversationControl `json:"conversationControl"`
	ConversationID      string                                                               `json:"conversationId"`
	CreatedAt           string                                                               `json:"createdAt"`
	DisplayTextRange    []int64                                                              `json:"displayTextRange"`
	// Lists edit-chain identifiers and the remaining edit window.
	Edit EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetEdit `json:"edit"`
	// Lists hashtags, symbols, links, and mentions from tweet text.
	Entities EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetEntities `json:"entities"`
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
	LimitedActions []EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetLimitedAction `json:"limitedActions"`
	// Attached media items, omitted when unavailable.
	Media []TweetMedia `json:"media"`
	// Complete Note Tweet content and rich-text metadata.
	NoteTweet EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetNoteTweet `json:"noteTweet"`
	// Describes public place metadata on a geotagged tweet.
	Place             EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetPlace `json:"place"`
	PossiblySensitive bool                                                   `json:"possiblySensitive"`
	// Public metadata whose fields are defined by X.
	PostCta map[string]any `json:"postCta"`
	// Engagement counts retained from a prior tweet edit.
	PreviousCounts EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetPreviousCounts `json:"previousCounts"`
	// Quoted tweet ID.
	QuotedTweetID string `json:"quotedTweetId"`
	// Public post and user referenced by this reaction.
	ReactionContext EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetReactionContext `json:"reactionContext"`
	// Public metadata whose fields are defined by X.
	Scopes map[string]any `json:"scopes"`
	Source string         `json:"source"`
	// Public visibility notice attached to an available tweet.
	Tombstone EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetTombstone `json:"tombstone"`
	Type      string                                                     `json:"type"`
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
		QuotedTweetID       respjson.Field
		ReactionContext     respjson.Field
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
func (r EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweet) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Describes an X Article preview and its lifecycle metadata.
type EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetArticle struct {
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
func (r EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetArticle) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetArticle) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Describes a public card and its referenced profiles.
type EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetCard struct {
	ID string `json:"id"`
	// Public metadata whose fields are defined by X.
	BindingValues map[string]any `json:"bindingValues"`
	Name          string         `json:"name"`
	// Public metadata whose fields are defined by X.
	Platform map[string]any `json:"platform"`
	URL      string         `json:"url"`
	// Unresolved card user references.
	UserReferenceErrors []EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetCardUserReferenceError `json:"userReferenceErrors"`
	UserReferences      []UserProfile                                                             `json:"userReferences"`
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
func (r EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetCard) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetCard) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetCardUserReferenceError struct {
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
func (r EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetCardUserReferenceError) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetCardUserReferenceError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Community Note presentation metadata returned by X.
type EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetCommunityNote struct {
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
func (r EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetCommunityNote) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetCommunityNote) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Public reply policy and conversation owner.
type EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetConversationControl struct {
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
func (r EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetConversationControl) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetConversationControl) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Lists edit-chain identifiers and the remaining edit window.
type EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetEdit struct {
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
func (r EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetEdit) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetEdit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Lists hashtags, symbols, links, and mentions from tweet text.
type EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetEntities struct {
	Hashtags     []EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetEntitiesHashtag     `json:"hashtags"`
	Smarttags    []EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetEntitiesSmarttag    `json:"smarttags"`
	Symbols      []EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetEntitiesSymbol      `json:"symbols"`
	Timestamps   []EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetEntitiesTimestamp   `json:"timestamps"`
	URLs         []EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetEntitiesURL         `json:"urls"`
	UserMentions []EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetEntitiesUserMention `json:"user_mentions"`
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
func (r EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetEntities) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetEntities) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Provides hashtag text and source offsets within a tweet.
type EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetEntitiesHashtag struct {
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
func (r EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetEntitiesHashtag) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetEntitiesHashtag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indexed smart-tag, cashtag, or video timestamp metadata.
type EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetEntitiesSmarttag struct {
	Indices []int64                                                              `json:"indices"`
	Seconds float64                                                              `json:"seconds"`
	Tag     EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetEntitiesSmarttagTag `json:"tag"`
	Text    string                                                               `json:"text"`
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
func (r EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetEntitiesSmarttag) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetEntitiesSmarttag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetEntitiesSmarttagTag struct {
	Info EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetEntitiesSmarttagTagInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetEntitiesSmarttagTag) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetEntitiesSmarttagTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetEntitiesSmarttagTagInfo struct {
	Info EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetEntitiesSmarttagTagInfoInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetEntitiesSmarttagTagInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetEntitiesSmarttagTagInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetEntitiesSmarttagTagInfoInfo struct {
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
func (r EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetEntitiesSmarttagTagInfoInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetEntitiesSmarttagTagInfoInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indexed smart-tag, cashtag, or video timestamp metadata.
type EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetEntitiesSymbol struct {
	Indices []int64                                                            `json:"indices"`
	Seconds float64                                                            `json:"seconds"`
	Tag     EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetEntitiesSymbolTag `json:"tag"`
	Text    string                                                             `json:"text"`
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
func (r EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetEntitiesSymbol) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetEntitiesSymbol) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetEntitiesSymbolTag struct {
	Info EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetEntitiesSymbolTagInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetEntitiesSymbolTag) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetEntitiesSymbolTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetEntitiesSymbolTagInfo struct {
	Info EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetEntitiesSymbolTagInfoInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetEntitiesSymbolTagInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetEntitiesSymbolTagInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetEntitiesSymbolTagInfoInfo struct {
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
func (r EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetEntitiesSymbolTagInfoInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetEntitiesSymbolTagInfoInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indexed smart-tag, cashtag, or video timestamp metadata.
type EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetEntitiesTimestamp struct {
	Indices []int64                                                               `json:"indices"`
	Seconds float64                                                               `json:"seconds"`
	Tag     EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetEntitiesTimestampTag `json:"tag"`
	Text    string                                                                `json:"text"`
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
func (r EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetEntitiesTimestamp) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetEntitiesTimestamp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetEntitiesTimestampTag struct {
	Info EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetEntitiesTimestampTagInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetEntitiesTimestampTag) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetEntitiesTimestampTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetEntitiesTimestampTagInfo struct {
	Info EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetEntitiesTimestampTagInfoInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetEntitiesTimestampTagInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetEntitiesTimestampTagInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetEntitiesTimestampTagInfoInfo struct {
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
func (r EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetEntitiesTimestampTagInfoInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetEntitiesTimestampTagInfoInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Provides shortened, display, and expanded URLs from tweet text.
type EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetEntitiesURL struct {
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
func (r EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetEntitiesURL) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetEntitiesURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Provides profile identity and source offsets for a mention.
type EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetEntitiesUserMention struct {
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
func (r EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetEntitiesUserMention) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetEntitiesUserMention) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetLimitedAction struct {
	Action string                                                               `json:"action"`
	Prompt EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetLimitedActionPrompt `json:"prompt"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Action      respjson.Field
		Prompt      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetLimitedAction) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetLimitedAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetLimitedActionPrompt struct {
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
func (r EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetLimitedActionPrompt) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetLimitedActionPrompt) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Complete Note Tweet content and rich-text metadata.
type EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetNoteTweet struct {
	Text string `json:"text" api:"required"`
	ID   string `json:"id"`
	// Lists hashtags, symbols, links, and mentions from tweet text.
	Entities EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetNoteTweetEntities `json:"entities"`
	// Inline media positions in the Note Tweet text.
	InlineMedia  []EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetNoteTweetInlineMedia `json:"inlineMedia"`
	IsExpandable bool                                                                    `json:"isExpandable"`
	RichtextTags []EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetNoteTweetRichtextTag `json:"richtextTags"`
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
func (r EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetNoteTweet) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetNoteTweet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Lists hashtags, symbols, links, and mentions from tweet text.
type EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetNoteTweetEntities struct {
	Hashtags     []EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesHashtag     `json:"hashtags"`
	Smarttags    []EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesSmarttag    `json:"smarttags"`
	Symbols      []EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesSymbol      `json:"symbols"`
	Timestamps   []EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesTimestamp   `json:"timestamps"`
	URLs         []EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesURL         `json:"urls"`
	UserMentions []EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesUserMention `json:"user_mentions"`
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
func (r EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetNoteTweetEntities) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetNoteTweetEntities) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Provides hashtag text and source offsets within a tweet.
type EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesHashtag struct {
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
func (r EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesHashtag) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesHashtag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indexed smart-tag, cashtag, or video timestamp metadata.
type EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesSmarttag struct {
	Indices []int64                                                                       `json:"indices"`
	Seconds float64                                                                       `json:"seconds"`
	Tag     EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesSmarttagTag `json:"tag"`
	Text    string                                                                        `json:"text"`
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
func (r EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesSmarttag) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesSmarttag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesSmarttagTag struct {
	Info EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesSmarttagTagInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesSmarttagTag) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesSmarttagTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesSmarttagTagInfo struct {
	Info EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesSmarttagTagInfoInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesSmarttagTagInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesSmarttagTagInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesSmarttagTagInfoInfo struct {
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
func (r EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesSmarttagTagInfoInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesSmarttagTagInfoInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indexed smart-tag, cashtag, or video timestamp metadata.
type EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesSymbol struct {
	Indices []int64                                                                     `json:"indices"`
	Seconds float64                                                                     `json:"seconds"`
	Tag     EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesSymbolTag `json:"tag"`
	Text    string                                                                      `json:"text"`
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
func (r EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesSymbol) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesSymbol) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesSymbolTag struct {
	Info EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesSymbolTagInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesSymbolTag) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesSymbolTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesSymbolTagInfo struct {
	Info EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesSymbolTagInfoInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesSymbolTagInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesSymbolTagInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesSymbolTagInfoInfo struct {
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
func (r EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesSymbolTagInfoInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesSymbolTagInfoInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indexed smart-tag, cashtag, or video timestamp metadata.
type EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesTimestamp struct {
	Indices []int64                                                                        `json:"indices"`
	Seconds float64                                                                        `json:"seconds"`
	Tag     EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesTimestampTag `json:"tag"`
	Text    string                                                                         `json:"text"`
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
func (r EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesTimestamp) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesTimestamp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesTimestampTag struct {
	Info EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesTimestampTagInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesTimestampTag) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesTimestampTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesTimestampTagInfo struct {
	Info EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesTimestampTagInfoInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesTimestampTagInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesTimestampTagInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesTimestampTagInfoInfo struct {
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
func (r EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesTimestampTagInfoInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesTimestampTagInfoInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Provides shortened, display, and expanded URLs from tweet text.
type EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesURL struct {
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
func (r EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesURL) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Provides profile identity and source offsets for a mention.
type EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesUserMention struct {
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
func (r EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesUserMention) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesUserMention) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetNoteTweetInlineMedia struct {
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
func (r EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetNoteTweetInlineMedia) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetNoteTweetInlineMedia) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetNoteTweetRichtextTag struct {
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
func (r EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetNoteTweetRichtextTag) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetNoteTweetRichtextTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Describes public place metadata on a geotagged tweet.
type EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetPlace struct {
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
func (r EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetPlace) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetPlace) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Engagement counts retained from a prior tweet edit.
type EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetPreviousCounts struct {
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
func (r EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetPreviousCounts) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetPreviousCounts) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Public post and user referenced by this reaction.
type EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetReactionContext struct {
	// Referenced post ID.
	ReactedToPostID string `json:"reactedToPostId"`
	// X user profile with bio, follower counts, and verification status.
	ReactedToUser UserProfile `json:"reactedToUser"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ReactedToPostID respjson.Field
		ReactedToUser   respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetReactionContext) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetReactionContext) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Public visibility notice attached to an available tweet.
type EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetTombstone struct {
	Text EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetTombstoneText `json:"text"`
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
func (r EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetTombstone) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetTombstone) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetTombstoneText struct {
	Entities []EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetTombstoneTextEntity `json:"entities"`
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
func (r EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetTombstoneText) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetTombstoneText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetTombstoneTextEntity struct {
	FromIndex int64                                                                   `json:"fromIndex"`
	Ref       EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetTombstoneTextEntityRef `json:"ref"`
	ToIndex   int64                                                                   `json:"toIndex"`
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
func (r EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetTombstoneTextEntity) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetTombstoneTextEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetTombstoneTextEntityRef struct {
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
func (r EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetTombstoneTextEntityRef) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetTombstoneTextEntityRef) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Public visibility notice attached to an available tweet.
type EmbeddedTweetQuotedTweetQuotedTweetTombstone struct {
	Text EmbeddedTweetQuotedTweetQuotedTweetTombstoneText `json:"text"`
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
func (r EmbeddedTweetQuotedTweetQuotedTweetTombstone) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetQuotedTweetQuotedTweetTombstone) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetQuotedTweetTombstoneText struct {
	Entities []EmbeddedTweetQuotedTweetQuotedTweetTombstoneTextEntity `json:"entities"`
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
func (r EmbeddedTweetQuotedTweetQuotedTweetTombstoneText) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetQuotedTweetQuotedTweetTombstoneText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetQuotedTweetTombstoneTextEntity struct {
	FromIndex int64                                                     `json:"fromIndex"`
	Ref       EmbeddedTweetQuotedTweetQuotedTweetTombstoneTextEntityRef `json:"ref"`
	ToIndex   int64                                                     `json:"toIndex"`
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
func (r EmbeddedTweetQuotedTweetQuotedTweetTombstoneTextEntity) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetQuotedTweetQuotedTweetTombstoneTextEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetQuotedTweetTombstoneTextEntityRef struct {
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
func (r EmbeddedTweetQuotedTweetQuotedTweetTombstoneTextEntityRef) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetQuotedTweetTombstoneTextEntityRef) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Public post and user referenced by this reaction.
type EmbeddedTweetQuotedTweetReactionContext struct {
	// Referenced post ID.
	ReactedToPostID string `json:"reactedToPostId"`
	// X user profile with bio, follower counts, and verification status.
	ReactedToUser UserProfile `json:"reactedToUser"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ReactedToPostID respjson.Field
		ReactedToUser   respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetQuotedTweetReactionContext) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetQuotedTweetReactionContext) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Nested tweet context at depth 3.
type EmbeddedTweetQuotedTweetRetweetedTweet struct {
	ID            string `json:"id" api:"required"`
	BookmarkCount int64  `json:"bookmarkCount" api:"required"`
	LikeCount     int64  `json:"likeCount" api:"required"`
	QuoteCount    int64  `json:"quoteCount" api:"required"`
	ReplyCount    int64  `json:"replyCount" api:"required"`
	RetweetCount  int64  `json:"retweetCount" api:"required"`
	Text          string `json:"text" api:"required"`
	ViewCount     int64  `json:"viewCount" api:"required"`
	// Describes an X Article preview and its lifecycle metadata.
	Article EmbeddedTweetQuotedTweetRetweetedTweetArticle `json:"article"`
	// X user profile with bio, follower counts, and verification status.
	Author UserProfile `json:"author"`
	// Describes a public card and its referenced profiles.
	Card EmbeddedTweetQuotedTweetRetweetedTweetCard `json:"card"`
	// Community ID.
	CommunityID string `json:"communityId"`
	// Community Note presentation metadata returned by X.
	CommunityNote EmbeddedTweetQuotedTweetRetweetedTweetCommunityNote `json:"communityNote"`
	// Content disclosure metadata shown by X when a tweet is labeled as paid
	// partnership content or AI-generated media.
	ContentDisclosure ContentDisclosure `json:"contentDisclosure"`
	// Public reply policy and conversation owner.
	ConversationControl EmbeddedTweetQuotedTweetRetweetedTweetConversationControl `json:"conversationControl"`
	ConversationID      string                                                    `json:"conversationId"`
	CreatedAt           string                                                    `json:"createdAt"`
	DisplayTextRange    []int64                                                   `json:"displayTextRange"`
	// Lists edit-chain identifiers and the remaining edit window.
	Edit EmbeddedTweetQuotedTweetRetweetedTweetEdit `json:"edit"`
	// Lists hashtags, symbols, links, and mentions from tweet text.
	Entities EmbeddedTweetQuotedTweetRetweetedTweetEntities `json:"entities"`
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
	LimitedActions []EmbeddedTweetQuotedTweetRetweetedTweetLimitedAction `json:"limitedActions"`
	// Attached media items, omitted when unavailable.
	Media []TweetMedia `json:"media"`
	// Complete Note Tweet content and rich-text metadata.
	NoteTweet EmbeddedTweetQuotedTweetRetweetedTweetNoteTweet `json:"noteTweet"`
	// Describes public place metadata on a geotagged tweet.
	Place             EmbeddedTweetQuotedTweetRetweetedTweetPlace `json:"place"`
	PossiblySensitive bool                                        `json:"possiblySensitive"`
	// Public metadata whose fields are defined by X.
	PostCta map[string]any `json:"postCta"`
	// Engagement counts retained from a prior tweet edit.
	PreviousCounts EmbeddedTweetQuotedTweetRetweetedTweetPreviousCounts `json:"previousCounts"`
	// Final nested tweet context at depth 4.
	QuotedTweet EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweet `json:"quoted_tweet"`
	// Quoted tweet ID.
	QuotedTweetID string `json:"quotedTweetId"`
	// Public post and user referenced by this reaction.
	ReactionContext EmbeddedTweetQuotedTweetRetweetedTweetReactionContext `json:"reactionContext"`
	// Final nested tweet context at depth 4.
	RetweetedTweet EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweet `json:"retweeted_tweet"`
	// Public metadata whose fields are defined by X.
	Scopes map[string]any `json:"scopes"`
	Source string         `json:"source"`
	// Public visibility notice attached to an available tweet.
	Tombstone EmbeddedTweetQuotedTweetRetweetedTweetTombstone `json:"tombstone"`
	Type      string                                          `json:"type"`
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
func (r EmbeddedTweetQuotedTweetRetweetedTweet) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetQuotedTweetRetweetedTweet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Describes an X Article preview and its lifecycle metadata.
type EmbeddedTweetQuotedTweetRetweetedTweetArticle struct {
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
func (r EmbeddedTweetQuotedTweetRetweetedTweetArticle) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetQuotedTweetRetweetedTweetArticle) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Describes a public card and its referenced profiles.
type EmbeddedTweetQuotedTweetRetweetedTweetCard struct {
	ID string `json:"id"`
	// Public metadata whose fields are defined by X.
	BindingValues map[string]any `json:"bindingValues"`
	Name          string         `json:"name"`
	// Public metadata whose fields are defined by X.
	Platform map[string]any `json:"platform"`
	URL      string         `json:"url"`
	// Unresolved card user references.
	UserReferenceErrors []EmbeddedTweetQuotedTweetRetweetedTweetCardUserReferenceError `json:"userReferenceErrors"`
	UserReferences      []UserProfile                                                  `json:"userReferences"`
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
func (r EmbeddedTweetQuotedTweetRetweetedTweetCard) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetQuotedTweetRetweetedTweetCard) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetRetweetedTweetCardUserReferenceError struct {
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
func (r EmbeddedTweetQuotedTweetRetweetedTweetCardUserReferenceError) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetRetweetedTweetCardUserReferenceError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Community Note presentation metadata returned by X.
type EmbeddedTweetQuotedTweetRetweetedTweetCommunityNote struct {
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
func (r EmbeddedTweetQuotedTweetRetweetedTweetCommunityNote) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetQuotedTweetRetweetedTweetCommunityNote) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Public reply policy and conversation owner.
type EmbeddedTweetQuotedTweetRetweetedTweetConversationControl struct {
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
func (r EmbeddedTweetQuotedTweetRetweetedTweetConversationControl) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetRetweetedTweetConversationControl) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Lists edit-chain identifiers and the remaining edit window.
type EmbeddedTweetQuotedTweetRetweetedTweetEdit struct {
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
func (r EmbeddedTweetQuotedTweetRetweetedTweetEdit) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetQuotedTweetRetweetedTweetEdit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Lists hashtags, symbols, links, and mentions from tweet text.
type EmbeddedTweetQuotedTweetRetweetedTweetEntities struct {
	Hashtags     []EmbeddedTweetQuotedTweetRetweetedTweetEntitiesHashtag     `json:"hashtags"`
	Smarttags    []EmbeddedTweetQuotedTweetRetweetedTweetEntitiesSmarttag    `json:"smarttags"`
	Symbols      []EmbeddedTweetQuotedTweetRetweetedTweetEntitiesSymbol      `json:"symbols"`
	Timestamps   []EmbeddedTweetQuotedTweetRetweetedTweetEntitiesTimestamp   `json:"timestamps"`
	URLs         []EmbeddedTweetQuotedTweetRetweetedTweetEntitiesURL         `json:"urls"`
	UserMentions []EmbeddedTweetQuotedTweetRetweetedTweetEntitiesUserMention `json:"user_mentions"`
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
func (r EmbeddedTweetQuotedTweetRetweetedTweetEntities) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetQuotedTweetRetweetedTweetEntities) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Provides hashtag text and source offsets within a tweet.
type EmbeddedTweetQuotedTweetRetweetedTweetEntitiesHashtag struct {
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
func (r EmbeddedTweetQuotedTweetRetweetedTweetEntitiesHashtag) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetQuotedTweetRetweetedTweetEntitiesHashtag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indexed smart-tag, cashtag, or video timestamp metadata.
type EmbeddedTweetQuotedTweetRetweetedTweetEntitiesSmarttag struct {
	Indices []int64                                                   `json:"indices"`
	Seconds float64                                                   `json:"seconds"`
	Tag     EmbeddedTweetQuotedTweetRetweetedTweetEntitiesSmarttagTag `json:"tag"`
	Text    string                                                    `json:"text"`
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
func (r EmbeddedTweetQuotedTweetRetweetedTweetEntitiesSmarttag) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetQuotedTweetRetweetedTweetEntitiesSmarttag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetRetweetedTweetEntitiesSmarttagTag struct {
	Info EmbeddedTweetQuotedTweetRetweetedTweetEntitiesSmarttagTagInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetQuotedTweetRetweetedTweetEntitiesSmarttagTag) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetRetweetedTweetEntitiesSmarttagTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetRetweetedTweetEntitiesSmarttagTagInfo struct {
	Info EmbeddedTweetQuotedTweetRetweetedTweetEntitiesSmarttagTagInfoInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetQuotedTweetRetweetedTweetEntitiesSmarttagTagInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetRetweetedTweetEntitiesSmarttagTagInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetRetweetedTweetEntitiesSmarttagTagInfoInfo struct {
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
func (r EmbeddedTweetQuotedTweetRetweetedTweetEntitiesSmarttagTagInfoInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetRetweetedTweetEntitiesSmarttagTagInfoInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indexed smart-tag, cashtag, or video timestamp metadata.
type EmbeddedTweetQuotedTweetRetweetedTweetEntitiesSymbol struct {
	Indices []int64                                                 `json:"indices"`
	Seconds float64                                                 `json:"seconds"`
	Tag     EmbeddedTweetQuotedTweetRetweetedTweetEntitiesSymbolTag `json:"tag"`
	Text    string                                                  `json:"text"`
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
func (r EmbeddedTweetQuotedTweetRetweetedTweetEntitiesSymbol) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetQuotedTweetRetweetedTweetEntitiesSymbol) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetRetweetedTweetEntitiesSymbolTag struct {
	Info EmbeddedTweetQuotedTweetRetweetedTweetEntitiesSymbolTagInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetQuotedTweetRetweetedTweetEntitiesSymbolTag) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetQuotedTweetRetweetedTweetEntitiesSymbolTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetRetweetedTweetEntitiesSymbolTagInfo struct {
	Info EmbeddedTweetQuotedTweetRetweetedTweetEntitiesSymbolTagInfoInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetQuotedTweetRetweetedTweetEntitiesSymbolTagInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetRetweetedTweetEntitiesSymbolTagInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetRetweetedTweetEntitiesSymbolTagInfoInfo struct {
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
func (r EmbeddedTweetQuotedTweetRetweetedTweetEntitiesSymbolTagInfoInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetRetweetedTweetEntitiesSymbolTagInfoInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indexed smart-tag, cashtag, or video timestamp metadata.
type EmbeddedTweetQuotedTweetRetweetedTweetEntitiesTimestamp struct {
	Indices []int64                                                    `json:"indices"`
	Seconds float64                                                    `json:"seconds"`
	Tag     EmbeddedTweetQuotedTweetRetweetedTweetEntitiesTimestampTag `json:"tag"`
	Text    string                                                     `json:"text"`
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
func (r EmbeddedTweetQuotedTweetRetweetedTweetEntitiesTimestamp) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetQuotedTweetRetweetedTweetEntitiesTimestamp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetRetweetedTweetEntitiesTimestampTag struct {
	Info EmbeddedTweetQuotedTweetRetweetedTweetEntitiesTimestampTagInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetQuotedTweetRetweetedTweetEntitiesTimestampTag) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetRetweetedTweetEntitiesTimestampTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetRetweetedTweetEntitiesTimestampTagInfo struct {
	Info EmbeddedTweetQuotedTweetRetweetedTweetEntitiesTimestampTagInfoInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetQuotedTweetRetweetedTweetEntitiesTimestampTagInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetRetweetedTweetEntitiesTimestampTagInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetRetweetedTweetEntitiesTimestampTagInfoInfo struct {
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
func (r EmbeddedTweetQuotedTweetRetweetedTweetEntitiesTimestampTagInfoInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetRetweetedTweetEntitiesTimestampTagInfoInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Provides shortened, display, and expanded URLs from tweet text.
type EmbeddedTweetQuotedTweetRetweetedTweetEntitiesURL struct {
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
func (r EmbeddedTweetQuotedTweetRetweetedTweetEntitiesURL) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetQuotedTweetRetweetedTweetEntitiesURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Provides profile identity and source offsets for a mention.
type EmbeddedTweetQuotedTweetRetweetedTweetEntitiesUserMention struct {
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
func (r EmbeddedTweetQuotedTweetRetweetedTweetEntitiesUserMention) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetRetweetedTweetEntitiesUserMention) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetRetweetedTweetLimitedAction struct {
	Action string                                                    `json:"action"`
	Prompt EmbeddedTweetQuotedTweetRetweetedTweetLimitedActionPrompt `json:"prompt"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Action      respjson.Field
		Prompt      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetQuotedTweetRetweetedTweetLimitedAction) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetQuotedTweetRetweetedTweetLimitedAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetRetweetedTweetLimitedActionPrompt struct {
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
func (r EmbeddedTweetQuotedTweetRetweetedTweetLimitedActionPrompt) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetRetweetedTweetLimitedActionPrompt) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Complete Note Tweet content and rich-text metadata.
type EmbeddedTweetQuotedTweetRetweetedTweetNoteTweet struct {
	Text string `json:"text" api:"required"`
	ID   string `json:"id"`
	// Lists hashtags, symbols, links, and mentions from tweet text.
	Entities EmbeddedTweetQuotedTweetRetweetedTweetNoteTweetEntities `json:"entities"`
	// Inline media positions in the Note Tweet text.
	InlineMedia  []EmbeddedTweetQuotedTweetRetweetedTweetNoteTweetInlineMedia `json:"inlineMedia"`
	IsExpandable bool                                                         `json:"isExpandable"`
	RichtextTags []EmbeddedTweetQuotedTweetRetweetedTweetNoteTweetRichtextTag `json:"richtextTags"`
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
func (r EmbeddedTweetQuotedTweetRetweetedTweetNoteTweet) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetQuotedTweetRetweetedTweetNoteTweet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Lists hashtags, symbols, links, and mentions from tweet text.
type EmbeddedTweetQuotedTweetRetweetedTweetNoteTweetEntities struct {
	Hashtags     []EmbeddedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesHashtag     `json:"hashtags"`
	Smarttags    []EmbeddedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesSmarttag    `json:"smarttags"`
	Symbols      []EmbeddedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesSymbol      `json:"symbols"`
	Timestamps   []EmbeddedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesTimestamp   `json:"timestamps"`
	URLs         []EmbeddedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesURL         `json:"urls"`
	UserMentions []EmbeddedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesUserMention `json:"user_mentions"`
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
func (r EmbeddedTweetQuotedTweetRetweetedTweetNoteTweetEntities) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetQuotedTweetRetweetedTweetNoteTweetEntities) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Provides hashtag text and source offsets within a tweet.
type EmbeddedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesHashtag struct {
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
func (r EmbeddedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesHashtag) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesHashtag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indexed smart-tag, cashtag, or video timestamp metadata.
type EmbeddedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesSmarttag struct {
	Indices []int64                                                            `json:"indices"`
	Seconds float64                                                            `json:"seconds"`
	Tag     EmbeddedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesSmarttagTag `json:"tag"`
	Text    string                                                             `json:"text"`
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
func (r EmbeddedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesSmarttag) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesSmarttag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesSmarttagTag struct {
	Info EmbeddedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesSmarttagTagInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesSmarttagTag) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesSmarttagTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesSmarttagTagInfo struct {
	Info EmbeddedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesSmarttagTagInfoInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesSmarttagTagInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesSmarttagTagInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesSmarttagTagInfoInfo struct {
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
func (r EmbeddedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesSmarttagTagInfoInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesSmarttagTagInfoInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indexed smart-tag, cashtag, or video timestamp metadata.
type EmbeddedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesSymbol struct {
	Indices []int64                                                          `json:"indices"`
	Seconds float64                                                          `json:"seconds"`
	Tag     EmbeddedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesSymbolTag `json:"tag"`
	Text    string                                                           `json:"text"`
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
func (r EmbeddedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesSymbol) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesSymbol) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesSymbolTag struct {
	Info EmbeddedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesSymbolTagInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesSymbolTag) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesSymbolTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesSymbolTagInfo struct {
	Info EmbeddedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesSymbolTagInfoInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesSymbolTagInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesSymbolTagInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesSymbolTagInfoInfo struct {
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
func (r EmbeddedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesSymbolTagInfoInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesSymbolTagInfoInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indexed smart-tag, cashtag, or video timestamp metadata.
type EmbeddedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesTimestamp struct {
	Indices []int64                                                             `json:"indices"`
	Seconds float64                                                             `json:"seconds"`
	Tag     EmbeddedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesTimestampTag `json:"tag"`
	Text    string                                                              `json:"text"`
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
func (r EmbeddedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesTimestamp) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesTimestamp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesTimestampTag struct {
	Info EmbeddedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesTimestampTagInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesTimestampTag) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesTimestampTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesTimestampTagInfo struct {
	Info EmbeddedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesTimestampTagInfoInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesTimestampTagInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesTimestampTagInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesTimestampTagInfoInfo struct {
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
func (r EmbeddedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesTimestampTagInfoInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesTimestampTagInfoInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Provides shortened, display, and expanded URLs from tweet text.
type EmbeddedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesURL struct {
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
func (r EmbeddedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesURL) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Provides profile identity and source offsets for a mention.
type EmbeddedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesUserMention struct {
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
func (r EmbeddedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesUserMention) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesUserMention) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetRetweetedTweetNoteTweetInlineMedia struct {
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
func (r EmbeddedTweetQuotedTweetRetweetedTweetNoteTweetInlineMedia) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetRetweetedTweetNoteTweetInlineMedia) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetRetweetedTweetNoteTweetRichtextTag struct {
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
func (r EmbeddedTweetQuotedTweetRetweetedTweetNoteTweetRichtextTag) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetRetweetedTweetNoteTweetRichtextTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Describes public place metadata on a geotagged tweet.
type EmbeddedTweetQuotedTweetRetweetedTweetPlace struct {
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
func (r EmbeddedTweetQuotedTweetRetweetedTweetPlace) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetQuotedTweetRetweetedTweetPlace) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Engagement counts retained from a prior tweet edit.
type EmbeddedTweetQuotedTweetRetweetedTweetPreviousCounts struct {
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
func (r EmbeddedTweetQuotedTweetRetweetedTweetPreviousCounts) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetQuotedTweetRetweetedTweetPreviousCounts) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Final nested tweet context at depth 4.
type EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweet struct {
	ID            string `json:"id" api:"required"`
	BookmarkCount int64  `json:"bookmarkCount" api:"required"`
	LikeCount     int64  `json:"likeCount" api:"required"`
	QuoteCount    int64  `json:"quoteCount" api:"required"`
	ReplyCount    int64  `json:"replyCount" api:"required"`
	RetweetCount  int64  `json:"retweetCount" api:"required"`
	Text          string `json:"text" api:"required"`
	ViewCount     int64  `json:"viewCount" api:"required"`
	// Describes an X Article preview and its lifecycle metadata.
	Article EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetArticle `json:"article"`
	// X user profile with bio, follower counts, and verification status.
	Author UserProfile `json:"author"`
	// Describes a public card and its referenced profiles.
	Card EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetCard `json:"card"`
	// Community ID.
	CommunityID string `json:"communityId"`
	// Community Note presentation metadata returned by X.
	CommunityNote EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetCommunityNote `json:"communityNote"`
	// Content disclosure metadata shown by X when a tweet is labeled as paid
	// partnership content or AI-generated media.
	ContentDisclosure ContentDisclosure `json:"contentDisclosure"`
	// Public reply policy and conversation owner.
	ConversationControl EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetConversationControl `json:"conversationControl"`
	ConversationID      string                                                               `json:"conversationId"`
	CreatedAt           string                                                               `json:"createdAt"`
	DisplayTextRange    []int64                                                              `json:"displayTextRange"`
	// Lists edit-chain identifiers and the remaining edit window.
	Edit EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetEdit `json:"edit"`
	// Lists hashtags, symbols, links, and mentions from tweet text.
	Entities EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetEntities `json:"entities"`
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
	LimitedActions []EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetLimitedAction `json:"limitedActions"`
	// Attached media items, omitted when unavailable.
	Media []TweetMedia `json:"media"`
	// Complete Note Tweet content and rich-text metadata.
	NoteTweet EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetNoteTweet `json:"noteTweet"`
	// Describes public place metadata on a geotagged tweet.
	Place             EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetPlace `json:"place"`
	PossiblySensitive bool                                                   `json:"possiblySensitive"`
	// Public metadata whose fields are defined by X.
	PostCta map[string]any `json:"postCta"`
	// Engagement counts retained from a prior tweet edit.
	PreviousCounts EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetPreviousCounts `json:"previousCounts"`
	// Quoted tweet ID.
	QuotedTweetID string `json:"quotedTweetId"`
	// Public post and user referenced by this reaction.
	ReactionContext EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetReactionContext `json:"reactionContext"`
	// Public metadata whose fields are defined by X.
	Scopes map[string]any `json:"scopes"`
	Source string         `json:"source"`
	// Public visibility notice attached to an available tweet.
	Tombstone EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetTombstone `json:"tombstone"`
	Type      string                                                     `json:"type"`
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
		QuotedTweetID       respjson.Field
		ReactionContext     respjson.Field
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
func (r EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweet) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Describes an X Article preview and its lifecycle metadata.
type EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetArticle struct {
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
func (r EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetArticle) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetArticle) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Describes a public card and its referenced profiles.
type EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetCard struct {
	ID string `json:"id"`
	// Public metadata whose fields are defined by X.
	BindingValues map[string]any `json:"bindingValues"`
	Name          string         `json:"name"`
	// Public metadata whose fields are defined by X.
	Platform map[string]any `json:"platform"`
	URL      string         `json:"url"`
	// Unresolved card user references.
	UserReferenceErrors []EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetCardUserReferenceError `json:"userReferenceErrors"`
	UserReferences      []UserProfile                                                             `json:"userReferences"`
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
func (r EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetCard) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetCard) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetCardUserReferenceError struct {
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
func (r EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetCardUserReferenceError) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetCardUserReferenceError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Community Note presentation metadata returned by X.
type EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetCommunityNote struct {
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
func (r EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetCommunityNote) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetCommunityNote) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Public reply policy and conversation owner.
type EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetConversationControl struct {
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
func (r EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetConversationControl) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetConversationControl) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Lists edit-chain identifiers and the remaining edit window.
type EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetEdit struct {
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
func (r EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetEdit) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetEdit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Lists hashtags, symbols, links, and mentions from tweet text.
type EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetEntities struct {
	Hashtags     []EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetEntitiesHashtag     `json:"hashtags"`
	Smarttags    []EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetEntitiesSmarttag    `json:"smarttags"`
	Symbols      []EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetEntitiesSymbol      `json:"symbols"`
	Timestamps   []EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetEntitiesTimestamp   `json:"timestamps"`
	URLs         []EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetEntitiesURL         `json:"urls"`
	UserMentions []EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetEntitiesUserMention `json:"user_mentions"`
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
func (r EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetEntities) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetEntities) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Provides hashtag text and source offsets within a tweet.
type EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetEntitiesHashtag struct {
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
func (r EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetEntitiesHashtag) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetEntitiesHashtag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indexed smart-tag, cashtag, or video timestamp metadata.
type EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetEntitiesSmarttag struct {
	Indices []int64                                                              `json:"indices"`
	Seconds float64                                                              `json:"seconds"`
	Tag     EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetEntitiesSmarttagTag `json:"tag"`
	Text    string                                                               `json:"text"`
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
func (r EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetEntitiesSmarttag) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetEntitiesSmarttag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetEntitiesSmarttagTag struct {
	Info EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetEntitiesSmarttagTagInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetEntitiesSmarttagTag) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetEntitiesSmarttagTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetEntitiesSmarttagTagInfo struct {
	Info EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetEntitiesSmarttagTagInfoInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetEntitiesSmarttagTagInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetEntitiesSmarttagTagInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetEntitiesSmarttagTagInfoInfo struct {
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
func (r EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetEntitiesSmarttagTagInfoInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetEntitiesSmarttagTagInfoInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indexed smart-tag, cashtag, or video timestamp metadata.
type EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetEntitiesSymbol struct {
	Indices []int64                                                            `json:"indices"`
	Seconds float64                                                            `json:"seconds"`
	Tag     EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetEntitiesSymbolTag `json:"tag"`
	Text    string                                                             `json:"text"`
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
func (r EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetEntitiesSymbol) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetEntitiesSymbol) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetEntitiesSymbolTag struct {
	Info EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetEntitiesSymbolTagInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetEntitiesSymbolTag) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetEntitiesSymbolTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetEntitiesSymbolTagInfo struct {
	Info EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetEntitiesSymbolTagInfoInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetEntitiesSymbolTagInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetEntitiesSymbolTagInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetEntitiesSymbolTagInfoInfo struct {
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
func (r EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetEntitiesSymbolTagInfoInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetEntitiesSymbolTagInfoInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indexed smart-tag, cashtag, or video timestamp metadata.
type EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetEntitiesTimestamp struct {
	Indices []int64                                                               `json:"indices"`
	Seconds float64                                                               `json:"seconds"`
	Tag     EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetEntitiesTimestampTag `json:"tag"`
	Text    string                                                                `json:"text"`
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
func (r EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetEntitiesTimestamp) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetEntitiesTimestamp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetEntitiesTimestampTag struct {
	Info EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetEntitiesTimestampTagInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetEntitiesTimestampTag) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetEntitiesTimestampTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetEntitiesTimestampTagInfo struct {
	Info EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetEntitiesTimestampTagInfoInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetEntitiesTimestampTagInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetEntitiesTimestampTagInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetEntitiesTimestampTagInfoInfo struct {
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
func (r EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetEntitiesTimestampTagInfoInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetEntitiesTimestampTagInfoInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Provides shortened, display, and expanded URLs from tweet text.
type EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetEntitiesURL struct {
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
func (r EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetEntitiesURL) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetEntitiesURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Provides profile identity and source offsets for a mention.
type EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetEntitiesUserMention struct {
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
func (r EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetEntitiesUserMention) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetEntitiesUserMention) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetLimitedAction struct {
	Action string                                                               `json:"action"`
	Prompt EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetLimitedActionPrompt `json:"prompt"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Action      respjson.Field
		Prompt      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetLimitedAction) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetLimitedAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetLimitedActionPrompt struct {
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
func (r EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetLimitedActionPrompt) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetLimitedActionPrompt) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Complete Note Tweet content and rich-text metadata.
type EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetNoteTweet struct {
	Text string `json:"text" api:"required"`
	ID   string `json:"id"`
	// Lists hashtags, symbols, links, and mentions from tweet text.
	Entities EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetNoteTweetEntities `json:"entities"`
	// Inline media positions in the Note Tweet text.
	InlineMedia  []EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetNoteTweetInlineMedia `json:"inlineMedia"`
	IsExpandable bool                                                                    `json:"isExpandable"`
	RichtextTags []EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetNoteTweetRichtextTag `json:"richtextTags"`
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
func (r EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetNoteTweet) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetNoteTweet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Lists hashtags, symbols, links, and mentions from tweet text.
type EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetNoteTweetEntities struct {
	Hashtags     []EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesHashtag     `json:"hashtags"`
	Smarttags    []EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesSmarttag    `json:"smarttags"`
	Symbols      []EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesSymbol      `json:"symbols"`
	Timestamps   []EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesTimestamp   `json:"timestamps"`
	URLs         []EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesURL         `json:"urls"`
	UserMentions []EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesUserMention `json:"user_mentions"`
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
func (r EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetNoteTweetEntities) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetNoteTweetEntities) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Provides hashtag text and source offsets within a tweet.
type EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesHashtag struct {
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
func (r EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesHashtag) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesHashtag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indexed smart-tag, cashtag, or video timestamp metadata.
type EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesSmarttag struct {
	Indices []int64                                                                       `json:"indices"`
	Seconds float64                                                                       `json:"seconds"`
	Tag     EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesSmarttagTag `json:"tag"`
	Text    string                                                                        `json:"text"`
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
func (r EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesSmarttag) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesSmarttag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesSmarttagTag struct {
	Info EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesSmarttagTagInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesSmarttagTag) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesSmarttagTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesSmarttagTagInfo struct {
	Info EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesSmarttagTagInfoInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesSmarttagTagInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesSmarttagTagInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesSmarttagTagInfoInfo struct {
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
func (r EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesSmarttagTagInfoInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesSmarttagTagInfoInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indexed smart-tag, cashtag, or video timestamp metadata.
type EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesSymbol struct {
	Indices []int64                                                                     `json:"indices"`
	Seconds float64                                                                     `json:"seconds"`
	Tag     EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesSymbolTag `json:"tag"`
	Text    string                                                                      `json:"text"`
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
func (r EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesSymbol) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesSymbol) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesSymbolTag struct {
	Info EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesSymbolTagInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesSymbolTag) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesSymbolTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesSymbolTagInfo struct {
	Info EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesSymbolTagInfoInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesSymbolTagInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesSymbolTagInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesSymbolTagInfoInfo struct {
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
func (r EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesSymbolTagInfoInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesSymbolTagInfoInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indexed smart-tag, cashtag, or video timestamp metadata.
type EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesTimestamp struct {
	Indices []int64                                                                        `json:"indices"`
	Seconds float64                                                                        `json:"seconds"`
	Tag     EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesTimestampTag `json:"tag"`
	Text    string                                                                         `json:"text"`
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
func (r EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesTimestamp) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesTimestamp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesTimestampTag struct {
	Info EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesTimestampTagInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesTimestampTag) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesTimestampTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesTimestampTagInfo struct {
	Info EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesTimestampTagInfoInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesTimestampTagInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesTimestampTagInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesTimestampTagInfoInfo struct {
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
func (r EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesTimestampTagInfoInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesTimestampTagInfoInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Provides shortened, display, and expanded URLs from tweet text.
type EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesURL struct {
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
func (r EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesURL) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Provides profile identity and source offsets for a mention.
type EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesUserMention struct {
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
func (r EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesUserMention) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesUserMention) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetNoteTweetInlineMedia struct {
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
func (r EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetNoteTweetInlineMedia) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetNoteTweetInlineMedia) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetNoteTweetRichtextTag struct {
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
func (r EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetNoteTweetRichtextTag) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetNoteTweetRichtextTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Describes public place metadata on a geotagged tweet.
type EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetPlace struct {
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
func (r EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetPlace) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetPlace) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Engagement counts retained from a prior tweet edit.
type EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetPreviousCounts struct {
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
func (r EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetPreviousCounts) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetPreviousCounts) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Public post and user referenced by this reaction.
type EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetReactionContext struct {
	// Referenced post ID.
	ReactedToPostID string `json:"reactedToPostId"`
	// X user profile with bio, follower counts, and verification status.
	ReactedToUser UserProfile `json:"reactedToUser"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ReactedToPostID respjson.Field
		ReactedToUser   respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetReactionContext) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetReactionContext) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Public visibility notice attached to an available tweet.
type EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetTombstone struct {
	Text EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetTombstoneText `json:"text"`
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
func (r EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetTombstone) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetTombstone) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetTombstoneText struct {
	Entities []EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetTombstoneTextEntity `json:"entities"`
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
func (r EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetTombstoneText) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetTombstoneText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetTombstoneTextEntity struct {
	FromIndex int64                                                                   `json:"fromIndex"`
	Ref       EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetTombstoneTextEntityRef `json:"ref"`
	ToIndex   int64                                                                   `json:"toIndex"`
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
func (r EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetTombstoneTextEntity) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetTombstoneTextEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetTombstoneTextEntityRef struct {
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
func (r EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetTombstoneTextEntityRef) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetTombstoneTextEntityRef) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Public post and user referenced by this reaction.
type EmbeddedTweetQuotedTweetRetweetedTweetReactionContext struct {
	// Referenced post ID.
	ReactedToPostID string `json:"reactedToPostId"`
	// X user profile with bio, follower counts, and verification status.
	ReactedToUser UserProfile `json:"reactedToUser"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ReactedToPostID respjson.Field
		ReactedToUser   respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetQuotedTweetRetweetedTweetReactionContext) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetQuotedTweetRetweetedTweetReactionContext) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Final nested tweet context at depth 4.
type EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweet struct {
	ID            string `json:"id" api:"required"`
	BookmarkCount int64  `json:"bookmarkCount" api:"required"`
	LikeCount     int64  `json:"likeCount" api:"required"`
	QuoteCount    int64  `json:"quoteCount" api:"required"`
	ReplyCount    int64  `json:"replyCount" api:"required"`
	RetweetCount  int64  `json:"retweetCount" api:"required"`
	Text          string `json:"text" api:"required"`
	ViewCount     int64  `json:"viewCount" api:"required"`
	// Describes an X Article preview and its lifecycle metadata.
	Article EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetArticle `json:"article"`
	// X user profile with bio, follower counts, and verification status.
	Author UserProfile `json:"author"`
	// Describes a public card and its referenced profiles.
	Card EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetCard `json:"card"`
	// Community ID.
	CommunityID string `json:"communityId"`
	// Community Note presentation metadata returned by X.
	CommunityNote EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetCommunityNote `json:"communityNote"`
	// Content disclosure metadata shown by X when a tweet is labeled as paid
	// partnership content or AI-generated media.
	ContentDisclosure ContentDisclosure `json:"contentDisclosure"`
	// Public reply policy and conversation owner.
	ConversationControl EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetConversationControl `json:"conversationControl"`
	ConversationID      string                                                                  `json:"conversationId"`
	CreatedAt           string                                                                  `json:"createdAt"`
	DisplayTextRange    []int64                                                                 `json:"displayTextRange"`
	// Lists edit-chain identifiers and the remaining edit window.
	Edit EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetEdit `json:"edit"`
	// Lists hashtags, symbols, links, and mentions from tweet text.
	Entities EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetEntities `json:"entities"`
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
	LimitedActions []EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetLimitedAction `json:"limitedActions"`
	// Attached media items, omitted when unavailable.
	Media []TweetMedia `json:"media"`
	// Complete Note Tweet content and rich-text metadata.
	NoteTweet EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetNoteTweet `json:"noteTweet"`
	// Describes public place metadata on a geotagged tweet.
	Place             EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetPlace `json:"place"`
	PossiblySensitive bool                                                      `json:"possiblySensitive"`
	// Public metadata whose fields are defined by X.
	PostCta map[string]any `json:"postCta"`
	// Engagement counts retained from a prior tweet edit.
	PreviousCounts EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetPreviousCounts `json:"previousCounts"`
	// Quoted tweet ID.
	QuotedTweetID string `json:"quotedTweetId"`
	// Public post and user referenced by this reaction.
	ReactionContext EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetReactionContext `json:"reactionContext"`
	// Public metadata whose fields are defined by X.
	Scopes map[string]any `json:"scopes"`
	Source string         `json:"source"`
	// Public visibility notice attached to an available tweet.
	Tombstone EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetTombstone `json:"tombstone"`
	Type      string                                                        `json:"type"`
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
		QuotedTweetID       respjson.Field
		ReactionContext     respjson.Field
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
func (r EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweet) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Describes an X Article preview and its lifecycle metadata.
type EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetArticle struct {
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
func (r EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetArticle) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetArticle) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Describes a public card and its referenced profiles.
type EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetCard struct {
	ID string `json:"id"`
	// Public metadata whose fields are defined by X.
	BindingValues map[string]any `json:"bindingValues"`
	Name          string         `json:"name"`
	// Public metadata whose fields are defined by X.
	Platform map[string]any `json:"platform"`
	URL      string         `json:"url"`
	// Unresolved card user references.
	UserReferenceErrors []EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetCardUserReferenceError `json:"userReferenceErrors"`
	UserReferences      []UserProfile                                                                `json:"userReferences"`
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
func (r EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetCard) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetCard) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetCardUserReferenceError struct {
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
func (r EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetCardUserReferenceError) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetCardUserReferenceError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Community Note presentation metadata returned by X.
type EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetCommunityNote struct {
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
func (r EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetCommunityNote) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetCommunityNote) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Public reply policy and conversation owner.
type EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetConversationControl struct {
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
func (r EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetConversationControl) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetConversationControl) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Lists edit-chain identifiers and the remaining edit window.
type EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetEdit struct {
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
func (r EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetEdit) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetEdit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Lists hashtags, symbols, links, and mentions from tweet text.
type EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetEntities struct {
	Hashtags     []EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetEntitiesHashtag     `json:"hashtags"`
	Smarttags    []EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetEntitiesSmarttag    `json:"smarttags"`
	Symbols      []EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetEntitiesSymbol      `json:"symbols"`
	Timestamps   []EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetEntitiesTimestamp   `json:"timestamps"`
	URLs         []EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetEntitiesURL         `json:"urls"`
	UserMentions []EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetEntitiesUserMention `json:"user_mentions"`
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
func (r EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetEntities) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetEntities) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Provides hashtag text and source offsets within a tweet.
type EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetEntitiesHashtag struct {
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
func (r EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetEntitiesHashtag) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetEntitiesHashtag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indexed smart-tag, cashtag, or video timestamp metadata.
type EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetEntitiesSmarttag struct {
	Indices []int64                                                                 `json:"indices"`
	Seconds float64                                                                 `json:"seconds"`
	Tag     EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetEntitiesSmarttagTag `json:"tag"`
	Text    string                                                                  `json:"text"`
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
func (r EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetEntitiesSmarttag) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetEntitiesSmarttag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetEntitiesSmarttagTag struct {
	Info EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetEntitiesSmarttagTagInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetEntitiesSmarttagTag) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetEntitiesSmarttagTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetEntitiesSmarttagTagInfo struct {
	Info EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetEntitiesSmarttagTagInfoInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetEntitiesSmarttagTagInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetEntitiesSmarttagTagInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetEntitiesSmarttagTagInfoInfo struct {
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
func (r EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetEntitiesSmarttagTagInfoInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetEntitiesSmarttagTagInfoInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indexed smart-tag, cashtag, or video timestamp metadata.
type EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetEntitiesSymbol struct {
	Indices []int64                                                               `json:"indices"`
	Seconds float64                                                               `json:"seconds"`
	Tag     EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetEntitiesSymbolTag `json:"tag"`
	Text    string                                                                `json:"text"`
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
func (r EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetEntitiesSymbol) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetEntitiesSymbol) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetEntitiesSymbolTag struct {
	Info EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetEntitiesSymbolTagInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetEntitiesSymbolTag) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetEntitiesSymbolTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetEntitiesSymbolTagInfo struct {
	Info EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetEntitiesSymbolTagInfoInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetEntitiesSymbolTagInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetEntitiesSymbolTagInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetEntitiesSymbolTagInfoInfo struct {
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
func (r EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetEntitiesSymbolTagInfoInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetEntitiesSymbolTagInfoInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indexed smart-tag, cashtag, or video timestamp metadata.
type EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetEntitiesTimestamp struct {
	Indices []int64                                                                  `json:"indices"`
	Seconds float64                                                                  `json:"seconds"`
	Tag     EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetEntitiesTimestampTag `json:"tag"`
	Text    string                                                                   `json:"text"`
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
func (r EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetEntitiesTimestamp) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetEntitiesTimestamp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetEntitiesTimestampTag struct {
	Info EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetEntitiesTimestampTagInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetEntitiesTimestampTag) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetEntitiesTimestampTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetEntitiesTimestampTagInfo struct {
	Info EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetEntitiesTimestampTagInfoInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetEntitiesTimestampTagInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetEntitiesTimestampTagInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetEntitiesTimestampTagInfoInfo struct {
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
func (r EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetEntitiesTimestampTagInfoInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetEntitiesTimestampTagInfoInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Provides shortened, display, and expanded URLs from tweet text.
type EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetEntitiesURL struct {
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
func (r EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetEntitiesURL) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetEntitiesURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Provides profile identity and source offsets for a mention.
type EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetEntitiesUserMention struct {
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
func (r EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetEntitiesUserMention) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetEntitiesUserMention) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetLimitedAction struct {
	Action string                                                                  `json:"action"`
	Prompt EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetLimitedActionPrompt `json:"prompt"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Action      respjson.Field
		Prompt      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetLimitedAction) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetLimitedAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetLimitedActionPrompt struct {
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
func (r EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetLimitedActionPrompt) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetLimitedActionPrompt) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Complete Note Tweet content and rich-text metadata.
type EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetNoteTweet struct {
	Text string `json:"text" api:"required"`
	ID   string `json:"id"`
	// Lists hashtags, symbols, links, and mentions from tweet text.
	Entities EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetNoteTweetEntities `json:"entities"`
	// Inline media positions in the Note Tweet text.
	InlineMedia  []EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetNoteTweetInlineMedia `json:"inlineMedia"`
	IsExpandable bool                                                                       `json:"isExpandable"`
	RichtextTags []EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetNoteTweetRichtextTag `json:"richtextTags"`
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
func (r EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetNoteTweet) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetNoteTweet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Lists hashtags, symbols, links, and mentions from tweet text.
type EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetNoteTweetEntities struct {
	Hashtags     []EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesHashtag     `json:"hashtags"`
	Smarttags    []EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesSmarttag    `json:"smarttags"`
	Symbols      []EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesSymbol      `json:"symbols"`
	Timestamps   []EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesTimestamp   `json:"timestamps"`
	URLs         []EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesURL         `json:"urls"`
	UserMentions []EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesUserMention `json:"user_mentions"`
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
func (r EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetNoteTweetEntities) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetNoteTweetEntities) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Provides hashtag text and source offsets within a tweet.
type EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesHashtag struct {
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
func (r EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesHashtag) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesHashtag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indexed smart-tag, cashtag, or video timestamp metadata.
type EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesSmarttag struct {
	Indices []int64                                                                          `json:"indices"`
	Seconds float64                                                                          `json:"seconds"`
	Tag     EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesSmarttagTag `json:"tag"`
	Text    string                                                                           `json:"text"`
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
func (r EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesSmarttag) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesSmarttag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesSmarttagTag struct {
	Info EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesSmarttagTagInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesSmarttagTag) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesSmarttagTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesSmarttagTagInfo struct {
	Info EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesSmarttagTagInfoInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesSmarttagTagInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesSmarttagTagInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesSmarttagTagInfoInfo struct {
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
func (r EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesSmarttagTagInfoInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesSmarttagTagInfoInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indexed smart-tag, cashtag, or video timestamp metadata.
type EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesSymbol struct {
	Indices []int64                                                                        `json:"indices"`
	Seconds float64                                                                        `json:"seconds"`
	Tag     EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesSymbolTag `json:"tag"`
	Text    string                                                                         `json:"text"`
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
func (r EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesSymbol) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesSymbol) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesSymbolTag struct {
	Info EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesSymbolTagInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesSymbolTag) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesSymbolTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesSymbolTagInfo struct {
	Info EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesSymbolTagInfoInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesSymbolTagInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesSymbolTagInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesSymbolTagInfoInfo struct {
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
func (r EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesSymbolTagInfoInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesSymbolTagInfoInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indexed smart-tag, cashtag, or video timestamp metadata.
type EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesTimestamp struct {
	Indices []int64                                                                           `json:"indices"`
	Seconds float64                                                                           `json:"seconds"`
	Tag     EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesTimestampTag `json:"tag"`
	Text    string                                                                            `json:"text"`
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
func (r EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesTimestamp) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesTimestamp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesTimestampTag struct {
	Info EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesTimestampTagInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesTimestampTag) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesTimestampTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesTimestampTagInfo struct {
	Info EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesTimestampTagInfoInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesTimestampTagInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesTimestampTagInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesTimestampTagInfoInfo struct {
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
func (r EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesTimestampTagInfoInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesTimestampTagInfoInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Provides shortened, display, and expanded URLs from tweet text.
type EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesURL struct {
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
func (r EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesURL) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Provides profile identity and source offsets for a mention.
type EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesUserMention struct {
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
func (r EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesUserMention) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesUserMention) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetNoteTweetInlineMedia struct {
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
func (r EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetNoteTweetInlineMedia) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetNoteTweetInlineMedia) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetNoteTweetRichtextTag struct {
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
func (r EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetNoteTweetRichtextTag) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetNoteTweetRichtextTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Describes public place metadata on a geotagged tweet.
type EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetPlace struct {
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
func (r EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetPlace) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetPlace) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Engagement counts retained from a prior tweet edit.
type EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetPreviousCounts struct {
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
func (r EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetPreviousCounts) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetPreviousCounts) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Public post and user referenced by this reaction.
type EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetReactionContext struct {
	// Referenced post ID.
	ReactedToPostID string `json:"reactedToPostId"`
	// X user profile with bio, follower counts, and verification status.
	ReactedToUser UserProfile `json:"reactedToUser"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ReactedToPostID respjson.Field
		ReactedToUser   respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetReactionContext) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetReactionContext) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Public visibility notice attached to an available tweet.
type EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetTombstone struct {
	Text EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetTombstoneText `json:"text"`
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
func (r EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetTombstone) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetTombstone) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetTombstoneText struct {
	Entities []EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetTombstoneTextEntity `json:"entities"`
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
func (r EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetTombstoneText) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetTombstoneText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetTombstoneTextEntity struct {
	FromIndex int64                                                                      `json:"fromIndex"`
	Ref       EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetTombstoneTextEntityRef `json:"ref"`
	ToIndex   int64                                                                      `json:"toIndex"`
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
func (r EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetTombstoneTextEntity) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetTombstoneTextEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetTombstoneTextEntityRef struct {
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
func (r EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetTombstoneTextEntityRef) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetTombstoneTextEntityRef) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Public visibility notice attached to an available tweet.
type EmbeddedTweetQuotedTweetRetweetedTweetTombstone struct {
	Text EmbeddedTweetQuotedTweetRetweetedTweetTombstoneText `json:"text"`
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
func (r EmbeddedTweetQuotedTweetRetweetedTweetTombstone) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetQuotedTweetRetweetedTweetTombstone) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetRetweetedTweetTombstoneText struct {
	Entities []EmbeddedTweetQuotedTweetRetweetedTweetTombstoneTextEntity `json:"entities"`
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
func (r EmbeddedTweetQuotedTweetRetweetedTweetTombstoneText) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetQuotedTweetRetweetedTweetTombstoneText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetRetweetedTweetTombstoneTextEntity struct {
	FromIndex int64                                                        `json:"fromIndex"`
	Ref       EmbeddedTweetQuotedTweetRetweetedTweetTombstoneTextEntityRef `json:"ref"`
	ToIndex   int64                                                        `json:"toIndex"`
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
func (r EmbeddedTweetQuotedTweetRetweetedTweetTombstoneTextEntity) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetRetweetedTweetTombstoneTextEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetRetweetedTweetTombstoneTextEntityRef struct {
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
func (r EmbeddedTweetQuotedTweetRetweetedTweetTombstoneTextEntityRef) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetQuotedTweetRetweetedTweetTombstoneTextEntityRef) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Public visibility notice attached to an available tweet.
type EmbeddedTweetQuotedTweetTombstone struct {
	Text EmbeddedTweetQuotedTweetTombstoneText `json:"text"`
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
func (r EmbeddedTweetQuotedTweetTombstone) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetQuotedTweetTombstone) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetTombstoneText struct {
	Entities []EmbeddedTweetQuotedTweetTombstoneTextEntity `json:"entities"`
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
func (r EmbeddedTweetQuotedTweetTombstoneText) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetQuotedTweetTombstoneText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetTombstoneTextEntity struct {
	FromIndex int64                                          `json:"fromIndex"`
	Ref       EmbeddedTweetQuotedTweetTombstoneTextEntityRef `json:"ref"`
	ToIndex   int64                                          `json:"toIndex"`
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
func (r EmbeddedTweetQuotedTweetTombstoneTextEntity) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetQuotedTweetTombstoneTextEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetQuotedTweetTombstoneTextEntityRef struct {
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
func (r EmbeddedTweetQuotedTweetTombstoneTextEntityRef) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetQuotedTweetTombstoneTextEntityRef) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Public post and user referenced by this reaction.
type EmbeddedTweetReactionContext struct {
	// Referenced post ID.
	ReactedToPostID string `json:"reactedToPostId"`
	// X user profile with bio, follower counts, and verification status.
	ReactedToUser UserProfile `json:"reactedToUser"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ReactedToPostID respjson.Field
		ReactedToUser   respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetReactionContext) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetReactionContext) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Nested tweet context at depth 2.
type EmbeddedTweetRetweetedTweet struct {
	ID            string `json:"id" api:"required"`
	BookmarkCount int64  `json:"bookmarkCount" api:"required"`
	LikeCount     int64  `json:"likeCount" api:"required"`
	QuoteCount    int64  `json:"quoteCount" api:"required"`
	ReplyCount    int64  `json:"replyCount" api:"required"`
	RetweetCount  int64  `json:"retweetCount" api:"required"`
	Text          string `json:"text" api:"required"`
	ViewCount     int64  `json:"viewCount" api:"required"`
	// Describes an X Article preview and its lifecycle metadata.
	Article EmbeddedTweetRetweetedTweetArticle `json:"article"`
	// X user profile with bio, follower counts, and verification status.
	Author UserProfile `json:"author"`
	// Describes a public card and its referenced profiles.
	Card EmbeddedTweetRetweetedTweetCard `json:"card"`
	// Community ID.
	CommunityID string `json:"communityId"`
	// Community Note presentation metadata returned by X.
	CommunityNote EmbeddedTweetRetweetedTweetCommunityNote `json:"communityNote"`
	// Content disclosure metadata shown by X when a tweet is labeled as paid
	// partnership content or AI-generated media.
	ContentDisclosure ContentDisclosure `json:"contentDisclosure"`
	// Public reply policy and conversation owner.
	ConversationControl EmbeddedTweetRetweetedTweetConversationControl `json:"conversationControl"`
	ConversationID      string                                         `json:"conversationId"`
	CreatedAt           string                                         `json:"createdAt"`
	DisplayTextRange    []int64                                        `json:"displayTextRange"`
	// Lists edit-chain identifiers and the remaining edit window.
	Edit EmbeddedTweetRetweetedTweetEdit `json:"edit"`
	// Lists hashtags, symbols, links, and mentions from tweet text.
	Entities EmbeddedTweetRetweetedTweetEntities `json:"entities"`
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
	LimitedActions []EmbeddedTweetRetweetedTweetLimitedAction `json:"limitedActions"`
	// Attached media items, omitted when unavailable.
	Media []TweetMedia `json:"media"`
	// Complete Note Tweet content and rich-text metadata.
	NoteTweet EmbeddedTweetRetweetedTweetNoteTweet `json:"noteTweet"`
	// Describes public place metadata on a geotagged tweet.
	Place             EmbeddedTweetRetweetedTweetPlace `json:"place"`
	PossiblySensitive bool                             `json:"possiblySensitive"`
	// Public metadata whose fields are defined by X.
	PostCta map[string]any `json:"postCta"`
	// Engagement counts retained from a prior tweet edit.
	PreviousCounts EmbeddedTweetRetweetedTweetPreviousCounts `json:"previousCounts"`
	// Nested tweet context at depth 3.
	QuotedTweet EmbeddedTweetRetweetedTweetQuotedTweet `json:"quoted_tweet"`
	// Quoted tweet ID.
	QuotedTweetID string `json:"quotedTweetId"`
	// Public post and user referenced by this reaction.
	ReactionContext EmbeddedTweetRetweetedTweetReactionContext `json:"reactionContext"`
	// Nested tweet context at depth 3.
	RetweetedTweet EmbeddedTweetRetweetedTweetRetweetedTweet `json:"retweeted_tweet"`
	// Public metadata whose fields are defined by X.
	Scopes map[string]any `json:"scopes"`
	Source string         `json:"source"`
	// Public visibility notice attached to an available tweet.
	Tombstone EmbeddedTweetRetweetedTweetTombstone `json:"tombstone"`
	Type      string                               `json:"type"`
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
func (r EmbeddedTweetRetweetedTweet) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetRetweetedTweet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Describes an X Article preview and its lifecycle metadata.
type EmbeddedTweetRetweetedTweetArticle struct {
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
func (r EmbeddedTweetRetweetedTweetArticle) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetRetweetedTweetArticle) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Describes a public card and its referenced profiles.
type EmbeddedTweetRetweetedTweetCard struct {
	ID string `json:"id"`
	// Public metadata whose fields are defined by X.
	BindingValues map[string]any `json:"bindingValues"`
	Name          string         `json:"name"`
	// Public metadata whose fields are defined by X.
	Platform map[string]any `json:"platform"`
	URL      string         `json:"url"`
	// Unresolved card user references.
	UserReferenceErrors []EmbeddedTweetRetweetedTweetCardUserReferenceError `json:"userReferenceErrors"`
	UserReferences      []UserProfile                                       `json:"userReferences"`
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
func (r EmbeddedTweetRetweetedTweetCard) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetRetweetedTweetCard) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetCardUserReferenceError struct {
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
func (r EmbeddedTweetRetweetedTweetCardUserReferenceError) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetRetweetedTweetCardUserReferenceError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Community Note presentation metadata returned by X.
type EmbeddedTweetRetweetedTweetCommunityNote struct {
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
func (r EmbeddedTweetRetweetedTweetCommunityNote) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetRetweetedTweetCommunityNote) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Public reply policy and conversation owner.
type EmbeddedTweetRetweetedTweetConversationControl struct {
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
func (r EmbeddedTweetRetweetedTweetConversationControl) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetRetweetedTweetConversationControl) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Lists edit-chain identifiers and the remaining edit window.
type EmbeddedTweetRetweetedTweetEdit struct {
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
func (r EmbeddedTweetRetweetedTweetEdit) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetRetweetedTweetEdit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Lists hashtags, symbols, links, and mentions from tweet text.
type EmbeddedTweetRetweetedTweetEntities struct {
	Hashtags     []EmbeddedTweetRetweetedTweetEntitiesHashtag     `json:"hashtags"`
	Smarttags    []EmbeddedTweetRetweetedTweetEntitiesSmarttag    `json:"smarttags"`
	Symbols      []EmbeddedTweetRetweetedTweetEntitiesSymbol      `json:"symbols"`
	Timestamps   []EmbeddedTweetRetweetedTweetEntitiesTimestamp   `json:"timestamps"`
	URLs         []EmbeddedTweetRetweetedTweetEntitiesURL         `json:"urls"`
	UserMentions []EmbeddedTweetRetweetedTweetEntitiesUserMention `json:"user_mentions"`
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
func (r EmbeddedTweetRetweetedTweetEntities) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetRetweetedTweetEntities) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Provides hashtag text and source offsets within a tweet.
type EmbeddedTweetRetweetedTweetEntitiesHashtag struct {
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
func (r EmbeddedTweetRetweetedTweetEntitiesHashtag) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetRetweetedTweetEntitiesHashtag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indexed smart-tag, cashtag, or video timestamp metadata.
type EmbeddedTweetRetweetedTweetEntitiesSmarttag struct {
	Indices []int64                                        `json:"indices"`
	Seconds float64                                        `json:"seconds"`
	Tag     EmbeddedTweetRetweetedTweetEntitiesSmarttagTag `json:"tag"`
	Text    string                                         `json:"text"`
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
func (r EmbeddedTweetRetweetedTweetEntitiesSmarttag) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetRetweetedTweetEntitiesSmarttag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetEntitiesSmarttagTag struct {
	Info EmbeddedTweetRetweetedTweetEntitiesSmarttagTagInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetRetweetedTweetEntitiesSmarttagTag) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetRetweetedTweetEntitiesSmarttagTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetEntitiesSmarttagTagInfo struct {
	Info EmbeddedTweetRetweetedTweetEntitiesSmarttagTagInfoInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetRetweetedTweetEntitiesSmarttagTagInfo) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetRetweetedTweetEntitiesSmarttagTagInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetEntitiesSmarttagTagInfoInfo struct {
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
func (r EmbeddedTweetRetweetedTweetEntitiesSmarttagTagInfoInfo) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetRetweetedTweetEntitiesSmarttagTagInfoInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indexed smart-tag, cashtag, or video timestamp metadata.
type EmbeddedTweetRetweetedTweetEntitiesSymbol struct {
	Indices []int64                                      `json:"indices"`
	Seconds float64                                      `json:"seconds"`
	Tag     EmbeddedTweetRetweetedTweetEntitiesSymbolTag `json:"tag"`
	Text    string                                       `json:"text"`
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
func (r EmbeddedTweetRetweetedTweetEntitiesSymbol) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetRetweetedTweetEntitiesSymbol) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetEntitiesSymbolTag struct {
	Info EmbeddedTweetRetweetedTweetEntitiesSymbolTagInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetRetweetedTweetEntitiesSymbolTag) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetRetweetedTweetEntitiesSymbolTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetEntitiesSymbolTagInfo struct {
	Info EmbeddedTweetRetweetedTweetEntitiesSymbolTagInfoInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetRetweetedTweetEntitiesSymbolTagInfo) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetRetweetedTweetEntitiesSymbolTagInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetEntitiesSymbolTagInfoInfo struct {
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
func (r EmbeddedTweetRetweetedTweetEntitiesSymbolTagInfoInfo) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetRetweetedTweetEntitiesSymbolTagInfoInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indexed smart-tag, cashtag, or video timestamp metadata.
type EmbeddedTweetRetweetedTweetEntitiesTimestamp struct {
	Indices []int64                                         `json:"indices"`
	Seconds float64                                         `json:"seconds"`
	Tag     EmbeddedTweetRetweetedTweetEntitiesTimestampTag `json:"tag"`
	Text    string                                          `json:"text"`
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
func (r EmbeddedTweetRetweetedTweetEntitiesTimestamp) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetRetweetedTweetEntitiesTimestamp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetEntitiesTimestampTag struct {
	Info EmbeddedTweetRetweetedTweetEntitiesTimestampTagInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetRetweetedTweetEntitiesTimestampTag) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetRetweetedTweetEntitiesTimestampTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetEntitiesTimestampTagInfo struct {
	Info EmbeddedTweetRetweetedTweetEntitiesTimestampTagInfoInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetRetweetedTweetEntitiesTimestampTagInfo) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetRetweetedTweetEntitiesTimestampTagInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetEntitiesTimestampTagInfoInfo struct {
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
func (r EmbeddedTweetRetweetedTweetEntitiesTimestampTagInfoInfo) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetRetweetedTweetEntitiesTimestampTagInfoInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Provides shortened, display, and expanded URLs from tweet text.
type EmbeddedTweetRetweetedTweetEntitiesURL struct {
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
func (r EmbeddedTweetRetweetedTweetEntitiesURL) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetRetweetedTweetEntitiesURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Provides profile identity and source offsets for a mention.
type EmbeddedTweetRetweetedTweetEntitiesUserMention struct {
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
func (r EmbeddedTweetRetweetedTweetEntitiesUserMention) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetRetweetedTweetEntitiesUserMention) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetLimitedAction struct {
	Action string                                         `json:"action"`
	Prompt EmbeddedTweetRetweetedTweetLimitedActionPrompt `json:"prompt"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Action      respjson.Field
		Prompt      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetRetweetedTweetLimitedAction) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetRetweetedTweetLimitedAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetLimitedActionPrompt struct {
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
func (r EmbeddedTweetRetweetedTweetLimitedActionPrompt) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetRetweetedTweetLimitedActionPrompt) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Complete Note Tweet content and rich-text metadata.
type EmbeddedTweetRetweetedTweetNoteTweet struct {
	Text string `json:"text" api:"required"`
	ID   string `json:"id"`
	// Lists hashtags, symbols, links, and mentions from tweet text.
	Entities EmbeddedTweetRetweetedTweetNoteTweetEntities `json:"entities"`
	// Inline media positions in the Note Tweet text.
	InlineMedia  []EmbeddedTweetRetweetedTweetNoteTweetInlineMedia `json:"inlineMedia"`
	IsExpandable bool                                              `json:"isExpandable"`
	RichtextTags []EmbeddedTweetRetweetedTweetNoteTweetRichtextTag `json:"richtextTags"`
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
func (r EmbeddedTweetRetweetedTweetNoteTweet) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetRetweetedTweetNoteTweet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Lists hashtags, symbols, links, and mentions from tweet text.
type EmbeddedTweetRetweetedTweetNoteTweetEntities struct {
	Hashtags     []EmbeddedTweetRetweetedTweetNoteTweetEntitiesHashtag     `json:"hashtags"`
	Smarttags    []EmbeddedTweetRetweetedTweetNoteTweetEntitiesSmarttag    `json:"smarttags"`
	Symbols      []EmbeddedTweetRetweetedTweetNoteTweetEntitiesSymbol      `json:"symbols"`
	Timestamps   []EmbeddedTweetRetweetedTweetNoteTweetEntitiesTimestamp   `json:"timestamps"`
	URLs         []EmbeddedTweetRetweetedTweetNoteTweetEntitiesURL         `json:"urls"`
	UserMentions []EmbeddedTweetRetweetedTweetNoteTweetEntitiesUserMention `json:"user_mentions"`
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
func (r EmbeddedTweetRetweetedTweetNoteTweetEntities) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetRetweetedTweetNoteTweetEntities) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Provides hashtag text and source offsets within a tweet.
type EmbeddedTweetRetweetedTweetNoteTweetEntitiesHashtag struct {
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
func (r EmbeddedTweetRetweetedTweetNoteTweetEntitiesHashtag) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetRetweetedTweetNoteTweetEntitiesHashtag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indexed smart-tag, cashtag, or video timestamp metadata.
type EmbeddedTweetRetweetedTweetNoteTweetEntitiesSmarttag struct {
	Indices []int64                                                 `json:"indices"`
	Seconds float64                                                 `json:"seconds"`
	Tag     EmbeddedTweetRetweetedTweetNoteTweetEntitiesSmarttagTag `json:"tag"`
	Text    string                                                  `json:"text"`
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
func (r EmbeddedTweetRetweetedTweetNoteTweetEntitiesSmarttag) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetRetweetedTweetNoteTweetEntitiesSmarttag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetNoteTweetEntitiesSmarttagTag struct {
	Info EmbeddedTweetRetweetedTweetNoteTweetEntitiesSmarttagTagInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetRetweetedTweetNoteTweetEntitiesSmarttagTag) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetRetweetedTweetNoteTweetEntitiesSmarttagTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetNoteTweetEntitiesSmarttagTagInfo struct {
	Info EmbeddedTweetRetweetedTweetNoteTweetEntitiesSmarttagTagInfoInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetRetweetedTweetNoteTweetEntitiesSmarttagTagInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetNoteTweetEntitiesSmarttagTagInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetNoteTweetEntitiesSmarttagTagInfoInfo struct {
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
func (r EmbeddedTweetRetweetedTweetNoteTweetEntitiesSmarttagTagInfoInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetNoteTweetEntitiesSmarttagTagInfoInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indexed smart-tag, cashtag, or video timestamp metadata.
type EmbeddedTweetRetweetedTweetNoteTweetEntitiesSymbol struct {
	Indices []int64                                               `json:"indices"`
	Seconds float64                                               `json:"seconds"`
	Tag     EmbeddedTweetRetweetedTweetNoteTweetEntitiesSymbolTag `json:"tag"`
	Text    string                                                `json:"text"`
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
func (r EmbeddedTweetRetweetedTweetNoteTweetEntitiesSymbol) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetRetweetedTweetNoteTweetEntitiesSymbol) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetNoteTweetEntitiesSymbolTag struct {
	Info EmbeddedTweetRetweetedTweetNoteTweetEntitiesSymbolTagInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetRetweetedTweetNoteTweetEntitiesSymbolTag) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetRetweetedTweetNoteTweetEntitiesSymbolTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetNoteTweetEntitiesSymbolTagInfo struct {
	Info EmbeddedTweetRetweetedTweetNoteTweetEntitiesSymbolTagInfoInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetRetweetedTweetNoteTweetEntitiesSymbolTagInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetNoteTweetEntitiesSymbolTagInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetNoteTweetEntitiesSymbolTagInfoInfo struct {
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
func (r EmbeddedTweetRetweetedTweetNoteTweetEntitiesSymbolTagInfoInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetNoteTweetEntitiesSymbolTagInfoInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indexed smart-tag, cashtag, or video timestamp metadata.
type EmbeddedTweetRetweetedTweetNoteTweetEntitiesTimestamp struct {
	Indices []int64                                                  `json:"indices"`
	Seconds float64                                                  `json:"seconds"`
	Tag     EmbeddedTweetRetweetedTweetNoteTweetEntitiesTimestampTag `json:"tag"`
	Text    string                                                   `json:"text"`
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
func (r EmbeddedTweetRetweetedTweetNoteTweetEntitiesTimestamp) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetRetweetedTweetNoteTweetEntitiesTimestamp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetNoteTweetEntitiesTimestampTag struct {
	Info EmbeddedTweetRetweetedTweetNoteTweetEntitiesTimestampTagInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetRetweetedTweetNoteTweetEntitiesTimestampTag) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetRetweetedTweetNoteTweetEntitiesTimestampTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetNoteTweetEntitiesTimestampTagInfo struct {
	Info EmbeddedTweetRetweetedTweetNoteTweetEntitiesTimestampTagInfoInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetRetweetedTweetNoteTweetEntitiesTimestampTagInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetNoteTweetEntitiesTimestampTagInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetNoteTweetEntitiesTimestampTagInfoInfo struct {
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
func (r EmbeddedTweetRetweetedTweetNoteTweetEntitiesTimestampTagInfoInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetNoteTweetEntitiesTimestampTagInfoInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Provides shortened, display, and expanded URLs from tweet text.
type EmbeddedTweetRetweetedTweetNoteTweetEntitiesURL struct {
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
func (r EmbeddedTweetRetweetedTweetNoteTweetEntitiesURL) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetRetweetedTweetNoteTweetEntitiesURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Provides profile identity and source offsets for a mention.
type EmbeddedTweetRetweetedTweetNoteTweetEntitiesUserMention struct {
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
func (r EmbeddedTweetRetweetedTweetNoteTweetEntitiesUserMention) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetRetweetedTweetNoteTweetEntitiesUserMention) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetNoteTweetInlineMedia struct {
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
func (r EmbeddedTweetRetweetedTweetNoteTweetInlineMedia) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetRetweetedTweetNoteTweetInlineMedia) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetNoteTweetRichtextTag struct {
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
func (r EmbeddedTweetRetweetedTweetNoteTweetRichtextTag) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetRetweetedTweetNoteTweetRichtextTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Describes public place metadata on a geotagged tweet.
type EmbeddedTweetRetweetedTweetPlace struct {
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
func (r EmbeddedTweetRetweetedTweetPlace) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetRetweetedTweetPlace) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Engagement counts retained from a prior tweet edit.
type EmbeddedTweetRetweetedTweetPreviousCounts struct {
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
func (r EmbeddedTweetRetweetedTweetPreviousCounts) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetRetweetedTweetPreviousCounts) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Nested tweet context at depth 3.
type EmbeddedTweetRetweetedTweetQuotedTweet struct {
	ID            string `json:"id" api:"required"`
	BookmarkCount int64  `json:"bookmarkCount" api:"required"`
	LikeCount     int64  `json:"likeCount" api:"required"`
	QuoteCount    int64  `json:"quoteCount" api:"required"`
	ReplyCount    int64  `json:"replyCount" api:"required"`
	RetweetCount  int64  `json:"retweetCount" api:"required"`
	Text          string `json:"text" api:"required"`
	ViewCount     int64  `json:"viewCount" api:"required"`
	// Describes an X Article preview and its lifecycle metadata.
	Article EmbeddedTweetRetweetedTweetQuotedTweetArticle `json:"article"`
	// X user profile with bio, follower counts, and verification status.
	Author UserProfile `json:"author"`
	// Describes a public card and its referenced profiles.
	Card EmbeddedTweetRetweetedTweetQuotedTweetCard `json:"card"`
	// Community ID.
	CommunityID string `json:"communityId"`
	// Community Note presentation metadata returned by X.
	CommunityNote EmbeddedTweetRetweetedTweetQuotedTweetCommunityNote `json:"communityNote"`
	// Content disclosure metadata shown by X when a tweet is labeled as paid
	// partnership content or AI-generated media.
	ContentDisclosure ContentDisclosure `json:"contentDisclosure"`
	// Public reply policy and conversation owner.
	ConversationControl EmbeddedTweetRetweetedTweetQuotedTweetConversationControl `json:"conversationControl"`
	ConversationID      string                                                    `json:"conversationId"`
	CreatedAt           string                                                    `json:"createdAt"`
	DisplayTextRange    []int64                                                   `json:"displayTextRange"`
	// Lists edit-chain identifiers and the remaining edit window.
	Edit EmbeddedTweetRetweetedTweetQuotedTweetEdit `json:"edit"`
	// Lists hashtags, symbols, links, and mentions from tweet text.
	Entities EmbeddedTweetRetweetedTweetQuotedTweetEntities `json:"entities"`
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
	LimitedActions []EmbeddedTweetRetweetedTweetQuotedTweetLimitedAction `json:"limitedActions"`
	// Attached media items, omitted when unavailable.
	Media []TweetMedia `json:"media"`
	// Complete Note Tweet content and rich-text metadata.
	NoteTweet EmbeddedTweetRetweetedTweetQuotedTweetNoteTweet `json:"noteTweet"`
	// Describes public place metadata on a geotagged tweet.
	Place             EmbeddedTweetRetweetedTweetQuotedTweetPlace `json:"place"`
	PossiblySensitive bool                                        `json:"possiblySensitive"`
	// Public metadata whose fields are defined by X.
	PostCta map[string]any `json:"postCta"`
	// Engagement counts retained from a prior tweet edit.
	PreviousCounts EmbeddedTweetRetweetedTweetQuotedTweetPreviousCounts `json:"previousCounts"`
	// Final nested tweet context at depth 4.
	QuotedTweet EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweet `json:"quoted_tweet"`
	// Quoted tweet ID.
	QuotedTweetID string `json:"quotedTweetId"`
	// Public post and user referenced by this reaction.
	ReactionContext EmbeddedTweetRetweetedTweetQuotedTweetReactionContext `json:"reactionContext"`
	// Final nested tweet context at depth 4.
	RetweetedTweet EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweet `json:"retweeted_tweet"`
	// Public metadata whose fields are defined by X.
	Scopes map[string]any `json:"scopes"`
	Source string         `json:"source"`
	// Public visibility notice attached to an available tweet.
	Tombstone EmbeddedTweetRetweetedTweetQuotedTweetTombstone `json:"tombstone"`
	Type      string                                          `json:"type"`
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
func (r EmbeddedTweetRetweetedTweetQuotedTweet) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetRetweetedTweetQuotedTweet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Describes an X Article preview and its lifecycle metadata.
type EmbeddedTweetRetweetedTweetQuotedTweetArticle struct {
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
func (r EmbeddedTweetRetweetedTweetQuotedTweetArticle) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetRetweetedTweetQuotedTweetArticle) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Describes a public card and its referenced profiles.
type EmbeddedTweetRetweetedTweetQuotedTweetCard struct {
	ID string `json:"id"`
	// Public metadata whose fields are defined by X.
	BindingValues map[string]any `json:"bindingValues"`
	Name          string         `json:"name"`
	// Public metadata whose fields are defined by X.
	Platform map[string]any `json:"platform"`
	URL      string         `json:"url"`
	// Unresolved card user references.
	UserReferenceErrors []EmbeddedTweetRetweetedTweetQuotedTweetCardUserReferenceError `json:"userReferenceErrors"`
	UserReferences      []UserProfile                                                  `json:"userReferences"`
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
func (r EmbeddedTweetRetweetedTweetQuotedTweetCard) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetRetweetedTweetQuotedTweetCard) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetQuotedTweetCardUserReferenceError struct {
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
func (r EmbeddedTweetRetweetedTweetQuotedTweetCardUserReferenceError) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetQuotedTweetCardUserReferenceError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Community Note presentation metadata returned by X.
type EmbeddedTweetRetweetedTweetQuotedTweetCommunityNote struct {
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
func (r EmbeddedTweetRetweetedTweetQuotedTweetCommunityNote) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetRetweetedTweetQuotedTweetCommunityNote) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Public reply policy and conversation owner.
type EmbeddedTweetRetweetedTweetQuotedTweetConversationControl struct {
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
func (r EmbeddedTweetRetweetedTweetQuotedTweetConversationControl) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetQuotedTweetConversationControl) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Lists edit-chain identifiers and the remaining edit window.
type EmbeddedTweetRetweetedTweetQuotedTweetEdit struct {
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
func (r EmbeddedTweetRetweetedTweetQuotedTweetEdit) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetRetweetedTweetQuotedTweetEdit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Lists hashtags, symbols, links, and mentions from tweet text.
type EmbeddedTweetRetweetedTweetQuotedTweetEntities struct {
	Hashtags     []EmbeddedTweetRetweetedTweetQuotedTweetEntitiesHashtag     `json:"hashtags"`
	Smarttags    []EmbeddedTweetRetweetedTweetQuotedTweetEntitiesSmarttag    `json:"smarttags"`
	Symbols      []EmbeddedTweetRetweetedTweetQuotedTweetEntitiesSymbol      `json:"symbols"`
	Timestamps   []EmbeddedTweetRetweetedTweetQuotedTweetEntitiesTimestamp   `json:"timestamps"`
	URLs         []EmbeddedTweetRetweetedTweetQuotedTweetEntitiesURL         `json:"urls"`
	UserMentions []EmbeddedTweetRetweetedTweetQuotedTweetEntitiesUserMention `json:"user_mentions"`
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
func (r EmbeddedTweetRetweetedTweetQuotedTweetEntities) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetRetweetedTweetQuotedTweetEntities) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Provides hashtag text and source offsets within a tweet.
type EmbeddedTweetRetweetedTweetQuotedTweetEntitiesHashtag struct {
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
func (r EmbeddedTweetRetweetedTweetQuotedTweetEntitiesHashtag) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetRetweetedTweetQuotedTweetEntitiesHashtag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indexed smart-tag, cashtag, or video timestamp metadata.
type EmbeddedTweetRetweetedTweetQuotedTweetEntitiesSmarttag struct {
	Indices []int64                                                   `json:"indices"`
	Seconds float64                                                   `json:"seconds"`
	Tag     EmbeddedTweetRetweetedTweetQuotedTweetEntitiesSmarttagTag `json:"tag"`
	Text    string                                                    `json:"text"`
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
func (r EmbeddedTweetRetweetedTweetQuotedTweetEntitiesSmarttag) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetRetweetedTweetQuotedTweetEntitiesSmarttag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetQuotedTweetEntitiesSmarttagTag struct {
	Info EmbeddedTweetRetweetedTweetQuotedTweetEntitiesSmarttagTagInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetRetweetedTweetQuotedTweetEntitiesSmarttagTag) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetQuotedTweetEntitiesSmarttagTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetQuotedTweetEntitiesSmarttagTagInfo struct {
	Info EmbeddedTweetRetweetedTweetQuotedTweetEntitiesSmarttagTagInfoInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetRetweetedTweetQuotedTweetEntitiesSmarttagTagInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetQuotedTweetEntitiesSmarttagTagInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetQuotedTweetEntitiesSmarttagTagInfoInfo struct {
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
func (r EmbeddedTweetRetweetedTweetQuotedTweetEntitiesSmarttagTagInfoInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetQuotedTweetEntitiesSmarttagTagInfoInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indexed smart-tag, cashtag, or video timestamp metadata.
type EmbeddedTweetRetweetedTweetQuotedTweetEntitiesSymbol struct {
	Indices []int64                                                 `json:"indices"`
	Seconds float64                                                 `json:"seconds"`
	Tag     EmbeddedTweetRetweetedTweetQuotedTweetEntitiesSymbolTag `json:"tag"`
	Text    string                                                  `json:"text"`
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
func (r EmbeddedTweetRetweetedTweetQuotedTweetEntitiesSymbol) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetRetweetedTweetQuotedTweetEntitiesSymbol) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetQuotedTweetEntitiesSymbolTag struct {
	Info EmbeddedTweetRetweetedTweetQuotedTweetEntitiesSymbolTagInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetRetweetedTweetQuotedTweetEntitiesSymbolTag) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetRetweetedTweetQuotedTweetEntitiesSymbolTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetQuotedTweetEntitiesSymbolTagInfo struct {
	Info EmbeddedTweetRetweetedTweetQuotedTweetEntitiesSymbolTagInfoInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetRetweetedTweetQuotedTweetEntitiesSymbolTagInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetQuotedTweetEntitiesSymbolTagInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetQuotedTweetEntitiesSymbolTagInfoInfo struct {
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
func (r EmbeddedTweetRetweetedTweetQuotedTweetEntitiesSymbolTagInfoInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetQuotedTweetEntitiesSymbolTagInfoInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indexed smart-tag, cashtag, or video timestamp metadata.
type EmbeddedTweetRetweetedTweetQuotedTweetEntitiesTimestamp struct {
	Indices []int64                                                    `json:"indices"`
	Seconds float64                                                    `json:"seconds"`
	Tag     EmbeddedTweetRetweetedTweetQuotedTweetEntitiesTimestampTag `json:"tag"`
	Text    string                                                     `json:"text"`
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
func (r EmbeddedTweetRetweetedTweetQuotedTweetEntitiesTimestamp) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetRetweetedTweetQuotedTweetEntitiesTimestamp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetQuotedTweetEntitiesTimestampTag struct {
	Info EmbeddedTweetRetweetedTweetQuotedTweetEntitiesTimestampTagInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetRetweetedTweetQuotedTweetEntitiesTimestampTag) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetQuotedTweetEntitiesTimestampTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetQuotedTweetEntitiesTimestampTagInfo struct {
	Info EmbeddedTweetRetweetedTweetQuotedTweetEntitiesTimestampTagInfoInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetRetweetedTweetQuotedTweetEntitiesTimestampTagInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetQuotedTweetEntitiesTimestampTagInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetQuotedTweetEntitiesTimestampTagInfoInfo struct {
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
func (r EmbeddedTweetRetweetedTweetQuotedTweetEntitiesTimestampTagInfoInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetQuotedTweetEntitiesTimestampTagInfoInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Provides shortened, display, and expanded URLs from tweet text.
type EmbeddedTweetRetweetedTweetQuotedTweetEntitiesURL struct {
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
func (r EmbeddedTweetRetweetedTweetQuotedTweetEntitiesURL) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetRetweetedTweetQuotedTweetEntitiesURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Provides profile identity and source offsets for a mention.
type EmbeddedTweetRetweetedTweetQuotedTweetEntitiesUserMention struct {
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
func (r EmbeddedTweetRetweetedTweetQuotedTweetEntitiesUserMention) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetQuotedTweetEntitiesUserMention) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetQuotedTweetLimitedAction struct {
	Action string                                                    `json:"action"`
	Prompt EmbeddedTweetRetweetedTweetQuotedTweetLimitedActionPrompt `json:"prompt"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Action      respjson.Field
		Prompt      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetRetweetedTweetQuotedTweetLimitedAction) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetRetweetedTweetQuotedTweetLimitedAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetQuotedTweetLimitedActionPrompt struct {
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
func (r EmbeddedTweetRetweetedTweetQuotedTweetLimitedActionPrompt) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetQuotedTweetLimitedActionPrompt) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Complete Note Tweet content and rich-text metadata.
type EmbeddedTweetRetweetedTweetQuotedTweetNoteTweet struct {
	Text string `json:"text" api:"required"`
	ID   string `json:"id"`
	// Lists hashtags, symbols, links, and mentions from tweet text.
	Entities EmbeddedTweetRetweetedTweetQuotedTweetNoteTweetEntities `json:"entities"`
	// Inline media positions in the Note Tweet text.
	InlineMedia  []EmbeddedTweetRetweetedTweetQuotedTweetNoteTweetInlineMedia `json:"inlineMedia"`
	IsExpandable bool                                                         `json:"isExpandable"`
	RichtextTags []EmbeddedTweetRetweetedTweetQuotedTweetNoteTweetRichtextTag `json:"richtextTags"`
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
func (r EmbeddedTweetRetweetedTweetQuotedTweetNoteTweet) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetRetweetedTweetQuotedTweetNoteTweet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Lists hashtags, symbols, links, and mentions from tweet text.
type EmbeddedTweetRetweetedTweetQuotedTweetNoteTweetEntities struct {
	Hashtags     []EmbeddedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesHashtag     `json:"hashtags"`
	Smarttags    []EmbeddedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesSmarttag    `json:"smarttags"`
	Symbols      []EmbeddedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesSymbol      `json:"symbols"`
	Timestamps   []EmbeddedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesTimestamp   `json:"timestamps"`
	URLs         []EmbeddedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesURL         `json:"urls"`
	UserMentions []EmbeddedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesUserMention `json:"user_mentions"`
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
func (r EmbeddedTweetRetweetedTweetQuotedTweetNoteTweetEntities) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetRetweetedTweetQuotedTweetNoteTweetEntities) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Provides hashtag text and source offsets within a tweet.
type EmbeddedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesHashtag struct {
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
func (r EmbeddedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesHashtag) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesHashtag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indexed smart-tag, cashtag, or video timestamp metadata.
type EmbeddedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesSmarttag struct {
	Indices []int64                                                            `json:"indices"`
	Seconds float64                                                            `json:"seconds"`
	Tag     EmbeddedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesSmarttagTag `json:"tag"`
	Text    string                                                             `json:"text"`
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
func (r EmbeddedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesSmarttag) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesSmarttag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesSmarttagTag struct {
	Info EmbeddedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesSmarttagTagInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesSmarttagTag) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesSmarttagTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesSmarttagTagInfo struct {
	Info EmbeddedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesSmarttagTagInfoInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesSmarttagTagInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesSmarttagTagInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesSmarttagTagInfoInfo struct {
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
func (r EmbeddedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesSmarttagTagInfoInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesSmarttagTagInfoInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indexed smart-tag, cashtag, or video timestamp metadata.
type EmbeddedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesSymbol struct {
	Indices []int64                                                          `json:"indices"`
	Seconds float64                                                          `json:"seconds"`
	Tag     EmbeddedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesSymbolTag `json:"tag"`
	Text    string                                                           `json:"text"`
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
func (r EmbeddedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesSymbol) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesSymbol) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesSymbolTag struct {
	Info EmbeddedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesSymbolTagInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesSymbolTag) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesSymbolTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesSymbolTagInfo struct {
	Info EmbeddedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesSymbolTagInfoInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesSymbolTagInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesSymbolTagInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesSymbolTagInfoInfo struct {
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
func (r EmbeddedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesSymbolTagInfoInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesSymbolTagInfoInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indexed smart-tag, cashtag, or video timestamp metadata.
type EmbeddedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesTimestamp struct {
	Indices []int64                                                             `json:"indices"`
	Seconds float64                                                             `json:"seconds"`
	Tag     EmbeddedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesTimestampTag `json:"tag"`
	Text    string                                                              `json:"text"`
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
func (r EmbeddedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesTimestamp) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesTimestamp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesTimestampTag struct {
	Info EmbeddedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesTimestampTagInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesTimestampTag) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesTimestampTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesTimestampTagInfo struct {
	Info EmbeddedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesTimestampTagInfoInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesTimestampTagInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesTimestampTagInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesTimestampTagInfoInfo struct {
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
func (r EmbeddedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesTimestampTagInfoInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesTimestampTagInfoInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Provides shortened, display, and expanded URLs from tweet text.
type EmbeddedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesURL struct {
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
func (r EmbeddedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesURL) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Provides profile identity and source offsets for a mention.
type EmbeddedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesUserMention struct {
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
func (r EmbeddedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesUserMention) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesUserMention) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetQuotedTweetNoteTweetInlineMedia struct {
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
func (r EmbeddedTweetRetweetedTweetQuotedTweetNoteTweetInlineMedia) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetQuotedTweetNoteTweetInlineMedia) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetQuotedTweetNoteTweetRichtextTag struct {
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
func (r EmbeddedTweetRetweetedTweetQuotedTweetNoteTweetRichtextTag) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetQuotedTweetNoteTweetRichtextTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Describes public place metadata on a geotagged tweet.
type EmbeddedTweetRetweetedTweetQuotedTweetPlace struct {
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
func (r EmbeddedTweetRetweetedTweetQuotedTweetPlace) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetRetweetedTweetQuotedTweetPlace) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Engagement counts retained from a prior tweet edit.
type EmbeddedTweetRetweetedTweetQuotedTweetPreviousCounts struct {
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
func (r EmbeddedTweetRetweetedTweetQuotedTweetPreviousCounts) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetRetweetedTweetQuotedTweetPreviousCounts) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Final nested tweet context at depth 4.
type EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweet struct {
	ID            string `json:"id" api:"required"`
	BookmarkCount int64  `json:"bookmarkCount" api:"required"`
	LikeCount     int64  `json:"likeCount" api:"required"`
	QuoteCount    int64  `json:"quoteCount" api:"required"`
	ReplyCount    int64  `json:"replyCount" api:"required"`
	RetweetCount  int64  `json:"retweetCount" api:"required"`
	Text          string `json:"text" api:"required"`
	ViewCount     int64  `json:"viewCount" api:"required"`
	// Describes an X Article preview and its lifecycle metadata.
	Article EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetArticle `json:"article"`
	// X user profile with bio, follower counts, and verification status.
	Author UserProfile `json:"author"`
	// Describes a public card and its referenced profiles.
	Card EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetCard `json:"card"`
	// Community ID.
	CommunityID string `json:"communityId"`
	// Community Note presentation metadata returned by X.
	CommunityNote EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetCommunityNote `json:"communityNote"`
	// Content disclosure metadata shown by X when a tweet is labeled as paid
	// partnership content or AI-generated media.
	ContentDisclosure ContentDisclosure `json:"contentDisclosure"`
	// Public reply policy and conversation owner.
	ConversationControl EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetConversationControl `json:"conversationControl"`
	ConversationID      string                                                               `json:"conversationId"`
	CreatedAt           string                                                               `json:"createdAt"`
	DisplayTextRange    []int64                                                              `json:"displayTextRange"`
	// Lists edit-chain identifiers and the remaining edit window.
	Edit EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetEdit `json:"edit"`
	// Lists hashtags, symbols, links, and mentions from tweet text.
	Entities EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetEntities `json:"entities"`
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
	LimitedActions []EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetLimitedAction `json:"limitedActions"`
	// Attached media items, omitted when unavailable.
	Media []TweetMedia `json:"media"`
	// Complete Note Tweet content and rich-text metadata.
	NoteTweet EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetNoteTweet `json:"noteTweet"`
	// Describes public place metadata on a geotagged tweet.
	Place             EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetPlace `json:"place"`
	PossiblySensitive bool                                                   `json:"possiblySensitive"`
	// Public metadata whose fields are defined by X.
	PostCta map[string]any `json:"postCta"`
	// Engagement counts retained from a prior tweet edit.
	PreviousCounts EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetPreviousCounts `json:"previousCounts"`
	// Quoted tweet ID.
	QuotedTweetID string `json:"quotedTweetId"`
	// Public post and user referenced by this reaction.
	ReactionContext EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetReactionContext `json:"reactionContext"`
	// Public metadata whose fields are defined by X.
	Scopes map[string]any `json:"scopes"`
	Source string         `json:"source"`
	// Public visibility notice attached to an available tweet.
	Tombstone EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetTombstone `json:"tombstone"`
	Type      string                                                     `json:"type"`
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
		QuotedTweetID       respjson.Field
		ReactionContext     respjson.Field
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
func (r EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweet) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Describes an X Article preview and its lifecycle metadata.
type EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetArticle struct {
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
func (r EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetArticle) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetArticle) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Describes a public card and its referenced profiles.
type EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetCard struct {
	ID string `json:"id"`
	// Public metadata whose fields are defined by X.
	BindingValues map[string]any `json:"bindingValues"`
	Name          string         `json:"name"`
	// Public metadata whose fields are defined by X.
	Platform map[string]any `json:"platform"`
	URL      string         `json:"url"`
	// Unresolved card user references.
	UserReferenceErrors []EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetCardUserReferenceError `json:"userReferenceErrors"`
	UserReferences      []UserProfile                                                             `json:"userReferences"`
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
func (r EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetCard) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetCard) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetCardUserReferenceError struct {
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
func (r EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetCardUserReferenceError) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetCardUserReferenceError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Community Note presentation metadata returned by X.
type EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetCommunityNote struct {
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
func (r EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetCommunityNote) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetCommunityNote) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Public reply policy and conversation owner.
type EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetConversationControl struct {
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
func (r EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetConversationControl) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetConversationControl) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Lists edit-chain identifiers and the remaining edit window.
type EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetEdit struct {
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
func (r EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetEdit) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetEdit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Lists hashtags, symbols, links, and mentions from tweet text.
type EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetEntities struct {
	Hashtags     []EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetEntitiesHashtag     `json:"hashtags"`
	Smarttags    []EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetEntitiesSmarttag    `json:"smarttags"`
	Symbols      []EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetEntitiesSymbol      `json:"symbols"`
	Timestamps   []EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetEntitiesTimestamp   `json:"timestamps"`
	URLs         []EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetEntitiesURL         `json:"urls"`
	UserMentions []EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetEntitiesUserMention `json:"user_mentions"`
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
func (r EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetEntities) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetEntities) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Provides hashtag text and source offsets within a tweet.
type EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetEntitiesHashtag struct {
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
func (r EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetEntitiesHashtag) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetEntitiesHashtag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indexed smart-tag, cashtag, or video timestamp metadata.
type EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetEntitiesSmarttag struct {
	Indices []int64                                                              `json:"indices"`
	Seconds float64                                                              `json:"seconds"`
	Tag     EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetEntitiesSmarttagTag `json:"tag"`
	Text    string                                                               `json:"text"`
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
func (r EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetEntitiesSmarttag) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetEntitiesSmarttag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetEntitiesSmarttagTag struct {
	Info EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetEntitiesSmarttagTagInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetEntitiesSmarttagTag) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetEntitiesSmarttagTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetEntitiesSmarttagTagInfo struct {
	Info EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetEntitiesSmarttagTagInfoInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetEntitiesSmarttagTagInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetEntitiesSmarttagTagInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetEntitiesSmarttagTagInfoInfo struct {
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
func (r EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetEntitiesSmarttagTagInfoInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetEntitiesSmarttagTagInfoInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indexed smart-tag, cashtag, or video timestamp metadata.
type EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetEntitiesSymbol struct {
	Indices []int64                                                            `json:"indices"`
	Seconds float64                                                            `json:"seconds"`
	Tag     EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetEntitiesSymbolTag `json:"tag"`
	Text    string                                                             `json:"text"`
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
func (r EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetEntitiesSymbol) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetEntitiesSymbol) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetEntitiesSymbolTag struct {
	Info EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetEntitiesSymbolTagInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetEntitiesSymbolTag) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetEntitiesSymbolTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetEntitiesSymbolTagInfo struct {
	Info EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetEntitiesSymbolTagInfoInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetEntitiesSymbolTagInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetEntitiesSymbolTagInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetEntitiesSymbolTagInfoInfo struct {
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
func (r EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetEntitiesSymbolTagInfoInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetEntitiesSymbolTagInfoInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indexed smart-tag, cashtag, or video timestamp metadata.
type EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetEntitiesTimestamp struct {
	Indices []int64                                                               `json:"indices"`
	Seconds float64                                                               `json:"seconds"`
	Tag     EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetEntitiesTimestampTag `json:"tag"`
	Text    string                                                                `json:"text"`
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
func (r EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetEntitiesTimestamp) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetEntitiesTimestamp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetEntitiesTimestampTag struct {
	Info EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetEntitiesTimestampTagInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetEntitiesTimestampTag) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetEntitiesTimestampTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetEntitiesTimestampTagInfo struct {
	Info EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetEntitiesTimestampTagInfoInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetEntitiesTimestampTagInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetEntitiesTimestampTagInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetEntitiesTimestampTagInfoInfo struct {
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
func (r EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetEntitiesTimestampTagInfoInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetEntitiesTimestampTagInfoInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Provides shortened, display, and expanded URLs from tweet text.
type EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetEntitiesURL struct {
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
func (r EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetEntitiesURL) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetEntitiesURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Provides profile identity and source offsets for a mention.
type EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetEntitiesUserMention struct {
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
func (r EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetEntitiesUserMention) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetEntitiesUserMention) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetLimitedAction struct {
	Action string                                                               `json:"action"`
	Prompt EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetLimitedActionPrompt `json:"prompt"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Action      respjson.Field
		Prompt      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetLimitedAction) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetLimitedAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetLimitedActionPrompt struct {
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
func (r EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetLimitedActionPrompt) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetLimitedActionPrompt) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Complete Note Tweet content and rich-text metadata.
type EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetNoteTweet struct {
	Text string `json:"text" api:"required"`
	ID   string `json:"id"`
	// Lists hashtags, symbols, links, and mentions from tweet text.
	Entities EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetNoteTweetEntities `json:"entities"`
	// Inline media positions in the Note Tweet text.
	InlineMedia  []EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetNoteTweetInlineMedia `json:"inlineMedia"`
	IsExpandable bool                                                                    `json:"isExpandable"`
	RichtextTags []EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetNoteTweetRichtextTag `json:"richtextTags"`
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
func (r EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetNoteTweet) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetNoteTweet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Lists hashtags, symbols, links, and mentions from tweet text.
type EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetNoteTweetEntities struct {
	Hashtags     []EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetNoteTweetEntitiesHashtag     `json:"hashtags"`
	Smarttags    []EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetNoteTweetEntitiesSmarttag    `json:"smarttags"`
	Symbols      []EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetNoteTweetEntitiesSymbol      `json:"symbols"`
	Timestamps   []EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetNoteTweetEntitiesTimestamp   `json:"timestamps"`
	URLs         []EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetNoteTweetEntitiesURL         `json:"urls"`
	UserMentions []EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetNoteTweetEntitiesUserMention `json:"user_mentions"`
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
func (r EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetNoteTweetEntities) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetNoteTweetEntities) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Provides hashtag text and source offsets within a tweet.
type EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetNoteTweetEntitiesHashtag struct {
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
func (r EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetNoteTweetEntitiesHashtag) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetNoteTweetEntitiesHashtag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indexed smart-tag, cashtag, or video timestamp metadata.
type EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetNoteTweetEntitiesSmarttag struct {
	Indices []int64                                                                       `json:"indices"`
	Seconds float64                                                                       `json:"seconds"`
	Tag     EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetNoteTweetEntitiesSmarttagTag `json:"tag"`
	Text    string                                                                        `json:"text"`
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
func (r EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetNoteTweetEntitiesSmarttag) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetNoteTweetEntitiesSmarttag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetNoteTweetEntitiesSmarttagTag struct {
	Info EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetNoteTweetEntitiesSmarttagTagInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetNoteTweetEntitiesSmarttagTag) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetNoteTweetEntitiesSmarttagTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetNoteTweetEntitiesSmarttagTagInfo struct {
	Info EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetNoteTweetEntitiesSmarttagTagInfoInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetNoteTweetEntitiesSmarttagTagInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetNoteTweetEntitiesSmarttagTagInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetNoteTweetEntitiesSmarttagTagInfoInfo struct {
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
func (r EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetNoteTweetEntitiesSmarttagTagInfoInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetNoteTweetEntitiesSmarttagTagInfoInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indexed smart-tag, cashtag, or video timestamp metadata.
type EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetNoteTweetEntitiesSymbol struct {
	Indices []int64                                                                     `json:"indices"`
	Seconds float64                                                                     `json:"seconds"`
	Tag     EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetNoteTweetEntitiesSymbolTag `json:"tag"`
	Text    string                                                                      `json:"text"`
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
func (r EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetNoteTweetEntitiesSymbol) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetNoteTweetEntitiesSymbol) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetNoteTweetEntitiesSymbolTag struct {
	Info EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetNoteTweetEntitiesSymbolTagInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetNoteTweetEntitiesSymbolTag) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetNoteTweetEntitiesSymbolTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetNoteTweetEntitiesSymbolTagInfo struct {
	Info EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetNoteTweetEntitiesSymbolTagInfoInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetNoteTweetEntitiesSymbolTagInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetNoteTweetEntitiesSymbolTagInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetNoteTweetEntitiesSymbolTagInfoInfo struct {
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
func (r EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetNoteTweetEntitiesSymbolTagInfoInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetNoteTweetEntitiesSymbolTagInfoInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indexed smart-tag, cashtag, or video timestamp metadata.
type EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetNoteTweetEntitiesTimestamp struct {
	Indices []int64                                                                        `json:"indices"`
	Seconds float64                                                                        `json:"seconds"`
	Tag     EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetNoteTweetEntitiesTimestampTag `json:"tag"`
	Text    string                                                                         `json:"text"`
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
func (r EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetNoteTweetEntitiesTimestamp) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetNoteTweetEntitiesTimestamp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetNoteTweetEntitiesTimestampTag struct {
	Info EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetNoteTweetEntitiesTimestampTagInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetNoteTweetEntitiesTimestampTag) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetNoteTweetEntitiesTimestampTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetNoteTweetEntitiesTimestampTagInfo struct {
	Info EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetNoteTweetEntitiesTimestampTagInfoInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetNoteTweetEntitiesTimestampTagInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetNoteTweetEntitiesTimestampTagInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetNoteTweetEntitiesTimestampTagInfoInfo struct {
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
func (r EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetNoteTweetEntitiesTimestampTagInfoInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetNoteTweetEntitiesTimestampTagInfoInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Provides shortened, display, and expanded URLs from tweet text.
type EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetNoteTweetEntitiesURL struct {
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
func (r EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetNoteTweetEntitiesURL) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetNoteTweetEntitiesURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Provides profile identity and source offsets for a mention.
type EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetNoteTweetEntitiesUserMention struct {
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
func (r EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetNoteTweetEntitiesUserMention) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetNoteTweetEntitiesUserMention) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetNoteTweetInlineMedia struct {
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
func (r EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetNoteTweetInlineMedia) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetNoteTweetInlineMedia) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetNoteTweetRichtextTag struct {
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
func (r EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetNoteTweetRichtextTag) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetNoteTweetRichtextTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Describes public place metadata on a geotagged tweet.
type EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetPlace struct {
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
func (r EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetPlace) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetPlace) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Engagement counts retained from a prior tweet edit.
type EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetPreviousCounts struct {
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
func (r EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetPreviousCounts) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetPreviousCounts) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Public post and user referenced by this reaction.
type EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetReactionContext struct {
	// Referenced post ID.
	ReactedToPostID string `json:"reactedToPostId"`
	// X user profile with bio, follower counts, and verification status.
	ReactedToUser UserProfile `json:"reactedToUser"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ReactedToPostID respjson.Field
		ReactedToUser   respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetReactionContext) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetReactionContext) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Public visibility notice attached to an available tweet.
type EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetTombstone struct {
	Text EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetTombstoneText `json:"text"`
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
func (r EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetTombstone) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetTombstone) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetTombstoneText struct {
	Entities []EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetTombstoneTextEntity `json:"entities"`
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
func (r EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetTombstoneText) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetTombstoneText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetTombstoneTextEntity struct {
	FromIndex int64                                                                   `json:"fromIndex"`
	Ref       EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetTombstoneTextEntityRef `json:"ref"`
	ToIndex   int64                                                                   `json:"toIndex"`
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
func (r EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetTombstoneTextEntity) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetTombstoneTextEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetTombstoneTextEntityRef struct {
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
func (r EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetTombstoneTextEntityRef) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetTombstoneTextEntityRef) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Public post and user referenced by this reaction.
type EmbeddedTweetRetweetedTweetQuotedTweetReactionContext struct {
	// Referenced post ID.
	ReactedToPostID string `json:"reactedToPostId"`
	// X user profile with bio, follower counts, and verification status.
	ReactedToUser UserProfile `json:"reactedToUser"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ReactedToPostID respjson.Field
		ReactedToUser   respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetRetweetedTweetQuotedTweetReactionContext) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetRetweetedTweetQuotedTweetReactionContext) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Final nested tweet context at depth 4.
type EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweet struct {
	ID            string `json:"id" api:"required"`
	BookmarkCount int64  `json:"bookmarkCount" api:"required"`
	LikeCount     int64  `json:"likeCount" api:"required"`
	QuoteCount    int64  `json:"quoteCount" api:"required"`
	ReplyCount    int64  `json:"replyCount" api:"required"`
	RetweetCount  int64  `json:"retweetCount" api:"required"`
	Text          string `json:"text" api:"required"`
	ViewCount     int64  `json:"viewCount" api:"required"`
	// Describes an X Article preview and its lifecycle metadata.
	Article EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetArticle `json:"article"`
	// X user profile with bio, follower counts, and verification status.
	Author UserProfile `json:"author"`
	// Describes a public card and its referenced profiles.
	Card EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetCard `json:"card"`
	// Community ID.
	CommunityID string `json:"communityId"`
	// Community Note presentation metadata returned by X.
	CommunityNote EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetCommunityNote `json:"communityNote"`
	// Content disclosure metadata shown by X when a tweet is labeled as paid
	// partnership content or AI-generated media.
	ContentDisclosure ContentDisclosure `json:"contentDisclosure"`
	// Public reply policy and conversation owner.
	ConversationControl EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetConversationControl `json:"conversationControl"`
	ConversationID      string                                                                  `json:"conversationId"`
	CreatedAt           string                                                                  `json:"createdAt"`
	DisplayTextRange    []int64                                                                 `json:"displayTextRange"`
	// Lists edit-chain identifiers and the remaining edit window.
	Edit EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetEdit `json:"edit"`
	// Lists hashtags, symbols, links, and mentions from tweet text.
	Entities EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetEntities `json:"entities"`
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
	LimitedActions []EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetLimitedAction `json:"limitedActions"`
	// Attached media items, omitted when unavailable.
	Media []TweetMedia `json:"media"`
	// Complete Note Tweet content and rich-text metadata.
	NoteTweet EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetNoteTweet `json:"noteTweet"`
	// Describes public place metadata on a geotagged tweet.
	Place             EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetPlace `json:"place"`
	PossiblySensitive bool                                                      `json:"possiblySensitive"`
	// Public metadata whose fields are defined by X.
	PostCta map[string]any `json:"postCta"`
	// Engagement counts retained from a prior tweet edit.
	PreviousCounts EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetPreviousCounts `json:"previousCounts"`
	// Quoted tweet ID.
	QuotedTweetID string `json:"quotedTweetId"`
	// Public post and user referenced by this reaction.
	ReactionContext EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetReactionContext `json:"reactionContext"`
	// Public metadata whose fields are defined by X.
	Scopes map[string]any `json:"scopes"`
	Source string         `json:"source"`
	// Public visibility notice attached to an available tweet.
	Tombstone EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetTombstone `json:"tombstone"`
	Type      string                                                        `json:"type"`
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
		QuotedTweetID       respjson.Field
		ReactionContext     respjson.Field
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
func (r EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweet) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Describes an X Article preview and its lifecycle metadata.
type EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetArticle struct {
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
func (r EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetArticle) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetArticle) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Describes a public card and its referenced profiles.
type EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetCard struct {
	ID string `json:"id"`
	// Public metadata whose fields are defined by X.
	BindingValues map[string]any `json:"bindingValues"`
	Name          string         `json:"name"`
	// Public metadata whose fields are defined by X.
	Platform map[string]any `json:"platform"`
	URL      string         `json:"url"`
	// Unresolved card user references.
	UserReferenceErrors []EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetCardUserReferenceError `json:"userReferenceErrors"`
	UserReferences      []UserProfile                                                                `json:"userReferences"`
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
func (r EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetCard) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetCard) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetCardUserReferenceError struct {
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
func (r EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetCardUserReferenceError) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetCardUserReferenceError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Community Note presentation metadata returned by X.
type EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetCommunityNote struct {
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
func (r EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetCommunityNote) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetCommunityNote) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Public reply policy and conversation owner.
type EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetConversationControl struct {
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
func (r EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetConversationControl) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetConversationControl) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Lists edit-chain identifiers and the remaining edit window.
type EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetEdit struct {
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
func (r EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetEdit) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetEdit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Lists hashtags, symbols, links, and mentions from tweet text.
type EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetEntities struct {
	Hashtags     []EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetEntitiesHashtag     `json:"hashtags"`
	Smarttags    []EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetEntitiesSmarttag    `json:"smarttags"`
	Symbols      []EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetEntitiesSymbol      `json:"symbols"`
	Timestamps   []EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetEntitiesTimestamp   `json:"timestamps"`
	URLs         []EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetEntitiesURL         `json:"urls"`
	UserMentions []EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetEntitiesUserMention `json:"user_mentions"`
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
func (r EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetEntities) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetEntities) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Provides hashtag text and source offsets within a tweet.
type EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetEntitiesHashtag struct {
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
func (r EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetEntitiesHashtag) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetEntitiesHashtag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indexed smart-tag, cashtag, or video timestamp metadata.
type EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetEntitiesSmarttag struct {
	Indices []int64                                                                 `json:"indices"`
	Seconds float64                                                                 `json:"seconds"`
	Tag     EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetEntitiesSmarttagTag `json:"tag"`
	Text    string                                                                  `json:"text"`
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
func (r EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetEntitiesSmarttag) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetEntitiesSmarttag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetEntitiesSmarttagTag struct {
	Info EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetEntitiesSmarttagTagInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetEntitiesSmarttagTag) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetEntitiesSmarttagTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetEntitiesSmarttagTagInfo struct {
	Info EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetEntitiesSmarttagTagInfoInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetEntitiesSmarttagTagInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetEntitiesSmarttagTagInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetEntitiesSmarttagTagInfoInfo struct {
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
func (r EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetEntitiesSmarttagTagInfoInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetEntitiesSmarttagTagInfoInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indexed smart-tag, cashtag, or video timestamp metadata.
type EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetEntitiesSymbol struct {
	Indices []int64                                                               `json:"indices"`
	Seconds float64                                                               `json:"seconds"`
	Tag     EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetEntitiesSymbolTag `json:"tag"`
	Text    string                                                                `json:"text"`
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
func (r EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetEntitiesSymbol) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetEntitiesSymbol) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetEntitiesSymbolTag struct {
	Info EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetEntitiesSymbolTagInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetEntitiesSymbolTag) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetEntitiesSymbolTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetEntitiesSymbolTagInfo struct {
	Info EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetEntitiesSymbolTagInfoInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetEntitiesSymbolTagInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetEntitiesSymbolTagInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetEntitiesSymbolTagInfoInfo struct {
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
func (r EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetEntitiesSymbolTagInfoInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetEntitiesSymbolTagInfoInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indexed smart-tag, cashtag, or video timestamp metadata.
type EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetEntitiesTimestamp struct {
	Indices []int64                                                                  `json:"indices"`
	Seconds float64                                                                  `json:"seconds"`
	Tag     EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetEntitiesTimestampTag `json:"tag"`
	Text    string                                                                   `json:"text"`
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
func (r EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetEntitiesTimestamp) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetEntitiesTimestamp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetEntitiesTimestampTag struct {
	Info EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetEntitiesTimestampTagInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetEntitiesTimestampTag) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetEntitiesTimestampTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetEntitiesTimestampTagInfo struct {
	Info EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetEntitiesTimestampTagInfoInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetEntitiesTimestampTagInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetEntitiesTimestampTagInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetEntitiesTimestampTagInfoInfo struct {
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
func (r EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetEntitiesTimestampTagInfoInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetEntitiesTimestampTagInfoInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Provides shortened, display, and expanded URLs from tweet text.
type EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetEntitiesURL struct {
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
func (r EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetEntitiesURL) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetEntitiesURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Provides profile identity and source offsets for a mention.
type EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetEntitiesUserMention struct {
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
func (r EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetEntitiesUserMention) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetEntitiesUserMention) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetLimitedAction struct {
	Action string                                                                  `json:"action"`
	Prompt EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetLimitedActionPrompt `json:"prompt"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Action      respjson.Field
		Prompt      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetLimitedAction) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetLimitedAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetLimitedActionPrompt struct {
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
func (r EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetLimitedActionPrompt) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetLimitedActionPrompt) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Complete Note Tweet content and rich-text metadata.
type EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetNoteTweet struct {
	Text string `json:"text" api:"required"`
	ID   string `json:"id"`
	// Lists hashtags, symbols, links, and mentions from tweet text.
	Entities EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetNoteTweetEntities `json:"entities"`
	// Inline media positions in the Note Tweet text.
	InlineMedia  []EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetNoteTweetInlineMedia `json:"inlineMedia"`
	IsExpandable bool                                                                       `json:"isExpandable"`
	RichtextTags []EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetNoteTweetRichtextTag `json:"richtextTags"`
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
func (r EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetNoteTweet) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetNoteTweet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Lists hashtags, symbols, links, and mentions from tweet text.
type EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetNoteTweetEntities struct {
	Hashtags     []EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesHashtag     `json:"hashtags"`
	Smarttags    []EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesSmarttag    `json:"smarttags"`
	Symbols      []EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesSymbol      `json:"symbols"`
	Timestamps   []EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesTimestamp   `json:"timestamps"`
	URLs         []EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesURL         `json:"urls"`
	UserMentions []EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesUserMention `json:"user_mentions"`
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
func (r EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetNoteTweetEntities) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetNoteTweetEntities) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Provides hashtag text and source offsets within a tweet.
type EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesHashtag struct {
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
func (r EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesHashtag) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesHashtag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indexed smart-tag, cashtag, or video timestamp metadata.
type EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesSmarttag struct {
	Indices []int64                                                                          `json:"indices"`
	Seconds float64                                                                          `json:"seconds"`
	Tag     EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesSmarttagTag `json:"tag"`
	Text    string                                                                           `json:"text"`
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
func (r EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesSmarttag) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesSmarttag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesSmarttagTag struct {
	Info EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesSmarttagTagInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesSmarttagTag) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesSmarttagTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesSmarttagTagInfo struct {
	Info EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesSmarttagTagInfoInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesSmarttagTagInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesSmarttagTagInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesSmarttagTagInfoInfo struct {
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
func (r EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesSmarttagTagInfoInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesSmarttagTagInfoInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indexed smart-tag, cashtag, or video timestamp metadata.
type EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesSymbol struct {
	Indices []int64                                                                        `json:"indices"`
	Seconds float64                                                                        `json:"seconds"`
	Tag     EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesSymbolTag `json:"tag"`
	Text    string                                                                         `json:"text"`
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
func (r EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesSymbol) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesSymbol) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesSymbolTag struct {
	Info EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesSymbolTagInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesSymbolTag) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesSymbolTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesSymbolTagInfo struct {
	Info EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesSymbolTagInfoInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesSymbolTagInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesSymbolTagInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesSymbolTagInfoInfo struct {
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
func (r EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesSymbolTagInfoInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesSymbolTagInfoInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indexed smart-tag, cashtag, or video timestamp metadata.
type EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesTimestamp struct {
	Indices []int64                                                                           `json:"indices"`
	Seconds float64                                                                           `json:"seconds"`
	Tag     EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesTimestampTag `json:"tag"`
	Text    string                                                                            `json:"text"`
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
func (r EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesTimestamp) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesTimestamp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesTimestampTag struct {
	Info EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesTimestampTagInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesTimestampTag) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesTimestampTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesTimestampTagInfo struct {
	Info EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesTimestampTagInfoInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesTimestampTagInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesTimestampTagInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesTimestampTagInfoInfo struct {
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
func (r EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesTimestampTagInfoInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesTimestampTagInfoInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Provides shortened, display, and expanded URLs from tweet text.
type EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesURL struct {
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
func (r EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesURL) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Provides profile identity and source offsets for a mention.
type EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesUserMention struct {
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
func (r EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesUserMention) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesUserMention) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetNoteTweetInlineMedia struct {
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
func (r EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetNoteTweetInlineMedia) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetNoteTweetInlineMedia) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetNoteTweetRichtextTag struct {
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
func (r EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetNoteTweetRichtextTag) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetNoteTweetRichtextTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Describes public place metadata on a geotagged tweet.
type EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetPlace struct {
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
func (r EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetPlace) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetPlace) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Engagement counts retained from a prior tweet edit.
type EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetPreviousCounts struct {
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
func (r EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetPreviousCounts) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetPreviousCounts) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Public post and user referenced by this reaction.
type EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetReactionContext struct {
	// Referenced post ID.
	ReactedToPostID string `json:"reactedToPostId"`
	// X user profile with bio, follower counts, and verification status.
	ReactedToUser UserProfile `json:"reactedToUser"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ReactedToPostID respjson.Field
		ReactedToUser   respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetReactionContext) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetReactionContext) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Public visibility notice attached to an available tweet.
type EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetTombstone struct {
	Text EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetTombstoneText `json:"text"`
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
func (r EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetTombstone) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetTombstone) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetTombstoneText struct {
	Entities []EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetTombstoneTextEntity `json:"entities"`
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
func (r EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetTombstoneText) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetTombstoneText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetTombstoneTextEntity struct {
	FromIndex int64                                                                      `json:"fromIndex"`
	Ref       EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetTombstoneTextEntityRef `json:"ref"`
	ToIndex   int64                                                                      `json:"toIndex"`
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
func (r EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetTombstoneTextEntity) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetTombstoneTextEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetTombstoneTextEntityRef struct {
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
func (r EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetTombstoneTextEntityRef) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetTombstoneTextEntityRef) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Public visibility notice attached to an available tweet.
type EmbeddedTweetRetweetedTweetQuotedTweetTombstone struct {
	Text EmbeddedTweetRetweetedTweetQuotedTweetTombstoneText `json:"text"`
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
func (r EmbeddedTweetRetweetedTweetQuotedTweetTombstone) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetRetweetedTweetQuotedTweetTombstone) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetQuotedTweetTombstoneText struct {
	Entities []EmbeddedTweetRetweetedTweetQuotedTweetTombstoneTextEntity `json:"entities"`
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
func (r EmbeddedTweetRetweetedTweetQuotedTweetTombstoneText) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetRetweetedTweetQuotedTweetTombstoneText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetQuotedTweetTombstoneTextEntity struct {
	FromIndex int64                                                        `json:"fromIndex"`
	Ref       EmbeddedTweetRetweetedTweetQuotedTweetTombstoneTextEntityRef `json:"ref"`
	ToIndex   int64                                                        `json:"toIndex"`
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
func (r EmbeddedTweetRetweetedTweetQuotedTweetTombstoneTextEntity) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetQuotedTweetTombstoneTextEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetQuotedTweetTombstoneTextEntityRef struct {
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
func (r EmbeddedTweetRetweetedTweetQuotedTweetTombstoneTextEntityRef) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetQuotedTweetTombstoneTextEntityRef) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Public post and user referenced by this reaction.
type EmbeddedTweetRetweetedTweetReactionContext struct {
	// Referenced post ID.
	ReactedToPostID string `json:"reactedToPostId"`
	// X user profile with bio, follower counts, and verification status.
	ReactedToUser UserProfile `json:"reactedToUser"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ReactedToPostID respjson.Field
		ReactedToUser   respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetRetweetedTweetReactionContext) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetRetweetedTweetReactionContext) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Nested tweet context at depth 3.
type EmbeddedTweetRetweetedTweetRetweetedTweet struct {
	ID            string `json:"id" api:"required"`
	BookmarkCount int64  `json:"bookmarkCount" api:"required"`
	LikeCount     int64  `json:"likeCount" api:"required"`
	QuoteCount    int64  `json:"quoteCount" api:"required"`
	ReplyCount    int64  `json:"replyCount" api:"required"`
	RetweetCount  int64  `json:"retweetCount" api:"required"`
	Text          string `json:"text" api:"required"`
	ViewCount     int64  `json:"viewCount" api:"required"`
	// Describes an X Article preview and its lifecycle metadata.
	Article EmbeddedTweetRetweetedTweetRetweetedTweetArticle `json:"article"`
	// X user profile with bio, follower counts, and verification status.
	Author UserProfile `json:"author"`
	// Describes a public card and its referenced profiles.
	Card EmbeddedTweetRetweetedTweetRetweetedTweetCard `json:"card"`
	// Community ID.
	CommunityID string `json:"communityId"`
	// Community Note presentation metadata returned by X.
	CommunityNote EmbeddedTweetRetweetedTweetRetweetedTweetCommunityNote `json:"communityNote"`
	// Content disclosure metadata shown by X when a tweet is labeled as paid
	// partnership content or AI-generated media.
	ContentDisclosure ContentDisclosure `json:"contentDisclosure"`
	// Public reply policy and conversation owner.
	ConversationControl EmbeddedTweetRetweetedTweetRetweetedTweetConversationControl `json:"conversationControl"`
	ConversationID      string                                                       `json:"conversationId"`
	CreatedAt           string                                                       `json:"createdAt"`
	DisplayTextRange    []int64                                                      `json:"displayTextRange"`
	// Lists edit-chain identifiers and the remaining edit window.
	Edit EmbeddedTweetRetweetedTweetRetweetedTweetEdit `json:"edit"`
	// Lists hashtags, symbols, links, and mentions from tweet text.
	Entities EmbeddedTweetRetweetedTweetRetweetedTweetEntities `json:"entities"`
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
	LimitedActions []EmbeddedTweetRetweetedTweetRetweetedTweetLimitedAction `json:"limitedActions"`
	// Attached media items, omitted when unavailable.
	Media []TweetMedia `json:"media"`
	// Complete Note Tweet content and rich-text metadata.
	NoteTweet EmbeddedTweetRetweetedTweetRetweetedTweetNoteTweet `json:"noteTweet"`
	// Describes public place metadata on a geotagged tweet.
	Place             EmbeddedTweetRetweetedTweetRetweetedTweetPlace `json:"place"`
	PossiblySensitive bool                                           `json:"possiblySensitive"`
	// Public metadata whose fields are defined by X.
	PostCta map[string]any `json:"postCta"`
	// Engagement counts retained from a prior tweet edit.
	PreviousCounts EmbeddedTweetRetweetedTweetRetweetedTweetPreviousCounts `json:"previousCounts"`
	// Final nested tweet context at depth 4.
	QuotedTweet EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweet `json:"quoted_tweet"`
	// Quoted tweet ID.
	QuotedTweetID string `json:"quotedTweetId"`
	// Public post and user referenced by this reaction.
	ReactionContext EmbeddedTweetRetweetedTweetRetweetedTweetReactionContext `json:"reactionContext"`
	// Final nested tweet context at depth 4.
	RetweetedTweet EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweet `json:"retweeted_tweet"`
	// Public metadata whose fields are defined by X.
	Scopes map[string]any `json:"scopes"`
	Source string         `json:"source"`
	// Public visibility notice attached to an available tweet.
	Tombstone EmbeddedTweetRetweetedTweetRetweetedTweetTombstone `json:"tombstone"`
	Type      string                                             `json:"type"`
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
func (r EmbeddedTweetRetweetedTweetRetweetedTweet) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetRetweetedTweetRetweetedTweet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Describes an X Article preview and its lifecycle metadata.
type EmbeddedTweetRetweetedTweetRetweetedTweetArticle struct {
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
func (r EmbeddedTweetRetweetedTweetRetweetedTweetArticle) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetRetweetedTweetRetweetedTweetArticle) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Describes a public card and its referenced profiles.
type EmbeddedTweetRetweetedTweetRetweetedTweetCard struct {
	ID string `json:"id"`
	// Public metadata whose fields are defined by X.
	BindingValues map[string]any `json:"bindingValues"`
	Name          string         `json:"name"`
	// Public metadata whose fields are defined by X.
	Platform map[string]any `json:"platform"`
	URL      string         `json:"url"`
	// Unresolved card user references.
	UserReferenceErrors []EmbeddedTweetRetweetedTweetRetweetedTweetCardUserReferenceError `json:"userReferenceErrors"`
	UserReferences      []UserProfile                                                     `json:"userReferences"`
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
func (r EmbeddedTweetRetweetedTweetRetweetedTweetCard) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetRetweetedTweetRetweetedTweetCard) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetRetweetedTweetCardUserReferenceError struct {
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
func (r EmbeddedTweetRetweetedTweetRetweetedTweetCardUserReferenceError) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetRetweetedTweetCardUserReferenceError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Community Note presentation metadata returned by X.
type EmbeddedTweetRetweetedTweetRetweetedTweetCommunityNote struct {
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
func (r EmbeddedTweetRetweetedTweetRetweetedTweetCommunityNote) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetRetweetedTweetRetweetedTweetCommunityNote) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Public reply policy and conversation owner.
type EmbeddedTweetRetweetedTweetRetweetedTweetConversationControl struct {
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
func (r EmbeddedTweetRetweetedTweetRetweetedTweetConversationControl) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetRetweetedTweetConversationControl) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Lists edit-chain identifiers and the remaining edit window.
type EmbeddedTweetRetweetedTweetRetweetedTweetEdit struct {
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
func (r EmbeddedTweetRetweetedTweetRetweetedTweetEdit) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetRetweetedTweetRetweetedTweetEdit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Lists hashtags, symbols, links, and mentions from tweet text.
type EmbeddedTweetRetweetedTweetRetweetedTweetEntities struct {
	Hashtags     []EmbeddedTweetRetweetedTweetRetweetedTweetEntitiesHashtag     `json:"hashtags"`
	Smarttags    []EmbeddedTweetRetweetedTweetRetweetedTweetEntitiesSmarttag    `json:"smarttags"`
	Symbols      []EmbeddedTweetRetweetedTweetRetweetedTweetEntitiesSymbol      `json:"symbols"`
	Timestamps   []EmbeddedTweetRetweetedTweetRetweetedTweetEntitiesTimestamp   `json:"timestamps"`
	URLs         []EmbeddedTweetRetweetedTweetRetweetedTweetEntitiesURL         `json:"urls"`
	UserMentions []EmbeddedTweetRetweetedTweetRetweetedTweetEntitiesUserMention `json:"user_mentions"`
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
func (r EmbeddedTweetRetweetedTweetRetweetedTweetEntities) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetRetweetedTweetRetweetedTweetEntities) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Provides hashtag text and source offsets within a tweet.
type EmbeddedTweetRetweetedTweetRetweetedTweetEntitiesHashtag struct {
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
func (r EmbeddedTweetRetweetedTweetRetweetedTweetEntitiesHashtag) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetRetweetedTweetRetweetedTweetEntitiesHashtag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indexed smart-tag, cashtag, or video timestamp metadata.
type EmbeddedTweetRetweetedTweetRetweetedTweetEntitiesSmarttag struct {
	Indices []int64                                                      `json:"indices"`
	Seconds float64                                                      `json:"seconds"`
	Tag     EmbeddedTweetRetweetedTweetRetweetedTweetEntitiesSmarttagTag `json:"tag"`
	Text    string                                                       `json:"text"`
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
func (r EmbeddedTweetRetweetedTweetRetweetedTweetEntitiesSmarttag) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetRetweetedTweetEntitiesSmarttag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetRetweetedTweetEntitiesSmarttagTag struct {
	Info EmbeddedTweetRetweetedTweetRetweetedTweetEntitiesSmarttagTagInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetRetweetedTweetRetweetedTweetEntitiesSmarttagTag) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetRetweetedTweetEntitiesSmarttagTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetRetweetedTweetEntitiesSmarttagTagInfo struct {
	Info EmbeddedTweetRetweetedTweetRetweetedTweetEntitiesSmarttagTagInfoInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetRetweetedTweetRetweetedTweetEntitiesSmarttagTagInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetRetweetedTweetEntitiesSmarttagTagInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetRetweetedTweetEntitiesSmarttagTagInfoInfo struct {
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
func (r EmbeddedTweetRetweetedTweetRetweetedTweetEntitiesSmarttagTagInfoInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetRetweetedTweetEntitiesSmarttagTagInfoInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indexed smart-tag, cashtag, or video timestamp metadata.
type EmbeddedTweetRetweetedTweetRetweetedTweetEntitiesSymbol struct {
	Indices []int64                                                    `json:"indices"`
	Seconds float64                                                    `json:"seconds"`
	Tag     EmbeddedTweetRetweetedTweetRetweetedTweetEntitiesSymbolTag `json:"tag"`
	Text    string                                                     `json:"text"`
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
func (r EmbeddedTweetRetweetedTweetRetweetedTweetEntitiesSymbol) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetRetweetedTweetRetweetedTweetEntitiesSymbol) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetRetweetedTweetEntitiesSymbolTag struct {
	Info EmbeddedTweetRetweetedTweetRetweetedTweetEntitiesSymbolTagInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetRetweetedTweetRetweetedTweetEntitiesSymbolTag) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetRetweetedTweetEntitiesSymbolTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetRetweetedTweetEntitiesSymbolTagInfo struct {
	Info EmbeddedTweetRetweetedTweetRetweetedTweetEntitiesSymbolTagInfoInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetRetweetedTweetRetweetedTweetEntitiesSymbolTagInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetRetweetedTweetEntitiesSymbolTagInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetRetweetedTweetEntitiesSymbolTagInfoInfo struct {
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
func (r EmbeddedTweetRetweetedTweetRetweetedTweetEntitiesSymbolTagInfoInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetRetweetedTweetEntitiesSymbolTagInfoInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indexed smart-tag, cashtag, or video timestamp metadata.
type EmbeddedTweetRetweetedTweetRetweetedTweetEntitiesTimestamp struct {
	Indices []int64                                                       `json:"indices"`
	Seconds float64                                                       `json:"seconds"`
	Tag     EmbeddedTweetRetweetedTweetRetweetedTweetEntitiesTimestampTag `json:"tag"`
	Text    string                                                        `json:"text"`
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
func (r EmbeddedTweetRetweetedTweetRetweetedTweetEntitiesTimestamp) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetRetweetedTweetEntitiesTimestamp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetRetweetedTweetEntitiesTimestampTag struct {
	Info EmbeddedTweetRetweetedTweetRetweetedTweetEntitiesTimestampTagInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetRetweetedTweetRetweetedTweetEntitiesTimestampTag) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetRetweetedTweetEntitiesTimestampTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetRetweetedTweetEntitiesTimestampTagInfo struct {
	Info EmbeddedTweetRetweetedTweetRetweetedTweetEntitiesTimestampTagInfoInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetRetweetedTweetRetweetedTweetEntitiesTimestampTagInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetRetweetedTweetEntitiesTimestampTagInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetRetweetedTweetEntitiesTimestampTagInfoInfo struct {
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
func (r EmbeddedTweetRetweetedTweetRetweetedTweetEntitiesTimestampTagInfoInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetRetweetedTweetEntitiesTimestampTagInfoInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Provides shortened, display, and expanded URLs from tweet text.
type EmbeddedTweetRetweetedTweetRetweetedTweetEntitiesURL struct {
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
func (r EmbeddedTweetRetweetedTweetRetweetedTweetEntitiesURL) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetRetweetedTweetRetweetedTweetEntitiesURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Provides profile identity and source offsets for a mention.
type EmbeddedTweetRetweetedTweetRetweetedTweetEntitiesUserMention struct {
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
func (r EmbeddedTweetRetweetedTweetRetweetedTweetEntitiesUserMention) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetRetweetedTweetEntitiesUserMention) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetRetweetedTweetLimitedAction struct {
	Action string                                                       `json:"action"`
	Prompt EmbeddedTweetRetweetedTweetRetweetedTweetLimitedActionPrompt `json:"prompt"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Action      respjson.Field
		Prompt      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetRetweetedTweetRetweetedTweetLimitedAction) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetRetweetedTweetRetweetedTweetLimitedAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetRetweetedTweetLimitedActionPrompt struct {
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
func (r EmbeddedTweetRetweetedTweetRetweetedTweetLimitedActionPrompt) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetRetweetedTweetLimitedActionPrompt) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Complete Note Tweet content and rich-text metadata.
type EmbeddedTweetRetweetedTweetRetweetedTweetNoteTweet struct {
	Text string `json:"text" api:"required"`
	ID   string `json:"id"`
	// Lists hashtags, symbols, links, and mentions from tweet text.
	Entities EmbeddedTweetRetweetedTweetRetweetedTweetNoteTweetEntities `json:"entities"`
	// Inline media positions in the Note Tweet text.
	InlineMedia  []EmbeddedTweetRetweetedTweetRetweetedTweetNoteTweetInlineMedia `json:"inlineMedia"`
	IsExpandable bool                                                            `json:"isExpandable"`
	RichtextTags []EmbeddedTweetRetweetedTweetRetweetedTweetNoteTweetRichtextTag `json:"richtextTags"`
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
func (r EmbeddedTweetRetweetedTweetRetweetedTweetNoteTweet) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetRetweetedTweetRetweetedTweetNoteTweet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Lists hashtags, symbols, links, and mentions from tweet text.
type EmbeddedTweetRetweetedTweetRetweetedTweetNoteTweetEntities struct {
	Hashtags     []EmbeddedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesHashtag     `json:"hashtags"`
	Smarttags    []EmbeddedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesSmarttag    `json:"smarttags"`
	Symbols      []EmbeddedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesSymbol      `json:"symbols"`
	Timestamps   []EmbeddedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesTimestamp   `json:"timestamps"`
	URLs         []EmbeddedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesURL         `json:"urls"`
	UserMentions []EmbeddedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesUserMention `json:"user_mentions"`
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
func (r EmbeddedTweetRetweetedTweetRetweetedTweetNoteTweetEntities) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetRetweetedTweetNoteTweetEntities) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Provides hashtag text and source offsets within a tweet.
type EmbeddedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesHashtag struct {
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
func (r EmbeddedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesHashtag) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesHashtag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indexed smart-tag, cashtag, or video timestamp metadata.
type EmbeddedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesSmarttag struct {
	Indices []int64                                                               `json:"indices"`
	Seconds float64                                                               `json:"seconds"`
	Tag     EmbeddedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesSmarttagTag `json:"tag"`
	Text    string                                                                `json:"text"`
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
func (r EmbeddedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesSmarttag) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesSmarttag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesSmarttagTag struct {
	Info EmbeddedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesSmarttagTagInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesSmarttagTag) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesSmarttagTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesSmarttagTagInfo struct {
	Info EmbeddedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesSmarttagTagInfoInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesSmarttagTagInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesSmarttagTagInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesSmarttagTagInfoInfo struct {
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
func (r EmbeddedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesSmarttagTagInfoInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesSmarttagTagInfoInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indexed smart-tag, cashtag, or video timestamp metadata.
type EmbeddedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesSymbol struct {
	Indices []int64                                                             `json:"indices"`
	Seconds float64                                                             `json:"seconds"`
	Tag     EmbeddedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesSymbolTag `json:"tag"`
	Text    string                                                              `json:"text"`
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
func (r EmbeddedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesSymbol) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesSymbol) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesSymbolTag struct {
	Info EmbeddedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesSymbolTagInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesSymbolTag) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesSymbolTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesSymbolTagInfo struct {
	Info EmbeddedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesSymbolTagInfoInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesSymbolTagInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesSymbolTagInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesSymbolTagInfoInfo struct {
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
func (r EmbeddedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesSymbolTagInfoInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesSymbolTagInfoInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indexed smart-tag, cashtag, or video timestamp metadata.
type EmbeddedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesTimestamp struct {
	Indices []int64                                                                `json:"indices"`
	Seconds float64                                                                `json:"seconds"`
	Tag     EmbeddedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesTimestampTag `json:"tag"`
	Text    string                                                                 `json:"text"`
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
func (r EmbeddedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesTimestamp) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesTimestamp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesTimestampTag struct {
	Info EmbeddedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesTimestampTagInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesTimestampTag) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesTimestampTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesTimestampTagInfo struct {
	Info EmbeddedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesTimestampTagInfoInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesTimestampTagInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesTimestampTagInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesTimestampTagInfoInfo struct {
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
func (r EmbeddedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesTimestampTagInfoInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesTimestampTagInfoInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Provides shortened, display, and expanded URLs from tweet text.
type EmbeddedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesURL struct {
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
func (r EmbeddedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesURL) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Provides profile identity and source offsets for a mention.
type EmbeddedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesUserMention struct {
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
func (r EmbeddedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesUserMention) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesUserMention) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetRetweetedTweetNoteTweetInlineMedia struct {
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
func (r EmbeddedTweetRetweetedTweetRetweetedTweetNoteTweetInlineMedia) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetRetweetedTweetNoteTweetInlineMedia) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetRetweetedTweetNoteTweetRichtextTag struct {
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
func (r EmbeddedTweetRetweetedTweetRetweetedTweetNoteTweetRichtextTag) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetRetweetedTweetNoteTweetRichtextTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Describes public place metadata on a geotagged tweet.
type EmbeddedTweetRetweetedTweetRetweetedTweetPlace struct {
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
func (r EmbeddedTweetRetweetedTweetRetweetedTweetPlace) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetRetweetedTweetRetweetedTweetPlace) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Engagement counts retained from a prior tweet edit.
type EmbeddedTweetRetweetedTweetRetweetedTweetPreviousCounts struct {
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
func (r EmbeddedTweetRetweetedTweetRetweetedTweetPreviousCounts) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetRetweetedTweetRetweetedTweetPreviousCounts) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Final nested tweet context at depth 4.
type EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweet struct {
	ID            string `json:"id" api:"required"`
	BookmarkCount int64  `json:"bookmarkCount" api:"required"`
	LikeCount     int64  `json:"likeCount" api:"required"`
	QuoteCount    int64  `json:"quoteCount" api:"required"`
	ReplyCount    int64  `json:"replyCount" api:"required"`
	RetweetCount  int64  `json:"retweetCount" api:"required"`
	Text          string `json:"text" api:"required"`
	ViewCount     int64  `json:"viewCount" api:"required"`
	// Describes an X Article preview and its lifecycle metadata.
	Article EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetArticle `json:"article"`
	// X user profile with bio, follower counts, and verification status.
	Author UserProfile `json:"author"`
	// Describes a public card and its referenced profiles.
	Card EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetCard `json:"card"`
	// Community ID.
	CommunityID string `json:"communityId"`
	// Community Note presentation metadata returned by X.
	CommunityNote EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetCommunityNote `json:"communityNote"`
	// Content disclosure metadata shown by X when a tweet is labeled as paid
	// partnership content or AI-generated media.
	ContentDisclosure ContentDisclosure `json:"contentDisclosure"`
	// Public reply policy and conversation owner.
	ConversationControl EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetConversationControl `json:"conversationControl"`
	ConversationID      string                                                                  `json:"conversationId"`
	CreatedAt           string                                                                  `json:"createdAt"`
	DisplayTextRange    []int64                                                                 `json:"displayTextRange"`
	// Lists edit-chain identifiers and the remaining edit window.
	Edit EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetEdit `json:"edit"`
	// Lists hashtags, symbols, links, and mentions from tweet text.
	Entities EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetEntities `json:"entities"`
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
	LimitedActions []EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetLimitedAction `json:"limitedActions"`
	// Attached media items, omitted when unavailable.
	Media []TweetMedia `json:"media"`
	// Complete Note Tweet content and rich-text metadata.
	NoteTweet EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetNoteTweet `json:"noteTweet"`
	// Describes public place metadata on a geotagged tweet.
	Place             EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetPlace `json:"place"`
	PossiblySensitive bool                                                      `json:"possiblySensitive"`
	// Public metadata whose fields are defined by X.
	PostCta map[string]any `json:"postCta"`
	// Engagement counts retained from a prior tweet edit.
	PreviousCounts EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetPreviousCounts `json:"previousCounts"`
	// Quoted tweet ID.
	QuotedTweetID string `json:"quotedTweetId"`
	// Public post and user referenced by this reaction.
	ReactionContext EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetReactionContext `json:"reactionContext"`
	// Public metadata whose fields are defined by X.
	Scopes map[string]any `json:"scopes"`
	Source string         `json:"source"`
	// Public visibility notice attached to an available tweet.
	Tombstone EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetTombstone `json:"tombstone"`
	Type      string                                                        `json:"type"`
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
		QuotedTweetID       respjson.Field
		ReactionContext     respjson.Field
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
func (r EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweet) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Describes an X Article preview and its lifecycle metadata.
type EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetArticle struct {
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
func (r EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetArticle) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetArticle) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Describes a public card and its referenced profiles.
type EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetCard struct {
	ID string `json:"id"`
	// Public metadata whose fields are defined by X.
	BindingValues map[string]any `json:"bindingValues"`
	Name          string         `json:"name"`
	// Public metadata whose fields are defined by X.
	Platform map[string]any `json:"platform"`
	URL      string         `json:"url"`
	// Unresolved card user references.
	UserReferenceErrors []EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetCardUserReferenceError `json:"userReferenceErrors"`
	UserReferences      []UserProfile                                                                `json:"userReferences"`
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
func (r EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetCard) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetCard) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetCardUserReferenceError struct {
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
func (r EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetCardUserReferenceError) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetCardUserReferenceError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Community Note presentation metadata returned by X.
type EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetCommunityNote struct {
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
func (r EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetCommunityNote) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetCommunityNote) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Public reply policy and conversation owner.
type EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetConversationControl struct {
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
func (r EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetConversationControl) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetConversationControl) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Lists edit-chain identifiers and the remaining edit window.
type EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetEdit struct {
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
func (r EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetEdit) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetEdit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Lists hashtags, symbols, links, and mentions from tweet text.
type EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetEntities struct {
	Hashtags     []EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetEntitiesHashtag     `json:"hashtags"`
	Smarttags    []EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetEntitiesSmarttag    `json:"smarttags"`
	Symbols      []EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetEntitiesSymbol      `json:"symbols"`
	Timestamps   []EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetEntitiesTimestamp   `json:"timestamps"`
	URLs         []EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetEntitiesURL         `json:"urls"`
	UserMentions []EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetEntitiesUserMention `json:"user_mentions"`
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
func (r EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetEntities) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetEntities) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Provides hashtag text and source offsets within a tweet.
type EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetEntitiesHashtag struct {
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
func (r EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetEntitiesHashtag) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetEntitiesHashtag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indexed smart-tag, cashtag, or video timestamp metadata.
type EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetEntitiesSmarttag struct {
	Indices []int64                                                                 `json:"indices"`
	Seconds float64                                                                 `json:"seconds"`
	Tag     EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetEntitiesSmarttagTag `json:"tag"`
	Text    string                                                                  `json:"text"`
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
func (r EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetEntitiesSmarttag) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetEntitiesSmarttag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetEntitiesSmarttagTag struct {
	Info EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetEntitiesSmarttagTagInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetEntitiesSmarttagTag) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetEntitiesSmarttagTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetEntitiesSmarttagTagInfo struct {
	Info EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetEntitiesSmarttagTagInfoInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetEntitiesSmarttagTagInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetEntitiesSmarttagTagInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetEntitiesSmarttagTagInfoInfo struct {
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
func (r EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetEntitiesSmarttagTagInfoInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetEntitiesSmarttagTagInfoInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indexed smart-tag, cashtag, or video timestamp metadata.
type EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetEntitiesSymbol struct {
	Indices []int64                                                               `json:"indices"`
	Seconds float64                                                               `json:"seconds"`
	Tag     EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetEntitiesSymbolTag `json:"tag"`
	Text    string                                                                `json:"text"`
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
func (r EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetEntitiesSymbol) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetEntitiesSymbol) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetEntitiesSymbolTag struct {
	Info EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetEntitiesSymbolTagInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetEntitiesSymbolTag) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetEntitiesSymbolTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetEntitiesSymbolTagInfo struct {
	Info EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetEntitiesSymbolTagInfoInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetEntitiesSymbolTagInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetEntitiesSymbolTagInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetEntitiesSymbolTagInfoInfo struct {
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
func (r EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetEntitiesSymbolTagInfoInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetEntitiesSymbolTagInfoInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indexed smart-tag, cashtag, or video timestamp metadata.
type EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetEntitiesTimestamp struct {
	Indices []int64                                                                  `json:"indices"`
	Seconds float64                                                                  `json:"seconds"`
	Tag     EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetEntitiesTimestampTag `json:"tag"`
	Text    string                                                                   `json:"text"`
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
func (r EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetEntitiesTimestamp) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetEntitiesTimestamp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetEntitiesTimestampTag struct {
	Info EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetEntitiesTimestampTagInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetEntitiesTimestampTag) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetEntitiesTimestampTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetEntitiesTimestampTagInfo struct {
	Info EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetEntitiesTimestampTagInfoInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetEntitiesTimestampTagInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetEntitiesTimestampTagInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetEntitiesTimestampTagInfoInfo struct {
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
func (r EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetEntitiesTimestampTagInfoInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetEntitiesTimestampTagInfoInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Provides shortened, display, and expanded URLs from tweet text.
type EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetEntitiesURL struct {
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
func (r EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetEntitiesURL) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetEntitiesURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Provides profile identity and source offsets for a mention.
type EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetEntitiesUserMention struct {
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
func (r EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetEntitiesUserMention) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetEntitiesUserMention) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetLimitedAction struct {
	Action string                                                                  `json:"action"`
	Prompt EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetLimitedActionPrompt `json:"prompt"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Action      respjson.Field
		Prompt      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetLimitedAction) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetLimitedAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetLimitedActionPrompt struct {
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
func (r EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetLimitedActionPrompt) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetLimitedActionPrompt) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Complete Note Tweet content and rich-text metadata.
type EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetNoteTweet struct {
	Text string `json:"text" api:"required"`
	ID   string `json:"id"`
	// Lists hashtags, symbols, links, and mentions from tweet text.
	Entities EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetNoteTweetEntities `json:"entities"`
	// Inline media positions in the Note Tweet text.
	InlineMedia  []EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetNoteTweetInlineMedia `json:"inlineMedia"`
	IsExpandable bool                                                                       `json:"isExpandable"`
	RichtextTags []EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetNoteTweetRichtextTag `json:"richtextTags"`
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
func (r EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetNoteTweet) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetNoteTweet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Lists hashtags, symbols, links, and mentions from tweet text.
type EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetNoteTweetEntities struct {
	Hashtags     []EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesHashtag     `json:"hashtags"`
	Smarttags    []EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesSmarttag    `json:"smarttags"`
	Symbols      []EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesSymbol      `json:"symbols"`
	Timestamps   []EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesTimestamp   `json:"timestamps"`
	URLs         []EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesURL         `json:"urls"`
	UserMentions []EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesUserMention `json:"user_mentions"`
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
func (r EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetNoteTweetEntities) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetNoteTweetEntities) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Provides hashtag text and source offsets within a tweet.
type EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesHashtag struct {
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
func (r EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesHashtag) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesHashtag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indexed smart-tag, cashtag, or video timestamp metadata.
type EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesSmarttag struct {
	Indices []int64                                                                          `json:"indices"`
	Seconds float64                                                                          `json:"seconds"`
	Tag     EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesSmarttagTag `json:"tag"`
	Text    string                                                                           `json:"text"`
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
func (r EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesSmarttag) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesSmarttag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesSmarttagTag struct {
	Info EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesSmarttagTagInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesSmarttagTag) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesSmarttagTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesSmarttagTagInfo struct {
	Info EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesSmarttagTagInfoInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesSmarttagTagInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesSmarttagTagInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesSmarttagTagInfoInfo struct {
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
func (r EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesSmarttagTagInfoInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesSmarttagTagInfoInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indexed smart-tag, cashtag, or video timestamp metadata.
type EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesSymbol struct {
	Indices []int64                                                                        `json:"indices"`
	Seconds float64                                                                        `json:"seconds"`
	Tag     EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesSymbolTag `json:"tag"`
	Text    string                                                                         `json:"text"`
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
func (r EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesSymbol) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesSymbol) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesSymbolTag struct {
	Info EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesSymbolTagInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesSymbolTag) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesSymbolTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesSymbolTagInfo struct {
	Info EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesSymbolTagInfoInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesSymbolTagInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesSymbolTagInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesSymbolTagInfoInfo struct {
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
func (r EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesSymbolTagInfoInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesSymbolTagInfoInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indexed smart-tag, cashtag, or video timestamp metadata.
type EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesTimestamp struct {
	Indices []int64                                                                           `json:"indices"`
	Seconds float64                                                                           `json:"seconds"`
	Tag     EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesTimestampTag `json:"tag"`
	Text    string                                                                            `json:"text"`
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
func (r EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesTimestamp) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesTimestamp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesTimestampTag struct {
	Info EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesTimestampTagInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesTimestampTag) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesTimestampTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesTimestampTagInfo struct {
	Info EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesTimestampTagInfoInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesTimestampTagInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesTimestampTagInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesTimestampTagInfoInfo struct {
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
func (r EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesTimestampTagInfoInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesTimestampTagInfoInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Provides shortened, display, and expanded URLs from tweet text.
type EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesURL struct {
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
func (r EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesURL) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Provides profile identity and source offsets for a mention.
type EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesUserMention struct {
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
func (r EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesUserMention) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesUserMention) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetNoteTweetInlineMedia struct {
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
func (r EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetNoteTweetInlineMedia) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetNoteTweetInlineMedia) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetNoteTweetRichtextTag struct {
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
func (r EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetNoteTweetRichtextTag) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetNoteTweetRichtextTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Describes public place metadata on a geotagged tweet.
type EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetPlace struct {
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
func (r EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetPlace) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetPlace) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Engagement counts retained from a prior tweet edit.
type EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetPreviousCounts struct {
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
func (r EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetPreviousCounts) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetPreviousCounts) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Public post and user referenced by this reaction.
type EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetReactionContext struct {
	// Referenced post ID.
	ReactedToPostID string `json:"reactedToPostId"`
	// X user profile with bio, follower counts, and verification status.
	ReactedToUser UserProfile `json:"reactedToUser"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ReactedToPostID respjson.Field
		ReactedToUser   respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetReactionContext) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetReactionContext) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Public visibility notice attached to an available tweet.
type EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetTombstone struct {
	Text EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetTombstoneText `json:"text"`
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
func (r EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetTombstone) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetTombstone) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetTombstoneText struct {
	Entities []EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetTombstoneTextEntity `json:"entities"`
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
func (r EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetTombstoneText) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetTombstoneText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetTombstoneTextEntity struct {
	FromIndex int64                                                                      `json:"fromIndex"`
	Ref       EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetTombstoneTextEntityRef `json:"ref"`
	ToIndex   int64                                                                      `json:"toIndex"`
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
func (r EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetTombstoneTextEntity) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetTombstoneTextEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetTombstoneTextEntityRef struct {
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
func (r EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetTombstoneTextEntityRef) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetTombstoneTextEntityRef) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Public post and user referenced by this reaction.
type EmbeddedTweetRetweetedTweetRetweetedTweetReactionContext struct {
	// Referenced post ID.
	ReactedToPostID string `json:"reactedToPostId"`
	// X user profile with bio, follower counts, and verification status.
	ReactedToUser UserProfile `json:"reactedToUser"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ReactedToPostID respjson.Field
		ReactedToUser   respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetRetweetedTweetRetweetedTweetReactionContext) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetRetweetedTweetRetweetedTweetReactionContext) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Final nested tweet context at depth 4.
type EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweet struct {
	ID            string `json:"id" api:"required"`
	BookmarkCount int64  `json:"bookmarkCount" api:"required"`
	LikeCount     int64  `json:"likeCount" api:"required"`
	QuoteCount    int64  `json:"quoteCount" api:"required"`
	ReplyCount    int64  `json:"replyCount" api:"required"`
	RetweetCount  int64  `json:"retweetCount" api:"required"`
	Text          string `json:"text" api:"required"`
	ViewCount     int64  `json:"viewCount" api:"required"`
	// Describes an X Article preview and its lifecycle metadata.
	Article EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetArticle `json:"article"`
	// X user profile with bio, follower counts, and verification status.
	Author UserProfile `json:"author"`
	// Describes a public card and its referenced profiles.
	Card EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetCard `json:"card"`
	// Community ID.
	CommunityID string `json:"communityId"`
	// Community Note presentation metadata returned by X.
	CommunityNote EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetCommunityNote `json:"communityNote"`
	// Content disclosure metadata shown by X when a tweet is labeled as paid
	// partnership content or AI-generated media.
	ContentDisclosure ContentDisclosure `json:"contentDisclosure"`
	// Public reply policy and conversation owner.
	ConversationControl EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetConversationControl `json:"conversationControl"`
	ConversationID      string                                                                     `json:"conversationId"`
	CreatedAt           string                                                                     `json:"createdAt"`
	DisplayTextRange    []int64                                                                    `json:"displayTextRange"`
	// Lists edit-chain identifiers and the remaining edit window.
	Edit EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetEdit `json:"edit"`
	// Lists hashtags, symbols, links, and mentions from tweet text.
	Entities EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetEntities `json:"entities"`
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
	LimitedActions []EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetLimitedAction `json:"limitedActions"`
	// Attached media items, omitted when unavailable.
	Media []TweetMedia `json:"media"`
	// Complete Note Tweet content and rich-text metadata.
	NoteTweet EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetNoteTweet `json:"noteTweet"`
	// Describes public place metadata on a geotagged tweet.
	Place             EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetPlace `json:"place"`
	PossiblySensitive bool                                                         `json:"possiblySensitive"`
	// Public metadata whose fields are defined by X.
	PostCta map[string]any `json:"postCta"`
	// Engagement counts retained from a prior tweet edit.
	PreviousCounts EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetPreviousCounts `json:"previousCounts"`
	// Quoted tweet ID.
	QuotedTweetID string `json:"quotedTweetId"`
	// Public post and user referenced by this reaction.
	ReactionContext EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetReactionContext `json:"reactionContext"`
	// Public metadata whose fields are defined by X.
	Scopes map[string]any `json:"scopes"`
	Source string         `json:"source"`
	// Public visibility notice attached to an available tweet.
	Tombstone EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetTombstone `json:"tombstone"`
	Type      string                                                           `json:"type"`
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
		QuotedTweetID       respjson.Field
		ReactionContext     respjson.Field
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
func (r EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweet) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Describes an X Article preview and its lifecycle metadata.
type EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetArticle struct {
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
func (r EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetArticle) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetArticle) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Describes a public card and its referenced profiles.
type EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetCard struct {
	ID string `json:"id"`
	// Public metadata whose fields are defined by X.
	BindingValues map[string]any `json:"bindingValues"`
	Name          string         `json:"name"`
	// Public metadata whose fields are defined by X.
	Platform map[string]any `json:"platform"`
	URL      string         `json:"url"`
	// Unresolved card user references.
	UserReferenceErrors []EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetCardUserReferenceError `json:"userReferenceErrors"`
	UserReferences      []UserProfile                                                                   `json:"userReferences"`
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
func (r EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetCard) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetCard) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetCardUserReferenceError struct {
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
func (r EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetCardUserReferenceError) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetCardUserReferenceError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Community Note presentation metadata returned by X.
type EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetCommunityNote struct {
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
func (r EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetCommunityNote) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetCommunityNote) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Public reply policy and conversation owner.
type EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetConversationControl struct {
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
func (r EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetConversationControl) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetConversationControl) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Lists edit-chain identifiers and the remaining edit window.
type EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetEdit struct {
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
func (r EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetEdit) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetEdit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Lists hashtags, symbols, links, and mentions from tweet text.
type EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetEntities struct {
	Hashtags     []EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetEntitiesHashtag     `json:"hashtags"`
	Smarttags    []EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetEntitiesSmarttag    `json:"smarttags"`
	Symbols      []EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetEntitiesSymbol      `json:"symbols"`
	Timestamps   []EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetEntitiesTimestamp   `json:"timestamps"`
	URLs         []EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetEntitiesURL         `json:"urls"`
	UserMentions []EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetEntitiesUserMention `json:"user_mentions"`
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
func (r EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetEntities) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetEntities) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Provides hashtag text and source offsets within a tweet.
type EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetEntitiesHashtag struct {
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
func (r EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetEntitiesHashtag) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetEntitiesHashtag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indexed smart-tag, cashtag, or video timestamp metadata.
type EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetEntitiesSmarttag struct {
	Indices []int64                                                                    `json:"indices"`
	Seconds float64                                                                    `json:"seconds"`
	Tag     EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetEntitiesSmarttagTag `json:"tag"`
	Text    string                                                                     `json:"text"`
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
func (r EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetEntitiesSmarttag) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetEntitiesSmarttag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetEntitiesSmarttagTag struct {
	Info EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetEntitiesSmarttagTagInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetEntitiesSmarttagTag) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetEntitiesSmarttagTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetEntitiesSmarttagTagInfo struct {
	Info EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetEntitiesSmarttagTagInfoInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetEntitiesSmarttagTagInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetEntitiesSmarttagTagInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetEntitiesSmarttagTagInfoInfo struct {
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
func (r EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetEntitiesSmarttagTagInfoInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetEntitiesSmarttagTagInfoInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indexed smart-tag, cashtag, or video timestamp metadata.
type EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetEntitiesSymbol struct {
	Indices []int64                                                                  `json:"indices"`
	Seconds float64                                                                  `json:"seconds"`
	Tag     EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetEntitiesSymbolTag `json:"tag"`
	Text    string                                                                   `json:"text"`
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
func (r EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetEntitiesSymbol) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetEntitiesSymbol) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetEntitiesSymbolTag struct {
	Info EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetEntitiesSymbolTagInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetEntitiesSymbolTag) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetEntitiesSymbolTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetEntitiesSymbolTagInfo struct {
	Info EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetEntitiesSymbolTagInfoInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetEntitiesSymbolTagInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetEntitiesSymbolTagInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetEntitiesSymbolTagInfoInfo struct {
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
func (r EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetEntitiesSymbolTagInfoInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetEntitiesSymbolTagInfoInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indexed smart-tag, cashtag, or video timestamp metadata.
type EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetEntitiesTimestamp struct {
	Indices []int64                                                                     `json:"indices"`
	Seconds float64                                                                     `json:"seconds"`
	Tag     EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetEntitiesTimestampTag `json:"tag"`
	Text    string                                                                      `json:"text"`
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
func (r EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetEntitiesTimestamp) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetEntitiesTimestamp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetEntitiesTimestampTag struct {
	Info EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetEntitiesTimestampTagInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetEntitiesTimestampTag) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetEntitiesTimestampTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetEntitiesTimestampTagInfo struct {
	Info EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetEntitiesTimestampTagInfoInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetEntitiesTimestampTagInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetEntitiesTimestampTagInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetEntitiesTimestampTagInfoInfo struct {
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
func (r EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetEntitiesTimestampTagInfoInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetEntitiesTimestampTagInfoInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Provides shortened, display, and expanded URLs from tweet text.
type EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetEntitiesURL struct {
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
func (r EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetEntitiesURL) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetEntitiesURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Provides profile identity and source offsets for a mention.
type EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetEntitiesUserMention struct {
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
func (r EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetEntitiesUserMention) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetEntitiesUserMention) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetLimitedAction struct {
	Action string                                                                     `json:"action"`
	Prompt EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetLimitedActionPrompt `json:"prompt"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Action      respjson.Field
		Prompt      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetLimitedAction) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetLimitedAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetLimitedActionPrompt struct {
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
func (r EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetLimitedActionPrompt) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetLimitedActionPrompt) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Complete Note Tweet content and rich-text metadata.
type EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetNoteTweet struct {
	Text string `json:"text" api:"required"`
	ID   string `json:"id"`
	// Lists hashtags, symbols, links, and mentions from tweet text.
	Entities EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetNoteTweetEntities `json:"entities"`
	// Inline media positions in the Note Tweet text.
	InlineMedia  []EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetNoteTweetInlineMedia `json:"inlineMedia"`
	IsExpandable bool                                                                          `json:"isExpandable"`
	RichtextTags []EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetNoteTweetRichtextTag `json:"richtextTags"`
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
func (r EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetNoteTweet) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetNoteTweet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Lists hashtags, symbols, links, and mentions from tweet text.
type EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetNoteTweetEntities struct {
	Hashtags     []EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesHashtag     `json:"hashtags"`
	Smarttags    []EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesSmarttag    `json:"smarttags"`
	Symbols      []EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesSymbol      `json:"symbols"`
	Timestamps   []EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesTimestamp   `json:"timestamps"`
	URLs         []EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesURL         `json:"urls"`
	UserMentions []EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesUserMention `json:"user_mentions"`
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
func (r EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetNoteTweetEntities) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetNoteTweetEntities) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Provides hashtag text and source offsets within a tweet.
type EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesHashtag struct {
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
func (r EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesHashtag) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesHashtag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indexed smart-tag, cashtag, or video timestamp metadata.
type EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesSmarttag struct {
	Indices []int64                                                                             `json:"indices"`
	Seconds float64                                                                             `json:"seconds"`
	Tag     EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesSmarttagTag `json:"tag"`
	Text    string                                                                              `json:"text"`
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
func (r EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesSmarttag) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesSmarttag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesSmarttagTag struct {
	Info EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesSmarttagTagInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesSmarttagTag) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesSmarttagTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesSmarttagTagInfo struct {
	Info EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesSmarttagTagInfoInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesSmarttagTagInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesSmarttagTagInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesSmarttagTagInfoInfo struct {
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
func (r EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesSmarttagTagInfoInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesSmarttagTagInfoInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indexed smart-tag, cashtag, or video timestamp metadata.
type EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesSymbol struct {
	Indices []int64                                                                           `json:"indices"`
	Seconds float64                                                                           `json:"seconds"`
	Tag     EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesSymbolTag `json:"tag"`
	Text    string                                                                            `json:"text"`
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
func (r EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesSymbol) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesSymbol) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesSymbolTag struct {
	Info EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesSymbolTagInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesSymbolTag) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesSymbolTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesSymbolTagInfo struct {
	Info EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesSymbolTagInfoInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesSymbolTagInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesSymbolTagInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesSymbolTagInfoInfo struct {
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
func (r EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesSymbolTagInfoInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesSymbolTagInfoInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indexed smart-tag, cashtag, or video timestamp metadata.
type EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesTimestamp struct {
	Indices []int64                                                                              `json:"indices"`
	Seconds float64                                                                              `json:"seconds"`
	Tag     EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesTimestampTag `json:"tag"`
	Text    string                                                                               `json:"text"`
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
func (r EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesTimestamp) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesTimestamp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesTimestampTag struct {
	Info EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesTimestampTagInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesTimestampTag) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesTimestampTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesTimestampTagInfo struct {
	Info EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesTimestampTagInfoInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesTimestampTagInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesTimestampTagInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesTimestampTagInfoInfo struct {
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
func (r EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesTimestampTagInfoInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesTimestampTagInfoInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Provides shortened, display, and expanded URLs from tweet text.
type EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesURL struct {
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
func (r EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesURL) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Provides profile identity and source offsets for a mention.
type EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesUserMention struct {
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
func (r EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesUserMention) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesUserMention) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetNoteTweetInlineMedia struct {
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
func (r EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetNoteTweetInlineMedia) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetNoteTweetInlineMedia) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetNoteTweetRichtextTag struct {
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
func (r EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetNoteTweetRichtextTag) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetNoteTweetRichtextTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Describes public place metadata on a geotagged tweet.
type EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetPlace struct {
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
func (r EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetPlace) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetPlace) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Engagement counts retained from a prior tweet edit.
type EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetPreviousCounts struct {
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
func (r EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetPreviousCounts) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetPreviousCounts) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Public post and user referenced by this reaction.
type EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetReactionContext struct {
	// Referenced post ID.
	ReactedToPostID string `json:"reactedToPostId"`
	// X user profile with bio, follower counts, and verification status.
	ReactedToUser UserProfile `json:"reactedToUser"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ReactedToPostID respjson.Field
		ReactedToUser   respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetReactionContext) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetReactionContext) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Public visibility notice attached to an available tweet.
type EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetTombstone struct {
	Text EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetTombstoneText `json:"text"`
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
func (r EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetTombstone) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetTombstone) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetTombstoneText struct {
	Entities []EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetTombstoneTextEntity `json:"entities"`
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
func (r EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetTombstoneText) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetTombstoneText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetTombstoneTextEntity struct {
	FromIndex int64                                                                         `json:"fromIndex"`
	Ref       EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetTombstoneTextEntityRef `json:"ref"`
	ToIndex   int64                                                                         `json:"toIndex"`
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
func (r EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetTombstoneTextEntity) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetTombstoneTextEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetTombstoneTextEntityRef struct {
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
func (r EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetTombstoneTextEntityRef) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetTombstoneTextEntityRef) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Public visibility notice attached to an available tweet.
type EmbeddedTweetRetweetedTweetRetweetedTweetTombstone struct {
	Text EmbeddedTweetRetweetedTweetRetweetedTweetTombstoneText `json:"text"`
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
func (r EmbeddedTweetRetweetedTweetRetweetedTweetTombstone) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetRetweetedTweetRetweetedTweetTombstone) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetRetweetedTweetTombstoneText struct {
	Entities []EmbeddedTweetRetweetedTweetRetweetedTweetTombstoneTextEntity `json:"entities"`
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
func (r EmbeddedTweetRetweetedTweetRetweetedTweetTombstoneText) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetRetweetedTweetRetweetedTweetTombstoneText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetRetweetedTweetTombstoneTextEntity struct {
	FromIndex int64                                                           `json:"fromIndex"`
	Ref       EmbeddedTweetRetweetedTweetRetweetedTweetTombstoneTextEntityRef `json:"ref"`
	ToIndex   int64                                                           `json:"toIndex"`
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
func (r EmbeddedTweetRetweetedTweetRetweetedTweetTombstoneTextEntity) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetRetweetedTweetTombstoneTextEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetRetweetedTweetTombstoneTextEntityRef struct {
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
func (r EmbeddedTweetRetweetedTweetRetweetedTweetTombstoneTextEntityRef) RawJSON() string {
	return r.JSON.raw
}
func (r *EmbeddedTweetRetweetedTweetRetweetedTweetTombstoneTextEntityRef) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Public visibility notice attached to an available tweet.
type EmbeddedTweetRetweetedTweetTombstone struct {
	Text EmbeddedTweetRetweetedTweetTombstoneText `json:"text"`
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
func (r EmbeddedTweetRetweetedTweetTombstone) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetRetweetedTweetTombstone) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetTombstoneText struct {
	Entities []EmbeddedTweetRetweetedTweetTombstoneTextEntity `json:"entities"`
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
func (r EmbeddedTweetRetweetedTweetTombstoneText) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetRetweetedTweetTombstoneText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetTombstoneTextEntity struct {
	FromIndex int64                                             `json:"fromIndex"`
	Ref       EmbeddedTweetRetweetedTweetTombstoneTextEntityRef `json:"ref"`
	ToIndex   int64                                             `json:"toIndex"`
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
func (r EmbeddedTweetRetweetedTweetTombstoneTextEntity) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetRetweetedTweetTombstoneTextEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetRetweetedTweetTombstoneTextEntityRef struct {
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
func (r EmbeddedTweetRetweetedTweetTombstoneTextEntityRef) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetRetweetedTweetTombstoneTextEntityRef) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Public visibility notice attached to an available tweet.
type EmbeddedTweetTombstone struct {
	Text EmbeddedTweetTombstoneText `json:"text"`
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
func (r EmbeddedTweetTombstone) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetTombstone) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetTombstoneText struct {
	Entities []EmbeddedTweetTombstoneTextEntity `json:"entities"`
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
func (r EmbeddedTweetTombstoneText) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetTombstoneText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetTombstoneTextEntity struct {
	FromIndex int64                               `json:"fromIndex"`
	Ref       EmbeddedTweetTombstoneTextEntityRef `json:"ref"`
	ToIndex   int64                               `json:"toIndex"`
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
func (r EmbeddedTweetTombstoneTextEntity) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetTombstoneTextEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddedTweetTombstoneTextEntityRef struct {
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
func (r EmbeddedTweetTombstoneTextEntityRef) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetTombstoneTextEntityRef) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Type of monitor event fired when account activity occurs.
type EventType string

const (
	EventTypeTweetNew                  EventType = "tweet.new"
	EventTypeTweetReply                EventType = "tweet.reply"
	EventTypeTweetRetweet              EventType = "tweet.retweet"
	EventTypeTweetQuote                EventType = "tweet.quote"
	EventTypeTweetMedia                EventType = "tweet.media"
	EventTypeTweetLink                 EventType = "tweet.link"
	EventTypeTweetPoll                 EventType = "tweet.poll"
	EventTypeTweetMention              EventType = "tweet.mention"
	EventTypeTweetHashtag              EventType = "tweet.hashtag"
	EventTypeTweetLongform             EventType = "tweet.longform"
	EventTypeProfileAvatarChanged      EventType = "profile.avatar.changed"
	EventTypeProfileBannerChanged      EventType = "profile.banner.changed"
	EventTypeProfileNameChanged        EventType = "profile.name.changed"
	EventTypeProfileUsernameChanged    EventType = "profile.username.changed"
	EventTypeProfileBioChanged         EventType = "profile.bio.changed"
	EventTypeProfileLocationChanged    EventType = "profile.location.changed"
	EventTypeProfileURLChanged         EventType = "profile.url.changed"
	EventTypeProfileVerifiedChanged    EventType = "profile.verified.changed"
	EventTypeProfileProtectedChanged   EventType = "profile.protected.changed"
	EventTypeProfilePinnedTweetChanged EventType = "profile.pinned_tweet.changed"
	EventTypeProfileUnavailableChanged EventType = "profile.unavailable.changed"
)

// Automatic search, user Tweet, and reply coverage preserves shape, filters,
// aliases, and billing. Follow next_cursor while the response reports more pages.
// An empty filtered page can still require continuation. Unprefixed cursors are
// legacy.
type PaginatedTweets struct {
	HasNextPage   bool          `json:"has_next_page" api:"required"`
	NextCursor    string        `json:"next_cursor" api:"required"`
	Tweets        []SearchTweet `json:"tweets" api:"required"`
	FilteredCount int64         `json:"filtered_count"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		HasNextPage   respjson.Field
		NextCursor    respjson.Field
		Tweets        respjson.Field
		FilteredCount respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaginatedTweets) RawJSON() string { return r.JSON.raw }
func (r *PaginatedTweets) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Profile coverage preserves shape, billing, aliases, and filters. Follow
// next_cursor while the response reports more pages. Unprefixed cursors remain
// legacy.
type PaginatedUsers struct {
	HasNextPage   bool          `json:"has_next_page" api:"required"`
	NextCursor    string        `json:"next_cursor" api:"required"`
	Users         []UserProfile `json:"users" api:"required"`
	FilteredCount int64         `json:"filtered_count"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		HasNextPage   respjson.Field
		NextCursor    respjson.Field
		Users         respjson.Field
		FilteredCount respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaginatedUsers) RawJSON() string { return r.JSON.raw }
func (r *PaginatedUsers) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Tweet returned from search results with inline author info. A zero metric can
// mean X did not report the count.
type SearchTweet struct {
	ID            string `json:"id" api:"required"`
	BookmarkCount int64  `json:"bookmarkCount" api:"required"`
	LikeCount     int64  `json:"likeCount" api:"required"`
	QuoteCount    int64  `json:"quoteCount" api:"required"`
	ReplyCount    int64  `json:"replyCount" api:"required"`
	RetweetCount  int64  `json:"retweetCount" api:"required"`
	Text          string `json:"text" api:"required"`
	ViewCount     int64  `json:"viewCount" api:"required"`
	// Describes an X Article preview and its lifecycle metadata.
	Article SearchTweetArticle `json:"article"`
	// X user profile with bio, follower counts, and verification status.
	Author UserProfile `json:"author"`
	// Describes a public card and its referenced profiles.
	Card SearchTweetCard `json:"card"`
	// Community ID.
	CommunityID string `json:"communityId"`
	// Community Note presentation metadata returned by X.
	CommunityNote SearchTweetCommunityNote `json:"communityNote"`
	// Content disclosure metadata shown by X when a tweet is labeled as paid
	// partnership content or AI-generated media.
	ContentDisclosure ContentDisclosure `json:"contentDisclosure"`
	// Public reply policy and conversation owner.
	ConversationControl SearchTweetConversationControl `json:"conversationControl"`
	ConversationID      string                         `json:"conversationId"`
	CreatedAt           string                         `json:"createdAt"`
	DisplayTextRange    []int64                        `json:"displayTextRange"`
	// Lists edit-chain identifiers and the remaining edit window.
	Edit SearchTweetEdit `json:"edit"`
	// Lists hashtags, symbols, links, and mentions from tweet text.
	Entities SearchTweetEntities `json:"entities"`
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
	LimitedActions []SearchTweetLimitedAction `json:"limitedActions"`
	// Attached media items, omitted when unavailable.
	Media []TweetMedia `json:"media"`
	// Complete Note Tweet content and rich-text metadata.
	NoteTweet SearchTweetNoteTweet `json:"noteTweet"`
	// Describes public place metadata on a geotagged tweet.
	Place             SearchTweetPlace `json:"place"`
	PossiblySensitive bool             `json:"possiblySensitive"`
	// Public metadata whose fields are defined by X.
	PostCta map[string]any `json:"postCta"`
	// Engagement counts retained from a prior tweet edit.
	PreviousCounts SearchTweetPreviousCounts `json:"previousCounts"`
	// Quoted or retweeted tweet context.
	QuotedTweet EmbeddedTweet `json:"quoted_tweet"`
	// Quoted tweet ID.
	QuotedTweetID string `json:"quotedTweetId"`
	// Public post and user referenced by this reaction.
	ReactionContext SearchTweetReactionContext `json:"reactionContext"`
	// Quoted or retweeted tweet context.
	RetweetedTweet EmbeddedTweet `json:"retweeted_tweet"`
	// Public metadata whose fields are defined by X.
	Scopes map[string]any `json:"scopes"`
	Source string         `json:"source"`
	// Public visibility notice attached to an available tweet.
	Tombstone SearchTweetTombstone `json:"tombstone"`
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
func (r SearchTweet) RawJSON() string { return r.JSON.raw }
func (r *SearchTweet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Describes an X Article preview and its lifecycle metadata.
type SearchTweetArticle struct {
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
func (r SearchTweetArticle) RawJSON() string { return r.JSON.raw }
func (r *SearchTweetArticle) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Describes a public card and its referenced profiles.
type SearchTweetCard struct {
	ID string `json:"id"`
	// Public metadata whose fields are defined by X.
	BindingValues map[string]any `json:"bindingValues"`
	Name          string         `json:"name"`
	// Public metadata whose fields are defined by X.
	Platform map[string]any `json:"platform"`
	URL      string         `json:"url"`
	// Unresolved card user references.
	UserReferenceErrors []SearchTweetCardUserReferenceError `json:"userReferenceErrors"`
	UserReferences      []UserProfile                       `json:"userReferences"`
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
func (r SearchTweetCard) RawJSON() string { return r.JSON.raw }
func (r *SearchTweetCard) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SearchTweetCardUserReferenceError struct {
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
func (r SearchTweetCardUserReferenceError) RawJSON() string { return r.JSON.raw }
func (r *SearchTweetCardUserReferenceError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Community Note presentation metadata returned by X.
type SearchTweetCommunityNote struct {
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
func (r SearchTweetCommunityNote) RawJSON() string { return r.JSON.raw }
func (r *SearchTweetCommunityNote) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Public reply policy and conversation owner.
type SearchTweetConversationControl struct {
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
func (r SearchTweetConversationControl) RawJSON() string { return r.JSON.raw }
func (r *SearchTweetConversationControl) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Lists edit-chain identifiers and the remaining edit window.
type SearchTweetEdit struct {
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
func (r SearchTweetEdit) RawJSON() string { return r.JSON.raw }
func (r *SearchTweetEdit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Lists hashtags, symbols, links, and mentions from tweet text.
type SearchTweetEntities struct {
	Hashtags     []SearchTweetEntitiesHashtag     `json:"hashtags"`
	Smarttags    []SearchTweetEntitiesSmarttag    `json:"smarttags"`
	Symbols      []SearchTweetEntitiesSymbol      `json:"symbols"`
	Timestamps   []SearchTweetEntitiesTimestamp   `json:"timestamps"`
	URLs         []SearchTweetEntitiesURL         `json:"urls"`
	UserMentions []SearchTweetEntitiesUserMention `json:"user_mentions"`
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
func (r SearchTweetEntities) RawJSON() string { return r.JSON.raw }
func (r *SearchTweetEntities) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Provides hashtag text and source offsets within a tweet.
type SearchTweetEntitiesHashtag struct {
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
func (r SearchTweetEntitiesHashtag) RawJSON() string { return r.JSON.raw }
func (r *SearchTweetEntitiesHashtag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indexed smart-tag, cashtag, or video timestamp metadata.
type SearchTweetEntitiesSmarttag struct {
	Indices []int64                        `json:"indices"`
	Seconds float64                        `json:"seconds"`
	Tag     SearchTweetEntitiesSmarttagTag `json:"tag"`
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
func (r SearchTweetEntitiesSmarttag) RawJSON() string { return r.JSON.raw }
func (r *SearchTweetEntitiesSmarttag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SearchTweetEntitiesSmarttagTag struct {
	Info SearchTweetEntitiesSmarttagTagInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SearchTweetEntitiesSmarttagTag) RawJSON() string { return r.JSON.raw }
func (r *SearchTweetEntitiesSmarttagTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SearchTweetEntitiesSmarttagTagInfo struct {
	Info SearchTweetEntitiesSmarttagTagInfoInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SearchTweetEntitiesSmarttagTagInfo) RawJSON() string { return r.JSON.raw }
func (r *SearchTweetEntitiesSmarttagTagInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SearchTweetEntitiesSmarttagTagInfoInfo struct {
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
func (r SearchTweetEntitiesSmarttagTagInfoInfo) RawJSON() string { return r.JSON.raw }
func (r *SearchTweetEntitiesSmarttagTagInfoInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indexed smart-tag, cashtag, or video timestamp metadata.
type SearchTweetEntitiesSymbol struct {
	Indices []int64                      `json:"indices"`
	Seconds float64                      `json:"seconds"`
	Tag     SearchTweetEntitiesSymbolTag `json:"tag"`
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
func (r SearchTweetEntitiesSymbol) RawJSON() string { return r.JSON.raw }
func (r *SearchTweetEntitiesSymbol) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SearchTweetEntitiesSymbolTag struct {
	Info SearchTweetEntitiesSymbolTagInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SearchTweetEntitiesSymbolTag) RawJSON() string { return r.JSON.raw }
func (r *SearchTweetEntitiesSymbolTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SearchTweetEntitiesSymbolTagInfo struct {
	Info SearchTweetEntitiesSymbolTagInfoInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SearchTweetEntitiesSymbolTagInfo) RawJSON() string { return r.JSON.raw }
func (r *SearchTweetEntitiesSymbolTagInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SearchTweetEntitiesSymbolTagInfoInfo struct {
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
func (r SearchTweetEntitiesSymbolTagInfoInfo) RawJSON() string { return r.JSON.raw }
func (r *SearchTweetEntitiesSymbolTagInfoInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indexed smart-tag, cashtag, or video timestamp metadata.
type SearchTweetEntitiesTimestamp struct {
	Indices []int64                         `json:"indices"`
	Seconds float64                         `json:"seconds"`
	Tag     SearchTweetEntitiesTimestampTag `json:"tag"`
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
func (r SearchTweetEntitiesTimestamp) RawJSON() string { return r.JSON.raw }
func (r *SearchTweetEntitiesTimestamp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SearchTweetEntitiesTimestampTag struct {
	Info SearchTweetEntitiesTimestampTagInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SearchTweetEntitiesTimestampTag) RawJSON() string { return r.JSON.raw }
func (r *SearchTweetEntitiesTimestampTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SearchTweetEntitiesTimestampTagInfo struct {
	Info SearchTweetEntitiesTimestampTagInfoInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SearchTweetEntitiesTimestampTagInfo) RawJSON() string { return r.JSON.raw }
func (r *SearchTweetEntitiesTimestampTagInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SearchTweetEntitiesTimestampTagInfoInfo struct {
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
func (r SearchTweetEntitiesTimestampTagInfoInfo) RawJSON() string { return r.JSON.raw }
func (r *SearchTweetEntitiesTimestampTagInfoInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Provides shortened, display, and expanded URLs from tweet text.
type SearchTweetEntitiesURL struct {
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
func (r SearchTweetEntitiesURL) RawJSON() string { return r.JSON.raw }
func (r *SearchTweetEntitiesURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Provides profile identity and source offsets for a mention.
type SearchTweetEntitiesUserMention struct {
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
func (r SearchTweetEntitiesUserMention) RawJSON() string { return r.JSON.raw }
func (r *SearchTweetEntitiesUserMention) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SearchTweetLimitedAction struct {
	Action string                         `json:"action"`
	Prompt SearchTweetLimitedActionPrompt `json:"prompt"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Action      respjson.Field
		Prompt      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SearchTweetLimitedAction) RawJSON() string { return r.JSON.raw }
func (r *SearchTweetLimitedAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SearchTweetLimitedActionPrompt struct {
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
func (r SearchTweetLimitedActionPrompt) RawJSON() string { return r.JSON.raw }
func (r *SearchTweetLimitedActionPrompt) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Complete Note Tweet content and rich-text metadata.
type SearchTweetNoteTweet struct {
	Text string `json:"text" api:"required"`
	ID   string `json:"id"`
	// Lists hashtags, symbols, links, and mentions from tweet text.
	Entities SearchTweetNoteTweetEntities `json:"entities"`
	// Inline media positions in the Note Tweet text.
	InlineMedia  []SearchTweetNoteTweetInlineMedia `json:"inlineMedia"`
	IsExpandable bool                              `json:"isExpandable"`
	RichtextTags []SearchTweetNoteTweetRichtextTag `json:"richtextTags"`
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
func (r SearchTweetNoteTweet) RawJSON() string { return r.JSON.raw }
func (r *SearchTweetNoteTweet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Lists hashtags, symbols, links, and mentions from tweet text.
type SearchTweetNoteTweetEntities struct {
	Hashtags     []SearchTweetNoteTweetEntitiesHashtag     `json:"hashtags"`
	Smarttags    []SearchTweetNoteTweetEntitiesSmarttag    `json:"smarttags"`
	Symbols      []SearchTweetNoteTweetEntitiesSymbol      `json:"symbols"`
	Timestamps   []SearchTweetNoteTweetEntitiesTimestamp   `json:"timestamps"`
	URLs         []SearchTweetNoteTweetEntitiesURL         `json:"urls"`
	UserMentions []SearchTweetNoteTweetEntitiesUserMention `json:"user_mentions"`
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
func (r SearchTweetNoteTweetEntities) RawJSON() string { return r.JSON.raw }
func (r *SearchTweetNoteTweetEntities) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Provides hashtag text and source offsets within a tweet.
type SearchTweetNoteTweetEntitiesHashtag struct {
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
func (r SearchTweetNoteTweetEntitiesHashtag) RawJSON() string { return r.JSON.raw }
func (r *SearchTweetNoteTweetEntitiesHashtag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indexed smart-tag, cashtag, or video timestamp metadata.
type SearchTweetNoteTweetEntitiesSmarttag struct {
	Indices []int64                                 `json:"indices"`
	Seconds float64                                 `json:"seconds"`
	Tag     SearchTweetNoteTweetEntitiesSmarttagTag `json:"tag"`
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
func (r SearchTweetNoteTweetEntitiesSmarttag) RawJSON() string { return r.JSON.raw }
func (r *SearchTweetNoteTweetEntitiesSmarttag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SearchTweetNoteTweetEntitiesSmarttagTag struct {
	Info SearchTweetNoteTweetEntitiesSmarttagTagInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SearchTweetNoteTweetEntitiesSmarttagTag) RawJSON() string { return r.JSON.raw }
func (r *SearchTweetNoteTweetEntitiesSmarttagTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SearchTweetNoteTweetEntitiesSmarttagTagInfo struct {
	Info SearchTweetNoteTweetEntitiesSmarttagTagInfoInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SearchTweetNoteTweetEntitiesSmarttagTagInfo) RawJSON() string { return r.JSON.raw }
func (r *SearchTweetNoteTweetEntitiesSmarttagTagInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SearchTweetNoteTweetEntitiesSmarttagTagInfoInfo struct {
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
func (r SearchTweetNoteTweetEntitiesSmarttagTagInfoInfo) RawJSON() string { return r.JSON.raw }
func (r *SearchTweetNoteTweetEntitiesSmarttagTagInfoInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indexed smart-tag, cashtag, or video timestamp metadata.
type SearchTweetNoteTweetEntitiesSymbol struct {
	Indices []int64                               `json:"indices"`
	Seconds float64                               `json:"seconds"`
	Tag     SearchTweetNoteTweetEntitiesSymbolTag `json:"tag"`
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
func (r SearchTweetNoteTweetEntitiesSymbol) RawJSON() string { return r.JSON.raw }
func (r *SearchTweetNoteTweetEntitiesSymbol) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SearchTweetNoteTweetEntitiesSymbolTag struct {
	Info SearchTweetNoteTweetEntitiesSymbolTagInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SearchTweetNoteTweetEntitiesSymbolTag) RawJSON() string { return r.JSON.raw }
func (r *SearchTweetNoteTweetEntitiesSymbolTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SearchTweetNoteTweetEntitiesSymbolTagInfo struct {
	Info SearchTweetNoteTweetEntitiesSymbolTagInfoInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SearchTweetNoteTweetEntitiesSymbolTagInfo) RawJSON() string { return r.JSON.raw }
func (r *SearchTweetNoteTweetEntitiesSymbolTagInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SearchTweetNoteTweetEntitiesSymbolTagInfoInfo struct {
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
func (r SearchTweetNoteTweetEntitiesSymbolTagInfoInfo) RawJSON() string { return r.JSON.raw }
func (r *SearchTweetNoteTweetEntitiesSymbolTagInfoInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indexed smart-tag, cashtag, or video timestamp metadata.
type SearchTweetNoteTweetEntitiesTimestamp struct {
	Indices []int64                                  `json:"indices"`
	Seconds float64                                  `json:"seconds"`
	Tag     SearchTweetNoteTweetEntitiesTimestampTag `json:"tag"`
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
func (r SearchTweetNoteTweetEntitiesTimestamp) RawJSON() string { return r.JSON.raw }
func (r *SearchTweetNoteTweetEntitiesTimestamp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SearchTweetNoteTweetEntitiesTimestampTag struct {
	Info SearchTweetNoteTweetEntitiesTimestampTagInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SearchTweetNoteTweetEntitiesTimestampTag) RawJSON() string { return r.JSON.raw }
func (r *SearchTweetNoteTweetEntitiesTimestampTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SearchTweetNoteTweetEntitiesTimestampTagInfo struct {
	Info SearchTweetNoteTweetEntitiesTimestampTagInfoInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SearchTweetNoteTweetEntitiesTimestampTagInfo) RawJSON() string { return r.JSON.raw }
func (r *SearchTweetNoteTweetEntitiesTimestampTagInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SearchTweetNoteTweetEntitiesTimestampTagInfoInfo struct {
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
func (r SearchTweetNoteTweetEntitiesTimestampTagInfoInfo) RawJSON() string { return r.JSON.raw }
func (r *SearchTweetNoteTweetEntitiesTimestampTagInfoInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Provides shortened, display, and expanded URLs from tweet text.
type SearchTweetNoteTweetEntitiesURL struct {
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
func (r SearchTweetNoteTweetEntitiesURL) RawJSON() string { return r.JSON.raw }
func (r *SearchTweetNoteTweetEntitiesURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Provides profile identity and source offsets for a mention.
type SearchTweetNoteTweetEntitiesUserMention struct {
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
func (r SearchTweetNoteTweetEntitiesUserMention) RawJSON() string { return r.JSON.raw }
func (r *SearchTweetNoteTweetEntitiesUserMention) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SearchTweetNoteTweetInlineMedia struct {
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
func (r SearchTweetNoteTweetInlineMedia) RawJSON() string { return r.JSON.raw }
func (r *SearchTweetNoteTweetInlineMedia) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SearchTweetNoteTweetRichtextTag struct {
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
func (r SearchTweetNoteTweetRichtextTag) RawJSON() string { return r.JSON.raw }
func (r *SearchTweetNoteTweetRichtextTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Describes public place metadata on a geotagged tweet.
type SearchTweetPlace struct {
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
func (r SearchTweetPlace) RawJSON() string { return r.JSON.raw }
func (r *SearchTweetPlace) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Engagement counts retained from a prior tweet edit.
type SearchTweetPreviousCounts struct {
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
func (r SearchTweetPreviousCounts) RawJSON() string { return r.JSON.raw }
func (r *SearchTweetPreviousCounts) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Public post and user referenced by this reaction.
type SearchTweetReactionContext struct {
	// Referenced post ID.
	ReactedToPostID string `json:"reactedToPostId"`
	// X user profile with bio, follower counts, and verification status.
	ReactedToUser UserProfile `json:"reactedToUser"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ReactedToPostID respjson.Field
		ReactedToUser   respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SearchTweetReactionContext) RawJSON() string { return r.JSON.raw }
func (r *SearchTweetReactionContext) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Public visibility notice attached to an available tweet.
type SearchTweetTombstone struct {
	Text SearchTweetTombstoneText `json:"text"`
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
func (r SearchTweetTombstone) RawJSON() string { return r.JSON.raw }
func (r *SearchTweetTombstone) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SearchTweetTombstoneText struct {
	Entities []SearchTweetTombstoneTextEntity `json:"entities"`
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
func (r SearchTweetTombstoneText) RawJSON() string { return r.JSON.raw }
func (r *SearchTweetTombstoneText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SearchTweetTombstoneTextEntity struct {
	FromIndex int64                             `json:"fromIndex"`
	Ref       SearchTweetTombstoneTextEntityRef `json:"ref"`
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
func (r SearchTweetTombstoneTextEntity) RawJSON() string { return r.JSON.raw }
func (r *SearchTweetTombstoneTextEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SearchTweetTombstoneTextEntityRef struct {
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
func (r SearchTweetTombstoneTextEntityRef) RawJSON() string { return r.JSON.raw }
func (r *SearchTweetTombstoneTextEntityRef) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Tweet media.
type TweetMedia struct {
	// Preview URL.
	MediaURL string `json:"mediaUrl" api:"required"`
	// Any of "photo", "video", "animated_gif".
	Type TweetMediaType `json:"type" api:"required"`
	// Tweet media link.
	URL string `json:"url" api:"required"`
	// Media entity ID.
	ID string `json:"id"`
	// Adult-content warning.
	AdultContent bool `json:"adultContent"`
	// Direct download permission.
	AllowDownload bool `json:"allowDownload"`
	// Accessibility text.
	AltText string `json:"altText"`
	// Video width and height ratio.
	AspectRatio []int64 `json:"aspectRatio"`
	// Availability reason.
	AvailabilityReason string `json:"availabilityReason"`
	// Availability state.
	AvailabilityStatus string `json:"availabilityStatus"`
	// Media description.
	Description string `json:"description"`
	// Display URL.
	DisplayURL string `json:"displayUrl"`
	// Video duration in milliseconds.
	DurationMillis int64 `json:"durationMillis"`
	// Embeddable status.
	Embeddable bool `json:"embeddable"`
	// Expanded media URL.
	ExpandedURL string `json:"expandedUrl"`
	// Face crop rectangles by size.
	FaceRects map[string][]TweetMediaFaceRect `json:"faceRects"`
	// Suggested image crops.
	FocusRects []TweetMediaFocusRect `json:"focusRects"`
	// Graphic-violence warning.
	GraphicViolence bool `json:"graphicViolence"`
	// Grok post ID associated with the media.
	GrokPostID string `json:"grokPostId"`
	// Original height.
	Height int64 `json:"height"`
	// Tweet text offsets.
	Indices []int64 `json:"indices"`
	// Stable X media key.
	MediaKey string `json:"mediaKey"`
	// Monetization status.
	Monetizable           bool `json:"monetizable"`
	OtherSensitiveContent bool `json:"otherSensitiveContent"`
	// Named media renditions and resize modes.
	Sizes map[string]TweetMediaSize `json:"sizes"`
	// Source tweet ID for copied media.
	SourceStatusID string `json:"sourceStatusId"`
	// Source profile ID for copied media.
	SourceUserID string `json:"sourceUserId"`
	// Public profiles tagged in the media.
	Tags []TweetMediaTag `json:"tags"`
	// Media title.
	Title string `json:"title"`
	// Video encodings in source order.
	VideoVariants []TweetMediaVideoVariant `json:"videoVariants"`
	// Public destination URL.
	VisitSiteURL string `json:"visitSiteUrl"`
	// Public media action URL.
	WatchNowURL string `json:"watchNowUrl"`
	// Original width.
	Width int64 `json:"width"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MediaURL              respjson.Field
		Type                  respjson.Field
		URL                   respjson.Field
		ID                    respjson.Field
		AdultContent          respjson.Field
		AllowDownload         respjson.Field
		AltText               respjson.Field
		AspectRatio           respjson.Field
		AvailabilityReason    respjson.Field
		AvailabilityStatus    respjson.Field
		Description           respjson.Field
		DisplayURL            respjson.Field
		DurationMillis        respjson.Field
		Embeddable            respjson.Field
		ExpandedURL           respjson.Field
		FaceRects             respjson.Field
		FocusRects            respjson.Field
		GraphicViolence       respjson.Field
		GrokPostID            respjson.Field
		Height                respjson.Field
		Indices               respjson.Field
		MediaKey              respjson.Field
		Monetizable           respjson.Field
		OtherSensitiveContent respjson.Field
		Sizes                 respjson.Field
		SourceStatusID        respjson.Field
		SourceUserID          respjson.Field
		Tags                  respjson.Field
		Title                 respjson.Field
		VideoVariants         respjson.Field
		VisitSiteURL          respjson.Field
		WatchNowURL           respjson.Field
		Width                 respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TweetMedia) RawJSON() string { return r.JSON.raw }
func (r *TweetMedia) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TweetMediaType string

const (
	TweetMediaTypePhoto       TweetMediaType = "photo"
	TweetMediaTypeVideo       TweetMediaType = "video"
	TweetMediaTypeAnimatedGif TweetMediaType = "animated_gif"
)

type TweetMediaFaceRect struct {
	H int64 `json:"h" api:"required"`
	W int64 `json:"w" api:"required"`
	X int64 `json:"x" api:"required"`
	Y int64 `json:"y" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		H           respjson.Field
		W           respjson.Field
		X           respjson.Field
		Y           respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TweetMediaFaceRect) RawJSON() string { return r.JSON.raw }
func (r *TweetMediaFaceRect) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TweetMediaFocusRect struct {
	H int64 `json:"h" api:"required"`
	W int64 `json:"w" api:"required"`
	X int64 `json:"x" api:"required"`
	Y int64 `json:"y" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		H           respjson.Field
		W           respjson.Field
		X           respjson.Field
		Y           respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TweetMediaFocusRect) RawJSON() string { return r.JSON.raw }
func (r *TweetMediaFocusRect) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TweetMediaSize struct {
	H      int64  `json:"h" api:"required"`
	Resize string `json:"resize" api:"required"`
	W      int64  `json:"w" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		H           respjson.Field
		Resize      respjson.Field
		W           respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TweetMediaSize) RawJSON() string { return r.JSON.raw }
func (r *TweetMediaSize) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TweetMediaTag struct {
	Name       string `json:"name"`
	ScreenName string `json:"screen_name"`
	Type       string `json:"type"`
	UserID     string `json:"user_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		ScreenName  respjson.Field
		Type        respjson.Field
		UserID      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TweetMediaTag) RawJSON() string { return r.JSON.raw }
func (r *TweetMediaTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TweetMediaVideoVariant struct {
	ContentType string `json:"contentType" api:"required"`
	URL         string `json:"url" api:"required"`
	Bitrate     int64  `json:"bitrate"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ContentType respjson.Field
		URL         respjson.Field
		Bitrate     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TweetMediaVideoVariant) RawJSON() string { return r.JSON.raw }
func (r *TweetMediaVideoVariant) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// X user profile with bio, follower counts, and verification status.
type UserProfile struct {
	ID       string `json:"id" api:"required"`
	Name     string `json:"name" api:"required"`
	Username string `json:"username" api:"required"`
	// X's best-effort public label inferred from aggregated account-access IP
	// addresses. It does not state nationality, residence, identity, registration,
	// post location, or exact location.
	AccountBasedIn UserProfileAccountBasedIn `json:"accountBasedIn" api:"nullable"`
	// Organization affiliation label shown on an X profile.
	AffiliatesHighlightedLabel     UserProfileAffiliatesHighlightedLabel `json:"affiliatesHighlightedLabel"`
	AutomatedBy                    string                                `json:"automatedBy"`
	BusinessAccountAffiliatesCount int64                                 `json:"businessAccountAffiliatesCount"`
	// Community role when returned by community member reads
	CommunityRole             string `json:"communityRole"`
	CoverPicture              string `json:"coverPicture"`
	CreatedAt                 string `json:"createdAt"`
	CreatorSubscriptionsCount int64  `json:"creatorSubscriptionsCount"`
	Description               string `json:"description"`
	FavouritesCount           int64  `json:"favouritesCount"`
	Followers                 int64  `json:"followers"`
	Following                 int64  `json:"following"`
	// Public profile bio translation returned by X
	GrokTranslatedBio               map[string]any `json:"grokTranslatedBio"`
	HasCustomTimelines              bool           `json:"hasCustomTimelines"`
	HasGraduatedAccess              bool           `json:"hasGraduatedAccess"`
	HasHiddenSubscriptionsOnProfile bool           `json:"hasHiddenSubscriptionsOnProfile"`
	// Profile highlight availability and count metadata.
	HighlightsInfo UserProfileHighlightsInfo `json:"highlightsInfo"`
	// Identity verification metadata displayed by X.
	IdentityVerification UserProfileIdentityVerification `json:"identityVerification"`
	IsAutomated          bool                            `json:"isAutomated"`
	// Whether X shows a blue verification badge
	IsBlueVerified        bool `json:"isBlueVerified"`
	IsProfileTranslatable bool `json:"isProfileTranslatable"`
	IsTranslator          bool `json:"isTranslator"`
	// Whether X marks the profile as verified
	IsVerified bool `json:"isVerified"`
	// Account owner's public profile location text
	Location                 string   `json:"location"`
	MediaCount               int64    `json:"mediaCount"`
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
	ProfilePicture             string `json:"profilePicture"`
	ProfileSortEnabled         bool   `json:"profileSortEnabled"`
	ProfileTranslatorType      string `json:"profileTranslatorType"`
	// Whether the profile protects its posts
	Protected           bool  `json:"protected"`
	StatusesCount       int64 `json:"statusesCount"`
	SuperFollowEligible bool  `json:"superFollowEligible"`
	// Whether X marks the subscription profile as active.
	SuperFollowsUserProfileActive bool `json:"superFollowsUserProfileActive"`
	// Public payment and creator-support handles shown on X.
	TipJar              UserProfileTipJar `json:"tipJar"`
	Unavailable         bool              `json:"unavailable"`
	UnavailableReason   string            `json:"unavailableReason"`
	URL                 string            `json:"url"`
	Verified            bool              `json:"verified"`
	VerifiedType        string            `json:"verifiedType"`
	WithheldInCountries []string          `json:"withheldInCountries"`
	// Whether X withholds a post or user
	WithheldScope string `json:"withheldScope"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                              respjson.Field
		Name                            respjson.Field
		Username                        respjson.Field
		AccountBasedIn                  respjson.Field
		AffiliatesHighlightedLabel      respjson.Field
		AutomatedBy                     respjson.Field
		BusinessAccountAffiliatesCount  respjson.Field
		CommunityRole                   respjson.Field
		CoverPicture                    respjson.Field
		CreatedAt                       respjson.Field
		CreatorSubscriptionsCount       respjson.Field
		Description                     respjson.Field
		FavouritesCount                 respjson.Field
		Followers                       respjson.Field
		Following                       respjson.Field
		GrokTranslatedBio               respjson.Field
		HasCustomTimelines              respjson.Field
		HasGraduatedAccess              respjson.Field
		HasHiddenSubscriptionsOnProfile respjson.Field
		HighlightsInfo                  respjson.Field
		IdentityVerification            respjson.Field
		IsAutomated                     respjson.Field
		IsBlueVerified                  respjson.Field
		IsProfileTranslatable           respjson.Field
		IsTranslator                    respjson.Field
		IsVerified                      respjson.Field
		Location                        respjson.Field
		MediaCount                      respjson.Field
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
		ProfileSortEnabled              respjson.Field
		ProfileTranslatorType           respjson.Field
		Protected                       respjson.Field
		StatusesCount                   respjson.Field
		SuperFollowEligible             respjson.Field
		SuperFollowsUserProfileActive   respjson.Field
		TipJar                          respjson.Field
		Unavailable                     respjson.Field
		UnavailableReason               respjson.Field
		URL                             respjson.Field
		Verified                        respjson.Field
		VerifiedType                    respjson.Field
		WithheldInCountries             respjson.Field
		WithheldScope                   respjson.Field
		ExtraFields                     map[string]respjson.Field
		raw                             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserProfile) RawJSON() string { return r.JSON.raw }
func (r *UserProfile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// X's best-effort public label inferred from aggregated account-access IP
// addresses. It does not state nationality, residence, identity, registration,
// post location, or exact location.
type UserProfileAccountBasedIn struct {
	// Any of "country", "region".
	Level      string    `json:"level" api:"required"`
	ObservedAt time.Time `json:"observedAt" api:"required" format:"date-time"`
	Value      string    `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Level       respjson.Field
		ObservedAt  respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserProfileAccountBasedIn) RawJSON() string { return r.JSON.raw }
func (r *UserProfileAccountBasedIn) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Organization affiliation label shown on an X profile.
type UserProfileAffiliatesHighlightedLabel struct {
	BadgeURL    string `json:"badgeUrl"`
	Description string `json:"description"`
	// Public text, ranges, references, and mention data.
	LongDescription      map[string]any `json:"longDescription"`
	URL                  string         `json:"url"`
	URLType              string         `json:"urlType"`
	UserLabelDisplayType string         `json:"userLabelDisplayType"`
	UserLabelType        string         `json:"userLabelType"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BadgeURL             respjson.Field
		Description          respjson.Field
		LongDescription      respjson.Field
		URL                  respjson.Field
		URLType              respjson.Field
		UserLabelDisplayType respjson.Field
		UserLabelType        respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserProfileAffiliatesHighlightedLabel) RawJSON() string { return r.JSON.raw }
func (r *UserProfileAffiliatesHighlightedLabel) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Profile highlight availability and count metadata.
type UserProfileHighlightsInfo struct {
	CanHighlightTweets bool   `json:"canHighlightTweets"`
	HighlightedTweets  string `json:"highlightedTweets"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CanHighlightTweets respjson.Field
		HighlightedTweets  respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserProfileHighlightsInfo) RawJSON() string { return r.JSON.raw }
func (r *UserProfileHighlightsInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Identity verification metadata displayed by X.
type UserProfileIdentityVerification struct {
	Description        string `json:"description"`
	IsIdentityVerified bool   `json:"isIdentityVerified"`
	VerifiedSinceMsec  string `json:"verifiedSinceMsec"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Description        respjson.Field
		IsIdentityVerified respjson.Field
		VerifiedSinceMsec  respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserProfileIdentityVerification) RawJSON() string { return r.JSON.raw }
func (r *UserProfileIdentityVerification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Public payment and creator-support handles shown on X.
type UserProfileTipJar struct {
	BandcampHandle string `json:"bandcampHandle"`
	BitcoinHandle  string `json:"bitcoinHandle"`
	CashAppHandle  string `json:"cashAppHandle"`
	EthereumHandle string `json:"ethereumHandle"`
	GofundmeHandle string `json:"gofundmeHandle"`
	IsEnabled      bool   `json:"isEnabled"`
	PatreonHandle  string `json:"patreonHandle"`
	PayPalHandle   string `json:"payPalHandle"`
	VenmoHandle    string `json:"venmoHandle"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BandcampHandle respjson.Field
		BitcoinHandle  respjson.Field
		CashAppHandle  respjson.Field
		EthereumHandle respjson.Field
		GofundmeHandle respjson.Field
		IsEnabled      respjson.Field
		PatreonHandle  respjson.Field
		PayPalHandle   respjson.Field
		VenmoHandle    respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserProfileTipJar) RawJSON() string { return r.JSON.raw }
func (r *UserProfileTipJar) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
