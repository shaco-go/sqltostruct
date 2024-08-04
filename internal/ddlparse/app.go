package ddlparse

import (
	_ "embed"
	"github.com/shaco-go/sqltostruct/internal/ddlparse/conf"
	"github.com/shaco-go/sqltostruct/internal/ddlparse/ddl"
	"github.com/shaco-go/sqltostruct/internal/ddlparse/tpl"
	"go/format"
	"strings"
	"text/template"
)

//go:embed tpl/model.tpl
var tplFile string

type App struct {
	table      *ddl.Table
	conf       *conf.Conf
	tplFuncMap template.FuncMap
}

func NewApp(conf *conf.Conf) *App {
	conf.Init()
	var app = &App{
		conf:       conf,
		table:      ddl.NewTable(),
		tplFuncMap: tpl.NewFuncMap(),
	}
	return app
}

// Build 生成模板
func (a *App) Build(s string) (string, error) {
	// 解析ddl
	err := a.table.Parse(s)
	if err != nil {
		return "", err
	}
	files, err := template.New("template").Funcs(a.tplFuncMap).Parse(tplFile)
	if err != nil {
		return "", err
	}
	var builder strings.Builder
	err = files.Execute(&builder, struct {
		Table *ddl.Table
		Conf  *conf.Conf
		test  string
	}{
		Table: a.table,
		Conf:  a.conf,
	})
	if err != nil {
		return "", err
	}
	source, err := format.Source([]byte(builder.String()))
	if err != nil {
		return "", err
	}
	return string(source), nil
	//return builder.String(), nil
}
