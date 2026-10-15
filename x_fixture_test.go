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

func TestXAutocompleteSearchFixture(t *testing.T) {
	client, request := fixtureClient(t, "SearchSuggestions.json")
	res, err := client.X.AutocompleteSearch(context.Background(), xtwitterscraper.XAutocompleteSearchParams{
		Q:      "nasa",
		Fields: xtwitterscraper.String("id,text,author.username"),
		Limit:  xtwitterscraper.Int(5),
		Omit:   xtwitterscraper.String("entities,media"),
		Types:  xtwitterscraper.String("users,hashtags"),
	})
	if err != nil {
		t.Fatal(err)
	}
	request.expect(t, http.MethodGet, "/api/v1/x/search/autocomplete", "fields=id%2Ctext%2Cauthor.username&limit=5&omit=entities%2Cmedia&q=nasa&types=users%2Chashtags")
	expectDecoded(t, res)
	if res.Query != "nasa" ||
		len(res.Users) != 1 ||
		len(res.Topics) != 1 ||
		len(res.Hashtags) != 0 ||
		len(res.Cashtags) != 0 ||
		res.Count != 2 ||
		res.Users[0].ID != "11348282" ||
		res.Users[0].Username != "NASA" ||
		res.Users[0].Name != "NASA" ||
		!res.Users[0].Verified ||
		!res.Users[0].IsBlueVerified ||
		res.Users[0].IsVerified ||
		res.Users[0].VerifiedType != "Government" ||
		res.Users[0].Protected ||
		res.Users[0].Location != "Pale Blue Dot" ||
		res.Users[0].ProfilePicture != "https://pbs.twimg.com/profile_images/1321163587679784960/0ZxKlEKB_normal.jpg" ||
		len(res.Users[0].Badges) != 0 ||
		res.Users[0].IsPersonaMediaGenable ||
		res.Users[0].Score != 0 ||
		len(res.Users[0].Tokens) != 0 ||
		res.Users[0].Inline ||
		res.Topics[0].Topic != "nasa" ||
		res.Topics[0].Score != 0 ||
		len(res.Topics[0].Tokens) != 0 {
		t.Fatalf("decoded %s", res.RawJSON())
	}
}

func TestXGetHashflagsFixture(t *testing.T) {
	client, request := fixtureClient(t, "getXHashflags.json")
	res, err := client.X.GetHashflags(context.Background(), xtwitterscraper.XGetHashflagsParams{
		ActiveOnly: xtwitterscraper.Bool(true),
		Fields:     xtwitterscraper.String("id,text,author.username"),
		Limit:      xtwitterscraper.Int(1),
		Omit:       xtwitterscraper.String("entities,media"),
		Q:          xtwitterscraper.String("nasa"),
	})
	if err != nil {
		t.Fatal(err)
	}
	request.expect(t, http.MethodGet, "/api/v1/x/hashflags", "activeOnly=true&fields=id%2Ctext%2Cauthor.username&limit=1&omit=entities%2Cmedia&q=nasa")
	expectDecoded(t, res)
	if len(res.Hashflags) != 1 ||
		res.Count != 1 ||
		res.Total != 1 ||
		res.Message != "No hashflag runs now. Send activeOnly=false to list the rest." ||
		res.Hashflags[0].Hashtag != "AION2" ||
		res.Hashflags[0].AssetURL != "https://abs.twimg.com/hashflags/BF-16692_AIon2_jp_Hashmoji_v3/BF-16692_AIon2_jp_Hashmoji_v3.png" ||
		!res.Hashflags[0].StartsAt.Equal(time.Date(2026, time.October, 2, 10, 0, 0, 0, time.UTC)) ||
		!res.Hashflags[0].EndsAt.Equal(time.Date(2027, time.January, 3, 14, 59, 0, 0, time.UTC)) ||
		res.Hashflags[0].IsHashfettiEnabled ||
		len(res.Hashflags[0].Animations) != 0 {
		t.Fatalf("decoded %s", res.RawJSON())
	}
}

