package xtwitterscraper_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/Xquik-dev/x-twitter-scraper-go"
)

func TestCreditTopupBalanceFixture(t *testing.T) {
	client, request := fixtureClient(t, "topUpCredits.json")
	res, err := client.Credits.TopupBalance(context.Background(), xtwitterscraper.CreditTopupBalanceParams{
		Currency: "EUR",
		Dollars:  10,
		Locale:   xtwitterscraper.String("en"),
	})
	if err != nil {
		t.Fatal(err)
	}
	request.expect(t, http.MethodPost, "/api/v1/credits/topup", "")
	request.expectBody(t, `{"currency":"EUR","dollars":10,"locale":"en"}`)
	expectDecoded(t, res)
	if res.ChargedCurrency != "EUR" || res.ChargedMinorUnits != 892 ||
		res.URL != "https://xquik.com/api/v1/credits/topup/redirect?session_id=checkout_example" {
		t.Fatalf("decoded %s", res.RawJSON())
	}
}

func TestCreditTopupBalanceLeavesOutUnsetCurrency(t *testing.T) {
	client, request := fixtureClient(t, "topUpCredits.json")
	if _, err := client.Credits.TopupBalance(context.Background(), xtwitterscraper.CreditTopupBalanceParams{Dollars: 10}); err != nil {
		t.Fatal(err)
	}
	request.expectBody(t, `{"dollars":10}`)
}

func TestCreditGetTopupStatusFixture(t *testing.T) {
	client, request := fixtureClient(t, "getCreditTopupStatus.json")
	res, err := client.Credits.GetTopupStatus(context.Background(), xtwitterscraper.CreditGetTopupStatusParams{SessionID: "checkout_example"})
	if err != nil {
		t.Fatal(err)
	}
	request.expect(t, http.MethodGet, "/api/v1/credits/topup/status", "session_id=checkout_example")
	expectDecoded(t, res)
	if res.Status != "paid" || res.AmountDollars != 25 || res.ChargedCurrency != "EUR" || res.ChargedMinorUnits != 2231 || res.Credits != "166666" {
		t.Fatalf("decoded %s", res.RawJSON())
	}
}
