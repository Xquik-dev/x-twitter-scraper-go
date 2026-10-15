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

func TestXSpaceGetFixture(t *testing.T) {
	client, request := fixtureClient(t, "getSpace.json")
	res, err := client.X.Spaces.Get(context.Background(), "1OGwblLjnQMKB", xtwitterscraper.XSpaceGetParams{
		Fields: xtwitterscraper.String("id,text,author.username"),
		Omit:   xtwitterscraper.String("entities,media"),
	})
	if err != nil {
		t.Fatal(err)
	}
	request.expect(t, http.MethodGet, "/api/v1/x/spaces/1OGwblLjnQMKB", "fields=id%2Ctext%2Cauthor.username&omit=entities%2Cmedia")
	expectDecoded(t, res)
	if res.Space.ID != "1OGwblLjnQMKB" ||
		res.Space.Title != "Extending the Network to the Edge" ||
		res.Space.State != "Ended" ||
		!res.Space.StartedAt.Equal(time.Date(2026, time.March, 11, 17, 0, 26, 227000000, time.UTC)) ||
		!res.Space.EndedAt.Equal(time.Date(2026, time.March, 11, 17, 49, 29, 190000000, time.UTC)) ||
		res.Space.TotalLiveListeners != 841 ||
		res.Space.TotalReplayWatched != 609 ||
		len(res.Space.Admins) != 1 ||
		len(res.Space.Speakers) != 0 ||
		len(res.Space.Listeners) != 0 ||
		res.Space.Creator.ID != "951329744804392960" ||
		res.Space.Creator.Username != "solana" ||
		res.Space.Creator.Name != "Solana" ||
		res.Space.Admins[0].ID != "951329744804392960" ||
		res.Space.Admins[0].Username != "solana" ||
		res.Space.Admins[0].Name != "Solana" {
		t.Fatalf("decoded %s", res.RawJSON())
	}
}

func TestXSpaceGetReplayFixture(t *testing.T) {
	client, request := fixtureClient(t, "getSpaceReplay.json")
	res, err := client.X.Spaces.GetReplay(context.Background(), "1OGwblLjnQMKB", xtwitterscraper.XSpaceGetReplayParams{
		Fields: xtwitterscraper.String("id,text,author.username"),
		Omit:   xtwitterscraper.String("entities,media"),
	})
	if err != nil {
		t.Fatal(err)
	}
	request.expect(t, http.MethodGet, "/api/v1/x/spaces/1OGwblLjnQMKB/replay", "fields=id%2Ctext%2Cauthor.username&omit=entities%2Cmedia")
	expectDecoded(t, res)
	if res.Replay.URL != "https://prod-fastly-us-east-1.video.pscp.tv/Transcoding/v1/hls/example/audio-space/playlist.m3u8?type=replay" ||
		res.Replay.StreamType != "HLS" ||
		res.Replay.Status != "LIVE_PUBLIC" {
		t.Fatalf("decoded %s", res.RawJSON())
	}
}

func TestXSpaceSearchFixture(t *testing.T) {
	client, request := fixtureClient(t, "searchSpaces.json")
	res, err := client.X.Spaces.Search(context.Background(), xtwitterscraper.XSpaceSearchParams{
		Q:         "from:solana",
		Cursor:    xtwitterscraper.String("cursor"),
		Fields:    xtwitterscraper.String("id,text,author.username"),
		Omit:      xtwitterscraper.String("entities,media"),
		PageSize:  xtwitterscraper.Int(10),
		QueryType: xtwitterscraper.XSpaceSearchParamsQueryTypeLatest,
	})
	if err != nil {
		t.Fatal(err)
	}
	request.expect(t, http.MethodGet, "/api/v1/x/spaces/search", "cursor=cursor&fields=id%2Ctext%2Cauthor.username&omit=entities%2Cmedia&pageSize=10&q=from%3Asolana&queryType=Latest")
	expectDecoded(t, res)
	if len(res.Spaces) != 1 ||
		!res.HasNextPage ||
		res.NextCursor != "DAACCgACGRElMJcAAA" ||
		res.Spaces[0].ID != "1OGwblLjnQMKB" ||
		res.Spaces[0].Title != "Extending the Network to the Edge" ||
		res.Spaces[0].State != "Ended" ||
		!res.Spaces[0].StartedAt.Equal(time.Date(2026, time.March, 11, 17, 0, 26, 227000000, time.UTC)) ||
		!res.Spaces[0].EndedAt.Equal(time.Date(2026, time.March, 11, 17, 49, 29, 190000000, time.UTC)) ||
		res.Spaces[0].TotalLiveListeners != 841 ||
		res.Spaces[0].TotalReplayWatched != 609 ||
		len(res.Spaces[0].Admins) != 1 ||
		len(res.Spaces[0].Speakers) != 0 ||
		len(res.Spaces[0].Listeners) != 0 ||
		res.Spaces[0].Admins[0].ID != "951329744804392960" ||
		res.Spaces[0].Admins[0].Username != "solana" ||
		res.Spaces[0].Admins[0].Name != "Solana" {
		t.Fatalf("decoded %s", res.RawJSON())
	}
}
