// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

package xtwitterscraper_test

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"mime"
	"net/http"
	"os"
	"reflect"
	"strings"
	"testing"
	"testing/iotest"

	"github.com/Xquik-dev/x-twitter-scraper-go"
	"github.com/Xquik-dev/x-twitter-scraper-go/internal/testutil"
	"github.com/Xquik-dev/x-twitter-scraper-go/option"
	"github.com/Xquik-dev/x-twitter-scraper-go/packages/param"
)

func TestXMediaUploadURLRequest(t *testing.T) {
	sent := false
	client := xtwitterscraper.NewClient(
		option.WithBaseURL("https://example.com/api/v1/"),
		option.WithMaxRetries(0),
		option.WithHTTPClient(testHTTPDoer(func(req *http.Request) (*http.Response, error) {
			sent = true
			if req.Method != http.MethodPost || req.URL.Path != "/api/v1/x/media" {
				t.Fatalf("request = %s %s", req.Method, req.URL.Path)
			}
			if got := req.Header.Get("Idempotency-Key"); got != "upload-request" {
				t.Fatalf("Idempotency-Key = %q", got)
			}
			if got := req.Header.Get("Content-Type"); got != "application/json" {
				t.Fatalf("Content-Type = %q", got)
			}
			var got map[string]string
			if err := json.NewDecoder(req.Body).Decode(&got); err != nil {
				t.Fatal(err)
			}
			want := map[string]string{
				"account": "@elonmusk",
				"url":     "https://example.com/image.png",
			}
			if !reflect.DeepEqual(got, want) {
				t.Fatalf("JSON body = %v, want %v", got, want)
			}
			return &http.Response{
				StatusCode: http.StatusOK,
				Header:     http.Header{"Content-Type": {"application/json"}},
				Body:       io.NopCloser(strings.NewReader(`{}`)),
			}, nil
		})),
	)
	_, err := client.X.Media.Upload(context.Background(), xtwitterscraper.XMediaUploadParams{
		Account:        "@elonmusk",
		URL:            xtwitterscraper.String("https://example.com/image.png"),
		IdempotencyKey: "upload-request",
	})
	if err != nil {
		t.Fatal(err)
	}
	if !sent {
		t.Fatal("media request was not sent")
	}
}

