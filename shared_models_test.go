// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

package xtwitterscraper_test

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/Xquik-dev/x-twitter-scraper-go"
	"github.com/Xquik-dev/x-twitter-scraper-go/packages/respjson"
	"github.com/Xquik-dev/x-twitter-scraper-go/shared"
)

type fieldValue struct {
	path string
	want reflect.Value
}

func field[T any](path string, want T) fieldValue {
	return fieldValue{path, reflect.ValueOf(want)}
}

func valueAt(value reflect.Value, path string) (reflect.Value, bool) {
	for _, part := range strings.Split(path, ".") {
		name, rest, indexed := strings.Cut(part, "[")
		value = reflect.Indirect(value).FieldByName(name)
		if !value.IsValid() {
			return reflect.Value{}, false
		}
		if indexed {
			index, err := strconv.Atoi(strings.TrimSuffix(rest, "]"))
			if err != nil || index >= value.Len() {
				return reflect.Value{}, false
			}
			value = value.Index(index)
		}
	}
	if timestamp, ok := reflect.TypeAssert[time.Time](value); ok {
		return reflect.ValueOf(timestamp.UTC().Format(time.RFC3339Nano)), true
	}
	if value.Kind() == reflect.String {
		return reflect.ValueOf(value.String()), true
	}
	return value, true
}

func expectValues[T any](t *testing.T, root T, values []fieldValue) {
	t.Helper()
	for _, value := range values {
		got, ok := valueAt(reflect.ValueOf(root), value.path)
		if !ok || !reflect.DeepEqual(got.Interface(), value.want.Interface()) {
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
		field("AccountBasedInAccurate", true),
		field("AccountBasedInUnavailable", true),
		field("AdvertiserAccountServiceLevels", []string{"smb", "media_studio"}),
		field("AdvertiserAccountType", "none"),
		field("AffiliateUsername", "BBCNews"),
		field("Birthdate.Day", int64(1)),
		field("Birthdate.Month", int64(10)),
		field("Birthdate.Year", int64(1958)),
		field("BusinessProfileState", "none"),
		field("ConnectedVia", "United States App Store"),
		field("ConnectedViaAccurate", true),
		field("ContactInfo.Emails[0].Domain", "example.com"),
		field("ContactInfo.Emails[0].Source", "bio"),
		field("ContactInfo.Emails[0].Value", "press@example.com"),
		field("ContactInfo.Emails[0].WrittenAs", "press at example dot com"),
		field("ContactInfo.Links[0].Network", "website"),
		field("ContactInfo.Links[0].Source", "url"),
		field("ContactInfo.Links[0].URL", "http://www.nasa.gov/"),
		field("ContactInfo.Phones[0].Source", "location"),
		field("ContactInfo.Phones[0].Value", "+12025550123"),
		field("ContactInfo.Phones[0].WrittenAs", "202 555 0123"),
		field("DefaultProfile", true),
		field("DefaultProfileImage", true),
		field("Email", "press@example.com"),
		field("FastFollowersCount", int64(12)),
		field("GeoEnabled", true),
		field("HasExtendedProfile", true),
		field("IdentityVerification.OverrideVerifiedYear", int64(2009)),
		field("IsTranslationEnabled", true),
		field("ListedCount", int64(3051)),
		field("NormalFollowersCount", int64(92376175)),
		field("ProfilePictureOriginal", "https://pbs.twimg.com/profile_images/1321163587679784960/0ZxKlEKB.jpg"),
		field("RedirectsToUsername", "spacexsi"),
		field("SubscribersCount", int64(7)),
		field("UsernameChanges.Count", int64(2)),
		field("UsernameChanges.LastChangedAt", "2016-05-14T00:22:47.057Z"),
		field("VerificationReason.Entities[0].FromIndex", int64(91)),
		field("VerificationReason.Entities[0].ToIndex", int64(101)),
		field("VerificationReason.Entities[0].URL", "https://help.twitter.com/en/rules-and-policies/profile-labels"),
		field("VerificationReason.Entities[0].URLType", "ExternalUrl"),
		field("VerificationReason.Text", "This account is verified because it is a government or multilateral organization account.  Learn more"),
	})
}

