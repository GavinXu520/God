package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func setUserApi(router *gin.Engine) {

	router.GET("/welcome", func(ctx *gin.Context) {
		firstname := ctx.DefaultQuery("firstname", "Guest")
		lastname := ctx.Query("lastname")

		ctx.String(http.StatusOK, "Hello %s %s", firstname, lastname)
	})

}
