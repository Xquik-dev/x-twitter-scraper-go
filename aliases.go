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

// Paginated tweet results. The item count can be lower than pageSize when the
// source returns fewer tweets, filters remove tweets, or remaining credits cover
// fewer results. Follow next_cursor while has_next_page is true. An empty page can
// still have has_next_page true after filtering. Zero affordable results returns
// 402 insufficient_credits.
//
// This is an alias to an internal type.
type PaginatedTweets = shared.PaginatedTweets

// Paginated user profiles. The item count can be lower than pageSize when the
// source returns fewer profiles or remaining credits cover fewer results. Follow
// next_cursor while has_next_page is true. A relationship can naturally contain
// fewer profiles than requested. Zero affordable results returns 402
// insufficient_credits.
//
// This is an alias to an internal type.
type PaginatedUsers = shared.PaginatedUsers

// Tweet returned from search results with inline author info. A zero metric can
// mean X did not report the count.
//
// This is an alias to an internal type.
type SearchTweet = shared.SearchTweet

// Content disclosure metadata shown by X when a tweet is labeled as paid
// partnership content or AI-generated media.
//
// This is an alias to an internal type.
type SearchTweetContentDisclosure = shared.SearchTweetContentDisclosure

// This is an alias to an internal type.
type SearchTweetContentDisclosureAdvertising = shared.SearchTweetContentDisclosureAdvertising

// This is an alias to an internal type.
type SearchTweetContentDisclosureAIGenerated = shared.SearchTweetContentDisclosureAIGenerated

// Normalized media attached to a tweet.
//
// This is an alias to an internal type.
type SearchTweetMedia = shared.SearchTweetMedia

// This is an alias to an internal type.
type SearchTweetMediaVideoVariant = shared.SearchTweetMediaVideoVariant

// Quoted or retweeted tweet context. Every object includes id, text, and
// engagement metrics. A zero metric can mean X did not report the count. Author,
// media, and conversation fields appear when available.
//
// This is an alias to an internal type.
type SearchTweetQuotedTweet = shared.SearchTweetQuotedTweet

// Content disclosure metadata shown by X when a tweet is labeled as paid
// partnership content or AI-generated media.
//
// This is an alias to an internal type.
type SearchTweetQuotedTweetContentDisclosure = shared.SearchTweetQuotedTweetContentDisclosure

// This is an alias to an internal type.
type SearchTweetQuotedTweetContentDisclosureAdvertising = shared.SearchTweetQuotedTweetContentDisclosureAdvertising

// This is an alias to an internal type.
type SearchTweetQuotedTweetContentDisclosureAIGenerated = shared.SearchTweetQuotedTweetContentDisclosureAIGenerated

// Normalized media attached to a tweet.
//
// This is an alias to an internal type.
type SearchTweetQuotedTweetMedia = shared.SearchTweetQuotedTweetMedia

// This is an alias to an internal type.
type SearchTweetQuotedTweetMediaVideoVariant = shared.SearchTweetQuotedTweetMediaVideoVariant

// Quoted or retweeted tweet context. Every object includes id, text, and
// engagement metrics. A zero metric can mean X did not report the count. Author,
// media, and conversation fields appear when available.
//
// This is an alias to an internal type.
type SearchTweetRetweetedTweet = shared.SearchTweetRetweetedTweet

// Content disclosure metadata shown by X when a tweet is labeled as paid
// partnership content or AI-generated media.
//
// This is an alias to an internal type.
type SearchTweetRetweetedTweetContentDisclosure = shared.SearchTweetRetweetedTweetContentDisclosure

// This is an alias to an internal type.
type SearchTweetRetweetedTweetContentDisclosureAdvertising = shared.SearchTweetRetweetedTweetContentDisclosureAdvertising

// This is an alias to an internal type.
type SearchTweetRetweetedTweetContentDisclosureAIGenerated = shared.SearchTweetRetweetedTweetContentDisclosureAIGenerated

// Normalized media attached to a tweet.
//
// This is an alias to an internal type.
type SearchTweetRetweetedTweetMedia = shared.SearchTweetRetweetedTweetMedia

// This is an alias to an internal type.
type SearchTweetRetweetedTweetMediaVideoVariant = shared.SearchTweetRetweetedTweetMediaVideoVariant

// X user profile with bio, follower counts, and verification status.
//
// This is an alias to an internal type.
type UserProfile = shared.UserProfile
