// Copyright the Xquik contributors.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"strings"
	"testing"
)

func TestValidate(t *testing.T) {
	const valid = `{"version":"2.1.0","runs":[{"results":[],"invocations":[{"executionSuccessful":true,"toolExecutionNotifications":[{"level":"none","descriptor":{"id":"go/diagnostics/successfully-extracted-files"}}]}]}]}`
	for _, test := range []struct {
		name, report string
		pass         bool
	}{
		{"complete", valid, true},
		{"malformed", "{", false},
		{"trailing", valid + "{}", false},
		{"empty", "{}", false},
		{"version", strings.Replace(valid, "2.1.0", "2.0.0", 1), false},
		{"findings", strings.Replace(valid, `"results":[]`, `"results":[{}]`, 1), false},
		{"missing results", strings.Replace(valid, `"results":[],`, "", 1), false},
		{"no invocation", `{"version":"2.1.0","runs":[{"results":[],"invocations":[]}]}`, false},
		{"failed", strings.Replace(valid, "true", "false", 1), false},
		{"warning", strings.Replace(valid, `"none"`, `"warning"`, 1), false},
		{"error", strings.Replace(valid, `"none"`, `"error"`, 1), false},
		{"no extraction", strings.Replace(valid, "successfully-extracted-files", "other", 1), false},
	} {
		t.Run(test.name, func(t *testing.T) {
			if err := validate([]byte(test.report)); (err == nil) != test.pass {
				t.Fatalf("validation error=%v, want pass=%v", err, test.pass)
			}
		})
	}
}
