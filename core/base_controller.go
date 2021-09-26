package core

import (
	"exam/security"
	"exam/utils"
	"exam/utils/json"
	"github.com/beego/beego/v2/server/web"
	"github.com/mitchellh/mapstructure"
	"strings"
)

type BaseController struct {
	web.Controller
}

func (t BaseController) Success(data ...interface{}) {
	t.Data["json"] = Success(data)
	_ = t.ServeJSON()
	t.StopRun()
}

func (t BaseController) Error(msg string) {
	t.Data["json"] = Err(msg)
	_ = t.ServeJSON()
	t.StopRun()
}

func (t BaseController) GetJsonParam() map[string]interface{} {
	v := make(map[string]interface{})
	err := json.FromBytes(t.Ctx.Input.RequestBody, &v)
	utils.TryThrowError(err)
	return v
}

func (t BaseController) ParseFromJsonParam(v interface{}) interface{} {
	result := make(map[string]interface{})
	err := json.FromBytes(t.Ctx.Input.RequestBody, &result)
	utils.TryThrowError(err)

	err = mapstructure.Decode(result, &v)
	utils.TryThrowError(err)
	return v
}

func (t BaseController) GetAccessToken() string {
	authorization := t.Ctx.Input.Header("Authorization")
	accessToken := strings.Split(authorization, "Bearer ")[1]
	return accessToken
}

func (t BaseController) GetLoginUser() security.LoginUser {
	data := t.Ctx.Input.GetData(security.LoginUserKey)
	return data.(security.LoginUser)
}
