package controller

import (
	"God/core/common/rescode"
	"net/http"

	"God/core/service"

	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
}

var userService = &service.UserService{}

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
