package middleware

import (
	"exam/core/web"
	"exam/security"
	"exam/utils/json"
	"exam/utils/redis"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func AuthMiddleWare(ctx *gin.Context) {
	if ctx.Request.RequestURI != "/exam/auth/v1/login" {
		authorization := ctx.GetHeader("Authorization")
		var accessToken string
		bearer := strings.Split(authorization, "Bearer ")
		if bearer == nil || len(bearer) < 2 {
			writeUnauthorized(ctx)
			return
		}
		if accessToken = bearer[1]; accessToken == "" {
			writeUnauthorized(ctx)
			return
		}
		result, err := redis.Get(fmt.Sprintf(security.LoginUserRedisKey, accessToken))
		if err != nil {
			writeUnauthorized(ctx)
			return
		}
		loginUser := security.LoginUser{}
		json.FromStr(result, &loginUser)
		ctx.Set(security.LoginUserRedisKey, loginUser)
	}
}

func writeUnauthorized(ctx *gin.Context) {
	ctx.Abort()
	web.ErrorWithStatusCode(ctx, http.StatusUnauthorized, "Unauthorized")
}
