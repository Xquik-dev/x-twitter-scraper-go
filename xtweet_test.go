// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

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

func TestXTweetNewWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
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
	_, err := client.X.Tweets.New(context.TODO(), xtwitterscraper.XTweetNewParams{
		Account:        "@elonmusk",
		AttachmentURL:  xtwitterscraper.String("https://x.com/elonmusk/status/1234567890"),
		CommunityID:    xtwitterscraper.String("1500000000000000000"),
		IsNoteTweet:    xtwitterscraper.Bool(false),
		Media:          []string{"https://example.com/video.mp4"},
		ReplyToTweetID: xtwitterscraper.String("1234567890"),
		Text:           xtwitterscraper.String("Just launched our new feature!"),
	})
	if err != nil {
		var apierr *xtwitterscraper.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestXTweetGet(t *testing.T) {
	t.Skip("Mock server tests are disabled")
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
	_, err := client.X.Tweets.Get(context.TODO(), "id")
	if err != nil {
		var apierr *xtwitterscraper.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestXTweetList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
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
	_, err := client.X.Tweets.List(context.TODO(), xtwitterscraper.XTweetListParams{
		IDs: "ids",
	})
	if err != nil {
		var apierr *xtwitterscraper.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestXTweetDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
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
	_, err := client.X.Tweets.Delete(
		context.TODO(),
		"id",
		xtwitterscraper.XTweetDeleteParams{
			Account: "@elonmusk",
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

func TestXTweetGetFavoritersWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
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
	_, err := client.X.Tweets.GetFavoriters(
		context.TODO(),
		"id",
		xtwitterscraper.XTweetGetFavoritersParams{
			Cursor:   xtwitterscraper.String("cursor"),
			PageSize: xtwitterscraper.Int(20),
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

func TestXTweetGetQuotesWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
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
	_, err := client.X.Tweets.GetQuotes(
		context.TODO(),
		"id",
		xtwitterscraper.XTweetGetQuotesParams{
			AnyWords:          xtwitterscraper.String("anyWords"),
			Cashtags:          xtwitterscraper.String("cashtags"),
			ConversationID:    xtwitterscraper.String("conversationId"),
			Cursor:            xtwitterscraper.String("cursor"),
			ExactPhrase:       xtwitterscraper.String("exactPhrase"),
			ExcludeWords:      xtwitterscraper.String("excludeWords"),
			FromUser:          xtwitterscraper.String("fromUser"),
			Hashtags:          xtwitterscraper.String("hashtags"),
			IncludeReplies:    xtwitterscraper.Bool(true),
			InReplyToTweetID:  xtwitterscraper.String("inReplyToTweetId"),
			Language:          xtwitterscraper.String("language"),
			MediaType:         xtwitterscraper.XTweetGetQuotesParamsMediaTypeImages,
			Mentioning:        xtwitterscraper.String("mentioning"),
			MinFaves:          xtwitterscraper.Int(0),
			MinQuotes:         xtwitterscraper.Int(0),
			MinReplies:        xtwitterscraper.Int(0),
			MinRetweets:       xtwitterscraper.Int(0),
			PageSize:          xtwitterscraper.Int(1),
			Quotes:            xtwitterscraper.XTweetGetQuotesParamsQuotesInclude,
			QuotesOfTweetID:   xtwitterscraper.String("quotesOfTweetId"),
			Replies:           xtwitterscraper.XTweetGetQuotesParamsRepliesInclude,
			Retweets:          xtwitterscraper.XTweetGetQuotesParamsRetweetsInclude,
			RetweetsOfTweetID: xtwitterscraper.String("retweetsOfTweetId"),
			SinceDate:         xtwitterscraper.Time(time.Now()),
			SinceTime:         xtwitterscraper.String("sinceTime"),
			ToUser:            xtwitterscraper.String("toUser"),
			UntilDate:         xtwitterscraper.Time(time.Now()),
			UntilTime:         xtwitterscraper.String("untilTime"),
			URL:               xtwitterscraper.String("url"),
			VerifiedOnly:      xtwitterscraper.Bool(true),
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

func TestXTweetGetRepliesWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
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
	_, err := client.X.Tweets.GetReplies(
		context.TODO(),
		"id",
		xtwitterscraper.XTweetGetRepliesParams{
			AnyWords:          xtwitterscraper.String("anyWords"),
			Cashtags:          xtwitterscraper.String("cashtags"),
			ConversationID:    xtwitterscraper.String("conversationId"),
			Cursor:            xtwitterscraper.String("cursor"),
			ExactPhrase:       xtwitterscraper.String("exactPhrase"),
			ExcludeWords:      xtwitterscraper.String("excludeWords"),
			FromUser:          xtwitterscraper.String("fromUser"),
			Hashtags:          xtwitterscraper.String("hashtags"),
			InReplyToTweetID:  xtwitterscraper.String("inReplyToTweetId"),
			Language:          xtwitterscraper.String("language"),
			MediaType:         xtwitterscraper.XTweetGetRepliesParamsMediaTypeImages,
			Mentioning:        xtwitterscraper.String("mentioning"),
			MinFaves:          xtwitterscraper.Int(0),
			MinQuotes:         xtwitterscraper.Int(0),
			MinReplies:        xtwitterscraper.Int(0),
			MinRetweets:       xtwitterscraper.Int(0),
			PageSize:          xtwitterscraper.Int(1),
			Quotes:            xtwitterscraper.XTweetGetRepliesParamsQuotesInclude,
			QuotesOfTweetID:   xtwitterscraper.String("quotesOfTweetId"),
			Replies:           xtwitterscraper.XTweetGetRepliesParamsRepliesInclude,
			Retweets:          xtwitterscraper.XTweetGetRepliesParamsRetweetsInclude,
			RetweetsOfTweetID: xtwitterscraper.String("retweetsOfTweetId"),
			SinceDate:         xtwitterscraper.Time(time.Now()),
			SinceTime:         xtwitterscraper.String("sinceTime"),
			ToUser:            xtwitterscraper.String("toUser"),
			UntilDate:         xtwitterscraper.Time(time.Now()),
			UntilTime:         xtwitterscraper.String("untilTime"),
			URL:               xtwitterscraper.String("url"),
			VerifiedOnly:      xtwitterscraper.Bool(true),
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

func TestXTweetGetRetweetersWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
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
	_, err := client.X.Tweets.GetRetweeters(
		context.TODO(),
		"id",
		xtwitterscraper.XTweetGetRetweetersParams{
			Cursor:   xtwitterscraper.String("cursor"),
			PageSize: xtwitterscraper.Int(20),
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

func TestXTweetGetThreadWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
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
	_, err := client.X.Tweets.GetThread(
		context.TODO(),
		"id",
		xtwitterscraper.XTweetGetThreadParams{
			Cursor:   xtwitterscraper.String("cursor"),
			PageSize: xtwitterscraper.Int(1),
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

func TestXTweetSearchWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
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
	_, err := client.X.Tweets.Search(context.TODO(), xtwitterscraper.XTweetSearchParams{
		Q:                 "q",
		AdvancedQuery:     xtwitterscraper.String("advancedQuery"),
		AnyWords:          xtwitterscraper.String("anyWords"),
		BoundingBox:       xtwitterscraper.String("boundingBox"),
		Cashtags:          xtwitterscraper.String("cashtags"),
		ConversationID:    xtwitterscraper.String("conversationId"),
		Cursor:            xtwitterscraper.String("cursor"),
		ExactPhrase:       xtwitterscraper.String("exactPhrase"),
		ExcludeWords:      xtwitterscraper.String("excludeWords"),
		FromUser:          xtwitterscraper.String("fromUser"),
		Hashtags:          xtwitterscraper.String("hashtags"),
		InReplyToTweetID:  xtwitterscraper.String("inReplyToTweetId"),
		Language:          xtwitterscraper.String("language"),
		Limit:             xtwitterscraper.Int(200),
		ListID:            xtwitterscraper.String("listId"),
		MediaType:         xtwitterscraper.XTweetSearchParamsMediaTypeImages,
		Mentioning:        xtwitterscraper.String("mentioning"),
		MinFaves:          xtwitterscraper.Int(0),
		MinQuotes:         xtwitterscraper.Int(0),
		MinReplies:        xtwitterscraper.Int(0),
		MinRetweets:       xtwitterscraper.Int(0),
		Place:             xtwitterscraper.String("place"),
		PlaceCountry:      xtwitterscraper.String("placeCountry"),
		PointRadius:       xtwitterscraper.String("pointRadius"),
		QueryType:         xtwitterscraper.XTweetSearchParamsQueryTypeLatest,
		Quotes:            xtwitterscraper.XTweetSearchParamsQuotesInclude,
		QuotesOfTweetID:   xtwitterscraper.String("quotesOfTweetId"),
		Replies:           xtwitterscraper.XTweetSearchParamsRepliesInclude,
		Retweets:          xtwitterscraper.XTweetSearchParamsRetweetsInclude,
		RetweetsOfTweetID: xtwitterscraper.String("retweetsOfTweetId"),
		SinceDate:         xtwitterscraper.Time(time.Now()),
		SinceTime:         xtwitterscraper.String("sinceTime"),
		ToUser:            xtwitterscraper.String("toUser"),
		UntilDate:         xtwitterscraper.Time(time.Now()),
		UntilTime:         xtwitterscraper.String("untilTime"),
		URL:               xtwitterscraper.String("url"),
		VerifiedOnly:      xtwitterscraper.Bool(true),
	})
	if err != nil {
		var apierr *xtwitterscraper.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
