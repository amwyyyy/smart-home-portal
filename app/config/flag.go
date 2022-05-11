package config

// 此处是处理默认参数
import (
	"com/denis/smarthome/app/helper"
	"os/exec"
	"path/filepath"
)

// Root 项目的根目录
var Root string

var configDir string

// 初始化配置文件
func init() {
	// 初始化所有命令行传入的参数
	flags()

	// 初始化 app 根目录
	root()

	initCfg()
}

func flags() {
	Root = ""
	configDir = ""

	// 单元测试需要把注释打开，以后优化
	//Root = "../../"
}

// 获取项目的根目录

// 项目的根目录，静态文件目录，模板加载目录等都会根据此目录来配置
// 会根据优先级从上到下获取，上面的有配置就会忽略下面的：
// 1. 启动时候带的参数，如： ./app --AppRoot=/www/blog
// 2. 当前可执行文件执行的路径，如: /www/web/blog/go-blog 此时的根目录就是 /www/web/blog
func root() {

	// 如果命令行设置了，使用命令行设置的
	// 如果没有设置，就是用当前可执行文件所在的目录
	if Root == "" {
		// 如果没有设置，就是用当前可执行文件所在的目录

		// 验证下文件是否存在，我们执行的就是自己，也不存在不存在的问题了
		file, _ := exec.LookPath("/")

		// 获取文件的绝对路径
		path, _ := filepath.Abs(file)

		Root = path
	}

	root, err := filepath.Abs(Root)
	helper.PanicErr(err)

	Root = root

}
