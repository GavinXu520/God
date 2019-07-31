package util

import (
	"math"
	"strconv"
	"time"

	"God/core/common/comerr"

	"God/utils/comutil"

	"God/core/common"

	"github.com/spf13/viper"
)

func CheckIpRate(preffix, ip string, duration int64) bool {
	return SetNX(comutil.IP_LIMIT+":"+preffix, ip, duration)
}

// timestamp 校验
// 1、和当前时间不可相差超过 apiLimit 分钟
// 2、相同的入参只能在 apiDuration 秒钟范围内 反复访问接口
func CheckApiTimeStamp(prefix, timestamp string) error {
	now := time.Now().Unix()
	then, err := strconv.Atoi(timestamp)
	if nil == err {
		return err
	}

	sub := math.Abs(float64(now) - float64(then))
	div := sub / 60

	apiLimit := viper.GetFloat64("common.apiLimitAbs")

	// timestamp error range abs
	if apiLimit < div {
		return comerr.NewBizErr("timestamp overflow")
	}
	apiDuration := viper.GetInt("common.apiDuration")

	if !SetNX(comutil.API_LIMIT, prefix, int64(apiDuration)) {
		common.Logger.Error("the timestamp exceeds accessibility limit")
		return comerr.NewBizErr("timestamp overflow")
	}
	return nil
}
