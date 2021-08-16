package controller

import (
	"errors"
	"fmt"
	"go_web/web_app/dict"
	"go_web/web_app/models"
	"go_web/web_app/services"

	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

var tokenService services.TokenService

func RefreshTokenHandler(ctx *gin.Context) {
	var param models.ParamRefreshToken
	if err := ctx.ShouldBindJSON(&param); err != nil {
		zap.L().Error("RefreshToken with invalid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			// 非validator.ValidationErrors类型错误直接返回
			ReturnErrWithMessage(ctx, dict.CodeInvalidParam, err.Error())
			return
		}
		ReturnErrWithMessage(ctx, dict.CodeInvalidParam, models.RemoveTopStruct(errs.Translate(models.Trans)))
		return
	}
	accessToken, refreshToken, err := tokenService.RefreshToken(param.AccessToken, param.RefreshToken)
	fmt.Println(accessToken)
	fmt.Println(refreshToken)
	if err != nil {
		if errors.Is(err, dict.ErrorAccessTokenValid) {
			ReturnErr(ctx, dict.CodeValidAccessToken)
			return
		}
		ReturnErr(ctx, dict.CodeInvalidRefreshToken)
		return
	}
	var data map[string]interface{}
	data = make(map[string]interface{})
	data["access_token"] = accessToken
	data["refresh_token"] = refreshToken

	ReturnOk(ctx, data)

}
