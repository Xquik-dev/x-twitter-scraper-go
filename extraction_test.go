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
			Cursor:       xtwitterscraper.String("cursor"),
			FieldStyle:   xtwitterscraper.ExtractionGetParamsFieldStyleSource,
			IncludeRaw:   xtwitterscraper.Bool(true),
			Limit:        xtwitterscraper.Int(1),
			OutputMode:   xtwitterscraper.ExtractionGetParamsOutputModeCompact,
			OutputPreset: xtwitterscraper.ExtractionGetParamsOutputPresetNested,
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
		ToolType:              xtwitterscraper.ExtractionEstimateCostParamsToolTypeFollowerExplorer,
		AdvancedQuery:         xtwitterscraper.String("min_faves:100"),
		AnyWords:              xtwitterscraper.String("ChatGPT AI model"),
		BioContains:           xtwitterscraper.String("bioContains"),
		BlueVerifiedOnly:      xtwitterscraper.Bool(true),
		BoundingBox:           xtwitterscraper.String("-74.1 40.6 -73.9 40.8"),
		CardName:              xtwitterscraper.String("cardName"),
		Cashtags:              xtwitterscraper.String("$TSLA $NVDA"),
		CollectionStrategy:    xtwitterscraper.ExtractionEstimateCostParamsCollectionStrategyAuto,
		ConversationID:        xtwitterscraper.String("1234567890"),
		DedupeAcrossTargets:   xtwitterscraper.Bool(true),
		DedupeMode:            xtwitterscraper.ExtractionEstimateCostParamsDedupeModeNone,
		ExactPhrase:           xtwitterscraper.String("artificial intelligence"),
		ExcludeOriginalAuthor: xtwitterscraper.Bool(true),
		ExcludeSource:         xtwitterscraper.String("excludeSource"),
		ExcludeWords:          xtwitterscraper.String("spam"),
		FromUser:              xtwitterscraper.String("nasa"),
		Geocode:               xtwitterscraper.String("geocode"),
		Hashtags:              xtwitterscraper.String("#AI startups"),
		HasLocation:           xtwitterscraper.Bool(true),
		HasMediaOnly:          xtwitterscraper.Bool(true),
		HasWebsite:            xtwitterscraper.Bool(true),
		IncludeOriginalPost:   xtwitterscraper.Bool(true),
		IncludeSearchTerms:    xtwitterscraper.Bool(true),
		IncludeTargetMetadata: xtwitterscraper.Bool(true),
		InReplyToTweetID:      xtwitterscraper.String("1234567890"),
		Language:              xtwitterscraper.String("en"),
		ListID:                xtwitterscraper.String("1234567890"),
		LocationContains:      xtwitterscraper.String("locationContains"),
		MaxDepth:              xtwitterscraper.Int(1),
		MaxFollowers:          xtwitterscraper.Int(0),
		MaxFollowing:          xtwitterscraper.Int(0),
		MaxID:                 xtwitterscraper.String("maxId"),
		MaxItemsPerTarget:     xtwitterscraper.Int(1),
		MaxLikes:              xtwitterscraper.Int(0),
		MaxPagesPerTarget:     xtwitterscraper.Int(1),
		MaxPosts:              xtwitterscraper.Int(0),
		MaxQuotes:             xtwitterscraper.Int(0),
		MaxReplies:            xtwitterscraper.Int(0),
		MaxRetweets:           xtwitterscraper.Int(0),
		MediaType:             xtwitterscraper.ExtractionEstimateCostParamsMediaTypeImages,
		Mentioning:            xtwitterscraper.String("example_user"),
		MinAccountAgeDays:     xtwitterscraper.Int(0),
		MinBookmarks:          xtwitterscraper.Int(0),
		MinFaves:              xtwitterscraper.Int(10),
		MinFollowers:          xtwitterscraper.Int(0),
		MinFollowing:          xtwitterscraper.Int(0),
		MinPosts:              xtwitterscraper.Int(0),
		MinQuotes:             xtwitterscraper.Int(2),
		MinReplies:            xtwitterscraper.Int(3),
		MinRetweets:           xtwitterscraper.Int(5),
		MinViews:              xtwitterscraper.Int(0),
		NativeRetweets:        xtwitterscraper.Bool(true),
		Near:                  xtwitterscraper.String("near"),
		News:                  xtwitterscraper.Bool(true),
		OverlapMode:           xtwitterscraper.Bool(true),
		Place:                 xtwitterscraper.String("96683cc9126741d1"),
		PlaceCountry:          xtwitterscraper.String("US"),
		PointRadius:           xtwitterscraper.String("-73.99 40.73 25mi"),
		QueryType:             xtwitterscraper.ExtractionEstimateCostParamsQueryTypeLatest,
		Quotes:                xtwitterscraper.ExtractionEstimateCostParamsQuotesInclude,
		QuotesOfTweetID:       xtwitterscraper.String("1234567890"),
		RelationTargets: []xtwitterscraper.ExtractionEstimateCostParamsRelationTarget{{
			Relation: "community_members",
			Value:    "x",
		}},
		Replies:           xtwitterscraper.ExtractionEstimateCostParamsRepliesInclude,
		ResultsLimit:      xtwitterscraper.Int(1000),
		Retweets:          xtwitterscraper.ExtractionEstimateCostParamsRetweetsExclude,
		RetweetsOfTweetID: xtwitterscraper.String("1234567890"),
		Safe:              xtwitterscraper.Bool(true),
		Scope:             xtwitterscraper.ExtractionEstimateCostParamsScopeAll,
		SearchQueries:     []string{"string"},
		SearchQuery:       xtwitterscraper.String("AI trends 2025"),
		SinceDate:         xtwitterscraper.Time(time.Now()),
		SinceID:           xtwitterscraper.String("sinceId"),
		SinceTime: xtwitterscraper.ExtractionEstimateCostParamsSinceTimeUnion{
			OfTime: xtwitterscraper.Time(time.Now()),
		},
		Sort:               xtwitterscraper.ExtractionEstimateCostParamsSortRelevance,
		Source:             xtwitterscraper.String("source"),
		StartCursor:        xtwitterscraper.String("x"),
		TargetCommunityID:  xtwitterscraper.String("1500000000000000000"),
		TargetCommunityIDs: []string{"string"},
		TargetListID:       xtwitterscraper.String("1234567890"),
		TargetListIDs:      []string{"string"},
		Targets: []xtwitterscraper.ExtractionEstimateCostParamsTargetUnion{{
			OfString: xtwitterscraper.String("string"),
		}},
		TargetSpaceID:   xtwitterscraper.String("1vOGwMdBqpwGB"),
		TargetTweetID:   xtwitterscraper.String("1234567890"),
		TargetTweetIDs:  []string{"string"},
		TargetUsername:  xtwitterscraper.String("elonmusk"),
		TargetUsernames: []string{"string"},
		ToUser:          xtwitterscraper.String("openai"),
		UntilDate:       xtwitterscraper.Time(time.Now()),
		UntilTime: xtwitterscraper.ExtractionEstimateCostParamsUntilTimeUnion{
			OfTime: xtwitterscraper.Time(time.Now()),
		},
		URL:              xtwitterscraper.String("example.com"),
		UsernameContains: xtwitterscraper.String("usernameContains"),
		VerifiedOnly:     xtwitterscraper.Bool(false),
		VerifiedType:     xtwitterscraper.String("verifiedType"),
		Within:           xtwitterscraper.String("within"),
		WithinTime:       xtwitterscraper.String("withinTime"),
	})
	if err != nil {
		var apierr *xtwitterscraper.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestExtractionExportResultsWithOptionalParams(t *testing.T) {
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
			Format:         xtwitterscraper.ExtractionExportResultsParamsFormatCsv,
			HasDescription: xtwitterscraper.Bool(true),
			HasLocation:    xtwitterscraper.Bool(true),
			HasMedia:       xtwitterscraper.Bool(true),
			Lang:           xtwitterscraper.String("lang"),
			MaxFollowers:   xtwitterscraper.Int(0),
			MaxFollowing:   xtwitterscraper.Int(0),
			MaxPosts:       xtwitterscraper.Int(0),
			MinFollowers:   xtwitterscraper.Int(0),
			MinFollowing:   xtwitterscraper.Int(0),
			MinLikes:       xtwitterscraper.Int(0),
			MinPosts:       xtwitterscraper.Int(0),
			MinReplies:     xtwitterscraper.Int(0),
			MinRetweets:    xtwitterscraper.Int(0),
			MinViews:       xtwitterscraper.Int(0),
			Search:         xtwitterscraper.String("search"),
			SinceDate:      xtwitterscraper.Time(time.Now()),
			UntilDate:      xtwitterscraper.Time(time.Now()),
			Verified:       xtwitterscraper.Bool(true),
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
		ToolType:              xtwitterscraper.ExtractionRunParamsToolTypeFollowerExplorer,
		DryRun:                xtwitterscraper.Bool(true),
		AdvancedQuery:         xtwitterscraper.String("min_faves:100"),
		AnyWords:              xtwitterscraper.String("ChatGPT AI model"),
		BioContains:           xtwitterscraper.String("bioContains"),
		BlueVerifiedOnly:      xtwitterscraper.Bool(true),
		BoundingBox:           xtwitterscraper.String("-74.1 40.6 -73.9 40.8"),
		CardName:              xtwitterscraper.String("cardName"),
		Cashtags:              xtwitterscraper.String("$TSLA $NVDA"),
		CollectionStrategy:    xtwitterscraper.ExtractionRunParamsCollectionStrategyAuto,
		ConversationID:        xtwitterscraper.String("1234567890"),
		DedupeAcrossTargets:   xtwitterscraper.Bool(true),
		DedupeMode:            xtwitterscraper.ExtractionRunParamsDedupeModeNone,
		ExactPhrase:           xtwitterscraper.String("artificial intelligence"),
		ExcludeOriginalAuthor: xtwitterscraper.Bool(true),
		ExcludeSource:         xtwitterscraper.String("excludeSource"),
		ExcludeWords:          xtwitterscraper.String("spam"),
		FromUser:              xtwitterscraper.String("nasa"),
		Geocode:               xtwitterscraper.String("geocode"),
		Hashtags:              xtwitterscraper.String("#AI startups"),
		HasLocation:           xtwitterscraper.Bool(true),
		HasMediaOnly:          xtwitterscraper.Bool(true),
		HasWebsite:            xtwitterscraper.Bool(true),
		IncludeOriginalPost:   xtwitterscraper.Bool(true),
		IncludeSearchTerms:    xtwitterscraper.Bool(true),
		IncludeTargetMetadata: xtwitterscraper.Bool(true),
		InReplyToTweetID:      xtwitterscraper.String("1234567890"),
		Language:              xtwitterscraper.String("en"),
		ListID:                xtwitterscraper.String("1234567890"),
		LocationContains:      xtwitterscraper.String("locationContains"),
		MaxDepth:              xtwitterscraper.Int(1),
		MaxFollowers:          xtwitterscraper.Int(0),
		MaxFollowing:          xtwitterscraper.Int(0),
		MaxID:                 xtwitterscraper.String("maxId"),
		MaxItemsPerTarget:     xtwitterscraper.Int(1),
		MaxLikes:              xtwitterscraper.Int(0),
		MaxPagesPerTarget:     xtwitterscraper.Int(1),
		MaxPosts:              xtwitterscraper.Int(0),
		MaxQuotes:             xtwitterscraper.Int(0),
		MaxReplies:            xtwitterscraper.Int(0),
		MaxRetweets:           xtwitterscraper.Int(0),
		MediaType:             xtwitterscraper.ExtractionRunParamsMediaTypeImages,
		Mentioning:            xtwitterscraper.String("example_user"),
		MinAccountAgeDays:     xtwitterscraper.Int(0),
		MinBookmarks:          xtwitterscraper.Int(0),
		MinFaves:              xtwitterscraper.Int(10),
		MinFollowers:          xtwitterscraper.Int(0),
		MinFollowing:          xtwitterscraper.Int(0),
		MinPosts:              xtwitterscraper.Int(0),
		MinQuotes:             xtwitterscraper.Int(2),
		MinReplies:            xtwitterscraper.Int(3),
		MinRetweets:           xtwitterscraper.Int(5),
		MinViews:              xtwitterscraper.Int(0),
		NativeRetweets:        xtwitterscraper.Bool(true),
		Near:                  xtwitterscraper.String("near"),
		News:                  xtwitterscraper.Bool(true),
		OverlapMode:           xtwitterscraper.Bool(true),
		Place:                 xtwitterscraper.String("96683cc9126741d1"),
		PlaceCountry:          xtwitterscraper.String("US"),
		PointRadius:           xtwitterscraper.String("-73.99 40.73 25mi"),
		QueryType:             xtwitterscraper.ExtractionRunParamsQueryTypeLatest,
		Quotes:                xtwitterscraper.ExtractionRunParamsQuotesInclude,
		QuotesOfTweetID:       xtwitterscraper.String("1234567890"),
		RelationTargets: []xtwitterscraper.ExtractionRunParamsRelationTarget{{
			Relation: "community_members",
			Value:    "x",
		}},
		Replies:           xtwitterscraper.ExtractionRunParamsRepliesInclude,
		ResultsLimit:      xtwitterscraper.Int(1000),
		Retweets:          xtwitterscraper.ExtractionRunParamsRetweetsExclude,
		RetweetsOfTweetID: xtwitterscraper.String("1234567890"),
		Safe:              xtwitterscraper.Bool(true),
		Scope:             xtwitterscraper.ExtractionRunParamsScopeAll,
		SearchQueries:     []string{"string"},
		SearchQuery:       xtwitterscraper.String("AI trends 2025"),
		SinceDate:         xtwitterscraper.Time(time.Now()),
		SinceID:           xtwitterscraper.String("sinceId"),
		SinceTime: xtwitterscraper.ExtractionRunParamsSinceTimeUnion{
			OfTime: xtwitterscraper.Time(time.Now()),
		},
		Sort:               xtwitterscraper.ExtractionRunParamsSortRelevance,
		Source:             xtwitterscraper.String("source"),
		StartCursor:        xtwitterscraper.String("x"),
		TargetCommunityID:  xtwitterscraper.String("1500000000000000000"),
		TargetCommunityIDs: []string{"string"},
		TargetListID:       xtwitterscraper.String("1234567890"),
		TargetListIDs:      []string{"string"},
		Targets: []xtwitterscraper.ExtractionRunParamsTargetUnion{{
			OfString: xtwitterscraper.String("string"),
		}},
		TargetSpaceID:   xtwitterscraper.String("1vOGwMdBqpwGB"),
		TargetTweetID:   xtwitterscraper.String("1234567890"),
		TargetTweetIDs:  []string{"string"},
		TargetUsername:  xtwitterscraper.String("elonmusk"),
		TargetUsernames: []string{"string"},
		ToUser:          xtwitterscraper.String("openai"),
		UntilDate:       xtwitterscraper.Time(time.Now()),
		UntilTime: xtwitterscraper.ExtractionRunParamsUntilTimeUnion{
			OfTime: xtwitterscraper.Time(time.Now()),
		},
		URL:              xtwitterscraper.String("example.com"),
		UsernameContains: xtwitterscraper.String("usernameContains"),
		VerifiedOnly:     xtwitterscraper.Bool(false),
		VerifiedType:     xtwitterscraper.String("verifiedType"),
		Within:           xtwitterscraper.String("within"),
		WithinTime:       xtwitterscraper.String("withinTime"),
	})
	if err != nil {
		var apierr *xtwitterscraper.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
