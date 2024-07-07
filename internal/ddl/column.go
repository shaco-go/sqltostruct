package ddl

import (
	"github.com/pingcap/tidb/pkg/parser/ast"
	"github.com/pingcap/tidb/pkg/parser/format"
	"github.com/pingcap/tidb/pkg/parser/types"
	"strings"
)

type Column struct {
	Type    string
	Null    bool
	Default string
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

	var sb strings.Builder
	ctx := format.NewRestoreCtx(format.DefaultRestoreFlags, &sb)
	err := def.Restore(ctx)
	if err == nil {
		c.Sql = sb.String()
	}
}
