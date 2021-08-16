package redis

import (
	"context"
	"go_web/web_app/dict"

	"github.com/go-redis/redis/v8"
)

// 根据帖子发表时间存储
func SavePostTime(postId int64, ts int64) error {
	ctx := context.Background()
	return rdb.ZAdd(ctx, dict.GetSavePostTimeKey(), &redis.Z{
		Score:  float64(ts),
		Member: postId,
	}).Err()
}
