// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

package xtwitterscraper

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

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
// You should not instantiate this service directly, and instead use the
// [NewXService] method instead.
type XService struct {
	options []option.RequestOption
	// X write actions (tweets, likes, follows, DMs)
	WriteActions XWriteActionService
	// AI analysis of posts, at 2 credits per analyzed post
	Analysis XAnalysisService
	Tweets   XTweetService
	Users    XUserService
	// Look up, search, and explore user profiles and relationships
	Followers XFollowerService
	// Public X job listings, search and details
	Jobs  XJobService
	Dm    XDmService
	Media XMediaService
	// X write actions (tweets, likes, follows, DMs)
	Profile     XProfileService
	Communities XCommunityService
	// Connected X account management
	Accounts XAccountService
	// Connected X account management
	AccountConnectionChallenges XAccountConnectionChallengeService
	// Connected X account management
	AccountConnectionAttempts XAccountConnectionAttemptService
	// Look up, search, and analyze individual tweets
	Bookmarks XBookmarkService
	// X List followers, members, and tweets
	Lists XListService
	// X live video details, watch counts, and broadcaster
	Broadcasts XBroadcastService
	// Community Notes contributors, their notes, ratings & deleted notes
	CommunityNotes XCommunityNoteService
	// X Space details, host, and speakers
	Spaces XSpaceService
	// NFL & college football games, scoreboards, box scores & play-by-play
	Sports XSportService
}

// NewXService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewXService(opts ...option.RequestOption) (r XService) {
	r = XService{}
	r.options = opts
	r.WriteActions = NewXWriteActionService(opts...)
	r.Analysis = NewXAnalysisService(opts...)
	r.Tweets = NewXTweetService(opts...)
	r.Users = NewXUserService(opts...)
	r.Followers = NewXFollowerService(opts...)
	r.Jobs = NewXJobService(opts...)
	r.Dm = NewXDmService(opts...)
	r.Media = NewXMediaService(opts...)
	r.Profile = NewXProfileService(opts...)
	r.Communities = NewXCommunityService(opts...)
	r.Accounts = NewXAccountService(opts...)
	r.AccountConnectionChallenges = NewXAccountConnectionChallengeService(opts...)
	r.AccountConnectionAttempts = NewXAccountConnectionAttemptService(opts...)
	r.Bookmarks = NewXBookmarkService(opts...)
	r.Lists = NewXListService(opts...)
	r.Broadcasts = NewXBroadcastService(opts...)
	r.CommunityNotes = NewXCommunityNoteService(opts...)
	r.Spaces = NewXSpaceService(opts...)
	r.Sports = NewXSportService(opts...)
	return
}

