// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

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
	"github.com/Xquik-dev/x-twitter-scraper-go/shared"
	"github.com/Xquik-dev/x-twitter-scraper-go/shared/constant"
)

// AI analysis of posts, at 2 credits per analyzed post
//
// XAnalysisService contains methods and other services that help with interacting
// with the x-twitter-scraper API.
//
// You should not instantiate this service directly, and instead use the
// [NewXAnalysisService] method instead.
type XAnalysisService struct {
	options []option.RequestOption
}

// NewXAnalysisService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewXAnalysisService(opts ...option.RequestOption) (r XAnalysisService) {
	r = XAnalysisService{}
	r.options = opts
	return
}

// Answers whether each post is about your brand, the author's attitude toward it &
// the author's relationship with it. Name the brand in analysis.targets. Each
// analyzed post costs 2 credits.
func (r *XAnalysisService) Brand(ctx context.Context, body XAnalysisBrandParams, opts ...option.RequestOption) (res *XAnalysisBrandResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "x/analysis/brand"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Answers 1 to 8 questions you write in analysis.questions. A choice picks a
// category, such as a topic. A score picks a level. A probability answers yes or
// no. Without questions it answers the sentiment questions. Each analyzed post
// costs 2 credits.
func (r *XAnalysisService) Classify(ctx context.Context, body XAnalysisClassifyParams, opts ...option.RequestOption) (res *XAnalysisClassifyResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "x/analysis/classify"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Answers each post's stance on an asset, its content type, its conviction & its
// relevance. The summary counts bullish & bearish posts per cashtag. Name the
// asset in analysis.targets. Each analyzed post costs 2 credits.
func (r *XAnalysisService) MarketSignals(ctx context.Context, body XAnalysisMarketSignalsParams, opts ...option.RequestOption) (res *XAnalysisMarketSignalsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "x/analysis/market-signals"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Answers each post's format, such as reporting, commentary or speculation, its
// source attribution & its relevance to your topic. Name the topic in
// analysis.targets. Each analyzed post costs 2 credits.
func (r *XAnalysisService) News(ctx context.Context, body XAnalysisNewsParams, opts ...option.RequestOption) (res *XAnalysisNewsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "x/analysis/news"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Answers each post's attitude, its intensity & its sarcasm probability. Send post
// IDs, an X search, an account, a List, a post's replies or quotes, or your own
// texts. Each analyzed post costs 2 credits. A post without an analysis costs
// nothing.
func (r *XAnalysisService) Sentiment(ctx context.Context, body XAnalysisSentimentParams, opts ...option.RequestOption) (res *XAnalysisSentimentResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "x/analysis/sentiment"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Answers each post's hook, clarity, payoff, humor, spam risk & likely reaction, &
// adds a Viral Score from 0 to 100 with a verdict. The score rates the wording. It
// does not predict likes or views. Send drafts in texts to score them before
// posting. Each analyzed post costs 2 credits.
func (r *XAnalysisService) ViralScore(ctx context.Context, body XAnalysisViralScoreParams, opts ...option.RequestOption) (res *XAnalysisViralScoreResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "x/analysis/viral-score"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Analyzed posts, the posts without an analysis & the totals of the answers. Each
// row in results costs 2 credits.
type XAnalysisBrandResponse struct {
	// Totals of the answers across the analyzed posts.
	AnalysisSummary XAnalysisBrandResponseAnalysisSummary `json:"analysisSummary" api:"required"`
	// Whether the search has more posts. Send next_cursor as cursor.
	HasNextPage bool `json:"has_next_page" api:"required"`
	// Cursor of the search's next page, else empty.
	NextCursor string `json:"next_cursor" api:"required"`
	// Each analyzed post with its answers.
	Results []XAnalysisBrandResponseResult `json:"results" api:"required"`
	// Each post without an analysis, by ID & reason. These cost nothing.
	Unanalyzed []XAnalysisBrandResponseUnanalyzed `json:"unanalyzed" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AnalysisSummary respjson.Field
		HasNextPage     respjson.Field
		NextCursor      respjson.Field
		Results         respjson.Field
		Unanalyzed      respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XAnalysisBrandResponse) RawJSON() string { return r.JSON.raw }
func (r *XAnalysisBrandResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// Totals of the answers across the analyzed posts.
type XAnalysisBrandResponseAnalysisSummary struct {
	// Posts & answers per cashtag, most posts first, up to 20. signal is present for
	// the market questions.
	Cashtags []XAnalysisBrandResponseAnalysisSummaryCashtag `json:"cashtags" api:"required"`
	// Likes, reposts, replies & quotes of the analyzed posts.
	Engagement int64 `json:"engagement" api:"required"`
	// Posts per comparison status, & up to 50 changed posts with their changes. Null
	// when the answer holds no post.
	Monitor XAnalysisBrandResponseAnalysisSummaryMonitor `json:"monitor" api:"required"`
	// Totals by question ID. A choice has counts, shares, engagementShares & top posts
	// per category. A score has mean, engagementWeightedMean & levels. A probability
	// has mean, yes & no.
	Questions map[string]PublicMetadata `json:"questions" api:"required"`
	// Posts by analysis outcome.
	Rows          XAnalysisBrandResponseAnalysisSummaryRows `json:"rows" api:"required"`
	SchemaVersion int64                                     `json:"schemaVersion" api:"required"`
	// Posts per linked hostname, most posts first, up to 20.
	SourceDomains []XAnalysisBrandResponseAnalysisSummarySourceDomain `json:"sourceDomains" api:"required"`
	// Mentions & answers per target in analysis.targets.
	Targets []XAnalysisBrandResponseAnalysisSummaryTarget `json:"targets" api:"required"`
	// Only from /x/analysis/viral-score. Scored posts, the average score, verdict
	// counts, the scores against actual engagement, a report per author & the
	// leaderboards.
	Viral PublicMetadata `json:"viral"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Cashtags      respjson.Field
		Engagement    respjson.Field
		Monitor       respjson.Field
		Questions     respjson.Field
		Rows          respjson.Field
		SchemaVersion respjson.Field
		SourceDomains respjson.Field
		Targets       respjson.Field
		Viral         respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XAnalysisBrandResponseAnalysisSummary) RawJSON() string { return r.JSON.raw }
func (r *XAnalysisBrandResponseAnalysisSummary) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type XAnalysisBrandResponseAnalysisSummaryCashtag struct {
	Cashtag string `json:"cashtag" api:"required"`
	// Category counts of the cashtag's posts by question ID.
	Choices PublicMetadata                                     `json:"choices" api:"required"`
	Rows    int64                                              `json:"rows" api:"required"`
	Signal  XAnalysisBrandResponseAnalysisSummaryCashtagSignal `json:"signal"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Cashtag     respjson.Field
		Choices     respjson.Field
		Rows        respjson.Field
		Signal      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XAnalysisBrandResponseAnalysisSummaryCashtag) RawJSON() string { return r.JSON.raw }
func (r *XAnalysisBrandResponseAnalysisSummaryCashtag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type XAnalysisBrandResponseAnalysisSummaryCashtagSignal struct {
	Bearish int64 `json:"bearish" api:"required"`
	Bullish int64 `json:"bullish" api:"required"`
	// Bullish minus bearish posts, over all posts.
	Score float64 `json:"score" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Bearish     respjson.Field
		Bullish     respjson.Field
		Score       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XAnalysisBrandResponseAnalysisSummaryCashtagSignal) RawJSON() string { return r.JSON.raw }
func (r *XAnalysisBrandResponseAnalysisSummaryCashtagSignal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// Posts per comparison status, & up to 50 changed posts with their changes. Null
// when the answer holds no post.
type XAnalysisBrandResponseAnalysisSummaryMonitor struct {
	ChangedRows []XAnalysisBrandResponseAnalysisSummaryMonitorChangedRow `json:"changedRows" api:"required"`
	// Posts per monitor status.
	Statuses map[string]int64 `json:"statuses" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChangedRows respjson.Field
		Statuses    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XAnalysisBrandResponseAnalysisSummaryMonitor) RawJSON() string { return r.JSON.raw }
func (r *XAnalysisBrandResponseAnalysisSummaryMonitor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type XAnalysisBrandResponseAnalysisSummaryMonitorChangedRow struct {
	Changes []XAnalysisBrandResponseAnalysisSummaryMonitorChangedRowChange `json:"changes" api:"required"`
	TweetID string                                                         `json:"tweetId" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Changes     respjson.Field
		TweetID     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XAnalysisBrandResponseAnalysisSummaryMonitorChangedRow) RawJSON() string { return r.JSON.raw }
func (r *XAnalysisBrandResponseAnalysisSummaryMonitorChangedRow) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type XAnalysisBrandResponseAnalysisSummaryMonitorChangedRowChange struct {
	Current    string `json:"current" api:"required"`
	Previous   string `json:"previous" api:"required"`
	QuestionID string `json:"questionId" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Current     respjson.Field
		Previous    respjson.Field
		QuestionID  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XAnalysisBrandResponseAnalysisSummaryMonitorChangedRowChange) RawJSON() string {
	return r.JSON.raw
}
func (r *XAnalysisBrandResponseAnalysisSummaryMonitorChangedRowChange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// Posts by analysis outcome.
type XAnalysisBrandResponseAnalysisSummaryRows struct {
	Analyzed int64 `json:"analyzed" api:"required"`
	Failed   int64 `json:"failed" api:"required"`
	Skipped  int64 `json:"skipped" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Analyzed    respjson.Field
		Failed      respjson.Field
		Skipped     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XAnalysisBrandResponseAnalysisSummaryRows) RawJSON() string { return r.JSON.raw }
func (r *XAnalysisBrandResponseAnalysisSummaryRows) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type XAnalysisBrandResponseAnalysisSummarySourceDomain struct {
	Domain string `json:"domain" api:"required"`
	Posts  int64  `json:"posts" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Domain      respjson.Field
		Posts       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XAnalysisBrandResponseAnalysisSummarySourceDomain) RawJSON() string { return r.JSON.raw }
func (r *XAnalysisBrandResponseAnalysisSummarySourceDomain) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type XAnalysisBrandResponseAnalysisSummaryTarget struct {
	// Category counts by question ID.
	Choices    PublicMetadata `json:"choices" api:"required"`
	Engagement int64          `json:"engagement" api:"required"`
	Mentions   int64          `json:"mentions" api:"required"`
	Name       string         `json:"name" api:"required"`
	Share      float64        `json:"share" api:"required"`
	// The most engaged posts per question ID & category.
	Top PublicMetadata `json:"top" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Choices     respjson.Field
		Engagement  respjson.Field
		Mentions    respjson.Field
		Name        respjson.Field
		Share       respjson.Field
		Top         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XAnalysisBrandResponseAnalysisSummaryTarget) RawJSON() string { return r.JSON.raw }
func (r *XAnalysisBrandResponseAnalysisSummaryTarget) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// 1 analyzed post with its answers.
type XAnalysisBrandResponseResult struct {
	// The answers in full, with their probabilities.
	Analysis XAnalysisBrandResponseResultAnalysis `json:"analysis" api:"required"`
	// Each question's decision by question ID. A choice gives its category, a score
	// its level from 0, which can have decimals, & a probability its value from 0 to
	// 1.
	Answers map[string]XAnalysisBrandResponseResultAnswersUnion `json:"answers" api:"required"`
	// Cashtags in the post's text, in upper case.
	Cashtags []string `json:"cashtags" api:"required"`
	// How the post's answers compare with its row in baseline, as the AI Actors
	// compare a run with an earlier dataset.
	Monitor XAnalysisBrandResponseResultMonitor `json:"monitor" api:"required"`
	// Hostnames the post links to, without t.co.
	SourceDomains []string `json:"sourceDomains" api:"required"`
	// The post, as /x/tweets/search returns it. Your own texts have the IDs text:1,
	// text:2 & so on.
	Tweet shared.SearchTweet `json:"tweet" api:"required"`
	// Only from /x/analysis/viral-score. The score with its weights version, the caps
	// that applied & the engagement comparison.
	Viral PublicMetadata `json:"viral"`
	// Only from /x/analysis/viral-score. log10 of 1 plus the weighted engagement per
	// 1,000 followers. Null without a follower count.
	ViralActualEngagementRate float64 `json:"viralActualEngagementRate" api:"nullable"`
	// Only from /x/analysis/viral-score. The post's public counts under X's published
	// ranking weights, per 1,000 views or followers.
	ViralAlgorithmScore float64 `json:"viralAlgorithmScore" api:"nullable"`
	// Only from /x/analysis/viral-score. How viral the wording reads, or null when a
	// trait has no answer.
	ViralScore int64 `json:"viralScore" api:"nullable"`
	// Only from /x/analysis/viral-score. send_it from 70, edit_first from 40, else
	// sleep_on_it.
	//
	// Any of "send_it", "edit_first", "sleep_on_it".
	ViralVerdict XAnalysisBrandResponseResultViralVerdict `json:"viralVerdict" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Analysis                  respjson.Field
		Answers                   respjson.Field
		Cashtags                  respjson.Field
		Monitor                   respjson.Field
		SourceDomains             respjson.Field
		Tweet                     respjson.Field
		Viral                     respjson.Field
		ViralActualEngagementRate respjson.Field
		ViralAlgorithmScore       respjson.Field
		ViralScore                respjson.Field
		ViralVerdict              respjson.Field
		ExtraFields               map[string]respjson.Field
		raw                       string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XAnalysisBrandResponseResult) RawJSON() string { return r.JSON.raw }
func (r *XAnalysisBrandResponseResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// The answers in full, with their probabilities.
type XAnalysisBrandResponseResultAnalysis struct {
	// 1 answer per question, in question order.
	Answers []XAnalysisBrandResponseResultAnalysisAnswer `json:"answers" api:"required"`
	// What the AI read beside the post's text.
	ContextAvailability XAnalysisBrandResponseResultAnalysisContextAvailability `json:"contextAvailability" api:"required"`
	// Bytes of post & settings the AI read.
	ContextBytes int64 `json:"contextBytes" api:"required"`
	// The ID & version of each question asked.
	Questions []XAnalysisBrandResponseResultAnalysisQuestion `json:"questions" api:"required"`
	Status    constant.Succeeded                             `json:"status" default:"succeeded"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Answers             respjson.Field
		ContextAvailability respjson.Field
		ContextBytes        respjson.Field
		Questions           respjson.Field
		Status              respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XAnalysisBrandResponseResultAnalysis) RawJSON() string { return r.JSON.raw }
func (r *XAnalysisBrandResponseResultAnalysis) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// 1 question's answer. A probability has probability. A choice or score has value,
// confidence & the probability of each category or level.
type XAnalysisBrandResponseResultAnalysisAnswer struct {
	QuestionID      string `json:"questionId" api:"required"`
	QuestionVersion string `json:"questionVersion" api:"required"`
	// Any of "choice", "score", "probability".
	Type       XAnalysisBrandResponseResultAnalysisAnswerType `json:"type" api:"required"`
	Confidence float64                                        `json:"confidence"`
	// For a score. The levels the question named, from 0.
	Levels []string `json:"levels"`
	// Probability of each category, or of each level by index.
	Probabilities map[string]float64 `json:"probabilities"`
	// For a probability. The chance the answer is yes.
	Probability float64 `json:"probability"`
	// The chosen category, or the level from 0, which can have decimals, such as 1.4
	// between levels 1 & 2.
	Value XAnalysisBrandResponseResultAnalysisAnswerValueUnion `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		QuestionID      respjson.Field
		QuestionVersion respjson.Field
		Type            respjson.Field
		Confidence      respjson.Field
		Levels          respjson.Field
		Probabilities   respjson.Field
		Probability     respjson.Field
		Value           respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XAnalysisBrandResponseResultAnalysisAnswer) RawJSON() string { return r.JSON.raw }
func (r *XAnalysisBrandResponseResultAnalysisAnswer) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type XAnalysisBrandResponseResultAnalysisAnswerType string

const (
	XAnalysisBrandResponseResultAnalysisAnswerTypeChoice      XAnalysisBrandResponseResultAnalysisAnswerType = "choice"
	XAnalysisBrandResponseResultAnalysisAnswerTypeScore       XAnalysisBrandResponseResultAnalysisAnswerType = "score"
	XAnalysisBrandResponseResultAnalysisAnswerTypeProbability XAnalysisBrandResponseResultAnalysisAnswerType = "probability"
)

// XAnalysisBrandResponseResultAnalysisAnswerValueUnion contains all possible
// properties and values from [float64], [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfFloat OfString]
type XAnalysisBrandResponseResultAnalysisAnswerValueUnion struct {
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	JSON     struct {
		OfFloat  respjson.Field
		OfString respjson.Field
		raw      string
	} `json:"-"`
}

func (u XAnalysisBrandResponseResultAnalysisAnswerValueUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v, nil)
	return
}

func (u XAnalysisBrandResponseResultAnalysisAnswerValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v, nil)
	return
}

// Returns the unmodified JSON received from the API
func (u XAnalysisBrandResponseResultAnalysisAnswerValueUnion) RawJSON() string { return u.JSON.raw }

func (r *XAnalysisBrandResponseResultAnalysisAnswerValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// What the AI read beside the post's text.
type XAnalysisBrandResponseResultAnalysisContextAvailability struct {
	// Any of "text_blocks", "summary", "not_supplied".
	Article XAnalysisBrandResponseResultAnalysisContextAvailabilityArticle `json:"article" api:"required"`
	// Any of "available", "missing".
	Author XAnalysisBrandResponseResultAnalysisContextAvailabilityAuthor `json:"author" api:"required"`
	// Any of "text_descriptions_only", "not_supplied".
	Media XAnalysisBrandResponseResultAnalysisContextAvailabilityMedia `json:"media" api:"required"`
	// truncated when the post was cut to fit.
	//
	// Any of "full", "truncated".
	PostText XAnalysisBrandResponseResultAnalysisContextAvailabilityPostText `json:"postText" api:"required"`
	// Any of "available", "missing", "not_applicable".
	Quote XAnalysisBrandResponseResultAnalysisContextAvailabilityQuote `json:"quote" api:"required"`
	// Any of "available", "missing", "not_applicable".
	Reply XAnalysisBrandResponseResultAnalysisContextAvailabilityReply `json:"reply" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Article     respjson.Field
		Author      respjson.Field
		Media       respjson.Field
		PostText    respjson.Field
		Quote       respjson.Field
		Reply       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XAnalysisBrandResponseResultAnalysisContextAvailability) RawJSON() string { return r.JSON.raw }
func (r *XAnalysisBrandResponseResultAnalysisContextAvailability) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type XAnalysisBrandResponseResultAnalysisContextAvailabilityArticle string

const (
	XAnalysisBrandResponseResultAnalysisContextAvailabilityArticleTextBlocks  XAnalysisBrandResponseResultAnalysisContextAvailabilityArticle = "text_blocks"
	XAnalysisBrandResponseResultAnalysisContextAvailabilityArticleSummary     XAnalysisBrandResponseResultAnalysisContextAvailabilityArticle = "summary"
	XAnalysisBrandResponseResultAnalysisContextAvailabilityArticleNotSupplied XAnalysisBrandResponseResultAnalysisContextAvailabilityArticle = "not_supplied"
)

type XAnalysisBrandResponseResultAnalysisContextAvailabilityAuthor string

const (
	XAnalysisBrandResponseResultAnalysisContextAvailabilityAuthorAvailable XAnalysisBrandResponseResultAnalysisContextAvailabilityAuthor = "available"
	XAnalysisBrandResponseResultAnalysisContextAvailabilityAuthorMissing   XAnalysisBrandResponseResultAnalysisContextAvailabilityAuthor = "missing"
)

type XAnalysisBrandResponseResultAnalysisContextAvailabilityMedia string

const (
	XAnalysisBrandResponseResultAnalysisContextAvailabilityMediaTextDescriptionsOnly XAnalysisBrandResponseResultAnalysisContextAvailabilityMedia = "text_descriptions_only"
	XAnalysisBrandResponseResultAnalysisContextAvailabilityMediaNotSupplied          XAnalysisBrandResponseResultAnalysisContextAvailabilityMedia = "not_supplied"
)

// truncated when the post was cut to fit.
type XAnalysisBrandResponseResultAnalysisContextAvailabilityPostText string

const (
	XAnalysisBrandResponseResultAnalysisContextAvailabilityPostTextFull      XAnalysisBrandResponseResultAnalysisContextAvailabilityPostText = "full"
	XAnalysisBrandResponseResultAnalysisContextAvailabilityPostTextTruncated XAnalysisBrandResponseResultAnalysisContextAvailabilityPostText = "truncated"
)

type XAnalysisBrandResponseResultAnalysisContextAvailabilityQuote string

const (
	XAnalysisBrandResponseResultAnalysisContextAvailabilityQuoteAvailable     XAnalysisBrandResponseResultAnalysisContextAvailabilityQuote = "available"
	XAnalysisBrandResponseResultAnalysisContextAvailabilityQuoteMissing       XAnalysisBrandResponseResultAnalysisContextAvailabilityQuote = "missing"
	XAnalysisBrandResponseResultAnalysisContextAvailabilityQuoteNotApplicable XAnalysisBrandResponseResultAnalysisContextAvailabilityQuote = "not_applicable"
)

type XAnalysisBrandResponseResultAnalysisContextAvailabilityReply string

const (
	XAnalysisBrandResponseResultAnalysisContextAvailabilityReplyAvailable     XAnalysisBrandResponseResultAnalysisContextAvailabilityReply = "available"
	XAnalysisBrandResponseResultAnalysisContextAvailabilityReplyMissing       XAnalysisBrandResponseResultAnalysisContextAvailabilityReply = "missing"
	XAnalysisBrandResponseResultAnalysisContextAvailabilityReplyNotApplicable XAnalysisBrandResponseResultAnalysisContextAvailabilityReply = "not_applicable"
)

type XAnalysisBrandResponseResultAnalysisQuestion struct {
	ID      string `json:"id" api:"required"`
	Version string `json:"version" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Version     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XAnalysisBrandResponseResultAnalysisQuestion) RawJSON() string { return r.JSON.raw }
func (r *XAnalysisBrandResponseResultAnalysisQuestion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// XAnalysisBrandResponseResultAnswersUnion contains all possible properties and
// values from [float64], [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfFloat OfString]
type XAnalysisBrandResponseResultAnswersUnion struct {
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	JSON     struct {
		OfFloat  respjson.Field
		OfString respjson.Field
		raw      string
	} `json:"-"`
}

func (u XAnalysisBrandResponseResultAnswersUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v, nil)
	return
}

func (u XAnalysisBrandResponseResultAnswersUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v, nil)
	return
}

// Returns the unmodified JSON received from the API
func (u XAnalysisBrandResponseResultAnswersUnion) RawJSON() string { return u.JSON.raw }

func (r *XAnalysisBrandResponseResultAnswersUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// How the post's answers compare with its row in baseline, as the AI Actors
// compare a run with an earlier dataset.
type XAnalysisBrandResponseResultMonitor struct {
	// Always null. The baseline is the rows you sent.
	BaselineDatasetID json.RawMessage `json:"baselineDatasetId" api:"required"`
	// IDs of the questions whose decision changed.
	ChangedQuestionIDs []string `json:"changedQuestionIds" api:"required"`
	// Each changed decision, earlier & now.
	Changes []XAnalysisBrandResponseResultMonitorChange `json:"changes" api:"required"`
	// Hash of the questions, targets & context. Rows with another hash are
	// not_comparable.
	ConfigurationFingerprint string `json:"configurationFingerprint" api:"required"`
	// first_run without baseline. new_to_baseline for a post baseline lacks. changed
	// when an answer clearly left its earlier decision. not_comparable when other
	// settings made the earlier row or it answers none of these questions.
	//
	// Any of "first_run", "new_to_baseline", "unchanged", "changed", "not_comparable",
	// "analysis_unavailable".
	Status XAnalysisBrandResponseResultMonitorStatus `json:"status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BaselineDatasetID        respjson.Field
		ChangedQuestionIDs       respjson.Field
		Changes                  respjson.Field
		ConfigurationFingerprint respjson.Field
		Status                   respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XAnalysisBrandResponseResultMonitor) RawJSON() string { return r.JSON.raw }
func (r *XAnalysisBrandResponseResultMonitor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type XAnalysisBrandResponseResultMonitorChange struct {
	Current    string `json:"current" api:"required"`
	Previous   string `json:"previous" api:"required"`
	QuestionID string `json:"questionId" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Current     respjson.Field
		Previous    respjson.Field
		QuestionID  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XAnalysisBrandResponseResultMonitorChange) RawJSON() string { return r.JSON.raw }
func (r *XAnalysisBrandResponseResultMonitorChange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// first_run without baseline. new_to_baseline for a post baseline lacks. changed
// when an answer clearly left its earlier decision. not_comparable when other
// settings made the earlier row or it answers none of these questions.
type XAnalysisBrandResponseResultMonitorStatus string

const (
	XAnalysisBrandResponseResultMonitorStatusFirstRun            XAnalysisBrandResponseResultMonitorStatus = "first_run"
	XAnalysisBrandResponseResultMonitorStatusNewToBaseline       XAnalysisBrandResponseResultMonitorStatus = "new_to_baseline"
	XAnalysisBrandResponseResultMonitorStatusUnchanged           XAnalysisBrandResponseResultMonitorStatus = "unchanged"
	XAnalysisBrandResponseResultMonitorStatusChanged             XAnalysisBrandResponseResultMonitorStatus = "changed"
	XAnalysisBrandResponseResultMonitorStatusNotComparable       XAnalysisBrandResponseResultMonitorStatus = "not_comparable"
	XAnalysisBrandResponseResultMonitorStatusAnalysisUnavailable XAnalysisBrandResponseResultMonitorStatus = "analysis_unavailable"
)

// Only from /x/analysis/viral-score. send_it from 70, edit_first from 40, else
// sleep_on_it.
type XAnalysisBrandResponseResultViralVerdict string

const (
	XAnalysisBrandResponseResultViralVerdictSendIt    XAnalysisBrandResponseResultViralVerdict = "send_it"
	XAnalysisBrandResponseResultViralVerdictEditFirst XAnalysisBrandResponseResultViralVerdict = "edit_first"
	XAnalysisBrandResponseResultViralVerdictSleepOnIt XAnalysisBrandResponseResultViralVerdict = "sleep_on_it"
)

// A post without an analysis. It costs nothing.
type XAnalysisBrandResponseUnanalyzed struct {
	// Post ID. Your own texts have the IDs text:1, text:2 & so on.
	ID string `json:"id" api:"required"`
	// post_unavailable when X returns no such post. missing_text for a post without
	// text. insufficient_credits when credits ran out. context_limit when the settings
	// leave the post no room. The other reasons mean the analysis failed. Send the
	// post again.
	Reason string `json:"reason" api:"required"`
	// Any of "failed", "skipped".
	Status XAnalysisBrandResponseUnanalyzedStatus `json:"status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Reason      respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XAnalysisBrandResponseUnanalyzed) RawJSON() string { return r.JSON.raw }
func (r *XAnalysisBrandResponseUnanalyzed) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type XAnalysisBrandResponseUnanalyzedStatus string

const (
	XAnalysisBrandResponseUnanalyzedStatusFailed  XAnalysisBrandResponseUnanalyzedStatus = "failed"
	XAnalysisBrandResponseUnanalyzedStatusSkipped XAnalysisBrandResponseUnanalyzedStatus = "skipped"
)

// Analyzed posts, the posts without an analysis & the totals of the answers. Each
// row in results costs 2 credits.
type XAnalysisClassifyResponse struct {
	// Totals of the answers across the analyzed posts.
	AnalysisSummary XAnalysisClassifyResponseAnalysisSummary `json:"analysisSummary" api:"required"`
	// Whether the search has more posts. Send next_cursor as cursor.
	HasNextPage bool `json:"has_next_page" api:"required"`
	// Cursor of the search's next page, else empty.
	NextCursor string `json:"next_cursor" api:"required"`
	// Each analyzed post with its answers.
	Results []XAnalysisClassifyResponseResult `json:"results" api:"required"`
	// Each post without an analysis, by ID & reason. These cost nothing.
	Unanalyzed []XAnalysisClassifyResponseUnanalyzed `json:"unanalyzed" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AnalysisSummary respjson.Field
		HasNextPage     respjson.Field
		NextCursor      respjson.Field
		Results         respjson.Field
		Unanalyzed      respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XAnalysisClassifyResponse) RawJSON() string { return r.JSON.raw }
func (r *XAnalysisClassifyResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// Totals of the answers across the analyzed posts.
type XAnalysisClassifyResponseAnalysisSummary struct {
	// Posts & answers per cashtag, most posts first, up to 20. signal is present for
	// the market questions.
	Cashtags []XAnalysisClassifyResponseAnalysisSummaryCashtag `json:"cashtags" api:"required"`
	// Likes, reposts, replies & quotes of the analyzed posts.
	Engagement int64 `json:"engagement" api:"required"`
	// Posts per comparison status, & up to 50 changed posts with their changes. Null
	// when the answer holds no post.
	Monitor XAnalysisClassifyResponseAnalysisSummaryMonitor `json:"monitor" api:"required"`
	// Totals by question ID. A choice has counts, shares, engagementShares & top posts
	// per category. A score has mean, engagementWeightedMean & levels. A probability
	// has mean, yes & no.
	Questions map[string]PublicMetadata `json:"questions" api:"required"`
	// Posts by analysis outcome.
	Rows          XAnalysisClassifyResponseAnalysisSummaryRows `json:"rows" api:"required"`
	SchemaVersion int64                                        `json:"schemaVersion" api:"required"`
	// Posts per linked hostname, most posts first, up to 20.
	SourceDomains []XAnalysisClassifyResponseAnalysisSummarySourceDomain `json:"sourceDomains" api:"required"`
	// Mentions & answers per target in analysis.targets.
	Targets []XAnalysisClassifyResponseAnalysisSummaryTarget `json:"targets" api:"required"`
	// Only from /x/analysis/viral-score. Scored posts, the average score, verdict
	// counts, the scores against actual engagement, a report per author & the
	// leaderboards.
	Viral PublicMetadata `json:"viral"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Cashtags      respjson.Field
		Engagement    respjson.Field
		Monitor       respjson.Field
		Questions     respjson.Field
		Rows          respjson.Field
		SchemaVersion respjson.Field
		SourceDomains respjson.Field
		Targets       respjson.Field
		Viral         respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XAnalysisClassifyResponseAnalysisSummary) RawJSON() string { return r.JSON.raw }
func (r *XAnalysisClassifyResponseAnalysisSummary) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type XAnalysisClassifyResponseAnalysisSummaryCashtag struct {
	Cashtag string `json:"cashtag" api:"required"`
	// Category counts of the cashtag's posts by question ID.
	Choices PublicMetadata                                        `json:"choices" api:"required"`
	Rows    int64                                                 `json:"rows" api:"required"`
	Signal  XAnalysisClassifyResponseAnalysisSummaryCashtagSignal `json:"signal"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Cashtag     respjson.Field
		Choices     respjson.Field
		Rows        respjson.Field
		Signal      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XAnalysisClassifyResponseAnalysisSummaryCashtag) RawJSON() string { return r.JSON.raw }
func (r *XAnalysisClassifyResponseAnalysisSummaryCashtag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type XAnalysisClassifyResponseAnalysisSummaryCashtagSignal struct {
	Bearish int64 `json:"bearish" api:"required"`
	Bullish int64 `json:"bullish" api:"required"`
	// Bullish minus bearish posts, over all posts.
	Score float64 `json:"score" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Bearish     respjson.Field
		Bullish     respjson.Field
		Score       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XAnalysisClassifyResponseAnalysisSummaryCashtagSignal) RawJSON() string { return r.JSON.raw }
func (r *XAnalysisClassifyResponseAnalysisSummaryCashtagSignal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// Posts per comparison status, & up to 50 changed posts with their changes. Null
// when the answer holds no post.
type XAnalysisClassifyResponseAnalysisSummaryMonitor struct {
	ChangedRows []XAnalysisClassifyResponseAnalysisSummaryMonitorChangedRow `json:"changedRows" api:"required"`
	// Posts per monitor status.
	Statuses map[string]int64 `json:"statuses" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChangedRows respjson.Field
		Statuses    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XAnalysisClassifyResponseAnalysisSummaryMonitor) RawJSON() string { return r.JSON.raw }
func (r *XAnalysisClassifyResponseAnalysisSummaryMonitor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type XAnalysisClassifyResponseAnalysisSummaryMonitorChangedRow struct {
	Changes []XAnalysisClassifyResponseAnalysisSummaryMonitorChangedRowChange `json:"changes" api:"required"`
	TweetID string                                                            `json:"tweetId" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Changes     respjson.Field
		TweetID     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XAnalysisClassifyResponseAnalysisSummaryMonitorChangedRow) RawJSON() string {
	return r.JSON.raw
}
func (r *XAnalysisClassifyResponseAnalysisSummaryMonitorChangedRow) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type XAnalysisClassifyResponseAnalysisSummaryMonitorChangedRowChange struct {
	Current    string `json:"current" api:"required"`
	Previous   string `json:"previous" api:"required"`
	QuestionID string `json:"questionId" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Current     respjson.Field
		Previous    respjson.Field
		QuestionID  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XAnalysisClassifyResponseAnalysisSummaryMonitorChangedRowChange) RawJSON() string {
	return r.JSON.raw
}
func (r *XAnalysisClassifyResponseAnalysisSummaryMonitorChangedRowChange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// Posts by analysis outcome.
type XAnalysisClassifyResponseAnalysisSummaryRows struct {
	Analyzed int64 `json:"analyzed" api:"required"`
	Failed   int64 `json:"failed" api:"required"`
	Skipped  int64 `json:"skipped" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Analyzed    respjson.Field
		Failed      respjson.Field
		Skipped     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XAnalysisClassifyResponseAnalysisSummaryRows) RawJSON() string { return r.JSON.raw }
func (r *XAnalysisClassifyResponseAnalysisSummaryRows) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type XAnalysisClassifyResponseAnalysisSummarySourceDomain struct {
	Domain string `json:"domain" api:"required"`
	Posts  int64  `json:"posts" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Domain      respjson.Field
		Posts       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XAnalysisClassifyResponseAnalysisSummarySourceDomain) RawJSON() string { return r.JSON.raw }
func (r *XAnalysisClassifyResponseAnalysisSummarySourceDomain) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type XAnalysisClassifyResponseAnalysisSummaryTarget struct {
	// Category counts by question ID.
	Choices    PublicMetadata `json:"choices" api:"required"`
	Engagement int64          `json:"engagement" api:"required"`
	Mentions   int64          `json:"mentions" api:"required"`
	Name       string         `json:"name" api:"required"`
	Share      float64        `json:"share" api:"required"`
	// The most engaged posts per question ID & category.
	Top PublicMetadata `json:"top" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Choices     respjson.Field
		Engagement  respjson.Field
		Mentions    respjson.Field
		Name        respjson.Field
		Share       respjson.Field
		Top         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XAnalysisClassifyResponseAnalysisSummaryTarget) RawJSON() string { return r.JSON.raw }
func (r *XAnalysisClassifyResponseAnalysisSummaryTarget) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// 1 analyzed post with its answers.
type XAnalysisClassifyResponseResult struct {
	// The answers in full, with their probabilities.
	Analysis XAnalysisClassifyResponseResultAnalysis `json:"analysis" api:"required"`
	// Each question's decision by question ID. A choice gives its category, a score
	// its level from 0, which can have decimals, & a probability its value from 0 to
	// 1.
	Answers map[string]XAnalysisClassifyResponseResultAnswersUnion `json:"answers" api:"required"`
	// Cashtags in the post's text, in upper case.
	Cashtags []string `json:"cashtags" api:"required"`
	// How the post's answers compare with its row in baseline, as the AI Actors
	// compare a run with an earlier dataset.
	Monitor XAnalysisClassifyResponseResultMonitor `json:"monitor" api:"required"`
	// Hostnames the post links to, without t.co.
	SourceDomains []string `json:"sourceDomains" api:"required"`
	// The post, as /x/tweets/search returns it. Your own texts have the IDs text:1,
	// text:2 & so on.
	Tweet shared.SearchTweet `json:"tweet" api:"required"`
	// Only from /x/analysis/viral-score. The score with its weights version, the caps
	// that applied & the engagement comparison.
	Viral PublicMetadata `json:"viral"`
	// Only from /x/analysis/viral-score. log10 of 1 plus the weighted engagement per
	// 1,000 followers. Null without a follower count.
	ViralActualEngagementRate float64 `json:"viralActualEngagementRate" api:"nullable"`
	// Only from /x/analysis/viral-score. The post's public counts under X's published
	// ranking weights, per 1,000 views or followers.
	ViralAlgorithmScore float64 `json:"viralAlgorithmScore" api:"nullable"`
	// Only from /x/analysis/viral-score. How viral the wording reads, or null when a
	// trait has no answer.
	ViralScore int64 `json:"viralScore" api:"nullable"`
	// Only from /x/analysis/viral-score. send_it from 70, edit_first from 40, else
	// sleep_on_it.
	//
	// Any of "send_it", "edit_first", "sleep_on_it".
	ViralVerdict XAnalysisClassifyResponseResultViralVerdict `json:"viralVerdict" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Analysis                  respjson.Field
		Answers                   respjson.Field
		Cashtags                  respjson.Field
		Monitor                   respjson.Field
		SourceDomains             respjson.Field
		Tweet                     respjson.Field
		Viral                     respjson.Field
		ViralActualEngagementRate respjson.Field
		ViralAlgorithmScore       respjson.Field
		ViralScore                respjson.Field
		ViralVerdict              respjson.Field
		ExtraFields               map[string]respjson.Field
		raw                       string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XAnalysisClassifyResponseResult) RawJSON() string { return r.JSON.raw }
func (r *XAnalysisClassifyResponseResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// The answers in full, with their probabilities.
type XAnalysisClassifyResponseResultAnalysis struct {
	// 1 answer per question, in question order.
	Answers []XAnalysisClassifyResponseResultAnalysisAnswer `json:"answers" api:"required"`
	// What the AI read beside the post's text.
	ContextAvailability XAnalysisClassifyResponseResultAnalysisContextAvailability `json:"contextAvailability" api:"required"`
	// Bytes of post & settings the AI read.
	ContextBytes int64 `json:"contextBytes" api:"required"`
	// The ID & version of each question asked.
	Questions []XAnalysisClassifyResponseResultAnalysisQuestion `json:"questions" api:"required"`
	Status    constant.Succeeded                                `json:"status" default:"succeeded"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Answers             respjson.Field
		ContextAvailability respjson.Field
		ContextBytes        respjson.Field
		Questions           respjson.Field
		Status              respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XAnalysisClassifyResponseResultAnalysis) RawJSON() string { return r.JSON.raw }
func (r *XAnalysisClassifyResponseResultAnalysis) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// 1 question's answer. A probability has probability. A choice or score has value,
// confidence & the probability of each category or level.
type XAnalysisClassifyResponseResultAnalysisAnswer struct {
	QuestionID      string `json:"questionId" api:"required"`
	QuestionVersion string `json:"questionVersion" api:"required"`
	// Any of "choice", "score", "probability".
	Type       XAnalysisClassifyResponseResultAnalysisAnswerType `json:"type" api:"required"`
	Confidence float64                                           `json:"confidence"`
	// For a score. The levels the question named, from 0.
	Levels []string `json:"levels"`
	// Probability of each category, or of each level by index.
	Probabilities map[string]float64 `json:"probabilities"`
	// For a probability. The chance the answer is yes.
	Probability float64 `json:"probability"`
	// The chosen category, or the level from 0, which can have decimals, such as 1.4
	// between levels 1 & 2.
	Value XAnalysisClassifyResponseResultAnalysisAnswerValueUnion `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		QuestionID      respjson.Field
		QuestionVersion respjson.Field
		Type            respjson.Field
		Confidence      respjson.Field
		Levels          respjson.Field
		Probabilities   respjson.Field
		Probability     respjson.Field
		Value           respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XAnalysisClassifyResponseResultAnalysisAnswer) RawJSON() string { return r.JSON.raw }
func (r *XAnalysisClassifyResponseResultAnalysisAnswer) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type XAnalysisClassifyResponseResultAnalysisAnswerType string

const (
	XAnalysisClassifyResponseResultAnalysisAnswerTypeChoice      XAnalysisClassifyResponseResultAnalysisAnswerType = "choice"
	XAnalysisClassifyResponseResultAnalysisAnswerTypeScore       XAnalysisClassifyResponseResultAnalysisAnswerType = "score"
	XAnalysisClassifyResponseResultAnalysisAnswerTypeProbability XAnalysisClassifyResponseResultAnalysisAnswerType = "probability"
)

// XAnalysisClassifyResponseResultAnalysisAnswerValueUnion contains all possible
// properties and values from [float64], [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfFloat OfString]
type XAnalysisClassifyResponseResultAnalysisAnswerValueUnion struct {
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	JSON     struct {
		OfFloat  respjson.Field
		OfString respjson.Field
		raw      string
	} `json:"-"`
}

func (u XAnalysisClassifyResponseResultAnalysisAnswerValueUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v, nil)
	return
}

func (u XAnalysisClassifyResponseResultAnalysisAnswerValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v, nil)
	return
}

// Returns the unmodified JSON received from the API
func (u XAnalysisClassifyResponseResultAnalysisAnswerValueUnion) RawJSON() string { return u.JSON.raw }

func (r *XAnalysisClassifyResponseResultAnalysisAnswerValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// What the AI read beside the post's text.
type XAnalysisClassifyResponseResultAnalysisContextAvailability struct {
	// Any of "text_blocks", "summary", "not_supplied".
	Article XAnalysisClassifyResponseResultAnalysisContextAvailabilityArticle `json:"article" api:"required"`
	// Any of "available", "missing".
	Author XAnalysisClassifyResponseResultAnalysisContextAvailabilityAuthor `json:"author" api:"required"`
	// Any of "text_descriptions_only", "not_supplied".
	Media XAnalysisClassifyResponseResultAnalysisContextAvailabilityMedia `json:"media" api:"required"`
	// truncated when the post was cut to fit.
	//
	// Any of "full", "truncated".
	PostText XAnalysisClassifyResponseResultAnalysisContextAvailabilityPostText `json:"postText" api:"required"`
	// Any of "available", "missing", "not_applicable".
	Quote XAnalysisClassifyResponseResultAnalysisContextAvailabilityQuote `json:"quote" api:"required"`
	// Any of "available", "missing", "not_applicable".
	Reply XAnalysisClassifyResponseResultAnalysisContextAvailabilityReply `json:"reply" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Article     respjson.Field
		Author      respjson.Field
		Media       respjson.Field
		PostText    respjson.Field
		Quote       respjson.Field
		Reply       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XAnalysisClassifyResponseResultAnalysisContextAvailability) RawJSON() string {
	return r.JSON.raw
}
func (r *XAnalysisClassifyResponseResultAnalysisContextAvailability) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type XAnalysisClassifyResponseResultAnalysisContextAvailabilityArticle string

const (
	XAnalysisClassifyResponseResultAnalysisContextAvailabilityArticleTextBlocks  XAnalysisClassifyResponseResultAnalysisContextAvailabilityArticle = "text_blocks"
	XAnalysisClassifyResponseResultAnalysisContextAvailabilityArticleSummary     XAnalysisClassifyResponseResultAnalysisContextAvailabilityArticle = "summary"
	XAnalysisClassifyResponseResultAnalysisContextAvailabilityArticleNotSupplied XAnalysisClassifyResponseResultAnalysisContextAvailabilityArticle = "not_supplied"
)

type XAnalysisClassifyResponseResultAnalysisContextAvailabilityAuthor string

const (
	XAnalysisClassifyResponseResultAnalysisContextAvailabilityAuthorAvailable XAnalysisClassifyResponseResultAnalysisContextAvailabilityAuthor = "available"
	XAnalysisClassifyResponseResultAnalysisContextAvailabilityAuthorMissing   XAnalysisClassifyResponseResultAnalysisContextAvailabilityAuthor = "missing"
)

type XAnalysisClassifyResponseResultAnalysisContextAvailabilityMedia string

const (
	XAnalysisClassifyResponseResultAnalysisContextAvailabilityMediaTextDescriptionsOnly XAnalysisClassifyResponseResultAnalysisContextAvailabilityMedia = "text_descriptions_only"
	XAnalysisClassifyResponseResultAnalysisContextAvailabilityMediaNotSupplied          XAnalysisClassifyResponseResultAnalysisContextAvailabilityMedia = "not_supplied"
)

// truncated when the post was cut to fit.
type XAnalysisClassifyResponseResultAnalysisContextAvailabilityPostText string

const (
	XAnalysisClassifyResponseResultAnalysisContextAvailabilityPostTextFull      XAnalysisClassifyResponseResultAnalysisContextAvailabilityPostText = "full"
	XAnalysisClassifyResponseResultAnalysisContextAvailabilityPostTextTruncated XAnalysisClassifyResponseResultAnalysisContextAvailabilityPostText = "truncated"
)

type XAnalysisClassifyResponseResultAnalysisContextAvailabilityQuote string

const (
	XAnalysisClassifyResponseResultAnalysisContextAvailabilityQuoteAvailable     XAnalysisClassifyResponseResultAnalysisContextAvailabilityQuote = "available"
	XAnalysisClassifyResponseResultAnalysisContextAvailabilityQuoteMissing       XAnalysisClassifyResponseResultAnalysisContextAvailabilityQuote = "missing"
	XAnalysisClassifyResponseResultAnalysisContextAvailabilityQuoteNotApplicable XAnalysisClassifyResponseResultAnalysisContextAvailabilityQuote = "not_applicable"
)

type XAnalysisClassifyResponseResultAnalysisContextAvailabilityReply string

const (
	XAnalysisClassifyResponseResultAnalysisContextAvailabilityReplyAvailable     XAnalysisClassifyResponseResultAnalysisContextAvailabilityReply = "available"
	XAnalysisClassifyResponseResultAnalysisContextAvailabilityReplyMissing       XAnalysisClassifyResponseResultAnalysisContextAvailabilityReply = "missing"
	XAnalysisClassifyResponseResultAnalysisContextAvailabilityReplyNotApplicable XAnalysisClassifyResponseResultAnalysisContextAvailabilityReply = "not_applicable"
)

type XAnalysisClassifyResponseResultAnalysisQuestion struct {
	ID      string `json:"id" api:"required"`
	Version string `json:"version" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Version     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XAnalysisClassifyResponseResultAnalysisQuestion) RawJSON() string { return r.JSON.raw }
func (r *XAnalysisClassifyResponseResultAnalysisQuestion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// XAnalysisClassifyResponseResultAnswersUnion contains all possible properties and
// values from [float64], [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfFloat OfString]
type XAnalysisClassifyResponseResultAnswersUnion struct {
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	JSON     struct {
		OfFloat  respjson.Field
		OfString respjson.Field
		raw      string
	} `json:"-"`
}

func (u XAnalysisClassifyResponseResultAnswersUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v, nil)
	return
}

func (u XAnalysisClassifyResponseResultAnswersUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v, nil)
	return
}

// Returns the unmodified JSON received from the API
func (u XAnalysisClassifyResponseResultAnswersUnion) RawJSON() string { return u.JSON.raw }

func (r *XAnalysisClassifyResponseResultAnswersUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// How the post's answers compare with its row in baseline, as the AI Actors
// compare a run with an earlier dataset.
type XAnalysisClassifyResponseResultMonitor struct {
	// Always null. The baseline is the rows you sent.
	BaselineDatasetID json.RawMessage `json:"baselineDatasetId" api:"required"`
	// IDs of the questions whose decision changed.
	ChangedQuestionIDs []string `json:"changedQuestionIds" api:"required"`
	// Each changed decision, earlier & now.
	Changes []XAnalysisClassifyResponseResultMonitorChange `json:"changes" api:"required"`
	// Hash of the questions, targets & context. Rows with another hash are
	// not_comparable.
	ConfigurationFingerprint string `json:"configurationFingerprint" api:"required"`
	// first_run without baseline. new_to_baseline for a post baseline lacks. changed
	// when an answer clearly left its earlier decision. not_comparable when other
	// settings made the earlier row or it answers none of these questions.
	//
	// Any of "first_run", "new_to_baseline", "unchanged", "changed", "not_comparable",
	// "analysis_unavailable".
	Status XAnalysisClassifyResponseResultMonitorStatus `json:"status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BaselineDatasetID        respjson.Field
		ChangedQuestionIDs       respjson.Field
		Changes                  respjson.Field
		ConfigurationFingerprint respjson.Field
		Status                   respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XAnalysisClassifyResponseResultMonitor) RawJSON() string { return r.JSON.raw }
func (r *XAnalysisClassifyResponseResultMonitor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type XAnalysisClassifyResponseResultMonitorChange struct {
	Current    string `json:"current" api:"required"`
	Previous   string `json:"previous" api:"required"`
	QuestionID string `json:"questionId" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Current     respjson.Field
		Previous    respjson.Field
		QuestionID  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XAnalysisClassifyResponseResultMonitorChange) RawJSON() string { return r.JSON.raw }
func (r *XAnalysisClassifyResponseResultMonitorChange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// first_run without baseline. new_to_baseline for a post baseline lacks. changed
// when an answer clearly left its earlier decision. not_comparable when other
// settings made the earlier row or it answers none of these questions.
type XAnalysisClassifyResponseResultMonitorStatus string

const (
	XAnalysisClassifyResponseResultMonitorStatusFirstRun            XAnalysisClassifyResponseResultMonitorStatus = "first_run"
	XAnalysisClassifyResponseResultMonitorStatusNewToBaseline       XAnalysisClassifyResponseResultMonitorStatus = "new_to_baseline"
	XAnalysisClassifyResponseResultMonitorStatusUnchanged           XAnalysisClassifyResponseResultMonitorStatus = "unchanged"
	XAnalysisClassifyResponseResultMonitorStatusChanged             XAnalysisClassifyResponseResultMonitorStatus = "changed"
	XAnalysisClassifyResponseResultMonitorStatusNotComparable       XAnalysisClassifyResponseResultMonitorStatus = "not_comparable"
	XAnalysisClassifyResponseResultMonitorStatusAnalysisUnavailable XAnalysisClassifyResponseResultMonitorStatus = "analysis_unavailable"
)

// Only from /x/analysis/viral-score. send_it from 70, edit_first from 40, else
// sleep_on_it.
type XAnalysisClassifyResponseResultViralVerdict string

const (
	XAnalysisClassifyResponseResultViralVerdictSendIt    XAnalysisClassifyResponseResultViralVerdict = "send_it"
	XAnalysisClassifyResponseResultViralVerdictEditFirst XAnalysisClassifyResponseResultViralVerdict = "edit_first"
	XAnalysisClassifyResponseResultViralVerdictSleepOnIt XAnalysisClassifyResponseResultViralVerdict = "sleep_on_it"
)

// A post without an analysis. It costs nothing.
type XAnalysisClassifyResponseUnanalyzed struct {
	// Post ID. Your own texts have the IDs text:1, text:2 & so on.
	ID string `json:"id" api:"required"`
	// post_unavailable when X returns no such post. missing_text for a post without
	// text. insufficient_credits when credits ran out. context_limit when the settings
	// leave the post no room. The other reasons mean the analysis failed. Send the
	// post again.
	Reason string `json:"reason" api:"required"`
	// Any of "failed", "skipped".
	Status XAnalysisClassifyResponseUnanalyzedStatus `json:"status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Reason      respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XAnalysisClassifyResponseUnanalyzed) RawJSON() string { return r.JSON.raw }
func (r *XAnalysisClassifyResponseUnanalyzed) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type XAnalysisClassifyResponseUnanalyzedStatus string

const (
	XAnalysisClassifyResponseUnanalyzedStatusFailed  XAnalysisClassifyResponseUnanalyzedStatus = "failed"
	XAnalysisClassifyResponseUnanalyzedStatusSkipped XAnalysisClassifyResponseUnanalyzedStatus = "skipped"
)

// Analyzed posts, the posts without an analysis & the totals of the answers. Each
// row in results costs 2 credits.
type XAnalysisMarketSignalsResponse struct {
	// Totals of the answers across the analyzed posts.
	AnalysisSummary XAnalysisMarketSignalsResponseAnalysisSummary `json:"analysisSummary" api:"required"`
	// Whether the search has more posts. Send next_cursor as cursor.
	HasNextPage bool `json:"has_next_page" api:"required"`
	// Cursor of the search's next page, else empty.
	NextCursor string `json:"next_cursor" api:"required"`
	// Each analyzed post with its answers.
	Results []XAnalysisMarketSignalsResponseResult `json:"results" api:"required"`
	// Each post without an analysis, by ID & reason. These cost nothing.
	Unanalyzed []XAnalysisMarketSignalsResponseUnanalyzed `json:"unanalyzed" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AnalysisSummary respjson.Field
		HasNextPage     respjson.Field
		NextCursor      respjson.Field
		Results         respjson.Field
		Unanalyzed      respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XAnalysisMarketSignalsResponse) RawJSON() string { return r.JSON.raw }
func (r *XAnalysisMarketSignalsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// Totals of the answers across the analyzed posts.
type XAnalysisMarketSignalsResponseAnalysisSummary struct {
	// Posts & answers per cashtag, most posts first, up to 20. signal is present for
	// the market questions.
	Cashtags []XAnalysisMarketSignalsResponseAnalysisSummaryCashtag `json:"cashtags" api:"required"`
	// Likes, reposts, replies & quotes of the analyzed posts.
	Engagement int64 `json:"engagement" api:"required"`
	// Posts per comparison status, & up to 50 changed posts with their changes. Null
	// when the answer holds no post.
	Monitor XAnalysisMarketSignalsResponseAnalysisSummaryMonitor `json:"monitor" api:"required"`
	// Totals by question ID. A choice has counts, shares, engagementShares & top posts
	// per category. A score has mean, engagementWeightedMean & levels. A probability
	// has mean, yes & no.
	Questions map[string]PublicMetadata `json:"questions" api:"required"`
	// Posts by analysis outcome.
	Rows          XAnalysisMarketSignalsResponseAnalysisSummaryRows `json:"rows" api:"required"`
	SchemaVersion int64                                             `json:"schemaVersion" api:"required"`
	// Posts per linked hostname, most posts first, up to 20.
	SourceDomains []XAnalysisMarketSignalsResponseAnalysisSummarySourceDomain `json:"sourceDomains" api:"required"`
	// Mentions & answers per target in analysis.targets.
	Targets []XAnalysisMarketSignalsResponseAnalysisSummaryTarget `json:"targets" api:"required"`
	// Only from /x/analysis/viral-score. Scored posts, the average score, verdict
	// counts, the scores against actual engagement, a report per author & the
	// leaderboards.
	Viral PublicMetadata `json:"viral"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Cashtags      respjson.Field
		Engagement    respjson.Field
		Monitor       respjson.Field
		Questions     respjson.Field
		Rows          respjson.Field
		SchemaVersion respjson.Field
		SourceDomains respjson.Field
		Targets       respjson.Field
		Viral         respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XAnalysisMarketSignalsResponseAnalysisSummary) RawJSON() string { return r.JSON.raw }
func (r *XAnalysisMarketSignalsResponseAnalysisSummary) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type XAnalysisMarketSignalsResponseAnalysisSummaryCashtag struct {
	Cashtag string `json:"cashtag" api:"required"`
	// Category counts of the cashtag's posts by question ID.
	Choices PublicMetadata                                             `json:"choices" api:"required"`
	Rows    int64                                                      `json:"rows" api:"required"`
	Signal  XAnalysisMarketSignalsResponseAnalysisSummaryCashtagSignal `json:"signal"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Cashtag     respjson.Field
		Choices     respjson.Field
		Rows        respjson.Field
		Signal      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XAnalysisMarketSignalsResponseAnalysisSummaryCashtag) RawJSON() string { return r.JSON.raw }
func (r *XAnalysisMarketSignalsResponseAnalysisSummaryCashtag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type XAnalysisMarketSignalsResponseAnalysisSummaryCashtagSignal struct {
	Bearish int64 `json:"bearish" api:"required"`
	Bullish int64 `json:"bullish" api:"required"`
	// Bullish minus bearish posts, over all posts.
	Score float64 `json:"score" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Bearish     respjson.Field
		Bullish     respjson.Field
		Score       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XAnalysisMarketSignalsResponseAnalysisSummaryCashtagSignal) RawJSON() string {
	return r.JSON.raw
}
func (r *XAnalysisMarketSignalsResponseAnalysisSummaryCashtagSignal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// Posts per comparison status, & up to 50 changed posts with their changes. Null
// when the answer holds no post.
type XAnalysisMarketSignalsResponseAnalysisSummaryMonitor struct {
	ChangedRows []XAnalysisMarketSignalsResponseAnalysisSummaryMonitorChangedRow `json:"changedRows" api:"required"`
	// Posts per monitor status.
	Statuses map[string]int64 `json:"statuses" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChangedRows respjson.Field
		Statuses    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XAnalysisMarketSignalsResponseAnalysisSummaryMonitor) RawJSON() string { return r.JSON.raw }
func (r *XAnalysisMarketSignalsResponseAnalysisSummaryMonitor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type XAnalysisMarketSignalsResponseAnalysisSummaryMonitorChangedRow struct {
	Changes []XAnalysisMarketSignalsResponseAnalysisSummaryMonitorChangedRowChange `json:"changes" api:"required"`
	TweetID string                                                                 `json:"tweetId" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Changes     respjson.Field
		TweetID     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XAnalysisMarketSignalsResponseAnalysisSummaryMonitorChangedRow) RawJSON() string {
	return r.JSON.raw
}
func (r *XAnalysisMarketSignalsResponseAnalysisSummaryMonitorChangedRow) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type XAnalysisMarketSignalsResponseAnalysisSummaryMonitorChangedRowChange struct {
	Current    string `json:"current" api:"required"`
	Previous   string `json:"previous" api:"required"`
	QuestionID string `json:"questionId" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Current     respjson.Field
		Previous    respjson.Field
		QuestionID  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XAnalysisMarketSignalsResponseAnalysisSummaryMonitorChangedRowChange) RawJSON() string {
	return r.JSON.raw
}
func (r *XAnalysisMarketSignalsResponseAnalysisSummaryMonitorChangedRowChange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// Posts by analysis outcome.
type XAnalysisMarketSignalsResponseAnalysisSummaryRows struct {
	Analyzed int64 `json:"analyzed" api:"required"`
	Failed   int64 `json:"failed" api:"required"`
	Skipped  int64 `json:"skipped" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Analyzed    respjson.Field
		Failed      respjson.Field
		Skipped     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XAnalysisMarketSignalsResponseAnalysisSummaryRows) RawJSON() string { return r.JSON.raw }
func (r *XAnalysisMarketSignalsResponseAnalysisSummaryRows) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type XAnalysisMarketSignalsResponseAnalysisSummarySourceDomain struct {
	Domain string `json:"domain" api:"required"`
	Posts  int64  `json:"posts" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Domain      respjson.Field
		Posts       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XAnalysisMarketSignalsResponseAnalysisSummarySourceDomain) RawJSON() string {
	return r.JSON.raw
}
func (r *XAnalysisMarketSignalsResponseAnalysisSummarySourceDomain) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type XAnalysisMarketSignalsResponseAnalysisSummaryTarget struct {
	// Category counts by question ID.
	Choices    PublicMetadata `json:"choices" api:"required"`
	Engagement int64          `json:"engagement" api:"required"`
	Mentions   int64          `json:"mentions" api:"required"`
	Name       string         `json:"name" api:"required"`
	Share      float64        `json:"share" api:"required"`
	// The most engaged posts per question ID & category.
	Top PublicMetadata `json:"top" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Choices     respjson.Field
		Engagement  respjson.Field
		Mentions    respjson.Field
		Name        respjson.Field
		Share       respjson.Field
		Top         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XAnalysisMarketSignalsResponseAnalysisSummaryTarget) RawJSON() string { return r.JSON.raw }
func (r *XAnalysisMarketSignalsResponseAnalysisSummaryTarget) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// 1 analyzed post with its answers.
type XAnalysisMarketSignalsResponseResult struct {
	// The answers in full, with their probabilities.
	Analysis XAnalysisMarketSignalsResponseResultAnalysis `json:"analysis" api:"required"`
	// Each question's decision by question ID. A choice gives its category, a score
	// its level from 0, which can have decimals, & a probability its value from 0 to
	// 1.
	Answers map[string]XAnalysisMarketSignalsResponseResultAnswersUnion `json:"answers" api:"required"`
	// Cashtags in the post's text, in upper case.
	Cashtags []string `json:"cashtags" api:"required"`
	// How the post's answers compare with its row in baseline, as the AI Actors
	// compare a run with an earlier dataset.
	Monitor XAnalysisMarketSignalsResponseResultMonitor `json:"monitor" api:"required"`
	// Hostnames the post links to, without t.co.
	SourceDomains []string `json:"sourceDomains" api:"required"`
	// The post, as /x/tweets/search returns it. Your own texts have the IDs text:1,
	// text:2 & so on.
	Tweet shared.SearchTweet `json:"tweet" api:"required"`
	// Only from /x/analysis/viral-score. The score with its weights version, the caps
	// that applied & the engagement comparison.
	Viral PublicMetadata `json:"viral"`
	// Only from /x/analysis/viral-score. log10 of 1 plus the weighted engagement per
	// 1,000 followers. Null without a follower count.
	ViralActualEngagementRate float64 `json:"viralActualEngagementRate" api:"nullable"`
	// Only from /x/analysis/viral-score. The post's public counts under X's published
	// ranking weights, per 1,000 views or followers.
	ViralAlgorithmScore float64 `json:"viralAlgorithmScore" api:"nullable"`
	// Only from /x/analysis/viral-score. How viral the wording reads, or null when a
	// trait has no answer.
	ViralScore int64 `json:"viralScore" api:"nullable"`
	// Only from /x/analysis/viral-score. send_it from 70, edit_first from 40, else
	// sleep_on_it.
	//
	// Any of "send_it", "edit_first", "sleep_on_it".
	ViralVerdict XAnalysisMarketSignalsResponseResultViralVerdict `json:"viralVerdict" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Analysis                  respjson.Field
		Answers                   respjson.Field
		Cashtags                  respjson.Field
		Monitor                   respjson.Field
		SourceDomains             respjson.Field
		Tweet                     respjson.Field
		Viral                     respjson.Field
		ViralActualEngagementRate respjson.Field
		ViralAlgorithmScore       respjson.Field
		ViralScore                respjson.Field
		ViralVerdict              respjson.Field
		ExtraFields               map[string]respjson.Field
		raw                       string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XAnalysisMarketSignalsResponseResult) RawJSON() string { return r.JSON.raw }
func (r *XAnalysisMarketSignalsResponseResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// The answers in full, with their probabilities.
type XAnalysisMarketSignalsResponseResultAnalysis struct {
	// 1 answer per question, in question order.
	Answers []XAnalysisMarketSignalsResponseResultAnalysisAnswer `json:"answers" api:"required"`
	// What the AI read beside the post's text.
	ContextAvailability XAnalysisMarketSignalsResponseResultAnalysisContextAvailability `json:"contextAvailability" api:"required"`
	// Bytes of post & settings the AI read.
	ContextBytes int64 `json:"contextBytes" api:"required"`
	// The ID & version of each question asked.
	Questions []XAnalysisMarketSignalsResponseResultAnalysisQuestion `json:"questions" api:"required"`
	Status    constant.Succeeded                                     `json:"status" default:"succeeded"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Answers             respjson.Field
		ContextAvailability respjson.Field
		ContextBytes        respjson.Field
		Questions           respjson.Field
		Status              respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XAnalysisMarketSignalsResponseResultAnalysis) RawJSON() string { return r.JSON.raw }
func (r *XAnalysisMarketSignalsResponseResultAnalysis) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// 1 question's answer. A probability has probability. A choice or score has value,
// confidence & the probability of each category or level.
type XAnalysisMarketSignalsResponseResultAnalysisAnswer struct {
	QuestionID      string `json:"questionId" api:"required"`
	QuestionVersion string `json:"questionVersion" api:"required"`
	// Any of "choice", "score", "probability".
	Type       XAnalysisMarketSignalsResponseResultAnalysisAnswerType `json:"type" api:"required"`
	Confidence float64                                                `json:"confidence"`
	// For a score. The levels the question named, from 0.
	Levels []string `json:"levels"`
	// Probability of each category, or of each level by index.
	Probabilities map[string]float64 `json:"probabilities"`
	// For a probability. The chance the answer is yes.
	Probability float64 `json:"probability"`
	// The chosen category, or the level from 0, which can have decimals, such as 1.4
	// between levels 1 & 2.
	Value XAnalysisMarketSignalsResponseResultAnalysisAnswerValueUnion `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		QuestionID      respjson.Field
		QuestionVersion respjson.Field
		Type            respjson.Field
		Confidence      respjson.Field
		Levels          respjson.Field
		Probabilities   respjson.Field
		Probability     respjson.Field
		Value           respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XAnalysisMarketSignalsResponseResultAnalysisAnswer) RawJSON() string { return r.JSON.raw }
func (r *XAnalysisMarketSignalsResponseResultAnalysisAnswer) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type XAnalysisMarketSignalsResponseResultAnalysisAnswerType string

const (
	XAnalysisMarketSignalsResponseResultAnalysisAnswerTypeChoice      XAnalysisMarketSignalsResponseResultAnalysisAnswerType = "choice"
	XAnalysisMarketSignalsResponseResultAnalysisAnswerTypeScore       XAnalysisMarketSignalsResponseResultAnalysisAnswerType = "score"
	XAnalysisMarketSignalsResponseResultAnalysisAnswerTypeProbability XAnalysisMarketSignalsResponseResultAnalysisAnswerType = "probability"
)

// XAnalysisMarketSignalsResponseResultAnalysisAnswerValueUnion contains all
// possible properties and values from [float64], [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfFloat OfString]
type XAnalysisMarketSignalsResponseResultAnalysisAnswerValueUnion struct {
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	JSON     struct {
		OfFloat  respjson.Field
		OfString respjson.Field
		raw      string
	} `json:"-"`
}

func (u XAnalysisMarketSignalsResponseResultAnalysisAnswerValueUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v, nil)
	return
}

func (u XAnalysisMarketSignalsResponseResultAnalysisAnswerValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v, nil)
	return
}

// Returns the unmodified JSON received from the API
func (u XAnalysisMarketSignalsResponseResultAnalysisAnswerValueUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *XAnalysisMarketSignalsResponseResultAnalysisAnswerValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// What the AI read beside the post's text.
type XAnalysisMarketSignalsResponseResultAnalysisContextAvailability struct {
	// Any of "text_blocks", "summary", "not_supplied".
	Article XAnalysisMarketSignalsResponseResultAnalysisContextAvailabilityArticle `json:"article" api:"required"`
	// Any of "available", "missing".
	Author XAnalysisMarketSignalsResponseResultAnalysisContextAvailabilityAuthor `json:"author" api:"required"`
	// Any of "text_descriptions_only", "not_supplied".
	Media XAnalysisMarketSignalsResponseResultAnalysisContextAvailabilityMedia `json:"media" api:"required"`
	// truncated when the post was cut to fit.
	//
	// Any of "full", "truncated".
	PostText XAnalysisMarketSignalsResponseResultAnalysisContextAvailabilityPostText `json:"postText" api:"required"`
	// Any of "available", "missing", "not_applicable".
	Quote XAnalysisMarketSignalsResponseResultAnalysisContextAvailabilityQuote `json:"quote" api:"required"`
	// Any of "available", "missing", "not_applicable".
	Reply XAnalysisMarketSignalsResponseResultAnalysisContextAvailabilityReply `json:"reply" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Article     respjson.Field
		Author      respjson.Field
		Media       respjson.Field
		PostText    respjson.Field
		Quote       respjson.Field
		Reply       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XAnalysisMarketSignalsResponseResultAnalysisContextAvailability) RawJSON() string {
	return r.JSON.raw
}
func (r *XAnalysisMarketSignalsResponseResultAnalysisContextAvailability) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type XAnalysisMarketSignalsResponseResultAnalysisContextAvailabilityArticle string

const (
	XAnalysisMarketSignalsResponseResultAnalysisContextAvailabilityArticleTextBlocks  XAnalysisMarketSignalsResponseResultAnalysisContextAvailabilityArticle = "text_blocks"
	XAnalysisMarketSignalsResponseResultAnalysisContextAvailabilityArticleSummary     XAnalysisMarketSignalsResponseResultAnalysisContextAvailabilityArticle = "summary"
	XAnalysisMarketSignalsResponseResultAnalysisContextAvailabilityArticleNotSupplied XAnalysisMarketSignalsResponseResultAnalysisContextAvailabilityArticle = "not_supplied"
)

type XAnalysisMarketSignalsResponseResultAnalysisContextAvailabilityAuthor string

const (
	XAnalysisMarketSignalsResponseResultAnalysisContextAvailabilityAuthorAvailable XAnalysisMarketSignalsResponseResultAnalysisContextAvailabilityAuthor = "available"
	XAnalysisMarketSignalsResponseResultAnalysisContextAvailabilityAuthorMissing   XAnalysisMarketSignalsResponseResultAnalysisContextAvailabilityAuthor = "missing"
)

type XAnalysisMarketSignalsResponseResultAnalysisContextAvailabilityMedia string

const (
	XAnalysisMarketSignalsResponseResultAnalysisContextAvailabilityMediaTextDescriptionsOnly XAnalysisMarketSignalsResponseResultAnalysisContextAvailabilityMedia = "text_descriptions_only"
	XAnalysisMarketSignalsResponseResultAnalysisContextAvailabilityMediaNotSupplied          XAnalysisMarketSignalsResponseResultAnalysisContextAvailabilityMedia = "not_supplied"
)

// truncated when the post was cut to fit.
type XAnalysisMarketSignalsResponseResultAnalysisContextAvailabilityPostText string

const (
	XAnalysisMarketSignalsResponseResultAnalysisContextAvailabilityPostTextFull      XAnalysisMarketSignalsResponseResultAnalysisContextAvailabilityPostText = "full"
	XAnalysisMarketSignalsResponseResultAnalysisContextAvailabilityPostTextTruncated XAnalysisMarketSignalsResponseResultAnalysisContextAvailabilityPostText = "truncated"
)

type XAnalysisMarketSignalsResponseResultAnalysisContextAvailabilityQuote string

const (
	XAnalysisMarketSignalsResponseResultAnalysisContextAvailabilityQuoteAvailable     XAnalysisMarketSignalsResponseResultAnalysisContextAvailabilityQuote = "available"
	XAnalysisMarketSignalsResponseResultAnalysisContextAvailabilityQuoteMissing       XAnalysisMarketSignalsResponseResultAnalysisContextAvailabilityQuote = "missing"
	XAnalysisMarketSignalsResponseResultAnalysisContextAvailabilityQuoteNotApplicable XAnalysisMarketSignalsResponseResultAnalysisContextAvailabilityQuote = "not_applicable"
)

type XAnalysisMarketSignalsResponseResultAnalysisContextAvailabilityReply string

const (
	XAnalysisMarketSignalsResponseResultAnalysisContextAvailabilityReplyAvailable     XAnalysisMarketSignalsResponseResultAnalysisContextAvailabilityReply = "available"
	XAnalysisMarketSignalsResponseResultAnalysisContextAvailabilityReplyMissing       XAnalysisMarketSignalsResponseResultAnalysisContextAvailabilityReply = "missing"
	XAnalysisMarketSignalsResponseResultAnalysisContextAvailabilityReplyNotApplicable XAnalysisMarketSignalsResponseResultAnalysisContextAvailabilityReply = "not_applicable"
)

type XAnalysisMarketSignalsResponseResultAnalysisQuestion struct {
	ID      string `json:"id" api:"required"`
	Version string `json:"version" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Version     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XAnalysisMarketSignalsResponseResultAnalysisQuestion) RawJSON() string { return r.JSON.raw }
func (r *XAnalysisMarketSignalsResponseResultAnalysisQuestion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// XAnalysisMarketSignalsResponseResultAnswersUnion contains all possible
// properties and values from [float64], [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfFloat OfString]
type XAnalysisMarketSignalsResponseResultAnswersUnion struct {
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	JSON     struct {
		OfFloat  respjson.Field
		OfString respjson.Field
		raw      string
	} `json:"-"`
}

func (u XAnalysisMarketSignalsResponseResultAnswersUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v, nil)
	return
}

func (u XAnalysisMarketSignalsResponseResultAnswersUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v, nil)
	return
}

// Returns the unmodified JSON received from the API
func (u XAnalysisMarketSignalsResponseResultAnswersUnion) RawJSON() string { return u.JSON.raw }

func (r *XAnalysisMarketSignalsResponseResultAnswersUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// How the post's answers compare with its row in baseline, as the AI Actors
// compare a run with an earlier dataset.
type XAnalysisMarketSignalsResponseResultMonitor struct {
	// Always null. The baseline is the rows you sent.
	BaselineDatasetID json.RawMessage `json:"baselineDatasetId" api:"required"`
	// IDs of the questions whose decision changed.
	ChangedQuestionIDs []string `json:"changedQuestionIds" api:"required"`
	// Each changed decision, earlier & now.
	Changes []XAnalysisMarketSignalsResponseResultMonitorChange `json:"changes" api:"required"`
	// Hash of the questions, targets & context. Rows with another hash are
	// not_comparable.
	ConfigurationFingerprint string `json:"configurationFingerprint" api:"required"`
	// first_run without baseline. new_to_baseline for a post baseline lacks. changed
	// when an answer clearly left its earlier decision. not_comparable when other
	// settings made the earlier row or it answers none of these questions.
	//
	// Any of "first_run", "new_to_baseline", "unchanged", "changed", "not_comparable",
	// "analysis_unavailable".
	Status XAnalysisMarketSignalsResponseResultMonitorStatus `json:"status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BaselineDatasetID        respjson.Field
		ChangedQuestionIDs       respjson.Field
		Changes                  respjson.Field
		ConfigurationFingerprint respjson.Field
		Status                   respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XAnalysisMarketSignalsResponseResultMonitor) RawJSON() string { return r.JSON.raw }
func (r *XAnalysisMarketSignalsResponseResultMonitor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type XAnalysisMarketSignalsResponseResultMonitorChange struct {
	Current    string `json:"current" api:"required"`
	Previous   string `json:"previous" api:"required"`
	QuestionID string `json:"questionId" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Current     respjson.Field
		Previous    respjson.Field
		QuestionID  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XAnalysisMarketSignalsResponseResultMonitorChange) RawJSON() string { return r.JSON.raw }
func (r *XAnalysisMarketSignalsResponseResultMonitorChange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// first_run without baseline. new_to_baseline for a post baseline lacks. changed
// when an answer clearly left its earlier decision. not_comparable when other
// settings made the earlier row or it answers none of these questions.
type XAnalysisMarketSignalsResponseResultMonitorStatus string

const (
	XAnalysisMarketSignalsResponseResultMonitorStatusFirstRun            XAnalysisMarketSignalsResponseResultMonitorStatus = "first_run"
	XAnalysisMarketSignalsResponseResultMonitorStatusNewToBaseline       XAnalysisMarketSignalsResponseResultMonitorStatus = "new_to_baseline"
	XAnalysisMarketSignalsResponseResultMonitorStatusUnchanged           XAnalysisMarketSignalsResponseResultMonitorStatus = "unchanged"
	XAnalysisMarketSignalsResponseResultMonitorStatusChanged             XAnalysisMarketSignalsResponseResultMonitorStatus = "changed"
	XAnalysisMarketSignalsResponseResultMonitorStatusNotComparable       XAnalysisMarketSignalsResponseResultMonitorStatus = "not_comparable"
	XAnalysisMarketSignalsResponseResultMonitorStatusAnalysisUnavailable XAnalysisMarketSignalsResponseResultMonitorStatus = "analysis_unavailable"
)

// Only from /x/analysis/viral-score. send_it from 70, edit_first from 40, else
// sleep_on_it.
type XAnalysisMarketSignalsResponseResultViralVerdict string

const (
	XAnalysisMarketSignalsResponseResultViralVerdictSendIt    XAnalysisMarketSignalsResponseResultViralVerdict = "send_it"
	XAnalysisMarketSignalsResponseResultViralVerdictEditFirst XAnalysisMarketSignalsResponseResultViralVerdict = "edit_first"
	XAnalysisMarketSignalsResponseResultViralVerdictSleepOnIt XAnalysisMarketSignalsResponseResultViralVerdict = "sleep_on_it"
)

// A post without an analysis. It costs nothing.
type XAnalysisMarketSignalsResponseUnanalyzed struct {
	// Post ID. Your own texts have the IDs text:1, text:2 & so on.
	ID string `json:"id" api:"required"`
	// post_unavailable when X returns no such post. missing_text for a post without
	// text. insufficient_credits when credits ran out. context_limit when the settings
	// leave the post no room. The other reasons mean the analysis failed. Send the
	// post again.
	Reason string `json:"reason" api:"required"`
	// Any of "failed", "skipped".
	Status XAnalysisMarketSignalsResponseUnanalyzedStatus `json:"status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Reason      respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XAnalysisMarketSignalsResponseUnanalyzed) RawJSON() string { return r.JSON.raw }
func (r *XAnalysisMarketSignalsResponseUnanalyzed) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type XAnalysisMarketSignalsResponseUnanalyzedStatus string

const (
	XAnalysisMarketSignalsResponseUnanalyzedStatusFailed  XAnalysisMarketSignalsResponseUnanalyzedStatus = "failed"
	XAnalysisMarketSignalsResponseUnanalyzedStatusSkipped XAnalysisMarketSignalsResponseUnanalyzedStatus = "skipped"
)

// Analyzed posts, the posts without an analysis & the totals of the answers. Each
// row in results costs 2 credits.
type XAnalysisNewsResponse struct {
	// Totals of the answers across the analyzed posts.
	AnalysisSummary XAnalysisNewsResponseAnalysisSummary `json:"analysisSummary" api:"required"`
	// Whether the search has more posts. Send next_cursor as cursor.
	HasNextPage bool `json:"has_next_page" api:"required"`
	// Cursor of the search's next page, else empty.
	NextCursor string `json:"next_cursor" api:"required"`
	// Each analyzed post with its answers.
	Results []XAnalysisNewsResponseResult `json:"results" api:"required"`
	// Each post without an analysis, by ID & reason. These cost nothing.
	Unanalyzed []XAnalysisNewsResponseUnanalyzed `json:"unanalyzed" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AnalysisSummary respjson.Field
		HasNextPage     respjson.Field
		NextCursor      respjson.Field
		Results         respjson.Field
		Unanalyzed      respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XAnalysisNewsResponse) RawJSON() string { return r.JSON.raw }
func (r *XAnalysisNewsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// Totals of the answers across the analyzed posts.
type XAnalysisNewsResponseAnalysisSummary struct {
	// Posts & answers per cashtag, most posts first, up to 20. signal is present for
	// the market questions.
	Cashtags []XAnalysisNewsResponseAnalysisSummaryCashtag `json:"cashtags" api:"required"`
	// Likes, reposts, replies & quotes of the analyzed posts.
	Engagement int64 `json:"engagement" api:"required"`
	// Posts per comparison status, & up to 50 changed posts with their changes. Null
	// when the answer holds no post.
	Monitor XAnalysisNewsResponseAnalysisSummaryMonitor `json:"monitor" api:"required"`
	// Totals by question ID. A choice has counts, shares, engagementShares & top posts
	// per category. A score has mean, engagementWeightedMean & levels. A probability
	// has mean, yes & no.
	Questions map[string]PublicMetadata `json:"questions" api:"required"`
	// Posts by analysis outcome.
	Rows          XAnalysisNewsResponseAnalysisSummaryRows `json:"rows" api:"required"`
	SchemaVersion int64                                    `json:"schemaVersion" api:"required"`
	// Posts per linked hostname, most posts first, up to 20.
	SourceDomains []XAnalysisNewsResponseAnalysisSummarySourceDomain `json:"sourceDomains" api:"required"`
	// Mentions & answers per target in analysis.targets.
	Targets []XAnalysisNewsResponseAnalysisSummaryTarget `json:"targets" api:"required"`
	// Only from /x/analysis/viral-score. Scored posts, the average score, verdict
	// counts, the scores against actual engagement, a report per author & the
	// leaderboards.
	Viral PublicMetadata `json:"viral"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Cashtags      respjson.Field
		Engagement    respjson.Field
		Monitor       respjson.Field
		Questions     respjson.Field
		Rows          respjson.Field
		SchemaVersion respjson.Field
		SourceDomains respjson.Field
		Targets       respjson.Field
		Viral         respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XAnalysisNewsResponseAnalysisSummary) RawJSON() string { return r.JSON.raw }
func (r *XAnalysisNewsResponseAnalysisSummary) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type XAnalysisNewsResponseAnalysisSummaryCashtag struct {
	Cashtag string `json:"cashtag" api:"required"`
	// Category counts of the cashtag's posts by question ID.
	Choices PublicMetadata                                    `json:"choices" api:"required"`
	Rows    int64                                             `json:"rows" api:"required"`
	Signal  XAnalysisNewsResponseAnalysisSummaryCashtagSignal `json:"signal"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Cashtag     respjson.Field
		Choices     respjson.Field
		Rows        respjson.Field
		Signal      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XAnalysisNewsResponseAnalysisSummaryCashtag) RawJSON() string { return r.JSON.raw }
func (r *XAnalysisNewsResponseAnalysisSummaryCashtag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type XAnalysisNewsResponseAnalysisSummaryCashtagSignal struct {
	Bearish int64 `json:"bearish" api:"required"`
	Bullish int64 `json:"bullish" api:"required"`
	// Bullish minus bearish posts, over all posts.
	Score float64 `json:"score" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Bearish     respjson.Field
		Bullish     respjson.Field
		Score       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XAnalysisNewsResponseAnalysisSummaryCashtagSignal) RawJSON() string { return r.JSON.raw }
func (r *XAnalysisNewsResponseAnalysisSummaryCashtagSignal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// Posts per comparison status, & up to 50 changed posts with their changes. Null
// when the answer holds no post.
type XAnalysisNewsResponseAnalysisSummaryMonitor struct {
	ChangedRows []XAnalysisNewsResponseAnalysisSummaryMonitorChangedRow `json:"changedRows" api:"required"`
	// Posts per monitor status.
	Statuses map[string]int64 `json:"statuses" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChangedRows respjson.Field
		Statuses    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XAnalysisNewsResponseAnalysisSummaryMonitor) RawJSON() string { return r.JSON.raw }
func (r *XAnalysisNewsResponseAnalysisSummaryMonitor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type XAnalysisNewsResponseAnalysisSummaryMonitorChangedRow struct {
	Changes []XAnalysisNewsResponseAnalysisSummaryMonitorChangedRowChange `json:"changes" api:"required"`
	TweetID string                                                        `json:"tweetId" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Changes     respjson.Field
		TweetID     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XAnalysisNewsResponseAnalysisSummaryMonitorChangedRow) RawJSON() string { return r.JSON.raw }
func (r *XAnalysisNewsResponseAnalysisSummaryMonitorChangedRow) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type XAnalysisNewsResponseAnalysisSummaryMonitorChangedRowChange struct {
	Current    string `json:"current" api:"required"`
	Previous   string `json:"previous" api:"required"`
	QuestionID string `json:"questionId" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Current     respjson.Field
		Previous    respjson.Field
		QuestionID  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XAnalysisNewsResponseAnalysisSummaryMonitorChangedRowChange) RawJSON() string {
	return r.JSON.raw
}
func (r *XAnalysisNewsResponseAnalysisSummaryMonitorChangedRowChange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// Posts by analysis outcome.
type XAnalysisNewsResponseAnalysisSummaryRows struct {
	Analyzed int64 `json:"analyzed" api:"required"`
	Failed   int64 `json:"failed" api:"required"`
	Skipped  int64 `json:"skipped" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Analyzed    respjson.Field
		Failed      respjson.Field
		Skipped     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XAnalysisNewsResponseAnalysisSummaryRows) RawJSON() string { return r.JSON.raw }
func (r *XAnalysisNewsResponseAnalysisSummaryRows) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type XAnalysisNewsResponseAnalysisSummarySourceDomain struct {
	Domain string `json:"domain" api:"required"`
	Posts  int64  `json:"posts" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Domain      respjson.Field
		Posts       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XAnalysisNewsResponseAnalysisSummarySourceDomain) RawJSON() string { return r.JSON.raw }
func (r *XAnalysisNewsResponseAnalysisSummarySourceDomain) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type XAnalysisNewsResponseAnalysisSummaryTarget struct {
	// Category counts by question ID.
	Choices    PublicMetadata `json:"choices" api:"required"`
	Engagement int64          `json:"engagement" api:"required"`
	Mentions   int64          `json:"mentions" api:"required"`
	Name       string         `json:"name" api:"required"`
	Share      float64        `json:"share" api:"required"`
	// The most engaged posts per question ID & category.
	Top PublicMetadata `json:"top" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Choices     respjson.Field
		Engagement  respjson.Field
		Mentions    respjson.Field
		Name        respjson.Field
		Share       respjson.Field
		Top         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XAnalysisNewsResponseAnalysisSummaryTarget) RawJSON() string { return r.JSON.raw }
func (r *XAnalysisNewsResponseAnalysisSummaryTarget) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// 1 analyzed post with its answers.
type XAnalysisNewsResponseResult struct {
	// The answers in full, with their probabilities.
	Analysis XAnalysisNewsResponseResultAnalysis `json:"analysis" api:"required"`
	// Each question's decision by question ID. A choice gives its category, a score
	// its level from 0, which can have decimals, & a probability its value from 0 to
	// 1.
	Answers map[string]XAnalysisNewsResponseResultAnswersUnion `json:"answers" api:"required"`
	// Cashtags in the post's text, in upper case.
	Cashtags []string `json:"cashtags" api:"required"`
	// How the post's answers compare with its row in baseline, as the AI Actors
	// compare a run with an earlier dataset.
	Monitor XAnalysisNewsResponseResultMonitor `json:"monitor" api:"required"`
	// Hostnames the post links to, without t.co.
	SourceDomains []string `json:"sourceDomains" api:"required"`
	// The post, as /x/tweets/search returns it. Your own texts have the IDs text:1,
	// text:2 & so on.
	Tweet shared.SearchTweet `json:"tweet" api:"required"`
	// Only from /x/analysis/viral-score. The score with its weights version, the caps
	// that applied & the engagement comparison.
	Viral PublicMetadata `json:"viral"`
	// Only from /x/analysis/viral-score. log10 of 1 plus the weighted engagement per
	// 1,000 followers. Null without a follower count.
	ViralActualEngagementRate float64 `json:"viralActualEngagementRate" api:"nullable"`
	// Only from /x/analysis/viral-score. The post's public counts under X's published
	// ranking weights, per 1,000 views or followers.
	ViralAlgorithmScore float64 `json:"viralAlgorithmScore" api:"nullable"`
	// Only from /x/analysis/viral-score. How viral the wording reads, or null when a
	// trait has no answer.
	ViralScore int64 `json:"viralScore" api:"nullable"`
	// Only from /x/analysis/viral-score. send_it from 70, edit_first from 40, else
	// sleep_on_it.
	//
	// Any of "send_it", "edit_first", "sleep_on_it".
	ViralVerdict XAnalysisNewsResponseResultViralVerdict `json:"viralVerdict" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Analysis                  respjson.Field
		Answers                   respjson.Field
		Cashtags                  respjson.Field
		Monitor                   respjson.Field
		SourceDomains             respjson.Field
		Tweet                     respjson.Field
		Viral                     respjson.Field
		ViralActualEngagementRate respjson.Field
		ViralAlgorithmScore       respjson.Field
		ViralScore                respjson.Field
		ViralVerdict              respjson.Field
		ExtraFields               map[string]respjson.Field
		raw                       string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XAnalysisNewsResponseResult) RawJSON() string { return r.JSON.raw }
func (r *XAnalysisNewsResponseResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// The answers in full, with their probabilities.
type XAnalysisNewsResponseResultAnalysis struct {
	// 1 answer per question, in question order.
	Answers []XAnalysisNewsResponseResultAnalysisAnswer `json:"answers" api:"required"`
	// What the AI read beside the post's text.
	ContextAvailability XAnalysisNewsResponseResultAnalysisContextAvailability `json:"contextAvailability" api:"required"`
	// Bytes of post & settings the AI read.
	ContextBytes int64 `json:"contextBytes" api:"required"`
	// The ID & version of each question asked.
	Questions []XAnalysisNewsResponseResultAnalysisQuestion `json:"questions" api:"required"`
	Status    constant.Succeeded                            `json:"status" default:"succeeded"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Answers             respjson.Field
		ContextAvailability respjson.Field
		ContextBytes        respjson.Field
		Questions           respjson.Field
		Status              respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XAnalysisNewsResponseResultAnalysis) RawJSON() string { return r.JSON.raw }
func (r *XAnalysisNewsResponseResultAnalysis) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// 1 question's answer. A probability has probability. A choice or score has value,
// confidence & the probability of each category or level.
type XAnalysisNewsResponseResultAnalysisAnswer struct {
	QuestionID      string `json:"questionId" api:"required"`
	QuestionVersion string `json:"questionVersion" api:"required"`
	// Any of "choice", "score", "probability".
	Type       XAnalysisNewsResponseResultAnalysisAnswerType `json:"type" api:"required"`
	Confidence float64                                       `json:"confidence"`
	// For a score. The levels the question named, from 0.
	Levels []string `json:"levels"`
	// Probability of each category, or of each level by index.
	Probabilities map[string]float64 `json:"probabilities"`
	// For a probability. The chance the answer is yes.
	Probability float64 `json:"probability"`
	// The chosen category, or the level from 0, which can have decimals, such as 1.4
	// between levels 1 & 2.
	Value XAnalysisNewsResponseResultAnalysisAnswerValueUnion `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		QuestionID      respjson.Field
		QuestionVersion respjson.Field
		Type            respjson.Field
		Confidence      respjson.Field
		Levels          respjson.Field
		Probabilities   respjson.Field
		Probability     respjson.Field
		Value           respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XAnalysisNewsResponseResultAnalysisAnswer) RawJSON() string { return r.JSON.raw }
func (r *XAnalysisNewsResponseResultAnalysisAnswer) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type XAnalysisNewsResponseResultAnalysisAnswerType string

const (
	XAnalysisNewsResponseResultAnalysisAnswerTypeChoice      XAnalysisNewsResponseResultAnalysisAnswerType = "choice"
	XAnalysisNewsResponseResultAnalysisAnswerTypeScore       XAnalysisNewsResponseResultAnalysisAnswerType = "score"
	XAnalysisNewsResponseResultAnalysisAnswerTypeProbability XAnalysisNewsResponseResultAnalysisAnswerType = "probability"
)

// XAnalysisNewsResponseResultAnalysisAnswerValueUnion contains all possible
// properties and values from [float64], [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfFloat OfString]
type XAnalysisNewsResponseResultAnalysisAnswerValueUnion struct {
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	JSON     struct {
		OfFloat  respjson.Field
		OfString respjson.Field
		raw      string
	} `json:"-"`
}

func (u XAnalysisNewsResponseResultAnalysisAnswerValueUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v, nil)
	return
}

func (u XAnalysisNewsResponseResultAnalysisAnswerValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v, nil)
	return
}

// Returns the unmodified JSON received from the API
func (u XAnalysisNewsResponseResultAnalysisAnswerValueUnion) RawJSON() string { return u.JSON.raw }

func (r *XAnalysisNewsResponseResultAnalysisAnswerValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// What the AI read beside the post's text.
type XAnalysisNewsResponseResultAnalysisContextAvailability struct {
	// Any of "text_blocks", "summary", "not_supplied".
	Article XAnalysisNewsResponseResultAnalysisContextAvailabilityArticle `json:"article" api:"required"`
	// Any of "available", "missing".
	Author XAnalysisNewsResponseResultAnalysisContextAvailabilityAuthor `json:"author" api:"required"`
	// Any of "text_descriptions_only", "not_supplied".
	Media XAnalysisNewsResponseResultAnalysisContextAvailabilityMedia `json:"media" api:"required"`
	// truncated when the post was cut to fit.
	//
	// Any of "full", "truncated".
	PostText XAnalysisNewsResponseResultAnalysisContextAvailabilityPostText `json:"postText" api:"required"`
	// Any of "available", "missing", "not_applicable".
	Quote XAnalysisNewsResponseResultAnalysisContextAvailabilityQuote `json:"quote" api:"required"`
	// Any of "available", "missing", "not_applicable".
	Reply XAnalysisNewsResponseResultAnalysisContextAvailabilityReply `json:"reply" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Article     respjson.Field
		Author      respjson.Field
		Media       respjson.Field
		PostText    respjson.Field
		Quote       respjson.Field
		Reply       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XAnalysisNewsResponseResultAnalysisContextAvailability) RawJSON() string { return r.JSON.raw }
func (r *XAnalysisNewsResponseResultAnalysisContextAvailability) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type XAnalysisNewsResponseResultAnalysisContextAvailabilityArticle string

const (
	XAnalysisNewsResponseResultAnalysisContextAvailabilityArticleTextBlocks  XAnalysisNewsResponseResultAnalysisContextAvailabilityArticle = "text_blocks"
	XAnalysisNewsResponseResultAnalysisContextAvailabilityArticleSummary     XAnalysisNewsResponseResultAnalysisContextAvailabilityArticle = "summary"
	XAnalysisNewsResponseResultAnalysisContextAvailabilityArticleNotSupplied XAnalysisNewsResponseResultAnalysisContextAvailabilityArticle = "not_supplied"
)

type XAnalysisNewsResponseResultAnalysisContextAvailabilityAuthor string

const (
	XAnalysisNewsResponseResultAnalysisContextAvailabilityAuthorAvailable XAnalysisNewsResponseResultAnalysisContextAvailabilityAuthor = "available"
	XAnalysisNewsResponseResultAnalysisContextAvailabilityAuthorMissing   XAnalysisNewsResponseResultAnalysisContextAvailabilityAuthor = "missing"
)

type XAnalysisNewsResponseResultAnalysisContextAvailabilityMedia string

const (
	XAnalysisNewsResponseResultAnalysisContextAvailabilityMediaTextDescriptionsOnly XAnalysisNewsResponseResultAnalysisContextAvailabilityMedia = "text_descriptions_only"
	XAnalysisNewsResponseResultAnalysisContextAvailabilityMediaNotSupplied          XAnalysisNewsResponseResultAnalysisContextAvailabilityMedia = "not_supplied"
)

// truncated when the post was cut to fit.
type XAnalysisNewsResponseResultAnalysisContextAvailabilityPostText string

const (
	XAnalysisNewsResponseResultAnalysisContextAvailabilityPostTextFull      XAnalysisNewsResponseResultAnalysisContextAvailabilityPostText = "full"
	XAnalysisNewsResponseResultAnalysisContextAvailabilityPostTextTruncated XAnalysisNewsResponseResultAnalysisContextAvailabilityPostText = "truncated"
)

type XAnalysisNewsResponseResultAnalysisContextAvailabilityQuote string

const (
	XAnalysisNewsResponseResultAnalysisContextAvailabilityQuoteAvailable     XAnalysisNewsResponseResultAnalysisContextAvailabilityQuote = "available"
	XAnalysisNewsResponseResultAnalysisContextAvailabilityQuoteMissing       XAnalysisNewsResponseResultAnalysisContextAvailabilityQuote = "missing"
	XAnalysisNewsResponseResultAnalysisContextAvailabilityQuoteNotApplicable XAnalysisNewsResponseResultAnalysisContextAvailabilityQuote = "not_applicable"
)

type XAnalysisNewsResponseResultAnalysisContextAvailabilityReply string

const (
	XAnalysisNewsResponseResultAnalysisContextAvailabilityReplyAvailable     XAnalysisNewsResponseResultAnalysisContextAvailabilityReply = "available"
	XAnalysisNewsResponseResultAnalysisContextAvailabilityReplyMissing       XAnalysisNewsResponseResultAnalysisContextAvailabilityReply = "missing"
	XAnalysisNewsResponseResultAnalysisContextAvailabilityReplyNotApplicable XAnalysisNewsResponseResultAnalysisContextAvailabilityReply = "not_applicable"
)

type XAnalysisNewsResponseResultAnalysisQuestion struct {
	ID      string `json:"id" api:"required"`
	Version string `json:"version" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Version     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XAnalysisNewsResponseResultAnalysisQuestion) RawJSON() string { return r.JSON.raw }
func (r *XAnalysisNewsResponseResultAnalysisQuestion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// XAnalysisNewsResponseResultAnswersUnion contains all possible properties and
// values from [float64], [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfFloat OfString]
type XAnalysisNewsResponseResultAnswersUnion struct {
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	JSON     struct {
		OfFloat  respjson.Field
		OfString respjson.Field
		raw      string
	} `json:"-"`
}

func (u XAnalysisNewsResponseResultAnswersUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v, nil)
	return
}

func (u XAnalysisNewsResponseResultAnswersUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v, nil)
	return
}

// Returns the unmodified JSON received from the API
func (u XAnalysisNewsResponseResultAnswersUnion) RawJSON() string { return u.JSON.raw }

func (r *XAnalysisNewsResponseResultAnswersUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// How the post's answers compare with its row in baseline, as the AI Actors
// compare a run with an earlier dataset.
type XAnalysisNewsResponseResultMonitor struct {
	// Always null. The baseline is the rows you sent.
	BaselineDatasetID json.RawMessage `json:"baselineDatasetId" api:"required"`
	// IDs of the questions whose decision changed.
	ChangedQuestionIDs []string `json:"changedQuestionIds" api:"required"`
	// Each changed decision, earlier & now.
	Changes []XAnalysisNewsResponseResultMonitorChange `json:"changes" api:"required"`
	// Hash of the questions, targets & context. Rows with another hash are
	// not_comparable.
	ConfigurationFingerprint string `json:"configurationFingerprint" api:"required"`
	// first_run without baseline. new_to_baseline for a post baseline lacks. changed
	// when an answer clearly left its earlier decision. not_comparable when other
	// settings made the earlier row or it answers none of these questions.
	//
	// Any of "first_run", "new_to_baseline", "unchanged", "changed", "not_comparable",
	// "analysis_unavailable".
	Status XAnalysisNewsResponseResultMonitorStatus `json:"status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BaselineDatasetID        respjson.Field
		ChangedQuestionIDs       respjson.Field
		Changes                  respjson.Field
		ConfigurationFingerprint respjson.Field
		Status                   respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XAnalysisNewsResponseResultMonitor) RawJSON() string { return r.JSON.raw }
func (r *XAnalysisNewsResponseResultMonitor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type XAnalysisNewsResponseResultMonitorChange struct {
	Current    string `json:"current" api:"required"`
	Previous   string `json:"previous" api:"required"`
	QuestionID string `json:"questionId" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Current     respjson.Field
		Previous    respjson.Field
		QuestionID  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XAnalysisNewsResponseResultMonitorChange) RawJSON() string { return r.JSON.raw }
func (r *XAnalysisNewsResponseResultMonitorChange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// first_run without baseline. new_to_baseline for a post baseline lacks. changed
// when an answer clearly left its earlier decision. not_comparable when other
// settings made the earlier row or it answers none of these questions.
type XAnalysisNewsResponseResultMonitorStatus string

const (
	XAnalysisNewsResponseResultMonitorStatusFirstRun            XAnalysisNewsResponseResultMonitorStatus = "first_run"
	XAnalysisNewsResponseResultMonitorStatusNewToBaseline       XAnalysisNewsResponseResultMonitorStatus = "new_to_baseline"
	XAnalysisNewsResponseResultMonitorStatusUnchanged           XAnalysisNewsResponseResultMonitorStatus = "unchanged"
	XAnalysisNewsResponseResultMonitorStatusChanged             XAnalysisNewsResponseResultMonitorStatus = "changed"
	XAnalysisNewsResponseResultMonitorStatusNotComparable       XAnalysisNewsResponseResultMonitorStatus = "not_comparable"
	XAnalysisNewsResponseResultMonitorStatusAnalysisUnavailable XAnalysisNewsResponseResultMonitorStatus = "analysis_unavailable"
)

// Only from /x/analysis/viral-score. send_it from 70, edit_first from 40, else
// sleep_on_it.
type XAnalysisNewsResponseResultViralVerdict string

const (
	XAnalysisNewsResponseResultViralVerdictSendIt    XAnalysisNewsResponseResultViralVerdict = "send_it"
	XAnalysisNewsResponseResultViralVerdictEditFirst XAnalysisNewsResponseResultViralVerdict = "edit_first"
	XAnalysisNewsResponseResultViralVerdictSleepOnIt XAnalysisNewsResponseResultViralVerdict = "sleep_on_it"
)

// A post without an analysis. It costs nothing.
type XAnalysisNewsResponseUnanalyzed struct {
	// Post ID. Your own texts have the IDs text:1, text:2 & so on.
	ID string `json:"id" api:"required"`
	// post_unavailable when X returns no such post. missing_text for a post without
	// text. insufficient_credits when credits ran out. context_limit when the settings
	// leave the post no room. The other reasons mean the analysis failed. Send the
	// post again.
	Reason string `json:"reason" api:"required"`
	// Any of "failed", "skipped".
	Status XAnalysisNewsResponseUnanalyzedStatus `json:"status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Reason      respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XAnalysisNewsResponseUnanalyzed) RawJSON() string { return r.JSON.raw }
func (r *XAnalysisNewsResponseUnanalyzed) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type XAnalysisNewsResponseUnanalyzedStatus string

const (
	XAnalysisNewsResponseUnanalyzedStatusFailed  XAnalysisNewsResponseUnanalyzedStatus = "failed"
	XAnalysisNewsResponseUnanalyzedStatusSkipped XAnalysisNewsResponseUnanalyzedStatus = "skipped"
)

// Analyzed posts, the posts without an analysis & the totals of the answers. Each
// row in results costs 2 credits.
type XAnalysisSentimentResponse struct {
	// Totals of the answers across the analyzed posts.
	AnalysisSummary XAnalysisSentimentResponseAnalysisSummary `json:"analysisSummary" api:"required"`
	// Whether the search has more posts. Send next_cursor as cursor.
	HasNextPage bool `json:"has_next_page" api:"required"`
	// Cursor of the search's next page, else empty.
	NextCursor string `json:"next_cursor" api:"required"`
	// Each analyzed post with its answers.
	Results []XAnalysisSentimentResponseResult `json:"results" api:"required"`
	// Each post without an analysis, by ID & reason. These cost nothing.
	Unanalyzed []XAnalysisSentimentResponseUnanalyzed `json:"unanalyzed" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AnalysisSummary respjson.Field
		HasNextPage     respjson.Field
		NextCursor      respjson.Field
		Results         respjson.Field
		Unanalyzed      respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XAnalysisSentimentResponse) RawJSON() string { return r.JSON.raw }
func (r *XAnalysisSentimentResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// Totals of the answers across the analyzed posts.
type XAnalysisSentimentResponseAnalysisSummary struct {
	// Posts & answers per cashtag, most posts first, up to 20. signal is present for
	// the market questions.
	Cashtags []XAnalysisSentimentResponseAnalysisSummaryCashtag `json:"cashtags" api:"required"`
	// Likes, reposts, replies & quotes of the analyzed posts.
	Engagement int64 `json:"engagement" api:"required"`
	// Posts per comparison status, & up to 50 changed posts with their changes. Null
	// when the answer holds no post.
	Monitor XAnalysisSentimentResponseAnalysisSummaryMonitor `json:"monitor" api:"required"`
	// Totals by question ID. A choice has counts, shares, engagementShares & top posts
	// per category. A score has mean, engagementWeightedMean & levels. A probability
	// has mean, yes & no.
	Questions map[string]PublicMetadata `json:"questions" api:"required"`
	// Posts by analysis outcome.
	Rows          XAnalysisSentimentResponseAnalysisSummaryRows `json:"rows" api:"required"`
	SchemaVersion int64                                         `json:"schemaVersion" api:"required"`
	// Posts per linked hostname, most posts first, up to 20.
	SourceDomains []XAnalysisSentimentResponseAnalysisSummarySourceDomain `json:"sourceDomains" api:"required"`
	// Mentions & answers per target in analysis.targets.
	Targets []XAnalysisSentimentResponseAnalysisSummaryTarget `json:"targets" api:"required"`
	// Only from /x/analysis/viral-score. Scored posts, the average score, verdict
	// counts, the scores against actual engagement, a report per author & the
	// leaderboards.
	Viral PublicMetadata `json:"viral"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Cashtags      respjson.Field
		Engagement    respjson.Field
		Monitor       respjson.Field
		Questions     respjson.Field
		Rows          respjson.Field
		SchemaVersion respjson.Field
		SourceDomains respjson.Field
		Targets       respjson.Field
		Viral         respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XAnalysisSentimentResponseAnalysisSummary) RawJSON() string { return r.JSON.raw }
func (r *XAnalysisSentimentResponseAnalysisSummary) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type XAnalysisSentimentResponseAnalysisSummaryCashtag struct {
	Cashtag string `json:"cashtag" api:"required"`
	// Category counts of the cashtag's posts by question ID.
	Choices PublicMetadata                                         `json:"choices" api:"required"`
	Rows    int64                                                  `json:"rows" api:"required"`
	Signal  XAnalysisSentimentResponseAnalysisSummaryCashtagSignal `json:"signal"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Cashtag     respjson.Field
		Choices     respjson.Field
		Rows        respjson.Field
		Signal      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XAnalysisSentimentResponseAnalysisSummaryCashtag) RawJSON() string { return r.JSON.raw }
func (r *XAnalysisSentimentResponseAnalysisSummaryCashtag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type XAnalysisSentimentResponseAnalysisSummaryCashtagSignal struct {
	Bearish int64 `json:"bearish" api:"required"`
	Bullish int64 `json:"bullish" api:"required"`
	// Bullish minus bearish posts, over all posts.
	Score float64 `json:"score" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Bearish     respjson.Field
		Bullish     respjson.Field
		Score       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XAnalysisSentimentResponseAnalysisSummaryCashtagSignal) RawJSON() string { return r.JSON.raw }
func (r *XAnalysisSentimentResponseAnalysisSummaryCashtagSignal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// Posts per comparison status, & up to 50 changed posts with their changes. Null
// when the answer holds no post.
type XAnalysisSentimentResponseAnalysisSummaryMonitor struct {
	ChangedRows []XAnalysisSentimentResponseAnalysisSummaryMonitorChangedRow `json:"changedRows" api:"required"`
	// Posts per monitor status.
	Statuses map[string]int64 `json:"statuses" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChangedRows respjson.Field
		Statuses    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XAnalysisSentimentResponseAnalysisSummaryMonitor) RawJSON() string { return r.JSON.raw }
func (r *XAnalysisSentimentResponseAnalysisSummaryMonitor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type XAnalysisSentimentResponseAnalysisSummaryMonitorChangedRow struct {
	Changes []XAnalysisSentimentResponseAnalysisSummaryMonitorChangedRowChange `json:"changes" api:"required"`
	TweetID string                                                             `json:"tweetId" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Changes     respjson.Field
		TweetID     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XAnalysisSentimentResponseAnalysisSummaryMonitorChangedRow) RawJSON() string {
	return r.JSON.raw
}
func (r *XAnalysisSentimentResponseAnalysisSummaryMonitorChangedRow) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type XAnalysisSentimentResponseAnalysisSummaryMonitorChangedRowChange struct {
	Current    string `json:"current" api:"required"`
	Previous   string `json:"previous" api:"required"`
	QuestionID string `json:"questionId" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Current     respjson.Field
		Previous    respjson.Field
		QuestionID  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XAnalysisSentimentResponseAnalysisSummaryMonitorChangedRowChange) RawJSON() string {
	return r.JSON.raw
}
func (r *XAnalysisSentimentResponseAnalysisSummaryMonitorChangedRowChange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// Posts by analysis outcome.
type XAnalysisSentimentResponseAnalysisSummaryRows struct {
	Analyzed int64 `json:"analyzed" api:"required"`
	Failed   int64 `json:"failed" api:"required"`
	Skipped  int64 `json:"skipped" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Analyzed    respjson.Field
		Failed      respjson.Field
		Skipped     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XAnalysisSentimentResponseAnalysisSummaryRows) RawJSON() string { return r.JSON.raw }
func (r *XAnalysisSentimentResponseAnalysisSummaryRows) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type XAnalysisSentimentResponseAnalysisSummarySourceDomain struct {
	Domain string `json:"domain" api:"required"`
	Posts  int64  `json:"posts" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Domain      respjson.Field
		Posts       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XAnalysisSentimentResponseAnalysisSummarySourceDomain) RawJSON() string { return r.JSON.raw }
func (r *XAnalysisSentimentResponseAnalysisSummarySourceDomain) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type XAnalysisSentimentResponseAnalysisSummaryTarget struct {
	// Category counts by question ID.
	Choices    PublicMetadata `json:"choices" api:"required"`
	Engagement int64          `json:"engagement" api:"required"`
	Mentions   int64          `json:"mentions" api:"required"`
	Name       string         `json:"name" api:"required"`
	Share      float64        `json:"share" api:"required"`
	// The most engaged posts per question ID & category.
	Top PublicMetadata `json:"top" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Choices     respjson.Field
		Engagement  respjson.Field
		Mentions    respjson.Field
		Name        respjson.Field
		Share       respjson.Field
		Top         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XAnalysisSentimentResponseAnalysisSummaryTarget) RawJSON() string { return r.JSON.raw }
func (r *XAnalysisSentimentResponseAnalysisSummaryTarget) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// 1 analyzed post with its answers.
type XAnalysisSentimentResponseResult struct {
	// The answers in full, with their probabilities.
	Analysis XAnalysisSentimentResponseResultAnalysis `json:"analysis" api:"required"`
	// Each question's decision by question ID. A choice gives its category, a score
	// its level from 0, which can have decimals, & a probability its value from 0 to
	// 1.
	Answers map[string]XAnalysisSentimentResponseResultAnswersUnion `json:"answers" api:"required"`
	// Cashtags in the post's text, in upper case.
	Cashtags []string `json:"cashtags" api:"required"`
	// How the post's answers compare with its row in baseline, as the AI Actors
	// compare a run with an earlier dataset.
	Monitor XAnalysisSentimentResponseResultMonitor `json:"monitor" api:"required"`
	// Hostnames the post links to, without t.co.
	SourceDomains []string `json:"sourceDomains" api:"required"`
	// The post, as /x/tweets/search returns it. Your own texts have the IDs text:1,
	// text:2 & so on.
	Tweet shared.SearchTweet `json:"tweet" api:"required"`
	// Only from /x/analysis/viral-score. The score with its weights version, the caps
	// that applied & the engagement comparison.
	Viral PublicMetadata `json:"viral"`
	// Only from /x/analysis/viral-score. log10 of 1 plus the weighted engagement per
	// 1,000 followers. Null without a follower count.
	ViralActualEngagementRate float64 `json:"viralActualEngagementRate" api:"nullable"`
	// Only from /x/analysis/viral-score. The post's public counts under X's published
	// ranking weights, per 1,000 views or followers.
	ViralAlgorithmScore float64 `json:"viralAlgorithmScore" api:"nullable"`
	// Only from /x/analysis/viral-score. How viral the wording reads, or null when a
	// trait has no answer.
	ViralScore int64 `json:"viralScore" api:"nullable"`
	// Only from /x/analysis/viral-score. send_it from 70, edit_first from 40, else
	// sleep_on_it.
	//
	// Any of "send_it", "edit_first", "sleep_on_it".
	ViralVerdict XAnalysisSentimentResponseResultViralVerdict `json:"viralVerdict" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Analysis                  respjson.Field
		Answers                   respjson.Field
		Cashtags                  respjson.Field
		Monitor                   respjson.Field
		SourceDomains             respjson.Field
		Tweet                     respjson.Field
		Viral                     respjson.Field
		ViralActualEngagementRate respjson.Field
		ViralAlgorithmScore       respjson.Field
		ViralScore                respjson.Field
		ViralVerdict              respjson.Field
		ExtraFields               map[string]respjson.Field
		raw                       string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XAnalysisSentimentResponseResult) RawJSON() string { return r.JSON.raw }
func (r *XAnalysisSentimentResponseResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// The answers in full, with their probabilities.
type XAnalysisSentimentResponseResultAnalysis struct {
	// 1 answer per question, in question order.
	Answers []XAnalysisSentimentResponseResultAnalysisAnswer `json:"answers" api:"required"`
	// What the AI read beside the post's text.
	ContextAvailability XAnalysisSentimentResponseResultAnalysisContextAvailability `json:"contextAvailability" api:"required"`
	// Bytes of post & settings the AI read.
	ContextBytes int64 `json:"contextBytes" api:"required"`
	// The ID & version of each question asked.
	Questions []XAnalysisSentimentResponseResultAnalysisQuestion `json:"questions" api:"required"`
	Status    constant.Succeeded                                 `json:"status" default:"succeeded"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Answers             respjson.Field
		ContextAvailability respjson.Field
		ContextBytes        respjson.Field
		Questions           respjson.Field
		Status              respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XAnalysisSentimentResponseResultAnalysis) RawJSON() string { return r.JSON.raw }
func (r *XAnalysisSentimentResponseResultAnalysis) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// 1 question's answer. A probability has probability. A choice or score has value,
// confidence & the probability of each category or level.
type XAnalysisSentimentResponseResultAnalysisAnswer struct {
	QuestionID      string `json:"questionId" api:"required"`
	QuestionVersion string `json:"questionVersion" api:"required"`
	// Any of "choice", "score", "probability".
	Type       XAnalysisSentimentResponseResultAnalysisAnswerType `json:"type" api:"required"`
	Confidence float64                                            `json:"confidence"`
	// For a score. The levels the question named, from 0.
	Levels []string `json:"levels"`
	// Probability of each category, or of each level by index.
	Probabilities map[string]float64 `json:"probabilities"`
	// For a probability. The chance the answer is yes.
	Probability float64 `json:"probability"`
	// The chosen category, or the level from 0, which can have decimals, such as 1.4
	// between levels 1 & 2.
	Value XAnalysisSentimentResponseResultAnalysisAnswerValueUnion `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		QuestionID      respjson.Field
		QuestionVersion respjson.Field
		Type            respjson.Field
		Confidence      respjson.Field
		Levels          respjson.Field
		Probabilities   respjson.Field
		Probability     respjson.Field
		Value           respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XAnalysisSentimentResponseResultAnalysisAnswer) RawJSON() string { return r.JSON.raw }
func (r *XAnalysisSentimentResponseResultAnalysisAnswer) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type XAnalysisSentimentResponseResultAnalysisAnswerType string

const (
	XAnalysisSentimentResponseResultAnalysisAnswerTypeChoice      XAnalysisSentimentResponseResultAnalysisAnswerType = "choice"
	XAnalysisSentimentResponseResultAnalysisAnswerTypeScore       XAnalysisSentimentResponseResultAnalysisAnswerType = "score"
	XAnalysisSentimentResponseResultAnalysisAnswerTypeProbability XAnalysisSentimentResponseResultAnalysisAnswerType = "probability"
)

// XAnalysisSentimentResponseResultAnalysisAnswerValueUnion contains all possible
// properties and values from [float64], [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfFloat OfString]
type XAnalysisSentimentResponseResultAnalysisAnswerValueUnion struct {
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	JSON     struct {
		OfFloat  respjson.Field
		OfString respjson.Field
		raw      string
	} `json:"-"`
}

func (u XAnalysisSentimentResponseResultAnalysisAnswerValueUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v, nil)
	return
}

func (u XAnalysisSentimentResponseResultAnalysisAnswerValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v, nil)
	return
}

// Returns the unmodified JSON received from the API
func (u XAnalysisSentimentResponseResultAnalysisAnswerValueUnion) RawJSON() string { return u.JSON.raw }

func (r *XAnalysisSentimentResponseResultAnalysisAnswerValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// What the AI read beside the post's text.
type XAnalysisSentimentResponseResultAnalysisContextAvailability struct {
	// Any of "text_blocks", "summary", "not_supplied".
	Article XAnalysisSentimentResponseResultAnalysisContextAvailabilityArticle `json:"article" api:"required"`
	// Any of "available", "missing".
	Author XAnalysisSentimentResponseResultAnalysisContextAvailabilityAuthor `json:"author" api:"required"`
	// Any of "text_descriptions_only", "not_supplied".
	Media XAnalysisSentimentResponseResultAnalysisContextAvailabilityMedia `json:"media" api:"required"`
	// truncated when the post was cut to fit.
	//
	// Any of "full", "truncated".
	PostText XAnalysisSentimentResponseResultAnalysisContextAvailabilityPostText `json:"postText" api:"required"`
	// Any of "available", "missing", "not_applicable".
	Quote XAnalysisSentimentResponseResultAnalysisContextAvailabilityQuote `json:"quote" api:"required"`
	// Any of "available", "missing", "not_applicable".
	Reply XAnalysisSentimentResponseResultAnalysisContextAvailabilityReply `json:"reply" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Article     respjson.Field
		Author      respjson.Field
		Media       respjson.Field
		PostText    respjson.Field
		Quote       respjson.Field
		Reply       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XAnalysisSentimentResponseResultAnalysisContextAvailability) RawJSON() string {
	return r.JSON.raw
}
func (r *XAnalysisSentimentResponseResultAnalysisContextAvailability) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type XAnalysisSentimentResponseResultAnalysisContextAvailabilityArticle string

const (
	XAnalysisSentimentResponseResultAnalysisContextAvailabilityArticleTextBlocks  XAnalysisSentimentResponseResultAnalysisContextAvailabilityArticle = "text_blocks"
	XAnalysisSentimentResponseResultAnalysisContextAvailabilityArticleSummary     XAnalysisSentimentResponseResultAnalysisContextAvailabilityArticle = "summary"
	XAnalysisSentimentResponseResultAnalysisContextAvailabilityArticleNotSupplied XAnalysisSentimentResponseResultAnalysisContextAvailabilityArticle = "not_supplied"
)

type XAnalysisSentimentResponseResultAnalysisContextAvailabilityAuthor string

const (
	XAnalysisSentimentResponseResultAnalysisContextAvailabilityAuthorAvailable XAnalysisSentimentResponseResultAnalysisContextAvailabilityAuthor = "available"
	XAnalysisSentimentResponseResultAnalysisContextAvailabilityAuthorMissing   XAnalysisSentimentResponseResultAnalysisContextAvailabilityAuthor = "missing"
)

type XAnalysisSentimentResponseResultAnalysisContextAvailabilityMedia string

const (
	XAnalysisSentimentResponseResultAnalysisContextAvailabilityMediaTextDescriptionsOnly XAnalysisSentimentResponseResultAnalysisContextAvailabilityMedia = "text_descriptions_only"
	XAnalysisSentimentResponseResultAnalysisContextAvailabilityMediaNotSupplied          XAnalysisSentimentResponseResultAnalysisContextAvailabilityMedia = "not_supplied"
)

// truncated when the post was cut to fit.
type XAnalysisSentimentResponseResultAnalysisContextAvailabilityPostText string

const (
	XAnalysisSentimentResponseResultAnalysisContextAvailabilityPostTextFull      XAnalysisSentimentResponseResultAnalysisContextAvailabilityPostText = "full"
	XAnalysisSentimentResponseResultAnalysisContextAvailabilityPostTextTruncated XAnalysisSentimentResponseResultAnalysisContextAvailabilityPostText = "truncated"
)

type XAnalysisSentimentResponseResultAnalysisContextAvailabilityQuote string

const (
	XAnalysisSentimentResponseResultAnalysisContextAvailabilityQuoteAvailable     XAnalysisSentimentResponseResultAnalysisContextAvailabilityQuote = "available"
	XAnalysisSentimentResponseResultAnalysisContextAvailabilityQuoteMissing       XAnalysisSentimentResponseResultAnalysisContextAvailabilityQuote = "missing"
	XAnalysisSentimentResponseResultAnalysisContextAvailabilityQuoteNotApplicable XAnalysisSentimentResponseResultAnalysisContextAvailabilityQuote = "not_applicable"
)

type XAnalysisSentimentResponseResultAnalysisContextAvailabilityReply string

const (
	XAnalysisSentimentResponseResultAnalysisContextAvailabilityReplyAvailable     XAnalysisSentimentResponseResultAnalysisContextAvailabilityReply = "available"
	XAnalysisSentimentResponseResultAnalysisContextAvailabilityReplyMissing       XAnalysisSentimentResponseResultAnalysisContextAvailabilityReply = "missing"
	XAnalysisSentimentResponseResultAnalysisContextAvailabilityReplyNotApplicable XAnalysisSentimentResponseResultAnalysisContextAvailabilityReply = "not_applicable"
)

type XAnalysisSentimentResponseResultAnalysisQuestion struct {
	ID      string `json:"id" api:"required"`
	Version string `json:"version" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Version     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XAnalysisSentimentResponseResultAnalysisQuestion) RawJSON() string { return r.JSON.raw }
func (r *XAnalysisSentimentResponseResultAnalysisQuestion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// XAnalysisSentimentResponseResultAnswersUnion contains all possible properties
// and values from [float64], [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfFloat OfString]
type XAnalysisSentimentResponseResultAnswersUnion struct {
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	JSON     struct {
		OfFloat  respjson.Field
		OfString respjson.Field
		raw      string
	} `json:"-"`
}

func (u XAnalysisSentimentResponseResultAnswersUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v, nil)
	return
}

func (u XAnalysisSentimentResponseResultAnswersUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v, nil)
	return
}

// Returns the unmodified JSON received from the API
func (u XAnalysisSentimentResponseResultAnswersUnion) RawJSON() string { return u.JSON.raw }

func (r *XAnalysisSentimentResponseResultAnswersUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// How the post's answers compare with its row in baseline, as the AI Actors
// compare a run with an earlier dataset.
type XAnalysisSentimentResponseResultMonitor struct {
	// Always null. The baseline is the rows you sent.
	BaselineDatasetID json.RawMessage `json:"baselineDatasetId" api:"required"`
	// IDs of the questions whose decision changed.
	ChangedQuestionIDs []string `json:"changedQuestionIds" api:"required"`
	// Each changed decision, earlier & now.
	Changes []XAnalysisSentimentResponseResultMonitorChange `json:"changes" api:"required"`
	// Hash of the questions, targets & context. Rows with another hash are
	// not_comparable.
	ConfigurationFingerprint string `json:"configurationFingerprint" api:"required"`
	// first_run without baseline. new_to_baseline for a post baseline lacks. changed
	// when an answer clearly left its earlier decision. not_comparable when other
	// settings made the earlier row or it answers none of these questions.
	//
	// Any of "first_run", "new_to_baseline", "unchanged", "changed", "not_comparable",
	// "analysis_unavailable".
	Status XAnalysisSentimentResponseResultMonitorStatus `json:"status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BaselineDatasetID        respjson.Field
		ChangedQuestionIDs       respjson.Field
		Changes                  respjson.Field
		ConfigurationFingerprint respjson.Field
		Status                   respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XAnalysisSentimentResponseResultMonitor) RawJSON() string { return r.JSON.raw }
func (r *XAnalysisSentimentResponseResultMonitor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type XAnalysisSentimentResponseResultMonitorChange struct {
	Current    string `json:"current" api:"required"`
	Previous   string `json:"previous" api:"required"`
	QuestionID string `json:"questionId" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Current     respjson.Field
		Previous    respjson.Field
		QuestionID  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XAnalysisSentimentResponseResultMonitorChange) RawJSON() string { return r.JSON.raw }
func (r *XAnalysisSentimentResponseResultMonitorChange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// first_run without baseline. new_to_baseline for a post baseline lacks. changed
// when an answer clearly left its earlier decision. not_comparable when other
// settings made the earlier row or it answers none of these questions.
type XAnalysisSentimentResponseResultMonitorStatus string

const (
	XAnalysisSentimentResponseResultMonitorStatusFirstRun            XAnalysisSentimentResponseResultMonitorStatus = "first_run"
	XAnalysisSentimentResponseResultMonitorStatusNewToBaseline       XAnalysisSentimentResponseResultMonitorStatus = "new_to_baseline"
	XAnalysisSentimentResponseResultMonitorStatusUnchanged           XAnalysisSentimentResponseResultMonitorStatus = "unchanged"
	XAnalysisSentimentResponseResultMonitorStatusChanged             XAnalysisSentimentResponseResultMonitorStatus = "changed"
	XAnalysisSentimentResponseResultMonitorStatusNotComparable       XAnalysisSentimentResponseResultMonitorStatus = "not_comparable"
	XAnalysisSentimentResponseResultMonitorStatusAnalysisUnavailable XAnalysisSentimentResponseResultMonitorStatus = "analysis_unavailable"
)

// Only from /x/analysis/viral-score. send_it from 70, edit_first from 40, else
// sleep_on_it.
type XAnalysisSentimentResponseResultViralVerdict string

const (
	XAnalysisSentimentResponseResultViralVerdictSendIt    XAnalysisSentimentResponseResultViralVerdict = "send_it"
	XAnalysisSentimentResponseResultViralVerdictEditFirst XAnalysisSentimentResponseResultViralVerdict = "edit_first"
	XAnalysisSentimentResponseResultViralVerdictSleepOnIt XAnalysisSentimentResponseResultViralVerdict = "sleep_on_it"
)

// A post without an analysis. It costs nothing.
type XAnalysisSentimentResponseUnanalyzed struct {
	// Post ID. Your own texts have the IDs text:1, text:2 & so on.
	ID string `json:"id" api:"required"`
	// post_unavailable when X returns no such post. missing_text for a post without
	// text. insufficient_credits when credits ran out. context_limit when the settings
	// leave the post no room. The other reasons mean the analysis failed. Send the
	// post again.
	Reason string `json:"reason" api:"required"`
	// Any of "failed", "skipped".
	Status XAnalysisSentimentResponseUnanalyzedStatus `json:"status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Reason      respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XAnalysisSentimentResponseUnanalyzed) RawJSON() string { return r.JSON.raw }
func (r *XAnalysisSentimentResponseUnanalyzed) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type XAnalysisSentimentResponseUnanalyzedStatus string

const (
	XAnalysisSentimentResponseUnanalyzedStatusFailed  XAnalysisSentimentResponseUnanalyzedStatus = "failed"
	XAnalysisSentimentResponseUnanalyzedStatusSkipped XAnalysisSentimentResponseUnanalyzedStatus = "skipped"
)

// Analyzed posts, the posts without an analysis & the totals of the answers. Each
// row in results costs 2 credits.
type XAnalysisViralScoreResponse struct {
	// Totals of the answers across the analyzed posts.
	AnalysisSummary XAnalysisViralScoreResponseAnalysisSummary `json:"analysisSummary" api:"required"`
	// Whether the search has more posts. Send next_cursor as cursor.
	HasNextPage bool `json:"has_next_page" api:"required"`
	// Cursor of the search's next page, else empty.
	NextCursor string `json:"next_cursor" api:"required"`
	// Each analyzed post with its answers.
	Results []XAnalysisViralScoreResponseResult `json:"results" api:"required"`
	// Each post without an analysis, by ID & reason. These cost nothing.
	Unanalyzed []XAnalysisViralScoreResponseUnanalyzed `json:"unanalyzed" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AnalysisSummary respjson.Field
		HasNextPage     respjson.Field
		NextCursor      respjson.Field
		Results         respjson.Field
		Unanalyzed      respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XAnalysisViralScoreResponse) RawJSON() string { return r.JSON.raw }
func (r *XAnalysisViralScoreResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// Totals of the answers across the analyzed posts.
type XAnalysisViralScoreResponseAnalysisSummary struct {
	// Posts & answers per cashtag, most posts first, up to 20. signal is present for
	// the market questions.
	Cashtags []XAnalysisViralScoreResponseAnalysisSummaryCashtag `json:"cashtags" api:"required"`
	// Likes, reposts, replies & quotes of the analyzed posts.
	Engagement int64 `json:"engagement" api:"required"`
	// Posts per comparison status, & up to 50 changed posts with their changes. Null
	// when the answer holds no post.
	Monitor XAnalysisViralScoreResponseAnalysisSummaryMonitor `json:"monitor" api:"required"`
	// Totals by question ID. A choice has counts, shares, engagementShares & top posts
	// per category. A score has mean, engagementWeightedMean & levels. A probability
	// has mean, yes & no.
	Questions map[string]PublicMetadata `json:"questions" api:"required"`
	// Posts by analysis outcome.
	Rows          XAnalysisViralScoreResponseAnalysisSummaryRows `json:"rows" api:"required"`
	SchemaVersion int64                                          `json:"schemaVersion" api:"required"`
	// Posts per linked hostname, most posts first, up to 20.
	SourceDomains []XAnalysisViralScoreResponseAnalysisSummarySourceDomain `json:"sourceDomains" api:"required"`
	// Mentions & answers per target in analysis.targets.
	Targets []XAnalysisViralScoreResponseAnalysisSummaryTarget `json:"targets" api:"required"`
	// Only from /x/analysis/viral-score. Scored posts, the average score, verdict
	// counts, the scores against actual engagement, a report per author & the
	// leaderboards.
	Viral PublicMetadata `json:"viral"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Cashtags      respjson.Field
		Engagement    respjson.Field
		Monitor       respjson.Field
		Questions     respjson.Field
		Rows          respjson.Field
		SchemaVersion respjson.Field
		SourceDomains respjson.Field
		Targets       respjson.Field
		Viral         respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XAnalysisViralScoreResponseAnalysisSummary) RawJSON() string { return r.JSON.raw }
func (r *XAnalysisViralScoreResponseAnalysisSummary) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type XAnalysisViralScoreResponseAnalysisSummaryCashtag struct {
	Cashtag string `json:"cashtag" api:"required"`
	// Category counts of the cashtag's posts by question ID.
	Choices PublicMetadata                                          `json:"choices" api:"required"`
	Rows    int64                                                   `json:"rows" api:"required"`
	Signal  XAnalysisViralScoreResponseAnalysisSummaryCashtagSignal `json:"signal"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Cashtag     respjson.Field
		Choices     respjson.Field
		Rows        respjson.Field
		Signal      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XAnalysisViralScoreResponseAnalysisSummaryCashtag) RawJSON() string { return r.JSON.raw }
func (r *XAnalysisViralScoreResponseAnalysisSummaryCashtag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type XAnalysisViralScoreResponseAnalysisSummaryCashtagSignal struct {
	Bearish int64 `json:"bearish" api:"required"`
	Bullish int64 `json:"bullish" api:"required"`
	// Bullish minus bearish posts, over all posts.
	Score float64 `json:"score" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Bearish     respjson.Field
		Bullish     respjson.Field
		Score       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XAnalysisViralScoreResponseAnalysisSummaryCashtagSignal) RawJSON() string { return r.JSON.raw }
func (r *XAnalysisViralScoreResponseAnalysisSummaryCashtagSignal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// Posts per comparison status, & up to 50 changed posts with their changes. Null
// when the answer holds no post.
type XAnalysisViralScoreResponseAnalysisSummaryMonitor struct {
	ChangedRows []XAnalysisViralScoreResponseAnalysisSummaryMonitorChangedRow `json:"changedRows" api:"required"`
	// Posts per monitor status.
	Statuses map[string]int64 `json:"statuses" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChangedRows respjson.Field
		Statuses    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XAnalysisViralScoreResponseAnalysisSummaryMonitor) RawJSON() string { return r.JSON.raw }
func (r *XAnalysisViralScoreResponseAnalysisSummaryMonitor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type XAnalysisViralScoreResponseAnalysisSummaryMonitorChangedRow struct {
	Changes []XAnalysisViralScoreResponseAnalysisSummaryMonitorChangedRowChange `json:"changes" api:"required"`
	TweetID string                                                              `json:"tweetId" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Changes     respjson.Field
		TweetID     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XAnalysisViralScoreResponseAnalysisSummaryMonitorChangedRow) RawJSON() string {
	return r.JSON.raw
}
func (r *XAnalysisViralScoreResponseAnalysisSummaryMonitorChangedRow) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type XAnalysisViralScoreResponseAnalysisSummaryMonitorChangedRowChange struct {
	Current    string `json:"current" api:"required"`
	Previous   string `json:"previous" api:"required"`
	QuestionID string `json:"questionId" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Current     respjson.Field
		Previous    respjson.Field
		QuestionID  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XAnalysisViralScoreResponseAnalysisSummaryMonitorChangedRowChange) RawJSON() string {
	return r.JSON.raw
}
func (r *XAnalysisViralScoreResponseAnalysisSummaryMonitorChangedRowChange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// Posts by analysis outcome.
type XAnalysisViralScoreResponseAnalysisSummaryRows struct {
	Analyzed int64 `json:"analyzed" api:"required"`
	Failed   int64 `json:"failed" api:"required"`
	Skipped  int64 `json:"skipped" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Analyzed    respjson.Field
		Failed      respjson.Field
		Skipped     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XAnalysisViralScoreResponseAnalysisSummaryRows) RawJSON() string { return r.JSON.raw }
func (r *XAnalysisViralScoreResponseAnalysisSummaryRows) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type XAnalysisViralScoreResponseAnalysisSummarySourceDomain struct {
	Domain string `json:"domain" api:"required"`
	Posts  int64  `json:"posts" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Domain      respjson.Field
		Posts       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XAnalysisViralScoreResponseAnalysisSummarySourceDomain) RawJSON() string { return r.JSON.raw }
func (r *XAnalysisViralScoreResponseAnalysisSummarySourceDomain) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type XAnalysisViralScoreResponseAnalysisSummaryTarget struct {
	// Category counts by question ID.
	Choices    PublicMetadata `json:"choices" api:"required"`
	Engagement int64          `json:"engagement" api:"required"`
	Mentions   int64          `json:"mentions" api:"required"`
	Name       string         `json:"name" api:"required"`
	Share      float64        `json:"share" api:"required"`
	// The most engaged posts per question ID & category.
	Top PublicMetadata `json:"top" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Choices     respjson.Field
		Engagement  respjson.Field
		Mentions    respjson.Field
		Name        respjson.Field
		Share       respjson.Field
		Top         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XAnalysisViralScoreResponseAnalysisSummaryTarget) RawJSON() string { return r.JSON.raw }
func (r *XAnalysisViralScoreResponseAnalysisSummaryTarget) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// 1 analyzed post with its answers.
type XAnalysisViralScoreResponseResult struct {
	// The answers in full, with their probabilities.
	Analysis XAnalysisViralScoreResponseResultAnalysis `json:"analysis" api:"required"`
	// Each question's decision by question ID. A choice gives its category, a score
	// its level from 0, which can have decimals, & a probability its value from 0 to
	// 1.
	Answers map[string]XAnalysisViralScoreResponseResultAnswersUnion `json:"answers" api:"required"`
	// Cashtags in the post's text, in upper case.
	Cashtags []string `json:"cashtags" api:"required"`
	// How the post's answers compare with its row in baseline, as the AI Actors
	// compare a run with an earlier dataset.
	Monitor XAnalysisViralScoreResponseResultMonitor `json:"monitor" api:"required"`
	// Hostnames the post links to, without t.co.
	SourceDomains []string `json:"sourceDomains" api:"required"`
	// The post, as /x/tweets/search returns it. Your own texts have the IDs text:1,
	// text:2 & so on.
	Tweet shared.SearchTweet `json:"tweet" api:"required"`
	// Only from /x/analysis/viral-score. The score with its weights version, the caps
	// that applied & the engagement comparison.
	Viral PublicMetadata `json:"viral"`
	// Only from /x/analysis/viral-score. log10 of 1 plus the weighted engagement per
	// 1,000 followers. Null without a follower count.
	ViralActualEngagementRate float64 `json:"viralActualEngagementRate" api:"nullable"`
	// Only from /x/analysis/viral-score. The post's public counts under X's published
	// ranking weights, per 1,000 views or followers.
	ViralAlgorithmScore float64 `json:"viralAlgorithmScore" api:"nullable"`
	// Only from /x/analysis/viral-score. How viral the wording reads, or null when a
	// trait has no answer.
	ViralScore int64 `json:"viralScore" api:"nullable"`
	// Only from /x/analysis/viral-score. send_it from 70, edit_first from 40, else
	// sleep_on_it.
	//
	// Any of "send_it", "edit_first", "sleep_on_it".
	ViralVerdict XAnalysisViralScoreResponseResultViralVerdict `json:"viralVerdict" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Analysis                  respjson.Field
		Answers                   respjson.Field
		Cashtags                  respjson.Field
		Monitor                   respjson.Field
		SourceDomains             respjson.Field
		Tweet                     respjson.Field
		Viral                     respjson.Field
		ViralActualEngagementRate respjson.Field
		ViralAlgorithmScore       respjson.Field
		ViralScore                respjson.Field
		ViralVerdict              respjson.Field
		ExtraFields               map[string]respjson.Field
		raw                       string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XAnalysisViralScoreResponseResult) RawJSON() string { return r.JSON.raw }
func (r *XAnalysisViralScoreResponseResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// The answers in full, with their probabilities.
type XAnalysisViralScoreResponseResultAnalysis struct {
	// 1 answer per question, in question order.
	Answers []XAnalysisViralScoreResponseResultAnalysisAnswer `json:"answers" api:"required"`
	// What the AI read beside the post's text.
	ContextAvailability XAnalysisViralScoreResponseResultAnalysisContextAvailability `json:"contextAvailability" api:"required"`
	// Bytes of post & settings the AI read.
	ContextBytes int64 `json:"contextBytes" api:"required"`
	// The ID & version of each question asked.
	Questions []XAnalysisViralScoreResponseResultAnalysisQuestion `json:"questions" api:"required"`
	Status    constant.Succeeded                                  `json:"status" default:"succeeded"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Answers             respjson.Field
		ContextAvailability respjson.Field
		ContextBytes        respjson.Field
		Questions           respjson.Field
		Status              respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XAnalysisViralScoreResponseResultAnalysis) RawJSON() string { return r.JSON.raw }
func (r *XAnalysisViralScoreResponseResultAnalysis) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// 1 question's answer. A probability has probability. A choice or score has value,
// confidence & the probability of each category or level.
type XAnalysisViralScoreResponseResultAnalysisAnswer struct {
	QuestionID      string `json:"questionId" api:"required"`
	QuestionVersion string `json:"questionVersion" api:"required"`
	// Any of "choice", "score", "probability".
	Type       XAnalysisViralScoreResponseResultAnalysisAnswerType `json:"type" api:"required"`
	Confidence float64                                             `json:"confidence"`
	// For a score. The levels the question named, from 0.
	Levels []string `json:"levels"`
	// Probability of each category, or of each level by index.
	Probabilities map[string]float64 `json:"probabilities"`
	// For a probability. The chance the answer is yes.
	Probability float64 `json:"probability"`
	// The chosen category, or the level from 0, which can have decimals, such as 1.4
	// between levels 1 & 2.
	Value XAnalysisViralScoreResponseResultAnalysisAnswerValueUnion `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		QuestionID      respjson.Field
		QuestionVersion respjson.Field
		Type            respjson.Field
		Confidence      respjson.Field
		Levels          respjson.Field
		Probabilities   respjson.Field
		Probability     respjson.Field
		Value           respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XAnalysisViralScoreResponseResultAnalysisAnswer) RawJSON() string { return r.JSON.raw }
func (r *XAnalysisViralScoreResponseResultAnalysisAnswer) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type XAnalysisViralScoreResponseResultAnalysisAnswerType string

const (
	XAnalysisViralScoreResponseResultAnalysisAnswerTypeChoice      XAnalysisViralScoreResponseResultAnalysisAnswerType = "choice"
	XAnalysisViralScoreResponseResultAnalysisAnswerTypeScore       XAnalysisViralScoreResponseResultAnalysisAnswerType = "score"
	XAnalysisViralScoreResponseResultAnalysisAnswerTypeProbability XAnalysisViralScoreResponseResultAnalysisAnswerType = "probability"
)

// XAnalysisViralScoreResponseResultAnalysisAnswerValueUnion contains all possible
// properties and values from [float64], [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfFloat OfString]
type XAnalysisViralScoreResponseResultAnalysisAnswerValueUnion struct {
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	JSON     struct {
		OfFloat  respjson.Field
		OfString respjson.Field
		raw      string
	} `json:"-"`
}

func (u XAnalysisViralScoreResponseResultAnalysisAnswerValueUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v, nil)
	return
}

func (u XAnalysisViralScoreResponseResultAnalysisAnswerValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v, nil)
	return
}

// Returns the unmodified JSON received from the API
func (u XAnalysisViralScoreResponseResultAnalysisAnswerValueUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *XAnalysisViralScoreResponseResultAnalysisAnswerValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// What the AI read beside the post's text.
type XAnalysisViralScoreResponseResultAnalysisContextAvailability struct {
	// Any of "text_blocks", "summary", "not_supplied".
	Article XAnalysisViralScoreResponseResultAnalysisContextAvailabilityArticle `json:"article" api:"required"`
	// Any of "available", "missing".
	Author XAnalysisViralScoreResponseResultAnalysisContextAvailabilityAuthor `json:"author" api:"required"`
	// Any of "text_descriptions_only", "not_supplied".
	Media XAnalysisViralScoreResponseResultAnalysisContextAvailabilityMedia `json:"media" api:"required"`
	// truncated when the post was cut to fit.
	//
	// Any of "full", "truncated".
	PostText XAnalysisViralScoreResponseResultAnalysisContextAvailabilityPostText `json:"postText" api:"required"`
	// Any of "available", "missing", "not_applicable".
	Quote XAnalysisViralScoreResponseResultAnalysisContextAvailabilityQuote `json:"quote" api:"required"`
	// Any of "available", "missing", "not_applicable".
	Reply XAnalysisViralScoreResponseResultAnalysisContextAvailabilityReply `json:"reply" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Article     respjson.Field
		Author      respjson.Field
		Media       respjson.Field
		PostText    respjson.Field
		Quote       respjson.Field
		Reply       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XAnalysisViralScoreResponseResultAnalysisContextAvailability) RawJSON() string {
	return r.JSON.raw
}
func (r *XAnalysisViralScoreResponseResultAnalysisContextAvailability) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type XAnalysisViralScoreResponseResultAnalysisContextAvailabilityArticle string

const (
	XAnalysisViralScoreResponseResultAnalysisContextAvailabilityArticleTextBlocks  XAnalysisViralScoreResponseResultAnalysisContextAvailabilityArticle = "text_blocks"
	XAnalysisViralScoreResponseResultAnalysisContextAvailabilityArticleSummary     XAnalysisViralScoreResponseResultAnalysisContextAvailabilityArticle = "summary"
	XAnalysisViralScoreResponseResultAnalysisContextAvailabilityArticleNotSupplied XAnalysisViralScoreResponseResultAnalysisContextAvailabilityArticle = "not_supplied"
)

type XAnalysisViralScoreResponseResultAnalysisContextAvailabilityAuthor string

const (
	XAnalysisViralScoreResponseResultAnalysisContextAvailabilityAuthorAvailable XAnalysisViralScoreResponseResultAnalysisContextAvailabilityAuthor = "available"
	XAnalysisViralScoreResponseResultAnalysisContextAvailabilityAuthorMissing   XAnalysisViralScoreResponseResultAnalysisContextAvailabilityAuthor = "missing"
)

type XAnalysisViralScoreResponseResultAnalysisContextAvailabilityMedia string

const (
	XAnalysisViralScoreResponseResultAnalysisContextAvailabilityMediaTextDescriptionsOnly XAnalysisViralScoreResponseResultAnalysisContextAvailabilityMedia = "text_descriptions_only"
	XAnalysisViralScoreResponseResultAnalysisContextAvailabilityMediaNotSupplied          XAnalysisViralScoreResponseResultAnalysisContextAvailabilityMedia = "not_supplied"
)

// truncated when the post was cut to fit.
type XAnalysisViralScoreResponseResultAnalysisContextAvailabilityPostText string

const (
	XAnalysisViralScoreResponseResultAnalysisContextAvailabilityPostTextFull      XAnalysisViralScoreResponseResultAnalysisContextAvailabilityPostText = "full"
	XAnalysisViralScoreResponseResultAnalysisContextAvailabilityPostTextTruncated XAnalysisViralScoreResponseResultAnalysisContextAvailabilityPostText = "truncated"
)

type XAnalysisViralScoreResponseResultAnalysisContextAvailabilityQuote string

const (
	XAnalysisViralScoreResponseResultAnalysisContextAvailabilityQuoteAvailable     XAnalysisViralScoreResponseResultAnalysisContextAvailabilityQuote = "available"
	XAnalysisViralScoreResponseResultAnalysisContextAvailabilityQuoteMissing       XAnalysisViralScoreResponseResultAnalysisContextAvailabilityQuote = "missing"
	XAnalysisViralScoreResponseResultAnalysisContextAvailabilityQuoteNotApplicable XAnalysisViralScoreResponseResultAnalysisContextAvailabilityQuote = "not_applicable"
)

type XAnalysisViralScoreResponseResultAnalysisContextAvailabilityReply string

const (
	XAnalysisViralScoreResponseResultAnalysisContextAvailabilityReplyAvailable     XAnalysisViralScoreResponseResultAnalysisContextAvailabilityReply = "available"
	XAnalysisViralScoreResponseResultAnalysisContextAvailabilityReplyMissing       XAnalysisViralScoreResponseResultAnalysisContextAvailabilityReply = "missing"
	XAnalysisViralScoreResponseResultAnalysisContextAvailabilityReplyNotApplicable XAnalysisViralScoreResponseResultAnalysisContextAvailabilityReply = "not_applicable"
)

type XAnalysisViralScoreResponseResultAnalysisQuestion struct {
	ID      string `json:"id" api:"required"`
	Version string `json:"version" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Version     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XAnalysisViralScoreResponseResultAnalysisQuestion) RawJSON() string { return r.JSON.raw }
func (r *XAnalysisViralScoreResponseResultAnalysisQuestion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// XAnalysisViralScoreResponseResultAnswersUnion contains all possible properties
// and values from [float64], [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfFloat OfString]
type XAnalysisViralScoreResponseResultAnswersUnion struct {
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	JSON     struct {
		OfFloat  respjson.Field
		OfString respjson.Field
		raw      string
	} `json:"-"`
}

func (u XAnalysisViralScoreResponseResultAnswersUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v, nil)
	return
}

func (u XAnalysisViralScoreResponseResultAnswersUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v, nil)
	return
}

// Returns the unmodified JSON received from the API
func (u XAnalysisViralScoreResponseResultAnswersUnion) RawJSON() string { return u.JSON.raw }

func (r *XAnalysisViralScoreResponseResultAnswersUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// How the post's answers compare with its row in baseline, as the AI Actors
// compare a run with an earlier dataset.
type XAnalysisViralScoreResponseResultMonitor struct {
	// Always null. The baseline is the rows you sent.
	BaselineDatasetID json.RawMessage `json:"baselineDatasetId" api:"required"`
	// IDs of the questions whose decision changed.
	ChangedQuestionIDs []string `json:"changedQuestionIds" api:"required"`
	// Each changed decision, earlier & now.
	Changes []XAnalysisViralScoreResponseResultMonitorChange `json:"changes" api:"required"`
	// Hash of the questions, targets & context. Rows with another hash are
	// not_comparable.
	ConfigurationFingerprint string `json:"configurationFingerprint" api:"required"`
	// first_run without baseline. new_to_baseline for a post baseline lacks. changed
	// when an answer clearly left its earlier decision. not_comparable when other
	// settings made the earlier row or it answers none of these questions.
	//
	// Any of "first_run", "new_to_baseline", "unchanged", "changed", "not_comparable",
	// "analysis_unavailable".
	Status XAnalysisViralScoreResponseResultMonitorStatus `json:"status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BaselineDatasetID        respjson.Field
		ChangedQuestionIDs       respjson.Field
		Changes                  respjson.Field
		ConfigurationFingerprint respjson.Field
		Status                   respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XAnalysisViralScoreResponseResultMonitor) RawJSON() string { return r.JSON.raw }
func (r *XAnalysisViralScoreResponseResultMonitor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type XAnalysisViralScoreResponseResultMonitorChange struct {
	Current    string `json:"current" api:"required"`
	Previous   string `json:"previous" api:"required"`
	QuestionID string `json:"questionId" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Current     respjson.Field
		Previous    respjson.Field
		QuestionID  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XAnalysisViralScoreResponseResultMonitorChange) RawJSON() string { return r.JSON.raw }
func (r *XAnalysisViralScoreResponseResultMonitorChange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

// first_run without baseline. new_to_baseline for a post baseline lacks. changed
// when an answer clearly left its earlier decision. not_comparable when other
// settings made the earlier row or it answers none of these questions.
type XAnalysisViralScoreResponseResultMonitorStatus string

const (
	XAnalysisViralScoreResponseResultMonitorStatusFirstRun            XAnalysisViralScoreResponseResultMonitorStatus = "first_run"
	XAnalysisViralScoreResponseResultMonitorStatusNewToBaseline       XAnalysisViralScoreResponseResultMonitorStatus = "new_to_baseline"
	XAnalysisViralScoreResponseResultMonitorStatusUnchanged           XAnalysisViralScoreResponseResultMonitorStatus = "unchanged"
	XAnalysisViralScoreResponseResultMonitorStatusChanged             XAnalysisViralScoreResponseResultMonitorStatus = "changed"
	XAnalysisViralScoreResponseResultMonitorStatusNotComparable       XAnalysisViralScoreResponseResultMonitorStatus = "not_comparable"
	XAnalysisViralScoreResponseResultMonitorStatusAnalysisUnavailable XAnalysisViralScoreResponseResultMonitorStatus = "analysis_unavailable"
)

// Only from /x/analysis/viral-score. send_it from 70, edit_first from 40, else
// sleep_on_it.
type XAnalysisViralScoreResponseResultViralVerdict string

const (
	XAnalysisViralScoreResponseResultViralVerdictSendIt    XAnalysisViralScoreResponseResultViralVerdict = "send_it"
	XAnalysisViralScoreResponseResultViralVerdictEditFirst XAnalysisViralScoreResponseResultViralVerdict = "edit_first"
	XAnalysisViralScoreResponseResultViralVerdictSleepOnIt XAnalysisViralScoreResponseResultViralVerdict = "sleep_on_it"
)

// A post without an analysis. It costs nothing.
type XAnalysisViralScoreResponseUnanalyzed struct {
	// Post ID. Your own texts have the IDs text:1, text:2 & so on.
	ID string `json:"id" api:"required"`
	// post_unavailable when X returns no such post. missing_text for a post without
	// text. insufficient_credits when credits ran out. context_limit when the settings
	// leave the post no room. The other reasons mean the analysis failed. Send the
	// post again.
	Reason string `json:"reason" api:"required"`
	// Any of "failed", "skipped".
	Status XAnalysisViralScoreResponseUnanalyzedStatus `json:"status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Reason      respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XAnalysisViralScoreResponseUnanalyzed) RawJSON() string { return r.JSON.raw }
func (r *XAnalysisViralScoreResponseUnanalyzed) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, &r.JSON.raw)
}

type XAnalysisViralScoreResponseUnanalyzedStatus string

const (
	XAnalysisViralScoreResponseUnanalyzedStatusFailed  XAnalysisViralScoreResponseUnanalyzedStatus = "failed"
	XAnalysisViralScoreResponseUnanalyzedStatusSkipped XAnalysisViralScoreResponseUnanalyzedStatus = "skipped"
)

type XAnalysisBrandParams struct {
	// The next_cursor of the page before. Send the same search with it.
	Cursor param.Opt[string] `json:"cursor,omitzero"`
	// Maximum posts to analyze from the search. A low credit balance can return fewer.
	Limit param.Opt[int64] `json:"limit,omitzero"`
	// Analyze a public List's posts. Send its ID or URL.
	ListID param.Opt[string] `json:"listId,omitzero"`
	// X search to analyze, with the operators of /x/tweets/search, such as from:nasa
	// or lang:en. X searches up to 512 characters, counted with the other search
	// fields, as /x/tweets/search counts q.
	Query param.Opt[string] `json:"query,omitzero"`
	// Analyze the quotes of 1 post. Send its ID or URL.
	QuotesOf param.Opt[string] `json:"quotesOf,omitzero"`
	// Analyze the replies in 1 post's conversation. Send its ID or URL.
	RepliesTo param.Opt[string] `json:"repliesTo,omitzero"`
	// Analyze 1 account's posts. Send a handle, with or without @, or a profile URL.
	Username param.Opt[string] `json:"username,omitzero"`
	// Analysis settings, the same as the analysis input of Xquik's AI Actors. Leave it
	// out for the route's own questions.
	Analysis XAnalysisBrandParamsAnalysis `json:"analysis,omitzero"`
	// Rows of an earlier answer to compare with, such as its results. Each needs the
	// post's id, or tweet.id, & its answers. Send monitor too, so other settings show
	// as not_comparable. Each result's monitor says whether the post is new, changed
	// or unchanged. The comparison costs nothing extra.
	Baseline []XAnalysisBrandParamsBaseline `json:"baseline,omitzero"`
	// Search order. Default Latest.
	//
	// Any of "Latest", "Top".
	QueryType XAnalysisBrandParamsQueryType `json:"queryType,omitzero"`
	// Inclusive start of the search, such as 2026-09-25T00:00:00Z. Unix seconds also
	// work. A time without an offset is UTC.
	SinceTime XAnalysisBrandParamsSinceTimeUnion `json:"sinceTime,omitzero"`
	// Your own texts to analyze, such as drafts. Reads nothing from X. Each text costs
	// the same as 1 analyzed post.
	Texts []string `json:"texts,omitzero"`
	// Post IDs or post URLs to analyze, up to 100.
	TweetIDs []string `json:"tweetIds,omitzero"`
	// Exclusive end of the search, such as 2026-09-26T00:00:00Z. Unix seconds also
	// work. A time without an offset is UTC.
	UntilTime XAnalysisBrandParamsUntilTimeUnion `json:"untilTime,omitzero"`
	paramObj
}

func (r XAnalysisBrandParams) MarshalJSON() (data []byte, err error) {
	type shadow XAnalysisBrandParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *XAnalysisBrandParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, nil)
}

// Analysis settings, the same as the analysis input of Xquik's AI Actors. Leave it
// out for the route's own questions.
type XAnalysisBrandParamsAnalysis struct {
	// Posts analyzed at the same time.
	Concurrency param.Opt[int64] `json:"concurrency,omitzero"`
	// Background the AI reads with every post.
	Context param.Opt[string] `json:"context,omitzero"`
	// Bytes of post, quote & settings the AI reads. Xquik cuts longer posts to fit.
	// Empty settings take 219 bytes, so a smaller value gets a 400.
	MaxContextBytes param.Opt[int64] `json:"maxContextBytes,omitzero"`
	// A ready set of questions, used when questions is left out. complaints,
	// competitors, purchase_intent & product_feedback need targets. viral_lite scores
	// how viral the wording reads.
	//
	// Any of "brand", "complaints", "competitors", "purchase_intent",
	// "product_feedback", "news", "sentiment", "market", "viral_lite".
	Preset string `json:"preset,omitzero"`
	// Your own questions, which replace the route's questions. Write question IDs &
	// category names in lower case without hyphens.
	Questions []XAnalysisBrandParamsAnalysisQuestion `json:"questions,omitzero"`
	// Brands, people, assets or topics the questions are about, each with its aliases.
	Targets []XAnalysisBrandParamsAnalysisTarget `json:"targets,omitzero"`
	paramObj
}

func (r XAnalysisBrandParamsAnalysis) MarshalJSON() (data []byte, err error) {
	type shadow XAnalysisBrandParamsAnalysis
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *XAnalysisBrandParamsAnalysis) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, nil)
}

func init() {
	apijson.RegisterFieldValidator[XAnalysisBrandParamsAnalysis](
		"preset", "brand", "complaints", "competitors", "purchase_intent", "product_feedback", "news", "sentiment", "market", "viral_lite",
	)
}

// A question the AI answers for each post. A choice picks 1 of 2 to 255
// categories, a score picks 1 of 2 or more ordered levels & a probability answers
// yes or no.
//
// The properties ID, Instructions, Type, Version are required.
type XAnalysisBrandParamsAnalysisQuestion struct {
	// Name of the answer, such as request. A name the response renames or drops gets a
	// 400, such as created_at, favorite_count, muting or blocking.
	ID string `json:"id" api:"required"`
	// The question, as the AI reads it.
	Instructions string `json:"instructions" api:"required"`
	// Any of "choice", "score", "probability".
	Type string `json:"type,omitzero" api:"required"`
	// Your version of the question, such as 1.
	Version string `json:"version" api:"required"`
	// For a choice. Each category name with its description, or null.
	Categories map[string]param.Opt[string] `json:"categories,omitzero"`
	// For a probability. What counts as yes & as no.
	Criteria XAnalysisBrandParamsAnalysisQuestionCriteria `json:"criteria,omitzero"`
	// For a score. Each level's description, lowest first.
	Levels []string `json:"levels,omitzero"`
	paramObj
}

func (r XAnalysisBrandParamsAnalysisQuestion) MarshalJSON() (data []byte, err error) {
	type shadow XAnalysisBrandParamsAnalysisQuestion
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *XAnalysisBrandParamsAnalysisQuestion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, nil)
}

func init() {
	apijson.RegisterFieldValidator[XAnalysisBrandParamsAnalysisQuestion](
		"type", "choice", "score", "probability",
	)
}

// For a probability. What counts as yes & as no.
//
// The properties No, Yes are required.
type XAnalysisBrandParamsAnalysisQuestionCriteria struct {
	No  string `json:"no" api:"required"`
	Yes string `json:"yes" api:"required"`
	paramObj
}

func (r XAnalysisBrandParamsAnalysisQuestionCriteria) MarshalJSON() (data []byte, err error) {
	type shadow XAnalysisBrandParamsAnalysisQuestionCriteria
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *XAnalysisBrandParamsAnalysisQuestionCriteria) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, nil)
}

// The property Name is required.
type XAnalysisBrandParamsAnalysisTarget struct {
	Name    string   `json:"name" api:"required"`
	Aliases []string `json:"aliases,omitzero"`
	paramObj
}

func (r XAnalysisBrandParamsAnalysisTarget) MarshalJSON() (data []byte, err error) {
	type shadow XAnalysisBrandParamsAnalysisTarget
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *XAnalysisBrandParamsAnalysisTarget) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, nil)
}

// The property Answers is required.
type XAnalysisBrandParamsBaseline struct {
	// The earlier answers by question ID.
	Answers map[string]XAnalysisBrandParamsBaselineAnswersUnion `json:"answers,omitzero" api:"required"`
	// Post ID, or text:1 & so on. Otherwise tweet.id counts.
	ID param.Opt[string] `json:"id,omitzero"`
	// The earlier row's monitor, with its fingerprint.
	Monitor XAnalysisBrandParamsBaselineMonitor `json:"monitor,omitzero"`
	Tweet   XAnalysisBrandParamsBaselineTweet   `json:"tweet,omitzero"`
	paramObj
}

func (r XAnalysisBrandParamsBaseline) MarshalJSON() (data []byte, err error) {
	type shadow XAnalysisBrandParamsBaseline
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *XAnalysisBrandParamsBaseline) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, nil)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type XAnalysisBrandParamsBaselineAnswersUnion struct {
	OfString param.Opt[string]  `json:",omitzero,inline"`
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	paramUnion
}

func (u XAnalysisBrandParamsBaselineAnswersUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfFloat)
}
func (u *XAnalysisBrandParamsBaselineAnswersUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u, nil)
}

// The earlier row's monitor, with its fingerprint.
type XAnalysisBrandParamsBaselineMonitor struct {
	ConfigurationFingerprint param.Opt[string] `json:"configurationFingerprint,omitzero"`
	paramObj
}

func (r XAnalysisBrandParamsBaselineMonitor) MarshalJSON() (data []byte, err error) {
	type shadow XAnalysisBrandParamsBaselineMonitor
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *XAnalysisBrandParamsBaselineMonitor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, nil)
}

// The property ID is required.
type XAnalysisBrandParamsBaselineTweet struct {
	ID string `json:"id" api:"required"`
	paramObj
}

func (r XAnalysisBrandParamsBaselineTweet) MarshalJSON() (data []byte, err error) {
	type shadow XAnalysisBrandParamsBaselineTweet
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *XAnalysisBrandParamsBaselineTweet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, nil)
}

// Search order. Default Latest.
type XAnalysisBrandParamsQueryType string

const (
	XAnalysisBrandParamsQueryTypeLatest XAnalysisBrandParamsQueryType = "Latest"
	XAnalysisBrandParamsQueryTypeTop    XAnalysisBrandParamsQueryType = "Top"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type XAnalysisBrandParamsSinceTimeUnion struct {
	OfString param.Opt[string] `json:",omitzero,inline"`
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	paramUnion
}

func (u XAnalysisBrandParamsSinceTimeUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfInt)
}
func (u *XAnalysisBrandParamsSinceTimeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u, nil)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type XAnalysisBrandParamsUntilTimeUnion struct {
	OfString param.Opt[string] `json:",omitzero,inline"`
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	paramUnion
}

func (u XAnalysisBrandParamsUntilTimeUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfInt)
}
func (u *XAnalysisBrandParamsUntilTimeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u, nil)
}

type XAnalysisClassifyParams struct {
	// The next_cursor of the page before. Send the same search with it.
	Cursor param.Opt[string] `json:"cursor,omitzero"`
	// Maximum posts to analyze from the search. A low credit balance can return fewer.
	Limit param.Opt[int64] `json:"limit,omitzero"`
	// Analyze a public List's posts. Send its ID or URL.
	ListID param.Opt[string] `json:"listId,omitzero"`
	// X search to analyze, with the operators of /x/tweets/search, such as from:nasa
	// or lang:en. X searches up to 512 characters, counted with the other search
	// fields, as /x/tweets/search counts q.
	Query param.Opt[string] `json:"query,omitzero"`
	// Analyze the quotes of 1 post. Send its ID or URL.
	QuotesOf param.Opt[string] `json:"quotesOf,omitzero"`
	// Analyze the replies in 1 post's conversation. Send its ID or URL.
	RepliesTo param.Opt[string] `json:"repliesTo,omitzero"`
	// Analyze 1 account's posts. Send a handle, with or without @, or a profile URL.
	Username param.Opt[string] `json:"username,omitzero"`
	// Analysis settings, the same as the analysis input of Xquik's AI Actors. Leave it
	// out for the route's own questions.
	Analysis XAnalysisClassifyParamsAnalysis `json:"analysis,omitzero"`
	// Rows of an earlier answer to compare with, such as its results. Each needs the
	// post's id, or tweet.id, & its answers. Send monitor too, so other settings show
	// as not_comparable. Each result's monitor says whether the post is new, changed
	// or unchanged. The comparison costs nothing extra.
	Baseline []XAnalysisClassifyParamsBaseline `json:"baseline,omitzero"`
	// Search order. Default Latest.
	//
	// Any of "Latest", "Top".
	QueryType XAnalysisClassifyParamsQueryType `json:"queryType,omitzero"`
	// Inclusive start of the search, such as 2026-09-25T00:00:00Z. Unix seconds also
	// work. A time without an offset is UTC.
	SinceTime XAnalysisClassifyParamsSinceTimeUnion `json:"sinceTime,omitzero"`
	// Your own texts to analyze, such as drafts. Reads nothing from X. Each text costs
	// the same as 1 analyzed post.
	Texts []string `json:"texts,omitzero"`
	// Post IDs or post URLs to analyze, up to 100.
	TweetIDs []string `json:"tweetIds,omitzero"`
	// Exclusive end of the search, such as 2026-09-26T00:00:00Z. Unix seconds also
	// work. A time without an offset is UTC.
	UntilTime XAnalysisClassifyParamsUntilTimeUnion `json:"untilTime,omitzero"`
	paramObj
}

func (r XAnalysisClassifyParams) MarshalJSON() (data []byte, err error) {
	type shadow XAnalysisClassifyParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *XAnalysisClassifyParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, nil)
}

// Analysis settings, the same as the analysis input of Xquik's AI Actors. Leave it
// out for the route's own questions.
type XAnalysisClassifyParamsAnalysis struct {
	// Posts analyzed at the same time.
	Concurrency param.Opt[int64] `json:"concurrency,omitzero"`
	// Background the AI reads with every post.
	Context param.Opt[string] `json:"context,omitzero"`
	// Bytes of post, quote & settings the AI reads. Xquik cuts longer posts to fit.
	// Empty settings take 219 bytes, so a smaller value gets a 400.
	MaxContextBytes param.Opt[int64] `json:"maxContextBytes,omitzero"`
	// A ready set of questions, used when questions is left out. complaints,
	// competitors, purchase_intent & product_feedback need targets. viral_lite scores
	// how viral the wording reads.
	//
	// Any of "brand", "complaints", "competitors", "purchase_intent",
	// "product_feedback", "news", "sentiment", "market", "viral_lite".
	Preset string `json:"preset,omitzero"`
	// Your own questions, which replace the route's questions. Write question IDs &
	// category names in lower case without hyphens.
	Questions []XAnalysisClassifyParamsAnalysisQuestion `json:"questions,omitzero"`
	// Brands, people, assets or topics the questions are about, each with its aliases.
	Targets []XAnalysisClassifyParamsAnalysisTarget `json:"targets,omitzero"`
	paramObj
}

func (r XAnalysisClassifyParamsAnalysis) MarshalJSON() (data []byte, err error) {
	type shadow XAnalysisClassifyParamsAnalysis
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *XAnalysisClassifyParamsAnalysis) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, nil)
}

func init() {
	apijson.RegisterFieldValidator[XAnalysisClassifyParamsAnalysis](
		"preset", "brand", "complaints", "competitors", "purchase_intent", "product_feedback", "news", "sentiment", "market", "viral_lite",
	)
}

// A question the AI answers for each post. A choice picks 1 of 2 to 255
// categories, a score picks 1 of 2 or more ordered levels & a probability answers
// yes or no.
//
// The properties ID, Instructions, Type, Version are required.
type XAnalysisClassifyParamsAnalysisQuestion struct {
	// Name of the answer, such as request. A name the response renames or drops gets a
	// 400, such as created_at, favorite_count, muting or blocking.
	ID string `json:"id" api:"required"`
	// The question, as the AI reads it.
	Instructions string `json:"instructions" api:"required"`
	// Any of "choice", "score", "probability".
	Type string `json:"type,omitzero" api:"required"`
	// Your version of the question, such as 1.
	Version string `json:"version" api:"required"`
	// For a choice. Each category name with its description, or null.
	Categories map[string]param.Opt[string] `json:"categories,omitzero"`
	// For a probability. What counts as yes & as no.
	Criteria XAnalysisClassifyParamsAnalysisQuestionCriteria `json:"criteria,omitzero"`
	// For a score. Each level's description, lowest first.
	Levels []string `json:"levels,omitzero"`
	paramObj
}

func (r XAnalysisClassifyParamsAnalysisQuestion) MarshalJSON() (data []byte, err error) {
	type shadow XAnalysisClassifyParamsAnalysisQuestion
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *XAnalysisClassifyParamsAnalysisQuestion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, nil)
}

func init() {
	apijson.RegisterFieldValidator[XAnalysisClassifyParamsAnalysisQuestion](
		"type", "choice", "score", "probability",
	)
}

// For a probability. What counts as yes & as no.
//
// The properties No, Yes are required.
type XAnalysisClassifyParamsAnalysisQuestionCriteria struct {
	No  string `json:"no" api:"required"`
	Yes string `json:"yes" api:"required"`
	paramObj
}

func (r XAnalysisClassifyParamsAnalysisQuestionCriteria) MarshalJSON() (data []byte, err error) {
	type shadow XAnalysisClassifyParamsAnalysisQuestionCriteria
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *XAnalysisClassifyParamsAnalysisQuestionCriteria) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, nil)
}

// The property Name is required.
type XAnalysisClassifyParamsAnalysisTarget struct {
	Name    string   `json:"name" api:"required"`
	Aliases []string `json:"aliases,omitzero"`
	paramObj
}

func (r XAnalysisClassifyParamsAnalysisTarget) MarshalJSON() (data []byte, err error) {
	type shadow XAnalysisClassifyParamsAnalysisTarget
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *XAnalysisClassifyParamsAnalysisTarget) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, nil)
}

