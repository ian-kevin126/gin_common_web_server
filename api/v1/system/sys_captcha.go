package system

import (
	"ewa_admin_server/global"
	"ewa_admin_server/model/common/response"
	systemRes "ewa_admin_server/model/system/response"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
)

// 当开启多服务器部署时，替换下面的配置，使用redis共享存储验证码
// var store = captcha.NewDefaultRedisStore()
var store = base64Captcha.DefaultMemStore

type BaseApi struct{}

func (b *BaseApi) Captcha(c *gin.Context) {
	// 判断验证码是否开启，简单防爆限制的实现。它的作用是判断当前请求是否超过了允许的阈值，以此来判断请求是否合法。其主要思路是通过缓存记录
	// 用户的请求次数，并在超过阈值后拒绝请求。如果开启了防爆，则将客户端 IP 存入缓存，在缓存中进行相应的统计和设置过期时间。

	openCaptcha := global.EWA_CONFIG.Captcha.OpenCaptcha // 是否开启防爆次数

	// 定义了一个变量openCaptchaTimeOut，该变量表示缓存的超时时间（单位为秒），即当一个 IP 被加入到缓存后多长时间失效。
	openCaptchaTimeOut := global.EWA_CONFIG.Captcha.OpenCaptchaTimeOut // 缓存超时时间

	/**
	ClientIP 实现了一种算法来返回真实的客户端 IP。它在后台调用 c.RemoteIP() 来检查远程 IP 是否是受信任的代理。
	如果是，它将尝试解析 Engine.RemoteIPHeaders 中定义的标头（默认为 [X-Forwarded-For, X-Real-Ip]）。如果标头在语法上
	无效或远程 IP 不对应于受信任的代理，则返回远程 IP（来自 Request.RemoteAddr）。
	*/
	// 定义了一个变量key，该变量从当前上下文对象c中获取客户端的 IP 地址并命名为 key。
	key := c.ClientIP()

	// 将之前设置的 key 传入，Get 方法会尝试从缓存中读取对应 key 的数据。
	v, ok := global.BlackCache.Get(key)

	if !ok {
		// 如果不存在，则说明这是第一次请求，将 key 加入到缓存中，并设置其有效期为 time.Duration(openCaptchaTimeOut) 秒。
		global.BlackCache.Set(key, 1, time.Second*time.Duration(openCaptchaTimeOut))
	}

	// 是否需要进行验证码验证。如果需要，oc 变量将被设置为 true，否则为 false。
	var oc bool
	// 配置文件中开启防爆次数的值openCaptcha，如果该值为 0 或者小于当前请求次数，则 oc 将被设置为 true。
	if openCaptcha == 0 || openCaptcha < interfaceToInt(v) {
		oc = true
	}

	// 声明了一个默认数字验证码的 driver（可以配置字符、公式、验证码），其中包括验证码图片高度、宽度、验证码字符数、噪声复杂度和随机因子等参数。
	driver := base64Captcha.NewDriverDigit(
		global.EWA_CONFIG.Captcha.ImgHeight,
		global.EWA_CONFIG.Captcha.ImgWidth,
		global.EWA_CONFIG.Captcha.KeyLong,
		0.7,
		80,
	)

	// 根据 driver 和 store（内存）创建了一个验证码实例 cp。
	// cp := base64Captcha.NewCaptcha(driver, store.UseWithCtx(c))   // v8下使用redis
	cp := base64Captcha.NewCaptcha(driver, store)

	id, b64s, err := cp.Generate()

	if err != nil {
		global.EWA_LOG.Error("验证码获取失败!", zap.Error(err))
		response.FailWithMessage("验证码获取失败", c)
		return
	}

	// 将验证码信息返回给客户端，其中包括验证码 id、base64 编码后的图片字符串、验证码长度和验证码是否需要验证等信息。
	response.OkWithDetailed(systemRes.SysCaptchaResponse{
		CaptchaId:     id,
		PicPath:       b64s,
		CaptchaLength: global.EWA_CONFIG.Captcha.KeyLong,
		OpenCaptcha:   oc,
	}, "验证码获取成功", c)
}

// 类型转换
func interfaceToInt(v interface{}) (i int) {
	switch v := v.(type) {
	case int:
		i = v
	default:
		i = 0
	}

	return
}
