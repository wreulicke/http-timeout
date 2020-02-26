package linter

import (
	"go/ast"
	"go/types"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var Analyzer = &analysis.Analyzer{
	Name: "no-timeout-http-server",
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	nodeFilter := []ast.Node{
		(*ast.CompositeLit)(nil),
	}
	inspect.Preorder(nodeFilter, func(n ast.Node) {
		switch n := n.(type) {
		case *ast.CompositeLit:
			s, ok := n.Type.(*ast.SelectorExpr)
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
			if s.Sel.Name != "Server" {
				return
			}
			var foundReadTimeout bool
			var foundWriteTimeout bool
			for _, e := range n.Elts {
				x := e.(*ast.KeyValueExpr)
				switch x.Key.(*ast.Ident).Name {
				case "ReadTimeout":
					foundReadTimeout = true
				case "WriteTimeout":
					foundWriteTimeout = true
				}
			}
			if !foundReadTimeout {
				pass.Reportf(n.Pos(), "Do not use net/http.Server with no ReadTimeout")
			}
			if !foundWriteTimeout {
				pass.Reportf(n.Pos(), "Do not use net/http.Server with no WriteTimeout")
			}
		}

	})
	return nil, nil
}
