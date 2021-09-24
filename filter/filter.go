package filter

import (
	"exam/core"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
	"net/http"
	"strings"
)

func Init() {
	beego.InsertFilter("/*", beego.BeforeExec, func(ctx *context.Context) {
		ctx.Input.SetData("user", "123")
		if ctx.Request.RequestURI != "/exam/auth/v1/login" {
			authorization := ctx.Input.Header("Authorization")
			if authorization == "" || strings.Split(authorization, "Bearer ")[1] == "" {
				ctx.Output.SetStatus(http.StatusUnauthorized)
				_ = ctx.Output.JSON(core.Err("Unauthorized"), false, false)
				return
			}
		}
	})

}
