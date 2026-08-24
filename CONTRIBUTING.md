# Contributing

Read [GOVERNANCE.md](GOVERNANCE.md) before proposing major changes.

Follow the shared [Xquik contribution policy][contribution-policy].

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

## Submit changes

Use clear Conventional Commit subjects and sign commits with `git commit --signoff`.
Follow the shared [review policy][review-policy].

## Report Security Issues

Never disclose suspected vulnerabilities in public issues.

Follow [SECURITY.md](SECURITY.md) for private reporting.

[contribution-policy]: https://github.com/Xquik-dev/.github/blob/main/CONTRIBUTING.md
[review-policy]: https://github.com/Xquik-dev/.github/blob/main/REVIEWING.md

Xquik is an independent third-party service. Not affiliated with X Corp. "Twitter" and "X" are trademarks of X Corp.
