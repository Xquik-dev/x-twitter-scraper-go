// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

package shared

import (
	"encoding/json"
	"time"

	"github.com/Xquik-dev/x-twitter-scraper-go/internal/apijson"
	"github.com/Xquik-dev/x-twitter-scraper-go/packages/respjson"
)

type anyMonitorEventData interface {
	implMonitorEventDataUnion()
}

func (MonitorPost) implMonitorEventDataUnion()             {}
func (MonitorProfileChange) implMonitorEventDataUnion()    {}
func (MonitorFollow) implMonitorEventDataUnion()           {}
func (MonitorPostMetricChange) implMonitorEventDataUnion() {}
func (MonitorPostEdit) implMonitorEventDataUnion()         {}
func (MonitorPostDeleted) implMonitorEventDataUnion()      {}
func (MonitorLegacyPost) implMonitorEventDataUnion()       {}

// AsAny returns the variant the event's type picks, or nil for a type it does not
// know. Every post type picks [MonitorPost]. Read an older stored post with
// [MonitorEventDataUnion.AsMonitorLegacyPost]. Use the following switch statement
// to find the correct variant.
//
//	variant, err := event.Data.AsAny(event.Type)
//	switch variant := variant.(type) {
//	case shared.MonitorPost:
//	case shared.MonitorProfileChange:
//	case shared.MonitorFollow:
//	case shared.MonitorPostMetricChange:
//	case shared.MonitorPostEdit:
//	case shared.MonitorPostDeleted:
//	case shared.MonitorLegacyPost:
//	default:
//	}
func (u MonitorEventDataUnion) AsAny(eventType EventType) (anyMonitorEventData, error) {
	switch eventType {
	case "tweet.new", "tweet.reply", "tweet.retweet", "tweet.quote", "tweet.media",
		"tweet.link", "tweet.poll", "tweet.mention", "tweet.hashtag", "tweet.longform",
		"tweet.deleted", "broadcast.started", "space.started", "post.reply.new",
		"post.quote.new":
		return u.AsMonitorPost()
	case "profile.avatar.changed", "profile.banner.changed", "profile.name.changed",
		"profile.username.changed", "profile.bio.changed", "profile.location.changed",
		"profile.url.changed", "profile.verified.changed", "profile.protected.changed",
		"profile.pinned_tweet.changed", "profile.unavailable.changed",
		"profile.followers.changed", "profile.following.changed":
		return u.AsMonitorProfileChange()
	case "following.added", "following.removed", "follower.added", "follower.removed":
		return u.AsMonitorFollow()
	case "post.likes.changed", "post.retweets.changed", "post.replies.changed",
		"post.quotes.changed", "post.bookmarks.changed", "post.views.changed":
		return u.AsMonitorPostMetricChange()
	case "post.edited":
		return u.AsMonitorPostEdit()
	case "post.deleted":
		return u.AsMonitorPostDeleted()
	}
	return nil, nil
}

// AsMonitorPost decodes the union as a [MonitorPost].
func (u MonitorEventDataUnion) AsMonitorPost() (v MonitorPost, err error) {
	err = apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v, &v.JSON.raw)
	return
}

// AsMonitorProfileChange decodes the union as a [MonitorProfileChange].
func (u MonitorEventDataUnion) AsMonitorProfileChange() (v MonitorProfileChange, err error) {
	err = apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v, &v.JSON.raw)
	return
}

// AsMonitorFollow decodes the union as a [MonitorFollow].
func (u MonitorEventDataUnion) AsMonitorFollow() (v MonitorFollow, err error) {
	err = apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v, &v.JSON.raw)
	return
}

// AsMonitorPostMetricChange decodes the union as a [MonitorPostMetricChange].
func (u MonitorEventDataUnion) AsMonitorPostMetricChange() (v MonitorPostMetricChange, err error) {
	err = apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v, &v.JSON.raw)
	return
}

// AsMonitorPostEdit decodes the union as a [MonitorPostEdit].
func (u MonitorEventDataUnion) AsMonitorPostEdit() (v MonitorPostEdit, err error) {
	err = apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v, &v.JSON.raw)
	return
}

// AsMonitorPostDeleted decodes the union as a [MonitorPostDeleted].
func (u MonitorEventDataUnion) AsMonitorPostDeleted() (v MonitorPostDeleted, err error) {
	err = apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v, &v.JSON.raw)
	return
}

// AsMonitorLegacyPost decodes the union as a [MonitorLegacyPost].
func (u MonitorEventDataUnion) AsMonitorLegacyPost() (v MonitorLegacyPost, err error) {
	err = apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v, &v.JSON.raw)
	return
}

// MonitorProfileValueUnion contains all possible properties and values from
// [string], [bool], [int64], [[]string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString, OfBool, OfInt or OfStringArray.
type MonitorProfileValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [int64] instead of an object.
	OfInt int64 `json:",inline"`
	// This field will be present if the value is a [[]string] instead of an object.
	OfStringArray []string `json:",inline"`
	JSON          struct {
		OfString      respjson.Field
		OfBool        respjson.Field
		OfInt         respjson.Field
		OfStringArray respjson.Field
		raw           string
	} `json:"-"`
}

// AsString decodes the union as a string.
func (u MonitorProfileValueUnion) AsString() (v string, err error) {
	err = apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v, nil)
	return
}

// AsBool decodes the union as a bool.
func (u MonitorProfileValueUnion) AsBool() (v bool, err error) {
	err = apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v, nil)
	return
}

// AsInt decodes the union as an int64.
func (u MonitorProfileValueUnion) AsInt() (v int64, err error) {
	err = apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v, nil)
	return
}

// AsStringArray decodes the union as a []string.
func (u MonitorProfileValueUnion) AsStringArray() (v []string, err error) {
	err = apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v, nil)
	return
}

// RawJSON returns the unmodified JSON received from the API.
func (u MonitorProfileValueUnion) RawJSON() string { return u.JSON.raw }

// UnmarshalJSON decodes the union and keeps its raw JSON.
func (u *MonitorProfileValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u, &u.JSON.raw)
}

// MonitorFollowReason is why the account left the list, on following.removed and
// follower.removed.
type MonitorFollowReason string

// The reasons an account leaves a following or follower list.
const (
	MonitorFollowReasonUnavailable MonitorFollowReason = "unavailable"
	MonitorFollowReasonUnfollowed  MonitorFollowReason = "unfollowed"
)

// MonitorPostMetricChangeMetric is the post metric that changed.
type MonitorPostMetricChangeMetric string

