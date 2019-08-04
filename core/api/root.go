package api

import "github.com/gin-gonic/gin"

func SetUpApi(router *gin.Engine) {

	root := router.Group("/api/v1")

	// SetUp user api
	setUserApi(root)
	// SetUp customer api
	setCustomerApi(root)
}
