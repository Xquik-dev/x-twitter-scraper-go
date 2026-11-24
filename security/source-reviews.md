# Approved source findings

The user approved these 123 exact findings on 2026-09-12.
On 2026-10-13, `scripts/move_source_reviews.go` moved 68 of them to new line numbers & file hashes.
It moves a finding only when its code & surrounding lines are unchanged, and it changed no reason.
On 2026-10-14, the comment sweep moved 80 more the same way. It also re-pinned the anonymous field decoder finding in `internal/apijson/decoder.go` by hand, because only its `// ignore errors` comment left the excerpt. Its reason is unchanged.
3 records name `cmd/reproducible-build/upload_test.go`, which no longer exists, so nothing can match them.
On 2026-10-13, 5 findings in `internal/scripttest` were reviewed and added. Each runs only in tests, with input the test controls.
On 2026-10-15, the CodeQL 2.27.2 update changed 1 line in `internal/scripttest/codeql_security_test.go`. The same script moved its 2 findings to the new file hash. Their code & reasons are unchanged.
On 2026-10-15, 4 findings in `cmd/api-coverage/main.go` were reviewed and added. The API coverage check reads the spec file or URL that its runner chooses.
On 2026-10-15, the API port moved 46 findings the same way, 4 in `xtweet.go`, 6 in `xuser.go`, 4 in `xlist.go` and 32 in `shared/shared.go`. Their code & reasons are unchanged.
On 2026-10-15, 4 findings in `shared/shared.go` were reviewed and added. They are the As variant conversions of the 2 Community unions, with the reason the other generated As methods carry.
On 2026-10-15, 24 findings in `xanalysis.go` were reviewed and added. They are the As variant conversions of the 2 answer unions in each of the 6 analysis responses, with the reason the other generated As methods carry.
On 2026-10-16, the API fields check moved the spec loader from `cmd/api-coverage/main.go` to `internal/apicheck/apicheck.go`. The 4 findings of the old file left with its code. The 1 finding gosec reports in the new file was reviewed with the reason the old file read finding carried.
On 2026-10-16, the field fixes moved 53 findings the same way in `shared/shared.go`, `xaccount.go`, `xlist.go`, `xtweet.go` and `xuser.go`. Their code & reasons are unchanged.
On 2026-10-17, the `OverrideVerifiedYear` field addition moved 36 findings in `shared/shared.go` to its new source hash.
On 2026-10-17, the embedded post fields and the bulk retry `AccountIDs` moved 39 findings the same way, in `shared/shared.go` and `xaccount.go`. Their code & reasons are unchanged.
On 2026-10-17, the bulk retry doc comment fix moved 3 findings in `xaccount.go` to its new file hash. Their code & reasons are unchanged.
On 2026-10-22, the `softIntervention` field, the `X-Xquik-SDK` header and its test moved 46 findings the same way, 32 in `shared/shared.go`, 4 in `xtweet.go`, 5 in `internal/requestconfig/requestconfig.go` and 1 in `client_test.go`. Their code & reasons are unchanged.
On 2026-10-22, the `SoftInterventionText` type moved 36 findings in `shared/shared.go` the same way. Their code & reasons are unchanged.
On 2026-11-03, the activity reads, the shadowban check and the post and profile fields moved 46 findings the same way, 36 in `shared/shared.go`, 6 in `xuser.go` and 4 in `xtweet.go`. Their code & reasons are unchanged.
On 2026-11-03, 4 findings in `activity.go` were reviewed and added. They are the As variant conversions of the results and credits unions of `ActivityLogRow`, with the reason the other generated As methods carry.
On 2026-11-03, the 5 shadowban checks moved 6 findings in `xuser.go` to its new file hash, through the same script. Their code, lines & reasons are unchanged.
On 2026-11-05, the API's new fields moved 34 findings the same way, 24 in `xanalysis.go`, 4 in `xtweet.go` and 6 in `xuser.go`. Their code & reasons are unchanged.
On 2026-11-06, the bookmark, block and mute writes moved 10 findings the same way, 4 in `xtweet.go` and 6 in `xuser.go`. Their code & reasons are unchanged.
On 2026-11-07, the insights & shadowban `limit` params moved 6 findings in `xuser.go` to its new file hash, through the same script. Their code, lines & reasons are unchanged.
On 2026-11-08, the image and topic filters moved 46 findings the same way, 36 in `shared/shared.go`, 4 in `xtweet.go` and 6 in `xuser.go`. Their code & reasons are unchanged.
On 2026-11-08, the follower audit moved 6 findings in `xuser.go` the same way. Their code & reasons are unchanged.
On 2026-11-08, the `Prefer` header on every write moved 10 findings the same way, 4 in `xtweet.go` and 6 in `xuser.go`. Their code & reasons are unchanged.
On 2026-11-08, the likers and batch comments moved 10 findings the same way, 4 in `xtweet.go` and 6 in `xuser.go`. Their code & reasons are unchanged.
On 2026-11-09, the X account connect links moved 7 findings the same way, 4 in `xaccountconnectionattempt.go` and 3 in `xaccount.go`. Their code & reasons are unchanged.
On 2026-11-09, 3 findings in `xaccount.go` were reviewed and added. They are the As variant conversions of `XAccountReauthResponseUnion`, with the reason the other generated As methods carry.
On 2026-11-13, the `hiddenTimeline` and `windowComplete` fields moved 50 findings the same way, 36 in `shared/shared.go`, 4 in `xtweet.go`, 6 in `xuser.go` and 4 in `xlist.go`, through a port of the same script whose fingerprints match all 170 records. Their code & reasons are unchanged.
On 2026-11-13, the `OwnThreads` mentions param moved 6 findings in `xuser.go` the same way. Their code & reasons are unchanged.
On 2026-11-13, the analysis `Topic`, `ImageShows`, `TopicFilter` and `ImageFilter` fields moved 24 findings in `xanalysis.go` the same way. Their code & reasons are unchanged.
On 2026-11-17, the analysis sources, `UnprocessedIDs` and `EndReason` fields and the CodeQL version floor moved 76 findings the same way, 24 in `xanalysis.go`, 36 in `shared/shared.go`, 4 in `xtweet.go`, 6 in `xuser.go`, 4 in `xlist.go` and 2 in `internal/scripttest/codeql_security_test.go`. Their code & reasons are unchanged.
On 2026-11-17, the release fixture's file overrides moved 8 findings in `internal/scripttest/release_local_test.go` and `internal/scripttest/temp_cleanup_test.go` the same way, and keeping the cleanup cases' check-local runs local moved 6 in `internal/scripttest/temp_cleanup_test.go` to its new file hash. Their code & reasons are unchanged.
On 2026-11-24, the `communityUnavailable` and `rankings` fields moved 40 findings the same way, 36 in `shared/shared.go` and 4 in `xtweet.go`. Their code & reasons are unchanged.
The temp cleanup tests have 6 reviewed findings for test-controlled script runners, executable fixtures, retained coverage files and restoring owner access to isolated module caches. Their records match the current source hashes and line numbers.
The checker retains every report and rejects new or changed findings.
The index contains SHA-256 digests of Go `%q` formatting for 8 strings.
Their order is rule, file, line, code, source hash, severity, confidence, details.
These records preserve the complete reviewed source excerpts and reasons.

## G115 in internal/encoding/json/decode_upstream_test.go:335

Upstream test-only custom byte marshaler. Callers supply bounded literal fixture values. Not used to validate security-sensitive input or included in library builds. This review does not claim the test helper validates arbitrary integer ranges.

Source SHA-256: `c9e5bf874a2aa1daf49b0ff1253ba34d4c0f77099dc6ef733e271f9117cf565c`
Severity: HIGH. Confidence: MEDIUM.
Finding: integer overflow conversion int64 -> byte

```go
334: 	}
335: 	*b = byteWithMarshalText(i)
336: 	return nil
```

## G115 in internal/encoding/json/decode_upstream_test.go:307

Upstream test-only custom byte marshaler. Callers supply bounded literal fixture values. Not used to validate security-sensitive input or included in library builds. This review does not claim the test helper validates arbitrary integer ranges.

Source SHA-256: `c9e5bf874a2aa1daf49b0ff1253ba34d4c0f77099dc6ef733e271f9117cf565c`
Severity: HIGH. Confidence: MEDIUM.
Finding: integer overflow conversion int64 -> byte

```go
306: 	}
307: 	*b = byteWithMarshalJSON(i)
308: 	return nil
```

## G115 in internal/encoding/json/decode_upstream_test.go:160

Upstream test-only custom byte marshaler. Callers supply bounded literal fixture values. Not used to validate security-sensitive input or included in library builds. This review does not claim the test helper validates arbitrary integer ranges.

Source SHA-256: `c9e5bf874a2aa1daf49b0ff1253ba34d4c0f77099dc6ef733e271f9117cf565c`
Severity: HIGH. Confidence: MEDIUM.
Finding: integer overflow conversion int -> uint8

```go
159: 	}
160: 	*u8 = u8marshal(n)
161: 	return nil
```

## G404 in internal/requestconfig/requestconfig.go:319

Bounded retry jitter, not a credential, key, authentication value, or security token. The arithmetic overflow was repaired and verified on ARM64 and AMD64. Cryptographic unpredictability is not a requirement for jitter.

Source SHA-256: `568a3ad6f55cae80a238e5a9819f90666a613b25966335eff089d01092c575e8`
Severity: HIGH. Confidence: MEDIUM.
Finding: Use of weak random number generator (math/rand or math/rand/v2 instead of crypto/rand)

```go
318: 	delay := 500 * time.Millisecond << min(max(retryCount, 0), 4)
319: 	return delay - time.Duration(rand.Int63n(int64(delay/4)))
320: }
```

## G404 in internal/encoding/json/scanner_upstream_test.go:292

Retry jitter or randomized test fixture data. Values do not protect credentials, authenticate requests, or generate keys.

