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
	"mime/multipart"
	"net/http"
	"os"
	"reflect"
	"strings"
	"testing"
	"testing/iotest"

	"github.com/Xquik-dev/x-twitter-scraper-go"
	"github.com/Xquik-dev/x-twitter-scraper-go/internal/testutil"
	"github.com/Xquik-dev/x-twitter-scraper-go/option"
)

func TestXProfileUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.X.Profile.Update(context.TODO(), xtwitterscraper.XProfileUpdateParams{
		Account:        "@elonmusk",
		IdempotencyKey: "Idempotency-Key",
		Description:    xtwitterscraper.String("description_value"),
		Location:       xtwitterscraper.String("location_value"),
		Name:           xtwitterscraper.String("Example Name"),
		URL:            xtwitterscraper.String("https://xquik.com/example"),
	})
	if err != nil {
		var apierr *xtwitterscraper.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestXProfileUpdateAvatar(t *testing.T) {
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
	_, err := client.X.Profile.UpdateAvatar(context.TODO(), xtwitterscraper.XProfileUpdateAvatarParams{
		Account:        "@elonmusk",
		URL:            xtwitterscraper.String("https://example.com/avatar.png"),
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

func TestXProfileUpdateBanner(t *testing.T) {
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
	_, err := client.X.Profile.UpdateBanner(context.TODO(), xtwitterscraper.XProfileUpdateBannerParams{
		Account:        "@elonmusk",
		URL:            xtwitterscraper.String("https://example.com/banner.png"),
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

func TestXProfileUpdateAvatarURLRequest(t *testing.T) {
	sent := false
	client := xtwitterscraper.NewClient(
		option.WithBaseURL("https://example.com/api/v1/"),
		option.WithMaxRetries(0),
		option.WithHTTPClient(testHTTPDoer(func(req *http.Request) (*http.Response, error) {
			sent = true
			if req.Method != http.MethodPatch || req.URL.Path != "/api/v1/x/profile/avatar" {
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
				"url":     "https://example.com/avatar.png",
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
	_, err := client.X.Profile.UpdateAvatar(context.Background(), xtwitterscraper.XProfileUpdateAvatarParams{
		Account:        "@elonmusk",
		URL:            xtwitterscraper.String("https://example.com/avatar.png"),
		IdempotencyKey: "upload-request",
	})
	if err != nil {
		t.Fatal(err)
	}
	if !sent {
		t.Fatal("avatar request was not sent")
	}
}

func TestXProfileUpdateAvatarFileRequest(t *testing.T) {
	tests := map[string]struct {
		content     string
		filename    string
		contentType string
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
		"unknown binary": {
			content:     "\x00\x01\x02\xffunknown",
			filename:    "upload",
			contentType: "application/octet-stream",
		},
		"empty file": {
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
					if req.Method != http.MethodPatch || req.URL.Path != "/api/v1/x/profile/avatar" {
						t.Fatalf("request = %s %s", req.Method, req.URL.Path)
					}
					if got := req.Header.Get("Idempotency-Key"); got != "upload-request" {
						t.Fatalf("Idempotency-Key = %q", got)
					}
					mediaType, params, err := mime.ParseMediaType(req.Header.Get("Content-Type"))
					if err != nil {
						t.Fatal(err)
					}
					if mediaType != "multipart/form-data" {
						t.Fatalf("Content-Type = %q", mediaType)
					}
					form, err := multipart.NewReader(req.Body, params["boundary"]).ReadForm(1 << 20)
					if err != nil {
						t.Fatal(err)
					}
					defer form.RemoveAll()
					want := map[string][]string{"account": {"@elonmusk"}}
					if !reflect.DeepEqual(form.Value, want) {
						t.Fatalf("form fields = %v, want %v", form.Value, want)
					}
					files := form.File["file"]
					if len(form.File) != 1 || len(files) != 1 {
						t.Fatalf("form files = %v", form.File)
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
			_, err := client.X.Profile.UpdateAvatar(context.Background(), xtwitterscraper.XProfileUpdateAvatarParams{
				Account:        "@elonmusk",
				File:           iotest.OneByteReader(strings.NewReader(test.content)),
				IdempotencyKey: "upload-request",
			})
			if err != nil {
				t.Fatal(err)
			}
			if !sent {
				t.Fatal("avatar request was not sent")
			}
		})
	}
}

func TestXProfileUpdateAvatarMarshalMultipart(t *testing.T) {
	tests := map[string]xtwitterscraper.XProfileUpdateAvatarParams{
		"without file": {
			Account: "@elonmusk",
			URL:     xtwitterscraper.String("https://example.com/avatar.png"),
		},
		"with file": {
			Account: "@elonmusk",
			File:    strings.NewReader("\x89PNG\r\n\x1a\nimage"),
		},
		"empty params": {},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			data, contentType, err := test.MarshalMultipart()
			if err != nil {
				t.Fatal(err)
			}
			mediaType, params, err := mime.ParseMediaType(contentType)
			if err != nil {
				t.Fatal(err)
			}
			if mediaType != "multipart/form-data" || params["boundary"] == "" {
				t.Fatalf("Content-Type = %q", contentType)
			}
			form, err := multipart.NewReader(strings.NewReader(string(data)), params["boundary"]).ReadForm(1 << 20)
			if err != nil {
				t.Fatal(err)
			}
			defer form.RemoveAll()
		})
	}
}

func TestXProfileUpdateAvatarInvalidParams(t *testing.T) {
	tests := map[string]struct {
		params  xtwitterscraper.XProfileUpdateAvatarParams
		wantErr error
	}{
		"url and file": {
			params: xtwitterscraper.XProfileUpdateAvatarParams{
				URL:  xtwitterscraper.String("https://example.com/avatar.png"),
				File: strings.NewReader("image"),
			},
			wantErr: xtwitterscraper.ErrMediaHasURLAndFile,
		},
		"empty url and file": {
			params: xtwitterscraper.XProfileUpdateAvatarParams{
				URL:  xtwitterscraper.String(""),
				File: strings.NewReader("image"),
			},
			wantErr: xtwitterscraper.ErrMediaHasURLAndFile,
		},
		"neither url nor file": {
			wantErr: xtwitterscraper.ErrMediaNeedsURLOrFile,
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
			_, err := client.X.Profile.UpdateAvatar(context.Background(), test.params)
			if !errors.Is(err, test.wantErr) {
				t.Fatalf("error = %v, want %v", err, test.wantErr)
			}
			if sent {
				t.Fatal("invalid avatar request was sent")
			}
		})
	}
}

func TestXProfileUpdateAvatarFileReadError(t *testing.T) {
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
			_, err := client.X.Profile.UpdateAvatar(context.Background(), xtwitterscraper.XProfileUpdateAvatarParams{
				Account:        "@elonmusk",
				File:           file,
				IdempotencyKey: "upload-request",
			})
			if !errors.Is(err, wantErr) {
				t.Fatalf("error = %v, want %v", err, wantErr)
			}
			if sent {
				t.Fatal("avatar request with unreadable file was sent")
			}
		})
	}
}

func TestXProfileUpdateBannerURLRequest(t *testing.T) {
	sent := false
	client := xtwitterscraper.NewClient(
		option.WithBaseURL("https://example.com/api/v1/"),
		option.WithMaxRetries(0),
		option.WithHTTPClient(testHTTPDoer(func(req *http.Request) (*http.Response, error) {
			sent = true
			if req.Method != http.MethodPatch || req.URL.Path != "/api/v1/x/profile/banner" {
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
				"url":     "https://example.com/banner.png",
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
	_, err := client.X.Profile.UpdateBanner(context.Background(), xtwitterscraper.XProfileUpdateBannerParams{
		Account:        "@elonmusk",
		URL:            xtwitterscraper.String("https://example.com/banner.png"),
		IdempotencyKey: "upload-request",
	})
	if err != nil {
		t.Fatal(err)
	}
	if !sent {
		t.Fatal("banner request was not sent")
	}
}

func TestXProfileUpdateBannerFileRequest(t *testing.T) {
	tests := map[string]struct {
		content     string
		filename    string
		contentType string
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
		"unknown binary": {
			content:     "\x00\x01\x02\xffunknown",
			filename:    "upload",
			contentType: "application/octet-stream",
		},
		"empty file": {
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
					if req.Method != http.MethodPatch || req.URL.Path != "/api/v1/x/profile/banner" {
						t.Fatalf("request = %s %s", req.Method, req.URL.Path)
					}
					if got := req.Header.Get("Idempotency-Key"); got != "upload-request" {
						t.Fatalf("Idempotency-Key = %q", got)
					}
					mediaType, params, err := mime.ParseMediaType(req.Header.Get("Content-Type"))
					if err != nil {
						t.Fatal(err)
					}
					if mediaType != "multipart/form-data" {
						t.Fatalf("Content-Type = %q", mediaType)
					}
					form, err := multipart.NewReader(req.Body, params["boundary"]).ReadForm(1 << 20)
					if err != nil {
						t.Fatal(err)
					}
					defer form.RemoveAll()
					want := map[string][]string{"account": {"@elonmusk"}}
					if !reflect.DeepEqual(form.Value, want) {
						t.Fatalf("form fields = %v, want %v", form.Value, want)
					}
					files := form.File["file"]
					if len(form.File) != 1 || len(files) != 1 {
						t.Fatalf("form files = %v", form.File)
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
			_, err := client.X.Profile.UpdateBanner(context.Background(), xtwitterscraper.XProfileUpdateBannerParams{
				Account:        "@elonmusk",
				File:           iotest.OneByteReader(strings.NewReader(test.content)),
				IdempotencyKey: "upload-request",
			})
			if err != nil {
				t.Fatal(err)
			}
			if !sent {
				t.Fatal("banner request was not sent")
			}
		})
	}
}

func TestXProfileUpdateBannerMarshalMultipart(t *testing.T) {
	tests := map[string]xtwitterscraper.XProfileUpdateBannerParams{
		"without file": {
			Account: "@elonmusk",
			URL:     xtwitterscraper.String("https://example.com/banner.png"),
		},
		"with file": {
			Account: "@elonmusk",
			File:    strings.NewReader("\x89PNG\r\n\x1a\nimage"),
		},
		"empty params": {},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			data, contentType, err := test.MarshalMultipart()
			if err != nil {
				t.Fatal(err)
			}
			mediaType, params, err := mime.ParseMediaType(contentType)
			if err != nil {
				t.Fatal(err)
			}
			if mediaType != "multipart/form-data" || params["boundary"] == "" {
				t.Fatalf("Content-Type = %q", contentType)
			}
			form, err := multipart.NewReader(strings.NewReader(string(data)), params["boundary"]).ReadForm(1 << 20)
			if err != nil {
				t.Fatal(err)
			}
			defer form.RemoveAll()
		})
	}
}

func TestXProfileUpdateBannerInvalidParams(t *testing.T) {
	tests := map[string]struct {
		params  xtwitterscraper.XProfileUpdateBannerParams
		wantErr error
	}{
		"url and file": {
			params: xtwitterscraper.XProfileUpdateBannerParams{
				URL:  xtwitterscraper.String("https://example.com/banner.png"),
				File: strings.NewReader("image"),
			},
			wantErr: xtwitterscraper.ErrMediaHasURLAndFile,
		},
		"empty url and file": {
			params: xtwitterscraper.XProfileUpdateBannerParams{
				URL:  xtwitterscraper.String(""),
				File: strings.NewReader("image"),
			},
			wantErr: xtwitterscraper.ErrMediaHasURLAndFile,
		},
		"neither url nor file": {
			wantErr: xtwitterscraper.ErrMediaNeedsURLOrFile,
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
			_, err := client.X.Profile.UpdateBanner(context.Background(), test.params)
			if !errors.Is(err, test.wantErr) {
				t.Fatalf("error = %v, want %v", err, test.wantErr)
			}
			if sent {
				t.Fatal("invalid banner request was sent")
			}
		})
	}
}

func TestXProfileUpdateBannerFileReadError(t *testing.T) {
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
			_, err := client.X.Profile.UpdateBanner(context.Background(), xtwitterscraper.XProfileUpdateBannerParams{
				Account:        "@elonmusk",
				File:           file,
				IdempotencyKey: "upload-request",
			})
			if !errors.Is(err, wantErr) {
				t.Fatalf("error = %v, want %v", err, wantErr)
			}
			if sent {
				t.Fatal("banner request with unreadable file was sent")
			}
		})
	}
}
