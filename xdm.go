// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package xtwitterscraper

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/stainless-sdks/x-twitter-scraper-go/internal/apijson"
	"github.com/stainless-sdks/x-twitter-scraper-go/internal/apiquery"
	"github.com/stainless-sdks/x-twitter-scraper-go/internal/requestconfig"
	"github.com/stainless-sdks/x-twitter-scraper-go/option"
	"github.com/stainless-sdks/x-twitter-scraper-go/packages/param"
	"github.com/stainless-sdks/x-twitter-scraper-go/packages/respjson"
)

// XDmService contains methods and other services that help with interacting with
// the x-twitter-scraper API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewXDmService] method instead.
type XDmService struct {
	options []option.RequestOption
}

// NewXDmService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewXDmService(opts ...option.RequestOption) (r XDmService) {
	r = XDmService{}
	r.options = opts
	return
}

// Get DM conversation history
func (r *XDmService) GetHistory(ctx context.Context, userID string, query XDmGetHistoryParams, opts ...option.RequestOption) (res *XDmGetHistoryResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if userID == "" {
		err = errors.New("missing required userId parameter")
		return nil, err
	}
	path := fmt.Sprintf("x/dm/%s/history", url.PathEscape(userID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Send direct message
func (r *XDmService) Send(ctx context.Context, userID string, body XDmSendParams, opts ...option.RequestOption) (res *XDmSendResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if userID == "" {
		err = errors.New("missing required userId parameter")
		return nil, err
	}
	path := fmt.Sprintf("x/dm/%s", url.PathEscape(userID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type XDmGetHistoryResponse struct {
	HasNextPage bool                           `json:"has_next_page" api:"required"`
	Messages    []XDmGetHistoryResponseMessage `json:"messages" api:"required"`
	NextCursor  string                         `json:"next_cursor" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		HasNextPage respjson.Field
		Messages    respjson.Field
		NextCursor  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XDmGetHistoryResponse) RawJSON() string { return r.JSON.raw }
func (r *XDmGetHistoryResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type XDmGetHistoryResponseMessage struct {
	ID         string `json:"id" api:"required"`
	CreatedAt  string `json:"createdAt"`
	ReceiverID string `json:"receiverId"`
	SenderID   string `json:"senderId"`
	Text       string `json:"text"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		ReceiverID  respjson.Field
		SenderID    respjson.Field
		Text        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XDmGetHistoryResponseMessage) RawJSON() string { return r.JSON.raw }
func (r *XDmGetHistoryResponseMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type XDmSendResponse struct {
	MessageID string `json:"messageId" api:"required"`
	Success   bool   `json:"success" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MessageID   respjson.Field
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XDmSendResponse) RawJSON() string { return r.JSON.raw }
func (r *XDmSendResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type XDmGetHistoryParams struct {
	// Pagination cursor for DM history
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Legacy pagination cursor (backward compat)
	MaxID param.Opt[string] `query:"maxId,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [XDmGetHistoryParams]'s query parameters as `url.Values`.
func (r XDmGetHistoryParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type XDmSendParams struct {
	// X account (@username or ID) sending the DM
	Account          string            `json:"account" api:"required"`
	Text             string            `json:"text" api:"required"`
	ReplyToMessageID param.Opt[string] `json:"reply_to_message_id,omitzero"`
	MediaIDs         []string          `json:"media_ids,omitzero"`
	paramObj
}

func (r XDmSendParams) MarshalJSON() (data []byte, err error) {
	type shadow XDmSendParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *XDmSendParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
