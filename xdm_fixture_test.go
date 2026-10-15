// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

package xtwitterscraper_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/Xquik-dev/x-twitter-scraper-go"
	"github.com/Xquik-dev/x-twitter-scraper-go/option"
)

func TestXDmDeleteMessageFixture(t *testing.T) {
	client, request := fixtureClient(t, "deleteDm.json")
	res, err := client.X.Dm.DeleteMessage(context.Background(), "1234567890123456789", xtwitterscraper.XDmDeleteMessageParams{
		Account:        "@elonmusk",
		UserID:         "nasa",
		IdempotencyKey: "delete-dm-1",
	})
	if err != nil {
		t.Fatal(err)
	}
	request.expect(t, http.MethodDelete, "/api/v1/x/dm/nasa/messages/1234567890123456789", "account=%40elonmusk")
	expectDecoded(t, res)
	if request.header.Get("Idempotency-Key") != "delete-dm-1" ||
		res.ID != "12345" ||
		res.WriteActionID != "12345" ||
		res.Action != "like" ||
		res.Status != "success" ||
		!res.Terminal ||
		res.Retryable ||
		res.SafeToRetry ||
		res.StatusURL != "/api/v1/x/write-actions/12345" ||
		!res.Charged ||
		res.ChargedCredits != "10" ||
		res.TargetID != "9876543210" ||
		!res.SendDispatched ||
		!res.Success ||
		res.Billing.Status != "charged" ||
		!res.Billing.Charged ||
		res.Billing.PlannedCredits != "10" ||
		res.Billing.ChargedCredits != "10" ||
		res.Request.Hash != "0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef" ||
		res.Account.ID != "42" ||
		res.Account.Username != "example" ||
		res.Target.Type != "tweet" ||
		res.Target.ID != "9876543210" ||
		res.Result.Type != "state_change" ||
		res.Result.ID != "9876543210" {
		t.Fatalf("decoded %s", res.RawJSON())
	}
}

func TestXDmDeleteMessageRequiresBothIDs(t *testing.T) {
	client := xtwitterscraper.NewClient(option.WithBaseURL("http://127.0.0.1:0/api/v1/"))
	for _, test := range []struct{ userID, messageID, want string }{
		{"", "1234567890123456789", "missing required userId parameter"},
		{"nasa", "", "missing required messageId parameter"},
	} {
		t.Run(test.want, func(t *testing.T) {
			_, err := client.X.Dm.DeleteMessage(context.Background(), test.messageID, xtwitterscraper.XDmDeleteMessageParams{
				UserID:         test.userID,
				Account:        "@elonmusk",
				IdempotencyKey: "delete-dm-1",
			})
			if err == nil || err.Error() != test.want {
				t.Fatalf("err = %v, want %q", err, test.want)
			}
		})
	}
}
