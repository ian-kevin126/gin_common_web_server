package global

import (
	"ewa_admin_server/config"

	"github.com/spf13/viper"
)

var (
	EWA_CONFIG config.Configuration
	EWA_VIPER  *viper.Viper
)
