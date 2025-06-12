package stateCode

const (
	LangZh   = 111
	LangEn   = 112
	LangZhTw = 113

	CommonSuccess      = 0
	CommonErrServerErr = 1000
	ParamterEmptyErr   = 1001

	TokeEmptyErr = 1101
	TokenErr     = 1102

	PNameEmpty   = 1201
	ChainIdEmpty = 1202
	ChainIdErr   = 1203

	NameEmptyErr      = 1301
	PasswordEmptyErr  = 1302
	NameOrPasswordErr = 1303
)

var Msg = map[int]map[int]string{
	CommonSuccess: {
		LangZh:   "成功",
		LangZhTw: "成功",
		LangEn:   "success",
	},
	CommonErrServerErr: {
		LangZh:   "服务器繁忙，请稍后重试",
		LangZhTw: "服務器繁忙，請稍後重試",
		LangEn:   "server is busy, please try again later",
	},
	ParamterEmptyErr: {
		LangZh:   "参数不能为空",
		LangZhTw: "参数不能為空",
		LangEn:   "parameter is empty",
	},
	TokeEmptyErr: {
		LangZh:   "token 不能为空",
		LangZhTw: "token 不能為空",
		LangEn:   "token required",
	},
	TokenErr: {
		LangZh:   "token错误",
		LangZhTw: "token錯誤",
		LangEn:   "token invalid",
	},
	PNameEmpty: {
		LangZh:   "sp_name 不能为空",
		LangZhTw: "sp_name 不能為空",
		LangEn:   "sp_name required",
	},
	ChainIdEmpty: {
		LangZh:   "chain_id 不能为空",
		LangZhTw: "chain_id 不能為空",
		LangEn:   "chain_id required",
	},
	ChainIdErr: {
		LangZh:   "chain_id 错误",
		LangZhTw: "chain_id 錯誤",
		LangEn:   "chain_id error",
	},
	NameEmptyErr: {
		LangZh:   "name 不能为空",
		LangZhTw: "name 不能為空",
		LangEn:   "name required",
	},
	PasswordEmptyErr: {
		LangZh:   "password 不能为空",
		LangZhTw: "password 不能為空",
		LangEn:   "password required",
	},
	NameOrPasswordErr: {
		LangZh:   "用户名或密码错误",
		LangZhTw: "用戶名或密碼錯誤",
		LangEn:   "name or password error",
	},
}

func GetMsg(c int, lang int) string {
	_, ok := Msg[c]
	if ok {
		msg, ok := Msg[c][lang]
		if ok {
			return msg
		}
		return Msg[CommonErrServerErr][lang]
	}
	return Msg[CommonErrServerErr][lang]
}
