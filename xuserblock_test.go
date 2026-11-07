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

func TestXUserBlockNew(t *testing.T) {
	expectWrite(t, "blockUser.json", http.MethodPost, "/api/v1/x/users/44196397/block", "block", func(client xtwitterscraper.Client, prefer param.Opt[string]) (*xtwitterscraper.XWriteActionGetResponse, error) {
		return client.X.Users.Block.New(context.Background(), "44196397", xtwitterscraper.XUserBlockNewParams{
			Account:        "@elonmusk",
			IdempotencyKey: "write-1",
			Prefer:         prefer,
		})
	})
}

func TestXUserBlockDelete(t *testing.T) {
	expectWrite(t, "unblockUser.json", http.MethodDelete, "/api/v1/x/users/44196397/block", "unblock", func(client xtwitterscraper.Client, prefer param.Opt[string]) (*xtwitterscraper.XWriteActionGetResponse, error) {
		return client.X.Users.Block.Delete(context.Background(), "44196397", xtwitterscraper.XUserBlockDeleteParams{
			Account:        "@elonmusk",
			IdempotencyKey: "write-1",
			Prefer:         prefer,
		})
	})
}
