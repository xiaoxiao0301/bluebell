package services

import (
	"go_web/web_app/dao/mysql"
	"go_web/web_app/dao/redis"
	"go_web/web_app/models"
	"go_web/web_app/pkg/snowflake"
	"strconv"

	"go.uber.org/zap"
)

type PostService struct {
}

// SavePost 存储帖子
func (p *PostService) SavePost(param *models.ParamPost, userId int64) (err error) {
	param.Id = snowflake.GenID()
	param.AuthorId = userId
	return mysql.SavePost(param)
}

// GetPostDetail 帖子详情
func (p *PostService) GetPostDetail(id int64) (post *models.PostModel, err error) {
	return mysql.GetPost(id)
}

// GetPostList 分页获取帖子列表
func (p *PostService) GetPostList(param *models.ParamPage) (data []*models.PostListDetail, err error) {
	posts, err := mysql.GetPosts(param)
	if err != nil {
		zap.L().Error("GetPostList mysql.GetPosts err:", zap.Error(err))
		return nil, err
	}
	data = make([]*models.PostListDetail, 0, len(posts))
	for _, post := range posts {
		userIdStr := strconv.Itoa(int(post.AuthorId))
		user, err := redis.GetUser(userIdStr)
		if err != nil {
			zap.L().Error("GetPostList redis.GetUser err:", zap.Error(err))
			return nil, err
		}

		categoryIdStr := strconv.Itoa(int(post.CategoryId))
		category, err := redis.GetCategoryDetail(categoryIdStr)
		if err != nil {
			zap.L().Error("GetPostList redis.GetCategoryDetail err:", zap.Error(err))
			return nil, err
		}
		postDetail := &models.PostListDetail{
			UserModel:     user,
			PostModel:     post,
			CategoryModel: category,
		}
		data = append(data, postDetail)
	}
	return
}
