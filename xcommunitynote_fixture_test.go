// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

package xtwitterscraper_test

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/Xquik-dev/x-twitter-scraper-go"
)

func TestXCommunityNoteGetContributorFixture(t *testing.T) {
	client, request := fixtureClient(t, "getCommunityNotesContributor.json")
	res, err := client.X.CommunityNotes.GetContributor(context.Background(), "ardent-beach-friarbird", xtwitterscraper.XCommunityNoteGetContributorParams{
		Fields: xtwitterscraper.String("id,text,author.username"),
		Omit:   xtwitterscraper.String("entities,media"),
	})
	if err != nil {
		t.Fatal(err)
	}
	request.expect(t, http.MethodGet, "/api/v1/x/community-notes/contributors/ardent-beach-friarbird", "fields=id%2Ctext%2Cauthor.username&omit=entities%2Cmedia")
	expectDecoded(t, res)
	if res.Contributor.Alias != "ardent-beach-friarbird" ||
		res.Contributor.DeletedNotesCount != 0 ||
		res.Contributor.IsAPIContributor ||
		res.Contributor.IsTopWriter ||
		res.Contributor.NotesCount.AwaitingMoreRatings != 18 ||
		res.Contributor.NotesCount.CurrentlyRatedHelpful != 1 ||
		res.Contributor.NotesCount.CurrentlyRatedNotHelpful != 2 ||
		!res.Contributor.NotesCount.UpdatedAt.Equal(time.Date(2026, time.October, 12, 22, 18, 40, 96000000, time.UTC)) ||
		res.Contributor.RatingsCount.AwaitingMoreRatings != 703 ||
		res.Contributor.RatingsCount.RatedAfterDecision != 154 ||
		!res.Contributor.RatingsCount.UpdatedAt.Equal(time.Date(2026, time.October, 12, 22, 50, 13, 228000000, time.UTC)) ||
		res.Contributor.RatingsCount.Successful.HelpfulCount != 27 ||
		res.Contributor.RatingsCount.Successful.NotHelpfulCount != 22 ||
		res.Contributor.RatingsCount.Successful.Total != 49 ||
		res.Contributor.RatingsCount.Unsuccessful.HelpfulCount != 15 ||
		res.Contributor.RatingsCount.Unsuccessful.NotHelpfulCount != 4 {
		t.Fatalf("decoded %s", res.RawJSON())
	}
}

func TestXCommunityNoteListContributorDeletedNotesFixture(t *testing.T) {
	client, request := fixtureClient(t, "listCommunityNotesContributorDeletedNotes.json")
	res, err := client.X.CommunityNotes.ListContributorDeletedNotes(context.Background(), "ardent-beach-friarbird", xtwitterscraper.XCommunityNoteListContributorDeletedNotesParams{
		Cursor:   xtwitterscraper.String("cursor"),
		Fields:   xtwitterscraper.String("id,text,author.username"),
		Omit:     xtwitterscraper.String("entities,media"),
		PageSize: xtwitterscraper.Int(10),
	})
	if err != nil {
		t.Fatal(err)
	}
	request.expect(t, http.MethodGet, "/api/v1/x/community-notes/contributors/ardent-beach-friarbird/deleted-notes", "cursor=cursor&fields=id%2Ctext%2Cauthor.username&omit=entities%2Cmedia&pageSize=10")
	expectDecoded(t, res)
	if len(res.Notes) != 1 ||
		res.HasNextPage ||
		res.NextCursor != "" ||
		res.Notes[0].ID != "2078950499982696756" ||
		!res.Notes[0].CreatedAt.Equal(time.Date(2026, time.July, 19, 21, 9, 58, 885000000, time.UTC)) ||
		res.Notes[0].RatingStatus != "CurrentlyRatedNotHelpful" ||
		len(res.Notes[0].NotHelpfulTags) != 2 ||
		res.Notes[0].DecidedBy != "CoreModel (v1.1)" ||
		!res.Notes[0].FullyVisibleModel ||
		res.Notes[0].AbovePcrhThreshold ||
		res.Notes[0].AppealStatus != "NotAppealed" ||
		!res.Notes[0].Deleted ||
		res.Notes[0].Tombstone != "This note has been deleted" ||
		res.Notes[0].IsAPIAuthor ||
		res.Notes[0].IsMediaNote ||
		res.Notes[0].MediaNoteMatchCount != 0 ||
		res.Notes[0].IsURLNote ||
		res.Notes[0].NotHelpfulTags[0] != "MissingKeyPoints" {
		t.Fatalf("decoded %s", res.RawJSON())
	}
}

func TestXCommunityNoteListContributorNotesFixture(t *testing.T) {
	client, request := fixtureClient(t, "listCommunityNotesContributorNotes.json")
	res, err := client.X.CommunityNotes.ListContributorNotes(context.Background(), "ardent-beach-friarbird", xtwitterscraper.XCommunityNoteListContributorNotesParams{
		Cursor:   xtwitterscraper.String("cursor"),
		Fields:   xtwitterscraper.String("id,text,author.username"),
		Omit:     xtwitterscraper.String("entities,media"),
		PageSize: xtwitterscraper.Int(5),
	})
	if err != nil {
		t.Fatal(err)
	}
	request.expect(t, http.MethodGet, "/api/v1/x/community-notes/contributors/ardent-beach-friarbird/notes", "cursor=cursor&fields=id%2Ctext%2Cauthor.username&omit=entities%2Cmedia&pageSize=5")
	expectDecoded(t, res)
	if res.Notes[0].Post.ID != "2026359315657805939" ||
		len(res.Notes) != 1 ||
		!res.HasNextPage ||
		res.NextCursor != "-1767098234260" ||
		res.Notes[0].ID != "2026375778254868976" ||
		!res.Notes[0].CreatedAt.Equal(time.Date(2026, time.February, 24, 19, 16, 48, 435000000, time.UTC)) ||
		res.Notes[0].Language != "en" ||
		res.Notes[0].Classification != "MisinformedOrPotentiallyMisleading" ||
		len(res.Notes[0].Entities) != 1 ||
		res.Notes[0].RatingStatus != "CurrentlyRatedHelpful" ||
		len(res.Notes[0].HelpfulTags) != 2 ||
		res.Notes[0].DecidedBy != "CoreModel (v1.1)" ||
		!res.Notes[0].FullyVisibleModel ||
		res.Notes[0].AbovePcrhThreshold ||
		res.Notes[0].AppealStatus != "NotAppealed" ||
		res.Notes[0].Deleted ||
		res.Notes[0].IsAPIAuthor ||
		res.Notes[0].IsMediaNote ||
		res.Notes[0].MediaNoteMatchCount != 0 ||
		res.Notes[0].IsURLNote ||
		res.Notes[0].IsTranslatable ||
		res.Notes[0].TimeToHelpfulFromNoteMs != 132890587 ||
		res.Notes[0].TimeToHelpfulFromPostMs != 136815576 ||
		res.Notes[0].PostID != "2026359315657805939" ||
		res.Notes[0].Contributor.Alias != "ardent-beach-friarbird" {
		t.Fatalf("decoded %s", res.RawJSON())
	}
}
