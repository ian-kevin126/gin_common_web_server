package global

import (
	"ewa_admin_server/config"

	"github.com/go-redis/redis/v8"

	"github.com/songzhibin97/gkit/cache/local_cache"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	EWA_CONFIG config.Configuration
	EWA_VIPER  *viper.Viper
	EWA_LOG    *zap.Logger
	EWA_DB     *gorm.DB
	EWA_REDIS  *redis.Client

	// BlackCache 本地缓存，用于将一些经常访问但不容易改动的数据存储到本地，减少网络请求次数提高性能。
	BlackCache local_cache.Cache
)