Source SHA-256: `7c20815e7d26acca191b85bf6c2ce560860b619531666781612f645f391be88c`
Severity: HIGH. Confidence: MEDIUM.
Finding: Use of weak random number generator (math/rand or math/rand/v2 instead of crypto/rand)

```go
291: func genMap(n int) map[string]any {
292: 	f := int(math.Abs(rand.NormFloat64()) * math.Min(10, float64(n/2)))
293: 	if f > n {
```

## G404 in internal/encoding/json/scanner_upstream_test.go:277

Retry jitter or randomized test fixture data. Values do not protect credentials, authenticate requests, or generate keys.

Source SHA-256: `7c20815e7d26acca191b85bf6c2ce560860b619531666781612f645f391be88c`
Severity: HIGH. Confidence: MEDIUM.
Finding: Use of weak random number generator (math/rand or math/rand/v2 instead of crypto/rand)

```go
276: func genArray(n int) []any {
277: 	f := int(math.Abs(rand.NormFloat64()) * math.Min(10, float64(n/2)))
278: 	if f > n {
```

## G404 in internal/encoding/json/scanner_upstream_test.go:267

Retry jitter or randomized test fixture data. Values do not protect credentials, authenticate requests, or generate keys.

Source SHA-256: `7c20815e7d26acca191b85bf6c2ce560860b619531666781612f645f391be88c`
Severity: HIGH. Confidence: MEDIUM.
Finding: Use of weak random number generator (math/rand or math/rand/v2 instead of crypto/rand)

```go
266: 	for i := range c {
267: 		f := math.Abs(rand.NormFloat64()*64 + 32)
268: 		if f > 0x10ffff {
```

## G404 in internal/encoding/json/scanner_upstream_test.go:264

Retry jitter or randomized test fixture data. Values do not protect credentials, authenticate requests, or generate keys.

Source SHA-256: `7c20815e7d26acca191b85bf6c2ce560860b619531666781612f645f391be88c`
Severity: HIGH. Confidence: MEDIUM.
Finding: Use of weak random number generator (math/rand or math/rand/v2 instead of crypto/rand)

```go
263: func genString(stddev float64) string {
264: 	n := int(math.Abs(rand.NormFloat64()*stddev + stddev/2))
265: 	c := make([]rune, n)
```

## G404 in internal/encoding/json/scanner_upstream_test.go:256

Retry jitter or randomized test fixture data. Values do not protect credentials, authenticate requests, or generate keys.

Source SHA-256: `7c20815e7d26acca191b85bf6c2ce560860b619531666781612f645f391be88c`
Severity: HIGH. Confidence: MEDIUM.
Finding: Use of weak random number generator (math/rand or math/rand/v2 instead of crypto/rand)

```go
255: 	case 1:
256: 		return rand.NormFloat64()
257: 	case 2:
```

## G404 in internal/encoding/json/scanner_upstream_test.go:254

Retry jitter or randomized test fixture data. Values do not protect credentials, authenticate requests, or generate keys.

Source SHA-256: `7c20815e7d26acca191b85bf6c2ce560860b619531666781612f645f391be88c`
Severity: HIGH. Confidence: MEDIUM.
Finding: Use of weak random number generator (math/rand or math/rand/v2 instead of crypto/rand)

```go
253: 	case 0:
254: 		return rand.Intn(2) == 0
255: 	case 1:
```

## G404 in internal/encoding/json/scanner_upstream_test.go:252

Retry jitter or randomized test fixture data. Values do not protect credentials, authenticate requests, or generate keys.

Source SHA-256: `7c20815e7d26acca191b85bf6c2ce560860b619531666781612f645f391be88c`
Severity: HIGH. Confidence: MEDIUM.
Finding: Use of weak random number generator (math/rand or math/rand/v2 instead of crypto/rand)

```go
251: 	}
252: 	switch rand.Intn(3) {
253: 	case 0:
```

## G404 in internal/encoding/json/scanner_upstream_test.go:245

Retry jitter or randomized test fixture data. Values do not protect credentials, authenticate requests, or generate keys.

Source SHA-256: `7c20815e7d26acca191b85bf6c2ce560860b619531666781612f645f391be88c`
Severity: HIGH. Confidence: MEDIUM.
Finding: Use of weak random number generator (math/rand or math/rand/v2 instead of crypto/rand)

```go
244: 	if n > 1 {
245: 		switch rand.Intn(2) {
246: 		case 0:
```

## G101 in guestwallet.go:385

Public generated credential-storage notice, not a credential value.

Source SHA-256: `5c647a919b599d495c0d68381ebbc9e1e5d9ea8012963508a57296190c05adc1`
Severity: HIGH. Confidence: LOW.
Finding: Potential hardcoded credentials

```go
384: const (
385: 	GuestWalletTopupResponseCredentialNoticeStoreAPIKeyAndTheIdempotencyKeySecurelyBeforeSharingCheckoutURLNoEmailRecoveryIsAvailable GuestWalletTopupResponseCredentialNotice = "Store api_key and the Idempotency-Key securely before sharing checkout_url. No email recovery is available."
386: )
```

## G703 in cmd/reproducible-build/main.go:71

Explicit local CLI output paths. No remote path source or claimed confinement boundary.

Source SHA-256: `1ab18bc767ae2cdc5a96bf94aa3822ec58fd5b44e3bf7bab31647b83fd225487`
Severity: HIGH. Confidence: HIGH.
Finding: Path traversal via taint analysis

```go
70: 		}
71: 		if err := os.WriteFile(args[0], first, 0o644); err != nil {
72: 			return err
```

## G703 in cmd/reproducible-build/main.go:68

Explicit local CLI output paths. No remote path source or claimed confinement boundary.

Source SHA-256: `1ab18bc767ae2cdc5a96bf94aa3822ec58fd5b44e3bf7bab31647b83fd225487`
Severity: HIGH. Confidence: HIGH.
Finding: Path traversal via taint analysis

```go
67: 	if len(args) == 1 {
68: 		if err := os.MkdirAll(filepath.Dir(args[0]), 0o755); err != nil {
69: 			return err
```

## G703 in cmd/coverage-gain/main.go:58

Explicit local CLI output paths. No remote path source or claimed confinement boundary.

Source SHA-256: `62def279c8417dc1ba3c8e05cd2a05e5ff853ef1290e6852437c4fe1533c2baa`
Severity: HIGH. Confidence: HIGH.
Finding: Path traversal via taint analysis

```go
57: 	}
58: 	if err := os.WriteFile(filepath.Join(directory, "summary.json"), data, 0o600); err != nil {
59: 		return err
```

## G705 in cmd/reproducible-build/upload_test.go:36

Local upload test returns a JSON fixture to a CLI client. No browser rendering or HTML sink in the tested flow.

Source SHA-256: `03b01546690072902765e00cb498db6dc822589cd5d89bc891d9e4fc5e82150a`
Severity: MEDIUM. Confidence: HIGH.
Finding: XSS via taint analysis

```go
35: 					}
36: 					fmt.Fprintf(w, `{"url":"http://%s/upload?signed-fixture"}`, r.Host)
37: 					return
```

## G204 in cmd/reproducible-build/upload_test.go:47

Fixed executable and argument array. Build helpers consume local Git arguments; upload test invokes the fixed repository script. No shell interpolation of remote data.

Source SHA-256: `03b01546690072902765e00cb498db6dc822589cd5d89bc891d9e4fc5e82150a`
Severity: MEDIUM. Confidence: HIGH.
Finding: Subprocess launched with variable

```go
46: 			defer server.Close()
47: 			command := exec.Command("bash", "-x", script)
48: 			command.Dir = directory
```

## G204 in cmd/reproducible-build/main_test.go:21

Fixed executable and argument array. Build helpers consume local Git arguments; upload test invokes the fixed repository script. No shell interpolation of remote data.

Source SHA-256: `b39ec9ad06af8abe5c33d26e19df7080b69fdeab5ca59d298a7f5d10c5ba38be`
Severity: MEDIUM. Confidence: HIGH.
Finding: Subprocess launched with variable

```go
20: 		t.Helper()
21: 		data, err := exec.Command("git", args...).CombinedOutput()
22: 		if err != nil {
```

## G204 in cmd/reproducible-build/main.go:23

Fixed executable and argument array. Build helpers consume local Git arguments; upload test invokes the fixed repository script. No shell interpolation of remote data.

Source SHA-256: `1ab18bc767ae2cdc5a96bf94aa3822ec58fd5b44e3bf7bab31647b83fd225487`
Severity: MEDIUM. Confidence: HIGH.
Finding: Subprocess launched with variable

```go
22: 	git := func(args ...string) ([]byte, error) {
23: 		command := exec.Command("git", args...)
24: 		command.Env = append(os.Environ(), "GIT_INDEX_FILE="+filepath.Join(directory, "index"), "TZ=UTC")
```

## G204 in cmd/coverage-gain/main.go:22

Fixed executable and argument array. Build helpers consume local Git arguments; upload test invokes the fixed repository script. No shell interpolation of remote data.

Source SHA-256: `62def279c8417dc1ba3c8e05cd2a05e5ff853ef1290e6852437c4fe1533c2baa`
Severity: MEDIUM. Confidence: HIGH.
Finding: Subprocess launched with variable

```go
21: func git(args ...string) (string, error) {
22: 	command := exec.Command("git", args...)
23: 	command.Env = append(os.Environ(), "LC_ALL=C")
```

## G204 in cmd/coverage-gain/main_test.go:26

Fixed executable and argument array. Build helpers consume local Git arguments; upload test invokes the fixed repository script. No shell interpolation of remote data.

Source SHA-256: `0aafc65741292ff8e02f31138dd3a6c2dd63fa5f4e34eef7d7c66ec0e1a11370`
Severity: MEDIUM. Confidence: HIGH.
Finding: Subprocess launched with a potential tainted input or cmd arguments

