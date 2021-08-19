package routes

import (
	"fmt"
	"go_web/web_app/controller"
	"go_web/web_app/logger"
	"go_web/web_app/middleware"
	"go_web/web_app/models"
	"net/http"

	_ "go_web/web_app/docs" // 引入 swagger生成的docs https://github.com/swaggo/gin-swagger

	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	"github.com/gin-gonic/gin"
)

func SetUp() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	// swagger路由
	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	// 测试路由
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

	// 保存社区
	r.POST("category", controller.CategoryStore)
	// 社区列表
	r.GET("categories", controller.CategoryListHandler)
	// 社区详情
	r.GET("category/:id", controller.CategoryDetailHandler)

	// 路由分组
	ar := r.Group("/")
	ar.Use(middleware.AuthUserToken())
	{
		// 帖子相关
		ar.POST("post", controller.PostStore)
		ar.GET("post/:id", controller.PostShow)
		ar.GET("posts", controller.PostIndex)
		ar.POST("post/vote", controller.PostVoteStore)
		// 新的帖子接口可以根据参数 order=time or score 来排序
		ar.GET("v2/posts", controller.NewPostsIndex)
		ar.GET("category/:id/posts", controller.GetCategoryIdPosts)

	}
	r.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": 404,
		})
	})

	return r
}
