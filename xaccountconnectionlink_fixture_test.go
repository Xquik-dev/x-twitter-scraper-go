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
)

func TestConnectionLinkNewSendsTheFieldsAndReturnsTheURL(t *testing.T) {
	client, request := fixtureClient(t, "createXAccountConnectionLink.json")
	res, err := client.X.AccountConnectionLinks.New(context.Background(), xtwitterscraper.XAccountConnectionLinkNewParams{
		AccountID: xtwitterscraper.String("42"),
		Reference: xtwitterscraper.String("crm-user-7"),
		ReturnURL: xtwitterscraper.String("https://app.example.com/connected"),
	})
	if err != nil {
		t.Fatal(err)
	}
	request.expect(t, http.MethodPost, "/api/v1/x/account-connection-links", "")
	request.expectBody(t, `{"accountId":"42","reference":"crm-user-7","returnUrl":"https://app.example.com/connected"}`)
	expectDecoded(t, res)
	if res.ID != "xlink_AAAAAAAAAAAAAAAA" || res.URL != "https://xquik.com/connect-x#token=CONNECT_LINK_TOKEN" ||
		res.Purpose != xtwitterscraper.XAccountConnectionLinkPurposeConnect ||
		res.Status != xtwitterscraper.XAccountConnectionLinkStatusPending ||
		res.Reference != "crm-user-7" || res.ReturnURL != "https://app.example.com/connected" ||
		string(res.Object) != "x_account_connection_link" {
		t.Fatalf("decoded %s", res.RawJSON())
	}
	if res.JSON.AccountID.Valid() || res.JSON.CompletedAt.Valid() || res.JSON.Failure.Valid() {
		t.Fatalf("a null field read as set in %s", res.RawJSON())
	}
	if !res.ExpiresAt.Equal(time.Date(2026, 11, 8, 12, 30, 0, 0, time.UTC)) {
		t.Fatalf("expiresAt = %s", res.ExpiresAt)
	}
}

func TestConnectionLinkNewTakesNoFields(t *testing.T) {
	client, request := fixtureClient(t, "createXAccountConnectionLink.json")
	if _, err := client.X.AccountConnectionLinks.New(context.Background(), xtwitterscraper.XAccountConnectionLinkNewParams{}); err != nil {
		t.Fatal(err)
	}
	request.expectBody(t, `{}`)
}

func TestConnectionLinkGetReadsTheConnectedAccountWithNoURL(t *testing.T) {
	client, request := fixtureClient(t, "getXAccountConnectionLink.json")
	res, err := client.X.AccountConnectionLinks.Get(context.Background(), "xlink_AAAAAAAAAAAAAAAA")
	if err != nil {
		t.Fatal(err)
	}
	request.expect(t, http.MethodGet, "/api/v1/x/account-connection-links/xlink_AAAAAAAAAAAAAAAA", "")
	expectDecoded(t, res)
	if res.Status != xtwitterscraper.XAccountConnectionLinkStatusConnected || res.AccountID != "42" ||
		res.XUserID != "1234567890" || res.XUsername != "your_x_username" || res.Health != "healthy" ||
		res.CompletedAt.IsZero() || res.URL != "" {
		t.Fatalf("decoded %s", res.RawJSON())
	}
}

func TestConnectionLinkDeleteRevokesTheLink(t *testing.T) {
	client, request := fixtureClient(t, "revokeXAccountConnectionLink.json")
	res, err := client.X.AccountConnectionLinks.Delete(context.Background(), "xlink_AAAAAAAAAAAAAAAA")
	if err != nil {
		t.Fatal(err)
	}
	request.expect(t, http.MethodDelete, "/api/v1/x/account-connection-links/xlink_AAAAAAAAAAAAAAAA", "")
	expectDecoded(t, res)
	if !res.Success {
		t.Fatalf("decoded %s", res.RawJSON())
	}
}

func TestConnectionLinkReadsRefuseAnEmptyID(t *testing.T) {
	client, _ := fixtureClient(t, "revokeXAccountConnectionLink.json")
	if _, err := client.X.AccountConnectionLinks.Get(context.Background(), ""); err == nil {
		t.Fatal("Get took an empty id")
	}
	if _, err := client.X.AccountConnectionLinks.Delete(context.Background(), ""); err == nil {
		t.Fatal("Delete took an empty id")
	}
}