// The property Answers is required.
type XAnalysisClassifyParamsBaseline struct {
	// The earlier answers by question ID.
	Answers map[string]XAnalysisClassifyParamsBaselineAnswersUnion `json:"answers,omitzero" api:"required"`
	// Post ID, or text:1 & so on. Otherwise tweet.id counts.
	ID param.Opt[string] `json:"id,omitzero"`
	// The earlier row's monitor, with its fingerprint.
	Monitor XAnalysisClassifyParamsBaselineMonitor `json:"monitor,omitzero"`
	Tweet   XAnalysisClassifyParamsBaselineTweet   `json:"tweet,omitzero"`
	paramObj
}

func (r XAnalysisClassifyParamsBaseline) MarshalJSON() (data []byte, err error) {
	type shadow XAnalysisClassifyParamsBaseline
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *XAnalysisClassifyParamsBaseline) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, nil)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type XAnalysisClassifyParamsBaselineAnswersUnion struct {
	OfString param.Opt[string]  `json:",omitzero,inline"`
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	paramUnion
}

func (u XAnalysisClassifyParamsBaselineAnswersUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfFloat)
}
func (u *XAnalysisClassifyParamsBaselineAnswersUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u, nil)
}

// The earlier row's monitor, with its fingerprint.
type XAnalysisClassifyParamsBaselineMonitor struct {
	ConfigurationFingerprint param.Opt[string] `json:"configurationFingerprint,omitzero"`
	paramObj
}

