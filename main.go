package main

import (
	"fmt"
	"github.com/shaco-go/sqltostruct/internal/ddl"
	"os"
)

func main() {
	s := "create table user\n(\n    id         int(11) auto_increment primary key comment '自增id',\n    name       varchar(20)    default '' not null comment '姓名',\n    age        tinyint        default 6666  not null comment '年龄',\n    price      decimal(10, 2) default 0  not null comment '价格',\n    created_at date                      null comment '创建时间'\n)"
	table, err := ddl.Parse(s)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(os.Getwd())
	table.Template("./internal/ddl/template")
}
