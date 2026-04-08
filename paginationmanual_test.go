// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package xtwitterscraper_test

import (
	"context"
	"os"
	"testing"

	"github.com/stainless-sdks/x-twitter-scraper-go"
	"github.com/stainless-sdks/x-twitter-scraper-go/internal/testutil"
	"github.com/stainless-sdks/x-twitter-scraper-go/option"
)

func TestManualPagination(t *testing.T) {
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
	page, err := client.X.Communities.Tweets.List(context.TODO(), xtwitterscraper.XCommunityTweetListParams{
		Q: "q",
	})
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	for _, tweet := range page.data {
		t.Logf("%+v\n", tweet.HasNextPage)
	}
	// The mock server isn't going to give us real pagination
	page, err = page.GetNextPage()
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	if page != nil {
		for _, tweet := range page.data {
			t.Logf("%+v\n", tweet.HasNextPage)
		}
	}
}
