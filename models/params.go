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

// ParamCategory 创建分类请求参数
type ParamCategory struct {
	CategoryName string `json:"category_name" binding:"required"`
	Introduction string `json:"introduction" binding:"required"`
}

// ParamPost 创建帖子请求参数
type ParamPost struct {
	Id         int64  `json:"post_id"`
	AuthorId   int64  `json:"author_id"`
	CategoryId int64  `json:"category_id,string" binding:"required"`
	Status     int32  `json:"status"`
	Title      string `json:"title" binding:"required"`
	Content    string `json:"content" binding:"required"`
}

// ParamPage 列表分页请求参数
type ParamPage struct {
	Page int `json:"page" form:"page" binding:"required"` // 当前页码
	Size int `json:"size" form:"size" binding:"required"` // 每页个数
}

// ParamVote 帖子投票请求参数
type ParamVote struct {
	PostId int64 `json:"post_id,string" binding:"required"`            // 投票的帖子id
	Value  *int8 `json:"value,string" binding:"required,oneof=-1 0 1"` // 投票结果， 1 赞成 0 取消 -1 反对
}

// ParamNewPost 新版帖子列表接口
type ParamNewPostList struct {
	Order string `form:"order" binding:"required,oneof=time score"` // 按照时间或者分数排序
	Sorts string `form:"sorts" binding:"required,oneof=asc desc"`   // 升序还是降序 asc 升序 desc 降序
	Page  int    `form:"page" form:"page" binding:"required"`       // 当前页码
	Size  int    `form:"size" form:"size" binding:"required"`       // 每页个数
}

// SignUpParamStructLevelValidation 自定义SignUpParam结构体校验函数
func SignUpParamStructLevelValidation(sl validator.StructLevel) {
	su := sl.Current().Interface().(ParamSignUp)

	if su.Password != su.RePassword {
		// 输出错误提示信息，最后一个参数就是传递的param
		sl.ReportError(su.RePassword, "re_password", "RePassword", "eqfield", "password")
	}
}
