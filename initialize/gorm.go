package initialize

import (
	"ewa_admin_server/global"
	"ewa_admin_server/model/system"
	"os"

	"go.uber.org/zap"

	"gorm.io/gorm"
)

// Gorm 初始化数据库并产生数据库全局变量
func Gorm() *gorm.DB {
	switch global.EWA_CONFIG.App.DbType {
	case "mysql":
		return GormMysql()
	case "pgsql":
		return GormPgSql()
	default:
		return GormMysql()
	}
}

// RegisterDBTables 注册数据库表
func RegisterDBTables(db *gorm.DB) {
	err := db.AutoMigrate(
		// 系统模块表
		system.JwtBlacklist{}, // JWT 黑名单表
	)

	if err != nil {
		global.EWA_LOG.Error("register table failed", zap.Error(err))
		os.Exit(0)
	}

	global.EWA_LOG.Info("register table success")
}
