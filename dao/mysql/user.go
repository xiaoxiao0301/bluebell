package mysql

import (
	"go_web/web_app/models"

	"go.uber.org/zap"
)

// CheckUserByUsername 根据用户名判断用户是否已经存在
func CheckUserByUsername(username string) bool {
	sqlStr := `select count(id) from user where username = ?`
	var count int
	if err := db.Get(&count, sqlStr, username); err != nil {
		zap.L().Error("CheckUserByUsername exec error:", zap.Error(err))
		return true
	}
	if count > 0 {
		return true
	}
	return false
}

// SaveUser 将用户信息插入数据库中
func SaveUser(userRows *models.UserRow) (err error) {
	sqlStr := `insert into user(user_id, username, password)values(?, ?, ?)`
	_, err = db.Exec(sqlStr, userRows.UserId, userRows.Username, userRows.Password)
	return
}

// GetUserInfoByUsername 根据用户名获取用户信息
func GetUserInfoByUsername(username string) (*models.UserModel, error) {
	var userinfo models.UserModel
	sqlStr := `select * from user where username = ?`
	if err := db.Get(&userinfo, sqlStr, username); err != nil {
		return nil, err
	}
	return &userinfo, nil
}