func (r XAnalysisClassifyParamsBaselineMonitor) MarshalJSON() (data []byte, err error) {
	type shadow XAnalysisClassifyParamsBaselineMonitor
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *XAnalysisClassifyParamsBaselineMonitor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, nil)
}

// The property ID is required.
type XAnalysisClassifyParamsBaselineTweet struct {
	ID string `json:"id" api:"required"`
	paramObj
}

func (r XAnalysisClassifyParamsBaselineTweet) MarshalJSON() (data []byte, err error) {
	type shadow XAnalysisClassifyParamsBaselineTweet
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *XAnalysisClassifyParamsBaselineTweet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, nil)
}

// Search order. Default Latest.
type XAnalysisClassifyParamsQueryType string

const (
	XAnalysisClassifyParamsQueryTypeLatest XAnalysisClassifyParamsQueryType = "Latest"
	XAnalysisClassifyParamsQueryTypeTop    XAnalysisClassifyParamsQueryType = "Top"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type XAnalysisClassifyParamsSinceTimeUnion struct {
	OfString param.Opt[string] `json:",omitzero,inline"`
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	paramUnion
}

func (u XAnalysisClassifyParamsSinceTimeUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfInt)
}
func (u *XAnalysisClassifyParamsSinceTimeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u, nil)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type XAnalysisClassifyParamsUntilTimeUnion struct {
	OfString param.Opt[string] `json:",omitzero,inline"`
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	paramUnion
}

