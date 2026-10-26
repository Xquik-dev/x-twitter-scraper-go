// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

// Package requestconfig builds, sends and retries the HTTP request behind each
// service method.
package requestconfig

import (
	"bytes"
	"context"
	"crypto/rand"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/big"
	"mime"
	"net/http"
	"net/url"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/Xquik-dev/x-twitter-scraper-go/internal"
	"github.com/Xquik-dev/x-twitter-scraper-go/internal/apierror"
	"github.com/Xquik-dev/x-twitter-scraper-go/internal/apiform"
	"github.com/Xquik-dev/x-twitter-scraper-go/internal/apiquery"
	"github.com/Xquik-dev/x-twitter-scraper-go/packages/param"
)

func getDefaultHeaders() map[string]string {
	return map[string]string{
		"User-Agent":  fmt.Sprintf("XTwitterScraper/Go %s", internal.PackageVersion),
		"X-Xquik-SDK": "go",
	}
}

func getNormalizedOS() string {
	return normalizeOS(runtime.GOOS)
}

func normalizeOS(goos string) string {
	switch goos {
	case "ios":
		return "iOS"
	case "android":
		return "Android"
	case "darwin":
		return "MacOS"
	case "windows":
		return "Windows"
	case "freebsd":
		return "FreeBSD"
	case "openbsd":
		return "OpenBSD"
	case "linux":
		return "Linux"
	default:
		return fmt.Sprintf("Other:%s", goos)
	}
}

func getNormalizedArchitecture() string {
	return normalizeArchitecture(runtime.GOARCH)
}

func normalizeArchitecture(goarch string) string {
	switch goarch {
	case "386":
		return "x32"
	case "amd64":
		return "x64"
	case "arm":
		return "arm"
	case "arm64":
		return "arm64"
	default:
		return fmt.Sprintf("other:%s", goarch)
	}
}

func getPlatformProperties() map[string]string {
	return map[string]string{
		"X-Stainless-Lang":            "go",
		"X-Stainless-Package-Version": internal.PackageVersion,
		"X-Stainless-OS":              getNormalizedOS(),
		"X-Stainless-Arch":            getNormalizedArchitecture(),
		"X-Stainless-Runtime":         "go",
		"X-Stainless-Runtime-Version": runtime.Version(),
	}
}

type RequestOption interface {
	Apply(*RequestConfig) error
}

type RequestOptionFunc func(*RequestConfig) error
type PreRequestOptionFunc func(*RequestConfig) error

func (s RequestOptionFunc) Apply(r *RequestConfig) error    { return s(r) }
func (s PreRequestOptionFunc) Apply(r *RequestConfig) error { return s(r) }

type optionalFiles interface {
	HasFiles() bool
}

func sendsMultipart(body apiform.Marshaler) bool {
	files, ok := body.(optionalFiles)
	return !ok || files.HasFiles()
}

