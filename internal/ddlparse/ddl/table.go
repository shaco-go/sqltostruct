package ddl

import (
	"github.com/pingcap/tidb/pkg/parser"
	"github.com/pingcap/tidb/pkg/parser/ast"
	"github.com/pingcap/tidb/pkg/parser/format"
	"github.com/pingcap/tidb/pkg/parser/test_driver"
	"github.com/pingcap/tidb/pkg/parser/types"
	"github.com/shaco-go/sqltostruct/internal/ddlparse/errors"
	"strings"
)

type Table struct {
	TableName string
	Columns   []Column
	Comment   string
}

func NewTable() *Table {
	return &Table{}
}

type Column struct {
	Field    string
	Type     string
	Null     bool
	Unsigned bool
	Comment  string
	FullSql  string
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
		t.parseColumn(val)
	}
	return in, false
}

func (t *Table) Leave(in ast.Node) (ast.Node, bool) {
	return in, true
}

// 解析列
func (t *Table) parseColumn(dom *ast.ColumnDef) {
	c := Column{
		Field: dom.Name.String(),
		Type:  types.TypeStr(dom.Tp.GetType()),
	}
	for _, item := range dom.Options {
		switch item.Tp {
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
	err := dom.Restore(ctx)
	if err == nil {
		c.FullSql = sb.String()
	}
	t.Columns = append(t.Columns, c)
}

// Parse 解析table
func (t *Table) Parse(ddl string) error {
	p := parser.New()
	stmt, _, err := p.ParseSQL(ddl)
	if err != nil {
		return err
	}
	if len(stmt) == 0 {
		return errors.DDLParseFail
	}
	dom := stmt[0]
	_, ok := dom.Accept(t)
	if !ok {
		return errors.DDLParseFail
	}
	return nil
}
