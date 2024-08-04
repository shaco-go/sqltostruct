package utils

import (
	"github.com/samber/lo"
	"github.com/shaco-go/sqltostruct/internal/ddlparse/conf"
)

// ToNamingStyle 转命名风格
func ToNamingStyle(s string, t conf.NamingStyle) string {
	switch t {
	case conf.PascalCase:
		return lo.PascalCase(s)
	case conf.CamelCase:
		return lo.CamelCase(s)
	case conf.KebabCase:
		return lo.KebabCase(s)
	case conf.SnakeCase:
		return lo.SnakeCase(s)
	default:
		return s
	}
}
