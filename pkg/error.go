package pkg

import (
	"errors"
	"net/http"
)

type Error struct {
	Err  error
	Msg  string
	Code int
}

func PanicIfErr(err error) {
	if err != nil {
		PanicError(http.StatusInternalServerError, err)
	}
}

func PanicError(code int, err error) {
	panic(&Error{
		Err:  err,
		Msg:  "请求出错，请稍后尝试",
		Code: code,
	})
}

func PanicErrorWithMsg(code int, msg string) {
	panic(&Error{
		Err:  errors.New(msg),
		Msg:  msg,
		Code: code,
	})
}
