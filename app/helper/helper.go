package helper

import (
	"com/denis/smarthome/app/service/dto"
	sentinel "github.com/alibaba/sentinel-golang/api"
	"github.com/gin-gonic/gin"
	"runtime"
)

// Panicf 用来打印指定格式的 panic 信息，
// 需要注意，只能直接调用，间接调用文件的位置会定位错误（层级不准确）
func Panicf(format string, values ...interface{}) {
	if _, file, line, ok := runtime.Caller(1); ok {
		format += "\n\t in %d line of the %q file"
		values = append(values, line, file)
	}

	Panicf(format, values...)
}

// PanicErr 用来断言错误，如果存在错误就 panic
func PanicErr(err error) {
	if err != nil {
		var format string
		var values []interface{}
		if _, file, line, ok := runtime.Caller(1); ok {
			format = "\n\t in %d line of the %q file"
			values = append(values, line, file, err)
		}
		Panicf(format, values...)
	}
}

// SentinelWrap 接口访问流控
func SentinelWrap(c *gin.Context, f func(c *gin.Context)) {
	e, b := sentinel.Entry(c.Request.RequestURI)
	if b != nil {
		c.JSON(505, dto.BuildFailResp("访问频率过快"))
	} else {
		defer e.Exit()
		f(c)
	}
}
