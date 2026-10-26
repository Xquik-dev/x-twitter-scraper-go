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
	"github.com/Xquik-dev/x-twitter-scraper-go/packages/param"
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

type profileImageParams struct {
	URL  param.Opt[string]
	File io.Reader
}

type profileImageUpload struct {
	path    string
	url     string
	update  func(xtwitterscraper.Client, profileImageParams) error
	marshal func(profileImageParams) ([]byte, string, error)
}

var profileImageUploads = map[string]profileImageUpload{
	"avatar": {
		path: "/api/v1/x/profile/avatar",
		url:  "https://example.com/avatar.png",
		update: func(client xtwitterscraper.Client, params profileImageParams) error {
			_, err := client.X.Profile.UpdateAvatar(context.Background(), xtwitterscraper.XProfileUpdateAvatarParams{
				Account:        "@elonmusk",
				URL:            params.URL,
				File:           params.File,
				IdempotencyKey: "upload-request",
			})
			return err
		},
		marshal: func(params profileImageParams) ([]byte, string, error) {
			return xtwitterscraper.XProfileUpdateAvatarParams{Account: "@elonmusk", URL: params.URL, File: params.File}.MarshalMultipart()
		},
	},
	"banner": {
		path: "/api/v1/x/profile/banner",
		url:  "https://example.com/banner.png",
		update: func(client xtwitterscraper.Client, params profileImageParams) error {
			_, err := client.X.Profile.UpdateBanner(context.Background(), xtwitterscraper.XProfileUpdateBannerParams{
				Account:        "@elonmusk",
				URL:            params.URL,
				File:           params.File,
				IdempotencyKey: "upload-request",
			})
			return err
		},
		marshal: func(params profileImageParams) ([]byte, string, error) {
			return xtwitterscraper.XProfileUpdateBannerParams{Account: "@elonmusk", URL: params.URL, File: params.File}.MarshalMultipart()
		},
	},
}

type diskReader struct {
	io.Reader
	name string
}

func (r diskReader) Name() string { return r.name }

func profileImageClient(t *testing.T, upload profileImageUpload, check func(*http.Request)) (xtwitterscraper.Client, *bool) {
	sent := false
	client := xtwitterscraper.NewClient(
		option.WithBaseURL("https://example.com/api/v1/"),
		option.WithMaxRetries(0),
		option.WithHTTPClient(testHTTPDoer(func(req *http.Request) (*http.Response, error) {
			sent = true
			if check == nil {
				return nil, errors.New("unexpected request")
			}
			if req.Method != http.MethodPatch || req.URL.Path != upload.path {
				t.Fatalf("request = %s %s", req.Method, req.URL.Path)
			}
			if got := req.Header.Get("Idempotency-Key"); got != "upload-request" {
				t.Fatalf("Idempotency-Key = %q", got)
			}
			check(req)
			return &http.Response{
				StatusCode: http.StatusOK,
				Header:     http.Header{"Content-Type": {"application/json"}},
				Body:       io.NopCloser(strings.NewReader(`{}`)),
			}, nil
		})),
	)
	return client, &sent
}

func TestXProfileImageURLRequest(t *testing.T) {
	for name, upload := range profileImageUploads {
		t.Run(name, func(t *testing.T) {
			client, sent := profileImageClient(t, upload, func(req *http.Request) {
				if got := req.Header.Get("Content-Type"); got != "application/json" {
					t.Fatalf("Content-Type = %q", got)
				}
				var got map[string]string
				if err := json.NewDecoder(req.Body).Decode(&got); err != nil {
					t.Fatal(err)
				}
				want := map[string]string{"account": "@elonmusk", "url": upload.url}
				if !reflect.DeepEqual(got, want) {
					t.Fatalf("JSON body = %v, want %v", got, want)
				}
			})
			if err := upload.update(client, profileImageParams{URL: xtwitterscraper.String(upload.url)}); err != nil {
				t.Fatal(err)
			}
			if !*sent {
				t.Fatal("request was not sent")
			}
		})
	}
}