var tweetFields = []fieldValue{
	field("AuthorUnavailable.Message", "This account is suspended."),
	field("AuthorUnavailable.Reason", "Suspended"),
	field("CashtagAttachments[0].CandleSize", "5m"),
	field("CashtagAttachments[0].Cashtag", "$TSLA"),
	field("CashtagAttachments[0].EndTime", "2026-10-09T13:42:08.796Z"),
	field("CashtagAttachments[0].SizeVariant", int64(1)),
	field("CashtagAttachments[0].StartTime", "2026-10-09T13:30:00Z"),
	field("CashtagAttachments[0].Timeframe", "1D"),
	field("CommunityName", "Build in Public"),
	field("ConversationControl.AllowedCountryCodes", []string{"US", "CA"}),
	field("ConversationControl.Mode", "Community"),
	field("ConversationControl.OwnerID", "11348282"),
	field("Coordinates.Coordinates", []float64{-122.4194, 37.7749}),
	field("Coordinates.Type", "Point"),
	field("Entities.Hashtags[0].Tag.ID", "hashtag-nasa"),
	field("Entities.Hashtags[0].Tag.Info.Info.Active", true),
	field("Entities.Hashtags[0].Tag.Info.Info.Name", "NASA"),
	field("Entities.Hashtags[0].Tag.Info.Info.Ticker", "NASA"),
	field("Entities.Smarttags[0].Tag.ID", "smarttag-1"),
	field("Entities.Smarttags[0].Tag.Info.Info.Active", true),
	field("Entities.Symbols[0].Tag.ID", "symbol-tsla"),
	field("Entities.Symbols[0].Tag.Info.Info.Active", true),
	field("Entities.Timestamps[0].Tag.ID", "timestamp-75"),
	field("Entities.Timestamps[0].Tag.Info.Info.Active", true),
	field("ExclusiveContent.CreatorUserName", "NASA"),
	field("ExclusiveContent.SubscribeURL", "https://x.com/NASA/creator-subscriptions/subscribe"),
	field("Geo.Coordinates", []float64{37.7749, -122.4194}),
	field("Geo.Type", "Point"),
	field("GrokTranslatedPost", xtwitterscraper.PublicMetadata{"language": json.RawMessage(`"es"`), "text": json.RawMessage(`"La galaxia espiral."`)}),
	field("HasCommunityNotes", true),
	field("Interstitial.DisplayType", "NonCompliant"),
	field("Interstitial.Text.Entities[0].FromIndex", int64(164)),
	field("Interstitial.Text.Entities[0].Ref.ScreenName", "XSafety"),
	field("Interstitial.Text.Entities[0].Ref.Text", "Learn more"),
	field("Interstitial.Text.Entities[0].Ref.Type", "TimelineUrl"),
	field("Interstitial.Text.Entities[0].Ref.URL", "https://help.x.com/rules-and-policies/public-interest"),
	field("Interstitial.Text.Entities[0].Ref.URLType", "ExternalUrl"),
	field("Interstitial.Text.Entities[0].ToIndex", int64(174)),
	field("Interstitial.Text.Rtl", true),
	field("Interstitial.Text.Text", "This Post violated the X Rules about glorifying violence. However, X has determined that it may be in the public's interest for the Post to remain accessible. Learn more"),
	field("IsPinned", true),
	field("IsRetweet", true),
	field("Media[0].GrokImageAnnotation.Prompt", "A spiral galaxy"),
	field("Media[0].GrokImageAnnotation.UpsampledPrompt", "A detailed spiral galaxy with blue and red arms"),
	field("Media[0].SourceUser.Username", "NASAJPL"),
	field("NoteTweet.NoteID", "note-77"),
	field("ParentTweet.ID", "2102761519985332440"),
	field("ParentTweet.Text", "The post this one replies to."),
	field("QuotedTweetPermalink.DisplayURL", "x.com/NASA/status/2102…"),
	field("QuotedTweetPermalink.ExpandedURL", "https://x.com/NASA/status/2102761519985332000"),
	field("QuotedTweetPermalink.URL", "https://t.co/quoted123"),
	field("QuotedTweetUnavailable.Reason", "Protected"),
	field("SoftIntervention.DisplayType", "Fosnr"),
	field("SoftIntervention.EngagementNudge", true),
	field("SoftIntervention.Text.Entities[0].FromIndex", int64(0)),
	field("SoftIntervention.Text.Entities[0].Ref.Type", "TimelineUrl"),
	field("SoftIntervention.Text.Entities[0].Ref.URL", "https://help.x.com/rules-and-policies/public-interest"),
	field("SoftIntervention.Text.Entities[0].Ref.URLType", "ExternalUrl"),
	field("SoftIntervention.Text.Entities[0].ToIndex", int64(4)),
	field("SoftIntervention.Text.Rtl", true),
	field("SoftIntervention.Text.Text", "This Post may violate X's rules against Abuse."),
	field("SoftIntervention.URL.URL", "https://help.x.com/rules-and-policies/enforcement-options"),
	field("SoftIntervention.URL.URLType", "ExternalUrl"),
	field("SportsContext.Competitors[0].Color", "007A33"),
	field("SportsContext.Competitors[0].IsActive", true),
	field("SportsContext.Competitors[0].IsWinner", true),
	field("SportsContext.League", "NBA"),
	field("SubscriberBadge.CreatorID", "44196397"),
	field("TextHighlights[0].EndIndex", int64(13)),
	field("TextHighlights[0].StartIndex", int64(5)),
	field("Tombstone.Text.Entities[0].Ref.ScreenName", "XSupport"),
	field("Tombstone.Text.Entities[0].Ref.Text", "Learn more"),
	field("UnavailableAuthorID", "783214"),
}

