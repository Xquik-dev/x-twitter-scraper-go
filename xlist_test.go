// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

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

func TestXListGetFollowersWithOptionalParams(t *testing.T) {
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
	_, err := client.X.Lists.GetFollowers(
		context.TODO(),
		"id",
		xtwitterscraper.XListGetFollowersParams{
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
			Mode:              xtwitterscraper.XListGetFollowersParamsModeStandard,
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

func TestXListGetMembersWithOptionalParams(t *testing.T) {
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
	_, err := client.X.Lists.GetMembers(
		context.TODO(),
		"id",
		xtwitterscraper.XListGetMembersParams{
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
			Mode:              xtwitterscraper.XListGetMembersParamsModeStandard,
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

func TestXListGetTweetsWithOptionalParams(t *testing.T) {
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
	_, err := client.X.Lists.GetTweets(
		context.TODO(),
		"id",
		xtwitterscraper.XListGetTweetsParams{
			AnyWords:         xtwitterscraper.String("anyWords"),
			BlueVerifiedOnly: xtwitterscraper.Bool(true),
			Cashtags:         xtwitterscraper.String("cashtags"),
			Cursor:           xtwitterscraper.String("cursor"),
			ExactPhrase:      xtwitterscraper.String("exactPhrase"),
			ExcludeWords:     xtwitterscraper.String("excludeWords"),
			FromUser:         xtwitterscraper.String("fromUser"),
			Hashtags:         xtwitterscraper.String("hashtags"),
			IncludeReplies:   xtwitterscraper.Bool(true),
			Language:         xtwitterscraper.String("language"),
			MaxFaves:         xtwitterscraper.Int(0),
			MaxQuotes:        xtwitterscraper.Int(0),
			MaxReplies:       xtwitterscraper.Int(0),
			MaxRetweets:      xtwitterscraper.Int(0),
			MediaType:        xtwitterscraper.XListGetTweetsParamsMediaTypeImages,
			Mentioning:       xtwitterscraper.String("mentioning"),
			MinBookmarks:     xtwitterscraper.Int(0),
			MinLikes:         xtwitterscraper.Int(0),
			MinQuotes:        xtwitterscraper.Int(0),
			MinReplies:       xtwitterscraper.Int(0),
			MinRetweets:      xtwitterscraper.Int(0),
			MinViews:         xtwitterscraper.Int(0),
			Mode:             xtwitterscraper.XListGetTweetsParamsModeStandard,
			NativeRetweets:   xtwitterscraper.Bool(true),
			PageSize:         xtwitterscraper.Int(1),
			Replies:          xtwitterscraper.XListGetTweetsParamsRepliesInclude,
			Retweets:         xtwitterscraper.XListGetTweetsParamsRetweetsInclude,
			SinceDate:        xtwitterscraper.Time(time.Now()),
			SinceTime:        xtwitterscraper.String("sinceTime"),
			ToUser:           xtwitterscraper.String("toUser"),
			UntilDate:        xtwitterscraper.Time(time.Now()),
			UntilTime:        xtwitterscraper.String("untilTime"),
			VerifiedOnly:     xtwitterscraper.Bool(true),
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
