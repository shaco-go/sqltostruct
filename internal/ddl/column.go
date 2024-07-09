package ddl

import (
	"github.com/pingcap/tidb/pkg/parser/ast"
	"github.com/pingcap/tidb/pkg/parser/format"
	"github.com/pingcap/tidb/pkg/parser/test_driver"
	"github.com/pingcap/tidb/pkg/parser/types"
	"strings"
)

type Column struct {
	Type    string
	Null    bool
	Default any
	Name    string
	Comment string
	Len     int
	Sql     string
}

func NewColumn(dom *ast.ColumnDef) *Column {
	var column Column
	column.Parse(dom)
	return &column
}

func (c *Column) Parse(def *ast.ColumnDef) {
	c.Type = types.TypeStr(def.Tp.GetType())
	c.Name = def.Name.String()
	c.Len = def.Tp.GetFlen()
	for _, item := range def.Options {
		switch item.Tp {
		case ast.ColumnOptionDefaultValue: //默认值
			if val, ok := item.Expr.(*test_driver.ValueExpr); ok {
				c.Default = val.Datum.GetValue()
			}
		case ast.ColumnOptionNull:
			c.Null = true
		case ast.ColumnOptionComment:
			if val, ok := item.Expr.(*test_driver.ValueExpr); ok {
				c.Comment = val.Datum.GetString()
			}
		}
	}

	var sb strings.Builder
	ctx := format.NewRestoreCtx(format.DefaultRestoreFlags, &sb)
	err := def.Restore(ctx)
	if err == nil {
		c.Sql = sb.String()
	}
}
