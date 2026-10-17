# Contributing

Read [GOVERNANCE.md](GOVERNANCE.md) before proposing major changes.

Follow the shared [Xquik contribution policy][contribution-policy].

## Set Up

Install Go 1.27.2, Bun, jq and uv.

```sh
./scripts/bootstrap
```

Use a local module replacement when testing another project.

```sh
go mod edit \
  -replace github.com/Xquik-dev/x-twitter-scraper-go=/path/to/x-twitter-scraper-go
```

Never commit credentials or runtime environment files.

## Generated Code

Most SDK files come from the public OpenAPI contract.

Preserve generated method names and response contracts.

Avoid manual generated-file changes when a generator fix exists.

Place stable examples outside generated directories.

## Verify Changes

Run focused tests while editing.

Run `bin/check-local` on your pushed branch head. GitHub Actions does not run here.
It runs these checks in order and stops at the first failure.
They are API coverage, API fields, module integrity, lint, REUSE licensing, tests, statement and branch coverage,
coverage against the parent, race detection, the dependency audit, the signing tool audit,
the source scan, reproducible archives, CodeQL, and fuzzing.
REUSE is required. It runs at a pinned version through `uv`, never from another install.
CodeQL is required. A missing CLI, any finding or any file CodeQL could not read fails the run.
It fuzzes for 15 seconds. `scripts/fuzz` runs the full 2 minute campaign on its own.
The parent comparison checks statement, branch, function, and line coverage against the parent commit.
A parent without a recorded baseline is measured first in a temporary clone, once per parent.
On a clean HEAD that GitHub already has, it posts the `xquik-local-checks` commit status.
The status lists only the checks that passed and says success only when all passed.
Pass `--no-status` to skip it.
Every check removes its temp files from `TMPDIR` when it passes, fails or stops on a signal.
Scripts create them with `temp_path` from `scripts/temp-paths`, which removes them on exit.

The dependency audit, `scripts/dependency-audit`, runs govulncheck over every package and test.
Every finding counts, whether govulncheck traced a call to it or not.
A finding in a module that ships to callers always blocks, with or without a fix.
A module ships when a package that callers can import depends on it.
A finding in a module that only commands or tests use blocks when a fix exists.
Without a fix, it passes only while `audit-exceptions.json` lists it.
Each entry names the advisory `id`, the `module`, a `reason`, and a `recheck` date.
The reason is `build-time only` or `peer dependency provided by the host`.
The recheck date falls within 30 days. The entry blocks again on that date.
The list is empty today.

The source scan, `scripts/source-security`, runs gosec over every package and test.
Every finding must match a reviewed record in `security/source-reviews.md` and `security/source-reviews.json`.
A new finding, or a reviewed one whose code or line moved, fails the run.

Run `bin/check-local` before requesting review.
Run `scripts/fuzz` as well when you change `internal/apiform`.

```sh
bin/check-local
scripts/fuzz
```

SDK commits are exempt from LOC reductions and per-commit coverage gains.

Preserve coverage and report all four metrics against the verified parent.

SDK checks have no time limit. Builds run separately.

Full checks run the complete 2-minute fuzz campaign directly.
SDK elapsed time remains visible and never blocks delivery by itself.

Install official CodeQL CLI 2.27.2 and verify its published archive checksum.
Set `XQUIK_CODEQL_BIN` to the executable or use its default location.
`~/.cache/xquik-codeql/2.27.2/codeql/codeql`.
Download the pinned query pack before validation.

```sh
"$XQUIK_CODEQL_BIN" pack download codeql/go-queries@1.6.12
```

Each full check extracts a fresh database and runs security-extended queries.
Findings, extraction errors, and incomplete results block release.

Add regression tests for every fixed defect.

## API coverage

`bin/check-local` starts with `go run ./cmd/api-coverage`.
It compares the operations in the public API spec with the requests the SDK sends.
It fails when the API has an operation that the SDK neither implements nor lists in `excluded-operations.json`.
An operation that the SDK has and the spec lacks is fine.

The check reads the spec from https://xquik.com/openapi.json.
Set `XQUIK_OPENAPI_SPEC` to a local file path or another URL to read another copy. The spec must be JSON.

It reads the SDK operations from the SDK source.
Each method of a `Service` type in the root package sends its request through `requestconfig.ExecuteNewRequest`.
The request names an `http.Method` constant and a path string or `fmt.Sprintf` format.
The check puts the path of the default base URL from `option.WithEnvironmentProduction` before each path.
It fails on a request it cannot read and names its file and line.
Path parameter names do not matter, so `{id}` matches `{tweetId}`.

Fix a failure by adding each missing operation to the SDK.
An operation that the SDK should not have goes in `excluded-operations.json` with a reason.
Remove an entry when the API drops its operation or the SDK implements it. The check fails on a stale entry.

