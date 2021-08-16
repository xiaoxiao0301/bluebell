package mysql

import (
	"fmt"
	"go_web/web_app/conf"

	"go.uber.org/zap"

	_ "github.com/go-sql-driver/mysql"

	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func Init(mysqlConf *conf.MysqlConfig) (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		mysqlConf.User,
		mysqlConf.PassWord,
		mysqlConf.Host,
		mysqlConf.Port,
		mysqlConf.DatabaseName,
	)
	// 也可以使用MustConnect连接不成功就panic
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		zap.L().Error("connect DB failed", zap.Error(err))
		return
	}
	db.SetMaxOpenConns(mysqlConf.MaxOpenConn)
	db.SetMaxIdleConns(mysqlConf.MaxIdleConn)
	return
}

func Close() {
	_ = db.Close()
}
