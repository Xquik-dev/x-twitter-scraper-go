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
		IdempotencyKey: "Idempotency-Key",
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
			Account:        "@elonmusk",
			IdempotencyKey: "Idempotency-Key",
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
			BioContains:       xtwitterscraper.String("bioContains"),
			Cursor:            xtwitterscraper.String("cursor"),
			HasLocation:       xtwitterscraper.Bool(true),
			HasWebsite:        xtwitterscraper.Bool(true),
			LocationContains:  xtwitterscraper.String("locationContains"),
			MaxFollowers:      xtwitterscraper.Int(0),
			MaxFollowing:      xtwitterscraper.Int(0),
			MaxStatuses:       xtwitterscraper.Int(0),
			MinAccountAgeDays: xtwitterscraper.Int(0),
			MinFollowers:      xtwitterscraper.Int(0),
			MinFollowing:      xtwitterscraper.Int(0),
			MinStatuses:       xtwitterscraper.Int(0),
			PageSize:          xtwitterscraper.Int(20),
			UsernameContains:  xtwitterscraper.String("usernameContains"),
			VerifiedOnly:      xtwitterscraper.Bool(true),
			VerifiedType:      xtwitterscraper.String("verifiedType"),
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
			BlueVerifiedOnly:  xtwitterscraper.Bool(true),
			CardName:          xtwitterscraper.String("cardName"),
			Cashtags:          xtwitterscraper.String("cashtags"),
			ConversationID:    xtwitterscraper.String("conversationId"),
			Cursor:            xtwitterscraper.String("cursor"),
			ExactPhrase:       xtwitterscraper.String("exactPhrase"),
			ExcludeSource:     xtwitterscraper.String("excludeSource"),
			ExcludeWords:      xtwitterscraper.String("excludeWords"),
			FromUser:          xtwitterscraper.String("fromUser"),
			Geocode:           xtwitterscraper.String("geocode"),
			Hashtags:          xtwitterscraper.String("hashtags"),
			IncludeReplies:    xtwitterscraper.Bool(true),
			InReplyToTweetID:  xtwitterscraper.String("inReplyToTweetId"),
			Language:          xtwitterscraper.String("language"),
			MaxFaves:          xtwitterscraper.Int(0),
			MaxID:             xtwitterscraper.String("maxId"),
			MaxQuotes:         xtwitterscraper.Int(0),
			MaxReplies:        xtwitterscraper.Int(0),
			MaxRetweets:       xtwitterscraper.Int(0),
			MediaType:         xtwitterscraper.XTweetGetQuotesParamsMediaTypeImages,
			Mentioning:        xtwitterscraper.String("mentioning"),
			MinBookmarks:      xtwitterscraper.Int(0),
			MinFaves:          xtwitterscraper.Int(0),
			MinQuotes:         xtwitterscraper.Int(0),
			MinReplies:        xtwitterscraper.Int(0),
			MinRetweets:       xtwitterscraper.Int(0),
			MinViews:          xtwitterscraper.Int(0),
			NativeRetweets:    xtwitterscraper.Bool(true),
			Near:              xtwitterscraper.String("near"),
			News:              xtwitterscraper.Bool(true),
			PageSize:          xtwitterscraper.Int(1),
			Quotes:            xtwitterscraper.XTweetGetQuotesParamsQuotesInclude,
			QuotesOfTweetID:   xtwitterscraper.String("quotesOfTweetId"),
			Replies:           xtwitterscraper.XTweetGetQuotesParamsRepliesInclude,
			Retweets:          xtwitterscraper.XTweetGetQuotesParamsRetweetsInclude,
			RetweetsOfTweetID: xtwitterscraper.String("retweetsOfTweetId"),
			Safe:              xtwitterscraper.Bool(true),
			SinceDate:         xtwitterscraper.Time(time.Now()),
			SinceID:           xtwitterscraper.String("sinceId"),
			SinceTime:         xtwitterscraper.String("sinceTime"),
			Source:            xtwitterscraper.String("source"),
			ToUser:            xtwitterscraper.String("toUser"),
			UntilDate:         xtwitterscraper.Time(time.Now()),
			UntilTime:         xtwitterscraper.String("untilTime"),
			URL:               xtwitterscraper.String("url"),
			VerifiedOnly:      xtwitterscraper.Bool(true),
			Within:            xtwitterscraper.String("within"),
			WithinTime:        xtwitterscraper.String("withinTime"),
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
			AnyWords:              xtwitterscraper.String("anyWords"),
			BlueVerifiedOnly:      xtwitterscraper.Bool(true),
			CardName:              xtwitterscraper.String("cardName"),
			Cashtags:              xtwitterscraper.String("cashtags"),
			ConversationID:        xtwitterscraper.String("conversationId"),
			Cursor:                xtwitterscraper.String("cursor"),
			ExactPhrase:           xtwitterscraper.String("exactPhrase"),
			ExcludeOriginalAuthor: xtwitterscraper.Bool(true),
			ExcludeSource:         xtwitterscraper.String("excludeSource"),
			ExcludeWords:          xtwitterscraper.String("excludeWords"),
			FromUser:              xtwitterscraper.String("fromUser"),
			Geocode:               xtwitterscraper.String("geocode"),
			Hashtags:              xtwitterscraper.String("hashtags"),
			HasMediaOnly:          xtwitterscraper.Bool(true),
			IncludeOriginalPost:   xtwitterscraper.Bool(true),
			InReplyToTweetID:      xtwitterscraper.String("inReplyToTweetId"),
			Language:              xtwitterscraper.String("language"),
			Limit:                 xtwitterscraper.Int(1),
			MaxDepth:              xtwitterscraper.Int(1),
			MaxFaves:              xtwitterscraper.Int(0),
			MaxID:                 xtwitterscraper.String("maxId"),
			MaxQuotes:             xtwitterscraper.Int(0),
			MaxReplies:            xtwitterscraper.Int(0),
			MaxRetweets:           xtwitterscraper.Int(0),
			MediaType:             xtwitterscraper.XTweetGetRepliesParamsMediaTypeImages,
			Mentioning:            xtwitterscraper.String("mentioning"),
			MinBookmarks:          xtwitterscraper.Int(0),
			MinFaves:              xtwitterscraper.Int(0),
			MinQuotes:             xtwitterscraper.Int(0),
			MinReplies:            xtwitterscraper.Int(0),
			MinRetweets:           xtwitterscraper.Int(0),
			MinViews:              xtwitterscraper.Int(0),
			Mode:                  xtwitterscraper.XTweetGetRepliesParamsModeStandard,
			NativeRetweets:        xtwitterscraper.Bool(true),
			Near:                  xtwitterscraper.String("near"),
			News:                  xtwitterscraper.Bool(true),
			PageSize:              xtwitterscraper.Int(1),
			Quotes:                xtwitterscraper.XTweetGetRepliesParamsQuotesInclude,
			QuotesOfTweetID:       xtwitterscraper.String("quotesOfTweetId"),
			Replies:               xtwitterscraper.XTweetGetRepliesParamsRepliesInclude,
			Retweets:              xtwitterscraper.XTweetGetRepliesParamsRetweetsInclude,
			RetweetsOfTweetID:     xtwitterscraper.String("retweetsOfTweetId"),
			Safe:                  xtwitterscraper.Bool(true),
			Scope:                 xtwitterscraper.XTweetGetRepliesParamsScopeAll,
			SinceDate:             xtwitterscraper.Time(time.Now()),
			SinceID:               xtwitterscraper.String("sinceId"),
			SinceTime:             xtwitterscraper.String("sinceTime"),
			Sort:                  xtwitterscraper.XTweetGetRepliesParamsSortRelevance,
			Source:                xtwitterscraper.String("source"),
			ToUser:                xtwitterscraper.String("toUser"),
			UntilDate:             xtwitterscraper.Time(time.Now()),
			UntilTime:             xtwitterscraper.String("untilTime"),
			URL:                   xtwitterscraper.String("url"),
			VerifiedOnly:          xtwitterscraper.Bool(true),
			Within:                xtwitterscraper.String("within"),
			WithinTime:            xtwitterscraper.String("withinTime"),
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
			BioContains:       xtwitterscraper.String("bioContains"),
			Cursor:            xtwitterscraper.String("cursor"),
			HasLocation:       xtwitterscraper.Bool(true),
			HasWebsite:        xtwitterscraper.Bool(true),
			LocationContains:  xtwitterscraper.String("locationContains"),
			MaxFollowers:      xtwitterscraper.Int(0),
			MaxFollowing:      xtwitterscraper.Int(0),
			MaxStatuses:       xtwitterscraper.Int(0),
			MinAccountAgeDays: xtwitterscraper.Int(0),
			MinFollowers:      xtwitterscraper.Int(0),
			MinFollowing:      xtwitterscraper.Int(0),
			MinStatuses:       xtwitterscraper.Int(0),
			PageSize:          xtwitterscraper.Int(20),
			UsernameContains:  xtwitterscraper.String("usernameContains"),
			VerifiedOnly:      xtwitterscraper.Bool(true),
			VerifiedType:      xtwitterscraper.String("verifiedType"),
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
			AnyWords:          xtwitterscraper.String("anyWords"),
			BlueVerifiedOnly:  xtwitterscraper.Bool(true),
			Cashtags:          xtwitterscraper.String("cashtags"),
			ConversationID:    xtwitterscraper.String("conversationId"),
			Cursor:            xtwitterscraper.String("cursor"),
			ExactPhrase:       xtwitterscraper.String("exactPhrase"),
			ExcludeWords:      xtwitterscraper.String("excludeWords"),
			FromUser:          xtwitterscraper.String("fromUser"),
			Hashtags:          xtwitterscraper.String("hashtags"),
			InReplyToTweetID:  xtwitterscraper.String("inReplyToTweetId"),
			Language:          xtwitterscraper.String("language"),
			MaxFaves:          xtwitterscraper.Int(0),
			MaxQuotes:         xtwitterscraper.Int(0),
			MaxReplies:        xtwitterscraper.Int(0),
			MaxRetweets:       xtwitterscraper.Int(0),
			MediaType:         xtwitterscraper.XTweetGetThreadParamsMediaTypeImages,
			Mentioning:        xtwitterscraper.String("mentioning"),
			MinBookmarks:      xtwitterscraper.Int(0),
			MinFaves:          xtwitterscraper.Int(0),
			MinQuotes:         xtwitterscraper.Int(0),
			MinReplies:        xtwitterscraper.Int(0),
			MinRetweets:       xtwitterscraper.Int(0),
			MinViews:          xtwitterscraper.Int(0),
			PageSize:          xtwitterscraper.Int(1),
			Quotes:            xtwitterscraper.XTweetGetThreadParamsQuotesInclude,
			QuotesOfTweetID:   xtwitterscraper.String("quotesOfTweetId"),
			Replies:           xtwitterscraper.XTweetGetThreadParamsRepliesInclude,
			Retweets:          xtwitterscraper.XTweetGetThreadParamsRetweetsInclude,
			RetweetsOfTweetID: xtwitterscraper.String("retweetsOfTweetId"),
			SinceDate:         xtwitterscraper.Time(time.Now()),
			ToUser:            xtwitterscraper.String("toUser"),
			UntilDate:         xtwitterscraper.Time(time.Now()),
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
		BlueVerifiedOnly:  xtwitterscraper.Bool(true),
		BoundingBox:       xtwitterscraper.String("boundingBox"),
		CardName:          xtwitterscraper.String("cardName"),
		Cashtags:          xtwitterscraper.String("cashtags"),
		ConversationID:    xtwitterscraper.String("conversationId"),
		Cursor:            xtwitterscraper.String("cursor"),
		ExactPhrase:       xtwitterscraper.String("exactPhrase"),
		ExcludeSource:     xtwitterscraper.String("excludeSource"),
		ExcludeWords:      xtwitterscraper.String("excludeWords"),
		FromUser:          xtwitterscraper.String("fromUser"),
		Geocode:           xtwitterscraper.String("geocode"),
		Hashtags:          xtwitterscraper.String("hashtags"),
		InReplyToTweetID:  xtwitterscraper.String("inReplyToTweetId"),
		Language:          xtwitterscraper.String("language"),
		Limit:             xtwitterscraper.Int(1),
		ListID:            xtwitterscraper.String("listId"),
		MaxFaves:          xtwitterscraper.Int(0),
		MaxID:             xtwitterscraper.String("maxId"),
		MaxQuotes:         xtwitterscraper.Int(0),
		MaxReplies:        xtwitterscraper.Int(0),
		MaxRetweets:       xtwitterscraper.Int(0),
		MediaType:         xtwitterscraper.XTweetSearchParamsMediaTypeImages,
		Mentioning:        xtwitterscraper.String("mentioning"),
		MinBookmarks:      xtwitterscraper.Int(0),
		MinFaves:          xtwitterscraper.Int(0),
		MinQuotes:         xtwitterscraper.Int(0),
		MinReplies:        xtwitterscraper.Int(0),
		MinRetweets:       xtwitterscraper.Int(0),
		MinViews:          xtwitterscraper.Int(0),
		Mode:              xtwitterscraper.XTweetSearchParamsModeStandard,
		NativeRetweets:    xtwitterscraper.Bool(true),
		Near:              xtwitterscraper.String("near"),
		News:              xtwitterscraper.Bool(true),
		Place:             xtwitterscraper.String("place"),
		PlaceCountry:      xtwitterscraper.String("placeCountry"),
		PointRadius:       xtwitterscraper.String("pointRadius"),
		QueryType:         xtwitterscraper.XTweetSearchParamsQueryTypeLatest,
		Quotes:            xtwitterscraper.XTweetSearchParamsQuotesInclude,
		QuotesOfTweetID:   xtwitterscraper.String("quotesOfTweetId"),
		Replies:           xtwitterscraper.XTweetSearchParamsRepliesInclude,
		Retweets:          xtwitterscraper.XTweetSearchParamsRetweetsInclude,
		RetweetsOfTweetID: xtwitterscraper.String("retweetsOfTweetId"),
		Safe:              xtwitterscraper.Bool(true),
		SinceDate:         xtwitterscraper.Time(time.Now()),
		SinceID:           xtwitterscraper.String("sinceId"),
		SinceTime:         xtwitterscraper.String("sinceTime"),
		Source:            xtwitterscraper.String("source"),
		ToUser:            xtwitterscraper.String("toUser"),
		UntilDate:         xtwitterscraper.Time(time.Now()),
		UntilTime:         xtwitterscraper.String("untilTime"),
		URL:               xtwitterscraper.String("url"),
		VerifiedOnly:      xtwitterscraper.Bool(true),
		Within:            xtwitterscraper.String("within"),
		WithinTime:        xtwitterscraper.String("withinTime"),
	})
	if err != nil {
		var apierr *xtwitterscraper.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
