// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

package xtwitterscraper_test

import (
	"slices"
	"testing"

	"github.com/Xquik-dev/x-twitter-scraper-go"
	"github.com/Xquik-dev/x-twitter-scraper-go/packages/param"
)

type writeCall func(xtwitterscraper.Client, param.Opt[string]) (*xtwitterscraper.XWriteActionGetResponse, error)

func expectWrite(t *testing.T, fixture, method, path string, action xtwitterscraper.XWriteActionGetResponseAction, call writeCall) {
	t.Helper()
	client, request := fixtureClient(t, fixture)
	res, err := call(client, xtwitterscraper.String("wait=50"))
	if err != nil {
		t.Fatal(err)
	}
	request.expect(t, method, path, "")
	request.expectBody(t, `{"account":"@elonmusk"}`)
	expectDecoded(t, res)
	if request.header.Get("Idempotency-Key") != "write-1" || request.header.Get("Prefer") != "wait=50" {
		t.Fatalf("headers = %v", request.header)
	}
	if res.ID != "12345" ||
		res.Action != action ||
		res.Status != "success" ||
		!res.Terminal ||
		res.StatusURL != "/api/v1/x/write-actions/12345" ||
		res.ChargedCredits != "10" ||
		res.Account.Username != "example" ||
		!res.Success {
		t.Fatalf("decoded %s", res.RawJSON())
	}
	if _, err := call(client, param.Opt[string]{}); err != nil {
		t.Fatal(err)
	}
	if prefer, sent := request.header["Prefer"]; sent {
		t.Fatalf("an omitted Prefer sent %v", prefer)
	}
}

func actionNames[T ~string](actions ...T) []string {
	names := make([]string, 0, len(actions))
	for _, action := range actions {
		names = append(names, string(action))
	}
	return names
}

