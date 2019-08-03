package util

import (
	"math"
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
func CheckApiTimeStamp(prefix string, timestamp int) error {
	now := time.Now().Unix()

	sub := math.Abs(float64(now) - float64(timestamp))
	div := sub / 60

	apiLimit := viper.GetFloat64("common.apiLimitAbs")

	// timestamp error range abs
	if apiLimit < div {
		common.Logger.Error("the timestamp overflow the range", "nowTimestamp",
			now, "inputTimestamp", timestamp, "SubAbs %d m", div, "apiLimit %d m", apiLimit)
		return comerr.NewBizErr("timestamp overflow")
	}
	// TODO
	//apiDuration := viper.GetInt("common.apiDuration")
	//
	//if SetNX(comutil.API_LIMIT, prefix, int64(apiDuration)) {
	//	common.Logger.Error("the timestamp exceeds accessibility limit")
	//	return comerr.NewBizErr("timestamp overflow")
	//}
	return nil
}
