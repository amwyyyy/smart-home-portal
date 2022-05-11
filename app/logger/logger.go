package logger

import (
	"com/denis/smarthome/app/config"
	"com/denis/smarthome/app/helper"
	"fmt"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"os"
	"path"
)

const (
	// *******
	// 日志模式
	// *******

	// LogClose 关闭日志
	LogClose string = "close"
	// LogFile 写入到日志文件
	LogFile string = "file"
	// LogStdout 写入到标准输出
	LogStdout string = "stdout"
)

var cfg = config.Cfg
var root = config.Root

// LogInfo 普通日志输出
var LogInfo = NewLogInfo()

// LogErr 错误日志输出
var LogErr = NewLogError()

// Logger 日志结构
type Logger struct {
	*LogWriter
	*log.Logger
}

// NewLogInfo 初始化普通日志
func NewLogInfo() *Logger {
	l := Logger{}

	l.Logger = log.New(l.LogWriter, "[info] ", log.LstdFlags|log.Lshortfile)

	// 设置日志文件滚动
	l.SetOutput(&lumberjack.Logger{
		Filename:   path.Join(root, cfg.Log.Dir, "app.log"),
		MaxSize:    50, // megabytes
		MaxBackups: 3,
		MaxAge:     28,   //days
		Compress:   true, // disabled by default
	})

	return &l
}

// NewLogError 初始化错误日志
func NewLogError() *Logger {
	l := Logger{}

	l.Logger = log.New(l.LogWriter, "[error] ", log.LstdFlags|log.Lshortfile)

	// 设置日志文件滚动
	l.SetOutput(&lumberjack.Logger{
		Filename:   path.Join(root, cfg.Log.Dir, "app.log"),
		MaxSize:    50, // megabytes
		MaxBackups: 3,
		MaxAge:     28,   //days
		Compress:   true, // disabled by default
	})

	return &l
}

// LogWriter 定义日志输出的结构
type LogWriter struct {
	// 日志文件的名称
	fileName string

	// 系统日志打开的文件
	*os.File
}

// NewAccessLog 初始化访问日志
func NewAccessLog() *lumberjack.Logger {
	// 返回滚动的日志文件
	return &lumberjack.Logger{
		Filename:   path.Join(root, cfg.Log.Dir, "access.log"),
		MaxSize:    50, // megabytes
		MaxBackups: 3,
		MaxAge:     28,   //days
		Compress:   true, // disabled by default
	}
}

func (lw *LogWriter) setWriter() {
	if cfg.Debug {
		lw.File = os.Stdout
		return
	}

	if cfg.Log.Mode == "file" {
		lw.File = lw.newFile()
		return
	}

	if cfg.Log.Mode == "close" {
		lw.File = nil
		return
	}

}

// GetFileFullName 获取日志文件的名称
func (lw *LogWriter) GetFileFullName() string {
	return path.Join(root, cfg.Log.Dir, lw.fileName)
}

func (lw *LogWriter) newFile() *os.File {
	dir := path.Join(root, cfg.Log.Dir)
	fileStat, err := os.Stat(dir)

	if err != nil {
		if !os.IsNotExist(err) {
			helper.Panicf("%v", err)
		}

		err := os.MkdirAll(dir, 0755)
		if err != nil {
			return nil
		}
	} else {
		if !fileStat.IsDir() {
			helper.Panicf("%q is an existing file", dir)
		}
	}

	logFile := path.Join(dir, lw.fileName)

	fileStat, err = os.Stat(logFile)

	if err != nil {
		if !os.IsNotExist(err) {
			helper.Panicf("faild to open %q file, %v", err)
		}
	}

	file, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)

	helper.PanicErr(err)

	return file
}

// Info 打印 info 信息
func Info(vals ...interface{}) {
	err := LogInfo.Output(2, fmt.Sprintln(vals...))
	if err != nil {
		return
	}
}

// Infof 指定格式的 info 日志
func Infof(format string, vals ...interface{}) {
	err := LogInfo.Output(2, fmt.Sprintf(format, vals...))
	if err != nil {
		return
	}
}

// Error 错误日志
func Error(vals ...interface{}) {
	err := LogErr.Output(2, fmt.Sprintln(vals...))
	if err != nil {
		return
	}
}

// Errorf 指定格式的错误日志
func Errorf(format string, vals ...interface{}) {
	err := LogErr.Output(2, fmt.Sprintf(format, vals...))
	if err != nil {
		return
	}
}
