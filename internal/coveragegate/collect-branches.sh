#!/usr/bin/env bash

# SPDX-FileCopyrightText: 2026 Xquik contributors
#
# SPDX-License-Identifier: Apache-2.0

set -euo pipefail

readonly gocove_version="v0.0.0-20260427173940-a0dceb9dadca"
readonly reports="${XQUIK_COVERAGE_DIR:?Coverage report directory is required.}"
readonly child_counters="$reports/branch-child-counters"
mkdir "$child_counters"
readonly temp_root="${TMPDIR:-/tmp}"
readonly native_cache="$(go env GOCACHE)"
branch_workspace="$(mktemp -d "${temp_root%/}/xquik-gocove.XXXXXX")"
readonly branch_meta="$branch_workspace/meta"
readonly branch_cache="$branch_workspace/go-cache"
mkdir -p "$branch_meta" "$branch_cache"

cleanup() {
  case "$branch_workspace" in
    "${temp_root%/}"/xquik-gocove.*)
      rm -rf -- "$branch_cache"
      ;;
    *)
      echo "Refusing to remove unexpected path: $branch_workspace" >&2
      ;;
  esac
}
trap cleanup EXIT

GOBIN="$branch_workspace" go install "github.com/srvgit/gocove/cmd/gocove@${gocove_version}"
export GOCOVE_META_DIR="$branch_meta"
export GOCACHE="$branch_cache"
export XQUIK_GO_COVERAGE_DIR="$child_counters"
export XQUIK_GO_NATIVE_CACHE="$native_cache"
cat > "$branch_workspace/run-test" <<'RUNNER'
#!/usr/bin/env bash
set -euo pipefail
export GOCACHE="${XQUIK_GO_NATIVE_CACHE:?Native Go cache is required.}"
exec "$@"
RUNNER
chmod +x "$branch_workspace/run-test"

echo "==> Collecting branch coverage"
# Rebuild every package. gocove does not emit catalogs for Go cache hits.
"$branch_workspace/gocove" test -a -count=1 -covermode=set -coverpkg=./... "-exec=$branch_workspace/run-test" ./...
"$branch_workspace/gocove" \
  report \
  --meta-dir="$branch_meta" \
  --format=json > "$branch_meta/report.json"

cp "$branch_meta/report.json" "$reports/branches.json"
cp "$branch_meta/_gocove/cover.profile" "$reports/unit-branch-coverage.out"
cp -R "$branch_meta" "$reports/branch-metadata"
if compgen -G "$child_counters/*" > /dev/null; then
  go tool covdata textfmt -i="$child_counters" -o="$reports/branch-child-coverage.out"
fi
