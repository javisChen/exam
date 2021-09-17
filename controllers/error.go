package controllers

import (
	"exam/core"
)

type ErrorController struct {
	core.BaseController
}

func (c *ErrorController) Error404() {
}

func (c *ErrorController) Error501() {
}

func (c *ErrorController) ErrorDb() {
	c.Data["content"] = "database is now down"
	c.TplName = "dberror.tpl"
}