func (u XAnalysisClassifyParamsUntilTimeUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfInt)
}
func (u *XAnalysisClassifyParamsUntilTimeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u, nil)
}

type XAnalysisMarketSignalsParams struct {
	// The next_cursor of the page before. Send the same search with it.
	Cursor param.Opt[string] `json:"cursor,omitzero"`
	// Maximum posts to analyze from the search. A low credit balance can return fewer.
	Limit param.Opt[int64] `json:"limit,omitzero"`
	// Analyze a public List's posts. Send its ID or URL.
	ListID param.Opt[string] `json:"listId,omitzero"`
	// X search to analyze, with the operators of /x/tweets/search, such as from:nasa
	// or lang:en. X searches up to 512 characters, counted with the other search
	// fields, as /x/tweets/search counts q.
	Query param.Opt[string] `json:"query,omitzero"`
	// Analyze the quotes of 1 post. Send its ID or URL.
	QuotesOf param.Opt[string] `json:"quotesOf,omitzero"`
	// Analyze the replies in 1 post's conversation. Send its ID or URL.
	RepliesTo param.Opt[string] `json:"repliesTo,omitzero"`
	// Analyze 1 account's posts. Send a handle, with or without @, or a profile URL.
	Username param.Opt[string] `json:"username,omitzero"`
	// Analysis settings, the same as the analysis input of Xquik's AI Actors. Leave it
	// out for the route's own questions.
	Analysis XAnalysisMarketSignalsParamsAnalysis `json:"analysis,omitzero"`
	// Rows of an earlier answer to compare with, such as its results. Each needs the
	// post's id, or tweet.id, & its answers. Send monitor too, so other settings show
	// as not_comparable. Each result's monitor says whether the post is new, changed
	// or unchanged. The comparison costs nothing extra.
	Baseline []XAnalysisMarketSignalsParamsBaseline `json:"baseline,omitzero"`
	// Search order. Default Latest.
	//
	// Any of "Latest", "Top".
	QueryType XAnalysisMarketSignalsParamsQueryType `json:"queryType,omitzero"`
	// Inclusive start of the search, such as 2026-09-25T00:00:00Z. Unix seconds also
	// work. A time without an offset is UTC.
	SinceTime XAnalysisMarketSignalsParamsSinceTimeUnion `json:"sinceTime,omitzero"`
	// Your own texts to analyze, such as drafts. Reads nothing from X. Each text costs
	// the same as 1 analyzed post.
	Texts []string `json:"texts,omitzero"`
	// Post IDs or post URLs to analyze, up to 100.
	TweetIDs []string `json:"tweetIds,omitzero"`
	// Exclusive end of the search, such as 2026-09-26T00:00:00Z. Unix seconds also
	// work. A time without an offset is UTC.
	UntilTime XAnalysisMarketSignalsParamsUntilTimeUnion `json:"untilTime,omitzero"`
	paramObj
}

