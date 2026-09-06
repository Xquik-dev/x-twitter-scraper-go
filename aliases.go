// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

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

// Quoted or retweeted tweet context.
//
// This is an alias to an internal type.
type EmbeddedTweet = shared.EmbeddedTweet

// Describes an X Article preview and its lifecycle metadata.
//
// This is an alias to an internal type.
type EmbeddedTweetArticle = shared.EmbeddedTweetArticle

// Describes a public card and its referenced profiles.
//
// This is an alias to an internal type.
type EmbeddedTweetCard = shared.EmbeddedTweetCard

// This is an alias to an internal type.
type EmbeddedTweetCardUserReferenceError = shared.EmbeddedTweetCardUserReferenceError

// Community Note presentation metadata returned by X.
//
// This is an alias to an internal type.
type EmbeddedTweetCommunityNote = shared.EmbeddedTweetCommunityNote

// Public reply policy and conversation owner.
//
// This is an alias to an internal type.
type EmbeddedTweetConversationControl = shared.EmbeddedTweetConversationControl

// Lists edit-chain identifiers and the remaining edit window.
//
// This is an alias to an internal type.
type EmbeddedTweetEdit = shared.EmbeddedTweetEdit

// Lists hashtags, symbols, links, and mentions from tweet text.
//
// This is an alias to an internal type.
type EmbeddedTweetEntities = shared.EmbeddedTweetEntities

// Provides hashtag text and source offsets within a tweet.
//
// This is an alias to an internal type.
type EmbeddedTweetEntitiesHashtag = shared.EmbeddedTweetEntitiesHashtag

// Indexed smart-tag, cashtag, or video timestamp metadata.
//
// This is an alias to an internal type.
type EmbeddedTweetEntitiesSmarttag = shared.EmbeddedTweetEntitiesSmarttag

// This is an alias to an internal type.
type EmbeddedTweetEntitiesSmarttagTag = shared.EmbeddedTweetEntitiesSmarttagTag

// This is an alias to an internal type.
type EmbeddedTweetEntitiesSmarttagTagInfo = shared.EmbeddedTweetEntitiesSmarttagTagInfo

// This is an alias to an internal type.
type EmbeddedTweetEntitiesSmarttagTagInfoInfo = shared.EmbeddedTweetEntitiesSmarttagTagInfoInfo

// Indexed smart-tag, cashtag, or video timestamp metadata.
//
// This is an alias to an internal type.
type EmbeddedTweetEntitiesSymbol = shared.EmbeddedTweetEntitiesSymbol

// This is an alias to an internal type.
type EmbeddedTweetEntitiesSymbolTag = shared.EmbeddedTweetEntitiesSymbolTag

// This is an alias to an internal type.
type EmbeddedTweetEntitiesSymbolTagInfo = shared.EmbeddedTweetEntitiesSymbolTagInfo

// This is an alias to an internal type.
type EmbeddedTweetEntitiesSymbolTagInfoInfo = shared.EmbeddedTweetEntitiesSymbolTagInfoInfo

// Indexed smart-tag, cashtag, or video timestamp metadata.
//
// This is an alias to an internal type.
type EmbeddedTweetEntitiesTimestamp = shared.EmbeddedTweetEntitiesTimestamp

// This is an alias to an internal type.
type EmbeddedTweetEntitiesTimestampTag = shared.EmbeddedTweetEntitiesTimestampTag

// This is an alias to an internal type.
type EmbeddedTweetEntitiesTimestampTagInfo = shared.EmbeddedTweetEntitiesTimestampTagInfo

// This is an alias to an internal type.
type EmbeddedTweetEntitiesTimestampTagInfoInfo = shared.EmbeddedTweetEntitiesTimestampTagInfoInfo

// Provides shortened, display, and expanded URLs from tweet text.
//
// This is an alias to an internal type.
type EmbeddedTweetEntitiesURL = shared.EmbeddedTweetEntitiesURL

// Provides profile identity and source offsets for a mention.
//
// This is an alias to an internal type.
type EmbeddedTweetEntitiesUserMention = shared.EmbeddedTweetEntitiesUserMention

// This is an alias to an internal type.
type EmbeddedTweetLimitedAction = shared.EmbeddedTweetLimitedAction

// This is an alias to an internal type.
type EmbeddedTweetLimitedActionPrompt = shared.EmbeddedTweetLimitedActionPrompt

// Complete Note Tweet content and rich-text metadata.
//
// This is an alias to an internal type.
type EmbeddedTweetNoteTweet = shared.EmbeddedTweetNoteTweet

// Lists hashtags, symbols, links, and mentions from tweet text.
//
// This is an alias to an internal type.
type EmbeddedTweetNoteTweetEntities = shared.EmbeddedTweetNoteTweetEntities

// Provides hashtag text and source offsets within a tweet.
//
// This is an alias to an internal type.
type EmbeddedTweetNoteTweetEntitiesHashtag = shared.EmbeddedTweetNoteTweetEntitiesHashtag

// Indexed smart-tag, cashtag, or video timestamp metadata.
//
// This is an alias to an internal type.
type EmbeddedTweetNoteTweetEntitiesSmarttag = shared.EmbeddedTweetNoteTweetEntitiesSmarttag

// This is an alias to an internal type.
type EmbeddedTweetNoteTweetEntitiesSmarttagTag = shared.EmbeddedTweetNoteTweetEntitiesSmarttagTag

// This is an alias to an internal type.
type EmbeddedTweetNoteTweetEntitiesSmarttagTagInfo = shared.EmbeddedTweetNoteTweetEntitiesSmarttagTagInfo

// This is an alias to an internal type.
type EmbeddedTweetNoteTweetEntitiesSmarttagTagInfoInfo = shared.EmbeddedTweetNoteTweetEntitiesSmarttagTagInfoInfo

// Indexed smart-tag, cashtag, or video timestamp metadata.
//
// This is an alias to an internal type.
type EmbeddedTweetNoteTweetEntitiesSymbol = shared.EmbeddedTweetNoteTweetEntitiesSymbol

// This is an alias to an internal type.
type EmbeddedTweetNoteTweetEntitiesSymbolTag = shared.EmbeddedTweetNoteTweetEntitiesSymbolTag

// This is an alias to an internal type.
type EmbeddedTweetNoteTweetEntitiesSymbolTagInfo = shared.EmbeddedTweetNoteTweetEntitiesSymbolTagInfo

// This is an alias to an internal type.
type EmbeddedTweetNoteTweetEntitiesSymbolTagInfoInfo = shared.EmbeddedTweetNoteTweetEntitiesSymbolTagInfoInfo

// Indexed smart-tag, cashtag, or video timestamp metadata.
//
// This is an alias to an internal type.
type EmbeddedTweetNoteTweetEntitiesTimestamp = shared.EmbeddedTweetNoteTweetEntitiesTimestamp

// This is an alias to an internal type.
type EmbeddedTweetNoteTweetEntitiesTimestampTag = shared.EmbeddedTweetNoteTweetEntitiesTimestampTag

// This is an alias to an internal type.
type EmbeddedTweetNoteTweetEntitiesTimestampTagInfo = shared.EmbeddedTweetNoteTweetEntitiesTimestampTagInfo

// This is an alias to an internal type.
type EmbeddedTweetNoteTweetEntitiesTimestampTagInfoInfo = shared.EmbeddedTweetNoteTweetEntitiesTimestampTagInfoInfo

// Provides shortened, display, and expanded URLs from tweet text.
//
// This is an alias to an internal type.
type EmbeddedTweetNoteTweetEntitiesURL = shared.EmbeddedTweetNoteTweetEntitiesURL

// Provides profile identity and source offsets for a mention.
//
// This is an alias to an internal type.
type EmbeddedTweetNoteTweetEntitiesUserMention = shared.EmbeddedTweetNoteTweetEntitiesUserMention

// This is an alias to an internal type.
type EmbeddedTweetNoteTweetInlineMedia = shared.EmbeddedTweetNoteTweetInlineMedia

// This is an alias to an internal type.
type EmbeddedTweetNoteTweetRichtextTag = shared.EmbeddedTweetNoteTweetRichtextTag

// Describes public place metadata on a geotagged tweet.
//
// This is an alias to an internal type.
type EmbeddedTweetPlace = shared.EmbeddedTweetPlace

// Engagement counts retained from a prior tweet edit.
//
// This is an alias to an internal type.
type EmbeddedTweetPreviousCounts = shared.EmbeddedTweetPreviousCounts

// Nested tweet context at depth 2.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweet = shared.EmbeddedTweetQuotedTweet

// Describes an X Article preview and its lifecycle metadata.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetArticle = shared.EmbeddedTweetQuotedTweetArticle

// Describes a public card and its referenced profiles.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetCard = shared.EmbeddedTweetQuotedTweetCard

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetCardUserReferenceError = shared.EmbeddedTweetQuotedTweetCardUserReferenceError

// Community Note presentation metadata returned by X.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetCommunityNote = shared.EmbeddedTweetQuotedTweetCommunityNote

// Public reply policy and conversation owner.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetConversationControl = shared.EmbeddedTweetQuotedTweetConversationControl

// Lists edit-chain identifiers and the remaining edit window.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetEdit = shared.EmbeddedTweetQuotedTweetEdit

// Lists hashtags, symbols, links, and mentions from tweet text.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetEntities = shared.EmbeddedTweetQuotedTweetEntities

// Provides hashtag text and source offsets within a tweet.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetEntitiesHashtag = shared.EmbeddedTweetQuotedTweetEntitiesHashtag

// Indexed smart-tag, cashtag, or video timestamp metadata.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetEntitiesSmarttag = shared.EmbeddedTweetQuotedTweetEntitiesSmarttag

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetEntitiesSmarttagTag = shared.EmbeddedTweetQuotedTweetEntitiesSmarttagTag

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetEntitiesSmarttagTagInfo = shared.EmbeddedTweetQuotedTweetEntitiesSmarttagTagInfo

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetEntitiesSmarttagTagInfoInfo = shared.EmbeddedTweetQuotedTweetEntitiesSmarttagTagInfoInfo

// Indexed smart-tag, cashtag, or video timestamp metadata.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetEntitiesSymbol = shared.EmbeddedTweetQuotedTweetEntitiesSymbol

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetEntitiesSymbolTag = shared.EmbeddedTweetQuotedTweetEntitiesSymbolTag

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetEntitiesSymbolTagInfo = shared.EmbeddedTweetQuotedTweetEntitiesSymbolTagInfo

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetEntitiesSymbolTagInfoInfo = shared.EmbeddedTweetQuotedTweetEntitiesSymbolTagInfoInfo

// Indexed smart-tag, cashtag, or video timestamp metadata.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetEntitiesTimestamp = shared.EmbeddedTweetQuotedTweetEntitiesTimestamp

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetEntitiesTimestampTag = shared.EmbeddedTweetQuotedTweetEntitiesTimestampTag

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetEntitiesTimestampTagInfo = shared.EmbeddedTweetQuotedTweetEntitiesTimestampTagInfo

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetEntitiesTimestampTagInfoInfo = shared.EmbeddedTweetQuotedTweetEntitiesTimestampTagInfoInfo

// Provides shortened, display, and expanded URLs from tweet text.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetEntitiesURL = shared.EmbeddedTweetQuotedTweetEntitiesURL

// Provides profile identity and source offsets for a mention.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetEntitiesUserMention = shared.EmbeddedTweetQuotedTweetEntitiesUserMention

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetLimitedAction = shared.EmbeddedTweetQuotedTweetLimitedAction

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetLimitedActionPrompt = shared.EmbeddedTweetQuotedTweetLimitedActionPrompt

// Complete Note Tweet content and rich-text metadata.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetNoteTweet = shared.EmbeddedTweetQuotedTweetNoteTweet

// Lists hashtags, symbols, links, and mentions from tweet text.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetNoteTweetEntities = shared.EmbeddedTweetQuotedTweetNoteTweetEntities

// Provides hashtag text and source offsets within a tweet.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetNoteTweetEntitiesHashtag = shared.EmbeddedTweetQuotedTweetNoteTweetEntitiesHashtag

// Indexed smart-tag, cashtag, or video timestamp metadata.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetNoteTweetEntitiesSmarttag = shared.EmbeddedTweetQuotedTweetNoteTweetEntitiesSmarttag

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetNoteTweetEntitiesSmarttagTag = shared.EmbeddedTweetQuotedTweetNoteTweetEntitiesSmarttagTag

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetNoteTweetEntitiesSmarttagTagInfo = shared.EmbeddedTweetQuotedTweetNoteTweetEntitiesSmarttagTagInfo

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetNoteTweetEntitiesSmarttagTagInfoInfo = shared.EmbeddedTweetQuotedTweetNoteTweetEntitiesSmarttagTagInfoInfo

