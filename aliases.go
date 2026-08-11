// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package xtwitterscraper

import (
	"github.com/Xquik-dev/x-twitter-scraper-go/internal/apierror"
	"github.com/Xquik-dev/x-twitter-scraper-go/packages/param"
	"github.com/Xquik-dev/x-twitter-scraper-go/shared"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

type Error = apierror.Error

// Content disclosure metadata shown by X when a tweet is labeled as paid
// partnership content or AI-generated media.
//
// This is an alias to an internal type.
type ContentDisclosure = shared.ContentDisclosure

// This is an alias to an internal type.
type ContentDisclosureAdvertising = shared.ContentDisclosureAdvertising

// This is an alias to an internal type.
type ContentDisclosureAIGenerated = shared.ContentDisclosureAIGenerated

// Quoted or retweeted tweet context. Every object includes id, text, and
// engagement metrics. A zero metric can mean X did not report the count. Author,
// media, and conversation fields appear when available.
//
// This is an alias to an internal type.
type EmbeddedTweet = shared.EmbeddedTweet

// Article metadata attached to a tweet.
//
// This is an alias to an internal type.
type EmbeddedTweetArticle = shared.EmbeddedTweetArticle

// Public card metadata attached to a tweet.
//
// This is an alias to an internal type.
type EmbeddedTweetCard = shared.EmbeddedTweetCard

// Community Note presentation metadata returned by X.
//
// This is an alias to an internal type.
type EmbeddedTweetCommunityNote = shared.EmbeddedTweetCommunityNote

// Edit history metadata returned by X.
//
// This is an alias to an internal type.
type EmbeddedTweetEdit = shared.EmbeddedTweetEdit

// Complete Note Tweet content and rich-text metadata.
//
// This is an alias to an internal type.
type EmbeddedTweetNoteTweet = shared.EmbeddedTweetNoteTweet

// This is an alias to an internal type.
type EmbeddedTweetNoteTweetRichtextTag = shared.EmbeddedTweetNoteTweetRichtextTag

// Public place metadata attached to a tweet.
//
// This is an alias to an internal type.
type EmbeddedTweetPlace = shared.EmbeddedTweetPlace

// Engagement counts retained from a prior tweet edit.
//
// This is an alias to an internal type.
type EmbeddedTweetPreviousCounts = shared.EmbeddedTweetPreviousCounts

// Type of monitor event fired when account activity occurs.
//
// This is an alias to an internal type.
type EventType = shared.EventType

// Equals "tweet.new"
const EventTypeTweetNew = shared.EventTypeTweetNew

// Equals "tweet.reply"
const EventTypeTweetReply = shared.EventTypeTweetReply

// Equals "tweet.retweet"
const EventTypeTweetRetweet = shared.EventTypeTweetRetweet

// Equals "tweet.quote"
const EventTypeTweetQuote = shared.EventTypeTweetQuote

// Equals "tweet.media"
const EventTypeTweetMedia = shared.EventTypeTweetMedia

// Equals "tweet.link"
const EventTypeTweetLink = shared.EventTypeTweetLink

// Equals "tweet.poll"
const EventTypeTweetPoll = shared.EventTypeTweetPoll

// Equals "tweet.mention"
const EventTypeTweetMention = shared.EventTypeTweetMention

// Equals "tweet.hashtag"
const EventTypeTweetHashtag = shared.EventTypeTweetHashtag

// Equals "tweet.longform"
const EventTypeTweetLongform = shared.EventTypeTweetLongform

// Equals "profile.avatar.changed"
const EventTypeProfileAvatarChanged = shared.EventTypeProfileAvatarChanged

// Equals "profile.banner.changed"
const EventTypeProfileBannerChanged = shared.EventTypeProfileBannerChanged

// Equals "profile.name.changed"
const EventTypeProfileNameChanged = shared.EventTypeProfileNameChanged

// Equals "profile.username.changed"
const EventTypeProfileUsernameChanged = shared.EventTypeProfileUsernameChanged

// Equals "profile.bio.changed"
const EventTypeProfileBioChanged = shared.EventTypeProfileBioChanged

// Equals "profile.location.changed"
const EventTypeProfileLocationChanged = shared.EventTypeProfileLocationChanged

// Equals "profile.url.changed"
const EventTypeProfileURLChanged = shared.EventTypeProfileURLChanged

// Equals "profile.verified.changed"
const EventTypeProfileVerifiedChanged = shared.EventTypeProfileVerifiedChanged

// Equals "profile.protected.changed"
const EventTypeProfileProtectedChanged = shared.EventTypeProfileProtectedChanged

// Equals "profile.pinned_tweet.changed"
const EventTypeProfilePinnedTweetChanged = shared.EventTypeProfilePinnedTweetChanged

// Equals "profile.unavailable.changed"
const EventTypeProfileUnavailableChanged = shared.EventTypeProfileUnavailableChanged

// No-mode search, user Tweet, user reply, and direct reply reads use automatic
// coverage. Shape, filters, aliases, and billing stay compatible. Unprefixed
// cursors remain legacy. Follow next_cursor while has_next_page is true. An empty
// filtered page can still have has_next_page true.
//
// This is an alias to an internal type.
type PaginatedTweets = shared.PaginatedTweets

// Paginated user profiles. No-mode follower, following, and verified follower
// requests merge independent views automatically. Response fields, page size,
// aliases, filters, and per-returned-profile billing stay unchanged. Existing
// unprefixed cursors retain legacy behavior. Follow next_cursor while
// has_next_page is true.
//
// This is an alias to an internal type.
type PaginatedUsers = shared.PaginatedUsers

// Tweet returned from search results with inline author info. A zero metric can
// mean X did not report the count.
//
// This is an alias to an internal type.
type SearchTweet = shared.SearchTweet

// Article metadata attached to a tweet.
//
// This is an alias to an internal type.
type SearchTweetArticle = shared.SearchTweetArticle

// Public card metadata attached to a tweet.
//
// This is an alias to an internal type.
type SearchTweetCard = shared.SearchTweetCard

// Community Note presentation metadata returned by X.
//
// This is an alias to an internal type.
type SearchTweetCommunityNote = shared.SearchTweetCommunityNote

// Edit history metadata returned by X.
//
// This is an alias to an internal type.
type SearchTweetEdit = shared.SearchTweetEdit

// Complete Note Tweet content and rich-text metadata.
//
// This is an alias to an internal type.
type SearchTweetNoteTweet = shared.SearchTweetNoteTweet

// This is an alias to an internal type.
type SearchTweetNoteTweetRichtextTag = shared.SearchTweetNoteTweetRichtextTag

// Public place metadata attached to a tweet.
//
// This is an alias to an internal type.
type SearchTweetPlace = shared.SearchTweetPlace

// Engagement counts retained from a prior tweet edit.
//
// This is an alias to an internal type.
type SearchTweetPreviousCounts = shared.SearchTweetPreviousCounts

// Normalized media attached to a tweet.
//
// This is an alias to an internal type.
type TweetMedia = shared.TweetMedia

// This is an alias to an internal type.
type TweetMediaType = shared.TweetMediaType

// Equals "photo"
const TweetMediaTypePhoto = shared.TweetMediaTypePhoto

// Equals "video"
const TweetMediaTypeVideo = shared.TweetMediaTypeVideo

// Equals "animated_gif"
const TweetMediaTypeAnimatedGif = shared.TweetMediaTypeAnimatedGif

// This is an alias to an internal type.
type TweetMediaFaceRect = shared.TweetMediaFaceRect

// This is an alias to an internal type.
type TweetMediaFocusRect = shared.TweetMediaFocusRect

// This is an alias to an internal type.
type TweetMediaSize = shared.TweetMediaSize

// This is an alias to an internal type.
type TweetMediaVideoVariant = shared.TweetMediaVideoVariant

// X user profile with bio, follower counts, and verification status.
//
// This is an alias to an internal type.
type UserProfile = shared.UserProfile

// Organization affiliation label shown on an X profile.
//
// This is an alias to an internal type.
type UserProfileAffiliatesHighlightedLabel = shared.UserProfileAffiliatesHighlightedLabel

// Profile highlight availability and count metadata.
//
// This is an alias to an internal type.
type UserProfileHighlightsInfo = shared.UserProfileHighlightsInfo

// Identity verification metadata displayed by X.
//
// This is an alias to an internal type.
type UserProfileIdentityVerification = shared.UserProfileIdentityVerification
