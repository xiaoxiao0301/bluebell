package models

// CategoryRow 行记录
type CategoryRow struct {
	Id   int64  `json:"id" db:"category_id"`
	Name string `json:"name" db:"category_name"`
}

// CategoryModel 详细信息
type CategoryModel struct {
	Id           int64     `json:"id" db:"id"`
	CategoryId   int64     `json:"category_id,string" db:"category_id"`
	CategoryName string    `json:"category_name" db:"category_name"`
	Introduction string    `json:"introduction" db:"introduction"`
	CreatedTime  LocalTime `json:"created_time" db:"created_time"`
	UpdatedTime  LocalTime `json:"updated_time" db:"updated_time"`
}
