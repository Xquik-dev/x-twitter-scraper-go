// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

package xtwitterscraper_test

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/Xquik-dev/x-twitter-scraper-go"
)

func decodeFixture(t *testing.T, name string, model interface{ UnmarshalJSON([]byte) error }) {
	t.Helper()
	payload, err := fixtures.ReadFile("testdata/fixtures/" + name)
	if err != nil {
		t.Fatal(err)
	}
	if err := model.UnmarshalJSON(payload); err != nil {
		t.Fatal(err)
	}
}

func metadataValue[T any](t *testing.T, metadata map[string]json.RawMessage, key string) T {
	t.Helper()
	var value T
	raw, ok := metadata[key]
	if !ok {
		t.Fatalf("%q is missing from %s", key, metadata)
	}
	if err := json.Unmarshal(raw, &value); err != nil {
		t.Fatalf("%q holds %s, which is not a %T: %v", key, raw, value, err)
	}
	return value
}

func TestPublicMetadataReadsRecordedPayloads(t *testing.T) {
	t.Run("event data", func(t *testing.T) {
		var event xtwitterscraper.EventDetail
		decodeFixture(t, "getEvent.json", &event)
		post, err := event.Data.AsMonitorPost()
		if err != nil || post.JSON.ExtraFields["tweetId"].Raw() != `"1234567890"` {
			t.Fatalf("post = %+v, err = %v", post, err)
		}
	})
	t.Run("article metadata", func(t *testing.T) {
		var res xtwitterscraper.XGetArticleResponse
		decodeFixture(t, "getArticle.json", &res)
		if got := metadataValue[int64](t, res.Article.Metadata, "first_published_at_secs"); got != 1736942400 {
			t.Fatalf("first_published_at_secs = %d", got)
		}
		if got := metadataValue[int64](t, res.Article.LifecycleState, "modified_at_secs"); got != 1736946000 {
			t.Fatalf("modified_at_secs = %d", got)
		}
	})
	t.Run("write payload", func(t *testing.T) {
		var res xtwitterscraper.XDmDeleteMessageResponse
		decodeFixture(t, "deleteDm.json", &res)
		if got := metadataValue[string](t, res.Request.Payload, "tweet_id"); got != "9876543210" {
			t.Fatalf("tweet_id = %q", got)
		}
	})
	t.Run("analysis question totals", func(t *testing.T) {
		var res xtwitterscraper.XAnalysisSentimentResponse
		decodeFixture(t, "TweetAnalysisPage.json", &res)
		sentiment := res.AnalysisSummary.Questions["sentiment"]
		if got := metadataValue[string](t, sentiment, "type"); got != "choice" {
			t.Fatalf("type = %q", got)
		}
		counts := metadataValue[map[string]int64](t, sentiment, "counts")
		want := map[string]int64{"positive": 0, "negative": 0, "mixed": 1, "neutral": 0, "unclear": 0}
		if !reflect.DeepEqual(counts, want) {
			t.Fatalf("counts = %v, want %v", counts, want)
		}
		monitor := res.Results[0].Monitor
		if monitor.BaselineDatasetID != nil || monitor.JSON.BaselineDatasetID.Raw() != "null" {
			t.Fatalf("baselineDatasetId = %s, metadata %q, want null", monitor.BaselineDatasetID, monitor.JSON.BaselineDatasetID.Raw())
		}
	})
}

func TestWriteActionReadsTheContractMediaAndDetails(t *testing.T) {
	var res xtwitterscraper.XWriteActionGetResponse
	payload := `{"media":{"count":1,"kind":"image"},"details":{"suggestion":"Poll the action status URL.","retry":{"after":30}}}`
	if err := res.UnmarshalJSON([]byte(payload)); err != nil {
		t.Fatal(err)
	}
	if got := metadataValue[int64](t, res.Media, "count"); got != 1 {
		t.Fatalf("media count = %d", got)
	}
	if got := metadataValue[string](t, res.Media, "kind"); got != "image" {
		t.Fatalf("media kind = %q", got)
	}
	if got := metadataValue[string](t, res.Details, "suggestion"); got != "Poll the action status URL." {
		t.Fatalf("suggestion = %q", got)
	}
	if got := string(res.Details["retry"]); got != `{"after":30}` {
		t.Fatalf("retry = %s, want the nested object as sent", got)
	}
}

func TestUnknownModelFieldsKeepTheirJSON(t *testing.T) {
	var metadata xtwitterscraper.RadarItemMetadata
	payload := `{"author":"founder","mrr":12500,"customers":{"active":42}}`
	if err := metadata.UnmarshalJSON([]byte(payload)); err != nil {
		t.Fatal(err)
	}
	if metadata.Author != "founder" || !metadata.JSON.ExtraFields["mrr"].Valid() {
		t.Fatalf("decoded %s", metadata.RawJSON())
	}
	if got := metadataValue[int64](t, metadata.ExtraFields, "mrr"); got != 12500 {
		t.Fatalf("mrr = %d", got)
	}
	if got := metadataValue[map[string]int64](t, metadata.ExtraFields, "customers"); got["active"] != 42 {
		t.Fatalf("customers = %v", got)
	}
	if _, typed := metadata.ExtraFields["author"]; typed {
		t.Fatal("a typed field also landed in ExtraFields")
	}
}

func TestComposeParamsSendTheirExtraFields(t *testing.T) {
	params := xtwitterscraper.ComposeNewParamsBodyComposePrepareRequest{
		Topic:       "launch",
		ExtraFields: map[string]json.RawMessage{"tone": json.RawMessage(`{"style": "dry"}`)},
	}
	body, err := params.MarshalJSON()
	if err != nil {
		t.Fatal(err)
	}
	got, err := canonicalJSON(body)
	if err != nil {
		t.Fatalf("body %s is not JSON: %v", body, err)
	}
	if want := `{"step":"compose","tone":{"style":"dry"},"topic":"launch"}`; got != want {
		t.Fatalf("body = %s, want %s", got, want)
	}
}
