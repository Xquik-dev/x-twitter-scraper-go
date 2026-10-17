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
)

// Public X job listings, search and details
//
// XJobService contains methods and other services that help with interacting with
// the x-twitter-scraper API.
//
// You should not instantiate this service directly, and instead use the
// [NewXJobService] method instead.
type XJobService struct {
	options []option.RequestOption
}

// NewXJobService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewXJobService(opts ...option.RequestOption) (r XJobService) {
	r = XJobService{}
	r.options = opts
	return
}

// Returns a public job listing with its description as text & Markdown, work type,
// seniority, salary, company & apply link. A job that closed returns job_not_found
// & is free.
func (r *XJobService) Get(ctx context.Context, id string, query XJobGetParams, opts ...option.RequestOption) (res *XJobGetResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("x/jobs/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Returns the places whose names match a query, each with the location ID that GET
// /x/jobs/search takes. No match returns an empty list, free.
func (r *XJobService) ListLocations(ctx context.Context, query XJobListLocationsParams, opts ...option.RequestOption) (res *XJobListLocationsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "x/jobs/locations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Returns job listings with title, location, salary, company & apply link. Add
// filters to narrow them. A closed job does not appear & costs nothing. Follow
// next_cursor for more.
func (r *XJobService) Search(ctx context.Context, query XJobSearchParams, opts ...option.RequestOption) (res *XJobSearchResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "x/jobs/search"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Public X job listing. A search returns the fields down to the salary, plus
// company & account. A lookup adds work type, seniority, function, team &
// descriptions. A field X leaves empty is left out.
type Job struct {
	// Job ID.
	ID string `json:"id" api:"required"`
	// The job page on X.
	URL string `json:"url" api:"required"`
	// The X account that posted a job. A job has none when X imported it from a job
	// board or no longer shows the account.
	Account JobAccount `json:"account"`
	// Where applicants apply, off X.
	ApplyURL string `json:"applyUrl"`
	// The company profile of a job X imported from a job board.
	Company JobCompany `json:"company"`
	// Full description as plain text, 1 line per paragraph.
	Description string `json:"description"`
	// Full description as Markdown, with headings, lists, links, bold & italics.
	DescriptionMarkdown string `json:"descriptionMarkdown"`
	// Contract type (full_time, full_time_contract, part_time or contract_to_hire).
	EmploymentType string `json:"employmentType"`
	// Whether X features the job.
	Featured bool `json:"featured"`
	// Salary as X shows it.
	FormattedSalary string `json:"formattedSalary"`
	// Job function, such as software_engineering.
	JobFunction string `json:"jobFunction"`
	// Location as the employer wrote it.
	Location string `json:"location"`
	// Where the work happens (onsite, remote or hybrid).
	LocationType string `json:"locationType"`
	// When the employer posted the job.
	PostedAt time.Time `json:"postedAt" format:"date-time"`
	// ISO 4217 currency code of the salary.
	SalaryCurrencyCode string `json:"salaryCurrencyCode"`
	// Pay period of the salary (annually or hourly).
	SalaryInterval string `json:"salaryInterval"`
	// Highest salary, in salaryCurrencyCode per salaryInterval.
	SalaryMax float64 `json:"salaryMax"`
	// Lowest salary, in salaryCurrencyCode per salaryInterval.
	SalaryMin float64 `json:"salaryMin"`
	// Seniority (intern, entry_level, junior, mid_level, senior, lead, manager or
	// executive).
	SeniorityLevel string `json:"seniorityLevel"`
	// Short summary of the job, as plain text.
	ShortDescription string `json:"shortDescription"`
	// Team inside the company.
	Team string `json:"team"`
	// Job title.
	Title string `json:"title"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                  respjson.Field
		URL                 respjson.Field
		Account             respjson.Field
		ApplyURL            respjson.Field
		Company             respjson.Field
		Description         respjson.Field
		DescriptionMarkdown respjson.Field
		EmploymentType      respjson.Field
		Featured            respjson.Field
		FormattedSalary     respjson.Field
		JobFunction         respjson.Field
		Location            respjson.Field
		LocationType        respjson.Field
		PostedAt            respjson.Field
		SalaryCurrencyCode  respjson.Field
		SalaryInterval      respjson.Field
		SalaryMax           respjson.Field
		SalaryMin           respjson.Field
		SeniorityLevel      respjson.Field
		ShortDescription    respjson.Field
		Team                respjson.Field
		Title               respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Job) RawJSON() string { return r.JSON.raw }
func (r *Job) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The X account that posted a job. A job has none when X imported it from a job
// board or no longer shows the account.
type JobAccount struct {
	// X user ID of the account.
	ID string `json:"id" api:"required"`
	// Whether X shows a verification badge on the account.
	Verified bool `json:"verified" api:"required"`
	// Bio of the account. A job lookup returns it.
	Description string `json:"description"`
	// Display name of the account that posted the job.
	Name string `json:"name"`
	// Avatar shape, such as Circle or Square.
	ProfileImageShape string `json:"profileImageShape"`
	// Avatar URL of the account that posted the job.
	ProfilePicture string `json:"profilePicture"`
	// The job account's avatar at full size.
	ProfilePictureOriginal string `json:"profilePictureOriginal"`
	// Handle without @.
	Username string `json:"username"`
	// Badge type, such as Business or Government.
	VerifiedType string `json:"verifiedType"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                     respjson.Field
		Verified               respjson.Field
		Description            respjson.Field
		Name                   respjson.Field
		ProfileImageShape      respjson.Field
		ProfilePicture         respjson.Field
		ProfilePictureOriginal respjson.Field
		Username               respjson.Field
		VerifiedType           respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r JobAccount) RawJSON() string { return r.JSON.raw }
func (r *JobAccount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The company profile of a job X imported from a job board.
type JobCompany struct {
	// X company profile ID.
	ID string `json:"id" api:"required"`
	// Company logo URL.
	LogoURL string `json:"logoUrl"`
	// Company name.
	Name string `json:"name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		LogoURL     respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r JobCompany) RawJSON() string { return r.JSON.raw }
func (r *JobCompany) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type XJobGetResponse struct {
	// Public X job listing. A search returns the fields down to the salary, plus
	// company & account. A lookup adds work type, seniority, function, team &
	// descriptions. A field X leaves empty is left out.
	Job Job `json:"job" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Job         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XJobGetResponse) RawJSON() string { return r.JSON.raw }
func (r *XJobGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type XJobGetResponseJob = Job

type XJobGetResponseJobAccount = JobAccount

type XJobGetResponseJobCompany = JobCompany

// The places whose names match a query.
type XJobListLocationsResponse struct {
	// Places returned, each charged 1 credit.
	Count int64 `json:"count" api:"required"`
	// Matching places, best match first.
	Locations []XJobListLocationsResponseLocation `json:"locations" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Locations   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XJobListLocationsResponse) RawJSON() string { return r.JSON.raw }
func (r *XJobListLocationsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A place a job search can name by its ID.
type XJobListLocationsResponseLocation struct {
	// Location ID, which the job search locationId takes.
	ID string `json:"id" api:"required"`
	// Place name with its region or country code.
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
func (r XJobListLocationsResponseLocation) RawJSON() string { return r.JSON.raw }
func (r *XJobListLocationsResponseLocation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// 1 page of job listings. Follow next_cursor while the response reports more
// pages.
type XJobSearchResponse struct {
	HasNextPage bool   `json:"has_next_page" api:"required"`
	Jobs        []Job  `json:"jobs" api:"required"`
	NextCursor  string `json:"next_cursor" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		HasNextPage respjson.Field
		Jobs        respjson.Field
		NextCursor  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XJobSearchResponse) RawJSON() string { return r.JSON.raw }
func (r *XJobSearchResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type XJobSearchResponseJob = Job

type XJobSearchResponseJobAccount = JobAccount

type XJobSearchResponseJobCompany = JobCompany

type XJobGetParams struct {
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

// URLQuery serializes [XJobGetParams]'s query parameters as `url.Values`.
func (r XJobGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type XJobListLocationsParams struct {
	// Place name or its start, such as new york. Its alias is query.
	Q string `query:"q" api:"required" json:"-"`
	// Fields to keep in each row, comma-separated. Use dots for nested fields, such as
	// author.username. A list keeps the field in each item. Case & underscores do not
	// matter. A row that lacks a field leaves it out. Cursors, counts & diagnostics
	// always stay. An unknown name answers 400 unknown_field with the valid names,
	// before any read. The price stays the same. SDK row types then hold only the
	// named fields.
	Fields param.Opt[string] `query:"fields,omitzero" json:"-"`
	// Most places to return. Without it, every match returns. Its aliases are
	// pageSize, count, max_results, maxItems, max_items & per_page.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Fields to drop from each row, comma-separated. It takes the same names as
	// fields. With fields, it drops from the kept fields.
	Omit param.Opt[string] `query:"omit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [XJobListLocationsParams]'s query parameters as `url.Values`.
func (r XJobListLocationsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type XJobSearchParams struct {
	// Keyword, such as a job title or skill. Its alias is query.
	Q string `query:"q" api:"required" json:"-"`
	// Company name to match.
	Company param.Opt[string] `query:"company,omitzero" json:"-"`
	// Pagination cursor for collection results. Also read as next_cursor, next_token,
	// pagination_token or continuation_token.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Comma-separated list of full_time, full_time_contract, part_time &
	// contract_to_hire. Any case works, with spaces or hyphens between words.
	EmploymentType param.Opt[string] `query:"employmentType,omitzero" json:"-"`
	// Fields to keep in each row, comma-separated. Use dots for nested fields, such as
	// author.username. A list keeps the field in each item. Case & underscores do not
	// matter. A row that lacks a field leaves it out. Cursors, counts & diagnostics
	// always stay. An unknown name answers 400 unknown_field with the valid names,
	// before any read. The price stays the same. SDK row types then hold only the
	// named fields.
	Fields param.Opt[string] `query:"fields,omitzero" json:"-"`
	// City, region or country as text, such as Berlin.
	Location param.Opt[string] `query:"location,omitzero" json:"-"`
	// Location ID from GET /x/jobs/locations. Replaces location when both are sent.
	LocationID param.Opt[string] `query:"locationId,omitzero" json:"-"`
	// Where the work happens. Comma-separated list of onsite, remote & hybrid. Any
	// case works, with spaces or hyphens between words.
	LocationType param.Opt[string] `query:"locationType,omitzero" json:"-"`
	// Fields to drop from each row, comma-separated. It takes the same names as
	// fields. With fields, it drops from the kept fields.
	Omit param.Opt[string] `query:"omit,omitzero" json:"-"`
	// Items per page (1-100, default 20). Source, filters or credits can return fewer.
	// Follow next_cursor for more pages. Its aliases are limit, count, max_results,
	// maxItems, max_items & per_page.
	PageSize param.Opt[int64] `query:"pageSize,omitzero" json:"-"`
	// Comma-separated list of intern, entry_level, junior, mid_level, senior, lead,
	// manager & executive. Any case works, with spaces or hyphens between words.
	Seniority param.Opt[string] `query:"seniority,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [XJobSearchParams]'s query parameters as `url.Values`.
func (r XJobSearchParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