var _ = []shared.SoftIntervention{
	xtwitterscraper.TweetDetail{}.SoftIntervention,
	shared.SearchTweet{}.SoftIntervention,
	shared.EmbeddedTweet{}.SoftIntervention,
	shared.EmbeddedTweetQuotedTweet{}.SoftIntervention,
	shared.EmbeddedTweetQuotedTweetQuotedTweet{}.SoftIntervention,
	shared.EmbeddedTweetQuotedTweetQuotedTweetQuotedTweet{}.SoftIntervention,
	shared.EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweet{}.SoftIntervention,
	shared.EmbeddedTweetQuotedTweetRetweetedTweet{}.SoftIntervention,
	shared.EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweet{}.SoftIntervention,
	shared.EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweet{}.SoftIntervention,
	shared.EmbeddedTweetRetweetedTweet{}.SoftIntervention,
	shared.EmbeddedTweetRetweetedTweetQuotedTweet{}.SoftIntervention,
	shared.EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweet{}.SoftIntervention,
	shared.EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweet{}.SoftIntervention,
	shared.EmbeddedTweetRetweetedTweetRetweetedTweet{}.SoftIntervention,
	shared.EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweet{}.SoftIntervention,
	shared.EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweet{}.SoftIntervention,
}

var _ shared.SoftInterventionText = shared.SoftIntervention{}.Text

var _ []shared.SearchTweetInterstitialTextEntity = shared.SoftInterventionText{}.Entities

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

