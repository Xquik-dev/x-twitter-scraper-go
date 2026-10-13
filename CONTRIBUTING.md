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
They are module integrity, lint, REUSE licensing, tests, statement and branch coverage,
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

Install official CodeQL CLI 2.27.0 and verify its published archive checksum.
Set `XQUIK_CODEQL_BIN` to the executable or use its default location.
`~/.cache/xquik-codeql/2.27.0/codeql/codeql`.
Download the pinned query pack before validation.

```sh
"$XQUIK_CODEQL_BIN" pack download codeql/go-queries@1.6.10
```

Each full check extracts a fresh database and runs security-extended queries.
Findings, extraction errors, and incomplete results block release.

Add regression tests for every fixed defect.

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
JSON output identifies success or failure. Raw module evidence remains available.

`bin/release-local` attaches no signed source archive.
To attach one, run `bun run release:publish` on the tagged commit afterwards.
It uses locked Sigstore 4.5.0 tooling and GitHub CLI authentication.
Use uv 0.13.0. The command selects Python 3.14.7 automatically.
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
