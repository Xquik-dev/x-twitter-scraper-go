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

// Equals "follower.gained"
const EventTypeFollowerGained = shared.EventTypeFollowerGained

// Equals "follower.lost"
const EventTypeFollowerLost = shared.EventTypeFollowerLost

// Paginated list of tweets with cursor-based navigation.
//
// This is an alias to an internal type.
type PaginatedTweets = shared.PaginatedTweets

// Paginated list of user profiles with cursor-based navigation.
//
// This is an alias to an internal type.
type PaginatedUsers = shared.PaginatedUsers