func (r XAnalysisMarketSignalsParams) MarshalJSON() (data []byte, err error) {
	type shadow XAnalysisMarketSignalsParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *XAnalysisMarketSignalsParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, nil)
}

// Analysis settings, the same as the analysis input of Xquik's AI Actors. Leave it
// out for the route's own questions.
type XAnalysisMarketSignalsParamsAnalysis struct {
	// Posts analyzed at the same time.
	Concurrency param.Opt[int64] `json:"concurrency,omitzero"`
	// Background the AI reads with every post.
	Context param.Opt[string] `json:"context,omitzero"`
	// Bytes of post, quote & settings the AI reads. Xquik cuts longer posts to fit.
	// Empty settings take 219 bytes, so a smaller value gets a 400.
	MaxContextBytes param.Opt[int64] `json:"maxContextBytes,omitzero"`
	// A ready set of questions, used when questions is left out. complaints,
	// competitors, purchase_intent & product_feedback need targets. viral_lite scores
	// how viral the wording reads.
	//
	// Any of "brand", "complaints", "competitors", "purchase_intent",
	// "product_feedback", "news", "sentiment", "market", "viral_lite".
	Preset string `json:"preset,omitzero"`
	// Your own questions, which replace the route's questions. Write question IDs &
	// category names in lower case without hyphens.
	Questions []XAnalysisMarketSignalsParamsAnalysisQuestion `json:"questions,omitzero"`
	// Brands, people, assets or topics the questions are about, each with its aliases.
	Targets []XAnalysisMarketSignalsParamsAnalysisTarget `json:"targets,omitzero"`
	paramObj
}

