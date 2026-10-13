// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

package xtwitterscraper_test

import (
	"net/url"
	"testing"

	"github.com/Xquik-dev/x-twitter-scraper-go"
)

// Code written for 0.18 sets MinFaves on these 9 params types. It must still
// compile and still send the minFaves query key, which the API reads as
// minLikes.
func TestMinFavesFromVersion018StillSendsMinFaves(t *testing.T) {
	cases := map[string]interface{ URLQuery() (url.Values, error) }{
		"XTweetGetQuotesParams":  xtwitterscraper.XTweetGetQuotesParams{MinFaves: xtwitterscraper.Int(100)},
		"XTweetGetRepliesParams": xtwitterscraper.XTweetGetRepliesParams{MinFaves: xtwitterscraper.Int(100)},
		"XTweetGetThreadParams":  xtwitterscraper.XTweetGetThreadParams{MinFaves: xtwitterscraper.Int(100)},
		"XTweetSearchParams": xtwitterscraper.XTweetSearchParams{
			Q:        "from:elonmusk",
			MinFaves: xtwitterscraper.Int(100),
		},
		"XUserGetLikesParams":    xtwitterscraper.XUserGetLikesParams{MinFaves: xtwitterscraper.Int(100)},
		"XUserGetMediaParams":    xtwitterscraper.XUserGetMediaParams{MinFaves: xtwitterscraper.Int(100)},
		"XUserGetMentionsParams": xtwitterscraper.XUserGetMentionsParams{MinFaves: xtwitterscraper.Int(100)},
		"XUserGetRepliesParams":  xtwitterscraper.XUserGetRepliesParams{MinFaves: xtwitterscraper.Int(100)},
		"XUserGetTweetsParams":   xtwitterscraper.XUserGetTweetsParams{MinFaves: xtwitterscraper.Int(100)},
	}
	for name, params := range cases {
		t.Run(name, func(t *testing.T) {
			query, err := params.URLQuery()
			if err != nil {
				t.Fatal(err)
			}
			if got := query.Get("minFaves"); got != "100" {
				t.Errorf("minFaves is %q, want 100", got)
			}
			if query.Has("minLikes") {
				t.Errorf("minLikes is set to %q, want it unset", query.Get("minLikes"))
			}
		})
	}
}
