// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

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

func TestXJobGetWithOptionalParams(t *testing.T) {
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
	_, err := client.X.Jobs.Get(
		context.TODO(),
		"id",
		xtwitterscraper.XJobGetParams{
			Fields: xtwitterscraper.String("id,text,author.username"),
			Omit:   xtwitterscraper.String("entities,media"),
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

func TestXJobListLocationsWithOptionalParams(t *testing.T) {
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
	_, err := client.X.Jobs.ListLocations(context.TODO(), xtwitterscraper.XJobListLocationsParams{
		Q:      "new york",
		Fields: xtwitterscraper.String("id,text,author.username"),
		Limit:  xtwitterscraper.Int(5),
		Omit:   xtwitterscraper.String("entities,media"),
	})
	if err != nil {
		var apierr *xtwitterscraper.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestXJobSearchWithOptionalParams(t *testing.T) {
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
	_, err := client.X.Jobs.Search(context.TODO(), xtwitterscraper.XJobSearchParams{
		Q:              "engineer",
		Company:        xtwitterscraper.String("SpaceX"),
		Cursor:         xtwitterscraper.String("cursor"),
		EmploymentType: xtwitterscraper.String("full_time"),
		Fields:         xtwitterscraper.String("id,text,author.username"),
		Location:       xtwitterscraper.String("Berlin"),
		LocationID:     xtwitterscraper.String("1720462281849397612"),
		LocationType:   xtwitterscraper.String("remote,hybrid"),
		Omit:           xtwitterscraper.String("entities,media"),
		PageSize:       xtwitterscraper.Int(1),
		Seniority:      xtwitterscraper.String("senior,lead"),
	})
	if err != nil {
		var apierr *xtwitterscraper.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
