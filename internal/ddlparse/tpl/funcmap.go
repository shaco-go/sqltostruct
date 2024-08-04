package tpl

import (
	"fmt"
	"github.com/samber/lo"
	"github.com/shaco-go/sqltostruct/internal/ddlparse/conf"
	"github.com/shaco-go/sqltostruct/internal/ddlparse/ddl"
	"github.com/shaco-go/sqltostruct/internal/ddlparse/utils"
	"html/template"
	"strings"
)

func NewFuncMap() template.FuncMap {
	return template.FuncMap{
		"pascalCase":     lo.PascalCase,
		"getFirstLetter": getFirstLetter,
		"getType":        getType,
		"getTag":         getTag,
		"getComment":     getComment,
	}
}

// 获取第一个字母
func getFirstLetter(str string) string {
	bytes := []byte(str)
	return strings.ToLower(string(bytes[0]))
}

func getType(col ddl.Column, config *conf.Conf) string {
	if val, ok := config.MappingRela[conf.ToType(col.Type, col.Null, col.Unsigned)]; ok {
		return val

	}
	if val, ok := config.MappingRela[col.Type]; ok {
		return val
	}
	return "any"
}

func getTag(col ddl.Column, config *conf.Conf) string {
	tags := lo.Filter(config.Tags, func(item conf.Tag, index int) bool {
		return item.Enable
	})
	var tagTpl []string
	if len(tags) > 0 {
		for _, tag := range tags {
			switch tag.Name {
			case "gorm":
				tagTpl = append(tagTpl, fmt.Sprintf(`gorm:"column:%s"`, utils.ToNamingStyle(col.Field, config.NamingStyle)))
			default:
				tagTpl = append(tagTpl, fmt.Sprintf(`%s:"%s"`, tag.Name, utils.ToNamingStyle(col.Field, config.NamingStyle)))
			}
		}
	}
	if len(tagTpl) > 0 {
		return fmt.Sprintf("`%s`", strings.Join(tagTpl, " "))
	}
	return ""
}

func getComment(col ddl.Column, config *conf.Conf) string {
	if config.Comment {
		return "// " + col.Comment
	}
	return ""
}
