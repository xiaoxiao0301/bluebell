package mysql

import (
	"database/sql"
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
