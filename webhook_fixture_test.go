// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

package xtwitterscraper_test

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/Xquik-dev/x-twitter-scraper-go"
	"github.com/Xquik-dev/x-twitter-scraper-go/shared"
)

func TestWebhookNewTelegramLinkFixture(t *testing.T) {
	client, request := fixtureClient(t, "createTelegramLink.json")
	res, err := client.Webhooks.NewTelegramLink(context.Background(), xtwitterscraper.WebhookNewTelegramLinkParams{
		EventTypes: []shared.EventType{shared.EventTypeTweetNew, shared.EventTypeTweetReply},
	})
	if err != nil {
		t.Fatal(err)
	}
	request.expect(t, http.MethodPost, "/api/v1/webhooks/telegram", "")
	request.expectBody(t, `{"eventTypes": ["tweet.new", "tweet.reply"]}`)
	expectDecoded(t, res)
	if res.URL != "https://t.me/XquikBot?start=Q2hhdExpbmtDb2RlRXhhbXBsZTAx" ||
		!res.ExpiresAt.Equal(time.Date(2025, time.January, 15, 12, 15, 0, 0, time.UTC)) {
		t.Fatalf("decoded %s", res.RawJSON())
	}
}
