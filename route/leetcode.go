package route

import (
	"go-get-offer/leetcode/leetcode1"
	"go-get-offer/leetcode/leetcode11"
	"go-get-offer/leetcode/leetcode128"
	"go-get-offer/leetcode/leetcode15"
	"go-get-offer/leetcode/leetcode283"
	"go-get-offer/leetcode/leetcode3"
	"go-get-offer/leetcode/leetcode322"
	"go-get-offer/leetcode/leetcode46"
	"go-get-offer/leetcode/leetcode51"
	"go-get-offer/leetcode/leetcode54"
	"go-get-offer/leetcode/leetcode56"
	"go-get-offer/leetcode/leetcode560"
	"go-get-offer/leetcode/leetcode752"
	"go-get-offer/leetcode/leettree"
)

/*
leetcode
*/
func Leetcode() {
	// 小抄
	Setting.handlerMap["leetcode1"] = leetcode1.Run          // 两数之和
	Setting.handlerMap["leetcode111"] = leettree.Leetcode111 // 二叉树的最小深度
	Setting.handlerMap["leetcode322"] = leetcode322.Run      // 零钱兑换
	Setting.handlerMap["leetcode46"] = leetcode46.Run        // 全排列
	Setting.handlerMap["leetcode51"] = leetcode51.Run        // N 皇后
	Setting.handlerMap["leetcode56"] = leetcode56.Run        // 合并区间
	Setting.handlerMap["leetcode752"] = leetcode752.Run      // 打开转盘锁
	Setting.handlerMap["leetcode15"] = leetcode15.Run        // 三数之和
	Setting.handlerMap["leetcode3"] = leetcode3.Run          // 无重复字符的最长子串
	Setting.handlerMap["leetcode560"] = leetcode560.Run      // 和为 K 的子数组

	// 随缘
	Setting.handlerMap["leetcode54"] = leetcode54.Run        // 螺旋矩阵
	Setting.handlerMap["leetcode199"] = leettree.Leetcode199 // 二叉树的右视图
	Setting.handlerMap["leetcode654"] = leettree.Leetcode654 // 最大二叉树
	Setting.handlerMap["leetcode283"] = leetcode283.Run      // 移动零
	Setting.handlerMap["leetcode128"] = leetcode128.Run      // 最长连续序列
	Setting.handlerMap["leetcode11"] = leetcode11.Run        // 盛最多水的容器
}
