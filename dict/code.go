package dict

type ResponseCode int

const (
	CodeSuccess ResponseCode = 1000 + iota
	CodeInvalidParam
	CodeUserExists
	CodeUserNotExists
	CodeInvalidPassword
	CodeServerBusy
	CodeNetWorkBusy
	CodeNeedLogin
	CodeInvalidToken
	CodeInvalidRefreshToken
	CodeValidAccessToken
	CodeNotQueryResult
	CodeVotedEqualResult
	CodeVoteTimeExpires
)

var codeText = map[ResponseCode]string{
	CodeSuccess:             "success",
	CodeInvalidParam:        "参数错误",
	CodeUserExists:          "用户已存在",
	CodeUserNotExists:       "用户不存在",
	CodeInvalidPassword:     "密码错误",
	CodeServerBusy:          "服务繁忙",
	CodeNetWorkBusy:         "网络繁忙，稍后再试",
	CodeNeedLogin:           "请登录",
	CodeInvalidToken:        "无效token",
	CodeInvalidRefreshToken: "无效的refresh_token",
	CodeValidAccessToken:    "access_token未过期",
	CodeNotQueryResult:      "暂无数据",
	CodeVotedEqualResult:    "不能投重复票",
	CodeVoteTimeExpires:     "投票时间已过",
}

// Message 返回错误码对应的文字信息
func (c ResponseCode) Message() string {
	message, ok := codeText[c]
	if !ok {
		return codeText[CodeServerBusy]
	}
	return message
}
