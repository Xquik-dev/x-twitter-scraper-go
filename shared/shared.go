// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

import (
	"github.com/Xquik-dev/x-twitter-scraper-go/packages/param"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

// Type of monitor event fired when account activity occurs.
type EventType string

const (
	EventTypeTweetNew       EventType = "tweet.new"
	EventTypeTweetReply     EventType = "tweet.reply"
	EventTypeTweetRetweet   EventType = "tweet.retweet"
	EventTypeTweetQuote     EventType = "tweet.quote"
	EventTypeFollowerGained EventType = "follower.gained"
	EventTypeFollowerLost   EventType = "follower.lost"
)
