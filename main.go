package main

import (
	_ "embed"
	"fmt"
	"github.com/shaco-go/sqltostruct/internal/ddlparse"
	"github.com/shaco-go/sqltostruct/internal/ddlparse/conf"
)

//go:embed model
var model string

func main() {
	config := conf.NewDefConf()
	app := ddlparse.NewApp(config)
	tmp, err := app.Build(model)
	if err != nil {
		panic(err)
	}
	fmt.Println(tmp)
}
