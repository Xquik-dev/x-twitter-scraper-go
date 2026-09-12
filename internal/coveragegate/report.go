// Copyright the Xquik contributors.
// SPDX-License-Identifier: Apache-2.0

package coveragegate

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/fs"
	"math"
	"path"

	"github.com/Xquik-dev/x-twitter-scraper-go/internal/coveragebranch"
	"golang.org/x/tools/cover"
)

// Report binds measured counts to their revision and instrumentation scope.
type Report struct {
	Revision string
	Scope    string
	Total    map[string]Metric
}

// Baseline rejects missing, malformed, or mismatched baseline evidence.
func Baseline(files fs.FS, revision, scope string) (Report, error) {
	var report Report
	data, err := fs.ReadFile(files, "report.json")
	if err != nil {
		return report, err
	}
	if err := json.Unmarshal(data, &report); err != nil {
		return report, err
	}
	if revision == "" || scope == "" || report.Revision != revision || report.Scope != scope {
		return report, fmt.Errorf("coverage baseline revision or instrumentation scope differs")
	}
	return report, nil
}

// Measure reads native Go, Cobertura, and branch instrumentation reports.
func Measure(files fs.FS) (map[string]Metric, error) {
	inputs := make(map[string][]byte)
	for _, name := range []string{"coverage.out", "coverage.xml", "branches.json", "branches.cover"} {
		data, err := fs.ReadFile(files, name)
		if err != nil {
			return nil, fmt.Errorf("read %s: %w", name, err)
		}
		inputs[name] = data
	}
	profiles, err := cover.ParseProfilesFromReader(bytes.NewReader(inputs["coverage.out"]))
	if err != nil {
		return nil, fmt.Errorf("read statement coverage: %w", err)
	}
	var statements Metric
	sources := make(map[string]bool, len(profiles))
	for _, profile := range profiles {
		sources[profile.FileName] = true
		for _, block := range profile.Blocks {
			statements.Total += int64(block.NumStmt)
			if block.Count > 0 {
				statements.Covered += int64(block.NumStmt)
			}
		}
	}
	var document struct {
		XMLName  xml.Name `xml:"coverage"`
		Covered  int64    `xml:"lines-covered,attr"`
		Total    int64    `xml:"lines-valid,attr"`
		Packages []struct {
			Name    string `xml:"name,attr"`
			Classes []struct {
				Filename string `xml:"filename,attr"`
				Methods  []struct {
					Rate float64 `xml:"line-rate,attr"`
				} `xml:"methods>method"`
			} `xml:"classes>class"`
		} `xml:"packages>package"`
	}
	if err := xml.Unmarshal(inputs["coverage.xml"], &document); err != nil {
		return nil, fmt.Errorf("read line and function coverage: %w", err)
	}
	var functions Metric
	for _, pkg := range document.Packages {
		for _, class := range pkg.Classes {
			filename := pkg.Name + "/" + path.Base(class.Filename)
			if !sources[filename] {
				return nil, fmt.Errorf("XML coverage contains unexpected or duplicate source %s", filename)
			}
			delete(sources, filename)
			functions.Total += int64(len(class.Methods))
			for _, method := range class.Methods {
				if math.IsNaN(method.Rate) || math.IsInf(method.Rate, 0) || method.Rate < 0 || method.Rate > 1 {
					return nil, fmt.Errorf("invalid method coverage rate %v", method.Rate)
				}
				if method.Rate > 0 {
					functions.Covered++
				}
			}
		}
	}
	if len(sources) != 0 {
		return nil, fmt.Errorf("XML coverage omits %d native source files", len(sources))
	}
	covered, total, err := coveragebranch.Measure(inputs["branches.json"], bytes.NewReader(inputs["branches.cover"]))
	if err != nil {
		return nil, err
	}
	return map[string]Metric{
		"statements": statements,
		"functions":  functions,
		"lines":      {Total: document.Total, Covered: document.Covered},
		"branches":   {Total: int64(total), Covered: int64(covered)},
	}, nil
}
