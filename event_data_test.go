// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

package xtwitterscraper_test

import (
	"slices"
	"testing"
	"time"

	"github.com/Xquik-dev/x-twitter-scraper-go"
)

func decodeEventData(t *testing.T, eventType, data string) xtwitterscraper.EventDetail {
	t.Helper()
	var event xtwitterscraper.EventDetail
	body := `{"id":"1","type":"` + eventType + `","monitorId":"10","monitorType":"account","occurredAt":"2025-01-15T12:00:00Z","data":` + data + `}`
	if err := event.UnmarshalJSON([]byte(body)); err != nil {
		t.Fatal(err)
	}
	return event
}

func TestEventDataPicksTheVariantOfTheEventType(t *testing.T) {
	t.Run("post", func(t *testing.T) {
		event := decodeEventData(t, "tweet.new", `{"id":"1234567890","text":"The future is now.","likeCount":42,"author":{"id":"44196397","username":"elonmusk","userName":"elonmusk","verified":true,"rawVerified":false}}`)
		variant, err := event.Data.AsAny(event.Type)
		post, ok := variant.(xtwitterscraper.MonitorPost)
		if err != nil || !ok || post.ID != "1234567890" || post.Text != "The future is now." || post.LikeCount != 42 ||
			post.Author.Username != "elonmusk" || post.Author.JSON.UserName.Raw() != `"elonmusk"` || !post.Author.Verified || post.Author.JSON.RawVerified.Raw() != "false" {
			t.Fatalf("variant = %#v, err = %v", variant, err)
		}
	})
	t.Run("profile change", func(t *testing.T) {
		event := decodeEventData(t, "profile.followers.changed", `{"field":"followers","userId":"44196397","username":"elonmusk","change":-3,"previous":120,"current":117}`)
		variant, err := event.Data.AsAny(event.Type)
		change, ok := variant.(xtwitterscraper.MonitorProfileChange)
		if err != nil || !ok || change.Field != "followers" || change.UserID != "44196397" || change.Change != -3 ||
			change.Previous.OfInt != 120 || change.Current.OfInt != 117 {
			t.Fatalf("variant = %#v, err = %v", variant, err)
		}
	})
	t.Run("profile text and pinned posts", func(t *testing.T) {
		event := decodeEventData(t, "profile.pinned_tweet.changed", `{"field":"pinnedTweetIds","userId":"44196397","username":"elonmusk","previous":["1"],"current":["2","3"]}`)
		change, err := event.Data.AsMonitorProfileChange()
		if err != nil || !slices.Equal(change.Previous.OfStringArray, []string{"1"}) || !slices.Equal(change.Current.OfStringArray, []string{"2", "3"}) {
			t.Fatalf("change = %#v, err = %v", change, err)
		}
		event = decodeEventData(t, "profile.bio.changed", `{"field":"description","userId":"44196397","username":"elonmusk","previous":"old bio","current":"new bio"}`)
		if change, err = event.Data.AsMonitorProfileChange(); err != nil || change.Previous.OfString != "old bio" || change.Current.OfString != "new bio" {
			t.Fatalf("change = %#v, err = %v", change, err)
		}
		event = decodeEventData(t, "profile.verified.changed", `{"field":"verified","userId":"44196397","username":"elonmusk","previous":false,"current":true}`)
		if change, err = event.Data.AsMonitorProfileChange(); err != nil || change.Previous.OfBool || !change.Current.OfBool {
			t.Fatalf("change = %#v, err = %v", change, err)
		}
	})
	t.Run("follow", func(t *testing.T) {
		event := decodeEventData(t, "follower.added", `{"id":"2244994945","followedAt":"2026-10-13T02:36:54.961Z","followedAtExact":true}`)
		variant, err := event.Data.AsAny(event.Type)
		follow, ok := variant.(xtwitterscraper.MonitorFollow)
		if err != nil || !ok || follow.ID != "2244994945" || !follow.FollowedAtExact ||
			!follow.FollowedAt.Equal(time.Date(2026, time.October, 13, 2, 36, 54, 961000000, time.UTC)) {
			t.Fatalf("variant = %#v, err = %v", variant, err)
		}
		event = decodeEventData(t, "following.removed", `{"id":"2244994945","reason":"unfollowed"}`)
		if follow, err = event.Data.AsMonitorFollow(); err != nil || follow.Reason != xtwitterscraper.MonitorFollowReasonUnfollowed {
			t.Fatalf("follow = %#v, err = %v", follow, err)
		}
	})
	t.Run("post metric change", func(t *testing.T) {
		event := decodeEventData(t, "post.likes.changed", `{"metric":"likes","previous":17096,"current":17200,"change":104}`)
		variant, err := event.Data.AsAny(event.Type)
		metric, ok := variant.(xtwitterscraper.MonitorPostMetricChange)
		if err != nil || !ok || metric.Metric != xtwitterscraper.MonitorPostMetricChangeMetricLikes ||
			metric.Previous != 17096 || metric.Current != 17200 || metric.Change != 104 {
			t.Fatalf("variant = %#v, err = %v", variant, err)
		}
	})
	t.Run("post edit", func(t *testing.T) {
		event := decodeEventData(t, "post.edited", `{"previousText":"We cut serving costs by 20%.","text":"We cut serving costs by 25%.","versionId":"2082577277246973001"}`)
		variant, err := event.Data.AsAny(event.Type)
		edit, ok := variant.(xtwitterscraper.MonitorPostEdit)
		if err != nil || !ok || edit.PreviousText != "We cut serving costs by 20%." || edit.Text != "We cut serving costs by 25%." || edit.VersionID != "2082577277246973001" {
			t.Fatalf("variant = %#v, err = %v", variant, err)
		}
	})
	t.Run("post deleted", func(t *testing.T) {
		event := decodeEventData(t, "post.deleted", `{"text":"We cut serving costs by 20%."}`)
		variant, err := event.Data.AsAny(event.Type)
		deleted, ok := variant.(xtwitterscraper.MonitorPostDeleted)
		if err != nil || !ok || deleted.Text != "We cut serving costs by 20%." {
			t.Fatalf("variant = %#v, err = %v", variant, err)
		}
	})
	t.Run("legacy post", func(t *testing.T) {
		event := decodeEventData(t, "tweet.new", `{"id":"1234567890","text":"Just launched!","author":null,"quoted_tweet":{"id":"9"}}`)
		legacy, err := event.Data.AsMonitorLegacyPost()
		if err != nil || legacy.ID != "1234567890" || legacy.Text != "Just launched!" || legacy.Author != nil ||
			string(legacy.QuotedTweet["id"]) != `"9"` {
			t.Fatalf("legacy = %#v, err = %v", legacy, err)
		}
	})
	t.Run("unknown type", func(t *testing.T) {
		event := decodeEventData(t, "account.renamed", `{"id":"1"}`)
		if variant, err := event.Data.AsAny(event.Type); variant != nil || err != nil {
			t.Fatalf("variant = %#v, err = %v", variant, err)
		}
	})
}

func TestEventListDecodesPostEventData(t *testing.T) {
	var res xtwitterscraper.EventListResponse
	body := `{"events":[{"id":"9002","type":"post.views.changed","monitorId":"31","monitorType":"post","occurredAt":"2026-10-22T11:05:00.000Z","data":{"metric":"views","previous":100,"current":250,"change":150}}],"hasMore":false}`
	if err := res.UnmarshalJSON([]byte(body)); err != nil {
		t.Fatal(err)
	}
	metric, err := res.Events[0].Data.AsMonitorPostMetricChange()
	if err != nil || metric.Metric != xtwitterscraper.MonitorPostMetricChangeMetricViews || metric.Change != 150 || res.Events[0].Data.Change != 150 {
		t.Fatalf("metric = %#v, err = %v", metric, err)
	}
}