func NewRequestConfig[D any](ctx context.Context, method string, u string, body param.ParamStruct, dst D, opts ...RequestOption) (*RequestConfig, error) {
	var reader io.Reader

	contentType := "application/json"
	hasSerializationFunc := false

	if multipartBody, ok := body.(apiform.Marshaler); ok && sendsMultipart(multipartBody) {
		var (
			content []byte
			err     error
		)
		content, contentType, err = multipartBody.MarshalMultipart()
		if err != nil {
			return nil, err
		}
		reader = bytes.NewBuffer(content)
		hasSerializationFunc = true
	} else if body, ok := body.(json.Marshaler); ok {
		content, err := body.MarshalJSON()
		if err != nil {
			return nil, err
		}
		reader = bytes.NewBuffer(content)
		hasSerializationFunc = true
	}
	if body, ok := body.(apiquery.Queryer); ok {
		hasSerializationFunc = true
		q, err := body.URLQuery()
		if err != nil {
			return nil, err
		}
		params := q.Encode()
		if params != "" {
			parsed, err := url.Parse(u)
			if err != nil {
				return nil, err
			}
			if parsed.RawQuery != "" {
				parsed.RawQuery = parsed.RawQuery + "&" + params
				u = parsed.String()
			} else {
				u = u + "?" + params
			}
		}
	}
	if body != nil && !hasSerializationFunc {
		buf := new(bytes.Buffer)
		enc := json.NewEncoder(buf)
		enc.SetEscapeHTML(true)
		if err := enc.Encode(body); err != nil {
			return nil, err
		}
		reader = buf
	}

	req, err := http.NewRequestWithContext(ctx, method, u, nil)
	if err != nil {
		return nil, err
	}
	if reader != nil {
		req.Header.Set("Content-Type", contentType)
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("X-Stainless-Retry-Count", "0")
	req.Header.Set("X-Stainless-Timeout", "0")
	for k, v := range getDefaultHeaders() {
		req.Header.Add(k, v)
	}

	for k, v := range getPlatformProperties() {
		req.Header.Add(k, v)
	}
	cfg := RequestConfig{
		MaxRetries: 2,
		Context:    ctx,
		Request:    req,
		HTTPClient: http.DefaultClient,
		Body:       reader,
	}
	cfg.ResponseBodyInto = Into(dst)
	cfg.Security = Security{
		APIKey:      true,
		OAuthBearer: true,
	}
	err = cfg.Apply(opts...)
	if err != nil {
		return nil, err
	}

	ApplySecurity(cfg)

	if req.Header.Get("X-Stainless-Timeout") == "0" {
		if cfg.RequestTimeout == time.Duration(0) {
			req.Header.Del("X-Stainless-Timeout")
		} else {
			req.Header.Set("X-Stainless-Timeout", strconv.Itoa(int(cfg.RequestTimeout.Seconds())))
		}
	}

	return &cfg, nil
}

type HTTPDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

type RequestConfig struct {
	MaxRetries       int
	RequestTimeout   time.Duration
	Context          context.Context
	Request          *http.Request
	BaseURL          *url.URL
	DefaultBaseURL   *url.URL
	CustomHTTPDoer   HTTPDoer
	HTTPClient       *http.Client
	Middlewares      []middleware
	APIKey           string
	BearerToken      string
	Security         Security
	ResponseBodyInto ResponseBody
	ResponseInto     **http.Response
	Body             io.Reader
}

type middleware = func(*http.Request, middlewareNext) (*http.Response, error)

type middlewareNext = func(*http.Request) (*http.Response, error)

func applyMiddleware(middleware middleware, next middlewareNext) middlewareNext {
	return func(req *http.Request) (res *http.Response, err error) {
		return middleware(req, next)
	}
}

func shouldRetry(req *http.Request, res *http.Response) bool {
	if req.Body != nil && req.GetBody == nil {
		return false
	}

	if res == nil {
		return true
	}

	if res.Header.Get("x-should-retry") == "true" {
		return true
	}
	if res.Header.Get("x-should-retry") == "false" {
		return false
	}

	return res.StatusCode == http.StatusRequestTimeout ||
		res.StatusCode == http.StatusConflict ||
		res.StatusCode == http.StatusTooManyRequests ||
		res.StatusCode >= http.StatusInternalServerError
}

func parseRetryAfterHeader(resp *http.Response) (time.Duration, bool) {
	if resp == nil {
		return 0, false
	}
	if milliseconds, err := strconv.ParseFloat(resp.Header.Get("Retry-After-Ms"), 64); err == nil {
		return time.Duration(milliseconds * float64(time.Millisecond)), true
	}
	value := resp.Header.Get("Retry-After")
	if seconds, err := strconv.ParseFloat(value, 64); err == nil {
		return time.Duration(seconds * float64(time.Second)), true
	}
	date, err := time.Parse(time.RFC1123, value)
	if err != nil {
		return 0, false
	}
	return time.Until(date), true
}

func isBeforeContextDeadline(t time.Time, ctx context.Context) bool {
	d, ok := ctx.Deadline()
	return !ok || t.Before(d)
}

type bodyWithTimeout struct {
	stop func()
	rc   io.ReadCloser
}

func (b *bodyWithTimeout) Read(p []byte) (n int, err error) {
	return b.rc.Read(p)
}

func (b *bodyWithTimeout) Close() error {
	err := b.rc.Close()
	b.stop()
	return err
}

func retryDelay(res *http.Response, retryCount int) time.Duration {
	if retryAfterDelay, ok := parseRetryAfterHeader(res); ok {
		return max(0, retryAfterDelay)
	}

	delay := 500 * time.Millisecond << min(max(retryCount, 0), 4)
	jitter, err := rand.Int(rand.Reader, big.NewInt(int64(delay/4)))
	if err != nil {
		return delay
	}
	return delay - time.Duration(jitter.Int64())
}

func (cfg *RequestConfig) Execute() (err error) {
	if cfg.BaseURL == nil {
		if cfg.DefaultBaseURL != nil {
			cfg.BaseURL = cfg.DefaultBaseURL
		} else {
			return fmt.Errorf("requestconfig: base url is not set")
		}
	}

	cfg.Request.URL, err = cfg.BaseURL.Parse(strings.TrimLeft(cfg.Request.URL.String(), "/"))
	if err != nil {
		return err
	}

	if cfg.Body != nil && cfg.Request.Body == nil {
		bodyRequest, err := http.NewRequestWithContext(cfg.Request.Context(), cfg.Request.Method, cfg.Request.URL.String(), cfg.Body)
		if err != nil {
			return err
		}
		cfg.Request.Body = bodyRequest.Body
		if bodyRequest.GetBody != nil {
			cfg.Request.GetBody = bodyRequest.GetBody
			cfg.Request.ContentLength = bodyRequest.ContentLength
		}
	}

	handler := cfg.HTTPClient.Do
	if cfg.CustomHTTPDoer != nil {
		handler = cfg.CustomHTTPDoer.Do
	}
	for i := len(cfg.Middlewares) - 1; i >= 0; i -= 1 {
		handler = applyMiddleware(cfg.Middlewares[i], handler)
	}

	shouldSendRetryCount := cfg.Request.Header.Get("X-Stainless-Retry-Count") == "0"

	var res *http.Response
	var cancel context.CancelFunc
	for retryCount := 0; retryCount <= cfg.MaxRetries; retryCount += 1 {
		ctx := cfg.Request.Context()
		if cfg.RequestTimeout != time.Duration(0) && isBeforeContextDeadline(time.Now().Add(cfg.RequestTimeout), ctx) {
			ctx, cancel = context.WithTimeout(ctx, cfg.RequestTimeout)
			defer func() {
				if cancel != nil {
					cancel()
				}
			}()
		}

		req := cfg.Request.Clone(ctx)
		if shouldSendRetryCount {
			req.Header.Set("X-Stainless-Retry-Count", strconv.Itoa(retryCount))
		}

		res, err = handler(req)
		if ctx != nil && ctx.Err() != nil {
			return ctx.Err()
		}
		if !shouldRetry(cfg.Request, res) || retryCount >= cfg.MaxRetries {
			break
		}

		if cfg.Request.GetBody != nil {
			cfg.Request.Body, err = cfg.Request.GetBody()
			if err != nil {
				return err
			}
		}

		if cfg.Request.GetBody == nil && cfg.Request.Body != nil {
			break
		}

		if res != nil && res.Body != nil {
			if err := res.Body.Close(); err != nil {
				return err
			}
		}

		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(retryDelay(res, retryCount)):
		}
	}

	if cfg.ResponseInto != nil {
		*cfg.ResponseInto = res
	}
	if cfg.ResponseBodyInto.raw != nil {
		*cfg.ResponseBodyInto.raw = res
	}

	if err != nil {
		return err
	}

	if res.StatusCode >= 400 {
		contents, err := io.ReadAll(res.Body)
		err = errors.Join(err, res.Body.Close())
		if err != nil {
			return err
		}

		res.Body = io.NopCloser(bytes.NewBuffer(contents))

		aerr := apierror.Error{Request: cfg.Request, Response: res, StatusCode: res.StatusCode}
		err = aerr.UnmarshalJSON(contents)
		if err != nil {
			return err
		}
		return &aerr
	}

	if cfg.ResponseBodyInto.decode == nil {
		if cancel != nil {
			res.Body = &bodyWithTimeout{rc: res.Body, stop: cancel}
			cancel = nil
		}
		return nil
	}

	contents, err := io.ReadAll(res.Body)
	err = errors.Join(err, res.Body.Close())
	if err != nil {
		return fmt.Errorf("error reading response body: %w", err)
	}

	return cfg.ResponseBodyInto.decode(res.Header.Get("content-type"), contents)
}

