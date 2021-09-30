package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

type ResponseCode int

const (
	SuccessMsg string = "Success"
	ErrorMsg   string = "Fail"

	SuccessCode ResponseCode = 0
	ErrorCode   ResponseCode = 50000
)

func Result(code ResponseCode, data interface{}, msg string) *Response {
	// 开始时间
	return &Response{
		int(code),
		data,
		msg,
	}
}

func Ok(c *gin.Context, data ...interface{}) {
	write(c, http.StatusOK, Result(SuccessCode, data, SuccessMsg))
}

func Error(c *gin.Context) {
	write(c, http.StatusOK, Result(ErrorCode, nil, ErrorMsg))
}

func ErrorWithMsg(c *gin.Context, message string) {
	write(c, http.StatusOK, Result(ErrorCode, nil, message))
}

func ErrorWithStatusCode(c *gin.Context, statusCode int, message string) {
	write(c, statusCode, Result(ErrorCode, nil, message))
}

func write(c *gin.Context, statusCode int, result *Response) {
	c.JSON(statusCode, result)
}
