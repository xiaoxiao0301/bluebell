package redis

import (
	"context"
	"go_web/web_app/dict"

	"github.com/go-redis/redis/v8"
)

// SavePostTime 根据帖子发表时间存储
func SavePostTime(postId int64, ts int64) error {
	ctx := context.Background()
	return rdb.ZAdd(ctx, dict.GetSavePostTimeKey(), &redis.Z{
		Score:  float64(ts),
		Member: postId,
	}).Err()
}

// SavePostScore 存储每个帖子的得分，基础是发表时间
func SavePostScore(postId int64, ts int64) error {
	ctx := context.Background()
	return rdb.ZAdd(ctx, dict.GetSavePostScoreKey(), &redis.Z{
		Score:  float64(ts),
		Member: postId,
	}).Err()
}

// GetPostPublishTime 获取帖子的发表时间
func GetPostPublishTime(postId string) (float64, error) {
	ctx := context.Background()
	return rdb.ZScore(ctx, dict.GetSavePostTimeKey(), postId).Result()
}

// GetCategoryPosts 获取分类下的帖子信息
func GetCategoryPosts(categoryId string) ([]string, error) {
	ctx := context.Background()
	// [],nil or [有值] nil
	ids, err := rdb.SMembers(ctx, dict.GetSaveCategoryPostsCountKey(categoryId)).Result()
	if len(ids) > 0 {
		return ids, nil
	} else if len(ids) == 0 {
		return ids, dict.ErrorNotQueryResult
	}
	return ids, err

}
