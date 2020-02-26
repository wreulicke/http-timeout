package linter

import (
	"go/ast"
	"go/types"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var Analyzer = &analysis.Analyzer{
	Name: "no-default-http-client",
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	nodeFilter := []ast.Node{
		(*ast.CallExpr)(nil),
		(*ast.SelectorExpr)(nil),
	}
	inspect.Preorder(nodeFilter, func(n ast.Node) {
		switch n := n.(type) {
		case *ast.SelectorExpr:
			pkgIdent, ok := n.X.(*ast.Ident)
			if !ok {
				return
			}
			pkgName, ok := pass.TypesInfo.Uses[pkgIdent].(*types.PkgName)
			if !ok || pkgName.Imported().Path() != "net/http" {
				return
			}
			if n.Sel.Name == "DefaultClient" {
				pass.Reportf(n.Pos(), "Do not use net/http.%s because default client has no timeout", n.Sel.Name)
			}
		case *ast.CallExpr:
			s, ok := n.Fun.(*ast.SelectorExpr)
			if !ok {
				return
			}
			pkgIdent, ok := s.X.(*ast.Ident)
			if !ok {
				return
			}
			pkgName, ok := pass.TypesInfo.Uses[pkgIdent].(*types.PkgName)
			if !ok || pkgName.Imported().Path() != "net/http" {
				return
			}
			switch s.Sel.Name {
			case "Get":
				pass.Reportf(n.Pos(), "Do not use net/http.%s because default client has no timeout", s.Sel.Name)
			case "Post":
				pass.Reportf(n.Pos(), "Do not use net/http.%s because default client has no timeout", s.Sel.Name)
			case "PostForm":
				pass.Reportf(n.Pos(), "Do not use net/http.%s because default client has no timeout", s.Sel.Name)
			}
		}

	})
	return nil, nil
}
