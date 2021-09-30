package main

import (
	"exam/core/web"
	"exam/filter"
	"exam/routers"
	_ "exam/routers"
	"github.com/beego/beego/v2/adapter/logs"
	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

func main() {

	filter.Init()

	logs.SetLevel(logs.LevelDebug)
	_ = orm.RegisterDriver("mysql", orm.DRMySQL)
	dbUrl, _ := beego.AppConfig.String("datasource.url")
	_ = orm.RegisterDataBase("default", "mysql", dbUrl)
	orm.Debug = true

	r := gin.New()

	r.Use(gin.Logger())

	r.NoRoute(func(context *gin.Context) {
		web.ErrorWithStatusCode(context, http.StatusNotFound, "resource not found")
	})

	r.NoMethod(func(context *gin.Context) {
		web.ErrorWithStatusCode(context, http.StatusMethodNotAllowed, "method not allowed")
	})

	routers.Init(r)

	r.Run()
}