func (r XAnalysisMarketSignalsParamsAnalysis) MarshalJSON() (data []byte, err error) {
	type shadow XAnalysisMarketSignalsParamsAnalysis
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *XAnalysisMarketSignalsParamsAnalysis) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, nil)
}

func init() {
	apijson.RegisterFieldValidator[XAnalysisMarketSignalsParamsAnalysis](
		"preset", "brand", "complaints", "competitors", "purchase_intent", "product_feedback", "news", "sentiment", "market", "viral_lite",
	)
}

// A question the AI answers for each post. A choice picks 1 of 2 to 255
// categories, a score picks 1 of 2 or more ordered levels & a probability answers
// yes or no.
//
// The properties ID, Instructions, Type, Version are required.
type XAnalysisMarketSignalsParamsAnalysisQuestion struct {
	// Name of the answer, such as request. A name the response renames or drops gets a
	// 400, such as created_at, favorite_count, muting or blocking.
	ID string `json:"id" api:"required"`
	// The question, as the AI reads it.
	Instructions string `json:"instructions" api:"required"`
	// Any of "choice", "score", "probability".
	Type string `json:"type,omitzero" api:"required"`
	// Your version of the question, such as 1.
	Version string `json:"version" api:"required"`
	// For a choice. Each category name with its description, or null.
	Categories map[string]param.Opt[string] `json:"categories,omitzero"`
	// For a probability. What counts as yes & as no.
	Criteria XAnalysisMarketSignalsParamsAnalysisQuestionCriteria `json:"criteria,omitzero"`
	// For a score. Each level's description, lowest first.
	Levels []string `json:"levels,omitzero"`
	paramObj
}

