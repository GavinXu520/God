package common

import (
	"database/sql"

	"github.com/spf13/viper"
)

func SetupDB() {
	// db config
	dialect := viper.GetString("db.dialect")
	user := viper.GetString("db.user")
	password := viper.GetString("db.password")
	database := viper.GetString("db.database")
	host := viper.GetString("db.host")
	port := viper.GetString("db.port")
	maxIdle := viper.GetInt("db.maxIdle")
	maxOpen := viper.GetInt("db.maxOpen")

	url := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + database + "?charset=utf8&parseTime=True&loc=Local"

	db, err := sql.Open(dialect, url)
	if err != nil {
		panic("Failed to connect database")
	}

	db.SetMaxIdleConns(maxIdle)
	db.SetMaxOpenConns(maxOpen)
	connectErr := db.Ping() // test db connect
	if connectErr != nil {
		panic("Failed to test connect database, err: " + connectErr.Error())
	}

	// read and write is only one instance
	DB = db
}

var (
	DB *sql.DB
)
