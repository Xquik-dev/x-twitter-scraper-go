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

	"github.com/stainless-sdks/x-twitter-scraper-go"
	"github.com/stainless-sdks/x-twitter-scraper-go/internal/testutil"
	"github.com/stainless-sdks/x-twitter-scraper-go/option"
)

func TestExtractionGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Extractions.Get(
		context.TODO(),
		"id",
		xtwitterscraper.ExtractionGetParams{
			After: xtwitterscraper.String("after"),
			Limit: xtwitterscraper.Int(1),
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
	_, err := client.Extractions.List(context.TODO(), xtwitterscraper.ExtractionListParams{
		After:    xtwitterscraper.String("after"),
		Limit:    xtwitterscraper.Int(1),
		Status:   xtwitterscraper.ExtractionListParamsStatusRunning,
		ToolType: xtwitterscraper.ExtractionListParamsToolTypeArticleExtractor,
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
	_, err := client.Extractions.EstimateCost(context.TODO(), xtwitterscraper.ExtractionEstimateCostParams{
		ToolType:          xtwitterscraper.ExtractionEstimateCostParamsToolTypeArticleExtractor,
		AdvancedQuery:     xtwitterscraper.String("advancedQuery"),
		ExactPhrase:       xtwitterscraper.String("exactPhrase"),
		ExcludeWords:      xtwitterscraper.String("excludeWords"),
		SearchQuery:       xtwitterscraper.String("searchQuery"),
		TargetCommunityID: xtwitterscraper.String("targetCommunityId"),
		TargetListID:      xtwitterscraper.String("targetListId"),
		TargetSpaceID:     xtwitterscraper.String("targetSpaceId"),
		TargetTweetID:     xtwitterscraper.String("targetTweetId"),
		TargetUsername:    xtwitterscraper.String("targetUsername"),
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
	_, err := client.Extractions.Run(context.TODO(), xtwitterscraper.ExtractionRunParams{
		ToolType:          xtwitterscraper.ExtractionRunParamsToolTypeArticleExtractor,
		AdvancedQuery:     xtwitterscraper.String("advancedQuery"),
		ExactPhrase:       xtwitterscraper.String("exactPhrase"),
		ExcludeWords:      xtwitterscraper.String("excludeWords"),
		SearchQuery:       xtwitterscraper.String("searchQuery"),
		TargetCommunityID: xtwitterscraper.String("targetCommunityId"),
		TargetListID:      xtwitterscraper.String("targetListId"),
		TargetSpaceID:     xtwitterscraper.String("targetSpaceId"),
		TargetTweetID:     xtwitterscraper.String("targetTweetId"),
		TargetUsername:    xtwitterscraper.String("targetUsername"),
	})
	if err != nil {
		var apierr *xtwitterscraper.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
