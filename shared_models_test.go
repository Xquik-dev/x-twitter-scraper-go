// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

package xtwitterscraper_test

import (
	"context"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/Xquik-dev/x-twitter-scraper-go"
	"github.com/Xquik-dev/x-twitter-scraper-go/shared"
)

type fieldValue struct {
	path string
	want any
}

func valueAt(root any, path string) (any, bool) {
	value := reflect.ValueOf(root)
	for _, part := range strings.Split(path, ".") {
		name, rest, indexed := strings.Cut(part, "[")
		value = reflect.Indirect(value).FieldByName(name)
		if !value.IsValid() {
			return nil, false
		}
		if indexed {
			index, err := strconv.Atoi(strings.TrimSuffix(rest, "]"))
			if err != nil || index >= value.Len() {
				return nil, false
			}
			value = value.Index(index)
		}
	}
	switch typed := value.Interface().(type) {
	case time.Time:
		return typed.UTC().Format(time.RFC3339Nano), true
	}
	if value.Kind() == reflect.String {
		return value.String(), true
	}
	return value.Interface(), true
}

func expectValues(t *testing.T, root any, values []fieldValue) {
	t.Helper()
	for _, value := range values {
		got, ok := valueAt(root, value.path)
		if !ok || !reflect.DeepEqual(got, value.want) {
			t.Errorf("%s = %#v, want %#v", value.path, got, value.want)
		}
	}
}

func TestUserProfileDecodesEveryField(t *testing.T) {
	payload, err := fixtures.ReadFile("testdata/fixtures/UserProfile.json")
	if err != nil {
		t.Fatal(err)
	}
	var profile shared.UserProfile
	if err := profile.UnmarshalJSON(payload); err != nil {
		t.Fatal(err)
	}
	expectDecoded(t, profile)
	expectValues(t, profile, []fieldValue{
		{"AccountBasedInAccurate", true},
		{"AccountBasedInUnavailable", true},
		{"AdvertiserAccountServiceLevels", []string{"smb", "media_studio"}},
		{"AdvertiserAccountType", "none"},
		{"AffiliateUsername", "BBCNews"},
		{"Birthdate.Day", int64(1)},
		{"Birthdate.Month", int64(10)},
		{"Birthdate.Year", int64(1958)},
		{"BusinessProfileState", "none"},
		{"ConnectedVia", "United States App Store"},
		{"ConnectedViaAccurate", true},
		{"ContactInfo.Emails[0].Domain", "example.com"},
		{"ContactInfo.Emails[0].Source", "bio"},
		{"ContactInfo.Emails[0].Value", "press@example.com"},
		{"ContactInfo.Emails[0].WrittenAs", "press at example dot com"},
		{"ContactInfo.Links[0].Network", "website"},
		{"ContactInfo.Links[0].Source", "url"},
		{"ContactInfo.Links[0].URL", "http://www.nasa.gov/"},
		{"ContactInfo.Phones[0].Source", "location"},
		{"ContactInfo.Phones[0].Value", "+12025550123"},
		{"ContactInfo.Phones[0].WrittenAs", "202 555 0123"},
		{"DefaultProfile", true},
		{"DefaultProfileImage", true},
		{"Email", "press@example.com"},
		{"FastFollowersCount", int64(12)},
		{"GeoEnabled", true},
		{"HasExtendedProfile", true},
		{"IsTranslationEnabled", true},
		{"ListedCount", int64(3051)},
		{"NormalFollowersCount", int64(92376175)},
		{"ProfilePictureOriginal", "https://pbs.twimg.com/profile_images/1321163587679784960/0ZxKlEKB.jpg"},
		{"RedirectsToUsername", "spacexsi"},
		{"SubscribersCount", int64(7)},
		{"UsernameChanges.Count", int64(2)},
		{"UsernameChanges.LastChangedAt", "2016-05-14T00:22:47.057Z"},
		{"VerificationReason.Entities[0].FromIndex", int64(91)},
		{"VerificationReason.Entities[0].ToIndex", int64(101)},
		{"VerificationReason.Entities[0].URL", "https://help.twitter.com/en/rules-and-policies/profile-labels"},
		{"VerificationReason.Entities[0].URLType", "ExternalUrl"},
		{"VerificationReason.Text", "This account is verified because it is a government or multilateral organization account.  Learn more"},
	})
}

