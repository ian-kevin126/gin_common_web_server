package initialize

import (
	"ewa_admin_server/global"
	"ewa_admin_server/initialize/internal"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// GormMysql 初始化Mysql数据库
func GormMysql() *gorm.DB {
	m := global.EWA_CONFIG.MySQL
	if m.Dbname == "" {
		return nil
	}

	// 创建 mysql.Config 实例，其中包含了连接数据库所需的信息，比如 DSN (数据源名称)，字符串类型字段的默认长度以及自动根据版本进行初始化等参数。
	mysqlConfig := mysql.Config{
		DSN:                       m.Dsn(), // DSN data source name
		DefaultStringSize:         191,     // string 类型字段的默认长度
		SkipInitializeWithVersion: false,   // 根据版本自动配置
	}

	// 打开数据库连接
	db, err := gorm.Open(mysql.New(mysqlConfig), internal.Gorm.Config(m.Prefix, m.Singular))

	// 将引擎设置为我们配置的引擎，并设置每个连接的最大空闲数和最大连接数。
	if err != nil {
		return nil
	} else {
		db.InstanceSet("gorm:table_options", "ENGINE="+m.Engine)
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)

		fmt.Println("====3-gorm====: gorm link mysql success")
		return db
	}
}
