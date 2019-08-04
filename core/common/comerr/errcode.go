package comerr

import (
	"God/core/common"
	"God/utils"
)

type ErrCode struct {
	Code    string
	Message string
	Success bool
}

func (e *ErrCode) Result(data interface{}, sign string) *common.Result {
	return &common.Result{
		Code:    e.Code,
		Message: e.Message,
		Success: e.Success,
		Data:    data,
		Sign:    sign,
	}
}

func (e *ErrCode) ResultEmpty() *common.Result {
	return &common.Result{
		Code:    e.Code,
		Message: e.Message,
		Success: e.Success,
		Data:    "",
		Sign:    util.Md5(""),
	}
}

func (e *ErrCode) ResultWithMsg(msg string) *common.Result {
	return &common.Result{
		Code:    e.Code,
		Message: msg,
		Success: e.Success,
		Data:    "",
		Sign:    util.Md5(""),
	}
}

func (e *ErrCode) ResultWithMsgData(msg string, data interface{}) *common.Result {
	return &common.Result{
		Code:    e.Code,
		Message: msg,
		Success: e.Success,
		Data:    data,
		Sign:    util.Md5(""),
	}
}

var (
	OK = &ErrCode{"0000", "ok", true}

	SYSTEMBUSY_ERROR = &ErrCode{"1000", "system error", false}

	REQUEST_PARAM_ERR = &ErrCode{"1005", "request params valid", false}

	LIMIT_REQUEST = &ErrCode{"1014", "request is frequent", false}

	NEED_RELOGIN = &ErrCode{"2000", "you have not login", false}

	EMPTY_RESULT = &ErrCode{"6003", "result is empty", false}
)
