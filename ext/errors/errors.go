package errors

import (
	"fmt"
	"strings"

	"github.com/go-errors/errors"
)

// Error Http返回错误
type Error struct {
	err  *errors.Error
	code int      // 前端错误提示
	msg  []string // 添加提示。不返回给前端只作为日志打印

}

// New 新建error  ineffectual assignment
func New(e interface{}) *Error {
	return NewWithMsg(e, "", 0)
}

// NewWithMsg 返回
func NewWithMsg(e interface{}, msg string, code int) *Error {
	switch eIns := e.(type) {
	case *Error:
		eIns.err = errors.New(eIns.err)
		if msg != "" {
			eIns.msg = append(eIns.msg, msg)
		}
		// 保留最初的status
		if eIns.code <= 0 && code > 0 {
			eIns.code = code
		}
		return eIns
	default:
		err := &Error{
			err: errors.New(eIns),
		}
		if msg != "" {
			err.msg = []string{msg}
		}
		if code > 0 {
			err.code = code
		}
		return err
	}
}

// String  用于打印错误
func (e *Error) Error() *ErrInfo {
	return &ErrInfo{
		Err:  e.err.ErrorStack(),
		Msg:  strings.Join(e.msg, "\n"),
		Code: e.code,
	}
}

// ErrorString  用于打印错误
func (e *Error) ErrorString() string {
	return e.Error().String()
}

// ErrorStack 转为string
func (e *Error) ErrorStack() string {
	return fmt.Sprintf("%+v", e.Error())
}

// Code 状态码
func (e *Error) Code() int {
	return e.code
}

// Msg 错误消息
func (e *Error) Msg() string {
	return strings.Join(e.msg, "\n")
}

// ErrInfo 转换错误格式为打印模式
type ErrInfo struct {
	Err  string
	Msg  string
	Code int
}

// String 打印
func (i *ErrInfo) String() string {
	return fmt.Sprintf("Err:%s\nMsg:%s\nCode:%d", i.Err, i.Msg, i.Code)
}
