// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

package xtwitterscraper

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/Xquik-dev/x-twitter-scraper-go/internal/apiform"
	"github.com/Xquik-dev/x-twitter-scraper-go/internal/apijson"
	"github.com/Xquik-dev/x-twitter-scraper-go/internal/requestconfig"
	"github.com/Xquik-dev/x-twitter-scraper-go/option"
	"github.com/Xquik-dev/x-twitter-scraper-go/packages/param"
	"github.com/Xquik-dev/x-twitter-scraper-go/packages/respjson"
)

// Support ticket management
//
// SupportTicketService contains methods and other services that help with
// interacting with the x-twitter-scraper API.
//
// You should not instantiate this service directly, and instead use the
// [NewSupportTicketService] method instead.
type SupportTicketService struct {
	options []option.RequestOption
}

// NewSupportTicketService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSupportTicketService(opts ...option.RequestOption) (r SupportTicketService) {
	r = SupportTicketService{}
	r.options = opts
	return
}

// Creates a support ticket with an initial message.
func (r *SupportTicketService) New(ctx context.Context, params SupportTicketNewParams, opts ...option.RequestOption) (res *SupportTicketNewResponse, err error) {
	if err = validateTicketMessage(params.Body, params.Attachments); err != nil {
		return nil, err
	}
	if !param.IsOmitted(params.IdempotencyKey) {
		opts = append(opts, option.WithHeader("Idempotency-Key", fmt.Sprintf("%v", params.IdempotencyKey.Value)))
	}
	opts = slices.Concat(r.options, opts)
	path := "support/tickets"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Returns one support ticket with its message history.
func (r *SupportTicketService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *SupportTicketGetResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("support/tickets/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Changes the open or closed state of a support ticket.
func (r *SupportTicketService) Update(ctx context.Context, id string, body SupportTicketUpdateParams, opts ...option.RequestOption) (res *SupportTicketUpdateResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("support/tickets/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Returns support tickets owned by the authenticated user.
func (r *SupportTicketService) List(ctx context.Context, opts ...option.RequestOption) (res *SupportTicketListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "support/tickets"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Adds a message to an existing support ticket.
func (r *SupportTicketService) Reply(ctx context.Context, id string, params SupportTicketReplyParams, opts ...option.RequestOption) (res *SupportTicketReplyResponse, err error) {
	if !param.IsOmitted(params.IdempotencyKey) {
		opts = append(opts, option.WithHeader("Idempotency-Key", fmt.Sprintf("%v", params.IdempotencyKey.Value)))
	}
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	if err = validateTicketMessage(params.Body, params.Attachments); err != nil {
		return nil, err
	}
	path := fmt.Sprintf("support/tickets/%s/messages", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

type SupportTicketNewResponse struct {
	Attachments []SupportTicketNewResponseAttachment `json:"attachments" api:"required"`
	PublicID    string                               `json:"publicId" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Attachments respjson.Field
		PublicID    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SupportTicketNewResponse) RawJSON() string { return r.JSON.raw }
func (r *SupportTicketNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// Attachment identifier and initial processing state.
type SupportTicketNewResponseAttachment struct {
	PublicID string `json:"publicId" api:"required"`
	// Any of "pending", "ready", "failed".
	Status string `json:"status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PublicID    respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SupportTicketNewResponseAttachment) RawJSON() string { return r.JSON.raw }
func (r *SupportTicketNewResponseAttachment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type SupportTicketGetResponse struct {
	CreatedAt time.Time                         `json:"createdAt" api:"required" format:"date-time"`
	Messages  []SupportTicketGetResponseMessage `json:"messages" api:"required"`
	PublicID  string                            `json:"publicId" api:"required"`
	// Any of "open", "in_progress", "resolved", "closed".
	Status    SupportTicketGetResponseStatus `json:"status" api:"required"`
	Subject   string                         `json:"subject" api:"required"`
	UpdatedAt time.Time                      `json:"updatedAt" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt   respjson.Field
		Messages    respjson.Field
		PublicID    respjson.Field
		Status      respjson.Field
		Subject     respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SupportTicketGetResponse) RawJSON() string { return r.JSON.raw }
func (r *SupportTicketGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type SupportTicketGetResponseMessage struct {
	Attachments []SupportTicketGetResponseMessageAttachment `json:"attachments" api:"required"`
	Body        string                                      `json:"body" api:"required"`
	CreatedAt   time.Time                                   `json:"createdAt" api:"required" format:"date-time"`
	// Any of "user", "support", "system".
	Sender string `json:"sender" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Attachments respjson.Field
		Body        respjson.Field
		CreatedAt   respjson.Field
		Sender      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SupportTicketGetResponseMessage) RawJSON() string { return r.JSON.raw }
func (r *SupportTicketGetResponseMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// Downloadable image or video attached to a support message.
type SupportTicketGetResponseMessageAttachment struct {
	// Validated media type.
	//
	// Any of "image/jpeg", "image/png", "image/gif", "image/webp", "video/mp4",
	// "video/quicktime", "video/webm".
	ContentType string `json:"contentType" api:"required"`
	Filename    string `json:"filename" api:"required"`
	// Attachment media class.
	//
	// Any of "image", "video".
	Kind      string `json:"kind" api:"required"`
	PublicID  string `json:"publicId" api:"required"`
	SizeBytes int64  `json:"sizeBytes" api:"required"`
	// Storage processing state.
	//
	// Any of "pending", "ready", "failed".
	Status string `json:"status" api:"required"`
	URL    string `json:"url" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ContentType respjson.Field
		Filename    respjson.Field
		Kind        respjson.Field
		PublicID    respjson.Field
		SizeBytes   respjson.Field
		Status      respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SupportTicketGetResponseMessageAttachment) RawJSON() string { return r.JSON.raw }
func (r *SupportTicketGetResponseMessageAttachment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type SupportTicketGetResponseStatus string

const (
	SupportTicketGetResponseStatusOpen       SupportTicketGetResponseStatus = "open"
	SupportTicketGetResponseStatusInProgress SupportTicketGetResponseStatus = "in_progress"
	SupportTicketGetResponseStatusResolved   SupportTicketGetResponseStatus = "resolved"
	SupportTicketGetResponseStatusClosed     SupportTicketGetResponseStatus = "closed"
)

type SupportTicketUpdateResponse struct {
	PublicID string `json:"publicId" api:"required"`
	// Any of "open", "resolved", "closed".
	Status SupportTicketUpdateResponseStatus `json:"status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PublicID    respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SupportTicketUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *SupportTicketUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type SupportTicketUpdateResponseStatus string

const (
	SupportTicketUpdateResponseStatusOpen     SupportTicketUpdateResponseStatus = "open"
	SupportTicketUpdateResponseStatusResolved SupportTicketUpdateResponseStatus = "resolved"
	SupportTicketUpdateResponseStatusClosed   SupportTicketUpdateResponseStatus = "closed"
)

type SupportTicketListResponse struct {
	HasMore    bool                              `json:"hasMore" api:"required"`
	Tickets    []SupportTicketListResponseTicket `json:"tickets" api:"required"`
	NextCursor string                            `json:"nextCursor"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		HasMore     respjson.Field
		Tickets     respjson.Field
		NextCursor  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SupportTicketListResponse) RawJSON() string { return r.JSON.raw }
func (r *SupportTicketListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type SupportTicketListResponseTicket struct {
	CreatedAt    time.Time `json:"createdAt" api:"required" format:"date-time"`
	MessageCount int64     `json:"messageCount" api:"required"`
	PublicID     string    `json:"publicId" api:"required"`
	// Any of "open", "in_progress", "resolved", "closed".
	Status    string    `json:"status" api:"required"`
	Subject   string    `json:"subject" api:"required"`
	UpdatedAt time.Time `json:"updatedAt" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt    respjson.Field
		MessageCount respjson.Field
		PublicID     respjson.Field
		Status       respjson.Field
		Subject      respjson.Field
		UpdatedAt    respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SupportTicketListResponseTicket) RawJSON() string { return r.JSON.raw }
func (r *SupportTicketListResponseTicket) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type SupportTicketReplyResponse struct {
	Attachments []SupportTicketReplyResponseAttachment `json:"attachments" api:"required"`
	PublicID    string                                 `json:"publicId" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Attachments respjson.Field
		PublicID    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SupportTicketReplyResponse) RawJSON() string { return r.JSON.raw }
func (r *SupportTicketReplyResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// Attachment identifier and initial processing state.
type SupportTicketReplyResponseAttachment struct {
	PublicID string `json:"publicId" api:"required"`
	// Any of "pending", "ready", "failed".
	Status string `json:"status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PublicID    respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SupportTicketReplyResponseAttachment) RawJSON() string { return r.JSON.raw }
func (r *SupportTicketReplyResponseAttachment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

var ErrTicketNeedsBodyOrAttachments = errors.New("missing required body or attachments parameter")

func validateTicketMessage(body param.Opt[string], attachments []io.Reader) error {
	if body.Or("") == "" && len(attachments) == 0 {
		return ErrTicketNeedsBodyOrAttachments
	}
	return nil
}

type SupportTicketNewParams struct {
	// Non-empty support ticket subject.
	Subject string `json:"subject" api:"required"`
	// Up to 4 image or video attachments. Each image can be up to 10 MiB and each
	// video up to 25 MiB. The request can be up to 30 MiB.
	Attachments []io.Reader `json:"attachments,omitzero" format:"binary"`
	// Non-empty support message text. Optional when attachments are provided.
	Body           param.Opt[string] `json:"body,omitzero"`
	IdempotencyKey param.Opt[string] `header:"Idempotency-Key,omitzero" json:"-"`
	paramObj
}

func (r SupportTicketNewParams) MarshalJSON() (data []byte, err error) {
	type shadow SupportTicketNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SupportTicketNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, nil)
}

func (r SupportTicketNewParams) HasFiles() bool {
	return len(r.Attachments) > 0
}

func (r SupportTicketNewParams) MarshalMultipart() (data []byte, contentType string, err error) {
	return apiform.MarshalMultipart(r)
}

type SupportTicketUpdateParams struct {
	// Any of "open", "resolved", "closed".
	Status SupportTicketUpdateParamsStatus `json:"status,omitzero" api:"required"`
	paramObj
}

func (r SupportTicketUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow SupportTicketUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SupportTicketUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, nil)
}

type SupportTicketUpdateParamsStatus string

const (
	SupportTicketUpdateParamsStatusOpen     SupportTicketUpdateParamsStatus = "open"
	SupportTicketUpdateParamsStatusResolved SupportTicketUpdateParamsStatus = "resolved"
	SupportTicketUpdateParamsStatusClosed   SupportTicketUpdateParamsStatus = "closed"
)

type SupportTicketReplyParams struct {
	// Up to 4 image or video attachments. Each image can be up to 10 MiB and each
	// video up to 25 MiB. The request can be up to 30 MiB.
	Attachments []io.Reader `json:"attachments,omitzero" format:"binary"`
	// Non-empty support message text. Optional when attachments are provided.
	Body           param.Opt[string] `json:"body,omitzero"`
	IdempotencyKey param.Opt[string] `header:"Idempotency-Key,omitzero" json:"-"`
	paramObj
}

func (r SupportTicketReplyParams) MarshalJSON() (data []byte, err error) {
	type shadow SupportTicketReplyParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SupportTicketReplyParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, nil)
}

func (r SupportTicketReplyParams) HasFiles() bool {
	return len(r.Attachments) > 0
}

func (r SupportTicketReplyParams) MarshalMultipart() (data []byte, contentType string, err error) {
	return apiform.MarshalMultipart(r)
}
