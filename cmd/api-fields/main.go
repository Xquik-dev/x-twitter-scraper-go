// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"os"
	"regexp"
	"slices"
	"strings"

	"github.com/Xquik-dev/x-twitter-scraper-go/internal/apicheck"
	"github.com/Xquik-dev/x-twitter-scraper-go/option"
)

const (
	acceptedFile = "accepted-field-differences.json"
	nextStep     = "Fix each field above in the SDK, or list it in accepted-field-differences.json with a reason."
)

var differencePattern = regexp.MustCompile(`^([A-Z]+) (/\S*) (request|response) (missing|extra|enum|required|nullable) (\S+)$`)

func parseDifference(text string) (string, bool) {
	match := differencePattern.FindStringSubmatch(text)
	if match == nil || !apicheck.KnownMethod(match[1]) {
		return "", false
	}
	return strings.Join([]string{apicheck.OperationKey(match[1], match[2]), match[3], match[4], match[5]}, " "), true
}

var accepted = apicheck.GroupFile{
	Name:      acceptedFile,
	ItemsKey:  "differences",
	Item:      "difference",
	ItemShape: "an upper case HTTP method, a path starting with /, request or response, missing, extra, enum, required or nullable, and a wire path, separated by spaces",
	Parse:     parseDifference,
}

type differences struct {
	key   string
	label string
	found map[string]string
}

func (d differences) add(side, kind, wirePath string) {
	d.found[strings.Join([]string{d.key, side, kind, wirePath}, " ")] = strings.Join([]string{d.label, side, kind, wirePath}, " ")
}

func join(parent, name string) string {
	if parent == "" {
		return name
	}
	return parent + "." + name
}

func (d differences) compare(spec, sdk *shape, wirePath, side string) {
	if spec.opaque() || sdk.opaque() {
		return
	}
	switch {
	case spec.kind != sdk.kind:
		d.add(side, "missing", bodyPath(wirePath))
	case spec.kind == object:
		d.compareFields(spec, sdk, wirePath, side)
	case spec.kind == array:
		d.compare(spec.items, sdk.items, wirePath+"[]", side)
	case spec.kind == dictionary:
		d.compare(spec.values, sdk.values, wirePath+"{}", side)
	case spec.enum != nil && sdk.enum != nil && !slices.Equal(spec.enum, sdk.enum):
		d.add(side, "enum", bodyPath(wirePath))
	}
}

func bodyPath(wirePath string) string {
	if wirePath == "" {
		return "(body)"
	}
	return wirePath
}

func (d differences) compareFields(spec, sdk *shape, wirePath, side string) {
	for name, want := range spec.fields {
		here := join(wirePath, name)
		got, ok := sdk.fields[name]
		if !ok {
			d.add(side, "missing", here)
			continue
		}
		if want.required != unencoded && got.required != unencoded && want.required != got.required {
			d.add(side, "required", here)
		}
		if want.nullable != unencoded && got.nullable != unencoded && want.nullable != got.nullable {
			d.add(side, "nullable", here)
		}
		d.compare(want.shape, got.shape, here, side)
	}
	if spec.open {
		return
	}
	for name := range sdk.fields {
		if _, ok := spec.fields[name]; !ok {
			d.add(side, "extra", join(wirePath, name))
		}
	}
}

func (d differences) compareBody(spec, sdk *shape, side string) {
	switch {
	case spec != nil && sdk != nil:
		d.compare(canonical(spec), canonical(sdk), "", side)
	case spec != nil:
		d.addAll(canonical(spec), side, "missing")
	case sdk != nil:
		d.addAll(canonical(sdk), side, "extra")
	}
}

func (d differences) addAll(only *shape, side, kind string) {
	if only.kind != object {
		if !only.opaque() {
			d.add(side, kind, bodyPath(""))
		}
		return
	}
	for name := range only.fields {
		d.add(side, kind, name)
	}
}

func check(root fs.FS, baseURL option.RequestOption, source string, client *http.Client) ([]string, string, error) {
	allowed, err := accepted.Read(root)
	if err != nil {
		return nil, "", err
	}
	base, err := apicheck.Fields.BasePath(baseURL)
	if err != nil {
		return nil, "", err
	}
	implemented, unreadable, err := apicheck.Fields.SDKOperations(root, base)
	if err != nil || len(unreadable) > 0 {
		return unreadable, "", err
	}
	sdk, err := readSDK(root)
	if err != nil {
		return nil, "", err
	}
	spec, err := apicheck.LoadSpec(source, client)
	if err != nil {
		return nil, "", err
	}
	reader := specReader{document: spec.Document, problems: map[string]bool{}}
	found := map[string]string{}
	compared := 0
	var problems []string
	for key, operation := range implemented {
		api, ok := spec.Operations[key]
		if !ok {
			continue
		}
		got, readable := sdk.operation(operation)
		if !readable {
			problems = append(problems, fmt.Sprintf("%s:%d passes request params whose type the field check cannot read. Pass the method's own params argument.", operation.Position.Filename, operation.Position.Line))
			continue
		}
		compared++
		want := reader.operation(api)
		operationDifferences := differences{key: key, label: api.Method + " " + api.Path, found: found}
		operationDifferences.compare(canonical(want.query), canonical(got.query), "", "request")
		operationDifferences.compareBody(want.body, got.body, "request")
		operationDifferences.compareBody(want.response, got.response, "response")
	}
	for _, found := range []map[string]bool{reader.problems, sdk.problems} {
		for problem := range found {
			problems = append(problems, problem)
		}
	}
	if len(problems) > 0 {
		slices.Sort(problems)
		return problems, "", nil
	}
	var lines, stale []string
	for key, label := range found {
		if _, ok := allowed[key]; !ok {
			lines = append(lines, label)
		}
	}
	for key, text := range allowed {
		if _, ok := found[key]; !ok {
			stale = append(stale, fmt.Sprintf("%s lists %s, which the check no longer finds. Remove it.", acceptedFile, text))
		}
	}
	slices.Sort(lines)
	slices.Sort(stale)
	if len(lines) > 0 {
		lines = append(lines, nextStep)
	}
	if failures := append(lines, stale...); len(failures) > 0 {
		return failures, "", nil
	}
	return nil, fmt.Sprintf("The SDK's fields match the API for %s, and %s lists %s.", count(compared, "operation"), acceptedFile, count(len(allowed), "accepted difference")), nil
}

func count(number int, noun string) string {
	if number == 1 {
		return fmt.Sprintf("1 %s", noun)
	}
	return fmt.Sprintf("%d %ss", number, noun)
}

func run(source string, root fs.FS, baseURL option.RequestOption, client *http.Client, out, failures io.Writer) int {
	return apicheck.Run(check, source, root, baseURL, client, out, failures)
}

func main() {
	if code := run(os.Getenv(apicheck.SpecVariable), os.DirFS("."), option.WithEnvironmentProduction(), http.DefaultClient, os.Stdout, os.Stderr); code != 0 {
		os.Exit(code)
	}
}
