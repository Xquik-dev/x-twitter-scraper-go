// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

package xtwitterscraper_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/Xquik-dev/x-twitter-scraper-go"
	"github.com/Xquik-dev/x-twitter-scraper-go/packages/param"
)

type filteredRead func(client xtwitterscraper.Client, topic, imageShows param.Opt[string]) (page xtwitterscraper.PaginatedTweets, decoded any, err error)

var filteredReads = map[string]struct {
	path string
	read filteredRead
}{
	"X.Tweets.Search": {"/api/v1/x/tweets/search", func(client xtwitterscraper.Client, topic, imageShows param.Opt[string]) (xtwitterscraper.PaginatedTweets, any, error) {
		res, err := client.X.Tweets.Search(context.Background(), xtwitterscraper.XTweetSearchParams{Topic: topic, ImageShows: imageShows})
		if err != nil {
			return xtwitterscraper.PaginatedTweets{}, nil, err
		}
		return xtwitterscraper.PaginatedTweets{Tweets: res.Tweets, ImageFilter: res.ImageFilter, TopicFilter: res.TopicFilter}, res.AsPaginatedTweets(), nil
	}},
	"X.Users.GetTweets": {"/api/v1/x/users/nasa/tweets", func(client xtwitterscraper.Client, topic, imageShows param.Opt[string]) (xtwitterscraper.PaginatedTweets, any, error) {
		res, err := client.X.Users.GetTweets(context.Background(), "nasa", xtwitterscraper.XUserGetTweetsParams{Topic: topic, ImageShows: imageShows})
		if err != nil {
			return xtwitterscraper.PaginatedTweets{}, nil, err
		}
		return *res, res, nil
	}},
	"X.Users.GetReplies": {"/api/v1/x/users/nasa/replies", func(client xtwitterscraper.Client, topic, imageShows param.Opt[string]) (xtwitterscraper.PaginatedTweets, any, error) {
		res, err := client.X.Users.GetReplies(context.Background(), "nasa", xtwitterscraper.XUserGetRepliesParams{Topic: topic, ImageShows: imageShows})
		if err != nil {
			return xtwitterscraper.PaginatedTweets{}, nil, err
		}
		return *res, res, nil
	}},
	"X.Users.GetMedia": {"/api/v1/x/users/nasa/media", func(client xtwitterscraper.Client, topic, imageShows param.Opt[string]) (xtwitterscraper.PaginatedTweets, any, error) {
		res, err := client.X.Users.GetMedia(context.Background(), "nasa", xtwitterscraper.XUserGetMediaParams{Topic: topic, ImageShows: imageShows})
		if err != nil {
			return xtwitterscraper.PaginatedTweets{}, nil, err
		}
		return *res, res, nil
	}},
}

const chartQuestion = "Does a picture attached to this post show a chart or graph that plots data, such as a line, bar, pie or candlestick chart? Judge the pictures in tweet.images."

func TestEachPostReadSendsTopicAndDecodesTheTopicFilter(t *testing.T) {
	for name, test := range filteredReads {
		t.Run(name, func(t *testing.T) {
			client, request := fixtureClient(t, "topicFilteredTweets.json")
			page, decoded, err := test.read(client, xtwitterscraper.String("the Rust programming language"), param.Opt[string]{})
			if err != nil {
				t.Fatal(err)
			}
			request.expect(t, http.MethodGet, test.path, "topic=the+Rust+programming+language")
			expectDecoded(t, decoded)
			filter := page.TopicFilter
			if filter.Topic != "the Rust programming language" ||
				filter.Question != "Is this post about the Rust programming language?" ||
				filter.Cutoff != 0.5 || filter.Checked != 20 || filter.Kept != 1 ||
				len(filter.Unchecked) != 1 || filter.Unchecked[0].ID != "1234567891" || filter.Unchecked[0].Reason != "check_failed" {
				t.Fatalf("topicFilter = %+v", filter)
			}
			if len(page.Tweets) != 1 || page.Tweets[0].TopicScore != 0.93 || page.Tweets[0].ImageScore != 0 || page.ImageFilter.Checked != 0 {
				t.Fatalf("tweets = %+v, imageFilter = %+v", page.Tweets, page.ImageFilter)
			}
		})
	}
}

func TestEachPostReadSendsImageShowsAndDecodesTheImageFilter(t *testing.T) {
	for name, test := range filteredReads {
		t.Run(name, func(t *testing.T) {
			client, request := fixtureClient(t, "imageFilteredTweets.json")
			page, decoded, err := test.read(client, param.Opt[string]{}, xtwitterscraper.String("charts"))
			if err != nil {
				t.Fatal(err)
			}
			request.expect(t, http.MethodGet, test.path, "imageShows=charts")
			expectDecoded(t, decoded)
			filter := page.ImageFilter
			if filter.ImageShows != "charts" || filter.Question != chartQuestion ||
				filter.Cutoff != 0.5 || filter.Checked != 16 || filter.Kept != 4 ||
				len(filter.Unchecked) != 1 || filter.Unchecked[0].ID != "1234567891" || filter.Unchecked[0].Reason != "no_safe_photos" {
				t.Fatalf("imageFilter = %+v", filter)
			}
			if len(page.Tweets) != 1 || page.Tweets[0].ImageScore != 0.88 || page.Tweets[0].TopicScore != 0 || page.TopicFilter.Checked != 0 {
				t.Fatalf("tweets = %+v, topicFilter = %+v", page.Tweets, page.TopicFilter)
			}
		})
	}
}