func TestResponsesDecodeEmbeddedTweetFields(t *testing.T) {
	payload, err := fixtures.ReadFile("testdata/fixtures/EmbeddedTweet.json")
	if err != nil {
		t.Fatal(err)
	}
	tweetJSON := fmt.Sprintf(`{"id":"1234567891","text":"Repost","bookmarkCount":0,"likeCount":0,"quoteCount":0,"replyCount":0,"retweetCount":0,"viewCount":0,"parentTweet":%s,"quoted_tweet":%s,"retweeted_tweet":%s}`, payload, payload, payload)
	values := []fieldValue{
		field("ID", "2102761519985332442"),
		field("BookmarkCount", int64(599)),
		field("LikeCount", int64(9775)),
		field("QuoteCount", int64(85)),
		field("ReplyCount", int64(271)),
		field("RetweetCount", int64(1364)),
		field("ViewCount", int64(962488)),
		field("Article.Title", "Release notes"),
		field("Author.ID", "11348282"),
		field("Author.Username", "NASA"),
		field("Card.Name", "summary"),
		field("CommunityID", "community-1"),
		field("GrokShareAttachment", xtwitterscraper.PublicMetadata{"id": json.RawMessage(`"grok-1"`)}),
		field("JetfuelAttachment", xtwitterscraper.PublicMetadata{"id": json.RawMessage(`"jetfuel-1"`)}),
		field("LimitedActions[0].Action", "Reply"),
		field("Place.ID", "example"),
		field("PostCta", xtwitterscraper.PublicMetadata{"action": json.RawMessage(`"open"`)}),
		field("PreviousCounts.LikeCount", int64(1)),
		field("ReactionContext.ReactedToPostID", "1234567890"),
		field("ReactionContext.ReactedToUser.Username", "NASA"),
		field("RetweetedAt", "2026-09-11T20:24:45Z"),
		field("Scopes", xtwitterscraper.PublicMetadata{"followers": json.RawMessage("true")}),
		field("Type", "tweet"),
	}
	for _, field := range tweetFields {
		if !strings.HasPrefix(field.path, "ParentTweet.") {
			values = append(values, field)
		}
	}
	cases := []struct {
		name     string
		response interface{ UnmarshalJSON([]byte) error }
		envelope string
		path     string
	}{
		{"replies", &xtwitterscraper.XTweetGetRepliesResponse{}, `{"tweets":[%s]}`, "Tweets[0]"},
		{"nested replies", &xtwitterscraper.XTweetGetRepliesResponse{}, `{"nested_replies":[%s]}`, "NestedReplies[0]"},
		{"tweet", &xtwitterscraper.XTweetGetResponse{}, `{"tweet":%s}`, "Tweet"},
		{"users", &shared.PaginatedTweets{}, `{"tweets":[%s]}`, "Tweets[0]"},
		{"space", &xtwitterscraper.XSpaceGetResponse{}, `{"space":{"tweet":%s}}`, "Space.Tweet"},
		{"space sharing", &xtwitterscraper.XSpaceGetResponse{}, `{"space":{"sharings":[{"tweet":%s}]}}`, "Space.Sharings[0].Tweet"},
		{"space search", &xtwitterscraper.XSpaceSearchResponse{}, `{"spaces":[{"tweet":%s}]}`, "Spaces[0].Tweet"},
		{"space search sharing", &xtwitterscraper.XSpaceSearchResponse{}, `{"spaces":[{"sharings":[{"tweet":%s}]}]}`, "Spaces[0].Sharings[0].Tweet"},
		{"brand", &xtwitterscraper.XAnalysisBrandResponse{}, `{"results":[{"tweet":%s}]}`, "Results[0].Tweet"},
		{"classify", &xtwitterscraper.XAnalysisClassifyResponse{}, `{"results":[{"tweet":%s}]}`, "Results[0].Tweet"},
		{"market signals", &xtwitterscraper.XAnalysisMarketSignalsResponse{}, `{"results":[{"tweet":%s}]}`, "Results[0].Tweet"},
		{"news", &xtwitterscraper.XAnalysisNewsResponse{}, `{"results":[{"tweet":%s}]}`, "Results[0].Tweet"},
		{"sentiment", &xtwitterscraper.XAnalysisSentimentResponse{}, `{"results":[{"tweet":%s}]}`, "Results[0].Tweet"},
		{"viral score", &xtwitterscraper.XAnalysisViralScoreResponse{}, `{"results":[{"tweet":%s}]}`, "Results[0].Tweet"},
	}
	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			if err := test.response.UnmarshalJSON([]byte(fmt.Sprintf(test.envelope, tweetJSON))); err != nil {
				t.Fatal(err)
			}
			expectDecoded(t, test.response)
			for _, field := range []string{"ParentTweet", "QuotedTweet", "RetweetedTweet"} {
				embedded, ok := valueAt(reflect.ValueOf(test.response), test.path+"."+field)
				if !ok {
					t.Fatalf("missing %s.%s", test.path, field)
				}
				tweet, ok := reflect.TypeAssert[shared.EmbeddedTweet](embedded)
				if !ok {
					t.Fatalf("%s.%s is a %s", test.path, field, embedded.Type())
				}
				expectValues(t, tweet, values)
				if !tweet.JSON.IsPinned.Valid() || !tweet.JSON.Author.Valid() || !tweet.JSON.ParentTweet.Valid() || !tweet.JSON.QuotedTweet.Valid() || !tweet.JSON.RetweetedTweet.Valid() {
					t.Fatalf("missing field metadata in %s", tweet.RawJSON())
				}
				if tweet.ParentTweet == nil {
					t.Fatalf("missing parent post in %s", tweet.RawJSON())
				}
				for _, nested := range []struct {
					id          string
					text        string
					retweetedID string
					quotedText  string
				}{
					{tweet.ParentTweet.ID, tweet.ParentTweet.Text, tweet.ParentTweet.RetweetedTweet.ID, tweet.ParentTweet.RetweetedTweet.QuotedTweet.Text},
					{tweet.QuotedTweet.ID, tweet.QuotedTweet.Text, tweet.QuotedTweet.RetweetedTweet.ID, tweet.QuotedTweet.RetweetedTweet.QuotedTweet.Text},
					{tweet.RetweetedTweet.ID, tweet.RetweetedTweet.Text, tweet.RetweetedTweet.RetweetedTweet.ID, tweet.RetweetedTweet.RetweetedTweet.QuotedTweet.Text},
				} {
					if nested.id != "1234567890" || nested.text != "Just launched!" || nested.retweetedID != "nested-1" || nested.quotedText != "Deeply nested post" {
						t.Fatalf("decoded nested post = %#v", nested)
					}
				}
			}
		})
	}
}

