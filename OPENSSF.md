# OpenSSF Best Practices Evidence

This register tracks the Gold assessment for this repository.

The official entry is [bestpractices.dev project 13734][badge].

Assessment date: 2026-07-23.

## Eligibility

This active, released public SDK qualifies for the OpenSSF Best Practices badge.
No OpenSSF-defined ineligibility applies.

## Verified Technical Controls

| Area | Evidence |
| --- | --- |
| License | Apache-2.0, BSD-3-Clause, and REUSE 3.3 metadata |
| Contribution process | DCO sign-off and independent review rules |
| Governance | Public roles, decisions, releases, and continuity policy |
| Security reporting | Private reporting, response targets, boundaries, and threat model |
| Tests | Full tests, regression tests, race detection, and weekly fuzzing |
| Statement coverage | `./scripts/coverage` enforces 90% |
| Branch coverage | `./scripts/branch-coverage` enforces 80% |
| Static analysis | `go vet` and CodeQL security-extended queries |
| Dependency review | Dependabot, `go mod verify`, and `govulncheck` |
| Licensing gate | Pinned REUSE action checks every repository file |
| Reproducibility | 2 normalized module archives must have identical hashes |
| CI | Pull requests and pushes run pinned, least-privilege workflows |
| Two-factor authentication | The Xquik-dev organization requires 2FA |

The current Go 1.26.6 statement result is 92.53%.

The current Go 1.26.6 branch result is 81.15%.

## Branch Measurement

Go's standard coverage tool does not report branches.
The gate pins the MIT-licensed `gocove` beta at commit `a0dceb9dadca`.
That tool instruments every package and emits a native Go profile.

The pinned release can join nested ranges to an outer block.
Executed switch and else branches can then appear uncovered.
The verifier requires strict overlap with positive profile ranges.
It rejects boundary-only parents and rebuilds packages without cached catalogs.
The gate includes all 1,236 reported production branches.

## Release Provenance

Release v0.18.0 includes a reproducible source archive and Sigstore bundle.
GitHub verification confirms the digest, workflow, tag & hosted runner.
The archive digest matches its attached SLSA provenance subject.

## Outstanding Gold Blockers

Human and organizational evidence remains incomplete.

Do not claim Gold while any mandatory criterion remains unmet.

| Gold Requirement | Current Evidence | Required Action |
| --- | --- | --- |
| Access continuity | Public evidence does not prove 2 release-capable maintainers | Grant and verify another maintainer's access |
| Bus factor | Git history shows one significant contributor | Add another significant contributor |
| Unassociated contributors | Fewer than 2 qualifying contributors are independent | Accept qualifying external contributions |
| Independent review | History does not prove 50% qualifying review coverage | Require and record independent reviews |
| Human security review | No completed review exists within 5 years | Commission and publish a scoped review |

Gold eligibility still requires review by a different human.

## Maintenance

Run these evidence commands before releases:

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
gh attestation verify ARCHIVE \
  --repo Xquik-dev/x-twitter-scraper-go \
  --signer-workflow Xquik-dev/x-twitter-scraper-go/.github/workflows/release-provenance.yml
```

Reassess the register before every major release.

Update bestpractices.dev only with public evidence.

[badge]: https://www.bestpractices.dev/projects/13734

Xquik is an independent third-party service. Not affiliated with X Corp. "Twitter" and "X" are trademarks of X Corp.