func TestXMediaUploadFileRequest(t *testing.T) {
	tests := map[string]struct {
		content     string
		filename    string
		contentType string
		isLongVideo param.Opt[bool]
	}{
		"png": {
			content:     "\x89PNG\r\n\x1a\n" + strings.Repeat("image data", 100),
			filename:    "upload.png",
			contentType: "image/png",
		},
		"jpeg": {
			content:     "\xff\xd8\xffimage",
			filename:    "upload.jpg",
			contentType: "image/jpeg",
		},
		"gif87a": {
			content:     "GIF87aimage",
			filename:    "upload.gif",
			contentType: "image/gif",
		},
		"gif89a": {
			content:     "GIF89aimage",
			filename:    "upload.gif",
			contentType: "image/gif",
		},
		"webp": {
			content:     "RIFF\x10\x00\x00\x00WEBPimage",
			filename:    "upload.webp",
			contentType: "image/webp",
		},
		"avif": {
			content:     "\x00\x00\x00\x20ftypavifimage",
			filename:    "upload.avif",
			contentType: "image/avif",
		},
		"avis": {
			content:     "\x00\x00\x00\x20ftypavisimage",
			filename:    "upload.avif",
			contentType: "image/avif",
		},
		"mp4 isom": {
			content:     "\x00\x00\x00\x20ftypisomvideo",
			filename:    "upload.mp4",
			contentType: "video/mp4",
		},
		"mp4 iso2": {
			content:     "\x00\x00\x00\x20ftypiso2video",
			filename:    "upload.mp4",
			contentType: "video/mp4",
		},
		"mp4 mp41": {
			content:     "\x00\x00\x00\x20ftypmp41video",
			filename:    "upload.mp4",
			contentType: "video/mp4",
		},
		"mp4 mp42": {
			content:     "\x00\x00\x00\x20ftypmp42video",
			filename:    "upload.mp4",
			contentType: "video/mp4",
		},
		"mp4 avc1": {
			content:     "\x00\x00\x00\x20ftypavc1video",
			filename:    "upload.mp4",
			contentType: "video/mp4",
		},
		"mp4 M4V": {
			content:     "\x00\x00\x00\x20ftypM4V video",
			filename:    "upload.mp4",
			contentType: "video/mp4",
		},
		"long video": {
			content:     "\x00\x00\x00\x20ftypisomvideo",
			filename:    "upload.mp4",
			contentType: "video/mp4",
			isLongVideo: xtwitterscraper.Bool(true),
		},
		"short video": {
			content:     "\x00\x00\x00\x20ftypisomvideo",
			filename:    "upload.mp4",
			contentType: "video/mp4",
			isLongVideo: xtwitterscraper.Bool(false),
		},
		"unknown binary": {
			content:     "\x00\x01\x02\xffunknown",
			filename:    "upload",
			contentType: "application/octet-stream",
		},
		"text": {
			content:     "plain text",
			filename:    "upload",
			contentType: "application/octet-stream",
		},
		"empty file": {
			filename:    "upload",
			contentType: "application/octet-stream",
		},
		"truncated png": {
			content:     "\x89PNG\r\n\x1a",
			filename:    "upload",
			contentType: "application/octet-stream",
		},
		"truncated ftyp": {
			content:     "\x00\x00\x00\x20ftypiso",
			filename:    "upload",
			contentType: "application/octet-stream",
		},
		"unsupported ftyp brand": {
			content:     "\x00\x00\x00\x20ftypqt  video",
			filename:    "upload",
			contentType: "application/octet-stream",
		},
		"brand without ftyp": {
			content:     "\x00\x00\x00\x20xxxxisomvideo",
			filename:    "upload",
			contentType: "application/octet-stream",
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			sent := false
			client := xtwitterscraper.NewClient(
				option.WithBaseURL("https://example.com/api/v1/"),
				option.WithMaxRetries(0),
				option.WithHTTPClient(testHTTPDoer(func(req *http.Request) (*http.Response, error) {
					sent = true
					if req.Method != http.MethodPost || req.URL.Path != "/api/v1/x/media" {
						t.Fatalf("request = %s %s", req.Method, req.URL.Path)
					}
					if got := req.Header.Get("Idempotency-Key"); got != "upload-request" {
						t.Fatalf("Idempotency-Key = %q", got)
					}
					mediaType, _, err := mime.ParseMediaType(req.Header.Get("Content-Type"))
					if err != nil {
						t.Fatal(err)
					}
					if mediaType != "multipart/form-data" {
						t.Fatalf("Content-Type = %q", mediaType)
					}
					if err := req.ParseMultipartForm(1 << 20); err != nil {
						t.Fatal(err)
					}
					defer req.MultipartForm.RemoveAll()
					want := map[string][]string{"account": {"@elonmusk"}}
					if !param.IsOmitted(test.isLongVideo) {
						want["is_long_video"] = []string{test.isLongVideo.String()}
					}
					if !reflect.DeepEqual(req.MultipartForm.Value, want) {
						t.Fatalf("form fields = %v, want %v", req.MultipartForm.Value, want)
					}
					files := req.MultipartForm.File["file"]
					if len(req.MultipartForm.File) != 1 || len(files) != 1 {
						t.Fatalf("form files = %v", req.MultipartForm.File)
					}
					file := files[0]
					if file.Filename != test.filename || file.Header.Get("Content-Type") != test.contentType {
						t.Fatalf("file headers = %v", file.Header)
					}
					reader, err := file.Open()
					if err != nil {
						t.Fatal(err)
					}
					defer reader.Close()
					content, err := io.ReadAll(reader)
					if err != nil || string(content) != test.content {
						t.Fatalf("file = %q, error = %v", content, err)
					}
					return &http.Response{
						StatusCode: http.StatusOK,
						Header:     http.Header{"Content-Type": {"application/json"}},
						Body:       io.NopCloser(strings.NewReader(`{}`)),
					}, nil
				})),
			)
			_, err := client.X.Media.Upload(context.Background(), xtwitterscraper.XMediaUploadParams{
				Account:        "@elonmusk",
				File:           iotest.OneByteReader(strings.NewReader(test.content)),
				IsLongVideo:    test.isLongVideo,
				IdempotencyKey: "upload-request",
			})
			if err != nil {
				t.Fatal(err)
			}
			if !sent {
				t.Fatal("media request was not sent")
			}
		})
	}
}

