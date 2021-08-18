package redis

import (
	"context"
	"encoding/json"
	"errors"
	"go_web/web_app/dict"
	"go_web/web_app/models"
	"strconv"

	"github.com/go-redis/redis/v8"
)

// SaveUser 将用户信息存储到缓存中
func SaveUser(user *models.UserModel) (err error) {
	userJson, _ := json.Marshal(user)
	ctx := context.Background()
	idStr := strconv.Itoa(int(user.UserId))
	return rdb.Set(ctx, dict.GetSaveUserKey(idStr), userJson, 0).Err()
}

// GetUser 从缓存中获取用户信息
func GetUser(id string) (user *models.UserModel, err error) {
	ctx := context.Background()
	data, err := rdb.Get(ctx, dict.GetSaveUserKey(id)).Result()
	if errors.Is(err, redis.Nil) {
		return nil, dict.ErrorNotQueryResult
	}
	user = new(models.UserModel)
	err = json.Unmarshal([]byte(data), user)
	if err != nil {
		return nil, err
	}
	return
}
