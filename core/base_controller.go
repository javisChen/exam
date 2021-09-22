package core

import (
	"exam/utils"
	json2 "exam/utils/json"
	"github.com/beego/beego/v2/server/web"
	"github.com/mitchellh/mapstructure"
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
	v := make(map[string]interface{})
	err := json2.FromBytes(t.Ctx.Input.RequestBody, &v)
	utils.TryThrowError(err)
	return v
}

func (t BaseController) ParseFromJsonParam(v interface{}) interface{} {
	result := make(map[string]interface{})
	err := json2.FromBytes(t.Ctx.Input.RequestBody, &result)
	utils.TryThrowError(err)

	err = mapstructure.Decode(result, &v)
	utils.TryThrowError(err)
	return v
}
