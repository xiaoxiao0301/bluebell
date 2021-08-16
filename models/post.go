package models

// 内存对齐

// PostModel 帖子信息
type PostModel struct {
	Id          int64     `json:"id" db:"id"`
	PostId      int64     `json:"post_id,string" db:"post_id"`
	AuthorId    int64     `json:"author_id,string" db:"author_id"`
	CategoryId  int64     `json:"category_id,string" db:"author_id"`
	Status      int32     `json:"status" db:"status"`
	Title       string    `json:"title" db:"title"`
	Content     string    `json:"content" db:"content"`
	CreatedTime LocalTime `json:"created_time" db:"created_time"`
	UpdatedTime LocalTime `json:"updated_time" db:"updated_time"`
}
