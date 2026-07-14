// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package xtwitterscraper

import (
	"context"
	"net/http"
	"slices"

	"github.com/Xquik-dev/x-twitter-scraper-go/internal/apijson"
	"github.com/Xquik-dev/x-twitter-scraper-go/internal/requestconfig"
	"github.com/Xquik-dev/x-twitter-scraper-go/option"
	"github.com/Xquik-dev/x-twitter-scraper-go/packages/param"
	"github.com/Xquik-dev/x-twitter-scraper-go/packages/respjson"
)

// Media upload and download
//
// XMediaService contains methods and other services that help with interacting
// with the x-twitter-scraper API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewXMediaService] method instead.
type XMediaService struct {
	options []option.RequestOption
}

// NewXMediaService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewXMediaService(opts ...option.RequestOption) (r XMediaService) {
	r = XMediaService{}
	r.options = opts
	return
}

// Download images and videos from tweets
func (r *XMediaService) Download(ctx context.Context, body XMediaDownloadParams, opts ...option.RequestOption) (res *XMediaDownloadResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "x/media/download"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Upload media
func (r *XMediaService) Upload(ctx context.Context, body XMediaUploadParams, opts ...option.RequestOption) (res *XMediaUploadResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "x/media"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type XMediaDownloadResponse struct {
	CacheHit    bool   `json:"cacheHit"`
	GalleryURL  string `json:"galleryUrl"`
	TotalMedia  int64  `json:"totalMedia"`
	TotalTweets int64  `json:"totalTweets"`
	TweetID     string `json:"tweetId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CacheHit    respjson.Field
		GalleryURL  respjson.Field
		TotalMedia  respjson.Field
		TotalTweets respjson.Field
		TweetID     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XMediaDownloadResponse) RawJSON() string { return r.JSON.raw }
func (r *XMediaDownloadResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type XMediaUploadResponse struct {
	MediaID string `json:"mediaId" api:"required"`
	// Public media URL for tweet `media` arrays.
	MediaURL string `json:"mediaUrl" api:"required" format:"uri"`
	Success  bool   `json:"success" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MediaID     respjson.Field
		MediaURL    respjson.Field
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XMediaUploadResponse) RawJSON() string { return r.JSON.raw }
func (r *XMediaUploadResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type XMediaDownloadParams struct {
	// Numeric tweet ID alias for tweetInput
	TweetID param.Opt[string] `json:"tweetId,omitzero"`
	// Tweet URL or ID (single tweet)
	TweetInput param.Opt[string] `json:"tweetInput,omitzero"`
	// Tweet URL alias for tweetInput
	TweetURL param.Opt[string] `json:"tweetUrl,omitzero"`
	// Array of tweet URLs or IDs (bulk, max 50 string items)
	TweetIDs []string `json:"tweetIds,omitzero"`
	paramObj
}

func (r XMediaDownloadParams) MarshalJSON() (data []byte, err error) {
	type shadow XMediaDownloadParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *XMediaDownloadParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type XMediaUploadParams struct {
	// X account (@username or ID) uploading media from URL
	Account string `json:"account" api:"required"`
	// HTTPS URL to download and upload as media
	URL string `json:"url" api:"required" format:"uri"`
	paramObj
}

func (r XMediaUploadParams) MarshalJSON() (data []byte, err error) {
	type shadow XMediaUploadParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *XMediaUploadParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
