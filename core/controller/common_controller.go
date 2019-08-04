package controller

import "God/core/service"

func checkSmsCode(mobileNo, smsCode string) (bool, error) {
	return service.CheckSmsCode(mobileNo, smsCode)
}
