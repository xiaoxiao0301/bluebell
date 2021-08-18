package dict

// ctx 相关

// ContextUserIdKey 用户idKey
const ContextUserIdKey = "userId"

// redis 相关
const (
	KeyPrefix             = "bluebell:"              // 缓存前缀
	UserKey               = "user:"                  // 用户key => user:id
	CategoryPostsCountKey = "category:post:numbers:" // 当前分类下帖子总数
	CategoryKeyPrefix     = "category:"              // 分类key => category:id
	PostKeyPrefix         = "post:"                  // 帖子前缀
	PostTimeKey           = "time"                   // 按照帖子发表时间存储 zSet
	UserVotedPostKey      = "user:vote:"             // 记录用户给某个帖子投票 zSet
	PostVotedScore        = "score"                  // 记录帖子的得分
)

func GetSaveUserKey(userId string) string {
	return KeyPrefix + UserKey + userId
}

func GetSaveCategoryKey(categoryId string) string {
	return KeyPrefix + CategoryKeyPrefix + categoryId
}

func GetSaveCategoryPostsCountKey(categoryId string) string {
	return KeyPrefix + CategoryPostsCountKey + categoryId
}

func GetSavePostTimeKey() string {
	return KeyPrefix + PostKeyPrefix + PostTimeKey
}

func GetUserVotedPostKey(postId string) string {
	return KeyPrefix + PostKeyPrefix + UserVotedPostKey + postId
}

func GetSavePostScoreKey() string {
	return KeyPrefix + PostKeyPrefix + PostVotedScore
}
