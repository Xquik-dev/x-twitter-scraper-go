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

func TestXUserGetAffiliatesFixture(t *testing.T) {
	client, request := fixtureClient(t, "PaginatedUsers.json")
	res, err := client.X.Users.GetAffiliates(context.Background(), "nasa", xtwitterscraper.XUserGetAffiliatesParams{
		BioContains:       xtwitterscraper.String("bioContains"),
		BlueVerifiedOnly:  xtwitterscraper.Bool(true),
		BusinessEmailOnly: xtwitterscraper.Bool(true),
		Cursor:            xtwitterscraper.String("cursor"),
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
		MinFollowers:      xtwitterscraper.Int(1000),
		MinFollowing:      xtwitterscraper.Int(1),
		MinStatuses:       xtwitterscraper.Int(1),
		Omit:              xtwitterscraper.String("entities,media"),
		PageSize:          xtwitterscraper.Int(50),
		UsernameContains:  xtwitterscraper.String("usernameContains"),
		VerifiedOnly:      xtwitterscraper.Bool(true),
		VerifiedType:      xtwitterscraper.String("verifiedType"),
	})
	if err != nil {
		t.Fatal(err)
	}
	request.expect(t, http.MethodGet, "/api/v1/x/users/nasa/affiliates", "bioContains=bioContains&blueVerifiedOnly=true&businessEmailOnly=true&cursor=cursor&emailDomains=acme.com%2Cstanford.edu&fields=id%2Ctext%2Cauthor.username&hasEmail=true&hasLocation=true&hasWebsite=true&locationContains=locationContains&maxFollowers=1&maxFollowing=1&maxStatuses=1&minAccountAgeDays=1&minFollowers=1000&minFollowing=1&minStatuses=1&omit=entities%2Cmedia&pageSize=50&usernameContains=usernameContains&verifiedOnly=true&verifiedType=verifiedType")
	expectDecoded(t, res)
	if res.HasNextPage ||
		res.NextCursor != "" {
		t.Fatalf("decoded %s", res.RawJSON())
	}
}

