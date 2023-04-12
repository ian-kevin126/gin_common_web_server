package main

import (
	"ewa_admin_server/core"
	"ewa_admin_server/global"
	"ewa_admin_server/initialize"
	"fmt"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

const AppMode = "debug" // 运行环境，主要有三种：debug、test、release

func main() {
	gin.SetMode(AppMode)

	//	1.配置初始化
	global.EWA_VIPER = core.InitializeViper()

	//	2.其他初始化
	initialize.OtherInit()

	//	3.日志
	global.EWA_LOG = core.InitializeZap()
	zap.ReplaceGlobals(global.EWA_LOG)

	global.EWA_LOG.Info("server run success on ", zap.String("zap_log", "zap_log"))

	//  4.数据库连接
	global.EWA_DB = initialize.Gorm()

	if global.EWA_DB != nil {
		initialize.RegisterDBTables(global.EWA_DB)
		fmt.Println("====init table success====")
		// 程序结束前关闭数据库链接
		db, _ := global.EWA_DB.DB()
		defer db.Close()
	}

	//	5.启动服务
	core.RunServer()
}
