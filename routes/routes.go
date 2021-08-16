package routes

import (
	"go_web/web_app/controller"
	"go_web/web_app/logger"
	"go_web/web_app/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetUp() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	// 注册
	r.POST("signup", controller.SignUpHandler)

	// 登录
	r.POST("login", controller.LoginHandler)

	// 刷新token
	r.POST("refresh", controller.RefreshTokenHandler)

	// 分类列表
	r.GET("categories", controller.CategoryListHandler)
	r.GET("category/:id", controller.CategoryDetailHandler)

	// 路由分组
	ar := r.Group("/")
	ar.Use(middleware.AuthUserToken())
	{
		ar.GET("/test", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"message": "test",
			})
		})
	}
	r.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": 404,
		})
	})

	return r
}
