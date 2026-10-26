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

func expectPostMonitor(t *testing.T, res *xtwitterscraper.PostMonitor) {
	t.Helper()
	expectDecoded(t, res)
	if res.ID != "31" ||
		res.Name != "Launch post" ||
		res.PostID != "2104984505677430978" ||
		res.PostURL != "https://x.com/xquik/status/2104984505677430978" ||
		res.AuthorUsername != "xquik" ||
		res.AuthorAvatarURL != "https://pbs.twimg.com/profile_images/1/xquik_normal.jpg" ||
		res.AuthorVerified != "blue" ||
		len(res.Media) != 1 || res.Media[0].Kind != "photo" || res.Media[0].URL != "https://pbs.twimg.com/media/GxPhoto.jpg" ||
		!res.Baseline.At.Equal(time.Date(2026, time.October, 22, 11, 0, 0, 0, time.UTC)) ||
		res.Baseline.Views != 541100 || res.Baseline.Likes != 1885 || res.Baseline.Retweets != 112 ||
		res.Baseline.Replies != 85 || res.Baseline.Quotes != 87 || res.Baseline.Bookmarks != 470 ||
		!res.PostedAt.Equal(time.Date(2026, time.September, 29, 17, 19, 49, 0, time.UTC)) ||
		len(res.EventTypes) != 3 || res.EventTypes[2] != "post.likes.changed" ||
		!res.IsActive ||
		res.PausedReason != "x_post_deleted" ||
		res.Metrics.Views.Current != 1885 ||
		res.Metrics.Bookmarks.Gain != 12 {
		t.Fatalf("decoded %s", res.RawJSON())
	}
}

func TestMonitorPostNewFixture(t *testing.T) {
	client, request := fixtureClient(t, "getPostMonitor.json")
	res, err := client.Monitors.Posts.New(context.Background(), xtwitterscraper.MonitorPostNewParams{
		Post:       "https://x.com/xquik/status/2104984505677430978",
		Name:       xtwitterscraper.String("Launch post"),
		EventTypes: []string{"post.reply.new", "post.likes.changed"},
	})
	if err != nil {
		t.Fatal(err)
	}
	request.expect(t, http.MethodPost, "/api/v1/monitors/posts", "")
	request.expectBody(t, `{"post":"https://x.com/xquik/status/2104984505677430978","name":"Launch post","eventTypes":["post.reply.new","post.likes.changed"]}`)
	expectPostMonitor(t, res)
}

func TestMonitorPostGetFixture(t *testing.T) {
	client, request := fixtureClient(t, "getPostMonitor.json")
	res, err := client.Monitors.Posts.Get(context.Background(), "31")
	if err != nil {
		t.Fatal(err)
	}
	request.expect(t, http.MethodGet, "/api/v1/monitors/posts/31", "")
	expectPostMonitor(t, res)
}

func TestMonitorPostUpdateFixture(t *testing.T) {
	client, request := fixtureClient(t, "getPostMonitor.json")
	res, err := client.Monitors.Posts.Update(context.Background(), "31", xtwitterscraper.MonitorPostUpdateParams{
		IsActive:   xtwitterscraper.Bool(false),
		EventTypes: []string{"post.deleted"},
	})
	if err != nil {
		t.Fatal(err)
	}
	request.expect(t, http.MethodPatch, "/api/v1/monitors/posts/31", "")
	request.expectBody(t, `{"isActive":false,"eventTypes":["post.deleted"]}`)
	expectPostMonitor(t, res)
}

func TestMonitorPostListFixture(t *testing.T) {
	client, request := fixtureClient(t, "listPostMonitors.json")
	res, err := client.Monitors.Posts.List(context.Background(), xtwitterscraper.MonitorPostListParams{
		Cursor: xtwitterscraper.String("abc123"),
		Limit:  xtwitterscraper.Int(20),
	})
	if err != nil {
		t.Fatal(err)
	}
	request.expect(t, http.MethodGet, "/api/v1/monitors/posts", "cursor=abc123&limit=20")
	expectDecoded(t, res)
	if res.Total != 1 || !res.HasMore || res.NextCursor != "abc123" || len(res.Monitors) != 1 {
		t.Fatalf("decoded %s", res.RawJSON())
	}
	expectPostMonitor(t, &res.Monitors[0])
}

func TestMonitorPostDeleteFixture(t *testing.T) {
	client, request := fixtureClient(t, "deletePostMonitor.json")
	res, err := client.Monitors.Posts.Delete(context.Background(), "31")
	if err != nil {
		t.Fatal(err)
	}
	request.expect(t, http.MethodDelete, "/api/v1/monitors/posts/31", "")
	expectDecoded(t, res)
	if res.DeletionStatus != "deleting" || res.StatusURL != "/api/v1/monitors/posts/31" || !res.Success {
		t.Fatalf("decoded %s", res.RawJSON())
	}
}

func TestMonitorPostGetMetricsFixture(t *testing.T) {
	client, request := fixtureClient(t, "getPostMonitorMetrics.json")
	res, err := client.Monitors.Posts.GetMetrics(context.Background(), "31", xtwitterscraper.MonitorPostGetMetricsParams{
		Hours: xtwitterscraper.Int(12),
	})
	if err != nil {
		t.Fatal(err)
	}
	request.expect(t, http.MethodGet, "/api/v1/monitors/posts/31/metrics", "hours=12")
	expectDecoded(t, res)
	if res.Hours != 12 || len(res.Buckets) != 1 ||
		!res.Buckets[0].At.Equal(time.Date(2026, time.October, 22, 11, 0, 0, 0, time.UTC)) ||
		res.Buckets[0].Metrics.Views.Value != 541600 ||
		res.Buckets[0].Metrics.Views.Gain != 500 ||
		res.Buckets[0].Metrics.Bookmarks.Value != 471 {
		t.Fatalf("decoded %s", res.RawJSON())
	}
}

func TestMonitorPostRejectsAnEmptyID(t *testing.T) {
	client, _ := fixtureClient(t, "getPostMonitor.json")
	ctx := context.Background()
	if _, err := client.Monitors.Posts.Get(ctx, ""); err == nil {
		t.Fatal("Get sent a request without an id")
	}
	if _, err := client.Monitors.Posts.Update(ctx, "", xtwitterscraper.MonitorPostUpdateParams{}); err == nil {
		t.Fatal("Update sent a request without an id")
	}
	if _, err := client.Monitors.Posts.Delete(ctx, ""); err == nil {
		t.Fatal("Delete sent a request without an id")
	}
	if _, err := client.Monitors.Posts.GetMetrics(ctx, "", xtwitterscraper.MonitorPostGetMetricsParams{}); err == nil {
		t.Fatal("GetMetrics sent a request without an id")
	}
}

func TestEventListFiltersByPostMonitorFixture(t *testing.T) {
	client, request := fixtureClient(t, "listPostEvents.json")
	res, err := client.Events.List(context.Background(), xtwitterscraper.EventListParams{
		PostMonitorID: xtwitterscraper.String("31"),
	})
	if err != nil {
		t.Fatal(err)
	}
	request.expect(t, http.MethodGet, "/api/v1/events", "postMonitorId=31")
	expectDecoded(t, res)
	event := res.Events[0]
	if event.MonitorType != xtwitterscraper.EventMonitorTypePost ||
		event.Type != "post.reply.new" ||
		event.PostID != "2104984505677430978" ||
		event.PostURL != "https://x.com/xquik/status/2104984505677430978" ||
		event.AuthorID != "1661816432185823232" ||
		event.AuthorUsername != "xquik" {
		t.Fatalf("decoded %s", res.RawJSON())
	}
}