type ResponseBody struct {
	raw    **http.Response
	decode func(contentType string, contents []byte) error
}

func Into[D any](dst D) ResponseBody {
	value := reflect.ValueOf(dst)
	if !value.IsValid() {
		return ResponseBody{}
	}
	if body, ok := reflect.TypeAssert[ResponseBody](value); ok {
		return body
	}
	if raw, ok := reflect.TypeAssert[**http.Response](value); ok {
		return ResponseBody{raw: raw}
	}
	return ResponseBody{decode: func(contentType string, contents []byte) error {
		return decodeBody(dst, contentType, contents)
	}}
}

func decodeBody[D any](dst D, contentType string, contents []byte) error {
	value := reflect.ValueOf(dst)
	if bytes, ok := reflect.TypeAssert[*[]byte](value); ok {
		*bytes = contents
		return nil
	}
	mediaType, _, err := mime.ParseMediaType(contentType)
	if (err != nil && !errors.Is(err, mime.ErrInvalidMediaParameter)) || !strings.Contains(mediaType, "application/json") && !strings.HasSuffix(mediaType, "+json") {
		if text, ok := reflect.TypeAssert[*string](value); ok {
			*text = string(contents)
			return nil
		}
		if text, ok := reflect.TypeAssert[**string](value); ok {
			body := string(contents)
			*text = &body
			return nil
		}
		return fmt.Errorf("expected destination type of 'string' or '[]byte' for responses with content-type '%s' that is not 'application/json'", contentType)
	}
	if err := json.NewDecoder(bytes.NewReader(contents)).Decode(dst); err != nil {
		return fmt.Errorf("error parsing response json: %w", err)
	}
	return nil
}

