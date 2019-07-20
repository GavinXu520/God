package api

import "github.com/gin-gonic/gin"

func SetUpApi(router *gin.Engine) {

	// SetUp user api
	setUserApi(router)
}
