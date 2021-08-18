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

// PostIndex 帖子列表
func PostIndex(ctx *gin.Context) {
	var param models.ParamPage
	if err := ctx.ShouldBindQuery(&param); err != nil {
		zap.L().Error("PostIndex with invalid param, err:", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ReturnErr(ctx, dict.CodeInvalidParam)
			return
		}
		ReturnErrWithMessage(ctx, dict.CodeInvalidParam, models.RemoveTopStruct(errs.Translate(models.Trans)))
		return
	}
	posts, err := PostService.GetPostList(&param)
	if err != nil {
		if err == dict.ErrorNotQueryResult {
			ReturnOk(ctx, nil)
			return
		}
		zap.L().Error("PostService.GetPostList err:", zap.Error(err))
		ReturnErr(ctx, dict.CodeNetWorkBusy)
		return
	}
	ReturnOk(ctx, posts)
}

// PostVoteStore 帖子投票
func PostVoteStore(ctx *gin.Context) {
	var param models.ParamVote
	if err := ctx.ShouldBindJSON(&param); err != nil {
		zap.L().Error("PostVoteStore with invalid param", zap.Error(err))
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

	err = PostService.VotePost(&param, userId)
	if err != nil {
		// 投了相同值得票
		if err == dict.ErrorVoteEqualValue {
			ReturnErr(ctx, dict.CodeVotedEqualResult)
			return
		}
		// 投票时间过了
		if err == dict.ErrorVoteTimeExpires {
			ReturnErr(ctx, dict.CodeVoteTimeExpires)
			return
		}
		zap.L().Error("PostService.VotePost err:", zap.Error(err))
		ReturnErr(ctx, dict.CodeNetWorkBusy)
		return
	}
	ReturnOk(ctx, nil)
}

// NewPostsIndex 新版帖子接口
func NewPostsIndex(ctx *gin.Context) {
	// 请求URL v2/posts?page=1&size=10&order=time&sorts=asc
	var param models.ParamNewPostList
	if err := ctx.ShouldBindQuery(&param); err != nil {
		zap.L().Error("NewPostsIndex with invalid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ReturnErr(ctx, dict.CodeInvalidParam)
			return
		}
		ReturnErrWithMessage(ctx, dict.CodeInvalidParam, models.RemoveTopStruct(errs.Translate(models.Trans)))
		return
	}
	posts, err := PostService.NewPostLists(&param)
	if err != nil {
		if err == dict.ErrorNotQueryResult {
			ReturnOk(ctx, "暂无数据")
			return
		}
		zap.L().Error("PostService.NewPostLists error:", zap.Error(err))
		ReturnErr(ctx, dict.CodeNetWorkBusy)
		return
	}
	ReturnOk(ctx, posts)
}
