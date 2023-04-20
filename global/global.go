package global

import (
	"ewa_admin_server/config"

	"github.com/go-redis/redis/v8"

	"go.uber.org/zap"
	"gorm.io/gorm"

	"github.com/spf13/viper"
)

var (
	EWA_CONFIG config.Configuration
	EWA_VIPER  *viper.Viper
	EWA_LOG    *zap.Logger
	EWA_DB     *gorm.DB
	EWA_REDIS  *redis.Client
)
