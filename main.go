package main

import (
	"com/denis/smarthome/app"
)

func main() {
	// 初始化 gin 引擎
	engine := app.Init()

	// 启动服务
	app.Run(engine)
}
