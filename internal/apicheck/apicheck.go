// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

// Package apicheck compares the operations in the Xquik OpenAPI spec with the
// requests the SDK's service methods send. The api-coverage and api-fields
// commands run it.
package apicheck

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
	"strconv"
	"strings"
	"time"

	"github.com/Xquik-dev/x-twitter-scraper-go/internal/requestconfig"
	"github.com/Xquik-dev/x-twitter-scraper-go/option"
)

const (
	DefaultSpecSource = "https://xquik.com/openapi.json"
	SpecVariable      = "XQUIK_OPENAPI_SPEC"
	requestFunction   = "ExecuteNewRequest"
)

type Check string

const (
	Coverage Check = "coverage check"
	Fields   Check = "field check"
)

var (
	placeholderPattern = regexp.MustCompile(`\{[^{}]*\}`)
	verbPattern        = regexp.MustCompile(`%[sdv]`)
	httpMethods        = []string{"get", "put", "post", "delete", "options", "head", "patch", "trace"}
	methodConstants    = map[string]string{
		"MethodGet": "GET", "MethodPut": "PUT", "MethodPost": "POST", "MethodDelete": "DELETE",
		"MethodOptions": "OPTIONS", "MethodHead": "HEAD", "MethodPatch": "PATCH", "MethodTrace": "TRACE",
	}
)

type SpecOperation struct {
	Method      string
	Path        string
	OperationID string
	Item        PathItem
	Operation   Operation
}

type PathItem struct {
	Servers    json.RawMessage   `json:"servers"`
	Parameters []json.RawMessage `json:"parameters"`
}

type Operation struct {
	OperationID string            `json:"operationId"`
	Servers     json.RawMessage   `json:"servers"`
	Parameters  []json.RawMessage `json:"parameters"`
	RequestBody json.RawMessage   `json:"requestBody"`
	Responses   json.RawMessage   `json:"responses"`
}

func (o SpecOperation) Label() string {
	label := o.Method + " " + o.Path
	if o.OperationID != "" {
		label += " (" + o.OperationID + ")"
	}
	return label
}

type Spec struct {
	Document   json.RawMessage
	Operations map[string]SpecOperation
}

type failure string

func (f failure) Error() string { return string(f) }

func Decode[T any](raw json.RawMessage) (T, bool) {
	var value T
	err := json.Unmarshal(raw, &value)
	return value, err == nil
}

func isObject(raw json.RawMessage) bool {
	return len(raw) > 0 && raw[0] == '{'
}

type SDKOperation struct {
	Position token.Position
	Function *ast.FuncDecl
	Call     *ast.CallExpr
}

type GroupFile struct {
	Name      string
	ItemsKey  string
	Item      string
	ItemShape string
	Parse     func(string) (string, bool)
}

func Source(configured string) string {
	if configured == "" {
		return DefaultSpecSource
	}
	return configured
}

func OperationKey(method, operationPath string) string {
	return method + " " + placeholderPattern.ReplaceAllString(operationPath, "{}")
}

func JoinPath(base, operationPath string) string {
	return strings.TrimSuffix(base, "/") + "/" + strings.TrimPrefix(operationPath, "/")
}

func KnownMethod(method string) bool {
	for _, known := range methodConstants {
		if known == method {
			return true
		}
	}
	return false
}

func UnreadableSpec(source string) error {
	return failure(fmt.Sprintf("Could not read the API spec from %s. Check network access, or set %s to a local copy of %s.", source, SpecVariable, DefaultSpecSource))
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

func serverPath(candidates ...json.RawMessage) (string, error) {
	for _, value := range candidates {
		if len(value) == 0 || string(value) == "null" {
			continue
		}
		servers, ok := Decode[[]json.RawMessage](value)
		if !ok {
			return "", errors.New("servers is not a list")
		}
		if len(servers) == 0 {
			continue
		}
		first, _ := Decode[struct {
			URL string `json:"url"`
		}](servers[0])
		parsed, err := url.Parse(first.URL)
		if err != nil {
			return "", err
		}
		return strings.TrimSuffix(parsed.Path, "/"), nil
	}
	return "", nil
}

func LoadSpec(source string, client *http.Client) (*Spec, error) {
	data, err := fetchSpec(source, client)
	if err != nil {
		return nil, UnreadableSpec(source)
	}
	var document struct {
		Servers json.RawMessage            `json:"servers"`
		Paths   map[string]json.RawMessage `json:"paths"`
	}
	if err := json.Unmarshal(data, &document); err != nil || len(document.Paths) == 0 {
		return nil, UnreadableSpec(source)
	}
	operations := map[string]SpecOperation{}
	for key, value := range document.Paths {
		if !isObject(value) {
			return nil, UnreadableSpec(source)
		}
		methods, _ := Decode[map[string]json.RawMessage](value)
		item, _ := Decode[PathItem](value)
		for _, method := range httpMethods {
			raw, present := methods[method]
			if !present {
				continue
			}
			if !isObject(raw) {
				return nil, UnreadableSpec(source)
			}
			operation, _ := Decode[Operation](raw)
			base, err := serverPath(operation.Servers, item.Servers, document.Servers)
			if err != nil {
				return nil, UnreadableSpec(source)
			}
			full := JoinPath(base, key)
			upper := strings.ToUpper(method)
			operations[OperationKey(upper, full)] = SpecOperation{Method: upper, Path: full, OperationID: operation.OperationID, Item: item, Operation: operation}
		}
	}
	return &Spec{Document: data, Operations: operations}, nil
}

func (c Check) BasePath(baseURL option.RequestOption) (string, error) {
	var config requestconfig.RequestConfig
	if err := baseURL.Apply(&config); err != nil || config.DefaultBaseURL == nil {
		return "", failure(fmt.Sprintf("The %s could not read the default base URL from option.WithEnvironmentProduction. Fix that option.", c))
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
	if ident, ok := expression.(*ast.Ident); ok {
		return ident.Name
	}
	return ""
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

func (c Check) serviceOperations(files *token.FileSet, function *ast.FuncDecl, base string, found map[string]SDKOperation) []string {
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
			unreadable = append(unreadable, fmt.Sprintf("%s:%d sends a request whose method or path the %s cannot read. Use an http.Method constant and a path string or fmt.Sprintf format.", position.Filename, position.Line, c))
			return true
		}
		key := OperationKey(method, JoinPath(base, target))
		if _, seen := found[key]; !seen {
			found[key] = SDKOperation{Position: position, Function: function, Call: call}
		}
		return true
	})
	return unreadable
}

