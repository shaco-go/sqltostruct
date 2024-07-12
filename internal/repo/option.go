package repo

import (
	"encoding/json"
	"fmt"
	"github.com/shaco-go/sqltostruct/global"
	"github.com/shaco-go/sqltostruct/internal/model"
	"github.com/shaco-go/sqltostruct/internal/request"
	"gorm.io/datatypes"
	"strings"
)

type Option struct {
}

func NewOption() *Option {
	return &Option{}
}

// Get 获取配置
func (o Option) Get() *model.Option {
	var option *model.Option
	_ = global.DB.Where("id = 1").Find(&option).Error
	if option.ID == 0 {
		// 初始化一下数据
		option = o.InitData()
		if err := global.DB.Create(option).Error; err != nil {
			fmt.Println("db create err:", err)
		}
	}
	return option
}

func (o Option) Save(req request.OptionSaveReq) error {
	var tag = make(datatypes.JSONMap)
	_ = tag.UnmarshalJSON([]byte(req.Tag))
	data := o.Get()
	data.IsNull = req.IsNull
	data.IsDefault = req.IsDefault
	data.IsUnsigned = req.IsUnsigned
	data.IsLen = req.IsLen
	data.IsComment = req.IsComment
	data.Style = req.Style
	data.IsGorm = req.IsGorm
	data.ColumnMap = o.toUpper(req.ColumnMap)
	data.ColumnNullMap = o.toUpper(req.ColumnNullMap)
	data.Tag = tag
	return global.DB.
		Select("*").
		Updates(data).Error
}

func (o Option) toUpper(data string) datatypes.JSONMap {
	var m map[string]any
	var newM = make(map[string]any, len(m))
	_ = json.Unmarshal([]byte(data), &m)
	for k, v := range m {
		newM[strings.ToUpper(k)] = v
	}
	return newM
}

// InitData 初始化数据
func (o Option) InitData() *model.Option {
	return &model.Option{
		ID:         1,
		IsNull:     false,
		IsDefault:  false,
		IsUnsigned: false,
		IsLen:      true,
		IsGorm:     true,
		IsComment:  true,
		Style:      model.SnakeCase,
		ColumnMap: datatypes.JSONMap{
			"BIT":        "[]byte",
			"TINYINT":    "int8",
			"BOOL":       "bool",
			"BOOLEAN":    "bool",
			"SMALLINT":   "int16",
			"MEDIUMINT":  "int32",
			"INT":        "int32",
			"INTEGER":    "int32",
			"BIGINT":     "int64",
			"FLOAT":      "float32",
			"DOUBLE":     "float64",
			"DECIMAL":    "float64",
			"DATE":       "time.Time",
			"DATETIME":   "time.Time",
			"TIMESTAMP":  "time.Time",
			"TIME":       "time.Time",
			"YEAR":       "int16",
			"CHAR":       "string",
			"VARCHAR":    "string",
			"BINARY":     "[]byte",
			"VARBINARY":  "[]byte",
			"TINYBLOB":   "[]byte",
			"BLOB":       "[]byte",
			"MEDIUMBLOB": "[]byte",
			"LONGBLOB":   "[]byte",
			"TINYTEXT":   "string",
			"TEXT":       "string",
			"MEDIUMTEXT": "string",
			"LONGTEXT":   "string",
			"ENUM":       "string",
			"SET":        "string",
		},
		ColumnNullMap: datatypes.JSONMap{
			"BIT":        "[]byte",
			"TINYINT":    "*int8",
			"BOOL":       "*bool",
			"BOOLEAN":    "*bool",
			"SMALLINT":   "*int16",
			"MEDIUMINT":  "*int32",
			"INT":        "*int32",
			"INTEGER":    "*int32",
			"BIGINT":     "*int64",
			"FLOAT":      "*float32",
			"DOUBLE":     "*float64",
			"DECIMAL":    "*float64",
			"DATE":       "*time.Time",
			"DATETIME":   "*time.Time",
			"TIMESTAMP":  "*time.Time",
			"TIME":       "*time.Time",
			"YEAR":       "*int16",
			"CHAR":       "*string",
			"VARCHAR":    "*string",
			"BINARY":     "[]byte",
			"VARBINARY":  "[]byte",
			"TINYBLOB":   "[]byte",
			"BLOB":       "[]byte",
			"MEDIUMBLOB": "[]byte",
			"LONGBLOB":   "[]byte",
			"TINYTEXT":   "*string",
			"TEXT":       "*string",
			"MEDIUMTEXT": "*string",
			"LONGTEXT":   "*string",
			"ENUM":       "*string",
			"SET":        "*string",
		},
		Tag: datatypes.JSONMap{
			"json": true,
			"form": true,
		},
	}
}
