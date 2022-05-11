package controllers

import (
	"com/denis/smarthome/app/helper"
	"com/denis/smarthome/app/logger"
	"com/denis/smarthome/app/service"
	"com/denis/smarthome/app/service/dto"
	"com/denis/smarthome/app/task"
	"com/denis/smarthome/app/utils"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"strings"
)

// ReceiveMsg 接收易快报出口消息
func ReceiveMsg(c *gin.Context) {
	helper.SentinelWrap(c, func(c *gin.Context) {
		req := dto.EkuaibaoRequest{}
		if err := json.NewDecoder(c.Request.Body).Decode(&req); err != nil {
			logger.Errorf("Error decoding request: %s", err)
			c.JSON(400, dto.BuildFailResp("request param error"))
			return
		}

		logger.Info(utils.ToString(req))

		// 只处理出差和外出单
		if !strings.Contains(req.Title, "出差") && !strings.Contains(req.Title, "外出") {
			c.JSON(200, dto.BuildResp())
			return
		}

		// 检查时间合法性
		oaService := service.NewOaService{}
		err := oaService.CheckTime(req.FlowId)
		if err != nil {
			c.JSON(400, dto.BuildFailResp(err.Error()))
			return
		}

		utils.Remove(req.FlowId)
		result := utils.Add(req.FlowId)

		if result {
			c.JSON(200, dto.BuildResp())
		} else {
			c.JSON(505, dto.BuildFailResp("system error"))
		}
	})
}

// AddReview 补单，由于系统问题而未添加的单
func AddReview(c *gin.Context) {
	helper.SentinelWrap(c, func(c *gin.Context) {
		flowId := c.Query("flowId")
		logger.Info("manual addReview flowId:", flowId)

		transformService := service.TransformService{}
		flowDetail := transformService.GetFlowDetails(flowId)
		if flowDetail.ErrorCode == 0 {
			_ = task.AddReview(flowDetail)
		}
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
