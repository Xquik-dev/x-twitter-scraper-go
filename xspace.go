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
	"github.com/Xquik-dev/x-twitter-scraper-go/shared"
)

// X Space details, host, and speakers
//
// XSpaceService contains methods and other services that help with interacting
// with the x-twitter-scraper API.
//
// You should not instantiate this service directly, and instead use the
// [NewXSpaceService] method instead.
type XSpaceService struct {
	options []option.RequestOption
}

// NewXSpaceService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewXSpaceService(opts ...option.RequestOption) (r XSpaceService) {
	r = XSpaceService{}
	r.options = opts
	return
}

// Returns 1 X Space as X shows it, live, scheduled or ended. Accepts the Space ID
// or the URL-encoded Space URL. A Space X limits to subscribers or a community
// answers 404.
func (r *XSpaceService) Get(ctx context.Context, id string, query XSpaceGetParams, opts ...option.RequestOption) (res *XSpaceGetResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("x/spaces/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Returns the recording X plays back once a Space is over, in state Ended or
// TimedOut. Accepts the Space ID or the URL-encoded Space URL. Check state and
// isAvailableForReplay with the Space lookup first. A live or scheduled Space
// answers 409, and a cancelled one 404.
func (r *XSpaceService) GetReplay(ctx context.Context, id string, query XSpaceGetReplayParams, opts ...option.RequestOption) (res *XSpaceGetReplayResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("x/spaces/%s/replay", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Returns public Spaces each once, as GET /x/spaces/{id} serves them. Latest & Top
// find the Spaces that matching posts share, live, scheduled or ended. Live &
// Upcoming read X's own Spaces search, in X's order. It finds live or scheduled
// Spaces by the words in their titles. Costs 1 credit per Space.
func (r *XSpaceService) Search(ctx context.Context, query XSpaceSearchParams, opts ...option.RequestOption) (res *XSpaceSearchResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "x/spaces/search"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type XSpaceGetResponse struct {
	// X Space
	Space XSpaceGetResponseSpace `json:"space" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Space       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XSpaceGetResponse) RawJSON() string { return r.JSON.raw }
func (r *XSpaceGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// X Space
type XSpaceGetResponseSpace struct {
	// Space ID
	ID string `json:"id" api:"required"`
	// Hosts and co-hosts
	Admins []XSpaceGetResponseSpaceAdmin `json:"admins" api:"required"`
	// Accounts X lists as listeners
	Listeners []XSpaceGetResponseSpaceListener `json:"listeners" api:"required"`
	// Accounts X lists as speakers
	Speakers []XSpaceGetResponseSpaceSpeaker `json:"speakers" api:"required"`
	// Audio or video kind, as X names it
	ContentType string `json:"contentType"`
	// X's code for who may speak
	ConversationControls int64 `json:"conversationControls"`
	// When the Space was made
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Public X profile.
	Creator shared.UserProfile `json:"creator"`
	// Whether X lets no more accounts join
	DisallowJoin bool `json:"disallowJoin"`
	// When the Space ended
	EndedAt time.Time `json:"endedAt" format:"date-time"`
	// Whether listeners may clip the Space
	IsAvailableForClipping bool `json:"isAvailableForClipping"`
	// Whether X offers a recording
	IsAvailableForReplay bool `json:"isAvailableForReplay"`
	// Whether only X employees may join
	IsEmployeeOnly bool `json:"isEmployeeOnly"`
	// Whether the host locked the Space
	IsLocked bool `json:"isLocked"`
	// Most hosts and co-hosts the Space allows
	MaxAdminCapacity int64 `json:"maxAdminCapacity"`
	// X's media key for the Space's audio
	MediaKey string `json:"mediaKey"`
	// IDs of the accounts the host tagged
	MentionedUserIDs []string `json:"mentionedUserIds"`
	// Whether anonymous listening is off
	NoIncognito bool `json:"noIncognito"`
	// Accounts in the Space now, while it is live
	ParticipantCount int64 `json:"participantCount"`
	// When the Space was set to start
	ScheduledStart time.Time `json:"scheduledStart" format:"date-time"`
	// Posts the hosts and speakers shared in the Space, in X's order. Present only
	// when X lists them.
	Sharings []XSpaceGetResponseSpaceSharing `json:"sharings"`
	// When the Space went live
	StartedAt time.Time `json:"startedAt" format:"date-time"`
	// NotStarted (scheduled), Running (live), Ended, Canceled (scheduled, never
	// started) or TimedOut (closed by X, not by its host)
	State string `json:"state"`
	// Space title
	Title string `json:"title"`
	// Accounts that listened live
	TotalLiveListeners int64 `json:"totalLiveListeners"`
	// Accounts that played the recording
	TotalReplayWatched int64 `json:"totalReplayWatched"`
	// Search tweet with author. Zero metrics may mean X omitted counts.
	Tweet shared.SearchTweet `json:"tweet"`
	// ID of the post that announces the Space
	TweetID string `json:"tweetId"`
	// When X last changed the Space
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                     respjson.Field
		Admins                 respjson.Field
		Listeners              respjson.Field
		Speakers               respjson.Field
		ContentType            respjson.Field
		ConversationControls   respjson.Field
		CreatedAt              respjson.Field
		Creator                respjson.Field
		DisallowJoin           respjson.Field
		EndedAt                respjson.Field
		IsAvailableForClipping respjson.Field
		IsAvailableForReplay   respjson.Field
		IsEmployeeOnly         respjson.Field
		IsLocked               respjson.Field
		MaxAdminCapacity       respjson.Field
		MediaKey               respjson.Field
		MentionedUserIDs       respjson.Field
		NoIncognito            respjson.Field
		ParticipantCount       respjson.Field
		ScheduledStart         respjson.Field
		Sharings               respjson.Field
		StartedAt              respjson.Field
		State                  respjson.Field
		Title                  respjson.Field
		TotalLiveListeners     respjson.Field
		TotalReplayWatched     respjson.Field
		Tweet                  respjson.Field
		TweetID                respjson.Field
		UpdatedAt              respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XSpaceGetResponseSpace) RawJSON() string { return r.JSON.raw }
func (r *XSpaceGetResponseSpace) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// An account on a Space's stage or in its audience, as X lists it there.
type XSpaceGetResponseSpaceAdmin struct {
	ID string `json:"id" api:"required"`
	// Name the account shows in the Space
	Name     string `json:"name" api:"required"`
	Username string `json:"username" api:"required"`
	// Organization affiliation label shown on an X profile.
	AffiliatesHighlightedLabel XSpaceGetResponseSpaceAdminAffiliatesHighlightedLabel `json:"affiliatesHighlightedLabel"`
	// Whether the account has X Premium verification
	IsBlueVerified bool `json:"isBlueVerified"`
	// Whether a host muted the account
	IsMutedByAdmin bool `json:"isMutedByAdmin"`
	// Whether the account muted itself
	IsMutedByGuest bool `json:"isMutedByGuest"`
	// Whether X shows a verified badge in the Space
	IsVerified bool `json:"isVerified"`
	// When the account joined the Space
	JoinedAt time.Time `json:"joinedAt" format:"date-time"`
	// The account's ID in X's broadcast service
	PeriscopeUserID string `json:"periscopeUserId"`
	// Profile picture URL
	ProfilePicture string `json:"profilePicture"`
	// The participant's profile picture at full size
	ProfilePictureOriginal string `json:"profilePictureOriginal"`
	// Business or Government badge type
	VerifiedType string `json:"verifiedType"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                         respjson.Field
		Name                       respjson.Field
		Username                   respjson.Field
		AffiliatesHighlightedLabel respjson.Field
		IsBlueVerified             respjson.Field
		IsMutedByAdmin             respjson.Field
		IsMutedByGuest             respjson.Field
		IsVerified                 respjson.Field
		JoinedAt                   respjson.Field
		PeriscopeUserID            respjson.Field
		ProfilePicture             respjson.Field
		ProfilePictureOriginal     respjson.Field
		VerifiedType               respjson.Field
		ExtraFields                map[string]respjson.Field
		raw                        string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XSpaceGetResponseSpaceAdmin) RawJSON() string { return r.JSON.raw }
func (r *XSpaceGetResponseSpaceAdmin) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// Organization affiliation label shown on an X profile.
type XSpaceGetResponseSpaceAdminAffiliatesHighlightedLabel struct {
	BadgeURL    string `json:"badgeUrl"`
	Description string `json:"description"`
	// Public text, ranges, references, and mention data.
	LongDescription      PublicMetadata `json:"longDescription"`
	URL                  string         `json:"url"`
	URLType              string         `json:"urlType"`
	UserLabelDisplayType string         `json:"userLabelDisplayType"`
	UserLabelType        string         `json:"userLabelType"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BadgeURL             respjson.Field
		Description          respjson.Field
		LongDescription      respjson.Field
		URL                  respjson.Field
		URLType              respjson.Field
		UserLabelDisplayType respjson.Field
		UserLabelType        respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XSpaceGetResponseSpaceAdminAffiliatesHighlightedLabel) RawJSON() string { return r.JSON.raw }
func (r *XSpaceGetResponseSpaceAdminAffiliatesHighlightedLabel) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// An account on a Space's stage or in its audience, as X lists it there.
type XSpaceGetResponseSpaceListener struct {
	ID string `json:"id" api:"required"`
	// Name the account shows in the Space
	Name     string `json:"name" api:"required"`
	Username string `json:"username" api:"required"`
	// Organization affiliation label shown on an X profile.
	AffiliatesHighlightedLabel XSpaceGetResponseSpaceListenerAffiliatesHighlightedLabel `json:"affiliatesHighlightedLabel"`
	// Whether the account has X Premium verification
	IsBlueVerified bool `json:"isBlueVerified"`
	// Whether a host muted the account
	IsMutedByAdmin bool `json:"isMutedByAdmin"`
	// Whether the account muted itself
	IsMutedByGuest bool `json:"isMutedByGuest"`
	// Whether X shows a verified badge in the Space
	IsVerified bool `json:"isVerified"`
	// When the account joined the Space
	JoinedAt time.Time `json:"joinedAt" format:"date-time"`
	// The account's ID in X's broadcast service
	PeriscopeUserID string `json:"periscopeUserId"`
	// Profile picture URL
	ProfilePicture string `json:"profilePicture"`
	// The participant's profile picture at full size
	ProfilePictureOriginal string `json:"profilePictureOriginal"`
	// Business or Government badge type
	VerifiedType string `json:"verifiedType"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                         respjson.Field
		Name                       respjson.Field
		Username                   respjson.Field
		AffiliatesHighlightedLabel respjson.Field
		IsBlueVerified             respjson.Field
		IsMutedByAdmin             respjson.Field
		IsMutedByGuest             respjson.Field
		IsVerified                 respjson.Field
		JoinedAt                   respjson.Field
		PeriscopeUserID            respjson.Field
		ProfilePicture             respjson.Field
		ProfilePictureOriginal     respjson.Field
		VerifiedType               respjson.Field
		ExtraFields                map[string]respjson.Field
		raw                        string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XSpaceGetResponseSpaceListener) RawJSON() string { return r.JSON.raw }
func (r *XSpaceGetResponseSpaceListener) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// Organization affiliation label shown on an X profile.
type XSpaceGetResponseSpaceListenerAffiliatesHighlightedLabel struct {
	BadgeURL    string `json:"badgeUrl"`
	Description string `json:"description"`
	// Public text, ranges, references, and mention data.
	LongDescription      PublicMetadata `json:"longDescription"`
	URL                  string         `json:"url"`
	URLType              string         `json:"urlType"`
	UserLabelDisplayType string         `json:"userLabelDisplayType"`
	UserLabelType        string         `json:"userLabelType"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BadgeURL             respjson.Field
		Description          respjson.Field
		LongDescription      respjson.Field
		URL                  respjson.Field
		URLType              respjson.Field
		UserLabelDisplayType respjson.Field
		UserLabelType        respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XSpaceGetResponseSpaceListenerAffiliatesHighlightedLabel) RawJSON() string { return r.JSON.raw }
func (r *XSpaceGetResponseSpaceListenerAffiliatesHighlightedLabel) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// An account on a Space's stage or in its audience, as X lists it there.
type XSpaceGetResponseSpaceSpeaker struct {
	ID string `json:"id" api:"required"`
	// Name the account shows in the Space
	Name     string `json:"name" api:"required"`
	Username string `json:"username" api:"required"`
	// Organization affiliation label shown on an X profile.
	AffiliatesHighlightedLabel XSpaceGetResponseSpaceSpeakerAffiliatesHighlightedLabel `json:"affiliatesHighlightedLabel"`
	// Whether the account has X Premium verification
	IsBlueVerified bool `json:"isBlueVerified"`
	// Whether a host muted the account
	IsMutedByAdmin bool `json:"isMutedByAdmin"`
	// Whether the account muted itself
	IsMutedByGuest bool `json:"isMutedByGuest"`
	// Whether X shows a verified badge in the Space
	IsVerified bool `json:"isVerified"`
	// When the account joined the Space
	JoinedAt time.Time `json:"joinedAt" format:"date-time"`
	// The account's ID in X's broadcast service
	PeriscopeUserID string `json:"periscopeUserId"`
	// Profile picture URL
	ProfilePicture string `json:"profilePicture"`
	// The participant's profile picture at full size
	ProfilePictureOriginal string `json:"profilePictureOriginal"`
	// Business or Government badge type
	VerifiedType string `json:"verifiedType"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                         respjson.Field
		Name                       respjson.Field
		Username                   respjson.Field
		AffiliatesHighlightedLabel respjson.Field
		IsBlueVerified             respjson.Field
		IsMutedByAdmin             respjson.Field
		IsMutedByGuest             respjson.Field
		IsVerified                 respjson.Field
		JoinedAt                   respjson.Field
		PeriscopeUserID            respjson.Field
		ProfilePicture             respjson.Field
		ProfilePictureOriginal     respjson.Field
		VerifiedType               respjson.Field
		ExtraFields                map[string]respjson.Field
		raw                        string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XSpaceGetResponseSpaceSpeaker) RawJSON() string { return r.JSON.raw }
func (r *XSpaceGetResponseSpaceSpeaker) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// Organization affiliation label shown on an X profile.
type XSpaceGetResponseSpaceSpeakerAffiliatesHighlightedLabel struct {
	BadgeURL    string `json:"badgeUrl"`
	Description string `json:"description"`
	// Public text, ranges, references, and mention data.
	LongDescription      PublicMetadata `json:"longDescription"`
	URL                  string         `json:"url"`
	URLType              string         `json:"urlType"`
	UserLabelDisplayType string         `json:"userLabelDisplayType"`
	UserLabelType        string         `json:"userLabelType"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BadgeURL             respjson.Field
		Description          respjson.Field
		LongDescription      respjson.Field
		URL                  respjson.Field
		URLType              respjson.Field
		UserLabelDisplayType respjson.Field
		UserLabelType        respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XSpaceGetResponseSpaceSpeakerAffiliatesHighlightedLabel) RawJSON() string { return r.JSON.raw }
func (r *XSpaceGetResponseSpaceSpeakerAffiliatesHighlightedLabel) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// A post shared in a Space, & the account that shared it.
type XSpaceGetResponseSpaceSharing struct {
	// X's ID of the sharing, not of the post
	ID string `json:"id"`
	// When the post was shared
	SharedAt time.Time `json:"sharedAt" format:"date-time"`
	// Public X profile.
	SharedBy shared.UserProfile `json:"sharedBy"`
	// Search tweet with author. Zero metrics may mean X omitted counts.
	Tweet shared.SearchTweet `json:"tweet"`
	// When X last changed the sharing
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		SharedAt    respjson.Field
		SharedBy    respjson.Field
		Tweet       respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XSpaceGetResponseSpaceSharing) RawJSON() string { return r.JSON.raw }
func (r *XSpaceGetResponseSpaceSharing) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type XSpaceGetReplayResponse struct {
	// Recording of an ended Space
	Replay XSpaceGetReplayResponseReplay `json:"replay" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Replay      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XSpaceGetReplayResponse) RawJSON() string { return r.JSON.raw }
func (r *XSpaceGetReplayResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// Recording of an ended Space
type XSpaceGetReplayResponseReplay struct {
	// HLS playlist of the recording
	URL string `json:"url" api:"required" format:"uri"`
	// Stream status, as X names it
	Status string `json:"status"`
	// Stream format, as X names it
	StreamType string `json:"streamType"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		URL         respjson.Field
		Status      respjson.Field
		StreamType  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XSpaceGetReplayResponseReplay) RawJSON() string { return r.JSON.raw }
func (r *XSpaceGetReplayResponseReplay) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type XSpaceSearchResponse struct {
	HasNextPage bool   `json:"has_next_page" api:"required"`
	NextCursor  string `json:"next_cursor" api:"required"`
	// Public Spaces, in the order the posts share them
	Spaces []XSpaceSearchResponseSpace `json:"spaces" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		HasNextPage respjson.Field
		NextCursor  respjson.Field
		Spaces      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XSpaceSearchResponse) RawJSON() string { return r.JSON.raw }
func (r *XSpaceSearchResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// X Space
type XSpaceSearchResponseSpace struct {
	// Space ID
	ID string `json:"id" api:"required"`
	// Hosts and co-hosts
	Admins []XSpaceSearchResponseSpaceAdmin `json:"admins" api:"required"`
	// Accounts X lists as listeners
	Listeners []XSpaceSearchResponseSpaceListener `json:"listeners" api:"required"`
	// Accounts X lists as speakers
	Speakers []XSpaceSearchResponseSpaceSpeaker `json:"speakers" api:"required"`
	// Audio or video kind, as X names it
	ContentType string `json:"contentType"`
	// X's code for who may speak
	ConversationControls int64 `json:"conversationControls"`
	// When the Space was made
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Public X profile.
	Creator shared.UserProfile `json:"creator"`
	// Whether X lets no more accounts join
	DisallowJoin bool `json:"disallowJoin"`
	// When the Space ended
	EndedAt time.Time `json:"endedAt" format:"date-time"`
	// Whether listeners may clip the Space
	IsAvailableForClipping bool `json:"isAvailableForClipping"`
	// Whether X offers a recording
	IsAvailableForReplay bool `json:"isAvailableForReplay"`
	// Whether only X employees may join
	IsEmployeeOnly bool `json:"isEmployeeOnly"`
	// Whether the host locked the Space
	IsLocked bool `json:"isLocked"`
	// Most hosts and co-hosts the Space allows
	MaxAdminCapacity int64 `json:"maxAdminCapacity"`
	// X's media key for the Space's audio
	MediaKey string `json:"mediaKey"`
	// IDs of the accounts the host tagged
	MentionedUserIDs []string `json:"mentionedUserIds"`
	// Whether anonymous listening is off
	NoIncognito bool `json:"noIncognito"`
	// Accounts in the Space now, while it is live
	ParticipantCount int64 `json:"participantCount"`
	// When the Space was set to start
	ScheduledStart time.Time `json:"scheduledStart" format:"date-time"`
	// Posts the hosts and speakers shared in the Space, in X's order. Present only
	// when X lists them.
	Sharings []XSpaceSearchResponseSpaceSharing `json:"sharings"`
	// When the Space went live
	StartedAt time.Time `json:"startedAt" format:"date-time"`
	// NotStarted (scheduled), Running (live), Ended, Canceled (scheduled, never
	// started) or TimedOut (closed by X, not by its host)
	State string `json:"state"`
	// Space title
	Title string `json:"title"`
	// Accounts that listened live
	TotalLiveListeners int64 `json:"totalLiveListeners"`
	// Accounts that played the recording
	TotalReplayWatched int64 `json:"totalReplayWatched"`
	// Search tweet with author. Zero metrics may mean X omitted counts.
	Tweet shared.SearchTweet `json:"tweet"`
	// ID of the post that announces the Space
	TweetID string `json:"tweetId"`
	// When X last changed the Space
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                     respjson.Field
		Admins                 respjson.Field
		Listeners              respjson.Field
		Speakers               respjson.Field
		ContentType            respjson.Field
		ConversationControls   respjson.Field
		CreatedAt              respjson.Field
		Creator                respjson.Field
		DisallowJoin           respjson.Field
		EndedAt                respjson.Field
		IsAvailableForClipping respjson.Field
		IsAvailableForReplay   respjson.Field
		IsEmployeeOnly         respjson.Field
		IsLocked               respjson.Field
		MaxAdminCapacity       respjson.Field
		MediaKey               respjson.Field
		MentionedUserIDs       respjson.Field
		NoIncognito            respjson.Field
		ParticipantCount       respjson.Field
		ScheduledStart         respjson.Field
		Sharings               respjson.Field
		StartedAt              respjson.Field
		State                  respjson.Field
		Title                  respjson.Field
		TotalLiveListeners     respjson.Field
		TotalReplayWatched     respjson.Field
		Tweet                  respjson.Field
		TweetID                respjson.Field
		UpdatedAt              respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XSpaceSearchResponseSpace) RawJSON() string { return r.JSON.raw }
func (r *XSpaceSearchResponseSpace) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// An account on a Space's stage or in its audience, as X lists it there.
type XSpaceSearchResponseSpaceAdmin struct {
	ID string `json:"id" api:"required"`
	// Name the account shows in the Space
	Name     string `json:"name" api:"required"`
	Username string `json:"username" api:"required"`
	// Organization affiliation label shown on an X profile.
	AffiliatesHighlightedLabel XSpaceSearchResponseSpaceAdminAffiliatesHighlightedLabel `json:"affiliatesHighlightedLabel"`
	// Whether the account has X Premium verification
	IsBlueVerified bool `json:"isBlueVerified"`
	// Whether a host muted the account
	IsMutedByAdmin bool `json:"isMutedByAdmin"`
	// Whether the account muted itself
	IsMutedByGuest bool `json:"isMutedByGuest"`
	// Whether X shows a verified badge in the Space
	IsVerified bool `json:"isVerified"`
	// When the account joined the Space
	JoinedAt time.Time `json:"joinedAt" format:"date-time"`
	// The account's ID in X's broadcast service
	PeriscopeUserID string `json:"periscopeUserId"`
	// Profile picture URL
	ProfilePicture string `json:"profilePicture"`
	// The participant's profile picture at full size
	ProfilePictureOriginal string `json:"profilePictureOriginal"`
	// Business or Government badge type
	VerifiedType string `json:"verifiedType"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                         respjson.Field
		Name                       respjson.Field
		Username                   respjson.Field
		AffiliatesHighlightedLabel respjson.Field
		IsBlueVerified             respjson.Field
		IsMutedByAdmin             respjson.Field
		IsMutedByGuest             respjson.Field
		IsVerified                 respjson.Field
		JoinedAt                   respjson.Field
		PeriscopeUserID            respjson.Field
		ProfilePicture             respjson.Field
		ProfilePictureOriginal     respjson.Field
		VerifiedType               respjson.Field
		ExtraFields                map[string]respjson.Field
		raw                        string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XSpaceSearchResponseSpaceAdmin) RawJSON() string { return r.JSON.raw }
func (r *XSpaceSearchResponseSpaceAdmin) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// Organization affiliation label shown on an X profile.
type XSpaceSearchResponseSpaceAdminAffiliatesHighlightedLabel struct {
	BadgeURL    string `json:"badgeUrl"`
	Description string `json:"description"`
	// Public text, ranges, references, and mention data.
	LongDescription      PublicMetadata `json:"longDescription"`
	URL                  string         `json:"url"`
	URLType              string         `json:"urlType"`
	UserLabelDisplayType string         `json:"userLabelDisplayType"`
	UserLabelType        string         `json:"userLabelType"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BadgeURL             respjson.Field
		Description          respjson.Field
		LongDescription      respjson.Field
		URL                  respjson.Field
		URLType              respjson.Field
		UserLabelDisplayType respjson.Field
		UserLabelType        respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XSpaceSearchResponseSpaceAdminAffiliatesHighlightedLabel) RawJSON() string { return r.JSON.raw }
func (r *XSpaceSearchResponseSpaceAdminAffiliatesHighlightedLabel) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// An account on a Space's stage or in its audience, as X lists it there.
type XSpaceSearchResponseSpaceListener struct {
	ID string `json:"id" api:"required"`
	// Name the account shows in the Space
	Name     string `json:"name" api:"required"`
	Username string `json:"username" api:"required"`
	// Organization affiliation label shown on an X profile.
	AffiliatesHighlightedLabel XSpaceSearchResponseSpaceListenerAffiliatesHighlightedLabel `json:"affiliatesHighlightedLabel"`
	// Whether the account has X Premium verification
	IsBlueVerified bool `json:"isBlueVerified"`
	// Whether a host muted the account
	IsMutedByAdmin bool `json:"isMutedByAdmin"`
	// Whether the account muted itself
	IsMutedByGuest bool `json:"isMutedByGuest"`
	// Whether X shows a verified badge in the Space
	IsVerified bool `json:"isVerified"`
	// When the account joined the Space
	JoinedAt time.Time `json:"joinedAt" format:"date-time"`
	// The account's ID in X's broadcast service
	PeriscopeUserID string `json:"periscopeUserId"`
	// Profile picture URL
	ProfilePicture string `json:"profilePicture"`
	// The participant's profile picture at full size
	ProfilePictureOriginal string `json:"profilePictureOriginal"`
	// Business or Government badge type
	VerifiedType string `json:"verifiedType"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                         respjson.Field
		Name                       respjson.Field
		Username                   respjson.Field
		AffiliatesHighlightedLabel respjson.Field
		IsBlueVerified             respjson.Field
		IsMutedByAdmin             respjson.Field
		IsMutedByGuest             respjson.Field
		IsVerified                 respjson.Field
		JoinedAt                   respjson.Field
		PeriscopeUserID            respjson.Field
		ProfilePicture             respjson.Field
		ProfilePictureOriginal     respjson.Field
		VerifiedType               respjson.Field
		ExtraFields                map[string]respjson.Field
		raw                        string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XSpaceSearchResponseSpaceListener) RawJSON() string { return r.JSON.raw }
func (r *XSpaceSearchResponseSpaceListener) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// Organization affiliation label shown on an X profile.
type XSpaceSearchResponseSpaceListenerAffiliatesHighlightedLabel struct {
	BadgeURL    string `json:"badgeUrl"`
	Description string `json:"description"`
	// Public text, ranges, references, and mention data.
	LongDescription      PublicMetadata `json:"longDescription"`
	URL                  string         `json:"url"`
	URLType              string         `json:"urlType"`
	UserLabelDisplayType string         `json:"userLabelDisplayType"`
	UserLabelType        string         `json:"userLabelType"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BadgeURL             respjson.Field
		Description          respjson.Field
		LongDescription      respjson.Field
		URL                  respjson.Field
		URLType              respjson.Field
		UserLabelDisplayType respjson.Field
		UserLabelType        respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XSpaceSearchResponseSpaceListenerAffiliatesHighlightedLabel) RawJSON() string {
	return r.JSON.raw
}
func (r *XSpaceSearchResponseSpaceListenerAffiliatesHighlightedLabel) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// An account on a Space's stage or in its audience, as X lists it there.
type XSpaceSearchResponseSpaceSpeaker struct {
	ID string `json:"id" api:"required"`
	// Name the account shows in the Space
	Name     string `json:"name" api:"required"`
	Username string `json:"username" api:"required"`
	// Organization affiliation label shown on an X profile.
	AffiliatesHighlightedLabel XSpaceSearchResponseSpaceSpeakerAffiliatesHighlightedLabel `json:"affiliatesHighlightedLabel"`
	// Whether the account has X Premium verification
	IsBlueVerified bool `json:"isBlueVerified"`
	// Whether a host muted the account
	IsMutedByAdmin bool `json:"isMutedByAdmin"`
	// Whether the account muted itself
	IsMutedByGuest bool `json:"isMutedByGuest"`
	// Whether X shows a verified badge in the Space
	IsVerified bool `json:"isVerified"`
	// When the account joined the Space
	JoinedAt time.Time `json:"joinedAt" format:"date-time"`
	// The account's ID in X's broadcast service
	PeriscopeUserID string `json:"periscopeUserId"`
	// Profile picture URL
	ProfilePicture string `json:"profilePicture"`
	// The participant's profile picture at full size
	ProfilePictureOriginal string `json:"profilePictureOriginal"`
	// Business or Government badge type
	VerifiedType string `json:"verifiedType"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                         respjson.Field
		Name                       respjson.Field
		Username                   respjson.Field
		AffiliatesHighlightedLabel respjson.Field
		IsBlueVerified             respjson.Field
		IsMutedByAdmin             respjson.Field
		IsMutedByGuest             respjson.Field
		IsVerified                 respjson.Field
		JoinedAt                   respjson.Field
		PeriscopeUserID            respjson.Field
		ProfilePicture             respjson.Field
		ProfilePictureOriginal     respjson.Field
		VerifiedType               respjson.Field
		ExtraFields                map[string]respjson.Field
		raw                        string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XSpaceSearchResponseSpaceSpeaker) RawJSON() string { return r.JSON.raw }
func (r *XSpaceSearchResponseSpaceSpeaker) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// Organization affiliation label shown on an X profile.
type XSpaceSearchResponseSpaceSpeakerAffiliatesHighlightedLabel struct {
	BadgeURL    string `json:"badgeUrl"`
	Description string `json:"description"`
	// Public text, ranges, references, and mention data.
	LongDescription      PublicMetadata `json:"longDescription"`
	URL                  string         `json:"url"`
	URLType              string         `json:"urlType"`
	UserLabelDisplayType string         `json:"userLabelDisplayType"`
	UserLabelType        string         `json:"userLabelType"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BadgeURL             respjson.Field
		Description          respjson.Field
		LongDescription      respjson.Field
		URL                  respjson.Field
		URLType              respjson.Field
		UserLabelDisplayType respjson.Field
		UserLabelType        respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XSpaceSearchResponseSpaceSpeakerAffiliatesHighlightedLabel) RawJSON() string {
	return r.JSON.raw
}
func (r *XSpaceSearchResponseSpaceSpeakerAffiliatesHighlightedLabel) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// A post shared in a Space, & the account that shared it.
type XSpaceSearchResponseSpaceSharing struct {
	// X's ID of the sharing, not of the post
	ID string `json:"id"`
	// When the post was shared
	SharedAt time.Time `json:"sharedAt" format:"date-time"`
	// Public X profile.
	SharedBy shared.UserProfile `json:"sharedBy"`
	// Search tweet with author. Zero metrics may mean X omitted counts.
	Tweet shared.SearchTweet `json:"tweet"`
	// When X last changed the sharing
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		SharedAt    respjson.Field
		SharedBy    respjson.Field
		Tweet       respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XSpaceSearchResponseSpaceSharing) RawJSON() string { return r.JSON.raw }
func (r *XSpaceSearchResponseSpaceSharing) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type XSpaceGetParams struct {
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

// URLQuery serializes [XSpaceGetParams]'s query parameters as `url.Values`.
func (r XSpaceGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type XSpaceGetReplayParams struct {
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

// URLQuery serializes [XSpaceGetReplayParams]'s query parameters as `url.Values`.
func (r XSpaceGetReplayParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type XSpaceSearchParams struct {
	// Words, hashtags or X search operators, such as from:solana. Latest & Top add
	// filter:spaces. Live & Upcoming match title words only, without operators. Alias
	// query.
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
	// For Latest & Top, the posts a page reads, 20 at a time (1-100, default 20).
	// Posts often share the same Space, so a page returns fewer Spaces. Send 100 to
	// read a search fastest. For Live & Upcoming, the most Spaces a page returns. Its
	// aliases are limit, count, max_results, maxItems, max_items & per_page.
	PageSize param.Opt[int64] `query:"pageSize,omitzero" json:"-"`
	// Latest finds Spaces in the newest posts. Top finds them in the most engaging
	// posts. Live lists Spaces on air now. Upcoming lists scheduled ones, as X's
	// Spaces search does. Any case works. Aliases result_type, sort_order, type,
	// search_type, product, category & section.
	//
	// Any of "Latest", "Top", "Live", "Upcoming".
	QueryType XSpaceSearchParamsQueryType `query:"queryType,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [XSpaceSearchParams]'s query parameters as `url.Values`.
func (r XSpaceSearchParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Latest finds Spaces in the newest posts. Top finds them in the most engaging
// posts. Live lists Spaces on air now. Upcoming lists scheduled ones, as X's
// Spaces search does. Any case works. Aliases result_type, sort_order, type,
// search_type, product, category & section.
type XSpaceSearchParamsQueryType string

const (
	XSpaceSearchParamsQueryTypeLatest   XSpaceSearchParamsQueryType = "Latest"
	XSpaceSearchParamsQueryTypeTop      XSpaceSearchParamsQueryType = "Top"
	XSpaceSearchParamsQueryTypeLive     XSpaceSearchParamsQueryType = "Live"
	XSpaceSearchParamsQueryTypeUpcoming XSpaceSearchParamsQueryType = "Upcoming"
)