// Indexed smart-tag, cashtag, or video timestamp metadata.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetNoteTweetEntitiesSymbol = shared.EmbeddedTweetQuotedTweetNoteTweetEntitiesSymbol

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetNoteTweetEntitiesSymbolTag = shared.EmbeddedTweetQuotedTweetNoteTweetEntitiesSymbolTag

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetNoteTweetEntitiesSymbolTagInfo = shared.EmbeddedTweetQuotedTweetNoteTweetEntitiesSymbolTagInfo

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetNoteTweetEntitiesSymbolTagInfoInfo = shared.EmbeddedTweetQuotedTweetNoteTweetEntitiesSymbolTagInfoInfo

// Indexed smart-tag, cashtag, or video timestamp metadata.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetNoteTweetEntitiesTimestamp = shared.EmbeddedTweetQuotedTweetNoteTweetEntitiesTimestamp

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetNoteTweetEntitiesTimestampTag = shared.EmbeddedTweetQuotedTweetNoteTweetEntitiesTimestampTag

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetNoteTweetEntitiesTimestampTagInfo = shared.EmbeddedTweetQuotedTweetNoteTweetEntitiesTimestampTagInfo

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetNoteTweetEntitiesTimestampTagInfoInfo = shared.EmbeddedTweetQuotedTweetNoteTweetEntitiesTimestampTagInfoInfo

// Provides shortened, display, and expanded URLs from tweet text.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetNoteTweetEntitiesURL = shared.EmbeddedTweetQuotedTweetNoteTweetEntitiesURL

// Provides profile identity and source offsets for a mention.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetNoteTweetEntitiesUserMention = shared.EmbeddedTweetQuotedTweetNoteTweetEntitiesUserMention

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetNoteTweetInlineMedia = shared.EmbeddedTweetQuotedTweetNoteTweetInlineMedia

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetNoteTweetRichtextTag = shared.EmbeddedTweetQuotedTweetNoteTweetRichtextTag

// Describes public place metadata on a geotagged tweet.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetPlace = shared.EmbeddedTweetQuotedTweetPlace

// Engagement counts retained from a prior tweet edit.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetPreviousCounts = shared.EmbeddedTweetQuotedTweetPreviousCounts

// Nested tweet context at depth 3.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetQuotedTweet = shared.EmbeddedTweetQuotedTweetQuotedTweet

// Describes an X Article preview and its lifecycle metadata.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetQuotedTweetArticle = shared.EmbeddedTweetQuotedTweetQuotedTweetArticle

// Describes a public card and its referenced profiles.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetQuotedTweetCard = shared.EmbeddedTweetQuotedTweetQuotedTweetCard

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetQuotedTweetCardUserReferenceError = shared.EmbeddedTweetQuotedTweetQuotedTweetCardUserReferenceError

// Community Note presentation metadata returned by X.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetQuotedTweetCommunityNote = shared.EmbeddedTweetQuotedTweetQuotedTweetCommunityNote

// Public reply policy and conversation owner.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetQuotedTweetConversationControl = shared.EmbeddedTweetQuotedTweetQuotedTweetConversationControl

// Lists edit-chain identifiers and the remaining edit window.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetQuotedTweetEdit = shared.EmbeddedTweetQuotedTweetQuotedTweetEdit

// Lists hashtags, symbols, links, and mentions from tweet text.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetQuotedTweetEntities = shared.EmbeddedTweetQuotedTweetQuotedTweetEntities

// Provides hashtag text and source offsets within a tweet.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetQuotedTweetEntitiesHashtag = shared.EmbeddedTweetQuotedTweetQuotedTweetEntitiesHashtag

// Indexed smart-tag, cashtag, or video timestamp metadata.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetQuotedTweetEntitiesSmarttag = shared.EmbeddedTweetQuotedTweetQuotedTweetEntitiesSmarttag

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetQuotedTweetEntitiesSmarttagTag = shared.EmbeddedTweetQuotedTweetQuotedTweetEntitiesSmarttagTag

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetQuotedTweetEntitiesSmarttagTagInfo = shared.EmbeddedTweetQuotedTweetQuotedTweetEntitiesSmarttagTagInfo

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetQuotedTweetEntitiesSmarttagTagInfoInfo = shared.EmbeddedTweetQuotedTweetQuotedTweetEntitiesSmarttagTagInfoInfo

// Indexed smart-tag, cashtag, or video timestamp metadata.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetQuotedTweetEntitiesSymbol = shared.EmbeddedTweetQuotedTweetQuotedTweetEntitiesSymbol

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetQuotedTweetEntitiesSymbolTag = shared.EmbeddedTweetQuotedTweetQuotedTweetEntitiesSymbolTag

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetQuotedTweetEntitiesSymbolTagInfo = shared.EmbeddedTweetQuotedTweetQuotedTweetEntitiesSymbolTagInfo

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetQuotedTweetEntitiesSymbolTagInfoInfo = shared.EmbeddedTweetQuotedTweetQuotedTweetEntitiesSymbolTagInfoInfo

// Indexed smart-tag, cashtag, or video timestamp metadata.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetQuotedTweetEntitiesTimestamp = shared.EmbeddedTweetQuotedTweetQuotedTweetEntitiesTimestamp

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetQuotedTweetEntitiesTimestampTag = shared.EmbeddedTweetQuotedTweetQuotedTweetEntitiesTimestampTag

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetQuotedTweetEntitiesTimestampTagInfo = shared.EmbeddedTweetQuotedTweetQuotedTweetEntitiesTimestampTagInfo

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetQuotedTweetEntitiesTimestampTagInfoInfo = shared.EmbeddedTweetQuotedTweetQuotedTweetEntitiesTimestampTagInfoInfo

// Provides shortened, display, and expanded URLs from tweet text.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetQuotedTweetEntitiesURL = shared.EmbeddedTweetQuotedTweetQuotedTweetEntitiesURL

// Provides profile identity and source offsets for a mention.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetQuotedTweetEntitiesUserMention = shared.EmbeddedTweetQuotedTweetQuotedTweetEntitiesUserMention

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetQuotedTweetLimitedAction = shared.EmbeddedTweetQuotedTweetQuotedTweetLimitedAction

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetQuotedTweetLimitedActionPrompt = shared.EmbeddedTweetQuotedTweetQuotedTweetLimitedActionPrompt

// Complete Note Tweet content and rich-text metadata.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetQuotedTweetNoteTweet = shared.EmbeddedTweetQuotedTweetQuotedTweetNoteTweet

// Lists hashtags, symbols, links, and mentions from tweet text.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetQuotedTweetNoteTweetEntities = shared.EmbeddedTweetQuotedTweetQuotedTweetNoteTweetEntities

// Provides hashtag text and source offsets within a tweet.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetQuotedTweetNoteTweetEntitiesHashtag = shared.EmbeddedTweetQuotedTweetQuotedTweetNoteTweetEntitiesHashtag

// Indexed smart-tag, cashtag, or video timestamp metadata.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetQuotedTweetNoteTweetEntitiesSmarttag = shared.EmbeddedTweetQuotedTweetQuotedTweetNoteTweetEntitiesSmarttag

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetQuotedTweetNoteTweetEntitiesSmarttagTag = shared.EmbeddedTweetQuotedTweetQuotedTweetNoteTweetEntitiesSmarttagTag

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetQuotedTweetNoteTweetEntitiesSmarttagTagInfo = shared.EmbeddedTweetQuotedTweetQuotedTweetNoteTweetEntitiesSmarttagTagInfo

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetQuotedTweetNoteTweetEntitiesSmarttagTagInfoInfo = shared.EmbeddedTweetQuotedTweetQuotedTweetNoteTweetEntitiesSmarttagTagInfoInfo

// Indexed smart-tag, cashtag, or video timestamp metadata.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetQuotedTweetNoteTweetEntitiesSymbol = shared.EmbeddedTweetQuotedTweetQuotedTweetNoteTweetEntitiesSymbol

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetQuotedTweetNoteTweetEntitiesSymbolTag = shared.EmbeddedTweetQuotedTweetQuotedTweetNoteTweetEntitiesSymbolTag

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetQuotedTweetNoteTweetEntitiesSymbolTagInfo = shared.EmbeddedTweetQuotedTweetQuotedTweetNoteTweetEntitiesSymbolTagInfo

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetQuotedTweetNoteTweetEntitiesSymbolTagInfoInfo = shared.EmbeddedTweetQuotedTweetQuotedTweetNoteTweetEntitiesSymbolTagInfoInfo

// Indexed smart-tag, cashtag, or video timestamp metadata.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetQuotedTweetNoteTweetEntitiesTimestamp = shared.EmbeddedTweetQuotedTweetQuotedTweetNoteTweetEntitiesTimestamp

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetQuotedTweetNoteTweetEntitiesTimestampTag = shared.EmbeddedTweetQuotedTweetQuotedTweetNoteTweetEntitiesTimestampTag

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetQuotedTweetNoteTweetEntitiesTimestampTagInfo = shared.EmbeddedTweetQuotedTweetQuotedTweetNoteTweetEntitiesTimestampTagInfo

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetQuotedTweetNoteTweetEntitiesTimestampTagInfoInfo = shared.EmbeddedTweetQuotedTweetQuotedTweetNoteTweetEntitiesTimestampTagInfoInfo

// Provides shortened, display, and expanded URLs from tweet text.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetQuotedTweetNoteTweetEntitiesURL = shared.EmbeddedTweetQuotedTweetQuotedTweetNoteTweetEntitiesURL

// Provides profile identity and source offsets for a mention.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetQuotedTweetNoteTweetEntitiesUserMention = shared.EmbeddedTweetQuotedTweetQuotedTweetNoteTweetEntitiesUserMention

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetQuotedTweetNoteTweetInlineMedia = shared.EmbeddedTweetQuotedTweetQuotedTweetNoteTweetInlineMedia

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetQuotedTweetNoteTweetRichtextTag = shared.EmbeddedTweetQuotedTweetQuotedTweetNoteTweetRichtextTag

// Describes public place metadata on a geotagged tweet.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetQuotedTweetPlace = shared.EmbeddedTweetQuotedTweetQuotedTweetPlace

// Engagement counts retained from a prior tweet edit.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetQuotedTweetPreviousCounts = shared.EmbeddedTweetQuotedTweetQuotedTweetPreviousCounts

// Final nested tweet context at depth 4.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetQuotedTweetQuotedTweet = shared.EmbeddedTweetQuotedTweetQuotedTweetQuotedTweet

// Describes an X Article preview and its lifecycle metadata.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetArticle = shared.EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetArticle

// Describes a public card and its referenced profiles.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetCard = shared.EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetCard

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetCardUserReferenceError = shared.EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetCardUserReferenceError

// Community Note presentation metadata returned by X.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetCommunityNote = shared.EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetCommunityNote

// Public reply policy and conversation owner.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetConversationControl = shared.EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetConversationControl

// Lists edit-chain identifiers and the remaining edit window.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetEdit = shared.EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetEdit

// Lists hashtags, symbols, links, and mentions from tweet text.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetEntities = shared.EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetEntities

// Provides hashtag text and source offsets within a tweet.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetEntitiesHashtag = shared.EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetEntitiesHashtag

// Indexed smart-tag, cashtag, or video timestamp metadata.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetEntitiesSmarttag = shared.EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetEntitiesSmarttag

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetEntitiesSmarttagTag = shared.EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetEntitiesSmarttagTag

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetEntitiesSmarttagTagInfo = shared.EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetEntitiesSmarttagTagInfo

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetEntitiesSmarttagTagInfoInfo = shared.EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetEntitiesSmarttagTagInfoInfo

// Indexed smart-tag, cashtag, or video timestamp metadata.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetEntitiesSymbol = shared.EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetEntitiesSymbol

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetEntitiesSymbolTag = shared.EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetEntitiesSymbolTag

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetEntitiesSymbolTagInfo = shared.EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetEntitiesSymbolTagInfo

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetEntitiesSymbolTagInfoInfo = shared.EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetEntitiesSymbolTagInfoInfo

// Indexed smart-tag, cashtag, or video timestamp metadata.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetEntitiesTimestamp = shared.EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetEntitiesTimestamp

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetEntitiesTimestampTag = shared.EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetEntitiesTimestampTag

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetEntitiesTimestampTagInfo = shared.EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetEntitiesTimestampTagInfo

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetEntitiesTimestampTagInfoInfo = shared.EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetEntitiesTimestampTagInfoInfo

// Provides shortened, display, and expanded URLs from tweet text.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetEntitiesURL = shared.EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetEntitiesURL

// Provides profile identity and source offsets for a mention.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetEntitiesUserMention = shared.EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetEntitiesUserMention

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetLimitedAction = shared.EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetLimitedAction

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetLimitedActionPrompt = shared.EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetLimitedActionPrompt

// Complete Note Tweet content and rich-text metadata.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetNoteTweet = shared.EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetNoteTweet

// Lists hashtags, symbols, links, and mentions from tweet text.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetNoteTweetEntities = shared.EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetNoteTweetEntities

// Provides hashtag text and source offsets within a tweet.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetNoteTweetEntitiesHashtag = shared.EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetNoteTweetEntitiesHashtag

// Indexed smart-tag, cashtag, or video timestamp metadata.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetNoteTweetEntitiesSmarttag = shared.EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetNoteTweetEntitiesSmarttag

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetNoteTweetEntitiesSmarttagTag = shared.EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetNoteTweetEntitiesSmarttagTag

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetNoteTweetEntitiesSmarttagTagInfo = shared.EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetNoteTweetEntitiesSmarttagTagInfo

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetNoteTweetEntitiesSmarttagTagInfoInfo = shared.EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetNoteTweetEntitiesSmarttagTagInfoInfo

// Indexed smart-tag, cashtag, or video timestamp metadata.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetNoteTweetEntitiesSymbol = shared.EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetNoteTweetEntitiesSymbol

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetNoteTweetEntitiesSymbolTag = shared.EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetNoteTweetEntitiesSymbolTag

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetNoteTweetEntitiesSymbolTagInfo = shared.EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetNoteTweetEntitiesSymbolTagInfo

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetNoteTweetEntitiesSymbolTagInfoInfo = shared.EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetNoteTweetEntitiesSymbolTagInfoInfo

// Indexed smart-tag, cashtag, or video timestamp metadata.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetNoteTweetEntitiesTimestamp = shared.EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetNoteTweetEntitiesTimestamp

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetNoteTweetEntitiesTimestampTag = shared.EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetNoteTweetEntitiesTimestampTag

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetNoteTweetEntitiesTimestampTagInfo = shared.EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetNoteTweetEntitiesTimestampTagInfo

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetNoteTweetEntitiesTimestampTagInfoInfo = shared.EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetNoteTweetEntitiesTimestampTagInfoInfo

// Provides shortened, display, and expanded URLs from tweet text.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetNoteTweetEntitiesURL = shared.EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetNoteTweetEntitiesURL

// Provides profile identity and source offsets for a mention.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetNoteTweetEntitiesUserMention = shared.EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetNoteTweetEntitiesUserMention

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetNoteTweetInlineMedia = shared.EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetNoteTweetInlineMedia

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetNoteTweetRichtextTag = shared.EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetNoteTweetRichtextTag

// Describes public place metadata on a geotagged tweet.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetPlace = shared.EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetPlace

// Engagement counts retained from a prior tweet edit.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetPreviousCounts = shared.EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetPreviousCounts

// Public post and user referenced by this reaction.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetReactionContext = shared.EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetReactionContext

// Public visibility notice attached to an available tweet.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetTombstone = shared.EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetTombstone

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetTombstoneText = shared.EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetTombstoneText

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetTombstoneTextEntity = shared.EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetTombstoneTextEntity

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetTombstoneTextEntityRef = shared.EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetTombstoneTextEntityRef

// Public post and user referenced by this reaction.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetQuotedTweetReactionContext = shared.EmbeddedTweetQuotedTweetQuotedTweetReactionContext

// Final nested tweet context at depth 4.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweet = shared.EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweet

// Describes an X Article preview and its lifecycle metadata.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetArticle = shared.EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetArticle

// Describes a public card and its referenced profiles.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetCard = shared.EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetCard

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetCardUserReferenceError = shared.EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetCardUserReferenceError

// Community Note presentation metadata returned by X.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetCommunityNote = shared.EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetCommunityNote

// Public reply policy and conversation owner.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetConversationControl = shared.EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetConversationControl

// Lists edit-chain identifiers and the remaining edit window.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetEdit = shared.EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetEdit

// Lists hashtags, symbols, links, and mentions from tweet text.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetEntities = shared.EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetEntities

// Provides hashtag text and source offsets within a tweet.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetEntitiesHashtag = shared.EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetEntitiesHashtag

// Indexed smart-tag, cashtag, or video timestamp metadata.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetEntitiesSmarttag = shared.EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetEntitiesSmarttag

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetEntitiesSmarttagTag = shared.EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetEntitiesSmarttagTag

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetEntitiesSmarttagTagInfo = shared.EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetEntitiesSmarttagTagInfo

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetEntitiesSmarttagTagInfoInfo = shared.EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetEntitiesSmarttagTagInfoInfo

// Indexed smart-tag, cashtag, or video timestamp metadata.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetEntitiesSymbol = shared.EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetEntitiesSymbol

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetEntitiesSymbolTag = shared.EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetEntitiesSymbolTag

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetEntitiesSymbolTagInfo = shared.EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetEntitiesSymbolTagInfo

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetEntitiesSymbolTagInfoInfo = shared.EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetEntitiesSymbolTagInfoInfo

// Indexed smart-tag, cashtag, or video timestamp metadata.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetEntitiesTimestamp = shared.EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetEntitiesTimestamp

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetEntitiesTimestampTag = shared.EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetEntitiesTimestampTag

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetEntitiesTimestampTagInfo = shared.EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetEntitiesTimestampTagInfo

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetEntitiesTimestampTagInfoInfo = shared.EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetEntitiesTimestampTagInfoInfo

// Provides shortened, display, and expanded URLs from tweet text.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetEntitiesURL = shared.EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetEntitiesURL

// Provides profile identity and source offsets for a mention.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetEntitiesUserMention = shared.EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetEntitiesUserMention

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetLimitedAction = shared.EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetLimitedAction

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetLimitedActionPrompt = shared.EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetLimitedActionPrompt

// Complete Note Tweet content and rich-text metadata.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetNoteTweet = shared.EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetNoteTweet

// Lists hashtags, symbols, links, and mentions from tweet text.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetNoteTweetEntities = shared.EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetNoteTweetEntities

// Provides hashtag text and source offsets within a tweet.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesHashtag = shared.EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesHashtag

// Indexed smart-tag, cashtag, or video timestamp metadata.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesSmarttag = shared.EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesSmarttag

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesSmarttagTag = shared.EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesSmarttagTag

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesSmarttagTagInfo = shared.EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesSmarttagTagInfo

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesSmarttagTagInfoInfo = shared.EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesSmarttagTagInfoInfo

// Indexed smart-tag, cashtag, or video timestamp metadata.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesSymbol = shared.EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesSymbol

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesSymbolTag = shared.EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesSymbolTag

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesSymbolTagInfo = shared.EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesSymbolTagInfo

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesSymbolTagInfoInfo = shared.EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesSymbolTagInfoInfo

// Indexed smart-tag, cashtag, or video timestamp metadata.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesTimestamp = shared.EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesTimestamp

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesTimestampTag = shared.EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesTimestampTag

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesTimestampTagInfo = shared.EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesTimestampTagInfo

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesTimestampTagInfoInfo = shared.EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesTimestampTagInfoInfo

// Provides shortened, display, and expanded URLs from tweet text.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesURL = shared.EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesURL

// Provides profile identity and source offsets for a mention.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesUserMention = shared.EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesUserMention

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetNoteTweetInlineMedia = shared.EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetNoteTweetInlineMedia

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetNoteTweetRichtextTag = shared.EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetNoteTweetRichtextTag

// Describes public place metadata on a geotagged tweet.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetPlace = shared.EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetPlace

// Engagement counts retained from a prior tweet edit.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetPreviousCounts = shared.EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetPreviousCounts

// Public post and user referenced by this reaction.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetReactionContext = shared.EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetReactionContext

// Public visibility notice attached to an available tweet.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetTombstone = shared.EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetTombstone

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetTombstoneText = shared.EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetTombstoneText

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetTombstoneTextEntity = shared.EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetTombstoneTextEntity

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetTombstoneTextEntityRef = shared.EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetTombstoneTextEntityRef

// Public visibility notice attached to an available tweet.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetQuotedTweetTombstone = shared.EmbeddedTweetQuotedTweetQuotedTweetTombstone

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetQuotedTweetTombstoneText = shared.EmbeddedTweetQuotedTweetQuotedTweetTombstoneText

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetQuotedTweetTombstoneTextEntity = shared.EmbeddedTweetQuotedTweetQuotedTweetTombstoneTextEntity

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetQuotedTweetTombstoneTextEntityRef = shared.EmbeddedTweetQuotedTweetQuotedTweetTombstoneTextEntityRef

// Public post and user referenced by this reaction.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetReactionContext = shared.EmbeddedTweetQuotedTweetReactionContext

// Nested tweet context at depth 3.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetRetweetedTweet = shared.EmbeddedTweetQuotedTweetRetweetedTweet

// Describes an X Article preview and its lifecycle metadata.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetRetweetedTweetArticle = shared.EmbeddedTweetQuotedTweetRetweetedTweetArticle

// Describes a public card and its referenced profiles.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetRetweetedTweetCard = shared.EmbeddedTweetQuotedTweetRetweetedTweetCard

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetRetweetedTweetCardUserReferenceError = shared.EmbeddedTweetQuotedTweetRetweetedTweetCardUserReferenceError

// Community Note presentation metadata returned by X.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetRetweetedTweetCommunityNote = shared.EmbeddedTweetQuotedTweetRetweetedTweetCommunityNote

// Public reply policy and conversation owner.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetRetweetedTweetConversationControl = shared.EmbeddedTweetQuotedTweetRetweetedTweetConversationControl

// Lists edit-chain identifiers and the remaining edit window.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetRetweetedTweetEdit = shared.EmbeddedTweetQuotedTweetRetweetedTweetEdit

// Lists hashtags, symbols, links, and mentions from tweet text.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetRetweetedTweetEntities = shared.EmbeddedTweetQuotedTweetRetweetedTweetEntities

// Provides hashtag text and source offsets within a tweet.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetRetweetedTweetEntitiesHashtag = shared.EmbeddedTweetQuotedTweetRetweetedTweetEntitiesHashtag

// Indexed smart-tag, cashtag, or video timestamp metadata.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetRetweetedTweetEntitiesSmarttag = shared.EmbeddedTweetQuotedTweetRetweetedTweetEntitiesSmarttag

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetRetweetedTweetEntitiesSmarttagTag = shared.EmbeddedTweetQuotedTweetRetweetedTweetEntitiesSmarttagTag

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetRetweetedTweetEntitiesSmarttagTagInfo = shared.EmbeddedTweetQuotedTweetRetweetedTweetEntitiesSmarttagTagInfo

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetRetweetedTweetEntitiesSmarttagTagInfoInfo = shared.EmbeddedTweetQuotedTweetRetweetedTweetEntitiesSmarttagTagInfoInfo

// Indexed smart-tag, cashtag, or video timestamp metadata.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetRetweetedTweetEntitiesSymbol = shared.EmbeddedTweetQuotedTweetRetweetedTweetEntitiesSymbol

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetRetweetedTweetEntitiesSymbolTag = shared.EmbeddedTweetQuotedTweetRetweetedTweetEntitiesSymbolTag

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetRetweetedTweetEntitiesSymbolTagInfo = shared.EmbeddedTweetQuotedTweetRetweetedTweetEntitiesSymbolTagInfo

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetRetweetedTweetEntitiesSymbolTagInfoInfo = shared.EmbeddedTweetQuotedTweetRetweetedTweetEntitiesSymbolTagInfoInfo

// Indexed smart-tag, cashtag, or video timestamp metadata.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetRetweetedTweetEntitiesTimestamp = shared.EmbeddedTweetQuotedTweetRetweetedTweetEntitiesTimestamp

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetRetweetedTweetEntitiesTimestampTag = shared.EmbeddedTweetQuotedTweetRetweetedTweetEntitiesTimestampTag

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetRetweetedTweetEntitiesTimestampTagInfo = shared.EmbeddedTweetQuotedTweetRetweetedTweetEntitiesTimestampTagInfo

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetRetweetedTweetEntitiesTimestampTagInfoInfo = shared.EmbeddedTweetQuotedTweetRetweetedTweetEntitiesTimestampTagInfoInfo

// Provides shortened, display, and expanded URLs from tweet text.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetRetweetedTweetEntitiesURL = shared.EmbeddedTweetQuotedTweetRetweetedTweetEntitiesURL

// Provides profile identity and source offsets for a mention.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetRetweetedTweetEntitiesUserMention = shared.EmbeddedTweetQuotedTweetRetweetedTweetEntitiesUserMention

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetRetweetedTweetLimitedAction = shared.EmbeddedTweetQuotedTweetRetweetedTweetLimitedAction

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetRetweetedTweetLimitedActionPrompt = shared.EmbeddedTweetQuotedTweetRetweetedTweetLimitedActionPrompt

// Complete Note Tweet content and rich-text metadata.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetRetweetedTweetNoteTweet = shared.EmbeddedTweetQuotedTweetRetweetedTweetNoteTweet

// Lists hashtags, symbols, links, and mentions from tweet text.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetRetweetedTweetNoteTweetEntities = shared.EmbeddedTweetQuotedTweetRetweetedTweetNoteTweetEntities

