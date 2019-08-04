package api

import (
	"God/core/controller"

	"github.com/gin-gonic/gin"
)

var customer = &controller.CustomerController{}

func setCustomerApi(rootApi *gin.RouterGroup) {

	api := rootApi.Group("/customer")

	// 立即借钱
	api.POST("/loan", customer.Loan)

}