func (r XAnalysisMarketSignalsParamsAnalysisQuestion) MarshalJSON() (data []byte, err error) {
	type shadow XAnalysisMarketSignalsParamsAnalysisQuestion
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *XAnalysisMarketSignalsParamsAnalysisQuestion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, nil)
}

func init() {
	apijson.RegisterFieldValidator[XAnalysisMarketSignalsParamsAnalysisQuestion](
		"type", "choice", "score", "probability",
	)
}

// For a probability. What counts as yes & as no.
//
// The properties No, Yes are required.
type XAnalysisMarketSignalsParamsAnalysisQuestionCriteria struct {
	No  string `json:"no" api:"required"`
	Yes string `json:"yes" api:"required"`
	paramObj
}

func (r XAnalysisMarketSignalsParamsAnalysisQuestionCriteria) MarshalJSON() (data []byte, err error) {
	type shadow XAnalysisMarketSignalsParamsAnalysisQuestionCriteria
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *XAnalysisMarketSignalsParamsAnalysisQuestionCriteria) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, nil)
}

// The property Name is required.
type XAnalysisMarketSignalsParamsAnalysisTarget struct {
	Name    string   `json:"name" api:"required"`
	Aliases []string `json:"aliases,omitzero"`
	paramObj
}

func (r XAnalysisMarketSignalsParamsAnalysisTarget) MarshalJSON() (data []byte, err error) {
	type shadow XAnalysisMarketSignalsParamsAnalysisTarget
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *XAnalysisMarketSignalsParamsAnalysisTarget) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, nil)
}

// The property Answers is required.
type XAnalysisMarketSignalsParamsBaseline struct {
	// The earlier answers by question ID.
	Answers map[string]XAnalysisMarketSignalsParamsBaselineAnswersUnion `json:"answers,omitzero" api:"required"`
	// Post ID, or text:1 & so on. Otherwise tweet.id counts.
	ID param.Opt[string] `json:"id,omitzero"`
	// The earlier row's monitor, with its fingerprint.
	Monitor XAnalysisMarketSignalsParamsBaselineMonitor `json:"monitor,omitzero"`
	Tweet   XAnalysisMarketSignalsParamsBaselineTweet   `json:"tweet,omitzero"`
	paramObj
}

func (r XAnalysisMarketSignalsParamsBaseline) MarshalJSON() (data []byte, err error) {
	type shadow XAnalysisMarketSignalsParamsBaseline
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *XAnalysisMarketSignalsParamsBaseline) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, nil)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type XAnalysisMarketSignalsParamsBaselineAnswersUnion struct {
	OfString param.Opt[string]  `json:",omitzero,inline"`
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	paramUnion
}

func (u XAnalysisMarketSignalsParamsBaselineAnswersUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfFloat)
}
func (u *XAnalysisMarketSignalsParamsBaselineAnswersUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u, nil)
}

// The earlier row's monitor, with its fingerprint.
type XAnalysisMarketSignalsParamsBaselineMonitor struct {
	ConfigurationFingerprint param.Opt[string] `json:"configurationFingerprint,omitzero"`
	paramObj
}

func (r XAnalysisMarketSignalsParamsBaselineMonitor) MarshalJSON() (data []byte, err error) {
	type shadow XAnalysisMarketSignalsParamsBaselineMonitor
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *XAnalysisMarketSignalsParamsBaselineMonitor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, nil)
}

// The property ID is required.
type XAnalysisMarketSignalsParamsBaselineTweet struct {
	ID string `json:"id" api:"required"`
	paramObj
}

func (r XAnalysisMarketSignalsParamsBaselineTweet) MarshalJSON() (data []byte, err error) {
	type shadow XAnalysisMarketSignalsParamsBaselineTweet
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *XAnalysisMarketSignalsParamsBaselineTweet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, nil)
}

// Search order. Default Latest.
type XAnalysisMarketSignalsParamsQueryType string

const (
	XAnalysisMarketSignalsParamsQueryTypeLatest XAnalysisMarketSignalsParamsQueryType = "Latest"
	XAnalysisMarketSignalsParamsQueryTypeTop    XAnalysisMarketSignalsParamsQueryType = "Top"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type XAnalysisMarketSignalsParamsSinceTimeUnion struct {
	OfString param.Opt[string] `json:",omitzero,inline"`
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	paramUnion
}

func (u XAnalysisMarketSignalsParamsSinceTimeUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfInt)
}
func (u *XAnalysisMarketSignalsParamsSinceTimeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u, nil)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type XAnalysisMarketSignalsParamsUntilTimeUnion struct {
	OfString param.Opt[string] `json:",omitzero,inline"`
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	paramUnion
}

func (u XAnalysisMarketSignalsParamsUntilTimeUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfInt)
}
func (u *XAnalysisMarketSignalsParamsUntilTimeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u, nil)
}

type XAnalysisNewsParams struct {
	// The next_cursor of the page before. Send the same search with it.
	Cursor param.Opt[string] `json:"cursor,omitzero"`
	// Maximum posts to analyze from the search. A low credit balance can return fewer.
	Limit param.Opt[int64] `json:"limit,omitzero"`
	// Analyze a public List's posts. Send its ID or URL.
	ListID param.Opt[string] `json:"listId,omitzero"`
	// X search to analyze, with the operators of /x/tweets/search, such as from:nasa
	// or lang:en. X searches up to 512 characters, counted with the other search
	// fields, as /x/tweets/search counts q.
	Query param.Opt[string] `json:"query,omitzero"`
	// Analyze the quotes of 1 post. Send its ID or URL.
	QuotesOf param.Opt[string] `json:"quotesOf,omitzero"`
	// Analyze the replies in 1 post's conversation. Send its ID or URL.
	RepliesTo param.Opt[string] `json:"repliesTo,omitzero"`
	// Analyze 1 account's posts. Send a handle, with or without @, or a profile URL.
	Username param.Opt[string] `json:"username,omitzero"`
	// Analysis settings, the same as the analysis input of Xquik's AI Actors. Leave it
	// out for the route's own questions.
	Analysis XAnalysisNewsParamsAnalysis `json:"analysis,omitzero"`
	// Rows of an earlier answer to compare with, such as its results. Each needs the
	// post's id, or tweet.id, & its answers. Send monitor too, so other settings show
	// as not_comparable. Each result's monitor says whether the post is new, changed
	// or unchanged. The comparison costs nothing extra.
	Baseline []XAnalysisNewsParamsBaseline `json:"baseline,omitzero"`
	// Search order. Default Latest.
	//
	// Any of "Latest", "Top".
	QueryType XAnalysisNewsParamsQueryType `json:"queryType,omitzero"`
	// Inclusive start of the search, such as 2026-09-25T00:00:00Z. Unix seconds also
	// work. A time without an offset is UTC.
	SinceTime XAnalysisNewsParamsSinceTimeUnion `json:"sinceTime,omitzero"`
	// Your own texts to analyze, such as drafts. Reads nothing from X. Each text costs
	// the same as 1 analyzed post.
	Texts []string `json:"texts,omitzero"`
	// Post IDs or post URLs to analyze, up to 100.
	TweetIDs []string `json:"tweetIds,omitzero"`
	// Exclusive end of the search, such as 2026-09-26T00:00:00Z. Unix seconds also
	// work. A time without an offset is UTC.
	UntilTime XAnalysisNewsParamsUntilTimeUnion `json:"untilTime,omitzero"`
	paramObj
}

func (r XAnalysisNewsParams) MarshalJSON() (data []byte, err error) {
	type shadow XAnalysisNewsParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *XAnalysisNewsParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, nil)
}

// Analysis settings, the same as the analysis input of Xquik's AI Actors. Leave it
// out for the route's own questions.
type XAnalysisNewsParamsAnalysis struct {
	// Posts analyzed at the same time.
	Concurrency param.Opt[int64] `json:"concurrency,omitzero"`
	// Background the AI reads with every post.
	Context param.Opt[string] `json:"context,omitzero"`
	// Bytes of post, quote & settings the AI reads. Xquik cuts longer posts to fit.
	// Empty settings take 219 bytes, so a smaller value gets a 400.
	MaxContextBytes param.Opt[int64] `json:"maxContextBytes,omitzero"`
	// A ready set of questions, used when questions is left out. complaints,
	// competitors, purchase_intent & product_feedback need targets. viral_lite scores
	// how viral the wording reads.
	//
	// Any of "brand", "complaints", "competitors", "purchase_intent",
	// "product_feedback", "news", "sentiment", "market", "viral_lite".
	Preset string `json:"preset,omitzero"`
	// Your own questions, which replace the route's questions. Write question IDs &
	// category names in lower case without hyphens.
	Questions []XAnalysisNewsParamsAnalysisQuestion `json:"questions,omitzero"`
	// Brands, people, assets or topics the questions are about, each with its aliases.
	Targets []XAnalysisNewsParamsAnalysisTarget `json:"targets,omitzero"`
	paramObj
}

