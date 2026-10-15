// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"fmt"
	"go/ast"
	"go/token"
	"go/types"
	"io/fs"
	"reflect"
	"slices"
	"strconv"
	"strings"

	"github.com/Xquik-dev/x-twitter-scraper-go/internal/apicheck"
)

const (
	rootPackage   = "main"
	sharedPackage = "shared"
	constPackage  = "constant"
)

var (
	packageDirs  = map[string]string{rootPackage: ".", sharedPackage: "shared", constPackage: "shared/constant"}
	scalarTypes  = []string{"string", "bool", "byte", "rune", "int", "int8", "int16", "int32", "int64", "uint", "uint8", "uint16", "uint32", "uint64", "float32", "float64"}
	scalarImport = []string{"time", "io"}
	paramMarkers = []string{"APIObject", "APIUnion"}
)

type goPackage struct {
	types  map[string]ast.Expr
	values map[string][]string
}

type sdkReader struct {
	packages map[string]*goPackage
	problems map[string]bool
}

func readSDK(root fs.FS) (sdkReader, error) {
	reader := sdkReader{packages: map[string]*goPackage{}, problems: map[string]bool{}}
	files := token.NewFileSet()
	for name, dir := range packageDirs {
		parsed, err := apicheck.Fields.ParseGoFiles(root, dir, files)
		if err != nil {
			return reader, err
		}
		current := &goPackage{types: map[string]ast.Expr{}, values: map[string][]string{}}
		for _, file := range parsed {
			current.collect(file)
		}
		reader.packages[name] = current
	}
	return reader, nil
}

func literal(expression ast.Expr) (string, bool) {
	value, ok := expression.(*ast.BasicLit)
	if !ok {
		return "", false
	}
	if text, err := strconv.Unquote(value.Value); err == nil {
		return text, true
	}
	return value.Value, true
}

func (p *goPackage) collect(file *ast.File) {
	for _, declaration := range file.Decls {
		switch typed := declaration.(type) {
		case *ast.GenDecl:
			for _, spec := range typed.Specs {
				switch value := spec.(type) {
				case *ast.TypeSpec:
					p.types[value.Name.Name] = value.Type
				case *ast.ValueSpec:
					p.constant(typed.Tok, value)
				}
			}
		case *ast.FuncDecl:
			p.defaultValue(typed)
		}
	}
}

func (p *goPackage) constant(tok token.Token, spec *ast.ValueSpec) {
	name, ok := spec.Type.(*ast.Ident)
	if tok != token.CONST || !ok {
		return
	}
	for _, expression := range spec.Values {
		if value, ok := literal(expression); ok {
			p.values[name.Name] = append(p.values[name.Name], value)
		}
	}
}

func (p *goPackage) defaultValue(function *ast.FuncDecl) {
	if function.Name.Name != "Default" || function.Recv == nil || function.Body == nil || len(function.Body.List) != 1 {
		return
	}
	statement, ok := function.Body.List[0].(*ast.ReturnStmt)
	receiver, named := function.Recv.List[0].Type.(*ast.Ident)
	if !ok || !named || len(statement.Results) != 1 {
		return
	}
	if value, ok := literal(statement.Results[0]); ok {
		p.values[receiver.Name] = []string{value}
	}
}

func (r sdkReader) unresolved(name string) *shape {
	r.problems[fmt.Sprintf("The field check cannot resolve the Go type %s. Use a type from the SDK, shared, shared/constant, time or io packages.", name)] = true
	return opaque(unknown)
}

func (r sdkReader) shape(pkg string, expression ast.Expr, tag string, stack []string) *shape {
	switch typed := expression.(type) {
	case *ast.Ident:
		switch {
		case slices.Contains(scalarTypes, typed.Name):
			return scalarOf(nil)
		case typed.Name == "any":
			return opaque(unknown)
		}
		return r.named(pkg, typed.Name, tag, stack)
	case *ast.SelectorExpr:
		owner := types.ExprString(typed.X)
		switch {
		case slices.Contains(scalarImport, owner):
			return scalarOf(nil)
		case owner == sharedPackage || owner == constPackage:
			return r.named(owner, typed.Sel.Name, tag, stack)
		case owner == "param" && slices.Contains(paramMarkers, typed.Sel.Name):
			return objectOf(map[string]field{})
		}
	case *ast.StarExpr:
		return r.shape(pkg, typed.X, tag, stack)
	case *ast.ArrayType:
		return &shape{kind: array, items: r.shape(pkg, typed.Elt, tag, stack)}
	case *ast.MapType:
		return opaque(free)
	case *ast.InterfaceType:
		return opaque(unknown)
	case *ast.IndexExpr:
		if selector, ok := typed.X.(*ast.SelectorExpr); ok && isPackage(selector.X, "param") && selector.Sel.Name == "Opt" {
			return r.shape(pkg, typed.Index, tag, stack)
		}
	case *ast.StructType:
		return r.structure(pkg, typed, tag, false, stack)
	}
	return r.unresolved(types.ExprString(expression))
}

