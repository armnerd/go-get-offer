package route

import (
	"go-get-offer/language"
)

/*
	Go相关
*/
func Language() {
	Setting.handlerMap["go1"] = language.Solution1 // 多态与协程
	Setting.handlerMap["go2"] = language.Solution2 // 交替打印字符
	Setting.handlerMap["go3"] = language.Solution3 // 协程顺序打印
}
