package mysql

import (
	"database/sql"
	"go_web/web_app/dao/redis"
	"go_web/web_app/dict"
	"go_web/web_app/models"
	"strconv"
)

// SavePost 存储帖子
func SavePost(param *models.ParamPost) (err error) {
	sqlStr := `insert into post(post_id, title, content, author_id, category_id) values(?, ?, ?, ?, ?)`
	_, err = db.Exec(sqlStr, param.Id, param.Title, param.Content, param.AuthorId, param.CategoryId)
	if err == nil {
		post, _ := GetPost(param.Id)
		categoryIdStr := strconv.Itoa(int(post.CategoryId))
		err = redis.SavePostTime(param.Id, post.CreatedTime.Unix())
		if err == nil {
			err = redis.SavePostScore(param.Id, post.CreatedTime.Unix())
			if err == nil {
				return redis.SaveCategoryPostCounts(categoryIdStr, post.PostId)
			}
		}
	}
	return
}

// GetPost 获取帖子详情
func GetPost(id int64) (post *models.PostModel, err error) {
	sqlStr := `select * from post where post_id = ?`
	post = new(models.PostModel)
	err = db.Get(post, sqlStr, id)
	if err != nil {
		if err == sql.ErrNoRows {
			err = dict.ErrorNotQueryResult
		}
	}
	return
}

// GetPosts 获取帖子列表
func GetPosts(param *models.ParamPage) (posts []*models.PostModel, err error) {
	sqlStr := `select * from post order by created_time desc limit ?,?`
	offset := calculatePageAndOffset(param.Page, param.Size)
	err = db.Select(&posts, sqlStr, offset, param.Size)
	if err != nil {
		if err == sql.ErrNoRows {
			err = dict.ErrorNotQueryResult
		}
	}
	return
}
