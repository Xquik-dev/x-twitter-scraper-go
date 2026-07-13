// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package xtwitterscraper

import (
	"github.com/Xquik-dev/x-twitter-scraper-go/option"
)

// SupportService contains methods and other services that help with interacting
// with the x-twitter-scraper API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSupportService] method instead.
type SupportService struct {
	options []option.RequestOption
	// Support ticket management
	Tickets SupportTicketService
}

// NewSupportService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSupportService(opts ...option.RequestOption) (r SupportService) {
	r = SupportService{}
	r.options = opts
	r.Tickets = NewSupportTicketService(opts...)
	return
}
