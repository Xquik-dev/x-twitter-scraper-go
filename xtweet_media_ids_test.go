// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

package xtwitterscraper_test

import (
	"context"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/Xquik-dev/x-twitter-scraper-go"
	"github.com/Xquik-dev/x-twitter-scraper-go/option"
)

func TestXTweetNewSendsMediaIDs(t *testing.T) {
	request := &capturedRequest{}
	client := xtwitterscraper.NewClient(
		option.WithBaseURL("https://example.com/api/v1/"),
		option.WithMaxRetries(0),
		option.WithHTTPClient(testHTTPDoer(func(req *http.Request) (*http.Response, error) {
			body, err := io.ReadAll(req.Body)
			if err != nil {
				return nil, err
			}
			*request = capturedRequest{method: req.Method, path: req.URL.Path, query: req.URL.Query(), body: body}
			return &http.Response{
				StatusCode: http.StatusOK,
				Header:     http.Header{"Content-Type": {"application/json"}},
				Body:       io.NopCloser(strings.NewReader(`{"id":"12345","status":"accepted"}`)),
			}, nil
		})),
	)
	_, err := client.X.Tweets.New(context.Background(), xtwitterscraper.XTweetNewParams{
		Account:        "@elonmusk",
		IdempotencyKey: "write-1",
		Text:           xtwitterscraper.String("Just launched our new feature!"),
		MediaIDs:       []string{"2104371968258560001"},
	})
	if err != nil {
		t.Fatal(err)
	}
	request.expect(t, http.MethodPost, "/api/v1/x/tweets", "")
	request.expectBody(t, `{"account":"@elonmusk","text":"Just launched our new feature!","media_ids":["2104371968258560001"]}`)
}
