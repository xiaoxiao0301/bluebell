package services

import "go_web/web_app/pkg/jwt"

// TokenService 认证服务层
type TokenService struct {
}

// RefreshToken 更新用户的access_token
func (token *TokenService) RefreshToken(o_access_token, o_refresh_token string) (access_token, refresh_token string, err error) {
	return jwt.RefreshToken(o_access_token, o_refresh_token)
}
