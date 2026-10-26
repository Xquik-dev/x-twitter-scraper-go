// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

package xtwitterscraper_test

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/Xquik-dev/x-twitter-scraper-go"
	"github.com/Xquik-dev/x-twitter-scraper-go/option"
)

const eventStreamFixture = "retry: 3000\n\n" +
	"id: 42\nevent: tweet.new\ndata: {\"id\":\"42\",\"type\":\"tweet.new\",\"monitorId\":\"7\",\"monitorType\":\"account\",\"occurredAt\":\"2025-01-15T12:00:00.000Z\",\"data\":{}}\n\n" +
	": keep-alive\n\n" +
	"id: 43\nevent: tweet.reply\ndata: {\"id\":\"43\",\"type\":\"tweet.reply\",\"monitorId\":\"7\",\"monitorType\":\"account\",\"occurredAt\":\"2025-01-15T12:00:05.000Z\",\"data\":{\"tweetId\":\"1893456789012345678\"}}\n\n"

func TestEventStreamStreamingFixture(t *testing.T) {
	var request *http.Request
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		request = r.Clone(context.Background())
		w.Header().Set("Content-Type", "text/event-stream")
		if _, err := w.Write([]byte(eventStreamFixture)); err != nil {
			t.Error(err)
		}
	}))
	defer server.Close()
	client := xtwitterscraper.NewClient(option.WithBaseURL(server.URL+"/api/v1/"), option.WithAPIKey("My API Key"), option.WithMaxRetries(0))
	stream := client.Events.StreamStreaming(context.Background(), xtwitterscraper.EventStreamParams{
		LastEventID: xtwitterscraper.String("41"),
		Include:     xtwitterscraper.EventStreamParamsIncludePostMetrics,
		LookBack:    xtwitterscraper.Bool(true),
	})
	defer stream.Close()
	var events []xtwitterscraper.Event
	for stream.Next() {
		events = append(events, stream.Current())
	}
	if err := stream.Err(); err != nil {
		t.Fatal(err)
	}
	if request.Method != http.MethodGet || request.URL.Path != "/api/v1/events/stream" || request.URL.Query().Encode() != "include=post_metrics&lookBack=true" ||
		request.Header.Get("Last-Event-ID") != "41" || request.Header.Get("Accept") != "text/event-stream" {
		t.Fatalf("request = %s %s, headers %v", request.Method, request.URL, request.Header)
	}
	if len(events) != 2 ||
		events[0].ID != "42" ||
		events[0].Type != "tweet.new" ||
		events[0].MonitorID != "7" ||
		events[0].MonitorType != xtwitterscraper.EventMonitorTypeAccount ||
		!events[0].OccurredAt.Equal(time.Date(2025, time.January, 15, 12, 0, 0, 0, time.UTC)) ||
		events[1].ID != "43" ||
		events[1].Type != "tweet.reply" ||
		events[1].Data.JSON.ExtraFields["tweetId"].Raw() != `"1893456789012345678"` {
		t.Fatalf("events = %+v", events)
	}
}

func TestEventStreamStreamingReportsAPIErrors(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusTooManyRequests)
		if _, err := w.Write([]byte(`{"error":"too_many_streams","message":"Close a stream first. Keep up to 5 open per account."}`)); err != nil {
			t.Error(err)
		}
	}))
	defer server.Close()
	client := xtwitterscraper.NewClient(option.WithBaseURL(server.URL+"/api/v1/"), option.WithAPIKey("My API Key"), option.WithMaxRetries(0))
	stream := client.Events.StreamStreaming(context.Background(), xtwitterscraper.EventStreamParams{})
	var apierr *xtwitterscraper.Error
	if stream.Next() || !errors.As(stream.Err(), &apierr) || apierr.StatusCode != http.StatusTooManyRequests {
		t.Fatalf("err = %v", stream.Err())
	}
	if err := stream.Close(); err != nil {
		t.Fatal(err)
	}
}