// Provides hashtag text and source offsets within a tweet.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesHashtag = shared.EmbeddedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesHashtag

// Indexed smart-tag, cashtag, or video timestamp metadata.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesSmarttag = shared.EmbeddedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesSmarttag

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesSmarttagTag = shared.EmbeddedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesSmarttagTag

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesSmarttagTagInfo = shared.EmbeddedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesSmarttagTagInfo

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesSmarttagTagInfoInfo = shared.EmbeddedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesSmarttagTagInfoInfo

// Indexed smart-tag, cashtag, or video timestamp metadata.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesSymbol = shared.EmbeddedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesSymbol

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesSymbolTag = shared.EmbeddedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesSymbolTag

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesSymbolTagInfo = shared.EmbeddedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesSymbolTagInfo

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesSymbolTagInfoInfo = shared.EmbeddedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesSymbolTagInfoInfo

// Indexed smart-tag, cashtag, or video timestamp metadata.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesTimestamp = shared.EmbeddedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesTimestamp

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesTimestampTag = shared.EmbeddedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesTimestampTag

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesTimestampTagInfo = shared.EmbeddedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesTimestampTagInfo

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesTimestampTagInfoInfo = shared.EmbeddedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesTimestampTagInfoInfo

// Provides shortened, display, and expanded URLs from tweet text.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesURL = shared.EmbeddedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesURL

// Provides profile identity and source offsets for a mention.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesUserMention = shared.EmbeddedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesUserMention

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetRetweetedTweetNoteTweetInlineMedia = shared.EmbeddedTweetQuotedTweetRetweetedTweetNoteTweetInlineMedia

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetRetweetedTweetNoteTweetRichtextTag = shared.EmbeddedTweetQuotedTweetRetweetedTweetNoteTweetRichtextTag

// Describes public place metadata on a geotagged tweet.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetRetweetedTweetPlace = shared.EmbeddedTweetQuotedTweetRetweetedTweetPlace

// Engagement counts retained from a prior tweet edit.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetRetweetedTweetPreviousCounts = shared.EmbeddedTweetQuotedTweetRetweetedTweetPreviousCounts

// Final nested tweet context at depth 4.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweet = shared.EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweet

// Describes an X Article preview and its lifecycle metadata.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetArticle = shared.EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetArticle

// Describes a public card and its referenced profiles.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetCard = shared.EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetCard

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetCardUserReferenceError = shared.EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetCardUserReferenceError

// Community Note presentation metadata returned by X.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetCommunityNote = shared.EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetCommunityNote

// Public reply policy and conversation owner.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetConversationControl = shared.EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetConversationControl

// Lists edit-chain identifiers and the remaining edit window.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetEdit = shared.EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetEdit

// Lists hashtags, symbols, links, and mentions from tweet text.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetEntities = shared.EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetEntities

// Provides hashtag text and source offsets within a tweet.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetEntitiesHashtag = shared.EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetEntitiesHashtag

// Indexed smart-tag, cashtag, or video timestamp metadata.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetEntitiesSmarttag = shared.EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetEntitiesSmarttag

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetEntitiesSmarttagTag = shared.EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetEntitiesSmarttagTag

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetEntitiesSmarttagTagInfo = shared.EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetEntitiesSmarttagTagInfo

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetEntitiesSmarttagTagInfoInfo = shared.EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetEntitiesSmarttagTagInfoInfo

// Indexed smart-tag, cashtag, or video timestamp metadata.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetEntitiesSymbol = shared.EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetEntitiesSymbol

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetEntitiesSymbolTag = shared.EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetEntitiesSymbolTag

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetEntitiesSymbolTagInfo = shared.EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetEntitiesSymbolTagInfo

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetEntitiesSymbolTagInfoInfo = shared.EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetEntitiesSymbolTagInfoInfo

// Indexed smart-tag, cashtag, or video timestamp metadata.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetEntitiesTimestamp = shared.EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetEntitiesTimestamp

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetEntitiesTimestampTag = shared.EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetEntitiesTimestampTag

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetEntitiesTimestampTagInfo = shared.EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetEntitiesTimestampTagInfo

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetEntitiesTimestampTagInfoInfo = shared.EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetEntitiesTimestampTagInfoInfo

// Provides shortened, display, and expanded URLs from tweet text.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetEntitiesURL = shared.EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetEntitiesURL

// Provides profile identity and source offsets for a mention.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetEntitiesUserMention = shared.EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetEntitiesUserMention

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetLimitedAction = shared.EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetLimitedAction

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetLimitedActionPrompt = shared.EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetLimitedActionPrompt

// Complete Note Tweet content and rich-text metadata.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetNoteTweet = shared.EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetNoteTweet

// Lists hashtags, symbols, links, and mentions from tweet text.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetNoteTweetEntities = shared.EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetNoteTweetEntities

// Provides hashtag text and source offsets within a tweet.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesHashtag = shared.EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesHashtag

// Indexed smart-tag, cashtag, or video timestamp metadata.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesSmarttag = shared.EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesSmarttag

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesSmarttagTag = shared.EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesSmarttagTag

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesSmarttagTagInfo = shared.EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesSmarttagTagInfo

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesSmarttagTagInfoInfo = shared.EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesSmarttagTagInfoInfo

// Indexed smart-tag, cashtag, or video timestamp metadata.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesSymbol = shared.EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesSymbol

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesSymbolTag = shared.EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesSymbolTag

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesSymbolTagInfo = shared.EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesSymbolTagInfo

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesSymbolTagInfoInfo = shared.EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesSymbolTagInfoInfo

// Indexed smart-tag, cashtag, or video timestamp metadata.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesTimestamp = shared.EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesTimestamp

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesTimestampTag = shared.EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesTimestampTag

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesTimestampTagInfo = shared.EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesTimestampTagInfo

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesTimestampTagInfoInfo = shared.EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesTimestampTagInfoInfo

// Provides shortened, display, and expanded URLs from tweet text.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesURL = shared.EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesURL

// Provides profile identity and source offsets for a mention.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesUserMention = shared.EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesUserMention

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetNoteTweetInlineMedia = shared.EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetNoteTweetInlineMedia

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetNoteTweetRichtextTag = shared.EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetNoteTweetRichtextTag

// Describes public place metadata on a geotagged tweet.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetPlace = shared.EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetPlace

// Engagement counts retained from a prior tweet edit.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetPreviousCounts = shared.EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetPreviousCounts

// Public post and user referenced by this reaction.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetReactionContext = shared.EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetReactionContext

// Public visibility notice attached to an available tweet.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetTombstone = shared.EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetTombstone

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetTombstoneText = shared.EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetTombstoneText

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetTombstoneTextEntity = shared.EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetTombstoneTextEntity

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetTombstoneTextEntityRef = shared.EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetTombstoneTextEntityRef

// Public post and user referenced by this reaction.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetRetweetedTweetReactionContext = shared.EmbeddedTweetQuotedTweetRetweetedTweetReactionContext

// Final nested tweet context at depth 4.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweet = shared.EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweet

// Describes an X Article preview and its lifecycle metadata.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetArticle = shared.EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetArticle

// Describes a public card and its referenced profiles.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetCard = shared.EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetCard

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetCardUserReferenceError = shared.EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetCardUserReferenceError

// Community Note presentation metadata returned by X.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetCommunityNote = shared.EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetCommunityNote

// Public reply policy and conversation owner.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetConversationControl = shared.EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetConversationControl

// Lists edit-chain identifiers and the remaining edit window.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetEdit = shared.EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetEdit

// Lists hashtags, symbols, links, and mentions from tweet text.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetEntities = shared.EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetEntities

// Provides hashtag text and source offsets within a tweet.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetEntitiesHashtag = shared.EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetEntitiesHashtag

// Indexed smart-tag, cashtag, or video timestamp metadata.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetEntitiesSmarttag = shared.EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetEntitiesSmarttag

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetEntitiesSmarttagTag = shared.EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetEntitiesSmarttagTag

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetEntitiesSmarttagTagInfo = shared.EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetEntitiesSmarttagTagInfo

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetEntitiesSmarttagTagInfoInfo = shared.EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetEntitiesSmarttagTagInfoInfo

// Indexed smart-tag, cashtag, or video timestamp metadata.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetEntitiesSymbol = shared.EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetEntitiesSymbol

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetEntitiesSymbolTag = shared.EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetEntitiesSymbolTag

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetEntitiesSymbolTagInfo = shared.EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetEntitiesSymbolTagInfo

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetEntitiesSymbolTagInfoInfo = shared.EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetEntitiesSymbolTagInfoInfo

// Indexed smart-tag, cashtag, or video timestamp metadata.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetEntitiesTimestamp = shared.EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetEntitiesTimestamp

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetEntitiesTimestampTag = shared.EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetEntitiesTimestampTag

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetEntitiesTimestampTagInfo = shared.EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetEntitiesTimestampTagInfo

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetEntitiesTimestampTagInfoInfo = shared.EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetEntitiesTimestampTagInfoInfo

// Provides shortened, display, and expanded URLs from tweet text.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetEntitiesURL = shared.EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetEntitiesURL

// Provides profile identity and source offsets for a mention.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetEntitiesUserMention = shared.EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetEntitiesUserMention

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetLimitedAction = shared.EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetLimitedAction

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetLimitedActionPrompt = shared.EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetLimitedActionPrompt

// Complete Note Tweet content and rich-text metadata.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetNoteTweet = shared.EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetNoteTweet

// Lists hashtags, symbols, links, and mentions from tweet text.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetNoteTweetEntities = shared.EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetNoteTweetEntities

// Provides hashtag text and source offsets within a tweet.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesHashtag = shared.EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesHashtag

// Indexed smart-tag, cashtag, or video timestamp metadata.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesSmarttag = shared.EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesSmarttag

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesSmarttagTag = shared.EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesSmarttagTag

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesSmarttagTagInfo = shared.EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesSmarttagTagInfo

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesSmarttagTagInfoInfo = shared.EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesSmarttagTagInfoInfo

// Indexed smart-tag, cashtag, or video timestamp metadata.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesSymbol = shared.EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesSymbol

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesSymbolTag = shared.EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesSymbolTag

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesSymbolTagInfo = shared.EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesSymbolTagInfo

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesSymbolTagInfoInfo = shared.EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesSymbolTagInfoInfo

// Indexed smart-tag, cashtag, or video timestamp metadata.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesTimestamp = shared.EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesTimestamp

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesTimestampTag = shared.EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesTimestampTag

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesTimestampTagInfo = shared.EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesTimestampTagInfo

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesTimestampTagInfoInfo = shared.EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesTimestampTagInfoInfo

// Provides shortened, display, and expanded URLs from tweet text.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesURL = shared.EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesURL

// Provides profile identity and source offsets for a mention.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesUserMention = shared.EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesUserMention

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetNoteTweetInlineMedia = shared.EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetNoteTweetInlineMedia

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetNoteTweetRichtextTag = shared.EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetNoteTweetRichtextTag

// Describes public place metadata on a geotagged tweet.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetPlace = shared.EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetPlace

// Engagement counts retained from a prior tweet edit.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetPreviousCounts = shared.EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetPreviousCounts

// Public post and user referenced by this reaction.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetReactionContext = shared.EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetReactionContext

// Public visibility notice attached to an available tweet.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetTombstone = shared.EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetTombstone

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetTombstoneText = shared.EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetTombstoneText

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetTombstoneTextEntity = shared.EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetTombstoneTextEntity

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetTombstoneTextEntityRef = shared.EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetTombstoneTextEntityRef

// Public visibility notice attached to an available tweet.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetRetweetedTweetTombstone = shared.EmbeddedTweetQuotedTweetRetweetedTweetTombstone

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetRetweetedTweetTombstoneText = shared.EmbeddedTweetQuotedTweetRetweetedTweetTombstoneText

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetRetweetedTweetTombstoneTextEntity = shared.EmbeddedTweetQuotedTweetRetweetedTweetTombstoneTextEntity

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetRetweetedTweetTombstoneTextEntityRef = shared.EmbeddedTweetQuotedTweetRetweetedTweetTombstoneTextEntityRef

// Public visibility notice attached to an available tweet.
//
// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetTombstone = shared.EmbeddedTweetQuotedTweetTombstone

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetTombstoneText = shared.EmbeddedTweetQuotedTweetTombstoneText

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetTombstoneTextEntity = shared.EmbeddedTweetQuotedTweetTombstoneTextEntity

// This is an alias to an internal type.
type EmbeddedTweetQuotedTweetTombstoneTextEntityRef = shared.EmbeddedTweetQuotedTweetTombstoneTextEntityRef

// Public post and user referenced by this reaction.
//
// This is an alias to an internal type.
type EmbeddedTweetReactionContext = shared.EmbeddedTweetReactionContext

// Nested tweet context at depth 2.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweet = shared.EmbeddedTweetRetweetedTweet