func TestEveryWriteRecordNamesTheBookmarkBlockAndMuteActions(t *testing.T) {
	want := []string{"bookmark", "unbookmark", "block", "unblock", "mute", "unmute"}
	for name, got := range map[string][]string{
		"XCommunityDeleteResponseAction":        actionNames(xtwitterscraper.XCommunityDeleteResponseActionBookmark, xtwitterscraper.XCommunityDeleteResponseActionUnbookmark, xtwitterscraper.XCommunityDeleteResponseActionBlock, xtwitterscraper.XCommunityDeleteResponseActionUnblock, xtwitterscraper.XCommunityDeleteResponseActionMute, xtwitterscraper.XCommunityDeleteResponseActionUnmute),
		"XCommunityJoinDeleteAllResponseAction": actionNames(xtwitterscraper.XCommunityJoinDeleteAllResponseActionBookmark, xtwitterscraper.XCommunityJoinDeleteAllResponseActionUnbookmark, xtwitterscraper.XCommunityJoinDeleteAllResponseActionBlock, xtwitterscraper.XCommunityJoinDeleteAllResponseActionUnblock, xtwitterscraper.XCommunityJoinDeleteAllResponseActionMute, xtwitterscraper.XCommunityJoinDeleteAllResponseActionUnmute),
		"XCommunityJoinNewResponseAction":       actionNames(xtwitterscraper.XCommunityJoinNewResponseActionBookmark, xtwitterscraper.XCommunityJoinNewResponseActionUnbookmark, xtwitterscraper.XCommunityJoinNewResponseActionBlock, xtwitterscraper.XCommunityJoinNewResponseActionUnblock, xtwitterscraper.XCommunityJoinNewResponseActionMute, xtwitterscraper.XCommunityJoinNewResponseActionUnmute),
		"XCommunityNewResponseAction":           actionNames(xtwitterscraper.XCommunityNewResponseActionBookmark, xtwitterscraper.XCommunityNewResponseActionUnbookmark, xtwitterscraper.XCommunityNewResponseActionBlock, xtwitterscraper.XCommunityNewResponseActionUnblock, xtwitterscraper.XCommunityNewResponseActionMute, xtwitterscraper.XCommunityNewResponseActionUnmute),
		"XDmDeleteMessageResponseAction":        actionNames(xtwitterscraper.XDmDeleteMessageResponseActionBookmark, xtwitterscraper.XDmDeleteMessageResponseActionUnbookmark, xtwitterscraper.XDmDeleteMessageResponseActionBlock, xtwitterscraper.XDmDeleteMessageResponseActionUnblock, xtwitterscraper.XDmDeleteMessageResponseActionMute, xtwitterscraper.XDmDeleteMessageResponseActionUnmute),
		"XDmSendResponseAction":                 actionNames(xtwitterscraper.XDmSendResponseActionBookmark, xtwitterscraper.XDmSendResponseActionUnbookmark, xtwitterscraper.XDmSendResponseActionBlock, xtwitterscraper.XDmSendResponseActionUnblock, xtwitterscraper.XDmSendResponseActionMute, xtwitterscraper.XDmSendResponseActionUnmute),
		"XMediaUploadResponseAction":            actionNames(xtwitterscraper.XMediaUploadResponseActionBookmark, xtwitterscraper.XMediaUploadResponseActionUnbookmark, xtwitterscraper.XMediaUploadResponseActionBlock, xtwitterscraper.XMediaUploadResponseActionUnblock, xtwitterscraper.XMediaUploadResponseActionMute, xtwitterscraper.XMediaUploadResponseActionUnmute),
		"XProfileUpdateAvatarResponseAction":    actionNames(xtwitterscraper.XProfileUpdateAvatarResponseActionBookmark, xtwitterscraper.XProfileUpdateAvatarResponseActionUnbookmark, xtwitterscraper.XProfileUpdateAvatarResponseActionBlock, xtwitterscraper.XProfileUpdateAvatarResponseActionUnblock, xtwitterscraper.XProfileUpdateAvatarResponseActionMute, xtwitterscraper.XProfileUpdateAvatarResponseActionUnmute),
		"XProfileUpdateBannerResponseAction":    actionNames(xtwitterscraper.XProfileUpdateBannerResponseActionBookmark, xtwitterscraper.XProfileUpdateBannerResponseActionUnbookmark, xtwitterscraper.XProfileUpdateBannerResponseActionBlock, xtwitterscraper.XProfileUpdateBannerResponseActionUnblock, xtwitterscraper.XProfileUpdateBannerResponseActionMute, xtwitterscraper.XProfileUpdateBannerResponseActionUnmute),
		"XProfileUpdateResponseAction":          actionNames(xtwitterscraper.XProfileUpdateResponseActionBookmark, xtwitterscraper.XProfileUpdateResponseActionUnbookmark, xtwitterscraper.XProfileUpdateResponseActionBlock, xtwitterscraper.XProfileUpdateResponseActionUnblock, xtwitterscraper.XProfileUpdateResponseActionMute, xtwitterscraper.XProfileUpdateResponseActionUnmute),
		"XTweetDeleteResponseAction":            actionNames(xtwitterscraper.XTweetDeleteResponseActionBookmark, xtwitterscraper.XTweetDeleteResponseActionUnbookmark, xtwitterscraper.XTweetDeleteResponseActionBlock, xtwitterscraper.XTweetDeleteResponseActionUnblock, xtwitterscraper.XTweetDeleteResponseActionMute, xtwitterscraper.XTweetDeleteResponseActionUnmute),
		"XTweetLikeDeleteResponseAction":        actionNames(xtwitterscraper.XTweetLikeDeleteResponseActionBookmark, xtwitterscraper.XTweetLikeDeleteResponseActionUnbookmark, xtwitterscraper.XTweetLikeDeleteResponseActionBlock, xtwitterscraper.XTweetLikeDeleteResponseActionUnblock, xtwitterscraper.XTweetLikeDeleteResponseActionMute, xtwitterscraper.XTweetLikeDeleteResponseActionUnmute),
		"XTweetLikeNewResponseAction":           actionNames(xtwitterscraper.XTweetLikeNewResponseActionBookmark, xtwitterscraper.XTweetLikeNewResponseActionUnbookmark, xtwitterscraper.XTweetLikeNewResponseActionBlock, xtwitterscraper.XTweetLikeNewResponseActionUnblock, xtwitterscraper.XTweetLikeNewResponseActionMute, xtwitterscraper.XTweetLikeNewResponseActionUnmute),
		"XTweetNewResponseAction":               actionNames(xtwitterscraper.XTweetNewResponseActionBookmark, xtwitterscraper.XTweetNewResponseActionUnbookmark, xtwitterscraper.XTweetNewResponseActionBlock, xtwitterscraper.XTweetNewResponseActionUnblock, xtwitterscraper.XTweetNewResponseActionMute, xtwitterscraper.XTweetNewResponseActionUnmute),
		"XTweetRetweetDeleteResponseAction":     actionNames(xtwitterscraper.XTweetRetweetDeleteResponseActionBookmark, xtwitterscraper.XTweetRetweetDeleteResponseActionUnbookmark, xtwitterscraper.XTweetRetweetDeleteResponseActionBlock, xtwitterscraper.XTweetRetweetDeleteResponseActionUnblock, xtwitterscraper.XTweetRetweetDeleteResponseActionMute, xtwitterscraper.XTweetRetweetDeleteResponseActionUnmute),
		"XTweetRetweetNewResponseAction":        actionNames(xtwitterscraper.XTweetRetweetNewResponseActionBookmark, xtwitterscraper.XTweetRetweetNewResponseActionUnbookmark, xtwitterscraper.XTweetRetweetNewResponseActionBlock, xtwitterscraper.XTweetRetweetNewResponseActionUnblock, xtwitterscraper.XTweetRetweetNewResponseActionMute, xtwitterscraper.XTweetRetweetNewResponseActionUnmute),
		"XUserFollowDeleteAllResponseAction":    actionNames(xtwitterscraper.XUserFollowDeleteAllResponseActionBookmark, xtwitterscraper.XUserFollowDeleteAllResponseActionUnbookmark, xtwitterscraper.XUserFollowDeleteAllResponseActionBlock, xtwitterscraper.XUserFollowDeleteAllResponseActionUnblock, xtwitterscraper.XUserFollowDeleteAllResponseActionMute, xtwitterscraper.XUserFollowDeleteAllResponseActionUnmute),
		"XUserFollowNewResponseAction":          actionNames(xtwitterscraper.XUserFollowNewResponseActionBookmark, xtwitterscraper.XUserFollowNewResponseActionUnbookmark, xtwitterscraper.XUserFollowNewResponseActionBlock, xtwitterscraper.XUserFollowNewResponseActionUnblock, xtwitterscraper.XUserFollowNewResponseActionMute, xtwitterscraper.XUserFollowNewResponseActionUnmute),
		"XUserRemoveFollowerResponseAction":     actionNames(xtwitterscraper.XUserRemoveFollowerResponseActionBookmark, xtwitterscraper.XUserRemoveFollowerResponseActionUnbookmark, xtwitterscraper.XUserRemoveFollowerResponseActionBlock, xtwitterscraper.XUserRemoveFollowerResponseActionUnblock, xtwitterscraper.XUserRemoveFollowerResponseActionMute, xtwitterscraper.XUserRemoveFollowerResponseActionUnmute),
		"XWriteActionGetResponseAction":         actionNames(xtwitterscraper.XWriteActionGetResponseActionBookmark, xtwitterscraper.XWriteActionGetResponseActionUnbookmark, xtwitterscraper.XWriteActionGetResponseActionBlock, xtwitterscraper.XWriteActionGetResponseActionUnblock, xtwitterscraper.XWriteActionGetResponseActionMute, xtwitterscraper.XWriteActionGetResponseActionUnmute),
	} {
		if !slices.Equal(got, want) {
			t.Errorf("%s = %v, want %v", name, got, want)
		}
	}
}