func TestXMediaUploadInvalidParams(t *testing.T) {
	tests := map[string]struct {
		params  xtwitterscraper.XMediaUploadParams
		wantErr string
	}{
		"url and file": {
			params: xtwitterscraper.XMediaUploadParams{
				URL:  xtwitterscraper.String("https://example.com/image.png"),
				File: strings.NewReader("image"),
			},
			wantErr: "cannot use url and file parameters together",
		},
		"empty url and file": {
			params: xtwitterscraper.XMediaUploadParams{
				URL:  xtwitterscraper.String(""),
				File: strings.NewReader("image"),
			},
			wantErr: "cannot use url and file parameters together",
		},
		"neither url nor file": {
			wantErr: "missing required url or file parameter",
		},
		"long video with url": {
			params: xtwitterscraper.XMediaUploadParams{
				URL:         xtwitterscraper.String("https://example.com/video.mp4"),
				IsLongVideo: xtwitterscraper.Bool(true),
			},
			wantErr: "is_long_video parameter requires a file parameter",
		},
		"short video with url": {
			params: xtwitterscraper.XMediaUploadParams{
				URL:         xtwitterscraper.String("https://example.com/video.mp4"),
				IsLongVideo: xtwitterscraper.Bool(false),
			},
			wantErr: "is_long_video parameter requires a file parameter",
		},
		"long video without file": {
			params: xtwitterscraper.XMediaUploadParams{
				IsLongVideo: xtwitterscraper.Bool(true),
			},
			wantErr: "missing required url or file parameter",
		},
		"short video without file": {
			params: xtwitterscraper.XMediaUploadParams{
				IsLongVideo: xtwitterscraper.Bool(false),
			},
			wantErr: "missing required url or file parameter",
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			sent := false
			client := xtwitterscraper.NewClient(
				option.WithBaseURL("https://example.com/api/v1/"),
				option.WithMaxRetries(0),
				option.WithHTTPClient(testHTTPDoer(func(req *http.Request) (*http.Response, error) {
					sent = true
					return nil, errors.New("unexpected request")
				})),
			)
			test.params.Account = "@elonmusk"
			test.params.IdempotencyKey = "upload-request"
			_, err := client.X.Media.Upload(context.Background(), test.params)
			if err == nil || err.Error() != test.wantErr {
				t.Fatalf("error = %v, want %q", err, test.wantErr)
			}
			if sent {
				t.Fatal("invalid media request was sent")
			}
		})
	}
}

func TestXMediaUploadFileReadError(t *testing.T) {
	wantErr := errors.New("file read failed")
	tests := map[string]io.Reader{
		"before detection": iotest.ErrReader(wantErr),
		"after detection":  io.MultiReader(strings.NewReader("\x89PNG\r\n\x1a\nimage data"), iotest.ErrReader(wantErr)),
	}
	for name, file := range tests {
		t.Run(name, func(t *testing.T) {
			sent := false
			client := xtwitterscraper.NewClient(
				option.WithBaseURL("https://example.com/api/v1/"),
				option.WithMaxRetries(0),
				option.WithHTTPClient(testHTTPDoer(func(req *http.Request) (*http.Response, error) {
					sent = true
					return nil, errors.New("unexpected request")
				})),
			)
			_, err := client.X.Media.Upload(context.Background(), xtwitterscraper.XMediaUploadParams{
				Account:        "@elonmusk",
				File:           file,
				IdempotencyKey: "upload-request",
			})
			if !errors.Is(err, wantErr) {
				t.Fatalf("error = %v, want %v", err, wantErr)
			}
			if sent {
				t.Fatal("media request with unreadable file was sent")
			}
		})
	}
}

func TestXMediaDownloadWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := xtwitterscraper.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
		option.WithBearerToken("My Bearer Token"),
	)
	_, err := client.X.Media.Download(context.TODO(), xtwitterscraper.XMediaDownloadParams{
		TweetID:    xtwitterscraper.String("1234567890"),
		TweetIDs:   []string{"1234567890", "1234567891"},
		TweetInput: xtwitterscraper.String("https://x.com/elonmusk/status/1234567890"),
		TweetURL:   xtwitterscraper.String("https://x.com/elonmusk/status/1234567890"),
	})
	if err != nil {
		var apierr *xtwitterscraper.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestXMediaUpload(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := xtwitterscraper.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
		option.WithBearerToken("My Bearer Token"),
	)
	_, err := client.X.Media.Upload(context.TODO(), xtwitterscraper.XMediaUploadParams{
		Account:        "@elonmusk",
		URL:            xtwitterscraper.String("https://example.com/image.png"),
		IdempotencyKey: "Idempotency-Key",
	})
	if err != nil {
		var apierr *xtwitterscraper.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