func (r sdkReader) named(pkg, name, tag string, stack []string) *shape {
	current := r.packages[pkg]
	if values, ok := current.values[name]; ok {
		return scalarOf(values)
	}
	expression, ok := current.types[name]
	if !ok && pkg == rootPackage {
		return r.unresolved(name)
	}
	if !ok {
		return r.unresolved(pkg + "." + name)
	}
	structure, ok := expression.(*ast.StructType)
	if !ok {
		return r.shape(pkg, expression, tag, stack)
	}
	key := pkg + "." + name
	if slices.Contains(stack, key) {
		return opaque(cycle)
	}
	return r.structure(pkg, structure, tag, strings.HasSuffix(name, "Union"), append(slices.Clone(stack), key))
}

func (r sdkReader) structure(pkg string, structure *ast.StructType, tag string, partial bool, stack []string) *shape {
	result := objectOf(map[string]field{})
	var variants []*shape
	for _, member := range structure.Fields.List {
		var tags reflect.StructTag
		if member.Tag != nil {
			if text, err := strconv.Unquote(member.Tag.Value); err == nil {
				tags = reflect.StructTag(text)
			}
		}
		wire, tagged := tags.Lookup(tag)
		options := strings.Split(wire, ",")
		name := options[0]
		anonymous := len(member.Names) == 0
		if anonymous && (!tagged || name == "") {
			embedded := r.shape(pkg, member.Type, tag, stack)
			switch embedded.kind {
			case object:
				for key, value := range embedded.fields {
					if _, taken := result.fields[key]; !taken {
						result.fields[key] = value
					}
				}
			case union:
				variants = append(variants, embedded)
			}
			continue
		}
		if !anonymous && slices.Contains(options[1:], "inline") && strings.HasPrefix(member.Names[0].Name, "Of") {
			variants = append(variants, r.shape(pkg, member.Type, tag, stack))
			continue
		}
		if !tagged || name == "" || name == "-" {
			continue
		}
		api := strings.Split(tags.Get("api"), ",")
		required := flagOf(slices.Contains(api, "required"))
		if selector, ok := member.Type.(*ast.SelectorExpr); partial || (ok && isPackage(selector.X, constPackage)) {
			required = unencoded
		}
		nullable := unencoded
		if slices.Contains(api, "nullable") {
			nullable = on
		}
		result.fields[name] = field{required: required, nullable: nullable, shape: r.shape(pkg, member.Type, tag, stack)}
	}
	if len(variants) > 0 && len(result.fields) == 0 {
		return unionOf(variants)
	}
	return result
}

func isPackage(expression ast.Expr, name string) bool {
	ident, ok := expression.(*ast.Ident)
	return ok && ident.Name == name
}

func (r sdkReader) operation(operation apicheck.SDKOperation) (operationShapes, bool) {
	shapes := operationShapes{query: objectOf(map[string]field{})}
	function := operation.Function
	if results := function.Type.Results; results != nil && len(results.List) > 1 && types.ExprString(results.List[0].Type) != "*http.Response" {
		shapes.response = r.shape(rootPackage, results.List[0].Type, "json", nil)
	}
	if len(operation.Call.Args) < 4 || isPackage(operation.Call.Args[3], "nil") {
		return shapes, true
	}
	argument, ok := operation.Call.Args[3].(*ast.Ident)
	if !ok {
		return shapes, false
	}
	for _, parameter := range function.Type.Params.List {
		for _, name := range parameter.Names {
			if name.Name != argument.Name {
				continue
			}
			shapes.query = r.shape(rootPackage, parameter.Type, "query", nil)
			if body := r.shape(rootPackage, parameter.Type, "json", nil); body.kind != object || len(body.fields) > 0 {
				shapes.body = body
			}
			return shapes, true
		}
	}
	return shapes, false
}
