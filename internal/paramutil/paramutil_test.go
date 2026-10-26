// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

package paramutil

import (
	"encoding/json"
	"testing"

	"github.com/Xquik-dev/x-twitter-scraper-go/packages/param"
	"github.com/Xquik-dev/x-twitter-scraper-go/packages/respjson"
)

type testObject struct {
	param.APIObject
}

func TestFieldUtilities(t *testing.T) {
	value := param.NewOpt("tweet")
	if got := AddrIfPresent(value); got == nil || *got != "tweet" {
		t.Fatalf("AddrIfPresent() = %v", got)
	}
	if got := AddrIfPresent(param.Opt[string]{}); got != nil {
		t.Fatalf("AddrIfPresent(omitted) = %v", got)
	}

	valid := ToOpt("timeline", respjson.NewField(`"timeline"`))
	if !valid.Valid() || valid.Value != "timeline" {
		t.Fatalf("ToOpt(valid) = %+v", valid)
	}
	null := ToOpt("timeline", respjson.NewField(respjson.Null))
	if null.Valid() {
		t.Fatalf("ToOpt(null) = %+v", null)
	}
	omitted := ToOpt("timeline", respjson.Field{})
	if omitted.Valid() {
		t.Fatalf("ToOpt(omitted) = %+v", omitted)
	}

	if Valid(testObject{}) {
		t.Fatal("zero object was valid")
	}
	if Valid(param.NullStruct[testObject]()) {
		t.Fatal("null object was valid")
	}
	if !Valid(param.Override[testObject](json.RawMessage(`{"query":"tweet"}`))) {
		t.Fatal("overridden object was invalid")
	}
}

func TestSyncMap(t *testing.T) {
	var values SyncMap[string, int]
	if _, ok := values.Load("tweet"); ok {
		t.Fatal("empty map loaded a value")
	}
	if value, loaded := values.LoadOrStore("tweet", 1); loaded || value != 1 {
		t.Fatalf("LoadOrStore() = %d, %t", value, loaded)
	}
	if value, loaded := values.LoadOrStore("tweet", 2); !loaded || value != 1 {
		t.Fatalf("LoadOrStore() = %d, %t", value, loaded)
	}
	values.Store("tweet", 3)
	if value, ok := values.Load("tweet"); !ok || value != 3 {
		t.Fatalf("Load() = %d, %t", value, ok)
	}
}
