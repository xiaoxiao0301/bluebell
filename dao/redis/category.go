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

// SaveCategory 将分类信息缓存
func SaveCategory(category *models.CategoryModel) (err error) {
	strId := strconv.Itoa(int(category.CategoryId))
	ctx := context.Background()
	categoryJson, _ := json.Marshal(category)
	return rdb.Set(ctx, dict.GetSaveCategoryKey(strId), categoryJson, 0).Err()
}

// GetCategoryDetail 获取分类详细信息
func GetCategoryDetail(categoryId string) (category *models.CategoryModel, err error) {
	ctx := context.Background()
	categoryJson, err := rdb.Get(ctx, dict.GetSaveCategoryKey(categoryId)).Result()
	if errors.Is(err, redis.Nil) {
		return nil, dict.ErrorNotQueryResult
	}
	category = new(models.CategoryModel)
	err = json.Unmarshal([]byte(categoryJson), category)
	if err != nil {
		return nil, err
	}
	return
}

// SaveCategoryPostCounts 统计分类下帖子的数量
func SaveCategoryPostCounts(categoryId string, postId int64) error {
	ctx := context.Background()
	return rdb.SAdd(ctx, dict.GetSaveCategoryPostsCountKey(categoryId), postId).Err()
}
