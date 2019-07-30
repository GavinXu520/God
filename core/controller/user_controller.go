package controller

import (
	"God/core/common/rescode"
	"net/http"

	"God/core/service"

	"strconv"

	"God/utils"

	"God/core/entity"

	"time"

	"github.com/gin-gonic/gin"
)

type UserController struct {
}

var userService = &service.UserService{}

func (self *UserController) Register(ctx *gin.Context) {

	ip := util.GetRealRemoteIp(ctx)
	// TODO redis 锁校验IP
	if !util.CheckIpRate("register", ip, time.Second*10) {
		ctx.JSON(http.StatusOK, rescode.LIMIT_REQUEST.Result(nil))
		return
	}

	sign := ctx.Query("sign")

	// fetch request header contents
	terminalidStr := util.GetValByHeader(ctx, "terminalid")
	devicecodeStr := util.GetValByHeader(ctx, "devicecode")
	versionStr := util.GetValByHeader(ctx, "version")

	var res entity.RegisterReq

	if err := ctx.BindJSON(&res); nil != err {
		ctx.JSON(http.StatusOK, rescode.REQUEST_PARAM_ERR.ResultWithMsg(err.Error()))
		return
	}

	if !util.CheckPhoneNo(res.MobileNo) {
		ctx.JSON(http.StatusOK, rescode.REQUEST_PARAM_ERR.ResultWithMsg("the mobile no. is wrong!!"))
		return
	}

	queue := make(entity.SignQueue, 0)
	queue = queue.AppendSignData("terminalid", terminalidStr)
	queue = queue.AppendSignData("devicecode", devicecodeStr)
	queue = queue.AppendSignData("version", versionStr)
	queue = queue.AppendSignData("timestamp", res.Timestamp)
	queue = queue.AppendSignData("mobileNo", res.MobileNo)
	queue = queue.AppendSignData("loginPassword", res.LoginPassword)
	queue = queue.AppendSignData("tradePassword", res.TradePassword)

	if !queue.CheckSign(sign) {
		ctx.JSON(http.StatusOK, rescode.REQUEST_PARAM_ERR.ResultWithMsg("failed to check the sign!!"))
		return
	}

	// todo

}

func (self *UserController) GetUserInfo(ctx *gin.Context) {

	id := ctx.Query("id")

	if idInt, err := strconv.Atoi(id); nil != err {
		ctx.JSON(http.StatusOK, rescode.REQUEST_PARAM_ERR.ResultWithMsg(err.Error()))
		return
	} else {
		userInfo, err := userService.GetUserInfo(uint32(idInt))
		if nil != err {
			ctx.JSON(http.StatusOK, rescode.REQUEST_PARAM_ERR.ResultWithMsg(err.Error()))
			return
		}
		ctx.JSON(http.StatusOK, rescode.OK.Result(userInfo))
	}
}
