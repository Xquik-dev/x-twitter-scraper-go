// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

package xtwitterscraper_test

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/Xquik-dev/x-twitter-scraper-go"
)

func TestMonitorEventDataMarshalsTheDecodedPayload(t *testing.T) {
	var event xtwitterscraper.EventDetail
	decodeFixture(t, "getEvent.json", &event)
	got, err := json.Marshal(event.Data)
	if err != nil {
		t.Fatal(err)
	}
	payload, err := fixtures.ReadFile("testdata/fixtures/getEvent.json")
	if err != nil {
		t.Fatal(err)
	}
	var fixture struct {
		Data json.RawMessage `json:"data"`
	}
	if err := json.Unmarshal(payload, &fixture); err != nil {
		t.Fatal(err)
	}
	var want bytes.Buffer
	if err := json.Compact(&want, fixture.Data); err != nil {
		t.Fatal(err)
	}
	if string(got) != want.String() {
		t.Fatalf("data marshals to %s, want %s", got, want.String())
	}
}

func TestEventDataAsAnyPicksTheVariantOfTheEventType(t *testing.T) {
	var detail xtwitterscraper.EventDetail
	decodeFixture(t, "getEvent.json", &detail)
	variant, err := detail.DataAsAny()
	if post, ok := variant.(xtwitterscraper.MonitorPost); err != nil || !ok || post.JSON.ExtraFields["tweetId"].Raw() != `"1234567890"` {
		t.Fatalf("detail variant = %#v, err = %v", variant, err)
	}

	var event xtwitterscraper.Event
	if err := event.UnmarshalJSON([]byte(`{"id":"43","type":"following.added","monitorId":"7","monitorType":"account","occurredAt":"2025-01-15T12:00:05Z","data":{"id":"9"}}`)); err != nil {
		t.Fatal(err)
	}
	variant, err = event.DataAsAny()
	if _, ok := variant.(xtwitterscraper.MonitorFollow); err != nil || !ok {
		t.Fatalf("event variant = %#v, err = %v", variant, err)
	}
}
