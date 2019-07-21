package main

import (
	"God/core"
	"God/core/common"

	_ "github.com/go-sql-driver/mysql" // must init
)

func init() {
	common.SetupConfig()
	common.SetUpLogger()
	common.SetupDB()
	common.SetUpRedis()
	common.SetupTimer()
}

func main() {
	// Start with: go run main.go -conf ./config/${fileName}.json
	core.SetupServer()
}
