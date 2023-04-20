package main

import (
	"ewa_admin_server/core"
	"ewa_admin_server/global"
	"ewa_admin_server/initialize"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

const AppMode = "debug" // 运行环境，主要有三种：debug、test、release

func main() {
	gin.SetMode(AppMode)

	//	TODO：1.配置初始化
	global.EWA_VIPER = core.InitializeViper()

	//	TODO：2.日志
	global.EWA_LOG = core.InitializeZap()
	zap.ReplaceGlobals(global.EWA_LOG)

	global.EWA_LOG.Info("server run success on ", zap.String("zap_log", "zap_log"))

	//  TODO：3.数据库连接
	global.EWA_DB = initialize.Gorm()

	//	TODO：4.其他初始化
	initialize.OtherInit()

	//	TODO：5.启动服务
	core.RunServer()
}
