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
	"github.com/Xquik-dev/x-twitter-scraper-go/shared"
)

func TestXCommunityFindFixture(t *testing.T) {
	client, request := fixtureClient(t, "findCommunities.json")
	res, err := client.X.Communities.Find(context.Background(), xtwitterscraper.XCommunityFindParams{
		Q:        "ai",
		Cursor:   xtwitterscraper.String("cursor"),
		Fields:   xtwitterscraper.String("id,text,author.username"),
		Omit:     xtwitterscraper.String("entities,media"),
		PageSize: xtwitterscraper.Int(20),
	})
	if err != nil {
		t.Fatal(err)
	}
	request.expect(t, http.MethodGet, "/api/v1/x/communities/find", "cursor=cursor&fields=id%2Ctext%2Cauthor.username&omit=entities%2Cmedia&pageSize=20&q=ai")
	expectDecoded(t, res)
	if len(res.Communities) != 1 ||
		!res.HasNextPage ||
		res.NextCursor != "WzM1LjUxMjQ5N10=" ||
		res.Communities[0].ID != "1472105760389668865" ||
		res.Communities[0].Name != "Tech Twitter" ||
		res.Communities[0].MemberCount != 121468 ||
		res.Communities[0].IsNsfw ||
		res.Communities[0].BannerURL != "https://pbs.twimg.com/community_banner_img/example.jpg" ||
		len(res.Communities[0].MembersFacepile) != 1 ||
		res.Communities[0].PrimaryTopic.Name != "Technology" ||
		res.Communities[0].MembersFacepile[0].ID != "2864203879" ||
		res.Communities[0].MembersFacepile[0].ProfilePicture != "https://pbs.twimg.com/profile_images/example.jpg" {
		t.Fatalf("decoded %s", res.RawJSON())
	}
}

func TestXCommunityListPopularFixture(t *testing.T) {
	client, request := fixtureClient(t, "getPopularCommunities.json")
	res, err := client.X.Communities.ListPopular(context.Background(), xtwitterscraper.XCommunityListPopularParams{
		Cursor:   xtwitterscraper.String("cursor"),
		Fields:   xtwitterscraper.String("id,text,author.username"),
		Omit:     xtwitterscraper.String("entities,media"),
		PageSize: xtwitterscraper.Int(20),
		Topic:    xtwitterscraper.String("technology"),
		TopicID:  xtwitterscraper.String("topicId"),
	})
	if err != nil {
		t.Fatal(err)
	}
	request.expect(t, http.MethodGet, "/api/v1/x/communities/popular", "cursor=cursor&fields=id%2Ctext%2Cauthor.username&omit=entities%2Cmedia&pageSize=20&topic=technology&topicId=topicId")
	expectDecoded(t, res)
	if len(res.Communities) != 1 ||
		!res.HasNextPage ||
		res.NextCursor != "WzM1LjUxMjQ5N10=" ||
		res.Communities[0].ID != "1472105760389668865" ||
		res.Communities[0].Name != "Tech Twitter" ||
		res.Communities[0].MemberCount != 121468 ||
		res.Communities[0].IsNsfw ||
		res.Communities[0].BannerURL != "https://pbs.twimg.com/community_banner_img/example.jpg" ||
		len(res.Communities[0].MembersFacepile) != 1 ||
		res.Communities[0].PrimaryTopic.Name != "Technology" ||
		res.Communities[0].MembersFacepile[0].ID != "2864203879" ||
		res.Communities[0].MembersFacepile[0].ProfilePicture != "https://pbs.twimg.com/profile_images/example.jpg" {
		t.Fatalf("decoded %s", res.RawJSON())
	}
}

func TestXCommunityListSuggestedFixture(t *testing.T) {
	client, request := fixtureClient(t, "getSuggestedCommunities.json")
	res, err := client.X.Communities.ListSuggested(context.Background(), xtwitterscraper.XCommunityListSuggestedParams{
		Fields: xtwitterscraper.String("id,text,author.username"),
		Limit:  xtwitterscraper.Int(10),
		Omit:   xtwitterscraper.String("entities,media"),
	})
	if err != nil {
		t.Fatal(err)
	}
	request.expect(t, http.MethodGet, "/api/v1/x/communities/suggested", "fields=id%2Ctext%2Cauthor.username&limit=10&omit=entities%2Cmedia")
	expectDecoded(t, res)
	if len(res.Communities) != 1 ||
		res.Count != 1 ||
		res.Communities[0].ID != "1472105760389668865" ||
		res.Communities[0].Name != "Tech Twitter" ||
		res.Communities[0].MemberCount != 121468 ||
		res.Communities[0].IsNsfw ||
		res.Communities[0].BannerURL != "https://pbs.twimg.com/community_banner_img/example.jpg" ||
		len(res.Communities[0].MembersFacepile) != 1 ||
		res.Communities[0].PrimaryTopic.Name != "Technology" ||
		res.Communities[0].MembersFacepile[0].ID != "2864203879" ||
		res.Communities[0].MembersFacepile[0].ProfilePicture != "https://pbs.twimg.com/profile_images/example.jpg" {
		t.Fatalf("decoded %s", res.RawJSON())
	}
}

