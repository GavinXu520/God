package main

import (
	"God/core"
	"God/core/common"
)

func init() {
	common.SetupConfig()
	common.SetUpLogger()
	//common.SetupDB()
	//common.SetUpRedis()
	common.SetupTimer()
}

func main() {
	core.SetupServer()
}
