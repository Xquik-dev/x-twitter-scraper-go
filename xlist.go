// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

package xtwitterscraper

import (
	"context"
	"encoding/json"
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

// X List followers, members, and tweets
//
// XListService contains methods and other services that help with interacting with
// the x-twitter-scraper API.
//
// You should not instantiate this service directly, and instead use the
// [NewXListService] method instead.
type XListService struct {
	options []option.RequestOption
}

// NewXListService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewXListService(opts ...option.RequestOption) (r XListService) {
	r = XListService{}
	r.options = opts
	return
}

// Returns 1 public X List as its page on X shows it. Accepts the List ID or the
// URL-encoded List URL, such as x.com/i/lists/4143408. A private List answers 404
// list_not_found.
func (r *XListService) Get(ctx context.Context, id string, query XListGetParams, opts ...option.RequestOption) (res *XListGetResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("x/lists/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Returns List followers with resumable or standard pagination.
func (r *XListService) GetFollowers(ctx context.Context, id string, query XListGetFollowersParams, opts ...option.RequestOption) (res *XListGetFollowersResponseUnion, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("x/lists/%s/followers", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Returns List members with resumable or standard pagination.
func (r *XListService) GetMembers(ctx context.Context, id string, query XListGetMembersParams, opts ...option.RequestOption) (res *XListGetMembersResponseUnion, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("x/lists/%s/members", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Returns 1 page of public Lists with member counts, banners & owner profiles. X
// may rank the Lists differently per request. Pass next_cursor as cursor for the
// next page.
func (r *XListService) GetSearch(ctx context.Context, query XListGetSearchParams, opts ...option.RequestOption) (res *XListGetSearchResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "x/lists/search"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Omit mode for resumable maximum coverage. Pass next_cursor unchanged. Standard
// keeps legacy pagination.
func (r *XListService) GetTweets(ctx context.Context, id string, query XListGetTweetsParams, opts ...option.RequestOption) (res *shared.PaginatedTweets, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("x/lists/%s/tweets", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// XListGetFollowersResponseUnion contains all possible properties and values from
// [shared.PaginatedUsers], [XListGetFollowersResponseUserListCoverageResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type XListGetFollowersResponseUnion struct {
	HasNextPage bool   `json:"has_next_page"`
	NextCursor  string `json:"next_cursor"`
	// Present on the last page of a filtered automatic read that sent no maxScan, once
	// the list scanned 100,000 profiles across its pages. That page is the list's
	// last.
	ScanLimit XListGetFollowersResponseScanLimit `json:"scanLimit"`
	// This field is from variant [shared.PaginatedUsers],
	// [XListGetFollowersResponseUserListCoverageResponse].
	Users []shared.UserProfile `json:"users"`
	// This field is from variant [shared.PaginatedUsers],
	// [XListGetFollowersResponseUserListCoverageResponse].
	FilteredCount int64 `json:"filtered_count"`
	// This field is from variant [XListGetFollowersResponseUserListCoverageResponse].
	Diagnostic XListGetFollowersResponseUserListCoverageResponseDiagnostic `json:"diagnostic"`
	JSON       struct {
		HasNextPage   respjson.Field
		NextCursor    respjson.Field
		ScanLimit     respjson.Field
		Users         respjson.Field
		FilteredCount respjson.Field
		Diagnostic    respjson.Field
		raw           string
	} `json:"-"`
}

func (u XListGetFollowersResponseUnion) AsPaginatedUsers() (v shared.PaginatedUsers) {
	_ = v.UnmarshalJSON(json.RawMessage(u.JSON.raw))
	return
}

func (u XListGetFollowersResponseUnion) AsXListGetFollowersResponseUserListCoverageResponse() (v XListGetFollowersResponseUserListCoverageResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v, &v.JSON.raw)
	return
}

// Returns the unmodified JSON received from the API
func (u XListGetFollowersResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *XListGetFollowersResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// Present on the last page of a filtered automatic read that sent no maxScan, once
// the list scanned 100,000 profiles across its pages. That page is the list's
// last.
type XListGetFollowersResponseScanLimit struct {
	// The profiles a filtered read scans of a list without maxScan.
	//
	// Any of 100000.
	Limit int64 `json:"limit" api:"required"`
	// What stopped the list & how to read further.
	Message string `json:"message" api:"required"`
	// Profiles the list read across its pages, kept or dropped by the filters. The
	// last page may read past the limit.
	RowsRead int64 `json:"rowsRead" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Limit       respjson.Field
		Message     respjson.Field
		RowsRead    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XListGetFollowersResponseScanLimit) RawJSON() string { return r.JSON.raw }
func (r *XListGetFollowersResponseScanLimit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// Terminal relationship coverage response with diagnostics.
type XListGetFollowersResponseUserListCoverageResponse struct {
	// Coverage evidence across parallel relationship strategies.
	Diagnostic XListGetFollowersResponseUserListCoverageResponseDiagnostic `json:"diagnostic" api:"required"`
	// Any of false.
	HasNextPage bool `json:"has_next_page"`
	// Any of "".
	NextCursor string `json:"next_cursor"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Diagnostic  respjson.Field
		HasNextPage respjson.Field
		NextCursor  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.PaginatedUsers
}

// Returns the unmodified JSON received from the API
func (r XListGetFollowersResponseUserListCoverageResponse) RawJSON() string { return r.JSON.raw }
func (r *XListGetFollowersResponseUserListCoverageResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// Coverage evidence across parallel relationship strategies.
type XListGetFollowersResponseUserListCoverageResponseDiagnostic struct {
	// True after all active strategies exhaust their sources.
	Complete            bool  `json:"complete" api:"required"`
	CursorFailureCount  int64 `json:"cursorFailureCount" api:"required"`
	DeadlineReached     bool  `json:"deadlineReached" api:"required"`
	DuplicateCount      int64 `json:"duplicateCount" api:"required"`
	FailedStrategyCount int64 `json:"failedStrategyCount" api:"required"`
	MalformedCount      int64 `json:"malformedCount" api:"required"`
	PagesFetched        int64 `json:"pagesFetched" api:"required"`
	// True when credits or the requested limit reduce output.
	ResponseTruncated    bool                                                                  `json:"responseTruncated" api:"required"`
	ResultLimitReached   bool                                                                  `json:"resultLimitReached" api:"required"`
	ReturnedUsers        int64                                                                 `json:"returnedUsers" api:"required"`
	StalledStrategyCount int64                                                                 `json:"stalledStrategyCount" api:"required"`
	Strategies           []XListGetFollowersResponseUserListCoverageResponseDiagnosticStrategy `json:"strategies" api:"required"`
	StrategyCount        int64                                                                 `json:"strategyCount" api:"required"`
	UniqueUsers          int64                                                                 `json:"uniqueUsers" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Complete             respjson.Field
		CursorFailureCount   respjson.Field
		DeadlineReached      respjson.Field
		DuplicateCount       respjson.Field
		FailedStrategyCount  respjson.Field
		MalformedCount       respjson.Field
		PagesFetched         respjson.Field
		ResponseTruncated    respjson.Field
		ResultLimitReached   respjson.Field
		ReturnedUsers        respjson.Field
		StalledStrategyCount respjson.Field
		Strategies           respjson.Field
		StrategyCount        respjson.Field
		UniqueUsers          respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XListGetFollowersResponseUserListCoverageResponseDiagnostic) RawJSON() string {
	return r.JSON.raw
}
func (r *XListGetFollowersResponseUserListCoverageResponseDiagnostic) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// Result counts and stop reason for one relationship strategy.
type XListGetFollowersResponseUserListCoverageResponseDiagnosticStrategy struct {
	DuplicateCount int64 `json:"duplicateCount" api:"required"`
	PagesFetched   int64 `json:"pagesFetched" api:"required"`
	// Reason a coverage strategy stopped.
	//
	// Any of "cursor_failure", "deadline", "exhausted", "failed", "page_limit",
	// "result_limit", "stalled".
	StopReason  string `json:"stopReason" api:"required"`
	Strategy    int64  `json:"strategy" api:"required"`
	UniqueAdded int64  `json:"uniqueAdded" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DuplicateCount respjson.Field
		PagesFetched   respjson.Field
		StopReason     respjson.Field
		Strategy       respjson.Field
		UniqueAdded    respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XListGetFollowersResponseUserListCoverageResponseDiagnosticStrategy) RawJSON() string {
	return r.JSON.raw
}
func (r *XListGetFollowersResponseUserListCoverageResponseDiagnosticStrategy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// XListGetMembersResponseUnion contains all possible properties and values from
// [shared.PaginatedUsers], [XListGetMembersResponseUserListCoverageResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type XListGetMembersResponseUnion struct {
	HasNextPage bool   `json:"has_next_page"`
	NextCursor  string `json:"next_cursor"`
	// Present on the last page of a filtered automatic read that sent no maxScan, once
	// the list scanned 100,000 profiles across its pages. That page is the list's
	// last.
	ScanLimit XListGetMembersResponseScanLimit `json:"scanLimit"`
	// This field is from variant [shared.PaginatedUsers],
	// [XListGetMembersResponseUserListCoverageResponse].
	Users []shared.UserProfile `json:"users"`
	// This field is from variant [shared.PaginatedUsers],
	// [XListGetMembersResponseUserListCoverageResponse].
	FilteredCount int64 `json:"filtered_count"`
	// This field is from variant [XListGetMembersResponseUserListCoverageResponse].
	Diagnostic XListGetMembersResponseUserListCoverageResponseDiagnostic `json:"diagnostic"`
	JSON       struct {
		HasNextPage   respjson.Field
		NextCursor    respjson.Field
		ScanLimit     respjson.Field
		Users         respjson.Field
		FilteredCount respjson.Field
		Diagnostic    respjson.Field
		raw           string
	} `json:"-"`
}

func (u XListGetMembersResponseUnion) AsPaginatedUsers() (v shared.PaginatedUsers) {
	_ = v.UnmarshalJSON(json.RawMessage(u.JSON.raw))
	return
}

func (u XListGetMembersResponseUnion) AsXListGetMembersResponseUserListCoverageResponse() (v XListGetMembersResponseUserListCoverageResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v, &v.JSON.raw)
	return
}

// Returns the unmodified JSON received from the API
func (u XListGetMembersResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *XListGetMembersResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// Present on the last page of a filtered automatic read that sent no maxScan, once
// the list scanned 100,000 profiles across its pages. That page is the list's
// last.
type XListGetMembersResponseScanLimit struct {
	// The profiles a filtered read scans of a list without maxScan.
	//
	// Any of 100000.
	Limit int64 `json:"limit" api:"required"`
	// What stopped the list & how to read further.
	Message string `json:"message" api:"required"`
	// Profiles the list read across its pages, kept or dropped by the filters. The
	// last page may read past the limit.
	RowsRead int64 `json:"rowsRead" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Limit       respjson.Field
		Message     respjson.Field
		RowsRead    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XListGetMembersResponseScanLimit) RawJSON() string { return r.JSON.raw }
func (r *XListGetMembersResponseScanLimit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// Terminal relationship coverage response with diagnostics.
type XListGetMembersResponseUserListCoverageResponse struct {
	// Coverage evidence across parallel relationship strategies.
	Diagnostic XListGetMembersResponseUserListCoverageResponseDiagnostic `json:"diagnostic" api:"required"`
	// Any of false.
	HasNextPage bool `json:"has_next_page"`
	// Any of "".
	NextCursor string `json:"next_cursor"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Diagnostic  respjson.Field
		HasNextPage respjson.Field
		NextCursor  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.PaginatedUsers
}

// Returns the unmodified JSON received from the API
func (r XListGetMembersResponseUserListCoverageResponse) RawJSON() string { return r.JSON.raw }
func (r *XListGetMembersResponseUserListCoverageResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// Coverage evidence across parallel relationship strategies.
type XListGetMembersResponseUserListCoverageResponseDiagnostic struct {
	// True after all active strategies exhaust their sources.
	Complete            bool  `json:"complete" api:"required"`
	CursorFailureCount  int64 `json:"cursorFailureCount" api:"required"`
	DeadlineReached     bool  `json:"deadlineReached" api:"required"`
	DuplicateCount      int64 `json:"duplicateCount" api:"required"`
	FailedStrategyCount int64 `json:"failedStrategyCount" api:"required"`
	MalformedCount      int64 `json:"malformedCount" api:"required"`
	PagesFetched        int64 `json:"pagesFetched" api:"required"`
	// True when credits or the requested limit reduce output.
	ResponseTruncated    bool                                                                `json:"responseTruncated" api:"required"`
	ResultLimitReached   bool                                                                `json:"resultLimitReached" api:"required"`
	ReturnedUsers        int64                                                               `json:"returnedUsers" api:"required"`
	StalledStrategyCount int64                                                               `json:"stalledStrategyCount" api:"required"`
	Strategies           []XListGetMembersResponseUserListCoverageResponseDiagnosticStrategy `json:"strategies" api:"required"`
	StrategyCount        int64                                                               `json:"strategyCount" api:"required"`
	UniqueUsers          int64                                                               `json:"uniqueUsers" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Complete             respjson.Field
		CursorFailureCount   respjson.Field
		DeadlineReached      respjson.Field
		DuplicateCount       respjson.Field
		FailedStrategyCount  respjson.Field
		MalformedCount       respjson.Field
		PagesFetched         respjson.Field
		ResponseTruncated    respjson.Field
		ResultLimitReached   respjson.Field
		ReturnedUsers        respjson.Field
		StalledStrategyCount respjson.Field
		Strategies           respjson.Field
		StrategyCount        respjson.Field
		UniqueUsers          respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XListGetMembersResponseUserListCoverageResponseDiagnostic) RawJSON() string {
	return r.JSON.raw
}
func (r *XListGetMembersResponseUserListCoverageResponseDiagnostic) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// Result counts and stop reason for one relationship strategy.
type XListGetMembersResponseUserListCoverageResponseDiagnosticStrategy struct {
	DuplicateCount int64 `json:"duplicateCount" api:"required"`
	PagesFetched   int64 `json:"pagesFetched" api:"required"`
	// Reason a coverage strategy stopped.
	//
	// Any of "cursor_failure", "deadline", "exhausted", "failed", "page_limit",
	// "result_limit", "stalled".
	StopReason  string `json:"stopReason" api:"required"`
	Strategy    int64  `json:"strategy" api:"required"`
	UniqueAdded int64  `json:"uniqueAdded" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DuplicateCount respjson.Field
		PagesFetched   respjson.Field
		StopReason     respjson.Field
		Strategy       respjson.Field
		UniqueAdded    respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XListGetMembersResponseUserListCoverageResponseDiagnosticStrategy) RawJSON() string {
	return r.JSON.raw
}
func (r *XListGetMembersResponseUserListCoverageResponseDiagnosticStrategy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type XListGetResponse struct {
	// Public X List with its owner's public profile. Leaves out what X tells the
	// reading account, such as whether it follows the List.
	List XListGetResponseList `json:"list" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		List        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XListGetResponse) RawJSON() string { return r.JSON.raw }
func (r *XListGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// Public X List with its owner's public profile. Leaves out what X tells the
// reading account, such as whether it follows the List.
type XListGetResponseList struct {
	// Numeric ID of the List.
	ID string `json:"id" api:"required"`
	// Name the owner gave the List.
	Name string `json:"name" api:"required"`
	// URL of the banner the owner uploaded. Absent when the owner uploaded none.
	BannerURL string `json:"bannerUrl"`
	// When the owner made the List.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Banner the owner uploaded.
	CustomBanner XListGetResponseListCustomBanner `json:"customBanner"`
	// Banner X shows when the owner uploaded none. X picks it per reading account, so
	// 2 calls can return different ones.
	DefaultBanner XListGetResponseListDefaultBanner `json:"defaultBanner"`
	// About text of the List. Empty when the owner wrote none.
	Description string `json:"description"`
	// Avatar URLs of a few List members, which X may pick per request.
	FacepileURLs []string `json:"facepileUrls"`
	// Accounts on the List.
	MemberCount int64 `json:"memberCount"`
	// Member count as X words it, such as 35 members.
	MembersContext string `json:"membersContext"`
	// Public profile of the account that owns the List.
	Owner shared.UserProfile `json:"owner"`
	// Accounts that follow the List, which X calls followers.
	SubscriberCount int64 `json:"subscriberCount"`
	// List page on x.com.
	URL string `json:"url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		Name            respjson.Field
		BannerURL       respjson.Field
		CreatedAt       respjson.Field
		CustomBanner    respjson.Field
		DefaultBanner   respjson.Field
		Description     respjson.Field
		FacepileURLs    respjson.Field
		MemberCount     respjson.Field
		MembersContext  respjson.Field
		Owner           respjson.Field
		SubscriberCount respjson.Field
		URL             respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XListGetResponseList) RawJSON() string { return r.JSON.raw }
func (r *XListGetResponseList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// Banner the owner uploaded.
type XListGetResponseListCustomBanner struct {
	// Banner image URL.
	URL string `json:"url" api:"required"`
	// Banner height in pixels.
	Height int64 `json:"height"`
	// Media ID of an uploaded banner.
	MediaID string `json:"mediaId"`
	// Stable media key of an uploaded banner.
	MediaKey string `json:"mediaKey"`
	// Dominant colors of an uploaded banner and their shares.
	Palette []XListGetResponseListCustomBannerPalette `json:"palette"`
	// Area X marks as the subject of the banner, in pixels.
	SalientRect XListGetResponseListCustomBannerSalientRect `json:"salientRect"`
	// Banner width in pixels.
	Width int64 `json:"width"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		URL         respjson.Field
		Height      respjson.Field
		MediaID     respjson.Field
		MediaKey    respjson.Field
		Palette     respjson.Field
		SalientRect respjson.Field
		Width       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XListGetResponseListCustomBanner) RawJSON() string { return r.JSON.raw }
func (r *XListGetResponseListCustomBanner) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type XListGetResponseListCustomBannerPalette struct {
	// Share of the banner
	Percentage float64 `json:"percentage"`
	// The color's red, green and blue values.
	Rgb XListGetResponseListCustomBannerPaletteRgb `json:"rgb"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Percentage  respjson.Field
		Rgb         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XListGetResponseListCustomBannerPalette) RawJSON() string { return r.JSON.raw }
func (r *XListGetResponseListCustomBannerPalette) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// The color's red, green and blue values.
type XListGetResponseListCustomBannerPaletteRgb struct {
	Blue  int64 `json:"blue"`
	Green int64 `json:"green"`
	Red   int64 `json:"red"`
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
func (r XListGetResponseListCustomBannerPaletteRgb) RawJSON() string { return r.JSON.raw }
func (r *XListGetResponseListCustomBannerPaletteRgb) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// Area X marks as the subject of the banner, in pixels.
type XListGetResponseListCustomBannerSalientRect struct {
	// Height of the marked area.
	Height int64 `json:"height"`
	// Pixels from the banner's left.
	Left int64 `json:"left"`
	// Pixels from the banner's top.
	Top int64 `json:"top"`
	// Width of the marked area.
	Width int64 `json:"width"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Height      respjson.Field
		Left        respjson.Field
		Top         respjson.Field
		Width       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XListGetResponseListCustomBannerSalientRect) RawJSON() string { return r.JSON.raw }
func (r *XListGetResponseListCustomBannerSalientRect) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// Banner X shows when the owner uploaded none. X picks it per reading account, so
// 2 calls can return different ones.
type XListGetResponseListDefaultBanner struct {
	// Banner image URL.
	URL string `json:"url" api:"required"`
	// Banner height in pixels.
	Height int64 `json:"height"`
	// Media ID of an uploaded banner.
	MediaID string `json:"mediaId"`
	// Stable media key of an uploaded banner.
	MediaKey string `json:"mediaKey"`
	// Dominant colors of an uploaded banner and their shares.
	Palette []XListGetResponseListDefaultBannerPalette `json:"palette"`
	// Area X marks as the subject of the banner, in pixels.
	SalientRect XListGetResponseListDefaultBannerSalientRect `json:"salientRect"`
	// Banner width in pixels.
	Width int64 `json:"width"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		URL         respjson.Field
		Height      respjson.Field
		MediaID     respjson.Field
		MediaKey    respjson.Field
		Palette     respjson.Field
		SalientRect respjson.Field
		Width       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XListGetResponseListDefaultBanner) RawJSON() string { return r.JSON.raw }
func (r *XListGetResponseListDefaultBanner) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type XListGetResponseListDefaultBannerPalette struct {
	// Share of the banner
	Percentage float64 `json:"percentage"`
	// The color's red, green and blue values.
	Rgb XListGetResponseListDefaultBannerPaletteRgb `json:"rgb"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Percentage  respjson.Field
		Rgb         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XListGetResponseListDefaultBannerPalette) RawJSON() string { return r.JSON.raw }
func (r *XListGetResponseListDefaultBannerPalette) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// The color's red, green and blue values.
type XListGetResponseListDefaultBannerPaletteRgb struct {
	Blue  int64 `json:"blue"`
	Green int64 `json:"green"`
	Red   int64 `json:"red"`
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
func (r XListGetResponseListDefaultBannerPaletteRgb) RawJSON() string { return r.JSON.raw }
func (r *XListGetResponseListDefaultBannerPaletteRgb) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// Area X marks as the subject of the banner, in pixels.
type XListGetResponseListDefaultBannerSalientRect struct {
	// Height of the marked area.
	Height int64 `json:"height"`
	// Pixels from the banner's left.
	Left int64 `json:"left"`
	// Pixels from the banner's top.
	Top int64 `json:"top"`
	// Width of the marked area.
	Width int64 `json:"width"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Height      respjson.Field
		Left        respjson.Field
		Top         respjson.Field
		Width       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XListGetResponseListDefaultBannerSalientRect) RawJSON() string { return r.JSON.raw }
func (r *XListGetResponseListDefaultBannerSalientRect) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// 1 page of public X Lists. Follow next_cursor while the response reports more
// pages. Each returned List costs 1 credit.
type XListGetSearchResponse struct {
	HasNextPage bool                         `json:"has_next_page" api:"required"`
	Lists       []XListGetSearchResponseList `json:"lists" api:"required"`
	NextCursor  string                       `json:"next_cursor" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		HasNextPage respjson.Field
		Lists       respjson.Field
		NextCursor  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XListGetSearchResponse) RawJSON() string { return r.JSON.raw }
func (r *XListGetSearchResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// Public X List with its owner's public profile. Leaves out what X tells the
// reading account, such as whether it follows the List.
type XListGetSearchResponseList struct {
	// Numeric ID of the List.
	ID string `json:"id" api:"required"`
	// Name the owner gave the List.
	Name string `json:"name" api:"required"`
	// URL of the banner the owner uploaded. Absent when the owner uploaded none.
	BannerURL string `json:"bannerUrl"`
	// When the owner made the List.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Banner the owner uploaded.
	CustomBanner XListGetSearchResponseListCustomBanner `json:"customBanner"`
	// Banner X shows when the owner uploaded none. X picks it per reading account, so
	// 2 calls can return different ones.
	DefaultBanner XListGetSearchResponseListDefaultBanner `json:"defaultBanner"`
	// About text of the List. Empty when the owner wrote none.
	Description string `json:"description"`
	// Avatar URLs of a few List members, which X may pick per request.
	FacepileURLs []string `json:"facepileUrls"`
	// Accounts on the List.
	MemberCount int64 `json:"memberCount"`
	// Member count as X words it, such as 35 members.
	MembersContext string `json:"membersContext"`
	// Public profile of the account that owns the List.
	Owner shared.UserProfile `json:"owner"`
	// Accounts that follow the List, which X calls followers.
	SubscriberCount int64 `json:"subscriberCount"`
	// List page on x.com.
	URL string `json:"url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		Name            respjson.Field
		BannerURL       respjson.Field
		CreatedAt       respjson.Field
		CustomBanner    respjson.Field
		DefaultBanner   respjson.Field
		Description     respjson.Field
		FacepileURLs    respjson.Field
		MemberCount     respjson.Field
		MembersContext  respjson.Field
		Owner           respjson.Field
		SubscriberCount respjson.Field
		URL             respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XListGetSearchResponseList) RawJSON() string { return r.JSON.raw }
func (r *XListGetSearchResponseList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// Banner the owner uploaded.
type XListGetSearchResponseListCustomBanner struct {
	// Banner image URL.
	URL string `json:"url" api:"required"`
	// Banner height in pixels.
	Height int64 `json:"height"`
	// Media ID of an uploaded banner.
	MediaID string `json:"mediaId"`
	// Stable media key of an uploaded banner.
	MediaKey string `json:"mediaKey"`
	// Dominant colors of an uploaded banner and their shares.
	Palette []XListGetSearchResponseListCustomBannerPalette `json:"palette"`
	// Area X marks as the subject of the banner, in pixels.
	SalientRect XListGetSearchResponseListCustomBannerSalientRect `json:"salientRect"`
	// Banner width in pixels.
	Width int64 `json:"width"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		URL         respjson.Field
		Height      respjson.Field
		MediaID     respjson.Field
		MediaKey    respjson.Field
		Palette     respjson.Field
		SalientRect respjson.Field
		Width       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XListGetSearchResponseListCustomBanner) RawJSON() string { return r.JSON.raw }
func (r *XListGetSearchResponseListCustomBanner) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type XListGetSearchResponseListCustomBannerPalette struct {
	// Share of the banner
	Percentage float64 `json:"percentage"`
	// The color's red, green and blue values.
	Rgb XListGetSearchResponseListCustomBannerPaletteRgb `json:"rgb"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Percentage  respjson.Field
		Rgb         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XListGetSearchResponseListCustomBannerPalette) RawJSON() string { return r.JSON.raw }
func (r *XListGetSearchResponseListCustomBannerPalette) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// The color's red, green and blue values.
type XListGetSearchResponseListCustomBannerPaletteRgb struct {
	Blue  int64 `json:"blue"`
	Green int64 `json:"green"`
	Red   int64 `json:"red"`
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
func (r XListGetSearchResponseListCustomBannerPaletteRgb) RawJSON() string { return r.JSON.raw }
func (r *XListGetSearchResponseListCustomBannerPaletteRgb) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// Area X marks as the subject of the banner, in pixels.
type XListGetSearchResponseListCustomBannerSalientRect struct {
	// Height of the marked area.
	Height int64 `json:"height"`
	// Pixels from the banner's left.
	Left int64 `json:"left"`
	// Pixels from the banner's top.
	Top int64 `json:"top"`
	// Width of the marked area.
	Width int64 `json:"width"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Height      respjson.Field
		Left        respjson.Field
		Top         respjson.Field
		Width       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XListGetSearchResponseListCustomBannerSalientRect) RawJSON() string { return r.JSON.raw }
func (r *XListGetSearchResponseListCustomBannerSalientRect) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// Banner X shows when the owner uploaded none. X picks it per reading account, so
// 2 calls can return different ones.
type XListGetSearchResponseListDefaultBanner struct {
	// Banner image URL.
	URL string `json:"url" api:"required"`
	// Banner height in pixels.
	Height int64 `json:"height"`
	// Media ID of an uploaded banner.
	MediaID string `json:"mediaId"`
	// Stable media key of an uploaded banner.
	MediaKey string `json:"mediaKey"`
	// Dominant colors of an uploaded banner and their shares.
	Palette []XListGetSearchResponseListDefaultBannerPalette `json:"palette"`
	// Area X marks as the subject of the banner, in pixels.
	SalientRect XListGetSearchResponseListDefaultBannerSalientRect `json:"salientRect"`
	// Banner width in pixels.
	Width int64 `json:"width"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		URL         respjson.Field
		Height      respjson.Field
		MediaID     respjson.Field
		MediaKey    respjson.Field
		Palette     respjson.Field
		SalientRect respjson.Field
		Width       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XListGetSearchResponseListDefaultBanner) RawJSON() string { return r.JSON.raw }
func (r *XListGetSearchResponseListDefaultBanner) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type XListGetSearchResponseListDefaultBannerPalette struct {
	// Share of the banner
	Percentage float64 `json:"percentage"`
	// The color's red, green and blue values.
	Rgb XListGetSearchResponseListDefaultBannerPaletteRgb `json:"rgb"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Percentage  respjson.Field
		Rgb         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XListGetSearchResponseListDefaultBannerPalette) RawJSON() string { return r.JSON.raw }
func (r *XListGetSearchResponseListDefaultBannerPalette) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// The color's red, green and blue values.
type XListGetSearchResponseListDefaultBannerPaletteRgb struct {
	Blue  int64 `json:"blue"`
	Green int64 `json:"green"`
	Red   int64 `json:"red"`
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
func (r XListGetSearchResponseListDefaultBannerPaletteRgb) RawJSON() string { return r.JSON.raw }
func (r *XListGetSearchResponseListDefaultBannerPaletteRgb) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// Area X marks as the subject of the banner, in pixels.
type XListGetSearchResponseListDefaultBannerSalientRect struct {
	// Height of the marked area.
	Height int64 `json:"height"`
	// Pixels from the banner's left.
	Left int64 `json:"left"`
	// Pixels from the banner's top.
	Top int64 `json:"top"`
	// Width of the marked area.
	Width int64 `json:"width"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Height      respjson.Field
		Left        respjson.Field
		Top         respjson.Field
		Width       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XListGetSearchResponseListDefaultBannerSalientRect) RawJSON() string { return r.JSON.raw }
func (r *XListGetSearchResponseListDefaultBannerSalientRect) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type XListGetFollowersParams struct {
	// Legacy cursor alias. Prefer cursor.
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// Match any comma-separated or line-separated bio term, ignoring case.
	BioContains param.Opt[string] `query:"bioContains,omitzero" json:"-"`
	// Only return profiles X marks as blue verified. X marks many business &
	// government accounts blue too. Use verifiedType for 1 badge type.
	BlueVerifiedOnly param.Opt[bool] `query:"blueVerifiedOnly,omitzero" json:"-"`
	// Only return profiles that publish an email outside free webmail, such as Gmail.
	BusinessEmailOnly param.Opt[bool] `query:"businessEmailOnly,omitzero" json:"-"`
	// Cursor from the previous response. Xquik cursors resume automatic coverage.
	// Existing unprefixed cursors keep legacy standard behavior.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Only return profiles with an email at one of these domains. Separate domains
	// with commas. Subdomains match.
	EmailDomains param.Opt[string] `query:"emailDomains,omitzero" json:"-"`
	// Fields to keep in each row, comma-separated. Use dots for nested fields, such as
	// author.username. A list keeps the field in each item. Case & underscores do not
	// matter. A row that lacks a field leaves it out. Cursors, counts & diagnostics
	// always stay. An unknown name answers 400 unknown_field with the valid names,
	// before any read. The price stays the same. SDK row types then hold only the
	// named fields.
	Fields param.Opt[string] `query:"fields,omitzero" json:"-"`
	// Only return profiles that publish an email.
	HasEmail param.Opt[bool] `query:"hasEmail,omitzero" json:"-"`
	// Only return profiles with a location.
	HasLocation param.Opt[bool] `query:"hasLocation,omitzero" json:"-"`
	// Only return profiles with a website.
	HasWebsite param.Opt[bool] `query:"hasWebsite,omitzero" json:"-"`
	// Match a location substring, ignoring case.
	LocationContains param.Opt[string] `query:"locationContains,omitzero" json:"-"`
	// Maximum follower count. Missing counts pass this maximum.
	MaxFollowers param.Opt[int64] `query:"maxFollowers,omitzero" json:"-"`
	// Profiles may follow at most this many accounts.
	MaxFollowing param.Opt[int64] `query:"maxFollowing,omitzero" json:"-"`
	// Most profiles 1 request reads, kept or dropped by the filters. The request then
	// answers the profiles it kept, the count it dropped & the cursor to read on. Each
	// request of a cursor may send its own value. Without it, a filtered read scans at
	// most 100,000 profiles of the list across all its pages, & the page that reaches
	// them ends the list with `scanLimit`. An unfiltered read has no scan limit.
	// Automatic mode only.
	MaxScan param.Opt[int64] `query:"maxScan,omitzero" json:"-"`
	// Maximum post count. maxPosts is also accepted.
	MaxStatuses param.Opt[int64] `query:"maxStatuses,omitzero" json:"-"`
	// Minimum account age in whole days.
	MinAccountAgeDays param.Opt[int64] `query:"minAccountAgeDays,omitzero" json:"-"`
	// Minimum follower count. Filtering happens before billing.
	MinFollowers param.Opt[int64] `query:"minFollowers,omitzero" json:"-"`
	// Profiles must follow at least this many accounts.
	MinFollowing param.Opt[int64] `query:"minFollowing,omitzero" json:"-"`
	// Minimum post count. minPosts is also accepted.
	MinStatuses param.Opt[int64] `query:"minStatuses,omitzero" json:"-"`
	// Fields to drop from each row, comma-separated. It takes the same names as
	// fields. With fields, it drops from the kept fields.
	Omit param.Opt[string] `query:"omit,omitzero" json:"-"`
	// Maximum user profiles: automatic 300; standard 200. Sources return fewer
	// profiles. Follow next_cursor while the response reports more pages.
	PageSize param.Opt[int64] `query:"pageSize,omitzero" json:"-"`
	// Match a username substring, ignoring case.
	UsernameContains param.Opt[string] `query:"usernameContains,omitzero" json:"-"`
	// Only return verified profiles.
	VerifiedOnly param.Opt[bool] `query:"verifiedOnly,omitzero" json:"-"`
	// Match the verification type exactly, ignoring case.
	VerifiedType param.Opt[string] `query:"verifiedType,omitzero" json:"-"`
	// Omit mode for resumable maximum coverage. Standard keeps legacy pagination.
	// Coverage returns diagnostics once and rejects cursors.
	//
	// Any of "standard", "coverage".
	Mode XListGetFollowersParamsMode `query:"mode,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [XListGetFollowersParams]'s query parameters as
// `url.Values`.
func (r XListGetFollowersParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Omit mode for resumable maximum coverage. Standard keeps legacy pagination.
// Coverage returns diagnostics once and rejects cursors.
type XListGetFollowersParamsMode string

const (
	XListGetFollowersParamsModeStandard XListGetFollowersParamsMode = "standard"
	XListGetFollowersParamsModeCoverage XListGetFollowersParamsMode = "coverage"
)

type XListGetMembersParams struct {
	// Legacy cursor alias. Prefer cursor.
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// Match any comma-separated or line-separated bio term, ignoring case.
	BioContains param.Opt[string] `query:"bioContains,omitzero" json:"-"`
	// Only return profiles X marks as blue verified. X marks many business &
	// government accounts blue too. Use verifiedType for 1 badge type.
	BlueVerifiedOnly param.Opt[bool] `query:"blueVerifiedOnly,omitzero" json:"-"`
	// Only return profiles that publish an email outside free webmail, such as Gmail.
	BusinessEmailOnly param.Opt[bool] `query:"businessEmailOnly,omitzero" json:"-"`
	// Cursor from the previous response. Xquik cursors resume automatic coverage.
	// Existing unprefixed cursors keep legacy standard behavior.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Only return profiles with an email at one of these domains. Separate domains
	// with commas. Subdomains match.
	EmailDomains param.Opt[string] `query:"emailDomains,omitzero" json:"-"`
	// Fields to keep in each row, comma-separated. Use dots for nested fields, such as
	// author.username. A list keeps the field in each item. Case & underscores do not
	// matter. A row that lacks a field leaves it out. Cursors, counts & diagnostics
	// always stay. An unknown name answers 400 unknown_field with the valid names,
	// before any read. The price stays the same. SDK row types then hold only the
	// named fields.
	Fields param.Opt[string] `query:"fields,omitzero" json:"-"`
	// Only return profiles that publish an email.
	HasEmail param.Opt[bool] `query:"hasEmail,omitzero" json:"-"`
	// Only return profiles with a location.
	HasLocation param.Opt[bool] `query:"hasLocation,omitzero" json:"-"`
	// Only return profiles with a website.
	HasWebsite param.Opt[bool] `query:"hasWebsite,omitzero" json:"-"`
	// Match a location substring, ignoring case.
	LocationContains param.Opt[string] `query:"locationContains,omitzero" json:"-"`
	// Maximum follower count. Missing counts pass this maximum.
	MaxFollowers param.Opt[int64] `query:"maxFollowers,omitzero" json:"-"`
	// Profiles may follow at most this many accounts.
	MaxFollowing param.Opt[int64] `query:"maxFollowing,omitzero" json:"-"`
	// Most profiles 1 request reads, kept or dropped by the filters. The request then
	// answers the profiles it kept, the count it dropped & the cursor to read on. Each
	// request of a cursor may send its own value. Without it, a filtered read scans at
	// most 100,000 profiles of the list across all its pages, & the page that reaches
	// them ends the list with `scanLimit`. An unfiltered read has no scan limit.
	// Automatic mode only.
	MaxScan param.Opt[int64] `query:"maxScan,omitzero" json:"-"`
	// Maximum post count. maxPosts is also accepted.
	MaxStatuses param.Opt[int64] `query:"maxStatuses,omitzero" json:"-"`
	// Minimum account age in whole days.
	MinAccountAgeDays param.Opt[int64] `query:"minAccountAgeDays,omitzero" json:"-"`
	// Minimum follower count. Filtering happens before billing.
	MinFollowers param.Opt[int64] `query:"minFollowers,omitzero" json:"-"`
	// Profiles must follow at least this many accounts.
	MinFollowing param.Opt[int64] `query:"minFollowing,omitzero" json:"-"`
	// Minimum post count. minPosts is also accepted.
	MinStatuses param.Opt[int64] `query:"minStatuses,omitzero" json:"-"`
	// Fields to drop from each row, comma-separated. It takes the same names as
	// fields. With fields, it drops from the kept fields.
	Omit param.Opt[string] `query:"omit,omitzero" json:"-"`
	// Maximum user profiles: automatic 300; standard 200. Sources return fewer
	// profiles. Follow next_cursor while the response reports more pages.
	PageSize param.Opt[int64] `query:"pageSize,omitzero" json:"-"`
	// Match a username substring, ignoring case.
	UsernameContains param.Opt[string] `query:"usernameContains,omitzero" json:"-"`
	// Only return verified profiles.
	VerifiedOnly param.Opt[bool] `query:"verifiedOnly,omitzero" json:"-"`
	// Match the verification type exactly, ignoring case.
	VerifiedType param.Opt[string] `query:"verifiedType,omitzero" json:"-"`
	// Omit mode for resumable maximum coverage. Standard keeps legacy pagination.
	// Coverage returns diagnostics once and rejects cursors.
	//
	// Any of "standard", "coverage".
	Mode XListGetMembersParamsMode `query:"mode,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [XListGetMembersParams]'s query parameters as `url.Values`.
func (r XListGetMembersParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Omit mode for resumable maximum coverage. Standard keeps legacy pagination.
// Coverage returns diagnostics once and rejects cursors.
type XListGetMembersParamsMode string

const (
	XListGetMembersParamsModeStandard XListGetMembersParamsMode = "standard"
	XListGetMembersParamsModeCoverage XListGetMembersParamsMode = "coverage"
)

type XListGetTweetsParams struct {
	// Raw advanced search query appended as-is.
	AdvancedQuery param.Opt[string] `query:"advancedQuery,omitzero" json:"-"`
	// Legacy cursor alias. Prefer cursor.
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// Words or quoted phrases where any one can match. Separate with spaces, commas,
	// or lines.
	AnyWords param.Opt[string] `query:"anyWords,omitzero" json:"-"`
	// Only return tweets from Blue-verified authors.
	BlueVerifiedOnly param.Opt[bool] `query:"blueVerifiedOnly,omitzero" json:"-"`
	// Geo bounding box, such as -74.1 40.6 -73.9 40.8.
	BoundingBox param.Opt[string] `query:"boundingBox,omitzero" json:"-"`
	// Keep only Tweets with this card type, such as poll2choice_text_only. Tweet
	// search checks each Tweet's card.
	CardName param.Opt[string] `query:"cardName,omitzero" json:"-"`
	// Cashtags separated by spaces, commas, or lines.
	Cashtags param.Opt[string] `query:"cashtags,omitzero" json:"-"`
	// Conversation ID filter.
	ConversationID param.Opt[string] `query:"conversationId,omitzero" json:"-"`
	// Cursor from the previous response. Xquik cursors resume automatic coverage.
	// Existing unprefixed cursors keep legacy standard behavior.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Match this literal phrase, including any hyphens.
	ExactPhrase param.Opt[string] `query:"exactPhrase,omitzero" json:"-"`
	// X search no longer supports excludeSource. A Tweet search with it answers 424 &
	// charges nothing.
	ExcludeSource param.Opt[string] `query:"excludeSource,omitzero" json:"-"`
	// Words or quoted phrases to exclude. Separate with spaces, commas, or lines.
	ExcludeWords param.Opt[string] `query:"excludeWords,omitzero" json:"-"`
	// Fields to keep in each row, comma-separated. Use dots for nested fields, such as
	// author.username. A list keeps the field in each item. Case & underscores do not
	// matter. A row that lacks a field leaves it out. Cursors, counts & diagnostics
	// always stay. An unknown name answers 400 unknown_field with the valid names,
	// before any read. The price stays the same. SDK row types then hold only the
	// named fields.
	Fields param.Opt[string] `query:"fields,omitzero" json:"-"`
	// Filter by author username.
	FromUser param.Opt[string] `query:"fromUser,omitzero" json:"-"`
	// X search no longer supports geocode. A Tweet search with it answers 424 &
	// charges nothing.
	Geocode param.Opt[string] `query:"geocode,omitzero" json:"-"`
	// Hashtags separated by spaces, commas, or lines.
	Hashtags param.Opt[string] `query:"hashtags,omitzero" json:"-"`
	// Include reply tweets unless replies specifies another mode.
	IncludeReplies param.Opt[bool] `query:"includeReplies,omitzero" json:"-"`
	// Only replies to this tweet ID.
	InReplyToTweetID param.Opt[string] `query:"inReplyToTweetId,omitzero" json:"-"`
	// Words the Tweets must match, in X search syntax.
	Keywords param.Opt[string] `query:"keywords,omitzero" json:"-"`
	// Filter by language. Alias `lang` is accepted.
	Language param.Opt[string] `query:"language,omitzero" json:"-"`
	// Maximum likes threshold. maxLikes is also accepted.
	MaxFaves param.Opt[int64] `query:"maxFaves,omitzero" json:"-"`
	// Return Tweets at or below this ID.
	MaxID param.Opt[string] `query:"maxId,omitzero" json:"-"`
	// Maximum quotes threshold.
	MaxQuotes param.Opt[int64] `query:"maxQuotes,omitzero" json:"-"`
	// Maximum replies threshold.
	MaxReplies param.Opt[int64] `query:"maxReplies,omitzero" json:"-"`
	// Maximum retweets threshold.
	MaxRetweets param.Opt[int64] `query:"maxRetweets,omitzero" json:"-"`
	// Filter tweets mentioning a username.
	Mentioning param.Opt[string] `query:"mentioning,omitzero" json:"-"`
	// Minimum bookmark count threshold.
	MinBookmarks param.Opt[int64] `query:"minBookmarks,omitzero" json:"-"`
	// Minimum likes. Aliases: minFaves, min_likes, min_faves.
	MinLikes param.Opt[int64] `query:"minLikes,omitzero" json:"-"`
	// Minimum quote count threshold.
	MinQuotes param.Opt[int64] `query:"minQuotes,omitzero" json:"-"`
	// Minimum replies threshold.
	MinReplies param.Opt[int64] `query:"minReplies,omitzero" json:"-"`
	// Minimum retweets threshold.
	MinRetweets param.Opt[int64] `query:"minRetweets,omitzero" json:"-"`
	// Minimum view count threshold.
	MinViews param.Opt[int64] `query:"minViews,omitzero" json:"-"`
	// Only return native reposts.
	NativeRetweets param.Opt[bool] `query:"nativeRetweets,omitzero" json:"-"`
	// A place name. A country, such as United States, works like placeCountry & keeps
	// Tweets tagged there. A city keeps Tweets about it.
	Near param.Opt[string] `query:"near,omitzero" json:"-"`
	// X no longer searches filter:news.
	News param.Opt[bool] `query:"news,omitzero" json:"-"`
	// Fields to drop from each row, comma-separated. It takes the same names as
	// fields. With fields, it drops from the kept fields.
	Omit param.Opt[string] `query:"omit,omitzero" json:"-"`
	// Automatic pages accept 1-300 Tweets. Standard pages keep 1-100. Default 20.
	// Follow next_cursor while the response reports more pages. Deprecated aliases
	// remain accepted.
	PageSize param.Opt[int64] `query:"pageSize,omitzero" json:"-"`
	// Search within a place ID.
	Place param.Opt[string] `query:"place,omitzero" json:"-"`
	// Search within a country code.
	PlaceCountry param.Opt[string] `query:"placeCountry,omitzero" json:"-"`
	// Geo point radius, such as -73.99 40.73 25mi.
	PointRadius param.Opt[string] `query:"pointRadius,omitzero" json:"-"`
	// Only quotes of this tweet ID.
	QuotesOfTweetID param.Opt[string] `query:"quotesOfTweetId,omitzero" json:"-"`
	// Only retweets of this tweet ID.
	RetweetsOfTweetID param.Opt[string] `query:"retweetsOfTweetId,omitzero" json:"-"`
	// Enable the safe-search filter.
	Safe param.Opt[bool] `query:"safe,omitzero" json:"-"`
	// Start date in YYYY-MM-DD format.
	SinceDate param.Opt[time.Time] `query:"sinceDate,omitzero" format:"date" json:"-"`
	// Return Tweets whose IDs exceed this ID.
	SinceID param.Opt[string] `query:"sinceId,omitzero" json:"-"`
	// Inclusive ISO bound for Tweet creation time.
	SinceTime param.Opt[string] `query:"sinceTime,omitzero" json:"-"`
	// X search no longer supports source. A Tweet search with it answers 424 & charges
	// nothing.
	Source param.Opt[string] `query:"source,omitzero" json:"-"`
	// Filter replies sent to a username.
	ToUser param.Opt[string] `query:"toUser,omitzero" json:"-"`
	// End date in YYYY-MM-DD format.
	UntilDate param.Opt[time.Time] `query:"untilDate,omitzero" format:"date" json:"-"`
	// Exclusive ISO bound for Tweet creation time.
	UntilTime param.Opt[string] `query:"untilTime,omitzero" json:"-"`
	// URL substring or domain filter.
	URL param.Opt[string] `query:"url,omitzero" json:"-"`
	// Only return tweets from verified authors.
	VerifiedOnly param.Opt[bool] `query:"verifiedOnly,omitzero" json:"-"`
	// X search no longer supports within. A Tweet search with it answers 424 & charges
	// nothing. Use near alone.
	Within param.Opt[string] `query:"within,omitzero" json:"-"`
	// Keeps Tweets from a recent window, such as 90m or 7d. Use a whole number & s, m,
	// h or d.
	WithinTime param.Opt[string] `query:"withinTime,omitzero" json:"-"`
	// Filter media. Aliases: has_video, has_media.
	//
	// Any of "images", "videos", "gifs", "media", "links", "none".
	MediaType XListGetTweetsParamsMediaType `query:"mediaType,omitzero" json:"-"`
	// Omit mode for resumable maximum coverage. Standard keeps legacy pagination.
	// Coverage returns diagnostics once and rejects cursors.
	//
	// Any of "standard", "coverage".
	Mode XListGetTweetsParamsMode `query:"mode,omitzero" json:"-"`
	// Only when the caller requests a quote mode.
	//
	// Any of "include", "exclude", "only".
	Quotes XListGetTweetsParamsQuotes `query:"quotes,omitzero" json:"-"`
	// Only when the caller requests a reply mode.
	//
	// Any of "include", "exclude", "only".
	Replies XListGetTweetsParamsReplies `query:"replies,omitzero" json:"-"`
	// Only when the caller requests a repost mode.
	//
	// Any of "include", "exclude", "only".
	Retweets XListGetTweetsParamsRetweets `query:"retweets,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [XListGetTweetsParams]'s query parameters as `url.Values`.
func (r XListGetTweetsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter media. Aliases: has_video, has_media.
type XListGetTweetsParamsMediaType string

const (
	XListGetTweetsParamsMediaTypeImages XListGetTweetsParamsMediaType = "images"
	XListGetTweetsParamsMediaTypeVideos XListGetTweetsParamsMediaType = "videos"
	XListGetTweetsParamsMediaTypeGifs   XListGetTweetsParamsMediaType = "gifs"
	XListGetTweetsParamsMediaTypeMedia  XListGetTweetsParamsMediaType = "media"
	XListGetTweetsParamsMediaTypeLinks  XListGetTweetsParamsMediaType = "links"
	XListGetTweetsParamsMediaTypeNone   XListGetTweetsParamsMediaType = "none"
)

// Omit mode for resumable maximum coverage. Standard keeps legacy pagination.
// Coverage returns diagnostics once and rejects cursors.
type XListGetTweetsParamsMode string

const (
	XListGetTweetsParamsModeStandard XListGetTweetsParamsMode = "standard"
	XListGetTweetsParamsModeCoverage XListGetTweetsParamsMode = "coverage"
)

// Only when the caller requests a reply mode.
type XListGetTweetsParamsReplies string

const (
	XListGetTweetsParamsRepliesInclude XListGetTweetsParamsReplies = "include"
	XListGetTweetsParamsRepliesExclude XListGetTweetsParamsReplies = "exclude"
	XListGetTweetsParamsRepliesOnly    XListGetTweetsParamsReplies = "only"
)

// Only when the caller requests a repost mode.
type XListGetTweetsParamsRetweets string

const (
	XListGetTweetsParamsRetweetsInclude XListGetTweetsParamsRetweets = "include"
	XListGetTweetsParamsRetweetsExclude XListGetTweetsParamsRetweets = "exclude"
	XListGetTweetsParamsRetweetsOnly    XListGetTweetsParamsRetweets = "only"
)

type XListGetParams struct {
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

// URLQuery serializes [XListGetParams]'s query parameters as `url.Values`.
func (r XListGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type XListGetSearchParams struct {
	// Keywords to find Lists by, such as a topic. Alias query.
	Q string `query:"q" api:"required" json:"-"`
	// Pagination cursor for collection results. Also read as next_cursor, next_token,
	// pagination_token or continuation_token.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
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
	// Items per page (1-100, default 20). Source, filters or credits can return fewer.
	// Follow next_cursor for more pages. Its aliases are limit, count, max_results,
	// maxItems, max_items & per_page.
	PageSize param.Opt[int64] `query:"pageSize,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [XListGetSearchParams]'s query parameters as `url.Values`.
func (r XListGetSearchParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Only when the caller requests a quote mode.
type XListGetTweetsParamsQuotes string

const (
	XListGetTweetsParamsQuotesInclude XListGetTweetsParamsQuotes = "include"
	XListGetTweetsParamsQuotesExclude XListGetTweetsParamsQuotes = "exclude"
	XListGetTweetsParamsQuotesOnly    XListGetTweetsParamsQuotes = "only"
)
