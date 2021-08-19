package main

import (
	"context"
	"fmt"
	"go_web/web_app/conf"
	"go_web/web_app/dao/mysql"
	"go_web/web_app/dao/redis"
	"go_web/web_app/logger"
	"go_web/web_app/models"
	"go_web/web_app/pkg/snowflake"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	router "go_web/web_app/routes"

	"go.uber.org/zap"
)

// @title BlueBell Api
// @version 1.0
// @description 使用gin开发简单帖子展示系统

// @host localhost:8080
// @BasePath /

func main() {
	// 加载配置文件
	if err := conf.Init(); err != nil {
		fmt.Printf("init config setting faild, err:%v\n", err)
		return
	}
	// 初始化日志
	if err := logger.Init(conf.Conf.LogConfig, conf.Conf.Mode); err != nil {
		fmt.Printf("init config logger faild, err:%v\n", err)
		return
	}
	defer logger.Sync()
	// 初始化MySQL连接
	if err := mysql.Init(conf.Conf.MysqlConfig); err != nil {
		fmt.Printf("init config mysql faild, err:%v\n", err)
		return
	}
	defer mysql.Close()
	// 初始化Redis连接
	if err := redis.Init(conf.Conf.RedisConfig); err != nil {
		fmt.Printf("init config redis faild, err:%v\n", err)
		return
	}
	defer redis.Close()

	if err := snowflake.Init(conf.Conf.StartTime, conf.Conf.MachineID); err != nil {
		fmt.Printf("init snowflake faild, err:%v\n", err)
		return
	}

	if err := models.InitTrans("zh"); err != nil {
		fmt.Printf("init trans faild, err:%v\n", err)
		return
	}

	// 注册路由
	r := router.SetUp()
	// 启动服务
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", conf.Conf.Port),
		Handler: r,
	}

	go func() {
		// 开启一个goroutine启动服务
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			zap.L().Fatal("listen err :", zap.Error(err))
		}
	}()

	// 等待中断信号来优雅地关闭服务器，为关闭服务器操作设置一个5秒的超时
	quit := make(chan os.Signal, 1) // 创建一个接收信号的通道
	// kill 默认会发送 syscall.SIGTERM 信号
	// kill -2 发送 syscall.SIGINT 信号，我们常用的Ctrl+C就是触发系统SIGINT信号
	// kill -9 发送 syscall.SIGKILL 信号，但是不能被捕获，所以不需要添加它
	// signal.Notify把收到的 syscall.SIGINT或syscall.SIGTERM 信号转发给quit
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) // 此处不会阻塞
	<-quit                                               // 阻塞在此，当接收到上述两种信号时才会往下执行
	zap.L().Info("Shutdown Server ...")
	// 创建一个5秒超时的context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// 5秒内优雅关闭服务（将未处理完的请求处理完再关闭服务），超过5秒就超时退出
	if err := srv.Shutdown(ctx); err != nil {
		zap.L().Fatal("Server Shutdown: ", zap.Error(err))
	}

	zap.L().Info("Server exiting")
}