func TestXCommunityListTopicsFixture(t *testing.T) {
	client, request := fixtureClient(t, "getCommunityTopics.json")
	res, err := client.X.Communities.ListTopics(context.Background(), xtwitterscraper.XCommunityListTopicsParams{
		Fields: xtwitterscraper.String("id,text,author.username"),
		Limit:  xtwitterscraper.Int(5),
		Omit:   xtwitterscraper.String("entities,media"),
	})
	if err != nil {
		t.Fatal(err)
	}
	request.expect(t, http.MethodGet, "/api/v1/x/communities/topics", "fields=id%2Ctext%2Cauthor.username&limit=5&omit=entities%2Cmedia")
	expectDecoded(t, res)
	if len(res.Topics) != 1 ||
		res.Count != 1 ||
		res.Topics[0].ID != "301" ||
		res.Topics[0].Name != "Technology" ||
		len(res.Topics[0].Subtopics) != 2 ||
		res.Topics[0].Subtopics[0].ID != "302" ||
		res.Topics[0].Subtopics[0].Name != "Artificial Intelligence" {
		t.Fatalf("decoded %s", res.RawJSON())
	}
}

func TestXCommunitySearchMembersFixture(t *testing.T) {
	client, request := fixtureClient(t, "searchCommunityMembers.json")
	res, err := client.X.Communities.SearchMembers(context.Background(), "1500000000000000000", xtwitterscraper.XCommunitySearchMembersParams{
		Q:                 "marckohlbrugge",
		BioContains:       xtwitterscraper.String("bioContains"),
		BlueVerifiedOnly:  xtwitterscraper.Bool(true),
		BusinessEmailOnly: xtwitterscraper.Bool(true),
		EmailDomains:      xtwitterscraper.String("acme.com,stanford.edu"),
		Fields:            xtwitterscraper.String("id,text,author.username"),
		HasEmail:          xtwitterscraper.Bool(true),
		HasLocation:       xtwitterscraper.Bool(true),
		HasWebsite:        xtwitterscraper.Bool(true),
		LocationContains:  xtwitterscraper.String("locationContains"),
		MaxFollowers:      xtwitterscraper.Int(1),
		MaxFollowing:      xtwitterscraper.Int(1),
		MaxStatuses:       xtwitterscraper.Int(1),
		MinAccountAgeDays: xtwitterscraper.Int(1),
		MinFollowers:      xtwitterscraper.Int(100),
		MinFollowing:      xtwitterscraper.Int(1),
		MinStatuses:       xtwitterscraper.Int(1),
		Omit:              xtwitterscraper.String("entities,media"),
		UsernameContains:  xtwitterscraper.String("usernameContains"),
		VerifiedOnly:      xtwitterscraper.Bool(true),
		VerifiedType:      xtwitterscraper.String("verifiedType"),
	})
	if err != nil {
		t.Fatal(err)
	}
	request.expect(t, http.MethodGet, "/api/v1/x/communities/1500000000000000000/members/search", "bioContains=bioContains&blueVerifiedOnly=true&businessEmailOnly=true&emailDomains=acme.com%2Cstanford.edu&fields=id%2Ctext%2Cauthor.username&hasEmail=true&hasLocation=true&hasWebsite=true&locationContains=locationContains&maxFollowers=1&maxFollowing=1&maxStatuses=1&minAccountAgeDays=1&minFollowers=100&minFollowing=1&minStatuses=1&omit=entities%2Cmedia&q=marckohlbrugge&usernameContains=usernameContains&verifiedOnly=true&verifiedType=verifiedType")
	expectDecoded(t, res)
	if res.HasNextPage ||
		res.NextCursor != "" {
		t.Fatalf("decoded %s", res.RawJSON())
	}
}

