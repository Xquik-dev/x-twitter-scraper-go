// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

package xtwitterscraper_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/Xquik-dev/x-twitter-scraper-go"
)

func TestXUserGetFollowerAuditFixture(t *testing.T) {
	client, request := fixtureClient(t, "getUserFollowerAudit.json")
	res, err := client.X.Users.GetFollowerAudit(context.Background(), "nasa", xtwitterscraper.XUserGetFollowerAuditParams{
		Fields: xtwitterscraper.String("audit"),
		Limit:  xtwitterscraper.Int(200),
		Omit:   xtwitterscraper.String("followers"),
	})
	if err != nil {
		t.Fatal(err)
	}
	request.expect(t, http.MethodGet, "/api/v1/x/users/nasa/follower-audit", "fields=audit&limit=200&omit=followers")
	expectDecoded(t, res)
	audit := res.Audit
	if audit.Account.UserName != "NASA" || audit.Account.Followers != 92376187 || audit.Account.Sample.Posts != 20 ||
		audit.Coverage.Percent != 0.0002 || audit.Coverage.Read != 200 || audit.Coverage.Total != 92376187 ||
		audit.Quality.Grade != xtwitterscraper.FollowerAuditQualityGradeB || audit.Quality.Score != 71 ||
		audit.Verified.Count != 9 || audit.Verified.Percent != 4.5 {
		t.Fatalf("decoded %s", res.RawJSON())
	}
	if len(audit.Classes) == 0 || audit.Classes[0].Class != xtwitterscraper.FollowerClassLikelyReal || audit.Classes[0].Count != 122 || audit.Classes[0].Percent != 61 ||
		len(audit.Signals) == 0 || audit.Signals[0].ID != "defaultPicture" || audit.Signals[0].Count != 36 || audit.Signals[0].Points != 2 ||
		len(audit.ArrivalBands) == 0 || audit.ArrivalBands[0].From != 1 || audit.ArrivalBands[0].To != 100 || audit.ArrivalBands[0].Read != 100 || audit.ArrivalBands[0].Count != 13 ||
		len(audit.Cohorts) == 0 || audit.Cohorts[0].WeekStartMs != 1761523200000 || audit.Cohorts[0].LikelyFake != 6 || audit.Cohorts[0].Percent != 3.5 ||
		len(audit.Languages) == 0 || audit.Languages[0].Value != "en" || audit.Languages[0].Count != 141 ||
		len(audit.Places) == 0 || audit.Places[0].Value != "United States" || audit.Places[0].Percent != 11 {
		t.Fatalf("decoded %s", res.RawJSON())
	}
	for name, followers := range map[string][]xtwitterscraper.AuditedFollower{"followers": audit.Followers, "topFollowers": audit.TopFollowers} {
		if len(followers) != 1 {
			t.Fatalf("%s = %d, want 1", name, len(followers))
		}
		follower := followers[0]
		if follower.ID != "1852412278803021824" || follower.UserName != "orbit_watcher" || follower.Name != "Orbit Watcher" ||
			follower.Class != xtwitterscraper.FollowerClassLikelyReal || follower.Score != 1 ||
			len(follower.Signals) != 1 || follower.Signals[0] != "noBanner" ||
			follower.Followers != 312 || follower.Following != 410 || follower.Likes != 2840 || follower.Posts != 1290 ||
			follower.Position != 1 || follower.JoinedAtMs != 1561939200000 ||
			follower.PictureURL != "https://pbs.twimg.com/profile_images/1/photo.jpg" || follower.JSON.Owner.Valid() {
			t.Fatalf("%s[0] = %s", name, follower.RawJSON())
		}
	}
}

func TestXUserGetFollowerAuditNeedsAnID(t *testing.T) {
	client, _ := fixtureClient(t, "getUserFollowerAudit.json")
	if _, err := client.X.Users.GetFollowerAudit(context.Background(), "", xtwitterscraper.XUserGetFollowerAuditParams{}); err == nil || err.Error() != "missing required id parameter" {
		t.Fatalf("err = %v", err)
	}
}
