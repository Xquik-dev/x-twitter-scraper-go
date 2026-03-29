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

	"github.com/stainless-sdks/x-twitter-scraper-go/internal/apijson"
	"github.com/stainless-sdks/x-twitter-scraper-go/internal/apiquery"
	"github.com/stainless-sdks/x-twitter-scraper-go/internal/requestconfig"
	"github.com/stainless-sdks/x-twitter-scraper-go/option"
	"github.com/stainless-sdks/x-twitter-scraper-go/packages/param"
	"github.com/stainless-sdks/x-twitter-scraper-go/packages/respjson"
)

// Tweet composition, drafts, writing styles & radar
//
// StyleService contains methods and other services that help with interacting with
// the x-twitter-scraper API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewStyleService] method instead.
type StyleService struct {
	options []option.RequestOption
}

// NewStyleService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewStyleService(opts ...option.RequestOption) (r StyleService) {
	r = StyleService{}
	r.options = opts
	return
}

// Get cached style profile
func (r *StyleService) Get(ctx context.Context, username string, opts ...option.RequestOption) (res *StyleGetResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if username == "" {
		err = errors.New("missing required username parameter")
		return nil, err
	}
	path := fmt.Sprintf("styles/%s", url.PathEscape(username))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Save style profile with custom tweets
func (r *StyleService) Update(ctx context.Context, username string, body StyleUpdateParams, opts ...option.RequestOption) (res *StyleUpdateResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if username == "" {
		err = errors.New("missing required username parameter")
		return nil, err
	}
	path := fmt.Sprintf("styles/%s", url.PathEscape(username))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

// List cached style profiles
func (r *StyleService) List(ctx context.Context, opts ...option.RequestOption) (res *StyleListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "styles"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Delete a style profile
func (r *StyleService) Delete(ctx context.Context, username string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if username == "" {
		err = errors.New("missing required username parameter")
		return err
	}
	path := fmt.Sprintf("styles/%s", url.PathEscape(username))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Analyze writing style from recent tweets
func (r *StyleService) Analyze(ctx context.Context, body StyleAnalyzeParams, opts ...option.RequestOption) (res *StyleAnalyzeResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "styles"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Compare two style profiles
func (r *StyleService) Compare(ctx context.Context, query StyleCompareParams, opts ...option.RequestOption) (res *StyleCompareResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "styles/compare"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Get engagement metrics for style tweets
func (r *StyleService) GetPerformance(ctx context.Context, username string, opts ...option.RequestOption) (res *StyleGetPerformanceResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if username == "" {
		err = errors.New("missing required username parameter")
		return nil, err
	}
	path := fmt.Sprintf("styles/%s/performance", url.PathEscape(username))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type StyleGetResponse struct {
	FetchedAt    time.Time               `json:"fetchedAt" api:"required" format:"date-time"`
	IsOwnAccount bool                    `json:"isOwnAccount" api:"required"`
	TweetCount   int64                   `json:"tweetCount" api:"required"`
	Tweets       []StyleGetResponseTweet `json:"tweets" api:"required"`
	XUsername    string                  `json:"xUsername" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FetchedAt    respjson.Field
		IsOwnAccount respjson.Field
		TweetCount   respjson.Field
		Tweets       respjson.Field
		XUsername    respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StyleGetResponse) RawJSON() string { return r.JSON.raw }
func (r *StyleGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StyleGetResponseTweet struct {
	ID             string `json:"id" api:"required"`
	Text           string `json:"text" api:"required"`
	AuthorUsername string `json:"authorUsername"`
	CreatedAt      string `json:"createdAt"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		Text           respjson.Field
		AuthorUsername respjson.Field
		CreatedAt      respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StyleGetResponseTweet) RawJSON() string { return r.JSON.raw }
func (r *StyleGetResponseTweet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StyleUpdateResponse struct {
	FetchedAt    time.Time                  `json:"fetchedAt" api:"required" format:"date-time"`
	IsOwnAccount bool                       `json:"isOwnAccount" api:"required"`
	TweetCount   int64                      `json:"tweetCount" api:"required"`
	Tweets       []StyleUpdateResponseTweet `json:"tweets" api:"required"`
	XUsername    string                     `json:"xUsername" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FetchedAt    respjson.Field
		IsOwnAccount respjson.Field
		TweetCount   respjson.Field
		Tweets       respjson.Field
		XUsername    respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StyleUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *StyleUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StyleUpdateResponseTweet struct {
	ID             string `json:"id" api:"required"`
	Text           string `json:"text" api:"required"`
	AuthorUsername string `json:"authorUsername"`
	CreatedAt      string `json:"createdAt"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		Text           respjson.Field
		AuthorUsername respjson.Field
		CreatedAt      respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StyleUpdateResponseTweet) RawJSON() string { return r.JSON.raw }
func (r *StyleUpdateResponseTweet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StyleListResponse struct {
	Styles []StyleListResponseStyle `json:"styles" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Styles      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StyleListResponse) RawJSON() string { return r.JSON.raw }
func (r *StyleListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StyleListResponseStyle struct {
	FetchedAt    time.Time `json:"fetchedAt" api:"required" format:"date-time"`
	IsOwnAccount bool      `json:"isOwnAccount" api:"required"`
	TweetCount   int64     `json:"tweetCount" api:"required"`
	XUsername    string    `json:"xUsername" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FetchedAt    respjson.Field
		IsOwnAccount respjson.Field
		TweetCount   respjson.Field
		XUsername    respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StyleListResponseStyle) RawJSON() string { return r.JSON.raw }
func (r *StyleListResponseStyle) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StyleAnalyzeResponse struct {
	FetchedAt    time.Time                   `json:"fetchedAt" api:"required" format:"date-time"`
	IsOwnAccount bool                        `json:"isOwnAccount" api:"required"`
	TweetCount   int64                       `json:"tweetCount" api:"required"`
	Tweets       []StyleAnalyzeResponseTweet `json:"tweets" api:"required"`
	XUsername    string                      `json:"xUsername" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FetchedAt    respjson.Field
		IsOwnAccount respjson.Field
		TweetCount   respjson.Field
		Tweets       respjson.Field
		XUsername    respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StyleAnalyzeResponse) RawJSON() string { return r.JSON.raw }
func (r *StyleAnalyzeResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StyleAnalyzeResponseTweet struct {
	ID             string `json:"id" api:"required"`
	Text           string `json:"text" api:"required"`
	AuthorUsername string `json:"authorUsername"`
	CreatedAt      string `json:"createdAt"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		Text           respjson.Field
		AuthorUsername respjson.Field
		CreatedAt      respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StyleAnalyzeResponseTweet) RawJSON() string { return r.JSON.raw }
func (r *StyleAnalyzeResponseTweet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StyleCompareResponse struct {
	Style1 StyleCompareResponseStyle1 `json:"style1" api:"required"`
	Style2 StyleCompareResponseStyle2 `json:"style2" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Style1      respjson.Field
		Style2      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StyleCompareResponse) RawJSON() string { return r.JSON.raw }
func (r *StyleCompareResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StyleCompareResponseStyle1 struct {
	FetchedAt    time.Time                         `json:"fetchedAt" api:"required" format:"date-time"`
	IsOwnAccount bool                              `json:"isOwnAccount" api:"required"`
	TweetCount   int64                             `json:"tweetCount" api:"required"`
	Tweets       []StyleCompareResponseStyle1Tweet `json:"tweets" api:"required"`
	XUsername    string                            `json:"xUsername" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FetchedAt    respjson.Field
		IsOwnAccount respjson.Field
		TweetCount   respjson.Field
		Tweets       respjson.Field
		XUsername    respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StyleCompareResponseStyle1) RawJSON() string { return r.JSON.raw }
func (r *StyleCompareResponseStyle1) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StyleCompareResponseStyle1Tweet struct {
	ID             string `json:"id" api:"required"`
	Text           string `json:"text" api:"required"`
	AuthorUsername string `json:"authorUsername"`
	CreatedAt      string `json:"createdAt"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		Text           respjson.Field
		AuthorUsername respjson.Field
		CreatedAt      respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StyleCompareResponseStyle1Tweet) RawJSON() string { return r.JSON.raw }
func (r *StyleCompareResponseStyle1Tweet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StyleCompareResponseStyle2 struct {
	FetchedAt    time.Time                         `json:"fetchedAt" api:"required" format:"date-time"`
	IsOwnAccount bool                              `json:"isOwnAccount" api:"required"`
	TweetCount   int64                             `json:"tweetCount" api:"required"`
	Tweets       []StyleCompareResponseStyle2Tweet `json:"tweets" api:"required"`
	XUsername    string                            `json:"xUsername" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FetchedAt    respjson.Field
		IsOwnAccount respjson.Field
		TweetCount   respjson.Field
		Tweets       respjson.Field
		XUsername    respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StyleCompareResponseStyle2) RawJSON() string { return r.JSON.raw }
func (r *StyleCompareResponseStyle2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StyleCompareResponseStyle2Tweet struct {
	ID             string `json:"id" api:"required"`
	Text           string `json:"text" api:"required"`
	AuthorUsername string `json:"authorUsername"`
	CreatedAt      string `json:"createdAt"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		Text           respjson.Field
		AuthorUsername respjson.Field
		CreatedAt      respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StyleCompareResponseStyle2Tweet) RawJSON() string { return r.JSON.raw }
func (r *StyleCompareResponseStyle2Tweet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StyleGetPerformanceResponse struct {
	TweetCount int64                              `json:"tweetCount" api:"required"`
	Tweets     []StyleGetPerformanceResponseTweet `json:"tweets" api:"required"`
	XUsername  string                             `json:"xUsername" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		TweetCount  respjson.Field
		Tweets      respjson.Field
		XUsername   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StyleGetPerformanceResponse) RawJSON() string { return r.JSON.raw }
func (r *StyleGetPerformanceResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StyleGetPerformanceResponseTweet struct {
	ID           string `json:"id" api:"required"`
	Text         string `json:"text" api:"required"`
	CreatedAt    string `json:"createdAt"`
	LikeCount    int64  `json:"likeCount"`
	ReplyCount   int64  `json:"replyCount"`
	RetweetCount int64  `json:"retweetCount"`
	ViewCount    int64  `json:"viewCount"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Text         respjson.Field
		CreatedAt    respjson.Field
		LikeCount    respjson.Field
		ReplyCount   respjson.Field
		RetweetCount respjson.Field
		ViewCount    respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StyleGetPerformanceResponseTweet) RawJSON() string { return r.JSON.raw }
func (r *StyleGetPerformanceResponseTweet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StyleUpdateParams struct {
	// Display label for the style
	Label string `json:"label" api:"required"`
	// Array of tweet objects
	Tweets []StyleUpdateParamsTweet `json:"tweets,omitzero" api:"required"`
	paramObj
}

func (r StyleUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow StyleUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *StyleUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Text is required.
type StyleUpdateParamsTweet struct {
	Text string `json:"text" api:"required"`
	paramObj
}

func (r StyleUpdateParamsTweet) MarshalJSON() (data []byte, err error) {
	type shadow StyleUpdateParamsTweet
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *StyleUpdateParamsTweet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StyleAnalyzeParams struct {
	// X username to analyze
	Username string `json:"username" api:"required"`
	paramObj
}

func (r StyleAnalyzeParams) MarshalJSON() (data []byte, err error) {
	type shadow StyleAnalyzeParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *StyleAnalyzeParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StyleCompareParams struct {
	// First username to compare
	Username1 string `query:"username1" api:"required" json:"-"`
	// Second username to compare
	Username2 string `query:"username2" api:"required" json:"-"`
	paramObj
}

// URLQuery serializes [StyleCompareParams]'s query parameters as `url.Values`.
func (r StyleCompareParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
