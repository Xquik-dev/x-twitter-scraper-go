// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package xtwitterscraper

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/stainless-sdks/x-twitter-scraper-go/internal/apijson"
	"github.com/stainless-sdks/x-twitter-scraper-go/internal/apiquery"
	"github.com/stainless-sdks/x-twitter-scraper-go/internal/requestconfig"
	"github.com/stainless-sdks/x-twitter-scraper-go/option"
	"github.com/stainless-sdks/x-twitter-scraper-go/packages/param"
	"github.com/stainless-sdks/x-twitter-scraper-go/packages/respjson"
	"github.com/stainless-sdks/x-twitter-scraper-go/shared"
)

// X data lookups (subscription required)
//
// XService contains methods and other services that help with interacting with the
// x-twitter-scraper API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewXService] method instead.
type XService struct {
	options []option.RequestOption
	Tweets  XTweetService
	// X data lookups (subscription required)
	Users XUserService
	// X data lookups (subscription required)
	Followers XFollowerService
	Dm        XDmService
	// Media upload & download
	Media XMediaService
	// X write actions (tweets, likes, follows, DMs)
	Profile     XProfileService
	Communities XCommunityService
	// Connected X account management
	Accounts XAccountService
	// X data lookups (subscription required)
	Bookmarks XBookmarkService
	// X data lookups (subscription required)
	Lists XListService
}

// NewXService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewXService(opts ...option.RequestOption) (r XService) {
	r = XService{}
	r.options = opts
	r.Tweets = NewXTweetService(opts...)
	r.Users = NewXUserService(opts...)
	r.Followers = NewXFollowerService(opts...)
	r.Dm = NewXDmService(opts...)
	r.Media = NewXMediaService(opts...)
	r.Profile = NewXProfileService(opts...)
	r.Communities = NewXCommunityService(opts...)
	r.Accounts = NewXAccountService(opts...)
	r.Bookmarks = NewXBookmarkService(opts...)
	r.Lists = NewXListService(opts...)
	return
}

// Retrieve the full content of an X Article (long-form post) by tweet ID.
func (r *XService) GetArticle(ctx context.Context, tweetID string, opts ...option.RequestOption) (res *XGetArticleResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if tweetID == "" {
		err = errors.New("missing required tweetId parameter")
		return nil, err
	}
	path := fmt.Sprintf("x/articles/%s", url.PathEscape(tweetID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Get home timeline
func (r *XService) GetHomeTimeline(ctx context.Context, query XGetHomeTimelineParams, opts ...option.RequestOption) (res *shared.PaginatedTweets, err error) {
	opts = slices.Concat(r.options, opts)
	path := "x/timeline"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Get notifications
func (r *XService) GetNotifications(ctx context.Context, query XGetNotificationsParams, opts ...option.RequestOption) (res *XGetNotificationsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "x/notifications"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Get trending topics
func (r *XService) GetTrends(ctx context.Context, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "x/trends"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, nil, opts...)
	return err
}

type XGetArticleResponse struct {
	Article XGetArticleResponseArticle `json:"article" api:"required"`
	Author  TweetAuthor                `json:"author"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Article     respjson.Field
		Author      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XGetArticleResponse) RawJSON() string { return r.JSON.raw }
func (r *XGetArticleResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type XGetArticleResponseArticle struct {
	Contents      []XGetArticleResponseArticleContent `json:"contents"`
	CoverImageURL string                              `json:"coverImageUrl"`
	CreatedAt     string                              `json:"createdAt"`
	LikeCount     int64                               `json:"likeCount"`
	PreviewText   string                              `json:"previewText"`
	QuoteCount    int64                               `json:"quoteCount"`
	ReplyCount    int64                               `json:"replyCount"`
	Title         string                              `json:"title"`
	ViewCount     int64                               `json:"viewCount"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Contents      respjson.Field
		CoverImageURL respjson.Field
		CreatedAt     respjson.Field
		LikeCount     respjson.Field
		PreviewText   respjson.Field
		QuoteCount    respjson.Field
		ReplyCount    respjson.Field
		Title         respjson.Field
		ViewCount     respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XGetArticleResponseArticle) RawJSON() string { return r.JSON.raw }
func (r *XGetArticleResponseArticle) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type XGetArticleResponseArticleContent struct {
	Height int64  `json:"height"`
	Text   string `json:"text"`
	// Block type: unstyled, header-one, header-two, header-three, unordered-list-item,
	// ordered-list-item, image, gif, divider
	Type string `json:"type"`
	// Media URL for image/gif blocks
	URL   string `json:"url"`
	Width int64  `json:"width"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Height      respjson.Field
		Text        respjson.Field
		Type        respjson.Field
		URL         respjson.Field
		Width       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XGetArticleResponseArticleContent) RawJSON() string { return r.JSON.raw }
func (r *XGetArticleResponseArticleContent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type XGetNotificationsResponse struct {
	HasNextPage   bool                                    `json:"has_next_page" api:"required"`
	NextCursor    string                                  `json:"next_cursor" api:"required"`
	Notifications []XGetNotificationsResponseNotification `json:"notifications" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		HasNextPage   respjson.Field
		NextCursor    respjson.Field
		Notifications respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XGetNotificationsResponse) RawJSON() string { return r.JSON.raw }
func (r *XGetNotificationsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type XGetNotificationsResponseNotification struct {
	ID        string `json:"id" api:"required"`
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
	Type      string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Message     respjson.Field
		Timestamp   respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XGetNotificationsResponseNotification) RawJSON() string { return r.JSON.raw }
func (r *XGetNotificationsResponseNotification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type XGetHomeTimelineParams struct {
	// Pagination cursor from previous response
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Comma-separated tweet IDs to exclude from results
	SeenTweetIDs param.Opt[string] `query:"seenTweetIds,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [XGetHomeTimelineParams]'s query parameters as `url.Values`.
func (r XGetHomeTimelineParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type XGetNotificationsParams struct {
	// Pagination cursor from previous response
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Notification type filter
	//
	// Any of "All", "Verified", "Mentions".
	Type XGetNotificationsParamsType `query:"type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [XGetNotificationsParams]'s query parameters as
// `url.Values`.
func (r XGetNotificationsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Notification type filter
type XGetNotificationsParamsType string

const (
	XGetNotificationsParamsTypeAll      XGetNotificationsParamsType = "All"
	XGetNotificationsParamsTypeVerified XGetNotificationsParamsType = "Verified"
	XGetNotificationsParamsTypeMentions XGetNotificationsParamsType = "Mentions"
)
