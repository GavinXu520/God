package util

import (
	"God/core/common"
	"time"
)

func SetNX(businessName, uniqueKey string, second int64) bool {
	key := businessName + ":" + uniqueKey
	rs := common.Redis.SetNX(key, "", time.Duration(second)*time.Second)
	if rs.Val() == true {
		return true
	}
	return false
}