// Describes an X Article preview and its lifecycle metadata.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetArticle = shared.EmbeddedTweetRetweetedTweetArticle

// Describes a public card and its referenced profiles.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetCard = shared.EmbeddedTweetRetweetedTweetCard

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetCardUserReferenceError = shared.EmbeddedTweetRetweetedTweetCardUserReferenceError

// Community Note presentation metadata returned by X.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetCommunityNote = shared.EmbeddedTweetRetweetedTweetCommunityNote

// Public reply policy and conversation owner.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetConversationControl = shared.EmbeddedTweetRetweetedTweetConversationControl

// Lists edit-chain identifiers and the remaining edit window.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetEdit = shared.EmbeddedTweetRetweetedTweetEdit

// Lists hashtags, symbols, links, and mentions from tweet text.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetEntities = shared.EmbeddedTweetRetweetedTweetEntities

// Provides hashtag text and source offsets within a tweet.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetEntitiesHashtag = shared.EmbeddedTweetRetweetedTweetEntitiesHashtag

// Indexed smart-tag, cashtag, or video timestamp metadata.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetEntitiesSmarttag = shared.EmbeddedTweetRetweetedTweetEntitiesSmarttag

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetEntitiesSmarttagTag = shared.EmbeddedTweetRetweetedTweetEntitiesSmarttagTag

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetEntitiesSmarttagTagInfo = shared.EmbeddedTweetRetweetedTweetEntitiesSmarttagTagInfo

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetEntitiesSmarttagTagInfoInfo = shared.EmbeddedTweetRetweetedTweetEntitiesSmarttagTagInfoInfo

// Indexed smart-tag, cashtag, or video timestamp metadata.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetEntitiesSymbol = shared.EmbeddedTweetRetweetedTweetEntitiesSymbol

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetEntitiesSymbolTag = shared.EmbeddedTweetRetweetedTweetEntitiesSymbolTag

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetEntitiesSymbolTagInfo = shared.EmbeddedTweetRetweetedTweetEntitiesSymbolTagInfo

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetEntitiesSymbolTagInfoInfo = shared.EmbeddedTweetRetweetedTweetEntitiesSymbolTagInfoInfo

// Indexed smart-tag, cashtag, or video timestamp metadata.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetEntitiesTimestamp = shared.EmbeddedTweetRetweetedTweetEntitiesTimestamp

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetEntitiesTimestampTag = shared.EmbeddedTweetRetweetedTweetEntitiesTimestampTag

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetEntitiesTimestampTagInfo = shared.EmbeddedTweetRetweetedTweetEntitiesTimestampTagInfo

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetEntitiesTimestampTagInfoInfo = shared.EmbeddedTweetRetweetedTweetEntitiesTimestampTagInfoInfo

// Provides shortened, display, and expanded URLs from tweet text.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetEntitiesURL = shared.EmbeddedTweetRetweetedTweetEntitiesURL

// Provides profile identity and source offsets for a mention.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetEntitiesUserMention = shared.EmbeddedTweetRetweetedTweetEntitiesUserMention

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetLimitedAction = shared.EmbeddedTweetRetweetedTweetLimitedAction

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetLimitedActionPrompt = shared.EmbeddedTweetRetweetedTweetLimitedActionPrompt

// Complete Note Tweet content and rich-text metadata.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetNoteTweet = shared.EmbeddedTweetRetweetedTweetNoteTweet

// Lists hashtags, symbols, links, and mentions from tweet text.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetNoteTweetEntities = shared.EmbeddedTweetRetweetedTweetNoteTweetEntities

// Provides hashtag text and source offsets within a tweet.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetNoteTweetEntitiesHashtag = shared.EmbeddedTweetRetweetedTweetNoteTweetEntitiesHashtag

// Indexed smart-tag, cashtag, or video timestamp metadata.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetNoteTweetEntitiesSmarttag = shared.EmbeddedTweetRetweetedTweetNoteTweetEntitiesSmarttag

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetNoteTweetEntitiesSmarttagTag = shared.EmbeddedTweetRetweetedTweetNoteTweetEntitiesSmarttagTag

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetNoteTweetEntitiesSmarttagTagInfo = shared.EmbeddedTweetRetweetedTweetNoteTweetEntitiesSmarttagTagInfo

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetNoteTweetEntitiesSmarttagTagInfoInfo = shared.EmbeddedTweetRetweetedTweetNoteTweetEntitiesSmarttagTagInfoInfo

// Indexed smart-tag, cashtag, or video timestamp metadata.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetNoteTweetEntitiesSymbol = shared.EmbeddedTweetRetweetedTweetNoteTweetEntitiesSymbol

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetNoteTweetEntitiesSymbolTag = shared.EmbeddedTweetRetweetedTweetNoteTweetEntitiesSymbolTag

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetNoteTweetEntitiesSymbolTagInfo = shared.EmbeddedTweetRetweetedTweetNoteTweetEntitiesSymbolTagInfo

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetNoteTweetEntitiesSymbolTagInfoInfo = shared.EmbeddedTweetRetweetedTweetNoteTweetEntitiesSymbolTagInfoInfo

// Indexed smart-tag, cashtag, or video timestamp metadata.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetNoteTweetEntitiesTimestamp = shared.EmbeddedTweetRetweetedTweetNoteTweetEntitiesTimestamp

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetNoteTweetEntitiesTimestampTag = shared.EmbeddedTweetRetweetedTweetNoteTweetEntitiesTimestampTag

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetNoteTweetEntitiesTimestampTagInfo = shared.EmbeddedTweetRetweetedTweetNoteTweetEntitiesTimestampTagInfo

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetNoteTweetEntitiesTimestampTagInfoInfo = shared.EmbeddedTweetRetweetedTweetNoteTweetEntitiesTimestampTagInfoInfo

// Provides shortened, display, and expanded URLs from tweet text.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetNoteTweetEntitiesURL = shared.EmbeddedTweetRetweetedTweetNoteTweetEntitiesURL

// Provides profile identity and source offsets for a mention.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetNoteTweetEntitiesUserMention = shared.EmbeddedTweetRetweetedTweetNoteTweetEntitiesUserMention

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetNoteTweetInlineMedia = shared.EmbeddedTweetRetweetedTweetNoteTweetInlineMedia

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetNoteTweetRichtextTag = shared.EmbeddedTweetRetweetedTweetNoteTweetRichtextTag

// Describes public place metadata on a geotagged tweet.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetPlace = shared.EmbeddedTweetRetweetedTweetPlace

// Engagement counts retained from a prior tweet edit.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetPreviousCounts = shared.EmbeddedTweetRetweetedTweetPreviousCounts

// Nested tweet context at depth 3.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetQuotedTweet = shared.EmbeddedTweetRetweetedTweetQuotedTweet

// Describes an X Article preview and its lifecycle metadata.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetQuotedTweetArticle = shared.EmbeddedTweetRetweetedTweetQuotedTweetArticle

// Describes a public card and its referenced profiles.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetQuotedTweetCard = shared.EmbeddedTweetRetweetedTweetQuotedTweetCard

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetQuotedTweetCardUserReferenceError = shared.EmbeddedTweetRetweetedTweetQuotedTweetCardUserReferenceError

// Community Note presentation metadata returned by X.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetQuotedTweetCommunityNote = shared.EmbeddedTweetRetweetedTweetQuotedTweetCommunityNote

// Public reply policy and conversation owner.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetQuotedTweetConversationControl = shared.EmbeddedTweetRetweetedTweetQuotedTweetConversationControl

// Lists edit-chain identifiers and the remaining edit window.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetQuotedTweetEdit = shared.EmbeddedTweetRetweetedTweetQuotedTweetEdit

// Lists hashtags, symbols, links, and mentions from tweet text.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetQuotedTweetEntities = shared.EmbeddedTweetRetweetedTweetQuotedTweetEntities

// Provides hashtag text and source offsets within a tweet.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetQuotedTweetEntitiesHashtag = shared.EmbeddedTweetRetweetedTweetQuotedTweetEntitiesHashtag

// Indexed smart-tag, cashtag, or video timestamp metadata.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetQuotedTweetEntitiesSmarttag = shared.EmbeddedTweetRetweetedTweetQuotedTweetEntitiesSmarttag

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetQuotedTweetEntitiesSmarttagTag = shared.EmbeddedTweetRetweetedTweetQuotedTweetEntitiesSmarttagTag

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetQuotedTweetEntitiesSmarttagTagInfo = shared.EmbeddedTweetRetweetedTweetQuotedTweetEntitiesSmarttagTagInfo

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetQuotedTweetEntitiesSmarttagTagInfoInfo = shared.EmbeddedTweetRetweetedTweetQuotedTweetEntitiesSmarttagTagInfoInfo

// Indexed smart-tag, cashtag, or video timestamp metadata.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetQuotedTweetEntitiesSymbol = shared.EmbeddedTweetRetweetedTweetQuotedTweetEntitiesSymbol

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetQuotedTweetEntitiesSymbolTag = shared.EmbeddedTweetRetweetedTweetQuotedTweetEntitiesSymbolTag

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetQuotedTweetEntitiesSymbolTagInfo = shared.EmbeddedTweetRetweetedTweetQuotedTweetEntitiesSymbolTagInfo

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetQuotedTweetEntitiesSymbolTagInfoInfo = shared.EmbeddedTweetRetweetedTweetQuotedTweetEntitiesSymbolTagInfoInfo

// Indexed smart-tag, cashtag, or video timestamp metadata.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetQuotedTweetEntitiesTimestamp = shared.EmbeddedTweetRetweetedTweetQuotedTweetEntitiesTimestamp

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetQuotedTweetEntitiesTimestampTag = shared.EmbeddedTweetRetweetedTweetQuotedTweetEntitiesTimestampTag

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetQuotedTweetEntitiesTimestampTagInfo = shared.EmbeddedTweetRetweetedTweetQuotedTweetEntitiesTimestampTagInfo

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetQuotedTweetEntitiesTimestampTagInfoInfo = shared.EmbeddedTweetRetweetedTweetQuotedTweetEntitiesTimestampTagInfoInfo

// Provides shortened, display, and expanded URLs from tweet text.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetQuotedTweetEntitiesURL = shared.EmbeddedTweetRetweetedTweetQuotedTweetEntitiesURL

// Provides profile identity and source offsets for a mention.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetQuotedTweetEntitiesUserMention = shared.EmbeddedTweetRetweetedTweetQuotedTweetEntitiesUserMention

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetQuotedTweetLimitedAction = shared.EmbeddedTweetRetweetedTweetQuotedTweetLimitedAction

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetQuotedTweetLimitedActionPrompt = shared.EmbeddedTweetRetweetedTweetQuotedTweetLimitedActionPrompt

// Complete Note Tweet content and rich-text metadata.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetQuotedTweetNoteTweet = shared.EmbeddedTweetRetweetedTweetQuotedTweetNoteTweet

// Lists hashtags, symbols, links, and mentions from tweet text.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetQuotedTweetNoteTweetEntities = shared.EmbeddedTweetRetweetedTweetQuotedTweetNoteTweetEntities

// Provides hashtag text and source offsets within a tweet.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesHashtag = shared.EmbeddedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesHashtag

// Indexed smart-tag, cashtag, or video timestamp metadata.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesSmarttag = shared.EmbeddedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesSmarttag

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesSmarttagTag = shared.EmbeddedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesSmarttagTag

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesSmarttagTagInfo = shared.EmbeddedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesSmarttagTagInfo

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesSmarttagTagInfoInfo = shared.EmbeddedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesSmarttagTagInfoInfo

// Indexed smart-tag, cashtag, or video timestamp metadata.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesSymbol = shared.EmbeddedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesSymbol

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesSymbolTag = shared.EmbeddedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesSymbolTag

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesSymbolTagInfo = shared.EmbeddedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesSymbolTagInfo

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesSymbolTagInfoInfo = shared.EmbeddedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesSymbolTagInfoInfo

// Indexed smart-tag, cashtag, or video timestamp metadata.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesTimestamp = shared.EmbeddedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesTimestamp

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesTimestampTag = shared.EmbeddedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesTimestampTag

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesTimestampTagInfo = shared.EmbeddedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesTimestampTagInfo

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesTimestampTagInfoInfo = shared.EmbeddedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesTimestampTagInfoInfo

// Provides shortened, display, and expanded URLs from tweet text.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesURL = shared.EmbeddedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesURL

// Provides profile identity and source offsets for a mention.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesUserMention = shared.EmbeddedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesUserMention

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetQuotedTweetNoteTweetInlineMedia = shared.EmbeddedTweetRetweetedTweetQuotedTweetNoteTweetInlineMedia

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetQuotedTweetNoteTweetRichtextTag = shared.EmbeddedTweetRetweetedTweetQuotedTweetNoteTweetRichtextTag

// Describes public place metadata on a geotagged tweet.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetQuotedTweetPlace = shared.EmbeddedTweetRetweetedTweetQuotedTweetPlace

