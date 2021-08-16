package controller

import (
	"errors"
	"go_web/web_app/dict"
	"go_web/web_app/services"
	"strconv"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

var categoryService services.CategoryService

// CategoryListHandler 分类列表
func CategoryListHandler(ctx *gin.Context) {
	categoryList, err := categoryService.CategoryList()
	if err != nil {
		if errors.Is(err, dict.ErrorNotQueryResult) {
			ReturnErr(ctx, dict.CodeNotQueryResult)
			return
		}
		zap.L().Error("获取分类列表出错", zap.Error(err))
		ReturnErr(ctx, dict.CodeNetWorkBusy)
		return
	}
	ReturnOk(ctx, categoryList)
}

// CategoryDetailHandler 分类详情
func CategoryDetailHandler(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ReturnErr(ctx, dict.CodeInvalidParam)
		return
	}
	category, err := categoryService.CategoryDetail(id)
	if err != nil {
		if errors.Is(err, dict.ErrorNotQueryResult) {
			ReturnErr(ctx, dict.CodeNotQueryResult)
			return
		}
		zap.L().Error("获取分类详情出错", zap.Error(err))
		ReturnErr(ctx, dict.CodeNetWorkBusy)
		return
	}
	ReturnOk(ctx, category)
}