func ExecuteNewRequest[D any](ctx context.Context, method string, u string, body param.ParamStruct, dst D, opts ...RequestOption) error {
	cfg, err := NewRequestConfig(ctx, method, u, body, dst, opts...)
	if err != nil {
		return err
	}
	return cfg.Execute()
}

func (cfg *RequestConfig) Clone(ctx context.Context) *RequestConfig {
	if cfg == nil {
		return nil
	}
	req := cfg.Request.Clone(ctx)
	var err error
	if req.Body != nil {
		req.Body, err = req.GetBody()
	}
	if err != nil {
		return nil
	}
	new := &RequestConfig{
		MaxRetries:     cfg.MaxRetries,
		RequestTimeout: cfg.RequestTimeout,
		Context:        ctx,
		Request:        req,
		BaseURL:        cfg.BaseURL,
		HTTPClient:     cfg.HTTPClient,
		Middlewares:    cfg.Middlewares,
		APIKey:         cfg.APIKey,
		BearerToken:    cfg.BearerToken,
	}

	return new
}

func (cfg *RequestConfig) Apply(opts ...RequestOption) error {
	for _, opt := range opts {
		err := opt.Apply(cfg)
		if err != nil {
			return err
		}
	}
	return nil
}

func PreRequestOptions(opts ...RequestOption) (RequestConfig, error) {
	cfg := RequestConfig{}
	for _, opt := range opts {
		if opt, ok := opt.(PreRequestOptionFunc); ok {
			err := opt.Apply(&cfg)
			if err != nil {
				return cfg, err
			}
		}
	}
	return cfg, nil
}

func WithDefaultBaseURL(baseURL string) RequestOption {
	u, err := url.Parse(baseURL)
	return RequestOptionFunc(func(r *RequestConfig) error {
		if err != nil {
			return err
		}
		r.DefaultBaseURL = u
		return nil
	})
}

type Security struct {
	APIKey      bool
	OAuthBearer bool
}

func WithSecurity(security Security) RequestOption {
	return RequestOptionFunc(func(r *RequestConfig) error {
		r.Security = security
		return nil
	})
}

func WithAPIKeySecurity() RequestOption {
	return RequestOptionFunc(func(r *RequestConfig) error {
		r.Security = Security{
			APIKey:      true,
			OAuthBearer: false,
		}
		return nil
	})
}

func WithOAuthBearerSecurity() RequestOption {
	return RequestOptionFunc(func(r *RequestConfig) error {
		r.Security = Security{
			APIKey:      false,
			OAuthBearer: true,
		}
		return nil
	})
}

func ApplySecurity(r RequestConfig) {
	if r.Security.APIKey && r.APIKey != "" && r.Request.Header.Get("x-api-key") == "" {
		r.Request.Header.Set("x-api-key", r.APIKey)
	}

	if r.Security.OAuthBearer && r.BearerToken != "" && r.Request.Header.Get("Authorization") == "" {
		r.Request.Header.Set("authorization", fmt.Sprintf("Bearer %s", r.BearerToken))
	}
}
