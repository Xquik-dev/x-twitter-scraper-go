# Agent instructions

Every commit reduces handwritten test LOC and non-test code LOC separately against its parent.
Both deltas must be negative. Neither category offsets the other.
Count every language and location, including shell scripts and development tooling.
Test-only helpers and scripts count as tests. Build, runtime, and maintenance code count as non-test code.
Count each line once. Report both categories' before, after, and delta using consistent rules.
Use useful simplifications within each category to offset additions.
Exclude formatting, minification, generation, moves, renames, reclassification, and deleted valuable tests or docs.
Preserve coverage, assertions, guards, diagnostics, and behavior.
Run `bun run check:all` before every commit, push, or deployment.
