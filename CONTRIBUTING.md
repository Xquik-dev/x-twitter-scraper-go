# Contributing

Read [GOVERNANCE.md](GOVERNANCE.md) before proposing major changes.

Follow the shared [Xquik contribution policy][contribution-policy].

## Set Up

Install Go 1.27.1 and Bun.

```sh
./scripts/bootstrap
bun run install:licenses
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

Run every gate before requesting review.

```sh
bun run check:all
bun run build
```

SDK commits are exempt from LOC reductions and per-commit coverage gains.

Preserve coverage and report all four metrics against the verified parent.

SDK checks have no time limit; builds run separately.

Full checks run the complete 2-minute fuzz campaign directly.
SDK elapsed time remains visible and never blocks delivery by itself.

Install official CodeQL CLI 2.27.0 and verify its published archive checksum.
Set `XQUIK_CODEQL_BIN` to the executable or use its default location:
`~/.cache/xquik-codeql/2.27.0/codeql/codeql`.
Download the pinned query pack before validation:

```sh
"$XQUIK_CODEQL_BIN" pack download codeql/go-queries@1.6.10
```

Each full check extracts a fresh database and runs security-extended queries.
Findings, extraction errors, and incomplete results block release.

Add regression tests for every fixed defect.

## Submit changes

Use clear Conventional Commit subjects and sign commits with `git commit --signoff`.
Verify published modules with `bun run release:verify vVERSION COMMIT_SHA`.
The command uses a fresh module cache and the public Go checksum database.
It checks the exact version and source commit, then installs and constructs a client.
JSON output identifies success or failure. Raw module evidence remains available.

Prepare releases with `bun run release:prepare` from a clean, committed checkout.
It runs complete validation and creates a reproducible source archive.
After merging, create the matching `vVERSION` tag on that commit.
Try the release workflow when Actions is available; otherwise run `bun run release:publish`.
The local publisher uses locked Sigstore 4.5.0 tooling and GitHub CLI authentication.
Use uv 0.12.13; the command selects Python 3.14.7 automatically.
Set `XQUIK_RELEASE_IDENTITY` and `XQUIK_RELEASE_ISSUER` to the authorized signer.
Signing uses Sigstore authentication and records local build provenance publicly.
Local attestations identify the signer, not a GitHub Actions workflow.
The publisher verifies signatures before uploading and rejects conflicting existing artifacts.
It verifies public module installation before publishing the GitHub draft.
Retry the same command after partial publication; matching artifacts are retained.
Follow the shared [review policy][review-policy].

## Report Security Issues

Never disclose suspected vulnerabilities in public issues.

Follow [SECURITY.md](SECURITY.md) for private reporting.

[contribution-policy]: https://github.com/Xquik-dev/.github/blob/main/CONTRIBUTING.md
[review-policy]: https://github.com/Xquik-dev/.github/blob/main/REVIEWING.md

Xquik is an independent third-party service. Not affiliated with X Corp. "Twitter" and "X" are trademarks of X Corp.
