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
			Wait:         xtwitterscraper.Int(0),
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
		Status:   xtwitterscraper.ExtractionListParamsStatusPending,
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

func TestExtractionCancel(t *testing.T) {
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
	_, err := client.Extractions.Cancel(context.TODO(), "id")
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
		AdvancedQuery:         xtwitterscraper.String("advancedQuery"),
		AnyWords:              xtwitterscraper.String("anyWords"),
		BioContains:           xtwitterscraper.String("bioContains"),
		BlueVerifiedOnly:      xtwitterscraper.Bool(true),
		BoundingBox:           xtwitterscraper.String("boundingBox"),
		CardName:              xtwitterscraper.String("cardName"),
		Cashtags:              xtwitterscraper.String("cashtags"),
		CollectionStrategy:    xtwitterscraper.ExtractionEstimateCostParamsCollectionStrategyAuto,
		ConversationID:        xtwitterscraper.String("conversationId"),
		DedupeAcrossTargets:   xtwitterscraper.Bool(true),
		DedupeMode:            xtwitterscraper.ExtractionEstimateCostParamsDedupeModeNone,
		ExactPhrase:           xtwitterscraper.String("exactPhrase"),
		ExcludeOriginalAuthor: xtwitterscraper.Bool(true),
		ExcludeSource:         xtwitterscraper.String("excludeSource"),
		ExcludeWords:          xtwitterscraper.String("excludeWords"),
		FromUser:              xtwitterscraper.String("fromUser"),
		Geocode:               xtwitterscraper.String("geocode"),
		Hashtags:              xtwitterscraper.String("hashtags"),
		HasLocation:           xtwitterscraper.Bool(true),
		HasMediaOnly:          xtwitterscraper.Bool(true),
		HasWebsite:            xtwitterscraper.Bool(true),
		IncludeOriginalPost:   xtwitterscraper.Bool(true),
		IncludeSearchTerms:    xtwitterscraper.Bool(true),
		IncludeTargetMetadata: xtwitterscraper.Bool(true),
		InReplyToTweetID:      xtwitterscraper.String("inReplyToTweetId"),
		Language:              xtwitterscraper.String("language"),
		ListID:                xtwitterscraper.String("listId"),
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
		Mentioning:            xtwitterscraper.String("mentioning"),
		MinAccountAgeDays:     xtwitterscraper.Int(0),
		MinBookmarks:          xtwitterscraper.Int(0),
		MinFaves:              xtwitterscraper.Int(0),
		MinFollowers:          xtwitterscraper.Int(0),
		MinFollowing:          xtwitterscraper.Int(0),
		MinPosts:              xtwitterscraper.Int(0),
		MinQuotes:             xtwitterscraper.Int(0),
		MinReplies:            xtwitterscraper.Int(0),
		MinRetweets:           xtwitterscraper.Int(0),
		MinViews:              xtwitterscraper.Int(0),
		NativeRetweets:        xtwitterscraper.Bool(true),
		Near:                  xtwitterscraper.String("near"),
		News:                  xtwitterscraper.Bool(true),
		OverlapMode:           xtwitterscraper.Bool(true),
		Place:                 xtwitterscraper.String("place"),
		PlaceCountry:          xtwitterscraper.String("placeCountry"),
		PointRadius:           xtwitterscraper.String("pointRadius"),
		QueryType:             xtwitterscraper.ExtractionEstimateCostParamsQueryTypeLatest,
		Quotes:                xtwitterscraper.ExtractionEstimateCostParamsQuotesInclude,
		QuotesOfTweetID:       xtwitterscraper.String("quotesOfTweetId"),
		RelationTargets: []xtwitterscraper.ExtractionEstimateCostParamsRelationTarget{{
			Relation: "community_members",
			Value:    "x",
		}},
		Replies:           xtwitterscraper.ExtractionEstimateCostParamsRepliesInclude,
		ResultsLimit:      xtwitterscraper.Int(1),
		Retweets:          xtwitterscraper.ExtractionEstimateCostParamsRetweetsInclude,
		RetweetsOfTweetID: xtwitterscraper.String("retweetsOfTweetId"),
		Safe:              xtwitterscraper.Bool(true),
		Scope:             xtwitterscraper.ExtractionEstimateCostParamsScopeAll,
		SearchQueries:     []string{"string"},
		SearchQuery:       xtwitterscraper.String("searchQuery"),
		SinceDate:         xtwitterscraper.Time(time.Now()),
		SinceID:           xtwitterscraper.String("sinceId"),
		SinceTime: xtwitterscraper.ExtractionEstimateCostParamsSinceTimeUnion{
			OfTime: xtwitterscraper.Time(time.Now()),
		},
		Sort:               xtwitterscraper.ExtractionEstimateCostParamsSortRelevance,
		Source:             xtwitterscraper.String("source"),
		StartCursor:        xtwitterscraper.String("x"),
		TargetCommunityID:  xtwitterscraper.String("targetCommunityId"),
		TargetCommunityIDs: []string{"string"},
		TargetListID:       xtwitterscraper.String("targetListId"),
		TargetListIDs:      []string{"string"},
		Targets: []xtwitterscraper.ExtractionEstimateCostParamsTargetUnion{{
			OfString: xtwitterscraper.String("string"),
		}},
		TargetSpaceID:   xtwitterscraper.String("targetSpaceId"),
		TargetTweetID:   xtwitterscraper.String("targetTweetId"),
		TargetTweetIDs:  []string{"string"},
		TargetUsername:  xtwitterscraper.String("elonmusk"),
		TargetUsernames: []string{"string"},
		ToUser:          xtwitterscraper.String("toUser"),
		UntilDate:       xtwitterscraper.Time(time.Now()),
		UntilTime: xtwitterscraper.ExtractionEstimateCostParamsUntilTimeUnion{
			OfTime: xtwitterscraper.Time(time.Now()),
		},
		URL:              xtwitterscraper.String("url"),
		UsernameContains: xtwitterscraper.String("usernameContains"),
		VerifiedOnly:     xtwitterscraper.Bool(true),
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
		AdvancedQuery:         xtwitterscraper.String("advancedQuery"),
		AnyWords:              xtwitterscraper.String("anyWords"),
		BioContains:           xtwitterscraper.String("bioContains"),
		BlueVerifiedOnly:      xtwitterscraper.Bool(true),
		BoundingBox:           xtwitterscraper.String("boundingBox"),
		CardName:              xtwitterscraper.String("cardName"),
		Cashtags:              xtwitterscraper.String("cashtags"),
		CollectionStrategy:    xtwitterscraper.ExtractionRunParamsCollectionStrategyAuto,
		ConversationID:        xtwitterscraper.String("conversationId"),
		DedupeAcrossTargets:   xtwitterscraper.Bool(true),
		DedupeMode:            xtwitterscraper.ExtractionRunParamsDedupeModeNone,
		ExactPhrase:           xtwitterscraper.String("exactPhrase"),
		ExcludeOriginalAuthor: xtwitterscraper.Bool(true),
		ExcludeSource:         xtwitterscraper.String("excludeSource"),
		ExcludeWords:          xtwitterscraper.String("excludeWords"),
		FromUser:              xtwitterscraper.String("fromUser"),
		Geocode:               xtwitterscraper.String("geocode"),
		Hashtags:              xtwitterscraper.String("hashtags"),
		HasLocation:           xtwitterscraper.Bool(true),
		HasMediaOnly:          xtwitterscraper.Bool(true),
		HasWebsite:            xtwitterscraper.Bool(true),
		IncludeOriginalPost:   xtwitterscraper.Bool(true),
		IncludeSearchTerms:    xtwitterscraper.Bool(true),
		IncludeTargetMetadata: xtwitterscraper.Bool(true),
		InReplyToTweetID:      xtwitterscraper.String("inReplyToTweetId"),
		Language:              xtwitterscraper.String("language"),
		ListID:                xtwitterscraper.String("listId"),
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
		Mentioning:            xtwitterscraper.String("mentioning"),
		MinAccountAgeDays:     xtwitterscraper.Int(0),
		MinBookmarks:          xtwitterscraper.Int(0),
		MinFaves:              xtwitterscraper.Int(0),
		MinFollowers:          xtwitterscraper.Int(0),
		MinFollowing:          xtwitterscraper.Int(0),
		MinPosts:              xtwitterscraper.Int(0),
		MinQuotes:             xtwitterscraper.Int(0),
		MinReplies:            xtwitterscraper.Int(0),
		MinRetweets:           xtwitterscraper.Int(0),
		MinViews:              xtwitterscraper.Int(0),
		NativeRetweets:        xtwitterscraper.Bool(true),
		Near:                  xtwitterscraper.String("near"),
		News:                  xtwitterscraper.Bool(true),
		OverlapMode:           xtwitterscraper.Bool(true),
		Place:                 xtwitterscraper.String("place"),
		PlaceCountry:          xtwitterscraper.String("placeCountry"),
		PointRadius:           xtwitterscraper.String("pointRadius"),
		QueryType:             xtwitterscraper.ExtractionRunParamsQueryTypeLatest,
		Quotes:                xtwitterscraper.ExtractionRunParamsQuotesInclude,
		QuotesOfTweetID:       xtwitterscraper.String("quotesOfTweetId"),
		RelationTargets: []xtwitterscraper.ExtractionRunParamsRelationTarget{{
			Relation: "community_members",
			Value:    "x",
		}},
		Replies:           xtwitterscraper.ExtractionRunParamsRepliesInclude,
		ResultsLimit:      xtwitterscraper.Int(1),
		Retweets:          xtwitterscraper.ExtractionRunParamsRetweetsInclude,
		RetweetsOfTweetID: xtwitterscraper.String("retweetsOfTweetId"),
		Safe:              xtwitterscraper.Bool(true),
		Scope:             xtwitterscraper.ExtractionRunParamsScopeAll,
		SearchQueries:     []string{"string"},
		SearchQuery:       xtwitterscraper.String("searchQuery"),
		SinceDate:         xtwitterscraper.Time(time.Now()),
		SinceID:           xtwitterscraper.String("sinceId"),
		SinceTime: xtwitterscraper.ExtractionRunParamsSinceTimeUnion{
			OfTime: xtwitterscraper.Time(time.Now()),
		},
		Sort:               xtwitterscraper.ExtractionRunParamsSortRelevance,
		Source:             xtwitterscraper.String("source"),
		StartCursor:        xtwitterscraper.String("x"),
		TargetCommunityID:  xtwitterscraper.String("targetCommunityId"),
		TargetCommunityIDs: []string{"string"},
		TargetListID:       xtwitterscraper.String("targetListId"),
		TargetListIDs:      []string{"string"},
		Targets: []xtwitterscraper.ExtractionRunParamsTargetUnion{{
			OfString: xtwitterscraper.String("string"),
		}},
		TargetSpaceID:   xtwitterscraper.String("targetSpaceId"),
		TargetTweetID:   xtwitterscraper.String("targetTweetId"),
		TargetTweetIDs:  []string{"string"},
		TargetUsername:  xtwitterscraper.String("elonmusk"),
		TargetUsernames: []string{"string"},
		ToUser:          xtwitterscraper.String("toUser"),
		UntilDate:       xtwitterscraper.Time(time.Now()),
		UntilTime: xtwitterscraper.ExtractionRunParamsUntilTimeUnion{
			OfTime: xtwitterscraper.Time(time.Now()),
		},
		URL:              xtwitterscraper.String("url"),
		UsernameContains: xtwitterscraper.String("usernameContains"),
		VerifiedOnly:     xtwitterscraper.Bool(true),
		VerifiedType:     xtwitterscraper.String("verifiedType"),
		Within:           xtwitterscraper.String("within"),
		WithinTime:       xtwitterscraper.String("withinTime"),
		IdempotencyKey:   xtwitterscraper.String("Idempotency-Key"),
	})
	if err != nil {
		var apierr *xtwitterscraper.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
