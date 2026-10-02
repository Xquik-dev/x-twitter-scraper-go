// Copyright the Xquik contributors.
// SPDX-License-Identifier: Apache-2.0

package coveragegate

import (
	_ "embed"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

//go:embed collect-coverage.sh
var statementCollector string

//go:embed collect-branches.sh
var branchCollector string

func collectNative(sourceDirectory, reportDirectory, script string, output io.Writer) error {
	command := exec.Command("bash", "-s")
	command.Dir = sourceDirectory
	for _, variable := range os.Environ() {
		name, _, _ := strings.Cut(variable, "=")
		if name != "GOCOVERDIR" && name != "XQUIK_GO_COVERAGE_DIR" {
			command.Env = append(command.Env, variable)
		}
	}
	command.Env = append(command.Env, "GOTOOLCHAIN="+runtime.Version(), "XQUIK_COVERAGE_DIR="+reportDirectory)
	command.Stdin = strings.NewReader(script)
	command.Stdout, command.Stderr = output, output
	return command.Run()
}

func mergeCollectedProfiles(reports *os.Root, unitName, childName, reportName string) error {
	unit, err := reports.ReadFile(unitName)
	if err != nil {
		return err
	}
	child, err := reports.ReadFile(childName)
	if err != nil && !os.IsNotExist(err) {
		return err
	}
	merged, err := MergeNativeCoverage(unit, child)
	if err != nil {
		return err
	}
	return reports.WriteFile(reportName, merged, 0o600)
}

// CollectStatements preserves native unit and child profiles and their merged report.
func CollectStatements(sourceDirectory, reportDirectory string, output io.Writer) error {
	if err := collectNative(sourceDirectory, reportDirectory, statementCollector, output); err != nil {
		return err
	}
	reports, err := os.OpenRoot(reportDirectory)
	if err != nil {
		return err
	}
	defer reports.Close()
	if err := mergeCollectedProfiles(reports, "unit-coverage.out", "child-coverage.out", "coverage.out"); err != nil {
		return err
	}
	profile, err := reports.Open("coverage.out")
	if err != nil {
		return err
	}
	defer profile.Close()
	report, err := reports.OpenFile("coverage.xml", os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0o600)
	if err != nil {
		return err
	}
	converter := exec.Command("go", "run", "github.com/boumenot/gocover-cobertura@v1.5.0", "-strict", "-by-files")
	converter.Dir = sourceDirectory
	converter.Env = append(os.Environ(), "GOTOOLCHAIN="+runtime.Version())
	converter.Stdin, converter.Stdout, converter.Stderr = profile, report, output
	conversionErr := converter.Run()
	return errors.Join(conversionErr, report.Close())
}

// CollectBranches preserves native branch execution and its complete source catalog.
func CollectBranches(sourceDirectory, reportDirectory string, output io.Writer) error {
	if err := collectNative(sourceDirectory, reportDirectory, branchCollector, output); err != nil {
		return err
	}
	reports, err := os.OpenRoot(reportDirectory)
	if err != nil {
		return err
	}
	defer reports.Close()
	return mergeCollectedProfiles(reports, "unit-branch-coverage.out", "branch-child-coverage.out", "branches.cover")
}

// Collect runs a native collector in the caller's source directory.
func Collect(kind string, output io.Writer) error {
	directory := os.Getenv("XQUIK_COVERAGE_DIR")
	if directory == "" {
		var err error
		directory, err = os.MkdirTemp("", "xquik-go-coverage.")
		if err != nil {
			return err
		}
	}
	fmt.Fprintf(output, "Native coverage reports: %s\n", directory)
	switch kind {
	case "--collect-statements":
		return CollectStatements(".", directory, output)
	case "--collect-branches":
		return CollectBranches(".", directory, output)
	default:
		return fmt.Errorf("unknown native collector: %s", kind)
	}
}
