package internal

import (
	"ewa_admin_server/global"
	"fmt"

	"gorm.io/gorm/logger"
)

type writer struct {
	logger.Writer
}

// NewWriter writer 构造函数
func NewWriter(w logger.Writer) *writer {
	return &writer{Writer: w}
}

// Printf 格式化打印日志
func (w *writer) Printf(message string, data ...interface{}) {
	var logZap bool
	switch global.EWA_CONFIG.App.DbType {
	case "mysql":
		logZap = global.EWA_CONFIG.MySQL.LogZap
	}
	if logZap {
		global.EWA_LOG.Info(fmt.Sprintf(message+"\n", data...))
	} else {
		w.Writer.Printf(message, data...)
	}
}
