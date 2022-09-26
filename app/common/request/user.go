package request

// 注册验证结构体
type Register struct {
	Name     string `form:"name" json:"name" binding:"required"`
	Mobile   string `form:"mobile" json:"mobile" binding:"required,mobile"`
	Password string `form:"password" json:"password" binding:"required"`
}

// 自定义注册错误信息
func (register Register) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"name.required":     "用户名称不能为空",
		"mobile.required":   "手机号码不能为空",
		"mobile.mobile":     "手机号码格式不正确",
		"password.required": "用户密码不能为空",
	}
}

// 登录验证结构体
type Login struct {
	Mobile   string `form:"mobile" json:"mobile" binding:"required,mobile"`
	Password string `form:"password" json:"password" binding:"required"`
}

// 自定义登录错误信息
func (login Login) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"mobile.required":   "手机号码不能为空",
		"mobile.mobile":     "手机号码格式不正确",
		"password.required": "用户密码不能为空",
	}
}

// 毒鸡汤结构体
type ChickenSoup struct {
	Text   string `form:"text" json:"text" binding:"required"`
	TypeId string `form:"typeId" json:"typeId" binding:"required"`
}

// 自定义登录错误信息
func (chickenSoup ChickenSoup) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"typeId.required": "鸡汤类型不能为空",
		"text.required":   "鸡汤内容不能为空",
	}
}