func TestXCommunityMediaListFixture(t *testing.T) {
	client, request := fixtureClient(t, "PaginatedTweets.json")
	res, err := client.X.Communities.Media.List(context.Background(), "1500000000000000000", xtwitterscraper.XCommunityMediaListParams{
		AnyWords:          xtwitterscraper.String("anyWords"),
		BlueVerifiedOnly:  xtwitterscraper.Bool(true),
		Cashtags:          xtwitterscraper.String("cashtags"),
		ConversationID:    xtwitterscraper.String("conversationId"),
		Cursor:            xtwitterscraper.String("cursor"),
		ExactPhrase:       xtwitterscraper.String("exactPhrase"),
		ExcludeWords:      xtwitterscraper.String("excludeWords"),
		Fields:            xtwitterscraper.String("id,text,author.username"),
		FromUser:          xtwitterscraper.String("fromUser"),
		Hashtags:          xtwitterscraper.String("hashtags"),
		InReplyToTweetID:  xtwitterscraper.String("inReplyToTweetId"),
		Language:          xtwitterscraper.String("language"),
		MaxFaves:          xtwitterscraper.Int(1),
		MaxID:             xtwitterscraper.String("maxId"),
		MaxQuotes:         xtwitterscraper.Int(1),
		MaxReplies:        xtwitterscraper.Int(1),
		MaxRetweets:       xtwitterscraper.Int(1),
		MediaType:         xtwitterscraper.XCommunityMediaListParamsMediaTypeImages,
		Mentioning:        xtwitterscraper.String("mentioning"),
		MinBookmarks:      xtwitterscraper.Int(1),
		MinLikes:          xtwitterscraper.Int(1),
		MinQuotes:         xtwitterscraper.Int(1),
		MinReplies:        xtwitterscraper.Int(1),
		MinRetweets:       xtwitterscraper.Int(1),
		MinViews:          xtwitterscraper.Int(1),
		NativeRetweets:    xtwitterscraper.Bool(true),
		Omit:              xtwitterscraper.String("entities,media"),
		PageSize:          xtwitterscraper.Int(20),
		Quotes:            xtwitterscraper.XCommunityMediaListParamsQuotesInclude,
		QuotesOfTweetID:   xtwitterscraper.String("quotesOfTweetId"),
		Replies:           xtwitterscraper.XCommunityMediaListParamsRepliesInclude,
		Retweets:          xtwitterscraper.XCommunityMediaListParamsRetweetsInclude,
		RetweetsOfTweetID: xtwitterscraper.String("retweetsOfTweetId"),
		SinceDate:         xtwitterscraper.Time(time.Date(2026, time.October, 13, 0, 0, 0, 0, time.UTC)),
		SinceID:           xtwitterscraper.String("sinceId"),
		ToUser:            xtwitterscraper.String("toUser"),
		UntilDate:         xtwitterscraper.Time(time.Date(2026, time.October, 13, 0, 0, 0, 0, time.UTC)),
		URL:               xtwitterscraper.String("url"),
		VerifiedOnly:      xtwitterscraper.Bool(true),
		WithinTime:        xtwitterscraper.String("7d"),
	})
	if err != nil {
		t.Fatal(err)
	}
	request.expect(t, http.MethodGet, "/api/v1/x/communities/1500000000000000000/media", "anyWords=anyWords&blueVerifiedOnly=true&cashtags=cashtags&conversationId=conversationId&cursor=cursor&exactPhrase=exactPhrase&excludeWords=excludeWords&fields=id%2Ctext%2Cauthor.username&fromUser=fromUser&hashtags=hashtags&inReplyToTweetId=inReplyToTweetId&language=language&maxFaves=1&maxId=maxId&maxQuotes=1&maxReplies=1&maxRetweets=1&mediaType=images&mentioning=mentioning&minBookmarks=1&minLikes=1&minQuotes=1&minReplies=1&minRetweets=1&minViews=1&nativeRetweets=true&omit=entities%2Cmedia&pageSize=20&quotes=include&quotesOfTweetId=quotesOfTweetId&replies=include&retweets=include&retweetsOfTweetId=retweetsOfTweetId&sinceDate=2026-10-13&sinceId=sinceId&toUser=toUser&untilDate=2026-10-13&url=url&verifiedOnly=true&withinTime=7d")
	expectDecoded(t, res)
	if !res.HasNextPage ||
		res.NextCursor != "DAAHCgABHTK6Vr___-sLAAIAAAATMjEwMTc3MDE4NjE1NjA1Mjk1OQgAAwAAAAIAAA" ||
		res.FilteredCount != 0 {
		t.Fatalf("decoded %s", res.RawJSON())
	}
}

func TestCommunityUnionVariants(t *testing.T) {
	payload, err := fixtures.ReadFile("testdata/fixtures/Community.json")
	if err != nil {
		t.Fatal(err)
	}
	var community shared.Community
	if err := community.UnmarshalJSON(payload); err != nil {
		t.Fatal(err)
	}
	expectDecoded(t, community)
	creator := community.Creator.AsUserProfile()
	if community.ID != "1493446837214187523" ||
		community.Creator.Username != "marckohlbrugge" ||
		creator.ID != "641633" ||
		creator.Name != "Marc Köhlbrugge" ||
		creator.Followers != 89054 ||
		community.MembersFacepile[0].AsUserProfile().Username != "marckohlbrugge" {
		t.Fatalf("decoded %s", community.RawJSON())
	}
	var summary shared.CommunityCreatorUnion
	if err := summary.UnmarshalJSON([]byte(`{"id":"641633","username":"marckohlbrugge","verified":true}`)); err != nil {
		t.Fatal(err)
	}
	if object := summary.AsCommunityCreatorObject(); object.ID != "641633" || object.Username != "marckohlbrugge" || !object.Verified {
		t.Fatalf("decoded %s", summary.RawJSON())
	}
	payload, err = fixtures.ReadFile("testdata/fixtures/findCommunities.json")
	if err != nil {
		t.Fatal(err)
	}
	var page shared.CommunityPage
	if err := page.UnmarshalJSON(payload); err != nil {
		t.Fatal(err)
	}
	expectDecoded(t, page)
	preview := page.Communities[0].MembersFacepile[0].AsCommunityMembersFacepileCommunityMemberPreview()
	if preview.ID != "2864203879" || preview.ProfilePicture != "https://pbs.twimg.com/profile_images/example.jpg" {
		t.Fatalf("decoded %s", page.RawJSON())
	}
}
