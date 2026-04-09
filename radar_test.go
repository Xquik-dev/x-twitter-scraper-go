// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package xtwitterscraper_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/x-twitter-scraper-go"
	"github.com/stainless-sdks/x-twitter-scraper-go/internal/testutil"
	"github.com/stainless-sdks/x-twitter-scraper-go/option"
)

func TestRadarGetTrendingTopicsWithOptionalParams(t *testing.T) {
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
	)
	_, err := client.Radar.GetTrendingTopics(context.TODO(), xtwitterscraper.RadarGetTrendingTopicsParams{
		Category: xtwitterscraper.String("category"),
		Count:    xtwitterscraper.Int(0),
		Hours:    xtwitterscraper.Int(0),
		Region:   xtwitterscraper.String("region"),
		Source:   xtwitterscraper.RadarGetTrendingTopicsParamsSourceGitHub,
	})
	if err != nil {
		var apierr *xtwitterscraper.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
