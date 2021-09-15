package main

import (
	_ "exam/routers"
	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	_ = orm.RegisterDriver("mysql", orm.DRMySQL)
	dbUrl, _ := beego.AppConfig.String("datasource.url")
	_ = orm.RegisterDataBase("default", "mysql", dbUrl)
	orm.Debug = true

	beego.Run()
}
