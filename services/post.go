package services

import (
	"go_web/web_app/dao/mysql"
	"go_web/web_app/dao/redis"
	"go_web/web_app/dict"
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
	return generaPostDetailWithUserAndCategory(posts)
}

// VotePost 给帖子投票
func (p *PostService) VotePost(param *models.ParamVote, userId int64) (err error) {
	// 本次投票的结果值
	newValue := *param.Value
	// 不能重复投赞同票
	postIdStr := strconv.Itoa(int(param.PostId))
	userIdStr := strconv.Itoa(int(userId))
	// 原来投票的值
	ov, err := redis.GetUserHasVoted(postIdStr, userIdStr)
	if err != nil {
		return err
	}
	return VoteHandler(userIdStr, postIdStr, ov, newValue)
}

// NewPostLists 新版帖子列表
func (p *PostService) NewPostLists(param *models.ParamNewPostList) (data []*models.PostListDetail, err error) {
	start := (param.Page - 1) * param.Size
	end := param.Size + start - 1

	var postIds []string
	order := param.Order
	sorts := param.Sorts
	if order == "time" {
		postIds, err = redis.GetPostsIdsByTime(sorts, int64(start), int64(end))
	} else {
		postIds, err = redis.GetPostsIdsByScore(sorts, int64(start), int64(end))
	}
	if err != nil {
		zap.L().Error("从缓存获取ids出错:", zap.Error(err))
		return make([]*models.PostListDetail, 0), err
	} else if err == nil {
		zap.L().Info("未获取数据", zap.String("order", order), zap.String("sorts", sorts),
			zap.Int("start", start), zap.Int("end", end))
		return make([]*models.PostListDetail, 0), nil
	}
	posts, err := mysql.GetPostsListByIds(postIds)
	if err != nil {
		zap.L().Error("mysql.GetPostsListByIds err:", zap.Error(err))
		return nil, err
	}
	return generaPostDetailWithUserAndCategory(posts)
}

// GetPosts 获取分类下的帖子
func (p *PostService) GetPosts(categoryId string) (data []*models.PostModel, err error) {
	data = make([]*models.PostModel, 0)
	postIds, err := redis.GetCategoryPosts(categoryId)
	if err != nil {
		if err == dict.ErrorNotQueryResult {
			return data, nil
		}
		zap.L().Error("redis.GetCategoryPosts error", zap.Error(err))
		return data, err
	}
	data, err = mysql.GetPostsListByIds(postIds)
	if err != nil {
		zap.L().Error("mysql.GetPostsListByIds error", zap.Error(err))
		return data, err
	}
	return
}

// 组装返回 models.PostListDetail 数据
func generaPostDetailWithUserAndCategory(posts []*models.PostModel) (data []*models.PostListDetail, err error) {
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
