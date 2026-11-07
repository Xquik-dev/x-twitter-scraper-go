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
)

func TestXUserMuteNew(t *testing.T) {
	expectWrite(t, "muteUser.json", http.MethodPost, "/api/v1/x/users/44196397/mute", "mute", func(client xtwitterscraper.Client, prefer param.Opt[string]) (*xtwitterscraper.XWriteActionGetResponse, error) {
		return client.X.Users.Mute.New(context.Background(), "44196397", xtwitterscraper.XUserMuteNewParams{
			Account:        "@elonmusk",
			IdempotencyKey: "write-1",
			Prefer:         prefer,
		})
	})
}

func TestXUserMuteDelete(t *testing.T) {
	expectWrite(t, "unmuteUser.json", http.MethodDelete, "/api/v1/x/users/44196397/mute", "unmute", func(client xtwitterscraper.Client, prefer param.Opt[string]) (*xtwitterscraper.XWriteActionGetResponse, error) {
		return client.X.Users.Mute.Delete(context.Background(), "44196397", xtwitterscraper.XUserMuteDeleteParams{
			Account:        "@elonmusk",
			IdempotencyKey: "write-1",
			Prefer:         prefer,
		})
	})
}
