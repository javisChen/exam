package core

import "github.com/beego/beego/v2/server/web"

type BaseController struct {
	web.Controller
}

func (t BaseController) Ok(data interface{}) {
	res := ok(data)
	t.Data["json"] = res
	_ = t.ServeJSON() //对json进行序列化输出
}
