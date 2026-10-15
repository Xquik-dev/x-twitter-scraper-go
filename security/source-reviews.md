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

## G404 in internal/requestconfig/requestconfig.go:310

Bounded retry jitter, not a credential, key, authentication value, or security token. The arithmetic overflow was repaired and verified on ARM64 and AMD64. Cryptographic unpredictability is not a requirement for jitter.

Source SHA-256: `38eb76cde02d28e4120d542130b8d00d950a1d187ef4516a2d57943112e2ad03`
Severity: HIGH. Confidence: MEDIUM.
Finding: Use of weak random number generator (math/rand or math/rand/v2 instead of crypto/rand)

```go
309: 	delay := 500 * time.Millisecond << min(max(retryCount, 0), 4)
310: 	return delay - time.Duration(rand.Int63n(int64(delay/4)))
311: }
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

## G104 in xuser.go:1038

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `9a4ca384472a6e1e2fa9037616a6c5b2a4ec476ce6b2ebb40f6a4c5aa3634203`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
1037: func (u XUserGetVerifiedFollowersResponseUnion) AsXUserGetVerifiedFollowersResponseUserListCoverageResponse() (v XUserGetVerifiedFollowersResponseUserListCoverageResponse) {
1038: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
1039: 	return
```

## G104 in xuser.go:1033

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `9a4ca384472a6e1e2fa9037616a6c5b2a4ec476ce6b2ebb40f6a4c5aa3634203`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
1032: func (u XUserGetVerifiedFollowersResponseUnion) AsPaginatedUsers() (v shared.PaginatedUsers) {
1033: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
1034: 	return
```

## G104 in xuser.go:891

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `9a4ca384472a6e1e2fa9037616a6c5b2a4ec476ce6b2ebb40f6a4c5aa3634203`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
890: func (u XUserGetFollowingResponseUnion) AsXUserGetFollowingResponseUserListCoverageResponse() (v XUserGetFollowingResponseUserListCoverageResponse) {
891: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
892: 	return
```

## G104 in xuser.go:886

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `9a4ca384472a6e1e2fa9037616a6c5b2a4ec476ce6b2ebb40f6a4c5aa3634203`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
885: func (u XUserGetFollowingResponseUnion) AsPaginatedUsers() (v shared.PaginatedUsers) {
886: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
887: 	return
```

## G104 in xuser.go:746

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `9a4ca384472a6e1e2fa9037616a6c5b2a4ec476ce6b2ebb40f6a4c5aa3634203`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
745: func (u XUserGetFollowersResponseUnion) AsXUserGetFollowersResponseUserListCoverageResponse() (v XUserGetFollowersResponseUserListCoverageResponse) {
746: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
747: 	return
```

## G104 in xuser.go:741

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `9a4ca384472a6e1e2fa9037616a6c5b2a4ec476ce6b2ebb40f6a4c5aa3634203`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
740: func (u XUserGetFollowersResponseUnion) AsPaginatedUsers() (v shared.PaginatedUsers) {
741: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
742: 	return
```

## G104 in xtweet.go:2426

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `0422001a58e2d2310ce06083430c9eaee6bcfcaa5f095e59c84897e91a9b1a78`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
2425: func (u XTweetSearchResponseUnion) AsXTweetSearchResponseTweetSearchCoverageResponse() (v XTweetSearchResponseTweetSearchCoverageResponse) {
2426: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
2427: 	return
```

## G104 in xtweet.go:2421

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `0422001a58e2d2310ce06083430c9eaee6bcfcaa5f095e59c84897e91a9b1a78`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
2420: func (u XTweetSearchResponseUnion) AsPaginatedTweets() (v shared.PaginatedTweets) {
2421: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
2422: 	return
```

## G104 in xtweet.go:1516

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `0422001a58e2d2310ce06083430c9eaee6bcfcaa5f095e59c84897e91a9b1a78`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
1515: func (u TweetDetailSportsContextScheduledAtMsUnion) AsString() (v string) {
1516: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
1517: 	return
```

## G104 in xtweet.go:1511

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `0422001a58e2d2310ce06083430c9eaee6bcfcaa5f095e59c84897e91a9b1a78`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
1510: func (u TweetDetailSportsContextScheduledAtMsUnion) AsFloat() (v float64) {
1511: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
1512: 	return
```

## G104 in xlist.go:283

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `d5e3115c1df6f96f01e29a18bfa91c6f5213784e76f8dd133fa9e9f7ca4c7f38`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
282: func (u XListGetMembersResponseUnion) AsXListGetMembersResponseUserListCoverageResponse() (v XListGetMembersResponseUserListCoverageResponse) {
283: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
284: 	return
```

## G104 in xlist.go:278

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `d5e3115c1df6f96f01e29a18bfa91c6f5213784e76f8dd133fa9e9f7ca4c7f38`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
277: func (u XListGetMembersResponseUnion) AsPaginatedUsers() (v shared.PaginatedUsers) {
278: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
279: 	return
```

## G104 in xlist.go:138

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `d5e3115c1df6f96f01e29a18bfa91c6f5213784e76f8dd133fa9e9f7ca4c7f38`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
137: func (u XListGetFollowersResponseUnion) AsXListGetFollowersResponseUserListCoverageResponse() (v XListGetFollowersResponseUserListCoverageResponse) {
138: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
139: 	return
```

## G104 in xlist.go:133

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `d5e3115c1df6f96f01e29a18bfa91c6f5213784e76f8dd133fa9e9f7ca4c7f38`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
132: func (u XListGetFollowersResponseUnion) AsPaginatedUsers() (v shared.PaginatedUsers) {
133: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
134: 	return
```

## G104 in xaccountconnectionattempt.go:149

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `ae5dbdc3272542e5b9e00beb62980fa94068c4ccc76a091825f935407f29c25c`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
148: func (u XAccountConnectionAttemptGetResponseUnion) AsRequiresEmailCode() (v XAccountConnectionAttemptGetResponseRequiresEmailCode) {
149: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
150: 	return
```

## G104 in xaccountconnectionattempt.go:144

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `ae5dbdc3272542e5b9e00beb62980fa94068c4ccc76a091825f935407f29c25c`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
143: func (u XAccountConnectionAttemptGetResponseUnion) AsFailed() (v XAccountConnectionAttemptGetResponseFailed) {
144: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
145: 	return
```

## G104 in xaccountconnectionattempt.go:139

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `ae5dbdc3272542e5b9e00beb62980fa94068c4ccc76a091825f935407f29c25c`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
138: func (u XAccountConnectionAttemptGetResponseUnion) AsSuccess() (v XAccountConnectionAttemptGetResponseSuccess) {
139: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
140: 	return
```

## G104 in xaccountconnectionattempt.go:134

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `ae5dbdc3272542e5b9e00beb62980fa94068c4ccc76a091825f935407f29c25c`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
133: func (u XAccountConnectionAttemptGetResponseUnion) AsPending() (v XAccountConnectionAttemptGetResponsePending) {
134: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
135: 	return
```

## G104 in xaccount.go:257

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `b7c2f0e1180f0b68f4a3e964f416bbfcf263bcc31005c738b13cb8b704f28f60`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
256: func (u XAccountNewResponseUnion) AsXAccountNewResponseXAccountConnectionChallenge() (v XAccountNewResponseXAccountConnectionChallenge) {
257: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
258: 	return
```

## G104 in xaccount.go:252

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `b7c2f0e1180f0b68f4a3e964f416bbfcf263bcc31005c738b13cb8b704f28f60`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
251: func (u XAccountNewResponseUnion) AsXAccountNewResponseXAccountConnectionAttemptPending() (v XAccountNewResponseXAccountConnectionAttemptPending) {
252: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
253: 	return
```

## G104 in xaccount.go:247

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `b7c2f0e1180f0b68f4a3e964f416bbfcf263bcc31005c738b13cb8b704f28f60`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
246: func (u XAccountNewResponseUnion) AsXAccountNewResponseSanitizedXAccount() (v XAccountNewResponseSanitizedXAccount) {
247: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
248: 	return
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

## G104 in shared/shared.go:22540

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `ac2d21aa28cc31d0b79d835cf77a2c9cd809b980fad2ab1212a296c9a3ad2b55`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
22539: func (u SearchTweetSportsContextScheduledAtMsUnion) AsString() (v string) {
22540: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
22541: 	return
```

## G104 in shared/shared.go:22535

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `ac2d21aa28cc31d0b79d835cf77a2c9cd809b980fad2ab1212a296c9a3ad2b55`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
22534: func (u SearchTweetSportsContextScheduledAtMsUnion) AsFloat() (v float64) {
22535: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
22536: 	return
```

## G104 in shared/shared.go:21163

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `ac2d21aa28cc31d0b79d835cf77a2c9cd809b980fad2ab1212a296c9a3ad2b55`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
21162: func (u EmbeddedTweetSportsContextScheduledAtMsUnion) AsString() (v string) {
21163: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
21164: 	return
```

## G104 in shared/shared.go:21158

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `ac2d21aa28cc31d0b79d835cf77a2c9cd809b980fad2ab1212a296c9a3ad2b55`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
21157: func (u EmbeddedTweetSportsContextScheduledAtMsUnion) AsFloat() (v float64) {
21158: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
21159: 	return
```

## G104 in shared/shared.go:20975

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `ac2d21aa28cc31d0b79d835cf77a2c9cd809b980fad2ab1212a296c9a3ad2b55`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
20974: func (u EmbeddedTweetRetweetedTweetSportsContextScheduledAtMsUnion) AsString() (v string) {
20975: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
20976: 	return
```

## G104 in shared/shared.go:20970

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `ac2d21aa28cc31d0b79d835cf77a2c9cd809b980fad2ab1212a296c9a3ad2b55`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
20969: func (u EmbeddedTweetRetweetedTweetSportsContextScheduledAtMsUnion) AsFloat() (v float64) {
20970: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
20971: 	return
```

## G104 in shared/shared.go:20783

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `ac2d21aa28cc31d0b79d835cf77a2c9cd809b980fad2ab1212a296c9a3ad2b55`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
20782: func (u EmbeddedTweetRetweetedTweetRetweetedTweetSportsContextScheduledAtMsUnion) AsString() (v string) {
20783: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
20784: 	return
```

## G104 in shared/shared.go:20778

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `ac2d21aa28cc31d0b79d835cf77a2c9cd809b980fad2ab1212a296c9a3ad2b55`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
20777: func (u EmbeddedTweetRetweetedTweetRetweetedTweetSportsContextScheduledAtMsUnion) AsFloat() (v float64) {
20778: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
20779: 	return
```

## G104 in shared/shared.go:20585

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `ac2d21aa28cc31d0b79d835cf77a2c9cd809b980fad2ab1212a296c9a3ad2b55`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
20584: func (u EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetSportsContextScheduledAtMsUnion) AsString() (v string) {
20585: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
20586: 	return
```

## G104 in shared/shared.go:20580

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `ac2d21aa28cc31d0b79d835cf77a2c9cd809b980fad2ab1212a296c9a3ad2b55`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
20579: func (u EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetSportsContextScheduledAtMsUnion) AsFloat() (v float64) {
20580: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
20581: 	return
```

## G104 in shared/shared.go:19166

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `ac2d21aa28cc31d0b79d835cf77a2c9cd809b980fad2ab1212a296c9a3ad2b55`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
19165: func (u EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetSportsContextScheduledAtMsUnion) AsString() (v string) {
19166: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
19167: 	return
```

## G104 in shared/shared.go:19161

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `ac2d21aa28cc31d0b79d835cf77a2c9cd809b980fad2ab1212a296c9a3ad2b55`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
19160: func (u EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetSportsContextScheduledAtMsUnion) AsFloat() (v float64) {
19161: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
19162: 	return
```

## G104 in shared/shared.go:16598

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `ac2d21aa28cc31d0b79d835cf77a2c9cd809b980fad2ab1212a296c9a3ad2b55`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
16597: func (u EmbeddedTweetRetweetedTweetQuotedTweetSportsContextScheduledAtMsUnion) AsString() (v string) {
16598: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
16599: 	return
```

## G104 in shared/shared.go:16593

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `ac2d21aa28cc31d0b79d835cf77a2c9cd809b980fad2ab1212a296c9a3ad2b55`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
16592: func (u EmbeddedTweetRetweetedTweetQuotedTweetSportsContextScheduledAtMsUnion) AsFloat() (v float64) {
16593: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
16594: 	return
```

## G104 in shared/shared.go:16400

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `ac2d21aa28cc31d0b79d835cf77a2c9cd809b980fad2ab1212a296c9a3ad2b55`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
16399: func (u EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetSportsContextScheduledAtMsUnion) AsString() (v string) {
16400: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
16401: 	return
```

## G104 in shared/shared.go:16395

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `ac2d21aa28cc31d0b79d835cf77a2c9cd809b980fad2ab1212a296c9a3ad2b55`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
16394: func (u EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetSportsContextScheduledAtMsUnion) AsFloat() (v float64) {
16395: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
16396: 	return
```

## G104 in shared/shared.go:14985

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `ac2d21aa28cc31d0b79d835cf77a2c9cd809b980fad2ab1212a296c9a3ad2b55`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
14984: func (u EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetSportsContextScheduledAtMsUnion) AsString() (v string) {
14985: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
14986: 	return
```

## G104 in shared/shared.go:14980

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `ac2d21aa28cc31d0b79d835cf77a2c9cd809b980fad2ab1212a296c9a3ad2b55`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
14979: func (u EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetSportsContextScheduledAtMsUnion) AsFloat() (v float64) {
14980: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
14981: 	return
```

## G104 in shared/shared.go:11332

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `ac2d21aa28cc31d0b79d835cf77a2c9cd809b980fad2ab1212a296c9a3ad2b55`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
11331: func (u EmbeddedTweetQuotedTweetSportsContextScheduledAtMsUnion) AsString() (v string) {
11332: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
11333: 	return
```

## G104 in shared/shared.go:11327

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `ac2d21aa28cc31d0b79d835cf77a2c9cd809b980fad2ab1212a296c9a3ad2b55`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
11326: func (u EmbeddedTweetQuotedTweetSportsContextScheduledAtMsUnion) AsFloat() (v float64) {
11327: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
11328: 	return
```

## G104 in shared/shared.go:11140

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `ac2d21aa28cc31d0b79d835cf77a2c9cd809b980fad2ab1212a296c9a3ad2b55`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
11139: func (u EmbeddedTweetQuotedTweetRetweetedTweetSportsContextScheduledAtMsUnion) AsString() (v string) {
11140: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
11141: 	return
```

## G104 in shared/shared.go:11135

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `ac2d21aa28cc31d0b79d835cf77a2c9cd809b980fad2ab1212a296c9a3ad2b55`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
11134: func (u EmbeddedTweetQuotedTweetRetweetedTweetSportsContextScheduledAtMsUnion) AsFloat() (v float64) {
11135: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
11136: 	return
```

## G104 in shared/shared.go:10942

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `ac2d21aa28cc31d0b79d835cf77a2c9cd809b980fad2ab1212a296c9a3ad2b55`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
10941: func (u EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetSportsContextScheduledAtMsUnion) AsString() (v string) {
10942: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
10943: 	return
```

## G104 in shared/shared.go:10937

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `ac2d21aa28cc31d0b79d835cf77a2c9cd809b980fad2ab1212a296c9a3ad2b55`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
10936: func (u EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetSportsContextScheduledAtMsUnion) AsFloat() (v float64) {
10937: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
10938: 	return
```

## G104 in shared/shared.go:9527

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `ac2d21aa28cc31d0b79d835cf77a2c9cd809b980fad2ab1212a296c9a3ad2b55`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
9526: func (u EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetSportsContextScheduledAtMsUnion) AsString() (v string) {
9527: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
9528: 	return
```

## G104 in shared/shared.go:9522

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `ac2d21aa28cc31d0b79d835cf77a2c9cd809b980fad2ab1212a296c9a3ad2b55`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
9521: func (u EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetSportsContextScheduledAtMsUnion) AsFloat() (v float64) {
9522: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
9523: 	return
```

## G104 in shared/shared.go:6973

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `ac2d21aa28cc31d0b79d835cf77a2c9cd809b980fad2ab1212a296c9a3ad2b55`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
6972: func (u EmbeddedTweetQuotedTweetQuotedTweetSportsContextScheduledAtMsUnion) AsString() (v string) {
6973: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
6974: 	return
```

## G104 in shared/shared.go:6968

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `ac2d21aa28cc31d0b79d835cf77a2c9cd809b980fad2ab1212a296c9a3ad2b55`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
6967: func (u EmbeddedTweetQuotedTweetQuotedTweetSportsContextScheduledAtMsUnion) AsFloat() (v float64) {
6968: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
6969: 	return
```

## G104 in shared/shared.go:6775

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `ac2d21aa28cc31d0b79d835cf77a2c9cd809b980fad2ab1212a296c9a3ad2b55`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
6774: func (u EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetSportsContextScheduledAtMsUnion) AsString() (v string) {
6775: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
6776: 	return
```

## G104 in shared/shared.go:6770

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `ac2d21aa28cc31d0b79d835cf77a2c9cd809b980fad2ab1212a296c9a3ad2b55`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
6769: func (u EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetSportsContextScheduledAtMsUnion) AsFloat() (v float64) {
6770: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
6771: 	return
```

## G104 in shared/shared.go:5366

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `ac2d21aa28cc31d0b79d835cf77a2c9cd809b980fad2ab1212a296c9a3ad2b55`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
5365: func (u EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetSportsContextScheduledAtMsUnion) AsString() (v string) {
5366: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
5367: 	return
```

## G104 in shared/shared.go:5361

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `ac2d21aa28cc31d0b79d835cf77a2c9cd809b980fad2ab1212a296c9a3ad2b55`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
5360: func (u EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetSportsContextScheduledAtMsUnion) AsFloat() (v float64) {
5361: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
5362: 	return
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

## G104 in internal/requestconfig/requestconfig_test.go:427

Literal valid method/URL or httptest-generated local URL. No untrusted input; the resulting request/URL is used by assertions.

Source SHA-256: `5076abbd4b0b83fd2113c3e0d3793910b137d315a462f8425bccb244fc55ea24`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
426:
427: 	cfg.DefaultBaseURL, _ = url.Parse("https://example.com/")
428: 	wantErr := errors.New("transport failed")
```

## G104 in internal/requestconfig/requestconfig_test.go:395

Literal valid method/URL or httptest-generated local URL. No untrusted input; the resulting request/URL is used by assertions.

Source SHA-256: `5076abbd4b0b83fd2113c3e0d3793910b137d315a462f8425bccb244fc55ea24`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
394: 	cfg := testConfig(t, http.MethodPost, "/resource", reader, &dst)
395: 	cfg.BaseURL, _ = url.Parse(server.URL + "/")
396: 	cfg.MaxRetries = 1
```

## G104 in internal/requestconfig/requestconfig_test.go:385

Local test HTTP handler writes fixture output. Client-side response/result assertions cover delivered bytes; this call is not a production sink.

Source SHA-256: `5076abbd4b0b83fd2113c3e0d3793910b137d315a462f8425bccb244fc55ea24`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
384: 		w.Header().Set("Content-Type", "application/json")
385: 		_, _ = io.WriteString(w, "{}")
386: 	}))
```

## G104 in internal/requestconfig/requestconfig_test.go:351

Literal valid method/URL or httptest-generated local URL. No untrusted input; the resulting request/URL is used by assertions.

Source SHA-256: `5076abbd4b0b83fd2113c3e0d3793910b137d315a462f8425bccb244fc55ea24`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
350: 		cfg := testConfig(t, http.MethodGet, "/resource", nil, &dst)
351: 		cfg.BaseURL, _ = url.Parse("https://example.com/")
352: 		cfg.MaxRetries = 0
```

## G104 in internal/requestconfig/requestconfig_test.go:341

Local test HTTP handler writes fixture output. Client-side response/result assertions cover delivered bytes; this call is not a production sink.

Source SHA-256: `5076abbd4b0b83fd2113c3e0d3793910b137d315a462f8425bccb244fc55ea24`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
340: 			w.WriteHeader(http.StatusBadRequest)
341: 			_, _ = io.WriteString(w, "{")
342: 		}, http.MethodGet, nil, nil)
```

## G104 in internal/requestconfig/requestconfig_test.go:327

Local test HTTP handler writes fixture output. Client-side response/result assertions cover delivered bytes; this call is not a production sink.

Source SHA-256: `5076abbd4b0b83fd2113c3e0d3793910b137d315a462f8425bccb244fc55ea24`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
326: 			w.WriteHeader(http.StatusBadRequest)
327: 			_, _ = io.WriteString(w, `{"message":"bad request"}`)
328: 		}, http.MethodGet, nil, nil)
```

## G104 in internal/requestconfig/requestconfig_test.go:311

Literal valid method/URL or httptest-generated local URL. No untrusted input; the resulting request/URL is used by assertions.

Source SHA-256: `5076abbd4b0b83fd2113c3e0d3793910b137d315a462f8425bccb244fc55ea24`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
310: 		cfg := testConfig(t, http.MethodGet, "/resource", nil, &responseBody)
311: 		cfg.BaseURL, _ = url.Parse(server.URL + "/")
312: 		cfg.MaxRetries = 0
```

## G104 in internal/requestconfig/requestconfig_test.go:284

Local test HTTP handler writes fixture output. Client-side response/result assertions cover delivered bytes; this call is not a production sink.

Source SHA-256: `5076abbd4b0b83fd2113c3e0d3793910b137d315a462f8425bccb244fc55ea24`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
283: 				w.Header().Set("Content-Type", test.contentType)
284: 				_, _ = io.WriteString(w, test.body)
285: 			}, http.MethodGet, nil, test.dst)
```

## G104 in internal/requestconfig/requestconfig_test.go:86

Literal valid method/URL or httptest-generated local URL. No untrusted input; the resulting request/URL is used by assertions.

Source SHA-256: `5076abbd4b0b83fd2113c3e0d3793910b137d315a462f8425bccb244fc55ea24`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
85: 	cfg := testConfig(t, method, "/resource", body, dst)
86: 	cfg.BaseURL, _ = url.Parse(server.URL + "/")
87: 	cfg.MaxRetries = 0
```

## G104 in internal/requestconfig/requestconfig.go:441

Media type classification intentionally ignores parameter parse errors. A valid media type remains usable when parameters fail; invalid types fall through to typed plaintext handling rather than accepting arbitrary structured data.

Source SHA-256: `38eb76cde02d28e4120d542130b8d00d950a1d187ef4516a2d57943112e2ad03`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
440: 	contentType := res.Header.Get("content-type")
441: 	mediaType, _, _ := mime.ParseMediaType(contentType)
442: 	isJSON := strings.Contains(mediaType, "application/json") || strings.HasSuffix(mediaType, "+json")
```

## G104 in internal/requestconfig/requestconfig.go:435

Response cleanup after read or before retry. Read/request errors are handled separately; cleanup error does not replace the primary request result. Close is still invoked, and raw-response ownership is transferred explicitly elsewhere.

Source SHA-256: `38eb76cde02d28e4120d542130b8d00d950a1d187ef4516a2d57943112e2ad03`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
434: 	contents, err := io.ReadAll(res.Body)
435: 	_ = res.Body.Close()
436: 	if err != nil {
```

## G104 in internal/requestconfig/requestconfig.go:410

Response cleanup after read or before retry. Read/request errors are handled separately; cleanup error does not replace the primary request result. Close is still invoked, and raw-response ownership is transferred explicitly elsewhere.

Source SHA-256: `38eb76cde02d28e4120d542130b8d00d950a1d187ef4516a2d57943112e2ad03`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
409: 		contents, err := io.ReadAll(res.Body)
410: 		_ = res.Body.Close()
411: 		if err != nil {
```

## G104 in internal/requestconfig/requestconfig.go:387

Response cleanup after read or before retry. Read/request errors are handled separately; cleanup error does not replace the primary request result. Close is still invoked, and raw-response ownership is transferred explicitly elsewhere.

Source SHA-256: `38eb76cde02d28e4120d542130b8d00d950a1d187ef4516a2d57943112e2ad03`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
386: 		if res != nil && res.Body != nil {
387: 			_ = res.Body.Close()
388: 		}
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

## G104 in client_test.go:44

This test asserts the observed User-Agent header. Endpoint-result correctness is covered separately; transport error is not the subject of this metadata assertion.

Source SHA-256: `09166219f05bbcc95489be75493a313b179a335bdb7a478fc1668b80937edd24`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
43: 	)
44: 	_, _ = client.Account.Get(context.Background())
45: 	if userAgent != fmt.Sprintf("XTwitterScraper/Go %s", internal.PackageVersion) {
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

## G204 in internal/scripttest/release_local_test.go:134

Test-only Git helper. The tests pass literal Git arguments, and the command runs in a temporary repository that the test created. No input comes from outside the test.

Source SHA-256: `fb3810d1a0546aa9bd864d5be01a5cb7951cb1b2b5c2fa7931a53816bad859bd`
Severity: MEDIUM. Confidence: HIGH.
Finding: Subprocess launched with variable

```go
133: 	fixture.t.Helper()
134: 	command := exec.Command("git", args...)
135: 	command.Dir, command.Env = directory, fixture.env
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

Source SHA-256: `f92d985ed9da1a10e51f03597feb569a22e373f92987f8c88f09f93c5f5ded08`
Severity: MEDIUM. Confidence: HIGH.
Finding: Subprocess launched with variable

```go
55: 	}
56: 	command := exec.Command(script)
57: 	command.Env = append(os.Environ(), "XQUIK_CODEQL_BIN="+codeql, "FAKE_REPORT="+reportPath, "TMPDIR="+root)
```

## G204 in internal/scripttest/release_local_test.go:145

Test-only runner for bin/release-local. The path points at the copy in the test's temporary repository, and the arguments are literal flags from each test case. Fake gh and curl commands stand in for GitHub and the module proxy.

Source SHA-256: `fb3810d1a0546aa9bd864d5be01a5cb7951cb1b2b5c2fa7931a53816bad859bd`
Severity: MEDIUM. Confidence: HIGH.
Finding: Subprocess launched with a potential tainted input or cmd arguments

```go
144: 	fixture.t.Helper()
145: 	command := exec.Command(filepath.Join(fixture.repo, "bin", "release-local"), args...)
146: 	command.Dir, command.Env = fixture.t.TempDir(), fixture.env
```

## G306 in internal/scripttest/codeql_security_test.go:50

Test-only stand-in for the CodeQL CLI, written into the test's temporary folder. It must be executable, so it needs mode 0o755. Its content is a literal in the test.

Source SHA-256: `f92d985ed9da1a10e51f03597feb569a22e373f92987f8c88f09f93c5f5ded08`
Severity: MEDIUM. Confidence: HIGH.
Finding: Expect WriteFile permissions to be 0600 or less

```go
49: 	reportPath := filepath.Join(root, "report.sarif")
50: 	if err := os.WriteFile(codeql, []byte(fakeCodeQL), 0o755); err != nil {
51: 		t.Fatal(err)
```

## G704 in cmd/api-coverage/main.go:134

The API coverage check fetches the spec URL that the developer running bin/check-local chooses with XQUIK_OPENAPI_SPEC, or https://xquik.com/openapi.json by default. No remote caller supplies it. The check runs only on a maintainer machine, reads the JSON answer and sends no credentials.

Source SHA-256: `fe30cc708a4e0af62916addd347d5fb061e24802fe925a56bb86b649f57c3664`
Severity: HIGH. Confidence: HIGH.
Finding: SSRF via taint analysis

```go
133: 	}
134: 	response, err := client.Do(request)
135: 	if err != nil {
```

## G704 in cmd/api-coverage/main.go:130

The API coverage check fetches the spec URL that the developer running bin/check-local chooses with XQUIK_OPENAPI_SPEC, or https://xquik.com/openapi.json by default. No remote caller supplies it. The check runs only on a maintainer machine, reads the JSON answer and sends no credentials.

Source SHA-256: `fe30cc708a4e0af62916addd347d5fb061e24802fe925a56bb86b649f57c3664`
Severity: HIGH. Confidence: HIGH.
Finding: SSRF via taint analysis

```go
129: 	defer cancel()
130: 	request, err := http.NewRequestWithContext(ctx, http.MethodGet, source, nil)
131: 	if err != nil {
```

## G703 in cmd/api-coverage/main.go:126

The API coverage check reads the spec file that the developer running bin/check-local names in XQUIK_OPENAPI_SPEC. No remote caller supplies the path, and the check only parses the file as JSON.

Source SHA-256: `fe30cc708a4e0af62916addd347d5fb061e24802fe925a56bb86b649f57c3664`
Severity: HIGH. Confidence: HIGH.
Finding: Path traversal via taint analysis

```go
125: 	if !strings.HasPrefix(source, "https://") && !strings.HasPrefix(source, "http://") {
126: 		return os.ReadFile(source)
127: 	}
```

## G304 in cmd/api-coverage/main.go:126

The API coverage check reads the spec file that the developer running bin/check-local names in XQUIK_OPENAPI_SPEC. No remote caller supplies the path, and the check only parses the file as JSON.

Source SHA-256: `fe30cc708a4e0af62916addd347d5fb061e24802fe925a56bb86b649f57c3664`
Severity: MEDIUM. Confidence: HIGH.
Finding: Potential file inclusion via variable

```go
125: 	if !strings.HasPrefix(source, "https://") && !strings.HasPrefix(source, "http://") {
126: 		return os.ReadFile(source)
127: 	}
```

## G104 in shared/shared.go:557

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `ac2d21aa28cc31d0b79d835cf77a2c9cd809b980fad2ab1212a296c9a3ad2b55`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
556: func (u CommunityMembersFacepileUnion) AsCommunityMembersFacepileCommunityMemberPreview() (v CommunityMembersFacepileCommunityMemberPreview) {
557: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
558: 	return
```

## G104 in shared/shared.go:552

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `ac2d21aa28cc31d0b79d835cf77a2c9cd809b980fad2ab1212a296c9a3ad2b55`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
551: func (u CommunityMembersFacepileUnion) AsUserProfile() (v UserProfile) {
552: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
553: 	return
```

## G104 in shared/shared.go:347

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `ac2d21aa28cc31d0b79d835cf77a2c9cd809b980fad2ab1212a296c9a3ad2b55`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
346: func (u CommunityCreatorUnion) AsCommunityCreatorObject() (v CommunityCreatorObject) {
347: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
348: 	return
```

## G104 in shared/shared.go:342

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `ac2d21aa28cc31d0b79d835cf77a2c9cd809b980fad2ab1212a296c9a3ad2b55`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
341: func (u CommunityCreatorUnion) AsUserProfile() (v UserProfile) {
342: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
343: 	return
```
