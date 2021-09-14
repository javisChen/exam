package controllers

import (
	"encoding/json"
	"exam/core"
)

type LoginController struct {
	core.BaseController
}

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (c *LoginController) Post() {
	body := c.Ctx.Input.RequestBody //这是获取到的json二进制数据
	var s User
	err := json.Unmarshal(body, &s)
	if err != nil {
		panic(err)
	}
	c.Ok(s)
	c.StopRun()
}
