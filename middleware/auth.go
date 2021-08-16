package middleware

import (
	"go_web/web_app/controller"
	"go_web/web_app/dict"
	"go_web/web_app/pkg/jwt"
	"strings"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

func AuthUserToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
		// 这里假设Token放在Header的Authorization中，并使用Bearer开头
		authHeader := ctx.Request.Header.Get("Authorization")
		// 没有携带请求头
		if authHeader == "" {
			controller.ReturnErr(ctx, dict.CodeNeedLogin)
			ctx.Abort()
			return
		}
		// 提前token Bearer Token
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			zap.L().Error("请求头格式错误", zap.Any("parts", parts))
			controller.ReturnErr(ctx, dict.CodeInvalidToken)
			ctx.Abort()
			return
		}
		claims, err := jwt.ParseToken(parts[1])
		if err != nil {
			zap.L().Error("token验证出错", zap.Error(err))
			controller.ReturnErr(ctx, dict.CodeInvalidToken)
			ctx.Abort()
			return
		}
		ctx.Set(dict.ContextUserIdKey, claims.UserId)
		ctx.Next()
	}
}
