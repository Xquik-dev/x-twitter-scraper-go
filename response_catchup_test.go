// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

package xtwitterscraper_test

import (
	"testing"

	"github.com/Xquik-dev/x-twitter-scraper-go"
)

func TestResponsesDecodeTheNewFields(t *testing.T) {
	cases := []struct {
		fixture string
		decode  func([]byte) (interface{ RawJSON() string }, error)
		values  []fieldValue
	}{
		{
			"getAccount.json",
			func(data []byte) (interface{ RawJSON() string }, error) {
				var v xtwitterscraper.AccountGetResponse
				err := v.UnmarshalJSON(data)
				return v, err
			},
			[]fieldValue{
				field("Subscription.Status", "active"),
				field("Subscription.Tier", "pro"),
				field("CreditInfo.AutoTopupStopped", false),
			},
		},
		{
			"updateAccount.json",
			func(data []byte) (interface{ RawJSON() string }, error) {
				var v xtwitterscraper.AccountUpdateLocaleResponse
				err := v.UnmarshalJSON(data)
				return v, err
			},
			nil,
		},
		{
			"getCredits.json",
			func(data []byte) (interface{ RawJSON() string }, error) {
				var v xtwitterscraper.CreditGetBalanceResponse
				err := v.UnmarshalJSON(data)
				return v, err
			},
			[]fieldValue{
				field("AutoTopupStopped", false),
			},
		},
		{
			"createDraft.json",
			func(data []byte) (interface{ RawJSON() string }, error) {
				var v xtwitterscraper.DraftDetail
				err := v.UnmarshalJSON(data)
				return v, err
			},
			nil,
		},
		{
			"getEvent.json",
			func(data []byte) (interface{ RawJSON() string }, error) {
				var v xtwitterscraper.EventDetail
				err := v.UnmarshalJSON(data)
				return v, err
			},
			nil,
		},
		{
			"getExtraction.json",
			func(data []byte) (interface{ RawJSON() string }, error) {
				var v xtwitterscraper.ExtractionGetResponse
				err := v.UnmarshalJSON(data)
				return v, err
			},
			nil,
		},
		{
			"getMonitor.json",
			func(data []byte) (interface{ RawJSON() string }, error) {
				var v xtwitterscraper.Monitor
				err := v.UnmarshalJSON(data)
				return v, err
			},
			nil,
		},
		{
			"listMonitors.json",
			func(data []byte) (interface{ RawJSON() string }, error) {
				var v xtwitterscraper.MonitorListResponse
				err := v.UnmarshalJSON(data)
				return v, err
			},
			[]fieldValue{
				field("HasMore", false),
			},
		},
		{
			"createKeywordMonitor.json",
			func(data []byte) (interface{ RawJSON() string }, error) {
				var v xtwitterscraper.MonitorKeywordNewResponse
				err := v.UnmarshalJSON(data)
				return v, err
			},
			nil,
		},
		{
			"getKeywordMonitor.json",
			func(data []byte) (interface{ RawJSON() string }, error) {
				var v xtwitterscraper.MonitorKeywordGetResponse
				err := v.UnmarshalJSON(data)
				return v, err
			},
			nil,
		},
		{
			"updateKeywordMonitor.json",
			func(data []byte) (interface{ RawJSON() string }, error) {
				var v xtwitterscraper.MonitorKeywordUpdateResponse
				err := v.UnmarshalJSON(data)
				return v, err
			},
			nil,
		},
		{
			"listKeywordMonitors.json",
			func(data []byte) (interface{ RawJSON() string }, error) {
				var v xtwitterscraper.MonitorKeywordListResponse
				err := v.UnmarshalJSON(data)
				return v, err
			},
			[]fieldValue{
				field("HasMore", false),
			},
		},
		{
			"listStyles.json",
			func(data []byte) (interface{ RawJSON() string }, error) {
				var v xtwitterscraper.StyleListResponse
				err := v.UnmarshalJSON(data)
				return v, err
			},
			[]fieldValue{
				field("HasMore", false),
			},
		},
		{
			"listTickets.json",
			func(data []byte) (interface{ RawJSON() string }, error) {
				var v xtwitterscraper.SupportTicketListResponse
				err := v.UnmarshalJSON(data)
				return v, err
			},
			[]fieldValue{
				field("HasMore", false),
			},
		},
		{
			"createWebhook.json",
			func(data []byte) (interface{ RawJSON() string }, error) {
				var v xtwitterscraper.WebhookNewResponse
				err := v.UnmarshalJSON(data)
				return v, err
			},
			nil,
		},
		{
			"updateWebhook.json",
			func(data []byte) (interface{ RawJSON() string }, error) {
				var v xtwitterscraper.Webhook
				err := v.UnmarshalJSON(data)
				return v, err
			},
			nil,
		},
		{
			"listWebhooks.json",
			func(data []byte) (interface{ RawJSON() string }, error) {
				var v xtwitterscraper.WebhookListResponse
				err := v.UnmarshalJSON(data)
				return v, err
			},
			[]fieldValue{
				field("HasMore", false),
			},
		},
		{
			"listWebhookDeliveries.json",
			func(data []byte) (interface{ RawJSON() string }, error) {
				var v xtwitterscraper.WebhookListDeliveriesResponse
				err := v.UnmarshalJSON(data)
				return v, err
			},
			[]fieldValue{
				field("HasMore", false),
			},
		},
		{
			"getArticle.json",
			func(data []byte) (interface{ RawJSON() string }, error) {
				var v xtwitterscraper.XGetArticleResponse
				err := v.UnmarshalJSON(data)
				return v, err
			},
			[]fieldValue{
				field("Article.ArticleID", "2100651802404544512"),
				field("Article.Contents[0].Links[0].Offset", int64(18)),
				field("Article.Contents[0].Links[0].Length", int64(9)),
				field("Article.Contents[0].Links[0].URL", "https://example.com/paragraph"),
				field("Article.RetweetCount", int64(40)),
				field("Article.BookmarkCount", int64(61)),
			},
		},
		{
			"getCommunityInfo.json",
			func(data []byte) (interface{ RawJSON() string }, error) {
				var v xtwitterscraper.XCommunityGetInfoResponse
				err := v.UnmarshalJSON(data)
				return v, err
			},
			nil,
		},
		{
			"checkFollow.json",
			func(data []byte) (interface{ RawJSON() string }, error) {
				var v xtwitterscraper.XFollowerCheckResponse
				err := v.UnmarshalJSON(data)
				return v, err
			},
			[]fieldValue{
				field("SourceID", "11348282"),
				field("TargetID", "34743251"),
			},
		},
		{
			"getListFollowers.json",
			func(data []byte) (interface{ RawJSON() string }, error) {
				var v xtwitterscraper.XListGetFollowersResponseUnion
				err := v.UnmarshalJSON(data)
				return v, err
			},
			nil,
		},
		{
			"getListMembers.json",
			func(data []byte) (interface{ RawJSON() string }, error) {
				var v xtwitterscraper.XListGetMembersResponseUnion
				err := v.UnmarshalJSON(data)
				return v, err
			},
			nil,
		},
		{
			"getTweetReplies.json",
			func(data []byte) (interface{ RawJSON() string }, error) {
				var v xtwitterscraper.XTweetGetRepliesResponse
				err := v.UnmarshalJSON(data)
				return v, err
			},
			nil,
		},
		{
			"getBatchUsers.json",
			func(data []byte) (interface{ RawJSON() string }, error) {
				var v xtwitterscraper.XUserGetBatchResponse
				err := v.UnmarshalJSON(data)
				return v, err
			},
			nil,
		},
		{
			"getUserFollowers.json",
			func(data []byte) (interface{ RawJSON() string }, error) {
				var v xtwitterscraper.XUserGetFollowersResponseUnion
				err := v.UnmarshalJSON(data)
				return v, err
			},
			nil,
		},
		{
			"getUserFollowing.json",
			func(data []byte) (interface{ RawJSON() string }, error) {
				var v xtwitterscraper.XUserGetFollowingResponseUnion
				err := v.UnmarshalJSON(data)
				return v, err
			},
			nil,
		},
		{
			"getUserVerifiedFollowers.json",
			func(data []byte) (interface{ RawJSON() string }, error) {
				var v xtwitterscraper.XUserGetVerifiedFollowersResponseUnion
				err := v.UnmarshalJSON(data)
				return v, err
			},
			nil,
		},
	}
	for _, test := range cases {
		t.Run(test.fixture, func(t *testing.T) {
			payload, err := fixtures.ReadFile("testdata/fixtures/" + test.fixture)
			if err != nil {
				t.Fatal(err)
			}
			decoded, err := test.decode(payload)
			if err != nil {
				t.Fatal(err)
			}
			expectDecoded(t, decoded)
			expectValues(t, decoded, test.values)
		})
	}
}

