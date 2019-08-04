package util

import (
	"math"
	"time"

	"God/utils/comutil"

	"God/core/common"

	"errors"

	"strings"

	"github.com/spf13/viper"
	redisErr "gopkg.in/redis.v4"
)

func CheckIpRate(preffix, ip string, duration int) bool {
	return SetNX(comutil.IP_LIMIT+":"+preffix, ip, duration)
}

// timestamp 校验
// 1、和当前时间不可相差超过 apiLimit 分钟
// 2、相同的入参只能在 apiDuration 秒钟范围内 反复访问接口
func CheckApiTimeStamp(prefix string, timestamp int) (int64, error) {
	now := time.Now().Unix()

	sub := math.Abs(float64(now) - float64(timestamp))
	div := sub / 60

	apiLimit := viper.GetFloat64("common.apiLimitAbs")

	// timestamp error range abs
	if apiLimit < div {
		common.Logger.Error("the timestamp overflow the range", "nowTimestamp",
			now, "inputTimestamp", timestamp, "SubAbs %d m", div, "apiLimit %d m", apiLimit)
		return now, errors.New("timestamp overflow")
	}
	// TODO
	//apiDuration := viper.GetInt("common.apiDuration")
	//
	//if SetNX(comutil.API_LIMIT, prefix, int64(apiDuration)) {
	//	common.Logger.Error("the timestamp exceeds accessibility limit")
	//	return comerr.NewBizErr("timestamp overflow")
	//}
	return 0, nil
}

func CheckTokenAndSession(token, sessionId, accountId string) (bool, string, error) {
	sessionId, err := GetKV(comutil.TOKEN, token)
	if nil != err && err.Error() != redisErr.Nil.Error() {
		return false, "", err
	} else if (nil != err && err.Error() == redisErr.Nil.Error()) || "" == strings.TrimSpace(sessionId) {
		// generate a new sessionId relate to curr session
		sessionId := RandStr()
		sessionLimit := viper.GetInt("common.sessionDuration")
		SetKV(comutil.SESSION, sessionId, accountId, sessionLimit)
		return true, sessionId, nil
	}

	if accountId, err := GetKV(comutil.SESSION, sessionId); nil != err && err.Error() != redisErr.Nil.Error() {
		return false, "", err
	} else if (nil != err && err.Error() == redisErr.Nil.Error()) || "" == strings.TrimSpace(accountId) {
		DelKV(comutil.TOKEN, token)
		DelKV(comutil.SESSION, sessionId)
		common.Logger.Error("the accountId is empty on session, need to  login")
		return false, "", nil
	} else {
		return true, sessionId, nil
	}
}
