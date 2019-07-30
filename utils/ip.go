package util

import (
	_ "fmt"
	"strconv"
	"strings"

	"God/core/common"

	"github.com/gin-gonic/gin"
	"github.com/go-errors/errors"
)

func Inet_aton(ip string) (int, error) {
	ips := strings.Split(ip, ".")
	E := errors.New("Not A IP.")
	if len(ips) != 4 {
		return 0, E
	}
	var intIP int
	for k, v := range ips {
		i, err := strconv.Atoi(v)
		if err != nil || i > 255 {
			return 0, E
		}
		intIP = intIP | i<<uint(8*(3-k))
	}
	common.Logger.Info("ip:", ip, ",the IP int is:", intIP)
	return intIP, nil
}

func GetRealRemoteIp(ctx *gin.Context) string {
	ip := ctx.Request.Header.Get("x-forwarded-for")
	if "" != strings.TrimSpace(ip) && "unknown" != strings.ToLower(ip) && CheckIP(ip) {
		ipArr := strings.Split(ip, ",") // 多次反向代理后会有多个ip值，第一个ip才是真实ip
		return ipArr[0]
	}

	if "" == strings.TrimSpace(ip) || "unknown" == strings.ToLower(ip) {
		ip = ctx.Request.Header.Get("Proxy-Client-IP")
	}

	if "" == strings.TrimSpace(ip) || "unknown" == strings.ToLower(ip) {
		ip = ctx.Request.Header.Get("WL-Proxy-Client-IP")
	}

	if "" == strings.TrimSpace(ip) || "unknown" == strings.ToLower(ip) {
		ip = ctx.Request.Header.Get("HTTP_CLIENT_IP")
	}

	if "" == strings.TrimSpace(ip) || "unknown" == strings.ToLower(ip) {
		ip = ctx.Request.Header.Get("HTTP_X_FORWARDED_FOR")
	}

	if "" == strings.TrimSpace(ip) || "unknown" == strings.ToLower(ip) {
		ip = ctx.Request.Header.Get("X-Real-IP")
	}

	if "" == strings.TrimSpace(ip) || "unknown" == strings.ToLower(ip) {
		ip = ctx.Request.RemoteAddr
	}
	return ip
}
