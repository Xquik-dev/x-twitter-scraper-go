// Copyright the Xquik contributors.
// SPDX-License-Identifier: Apache-2.0

package xtwitterscraper_test

import (
	"strings"
	"testing"
	"time"

	"github.com/Xquik-dev/x-twitter-scraper-go"
)

type namedReader struct {
	*strings.Reader
}

func (namedReader) Name() string {
	return "timeline.json"
}

func TestFieldHelpers(t *testing.T) {
	now := time.Unix(1_700_000_000, 0).UTC()
	if got := xtwitterscraper.String("tweet"); !got.Valid() || got.Value != "tweet" {
		t.Fatalf("String() = %+v", got)
	}
	if got := xtwitterscraper.Int(42); !got.Valid() || got.Value != 42 {
		t.Fatalf("Int() = %+v", got)
	}
	if got := xtwitterscraper.Bool(true); !got.Valid() || !got.Value {
		t.Fatalf("Bool() = %+v", got)
	}
	if got := xtwitterscraper.Float(1.5); !got.Valid() || got.Value != 1.5 {
		t.Fatalf("Float() = %+v", got)
	}
	if got := xtwitterscraper.Time(now); !got.Valid() || !got.Value.Equal(now) {
		t.Fatalf("Time() = %+v", got)
	}
	if got := xtwitterscraper.Opt("search"); !got.Valid() || got.Value != "search" {
		t.Fatalf("Opt() = %+v", got)
	}
	if got := xtwitterscraper.Ptr("value"); *got != "value" {
		t.Fatalf("Ptr() = %q", *got)
	}
	if *xtwitterscraper.IntPtr(1) != 1 ||
		!*xtwitterscraper.BoolPtr(true) ||
		*xtwitterscraper.FloatPtr(2.5) != 2.5 ||
		*xtwitterscraper.StringPtr("value") != "value" ||
		!xtwitterscraper.TimePtr(now).Equal(now) {
		t.Fatal("typed pointer helper returned the wrong value")
	}

	unnamed := xtwitterscraper.File(strings.NewReader("media"), "", "text/plain")
	if unnamed.Filename() != "" || unnamed.ContentType() != "text/plain" {
		t.Fatalf("unnamed file = %q, %q", unnamed.Filename(), unnamed.ContentType())
	}
	named := xtwitterscraper.File(namedReader{Reader: strings.NewReader("media")}, "", "application/json")
	if named.Filename() != "timeline.json" {
		t.Fatalf("named file = %q", named.Filename())
	}
	explicit := xtwitterscraper.File(strings.NewReader("media"), "followers.csv", "text/csv")
	if explicit.Filename() != "followers.csv" {
		t.Fatalf("explicit file = %q", explicit.Filename())
	}
}
