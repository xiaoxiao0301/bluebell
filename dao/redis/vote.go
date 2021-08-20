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

// GetPostsIdsByTime 根据发帖时间获取帖子ids
func GetPostsIdsByTime(sorts string, start, end int64) (ids []string, err error) {
	ctx := context.Background()
	if sorts == "asc" {
		// 升序，默认 zRange
		ids, err = rdb.ZRange(ctx, dict.GetSavePostTimeKey(), start, end).Result()
	} else {
		// 降序， zRevRange
		ids, err = rdb.ZRevRange(ctx, dict.GetSavePostTimeKey(), start, end).Result()
	}
	if len(ids) > 0 {
		return ids, nil
	} else if len(ids) == 0 {
		return ids, dict.ErrorNotQueryResult
	}
	return
}

// GetPostsIdsByScore 根据帖子分数获取帖子ids
func GetPostsIdsByScore(sorts string, start, end int64) (ids []string, err error) {
	ctx := context.Background()
	if sorts == "asc" {
		// 升序，默认 zRange
		ids, err = rdb.ZRange(ctx, dict.GetSavePostScoreKey(), start, end).Result()
	} else {
		// 降序， zRevRange
		ids, err = rdb.ZRevRange(ctx, dict.GetSavePostScoreKey(), start, end).Result()
	}
	if len(ids) > 0 {
		return ids, nil
	} else if len(ids) == 0 {
		return ids, dict.ErrorNotQueryResult
	}
	return
}
