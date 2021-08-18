package jwt

import (
	"go_web/web_app/dict"
	"time"

	"github.com/golang-jwt/jwt"
)

// https://pkg.go.dev/github.com/golang-jwt/jwt#pkg-overview
// https://github.com/dgrijalva/jwt-go 这个是原地址

var JWT_TOKEN_SIGN_KEY = []byte("MIICXAIBAAKBgQCTdAu6xVgOz431opYE7K/DJFHiBho93NLgXFqw4vWoKc1ApbBV")

const (
	JWT_ACCESS_TOEKN_EXPIRED_DURATION  = 2 * time.Hour
	JWT_REFRESH_TOEKN_EXPIRED_DURATION = 30 * 24 * time.Hour
)

type TokenAuthClaims struct {
	UserId   int64  `json:"user_id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

// CreateToken 生成access_token与refresh_token
func CreateToken(userId int64, username string) (accessToken, refreshToken string, err error) {
	tc := TokenAuthClaims{
		UserId:   userId,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(JWT_ACCESS_TOEKN_EXPIRED_DURATION).Unix(),
			Issuer:    "bluebell",
		},
	}

	accessTokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, tc)
	accessToken, err = accessTokenClaims.SignedString(JWT_TOKEN_SIGN_KEY)

	refreshTokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(JWT_REFRESH_TOEKN_EXPIRED_DURATION).Unix(),
		Issuer:    "bluebell",
	})
	refreshToken, err = refreshTokenClaims.SignedString(JWT_TOKEN_SIGN_KEY)
	return
}

// ParseToken 解析access_token
func ParseToken(tokenStr string) (tc *TokenAuthClaims, err error) {
	tc = new(TokenAuthClaims)
	tokenClaims, err := jwt.ParseWithClaims(tokenStr, tc, keyFunc)
	if err != nil {
		return nil, err
	}

	tc, ok := tokenClaims.Claims.(*TokenAuthClaims)
	if ok && tokenClaims.Valid {
		return tc, nil
	}
	return
}

// RefreshToken 利用refresh_token重新生成新的accessToken和refreshToken
func RefreshToken(accessToken, refreshToken string) (newAccessToken, newRefreshToken string, err error) {
	// 判断refresh_token是否有效
	_, err = jwt.Parse(refreshToken, keyFunc)
	if err != nil {
		return
	}
	// 从旧的access_token中解析出claims数据
	var authClaims *TokenAuthClaims
	authClaims = new(TokenAuthClaims)
	_, err = jwt.ParseWithClaims(accessToken, authClaims, keyFunc)
	if err != nil {
		v, _ := err.(*jwt.ValidationError)
		// access_token是过期错误且refresh_token没有过期就新创建一个access_token和refresh_token
		if v.Errors == jwt.ValidationErrorExpired {
			return CreateToken(authClaims.UserId, authClaims.Username)
		}
		return
	} else {
		return "", "", dict.ErrorAccessTokenValid
	}

}

func keyFunc(token *jwt.Token) (interface{}, error) {
	return JWT_TOKEN_SIGN_KEY, nil
}
