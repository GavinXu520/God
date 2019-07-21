package util

import (
	"God/core/common"
	"encoding/json"
	"fmt"
)

func Interface2Json(argsName string, args interface{}) {
	jsonByte, err := json.Marshal(args)
	if nil != err {
		panic("Failed to parse json, the args is: " + fmt.Sprint(args))
	}
	common.Logger.Info(argsName + ":=" + string(jsonByte) + " \n")
}

func Interface2JsonStr(args interface{}) (string, error) {
	jsonByte, err := json.Marshal(args)
	if nil != err {
		return "", err
	}
	return string(jsonByte), nil
}
