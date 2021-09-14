package core

type ServerResponse struct {
	Code string      `json:"code,omitempty"`
	Data interface{} `json:"data,omitempty"`
	Msg  string      `json:"msg,omitempty"`
}

func ok(data interface{}) ServerResponse {
	return ServerResponse{
		Code: "0",
		Data: data,
		Msg:  "success",
	}
}
