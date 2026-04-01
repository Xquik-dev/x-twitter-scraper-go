// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package xtwitterscraper

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/stainless-sdks/x-twitter-scraper-go/internal/apijson"
	"github.com/stainless-sdks/x-twitter-scraper-go/internal/apiquery"
	"github.com/stainless-sdks/x-twitter-scraper-go/internal/requestconfig"
	"github.com/stainless-sdks/x-twitter-scraper-go/option"
	"github.com/stainless-sdks/x-twitter-scraper-go/packages/param"
	"github.com/stainless-sdks/x-twitter-scraper-go/packages/respjson"
)

// X data lookups (subscription required)
//
// XBookmarkService contains methods and other services that help with interacting
// with the x-twitter-scraper API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewXBookmarkService] method instead.
type XBookmarkService struct {
	options []option.RequestOption
}

// NewXBookmarkService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewXBookmarkService(opts ...option.RequestOption) (r XBookmarkService) {
	r = XBookmarkService{}
	r.options = opts
	return
}

// Get bookmarked tweets
func (r *XBookmarkService) List(ctx context.Context, query XBookmarkListParams, opts ...option.RequestOption) (res *XBookmarkListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "x/bookmarks"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Get bookmark folders
func (r *XBookmarkService) GetFolders(ctx context.Context, opts ...option.RequestOption) (res *XBookmarkGetFoldersResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "x/bookmarks/folders"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type XBookmarkListResponse struct {
	HasNextPage bool                         `json:"has_next_page" api:"required"`
	NextCursor  string                       `json:"next_cursor" api:"required"`
	Tweets      []XBookmarkListResponseTweet `json:"tweets" api:"required"`
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
func (r XBookmarkListResponse) RawJSON() string { return r.JSON.raw }
func (r *XBookmarkListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type XBookmarkListResponseTweet struct {
	ID            string                           `json:"id" api:"required"`
	Text          string                           `json:"text" api:"required"`
	Author        XBookmarkListResponseTweetAuthor `json:"author"`
	BookmarkCount int64                            `json:"bookmarkCount"`
	CreatedAt     string                           `json:"createdAt"`
	LikeCount     int64                            `json:"likeCount"`
	QuoteCount    int64                            `json:"quoteCount"`
	ReplyCount    int64                            `json:"replyCount"`
	RetweetCount  int64                            `json:"retweetCount"`
	ViewCount     int64                            `json:"viewCount"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		Text          respjson.Field
		Author        respjson.Field
		BookmarkCount respjson.Field
		CreatedAt     respjson.Field
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
func (r XBookmarkListResponseTweet) RawJSON() string { return r.JSON.raw }
func (r *XBookmarkListResponseTweet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type XBookmarkListResponseTweetAuthor struct {
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
func (r XBookmarkListResponseTweetAuthor) RawJSON() string { return r.JSON.raw }
func (r *XBookmarkListResponseTweetAuthor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type XBookmarkGetFoldersResponse struct {
	Folders     []XBookmarkGetFoldersResponseFolder `json:"folders" api:"required"`
	HasNextPage bool                                `json:"has_next_page" api:"required"`
	NextCursor  string                              `json:"next_cursor" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Folders     respjson.Field
		HasNextPage respjson.Field
		NextCursor  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XBookmarkGetFoldersResponse) RawJSON() string { return r.JSON.raw }
func (r *XBookmarkGetFoldersResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type XBookmarkGetFoldersResponseFolder struct {
	ID   string `json:"id" api:"required"`
	Name string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XBookmarkGetFoldersResponseFolder) RawJSON() string { return r.JSON.raw }
func (r *XBookmarkGetFoldersResponseFolder) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type XBookmarkListParams struct {
	// Pagination cursor from previous response
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Optional bookmark folder ID
	FolderID param.Opt[string] `query:"folderId,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [XBookmarkListParams]'s query parameters as `url.Values`.
func (r XBookmarkListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
