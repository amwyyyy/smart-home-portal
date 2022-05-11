package config

import (
	"log"
	"reflect"
	"testing"
)

func TestNewConfig(t *testing.T) {
	// 手动初始化一个配置，需要从 config/config.toml 中获取，要保证一直
	cfg := manuallyGenerateConfiguration()

	log.Println("toml: ", Cfg)

	if !reflect.DeepEqual(cfg, Cfg) {
		t.Errorf("\nwant: \n%v \nhave: \n%v", cfg, Cfg)
	}
}

func manuallyGenerateConfiguration() Config {
	return Config{
		// 基本信息
		Name:  "ibox",
		URL:   "http://localhost",
		Host:  "",
		Port:  8080,
		Debug: true,

		// 日志
		Log: Log{
			Dir:    "logs",
			Mode:   "file",
			Access: true,
		},
	}
}
