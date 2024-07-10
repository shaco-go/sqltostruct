package main

import (
	"embed"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/shaco-go/sqltostruct/global"
	"github.com/shaco-go/sqltostruct/internal/core"
	"github.com/shaco-go/sqltostruct/internal/initialization"
	"github.com/shaco-go/sqltostruct/internal/middleare"
	"github.com/shaco-go/sqltostruct/internal/repo"
	"github.com/shaco-go/sqltostruct/internal/request"
	"github.com/shaco-go/sqltostruct/internal/response"
	"net/http"
)

var (
	optionRepo = repo.NewOption()
)

//go:embed ui/dist
var f embed.FS

func main() {
	// init db
	global.DB = initialization.NewGorm()

	// router
	e := gin.Default()
	e.Use(middleare.CORS())

	//st, _ := fs.Sub(f, "ui/dist")
	//e.StaticFS("/", http.FS(st))
	e.Use(middleare.Serve("/", middleare.EmbedFolder(f, "static/dist")))
	e.NoRoute(func(c *gin.Context) {
		data, err := f.ReadFile("ui/dist/index.html")
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		c.Data(http.StatusOK, "text/html; charset=utf-8", data)
	})

	e.GET("/config", func(c *gin.Context) {
		data := optionRepo.Get()
		response.OkWithData(data, c)
	})
	e.PATCH("/config", func(c *gin.Context) {
		var req request.OptionSaveReq
		if err := c.ShouldBind(&req); err != nil {
			response.FailWithMessage(err.Error(), c)
			return
		}
		if err := optionRepo.Save(req); err != nil {
			response.FailWithMessage(err.Error(), c)
			return
		}
		response.Ok(c)
	})
	e.POST("/generate", func(c *gin.Context) {
		sql := c.PostForm("sql")
		t, err := core.NewTemplate(sql)
		if err != nil {
			response.FailWithMessage(err.Error(), c)
			return
		}
		code, err := t.Generate()
		if err != nil {
			response.FailWithMessage(err.Error(), c)
			return
		}
		fmt.Println(code)
		response.OkWithData(gin.H{
			"code": code,
		}, c)
	})

	panic(e.Run(":7899"))
}
