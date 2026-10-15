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

func TestXTweetCheckRepostFixture(t *testing.T) {
	client, request := fixtureClient(t, "getTweetRepostCheck.json")
	res, err := client.X.Tweets.CheckRepost(context.Background(), "2098439298617856319", xtwitterscraper.XTweetCheckRepostParams{
		User:   "nasa",
		Fields: xtwitterscraper.String("id,text,author.username"),
		Omit:   xtwitterscraper.String("entities,media"),
	})
	if err != nil {
		t.Fatal(err)
	}
	request.expect(t, http.MethodGet, "/api/v1/x/tweets/2098439298617856319/retweeters/check", "fields=id%2Ctext%2Cauthor.username&omit=entities%2Cmedia&user=nasa")
	expectDecoded(t, res)
	if !res.IsRetweeted ||
		!res.RetweetedAt.Equal(time.Date(2026, time.September, 11, 16, 59, 50, 0, time.UTC)) ||
		res.RetweetID != "2098456494362460375" ||
		res.TweetID != "2098439298617856319" ||
		res.UserID != "3065618342" ||
		res.CheckedPosts != 20 ||
		res.Complete {
		t.Fatalf("decoded %s", res.RawJSON())
	}
}

func TestXTweetEmbedFixture(t *testing.T) {
	client, request := fixtureClient(t, "getTweetEmbed.json")
	res, err := client.X.Tweets.Embed(context.Background(), "2102882545427714266", xtwitterscraper.XTweetEmbedParams{
		Fields:     xtwitterscraper.String("id,text,author.username"),
		HideMedia:  xtwitterscraper.Bool(true),
		HideThread: xtwitterscraper.Bool(true),
		Omit:       xtwitterscraper.String("entities,media"),
		OmitScript: xtwitterscraper.Bool(true),
		Theme:      xtwitterscraper.XTweetEmbedParamsThemeDark,
	})
	if err != nil {
		t.Fatal(err)
	}
	request.expect(t, http.MethodGet, "/api/v1/x/tweets/2102882545427714266/embed", "fields=id%2Ctext%2Cauthor.username&hideMedia=true&hideThread=true&omit=entities%2Cmedia&omitScript=true&theme=dark")
	expectDecoded(t, res)
	if res.Embed.TweetID != "2102882545427714266" ||
		res.Embed.URL != "https://x.com/NASA/status/2102882545427714266" ||
		res.Embed.AuthorName != "NASA" ||
		res.Embed.AuthorURL != "https://x.com/NASA" ||
		res.Embed.HTML != "<blockquote class=\"twitter-tweet\">...</blockquote>" ||
		res.Embed.Width != 550 {
		t.Fatalf("decoded %s", res.RawJSON())
	}
}

