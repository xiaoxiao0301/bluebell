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

// CategoryStore 新建社区
// @Summary 新建社区
// @Description 新建社区
// @Tags 社区
// @Accept  application/json
// @Produce  application/json
// @Param category body models.ParamCategory true "社区"
// @Success 200 {object} _ResponseCommon
// @Router /category [post]
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

// CategoryListHandler 社区列表
// @Summary 社区列表
// @Description 社区列表
// @Tags 社区
// @Produce  application/json
// @Success 200 {object} models.CategoryRow
// @Router /categories [get]
func CategoryListHandler(ctx *gin.Context) {
	categoryList, err := CategoryService.CategoryList()
	if err != nil {
		if errors.Is(err, dict.ErrorNotQueryResult) {
			ReturnErr(ctx, dict.CodeNotQueryResult)
			return
		}
		zap.L().Error("获取社区列表出错", zap.Error(err))
		ReturnErr(ctx, dict.CodeNetWorkBusy)
		return
	}
	ReturnOk(ctx, categoryList)
}

// CategoryDetailHandler 社区详情
// @Summary 社区详情
// @Description 社区详情
// @Tags 社区
// @Produce  application/json
// @Param id path string true "社区ID" default(3730413906300928)
// @Success 200 {object} _ResponseCommon
// @Success 200 {object} models.CategoryModel
// @Router /category/{id} [get]
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
		zap.L().Error("获取社区详情出错", zap.Error(err))
		ReturnErr(ctx, dict.CodeNetWorkBusy)
		return
	}
	ReturnOk(ctx, category)
}

// GetCategoryIdPosts 获取某个社区下的所有帖子
// @Summary 获取某个社区下的所有帖子
// @Description 获取某个社区下的所有帖子
// @Tags 社区
// @Produce  application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param id path string true "社区ID" default(3730413906300928)
// @Success 200 {object} _ResponseCommon
// @Success 200 {object} models.PostModel
// @Security ApiKeyAuth
// @Router /category/{id}/posts [get]
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
