package app

import (
	"com/denis/smarthome/app/config"
	"com/denis/smarthome/app/http"
	"com/denis/smarthome/app/http/routes"
	"com/denis/smarthome/app/logger"
	"com/denis/smarthome/app/utils"
	sentinel "github.com/alibaba/sentinel-golang/api"
	"github.com/alibaba/sentinel-golang/core/flow"
	"github.com/alibaba/sentinel-golang/core/system"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"os"
)

var cfg = config.Cfg

// Init 初始化
func Init() *gin.Engine {
	// 初始化日志配置
	setDefaultConfig()

	// 初始化引擎
	engine := gin.New()

	// 注册全局的中间件
	engine.Use(gin.Logger())
	engine.Use(http.Recovery)

	// 配置路由
	routes.New(engine)

	// 初始化sentinel
	initSentinel(engine.Routes())

	// 启动定时任务
	//go task.Task()

	// 加载数据文件
	utils.Load()

	return engine
}

func initSentinel(routes gin.RoutesInfo) {
	if err := sentinel.InitWithConfigFile("config/sentinel.yml"); err != nil {
		logger.Error(err)
	}

	// 自适应流控
	_, err := system.LoadRules([]*system.Rule{
		{
			MetricType:   system.Load,
			TriggerCount: 2.0,
			Strategy:     system.BBR,
		},
	})
	if err != nil {
		logger.Errorf("Unexpected error: %+v", err)
		return
	}

	// initialize sentinel rules
	initRules(routes)
}

func initRules(routes gin.RoutesInfo) {
	rules := make([]*flow.Rule, 0)
	for _, r := range routes {
		rules = append(rules, &flow.Rule{
			Resource:               r.Path,
			TokenCalculateStrategy: flow.Direct,
			ControlBehavior:        flow.Reject,
			Threshold:              100,
			StatIntervalInMs:       1000,
		})
	}

	_, err := flow.LoadRules(rules)
	if err != nil {
		logger.Errorf("Unexpected error: %+v", err)
		return
	}
}

func setDefaultConfig() {
	if cfg.Debug {
		gin.SetMode(gin.DebugMode)

		return
	}

	gin.SetMode(gin.ReleaseMode)
	gin.DisableConsoleColor()

	if cfg.Log.Mode == logger.LogFile && cfg.Log.Access {
		gin.DefaultWriter = logger.NewAccessLog()
		return
	}

	gin.DefaultWriter = ioutil.Discard
}

// Run 启动服务
func Run(e *gin.Engine) {
	err := e.Run(config.Cfg.Addr())
	if err != nil {
		logger.Errorf(err.Error())
		os.Exit(-1)
	}
}
