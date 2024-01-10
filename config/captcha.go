package config

type Captcha struct {
	KeyLong            int `mapstructure:"key_long" json:"key_long" yaml:"key_long"`                                     // 验证码长度
	ImgWidth           int `mapstructure:"img_width" json:"img_width" yaml:"img_width"`                                  // 验证码宽度
	ImgHeight          int `mapstructure:"img_height" json:"img_height" yaml:"img_height"`                               // 验证码高度
	OpenCaptcha        int `mapstructure:"open_captcha" json:"open_captcha" yaml:"open_captcha"`                         // 防爆破验证码开启此数，0代表每次登录都需要验证码，其他数字代表错误密码此数，如3代表错误三次后出现验证码
	OpenCaptchaTimeOut int `mapstructure:"open_captcha_timeout" json:"open_captcha_timeout" yaml:"open_captcha_timeout"` // 防爆破验证码超时时间，单位：s(秒)
}
