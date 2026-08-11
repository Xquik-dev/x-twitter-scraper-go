// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

import (
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

// Quoted or retweeted tweet context. Every object includes id, text, and
// engagement metrics. A zero metric can mean X did not report the count. Author,
// media, and conversation fields appear when available.
type EmbeddedTweet struct {
	ID            string `json:"id" api:"required"`
	BookmarkCount int64  `json:"bookmarkCount" api:"required"`
	LikeCount     int64  `json:"likeCount" api:"required"`
	QuoteCount    int64  `json:"quoteCount" api:"required"`
	ReplyCount    int64  `json:"replyCount" api:"required"`
	RetweetCount  int64  `json:"retweetCount" api:"required"`
	Text          string `json:"text" api:"required"`
	ViewCount     int64  `json:"viewCount" api:"required"`
	// Article metadata attached to a tweet.
	Article EmbeddedTweetArticle `json:"article"`
	// X user profile with bio, follower counts, and verification status.
	Author UserProfile `json:"author"`
	// Public card metadata attached to a tweet.
	Card EmbeddedTweetCard `json:"card"`
	// Community Note presentation metadata returned by X.
	CommunityNote EmbeddedTweetCommunityNote `json:"communityNote"`
	// Content disclosure metadata shown by X when a tweet is labeled as paid
	// partnership content or AI-generated media.
	ContentDisclosure ContentDisclosure `json:"contentDisclosure"`
	ConversationID    string            `json:"conversationId"`
	CreatedAt         string            `json:"createdAt"`
	DisplayTextRange  []int64           `json:"displayTextRange"`
	// Edit history metadata returned by X.
	Edit              EmbeddedTweetEdit `json:"edit"`
	Entities          map[string]any    `json:"entities"`
	InReplyToID       string            `json:"inReplyToId"`
	InReplyToUserID   string            `json:"inReplyToUserId"`
	InReplyToUsername string            `json:"inReplyToUsername"`
	IsLimitedReply    bool              `json:"isLimitedReply"`
	IsNoteTweet       bool              `json:"isNoteTweet"`
	IsQuoteStatus     bool              `json:"isQuoteStatus"`
	IsReply           bool              `json:"isReply"`
	IsTranslatable    bool              `json:"isTranslatable"`
	Lang              string            `json:"lang"`
	Media             []TweetMedia      `json:"media"`
	// Complete Note Tweet content and rich-text metadata.
	NoteTweet EmbeddedTweetNoteTweet `json:"noteTweet"`
	// Public place metadata attached to a tweet.
	Place             EmbeddedTweetPlace `json:"place"`
	PossiblySensitive bool               `json:"possiblySensitive"`
	// Engagement counts retained from a prior tweet edit.
	PreviousCounts EmbeddedTweetPreviousCounts `json:"previousCounts"`
	// Quoted or retweeted tweet context. Every object includes id, text, and
	// engagement metrics. A zero metric can mean X did not report the count. Author,
	// media, and conversation fields appear when available.
	QuotedTweet *EmbeddedTweet `json:"quoted_tweet"`
	// Quoted or retweeted tweet context. Every object includes id, text, and
	// engagement metrics. A zero metric can mean X did not report the count. Author,
	// media, and conversation fields appear when available.
	RetweetedTweet *EmbeddedTweet `json:"retweeted_tweet"`
	Source         string         `json:"source"`
	Type           string         `json:"type"`
	URL            string         `json:"url"`
	ViewState      string         `json:"viewState"`
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
		Article           respjson.Field
		Author            respjson.Field
		Card              respjson.Field
		CommunityNote     respjson.Field
		ContentDisclosure respjson.Field
		ConversationID    respjson.Field
		CreatedAt         respjson.Field
		DisplayTextRange  respjson.Field
		Edit              respjson.Field
		Entities          respjson.Field
		InReplyToID       respjson.Field
		InReplyToUserID   respjson.Field
		InReplyToUsername respjson.Field
		IsLimitedReply    respjson.Field
		IsNoteTweet       respjson.Field
		IsQuoteStatus     respjson.Field
		IsReply           respjson.Field
		IsTranslatable    respjson.Field
		Lang              respjson.Field
		Media             respjson.Field
		NoteTweet         respjson.Field
		Place             respjson.Field
		PossiblySensitive respjson.Field
		PreviousCounts    respjson.Field
		QuotedTweet       respjson.Field
		RetweetedTweet    respjson.Field
		Source            respjson.Field
		Type              respjson.Field
		URL               respjson.Field
		ViewState         respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweet) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Article metadata attached to a tweet.
type EmbeddedTweetArticle struct {
	ID            string `json:"id"`
	CoverMediaURL string `json:"coverMediaUrl"`
	PreviewText   string `json:"previewText"`
	Title         string `json:"title"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		CoverMediaURL respjson.Field
		PreviewText   respjson.Field
		Title         respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetArticle) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetArticle) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Public card metadata attached to a tweet.
type EmbeddedTweetCard struct {
	ID            string         `json:"id"`
	BindingValues map[string]any `json:"bindingValues"`
	Name          string         `json:"name"`
	URL           string         `json:"url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		BindingValues respjson.Field
		Name          respjson.Field
		URL           respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetCard) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetCard) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Community Note presentation metadata returned by X.
