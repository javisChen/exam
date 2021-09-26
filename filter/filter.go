package filter

import (
	"exam/core"
	"exam/security"
	"exam/utils/json"
	"exam/utils/redis"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
	"net/http"
	"strings"
)

func Init() {
	beego.InsertFilter("/*", beego.BeforeExec, func(ctx *context.Context) {
		if ctx.Request.RequestURI != "/exam/auth/v1/login" {
			authorization := ctx.Input.Header("Authorization")
			var accessToken string
			if accessToken = strings.Split(authorization, "Bearer ")[1]; accessToken == "" {
				writeUnauthorized(ctx)
				return
			}
			result, err := redis.Get(fmt.Sprintf(security.LoginUserRedisKey, accessToken))
			if err != nil {
				writeUnauthorized(ctx)
			}
			loginUser := security.LoginUser{}
			json.FromStr(result, &loginUser)
			ctx.Input.SetData(security.LoginUserKey, loginUser)
		}
	})

}

func writeUnauthorized(ctx *context.Context) {
	ctx.Output.SetStatus(http.StatusUnauthorized)
	_ = ctx.Output.JSON(core.Err("Unauthorized"), false, false)
}
