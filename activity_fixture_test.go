// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

package xtwitterscraper_test

import (
	"context"
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/Xquik-dev/x-twitter-scraper-go"
)

func TestActivityListFixture(t *testing.T) {
	client, request := fixtureClient(t, "listActivity.json")
	res, err := client.Activity.List(context.Background(), xtwitterscraper.ActivityListParams{
		Limit:    xtwitterscraper.Int(50),
		Category: xtwitterscraper.String("monitors"),
		Status:   xtwitterscraper.String("failed,running"),
		Q:        xtwitterscraper.String("artemis"),
		Since:    xtwitterscraper.Time(time.Date(2026, time.October, 21, 0, 0, 0, 0, time.UTC)),
		Until:    xtwitterscraper.Time(time.Date(2026, time.October, 28, 0, 0, 0, 0, time.UTC)),
	})
	if err != nil {
		t.Fatal(err)
	}
	request.expect(t, http.MethodGet, "/api/v1/activity", "category=monitors&limit=50&q=artemis&since=2026-10-21T00%3A00%3A00Z&status=failed%2Crunning&until=2026-10-28T00%3A00%3A00Z")
	expectDecoded(t, res)
	rows := *res
	if len(rows) != 2 {
		t.Fatalf("decoded %d rows, want 2", len(rows))
	}
	counted, empty := rows[0], rows[1]
	if counted.Time != "2026-10-28T20:22:13.481Z" ||
		counted.Category != "monitor_change" ||
		counted.Type != "pause" ||
		counted.Activity != "Account monitor for @elonmusk paused" ||
		counted.Status != "completed" ||
		counted.Source != "elonmusk" ||
		counted.Target != "1846000000000000001" ||
		counted.Details != "200" ||
		counted.ID != "naF3xqHLWDTs" ||
		counted.Link != "https://x.com/elonmusk/status/1846000000000000001" ||
		empty.Target != "" ||
		empty.Link != "" {
		t.Fatalf("decoded %s and %s", counted.RawJSON(), empty.RawJSON())
	}
	if counted.Results.OfInt != 12 || counted.Results.AsInt() != 12 || !counted.Results.JSON.OfInt.Valid() || counted.Results.JSON.OfString.Valid() ||
		counted.Credits.OfInt != 1000 || counted.Credits.AsInt() != 1000 || !counted.Credits.JSON.OfInt.Valid() || counted.Credits.JSON.OfString.Valid() {
		t.Fatalf("decoded the counted row's numbers as %s and %s", counted.Results.RawJSON(), counted.Credits.RawJSON())
	}
	if empty.Results.AsString() != "" || !empty.Results.JSON.OfString.Valid() || empty.Results.JSON.OfInt.Valid() || empty.Results.RawJSON() != `""` ||
		empty.Credits.AsString() != "" || !empty.Credits.JSON.OfString.Valid() || empty.Credits.JSON.OfInt.Valid() || empty.Credits.RawJSON() != `""` {
		t.Fatalf("decoded the empty row's numbers as %s and %s", empty.Results.RawJSON(), empty.Credits.RawJSON())
	}
}

func TestActivityExportFixture(t *testing.T) {
	client, request := fixtureClient(t, "exportActivity.csv")
	res, err := client.Activity.Export(context.Background(), xtwitterscraper.ActivityExportParams{
		Format: xtwitterscraper.ActivityExportParamsFormatCsv,
		Limit:  xtwitterscraper.Int(50),
		Q:      xtwitterscraper.String("artemis"),
	})
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}
	request.expect(t, http.MethodGet, "/api/v1/activity", "format=csv&limit=50&q=artemis")
	if accept := request.header.Get("Accept"); accept != "application/octet-stream" {
		t.Fatalf("Accept = %q, want application/octet-stream", accept)
	}
	want := "time,category,type,activity,status,source,target,results,credits,details,id,link\n2026-10-28T20:22:13.481Z,monitor_change,pause,Account monitor for @elonmusk paused,completed,elonmusk,,,,,naF3xqHLWDTs,\n"
	if string(body) != want {
		t.Fatalf("body = %q, want %q", body, want)
	}
}

func TestActivityGetFixture(t *testing.T) {
	client, request := fixtureClient(t, "getActivity.json")
	res, err := client.Activity.Get(context.Background(), "naF3xqHLWDTs")
	if err != nil {
		t.Fatal(err)
	}
	request.expect(t, http.MethodGet, "/api/v1/activity/naF3xqHLWDTs", "")
	expectDecoded(t, res)
	if len(res.Related) != 1 {
		t.Fatalf("decoded %d related activities, want 1", len(res.Related))
	}
	related := res.Related[0]
	if res.Object != "activity" ||
		res.ID != "naF3xqHLWDTs" ||
		res.Kind != xtwitterscraper.ActivityRecordKindMonitorChange ||
		res.Type != "pause" ||
		res.Status != xtwitterscraper.ActivityRecordStatusCompleted ||
		res.URL != "https://dashboard.xquik.com/activity/naF3xqHLWDTs" ||
		res.RequestID != "9f2c1d7e4b8a4c55a1e0d3b2c6f7a8e9" ||
		!res.Timestamps.CreatedAt.Equal(time.Date(2026, time.October, 28, 20, 22, 13, 481000000, time.UTC)) ||
		res.Timestamps.JSON.UpdatedAt.Valid() ||
		res.Request.Method != xtwitterscraper.ActivityRecordRequestMethodPatch ||
		res.Request.Path != "/api/v1/monitors/42" ||
		res.Request.HTTPStatus != 200 ||
		res.Request.DurationMs != 37 ||
		res.Request.IPAddress != "203.0.113.9" ||
		res.Request.UserAgent != "my-agent/1.0" ||
		!res.Request.At.Equal(time.Date(2026, time.October, 28, 20, 22, 13, 519000000, time.UTC)) ||
		res.Caller.AuthMethod != xtwitterscraper.ActivityRecordCallerAuthMethodAPIKey ||
		res.Caller.APIKey.Name != "Production" ||
		res.Caller.APIKey.Prefix != "xq_ab12" ||
		res.Caller.JSON.OAuthClient.Raw() != "null" ||
		res.JSON.Charge.Raw() != "null" ||
		res.Details["monitorKind"] != "account" ||
		res.Details["subjectXId"] != "44196397" ||
		related.ID != "Q7wErTyUiOpA" ||
		related.Type != "create" ||
		related.Status != xtwitterscraper.ActivityRecordRelatedStatusCompleted ||
		!related.CreatedAt.Equal(time.Date(2026, time.October, 21, 9, 12, 40, 2000000, time.UTC)) ||
		related.URL != "https://dashboard.xquik.com/activity/Q7wErTyUiOpA" {
		t.Fatalf("decoded %s", res.RawJSON())
	}
}
