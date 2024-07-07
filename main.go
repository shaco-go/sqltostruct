package main

import (
	"fmt"
	"github.com/shaco-go/sqltostruct/internal/ddl"
)

func main() {
	s := ""
	table, err := ddl.Parse(s)
	fmt.Println(table, err)
}
