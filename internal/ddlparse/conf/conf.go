package conf

import (
	"strings"
)

type NamingStyle int

const (
	PascalCase NamingStyle = iota + 1
	CamelCase
	KebabCase
	SnakeCase
)

type Conf struct {
	NamingStyle NamingStyle `json:"naming_style"`
	Tags        []Tag       `json:"tags"`
	Mapping     []Mapping   `json:"mapping"`
	Comment     bool        `json:"comment"`
	MappingRela map[string]string
}

type Tag struct {
	Name      string `json:"name"`
	Enable    bool   `json:"enable"`
	IsDefault bool   `json:"is_default"`
}

type Mapping struct {
	Name     string `json:"name"`
	Type     string `json:"type"`
	Value    string `json:"value"`
	Null     bool   `json:"null"`
	Unsigned bool   `json:"unsigned"`
}

// Init 初始化
func (c *Conf) Init() {
	c.MappingRela = make(map[string]string, len(c.Mapping))
	for _, item := range c.Mapping {
		c.MappingRela[ToType(item.Type, item.Null, item.Unsigned)] = item.Value
	}
}

// ToType ddl转type
func ToType(t string, null, unsigned bool) string {
	s := t
	var tag []string
	if null {
		tag = append(tag, "null")
	}
	if unsigned {
		tag = append(tag, "unsigned")
	}
	if len(tag) > 0 {
		s += ":" + strings.Join(tag, ",")
	}
	return s
}
