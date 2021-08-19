package models

// CategoryRow 行记录
type CategoryRow struct {
	Id   int64  `json:"id" db:"category_id"`     // 社区ID
	Name string `json:"name" db:"category_name"` // 社区名称
}

// CategoryModel 详细信息
type CategoryModel struct {
	Id           int64     `json:"id" db:"id"`                          // ID
	CategoryId   int64     `json:"category_id,string" db:"category_id"` // 社区ID
	CategoryName string    `json:"category_name" db:"category_name"`    // 社区名称
	Introduction string    `json:"introduction" db:"introduction"`      // 社区简介
	CreatedTime  LocalTime `json:"created_time" db:"created_time"`      // 社区创建时间
	UpdatedTime  LocalTime `json:"updated_time" db:"updated_time"`      // 社区更新时间
}
