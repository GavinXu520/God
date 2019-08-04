package util

import (
	"God/core/common"
	"fmt"
	"time"

	"github.com/go-errors/errors"
	"github.com/spf13/viper"
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

func RateLimit(businessName, uniqueKey string, count, second int64, needAdd bool) (bool, error) {
	key := viper.GetString("rateLimit.key") + "_" + businessName + "_" + uniqueKey
	if key == "" {
		return false, errors.New("取不到reids中ratelimit")
	}
	//获取当前队列长度
	len := common.Redis.LLen(key).Val()
	//如果请求次数已经超过当前最大的请求次数
	if len >= count {
		firstTime, err := common.Redis.RPop(key).Int64()
		if err != nil {
			return false, errors.New("取不到reids中ratelimitTime")
		}
		//判断当前时间是否已经超过限制时间,恢复次数
		if time.Now().Unix() > firstTime+second {
			//队列push 将当前请求时间放入队列中
			if needAdd {
				common.Redis.LPush(key, time.Now().Unix())
				//设置过期时间
				common.Redis.Expire(key, time.Duration(second*int64(time.Second)))
			}
			return true, nil
		} else {
			common.Redis.RPush(key, firstTime)
			return false, nil
		}

	}

	//队列push 将当前请求时间放入队列中
	if needAdd {
		common.Redis.LPush(key, time.Now().Unix())
		common.Redis.Expire(key, time.Duration(second*int64(time.Second)))
	}

	return true, nil
}

/**
检验登录密码错误次数
返回值 含义 ： 是否保存成功  是否达到限制次数  错误示例
*/
func RateLimitUserPwdErr(businessName, uniqueKey string, count, second int64) (bool, bool, error) {

	key := viper.GetString("rateLimit.key") + "_" + businessName + "_" + uniqueKey
	//field :=  businessName + "_" + uniqueKey
	if key == "" {
		return false, false, errors.New("xxx")
	}
	//每一次都是先放入缓存
	intCmd := common.Redis.LPush(key, time.Now().Unix())

	limitNum, err := intCmd.Result()

	if nil != err {
		return false, false, errors.New("xxx")
	}

	//设置过期时间
	common.Redis.Expire(key, time.Duration(second*int64(time.Second)))

	//获取当前队列长度
	len := common.Redis.LLen(key).Val()
	fmt.Print("len:", len, ",limitNum:", limitNum)

	//如果请求次数已经超过当前最大的请求次数
	if len > count {
		firstTime, err := common.Redis.RPop(key).Int64()

		if err != nil {
			return false, false, errors.New("xxx")
		}
		//再把它压回去/不然每次都会是比 count小,影响到是否需要输入验证码的方法  IsCaptchaVerify
		if len == count {
			common.Redis.RPush(key, firstTime)
		}
		//判断当前时间是否已经超过限制时间,恢复次数 (单位: s)  time.Now().Unix()取的是 s
		if time.Now().Unix() > firstTime+second {
			//在规定时间内,未超过限制
			return true, false, nil
		} else {
			//否则 受到限制
			//common.Redis.RPush(key, firstTime)
			return true, true, nil
		}
	}

	//次数不满足, 未受到限制
	return true, false, nil
}

/**
delete the login pwd wrong num limit
*/
func DelLimitUserPwdErr(businessName, uniqueKey string) (bool, error) {
	key := viper.GetString("rateLimit.key") + "_" + businessName + "_" + uniqueKey
	if key == "" {
		return false, errors.New("xxx")
	}
	intCmd := common.Redis.Del(key)

	_, err := intCmd.Result()

	if nil != err {
		return false, errors.New("xxxx")
	}
	return true, nil
}

/**
check need to have imgCode
*/
func IsCaptchaVerify(businessName, uniqueKey string, count int64) (bool, error) {
	key := viper.GetString("rateLimit.key") + "_" + businessName + "_" + uniqueKey
	if key == "" {
		return false, errors.New("xxxx")
	}

	//获取当前队列长度
	len := common.Redis.LLen(key).Val()
	fmt.Print("len:", len)

	//如果请求次数已经超过当前最大的请求次数(需要输入图形验证码)
	if len >= count {
		return true, nil
	}
	return false, nil
}