// Engagement counts retained from a prior tweet edit.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetQuotedTweetPreviousCounts = shared.EmbeddedTweetRetweetedTweetQuotedTweetPreviousCounts

// Final nested tweet context at depth 4.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweet = shared.EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweet

// Describes an X Article preview and its lifecycle metadata.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetArticle = shared.EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetArticle

// Describes a public card and its referenced profiles.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetCard = shared.EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetCard

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetCardUserReferenceError = shared.EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetCardUserReferenceError

// Community Note presentation metadata returned by X.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetCommunityNote = shared.EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetCommunityNote

// Public reply policy and conversation owner.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetConversationControl = shared.EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetConversationControl

// Lists edit-chain identifiers and the remaining edit window.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetEdit = shared.EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetEdit

// Lists hashtags, symbols, links, and mentions from tweet text.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetEntities = shared.EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetEntities

// Provides hashtag text and source offsets within a tweet.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetEntitiesHashtag = shared.EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetEntitiesHashtag

// Indexed smart-tag, cashtag, or video timestamp metadata.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetEntitiesSmarttag = shared.EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetEntitiesSmarttag

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetEntitiesSmarttagTag = shared.EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetEntitiesSmarttagTag

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetEntitiesSmarttagTagInfo = shared.EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetEntitiesSmarttagTagInfo

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetEntitiesSmarttagTagInfoInfo = shared.EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetEntitiesSmarttagTagInfoInfo

// Indexed smart-tag, cashtag, or video timestamp metadata.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetEntitiesSymbol = shared.EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetEntitiesSymbol

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetEntitiesSymbolTag = shared.EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetEntitiesSymbolTag

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetEntitiesSymbolTagInfo = shared.EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetEntitiesSymbolTagInfo

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetEntitiesSymbolTagInfoInfo = shared.EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetEntitiesSymbolTagInfoInfo

// Indexed smart-tag, cashtag, or video timestamp metadata.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetEntitiesTimestamp = shared.EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetEntitiesTimestamp

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetEntitiesTimestampTag = shared.EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetEntitiesTimestampTag

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetEntitiesTimestampTagInfo = shared.EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetEntitiesTimestampTagInfo

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetEntitiesTimestampTagInfoInfo = shared.EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetEntitiesTimestampTagInfoInfo

// Provides shortened, display, and expanded URLs from tweet text.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetEntitiesURL = shared.EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetEntitiesURL

// Provides profile identity and source offsets for a mention.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetEntitiesUserMention = shared.EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetEntitiesUserMention

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetLimitedAction = shared.EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetLimitedAction

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetLimitedActionPrompt = shared.EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetLimitedActionPrompt

// Complete Note Tweet content and rich-text metadata.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetNoteTweet = shared.EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetNoteTweet

// Lists hashtags, symbols, links, and mentions from tweet text.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetNoteTweetEntities = shared.EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetNoteTweetEntities

// Provides hashtag text and source offsets within a tweet.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetNoteTweetEntitiesHashtag = shared.EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetNoteTweetEntitiesHashtag

// Indexed smart-tag, cashtag, or video timestamp metadata.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetNoteTweetEntitiesSmarttag = shared.EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetNoteTweetEntitiesSmarttag

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetNoteTweetEntitiesSmarttagTag = shared.EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetNoteTweetEntitiesSmarttagTag

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetNoteTweetEntitiesSmarttagTagInfo = shared.EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetNoteTweetEntitiesSmarttagTagInfo

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetNoteTweetEntitiesSmarttagTagInfoInfo = shared.EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetNoteTweetEntitiesSmarttagTagInfoInfo

// Indexed smart-tag, cashtag, or video timestamp metadata.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetNoteTweetEntitiesSymbol = shared.EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetNoteTweetEntitiesSymbol

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetNoteTweetEntitiesSymbolTag = shared.EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetNoteTweetEntitiesSymbolTag

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetNoteTweetEntitiesSymbolTagInfo = shared.EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetNoteTweetEntitiesSymbolTagInfo

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetNoteTweetEntitiesSymbolTagInfoInfo = shared.EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetNoteTweetEntitiesSymbolTagInfoInfo

// Indexed smart-tag, cashtag, or video timestamp metadata.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetNoteTweetEntitiesTimestamp = shared.EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetNoteTweetEntitiesTimestamp

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetNoteTweetEntitiesTimestampTag = shared.EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetNoteTweetEntitiesTimestampTag

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetNoteTweetEntitiesTimestampTagInfo = shared.EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetNoteTweetEntitiesTimestampTagInfo

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetNoteTweetEntitiesTimestampTagInfoInfo = shared.EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetNoteTweetEntitiesTimestampTagInfoInfo

// Provides shortened, display, and expanded URLs from tweet text.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetNoteTweetEntitiesURL = shared.EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetNoteTweetEntitiesURL

// Provides profile identity and source offsets for a mention.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetNoteTweetEntitiesUserMention = shared.EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetNoteTweetEntitiesUserMention

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetNoteTweetInlineMedia = shared.EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetNoteTweetInlineMedia

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetNoteTweetRichtextTag = shared.EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetNoteTweetRichtextTag

// Describes public place metadata on a geotagged tweet.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetPlace = shared.EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetPlace

// Engagement counts retained from a prior tweet edit.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetPreviousCounts = shared.EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetPreviousCounts

// Public post and user referenced by this reaction.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetReactionContext = shared.EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetReactionContext

// Public visibility notice attached to an available tweet.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetTombstone = shared.EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetTombstone

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetTombstoneText = shared.EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetTombstoneText

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetTombstoneTextEntity = shared.EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetTombstoneTextEntity

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetTombstoneTextEntityRef = shared.EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetTombstoneTextEntityRef

// Public post and user referenced by this reaction.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetQuotedTweetReactionContext = shared.EmbeddedTweetRetweetedTweetQuotedTweetReactionContext

// Final nested tweet context at depth 4.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweet = shared.EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweet

// Describes an X Article preview and its lifecycle metadata.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetArticle = shared.EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetArticle

// Describes a public card and its referenced profiles.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetCard = shared.EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetCard

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetCardUserReferenceError = shared.EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetCardUserReferenceError

// Community Note presentation metadata returned by X.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetCommunityNote = shared.EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetCommunityNote

// Public reply policy and conversation owner.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetConversationControl = shared.EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetConversationControl

// Lists edit-chain identifiers and the remaining edit window.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetEdit = shared.EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetEdit

// Lists hashtags, symbols, links, and mentions from tweet text.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetEntities = shared.EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetEntities

// Provides hashtag text and source offsets within a tweet.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetEntitiesHashtag = shared.EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetEntitiesHashtag

// Indexed smart-tag, cashtag, or video timestamp metadata.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetEntitiesSmarttag = shared.EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetEntitiesSmarttag

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetEntitiesSmarttagTag = shared.EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetEntitiesSmarttagTag

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetEntitiesSmarttagTagInfo = shared.EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetEntitiesSmarttagTagInfo

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetEntitiesSmarttagTagInfoInfo = shared.EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetEntitiesSmarttagTagInfoInfo

// Indexed smart-tag, cashtag, or video timestamp metadata.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetEntitiesSymbol = shared.EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetEntitiesSymbol

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetEntitiesSymbolTag = shared.EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetEntitiesSymbolTag

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetEntitiesSymbolTagInfo = shared.EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetEntitiesSymbolTagInfo

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetEntitiesSymbolTagInfoInfo = shared.EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetEntitiesSymbolTagInfoInfo

// Indexed smart-tag, cashtag, or video timestamp metadata.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetEntitiesTimestamp = shared.EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetEntitiesTimestamp

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetEntitiesTimestampTag = shared.EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetEntitiesTimestampTag

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetEntitiesTimestampTagInfo = shared.EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetEntitiesTimestampTagInfo

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetEntitiesTimestampTagInfoInfo = shared.EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetEntitiesTimestampTagInfoInfo

// Provides shortened, display, and expanded URLs from tweet text.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetEntitiesURL = shared.EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetEntitiesURL

// Provides profile identity and source offsets for a mention.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetEntitiesUserMention = shared.EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetEntitiesUserMention

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetLimitedAction = shared.EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetLimitedAction

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetLimitedActionPrompt = shared.EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetLimitedActionPrompt

// Complete Note Tweet content and rich-text metadata.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetNoteTweet = shared.EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetNoteTweet

// Lists hashtags, symbols, links, and mentions from tweet text.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetNoteTweetEntities = shared.EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetNoteTweetEntities

// Provides hashtag text and source offsets within a tweet.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesHashtag = shared.EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesHashtag

// Indexed smart-tag, cashtag, or video timestamp metadata.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesSmarttag = shared.EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesSmarttag

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesSmarttagTag = shared.EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesSmarttagTag

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesSmarttagTagInfo = shared.EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesSmarttagTagInfo

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesSmarttagTagInfoInfo = shared.EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesSmarttagTagInfoInfo

// Indexed smart-tag, cashtag, or video timestamp metadata.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesSymbol = shared.EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesSymbol

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesSymbolTag = shared.EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesSymbolTag

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesSymbolTagInfo = shared.EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesSymbolTagInfo

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesSymbolTagInfoInfo = shared.EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesSymbolTagInfoInfo

// Indexed smart-tag, cashtag, or video timestamp metadata.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesTimestamp = shared.EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesTimestamp

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesTimestampTag = shared.EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesTimestampTag

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesTimestampTagInfo = shared.EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesTimestampTagInfo

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesTimestampTagInfoInfo = shared.EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesTimestampTagInfoInfo

// Provides shortened, display, and expanded URLs from tweet text.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesURL = shared.EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesURL

// Provides profile identity and source offsets for a mention.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesUserMention = shared.EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetNoteTweetEntitiesUserMention

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetNoteTweetInlineMedia = shared.EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetNoteTweetInlineMedia

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetNoteTweetRichtextTag = shared.EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetNoteTweetRichtextTag

// Describes public place metadata on a geotagged tweet.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetPlace = shared.EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetPlace

// Engagement counts retained from a prior tweet edit.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetPreviousCounts = shared.EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetPreviousCounts

// Public post and user referenced by this reaction.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetReactionContext = shared.EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetReactionContext

// Public visibility notice attached to an available tweet.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetTombstone = shared.EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetTombstone

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetTombstoneText = shared.EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetTombstoneText

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetTombstoneTextEntity = shared.EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetTombstoneTextEntity

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetTombstoneTextEntityRef = shared.EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetTombstoneTextEntityRef

// Public visibility notice attached to an available tweet.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetQuotedTweetTombstone = shared.EmbeddedTweetRetweetedTweetQuotedTweetTombstone

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetQuotedTweetTombstoneText = shared.EmbeddedTweetRetweetedTweetQuotedTweetTombstoneText

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetQuotedTweetTombstoneTextEntity = shared.EmbeddedTweetRetweetedTweetQuotedTweetTombstoneTextEntity

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetQuotedTweetTombstoneTextEntityRef = shared.EmbeddedTweetRetweetedTweetQuotedTweetTombstoneTextEntityRef

// Public post and user referenced by this reaction.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetReactionContext = shared.EmbeddedTweetRetweetedTweetReactionContext

// Nested tweet context at depth 3.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetRetweetedTweet = shared.EmbeddedTweetRetweetedTweetRetweetedTweet

// Describes an X Article preview and its lifecycle metadata.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetRetweetedTweetArticle = shared.EmbeddedTweetRetweetedTweetRetweetedTweetArticle

// Describes a public card and its referenced profiles.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetRetweetedTweetCard = shared.EmbeddedTweetRetweetedTweetRetweetedTweetCard

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetRetweetedTweetCardUserReferenceError = shared.EmbeddedTweetRetweetedTweetRetweetedTweetCardUserReferenceError

// Community Note presentation metadata returned by X.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetRetweetedTweetCommunityNote = shared.EmbeddedTweetRetweetedTweetRetweetedTweetCommunityNote

// Public reply policy and conversation owner.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetRetweetedTweetConversationControl = shared.EmbeddedTweetRetweetedTweetRetweetedTweetConversationControl

// Lists edit-chain identifiers and the remaining edit window.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetRetweetedTweetEdit = shared.EmbeddedTweetRetweetedTweetRetweetedTweetEdit

// Lists hashtags, symbols, links, and mentions from tweet text.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetRetweetedTweetEntities = shared.EmbeddedTweetRetweetedTweetRetweetedTweetEntities

// Provides hashtag text and source offsets within a tweet.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetRetweetedTweetEntitiesHashtag = shared.EmbeddedTweetRetweetedTweetRetweetedTweetEntitiesHashtag