// The post metrics a post monitor reports changes for.
const (
	MonitorPostMetricChangeMetricBookmarks MonitorPostMetricChangeMetric = "bookmarks"
	MonitorPostMetricChangeMetricLikes     MonitorPostMetricChangeMetric = "likes"
	MonitorPostMetricChangeMetricQuotes    MonitorPostMetricChangeMetric = "quotes"
	MonitorPostMetricChangeMetricReplies   MonitorPostMetricChangeMetric = "replies"
	MonitorPostMetricChangeMetricRetweets  MonitorPostMetricChangeMetric = "retweets"
	MonitorPostMetricChangeMetricViews     MonitorPostMetricChangeMetric = "views"
)

// MonitorEventDataUnion contains all possible properties and values from
// [MonitorPost], [MonitorProfileChange], [MonitorFollow],
// [MonitorPostMetricChange], [MonitorPostEdit], [MonitorPostDeleted],
// [MonitorLegacyPost]. Use [MonitorEventDataUnion.AsAny] with the event's type
// to switch on the variant, or the methods beginning with 'As' to cast it to
// one variant.
type MonitorEventDataUnion struct {
	ID string `json:"id"`
	// This field is from variant [MonitorPost].
	BookmarkCount int64 `json:"bookmarkCount"`
	// This field is from variant [MonitorPost].
	LikeCount int64 `json:"likeCount"`
	// This field is from variant [MonitorPost].
	QuoteCount int64 `json:"quoteCount"`
	// This field is from variant [MonitorPost].
	ReplyCount int64 `json:"replyCount"`
	// This field is from variant [MonitorPost].
	RetweetCount int64  `json:"retweetCount"`
	Text         string `json:"text"`
	// This field is from variant [MonitorPost].
	ViewCount int64 `json:"viewCount"`
	// Describes an X Article preview and its lifecycle metadata.
	// This field is from variant [MonitorPost].
	Article SearchTweetArticle `json:"article"`
	// Public profile with retained legacy monitor fields.
	Author MonitorProfile `json:"author"`
	// Why the tweet has no author. X serves a suspended or deleted account as an
	// unavailable user with a reason such as Suspended or NotFound. Omitted when the
	// author is available.
	// This field is from variant [MonitorPost].
	AuthorUnavailable SearchTweetAuthorUnavailable `json:"authorUnavailable"`
	// Describes a public card and its referenced profiles.
	// This field is from variant [MonitorPost].
	Card MonitorPostCard `json:"card"`
	// Price charts the author attached under the post, 1 per cashtag. Absent when the
	// post has none. X may send a cashtag without a chart window.
	// This field is from variant [MonitorPost].
	CashtagAttachments []SearchTweetCashtagAttachment `json:"cashtagAttachments"`
	// Community ID.
	// This field is from variant [MonitorPost].
	CommunityID string `json:"communityId"`
	// Community name. Present when X sends it with the post.
	// This field is from variant [MonitorPost].
	CommunityName string `json:"communityName"`
	// Community Note presentation metadata returned by X.
	// This field is from variant [MonitorPost].
	CommunityNote SearchTweetCommunityNote `json:"communityNote"`
	// Content disclosure metadata shown by X when a tweet is labeled as paid
	// partnership content or AI-generated media.
	// This field is from variant [MonitorPost].
	ContentDisclosure ContentDisclosure `json:"contentDisclosure"`
	// Public reply policy and conversation owner.
	// This field is from variant [MonitorPost].
	ConversationControl SearchTweetConversationControl `json:"conversationControl"`
	// This field is from variant [MonitorPost].
	ConversationID string `json:"conversationId"`
	// Author-published point, with longitude before latitude.
	// This field is from variant [MonitorPost].
	Coordinates SearchTweetCoordinates `json:"coordinates"`
	// This field is from variant [MonitorPost].
	CreatedAt string `json:"createdAt"`
	// This field is from variant [MonitorPost].
	DisplayTextRange []int64 `json:"displayTextRange"`
	// Lists edit-chain identifiers and the remaining edit window.
	// This field is from variant [MonitorPost].
	Edit SearchTweetEdit `json:"edit"`
	// Lists hashtags, symbols, links, and mentions from tweet text.
	// This field is from variant [MonitorPost].
	Entities SearchTweetEntities `json:"entities"`
	// Marks a subscriber-only post with the creator whose subscription unlocks it.
	// Non-subscribers get the preview text & the subscribe link. Omitted on public
	// posts.
	// This field is from variant [MonitorPost].
	ExclusiveContent SearchTweetExclusiveContent `json:"exclusiveContent"`
	// Author-published point, with latitude before longitude.
	// This field is from variant [MonitorPost].
	Geo SearchTweetGeo `json:"geo"`
	// Public metadata whose fields are defined by X.
	// This field is from variant [MonitorPost].
	GrokShareAttachment PublicMetadata `json:"grokShareAttachment"`
	// Grok's translation of the post, as X shows it.
	// This field is from variant [MonitorPost].
	GrokTranslatedPost PublicMetadata `json:"grokTranslatedPost"`
	// Whether X holds Community Notes on the post, shown or not.
	// This field is from variant [MonitorPost].
	HasCommunityNotes bool `json:"hasCommunityNotes"`
	// This field is from variant [MonitorPost].
	InReplyToID string `json:"inReplyToId"`
	// This field is from variant [MonitorPost].
	InReplyToUserID string `json:"inReplyToUserId"`
	// This field is from variant [MonitorPost].
	InReplyToUsername string `json:"inReplyToUsername"`
	// The notice x.com covers the post with until the reader views it, such as X's
	// public interest notice on a post that broke its rules but stays up.
	// This field is from variant [MonitorPost].
	Interstitial SearchTweetInterstitial `json:"interstitial"`
	// This field is from variant [MonitorPost].
	IsLimitedReply bool `json:"isLimitedReply"`
	// This field is from variant [MonitorPost].
	IsNoteTweet bool `json:"isNoteTweet"`
	// This field is from variant [MonitorPost].
	IsPinned bool `json:"isPinned"`
	// This field is from variant [MonitorPost].
	IsQuoteStatus bool `json:"isQuoteStatus"`
	// This field is from variant [MonitorPost].
	IsReply bool `json:"isReply"`
	// This field is from variant [MonitorPost].
	IsRetweet bool `json:"isRetweet"`
	// This field is from variant [MonitorPost].
	IsTranslatable bool `json:"isTranslatable"`
	// Public metadata whose fields are defined by X.
	// This field is from variant [MonitorPost].
	JetfuelAttachment PublicMetadata `json:"jetfuelAttachment"`
	// This field is from variant [MonitorPost].
	Lang string `json:"lang"`
	// Public interaction restrictions and user-facing prompts.
	// This field is from variant [MonitorPost].
	LimitedActions []SearchTweetLimitedAction `json:"limitedActions"`
	// Attached media items, omitted when unavailable.
	// This field is from variant [MonitorPost].
	Media []TweetMedia `json:"media"`
	// Complete Note Tweet content and rich-text metadata.
	// This field is from variant [MonitorPost].
	NoteTweet SearchTweetNoteTweet `json:"noteTweet"`
	// A parent, quoted or retweeted tweet. It has the TweetPublicData fields and an
	// author. Nesting stops after 4 levels.
	// This field is from variant [MonitorPost].
	ParentTweet EmbeddedTweet `json:"parentTweet"`
	// Describes public place metadata on a geotagged tweet.
	// This field is from variant [MonitorPost].
	Place SearchTweetPlace `json:"place"`
	// This field is from variant [MonitorPost].
	PossiblySensitive bool `json:"possiblySensitive"`
	// Public metadata whose fields are defined by X.
	// This field is from variant [MonitorPost].
	PostCta PublicMetadata `json:"postCta"`
	// Engagement counts retained from a prior tweet edit.
	// This field is from variant [MonitorPost].
	PreviousCounts SearchTweetPreviousCounts `json:"previousCounts"`
	// Embedded post with public fields and legacy profile aliases.
	QuotedTweet MonitorEmbeddedPost `json:"quoted_tweet"`
	// Quoted tweet ID.
	// This field is from variant [MonitorPost].
	QuotedTweetID string `json:"quotedTweetId"`
	// The quoted post's link, kept when X withholds that post.
	// This field is from variant [MonitorPost].
	QuotedTweetPermalink SearchTweetQuotedTweetPermalink `json:"quotedTweetPermalink"`
	// Why X shows no quoted post. Omitted when X names no reason, as its legacy
	// answers never do.
	// This field is from variant [MonitorPost].
	QuotedTweetUnavailable SearchTweetQuotedTweetUnavailable `json:"quotedTweetUnavailable"`
	// Public post and user referenced by this reaction.
	// This field is from variant [MonitorPost].
	ReactionContext MonitorPostReactionContext `json:"reactionContext"`
	// Embedded post with public fields and legacy profile aliases.
	RetweetedTweet MonitorEmbeddedPost `json:"retweeted_tweet"`
	// Repost event time in UTC. Null when unavailable; omitted for other posts.
	// This field is from variant [MonitorPost].
	RetweetedAt time.Time `json:"retweetedAt" api:"nullable" format:"date-time"`
	// Public metadata whose fields are defined by X.
	// This field is from variant [MonitorPost].
	Scopes PublicMetadata `json:"scopes"`
	// The label X shows under a post whose reach it limits. It links to the page
	// that explains it.
	// This field is from variant [MonitorPost].
	SoftIntervention SoftIntervention `json:"softIntervention"`
	// This field is from variant [MonitorPost].
	Source string `json:"source"`
	// Sports game context attached to the post, when available.
	// This field is from variant [MonitorPost].
	SportsContext SearchTweetSportsContext `json:"sportsContext"`
	// The Subscriber badge X shows on a reply. The reply's author subscribes to the
	// creator who started the conversation. Omitted on a post without the badge.
	// This field is from variant [MonitorPost].
	SubscriberBadge SearchTweetSubscriberBadge `json:"subscriberBadge"`
	// Where `text` matches the search, on each row a search read. Offsets count code
	// points of `text`, as `displayTextRange` does. Empty when X marks nothing in the
	// post. X marks matches only in a post's first 280 characters.
	// This field is from variant [MonitorPost].
	TextHighlights []SearchTweetTextHighlight `json:"textHighlights"`
	// Public visibility notice attached to an available tweet.
	// This field is from variant [MonitorPost].
	Tombstone SearchTweetTombstone `json:"tombstone"`
	// This field is from variant [MonitorPost].
	Type string `json:"type"`
	// The author's ID when X withholds the author's profile.
	// This field is from variant [MonitorPost].
	UnavailableAuthorID string `json:"unavailableAuthorId"`
	// X did not send these counts for this post, so treat each 0 as unknown.
	// This field is from variant [MonitorPost].
	UnknownCounts []TweetUnknownCount `json:"unknownCounts"`
	// User IDs that left this conversation.
	// This field is from variant [MonitorPost].
	UnmentionedUserIDs []string `json:"unmentionedUserIds"`
	// This field is from variant [MonitorPost].
	URL string `json:"url"`
	// This field is from variant [MonitorPost].
	ViewState string `json:"viewState"`
	// The fields X did not send for this post. text reads empty, so treat it as
	// unknown rather than empty. Left out when X sent the text.
	// This field is from variant [MonitorPost].
	UnknownFields []string `json:"unknownFields"`
	// Changed profile field, followers or following on the 2 count events.
	// This field is from variant [MonitorProfileChange].
	Field string `json:"field"`
	// Permanent X user ID.
	// This field is from variant [MonitorProfileChange].
	UserID string `json:"userId"`
	// Username observed for this change.
	// This field is from variant [MonitorProfileChange].
	Username string `json:"username"`
	// Current count minus previous count. Only on the 2 count events.
	Change int64 `json:"change"`
	// A profile value, or a post metric count in [MonitorProfileValueUnion.OfInt].
	Current MonitorProfileValueUnion `json:"current"`
	// A profile value, or a post metric count in [MonitorProfileValueUnion.OfInt].
	Previous MonitorProfileValueUnion `json:"previous"`
	// X's time of the follow, on following.added & follower.added. It is missing
	// when X gives none, as for an account that returns to a list.
	// This field is from variant [MonitorFollow].
	FollowedAt time.Time `json:"followedAt" format:"date-time"`
	// Whether followedAt is X's exact time of this follow.
	// This field is from variant [MonitorFollow].
	FollowedAtExact bool `json:"followedAtExact"`
	// Why the account left the list, on following.removed & follower.removed.
	//
	// Any of "unavailable", "unfollowed".
	// This field is from variant [MonitorFollow].
	Reason MonitorFollowReason `json:"reason"`
	// The metric that changed.
	//
	// Any of "bookmarks", "likes", "quotes", "replies", "retweets", "views".
	// This field is from variant [MonitorPostMetricChange].
	Metric MonitorPostMetricChangeMetric `json:"metric"`
	// The text before the edit.
	// This field is from variant [MonitorPostEdit].
	PreviousText string `json:"previousText"`
	// X's ID of the newest version of the post.
	// This field is from variant [MonitorPostEdit].
	VersionID string `json:"versionId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                     respjson.Field
		BookmarkCount          respjson.Field
		LikeCount              respjson.Field
		QuoteCount             respjson.Field
		ReplyCount             respjson.Field
		RetweetCount           respjson.Field
		Text                   respjson.Field
		ViewCount              respjson.Field
		Article                respjson.Field
		Author                 respjson.Field
		AuthorUnavailable      respjson.Field
		Card                   respjson.Field
		CashtagAttachments     respjson.Field
		CommunityID            respjson.Field
		CommunityName          respjson.Field
		CommunityNote          respjson.Field
		ContentDisclosure      respjson.Field
		ConversationControl    respjson.Field
		ConversationID         respjson.Field
		Coordinates            respjson.Field
		CreatedAt              respjson.Field
		DisplayTextRange       respjson.Field
		Edit                   respjson.Field
		Entities               respjson.Field
		ExclusiveContent       respjson.Field
		Geo                    respjson.Field
		GrokShareAttachment    respjson.Field
		GrokTranslatedPost     respjson.Field
		HasCommunityNotes      respjson.Field
		InReplyToID            respjson.Field
		InReplyToUserID        respjson.Field
		InReplyToUsername      respjson.Field
		Interstitial           respjson.Field
		IsLimitedReply         respjson.Field
		IsNoteTweet            respjson.Field
		IsPinned               respjson.Field
		IsQuoteStatus          respjson.Field
		IsReply                respjson.Field
		IsRetweet              respjson.Field
		IsTranslatable         respjson.Field
		JetfuelAttachment      respjson.Field
		Lang                   respjson.Field
		LimitedActions         respjson.Field
		Media                  respjson.Field
		NoteTweet              respjson.Field
		ParentTweet            respjson.Field
		Place                  respjson.Field
		PossiblySensitive      respjson.Field
		PostCta                respjson.Field
		PreviousCounts         respjson.Field
		QuotedTweet            respjson.Field
		QuotedTweetID          respjson.Field
		QuotedTweetPermalink   respjson.Field
		QuotedTweetUnavailable respjson.Field
		ReactionContext        respjson.Field
		RetweetedTweet         respjson.Field
		RetweetedAt            respjson.Field
		Scopes                 respjson.Field
		SoftIntervention       respjson.Field
		Source                 respjson.Field
		SportsContext          respjson.Field
		SubscriberBadge        respjson.Field
		TextHighlights         respjson.Field
		Tombstone              respjson.Field
		Type                   respjson.Field
		UnavailableAuthorID    respjson.Field
		UnknownCounts          respjson.Field
		UnmentionedUserIDs     respjson.Field
		URL                    respjson.Field
		ViewState              respjson.Field
		UnknownFields          respjson.Field
		Field                  respjson.Field
		UserID                 respjson.Field
		Username               respjson.Field
		Change                 respjson.Field
		Current                respjson.Field
		Previous               respjson.Field
		FollowedAt             respjson.Field
		FollowedAtExact        respjson.Field
		Reason                 respjson.Field
		Metric                 respjson.Field
		PreviousText           respjson.Field
		VersionID              respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// RawJSON returns the unmodified JSON received from the API.
func (u MonitorEventDataUnion) RawJSON() string { return u.JSON.raw }
func (u *MonitorEventDataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u, &u.JSON.raw)
}

// MonitorPost is a monitored post with public fields and retained legacy
// fields.
type MonitorPost struct {
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
	// Public profile with retained legacy monitor fields.
	Author MonitorProfile `json:"author"`
	// Why the tweet has no author. X serves a suspended or deleted account as an
	// unavailable user with a reason such as Suspended or NotFound. Omitted when the
	// author is available.
	AuthorUnavailable SearchTweetAuthorUnavailable `json:"authorUnavailable"`
	// Describes a public card and its referenced profiles.
	Card MonitorPostCard `json:"card"`
	// Price charts the author attached under the post, 1 per cashtag. Absent when the
	// post has none. X may send a cashtag without a chart window.
	CashtagAttachments []SearchTweetCashtagAttachment `json:"cashtagAttachments"`
	// Community ID.
	CommunityID string `json:"communityId"`
	// Community name. Present when X sends it with the post.
	CommunityName string `json:"communityName"`
	// Community Note presentation metadata returned by X.
	CommunityNote SearchTweetCommunityNote `json:"communityNote"`
	// Content disclosure metadata shown by X when a tweet is labeled as paid
	// partnership content or AI-generated media.
	ContentDisclosure ContentDisclosure `json:"contentDisclosure"`
	// Public reply policy and conversation owner.
	ConversationControl SearchTweetConversationControl `json:"conversationControl"`
	ConversationID      string                         `json:"conversationId"`
	// Author-published point, with longitude before latitude.
	Coordinates      SearchTweetCoordinates `json:"coordinates"`
	CreatedAt        string                 `json:"createdAt"`
	DisplayTextRange []int64                `json:"displayTextRange"`
	// Lists edit-chain identifiers and the remaining edit window.
	Edit SearchTweetEdit `json:"edit"`
	// Lists hashtags, symbols, links, and mentions from tweet text.
	Entities SearchTweetEntities `json:"entities"`
	// Marks a subscriber-only post with the creator whose subscription unlocks it.
	// Non-subscribers get the preview text & the subscribe link. Omitted on public
	// posts.
	ExclusiveContent SearchTweetExclusiveContent `json:"exclusiveContent"`
	// Author-published point, with latitude before longitude.
	Geo SearchTweetGeo `json:"geo"`
	// Public metadata whose fields are defined by X.
	GrokShareAttachment PublicMetadata `json:"grokShareAttachment"`
	// Grok's translation of the post, as X shows it.
	GrokTranslatedPost PublicMetadata `json:"grokTranslatedPost"`
	// Whether X holds Community Notes on the post, shown or not.
	HasCommunityNotes bool   `json:"hasCommunityNotes"`
	InReplyToID       string `json:"inReplyToId"`
	InReplyToUserID   string `json:"inReplyToUserId"`
	InReplyToUsername string `json:"inReplyToUsername"`
	// The notice x.com covers the post with until the reader views it, such as X's
	// public interest notice on a post that broke its rules but stays up.
	Interstitial   SearchTweetInterstitial `json:"interstitial"`
	IsLimitedReply bool                    `json:"isLimitedReply"`
	IsNoteTweet    bool                    `json:"isNoteTweet"`
	IsPinned       bool                    `json:"isPinned"`
	IsQuoteStatus  bool                    `json:"isQuoteStatus"`
	IsReply        bool                    `json:"isReply"`
	IsRetweet      bool                    `json:"isRetweet"`
	IsTranslatable bool                    `json:"isTranslatable"`
	// Public metadata whose fields are defined by X.
	JetfuelAttachment PublicMetadata `json:"jetfuelAttachment"`
	Lang              string         `json:"lang"`
	// Public interaction restrictions and user-facing prompts.
	LimitedActions []SearchTweetLimitedAction `json:"limitedActions"`
	// Attached media items, omitted when unavailable.
	Media []TweetMedia `json:"media"`
	// Complete Note Tweet content and rich-text metadata.
	NoteTweet SearchTweetNoteTweet `json:"noteTweet"`
	// A parent, quoted or retweeted tweet. It has the TweetPublicData fields and an
	// author. Nesting stops after 4 levels.
	ParentTweet EmbeddedTweet `json:"parentTweet"`
	// Describes public place metadata on a geotagged tweet.
	Place             SearchTweetPlace `json:"place"`
	PossiblySensitive bool             `json:"possiblySensitive"`
	// Public metadata whose fields are defined by X.
	PostCta PublicMetadata `json:"postCta"`
	// Engagement counts retained from a prior tweet edit.
	PreviousCounts SearchTweetPreviousCounts `json:"previousCounts"`
	// Embedded post with public fields and legacy profile aliases.
	QuotedTweet MonitorEmbeddedPost `json:"quoted_tweet"`
	// Quoted tweet ID.
	QuotedTweetID string `json:"quotedTweetId"`
	// The quoted post's link, kept when X withholds that post.
	QuotedTweetPermalink SearchTweetQuotedTweetPermalink `json:"quotedTweetPermalink"`
	// Why X shows no quoted post. Omitted when X names no reason, as its legacy
	// answers never do.
	QuotedTweetUnavailable SearchTweetQuotedTweetUnavailable `json:"quotedTweetUnavailable"`
	// Public post and user referenced by this reaction.
	ReactionContext MonitorPostReactionContext `json:"reactionContext"`
	// Embedded post with public fields and legacy profile aliases.
	RetweetedTweet MonitorEmbeddedPost `json:"retweeted_tweet"`
	// Repost event time in UTC. Null when unavailable; omitted for other posts.
	RetweetedAt time.Time `json:"retweetedAt" api:"nullable" format:"date-time"`
	// Public metadata whose fields are defined by X.
	Scopes PublicMetadata `json:"scopes"`
	// The label X shows under a post whose reach it limits. It links to the page
	// that explains it.
	SoftIntervention SoftIntervention `json:"softIntervention"`
	Source           string           `json:"source"`
	// Sports game context attached to the post, when available.
	SportsContext SearchTweetSportsContext `json:"sportsContext"`
	// The Subscriber badge X shows on a reply. The reply's author subscribes to the
	// creator who started the conversation. Omitted on a post without the badge.
	SubscriberBadge SearchTweetSubscriberBadge `json:"subscriberBadge"`
	// Where `text` matches the search, on each row a search read. Offsets count code
	// points of `text`, as `displayTextRange` does. Empty when X marks nothing in the
	// post. X marks matches only in a post's first 280 characters.
	TextHighlights []SearchTweetTextHighlight `json:"textHighlights"`
	// Public visibility notice attached to an available tweet.
	Tombstone SearchTweetTombstone `json:"tombstone"`
	Type      string               `json:"type"`
	// The author's ID when X withholds the author's profile.
	UnavailableAuthorID string `json:"unavailableAuthorId"`
	// X did not send these counts for this post, so treat each 0 as unknown.
	UnknownCounts []TweetUnknownCount `json:"unknownCounts"`
	// User IDs that left this conversation.
	UnmentionedUserIDs []string `json:"unmentionedUserIds"`
	URL                string   `json:"url"`
	ViewState          string   `json:"viewState"`
	// The fields X did not send for this post. text reads empty, so treat it as
	// unknown rather than empty. Left out when X sent the text.
	UnknownFields []string `json:"unknownFields"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                     respjson.Field
		BookmarkCount          respjson.Field
		LikeCount              respjson.Field
		QuoteCount             respjson.Field
		ReplyCount             respjson.Field
		RetweetCount           respjson.Field
		Text                   respjson.Field
		ViewCount              respjson.Field
		Article                respjson.Field
		Author                 respjson.Field
		AuthorUnavailable      respjson.Field
		Card                   respjson.Field
		CashtagAttachments     respjson.Field
		CommunityID            respjson.Field
		CommunityName          respjson.Field
		CommunityNote          respjson.Field
		ContentDisclosure      respjson.Field
		ConversationControl    respjson.Field
		ConversationID         respjson.Field
		Coordinates            respjson.Field
		CreatedAt              respjson.Field
		DisplayTextRange       respjson.Field
		Edit                   respjson.Field
		Entities               respjson.Field
		ExclusiveContent       respjson.Field
		Geo                    respjson.Field
		GrokShareAttachment    respjson.Field
		GrokTranslatedPost     respjson.Field
		HasCommunityNotes      respjson.Field
		InReplyToID            respjson.Field
		InReplyToUserID        respjson.Field
		InReplyToUsername      respjson.Field
		Interstitial           respjson.Field
		IsLimitedReply         respjson.Field
		IsNoteTweet            respjson.Field
		IsPinned               respjson.Field
		IsQuoteStatus          respjson.Field
		IsReply                respjson.Field
		IsRetweet              respjson.Field
		IsTranslatable         respjson.Field
		JetfuelAttachment      respjson.Field
		Lang                   respjson.Field
		LimitedActions         respjson.Field
		Media                  respjson.Field
		NoteTweet              respjson.Field
		ParentTweet            respjson.Field
		Place                  respjson.Field
		PossiblySensitive      respjson.Field
		PostCta                respjson.Field
		PreviousCounts         respjson.Field
		QuotedTweet            respjson.Field
		QuotedTweetID          respjson.Field
		QuotedTweetPermalink   respjson.Field
		QuotedTweetUnavailable respjson.Field
		ReactionContext        respjson.Field
		RetweetedTweet         respjson.Field
		RetweetedAt            respjson.Field
		Scopes                 respjson.Field
		SoftIntervention       respjson.Field
		Source                 respjson.Field
		SportsContext          respjson.Field
		SubscriberBadge        respjson.Field
		TextHighlights         respjson.Field
		Tombstone              respjson.Field
		Type                   respjson.Field
		UnavailableAuthorID    respjson.Field
		UnknownCounts          respjson.Field
		UnmentionedUserIDs     respjson.Field
		URL                    respjson.Field
		ViewState              respjson.Field
		UnknownFields          respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// RawJSON returns the unmodified JSON received from the API.
func (r MonitorPost) RawJSON() string { return r.JSON.raw }
func (r *MonitorPost) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// MonitorProfileChange is a changed profile field with its previous and current
// values. The field of a count event is followers or following, with both
// counts as integers.
type MonitorProfileChange struct {
	// Changed profile field, followers or following on the 2 count events.
	Field string `json:"field" api:"required"`
	// Permanent X user ID.
	UserID string `json:"userId" api:"required"`
	// Username observed for this change.
	Username string `json:"username" api:"required"`
	// Current count minus previous count. Only on the 2 count events.
	Change int64 `json:"change"`
	// Profile text, flag, count, or pinned post IDs. Missing values are omitted.
	Current MonitorProfileValueUnion `json:"current"`
	// Profile text, flag, count, or pinned post IDs. Missing values are omitted.
	Previous MonitorProfileValueUnion `json:"previous"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Field       respjson.Field
		UserID      respjson.Field
		Username    respjson.Field
		Change      respjson.Field
		Current     respjson.Field
		Previous    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// RawJSON returns the unmodified JSON received from the API.
func (r MonitorProfileChange) RawJSON() string { return r.JSON.raw }
func (r *MonitorProfileChange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// MonitorFollow is the other account of a following or follower event. It holds
// the ID alone when X fails the profile lookup.
type MonitorFollow struct {
	// The other account's X user ID.
	ID string `json:"id" api:"required"`
	// X's time of the follow, on following.added & follower.added. It is missing
	// when X gives none, as for an account that returns to a list.
	FollowedAt time.Time `json:"followedAt" format:"date-time"`
	// Whether followedAt is X's exact time of this follow.
	FollowedAtExact bool `json:"followedAtExact"`
	// Why the account left the list, on following.removed & follower.removed.
	//
	// Any of "unavailable", "unfollowed".
	Reason MonitorFollowReason `json:"reason"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		FollowedAt      respjson.Field
		FollowedAtExact respjson.Field
		Reason          respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// RawJSON returns the unmodified JSON received from the API.
func (r MonitorFollow) RawJSON() string { return r.JSON.raw }
func (r *MonitorFollow) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// MonitorPostMetricChange is a post metric's change since its last event, on
// the 6 post metric changed events.
type MonitorPostMetricChange struct {
	// Current count minus previous count.
	Change int64 `json:"change" api:"required"`
	// The count X shows now.
	Current int64 `json:"current" api:"required"`
	// The metric that changed.
	//
	// Any of "bookmarks", "likes", "quotes", "replies", "retweets", "views".
	Metric MonitorPostMetricChangeMetric `json:"metric" api:"required"`
	// The count the last event of this metric reported.
	Previous int64 `json:"previous" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Change      respjson.Field
		Current     respjson.Field
		Metric      respjson.Field
		Previous    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// RawJSON returns the unmodified JSON received from the API.
func (r MonitorPostMetricChange) RawJSON() string { return r.JSON.raw }
func (r *MonitorPostMetricChange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// MonitorPostEdit is the post's text before and after an edit, on post.edited.
type MonitorPostEdit struct {
	// The text before the edit.
	PreviousText string `json:"previousText" api:"required"`
	// The text of the newest version.
	Text string `json:"text" api:"required"`
	// X's ID of the newest version of the post.
	VersionID string `json:"versionId" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PreviousText respjson.Field
		Text         respjson.Field
		VersionID    respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// RawJSON returns the unmodified JSON received from the API.
func (r MonitorPostEdit) RawJSON() string { return r.JSON.raw }
func (r *MonitorPostEdit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// MonitorPostDeleted is the last text the monitor saw, on post.deleted.
type MonitorPostDeleted struct {
	// The post's text before its author deleted it.
	Text string `json:"text" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Text        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// RawJSON returns the unmodified JSON received from the API.
func (r MonitorPostDeleted) RawJSON() string { return r.JSON.raw }
func (r *MonitorPostDeleted) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// MonitorLegacyPost is an older stored post. Fields can be missing or null.
type MonitorLegacyPost struct {
	ID   string `json:"id"`
	Text string `json:"text"`
	// Legacy author fields, null or missing on some stored posts.
	Author PublicMetadata `json:"author" api:"nullable"`
	// Legacy quoted post fields, null or missing on some stored posts.
	QuotedTweet PublicMetadata `json:"quoted_tweet" api:"nullable"`
	// Legacy retweeted post fields, null or missing on some stored posts.
	RetweetedTweet PublicMetadata `json:"retweeted_tweet" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		Text           respjson.Field
		Author         respjson.Field
		QuotedTweet    respjson.Field
		RetweetedTweet respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// RawJSON returns the unmodified JSON received from the API.
func (r MonitorLegacyPost) RawJSON() string { return r.JSON.raw }
func (r *MonitorLegacyPost) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// MonitorProfile is a public profile with retained legacy monitor fields.
type MonitorProfile struct {
	ID       string `json:"id" api:"required"`
	Name     string `json:"name" api:"required"`
	Username string `json:"username" api:"required"`
	// X's best-effort public label inferred from aggregated account-access IP
	// addresses. It does not state nationality, residence, identity, registration,
	// post location, or exact location.
	AccountBasedIn UserProfileAccountBasedIn `json:"accountBasedIn" api:"nullable"`
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
	AffiliatesHighlightedLabel UserProfileAffiliatesHighlightedLabel `json:"affiliatesHighlightedLabel"`
	// Handle of the organization X shows the account is affiliated with, from its
	// About page.
	AffiliateUsername string `json:"affiliateUsername"`
	AutomatedBy       string `json:"automatedBy"`
	// The parts of the birthday the account made public.
	Birthdate                      UserProfileBirthdate `json:"birthdate"`
	BusinessAccountAffiliatesCount int64                `json:"businessAccountAffiliatesCount"`
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
	ContactInfo               UserProfileContactInfo `json:"contactInfo"`
	CoverPicture              string                 `json:"coverPicture"`
	CreatedAt                 string                 `json:"createdAt"`
	CreatorSubscriptionsCount int64                  `json:"creatorSubscriptionsCount"`
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
	Followers          int64 `json:"followers"`
	Following          int64 `json:"following"`
	// Whether the profile may add a location to posts.
	GeoEnabled bool `json:"geoEnabled"`
	// Public profile bio translation returned by X
	GrokTranslatedBio  PublicMetadata `json:"grokTranslatedBio"`
	HasCustomTimelines bool           `json:"hasCustomTimelines"`
	// Whether the profile filled X's extended fields, such as a birth date. True or
	// false.
	HasExtendedProfile              bool `json:"hasExtendedProfile"`
	HasGraduatedAccess              bool `json:"hasGraduatedAccess"`
	HasHiddenSubscriptionsOnProfile bool `json:"hasHiddenSubscriptionsOnProfile"`
	// Profile highlight availability and count metadata.
	HighlightsInfo UserProfileHighlightsInfo `json:"highlightsInfo"`
	// Identity verification metadata displayed by X.
	IdentityVerification UserProfileIdentityVerification `json:"identityVerification"`
	IsAutomated          bool                            `json:"isAutomated"`
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
	Professional PublicMetadata `json:"professional"`
	// Structured profile bio with entity annotations
	ProfileBio PublicMetadata `json:"profile_bio"`
	// Original X profile banner field when available
	ProfileBannerURL           string `json:"profileBannerUrl"`
	ProfileDescriptionLanguage string `json:"profileDescriptionLanguage"`
	ProfileImageShape          string `json:"profileImageShape"`
	ProfileInterstitialType    string `json:"profileInterstitialType"`
	ProfilePicture             string `json:"profilePicture"`
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
	// has no match or lookup failed; otherwise omitted.
	RetweetedAt   time.Time `json:"retweetedAt" api:"nullable" format:"date-time"`
	StatusesCount int64     `json:"statusesCount"`
	// Accounts subscribed to the profile's paid Subscriptions. 0 when it has none.
	SubscribersCount    int64 `json:"subscribersCount"`
	SuperFollowEligible bool  `json:"superFollowEligible"`
	// Whether X marks the subscription profile as active.
	SuperFollowsUserProfileActive bool `json:"superFollowsUserProfileActive"`
	// Public payment and creator-support handles shown on X.
	TipJar            UserProfileTipJar `json:"tipJar"`
	Unavailable       bool              `json:"unavailable"`
	UnavailableReason string            `json:"unavailableReason"`
	URL               string            `json:"url"`
	// How often the account changed its username.
	UsernameChanges UserProfileUsernameChanges `json:"usernameChanges"`
	// X's text on why the profile is verified, with its links.
	VerificationReason  UserProfileVerificationReason `json:"verificationReason"`
	Verified            bool                          `json:"verified"`
	VerifiedType        string                        `json:"verifiedType"`
	WithheldInCountries []string                      `json:"withheldInCountries"`
	// Whether X withholds a post or user
	WithheldScope string `json:"withheldScope"`
	// Deprecated: Legacy username alias. Read username instead.
	UserName string `json:"userName"`
	// Deprecated: Original X flag. Read verified for the displayed badge.
	RawVerified bool `json:"rawVerified"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                              respjson.Field
		Name                            respjson.Field
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
		VerifiedType                    respjson.Field
		WithheldInCountries             respjson.Field
		WithheldScope                   respjson.Field
		UserName                        respjson.Field
		RawVerified                     respjson.Field
		ExtraFields                     map[string]respjson.Field
		raw                             string
	} `json:"-"`
}

// RawJSON returns the unmodified JSON received from the API.
func (r MonitorProfile) RawJSON() string { return r.JSON.raw }
func (r *MonitorProfile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// MonitorEmbeddedPost is an embedded post with public fields and legacy profile
// aliases.
type MonitorEmbeddedPost struct {
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
	// Public profile with retained legacy monitor fields.
	Author            MonitorProfile               `json:"author"`
	AuthorUnavailable SearchTweetAuthorUnavailable `json:"authorUnavailable"`
	// Describes a public card and its referenced profiles.
	Card               EmbeddedTweetCard              `json:"card"`
	CashtagAttachments []SearchTweetCashtagAttachment `json:"cashtagAttachments"`
	// Community ID.
	CommunityID   string `json:"communityId"`
	CommunityName string `json:"communityName"`
	// Community Note presentation metadata returned by X.
	CommunityNote EmbeddedTweetCommunityNote `json:"communityNote"`
	// Content disclosure metadata shown by X when a tweet is labeled as paid
	// partnership content or AI-generated media.
	ContentDisclosure ContentDisclosure `json:"contentDisclosure"`
	// Public reply policy and conversation owner.
	ConversationControl EmbeddedTweetConversationControl `json:"conversationControl"`
	ConversationID      string                           `json:"conversationId"`
	Coordinates         SearchTweetCoordinates           `json:"coordinates"`
	CreatedAt           string                           `json:"createdAt"`
	DisplayTextRange    []int64                          `json:"displayTextRange"`
	// Lists edit-chain identifiers and the remaining edit window.
	Edit EmbeddedTweetEdit `json:"edit"`
	// Lists hashtags, symbols, links, and mentions from tweet text.
	Entities         EmbeddedTweetEntities       `json:"entities"`
	ExclusiveContent SearchTweetExclusiveContent `json:"exclusiveContent"`
	Geo              SearchTweetGeo              `json:"geo"`
	// Public metadata whose fields are defined by X.
	GrokShareAttachment PublicMetadata          `json:"grokShareAttachment"`
	GrokTranslatedPost  PublicMetadata          `json:"grokTranslatedPost"`
	HasCommunityNotes   bool                    `json:"hasCommunityNotes"`
	InReplyToID         string                  `json:"inReplyToId"`
	InReplyToUserID     string                  `json:"inReplyToUserId"`
	InReplyToUsername   string                  `json:"inReplyToUsername"`
	Interstitial        SearchTweetInterstitial `json:"interstitial"`
	IsLimitedReply      bool                    `json:"isLimitedReply"`
	IsNoteTweet         bool                    `json:"isNoteTweet"`
	IsPinned            bool                    `json:"isPinned"`
	IsQuoteStatus       bool                    `json:"isQuoteStatus"`
	IsReply             bool                    `json:"isReply"`
	IsRetweet           bool                    `json:"isRetweet"`
	IsTranslatable      bool                    `json:"isTranslatable"`
	// Public metadata whose fields are defined by X.
	JetfuelAttachment PublicMetadata `json:"jetfuelAttachment"`
	Lang              string         `json:"lang"`
	// Public interaction restrictions and user-facing prompts.
	LimitedActions []EmbeddedTweetLimitedAction `json:"limitedActions"`
	// Attached media items, omitted when unavailable.
	Media []TweetMedia `json:"media"`
	// Complete Note Tweet content and rich-text metadata.
	NoteTweet   EmbeddedTweetNoteTweet `json:"noteTweet"`
	ParentTweet *EmbeddedTweet         `json:"parentTweet"`
	// Describes public place metadata on a geotagged tweet.
	Place             EmbeddedTweetPlace `json:"place"`
	PossiblySensitive bool               `json:"possiblySensitive"`
	// Public metadata whose fields are defined by X.
	PostCta PublicMetadata `json:"postCta"`
	// Engagement counts retained from a prior tweet edit.
	PreviousCounts EmbeddedTweetPreviousCounts `json:"previousCounts"`
	// Nested tweet context at depth 2.
	QuotedTweet EmbeddedTweetQuotedTweet `json:"quoted_tweet"`
	// Quoted tweet ID.
	QuotedTweetID          string                            `json:"quotedTweetId"`
	QuotedTweetPermalink   SearchTweetQuotedTweetPermalink   `json:"quotedTweetPermalink"`
	QuotedTweetUnavailable SearchTweetQuotedTweetUnavailable `json:"quotedTweetUnavailable"`
	// Public post and user referenced by this reaction.
	ReactionContext EmbeddedTweetReactionContext `json:"reactionContext"`
	// Nested tweet context at depth 2.
	RetweetedTweet EmbeddedTweetRetweetedTweet `json:"retweeted_tweet"`
	// Repost event time in UTC. Null when unavailable; omitted for other posts.
	RetweetedAt time.Time `json:"retweetedAt" api:"nullable" format:"date-time"`
	// Public metadata whose fields are defined by X.
	Scopes PublicMetadata `json:"scopes"`
	// The label X shows under a post whose reach it limits. It links to the page
	// that explains it.
	SoftIntervention SoftIntervention `json:"softIntervention"`
	Source           string           `json:"source"`
	// Sports game context attached to the post, when available.
	SportsContext   EmbeddedTweetSportsContext `json:"sportsContext"`
	SubscriberBadge SearchTweetSubscriberBadge `json:"subscriberBadge"`
	TextHighlights  []SearchTweetTextHighlight `json:"textHighlights"`
	// Public visibility notice attached to an available tweet.
	Tombstone           EmbeddedTweetTombstone `json:"tombstone"`
	Type                string                 `json:"type"`
	UnavailableAuthorID string                 `json:"unavailableAuthorId"`
	// X did not send these counts for this post, so treat each 0 as unknown.
	UnknownCounts []TweetUnknownCount `json:"unknownCounts"`
	// User IDs that left this conversation.
	UnmentionedUserIDs []string `json:"unmentionedUserIds"`
	URL                string   `json:"url"`
	ViewState          string   `json:"viewState"`
	// The fields X did not send for this post. text reads empty, so treat it as
	// unknown rather than empty. Left out when X sent the text.
	UnknownFields []string `json:"unknownFields"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                     respjson.Field
		BookmarkCount          respjson.Field
		LikeCount              respjson.Field
		QuoteCount             respjson.Field
		ReplyCount             respjson.Field
		RetweetCount           respjson.Field
		Text                   respjson.Field
		ViewCount              respjson.Field
		Article                respjson.Field
		Author                 respjson.Field
		AuthorUnavailable      respjson.Field
		Card                   respjson.Field
		CashtagAttachments     respjson.Field
		CommunityID            respjson.Field
		CommunityName          respjson.Field
		CommunityNote          respjson.Field
		ContentDisclosure      respjson.Field
		ConversationControl    respjson.Field
		ConversationID         respjson.Field
		Coordinates            respjson.Field
		CreatedAt              respjson.Field
		DisplayTextRange       respjson.Field
		Edit                   respjson.Field
		Entities               respjson.Field
		ExclusiveContent       respjson.Field
		Geo                    respjson.Field
		GrokShareAttachment    respjson.Field
		GrokTranslatedPost     respjson.Field
		HasCommunityNotes      respjson.Field
		InReplyToID            respjson.Field
		InReplyToUserID        respjson.Field
		InReplyToUsername      respjson.Field
		Interstitial           respjson.Field
		IsLimitedReply         respjson.Field
		IsNoteTweet            respjson.Field
		IsPinned               respjson.Field
		IsQuoteStatus          respjson.Field
		IsReply                respjson.Field
		IsRetweet              respjson.Field
		IsTranslatable         respjson.Field
		JetfuelAttachment      respjson.Field
		Lang                   respjson.Field
		LimitedActions         respjson.Field
		Media                  respjson.Field
		NoteTweet              respjson.Field
		ParentTweet            respjson.Field
		Place                  respjson.Field
		PossiblySensitive      respjson.Field
		PostCta                respjson.Field
		PreviousCounts         respjson.Field
		QuotedTweet            respjson.Field
		QuotedTweetID          respjson.Field
		QuotedTweetPermalink   respjson.Field
		QuotedTweetUnavailable respjson.Field
		ReactionContext        respjson.Field
		RetweetedTweet         respjson.Field
		RetweetedAt            respjson.Field
		Scopes                 respjson.Field
		SoftIntervention       respjson.Field
		Source                 respjson.Field
		SportsContext          respjson.Field
		SubscriberBadge        respjson.Field
		TextHighlights         respjson.Field
		Tombstone              respjson.Field
		Type                   respjson.Field
		UnavailableAuthorID    respjson.Field
		UnknownCounts          respjson.Field
		UnmentionedUserIDs     respjson.Field
		URL                    respjson.Field
		ViewState              respjson.Field
		UnknownFields          respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// RawJSON returns the unmodified JSON received from the API.
func (r MonitorEmbeddedPost) RawJSON() string { return r.JSON.raw }
func (r *MonitorEmbeddedPost) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// MonitorPostCard describes a public card and the monitor profiles it
// references.
type MonitorPostCard struct {
	ID string `json:"id"`
	// Public metadata whose fields are defined by X.
	BindingValues PublicMetadata `json:"bindingValues"`
	Name          string         `json:"name"`
	// Public metadata whose fields are defined by X.
	Platform PublicMetadata `json:"platform"`
	URL      string         `json:"url"`
	// Unresolved card user references.
	UserReferenceErrors []SearchTweetCardUserReferenceError `json:"userReferenceErrors"`
	UserReferences      []MonitorProfile                    `json:"userReferences"`
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

// RawJSON returns the unmodified JSON received from the API.
func (r MonitorPostCard) RawJSON() string { return r.JSON.raw }
func (r *MonitorPostCard) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// MonitorPostReactionContext holds the public post and user referenced by this
// reaction.
type MonitorPostReactionContext struct {
	// Embedded post with public fields and legacy profile aliases.
	ReactedToPost MonitorEmbeddedPost `json:"reactedToPost"`
	// Referenced post ID.
	ReactedToPostID string `json:"reactedToPostId"`
	// Public profile with retained legacy monitor fields.
	ReactedToUser MonitorProfile `json:"reactedToUser"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ReactedToPost   respjson.Field
		ReactedToPostID respjson.Field
		ReactedToUser   respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// RawJSON returns the unmodified JSON received from the API.
func (r MonitorPostReactionContext) RawJSON() string { return r.JSON.raw }
func (r *MonitorPostReactionContext) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}
