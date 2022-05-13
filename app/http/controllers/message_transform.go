package controllers

import (
	"com/denis/smarthome/app/helper"
	"com/denis/smarthome/app/logger"
	"com/denis/smarthome/app/service/dto"
	"com/denis/smarthome/app/utils"
	"github.com/gin-gonic/gin"
)

// ReceiveMsg 接收易快报出口消息
func ReceiveMsg(c *gin.Context) {
	helper.SentinelWrap(c, func(c *gin.Context) {

		c.JSON(200, dto.BuildResp())
	})
}

// DelFlowId 用于删除异常单据
func DelFlowId(c *gin.Context) {
	helper.SentinelWrap(c, func(c *gin.Context) {
		flowId := c.Query("flowId")
		logger.Info("manual DelFlowId flowId:", flowId)

		utils.Remove(flowId)
	})
}
