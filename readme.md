# Sqltostruct


## 优点
* 兼容非RDBMS 服务器提供的DDL
* 可自定义tag,配置对应的字段映射,包括null类型

## 待开发

- [ ] 支持json转struct
- [ ] 支持更多的配置
- [ ] 更新到本地文件,只更新未添加的字段
- [ ] 支持struct转sql,json

## 截图
![图片](https://i.imgur.com/N5u9w1C.png)

## 下载
[https://github.com/shaco-go/sqltostruct/releases](https://github.com/shaco-go/sqltostruct/releases)

## 注意
默认端口7788,如果端口被占用`-p`来制定端口
```shell
.\sqltostruct.exe -p 8888
```