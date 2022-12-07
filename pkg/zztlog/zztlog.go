package zztlog

import (
	"github.com/zztroot/zztlog"
)

var logger *zztlog.InitLog

func init() {
	//通过文件配置
	if err := zztlog.InitConfigFile("./conf/zztlog.json"); err != nil {
		logger.Error(err)
		// return
	}
	logger.Info("🟢 zztlog initialized")
}

func Debug(message ...interface{}) {
	logger.Debug(message)
}

func Info(message ...interface{}) {
	logger.Info(message)
}

func Warn(message ...interface{}) {
	logger.Warn(message)
}

func Error(message ...interface{}) {
	logger.Error(message)
}
