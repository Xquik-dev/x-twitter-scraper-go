// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

package xtwitterscraper_test

import (
	"encoding/json"
	"net/url"
	"strings"
	"testing"
	"time"

	"github.com/Xquik-dev/x-twitter-scraper-go"
)

func TestNewQueryParamsReachTheQuery(t *testing.T) {
	cases := map[string]struct {
		params interface{ URLQuery() (url.Values, error) }
		want   string
	}{
		"ExtractionExportResultsParams":        {xtwitterscraper.ExtractionExportResultsParams{BusinessEmailOnly: xtwitterscraper.Bool(true), EmailDomains: xtwitterscraper.String("emailDomains"), Format: xtwitterscraper.ExtractionExportResultsParamsFormatCsv, HasEmail: xtwitterscraper.Bool(true)}, "businessEmailOnly=true&emailDomains=emailDomains&format=csv&hasEmail=true"},
		"XCommunityGetMembersParams":           {xtwitterscraper.XCommunityGetMembersParams{BlueVerifiedOnly: xtwitterscraper.Bool(true), BusinessEmailOnly: xtwitterscraper.Bool(true), EmailDomains: xtwitterscraper.String("emailDomains"), HasEmail: xtwitterscraper.Bool(true)}, "blueVerifiedOnly=true&businessEmailOnly=true&emailDomains=emailDomains&hasEmail=true"},
		"XCommunityGetModeratorsParams":        {xtwitterscraper.XCommunityGetModeratorsParams{BlueVerifiedOnly: xtwitterscraper.Bool(true), BusinessEmailOnly: xtwitterscraper.Bool(true), EmailDomains: xtwitterscraper.String("emailDomains"), HasEmail: xtwitterscraper.Bool(true)}, "blueVerifiedOnly=true&businessEmailOnly=true&emailDomains=emailDomains&hasEmail=true"},
		"XCommunityGetSearchParams":            {xtwitterscraper.XCommunityGetSearchParams{AnyWords: xtwitterscraper.String("anyWords"), BlueVerifiedOnly: xtwitterscraper.Bool(true), Cashtags: xtwitterscraper.String("cashtags"), CommunityID: "communityId", ConversationID: xtwitterscraper.String("conversationId"), ExactPhrase: xtwitterscraper.String("exactPhrase"), ExcludeWords: xtwitterscraper.String("excludeWords"), FromUser: xtwitterscraper.String("fromUser"), Hashtags: xtwitterscraper.String("hashtags"), InReplyToTweetID: xtwitterscraper.String("inReplyToTweetId"), MaxFaves: xtwitterscraper.Int(2), MaxID: xtwitterscraper.String("maxId"), MaxQuotes: xtwitterscraper.Int(2), MaxReplies: xtwitterscraper.Int(2), MaxRetweets: xtwitterscraper.Int(2), Mentioning: xtwitterscraper.String("mentioning"), MinBookmarks: xtwitterscraper.Int(2), MinQuotes: xtwitterscraper.Int(2), NativeRetweets: xtwitterscraper.Bool(true), Q: "q", Quotes: xtwitterscraper.XCommunityGetSearchParamsQuotesInclude, QuotesOfTweetID: xtwitterscraper.String("quotesOfTweetId"), Replies: xtwitterscraper.XCommunityGetSearchParamsRepliesInclude, Retweets: xtwitterscraper.XCommunityGetSearchParamsRetweetsInclude, RetweetsOfTweetID: xtwitterscraper.String("retweetsOfTweetId"), SinceID: xtwitterscraper.String("sinceId"), ToUser: xtwitterscraper.String("toUser"), URL: xtwitterscraper.String("url"), WithinTime: xtwitterscraper.String("withinTime")}, "anyWords=anyWords&blueVerifiedOnly=true&cashtags=cashtags&communityId=communityId&conversationId=conversationId&exactPhrase=exactPhrase&excludeWords=excludeWords&fromUser=fromUser&hashtags=hashtags&inReplyToTweetId=inReplyToTweetId&maxFaves=2&maxId=maxId&maxQuotes=2&maxReplies=2&maxRetweets=2&mentioning=mentioning&minBookmarks=2&minQuotes=2&nativeRetweets=true&q=q&quotes=include&quotesOfTweetId=quotesOfTweetId&replies=include&retweets=include&retweetsOfTweetId=retweetsOfTweetId&sinceId=sinceId&toUser=toUser&url=url&withinTime=withinTime"},
		"XCommunityTweetListParams":            {xtwitterscraper.XCommunityTweetListParams{AnyWords: xtwitterscraper.String("anyWords"), BlueVerifiedOnly: xtwitterscraper.Bool(true), Cashtags: xtwitterscraper.String("cashtags"), CommunityID: "communityId", ConversationID: xtwitterscraper.String("conversationId"), ExactPhrase: xtwitterscraper.String("exactPhrase"), ExcludeWords: xtwitterscraper.String("excludeWords"), FromUser: xtwitterscraper.String("fromUser"), Hashtags: xtwitterscraper.String("hashtags"), InReplyToTweetID: xtwitterscraper.String("inReplyToTweetId"), MaxFaves: xtwitterscraper.Int(2), MaxID: xtwitterscraper.String("maxId"), MaxQuotes: xtwitterscraper.Int(2), MaxReplies: xtwitterscraper.Int(2), MaxRetweets: xtwitterscraper.Int(2), Mentioning: xtwitterscraper.String("mentioning"), MinBookmarks: xtwitterscraper.Int(2), MinQuotes: xtwitterscraper.Int(2), NativeRetweets: xtwitterscraper.Bool(true), Q: "q", Quotes: xtwitterscraper.XCommunityTweetListParamsQuotesInclude, QuotesOfTweetID: xtwitterscraper.String("quotesOfTweetId"), Replies: xtwitterscraper.XCommunityTweetListParamsRepliesInclude, Retweets: xtwitterscraper.XCommunityTweetListParamsRetweetsInclude, RetweetsOfTweetID: xtwitterscraper.String("retweetsOfTweetId"), SinceID: xtwitterscraper.String("sinceId"), ToUser: xtwitterscraper.String("toUser"), URL: xtwitterscraper.String("url"), WithinTime: xtwitterscraper.String("withinTime")}, "anyWords=anyWords&blueVerifiedOnly=true&cashtags=cashtags&communityId=communityId&conversationId=conversationId&exactPhrase=exactPhrase&excludeWords=excludeWords&fromUser=fromUser&hashtags=hashtags&inReplyToTweetId=inReplyToTweetId&maxFaves=2&maxId=maxId&maxQuotes=2&maxReplies=2&maxRetweets=2&mentioning=mentioning&minBookmarks=2&minQuotes=2&nativeRetweets=true&q=q&quotes=include&quotesOfTweetId=quotesOfTweetId&replies=include&retweets=include&retweetsOfTweetId=retweetsOfTweetId&sinceId=sinceId&toUser=toUser&url=url&withinTime=withinTime"},
		"XCommunityTweetListByCommunityParams": {xtwitterscraper.XCommunityTweetListByCommunityParams{AnyWords: xtwitterscraper.String("anyWords"), BlueVerifiedOnly: xtwitterscraper.Bool(true), Cashtags: xtwitterscraper.String("cashtags"), ConversationID: xtwitterscraper.String("conversationId"), ExactPhrase: xtwitterscraper.String("exactPhrase"), ExcludeWords: xtwitterscraper.String("excludeWords"), FromUser: xtwitterscraper.String("fromUser"), Hashtags: xtwitterscraper.String("hashtags"), InReplyToTweetID: xtwitterscraper.String("inReplyToTweetId"), MaxFaves: xtwitterscraper.Int(2), MaxID: xtwitterscraper.String("maxId"), MaxQuotes: xtwitterscraper.Int(2), MaxReplies: xtwitterscraper.Int(2), MaxRetweets: xtwitterscraper.Int(2), Mentioning: xtwitterscraper.String("mentioning"), MinBookmarks: xtwitterscraper.Int(2), MinQuotes: xtwitterscraper.Int(2), NativeRetweets: xtwitterscraper.Bool(true), QueryType: xtwitterscraper.XCommunityTweetListByCommunityParamsQueryTypeLatest, Quotes: xtwitterscraper.XCommunityTweetListByCommunityParamsQuotesInclude, QuotesOfTweetID: xtwitterscraper.String("quotesOfTweetId"), Replies: xtwitterscraper.XCommunityTweetListByCommunityParamsRepliesInclude, Retweets: xtwitterscraper.XCommunityTweetListByCommunityParamsRetweetsInclude, RetweetsOfTweetID: xtwitterscraper.String("retweetsOfTweetId"), SinceID: xtwitterscraper.String("sinceId"), ToUser: xtwitterscraper.String("toUser"), URL: xtwitterscraper.String("url"), WithinTime: xtwitterscraper.String("withinTime")}, "anyWords=anyWords&blueVerifiedOnly=true&cashtags=cashtags&conversationId=conversationId&exactPhrase=exactPhrase&excludeWords=excludeWords&fromUser=fromUser&hashtags=hashtags&inReplyToTweetId=inReplyToTweetId&maxFaves=2&maxId=maxId&maxQuotes=2&maxReplies=2&maxRetweets=2&mentioning=mentioning&minBookmarks=2&minQuotes=2&nativeRetweets=true&queryType=Latest&quotes=include&quotesOfTweetId=quotesOfTweetId&replies=include&retweets=include&retweetsOfTweetId=retweetsOfTweetId&sinceId=sinceId&toUser=toUser&url=url&withinTime=withinTime"},
		"XListGetFollowersParams":              {xtwitterscraper.XListGetFollowersParams{After: xtwitterscraper.String("after"), BlueVerifiedOnly: xtwitterscraper.Bool(true), BusinessEmailOnly: xtwitterscraper.Bool(true), EmailDomains: xtwitterscraper.String("emailDomains"), HasEmail: xtwitterscraper.Bool(true), MaxScan: xtwitterscraper.Int(2)}, "after=after&blueVerifiedOnly=true&businessEmailOnly=true&emailDomains=emailDomains&hasEmail=true&maxScan=2"},
		"XListGetMembersParams":                {xtwitterscraper.XListGetMembersParams{After: xtwitterscraper.String("after"), BlueVerifiedOnly: xtwitterscraper.Bool(true), BusinessEmailOnly: xtwitterscraper.Bool(true), EmailDomains: xtwitterscraper.String("emailDomains"), HasEmail: xtwitterscraper.Bool(true), MaxScan: xtwitterscraper.Int(2)}, "after=after&blueVerifiedOnly=true&businessEmailOnly=true&emailDomains=emailDomains&hasEmail=true&maxScan=2"},
		"XListGetTweetsParams":                 {xtwitterscraper.XListGetTweetsParams{AdvancedQuery: xtwitterscraper.String("advancedQuery"), After: xtwitterscraper.String("after"), BoundingBox: xtwitterscraper.String("boundingBox"), CardName: xtwitterscraper.String("cardName"), ConversationID: xtwitterscraper.String("conversationId"), ExcludeSource: xtwitterscraper.String("excludeSource"), Geocode: xtwitterscraper.String("geocode"), InReplyToTweetID: xtwitterscraper.String("inReplyToTweetId"), Keywords: xtwitterscraper.String("keywords"), MaxID: xtwitterscraper.String("maxId"), Near: xtwitterscraper.String("near"), News: xtwitterscraper.Bool(true), Place: xtwitterscraper.String("place"), PlaceCountry: xtwitterscraper.String("placeCountry"), PointRadius: xtwitterscraper.String("pointRadius"), Quotes: xtwitterscraper.XListGetTweetsParamsQuotesInclude, QuotesOfTweetID: xtwitterscraper.String("quotesOfTweetId"), RetweetsOfTweetID: xtwitterscraper.String("retweetsOfTweetId"), Safe: xtwitterscraper.Bool(true), SinceID: xtwitterscraper.String("sinceId"), Source: xtwitterscraper.String("source"), URL: xtwitterscraper.String("url"), Within: xtwitterscraper.String("within"), WithinTime: xtwitterscraper.String("withinTime")}, "advancedQuery=advancedQuery&after=after&boundingBox=boundingBox&cardName=cardName&conversationId=conversationId&excludeSource=excludeSource&geocode=geocode&inReplyToTweetId=inReplyToTweetId&keywords=keywords&maxId=maxId&near=near&news=true&place=place&placeCountry=placeCountry&pointRadius=pointRadius&quotes=include&quotesOfTweetId=quotesOfTweetId&retweetsOfTweetId=retweetsOfTweetId&safe=true&sinceId=sinceId&source=source&url=url&within=within&withinTime=withinTime"},
		"XTweetGetFavoritersParams":            {xtwitterscraper.XTweetGetFavoritersParams{BlueVerifiedOnly: xtwitterscraper.Bool(true), BusinessEmailOnly: xtwitterscraper.Bool(true), EmailDomains: xtwitterscraper.String("emailDomains"), HasEmail: xtwitterscraper.Bool(true)}, "blueVerifiedOnly=true&businessEmailOnly=true&emailDomains=emailDomains&hasEmail=true"},
		"XTweetGetQuotesParams":                {xtwitterscraper.XTweetGetQuotesParams{AdvancedQuery: xtwitterscraper.String("advancedQuery"), BoundingBox: xtwitterscraper.String("boundingBox"), Keywords: xtwitterscraper.String("keywords"), ListID: xtwitterscraper.String("listId"), Place: xtwitterscraper.String("place"), PlaceCountry: xtwitterscraper.String("placeCountry"), PointRadius: xtwitterscraper.String("pointRadius"), QueryType: xtwitterscraper.XTweetGetQuotesParamsQueryTypeLatest}, "advancedQuery=advancedQuery&boundingBox=boundingBox&keywords=keywords&listId=listId&place=place&placeCountry=placeCountry&pointRadius=pointRadius&queryType=Latest"},
		"XTweetGetRepliesParams":               {xtwitterscraper.XTweetGetRepliesParams{After: xtwitterscraper.String("after"), IncludeHidden: xtwitterscraper.Bool(true)}, "after=after&includeHidden=true"},
		"XTweetGetRetweetersParams":            {xtwitterscraper.XTweetGetRetweetersParams{BlueVerifiedOnly: xtwitterscraper.Bool(true), BusinessEmailOnly: xtwitterscraper.Bool(true), EmailDomains: xtwitterscraper.String("emailDomains"), EnrichProfiles: xtwitterscraper.Bool(true), HasEmail: xtwitterscraper.Bool(true)}, "blueVerifiedOnly=true&businessEmailOnly=true&emailDomains=emailDomains&enrichProfiles=true&hasEmail=true"},
		"XTweetGetThreadParams":                {xtwitterscraper.XTweetGetThreadParams{MaxID: xtwitterscraper.String("maxId"), NativeRetweets: xtwitterscraper.Bool(true), SinceID: xtwitterscraper.String("sinceId"), WithinTime: xtwitterscraper.String("withinTime")}, "maxId=maxId&nativeRetweets=true&sinceId=sinceId&withinTime=withinTime"},
		"XTweetSearchParams":                   {xtwitterscraper.XTweetSearchParams{Keywords: xtwitterscraper.String("keywords"), Q: "q"}, "keywords=keywords&q=q"},
		"XUserGetBatchParams":                  {xtwitterscraper.XUserGetBatchParams{BioContains: xtwitterscraper.String("bioContains"), BlueVerifiedOnly: xtwitterscraper.Bool(true), BusinessEmailOnly: xtwitterscraper.Bool(true), EmailDomains: xtwitterscraper.String("emailDomains"), HasEmail: xtwitterscraper.Bool(true), HasLocation: xtwitterscraper.Bool(true), HasWebsite: xtwitterscraper.Bool(true), IDs: "ids", LocationContains: xtwitterscraper.String("locationContains"), MaxFollowing: xtwitterscraper.Int(2), MaxStatuses: xtwitterscraper.Int(2), MinFollowing: xtwitterscraper.Int(2), MinStatuses: xtwitterscraper.Int(2), UsernameContains: xtwitterscraper.String("usernameContains"), Usernames: xtwitterscraper.String("usernames"), VerifiedType: xtwitterscraper.String("verifiedType")}, "bioContains=bioContains&blueVerifiedOnly=true&businessEmailOnly=true&emailDomains=emailDomains&hasEmail=true&hasLocation=true&hasWebsite=true&ids=ids&locationContains=locationContains&maxFollowing=2&maxStatuses=2&minFollowing=2&minStatuses=2&usernameContains=usernameContains&usernames=usernames&verifiedType=verifiedType"},
		"XUserGetFollowersParams":              {xtwitterscraper.XUserGetFollowersParams{BlueVerifiedOnly: xtwitterscraper.Bool(true), BusinessEmailOnly: xtwitterscraper.Bool(true), EmailDomains: xtwitterscraper.String("emailDomains"), EnrichProfiles: xtwitterscraper.Bool(true), HasEmail: xtwitterscraper.Bool(true), MaxScan: xtwitterscraper.Int(2)}, "blueVerifiedOnly=true&businessEmailOnly=true&emailDomains=emailDomains&enrichProfiles=true&hasEmail=true&maxScan=2"},
		"XUserGetFollowersYouKnowParams":       {xtwitterscraper.XUserGetFollowersYouKnowParams{BlueVerifiedOnly: xtwitterscraper.Bool(true), BusinessEmailOnly: xtwitterscraper.Bool(true), EmailDomains: xtwitterscraper.String("emailDomains"), HasEmail: xtwitterscraper.Bool(true)}, "blueVerifiedOnly=true&businessEmailOnly=true&emailDomains=emailDomains&hasEmail=true"},
		"XUserGetFollowingParams":              {xtwitterscraper.XUserGetFollowingParams{BlueVerifiedOnly: xtwitterscraper.Bool(true), BusinessEmailOnly: xtwitterscraper.Bool(true), EmailDomains: xtwitterscraper.String("emailDomains"), EnrichProfiles: xtwitterscraper.Bool(true), HasEmail: xtwitterscraper.Bool(true), MaxScan: xtwitterscraper.Int(2)}, "blueVerifiedOnly=true&businessEmailOnly=true&emailDomains=emailDomains&enrichProfiles=true&hasEmail=true&maxScan=2"},
		"XUserGetRepliesParams":                {xtwitterscraper.XUserGetRepliesParams{AdvancedQuery: xtwitterscraper.String("advancedQuery"), BoundingBox: xtwitterscraper.String("boundingBox"), Keywords: xtwitterscraper.String("keywords"), ListID: xtwitterscraper.String("listId"), Mode: xtwitterscraper.XUserGetRepliesParamsModeStandard, Place: xtwitterscraper.String("place"), PlaceCountry: xtwitterscraper.String("placeCountry"), PointRadius: xtwitterscraper.String("pointRadius"), SinceTime: xtwitterscraper.String("sinceTime"), UntilTime: xtwitterscraper.String("untilTime")}, "advancedQuery=advancedQuery&boundingBox=boundingBox&keywords=keywords&listId=listId&mode=standard&place=place&placeCountry=placeCountry&pointRadius=pointRadius&sinceTime=sinceTime&untilTime=untilTime"},
		"XUserGetSearchParams":                 {xtwitterscraper.XUserGetSearchParams{BlueVerifiedOnly: xtwitterscraper.Bool(true), BusinessEmailOnly: xtwitterscraper.Bool(true), EmailDomains: xtwitterscraper.String("emailDomains"), EnrichProfiles: xtwitterscraper.Bool(true), HasEmail: xtwitterscraper.Bool(true), Q: "q"}, "blueVerifiedOnly=true&businessEmailOnly=true&emailDomains=emailDomains&enrichProfiles=true&hasEmail=true&q=q"},
		"XUserGetTweetsParams":                 {xtwitterscraper.XUserGetTweetsParams{AdvancedQuery: xtwitterscraper.String("advancedQuery"), BoundingBox: xtwitterscraper.String("boundingBox"), Keywords: xtwitterscraper.String("keywords"), ListID: xtwitterscraper.String("listId"), Mode: xtwitterscraper.XUserGetTweetsParamsModeStandard, Place: xtwitterscraper.String("place"), PlaceCountry: xtwitterscraper.String("placeCountry"), PointRadius: xtwitterscraper.String("pointRadius"), SinceTime: xtwitterscraper.String("sinceTime"), UntilTime: xtwitterscraper.String("untilTime")}, "advancedQuery=advancedQuery&boundingBox=boundingBox&keywords=keywords&listId=listId&mode=standard&place=place&placeCountry=placeCountry&pointRadius=pointRadius&sinceTime=sinceTime&untilTime=untilTime"},
		"XUserGetVerifiedFollowersParams":      {xtwitterscraper.XUserGetVerifiedFollowersParams{BlueVerifiedOnly: xtwitterscraper.Bool(true), BusinessEmailOnly: xtwitterscraper.Bool(true), EmailDomains: xtwitterscraper.String("emailDomains"), EnrichProfiles: xtwitterscraper.Bool(true), HasEmail: xtwitterscraper.Bool(true), MaxScan: xtwitterscraper.Int(2)}, "blueVerifiedOnly=true&businessEmailOnly=true&emailDomains=emailDomains&enrichProfiles=true&hasEmail=true&maxScan=2"},
	}
	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			query, err := test.params.URLQuery()
			if err != nil || query.Encode() != test.want {
				t.Fatalf("query = %s, %v, want %s", query.Encode(), err, test.want)
			}
		})
	}
}

