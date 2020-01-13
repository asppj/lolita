package errors

import (
	"fmt"
	"strings"

	"github.com/go-errors/errors"
)

const (
	// StatusOK StatusOK
	StatusOK = 600200 // RFC 7231, 6.3.1
	// StatusBadRequest  StatusBadRequest
	StatusBadRequest = 600400 // RFC 7231, 6.5.1
	// StatusUnauthorized  StatusUnauthorized
	StatusUnauthorized = 600401 // RFC 7235, 3.1
	// StatusForbidden StatusForbidden
	StatusForbidden = 600403 // RFC 7231, 6.5.3
	// StatusNotFound StatusNotFound
	StatusNotFound = 600404 // RFC 7231, 6.5.4
	// StatusNotAcceptable StatusNotAcceptable
	StatusNotAcceptable = 600406 // RFC 7231, 6.5.6
	// StatusRequestTimeout StatusRequestTimeout
	StatusRequestTimeout = 600408 // RFC 7231, 6.5.7
	// StatusInternalServerError StatusInternalServerError
	StatusInternalServerError = 600500 // RFC 7231, 6.6.1
	// StatusServiceUnavailable StatusServiceUnavailable
	StatusServiceUnavailable = 600503 // RFC 7231, 6.6.4
)

// Error Http返回错误
type Error struct {
	err    *errors.Error
	status int
	msg    []string // 添加提示

}

// New 新建error  ineffectual assignment
func New(e interface{}) *Error {
	return NewWithMsg(e, "", 0)
}

// NewWithMsg 返回
func NewWithMsg(e interface{}, msg string, status int) *Error {
	switch eIns := e.(type) {
	case *Error:
		eIns.err = errors.New(eIns.err)
		if msg != "" {
			eIns.msg = append(eIns.msg, msg)
		}
		// 保留最初的status
		if eIns.status <= 0 && status > 0 {
			eIns.status = status
		}
		return eIns
	default:
		err := &Error{
			err: errors.New(eIns),
		}
		if msg != "" {
			err.msg = []string{msg}
		}
		if status > 0 {
			err.status = status
		}
		return err
	}
}

// ErrInfo 转换错误格式为打印模式
type ErrInfo struct {
	Err    string
	Msg    string
	Status int
}

// String  用于打印错误
func (e *Error) String() ErrInfo {
	return ErrInfo{
		Err:    e.err.ErrorStack(),
		Msg:    strings.Join(e.msg, "\n"),
		Status: e.status,
	}
}

// ErrorStack 转为string
func (e *Error) ErrorStack() string {
	return fmt.Sprintf("%+v", e.String())
}
