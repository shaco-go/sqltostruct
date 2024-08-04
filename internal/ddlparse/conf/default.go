package conf

// 默认的映射关系
func defMapping() []Mapping {
	var m = make([]Mapping, 0, 46)

	// 整数类型
	m = append(m, Mapping{Name: "tinyint", Type: "tinyint", Value: "int8"})
	m = append(m, Mapping{Name: "smallint", Type: "smallint", Value: "int16"})
	m = append(m, Mapping{Name: "mediumint", Type: "mediumint", Value: "int32"})
	m = append(m, Mapping{Name: "int", Type: "int", Value: "int32"})
	m = append(m, Mapping{Name: "bigint", Type: "bigint", Value: "int64"})

	// 引用类型的整数
	m = append(m, Mapping{Name: "tinyint:null", Type: "tinyint", Value: "*int8", Null: true})
	m = append(m, Mapping{Name: "smallint:null", Type: "smallint", Value: "*int16", Null: true})
	m = append(m, Mapping{Name: "mediumint:null", Type: "mediumint", Value: "*int32", Null: true})
	m = append(m, Mapping{Name: "int:null", Type: "int", Value: "*int32", Null: true})
	m = append(m, Mapping{Name: "bigint:null", Type: "bigint", Value: "*int64", Null: true})

	// 无符号整数类型
	m = append(m, Mapping{Name: "tinyint:unsigned", Type: "tinyint", Value: "uint8", Unsigned: true})
	m = append(m, Mapping{Name: "smallint:unsigned", Type: "smallint", Value: "uint16", Unsigned: true})
	m = append(m, Mapping{Name: "mediumint:unsigned", Type: "mediumint", Value: "uint32", Unsigned: true})
	m = append(m, Mapping{Name: "int:unsigned", Type: "int", Value: "uint32", Unsigned: true})
	m = append(m, Mapping{Name: "bigint:unsigned", Type: "bigint", Value: "uint64", Unsigned: true})

	// 引用类型的无符号整数
	m = append(m, Mapping{Name: "tinyint:null,unsigned", Type: "tinyint", Value: "*uint8", Null: true, Unsigned: true})
	m = append(m, Mapping{Name: "smallint:null,unsigned", Type: "smallint", Value: "*uint16", Null: true, Unsigned: true})
	m = append(m, Mapping{Name: "mediumint:null,unsigned", Type: "mediumint", Value: "*uint32", Null: true, Unsigned: true})
	m = append(m, Mapping{Name: "int:null,unsigned", Type: "int", Value: "*uint32", Null: true, Unsigned: true})
	m = append(m, Mapping{Name: "bigint:null,unsigned", Type: "bigint", Value: "*uint64", Null: true, Unsigned: true})

	// 浮点类型
	m = append(m, Mapping{Name: "float", Type: "float", Value: "float32"})
	m = append(m, Mapping{Name: "double", Type: "double", Value: "float64"})

	// 引用类型的浮点数
	m = append(m, Mapping{Name: "float:null", Type: "float", Value: "*float32", Null: true})
	m = append(m, Mapping{Name: "double:null", Type: "double", Value: "*float64", Null: true})

	// 定点类型
	m = append(m, Mapping{Name: "decimal", Type: "decimal", Value: "float64"})

	// 引用类型的定点数
	m = append(m, Mapping{Name: "decimal:null", Type: "decimal", Value: "*float64", Null: true})

	// 日期和时间类型
	m = append(m, Mapping{Name: "date", Type: "date", Value: "string"})
	m = append(m, Mapping{Name: "datetime", Type: "datetime", Value: "string"})
	m = append(m, Mapping{Name: "timestamp", Type: "timestamp", Value: "string"})
	m = append(m, Mapping{Name: "time", Type: "time", Value: "string"})
	m = append(m, Mapping{Name: "year", Type: "year", Value: "string"})

	// 引用日期和时间类型
	m = append(m, Mapping{Name: "date:null", Type: "date", Value: "string", Null: true})
	m = append(m, Mapping{Name: "datetime:null", Type: "datetime", Value: "string", Null: true})
	m = append(m, Mapping{Name: "timestamp:null", Type: "timestamp", Value: "string", Null: true})
	m = append(m, Mapping{Name: "time:null", Type: "time", Value: "string", Null: true})
	m = append(m, Mapping{Name: "year:null", Type: "year", Value: "string", Null: true})

	// 字符串类型
	m = append(m, Mapping{Name: "char", Type: "char", Value: "string"})
	m = append(m, Mapping{Name: "varchar", Type: "varchar", Value: "string"})

	// 文本类型
	m = append(m, Mapping{Name: "tinytext", Type: "tinytext", Value: "string"})
	m = append(m, Mapping{Name: "text", Type: "text", Value: "string"})
	m = append(m, Mapping{Name: "mediumtext", Type: "mediumtext", Value: "string"})
	m = append(m, Mapping{Name: "longtext", Type: "longtext", Value: "string"})

	// 二进制类型
	m = append(m, Mapping{Name: "binary", Type: "binary", Value: "[]byte"})
	m = append(m, Mapping{Name: "varbinary", Type: "varbinary", Value: "[]byte"})

	// 文本型二进制数据
	m = append(m, Mapping{Name: "tinyblob", Type: "tinyblob", Value: "[]byte"})
	m = append(m, Mapping{Name: "blob", Type: "blob", Value: "[]byte"})
	m = append(m, Mapping{Name: "mediumblob", Type: "mediumblob", Value: "[]byte"})
	m = append(m, Mapping{Name: "longblob", Type: "longblob", Value: "[]byte"})

	// 枚举和集合类型
	m = append(m, Mapping{Name: "enum", Type: "enum", Value: "string"})
	m = append(m, Mapping{Name: "set", Type: "set", Value: "string"})

	// JSON 类型
	m = append(m, Mapping{Name: "json", Type: "json", Value: "string"})

	return m
}

// 默认tags
func defTags() []Tag {
	return []Tag{
		{Name: "form", Enable: true, IsDefault: true},
		{Name: "json", Enable: true, IsDefault: true},
		{Name: "gorm", Enable: true, IsDefault: true},
	}
}

// NewDefConf 默认配置
func NewDefConf() *Conf {
	return &Conf{
		NamingStyle: SnakeCase,
		Tags:        defTags(),
		Mapping:     defMapping(),
		Comment:     true,
	}
}
