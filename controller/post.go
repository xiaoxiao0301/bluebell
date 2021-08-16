package controller

import (
	"go_web/web_app/dict"
	"go_web/web_app/models"
	"go_web/web_app/services"
	"strconv"

	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

var PostService services.PostService

// PostStore 存储帖子
func PostStore(ctx *gin.Context) {
	var param models.ParamPost
	if err := ctx.ShouldBindJSON(&param); err != nil {
		zap.L().Error("PostStore with invalid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ReturnErr(ctx, dict.CodeInvalidParam)
			return
		}
		ReturnErrWithMessage(ctx, dict.CodeInvalidParam, models.RemoveTopStruct(errs.Translate(models.Trans)))
		return
	}
	userId, err := UserServices.GetLoginUserId(ctx)
	if err != nil {
		ReturnErr(ctx, dict.CodeNeedLogin)
		return
	}
	err = PostService.SavePost(&param, userId)
	if err != nil {
		zap.L().Error("postService.SavePost error", zap.Error(err))
		ReturnErr(ctx, dict.CodeNetWorkBusy)
		return
	}
	ReturnOk(ctx, nil)
}

// PostShow 帖子详情
func PostShow(ctx *gin.Context) {
	idStr := ctx.Param("id")
	if idStr == "" {
		ReturnErr(ctx, dict.CodeInvalidParam)
		return
	}
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ReturnErr(ctx, dict.CodeInvalidParam)
		return
	}
	post, err := PostService.GetPostDetail(id)
	if err != nil {
		if err == dict.ErrorNotQueryResult {
			ReturnErr(ctx, dict.CodeNotQueryResult)
			return
		}
		ReturnErr(ctx, dict.CodeNetWorkBusy)
		return
	}
	ReturnOk(ctx, post)
}
