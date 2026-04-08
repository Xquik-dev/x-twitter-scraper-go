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
// XListService contains methods and other services that help with interacting with
// the x-twitter-scraper API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewXListService] method instead.
type XListService struct {
	options []option.RequestOption
}

// NewXListService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewXListService(opts ...option.RequestOption) (r XListService) {
	r = XListService{}
	r.options = opts
	return
}

// Get list followers
func (r *XListService) GetFollowers(ctx context.Context, id string, query XListGetFollowersParams, opts ...option.RequestOption) (res *XListGetFollowersResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("x/lists/%s/followers", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Get list members
func (r *XListService) GetMembers(ctx context.Context, id string, query XListGetMembersParams, opts ...option.RequestOption) (res *XListGetMembersResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("x/lists/%s/members", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Get list tweets
func (r *XListService) GetTweets(ctx context.Context, id string, query XListGetTweetsParams, opts ...option.RequestOption) (res *XListGetTweetsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("x/lists/%s/tweets", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Paginated list of user profiles with cursor-based navigation.
type XListGetFollowersResponse struct {
	HasNextPage bool                            `json:"has_next_page" api:"required"`
	NextCursor  string                          `json:"next_cursor" api:"required"`
	Users       []XListGetFollowersResponseUser `json:"users" api:"required"`
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
func (r XListGetFollowersResponse) RawJSON() string { return r.JSON.raw }
func (r *XListGetFollowersResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// X user profile with bio, follower counts, and verification status.
type XListGetFollowersResponseUser struct {
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
func (r XListGetFollowersResponseUser) RawJSON() string { return r.JSON.raw }
func (r *XListGetFollowersResponseUser) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Paginated list of user profiles with cursor-based navigation.
type XListGetMembersResponse struct {
	HasNextPage bool                          `json:"has_next_page" api:"required"`
	NextCursor  string                        `json:"next_cursor" api:"required"`
	Users       []XListGetMembersResponseUser `json:"users" api:"required"`
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
func (r XListGetMembersResponse) RawJSON() string { return r.JSON.raw }
func (r *XListGetMembersResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// X user profile with bio, follower counts, and verification status.
type XListGetMembersResponseUser struct {
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
func (r XListGetMembersResponseUser) RawJSON() string { return r.JSON.raw }
func (r *XListGetMembersResponseUser) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Paginated list of tweets with cursor-based navigation.
type XListGetTweetsResponse struct {
	HasNextPage bool                          `json:"has_next_page" api:"required"`
	NextCursor  string                        `json:"next_cursor" api:"required"`
	Tweets      []XListGetTweetsResponseTweet `json:"tweets" api:"required"`
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
func (r XListGetTweetsResponse) RawJSON() string { return r.JSON.raw }
func (r *XListGetTweetsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Tweet returned from search results with inline author info.
type XListGetTweetsResponseTweet struct {
	ID            string                            `json:"id" api:"required"`
	Text          string                            `json:"text" api:"required"`
	Author        XListGetTweetsResponseTweetAuthor `json:"author"`
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
func (r XListGetTweetsResponseTweet) RawJSON() string { return r.JSON.raw }
func (r *XListGetTweetsResponseTweet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type XListGetTweetsResponseTweetAuthor struct {
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
func (r XListGetTweetsResponseTweetAuthor) RawJSON() string { return r.JSON.raw }
func (r *XListGetTweetsResponseTweetAuthor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type XListGetFollowersParams struct {
	// Pagination cursor for list followers
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [XListGetFollowersParams]'s query parameters as
// `url.Values`.
func (r XListGetFollowersParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type XListGetMembersParams struct {
	// Pagination cursor for list members
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [XListGetMembersParams]'s query parameters as `url.Values`.
func (r XListGetMembersParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type XListGetTweetsParams struct {
	// Pagination cursor for list tweets
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Include replies (default false)
	IncludeReplies param.Opt[bool] `query:"includeReplies,omitzero" json:"-"`
	// Unix timestamp - filter after
	SinceTime param.Opt[string] `query:"sinceTime,omitzero" json:"-"`
	// Unix timestamp - filter before
	UntilTime param.Opt[string] `query:"untilTime,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [XListGetTweetsParams]'s query parameters as `url.Values`.
func (r XListGetTweetsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
