package xerr

import "fmt"

var message = map[string]map[uint32]string{}
var defaultLang = "en"

func init() {
	message[defaultLang] = map[uint32]string{
		OK:                  "SUCCESS",
		SERVER_COMMON_ERROR: "Network abnormality, please try again later",
		REUQEST_PARAM_ERROR: "Parameter error",
		TOKEN_EXPIRE_ERROR:  "Token has expired",
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
		defaultLang = "en"
	}
	if msg, ok := message[defaultLang][errcode]; ok {
		return fmt.Sprintf(msg, args...)
	} else {
		return ""
	}
}

// 携带客户端语言的错误提示
func GetErrMsgLang(lang string, errCode uint32, args ...any) string {
	if message[lang] == nil {
		lang = defaultLang
	}
	if msg, ok := message[lang][errCode]; ok {
		return fmt.Sprintf(msg, args...)
	} else {
		return ""
	}
}
