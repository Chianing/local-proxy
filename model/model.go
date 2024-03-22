package model

import "local-proxy/constants"

type ServerResponse struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

func buildResponse(code int, data interface{}, msg string) *ServerResponse {
	return &ServerResponse{
		Code: code,
		Data: data,
		Msg:  msg,
	}
}

func BuildSuccessResponse(data interface{}) *ServerResponse {
	return buildResponse(int(constants.CodeSuccess), data, "OK")
}

func BuildFailureResponse(msg string) *ServerResponse {
	return buildResponse(int(constants.CodeFailure), nil, msg)
}
