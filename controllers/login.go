package controllers

import (
	"exam/core"
	userDao "exam/dao/user"
	"exam/models/resp"
	"exam/security"
	"exam/utils/json"
	"exam/utils/redis"
	"exam/utils/uuid"
	"fmt"
	"time"
)

type LoginController struct {
	core.BaseController
}

func (c LoginController) Login() {
	jsonParam := c.GetJsonParam()
	fmt.Println("请求参数 >", jsonParam)
	phone := jsonParam["phone"].(string)
	password := jsonParam["password"].(string)

	var user, _ = userDao.SelectByPhone(phone)
	if user.Password != password {
		c.Error("用户名或密码有误")
	}

	uuidStr := uuid.UUID()
	loginUser := security.NewLoginUser(user.Id, user.Phone, user.Username)
	err := redis.SetWithExpire(fmt.Sprintf(security.LoginUserRedisKey, uuidStr), json.ToJSONStr(loginUser), 24, time.Hour)
	if err != nil {
		c.Error(err.Error())
	}
	c.Success(resp.LoginResp{AccessToken: uuidStr})
}

func (c LoginController) Logout() {
	accessToken := c.GetAccessToken()
	result, err := redis.Get(fmt.Sprintf(security.LoginUserRedisKey, accessToken))
	if err != nil {
		c.Error(err.Error())
	}
	_ = redis.Remove(fmt.Sprintf(security.LoginUserRedisKey, result))
	c.Success()
}
