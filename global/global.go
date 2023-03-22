package global

import (
	"ewa_admin_server/config"

	"go.uber.org/zap"

	"github.com/spf13/viper"
)

var (
	EWA_CONFIG config.Configuration
	EWA_VIPER  *viper.Viper
	EWA_LOG    *zap.Logger
)
