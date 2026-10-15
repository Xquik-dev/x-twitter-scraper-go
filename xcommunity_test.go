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

func TestXCommunityNewWithOptionalParams(t *testing.T) {
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
	_, err := client.X.Communities.New(context.TODO(), xtwitterscraper.XCommunityNewParams{
		Account:        "@elonmusk",
		Name:           "Example Name",
		IdempotencyKey: "Idempotency-Key",
		Description:    xtwitterscraper.String("A community for Tesla enthusiasts"),
	})
	if err != nil {
		var apierr *xtwitterscraper.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestXCommunityDelete(t *testing.T) {
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
	_, err := client.X.Communities.Delete(
		context.TODO(),
		"id",
		xtwitterscraper.XCommunityDeleteParams{
			Account:        "@elonmusk",
			CommunityName:  "Tesla Fans",
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

func TestXCommunityGetInfo(t *testing.T) {
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
	_, err := client.X.Communities.GetInfo(context.TODO(), "id")
	if err != nil {
		var apierr *xtwitterscraper.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestXCommunityGetMembersWithOptionalParams(t *testing.T) {
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
	_, err := client.X.Communities.GetMembers(
		context.TODO(),
		"id",
		xtwitterscraper.XCommunityGetMembersParams{
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

func TestXCommunityGetModeratorsWithOptionalParams(t *testing.T) {
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
	_, err := client.X.Communities.GetModerators(
		context.TODO(),
		"id",
		xtwitterscraper.XCommunityGetModeratorsParams{
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

func TestXCommunityGetSearchWithOptionalParams(t *testing.T) {
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
	_, err := client.X.Communities.GetSearch(context.TODO(), xtwitterscraper.XCommunityGetSearchParams{
		CommunityID:  "321669910225",
		Q:            "q",
		Cursor:       xtwitterscraper.String("cursor"),
		Language:     xtwitterscraper.String("language"),
		MediaType:    xtwitterscraper.XCommunityGetSearchParamsMediaTypeImages,
		MinLikes:     xtwitterscraper.Int(0),
		MinReplies:   xtwitterscraper.Int(0),
		MinRetweets:  xtwitterscraper.Int(0),
		MinViews:     xtwitterscraper.Int(0),
		PageSize:     xtwitterscraper.Int(1),
		QueryType:    xtwitterscraper.XCommunityGetSearchParamsQueryTypeLatest,
		SinceDate:    xtwitterscraper.Time(time.Now()),
		UntilDate:    xtwitterscraper.Time(time.Now()),
		VerifiedOnly: xtwitterscraper.Bool(true),
	})
	if err != nil {
		var apierr *xtwitterscraper.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestXCommunityFindWithOptionalParams(t *testing.T) {
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
	_, err := client.X.Communities.Find(context.TODO(), xtwitterscraper.XCommunityFindParams{
		Q:        "q",
		Cursor:   xtwitterscraper.String("cursor"),
		Fields:   xtwitterscraper.String("id,text,author.username"),
		Omit:     xtwitterscraper.String("entities,media"),
		PageSize: xtwitterscraper.Int(1),
	})
	if err != nil {
		var apierr *xtwitterscraper.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestXCommunityListPopularWithOptionalParams(t *testing.T) {
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
	_, err := client.X.Communities.ListPopular(context.TODO(), xtwitterscraper.XCommunityListPopularParams{
		Cursor:   xtwitterscraper.String("cursor"),
		Fields:   xtwitterscraper.String("id,text,author.username"),
		Omit:     xtwitterscraper.String("entities,media"),
		PageSize: xtwitterscraper.Int(1),
		Topic:    xtwitterscraper.String("topic"),
		TopicID:  xtwitterscraper.String("topicId"),
	})
	if err != nil {
		var apierr *xtwitterscraper.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestXCommunityListSuggestedWithOptionalParams(t *testing.T) {
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
	_, err := client.X.Communities.ListSuggested(context.TODO(), xtwitterscraper.XCommunityListSuggestedParams{
		Fields: xtwitterscraper.String("id,text,author.username"),
		Limit:  xtwitterscraper.Int(1),
		Omit:   xtwitterscraper.String("entities,media"),
	})
	if err != nil {
		var apierr *xtwitterscraper.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestXCommunityListTopicsWithOptionalParams(t *testing.T) {
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
	_, err := client.X.Communities.ListTopics(context.TODO(), xtwitterscraper.XCommunityListTopicsParams{
		Fields: xtwitterscraper.String("id,text,author.username"),
		Limit:  xtwitterscraper.Int(1),
		Omit:   xtwitterscraper.String("entities,media"),
	})
	if err != nil {
		var apierr *xtwitterscraper.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestXCommunitySearchMembersWithOptionalParams(t *testing.T) {
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
	_, err := client.X.Communities.SearchMembers(
		context.TODO(),
		"id",
		xtwitterscraper.XCommunitySearchMembersParams{
			Q:                 "q",
			BioContains:       xtwitterscraper.String("bioContains"),
			BlueVerifiedOnly:  xtwitterscraper.Bool(true),
			BusinessEmailOnly: xtwitterscraper.Bool(true),
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
