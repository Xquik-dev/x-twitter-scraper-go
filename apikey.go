// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package xtwitterscraper

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/stainless-sdks/x-twitter-scraper-go/internal/apijson"
	"github.com/stainless-sdks/x-twitter-scraper-go/internal/requestconfig"
	"github.com/stainless-sdks/x-twitter-scraper-go/option"
	"github.com/stainless-sdks/x-twitter-scraper-go/packages/param"
	"github.com/stainless-sdks/x-twitter-scraper-go/packages/respjson"
)

// API key management (session auth only)
//
// APIKeyService contains methods and other services that help with interacting
// with the x-twitter-scraper API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAPIKeyService] method instead.
type APIKeyService struct {
	options []option.RequestOption
}

// NewAPIKeyService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAPIKeyService(opts ...option.RequestOption) (r APIKeyService) {
	r = APIKeyService{}
	r.options = opts
	return
}

// Create API key
func (r *APIKeyService) New(ctx context.Context, body APIKeyNewParams, opts ...option.RequestOption) (res *APIKeyNewResponse, err error) {
	var preClientOpts = []option.RequestOption{requestconfig.WithSecurity(requestconfig.Security{})}
	opts = slices.Concat(preClientOpts, r.options, opts)
	path := "api-keys"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// List API keys
func (r *APIKeyService) List(ctx context.Context, opts ...option.RequestOption) (res *APIKeyListResponse, err error) {
	var preClientOpts = []option.RequestOption{requestconfig.WithSecurity(requestconfig.Security{})}
	opts = slices.Concat(preClientOpts, r.options, opts)
	path := "api-keys"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Revoke API key
func (r *APIKeyService) Revoke(ctx context.Context, id string, opts ...option.RequestOption) (res *APIKeyRevokeResponse, err error) {
	var preClientOpts = []option.RequestOption{requestconfig.WithSecurity(requestconfig.Security{})}
	opts = slices.Concat(preClientOpts, r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api-keys/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

type APIKey struct {
	ID         string    `json:"id" api:"required"`
	CreatedAt  time.Time `json:"createdAt" api:"required" format:"date-time"`
	IsActive   bool      `json:"isActive" api:"required"`
	Name       string    `json:"name" api:"required"`
	Prefix     string    `json:"prefix" api:"required"`
	LastUsedAt time.Time `json:"lastUsedAt" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		IsActive    respjson.Field
		Name        respjson.Field
		Prefix      respjson.Field
		LastUsedAt  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r APIKey) RawJSON() string { return r.JSON.raw }
func (r *APIKey) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APIKeyNewResponse struct {
	ID        string    `json:"id" api:"required"`
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	FullKey   string    `json:"fullKey" api:"required"`
	Name      string    `json:"name" api:"required"`
	Prefix    string    `json:"prefix" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		FullKey     respjson.Field
		Name        respjson.Field
		Prefix      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r APIKeyNewResponse) RawJSON() string { return r.JSON.raw }
func (r *APIKeyNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APIKeyListResponse struct {
	Keys []APIKey `json:"keys" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Keys        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r APIKeyListResponse) RawJSON() string { return r.JSON.raw }
func (r *APIKeyListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APIKeyRevokeResponse struct {
	Success bool `json:"success" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r APIKeyRevokeResponse) RawJSON() string { return r.JSON.raw }
func (r *APIKeyRevokeResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APIKeyNewParams struct {
	Name param.Opt[string] `json:"name,omitzero"`
	paramObj
}

func (r APIKeyNewParams) MarshalJSON() (data []byte, err error) {
	type shadow APIKeyNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *APIKeyNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
