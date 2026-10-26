// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

package xtwitterscraper

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"path"
	"slices"
	"time"

	"github.com/Xquik-dev/x-twitter-scraper-go/internal/apiform"
	"github.com/Xquik-dev/x-twitter-scraper-go/internal/apijson"
	"github.com/Xquik-dev/x-twitter-scraper-go/internal/requestconfig"
	"github.com/Xquik-dev/x-twitter-scraper-go/option"
	"github.com/Xquik-dev/x-twitter-scraper-go/packages/param"
	"github.com/Xquik-dev/x-twitter-scraper-go/packages/respjson"
	"github.com/Xquik-dev/x-twitter-scraper-go/shared/constant"
)

// XMediaService contains methods and other services that help with interacting
// with the x-twitter-scraper API.
//
// You should not instantiate this service directly, and instead use the
// [NewXMediaService] method instead.
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

// Download downloads images and videos attached to selected tweets.
func (r *XMediaService) Download(ctx context.Context, body XMediaDownloadParams, opts ...option.RequestOption) (res *XMediaDownloadResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "x/media/download"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Upload uploads media and returns an identifier for later writes.
func (r *XMediaService) Upload(ctx context.Context, params XMediaUploadParams, opts ...option.RequestOption) (res *XMediaUploadResponse, err error) {
	if err = validateMediaUpload(params.URL, params.File); err != nil {
		return nil, err
	}
	if !param.IsOmitted(params.IsLongVideo) && params.File == nil {
		return nil, ErrLongVideoNeedsFile
	}
	if !param.IsOmitted(params.IdempotencyKey) {
		opts = append(opts, option.WithHeader("Idempotency-Key", fmt.Sprintf("%v", params.IdempotencyKey)))
	}
	opts = slices.Concat(r.options, opts)
	path := "x/media"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

type XMediaDownloadResponse struct {
	CacheHit   bool   `json:"cacheHit"`
	GalleryURL string `json:"galleryUrl"`
	// The post's files saved to your gallery. Single downloads only.
	Items       []XMediaDownloadResponseItem `json:"items"`
	TotalMedia  int64                        `json:"totalMedia"`
	TotalTweets int64                        `json:"totalTweets"`
	TweetID     string                       `json:"tweetId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CacheHit    respjson.Field
		GalleryURL  respjson.Field
		Items       respjson.Field
		TotalMedia  respjson.Field
		TotalTweets respjson.Field
		TweetID     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// RawJSON returns the unmodified JSON received from the API.
func (r XMediaDownloadResponse) RawJSON() string { return r.JSON.raw }
func (r *XMediaDownloadResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// XMediaDownloadResponseItem holds 1 file of the post saved to your gallery.
type XMediaDownloadResponseItem struct {
	// The gallery file's ID.
	ID      string `json:"id" api:"required"`
	TweetID string `json:"tweetId" api:"required"`
	// Any of "photo", "video", "animated_gif".
	MediaType XMediaDownloadResponseItemMediaType `json:"mediaType" api:"required"`
	// The saved copy of the file.
	MediaURL string `json:"mediaUrl" api:"required"`
	// File size in bytes, as a string
	FileSize string `json:"fileSize" api:"required"`
	// When the file joined your gallery
	CreatedAt         time.Time `json:"createdAt" api:"required" format:"date-time"`
	AuthorUsername    string    `json:"authorUsername" api:"required"`
	AuthorDisplayName string    `json:"authorDisplayName" api:"required"`
	AuthorAvatar      string    `json:"authorAvatar" api:"required"`
	TweetText         string    `json:"tweetText" api:"required"`
	Likes             int64     `json:"likes" api:"required"`
	Retweets          int64     `json:"retweets" api:"required"`
	Views             int64     `json:"views" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		TweetID           respjson.Field
		MediaType         respjson.Field
		MediaURL          respjson.Field
		FileSize          respjson.Field
		CreatedAt         respjson.Field
		AuthorUsername    respjson.Field
		AuthorDisplayName respjson.Field
		AuthorAvatar      respjson.Field
		TweetText         respjson.Field
		Likes             respjson.Field
		Retweets          respjson.Field
		Views             respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// RawJSON returns the unmodified JSON received from the API.
func (r XMediaDownloadResponseItem) RawJSON() string { return r.JSON.raw }
func (r *XMediaDownloadResponseItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// XMediaDownloadResponseItemMediaType is the kind of saved file.
type XMediaDownloadResponseItemMediaType string

const (
	XMediaDownloadResponseItemMediaTypePhoto       XMediaDownloadResponseItemMediaType = "photo"
	XMediaDownloadResponseItemMediaTypeVideo       XMediaDownloadResponseItemMediaType = "video"
	XMediaDownloadResponseItemMediaTypeAnimatedGif XMediaDownloadResponseItemMediaType = "animated_gif"
)

// XMediaUploadResponse is a durable write record. Poll statusUrl until terminal is true. Reusing its
// Idempotency-Key returns this record. Create another action only when safeToRetry
// is true.
type XMediaUploadResponse struct {
	ID string `json:"id" api:"required"`
	// Connected account selected for the write.
	Account XMediaUploadResponseAccount `json:"account" api:"required"`
	// Any of "create_tweet", "delete_tweet", "like", "unlike", "retweet", "unretweet",
	// "follow", "unfollow", "remove_follower", "send_dm", "delete_dm", "upload_media",
	// "update_profile", "update_avatar", "update_banner", "create_community",
	// "delete_community", "join_community", "leave_community".
	Action XMediaUploadResponseAction `json:"action" api:"required"`
	// plannedCredits is the approved maximum. chargedCredits comes from the settled
	// credit ledger. Pending or failed writes are not charged.
	Billing        XMediaUploadResponseBilling `json:"billing" api:"required"`
	Charged        bool                        `json:"charged" api:"required"`
	ChargedCredits string                      `json:"chargedCredits" api:"required"`
	// Exact follow-up an API client or agent should perform.
	NextAction  XMediaUploadResponseNextAction `json:"nextAction" api:"required"`
	Object      constant.XWriteAction          `json:"object" default:"x_write_action"`
	PollAfterMs int64                          `json:"pollAfterMs" api:"required"`
	// Stable fingerprint and sanitized payload for replay checks.
	Request XMediaUploadResponseRequest `json:"request" api:"required"`
	// Confirmed result produced by the write, when available.
	Result XMediaUploadResponseResult `json:"result" api:"required"`
	// True only when a new attempt can reasonably succeed.
	Retryable bool `json:"retryable" api:"required"`
	// True only when no write was dispatched and a new idempotency key may be used.
	SafeToRetry    bool `json:"safeToRetry" api:"required"`
	SendDispatched bool `json:"sendDispatched" api:"required"`
	// Any of "accepted", "dispatching", "pending_confirmation", "success", "failed",
	// "expired".
	Status    XMediaUploadResponseStatus `json:"status" api:"required"`
	StatusURL string                     `json:"statusUrl" api:"required"`
	Success   bool                       `json:"success" api:"required"`
	// Existing X resource targeted by the write, when applicable.
	Target        XMediaUploadResponseTarget `json:"target" api:"required"`
	TargetID      string                     `json:"targetId" api:"required"`
	Terminal      bool                       `json:"terminal" api:"required"`
	WriteActionID string                     `json:"writeActionId" api:"required"`
	// Compatibility field for a confirmed community ID.
	CommunityID string `json:"communityId"`
	// Confirmed community name when available.
	CommunityName         string    `json:"communityName"`
	CompletedAt           time.Time `json:"completedAt" format:"date-time"`
	ConfirmationAttempts  int64     `json:"confirmationAttempts"`
	ConfirmationCheckedAt time.Time `json:"confirmationCheckedAt" format:"date-time"`
	ConfirmedAt           time.Time `json:"confirmedAt" format:"date-time"`
	CreatedAt             time.Time `json:"createdAt" format:"date-time"`
	// Structured recovery context for a failed write.
	Details PublicMetadata `json:"details"`
	Error   string         `json:"error"`
	// Deadline for resolving a non-terminal write. This is not the Idempotency-Key
	// retention deadline.
	ExpiresAt  time.Time `json:"expiresAt" format:"date-time"`
	Idempotent bool      `json:"idempotent"`
	// Media count, kind, size, and billing details when used.
	Media PublicMetadata `json:"media"`
	// Compatibility field for a confirmed media upload ID.
	MediaID string `json:"mediaId"`
	// Public media URL when the upload creates one.
	MediaURL string `json:"mediaUrl" format:"uri"`
	Message  string `json:"message"`
	// Compatibility field for a confirmed direct message ID.
	MessageID   string `json:"messageId"`
	RequestHash string `json:"requestHash"`
	RequestID   string `json:"requestId"`
	// Compatibility result ID for other write actions.
	ResultID string `json:"resultId"`
	// Dispatch timestamp when the write reached execution.
	SendDispatchedAt time.Time `json:"sendDispatchedAt" format:"date-time"`
	// Compatibility field for a confirmed tweet result ID.
	TweetID   string    `json:"tweetId"`
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                    respjson.Field
		Account               respjson.Field
		Action                respjson.Field
		Billing               respjson.Field
		Charged               respjson.Field
		ChargedCredits        respjson.Field
		NextAction            respjson.Field
		Object                respjson.Field
		PollAfterMs           respjson.Field
		Request               respjson.Field
		Result                respjson.Field
		Retryable             respjson.Field
		SafeToRetry           respjson.Field
		SendDispatched        respjson.Field
		Status                respjson.Field
		StatusURL             respjson.Field
		Success               respjson.Field
		Target                respjson.Field
		TargetID              respjson.Field
		Terminal              respjson.Field
		WriteActionID         respjson.Field
		CommunityID           respjson.Field
		CommunityName         respjson.Field
		CompletedAt           respjson.Field
		ConfirmationAttempts  respjson.Field
		ConfirmationCheckedAt respjson.Field
		ConfirmedAt           respjson.Field
		CreatedAt             respjson.Field
		Details               respjson.Field
		Error                 respjson.Field
		ExpiresAt             respjson.Field
		Idempotent            respjson.Field
		Media                 respjson.Field
		MediaID               respjson.Field
		MediaURL              respjson.Field
		Message               respjson.Field
		MessageID             respjson.Field
		RequestHash           respjson.Field
		RequestID             respjson.Field
		ResultID              respjson.Field
		SendDispatchedAt      respjson.Field
		TweetID               respjson.Field
		UpdatedAt             respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// RawJSON returns the unmodified JSON received from the API.
func (r XMediaUploadResponse) RawJSON() string { return r.JSON.raw }
func (r *XMediaUploadResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// XMediaUploadResponseAccount is a connected account selected for the write.
type XMediaUploadResponseAccount struct {
	ID       string `json:"id" api:"required"`
	Username string `json:"username" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Username    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// RawJSON returns the unmodified JSON received from the API.
func (r XMediaUploadResponseAccount) RawJSON() string { return r.JSON.raw }
func (r *XMediaUploadResponseAccount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type XMediaUploadResponseAction string

const (
	XMediaUploadResponseActionCreateTweet     XMediaUploadResponseAction = "create_tweet"
	XMediaUploadResponseActionDeleteTweet     XMediaUploadResponseAction = "delete_tweet"
	XMediaUploadResponseActionLike            XMediaUploadResponseAction = "like"
	XMediaUploadResponseActionUnlike          XMediaUploadResponseAction = "unlike"
	XMediaUploadResponseActionRetweet         XMediaUploadResponseAction = "retweet"
	XMediaUploadResponseActionUnretweet       XMediaUploadResponseAction = "unretweet"
	XMediaUploadResponseActionFollow          XMediaUploadResponseAction = "follow"
	XMediaUploadResponseActionUnfollow        XMediaUploadResponseAction = "unfollow"
	XMediaUploadResponseActionRemoveFollower  XMediaUploadResponseAction = "remove_follower"
	XMediaUploadResponseActionSendDm          XMediaUploadResponseAction = "send_dm"
	XMediaUploadResponseActionDeleteDm        XMediaUploadResponseAction = "delete_dm"
	XMediaUploadResponseActionUploadMedia     XMediaUploadResponseAction = "upload_media"
	XMediaUploadResponseActionUpdateProfile   XMediaUploadResponseAction = "update_profile"
	XMediaUploadResponseActionUpdateAvatar    XMediaUploadResponseAction = "update_avatar"
	XMediaUploadResponseActionUpdateBanner    XMediaUploadResponseAction = "update_banner"
	XMediaUploadResponseActionCreateCommunity XMediaUploadResponseAction = "create_community"
	XMediaUploadResponseActionDeleteCommunity XMediaUploadResponseAction = "delete_community"
	XMediaUploadResponseActionJoinCommunity   XMediaUploadResponseAction = "join_community"
	XMediaUploadResponseActionLeaveCommunity  XMediaUploadResponseAction = "leave_community"
)

// XMediaUploadResponseBilling is the type of [XMediaUploadResponse.Billing]. plannedCredits is the approved maximum. chargedCredits comes from the settled
// credit ledger. Pending or failed writes are not charged.
type XMediaUploadResponseBilling struct {
	Charged        bool   `json:"charged" api:"required"`
	ChargedCredits string `json:"chargedCredits" api:"required"`
	PlannedCredits string `json:"plannedCredits" api:"required"`
	// Any of "not_charged", "pending", "charged", "charge_failed", "refunded".
	Status string `json:"status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Charged        respjson.Field
		ChargedCredits respjson.Field
		PlannedCredits respjson.Field
		Status         respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// RawJSON returns the unmodified JSON received from the API.
func (r XMediaUploadResponseBilling) RawJSON() string { return r.JSON.raw }
func (r *XMediaUploadResponseBilling) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// XMediaUploadResponseNextAction is the type of [XMediaUploadResponse.NextAction]. Exact follow-up an API client or agent should perform.
type XMediaUploadResponseNextAction struct {
	// Any of "poll", "retry", "verify_result", "fix_request".
	Type                      string `json:"type" api:"required"`
	AfterMs                   int64  `json:"afterMs"`
	RequiresNewIdempotencyKey bool   `json:"requiresNewIdempotencyKey"`
	URL                       string `json:"url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type                      respjson.Field
		AfterMs                   respjson.Field
		RequiresNewIdempotencyKey respjson.Field
		URL                       respjson.Field
		ExtraFields               map[string]respjson.Field
		raw                       string
	} `json:"-"`
}

// RawJSON returns the unmodified JSON received from the API.
func (r XMediaUploadResponseNextAction) RawJSON() string { return r.JSON.raw }
func (r *XMediaUploadResponseNextAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// XMediaUploadResponseRequest holds stable fingerprint and sanitized payload for replay checks.
type XMediaUploadResponseRequest struct {
	// Stable hash of account, action, target, and payload.
	Hash string `json:"hash" api:"required"`
	// Exact sanitized payload dispatched for this action.
	Payload PublicMetadata `json:"payload" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Hash        respjson.Field
		Payload     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// RawJSON returns the unmodified JSON received from the API.
func (r XMediaUploadResponseRequest) RawJSON() string { return r.JSON.raw }
func (r *XMediaUploadResponseRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// XMediaUploadResponseResult is a confirmed result produced by the write, when available.
type XMediaUploadResponseResult struct {
	ID    string `json:"id"`
	State string `json:"state"`
	// Any of "tweet", "direct_message", "media", "community", "state_change".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		State       respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// RawJSON returns the unmodified JSON received from the API.
func (r XMediaUploadResponseResult) RawJSON() string { return r.JSON.raw }
func (r *XMediaUploadResponseResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type XMediaUploadResponseStatus string

const (
	XMediaUploadResponseStatusAccepted            XMediaUploadResponseStatus = "accepted"
	XMediaUploadResponseStatusDispatching         XMediaUploadResponseStatus = "dispatching"
	XMediaUploadResponseStatusPendingConfirmation XMediaUploadResponseStatus = "pending_confirmation"
	XMediaUploadResponseStatusSuccess             XMediaUploadResponseStatus = "success"
	XMediaUploadResponseStatusFailed              XMediaUploadResponseStatus = "failed"
	XMediaUploadResponseStatusExpired             XMediaUploadResponseStatus = "expired"
)

// XMediaUploadResponseTarget is an existing X resource targeted by the write, when applicable.
type XMediaUploadResponseTarget struct {
	ID string `json:"id" api:"required"`
	// Any of "tweet", "user", "community".
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// RawJSON returns the unmodified JSON received from the API.
func (r XMediaUploadResponseTarget) RawJSON() string { return r.JSON.raw }
func (r *XMediaUploadResponseTarget) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
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
	return apijson.UnmarshalRoot(data, r, nil)
}

var (
	ErrMediaNeedsURLOrFile = errors.New("missing required url or file parameter")
	ErrMediaHasURLAndFile  = errors.New("cannot use url and file parameters together")
	ErrLongVideoNeedsFile  = errors.New("is_long_video parameter requires a file parameter")
)

type XMediaUploadParams struct {
	// X account (@username or ID) uploading media
	Account string `json:"account" api:"required"`
	// Media file to upload, up to 15 MB. Required when is_long_video is provided.
	File io.Reader `json:"file,omitzero" format:"binary"`
	// Set true only for video/mp4 multipart uploads longer than 140 seconds.
	IsLongVideo param.Opt[bool] `json:"is_long_video,omitzero"`
	// HTTPS URL to download and upload as media
	URL            param.Opt[string] `json:"url,omitzero" format:"uri"`
	IdempotencyKey string            `header:"Idempotency-Key" api:"required" json:"-"`
	paramObj
}

func (r XMediaUploadParams) MarshalJSON() (data []byte, err error) {
	type shadow XMediaUploadParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *XMediaUploadParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, nil)
}

func (r XMediaUploadParams) HasFiles() bool {
	return r.File != nil
}

func (r XMediaUploadParams) MarshalMultipart() (data []byte, contentType string, err error) {
	r.File, err = mediaUploadFile(r.File)
	if err != nil {
		return nil, "", err
	}
	return apiform.MarshalMultipart(r)
}

func validateMediaUpload(url param.Opt[string], file io.Reader) error {
	hasURL := !param.IsOmitted(url)
	if hasURL && file != nil {
		return ErrMediaHasURLAndFile
	}
	if !hasURL && file == nil {
		return ErrMediaNeedsURLOrFile
	}
	return nil
}

func mediaUploadFile(file io.Reader) (io.Reader, error) {
	if file == nil {
		return nil, nil
	}
	if typed, ok := file.(interface{ ContentType() string }); ok && typed.ContentType() != "" {
		return file, nil
	}
	var header [12]byte
	n, err := io.ReadFull(file, header[:])
	if err != nil && err != io.EOF && err != io.ErrUnexpectedEOF {
		return nil, err
	}
	filename, contentType := mediaUploadFileType(header[:n])
	if name := uploadFilename(file); name != "" {
		filename = name
	}
	return File(io.MultiReader(bytes.NewReader(header[:n]), file), filename, contentType), nil
}

func uploadFilename(file io.Reader) string {
	if named, ok := file.(interface{ Filename() string }); ok && named.Filename() != "" {
		return named.Filename()
	}
	if named, ok := file.(interface{ Name() string }); ok && named.Name() != "" {
		return path.Base(named.Name())
	}
	return ""
}

func mediaUploadFileType(data []byte) (string, string) {
	switch {
	case bytes.HasPrefix(data, []byte("\x89PNG\r\n\x1a\n")):
		return "upload.png", "image/png"
	case bytes.HasPrefix(data, []byte("\xff\xd8\xff")):
		return "upload.jpg", "image/jpeg"
	case bytes.HasPrefix(data, []byte("GIF87a")), bytes.HasPrefix(data, []byte("GIF89a")):
		return "upload.gif", "image/gif"
	case len(data) >= 12 && string(data[:4]) == "RIFF" && string(data[8:12]) == "WEBP":
		return "upload.webp", "image/webp"
	case len(data) >= 12 && string(data[4:8]) == "ftyp":
		switch string(data[8:12]) {
		case "avif", "avis":
			return "upload.avif", "image/avif"
		case "isom", "iso2", "mp41", "mp42", "avc1", "M4V ":
			return "upload.mp4", "video/mp4"
		}
	}
	return "upload", "application/octet-stream"
}
