package ddl

import (
	"fmt"
	"github.com/pingcap/tidb/pkg/parser/ast"
	"github.com/samber/lo"
	"os"
	"path/filepath"
	"text/template"
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

func (t *Table) Template(filePath string) {
	files, err := template.New(filepath.Base(filePath)).Funcs(map[string]any{
		"pascalCase": lo.PascalCase,
	}).ParseFiles(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = files.Execute(os.Stdout, t)
	if err != nil {
		fmt.Println(err)
		return
	}
}