func TestNestedEmbeddedTweetModelsDecodeEveryField(t *testing.T) {
	payload, err := fixtures.ReadFile("testdata/fixtures/EmbeddedTweet.json")
	if err != nil {
		t.Fatal(err)
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(payload, &fields); err != nil {
		t.Fatal(err)
	}
	delete(fields, "quoted_tweet")
	delete(fields, "retweeted_tweet")
	payload, err = json.Marshal(fields)
	if err != nil {
		t.Fatal(err)
	}
	values := []fieldValue{
		field("Author.Username", "NASA"),
		field("ParentTweet.Text", "Just launched!"),
		field("ParentTweet.RetweetedTweet.ID", "nested-1"),
	}
	for _, field := range tweetFields {
		if !strings.HasPrefix(field.path, "ParentTweet.") {
			values = append(values, field)
		}
	}
	models := []interface{ UnmarshalJSON([]byte) error }{
		&shared.EmbeddedTweetQuotedTweet{},
		&shared.EmbeddedTweetQuotedTweetQuotedTweet{},
		&shared.EmbeddedTweetQuotedTweetQuotedTweetQuotedTweet{},
		&shared.EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweet{},
		&shared.EmbeddedTweetQuotedTweetRetweetedTweet{},
		&shared.EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweet{},
		&shared.EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweet{},
		&shared.EmbeddedTweetRetweetedTweet{},
		&shared.EmbeddedTweetRetweetedTweetQuotedTweet{},
		&shared.EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweet{},
		&shared.EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweet{},
		&shared.EmbeddedTweetRetweetedTweetRetweetedTweet{},
		&shared.EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweet{},
		&shared.EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweet{},
	}
	for _, model := range models {
		t.Run(reflect.TypeOf(model).Elem().Name(), func(t *testing.T) {
			if err := model.UnmarshalJSON(payload); err != nil {
				t.Fatal(err)
			}
			expectDecoded(t, model)
			expectValues(t, model, values)
		})
	}
}

func TestEmbeddedTweetOptionalFields(t *testing.T) {
	for _, payload := range []string{`{}`, `{"id":"1234567890","text":"Just launched!"}`, `{"author":null,"parentTweet":null,"quoted_tweet":null,"retweeted_tweet":null}`} {
		t.Run(payload, func(t *testing.T) {
			var tweet shared.EmbeddedTweet
			if err := tweet.UnmarshalJSON([]byte(payload)); err != nil {
				t.Fatal(err)
			}
			expectDecoded(t, tweet)
			if tweet.JSON.IsPinned.Valid() || tweet.JSON.CashtagAttachments.Valid() || tweet.JSON.Author.Valid() || tweet.ParentTweet != nil || tweet.JSON.QuotedTweet.Valid() || tweet.JSON.RetweetedTweet.Valid() {
				t.Fatalf("decoded %s", tweet.RawJSON())
			}
		})
	}
}

func TestTweetModelsUnknownCounts(t *testing.T) {
	cases := []struct {
		name    string
		raw     string
		want    []shared.TweetUnknownCount
		present bool
	}{
		{"known counts", `["bookmarkCount","viewCount"]`, []shared.TweetUnknownCount{shared.TweetUnknownCountBookmarkCount, shared.TweetUnknownCountViewCount}, true},
		{"future count", `["futureCount"]`, []shared.TweetUnknownCount{"futureCount"}, true},
		{"empty", `[]`, []shared.TweetUnknownCount{}, true},
		{"absent", "", nil, false},
	}
	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			payload := `{"id":"1234567890","text":"Post","bookmarkCount":0,"likeCount":0,"quoteCount":0,"replyCount":0,"retweetCount":0,"viewCount":0}`
			if test.present {
				payload = strings.TrimSuffix(payload, "}") + `,"unknownCounts":` + test.raw + "}"
			}
			models := []interface{ UnmarshalJSON([]byte) error }{
				&xtwitterscraper.TweetDetail{},
				&shared.SearchTweet{},
				&shared.EmbeddedTweet{},
				&shared.EmbeddedTweetQuotedTweet{},
				&shared.EmbeddedTweetQuotedTweetQuotedTweet{},
				&shared.EmbeddedTweetQuotedTweetQuotedTweetQuotedTweet{},
				&shared.EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweet{},
				&shared.EmbeddedTweetQuotedTweetRetweetedTweet{},
				&shared.EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweet{},
				&shared.EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweet{},
				&shared.EmbeddedTweetRetweetedTweet{},
				&shared.EmbeddedTweetRetweetedTweetQuotedTweet{},
				&shared.EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweet{},
				&shared.EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweet{},
				&shared.EmbeddedTweetRetweetedTweetRetweetedTweet{},
				&shared.EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweet{},
				&shared.EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweet{},
			}
			for _, model := range models {
				t.Run(reflect.TypeOf(model).Elem().Name(), func(t *testing.T) {
					if err := model.UnmarshalJSON([]byte(payload)); err != nil {
						t.Fatal(err)
					}
					expectDecoded(t, model)
					expectValues(t, model, []fieldValue{field("UnknownCounts", test.want)})
					metadata, ok := valueAt(reflect.ValueOf(model), "JSON.UnknownCounts")
					if !ok {
						t.Fatal("missing unknown counts metadata")
					}
					field, ok := reflect.TypeAssert[respjson.Field](metadata)
					if !ok {
						t.Fatalf("unknown counts metadata is a %s", metadata.Type())
					}
					if field.Valid() != test.present || field.Raw() != test.raw {
						t.Fatalf("unknown counts metadata = %#v, want present %v and raw %q", field, test.present, test.raw)
					}
				})
			}
		})
	}
}

