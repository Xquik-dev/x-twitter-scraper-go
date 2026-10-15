// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io"
	"io/fs"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/Xquik-dev/x-twitter-scraper-go/internal/requestconfig"
	"github.com/Xquik-dev/x-twitter-scraper-go/option"
)

const (
	defaultSpecSource = "https://xquik.com/openapi.json"
	specVariable      = "XQUIK_OPENAPI_SPEC"
	exclusionFile     = "excluded-operations.json"
	requestFunction   = "ExecuteNewRequest"
	missingNextStep   = "Add each missing operation above to the SDK, or list it in excluded-operations.json with a reason."
)

var (
	placeholderPattern = regexp.MustCompile(`\{[^{}]*\}`)
	verbPattern        = regexp.MustCompile(`%[sdv]`)
	operationPattern   = regexp.MustCompile(`^([A-Z]+) (/\S*)$`)
	methodConstants    = map[string]string{
		"MethodGet": "GET", "MethodPut": "PUT", "MethodPost": "POST", "MethodDelete": "DELETE",
		"MethodOptions": "OPTIONS", "MethodHead": "HEAD", "MethodPatch": "PATCH", "MethodTrace": "TRACE",
	}
)

type operations map[string]string

type server struct {
	URL string `json:"url"`
}

type specOperation struct {
	OperationID string   `json:"operationId"`
	Servers     []server `json:"servers"`
}

type pathItem struct {
	Servers []server       `json:"servers"`
	Get     *specOperation `json:"get"`
	Put     *specOperation `json:"put"`
	Post    *specOperation `json:"post"`
	Delete  *specOperation `json:"delete"`
	Options *specOperation `json:"options"`
	Head    *specOperation `json:"head"`
	Patch   *specOperation `json:"patch"`
	Trace   *specOperation `json:"trace"`
}

func (p pathItem) methods() map[string]*specOperation {
	return map[string]*specOperation{
		"GET": p.Get, "PUT": p.Put, "POST": p.Post, "DELETE": p.Delete,
		"OPTIONS": p.Options, "HEAD": p.Head, "PATCH": p.Patch, "TRACE": p.Trace,
	}
}

type spec struct {
	Servers []server            `json:"servers"`
	Paths   map[string]pathItem `json:"paths"`
}

type exclusionGroup struct {
	Reason     string   `json:"reason"`
	Operations []string `json:"operations"`
}

func operationKey(method, operationPath string) string {
	return method + " " + placeholderPattern.ReplaceAllString(operationPath, "{}")
}

func joinPath(base, operationPath string) string {
	return strings.TrimSuffix(base, "/") + "/" + strings.TrimPrefix(operationPath, "/")
}

func knownMethod(method string) bool {
	for _, known := range methodConstants {
		if known == method {
			return true
		}
	}
	return false
}

func serverPath(servers ...[]server) (string, error) {
	for _, candidates := range servers {
		if len(candidates) == 0 {
			continue
		}
		parsed, err := url.Parse(candidates[0].URL)
		if err != nil {
			return "", err
		}
		return strings.TrimSuffix(parsed.Path, "/"), nil
	}
	return "", nil
}

func unreadableSpec(source string) error {
	return fmt.Errorf("Could not read the API spec from %s. Check network access, or set %s to a local copy of %s.", source, specVariable, defaultSpecSource)
}

func fetchSpec(source string, client *http.Client) ([]byte, error) {
	if !strings.HasPrefix(source, "https://") && !strings.HasPrefix(source, "http://") {
		return os.ReadFile(source)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, source, nil)
	if err != nil {
		return nil, err
	}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return nil, errors.New(response.Status)
	}
	return io.ReadAll(response.Body)
}

func specOperations(source string, client *http.Client) (operations, error) {
	data, err := fetchSpec(source, client)
	if err != nil {
		return nil, unreadableSpec(source)
	}
	var document spec
	if err := json.Unmarshal(data, &document); err != nil || len(document.Paths) == 0 {
		return nil, unreadableSpec(source)
	}
	found := operations{}
	for key, item := range document.Paths {
		for method, op := range item.methods() {
			if op == nil {
				continue
			}
			base, err := serverPath(op.Servers, item.Servers, document.Servers)
			if err != nil {
				return nil, unreadableSpec(source)
			}
			label := method + " " + joinPath(base, key)
			if op.OperationID != "" {
				label += " (" + op.OperationID + ")"
			}
			found[operationKey(method, joinPath(base, key))] = label
		}
	}
	return found, nil
}