// Indexed smart-tag, cashtag, or video timestamp metadata.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetRetweetedTweetEntitiesSmarttag = shared.EmbeddedTweetRetweetedTweetRetweetedTweetEntitiesSmarttag

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetRetweetedTweetEntitiesSmarttagTag = shared.EmbeddedTweetRetweetedTweetRetweetedTweetEntitiesSmarttagTag

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetRetweetedTweetEntitiesSmarttagTagInfo = shared.EmbeddedTweetRetweetedTweetRetweetedTweetEntitiesSmarttagTagInfo

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetRetweetedTweetEntitiesSmarttagTagInfoInfo = shared.EmbeddedTweetRetweetedTweetRetweetedTweetEntitiesSmarttagTagInfoInfo

// Indexed smart-tag, cashtag, or video timestamp metadata.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetRetweetedTweetEntitiesSymbol = shared.EmbeddedTweetRetweetedTweetRetweetedTweetEntitiesSymbol

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetRetweetedTweetEntitiesSymbolTag = shared.EmbeddedTweetRetweetedTweetRetweetedTweetEntitiesSymbolTag

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetRetweetedTweetEntitiesSymbolTagInfo = shared.EmbeddedTweetRetweetedTweetRetweetedTweetEntitiesSymbolTagInfo

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetRetweetedTweetEntitiesSymbolTagInfoInfo = shared.EmbeddedTweetRetweetedTweetRetweetedTweetEntitiesSymbolTagInfoInfo

// Indexed smart-tag, cashtag, or video timestamp metadata.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetRetweetedTweetEntitiesTimestamp = shared.EmbeddedTweetRetweetedTweetRetweetedTweetEntitiesTimestamp

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetRetweetedTweetEntitiesTimestampTag = shared.EmbeddedTweetRetweetedTweetRetweetedTweetEntitiesTimestampTag

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetRetweetedTweetEntitiesTimestampTagInfo = shared.EmbeddedTweetRetweetedTweetRetweetedTweetEntitiesTimestampTagInfo

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetRetweetedTweetEntitiesTimestampTagInfoInfo = shared.EmbeddedTweetRetweetedTweetRetweetedTweetEntitiesTimestampTagInfoInfo

// Provides shortened, display, and expanded URLs from tweet text.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetRetweetedTweetEntitiesURL = shared.EmbeddedTweetRetweetedTweetRetweetedTweetEntitiesURL

// Provides profile identity and source offsets for a mention.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetRetweetedTweetEntitiesUserMention = shared.EmbeddedTweetRetweetedTweetRetweetedTweetEntitiesUserMention

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetRetweetedTweetLimitedAction = shared.EmbeddedTweetRetweetedTweetRetweetedTweetLimitedAction

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetRetweetedTweetLimitedActionPrompt = shared.EmbeddedTweetRetweetedTweetRetweetedTweetLimitedActionPrompt

// Complete Note Tweet content and rich-text metadata.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetRetweetedTweetNoteTweet = shared.EmbeddedTweetRetweetedTweetRetweetedTweetNoteTweet

// Lists hashtags, symbols, links, and mentions from tweet text.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetRetweetedTweetNoteTweetEntities = shared.EmbeddedTweetRetweetedTweetRetweetedTweetNoteTweetEntities

// Provides hashtag text and source offsets within a tweet.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesHashtag = shared.EmbeddedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesHashtag

// Indexed smart-tag, cashtag, or video timestamp metadata.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesSmarttag = shared.EmbeddedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesSmarttag

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesSmarttagTag = shared.EmbeddedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesSmarttagTag

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesSmarttagTagInfo = shared.EmbeddedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesSmarttagTagInfo

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesSmarttagTagInfoInfo = shared.EmbeddedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesSmarttagTagInfoInfo

// Indexed smart-tag, cashtag, or video timestamp metadata.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesSymbol = shared.EmbeddedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesSymbol

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesSymbolTag = shared.EmbeddedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesSymbolTag

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesSymbolTagInfo = shared.EmbeddedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesSymbolTagInfo

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesSymbolTagInfoInfo = shared.EmbeddedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesSymbolTagInfoInfo

// Indexed smart-tag, cashtag, or video timestamp metadata.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesTimestamp = shared.EmbeddedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesTimestamp

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesTimestampTag = shared.EmbeddedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesTimestampTag

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesTimestampTagInfo = shared.EmbeddedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesTimestampTagInfo

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesTimestampTagInfoInfo = shared.EmbeddedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesTimestampTagInfoInfo

// Provides shortened, display, and expanded URLs from tweet text.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesURL = shared.EmbeddedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesURL

// Provides profile identity and source offsets for a mention.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesUserMention = shared.EmbeddedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesUserMention

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetRetweetedTweetNoteTweetInlineMedia = shared.EmbeddedTweetRetweetedTweetRetweetedTweetNoteTweetInlineMedia

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetRetweetedTweetNoteTweetRichtextTag = shared.EmbeddedTweetRetweetedTweetRetweetedTweetNoteTweetRichtextTag

// Describes public place metadata on a geotagged tweet.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetRetweetedTweetPlace = shared.EmbeddedTweetRetweetedTweetRetweetedTweetPlace

// Engagement counts retained from a prior tweet edit.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetRetweetedTweetPreviousCounts = shared.EmbeddedTweetRetweetedTweetRetweetedTweetPreviousCounts

// Final nested tweet context at depth 4.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweet = shared.EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweet

// Describes an X Article preview and its lifecycle metadata.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetArticle = shared.EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetArticle

// Describes a public card and its referenced profiles.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetCard = shared.EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetCard

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetCardUserReferenceError = shared.EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetCardUserReferenceError

// Community Note presentation metadata returned by X.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetCommunityNote = shared.EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetCommunityNote

// Public reply policy and conversation owner.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetConversationControl = shared.EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetConversationControl

// Lists edit-chain identifiers and the remaining edit window.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetEdit = shared.EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetEdit

// Lists hashtags, symbols, links, and mentions from tweet text.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetEntities = shared.EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetEntities

// Provides hashtag text and source offsets within a tweet.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetEntitiesHashtag = shared.EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetEntitiesHashtag

// Indexed smart-tag, cashtag, or video timestamp metadata.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetEntitiesSmarttag = shared.EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetEntitiesSmarttag

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetEntitiesSmarttagTag = shared.EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetEntitiesSmarttagTag

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetEntitiesSmarttagTagInfo = shared.EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetEntitiesSmarttagTagInfo

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetEntitiesSmarttagTagInfoInfo = shared.EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetEntitiesSmarttagTagInfoInfo

// Indexed smart-tag, cashtag, or video timestamp metadata.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetEntitiesSymbol = shared.EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetEntitiesSymbol

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetEntitiesSymbolTag = shared.EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetEntitiesSymbolTag

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetEntitiesSymbolTagInfo = shared.EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetEntitiesSymbolTagInfo

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetEntitiesSymbolTagInfoInfo = shared.EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetEntitiesSymbolTagInfoInfo

// Indexed smart-tag, cashtag, or video timestamp metadata.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetEntitiesTimestamp = shared.EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetEntitiesTimestamp

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetEntitiesTimestampTag = shared.EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetEntitiesTimestampTag

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetEntitiesTimestampTagInfo = shared.EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetEntitiesTimestampTagInfo

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetEntitiesTimestampTagInfoInfo = shared.EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetEntitiesTimestampTagInfoInfo

// Provides shortened, display, and expanded URLs from tweet text.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetEntitiesURL = shared.EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetEntitiesURL

// Provides profile identity and source offsets for a mention.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetEntitiesUserMention = shared.EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetEntitiesUserMention

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetLimitedAction = shared.EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetLimitedAction

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetLimitedActionPrompt = shared.EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetLimitedActionPrompt

// Complete Note Tweet content and rich-text metadata.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetNoteTweet = shared.EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetNoteTweet

// Lists hashtags, symbols, links, and mentions from tweet text.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetNoteTweetEntities = shared.EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetNoteTweetEntities

// Provides hashtag text and source offsets within a tweet.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesHashtag = shared.EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesHashtag

// Indexed smart-tag, cashtag, or video timestamp metadata.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesSmarttag = shared.EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesSmarttag

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesSmarttagTag = shared.EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesSmarttagTag

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesSmarttagTagInfo = shared.EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesSmarttagTagInfo

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesSmarttagTagInfoInfo = shared.EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesSmarttagTagInfoInfo

// Indexed smart-tag, cashtag, or video timestamp metadata.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesSymbol = shared.EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesSymbol

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesSymbolTag = shared.EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesSymbolTag

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesSymbolTagInfo = shared.EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesSymbolTagInfo

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesSymbolTagInfoInfo = shared.EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesSymbolTagInfoInfo

// Indexed smart-tag, cashtag, or video timestamp metadata.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesTimestamp = shared.EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesTimestamp

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesTimestampTag = shared.EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesTimestampTag

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesTimestampTagInfo = shared.EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesTimestampTagInfo

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesTimestampTagInfoInfo = shared.EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesTimestampTagInfoInfo

// Provides shortened, display, and expanded URLs from tweet text.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesURL = shared.EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesURL

// Provides profile identity and source offsets for a mention.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesUserMention = shared.EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetNoteTweetEntitiesUserMention

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetNoteTweetInlineMedia = shared.EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetNoteTweetInlineMedia

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetNoteTweetRichtextTag = shared.EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetNoteTweetRichtextTag

// Describes public place metadata on a geotagged tweet.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetPlace = shared.EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetPlace

// Engagement counts retained from a prior tweet edit.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetPreviousCounts = shared.EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetPreviousCounts

// Public post and user referenced by this reaction.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetReactionContext = shared.EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetReactionContext

// Public visibility notice attached to an available tweet.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetTombstone = shared.EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetTombstone

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetTombstoneText = shared.EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetTombstoneText

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetTombstoneTextEntity = shared.EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetTombstoneTextEntity

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetTombstoneTextEntityRef = shared.EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetTombstoneTextEntityRef

// Public post and user referenced by this reaction.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetRetweetedTweetReactionContext = shared.EmbeddedTweetRetweetedTweetRetweetedTweetReactionContext

// Final nested tweet context at depth 4.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweet = shared.EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweet

// Describes an X Article preview and its lifecycle metadata.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetArticle = shared.EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetArticle

// Describes a public card and its referenced profiles.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetCard = shared.EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetCard

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetCardUserReferenceError = shared.EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetCardUserReferenceError

// Community Note presentation metadata returned by X.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetCommunityNote = shared.EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetCommunityNote

// Public reply policy and conversation owner.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetConversationControl = shared.EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetConversationControl

// Lists edit-chain identifiers and the remaining edit window.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetEdit = shared.EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetEdit

// Lists hashtags, symbols, links, and mentions from tweet text.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetEntities = shared.EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetEntities

// Provides hashtag text and source offsets within a tweet.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetEntitiesHashtag = shared.EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetEntitiesHashtag

// Indexed smart-tag, cashtag, or video timestamp metadata.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetEntitiesSmarttag = shared.EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetEntitiesSmarttag

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetEntitiesSmarttagTag = shared.EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetEntitiesSmarttagTag

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetEntitiesSmarttagTagInfo = shared.EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetEntitiesSmarttagTagInfo

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetEntitiesSmarttagTagInfoInfo = shared.EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetEntitiesSmarttagTagInfoInfo

// Indexed smart-tag, cashtag, or video timestamp metadata.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetEntitiesSymbol = shared.EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetEntitiesSymbol

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetEntitiesSymbolTag = shared.EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetEntitiesSymbolTag

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetEntitiesSymbolTagInfo = shared.EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetEntitiesSymbolTagInfo

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetEntitiesSymbolTagInfoInfo = shared.EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetEntitiesSymbolTagInfoInfo

// Indexed smart-tag, cashtag, or video timestamp metadata.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetEntitiesTimestamp = shared.EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetEntitiesTimestamp

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetEntitiesTimestampTag = shared.EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetEntitiesTimestampTag

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetEntitiesTimestampTagInfo = shared.EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetEntitiesTimestampTagInfo

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetEntitiesTimestampTagInfoInfo = shared.EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetEntitiesTimestampTagInfoInfo

// Provides shortened, display, and expanded URLs from tweet text.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetEntitiesURL = shared.EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetEntitiesURL

// Provides profile identity and source offsets for a mention.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetEntitiesUserMention = shared.EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetEntitiesUserMention

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetLimitedAction = shared.EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetLimitedAction

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetLimitedActionPrompt = shared.EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetLimitedActionPrompt

// Complete Note Tweet content and rich-text metadata.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetNoteTweet = shared.EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetNoteTweet

// Lists hashtags, symbols, links, and mentions from tweet text.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetNoteTweetEntities = shared.EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetNoteTweetEntities

// Provides hashtag text and source offsets within a tweet.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesHashtag = shared.EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesHashtag

// Indexed smart-tag, cashtag, or video timestamp metadata.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesSmarttag = shared.EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesSmarttag

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesSmarttagTag = shared.EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesSmarttagTag

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesSmarttagTagInfo = shared.EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesSmarttagTagInfo

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesSmarttagTagInfoInfo = shared.EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesSmarttagTagInfoInfo

