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

// Community Notes contributors, their notes, ratings & deleted notes
//
// XCommunityNoteService contains methods and other services that help with
// interacting with the x-twitter-scraper API.
//
// You should not instantiate this service directly, and instead use the
// [NewXCommunityNoteService] method instead.
type XCommunityNoteService struct {
	options []option.RequestOption
}

// NewXCommunityNoteService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewXCommunityNoteService(opts ...option.RequestOption) (r XCommunityNoteService) {
	r = XCommunityNoteService{}
	r.options = opts
	return
}

// Returns 1 Community Notes contributor by the alias X gives them. Includes note
// counts by rating, rating counts & the deleted notes count. Each lookup costs 1
// credit.
func (r *XCommunityNoteService) GetContributor(ctx context.Context, alias string, query XCommunityNoteGetContributorParams, opts ...option.RequestOption) (res *XCommunityNoteGetContributorResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if alias == "" {
		err = errors.New("missing required alias parameter")
		return nil, err
	}
	path := fmt.Sprintf("x/community-notes/contributors/%s", url.PathEscape(alias))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Returns the notes a contributor deleted, newest first. X keeps each note's ID,
// times, rating & tags. A notice replaces its text & post. Each note costs 1
// credit. A contributor who deleted none returns an empty list, free.
func (r *XCommunityNoteService) ListContributorDeletedNotes(ctx context.Context, alias string, query XCommunityNoteListContributorDeletedNotesParams, opts ...option.RequestOption) (res *XCommunityNoteListContributorDeletedNotesResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if alias == "" {
		err = errors.New("missing required alias parameter")
		return nil, err
	}
	path := fmt.Sprintf("x/community-notes/contributors/%s/deleted-notes", url.PathEscape(alias))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Returns a contributor's notes, newest first, each with the post it is on.
// Includes notes X has not rated helpful, which x.com shows under no post. Each
// note costs 1 credit. Follow next_cursor for more.
func (r *XCommunityNoteService) ListContributorNotes(ctx context.Context, alias string, query XCommunityNoteListContributorNotesParams, opts ...option.RequestOption) (res *XCommunityNoteListContributorNotesResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if alias == "" {
		err = errors.New("missing required alias parameter")
		return nil, err
	}
	path := fmt.Sprintf("x/community-notes/contributors/%s/notes", url.PathEscape(alias))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type XCommunityNoteGetContributorResponse struct {
	// A Community Notes contributor. X names contributors by an alias, never by their
	// X account. A field X leaves out is left out. Beside a note, X sends its writer
	// with only some of the counts.
	Contributor XCommunityNoteGetContributorResponseContributor `json:"contributor" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Contributor respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XCommunityNoteGetContributorResponse) RawJSON() string { return r.JSON.raw }
func (r *XCommunityNoteGetContributorResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// A Community Notes contributor. X names contributors by an alias, never by their
// X account. A field X leaves out is left out. Beside a note, X sends its writer
// with only some of the counts.
type XCommunityNoteGetContributorResponseContributor struct {
	// The alias X gives the contributor.
	Alias string `json:"alias" api:"required"`
	// Notes the contributor deleted.
	DeletedNotesCount int64 `json:"deletedNotesCount"`
	// Whether the contributor writes notes through X's API.
	IsAPIContributor bool `json:"isApiContributor"`
	// Whether X marks the contributor as a top writer.
	IsTopWriter bool `json:"isTopWriter"`
	// How many notes a contributor wrote, by X's current rating.
	NotesCount XCommunityNoteGetContributorResponseContributorNotesCount `json:"notesCount"`
	// How a contributor's ratings of other notes went.
	RatingsCount XCommunityNoteGetContributorResponseContributorRatingsCount `json:"ratingsCount"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Alias             respjson.Field
		DeletedNotesCount respjson.Field
		IsAPIContributor  respjson.Field
		IsTopWriter       respjson.Field
		NotesCount        respjson.Field
		RatingsCount      respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XCommunityNoteGetContributorResponseContributor) RawJSON() string { return r.JSON.raw }
func (r *XCommunityNoteGetContributorResponseContributor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// How many notes a contributor wrote, by X's current rating.
type XCommunityNoteGetContributorResponseContributorNotesCount struct {
	// Notes that need more ratings.
	AwaitingMoreRatings int64 `json:"awaitingMoreRatings"`
	// Notes X rates helpful now.
	CurrentlyRatedHelpful int64 `json:"currentlyRatedHelpful"`
	// Notes X rates not helpful now.
	CurrentlyRatedNotHelpful int64 `json:"currentlyRatedNotHelpful"`
	// When X last counted the notes.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AwaitingMoreRatings      respjson.Field
		CurrentlyRatedHelpful    respjson.Field
		CurrentlyRatedNotHelpful respjson.Field
		UpdatedAt                respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XCommunityNoteGetContributorResponseContributorNotesCount) RawJSON() string {
	return r.JSON.raw
}
func (r *XCommunityNoteGetContributorResponseContributorNotesCount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// How a contributor's ratings of other notes went.
type XCommunityNoteGetContributorResponseContributorRatingsCount struct {
	// Ratings on notes that need more ratings.
	AwaitingMoreRatings int64 `json:"awaitingMoreRatings"`
	// Ratings made after X rated the note.
	RatedAfterDecision int64 `json:"ratedAfterDecision"`
	// Ratings that matched the rating X gave the note.
	Successful XCommunityNoteGetContributorResponseContributorRatingsCountSuccessful `json:"successful"`
	// Ratings that went against the rating X gave the note.
	Unsuccessful XCommunityNoteGetContributorResponseContributorRatingsCountUnsuccessful `json:"unsuccessful"`
	// When X last counted the ratings.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AwaitingMoreRatings respjson.Field
		RatedAfterDecision  respjson.Field
		Successful          respjson.Field
		Unsuccessful        respjson.Field
		UpdatedAt           respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XCommunityNoteGetContributorResponseContributorRatingsCount) RawJSON() string {
	return r.JSON.raw
}
func (r *XCommunityNoteGetContributorResponseContributorRatingsCount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// Ratings that matched the rating X gave the note.
type XCommunityNoteGetContributorResponseContributorRatingsCountSuccessful struct {
	// Ratings that called a note helpful.
	HelpfulCount int64 `json:"helpfulCount"`
	// Ratings that called a note not helpful.
	NotHelpfulCount int64 `json:"notHelpfulCount"`
	// All ratings in the group, when X sends the sum.
	Total int64 `json:"total"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		HelpfulCount    respjson.Field
		NotHelpfulCount respjson.Field
		Total           respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XCommunityNoteGetContributorResponseContributorRatingsCountSuccessful) RawJSON() string {
	return r.JSON.raw
}
func (r *XCommunityNoteGetContributorResponseContributorRatingsCountSuccessful) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// Ratings that went against the rating X gave the note.
type XCommunityNoteGetContributorResponseContributorRatingsCountUnsuccessful struct {
	// Ratings that called a note helpful.
	HelpfulCount int64 `json:"helpfulCount"`
	// Ratings that called a note not helpful.
	NotHelpfulCount int64 `json:"notHelpfulCount"`
	// All ratings in the group, when X sends the sum.
	Total int64 `json:"total"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		HelpfulCount    respjson.Field
		NotHelpfulCount respjson.Field
		Total           respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XCommunityNoteGetContributorResponseContributorRatingsCountUnsuccessful) RawJSON() string {
	return r.JSON.raw
}
func (r *XCommunityNoteGetContributorResponseContributorRatingsCountUnsuccessful) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// 1 page of Community Notes. Follow next_cursor while the response reports more
// pages.
type XCommunityNoteListContributorDeletedNotesResponse struct {
	HasNextPage bool                                                    `json:"has_next_page" api:"required"`
	NextCursor  string                                                  `json:"next_cursor" api:"required"`
	Notes       []XCommunityNoteListContributorDeletedNotesResponseNote `json:"notes" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		HasNextPage respjson.Field
		NextCursor  respjson.Field
		Notes       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XCommunityNoteListContributorDeletedNotesResponse) RawJSON() string { return r.JSON.raw }
func (r *XCommunityNoteListContributorDeletedNotesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// A Community Note as its writer's page lists it, with the post it is on. A
// deleted note has no text, classification or post. A field X leaves out is left
// out.
type XCommunityNoteListContributorDeletedNotesResponseNote struct {
	// Note ID.
	ID string `json:"id" api:"required"`
	// Whether the note scores above X's bar for a possible helpful rating.
	AbovePcrhThreshold bool `json:"abovePcrhThreshold"`
	// Whether the rating is under appeal.
	AppealStatus string `json:"appealStatus"`
	// What the writer says of the post, such as MisinformedOrPotentiallyMisleading or
	// NotMisleading.
	Classification string `json:"classification"`
	// The note's writer, with the counts X sends beside a note. GET
	// /x/community-notes/contributors/{alias} serves every count.
	Contributor XCommunityNoteListContributorDeletedNotesResponseNoteContributor `json:"contributor"`
	// When the contributor wrote the note.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// The X scoring model that set the rating.
	DecidedBy string `json:"decidedBy"`
	// Whether the writer deleted the note.
	Deleted bool `json:"deleted"`
	// The link a link note is on, as X shows it.
	DisplayURL string `json:"displayUrl"`
	// Links, tags & accounts inside the text, by character range.
	Entities []XCommunityNoteListContributorDeletedNotesResponseNoteEntity `json:"entities"`
	// Whether X marks that scoring model as fully visible.
	FullyVisibleModel bool `json:"fullyVisibleModel"`
	// Why raters called the note helpful.
	HelpfulTags []string `json:"helpfulTags"`
	// Views of the note, when X shares them.
	ImpressionCount int64 `json:"impressionCount"`
	// Whether the note came through X's API.
	IsAPIAuthor bool `json:"isApiAuthor"`
	// Whether the note is about the post's media. X can show such a note on other
	// posts.
	IsMediaNote bool `json:"isMediaNote"`
	// Whether X offers a translation of the note.
	IsTranslatable bool `json:"isTranslatable"`
	// Whether the note is on a link.
	IsURLNote bool `json:"isUrlNote"`
	// Language code of the note.
	Language string `json:"language"`
	// Other posts X matched a media note to.
	MediaNoteMatchCount int64 `json:"mediaNoteMatchCount"`
	// Why raters called the note not helpful.
	NotHelpfulTags []string `json:"notHelpfulTags"`
	// The post the note is on. Null when X no longer shows it.
	Post shared.SearchTweet `json:"post" api:"nullable"`
	// ID of the post the note is on.
	PostID string `json:"postId"`
	// Why X shows no post, when X says.
	PostUnavailable XCommunityNoteListContributorDeletedNotesResponseNotePostUnavailable `json:"postUnavailable"`
	// X's current rating, such as NeedsMoreRatings, CurrentlyRatedHelpful or
	// CurrentlyRatedNotHelpful. x.com shows only helpful notes under posts.
	RatingStatus string `json:"ratingStatus"`
	// What the note says.
	Text string `json:"text"`
	// Milliseconds from the note's creation to its helpful rating.
	TimeToHelpfulFromNoteMs int64 `json:"timeToHelpfulFromNoteMs"`
	// Milliseconds from the post's creation to the note's helpful rating.
	TimeToHelpfulFromPostMs int64 `json:"timeToHelpfulFromPostMs"`
	// X's notice in place of a deleted note.
	Tombstone string `json:"tombstone"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                      respjson.Field
		AbovePcrhThreshold      respjson.Field
		AppealStatus            respjson.Field
		Classification          respjson.Field
		Contributor             respjson.Field
		CreatedAt               respjson.Field
		DecidedBy               respjson.Field
		Deleted                 respjson.Field
		DisplayURL              respjson.Field
		Entities                respjson.Field
		FullyVisibleModel       respjson.Field
		HelpfulTags             respjson.Field
		ImpressionCount         respjson.Field
		IsAPIAuthor             respjson.Field
		IsMediaNote             respjson.Field
		IsTranslatable          respjson.Field
		IsURLNote               respjson.Field
		Language                respjson.Field
		MediaNoteMatchCount     respjson.Field
		NotHelpfulTags          respjson.Field
		Post                    respjson.Field
		PostID                  respjson.Field
		PostUnavailable         respjson.Field
		RatingStatus            respjson.Field
		Text                    respjson.Field
		TimeToHelpfulFromNoteMs respjson.Field
		TimeToHelpfulFromPostMs respjson.Field
		Tombstone               respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XCommunityNoteListContributorDeletedNotesResponseNote) RawJSON() string { return r.JSON.raw }
func (r *XCommunityNoteListContributorDeletedNotesResponseNote) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// The note's writer, with the counts X sends beside a note. GET
// /x/community-notes/contributors/{alias} serves every count.
type XCommunityNoteListContributorDeletedNotesResponseNoteContributor struct {
	// The alias X gives the contributor.
	Alias string `json:"alias" api:"required"`
	// Notes the contributor deleted.
	DeletedNotesCount int64 `json:"deletedNotesCount"`
	// Whether the contributor writes notes through X's API.
	IsAPIContributor bool `json:"isApiContributor"`
	// Whether X marks the contributor as a top writer.
	IsTopWriter bool `json:"isTopWriter"`
	// How many notes a contributor wrote, by X's current rating.
	NotesCount XCommunityNoteListContributorDeletedNotesResponseNoteContributorNotesCount `json:"notesCount"`
	// How a contributor's ratings of other notes went.
	RatingsCount XCommunityNoteListContributorDeletedNotesResponseNoteContributorRatingsCount `json:"ratingsCount"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Alias             respjson.Field
		DeletedNotesCount respjson.Field
		IsAPIContributor  respjson.Field
		IsTopWriter       respjson.Field
		NotesCount        respjson.Field
		RatingsCount      respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XCommunityNoteListContributorDeletedNotesResponseNoteContributor) RawJSON() string {
	return r.JSON.raw
}
func (r *XCommunityNoteListContributorDeletedNotesResponseNoteContributor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// How many notes a contributor wrote, by X's current rating.
type XCommunityNoteListContributorDeletedNotesResponseNoteContributorNotesCount struct {
	// Notes that need more ratings.
	AwaitingMoreRatings int64 `json:"awaitingMoreRatings"`
	// Notes X rates helpful now.
	CurrentlyRatedHelpful int64 `json:"currentlyRatedHelpful"`
	// Notes X rates not helpful now.
	CurrentlyRatedNotHelpful int64 `json:"currentlyRatedNotHelpful"`
	// When X last counted the notes.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AwaitingMoreRatings      respjson.Field
		CurrentlyRatedHelpful    respjson.Field
		CurrentlyRatedNotHelpful respjson.Field
		UpdatedAt                respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XCommunityNoteListContributorDeletedNotesResponseNoteContributorNotesCount) RawJSON() string {
	return r.JSON.raw
}
func (r *XCommunityNoteListContributorDeletedNotesResponseNoteContributorNotesCount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// How a contributor's ratings of other notes went.
type XCommunityNoteListContributorDeletedNotesResponseNoteContributorRatingsCount struct {
	// Ratings on notes that need more ratings.
	AwaitingMoreRatings int64 `json:"awaitingMoreRatings"`
	// Ratings made after X rated the note.
	RatedAfterDecision int64 `json:"ratedAfterDecision"`
	// Ratings that matched the rating X gave the note.
	Successful XCommunityNoteListContributorDeletedNotesResponseNoteContributorRatingsCountSuccessful `json:"successful"`
	// Ratings that went against the rating X gave the note.
	Unsuccessful XCommunityNoteListContributorDeletedNotesResponseNoteContributorRatingsCountUnsuccessful `json:"unsuccessful"`
	// When X last counted the ratings.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AwaitingMoreRatings respjson.Field
		RatedAfterDecision  respjson.Field
		Successful          respjson.Field
		Unsuccessful        respjson.Field
		UpdatedAt           respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XCommunityNoteListContributorDeletedNotesResponseNoteContributorRatingsCount) RawJSON() string {
	return r.JSON.raw
}
func (r *XCommunityNoteListContributorDeletedNotesResponseNoteContributorRatingsCount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// Ratings that matched the rating X gave the note.
type XCommunityNoteListContributorDeletedNotesResponseNoteContributorRatingsCountSuccessful struct {
	// Ratings that called a note helpful.
	HelpfulCount int64 `json:"helpfulCount"`
	// Ratings that called a note not helpful.
	NotHelpfulCount int64 `json:"notHelpfulCount"`
	// All ratings in the group, when X sends the sum.
	Total int64 `json:"total"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		HelpfulCount    respjson.Field
		NotHelpfulCount respjson.Field
		Total           respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XCommunityNoteListContributorDeletedNotesResponseNoteContributorRatingsCountSuccessful) RawJSON() string {
	return r.JSON.raw
}
func (r *XCommunityNoteListContributorDeletedNotesResponseNoteContributorRatingsCountSuccessful) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// Ratings that went against the rating X gave the note.
type XCommunityNoteListContributorDeletedNotesResponseNoteContributorRatingsCountUnsuccessful struct {
	// Ratings that called a note helpful.
	HelpfulCount int64 `json:"helpfulCount"`
	// Ratings that called a note not helpful.
	NotHelpfulCount int64 `json:"notHelpfulCount"`
	// All ratings in the group, when X sends the sum.
	Total int64 `json:"total"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		HelpfulCount    respjson.Field
		NotHelpfulCount respjson.Field
		Total           respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XCommunityNoteListContributorDeletedNotesResponseNoteContributorRatingsCountUnsuccessful) RawJSON() string {
	return r.JSON.raw
}
func (r *XCommunityNoteListContributorDeletedNotesResponseNoteContributorRatingsCountUnsuccessful) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// A link, tag or account inside a note's text.
type XCommunityNoteListContributorDeletedNotesResponseNoteEntity struct {
	// Where it starts in the text.
	FromIndex int64 `json:"fromIndex"`
	// What it points at.
	Ref XCommunityNoteListContributorDeletedNotesResponseNoteEntityRef `json:"ref"`
	// Where it ends in the text.
	ToIndex int64 `json:"toIndex"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FromIndex   respjson.Field
		Ref         respjson.Field
		ToIndex     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XCommunityNoteListContributorDeletedNotesResponseNoteEntity) RawJSON() string {
	return r.JSON.raw
}
func (r *XCommunityNoteListContributorDeletedNotesResponseNoteEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// What it points at.
type XCommunityNoteListContributorDeletedNotesResponseNoteEntityRef struct {
	// A mentioned account's username.
	ScreenName string `json:"screenName"`
	// A hashtag's or cashtag's text.
	Text string `json:"text"`
	// X's kind, such as TimelineUrl.
	Type string `json:"type"`
	// The link.
	URL string `json:"url"`
	// How X opens the link, such as ExternalUrl.
	URLType string `json:"urlType"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ScreenName  respjson.Field
		Text        respjson.Field
		Type        respjson.Field
		URL         respjson.Field
		URLType     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XCommunityNoteListContributorDeletedNotesResponseNoteEntityRef) RawJSON() string {
	return r.JSON.raw
}
func (r *XCommunityNoteListContributorDeletedNotesResponseNoteEntityRef) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// Why X shows no post, when X says.
type XCommunityNoteListContributorDeletedNotesResponseNotePostUnavailable struct {
	// Why X shows no post, as X names it.
	Reason string `json:"reason" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Reason      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XCommunityNoteListContributorDeletedNotesResponseNotePostUnavailable) RawJSON() string {
	return r.JSON.raw
}
func (r *XCommunityNoteListContributorDeletedNotesResponseNotePostUnavailable) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// 1 page of Community Notes. Follow next_cursor while the response reports more
// pages.
type XCommunityNoteListContributorNotesResponse struct {
	HasNextPage bool                                             `json:"has_next_page" api:"required"`
	NextCursor  string                                           `json:"next_cursor" api:"required"`
	Notes       []XCommunityNoteListContributorNotesResponseNote `json:"notes" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		HasNextPage respjson.Field
		NextCursor  respjson.Field
		Notes       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XCommunityNoteListContributorNotesResponse) RawJSON() string { return r.JSON.raw }
func (r *XCommunityNoteListContributorNotesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// A Community Note as its writer's page lists it, with the post it is on. A
// deleted note has no text, classification or post. A field X leaves out is left
// out.
type XCommunityNoteListContributorNotesResponseNote struct {
	// Note ID.
	ID string `json:"id" api:"required"`
	// Whether the note scores above X's bar for a possible helpful rating.
	AbovePcrhThreshold bool `json:"abovePcrhThreshold"`
	// Whether the rating is under appeal.
	AppealStatus string `json:"appealStatus"`
	// What the writer says of the post, such as MisinformedOrPotentiallyMisleading or
	// NotMisleading.
	Classification string `json:"classification"`
	// The note's writer, with the counts X sends beside a note. GET
	// /x/community-notes/contributors/{alias} serves every count.
	Contributor XCommunityNoteListContributorNotesResponseNoteContributor `json:"contributor"`
	// When the contributor wrote the note.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// The X scoring model that set the rating.
	DecidedBy string `json:"decidedBy"`
	// Whether the writer deleted the note.
	Deleted bool `json:"deleted"`
	// The link a link note is on, as X shows it.
	DisplayURL string `json:"displayUrl"`
	// Links, tags & accounts inside the text, by character range.
	Entities []XCommunityNoteListContributorNotesResponseNoteEntity `json:"entities"`
	// Whether X marks that scoring model as fully visible.
	FullyVisibleModel bool `json:"fullyVisibleModel"`
	// Why raters called the note helpful.
	HelpfulTags []string `json:"helpfulTags"`
	// Views of the note, when X shares them.
	ImpressionCount int64 `json:"impressionCount"`
	// Whether the note came through X's API.
	IsAPIAuthor bool `json:"isApiAuthor"`
	// Whether the note is about the post's media. X can show such a note on other
	// posts.
	IsMediaNote bool `json:"isMediaNote"`
	// Whether X offers a translation of the note.
	IsTranslatable bool `json:"isTranslatable"`
	// Whether the note is on a link.
	IsURLNote bool `json:"isUrlNote"`
	// Language code of the note.
	Language string `json:"language"`
	// Other posts X matched a media note to.
	MediaNoteMatchCount int64 `json:"mediaNoteMatchCount"`
	// Why raters called the note not helpful.
	NotHelpfulTags []string `json:"notHelpfulTags"`
	// The post the note is on. Null when X no longer shows it.
	Post shared.SearchTweet `json:"post" api:"nullable"`
	// ID of the post the note is on.
	PostID string `json:"postId"`
	// Why X shows no post, when X says.
	PostUnavailable XCommunityNoteListContributorNotesResponseNotePostUnavailable `json:"postUnavailable"`
	// X's current rating, such as NeedsMoreRatings, CurrentlyRatedHelpful or
	// CurrentlyRatedNotHelpful. x.com shows only helpful notes under posts.
	RatingStatus string `json:"ratingStatus"`
	// What the note says.
	Text string `json:"text"`
	// Milliseconds from the note's creation to its helpful rating.
	TimeToHelpfulFromNoteMs int64 `json:"timeToHelpfulFromNoteMs"`
	// Milliseconds from the post's creation to the note's helpful rating.
	TimeToHelpfulFromPostMs int64 `json:"timeToHelpfulFromPostMs"`
	// X's notice in place of a deleted note.
	Tombstone string `json:"tombstone"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                      respjson.Field
		AbovePcrhThreshold      respjson.Field
		AppealStatus            respjson.Field
		Classification          respjson.Field
		Contributor             respjson.Field
		CreatedAt               respjson.Field
		DecidedBy               respjson.Field
		Deleted                 respjson.Field
		DisplayURL              respjson.Field
		Entities                respjson.Field
		FullyVisibleModel       respjson.Field
		HelpfulTags             respjson.Field
		ImpressionCount         respjson.Field
		IsAPIAuthor             respjson.Field
		IsMediaNote             respjson.Field
		IsTranslatable          respjson.Field
		IsURLNote               respjson.Field
		Language                respjson.Field
		MediaNoteMatchCount     respjson.Field
		NotHelpfulTags          respjson.Field
		Post                    respjson.Field
		PostID                  respjson.Field
		PostUnavailable         respjson.Field
		RatingStatus            respjson.Field
		Text                    respjson.Field
		TimeToHelpfulFromNoteMs respjson.Field
		TimeToHelpfulFromPostMs respjson.Field
		Tombstone               respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XCommunityNoteListContributorNotesResponseNote) RawJSON() string { return r.JSON.raw }
func (r *XCommunityNoteListContributorNotesResponseNote) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// The note's writer, with the counts X sends beside a note. GET
// /x/community-notes/contributors/{alias} serves every count.
type XCommunityNoteListContributorNotesResponseNoteContributor struct {
	// The alias X gives the contributor.
	Alias string `json:"alias" api:"required"`
	// Notes the contributor deleted.
	DeletedNotesCount int64 `json:"deletedNotesCount"`
	// Whether the contributor writes notes through X's API.
	IsAPIContributor bool `json:"isApiContributor"`
	// Whether X marks the contributor as a top writer.
	IsTopWriter bool `json:"isTopWriter"`
	// How many notes a contributor wrote, by X's current rating.
	NotesCount XCommunityNoteListContributorNotesResponseNoteContributorNotesCount `json:"notesCount"`
	// How a contributor's ratings of other notes went.
	RatingsCount XCommunityNoteListContributorNotesResponseNoteContributorRatingsCount `json:"ratingsCount"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Alias             respjson.Field
		DeletedNotesCount respjson.Field
		IsAPIContributor  respjson.Field
		IsTopWriter       respjson.Field
		NotesCount        respjson.Field
		RatingsCount      respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XCommunityNoteListContributorNotesResponseNoteContributor) RawJSON() string {
	return r.JSON.raw
}
func (r *XCommunityNoteListContributorNotesResponseNoteContributor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// How many notes a contributor wrote, by X's current rating.
type XCommunityNoteListContributorNotesResponseNoteContributorNotesCount struct {
	// Notes that need more ratings.
	AwaitingMoreRatings int64 `json:"awaitingMoreRatings"`
	// Notes X rates helpful now.
	CurrentlyRatedHelpful int64 `json:"currentlyRatedHelpful"`
	// Notes X rates not helpful now.
	CurrentlyRatedNotHelpful int64 `json:"currentlyRatedNotHelpful"`
	// When X last counted the notes.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AwaitingMoreRatings      respjson.Field
		CurrentlyRatedHelpful    respjson.Field
		CurrentlyRatedNotHelpful respjson.Field
		UpdatedAt                respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XCommunityNoteListContributorNotesResponseNoteContributorNotesCount) RawJSON() string {
	return r.JSON.raw
}
func (r *XCommunityNoteListContributorNotesResponseNoteContributorNotesCount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// How a contributor's ratings of other notes went.
type XCommunityNoteListContributorNotesResponseNoteContributorRatingsCount struct {
	// Ratings on notes that need more ratings.
	AwaitingMoreRatings int64 `json:"awaitingMoreRatings"`
	// Ratings made after X rated the note.
	RatedAfterDecision int64 `json:"ratedAfterDecision"`
	// Ratings that matched the rating X gave the note.
	Successful XCommunityNoteListContributorNotesResponseNoteContributorRatingsCountSuccessful `json:"successful"`
	// Ratings that went against the rating X gave the note.
	Unsuccessful XCommunityNoteListContributorNotesResponseNoteContributorRatingsCountUnsuccessful `json:"unsuccessful"`
	// When X last counted the ratings.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AwaitingMoreRatings respjson.Field
		RatedAfterDecision  respjson.Field
		Successful          respjson.Field
		Unsuccessful        respjson.Field
		UpdatedAt           respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XCommunityNoteListContributorNotesResponseNoteContributorRatingsCount) RawJSON() string {
	return r.JSON.raw
}
func (r *XCommunityNoteListContributorNotesResponseNoteContributorRatingsCount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// Ratings that matched the rating X gave the note.
type XCommunityNoteListContributorNotesResponseNoteContributorRatingsCountSuccessful struct {
	// Ratings that called a note helpful.
	HelpfulCount int64 `json:"helpfulCount"`
	// Ratings that called a note not helpful.
	NotHelpfulCount int64 `json:"notHelpfulCount"`
	// All ratings in the group, when X sends the sum.
	Total int64 `json:"total"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		HelpfulCount    respjson.Field
		NotHelpfulCount respjson.Field
		Total           respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XCommunityNoteListContributorNotesResponseNoteContributorRatingsCountSuccessful) RawJSON() string {
	return r.JSON.raw
}
func (r *XCommunityNoteListContributorNotesResponseNoteContributorRatingsCountSuccessful) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// Ratings that went against the rating X gave the note.
type XCommunityNoteListContributorNotesResponseNoteContributorRatingsCountUnsuccessful struct {
	// Ratings that called a note helpful.
	HelpfulCount int64 `json:"helpfulCount"`
	// Ratings that called a note not helpful.
	NotHelpfulCount int64 `json:"notHelpfulCount"`
	// All ratings in the group, when X sends the sum.
	Total int64 `json:"total"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		HelpfulCount    respjson.Field
		NotHelpfulCount respjson.Field
		Total           respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XCommunityNoteListContributorNotesResponseNoteContributorRatingsCountUnsuccessful) RawJSON() string {
	return r.JSON.raw
}
func (r *XCommunityNoteListContributorNotesResponseNoteContributorRatingsCountUnsuccessful) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// A link, tag or account inside a note's text.
type XCommunityNoteListContributorNotesResponseNoteEntity struct {
	// Where it starts in the text.
	FromIndex int64 `json:"fromIndex"`
	// What it points at.
	Ref XCommunityNoteListContributorNotesResponseNoteEntityRef `json:"ref"`
	// Where it ends in the text.
	ToIndex int64 `json:"toIndex"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FromIndex   respjson.Field
		Ref         respjson.Field
		ToIndex     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XCommunityNoteListContributorNotesResponseNoteEntity) RawJSON() string { return r.JSON.raw }
func (r *XCommunityNoteListContributorNotesResponseNoteEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// What it points at.
type XCommunityNoteListContributorNotesResponseNoteEntityRef struct {
	// A mentioned account's username.
	ScreenName string `json:"screenName"`
	// A hashtag's or cashtag's text.
	Text string `json:"text"`
	// X's kind, such as TimelineUrl.
	Type string `json:"type"`
	// The link.
	URL string `json:"url"`
	// How X opens the link, such as ExternalUrl.
	URLType string `json:"urlType"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ScreenName  respjson.Field
		Text        respjson.Field
		Type        respjson.Field
		URL         respjson.Field
		URLType     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XCommunityNoteListContributorNotesResponseNoteEntityRef) RawJSON() string { return r.JSON.raw }
func (r *XCommunityNoteListContributorNotesResponseNoteEntityRef) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// Why X shows no post, when X says.
type XCommunityNoteListContributorNotesResponseNotePostUnavailable struct {
	// Why X shows no post, as X names it.
	Reason string `json:"reason" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Reason      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XCommunityNoteListContributorNotesResponseNotePostUnavailable) RawJSON() string {
	return r.JSON.raw
}
func (r *XCommunityNoteListContributorNotesResponseNotePostUnavailable) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type XCommunityNoteGetContributorParams struct {
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

// URLQuery serializes [XCommunityNoteGetContributorParams]'s query parameters as `url.Values`.
func (r XCommunityNoteGetContributorParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type XCommunityNoteListContributorDeletedNotesParams struct {
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
	// Notes per page (1-10, default 10). A low credit balance can return fewer. Follow
	// next_cursor for more pages. Also read as limit, count, max_results, maxItems,
	// max_items & per_page.
	PageSize param.Opt[int64] `query:"pageSize,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [XCommunityNoteListContributorDeletedNotesParams]'s query parameters as `url.Values`.
func (r XCommunityNoteListContributorDeletedNotesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type XCommunityNoteListContributorNotesParams struct {
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
	// Notes per page (1-10, default 10). A low credit balance can return fewer. Follow
	// next_cursor for more pages. Also read as limit, count, max_results, maxItems,
	// max_items & per_page.
	PageSize param.Opt[int64] `query:"pageSize,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [XCommunityNoteListContributorNotesParams]'s query parameters as `url.Values`.
func (r XCommunityNoteListContributorNotesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
