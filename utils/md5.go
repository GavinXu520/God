package util

import (
	"God/core/common"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
)

func Md5(data string) string {
	hash := md5.New()        // init one MD5 instance
	hash.Write([]byte(data)) //  the data need to MD5
	cipherStr := hash.Sum(nil)
	common.Logger.Debug("Before MD5: ", data)
	common.Logger.Info("After MD5: ", hex.EncodeToString(cipherStr))
	return hex.EncodeToString(cipherStr)
}

func Base64(data string) string {
	common.Logger.Debug("Before base64: ", data)
	bytes := []byte(data)
	result := base64.StdEncoding.EncodeToString(bytes)
	common.Logger.Info("After base64: ", result)
	return result
}