func TestXGetNewsStoryVersionsFixture(t *testing.T) {
	client, request := fixtureClient(t, "getXNewsStoryVersions.json")
	res, err := client.X.GetNewsStoryVersions(context.Background(), "2110209733533598095", xtwitterscraper.XGetNewsStoryVersionsParams{
		Fields: xtwitterscraper.String("id,text,author.username"),
		Omit:   xtwitterscraper.String("entities,media"),
	})
	if err != nil {
		t.Fatal(err)
	}
	request.expect(t, http.MethodGet, "/api/v1/x/news/2110209733533598095/versions", "fields=id%2Ctext%2Cauthor.username&omit=entities%2Cmedia")
	expectDecoded(t, res)
	if res.ID != "2110209733533598095" ||
		res.Count != 1 ||
		len(res.Versions) != 1 ||
		res.Versions[0].ID != "2110211669641105725" ||
		!res.Versions[0].CreatedAt.Equal(time.Date(2026, time.October, 14, 3, 30, 42, 536000000, time.UTC)) ||
		res.Versions[0].Headline != "Elon Musk Pushes Starlink for India's Remote Areas Amid Regulatory Hurdles" {
		t.Fatalf("decoded %s", res.RawJSON())
	}
}

func TestXGetTrendLocationsFixture(t *testing.T) {
	client, request := fixtureClient(t, "getXTrendLocations.json")
	res, err := client.X.GetTrendLocations(context.Background(), xtwitterscraper.XGetTrendLocationsParams{
		Country:   xtwitterscraper.String("US"),
		Fields:    xtwitterscraper.String("id,text,author.username"),
		Limit:     xtwitterscraper.Int(1),
		Omit:      xtwitterscraper.String("entities,media"),
		PlaceType: xtwitterscraper.XGetTrendLocationsParamsPlaceTypeTown,
		Q:         xtwitterscraper.String("q"),
	})
	if err != nil {
		t.Fatal(err)
	}
	request.expect(t, http.MethodGet, "/api/v1/x/trends/locations", "country=US&fields=id%2Ctext%2Cauthor.username&limit=1&omit=entities%2Cmedia&placeType=town&q=q")
	expectDecoded(t, res)
	if len(res.Locations) != 1 ||
		res.Count != 1 ||
		res.Total != 1 ||
		res.Message != "Unknown placeType. Use country, town, worldwide or unknown." ||
		res.Locations[0].Woeid != 2344116 ||
		res.Locations[0].Name != "Istanbul" ||
		res.Locations[0].Country != "Turkey" ||
		res.Locations[0].CountryCode != "TR" ||
		res.Locations[0].PlaceType != "Town" ||
		res.Locations[0].PlaceTypeCode != 7 ||
		res.Locations[0].ParentWoeid != 23424969 ||
		res.Locations[0].URL != "http://where.yahooapis.com/v1/place/2344116" {
		t.Fatalf("decoded %s", res.RawJSON())
	}
}

func TestXGetTrendsHistoryFixture(t *testing.T) {
	client, request := fixtureClient(t, "getXTrendsHistory.json")
	res, err := client.X.GetTrendsHistory(context.Background(), xtwitterscraper.XGetTrendsHistoryParams{
		Count:    xtwitterscraper.Int(1),
		Country:  xtwitterscraper.String("country"),
		Cursor:   xtwitterscraper.String("2026-10-14T00:00:00.000Z~1791979200000"),
		Every:    xtwitterscraper.String("1h"),
		Fields:   xtwitterscraper.String("id,text,author.username"),
		From:     xtwitterscraper.String("2026-10-13"),
		Limit:    xtwitterscraper.Int(1),
		Location: xtwitterscraper.String("location"),
		Omit:     xtwitterscraper.String("entities,media"),
		To:       xtwitterscraper.String("2026-10-14"),
		Woeid:    xtwitterscraper.Int(1),
	})
	if err != nil {
		t.Fatal(err)
	}
	request.expect(t, http.MethodGet, "/api/v1/x/trends/history", "count=1&country=country&cursor=2026-10-14T00%3A00%3A00.000Z~1791979200000&every=1h&fields=id%2Ctext%2Cauthor.username&from=2026-10-13&limit=1&location=location&omit=entities%2Cmedia&to=2026-10-14&woeid=1")
	expectDecoded(t, res)
	if !res.FromTime.Equal(time.Date(2026, time.October, 13, 0, 0, 0, 0, time.UTC)) ||
		!res.ToTime.Equal(time.Date(2026, time.October, 14, 0, 0, 0, 0, time.UTC)) ||
		len(res.Snapshots) != 1 ||
		res.Count != 1 ||
		res.Woeid != 1 ||
		res.HasNextPage ||
		res.NextCursor != "" ||
		!res.Snapshots[0].CapturedAt.Equal(time.Date(2026, time.October, 13, 14, 0, 26, 0, time.UTC)) ||
		res.Snapshots[0].Count != 1 ||
		len(res.Snapshots[0].Trends) != 1 ||
		res.Snapshots[0].Trends[0].Name != "#AI" ||
		res.Snapshots[0].Trends[0].Description != "Artificial intelligence discussions" ||
		res.Snapshots[0].Trends[0].Query != "%23AI" ||
		res.Snapshots[0].Trends[0].Rank != 1 ||
		res.Snapshots[0].Trends[0].TweetVolume != 250000 ||
		res.Snapshots[0].Trends[0].URL != "https://x.com/search?q=%23AI" {
		t.Fatalf("decoded %s", res.RawJSON())
	}
}