func TestNewBodyParamsReachTheBody(t *testing.T) {
	cases := map[string]struct {
		params interface{ MarshalJSON() ([]byte, error) }
		want   map[string]any
	}{
		"ExtractionEstimateCostParams": {xtwitterscraper.ExtractionEstimateCostParams{BusinessEmailOnly: xtwitterscraper.Bool(true), EmailDomains: xtwitterscraper.String("emailDomains"), EnrichProfiles: xtwitterscraper.Bool(true), HasEmail: xtwitterscraper.Bool(true)}, map[string]any{"businessEmailOnly": true, "emailDomains": "emailDomains", "enrichProfiles": true, "hasEmail": true}},
		"ExtractionRunParams":          {xtwitterscraper.ExtractionRunParams{BusinessEmailOnly: xtwitterscraper.Bool(true), EmailDomains: xtwitterscraper.String("emailDomains"), EnrichProfiles: xtwitterscraper.Bool(true), HasEmail: xtwitterscraper.Bool(true)}, map[string]any{"businessEmailOnly": true, "emailDomains": "emailDomains", "enrichProfiles": true, "hasEmail": true}},
		"MonitorNewParams":             {xtwitterscraper.MonitorNewParams{Name: xtwitterscraper.String("name")}, map[string]any{"name": "name"}},
		"MonitorUpdateParams":          {xtwitterscraper.MonitorUpdateParams{Name: xtwitterscraper.String("name")}, map[string]any{"name": "name"}},
		"MonitorKeywordNewParams":      {xtwitterscraper.MonitorKeywordNewParams{Name: xtwitterscraper.String("name")}, map[string]any{"name": "name"}},
		"MonitorKeywordUpdateParams":   {xtwitterscraper.MonitorKeywordUpdateParams{Name: xtwitterscraper.String("name")}, map[string]any{"name": "name"}},
	}
	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			data, err := test.params.MarshalJSON()
			if err != nil {
				t.Fatal(err)
			}
			var body map[string]any
			if err := json.Unmarshal(data, &body); err != nil {
				t.Fatal(err)
			}
			for key, want := range test.want {
				if body[key] != want {
					t.Errorf("%s = %#v, want %#v in %s", key, body[key], want, data)
				}
			}
		})
	}
}

