package errcode

import (
	"fmt"
	"net/http"
)

type Error struct {
	code int `json:"code"`
	msg string `json:"msg"`
	details []string `json:"details"`
}

var codes = map[int]string{}

func NewError(code int,msg string) *Error {
	if _,ok := codes[code];ok {
		panic(fmt.Sprintf("错误码%d已经存在,请更换一个",code))
	}
	codes[code] = msg
	return &Error{code: code,msg: msg}
}

func (e *Error)Error() string {
	return fmt.Sprintf("错误码: %d, 错误信息: %s",e.Code(),e.Msg())
}

func (e *Error) Code() int {
	return e.code
}

func (e *Error) Msg() string {
	return e.msg
}

func (e *Error)Msgf(args []interface{}) string {
	return fmt.Sprintf(e.msg,args...)
}
func (e *Error)Details() []string {
	return e.details
}

func (e *Error)WithDetails(details ...string) *Error {
	e.details = []string{}
	for _,d := range details{
		e.details = append(e.details,d)
	}
	return e
}

func (e *Error) StatusCode() int {
	switch e.Code() {
	case Success.code:
		return http.StatusOK
	case ServerError.code:
		return http.StatusInternalServerError
	case InvalidParams.code:
		return http.StatusBadRequest
	case UnauthorizeAuthNotExist.code:
		fallthrough
	case UnauthorizedTokenError.code:
		fallthrough
	case UnauthorizedTokenGenerate.code:
		fallthrough
	case UnauthorizedTokenTimeout.code:
		return http.StatusUnauthorized
	case ToManyRequests.code:
		return http.StatusTooManyRequests
	}

	return http.StatusInternalServerError

}