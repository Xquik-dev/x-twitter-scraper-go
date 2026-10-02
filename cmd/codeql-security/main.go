// Copyright the Xquik contributors.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/url"
	"os"
	"os/exec"
	"path"
	"strings"
)

func validate(data []byte, expected []string) (int, error) {
	var report struct {
		Version string
		Runs    []struct {
			Results     []json.RawMessage
			Invocations []struct {
				ExecutionSuccessful        bool
				ToolExecutionNotifications []struct {
					Level      string
					Descriptor struct{ ID string }
					Locations  []struct {
						PhysicalLocation struct {
							ArtifactLocation struct {
								URI       string
								URIBaseID string
							}
						}
					}
				}
			}
		}
	}
	if err := json.Unmarshal(data, &report); err != nil {
		return 0, err
	}
	if report.Version != "2.1.0" || len(report.Runs) != 1 {
		return 0, fmt.Errorf("expected one SARIF 2.1.0 run")
	}
	run := report.Runs[0]
	if run.Results == nil || len(run.Results) != 0 || len(run.Invocations) == 0 {
		return 0, fmt.Errorf("CodeQL findings or incomplete analysis prevent release")
	}
	if len(expected) == 0 {
		return 0, fmt.Errorf("CodeQL source inventory has no tracked Go files")
	}
	sources := make(map[string]bool, len(expected))
	for _, name := range expected {
		sources[name] = false
	}
	for _, invocation := range run.Invocations {
		if !invocation.ExecutionSuccessful {
			return 0, fmt.Errorf("CodeQL invocation failed")
		}
		for _, notification := range invocation.ToolExecutionNotifications {
			switch notification.Level {
			case "none", "note":
			default:
				return 0, fmt.Errorf("CodeQL diagnostic requires review: %s", notification.Descriptor.ID)
			}
			if notification.Descriptor.ID == "go/diagnostics/successfully-extracted-files" {
				if len(notification.Locations) == 0 {
					return 0, fmt.Errorf("CodeQL source inventory diagnostic has no locations")
				}
				for _, location := range notification.Locations {
					artifact := location.PhysicalLocation.ArtifactLocation
					uri, err := url.Parse(artifact.URI)
					if err != nil || artifact.URIBaseID != "%SRCROOT%" || uri.Scheme != "" || uri.Host != "" || uri.RawQuery != "" || uri.Fragment != "" || uri.Path == "" || path.IsAbs(uri.Path) || path.Clean(uri.Path) != uri.Path || strings.Contains(uri.Path, "\\") || strings.HasPrefix(uri.Path, "../") {
						return 0, fmt.Errorf("CodeQL source inventory contains an invalid source location: %s", artifact.URI)
					}
					if strings.HasSuffix(uri.Path, ".go") {
						if _, exists := sources[uri.Path]; !exists {
							return 0, fmt.Errorf("CodeQL source inventory contains an untracked Go file: %s", uri.Path)
						}
						sources[uri.Path] = true
					}
				}
			}
		}
	}
	for _, name := range expected {
		if !sources[name] {
			return 0, fmt.Errorf("CodeQL source inventory is missing %s; enable extract_tests & extract every tracked Go file", name)
		}
	}
	return len(sources), nil
}

func main() {
	data, err := io.ReadAll(os.Stdin)
	count := 0
	if err == nil {
		var inventory []byte
		inventory, err = exec.Command("git", "ls-files", "-z", "--", "*.go").Output()
		if err == nil {
			expected := strings.Split(strings.TrimSuffix(string(inventory), "\x00"), "\x00")
			if len(inventory) == 0 {
				expected = nil
			}
			count, err = validate(data, expected)
		}
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Printf("CodeQL completed without findings; extracted Go files: %d\n", count)
}
