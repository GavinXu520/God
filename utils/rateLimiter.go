package util

import (
	"God/core/common"
	"time"
)

func SetNX(businessName, uniqueKey string, second int) bool {
	key := businessName + ":" + uniqueKey
	rs := common.Redis.SetNX(key, "", time.Duration(second)*time.Second)
	if rs.Val() == true {
		return true
	}
	return false
}

func SetKV(businessName, key, val string, second int) bool {
	key = businessName + ":" + key
	rs := common.Redis.SetNX(key, val, time.Duration(second)*time.Second)
	if rs.Val() == true {
		return true
	}
	return false
}
