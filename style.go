// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package xtwitterscraper

import (
	"context"
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

// List cached style profiles
func (r *StyleService) List(ctx context.Context, opts ...option.RequestOption) (res *StyleListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "styles"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
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

// Style profile summary with tweet count and ownership flag.
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

// Full style profile with sampled tweets used for tone analysis.
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
	// Full style profile with sampled tweets used for tone analysis.
	Style1 StyleCompareResponseStyle1 `json:"style1" api:"required"`
	// Full style profile with sampled tweets used for tone analysis.
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

// Full style profile with sampled tweets used for tone analysis.
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

// Full style profile with sampled tweets used for tone analysis.
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
