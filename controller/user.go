package controller

import (
	"errors"
	"go_web/web_app/dict"
	"go_web/web_app/models"
	"go_web/web_app/pkg/jwt"
	"go_web/web_app/services"

	"github.com/go-playground/validator/v10"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

var UserServices services.UserService

// SignUpHandler 注册处理
// @Summary 注册
// @Description 用户注册
// @Tags 用户
// @Accept  application/json
// @Produce  application/json
// @Param user body models.ParamSignUp  true "注册信息"
// @Success 200 {object} _ResponseCommon
// @Router /signup [post]
func SignUpHandler(ctx *gin.Context) {
	var param models.ParamSignUp
	if err := ctx.ShouldBindJSON(&param); err != nil {
		zap.L().Error("SignUp with invalid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			// 非validator.ValidationErrors类型错误直接返回
			ReturnErr(ctx, dict.CodeInvalidParam)
			return
		}
		ReturnErrWithMessage(ctx, dict.CodeInvalidParam, models.RemoveTopStruct(errs.Translate(models.Trans)))
		return
	}

	if err := UserServices.RegisterUser(&param); err != nil {
		if errors.Is(err, dict.ErrorUserExists) {
			ReturnErr(ctx, dict.CodeUserExists)
		} else {
			ReturnErrWithMessage(ctx, dict.CodeInvalidParam, err.Error())
		}
		return
	}

	ReturnOk(ctx, nil)
}

// LoginHandler 登录处理
// @Summary 登陆
// @Description 用户注册
// @Tags 用户
// @Accept  application/json
// @Produce  application/json
// @Param user body models.ParamLogin  true "登陆信息"
// @Success 200 {object} _ResponseCommon
// @Success 200 {object} _ResponseLoginSuccess
// @Router /login [post]
func LoginHandler(ctx *gin.Context) {
	var param models.ParamLogin
	if err := ctx.ShouldBindJSON(&param); err != nil {
		zap.L().Error("LoginHandler with invalid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			// 非validator.ValidationErrors类型错误直接返回
			ReturnErrWithMessage(ctx, dict.CodeInvalidParam, err.Error())
			return
		}
		ReturnErrWithMessage(ctx, dict.CodeInvalidParam, models.RemoveTopStruct(errs.Translate(models.Trans)))
		return
	}

	userinfo, err := UserServices.LoginUser(&param)
	if err != nil {
		if errors.Is(err, dict.ErrorUserNameOrPassword) {
			ReturnErrWithMessage(ctx, dict.CodeInvalidPassword, "用户名或密码错误")
		} else {
			ReturnErr(ctx, dict.CodeUserNotExists)
		}
		return
	}

	accessToken, refreshToken, err := jwt.CreateToken(userinfo.UserId, userinfo.Username)
	if err != nil {
		ReturnErr(ctx, dict.CodeNetWorkBusy)
		return
	}

	var data map[string]interface{}
	data = make(map[string]interface{})
	data["userid"] = userinfo.UserId
	data["username"] = userinfo.Username
	data["access_token"] = accessToken
	data["refresh_token"] = refreshToken

	ReturnOk(ctx, data)
}
