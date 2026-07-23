// Copyright the Xquik contributors.
// SPDX-License-Identifier: Apache-2.0

package param

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

type stringValue string

func (value stringValue) String() string {
	return "formatted:" + string(value)
}

func TestOptBehavior(t *testing.T) {
	valid := NewOpt("tweet")
	if !valid.Valid() || valid.Or("fallback") != "tweet" || valid.String() != "tweet" {
		t.Fatalf("valid option = %+v", valid)
	}
	omitted := Opt[string]{}
	if omitted.Valid() || omitted.Or("fallback") != "fallback" {
		t.Fatalf("omitted option = %+v", omitted)
	}
	null := Null[string]()
	if null.Valid() || null.String() != "null" {
		t.Fatalf("null option = %+v", null)
	}
	formatted := NewOpt(stringValue("tweet"))
	if got := formatted.String(); got != "formatted:tweet" {
		t.Fatalf("String() = %q", got)
	}
}

func TestOptJSONAndTimeLayout(t *testing.T) {
	valid := NewOpt("tweet")
	if got, err := valid.MarshalJSON(); err != nil || string(got) != `"tweet"` {
		t.Fatalf("MarshalJSON() = %s, %v", got, err)
	}
	if got, err := (Opt[string]{}).MarshalJSON(); err != nil || string(got) != "null" {
		t.Fatalf("omitted MarshalJSON() = %s, %v", got, err)
	}

	var decoded Opt[string]
	if err := json.Unmarshal([]byte(`"timeline"`), &decoded); err != nil ||
		!decoded.Valid() || decoded.Value != "timeline" {
		t.Fatalf("decoded = %+v, %v", decoded, err)
	}
	if err := json.Unmarshal([]byte(`null`), &decoded); err != nil || decoded.Valid() {
		t.Fatalf("decoded null = %+v, %v", decoded, err)
	}
	if err := json.Unmarshal([]byte(` null `), &decoded); err != nil || decoded.Valid() {
		t.Fatalf("decoded spaced null = %+v, %v", decoded, err)
	}
	if err := json.Unmarshal([]byte(`{`), &decoded); err == nil {
		t.Fatal("invalid JSON was accepted")
	}

	when := time.Date(2026, time.July, 23, 12, 30, 0, 0, time.UTC)
	timeValue := NewOpt(when)
	if got := timeValue.MarshalJSONWithTimeLayout("2006-01-02"); string(got) != `"2026-07-23"` {
		t.Fatalf("formatted time = %s", got)
	}
	if got := valid.MarshalJSONWithTimeLayout(time.RFC3339); got != nil {
		t.Fatalf("non-time layout = %s", got)
	}
	if got := Null[time.Time]().MarshalJSONWithTimeLayout(time.RFC3339); got != nil {
		t.Fatalf("null time layout = %s", got)
	}

	if got := fmt.Sprint(valid); got != "tweet" {
		t.Fatalf("formatted option = %q", got)
	}
}
