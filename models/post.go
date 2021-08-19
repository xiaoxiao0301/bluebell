package models

// 内存对齐

// PostModel 帖子信息
type PostModel struct {
	Id          int64     `json:"id" db:"id"`                          //id
	PostId      int64     `json:"post_id,string" db:"post_id"`         // 帖子ID
	AuthorId    int64     `json:"author_id,string" db:"author_id"`     // 帖子作者ID
	CategoryId  int64     `json:"category_id,string" db:"category_id"` // 帖子社区ID
	Status      int32     `json:"status" db:"status"`                  // 帖子状态
	Title       string    `json:"title" db:"title"`                    // 帖子标题
	Content     string    `json:"content" db:"content"`                //帖子内容
	CreatedTime LocalTime `json:"created_time" db:"created_time"`      // 帖子创建时间
	UpdatedTime LocalTime `json:"updated_time" db:"updated_time"`      // 帖子更新时间
}

// PostListDetail 帖子列表详情
type PostListDetail struct {
	*UserModel     `json:"user"`     // 用户信息
	*PostModel     `json:"post"`     // 帖子信息
	*CategoryModel `json:"category"` // 社区信息
}
