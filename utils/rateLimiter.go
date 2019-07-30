package util

import (
	"God/core/common"
	"time"
)

func SetNX(businessName, uniqueKey string, second int64) bool {
	key := businessName + ":" + uniqueKey
	rs := common.Redis.SetNX(key, "", time.Duration(int64(time.Second)*second))
	if rs.Val() == true {
		return true
	}
	return false
}
