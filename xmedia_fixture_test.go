package xtwitterscraper_test

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/Xquik-dev/x-twitter-scraper-go"
)

func TestXMediaDownloadFixture(t *testing.T) {
	client, request := fixtureClient(t, "downloadMedia.json")
	res, err := client.X.Media.Download(context.Background(), xtwitterscraper.XMediaDownloadParams{
		TweetInput: xtwitterscraper.String("1234567890"),
	})
	if err != nil {
		t.Fatal(err)
	}
	request.expect(t, http.MethodPost, "/api/v1/x/media/download", "")
	request.expectBody(t, `{"tweetInput":"1234567890"}`)
	expectDecoded(t, res)
	if res.TweetID != "1234567890" || res.GalleryURL != "https://xquik.com/gallery/abc123" || res.CacheHit || len(res.Items) != 1 {
		t.Fatalf("decoded %s", res.RawJSON())
	}
	item := res.Items[0]
	if item.ID != "42" || item.TweetID != "1234567890" || item.MediaType != "photo" ||
		item.MediaURL != "https://media.xquik.com/tweet_1234567890/0.jpg" || item.FileSize != "2097152" ||
		!item.CreatedAt.Equal(time.Date(2026, time.October, 16, 11, 0, 0, 0, time.UTC)) ||
		item.AuthorUsername != "elonmusk" || item.AuthorDisplayName != "Elon Musk" ||
		item.AuthorAvatar != "" || item.JSON.AuthorAvatar.Raw() != "null" ||
		item.TweetText != "Launch day." || item.Likes != 1200 || item.Retweets != 340 || item.Views != 56000 {
		t.Fatalf("decoded %s", item.RawJSON())
	}
}
