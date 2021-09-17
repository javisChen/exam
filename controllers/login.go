package controllers

import (
	"exam/core"
	"exam/models"
	"fmt"
)

type LoginController struct {
	core.BaseController
}

func (c LoginController) Login() {
	jsonParam := c.GetJsonParam()
	fmt.Println("请求参数 >", jsonParam)
	phone := jsonParam["phone"].(string)
	password := jsonParam["password"].(string)

	var user models.User
	_ = core.GetOrm().Raw("select username, password, phone from user where phone = ?", phone).QueryRow(&user)
	if user.Password != password {
		c.Error("用户名或密码有误")
	}
	c.SetSession("login_user", user)
	c.Success()
}

func (c LoginController) Logout() {
	c.SetSession("login_user", nil)
	c.Success()
}
