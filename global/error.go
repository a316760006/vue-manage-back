package global

// 项目中可能存在的错误都统一存放到这里，为每一种类型错误都定义一个错误码，便于在开发过程快速定位错误，前端也可以根据不同错误码实现不同逻辑的页面交互
type CustomError struct {
	ErrorCode int
	ErrorMsg  string
}

type CustomErrors struct {
	BusinessError CustomError
	ValidateError CustomError
	TokenError    CustomError
}

var Errors = CustomErrors{
	BusinessError: CustomError{40000, "业务错误"},
	ValidateError: CustomError{42200, "请求参数错误"},
	TokenError:    CustomError{40100, "登录授权失效"},
}
