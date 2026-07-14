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

	"github.com/Xquik-dev/x-twitter-scraper-go/internal/apijson"
	"github.com/Xquik-dev/x-twitter-scraper-go/internal/requestconfig"
	"github.com/Xquik-dev/x-twitter-scraper-go/option"
	"github.com/Xquik-dev/x-twitter-scraper-go/packages/respjson"
)

// X write actions (tweets, likes, follows, DMs)
//
// XWriteActionService contains methods and other services that help with
// interacting with the x-twitter-scraper API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewXWriteActionService] method instead.
type XWriteActionService struct {
	options []option.RequestOption
}

// NewXWriteActionService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewXWriteActionService(opts ...option.RequestOption) (r XWriteActionService) {
	r = XWriteActionService{}
	r.options = opts
	return
}

// Get write action status
func (r *XWriteActionService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *XWriteActionGetResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("x/write-actions/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type XWriteActionGetResponse struct {
	Action         string                       `json:"action" api:"required"`
	Charged        bool                         `json:"charged" api:"required"`
	ChargedCredits string                       `json:"chargedCredits" api:"required"`
	CreatedAt      time.Time                    `json:"createdAt" api:"required" format:"date-time"`
	Media          XWriteActionGetResponseMedia `json:"media" api:"required"`
	Retryable      bool                         `json:"retryable" api:"required"`
	SendDispatched bool                         `json:"sendDispatched" api:"required"`
	// Any of "success", "failed", "pending_confirmation".
	Status                XWriteActionGetResponseStatus `json:"status" api:"required"`
	WriteActionID         string                        `json:"writeActionId" api:"required"`
	ConfirmationAttempts  int64                         `json:"confirmationAttempts"`
	ConfirmationCheckedAt time.Time                     `json:"confirmationCheckedAt" format:"date-time"`
	ConfirmationSource    string                        `json:"confirmationSource" api:"nullable"`
	ConfirmedAt           time.Time                     `json:"confirmedAt" format:"date-time"`
	Message               string                        `json:"message"`
	MessageID             string                        `json:"messageId"`
	SendDispatchedAt      time.Time                     `json:"sendDispatchedAt" format:"date-time"`
	TargetID              string                        `json:"targetId" api:"nullable"`
	TweetID               string                        `json:"tweetId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Action                respjson.Field
		Charged               respjson.Field
		ChargedCredits        respjson.Field
		CreatedAt             respjson.Field
		Media                 respjson.Field
		Retryable             respjson.Field
		SendDispatched        respjson.Field
		Status                respjson.Field
		WriteActionID         respjson.Field
		ConfirmationAttempts  respjson.Field
		ConfirmationCheckedAt respjson.Field
		ConfirmationSource    respjson.Field
		ConfirmedAt           respjson.Field
		Message               respjson.Field
		MessageID             respjson.Field
		SendDispatchedAt      respjson.Field
		TargetID              respjson.Field
		TweetID               respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XWriteActionGetResponse) RawJSON() string { return r.JSON.raw }
func (r *XWriteActionGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type XWriteActionGetResponseMedia struct {
	Count   int64  `json:"count" api:"required"`
	Credits string `json:"credits" api:"required"`
	// Any of "none", "image", "video".
	Kind       string `json:"kind" api:"required"`
	TotalBytes string `json:"totalBytes" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Credits     respjson.Field
		Kind        respjson.Field
		TotalBytes  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XWriteActionGetResponseMedia) RawJSON() string { return r.JSON.raw }
func (r *XWriteActionGetResponseMedia) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type XWriteActionGetResponseStatus string

const (
	XWriteActionGetResponseStatusSuccess             XWriteActionGetResponseStatus = "success"
	XWriteActionGetResponseStatusFailed              XWriteActionGetResponseStatus = "failed"
	XWriteActionGetResponseStatusPendingConfirmation XWriteActionGetResponseStatus = "pending_confirmation"
)
