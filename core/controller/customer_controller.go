package controller

import (
	"God/core/common/comerr"
	"God/core/entity"
	util "God/utils"
	"fmt"
	"net/http"
	"strconv"

	"God/utils/comutil"

	"God/core/service"

	"github.com/gin-gonic/gin"
)

type CustomerController struct{}

var customService = &service.CustomerService{}

func getCustomHeader(ctx *gin.Context) (*entity.ReqHeader, error) {
	// fetch request header contents
	terminalidStr := util.GetValByHeader(ctx, "terminalid")
	devicecodeStr := util.GetValByHeader(ctx, "devicecode")
	versionStr := util.GetValByHeader(ctx, "version")
	terminalId, err := strconv.Atoi(terminalidStr)
	if nil != err {
		return nil, err
	}
	token := util.GetValByHeader(ctx, "token")
	sessionId := util.GetValByHeader(ctx, "sessionId")
	accountIdStr := util.GetValByHeader(ctx, "accountId")
	accountId, err := strconv.Atoi(accountIdStr)
	if nil != err {
		return nil, err
	}
	return &entity.ReqHeader{
		Terminalid: terminalId,
		Devicecode: devicecodeStr,
		Version:    versionStr,
		Token:      token,
		SessionId:  sessionId,
		AccountId:  accountId,
	}, nil
}

func (self *CustomerController) Loan(ctx *gin.Context) {

	ip := util.GetRealRemoteIp(ctx)
	// check IP rate per 10s
	if !util.CheckIpRate(comutil.LOAN, ip, 10) {
		ctx.JSON(http.StatusOK, comerr.LIMIT_REQUEST.ResultEmpty())
		return
	}

	var req entity.LoanReq

	if err := ctx.BindJSON(&req); nil != err {
		ctx.JSON(http.StatusOK, comerr.SYSTEMBUSY_ERROR.ResultWithMsg(err.Error()))
		return
	}

	// check timestamp
	if now, err := util.CheckApiTimeStamp(req.Sign, req.Data.Timestamp); nil != err {
		ctx.JSON(http.StatusOK, comerr.SYSTEMBUSY_ERROR.ResultWithMsgData("failed to check the timestamp: "+err.Error(), now))
		return
	}

	header, err := getCustomHeader(ctx)
	if nil != err {
		ctx.JSON(http.StatusOK, comerr.SYSTEMBUSY_ERROR.ResultWithMsg(err.Error()))
		return
	}

	ok, seId, err := util.CheckTokenAndSession(header.Token, header.SessionId, fmt.Sprint(header.AccountId))
	if nil != err {
		ctx.JSON(http.StatusOK, comerr.SYSTEMBUSY_ERROR.ResultWithMsg(err.Error()))
		return
	}
	if !ok {
		ctx.JSON(http.StatusOK, comerr.NEED_RELOGIN.ResultEmpty())
		return
	}

	// reset sessionId to header
	header.SessionId = seId

	queue := entity.NewSignQueue()
	//queue.AppendSignData("terminalid", fmt.Sprint(header.Terminalid))
	//queue.AppendSignData("devicecode", header.Devicecode)
	//queue.AppendSignData("version", header.Version)
	queue.AppendSignData("timestamp", fmt.Sprint(req.Data.Timestamp))

	// check sign
	if !queue.CheckSign(req.Sign) {
		ctx.JSON(http.StatusOK, comerr.REQUEST_PARAM_ERR.ResultWithMsg("failed to check the sign!!"))
		return
	}

	// To loan balances
	if res, err := customService.Loan(header, &req); nil != err {
		ctx.JSON(http.StatusOK, comerr.SYSTEMBUSY_ERROR.ResultWithMsg("failed to Loan: "+err.Error()))
		return
	} else {
		ctx.JSON(http.StatusOK, comerr.OK.ResultWithMsgData("Loan success", res))
	}
}
