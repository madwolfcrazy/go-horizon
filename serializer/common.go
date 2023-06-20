package serializer

import "fmt"

//Response  API通用返回样板
type Response struct {
	Code  int         `json:"code"`
	Data  interface{} `json:"data"`
	Error string      `json:"error"`
	Msg   string      `json:"msg"`
}

//ErrorResponse 输出错误
func ErrorResponse(err interface{}) Response {
	return Response{
		Code:  500,
		Msg:   "错误",
		Error: fmt.Sprint(err),
	}
}

//ErrorAuthResponse 登录验证错误
func ErrorAuthResponse(err interface{}) Response {
	return Response{
		Code:  403,
		Msg:   "错误",
		Error: fmt.Sprint(err),
	}
}

//Common 普通输出
func Common(v interface{}) Response {
	return Response{
		Code:  0,
		Data:  v,
		Error: "",
	}
}

//CommonOK 输出普通OK结果
func CommonOK() Response {
	return Response{
		Code:  0,
		Msg:   "OK",
		Error: "",
	}
}
