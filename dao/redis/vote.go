package redis

import (
	"context"
	"errors"
	"go_web/web_app/dict"

	"github.com/go-redis/redis/v8"
)

// GetUserHasVoted 获取用户的投票结果
func GetUserHasVoted(postId string, userId string) (v float64, err error) {
	ctx := context.Background()
	v, err = rdb.ZScore(ctx, dict.GetUserVotedPostKey(postId), userId).Result()
	if errors.Is(err, redis.Nil) {
		// 查询结果为空
		err = nil
		return
	}
	return
}

// SaveUserVotedPost 用户给某个帖子投票
func SaveUserVotedPost(postId string, userId string, value float64) error {
	ctx := context.Background()
	return rdb.ZIncrBy(ctx, dict.GetUserVotedPostKey(postId), value, userId).Err()
}

// SaveUserVotedPost 第一次投票添加
func SaveUserFirstVotedPost(postId string, userId string, value float64) error {
	ctx := context.Background()
	return rdb.ZAdd(ctx, dict.GetUserVotedPostKey(postId), &redis.Z{
		Score:  value,
		Member: userId,
	}).Err()
}

// IncrPostScore 给帖子加分或减分
func IncrPostScore(postId string, score float64) error {
	ctx := context.Background()
	return rdb.ZIncrBy(ctx, dict.GetSavePostScoreKey(), score, postId).Err()
}
