package utils

import (
	"com/denis/smarthome/app/helper"
	"com/denis/smarthome/app/logger"
	"github.com/patrickmn/go-cache"
	"os"
	"time"
)

var fileName = "store.log"

var c = cache.New(cache.NoExpiration, cache.NoExpiration)

// Load 从磁盘中加载数据
func Load() {

	if fileExist(fileName) {
		err := c.LoadFile(fileName)
		if err != nil {
			helper.PanicErr(err)
		}
	}

	logger.Info("load flowId stroe success.")

	go Flush()
}

// Add 添加数据
func Add(flowId string) bool {
	err := c.Add(flowId, 0, cache.NoExpiration)
	if err != nil {
		logger.Errorf("Add flowId err: %s", err.Error())
		return false
	}

	return true
}

// Remove 移除数据
func Remove(flowId string) {
	c.Delete(flowId)
}

// Flush 刷盘
func Flush() {
	logger.Info("start flowId stroe...")

	for true {
		err := c.SaveFile(fileName)
		if err != nil {
			logger.Errorf("flush flowId err: %s", err.Error())
			continue
		}

		time.Sleep(time.Duration(3) * time.Second)

	}
}

func fileExist(path string) bool {
	_, err := os.Lstat(path)
	return !os.IsNotExist(err)
}

func List() []string {
	var items []string
	for k := range c.Items() {
		items = append(items, k)
	}
	return items
}
