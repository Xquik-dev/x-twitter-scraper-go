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

func TestXTweetBookmarkNew(t *testing.T) {
	expectWrite(t, "bookmarkTweet.json", http.MethodPost, "/api/v1/x/tweets/1234567890/bookmark", "bookmark", func(client xtwitterscraper.Client, prefer param.Opt[string]) (*xtwitterscraper.XWriteActionGetResponse, error) {
		return client.X.Tweets.Bookmark.New(context.Background(), "1234567890", xtwitterscraper.XTweetBookmarkNewParams{
			Account:        "@elonmusk",
			IdempotencyKey: "write-1",
			Prefer:         prefer,
		})
	})
}

func TestXTweetBookmarkDelete(t *testing.T) {
	expectWrite(t, "unbookmarkTweet.json", http.MethodDelete, "/api/v1/x/tweets/1234567890/bookmark", "unbookmark", func(client xtwitterscraper.Client, prefer param.Opt[string]) (*xtwitterscraper.XWriteActionGetResponse, error) {
		return client.X.Tweets.Bookmark.Delete(context.Background(), "1234567890", xtwitterscraper.XTweetBookmarkDeleteParams{
			Account:        "@elonmusk",
			IdempotencyKey: "write-1",
			Prefer:         prefer,
		})
	})
}
