package controllers

import (
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Post() {
	data := c.Data
	fmt.Println(data)
	c.Ctx.WriteString("login success")
}