func TestSpecFieldsOf20261105ReachTheRequest(t *testing.T) {
	day := time.Date(2020, 1, 2, 0, 0, 0, 0, time.UTC)
	exportQuery, err := xtwitterscraper.ExtractionExportResultsParams{
		Format:       xtwitterscraper.ExtractionExportResultsParamsFormatCsv,
		Compare:      xtwitterscraper.String("run-0"),
		Change:       xtwitterscraper.ExtractionExportResultsParamsChangeRemoved,
		HasWebsite:   xtwitterscraper.Bool(true),
		HideFake:     xtwitterscraper.Bool(true),
		HideInactive: xtwitterscraper.Bool(false),
		JoinedSince:  xtwitterscraper.Time(day),
		JoinedUntil:  xtwitterscraper.Time(day.AddDate(4, 0, 0)),
		Location:     xtwitterscraper.String("Berlin"),
	}.URLQuery()
	want := "change=removed&compare=run-0&format=csv&hasWebsite=true&hideFake=true&hideInactive=false&joinedSince=2020-01-02&joinedUntil=2024-01-02&location=Berlin"
	if err != nil || exportQuery.Encode() != want {
		t.Fatalf("export query = %s, %v, want %s", exportQuery.Encode(), err, want)
	}
	searchQuery, err := xtwitterscraper.XTweetSearchParams{Q: "launch", TextOnlyTerms: xtwitterscraper.String("rocket")}.URLQuery()
	if err != nil || searchQuery.Get("textOnlyTerms") != "rocket" {
		t.Fatalf("search query = %s, %v, want textOnlyTerms=rocket", searchQuery.Encode(), err)
	}
	bodies := map[string]interface{ MarshalJSON() ([]byte, error) }{
		"textOnlyTerms run":      xtwitterscraper.ExtractionRunParams{TextOnlyTerms: xtwitterscraper.String("rocket")},
		"textOnlyTerms estimate": xtwitterscraper.ExtractionEstimateCostParams{TextOnlyTerms: xtwitterscraper.String("rocket")},
		"analysis images":        xtwitterscraper.XAnalysisSentimentParams{Analysis: xtwitterscraper.XAnalysisSentimentParamsAnalysis{Images: xtwitterscraper.Bool(false)}},
	}
	wants := map[string]string{
		"textOnlyTerms run":      `"textOnlyTerms":"rocket"`,
		"textOnlyTerms estimate": `"textOnlyTerms":"rocket"`,
		"analysis images":        `"analysis":{"images":false}`,
	}
	for name, params := range bodies {
		data, err := params.MarshalJSON()
		if err != nil || !strings.Contains(string(data), wants[name]) {
			t.Errorf("%s body = %s, %v, want %s", name, data, err, wants[name])
		}
	}
}

