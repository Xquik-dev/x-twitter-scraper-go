// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

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
	"github.com/Xquik-dev/x-twitter-scraper-go/shared"
)

// Look up, search, and analyze individual tweets
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

// Returns bookmarks from the connected X account.
func (r *XBookmarkService) List(ctx context.Context, query XBookmarkListParams, opts ...option.RequestOption) (res *shared.PaginatedTweets, err error) {
	opts = slices.Concat(r.options, opts)
	path := "x/bookmarks"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Returns bookmark folders from the connected X account.
func (r *XBookmarkService) GetFolders(ctx context.Context, opts ...option.RequestOption) (res *XBookmarkGetFoldersResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "x/bookmarks/folders"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type XBookmarkGetFoldersResponse struct {
	Folders []XBookmarkGetFoldersResponseFolder `json:"folders" api:"required"`
	// Whether another folder page is available
	HasNextPage bool `json:"has_next_page" api:"required"`
	// Cursor for the next folder page
	NextCursor string `json:"next_cursor" api:"required"`
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

// Bookmark folder and its optional public cover image.
type XBookmarkGetFoldersResponseFolder struct {
	// Folder ID.
	ID string `json:"id" api:"required"`
	// Public folder cover image metadata.
	Media XBookmarkGetFoldersResponseFolderMedia `json:"media"`
	// Folder name.
	Name string `json:"name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Media       respjson.Field
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

// Public folder cover image metadata.
type XBookmarkGetFoldersResponseFolderMedia struct {
	// Media object ID.
	ID string `json:"id"`
	// Media ID.
	MediaID string `json:"mediaId"`
	// Stable media key.
	MediaKey string `json:"mediaKey"`
	// Original image height.
	OriginalImageHeight int64 `json:"originalImageHeight"`
	// Original image URL.
	OriginalImageURL string `json:"originalImageUrl"`
	// Original image width.
	OriginalImageWidth int64 `json:"originalImageWidth"`
	// Dominant image colors and their proportions.
	Palette []XBookmarkGetFoldersResponseFolderMediaPalette `json:"palette"`
	// Media object type.
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                  respjson.Field
		MediaID             respjson.Field
		MediaKey            respjson.Field
		OriginalImageHeight respjson.Field
		OriginalImageURL    respjson.Field
		OriginalImageWidth  respjson.Field
		Palette             respjson.Field
		Type                respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XBookmarkGetFoldersResponseFolderMedia) RawJSON() string { return r.JSON.raw }
func (r *XBookmarkGetFoldersResponseFolderMedia) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type XBookmarkGetFoldersResponseFolderMediaPalette struct {
	Percentage float64                                          `json:"percentage"`
	Rgb        XBookmarkGetFoldersResponseFolderMediaPaletteRgb `json:"rgb"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Percentage  respjson.Field
		Rgb         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XBookmarkGetFoldersResponseFolderMediaPalette) RawJSON() string { return r.JSON.raw }
func (r *XBookmarkGetFoldersResponseFolderMediaPalette) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type XBookmarkGetFoldersResponseFolderMediaPaletteRgb struct {
	Blue  float64 `json:"blue"`
	Green float64 `json:"green"`
	Red   float64 `json:"red"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Blue        respjson.Field
		Green       respjson.Field
		Red         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XBookmarkGetFoldersResponseFolderMediaPaletteRgb) RawJSON() string { return r.JSON.raw }
func (r *XBookmarkGetFoldersResponseFolderMediaPaletteRgb) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type XBookmarkListParams struct {
	// Pagination cursor for bookmarks
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