func TestTweetResponseUnknownCounts(t *testing.T) {
	payload := []byte(`{"tweet":{"id":"1234567891","text":"Repost","bookmarkCount":0,"unknownCounts":["bookmarkCount","viewCount"],"likeCount":0,"quoteCount":0,"replyCount":0,"retweetCount":0,"viewCount":0,"retweeted_tweet":{"id":"1234567890","text":"Post","bookmarkCount":0,"unknownCounts":["bookmarkCount","viewCount"],"likeCount":0,"quoteCount":0,"replyCount":0,"retweetCount":0,"viewCount":0}}}`)
	var response xtwitterscraper.XTweetGetResponse
	if err := response.UnmarshalJSON(payload); err != nil {
		t.Fatal(err)
	}
	expectDecoded(t, response)
	want := []shared.TweetUnknownCount{shared.TweetUnknownCountBookmarkCount, shared.TweetUnknownCountViewCount}
	expectValues(t, response, []fieldValue{
		field("Tweet.UnknownCounts", want),
		field("Tweet.RetweetedTweet.UnknownCounts", want),
	})
	if !response.Tweet.JSON.UnknownCounts.Valid() || !response.Tweet.RetweetedTweet.JSON.UnknownCounts.Valid() {
		t.Fatalf("missing unknown counts metadata in %s", response.RawJSON())
	}
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
	expectValues(t, res, []fieldValue{field("Author.Followers", int64(92376233)), field("Author.Verified", true)})
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
