package core

type ResponseCode string

const (
	SuccessMsg string = "Success"

	SuccessCode ResponseCode = "0"
	ErrorCode   ResponseCode = "50000"
)

type ServerResponse struct {
	Code ResponseCode `json:"code,omitempty"`
	Data interface{}  `json:"data,omitempty"`
	Msg  string       `json:"msg,omitempty"`
}

func Success(data interface{}) ServerResponse {
	return ServerResponse{
		Code: SuccessCode,
		Data: data,
		Msg:  SuccessMsg,
	}
}

func Err(msg string) ServerResponse {
	return ServerResponse{
		Code: ErrorCode,
		Msg:  msg,
	}
}
