package initialize

import (
	"ewa_admin_server/global"
	"ewa_admin_server/utils"
	"fmt"
	"reflect"
	"strings"

	"github.com/songzhibin97/gkit/cache/local_cache"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func OtherInit() {
	initializeValidator()
	initializeJWTAndBlackList()
	fmt.Println(" ===== Other init ===== ")
}

/*
这段代码的作用是从配置文件中获取 JWT 的过期时间和缓冲时间，解析为 Duration 类型并存储到内存缓存 `global.BlackCache` 中。具体实现如下：

- 使用自定义工具包 utils 中的 `ParseDuration()` 方法将读取到的全局变量 `global.EWA_CONFIG.JWT.ExpiresTime` 按照固定
格式解析成 time.Duration 类型的变量 dr，表示 JWT 有效期。如果解析失败，则会触发 panic()，即程序终止。

- 同样使用 `ParseDuration()` 方法将读取到的全局变量 `global.EWA_CONFIG.JWT.BufferTime` 按照固定格式解析成
time.Duration 类型的变量，表示 JWT 缓冲期。如果解析失败，则同样会触发 panic()。

- 调用 local_cache 包中的 `NewCache()` 方法创建一个新的缓存对象，并设置了其默认过期时间为 JWT 有效期 `dr`。

通过将 JWT 的过期时间和缓冲时间解析为 time.Duration 后存储到内存中，可以避免每次访问数据库获取这些配置数据，提高了应用程序的性能和效率。
同时也保证了 JWT 验证的准确性和安全性。
*/
func initializeJWTAndBlackList() {
	dr, err := utils.ParseDuration(global.EWA_CONFIG.JWT.ExpiresTime)
	if err != nil {
		panic(err)
	}
	_, err = utils.ParseDuration(global.EWA_CONFIG.JWT.BufferTime)
	if err != nil {
		panic(err)
	}

	global.BlackCache = local_cache.NewCache(
		local_cache.SetDefaultExpire(dr),
	)
}

func initializeValidator() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// 注册自定义验证器
		_ = v.RegisterValidation("mobile", utils.ValidateMobile)

		// 注册自定义 json tag 函数
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})
	}
}
