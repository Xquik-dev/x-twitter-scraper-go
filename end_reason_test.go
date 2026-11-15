// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

package xtwitterscraper_test

import (
	"context"
	"testing"

	"github.com/Xquik-dev/x-twitter-scraper-go"
)

var (
	tweetPages = map[string]string{"PaginatedTweets.json": "", "PaginatedTweetsStalled.json": "stalled"}
	userPages  = map[string]string{"PaginatedUsers.json": "", "PaginatedUsersStalled.json": "stalled"}
)

var endedReads = map[string]struct {
	pages map[string]string
	read  func(client xtwitterscraper.Client) (endReason string, decoded any, err error)
}{
	"X.Bookmarks.List": {tweetPages, func(client xtwitterscraper.Client) (string, any, error) {
		res, err := client.X.Bookmarks.List(context.Background(), xtwitterscraper.XBookmarkListParams{})
		if err != nil {
			return "", nil, err
		}
		return res.EndReason, res, nil
	}},
	"X.Tweets.Search": {tweetPages, func(client xtwitterscraper.Client) (string, any, error) {
		res, err := client.X.Tweets.Search(context.Background(), xtwitterscraper.XTweetSearchParams{Q: "rust"})
		if err != nil {
			return "", nil, err
		}
		return res.EndReason, res.AsPaginatedTweets(), nil
	}},
	"X.Communities.GetMembers": {userPages, func(client xtwitterscraper.Client) (string, any, error) {
		res, err := client.X.Communities.GetMembers(context.Background(), "9", xtwitterscraper.XCommunityGetMembersParams{})
		if err != nil {
			return "", nil, err
		}
		return res.EndReason, res, nil
	}},
	"X.Users.GetFollowers": {userPages, func(client xtwitterscraper.Client) (string, any, error) {
		res, err := client.X.Users.GetFollowers(context.Background(), "44", xtwitterscraper.XUserGetFollowersParams{})
		if err != nil {
			return "", nil, err
		}
		return res.EndReason, res.AsPaginatedUsers(), nil
	}},
}

func TestALastPageSaysWhenXCutItShort(t *testing.T) {
	for name, test := range endedReads {
		for fixture, want := range test.pages {
			t.Run(name+"/"+fixture, func(t *testing.T) {
				client, _ := fixtureClient(t, fixture)
				endReason, decoded, err := test.read(client)
				if err != nil {
					t.Fatal(err)
				}
				expectDecoded(t, decoded)
				if endReason != want {
					t.Fatalf("EndReason = %q, want %q", endReason, want)
				}
			})
		}
	}
}
