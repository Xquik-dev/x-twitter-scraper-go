// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package xtwitterscraper

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/Xquik-dev/x-twitter-scraper-go/internal/requestconfig"
	"github.com/Xquik-dev/x-twitter-scraper-go/option"
	"github.com/Xquik-dev/x-twitter-scraper-go/packages/param"
)

// Support ticket management
//
// SupportAttachmentService contains methods and other services that help with
// interacting with the x-twitter-scraper API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSupportAttachmentService] method instead.
type SupportAttachmentService struct {
	options []option.RequestOption
}

// NewSupportAttachmentService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSupportAttachmentService(opts ...option.RequestOption) (r SupportAttachmentService) {
	r = SupportAttachmentService{}
	r.options = opts
	return
}

// Streams an authenticated user's support image or video. Video requests support
// one standard byte range for seeking and resumable playback.
func (r *SupportAttachmentService) Download(ctx context.Context, id string, query SupportAttachmentDownloadParams, opts ...option.RequestOption) (res *http.Response, err error) {
	if !param.IsOmitted(query.Range) {
		opts = append(opts, option.WithHeader("Range", fmt.Sprintf("%v", query.Range.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/octet-stream")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("support/attachments/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type SupportAttachmentDownloadParams struct {
	Range param.Opt[string] `header:"Range,omitzero" json:"-"`
	paramObj
}
