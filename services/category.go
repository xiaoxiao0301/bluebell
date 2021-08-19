package services

import (
	"errors"
	"go_web/web_app/dao/mysql"
	"go_web/web_app/dao/redis"
	"go_web/web_app/dict"
	"go_web/web_app/models"
	"go_web/web_app/pkg/snowflake"
	"strconv"
)

// 社区服务层
type CategoryService struct {
}

// CategoryList 获取数据库中社区列表信息
func (cs *CategoryService) CategoryList() ([]*models.CategoryRow, error) {
	return mysql.GetCategoryList()
}

// CategoryDetail 获取数据库中社区详细信息
func (cs *CategoryService) CategoryDetail(id int64) (*models.CategoryModel, error) {
	// 先从获取中获取，缓存没有再从数据库获取
	idStr := strconv.Itoa(int(id))
	category, err := redis.GetCategoryDetail(idStr)
	if err != nil {
		if errors.Is(err, dict.ErrorNotQueryResult) {
			return mysql.GetCategoryDetail(id)
		}
		return nil, err
	}

	return category, nil
}

// CategoryStore 存储社区信息
func (cs *CategoryService) CategoryStore(p *models.ParamCategory) (err error) {

	var category = &models.CategoryModel{
		CategoryId:   snowflake.GenID(),
		CategoryName: p.CategoryName,
		Introduction: p.Introduction,
	}

	return mysql.SaveCategory(category)
}
