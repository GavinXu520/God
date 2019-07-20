package init

import (
	"info-base/src/server/common"

	"time"

	"github.com/robfig/cron"
)

func SetupTimer() {
	c := cron.New()

	// health check per 30 minutes
	c.AddFunc("0 0/30 * * * *", func() {
		common.Logger.Info("The system health heartbeat check, time is : %s",
			time.Now().Format("2006-01-02 15:04:05"))
	})

	c.Start()
}
