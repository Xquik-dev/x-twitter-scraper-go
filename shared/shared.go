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
	// Whether the disclosure can be edited on X.
	CanEdit bool `json:"canEdit"`
	// Source of the AI-generated media disclosure.
	DetectionSource string `json:"detectionSource"`
	// True when X labels the tweet as containing AI-generated media.
	HasAIGeneratedMedia bool `json:"hasAiGeneratedMedia"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CanEdit             respjson.Field
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
	// X user profile with bio, follower counts, and verification status.
	Author UserProfile `json:"author"`
	// Content disclosure metadata shown by X when a tweet is labeled as paid
	// partnership content or AI-generated media.
	ContentDisclosure ContentDisclosure `json:"contentDisclosure"`
	ConversationID    string            `json:"conversationId"`
	CreatedAt         string            `json:"createdAt"`
	DisplayTextRange  []int64           `json:"displayTextRange"`
	Entities          map[string]any    `json:"entities"`
	InReplyToID       string            `json:"inReplyToId"`
	InReplyToUserID   string            `json:"inReplyToUserId"`
	InReplyToUsername string            `json:"inReplyToUsername"`
	IsLimitedReply    bool              `json:"isLimitedReply"`
	IsNoteTweet       bool              `json:"isNoteTweet"`
	IsQuoteStatus     bool              `json:"isQuoteStatus"`
	IsReply           bool              `json:"isReply"`
	Lang              string            `json:"lang"`
	Media             []TweetMedia      `json:"media"`
	Source            string            `json:"source"`
	Type              string            `json:"type"`
	URL               string            `json:"url"`
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
		Source            respjson.Field
		Type              respjson.Field
		URL               respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddedTweet) RawJSON() string { return r.JSON.raw }
func (r *EmbeddedTweet) UnmarshalJSON(data []byte) error {
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

// Paginated tweet results. The item count can be lower than pageSize when the
// source returns fewer tweets, filters remove tweets, or remaining credits cover
// fewer results. Follow next_cursor while has_next_page is true. An empty page can
// still have has_next_page true after filtering. Zero affordable results returns
// 402 insufficient_credits.
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

// Paginated user profiles. The item count can be lower than pageSize when the
// source returns fewer profiles or remaining credits cover fewer results. Follow
// next_cursor while has_next_page is true. A relationship can naturally contain
// fewer profiles than requested. Zero affordable results returns 402
// insufficient_credits.
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
	// X user profile with bio, follower counts, and verification status.
	Author UserProfile `json:"author"`
	// Content disclosure metadata shown by X when a tweet is labeled as paid
	// partnership content or AI-generated media.
	ContentDisclosure ContentDisclosure `json:"contentDisclosure"`
	// Root tweet ID for the search result conversation
	ConversationID string `json:"conversationId"`
	CreatedAt      string `json:"createdAt"`
	// Start and end offsets for rendered tweet text
	DisplayTextRange []int64 `json:"displayTextRange"`
	// Parsed search-result entities including URLs, mentions, hashtags, and media
	// markers
	Entities map[string]any `json:"entities"`
	// Tweet ID being replied to
	InReplyToID string `json:"inReplyToId"`
	// User ID being replied to
	InReplyToUserID string `json:"inReplyToUserId"`
	// Username being replied to
	InReplyToUsername string `json:"inReplyToUsername"`
	// Whether the tweet has limited reply permissions
	IsLimitedReply bool `json:"isLimitedReply"`
	// True for Note Tweets (long-form content, up to 25,000 characters)
	IsNoteTweet bool `json:"isNoteTweet"`
	// True when this search result quotes another tweet
	IsQuoteStatus bool `json:"isQuoteStatus"`
	// True when this search result is a reply
	IsReply bool `json:"isReply"`
	// Tweet language code
	Lang string `json:"lang"`
	// Search-result media attachments, omitted when no media is present
	Media []TweetMedia `json:"media"`
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
func (r SearchTweet) RawJSON() string { return r.JSON.raw }
func (r *SearchTweet) UnmarshalJSON(data []byte) error {
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
	// Available video encodings, ordered as returned
	VideoVariants []TweetMediaVideoVariant `json:"videoVariants"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MediaURL      respjson.Field
		Type          respjson.Field
		URL           respjson.Field
		VideoVariants respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
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
	ID          string `json:"id" api:"required"`
	Name        string `json:"name" api:"required"`
	Username    string `json:"username" api:"required"`
	AutomatedBy string `json:"automatedBy"`
	CanDm       bool   `json:"canDm"`
	// Community role when returned by community member reads
	CommunityRole      string `json:"communityRole"`
	CoverPicture       string `json:"coverPicture"`
	CreatedAt          string `json:"createdAt"`
	Description        string `json:"description"`
	FavouritesCount    int64  `json:"favouritesCount"`
	Followers          int64  `json:"followers"`
	Following          int64  `json:"following"`
	HasCustomTimelines bool   `json:"hasCustomTimelines"`
	IsAutomated        bool   `json:"isAutomated"`
	// Whether X shows a blue verification badge
	IsBlueVerified bool `json:"isBlueVerified"`
	IsTranslator   bool `json:"isTranslator"`
	// Whether X marks the profile as verified
	IsVerified        bool     `json:"isVerified"`
	Location          string   `json:"location"`
	MediaCount        int64    `json:"mediaCount"`
	PinnedTweetIDs    []string `json:"pinnedTweetIds"`
	PossiblySensitive bool     `json:"possiblySensitive"`
	// Structured profile bio with entity annotations
	ProfileBio map[string]any `json:"profile_bio"`
	// Original X profile banner field when available
	ProfileBannerURL string `json:"profileBannerUrl"`
	ProfilePicture   string `json:"profilePicture"`
	// Whether the profile protects its posts
	Protected         bool   `json:"protected"`
	StatusesCount     int64  `json:"statusesCount"`
	Unavailable       bool   `json:"unavailable"`
	UnavailableReason string `json:"unavailableReason"`
	URL               string `json:"url"`
	Verified          bool   `json:"verified"`
	VerifiedType      string `json:"verifiedType"`
	// Whether this profile follows the authenticated viewer
	ViewerFollowedBy bool `json:"viewerFollowedBy"`
	// Whether the authenticated viewer follows this profile
	ViewerFollowing     bool     `json:"viewerFollowing"`
	WithheldInCountries []string `json:"withheldInCountries"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                  respjson.Field
		Name                respjson.Field
		Username            respjson.Field
		AutomatedBy         respjson.Field
		CanDm               respjson.Field
		CommunityRole       respjson.Field
		CoverPicture        respjson.Field
		CreatedAt           respjson.Field
		Description         respjson.Field
		FavouritesCount     respjson.Field
		Followers           respjson.Field
		Following           respjson.Field
		HasCustomTimelines  respjson.Field
		IsAutomated         respjson.Field
		IsBlueVerified      respjson.Field
		IsTranslator        respjson.Field
		IsVerified          respjson.Field
		Location            respjson.Field
		MediaCount          respjson.Field
		PinnedTweetIDs      respjson.Field
		PossiblySensitive   respjson.Field
		ProfileBio          respjson.Field
		ProfileBannerURL    respjson.Field
		ProfilePicture      respjson.Field
		Protected           respjson.Field
		StatusesCount       respjson.Field
		Unavailable         respjson.Field
		UnavailableReason   respjson.Field
		URL                 respjson.Field
		Verified            respjson.Field
		VerifiedType        respjson.Field
		ViewerFollowedBy    respjson.Field
		ViewerFollowing     respjson.Field
		WithheldInCountries respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserProfile) RawJSON() string { return r.JSON.raw }
func (r *UserProfile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
