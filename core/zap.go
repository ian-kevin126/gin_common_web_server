package core

import (
	"ewa_admin_server/core/internal"
	"ewa_admin_server/global"
	"ewa_admin_server/utils"
	"fmt"

	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// InitializeZap Zap 获取 zap.Logger
func InitializeZap() (logger *zap.Logger) {
	if ok, _ := utils.PathExists(global.EWA_CONFIG.Zap.Director); !ok {
		// 判断是否有Director文件夹
		fmt.Printf("create %v directory\n", global.EWA_CONFIG.Zap.Director)
		_ = os.Mkdir(global.EWA_CONFIG.Zap.Director, os.ModePerm)
	}

	cores := internal.Zap.GetZapCores()
	logger = zap.New(zapcore.NewTee(cores...))

	if global.EWA_CONFIG.Zap.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}

	fmt.Println("====2-zap====: zap log init success")
	return logger
}
