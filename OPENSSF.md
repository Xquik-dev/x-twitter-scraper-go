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
| Tests | Full tests, regression tests, race detection, and fuzzing |
| Statement coverage | `scripts/collect-coverage` measures statement coverage |
| Branch coverage | `scripts/collect-branches` measures branch coverage |
| Static analysis | `go vet` and CodeQL security-extended queries |
| Dependency review | `go mod verify` and `scripts/dependency-audit`, which runs `govulncheck` |
| Licensing gate | `bin/check-local` runs REUSE lint on every repository file |
| Reproducibility | 2 normalized module archives must have identical hashes |
| Local checks | `bin/check-local` runs every check; no GitHub Actions workflow runs |
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

Releases after v0.18.1 come from `bin/release-local` on a maintainer's machine.
They carry no GitHub workflow attestation.

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

Run the evidence checks before releases:

```sh
bin/check-local
```

`bin/release-local` runs them again before it tags and publishes a release.

Reassess the register before every major release.

Update bestpractices.dev only with public evidence.

[badge]: https://www.bestpractices.dev/projects/13734

Xquik is an independent third-party service. Not affiliated with X Corp. "Twitter" and "X" are trademarks of X Corp.
