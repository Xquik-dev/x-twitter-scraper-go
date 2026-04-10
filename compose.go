// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package xtwitterscraper

import (
	"context"
	"net/http"
	"slices"

	"github.com/stainless-sdks/x-twitter-scraper-go/internal/apijson"
	"github.com/stainless-sdks/x-twitter-scraper-go/internal/requestconfig"
	"github.com/stainless-sdks/x-twitter-scraper-go/option"
	"github.com/stainless-sdks/x-twitter-scraper-go/packages/param"
	"github.com/stainless-sdks/x-twitter-scraper-go/packages/respjson"
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

// Compose, refine, or score a tweet
func (r *ComposeService) New(ctx context.Context, body ComposeNewParams, opts ...option.RequestOption) (res *ComposeNewResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "compose"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type ComposeNewResponse struct {
	// AI feedback on the draft
	Feedback string `json:"feedback"`
	// Engagement score (0-100)
	Score float64 `json:"score"`
	// Improvement suggestions
	Suggestions []string `json:"suggestions"`
	// Generated or refined tweet text
	Text        string         `json:"text"`
	ExtraFields map[string]any `json:"" api:"extrafields"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Feedback    respjson.Field
		Score       respjson.Field
		Suggestions respjson.Field
		Text        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ComposeNewResponse) RawJSON() string { return r.JSON.raw }
func (r *ComposeNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComposeNewParams struct {
	// Workflow step
	//
	// Any of "compose", "refine", "score".
	Step ComposeNewParamsStep `json:"step,omitzero" api:"required"`
	// Extra context or URLs (refine)
	AdditionalContext param.Opt[string] `json:"additionalContext,omitzero"`
	// Desired call to action (refine)
	CallToAction param.Opt[string] `json:"callToAction,omitzero"`
	// Tweet draft text to evaluate (score)
	Draft param.Opt[string] `json:"draft,omitzero"`
	// Whether a link is attached (score)
	HasLink param.Opt[bool] `json:"hasLink,omitzero"`
	// Whether media is attached (score)
	HasMedia param.Opt[bool] `json:"hasMedia,omitzero"`
	// Cached style username for voice matching (compose)
	StyleUsername param.Opt[string] `json:"styleUsername,omitzero"`
	// Desired tone (refine)
	Tone param.Opt[string] `json:"tone,omitzero"`
	// Tweet topic (compose, refine)
	Topic param.Opt[string] `json:"topic,omitzero"`
	// Optimization goal
	//
	// Any of "engagement", "followers", "authority", "conversation".
	Goal ComposeNewParamsGoal `json:"goal,omitzero"`
	// Media type (refine)
	//
	// Any of "photo", "video", "none".
	MediaType ComposeNewParamsMediaType `json:"mediaType,omitzero"`
	paramObj
}

func (r ComposeNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ComposeNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ComposeNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Workflow step
type ComposeNewParamsStep string

const (
	ComposeNewParamsStepCompose ComposeNewParamsStep = "compose"
	ComposeNewParamsStepRefine  ComposeNewParamsStep = "refine"
	ComposeNewParamsStepScore   ComposeNewParamsStep = "score"
)

// Optimization goal
type ComposeNewParamsGoal string

const (
	ComposeNewParamsGoalEngagement   ComposeNewParamsGoal = "engagement"
	ComposeNewParamsGoalFollowers    ComposeNewParamsGoal = "followers"
	ComposeNewParamsGoalAuthority    ComposeNewParamsGoal = "authority"
	ComposeNewParamsGoalConversation ComposeNewParamsGoal = "conversation"
)

// Media type (refine)
type ComposeNewParamsMediaType string

const (
	ComposeNewParamsMediaTypePhoto ComposeNewParamsMediaType = "photo"
	ComposeNewParamsMediaTypeVideo ComposeNewParamsMediaType = "video"
	ComposeNewParamsMediaTypeNone  ComposeNewParamsMediaType = "none"
)
