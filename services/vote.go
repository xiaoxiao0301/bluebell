package services

import (
	"go_web/web_app/dao/redis"
	"go_web/web_app/dict"
	"math"
	"time"
)

/**
投票的简单逻辑处理:
限制：
 1. 不能投重复票，举例，上次是赞同，这一次还是赞同
 2. 帖子一周之后不能投票了
计算得分逻辑
ov
  1 赞同
		0， 取消 -432         	ov-nv :  1 > 0
		-1， 反对 -432-432   	ov-nv :   2 > 0
  0 取消
        -1 -432          	ov-nv :   1 > 0
		1  +432          	ov-nv :   -1 < 0
  -1 反对
        0 +432          	ov-nv :  -1 < 0
        1 +432+432         	ov-nv :   -2 < 0

差值小于0 加分= 432 * 差值得绝对值  大于 0 减分
*/

const (
	VoteLimitTimeSeconds = 7 * 24 * 3600 // 帖子投票限制时间，帖子发表一周之内才能投票
	VotePostOneScore     = 432           // 帖子每一票增加432分
)

// VoteHandler 投票逻辑处理
func VoteHandler(userId string, postIdStr string, ov float64, nv int8) (err error) {
	// 被投票的帖子必须是一周之内
	postPublishTime, err := redis.GetPostPublishTime(postIdStr)
	if err != nil {
		return
	}
	if float64(time.Now().Unix())-postPublishTime > VoteLimitTimeSeconds {
		err = dict.ErrorVoteTimeExpires
		return
	}

	// 不能投重复票
	fnv := float64(nv)
	diffValue := ov - fnv
	if diffValue == 0 {
		return dict.ErrorVoteEqualValue
	}
	// 记录一下用户给某个帖子投票了
	switch ov {
	case -1:
		if nv == 0 {
			err = redis.SaveUserVotedPost(postIdStr, userId, 1)
		} else {
			err = redis.SaveUserVotedPost(postIdStr, userId, 2)
		}
	case 0:
		err = redis.SaveUserFirstVotedPost(postIdStr, userId, fnv)
	case 1:
		if nv == 0 {
			err = redis.SaveUserVotedPost(postIdStr, userId, -1)
		} else {
			err = redis.SaveUserVotedPost(postIdStr, userId, -2)
		}
	}
	if err != nil {
		return err
	}

	// 计算某个帖子得分
	score := 0.0
	if diffValue < 0 {
		score = VotePostOneScore * math.Abs(diffValue)
	} else {
		score = VotePostOneScore * diffValue * -1
	}
	return redis.IncrPostScore(postIdStr, score)

}
