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
	"github.com/Xquik-dev/x-twitter-scraper-go/shared"
)

// XService contains methods and other services that help with interacting with the
// x-twitter-scraper API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewXService] method instead.
type XService struct {
	options []option.RequestOption
	// X write actions (tweets, likes, follows, DMs)
	WriteActions XWriteActionService
	Tweets       XTweetService
	Users        XUserService
	// Look up, search, and explore user profiles and relationships
	Followers XFollowerService
	Dm        XDmService
	Media     XMediaService
	// X write actions (tweets, likes, follows, DMs)
	Profile     XProfileService
	Communities XCommunityService
	// Connected X account management
	Accounts XAccountService
	// Connected X account management
	AccountConnectionChallenges XAccountConnectionChallengeService
	// Look up, search, and analyze individual tweets
	Bookmarks XBookmarkService
	// X List followers, members, and tweets
	Lists XListService
}

// NewXService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewXService(opts ...option.RequestOption) (r XService) {
	r = XService{}
	r.options = opts
	r.WriteActions = NewXWriteActionService(opts...)
	r.Tweets = NewXTweetService(opts...)
	r.Users = NewXUserService(opts...)
	r.Followers = NewXFollowerService(opts...)
	r.Dm = NewXDmService(opts...)
	r.Media = NewXMediaService(opts...)
	r.Profile = NewXProfileService(opts...)
	r.Communities = NewXCommunityService(opts...)
	r.Accounts = NewXAccountService(opts...)
	r.AccountConnectionChallenges = NewXAccountConnectionChallengeService(opts...)
	r.Bookmarks = NewXBookmarkService(opts...)
	r.Lists = NewXListService(opts...)
	return
}

// Retrieve the full content of an X Article (long-form post) by numeric tweet ID.
// Returns article_not_found when the tweet is valid but is not an X Article.
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

// Get trending hashtags and topics from X by region
func (r *XService) GetTrends(ctx context.Context, query XGetTrendsParams, opts ...option.RequestOption) (res *XGetTrendsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "x/trends"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type XGetArticleResponse struct {
	Article XGetArticleResponseArticle `json:"article" api:"required"`
	// X Article author profile fields returned when available.
	Author XGetArticleResponseAuthor `json:"author"`
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
	// Plain text joined from all article blocks
	BodyText      string                              `json:"bodyText"`
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
		BodyText      respjson.Field
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
	Height int64 `json:"height"`
	// Inline text formatting ranges
	InlineStyleRanges []XGetArticleResponseArticleContentInlineStyleRange `json:"inlineStyleRanges"`
	// Preview image URL for media blocks
	PreviewURL string `json:"previewUrl"`
	Text       string `json:"text"`
	// Block type: paragraph, header-one, header-two, header-three, header-four,
	// header-five, header-six, unordered-list-item, ordered-list-item, blockquote,
	// code-block, media, divider
	Type string `json:"type"`
	// Media URL for media blocks
	URL   string `json:"url"`
	Width int64  `json:"width"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Height            respjson.Field
		InlineStyleRanges respjson.Field
		PreviewURL        respjson.Field
		Text              respjson.Field
		Type              respjson.Field
		URL               respjson.Field
		Width             respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XGetArticleResponseArticleContent) RawJSON() string { return r.JSON.raw }
func (r *XGetArticleResponseArticleContent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type XGetArticleResponseArticleContentInlineStyleRange struct {
	Length int64  `json:"length"`
	Offset int64  `json:"offset"`
	Style  string `json:"style"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Length      respjson.Field
		Offset      respjson.Field
		Style       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XGetArticleResponseArticleContentInlineStyleRange) RawJSON() string { return r.JSON.raw }
func (r *XGetArticleResponseArticleContentInlineStyleRange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// X Article author profile fields returned when available.
type XGetArticleResponseAuthor struct {
	ID               string `json:"id" api:"required"`
	Name             string `json:"name" api:"required"`
	Username         string `json:"username" api:"required"`
	CanDm            bool   `json:"canDm"`
	CreatedAt        string `json:"createdAt"`
	Description      string `json:"description"`
	FavouritesCount  int64  `json:"favouritesCount"`
	FollowersCount   int64  `json:"followersCount"`
	FollowingCount   int64  `json:"followingCount"`
	IsBlueVerified   bool   `json:"isBlueVerified"`
	IsTranslator     bool   `json:"isTranslator"`
	IsVerified       bool   `json:"isVerified"`
	Location         string `json:"location"`
	MediaCount       int64  `json:"mediaCount"`
	ProfileBannerURL string `json:"profileBannerUrl"`
	ProfilePicture   string `json:"profilePicture"`
	Protected        bool   `json:"protected"`
	StatusesCount    int64  `json:"statusesCount"`
	URL              string `json:"url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		Name             respjson.Field
		Username         respjson.Field
		CanDm            respjson.Field
		CreatedAt        respjson.Field
		Description      respjson.Field
		FavouritesCount  respjson.Field
		FollowersCount   respjson.Field
		FollowingCount   respjson.Field
		IsBlueVerified   respjson.Field
		IsTranslator     respjson.Field
		IsVerified       respjson.Field
		Location         respjson.Field
		MediaCount       respjson.Field
		ProfileBannerURL respjson.Field
		ProfilePicture   respjson.Field
		Protected        respjson.Field
		StatusesCount    respjson.Field
		URL              respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XGetArticleResponseAuthor) RawJSON() string { return r.JSON.raw }
func (r *XGetArticleResponseAuthor) UnmarshalJSON(data []byte) error {
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

type XGetTrendsResponse struct {
	Count  int64                     `json:"count" api:"required"`
	Trends []XGetTrendsResponseTrend `json:"trends" api:"required"`
	Woeid  int64                     `json:"woeid" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Trends      respjson.Field
		Woeid       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XGetTrendsResponse) RawJSON() string { return r.JSON.raw }
func (r *XGetTrendsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type XGetTrendsResponseTrend struct {
	Name        string `json:"name" api:"required"`
	Description string `json:"description"`
	Query       string `json:"query"`
	Rank        int64  `json:"rank"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Description respjson.Field
		Query       respjson.Field
		Rank        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XGetTrendsResponseTrend) RawJSON() string { return r.JSON.raw }
func (r *XGetTrendsResponseTrend) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type XGetHomeTimelineParams struct {
	// Pagination cursor for timeline
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Comma-separated tweet IDs to exclude from results. Empty entries are ignored.
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
	// Pagination cursor for notifications
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Notification type filter. Unrecognized values fall back to All.
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

// Notification type filter. Unrecognized values fall back to All.
type XGetNotificationsParamsType string

const (
	XGetNotificationsParamsTypeAll      XGetNotificationsParamsType = "All"
	XGetNotificationsParamsTypeVerified XGetNotificationsParamsType = "Verified"
	XGetNotificationsParamsTypeMentions XGetNotificationsParamsType = "Mentions"
)

type XGetTrendsParams struct {
	// Number of trending topics to return (1-50, default 30)
	Count param.Opt[int64] `query:"count,omitzero" json:"-"`
	// Region WOEID (1=Worldwide, 23424977=US, 23424975=UK, 23424969=Turkey)
	Woeid param.Opt[int64] `query:"woeid,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [XGetTrendsParams]'s query parameters as `url.Values`.
func (r XGetTrendsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
