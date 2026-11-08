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

type preferCall func(client xtwitterscraper.Client, prefer param.Opt[string]) error

func TestEveryWriteSendsPreferOnlyWhenSet(t *testing.T) {
	ctx := context.Background()
	const account, key = "@elonmusk", "write-1"
	cases := map[string]struct {
		method string
		path   string
		call   preferCall
	}{
		"X.Communities.New": {http.MethodPost, "/api/v1/x/communities", func(client xtwitterscraper.Client, prefer param.Opt[string]) error {
			_, err := client.X.Communities.New(ctx, xtwitterscraper.XCommunityNewParams{Account: account, Name: "Rustaceans", IdempotencyKey: key, Prefer: prefer})
			return err
		}},
		"X.Communities.Delete": {http.MethodDelete, "/api/v1/x/communities/1500000000000000000", func(client xtwitterscraper.Client, prefer param.Opt[string]) error {
			_, err := client.X.Communities.Delete(ctx, "1500000000000000000", xtwitterscraper.XCommunityDeleteParams{Account: account, CommunityName: "Rustaceans", IdempotencyKey: key, Prefer: prefer})
			return err
		}},
		"X.Communities.Join.New": {http.MethodPost, "/api/v1/x/communities/1500000000000000000/join", func(client xtwitterscraper.Client, prefer param.Opt[string]) error {
			_, err := client.X.Communities.Join.New(ctx, "1500000000000000000", xtwitterscraper.XCommunityJoinNewParams{Account: account, IdempotencyKey: key, Prefer: prefer})
			return err
		}},
		"X.Communities.Join.DeleteAll": {http.MethodDelete, "/api/v1/x/communities/1500000000000000000/join", func(client xtwitterscraper.Client, prefer param.Opt[string]) error {
			_, err := client.X.Communities.Join.DeleteAll(ctx, "1500000000000000000", xtwitterscraper.XCommunityJoinDeleteAllParams{Account: account, IdempotencyKey: key, Prefer: prefer})
			return err
		}},
		"X.Dm.Send": {http.MethodPost, "/api/v1/x/dm/44196397", func(client xtwitterscraper.Client, prefer param.Opt[string]) error {
			_, err := client.X.Dm.Send(ctx, "44196397", xtwitterscraper.XDmSendParams{Account: account, Text: "Hello", IdempotencyKey: key, Prefer: prefer})
			return err
		}},
		"X.Dm.DeleteMessage": {http.MethodDelete, "/api/v1/x/dm/44196397/messages/1600000000000000000", func(client xtwitterscraper.Client, prefer param.Opt[string]) error {
			_, err := client.X.Dm.DeleteMessage(ctx, "1600000000000000000", xtwitterscraper.XDmDeleteMessageParams{UserID: "44196397", Account: account, IdempotencyKey: key, Prefer: prefer})
			return err
		}},
		"X.Tweets.New": {http.MethodPost, "/api/v1/x/tweets", func(client xtwitterscraper.Client, prefer param.Opt[string]) error {
			_, err := client.X.Tweets.New(ctx, xtwitterscraper.XTweetNewParams{Account: account, IdempotencyKey: key, Prefer: prefer})
			return err
		}},
		"X.Tweets.Delete": {http.MethodDelete, "/api/v1/x/tweets/1234567890", func(client xtwitterscraper.Client, prefer param.Opt[string]) error {
			_, err := client.X.Tweets.Delete(ctx, "1234567890", xtwitterscraper.XTweetDeleteParams{Account: account, IdempotencyKey: key, Prefer: prefer})
			return err
		}},
		"X.Tweets.Like.New": {http.MethodPost, "/api/v1/x/tweets/1234567890/like", func(client xtwitterscraper.Client, prefer param.Opt[string]) error {
			_, err := client.X.Tweets.Like.New(ctx, "1234567890", xtwitterscraper.XTweetLikeNewParams{Account: account, IdempotencyKey: key, Prefer: prefer})
			return err
		}},
		"X.Tweets.Like.Delete": {http.MethodDelete, "/api/v1/x/tweets/1234567890/like", func(client xtwitterscraper.Client, prefer param.Opt[string]) error {
			_, err := client.X.Tweets.Like.Delete(ctx, "1234567890", xtwitterscraper.XTweetLikeDeleteParams{Account: account, IdempotencyKey: key, Prefer: prefer})
			return err
		}},
		"X.Tweets.Retweet.New": {http.MethodPost, "/api/v1/x/tweets/1234567890/retweet", func(client xtwitterscraper.Client, prefer param.Opt[string]) error {
			_, err := client.X.Tweets.Retweet.New(ctx, "1234567890", xtwitterscraper.XTweetRetweetNewParams{Account: account, IdempotencyKey: key, Prefer: prefer})
			return err
		}},
		"X.Tweets.Retweet.Delete": {http.MethodDelete, "/api/v1/x/tweets/1234567890/retweet", func(client xtwitterscraper.Client, prefer param.Opt[string]) error {
			_, err := client.X.Tweets.Retweet.Delete(ctx, "1234567890", xtwitterscraper.XTweetRetweetDeleteParams{Account: account, IdempotencyKey: key, Prefer: prefer})
			return err
		}},
		"X.Users.Follow.New": {http.MethodPost, "/api/v1/x/users/44196397/follow", func(client xtwitterscraper.Client, prefer param.Opt[string]) error {
			_, err := client.X.Users.Follow.New(ctx, "44196397", xtwitterscraper.XUserFollowNewParams{Account: account, IdempotencyKey: key, Prefer: prefer})
			return err
		}},
		"X.Users.Follow.DeleteAll": {http.MethodDelete, "/api/v1/x/users/44196397/follow", func(client xtwitterscraper.Client, prefer param.Opt[string]) error {
			_, err := client.X.Users.Follow.DeleteAll(ctx, "44196397", xtwitterscraper.XUserFollowDeleteAllParams{Account: account, IdempotencyKey: key, Prefer: prefer})
			return err
		}},
		"X.Users.RemoveFollower": {http.MethodPost, "/api/v1/x/users/44196397/remove-follower", func(client xtwitterscraper.Client, prefer param.Opt[string]) error {
			_, err := client.X.Users.RemoveFollower(ctx, "44196397", xtwitterscraper.XUserRemoveFollowerParams{Account: account, IdempotencyKey: key, Prefer: prefer})
			return err
		}},
		"X.Profile.Update": {http.MethodPatch, "/api/v1/x/profile", func(client xtwitterscraper.Client, prefer param.Opt[string]) error {
			_, err := client.X.Profile.Update(ctx, xtwitterscraper.XProfileUpdateParams{Account: account, IdempotencyKey: key, Prefer: prefer})
			return err
		}},
		"X.Profile.UpdateAvatar": {http.MethodPatch, "/api/v1/x/profile/avatar", func(client xtwitterscraper.Client, prefer param.Opt[string]) error {
			_, err := client.X.Profile.UpdateAvatar(ctx, xtwitterscraper.XProfileUpdateAvatarParams{Account: account, URL: xtwitterscraper.String("https://example.com/avatar.jpg"), IdempotencyKey: key, Prefer: prefer})
			return err
		}},
		"X.Profile.UpdateBanner": {http.MethodPatch, "/api/v1/x/profile/banner", func(client xtwitterscraper.Client, prefer param.Opt[string]) error {
			_, err := client.X.Profile.UpdateBanner(ctx, xtwitterscraper.XProfileUpdateBannerParams{Account: account, URL: xtwitterscraper.String("https://example.com/banner.jpg"), IdempotencyKey: key, Prefer: prefer})
			return err
		}},
		"X.Media.Upload": {http.MethodPost, "/api/v1/x/media", func(client xtwitterscraper.Client, prefer param.Opt[string]) error {
			_, err := client.X.Media.Upload(ctx, xtwitterscraper.XMediaUploadParams{Account: account, URL: xtwitterscraper.String("https://example.com/photo.jpg"), IdempotencyKey: key, Prefer: prefer})
			return err
		}},
	}
	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			client, request := fixtureClient(t, "muteUser.json")
			if err := test.call(client, xtwitterscraper.String("wait=50")); err != nil {
				t.Fatal(err)
			}
			if request.method != test.method || request.path != test.path {
				t.Fatalf("request = %s %s, want %s %s", request.method, request.path, test.method, test.path)
			}
			if request.header.Get("Prefer") != "wait=50" || request.header.Get("Idempotency-Key") != key {
				t.Fatalf("headers = %v", request.header)
			}
			if err := test.call(client, param.Opt[string]{}); err != nil {
				t.Fatal(err)
			}
			if prefer, sent := request.header["Prefer"]; sent {
				t.Fatalf("an omitted Prefer sent %v", prefer)
			}
		})
	}
}

func TestDrawRunSendsIdempotencyKeyOnlyWhenSet(t *testing.T) {
	client, request := fixtureClient(t, "muteUser.json")
	params := xtwitterscraper.DrawRunParams{TweetURL: "https://x.com/elonmusk/status/1234567890", IdempotencyKey: xtwitterscraper.String("draw-1")}
	if _, err := client.Draws.Run(context.Background(), params); err != nil {
		t.Fatal(err)
	}
	request.expect(t, http.MethodPost, "/api/v1/draws", "")
	request.expectBody(t, `{"tweetUrl":"https://x.com/elonmusk/status/1234567890"}`)
	if request.header.Get("Idempotency-Key") != "draw-1" {
		t.Fatalf("headers = %v", request.header)
	}
	params.IdempotencyKey = param.Opt[string]{}
	if _, err := client.Draws.Run(context.Background(), params); err != nil {
		t.Fatal(err)
	}
	if key, sent := request.header["Idempotency-Key"]; sent {
		t.Fatalf("an omitted Idempotency-Key sent %v", key)
	}
}
