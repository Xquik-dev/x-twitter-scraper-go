// Copyright the Xquik contributors.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"strings"
	"testing"
	"testing/fstest"
)

func TestExactReviewedFindings(t *testing.T) {
	const report = `{"Golang errors":{},"Stats":{"files":1,"found":1,"nosec":0},"Issues":[{"rule_id":"G104","file":"/repo/example.go","line":"1","code":"sample"}]}`
	const review = `["0be58b411a169910b7c8ac0db509b092d5d7cbe749be6007721c822dde350790"]`
	for _, test := range []struct {
		name, report, review, source string
		valid                        bool
	}{
		{"approved", report, review, "", true},
		{"changed source", report, review, "changed", false},
		{"changed rule", strings.ReplaceAll(report, "G104", "G101"), review, "", false},
		{"changed line", strings.Replace(report, `"line":"1"`, `"line":"2"`, 1), review, "", false},
		{"changed code", strings.ReplaceAll(report, "sample", "other"), review, "", false},
		{"changed severity", strings.Replace(report, `"rule_id"`, `"severity":"HIGH","rule_id"`, 1), review, "", false},
		{"changed confidence", strings.Replace(report, `"rule_id"`, `"confidence":"HIGH","rule_id"`, 1), review, "", false},
		{"changed details", strings.Replace(report, `"rule_id"`, `"details":"new","rule_id"`, 1), review, "", false},
		{"outside source", strings.ReplaceAll(report, "/repo/example.go", "/outside.go"), review, "", false},
		{"missing source", strings.ReplaceAll(report, "example.go", "missing.go"), review, "", false},
		{"missing review", report, `[]`, "", false},
		{"broken review", report, "{", "", false},
		{"broken report", "{", review, "", false},
		{"absent statistics", "{}", review, "", false},
		{"empty scan", strings.Replace(report, `"files":1`, `"files":0`, 1), review, "", false},
		{"negative files", strings.Replace(report, `"files":1`, `"files":-1`, 1), review, "", false},
		{"missing errors", strings.Replace(report, `"Golang errors":{},`, "", 1), review, "", false},
		{"missing reports", strings.Replace(report, `"found":1`, `"found":2`, 1), review, "", false},
		{"suppressed", strings.Replace(report, `"nosec":0`, `"nosec":1`, 1), review, "", false},
		{"suppressed issue", strings.Replace(report, `"rule_id"`, `"nosec":true,"rule_id"`, 1), review, "", false},
		{"analysis failure", strings.Replace(report, `"Golang errors":{}`, `"Golang errors":{"example.go":["failed"]}`, 1), review, "", false},
		{"extra document", report + "{}", review, "", false},
		{"truncated trailing output", report + "{", review, "", false},
	} {
		t.Run(test.name, func(t *testing.T) {
			sources := fstest.MapFS{"security/source-reviews.json": {Data: []byte(test.review)}, "example.go": {Data: []byte(test.source)}}
			if err := verify(strings.NewReader(test.report), sources, "/repo"); (err == nil) != test.valid {
				t.Fatalf("valid=%v, error=%v", test.valid, err)
			}
		})
	}
	if err := verify(strings.NewReader(report), fstest.MapFS{}, "/repo"); err == nil {
		t.Fatal("missing review file passed")
	}
	for _, root := range []string{"relative", "/repo"} {
		sources := fstest.MapFS{"security/source-reviews.json": {Data: []byte(review)}, "example.go": {}}
		input := report
		if root == "/repo" {
			input = `{"Golang errors":{},"Stats":{"files":1,"found":0,"nosec":0},"Issues":[]}`
		}
		if err := verify(strings.NewReader(input), sources, root); (err == nil) != (root == "/repo") {
			t.Fatal(err)
		}
	}
}
