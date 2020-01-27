package vars

import (
	"strconv"
	"strings"

	"github.com/rogpeppe/godef/go/ast"
	"github.com/rogpeppe/godef/go/parser"
	"github.com/rogpeppe/godef/go/token"
	"github.com/rogpeppe/godef/go/types"
)

type tmpVisitor struct {
	row, col int
	fs       *token.FileSet
	ty       *types.Type
}

func (v *tmpVisitor) Visit(n ast.Node) ast.Visitor {
	if n == nil || v.ty != nil {
		return nil
	}
	position := v.fs.Position(n.Pos())
	nLen := n.End() - n.Pos()
	if position.Line == v.row {
		if position.Column+int(nLen)+1 >= v.col {
			if exprStmt, ok := n.(*ast.ExprStmt); ok {
				if taExpr, sOk := exprStmt.X.(*ast.TypeAssertExpr); sOk {
					if _, bOk := taExpr.Type.(*ast.BadExpr); bOk {
						return v
					}
				}
			} else if _, ok = n.(*ast.TypeAssertExpr); ok {
				return v
			}

			pos := getPos(n)
			if pos == token.NoPos {
				return v
			}

			val, _ := n.(ast.Expr)
			_, typ := types.ExprType(val, types.DefaultImporter, v.fs)
			v.ty = &typ
			return nil
		}
	}
	return v
}

func getPos(n ast.Node) token.Pos {
	switch val := n.(type) {
	case *ast.CallExpr:
		return getPos(val.Fun)
	case *ast.SelectorExpr:
		return val.Sel.Pos()
	case *ast.Ident:
		return val.NamePos
	case *ast.BasicLit:
		return val.Pos()
	}

	return token.NoPos
}

func nodeNameGetter(n ast.Node) []string {
	name := []string{}
	switch val := n.(type) {
	case *ast.Ident:
		name = append(name, string(val.Name[0]))
	case *ast.ArrayType:
		name = append(name, nodeNameGetter(val)[0]+"s")
	case types.MultiValue:
		for _, v := range val.Types {
			name = append(name, nodeNameGetter(v)...)
		}
	default:
		name = append(name, "i")
	}
	return name
}

func ReplaceVarWithDef(fileName string, source interface{}) ([]string, error) {
	fs := token.NewFileSet()
	pkgScope := ast.NewScope(parser.Universe)
	f, err := parser.ParseFile(fs, fileName, source, 0, pkgScope, types.DefaultImportPathToName)
	if err != nil && !strings.Contains(err.Error(), "var") {
		return nil, err
	}

	splits := strings.Split(err.Error(), ":")
	rowLine := 0
	if fileName != "" {
		rowLine += 2
	}
	rowStr, colStr := splits[rowLine], splits[rowLine+1]
	row, _ := strconv.Atoi(rowStr)
	col, _ := strconv.Atoi(colStr)

	// find the selectorExpr
	tmpV := tmpVisitor{row: row, col: col, fs: fs}
	ast.Walk(&tmpV, f)

	// query result doc
	// var prefix string
	if tmpV.ty == nil {
		return nil, nil
	}

	names := nodeNameGetter(tmpV.ty.Node)
	return names, nil
}

func insertPrefix(prefix string, source interface{}) string {
	return ""
}
