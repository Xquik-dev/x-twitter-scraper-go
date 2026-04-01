// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package xtwitterscraper

import (
	"github.com/stainless-sdks/x-twitter-scraper-go/internal/apierror"
	"github.com/stainless-sdks/x-twitter-scraper-go/packages/param"
	"github.com/stainless-sdks/x-twitter-scraper-go/shared"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

type Error = apierror.Error

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
