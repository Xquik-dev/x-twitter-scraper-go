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
	"github.com/stainless-sdks/x-twitter-scraper-go/internal/requestconfig"
	"github.com/stainless-sdks/x-twitter-scraper-go/option"
	"github.com/stainless-sdks/x-twitter-scraper-go/packages/param"
	"github.com/stainless-sdks/x-twitter-scraper-go/packages/respjson"
)

// Support ticket management
//
// SupportTicketService contains methods and other services that help with
// interacting with the x-twitter-scraper API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSupportTicketService] method instead.
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

// Create a support ticket
func (r *SupportTicketService) New(ctx context.Context, body SupportTicketNewParams, opts ...option.RequestOption) (res *SupportTicketNewResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "support/tickets"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Get ticket with all messages
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

// Update ticket status
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

// List user's support tickets
func (r *SupportTicketService) List(ctx context.Context, opts ...option.RequestOption) (res *SupportTicketListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "support/tickets"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Reply to a support ticket
func (r *SupportTicketService) Reply(ctx context.Context, id string, body SupportTicketReplyParams, opts ...option.RequestOption) (res *SupportTicketReplyResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("support/tickets/%s/messages", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type SupportTicketNewResponse struct {
	PublicID string `json:"publicId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PublicID    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SupportTicketNewResponse) RawJSON() string { return r.JSON.raw }
func (r *SupportTicketNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SupportTicketGetResponse struct {
	CreatedAt time.Time                         `json:"createdAt" format:"date-time"`
	Messages  []SupportTicketGetResponseMessage `json:"messages"`
	PublicID  string                            `json:"publicId"`
	Status    string                            `json:"status"`
	Subject   string                            `json:"subject"`
	UpdatedAt time.Time                         `json:"updatedAt" format:"date-time"`
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
	return apijson.UnmarshalRoot(data, r)
}

type SupportTicketGetResponseMessage struct {
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	Sender    string    `json:"sender"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
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
	return apijson.UnmarshalRoot(data, r)
}

type SupportTicketUpdateResponse struct {
	PublicID string `json:"publicId"`
	Status   string `json:"status"`
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
	return apijson.UnmarshalRoot(data, r)
}

type SupportTicketListResponse struct {
	Tickets []SupportTicketListResponseTicket `json:"tickets"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Tickets     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SupportTicketListResponse) RawJSON() string { return r.JSON.raw }
func (r *SupportTicketListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SupportTicketListResponseTicket struct {
	CreatedAt    time.Time `json:"createdAt" format:"date-time"`
	MessageCount int64     `json:"messageCount"`
	PublicID     string    `json:"publicId"`
	Status       string    `json:"status"`
	Subject      string    `json:"subject"`
	UpdatedAt    time.Time `json:"updatedAt" format:"date-time"`
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
	return apijson.UnmarshalRoot(data, r)
}

type SupportTicketReplyResponse struct {
	PublicID string `json:"publicId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PublicID    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SupportTicketReplyResponse) RawJSON() string { return r.JSON.raw }
func (r *SupportTicketReplyResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SupportTicketNewParams struct {
	Body    string `json:"body" api:"required"`
	Subject string `json:"subject" api:"required"`
	paramObj
}

func (r SupportTicketNewParams) MarshalJSON() (data []byte, err error) {
	type shadow SupportTicketNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SupportTicketNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
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
	return apijson.UnmarshalRoot(data, r)
}

type SupportTicketUpdateParamsStatus string

const (
	SupportTicketUpdateParamsStatusOpen     SupportTicketUpdateParamsStatus = "open"
	SupportTicketUpdateParamsStatusResolved SupportTicketUpdateParamsStatus = "resolved"
	SupportTicketUpdateParamsStatusClosed   SupportTicketUpdateParamsStatus = "closed"
)

type SupportTicketReplyParams struct {
	Body string `json:"body" api:"required"`
	paramObj
}

func (r SupportTicketReplyParams) MarshalJSON() (data []byte, err error) {
	type shadow SupportTicketReplyParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SupportTicketReplyParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