```go
25: 		t.Helper()
26: 		command := exec.Command("git", append([]string{"-c", "user.name=Coverage test", "-c", "user.email=coverage@example.test", "-c", "commit.gpgsign=false"}, args...)...)
27: 		if data, err := command.CombinedOutput(); err != nil {
```

## G304 in cmd/coverage-gain/main_test.go:84

Test reads its own summary artifact within a temporary directory.

Source SHA-256: `0aafc65741292ff8e02f31138dd3a6c2dd63fa5f4e34eef7d7c66ec0e1a11370`
Severity: MEDIUM. Confidence: HIGH.
Finding: Potential file inclusion via variable

```go
83: 	}
84: 	report, err := os.ReadFile(filepath.Join(directory, "summary.json"))
85: 	if err != nil || !strings.Contains(string(report), `"Covered": 1`) {
```

## G107 in internal/testutil/testutil.go:15

Test driver selects the server URL; no remote user supplies this URL. HTTP body leak separately reproduced and repaired.

Source SHA-256: `1ae5148653d265ee654248ed19503d781e5443a0d98386e78949501e669ba194`
Severity: MEDIUM. Confidence: MEDIUM.
Finding: Potential HTTP request made with variable url

```go
14: func CheckTestServer(t *testing.T, url string) bool {
15: 	if response, err := http.Get(url); err != nil {
16: 		const SKIP_MOCK_TESTS = "SKIP_MOCK_TESTS"
```

## G302 in cmd/reproducible-build/main_test.go:37

Executable-bit fixture intentionally verifies preservation of executable files in module archives.

Source SHA-256: `b39ec9ad06af8abe5c33d26e19df7080b69fdeab5ca59d298a7f5d10c5ba38be`
Severity: MEDIUM. Confidence: HIGH.
Finding: Expect file permissions to be 0600 or less

```go
36: 	}
37: 	if err := os.Chmod("executable", 0o755); err != nil {
38: 		t.Fatal(err)
```

## G301 in cmd/reproducible-build/main.go:68

Caller-selected directory for source archive. Archive intended for publication; no secret file permission boundary.

Source SHA-256: `1ab18bc767ae2cdc5a96bf94aa3822ec58fd5b44e3bf7bab31647b83fd225487`
Severity: MEDIUM. Confidence: HIGH.
Finding: Expect directory permissions to be 0750 or less

```go
67: 	if len(args) == 1 {
68: 		if err := os.MkdirAll(filepath.Dir(args[0]), 0o755); err != nil {
69: 			return err
```

## G306 in cmd/reproducible-build/upload_test.go:27

Source archive or public fixture contents. Temporary fixture directories use testing.TempDir isolation.

Source SHA-256: `03b01546690072902765e00cb498db6dc822589cd5d89bc891d9e4fc5e82150a`
Severity: MEDIUM. Confidence: HIGH.
Finding: Expect WriteFile permissions to be 0600 or less

```go
26: 			directory := t.TempDir()
27: 			if err := os.WriteFile(filepath.Join(directory, "fixture.go"), []byte("package fixture\n"), 0o644); err != nil {
28: 				t.Fatal(err)
```

## G306 in cmd/reproducible-build/main_test.go:29

Source archive or public fixture contents. Temporary fixture directories use testing.TempDir isolation.

Source SHA-256: `b39ec9ad06af8abe5c33d26e19df7080b69fdeab5ca59d298a7f5d10c5ba38be`
Severity: MEDIUM. Confidence: HIGH.
Finding: Expect WriteFile permissions to be 0600 or less

```go
28: 		t.Helper()
29: 		if err := os.WriteFile(name, []byte(value), 0o644); err != nil {
30: 			t.Fatal(err)
```

## G306 in cmd/reproducible-build/main.go:71

Source archive or public fixture contents. Temporary fixture directories use testing.TempDir isolation.

Source SHA-256: `1ab18bc767ae2cdc5a96bf94aa3822ec58fd5b44e3bf7bab31647b83fd225487`
Severity: MEDIUM. Confidence: HIGH.
Finding: Expect WriteFile permissions to be 0600 or less

```go
70: 		}
71: 		if err := os.WriteFile(args[0], first, 0o644); err != nil {
72: 			return err
```

## G103 in internal/apijson/decoder.go:603

Only call assigns node.Raw into the generated JSON.raw metadata field. The pointer comes from the addressable current Go reflect.Value, without pointer arithmetic or stored unsafe pointers. reflect.Set performs typed assignment. RawJSON model tests cover metadata preservation.

Source SHA-256: `fdffb7b3c66a4b30e15a404f43391da29cf5c41df8dafd8ebd275f870f0810c4`
Severity: LOW. Confidence: HIGH.
Finding: Use of unsafe calls should be audited

```go
602: func setUnexportedField(field reflect.Value, value any) {
603: 	reflect.NewAt(field.Type(), unsafe.Pointer(field.UnsafeAddr())).Elem().Set(reflect.ValueOf(value))
604: }
```

