package main

import (
	"God/core"
	"God/core/common"
)

func init() {
	common.SetupConfig()
	common.SetUpLogger()
	common.SetupDB()
	//common.SetUpRedis()
	common.SetupTimer()
}

func main() {
	// Start with: go run main.go -conf ./config/${fileName}.json
	core.SetupServer()
}
