// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

package xtwitterscraper_test

import (
	"bytes"
	"encoding/json"
	"reflect"
	"slices"
	"testing"
	"time"

	"github.com/Xquik-dev/x-twitter-scraper-go"
)

func TestMonitorEventDataMarshalsTheDecodedPayload(t *testing.T) {
	var event xtwitterscraper.EventDetail
	decodeFixture(t, "getEvent.json", &event)
	got, err := json.Marshal(event.Data)
	if err != nil {
		t.Fatal(err)
	}
	payload, err := fixtures.ReadFile("testdata/fixtures/getEvent.json")
	if err != nil {
		t.Fatal(err)
	}
	var fixture struct {
		Data json.RawMessage `json:"data"`
	}
	if err := json.Unmarshal(payload, &fixture); err != nil {
		t.Fatal(err)
	}
	var want bytes.Buffer
	if err := json.Compact(&want, fixture.Data); err != nil {
		t.Fatal(err)
	}
	if string(got) != want.String() {
		t.Fatalf("data marshals to %s, want %s", got, want.String())
	}
}

func decodeEventData(t *testing.T, eventType, data string) (xtwitterscraper.Event, xtwitterscraper.EventDetail) {
	t.Helper()
	body := []byte(`{"id":"1","type":"` + eventType + `","monitorId":"10","monitorType":"account","occurredAt":"2025-01-15T12:00:00Z","data":` + data + `}`)
	var event xtwitterscraper.Event
	if err := event.UnmarshalJSON(body); err != nil {
		t.Fatal(err)
	}
	var detail xtwitterscraper.EventDetail
	if err := detail.UnmarshalJSON(body); err != nil {
		t.Fatal(err)
	}
	return event, detail
}

func eventDataVariant[T xtwitterscraper.MonitorEventDataVariant](t *testing.T, eventType, data string) T {
	t.Helper()
	event, detail := decodeEventData(t, eventType, data)
	fromEvent, eventErr := event.DataAsAny()
	fromDetail, detailErr := detail.DataAsAny()
	variant, ok := fromDetail.(T)
	if eventErr != nil || detailErr != nil || !ok || !reflect.DeepEqual(fromEvent, fromDetail) {
		t.Fatalf("event variant = %#v, err = %v, detail variant = %#v, err = %v", fromEvent, eventErr, fromDetail, detailErr)
	}
	return variant
}