func TestExtractionsSendTopicAndImageShows(t *testing.T) {
	cases := map[string]struct {
		fixture string
		path    string
		call    func(client xtwitterscraper.Client) error
		want    string
	}{
		"Run with topic": {"runExtraction.json", "/api/v1/extractions", func(client xtwitterscraper.Client) error {
			_, err := client.Extractions.Run(context.Background(), xtwitterscraper.ExtractionRunParams{ToolType: xtwitterscraper.ExtractionRunParamsToolTypeTweetSearchExtractor, SearchQuery: xtwitterscraper.String("rust"), Topic: xtwitterscraper.String("the Rust programming language")})
			return err
		}, `{"toolType":"tweet_search_extractor","searchQuery":"rust","topic":"the Rust programming language"}`},
		"Run with imageShows": {"runExtraction.json", "/api/v1/extractions", func(client xtwitterscraper.Client) error {
			_, err := client.Extractions.Run(context.Background(), xtwitterscraper.ExtractionRunParams{ToolType: xtwitterscraper.ExtractionRunParamsToolTypeUserMedia, TargetUsername: xtwitterscraper.String("nasa"), ImageShows: xtwitterscraper.String("charts")})
			return err
		}, `{"toolType":"user_media","targetUsername":"nasa","imageShows":"charts"}`},
		"EstimateCost with topic": {"estimateExtraction.json", "/api/v1/extractions/estimate", func(client xtwitterscraper.Client) error {
			_, err := client.Extractions.EstimateCost(context.Background(), xtwitterscraper.ExtractionEstimateCostParams{ToolType: xtwitterscraper.ExtractionEstimateCostParamsToolTypeTweetSearchExtractor, SearchQuery: xtwitterscraper.String("rust"), Topic: xtwitterscraper.String("the Rust programming language")})
			return err
		}, `{"toolType":"tweet_search_extractor","searchQuery":"rust","topic":"the Rust programming language"}`},
		"EstimateCost with imageShows": {"estimateExtraction.json", "/api/v1/extractions/estimate", func(client xtwitterscraper.Client) error {
			_, err := client.Extractions.EstimateCost(context.Background(), xtwitterscraper.ExtractionEstimateCostParams{ToolType: xtwitterscraper.ExtractionEstimateCostParamsToolTypeUserMedia, TargetUsername: xtwitterscraper.String("nasa"), ImageShows: xtwitterscraper.String("charts")})
			return err
		}, `{"toolType":"user_media","targetUsername":"nasa","imageShows":"charts"}`},
	}
	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			client, request := fixtureClient(t, test.fixture)
			if err := test.call(client); err != nil {
				t.Fatal(err)
			}
			request.expect(t, http.MethodPost, test.path, "")
			request.expectBody(t, test.want)
		})
	}
}

func TestExtractionJobsDecodeTheTopicAndImageFilters(t *testing.T) {
	client, request := fixtureClient(t, "listFilteredExtractions.json")
	res, err := client.Extractions.List(context.Background(), xtwitterscraper.ExtractionListParams{})
	if err != nil {
		t.Fatal(err)
	}
	request.expect(t, http.MethodGet, "/api/v1/extractions", "")
	expectDecoded(t, res)
	if len(res.Extractions) != 2 {
		t.Fatalf("decoded %s", res.RawJSON())
	}
	topic := res.Extractions[0].TopicFilter
	if topic.Topic != "the Rust programming language" || topic.Checked != 20 || topic.Kept != 1 ||
		!topic.JSON.Unchecked.Valid() || len(topic.Unchecked) != 0 || res.Extractions[0].JSON.ImageFilter.Valid() {
		t.Fatalf("topic job = %s", res.Extractions[0].RawJSON())
	}
	image := res.Extractions[1].ImageFilter
	if image.ImageShows != "charts" || image.Question != chartQuestion || image.Checked != 16 || image.Kept != 4 ||
		len(image.Unchecked) != 1 || image.Unchecked[0].Reason != "no_safe_photos" || res.Extractions[1].JSON.TopicFilter.Valid() {
		t.Fatalf("image job = %s", res.Extractions[1].RawJSON())
	}
}

func TestMediaDownloadSendsImageShowsAndDecodesTheImageFilter(t *testing.T) {
	client, request := fixtureClient(t, "downloadFilteredMedia.json")
	res, err := client.X.Media.Download(context.Background(), xtwitterscraper.XMediaDownloadParams{
		TweetInput: xtwitterscraper.String("1234567890"),
		ImageShows: xtwitterscraper.String("charts"),
	})
	if err != nil {
		t.Fatal(err)
	}
	request.expect(t, http.MethodPost, "/api/v1/x/media/download", "")
	request.expectBody(t, `{"tweetInput":"1234567890","imageShows":"charts"}`)
	expectDecoded(t, res)
	if res.ImageFilter.ImageShows != "charts" || res.ImageFilter.Checked != 1 || res.ImageFilter.Kept != 0 ||
		!res.JSON.Items.Valid() || len(res.Items) != 0 {
		t.Fatalf("decoded %s", res.RawJSON())
	}
}
