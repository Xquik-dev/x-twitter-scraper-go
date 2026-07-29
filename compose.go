// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package xtwitterscraper

import (
	"context"
	"encoding/json"
	"net/http"
	"slices"

	"github.com/Xquik-dev/x-twitter-scraper-go/internal/apijson"
	"github.com/Xquik-dev/x-twitter-scraper-go/internal/requestconfig"
	"github.com/Xquik-dev/x-twitter-scraper-go/option"
	"github.com/Xquik-dev/x-twitter-scraper-go/packages/param"
	"github.com/Xquik-dev/x-twitter-scraper-go/packages/respjson"
	"github.com/Xquik-dev/x-twitter-scraper-go/shared/constant"
)

// AI tweet composition, drafts, writing styles, and radar
//
// ComposeService contains methods and other services that help with interacting
// with the x-twitter-scraper API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewComposeService] method instead.
type ComposeService struct {
	options []option.RequestOption
}

// NewComposeService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewComposeService(opts ...option.RequestOption) (r ComposeService) {
	r = ComposeService{}
	r.options = opts
	return
}

// Run one step of Xquik's three-step writing workflow. Compose returns questions,
// editorial rules, and source-specific Radar recommendations. Refine returns
// goal-specific guidance. Score applies deterministic text checks. It does not
// predict reach or expose X ranking weights.
func (r *ComposeService) New(ctx context.Context, body ComposeNewParams, opts ...option.RequestOption) (res *ComposeNewResponseUnion, err error) {
	opts = slices.Concat(r.options, opts)
	path := "compose"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// ComposeNewResponseUnion contains all possible properties and values from
// [ComposeNewResponseComposePrepareResult],
// [ComposeNewResponseComposeRefineResult], [ComposeNewResponseComposeScoreResult].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ComposeNewResponseUnion struct {
	// This field is from variant [ComposeNewResponseComposePrepareResult].
	ContentRules []ComposeNewResponseComposePrepareResultContentRule `json:"contentRules"`
	// This field is from variant [ComposeNewResponseComposePrepareResult].
	EngagementMultipliers []ComposeNewResponseComposePrepareResultEngagementMultiplier `json:"engagementMultipliers"`
	// This field is from variant [ComposeNewResponseComposePrepareResult].
	EngagementVelocity string `json:"engagementVelocity"`
	// This field is from variant [ComposeNewResponseComposePrepareResult].
	FollowUpQuestions []string `json:"followUpQuestions"`
	IntentURL         string   `json:"intentUrl"`
	NextStep          string   `json:"nextStep"`
	// This field is from variant [ComposeNewResponseComposePrepareResult].
	RadarRecommendations []ComposeNewResponseComposePrepareResultRadarRecommendation `json:"radarRecommendations"`
	// This field is from variant [ComposeNewResponseComposePrepareResult].
	ScorerWeights []ComposeNewResponseComposePrepareResultScorerWeight `json:"scorerWeights"`
	// This field is from variant [ComposeNewResponseComposePrepareResult].
	Source string `json:"source"`
	// This field is from variant [ComposeNewResponseComposePrepareResult].
	TopPenalties []string `json:"topPenalties"`
	// This field is from variant [ComposeNewResponseComposePrepareResult].
	SavedStyles []ComposeNewResponseComposePrepareResultSavedStyle `json:"savedStyles"`
	// This field is from variant [ComposeNewResponseComposePrepareResult].
	StyleNote string `json:"styleNote"`
	// This field is from variant [ComposeNewResponseComposePrepareResult].
	StyleTweets []string `json:"styleTweets"`
	// This field is from variant [ComposeNewResponseComposeRefineResult].
	CompositionGuidance []string `json:"compositionGuidance"`
	// This field is from variant [ComposeNewResponseComposeRefineResult].
	ExamplePatterns []ComposeNewResponseComposeRefineResultExamplePattern `json:"examplePatterns"`
	// This field is from variant [ComposeNewResponseComposeScoreResult].
	Checklist []ComposeNewResponseComposeScoreResultChecklist `json:"checklist"`
	// This field is from variant [ComposeNewResponseComposeScoreResult].
	Passed bool `json:"passed"`
	// This field is from variant [ComposeNewResponseComposeScoreResult].
	PassedCount int64 `json:"passedCount"`
	// This field is from variant [ComposeNewResponseComposeScoreResult].
	TopSuggestion string `json:"topSuggestion"`
	// This field is from variant [ComposeNewResponseComposeScoreResult].
	TotalChecks int64 `json:"totalChecks"`
	JSON        struct {
		ContentRules          respjson.Field
		EngagementMultipliers respjson.Field
		EngagementVelocity    respjson.Field
		FollowUpQuestions     respjson.Field
		IntentURL             respjson.Field
		NextStep              respjson.Field
		RadarRecommendations  respjson.Field
		ScorerWeights         respjson.Field
		Source                respjson.Field
		TopPenalties          respjson.Field
		SavedStyles           respjson.Field
		StyleNote             respjson.Field
		StyleTweets           respjson.Field
		CompositionGuidance   respjson.Field
		ExamplePatterns       respjson.Field
		Checklist             respjson.Field
		Passed                respjson.Field
		PassedCount           respjson.Field
		TopSuggestion         respjson.Field
		TotalChecks           respjson.Field
		raw                   string
	} `json:"-"`
}

