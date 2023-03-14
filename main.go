package main

import (
	"ewa_admin_server/core"
	"ewa_admin_server/global"
	"fmt"
)

func main() {
	//	TODO：1.配置初始化
	global.EWA_VIPER = core.InitializeViper()
	fmt.Println("====app_name====: ", global.EWA_CONFIG.App.AppName)

	//	TODO：2.日志

	//  TODO：3.数据库连接

	//	TODO：4.其他初始化

	//	TODO：5.启动服务
	core.RunServer()
}
