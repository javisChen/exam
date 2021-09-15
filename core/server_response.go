package core

type ResponseCode string

const (
	SuccessMsg string = "success"

	SuccessCode ResponseCode = "0"
	ErrorCode   ResponseCode = "50000"
)

type ServerResponse struct {
	Code ResponseCode `json:"code,omitempty"`
	Data interface{}  `json:"data,omitempty"`
	Msg  string       `json:"msg,omitempty"`
}

func success(data interface{}) ServerResponse {
	return ServerResponse{
		Code: SuccessCode,
		Data: data,
		Msg:  SuccessMsg,
	}
}

func error(msg string) ServerResponse {
	return ServerResponse{
		Code: ErrorCode,
		Msg:  msg,
	}
}
