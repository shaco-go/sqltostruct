package core

import (
	"github.com/pingcap/tidb/pkg/parser/ast"
)

type Table struct {
	TableName string
	Comment   string
	Columns   []*Column
}

func (t *Table) Enter(in ast.Node) (ast.Node, bool) {
	// 表名
	if val, ok := in.(*ast.TableName); ok {
		t.TableName = val.Name.String()
	}
	// 表注释
	if val, ok := in.(*ast.TableOption); ok {
		if val.Tp == ast.TableOptionComment {
			t.Comment = val.StrValue
		}
	}
	// 解析列
	if val, ok := in.(*ast.ColumnDef); ok {
		t.Columns = append(t.Columns, NewColumn(val))
	}
	return in, false
}

func (t *Table) Leave(in ast.Node) (ast.Node, bool) {
	return in, true
}
