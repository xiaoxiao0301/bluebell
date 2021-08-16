package mysql

import (
	"database/sql"
	"go_web/web_app/dao/redis"
	"go_web/web_app/dict"
	"go_web/web_app/models"
)

// GetCategoryList 获取分类列表
func GetCategoryList() (list []*models.CategoryRow, err error) {
	sqlStr := `select category_id,category_name from category`
	err = db.Select(&list, sqlStr)
	if err != nil {
		if err == sql.ErrNoRows {
			err = dict.ErrorNotQueryResult
		}
	}
	return
}

// GetCategoryDetail 获取分类详细
func GetCategoryDetail(id int64) (category *models.CategoryModel, err error) {
	sqlStr := `select * from category where category_id = ?`
	category = new(models.CategoryModel)
	err = db.Get(category, sqlStr, id)
	if err != nil {
		if err == sql.ErrNoRows {
			err = dict.ErrorNotQueryResult
		}
	}
	return
}

// SaveCategory 存储分类信息
func SaveCategory(category *models.CategoryModel) (err error) {
	sqlStr := `insert into category(category_id, category_name, introduction) values(?, ?, ?)`
	_, err = db.Exec(sqlStr, category.CategoryId, category.CategoryName, category.Introduction)
	if err == nil {
		category, _ := GetCategoryDetail(category.CategoryId)
		return redis.SaveCategory(category)
	}
	return
}