`excluded-operations.json` is a list of groups.
Each group has a non-empty `reason` and at least 1 operation.
Write an operation as an upper case method, a space and the absolute path from the spec. List each operation once.

```json
[
  {
    "reason": "Searches the Xquik docs at https://xquik.com/ask, outside the SDK base URL https://xquik.com/api/v1.",
    "operations": ["GET /ask", "POST /ask"]
  }
]
```

## API fields

`bin/check-local` runs `go run ./cmd/api-fields` right after the coverage check.
It reads the spec the same way, from https://xquik.com/openapi.json or `XQUIK_OPENAPI_SPEC`.
For each operation that the SDK implements, it compares the request query params, the request body and the 2xx JSON response with the SDK's Go types.
It reads the JSON request body, or the form body when the operation takes no JSON.
It reads wire names from the `query` and `json` tags at every depth. Header and path params stay out.

It follows `$ref` and `allOf` and walks arrays through their items.
It merges `oneOf`, `anyOf` and several 2xx bodies into 1 object, where a field is required only when every variant requires it.
A free-form object, a cycle, a map or an `any` value stops the walk.
An extra SDK field is fine under a spec object whose `additionalProperties` is set and not `false`.
A field missing on 1 side counts once, at its top position.

It reports 5 kinds of difference.

- `missing` is a field in the spec that the SDK lacks or types as another kind, such as a string for an object.
- `extra` is a field in the SDK that the spec lacks.
- `enum` means both sides list enum values, and the values differ.
- `required` means the required flag differs.
- `nullable` means the SDK marks a field `api:"nullable"` that the spec does not mark nullable.

Each difference prints as a line such as `GET /api/v1/x/tweets/{id}/replies request extra cardName`.
That is the method, the absolute path, `request` or `response`, the kind and the wire path.
It also fails on a Go type or spec reference it cannot resolve.

Fix a difference in the SDK when the fix is additive.
Never remove or rename a public field or param to pass the check.
List a difference the SDK keeps on purpose in `accepted-field-differences.json` with its reason.
The file has the shape of `excluded-operations.json`, with `differences` in place of `operations`.
The check fails on an entry it no longer finds.

```json
[
  {
    "reason": "Deprecated alias of minLikes, which the API still reads. The SDK keeps it so existing calls still compile.",
    "differences": ["GET /api/v1/x/tweets/search request extra minFaves"]
  }
]
```

The SDK keeps plain `string` fields where the spec now lists enum values.
Turning a `string` into a named enum type breaks code that assigns a `string` to it. So the check compares enum values only when both sides list them.

## Submit changes

Use clear Conventional Commit subjects and sign commits with `git commit --signoff`.
Follow the shared [review policy][review-policy].

## Release

Maintainers release from their own machine with `bin/release-local`.
Run it on a clean `main` that matches `origin/main`.
It reads the version from `.release-please-manifest.json`.
It refuses when the `vVERSION` tag already exists on GitHub.
It runs `bin/check-local`, the dependency audit, and `go build ./...`.
It builds the module zip as proxy.golang.org would and prints its checksum.
Next it pushes the annotated tag and asks proxy.golang.org to fetch it.
pkg.go.dev lists the version once the proxy serves it.
Last, it creates the GitHub release with notes from `CHANGELOG.md`.
Go needs no registry credential. GitHub CLI uses its own login.

Run `bin/release-local --dry-run` to see what a release would publish.
It runs every step except the tag and the GitHub release.
It prints the module path, the version, the checksum, and every file in the module zip.

Verify published modules with `bun run release:verify vVERSION COMMIT_SHA`.
The command uses a fresh module cache and the public Go checksum database.
It checks the exact version and source commit, then installs and constructs a client.
JSON output identifies success or failure. The command removes its module cache when it exits.

`bin/release-local` attaches no signed source archive.
To attach one, run `bun run release:publish` on the tagged commit afterwards.
It uses locked Sigstore 4.5.0 tooling and GitHub CLI authentication.
Use uv 0.13.1. The command selects Python 3.14.8 automatically.
Set `XQUIK_RELEASE_IDENTITY` and `XQUIK_RELEASE_ISSUER` to the authorized signer.
Signing uses Sigstore authentication and records local build provenance publicly.
Its attestations identify that signer.
It verifies signatures before uploading and rejects conflicting existing artifacts.
Retry the same command after partial publication. Matching artifacts are retained.

## Report Security Issues

Never disclose suspected vulnerabilities in public issues.

Follow [SECURITY.md](SECURITY.md) for private reporting.

[contribution-policy]: https://github.com/Xquik-dev/.github/blob/main/CONTRIBUTING.md
[review-policy]: https://github.com/Xquik-dev/.github/blob/main/REVIEWING.md

Xquik is an independent third-party service. Not affiliated with X Corp. "Twitter" and "X" are trademarks of X Corp.
