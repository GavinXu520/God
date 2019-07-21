package util

import (
	_ "fmt"
	"strconv"
	"strings"

	"God/core/common"

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
