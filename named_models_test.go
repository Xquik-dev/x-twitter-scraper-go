// Copyright the Xquik contributors.
// SPDX-License-Identifier: Apache-2.0

package xtwitterscraper_test

import (
	"context"
	"testing"

	"github.com/Xquik-dev/x-twitter-scraper-go"
)

func TestJobOperationsReturnNamedModel(t *testing.T) {
	getClient, _ := fixtureClient(t, "getJob.json")
	get, err := getClient.X.Jobs.Get(context.Background(), "1899202395210649861", xtwitterscraper.XJobGetParams{})
	if err != nil {
		t.Fatal(err)
	}
	searchClient, _ := fixtureClient(t, "searchJobs.json")
	search, err := searchClient.X.Jobs.Search(context.Background(), xtwitterscraper.XJobSearchParams{Q: "engineer"})
	if err != nil {
		t.Fatal(err)
	}
	if len(search.Jobs) != 1 {
		t.Fatalf("jobs = %d, want 1", len(search.Jobs))
	}
	for _, job := range []xtwitterscraper.Job{get.Job, search.Jobs[0]} {
		expectDecoded(t, job)
		if job.ID != "1899202395210649861" || job.Title != "Engineer" || !job.JSON.ID.Valid() || job.RawJSON() == "" {
			t.Fatalf("decoded %s", job.RawJSON())
		}
	}
}

func TestSpaceOperationsReturnNamedModel(t *testing.T) {
	getClient, _ := fixtureClient(t, "getSpace.json")
	get, err := getClient.X.Spaces.Get(context.Background(), "1OGwblLjnQMKB", xtwitterscraper.XSpaceGetParams{})
	if err != nil {
		t.Fatal(err)
	}
	searchClient, _ := fixtureClient(t, "searchSpaces.json")
	search, err := searchClient.X.Spaces.Search(context.Background(), xtwitterscraper.XSpaceSearchParams{Q: "from:solana"})
	if err != nil {
		t.Fatal(err)
	}
	if len(search.Spaces) != 1 {
		t.Fatalf("spaces = %d, want 1", len(search.Spaces))
	}
	for _, space := range []xtwitterscraper.Space{get.Space, search.Spaces[0]} {
		expectDecoded(t, space)
		if space.ID != "1OGwblLjnQMKB" || space.State != "Ended" || !space.JSON.ID.Valid() || space.RawJSON() == "" {
			t.Fatalf("decoded %s", space.RawJSON())
		}
		participants := append([]xtwitterscraper.SpaceParticipant{}, space.Admins...)
		participants = append(participants, space.Speakers...)
		participants = append(participants, space.Listeners...)
		if len(participants) != 1 || participants[0].Username != "solana" {
			t.Fatalf("decoded participants %s", space.RawJSON())
		}
	}
}

func TestCommunityNoteOperationsReturnNamedModel(t *testing.T) {
	notesClient, _ := fixtureClient(t, "listCommunityNotesContributorNotes.json")
	notes, err := notesClient.X.CommunityNotes.ListContributorNotes(context.Background(), "ardent-beach-friarbird", xtwitterscraper.XCommunityNoteListContributorNotesParams{})
	if err != nil {
		t.Fatal(err)
	}
	deletedClient, _ := fixtureClient(t, "listCommunityNotesContributorDeletedNotes.json")
	deleted, err := deletedClient.X.CommunityNotes.ListContributorDeletedNotes(context.Background(), "ardent-beach-friarbird", xtwitterscraper.XCommunityNoteListContributorDeletedNotesParams{})
	if err != nil {
		t.Fatal(err)
	}
	pages := []*xtwitterscraper.PaginatedCommunityNotes{notes, deleted}
	for i, page := range pages {
		expectDecoded(t, page)
		if len(page.Notes) != 1 || !page.JSON.Notes.Valid() || page.RawJSON() == "" {
			t.Fatalf("decoded %s", page.RawJSON())
		}
		wantID := []string{"2026375778254868976", "2078950499982696756"}[i]
		if page.Notes[0].ID != wantID || page.Notes[0].Deleted != (i == 1) {
			t.Fatalf("decoded %s", page.RawJSON())
		}
	}
}

