// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package xtwitterscraper

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/Xquik-dev/x-twitter-scraper-go/internal/apijson"
	"github.com/Xquik-dev/x-twitter-scraper-go/internal/apiquery"
	"github.com/Xquik-dev/x-twitter-scraper-go/internal/requestconfig"
	"github.com/Xquik-dev/x-twitter-scraper-go/option"
	"github.com/Xquik-dev/x-twitter-scraper-go/packages/param"
	"github.com/Xquik-dev/x-twitter-scraper-go/packages/respjson"
)

// X data lookups (subscription required)
//
// XUserService contains methods and other services that help with interacting with
// the x-twitter-scraper API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewXUserService] method instead.
type XUserService struct {
	options []option.RequestOption
	Follow  XUserFollowService
}

// NewXUserService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewXUserService(opts ...option.RequestOption) (r XUserService) {
	r = XUserService{}
	r.options = opts
	r.Follow = NewXUserFollowService(opts...)
	return
}

// Get multiple users by IDs
func (r *XUserService) GetBatch(ctx context.Context, query XUserGetBatchParams, opts ...option.RequestOption) (res *XUserGetBatchResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "x/users/batch"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Get user followers
func (r *XUserService) GetFollowers(ctx context.Context, id string, query XUserGetFollowersParams, opts ...option.RequestOption) (res *XUserGetFollowersResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("x/users/%s/followers", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Get followers you know for a user
func (r *XUserService) GetFollowersYouKnow(ctx context.Context, id string, query XUserGetFollowersYouKnowParams, opts ...option.RequestOption) (res *XUserGetFollowersYouKnowResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("x/users/%s/followers-you-know", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Get users this user follows
func (r *XUserService) GetFollowing(ctx context.Context, id string, query XUserGetFollowingParams, opts ...option.RequestOption) (res *XUserGetFollowingResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("x/users/%s/following", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Get tweets liked by a user
func (r *XUserService) GetLikes(ctx context.Context, id string, query XUserGetLikesParams, opts ...option.RequestOption) (res *XUserGetLikesResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("x/users/%s/likes", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Get media tweets by a user
func (r *XUserService) GetMedia(ctx context.Context, id string, query XUserGetMediaParams, opts ...option.RequestOption) (res *XUserGetMediaResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("x/users/%s/media", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Get tweets mentioning a user
func (r *XUserService) GetMentions(ctx context.Context, id string, query XUserGetMentionsParams, opts ...option.RequestOption) (res *XUserGetMentionsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("x/users/%s/mentions", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Search users by name or username
func (r *XUserService) GetSearch(ctx context.Context, query XUserGetSearchParams, opts ...option.RequestOption) (res *XUserGetSearchResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "x/users/search"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Get recent tweets by a user
func (r *XUserService) GetTweets(ctx context.Context, id string, query XUserGetTweetsParams, opts ...option.RequestOption) (res *XUserGetTweetsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("x/users/%s/tweets", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Get verified followers
func (r *XUserService) GetVerifiedFollowers(ctx context.Context, id string, query XUserGetVerifiedFollowersParams, opts ...option.RequestOption) (res *XUserGetVerifiedFollowersResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("x/users/%s/verified-followers", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Paginated list of user profiles with cursor-based navigation.
type XUserGetBatchResponse struct {
	HasNextPage bool                        `json:"has_next_page" api:"required"`
	NextCursor  string                      `json:"next_cursor" api:"required"`
	Users       []XUserGetBatchResponseUser `json:"users" api:"required"`
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
func (r XUserGetBatchResponse) RawJSON() string { return r.JSON.raw }
func (r *XUserGetBatchResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// X user profile with bio, follower counts, and verification status.
type XUserGetBatchResponseUser struct {
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
func (r XUserGetBatchResponseUser) RawJSON() string { return r.JSON.raw }
func (r *XUserGetBatchResponseUser) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Paginated list of user profiles with cursor-based navigation.
type XUserGetFollowersResponse struct {
	HasNextPage bool                            `json:"has_next_page" api:"required"`
	NextCursor  string                          `json:"next_cursor" api:"required"`
	Users       []XUserGetFollowersResponseUser `json:"users" api:"required"`
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
func (r XUserGetFollowersResponse) RawJSON() string { return r.JSON.raw }
func (r *XUserGetFollowersResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// X user profile with bio, follower counts, and verification status.
type XUserGetFollowersResponseUser struct {
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
func (r XUserGetFollowersResponseUser) RawJSON() string { return r.JSON.raw }
func (r *XUserGetFollowersResponseUser) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Paginated list of user profiles with cursor-based navigation.
type XUserGetFollowersYouKnowResponse struct {
	HasNextPage bool                                   `json:"has_next_page" api:"required"`
	NextCursor  string                                 `json:"next_cursor" api:"required"`
	Users       []XUserGetFollowersYouKnowResponseUser `json:"users" api:"required"`
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
func (r XUserGetFollowersYouKnowResponse) RawJSON() string { return r.JSON.raw }
func (r *XUserGetFollowersYouKnowResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// X user profile with bio, follower counts, and verification status.
type XUserGetFollowersYouKnowResponseUser struct {
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
func (r XUserGetFollowersYouKnowResponseUser) RawJSON() string { return r.JSON.raw }
func (r *XUserGetFollowersYouKnowResponseUser) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Paginated list of user profiles with cursor-based navigation.
type XUserGetFollowingResponse struct {
	HasNextPage bool                            `json:"has_next_page" api:"required"`
	NextCursor  string                          `json:"next_cursor" api:"required"`
	Users       []XUserGetFollowingResponseUser `json:"users" api:"required"`
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
func (r XUserGetFollowingResponse) RawJSON() string { return r.JSON.raw }
func (r *XUserGetFollowingResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// X user profile with bio, follower counts, and verification status.
type XUserGetFollowingResponseUser struct {
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
func (r XUserGetFollowingResponseUser) RawJSON() string { return r.JSON.raw }
func (r *XUserGetFollowingResponseUser) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Paginated list of tweets with cursor-based navigation.
type XUserGetLikesResponse struct {
	HasNextPage bool                         `json:"has_next_page" api:"required"`
	NextCursor  string                       `json:"next_cursor" api:"required"`
	Tweets      []XUserGetLikesResponseTweet `json:"tweets" api:"required"`
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
func (r XUserGetLikesResponse) RawJSON() string { return r.JSON.raw }
func (r *XUserGetLikesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Tweet returned from search results with inline author info.
type XUserGetLikesResponseTweet struct {
	ID            string                           `json:"id" api:"required"`
	Text          string                           `json:"text" api:"required"`
	Author        XUserGetLikesResponseTweetAuthor `json:"author"`
	BookmarkCount int64                            `json:"bookmarkCount"`
	CreatedAt     string                           `json:"createdAt"`
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
func (r XUserGetLikesResponseTweet) RawJSON() string { return r.JSON.raw }
func (r *XUserGetLikesResponseTweet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type XUserGetLikesResponseTweetAuthor struct {
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
func (r XUserGetLikesResponseTweetAuthor) RawJSON() string { return r.JSON.raw }
func (r *XUserGetLikesResponseTweetAuthor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Paginated list of tweets with cursor-based navigation.
type XUserGetMediaResponse struct {
	HasNextPage bool                         `json:"has_next_page" api:"required"`
	NextCursor  string                       `json:"next_cursor" api:"required"`
	Tweets      []XUserGetMediaResponseTweet `json:"tweets" api:"required"`
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
func (r XUserGetMediaResponse) RawJSON() string { return r.JSON.raw }
func (r *XUserGetMediaResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Tweet returned from search results with inline author info.
type XUserGetMediaResponseTweet struct {
	ID            string                           `json:"id" api:"required"`
	Text          string                           `json:"text" api:"required"`
	Author        XUserGetMediaResponseTweetAuthor `json:"author"`
	BookmarkCount int64                            `json:"bookmarkCount"`
	CreatedAt     string                           `json:"createdAt"`
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
func (r XUserGetMediaResponseTweet) RawJSON() string { return r.JSON.raw }
func (r *XUserGetMediaResponseTweet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type XUserGetMediaResponseTweetAuthor struct {
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
func (r XUserGetMediaResponseTweetAuthor) RawJSON() string { return r.JSON.raw }
func (r *XUserGetMediaResponseTweetAuthor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Paginated list of tweets with cursor-based navigation.
type XUserGetMentionsResponse struct {
	HasNextPage bool                            `json:"has_next_page" api:"required"`
	NextCursor  string                          `json:"next_cursor" api:"required"`
	Tweets      []XUserGetMentionsResponseTweet `json:"tweets" api:"required"`
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
func (r XUserGetMentionsResponse) RawJSON() string { return r.JSON.raw }
func (r *XUserGetMentionsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Tweet returned from search results with inline author info.
type XUserGetMentionsResponseTweet struct {
	ID            string                              `json:"id" api:"required"`
	Text          string                              `json:"text" api:"required"`
	Author        XUserGetMentionsResponseTweetAuthor `json:"author"`
	BookmarkCount int64                               `json:"bookmarkCount"`
	CreatedAt     string                              `json:"createdAt"`
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
func (r XUserGetMentionsResponseTweet) RawJSON() string { return r.JSON.raw }
func (r *XUserGetMentionsResponseTweet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type XUserGetMentionsResponseTweetAuthor struct {
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
func (r XUserGetMentionsResponseTweetAuthor) RawJSON() string { return r.JSON.raw }
func (r *XUserGetMentionsResponseTweetAuthor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Paginated list of user profiles with cursor-based navigation.
type XUserGetSearchResponse struct {
	HasNextPage bool                         `json:"has_next_page" api:"required"`
	NextCursor  string                       `json:"next_cursor" api:"required"`
	Users       []XUserGetSearchResponseUser `json:"users" api:"required"`
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
func (r XUserGetSearchResponse) RawJSON() string { return r.JSON.raw }
func (r *XUserGetSearchResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// X user profile with bio, follower counts, and verification status.
type XUserGetSearchResponseUser struct {
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
func (r XUserGetSearchResponseUser) RawJSON() string { return r.JSON.raw }
func (r *XUserGetSearchResponseUser) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Paginated list of tweets with cursor-based navigation.
type XUserGetTweetsResponse struct {
	HasNextPage bool                          `json:"has_next_page" api:"required"`
	NextCursor  string                        `json:"next_cursor" api:"required"`
	Tweets      []XUserGetTweetsResponseTweet `json:"tweets" api:"required"`
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
func (r XUserGetTweetsResponse) RawJSON() string { return r.JSON.raw }
func (r *XUserGetTweetsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Tweet returned from search results with inline author info.
type XUserGetTweetsResponseTweet struct {
	ID            string                            `json:"id" api:"required"`
	Text          string                            `json:"text" api:"required"`
	Author        XUserGetTweetsResponseTweetAuthor `json:"author"`
	BookmarkCount int64                             `json:"bookmarkCount"`
	CreatedAt     string                            `json:"createdAt"`
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
func (r XUserGetTweetsResponseTweet) RawJSON() string { return r.JSON.raw }
func (r *XUserGetTweetsResponseTweet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type XUserGetTweetsResponseTweetAuthor struct {
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
func (r XUserGetTweetsResponseTweetAuthor) RawJSON() string { return r.JSON.raw }
func (r *XUserGetTweetsResponseTweetAuthor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Paginated list of user profiles with cursor-based navigation.
type XUserGetVerifiedFollowersResponse struct {
	HasNextPage bool                                    `json:"has_next_page" api:"required"`
	NextCursor  string                                  `json:"next_cursor" api:"required"`
	Users       []XUserGetVerifiedFollowersResponseUser `json:"users" api:"required"`
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
func (r XUserGetVerifiedFollowersResponse) RawJSON() string { return r.JSON.raw }
func (r *XUserGetVerifiedFollowersResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// X user profile with bio, follower counts, and verification status.
type XUserGetVerifiedFollowersResponseUser struct {
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
func (r XUserGetVerifiedFollowersResponseUser) RawJSON() string { return r.JSON.raw }
func (r *XUserGetVerifiedFollowersResponseUser) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type XUserGetBatchParams struct {
	// Comma-separated user IDs (max 100)
	IDs string `query:"ids" api:"required" json:"-"`
	paramObj
}

// URLQuery serializes [XUserGetBatchParams]'s query parameters as `url.Values`.
func (r XUserGetBatchParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type XUserGetFollowersParams struct {
	// Pagination cursor for followers list
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Items per page (20-200, default 200)
	PageSize param.Opt[int64] `query:"pageSize,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [XUserGetFollowersParams]'s query parameters as
// `url.Values`.
func (r XUserGetFollowersParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type XUserGetFollowersYouKnowParams struct {
	// Pagination cursor for followers-you-know
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [XUserGetFollowersYouKnowParams]'s query parameters as
// `url.Values`.
func (r XUserGetFollowersYouKnowParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type XUserGetFollowingParams struct {
	// Pagination cursor for following list
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Results per page (20-200, default 200)
	PageSize param.Opt[int64] `query:"pageSize,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [XUserGetFollowingParams]'s query parameters as
// `url.Values`.
func (r XUserGetFollowingParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type XUserGetLikesParams struct {
	// Pagination cursor for liked tweets
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [XUserGetLikesParams]'s query parameters as `url.Values`.
func (r XUserGetLikesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type XUserGetMediaParams struct {
	// Pagination cursor for media tweets
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [XUserGetMediaParams]'s query parameters as `url.Values`.
func (r XUserGetMediaParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type XUserGetMentionsParams struct {
	// Pagination cursor for mentions
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Unix timestamp - return mentions after this time
	SinceTime param.Opt[string] `query:"sinceTime,omitzero" json:"-"`
	// Unix timestamp - return mentions before this time
	UntilTime param.Opt[string] `query:"untilTime,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [XUserGetMentionsParams]'s query parameters as `url.Values`.
func (r XUserGetMentionsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type XUserGetSearchParams struct {
	// User search query
	Q string `query:"q" api:"required" json:"-"`
	// Pagination cursor for user search
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [XUserGetSearchParams]'s query parameters as `url.Values`.
func (r XUserGetSearchParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type XUserGetTweetsParams struct {
	// Pagination cursor for user tweets
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Include parent tweet for replies
	IncludeParentTweet param.Opt[bool] `query:"includeParentTweet,omitzero" json:"-"`
	// Include reply tweets
	IncludeReplies param.Opt[bool] `query:"includeReplies,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [XUserGetTweetsParams]'s query parameters as `url.Values`.
func (r XUserGetTweetsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type XUserGetVerifiedFollowersParams struct {
	// Pagination cursor for verified followers
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [XUserGetVerifiedFollowersParams]'s query parameters as
// `url.Values`.
func (r XUserGetVerifiedFollowersParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
