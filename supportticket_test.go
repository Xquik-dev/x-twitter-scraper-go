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

	"github.com/Xquik-dev/x-twitter-scraper-go"
	"github.com/Xquik-dev/x-twitter-scraper-go/internal/testutil"
	"github.com/Xquik-dev/x-twitter-scraper-go/option"
	"github.com/Xquik-dev/x-twitter-scraper-go/packages/param"
)

func TestSupportTicketRequestBodies(t *testing.T) {
	services := []struct {
		name    string
		path    string
		subject string
		send    func(xtwitterscraper.Client, param.Opt[string], []io.Reader) error
	}{
		{
			name:    "ticket",
			path:    "/api/v1/support/tickets",
			subject: "Cannot connect X account",
			send: func(client xtwitterscraper.Client, body param.Opt[string], attachments []io.Reader) error {
				_, err := client.Support.Tickets.New(context.Background(), xtwitterscraper.SupportTicketNewParams{
					Subject:        "Cannot connect X account",
					Body:           body,
					Attachments:    attachments,
					IdempotencyKey: xtwitterscraper.String("support-request"),
				})
				return err
			},
		},
		{
			name: "reply",
			path: "/api/v1/support/tickets/tkt_example/messages",
			send: func(client xtwitterscraper.Client, body param.Opt[string], attachments []io.Reader) error {
				_, err := client.Support.Tickets.Reply(context.Background(), "tkt_example", xtwitterscraper.SupportTicketReplyParams{
					Body:           body,
					Attachments:    attachments,
					IdempotencyKey: xtwitterscraper.String("support-request"),
				})
				return err
			},
		},
	}
	for _, service := range services {
		t.Run(service.name, func(t *testing.T) {
			tests := map[string]struct {
				body        param.Opt[string]
				attachments []io.Reader
				wantJSON    bool
				wantBody    []string
				wantFiles   []string
				filename    string
				contentType string
				wantErr     error
			}{
				"neither body nor attachments": {
					wantErr: xtwitterscraper.ErrTicketNeedsBodyOrAttachments,
				},
				"empty attachments": {
					attachments: []io.Reader{},
					wantErr:     xtwitterscraper.ErrTicketNeedsBodyOrAttachments,
				},
				"empty body": {
					body:    xtwitterscraper.String(""),
					wantErr: xtwitterscraper.ErrTicketNeedsBodyOrAttachments,
				},
				"empty body and empty attachments": {
					body:        xtwitterscraper.String(""),
					attachments: []io.Reader{},
					wantErr:     xtwitterscraper.ErrTicketNeedsBodyOrAttachments,
				},
				"null body": {
					body:    param.Null[string](),
					wantErr: xtwitterscraper.ErrTicketNeedsBodyOrAttachments,
				},
				"attachments only": {
					attachments: []io.Reader{strings.NewReader("first"), strings.NewReader("second"), strings.NewReader("third"), strings.NewReader("fourth")},
					wantFiles:   []string{"first", "second", "third", "fourth"},
				},
				"body only": {
					body:     xtwitterscraper.String("Please help."),
					wantJSON: true,
					wantBody: []string{"Please help."},
				},
				"body and empty attachments": {
					body:        xtwitterscraper.String("Please help."),
					attachments: []io.Reader{},
					wantJSON:    true,
					wantBody:    []string{"Please help."},
				},
				"body and attachments": {
					body:        xtwitterscraper.String("Please help."),
					attachments: []io.Reader{strings.NewReader("screenshot")},
					wantBody:    []string{"Please help."},
					wantFiles:   []string{"screenshot"},
				},
				"named attachment": {
					attachments: []io.Reader{xtwitterscraper.File(strings.NewReader("image"), "screenshot.png", "image/png")},
					wantFiles:   []string{"image"},
					filename:    "screenshot.png",
					contentType: "image/png",
				},
				"empty body and attachments": {
					body:        xtwitterscraper.String(""),
					attachments: []io.Reader{strings.NewReader("screenshot")},
					wantBody:    []string{""},
					wantFiles:   []string{"screenshot"},
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
							if test.wantErr != nil {
								return nil, errors.New("unexpected request")
							}
							if req.Method != http.MethodPost || req.URL.Path != service.path {
								t.Fatalf("request = %s %s", req.Method, req.URL.Path)
							}
							if got := req.Header.Get("Idempotency-Key"); got != "support-request" {
								t.Fatalf("Idempotency-Key = %q", got)
							}
							mediaType, params, err := mime.ParseMediaType(req.Header.Get("Content-Type"))
							if err != nil {
								t.Fatal(err)
							}
							if test.wantJSON {
								if mediaType != "application/json" {
									t.Fatalf("Content-Type = %q", mediaType)
								}
								var got map[string]any
								if err := json.NewDecoder(req.Body).Decode(&got); err != nil {
									t.Fatal(err)
								}
								want := map[string]any{"body": "Please help."}
								if test.attachments != nil {
									want["attachments"] = []any{}
								}
								if service.subject != "" {
									want["subject"] = service.subject
								}
								if !reflect.DeepEqual(got, want) {
									t.Fatalf("JSON body = %v, want %v", got, want)
								}
							} else {
								if mediaType != "multipart/form-data" {
									t.Fatalf("Content-Type = %q", mediaType)
								}
								form, err := multipart.NewReader(req.Body, params["boundary"]).ReadForm(1 << 20)
								if err != nil {
									t.Fatal(err)
								}
								defer form.RemoveAll()
								want := map[string][]string{}
								if test.wantBody != nil {
									want["body"] = test.wantBody
								}
								if service.subject != "" {
									want["subject"] = []string{service.subject}
								}
								if !reflect.DeepEqual(form.Value, want) {
									t.Fatalf("form fields = %v, want %v", form.Value, want)
								}
								files := form.File["attachments"]
								if len(form.File) != 1 || len(files) != len(test.wantFiles) {
									t.Fatalf("form files = %v", form.File)
								}
								for i, file := range files {
									filename := test.filename
									if filename == "" {
										filename = "anonymous_file"
									}
									contentType := test.contentType
									if contentType == "" {
										contentType = "application/octet-stream"
									}
									if file.Filename != filename || file.Header.Get("Content-Type") != contentType {
										t.Fatalf("attachment headers = %v", file.Header)
									}
									reader, err := file.Open()
									if err != nil {
										t.Fatal(err)
									}
									defer reader.Close()
									content, err := io.ReadAll(reader)
									if err != nil || string(content) != test.wantFiles[i] {
										t.Fatalf("attachment = %q, error = %v", content, err)
									}
								}
							}
							return &http.Response{
								StatusCode: http.StatusOK,
								Header:     http.Header{"Content-Type": {"application/json"}},
								Body:       io.NopCloser(strings.NewReader(`{"publicId":"tkt_example","attachments":[]}`)),
							}, nil
						})),
					)
					err := service.send(client, test.body, test.attachments)
					if test.wantErr != nil {
						if !errors.Is(err, test.wantErr) {
							t.Fatalf("error = %v, want %v", err, test.wantErr)
						}
						if sent {
							t.Fatal("invalid support request was sent")
						}
						return
					}
					if err != nil {
						t.Fatal(err)
					}
					if !sent {
						t.Fatal("support request was not sent")
					}
				})
			}
		})
	}
}

