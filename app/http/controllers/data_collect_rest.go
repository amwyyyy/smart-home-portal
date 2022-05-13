package controllers

import (
	"com/denis/smarthome/app/helper"
	"com/denis/smarthome/app/service/dto"
	"github.com/gin-gonic/gin"
)

// CollectPhoto 收集摄像头图片
func CollectPhoto(c *gin.Context) {
	helper.SentinelWrap(c, func(c *gin.Context) {
		c.JSON(200, dto.BuildResp())
	})
}

// CollectEnvInfo 收集环境信息（温度湿度）
func CollectEnvInfo(c *gin.Context) {

}

// CollectAirInfo 收集空气信息
func CollectAirInfo(c *gin.Context) {

}

// CollectHumanInfo 接收人体活动信息
func CollectHumanInfo(c *gin.Context) {

}

// CollectLightInfo 收集光照强度信息
func CollectLightInfo(c *gin.Context) {

}
