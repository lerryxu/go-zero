package xerr

import (
	"fmt"
)

/**
常用通用固定错误
*/

type Error struct {
	Code uint32 `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data,omitempty"`
}

// 返回给前端的错误码
func (e *Error) GetCode() uint32 {
	return e.Code
}

// 返回给前端显示端错误信息
func (e *Error) GetMsg() string {
	return e.Msg
}

func (e *Error) Error() string {
	return fmt.Sprintf("Code:%d, Msg:%s", e.Code, e.Msg)
}

func NewError(code uint32, msg string) *Error {
	return &Error{Code: code, Msg: msg}
}

func NewErrorCode(code uint32, args ...any) *Error {
	return &Error{Code: code, Msg: GetErrMsg(code, args...)}
}

func NewErrorMsg(msg string) *Error {
	return &Error{Code: SERVER_COMMON_ERROR, Msg: msg}
}

func NewErrorData(code uint32, data any) *Error {
	return &Error{Code: code, Data: data}
}