func (c Check) ParseGoFiles(root fs.FS, dir string, files *token.FileSet) ([]*ast.File, error) {
	entries, err := fs.ReadDir(root, dir)
	if err != nil {
		return nil, failure(fmt.Sprintf("The %s could not list the SDK files. Run it from the repository root.", c))
	}
	var parsed []*ast.File
	for _, entry := range entries {
		name := entry.Name()
		if entry.IsDir() || !strings.HasSuffix(name, ".go") || strings.HasSuffix(name, "_test.go") {
			continue
		}
		if dir != "." {
			name = dir + "/" + name
		}
		source, err := fs.ReadFile(root, name)
		if err != nil {
			return nil, failure(fmt.Sprintf("The %s could not read %s. Run it from the repository root.", c, name))
		}
		file, err := parser.ParseFile(files, name, source, parser.SkipObjectResolution)
		if err != nil {
			return nil, failure(fmt.Sprintf("The %s could not parse %s. Run go vet ./... and fix the syntax error.", c, name))
		}
		parsed = append(parsed, file)
	}
	return parsed, nil
}

func (c Check) SDKOperations(root fs.FS, base string) (map[string]SDKOperation, []string, error) {
	files := token.NewFileSet()
	parsed, err := c.ParseGoFiles(root, ".", files)
	if err != nil {
		return nil, nil, err
	}
	found := map[string]SDKOperation{}
	var unreadable []string
	for _, file := range parsed {
		for _, declaration := range file.Decls {
			function, ok := declaration.(*ast.FuncDecl)
			if ok && strings.HasSuffix(receiverName(function), "Service") {
				unreadable = append(unreadable, c.serviceOperations(files, function, base, found)...)
			}
		}
	}
	return found, unreadable, nil
}

func (g GroupFile) invalid(problem string) error {
	return failure(fmt.Sprintf("%s is not valid. %s.", g.Name, problem))
}

func (g GroupFile) Read(root fs.FS) (map[string]string, error) {
	data, err := fs.ReadFile(root, g.Name)
	if err != nil {
		return nil, g.invalid("The repository root has no readable copy. Create it, even as an empty list []")
	}
	shape := fmt.Sprintf("It must be a JSON list of groups, each with only a reason string and a list of strings named %s", g.ItemsKey)
	item := g.Item
	decoder := json.NewDecoder(bytes.NewReader(data))
	var groups []map[string]json.RawMessage
	if err := decoder.Decode(&groups); err != nil {
		var syntax *json.SyntaxError
		if errors.As(err, &syntax) {
			return nil, g.invalid(fmt.Sprintf("It is not valid JSON near byte %d", syntax.Offset))
		}
		return nil, g.invalid(shape)
	}
	var extra json.RawMessage
	if err := decoder.Decode(&extra); err != io.EOF {
		return nil, g.invalid("It holds more than 1 JSON value")
	}
	found := map[string]string{}
	for index, group := range groups {
		number := index + 1
		var reason string
		var items []string
		for name, value := range group {
			var err error
			switch name {
			case g.ItemsKey:
				err = json.Unmarshal(value, &items)
			case "reason":
				err = json.Unmarshal(value, &reason)
			default:
				return nil, g.invalid(shape)
			}
			if err != nil {
				return nil, g.invalid(shape)
			}
		}
		if strings.TrimSpace(reason) == "" {
			return nil, g.invalid(fmt.Sprintf("Group %d needs a non-empty reason", number))
		}
		if len(items) == 0 {
			return nil, g.invalid(fmt.Sprintf("Group %d needs at least 1 %s", number, item))
		}
		for _, text := range items {
			key, ok := g.Parse(text)
			if !ok {
				return nil, g.invalid(fmt.Sprintf("Group %d lists %q, which is not %s", number, text, g.ItemShape))
			}
			if earlier, duplicate := found[key]; duplicate {
				return nil, g.invalid(fmt.Sprintf("It lists %s and %s, which name the same %s. List it once", earlier, text, item))
			}
			found[key] = text
		}
	}
	return found, nil
}

func Run(check func(fs.FS, option.RequestOption, string, *http.Client) ([]string, string, error), source string, root fs.FS, baseURL option.RequestOption, client *http.Client, out, failures io.Writer) int {
	lines, summary, err := check(root, baseURL, Source(source), client)
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
