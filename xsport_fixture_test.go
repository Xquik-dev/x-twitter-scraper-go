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
	recent := res.Teams["689566294702555136"].RecentGames[0]
	if res.ID != "2055084502741811200" ||
		res.Home.Totals.ThirdDowns.Attempts != 14 ||
		res.Home.Totals.Passing.Rating != 111.85345 ||
		player.PlayerID != "1520565710329499648" ||
		player.Passing.Yards != 264 ||
		player.Defense.Sacks != 0.5 ||
		!res.Players["1520565710329499648"].XUser.BlueVerified ||
		recent.ID != "2083346094763671552" ||
		!recent.StartsAt.Equal(time.Date(2026, time.October, 17, 23, 30, 0, 0, time.UTC)) ||
		recent.Status != "Closed" ||
		recent.Away.TeamID != "898300623782727680" ||
		recent.Away.Points != 34 {
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

func TestXSportGetAudienceFixture(t *testing.T) {
	client, request := fixtureClient(t, "getSportsGameAudience.json")
	res, err := client.X.Sports.GetAudience(context.Background(), "2055084502741811200", xtwitterscraper.XSportGetAudienceParams{
		League: "ncaafb",
		Fields: xtwitterscraper.String("viewers,chat"),
	})
	if err != nil {
		t.Fatal(err)
	}
	request.expect(t, http.MethodGet, "/api/v1/x/sports/ncaafb/games/2055084502741811200/audience", "fields=viewers%2Cchat")
	expectDecoded(t, res)
	team := res.Teams["689566294702555136"]
	if res.Viewers.Count != 155 ||
		!res.Viewers.AsOf.Equal(time.Date(2026, 10, 20, 6, 19, 50, 0, time.UTC)) ||
		res.TeamPicks.Total != 7976 ||
		res.TeamPicks.Teams[1].Count != 2978 ||
		res.Chat.Available ||
		res.Chat.OpensIn.Hours != 12 ||
		team.Location != "Virginia Tech" ||
		team.Record.Wins != 7 {
		t.Fatalf("decoded %s", res.RawJSON())
	}
}

func TestXSportGetRosterFixture(t *testing.T) {
	client, request := fixtureClient(t, "getFootballRoster.json")
	res, err := client.X.Sports.GetRoster(context.Background(), "2055084502741811200", xtwitterscraper.XSportGetRosterParams{
		League: "nfl",
		Omit:   xtwitterscraper.String("teams"),
	})
	if err != nil {
		t.Fatal(err)
	}
	request.expect(t, http.MethodGet, "/api/v1/x/sports/nfl/games/2055084502741811200/roster", "omit=teams")
	expectDecoded(t, res)
	purdy := res.Players["1520565710329499648"]
	unnamed := res.Players["1785102096355790848"].XUser
	if res.League != "nfl" ||
		res.Rosters[0].TeamID != "689566294702555136" ||
		res.Rosters[0].PlayerIDs[1] != "2047931086345285632" ||
		res.Rosters[1].PlayerIDs[0] != "1785102096355790848" ||
		res.Teams["689566304366292992"].Abbreviation != "WAS" ||
		purdy.Position != "QB" ||
		purdy.Headshots.Medium != "https://pbs.twimg.com/semantic_core_img/1960514226075602944/OZ6rz5aH?format=png&name=orig" ||
		purdy.JSON.College.Raw() != "null" ||
		purdy.XUser.Username != "brockpurdy13" ||
		!purdy.XUser.JSON.Protected.Valid() ||
		purdy.XUser.Protected ||
		purdy.XUser.JSON.LabelType.Raw() != "null" ||
		res.Players["2047931086345285632"].XUser.LabelType != "BusinessLabel" ||
		unnamed.ID != "2503135158" ||
		unnamed.JSON.Username.Raw() != "null" ||
		unnamed.JSON.SuperFollowEligible.Raw() != "null" {
		t.Fatalf("decoded %s", res.RawJSON())
	}
}

func TestXSportGetPlayerPropsFixture(t *testing.T) {
	client, request := fixtureClient(t, "getFootballPlayerProps.json")
	res, err := client.X.Sports.GetPlayerProps(context.Background(), "2055084502741929984", xtwitterscraper.XSportGetPlayerPropsParams{
		League: "nfl",
		Fields: xtwitterscraper.String("categories"),
	})
	if err != nil {
		t.Fatal(err)
	}
	request.expect(t, http.MethodGet, "/api/v1/x/sports/nfl/games/2055084502741929984/player-props", "fields=categories")
	expectDecoded(t, res)
	passing := res.Categories[0]
	scorer := res.Categories[1].Props[0]
	if res.SportsbookURL != "https://playmgmsports.onelink.me/TkMx" ||
		passing.Key != "passing_yards" ||
		passing.Rank != 1 ||
		passing.Props[1].PlayerID != "1783744076115783681" ||
		passing.Props[0].Line != 248.5 ||
		passing.Props[0].Odds != -115 ||
		passing.Props[0].AlternateOdds != -115 ||
		!passing.Props[0].OpenForBetting ||
		res.Categories[1].Name != "First Touchdown Scorer" ||
		scorer.JSON.Line.Raw() != "null" ||
		scorer.Odds != 390 ||
		scorer.JSON.AlternateOdds.Raw() != "null" ||
		res.Players[scorer.PlayerID].Name != "D'Andre Swift" ||
		res.Players["1783744053755990016"].Jersey != "10" ||
		res.Players["1783744053755990016"].JSON.XUser.Raw() != "null" {
		t.Fatalf("decoded %s", res.RawJSON())
	}
}

func TestXSportGetPostsFixture(t *testing.T) {
	client, request := fixtureClient(t, "getSportsGamePosts.json")
	res, err := client.X.Sports.GetPosts(context.Background(), "2055084502741811200", xtwitterscraper.XSportGetPostsParams{
		League:    "nfl",
		Cursor:    xtwitterscraper.String("next"),
		PageSize:  xtwitterscraper.Int(20),
		QueryType: xtwitterscraper.XSportGetPostsParamsQueryTypeLatest,
	})
	if err != nil {
		t.Fatal(err)
	}
	request.expect(t, http.MethodGet, "/api/v1/x/sports/nfl/games/2055084502741811200/posts", "cursor=next&pageSize=20&queryType=Latest")
	expectDecoded(t, res)
	if res.Query != `(#WASvsSF OR #WASatSF OR "San Francisco 49ers" OR "Washington Commanders" OR @49ers OR @Commanders) since_time:1792455300 until_time:1792541700 -filter:nativeretweets` ||
		!res.HasNextPage ||
		res.NextCursor[:10] != "DAACCgACHV" ||
		res.Tweets[0].ID != "2102761519985332442" ||
		res.Tweets[0].Author.Username != "NASA" ||
		res.Tweets[0].LikeCount != 9775 {
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
	if _, err := client.X.Sports.GetAudience(ctx, "1", xtwitterscraper.XSportGetAudienceParams{}); err == nil {
		t.Fatal("GetAudience sent a request without a league")
	}
	if _, err := client.X.Sports.GetAudience(ctx, "", xtwitterscraper.XSportGetAudienceParams{League: "nfl"}); err == nil {
		t.Fatal("GetAudience sent a request without an id")
	}
	if _, err := client.X.Sports.GetRoster(ctx, "1", xtwitterscraper.XSportGetRosterParams{}); err == nil {
		t.Fatal("GetRoster sent a request without a league")
	}
	if _, err := client.X.Sports.GetRoster(ctx, "", xtwitterscraper.XSportGetRosterParams{League: "nfl"}); err == nil {
		t.Fatal("GetRoster sent a request without an id")
	}
	if _, err := client.X.Sports.GetPlayerProps(ctx, "1", xtwitterscraper.XSportGetPlayerPropsParams{}); err == nil {
		t.Fatal("GetPlayerProps sent a request without a league")
	}
	if _, err := client.X.Sports.GetPlayerProps(ctx, "", xtwitterscraper.XSportGetPlayerPropsParams{League: "nfl"}); err == nil {
		t.Fatal("GetPlayerProps sent a request without an id")
	}
	if _, err := client.X.Sports.GetPosts(ctx, "1", xtwitterscraper.XSportGetPostsParams{}); err == nil {
		t.Fatal("GetPosts sent a request without a league")
	}
	if _, err := client.X.Sports.GetPosts(ctx, "", xtwitterscraper.XSportGetPostsParams{League: "nfl"}); err == nil {
		t.Fatal("GetPosts sent a request without an id")
	}
}