func TestXUserGetArticlesFixture(t *testing.T) {
	client, request := fixtureClient(t, "PaginatedTweets.json")
	res, err := client.X.Users.GetArticles(context.Background(), "nasa", xtwitterscraper.XUserGetArticlesParams{
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
		Language:          xtwitterscraper.String("en"),
		MaxFaves:          xtwitterscraper.Int(1),
		MaxID:             xtwitterscraper.String("maxId"),
		MaxQuotes:         xtwitterscraper.Int(1),
		MaxReplies:        xtwitterscraper.Int(1),
		MaxRetweets:       xtwitterscraper.Int(1),
		MediaType:         xtwitterscraper.XUserGetArticlesParamsMediaTypeImages,
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
		Quotes:            xtwitterscraper.XUserGetArticlesParamsQuotesInclude,
		QuotesOfTweetID:   xtwitterscraper.String("quotesOfTweetId"),
		Replies:           xtwitterscraper.XUserGetArticlesParamsRepliesInclude,
		Retweets:          xtwitterscraper.XUserGetArticlesParamsRetweetsInclude,
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
	request.expect(t, http.MethodGet, "/api/v1/x/users/nasa/articles", "anyWords=anyWords&blueVerifiedOnly=true&cashtags=cashtags&conversationId=conversationId&cursor=cursor&exactPhrase=exactPhrase&excludeWords=excludeWords&fields=id%2Ctext%2Cauthor.username&fromUser=fromUser&hashtags=hashtags&inReplyToTweetId=inReplyToTweetId&language=en&maxFaves=1&maxId=maxId&maxQuotes=1&maxReplies=1&maxRetweets=1&mediaType=images&mentioning=mentioning&minBookmarks=1&minLikes=1&minQuotes=1&minReplies=1&minRetweets=1&minViews=1&nativeRetweets=true&omit=entities%2Cmedia&pageSize=20&quotes=include&quotesOfTweetId=quotesOfTweetId&replies=include&retweets=include&retweetsOfTweetId=retweetsOfTweetId&sinceDate=2026-10-13&sinceId=sinceId&toUser=toUser&untilDate=2026-10-13&url=url&verifiedOnly=true&withinTime=7d")
	expectDecoded(t, res)
	if !res.HasNextPage ||
		res.NextCursor != "DAAHCgABHTK6Vr___-sLAAIAAAATMjEwMTc3MDE4NjE1NjA1Mjk1OQgAAwAAAAIAAA" ||
		res.FilteredCount != 0 {
		t.Fatalf("decoded %s", res.RawJSON())
	}
}

func TestXUserGetBatchTweetsFixture(t *testing.T) {
	client, request := fixtureClient(t, "getBatchUserTweets.json")
	res, err := client.X.Users.GetBatchTweets(context.Background(), xtwitterscraper.XUserGetBatchTweetsParams{
		Fields:    xtwitterscraper.String("id,text,author.username"),
		IDs:       xtwitterscraper.String("ids"),
		Omit:      xtwitterscraper.String("entities,media"),
		PageSize:  xtwitterscraper.Int(20),
		Usernames: xtwitterscraper.String("nasa,@OpenAI"),
	})
	if err != nil {
		t.Fatal(err)
	}
	request.expect(t, http.MethodGet, "/api/v1/x/users/batch/tweets", "fields=id%2Ctext%2Cauthor.username&ids=ids&omit=entities%2Cmedia&pageSize=20&usernames=nasa%2C%40OpenAI")
	expectDecoded(t, res)
	if len(res.Tweets) != 1 ||
		res.HasNextPage ||
		res.NextCursor != "" ||
		len(res.UnavailableIDs) != 1 ||
		len(res.UnprocessedIDs) != 0 ||
		res.UnavailableIDs[0] != "1234567890" ||
		res.Tweets[0].ID != "1234567890" ||
		res.Tweets[0].Text != "Just launched our new feature!" ||
		res.Tweets[0].CreatedAt != "2025-01-15T12:00:00Z" ||
		res.Tweets[0].LikeCount != 42 ||
		res.Tweets[0].RetweetCount != 5 ||
		res.Tweets[0].ReplyCount != 3 ||
		res.Tweets[0].QuoteCount != 1 ||
		res.Tweets[0].ViewCount != 1500 ||
		res.Tweets[0].BookmarkCount != 2 {
		t.Fatalf("decoded %s", res.RawJSON())
	}
}

func TestXUserGetFollowerIDsFixture(t *testing.T) {
	client, request := fixtureClient(t, "PaginatedUserIds.json")
	res, err := client.X.Users.GetFollowerIDs(context.Background(), "11348282", xtwitterscraper.XUserGetFollowerIDsParams{
		Cursor:   xtwitterscraper.String("DAACCgACGRElMJcAAA"),
		PageSize: xtwitterscraper.Int(5000),
	})
	if err != nil {
		t.Fatal(err)
	}
	request.expect(t, http.MethodGet, "/api/v1/x/users/11348282/follower-ids", "cursor=DAACCgACGRElMJcAAA&pageSize=5000")
	expectDecoded(t, res)
	if len(res.IDs) != 2 ||
		!res.HasNextPage ||
		res.NextCursor != "1878079982871151535" ||
		res.IDs[0] != "9876543210" {
		t.Fatalf("decoded %s", res.RawJSON())
	}
}

func TestXUserGetFollowingIDsFixture(t *testing.T) {
	client, request := fixtureClient(t, "PaginatedUserIds.json")
	res, err := client.X.Users.GetFollowingIDs(context.Background(), "11348282", xtwitterscraper.XUserGetFollowingIDsParams{
		Cursor:   xtwitterscraper.String("cursor"),
		PageSize: xtwitterscraper.Int(1000),
	})
	if err != nil {
		t.Fatal(err)
	}
	request.expect(t, http.MethodGet, "/api/v1/x/users/11348282/following-ids", "cursor=cursor&pageSize=1000")
	expectDecoded(t, res)
	if len(res.IDs) != 2 ||
		!res.HasNextPage ||
		res.NextCursor != "1878079982871151535" ||
		res.IDs[0] != "9876543210" {
		t.Fatalf("decoded %s", res.RawJSON())
	}
}

func TestXUserGetHighlightsFixture(t *testing.T) {
	client, request := fixtureClient(t, "PaginatedTweets.json")
	res, err := client.X.Users.GetHighlights(context.Background(), "nasa", xtwitterscraper.XUserGetHighlightsParams{
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
		MediaType:         xtwitterscraper.XUserGetHighlightsParamsMediaTypeVideos,
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
		Quotes:            xtwitterscraper.XUserGetHighlightsParamsQuotesInclude,
		QuotesOfTweetID:   xtwitterscraper.String("quotesOfTweetId"),
		Replies:           xtwitterscraper.XUserGetHighlightsParamsRepliesInclude,
		Retweets:          xtwitterscraper.XUserGetHighlightsParamsRetweetsInclude,
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
	request.expect(t, http.MethodGet, "/api/v1/x/users/nasa/highlights", "anyWords=anyWords&blueVerifiedOnly=true&cashtags=cashtags&conversationId=conversationId&cursor=cursor&exactPhrase=exactPhrase&excludeWords=excludeWords&fields=id%2Ctext%2Cauthor.username&fromUser=fromUser&hashtags=hashtags&inReplyToTweetId=inReplyToTweetId&language=language&maxFaves=1&maxId=maxId&maxQuotes=1&maxReplies=1&maxRetweets=1&mediaType=videos&mentioning=mentioning&minBookmarks=1&minLikes=1&minQuotes=1&minReplies=1&minRetweets=1&minViews=1&nativeRetweets=true&omit=entities%2Cmedia&pageSize=20&quotes=include&quotesOfTweetId=quotesOfTweetId&replies=include&retweets=include&retweetsOfTweetId=retweetsOfTweetId&sinceDate=2026-10-13&sinceId=sinceId&toUser=toUser&untilDate=2026-10-13&url=url&verifiedOnly=true&withinTime=7d")
	expectDecoded(t, res)
	if !res.HasNextPage ||
		res.NextCursor != "DAAHCgABHTK6Vr___-sLAAIAAAATMjEwMTc3MDE4NjE1NjA1Mjk1OQgAAwAAAAIAAA" ||
		res.FilteredCount != 0 {
		t.Fatalf("decoded %s", res.RawJSON())
	}
}

func TestXUserGetInsightsFixture(t *testing.T) {
	client, request := fixtureClient(t, "getUserInsights.json")
	res, err := client.X.Users.GetInsights(context.Background(), "nasa", xtwitterscraper.XUserGetInsightsParams{
		Fields: xtwitterscraper.String("insights"),
		Omit:   xtwitterscraper.String("entities,media"),
	})
	if err != nil {
		t.Fatal(err)
	}
	request.expect(t, http.MethodGet, "/api/v1/x/users/nasa/insights", "fields=insights&omit=entities%2Cmedia")
	expectDecoded(t, res)
	if res.Message != "This account's posts are private, so the report has its counts only." ||
		res.Insights.UserID != "11348282" ||
		res.Insights.UserName != "NASA" ||
		res.Insights.Followers != 92376187 ||
		res.Insights.Following != 115 ||
		res.Insights.Posts != 74336 ||
		res.Insights.AccountAgeDays != 6858 ||
		res.Insights.PostsPrivate ||
		res.Insights.EngagementRate != 0.0047 ||
		res.Insights.EngagementPerView != 0.5516 ||
		res.Insights.PostsPerDay != 1.6814 ||
		res.Insights.TopPostID != "2102761519985332442" ||
		res.Insights.Sample.Posts != 10 ||
		!res.Insights.Sample.OldestAt.Equal(time.Date(2026, time.September, 23, 13, 15, 28, 0, time.UTC)) ||
		!res.Insights.Sample.NewestAt.Equal(time.Date(2026, time.September, 28, 22, 12, 5, 0, time.UTC)) ||
		res.Insights.Averages.Likes != 3684.2 ||
		res.Insights.Averages.Replies != 142.7 ||
		res.Insights.Averages.Reposts != 459.8 ||
		res.Insights.Averages.Quotes != 45 ||
		res.Insights.Averages.Bookmarks != 209.7 ||
		res.Insights.Averages.Views != 785247.8 ||
		res.Insights.Averages.Engagements != 4331.7 {
		t.Fatalf("decoded %s", res.RawJSON())
	}
}

func TestXUserGetListMembershipsFixture(t *testing.T) {
	client, request := fixtureClient(t, "PaginatedListsWithList.json")
	res, err := client.X.Users.GetListMemberships(context.Background(), "nasa", xtwitterscraper.XUserGetListMembershipsParams{
		Cursor:   xtwitterscraper.String("cursor"),
		Fields:   xtwitterscraper.String("id,text,author.username"),
		Omit:     xtwitterscraper.String("entities,media"),
		PageSize: xtwitterscraper.Int(50),
	})
	if err != nil {
		t.Fatal(err)
	}
	request.expect(t, http.MethodGet, "/api/v1/x/users/nasa/list-memberships", "cursor=cursor&fields=id%2Ctext%2Cauthor.username&omit=entities%2Cmedia&pageSize=50")
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

func TestXUserGetListsFixture(t *testing.T) {
	client, request := fixtureClient(t, "PaginatedListsWithList.json")
	res, err := client.X.Users.GetLists(context.Background(), "nasa", xtwitterscraper.XUserGetListsParams{
		Cursor:   xtwitterscraper.String("DAACCgACGRElMJcAAA"),
		Fields:   xtwitterscraper.String("id,text,author.username"),
		Omit:     xtwitterscraper.String("entities,media"),
		PageSize: xtwitterscraper.Int(1),
	})
	if err != nil {
		t.Fatal(err)
	}
	request.expect(t, http.MethodGet, "/api/v1/x/users/nasa/lists", "cursor=DAACCgACGRElMJcAAA&fields=id%2Ctext%2Cauthor.username&omit=entities%2Cmedia&pageSize=1")
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

func TestXUserGetLiveFixture(t *testing.T) {
	client, request := fixtureClient(t, "getUserLiveStatus.json")
	res, err := client.X.Users.GetLive(context.Background(), "nasa", xtwitterscraper.XUserGetLiveParams{
		Fields: xtwitterscraper.String("id,text,author.username"),
		Omit:   xtwitterscraper.String("entities,media"),
	})
	if err != nil {
		t.Fatal(err)
	}
	request.expect(t, http.MethodGet, "/api/v1/x/users/nasa/live", "fields=id%2Ctext%2Cauthor.username&omit=entities%2Cmedia")
	expectDecoded(t, res)
	if res.Live ||
		res.UserID != "11348282" ||
		res.UserName != "NASA" ||
		res.Broadcast.ID != "1DxLdZyVPMMxm" ||
		res.Broadcast.Title != "NASA's SpaceX Crew-12 Re-Entry & Splashdown" ||
		res.Broadcast.State != "Ended" ||
		!res.Broadcast.StartedAt.Equal(time.Date(2026, time.October, 8, 14, 20, 33, 863000000, time.UTC)) ||
		!res.Broadcast.EndedAt.Equal(time.Date(2026, time.October, 8, 16, 38, 16, 770000000, time.UTC)) ||
		res.Broadcast.TotalWatched != 220908 ||
		res.Broadcast.TotalWatching != 119617 ||
		res.Broadcast.Width != 3840 ||
		res.Broadcast.Height != 2160 {
		t.Fatalf("decoded %s", res.RawJSON())
	}
}

func TestXUserGetSimilarFixture(t *testing.T) {
	client, request := fixtureClient(t, "PaginatedUsers.json")
	res, err := client.X.Users.GetSimilar(context.Background(), "nasa", xtwitterscraper.XUserGetSimilarParams{
		BioContains:       xtwitterscraper.String("bioContains"),
		BlueVerifiedOnly:  xtwitterscraper.Bool(true),
		BusinessEmailOnly: xtwitterscraper.Bool(true),
		Cursor:            xtwitterscraper.String("cursor"),
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
		MinFollowers:      xtwitterscraper.Int(1),
		MinFollowing:      xtwitterscraper.Int(1),
		MinStatuses:       xtwitterscraper.Int(1),
		Omit:              xtwitterscraper.String("entities,media"),
		PageSize:          xtwitterscraper.Int(20),
		UsernameContains:  xtwitterscraper.String("usernameContains"),
		VerifiedOnly:      xtwitterscraper.Bool(true),
		VerifiedType:      xtwitterscraper.String("verifiedType"),
	})
	if err != nil {
		t.Fatal(err)
	}
	request.expect(t, http.MethodGet, "/api/v1/x/users/nasa/similar", "bioContains=bioContains&blueVerifiedOnly=true&businessEmailOnly=true&cursor=cursor&emailDomains=acme.com%2Cstanford.edu&fields=id%2Ctext%2Cauthor.username&hasEmail=true&hasLocation=true&hasWebsite=true&locationContains=locationContains&maxFollowers=1&maxFollowing=1&maxStatuses=1&minAccountAgeDays=1&minFollowers=1&minFollowing=1&minStatuses=1&omit=entities%2Cmedia&pageSize=20&usernameContains=usernameContains&verifiedOnly=true&verifiedType=verifiedType")
	expectDecoded(t, res)
	if res.HasNextPage ||
		res.NextCursor != "" {
		t.Fatalf("decoded %s", res.RawJSON())
	}
}

func TestXUserGetSubscriptionsFixture(t *testing.T) {
	client, request := fixtureClient(t, "PaginatedUsers.json")
	res, err := client.X.Users.GetSubscriptions(context.Background(), "nasa", xtwitterscraper.XUserGetSubscriptionsParams{
		BioContains:       xtwitterscraper.String("space"),
		BlueVerifiedOnly:  xtwitterscraper.Bool(true),
		BusinessEmailOnly: xtwitterscraper.Bool(true),
		Cursor:            xtwitterscraper.String("cursor"),
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
		MinFollowers:      xtwitterscraper.Int(1),
		MinFollowing:      xtwitterscraper.Int(1),
		MinStatuses:       xtwitterscraper.Int(1),
		Omit:              xtwitterscraper.String("entities,media"),
		PageSize:          xtwitterscraper.Int(1),
		UsernameContains:  xtwitterscraper.String("usernameContains"),
		VerifiedOnly:      xtwitterscraper.Bool(true),
		VerifiedType:      xtwitterscraper.String("verifiedType"),
	})
	if err != nil {
		t.Fatal(err)
	}
	request.expect(t, http.MethodGet, "/api/v1/x/users/nasa/subscriptions", "bioContains=space&blueVerifiedOnly=true&businessEmailOnly=true&cursor=cursor&emailDomains=acme.com%2Cstanford.edu&fields=id%2Ctext%2Cauthor.username&hasEmail=true&hasLocation=true&hasWebsite=true&locationContains=locationContains&maxFollowers=1&maxFollowing=1&maxStatuses=1&minAccountAgeDays=1&minFollowers=1&minFollowing=1&minStatuses=1&omit=entities%2Cmedia&pageSize=1&usernameContains=usernameContains&verifiedOnly=true&verifiedType=verifiedType")
	expectDecoded(t, res)
	if res.HasNextPage ||
		res.NextCursor != "" {
		t.Fatalf("decoded %s", res.RawJSON())
	}
}

func TestXUserTranslateBioFixture(t *testing.T) {
	client, request := fixtureClient(t, "translateUserBio.json")
	res, err := client.X.Users.TranslateBio(context.Background(), "nasa", xtwitterscraper.XUserTranslateBioParams{
		Language: "tr",
		Fields:   xtwitterscraper.String("id,text,author.username"),
		Omit:     xtwitterscraper.String("entities,media"),
	})
	if err != nil {
		t.Fatal(err)
	}
	request.expect(t, http.MethodGet, "/api/v1/x/users/nasa/translation", "fields=id%2Ctext%2Cauthor.username&language=tr&omit=entities%2Cmedia")
	expectDecoded(t, res)
	if res.Translation.UserID != "11348282" ||
		res.Translation.UserName != "NASA" ||
		res.Translation.Language != "es" ||
		res.Translation.SourceLanguage != "en" ||
		res.Translation.Text != "Haciendo lo aparentemente imposible, posible. ✨" {
		t.Fatalf("decoded %s", res.RawJSON())
	}
}