func (r XAnalysisNewsParamsAnalysis) MarshalJSON() (data []byte, err error) {
	type shadow XAnalysisNewsParamsAnalysis
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *XAnalysisNewsParamsAnalysis) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, nil)
}

func init() {
	apijson.RegisterFieldValidator[XAnalysisNewsParamsAnalysis](
		"preset", "brand", "complaints", "competitors", "purchase_intent", "product_feedback", "news", "sentiment", "market", "viral_lite",
	)
}

// A question the AI answers for each post. A choice picks 1 of 2 to 255
// categories, a score picks 1 of 2 or more ordered levels & a probability answers
// yes or no.
//
// The properties ID, Instructions, Type, Version are required.
type XAnalysisNewsParamsAnalysisQuestion struct {
	// Name of the answer, such as request. A name the response renames or drops gets a
	// 400, such as created_at, favorite_count, muting or blocking.
	ID string `json:"id" api:"required"`
	// The question, as the AI reads it.
	Instructions string `json:"instructions" api:"required"`
	// Any of "choice", "score", "probability".
	Type string `json:"type,omitzero" api:"required"`
	// Your version of the question, such as 1.
	Version string `json:"version" api:"required"`
	// For a choice. Each category name with its description, or null.
	Categories map[string]param.Opt[string] `json:"categories,omitzero"`
	// For a probability. What counts as yes & as no.
	Criteria XAnalysisNewsParamsAnalysisQuestionCriteria `json:"criteria,omitzero"`
	// For a score. Each level's description, lowest first.
	Levels []string `json:"levels,omitzero"`
	paramObj
}

func (r XAnalysisNewsParamsAnalysisQuestion) MarshalJSON() (data []byte, err error) {
	type shadow XAnalysisNewsParamsAnalysisQuestion
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *XAnalysisNewsParamsAnalysisQuestion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, nil)
}

func init() {
	apijson.RegisterFieldValidator[XAnalysisNewsParamsAnalysisQuestion](
		"type", "choice", "score", "probability",
	)
}

// For a probability. What counts as yes & as no.
//
// The properties No, Yes are required.
type XAnalysisNewsParamsAnalysisQuestionCriteria struct {
	No  string `json:"no" api:"required"`
	Yes string `json:"yes" api:"required"`
	paramObj
}

func (r XAnalysisNewsParamsAnalysisQuestionCriteria) MarshalJSON() (data []byte, err error) {
	type shadow XAnalysisNewsParamsAnalysisQuestionCriteria
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *XAnalysisNewsParamsAnalysisQuestionCriteria) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, nil)
}

// The property Name is required.
type XAnalysisNewsParamsAnalysisTarget struct {
	Name    string   `json:"name" api:"required"`
	Aliases []string `json:"aliases,omitzero"`
	paramObj
}

func (r XAnalysisNewsParamsAnalysisTarget) MarshalJSON() (data []byte, err error) {
	type shadow XAnalysisNewsParamsAnalysisTarget
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *XAnalysisNewsParamsAnalysisTarget) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, nil)
}

// The property Answers is required.
type XAnalysisNewsParamsBaseline struct {
	// The earlier answers by question ID.
	Answers map[string]XAnalysisNewsParamsBaselineAnswersUnion `json:"answers,omitzero" api:"required"`
	// Post ID, or text:1 & so on. Otherwise tweet.id counts.
	ID param.Opt[string] `json:"id,omitzero"`
	// The earlier row's monitor, with its fingerprint.
	Monitor XAnalysisNewsParamsBaselineMonitor `json:"monitor,omitzero"`
	Tweet   XAnalysisNewsParamsBaselineTweet   `json:"tweet,omitzero"`
	paramObj
}

func (r XAnalysisNewsParamsBaseline) MarshalJSON() (data []byte, err error) {
	type shadow XAnalysisNewsParamsBaseline
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *XAnalysisNewsParamsBaseline) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, nil)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type XAnalysisNewsParamsBaselineAnswersUnion struct {
	OfString param.Opt[string]  `json:",omitzero,inline"`
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	paramUnion
}

func (u XAnalysisNewsParamsBaselineAnswersUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfFloat)
}
func (u *XAnalysisNewsParamsBaselineAnswersUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u, nil)
}

// The earlier row's monitor, with its fingerprint.
type XAnalysisNewsParamsBaselineMonitor struct {
	ConfigurationFingerprint param.Opt[string] `json:"configurationFingerprint,omitzero"`
	paramObj
}

func (r XAnalysisNewsParamsBaselineMonitor) MarshalJSON() (data []byte, err error) {
	type shadow XAnalysisNewsParamsBaselineMonitor
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *XAnalysisNewsParamsBaselineMonitor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, nil)
}

// The property ID is required.
type XAnalysisNewsParamsBaselineTweet struct {
	ID string `json:"id" api:"required"`
	paramObj
}

func (r XAnalysisNewsParamsBaselineTweet) MarshalJSON() (data []byte, err error) {
	type shadow XAnalysisNewsParamsBaselineTweet
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *XAnalysisNewsParamsBaselineTweet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, nil)
}

// Search order. Default Latest.
type XAnalysisNewsParamsQueryType string

const (
	XAnalysisNewsParamsQueryTypeLatest XAnalysisNewsParamsQueryType = "Latest"
	XAnalysisNewsParamsQueryTypeTop    XAnalysisNewsParamsQueryType = "Top"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type XAnalysisNewsParamsSinceTimeUnion struct {
	OfString param.Opt[string] `json:",omitzero,inline"`
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	paramUnion
}

func (u XAnalysisNewsParamsSinceTimeUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfInt)
}
func (u *XAnalysisNewsParamsSinceTimeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u, nil)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type XAnalysisNewsParamsUntilTimeUnion struct {
	OfString param.Opt[string] `json:",omitzero,inline"`
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	paramUnion
}

func (u XAnalysisNewsParamsUntilTimeUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfInt)
}
func (u *XAnalysisNewsParamsUntilTimeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u, nil)
}

type XAnalysisSentimentParams struct {
	// The next_cursor of the page before. Send the same search with it.
	Cursor param.Opt[string] `json:"cursor,omitzero"`
	// Maximum posts to analyze from the search. A low credit balance can return fewer.
	Limit param.Opt[int64] `json:"limit,omitzero"`
	// Analyze a public List's posts. Send its ID or URL.
	ListID param.Opt[string] `json:"listId,omitzero"`
	// X search to analyze, with the operators of /x/tweets/search, such as from:nasa
	// or lang:en. X searches up to 512 characters, counted with the other search
	// fields, as /x/tweets/search counts q.
	Query param.Opt[string] `json:"query,omitzero"`
	// Analyze the quotes of 1 post. Send its ID or URL.
	QuotesOf param.Opt[string] `json:"quotesOf,omitzero"`
	// Analyze the replies in 1 post's conversation. Send its ID or URL.
	RepliesTo param.Opt[string] `json:"repliesTo,omitzero"`
	// Analyze 1 account's posts. Send a handle, with or without @, or a profile URL.
	Username param.Opt[string] `json:"username,omitzero"`
	// Analysis settings, the same as the analysis input of Xquik's AI Actors. Leave it
	// out for the route's own questions.
	Analysis XAnalysisSentimentParamsAnalysis `json:"analysis,omitzero"`
	// Rows of an earlier answer to compare with, such as its results. Each needs the
	// post's id, or tweet.id, & its answers. Send monitor too, so other settings show
	// as not_comparable. Each result's monitor says whether the post is new, changed
	// or unchanged. The comparison costs nothing extra.
	Baseline []XAnalysisSentimentParamsBaseline `json:"baseline,omitzero"`
	// Search order. Default Latest.
	//
	// Any of "Latest", "Top".
	QueryType XAnalysisSentimentParamsQueryType `json:"queryType,omitzero"`
	// Inclusive start of the search, such as 2026-09-25T00:00:00Z. Unix seconds also
	// work. A time without an offset is UTC.
	SinceTime XAnalysisSentimentParamsSinceTimeUnion `json:"sinceTime,omitzero"`
	// Your own texts to analyze, such as drafts. Reads nothing from X. Each text costs
	// the same as 1 analyzed post.
	Texts []string `json:"texts,omitzero"`
	// Post IDs or post URLs to analyze, up to 100.
	TweetIDs []string `json:"tweetIds,omitzero"`
	// Exclusive end of the search, such as 2026-09-26T00:00:00Z. Unix seconds also
	// work. A time without an offset is UTC.
	UntilTime XAnalysisSentimentParamsUntilTimeUnion `json:"untilTime,omitzero"`
	paramObj
}

func (r XAnalysisSentimentParams) MarshalJSON() (data []byte, err error) {
	type shadow XAnalysisSentimentParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *XAnalysisSentimentParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, nil)
}

// Analysis settings, the same as the analysis input of Xquik's AI Actors. Leave it
// out for the route's own questions.
type XAnalysisSentimentParamsAnalysis struct {
	// Posts analyzed at the same time.
	Concurrency param.Opt[int64] `json:"concurrency,omitzero"`
	// Background the AI reads with every post.
	Context param.Opt[string] `json:"context,omitzero"`
	// Bytes of post, quote & settings the AI reads. Xquik cuts longer posts to fit.
	// Empty settings take 219 bytes, so a smaller value gets a 400.
	MaxContextBytes param.Opt[int64] `json:"maxContextBytes,omitzero"`
	// A ready set of questions, used when questions is left out. complaints,
	// competitors, purchase_intent & product_feedback need targets. viral_lite scores
	// how viral the wording reads.
	//
	// Any of "brand", "complaints", "competitors", "purchase_intent",
	// "product_feedback", "news", "sentiment", "market", "viral_lite".
	Preset string `json:"preset,omitzero"`
	// Your own questions, which replace the route's questions. Write question IDs &
	// category names in lower case without hyphens.
	Questions []XAnalysisSentimentParamsAnalysisQuestion `json:"questions,omitzero"`
	// Brands, people, assets or topics the questions are about, each with its aliases.
	Targets []XAnalysisSentimentParamsAnalysisTarget `json:"targets,omitzero"`
	paramObj
}

func (r XAnalysisSentimentParamsAnalysis) MarshalJSON() (data []byte, err error) {
	type shadow XAnalysisSentimentParamsAnalysis
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *XAnalysisSentimentParamsAnalysis) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, nil)
}

func init() {
	apijson.RegisterFieldValidator[XAnalysisSentimentParamsAnalysis](
		"preset", "brand", "complaints", "competitors", "purchase_intent", "product_feedback", "news", "sentiment", "market", "viral_lite",
	)
}

// A question the AI answers for each post. A choice picks 1 of 2 to 255
// categories, a score picks 1 of 2 or more ordered levels & a probability answers
// yes or no.
//
// The properties ID, Instructions, Type, Version are required.
type XAnalysisSentimentParamsAnalysisQuestion struct {
	// Name of the answer, such as request. A name the response renames or drops gets a
	// 400, such as created_at, favorite_count, muting or blocking.
	ID string `json:"id" api:"required"`
	// The question, as the AI reads it.
	Instructions string `json:"instructions" api:"required"`
	// Any of "choice", "score", "probability".
	Type string `json:"type,omitzero" api:"required"`
	// Your version of the question, such as 1.
	Version string `json:"version" api:"required"`
	// For a choice. Each category name with its description, or null.
	Categories map[string]param.Opt[string] `json:"categories,omitzero"`
	// For a probability. What counts as yes & as no.
	Criteria XAnalysisSentimentParamsAnalysisQuestionCriteria `json:"criteria,omitzero"`
	// For a score. Each level's description, lowest first.
	Levels []string `json:"levels,omitzero"`
	paramObj
}

func (r XAnalysisSentimentParamsAnalysisQuestion) MarshalJSON() (data []byte, err error) {
	type shadow XAnalysisSentimentParamsAnalysisQuestion
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *XAnalysisSentimentParamsAnalysisQuestion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, nil)
}

func init() {
	apijson.RegisterFieldValidator[XAnalysisSentimentParamsAnalysisQuestion](
		"type", "choice", "score", "probability",
	)
}

// For a probability. What counts as yes & as no.
//
// The properties No, Yes are required.
type XAnalysisSentimentParamsAnalysisQuestionCriteria struct {
	No  string `json:"no" api:"required"`
	Yes string `json:"yes" api:"required"`
	paramObj
}

func (r XAnalysisSentimentParamsAnalysisQuestionCriteria) MarshalJSON() (data []byte, err error) {
	type shadow XAnalysisSentimentParamsAnalysisQuestionCriteria
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *XAnalysisSentimentParamsAnalysisQuestionCriteria) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, nil)
}

// The property Name is required.
type XAnalysisSentimentParamsAnalysisTarget struct {
	Name    string   `json:"name" api:"required"`
	Aliases []string `json:"aliases,omitzero"`
	paramObj
}

func (r XAnalysisSentimentParamsAnalysisTarget) MarshalJSON() (data []byte, err error) {
	type shadow XAnalysisSentimentParamsAnalysisTarget
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *XAnalysisSentimentParamsAnalysisTarget) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, nil)
}

// The property Answers is required.
type XAnalysisSentimentParamsBaseline struct {
	// The earlier answers by question ID.
	Answers map[string]XAnalysisSentimentParamsBaselineAnswersUnion `json:"answers,omitzero" api:"required"`
	// Post ID, or text:1 & so on. Otherwise tweet.id counts.
	ID param.Opt[string] `json:"id,omitzero"`
	// The earlier row's monitor, with its fingerprint.
	Monitor XAnalysisSentimentParamsBaselineMonitor `json:"monitor,omitzero"`
	Tweet   XAnalysisSentimentParamsBaselineTweet   `json:"tweet,omitzero"`
	paramObj
}

func (r XAnalysisSentimentParamsBaseline) MarshalJSON() (data []byte, err error) {
	type shadow XAnalysisSentimentParamsBaseline
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *XAnalysisSentimentParamsBaseline) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, nil)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type XAnalysisSentimentParamsBaselineAnswersUnion struct {
	OfString param.Opt[string]  `json:",omitzero,inline"`
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	paramUnion
}

func (u XAnalysisSentimentParamsBaselineAnswersUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfFloat)
}
func (u *XAnalysisSentimentParamsBaselineAnswersUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u, nil)
}

// The earlier row's monitor, with its fingerprint.
type XAnalysisSentimentParamsBaselineMonitor struct {
	ConfigurationFingerprint param.Opt[string] `json:"configurationFingerprint,omitzero"`
	paramObj
}

func (r XAnalysisSentimentParamsBaselineMonitor) MarshalJSON() (data []byte, err error) {
	type shadow XAnalysisSentimentParamsBaselineMonitor
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *XAnalysisSentimentParamsBaselineMonitor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, nil)
}

// The property ID is required.
type XAnalysisSentimentParamsBaselineTweet struct {
	ID string `json:"id" api:"required"`
	paramObj
}

func (r XAnalysisSentimentParamsBaselineTweet) MarshalJSON() (data []byte, err error) {
	type shadow XAnalysisSentimentParamsBaselineTweet
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *XAnalysisSentimentParamsBaselineTweet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, nil)
}

// Search order. Default Latest.
type XAnalysisSentimentParamsQueryType string

const (
	XAnalysisSentimentParamsQueryTypeLatest XAnalysisSentimentParamsQueryType = "Latest"
	XAnalysisSentimentParamsQueryTypeTop    XAnalysisSentimentParamsQueryType = "Top"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type XAnalysisSentimentParamsSinceTimeUnion struct {
	OfString param.Opt[string] `json:",omitzero,inline"`
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	paramUnion
}

func (u XAnalysisSentimentParamsSinceTimeUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfInt)
}
func (u *XAnalysisSentimentParamsSinceTimeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u, nil)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type XAnalysisSentimentParamsUntilTimeUnion struct {
	OfString param.Opt[string] `json:",omitzero,inline"`
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	paramUnion
}

func (u XAnalysisSentimentParamsUntilTimeUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfInt)
}
func (u *XAnalysisSentimentParamsUntilTimeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u, nil)
}

type XAnalysisViralScoreParams struct {
	// The next_cursor of the page before. Send the same search with it.
	Cursor param.Opt[string] `json:"cursor,omitzero"`
	// Maximum posts to analyze from the search. A low credit balance can return fewer.
	Limit param.Opt[int64] `json:"limit,omitzero"`
	// Analyze a public List's posts. Send its ID or URL.
	ListID param.Opt[string] `json:"listId,omitzero"`
	// X search to analyze, with the operators of /x/tweets/search, such as from:nasa
	// or lang:en. X searches up to 512 characters, counted with the other search
	// fields, as /x/tweets/search counts q.
	Query param.Opt[string] `json:"query,omitzero"`
	// Analyze the quotes of 1 post. Send its ID or URL.
	QuotesOf param.Opt[string] `json:"quotesOf,omitzero"`
	// Analyze the replies in 1 post's conversation. Send its ID or URL.
	RepliesTo param.Opt[string] `json:"repliesTo,omitzero"`
	// Analyze 1 account's posts. Send a handle, with or without @, or a profile URL.
	Username param.Opt[string] `json:"username,omitzero"`
	// Analysis settings, the same as the analysis input of Xquik's AI Actors. Leave it
	// out for the route's own questions.
	Analysis XAnalysisViralScoreParamsAnalysis `json:"analysis,omitzero"`
	// Rows of an earlier answer to compare with, such as its results. Each needs the
	// post's id, or tweet.id, & its answers. Send monitor too, so other settings show
	// as not_comparable. Each result's monitor says whether the post is new, changed
	// or unchanged. The comparison costs nothing extra.
	Baseline []XAnalysisViralScoreParamsBaseline `json:"baseline,omitzero"`
	// Search order. Default Latest.
	//
	// Any of "Latest", "Top".
	QueryType XAnalysisViralScoreParamsQueryType `json:"queryType,omitzero"`
	// Inclusive start of the search, such as 2026-09-25T00:00:00Z. Unix seconds also
	// work. A time without an offset is UTC.
	SinceTime XAnalysisViralScoreParamsSinceTimeUnion `json:"sinceTime,omitzero"`
	// Your own texts to analyze, such as drafts. Reads nothing from X. Each text costs
	// the same as 1 analyzed post.
	Texts []string `json:"texts,omitzero"`
	// Post IDs or post URLs to analyze, up to 100.
	TweetIDs []string `json:"tweetIds,omitzero"`
	// Exclusive end of the search, such as 2026-09-26T00:00:00Z. Unix seconds also
	// work. A time without an offset is UTC.
	UntilTime XAnalysisViralScoreParamsUntilTimeUnion `json:"untilTime,omitzero"`
	paramObj
}

func (r XAnalysisViralScoreParams) MarshalJSON() (data []byte, err error) {
	type shadow XAnalysisViralScoreParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *XAnalysisViralScoreParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, nil)
}

// Analysis settings, the same as the analysis input of Xquik's AI Actors. Leave it
// out for the route's own questions.
type XAnalysisViralScoreParamsAnalysis struct {
	// Posts analyzed at the same time.
	Concurrency param.Opt[int64] `json:"concurrency,omitzero"`
	// Background the AI reads with every post.
	Context param.Opt[string] `json:"context,omitzero"`
	// Bytes of post, quote & settings the AI reads. Xquik cuts longer posts to fit.
	// Empty settings take 219 bytes, so a smaller value gets a 400.
	MaxContextBytes param.Opt[int64] `json:"maxContextBytes,omitzero"`
	// A ready set of questions, used when questions is left out. complaints,
	// competitors, purchase_intent & product_feedback need targets. viral_lite scores
	// how viral the wording reads.
	//
	// Any of "brand", "complaints", "competitors", "purchase_intent",
	// "product_feedback", "news", "sentiment", "market", "viral_lite".
	Preset string `json:"preset,omitzero"`
	// Your own questions, which replace the route's questions. Write question IDs &
	// category names in lower case without hyphens.
	Questions []XAnalysisViralScoreParamsAnalysisQuestion `json:"questions,omitzero"`
	// Brands, people, assets or topics the questions are about, each with its aliases.
	Targets []XAnalysisViralScoreParamsAnalysisTarget `json:"targets,omitzero"`
	paramObj
}

func (r XAnalysisViralScoreParamsAnalysis) MarshalJSON() (data []byte, err error) {
	type shadow XAnalysisViralScoreParamsAnalysis
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *XAnalysisViralScoreParamsAnalysis) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, nil)
}

func init() {
	apijson.RegisterFieldValidator[XAnalysisViralScoreParamsAnalysis](
		"preset", "brand", "complaints", "competitors", "purchase_intent", "product_feedback", "news", "sentiment", "market", "viral_lite",
	)
}

// A question the AI answers for each post. A choice picks 1 of 2 to 255
// categories, a score picks 1 of 2 or more ordered levels & a probability answers
// yes or no.
//
// The properties ID, Instructions, Type, Version are required.
type XAnalysisViralScoreParamsAnalysisQuestion struct {
	// Name of the answer, such as request. A name the response renames or drops gets a
	// 400, such as created_at, favorite_count, muting or blocking.
	ID string `json:"id" api:"required"`
	// The question, as the AI reads it.
	Instructions string `json:"instructions" api:"required"`
	// Any of "choice", "score", "probability".
	Type string `json:"type,omitzero" api:"required"`
	// Your version of the question, such as 1.
	Version string `json:"version" api:"required"`
	// For a choice. Each category name with its description, or null.
	Categories map[string]param.Opt[string] `json:"categories,omitzero"`
	// For a probability. What counts as yes & as no.
	Criteria XAnalysisViralScoreParamsAnalysisQuestionCriteria `json:"criteria,omitzero"`
	// For a score. Each level's description, lowest first.
	Levels []string `json:"levels,omitzero"`
	paramObj
}

func (r XAnalysisViralScoreParamsAnalysisQuestion) MarshalJSON() (data []byte, err error) {
	type shadow XAnalysisViralScoreParamsAnalysisQuestion
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *XAnalysisViralScoreParamsAnalysisQuestion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, nil)
}

func init() {
	apijson.RegisterFieldValidator[XAnalysisViralScoreParamsAnalysisQuestion](
		"type", "choice", "score", "probability",
	)
}

// For a probability. What counts as yes & as no.
//
// The properties No, Yes are required.
type XAnalysisViralScoreParamsAnalysisQuestionCriteria struct {
	No  string `json:"no" api:"required"`
	Yes string `json:"yes" api:"required"`
	paramObj
}

func (r XAnalysisViralScoreParamsAnalysisQuestionCriteria) MarshalJSON() (data []byte, err error) {
	type shadow XAnalysisViralScoreParamsAnalysisQuestionCriteria
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *XAnalysisViralScoreParamsAnalysisQuestionCriteria) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, nil)
}

// The property Name is required.
type XAnalysisViralScoreParamsAnalysisTarget struct {
	Name    string   `json:"name" api:"required"`
	Aliases []string `json:"aliases,omitzero"`
	paramObj
}

func (r XAnalysisViralScoreParamsAnalysisTarget) MarshalJSON() (data []byte, err error) {
	type shadow XAnalysisViralScoreParamsAnalysisTarget
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *XAnalysisViralScoreParamsAnalysisTarget) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, nil)
}

// The property Answers is required.
type XAnalysisViralScoreParamsBaseline struct {
	// The earlier answers by question ID.
	Answers map[string]XAnalysisViralScoreParamsBaselineAnswersUnion `json:"answers,omitzero" api:"required"`
	// Post ID, or text:1 & so on. Otherwise tweet.id counts.
	ID param.Opt[string] `json:"id,omitzero"`
	// The earlier row's monitor, with its fingerprint.
	Monitor XAnalysisViralScoreParamsBaselineMonitor `json:"monitor,omitzero"`
	Tweet   XAnalysisViralScoreParamsBaselineTweet   `json:"tweet,omitzero"`
	paramObj
}

func (r XAnalysisViralScoreParamsBaseline) MarshalJSON() (data []byte, err error) {
	type shadow XAnalysisViralScoreParamsBaseline
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *XAnalysisViralScoreParamsBaseline) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, nil)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type XAnalysisViralScoreParamsBaselineAnswersUnion struct {
	OfString param.Opt[string]  `json:",omitzero,inline"`
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	paramUnion
}

func (u XAnalysisViralScoreParamsBaselineAnswersUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfFloat)
}
func (u *XAnalysisViralScoreParamsBaselineAnswersUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u, nil)
}

// The earlier row's monitor, with its fingerprint.
type XAnalysisViralScoreParamsBaselineMonitor struct {
	ConfigurationFingerprint param.Opt[string] `json:"configurationFingerprint,omitzero"`
	paramObj
}

func (r XAnalysisViralScoreParamsBaselineMonitor) MarshalJSON() (data []byte, err error) {
	type shadow XAnalysisViralScoreParamsBaselineMonitor
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *XAnalysisViralScoreParamsBaselineMonitor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, nil)
}

// The property ID is required.
type XAnalysisViralScoreParamsBaselineTweet struct {
	ID string `json:"id" api:"required"`
	paramObj
}

func (r XAnalysisViralScoreParamsBaselineTweet) MarshalJSON() (data []byte, err error) {
	type shadow XAnalysisViralScoreParamsBaselineTweet
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *XAnalysisViralScoreParamsBaselineTweet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r, nil)
}

// Search order. Default Latest.
type XAnalysisViralScoreParamsQueryType string

const (
	XAnalysisViralScoreParamsQueryTypeLatest XAnalysisViralScoreParamsQueryType = "Latest"
	XAnalysisViralScoreParamsQueryTypeTop    XAnalysisViralScoreParamsQueryType = "Top"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type XAnalysisViralScoreParamsSinceTimeUnion struct {
	OfString param.Opt[string] `json:",omitzero,inline"`
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	paramUnion
}

func (u XAnalysisViralScoreParamsSinceTimeUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfInt)
}
func (u *XAnalysisViralScoreParamsSinceTimeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u, nil)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type XAnalysisViralScoreParamsUntilTimeUnion struct {
	OfString param.Opt[string] `json:",omitzero,inline"`
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	paramUnion
}

func (u XAnalysisViralScoreParamsUntilTimeUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfInt)
}
func (u *XAnalysisViralScoreParamsUntilTimeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u, nil)
}
