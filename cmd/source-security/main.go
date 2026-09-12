// Copyright the Xquik contributors.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
)

func verify(input io.Reader, sources fs.FS, root string) error {
	data, err := fs.ReadFile(sources, "security/source-reviews.json")
	if err != nil {
		return err
	}
	var reviewed []string
	if err := json.Unmarshal(data, &reviewed); err != nil {
		return err
	}
	approved := map[string]bool{}
	for _, item := range reviewed {
		approved[item] = true
	}
	var report struct {
		Errors map[string]json.RawMessage `json:"Golang errors"`
		Stats  *struct{ Files, Found, Nosec int }
		Issues []struct {
			Rule                          string `json:"rule_id"`
			File, Line, Code              string
			Nosec                         bool
			Severity, Confidence, Details string
		}
	}
	decoder := json.NewDecoder(input)
	if err := decoder.Decode(&report); err != nil {
		return err
	}
	var extra any
	if err := decoder.Decode(&extra); err != io.EOF {
		return fmt.Errorf("invalid trailing scanner output: %v", err)
	}
	if report.Stats == nil || report.Stats.Files <= 0 || report.Stats.Nosec != 0 || report.Errors == nil || len(report.Errors) != 0 || report.Stats.Found != len(report.Issues) {
		return fmt.Errorf("incomplete or suppressed source scan")
	}
	for _, issue := range report.Issues {
		name, err := filepath.Rel(root, issue.File)
		if err != nil {
			return err
		}
		data, err := fs.ReadFile(sources, filepath.ToSlash(name))
		if err != nil {
			return err
		}
		item := [8]string{issue.Rule, filepath.ToSlash(name), issue.Line, issue.Code, fmt.Sprintf("%x", sha256.Sum256(data)), issue.Severity, issue.Confidence, issue.Details}
		fingerprint := fmt.Sprintf("%x", sha256.Sum256([]byte(fmt.Sprintf("%q", item))))
		if issue.Nosec || !approved[fingerprint] {
			return fmt.Errorf("unreviewed or changed finding: %s %s:%s", issue.Rule, name, issue.Line)
		}
	}
	fmt.Printf("Source scan: %d reports, exact approved exceptions only.\n", len(report.Issues))
	return nil
}

func main() {
	root, err := os.Getwd()
	if err == nil {
		err = verify(os.Stdin, os.DirFS(root), root)
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