func basePath(baseURL option.RequestOption) (string, error) {
	var config requestconfig.RequestConfig
	if err := baseURL.Apply(&config); err != nil || config.DefaultBaseURL == nil {
		return "", errors.New("The coverage check could not read the default base URL from option.WithEnvironmentProduction. Fix that option.")
	}
	return strings.TrimSuffix(config.DefaultBaseURL.Path, "/"), nil
}

func stringValue(expression ast.Expr) (string, bool) {
	literal, ok := expression.(*ast.BasicLit)
	if !ok || literal.Kind != token.STRING {
		return "", false
	}
	value, err := strconv.Unquote(literal.Value)
	return value, err == nil
}

func selectorOf(expression ast.Expr, pkg string) (string, bool) {
	selector, ok := expression.(*ast.SelectorExpr)
	if !ok {
		return "", false
	}
	ident, ok := selector.X.(*ast.Ident)
	return selector.Sel.Name, ok && ident.Name == pkg
}

func receiverName(function *ast.FuncDecl) string {
	if function.Recv == nil {
		return ""
	}
	expression := function.Recv.List[0].Type
	if star, ok := expression.(*ast.StarExpr); ok {
		expression = star.X
	}
	ident, _ := expression.(*ast.Ident)
	return ident.String()
}

func assignments(body *ast.BlockStmt, name string) []ast.Expr {
	var values []ast.Expr
	ast.Inspect(body, func(node ast.Node) bool {
		assignment, ok := node.(*ast.AssignStmt)
		if !ok || len(assignment.Lhs) != len(assignment.Rhs) {
			return true
		}
		for index, target := range assignment.Lhs {
			if ident, ok := target.(*ast.Ident); ok && ident.Name == name {
				values = append(values, assignment.Rhs[index])
			}
		}
		return true
	})
	return values
}

func requestPath(function *ast.FuncDecl, argument ast.Expr) (string, bool) {
	expression := argument
	if ident, ok := argument.(*ast.Ident); ok {
		values := assignments(function.Body, ident.Name)
		if len(values) != 1 {
			return "", false
		}
		expression = values[0]
	}
	if value, ok := stringValue(expression); ok {
		return value, true
	}
	call, ok := expression.(*ast.CallExpr)
	if !ok || len(call.Args) == 0 {
		return "", false
	}
	if name, ok := selectorOf(call.Fun, "fmt"); !ok || name != "Sprintf" {
		return "", false
	}
	format, ok := stringValue(call.Args[0])
	arguments := len(call.Args) - 1
	if !ok || len(verbPattern.FindAllString(format, -1)) != arguments || strings.Count(format, "%") != arguments {
		return "", false
	}
	return verbPattern.ReplaceAllString(format, "{}"), true
}

func requestMethod(argument ast.Expr) (string, bool) {
	name, ok := selectorOf(argument, "http")
	method, known := methodConstants[name]
	return method, ok && known
}

func serviceOperations(files *token.FileSet, function *ast.FuncDecl, base string, found operations) []string {
	var unreadable []string
	ast.Inspect(function.Body, func(node ast.Node) bool {
		call, ok := node.(*ast.CallExpr)
		if !ok {
			return true
		}
		selector, ok := call.Fun.(*ast.SelectorExpr)
		if !ok || selector.Sel.Name != requestFunction {
			return true
		}
		position := files.Position(call.Pos())
		if len(call.Args) < 3 {
			unreadable = append(unreadable, fmt.Sprintf("%s:%d calls %s without a method and a path. Pass both as arguments.", position.Filename, position.Line, requestFunction))
			return true
		}
		method, methodOK := requestMethod(call.Args[1])
		target, pathOK := requestPath(function, call.Args[2])
		if !methodOK || !pathOK {
			unreadable = append(unreadable, fmt.Sprintf("%s:%d sends a request whose method or path the coverage check cannot read. Use an http.Method constant and a path string or fmt.Sprintf format.", position.Filename, position.Line))
			return true
		}
		full := joinPath(base, target)
		found[operationKey(method, full)] = method + " " + full
		return true
	})
	return unreadable
}

func sdkOperations(root fs.FS, base string) (operations, []string, error) {
	entries, err := fs.ReadDir(root, ".")
	if err != nil {
		return nil, nil, errors.New("The coverage check could not list the SDK files. Run it from the repository root.")
	}
	found := operations{}
	var unreadable []string
	files := token.NewFileSet()
	for _, entry := range entries {
		name := entry.Name()
		if entry.IsDir() || !strings.HasSuffix(name, ".go") || strings.HasSuffix(name, "_test.go") {
			continue
		}
		source, err := fs.ReadFile(root, name)
		if err != nil {
			return nil, nil, fmt.Errorf("The coverage check could not read %s. Run it from the repository root.", name)
		}
		file, err := parser.ParseFile(files, name, source, parser.SkipObjectResolution)
		if err != nil {
			return nil, nil, fmt.Errorf("The coverage check could not parse %s. Run go vet ./... and fix the syntax error.", name)
		}
		for _, declaration := range file.Decls {
			function, ok := declaration.(*ast.FuncDecl)
			if ok && strings.HasSuffix(receiverName(function), "Service") {
				unreadable = append(unreadable, serviceOperations(files, function, base, found)...)
			}
		}
	}
	return found, unreadable, nil
}

