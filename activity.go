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
	"github.com/Xquik-dev/x-twitter-scraper-go/shared/constant"
)

// Activity log of the account
//
// ActivityService contains methods and other services that help with interacting
// with the x-twitter-scraper API.
//
// You should not instantiate this service directly, and instead use the
// [NewActivityService] method instead.
type ActivityService struct {
	options []option.RequestOption
}

// NewActivityService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewActivityService(opts ...option.RequestOption) (r ActivityService) {
	r = ActivityService{}
	r.options = opts
	return
}

// Returns the account's activity log as JSON rows, newest first, with the
// Activity page's filters. It covers write actions, extractions, monitor changes
// and events, webhooks, giveaways, media, support, API keys, X account connections
// and billing.
func (r *ActivityService) List(ctx context.Context, query ActivityListParams, opts ...option.RequestOption) (res *[]ActivityLogRow, err error) {
	opts = slices.Concat(r.options, opts)
	path := "activity"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Downloads the same activity rows as a CSV, Excel, Markdown, PDF or text file.
func (r *ActivityService) Export(ctx context.Context, query ActivityExportParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/octet-stream")}, opts...)
	path := "activity"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Returns 1 activity of the account with its full technical record. The record
// holds the request, the caller, the charge, every timestamp in UTC and related
// activity. It never holds an API key value, a webhook secret or another account's
// data. Only a signed-in dashboard session sees an api_key record, so an API key
// or an OAuth app gets 404 for its ID.
func (r *ActivityService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *ActivityRecord, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("activity/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// 1 Activity log row. A cut list ends with a row whose category is notice and
// whose activity says where it stopped.
type ActivityLogRow struct {
	// ISO 8601 time in UTC to the millisecond. Empty on the notice row.
	Time     string `json:"time" api:"required"`
	Category string `json:"category" api:"required"`
	Type     string `json:"type" api:"required"`
	Activity string `json:"activity" api:"required"`
	Status   string `json:"status" api:"required"`
	Source   string `json:"source" api:"required"`
	Target   string `json:"target" api:"required"`
	// A number, or an empty string when the activity has no results.
	Results ActivityLogRowResultsUnion `json:"results" api:"required"`
	// A number, or an empty string when the activity has no charge.
	Credits ActivityLogRowCreditsUnion `json:"credits" api:"required"`
	Details string                     `json:"details" api:"required"`
	// Activity ID. Activity.Get returns its full record.
	ID   string `json:"id" api:"required"`
	Link string `json:"link" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Time        respjson.Field
		Category    respjson.Field
		Type        respjson.Field
		Activity    respjson.Field
		Status      respjson.Field
		Source      respjson.Field
		Target      respjson.Field
		Results     respjson.Field
		Credits     respjson.Field
		Details     respjson.Field
		ID          respjson.Field
		Link        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ActivityLogRow) RawJSON() string { return r.JSON.raw }
func (r *ActivityLogRow) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ActivityLogRowResultsUnion contains all possible properties and values from
// [int64], [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a JSON object, OfInt or OfString is valid.
type ActivityLogRowResultsUnion struct {
	// This field will be present if the value is a [int64] instead of an object.
	OfInt int64 `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	JSON     struct {
		OfInt    respjson.Field
		OfString respjson.Field
		raw      string
	} `json:"-"`
}

func (u ActivityLogRowResultsUnion) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ActivityLogRowResultsUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ActivityLogRowResultsUnion) RawJSON() string { return u.JSON.raw }

func (r *ActivityLogRowResultsUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ActivityLogRowCreditsUnion contains all possible properties and values from
// [int64], [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a JSON object, OfInt or OfString is valid.
type ActivityLogRowCreditsUnion struct {
	// This field will be present if the value is a [int64] instead of an object.
	OfInt int64 `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	JSON     struct {
		OfInt    respjson.Field
		OfString respjson.Field
		raw      string
	} `json:"-"`
}

func (u ActivityLogRowCreditsUnion) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ActivityLogRowCreditsUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ActivityLogRowCreditsUnion) RawJSON() string { return u.JSON.raw }

func (r *ActivityLogRowCreditsUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Full technical record of 1 activity. Timestamps are ISO 8601 in UTC to the
// millisecond.
type ActivityRecord struct {
	Object constant.Activity `json:"object" default:"activity"`
	// Activity ID.
	ID string `json:"id" api:"required"`
	// Kind of the activity. Only a signed-in dashboard session sees an api_key
	// record.
	//
	// Any of "write_action", "extraction", "monitor_change", "api_key", "billing",
	// "giveaway_draw", "monitor_event", "media_download", "media_share",
	// "support_ticket", "webhook_delivery", "webhook_endpoint",
	// "x_account_connection".
	Kind ActivityRecordKind `json:"kind" api:"required"`
	// Type of the activity within its kind, the same as the type of its Activity log
	// row. Examples are a write action type, an extraction tool type, a monitor
	// change, a monitor event type or a ledger entry type.
	Type string `json:"type" api:"required"`
	// Any of "completed", "running", "failed", "canceled".
	Status ActivityRecordStatus `json:"status" api:"required"`
	// Dashboard page of this activity.
	URL string `json:"url" api:"required" format:"uri"`
	// ID of the API request behind this activity, or null.
	RequestID  string                   `json:"requestId" api:"required"`
	Timestamps ActivityRecordTimestamps `json:"timestamps" api:"required"`
	// The API request behind this activity. Null for activity from before Oct 30,
	// 2026, and for activity no API request starts, such as a monitor event, a
	// webhook delivery or a credit grant.
	Request ActivityRecordRequest `json:"request" api:"required"`
	Caller  ActivityRecordCaller  `json:"caller" api:"required"`
	// Credits charged for this activity, or null.
	Charge ActivityRecordCharge `json:"charge" api:"required"`
	// Fields of this kind of activity, such as its target. A field can hold JSON,
	// such as the data of a monitor event or the payload of a webhook delivery.
	Details map[string]any `json:"details" api:"required"`
	// Up to 5 other activities on the same subject within 7 days, newest first.
	Related []ActivityRecordRelated `json:"related" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Object      respjson.Field
		ID          respjson.Field
		Kind        respjson.Field
		Type        respjson.Field
		Status      respjson.Field
		URL         respjson.Field
		RequestID   respjson.Field
		Timestamps  respjson.Field
		Request     respjson.Field
		Caller      respjson.Field
		Charge      respjson.Field
		Details     respjson.Field
		Related     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ActivityRecord) RawJSON() string { return r.JSON.raw }
func (r *ActivityRecord) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Kind of the activity. Only a signed-in dashboard session sees an api_key record.
type ActivityRecordKind string

const (
	ActivityRecordKindWriteAction        ActivityRecordKind = "write_action"
	ActivityRecordKindExtraction         ActivityRecordKind = "extraction"
	ActivityRecordKindMonitorChange      ActivityRecordKind = "monitor_change"
	ActivityRecordKindAPIKey             ActivityRecordKind = "api_key"
	ActivityRecordKindBilling            ActivityRecordKind = "billing"
	ActivityRecordKindGiveawayDraw       ActivityRecordKind = "giveaway_draw"
	ActivityRecordKindMonitorEvent       ActivityRecordKind = "monitor_event"
	ActivityRecordKindMediaDownload      ActivityRecordKind = "media_download"
	ActivityRecordKindMediaShare         ActivityRecordKind = "media_share"
	ActivityRecordKindSupportTicket      ActivityRecordKind = "support_ticket"
	ActivityRecordKindWebhookDelivery    ActivityRecordKind = "webhook_delivery"
	ActivityRecordKindWebhookEndpoint    ActivityRecordKind = "webhook_endpoint"
	ActivityRecordKindXAccountConnection ActivityRecordKind = "x_account_connection"
)

type ActivityRecordStatus string

const (
	ActivityRecordStatusCompleted ActivityRecordStatus = "completed"
	ActivityRecordStatusRunning   ActivityRecordStatus = "running"
	ActivityRecordStatusFailed    ActivityRecordStatus = "failed"
	ActivityRecordStatusCanceled  ActivityRecordStatus = "canceled"
)

type ActivityRecordTimestamps struct {
	CreatedAt        time.Time `json:"createdAt" api:"required" format:"date-time"`
	UpdatedAt        time.Time `json:"updatedAt" format:"date-time"`
	CompletedAt      time.Time `json:"completedAt" format:"date-time"`
	ConfirmedAt      time.Time `json:"confirmedAt" format:"date-time"`
	ExpiresAt        time.Time `json:"expiresAt" format:"date-time"`
	OccurredAt       time.Time `json:"occurredAt" format:"date-time"`
	DiscoveredAt     time.Time `json:"discoveredAt" format:"date-time"`
	PostedAt         time.Time `json:"postedAt" format:"date-time"`
	FirstAttemptedAt time.Time `json:"firstAttemptedAt" format:"date-time"`
	DeliveredAt      time.Time `json:"deliveredAt" format:"date-time"`
	NextRetryAt      time.Time `json:"nextRetryAt" format:"date-time"`
	ExhaustedAt      time.Time `json:"exhaustedAt" format:"date-time"`
	DrawnAt          time.Time `json:"drawnAt" format:"date-time"`
	LastUsedAt       time.Time `json:"lastUsedAt" format:"date-time"`
	PausedAt         time.Time `json:"pausedAt" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt        respjson.Field
		UpdatedAt        respjson.Field
		CompletedAt      respjson.Field
		ConfirmedAt      respjson.Field
		ExpiresAt        respjson.Field
		OccurredAt       respjson.Field
		DiscoveredAt     respjson.Field
		PostedAt         respjson.Field
		FirstAttemptedAt respjson.Field
		DeliveredAt      respjson.Field
		NextRetryAt      respjson.Field
		ExhaustedAt      respjson.Field
		DrawnAt          respjson.Field
		LastUsedAt       respjson.Field
		PausedAt         respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ActivityRecordTimestamps) RawJSON() string { return r.JSON.raw }
func (r *ActivityRecordTimestamps) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The API request behind this activity. Null for activity from before Oct 30,
// 2026, and for activity no API request starts, such as a monitor event, a webhook
// delivery or a credit grant.
type ActivityRecordRequest struct {
	// Any of "DELETE", "PATCH", "POST", "PUT".
	Method     ActivityRecordRequestMethod `json:"method" api:"required"`
	Path       string                      `json:"path" api:"required"`
	HTTPStatus int64                       `json:"httpStatus" api:"required"`
	DurationMs int64                       `json:"durationMs" api:"required"`
	// IP address of the caller, or null.
	IPAddress string `json:"ipAddress" api:"required"`
	// User agent of the caller, or null.
	UserAgent string    `json:"userAgent" api:"required"`
	At        time.Time `json:"at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Method      respjson.Field
		Path        respjson.Field
		HTTPStatus  respjson.Field
		DurationMs  respjson.Field
		IPAddress   respjson.Field
		UserAgent   respjson.Field
		At          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ActivityRecordRequest) RawJSON() string { return r.JSON.raw }
func (r *ActivityRecordRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActivityRecordRequestMethod string

const (
	ActivityRecordRequestMethodDelete ActivityRecordRequestMethod = "DELETE"
	ActivityRecordRequestMethodPatch  ActivityRecordRequestMethod = "PATCH"
	ActivityRecordRequestMethodPost   ActivityRecordRequestMethod = "POST"
	ActivityRecordRequestMethodPut    ActivityRecordRequestMethod = "PUT"
)

type ActivityRecordCaller struct {
	// How the caller signed in, or null.
	//
	// Any of "api_key", "oauth", "session".
	AuthMethod ActivityRecordCallerAuthMethod `json:"authMethod" api:"required"`
	// Name and prefix of the API key used. Never its value. Null when no API key
	// made the request.
	APIKey ActivityRecordCallerAPIKey `json:"apiKey" api:"required"`
	// The OAuth app that made the request, or null.
	OAuthClient ActivityRecordCallerOAuthClient `json:"oauthClient" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AuthMethod  respjson.Field
		APIKey      respjson.Field
		OAuthClient respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ActivityRecordCaller) RawJSON() string { return r.JSON.raw }
func (r *ActivityRecordCaller) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// How the caller signed in, or null.
type ActivityRecordCallerAuthMethod string

const (
	ActivityRecordCallerAuthMethodAPIKey  ActivityRecordCallerAuthMethod = "api_key"
	ActivityRecordCallerAuthMethodOAuth   ActivityRecordCallerAuthMethod = "oauth"
	ActivityRecordCallerAuthMethodSession ActivityRecordCallerAuthMethod = "session"
)

// Name and prefix of the API key used. Never its value. Null when no API key made
// the request.
type ActivityRecordCallerAPIKey struct {
	Name   string `json:"name" api:"required"`
	Prefix string `json:"prefix" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Prefix      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ActivityRecordCallerAPIKey) RawJSON() string { return r.JSON.raw }
func (r *ActivityRecordCallerAPIKey) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The OAuth app that made the request, or null.
type ActivityRecordCallerOAuthClient struct {
	ID   string `json:"id" api:"required"`
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
func (r ActivityRecordCallerOAuthClient) RawJSON() string { return r.JSON.raw }
func (r *ActivityRecordCallerOAuthClient) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Credits charged for this activity, or null.
type ActivityRecordCharge struct {
	Credits      int64     `json:"credits" api:"required"`
	BalanceAfter int64     `json:"balanceAfter" api:"required"`
	At           time.Time `json:"at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Credits      respjson.Field
		BalanceAfter respjson.Field
		At           respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ActivityRecordCharge) RawJSON() string { return r.JSON.raw }
func (r *ActivityRecordCharge) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Another activity on the same subject.
type ActivityRecordRelated struct {
	ID   string `json:"id" api:"required"`
	Type string `json:"type" api:"required"`
	// Any of "completed", "running", "failed", "canceled".
	Status    ActivityRecordRelatedStatus `json:"status" api:"required"`
	CreatedAt time.Time                   `json:"createdAt" api:"required" format:"date-time"`
	URL       string                      `json:"url" api:"required" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Type        respjson.Field
		Status      respjson.Field
		CreatedAt   respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ActivityRecordRelated) RawJSON() string { return r.JSON.raw }
func (r *ActivityRecordRelated) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActivityRecordRelatedStatus string

const (
	ActivityRecordRelatedStatusCompleted ActivityRecordRelatedStatus = "completed"
	ActivityRecordRelatedStatusRunning   ActivityRecordRelatedStatus = "running"
	ActivityRecordRelatedStatusFailed    ActivityRecordRelatedStatus = "failed"
	ActivityRecordRelatedStatusCanceled  ActivityRecordRelatedStatus = "canceled"
)

type ActivityListParams struct {
	// Most rows to return, 100 by default. PDF stops at 10,000 rows.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Activity group, like the Activity page's tabs. all by default, or actions,
	// extractions, monitors, webhooks or billing.
	Category param.Opt[string] `query:"category,omitzero" json:"-"`
	// Comma-separated statuses.
	Status param.Opt[string] `query:"status,omitzero" json:"-"`
	// Search text.
	Q param.Opt[string] `query:"q,omitzero" json:"-"`
	// Oldest time, ISO 8601.
	Since param.Opt[time.Time] `query:"since,omitzero" format:"date-time" json:"-"`
	// Newest time, ISO 8601.
	Until param.Opt[time.Time] `query:"until,omitzero" format:"date-time" json:"-"`
	paramObj
}

// URLQuery serializes [ActivityListParams]'s query parameters as `url.Values`.
func (r ActivityListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ActivityExportParams struct {
	// File format of the export.
	//
	// Any of "json", "csv", "xlsx", "md", "md-document", "pdf", "txt".
	Format ActivityExportParamsFormat `query:"format,omitzero" api:"required" json:"-"`
	// Most rows to return, 100 by default. PDF stops at 10,000 rows.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Activity group, like the Activity page's tabs. all by default, or actions,
	// extractions, monitors, webhooks or billing.
	Category param.Opt[string] `query:"category,omitzero" json:"-"`
	// Comma-separated statuses.
	Status param.Opt[string] `query:"status,omitzero" json:"-"`
	// Search text.
	Q param.Opt[string] `query:"q,omitzero" json:"-"`
	// Oldest time, ISO 8601.
	Since param.Opt[time.Time] `query:"since,omitzero" format:"date-time" json:"-"`
	// Newest time, ISO 8601.
	Until param.Opt[time.Time] `query:"until,omitzero" format:"date-time" json:"-"`
	paramObj
}

// URLQuery serializes [ActivityExportParams]'s query parameters as `url.Values`.
func (r ActivityExportParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// File format of the export.
type ActivityExportParamsFormat string

const (
	ActivityExportParamsFormatJson       ActivityExportParamsFormat = "json"
	ActivityExportParamsFormatCsv        ActivityExportParamsFormat = "csv"
	ActivityExportParamsFormatXlsx       ActivityExportParamsFormat = "xlsx"
	ActivityExportParamsFormatMd         ActivityExportParamsFormat = "md"
	ActivityExportParamsFormatMdDocument ActivityExportParamsFormat = "md-document"
	ActivityExportParamsFormatPdf        ActivityExportParamsFormat = "pdf"
	ActivityExportParamsFormatTxt        ActivityExportParamsFormat = "txt"
)
