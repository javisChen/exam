package core

import (
	"encoding/json"
	"exam/utils"
	"github.com/beego/beego/v2/server/web"
)

type BaseController struct {
	web.Controller
}

func (t BaseController) Success(data ...interface{}) {
	t.Data["json"] = success(data)
	_ = t.ServeJSON()
	t.StopRun()
}

func (t BaseController) Error(msg string) {
	t.Data["json"] = error(msg)
	_ = t.ServeJSON()
	t.StopRun()
}

func (t BaseController) GetJsonParam() map[string]interface{} {
	m := make(map[string]interface{})
	err := json.Unmarshal(t.Ctx.Input.RequestBody, &m)
	utils.TryThrowError(err)
	return m
}