func (u ComposeNewResponseUnion) AsComposeNewResponseComposePrepareResult() (v ComposeNewResponseComposePrepareResult) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ComposeNewResponseUnion) AsComposeNewResponseComposeRefineResult() (v ComposeNewResponseComposeRefineResult) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ComposeNewResponseUnion) AsComposeNewResponseComposeScoreResult() (v ComposeNewResponseComposeScoreResult) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ComposeNewResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *ComposeNewResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComposeNewResponseComposePrepareResult struct {
	// Xquik editorial heuristics, ordered for the goal.
	ContentRules []ComposeNewResponseComposePrepareResultContentRule `json:"contentRules" api:"required"`
	// Published engagement signal names. Production multipliers are not published.
	EngagementMultipliers []ComposeNewResponseComposePrepareResultEngagementMultiplier `json:"engagementMultipliers" api:"required"`
	// Publication limit for timing and decay claims.
	EngagementVelocity string   `json:"engagementVelocity" api:"required"`
	FollowUpQuestions  []string `json:"followUpQuestions" api:"required"`
	// X post intent seeded with the topic.
	IntentURL string `json:"intentUrl" api:"required" format:"uri"`
	NextStep  string `json:"nextStep" api:"required"`
	// Sources and guidance for researching a fresh post angle.
	RadarRecommendations []ComposeNewResponseComposePrepareResultRadarRecommendation `json:"radarRecommendations" api:"required"`
	// Published signal names with unpublished weights as null.
	ScorerWeights []ComposeNewResponseComposePrepareResultScorerWeight `json:"scorerWeights" api:"required"`
	// Signal source and evidence limits.
	Source string `json:"source" api:"required"`
	// Negative engagement predictions in the public model.
	TopPenalties []string `json:"topPenalties" api:"required"`
	// Style analyses saved to the account.
	SavedStyles []ComposeNewResponseComposePrepareResultSavedStyle `json:"savedStyles"`
	// Next action when no cached style is available.
	StyleNote string `json:"styleNote"`
	// Cached examples for the requested style username.
	StyleTweets []string `json:"styleTweets"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ContentRules          respjson.Field
		EngagementMultipliers respjson.Field
		EngagementVelocity    respjson.Field
		FollowUpQuestions     respjson.Field
		IntentURL             respjson.Field
		NextStep              respjson.Field
		RadarRecommendations  respjson.Field
		ScorerWeights         respjson.Field
		Source                respjson.Field
		TopPenalties          respjson.Field
		SavedStyles           respjson.Field
		StyleNote             respjson.Field
		StyleTweets           respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ComposeNewResponseComposePrepareResult) RawJSON() string { return r.JSON.raw }
func (r *ComposeNewResponseComposePrepareResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComposeNewResponseComposePrepareResultContentRule struct {
	Rule string `json:"rule" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Rule        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ComposeNewResponseComposePrepareResultContentRule) RawJSON() string { return r.JSON.raw }
