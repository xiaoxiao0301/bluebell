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
// @Summary 存储帖子
// @Description 存储帖子
// @Tags 帖子
// @Accept application/json
// @Produce  application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param posts body models.ParamPost true "帖子"
// @Success 200 {object} _ResponseCommon
// @Security ApiKeyAuth
// @Router /post [post]
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
// @Summary 帖子详情
// @Description 帖子详情
// @Tags 帖子
// @Produce  application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param id path string true "帖子ID" default(3765906580705280)
// @Success 200 {object} models.PostModel
// @Security ApiKeyAuth
// @Router /post/{id} [get]
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
// @Summary 帖子列表
// @Description 帖子列表
// @Tags 帖子
// @Produce  application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param page query int true "页码" default(1)
// @Param size query int true "每页大小" default(10)
// @Success 200 {object} models.PostListDetail
// @Security ApiKeyAuth
// @Router /posts [get]
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
// @Summary 帖子投票
// @Description 用户可以给帖子投赞成或者反对票
// @Tags 帖子
// @Accept application/json
// @Produce  application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param vote body models.ParamVote true "投票"
// @Success 200 {object} _ResponseCommon
// @Success 200 {object} models.PostListDetail
// @Security ApiKeyAuth
// @Router /post/vote [post]
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
// @Summary 帖子列表
// @Description 可以根据发帖时间和帖子分数来获取帖子列表
// @Tags 帖子
// @Produce  application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param page query int true "页码" default(1)
// @Param size query int true "每页大小" default(10)
// @Param order query string true "排序依据, time 时间 score 得分" default(time)
// @Param sorts query string true "升序还是降序 asc 升序 desc 降序" default(asc)
// @Success 200 {object} models.PostListDetail
// @Security ApiKeyAuth
// @Router /v2/posts [get]
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
