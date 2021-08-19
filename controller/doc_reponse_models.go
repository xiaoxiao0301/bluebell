package controller

import "go_web/web_app/dict"

// _ResponseCommon 通用返回的数据结构
type _ResponseCommon struct {
	Code    dict.ResponseCode `json:"code"`    // 业务响应状态码
	Message string            `json:"message"` //提示信息
	Data    string            `json:"data"`    // 数据
}

// _ResponseLoginSuccess 登陆成功返回数据结构
type _ResponseLoginSuccess struct {
	Code         dict.ResponseCode `json:"code"`          // 业务响应状态码
	Message      string            `json:"message"`       //提示信息
	UserId       string            `json:"user_id"`       // 登陆用户ID
	Username     string            `json:"username"`      // 登陆用户昵称
	AccessToken  string            `json:"access_token"`  // jwt access_token 验证使用
	RefreshToken string            `json:"refresh_token"` // jwt refresh_token 刷新token
}

// _ResponseRefreshToken 刷新token返回数据结构
type _ResponseRefreshToken struct {
	Code         dict.ResponseCode `json:"code"`          // 业务响应状态码
	Message      string            `json:"message"`       //提示信息
	AccessToken  string            `json:"access_token"`  // jwt access_token 验证使用
	RefreshToken string            `json:"refresh_token"` // jwt refresh_token 刷新token
}
