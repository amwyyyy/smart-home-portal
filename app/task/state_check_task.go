package task

import (
	"com/denis/smarthome/app/logger"
	"github.com/robfig/cron/v3"
)

// Task 定时根据flowId获取单据状态
func Task() {
	c := cron.New(cron.WithSeconds())

	_, err := c.AddFunc("0 0/5 * * * ?", func() {
		logger.Info("task run!")

	})

	if err != nil {
		logger.Errorf("task run error:", err)
	}

	c.Start()
}
