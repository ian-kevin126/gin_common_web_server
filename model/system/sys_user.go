package system

import (
	"ewa_admin_server/utils"
)

type Register struct {
	Name     string `form:"name" json:"name" binding:"required"`
	Mobile   string `form:"mobile" json:"mobile" binding:"required,mobile"`
	Password string `form:"password" json:"password" binding:"required"`
}

// GetMessages 自定义错误信息
func (register Register) GetMessages() utils.ValidatorMessages {
	return utils.ValidatorMessages{
		"name.required":     "用户名称不能为空",
		"mobile.required":   "手机号码不能为空",
		"mobile.mobile":     "手机号码格式不正确",
		"password.required": "用户密码不能为空",
	}
}
