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

func TestXBroadcastGetFixture(t *testing.T) {
	client, request := fixtureClient(t, "getBroadcast.json")
	res, err := client.X.Broadcasts.Get(context.Background(), "1nGeLMWbkpaKX", xtwitterscraper.XBroadcastGetParams{
		Fields: xtwitterscraper.String("id,text,author.username"),
		Omit:   xtwitterscraper.String("entities,media"),
	})
	if err != nil {
		t.Fatal(err)
	}
	request.expect(t, http.MethodGet, "/api/v1/x/broadcasts/1nGeLMWbkpaKX", "fields=id%2Ctext%2Cauthor.username&omit=entities%2Cmedia")
	expectDecoded(t, res)
	if res.Broadcast.ID != "1nGeLMWbkpaKX" ||
		res.Broadcast.Title != "USSF-259 Mission" ||
		res.Broadcast.State != "Ended" ||
		!res.Broadcast.StartedAt.Equal(time.Date(2026, time.September, 17, 0, 54, 56, 506000000, time.UTC)) ||
		!res.Broadcast.EndedAt.Equal(time.Date(2026, time.September, 17, 1, 17, 35, 36000000, time.UTC)) ||
		res.Broadcast.TotalWatched != 161758 ||
		res.Broadcast.TotalWatching != 21639 ||
		res.Broadcast.Width != 3840 ||
		res.Broadcast.Height != 2160 ||
		res.Broadcast.Creator.ID != "34743251" ||
		res.Broadcast.Creator.Username != "SpaceX" ||
		res.Broadcast.Creator.Name != "SpaceX" {
		t.Fatalf("decoded %s", res.RawJSON())
	}
}
