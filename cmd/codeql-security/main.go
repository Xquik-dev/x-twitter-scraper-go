// Copyright the Xquik contributors.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

const extractionErrors = "go/diagnostics/extraction-errors"

func validate(data []byte) error {
	var report struct {
		Version string
		Runs    []struct {
			Tool struct {
				Scanner struct {
					Notifications []struct{ ID string }
				} `json:"driver"`
			}
			Results     []json.RawMessage
			Invocations []struct {
				ExecutionSuccessful        bool
				ToolExecutionNotifications []struct {
					Level      string
					Message    struct{ Text string }
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
	ran := false
	for _, descriptor := range run.Tool.Scanner.Notifications {
		ran = ran || descriptor.ID == extractionErrors
	}
	if !ran {
		return fmt.Errorf("The CodeQL scan lacks %s, so it cannot show that it read every file. Make the suite in scripts/codeql-security run that diagnostic query.", extractionErrors)
	}
	extracted := 0
	var unread []string
	for _, invocation := range run.Invocations {
		if !invocation.ExecutionSuccessful {
			return fmt.Errorf("CodeQL invocation failed")
		}
		for _, notification := range invocation.ToolExecutionNotifications {
			switch {
			case notification.Descriptor.ID == extractionErrors:
				unread = append(unread, fmt.Sprintf("%q", notification.Message.Text))
			case notification.Level != "none" && notification.Level != "note":
				return fmt.Errorf("CodeQL diagnostic %s needs review. It said %q. Fix the cause, then rerun bin/check-local.", notification.Descriptor.ID, notification.Message.Text)
			}
			if notification.Descriptor.ID == "go/diagnostics/successfully-extracted-files" {
				extracted++
			}
		}
	}
	if len(unread) > 0 {
		return fmt.Errorf("CodeQL could not read all Go code, so the scan is incomplete. It said %s. Fix that code so Go compiles it, then rerun bin/check-local.", strings.Join(unread, " and "))
	}
	if extracted == 0 {
		return fmt.Errorf("CodeQL did not confirm extracted Go files")
	}
	fmt.Printf("CodeQL completed without findings and extracted %d files.\n", extracted)
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