// Indexed smart-tag, cashtag, or video timestamp metadata.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesSymbol = shared.EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesSymbol

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesSymbolTag = shared.EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesSymbolTag

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesSymbolTagInfo = shared.EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesSymbolTagInfo

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesSymbolTagInfoInfo = shared.EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesSymbolTagInfoInfo

// Indexed smart-tag, cashtag, or video timestamp metadata.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesTimestamp = shared.EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesTimestamp

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesTimestampTag = shared.EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesTimestampTag

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesTimestampTagInfo = shared.EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesTimestampTagInfo

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesTimestampTagInfoInfo = shared.EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesTimestampTagInfoInfo

// Provides shortened, display, and expanded URLs from tweet text.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesURL = shared.EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesURL

// Provides profile identity and source offsets for a mention.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesUserMention = shared.EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetNoteTweetEntitiesUserMention

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetNoteTweetInlineMedia = shared.EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetNoteTweetInlineMedia

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetNoteTweetRichtextTag = shared.EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetNoteTweetRichtextTag

// Describes public place metadata on a geotagged tweet.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetPlace = shared.EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetPlace

// Engagement counts retained from a prior tweet edit.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetPreviousCounts = shared.EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetPreviousCounts

// Public post and user referenced by this reaction.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetReactionContext = shared.EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetReactionContext

// Public visibility notice attached to an available tweet.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetTombstone = shared.EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetTombstone

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetTombstoneText = shared.EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetTombstoneText

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetTombstoneTextEntity = shared.EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetTombstoneTextEntity

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetTombstoneTextEntityRef = shared.EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetTombstoneTextEntityRef

// Public visibility notice attached to an available tweet.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetRetweetedTweetTombstone = shared.EmbeddedTweetRetweetedTweetRetweetedTweetTombstone

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetRetweetedTweetTombstoneText = shared.EmbeddedTweetRetweetedTweetRetweetedTweetTombstoneText

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetRetweetedTweetTombstoneTextEntity = shared.EmbeddedTweetRetweetedTweetRetweetedTweetTombstoneTextEntity

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetRetweetedTweetTombstoneTextEntityRef = shared.EmbeddedTweetRetweetedTweetRetweetedTweetTombstoneTextEntityRef

// Public visibility notice attached to an available tweet.
//
// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetTombstone = shared.EmbeddedTweetRetweetedTweetTombstone

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetTombstoneText = shared.EmbeddedTweetRetweetedTweetTombstoneText

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetTombstoneTextEntity = shared.EmbeddedTweetRetweetedTweetTombstoneTextEntity

// This is an alias to an internal type.
type EmbeddedTweetRetweetedTweetTombstoneTextEntityRef = shared.EmbeddedTweetRetweetedTweetTombstoneTextEntityRef

// Public visibility notice attached to an available tweet.
//
// This is an alias to an internal type.
type EmbeddedTweetTombstone = shared.EmbeddedTweetTombstone

// This is an alias to an internal type.
type EmbeddedTweetTombstoneText = shared.EmbeddedTweetTombstoneText

// This is an alias to an internal type.
type EmbeddedTweetTombstoneTextEntity = shared.EmbeddedTweetTombstoneTextEntity

// This is an alias to an internal type.
type EmbeddedTweetTombstoneTextEntityRef = shared.EmbeddedTweetTombstoneTextEntityRef

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

// Automatic search, user Tweet, and reply coverage preserves shape, filters,
// aliases, and billing. Follow next_cursor while the response reports more pages.
// An empty filtered page can still require continuation. Unprefixed cursors are
// legacy.
//
// This is an alias to an internal type.
type PaginatedTweets = shared.PaginatedTweets

// Profile coverage preserves shape, billing, aliases, and filters. Follow
// next_cursor while the response reports more pages. Unprefixed cursors remain
// legacy.
//
// This is an alias to an internal type.
type PaginatedUsers = shared.PaginatedUsers

// Tweet returned from search results with inline author info. A zero metric can
// mean X did not report the count.
//
// This is an alias to an internal type.
type SearchTweet = shared.SearchTweet

// Describes an X Article preview and its lifecycle metadata.
//
// This is an alias to an internal type.
type SearchTweetArticle = shared.SearchTweetArticle

// Describes a public card and its referenced profiles.
//
// This is an alias to an internal type.
type SearchTweetCard = shared.SearchTweetCard

// This is an alias to an internal type.
type SearchTweetCardUserReferenceError = shared.SearchTweetCardUserReferenceError

// Community Note presentation metadata returned by X.
//
// This is an alias to an internal type.
type SearchTweetCommunityNote = shared.SearchTweetCommunityNote

// Public reply policy and conversation owner.
//
// This is an alias to an internal type.
type SearchTweetConversationControl = shared.SearchTweetConversationControl

// Lists edit-chain identifiers and the remaining edit window.
//
// This is an alias to an internal type.
type SearchTweetEdit = shared.SearchTweetEdit

// Lists hashtags, symbols, links, and mentions from tweet text.
//
// This is an alias to an internal type.
type SearchTweetEntities = shared.SearchTweetEntities

// Provides hashtag text and source offsets within a tweet.
//
// This is an alias to an internal type.
type SearchTweetEntitiesHashtag = shared.SearchTweetEntitiesHashtag

// Indexed smart-tag, cashtag, or video timestamp metadata.
//
// This is an alias to an internal type.
type SearchTweetEntitiesSmarttag = shared.SearchTweetEntitiesSmarttag

// This is an alias to an internal type.
type SearchTweetEntitiesSmarttagTag = shared.SearchTweetEntitiesSmarttagTag

// This is an alias to an internal type.
type SearchTweetEntitiesSmarttagTagInfo = shared.SearchTweetEntitiesSmarttagTagInfo

// This is an alias to an internal type.
type SearchTweetEntitiesSmarttagTagInfoInfo = shared.SearchTweetEntitiesSmarttagTagInfoInfo

// Indexed smart-tag, cashtag, or video timestamp metadata.
//
// This is an alias to an internal type.
type SearchTweetEntitiesSymbol = shared.SearchTweetEntitiesSymbol

// This is an alias to an internal type.
type SearchTweetEntitiesSymbolTag = shared.SearchTweetEntitiesSymbolTag

// This is an alias to an internal type.
type SearchTweetEntitiesSymbolTagInfo = shared.SearchTweetEntitiesSymbolTagInfo

// This is an alias to an internal type.
type SearchTweetEntitiesSymbolTagInfoInfo = shared.SearchTweetEntitiesSymbolTagInfoInfo

// Indexed smart-tag, cashtag, or video timestamp metadata.
//
// This is an alias to an internal type.
type SearchTweetEntitiesTimestamp = shared.SearchTweetEntitiesTimestamp

// This is an alias to an internal type.
type SearchTweetEntitiesTimestampTag = shared.SearchTweetEntitiesTimestampTag

// This is an alias to an internal type.
type SearchTweetEntitiesTimestampTagInfo = shared.SearchTweetEntitiesTimestampTagInfo

// This is an alias to an internal type.
type SearchTweetEntitiesTimestampTagInfoInfo = shared.SearchTweetEntitiesTimestampTagInfoInfo

// Provides shortened, display, and expanded URLs from tweet text.
//
// This is an alias to an internal type.
type SearchTweetEntitiesURL = shared.SearchTweetEntitiesURL

// Provides profile identity and source offsets for a mention.
//
// This is an alias to an internal type.
type SearchTweetEntitiesUserMention = shared.SearchTweetEntitiesUserMention

// This is an alias to an internal type.
type SearchTweetLimitedAction = shared.SearchTweetLimitedAction

// This is an alias to an internal type.
type SearchTweetLimitedActionPrompt = shared.SearchTweetLimitedActionPrompt

// Complete Note Tweet content and rich-text metadata.
//
// This is an alias to an internal type.
type SearchTweetNoteTweet = shared.SearchTweetNoteTweet

// Lists hashtags, symbols, links, and mentions from tweet text.
//
// This is an alias to an internal type.
type SearchTweetNoteTweetEntities = shared.SearchTweetNoteTweetEntities

// Provides hashtag text and source offsets within a tweet.
//
// This is an alias to an internal type.
type SearchTweetNoteTweetEntitiesHashtag = shared.SearchTweetNoteTweetEntitiesHashtag

// Indexed smart-tag, cashtag, or video timestamp metadata.
//
// This is an alias to an internal type.
type SearchTweetNoteTweetEntitiesSmarttag = shared.SearchTweetNoteTweetEntitiesSmarttag

// This is an alias to an internal type.
type SearchTweetNoteTweetEntitiesSmarttagTag = shared.SearchTweetNoteTweetEntitiesSmarttagTag

// This is an alias to an internal type.
type SearchTweetNoteTweetEntitiesSmarttagTagInfo = shared.SearchTweetNoteTweetEntitiesSmarttagTagInfo

// This is an alias to an internal type.
type SearchTweetNoteTweetEntitiesSmarttagTagInfoInfo = shared.SearchTweetNoteTweetEntitiesSmarttagTagInfoInfo

// Indexed smart-tag, cashtag, or video timestamp metadata.
//
// This is an alias to an internal type.
type SearchTweetNoteTweetEntitiesSymbol = shared.SearchTweetNoteTweetEntitiesSymbol

// This is an alias to an internal type.
type SearchTweetNoteTweetEntitiesSymbolTag = shared.SearchTweetNoteTweetEntitiesSymbolTag

// This is an alias to an internal type.
type SearchTweetNoteTweetEntitiesSymbolTagInfo = shared.SearchTweetNoteTweetEntitiesSymbolTagInfo

// This is an alias to an internal type.
type SearchTweetNoteTweetEntitiesSymbolTagInfoInfo = shared.SearchTweetNoteTweetEntitiesSymbolTagInfoInfo

// Indexed smart-tag, cashtag, or video timestamp metadata.
//
// This is an alias to an internal type.
type SearchTweetNoteTweetEntitiesTimestamp = shared.SearchTweetNoteTweetEntitiesTimestamp

// This is an alias to an internal type.
type SearchTweetNoteTweetEntitiesTimestampTag = shared.SearchTweetNoteTweetEntitiesTimestampTag

// This is an alias to an internal type.
type SearchTweetNoteTweetEntitiesTimestampTagInfo = shared.SearchTweetNoteTweetEntitiesTimestampTagInfo

// This is an alias to an internal type.
type SearchTweetNoteTweetEntitiesTimestampTagInfoInfo = shared.SearchTweetNoteTweetEntitiesTimestampTagInfoInfo

// Provides shortened, display, and expanded URLs from tweet text.
//
// This is an alias to an internal type.
type SearchTweetNoteTweetEntitiesURL = shared.SearchTweetNoteTweetEntitiesURL

// Provides profile identity and source offsets for a mention.
//
// This is an alias to an internal type.
type SearchTweetNoteTweetEntitiesUserMention = shared.SearchTweetNoteTweetEntitiesUserMention

// This is an alias to an internal type.
type SearchTweetNoteTweetInlineMedia = shared.SearchTweetNoteTweetInlineMedia

// This is an alias to an internal type.
type SearchTweetNoteTweetRichtextTag = shared.SearchTweetNoteTweetRichtextTag

// Describes public place metadata on a geotagged tweet.
//
// This is an alias to an internal type.
type SearchTweetPlace = shared.SearchTweetPlace

// Engagement counts retained from a prior tweet edit.
//
// This is an alias to an internal type.
type SearchTweetPreviousCounts = shared.SearchTweetPreviousCounts

// Public post and user referenced by this reaction.
//
// This is an alias to an internal type.
type SearchTweetReactionContext = shared.SearchTweetReactionContext

// Public visibility notice attached to an available tweet.
//
// This is an alias to an internal type.
type SearchTweetTombstone = shared.SearchTweetTombstone

// This is an alias to an internal type.
type SearchTweetTombstoneText = shared.SearchTweetTombstoneText

// This is an alias to an internal type.
type SearchTweetTombstoneTextEntity = shared.SearchTweetTombstoneTextEntity

// This is an alias to an internal type.
type SearchTweetTombstoneTextEntityRef = shared.SearchTweetTombstoneTextEntityRef

// Tweet media.
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
type TweetMediaTag = shared.TweetMediaTag

// This is an alias to an internal type.
type TweetMediaVideoVariant = shared.TweetMediaVideoVariant

// X user profile with bio, follower counts, and verification status.
//
// This is an alias to an internal type.
type UserProfile = shared.UserProfile

// X's best-effort public label inferred from aggregated account-access IP
// addresses. It does not state nationality, residence, identity, registration,
// post location, or exact location.
//
// This is an alias to an internal type.
type UserProfileAccountBasedIn = shared.UserProfileAccountBasedIn

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

// Public payment and creator-support handles shown on X.
//
// This is an alias to an internal type.
type UserProfileTipJar = shared.UserProfileTipJar
