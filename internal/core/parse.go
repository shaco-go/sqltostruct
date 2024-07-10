package core

import (
	"fmt"
	"github.com/pingcap/tidb/pkg/parser"
	_ "github.com/pingcap/tidb/pkg/parser/test_driver"
	"github.com/pkg/errors"
)

func Parse(sql string) (*Table, error) {
	p := parser.New()
	stmt, _, err := p.ParseSQL(sql)
	if err != nil {
		return nil, err
	}
	if len(stmt) == 0 {
		return nil, errors.New("sql parse fail")
	}
	dom := stmt[0]
	var table Table
	node, ok := dom.Accept(&table)
	fmt.Println(node)
	if !ok {
		return nil, errors.New("sql parse fail")
	}
	return &table, nil
}