func (r *ComposeNewResponseComposePrepareResultContentRule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComposeNewResponseComposePrepareResultEngagementMultiplier struct {
	// Human-readable published signal name.
	Action     string                                   `json:"action" api:"required"`
	Multiplier constant.ProductionWeightNotPublishedByX `json:"multiplier" default:"Production weight not published by X"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Action      respjson.Field
		Multiplier  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ComposeNewResponseComposePrepareResultEngagementMultiplier) RawJSON() string {
	return r.JSON.raw
}
func (r *ComposeNewResponseComposePrepareResultEngagementMultiplier) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComposeNewResponseComposePrepareResultRadarRecommendation struct {
	// Radar endpoint for this source.
	Endpoint string `json:"endpoint" api:"required"`
	// Source-specific drafting guidance.
	Guidance string `json:"guidance" api:"required"`
	// Any of "reddit", "github", "trustmrr", "hacker_news", "google_trends",
	// "wikipedia", "polymarket".
	Source string `json:"source" api:"required"`
	// Current-topic research this source supports.
	UseFor string `json:"useFor" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Endpoint    respjson.Field
		Guidance    respjson.Field
		Source      respjson.Field
		UseFor      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ComposeNewResponseComposePrepareResultRadarRecommendation) RawJSON() string {
	return r.JSON.raw
}
func (r *ComposeNewResponseComposePrepareResultRadarRecommendation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComposeNewResponseComposePrepareResultScorerWeight struct {
	// Signal direction and publication limit.
	Context string `json:"context" api:"required"`
	// Signal name from X's public ranking repository.
	Signal string `json:"signal" api:"required"`
	// X does not publish the production weight.
	Weight any `json:"weight" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Context     respjson.Field
		Signal      respjson.Field
		Weight      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ComposeNewResponseComposePrepareResultScorerWeight) RawJSON() string { return r.JSON.raw }
func (r *ComposeNewResponseComposePrepareResultScorerWeight) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComposeNewResponseComposePrepareResultSavedStyle struct {
	TweetCount int64  `json:"tweetCount" api:"required"`
	Username   string `json:"username" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		TweetCount  respjson.Field
		Username    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ComposeNewResponseComposePrepareResultSavedStyle) RawJSON() string { return r.JSON.raw }
func (r *ComposeNewResponseComposePrepareResultSavedStyle) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComposeNewResponseComposeRefineResult struct {
	// Goal, tone, media, and editorial guidance.
	CompositionGuidance []string                                              `json:"compositionGuidance" api:"required"`
	ExamplePatterns     []ComposeNewResponseComposeRefineResultExamplePattern `json:"examplePatterns" api:"required"`
	// X post intent seeded with the topic.
	IntentURL string `json:"intentUrl" api:"required" format:"uri"`
	NextStep  string `json:"nextStep" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CompositionGuidance respjson.Field
		ExamplePatterns     respjson.Field
		IntentURL           respjson.Field
		NextStep            respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ComposeNewResponseComposeRefineResult) RawJSON() string { return r.JSON.raw }
func (r *ComposeNewResponseComposeRefineResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComposeNewResponseComposeRefineResultExamplePattern struct {
	Description string `json:"description" api:"required"`
	Pattern     string `json:"pattern" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Description respjson.Field
		Pattern     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ComposeNewResponseComposeRefineResultExamplePattern) RawJSON() string { return r.JSON.raw }
func (r *ComposeNewResponseComposeRefineResultExamplePattern) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComposeNewResponseComposeScoreResult struct {
	// Deterministic editorial checks. Not a reach prediction.
	Checklist     []ComposeNewResponseComposeScoreResultChecklist `json:"checklist" api:"required"`
	NextStep      string                                          `json:"nextStep" api:"required"`
	Passed        bool                                            `json:"passed" api:"required"`
	PassedCount   int64                                           `json:"passedCount" api:"required"`
	TopSuggestion string                                          `json:"topSuggestion" api:"required"`
	TotalChecks   int64                                           `json:"totalChecks" api:"required"`
	// Present only when every check passes.
	IntentURL string `json:"intentUrl" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Checklist     respjson.Field
		NextStep      respjson.Field
		Passed        respjson.Field
		PassedCount   respjson.Field
		TopSuggestion respjson.Field
		TotalChecks   respjson.Field
		IntentURL     respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ComposeNewResponseComposeScoreResult) RawJSON() string { return r.JSON.raw }
func (r *ComposeNewResponseComposeScoreResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComposeNewResponseComposeScoreResultChecklist struct {
	Factor string `json:"factor" api:"required"`
	Passed bool   `json:"passed" api:"required"`
	// Present only when the check fails.
	Suggestion string `json:"suggestion"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Factor      respjson.Field
		Passed      respjson.Field
		Suggestion  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ComposeNewResponseComposeScoreResultChecklist) RawJSON() string { return r.JSON.raw }
func (r *ComposeNewResponseComposeScoreResultChecklist) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComposeNewParams struct {

	//
	// Request body variants
	//

	// This field is a request body variant, only one variant field can be set.
	OfComposePrepareRequest *ComposeNewParamsBodyComposePrepareRequest `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfComposeRefineRequest *ComposeNewParamsBodyComposeRefineRequest `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfComposeScoreRequest *ComposeNewParamsBodyComposeScoreRequest `json:",inline"`

	paramObj
}

func (u ComposeNewParams) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfComposePrepareRequest, u.OfComposeRefineRequest, u.OfComposeScoreRequest)
}
func (r *ComposeNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Step, Topic are required.
type ComposeNewParamsBodyComposePrepareRequest struct {
	// Subject for the post.
	Topic string `json:"topic" api:"required"`
	// Username from a style analysis saved to this account.
	StyleUsername param.Opt[string] `json:"styleUsername,omitzero"`
	// Editorial goal used to order the rules and questions.
	//
	// Any of "engagement", "followers", "authority", "conversation".
	Goal string `json:"goal,omitzero"`
	// This field can be elided, and will marshal its zero value as "compose".
	Step        constant.Compose `json:"step" default:"compose"`
	ExtraFields map[string]any   `json:"-"`
	paramObj
}

func (r ComposeNewParamsBodyComposePrepareRequest) MarshalJSON() (data []byte, err error) {
	type shadow ComposeNewParamsBodyComposePrepareRequest
	return param.MarshalWithExtras(r, (*shadow)(&r), r.ExtraFields)
}
func (r *ComposeNewParamsBodyComposePrepareRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ComposeNewParamsBodyComposePrepareRequest](
		"goal", "engagement", "followers", "authority", "conversation",
	)
}

