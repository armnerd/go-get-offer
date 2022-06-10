package route

import (
	"code/leetcode/leetcode54"
	"code/leetcode/leetcode56"
	"code/leetcode/leettree"
)

/*
	leetcode
*/
func Leetcode() {
	// 小抄
	Setting.handlerMap["leetcode56"] = leetcode56.Run // 合并区间

	// 随缘
	Setting.handlerMap["leetcode54"] = leetcode54.Run        // 螺旋矩阵
	Setting.handlerMap["leetcode199"] = leettree.Leetcode199 // 二叉树的右视图
	Setting.handlerMap["leetcode654"] = leettree.Leetcode654 // 最大二叉树
}
