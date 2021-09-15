package controllers

import (
	"exam/core"
	"exam/models"
)

type LoginController struct {
	core.BaseController
}

func (c LoginController) Login() {
	var users []models.User
	o := core.GetOrm()
	o.Raw("select * from user").QueryRows(&users)
	c.Success(users)
}

func (c LoginController) Logout() {
	c.Success()
}
