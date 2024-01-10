package initialize

import (
	"ewa_admin_server/global"
	"ewa_admin_server/initialize/internal"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// GormPgSql 初始化 Postgresql 数据库
func GormPgSql() *gorm.DB {
	p := global.EWA_CONFIG.Pgsql

	if p.Dbname == "" {
		return nil
	}
	pgsqlConfig := postgres.Config{
		DSN:                  p.Dsn(), // DSN data source name
		PreferSimpleProtocol: false,
	}

	db, err := gorm.Open(postgres.New(pgsqlConfig), internal.Gorm.Config(p.Prefix, p.Singular))

	if err != nil {
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(p.MaxIdleConns)
		sqlDB.SetMaxOpenConns(p.MaxOpenConns)

		fmt.Println("====3-gorm====: gorm link PostgreSQL success")

		return db
	}
}
