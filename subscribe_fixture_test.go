package xtwitterscraper_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/Xquik-dev/x-twitter-scraper-go"
)

func TestSubscribeNewFixture(t *testing.T) {
	client, request := fixtureClient(t, "subscribe.json")
	res, err := client.Subscribe.New(context.Background(), xtwitterscraper.SubscribeNewParams{
		Currency: "EUR",
		Tier:     "pro",
	})
	if err != nil {
		t.Fatal(err)
	}
	request.expect(t, http.MethodPost, "/api/v1/subscribe", "")
	request.expectBody(t, `{"currency":"EUR","tier":"pro"}`)
	expectDecoded(t, res)
	if res.ChargedCurrency != "EUR" || res.Status != "checkout_created" || res.URL != "https://xquik.com/billing/session" || res.Message != "Billing session created" {
		t.Fatalf("decoded %s", res.RawJSON())
	}
}