func TestCommunityInfoCreatorPreservesExtraFields(t *testing.T) {
	var community xtwitterscraper.XCommunityGetInfoResponseCommunity
	if err := community.UnmarshalJSON([]byte(`{"creator":{"id":"641633","username":"marckohlbrugge","newProfileField":42}}`)); err != nil {
		t.Fatal(err)
	}
	var creator xtwitterscraper.XCommunityGetInfoResponseCommunityCreator = community.Creator
	if field, ok := creator.JSON.ExtraFields["newProfileField"]; !ok || field.Raw() != "42" {
		t.Fatalf("creator extra field = %q, want 42", field.Raw())
	}
}

var (
	_ xtwitterscraper.Job                                        = xtwitterscraper.XJobGetResponseJob{}
	_ xtwitterscraper.Job                                        = xtwitterscraper.XJobSearchResponseJob{}
	_ xtwitterscraper.JobAccount                                 = xtwitterscraper.XJobGetResponseJobAccount{}
	_ xtwitterscraper.JobAccount                                 = xtwitterscraper.XJobSearchResponseJobAccount{}
	_ xtwitterscraper.JobCompany                                 = xtwitterscraper.XJobGetResponseJobCompany{}
	_ xtwitterscraper.JobCompany                                 = xtwitterscraper.XJobSearchResponseJobCompany{}
	_ xtwitterscraper.Space                                      = xtwitterscraper.XSpaceGetResponseSpace{}
	_ xtwitterscraper.SpaceParticipant                           = xtwitterscraper.XSpaceGetResponseSpaceAdmin{}
	_ xtwitterscraper.SpaceParticipantAffiliatesHighlightedLabel = xtwitterscraper.XSpaceGetResponseSpaceAdminAffiliatesHighlightedLabel{}
	_ xtwitterscraper.SpaceSharing                               = xtwitterscraper.XSpaceGetResponseSpaceSharing{}
	_ xtwitterscraper.SpaceParticipant                           = xtwitterscraper.XSpaceGetResponseSpaceListener{}
	_ xtwitterscraper.SpaceParticipantAffiliatesHighlightedLabel = xtwitterscraper.XSpaceGetResponseSpaceListenerAffiliatesHighlightedLabel{}
	_ xtwitterscraper.SpaceParticipant                           = xtwitterscraper.XSpaceGetResponseSpaceSpeaker{}
	_ xtwitterscraper.SpaceParticipantAffiliatesHighlightedLabel = xtwitterscraper.XSpaceGetResponseSpaceSpeakerAffiliatesHighlightedLabel{}
	_ xtwitterscraper.Space                                      = xtwitterscraper.XSpaceSearchResponseSpace{}
	_ xtwitterscraper.SpaceSharing                               = xtwitterscraper.XSpaceSearchResponseSpaceSharing{}
	_ xtwitterscraper.SpaceParticipant                           = xtwitterscraper.XSpaceSearchResponseSpaceAdmin{}
	_ xtwitterscraper.SpaceParticipantAffiliatesHighlightedLabel = xtwitterscraper.XSpaceSearchResponseSpaceAdminAffiliatesHighlightedLabel{}
	_ xtwitterscraper.SpaceParticipant                           = xtwitterscraper.XSpaceSearchResponseSpaceListener{}
	_ xtwitterscraper.SpaceParticipantAffiliatesHighlightedLabel = xtwitterscraper.XSpaceSearchResponseSpaceListenerAffiliatesHighlightedLabel{}
	_ xtwitterscraper.SpaceParticipant                           = xtwitterscraper.XSpaceSearchResponseSpaceSpeaker{}
	_ xtwitterscraper.SpaceParticipantAffiliatesHighlightedLabel = xtwitterscraper.XSpaceSearchResponseSpaceSpeakerAffiliatesHighlightedLabel{}
	_ xtwitterscraper.Broadcast                                  = xtwitterscraper.XBroadcastGetResponseBroadcast{}
	_ xtwitterscraper.BroadcastPeriscopeUser                     = xtwitterscraper.XBroadcastGetResponseBroadcastPeriscopeUser{}
	_ xtwitterscraper.Broadcast                                  = xtwitterscraper.XUserGetLiveResponseBroadcast{}
	_ xtwitterscraper.BroadcastPeriscopeUser                     = xtwitterscraper.XUserGetLiveResponseBroadcastPeriscopeUser{}
	_ xtwitterscraper.CommunityNotesContributor                  = xtwitterscraper.XCommunityNoteGetContributorResponseContributor{}
	_ xtwitterscraper.CommunityNotesContributorNotesCount        = xtwitterscraper.XCommunityNoteGetContributorResponseContributorNotesCount{}
	_ xtwitterscraper.CommunityNotesContributorRatingsCount      = xtwitterscraper.XCommunityNoteGetContributorResponseContributorRatingsCount{}
	_ xtwitterscraper.CommunityNotesRatingOutcome                = xtwitterscraper.XCommunityNoteGetContributorResponseContributorRatingsCountSuccessful{}
	_ xtwitterscraper.CommunityNotesRatingOutcome                = xtwitterscraper.XCommunityNoteGetContributorResponseContributorRatingsCountUnsuccessful{}
	_ xtwitterscraper.PaginatedCommunityNotes                    = xtwitterscraper.XCommunityNoteListContributorNotesResponse{}
	_ xtwitterscraper.PaginatedCommunityNotes                    = xtwitterscraper.XCommunityNoteListContributorDeletedNotesResponse{}
	_ xtwitterscraper.PaginatedCommunityNotesNote                = xtwitterscraper.XCommunityNoteListContributorNotesResponseNote{}
	_ xtwitterscraper.PaginatedCommunityNotesNote                = xtwitterscraper.XCommunityNoteListContributorDeletedNotesResponseNote{}
	_ xtwitterscraper.PaginatedCommunityNotesNoteEntity          = xtwitterscraper.XCommunityNoteListContributorNotesResponseNoteEntity{}
	_ xtwitterscraper.PaginatedCommunityNotesNoteEntity          = xtwitterscraper.XCommunityNoteListContributorDeletedNotesResponseNoteEntity{}
	_ xtwitterscraper.PaginatedCommunityNotesNoteEntityRef       = xtwitterscraper.XCommunityNoteListContributorNotesResponseNoteEntityRef{}
	_ xtwitterscraper.PaginatedCommunityNotesNoteEntityRef       = xtwitterscraper.XCommunityNoteListContributorDeletedNotesResponseNoteEntityRef{}
	_ xtwitterscraper.PaginatedCommunityNotesNotePostUnavailable = xtwitterscraper.XCommunityNoteListContributorNotesResponseNotePostUnavailable{}
	_ xtwitterscraper.PaginatedCommunityNotesNotePostUnavailable = xtwitterscraper.XCommunityNoteListContributorDeletedNotesResponseNotePostUnavailable{}
	_ xtwitterscraper.CommunityNotesContributor                  = xtwitterscraper.XCommunityNoteListContributorNotesResponseNoteContributor{}
	_ xtwitterscraper.CommunityNotesContributorNotesCount        = xtwitterscraper.XCommunityNoteListContributorNotesResponseNoteContributorNotesCount{}
	_ xtwitterscraper.CommunityNotesContributorRatingsCount      = xtwitterscraper.XCommunityNoteListContributorNotesResponseNoteContributorRatingsCount{}
	_ xtwitterscraper.CommunityNotesRatingOutcome                = xtwitterscraper.XCommunityNoteListContributorNotesResponseNoteContributorRatingsCountSuccessful{}
	_ xtwitterscraper.CommunityNotesRatingOutcome                = xtwitterscraper.XCommunityNoteListContributorNotesResponseNoteContributorRatingsCountUnsuccessful{}
	_ xtwitterscraper.CommunityNotesContributor                  = xtwitterscraper.XCommunityNoteListContributorDeletedNotesResponseNoteContributor{}
	_ xtwitterscraper.CommunityNotesContributorNotesCount        = xtwitterscraper.XCommunityNoteListContributorDeletedNotesResponseNoteContributorNotesCount{}
	_ xtwitterscraper.CommunityNotesContributorRatingsCount      = xtwitterscraper.XCommunityNoteListContributorDeletedNotesResponseNoteContributorRatingsCount{}
	_ xtwitterscraper.CommunityNotesRatingOutcome                = xtwitterscraper.XCommunityNoteListContributorDeletedNotesResponseNoteContributorRatingsCountSuccessful{}
	_ xtwitterscraper.CommunityNotesRatingOutcome                = xtwitterscraper.XCommunityNoteListContributorDeletedNotesResponseNoteContributorRatingsCountUnsuccessful{}
	_ xtwitterscraper.Community                                  = xtwitterscraper.XCommunityGetInfoResponseCommunity{}
	_ xtwitterscraper.CommunityCreatorUnion                      = xtwitterscraper.XCommunityGetInfoResponseCommunityCreator{}
	_ xtwitterscraper.CommunityPrimaryTopic                      = xtwitterscraper.XCommunityGetInfoResponseCommunityPrimaryTopic{}
	_ xtwitterscraper.CommunityRule                              = xtwitterscraper.XCommunityGetInfoResponseCommunityRule{}
)
