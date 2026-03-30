// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package xtwitterscraper

import (
	"bytes"
	"context"
	"io"
	"mime/multipart"
	"net/http"
	"slices"

	"github.com/stainless-sdks/x-twitter-scraper-go/internal/apiform"
	"github.com/stainless-sdks/x-twitter-scraper-go/internal/apijson"
	"github.com/stainless-sdks/x-twitter-scraper-go/internal/requestconfig"
	"github.com/stainless-sdks/x-twitter-scraper-go/option"
	"github.com/stainless-sdks/x-twitter-scraper-go/packages/param"
	"github.com/stainless-sdks/x-twitter-scraper-go/packages/respjson"
)

// X write actions (tweets, likes, follows, DMs)
//
// XProfileService contains methods and other services that help with interacting
// with the x-twitter-scraper API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewXProfileService] method instead.
type XProfileService struct {
	options []option.RequestOption
}

// NewXProfileService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewXProfileService(opts ...option.RequestOption) (r XProfileService) {
	r = XProfileService{}
	r.options = opts
	return
}

// Update X profile
func (r *XProfileService) Update(ctx context.Context, body XProfileUpdateParams, opts ...option.RequestOption) (res *XProfileUpdateResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "x/profile"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Update profile avatar
func (r *XProfileService) UpdateAvatar(ctx context.Context, body XProfileUpdateAvatarParams, opts ...option.RequestOption) (res *XProfileUpdateAvatarResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "x/profile/avatar"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Update profile banner
func (r *XProfileService) UpdateBanner(ctx context.Context, body XProfileUpdateBannerParams, opts ...option.RequestOption) (res *XProfileUpdateBannerResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "x/profile/banner"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

type XProfileUpdateResponse struct {
	Success bool `json:"success" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XProfileUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *XProfileUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type XProfileUpdateAvatarResponse struct {
	Success bool `json:"success" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XProfileUpdateAvatarResponse) RawJSON() string { return r.JSON.raw }
func (r *XProfileUpdateAvatarResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type XProfileUpdateBannerResponse struct {
	Success bool `json:"success" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r XProfileUpdateBannerResponse) RawJSON() string { return r.JSON.raw }
func (r *XProfileUpdateBannerResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type XProfileUpdateParams struct {
	// X account (@username or account ID)
	Account string `json:"account" api:"required"`
	// Bio description
	Description param.Opt[string] `json:"description,omitzero"`
	Location    param.Opt[string] `json:"location,omitzero"`
	// Display name
	Name param.Opt[string] `json:"name,omitzero"`
	// Website URL
	URL param.Opt[string] `json:"url,omitzero"`
	paramObj
}

func (r XProfileUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow XProfileUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *XProfileUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type XProfileUpdateAvatarParams struct {
	// X account (@username or account ID)
	Account string `json:"account" api:"required"`
	// Avatar image (max 716KB)
	File io.Reader `json:"file,omitzero" api:"required" format:"binary"`
	paramObj
}

func (r XProfileUpdateAvatarParams) MarshalMultipart() (data []byte, contentType string, err error) {
	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)
	err = apiform.MarshalRoot(r, writer)
	if err == nil {
		err = apiform.WriteExtras(writer, r.ExtraFields())
	}
	if err != nil {
		writer.Close()
		return nil, "", err
	}
	err = writer.Close()
	if err != nil {
		return nil, "", err
	}
	return buf.Bytes(), writer.FormDataContentType(), nil
}

type XProfileUpdateBannerParams struct {
	// X account (@username or account ID)
	Account string `json:"account" api:"required"`
	// Banner image (max 2MB)
	File io.Reader `json:"file,omitzero" api:"required" format:"binary"`
	paramObj
}

func (r XProfileUpdateBannerParams) MarshalMultipart() (data []byte, contentType string, err error) {
	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)
	err = apiform.MarshalRoot(r, writer)
	if err == nil {
		err = apiform.WriteExtras(writer, r.ExtraFields())
	}
	if err != nil {
		writer.Close()
		return nil, "", err
	}
	err = writer.Close()
	if err != nil {
		return nil, "", err
	}
	return buf.Bytes(), writer.FormDataContentType(), nil
}
