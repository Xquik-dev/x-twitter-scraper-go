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

func TestXListGetFixture(t *testing.T) {
	client, request := fixtureClient(t, "getList.json")
	res, err := client.X.Lists.Get(context.Background(), "4143408", xtwitterscraper.XListGetParams{
		Fields: xtwitterscraper.String("id,text,author.username"),
		Omit:   xtwitterscraper.String("entities,media"),
	})
	if err != nil {
		t.Fatal(err)
	}
	request.expect(t, http.MethodGet, "/api/v1/x/lists/4143408", "fields=id%2Ctext%2Cauthor.username&omit=entities%2Cmedia")
	expectDecoded(t, res)
	if res.List.ID != "4143408" ||
		res.List.Name != "NASA Centers" ||
		res.List.Description != "NASA has 10 centers" ||
		res.List.MemberCount != 11 ||
		res.List.SubscriberCount != 859 ||
		res.List.URL != "https://x.com/i/lists/4143408" ||
		!res.List.CreatedAt.Equal(time.Date(2009, time.December, 2, 22, 41, 2, 0, time.UTC)) ||
		res.List.BannerURL != "https://pbs.twimg.com/media/example.png" ||
		res.List.MembersContext != "11 members" ||
		res.List.Owner.ID != "11348282" ||
		res.List.Owner.Username != "NASA" ||
		res.List.Owner.Name != "NASA" {
		t.Fatalf("decoded %s", res.RawJSON())
	}
}

func TestXListGetSearchFixture(t *testing.T) {
	client, request := fixtureClient(t, "PaginatedListsWithList.json")
	res, err := client.X.Lists.GetSearch(context.Background(), xtwitterscraper.XListGetSearchParams{
		Q:        "nasa",
		Cursor:   xtwitterscraper.String("cursor"),
		Fields:   xtwitterscraper.String("id,text,author.username"),
		Omit:     xtwitterscraper.String("entities,media"),
		PageSize: xtwitterscraper.Int(20),
	})
	if err != nil {
		t.Fatal(err)
	}
	request.expect(t, http.MethodGet, "/api/v1/x/lists/search", "cursor=cursor&fields=id%2Ctext%2Cauthor.username&omit=entities%2Cmedia&pageSize=20&q=nasa")
	expectDecoded(t, res)
	if len(res.Lists) != 1 ||
		!res.HasNextPage ||
		res.NextCursor != "DAACCgACGRElMJcAAA" ||
		res.Lists[0].ID != "4143408" ||
		res.Lists[0].Name != "NASA Centers" ||
		res.Lists[0].Description != "NASA has 10 centers http://www.nasa.gov/about/sites" ||
		res.Lists[0].MemberCount != 11 ||
		res.Lists[0].SubscriberCount != 859 ||
		res.Lists[0].URL != "https://x.com/i/lists/4143408" ||
		!res.Lists[0].CreatedAt.Equal(time.Date(2009, time.December, 2, 22, 41, 2, 0, time.UTC)) ||
		res.Lists[0].MembersContext != "11 members" ||
		res.Lists[0].Owner.ID != "11348282" ||
		res.Lists[0].Owner.Username != "NASA" ||
		res.Lists[0].Owner.Name != "NASA" ||
		res.Lists[0].Owner.Followers != 92379503 ||
		res.Lists[0].Owner.Following != 115 ||
		!res.Lists[0].Owner.Verified ||
		!res.Lists[0].Owner.IsBlueVerified ||
		!res.Lists[0].Owner.IsVerified ||
		res.Lists[0].Owner.ProfilePicture != "https://pbs.twimg.com/profile_images/1321163587679784960/0ZxKlEKB_normal.jpg" ||
		res.Lists[0].Owner.CoverPicture != "https://pbs.twimg.com/profile_banners/11348282/1788992290" ||
		res.Lists[0].Owner.ProfileBannerURL != "https://pbs.twimg.com/profile_banners/11348282/1788992290" ||
		res.Lists[0].Owner.Description != "Making the seemingly impossible, possible. ✨" ||
		res.Lists[0].Owner.Location != "Pale Blue Dot" {
		t.Fatalf("decoded %s", res.RawJSON())
	}
}
