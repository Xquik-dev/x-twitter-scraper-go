// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

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

// XUserService contains methods and other services that help with interacting with
// the x-twitter-scraper API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewXUserService] method instead.
type XUserService struct {
	options []option.RequestOption
	// X write actions (tweets, likes, follows, DMs)
	Follow XUserFollowService
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

// Get user profile with follower counts and verification
func (r *XUserService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *shared.UserProfile, err error) {
	var preClientOpts = []option.RequestOption{requestconfig.WithSecurity(requestconfig.Security{
		APIKey:      true,
		OAuthBearer: true,
	})}
	opts = slices.Concat(preClientOpts, r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("x/users/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Remove follower
func (r *XUserService) RemoveFollower(ctx context.Context, id string, body XUserRemoveFollowerParams, opts ...option.RequestOption) (res *XUserRemoveFollowerResponse, err error) {
	var preClientOpts = []option.RequestOption{requestconfig.WithSecurity(requestconfig.Security{
		APIKey:      true,
		OAuthBearer: true,
	})}
	opts = slices.Concat(preClientOpts, r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("x/users/%s/remove-follower", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Look up multiple users by IDs in one call
func (r *XUserService) GetBatch(ctx context.Context, query XUserGetBatchParams, opts ...option.RequestOption) (res *XUserGetBatchResponse, err error) {
	var preClientOpts = []option.RequestOption{requestconfig.WithSecurity(requestconfig.Security{
		APIKey:      true,
		OAuthBearer: true,
	})}
	opts = slices.Concat(preClientOpts, r.options, opts)
	path := "x/users/batch"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// List followers of a user
func (r *XUserService) GetFollowers(ctx context.Context, id string, query XUserGetFollowersParams, opts ...option.RequestOption) (res *shared.PaginatedUsers, err error) {
	var preClientOpts = []option.RequestOption{requestconfig.WithSecurity(requestconfig.Security{
		APIKey:      true,
		OAuthBearer: true,
	})}
	opts = slices.Concat(preClientOpts, r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("x/users/%s/followers", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// List mutual followers between you and a user
func (r *XUserService) GetFollowersYouKnow(ctx context.Context, id string, query XUserGetFollowersYouKnowParams, opts ...option.RequestOption) (res *shared.PaginatedUsers, err error) {
	var preClientOpts = []option.RequestOption{requestconfig.WithSecurity(requestconfig.Security{
		APIKey:      true,
		OAuthBearer: true,
	})}
	opts = slices.Concat(preClientOpts, r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("x/users/%s/followers-you-know", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// List accounts a user follows
func (r *XUserService) GetFollowing(ctx context.Context, id string, query XUserGetFollowingParams, opts ...option.RequestOption) (res *shared.PaginatedUsers, err error) {
	var preClientOpts = []option.RequestOption{requestconfig.WithSecurity(requestconfig.Security{
		APIKey:      true,
		OAuthBearer: true,
	})}
	opts = slices.Concat(preClientOpts, r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("x/users/%s/following", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// List tweets liked by a user
func (r *XUserService) GetLikes(ctx context.Context, id string, query XUserGetLikesParams, opts ...option.RequestOption) (res *shared.PaginatedTweets, err error) {
	var preClientOpts = []option.RequestOption{requestconfig.WithSecurity(requestconfig.Security{
		APIKey:      true,
		OAuthBearer: true,
	})}
	opts = slices.Concat(preClientOpts, r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("x/users/%s/likes", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// List media tweets posted by a user
func (r *XUserService) GetMedia(ctx context.Context, id string, query XUserGetMediaParams, opts ...option.RequestOption) (res *shared.PaginatedTweets, err error) {
	var preClientOpts = []option.RequestOption{requestconfig.WithSecurity(requestconfig.Security{
		APIKey:      true,
		OAuthBearer: true,
	})}
	opts = slices.Concat(preClientOpts, r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("x/users/%s/media", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// List tweets mentioning a user
func (r *XUserService) GetMentions(ctx context.Context, id string, query XUserGetMentionsParams, opts ...option.RequestOption) (res *shared.PaginatedTweets, err error) {
	var preClientOpts = []option.RequestOption{requestconfig.WithSecurity(requestconfig.Security{
		APIKey:      true,
		OAuthBearer: true,
	})}
	opts = slices.Concat(preClientOpts, r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("x/users/%s/mentions", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Returns the user's timeline with replies included by default.
func (r *XUserService) GetReplies(ctx context.Context, id string, query XUserGetRepliesParams, opts ...option.RequestOption) (res *shared.PaginatedTweets, err error) {
	var preClientOpts = []option.RequestOption{requestconfig.WithSecurity(requestconfig.Security{
		APIKey:      true,
		OAuthBearer: true,
	})}
	opts = slices.Concat(preClientOpts, r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("x/users/%s/replies", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Search users by name or username
func (r *XUserService) GetSearch(ctx context.Context, query XUserGetSearchParams, opts ...option.RequestOption) (res *shared.PaginatedUsers, err error) {
	var preClientOpts = []option.RequestOption{requestconfig.WithSecurity(requestconfig.Security{
		APIKey:      true,
		OAuthBearer: true,
	})}
	opts = slices.Concat(preClientOpts, r.options, opts)
	path := "x/users/search"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// List recent tweets posted by a user
func (r *XUserService) GetTweets(ctx context.Context, id string, query XUserGetTweetsParams, opts ...option.RequestOption) (res *shared.PaginatedTweets, err error) {
	var preClientOpts = []option.RequestOption{requestconfig.WithSecurity(requestconfig.Security{
		APIKey:      true,
		OAuthBearer: true,
	})}
	opts = slices.Concat(preClientOpts, r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("x/users/%s/tweets", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// List verified followers of a user
func (r *XUserService) GetVerifiedFollowers(ctx context.Context, id string, query XUserGetVerifiedFollowersParams, opts ...option.RequestOption) (res *shared.PaginatedUsers, err error) {
	var preClientOpts = []option.RequestOption{requestconfig.WithSecurity(requestconfig.Security{
		APIKey:      true,
		OAuthBearer: true,
	})}
	opts = slices.Concat(preClientOpts, r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("x/users/%s/verified-followers", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type XUserRemoveFollowerResponse struct {
	Success bool `json:"success" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XUserRemoveFollowerResponse) RawJSON() string { return r.JSON.raw }
func (r *XUserRemoveFollowerResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Batch user lookup results. Duplicate requested IDs are ignored while preserving
// first-seen order. unavailable_ids identifies processed IDs with no returned
// profile. unprocessed_ids identifies IDs skipped when available credits limit
// processing.
type XUserGetBatchResponse struct {
	// Batch lookups never paginate.
	HasNextPage bool `json:"has_next_page" api:"required"`
	// Empty because batch lookups never paginate.
	NextCursor string `json:"next_cursor" api:"required"`
	// Number of requested IDs included in the lookup.
	ProcessedCount int64 `json:"processed_count" api:"required"`
	// Number of unique IDs requested.
	RequestedCount int64 `json:"requested_count" api:"required"`
	// Number of user profiles returned and charged.
	ReturnedCount int64 `json:"returned_count" api:"required"`
	// Processed IDs with no returned profile, in first-seen request order.
	UnavailableIDs []string `json:"unavailable_ids" api:"required"`
	// Requested IDs skipped because available credits limited processing. Retry these
	// IDs after adding credits.
	UnprocessedIDs []string             `json:"unprocessed_ids" api:"required"`
	Users          []shared.UserProfile `json:"users" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		HasNextPage    respjson.Field
		NextCursor     respjson.Field
		ProcessedCount respjson.Field
		RequestedCount respjson.Field
		ReturnedCount  respjson.Field
		UnavailableIDs respjson.Field
		UnprocessedIDs respjson.Field
		Users          respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XUserGetBatchResponse) RawJSON() string { return r.JSON.raw }
func (r *XUserGetBatchResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type XUserRemoveFollowerParams struct {
	// X account identifier (@username or account ID)
	Account string `json:"account" api:"required"`
	paramObj
}

func (r XUserRemoveFollowerParams) MarshalJSON() (data []byte, err error) {
	type shadow XUserRemoveFollowerParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *XUserRemoveFollowerParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type XUserGetBatchParams struct {
	// Comma-separated numeric user IDs (1-100 values). Duplicate IDs are ignored while
	// preserving first-seen order.
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
	// Legacy cursor alias. Prefer cursor.
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// Pagination cursor for followers list
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Legacy integer page size alias for following lists. Prefer pageSize.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Maximum user profiles requested from this page (20-200, default 200). The
	// response can contain fewer profiles because the source returned fewer or
	// remaining credits cover fewer results. Keep requesting next_cursor while
	// has_next_page is true. The deprecated limit and count aliases remain accepted.
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
	// Maximum user profiles requested from this page (20-200, default 200). The
	// response can contain fewer profiles because the source returned fewer or
	// remaining credits cover fewer results. Keep requesting next_cursor while
	// has_next_page is true. The deprecated limit and count aliases remain accepted.
	PageSize param.Opt[int64] `query:"pageSize,omitzero" json:"-"`
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
	// Legacy cursor alias. Prefer cursor.
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// Pagination cursor for following list
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Legacy page size alias. Prefer pageSize.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Maximum user profiles requested from this page (20-200, default 200). The
	// response can contain fewer profiles because the source returned fewer or
	// remaining credits cover fewer results. Keep requesting next_cursor while
	// has_next_page is true. The deprecated limit and count aliases remain accepted.
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
	// Words or quoted phrases where any one can match. Separate with spaces, commas,
	// or lines.
	AnyWords param.Opt[string] `query:"anyWords,omitzero" json:"-"`
	// Cashtags separated by spaces, commas, or lines.
	Cashtags param.Opt[string] `query:"cashtags,omitzero" json:"-"`
	// Conversation ID filter.
	ConversationID param.Opt[string] `query:"conversationId,omitzero" json:"-"`
	// Pagination cursor for liked tweets
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Exact phrase to match.
	ExactPhrase param.Opt[string] `query:"exactPhrase,omitzero" json:"-"`
	// Words or quoted phrases to exclude. Separate with spaces, commas, or lines.
	ExcludeWords param.Opt[string] `query:"excludeWords,omitzero" json:"-"`
	// Filter by author username.
	FromUser param.Opt[string] `query:"fromUser,omitzero" json:"-"`
	// Hashtags separated by spaces, commas, or lines.
	Hashtags param.Opt[string] `query:"hashtags,omitzero" json:"-"`
	// Only replies to this tweet ID.
	InReplyToTweetID param.Opt[string] `query:"inReplyToTweetId,omitzero" json:"-"`
	// Language code filter, e.g. en or tr.
	Language param.Opt[string] `query:"language,omitzero" json:"-"`
	// Filter tweets mentioning a username.
	Mentioning param.Opt[string] `query:"mentioning,omitzero" json:"-"`
	// Minimum likes threshold.
	MinFaves param.Opt[int64] `query:"minFaves,omitzero" json:"-"`
	// Minimum quote count threshold.
	MinQuotes param.Opt[int64] `query:"minQuotes,omitzero" json:"-"`
	// Minimum replies threshold.
	MinReplies param.Opt[int64] `query:"minReplies,omitzero" json:"-"`
	// Minimum retweets threshold.
	MinRetweets param.Opt[int64] `query:"minRetweets,omitzero" json:"-"`
	// Maximum items requested from this page (1-100, default 20). The response can
	// contain fewer items because the source returned fewer, filters removed items, or
	// remaining credits cover fewer results. Keep requesting next_cursor while
	// has_next_page is true, even when a page is empty. The deprecated limit and count
	// aliases remain accepted.
	PageSize param.Opt[int64] `query:"pageSize,omitzero" json:"-"`
	// Only quotes of this tweet ID.
	QuotesOfTweetID param.Opt[string] `query:"quotesOfTweetId,omitzero" json:"-"`
	// Only retweets of this tweet ID.
	RetweetsOfTweetID param.Opt[string] `query:"retweetsOfTweetId,omitzero" json:"-"`
	// Start date in YYYY-MM-DD format.
	SinceDate param.Opt[time.Time] `query:"sinceDate,omitzero" format:"date" json:"-"`
	// Filter replies sent to a username.
	ToUser param.Opt[string] `query:"toUser,omitzero" json:"-"`
	// End date in YYYY-MM-DD format.
	UntilDate param.Opt[time.Time] `query:"untilDate,omitzero" format:"date" json:"-"`
	// URL substring or domain filter.
	URL param.Opt[string] `query:"url,omitzero" json:"-"`
	// Only return tweets from verified authors.
	VerifiedOnly param.Opt[bool] `query:"verifiedOnly,omitzero" json:"-"`
	// Filter by media type.
	//
	// Any of "images", "videos", "gifs", "media", "links", "none".
	MediaType XUserGetLikesParamsMediaType `query:"mediaType,omitzero" json:"-"`
	// Quote mode.
	//
	// Any of "include", "exclude", "only".
	Quotes XUserGetLikesParamsQuotes `query:"quotes,omitzero" json:"-"`
	// Reply mode.
	//
	// Any of "include", "exclude", "only".
	Replies XUserGetLikesParamsReplies `query:"replies,omitzero" json:"-"`
	// Retweet mode.
	//
	// Any of "include", "exclude", "only".
	Retweets XUserGetLikesParamsRetweets `query:"retweets,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [XUserGetLikesParams]'s query parameters as `url.Values`.
func (r XUserGetLikesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by media type.
type XUserGetLikesParamsMediaType string

const (
	XUserGetLikesParamsMediaTypeImages XUserGetLikesParamsMediaType = "images"
	XUserGetLikesParamsMediaTypeVideos XUserGetLikesParamsMediaType = "videos"
	XUserGetLikesParamsMediaTypeGifs   XUserGetLikesParamsMediaType = "gifs"
	XUserGetLikesParamsMediaTypeMedia  XUserGetLikesParamsMediaType = "media"
	XUserGetLikesParamsMediaTypeLinks  XUserGetLikesParamsMediaType = "links"
	XUserGetLikesParamsMediaTypeNone   XUserGetLikesParamsMediaType = "none"
)

// Quote mode.
type XUserGetLikesParamsQuotes string

const (
	XUserGetLikesParamsQuotesInclude XUserGetLikesParamsQuotes = "include"
	XUserGetLikesParamsQuotesExclude XUserGetLikesParamsQuotes = "exclude"
	XUserGetLikesParamsQuotesOnly    XUserGetLikesParamsQuotes = "only"
)

// Reply mode.
type XUserGetLikesParamsReplies string

const (
	XUserGetLikesParamsRepliesInclude XUserGetLikesParamsReplies = "include"
	XUserGetLikesParamsRepliesExclude XUserGetLikesParamsReplies = "exclude"
	XUserGetLikesParamsRepliesOnly    XUserGetLikesParamsReplies = "only"
)

// Retweet mode.
type XUserGetLikesParamsRetweets string

const (
	XUserGetLikesParamsRetweetsInclude XUserGetLikesParamsRetweets = "include"
	XUserGetLikesParamsRetweetsExclude XUserGetLikesParamsRetweets = "exclude"
	XUserGetLikesParamsRetweetsOnly    XUserGetLikesParamsRetweets = "only"
)

type XUserGetMediaParams struct {
	// Words or quoted phrases where any one can match. Separate with spaces, commas,
	// or lines.
	AnyWords param.Opt[string] `query:"anyWords,omitzero" json:"-"`
	// Cashtags separated by spaces, commas, or lines.
	Cashtags param.Opt[string] `query:"cashtags,omitzero" json:"-"`
	// Conversation ID filter.
	ConversationID param.Opt[string] `query:"conversationId,omitzero" json:"-"`
	// Pagination cursor for media tweets
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Exact phrase to match.
	ExactPhrase param.Opt[string] `query:"exactPhrase,omitzero" json:"-"`
	// Words or quoted phrases to exclude. Separate with spaces, commas, or lines.
	ExcludeWords param.Opt[string] `query:"excludeWords,omitzero" json:"-"`
	// Filter by author username.
	FromUser param.Opt[string] `query:"fromUser,omitzero" json:"-"`
	// Hashtags separated by spaces, commas, or lines.
	Hashtags param.Opt[string] `query:"hashtags,omitzero" json:"-"`
	// Only replies to this tweet ID.
	InReplyToTweetID param.Opt[string] `query:"inReplyToTweetId,omitzero" json:"-"`
	// Language code filter, e.g. en or tr.
	Language param.Opt[string] `query:"language,omitzero" json:"-"`
	// Filter tweets mentioning a username.
	Mentioning param.Opt[string] `query:"mentioning,omitzero" json:"-"`
	// Minimum likes threshold.
	MinFaves param.Opt[int64] `query:"minFaves,omitzero" json:"-"`
	// Minimum quote count threshold.
	MinQuotes param.Opt[int64] `query:"minQuotes,omitzero" json:"-"`
	// Minimum replies threshold.
	MinReplies param.Opt[int64] `query:"minReplies,omitzero" json:"-"`
	// Minimum retweets threshold.
	MinRetweets param.Opt[int64] `query:"minRetweets,omitzero" json:"-"`
	// Maximum items requested from this page (1-100, default 20). The response can
	// contain fewer items because the source returned fewer, filters removed items, or
	// remaining credits cover fewer results. Keep requesting next_cursor while
	// has_next_page is true, even when a page is empty. The deprecated limit and count
	// aliases remain accepted.
	PageSize param.Opt[int64] `query:"pageSize,omitzero" json:"-"`
	// Only quotes of this tweet ID.
	QuotesOfTweetID param.Opt[string] `query:"quotesOfTweetId,omitzero" json:"-"`
	// Only retweets of this tweet ID.
	RetweetsOfTweetID param.Opt[string] `query:"retweetsOfTweetId,omitzero" json:"-"`
	// Start date in YYYY-MM-DD format.
	SinceDate param.Opt[time.Time] `query:"sinceDate,omitzero" format:"date" json:"-"`
	// Filter replies sent to a username.
	ToUser param.Opt[string] `query:"toUser,omitzero" json:"-"`
	// End date in YYYY-MM-DD format.
	UntilDate param.Opt[time.Time] `query:"untilDate,omitzero" format:"date" json:"-"`
	// URL substring or domain filter.
	URL param.Opt[string] `query:"url,omitzero" json:"-"`
	// Only return tweets from verified authors.
	VerifiedOnly param.Opt[bool] `query:"verifiedOnly,omitzero" json:"-"`
	// Filter by media type.
	//
	// Any of "images", "videos", "gifs", "media", "links", "none".
	MediaType XUserGetMediaParamsMediaType `query:"mediaType,omitzero" json:"-"`
	// Quote mode.
	//
	// Any of "include", "exclude", "only".
	Quotes XUserGetMediaParamsQuotes `query:"quotes,omitzero" json:"-"`
	// Reply mode.
	//
	// Any of "include", "exclude", "only".
	Replies XUserGetMediaParamsReplies `query:"replies,omitzero" json:"-"`
	// Retweet mode.
	//
	// Any of "include", "exclude", "only".
	Retweets XUserGetMediaParamsRetweets `query:"retweets,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [XUserGetMediaParams]'s query parameters as `url.Values`.
func (r XUserGetMediaParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by media type.
type XUserGetMediaParamsMediaType string

const (
	XUserGetMediaParamsMediaTypeImages XUserGetMediaParamsMediaType = "images"
	XUserGetMediaParamsMediaTypeVideos XUserGetMediaParamsMediaType = "videos"
	XUserGetMediaParamsMediaTypeGifs   XUserGetMediaParamsMediaType = "gifs"
	XUserGetMediaParamsMediaTypeMedia  XUserGetMediaParamsMediaType = "media"
	XUserGetMediaParamsMediaTypeLinks  XUserGetMediaParamsMediaType = "links"
	XUserGetMediaParamsMediaTypeNone   XUserGetMediaParamsMediaType = "none"
)

// Quote mode.
type XUserGetMediaParamsQuotes string

const (
	XUserGetMediaParamsQuotesInclude XUserGetMediaParamsQuotes = "include"
	XUserGetMediaParamsQuotesExclude XUserGetMediaParamsQuotes = "exclude"
	XUserGetMediaParamsQuotesOnly    XUserGetMediaParamsQuotes = "only"
)

// Reply mode.
type XUserGetMediaParamsReplies string

const (
	XUserGetMediaParamsRepliesInclude XUserGetMediaParamsReplies = "include"
	XUserGetMediaParamsRepliesExclude XUserGetMediaParamsReplies = "exclude"
	XUserGetMediaParamsRepliesOnly    XUserGetMediaParamsReplies = "only"
)

// Retweet mode.
type XUserGetMediaParamsRetweets string

const (
	XUserGetMediaParamsRetweetsInclude XUserGetMediaParamsRetweets = "include"
	XUserGetMediaParamsRetweetsExclude XUserGetMediaParamsRetweets = "exclude"
	XUserGetMediaParamsRetweetsOnly    XUserGetMediaParamsRetweets = "only"
)

type XUserGetMentionsParams struct {
	// Words or quoted phrases where any one can match. Separate with spaces, commas,
	// or lines.
	AnyWords param.Opt[string] `query:"anyWords,omitzero" json:"-"`
	// Cashtags separated by spaces, commas, or lines.
	Cashtags param.Opt[string] `query:"cashtags,omitzero" json:"-"`
	// Conversation ID filter.
	ConversationID param.Opt[string] `query:"conversationId,omitzero" json:"-"`
	// Pagination cursor for mentions
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Exact phrase to match.
	ExactPhrase param.Opt[string] `query:"exactPhrase,omitzero" json:"-"`
	// Words or quoted phrases to exclude. Separate with spaces, commas, or lines.
	ExcludeWords param.Opt[string] `query:"excludeWords,omitzero" json:"-"`
	// Filter by author username.
	FromUser param.Opt[string] `query:"fromUser,omitzero" json:"-"`
	// Hashtags separated by spaces, commas, or lines.
	Hashtags param.Opt[string] `query:"hashtags,omitzero" json:"-"`
	// Only replies to this tweet ID.
	InReplyToTweetID param.Opt[string] `query:"inReplyToTweetId,omitzero" json:"-"`
	// Language code filter, e.g. en or tr.
	Language param.Opt[string] `query:"language,omitzero" json:"-"`
	// Filter tweets mentioning a username.
	Mentioning param.Opt[string] `query:"mentioning,omitzero" json:"-"`
	// Minimum likes threshold.
	MinFaves param.Opt[int64] `query:"minFaves,omitzero" json:"-"`
	// Minimum quote count threshold.
	MinQuotes param.Opt[int64] `query:"minQuotes,omitzero" json:"-"`
	// Minimum replies threshold.
	MinReplies param.Opt[int64] `query:"minReplies,omitzero" json:"-"`
	// Minimum retweets threshold.
	MinRetweets param.Opt[int64] `query:"minRetweets,omitzero" json:"-"`
	// Maximum items requested from this page (1-100, default 20). The response can
	// contain fewer items because the source returned fewer, filters removed items, or
	// remaining credits cover fewer results. Keep requesting next_cursor while
	// has_next_page is true, even when a page is empty. The deprecated limit and count
	// aliases remain accepted.
	PageSize param.Opt[int64] `query:"pageSize,omitzero" json:"-"`
	// Only quotes of this tweet ID.
	QuotesOfTweetID param.Opt[string] `query:"quotesOfTweetId,omitzero" json:"-"`
	// Only retweets of this tweet ID.
	RetweetsOfTweetID param.Opt[string] `query:"retweetsOfTweetId,omitzero" json:"-"`
	// Start date in YYYY-MM-DD format.
	SinceDate param.Opt[time.Time] `query:"sinceDate,omitzero" format:"date" json:"-"`
	// Unix timestamp - return mentions after this time
	SinceTime param.Opt[string] `query:"sinceTime,omitzero" json:"-"`
	// Filter replies sent to a username.
	ToUser param.Opt[string] `query:"toUser,omitzero" json:"-"`
	// End date in YYYY-MM-DD format.
	UntilDate param.Opt[time.Time] `query:"untilDate,omitzero" format:"date" json:"-"`
	// Unix timestamp - return mentions before this time
	UntilTime param.Opt[string] `query:"untilTime,omitzero" json:"-"`
	// URL substring or domain filter.
	URL param.Opt[string] `query:"url,omitzero" json:"-"`
	// Only return tweets from verified authors.
	VerifiedOnly param.Opt[bool] `query:"verifiedOnly,omitzero" json:"-"`
	// Filter by media type.
	//
	// Any of "images", "videos", "gifs", "media", "links", "none".
	MediaType XUserGetMentionsParamsMediaType `query:"mediaType,omitzero" json:"-"`
	// Quote mode.
	//
	// Any of "include", "exclude", "only".
	Quotes XUserGetMentionsParamsQuotes `query:"quotes,omitzero" json:"-"`
	// Reply mode.
	//
	// Any of "include", "exclude", "only".
	Replies XUserGetMentionsParamsReplies `query:"replies,omitzero" json:"-"`
	// Retweet mode.
	//
	// Any of "include", "exclude", "only".
	Retweets XUserGetMentionsParamsRetweets `query:"retweets,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [XUserGetMentionsParams]'s query parameters as `url.Values`.
func (r XUserGetMentionsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by media type.
type XUserGetMentionsParamsMediaType string

const (
	XUserGetMentionsParamsMediaTypeImages XUserGetMentionsParamsMediaType = "images"
	XUserGetMentionsParamsMediaTypeVideos XUserGetMentionsParamsMediaType = "videos"
	XUserGetMentionsParamsMediaTypeGifs   XUserGetMentionsParamsMediaType = "gifs"
	XUserGetMentionsParamsMediaTypeMedia  XUserGetMentionsParamsMediaType = "media"
	XUserGetMentionsParamsMediaTypeLinks  XUserGetMentionsParamsMediaType = "links"
	XUserGetMentionsParamsMediaTypeNone   XUserGetMentionsParamsMediaType = "none"
)

// Quote mode.
type XUserGetMentionsParamsQuotes string

const (
	XUserGetMentionsParamsQuotesInclude XUserGetMentionsParamsQuotes = "include"
	XUserGetMentionsParamsQuotesExclude XUserGetMentionsParamsQuotes = "exclude"
	XUserGetMentionsParamsQuotesOnly    XUserGetMentionsParamsQuotes = "only"
)

// Reply mode.
type XUserGetMentionsParamsReplies string

const (
	XUserGetMentionsParamsRepliesInclude XUserGetMentionsParamsReplies = "include"
	XUserGetMentionsParamsRepliesExclude XUserGetMentionsParamsReplies = "exclude"
	XUserGetMentionsParamsRepliesOnly    XUserGetMentionsParamsReplies = "only"
)

// Retweet mode.
type XUserGetMentionsParamsRetweets string

const (
	XUserGetMentionsParamsRetweetsInclude XUserGetMentionsParamsRetweets = "include"
	XUserGetMentionsParamsRetweetsExclude XUserGetMentionsParamsRetweets = "exclude"
	XUserGetMentionsParamsRetweetsOnly    XUserGetMentionsParamsRetweets = "only"
)

type XUserGetRepliesParams struct {
	// Words or quoted phrases where any one can match. Separate with spaces, commas,
	// or lines.
	AnyWords param.Opt[string] `query:"anyWords,omitzero" json:"-"`
	// Cashtags separated by spaces, commas, or lines.
	Cashtags param.Opt[string] `query:"cashtags,omitzero" json:"-"`
	// Conversation ID filter.
	ConversationID param.Opt[string] `query:"conversationId,omitzero" json:"-"`
	// Pagination cursor for user replies
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Exact phrase to match.
	ExactPhrase param.Opt[string] `query:"exactPhrase,omitzero" json:"-"`
	// Words or quoted phrases to exclude. Separate with spaces, commas, or lines.
	ExcludeWords param.Opt[string] `query:"excludeWords,omitzero" json:"-"`
	// Filter by author username.
	FromUser param.Opt[string] `query:"fromUser,omitzero" json:"-"`
	// Hashtags separated by spaces, commas, or lines.
	Hashtags param.Opt[string] `query:"hashtags,omitzero" json:"-"`
	// Include parent tweet for replies
	IncludeParentTweet param.Opt[bool] `query:"includeParentTweet,omitzero" json:"-"`
	// Only replies to this tweet ID.
	InReplyToTweetID param.Opt[string] `query:"inReplyToTweetId,omitzero" json:"-"`
	// Language code filter, e.g. en or tr.
	Language param.Opt[string] `query:"language,omitzero" json:"-"`
	// Filter tweets mentioning a username.
	Mentioning param.Opt[string] `query:"mentioning,omitzero" json:"-"`
	// Minimum likes threshold.
	MinFaves param.Opt[int64] `query:"minFaves,omitzero" json:"-"`
	// Minimum quote count threshold.
	MinQuotes param.Opt[int64] `query:"minQuotes,omitzero" json:"-"`
	// Minimum replies threshold.
	MinReplies param.Opt[int64] `query:"minReplies,omitzero" json:"-"`
	// Minimum retweets threshold.
	MinRetweets param.Opt[int64] `query:"minRetweets,omitzero" json:"-"`
	// Maximum items requested from this page (1-100, default 20). The response can
	// contain fewer items because the source returned fewer, filters removed items, or
	// remaining credits cover fewer results. Keep requesting next_cursor while
	// has_next_page is true, even when a page is empty. The deprecated limit and count
	// aliases remain accepted.
	PageSize param.Opt[int64] `query:"pageSize,omitzero" json:"-"`
	// Only quotes of this tweet ID.
	QuotesOfTweetID param.Opt[string] `query:"quotesOfTweetId,omitzero" json:"-"`
	// Only retweets of this tweet ID.
	RetweetsOfTweetID param.Opt[string] `query:"retweetsOfTweetId,omitzero" json:"-"`
	// Start date in YYYY-MM-DD format.
	SinceDate param.Opt[time.Time] `query:"sinceDate,omitzero" format:"date" json:"-"`
	// Filter replies sent to a username.
	ToUser param.Opt[string] `query:"toUser,omitzero" json:"-"`
	// End date in YYYY-MM-DD format.
	UntilDate param.Opt[time.Time] `query:"untilDate,omitzero" format:"date" json:"-"`
	// URL substring or domain filter.
	URL param.Opt[string] `query:"url,omitzero" json:"-"`
	// Only return tweets from verified authors.
	VerifiedOnly param.Opt[bool] `query:"verifiedOnly,omitzero" json:"-"`
	// Filter by media type.
	//
	// Any of "images", "videos", "gifs", "media", "links", "none".
	MediaType XUserGetRepliesParamsMediaType `query:"mediaType,omitzero" json:"-"`
	// Quote mode.
	//
	// Any of "include", "exclude", "only".
	Quotes XUserGetRepliesParamsQuotes `query:"quotes,omitzero" json:"-"`
	// Reply mode.
	//
	// Any of "include", "exclude", "only".
	Replies XUserGetRepliesParamsReplies `query:"replies,omitzero" json:"-"`
	// Retweet mode.
	//
	// Any of "include", "exclude", "only".
	Retweets XUserGetRepliesParamsRetweets `query:"retweets,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [XUserGetRepliesParams]'s query parameters as `url.Values`.
func (r XUserGetRepliesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by media type.
type XUserGetRepliesParamsMediaType string

const (
	XUserGetRepliesParamsMediaTypeImages XUserGetRepliesParamsMediaType = "images"
	XUserGetRepliesParamsMediaTypeVideos XUserGetRepliesParamsMediaType = "videos"
	XUserGetRepliesParamsMediaTypeGifs   XUserGetRepliesParamsMediaType = "gifs"
	XUserGetRepliesParamsMediaTypeMedia  XUserGetRepliesParamsMediaType = "media"
	XUserGetRepliesParamsMediaTypeLinks  XUserGetRepliesParamsMediaType = "links"
	XUserGetRepliesParamsMediaTypeNone   XUserGetRepliesParamsMediaType = "none"
)

// Quote mode.
type XUserGetRepliesParamsQuotes string

const (
	XUserGetRepliesParamsQuotesInclude XUserGetRepliesParamsQuotes = "include"
	XUserGetRepliesParamsQuotesExclude XUserGetRepliesParamsQuotes = "exclude"
	XUserGetRepliesParamsQuotesOnly    XUserGetRepliesParamsQuotes = "only"
)

// Reply mode.
type XUserGetRepliesParamsReplies string

const (
	XUserGetRepliesParamsRepliesInclude XUserGetRepliesParamsReplies = "include"
	XUserGetRepliesParamsRepliesExclude XUserGetRepliesParamsReplies = "exclude"
	XUserGetRepliesParamsRepliesOnly    XUserGetRepliesParamsReplies = "only"
)

// Retweet mode.
type XUserGetRepliesParamsRetweets string

const (
	XUserGetRepliesParamsRetweetsInclude XUserGetRepliesParamsRetweets = "include"
	XUserGetRepliesParamsRetweetsExclude XUserGetRepliesParamsRetweets = "exclude"
	XUserGetRepliesParamsRetweetsOnly    XUserGetRepliesParamsRetweets = "only"
)

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
	// Words or quoted phrases where any one can match. Separate with spaces, commas,
	// or lines.
	AnyWords param.Opt[string] `query:"anyWords,omitzero" json:"-"`
	// Cashtags separated by spaces, commas, or lines.
	Cashtags param.Opt[string] `query:"cashtags,omitzero" json:"-"`
	// Conversation ID filter.
	ConversationID param.Opt[string] `query:"conversationId,omitzero" json:"-"`
	// Pagination cursor for user tweets
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Exact phrase to match.
	ExactPhrase param.Opt[string] `query:"exactPhrase,omitzero" json:"-"`
	// Words or quoted phrases to exclude. Separate with spaces, commas, or lines.
	ExcludeWords param.Opt[string] `query:"excludeWords,omitzero" json:"-"`
	// Filter by author username.
	FromUser param.Opt[string] `query:"fromUser,omitzero" json:"-"`
	// Hashtags separated by spaces, commas, or lines.
	Hashtags param.Opt[string] `query:"hashtags,omitzero" json:"-"`
	// Include parent tweet for replies
	IncludeParentTweet param.Opt[bool] `query:"includeParentTweet,omitzero" json:"-"`
	// Include reply tweets
	IncludeReplies param.Opt[bool] `query:"includeReplies,omitzero" json:"-"`
	// Only replies to this tweet ID.
	InReplyToTweetID param.Opt[string] `query:"inReplyToTweetId,omitzero" json:"-"`
	// Language code filter, e.g. en or tr.
	Language param.Opt[string] `query:"language,omitzero" json:"-"`
	// Filter tweets mentioning a username.
	Mentioning param.Opt[string] `query:"mentioning,omitzero" json:"-"`
	// Minimum likes threshold.
	MinFaves param.Opt[int64] `query:"minFaves,omitzero" json:"-"`
	// Minimum quote count threshold.
	MinQuotes param.Opt[int64] `query:"minQuotes,omitzero" json:"-"`
	// Minimum replies threshold.
	MinReplies param.Opt[int64] `query:"minReplies,omitzero" json:"-"`
	// Minimum retweets threshold.
	MinRetweets param.Opt[int64] `query:"minRetweets,omitzero" json:"-"`
	// Maximum items requested from this page (1-100, default 20). The response can
	// contain fewer items because the source returned fewer, filters removed items, or
	// remaining credits cover fewer results. Keep requesting next_cursor while
	// has_next_page is true, even when a page is empty. The deprecated limit and count
	// aliases remain accepted.
	PageSize param.Opt[int64] `query:"pageSize,omitzero" json:"-"`
	// Only quotes of this tweet ID.
	QuotesOfTweetID param.Opt[string] `query:"quotesOfTweetId,omitzero" json:"-"`
	// Only retweets of this tweet ID.
	RetweetsOfTweetID param.Opt[string] `query:"retweetsOfTweetId,omitzero" json:"-"`
	// Start date in YYYY-MM-DD format.
	SinceDate param.Opt[time.Time] `query:"sinceDate,omitzero" format:"date" json:"-"`
	// Filter replies sent to a username.
	ToUser param.Opt[string] `query:"toUser,omitzero" json:"-"`
	// End date in YYYY-MM-DD format.
	UntilDate param.Opt[time.Time] `query:"untilDate,omitzero" format:"date" json:"-"`
	// URL substring or domain filter.
	URL param.Opt[string] `query:"url,omitzero" json:"-"`
	// Only return tweets from verified authors.
	VerifiedOnly param.Opt[bool] `query:"verifiedOnly,omitzero" json:"-"`
	// Filter by media type.
	//
	// Any of "images", "videos", "gifs", "media", "links", "none".
	MediaType XUserGetTweetsParamsMediaType `query:"mediaType,omitzero" json:"-"`
	// Quote mode.
	//
	// Any of "include", "exclude", "only".
	Quotes XUserGetTweetsParamsQuotes `query:"quotes,omitzero" json:"-"`
	// Reply mode.
	//
	// Any of "include", "exclude", "only".
	Replies XUserGetTweetsParamsReplies `query:"replies,omitzero" json:"-"`
	// Retweet mode.
	//
	// Any of "include", "exclude", "only".
	Retweets XUserGetTweetsParamsRetweets `query:"retweets,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [XUserGetTweetsParams]'s query parameters as `url.Values`.
func (r XUserGetTweetsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by media type.
type XUserGetTweetsParamsMediaType string

const (
	XUserGetTweetsParamsMediaTypeImages XUserGetTweetsParamsMediaType = "images"
	XUserGetTweetsParamsMediaTypeVideos XUserGetTweetsParamsMediaType = "videos"
	XUserGetTweetsParamsMediaTypeGifs   XUserGetTweetsParamsMediaType = "gifs"
	XUserGetTweetsParamsMediaTypeMedia  XUserGetTweetsParamsMediaType = "media"
	XUserGetTweetsParamsMediaTypeLinks  XUserGetTweetsParamsMediaType = "links"
	XUserGetTweetsParamsMediaTypeNone   XUserGetTweetsParamsMediaType = "none"
)

// Quote mode.
type XUserGetTweetsParamsQuotes string

const (
	XUserGetTweetsParamsQuotesInclude XUserGetTweetsParamsQuotes = "include"
	XUserGetTweetsParamsQuotesExclude XUserGetTweetsParamsQuotes = "exclude"
	XUserGetTweetsParamsQuotesOnly    XUserGetTweetsParamsQuotes = "only"
)

// Reply mode.
type XUserGetTweetsParamsReplies string

const (
	XUserGetTweetsParamsRepliesInclude XUserGetTweetsParamsReplies = "include"
	XUserGetTweetsParamsRepliesExclude XUserGetTweetsParamsReplies = "exclude"
	XUserGetTweetsParamsRepliesOnly    XUserGetTweetsParamsReplies = "only"
)

// Retweet mode.
type XUserGetTweetsParamsRetweets string

const (
	XUserGetTweetsParamsRetweetsInclude XUserGetTweetsParamsRetweets = "include"
	XUserGetTweetsParamsRetweetsExclude XUserGetTweetsParamsRetweets = "exclude"
	XUserGetTweetsParamsRetweetsOnly    XUserGetTweetsParamsRetweets = "only"
)

type XUserGetVerifiedFollowersParams struct {
	// Pagination cursor for verified followers
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Maximum user profiles requested from this page (20-200, default 200). The
	// response can contain fewer profiles because the source returned fewer or
	// remaining credits cover fewer results. Keep requesting next_cursor while
	// has_next_page is true. The deprecated limit and count aliases remain accepted.
	PageSize param.Opt[int64] `query:"pageSize,omitzero" json:"-"`
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
