package models

// UserRow 插入数据库的一条记录
type UserRow struct {
	UserId   int64  `json:"user_id" db:"user_id"`
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
}

// UserModel 获取数据库用户信息
type UserModel struct {
	Id          int64     `json:"id" db:"id"`
	UserId      int64     `json:"user_id,string" db:"user_id"`
	Username    string    `json:"username" db:"username"`
	Password    string    `json:"password" db:"password"`
	Email       *string   `json:"email" db:"email"`
	Gender      string    `json:"gender" db:"gender"`
	CreatedTime LocalTime `json:"created_time" db:"created_time"`
	UpdatedTime LocalTime `json:"updated_time" db:"updated_time"`
}
