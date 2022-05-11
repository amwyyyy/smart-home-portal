package http

import (
	"com/denis/smarthome/app/logger"
	"com/denis/smarthome/app/service/dto"
	"net/http/httputil"

	"github.com/gin-gonic/gin"
)

// Recovery 用于页面出现 panic 时候的恢复
func Recovery(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			httprequest, _ := httputil.DumpRequest(c.Request, false)
			logger.LogErr.Printf("\n[Recovery] %s panic recovered: \n%s", string(httprequest), err)

			c.JSON(500, dto.BuildFailResp("server error"))
			return
		}
	}()
	c.Next()
}
