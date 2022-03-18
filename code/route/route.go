package route

// config
type config struct {
	handlerMap map[string]interface{}
}

// Setting 全局配置
var Setting config

// init 初始化配置
func init() {
	Setting.handlerMap = make(map[string]interface{})
	Offer()     // 剑指offer
	Leetcode()  // leetcode
	Interview() // 面试真题
}
