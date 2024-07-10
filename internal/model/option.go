package model

import "gorm.io/datatypes"

const (
	PascalCase = iota
	CamelCase
	KebabCase
	SnakeCase
)

type Option struct {
	ID            int64             `json:"id" gorm:"column:id;primary_key;AUTO_INCREMENT"`
	IsNull        bool              `json:"is_null" gorm:"column:is_null;"`
	IsDefault     bool              `json:"is_default" gorm:"column:is_default;"`
	IsUnsigned    bool              `json:"is_unsigned" gorm:"column:is_unsigned;"`
	IsLen         bool              `json:"is_len" gorm:"column:is_len;"`
	IsComment     bool              `json:"is_comment" gorm:"column:is_comment;"`
	IsGorm        bool              `json:"is_gorm" gorm:"column:is_gorm;"`
	Style         int               `json:"style" gorm:"column:style;"`
	ColumnMap     datatypes.JSONMap `json:"column_map" gorm:"column:column_map;"`
	ColumnNullMap datatypes.JSONMap `json:"column_null_map" gorm:"column:column_null_map;"`
	Tag           datatypes.JSONMap `json:"tag" gorm:"column:tag;"`
}

func (o *Option) TableName() string {
	return "option"
}
