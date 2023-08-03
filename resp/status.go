package resp

var (
	SUCCESS       = 1000
	NOAUTH        = 401
	PARAMERROR    = 1001
	CUSTOMIZERROR = 100
	TOKENERROR    = 9001
	TOKENEXPIRE   = 9002
)

var StatusMessage = map[int]string{
	1000: "成功",
	100:  "",
	400:  "账号或密码错误",
	401:  "没有权限",
	402:  "意外错误",
	1001: "参数有误",
	9001: "token鉴权失败",
	9002: "token过期",
	9003: "验证码错误",
	9004: "手机号格式有误",
	9005: "用户尚未注册",
	9006: "登录信息已经过期，请重新登录",
}
