package services

import (
	"go_web/web_app/dao/mysql"
	"go_web/web_app/models"
	"go_web/web_app/pkg/snowflake"
)

type PostService struct {
}

// SavePost 存储帖子
func (p *PostService) SavePost(param *models.ParamPost, userId int64) (err error) {
	param.Id = snowflake.GenID()
	param.AuthorId = userId
	return mysql.SavePost(param)
}

func (p *PostService) GetPostDetail(id int64) (post *models.PostModel, err error) {
	return mysql.GetPost(id)
}