var tweetFields = []fieldValue{
	{"AuthorUnavailable.Message", "This account is suspended."},
	{"AuthorUnavailable.Reason", "Suspended"},
	{"CashtagAttachments[0].CandleSize", "5m"},
	{"CashtagAttachments[0].Cashtag", "$TSLA"},
	{"CashtagAttachments[0].EndTime", "2026-10-09T13:42:08.796Z"},
	{"CashtagAttachments[0].SizeVariant", int64(1)},
	{"CashtagAttachments[0].StartTime", "2026-10-09T13:30:00Z"},
	{"CashtagAttachments[0].Timeframe", "1D"},
	{"CommunityName", "Build in Public"},
	{"ConversationControl.AllowedCountryCodes", []string{"US", "CA"}},
	{"ConversationControl.Mode", "Community"},
	{"ConversationControl.OwnerID", "11348282"},
	{"Coordinates.Coordinates", []float64{-122.4194, 37.7749}},
	{"Coordinates.Type", "Point"},
	{"Entities.Hashtags[0].Tag.ID", "hashtag-nasa"},
	{"Entities.Hashtags[0].Tag.Info.Info.Active", true},
	{"Entities.Hashtags[0].Tag.Info.Info.Name", "NASA"},
	{"Entities.Hashtags[0].Tag.Info.Info.Ticker", "NASA"},
	{"Entities.Smarttags[0].Tag.ID", "smarttag-1"},
	{"Entities.Smarttags[0].Tag.Info.Info.Active", true},
	{"Entities.Symbols[0].Tag.ID", "symbol-tsla"},
	{"Entities.Symbols[0].Tag.Info.Info.Active", true},
	{"Entities.Timestamps[0].Tag.ID", "timestamp-75"},
	{"Entities.Timestamps[0].Tag.Info.Info.Active", true},
	{"ExclusiveContent.CreatorUserName", "NASA"},
	{"ExclusiveContent.SubscribeURL", "https://x.com/NASA/creator-subscriptions/subscribe"},
	{"Geo.Coordinates", []float64{37.7749, -122.4194}},
	{"Geo.Type", "Point"},
	{"GrokTranslatedPost", map[string]any{"language": "es", "text": "La galaxia espiral."}},
	{"HasCommunityNotes", true},
	{"Interstitial.DisplayType", "NonCompliant"},
	{"Interstitial.Text.Entities[0].FromIndex", int64(164)},
	{"Interstitial.Text.Entities[0].Ref.ScreenName", "XSafety"},
	{"Interstitial.Text.Entities[0].Ref.Text", "Learn more"},
	{"Interstitial.Text.Entities[0].Ref.Type", "TimelineUrl"},
	{"Interstitial.Text.Entities[0].Ref.URL", "https://help.x.com/rules-and-policies/public-interest"},
	{"Interstitial.Text.Entities[0].Ref.URLType", "ExternalUrl"},
	{"Interstitial.Text.Entities[0].ToIndex", int64(174)},
	{"Interstitial.Text.Rtl", true},
	{"Interstitial.Text.Text", "This Post violated the X Rules about glorifying violence. However, X has determined that it may be in the public's interest for the Post to remain accessible. Learn more"},
	{"IsPinned", true},
	{"IsRetweet", true},
	{"Media[0].GrokImageAnnotation.Prompt", "A spiral galaxy"},
	{"Media[0].GrokImageAnnotation.UpsampledPrompt", "A detailed spiral galaxy with blue and red arms"},
	{"Media[0].SourceUser.Username", "NASAJPL"},
	{"NoteTweet.NoteID", "note-77"},
	{"ParentTweet.ID", "2102761519985332440"},
	{"ParentTweet.Text", "The post this one replies to."},
	{"QuotedTweetPermalink.DisplayURL", "x.com/NASA/status/2102…"},
	{"QuotedTweetPermalink.ExpandedURL", "https://x.com/NASA/status/2102761519985332000"},
	{"QuotedTweetPermalink.URL", "https://t.co/quoted123"},
	{"QuotedTweetUnavailable.Reason", "Protected"},
	{"SportsContext.Competitors[0].Color", "007A33"},
	{"SportsContext.Competitors[0].IsActive", true},
	{"SportsContext.Competitors[0].IsWinner", true},
	{"SportsContext.League", "NBA"},
	{"SubscriberBadge.CreatorID", "44196397"},
	{"TextHighlights[0].EndIndex", int64(13)},
	{"TextHighlights[0].StartIndex", int64(5)},
	{"Tombstone.Text.Entities[0].Ref.ScreenName", "XSupport"},
	{"Tombstone.Text.Entities[0].Ref.Text", "Learn more"},
	{"UnavailableAuthorID", "783214"},
}

func TestSearchTweetDecodesEveryField(t *testing.T) {
	payload, err := fixtures.ReadFile("testdata/fixtures/SearchTweet.json")
	if err != nil {
		t.Fatal(err)
	}
	var tweet shared.SearchTweet
	if err := tweet.UnmarshalJSON(payload); err != nil {
		t.Fatal(err)
	}
	expectDecoded(t, tweet)
	expectValues(t, tweet, tweetFields)
}

