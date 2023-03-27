package internal

import (
	"ewa_admin_server/global"
	"log"
	"os"
	"time"

	"gorm.io/gorm/schema"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DBBASE interface {
	GetLogMode() string
}

var Gorm = new(_gorm)

type _gorm struct{}

// Config gorm 自定义配置
func (g *_gorm) Config(prefix string, singular bool) *gorm.Config {

	// 将传入的字符串前缀和单复数形式参数应用到 GORM 的命名策略中，并禁用迁移过程中的外键约束，返回最终生成的 GORM 配置信息。
	config := &gorm.Config{
		// 命名策略
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   prefix,   // 表前缀，在表名前添加前缀，如添加用户模块的表前缀 user_
			SingularTable: singular, // 是否使用单数形式的表名，如果设置为 true，那么 User 模型会对应 users 表
		},
		// 是否在迁移时禁用外键约束，默认为 false，表示会根据模型之间的关联自动生成外键约束语句
		DisableForeignKeyConstraintWhenMigrating: true,
	}

	_default := logger.New(NewWriter(log.New(os.Stdout, "\r\n", log.LstdFlags)), logger.Config{
		SlowThreshold: 200 * time.Millisecond,
		LogLevel:      logger.Warn,
		Colorful:      true,
	})

	var logMode DBBASE
	switch global.EWA_CONFIG.App.DbType {
	case "mysql":
		logMode = &global.EWA_CONFIG.MySQL
	default:
		logMode = &global.EWA_CONFIG.MySQL
	}

	switch logMode.GetLogMode() {
	case "silent", "Silent":
		config.Logger = _default.LogMode(logger.Silent)
	case "error", "Error":
		config.Logger = _default.LogMode(logger.Error)
	case "warn", "Warn":
		config.Logger = _default.LogMode(logger.Warn)
	case "info", "Info":
		config.Logger = _default.LogMode(logger.Info)
	default:
		config.Logger = _default.LogMode(logger.Info)
	}

	return config
}
