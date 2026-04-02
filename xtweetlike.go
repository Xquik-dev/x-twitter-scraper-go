// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package xtwitterscraper

import (
	"github.com/Xquik-dev/x-twitter-scraper-go/option"
)

// XTweetLikeService contains methods and other services that help with interacting
// with the x-twitter-scraper API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewXTweetLikeService] method instead.
type XTweetLikeService struct {
	options []option.RequestOption
}

// NewXTweetLikeService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewXTweetLikeService(opts ...option.RequestOption) (r XTweetLikeService) {
	r = XTweetLikeService{}
	r.options = opts
	return
}
