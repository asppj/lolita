package api

import "github.com/asppj/t-go-opentrace/ext/errors"

// Response 返回格式
type Response struct {
	Success bool        `json:"success"`        // 请求是否成功
	Data    interface{} `json:"data"`           // 数据
	Code    int         `json:"code,omitempty"` // 错误码
	Msg     string      `json:"msg"`            // 错误提示
}

// newRes 新建
func newRes(success bool, code int, msg string, data interface{}) *Response {
	return &Response{
		Success: success,
		Data:    data,
		Code:    code,
		Msg:     msg,
	}
}

// DefaultRes 默认
func DefaultRes() *Response {
	return newRes(true, errors.StatusOK, SUCCESS, struct{}{})
}

// failedRes 失败
func failedRes() *Response {
	return newRes(false, errors.StatusUnCustomize, unCustomize, struct{}{})
}

// successRes 成功
func successRes() *Response {
	return DefaultRes()
}
