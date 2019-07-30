package util

import "github.com/gin-gonic/gin"

func GetValByHeader(ctx *gin.Context, key string) string {
	return ctx.Request.Header.Get(key)
}
