package controllers

import (
	"exam/core"
	"exam/models"
	"exam/models/resp"
	"exam/utils/json"
	"exam/utils/redis"
	"exam/utils/uuid"
	"fmt"
	"time"
)

var loginUserRedisKey = "login_user:tk:%s"

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

	uuidStr := uuid.UUID()
	user.Password = ""
	err := redis.SetWithExpire(fmt.Sprintf(loginUserRedisKey, uuidStr), json.ToJSONStr(user), 24, time.Hour)
	if err != nil {
		c.Error(err.Error())
	}
	c.Success(resp.LoginResp{AccessToken: uuidStr})
}

func (c LoginController) Logout() {
	accessToken := c.GetAccessToken()
	result, err := redis.Get(fmt.Sprintf(loginUserRedisKey, accessToken))
	if err != nil {
		c.Error(err.Error())
	}
	_ = redis.Remove(fmt.Sprintf(loginUserRedisKey, result))
	c.Success()
}