func TestAccountNewSendsTheConnectionAttemptID(t *testing.T) {
	client, request := fixtureClient(t, "reauthXAccountPending.json")
	_, err := client.X.Accounts.New(context.Background(), xtwitterscraper.XAccountNewParams{
		ConnectionAttemptID: xtwitterscraper.String("xatt_0123456789abcdef0123456789abcdef"),
		Email:               "account@example.invalid",
		Password:            "<ACCOUNT_PASSWORD>",
		TotpSecret:          "<TOTP_SECRET>",
		Username:            "your_x_username",
	})
	if err != nil {
		t.Fatal(err)
	}
	request.expectBody(t, `{"connection_attempt_id":"xatt_0123456789abcdef0123456789abcdef","email":"account@example.invalid","password":"<ACCOUNT_PASSWORD>","totp_secret":"<TOTP_SECRET>","username":"your_x_username"}`)
}

func TestAccountReauthReturnsTheAccount(t *testing.T) {
	client, request := fixtureClient(t, "reauthXAccount.json")
	res, err := client.X.Accounts.Reauth(context.Background(), "42", xtwitterscraper.XAccountReauthParams{
		ConnectionAttemptID: xtwitterscraper.String("xatt_0123456789abcdef0123456789abcdef"),
		Password:            "<ACCOUNT_PASSWORD>",
	})
	if err != nil {
		t.Fatal(err)
	}
	request.expect(t, http.MethodPost, "/api/v1/x/accounts/42/reauth", "")
	request.expectBody(t, `{"connection_attempt_id":"xatt_0123456789abcdef0123456789abcdef","password":"<ACCOUNT_PASSWORD>"}`)
	account := res.AsXAccountReauthResponseSanitizedXAccount()
	if res.Status != "active" || account.ID != "42" || account.Health != "healthy" || account.XUsername != "your_x_username" {
		t.Fatalf("decoded %s", res.RawJSON())
	}
}

func TestAccountReauthReturnsAPendingAttempt(t *testing.T) {
	client, _ := fixtureClient(t, "reauthXAccountPending.json")
	res, err := client.X.Accounts.Reauth(context.Background(), "42", xtwitterscraper.XAccountReauthParams{Password: "<ACCOUNT_PASSWORD>"})
	if err != nil {
		t.Fatal(err)
	}
	pending := res.AsXAccountReauthResponseXAccountConnectionAttemptPending()
	if res.Status != "pending" || pending.ID != "xatt_0123456789abcdef0123456789abcdef" || pending.PollAfterMs != 3000 {
		t.Fatalf("decoded %s", res.RawJSON())
	}
}

func TestAccountReauthReturnsAnEmailCodeChallenge(t *testing.T) {
	client, _ := fixtureClient(t, "reauthXAccountChallenge.json")
	res, err := client.X.Accounts.Reauth(context.Background(), "42", xtwitterscraper.XAccountReauthParams{Password: "<ACCOUNT_PASSWORD>"})
	if err != nil {
		t.Fatal(err)
	}
	challenge := res.AsXAccountReauthResponseXAccountConnectionChallenge()
	if res.Status != "requires_email_code" || challenge.Username != "your_x_username" || challenge.Message == "" || challenge.ExpiresAt.IsZero() {
		t.Fatalf("decoded %s", res.RawJSON())
	}
}

func TestConnectionAttemptSuccessReturnsTheAccountID(t *testing.T) {
	client, request := fixtureClient(t, "getXAccountConnectionAttemptSuccess.json")
	res, err := client.X.AccountConnectionAttempts.Get(context.Background(), "xatt_0123456789abcdef0123456789abcdef")
	if err != nil {
		t.Fatal(err)
	}
	request.expect(t, http.MethodGet, "/api/v1/x/account-connection-attempts/xatt_0123456789abcdef0123456789abcdef", "")
	if res.AccountID != "42" || res.AsSuccess().AccountID != "42" {
		t.Fatalf("decoded %s", res.RawJSON())
	}
}
