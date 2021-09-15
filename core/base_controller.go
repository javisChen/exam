package core

import "github.com/beego/beego/v2/server/web"

type BaseController struct {
	web.Controller
}

func (t BaseController) Success(data ...interface{}) {
	t.Data["json"] = success(data)
	_ = t.ServeJSON()
}

func (t BaseController) Error(msg string) {
	t.Data["json"] = error(msg)
	_ = t.ServeJSON()
}
