package util

import (
	"God/core/common"
	"time"
)

func GetKV(businessName, key string) (string, error) {
	key = businessName + ":" + key
	if val, err := common.Redis.Get(key).Result(); nil != err {
		return "", err
	} else {
		return val, nil
	}
}

func ExpireKV(businessName, key string, second int) (bool, error) {
	key = businessName + ":" + key
	if ok, err := common.Redis.Expire(key, time.Duration(second)*time.Second).Result(); nil != err {
		return false, err
	} else {
		return ok, nil
	}
}

func DelKV(businessName, key string) (bool, error) {
	key = businessName + ":" + key
	if _, err := common.Redis.Del(key).Result(); nil != err {
		return false, err
	} else {
		return true, nil
	}
}