## G104 in xuser.go:1205

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `0f32ca5693357ab12fb195b7cdbf74430eb86d8108551148a4e74e5969a2fbba`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
1204: func (u XUserGetVerifiedFollowersResponseUnion) AsXUserGetVerifiedFollowersResponseUserListCoverageResponse() (v XUserGetVerifiedFollowersResponseUserListCoverageResponse) {
1205: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
1206: 	return
```

## G104 in xuser.go:1200

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `0f32ca5693357ab12fb195b7cdbf74430eb86d8108551148a4e74e5969a2fbba`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
1199: func (u XUserGetVerifiedFollowersResponseUnion) AsPaginatedUsers() (v shared.PaginatedUsers) {
1200: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
1201: 	return
```

## G104 in xuser.go:1014

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `0f32ca5693357ab12fb195b7cdbf74430eb86d8108551148a4e74e5969a2fbba`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
1013: func (u XUserGetFollowingResponseUnion) AsXUserGetFollowingResponseUserListCoverageResponse() (v XUserGetFollowingResponseUserListCoverageResponse) {
1014: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
1015: 	return
```

## G104 in xuser.go:1009

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `0f32ca5693357ab12fb195b7cdbf74430eb86d8108551148a4e74e5969a2fbba`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
1008: func (u XUserGetFollowingResponseUnion) AsPaginatedUsers() (v shared.PaginatedUsers) {
1009: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
1010: 	return
```

## G104 in xuser.go:825

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `0f32ca5693357ab12fb195b7cdbf74430eb86d8108551148a4e74e5969a2fbba`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
824: func (u XUserGetFollowersResponseUnion) AsXUserGetFollowersResponseUserListCoverageResponse() (v XUserGetFollowersResponseUserListCoverageResponse) {
825: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
826: 	return
```

## G104 in xuser.go:820

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `0f32ca5693357ab12fb195b7cdbf74430eb86d8108551148a4e74e5969a2fbba`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
819: func (u XUserGetFollowersResponseUnion) AsPaginatedUsers() (v shared.PaginatedUsers) {
820: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
821: 	return
```

## G104 in xtweet.go:3090

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `5432ba4d19055df1215634cca40027e5dea11fc4d9a82088c15a4c36c02a3f23`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
3089: func (u XTweetSearchResponseUnion) AsXTweetSearchResponseTweetSearchCoverageResponse() (v XTweetSearchResponseTweetSearchCoverageResponse) {
3090: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
3091: 	return
```

## G104 in xtweet.go:3085

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `5432ba4d19055df1215634cca40027e5dea11fc4d9a82088c15a4c36c02a3f23`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
3084: func (u XTweetSearchResponseUnion) AsPaginatedTweets() (v shared.PaginatedTweets) {
3085: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
3086: 	return
```

## G104 in xtweet.go:2081

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `5432ba4d19055df1215634cca40027e5dea11fc4d9a82088c15a4c36c02a3f23`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
2080: func (u TweetDetailSportsContextScheduledAtMsUnion) AsString() (v string) {
2081: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
2082: 	return
```

## G104 in xtweet.go:2076

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `5432ba4d19055df1215634cca40027e5dea11fc4d9a82088c15a4c36c02a3f23`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
2075: func (u TweetDetailSportsContextScheduledAtMsUnion) AsFloat() (v float64) {
2076: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
2077: 	return
```

## G104 in xlist.go:342

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `91a0eaa45b5d90f4b57d49c70f0d99f9e2931e3796f809bec22cf934b941bedc`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
341: func (u XListGetMembersResponseUnion) AsXListGetMembersResponseUserListCoverageResponse() (v XListGetMembersResponseUserListCoverageResponse) {
342: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
343: 	return
```

## G104 in xlist.go:337

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `91a0eaa45b5d90f4b57d49c70f0d99f9e2931e3796f809bec22cf934b941bedc`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
336: func (u XListGetMembersResponseUnion) AsPaginatedUsers() (v shared.PaginatedUsers) {
337: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
338: 	return
```

## G104 in xlist.go:153

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `91a0eaa45b5d90f4b57d49c70f0d99f9e2931e3796f809bec22cf934b941bedc`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
152: func (u XListGetFollowersResponseUnion) AsXListGetFollowersResponseUserListCoverageResponse() (v XListGetFollowersResponseUserListCoverageResponse) {
153: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
154: 	return
```

## G104 in xlist.go:148

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `91a0eaa45b5d90f4b57d49c70f0d99f9e2931e3796f809bec22cf934b941bedc`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
147: func (u XListGetFollowersResponseUnion) AsPaginatedUsers() (v shared.PaginatedUsers) {
148: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
149: 	return
```

## G104 in xaccountconnectionattempt.go:152

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `f033ef2a3d7e7f4381a11f30b624c1f273e89792fe8c688b37b8667f49e342b3`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
151: func (u XAccountConnectionAttemptGetResponseUnion) AsRequiresEmailCode() (v XAccountConnectionAttemptGetResponseRequiresEmailCode) {
152: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
153: 	return
```

## G104 in xaccountconnectionattempt.go:147

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `f033ef2a3d7e7f4381a11f30b624c1f273e89792fe8c688b37b8667f49e342b3`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
146: func (u XAccountConnectionAttemptGetResponseUnion) AsFailed() (v XAccountConnectionAttemptGetResponseFailed) {
147: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
148: 	return
```

## G104 in xaccountconnectionattempt.go:142

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `f033ef2a3d7e7f4381a11f30b624c1f273e89792fe8c688b37b8667f49e342b3`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
141: func (u XAccountConnectionAttemptGetResponseUnion) AsSuccess() (v XAccountConnectionAttemptGetResponseSuccess) {
142: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
143: 	return
```

## G104 in xaccountconnectionattempt.go:137

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `f033ef2a3d7e7f4381a11f30b624c1f273e89792fe8c688b37b8667f49e342b3`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
136: func (u XAccountConnectionAttemptGetResponseUnion) AsPending() (v XAccountConnectionAttemptGetResponsePending) {
137: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
138: 	return
```

## G104 in xaccount.go:461

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `ff5eca87f765b5e33c517c66ceb8abd42ac884011e415d57baab5337ba8317dc`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
460: func (u XAccountReauthResponseUnion) AsXAccountReauthResponseXAccountConnectionChallenge() (v XAccountReauthResponseXAccountConnectionChallenge) {
461: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
462: 	return
```

## G104 in xaccount.go:456

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `ff5eca87f765b5e33c517c66ceb8abd42ac884011e415d57baab5337ba8317dc`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
455: func (u XAccountReauthResponseUnion) AsXAccountReauthResponseXAccountConnectionAttemptPending() (v XAccountReauthResponseXAccountConnectionAttemptPending) {
456: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
457: 	return
```

## G104 in xaccount.go:451

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `ff5eca87f765b5e33c517c66ceb8abd42ac884011e415d57baab5337ba8317dc`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
450: func (u XAccountReauthResponseUnion) AsXAccountReauthResponseSanitizedXAccount() (v XAccountReauthResponseSanitizedXAccount) {
451: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
452: 	return
```

## G104 in xaccount.go:261

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `ff5eca87f765b5e33c517c66ceb8abd42ac884011e415d57baab5337ba8317dc`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
260: func (u XAccountNewResponseUnion) AsXAccountNewResponseXAccountConnectionChallenge() (v XAccountNewResponseXAccountConnectionChallenge) {
261: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
262: 	return
```

## G104 in xaccount.go:256

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `ff5eca87f765b5e33c517c66ceb8abd42ac884011e415d57baab5337ba8317dc`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
255: func (u XAccountNewResponseUnion) AsXAccountNewResponseXAccountConnectionAttemptPending() (v XAccountNewResponseXAccountConnectionAttemptPending) {
256: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
257: 	return
```

## G104 in xaccount.go:251

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `ff5eca87f765b5e33c517c66ceb8abd42ac884011e415d57baab5337ba8317dc`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
250: func (u XAccountNewResponseUnion) AsXAccountNewResponseSanitizedXAccount() (v XAccountNewResponseSanitizedXAccount) {
251: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
252: 	return
```

## G104 in supportattachment_test.go:19

Local test HTTP handler writes fixture output. Client-side response/result assertions cover delivered bytes; this call is not a production sink.

Source SHA-256: `021679a0e76e5072ec09e80829ddb94e792702aef786f737c9df98d020971c1a`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
18: 		w.WriteHeader(200)
19: 		w.Write([]byte("abc"))
20: 	}))
```

## G104 in shared/shared.go:25107

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `1ed865dfa06f424a05f33f6815ed7672e0a7d2e81a7b533d0d985d3abcc0fa42`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
25106: func (u SearchTweetSportsContextScheduledAtMsUnion) AsString() (v string) {
25107: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
25108: 	return
```

## G104 in shared/shared.go:25102

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `1ed865dfa06f424a05f33f6815ed7672e0a7d2e81a7b533d0d985d3abcc0fa42`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
25101: func (u SearchTweetSportsContextScheduledAtMsUnion) AsFloat() (v float64) {
25102: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
25103: 	return
```

## G104 in shared/shared.go:22989

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `1ed865dfa06f424a05f33f6815ed7672e0a7d2e81a7b533d0d985d3abcc0fa42`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
22988: func (u EmbeddedTweetSportsContextScheduledAtMsUnion) AsString() (v string) {
22989: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
22990: 	return
```

## G104 in shared/shared.go:22984

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `1ed865dfa06f424a05f33f6815ed7672e0a7d2e81a7b533d0d985d3abcc0fa42`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
22983: func (u EmbeddedTweetSportsContextScheduledAtMsUnion) AsFloat() (v float64) {
22984: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
22985: 	return
```

## G104 in shared/shared.go:22789

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `1ed865dfa06f424a05f33f6815ed7672e0a7d2e81a7b533d0d985d3abcc0fa42`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
22788: func (u EmbeddedTweetRetweetedTweetSportsContextScheduledAtMsUnion) AsString() (v string) {
22789: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
22790: 	return
```

## G104 in shared/shared.go:22784

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `1ed865dfa06f424a05f33f6815ed7672e0a7d2e81a7b533d0d985d3abcc0fa42`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
22783: func (u EmbeddedTweetRetweetedTweetSportsContextScheduledAtMsUnion) AsFloat() (v float64) {
22784: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
22785: 	return
```

## G104 in shared/shared.go:22585

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `1ed865dfa06f424a05f33f6815ed7672e0a7d2e81a7b533d0d985d3abcc0fa42`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
22584: func (u EmbeddedTweetRetweetedTweetRetweetedTweetSportsContextScheduledAtMsUnion) AsString() (v string) {
22585: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
22586: 	return
```

## G104 in shared/shared.go:22580

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `1ed865dfa06f424a05f33f6815ed7672e0a7d2e81a7b533d0d985d3abcc0fa42`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
22579: func (u EmbeddedTweetRetweetedTweetRetweetedTweetSportsContextScheduledAtMsUnion) AsFloat() (v float64) {
22580: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
22581: 	return
```

## G104 in shared/shared.go:22375

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `1ed865dfa06f424a05f33f6815ed7672e0a7d2e81a7b533d0d985d3abcc0fa42`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
22374: func (u EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetSportsContextScheduledAtMsUnion) AsString() (v string) {
22375: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
22376: 	return
```

## G104 in shared/shared.go:22370

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `1ed865dfa06f424a05f33f6815ed7672e0a7d2e81a7b533d0d985d3abcc0fa42`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
22369: func (u EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetSportsContextScheduledAtMsUnion) AsFloat() (v float64) {
22370: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
22371: 	return
```

## G104 in shared/shared.go:20845

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `1ed865dfa06f424a05f33f6815ed7672e0a7d2e81a7b533d0d985d3abcc0fa42`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
20844: func (u EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetSportsContextScheduledAtMsUnion) AsString() (v string) {
20845: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
20846: 	return
```

## G104 in shared/shared.go:20840

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `1ed865dfa06f424a05f33f6815ed7672e0a7d2e81a7b533d0d985d3abcc0fa42`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
20839: func (u EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetSportsContextScheduledAtMsUnion) AsFloat() (v float64) {
20840: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
20841: 	return
```

## G104 in shared/shared.go:18077

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `1ed865dfa06f424a05f33f6815ed7672e0a7d2e81a7b533d0d985d3abcc0fa42`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
18076: func (u EmbeddedTweetRetweetedTweetQuotedTweetSportsContextScheduledAtMsUnion) AsString() (v string) {
18077: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
18078: 	return
```

## G104 in shared/shared.go:18072

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `1ed865dfa06f424a05f33f6815ed7672e0a7d2e81a7b533d0d985d3abcc0fa42`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
18071: func (u EmbeddedTweetRetweetedTweetQuotedTweetSportsContextScheduledAtMsUnion) AsFloat() (v float64) {
18072: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
18073: 	return
```

## G104 in shared/shared.go:17867

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `1ed865dfa06f424a05f33f6815ed7672e0a7d2e81a7b533d0d985d3abcc0fa42`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
17866: func (u EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetSportsContextScheduledAtMsUnion) AsString() (v string) {
17867: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
17868: 	return
```

## G104 in shared/shared.go:17862

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `1ed865dfa06f424a05f33f6815ed7672e0a7d2e81a7b533d0d985d3abcc0fa42`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
17861: func (u EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetSportsContextScheduledAtMsUnion) AsFloat() (v float64) {
17862: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
17863: 	return
```

## G104 in shared/shared.go:16341

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `1ed865dfa06f424a05f33f6815ed7672e0a7d2e81a7b533d0d985d3abcc0fa42`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
16340: func (u EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetSportsContextScheduledAtMsUnion) AsString() (v string) {
16341: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
16342: 	return
```

## G104 in shared/shared.go:16336

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `1ed865dfa06f424a05f33f6815ed7672e0a7d2e81a7b533d0d985d3abcc0fa42`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
16335: func (u EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetSportsContextScheduledAtMsUnion) AsFloat() (v float64) {
16336: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
16337: 	return
```

## G104 in shared/shared.go:12399

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `1ed865dfa06f424a05f33f6815ed7672e0a7d2e81a7b533d0d985d3abcc0fa42`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
12398: func (u EmbeddedTweetQuotedTweetSportsContextScheduledAtMsUnion) AsString() (v string) {
12399: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
12400: 	return
```

## G104 in shared/shared.go:12394

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `1ed865dfa06f424a05f33f6815ed7672e0a7d2e81a7b533d0d985d3abcc0fa42`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
12393: func (u EmbeddedTweetQuotedTweetSportsContextScheduledAtMsUnion) AsFloat() (v float64) {
12394: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
12395: 	return
```

## G104 in shared/shared.go:12195

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `1ed865dfa06f424a05f33f6815ed7672e0a7d2e81a7b533d0d985d3abcc0fa42`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
12194: func (u EmbeddedTweetQuotedTweetRetweetedTweetSportsContextScheduledAtMsUnion) AsString() (v string) {
12195: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
12196: 	return
```

## G104 in shared/shared.go:12190

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `1ed865dfa06f424a05f33f6815ed7672e0a7d2e81a7b533d0d985d3abcc0fa42`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
12189: func (u EmbeddedTweetQuotedTweetRetweetedTweetSportsContextScheduledAtMsUnion) AsFloat() (v float64) {
12190: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
12191: 	return
```

## G104 in shared/shared.go:11985

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `1ed865dfa06f424a05f33f6815ed7672e0a7d2e81a7b533d0d985d3abcc0fa42`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
11984: func (u EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetSportsContextScheduledAtMsUnion) AsString() (v string) {
11985: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
11986: 	return
```

## G104 in shared/shared.go:11980

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `1ed865dfa06f424a05f33f6815ed7672e0a7d2e81a7b533d0d985d3abcc0fa42`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
11979: func (u EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetSportsContextScheduledAtMsUnion) AsFloat() (v float64) {
11980: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
11981: 	return
```

## G104 in shared/shared.go:10459

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `1ed865dfa06f424a05f33f6815ed7672e0a7d2e81a7b533d0d985d3abcc0fa42`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
10458: func (u EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetSportsContextScheduledAtMsUnion) AsString() (v string) {
10459: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
10460: 	return
```

## G104 in shared/shared.go:10454

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `1ed865dfa06f424a05f33f6815ed7672e0a7d2e81a7b533d0d985d3abcc0fa42`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
10453: func (u EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetSportsContextScheduledAtMsUnion) AsFloat() (v float64) {
10454: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
10455: 	return
```

## G104 in shared/shared.go:7705

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `1ed865dfa06f424a05f33f6815ed7672e0a7d2e81a7b533d0d985d3abcc0fa42`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
7704: func (u EmbeddedTweetQuotedTweetQuotedTweetSportsContextScheduledAtMsUnion) AsString() (v string) {
7705: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
7706: 	return
```

## G104 in shared/shared.go:7700

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `1ed865dfa06f424a05f33f6815ed7672e0a7d2e81a7b533d0d985d3abcc0fa42`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
7699: func (u EmbeddedTweetQuotedTweetQuotedTweetSportsContextScheduledAtMsUnion) AsFloat() (v float64) {
7700: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
7701: 	return
```

## G104 in shared/shared.go:7495

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `1ed865dfa06f424a05f33f6815ed7672e0a7d2e81a7b533d0d985d3abcc0fa42`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
7494: func (u EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetSportsContextScheduledAtMsUnion) AsString() (v string) {
7495: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
7496: 	return
```

## G104 in shared/shared.go:7490

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `1ed865dfa06f424a05f33f6815ed7672e0a7d2e81a7b533d0d985d3abcc0fa42`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
7489: func (u EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetSportsContextScheduledAtMsUnion) AsFloat() (v float64) {
7490: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
7491: 	return
```

## G104 in shared/shared.go:5975

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `1ed865dfa06f424a05f33f6815ed7672e0a7d2e81a7b533d0d985d3abcc0fa42`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
5974: func (u EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetSportsContextScheduledAtMsUnion) AsString() (v string) {
5975: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
5976: 	return
```

## G104 in shared/shared.go:5970

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `1ed865dfa06f424a05f33f6815ed7672e0a7d2e81a7b533d0d985d3abcc0fa42`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
5969: func (u EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetSportsContextScheduledAtMsUnion) AsFloat() (v float64) {
5970: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
5971: 	return
```

## G104 in option/requestoption_test.go:166

Literal valid method/URL or httptest-generated local URL. No untrusted input; the resulting request/URL is used by assertions.

Source SHA-256: `1f0ee1564dbcadb14eeba9f20df17a7bb91a1a3528c1a57fb9cae41eb1ac9e1f`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
165: 		}
166: 		wantURL, _ := url.Parse("https://xquik.com/api/v1/")
167: 		if cfg.DefaultBaseURL.String() != wantURL.String() {
```

## G104 in mock_server_test.go:18

Local test HTTP handler writes fixture output. Client-side response/result assertions cover delivered bytes; this call is not a production sink.

Source SHA-256: `5bf8e9afeadda478b4c80c8119e6bfb366d9eb82f25a9c39e1bd38f5347b13f3`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
17: 		writer.Header().Set("Content-Type", "application/json")
18: 		_, _ = writer.Write([]byte("{}"))
19: 	}))
```

## G104 in internal/requestconfig/requestconfig_test.go:473

Literal valid method/URL or httptest-generated local URL. No untrusted input; the resulting request/URL is used by assertions.

Source SHA-256: `6bc9fbc9ee40fcabf1db6c8e9b380a060ad5c3ca533ca06bf4045bb2ff103959`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
472:
473: 	cfg.DefaultBaseURL, _ = url.Parse("https://example.com/")
474: 	wantErr := errors.New("transport failed")
```

## G104 in internal/requestconfig/requestconfig_test.go:441

Literal valid method/URL or httptest-generated local URL. No untrusted input; the resulting request/URL is used by assertions.

Source SHA-256: `6bc9fbc9ee40fcabf1db6c8e9b380a060ad5c3ca533ca06bf4045bb2ff103959`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
440: 	cfg := testConfig(t, http.MethodPost, "/resource", reader, &dst)
441: 	cfg.BaseURL, _ = url.Parse(server.URL + "/")
442: 	cfg.MaxRetries = 1
```

## G104 in internal/requestconfig/requestconfig_test.go:431

Local test HTTP handler writes fixture output. Client-side response/result assertions cover delivered bytes; this call is not a production sink.

Source SHA-256: `6bc9fbc9ee40fcabf1db6c8e9b380a060ad5c3ca533ca06bf4045bb2ff103959`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
430: 		w.Header().Set("Content-Type", "application/json")
431: 		_, _ = io.WriteString(w, "{}")
432: 	}))
```

## G104 in internal/requestconfig/requestconfig_test.go:397

Literal valid method/URL or httptest-generated local URL. No untrusted input; the resulting request/URL is used by assertions.

Source SHA-256: `6bc9fbc9ee40fcabf1db6c8e9b380a060ad5c3ca533ca06bf4045bb2ff103959`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
396: 		cfg := testConfig(t, http.MethodGet, "/resource", nil, &dst)
397: 		cfg.BaseURL, _ = url.Parse("https://example.com/")
398: 		cfg.MaxRetries = 0
```

## G104 in internal/requestconfig/requestconfig_test.go:387

Local test HTTP handler writes fixture output. Client-side response/result assertions cover delivered bytes; this call is not a production sink.

Source SHA-256: `6bc9fbc9ee40fcabf1db6c8e9b380a060ad5c3ca533ca06bf4045bb2ff103959`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
386: 			w.WriteHeader(http.StatusBadRequest)
387: 			_, _ = io.WriteString(w, "{")
388: 		}, http.MethodGet, nil, nil)
```

## G104 in internal/requestconfig/requestconfig_test.go:373

Local test HTTP handler writes fixture output. Client-side response/result assertions cover delivered bytes; this call is not a production sink.

Source SHA-256: `6bc9fbc9ee40fcabf1db6c8e9b380a060ad5c3ca533ca06bf4045bb2ff103959`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
372: 			w.WriteHeader(http.StatusBadRequest)
373: 			_, _ = io.WriteString(w, `{"message":"bad request"}`)
374: 		}, http.MethodGet, nil, nil)
```

## G104 in internal/requestconfig/requestconfig_test.go:357

Literal valid method/URL or httptest-generated local URL. No untrusted input; the resulting request/URL is used by assertions.

Source SHA-256: `6bc9fbc9ee40fcabf1db6c8e9b380a060ad5c3ca533ca06bf4045bb2ff103959`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
356: 		cfg := testConfig(t, http.MethodGet, "/resource", nil, &responseBody)
357: 		cfg.BaseURL, _ = url.Parse(server.URL + "/")
358: 		cfg.MaxRetries = 0
```

## G104 in internal/requestconfig/requestconfig_test.go:330

Local test HTTP handler writes fixture output. Client-side response/result assertions cover delivered bytes; this call is not a production sink.

Source SHA-256: `6bc9fbc9ee40fcabf1db6c8e9b380a060ad5c3ca533ca06bf4045bb2ff103959`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
329: 				w.Header().Set("Content-Type", test.contentType)
330: 				_, _ = io.WriteString(w, test.body)
331: 			}, http.MethodGet, nil, test.dst)
```

## G104 in internal/requestconfig/requestconfig_test.go:100

Literal valid method/URL or httptest-generated local URL. No untrusted input; the resulting request/URL is used by assertions.

Source SHA-256: `6bc9fbc9ee40fcabf1db6c8e9b380a060ad5c3ca533ca06bf4045bb2ff103959`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
99: 	cfg := testConfig(t, method, "/resource", body, dst)
100: 	cfg.BaseURL, _ = url.Parse(server.URL + "/")
101: 	cfg.MaxRetries = 0
```

## G104 in internal/requestconfig/requestconfig.go:450

Media type classification intentionally ignores parameter parse errors. A valid media type remains usable when parameters fail; invalid types fall through to typed plaintext handling rather than accepting arbitrary structured data.

Source SHA-256: `568a3ad6f55cae80a238e5a9819f90666a613b25966335eff089d01092c575e8`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
449: 	contentType := res.Header.Get("content-type")
450: 	mediaType, _, _ := mime.ParseMediaType(contentType)
451: 	isJSON := strings.Contains(mediaType, "application/json") || strings.HasSuffix(mediaType, "+json")
```

## G104 in internal/requestconfig/requestconfig.go:444

Response cleanup after read or before retry. Read/request errors are handled separately; cleanup error does not replace the primary request result. Close is still invoked, and raw-response ownership is transferred explicitly elsewhere.

Source SHA-256: `568a3ad6f55cae80a238e5a9819f90666a613b25966335eff089d01092c575e8`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
443: 	contents, err := io.ReadAll(res.Body)
444: 	_ = res.Body.Close()
445: 	if err != nil {
```

## G104 in internal/requestconfig/requestconfig.go:419

Response cleanup after read or before retry. Read/request errors are handled separately; cleanup error does not replace the primary request result. Close is still invoked, and raw-response ownership is transferred explicitly elsewhere.

Source SHA-256: `568a3ad6f55cae80a238e5a9819f90666a613b25966335eff089d01092c575e8`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
418: 		contents, err := io.ReadAll(res.Body)
419: 		_ = res.Body.Close()
420: 		if err != nil {
```

## G104 in internal/requestconfig/requestconfig.go:396

Response cleanup after read or before retry. Read/request errors are handled separately; cleanup error does not replace the primary request result. Close is still invoked, and raw-response ownership is transferred explicitly elsewhere.

Source SHA-256: `568a3ad6f55cae80a238e5a9819f90666a613b25966335eff089d01092c575e8`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
395: 		if res != nil && res.Body != nil {
396: 			_ = res.Body.Close()
397: 		}
```

## G104 in internal/requestconfig/helpers_test.go:229

Literal valid method/URL or httptest-generated local URL. No untrusted input; the resulting request/URL is used by assertions.

Source SHA-256: `f372ebd46b342159af8dae468782aae0a565da7fa7cd592d122ceb4e6dd3bc21`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
228: func TestMiddlewareAndExecuteNewRequest(t *testing.T) {
229: 	req, _ := http.NewRequest(http.MethodGet, "https://example.com", nil)
230: 	called := false
```

## G104 in internal/requestconfig/helpers_test.go:168

Literal valid method/URL or httptest-generated local URL. No untrusted input; the resulting request/URL is used by assertions.

Source SHA-256: `f372ebd46b342159af8dae468782aae0a565da7fa7cd592d122ceb4e6dd3bc21`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
167: func TestSecurityOptions(t *testing.T) {
168: 	req, _ := http.NewRequest(http.MethodGet, "https://example.com", nil)
169: 	cfg := &RequestConfig{Request: req, APIKey: "key", BearerToken: "token"}
```

## G104 in internal/requestconfig/helpers_test.go:130

Literal valid method/URL or httptest-generated local URL. No untrusted input; the resulting request/URL is used by assertions.

Source SHA-256: `f372ebd46b342159af8dae468782aae0a565da7fa7cd592d122ceb4e6dd3bc21`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
129: 	cfg := testConfig(t, http.MethodPost, "/resource", bytes.NewBufferString("body"), nil)
130: 	cfg.BaseURL, _ = url.Parse(server.URL + "/")
131: 	if err := cfg.Execute(); err != nil {
```

## G104 in internal/requestconfig/helpers_test.go:20

Literal valid method/URL or httptest-generated local URL. No untrusted input; the resulting request/URL is used by assertions.

Source SHA-256: `f372ebd46b342159af8dae468782aae0a565da7fa7cd592d122ceb4e6dd3bc21`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
19: func TestRetryHelpers(t *testing.T) {
20: 	req, _ := http.NewRequest(http.MethodPost, "https://example.com", strings.NewReader("body"))
21: 	req.GetBody = nil
```

## G104 in internal/encoding/json/stream_upstream_test.go:499

Local test HTTP handler writes fixture output. Client-side response/result assertions cover delivered bytes; this call is not a production sink.

Source SHA-256: `0d17a7dc6e3c583874aa5ae50cc43774ca410d6a329dc3a03fc94dda9cca0e88`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
498: 	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
499: 		w.Write([]byte(raw))
500: 	}))
```

## G104 in internal/encoding/json/stream_upstream_test.go:381

io.Pipe reader/writer cleanup in the test; pipe Close implementations return nil. Reads and output behavior are asserted separately.

Source SHA-256: `0d17a7dc6e3c583874aa5ae50cc43774ca410d6a329dc3a03fc94dda9cca0e88`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
380: 			r.Close()
381: 			w.Close()
382: 		})
```

## G104 in internal/encoding/json/stream_upstream_test.go:380

io.Pipe reader/writer cleanup in the test; pipe Close implementations return nil. Reads and output behavior are asserted separately.

Source SHA-256: `0d17a7dc6e3c583874aa5ae50cc43774ca410d6a329dc3a03fc94dda9cca0e88`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
379: 			}
380: 			r.Close()
381: 			w.Close()
```

## G104 in internal/encoding/json/encode_upstream_test.go:1333

Panic regression: deferred recover assertion or following failure statement rejects ordinary return. An error result is not the asserted behavior.

Source SHA-256: `a2a8b7e67da52b2e0998ec1bfa1bb47652c29639632599b4332432ece3b90ba7`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
1332: 	}()
1333: 	Marshal(&marshalPanic{})
1334: 	t.Error("Marshal should have panicked")
```

## G104 in internal/encoding/json/decode_upstream_test.go:2637

Panic regression: deferred recover assertion or following failure statement rejects ordinary return. An error result is not the asserted behavior.

Source SHA-256: `c9e5bf874a2aa1daf49b0ff1253ba34d4c0f77099dc6ef733e271f9117cf565c`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
2636: 	}()
2637: 	Unmarshal([]byte("{}"), &unmarshalPanic{})
2638: 	t.Fatalf("Unmarshal should have panicked")
```

## G104 in internal/encoding/json/decode_upstream_test.go:1368

Best-effort diagnostic serialization inside a branch that already fails the test. Errors cannot turn the test green.

Source SHA-256: `c9e5bf874a2aa1daf49b0ff1253ba34d4c0f77099dc6ef733e271f9117cf565c`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
1367: 				gotJSON, _ := Marshal(got)
1368: 				wantJSON, _ := Marshal(tt.out)
1369: 				t.Fatalf("%s: Decode:\n\tgot:  %#+v\n\twant: %#+v\n\n\tgotJSON:  %s\n\twantJSON: %s", tt.Where, got, tt.out, gotJSON, wantJSON)
```

## G104 in internal/encoding/json/decode_upstream_test.go:1367

Best-effort diagnostic serialization inside a branch that already fails the test. Errors cannot turn the test green.

Source SHA-256: `c9e5bf874a2aa1daf49b0ff1253ba34d4c0f77099dc6ef733e271f9117cf565c`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
1366: 			if got := v.Elem().Interface(); !reflect.DeepEqual(got, tt.out) {
1367: 				gotJSON, _ := Marshal(got)
1368: 				wantJSON, _ := Marshal(tt.out)
```

## G104 in internal/apiquery/query_test.go:411

Unescapes the output of url.Values.Encode, which produces valid escapes.

Source SHA-256: `22ae7514cbfbaff006034add312edb0d3f6ba3201384885ea958d6eb0af34d2b`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
410: 			}
411: 			str, _ := url.QueryUnescape(values.Encode())
412: 			if str != test.enc {
```

## G104 in internal/apiquery/coverage_test.go:107-110

Panic regression: deferred recover assertion or following failure statement rejects ordinary return. An error result is not the asserted behavior.

Source SHA-256: `0ae56f6c4de93b85d28b32a32e6f1c373d885420718b080a5dfceed4778b67a0`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
106: 			}()
107: 			_, _ = MarshalWithSettings(
108: 				[]string{"tweet"},
109: 				QuerySettings{ArrayFormat: format},
110: 			)
111: 		})
```

## G104 in internal/apijson/decoder.go:421

Lenient anonymous/nullable field decoding. Null optional unmarshalling has an unconditional nil-error path. Object field decoding records metadata and preserves raw JSON; it is not an authorization validator.

Source SHA-256: `fdffb7b3c66a4b30e15a404f43391da29cf5c41df8dafd8ebd275f870f0810c4`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
420: 			if itemNode.Type == gjson.Null && dest.IsValid() && dest.Type().Implements(reflect.TypeOf((*param.Optional)(nil)).Elem()) {
421: 				_ = dest.Addr().Interface().(json.Unmarshaler).UnmarshalJSON([]byte(itemNode.Raw))
422: 				continue
```

## G104 in internal/apijson/decoder.go:354

Lenient anonymous/nullable field decoding. Null optional unmarshalling has an unconditional nil-error path. Object field decoding records metadata and preserves raw JSON; it is not an authorization validator.

Source SHA-256: `fdffb7b3c66a4b30e15a404f43391da29cf5c41df8dafd8ebd275f870f0810c4`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
353: 		for _, decoder := range anonymousDecoders {
354: 			_ = decoder.fn(node, value.FieldByIndex(decoder.idx), state)
355: 		}
```

## G104 in internal/apierror/apierror.go:63

Best-effort byte-only diagnostic API. Dump errors produce no dump; no request is sent as an operation. Credentials are redacted on copies before dump. Request/response redaction regression suite covers original immutability.

Source SHA-256: `008a3c61e9827c66198203f6763349baf3f60b4d25c1dc244607941a03e206b6`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
62: 	}
63: 	out, _ := httputil.DumpResponse(&responseCopy, body)
64: 	return out
```

## G104 in internal/apierror/apierror.go:50

Best-effort byte-only diagnostic API. Dump errors produce no dump; no request is sent as an operation. Credentials are redacted on copies before dump. Request/response redaction regression suite covers original immutability.

Source SHA-256: `008a3c61e9827c66198203f6763349baf3f60b4d25c1dc244607941a03e206b6`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
49: 	}
50: 	out, _ := httputil.DumpRequestOut(&requestCopy, body)
51: 	return out
```

## G104 in extraction_test.go:242

Local test HTTP handler writes fixture output. Client-side response/result assertions cover delivered bytes; this call is not a production sink.

Source SHA-256: `dd6d585c59521d09ebe7bc70e23219d27a86e0f625e96b3e4bbba399e0b35b9a`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
241: 		w.WriteHeader(200)
242: 		w.Write([]byte("abc"))
243: 	}))
```

## G104 in draw_test.go:74

Local test HTTP handler writes fixture output. Client-side response/result assertions cover delivered bytes; this call is not a production sink.

Source SHA-256: `650c945d3a368ef560b4f21fe634f77a0c1137be2f423dd57aa787f0e7459ada`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
73: 		w.WriteHeader(200)
74: 		w.Write([]byte("abc"))
75: 	}))
```

## G104 in compose.go:130

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `48bd826bf6a5668993019cb18122fad40d8d76bdd7dbe2918e8320d68e977903`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
129: func (u ComposeNewResponseUnion) AsComposeNewResponseComposeScoreResult() (v ComposeNewResponseComposeScoreResult) {
130: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
131: 	return
```

## G104 in compose.go:125

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `48bd826bf6a5668993019cb18122fad40d8d76bdd7dbe2918e8320d68e977903`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
124: func (u ComposeNewResponseUnion) AsComposeNewResponseComposeRefineResult() (v ComposeNewResponseComposeRefineResult) {
125: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
126: 	return
```

## G104 in compose.go:120

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `48bd826bf6a5668993019cb18122fad40d8d76bdd7dbe2918e8320d68e977903`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
119: func (u ComposeNewResponseUnion) AsComposeNewResponseComposePrepareResult() (v ComposeNewResponseComposePrepareResult) {
120: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
121: 	return
```

## G104 in client_test.go:45

This test asserts the observed User-Agent header. Endpoint-result correctness is covered separately; transport error is not the subject of this metadata assertion.

Source SHA-256: `5c479c9498e43c1cb8e0472a8aa19797eec6e414acc41fab2fcc7bd1c769261b`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
44: 	)
45: 	_, _ = client.Account.Get(context.Background())
46: 	if userAgent != fmt.Sprintf("XTwitterScraper/Go %s", internal.PackageVersion) {
```

## G104 in client_convenience_test.go:23

Local test HTTP handler writes fixture output. Client-side response/result assertions cover delivered bytes; this call is not a production sink.

Source SHA-256: `b8c9511695deb46a3938509702c511400e3602b67ebcad22cadce3de9a24dc2c`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
22: 		w.Header().Set("Content-Type", "application/json")
23: 		_, _ = io.WriteString(w, "{}")
24: 	}))
```

## G204 in internal/scripttest/release_local_test.go:145

Test-only Git helper. The tests pass literal Git arguments, and the command runs in a temporary repository that the test created. No input comes from outside the test.

Source SHA-256: `6d1b837c6711591788603f40fad3d2e89637d2c2b1164f4a5e7e0c93110a62a1`
Severity: MEDIUM. Confidence: HIGH.
Finding: Subprocess launched with variable

```go
144: 	fixture.t.Helper()
145: 	command := exec.Command("git", args...)
146: 	command.Dir, command.Env = directory, fixture.env
```

## G204 in internal/scripttest/dependency_audit_test.go:110

Test-only runner for scripts/dependency-audit. The script path is this repository's file, and root is the test's temporary folder. A fake go command and fixed report files stand in for every outside system.

Source SHA-256: `a85797b83720f65194afcb3f705f182c8a4bc46c0181fe82200bf5ee3aea7335`
Severity: MEDIUM. Confidence: HIGH.
Finding: Subprocess launched with variable

```go
109: 	write("audit-exceptions.json", encode(exceptions), 0o644)
110: 	command := exec.Command(script, root)
111: 	command.Env = append(os.Environ(),
```

## G204 in internal/scripttest/codeql_security_test.go:56

Test-only runner for scripts/codeql-security. The script path is this repository's file, and a fake CodeQL command that the test wrote stands in for the real CLI.

Source SHA-256: `4c73b90ed74b421841316dbdb176fe063f2543ba74e80bc91ac0bbca9732ee2e`
Severity: MEDIUM. Confidence: HIGH.
Finding: Subprocess launched with variable

```go
55: 	}
56: 	command := exec.Command(script)
57: 	command.Env = append(os.Environ(), "XQUIK_CODEQL_BIN="+codeql, "FAKE_REPORT="+reportPath, "TMPDIR="+root)
```

## G204 in internal/scripttest/release_local_test.go:156

Test-only runner for bin/release-local. The path points at the copy in the test's temporary repository, and the arguments are literal flags from each test case. Fake gh and curl commands stand in for GitHub and the module proxy.

Source SHA-256: `6d1b837c6711591788603f40fad3d2e89637d2c2b1164f4a5e7e0c93110a62a1`
Severity: MEDIUM. Confidence: HIGH.
Finding: Subprocess launched with a potential tainted input or cmd arguments

```go
155: 	fixture.t.Helper()
156: 	command := exec.Command(filepath.Join(fixture.repo, "bin", "release-local"), args...)
157: 	command.Dir, command.Env = fixture.t.TempDir(), fixture.env
```

## G306 in internal/scripttest/codeql_security_test.go:50

Test-only stand-in for the CodeQL CLI, written into the test's temporary folder. It must be executable, so it needs mode 0o755. Its content is a literal in the test.

Source SHA-256: `4c73b90ed74b421841316dbdb176fe063f2543ba74e80bc91ac0bbca9732ee2e`
Severity: MEDIUM. Confidence: HIGH.
Finding: Expect WriteFile permissions to be 0600 or less

```go
49: 	reportPath := filepath.Join(root, "report.sarif")
50: 	if err := os.WriteFile(codeql, []byte(fakeCodeQL), 0o755); err != nil {
51: 		t.Fatal(err)
```

## G104 in shared/shared.go:779

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `1ed865dfa06f424a05f33f6815ed7672e0a7d2e81a7b533d0d985d3abcc0fa42`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
778: func (u CommunityMembersFacepileUnion) AsCommunityMembersFacepileCommunityMemberPreview() (v CommunityMembersFacepileCommunityMemberPreview) {
779: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
780: 	return
```

## G104 in shared/shared.go:774

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `1ed865dfa06f424a05f33f6815ed7672e0a7d2e81a7b533d0d985d3abcc0fa42`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
773: func (u CommunityMembersFacepileUnion) AsUserProfile() (v UserProfile) {
774: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
775: 	return
```

## G104 in shared/shared.go:501

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `1ed865dfa06f424a05f33f6815ed7672e0a7d2e81a7b533d0d985d3abcc0fa42`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
500: func (u CommunityCreatorUnion) AsCommunityCreatorObject() (v CommunityCreatorObject) {
501: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
502: 	return
```

## G104 in shared/shared.go:496

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `1ed865dfa06f424a05f33f6815ed7672e0a7d2e81a7b533d0d985d3abcc0fa42`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
495: func (u CommunityCreatorUnion) AsUserProfile() (v UserProfile) {
496: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
497: 	return
```

## G104 in xanalysis.go:4119

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `eb148537d9d58e8f8edc27f992856702ef09d05582bf0978426778307f3a7138`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
4118: func (u XAnalysisViralScoreResponseResultAnswersUnion) AsString() (v string) {
4119: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
4120: 	return
```

## G104 in xanalysis.go:4114

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `eb148537d9d58e8f8edc27f992856702ef09d05582bf0978426778307f3a7138`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
4113: func (u XAnalysisViralScoreResponseResultAnswersUnion) AsFloat() (v float64) {
4114: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
4115: 	return
```

## G104 in xanalysis.go:3976

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `eb148537d9d58e8f8edc27f992856702ef09d05582bf0978426778307f3a7138`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
3975: func (u XAnalysisViralScoreResponseResultAnalysisAnswerValueUnion) AsString() (v string) {
3976: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
3977: 	return
```

## G104 in xanalysis.go:3971

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `eb148537d9d58e8f8edc27f992856702ef09d05582bf0978426778307f3a7138`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
3970: func (u XAnalysisViralScoreResponseResultAnalysisAnswerValueUnion) AsFloat() (v float64) {
3971: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
3972: 	return
```

## G104 in xanalysis.go:3427

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `eb148537d9d58e8f8edc27f992856702ef09d05582bf0978426778307f3a7138`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
3426: func (u XAnalysisSentimentResponseResultAnswersUnion) AsString() (v string) {
3427: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
3428: 	return
```

## G104 in xanalysis.go:3422

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `eb148537d9d58e8f8edc27f992856702ef09d05582bf0978426778307f3a7138`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
3421: func (u XAnalysisSentimentResponseResultAnswersUnion) AsFloat() (v float64) {
3422: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
3423: 	return
```

## G104 in xanalysis.go:3286

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `eb148537d9d58e8f8edc27f992856702ef09d05582bf0978426778307f3a7138`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
3285: func (u XAnalysisSentimentResponseResultAnalysisAnswerValueUnion) AsString() (v string) {
3286: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
3287: 	return
```

## G104 in xanalysis.go:3281

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `eb148537d9d58e8f8edc27f992856702ef09d05582bf0978426778307f3a7138`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
3280: func (u XAnalysisSentimentResponseResultAnalysisAnswerValueUnion) AsFloat() (v float64) {
3281: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
3282: 	return
```

## G104 in xanalysis.go:2737

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `eb148537d9d58e8f8edc27f992856702ef09d05582bf0978426778307f3a7138`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
2736: func (u XAnalysisNewsResponseResultAnswersUnion) AsString() (v string) {
2737: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
2738: 	return
```

## G104 in xanalysis.go:2732

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `eb148537d9d58e8f8edc27f992856702ef09d05582bf0978426778307f3a7138`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
2731: func (u XAnalysisNewsResponseResultAnswersUnion) AsFloat() (v float64) {
2732: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
2733: 	return
```

## G104 in xanalysis.go:2598

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `eb148537d9d58e8f8edc27f992856702ef09d05582bf0978426778307f3a7138`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
2597: func (u XAnalysisNewsResponseResultAnalysisAnswerValueUnion) AsString() (v string) {
2598: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
2599: 	return
```

## G104 in xanalysis.go:2593

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `eb148537d9d58e8f8edc27f992856702ef09d05582bf0978426778307f3a7138`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
2592: func (u XAnalysisNewsResponseResultAnalysisAnswerValueUnion) AsFloat() (v float64) {
2593: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
2594: 	return
```

## G104 in xanalysis.go:2051

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `eb148537d9d58e8f8edc27f992856702ef09d05582bf0978426778307f3a7138`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
2050: func (u XAnalysisMarketSignalsResponseResultAnswersUnion) AsString() (v string) {
2051: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
2052: 	return
```

## G104 in xanalysis.go:2046

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `eb148537d9d58e8f8edc27f992856702ef09d05582bf0978426778307f3a7138`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
2045: func (u XAnalysisMarketSignalsResponseResultAnswersUnion) AsFloat() (v float64) {
2046: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
2047: 	return
```

## G104 in xanalysis.go:1908

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `eb148537d9d58e8f8edc27f992856702ef09d05582bf0978426778307f3a7138`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
1907: func (u XAnalysisMarketSignalsResponseResultAnalysisAnswerValueUnion) AsString() (v string) {
1908: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
1909: 	return
```

## G104 in xanalysis.go:1903

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `eb148537d9d58e8f8edc27f992856702ef09d05582bf0978426778307f3a7138`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
1902: func (u XAnalysisMarketSignalsResponseResultAnalysisAnswerValueUnion) AsFloat() (v float64) {
1903: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
1904: 	return
```

## G104 in xanalysis.go:1355

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `eb148537d9d58e8f8edc27f992856702ef09d05582bf0978426778307f3a7138`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
1354: func (u XAnalysisClassifyResponseResultAnswersUnion) AsString() (v string) {
1355: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
1356: 	return
```

## G104 in xanalysis.go:1350

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `eb148537d9d58e8f8edc27f992856702ef09d05582bf0978426778307f3a7138`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
1349: func (u XAnalysisClassifyResponseResultAnswersUnion) AsFloat() (v float64) {
1350: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
1351: 	return
```

## G104 in xanalysis.go:1214

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `eb148537d9d58e8f8edc27f992856702ef09d05582bf0978426778307f3a7138`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
1213: func (u XAnalysisClassifyResponseResultAnalysisAnswerValueUnion) AsString() (v string) {
1214: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
1215: 	return
```

## G104 in xanalysis.go:1209

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `eb148537d9d58e8f8edc27f992856702ef09d05582bf0978426778307f3a7138`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
1208: func (u XAnalysisClassifyResponseResultAnalysisAnswerValueUnion) AsFloat() (v float64) {
1209: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
1210: 	return
```

## G104 in xanalysis.go:665

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `eb148537d9d58e8f8edc27f992856702ef09d05582bf0978426778307f3a7138`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
664: func (u XAnalysisBrandResponseResultAnswersUnion) AsString() (v string) {
665: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
666: 	return
```

## G104 in xanalysis.go:660

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `eb148537d9d58e8f8edc27f992856702ef09d05582bf0978426778307f3a7138`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
659: func (u XAnalysisBrandResponseResultAnswersUnion) AsFloat() (v float64) {
660: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
661: 	return
```

## G104 in xanalysis.go:526

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `eb148537d9d58e8f8edc27f992856702ef09d05582bf0978426778307f3a7138`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
525: func (u XAnalysisBrandResponseResultAnalysisAnswerValueUnion) AsString() (v string) {
526: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
527: 	return
```

## G104 in xanalysis.go:521

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `eb148537d9d58e8f8edc27f992856702ef09d05582bf0978426778307f3a7138`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
520: func (u XAnalysisBrandResponseResultAnalysisAnswerValueUnion) AsFloat() (v float64) {
521: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
522: 	return
```

## G304 in internal/apicheck/apicheck.go:118

The API coverage and field checks read the spec file that the developer running bin/check-local names in XQUIK_OPENAPI_SPEC. No remote caller supplies the path, and the checks only parse the file as JSON.

Source SHA-256: `20e2b4585c3a6d6c4a6c0dcb8a2d8659930f1185c4530071d696d73d294cc0e4`
Severity: MEDIUM. Confidence: HIGH.
Finding: Potential file inclusion via variable

```go
117: 	if !strings.HasPrefix(source, "https://") && !strings.HasPrefix(source, "http://") {
118: 		return os.ReadFile(source)
119: 	}
```

## G204 in internal/scripttest/temp_cleanup_test.go:264

Test-only runner for the check scripts and bin/release-local. The paths point at this repository's scripts or the copy in the test's temporary repository, and the arguments are literal flags from each test case. Fake tools stand in for go, CodeQL, uv, gh and curl.

Source SHA-256: `e56c4cae1eda5a921d1f2ae82671c4168844fe2d5741193c50433b902ff0db92`
Severity: MEDIUM. Confidence: HIGH.
Finding: Subprocess launched with a potential tainted input or cmd arguments

```go
263: 	}
264: 	cmd := exec.Command(filepath.Join(fixture.repo, "bin", "release-local"), args...)
265: 	cmd.Dir = t.TempDir()
```

## G204 in internal/scripttest/temp_cleanup_test.go:123

Test-only runner for the check scripts and bin/release-local. The paths point at this repository's scripts or the copy in the test's temporary repository, and the arguments are literal flags from each test case. Fake tools stand in for go, CodeQL, uv, gh and curl.

Source SHA-256: `e56c4cae1eda5a921d1f2ae82671c4168844fe2d5741193c50433b902ff0db92`
Severity: MEDIUM. Confidence: HIGH.
Finding: Subprocess launched with a potential tainted input or cmd arguments

```go
122: 	}
123: 	cmd := exec.Command(command[0], command[1:]...)
124: 	cmd.Env = append(os.Environ(), env...)
```

## G306 in internal/scripttest/temp_cleanup_test.go:254

Test-only stand-in for bin/check-local, written into the test's temporary repository. It must be executable, so it needs mode 0o755. Each test supplies its content.

Source SHA-256: `e56c4cae1eda5a921d1f2ae82671c4168844fe2d5741193c50433b902ff0db92`
Severity: MEDIUM. Confidence: HIGH.
Finding: Expect WriteFile permissions to be 0600 or less

```go
253: 	checkLocal = "#!/bin/sh\nmktemp \"${GOTMPDIR:-$TMPDIR}/go-build.XXXXXX\" >/dev/null || exit 1\n" + checkLocal
254: 	if err := os.WriteFile(filepath.Join(fixture.repo, "bin", "check-local"), []byte(checkLocal), 0o755); err != nil {
255: 		t.Fatal(err)
```

## G306 in internal/scripttest/temp_cleanup_test.go:105

Test-only stand-ins for the check tools, written into the test's temporary folder. They must be executable, so they need mode 0o755. Their content is a literal in the test.

Source SHA-256: `e56c4cae1eda5a921d1f2ae82671c4168844fe2d5741193c50433b902ff0db92`
Severity: MEDIUM. Confidence: HIGH.
Finding: Expect WriteFile permissions to be 0600 or less

```go
104: 	for _, tool := range test.tools {
105: 		if err := os.WriteFile(filepath.Join(bin, tool), []byte(fakeTool), 0o755); err != nil {
106: 			t.Fatal(err)
```

## G302 in internal/scripttest/temp_cleanup_test.go:130

Test-only cleanup restores owner access to directories in the fake module cache. The cache is inside the test's temporary root. Mode 0o700 lets the test remove its read-only fixture and grants no access to other users.

Source SHA-256: `e56c4cae1eda5a921d1f2ae82671c4168844fe2d5741193c50433b902ff0db92`
Severity: MEDIUM. Confidence: HIGH.
Finding: Expect file permissions to be 0600 or less

```go
129: 		for _, directory := range []string{moduleCache, filepath.Join(moduleCache, "cache"), filepath.Join(moduleCache, "cache", "download")} {
130: 			if err := os.Chmod(directory, 0o700); err != nil && !errors.Is(err, os.ErrNotExist) {
131: 				t.Errorf("%v\n%s\n%s", err, cmd.Stdout, cmd.Stderr)
```

## G304 in internal/scripttest/temp_cleanup_test.go:327

Test-only read of retained coverage artifacts. The output directory comes from t.TempDir and the file names are the literal coverage.out and coverage.xml cases. No external input supplies either path.

Source SHA-256: `e56c4cae1eda5a921d1f2ae82671c4168844fe2d5741193c50433b902ff0db92`
Severity: MEDIUM. Confidence: HIGH.
Finding: Potential file inclusion via variable

```go
326: 				for _, name := range []string{"coverage.out", "coverage.xml"} {
327: 					data, err := os.ReadFile(filepath.Join(output, name))
328: 					if err != nil || string(data) != "mode: set\n" {
```

## G104 in activity.go:145

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `33d36922d62659e2005541cb579dbb4d2b2ea9564e29907e17e83dd633c97254`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
144: func (u ActivityLogRowResultsUnion) AsInt() (v int64) {
145: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
146: 	return
```

## G104 in activity.go:150

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `33d36922d62659e2005541cb579dbb4d2b2ea9564e29907e17e83dd633c97254`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
149: func (u ActivityLogRowResultsUnion) AsString() (v string) {
150: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
151: 	return
```

## G104 in activity.go:180

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `33d36922d62659e2005541cb579dbb4d2b2ea9564e29907e17e83dd633c97254`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
179: func (u ActivityLogRowCreditsUnion) AsInt() (v int64) {
180: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
181: 	return
```

## G104 in activity.go:185

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `33d36922d62659e2005541cb579dbb4d2b2ea9564e29907e17e83dd633c97254`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
184: func (u ActivityLogRowCreditsUnion) AsString() (v string) {
185: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
186: 	return
```
