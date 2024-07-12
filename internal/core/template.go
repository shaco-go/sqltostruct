package core

import (
	_ "embed"
	"fmt"
	"github.com/pkg/errors"
	"github.com/samber/lo"
	"github.com/shaco-go/sqltostruct/internal/model"
	"github.com/shaco-go/sqltostruct/internal/repo"
	"github.com/spf13/cast"
	"go/format"
	"sort"
	"strings"
	"text/template"
)

var (
	optionRepo = repo.NewOption()
)

//go:embed template
var templateStr string

type Template struct {
	Table   *Table
	Option  *model.Option
	Columns []string
}

func NewTemplate(sql string) (*Template, error) {
	table, err := Parse(sql)
	if err != nil {
		return nil, errors.New("解析SQL失败")
	}
	option := optionRepo.Get()
	var t = &Template{
		Table:  table,
		Option: option,
	}
	t.build()
	return t, nil
}

func (t *Template) getColumnByStyle(column string) string {
	switch t.Option.Style {
	case model.PascalCase:
		return lo.PascalCase(column)
	case model.CamelCase:
		return lo.CamelCase(column)
	case model.KebabCase:
		return lo.KebabCase(column)
	case model.SnakeCase:
		return lo.SnakeCase(column)
	}
	return column
}

func (t *Template) build() {
	// tag顺序
	var tagField []string
	for field, ok := range t.Option.Tag {
		if cast.ToBool(ok) {
			tagField = append(tagField, field)
		}
	}
	sort.Slice(tagField, func(i, j int) bool {
		return i < j
	})
	for _, item := range t.Table.Columns {
		var text = lo.PascalCase(item.Name) + " "
		// 类型
		var typeStr = "any"
		if t.Option.IsNull && item.Null {
			if val, ok := t.Option.ColumnNullMap[strings.ToUpper(item.Type)]; ok {
				typeStr = cast.ToString(val)
			}
		} else {
			if val, ok := t.Option.ColumnMap[strings.ToUpper(item.Type)]; ok {
				typeStr = cast.ToString(val)
			}
		}
		text += typeStr
		var tag []string
		if len(tagField) > 0 {
			for _, field := range tagField {
				tag = append(tag, fmt.Sprintf(`%s:"%s"`, field, t.getColumnByStyle(item.Name)))
			}
		}
		if t.Option.IsGorm {
			tag = append(tag, fmt.Sprintf(`gorm:"column:%s"`, t.getColumnByStyle(item.Name)))
		}
		if len(tag) > 0 {
			text += fmt.Sprintf(" `%s` ", strings.Join(tag, " "))
		}
		if t.Option.IsComment {
			text += fmt.Sprintf(" // %s", t.getColumnByStyle(item.Comment))
		}
		t.Columns = append(t.Columns, text)
	}
}

func (t *Template) Generate() (string, error) {
	files, err := template.New("template").Funcs(map[string]any{
		"pascalCase": lo.PascalCase,
		"getWord": func(word string) string {
			var a = []rune(word)
			return string(a[0])
		},
	}).Parse(templateStr)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	var builder strings.Builder
	err = files.Execute(&builder, t)
	source, _ := format.Source([]byte(builder.String()))
	return string(source), nil
}
