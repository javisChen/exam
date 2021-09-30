package web

import (
	"exam/security"
	"github.com/gin-gonic/gin"
	"strings"
)

func GetJsonParam(c *gin.Context) (*map[string]interface{}, error) {
	m := make(map[string]interface{})
	err := c.ShouldBindJSON(&m)
	if err != nil {
		return nil, err
	}
	return &m, err
}

func GetAccessToken(c *gin.Context) string {
	authorization := c.GetHeader("Authorization")
	accessToken := strings.Split(authorization, "Bearer ")[1]
	return accessToken
}

func GetLoginUser(c *gin.Context) security.LoginUser {
	value, _ := c.Get(security.LoginUserKey)
	return value.(security.LoginUser)
}
