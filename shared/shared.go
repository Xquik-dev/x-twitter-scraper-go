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

// Type of monitor event fired when account activity occurs.
type EventType string

const (
	EventTypeTweetNew     EventType = "tweet.new"
	EventTypeTweetReply   EventType = "tweet.reply"
	EventTypeTweetRetweet EventType = "tweet.retweet"
	EventTypeTweetQuote   EventType = "tweet.quote"
)

// Paginated list of tweets with cursor-based navigation.
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

// Paginated list of user profiles with cursor-based navigation.
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

// Tweet returned from search results with inline author info.
type SearchTweet struct {
	ID            string            `json:"id" api:"required"`
	Text          string            `json:"text" api:"required"`
	Author        SearchTweetAuthor `json:"author"`
	BookmarkCount int64             `json:"bookmarkCount"`
	CreatedAt     string            `json:"createdAt"`
	// True for Note Tweets (long-form content, up to 25,000 characters)
	IsNoteTweet  bool  `json:"isNoteTweet"`
	LikeCount    int64 `json:"likeCount"`
	QuoteCount   int64 `json:"quoteCount"`
	ReplyCount   int64 `json:"replyCount"`
	RetweetCount int64 `json:"retweetCount"`
	ViewCount    int64 `json:"viewCount"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		Text          respjson.Field
		Author        respjson.Field
		BookmarkCount respjson.Field
		CreatedAt     respjson.Field
		IsNoteTweet   respjson.Field
		LikeCount     respjson.Field
		QuoteCount    respjson.Field
		ReplyCount    respjson.Field
		RetweetCount  respjson.Field
		ViewCount     respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SearchTweet) RawJSON() string { return r.JSON.raw }
func (r *SearchTweet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SearchTweetAuthor struct {
	ID       string `json:"id" api:"required"`
	Name     string `json:"name" api:"required"`
	Username string `json:"username" api:"required"`
	Verified bool   `json:"verified"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		Username    respjson.Field
		Verified    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SearchTweetAuthor) RawJSON() string { return r.JSON.raw }
func (r *SearchTweetAuthor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// X user profile with bio, follower counts, and verification status.
type UserProfile struct {
	ID             string `json:"id" api:"required"`
	Name           string `json:"name" api:"required"`
	Username       string `json:"username" api:"required"`
	CreatedAt      string `json:"createdAt"`
	Description    string `json:"description"`
	Followers      int64  `json:"followers"`
	Following      int64  `json:"following"`
	Location       string `json:"location"`
	ProfilePicture string `json:"profilePicture"`
	StatusesCount  int64  `json:"statusesCount"`
	Verified       bool   `json:"verified"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		Name           respjson.Field
		Username       respjson.Field
		CreatedAt      respjson.Field
		Description    respjson.Field
		Followers      respjson.Field
		Following      respjson.Field
		Location       respjson.Field
		ProfilePicture respjson.Field
		StatusesCount  respjson.Field
		Verified       respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserProfile) RawJSON() string { return r.JSON.raw }
func (r *UserProfile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
