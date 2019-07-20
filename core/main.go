package core

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"God/core/api"

	"God/core/common"

	"github.com/gin-gonic/gin"
	"github.com/go-errors/errors"
	"github.com/spf13/viper"
)

func SetupServer() {

	banner := `
 _          __  _____   _       _____   _____       ___  ___   _____        _____   _____        _____   _____   _____        _____   _____   _____    _     _   _   _____   _____  
| |        / / | ____| | |     /  ___| /  _  \     /   |/   | | ____|      |_   _| /  _  \      /  ___| /  _  \ |  _  \      /  ___/ | ____| |  _  \  | |   / / | | /  ___| | ____| 
| |  __   / /  | |__   | |     | |     | | | |    / /|   /| | | |__          | |   | | | |      | |     | | | | | | | |      | |___  | |__   | |_| |  | |  / /  | | | |     | |__   
| | /  | / /   |  __|  | |     | |     | | | |   / / |__/ | | |  __|         | |   | | | |      | |  _  | | | | | | | |      \___  \ |  __|  |  _  /  | | / /   | | | |     |  __|  
| |/   |/ /    | |___  | |___  | |___  | |_| |  / /       | | | |___         | |   | |_| |      | |_| | | |_| | | |_| |       ___| | | |___  | | \ \  | |/ /    | | | |___  | |___  
|___/|___/     |_____| |_____| \_____| \_____/ /_/        |_| |_____|        |_|   \_____/      \_____/ \_____/ |_____/      /_____/ |_____| |_|  \_\ |___/     |_| \_____| |_____|  
		`

	port := viper.GetString("server.port")
	host := viper.GetString("server.host")

	//router := gin.Default()

	router := gin.New()

	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	//router.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	router.Use(gin.Recovery())

	// Add ErrorDeal middleware
	router.Use(gin.HandlerFunc(func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				msg := fmt.Sprintf("发生panic异常: %v\n", errors.Wrap(err, 2).ErrorStack())
				ctx.JSON(http.StatusInternalServerError, msg)
			}
		}()
		ctx.Next()
	}))

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, banner)
	})

	// todo SetUp All Api
	api.SetUpApi(router)

	srv := &http.Server{
		Addr:    host + ":" + port,
		Handler: router,
	}

	go func() {
		// server connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		} else {
			log.Printf("%s\n\n%s\n", "Good bye, see you next time !!", fmt.Sprintf("%s: Running at %s", time.Now().Format("Mon, 02 Jan 2006 15:04:05 GMT"), port))
		}
	}()

	// shutdown server
	shut_down(srv)
}

func shut_down(srv *http.Server) {
	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown God Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 6*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("God http Server Shutdown failed: ", err)
	}
	// catching ctx.Done(). timeout of 5 seconds.
	select {
	case <-ctx.Done():

		log.Println("Close all resources ...")

		if nil != common.DB {
			if err := common.DB.Close(); nil != err {
				log.Fatal("Failed to close db:", err)
			}
		}
		if nil != common.Redis {
			if err := common.Redis.Close(); nil != err {
				log.Fatal("Failed to close redis:", err)
			}
		}
		log.Println("Resources close success!!")
	}
	log.Println("Server success exit ...")
}