type EmbeddedTweetCommunityNote struct {
	ID             string `json:"id"`
	DestinationURL string `json:"destinationUrl"`
	Footer         string `json:"footer"`
	ShortTitle     string `json:"shortTitle"`
	Subtitle       string `json:"subtitle"`
	Title          string `json:"title"`
	VisualStyle    string `json:"visualStyle"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		DestinationURL respjson.Field
		Footer         respjson.Field
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

// Edit history metadata returned by X.
type EmbeddedTweetEdit struct {
	EditableUntilMsecs string   `json:"editableUntilMsecs"`
	EditTweetIDs       []string `json:"editTweetIds"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EditableUntilMsecs respjson.Field
		EditTweetIDs       respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweetEdit) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweetEdit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Complete Note Tweet content and rich-text metadata.
type EmbeddedTweetNoteTweet struct {
	Text         string                              `json:"text" api:"required"`
	ID           string                              `json:"id"`
	Entities     map[string]any                      `json:"entities"`
	IsExpandable bool                                `json:"isExpandable"`
	RichtextTags []EmbeddedTweetNoteTweetRichtextTag `json:"richtextTags"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Text         respjson.Field
		ID           respjson.Field
		Entities     respjson.Field
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

// Public place metadata attached to a tweet.
type EmbeddedTweetPlace struct {
	ID          string         `json:"id"`
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

// No-mode search, user Tweet, user reply, and direct reply reads use automatic
// coverage. Shape, filters, aliases, and billing stay compatible. Unprefixed
// cursors remain legacy. Follow next_cursor while has_next_page is true. An empty
// filtered page can still have has_next_page true.
type PaginatedTweets struct {
	HasNextPage bool          `json:"has_next_page" api:"required"`
	NextCursor  string        `json:"next_cursor" api:"required"`
	Tweets      []SearchTweet `json:"tweets" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		HasNextPage respjson.Field
		NextCursor  respjson.Field
		Tweets      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaginatedTweets) RawJSON() string { return r.JSON.raw }
