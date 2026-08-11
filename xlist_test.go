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
			Cursor:         xtwitterscraper.String("cursor"),
			IncludeReplies: xtwitterscraper.Bool(true),
			PageSize:       xtwitterscraper.Int(1),
			SinceTime:      xtwitterscraper.String("sinceTime"),
			UntilTime:      xtwitterscraper.String("untilTime"),
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