func TestEventDataAsAnyPicksTheVariantOfTheEventType(t *testing.T) {
	t.Run("post", func(t *testing.T) {
		post := eventDataVariant[xtwitterscraper.MonitorPost](t, "tweet.new", `{"id":"1234567890","text":"The future is now.","likeCount":42,"author":{"id":"44196397","username":"elonmusk","userName":"elonmusk","verified":true,"rawVerified":false}}`)
		if post.ID != "1234567890" || post.Text != "The future is now." || post.LikeCount != 42 ||
			post.Author.Username != "elonmusk" || post.Author.JSON.UserName.Raw() != `"elonmusk"` || !post.Author.Verified || post.Author.JSON.RawVerified.Raw() != "false" {
			t.Fatalf("post = %#v", post)
		}
	})
	t.Run("fixture post", func(t *testing.T) {
		var detail xtwitterscraper.EventDetail
		decodeFixture(t, "getEvent.json", &detail)
		variant, err := detail.DataAsAny()
		if post, ok := variant.(xtwitterscraper.MonitorPost); err != nil || !ok || post.JSON.ExtraFields["tweetId"].Raw() != `"1234567890"` {
			t.Fatalf("variant = %#v, err = %v", variant, err)
		}
	})
	t.Run("profile change", func(t *testing.T) {
		change := eventDataVariant[xtwitterscraper.MonitorProfileChange](t, "profile.followers.changed", `{"field":"followers","userId":"44196397","username":"elonmusk","change":-3,"previous":120,"current":117}`)
		if change.Field != "followers" || change.UserID != "44196397" || change.Change != -3 ||
			change.Previous.OfInt != 120 || change.Current.OfInt != 117 {
			t.Fatalf("change = %#v", change)
		}
	})
	t.Run("profile text and pinned posts", func(t *testing.T) {
		change := eventDataVariant[xtwitterscraper.MonitorProfileChange](t, "profile.pinned_tweet.changed", `{"field":"pinnedTweetIds","userId":"44196397","username":"elonmusk","previous":["1"],"current":["2","3"]}`)
		if !slices.Equal(change.Previous.OfStringArray, []string{"1"}) || !slices.Equal(change.Current.OfStringArray, []string{"2", "3"}) {
			t.Fatalf("change = %#v", change)
		}
		change = eventDataVariant[xtwitterscraper.MonitorProfileChange](t, "profile.bio.changed", `{"field":"description","userId":"44196397","username":"elonmusk","previous":"old bio","current":"new bio"}`)
		if change.Previous.OfString != "old bio" || change.Current.OfString != "new bio" {
			t.Fatalf("change = %#v", change)
		}
		change = eventDataVariant[xtwitterscraper.MonitorProfileChange](t, "profile.verified.changed", `{"field":"verified","userId":"44196397","username":"elonmusk","previous":false,"current":true}`)
		if change.Previous.OfBool || !change.Current.OfBool {
			t.Fatalf("change = %#v", change)
		}
	})
	t.Run("follow", func(t *testing.T) {
		follow := eventDataVariant[xtwitterscraper.MonitorFollow](t, "follower.added", `{"id":"2244994945","followedAt":"2026-10-13T02:36:54.961Z","followedAtExact":true}`)
		if follow.ID != "2244994945" || !follow.FollowedAtExact ||
			!follow.FollowedAt.Equal(time.Date(2026, time.October, 13, 2, 36, 54, 961000000, time.UTC)) {
			t.Fatalf("follow = %#v", follow)
		}
		follow = eventDataVariant[xtwitterscraper.MonitorFollow](t, "following.removed", `{"id":"2244994945","reason":"unfollowed"}`)
		if follow.Reason != xtwitterscraper.MonitorFollowReasonUnfollowed {
			t.Fatalf("follow = %#v", follow)
		}
	})
	t.Run("post metric change", func(t *testing.T) {
		metric := eventDataVariant[xtwitterscraper.MonitorPostMetricChange](t, "post.likes.changed", `{"metric":"likes","previous":17096,"current":17200,"change":104}`)
		if metric.Metric != xtwitterscraper.MonitorPostMetricChangeMetricLikes ||
			metric.Previous != 17096 || metric.Current != 17200 || metric.Change != 104 {
			t.Fatalf("metric = %#v", metric)
		}
	})
	t.Run("post edit", func(t *testing.T) {
		edit := eventDataVariant[xtwitterscraper.MonitorPostEdit](t, "post.edited", `{"previousText":"We cut serving costs by 20%.","text":"We cut serving costs by 25%.","versionId":"2082577277246973001"}`)
		if edit.PreviousText != "We cut serving costs by 20%." || edit.Text != "We cut serving costs by 25%." || edit.VersionID != "2082577277246973001" {
			t.Fatalf("edit = %#v", edit)
		}
	})
	t.Run("post deleted", func(t *testing.T) {
		deleted := eventDataVariant[xtwitterscraper.MonitorPostDeleted](t, "post.deleted", `{"text":"We cut serving costs by 20%."}`)
		if deleted.Text != "We cut serving costs by 20%." {
			t.Fatalf("deleted = %#v", deleted)
		}
	})
	t.Run("legacy post", func(t *testing.T) {
		_, detail := decodeEventData(t, "tweet.new", `{"id":"1234567890","text":"Just launched!","author":null,"quoted_tweet":{"id":"9"}}`)
		legacy, err := detail.Data.AsMonitorLegacyPost()
		if err != nil || legacy.ID != "1234567890" || legacy.Text != "Just launched!" || legacy.Author != nil ||
			string(legacy.QuotedTweet["id"]) != `"9"` {
			t.Fatalf("legacy = %#v, err = %v", legacy, err)
		}
	})
	t.Run("unknown type", func(t *testing.T) {
		event, detail := decodeEventData(t, "account.renamed", `{"id":"1"}`)
		if variant, err := event.DataAsAny(); variant != nil || err != nil {
			t.Fatalf("event variant = %#v, err = %v", variant, err)
		}
		if variant, err := detail.DataAsAny(); variant != nil || err != nil {
			t.Fatalf("detail variant = %#v, err = %v", variant, err)
		}
	})
}

func TestEventListDecodesPostEventData(t *testing.T) {
	var res xtwitterscraper.EventListResponse
	body := `{"events":[{"id":"9002","type":"post.views.changed","monitorId":"31","monitorType":"post","occurredAt":"2026-10-22T11:05:00.000Z","data":{"metric":"views","previous":100,"current":250,"change":150}}],"hasMore":false}`
	if err := res.UnmarshalJSON([]byte(body)); err != nil {
		t.Fatal(err)
	}
	variant, err := res.Events[0].DataAsAny()
	metric, ok := variant.(xtwitterscraper.MonitorPostMetricChange)
	if err != nil || !ok || metric.Metric != xtwitterscraper.MonitorPostMetricChangeMetricViews || metric.Change != 150 || res.Events[0].Data.Change != 150 {
		t.Fatalf("variant = %#v, err = %v", variant, err)
	}
}