// Returns the accounts, topics, hashtags & cashtags X suggests while someone types
// a search. X may rank them differently per request. Costs 1 credit per
// suggestion.
func (r *XService) AutocompleteSearch(ctx context.Context, query XAutocompleteSearchParams, opts ...option.RequestOption) (res *XAutocompleteSearchResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "x/search/autocomplete"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
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

// Returns X's hashflags, the custom emoji X shows after some hashtags. Lists the
// ones running now unless activeOnly is false. Each returned hashflag costs 1
// credit.
func (r *XService) GetHashflags(ctx context.Context, query XGetHashflagsParams, opts ...option.RequestOption) (res *XGetHashflagsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "x/hashflags"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Returns the home timeline for the connected X account.
func (r *XService) GetHomeTimeline(ctx context.Context, query XGetHomeTimelineParams, opts ...option.RequestOption) (res *shared.PaginatedTweets, err error) {
	opts = slices.Concat(r.options, opts)
	path := "x/timeline"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Returns each version X wrote of 1 news story, the AI trends x.com shows under
// Explore & News. Each version has its time, headline & text, oldest first, so you
// can see how the story changed. X lists every version at once. Each request costs
// 1 credit.
func (r *XService) GetNewsStoryVersions(ctx context.Context, id string, query XGetNewsStoryVersionsParams, opts ...option.RequestOption) (res *XGetNewsStoryVersionsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("x/news/%s/versions", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Returns notifications for the connected X account.
func (r *XService) GetNotifications(ctx context.Context, query XGetNotificationsParams, opts ...option.RequestOption) (res *XGetNotificationsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "x/notifications"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Returns the places X offers trends for. Pass a location's woeid to /x/trends.
// Each returned location costs 1 credit.
func (r *XService) GetTrendLocations(ctx context.Context, query XGetTrendLocationsParams, opts ...option.RequestOption) (res *XGetTrendLocationsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "x/trends/locations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Returns X's current trends for a place you choose. Pick it by woeid, country or
// location name. Without one, trends are worldwide. Add at to read what was
// trending on Twitter at a past time, such as yesterday. Xquik keeps a Twitter
// trends archive from 2026-10-13 02:20 UTC on, with 1 list per place every 5
// minutes. Each request costs 1 credit, however many trends it returns.
func (r *XService) GetTrends(ctx context.Context, query XGetTrendsParams, opts ...option.RequestOption) (res *XGetTrendsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "x/trends"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Returns the Twitter trends archive of a place, as the lists X showed there over
// a range of time. Use it for trending on Twitter yesterday, last week or any day
// since 2026-10-13 02:20 UTC. Xquik stores 1 list per place every 5 minutes and
// keeps it for 25 months. Pick the place by woeid, country or location name, the
// range with from & to, and the step with every. from & to take a time or a span
// back from now, such as 7d, so a schedule can read the last day on every run.
// Each step returns the first list stored in it, with its own time. Each returned
// list costs 1 credit, so 24 lists cost 24 credits. A step with no stored list
// returns none and costs nothing. 1 request returns 24 lists, or up to 100 with
// limit.
func (r *XService) GetTrendsHistory(ctx context.Context, query XGetTrendsHistoryParams, opts ...option.RequestOption) (res *XGetTrendsHistoryResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "x/trends/history"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Finds places on X by name. Pass a place's id as place to /x/tweets/search to
// read posts tagged there. X ranks the places, so their order can differ between
// requests. Each returned place costs 1 credit.
func (r *XService) SearchPlaces(ctx context.Context, query XSearchPlacesParams, opts ...option.RequestOption) (res *XSearchPlacesResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "x/places/search"
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
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type XGetArticleResponseArticle struct {
	// X's ID for the Article, which differs from its post's ID
	ArticleID string `json:"articleId"`
	// Article blocks as Markdown, with headings, lists, bold, italic, links, images &
	// embedded posts as their x.com URLs. Holds no HTML.
	BodyMarkdown string `json:"bodyMarkdown"`
	// Plain text joined from all article blocks
	BodyText string `json:"bodyText"`
	// Accounts that bookmarked the Article's post. 0 or more.
	BookmarkCount int64                               `json:"bookmarkCount"`
	Contents      []XGetArticleResponseArticleContent `json:"contents"`
	CoverImageURL string                              `json:"coverImageUrl"`
	// Public metadata whose fields X defines.
	CoverMedia PublicMetadata `json:"coverMedia"`
	CreatedAt  string         `json:"createdAt"`
	// Public metadata whose fields X defines.
	LifecycleState PublicMetadata `json:"lifecycleState"`
	LikeCount      int64          `json:"likeCount"`
	// Public metadata whose fields X defines.
	Metadata    PublicMetadata `json:"metadata"`
	PreviewText string         `json:"previewText"`
	QuoteCount  int64          `json:"quoteCount"`
	ReplyCount  int64          `json:"replyCount"`
	// Reposts of the Article's post. 0 or more.
	RetweetCount int64 `json:"retweetCount"`
	// X's summary of the Article, when X has one
	SummaryText string `json:"summaryText"`
	Title       string `json:"title"`
	ViewCount   int64  `json:"viewCount"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ArticleID      respjson.Field
		BodyMarkdown   respjson.Field
		BodyText       respjson.Field
		BookmarkCount  respjson.Field
		Contents       respjson.Field
		CoverImageURL  respjson.Field
		CoverMedia     respjson.Field
		CreatedAt      respjson.Field
		LifecycleState respjson.Field
		LikeCount      respjson.Field
		Metadata       respjson.Field
		PreviewText    respjson.Field
		QuoteCount     respjson.Field
		ReplyCount     respjson.Field
		RetweetCount   respjson.Field
		SummaryText    respjson.Field
		Title          respjson.Field
		ViewCount      respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XGetArticleResponseArticle) RawJSON() string { return r.JSON.raw }
func (r *XGetArticleResponseArticle) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type XGetArticleResponseArticleContent struct {
	// Alt text of the image or preview image
	AltText string `json:"altText"`
	// Caption the author wrote under a media block
	Caption string `json:"caption"`
	// Length of a video, in milliseconds
	DurationMs int64 `json:"durationMs"`
	Height     int64 `json:"height"`
	// Inline text formatting ranges
	InlineStyleRanges []XGetArticleResponseArticleContentInlineStyleRange `json:"inlineStyleRanges"`
	// Links in the block's text, with the text range each covers
	Links []XGetArticleResponseArticleContentLink `json:"links"`
	// What a media block holds
	//
	// Any of "image", "video", "gif".
	MediaType XGetArticleResponseArticleContentMediaType `json:"mediaType"`
	// Preview image URL for media blocks
	PreviewURL string `json:"previewUrl"`
	Text       string `json:"text"`
	// ID of the post a tweet block embeds
	TweetID string `json:"tweetId"`
	// Block type: paragraph, header-one, header-two, header-three, header-four,
	// header-five, header-six, unordered-list-item, ordered-list-item, blockquote,
	// code-block, media, divider
	Type string `json:"type"`
	// Media URL for media blocks
	URL string `json:"url"`
	// Files of a video or GIF block, in X's order
	Variants []shared.TweetMediaVideoVariant `json:"variants"`
	Width    int64                           `json:"width"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AltText           respjson.Field
		Caption           respjson.Field
		DurationMs        respjson.Field
		Height            respjson.Field
		InlineStyleRanges respjson.Field
		Links             respjson.Field
		MediaType         respjson.Field
		PreviewURL        respjson.Field
		Text              respjson.Field
		TweetID           respjson.Field
		Type              respjson.Field
		URL               respjson.Field
		Variants          respjson.Field
		Width             respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XGetArticleResponseArticleContent) RawJSON() string { return r.JSON.raw }
func (r *XGetArticleResponseArticleContent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type XGetArticleResponseArticleContentLink struct {
	Length int64  `json:"length"`
	Offset int64  `json:"offset"`
	URL    string `json:"url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Length      respjson.Field
		Offset      respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XGetArticleResponseArticleContentLink) RawJSON() string { return r.JSON.raw }
func (r *XGetArticleResponseArticleContentLink) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// What a media block holds
type XGetArticleResponseArticleContentMediaType string

const (
	XGetArticleResponseArticleContentMediaTypeImage XGetArticleResponseArticleContentMediaType = "image"
	XGetArticleResponseArticleContentMediaTypeVideo XGetArticleResponseArticleContentMediaType = "video"
	XGetArticleResponseArticleContentMediaTypeGif   XGetArticleResponseArticleContentMediaType = "gif"
)

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
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// X Article author profile fields returned when available.
type XGetArticleResponseAuthor struct {
	ID               string `json:"id" api:"required"`
	Name             string `json:"name" api:"required"`
	Username         string `json:"username" api:"required"`
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
	// The same profile picture at full size.
	ProfilePictureOriginal string `json:"profilePictureOriginal"`
	Protected              bool   `json:"protected"`
	StatusesCount          int64  `json:"statusesCount"`
	URL                    string `json:"url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                     respjson.Field
		Name                   respjson.Field
		Username               respjson.Field
		CreatedAt              respjson.Field
		Description            respjson.Field
		FavouritesCount        respjson.Field
		FollowersCount         respjson.Field
		FollowingCount         respjson.Field
		IsBlueVerified         respjson.Field
		IsTranslator           respjson.Field
		IsVerified             respjson.Field
		Location               respjson.Field
		MediaCount             respjson.Field
		ProfileBannerURL       respjson.Field
		ProfilePicture         respjson.Field
		ProfilePictureOriginal respjson.Field
		Protected              respjson.Field
		StatusesCount          respjson.Field
		URL                    respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XGetArticleResponseAuthor) RawJSON() string { return r.JSON.raw }
func (r *XGetArticleResponseAuthor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
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
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
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
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type XGetTrendsResponse struct {
	Count  int64                     `json:"count" api:"required"`
	Trends []XGetTrendsResponseTrend `json:"trends" api:"required"`
	Woeid  int64                     `json:"woeid" api:"required"`
	// When Xquik read this list from X. It comes from the stored list and is never
	// estimated.
	CapturedAt time.Time `json:"capturedAt" format:"date-time"`
	// The place country or location named. Absent when woeid picked the place or the
	// request named none.
	Location XGetTrendsResponseLocation `json:"location"`
	// The time at asked for, as Xquik read it. A span such as 24h shows here as the
	// time it counted back to. Present only with at.
	RequestedAt time.Time `json:"requestedAt" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Trends      respjson.Field
		Woeid       respjson.Field
		CapturedAt  respjson.Field
		Location    respjson.Field
		RequestedAt respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XGetTrendsResponse) RawJSON() string { return r.JSON.raw }
func (r *XGetTrendsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type XGetTrendsResponseTrend struct {
	Name        string `json:"name" api:"required"`
	Description string `json:"description"`
	// Promotion identifier from X. Null for organic trends.
	PromotedContent string `json:"promotedContent" api:"nullable"`
	Query           string `json:"query"`
	Rank            int64  `json:"rank"`
	// Approximate public post volume when X supplies it.
	TweetVolume int64 `json:"tweetVolume" api:"nullable"`
	// X search URL for the trend.
	URL string `json:"url" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name            respjson.Field
		Description     respjson.Field
		PromotedContent respjson.Field
		Query           respjson.Field
		Rank            respjson.Field
		TweetVolume     respjson.Field
		URL             respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XGetTrendsResponseTrend) RawJSON() string { return r.JSON.raw }
func (r *XGetTrendsResponseTrend) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// The place country or location named. Absent when woeid picked the place or the
// request named none.
type XGetTrendsResponseLocation struct {
	// Country name. Null for Worldwide.
	Country string `json:"country" api:"required"`
	// 2-letter country code. Null for Worldwide.
	CountryCode string `json:"countryCode" api:"required"`
	Name        string `json:"name" api:"required"`
	// WOEID of the country or region above it. Null for Worldwide.
	ParentWoeid int64 `json:"parentWoeid" api:"required"`
	// X's kind of place. Town, Country, Supername for Worldwide, or Unknown.
	PlaceType string `json:"placeType" api:"required"`
	// X's numeric code for the kind of place.
	PlaceTypeCode int64 `json:"placeTypeCode" api:"required"`
	// X's legacy Yahoo GeoPlanet reference. It may not resolve.
	URL string `json:"url" api:"required"`
	// Location ID. Pass it as woeid to /x/trends.
	Woeid int64 `json:"woeid" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Country       respjson.Field
		CountryCode   respjson.Field
		Name          respjson.Field
		ParentWoeid   respjson.Field
		PlaceType     respjson.Field
		PlaceTypeCode respjson.Field
		URL           respjson.Field
		Woeid         respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XGetTrendsResponseLocation) RawJSON() string { return r.JSON.raw }
func (r *XGetTrendsResponseLocation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// What X suggests for a partial search query. Each list keeps X's order. A list is
// empty when X suggests none of its type or types leaves it out.
type XAutocompleteSearchResponse struct {
	// Suggested stocks & crypto tokens.
	Cashtags []XAutocompleteSearchResponseCashtag `json:"cashtags" api:"required"`
	// Suggestions returned across all 4 lists. Each costs 1 credit.
	Count int64 `json:"count" api:"required"`
	// Suggested hashtags.
	Hashtags []XAutocompleteSearchResponseHashtag `json:"hashtags" api:"required"`
	// The query you sent, trimmed.
	Query string `json:"query" api:"required"`
	// Suggested search phrases that are not hashtags.
	Topics []XAutocompleteSearchResponseTopic `json:"topics" api:"required"`
	// Suggested accounts.
	Users []XAutocompleteSearchResponseUser `json:"users" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Cashtags    respjson.Field
		Count       respjson.Field
		Hashtags    respjson.Field
		Query       respjson.Field
		Topics      respjson.Field
		Users       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XAutocompleteSearchResponse) RawJSON() string { return r.JSON.raw }
func (r *XAutocompleteSearchResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// A stock or crypto token X suggests for the query.
type XAutocompleteSearchResponseCashtag struct {
	// Ticker symbol without the $.
	Ticker string `json:"ticker" api:"required"`
	// Logo of the token's chain. Tokens only.
	ChainLogoURL string `json:"chainLogoUrl" format:"uri"`
	// Token contract address. Tokens only.
	ContractAddress string `json:"contractAddress"`
	// Currency of price & marketCap.
	Currency string `json:"currency"`
	// Price change over the last 24 hours, in percent.
	DailyChangePercent float64 `json:"dailyChangePercent"`
	// Stock exchange. Stocks only.
	ExchangeShortName string `json:"exchangeShortName"`
	// Logo image URL.
	LogoURL string `json:"logoUrl" format:"uri"`
	// Market capitalization as decimal text.
	MarketCap string `json:"marketCap"`
	// Company or token name.
	Name string `json:"name"`
	// Latest price as decimal text, which keeps every digit.
	Price string `json:"price"`
	// X's tag for the asset. $TICKER for a stock, chain:contract for a token.
	SmartTag string `json:"smartTag"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Ticker             respjson.Field
		ChainLogoURL       respjson.Field
		ContractAddress    respjson.Field
		Currency           respjson.Field
		DailyChangePercent respjson.Field
		ExchangeShortName  respjson.Field
		LogoURL            respjson.Field
		MarketCap          respjson.Field
		Name               respjson.Field
		Price              respjson.Field
		SmartTag           respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XAutocompleteSearchResponseCashtag) RawJSON() string { return r.JSON.raw }
func (r *XAutocompleteSearchResponseCashtag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// A hashtag X suggests for the query.
type XAutocompleteSearchResponseHashtag struct {
	// Suggested hashtag, with its
	Hashtag string `json:"hashtag" api:"required"`
	// X's inline flag for the hashtag.
	Inline bool `json:"inline"`
	// X's rounded relevance score for the hashtag.
	Score float64 `json:"score"`
	// Normalized text tokens X matches the hashtag by.
	Tokens []string `json:"tokens"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Hashtag     respjson.Field
		Inline      respjson.Field
		Score       respjson.Field
		Tokens      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XAutocompleteSearchResponseHashtag) RawJSON() string { return r.JSON.raw }
func (r *XAutocompleteSearchResponseHashtag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// A search phrase X suggests for the query.
type XAutocompleteSearchResponseTopic struct {
	// Suggested search phrase.
	Topic string `json:"topic" api:"required"`
	// X's inline flag for the phrase.
	Inline bool `json:"inline"`
	// X's rounded relevance score for the phrase.
	Score float64 `json:"score"`
	// Normalized text tokens X matches the phrase by.
	Tokens []string `json:"tokens"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Topic       respjson.Field
		Inline      respjson.Field
		Score       respjson.Field
		Tokens      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XAutocompleteSearchResponseTopic) RawJSON() string { return r.JSON.raw }
func (r *XAutocompleteSearchResponseTopic) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// An account X suggests for the query.
type XAutocompleteSearchResponseUser struct {
	// X user ID.
	ID string `json:"id" api:"required"`
	// Display name.
	Name string `json:"name" api:"required"`
	// Handle without the @.
	Username string `json:"username" api:"required"`
	// Affiliation badges X shows beside the name.
	Badges []XAutocompleteSearchResponseUserBadge `json:"badges"`
	// X's inline flag for the account.
	Inline bool `json:"inline"`
	// Whether X shows a blue verification badge.
	IsBlueVerified bool `json:"isBlueVerified"`
	// X's flag for whether media can be generated from this account's persona.
	IsPersonaMediaGenable bool `json:"isPersonaMediaGenable"`
	// Whether the account holds X's legacy verification.
	IsVerified bool `json:"isVerified"`
	// Public profile location text. Omitted when empty.
	Location string `json:"location"`
	// Avatar URL.
	ProfilePicture string `json:"profilePicture" format:"uri"`
	// The same avatar at full size.
	ProfilePictureOriginal string `json:"profilePictureOriginal" format:"uri"`
	// Whether the account protects its posts.
	Protected bool `json:"protected"`
	// X's rounded relevance score for the account.
	Score float64 `json:"score"`
	// Normalized text tokens X matches the account by.
	Tokens []string `json:"tokens"`
	// Whether X shows any verification badge.
	Verified bool `json:"verified"`
	// Verification category, such as Government or Business. Omitted when X names
	// none.
	VerifiedType string `json:"verifiedType"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                     respjson.Field
		Name                   respjson.Field
		Username               respjson.Field
		Badges                 respjson.Field
		Inline                 respjson.Field
		IsBlueVerified         respjson.Field
		IsPersonaMediaGenable  respjson.Field
		IsVerified             respjson.Field
		Location               respjson.Field
		ProfilePicture         respjson.Field
		ProfilePictureOriginal respjson.Field
		Protected              respjson.Field
		Score                  respjson.Field
		Tokens                 respjson.Field
		Verified               respjson.Field
		VerifiedType           respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XAutocompleteSearchResponseUser) RawJSON() string { return r.JSON.raw }
func (r *XAutocompleteSearchResponseUser) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// 1 affiliation badge.
type XAutocompleteSearchResponseUserBadge struct {
	// Kind of badge.
	BadgeType string `json:"badgeType"`
	// Badge image URL.
	BadgeURL string `json:"badgeUrl" format:"uri"`
	// Name of the affiliated account.
	Description string `json:"description"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BadgeType   respjson.Field
		BadgeURL    respjson.Field
		Description respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XAutocompleteSearchResponseUserBadge) RawJSON() string { return r.JSON.raw }
func (r *XAutocompleteSearchResponseUserBadge) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type XGetHashflagsResponse struct {
	// Hashflags returned & charged.
	Count     int64                           `json:"count" api:"required"`
	Hashflags []XGetHashflagsResponseHashflag `json:"hashflags" api:"required"`
	// Hashflags that match the filters, before limit.
	Total int64 `json:"total" api:"required"`
	// What to send instead. Present only when no hashflag matches.
	Message string `json:"message"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Hashflags   respjson.Field
		Total       respjson.Field
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XGetHashflagsResponse) RawJSON() string { return r.JSON.raw }
func (r *XGetHashflagsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// A custom emoji X shows after a hashtag for a set time.
type XGetHashflagsResponseHashflag struct {
	// Animations X plays for the hashtag, such as on a like. Often empty.
	Animations []XGetHashflagsResponseHashflagAnimation `json:"animations" api:"required"`
	// Image of the emoji.
	AssetURL string `json:"assetUrl" api:"required" format:"uri"`
	// When X stops showing the emoji, in UTC.
	EndsAt time.Time `json:"endsAt" api:"required" format:"date-time"`
	// Hashtag text without the hash sign. X also lists some plain phrases.
	Hashtag string `json:"hashtag" api:"required"`
	// Whether X marks the hashflag for its hashfetti effect.
	IsHashfettiEnabled bool `json:"isHashfettiEnabled" api:"required"`
	// When X starts showing the emoji, in UTC.
	StartsAt time.Time `json:"startsAt" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Animations         respjson.Field
		AssetURL           respjson.Field
		EndsAt             respjson.Field
		Hashtag            respjson.Field
		IsHashfettiEnabled respjson.Field
		StartsAt           respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XGetHashflagsResponseHashflag) RawJSON() string { return r.JSON.raw }
func (r *XGetHashflagsResponseHashflag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type XGetHashflagsResponseHashflagAnimation struct {
	// Animation file.
	AssetURL string `json:"assetUrl" api:"required" format:"uri"`
	// What plays the animation, such as Like.
	Context string `json:"context" api:"required"`
	// X's order among the animations of a hashtag.
	Priority float64 `json:"priority" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AssetURL    respjson.Field
		Context     respjson.Field
		Priority    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XGetHashflagsResponseHashflagAnimation) RawJSON() string { return r.JSON.raw }
func (r *XGetHashflagsResponseHashflagAnimation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type XGetNewsStoryVersionsResponse struct {
	// The story's ID, as the request named it.
	ID string `json:"id" api:"required"`
	// Versions returned.
	Count    int64                                  `json:"count" api:"required"`
	Versions []XGetNewsStoryVersionsResponseVersion `json:"versions" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Count       respjson.Field
		Versions    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XGetNewsStoryVersionsResponse) RawJSON() string { return r.JSON.raw }
func (r *XGetNewsStoryVersionsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// 1 version X wrote of a news story.
type XGetNewsStoryVersionsResponseVersion struct {
	// The version's ID.
	ID string `json:"id" api:"required"`
	// When X wrote the version, in UTC, from its ID.
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// The story's headline in this version.
	Headline string `json:"headline" api:"required"`
	// The story's text in this version, without the time line X writes above it in the
	// reader's time zone.
	Text string `json:"text" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Headline    respjson.Field
		Text        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XGetNewsStoryVersionsResponseVersion) RawJSON() string { return r.JSON.raw }
func (r *XGetNewsStoryVersionsResponseVersion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type XGetTrendLocationsResponse struct {
	// Locations returned & charged.
	Count     int64                                `json:"count" api:"required"`
	Locations []XGetTrendLocationsResponseLocation `json:"locations" api:"required"`
	// Locations that match the filters, before limit.
	Total int64 `json:"total" api:"required"`
	// What to send instead. Present only when no location matches.
	Message string `json:"message"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Locations   respjson.Field
		Total       respjson.Field
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XGetTrendLocationsResponse) RawJSON() string { return r.JSON.raw }
func (r *XGetTrendLocationsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// A place X offers trends for.
type XGetTrendLocationsResponseLocation struct {
	// Country name. Null for Worldwide.
	Country string `json:"country" api:"required"`
	// 2-letter country code. Null for Worldwide.
	CountryCode string `json:"countryCode" api:"required"`
	Name        string `json:"name" api:"required"`
	// WOEID of the country or region above it. Null for Worldwide.
	ParentWoeid int64 `json:"parentWoeid" api:"required"`
	// X's kind of place. Town, Country, Supername for Worldwide, or Unknown.
	PlaceType string `json:"placeType" api:"required"`
	// X's numeric code for the kind of place.
	PlaceTypeCode int64 `json:"placeTypeCode" api:"required"`
	// X's legacy Yahoo GeoPlanet reference. It may not resolve.
	URL string `json:"url" api:"required"`
	// Location ID. Pass it as woeid to /x/trends.
	Woeid int64 `json:"woeid" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Country       respjson.Field
		CountryCode   respjson.Field
		Name          respjson.Field
		ParentWoeid   respjson.Field
		PlaceType     respjson.Field
		PlaceTypeCode respjson.Field
		URL           respjson.Field
		Woeid         respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XGetTrendLocationsResponseLocation) RawJSON() string { return r.JSON.raw }
func (r *XGetTrendLocationsResponseLocation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type XGetTrendsHistoryResponse struct {
	// Lists returned & charged.
	Count int64 `json:"count" api:"required"`
	// Start of the range as Xquik read it. A span such as 7d shows here as the time it
	// counted back to.
	FromTime time.Time `json:"fromTime" api:"required" format:"date-time"`
	// Whether the range holds more lists.
	HasNextPage bool `json:"has_next_page" api:"required"`
	// Send it as cursor for the next lists. Empty on the last page.
	NextCursor string `json:"next_cursor" api:"required"`
	// 1 stored list per step, oldest first. A step with no stored list is left out.
	Snapshots []XGetTrendsHistoryResponseSnapshot `json:"snapshots" api:"required"`
	// End of the range as Xquik read it, which the range leaves out. Without to, it is
	// the time of the first request.
	ToTime time.Time `json:"toTime" api:"required" format:"date-time"`
	Woeid  int64     `json:"woeid" api:"required"`
	// The place the request named by country or location. Absent when it sent a woeid
	// or named no place.
	Location XGetTrendsHistoryResponseLocation `json:"location"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		FromTime    respjson.Field
		HasNextPage respjson.Field
		NextCursor  respjson.Field
		Snapshots   respjson.Field
		ToTime      respjson.Field
		Woeid       respjson.Field
		Location    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XGetTrendsHistoryResponse) RawJSON() string { return r.JSON.raw }
func (r *XGetTrendsHistoryResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type XGetTrendsHistoryResponseSnapshot struct {
	// When Xquik read this stored list from X. Xquik never estimates it.
	CapturedAt time.Time `json:"capturedAt" api:"required" format:"date-time"`
	// Trends in this list.
	Count  int64                                    `json:"count" api:"required"`
	Trends []XGetTrendsHistoryResponseSnapshotTrend `json:"trends" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CapturedAt  respjson.Field
		Count       respjson.Field
		Trends      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XGetTrendsHistoryResponseSnapshot) RawJSON() string { return r.JSON.raw }
func (r *XGetTrendsHistoryResponseSnapshot) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type XGetTrendsHistoryResponseSnapshotTrend struct {
	Name        string `json:"name" api:"required"`
	Description string `json:"description"`
	// X's promotion ID of a promoted trend. Null for the others.
	PromotedContent string `json:"promotedContent" api:"nullable"`
	Query           string `json:"query"`
	// The trend's place in X's list at that time.
	Rank int64 `json:"rank"`
	// X's post count for the trend at that time, when X showed one.
	TweetVolume int64 `json:"tweetVolume" api:"nullable"`
	// The trend's search link on X.
	URL string `json:"url" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name            respjson.Field
		Description     respjson.Field
		PromotedContent respjson.Field
		Query           respjson.Field
		Rank            respjson.Field
		TweetVolume     respjson.Field
		URL             respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XGetTrendsHistoryResponseSnapshotTrend) RawJSON() string { return r.JSON.raw }
func (r *XGetTrendsHistoryResponseSnapshotTrend) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// The place the request named by country or location. Absent when it sent a woeid
// or named no place.
type XGetTrendsHistoryResponseLocation struct {
	// Country name. Null for Worldwide.
	Country string `json:"country" api:"required"`
	// 2-letter country code. Null for Worldwide.
	CountryCode string `json:"countryCode" api:"required"`
	Name        string `json:"name" api:"required"`
	// WOEID of the country or region above it. Null for Worldwide.
	ParentWoeid int64 `json:"parentWoeid" api:"required"`
	// X's kind of place. Town, Country, Supername for Worldwide, or Unknown.
	PlaceType string `json:"placeType" api:"required"`
	// X's numeric code for the kind of place.
	PlaceTypeCode int64 `json:"placeTypeCode" api:"required"`
	// X's legacy Yahoo GeoPlanet reference. It may not resolve.
	URL string `json:"url" api:"required"`
	// Location ID. Pass it as woeid to /x/trends.
	Woeid int64 `json:"woeid" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Country       respjson.Field
		CountryCode   respjson.Field
		Name          respjson.Field
		ParentWoeid   respjson.Field
		PlaceType     respjson.Field
		PlaceTypeCode respjson.Field
		URL           respjson.Field
		Woeid         respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XGetTrendsHistoryResponseLocation) RawJSON() string { return r.JSON.raw }
func (r *XGetTrendsHistoryResponseLocation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type XSearchPlacesResponse struct {
	// Places returned & charged.
	Count  int64                        `json:"count" api:"required"`
	Places []XSearchPlacesResponsePlace `json:"places" api:"required"`
	// Places X found, before limit.
	Total int64 `json:"total" api:"required"`
	// What to send instead. Present only when X finds no place.
	Message string `json:"message"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Places      respjson.Field
		Total       respjson.Field
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XSearchPlacesResponse) RawJSON() string { return r.JSON.raw }
func (r *XSearchPlacesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// A place X found, with the larger places it lies in.
type XSearchPlacesResponsePlace struct {
	// Place ID. Pass it as place to /x/tweets/search.
	ID string `json:"id" api:"required"`
	// Extra details X attaches to some places. Often empty.
	Attributes PublicMetadata `json:"attributes" api:"required"`
	// Outline of the place as GeoJSON. Each point holds longitude, then latitude.
	BoundingBox XSearchPlacesResponsePlaceBoundingBox `json:"boundingBox" api:"required"`
	// Center of the place. Longitude, then latitude.
	Centroid []float64 `json:"centroid" api:"required"`
	// Larger places this place lies in. Often 1.
	ContainedWithin []XSearchPlacesResponsePlaceContainedWithin `json:"containedWithin" api:"required"`
	// Country the place lies in.
	Country string `json:"country" api:"required"`
	// 2-letter code of the country the place lies in.
	CountryCode string `json:"countryCode" api:"required"`
	// Name with the region or country around the place.
	FullName string `json:"fullName" api:"required"`
	// Short name of the place.
	Name string `json:"name" api:"required"`
	// X's kind of place, such as city, admin, neighborhood, poi or country.
	PlaceType string `json:"placeType" api:"required"`
	// X's own reference URL for the place.
	URL string `json:"url" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		Attributes      respjson.Field
		BoundingBox     respjson.Field
		Centroid        respjson.Field
		ContainedWithin respjson.Field
		Country         respjson.Field
		CountryCode     respjson.Field
		FullName        respjson.Field
		Name            respjson.Field
		PlaceType       respjson.Field
		URL             respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XSearchPlacesResponsePlace) RawJSON() string { return r.JSON.raw }
func (r *XSearchPlacesResponsePlace) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// Outline of the place as GeoJSON. Each point holds longitude, then latitude.
type XSearchPlacesResponsePlaceBoundingBox struct {
	// Rings of points that draw the outline.
	Coordinates [][][]float64 `json:"coordinates" api:"required"`
	// GeoJSON shape of the outline.
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Coordinates respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XSearchPlacesResponsePlaceBoundingBox) RawJSON() string { return r.JSON.raw }
func (r *XSearchPlacesResponsePlaceBoundingBox) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// A place on X, such as a city or a neighborhood.
type XSearchPlacesResponsePlaceContainedWithin struct {
	// Place ID. Pass it as place to /x/tweets/search.
	ID string `json:"id" api:"required"`
	// Extra details X attaches to some places. Often empty.
	Attributes PublicMetadata `json:"attributes" api:"required"`
	// Outline of the place as GeoJSON. Each point holds longitude, then latitude.
	BoundingBox XSearchPlacesResponsePlaceContainedWithinBoundingBox `json:"boundingBox" api:"required"`
	// Center of the place. Longitude, then latitude.
	Centroid []float64 `json:"centroid" api:"required"`
	// Country the place lies in.
	Country string `json:"country" api:"required"`
	// 2-letter code of the country the place lies in.
	CountryCode string `json:"countryCode" api:"required"`
	// Name with the region or country around the place.
	FullName string `json:"fullName" api:"required"`
	// Short name of the place.
	Name string `json:"name" api:"required"`
	// X's kind of place, such as city, admin, neighborhood, poi or country.
	PlaceType string `json:"placeType" api:"required"`
	// X's own reference URL for the place.
	URL string `json:"url" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Attributes  respjson.Field
		BoundingBox respjson.Field
		Centroid    respjson.Field
		Country     respjson.Field
		CountryCode respjson.Field
		FullName    respjson.Field
		Name        respjson.Field
		PlaceType   respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XSearchPlacesResponsePlaceContainedWithin) RawJSON() string { return r.JSON.raw }
func (r *XSearchPlacesResponsePlaceContainedWithin) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// Outline of the place as GeoJSON. Each point holds longitude, then latitude.
type XSearchPlacesResponsePlaceContainedWithinBoundingBox struct {
	// Rings of points that draw the outline.
	Coordinates [][][]float64 `json:"coordinates" api:"required"`
	// GeoJSON shape of the outline.
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Coordinates respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XSearchPlacesResponsePlaceContainedWithinBoundingBox) RawJSON() string { return r.JSON.raw }
func (r *XSearchPlacesResponsePlaceContainedWithinBoundingBox) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
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
	// A past time, to read what was trending on Twitter then. Takes an ISO 8601 time,
	// a date such as 2026-10-13, or Unix seconds. A time without a zone reads as UTC.
	// It also takes a span back from now, as a number & s, m, h or d. So at=24h reads
	// what was trending on Twitter yesterday at this hour, on every run of a schedule.
	// The answer is the stored list nearest that time, with its own time in
	// capturedAt. History starts 2026-10-13 02:20 UTC. A time over 24 hours from every
	// stored list answers 400 at no charge.
	At param.Opt[string] `query:"at,omitzero" json:"-"`
	// Number of trending topics to return (1-50, default 30)
	Count param.Opt[int64] `query:"count,omitzero" json:"-"`
	// Country to read trends for, by name or 2-letter code, such as Turkey or TR. Case
	// insensitive.
	Country param.Opt[string] `query:"country,omitzero" json:"-"`
	// Fields to keep in each row, comma-separated. Use dots for nested fields, such as
	// author.username. A list keeps the field in each item. Case & underscores do not
	// matter. A row that lacks a field leaves it out. Cursors, counts & diagnostics
	// always stay. An unknown name answers 400 unknown_field with the valid names,
	// before any read. The price stays the same. SDK row types then hold only the
	// named fields.
	Fields param.Opt[string] `query:"fields,omitzero" json:"-"`
	// Town or country to read trends for, by name, such as Istanbul. Case insensitive.
	// Add country when 2 towns share a name. /x/trends/locations lists every name.
	Location param.Opt[string] `query:"location,omitzero" json:"-"`
	// Fields to drop from each row, comma-separated. It takes the same names as
	// fields. With fields, it drops from the kept fields.
	Omit param.Opt[string] `query:"omit,omitzero" json:"-"`
	// Region WOEID (1=Worldwide, 23424977=US, 23424975=UK, 23424969=Turkey). Wins over
	// country & location.
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

type XAutocompleteSearchParams struct {
	// Text typed so far. Start with # for hashtags or $ for cashtags. A # query
	// returns hashtags & the accounts X suggests for it. Its alias is query.
	Q string `query:"q" api:"required" json:"-"`
	// Fields to keep in each row, comma-separated. Use dots for nested fields, such as
	// author.username. A list keeps the field in each item. Case & underscores do not
	// matter. A row that lacks a field leaves it out. Cursors, counts & diagnostics
	// always stay. An unknown name answers 400 unknown_field with the valid names,
	// before any read. The price stays the same. SDK row types then hold only the
	// named fields.
	Fields param.Opt[string] `query:"fields,omitzero" json:"-"`
	// Most suggestions to return, from 1 to 100. Default 20. Accounts fill it first,
	// then topics, hashtags & cashtags. A low credit balance can return fewer. Its
	// aliases are pageSize, count, max_results, maxItems, max_items & per_page.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Fields to drop from each row, comma-separated. It takes the same names as
	// fields. With fields, it drops from the kept fields.
	Omit param.Opt[string] `query:"omit,omitzero" json:"-"`
	// Comma-separated suggestion types from users, topics, hashtags & cashtags. Omit
	// for all. With only hashtags or only cashtags, the query gets its # or $ when it
	// lacks one.
	Types param.Opt[string] `query:"types,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [XAutocompleteSearchParams]'s query parameters as `url.Values`.
func (r XAutocompleteSearchParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type XGetHashflagsParams struct {
	// Keeps hashflags that run now. Send false to add ended & upcoming ones.
	ActiveOnly param.Opt[bool] `query:"activeOnly,omitzero" json:"-"`
	// Fields to keep in each row, comma-separated. Use dots for nested fields, such as
	// author.username. A list keeps the field in each item. Case & underscores do not
	// matter. A row that lacks a field leaves it out. Cursors, counts & diagnostics
	// always stay. An unknown name answers 400 unknown_field with the valid names,
	// before any read. The price stays the same. SDK row types then hold only the
	// named fields.
	Fields param.Opt[string] `query:"fields,omitzero" json:"-"`
	// Maximum hashflags to return. Omit it for every match. A low credit balance can
	// return fewer. Aliases pageSize, count, max_results, maxItems, max_items &
	// per_page.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Fields to drop from each row, comma-separated. It takes the same names as
	// fields. With fields, it drops from the kept fields.
	Omit param.Opt[string] `query:"omit,omitzero" json:"-"`
	// Text the hashtag contains, such as football. Case insensitive. A leading hash
	// sign is ignored. Alias query.
	Q param.Opt[string] `query:"q,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [XGetHashflagsParams]'s query parameters as `url.Values`.
func (r XGetHashflagsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type XGetNewsStoryVersionsParams struct {
	// Fields to keep in each row, comma-separated. Use dots for nested fields, such as
	// author.username. A list keeps the field in each item. Case & underscores do not
	// matter. A row that lacks a field leaves it out. Cursors, counts & diagnostics
	// always stay. An unknown name answers 400 unknown_field with the valid names,
	// before any read. The price stays the same. SDK row types then hold only the
	// named fields.
	Fields param.Opt[string] `query:"fields,omitzero" json:"-"`
	// Fields to drop from each row, comma-separated. It takes the same names as
	// fields. With fields, it drops from the kept fields.
	Omit param.Opt[string] `query:"omit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [XGetNewsStoryVersionsParams]'s query parameters as `url.Values`.
func (r XGetNewsStoryVersionsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type XGetTrendLocationsParams struct {
	// Country name or 2-letter code, such as Turkey or TR. Case insensitive. Returns
	// the country & its towns.
	Country param.Opt[string] `query:"country,omitzero" json:"-"`
	// Fields to keep in each row, comma-separated. Use dots for nested fields, such as
	// author.username. A list keeps the field in each item. Case & underscores do not
	// matter. A row that lacks a field leaves it out. Cursors, counts & diagnostics
	// always stay. An unknown name answers 400 unknown_field with the valid names,
	// before any read. The price stays the same. SDK row types then hold only the
	// named fields.
	Fields param.Opt[string] `query:"fields,omitzero" json:"-"`
	// Maximum locations to return. Omit it for every match. A low credit balance can
	// return fewer. Aliases pageSize, count, max_results, maxItems, max_items &
	// per_page.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Fields to drop from each row, comma-separated. It takes the same names as
	// fields. With fields, it drops from the kept fields.
	Omit param.Opt[string] `query:"omit,omitzero" json:"-"`
	// Text the location name contains, such as york. Case insensitive. Alias query.
	Q param.Opt[string] `query:"q,omitzero" json:"-"`
	// Kind of place. Case insensitive. X labels a few towns unknown.
	//
	// Any of "country", "town", "worldwide", "unknown".
	PlaceType XGetTrendLocationsParamsPlaceType `query:"placeType,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [XGetTrendLocationsParams]'s query parameters as `url.Values`.
func (r XGetTrendLocationsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Kind of place. Case insensitive. X labels a few towns unknown.
type XGetTrendLocationsParamsPlaceType string

const (
	XGetTrendLocationsParamsPlaceTypeCountry   XGetTrendLocationsParamsPlaceType = "country"
	XGetTrendLocationsParamsPlaceTypeTown      XGetTrendLocationsParamsPlaceType = "town"
	XGetTrendLocationsParamsPlaceTypeWorldwide XGetTrendLocationsParamsPlaceType = "worldwide"
	XGetTrendLocationsParamsPlaceTypeUnknown   XGetTrendLocationsParamsPlaceType = "unknown"
)

type XGetTrendsHistoryParams struct {
	// Trends to return in each list, from 1 to 50.
	Count param.Opt[int64] `query:"count,omitzero" json:"-"`
	// Country to read trends history for, by name or 2-letter code, such as Turkey or
	// TR. Case insensitive.
	Country param.Opt[string] `query:"country,omitzero" json:"-"`
	// The next_cursor of the page before, for the next lists. It keeps the time of the
	// first request, so a range of spans stays the same on every page.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Step between 2 returned lists, as a number & m, h or d, such as 15m, 1h or 1d. A
	// step under 5m returns every stored list.
	Every param.Opt[string] `query:"every,omitzero" json:"-"`
	// Fields to keep in each row, comma-separated. Use dots for nested fields, such as
	// author.username. A list keeps the field in each item. Case & underscores do not
	// matter. A row that lacks a field leaves it out. Cursors, counts & diagnostics
	// always stay. An unknown name answers 400 unknown_field with the valid names,
	// before any read. The price stays the same. SDK row types then hold only the
	// named fields.
	Fields param.Opt[string] `query:"fields,omitzero" json:"-"`
	// Start of the range. Takes an ISO 8601 time, a date such as 2026-10-13, or Unix
	// seconds. A time without a zone reads as UTC. It also takes a span back from now,
	// as a number & s, m, h or d, so from=7d reads the last 7 days. Defaults to 24
	// hours before to. History starts 2026-10-13 02:20 UTC, and a range that begins
	// earlier returns the stored part. Alias sinceTime.
	From param.Opt[string] `query:"from,omitzero" json:"-"`
	// Maximum lists 1 request returns & charges. A wider range gives a next_cursor for
	// the next lists. A low credit balance can return fewer.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Town or country to read trends history for, by name, such as Istanbul. Case
	// insensitive. Add country when 2 towns share a name. /x/trends/locations lists
	// every name.
	Location param.Opt[string] `query:"location,omitzero" json:"-"`
	// Fields to drop from each row, comma-separated. It takes the same names as
	// fields. With fields, it drops from the kept fields.
	Omit param.Opt[string] `query:"omit,omitzero" json:"-"`
	// End of the range, which it leaves out. Takes the forms of from, so from=2d with
	// to=1d reads the day before the last 24 hours. Defaults to now. Alias untilTime.
	To param.Opt[string] `query:"to,omitzero" json:"-"`
	// WOEID of the place, such as 1 for Worldwide or 23424969 for Turkey. It wins over
	// country & location. /x/trends/locations lists every WOEID.
	Woeid param.Opt[int64] `query:"woeid,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [XGetTrendsHistoryParams]'s query parameters as `url.Values`.
func (r XGetTrendsHistoryParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type XSearchPlacesParams struct {
	// Place name to search for, such as istanbul. Alias query.
	Q string `query:"q" api:"required" json:"-"`
	// Fields to keep in each row, comma-separated. Use dots for nested fields, such as
	// author.username. A list keeps the field in each item. Case & underscores do not
	// matter. A row that lacks a field leaves it out. Cursors, counts & diagnostics
	// always stay. An unknown name answers 400 unknown_field with the valid names,
	// before any read. The price stays the same. SDK row types then hold only the
	// named fields.
	Fields param.Opt[string] `query:"fields,omitzero" json:"-"`
	// Maximum places to return. Omit it for every place X finds. A low credit balance
	// can return fewer. Aliases pageSize, count, max_results, maxItems, max_items &
	// per_page.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Fields to drop from each row, comma-separated. It takes the same names as
	// fields. With fields, it drops from the kept fields.
	Omit param.Opt[string] `query:"omit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [XSearchPlacesParams]'s query parameters as `url.Values`.
func (r XSearchPlacesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
