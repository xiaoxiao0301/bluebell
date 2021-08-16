package redis

import (
	"context"
	"fmt"
	"go_web/web_app/conf"
	"time"

	"github.com/go-redis/redis/v8"
)

// 声明一个全局的rdb变量
var rdb *redis.Client

// 初始化连接
func Init(redisConf *conf.RedisConfig) (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d",
			redisConf.Host,
			redisConf.Port,
		),
		Password: redisConf.PassWord, // no password set
		DB:       redisConf.Db,       // use default DB
		PoolSize: redisConf.PoolSize,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err = rdb.Ping(ctx).Result()
	return
}

func Close() {
	_ = rdb.Close()
}
