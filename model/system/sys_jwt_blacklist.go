package system

import "ewa_admin_server/model/common"

type JwtBlacklist struct {
	common.EWA_MODEL
	Jwt string `gorm:"type:text;comment:jwt"`
}