func TestSpecFieldsOf20261105DecodeFromTheResponse(t *testing.T) {
	var job xtwitterscraper.ExtractionJob
	if err := job.UnmarshalJSON([]byte(`{"id":"ext_1","url":"https://dashboard.xquik.com/activity/a1","toolType":"follower_explorer","status":"completed","totalResults":3,"createdAt":"2026-11-05T00:00:00Z"}`)); err != nil || job.URL != "https://dashboard.xquik.com/activity/a1" {
		t.Fatalf("job url = %q, %v", job.URL, err)
	}
	var item xtwitterscraper.RadarItemMetadata
	if err := item.UnmarshalJSON([]byte(`{"sentiment":"positive","windows":["now","day"]}`)); err != nil || item.Sentiment != "positive" || strings.Join(item.Windows, ",") != "now,day" {
		t.Fatalf("radar metadata = %q %v, %v", item.Sentiment, item.Windows, err)
	}
	var insights xtwitterscraper.XUserGetInsightsResponseInsights
	if err := insights.UnmarshalJSON([]byte(`{"medianViews":800791.5}`)); err != nil || insights.MedianViews != 800791.5 {
		t.Fatalf("medianViews = %v, %v", insights.MedianViews, err)
	}
	var context xtwitterscraper.XAnalysisSentimentResponseResultAnalysisContextAvailability
	if err := context.UnmarshalJSON([]byte(`{"media":"images"}`)); err != nil || context.Media != xtwitterscraper.XAnalysisSentimentResponseResultAnalysisContextAvailabilityMediaImages {
		t.Fatalf("context media = %q, %v", context.Media, err)
	}
}
