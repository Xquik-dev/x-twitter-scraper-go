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

func TestXSportListGamesFixture(t *testing.T) {
	client, request := fixtureClient(t, "listFootballGames.json")
	res, err := client.X.Sports.ListGames(context.Background(), "nfl", xtwitterscraper.XSportListGamesParams{
		From:   xtwitterscraper.String("2026-10-01"),
		To:     xtwitterscraper.String("2026-10-08"),
		Team:   xtwitterscraper.String("689566294702555136"),
		Fields: xtwitterscraper.String("games.id"),
		Omit:   xtwitterscraper.String("teams"),
	})
	if err != nil {
		t.Fatal(err)
	}
	request.expect(t, http.MethodGet, "/api/v1/x/sports/nfl/games", "fields=games.id&from=2026-10-01&omit=teams&team=689566294702555136&to=2026-10-08")
	expectDecoded(t, res)
	game := res.Games[0]
	team := res.Teams["689566294702555136"]
	if res.League != "nfl" ||
		game.ID != "2055084502741811200" ||
		game.Week != 6 ||
		!game.StartsAt.Equal(time.Date(2026, time.October, 20, 0, 15, 0, 0, time.UTC)) ||
		game.Home.Points != 23 ||
		game.TeamIDs[1] != "689566304366292992" ||
		game.Periods[0].Type != "Quarter" ||
		game.Odds.Away.Spread != "text" ||
		team.Abbreviation != "SF" ||
		team.Colors[0].R != 170 ||
		team.XUserID != "43403778" {
		t.Fatalf("decoded %s", res.RawJSON())
	}
}

func TestXSportGetScoreboardFixture(t *testing.T) {
	client, request := fixtureClient(t, "getFootballScoreboard.json")
	res, err := client.X.Sports.GetScoreboard(context.Background(), "2055084502741811200", xtwitterscraper.XSportGetScoreboardParams{
		League: "nfl",
		Fields: xtwitterscraper.String("id,home"),
	})
	if err != nil {
		t.Fatal(err)
	}
	request.expect(t, http.MethodGet, "/api/v1/x/sports/nfl/games/2055084502741811200", "fields=id%2Chome")
	expectDecoded(t, res)
	if res.ID != "2055084502741811200" ||
		res.Venue != "Levi's Stadium" ||
		res.Situation.YardsToGo != 22 ||
		res.Home.TimeoutsLeft != 2 ||
		res.Periods[0].HomePoints != 13 ||
		res.Periods[0].Type != "Quarter" ||
		res.LastPlay.ID != "170" ||
		res.Players["1520565710329499648"].XUser.Username != "brockpurdy13" {
		t.Fatalf("decoded %s", res.RawJSON())
	}
}

func TestXSportGetBoxScoreFixture(t *testing.T) {
	client, request := fixtureClient(t, "getFootballBoxScore.json")
	res, err := client.X.Sports.GetBoxScore(context.Background(), "2055084502741811200", xtwitterscraper.XSportGetBoxScoreParams{
		League: "ncaafb",
	})
	if err != nil {
		t.Fatal(err)
	}
	request.expect(t, http.MethodGet, "/api/v1/x/sports/ncaafb/games/2055084502741811200/box-score", "")
	expectDecoded(t, res)
	player := res.Home.Players[0]
	if res.ID != "2055084502741811200" ||
		res.Home.Totals.ThirdDowns.Attempts != 14 ||
		res.Home.Totals.Passing.Rating != 111.85345 ||
		player.PlayerID != "1520565710329499648" ||
		player.Passing.Yards != 264 ||
		player.Defense.Sacks != 0.5 ||
		!res.Players["1520565710329499648"].XUser.BlueVerified {
		t.Fatalf("decoded %s", res.RawJSON())
	}
}

func TestXSportGetPlaysFixture(t *testing.T) {
	client, request := fixtureClient(t, "getFootballPlays.json")
	res, err := client.X.Sports.GetPlays(context.Background(), "2055084502741811200", xtwitterscraper.XSportGetPlaysParams{
		League: "nfl",
		Omit:   xtwitterscraper.String("players"),
	})
	if err != nil {
		t.Fatal(err)
	}
	request.expect(t, http.MethodGet, "/api/v1/x/sports/nfl/games/2055084502741811200/plays", "omit=players")
	expectDecoded(t, res)
	drive := res.Quarters[0].Drives[0]
	play := drive.Plays[0]
	if res.HomeTeamID != "689566294702555136" ||
		drive.StartReason != "Kickoff" ||
		drive.NetYards != 64 ||
		play.Type != "Pass" ||
		play.Details[0].Players[0].PlayerRole != "Pass" ||
		play.Details[0].JSON.Yards.Raw() != "null" ||
		play.Post.ID != "2112339130483151117" {
		t.Fatalf("decoded %s", res.RawJSON())
	}
}

func TestXSportRejectsMissingPathParams(t *testing.T) {
	client, _ := fixtureClient(t, "getFootballScoreboard.json")
	ctx := context.Background()
	if _, err := client.X.Sports.ListGames(ctx, "", xtwitterscraper.XSportListGamesParams{}); err == nil {
		t.Fatal("ListGames sent a request without a league")
	}
	if _, err := client.X.Sports.GetScoreboard(ctx, "1", xtwitterscraper.XSportGetScoreboardParams{}); err == nil {
		t.Fatal("GetScoreboard sent a request without a league")
	}
	if _, err := client.X.Sports.GetBoxScore(ctx, "", xtwitterscraper.XSportGetBoxScoreParams{League: "nfl"}); err == nil {
		t.Fatal("GetBoxScore sent a request without an id")
	}
	if _, err := client.X.Sports.GetPlays(ctx, "", xtwitterscraper.XSportGetPlaysParams{League: "nfl"}); err == nil {
		t.Fatal("GetPlays sent a request without an id")
	}
}
