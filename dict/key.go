package dict

// ctx 相关
const ContextUserIdKey = "userId"

// redis 相关
const (
	KeyPrefix         = "bluebell:" // 缓存前缀
	UserKey           = "user:"     // 用户key => user:id
	CategoryKeyPrefix = "category:" // 分类key => category:id
	PostKeyPrefix     = "post:"     // 帖子前缀
	PostTimeKey       = "time"      // 按照帖子发表时间存储 zset
)

func GetSaveUserKey(userId string) string {
	return KeyPrefix + UserKey + userId
}

func GetSaveCategoryKey(categoryId string) string {
	return KeyPrefix + CategoryKeyPrefix + categoryId
}

func GetSavePostTimeKey() string {
	return KeyPrefix + PostKeyPrefix + PostTimeKey
}
