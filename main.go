package main

import (
	"God/core"
	"God/core/init"
)

func init() {
	init.SetupConfig()
	//init.SetErrorDeal()
	//init.SetUpLogger()
	init.SetupDB()
	init.SetUpRedis()
	init.SetupTimer()
}

func main() {
	core.SetupServer()
}
