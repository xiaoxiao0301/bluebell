package dict

import "errors"

var (
	ErrorUserExists         = errors.New("用户已经存在")
	ErrorUserNotExists      = errors.New("用户暂未注册")
	ErrorNeedLogin          = errors.New("请登录")
	ErrorUserNameOrPassword = errors.New("用户名或密码错误")
	ErrorAccessTokenValid   = errors.New("token有效")
	ErrorNotQueryResult     = errors.New("无查询结果")
	ErrorVoteEqualValue     = errors.New("不能投重复票")
	ErrorVoteTimeExpires    = errors.New("投票时间已过")
)
