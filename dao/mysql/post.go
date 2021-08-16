package mysql

import (
	"database/sql"
	"go_web/web_app/dao/redis"
	"go_web/web_app/dict"
	"go_web/web_app/models"
)

// 存储帖子
func SavePost(param *models.ParamPost) (err error) {
	sqlStr := `insert into post(post_id, title, content, author_id, category_id) values(?, ?, ?, ?, ?)`
	_, err = db.Exec(sqlStr, param.Id, param.Title, param.Content, param.AuthorId, param.CategoryId)
	if err == nil {
		post, _ := GetPost(param.Id)
		return redis.SavePostTime(param.Id, post.CreatedTime.Unix())
	}
	return
}

// 获取帖子详情
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
