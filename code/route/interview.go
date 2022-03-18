package route

import (
	"code/interview"
)

/*
	面试真题
*/
func Interview() {
	// golang
	Setting.handlerMap["interview1"] = interview.Solution1 // 多态与协程
	Setting.handlerMap["interview2"] = interview.Solution2 // 交替打印字符
	Setting.handlerMap["interview3"] = interview.Solution3 // 协程顺序打印
}
