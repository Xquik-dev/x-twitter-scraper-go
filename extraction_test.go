// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package xtwitterscraper_test

import (
	"bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/Xquik-dev/x-twitter-scraper-go"
	"github.com/Xquik-dev/x-twitter-scraper-go/internal/testutil"
	"github.com/Xquik-dev/x-twitter-scraper-go/option"
)

func TestExtractionGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Extractions.Get(
		context.TODO(),
		"id",
		xtwitterscraper.ExtractionGetParams{
			Cursor: xtwitterscraper.String("cursor"),
			Limit:  xtwitterscraper.Int(1),
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

func TestExtractionListWithOptionalParams(t *testing.T) {
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
	_, err := client.Extractions.List(context.TODO(), xtwitterscraper.ExtractionListParams{
		Cursor:   xtwitterscraper.String("cursor"),
		Limit:    xtwitterscraper.Int(1),
		Status:   xtwitterscraper.ExtractionListParamsStatusRunning,
		ToolType: xtwitterscraper.ExtractionListParamsToolTypeFollowerExplorer,
	})
	if err != nil {
		var apierr *xtwitterscraper.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestExtractionEstimateCostWithOptionalParams(t *testing.T) {
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
	_, err := client.Extractions.EstimateCost(context.TODO(), xtwitterscraper.ExtractionEstimateCostParams{
		ToolType:          xtwitterscraper.ExtractionEstimateCostParamsToolTypeFollowerExplorer,
		AdvancedQuery:     xtwitterscraper.String("min_faves:100"),
		AnyWords:          xtwitterscraper.String("ChatGPT AI model"),
		BoundingBox:       xtwitterscraper.String("-74.1 40.6 -73.9 40.8"),
		Cashtags:          xtwitterscraper.String("$TSLA $NVDA"),
		ConversationID:    xtwitterscraper.String("1234567890"),
		ExactPhrase:       xtwitterscraper.String("artificial intelligence"),
		ExcludeWords:      xtwitterscraper.String("spam"),
		FromUser:          xtwitterscraper.String("nasa"),
		Hashtags:          xtwitterscraper.String("#AI startups"),
		InReplyToTweetID:  xtwitterscraper.String("1234567890"),
		Language:          xtwitterscraper.String("en"),
		ListID:            xtwitterscraper.String("1234567890"),
		MediaType:         xtwitterscraper.ExtractionEstimateCostParamsMediaTypeImages,
		Mentioning:        xtwitterscraper.String("example_user"),
		MinFaves:          xtwitterscraper.Int(10),
		MinQuotes:         xtwitterscraper.Int(2),
		MinReplies:        xtwitterscraper.Int(3),
		MinRetweets:       xtwitterscraper.Int(5),
		Place:             xtwitterscraper.String("96683cc9126741d1"),
		PlaceCountry:      xtwitterscraper.String("US"),
		PointRadius:       xtwitterscraper.String("-73.99 40.73 25mi"),
		Quotes:            xtwitterscraper.ExtractionEstimateCostParamsQuotesInclude,
		QuotesOfTweetID:   xtwitterscraper.String("1234567890"),
		Replies:           xtwitterscraper.ExtractionEstimateCostParamsRepliesInclude,
		ResultsLimit:      xtwitterscraper.Int(1000),
		Retweets:          xtwitterscraper.ExtractionEstimateCostParamsRetweetsExclude,
		RetweetsOfTweetID: xtwitterscraper.String("1234567890"),
		SearchQuery:       xtwitterscraper.String("AI trends 2025"),
		SinceDate:         xtwitterscraper.Time(time.Now()),
		TargetCommunityID: xtwitterscraper.String("1500000000000000000"),
		TargetListID:      xtwitterscraper.String("1234567890"),
		TargetSpaceID:     xtwitterscraper.String("1vOGwMdBqpwGB"),
		TargetTweetID:     xtwitterscraper.String("1234567890"),
		TargetUsername:    xtwitterscraper.String("elonmusk"),
		ToUser:            xtwitterscraper.String("openai"),
		UntilDate:         xtwitterscraper.Time(time.Now()),
		URL:               xtwitterscraper.String("example.com"),
		VerifiedOnly:      xtwitterscraper.Bool(false),
	})
	if err != nil {
		var apierr *xtwitterscraper.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestExtractionExportResults(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("abc"))
	}))
	defer server.Close()
	baseURL := server.URL
	client := xtwitterscraper.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
		option.WithBearerToken("My Bearer Token"),
	)
	resp, err := client.Extractions.ExportResults(
		context.TODO(),
		"id",
		xtwitterscraper.ExtractionExportResultsParams{
			Format: xtwitterscraper.ExtractionExportResultsParamsFormatCsv,
		},
	)
	if err != nil {
		var apierr *xtwitterscraper.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		var apierr *xtwitterscraper.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
	if !bytes.Equal(b, []byte("abc")) {
		t.Fatalf("return value not %s: %s", "abc", b)
	}
}

func TestExtractionRunWithOptionalParams(t *testing.T) {
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
	_, err := client.Extractions.Run(context.TODO(), xtwitterscraper.ExtractionRunParams{
		ToolType:          xtwitterscraper.ExtractionRunParamsToolTypeFollowerExplorer,
		AdvancedQuery:     xtwitterscraper.String("min_faves:100"),
		AnyWords:          xtwitterscraper.String("ChatGPT AI model"),
		BoundingBox:       xtwitterscraper.String("-74.1 40.6 -73.9 40.8"),
		Cashtags:          xtwitterscraper.String("$TSLA $NVDA"),
		ConversationID:    xtwitterscraper.String("1234567890"),
		ExactPhrase:       xtwitterscraper.String("artificial intelligence"),
		ExcludeWords:      xtwitterscraper.String("spam"),
		FromUser:          xtwitterscraper.String("nasa"),
		Hashtags:          xtwitterscraper.String("#AI startups"),
		InReplyToTweetID:  xtwitterscraper.String("1234567890"),
		Language:          xtwitterscraper.String("en"),
		ListID:            xtwitterscraper.String("1234567890"),
		MediaType:         xtwitterscraper.ExtractionRunParamsMediaTypeImages,
		Mentioning:        xtwitterscraper.String("example_user"),
		MinFaves:          xtwitterscraper.Int(10),
		MinQuotes:         xtwitterscraper.Int(2),
		MinReplies:        xtwitterscraper.Int(3),
		MinRetweets:       xtwitterscraper.Int(5),
		Place:             xtwitterscraper.String("96683cc9126741d1"),
		PlaceCountry:      xtwitterscraper.String("US"),
		PointRadius:       xtwitterscraper.String("-73.99 40.73 25mi"),
		Quotes:            xtwitterscraper.ExtractionRunParamsQuotesInclude,
		QuotesOfTweetID:   xtwitterscraper.String("1234567890"),
		Replies:           xtwitterscraper.ExtractionRunParamsRepliesInclude,
		ResultsLimit:      xtwitterscraper.Int(1000),
		Retweets:          xtwitterscraper.ExtractionRunParamsRetweetsExclude,
		RetweetsOfTweetID: xtwitterscraper.String("1234567890"),
		SearchQuery:       xtwitterscraper.String("AI trends 2025"),
		SinceDate:         xtwitterscraper.Time(time.Now()),
		TargetCommunityID: xtwitterscraper.String("1500000000000000000"),
		TargetListID:      xtwitterscraper.String("1234567890"),
		TargetSpaceID:     xtwitterscraper.String("1vOGwMdBqpwGB"),
		TargetTweetID:     xtwitterscraper.String("1234567890"),
		TargetUsername:    xtwitterscraper.String("elonmusk"),
		ToUser:            xtwitterscraper.String("openai"),
		UntilDate:         xtwitterscraper.Time(time.Now()),
		URL:               xtwitterscraper.String("example.com"),
		VerifiedOnly:      xtwitterscraper.Bool(false),
	})
	if err != nil {
		var apierr *xtwitterscraper.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