func (r *PaginatedTweets) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Paginated user profiles. No-mode follower, following, and verified follower
// requests merge independent views automatically. Response fields, page size,
// aliases, filters, and per-returned-profile billing stay unchanged. Existing
// unprefixed cursors retain legacy behavior. Follow next_cursor while
// has_next_page is true.
type PaginatedUsers struct {
	HasNextPage bool          `json:"has_next_page" api:"required"`
	NextCursor  string        `json:"next_cursor" api:"required"`
	Users       []UserProfile `json:"users" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		HasNextPage respjson.Field
		NextCursor  respjson.Field
		Users       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
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
	// Article metadata attached to a tweet.
	Article SearchTweetArticle `json:"article"`
	// X user profile with bio, follower counts, and verification status.
	Author UserProfile `json:"author"`
	// Public card metadata attached to a tweet.
	Card SearchTweetCard `json:"card"`
	// Community Note presentation metadata returned by X.
	CommunityNote SearchTweetCommunityNote `json:"communityNote"`
	// Content disclosure metadata shown by X when a tweet is labeled as paid
	// partnership content or AI-generated media.
	ContentDisclosure ContentDisclosure `json:"contentDisclosure"`
	// Root tweet ID for the search result conversation
	ConversationID string `json:"conversationId"`
	CreatedAt      string `json:"createdAt"`
	// Rendered text's start and end offsets.
	DisplayTextRange []int64 `json:"displayTextRange"`
	// Edit history metadata returned by X.
	Edit SearchTweetEdit `json:"edit"`
	// Parsed search-result entities including URLs, mentions, hashtags, and media
	// markers
	Entities map[string]any `json:"entities"`
	// ID of the tweet this result replies to.
	InReplyToID string `json:"inReplyToId"`
	// ID of the user this result replies to.
	InReplyToUserID string `json:"inReplyToUserId"`
	// Username this result replies to.
	InReplyToUsername string `json:"inReplyToUsername"`
	// Whether the tweet has limited reply permissions
	IsLimitedReply bool `json:"isLimitedReply"`
	// True for Note Tweets (long-form content, up to 25,000 characters)
	IsNoteTweet bool `json:"isNoteTweet"`
	// True when this search result quotes another tweet
	IsQuoteStatus bool `json:"isQuoteStatus"`
	// True when this search result is a reply
	IsReply        bool `json:"isReply"`
	IsTranslatable bool `json:"isTranslatable"`
	// Search result language code.
	Lang string `json:"lang"`
	// Search-result media attachments, omitted when no media is present
	Media []TweetMedia `json:"media"`
	// Complete Note Tweet content and rich-text metadata.
	NoteTweet SearchTweetNoteTweet `json:"noteTweet"`
	// Public place metadata attached to a tweet.
	Place             SearchTweetPlace `json:"place"`
	PossiblySensitive bool             `json:"possiblySensitive"`
	// Engagement counts retained from a prior tweet edit.
	PreviousCounts SearchTweetPreviousCounts `json:"previousCounts"`
	// Quoted or retweeted tweet context. Every object includes id, text, and
	// engagement metrics. A zero metric can mean X did not report the count. Author,
	// media, and conversation fields appear when available.
	QuotedTweet EmbeddedTweet `json:"quoted_tweet"`
	// Quoted or retweeted tweet context. Every object includes id, text, and
	// engagement metrics. A zero metric can mean X did not report the count. Author,
	// media, and conversation fields appear when available.
	RetweetedTweet EmbeddedTweet `json:"retweeted_tweet"`
	// Client application used to post the tweet
	Source string `json:"source"`
	Type   string `json:"type"`
	// Search result permalink.
	URL       string `json:"url"`
	ViewState string `json:"viewState"`
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
		Article           respjson.Field
		Author            respjson.Field
		Card              respjson.Field
		CommunityNote     respjson.Field
		ContentDisclosure respjson.Field
		ConversationID    respjson.Field
		CreatedAt         respjson.Field
		DisplayTextRange  respjson.Field
		Edit              respjson.Field
		Entities          respjson.Field
		InReplyToID       respjson.Field
		InReplyToUserID   respjson.Field
		InReplyToUsername respjson.Field
		IsLimitedReply    respjson.Field
		IsNoteTweet       respjson.Field
		IsQuoteStatus     respjson.Field
		IsReply           respjson.Field
		IsTranslatable    respjson.Field
		Lang              respjson.Field
		Media             respjson.Field
		NoteTweet         respjson.Field
		Place             respjson.Field
		PossiblySensitive respjson.Field
		PreviousCounts    respjson.Field
		QuotedTweet       respjson.Field
		RetweetedTweet    respjson.Field
		Source            respjson.Field
		Type              respjson.Field
		URL               respjson.Field
		ViewState         respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SearchTweet) RawJSON() string { return r.JSON.raw }
func (r *SearchTweet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Article metadata attached to a tweet.
type SearchTweetArticle struct {
	ID            string `json:"id"`
	CoverMediaURL string `json:"coverMediaUrl"`
	PreviewText   string `json:"previewText"`
	Title         string `json:"title"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		CoverMediaURL respjson.Field
		PreviewText   respjson.Field
		Title         respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SearchTweetArticle) RawJSON() string { return r.JSON.raw }
