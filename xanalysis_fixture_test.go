// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

package xtwitterscraper_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/Xquik-dev/x-twitter-scraper-go"
)

func TestXAnalysisSentimentFixture(t *testing.T) {
	client, request := fixtureClient(t, "TweetAnalysisPage.json")
	res, err := client.X.Analysis.Sentiment(context.Background(), xtwitterscraper.XAnalysisSentimentParams{
		Query:     xtwitterscraper.String("Sony WH-1000XM6 lang:en"),
		QueryType: xtwitterscraper.XAnalysisSentimentParamsQueryTypeLatest,
		SinceTime: xtwitterscraper.XAnalysisSentimentParamsSinceTimeUnion{OfString: xtwitterscraper.String("2026-09-01T00:00:00Z")},
		UntilTime: xtwitterscraper.XAnalysisSentimentParamsUntilTimeUnion{OfInt: xtwitterscraper.Int(1790812799)},
		Limit:     xtwitterscraper.Int(20),
		Analysis: xtwitterscraper.XAnalysisSentimentParamsAnalysis{
			Context: xtwitterscraper.String("Consumer headphones & customer service."),
			Targets: []xtwitterscraper.XAnalysisSentimentParamsAnalysisTarget{{Name: "Sony", Aliases: []string{"WH-1000XM6"}}},
		},
		Baseline: []xtwitterscraper.XAnalysisSentimentParamsBaseline{{
			ID: xtwitterscraper.String("2102882545427714266"),
			Answers: map[string]xtwitterscraper.XAnalysisSentimentParamsBaselineAnswersUnion{
				"sentiment": {OfString: xtwitterscraper.String("positive")},
				"intensity": {OfFloat: xtwitterscraper.Float(1)},
			},
		}},
	})
	if err != nil {
		t.Fatal(err)
	}
	request.expect(t, http.MethodPost, "/api/v1/x/analysis/sentiment", "")
	expectDecoded(t, res)
	request.expectBody(t, `{
		"query": "Sony WH-1000XM6 lang:en",
		"queryType": "Latest",
		"sinceTime": "2026-09-01T00:00:00Z",
		"untilTime": 1790812799,
		"limit": 20,
		"analysis": {"context": "Consumer headphones & customer service.", "targets": [{"name": "Sony", "aliases": ["WH-1000XM6"]}]},
		"baseline": [{"id": "2102882545427714266", "answers": {"sentiment": "positive", "intensity": 1}}]
	}`)
	if len(res.Results) != 1 || len(res.Results[0].Analysis.Answers) != 3 || len(res.Unanalyzed) != 1 {
		t.Fatalf("decoded %s", res.RawJSON())
	}
	result := res.Results[0]
	answer := result.Analysis.Answers[0]
	if result.Tweet.ID != "1893456789012345678" ||
		result.Answers["sentiment"].AsString() != "mixed" ||
		result.Answers["intensity"].AsFloat() != 1 ||
		result.Answers["sarcasm"].AsFloat() != 0.03 ||
		answer.QuestionID != "sentiment" ||
		answer.Value.AsString() != "mixed" ||
		answer.Confidence != 0.88 ||
		answer.Probabilities["mixed"] != 0.88 ||
		result.Analysis.Answers[1].Value.AsFloat() != 1 ||
		result.Analysis.ContextBytes != 612 ||
		result.Analysis.ContextAvailability.Author != "available" ||
		result.Monitor.Status != "changed" ||
		result.Monitor.Changes[0].Previous != "positive" ||
		res.Unanalyzed[0].Reason != "post_unavailable" ||
		res.AnalysisSummary.Rows.Analyzed != 1 ||
		res.AnalysisSummary.Engagement != 37 ||
		res.HasNextPage ||
		res.NextCursor != "" {
		t.Fatalf("decoded %s", res.RawJSON())
	}
}

func TestXAnalysisViralScoreFixture(t *testing.T) {
	client, request := fixtureClient(t, "TweetAnalysisViralPage.json")
	res, err := client.X.Analysis.ViralScore(context.Background(), xtwitterscraper.XAnalysisViralScoreParams{
		TweetIDs: []string{"2102882545427714266", "https://x.com/NASA/status/2102761519985332442"},
	})
	if err != nil {
		t.Fatal(err)
	}
	request.expect(t, http.MethodPost, "/api/v1/x/analysis/viral-score", "")
	expectDecoded(t, res)
	request.expectBody(t, `{"tweetIds": ["2102882545427714266", "https://x.com/NASA/status/2102761519985332442"]}`)
	if len(res.Results) != 1 || len(res.AnalysisSummary.Cashtags) != 1 || len(res.AnalysisSummary.Targets) != 1 {
		t.Fatalf("decoded %s", res.RawJSON())
	}
	result := res.Results[0]
	summary := res.AnalysisSummary
	if result.JSON.ViralScore.Valid() ||
		result.JSON.ViralScore.Raw() != "null" ||
		result.ViralVerdict != "send_it" ||
		result.ViralAlgorithmScore != 18.4 ||
		result.ViralActualEngagementRate != 1.2874 ||
		result.Viral["score"] != float64(72) ||
		result.Viral["weights"] != "viral_lite:1" ||
		summary.JSON.Monitor.Valid() ||
		summary.JSON.Monitor.Raw() != "null" ||
		summary.Viral["averageScore"] != float64(58) ||
		summary.Cashtags[0].Cashtag != "$NVDA" ||
		summary.Cashtags[0].Rows != 3 ||
		summary.Cashtags[0].Signal.Bullish != 2 ||
		summary.Cashtags[0].Signal.Score != 0.3333 ||
		summary.Targets[0].Name != "Sony" ||
		summary.Targets[0].Mentions != 14 ||
		summary.Targets[0].Share != 0.7 ||
		summary.Targets[0].Engagement != 1200 {
		t.Fatalf("decoded %s", res.RawJSON())
	}
}

