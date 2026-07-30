// Copyright the Xquik contributors.
// SPDX-License-Identifier: Apache-2.0

package xtwitterscraper

import (
	"go/ast"
	"go/parser"
	"go/token"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
)

type rawJSONModel interface {
	UnmarshalJSON([]byte) error
	RawJSON() string
}

func TestGeneratedModelsPreserveRawJSON(t *testing.T) {
	raw := []byte(`{"_contract_probe":true}`)
	models := []rawJSONModel{
		(*AccountGetResponse)(nil),
		(*AccountGetResponseCreditInfo)(nil),
		(*AccountGetResponseMonitorBilling)(nil),
		(*AccountSetXUsernameResponse)(nil),
		(*AccountUpdateLocaleResponse)(nil),
		(*ComposeNewResponseComposePrepareResult)(nil),
		(*ComposeNewResponseComposePrepareResultContentRule)(nil),
		(*ComposeNewResponseComposePrepareResultEngagementMultiplier)(nil),
		(*ComposeNewResponseComposePrepareResultRadarRecommendation)(nil),
		(*ComposeNewResponseComposePrepareResultSavedStyle)(nil),
		(*ComposeNewResponseComposePrepareResultScorerWeight)(nil),
		(*ComposeNewResponseComposeRefineResult)(nil),
		(*ComposeNewResponseComposeRefineResultExamplePattern)(nil),
		(*ComposeNewResponseComposeScoreResult)(nil),
		(*ComposeNewResponseComposeScoreResultChecklist)(nil),
		(*ComposeNewResponseUnion)(nil),
		(*ContentDisclosure)(nil),
		(*ContentDisclosureAIGenerated)(nil),
		(*ContentDisclosureAdvertising)(nil),
		(*CreditGetBalanceResponse)(nil),
		(*CreditGetTopupStatusResponse)(nil),
		(*CreditTopupBalanceResponse)(nil),
		(*Delivery)(nil),
		(*Draft)(nil),
		(*DraftDetail)(nil),
		(*DraftListResponse)(nil),
		(*DrawDetail)(nil),
		(*DrawGetResponse)(nil),
		(*DrawListItem)(nil),
		(*DrawListResponse)(nil),
		(*DrawRunResponse)(nil),
		(*EmbeddedTweet)(nil),
		(*Error)(nil),
		(*Event)(nil),
		(*EventDetail)(nil),
		(*EventListResponse)(nil),
		(*ExtractionEstimateCostResponse)(nil),
		(*ExtractionGetResponse)(nil),
		(*ExtractionJob)(nil),
		(*ExtractionListResponse)(nil),
		(*ExtractionRunResponse)(nil),
		(*GuestWalletAmount)(nil),
		(*GuestWalletGetStatusResponse)(nil),
		(*GuestWalletGetStatusResponseLatestPurchase)(nil),
		(*GuestWalletGetStatusResponseTopUp)(nil),
		(*GuestWalletNewResponse)(nil),
		(*GuestWalletNewResponseAuthorization)(nil),
		(*GuestWalletTopupResponse)(nil),
		(*GuestWalletTopupResponseAuthorization)(nil),
		(*Monitor)(nil),
		(*MonitorDeactivateResponse)(nil),
		(*MonitorKeywordDeactivateResponse)(nil),
		(*MonitorKeywordGetResponse)(nil),
		(*MonitorKeywordListResponse)(nil),
		(*MonitorKeywordListResponseMonitor)(nil),
		(*MonitorKeywordNewResponse)(nil),
		(*MonitorKeywordUpdateResponse)(nil),
		(*MonitorListResponse)(nil),
		(*MonitorNewResponse)(nil),
		(*PaginatedTweets)(nil),
		(*PaginatedUsers)(nil),
		(*RadarGetTrendingTopicsResponse)(nil),
		(*RadarItem)(nil),
		(*RadarItemMetadata)(nil),
		(*SearchTweet)(nil),
		(*StyleCompareResponse)(nil),
		(*StyleGetPerformanceResponse)(nil),
		(*StyleGetPerformanceResponseTweet)(nil),
		(*StyleListResponse)(nil),
		(*StyleProfile)(nil),
		(*StyleProfileSummary)(nil),
		(*StyleProfileTweet)(nil),
		(*SubscribeNewResponse)(nil),
		(*SupportTicketGetResponse)(nil),
		(*SupportTicketGetResponseMessage)(nil),
		(*SupportTicketGetResponseMessageAttachment)(nil),
		(*SupportTicketListResponse)(nil),
		(*SupportTicketListResponseTicket)(nil),
		(*SupportTicketNewResponse)(nil),
		(*SupportTicketNewResponseAttachment)(nil),
		(*SupportTicketReplyResponse)(nil),
		(*SupportTicketReplyResponseAttachment)(nil),
		(*SupportTicketUpdateResponse)(nil),
		(*TrendListResponse)(nil),
		(*TrendListResponseTrend)(nil),
		(*TweetAuthor)(nil),
		(*TweetDetail)(nil),
		(*TweetMedia)(nil),
		(*TweetMediaVideoVariant)(nil),
		(*UserProfile)(nil),
		(*Webhook)(nil),
		(*WebhookDeactivateResponse)(nil),
		(*WebhookListDeliveriesResponse)(nil),
		(*WebhookListResponse)(nil),
		(*WebhookNewResponse)(nil),
		(*WebhookResumeResponse)(nil),
		(*WebhookTestResponse)(nil),
		(*Winner)(nil),
		(*XAccount)(nil),
		(*XAccountBulkRetryResponse)(nil),
		(*XAccountConnectionAttemptGetResponseUnion)(nil),
		(*XAccountConnectionAttemptGetResponseXAccountConnectionAttemptFailed)(nil),
		(*XAccountConnectionAttemptGetResponseXAccountConnectionAttemptPending)(nil),
		(*XAccountConnectionAttemptGetResponseXAccountConnectionAttemptSuccess)(nil),
		(*XAccountConnectionAttemptGetResponseXAccountConnectionChallenge)(nil),
		(*XAccountConnectionChallengeSubmitResponse)(nil),
		(*XAccountDeleteResponse)(nil),
		(*XAccountDetail)(nil),
		(*XAccountListResponse)(nil),
		(*XAccountReauthResponse)(nil),
		(*XBookmarkGetFoldersResponse)(nil),
		(*XBookmarkGetFoldersResponseFolder)(nil),
		(*XCommunityDeleteResponse)(nil),
		(*XCommunityDeleteResponseAccount)(nil),
		(*XCommunityDeleteResponseBilling)(nil),
		(*XCommunityDeleteResponseNextAction)(nil),
		(*XCommunityDeleteResponseRequest)(nil),
		(*XCommunityDeleteResponseResult)(nil),
		(*XCommunityDeleteResponseTarget)(nil),
		(*XCommunityGetInfoResponse)(nil),
		(*XCommunityGetInfoResponseCommunity)(nil),
		(*XCommunityGetInfoResponseCommunityCreator)(nil),
		(*XCommunityGetInfoResponseCommunityPrimaryTopic)(nil),
		(*XCommunityGetInfoResponseCommunityRule)(nil),
		(*XCommunityJoinDeleteAllResponse)(nil),
		(*XCommunityJoinDeleteAllResponseAccount)(nil),
		(*XCommunityJoinDeleteAllResponseBilling)(nil),
		(*XCommunityJoinDeleteAllResponseNextAction)(nil),
		(*XCommunityJoinDeleteAllResponseRequest)(nil),
		(*XCommunityJoinDeleteAllResponseResult)(nil),
		(*XCommunityJoinDeleteAllResponseTarget)(nil),
		(*XCommunityJoinNewResponse)(nil),
		(*XCommunityJoinNewResponseAccount)(nil),
		(*XCommunityJoinNewResponseBilling)(nil),
		(*XCommunityJoinNewResponseNextAction)(nil),
		(*XCommunityJoinNewResponseRequest)(nil),
		(*XCommunityJoinNewResponseResult)(nil),
		(*XCommunityJoinNewResponseTarget)(nil),
		(*XCommunityNewResponse)(nil),
		(*XCommunityNewResponseAccount)(nil),
		(*XCommunityNewResponseBilling)(nil),
		(*XCommunityNewResponseNextAction)(nil),
		(*XCommunityNewResponseRequest)(nil),
		(*XCommunityNewResponseResult)(nil),
		(*XCommunityNewResponseTarget)(nil),
		(*XDmGetHistoryResponse)(nil),
		(*XDmGetHistoryResponseMessage)(nil),
		(*XDmSendResponse)(nil),
		(*XDmSendResponseAccount)(nil),
		(*XDmSendResponseBilling)(nil),
		(*XDmSendResponseNextAction)(nil),
		(*XDmSendResponseRequest)(nil),
		(*XDmSendResponseResult)(nil),
		(*XDmSendResponseTarget)(nil),
		(*XFollowerCheckResponse)(nil),
		(*XGetArticleResponse)(nil),
		(*XGetArticleResponseArticle)(nil),
		(*XGetArticleResponseArticleContent)(nil),
		(*XGetArticleResponseArticleContentInlineStyleRange)(nil),
		(*XGetArticleResponseAuthor)(nil),
		(*XGetNotificationsResponse)(nil),
		(*XGetNotificationsResponseNotification)(nil),
		(*XGetTrendsResponse)(nil),
		(*XGetTrendsResponseTrend)(nil),
		(*XMediaDownloadResponse)(nil),
		(*XMediaUploadResponse)(nil),
		(*XMediaUploadResponseAccount)(nil),
		(*XMediaUploadResponseBilling)(nil),
		(*XMediaUploadResponseNextAction)(nil),
		(*XMediaUploadResponseRequest)(nil),
		(*XMediaUploadResponseResult)(nil),
		(*XMediaUploadResponseTarget)(nil),
		(*XProfileUpdateAvatarResponse)(nil),
		(*XProfileUpdateAvatarResponseAccount)(nil),
		(*XProfileUpdateAvatarResponseBilling)(nil),
		(*XProfileUpdateAvatarResponseNextAction)(nil),
		(*XProfileUpdateAvatarResponseRequest)(nil),
		(*XProfileUpdateAvatarResponseResult)(nil),
		(*XProfileUpdateAvatarResponseTarget)(nil),
		(*XProfileUpdateBannerResponse)(nil),
		(*XProfileUpdateBannerResponseAccount)(nil),
		(*XProfileUpdateBannerResponseBilling)(nil),
		(*XProfileUpdateBannerResponseNextAction)(nil),
		(*XProfileUpdateBannerResponseRequest)(nil),
		(*XProfileUpdateBannerResponseResult)(nil),
		(*XProfileUpdateBannerResponseTarget)(nil),
		(*XProfileUpdateResponse)(nil),
		(*XProfileUpdateResponseAccount)(nil),
		(*XProfileUpdateResponseBilling)(nil),
		(*XProfileUpdateResponseNextAction)(nil),
		(*XProfileUpdateResponseRequest)(nil),
		(*XProfileUpdateResponseResult)(nil),
		(*XProfileUpdateResponseTarget)(nil),
		(*XTweetDeleteResponse)(nil),
		(*XTweetDeleteResponseAccount)(nil),
		(*XTweetDeleteResponseBilling)(nil),
		(*XTweetDeleteResponseNextAction)(nil),
		(*XTweetDeleteResponseRequest)(nil),
		(*XTweetDeleteResponseResult)(nil),
		(*XTweetDeleteResponseTarget)(nil),
		(*XTweetGetResponse)(nil),
		(*XTweetLikeDeleteResponse)(nil),
		(*XTweetLikeDeleteResponseAccount)(nil),
		(*XTweetLikeDeleteResponseBilling)(nil),
		(*XTweetLikeDeleteResponseNextAction)(nil),
		(*XTweetLikeDeleteResponseRequest)(nil),
		(*XTweetLikeDeleteResponseResult)(nil),
		(*XTweetLikeDeleteResponseTarget)(nil),
		(*XTweetLikeNewResponse)(nil),
		(*XTweetLikeNewResponseAccount)(nil),
		(*XTweetLikeNewResponseBilling)(nil),
		(*XTweetLikeNewResponseNextAction)(nil),
		(*XTweetLikeNewResponseRequest)(nil),
		(*XTweetLikeNewResponseResult)(nil),
		(*XTweetLikeNewResponseTarget)(nil),
		(*XTweetNewResponse)(nil),
		(*XTweetNewResponseAccount)(nil),
		(*XTweetNewResponseBilling)(nil),
		(*XTweetNewResponseNextAction)(nil),
		(*XTweetNewResponseRequest)(nil),
		(*XTweetNewResponseResult)(nil),
		(*XTweetNewResponseTarget)(nil),
		(*XTweetRetweetDeleteResponse)(nil),
		(*XTweetRetweetDeleteResponseAccount)(nil),
		(*XTweetRetweetDeleteResponseBilling)(nil),
		(*XTweetRetweetDeleteResponseNextAction)(nil),
		(*XTweetRetweetDeleteResponseRequest)(nil),
		(*XTweetRetweetDeleteResponseResult)(nil),
		(*XTweetRetweetDeleteResponseTarget)(nil),
		(*XTweetRetweetNewResponse)(nil),
		(*XTweetRetweetNewResponseAccount)(nil),
		(*XTweetRetweetNewResponseBilling)(nil),
		(*XTweetRetweetNewResponseNextAction)(nil),
		(*XTweetRetweetNewResponseRequest)(nil),
		(*XTweetRetweetNewResponseResult)(nil),
		(*XTweetRetweetNewResponseTarget)(nil),
		(*XUserFollowDeleteAllResponse)(nil),
		(*XUserFollowDeleteAllResponseAccount)(nil),
		(*XUserFollowDeleteAllResponseBilling)(nil),
		(*XUserFollowDeleteAllResponseNextAction)(nil),
		(*XUserFollowDeleteAllResponseRequest)(nil),
		(*XUserFollowDeleteAllResponseResult)(nil),
		(*XUserFollowDeleteAllResponseTarget)(nil),
		(*XUserFollowNewResponse)(nil),
		(*XUserFollowNewResponseAccount)(nil),
		(*XUserFollowNewResponseBilling)(nil),
		(*XUserFollowNewResponseNextAction)(nil),
		(*XUserFollowNewResponseRequest)(nil),
		(*XUserFollowNewResponseResult)(nil),
		(*XUserFollowNewResponseTarget)(nil),
		(*XUserGetBatchResponse)(nil),
		(*XUserRemoveFollowerResponse)(nil),
		(*XUserRemoveFollowerResponseAccount)(nil),
		(*XUserRemoveFollowerResponseBilling)(nil),
		(*XUserRemoveFollowerResponseNextAction)(nil),
		(*XUserRemoveFollowerResponseRequest)(nil),
		(*XUserRemoveFollowerResponseResult)(nil),
		(*XUserRemoveFollowerResponseTarget)(nil),
		(*XWriteActionGetResponse)(nil),
		(*XWriteActionGetResponseAccount)(nil),
		(*XWriteActionGetResponseBilling)(nil),
		(*XWriteActionGetResponseNextAction)(nil),
		(*XWriteActionGetResponseRequest)(nil),
		(*XWriteActionGetResponseResult)(nil),
		(*XWriteActionGetResponseTarget)(nil),
	}
	registered := make(map[string]struct{}, len(models))
	for _, template := range models {
		modelType := reflect.TypeOf(template).Elem()
		if _, duplicate := registered[modelType.Name()]; duplicate {
			t.Fatalf("duplicate model registration: %s", modelType.Name())
		}
		registered[modelType.Name()] = struct{}{}
		t.Run(modelType.Name(), func(t *testing.T) {
			model := reflect.New(modelType).Interface().(rawJSONModel)
			if err := model.UnmarshalJSON(raw); err != nil {
				t.Fatalf("UnmarshalJSON() error = %v", err)
			}
			if got := model.RawJSON(); got != string(raw) {
				t.Fatalf("RawJSON() = %q, want %q", got, raw)
			}
		})
	}

	discovered := rawJSONReceiverNames(t)
	for name := range discovered {
		if _, ok := registered[name]; !ok {
			t.Errorf("generated model is not registered: %s", name)
		}
	}
	for name := range registered {
		if _, ok := discovered[name]; !ok {
			t.Errorf("registered model has no RawJSON method: %s", name)
		}
	}
	if len(registered) != len(discovered) {
		t.Fatalf(
			"registered %d generated models, discovered %d",
			len(registered),
			len(discovered),
		)
	}
}

