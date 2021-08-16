package models

import "github.com/go-playground/validator/v10"

// 定义请求的参数结构体

// ParamSignUp 注册请求参数
type ParamSignUp struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}

// ParamLogin 登录请求参数
type ParamLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// ParamRefreshToken 刷新token请求参数
type ParamRefreshToken struct {
	AccessToken  string `json:"access_token" binding:"required"`
	RefreshToken string `json:"refresh_token" binding:"required"`
}

// ParamPost 创建帖子请求参数
type ParamPost struct {
	Id         int64  `json:"post_id"`
	AuthorId   int64  `json:"author_id"`
	CategoryId int64  `json:"category_id" binding:"required"`
	Status     int32  `json:"status"`
	Title      string `json:"title" binding:"required"`
	Context    string `json:"context" binding:"required"`
}

// SignUpParamStructLevelValidation 自定义SignUpParam结构体校验函数
func SignUpParamStructLevelValidation(sl validator.StructLevel) {
	su := sl.Current().Interface().(ParamSignUp)

	if su.Password != su.RePassword {
		// 输出错误提示信息，最后一个参数就是传递的param
		sl.ReportError(su.RePassword, "re_password", "RePassword", "eqfield", "password")
	}
}