func TestXTweetGetHiddenRepliesFixture(t *testing.T) {
	client, request := fixtureClient(t, "PaginatedTweets.json")
	res, err := client.X.Tweets.GetHiddenReplies(context.Background(), "2102882545427714266", xtwitterscraper.XTweetGetHiddenRepliesParams{
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
		MediaType:         xtwitterscraper.XTweetGetHiddenRepliesParamsMediaTypeImages,
		Mentioning:        xtwitterscraper.String("mentioning"),
		MinBookmarks:      xtwitterscraper.Int(1),
		MinLikes:          xtwitterscraper.Int(5),
		MinQuotes:         xtwitterscraper.Int(1),
		MinReplies:        xtwitterscraper.Int(1),
		MinRetweets:       xtwitterscraper.Int(1),
		MinViews:          xtwitterscraper.Int(1),
		NativeRetweets:    xtwitterscraper.Bool(true),
		Omit:              xtwitterscraper.String("entities,media"),
		PageSize:          xtwitterscraper.Int(20),
		Quotes:            xtwitterscraper.XTweetGetHiddenRepliesParamsQuotesInclude,
		QuotesOfTweetID:   xtwitterscraper.String("quotesOfTweetId"),
		Replies:           xtwitterscraper.XTweetGetHiddenRepliesParamsRepliesOnly,
		Retweets:          xtwitterscraper.XTweetGetHiddenRepliesParamsRetweetsInclude,
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
	request.expect(t, http.MethodGet, "/api/v1/x/tweets/2102882545427714266/hidden-replies", "anyWords=anyWords&blueVerifiedOnly=true&cashtags=cashtags&conversationId=conversationId&cursor=cursor&exactPhrase=exactPhrase&excludeWords=excludeWords&fields=id%2Ctext%2Cauthor.username&fromUser=fromUser&hashtags=hashtags&inReplyToTweetId=inReplyToTweetId&language=language&maxFaves=1&maxId=maxId&maxQuotes=1&maxReplies=1&maxRetweets=1&mediaType=images&mentioning=mentioning&minBookmarks=1&minLikes=5&minQuotes=1&minReplies=1&minRetweets=1&minViews=1&nativeRetweets=true&omit=entities%2Cmedia&pageSize=20&quotes=include&quotesOfTweetId=quotesOfTweetId&replies=only&retweets=include&retweetsOfTweetId=retweetsOfTweetId&sinceDate=2026-10-13&sinceId=sinceId&toUser=toUser&untilDate=2026-10-13&url=url&verifiedOnly=true&withinTime=7d")
	expectDecoded(t, res)
	if !res.HasNextPage ||
		res.NextCursor != "DAAHCgABHTK6Vr___-sLAAIAAAATMjEwMTc3MDE4NjE1NjA1Mjk1OQgAAwAAAAIAAA" ||
		res.FilteredCount != 0 ||
		len(res.Tweets) != 2 ||
		res.Tweets[0].ID != "2102761519985332442" ||
		res.Tweets[0].RetweetCount != 1364 ||
		res.Tweets[0].ReplyCount != 271 ||
		res.Tweets[0].LikeCount != 9775 ||
		!res.ReadBackTo.Equal(time.Date(2026, time.October, 8, 0, 0, 0, 0, time.UTC)) ||
		res.SearchContext.SmartTag.ID != "cashtag-tsla" ||
		!res.SearchContext.SmartTag.Active ||
		!res.SearchContext.SmartTag.SupportsOhlc ||
		res.SearchedQuery != "$TSLA since:2026-10-08" ||
		!res.WindowComplete {
		t.Fatalf("decoded %s", res.RawJSON())
	}
}

func TestXTweetGetInsightsFixture(t *testing.T) {
	client, request := fixtureClient(t, "getTweetInsights.json")
	res, err := client.X.Tweets.GetInsights(context.Background(), "2102882545427714266", xtwitterscraper.XTweetGetInsightsParams{
		Fields: xtwitterscraper.String("insights"),
		Omit:   xtwitterscraper.String("insights.metrics.views"),
	})
	if err != nil {
		t.Fatal(err)
	}
	request.expect(t, http.MethodGet, "/api/v1/x/tweets/2102882545427714266/insights", "fields=insights&omit=insights.metrics.views")
	expectDecoded(t, res)
	if res.Insights.TweetID != "2102882545427714266" ||
		res.Insights.AuthorID != "11348282" ||
		res.Insights.AuthorUserName != "NASA" ||
		!res.Insights.CreatedAt.Equal(time.Date(2026, time.September, 23, 22, 7, 23, 0, time.UTC)) ||
		res.Insights.AgeHours != 133.88 ||
		res.Insights.Engagements != 10122 ||
		res.Insights.EngagementsPerHour != 75.605 ||
		res.Insights.EngagementPerView != 0.7839 {
		t.Fatalf("decoded %s", res.RawJSON())
	}
}

func TestXTweetGetRetweeterIDsFixture(t *testing.T) {
	client, request := fixtureClient(t, "PaginatedUserIds.json")
	res, err := client.X.Tweets.GetRetweeterIDs(context.Background(), "2098439298617856319", xtwitterscraper.XTweetGetRetweeterIDsParams{
		Cursor:   xtwitterscraper.String("DAACCgACGRElMJcAAA"),
		PageSize: xtwitterscraper.Int(100),
	})
	if err != nil {
		t.Fatal(err)
	}
	request.expect(t, http.MethodGet, "/api/v1/x/tweets/2098439298617856319/retweeter-ids", "cursor=DAACCgACGRElMJcAAA&pageSize=100")
	expectDecoded(t, res)
	if !res.HasNextPage ||
		res.NextCursor != "1878079982871151535" ||
		len(res.IDs) != 2 ||
		res.IDs[0] != "9876543210" {
		t.Fatalf("decoded %s", res.RawJSON())
	}
}

func TestXTweetGetSubtitlesFixture(t *testing.T) {
	client, request := fixtureClient(t, "getTweetSubtitles.json")
	res, err := client.X.Tweets.GetSubtitles(context.Background(), "2104593672624886205", xtwitterscraper.XTweetGetSubtitlesParams{
		Fields: xtwitterscraper.String("id,text,author.username"),
		Omit:   xtwitterscraper.String("entities,media"),
	})
	if err != nil {
		t.Fatal(err)
	}
	request.expect(t, http.MethodGet, "/api/v1/x/tweets/2104593672624886205/subtitles", "fields=id%2Ctext%2Cauthor.username&omit=entities%2Cmedia")
	expectDecoded(t, res)
	if res.Subtitles.TweetID != "2104593672624886205" ||
		len(res.Subtitles.Tracks) != 1 ||
		res.Subtitles.Tracks[0].MediaKey != "13_2104593621043314689" ||
		res.Subtitles.Tracks[0].Language != "en" ||
		res.Subtitles.Tracks[0].Name != "Planetary Defenders_Caption.en_US.srt" ||
		res.Subtitles.Tracks[0].AutoGenerated ||
		res.Subtitles.Tracks[0].Text != "This asteroid 2024 YR4 might get interesting." ||
		res.Subtitles.Tracks[0].Truncated {
		t.Fatalf("decoded %s", res.RawJSON())
	}
}

func TestXTweetListTopFixture(t *testing.T) {
	client, request := fixtureClient(t, "PaginatedTweets.json")
	res, err := client.X.Tweets.ListTop(context.Background(), xtwitterscraper.XTweetListTopParams{
		Country:  xtwitterscraper.String("US"),
		Cursor:   xtwitterscraper.String("cursor"),
		Fields:   xtwitterscraper.String("id,text,author.username"),
		Language: xtwitterscraper.String("en"),
		Metric:   xtwitterscraper.XTweetListTopParamsMetricViews,
		Omit:     xtwitterscraper.String("entities,media"),
		PageSize: xtwitterscraper.Int(20),
		Window:   xtwitterscraper.XTweetListTopParamsWindowWeek,
	})
	if err != nil {
		t.Fatal(err)
	}
	request.expect(t, http.MethodGet, "/api/v1/x/tweets/top", "country=US&cursor=cursor&fields=id%2Ctext%2Cauthor.username&language=en&metric=views&omit=entities%2Cmedia&pageSize=20&window=week")
	expectDecoded(t, res)
	if !res.HasNextPage ||
		res.NextCursor != "DAAHCgABHTK6Vr___-sLAAIAAAATMjEwMTc3MDE4NjE1NjA1Mjk1OQgAAwAAAAIAAA" ||
		res.FilteredCount != 0 ||
		len(res.Tweets) != 2 ||
		res.Tweets[0].ID != "2102761519985332442" ||
		res.Tweets[0].RetweetCount != 1364 ||
		res.Tweets[0].ReplyCount != 271 ||
		res.Tweets[0].LikeCount != 9775 {
		t.Fatalf("decoded %s", res.RawJSON())
	}
}

func TestXTweetResolveLinksFixture(t *testing.T) {
	client, request := fixtureClient(t, "resolveLinks.json")
	res, err := client.X.Tweets.ResolveLinks(context.Background(), xtwitterscraper.XTweetResolveLinksParams{
		URLs:   "https://t.co/Xe1nm2pkSX",
		Fields: xtwitterscraper.String("id,text,author.username"),
		Omit:   xtwitterscraper.String("entities,media"),
	})
	if err != nil {
		t.Fatal(err)
	}
	request.expect(t, http.MethodGet, "/api/v1/x/links/resolve", "fields=id%2Ctext%2Cauthor.username&omit=entities%2Cmedia&urls=https%3A%2F%2Ft.co%2FXe1nm2pkSX")
	expectDecoded(t, res)
	if len(res.Links) != 1 ||
		res.Links[0].URL != "https://t.co/Xe1nm2pkSX" ||
		res.Links[0].ResolvedURL != "https://go.nasa.gov/3VcvAQn" ||
		len(res.UnprocessedURLs) != 0 {
		t.Fatalf("decoded %s", res.RawJSON())
	}
}

func TestXTweetTranslateFixture(t *testing.T) {
	client, request := fixtureClient(t, "translateTweet.json")
	res, err := client.X.Tweets.Translate(context.Background(), "2102882545427714266", xtwitterscraper.XTweetTranslateParams{
		Language: "tr",
		Fields:   xtwitterscraper.String("id,text,author.username"),
		Omit:     xtwitterscraper.String("entities,media"),
	})
	if err != nil {
		t.Fatal(err)
	}
	request.expect(t, http.MethodGet, "/api/v1/x/tweets/2102882545427714266/translation", "fields=id%2Ctext%2Cauthor.username&language=tr&omit=entities%2Cmedia")
	expectDecoded(t, res)
	if res.Translation.TweetID != "2102882545427714266" ||
		res.Translation.Language != "tr" ||
		res.Translation.SourceLanguage != "en" ||
		res.Translation.Text != "Artemis Anlaşmaları’na Hoş Geldiniz, Hırvatistan https://t.co/Xe1nm2pkSX" ||
		len(res.Translation.Entities.URLs) != 1 ||
		res.Translation.Entities.URLs[0].URL != "https://t.co/Xe1nm2pkSX" ||
		res.Translation.Entities.URLs[0].ExpandedURL != "https://go.nasa.gov/3VcvAQn" ||
		res.Translation.Entities.URLs[0].DisplayURL != "go.nasa.gov/3VcvAQn" {
		t.Fatalf("decoded %s", res.RawJSON())
	}
}
