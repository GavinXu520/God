package init

import (
	"flag"
	"log"

	"github.com/spf13/viper"
)

func SetupConfig() {
	//go run main.go -conf ./config/default.json
	var (
		conf = flag.String("conf", "", "eg:go run main.go -conf ./config/default.json")
	)
	if *conf == "" {
		*conf = "./config/local"
	}

	flag.Parse()

	log.Printf("The config file is %v", *conf)
	viper.SetConfigFile(*conf)
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Failed to read config fail, err: %v", err)
	}
}