// The properties Goal, Step, Tone, Topic are required.
type ComposeNewParamsBodyComposeRefineRequest struct {
	// Editorial goal for the guidance.
	//
	// Any of "engagement", "followers", "authority", "conversation".
	Goal string `json:"goal,omitzero" api:"required"`
	// Requested writing tone.
	Tone string `json:"tone" api:"required"`
	// Subject for the post.
	Topic string `json:"topic" api:"required"`
	// Audience, constraints, sources, or other writing context.
	AdditionalContext param.Opt[string] `json:"additionalContext,omitzero"`
	// Specific action the draft should request.
	CallToAction param.Opt[string] `json:"callToAction,omitzero"`
	// Planned media type.
	//
	// Any of "photo", "video", "none".
	MediaType string `json:"mediaType,omitzero"`
	// This field can be elided, and will marshal its zero value as "refine".
	Step        constant.Refine `json:"step" default:"refine"`
	ExtraFields map[string]any  `json:"-"`
	paramObj
}

func (r ComposeNewParamsBodyComposeRefineRequest) MarshalJSON() (data []byte, err error) {
	type shadow ComposeNewParamsBodyComposeRefineRequest
	return param.MarshalWithExtras(r, (*shadow)(&r), r.ExtraFields)
}
func (r *ComposeNewParamsBodyComposeRefineRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ComposeNewParamsBodyComposeRefineRequest](
		"goal", "engagement", "followers", "authority", "conversation",
	)
	apijson.RegisterFieldValidator[ComposeNewParamsBodyComposeRefineRequest](
		"mediaType", "photo", "video", "none",
	)
}

// The properties Draft, Step are required.
type ComposeNewParamsBodyComposeScoreRequest struct {
	// Full post text for deterministic editorial checks.
	Draft string `json:"draft" api:"required"`
	// True when a separate link card is attached.
	HasLink param.Opt[bool] `json:"hasLink,omitzero"`
	// Accepted for backward compatibility. Text checks ignore this field.
	//
	// Deprecated: Ignored. Remove this field. Use hasLink for a separate link card.
	HasMedia param.Opt[bool] `json:"hasMedia,omitzero"`
	// This field can be elided, and will marshal its zero value as "score".
	Step        constant.Score `json:"step" default:"score"`
	ExtraFields map[string]any `json:"-"`
	paramObj
}

func (r ComposeNewParamsBodyComposeScoreRequest) MarshalJSON() (data []byte, err error) {
	type shadow ComposeNewParamsBodyComposeScoreRequest
	return param.MarshalWithExtras(r, (*shadow)(&r), r.ExtraFields)
}
func (r *ComposeNewParamsBodyComposeScoreRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
