package auth_handler

import (
	"exam/core/web"
	userDao "exam/dao/user"
	"exam/models/resp"
	"exam/security"
	"exam/utils/json"
	"exam/utils/redis"
	"exam/utils/uuid"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

type RequestContext struct {
	c *gin.Context
}

func Login(c *gin.Context) {
	jsonParam, err := web.GetJsonParam(c)
	if err != nil {
		web.ErrorWithMsg(c, err.Error())
	}
	fmt.Println("请求参数 >", jsonParam)
	phone := (*jsonParam)["phone"].(string)
	password := (*jsonParam)["password"].(string)

	var user, _ = userDao.SelectByPhone(phone)
	if user.Password != password {
		web.ErrorWithMsg(c, "用户名或密码有误")
	}

	uuidStr := uuid.UUID()
	loginUser := security.NewLoginUser(user.Id, user.Phone, user.Username)
	err = redis.SetWithExpire(fmt.Sprintf(security.LoginUserRedisKey, uuidStr), json.ToJSONStr(loginUser), 24, time.Hour)
	if err != nil {
		web.ErrorWithMsg(c, err.Error())
	}
	web.Ok(c, resp.LoginResp{AccessToken: uuidStr})
}

func Logout(c *gin.Context) {
	accessToken := web.GetAccessToken(c)
	result, err := redis.Get(fmt.Sprintf(security.LoginUserRedisKey, accessToken))
	if err != nil {
		web.ErrorWithMsg(c, err.Error())
	}
	_ = redis.Remove(fmt.Sprintf(security.LoginUserRedisKey, result))
	web.Ok(c)
}
