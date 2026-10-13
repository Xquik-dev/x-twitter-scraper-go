# Approved source findings

The user approved these 123 exact findings on 2026-09-12.
On 2026-10-13, `scripts/move_source_reviews.go` moved 68 of them to new line numbers & file hashes.
It moves a finding only when its code & surrounding lines are unchanged, and it changed no reason.
3 records name `cmd/reproducible-build/upload_test.go`, which no longer exists, so nothing can match them.
On 2026-10-13, 5 findings in `internal/scripttest` were reviewed and added. Each runs only in tests, with input the test controls.
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

## G404 in internal/requestconfig/requestconfig.go:345

Bounded retry jitter, not a credential, key, authentication value, or security token. The arithmetic overflow was repaired and verified on ARM64 and AMD64. Cryptographic unpredictability is not a requirement for jitter.

Source SHA-256: `5f88430c11b9e4ca9290ca8df712bcca7d76081b9d8b0e67921abef8d396fecc`
Severity: HIGH. Confidence: MEDIUM.
Finding: Use of weak random number generator (math/rand or math/rand/v2 instead of crypto/rand)

```go
344: 	delay := 500 * time.Millisecond << min(max(retryCount, 0), 4)
345: 	return delay - time.Duration(rand.Int63n(int64(delay/4)))
346: }
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

## G101 in guestwallet.go:386

Public generated credential-storage notice, not a credential value.

Source SHA-256: `d819b178d71877adad1c3e76696186b2f233cea6cc6b52763e12e13696eaa011`
Severity: HIGH. Confidence: LOW.
Finding: Potential hardcoded credentials

```go
385: const (
386: 	GuestWalletTopupResponseCredentialNoticeStoreAPIKeyAndTheIdempotencyKeySecurelyBeforeSharingCheckoutURLNoEmailRecoveryIsAvailable GuestWalletTopupResponseCredentialNotice = "Store api_key and the Idempotency-Key securely before sharing checkout_url. No email recovery is available."
387: )
```

## G703 in cmd/reproducible-build/main.go:72

Explicit local CLI output paths. No remote path source or claimed confinement boundary.

Source SHA-256: `d2af26f3cfd794448a4e9106077842b41b753238d2ab34e6c3306e8752ce478b`
Severity: HIGH. Confidence: HIGH.
Finding: Path traversal via taint analysis

```go
71: 		}
72: 		if err := os.WriteFile(args[0], first, 0o644); err != nil {
73: 			return err
```

## G703 in cmd/reproducible-build/main.go:69

Explicit local CLI output paths. No remote path source or claimed confinement boundary.

Source SHA-256: `d2af26f3cfd794448a4e9106077842b41b753238d2ab34e6c3306e8752ce478b`
Severity: HIGH. Confidence: HIGH.
Finding: Path traversal via taint analysis

```go
68: 	if len(args) == 1 {
69: 		if err := os.MkdirAll(filepath.Dir(args[0]), 0o755); err != nil {
70: 			return err
```

## G703 in cmd/coverage-gain/main.go:58

Explicit local CLI output paths. No remote path source or claimed confinement boundary.

Source SHA-256: `0eae09b979dfabc057b0b1322d7ea53fd2dda62436ae3870e242a02a12b106ce`
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

## G204 in cmd/reproducible-build/main.go:24

Fixed executable and argument array. Build helpers consume local Git arguments; upload test invokes the fixed repository script. No shell interpolation of remote data.

Source SHA-256: `d2af26f3cfd794448a4e9106077842b41b753238d2ab34e6c3306e8752ce478b`
Severity: MEDIUM. Confidence: HIGH.
Finding: Subprocess launched with variable

```go
23: 	git := func(args ...string) ([]byte, error) {
24: 		command := exec.Command("git", args...)
25: 		command.Env = append(os.Environ(), "GIT_INDEX_FILE="+filepath.Join(directory, "index"), "TZ=UTC")
```

## G204 in cmd/coverage-gain/main.go:22

Fixed executable and argument array. Build helpers consume local Git arguments; upload test invokes the fixed repository script. No shell interpolation of remote data.

Source SHA-256: `0eae09b979dfabc057b0b1322d7ea53fd2dda62436ae3870e242a02a12b106ce`
Severity: MEDIUM. Confidence: HIGH.
Finding: Subprocess launched with variable

```go
21: func git(args ...string) (string, error) {
22: 	command := exec.Command("git", args...)
23: 	command.Env = append(os.Environ(), "LC_ALL=C")
```

## G204 in cmd/coverage-gain/main_test.go:26

Fixed executable and argument array. Build helpers consume local Git arguments; upload test invokes the fixed repository script. No shell interpolation of remote data.

Source SHA-256: `ae4a756fbcd388fe4adb788c43e45e7d74a8c2961938e9e52a9b51156720395d`
Severity: MEDIUM. Confidence: HIGH.
Finding: Subprocess launched with a potential tainted input or cmd arguments

```go
25: 		t.Helper()
26: 		command := exec.Command("git", append([]string{"-c", "user.name=Coverage test", "-c", "user.email=coverage@example.test", "-c", "commit.gpgsign=false"}, args...)...)
27: 		if data, err := command.CombinedOutput(); err != nil {
```

## G304 in cmd/coverage-gain/main_test.go:84

Test reads its own summary artifact within a temporary directory.

Source SHA-256: `ae4a756fbcd388fe4adb788c43e45e7d74a8c2961938e9e52a9b51156720395d`
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

## G301 in cmd/reproducible-build/main.go:69

Caller-selected directory for source archive. Archive intended for publication; no secret file permission boundary.

Source SHA-256: `d2af26f3cfd794448a4e9106077842b41b753238d2ab34e6c3306e8752ce478b`
Severity: MEDIUM. Confidence: HIGH.
Finding: Expect directory permissions to be 0750 or less

```go
68: 	if len(args) == 1 {
69: 		if err := os.MkdirAll(filepath.Dir(args[0]), 0o755); err != nil {
70: 			return err
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

## G306 in cmd/reproducible-build/main.go:72

Source archive or public fixture contents. Temporary fixture directories use testing.TempDir isolation.

Source SHA-256: `d2af26f3cfd794448a4e9106077842b41b753238d2ab34e6c3306e8752ce478b`
Severity: MEDIUM. Confidence: HIGH.
Finding: Expect WriteFile permissions to be 0600 or less

```go
71: 		}
72: 		if err := os.WriteFile(args[0], first, 0o644); err != nil {
73: 			return err
```

## G103 in internal/apijson/decoder.go:661

Only call assigns node.Raw into the generated JSON.raw metadata field. The pointer comes from the addressable current Go reflect.Value, without pointer arithmetic or stored unsafe pointers. reflect.Set performs typed assignment. RawJSON model tests cover metadata preservation.

Source SHA-256: `bdd153f5d5964310c4f647c173c7dad095620a8e81dfcfe8cfa19dca8a846437`
Severity: LOW. Confidence: HIGH.
Finding: Use of unsafe calls should be audited

```go
660: func setUnexportedField(field reflect.Value, value any) {
661: 	reflect.NewAt(field.Type(), unsafe.Pointer(field.UnsafeAddr())).Elem().Set(reflect.ValueOf(value))
662: }
```

## G104 in xuser.go:866

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `30c51e40c26b7bf4ee2b8e789cfb33545869164bd7db9fd6dd5cd023dbefc683`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
865: func (u XUserGetVerifiedFollowersResponseUnion) AsXUserGetVerifiedFollowersResponseUserListCoverageResponse() (v XUserGetVerifiedFollowersResponseUserListCoverageResponse) {
866: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
867: 	return
```

## G104 in xuser.go:861

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `30c51e40c26b7bf4ee2b8e789cfb33545869164bd7db9fd6dd5cd023dbefc683`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
860: func (u XUserGetVerifiedFollowersResponseUnion) AsPaginatedUsers() (v shared.PaginatedUsers) {
861: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
862: 	return
```

## G104 in xuser.go:719

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `30c51e40c26b7bf4ee2b8e789cfb33545869164bd7db9fd6dd5cd023dbefc683`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
718: func (u XUserGetFollowingResponseUnion) AsXUserGetFollowingResponseUserListCoverageResponse() (v XUserGetFollowingResponseUserListCoverageResponse) {
719: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
720: 	return
```

## G104 in xuser.go:714

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `30c51e40c26b7bf4ee2b8e789cfb33545869164bd7db9fd6dd5cd023dbefc683`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
713: func (u XUserGetFollowingResponseUnion) AsPaginatedUsers() (v shared.PaginatedUsers) {
714: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
715: 	return
```

## G104 in xuser.go:574

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `30c51e40c26b7bf4ee2b8e789cfb33545869164bd7db9fd6dd5cd023dbefc683`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
573: func (u XUserGetFollowersResponseUnion) AsXUserGetFollowersResponseUserListCoverageResponse() (v XUserGetFollowersResponseUserListCoverageResponse) {
574: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
575: 	return
```

## G104 in xuser.go:569

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `30c51e40c26b7bf4ee2b8e789cfb33545869164bd7db9fd6dd5cd023dbefc683`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
568: func (u XUserGetFollowersResponseUnion) AsPaginatedUsers() (v shared.PaginatedUsers) {
569: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
570: 	return
```

## G104 in xtweet.go:2310

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `27cf093f601bb7eeb96460e2145b22ab6252de0c2a0cf147d2486de65e774774`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
2309: func (u XTweetSearchResponseUnion) AsXTweetSearchResponseTweetSearchCoverageResponse() (v XTweetSearchResponseTweetSearchCoverageResponse) {
2310: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
2311: 	return
```

## G104 in xtweet.go:2305

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `27cf093f601bb7eeb96460e2145b22ab6252de0c2a0cf147d2486de65e774774`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
2304: func (u XTweetSearchResponseUnion) AsPaginatedTweets() (v shared.PaginatedTweets) {
2305: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
2306: 	return
```

## G104 in xtweet.go:1400

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `27cf093f601bb7eeb96460e2145b22ab6252de0c2a0cf147d2486de65e774774`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
1399: func (u TweetDetailSportsContextScheduledAtMsUnion) AsString() (v string) {
1400: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
1401: 	return
```

## G104 in xtweet.go:1395

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `27cf093f601bb7eeb96460e2145b22ab6252de0c2a0cf147d2486de65e774774`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
1394: func (u TweetDetailSportsContextScheduledAtMsUnion) AsFloat() (v float64) {
1395: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
1396: 	return
```

## G104 in xlist.go:260

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `06609dba681a4d3ea04b529e8dffbf09625257d1c64616cf8cde12aef0eabc98`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
259: func (u XListGetMembersResponseUnion) AsXListGetMembersResponseUserListCoverageResponse() (v XListGetMembersResponseUserListCoverageResponse) {
260: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
261: 	return
```

## G104 in xlist.go:255

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `06609dba681a4d3ea04b529e8dffbf09625257d1c64616cf8cde12aef0eabc98`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
254: func (u XListGetMembersResponseUnion) AsPaginatedUsers() (v shared.PaginatedUsers) {
255: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
256: 	return
```

## G104 in xlist.go:115

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `06609dba681a4d3ea04b529e8dffbf09625257d1c64616cf8cde12aef0eabc98`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
114: func (u XListGetFollowersResponseUnion) AsXListGetFollowersResponseUserListCoverageResponse() (v XListGetFollowersResponseUserListCoverageResponse) {
115: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
116: 	return
```

## G104 in xlist.go:110

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `06609dba681a4d3ea04b529e8dffbf09625257d1c64616cf8cde12aef0eabc98`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
109: func (u XListGetFollowersResponseUnion) AsPaginatedUsers() (v shared.PaginatedUsers) {
110: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
111: 	return
```

## G104 in xaccountconnectionattempt.go:153

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `9feab0f3e510674b1f6a5c80d6a94ff976eac2e93732d1b053a66c06db1e0e5f`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
152: func (u XAccountConnectionAttemptGetResponseUnion) AsRequiresEmailCode() (v XAccountConnectionAttemptGetResponseRequiresEmailCode) {
153: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
154: 	return
```

## G104 in xaccountconnectionattempt.go:148

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `9feab0f3e510674b1f6a5c80d6a94ff976eac2e93732d1b053a66c06db1e0e5f`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
147: func (u XAccountConnectionAttemptGetResponseUnion) AsFailed() (v XAccountConnectionAttemptGetResponseFailed) {
148: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
149: 	return
```

## G104 in xaccountconnectionattempt.go:143

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `9feab0f3e510674b1f6a5c80d6a94ff976eac2e93732d1b053a66c06db1e0e5f`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
142: func (u XAccountConnectionAttemptGetResponseUnion) AsSuccess() (v XAccountConnectionAttemptGetResponseSuccess) {
143: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
144: 	return
```

## G104 in xaccountconnectionattempt.go:138

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `9feab0f3e510674b1f6a5c80d6a94ff976eac2e93732d1b053a66c06db1e0e5f`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
137: func (u XAccountConnectionAttemptGetResponseUnion) AsPending() (v XAccountConnectionAttemptGetResponsePending) {
138: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
139: 	return
```

## G104 in xaccount.go:258

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `15bf5077d0bd81ab9a87b47435d362ea8dfb0fc0b0fc88acd3ad06635a89e9e7`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
257: func (u XAccountNewResponseUnion) AsXAccountNewResponseXAccountConnectionChallenge() (v XAccountNewResponseXAccountConnectionChallenge) {
258: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
259: 	return
```

## G104 in xaccount.go:253

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `15bf5077d0bd81ab9a87b47435d362ea8dfb0fc0b0fc88acd3ad06635a89e9e7`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
252: func (u XAccountNewResponseUnion) AsXAccountNewResponseXAccountConnectionAttemptPending() (v XAccountNewResponseXAccountConnectionAttemptPending) {
253: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
254: 	return
```

## G104 in xaccount.go:248

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `15bf5077d0bd81ab9a87b47435d362ea8dfb0fc0b0fc88acd3ad06635a89e9e7`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
247: func (u XAccountNewResponseUnion) AsXAccountNewResponseSanitizedXAccount() (v XAccountNewResponseSanitizedXAccount) {
248: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
249: 	return
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

## G104 in shared/shared.go:21857

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `e24ce7712ba70d75ca0d11a2cd1ef5387d11c39e5606e59551c5bc13f86a1e2e`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
21856: func (u SearchTweetSportsContextScheduledAtMsUnion) AsString() (v string) {
21857: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
21858: 	return
```

## G104 in shared/shared.go:21852

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `e24ce7712ba70d75ca0d11a2cd1ef5387d11c39e5606e59551c5bc13f86a1e2e`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
21851: func (u SearchTweetSportsContextScheduledAtMsUnion) AsFloat() (v float64) {
21852: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
21853: 	return
```

## G104 in shared/shared.go:20480

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `e24ce7712ba70d75ca0d11a2cd1ef5387d11c39e5606e59551c5bc13f86a1e2e`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
20479: func (u EmbeddedTweetSportsContextScheduledAtMsUnion) AsString() (v string) {
20480: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
20481: 	return
```

## G104 in shared/shared.go:20475

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `e24ce7712ba70d75ca0d11a2cd1ef5387d11c39e5606e59551c5bc13f86a1e2e`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
20474: func (u EmbeddedTweetSportsContextScheduledAtMsUnion) AsFloat() (v float64) {
20475: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
20476: 	return
```

## G104 in shared/shared.go:20292

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `e24ce7712ba70d75ca0d11a2cd1ef5387d11c39e5606e59551c5bc13f86a1e2e`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
20291: func (u EmbeddedTweetRetweetedTweetSportsContextScheduledAtMsUnion) AsString() (v string) {
20292: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
20293: 	return
```

## G104 in shared/shared.go:20287

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `e24ce7712ba70d75ca0d11a2cd1ef5387d11c39e5606e59551c5bc13f86a1e2e`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
20286: func (u EmbeddedTweetRetweetedTweetSportsContextScheduledAtMsUnion) AsFloat() (v float64) {
20287: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
20288: 	return
```

## G104 in shared/shared.go:20100

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `e24ce7712ba70d75ca0d11a2cd1ef5387d11c39e5606e59551c5bc13f86a1e2e`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
20099: func (u EmbeddedTweetRetweetedTweetRetweetedTweetSportsContextScheduledAtMsUnion) AsString() (v string) {
20100: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
20101: 	return
```

## G104 in shared/shared.go:20095

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `e24ce7712ba70d75ca0d11a2cd1ef5387d11c39e5606e59551c5bc13f86a1e2e`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
20094: func (u EmbeddedTweetRetweetedTweetRetweetedTweetSportsContextScheduledAtMsUnion) AsFloat() (v float64) {
20095: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
20096: 	return
```

## G104 in shared/shared.go:19902

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `e24ce7712ba70d75ca0d11a2cd1ef5387d11c39e5606e59551c5bc13f86a1e2e`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
19901: func (u EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetSportsContextScheduledAtMsUnion) AsString() (v string) {
19902: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
19903: 	return
```

## G104 in shared/shared.go:19897

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `e24ce7712ba70d75ca0d11a2cd1ef5387d11c39e5606e59551c5bc13f86a1e2e`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
19896: func (u EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetSportsContextScheduledAtMsUnion) AsFloat() (v float64) {
19897: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
19898: 	return
```

## G104 in shared/shared.go:18483

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `e24ce7712ba70d75ca0d11a2cd1ef5387d11c39e5606e59551c5bc13f86a1e2e`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
18482: func (u EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetSportsContextScheduledAtMsUnion) AsString() (v string) {
18483: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
18484: 	return
```

## G104 in shared/shared.go:18478

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `e24ce7712ba70d75ca0d11a2cd1ef5387d11c39e5606e59551c5bc13f86a1e2e`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
18477: func (u EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetSportsContextScheduledAtMsUnion) AsFloat() (v float64) {
18478: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
18479: 	return
```

## G104 in shared/shared.go:15915

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `e24ce7712ba70d75ca0d11a2cd1ef5387d11c39e5606e59551c5bc13f86a1e2e`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
15914: func (u EmbeddedTweetRetweetedTweetQuotedTweetSportsContextScheduledAtMsUnion) AsString() (v string) {
15915: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
15916: 	return
```

## G104 in shared/shared.go:15910

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `e24ce7712ba70d75ca0d11a2cd1ef5387d11c39e5606e59551c5bc13f86a1e2e`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
15909: func (u EmbeddedTweetRetweetedTweetQuotedTweetSportsContextScheduledAtMsUnion) AsFloat() (v float64) {
15910: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
15911: 	return
```

## G104 in shared/shared.go:15717

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `e24ce7712ba70d75ca0d11a2cd1ef5387d11c39e5606e59551c5bc13f86a1e2e`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
15716: func (u EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetSportsContextScheduledAtMsUnion) AsString() (v string) {
15717: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
15718: 	return
```

## G104 in shared/shared.go:15712

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `e24ce7712ba70d75ca0d11a2cd1ef5387d11c39e5606e59551c5bc13f86a1e2e`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
15711: func (u EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetSportsContextScheduledAtMsUnion) AsFloat() (v float64) {
15712: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
15713: 	return
```

## G104 in shared/shared.go:14302

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `e24ce7712ba70d75ca0d11a2cd1ef5387d11c39e5606e59551c5bc13f86a1e2e`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
14301: func (u EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetSportsContextScheduledAtMsUnion) AsString() (v string) {
14302: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
14303: 	return
```

## G104 in shared/shared.go:14297

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `e24ce7712ba70d75ca0d11a2cd1ef5387d11c39e5606e59551c5bc13f86a1e2e`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
14296: func (u EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetSportsContextScheduledAtMsUnion) AsFloat() (v float64) {
14297: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
14298: 	return
```

## G104 in shared/shared.go:10649

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `e24ce7712ba70d75ca0d11a2cd1ef5387d11c39e5606e59551c5bc13f86a1e2e`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
10648: func (u EmbeddedTweetQuotedTweetSportsContextScheduledAtMsUnion) AsString() (v string) {
10649: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
10650: 	return
```

## G104 in shared/shared.go:10644

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `e24ce7712ba70d75ca0d11a2cd1ef5387d11c39e5606e59551c5bc13f86a1e2e`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
10643: func (u EmbeddedTweetQuotedTweetSportsContextScheduledAtMsUnion) AsFloat() (v float64) {
10644: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
10645: 	return
```

## G104 in shared/shared.go:10457

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `e24ce7712ba70d75ca0d11a2cd1ef5387d11c39e5606e59551c5bc13f86a1e2e`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
10456: func (u EmbeddedTweetQuotedTweetRetweetedTweetSportsContextScheduledAtMsUnion) AsString() (v string) {
10457: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
10458: 	return
```

## G104 in shared/shared.go:10452

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `e24ce7712ba70d75ca0d11a2cd1ef5387d11c39e5606e59551c5bc13f86a1e2e`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
10451: func (u EmbeddedTweetQuotedTweetRetweetedTweetSportsContextScheduledAtMsUnion) AsFloat() (v float64) {
10452: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
10453: 	return
```

## G104 in shared/shared.go:10259

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `e24ce7712ba70d75ca0d11a2cd1ef5387d11c39e5606e59551c5bc13f86a1e2e`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
10258: func (u EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetSportsContextScheduledAtMsUnion) AsString() (v string) {
10259: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
10260: 	return
```

## G104 in shared/shared.go:10254

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `e24ce7712ba70d75ca0d11a2cd1ef5387d11c39e5606e59551c5bc13f86a1e2e`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
10253: func (u EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetSportsContextScheduledAtMsUnion) AsFloat() (v float64) {
10254: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
10255: 	return
```

## G104 in shared/shared.go:8844

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `e24ce7712ba70d75ca0d11a2cd1ef5387d11c39e5606e59551c5bc13f86a1e2e`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
8843: func (u EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetSportsContextScheduledAtMsUnion) AsString() (v string) {
8844: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
8845: 	return
```

## G104 in shared/shared.go:8839

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `e24ce7712ba70d75ca0d11a2cd1ef5387d11c39e5606e59551c5bc13f86a1e2e`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
8838: func (u EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetSportsContextScheduledAtMsUnion) AsFloat() (v float64) {
8839: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
8840: 	return
```

## G104 in shared/shared.go:6290

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `e24ce7712ba70d75ca0d11a2cd1ef5387d11c39e5606e59551c5bc13f86a1e2e`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
6289: func (u EmbeddedTweetQuotedTweetQuotedTweetSportsContextScheduledAtMsUnion) AsString() (v string) {
6290: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
6291: 	return
```

## G104 in shared/shared.go:6285

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `e24ce7712ba70d75ca0d11a2cd1ef5387d11c39e5606e59551c5bc13f86a1e2e`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
6284: func (u EmbeddedTweetQuotedTweetQuotedTweetSportsContextScheduledAtMsUnion) AsFloat() (v float64) {
6285: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
6286: 	return
```

## G104 in shared/shared.go:6092

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `e24ce7712ba70d75ca0d11a2cd1ef5387d11c39e5606e59551c5bc13f86a1e2e`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
6091: func (u EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetSportsContextScheduledAtMsUnion) AsString() (v string) {
6092: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
6093: 	return
```

## G104 in shared/shared.go:6087

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `e24ce7712ba70d75ca0d11a2cd1ef5387d11c39e5606e59551c5bc13f86a1e2e`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
6086: func (u EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetSportsContextScheduledAtMsUnion) AsFloat() (v float64) {
6087: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
6088: 	return
```

## G104 in shared/shared.go:4683

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `e24ce7712ba70d75ca0d11a2cd1ef5387d11c39e5606e59551c5bc13f86a1e2e`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
4682: func (u EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetSportsContextScheduledAtMsUnion) AsString() (v string) {
4683: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
4684: 	return
```

## G104 in shared/shared.go:4678

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `e24ce7712ba70d75ca0d11a2cd1ef5387d11c39e5606e59551c5bc13f86a1e2e`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
4677: func (u EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetSportsContextScheduledAtMsUnion) AsFloat() (v float64) {
4678: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
4679: 	return
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

## G104 in internal/requestconfig/requestconfig.go:493

Media type classification intentionally ignores parameter parse errors. A valid media type remains usable when parameters fail; invalid types fall through to typed plaintext handling rather than accepting arbitrary structured data.

Source SHA-256: `5f88430c11b9e4ca9290ca8df712bcca7d76081b9d8b0e67921abef8d396fecc`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
492: 	contentType := res.Header.Get("content-type")
493: 	mediaType, _, _ := mime.ParseMediaType(contentType)
494: 	isJSON := strings.Contains(mediaType, "application/json") || strings.HasSuffix(mediaType, "+json")
```

## G104 in internal/requestconfig/requestconfig.go:486

Response cleanup after read or before retry. Read/request errors are handled separately; cleanup error does not replace the primary request result. Close is still invoked, and raw-response ownership is transferred explicitly elsewhere.

Source SHA-256: `5f88430c11b9e4ca9290ca8df712bcca7d76081b9d8b0e67921abef8d396fecc`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
485: 	contents, err := io.ReadAll(res.Body)
486: 	_ = res.Body.Close()
487: 	if err != nil {
```

## G104 in internal/requestconfig/requestconfig.go:455

Response cleanup after read or before retry. Read/request errors are handled separately; cleanup error does not replace the primary request result. Close is still invoked, and raw-response ownership is transferred explicitly elsewhere.

Source SHA-256: `5f88430c11b9e4ca9290ca8df712bcca7d76081b9d8b0e67921abef8d396fecc`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
454: 		contents, err := io.ReadAll(res.Body)
455: 		_ = res.Body.Close()
456: 		if err != nil {
```

## G104 in internal/requestconfig/requestconfig.go:427

Response cleanup after read or before retry. Read/request errors are handled separately; cleanup error does not replace the primary request result. Close is still invoked, and raw-response ownership is transferred explicitly elsewhere.

Source SHA-256: `5f88430c11b9e4ca9290ca8df712bcca7d76081b9d8b0e67921abef8d396fecc`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
426: 		if res != nil && res.Body != nil {
427: 			_ = res.Body.Close()
428: 		}
```

## G104 in internal/requestconfig/helpers_test.go:229

Literal valid method/URL or httptest-generated local URL. No untrusted input; the resulting request/URL is used by assertions.

Source SHA-256: `b4e56cfbb6f81a9e2008a023c83b676cb9858d9eefbd75aecb4c571b7aed5c9a`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
228: func TestMiddlewareAndExecuteNewRequest(t *testing.T) {
229: 	req, _ := http.NewRequest(http.MethodGet, "https://example.com", nil)
230: 	called := false
```

## G104 in internal/requestconfig/helpers_test.go:168

Literal valid method/URL or httptest-generated local URL. No untrusted input; the resulting request/URL is used by assertions.

Source SHA-256: `b4e56cfbb6f81a9e2008a023c83b676cb9858d9eefbd75aecb4c571b7aed5c9a`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
167: func TestSecurityOptions(t *testing.T) {
168: 	req, _ := http.NewRequest(http.MethodGet, "https://example.com", nil)
169: 	cfg := &RequestConfig{Request: req, APIKey: "key", BearerToken: "token"}
```

## G104 in internal/requestconfig/helpers_test.go:130

Literal valid method/URL or httptest-generated local URL. No untrusted input; the resulting request/URL is used by assertions.

Source SHA-256: `b4e56cfbb6f81a9e2008a023c83b676cb9858d9eefbd75aecb4c571b7aed5c9a`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
129: 	cfg := testConfig(t, http.MethodPost, "/resource", bytes.NewBufferString("body"), nil)
130: 	cfg.BaseURL, _ = url.Parse(server.URL + "/")
131: 	if err := cfg.Execute(); err != nil {
```

## G104 in internal/requestconfig/helpers_test.go:20

Literal valid method/URL or httptest-generated local URL. No untrusted input; the resulting request/URL is used by assertions.

Source SHA-256: `b4e56cfbb6f81a9e2008a023c83b676cb9858d9eefbd75aecb4c571b7aed5c9a`
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

## G104 in internal/apijson/decoder.go:469

Lenient anonymous/nullable field decoding. Null optional unmarshalling has an unconditional nil-error path. Object field decoding records metadata and preserves raw JSON; it is not an authorization validator.

Source SHA-256: `bdd153f5d5964310c4f647c173c7dad095620a8e81dfcfe8cfa19dca8a846437`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
468: 			if itemNode.Type == gjson.Null && dest.IsValid() && dest.Type().Implements(reflect.TypeOf((*param.Optional)(nil)).Elem()) {
469: 				_ = dest.Addr().Interface().(json.Unmarshaler).UnmarshalJSON([]byte(itemNode.Raw))
470: 				continue
```

## G104 in internal/apijson/decoder.go:400

Lenient anonymous/nullable field decoding. Null optional unmarshalling has an unconditional nil-error path. Object field decoding records metadata and preserves raw JSON; it is not an authorization validator.

Source SHA-256: `bdd153f5d5964310c4f647c173c7dad095620a8e81dfcfe8cfa19dca8a846437`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
399: 			// ignore errors
400: 			_ = decoder.fn(node, value.FieldByIndex(decoder.idx), state)
401: 		}
```

## G104 in internal/apierror/apierror.go:69

Best-effort byte-only diagnostic API. Dump errors produce no dump; no request is sent as an operation. Credentials are redacted on copies before dump. Request/response redaction regression suite covers original immutability.

Source SHA-256: `95f12584423660b571c8b91bfea6e648b958cf16f3924b25043c5353a6716434`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
68: 	}
69: 	out, _ := httputil.DumpResponse(&responseCopy, body)
70: 	return out
```

## G104 in internal/apierror/apierror.go:56

Best-effort byte-only diagnostic API. Dump errors produce no dump; no request is sent as an operation. Credentials are redacted on copies before dump. Request/response redaction regression suite covers original immutability.

Source SHA-256: `95f12584423660b571c8b91bfea6e648b958cf16f3924b25043c5353a6716434`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
55: 	}
56: 	out, _ := httputil.DumpRequestOut(&requestCopy, body)
57: 	return out
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

## G104 in compose.go:131

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `84be029f7cbe0da4e15472c3c1bff6b62f925a713402b2f6ad3f93dc3722f193`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
130: func (u ComposeNewResponseUnion) AsComposeNewResponseComposeScoreResult() (v ComposeNewResponseComposeScoreResult) {
131: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
132: 	return
```

## G104 in compose.go:126

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `84be029f7cbe0da4e15472c3c1bff6b62f925a713402b2f6ad3f93dc3722f193`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
125: func (u ComposeNewResponseUnion) AsComposeNewResponseComposeRefineResult() (v ComposeNewResponseComposeRefineResult) {
126: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
127: 	return
```

## G104 in compose.go:121

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `84be029f7cbe0da4e15472c3c1bff6b62f925a713402b2f6ad3f93dc3722f193`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
120: func (u ComposeNewResponseUnion) AsComposeNewResponseComposePrepareResult() (v ComposeNewResponseComposePrepareResult) {
121: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
122: 	return
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

## G204 in internal/scripttest/release_local_test.go:141

Test-only Git helper. The tests pass literal Git arguments, and the command runs in a temporary repository that the test created. No input comes from outside the test.

Source SHA-256: `f5a2e37d58b467e1f4629280271f180fb60c9aa11df697f65ac06c06ae6808d6`
Severity: MEDIUM. Confidence: HIGH.
Finding: Subprocess launched with variable

```go
140: 	fixture.t.Helper()
141: 	command := exec.Command("git", args...)
142: 	command.Dir, command.Env = directory, fixture.env
```

## G204 in internal/scripttest/dependency_audit_test.go:116

Test-only runner for scripts/dependency-audit. The script path is this repository's file, and root is the test's temporary folder. A fake go command and fixed report files stand in for every outside system.

Source SHA-256: `cfd0b9e20afadea6250ca048362229e94cbb4291ebbac0033be95affb44ed17b`
Severity: MEDIUM. Confidence: HIGH.
Finding: Subprocess launched with variable

```go
115: 	write("audit-exceptions.json", encode(exceptions), 0o644)
116: 	command := exec.Command(script, root)
117: 	command.Env = append(os.Environ(),
```

## G204 in internal/scripttest/codeql_security_test.go:60

Test-only runner for scripts/codeql-security. The script path is this repository's file, and a fake CodeQL command that the test wrote stands in for the real CLI.

Source SHA-256: `d297bdf991398be289adf92ac888d1dc7ba258a05bf4b56a04bed4c18b217a3e`
Severity: MEDIUM. Confidence: HIGH.
Finding: Subprocess launched with variable

```go
59: 	}
60: 	command := exec.Command(script)
61: 	command.Env = append(os.Environ(), "XQUIK_CODEQL_BIN="+codeql, "FAKE_REPORT="+reportPath, "TMPDIR="+root)
```

## G204 in internal/scripttest/release_local_test.go:152

Test-only runner for bin/release-local. The path points at the copy in the test's temporary repository, and the arguments are literal flags from each test case. Fake gh and curl commands stand in for GitHub and the module proxy.

Source SHA-256: `f5a2e37d58b467e1f4629280271f180fb60c9aa11df697f65ac06c06ae6808d6`
Severity: MEDIUM. Confidence: HIGH.
Finding: Subprocess launched with a potential tainted input or cmd arguments

```go
151: 	fixture.t.Helper()
152: 	command := exec.Command(filepath.Join(fixture.repo, "bin", "release-local"), args...)
153: 	command.Dir, command.Env = fixture.t.TempDir(), fixture.env
```

## G306 in internal/scripttest/codeql_security_test.go:54

Test-only stand-in for the CodeQL CLI, written into the test's temporary folder. It must be executable, so it needs mode 0o755. Its content is a literal in the test.

Source SHA-256: `d297bdf991398be289adf92ac888d1dc7ba258a05bf4b56a04bed4c18b217a3e`
Severity: MEDIUM. Confidence: HIGH.
Finding: Expect WriteFile permissions to be 0600 or less

```go
53: 	reportPath := filepath.Join(root, "report.sarif")
54: 	if err := os.WriteFile(codeql, []byte(fakeCodeQL), 0o755); err != nil {
55: 		t.Fatal(err)
```
