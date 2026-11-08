// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

package xtwitterscraper_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/Xquik-dev/x-twitter-scraper-go"
	"github.com/Xquik-dev/x-twitter-scraper-go/packages/param"
	"github.com/Xquik-dev/x-twitter-scraper-go/shared"
)

func TestMonitorUpsertSendsTheEventTypesAndLabelAndSaysItCreated(t *testing.T) {
	client, request := fixtureClient(t, "upsertMonitor.json")
	res, err := client.Monitors.Upsert(context.Background(), "elonmusk", xtwitterscraper.MonitorUpsertParams{
		EventTypes: []shared.EventType{shared.EventTypeTweetNew, shared.EventTypeTweetReply},
		Name:       xtwitterscraper.String("Launch watch"),
	})
	if err != nil {
		t.Fatal(err)
	}
	request.expect(t, http.MethodPut, "/api/v1/monitors/username/elonmusk", "")
	request.expectBody(t, `{"eventTypes":["tweet.new","tweet.reply"],"name":"Launch watch"}`)
	expectDecoded(t, res)
	if !res.Created || res.ID != "42" || res.Username != "elonmusk" || res.XUserID != "1234567890" ||
		res.Name != "Launch watch" || len(res.EventTypes) != 2 || !res.IsActive {
		t.Fatalf("decoded %s", res.RawJSON())
	}
}

func TestMonitorUpsertSendsANullNameToRemoveTheLabel(t *testing.T) {
	client, request := fixtureClient(t, "upsertMonitor.json")
	_, err := client.Monitors.Upsert(context.Background(), "1234567890", xtwitterscraper.MonitorUpsertParams{
		EventTypes: []shared.EventType{shared.EventTypeTweetNew},
		Name:       param.Null[string](),
	})
	if err != nil {
		t.Fatal(err)
	}
	request.expect(t, http.MethodPut, "/api/v1/monitors/username/1234567890", "")
	request.expectBody(t, `{"eventTypes":["tweet.new"],"name":null}`)
}
