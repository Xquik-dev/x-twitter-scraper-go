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

func TestComposeNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Compose.New(context.TODO(), xtwitterscraper.ComposeNewParams{
		Step:              xtwitterscraper.ComposeNewParamsStepCompose,
		AdditionalContext: xtwitterscraper.String("additionalContext"),
		CallToAction:      xtwitterscraper.String("callToAction"),
		Draft:             xtwitterscraper.String("draft"),
		Goal:              xtwitterscraper.ComposeNewParamsGoalEngagement,
		HasLink:           xtwitterscraper.Bool(true),
		HasMedia:          xtwitterscraper.Bool(true),
		MediaType:         xtwitterscraper.ComposeNewParamsMediaTypePhoto,
		StyleUsername:     xtwitterscraper.String("styleUsername"),
		Tone:              xtwitterscraper.String("tone"),
		Topic:             xtwitterscraper.String("topic"),
	})
	if err != nil {
		var apierr *xtwitterscraper.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
