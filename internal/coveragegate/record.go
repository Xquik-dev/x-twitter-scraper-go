// Copyright the Xquik contributors.
// SPDX-License-Identifier: Apache-2.0

package coveragegate

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

// InstrumentationScope identifies the native full-source collection contract.
const InstrumentationScope = "go-packages-v3"

// Scope identifies the native toolchain and full-source instrumentation.
func Scope() string {
	return runtime.Version() + "/" + runtime.GOOS + "/" + runtime.GOARCH + "/" + InstrumentationScope
}

func recordGit(directory string, command *exec.Cmd) (string, error) {
	command.Dir = directory
	command.Env = append(os.Environ(), "LC_ALL=C")
	data, err := command.Output()
	return strings.TrimSpace(string(data)), err
}

func cleanRevision(directory string) (string, error) {
	status, err := recordGit(directory, exec.Command("git", "status", "--porcelain", "--untracked-files=all"))
	if err != nil {
		return "", err
	}
	if status != "" {
		return "", fmt.Errorf("coverage baseline requires a clean checkout")
	}
	return recordGit(directory, exec.Command("git", "rev-parse", "--verify", "HEAD"))
}

// Record collects and preserves native coverage from one clean source revision.
func Record(sourceDirectory string, nativeOutput io.Writer) (Report, error) {
	revision, err := cleanRevision(sourceDirectory)
	if err != nil {
		return Report{}, err
	}
	common, err := recordGit(sourceDirectory, exec.Command("git", "rev-parse", "--path-format=absolute", "--git-common-dir"))
	if err != nil {
		return Report{}, err
	}
	root, err := recordGit(sourceDirectory, exec.Command("git", "rev-parse", "--show-toplevel"))
	if err != nil {
		return Report{}, err
	}
	toolchain := exec.Command("go", "env", "GOVERSION", "GOOS", "GOARCH")
	toolchain.Dir = root
	selected, err := toolchain.Output()
	if err != nil {
		return Report{}, fmt.Errorf("read native Go toolchain: %w; run with %s", err, Scope())
	}
	fields := strings.Fields(string(selected))
	nativeScope := strings.Join(fields, "/") + "/" + InstrumentationScope
	if len(fields) != 3 || nativeScope != Scope() {
		return Report{}, fmt.Errorf("native toolchain does not match recorder: got %s, need %s; run the recorder with the matching Go version and target", nativeScope, Scope())
	}
	gitDirectory, err := os.OpenRoot(common)
	if err != nil {
		return Report{}, err
	}
	defer gitDirectory.Close()
	if err := gitDirectory.MkdirAll("coverage-baselines/"+InstrumentationScope, 0o700); err != nil {
		return Report{}, err
	}
	baselines, err := gitDirectory.OpenRoot("coverage-baselines/" + InstrumentationScope)
	if err != nil {
		return Report{}, err
	}
	defer baselines.Close()
	if err := baselines.Mkdir(revision, 0o700); err != nil {
		if os.IsExist(err) {
			return Report{}, fmt.Errorf("coverage baseline already exists for %s; preserve its evidence", revision)
		}
		return Report{}, err
	}
	reports, err := baselines.OpenRoot(revision)
	if err != nil {
		return Report{}, err
	}
	defer reports.Close()
	directory := reports.Name()
	fmt.Fprintf(nativeOutput, "Native baseline reports: %s\n", directory)
	for _, collector := range []struct {
		name    string
		collect func(string, string, io.Writer) error
	}{
		{"statements", CollectStatements},
		{"branches", CollectBranches},
	} {
		collectionErr := collector.collect(root, directory, nativeOutput)
		current, err := cleanRevision(root)
		if err != nil || current != revision {
			return Report{}, fmt.Errorf("source changed during collection; retained reports are not a verified baseline")
		}
		if collectionErr != nil {
			return Report{}, fmt.Errorf("collect %s: %w", collector.name, collectionErr)
		}
	}
	metrics, err := Measure(reports.FS())
	if err != nil {
		return Report{}, err
	}
	if err := Compare(metrics, metrics); err != nil {
		return Report{}, err
	}
	report := Report{Revision: revision, Scope: Scope(), Total: metrics}
	pending, err := reports.OpenFile(".baseline-pending.json", os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0o600)
	if err != nil {
		return Report{}, err
	}
	encoder := json.NewEncoder(pending)
	encoder.SetIndent("", "  ")
	if err := errors.Join(encoder.Encode(report), pending.Close()); err != nil {
		return Report{}, err
	}
	if err := reports.Link(".baseline-pending.json", "report.json"); err != nil {
		return Report{}, fmt.Errorf("publish immutable coverage baseline: %w", err)
	}
	if err := reports.Remove(".baseline-pending.json"); err != nil {
		return Report{}, fmt.Errorf("remove published baseline's temporary alias: %w", err)
	}
	return report, nil
}
