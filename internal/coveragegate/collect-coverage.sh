#!/usr/bin/env bash
# Copyright the Xquik contributors.
# SPDX-License-Identifier: Apache-2.0

set -euo pipefail

readonly reports="${XQUIK_COVERAGE_DIR:?Coverage report directory is required.}"
readonly child_counters="$reports/child-counters"
mkdir "$child_counters"
export XQUIK_GO_COVERAGE_DIR="$child_counters"

echo "==> Collecting unit and child statement, function, and line coverage"
go test -count=1 -covermode=set -coverpkg=./... -coverprofile="$reports/unit-coverage.out" ./...
if compgen -G "$child_counters/*" > /dev/null; then
  go tool covdata textfmt -i="$child_counters" -o="$reports/child-coverage.out"
fi
