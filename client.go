// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

package xtwitterscraper

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"slices"
	"strings"

	"github.com/Xquik-dev/x-twitter-scraper-go/internal/requestconfig"
	"github.com/Xquik-dev/x-twitter-scraper-go/option"
	"github.com/Xquik-dev/x-twitter-scraper-go/packages/param"
)

// Client creates a struct with services and top level methods that help with
// interacting with the x-twitter-scraper API. You should not instantiate this
// client directly, and instead use the [NewClient] method instead.
type Client struct {
	options []option.RequestOption
	// Account info and settings
	Account AccountService
	// Subscription, billing, and credits
	Subscribe SubscribeService
	// AI tweet composition, drafts, writing styles, and radar
	Compose ComposeService
	// AI tweet composition, drafts, writing styles, and radar
	Drafts DraftService
	// AI tweet composition, drafts, writing styles, and radar
	Styles StyleService
	// AI tweet composition, drafts, writing styles, and radar
	Radar RadarService
	// X account monitoring with 1-second checks
	Monitors MonitorService
	// Activity events from monitored accounts
	Events EventService
	// Saved or bulk data extraction (23 tool types)
	Extractions ExtractionService
	// Giveaway draws from tweet replies
	Draws DrawService
	// Webhook endpoint management and delivery
	Webhooks WebhookService
	X        XService
	// Trending topics and hashtags by region
	Trends  TrendService
	Support SupportService
	// Subscription, billing, and credits
	Credits CreditService
	// Accountless prepaid access for paid read endpoints
	GuestWallets GuestWalletService
}

// DefaultClientOptions read from the environment (X_TWITTER_SCRAPER_API_KEY,
// X_TWITTER_SCRAPER_BEARER_TOKEN, X_TWITTER_SCRAPER_BASE_URL). This should be used
// to initialize new clients.
func DefaultClientOptions() []option.RequestOption {
	defaults := []option.RequestOption{option.WithHTTPClient(defaultHTTPClient()), option.WithEnvironmentProduction()}
	if o, ok := os.LookupEnv("X_TWITTER_SCRAPER_BASE_URL"); ok {
		defaults = append(defaults, option.WithBaseURL(o))
	}
	if o, ok := os.LookupEnv("X_TWITTER_SCRAPER_API_KEY"); ok {
		defaults = append(defaults, option.WithAPIKey(o))
	}
	if o, ok := os.LookupEnv("X_TWITTER_SCRAPER_BEARER_TOKEN"); ok {
		defaults = append(defaults, option.WithBearerToken(o))
	}
	if o, ok := os.LookupEnv("X_TWITTER_SCRAPER_CUSTOM_HEADERS"); ok {
		for _, line := range strings.Split(o, "\n") {
			colon := strings.Index(line, ":")
			if colon >= 0 {
				defaults = append(defaults, option.WithHeader(strings.TrimSpace(line[:colon]), strings.TrimSpace(line[colon+1:])))
			}
		}
	}
	return defaults
}

// NewClient generates a new client with the default option read from the
// environment (X_TWITTER_SCRAPER_API_KEY, X_TWITTER_SCRAPER_BEARER_TOKEN,
// X_TWITTER_SCRAPER_BASE_URL). The option passed in as arguments are applied after
// these default arguments, and all option will be passed down to the services and
// requests that this client makes.
func NewClient(opts ...option.RequestOption) (r Client) {
	opts = append(DefaultClientOptions(), opts...)

	r = Client{options: opts}

	r.Account = NewAccountService(opts...)
	r.Subscribe = NewSubscribeService(opts...)
	r.Compose = NewComposeService(opts...)
	r.Drafts = NewDraftService(opts...)
	r.Styles = NewStyleService(opts...)
	r.Radar = NewRadarService(opts...)
	r.Monitors = NewMonitorService(opts...)
	r.Events = NewEventService(opts...)
	r.Extractions = NewExtractionService(opts...)
	r.Draws = NewDrawService(opts...)
	r.Webhooks = NewWebhookService(opts...)
	r.X = NewXService(opts...)
	r.Trends = NewTrendService(opts...)
	r.Support = NewSupportService(opts...)
	r.Credits = NewCreditService(opts...)
	r.GuestWallets = NewGuestWalletService(opts...)

	return
}

// Execute makes a request with the given context, method, URL, request params,
// response, and request options. This is useful for hitting undocumented endpoints
// while retaining the base URL, auth, retries, and other options from the client.
//
// Any "…Params" object defined in this library can be used as the request
// argument. Note that 'path' arguments will not be forwarded into the url. A
// params struct of your own embeds [param.APIObject]. It is serialized into the
// body using [encoding/json], or with its MarshalJSON method when it has one. A
// URLQuery method adds the returned [url.Values] to the query string, and a
// MarshalMultipart method sends a multipart form. To send an [io.Reader] as is,
// such as a [bytes.Reader] over a byte slice, pass nil params and
// [option.WithRequestBody].
//
// Any model defined in this library can be used as the response argument, and
// the response body is deserialized into it with its UnmarshalJSON method. A
// [*json.RawMessage] receives the JSON body as is. A nil response leaves the
// body unread. For other destinations, such as a [*http.Response] or a byte
// slice, pass nil and [option.WithResponseBodyInto] or [option.WithResponseInto].
func (r *Client) Execute(ctx context.Context, method string, path string, params param.ParamStruct, res json.Unmarshaler, opts ...option.RequestOption) error {
	opts = slices.Concat(r.options, opts)
	return requestconfig.ExecuteNewRequest(ctx, method, path, params, res, opts...)
}

// Get makes a GET request with the given URL, params, and optionally deserializes
// to a response. See [Execute] documentation on the params and response.
func (r *Client) Get(ctx context.Context, path string, params param.ParamStruct, res json.Unmarshaler, opts ...option.RequestOption) error {
	return r.Execute(ctx, http.MethodGet, path, params, res, opts...)
}

// Post makes a POST request with the given URL, params, and optionally
// deserializes to a response. See [Execute] documentation on the params and
// response.
func (r *Client) Post(ctx context.Context, path string, params param.ParamStruct, res json.Unmarshaler, opts ...option.RequestOption) error {
	return r.Execute(ctx, http.MethodPost, path, params, res, opts...)
}

// Put makes a PUT request with the given URL, params, and optionally deserializes
// to a response. See [Execute] documentation on the params and response.
func (r *Client) Put(ctx context.Context, path string, params param.ParamStruct, res json.Unmarshaler, opts ...option.RequestOption) error {
	return r.Execute(ctx, http.MethodPut, path, params, res, opts...)
}

// Patch makes a PATCH request with the given URL, params, and optionally
// deserializes to a response. See [Execute] documentation on the params and
// response.
func (r *Client) Patch(ctx context.Context, path string, params param.ParamStruct, res json.Unmarshaler, opts ...option.RequestOption) error {
	return r.Execute(ctx, http.MethodPatch, path, params, res, opts...)
}

// Delete makes a DELETE request with the given URL, params, and optionally
// deserializes to a response. See [Execute] documentation on the params and
// response.
func (r *Client) Delete(ctx context.Context, path string, params param.ParamStruct, res json.Unmarshaler, opts ...option.RequestOption) error {
	return r.Execute(ctx, http.MethodDelete, path, params, res, opts...)
}
