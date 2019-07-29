package rescode

import (
	"God/core/common"
)

type ErrCode struct {
	Code int
	Msg  string
}

func (e *ErrCode) Result(data interface{}) *common.Result {
	return &common.Result{
		Status: e.Code,
		Msg:    e.Msg,
		Data:   data,
	}
}

func (e *ErrCode) ResultWithMsg(msg string) *common.Result {
	return &common.Result{
		Status: e.Code,
		Msg:    msg,
		Data:   "",
	}
}

func (e *ErrCode) ReplaceMsg(msg string) *ErrCode {
	return &ErrCode{
		Code: e.Code,
		Msg:  msg,
	}
}

var (
	OK = &ErrCode{0, "ok"}

	SYSTEMBUSY_ERROR = &ErrCode{-1, "system error"}

	REQUEST_PARAM_ERR = &ErrCode{1005, "request params valid"}

	LIMIT_REQUEST = &ErrCode{1014, "request is frequent"}
)
