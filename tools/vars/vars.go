package vars

import (
	"go/ast"
	"go/parser"
	"go/token"
	"strconv"
	"strings"
)

type tmpVisitor struct {
	row, col int
	fs       *token.FileSet
	vars     []ast.Node
}

func (v *tmpVisitor) Visit(n ast.Node) ast.Visitor {
	if n == nil {
		return v
	}
	position := v.fs.Position(n.Pos())
	nLen := n.End() - n.Pos()
	if position.Line == v.row && position.Column+int(nLen)+1 >= v.col {
		if exprStmt, ok := n.(*ast.ExprStmt); ok {
			if selExpr, sOk := exprStmt.X.(*ast.SelectorExpr); sOk && selExpr.Sel.Name == "_" {
				v.vars = append(v.vars, selExpr.X)
				ast.Print(v.fs, n)
			}
		}
	}
	return v
}

func ReplaceVarWithDef(source interface{}) (string, error) {
	fs := token.NewFileSet()
	f, err := parser.ParseFile(fs, "", source, parser.ParseComments)
	if err != nil && !strings.Contains(err.Error(), "var") {
		return "", err
	}

	splits := strings.Split(err.Error(), ":")
	rowStr, colStr := splits[0], splits[1]
	row, _ := strconv.Atoi(rowStr)
	col, _ := strconv.Atoi(colStr)

	// find the selectorExpr
	tmpV := tmpVisitor{row: row, col: col, fs: fs}
	ast.Walk(&tmpV, f)

	// query result doc
	var prefix string
	for _, v := range tmpV.vars {
		switch val := v.(type) {
		case *ast.CallExpr:
			prefix, err = dealCallExpr(val)
			break
		}
	}

	// insert prefix before line
	afterInsert := insertPrefix(prefix, source)
	return afterInsert, nil
}

func insertPrefix(prefix string, source interface{}) string {
	return ""
}
