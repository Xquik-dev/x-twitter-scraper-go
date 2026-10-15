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

func TestXUserGet(t *testing.T) {
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
	_, err := client.X.Users.Get(context.TODO(), "id")
	if err != nil {
		var apierr *xtwitterscraper.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestXUserRemoveFollower(t *testing.T) {
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
	_, err := client.X.Users.RemoveFollower(
		context.TODO(),
		"id",
		xtwitterscraper.XUserRemoveFollowerParams{
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

func TestXUserGetBatchWithOptionalParams(t *testing.T) {
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
	_, err := client.X.Users.GetBatch(context.TODO(), xtwitterscraper.XUserGetBatchParams{
		IDs:               "ids",
		MaxFollowers:      xtwitterscraper.Int(0),
		MinAccountAgeDays: xtwitterscraper.Int(0),
		MinFollowers:      xtwitterscraper.Int(0),
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

func TestXUserGetFollowersWithOptionalParams(t *testing.T) {
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
	_, err := client.X.Users.GetFollowers(
		context.TODO(),
		"id",
		xtwitterscraper.XUserGetFollowersParams{
			After:             xtwitterscraper.String("after"),
			BioContains:       xtwitterscraper.String("bioContains"),
			Cursor:            xtwitterscraper.String("cursor"),
			HasLocation:       xtwitterscraper.Bool(true),
			HasWebsite:        xtwitterscraper.Bool(true),
			Limit:             xtwitterscraper.Int(1),
			LocationContains:  xtwitterscraper.String("locationContains"),
			MaxFollowers:      xtwitterscraper.Int(0),
			MaxFollowing:      xtwitterscraper.Int(0),
			MaxStatuses:       xtwitterscraper.Int(0),
			MinAccountAgeDays: xtwitterscraper.Int(0),
			MinFollowers:      xtwitterscraper.Int(0),
			MinFollowing:      xtwitterscraper.Int(0),
			MinStatuses:       xtwitterscraper.Int(0),
			Mode:              xtwitterscraper.XUserGetFollowersParamsModeStandard,
			PageSize:          xtwitterscraper.Int(1),
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

func TestXUserGetFollowersYouKnowWithOptionalParams(t *testing.T) {
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
	_, err := client.X.Users.GetFollowersYouKnow(
		context.TODO(),
		"id",
		xtwitterscraper.XUserGetFollowersYouKnowParams{
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
			PageSize:          xtwitterscraper.Int(1),
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

func TestXUserGetFollowingWithOptionalParams(t *testing.T) {
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
	_, err := client.X.Users.GetFollowing(
		context.TODO(),
		"id",
		xtwitterscraper.XUserGetFollowingParams{
			After:             xtwitterscraper.String("after"),
			BioContains:       xtwitterscraper.String("bioContains"),
			Cursor:            xtwitterscraper.String("cursor"),
			HasLocation:       xtwitterscraper.Bool(true),
			HasWebsite:        xtwitterscraper.Bool(true),
			Limit:             xtwitterscraper.Int(1),
			LocationContains:  xtwitterscraper.String("locationContains"),
			MaxFollowers:      xtwitterscraper.Int(0),
			MaxFollowing:      xtwitterscraper.Int(0),
			MaxStatuses:       xtwitterscraper.Int(0),
			MinAccountAgeDays: xtwitterscraper.Int(0),
			MinFollowers:      xtwitterscraper.Int(0),
			MinFollowing:      xtwitterscraper.Int(0),
			MinStatuses:       xtwitterscraper.Int(0),
			Mode:              xtwitterscraper.XUserGetFollowingParamsModeStandard,
			PageSize:          xtwitterscraper.Int(1),
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

func TestXUserGetLikesWithOptionalParams(t *testing.T) {
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
	_, err := client.X.Users.GetLikes(
		context.TODO(),
		"id",
		xtwitterscraper.XUserGetLikesParams{
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
			InReplyToTweetID:  xtwitterscraper.String("inReplyToTweetId"),
			Language:          xtwitterscraper.String("language"),
			MaxFaves:          xtwitterscraper.Int(0),
			MaxID:             xtwitterscraper.String("maxId"),
			MaxQuotes:         xtwitterscraper.Int(0),
			MaxReplies:        xtwitterscraper.Int(0),
			MaxRetweets:       xtwitterscraper.Int(0),
			MediaType:         xtwitterscraper.XUserGetLikesParamsMediaTypeImages,
			Mentioning:        xtwitterscraper.String("mentioning"),
			MinBookmarks:      xtwitterscraper.Int(0),
			MinLikes:          xtwitterscraper.Int(0),
			MinQuotes:         xtwitterscraper.Int(0),
			MinReplies:        xtwitterscraper.Int(0),
			MinRetweets:       xtwitterscraper.Int(0),
			MinViews:          xtwitterscraper.Int(0),
			NativeRetweets:    xtwitterscraper.Bool(true),
			Near:              xtwitterscraper.String("near"),
			News:              xtwitterscraper.Bool(true),
			PageSize:          xtwitterscraper.Int(1),
			Quotes:            xtwitterscraper.XUserGetLikesParamsQuotesInclude,
			QuotesOfTweetID:   xtwitterscraper.String("quotesOfTweetId"),
			Replies:           xtwitterscraper.XUserGetLikesParamsRepliesInclude,
			Retweets:          xtwitterscraper.XUserGetLikesParamsRetweetsInclude,
			RetweetsOfTweetID: xtwitterscraper.String("retweetsOfTweetId"),
			Safe:              xtwitterscraper.Bool(true),
			SinceDate:         xtwitterscraper.Time(time.Now()),
			SinceID:           xtwitterscraper.String("sinceId"),
			Source:            xtwitterscraper.String("source"),
			ToUser:            xtwitterscraper.String("toUser"),
			UntilDate:         xtwitterscraper.Time(time.Now()),
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

func TestXUserGetMediaWithOptionalParams(t *testing.T) {
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
	_, err := client.X.Users.GetMedia(
		context.TODO(),
		"id",
		xtwitterscraper.XUserGetMediaParams{
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
			InReplyToTweetID:  xtwitterscraper.String("inReplyToTweetId"),
			Language:          xtwitterscraper.String("language"),
			MaxFaves:          xtwitterscraper.Int(0),
			MaxID:             xtwitterscraper.String("maxId"),
			MaxQuotes:         xtwitterscraper.Int(0),
			MaxReplies:        xtwitterscraper.Int(0),
			MaxRetweets:       xtwitterscraper.Int(0),
			MediaType:         xtwitterscraper.XUserGetMediaParamsMediaTypeImages,
			Mentioning:        xtwitterscraper.String("mentioning"),
			MinBookmarks:      xtwitterscraper.Int(0),
			MinLikes:          xtwitterscraper.Int(0),
			MinQuotes:         xtwitterscraper.Int(0),
			MinReplies:        xtwitterscraper.Int(0),
			MinRetweets:       xtwitterscraper.Int(0),
			MinViews:          xtwitterscraper.Int(0),
			NativeRetweets:    xtwitterscraper.Bool(true),
			Near:              xtwitterscraper.String("near"),
			News:              xtwitterscraper.Bool(true),
			PageSize:          xtwitterscraper.Int(1),
			Quotes:            xtwitterscraper.XUserGetMediaParamsQuotesInclude,
			QuotesOfTweetID:   xtwitterscraper.String("quotesOfTweetId"),
			Replies:           xtwitterscraper.XUserGetMediaParamsRepliesInclude,
			Retweets:          xtwitterscraper.XUserGetMediaParamsRetweetsInclude,
			RetweetsOfTweetID: xtwitterscraper.String("retweetsOfTweetId"),
			Safe:              xtwitterscraper.Bool(true),
			SinceDate:         xtwitterscraper.Time(time.Now()),
			SinceID:           xtwitterscraper.String("sinceId"),
			Source:            xtwitterscraper.String("source"),
			ToUser:            xtwitterscraper.String("toUser"),
			UntilDate:         xtwitterscraper.Time(time.Now()),
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

func TestXUserGetMentionsWithOptionalParams(t *testing.T) {
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
	_, err := client.X.Users.GetMentions(
		context.TODO(),
		"id",
		xtwitterscraper.XUserGetMentionsParams{
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
			InReplyToTweetID:  xtwitterscraper.String("inReplyToTweetId"),
			Language:          xtwitterscraper.String("language"),
			MaxFaves:          xtwitterscraper.Int(0),
			MaxID:             xtwitterscraper.String("maxId"),
			MaxQuotes:         xtwitterscraper.Int(0),
			MaxReplies:        xtwitterscraper.Int(0),
			MaxRetweets:       xtwitterscraper.Int(0),
			MediaType:         xtwitterscraper.XUserGetMentionsParamsMediaTypeImages,
			Mentioning:        xtwitterscraper.String("mentioning"),
			MinBookmarks:      xtwitterscraper.Int(0),
			MinLikes:          xtwitterscraper.Int(0),
			MinQuotes:         xtwitterscraper.Int(0),
			MinReplies:        xtwitterscraper.Int(0),
			MinRetweets:       xtwitterscraper.Int(0),
			MinViews:          xtwitterscraper.Int(0),
			NativeRetweets:    xtwitterscraper.Bool(true),
			Near:              xtwitterscraper.String("near"),
			News:              xtwitterscraper.Bool(true),
			PageSize:          xtwitterscraper.Int(1),
			Quotes:            xtwitterscraper.XUserGetMentionsParamsQuotesInclude,
			QuotesOfTweetID:   xtwitterscraper.String("quotesOfTweetId"),
			Replies:           xtwitterscraper.XUserGetMentionsParamsRepliesInclude,
			Retweets:          xtwitterscraper.XUserGetMentionsParamsRetweetsInclude,
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

func TestXUserGetRepliesWithOptionalParams(t *testing.T) {
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
	_, err := client.X.Users.GetReplies(
		context.TODO(),
		"id",
		xtwitterscraper.XUserGetRepliesParams{
			AnyWords:           xtwitterscraper.String("anyWords"),
			BlueVerifiedOnly:   xtwitterscraper.Bool(true),
			CardName:           xtwitterscraper.String("cardName"),
			Cashtags:           xtwitterscraper.String("cashtags"),
			ConversationID:     xtwitterscraper.String("conversationId"),
			Cursor:             xtwitterscraper.String("cursor"),
			ExactPhrase:        xtwitterscraper.String("exactPhrase"),
			ExcludeSource:      xtwitterscraper.String("excludeSource"),
			ExcludeWords:       xtwitterscraper.String("excludeWords"),
			FromUser:           xtwitterscraper.String("fromUser"),
			Geocode:            xtwitterscraper.String("geocode"),
			Hashtags:           xtwitterscraper.String("hashtags"),
			IncludeParentTweet: xtwitterscraper.Bool(true),
			InReplyToTweetID:   xtwitterscraper.String("inReplyToTweetId"),
			Language:           xtwitterscraper.String("language"),
			MaxFaves:           xtwitterscraper.Int(0),
			MaxID:              xtwitterscraper.String("maxId"),
			MaxQuotes:          xtwitterscraper.Int(0),
			MaxReplies:         xtwitterscraper.Int(0),
			MaxRetweets:        xtwitterscraper.Int(0),
			MediaType:          xtwitterscraper.XUserGetRepliesParamsMediaTypeImages,
			Mentioning:         xtwitterscraper.String("mentioning"),
			MinBookmarks:       xtwitterscraper.Int(0),
			MinLikes:           xtwitterscraper.Int(0),
			MinQuotes:          xtwitterscraper.Int(0),
			MinReplies:         xtwitterscraper.Int(0),
			MinRetweets:        xtwitterscraper.Int(0),
			MinViews:           xtwitterscraper.Int(0),
			NativeRetweets:     xtwitterscraper.Bool(true),
			Near:               xtwitterscraper.String("near"),
			News:               xtwitterscraper.Bool(true),
			PageSize:           xtwitterscraper.Int(1),
			Quotes:             xtwitterscraper.XUserGetRepliesParamsQuotesInclude,
			QuotesOfTweetID:    xtwitterscraper.String("quotesOfTweetId"),
			Replies:            xtwitterscraper.XUserGetRepliesParamsRepliesInclude,
			Retweets:           xtwitterscraper.XUserGetRepliesParamsRetweetsInclude,
			RetweetsOfTweetID:  xtwitterscraper.String("retweetsOfTweetId"),
			Safe:               xtwitterscraper.Bool(true),
			SinceDate:          xtwitterscraper.Time(time.Now()),
			SinceID:            xtwitterscraper.String("sinceId"),
			Source:             xtwitterscraper.String("source"),
			ToUser:             xtwitterscraper.String("toUser"),
			UntilDate:          xtwitterscraper.Time(time.Now()),
			URL:                xtwitterscraper.String("url"),
			VerifiedOnly:       xtwitterscraper.Bool(true),
			Within:             xtwitterscraper.String("within"),
			WithinTime:         xtwitterscraper.String("withinTime"),
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

func TestXUserGetSearchWithOptionalParams(t *testing.T) {
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
	_, err := client.X.Users.GetSearch(context.TODO(), xtwitterscraper.XUserGetSearchParams{
		Q:                 "q",
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
		PageSize:          xtwitterscraper.Int(1),
		UsernameContains:  xtwitterscraper.String("usernameContains"),
		VerifiedOnly:      xtwitterscraper.Bool(true),
		VerifiedType:      xtwitterscraper.String("verifiedType"),
	})
	if err != nil {
		var apierr *xtwitterscraper.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestXUserGetTweetsWithOptionalParams(t *testing.T) {
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
	_, err := client.X.Users.GetTweets(
		context.TODO(),
		"id",
		xtwitterscraper.XUserGetTweetsParams{
			AnyWords:           xtwitterscraper.String("anyWords"),
			BlueVerifiedOnly:   xtwitterscraper.Bool(true),
			CardName:           xtwitterscraper.String("cardName"),
			Cashtags:           xtwitterscraper.String("cashtags"),
			ConversationID:     xtwitterscraper.String("conversationId"),
			Cursor:             xtwitterscraper.String("cursor"),
			ExactPhrase:        xtwitterscraper.String("exactPhrase"),
			ExcludeSource:      xtwitterscraper.String("excludeSource"),
			ExcludeWords:       xtwitterscraper.String("excludeWords"),
			FromUser:           xtwitterscraper.String("fromUser"),
			Geocode:            xtwitterscraper.String("geocode"),
			Hashtags:           xtwitterscraper.String("hashtags"),
			IncludeParentTweet: xtwitterscraper.Bool(true),
			IncludeReplies:     xtwitterscraper.Bool(true),
			InReplyToTweetID:   xtwitterscraper.String("inReplyToTweetId"),
			Language:           xtwitterscraper.String("language"),
			MaxFaves:           xtwitterscraper.Int(0),
			MaxID:              xtwitterscraper.String("maxId"),
			MaxQuotes:          xtwitterscraper.Int(0),
			MaxReplies:         xtwitterscraper.Int(0),
			MaxRetweets:        xtwitterscraper.Int(0),
			MediaType:          xtwitterscraper.XUserGetTweetsParamsMediaTypeImages,
			Mentioning:         xtwitterscraper.String("mentioning"),
			MinBookmarks:       xtwitterscraper.Int(0),
			MinLikes:           xtwitterscraper.Int(0),
			MinQuotes:          xtwitterscraper.Int(0),
			MinReplies:         xtwitterscraper.Int(0),
			MinRetweets:        xtwitterscraper.Int(0),
			MinViews:           xtwitterscraper.Int(0),
			NativeRetweets:     xtwitterscraper.Bool(true),
			Near:               xtwitterscraper.String("near"),
			News:               xtwitterscraper.Bool(true),
			PageSize:           xtwitterscraper.Int(1),
			Quotes:             xtwitterscraper.XUserGetTweetsParamsQuotesInclude,
			QuotesOfTweetID:    xtwitterscraper.String("quotesOfTweetId"),
			Replies:            xtwitterscraper.XUserGetTweetsParamsRepliesInclude,
			Retweets:           xtwitterscraper.XUserGetTweetsParamsRetweetsInclude,
			RetweetsOfTweetID:  xtwitterscraper.String("retweetsOfTweetId"),
			Safe:               xtwitterscraper.Bool(true),
			SinceDate:          xtwitterscraper.Time(time.Now()),
			SinceID:            xtwitterscraper.String("sinceId"),
			Source:             xtwitterscraper.String("source"),
			ToUser:             xtwitterscraper.String("toUser"),
			UntilDate:          xtwitterscraper.Time(time.Now()),
			URL:                xtwitterscraper.String("url"),
			VerifiedOnly:       xtwitterscraper.Bool(true),
			Within:             xtwitterscraper.String("within"),
			WithinTime:         xtwitterscraper.String("withinTime"),
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

func TestXUserGetVerifiedFollowersWithOptionalParams(t *testing.T) {
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
	_, err := client.X.Users.GetVerifiedFollowers(
		context.TODO(),
		"id",
		xtwitterscraper.XUserGetVerifiedFollowersParams{
			After:             xtwitterscraper.String("after"),
			BioContains:       xtwitterscraper.String("bioContains"),
			Cursor:            xtwitterscraper.String("cursor"),
			HasLocation:       xtwitterscraper.Bool(true),
			HasWebsite:        xtwitterscraper.Bool(true),
			Limit:             xtwitterscraper.Int(1),
			LocationContains:  xtwitterscraper.String("locationContains"),
			MaxFollowers:      xtwitterscraper.Int(0),
			MaxFollowing:      xtwitterscraper.Int(0),
			MaxStatuses:       xtwitterscraper.Int(0),
			MinAccountAgeDays: xtwitterscraper.Int(0),
			MinFollowers:      xtwitterscraper.Int(0),
			MinFollowing:      xtwitterscraper.Int(0),
			MinStatuses:       xtwitterscraper.Int(0),
			Mode:              xtwitterscraper.XUserGetVerifiedFollowersParamsModeStandard,
			PageSize:          xtwitterscraper.Int(1),
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

func TestXUserGetAffiliatesWithOptionalParams(t *testing.T) {
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
	_, err := client.X.Users.GetAffiliates(
		context.TODO(),
		"id",
		xtwitterscraper.XUserGetAffiliatesParams{
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
			MaxFollowers:      xtwitterscraper.Int(0),
			MaxFollowing:      xtwitterscraper.Int(0),
			MaxStatuses:       xtwitterscraper.Int(0),
			MinAccountAgeDays: xtwitterscraper.Int(0),
			MinFollowers:      xtwitterscraper.Int(0),
			MinFollowing:      xtwitterscraper.Int(0),
			MinStatuses:       xtwitterscraper.Int(0),
			Omit:              xtwitterscraper.String("entities,media"),
			PageSize:          xtwitterscraper.Int(1),
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

func TestXUserGetArticlesWithOptionalParams(t *testing.T) {
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
	_, err := client.X.Users.GetArticles(
		context.TODO(),
		"id",
		xtwitterscraper.XUserGetArticlesParams{
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
			MediaType:         xtwitterscraper.XUserGetArticlesParamsMediaTypeImages,
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
			Quotes:            xtwitterscraper.XUserGetArticlesParamsQuotesInclude,
			QuotesOfTweetID:   xtwitterscraper.String("quotesOfTweetId"),
			Replies:           xtwitterscraper.XUserGetArticlesParamsRepliesInclude,
			Retweets:          xtwitterscraper.XUserGetArticlesParamsRetweetsInclude,
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

func TestXUserGetBatchTweetsWithOptionalParams(t *testing.T) {
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
	_, err := client.X.Users.GetBatchTweets(context.TODO(), xtwitterscraper.XUserGetBatchTweetsParams{
		Fields:    xtwitterscraper.String("id,text,author.username"),
		IDs:       xtwitterscraper.String("ids"),
		Omit:      xtwitterscraper.String("entities,media"),
		PageSize:  xtwitterscraper.Int(1),
		Usernames: xtwitterscraper.String("nasa,@OpenAI,x.com/verge"),
	})
	if err != nil {
		var apierr *xtwitterscraper.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestXUserGetFollowerIDsWithOptionalParams(t *testing.T) {
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
	_, err := client.X.Users.GetFollowerIDs(
		context.TODO(),
		"id",
		xtwitterscraper.XUserGetFollowerIDsParams{
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

func TestXUserGetFollowingIDsWithOptionalParams(t *testing.T) {
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
	_, err := client.X.Users.GetFollowingIDs(
		context.TODO(),
		"id",
		xtwitterscraper.XUserGetFollowingIDsParams{
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

func TestXUserGetHighlightsWithOptionalParams(t *testing.T) {
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
	_, err := client.X.Users.GetHighlights(
		context.TODO(),
		"id",
		xtwitterscraper.XUserGetHighlightsParams{
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
			MediaType:         xtwitterscraper.XUserGetHighlightsParamsMediaTypeImages,
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
			Quotes:            xtwitterscraper.XUserGetHighlightsParamsQuotesInclude,
			QuotesOfTweetID:   xtwitterscraper.String("quotesOfTweetId"),
			Replies:           xtwitterscraper.XUserGetHighlightsParamsRepliesInclude,
			Retweets:          xtwitterscraper.XUserGetHighlightsParamsRetweetsInclude,
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

func TestXUserGetInsightsWithOptionalParams(t *testing.T) {
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
	_, err := client.X.Users.GetInsights(
		context.TODO(),
		"id",
		xtwitterscraper.XUserGetInsightsParams{
			Fields: xtwitterscraper.String("id,text,author.username"),
			Omit:   xtwitterscraper.String("entities,media"),
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

func TestXUserGetListMembershipsWithOptionalParams(t *testing.T) {
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
	_, err := client.X.Users.GetListMemberships(
		context.TODO(),
		"id",
		xtwitterscraper.XUserGetListMembershipsParams{
			Cursor:   xtwitterscraper.String("cursor"),
			Fields:   xtwitterscraper.String("id,text,author.username"),
			Omit:     xtwitterscraper.String("entities,media"),
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

func TestXUserGetListsWithOptionalParams(t *testing.T) {
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
	_, err := client.X.Users.GetLists(
		context.TODO(),
		"id",
		xtwitterscraper.XUserGetListsParams{
			Cursor:   xtwitterscraper.String("cursor"),
			Fields:   xtwitterscraper.String("id,text,author.username"),
			Omit:     xtwitterscraper.String("entities,media"),
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

func TestXUserGetLiveWithOptionalParams(t *testing.T) {
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
	_, err := client.X.Users.GetLive(
		context.TODO(),
		"id",
		xtwitterscraper.XUserGetLiveParams{
			Fields: xtwitterscraper.String("id,text,author.username"),
			Omit:   xtwitterscraper.String("entities,media"),
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

func TestXUserGetSimilarWithOptionalParams(t *testing.T) {
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
	_, err := client.X.Users.GetSimilar(
		context.TODO(),
		"id",
		xtwitterscraper.XUserGetSimilarParams{
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
			MaxFollowers:      xtwitterscraper.Int(0),
			MaxFollowing:      xtwitterscraper.Int(0),
			MaxStatuses:       xtwitterscraper.Int(0),
			MinAccountAgeDays: xtwitterscraper.Int(0),
			MinFollowers:      xtwitterscraper.Int(0),
			MinFollowing:      xtwitterscraper.Int(0),
			MinStatuses:       xtwitterscraper.Int(0),
			Omit:              xtwitterscraper.String("entities,media"),
			PageSize:          xtwitterscraper.Int(1),
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

func TestXUserGetSubscriptionsWithOptionalParams(t *testing.T) {
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
	_, err := client.X.Users.GetSubscriptions(
		context.TODO(),
		"id",
		xtwitterscraper.XUserGetSubscriptionsParams{
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
			MaxFollowers:      xtwitterscraper.Int(0),
			MaxFollowing:      xtwitterscraper.Int(0),
			MaxStatuses:       xtwitterscraper.Int(0),
			MinAccountAgeDays: xtwitterscraper.Int(0),
			MinFollowers:      xtwitterscraper.Int(0),
			MinFollowing:      xtwitterscraper.Int(0),
			MinStatuses:       xtwitterscraper.Int(0),
			Omit:              xtwitterscraper.String("entities,media"),
			PageSize:          xtwitterscraper.Int(1),
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

func TestXUserTranslateBioWithOptionalParams(t *testing.T) {
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
	_, err := client.X.Users.TranslateBio(
		context.TODO(),
		"id",
		xtwitterscraper.XUserTranslateBioParams{
			Language: "es",
			Fields:   xtwitterscraper.String("id,text,author.username"),
			Omit:     xtwitterscraper.String("entities,media"),
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