func TestXAnalysisBrandFixture(t *testing.T) {
	client, request := fixtureClient(t, "TweetAnalysisPage.json")
	res, err := client.X.Analysis.Brand(context.Background(), xtwitterscraper.XAnalysisBrandParams{
		Username: xtwitterscraper.String("@NASA"),
		Limit:    xtwitterscraper.Int(20),
	})
	if err != nil {
		t.Fatal(err)
	}
	request.expect(t, http.MethodPost, "/api/v1/x/analysis/brand", "")
	expectDecoded(t, res)
	request.expectBody(t, `{"username": "@NASA", "limit": 20}`)
	if len(res.Results) != 1 || res.Results[0].Analysis.ContextBytes != 612 || res.Unanalyzed[0].ID != "1893456789012345679" {
		t.Fatalf("decoded %s", res.RawJSON())
	}
}

func TestXAnalysisNewsFixture(t *testing.T) {
	client, request := fixtureClient(t, "TweetAnalysisPage.json")
	res, err := client.X.Analysis.News(context.Background(), xtwitterscraper.XAnalysisNewsParams{
		ListID: xtwitterscraper.String("4143408"),
		Cursor: xtwitterscraper.String("xqcv1_tkGoEx1UJRqeog9a2BJdmzUZTMCoU_M8kb"),
	})
	if err != nil {
		t.Fatal(err)
	}
	request.expect(t, http.MethodPost, "/api/v1/x/analysis/news", "")
	expectDecoded(t, res)
	request.expectBody(t, `{"listId": "4143408", "cursor": "xqcv1_tkGoEx1UJRqeog9a2BJdmzUZTMCoU_M8kb"}`)
	if len(res.Results) != 1 || res.Results[0].Monitor.ConfigurationFingerprint != "9f2c4e1b7a5d3c8e6f0a1b2c3d4e5f60718293a4b5c6d7e8f90a1b2c3d4e5f6a" {
		t.Fatalf("decoded %s", res.RawJSON())
	}
}

func TestXAnalysisMarketSignalsFixture(t *testing.T) {
	client, request := fixtureClient(t, "TweetAnalysisPage.json")
	res, err := client.X.Analysis.MarketSignals(context.Background(), xtwitterscraper.XAnalysisMarketSignalsParams{
		QuotesOf: xtwitterscraper.String("2102882545427714266"),
	})
	if err != nil {
		t.Fatal(err)
	}
	request.expect(t, http.MethodPost, "/api/v1/x/analysis/market-signals", "")
	expectDecoded(t, res)
	request.expectBody(t, `{"quotesOf": "2102882545427714266"}`)
	if len(res.Results) != 1 || res.AnalysisSummary.SchemaVersion != 1 || res.AnalysisSummary.Rows.Skipped != 0 {
		t.Fatalf("decoded %s", res.RawJSON())
	}
}

func TestXAnalysisClassifyFixture(t *testing.T) {
	client, request := fixtureClient(t, "TweetAnalysisPage.json")
	res, err := client.X.Analysis.Classify(context.Background(), xtwitterscraper.XAnalysisClassifyParams{
		Texts: []string{"We shipped dark mode today. Tell us what breaks."},
		Analysis: xtwitterscraper.XAnalysisClassifyParamsAnalysis{
			Questions: []xtwitterscraper.XAnalysisClassifyParamsAnalysisQuestion{{
				ID:           "request",
				Instructions: "Does the post ask for something?",
				Type:         "probability",
				Version:      "1",
				Criteria:     xtwitterscraper.XAnalysisClassifyParamsAnalysisQuestionCriteria{Yes: "It asks for help or a change.", No: "It asks for nothing."},
			}},
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	request.expect(t, http.MethodPost, "/api/v1/x/analysis/classify", "")
	expectDecoded(t, res)
	request.expectBody(t, `{
		"texts": ["We shipped dark mode today. Tell us what breaks."],
		"analysis": {"questions": [{"id": "request", "instructions": "Does the post ask for something?", "type": "probability", "version": "1", "criteria": {"yes": "It asks for help or a change.", "no": "It asks for nothing."}}]}
	}`)
	if len(res.Results) != 1 || len(res.Results[0].Analysis.Answers) != 3 || res.Results[0].Analysis.Answers[2].Probability != 0.03 {
		t.Fatalf("decoded %s", res.RawJSON())
	}
}
