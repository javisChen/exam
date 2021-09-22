package controllers

import (
	"exam/core"
)

type ErrorController struct {
	core.BaseController
}

func (c *ErrorController) Error404() {
	c.Error("404 NOT FOUND")
}

func (c *ErrorController) Error501() {
	c.Error("501")
}

func (c *ErrorController) Error5050() {
	c.Error("500")
}

func (c *ErrorController) ErrorDb() {
	c.Data["content"] = "database is now down"
	c.TplName = "dberror.tpl"
}
