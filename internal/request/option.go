package request

type OptionSaveReq struct {
	IsNull        bool   `json:"is_null" form:"is_null"`
	IsDefault     bool   `json:"is_default" form:"is_default"`
	IsUnsigned    bool   `json:"is_unsigned" form:"is_unsigned"`
	IsLen         bool   `json:"is_len" form:"is_len"`
	IsComment     bool   `json:"is_comment" form:"is_comment"`
	IsGorm        bool   `json:"is_gorm" form:"is_gorm"`
	Style         int    `json:"style" form:"style"`
	ColumnMap     string `json:"column_map" form:"column_map"`
	ColumnNullMap string `json:"column_null_map" form:"column_null_map"`
	Tag           string `json:"tag" form:"tag"`
}
