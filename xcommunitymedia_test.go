// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

package xtwitterscraper_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/Xquik-dev/x-twitter-scraper-go"
	"github.com/Xquik-dev/x-twitter-scraper-go/internal/testutil"
	"github.com/Xquik-dev/x-twitter-scraper-go/option"
)

func TestXCommunityMediaListWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := xtwitterscraper.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
		option.WithBearerToken("My Bearer Token"),
	)
	_, err := client.X.Communities.Media.List(
		context.TODO(),
		"id",
		xtwitterscraper.XCommunityMediaListParams{
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
			MaxFaves:          xtwitterscraper.Int(0),
			MaxID:             xtwitterscraper.String("maxId"),
			MaxQuotes:         xtwitterscraper.Int(0),
			MaxReplies:        xtwitterscraper.Int(0),
			MaxRetweets:       xtwitterscraper.Int(0),
			MediaType:         xtwitterscraper.XCommunityMediaListParamsMediaTypeImages,
			Mentioning:        xtwitterscraper.String("mentioning"),
			MinBookmarks:      xtwitterscraper.Int(0),
			MinLikes:          xtwitterscraper.Int(0),
			MinQuotes:         xtwitterscraper.Int(0),
			MinReplies:        xtwitterscraper.Int(0),
			MinRetweets:       xtwitterscraper.Int(0),
			MinViews:          xtwitterscraper.Int(0),
			NativeRetweets:    xtwitterscraper.Bool(true),
			Omit:              xtwitterscraper.String("entities,media"),
			PageSize:          xtwitterscraper.Int(1),
			Quotes:            xtwitterscraper.XCommunityMediaListParamsQuotesInclude,
			QuotesOfTweetID:   xtwitterscraper.String("quotesOfTweetId"),
			Replies:           xtwitterscraper.XCommunityMediaListParamsRepliesInclude,
			Retweets:          xtwitterscraper.XCommunityMediaListParamsRetweetsInclude,
			RetweetsOfTweetID: xtwitterscraper.String("retweetsOfTweetId"),
			SinceDate:         xtwitterscraper.Time(time.Now()),
			SinceID:           xtwitterscraper.String("sinceId"),
			ToUser:            xtwitterscraper.String("toUser"),
			UntilDate:         xtwitterscraper.Time(time.Now()),
			URL:               xtwitterscraper.String("url"),
			VerifiedOnly:      xtwitterscraper.Bool(true),
			WithinTime:        xtwitterscraper.String("7d"),
		},
	)
	if err != nil {
		var apierr *xtwitterscraper.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
