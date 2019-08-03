package controller

import (
	"God/core/entity"
	"God/core/service"
	util "God/utils"
	"fmt"
	"net/http"
	"strconv"

	"github.com/spf13/viper"

	"God/core/common/comerr"

	"github.com/gin-gonic/gin"
)

type UserController struct {
}

var userService = &service.UserService{}

func (self *UserController) Register(ctx *gin.Context) {

	ip := util.GetRealRemoteIp(ctx)
	// check IP rate
	if !util.CheckIpRate("register", ip, 10) {
		ctx.JSON(http.StatusOK, comerr.LIMIT_REQUEST.Result(nil))
		return
	}

	sign := ctx.Query("sign")

	// fetch request header contents
	terminalidStr := util.GetValByHeader(ctx, "terminalid")
	devicecodeStr := util.GetValByHeader(ctx, "devicecode")
	versionStr := util.GetValByHeader(ctx, "version")
	terminalid, err := strconv.Atoi(terminalidStr)
	if nil != err {
		ctx.JSON(http.StatusOK, comerr.REQUEST_PARAM_ERR.ResultWithMsg("parse terminalid is failed: "+err.Error()))
		return
	}
	header := &entity.ReqHeader{
		Terminalid: terminalid,
		Devicecode: devicecodeStr,
		Version:    versionStr,
	}

	var req entity.RegisterReq

	if err := ctx.BindJSON(&req); nil != err {
		ctx.JSON(http.StatusOK, comerr.REQUEST_PARAM_ERR.ResultWithMsg(err.Error()))
		return
	}

	// check mobileNo
	if !util.CheckPhoneNo(req.MobileNo) {
		ctx.JSON(http.StatusOK, comerr.REQUEST_PARAM_ERR.ResultWithMsg("the mobile No. is wrong!!"))
		return
	}

	pwdLenLimit := viper.GetInt("common.pwdLenLimit")
	// check pwds length
	if !util.CheckPwdLen(req.LoginPassword, pwdLenLimit) {
		ctx.JSON(http.StatusOK, comerr.REQUEST_PARAM_ERR.ResultWithMsg(fmt.Sprintf("the LoginPassword length is wrong, must be %d char!!", pwdLenLimit)))
		return
	}
	if !util.CheckPwdLen(req.TradePassword, pwdLenLimit) {
		ctx.JSON(http.StatusOK, comerr.REQUEST_PARAM_ERR.ResultWithMsg(fmt.Sprintf("the TradePassword length is wrong, must be %d char!!", pwdLenLimit)))
		return
	}

	queue := make(entity.SignQueue, 0)
	queue = queue.AppendSignData("terminalid", terminalidStr)
	queue = queue.AppendSignData("devicecode", devicecodeStr)
	queue = queue.AppendSignData("version", versionStr)
	queue = queue.AppendSignData("timestamp", fmt.Sprint(req.Timestamp))
	queue = queue.AppendSignData("mobileNo", req.MobileNo)
	queue = queue.AppendSignData("loginPassword", req.LoginPassword)
	queue = queue.AppendSignData("tradePassword", req.TradePassword)

	// check sign
	if !queue.CheckSign(sign) {
		ctx.JSON(http.StatusOK, comerr.REQUEST_PARAM_ERR.ResultWithMsg("failed to check the sign!!"))
		return
	}

	// check timestamp
	if err := util.CheckApiTimeStamp(sign, req.Timestamp); nil != err {
		ctx.JSON(http.StatusOK, comerr.REQUEST_PARAM_ERR.ResultWithMsg("failed to check the timestamp: "+err.Error()))
		return
	}

	if accountId, err := userService.Register(header, &req); nil != err {
		ctx.JSON(http.StatusOK, comerr.SYSTEMBUSY_ERROR.ResultWithMsg("failed to register: "+err.Error()))
		return
	} else {
		ctx.JSON(http.StatusOK, comerr.OK.ResultWithMsg(fmt.Sprintf("Register success, the accountId: %d", accountId)))
	}
}

func (self *UserController) GetUserBase(ctx *gin.Context) {

	id := ctx.Query("id")

	if idInt, err := strconv.Atoi(id); nil != err {
		ctx.JSON(http.StatusOK, comerr.REQUEST_PARAM_ERR.ResultWithMsg(err.Error()))
		return
	} else {
		userInfo, err := userService.GetUserBase(uint32(idInt))
		if nil != err {
			ctx.JSON(http.StatusOK, comerr.SYSTEMBUSY_ERROR.ResultWithMsg(err.Error()))
			return
		} else if nil == userInfo {
			ctx.JSON(http.StatusOK, comerr.EMPTY_RESULT.ResultEmpty())
		} else {
			ctx.JSON(http.StatusOK, comerr.OK.Result(userInfo))
		}
	}
}
