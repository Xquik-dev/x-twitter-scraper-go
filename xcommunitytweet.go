// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package xtwitterscraper

import (
	"context"
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
// XCommunityTweetService contains methods and other services that help with
// interacting with the x-twitter-scraper API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewXCommunityTweetService] method instead.
type XCommunityTweetService struct {
	options []option.RequestOption
}

// NewXCommunityTweetService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewXCommunityTweetService(opts ...option.RequestOption) (r XCommunityTweetService) {
	r = XCommunityTweetService{}
	r.options = opts
	return
}

// Search tweets across all communities
func (r *XCommunityTweetService) List(ctx context.Context, query XCommunityTweetListParams, opts ...option.RequestOption) (res *XCommunityTweetListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "x/communities/tweets"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Paginated list of tweets with cursor-based navigation.
type XCommunityTweetListResponse struct {
	HasNextPage bool                               `json:"has_next_page" api:"required"`
	NextCursor  string                             `json:"next_cursor" api:"required"`
	Tweets      []XCommunityTweetListResponseTweet `json:"tweets" api:"required"`
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
func (r XCommunityTweetListResponse) RawJSON() string { return r.JSON.raw }
func (r *XCommunityTweetListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Tweet returned from search results with inline author info.
type XCommunityTweetListResponseTweet struct {
	ID            string                                 `json:"id" api:"required"`
	Text          string                                 `json:"text" api:"required"`
	Author        XCommunityTweetListResponseTweetAuthor `json:"author"`
	BookmarkCount int64                                  `json:"bookmarkCount"`
	CreatedAt     string                                 `json:"createdAt"`
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
func (r XCommunityTweetListResponseTweet) RawJSON() string { return r.JSON.raw }
func (r *XCommunityTweetListResponseTweet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type XCommunityTweetListResponseTweetAuthor struct {
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
func (r XCommunityTweetListResponseTweetAuthor) RawJSON() string { return r.JSON.raw }
func (r *XCommunityTweetListResponseTweetAuthor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type XCommunityTweetListParams struct {
	// Search query for cross-community tweets
	Q string `query:"q" api:"required" json:"-"`
	// Pagination cursor for cross-community results
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Sort order for cross-community results (Latest or Top)
	QueryType param.Opt[string] `query:"queryType,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [XCommunityTweetListParams]'s query parameters as
// `url.Values`.
func (r XCommunityTweetListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
