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
	"github.com/Xquik-dev/x-twitter-scraper-go/packages/param"
)

func TestXAnalysisBrandWithOptionalParams(t *testing.T) {
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
	_, err := client.X.Analysis.Brand(context.TODO(), xtwitterscraper.XAnalysisBrandParams{
		Analysis: xtwitterscraper.XAnalysisBrandParamsAnalysis{
			Concurrency:     xtwitterscraper.Int(16),
			Context:         xtwitterscraper.String("Consumer headphones & customer service."),
			MaxContextBytes: xtwitterscraper.Int(64000),
			Preset:          "brand",
			Questions: []xtwitterscraper.XAnalysisBrandParamsAnalysisQuestion{{
				ID:           "request",
				Instructions: "What does the author ask for?",
				Type:         "choice",
				Version:      "1",
				Categories:   map[string]param.Opt[string]{"foo": xtwitterscraper.String("Help with a problem.")},
				Criteria: xtwitterscraper.XAnalysisBrandParamsAnalysisQuestionCriteria{
					No:  "Asks for no refund.",
					Yes: "Asks for a refund.",
				},
				Levels: []string{"No urgency."},
			}},
			Targets: []xtwitterscraper.XAnalysisBrandParamsAnalysisTarget{{
				Name:    "Sony",
				Aliases: []string{"WH-1000XM6"},
			}},
		},
		Baseline: []xtwitterscraper.XAnalysisBrandParamsBaseline{{
			Answers: map[string]xtwitterscraper.XAnalysisBrandParamsBaselineAnswersUnion{"foo": {
				OfString: xtwitterscraper.String("positive"),
			}},
			ID: xtwitterscraper.String("1893456789012345678"),
			Monitor: xtwitterscraper.XAnalysisBrandParamsBaselineMonitor{
				ConfigurationFingerprint: xtwitterscraper.String("9f2c4e1b7a5d3c8e6f0a1b2c3d4e5f60718293a4b5c6d7e8f90a1b2c3d4e5f6a"),
			},
			Tweet: xtwitterscraper.XAnalysisBrandParamsBaselineTweet{
				ID: "1893456789012345678",
			},
		}},
		Cursor:    xtwitterscraper.String("DAACCgACGRElMJcAAA"),
		Limit:     xtwitterscraper.Int(20),
		ListID:    xtwitterscraper.String("1234567890"),
		Query:     xtwitterscraper.String("Sony WH-1000XM6 lang:en"),
		QueryType: xtwitterscraper.XAnalysisBrandParamsQueryTypeLatest,
		QuotesOf:  xtwitterscraper.String("1893456789012345678"),
		RepliesTo: xtwitterscraper.String("1893456789012345678"),
		SinceTime: xtwitterscraper.XAnalysisBrandParamsSinceTimeUnion{
			OfString: xtwitterscraper.String("2026-09-25T00:00:00Z"),
		},
		Texts:    []string{"We shipped dark mode today. Tell us what breaks."},
		TweetIDs: []string{"1893456789012345678"},
		UntilTime: xtwitterscraper.XAnalysisBrandParamsUntilTimeUnion{
			OfString: xtwitterscraper.String("2026-09-26T00:00:00Z"),
		},
		Username: xtwitterscraper.String("@NASA"),
	})
	if err != nil {
		var apierr *xtwitterscraper.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestXAnalysisClassifyWithOptionalParams(t *testing.T) {
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
	_, err := client.X.Analysis.Classify(context.TODO(), xtwitterscraper.XAnalysisClassifyParams{
		Analysis: xtwitterscraper.XAnalysisClassifyParamsAnalysis{
			Concurrency:     xtwitterscraper.Int(16),
			Context:         xtwitterscraper.String("Consumer headphones & customer service."),
			MaxContextBytes: xtwitterscraper.Int(64000),
			Preset:          "brand",
			Questions: []xtwitterscraper.XAnalysisClassifyParamsAnalysisQuestion{{
				ID:           "request",
				Instructions: "What does the author ask for?",
				Type:         "choice",
				Version:      "1",
				Categories:   map[string]param.Opt[string]{"foo": xtwitterscraper.String("Help with a problem.")},
				Criteria: xtwitterscraper.XAnalysisClassifyParamsAnalysisQuestionCriteria{
					No:  "Asks for no refund.",
					Yes: "Asks for a refund.",
				},
				Levels: []string{"No urgency."},
			}},
			Targets: []xtwitterscraper.XAnalysisClassifyParamsAnalysisTarget{{
				Name:    "Sony",
				Aliases: []string{"WH-1000XM6"},
			}},
		},
		Baseline: []xtwitterscraper.XAnalysisClassifyParamsBaseline{{
			Answers: map[string]xtwitterscraper.XAnalysisClassifyParamsBaselineAnswersUnion{"foo": {
				OfString: xtwitterscraper.String("positive"),
			}},
			ID: xtwitterscraper.String("1893456789012345678"),
			Monitor: xtwitterscraper.XAnalysisClassifyParamsBaselineMonitor{
				ConfigurationFingerprint: xtwitterscraper.String("9f2c4e1b7a5d3c8e6f0a1b2c3d4e5f60718293a4b5c6d7e8f90a1b2c3d4e5f6a"),
			},
			Tweet: xtwitterscraper.XAnalysisClassifyParamsBaselineTweet{
				ID: "1893456789012345678",
			},
		}},
		Cursor:    xtwitterscraper.String("DAACCgACGRElMJcAAA"),
		Limit:     xtwitterscraper.Int(20),
		ListID:    xtwitterscraper.String("1234567890"),
		Query:     xtwitterscraper.String("Sony WH-1000XM6 lang:en"),
		QueryType: xtwitterscraper.XAnalysisClassifyParamsQueryTypeLatest,
		QuotesOf:  xtwitterscraper.String("1893456789012345678"),
		RepliesTo: xtwitterscraper.String("1893456789012345678"),
		SinceTime: xtwitterscraper.XAnalysisClassifyParamsSinceTimeUnion{
			OfString: xtwitterscraper.String("2026-09-25T00:00:00Z"),
		},
		Texts:    []string{"We shipped dark mode today. Tell us what breaks."},
		TweetIDs: []string{"1893456789012345678"},
		UntilTime: xtwitterscraper.XAnalysisClassifyParamsUntilTimeUnion{
			OfString: xtwitterscraper.String("2026-09-26T00:00:00Z"),
		},
		Username: xtwitterscraper.String("@NASA"),
	})
	if err != nil {
		var apierr *xtwitterscraper.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestXAnalysisMarketSignalsWithOptionalParams(t *testing.T) {
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
	_, err := client.X.Analysis.MarketSignals(context.TODO(), xtwitterscraper.XAnalysisMarketSignalsParams{
		Analysis: xtwitterscraper.XAnalysisMarketSignalsParamsAnalysis{
			Concurrency:     xtwitterscraper.Int(16),
			Context:         xtwitterscraper.String("Consumer headphones & customer service."),
			MaxContextBytes: xtwitterscraper.Int(64000),
			Preset:          "brand",
			Questions: []xtwitterscraper.XAnalysisMarketSignalsParamsAnalysisQuestion{{
				ID:           "request",
				Instructions: "What does the author ask for?",
				Type:         "choice",
				Version:      "1",
				Categories:   map[string]param.Opt[string]{"foo": xtwitterscraper.String("Help with a problem.")},
				Criteria: xtwitterscraper.XAnalysisMarketSignalsParamsAnalysisQuestionCriteria{
					No:  "Asks for no refund.",
					Yes: "Asks for a refund.",
				},
				Levels: []string{"No urgency."},
			}},
			Targets: []xtwitterscraper.XAnalysisMarketSignalsParamsAnalysisTarget{{
				Name:    "Sony",
				Aliases: []string{"WH-1000XM6"},
			}},
		},
		Baseline: []xtwitterscraper.XAnalysisMarketSignalsParamsBaseline{{
			Answers: map[string]xtwitterscraper.XAnalysisMarketSignalsParamsBaselineAnswersUnion{"foo": {
				OfString: xtwitterscraper.String("positive"),
			}},
			ID: xtwitterscraper.String("1893456789012345678"),
			Monitor: xtwitterscraper.XAnalysisMarketSignalsParamsBaselineMonitor{
				ConfigurationFingerprint: xtwitterscraper.String("9f2c4e1b7a5d3c8e6f0a1b2c3d4e5f60718293a4b5c6d7e8f90a1b2c3d4e5f6a"),
			},
			Tweet: xtwitterscraper.XAnalysisMarketSignalsParamsBaselineTweet{
				ID: "1893456789012345678",
			},
		}},
		Cursor:    xtwitterscraper.String("DAACCgACGRElMJcAAA"),
		Limit:     xtwitterscraper.Int(20),
		ListID:    xtwitterscraper.String("1234567890"),
		Query:     xtwitterscraper.String("Sony WH-1000XM6 lang:en"),
		QueryType: xtwitterscraper.XAnalysisMarketSignalsParamsQueryTypeLatest,
		QuotesOf:  xtwitterscraper.String("1893456789012345678"),
		RepliesTo: xtwitterscraper.String("1893456789012345678"),
		SinceTime: xtwitterscraper.XAnalysisMarketSignalsParamsSinceTimeUnion{
			OfString: xtwitterscraper.String("2026-09-25T00:00:00Z"),
		},
		Texts:    []string{"We shipped dark mode today. Tell us what breaks."},
		TweetIDs: []string{"1893456789012345678"},
		UntilTime: xtwitterscraper.XAnalysisMarketSignalsParamsUntilTimeUnion{
			OfString: xtwitterscraper.String("2026-09-26T00:00:00Z"),
		},
		Username: xtwitterscraper.String("@NASA"),
	})
	if err != nil {
		var apierr *xtwitterscraper.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestXAnalysisNewsWithOptionalParams(t *testing.T) {
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
	_, err := client.X.Analysis.News(context.TODO(), xtwitterscraper.XAnalysisNewsParams{
		Analysis: xtwitterscraper.XAnalysisNewsParamsAnalysis{
			Concurrency:     xtwitterscraper.Int(16),
			Context:         xtwitterscraper.String("Consumer headphones & customer service."),
			MaxContextBytes: xtwitterscraper.Int(64000),
			Preset:          "brand",
			Questions: []xtwitterscraper.XAnalysisNewsParamsAnalysisQuestion{{
				ID:           "request",
				Instructions: "What does the author ask for?",
				Type:         "choice",
				Version:      "1",
				Categories:   map[string]param.Opt[string]{"foo": xtwitterscraper.String("Help with a problem.")},
				Criteria: xtwitterscraper.XAnalysisNewsParamsAnalysisQuestionCriteria{
					No:  "Asks for no refund.",
					Yes: "Asks for a refund.",
				},
				Levels: []string{"No urgency."},
			}},
			Targets: []xtwitterscraper.XAnalysisNewsParamsAnalysisTarget{{
				Name:    "Sony",
				Aliases: []string{"WH-1000XM6"},
			}},
		},
		Baseline: []xtwitterscraper.XAnalysisNewsParamsBaseline{{
			Answers: map[string]xtwitterscraper.XAnalysisNewsParamsBaselineAnswersUnion{"foo": {
				OfString: xtwitterscraper.String("positive"),
			}},
			ID: xtwitterscraper.String("1893456789012345678"),
			Monitor: xtwitterscraper.XAnalysisNewsParamsBaselineMonitor{
				ConfigurationFingerprint: xtwitterscraper.String("9f2c4e1b7a5d3c8e6f0a1b2c3d4e5f60718293a4b5c6d7e8f90a1b2c3d4e5f6a"),
			},
			Tweet: xtwitterscraper.XAnalysisNewsParamsBaselineTweet{
				ID: "1893456789012345678",
			},
		}},
		Cursor:    xtwitterscraper.String("DAACCgACGRElMJcAAA"),
		Limit:     xtwitterscraper.Int(20),
		ListID:    xtwitterscraper.String("1234567890"),
		Query:     xtwitterscraper.String("Sony WH-1000XM6 lang:en"),
		QueryType: xtwitterscraper.XAnalysisNewsParamsQueryTypeLatest,
		QuotesOf:  xtwitterscraper.String("1893456789012345678"),
		RepliesTo: xtwitterscraper.String("1893456789012345678"),
		SinceTime: xtwitterscraper.XAnalysisNewsParamsSinceTimeUnion{
			OfString: xtwitterscraper.String("2026-09-25T00:00:00Z"),
		},
		Texts:    []string{"We shipped dark mode today. Tell us what breaks."},
		TweetIDs: []string{"1893456789012345678"},
		UntilTime: xtwitterscraper.XAnalysisNewsParamsUntilTimeUnion{
			OfString: xtwitterscraper.String("2026-09-26T00:00:00Z"),
		},
		Username: xtwitterscraper.String("@NASA"),
	})
	if err != nil {
		var apierr *xtwitterscraper.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestXAnalysisSentimentWithOptionalParams(t *testing.T) {
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
	_, err := client.X.Analysis.Sentiment(context.TODO(), xtwitterscraper.XAnalysisSentimentParams{
		Analysis: xtwitterscraper.XAnalysisSentimentParamsAnalysis{
			Concurrency:     xtwitterscraper.Int(16),
			Context:         xtwitterscraper.String("Consumer headphones & customer service."),
			MaxContextBytes: xtwitterscraper.Int(64000),
			Preset:          "brand",
			Questions: []xtwitterscraper.XAnalysisSentimentParamsAnalysisQuestion{{
				ID:           "request",
				Instructions: "What does the author ask for?",
				Type:         "choice",
				Version:      "1",
				Categories:   map[string]param.Opt[string]{"foo": xtwitterscraper.String("Help with a problem.")},
				Criteria: xtwitterscraper.XAnalysisSentimentParamsAnalysisQuestionCriteria{
					No:  "Asks for no refund.",
					Yes: "Asks for a refund.",
				},
				Levels: []string{"No urgency."},
			}},
			Targets: []xtwitterscraper.XAnalysisSentimentParamsAnalysisTarget{{
				Name:    "Sony",
				Aliases: []string{"WH-1000XM6"},
			}},
		},
		Baseline: []xtwitterscraper.XAnalysisSentimentParamsBaseline{{
			Answers: map[string]xtwitterscraper.XAnalysisSentimentParamsBaselineAnswersUnion{"foo": {
				OfString: xtwitterscraper.String("positive"),
			}},
			ID: xtwitterscraper.String("1893456789012345678"),
			Monitor: xtwitterscraper.XAnalysisSentimentParamsBaselineMonitor{
				ConfigurationFingerprint: xtwitterscraper.String("9f2c4e1b7a5d3c8e6f0a1b2c3d4e5f60718293a4b5c6d7e8f90a1b2c3d4e5f6a"),
			},
			Tweet: xtwitterscraper.XAnalysisSentimentParamsBaselineTweet{
				ID: "1893456789012345678",
			},
		}},
		Cursor:    xtwitterscraper.String("DAACCgACGRElMJcAAA"),
		Limit:     xtwitterscraper.Int(20),
		ListID:    xtwitterscraper.String("1234567890"),
		Query:     xtwitterscraper.String("Sony WH-1000XM6 lang:en"),
		QueryType: xtwitterscraper.XAnalysisSentimentParamsQueryTypeLatest,
		QuotesOf:  xtwitterscraper.String("1893456789012345678"),
		RepliesTo: xtwitterscraper.String("1893456789012345678"),
		SinceTime: xtwitterscraper.XAnalysisSentimentParamsSinceTimeUnion{
			OfString: xtwitterscraper.String("2026-09-25T00:00:00Z"),
		},
		Texts:    []string{"We shipped dark mode today. Tell us what breaks."},
		TweetIDs: []string{"1893456789012345678"},
		UntilTime: xtwitterscraper.XAnalysisSentimentParamsUntilTimeUnion{
			OfString: xtwitterscraper.String("2026-09-26T00:00:00Z"),
		},
		Username: xtwitterscraper.String("@NASA"),
	})
	if err != nil {
		var apierr *xtwitterscraper.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestXAnalysisViralScoreWithOptionalParams(t *testing.T) {
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
	_, err := client.X.Analysis.ViralScore(context.TODO(), xtwitterscraper.XAnalysisViralScoreParams{
		Analysis: xtwitterscraper.XAnalysisViralScoreParamsAnalysis{
			Concurrency:     xtwitterscraper.Int(16),
			Context:         xtwitterscraper.String("Consumer headphones & customer service."),
			MaxContextBytes: xtwitterscraper.Int(64000),
			Preset:          "brand",
			Questions: []xtwitterscraper.XAnalysisViralScoreParamsAnalysisQuestion{{
				ID:           "request",
				Instructions: "What does the author ask for?",
				Type:         "choice",
				Version:      "1",
				Categories:   map[string]param.Opt[string]{"foo": xtwitterscraper.String("Help with a problem.")},
				Criteria: xtwitterscraper.XAnalysisViralScoreParamsAnalysisQuestionCriteria{
					No:  "Asks for no refund.",
					Yes: "Asks for a refund.",
				},
				Levels: []string{"No urgency."},
			}},
			Targets: []xtwitterscraper.XAnalysisViralScoreParamsAnalysisTarget{{
				Name:    "Sony",
				Aliases: []string{"WH-1000XM6"},
			}},
		},
		Baseline: []xtwitterscraper.XAnalysisViralScoreParamsBaseline{{
			Answers: map[string]xtwitterscraper.XAnalysisViralScoreParamsBaselineAnswersUnion{"foo": {
				OfString: xtwitterscraper.String("positive"),
			}},
			ID: xtwitterscraper.String("1893456789012345678"),
			Monitor: xtwitterscraper.XAnalysisViralScoreParamsBaselineMonitor{
				ConfigurationFingerprint: xtwitterscraper.String("9f2c4e1b7a5d3c8e6f0a1b2c3d4e5f60718293a4b5c6d7e8f90a1b2c3d4e5f6a"),
			},
			Tweet: xtwitterscraper.XAnalysisViralScoreParamsBaselineTweet{
				ID: "1893456789012345678",
			},
		}},
		Cursor:    xtwitterscraper.String("DAACCgACGRElMJcAAA"),
		Limit:     xtwitterscraper.Int(20),
		ListID:    xtwitterscraper.String("1234567890"),
		Query:     xtwitterscraper.String("Sony WH-1000XM6 lang:en"),
		QueryType: xtwitterscraper.XAnalysisViralScoreParamsQueryTypeLatest,
		QuotesOf:  xtwitterscraper.String("1893456789012345678"),
		RepliesTo: xtwitterscraper.String("1893456789012345678"),
		SinceTime: xtwitterscraper.XAnalysisViralScoreParamsSinceTimeUnion{
			OfString: xtwitterscraper.String("2026-09-25T00:00:00Z"),
		},
		Texts:    []string{"We shipped dark mode today. Tell us what breaks."},
		TweetIDs: []string{"1893456789012345678"},
		UntilTime: xtwitterscraper.XAnalysisViralScoreParamsUntilTimeUnion{
			OfString: xtwitterscraper.String("2026-09-26T00:00:00Z"),
		},
		Username: xtwitterscraper.String("@NASA"),
	})
	if err != nil {
		var apierr *xtwitterscraper.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