func TestSupportTicketMarshalMultipart(t *testing.T) {
	tests := map[string]func() ([]byte, string, error){
		"ticket without attachments": xtwitterscraper.SupportTicketNewParams{
			Subject: "Cannot connect X account",
			Body:    xtwitterscraper.String("Please help."),
		}.MarshalMultipart,
		"ticket with empty attachments": xtwitterscraper.SupportTicketNewParams{
			Subject:     "Cannot connect X account",
			Body:        xtwitterscraper.String("Please help."),
			Attachments: []io.Reader{},
		}.MarshalMultipart,
		"ticket with attachments": xtwitterscraper.SupportTicketNewParams{
			Subject:     "Cannot connect X account",
			Attachments: []io.Reader{strings.NewReader("screenshot")},
		}.MarshalMultipart,
		"reply without attachments": xtwitterscraper.SupportTicketReplyParams{
			Body: xtwitterscraper.String("Please help."),
		}.MarshalMultipart,
		"reply with empty attachments": xtwitterscraper.SupportTicketReplyParams{
			Body:        xtwitterscraper.String("Please help."),
			Attachments: []io.Reader{},
		}.MarshalMultipart,
		"reply with attachments": xtwitterscraper.SupportTicketReplyParams{
			Attachments: []io.Reader{strings.NewReader("screenshot")},
		}.MarshalMultipart,
	}
	for name, marshal := range tests {
		t.Run(name, func(t *testing.T) {
			data, contentType, err := marshal()
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

func TestSupportTicketNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Support.Tickets.New(context.TODO(), xtwitterscraper.SupportTicketNewParams{
		Body:           xtwitterscraper.String("I am unable to connect my X account. Please help."),
		Subject:        "Cannot connect X account",
		IdempotencyKey: xtwitterscraper.String("Idempotency-Key"),
	})
	if err != nil {
		var apierr *xtwitterscraper.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSupportTicketGet(t *testing.T) {
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
	_, err := client.Support.Tickets.Get(context.TODO(), "tkt_a1b2c3d4e5f6a1b2c3d4e5f6")
	if err != nil {
		var apierr *xtwitterscraper.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSupportTicketUpdate(t *testing.T) {
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
	_, err := client.Support.Tickets.Update(
		context.TODO(),
		"tkt_a1b2c3d4e5f6a1b2c3d4e5f6",
		xtwitterscraper.SupportTicketUpdateParams{
			Status: xtwitterscraper.SupportTicketUpdateParamsStatusResolved,
		},
	)
	if err != nil {
		var apierr *xtwitterscraper.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSupportTicketList(t *testing.T) {
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
	_, err := client.Support.Tickets.List(context.TODO())
	if err != nil {
		var apierr *xtwitterscraper.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSupportTicketReplyWithOptionalParams(t *testing.T) {
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
	_, err := client.Support.Tickets.Reply(
		context.TODO(),
		"tkt_a1b2c3d4e5f6a1b2c3d4e5f6",
		xtwitterscraper.SupportTicketReplyParams{
			Body:           xtwitterscraper.String("Thank you for the update."),
			IdempotencyKey: xtwitterscraper.String("Idempotency-Key"),
		},
	)
	if err != nil {
		var apierr *xtwitterscraper.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