func (r *SearchTweetArticle) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Public card metadata attached to a tweet.
type SearchTweetCard struct {
	ID            string         `json:"id"`
	BindingValues map[string]any `json:"bindingValues"`
	Name          string         `json:"name"`
	URL           string         `json:"url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		BindingValues respjson.Field
		Name          respjson.Field
		URL           respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SearchTweetCard) RawJSON() string { return r.JSON.raw }
func (r *SearchTweetCard) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Community Note presentation metadata returned by X.
type SearchTweetCommunityNote struct {
	ID             string `json:"id"`
	DestinationURL string `json:"destinationUrl"`
	Footer         string `json:"footer"`
	ShortTitle     string `json:"shortTitle"`
	Subtitle       string `json:"subtitle"`
	Title          string `json:"title"`
	VisualStyle    string `json:"visualStyle"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		DestinationURL respjson.Field
		Footer         respjson.Field
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

// Edit history metadata returned by X.
type SearchTweetEdit struct {
	EditableUntilMsecs string   `json:"editableUntilMsecs"`
	EditTweetIDs       []string `json:"editTweetIds"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EditableUntilMsecs respjson.Field
		EditTweetIDs       respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SearchTweetEdit) RawJSON() string { return r.JSON.raw }
func (r *SearchTweetEdit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Complete Note Tweet content and rich-text metadata.
type SearchTweetNoteTweet struct {
	Text         string                            `json:"text" api:"required"`
	ID           string                            `json:"id"`
	Entities     map[string]any                    `json:"entities"`
	IsExpandable bool                              `json:"isExpandable"`
	RichtextTags []SearchTweetNoteTweetRichtextTag `json:"richtextTags"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Text         respjson.Field
		ID           respjson.Field
		Entities     respjson.Field
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

// Public place metadata attached to a tweet.
type SearchTweetPlace struct {
	ID          string         `json:"id"`
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

// Normalized media attached to a tweet.
type TweetMedia struct {
	// Media preview URL
	MediaURL string `json:"mediaUrl" api:"required"`
	// Any of "photo", "video", "animated_gif".
	Type TweetMediaType `json:"type" api:"required"`
	// X media link from the tweet
	URL string `json:"url" api:"required"`
	// X media entity ID.
	ID string `json:"id"`
	// Whether X permits direct media download.
	AllowDownload bool `json:"allowDownload"`
	// Accessibility text supplied for the media.
	AltText string `json:"altText"`
	// Video aspect ratio as width and height.
	AspectRatio []int64 `json:"aspectRatio"`
	// Media availability state reported by X.
	AvailabilityStatus string `json:"availabilityStatus"`
	// Display-friendly media URL reported by X.
	DisplayURL string `json:"displayUrl"`
	// Video duration in milliseconds.
	DurationMillis int64 `json:"durationMillis"`
	// Expanded X media URL.
	ExpandedURL string `json:"expandedUrl"`
	// Face-aware crop rectangles grouped by media size.
	FaceRects map[string][]TweetMediaFaceRect `json:"faceRects"`
	// Suggested image crops reported by X.
	FocusRects []TweetMediaFocusRect `json:"focusRects"`
	// Original media height.
	Height int64 `json:"height"`
	// Media entity offsets in the tweet text.
	Indices []int64 `json:"indices"`
	// Stable X media key.
	MediaKey string `json:"mediaKey"`
	// Whether X reports the media as monetizable.
	Monetizable bool `json:"monetizable"`
	// Named media renditions and resize modes.
	Sizes map[string]TweetMediaSize `json:"sizes"`
	// Available video encodings, ordered as returned
	VideoVariants []TweetMediaVideoVariant `json:"videoVariants"`
	// Original media width.
	Width int64 `json:"width"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MediaURL           respjson.Field
		Type               respjson.Field
		URL                respjson.Field
		ID                 respjson.Field
		AllowDownload      respjson.Field
		AltText            respjson.Field
		AspectRatio        respjson.Field
		AvailabilityStatus respjson.Field
		DisplayURL         respjson.Field
		DurationMillis     respjson.Field
		ExpandedURL        respjson.Field
		FaceRects          respjson.Field
		FocusRects         respjson.Field
		Height             respjson.Field
		Indices            respjson.Field
		MediaKey           respjson.Field
		Monetizable        respjson.Field
		Sizes              respjson.Field
		VideoVariants      respjson.Field
		Width              respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
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
	// Organization affiliation label shown on an X profile.
	AffiliatesHighlightedLabel     UserProfileAffiliatesHighlightedLabel `json:"affiliatesHighlightedLabel"`
	AutomatedBy                    string                                `json:"automatedBy"`
	BusinessAccountAffiliatesCount int64                                 `json:"businessAccountAffiliatesCount"`
	// Community role when returned by community member reads
	CommunityRole                   string `json:"communityRole"`
	CoverPicture                    string `json:"coverPicture"`
	CreatedAt                       string `json:"createdAt"`
	CreatorSubscriptionsCount       int64  `json:"creatorSubscriptionsCount"`
	Description                     string `json:"description"`
	FavouritesCount                 int64  `json:"favouritesCount"`
	Followers                       int64  `json:"followers"`
	Following                       int64  `json:"following"`
	HasCustomTimelines              bool   `json:"hasCustomTimelines"`
	HasGraduatedAccess              bool   `json:"hasGraduatedAccess"`
	HasHiddenSubscriptionsOnProfile bool   `json:"hasHiddenSubscriptionsOnProfile"`
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
	IsVerified               bool     `json:"isVerified"`
	Location                 string   `json:"location"`
	MediaCount               int64    `json:"mediaCount"`
	ParodyCommentaryFanLabel string   `json:"parodyCommentaryFanLabel"`
	PinnedTweetIDs           []string `json:"pinnedTweetIds"`
	PossiblySensitive        bool     `json:"possiblySensitive"`
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
	Protected           bool     `json:"protected"`
	StatusesCount       int64    `json:"statusesCount"`
	SuperFollowEligible bool     `json:"superFollowEligible"`
	Unavailable         bool     `json:"unavailable"`
	UnavailableReason   string   `json:"unavailableReason"`
	URL                 string   `json:"url"`
	Verified            bool     `json:"verified"`
	VerifiedType        string   `json:"verifiedType"`
	WithheldInCountries []string `json:"withheldInCountries"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                              respjson.Field
		Name                            respjson.Field
		Username                        respjson.Field
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
		Unavailable                     respjson.Field
		UnavailableReason               respjson.Field
		URL                             respjson.Field
		Verified                        respjson.Field
		VerifiedType                    respjson.Field
		WithheldInCountries             respjson.Field
		ExtraFields                     map[string]respjson.Field
		raw                             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserProfile) RawJSON() string { return r.JSON.raw }
func (r *UserProfile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Organization affiliation label shown on an X profile.
type UserProfileAffiliatesHighlightedLabel struct {
	BadgeURL             string `json:"badgeUrl"`
	Description          string `json:"description"`
	URL                  string `json:"url"`
	URLType              string `json:"urlType"`
	UserLabelDisplayType string `json:"userLabelDisplayType"`
	UserLabelType        string `json:"userLabelType"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BadgeURL             respjson.Field
		Description          respjson.Field
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