func TestXSearchPlacesFixture(t *testing.T) {
	client, request := fixtureClient(t, "searchXPlaces.json")
	res, err := client.X.SearchPlaces(context.Background(), xtwitterscraper.XSearchPlacesParams{
		Q:      "paris",
		Fields: xtwitterscraper.String("id,text,author.username"),
		Limit:  xtwitterscraper.Int(3),
		Omit:   xtwitterscraper.String("entities,media"),
	})
	if err != nil {
		t.Fatal(err)
	}
	request.expect(t, http.MethodGet, "/api/v1/x/places/search", "fields=id%2Ctext%2Cauthor.username&limit=3&omit=entities%2Cmedia&q=paris")
	expectDecoded(t, res)
	if len(res.Places) != 1 ||
		res.Count != 1 ||
		res.Total != 1 ||
		res.Message != "X found no place for that text. Try a city or country name." ||
		res.Places[0].ID != "5e02a0f0d91c76d2" ||
		res.Places[0].Name != "İstanbul" ||
		res.Places[0].FullName != "İstanbul, Türkiye" ||
		res.Places[0].Country != "Turkey" ||
		res.Places[0].CountryCode != "TR" ||
		res.Places[0].PlaceType != "city" ||
		res.Places[0].URL != "https://api.twitter.com/1.1/geo/id/5e02a0f0d91c76d2.json" ||
		len(res.Places[0].Centroid) != 2 ||
		len(res.Places[0].ContainedWithin) != 0 ||
		res.Places[0].Centroid[0] != 28.84318392941324 ||
		res.Places[0].BoundingBox.Type != "Polygon" ||
		len(res.Places[0].BoundingBox.Coordinates) != 1 ||
		len(res.Places[0].BoundingBox.Coordinates[0]) != 4 ||
		len(res.Places[0].BoundingBox.Coordinates[0][0]) != 2 ||
		res.Places[0].BoundingBox.Coordinates[0][0][0] != 28.6321043 {
		t.Fatalf("decoded %s", res.RawJSON())
	}
}

func TestXGetTrendsFixture(t *testing.T) {
	client, request := fixtureClient(t, "getXTrends.json")
	res, err := client.X.GetTrends(context.Background(), xtwitterscraper.XGetTrendsParams{
		At:       xtwitterscraper.String("2026-10-13T14:00:00Z"),
		Count:    xtwitterscraper.Int(10),
		Country:  xtwitterscraper.String("TR"),
		Fields:   xtwitterscraper.String("trends,capturedAt"),
		Location: xtwitterscraper.String("Istanbul"),
		Omit:     xtwitterscraper.String("trends.url"),
		Woeid:    xtwitterscraper.Int(2344116),
	})
	if err != nil {
		t.Fatal(err)
	}
	request.expect(t, http.MethodGet, "/api/v1/x/trends", "at=2026-10-13T14%3A00%3A00Z&count=10&country=TR&fields=trends%2CcapturedAt&location=Istanbul&omit=trends.url&woeid=2344116")
	expectDecoded(t, res)
	if len(res.Trends) != 1 ||
		res.Count != 30 ||
		res.Woeid != 1 ||
		!res.CapturedAt.Equal(time.Date(2026, time.October, 13, 14, 0, 26, 0, time.UTC)) ||
		!res.RequestedAt.Equal(time.Date(2026, time.October, 13, 14, 0, 0, 0, time.UTC)) ||
		res.Location.Woeid != 2344116 ||
		res.Location.Name != "Istanbul" ||
		res.Location.Country != "Turkey" ||
		res.Location.CountryCode != "TR" ||
		res.Location.PlaceType != "Town" ||
		res.Location.PlaceTypeCode != 7 ||
		res.Location.ParentWoeid != 23424969 ||
		res.Location.URL != "http://where.yahooapis.com/v1/place/2344116" ||
		res.Trends[0].Name != "#AI" ||
		res.Trends[0].Description != "Artificial intelligence discussions" ||
		res.Trends[0].Query != "%23AI" ||
		res.Trends[0].Rank != 1 ||
		res.Trends[0].TweetVolume != 250000 ||
		res.Trends[0].URL != "https://x.com/search?q=%23AI" {
		t.Fatalf("decoded %s", res.RawJSON())
	}
}
