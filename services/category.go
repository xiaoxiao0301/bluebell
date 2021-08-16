package services

import (
	"go_web/web_app/dao/mysql"
	"go_web/web_app/models"
)

// 分类服务层
type CategoryService struct {
}

// CategoryList 获取数据库中分类列表信息
func (cs *CategoryService) CategoryList() ([]*models.CategoryRow, error) {
	return mysql.GetCategoryList()
}

// CategoryDetail 获取数据库中分类详细信息
func (cs *CategoryService) CategoryDetail(id int64) (*models.CategoryModel, error) {
	return mysql.GetCategoryDetail(id)
}
