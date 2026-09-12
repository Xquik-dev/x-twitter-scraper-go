# Approved source findings

The user approved these 123 exact findings on 2026-09-12.
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

## G404 in internal/requestconfig/requestconfig.go:347

Bounded retry jitter, not a credential, key, authentication value, or security token. The arithmetic overflow was repaired and verified on ARM64 and AMD64. Cryptographic unpredictability is not a requirement for jitter.

Source SHA-256: `dd7d082a1b1b8ed246bbac74b1176072589eeb6fe699174da2ce61188b5f5616`
Severity: HIGH. Confidence: MEDIUM.
Finding: Use of weak random number generator (math/rand or math/rand/v2 instead of crypto/rand)

```go
346: 	delay := 500 * time.Millisecond << min(max(retryCount, 0), 4)
347: 	return delay - time.Duration(rand.Int63n(int64(delay/4)))
348: }
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

## G101 in guestwallet.go:388

Public generated credential-storage notice, not a credential value.

Source SHA-256: `c7f19833a2d77b5463a36f06305e11b3bf8f8672b358acb89ffcbe9f2b412a9f`
Severity: HIGH. Confidence: LOW.
Finding: Potential hardcoded credentials

```go
387: const (
388: 	GuestWalletTopupResponseCredentialNoticeStoreAPIKeyAndTheIdempotencyKeySecurelyBeforeSharingCheckoutURLNoEmailRecoveryIsAvailable GuestWalletTopupResponseCredentialNotice = "Store api_key and the Idempotency-Key securely before sharing checkout_url. No email recovery is available."
389: )
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

## G703 in cmd/coverage-gain/main.go:55

Explicit local CLI output paths. No remote path source or claimed confinement boundary.

Source SHA-256: `231c1e43b633c36028f604e68e0894b7ef5e38057d56a5b23091d8a0cda532be`
Severity: HIGH. Confidence: HIGH.
Finding: Path traversal via taint analysis