func TestXTweetGetFixture(t *testing.T) {
	client, request := fixtureClient(t, "lookupTweet.json")
	res, err := client.X.Tweets.Get(context.Background(), "2102761519985332442")
	if err != nil {
		t.Fatal(err)
	}
	request.expect(t, http.MethodGet, "/api/v1/x/tweets/2102761519985332442", "")
	expectDecoded(t, res)
	expectValues(t, res.Tweet, tweetFields)
	expectValues(t, res, []fieldValue{{"Author.Followers", int64(92376233)}, {"Author.Verified", true}})
}

func TestReadsSendFieldsAndOmit(t *testing.T) {
	fields := xtwitterscraper.String("id,text,author.username")
	omit := xtwitterscraper.String("entities,media")
	cases := map[string]interface{ URLQuery() (url.Values, error) }{
		"TrendListParams":                      xtwitterscraper.TrendListParams{Fields: fields, Omit: omit},
		"XAutocompleteSearchParams":            xtwitterscraper.XAutocompleteSearchParams{Q: "nasa", Fields: fields, Omit: omit},
		"XCommunityFindParams":                 xtwitterscraper.XCommunityFindParams{Q: "nasa", Fields: fields, Omit: omit},
		"XCommunityGetMembersParams":           xtwitterscraper.XCommunityGetMembersParams{Fields: fields, Omit: omit},
		"XCommunityGetModeratorsParams":        xtwitterscraper.XCommunityGetModeratorsParams{Fields: fields, Omit: omit},
		"XCommunityGetSearchParams":            xtwitterscraper.XCommunityGetSearchParams{CommunityID: "nasa", Q: "nasa", Fields: fields, Omit: omit},
		"XCommunityListPopularParams":          xtwitterscraper.XCommunityListPopularParams{Fields: fields, Omit: omit},
		"XCommunityListSuggestedParams":        xtwitterscraper.XCommunityListSuggestedParams{Fields: fields, Omit: omit},
		"XCommunityListTopicsParams":           xtwitterscraper.XCommunityListTopicsParams{Fields: fields, Omit: omit},
		"XCommunityMediaListParams":            xtwitterscraper.XCommunityMediaListParams{Fields: fields, Omit: omit},
		"XCommunitySearchMembersParams":        xtwitterscraper.XCommunitySearchMembersParams{Q: "nasa", Fields: fields, Omit: omit},
		"XCommunityTweetListByCommunityParams": xtwitterscraper.XCommunityTweetListByCommunityParams{Fields: fields, Omit: omit},
		"XCommunityTweetListParams":            xtwitterscraper.XCommunityTweetListParams{CommunityID: "nasa", Q: "nasa", Fields: fields, Omit: omit},
		"XFollowerCheckParams":                 xtwitterscraper.XFollowerCheckParams{Source: "nasa", Target: "nasa", Fields: fields, Omit: omit},
		"XGetHashflagsParams":                  xtwitterscraper.XGetHashflagsParams{Fields: fields, Omit: omit},
		"XGetNewsStoryVersionsParams":          xtwitterscraper.XGetNewsStoryVersionsParams{Fields: fields, Omit: omit},
		"XGetTrendLocationsParams":             xtwitterscraper.XGetTrendLocationsParams{Fields: fields, Omit: omit},
		"XGetTrendsHistoryParams":              xtwitterscraper.XGetTrendsHistoryParams{Fields: fields, Omit: omit},
		"XGetTrendsParams":                     xtwitterscraper.XGetTrendsParams{Fields: fields, Omit: omit},
		"XListGetFollowersParams":              xtwitterscraper.XListGetFollowersParams{Fields: fields, Omit: omit},
		"XListGetMembersParams":                xtwitterscraper.XListGetMembersParams{Fields: fields, Omit: omit},
		"XListGetParams":                       xtwitterscraper.XListGetParams{Fields: fields, Omit: omit},
		"XListGetSearchParams":                 xtwitterscraper.XListGetSearchParams{Q: "nasa", Fields: fields, Omit: omit},
		"XListGetTweetsParams":                 xtwitterscraper.XListGetTweetsParams{Fields: fields, Omit: omit},
		"XSearchPlacesParams":                  xtwitterscraper.XSearchPlacesParams{Q: "nasa", Fields: fields, Omit: omit},
		"XTweetCheckRepostParams":              xtwitterscraper.XTweetCheckRepostParams{User: "nasa", Fields: fields, Omit: omit},
		"XTweetEmbedParams":                    xtwitterscraper.XTweetEmbedParams{Fields: fields, Omit: omit},
		"XTweetGetFavoritersParams":            xtwitterscraper.XTweetGetFavoritersParams{Fields: fields, Omit: omit},
		"XTweetGetHiddenRepliesParams":         xtwitterscraper.XTweetGetHiddenRepliesParams{Fields: fields, Omit: omit},
		"XTweetGetInsightsParams":              xtwitterscraper.XTweetGetInsightsParams{Fields: fields, Omit: omit},
		"XTweetGetQuotesParams":                xtwitterscraper.XTweetGetQuotesParams{Fields: fields, Omit: omit},
		"XTweetGetRepliesParams":               xtwitterscraper.XTweetGetRepliesParams{Fields: fields, Omit: omit},
		"XTweetGetRetweetersParams":            xtwitterscraper.XTweetGetRetweetersParams{Fields: fields, Omit: omit},
		"XTweetGetSubtitlesParams":             xtwitterscraper.XTweetGetSubtitlesParams{Fields: fields, Omit: omit},
		"XTweetGetThreadParams":                xtwitterscraper.XTweetGetThreadParams{Fields: fields, Omit: omit},
		"XTweetListParams":                     xtwitterscraper.XTweetListParams{IDs: "nasa", Fields: fields, Omit: omit},
		"XTweetListTopParams":                  xtwitterscraper.XTweetListTopParams{Fields: fields, Omit: omit},
		"XTweetResolveLinksParams":             xtwitterscraper.XTweetResolveLinksParams{URLs: "nasa", Fields: fields, Omit: omit},
		"XTweetSearchParams":                   xtwitterscraper.XTweetSearchParams{Q: "nasa", Fields: fields, Omit: omit},
		"XTweetTranslateParams":                xtwitterscraper.XTweetTranslateParams{Language: "nasa", Fields: fields, Omit: omit},
		"XUserGetAffiliatesParams":             xtwitterscraper.XUserGetAffiliatesParams{Fields: fields, Omit: omit},
		"XUserGetArticlesParams":               xtwitterscraper.XUserGetArticlesParams{Fields: fields, Omit: omit},
		"XUserGetBatchParams":                  xtwitterscraper.XUserGetBatchParams{IDs: "nasa", Fields: fields, Omit: omit},
		"XUserGetBatchTweetsParams":            xtwitterscraper.XUserGetBatchTweetsParams{Fields: fields, Omit: omit},
		"XUserGetFollowersParams":              xtwitterscraper.XUserGetFollowersParams{Fields: fields, Omit: omit},
		"XUserGetFollowersYouKnowParams":       xtwitterscraper.XUserGetFollowersYouKnowParams{Fields: fields, Omit: omit},
		"XUserGetFollowingParams":              xtwitterscraper.XUserGetFollowingParams{Fields: fields, Omit: omit},
		"XUserGetHighlightsParams":             xtwitterscraper.XUserGetHighlightsParams{Fields: fields, Omit: omit},
		"XUserGetInsightsParams":               xtwitterscraper.XUserGetInsightsParams{Fields: fields, Omit: omit},
		"XUserGetLikesParams":                  xtwitterscraper.XUserGetLikesParams{Fields: fields, Omit: omit},
		"XUserGetListMembershipsParams":        xtwitterscraper.XUserGetListMembershipsParams{Fields: fields, Omit: omit},
		"XUserGetListsParams":                  xtwitterscraper.XUserGetListsParams{Fields: fields, Omit: omit},
		"XUserGetLiveParams":                   xtwitterscraper.XUserGetLiveParams{Fields: fields, Omit: omit},
		"XUserGetMediaParams":                  xtwitterscraper.XUserGetMediaParams{Fields: fields, Omit: omit},
		"XUserGetMentionsParams":               xtwitterscraper.XUserGetMentionsParams{Fields: fields, Omit: omit},
		"XUserGetRepliesParams":                xtwitterscraper.XUserGetRepliesParams{Fields: fields, Omit: omit},
		"XUserGetSearchParams":                 xtwitterscraper.XUserGetSearchParams{Q: "nasa", Fields: fields, Omit: omit},
		"XUserGetSimilarParams":                xtwitterscraper.XUserGetSimilarParams{Fields: fields, Omit: omit},
		"XUserGetSubscriptionsParams":          xtwitterscraper.XUserGetSubscriptionsParams{Fields: fields, Omit: omit},
		"XUserGetTweetsParams":                 xtwitterscraper.XUserGetTweetsParams{Fields: fields, Omit: omit},
		"XUserGetVerifiedFollowersParams":      xtwitterscraper.XUserGetVerifiedFollowersParams{Fields: fields, Omit: omit},
		"XUserTranslateBioParams":              xtwitterscraper.XUserTranslateBioParams{Language: "nasa", Fields: fields, Omit: omit},
	}
	for name, params := range cases {
		t.Run(name, func(t *testing.T) {
			query, err := params.URLQuery()
			if err != nil {
				t.Fatal(err)
			}
			if query.Get("fields") != "id,text,author.username" || query.Get("omit") != "entities,media" {
				t.Fatalf("query = %s", query.Encode())
			}
		})
	}
}