func TestExtractionRunDecodesTheStartedJob(t *testing.T) {
	var res xtwitterscraper.ExtractionRunResponse
	payload := `{"id":"ext_123","pollAfterMs":1000,"status":"running","statusUrl":"/api/v1/extractions/a1b2c3d4","toolType":"follower_explorer","waitUrl":"/api/v1/extractions/id?wait=10&limit=1&outputMode=compact"}`
	if err := res.UnmarshalJSON([]byte(payload)); err != nil {
		t.Fatal(err)
	}
	expectDecoded(t, res)
	expectValues(t, res, []fieldValue{
		field("ID", "ext_123"),
		field("PollAfterMs", int64(1000)),
		field("Status", "running"),
		field("StatusURL", "/api/v1/extractions/a1b2c3d4"),
		field("ToolType", "follower_explorer"),
		field("WaitURL", "/api/v1/extractions/id?wait=10&limit=1&outputMode=compact"),
	})
}

func TestConnectionChallengeSubmitDecodesTheEmailCodeChallenge(t *testing.T) {
	var res xtwitterscraper.XAccountConnectionChallengeSubmitResponse
	payload := `{"object":"x_account_connection_challenge","id":"xch_8vGd8Y9JvH6dV0xA","expiresAt":"2026-05-08T12:10:00Z","message":"Enter the email verification code to continue.","username":"example_user"}`
	if err := res.UnmarshalJSON([]byte(payload)); err != nil {
		t.Fatal(err)
	}
	expectDecoded(t, res)
	expectValues(t, res, []fieldValue{
		field("ID", "xch_8vGd8Y9JvH6dV0xA"),
		field("Object", "x_account_connection_challenge"),
		field("ExpiresAt", "2026-05-08T12:10:00Z"),
		field("Message", "Enter the email verification code to continue."),
		field("Username", "example_user"),
	})
}

func TestCommunityInfoDecodesTheMembersFacepile(t *testing.T) {
	var res xtwitterscraper.XCommunityGetInfoResponse
	payload := `{"community":{"id":"1500000000000000000","name":"Tesla Fans","membersFacepile":[{"id":"44196397","username":"elonmusk","name":"Elon Musk","followers":150000000,"profilePicture":"https://pbs.twimg.com/a_normal.jpg"},{"id":"11348282","profilePictureOriginal":"https://pbs.twimg.com/b.jpg"}]}}`
	if err := res.UnmarshalJSON([]byte(payload)); err != nil {
		t.Fatal(err)
	}
	facepile := res.Community.MembersFacepile
	if len(facepile) != 2 {
		t.Fatalf("decoded %s", res.RawJSON())
	}
	profile := facepile[0].AsUserProfile()
	preview := facepile[1].AsCommunityMembersFacepileCommunityMemberPreview()
	expectValues(t, profile, []fieldValue{field("ID", "44196397"), field("Username", "elonmusk"), field("Followers", int64(150000000))})
	expectValues(t, preview, []fieldValue{field("ID", "11348282"), field("ProfilePictureOriginal", "https://pbs.twimg.com/b.jpg")})
}
