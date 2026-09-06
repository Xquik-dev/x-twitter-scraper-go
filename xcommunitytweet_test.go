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

func TestXCommunityTweetListWithOptionalParams(t *testing.T) {
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
	_, err := client.X.Communities.Tweets.List(context.TODO(), xtwitterscraper.XCommunityTweetListParams{
		CommunityID:  "321669910225",
		Q:            "q",
		Cursor:       xtwitterscraper.String("cursor"),
		Language:     xtwitterscraper.String("language"),
		MediaType:    xtwitterscraper.XCommunityTweetListParamsMediaTypeImages,
		MinLikes:     xtwitterscraper.Int(0),
		MinReplies:   xtwitterscraper.Int(0),
		MinRetweets:  xtwitterscraper.Int(0),
		MinViews:     xtwitterscraper.Int(0),
		PageSize:     xtwitterscraper.Int(1),
		QueryType:    xtwitterscraper.XCommunityTweetListParamsQueryTypeLatest,
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

func TestXCommunityTweetListByCommunityWithOptionalParams(t *testing.T) {
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
	_, err := client.X.Communities.Tweets.ListByCommunity(
		context.TODO(),
		"id",
		xtwitterscraper.XCommunityTweetListByCommunityParams{
			Cursor:       xtwitterscraper.String("cursor"),
			Language:     xtwitterscraper.String("language"),
			MediaType:    xtwitterscraper.XCommunityTweetListByCommunityParamsMediaTypeImages,
			MinLikes:     xtwitterscraper.Int(0),
			MinReplies:   xtwitterscraper.Int(0),
			MinRetweets:  xtwitterscraper.Int(0),
			MinViews:     xtwitterscraper.Int(0),
			PageSize:     xtwitterscraper.Int(1),
			SinceDate:    xtwitterscraper.Time(time.Now()),
			UntilDate:    xtwitterscraper.Time(time.Now()),
			VerifiedOnly: xtwitterscraper.Bool(true),
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