func invalidExclusions(format string, args ...any) error {
	return fmt.Errorf("%s is not valid. %s.", exclusionFile, fmt.Sprintf(format, args...))
}

func readExclusions(root fs.FS) (operations, error) {
	data, err := fs.ReadFile(root, exclusionFile)
	if err != nil {
		return nil, invalidExclusions("The repository root has no readable copy. Create it, even as an empty list []")
	}
	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	var groups []exclusionGroup
	if err := decoder.Decode(&groups); err != nil {
		var syntax *json.SyntaxError
		if errors.As(err, &syntax) {
			return nil, invalidExclusions("It is not valid JSON near byte %d", syntax.Offset)
		}
		return nil, invalidExclusions("It must be a JSON list of groups, each with only a reason string and an operations list of strings")
	}
	if _, err := decoder.Token(); err != io.EOF {
		return nil, invalidExclusions("It holds more than 1 JSON value")
	}
	excluded := operations{}
	for index, group := range groups {
		number := index + 1
		if strings.TrimSpace(group.Reason) == "" {
			return nil, invalidExclusions("Group %d needs a non-empty reason", number)
		}
		if len(group.Operations) == 0 {
			return nil, invalidExclusions("Group %d needs at least 1 operation", number)
		}
		for _, text := range group.Operations {
			match := operationPattern.FindStringSubmatch(text)
			if match == nil || !knownMethod(match[1]) {
				return nil, invalidExclusions("Group %d lists %q, which is not an upper case HTTP method, a space and a path starting with /", number, text)
			}
			key := operationKey(match[1], match[2])
			if earlier, duplicate := excluded[key]; duplicate {
				return nil, invalidExclusions("It lists %s and %s, which name the same operation. List it once", earlier, text)
			}
			excluded[key] = text
		}
	}
	return excluded, nil
}

func check(root fs.FS, baseURL option.RequestOption, source string, client *http.Client) ([]string, string, error) {
	excluded, err := readExclusions(root)
	if err != nil {
		return nil, "", err
	}
	base, err := basePath(baseURL)
	if err != nil {
		return nil, "", err
	}
	implemented, unreadable, err := sdkOperations(root, base)
	if err != nil || len(unreadable) > 0 {
		return unreadable, "", err
	}
	api, err := specOperations(source, client)
	if err != nil {
		return nil, "", err
	}
	var missing, stale []string
	for key, label := range api {
		if implemented[key] == "" && excluded[key] == "" {
			missing = append(missing, label+" is in the API but not in the SDK.")
		}
	}
	for key, text := range excluded {
		switch {
		case api[key] == "":
			stale = append(stale, fmt.Sprintf("%s lists %s, which the API no longer has. Remove it.", exclusionFile, text))
		case implemented[key] != "":
			stale = append(stale, fmt.Sprintf("%s lists %s, which the SDK already implements. Remove it.", exclusionFile, text))
		}
	}
	slices.Sort(missing)
	slices.Sort(stale)
	if len(missing) > 0 {
		missing = append(missing, missingNextStep)
	}
	if failures := append(missing, stale...); len(failures) > 0 {
		return failures, "", nil
	}
	covered := len(api) - len(excluded)
	return nil, fmt.Sprintf("The SDK covers %d of %d API operations, and %s lists the other %d.", covered, len(api), exclusionFile, len(excluded)), nil
}

func run(source string, root fs.FS, baseURL option.RequestOption, client *http.Client, out, failures io.Writer) int {
	if source == "" {
		source = defaultSpecSource
	}
	lines, summary, err := check(root, baseURL, source, client)
	if err != nil {
		lines = append(lines, err.Error())
	}
	for _, line := range lines {
		fmt.Fprintln(failures, line)
	}
	if len(lines) > 0 {
		return 1
	}
	fmt.Fprintln(out, summary)
	return 0
}

func main() {
	if code := run(os.Getenv(specVariable), os.DirFS("."), option.WithEnvironmentProduction(), http.DefaultClient, os.Stdout, os.Stderr); code != 0 {
		os.Exit(code)
	}
}
