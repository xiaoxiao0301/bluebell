package routes

import (
	"fmt"
	"go_web/web_app/controller"
	"go_web/web_app/logger"
	"go_web/web_app/middleware"
	"go_web/web_app/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetUp() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	r.GET("test", func(context *gin.Context) {
		t := "2021-08-16 14:25:16"
		fmt.Println(models.ConvertTimestampStringToSeconds(t))
	})

	// 注册
	r.POST("signup", controller.SignUpHandler)

	// 登录
	r.POST("login", controller.LoginHandler)

	// 刷新token
	r.POST("refresh", controller.RefreshTokenHandler)

	// 保存分类
	r.POST("category", controller.CategoryStore)
	// 分类列表
	r.GET("categories", controller.CategoryListHandler)
	// 分类详情
	r.GET("category/:id", controller.CategoryDetailHandler)

	// 路由分组
	ar := r.Group("/")
	ar.Use(middleware.AuthUserToken())
	{
		// 帖子相关
		ar.POST("post", controller.PostStore)
		ar.GET("post/:id", controller.PostShow)
		ar.GET("posts", controller.PostIndex)
	}
	r.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": 404,
		})
	})

	return r
}
