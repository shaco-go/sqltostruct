package main

import (
	"fmt"
	"github.com/shaco-go/sqltostruct/internal/ddl"
)

func main() {
	s := "create table Map\n(\n    id       int            null,\n    province varchar(255)   null,\n    lng      decimal(10, 6) null,\n    lat      decimal(10, 6) null\n);"
	table, err := ddl.Parse(s)
	fmt.Println(table, err)
}
