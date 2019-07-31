package controller

import (
	"God/core/entity"
	"God/core/service"
	"God/utils"
	"net/http"
	"strconv"

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

	var res entity.RegisterReq

	if err := ctx.BindJSON(&res); nil != err {
		ctx.JSON(http.StatusOK, comerr.REQUEST_PARAM_ERR.ResultWithMsg(err.Error()))
		return
	}

	// check mobileNo
	if !util.CheckPhoneNo(res.MobileNo) {
		ctx.JSON(http.StatusOK, comerr.REQUEST_PARAM_ERR.ResultWithMsg("the mobile No. is wrong!!"))
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

	// check sign
	if !queue.CheckSign(sign) {
		ctx.JSON(http.StatusOK, comerr.REQUEST_PARAM_ERR.ResultWithMsg("failed to check the sign!!"))
		return
	}

	// check timestamp
	if err := util.CheckApiTimeStamp(sign, res.Timestamp); nil != err {
		ctx.JSON(http.StatusOK, comerr.REQUEST_PARAM_ERR.ResultWithMsg("failed to check the timestamp: "+err.Error()))
		return
	}

}

func (self *UserController) GetUserInfo(ctx *gin.Context) {

	id := ctx.Query("id")

	if idInt, err := strconv.Atoi(id); nil != err {
		ctx.JSON(http.StatusOK, comerr.REQUEST_PARAM_ERR.ResultWithMsg(err.Error()))
		return
	} else {
		userInfo, err := userService.GetUserInfo(uint32(idInt))
		if nil != err {
			ctx.JSON(http.StatusOK, comerr.REQUEST_PARAM_ERR.ResultWithMsg(err.Error()))
			return
		}
		ctx.JSON(http.StatusOK, comerr.OK.Result(userInfo))
	}
}