```go
54: 	}
55: 	if err := os.WriteFile(filepath.Join(directory, "summary.json"), data, 0o600); err != nil {
56: 		return err
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

## G204 in cmd/coverage-gain/main.go:19

Fixed executable and argument array. Build helpers consume local Git arguments; upload test invokes the fixed repository script. No shell interpolation of remote data.

Source SHA-256: `231c1e43b633c36028f604e68e0894b7ef5e38057d56a5b23091d8a0cda532be`
Severity: MEDIUM. Confidence: HIGH.
Finding: Subprocess launched with variable

```go
18: func git(args ...string) (string, error) {
19: 	command := exec.Command("git", args...)
20: 	command.Env = append(os.Environ(), "LC_ALL=C")
```

## G204 in cmd/coverage-gain/main_test.go:26

Fixed executable and argument array. Build helpers consume local Git arguments; upload test invokes the fixed repository script. No shell interpolation of remote data.

Source SHA-256: `2fac5356ad419874263c4eea2e6250d9bec9cb121ab89dfc828ef722e1acf1a3`
Severity: MEDIUM. Confidence: HIGH.
Finding: Subprocess launched with a potential tainted input or cmd arguments

```go
25: 		t.Helper()
26: 		command := exec.Command("git", append([]string{"-c", "user.name=Coverage test", "-c", "user.email=coverage@example.test", "-c", "commit.gpgsign=false"}, args...)...)
27: 		if data, err := command.CombinedOutput(); err != nil {
```

## G304 in cmd/coverage-gain/main_test.go:84

Test reads its own summary artifact within a temporary directory.

Source SHA-256: `2fac5356ad419874263c4eea2e6250d9bec9cb121ab89dfc828ef722e1acf1a3`
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

## G104 in xuser.go:868

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `1d063a0039f2fc2a9ad119e7d2ee1d8347facc1b08aec58c591cb66bc933ccba`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
867: func (u XUserGetVerifiedFollowersResponseUnion) AsXUserGetVerifiedFollowersResponseUserListCoverageResponse() (v XUserGetVerifiedFollowersResponseUserListCoverageResponse) {
868: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
869: 	return
```

## G104 in xuser.go:863

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `1d063a0039f2fc2a9ad119e7d2ee1d8347facc1b08aec58c591cb66bc933ccba`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
862: func (u XUserGetVerifiedFollowersResponseUnion) AsPaginatedUsers() (v shared.PaginatedUsers) {
863: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
864: 	return
```

## G104 in xuser.go:721

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `1d063a0039f2fc2a9ad119e7d2ee1d8347facc1b08aec58c591cb66bc933ccba`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
720: func (u XUserGetFollowingResponseUnion) AsXUserGetFollowingResponseUserListCoverageResponse() (v XUserGetFollowingResponseUserListCoverageResponse) {
721: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
722: 	return
```

## G104 in xuser.go:716

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `1d063a0039f2fc2a9ad119e7d2ee1d8347facc1b08aec58c591cb66bc933ccba`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
715: func (u XUserGetFollowingResponseUnion) AsPaginatedUsers() (v shared.PaginatedUsers) {
716: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
717: 	return
```

## G104 in xuser.go:576

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `1d063a0039f2fc2a9ad119e7d2ee1d8347facc1b08aec58c591cb66bc933ccba`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
575: func (u XUserGetFollowersResponseUnion) AsXUserGetFollowersResponseUserListCoverageResponse() (v XUserGetFollowersResponseUserListCoverageResponse) {
576: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
577: 	return
```

## G104 in xuser.go:571

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `1d063a0039f2fc2a9ad119e7d2ee1d8347facc1b08aec58c591cb66bc933ccba`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
570: func (u XUserGetFollowersResponseUnion) AsPaginatedUsers() (v shared.PaginatedUsers) {
571: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
572: 	return
```

## G104 in xtweet.go:2312

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `06861080b25ee45bb1bb2ebcfaddf93037a3cc2a6c0ca83dd7e4b70b2fdd73e0`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
2311: func (u XTweetSearchResponseUnion) AsXTweetSearchResponseTweetSearchCoverageResponse() (v XTweetSearchResponseTweetSearchCoverageResponse) {
2312: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
2313: 	return
```

## G104 in xtweet.go:2307

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `06861080b25ee45bb1bb2ebcfaddf93037a3cc2a6c0ca83dd7e4b70b2fdd73e0`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
2306: func (u XTweetSearchResponseUnion) AsPaginatedTweets() (v shared.PaginatedTweets) {
2307: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
2308: 	return
```

## G104 in xtweet.go:1402

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `06861080b25ee45bb1bb2ebcfaddf93037a3cc2a6c0ca83dd7e4b70b2fdd73e0`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
1401: func (u TweetDetailSportsContextScheduledAtMsUnion) AsString() (v string) {
1402: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
1403: 	return
```

## G104 in xtweet.go:1397

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `06861080b25ee45bb1bb2ebcfaddf93037a3cc2a6c0ca83dd7e4b70b2fdd73e0`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
1396: func (u TweetDetailSportsContextScheduledAtMsUnion) AsFloat() (v float64) {
1397: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
1398: 	return
```

## G104 in xlist.go:262

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `e58c78c6bec4232000474884ee13a60b118fdb7621ff995c132af68bf9fc510d`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
261: func (u XListGetMembersResponseUnion) AsXListGetMembersResponseUserListCoverageResponse() (v XListGetMembersResponseUserListCoverageResponse) {
262: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
263: 	return
```

## G104 in xlist.go:257

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `e58c78c6bec4232000474884ee13a60b118fdb7621ff995c132af68bf9fc510d`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
256: func (u XListGetMembersResponseUnion) AsPaginatedUsers() (v shared.PaginatedUsers) {
257: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
258: 	return
```

## G104 in xlist.go:117

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `e58c78c6bec4232000474884ee13a60b118fdb7621ff995c132af68bf9fc510d`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
116: func (u XListGetFollowersResponseUnion) AsXListGetFollowersResponseUserListCoverageResponse() (v XListGetFollowersResponseUserListCoverageResponse) {
117: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
118: 	return
```

## G104 in xlist.go:112

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `e58c78c6bec4232000474884ee13a60b118fdb7621ff995c132af68bf9fc510d`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
111: func (u XListGetFollowersResponseUnion) AsPaginatedUsers() (v shared.PaginatedUsers) {
112: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
113: 	return
```

## G104 in xaccountconnectionattempt.go:155

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `db32cf9309638f2e6e60c21383b9e6c323c1e0180984b435d712892014a7dcef`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
154: func (u XAccountConnectionAttemptGetResponseUnion) AsRequiresEmailCode() (v XAccountConnectionAttemptGetResponseRequiresEmailCode) {
155: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
156: 	return
```

## G104 in xaccountconnectionattempt.go:150

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `db32cf9309638f2e6e60c21383b9e6c323c1e0180984b435d712892014a7dcef`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
149: func (u XAccountConnectionAttemptGetResponseUnion) AsFailed() (v XAccountConnectionAttemptGetResponseFailed) {
150: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
151: 	return
```

## G104 in xaccountconnectionattempt.go:145

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `db32cf9309638f2e6e60c21383b9e6c323c1e0180984b435d712892014a7dcef`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
144: func (u XAccountConnectionAttemptGetResponseUnion) AsSuccess() (v XAccountConnectionAttemptGetResponseSuccess) {
145: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
146: 	return
```

## G104 in xaccountconnectionattempt.go:140

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `db32cf9309638f2e6e60c21383b9e6c323c1e0180984b435d712892014a7dcef`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
139: func (u XAccountConnectionAttemptGetResponseUnion) AsPending() (v XAccountConnectionAttemptGetResponsePending) {
140: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
141: 	return
```

## G104 in xaccount.go:260

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `e9848f5ae238bd6b72833cd78f5b4e056809770e08bf6f1a5f0531ccb3f10460`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
259: func (u XAccountNewResponseUnion) AsXAccountNewResponseXAccountConnectionChallenge() (v XAccountNewResponseXAccountConnectionChallenge) {
260: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
261: 	return
```

## G104 in xaccount.go:255

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `e9848f5ae238bd6b72833cd78f5b4e056809770e08bf6f1a5f0531ccb3f10460`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
254: func (u XAccountNewResponseUnion) AsXAccountNewResponseXAccountConnectionAttemptPending() (v XAccountNewResponseXAccountConnectionAttemptPending) {
255: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
256: 	return
```

## G104 in xaccount.go:250

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `e9848f5ae238bd6b72833cd78f5b4e056809770e08bf6f1a5f0531ccb3f10460`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
249: func (u XAccountNewResponseUnion) AsXAccountNewResponseSanitizedXAccount() (v XAccountNewResponseSanitizedXAccount) {
250: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
251: 	return
```

## G104 in supportattachment_test.go:21

Local test HTTP handler writes fixture output. Client-side response/result assertions cover delivered bytes; this call is not a production sink.

Source SHA-256: `78d21614d3d3d7a2d96cf1c2ec482ed79fd8cb3fb7ef8767d4366c1580cc6b36`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
20: 		w.WriteHeader(200)
21: 		w.Write([]byte("abc"))
22: 	}))
```

## G104 in shared/shared.go:21859

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `3d46ca3e17c8d7abb409b11169013034d0dbc69f73217f491a3b5ede0f4cab32`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
21858: func (u SearchTweetSportsContextScheduledAtMsUnion) AsString() (v string) {
21859: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
21860: 	return
```

## G104 in shared/shared.go:21854

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `3d46ca3e17c8d7abb409b11169013034d0dbc69f73217f491a3b5ede0f4cab32`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
21853: func (u SearchTweetSportsContextScheduledAtMsUnion) AsFloat() (v float64) {
21854: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
21855: 	return
```

## G104 in shared/shared.go:20482

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `3d46ca3e17c8d7abb409b11169013034d0dbc69f73217f491a3b5ede0f4cab32`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
20481: func (u EmbeddedTweetSportsContextScheduledAtMsUnion) AsString() (v string) {
20482: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
20483: 	return
```

## G104 in shared/shared.go:20477

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `3d46ca3e17c8d7abb409b11169013034d0dbc69f73217f491a3b5ede0f4cab32`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
20476: func (u EmbeddedTweetSportsContextScheduledAtMsUnion) AsFloat() (v float64) {
20477: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
20478: 	return
```

## G104 in shared/shared.go:20294

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `3d46ca3e17c8d7abb409b11169013034d0dbc69f73217f491a3b5ede0f4cab32`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
20293: func (u EmbeddedTweetRetweetedTweetSportsContextScheduledAtMsUnion) AsString() (v string) {
20294: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
20295: 	return
```

## G104 in shared/shared.go:20289

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `3d46ca3e17c8d7abb409b11169013034d0dbc69f73217f491a3b5ede0f4cab32`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
20288: func (u EmbeddedTweetRetweetedTweetSportsContextScheduledAtMsUnion) AsFloat() (v float64) {
20289: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
20290: 	return
```

## G104 in shared/shared.go:20102

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `3d46ca3e17c8d7abb409b11169013034d0dbc69f73217f491a3b5ede0f4cab32`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
20101: func (u EmbeddedTweetRetweetedTweetRetweetedTweetSportsContextScheduledAtMsUnion) AsString() (v string) {
20102: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
20103: 	return
```

## G104 in shared/shared.go:20097

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `3d46ca3e17c8d7abb409b11169013034d0dbc69f73217f491a3b5ede0f4cab32`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
20096: func (u EmbeddedTweetRetweetedTweetRetweetedTweetSportsContextScheduledAtMsUnion) AsFloat() (v float64) {
20097: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
20098: 	return
```

## G104 in shared/shared.go:19904

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `3d46ca3e17c8d7abb409b11169013034d0dbc69f73217f491a3b5ede0f4cab32`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
19903: func (u EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetSportsContextScheduledAtMsUnion) AsString() (v string) {
19904: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
19905: 	return
```

## G104 in shared/shared.go:19899

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `3d46ca3e17c8d7abb409b11169013034d0dbc69f73217f491a3b5ede0f4cab32`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
19898: func (u EmbeddedTweetRetweetedTweetRetweetedTweetRetweetedTweetSportsContextScheduledAtMsUnion) AsFloat() (v float64) {
19899: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
19900: 	return
```

## G104 in shared/shared.go:18485

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `3d46ca3e17c8d7abb409b11169013034d0dbc69f73217f491a3b5ede0f4cab32`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
18484: func (u EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetSportsContextScheduledAtMsUnion) AsString() (v string) {
18485: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
18486: 	return
```

## G104 in shared/shared.go:18480

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `3d46ca3e17c8d7abb409b11169013034d0dbc69f73217f491a3b5ede0f4cab32`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
18479: func (u EmbeddedTweetRetweetedTweetRetweetedTweetQuotedTweetSportsContextScheduledAtMsUnion) AsFloat() (v float64) {
18480: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
18481: 	return
```

## G104 in shared/shared.go:15917

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `3d46ca3e17c8d7abb409b11169013034d0dbc69f73217f491a3b5ede0f4cab32`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
15916: func (u EmbeddedTweetRetweetedTweetQuotedTweetSportsContextScheduledAtMsUnion) AsString() (v string) {
15917: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
15918: 	return
```

## G104 in shared/shared.go:15912

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `3d46ca3e17c8d7abb409b11169013034d0dbc69f73217f491a3b5ede0f4cab32`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
15911: func (u EmbeddedTweetRetweetedTweetQuotedTweetSportsContextScheduledAtMsUnion) AsFloat() (v float64) {
15912: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
15913: 	return
```

## G104 in shared/shared.go:15719

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `3d46ca3e17c8d7abb409b11169013034d0dbc69f73217f491a3b5ede0f4cab32`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
15718: func (u EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetSportsContextScheduledAtMsUnion) AsString() (v string) {
15719: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
15720: 	return
```

## G104 in shared/shared.go:15714

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `3d46ca3e17c8d7abb409b11169013034d0dbc69f73217f491a3b5ede0f4cab32`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
15713: func (u EmbeddedTweetRetweetedTweetQuotedTweetRetweetedTweetSportsContextScheduledAtMsUnion) AsFloat() (v float64) {
15714: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
15715: 	return
```

## G104 in shared/shared.go:14304

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `3d46ca3e17c8d7abb409b11169013034d0dbc69f73217f491a3b5ede0f4cab32`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
14303: func (u EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetSportsContextScheduledAtMsUnion) AsString() (v string) {
14304: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
14305: 	return
```

## G104 in shared/shared.go:14299

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `3d46ca3e17c8d7abb409b11169013034d0dbc69f73217f491a3b5ede0f4cab32`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
14298: func (u EmbeddedTweetRetweetedTweetQuotedTweetQuotedTweetSportsContextScheduledAtMsUnion) AsFloat() (v float64) {
14299: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
14300: 	return
```

## G104 in shared/shared.go:10651

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `3d46ca3e17c8d7abb409b11169013034d0dbc69f73217f491a3b5ede0f4cab32`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
10650: func (u EmbeddedTweetQuotedTweetSportsContextScheduledAtMsUnion) AsString() (v string) {
10651: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
10652: 	return
```

## G104 in shared/shared.go:10646

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `3d46ca3e17c8d7abb409b11169013034d0dbc69f73217f491a3b5ede0f4cab32`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
10645: func (u EmbeddedTweetQuotedTweetSportsContextScheduledAtMsUnion) AsFloat() (v float64) {
10646: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
10647: 	return
```

## G104 in shared/shared.go:10459

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `3d46ca3e17c8d7abb409b11169013034d0dbc69f73217f491a3b5ede0f4cab32`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
10458: func (u EmbeddedTweetQuotedTweetRetweetedTweetSportsContextScheduledAtMsUnion) AsString() (v string) {
10459: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
10460: 	return
```

## G104 in shared/shared.go:10454

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `3d46ca3e17c8d7abb409b11169013034d0dbc69f73217f491a3b5ede0f4cab32`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
10453: func (u EmbeddedTweetQuotedTweetRetweetedTweetSportsContextScheduledAtMsUnion) AsFloat() (v float64) {
10454: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
10455: 	return
```

## G104 in shared/shared.go:10261

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `3d46ca3e17c8d7abb409b11169013034d0dbc69f73217f491a3b5ede0f4cab32`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
10260: func (u EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetSportsContextScheduledAtMsUnion) AsString() (v string) {
10261: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
10262: 	return
```

## G104 in shared/shared.go:10256

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `3d46ca3e17c8d7abb409b11169013034d0dbc69f73217f491a3b5ede0f4cab32`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
10255: func (u EmbeddedTweetQuotedTweetRetweetedTweetRetweetedTweetSportsContextScheduledAtMsUnion) AsFloat() (v float64) {
10256: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
10257: 	return
```

## G104 in shared/shared.go:8846

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `3d46ca3e17c8d7abb409b11169013034d0dbc69f73217f491a3b5ede0f4cab32`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
8845: func (u EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetSportsContextScheduledAtMsUnion) AsString() (v string) {
8846: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
8847: 	return
```

## G104 in shared/shared.go:8841

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `3d46ca3e17c8d7abb409b11169013034d0dbc69f73217f491a3b5ede0f4cab32`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
8840: func (u EmbeddedTweetQuotedTweetRetweetedTweetQuotedTweetSportsContextScheduledAtMsUnion) AsFloat() (v float64) {
8841: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
8842: 	return
```

## G104 in shared/shared.go:6292

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `3d46ca3e17c8d7abb409b11169013034d0dbc69f73217f491a3b5ede0f4cab32`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
6291: func (u EmbeddedTweetQuotedTweetQuotedTweetSportsContextScheduledAtMsUnion) AsString() (v string) {
6292: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
6293: 	return
```

## G104 in shared/shared.go:6287

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `3d46ca3e17c8d7abb409b11169013034d0dbc69f73217f491a3b5ede0f4cab32`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
6286: func (u EmbeddedTweetQuotedTweetQuotedTweetSportsContextScheduledAtMsUnion) AsFloat() (v float64) {
6287: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
6288: 	return
```

## G104 in shared/shared.go:6094

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `3d46ca3e17c8d7abb409b11169013034d0dbc69f73217f491a3b5ede0f4cab32`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
6093: func (u EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetSportsContextScheduledAtMsUnion) AsString() (v string) {
6094: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
6095: 	return
```

## G104 in shared/shared.go:6089

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `3d46ca3e17c8d7abb409b11169013034d0dbc69f73217f491a3b5ede0f4cab32`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
6088: func (u EmbeddedTweetQuotedTweetQuotedTweetRetweetedTweetSportsContextScheduledAtMsUnion) AsFloat() (v float64) {
6089: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
6090: 	return
```

## G104 in shared/shared.go:4685

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `3d46ca3e17c8d7abb409b11169013034d0dbc69f73217f491a3b5ede0f4cab32`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
4684: func (u EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetSportsContextScheduledAtMsUnion) AsString() (v string) {
4685: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
4686: 	return
```

## G104 in shared/shared.go:4680

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `3d46ca3e17c8d7abb409b11169013034d0dbc69f73217f491a3b5ede0f4cab32`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
4679: func (u EmbeddedTweetQuotedTweetQuotedTweetQuotedTweetSportsContextScheduledAtMsUnion) AsFloat() (v float64) {
4680: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
4681: 	return
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

## G104 in internal/requestconfig/requestconfig.go:495

Media type classification intentionally ignores parameter parse errors. A valid media type remains usable when parameters fail; invalid types fall through to typed plaintext handling rather than accepting arbitrary structured data.

Source SHA-256: `dd7d082a1b1b8ed246bbac74b1176072589eeb6fe699174da2ce61188b5f5616`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
494: 	contentType := res.Header.Get("content-type")
495: 	mediaType, _, _ := mime.ParseMediaType(contentType)
496: 	isJSON := strings.Contains(mediaType, "application/json") || strings.HasSuffix(mediaType, "+json")
```

## G104 in internal/requestconfig/requestconfig.go:488

Response cleanup after read or before retry. Read/request errors are handled separately; cleanup error does not replace the primary request result. Close is still invoked, and raw-response ownership is transferred explicitly elsewhere.

Source SHA-256: `dd7d082a1b1b8ed246bbac74b1176072589eeb6fe699174da2ce61188b5f5616`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
487: 	contents, err := io.ReadAll(res.Body)
488: 	_ = res.Body.Close()
489: 	if err != nil {
```

## G104 in internal/requestconfig/requestconfig.go:457

Response cleanup after read or before retry. Read/request errors are handled separately; cleanup error does not replace the primary request result. Close is still invoked, and raw-response ownership is transferred explicitly elsewhere.

Source SHA-256: `dd7d082a1b1b8ed246bbac74b1176072589eeb6fe699174da2ce61188b5f5616`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
456: 		contents, err := io.ReadAll(res.Body)
457: 		_ = res.Body.Close()
458: 		if err != nil {
```

## G104 in internal/requestconfig/requestconfig.go:429

Response cleanup after read or before retry. Read/request errors are handled separately; cleanup error does not replace the primary request result. Close is still invoked, and raw-response ownership is transferred explicitly elsewhere.

Source SHA-256: `dd7d082a1b1b8ed246bbac74b1176072589eeb6fe699174da2ce61188b5f5616`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
428: 		if res != nil && res.Body != nil {
429: 			_ = res.Body.Close()
430: 		}
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

## G104 in internal/apierror/apierror.go:71

Best-effort byte-only diagnostic API. Dump errors produce no dump; no request is sent as an operation. Credentials are redacted on copies before dump. Request/response redaction regression suite covers original immutability.

Source SHA-256: `0f429d3cd33047decd91f0debf0ba582ee38cda7e7e47312807b536511465ec5`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
70: 	}
71: 	out, _ := httputil.DumpResponse(&responseCopy, body)
72: 	return out
```

## G104 in internal/apierror/apierror.go:58

Best-effort byte-only diagnostic API. Dump errors produce no dump; no request is sent as an operation. Credentials are redacted on copies before dump. Request/response redaction regression suite covers original immutability.

Source SHA-256: `0f429d3cd33047decd91f0debf0ba582ee38cda7e7e47312807b536511465ec5`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
57: 	}
58: 	out, _ := httputil.DumpRequestOut(&requestCopy, body)
59: 	return out
```

## G104 in extraction_test.go:244

Local test HTTP handler writes fixture output. Client-side response/result assertions cover delivered bytes; this call is not a production sink.

Source SHA-256: `54090cdf4d4a63ab0336ea1f79ffa7112d54163ba2a118ce159b7fb726a1fba4`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
243: 		w.WriteHeader(200)
244: 		w.Write([]byte("abc"))
245: 	}))
```

## G104 in draw_test.go:76

Local test HTTP handler writes fixture output. Client-side response/result assertions cover delivered bytes; this call is not a production sink.

Source SHA-256: `57054125d50556fba99293b55e321da8f9bb38f47cf825c6f4ce0631de5dc6b4`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
75: 		w.WriteHeader(200)
76: 		w.Write([]byte("abc"))
77: 	}))
```

## G104 in compose.go:133

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `7925a0632008a7cc9fdc3ffa2d5e5f18876971c17250b9fa0508c257b2824446`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
132: func (u ComposeNewResponseUnion) AsComposeNewResponseComposeScoreResult() (v ComposeNewResponseComposeScoreResult) {
133: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
134: 	return
```

## G104 in compose.go:128

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `7925a0632008a7cc9fdc3ffa2d5e5f18876971c17250b9fa0508c257b2824446`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
127: func (u ComposeNewResponseUnion) AsComposeNewResponseComposeRefineResult() (v ComposeNewResponseComposeRefineResult) {
128: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
129: 	return
```

## G104 in compose.go:123

Generated best-effort As variant conversion, signature returns a value without an error channel. README response union contract and numeric/string/connection variant regressions preserve this behavior. Successful network JSON responses are parsed by encoding/json decoder before model conversion.

Source SHA-256: `7925a0632008a7cc9fdc3ffa2d5e5f18876971c17250b9fa0508c257b2824446`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
122: func (u ComposeNewResponseUnion) AsComposeNewResponseComposePrepareResult() (v ComposeNewResponseComposePrepareResult) {
123: 	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
124: 	return
```

## G104 in client_test.go:46

This test asserts the observed User-Agent header. Endpoint-result correctness is covered separately; transport error is not the subject of this metadata assertion.

Source SHA-256: `c9b7185c7ce3f7de9b5474de81ec3b0e90bb8dbbdd889ab7bd478cec18f1baca`
Severity: LOW. Confidence: HIGH.
Finding: Errors unhandled

```go
45: 	)
46: 	_, _ = client.Account.Get(context.Background())
47: 	if userAgent != fmt.Sprintf("XTwitterScraper/Go %s", internal.PackageVersion) {
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
