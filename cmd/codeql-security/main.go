// Copyright the Xquik contributors.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func validate(data []byte) error {
	var report struct {
		Version string
		Runs    []struct {
			Results     []json.RawMessage
			Invocations []struct {
				ExecutionSuccessful        bool
				ToolExecutionNotifications []struct {
					Level      string
					Descriptor struct{ ID string }
				}
			}
		}
	}
	if err := json.Unmarshal(data, &report); err != nil {
		return err
	}
	if report.Version != "2.1.0" || len(report.Runs) != 1 {
		return fmt.Errorf("expected one SARIF 2.1.0 run")
	}
	run := report.Runs[0]
	if run.Results == nil || len(run.Results) != 0 || len(run.Invocations) == 0 {
		return fmt.Errorf("CodeQL findings or incomplete analysis prevent release")
	}
	extracted := 0
	for _, invocation := range run.Invocations {
		if !invocation.ExecutionSuccessful {
			return fmt.Errorf("CodeQL invocation failed")
		}
		for _, notification := range invocation.ToolExecutionNotifications {
			switch notification.Level {
			case "none", "note":
			default:
				return fmt.Errorf("CodeQL diagnostic requires review: %s", notification.Descriptor.ID)
			}
			if notification.Descriptor.ID == "go/diagnostics/successfully-extracted-files" {
				extracted++
			}
		}
	}
	if extracted == 0 {
		return fmt.Errorf("CodeQL did not confirm extracted Go files")
	}
	fmt.Printf("CodeQL completed without findings; extracted Go files: %d\n", extracted)
	return nil
}

func main() {
	data, err := io.ReadAll(os.Stdin)
	if err == nil {
		err = validate(data)
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
