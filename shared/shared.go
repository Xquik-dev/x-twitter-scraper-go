// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

import (
	"github.com/stainless-sdks/x-twitter-scraper-go"
	"github.com/stainless-sdks/x-twitter-scraper-go/internal/apijson"
	"github.com/stainless-sdks/x-twitter-scraper-go/packages/param"
	"github.com/stainless-sdks/x-twitter-scraper-go/packages/respjson"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

type EventType string

const (
	EventTypeTweetNew       EventType = "tweet.new"
	EventTypeTweetReply     EventType = "tweet.reply"
	EventTypeTweetRetweet   EventType = "tweet.retweet"
	EventTypeTweetQuote     EventType = "tweet.quote"
	EventTypeFollowerGained EventType = "follower.gained"
	EventTypeFollowerLost   EventType = "follower.lost"
)

type PaginatedTweets struct {
	HasNextPage bool                          `json:"has_next_page" api:"required"`
	NextCursor  string                        `json:"next_cursor" api:"required"`
	Tweets      []xtwitterscraper.SearchTweet `json:"tweets" api:"required"`
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

type PaginatedUsers struct {
	HasNextPage bool                          `json:"has_next_page" api:"required"`
	NextCursor  string                        `json:"next_cursor" api:"required"`
	Users       []xtwitterscraper.UserProfile `json:"users" api:"required"`
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
