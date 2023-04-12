package system

import (
	"ewa_admin_server/global"
	"ewa_admin_server/model/system"

	"go.uber.org/zap"
)

/*
将数据库中存储的 JWT（JSON Web Token）黑名单数据加载到内存缓存 global.BlackCache 中。JWT 黑名单是指一些已失效的 JWT，
这些无效的 JWT 因为其包含的信息已经过期、被修改等原因导致不能再使用。通过将 JWT 黑名单数据加载到内存中，可以有效提高 JWT 验证的速度，
并减少频繁地访问数据库的开销。同时也保证了 JWT 验证的准确性和安全性。
*/

func LoadAll() {
	var data []string

	// 查询系统的 JWT 黑名单列表，将结果存储到 data 中。
	err := global.EWA_DB.Model(&system.JwtBlacklist{}).Select("jwt").Find(&data).Error

	if err != nil {
		global.EWA_LOG.Error("加载数据库jwt黑名单失败!", zap.Error(err))
		return
	}

	// jwt黑名单：遍历 jwt 列表，将每个 JWT 数据加入到内存缓存 global.BlackCache 中，键名是 JWT 值，键值为空结构体。
	for i := 0; i < len(data); i++ {
		global.BlackCache.SetDefault(data[i], struct{}{})
	}
}
