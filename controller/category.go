package controller

import (
	"errors"
	"go_web/web_app/dict"
	"go_web/web_app/models"
	"go_web/web_app/services"
	"strconv"

	"github.com/go-playground/validator/v10"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

var CategoryService services.CategoryService

// CategoryStore 新建分类
func CategoryStore(ctx *gin.Context) {
	var param models.ParamCategory

	if err := ctx.ShouldBindJSON(&param); err != nil {
		zap.L().Error("CategoryStore with invalid param, err :", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ReturnErr(ctx, dict.CodeInvalidParam)
			return
		}
		ReturnErrWithMessage(ctx, dict.CodeInvalidParam, models.RemoveTopStruct(errs.Translate(models.Trans)))
		return
	}
	err := CategoryService.CategoryStore(&param)
	if err != nil {
		zap.L().Error("categoryService.CategoryStore with err:", zap.Error(err))
		ReturnErr(ctx, dict.CodeNetWorkBusy)
		return
	}
	ReturnOk(ctx, nil)
}

// CategoryListHandler 分类列表
func CategoryListHandler(ctx *gin.Context) {
	categoryList, err := CategoryService.CategoryList()
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
	category, err := CategoryService.CategoryDetail(id)
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

// GetCategoryIdPosts 获取某个分类下的所有帖子
func GetCategoryIdPosts(ctx *gin.Context) {
	idStr := ctx.Param("id")
	if idStr == "" {
		ReturnErr(ctx, dict.CodeInvalidParam)
		return
	}
	posts, err := PostService.GetPosts(idStr)
	if err != nil {
		if errors.Is(err, dict.ErrorNotQueryResult) {
			ReturnErr(ctx, dict.CodeNotQueryResult)
			return
		}
		zap.L().Error("PostService.NewPostLists error:", zap.Error(err))
		ReturnErr(ctx, dict.CodeNetWorkBusy)
		return
	}
	ReturnOk(ctx, posts)
}
