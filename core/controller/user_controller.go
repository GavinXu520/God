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

	"God/utils/comutil"

	"github.com/gin-gonic/gin"
)

// omitempty 排除 0 值
type UserController struct{}

var userService = &service.UserService{}

type checkData struct {
	MobileNo  string
	Sign      string
	Timestamp int
}

func getUserHeader(ctx *gin.Context) (*entity.ReqHeader, error) {
	// fetch request header contents
	terminalidStr := util.GetValByHeader(ctx, "terminalid")
	devicecodeStr := util.GetValByHeader(ctx, "devicecode")
	versionStr := util.GetValByHeader(ctx, "version")
	terminalId, err := strconv.Atoi(terminalidStr)
	if nil != err {
		return nil, err
	}
	return &entity.ReqHeader{
		Terminalid: terminalId,
		Devicecode: devicecodeStr,
		Version:    versionStr,
	}, nil
}

func commonCheck(ctx *gin.Context, data *checkData, ip_limie_prefix string) bool {
	ip := util.GetRealRemoteIp(ctx)
	// check IP rate per 10s
	if !util.CheckIpRate(ip_limie_prefix, ip, 10) {
		ctx.JSON(http.StatusOK, comerr.LIMIT_REQUEST.ResultEmpty())
		return false
	}
	// check mobileNo
	if !util.CheckPhoneNo(data.MobileNo) {
		ctx.JSON(http.StatusOK, comerr.REQUEST_PARAM_ERR.ResultWithMsg("the mobile No. is wrong!!"))
		return false
	}
	// check timestamp
	if now, err := util.CheckApiTimeStamp(data.Sign, data.Timestamp); nil != err {
		ctx.JSON(http.StatusOK, comerr.SYSTEMBUSY_ERROR.ResultWithMsgData("failed to check the timestamp: "+err.Error(), now))
		return false
	}
	return true
}

func (self *UserController) Register(ctx *gin.Context) {

	var req entity.RegisterReq

	if err := ctx.BindJSON(&req); nil != err {
		ctx.JSON(http.StatusOK, comerr.SYSTEMBUSY_ERROR.ResultWithMsg(err.Error()))
		return
	}

	if !commonCheck(ctx, &checkData{MobileNo: req.Data.MobileNo, Sign: req.Sign, Timestamp: req.Data.Timestamp}, comutil.REGISTER) {
		return
	}

	header, err := getUserHeader(ctx)
	if nil != err {
		ctx.JSON(http.StatusOK, comerr.SYSTEMBUSY_ERROR.ResultWithMsg(err.Error()))
		return
	}

	pwdLenLimit := viper.GetInt("common.pwdLenLimit")
	// check pwds length
	if !util.CheckPwdLen(req.Data.LoginPassword, pwdLenLimit) {
		ctx.JSON(http.StatusOK, comerr.REQUEST_PARAM_ERR.ResultWithMsg(fmt.Sprintf("the LoginPassword length is wrong, must be %d char!!", pwdLenLimit)))
		return
	}
	if !util.CheckPwdLen(req.Data.TradePassword, pwdLenLimit) {
		ctx.JSON(http.StatusOK, comerr.REQUEST_PARAM_ERR.ResultWithMsg(fmt.Sprintf("the TradePassword length is wrong, must be %d char!!", pwdLenLimit)))
		return
	}

	queue := entity.NewSignQueue()
	//queue.AppendSignData("terminalid", fmt.Sprint(header.Terminalid))
	//queue.AppendSignData("devicecode", header.Devicecode)
	//queue.AppendSignData("version", header.Version)
	queue.AppendSignData("timestamp", fmt.Sprint(req.Data.Timestamp))
	queue.AppendSignData("mobileNo", req.Data.MobileNo)
	queue.AppendSignData("loginPassword", req.Data.LoginPassword)
	queue.AppendSignData("tradePassword", req.Data.TradePassword)

	// check sign
	if !queue.CheckSign(req.Sign) {
		ctx.JSON(http.StatusOK, comerr.REQUEST_PARAM_ERR.ResultWithMsg("failed to check the sign!!"))
		return
	}

	if res, err := userService.Register(header, &req); nil != err {
		ctx.JSON(http.StatusOK, comerr.SYSTEMBUSY_ERROR.ResultWithMsg("failed to register: "+err.Error()))
		return
	} else {
		ctx.JSON(http.StatusOK, comerr.OK.ResultWithMsgData("Register success", res))
	}
}