func TestXProfileImageFileRequest(t *testing.T) {
	tests := map[string]struct {
		file        func(content string) io.Reader
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
		"named file keeps its name": {
			file: func(content string) io.Reader {
				return xtwitterscraper.File(strings.NewReader(content), "me.png", "")
			},
			content:     "\x89PNG\r\n\x1a\nimage",
			filename:    "me.png",
			contentType: "image/png",
		},
		"disk file keeps its base name": {
			file: func(content string) io.Reader {
				return diskReader{Reader: strings.NewReader(content), name: "/tmp/photos/me.png"}
			},
			content:     "\x89PNG\r\n\x1a\nimage",
			filename:    "me.png",
			contentType: "image/png",
		},
		"file with an empty name gets the detected name": {
			file: func(content string) io.Reader {
				return diskReader{Reader: strings.NewReader(content)}
			},
			content:     "\x89PNG\r\n\x1a\nimage",
			filename:    "upload.png",
			contentType: "image/png",
		},
		"typed file keeps its name and type": {
			file: func(content string) io.Reader {
				return xtwitterscraper.File(strings.NewReader(content), "me.heic", "image/heic")
			},
			content:     "\x00\x00\x00\x18ftypheicimage",
			filename:    "me.heic",
			contentType: "image/heic",
		},
	}
	for name, upload := range profileImageUploads {
		for testName, test := range tests {
			t.Run(name+" "+testName, func(t *testing.T) {
				client, sent := profileImageClient(t, upload, func(req *http.Request) {
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
				})
				file := iotest.OneByteReader(strings.NewReader(test.content))
				if test.file != nil {
					file = test.file(test.content)
				}
				if err := upload.update(client, profileImageParams{File: file}); err != nil {
					t.Fatal(err)
				}
				if !*sent {
					t.Fatal("request was not sent")
				}
			})
		}
	}
}

func TestXProfileImageMarshalMultipart(t *testing.T) {
	tests := map[string]profileImageParams{
		"without file": {URL: xtwitterscraper.String("https://example.com/image.png")},
		"with file":    {File: strings.NewReader("\x89PNG\r\n\x1a\nimage")},
		"empty params": {},
	}
	for name, upload := range profileImageUploads {
		for testName, test := range tests {
			t.Run(name+" "+testName, func(t *testing.T) {
				data, contentType, err := upload.marshal(test)
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
}

func TestXProfileImageInvalidParams(t *testing.T) {
	tests := map[string]struct {
		params  profileImageParams
		wantErr error
	}{
		"url and file": {
			params:  profileImageParams{URL: xtwitterscraper.String("https://example.com/image.png"), File: strings.NewReader("image")},
			wantErr: xtwitterscraper.ErrMediaHasURLAndFile,
		},
		"empty url and file": {
			params:  profileImageParams{URL: xtwitterscraper.String(""), File: strings.NewReader("image")},
			wantErr: xtwitterscraper.ErrMediaHasURLAndFile,
		},
		"neither url nor file": {
			wantErr: xtwitterscraper.ErrMediaNeedsURLOrFile,
		},
	}
	for name, upload := range profileImageUploads {
		for testName, test := range tests {
			t.Run(name+" "+testName, func(t *testing.T) {
				client, sent := profileImageClient(t, upload, nil)
				if err := upload.update(client, test.params); !errors.Is(err, test.wantErr) {
					t.Fatalf("error = %v, want %v", err, test.wantErr)
				}
				if *sent {
					t.Fatal("invalid request was sent")
				}
			})
		}
	}
}

func TestXProfileImageFileReadError(t *testing.T) {
	wantErr := errors.New("file read failed")
	tests := map[string]func() io.Reader{
		"before detection": func() io.Reader { return iotest.ErrReader(wantErr) },
		"after detection": func() io.Reader {
			return io.MultiReader(strings.NewReader("\x89PNG\r\n\x1a\nimage data"), iotest.ErrReader(wantErr))
		},
	}
	for name, upload := range profileImageUploads {
		for testName, file := range tests {
			t.Run(name+" "+testName, func(t *testing.T) {
				client, sent := profileImageClient(t, upload, nil)
				if err := upload.update(client, profileImageParams{File: file()}); !errors.Is(err, wantErr) {
					t.Fatalf("error = %v, want %v", err, wantErr)
				}
				if *sent {
					t.Fatal("request with unreadable file was sent")
				}
			})
		}
	}
}