func rawJSONReceiverNames(t *testing.T) map[string]struct{} {
	t.Helper()

	names := rawJSONReceiverNamesInDirectory(t, ".")
	aliasTargets := map[string]map[string]struct{}{
		"apierror": rawJSONReceiverNamesInDirectory(t, "internal/apierror"),
		"shared":   rawJSONReceiverNamesInDirectory(t, "shared"),
	}

	aliasFile, err := parser.ParseFile(token.NewFileSet(), "aliases.go", nil, 0)
	if err != nil {
		t.Fatal(err)
	}
	for _, declaration := range aliasFile.Decls {
		general, ok := declaration.(*ast.GenDecl)
		if !ok || general.Tok != token.TYPE {
			continue
		}
		for _, specification := range general.Specs {
			alias, ok := specification.(*ast.TypeSpec)
			if !ok || !alias.Assign.IsValid() {
				continue
			}
			target, ok := alias.Type.(*ast.SelectorExpr)
			if !ok {
				continue
			}
			packageName, ok := target.X.(*ast.Ident)
			if !ok {
				continue
			}
			targetNames := aliasTargets[packageName.Name]
			if _, ok := targetNames[target.Sel.Name]; ok {
				names[alias.Name.Name] = struct{}{}
			}
		}
	}
	return names
}

func rawJSONReceiverNamesInDirectory(
	t *testing.T,
	directory string,
) map[string]struct{} {
	t.Helper()

	paths, err := filepath.Glob(filepath.Join(directory, "*.go"))
	if err != nil {
		t.Fatal(err)
	}
	names := make(map[string]struct{})
	for _, path := range paths {
		if strings.HasSuffix(path, "_test.go") {
			continue
		}
		file, err := parser.ParseFile(token.NewFileSet(), path, nil, 0)
		if err != nil {
			t.Fatalf("parse %s: %v", path, err)
		}
		for _, declaration := range file.Decls {
			function, ok := declaration.(*ast.FuncDecl)
			if !ok || function.Name.Name != "RawJSON" || function.Recv == nil {
				continue
			}
			receiver := function.Recv.List[0].Type
			if pointer, ok := receiver.(*ast.StarExpr); ok {
				receiver = pointer.X
			}
			identifier, ok := receiver.(*ast.Ident)
			if ok && ast.IsExported(identifier.Name) {
				names[identifier.Name] = struct{}{}
			}
		}
	}
	return names
}
