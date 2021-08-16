package services

import (
	"crypto/md5"
	"encoding/hex"
	"go_web/web_app/dao/mysql"
	"go_web/web_app/dict"
	"go_web/web_app/models"
	"go_web/web_app/pkg/snowflake"

	"go.uber.org/zap"
)

// secret md5加密salt
const secret = "jixiaoxiao.com"

// UserService 用户服务层
type UserService struct {
}

// RegisterUser 用户注册
func (services *UserService) RegisterUser(userParam *models.ParamSignUp) (err error) {
	// 判断用户是否已经存在
	exists := mysql.CheckUserByUsername(userParam.Username)
	if exists {
		return dict.ErrorUserExists
	}
	// 对明文密码进行加密处理
	userParam.Password = encryptPassword(userParam.Password)

	// 构造插入用户记录
	var user = &models.UserRow{
		UserId:   snowflake.GenID(),
		Username: userParam.Username,
		Password: userParam.Password,
	}

	// 用户数据入库
	return mysql.SaveUser(user)
}

// LoginUser 用户登录
func (service *UserService) LoginUser(userParam *models.ParamLogin) (userinfo *models.UserModel, err error) {
	// 获取数据库存储的密码和输入的密码进行校验
	userinfo, err = mysql.GetUserInfoByUsername(userParam.Username)
	if err != nil {
		zap.L().Error("获取用户信息失败:", zap.String("username", userParam.Username))
		return nil, dict.ErrorUserNotExists
	}
	if encryptPassword(userParam.Password) != userinfo.Password {
		return nil, dict.ErrorUserNameOrPassword
	}
	return userinfo, nil
}

// encryptPassword 密码进行md5加密
func encryptPassword(oldPassword string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(oldPassword)))
}
