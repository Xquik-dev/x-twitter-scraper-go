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

	"github.com/Xquik-dev/x-twitter-scraper-go/internal/apicheck"
	"github.com/Xquik-dev/x-twitter-scraper-go/option"
)

const (
	exclusionFile   = "excluded-operations.json"
	missingNextStep = "Add each missing operation above to the SDK, or list it in excluded-operations.json with a reason."
)

var operationPattern = regexp.MustCompile(`^([A-Z]+) (/\S*)$`)

func parseOperation(text string) (string, bool) {
	match := operationPattern.FindStringSubmatch(text)
	if match == nil || !apicheck.KnownMethod(match[1]) {
		return "", false
	}
	return apicheck.OperationKey(match[1], match[2]), true
}

var exclusions = apicheck.GroupFile{
	Name:      exclusionFile,
	ItemsKey:  "operations",
	Item:      "operation",
	ItemShape: "an upper case HTTP method, a space and a path starting with /",
	Parse:     parseOperation,
}

func check(root fs.FS, baseURL option.RequestOption, source string, client *http.Client) ([]string, string, error) {
	excluded, err := exclusions.Read(root)
	if err != nil {
		return nil, "", err
	}
	base, err := apicheck.Coverage.BasePath(baseURL)
	if err != nil {
		return nil, "", err
	}
	implemented, unreadable, err := apicheck.Coverage.SDKOperations(root, base)
	if err != nil || len(unreadable) > 0 {
		return unreadable, "", err
	}
	spec, err := apicheck.LoadSpec(source, client)
	if err != nil {
		return nil, "", err
	}
	var missing, stale []string
	for key, operation := range spec.Operations {
		_, done := implemented[key]
		if _, skipped := excluded[key]; !done && !skipped {
			missing = append(missing, operation.Label()+" is in the API but not in the SDK.")
		}
	}
	for key, text := range excluded {
		_, known := spec.Operations[key]
		_, done := implemented[key]
		switch {
		case !known:
			stale = append(stale, fmt.Sprintf("%s lists %s, which the API no longer has. Remove it.", exclusionFile, text))
		case done:
			stale = append(stale, fmt.Sprintf("%s lists %s, which the SDK already implements. Remove it.", exclusionFile, text))
		}
	}
	slices.Sort(missing)
	slices.Sort(stale)
	if len(missing) > 0 {
		missing = append(missing, missingNextStep)
	}
	if failures := append(missing, stale...); len(failures) > 0 {
		return failures, "", nil
	}
	total := len(spec.Operations)
	return nil, fmt.Sprintf("The SDK covers %d of %d API operations, and %s lists the other %d.", total-len(excluded), total, exclusionFile, len(excluded)), nil
}

func run(source string, root fs.FS, baseURL option.RequestOption, client *http.Client, out, failures io.Writer) int {
	return apicheck.Run(check, source, root, baseURL, client, out, failures)
}

func main() {
	if code := run(os.Getenv(apicheck.SpecVariable), os.DirFS("."), option.WithEnvironmentProduction(), http.DefaultClient, os.Stdout, os.Stderr); code != 0 {
		os.Exit(code)
	}
}
