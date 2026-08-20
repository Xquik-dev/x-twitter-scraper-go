# Contributing

Read [GOVERNANCE.md](GOVERNANCE.md) before proposing major changes.

## Set Up

Install Go 1.26.6 or newer.

```sh
./scripts/bootstrap
./scripts/lint
./scripts/test
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
go mod verify
./scripts/lint
./scripts/test
./scripts/coverage
./scripts/branch-coverage
go test -race ./...
go run golang.org/x/vuln/cmd/govulncheck@v1.6.0 ./...
uvx --from reuse==5.1.1 reuse lint
./scripts/reproducible-build
```

Statement coverage must remain at least 90%.

Branch coverage must remain at least 80%.

Add regression tests for every fixed defect.

## Submit Changes

Keep pull requests focused and explain user-visible behavior.

Link relevant issues and public API contracts.

Use clear Conventional Commit subjects when practical.

Sign every commit with the Developer Certificate of Origin.

```sh
git commit --signoff
```

The sign-off confirms the [Developer Certificate of Origin](https://developercertificate.org/).

Another human must review maintainer-authored, nontrivial changes.

Reviewers follow the shared [review policy][review-policy].

Address every review comment before merging.

## Report Security Issues

Never disclose suspected vulnerabilities in public issues.

Follow [SECURITY.md](SECURITY.md) for private reporting.

[review-policy]: https://github.com/Xquik-dev/.github/blob/main/REVIEWING.md

Xquik is an independent third-party service. Not affiliated with X Corp. "Twitter" and "X" are trademarks of X Corp.