func (self *UserController) LoginByPwd(ctx *gin.Context) {
	var req entity.LoginReq

	if err := ctx.BindJSON(&req); nil != err {
		ctx.JSON(http.StatusOK, comerr.SYSTEMBUSY_ERROR.ResultWithMsg(err.Error()))
		return
	}

	if !commonCheck(ctx, &checkData{MobileNo: req.Data.MobileNo, Sign: req.Sign, Timestamp: req.Data.Timestamp}, comutil.LOGIN) {
		return
	}

	header, err := getUserHeader(ctx)
	if nil != err {
		ctx.JSON(http.StatusOK, comerr.SYSTEMBUSY_ERROR.ResultWithMsg(err.Error()))
		return
	}

	pwdLenLimit := viper.GetInt("common.pwdLenLimit")
	// check pwds length
	if !util.CheckPwdLen(req.Data.LoginPassword, pwdLenLimit) {
		ctx.JSON(http.StatusOK, comerr.REQUEST_PARAM_ERR.ResultWithMsg(fmt.Sprintf("the LoginPassword length is wrong, must be %d char!!", pwdLenLimit)))
		return
	}

	queue := entity.NewSignQueue()
	//queue.AppendSignData("terminalid", fmt.Sprint(header.Terminalid))
	//queue.AppendSignData("devicecode", header.Devicecode)
	//queue.AppendSignData("version", header.Version)
	queue.AppendSignData("timestamp", fmt.Sprint(req.Data.Timestamp))
	queue.AppendSignData("mobileNo", req.Data.MobileNo)
	queue.AppendSignData("loginPassword", req.Data.LoginPassword)

	// check sign
	if !queue.CheckSign(req.Sign) {
		ctx.JSON(http.StatusOK, comerr.REQUEST_PARAM_ERR.ResultWithMsg("failed to check the sign!!"))
		return
	}
	if res, err := userService.LoginByMobile(header, &req, true); nil != err {
		ctx.JSON(http.StatusOK, comerr.SYSTEMBUSY_ERROR.ResultWithMsg("failed to login: "+err.Error()))
		return
	} else {
		ctx.JSON(http.StatusOK, comerr.OK.ResultWithMsgData("Login success", res))
	}
}

func (self *UserController) LoginByMobileSms(ctx *gin.Context) {
	var req entity.LoginReq

	if err := ctx.BindJSON(&req); nil != err {
		ctx.JSON(http.StatusOK, comerr.SYSTEMBUSY_ERROR.ResultWithMsg(err.Error()))
		return
	}

	if !commonCheck(ctx, &checkData{MobileNo: req.Data.MobileNo, Sign: req.Sign, Timestamp: req.Data.Timestamp}, comutil.LOGIN) {
		return
	}

	header, err := getUserHeader(ctx)
	if nil != err {
		ctx.JSON(http.StatusOK, comerr.SYSTEMBUSY_ERROR.ResultWithMsg(err.Error()))
		return
	}

	queue := entity.NewSignQueue()
	//queue.AppendSignData("terminalid", fmt.Sprint(header.Terminalid))
	//queue.AppendSignData("devicecode", header.Devicecode)
	//queue.AppendSignData("version", header.Version)
	queue.AppendSignData("timestamp", fmt.Sprint(req.Data.Timestamp))
	queue.AppendSignData("mobileNo", req.Data.MobileNo)
	queue.AppendSignData("smsCode", req.Data.SmsCode)

	// check sign
	if !queue.CheckSign(req.Sign) {
		ctx.JSON(http.StatusOK, comerr.REQUEST_PARAM_ERR.ResultWithMsg("failed to check the sign!!"))
		return
	}

	// check smscode
	if pass, err := checkSmsCode(req.Data.MobileNo, req.Data.SmsCode); nil != err {
		ctx.JSON(http.StatusOK, comerr.SYSTEMBUSY_ERROR.ResultWithMsg("failed to check smscode: "+err.Error()))
		return
	} else if !pass {
		ctx.JSON(http.StatusOK, comerr.SYSTEMBUSY_ERROR.ResultWithMsg("failed to login, the smscode validate is failed: "+err.Error()))
		return
	}

	if res, err := userService.LoginByMobile(header, &req, false); nil != err {
		ctx.JSON(http.StatusOK, comerr.SYSTEMBUSY_ERROR.ResultWithMsg("failed to login: "+err.Error()))
		return
	} else {
		ctx.JSON(http.StatusOK, comerr.OK.ResultWithMsgData("Login success", res))
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
			ctx.JSON(http.StatusOK, comerr.OK.Result(userInfo, ""))
		}
	}
}
