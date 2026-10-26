// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

package xtwitterscraper

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/Xquik-dev/x-twitter-scraper-go/internal/apijson"
	"github.com/Xquik-dev/x-twitter-scraper-go/internal/apiquery"
	"github.com/Xquik-dev/x-twitter-scraper-go/internal/requestconfig"
	"github.com/Xquik-dev/x-twitter-scraper-go/option"
	"github.com/Xquik-dev/x-twitter-scraper-go/packages/param"
	"github.com/Xquik-dev/x-twitter-scraper-go/packages/respjson"
	"github.com/Xquik-dev/x-twitter-scraper-go/shared"
)

// XSportService covers NFL & college football games, scoreboards, plays, box
// scores, audiences, rosters, player props & posts.
//
// XSportService contains methods and other services that help with interacting
// with the x-twitter-scraper API.
//
// You should not instantiate this service directly, and instead use the
// [NewXSportService] method instead.
type XSportService struct {
	options []option.RequestOption
}

// NewXSportService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewXSportService(opts ...option.RequestOption) (r XSportService) {
	r = XSportService{}
	r.options = opts
	return
}

// Lists NFL or college football games in a time window, with teams & scores. Use
// it to find game IDs & team IDs. Without from or to, X picks about 5 weeks around
// today. Each request costs 1 credit.
func (r *XSportService) ListGames(ctx context.Context, league string, query XSportListGamesParams, opts ...option.RequestOption) (res *FootballSchedule, err error) {
	opts = slices.Concat(r.options, opts)
	if league == "" {
		err = errors.New("missing required league parameter")
		return nil, err
	}
	path := fmt.Sprintf("x/sports/%s/games", url.PathEscape(league))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Returns 1 game's scoreboard. It has the teams, the score by period, the status &
// clock, down & distance, the venue & the latest play. Each request costs 1 credit.
func (r *XSportService) GetScoreboard(ctx context.Context, id string, params XSportGetScoreboardParams, opts ...option.RequestOption) (res *FootballScoreboard, err error) {
	opts = slices.Concat(r.options, opts)
	if params.League == "" {
		err = errors.New("missing required league parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("x/sports/%s/games/%s", url.PathEscape(params.League), url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Returns each team's totals & each player's stats for 1 game. Stats cover passing,
// rushing, receiving, defense, kicking & returns. A game that has not started has
// no stats yet. Each request costs 1 credit.
func (r *XSportService) GetBoxScore(ctx context.Context, id string, params XSportGetBoxScoreParams, opts ...option.RequestOption) (res *FootballBoxScore, err error) {
	opts = slices.Concat(r.options, opts)
	if params.League == "" {
		err = errors.New("missing required league parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("x/sports/%s/games/%s/box-score", url.PathEscape(params.League), url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Returns every drive & play of 1 game, by quarter. Each play has its clock, down &
// distance, details, penalties & linked posts. Each request costs 1 credit.
func (r *XSportService) GetPlays(ctx context.Context, id string, params XSportGetPlaysParams, opts ...option.RequestOption) (res *FootballPlays, err error) {
	opts = slices.Concat(r.options, opts)
	if params.League == "" {
		err = errors.New("missing required league parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("x/sports/%s/games/%s/plays", url.PathEscape(params.League), url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Returns who follows 1 game on X. It has live viewers, fan picks per team & the
// game chat's state. Each request costs 1 credit.
func (r *XSportService) GetAudience(ctx context.Context, id string, params XSportGetAudienceParams, opts ...option.RequestOption) (res *SportsGameAudience, err error) {
	opts = slices.Concat(r.options, opts)
	if params.League == "" {
		err = errors.New("missing required league parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("x/sports/%s/games/%s/audience", url.PathEscape(params.League), url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// GetRoster returns both teams' players for 1 NFL game. Each player has a jersey,
// position, headshot & X account. Each request costs 1 credit.
func (r *XSportService) GetRoster(ctx context.Context, id string, params XSportGetRosterParams, opts ...option.RequestOption) (res *FootballRoster, err error) {
	opts = slices.Concat(r.options, opts)
	if params.League == "" {
		err = errors.New("missing required league parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("x/sports/%s/games/%s/roster", url.PathEscape(params.League), url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// GetPlayerProps returns the player prop lines & American odds X shows for 1 NFL
// game. X lists props before & during a game. Other games return no categories.
// Each request costs 1 credit.
func (r *XSportService) GetPlayerProps(ctx context.Context, id string, params XSportGetPlayerPropsParams, opts ...option.RequestOption) (res *FootballProps, err error) {
	opts = slices.Concat(r.options, opts)
	if params.League == "" {
		err = errors.New("missing required league parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("x/sports/%s/games/%s/player-props", url.PathEscape(params.League), url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// GetPosts returns the posts X shows on a game's page, with the X search that
// finds them. Works before, during & after the game. Pages hold up to 20 posts, &
// each post costs 1 credit.
func (r *XSportService) GetPosts(ctx context.Context, id string, params XSportGetPostsParams, opts ...option.RequestOption) (res *SportsGamePosts, err error) {
	opts = slices.Concat(r.options, opts)
	if params.League == "" {
		err = errors.New("missing required league parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("x/sports/%s/games/%s/posts", url.PathEscape(params.League), url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// A team's side of a game & its points.
type FootballTeamScore struct {
	// The team, a key of teams.
	TeamID string `json:"teamId" api:"required"`
	// The team's points so far.
	Points int64 `json:"points" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		TeamID      respjson.Field
		Points      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FootballTeamScore) RawJSON() string { return r.JSON.raw }
func (r *FootballTeamScore) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A period of a listed game.
type FootballSchedulePeriod struct {
	// X's name for the period, such as Quarter.
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FootballSchedulePeriod) RawJSON() string { return r.JSON.raw }
func (r *FootballSchedulePeriod) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Betting odds titles for 1 team, as X shows them.
type FootballOdds struct {
	// The money line. X sends it for the away team only.
	MoneyLine string `json:"moneyLine" api:"required"`
	// The point spread.
	Spread string `json:"spread" api:"required"`
	// The over or under total.
	Total string `json:"total" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MoneyLine   respjson.Field
		Spread      respjson.Field
		Total       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FootballOdds) RawJSON() string { return r.JSON.raw }
func (r *FootballOdds) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Betting odds for each team, or null when X has none.
type FootballScheduleGameOdds struct {
	// Betting odds titles for 1 team, as X shows them.
	Home FootballOdds `json:"home" api:"required"`
	// Betting odds titles for 1 team, as X shows them.
	Away FootballOdds `json:"away" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Home        respjson.Field
		Away        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FootballScheduleGameOdds) RawJSON() string { return r.JSON.raw }
func (r *FootballScheduleGameOdds) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A listed game, with its teams & score.
type FootballScheduleGame struct {
	// The game ID the other sports routes take.
	ID string `json:"id" api:"required"`
	// This game's part of the season, such as RegularSeason.
	SeasonType string `json:"seasonType" api:"required"`
	// This game's week of the season.
	Week int64 `json:"week" api:"required"`
	// This game's kickoff time in UTC, or null when Xquik cannot read it.
	StartsAt time.Time `json:"startsAt" api:"required" format:"date-time"`
	// This game's status from X, such as Scheduled, InProgress or Closed.
	Status string `json:"status" api:"required"`
	// This game's clock, as X writes it.
	Clock string `json:"clock" api:"required"`
	// This game's current or last quarter.
	Quarter int64 `json:"quarter" api:"required"`
	// A team's side of a game & its points.
	Home FootballTeamScore `json:"home" api:"required"`
	// A team's side of a game & its points.
	Away FootballTeamScore `json:"away" api:"required"`
	// The game's teams, keys of teams.
	TeamIDs []string `json:"teamIds" api:"required"`
	// The periods played so far.
	Periods []FootballSchedulePeriod `json:"periods" api:"required"`
	// Betting odds for each team, or null when X has none.
	Odds FootballScheduleGameOdds `json:"odds" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		SeasonType  respjson.Field
		Week        respjson.Field
		StartsAt    respjson.Field
		Status      respjson.Field
		Clock       respjson.Field
		Quarter     respjson.Field
		Home        respjson.Field
		Away        respjson.Field
		TeamIDs     respjson.Field
		Periods     respjson.Field
		Odds        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FootballScheduleGame) RawJSON() string { return r.JSON.raw }
func (r *FootballScheduleGame) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SportsTeamColor struct {
	// Red, from 0 to 255.
	R int64 `json:"r" api:"required"`
	// Green, from 0 to 255.
	G int64 `json:"g" api:"required"`
	// Blue, from 0 to 255.
	B int64 `json:"b" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		R           respjson.Field
		G           respjson.Field
		B           respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SportsTeamColor) RawJSON() string { return r.JSON.raw }
func (r *SportsTeamColor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The team's logo images.
type SportsTeamLogos struct {
	// The logo X shows by default.
	Default string `json:"default" api:"required"`
	// The logo for light backgrounds.
	Light string `json:"light" api:"required"`
	// The logo for dark backgrounds.
	Dark string `json:"dark" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Default     respjson.Field
		Light       respjson.Field
		Dark        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SportsTeamLogos) RawJSON() string { return r.JSON.raw }
func (r *SportsTeamLogos) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The team's wins, losses & ties this season. X sends it only in college
// football games.
type SportsTeamRecord struct {
	// Games the team won this season.
	Wins int64 `json:"wins" api:"required"`
	// Games the team lost this season.
	Losses int64 `json:"losses" api:"required"`
	// Games the team tied this season.
	Ties int64 `json:"ties" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Wins        respjson.Field
		Losses      respjson.Field
		Ties        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SportsTeamRecord) RawJSON() string { return r.JSON.raw }
func (r *SportsTeamRecord) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A team, as X's sports pages show it.
type SportsTeam struct {
	// The team ID that games & players refer to.
	ID string `json:"id" api:"required"`
	// The team's full name.
	Name string `json:"name" api:"required"`
	// The team's short name.
	ShortName string `json:"shortName" api:"required"`
	// The team's abbreviation.
	Abbreviation string `json:"abbreviation" api:"required"`
	// Another ID X keeps for the team.
	AlternateID string `json:"alternateId" api:"required"`
	// The school or place the team plays for. X sends it in college football
	// schedules & in a game's audience.
	Location string `json:"location" api:"required"`
	// The team's colors, as red, green & blue from 0 to 255.
	Colors []SportsTeamColor `json:"colors" api:"required"`
	// The team's logo images.
	Logos SportsTeamLogos `json:"logos" api:"required"`
	// The user ID of the team's X account.
	XUserID string `json:"xUserId" api:"required"`
	// The team's wins, losses & ties this season. X sends it only in college
	// football games.
	Record SportsTeamRecord `json:"record" api:"required"`
	// The team's recent games, newest first. Games without a kickoff time come
	// last. X sends them only with the teams of a college football game. The list
	// is empty elsewhere.
	RecentGames []SportsTeamGame `json:"recentGames" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Name         respjson.Field
		ShortName    respjson.Field
		Abbreviation respjson.Field
		AlternateID  respjson.Field
		Location     respjson.Field
		Colors       respjson.Field
		Logos        respjson.Field
		XUserID      respjson.Field
		Record       respjson.Field
		RecentGames  respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SportsTeam) RawJSON() string { return r.JSON.raw }
func (r *SportsTeam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SportsTeamGame holds a recent game of a team, as X lists it with the team.
type SportsTeamGame struct {
	// The game ID, which the game routes take.
	ID string `json:"id" api:"required"`
	// When this game kicked off, in UTC. Null when Xquik can't read X's time.
	StartsAt time.Time `json:"startsAt" api:"required" format:"date-time"`
	// The stadium. X has sent none in recent games so far.
	Venue string `json:"venue" api:"required"`
	// This game's status as X names it, such as Closed.
	Status string `json:"status" api:"required"`
	// The game clock, as X writes it. X has sent none so far.
	Clock string `json:"clock" api:"required"`
	// The current or last quarter. X has sent none so far.
	Quarter int64 `json:"quarter" api:"required"`
	// A team's side of a game and its points.
	Home FootballTeamScore `json:"home" api:"required"`
	// A team's side of a game and its points.
	Away FootballTeamScore `json:"away" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		StartsAt    respjson.Field
		Venue       respjson.Field
		Status      respjson.Field
		Clock       respjson.Field
		Quarter     respjson.Field
		Home        respjson.Field
		Away        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// RawJSON returns the unmodified JSON received from the API.
func (r SportsTeamGame) RawJSON() string { return r.JSON.raw }
func (r *SportsTeamGame) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A league's games in a time window.
type FootballSchedule struct {
	// The league whose games this list holds.
	League string `json:"league" api:"required"`
	// The games, earliest first.
	Games []FootballScheduleGame `json:"games" api:"required"`
	// The teams the games name, by team ID.
	Teams map[string]SportsTeam `json:"teams" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		League      respjson.Field
		Games       respjson.Field
		Teams       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FootballSchedule) RawJSON() string { return r.JSON.raw }
func (r *FootballSchedule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Down & distance now, or null before kickoff.
type FootballScoreboardSituation struct {
	// The current down, from 1 to 4.
	Down int64 `json:"down" api:"required"`
	// The yards the offense needs now for a first down.
	YardsToGo int64 `json:"yardsToGo" api:"required"`
	// The ball's current yard line.
	Yardline int64 `json:"yardline" api:"required"`
	// The team whose half the ball is in, a key of teams.
	YardlineTeamID string `json:"yardlineTeamId" api:"required"`
	// The team with the ball, a key of teams.
	PossessionTeamID string `json:"possessionTeamId" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Down             respjson.Field
		YardsToGo        respjson.Field
		Yardline         respjson.Field
		YardlineTeamID   respjson.Field
		PossessionTeamID respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FootballScoreboardSituation) RawJSON() string { return r.JSON.raw }
func (r *FootballScoreboardSituation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A team's side of a scoreboard.
type FootballSide struct {
	// The team, a key of teams.
	TeamID string `json:"teamId" api:"required"`
	// The team's points so far.
	Points int64 `json:"points" api:"required"`
	// The timeouts the team has left.
	TimeoutsLeft int64 `json:"timeoutsLeft" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		TeamID       respjson.Field
		Points       respjson.Field
		TimeoutsLeft respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FootballSide) RawJSON() string { return r.JSON.raw }
func (r *FootballSide) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A period of a game & its points.
type FootballPeriod struct {
	// X's name for the period, such as Quarter.
	Type string `json:"type" api:"required"`
	// The period's number, or null when X sends no digits.
	Number int64 `json:"number" api:"required"`
	// The period's place in the game, from 1.
	Sequence int64 `json:"sequence" api:"required"`
	// The home team's points in the period.
	HomePoints int64 `json:"homePoints" api:"required"`
	// The away team's points in the period.
	AwayPoints int64 `json:"awayPoints" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		Number      respjson.Field
		Sequence    respjson.Field
		HomePoints  respjson.Field
		AwayPoints  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FootballPeriod) RawJSON() string { return r.JSON.raw }
func (r *FootballPeriod) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The points a scoring play added, or null.
type FootballPlayScore struct {
	// The home team's score once this scoring play counted.
	HomePoints int64 `json:"homePoints" api:"required"`
	// The away team's score once this scoring play counted.
	AwayPoints int64 `json:"awayPoints" api:"required"`
	// The points the play scored.
	Points int64 `json:"points" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		HomePoints  respjson.Field
		AwayPoints  respjson.Field
		Points      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FootballPlayScore) RawJSON() string { return r.JSON.raw }
func (r *FootballPlayScore) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Down & distance when the play started, or null.
type FootballPlayStart struct {
	// The down, from 1 to 4.
	Down int64 `json:"down" api:"required"`
	// The yards needed for a first down.
	YardsToGo int64 `json:"yardsToGo" api:"required"`
	// The yard line of the ball.
	Yardline int64 `json:"yardline" api:"required"`
	// Whose half the ball was in. Home or Away in NFL games, the team's
	// alternateId in college games.
	YardlineSide string `json:"yardlineSide" api:"required"`
	// The team with the ball. Home or Away in NFL games, the team's alternateId in
	// college games.
	Possession string `json:"possession" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Down         respjson.Field
		YardsToGo    respjson.Field
		Yardline     respjson.Field
		YardlineSide respjson.Field
		Possession   respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FootballPlayStart) RawJSON() string { return r.JSON.raw }
func (r *FootballPlayStart) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Flags X shows on the field for the play, or null.
type FootballPlayOnField struct {
	// Whether a penalty called the play back.
	CalledBack bool `json:"calledBack" api:"required"`
	// Whether a flag was thrown.
	PenaltyOnField bool `json:"penaltyOnField" api:"required"`
	// Whether a timeout followed the play.
	Timeout bool `json:"timeout" api:"required"`
	// Whether officials were reviewing the play.
	UnderReview bool `json:"underReview" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CalledBack     respjson.Field
		PenaltyOnField respjson.Field
		Timeout        respjson.Field
		UnderReview    respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FootballPlayOnField) RawJSON() string { return r.JSON.raw }
func (r *FootballPlayOnField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FootballPlayDetailPlayer struct {
	// The player, a key of players.
	PlayerID string `json:"playerId" api:"required"`
	// The player's role, such as Pass, Receive or Tackle.
	PlayerRole string `json:"playerRole" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PlayerID    respjson.Field
		PlayerRole  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FootballPlayDetailPlayer) RawJSON() string { return r.JSON.raw }
func (r *FootballPlayDetailPlayer) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The penalty this step called, or null.
type FootballPlayDetailPenalty struct {
	// The penalty's name.
	Description string `json:"description" api:"required"`
	// What the penalty does to the down.
	NextDown string `json:"nextDown" api:"required"`
	// Whether the penalty was accepted or declined.
	Result string `json:"result" api:"required"`
	// The penalized team, a key of teams.
	TeamID string `json:"teamId" api:"required"`
	// The penalty's yards.
	Yards int64 `json:"yards" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Description respjson.Field
		NextDown    respjson.Field
		Result      respjson.Field
		TeamID      respjson.Field
		Yards       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FootballPlayDetailPenalty) RawJSON() string { return r.JSON.raw }
func (r *FootballPlayDetailPenalty) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// 1 step of a play, such as a pass, a tackle or a penalty.
type FootballPlayDetail struct {
	// The step's place in the play, from 1.
	Sequence int64 `json:"sequence" api:"required"`
	// X's kind of step, such as Pass, Tackle or Penalty.
	Category string `json:"category" api:"required"`
	// What happened in the step.
	Description string `json:"description" api:"required"`
	// The direction X gives for the step.
	Direction string `json:"direction" api:"required"`
	// The yards the step gained or lost.
	Yards int64 `json:"yards" api:"required"`
	// The players in the step & their roles.
	Players []FootballPlayDetailPlayer `json:"players" api:"required"`
	// The penalty this step called, or null.
	Penalty FootballPlayDetailPenalty `json:"penalty" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Sequence    respjson.Field
		Category    respjson.Field
		Description respjson.Field
		Direction   respjson.Field
		Yards       respjson.Field
		Players     respjson.Field
		Penalty     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FootballPlayDetail) RawJSON() string { return r.JSON.raw }
func (r *FootballPlayDetail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The post X links to the play, or null.
type FootballPlayPost struct {
	// The post's ID.
	ID string `json:"id" api:"required"`
	// The username of the post's author, when X sends it.
	Username string `json:"username" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Username    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FootballPlayPost) RawJSON() string { return r.JSON.raw }
func (r *FootballPlayPost) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A play, with its situation, steps & linked posts.
type FootballPlay struct {
	// X's ID for the play within the game.
	ID string `json:"id" api:"required"`
	// The play's place in its drive, or null when X sends no digits.
	Sequence int64 `json:"sequence" api:"required"`
	// The game clock when the play started, as HH:MM:SS.
	Clock string `json:"clock" api:"required"`
	// X's kind of play, such as Pass, Rush, Punt or Conversion.
	Type string `json:"type" api:"required"`
	// Where the play went, such as Left or Middle.
	Direction string `json:"direction" api:"required"`
	// The play as the official play-by-play writes it.
	Description string `json:"description" api:"required"`
	// Whether the play scored points.
	Scoring bool `json:"scoring" api:"required"`
	// The home team's points after the play.
	HomePoints int64 `json:"homePoints" api:"required"`
	// The away team's points after the play.
	AwayPoints int64 `json:"awayPoints" api:"required"`
	// The points a scoring play added, or null.
	Score FootballPlayScore `json:"score" api:"required"`
	// Down & distance when the play started, or null.
	Start FootballPlayStart `json:"start" api:"required"`
	// Flags X shows on the field for the play, or null.
	OnField FootballPlayOnField `json:"onField" api:"required"`
	// The play's steps, in order.
	Details []FootballPlayDetail `json:"details" api:"required"`
	// The post X links to the play, or null.
	Post FootballPlayPost `json:"post" api:"required"`
	// The ID of a post with the play's video.
	VideoPostID string `json:"videoPostId" api:"required"`
	// The ID of a reply with the play's video.
	VideoReplyPostID string `json:"videoReplyPostId" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		Sequence         respjson.Field
		Clock            respjson.Field
		Type             respjson.Field
		Direction        respjson.Field
		Description      respjson.Field
		Scoring          respjson.Field
		HomePoints       respjson.Field
		AwayPoints       respjson.Field
		Score            respjson.Field
		Start            respjson.Field
		OnField          respjson.Field
		Details          respjson.Field
		Post             respjson.Field
		VideoPostID      respjson.Field
		VideoReplyPostID respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FootballPlay) RawJSON() string { return r.JSON.raw }
func (r *FootballPlay) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Photos of the player. Any of them can be null.
type FootballPlayerHeadshots struct {
	// The photo X shows on the player's card.
	Original string `json:"original" api:"required"`
	// The large photo.
	Large string `json:"large" api:"required"`
	// The medium photo.
	Medium string `json:"medium" api:"required"`
	// The small photo.
	Small string `json:"small" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Original    respjson.Field
		Large       respjson.Field
		Medium      respjson.Field
		Small       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FootballPlayerHeadshots) RawJSON() string { return r.JSON.raw }
func (r *FootballPlayerHeadshots) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The player's X account, or null without one.
type FootballPlayerXUser struct {
	// The account's user ID.
	ID string `json:"id" api:"required"`
	// The account's username, without @.
	Username string `json:"username" api:"required"`
	// Whether the account has a blue check.
	BlueVerified bool `json:"blueVerified" api:"required"`
	// Whether the account protects its posts. Only rosters carry it.
	Protected bool `json:"protected" api:"required"`
	// Whether the account can offer subscriptions. Only rosters carry it.
	SuperFollowEligible bool `json:"superFollowEligible" api:"required"`
	// The label X highlights on the account, such as BusinessLabel. Only rosters
	// carry it.
	LabelType string `json:"labelType" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                  respjson.Field
		Username            respjson.Field
		BlueVerified        respjson.Field
		Protected           respjson.Field
		SuperFollowEligible respjson.Field
		LabelType           respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FootballPlayerXUser) RawJSON() string { return r.JSON.raw }
func (r *FootballPlayerXUser) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A football player, as X's sports pages show them.
type FootballPlayer struct {
	// The player ID that plays & stats refer to.
	ID string `json:"id" api:"required"`
	// The player's team, a key of teams.
	TeamID string `json:"teamId" api:"required"`
	// The player's full name.
	Name string `json:"name" api:"required"`
	// The name play descriptions use.
	ShortName string `json:"shortName" api:"required"`
	// The jersey number, as text.
	Jersey string `json:"jersey" api:"required"`
	// The position abbreviation, such as QB or DE.
	Position string `json:"position" api:"required"`
	// The college the player came from.
	College string `json:"college" api:"required"`
	// Where the player was born.
	BirthPlace string `json:"birthPlace" api:"required"`
	// The birth date, as YYYY-MM-DD.
	BirthDate string `json:"birthDate" api:"required"`
	// Photos of the player. Any of them can be null.
	Headshots FootballPlayerHeadshots `json:"headshots" api:"required"`
	// The player's X account, or null without one.
	XUser FootballPlayerXUser `json:"xUser" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		TeamID      respjson.Field
		Name        respjson.Field
		ShortName   respjson.Field
		Jersey      respjson.Field
		Position    respjson.Field
		College     respjson.Field
		BirthPlace  respjson.Field
		BirthDate   respjson.Field
		Headshots   respjson.Field
		XUser       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FootballPlayer) RawJSON() string { return r.JSON.raw }
func (r *FootballPlayer) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// 1 game's scoreboard.
type FootballScoreboard struct {
	// The game ID from your request.
	ID string `json:"id" api:"required"`
	// The league from your request.
	League string `json:"league" api:"required"`
	// X's ID for the season. NFL games have one.
	SeasonID string `json:"seasonId" api:"required"`
	// The part of the season, such as RegularSeason.
	SeasonType string `json:"seasonType" api:"required"`
	// The week of the season.
	Week int64 `json:"week" api:"required"`
	// The kickoff time in UTC, or null when X sends none Xquik can read.
	StartsAt time.Time `json:"startsAt" api:"required" format:"date-time"`
	// The stadium.
	Venue string `json:"venue" api:"required"`
	// X's game status, such as Scheduled, InProgress or Closed.
	Status string `json:"status" api:"required"`
	// The game clock, as X writes it.
	Clock string `json:"clock" api:"required"`
	// The current or last quarter.
	Quarter int64 `json:"quarter" api:"required"`
	// Down & distance now, or null before kickoff.
	Situation FootballScoreboardSituation `json:"situation" api:"required"`
	// A team's side of a scoreboard.
	Home FootballSide `json:"home" api:"required"`
	// A team's side of a scoreboard.
	Away FootballSide `json:"away" api:"required"`
	// Points by period, in order.
	Periods []FootballPeriod `json:"periods" api:"required"`
	// The latest play, or null before kickoff.
	LastPlay FootballPlay `json:"lastPlay" api:"required"`
	// The teams this answer names, by team ID.
	Teams map[string]SportsTeam `json:"teams" api:"required"`
	// The players this answer names, by player ID.
	Players map[string]FootballPlayer `json:"players" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		League      respjson.Field
		SeasonID    respjson.Field
		SeasonType  respjson.Field
		Week        respjson.Field
		StartsAt    respjson.Field
		Venue       respjson.Field
		Status      respjson.Field
		Clock       respjson.Field
		Quarter     respjson.Field
		Situation   respjson.Field
		Home        respjson.Field
		Away        respjson.Field
		Periods     respjson.Field
		LastPlay    respjson.Field
		Teams       respjson.Field
		Players     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FootballScoreboard) RawJSON() string { return r.JSON.raw }
func (r *FootballScoreboard) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Third downs tried & converted.
type FootballTeamTotalsThirdDowns struct {
	// Third downs tried.
	Attempts int64 `json:"attempts" api:"required"`
	// Third downs converted.
	Successes int64 `json:"successes" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Attempts    respjson.Field
		Successes   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FootballTeamTotalsThirdDowns) RawJSON() string { return r.JSON.raw }
func (r *FootballTeamTotalsThirdDowns) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Fumble stats.
type FootballFumbleStats struct {
	// Fumbles the other team recovered.
	Lost int64 `json:"lost" api:"required"`
	// Touchdowns scored after recovering the other team's fumble.
	OpponentRecoveryTouchdowns int64 `json:"opponentRecoveryTouchdowns" api:"required"`
	// Touchdowns scored after recovering an own fumble.
	OwnRecoveryTouchdowns int64 `json:"ownRecoveryTouchdowns" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Lost                       respjson.Field
		OpponentRecoveryTouchdowns respjson.Field
		OwnRecoveryTouchdowns      respjson.Field
		ExtraFields                map[string]respjson.Field
		raw                        string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FootballFumbleStats) RawJSON() string { return r.JSON.raw }
func (r *FootballFumbleStats) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Passing stats.
type FootballPassingStats struct {
	// Passes thrown.
	Attempts int64 `json:"attempts" api:"required"`
	// Yards per pass.
	AvgYards float64 `json:"avgYards" api:"required"`
	// Passes caught.
	Completions int64 `json:"completions" api:"required"`
	// Passes intercepted.
	Interceptions int64 `json:"interceptions" api:"required"`
	// Passing yards minus sack yards.
	NetYards int64 `json:"netYards" api:"required"`
	// The passer rating.
	Rating float64 `json:"rating" api:"required"`
	// Yards lost to sacks.
	SackYards int64 `json:"sackYards" api:"required"`
	// Times sacked.
	Sacks int64 `json:"sacks" api:"required"`
	// Passing touchdowns.
	Touchdowns int64 `json:"touchdowns" api:"required"`
	// Passing yards.
	Yards int64 `json:"yards" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Attempts      respjson.Field
		AvgYards      respjson.Field
		Completions   respjson.Field
		Interceptions respjson.Field
		NetYards      respjson.Field
		Rating        respjson.Field
		SackYards     respjson.Field
		Sacks         respjson.Field
		Touchdowns    respjson.Field
		Yards         respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FootballPassingStats) RawJSON() string { return r.JSON.raw }
func (r *FootballPassingStats) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Penalties against the team.
type FootballTeamTotalsPenalties struct {
	// Penalties.
	Count int64 `json:"count" api:"required"`
	// Penalty yards.
	Yards int64 `json:"yards" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Yards       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FootballTeamTotalsPenalties) RawJSON() string { return r.JSON.raw }
func (r *FootballTeamTotalsPenalties) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Rushing stats.
type FootballRushingStats struct {
	// Rushes.
	Attempts int64 `json:"attempts" api:"required"`
	// Yards per rush.
	AvgYards float64 `json:"avgYards" api:"required"`
	// The longest rush, in yards.
	Longest int64 `json:"longest" api:"required"`
	// Rushing touchdowns.
	Touchdowns int64 `json:"touchdowns" api:"required"`
	// Rushing yards.
	Yards int64 `json:"yards" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Attempts    respjson.Field
		AvgYards    respjson.Field
		Longest     respjson.Field
		Touchdowns  respjson.Field
		Yards       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FootballRushingStats) RawJSON() string { return r.JSON.raw }
func (r *FootballRushingStats) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A team's totals for 1 game.
type FootballTeamTotals struct {
	// First downs.
	FirstDowns int64 `json:"firstDowns" api:"required"`
	// Third downs tried & converted.
	ThirdDowns FootballTeamTotalsThirdDowns `json:"thirdDowns" api:"required"`
	// The team's fumbles, or null.
	Fumbles FootballFumbleStats `json:"fumbles" api:"required"`
	// The team's passing, or null.
	Passing FootballPassingStats `json:"passing" api:"required"`
	// Penalties against the team.
	Penalties FootballTeamTotalsPenalties `json:"penalties" api:"required"`
	// The team's rushing, or null.
	Rushing FootballRushingStats `json:"rushing" api:"required"`
	// Time with the ball, in seconds.
	TimeOfPossessionSeconds int64 `json:"timeOfPossessionSeconds" api:"required"`
	// Total yards.
	TotalYards int64 `json:"totalYards" api:"required"`
	// Turnovers.
	Turnovers int64 `json:"turnovers" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FirstDowns              respjson.Field
		ThirdDowns              respjson.Field
		Fumbles                 respjson.Field
		Passing                 respjson.Field
		Penalties               respjson.Field
		Rushing                 respjson.Field
		TimeOfPossessionSeconds respjson.Field
		TotalYards              respjson.Field
		Turnovers               respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FootballTeamTotals) RawJSON() string { return r.JSON.raw }
func (r *FootballTeamTotals) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Receiving stats.
type FootballPlayerStatsReceiving struct {
	// Yards per catch.
	AvgYards float64 `json:"avgYards" api:"required"`
	// The longest catch, in yards.
	Longest int64 `json:"longest" api:"required"`
	// Catches.
	Receptions int64 `json:"receptions" api:"required"`
	// Passes thrown to the player.
	Targets int64 `json:"targets" api:"required"`
	// Receiving touchdowns.
	Touchdowns int64 `json:"touchdowns" api:"required"`
	// Receiving yards.
	Yards int64 `json:"yards" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AvgYards    respjson.Field
		Longest     respjson.Field
		Receptions  respjson.Field
		Targets     respjson.Field
		Touchdowns  respjson.Field
		Yards       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FootballPlayerStatsReceiving) RawJSON() string { return r.JSON.raw }
func (r *FootballPlayerStatsReceiving) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Defensive stats.
type FootballPlayerStatsDefense struct {
	// Solo & assisted tackles together.
	Combined float64 `json:"combined" api:"required"`
	// Fumbles forced.
	ForcedFumbles int64 `json:"forcedFumbles" api:"required"`
	// Interceptions.
	Interceptions int64 `json:"interceptions" api:"required"`
	// Passes defended.
	PassesDefended int64 `json:"passesDefended" api:"required"`
	// Sacks, with a shared sack as 0.5.
	Sacks float64 `json:"sacks" api:"required"`
	// Solo tackles.
	Tackles float64 `json:"tackles" api:"required"`
	// Tackles for a loss.
	TacklesForLoss float64 `json:"tacklesForLoss" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Combined       respjson.Field
		ForcedFumbles  respjson.Field
		Interceptions  respjson.Field
		PassesDefended respjson.Field
		Sacks          respjson.Field
		Tackles        respjson.Field
		TacklesForLoss respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FootballPlayerStatsDefense) RawJSON() string { return r.JSON.raw }
func (r *FootballPlayerStatsDefense) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Kick, punt or interception return stats.
type FootballReturnStats struct {
	// Yards per return.
	AvgYards float64 `json:"avgYards" api:"required"`
	// The longest return, in yards.
	Longest int64 `json:"longest" api:"required"`
	// Returns.
	Returns int64 `json:"returns" api:"required"`
	// Return touchdowns.
	Touchdowns int64 `json:"touchdowns" api:"required"`
	// Return yards.
	Yards int64 `json:"yards" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AvgYards    respjson.Field
		Longest     respjson.Field
		Returns     respjson.Field
		Touchdowns  respjson.Field
		Yards       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FootballReturnStats) RawJSON() string { return r.JSON.raw }
func (r *FootballReturnStats) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Punting stats.
type FootballPlayerStatsPunts struct {
	// Punts.
	Attempts int64 `json:"attempts" api:"required"`
	// Net yards per punt, after returns.
	AvgNetYards float64 `json:"avgNetYards" api:"required"`
	// Yards per punt.
	AvgYards float64 `json:"avgYards" api:"required"`
	// Punts downed inside the 20.
	Inside20 int64 `json:"inside20" api:"required"`
	// The longest punt, in yards.
	Longest int64 `json:"longest" api:"required"`
	// Punting yards.
	Yards int64 `json:"yards" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Attempts    respjson.Field
		AvgNetYards respjson.Field
		AvgYards    respjson.Field
		Inside20    respjson.Field
		Longest     respjson.Field
		Yards       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FootballPlayerStatsPunts) RawJSON() string { return r.JSON.raw }
func (r *FootballPlayerStatsPunts) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Field goal stats.
type FootballPlayerStatsFieldGoals struct {
	// Field goals tried.
	Attempts int64 `json:"attempts" api:"required"`
	// The longest made field goal, in yards.
	Longest int64 `json:"longest" api:"required"`
	// Field goals made.
	Made int64 `json:"made" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Attempts    respjson.Field
		Longest     respjson.Field
		Made        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FootballPlayerStatsFieldGoals) RawJSON() string { return r.JSON.raw }
func (r *FootballPlayerStatsFieldGoals) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Extra point kicks.
type FootballPlayerStatsExtraPoints struct {
	// Extra points tried.
	Attempts int64 `json:"attempts" api:"required"`
	// Extra points made.
	Made int64 `json:"made" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Attempts    respjson.Field
		Made        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FootballPlayerStatsExtraPoints) RawJSON() string { return r.JSON.raw }
func (r *FootballPlayerStatsExtraPoints) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// 2-point conversions.
type FootballPlayerStatsConversions struct {
	// Defensive 2-point returns.
	DefenseSuccesses int64 `json:"defenseSuccesses" api:"required"`
	// 2-point conversions by pass.
	PassSuccesses int64 `json:"passSuccesses" api:"required"`
	// 2-point conversions by rush.
	RushSuccesses int64 `json:"rushSuccesses" api:"required"`
	// 2-point conversions after a turnover.
	TurnoverSuccesses int64 `json:"turnoverSuccesses" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DefenseSuccesses  respjson.Field
		PassSuccesses     respjson.Field
		RushSuccesses     respjson.Field
		TurnoverSuccesses respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FootballPlayerStatsConversions) RawJSON() string { return r.JSON.raw }
func (r *FootballPlayerStatsConversions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// 1 player's stats for 1 game. Only the stat groups X sent are present.
type FootballPlayerStats struct {
	// The player these stats belong to, a key of players.
	PlayerID string `json:"playerId" api:"required"`
	// Passing stats.
	Passing FootballPassingStats `json:"passing"`
	// Rushing stats.
	Rushing FootballRushingStats `json:"rushing"`
	// Receiving stats.
	Receiving FootballPlayerStatsReceiving `json:"receiving"`
	// Defensive stats.
	Defense FootballPlayerStatsDefense `json:"defense"`
	// Fumble stats.
	Fumbles FootballFumbleStats `json:"fumbles"`
	// Kick, punt or interception return stats.
	KickReturns FootballReturnStats `json:"kickReturns"`
	// Kick, punt or interception return stats.
	PuntReturns FootballReturnStats `json:"puntReturns"`
	// Kick, punt or interception return stats.
	InterceptionReturns FootballReturnStats `json:"interceptionReturns"`
	// Punting stats.
	Punts FootballPlayerStatsPunts `json:"punts"`
	// Field goal stats.
	FieldGoals FootballPlayerStatsFieldGoals `json:"fieldGoals"`
	// Extra point kicks.
	ExtraPoints FootballPlayerStatsExtraPoints `json:"extraPoints"`
	// 2-point conversions.
	Conversions FootballPlayerStatsConversions `json:"conversions"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PlayerID            respjson.Field
		Passing             respjson.Field
		Rushing             respjson.Field
		Receiving           respjson.Field
		Defense             respjson.Field
		Fumbles             respjson.Field
		KickReturns         respjson.Field
		PuntReturns         respjson.Field
		InterceptionReturns respjson.Field
		Punts               respjson.Field
		FieldGoals          respjson.Field
		ExtraPoints         respjson.Field
		Conversions         respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FootballPlayerStats) RawJSON() string { return r.JSON.raw }
func (r *FootballPlayerStats) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// 1 team's totals & player stats.
type FootballTeamBox struct {
	// The team these totals belong to, a key of teams.
	TeamID string `json:"teamId" api:"required"`
	// The team's totals, or null before kickoff.
	Totals FootballTeamTotals `json:"totals" api:"required"`
	// Each player with stats, 1 entry per player.
	Players []FootballPlayerStats `json:"players" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		TeamID      respjson.Field
		Totals      respjson.Field
		Players     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FootballTeamBox) RawJSON() string { return r.JSON.raw }
func (r *FootballTeamBox) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// 1 game's team totals & player stats.
type FootballBoxScore struct {
	// The ID of the game this box score covers.
	ID string `json:"id" api:"required"`
	// The league of the game this box score covers.
	League string `json:"league" api:"required"`
	// 1 team's totals & player stats.
	Home FootballTeamBox `json:"home" api:"required"`
	// 1 team's totals & player stats.
	Away FootballTeamBox `json:"away" api:"required"`
	// The teams this box score names, by team ID.
	Teams map[string]SportsTeam `json:"teams" api:"required"`
	// The players this box score names, by player ID.
	Players map[string]FootballPlayer `json:"players" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		League      respjson.Field
		Home        respjson.Field
		Away        respjson.Field
		Teams       respjson.Field
		Players     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FootballBoxScore) RawJSON() string { return r.JSON.raw }
func (r *FootballBoxScore) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A drive, with its result & plays.
type FootballDrive struct {
	// X's ID for the drive within the game.
	ID string `json:"id" api:"required"`
	// The drive's place in the game, from 1.
	Sequence int64 `json:"sequence" api:"required"`
	// The team that drove. Home or Away in NFL games, the team's alternateId in
	// college games.
	Offense string `json:"offense" api:"required"`
	// The team on defense. Home or Away in NFL games, the team's alternateId in
	// college games.
	Defense string `json:"defense" api:"required"`
	// The game clock when the drive started, as HH:MM:SS.
	StartClock string `json:"startClock" api:"required"`
	// The game clock when the drive ended, as HH:MM:SS.
	EndClock string `json:"endClock" api:"required"`
	// How the drive started, such as Kickoff or Punt.
	StartReason string `json:"startReason" api:"required"`
	// How the drive ended, such as Touchdown or Punt.
	EndReason string `json:"endReason" api:"required"`
	// The drive's plays.
	PlayCount int64 `json:"playCount" api:"required"`
	// The first downs the drive gained.
	FirstDowns int64 `json:"firstDowns" api:"required"`
	// The yards the drive gained.
	NetYards int64 `json:"netYards" api:"required"`
	// Whether the drive scored.
	Scoring bool `json:"scoring" api:"required"`
	// The drive's plays, in order.
	Plays []FootballPlay `json:"plays" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Sequence    respjson.Field
		Offense     respjson.Field
		Defense     respjson.Field
		StartClock  respjson.Field
		EndClock    respjson.Field
		StartReason respjson.Field
		EndReason   respjson.Field
		PlayCount   respjson.Field
		FirstDowns  respjson.Field
		NetYards    respjson.Field
		Scoring     respjson.Field
		Plays       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FootballDrive) RawJSON() string { return r.JSON.raw }
func (r *FootballDrive) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A quarter & its drives.
type FootballQuarter struct {
	// X's ID for the quarter within the game.
	ID string `json:"id" api:"required"`
	// The quarter's number.
	Number int64 `json:"number" api:"required"`
	// The quarter's place in the game, from 1.
	Sequence int64 `json:"sequence" api:"required"`
	// The quarter's drives, in order.
	Drives []FootballDrive `json:"drives" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Number      respjson.Field
		Sequence    respjson.Field
		Drives      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FootballQuarter) RawJSON() string { return r.JSON.raw }
func (r *FootballQuarter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// 1 game's drives & plays, by quarter.
type FootballPlays struct {
	// The ID of the game these plays come from.
	ID string `json:"id" api:"required"`
	// The league of the game these plays come from.
	League string `json:"league" api:"required"`
	// The home team, a key of teams.
	HomeTeamID string `json:"homeTeamId" api:"required"`
	// The away team, a key of teams.
	AwayTeamID string `json:"awayTeamId" api:"required"`
	// The game's quarters, in order. Empty before kickoff.
	Quarters []FootballQuarter `json:"quarters" api:"required"`
	// The teams these plays name, by team ID.
	Teams map[string]SportsTeam `json:"teams" api:"required"`
	// The players these plays name, by player ID.
	Players map[string]FootballPlayer `json:"players" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		League      respjson.Field
		HomeTeamID  respjson.Field
		AwayTeamID  respjson.Field
		Quarters    respjson.Field
		Teams       respjson.Field
		Players     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FootballPlays) RawJSON() string { return r.JSON.raw }
func (r *FootballPlays) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// People watching the game on X at X's last count. Null before X counts any, as
// before kickoff.
type SportsGameAudienceViewers struct {
	// How many people watched the game on X.
	Count int64 `json:"count" api:"required"`
	// When X counted them.
	AsOf time.Time `json:"asOf" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		AsOf        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SportsGameAudienceViewers) RawJSON() string { return r.JSON.raw }
func (r *SportsGameAudienceViewers) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// 1 team's picks.
type SportsGameAudienceTeamPick struct {
	// The team ID, a key in teams.
	TeamID string `json:"teamId" api:"required"`
	// How many people picked this team.
	Count int64 `json:"count" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		TeamID      respjson.Field
		Count       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SportsGameAudienceTeamPick) RawJSON() string { return r.JSON.raw }
func (r *SportsGameAudienceTeamPick) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// How many people on X picked a team to root for.
type SportsGameAudienceTeamPicks struct {
	// Picks for every team in the game.
	Total int64 `json:"total" api:"required"`
	// Picks per team.
	Teams []SportsGameAudienceTeamPick `json:"teams" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Total       respjson.Field
		Teams       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SportsGameAudienceTeamPicks) RawJSON() string { return r.JSON.raw }
func (r *SportsGameAudienceTeamPicks) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Time left until the chat opens. Null while it is open & after the game.
type SportsGameAudienceChatOpensIn struct {
	// Whole days left.
	Days int64 `json:"days" api:"required"`
	// Hours left after the days.
	Hours int64 `json:"hours" api:"required"`
	// Minutes left after the hours.
	Minutes int64 `json:"minutes" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Days        respjson.Field
		Hours       respjson.Field
		Minutes     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SportsGameAudienceChatOpensIn) RawJSON() string { return r.JSON.raw }
func (r *SportsGameAudienceChatOpensIn) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The game's live chat on X.
type SportsGameAudienceChat struct {
	// Whether the chat is open now.
	Available bool `json:"available" api:"required"`
	// X's ID for the game's chat room.
	BroadcastID string `json:"broadcastId" api:"required"`
	// Time left until the chat opens. Null while it is open & after the game.
	OpensIn SportsGameAudienceChatOpensIn `json:"opensIn" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Available   respjson.Field
		BroadcastID respjson.Field
		OpensIn     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SportsGameAudienceChat) RawJSON() string { return r.JSON.raw }
func (r *SportsGameAudienceChat) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Who watches & roots for 1 game on X.
type SportsGameAudience struct {
	// The ID of the game, from your request.
	ID string `json:"id" api:"required"`
	// The game's league, from your request.
	League string `json:"league" api:"required"`
	// People watching the game on X at X's last count. Null before X counts any,
	// as before kickoff.
	Viewers SportsGameAudienceViewers `json:"viewers" api:"required"`
	// How many people on X picked a team to root for.
	TeamPicks SportsGameAudienceTeamPicks `json:"teamPicks" api:"required"`
	// The game's live chat on X.
	Chat SportsGameAudienceChat `json:"chat" api:"required"`
	// The game's teams, by team ID. X's team picker sends no short name, alternate
	// ID, X account or record, so those are null.
	Teams map[string]SportsTeam `json:"teams" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		League      respjson.Field
		Viewers     respjson.Field
		TeamPicks   respjson.Field
		Chat        respjson.Field
		Teams       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SportsGameAudience) RawJSON() string { return r.JSON.raw }
func (r *SportsGameAudience) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// FootballRosterTeam is 1 team's roster.
type FootballRosterTeam struct {
	// The team this roster belongs to, a key of teams.
	TeamID string `json:"teamId" api:"required"`
	// The team's players in X's order, keys of players.
	PlayerIDs []string `json:"playerIds" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		TeamID      respjson.Field
		PlayerIDs   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// RawJSON returns the unmodified JSON received from the API.
func (r FootballRosterTeam) RawJSON() string { return r.JSON.raw }
func (r *FootballRosterTeam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// FootballRoster holds both teams' rosters for 1 NFL game.
type FootballRoster struct {
	// The ID of the game these rosters cover.
	ID string `json:"id" api:"required"`
	// The league of the game these rosters cover.
	League string `json:"league" api:"required"`
	// Each team's roster, in X's order.
	Rosters []FootballRosterTeam `json:"rosters" api:"required"`
	// The 2 teams, by team ID.
	Teams map[string]SportsTeam `json:"teams" api:"required"`
	// Every rostered player, by player ID. Rosters carry no college or birth details
	// & only a medium headshot, so the rest are null.
	Players map[string]FootballPlayer `json:"players" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		League      respjson.Field
		Rosters     respjson.Field
		Teams       respjson.Field
		Players     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// RawJSON returns the unmodified JSON received from the API.
func (r FootballRoster) RawJSON() string { return r.JSON.raw }
func (r *FootballRoster) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// FootballPlayerProp is 1 player's line & odds in a prop category.
type FootballPlayerProp struct {
	// The player the line is for, a key of players.
	PlayerID string `json:"playerId" api:"required"`
	// The line's place in its category, from 1.
	Rank float64 `json:"rank" api:"required"`
	// The line the odds price, such as 248.5 yards. Null when the prop has no line,
	// such as the first touchdown scorer.
	Line float64 `json:"line" api:"required"`
	// The American odds X shows first, such as -115 or 240.
	Odds float64 `json:"odds" api:"required"`
	// The American odds X shows second. Null for a prop with 1 price.
	AlternateOdds float64 `json:"alternateOdds" api:"required"`
	// Whether the sportsbook takes bets on the line.
	OpenForBetting bool `json:"openForBetting" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PlayerID       respjson.Field
		Rank           respjson.Field
		Line           respjson.Field
		Odds           respjson.Field
		AlternateOdds  respjson.Field
		OpenForBetting respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// RawJSON returns the unmodified JSON received from the API.
func (r FootballPlayerProp) RawJSON() string { return r.JSON.raw }
func (r *FootballPlayerProp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// FootballPropCategory is 1 kind of player prop & its lines.
type FootballPropCategory struct {
	// X's key for the prop, such as passing_yards.
	Key string `json:"key" api:"required"`
	// The name X shows for the prop.
	Name string `json:"name" api:"required"`
	// The category's place in X's order, from 1.
	Rank float64 `json:"rank" api:"required"`
	// Each player's line, in X's order.
	Props []FootballPlayerProp `json:"props" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Key         respjson.Field
		Name        respjson.Field
		Rank        respjson.Field
		Props       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// RawJSON returns the unmodified JSON received from the API.
func (r FootballPropCategory) RawJSON() string { return r.JSON.raw }
func (r *FootballPropCategory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// FootballProps holds the player props X shows for 1 NFL game.
type FootballProps struct {
	// The ID of the game these props cover.
	ID string `json:"id" api:"required"`
	// The league of the game these props cover.
	League string `json:"league" api:"required"`
	// The sportsbook page X links for these props.
	SportsbookURL string `json:"sportsbookUrl" api:"required" format:"uri"`
	// Each kind of prop, such as passing yards, in X's order. A game without props
	// has none.
	Categories []FootballPropCategory `json:"categories" api:"required"`
	// The players these props name, by player ID. X sends only each player's name,
	// jersey & team, so other fields are null.
	Players map[string]FootballPlayer `json:"players" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		League        respjson.Field
		SportsbookURL respjson.Field
		Categories    respjson.Field
		Players       respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// RawJSON returns the unmodified JSON received from the API.
func (r FootballProps) RawJSON() string { return r.JSON.raw }
func (r *FootballProps) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SportsGamePosts is a page of 1 game's posts, in Xquik's post shape. Pass
// next_cursor as cursor for the next page.
type SportsGamePosts struct {
	// The X search that finds this game's posts for the order asked. X builds it from
	// the teams, hashtags, accounts & star players, with the game's time window. Send
	// it to GET /x/tweets/search for more filters. Null when X sends a timeline ID
	// Xquik cannot read.
	Query string `json:"query" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Query       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.PaginatedTweets
}

// RawJSON returns the unmodified JSON received from the API.
func (r SportsGamePosts) RawJSON() string { return r.JSON.raw }
func (r *SportsGamePosts) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type XSportListGamesParams struct {
	// Start of the window. Takes an ISO 8601 date or time, such as 2026-10-01, or Unix
	// seconds. A time without a zone reads as UTC. Defaults to 7 days before to. Alias
	// sinceTime.
	From param.Opt[string] `query:"from,omitzero" json:"-"`
	// End of the window. Takes the forms of from. Defaults to 7 days after from. Alias
	// untilTime.
	To param.Opt[string] `query:"to,omitzero" json:"-"`
	// Keeps only 1 team's games, by team ID. Each game lists its teamIds. Aliases
	// teamId & team_id.
	Team param.Opt[string] `query:"team,omitzero" json:"-"`
	// Fields to keep in each row, comma-separated. Use dots for nested fields, such as
	// author.username. A list keeps the field in each item. Case & underscores do not
	// matter. A row that lacks a field leaves it out. Cursors, counts & diagnostics
	// always stay. An unknown name answers 400 unknown_field with the valid names,
	// before any read. The price stays the same. SDK row types then hold only the
	// named fields.
	Fields param.Opt[string] `query:"fields,omitzero" json:"-"`
	// Fields to drop from each row, comma-separated. It takes the same names as
	// fields. With fields, it drops from the kept fields.
	Omit param.Opt[string] `query:"omit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [XSportListGamesParams]'s query parameters as `url.Values`.
func (r XSportListGamesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type XSportGetScoreboardParams struct {
	// nfl for the NFL or ncaafb for college football, in any case. Another value gets
	// a free 400 `invalid_league` before X is read.
	League string `path:"league" api:"required" json:"-"`
	// Fields to keep in each row, comma-separated. Use dots for nested fields, such as
	// author.username. A list keeps the field in each item. Case & underscores do not
	// matter. A row that lacks a field leaves it out. Cursors, counts & diagnostics
	// always stay. An unknown name answers 400 unknown_field with the valid names,
	// before any read. The price stays the same. SDK row types then hold only the
	// named fields.
	Fields param.Opt[string] `query:"fields,omitzero" json:"-"`
	// Fields to drop from each row, comma-separated. It takes the same names as
	// fields. With fields, it drops from the kept fields.
	Omit param.Opt[string] `query:"omit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [XSportGetScoreboardParams]'s query parameters as `url.Values`.
func (r XSportGetScoreboardParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type XSportGetBoxScoreParams struct {
	// nfl for the NFL or ncaafb for college football, in any case. Another value gets
	// a free 400 `invalid_league` before X is read.
	League string `path:"league" api:"required" json:"-"`
	// Fields to keep in each row, comma-separated. Use dots for nested fields, such as
	// author.username. A list keeps the field in each item. Case & underscores do not
	// matter. A row that lacks a field leaves it out. Cursors, counts & diagnostics
	// always stay. An unknown name answers 400 unknown_field with the valid names,
	// before any read. The price stays the same. SDK row types then hold only the
	// named fields.
	Fields param.Opt[string] `query:"fields,omitzero" json:"-"`
	// Fields to drop from each row, comma-separated. It takes the same names as
	// fields. With fields, it drops from the kept fields.
	Omit param.Opt[string] `query:"omit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [XSportGetBoxScoreParams]'s query parameters as `url.Values`.
func (r XSportGetBoxScoreParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type XSportGetPlaysParams struct {
	// nfl for the NFL or ncaafb for college football, in any case. Another value gets
	// a free 400 `invalid_league` before X is read.
	League string `path:"league" api:"required" json:"-"`
	// Fields to keep in each row, comma-separated. Use dots for nested fields, such as
	// author.username. A list keeps the field in each item. Case & underscores do not
	// matter. A row that lacks a field leaves it out. Cursors, counts & diagnostics
	// always stay. An unknown name answers 400 unknown_field with the valid names,
	// before any read. The price stays the same. SDK row types then hold only the
	// named fields.
	Fields param.Opt[string] `query:"fields,omitzero" json:"-"`
	// Fields to drop from each row, comma-separated. It takes the same names as
	// fields. With fields, it drops from the kept fields.
	Omit param.Opt[string] `query:"omit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [XSportGetPlaysParams]'s query parameters as `url.Values`.
func (r XSportGetPlaysParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type XSportGetAudienceParams struct {
	// nfl for the NFL or ncaafb for college football, in any case. Another value gets
	// a free 400 `invalid_league` before X is read.
	League string `path:"league" api:"required" json:"-"`
	// Fields to keep in each row, comma-separated. Use dots for nested fields, such as
	// author.username. A list keeps the field in each item. Case & underscores do not
	// matter. A row that lacks a field leaves it out. Cursors, counts & diagnostics
	// always stay. An unknown name answers 400 unknown_field with the valid names,
	// before any read. The price stays the same. SDK row types then hold only the
	// named fields.
	Fields param.Opt[string] `query:"fields,omitzero" json:"-"`
	// Fields to drop from each row, comma-separated. It takes the same names as
	// fields. With fields, it drops from the kept fields.
	Omit param.Opt[string] `query:"omit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [XSportGetAudienceParams]'s query parameters as `url.Values`.
func (r XSportGetAudienceParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type XSportGetRosterParams struct {
	// nfl, in any case. X has no college rosters or props, so ncaafb gets a free 400
	// `unsupported_league`. Another value gets a free 400 `invalid_league`. Both come
	// before X is read.
	League string `path:"league" api:"required" json:"-"`
	// Fields to keep in each row, comma-separated. Use dots for nested fields, such as
	// author.username. A list keeps the field in each item. Case & underscores do not
	// matter. A row that lacks a field leaves it out. Cursors, counts & diagnostics
	// always stay. An unknown name answers 400 unknown_field with the valid names,
	// before any read. The price stays the same. SDK row types then hold only the
	// named fields.
	Fields param.Opt[string] `query:"fields,omitzero" json:"-"`
	// Fields to drop from each row, comma-separated. It takes the same names as
	// fields. With fields, it drops from the kept fields.
	Omit param.Opt[string] `query:"omit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [XSportGetRosterParams]'s query parameters as `url.Values`.
func (r XSportGetRosterParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type XSportGetPlayerPropsParams struct {
	// nfl, in any case. X has no college rosters or props, so ncaafb gets a free 400
	// `unsupported_league`. Another value gets a free 400 `invalid_league`. Both come
	// before X is read.
	League string `path:"league" api:"required" json:"-"`
	// Fields to keep in each row, comma-separated. Use dots for nested fields, such as
	// author.username. A list keeps the field in each item. Case & underscores do not
	// matter. A row that lacks a field leaves it out. Cursors, counts & diagnostics
	// always stay. An unknown name answers 400 unknown_field with the valid names,
	// before any read. The price stays the same. SDK row types then hold only the
	// named fields.
	Fields param.Opt[string] `query:"fields,omitzero" json:"-"`
	// Fields to drop from each row, comma-separated. It takes the same names as
	// fields. With fields, it drops from the kept fields.
	Omit param.Opt[string] `query:"omit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [XSportGetPlayerPropsParams]'s query parameters as
// `url.Values`.
func (r XSportGetPlayerPropsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type XSportGetPostsParams struct {
	// nfl for the NFL or ncaafb for college football, in any case. Another value gets
	// a free 400 `invalid_league` before X is read.
	League string `path:"league" api:"required" json:"-"`
	// Pagination cursor for collection results. Also read as next_cursor, next_token,
	// pagination_token or continuation_token.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Fields to keep in each row, comma-separated. Use dots for nested fields, such as
	// author.username. A list keeps the field in each item. Case & underscores do not
	// matter. A row that lacks a field leaves it out. Cursors, counts & diagnostics
	// always stay. An unknown name answers 400 unknown_field with the valid names,
	// before any read. The price stays the same. SDK row types then hold only the
	// named fields.
	Fields param.Opt[string] `query:"fields,omitzero" json:"-"`
	// Fields to drop from each row, comma-separated. It takes the same names as
	// fields. With fields, it drops from the kept fields.
	Omit param.Opt[string] `query:"omit,omitzero" json:"-"`
	// Items per page (1-100, default 20). Source, filters or credits can return fewer.
	// Follow next_cursor for more pages. Its aliases are limit, count, max_results,
	// maxItems, max_items & per_page.
	PageSize param.Opt[int64] `query:"pageSize,omitzero" json:"-"`
	// Top for the posts X ranks highest, or Latest for the newest. Any case works, &
	// so do Relevance & Recency. Alias order or sort. Send the same value with each
	// cursor.
	//
	// Any of "Top", "Latest".
	QueryType XSportGetPostsParamsQueryType `query:"queryType,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [XSportGetPostsParams]'s query parameters as `url.Values`.
func (r XSportGetPostsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// XSportGetPostsParamsQueryType picks the order of a game's posts. Top is for the
// posts X ranks highest, and Latest is for the newest.
type XSportGetPostsParamsQueryType string

const (
	XSportGetPostsParamsQueryTypeTop    XSportGetPostsParamsQueryType = "Top"
	XSportGetPostsParamsQueryTypeLatest XSportGetPostsParamsQueryType = "Latest"
)
