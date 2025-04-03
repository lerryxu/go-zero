package xerr

import "fmt"

var message = map[string]map[uint32]string{}
var defaultLang = "zh-CN"

func init() {
	message[defaultLang] = map[uint32]string{
		OK:                  "SUCCESS",
		SERVER_COMMON_ERROR: "服务器开小差啦,稍后再来试一试",
		REUQEST_PARAM_ERROR: "参数错误",
		TOKEN_EXPIRE_ERROR:  "Token已过期",
	}
}

// 多语言错误配置
func InitLangMsg(lang string, src map[uint32]string) {
	if message[lang] == nil {
		message[lang] = map[uint32]string{}
	}
	for k, v := range src {
		message[lang][k] = v
	}
}

// 设置默认语言
func SetLang(lang string) {
	defaultLang = lang
}

func GetErrMsg(errcode uint32, args ...any) string {
	if message[defaultLang] == nil {
		defaultLang = "zh-CN"
	}
	if msg, ok := message[defaultLang][errcode]; ok {
		return fmt.Sprintf(msg, args...)
	} else {
		return ""
	}
}
