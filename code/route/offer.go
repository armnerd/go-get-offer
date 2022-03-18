package route

import (
	"code/offer/link"
	"code/offer/question10"
	"code/offer/question11"
	"code/offer/question12"
	"code/offer/question13"
	"code/offer/question14"
	"code/offer/question15"
	"code/offer/question16"
	"code/offer/question17"
	"code/offer/question19"
	"code/offer/question20"
	"code/offer/question29"
	"code/offer/question3"
	"code/offer/question30"
	"code/offer/question31"
	"code/offer/question38"
	"code/offer/question4"
	"code/offer/question40"
	"code/offer/question41"
	"code/offer/question42"
	"code/offer/question43"
	"code/offer/question44"
	"code/offer/question45"
	"code/offer/question46"
	"code/offer/question47"
	"code/offer/question48"
	"code/offer/question49"
	"code/offer/question5"
	"code/offer/question50"
	"code/offer/question51"
	"code/offer/question53"
	"code/offer/question56"
	"code/offer/question57"
	"code/offer/question58"
	"code/offer/question59"
	"code/offer/question61"
	"code/offer/question62"
	"code/offer/question63"
	"code/offer/question64"
	"code/offer/question65"
	"code/offer/question66"
	"code/offer/question67"
	"code/offer/question69"
	"code/offer/question70"
	"code/offer/question71"
	"code/offer/question73"
	"code/offer/question74"
	"code/offer/question75"
	"code/offer/question81"
	"code/offer/question83"
	"code/offer/question85"
	"code/offer/question9"
	"code/offer/tree"
)

/*
	剑指offer
*/
func Offer() {
	// 链表
	Setting.handlerMap["offer6"] = link.Question6   // 从尾到头打印链表
	Setting.handlerMap["offer24"] = link.Question24 // 反转链表
	Setting.handlerMap["offer25"] = link.Question25 // 合并两个排序的链表
	Setting.handlerMap["offer52"] = link.Question52 // 两个链表的第一个公共结点
	Setting.handlerMap["offer23"] = link.Question23 // 链表中环的入口结点
	Setting.handlerMap["offer22"] = link.Question22 // 链表中倒数最后k个结点
	Setting.handlerMap["offer35"] = link.Question35 // 复杂链表的复制
	Setting.handlerMap["offer76"] = link.Question76 // 删除链表中重复的结点
	Setting.handlerMap["offer18"] = link.Question18 // 删除链表的节点

	// 树
	Setting.handlerMap["offer55"] = tree.Question55 // 二叉树的深度
	Setting.handlerMap["offer77"] = tree.Question77 // 按之字形顺序打印二叉树
	Setting.handlerMap["offer54"] = tree.Question54 // 二叉搜索树的第k个结点
	Setting.handlerMap["offer7"] = tree.Question7   // 重建二叉树
	Setting.handlerMap["offer26"] = tree.Question26 // 树的子结构
	Setting.handlerMap["offer27"] = tree.Question27 // 二叉树的镜像
	Setting.handlerMap["offer32"] = tree.Question32 // 从上往下打印二叉树
	Setting.handlerMap["offer33"] = tree.Question33 // 二叉搜索树的后序遍历序列
	Setting.handlerMap["offer82"] = tree.Question82 // 二叉树中和为某一值的路径(一)
	Setting.handlerMap["offer34"] = tree.Question34 // 二叉树中和为某一值的路径(二)
	Setting.handlerMap["offer36"] = tree.Question36 // 二叉搜索树与双向链表
	Setting.handlerMap["offer79"] = tree.Question79 // 判断是不是平衡二叉树
	Setting.handlerMap["offer8"] = tree.Question8   // 二叉树的下一个结点
	Setting.handlerMap["offer28"] = tree.Question28 // 对称的二叉树
	Setting.handlerMap["offer78"] = tree.Question78 // 把二叉树打印成多行
	Setting.handlerMap["offer37"] = tree.Question37 // 序列化二叉树
	Setting.handlerMap["offer84"] = tree.Question84 // 二叉树中和为某一值的路径(三)
	Setting.handlerMap["offer86"] = tree.Question86 // 在二叉树中找到两个节点的最近公共祖先
	Setting.handlerMap["offer68"] = tree.Question68 // 二叉搜索树的最近公共祖先

	// 队列 & 栈
	Setting.handlerMap["offer9"] = question9.Run   // 用两个栈实现队列
	Setting.handlerMap["offer30"] = question30.Run // 包含min函数的栈
	Setting.handlerMap["offer31"] = question31.Run // 栈的压入、弹出序列
	Setting.handlerMap["offer73"] = question73.Run // 翻转单词序列
	Setting.handlerMap["offer59"] = question59.Run // 滑动窗口的最大值

	// 搜索算法
	Setting.handlerMap["offer53"] = question53.Run // 数字在升序数组中出现的次数
	Setting.handlerMap["offer4"] = question4.Run   // 二维数组中的查找
	Setting.handlerMap["offer11"] = question11.Run // 旋转数组的最小数字
	Setting.handlerMap["offer38"] = question38.Run // 字符串的排列
	Setting.handlerMap["offer44"] = question44.Run // 数字序列中某一位的数字

	// 动态规划
	Setting.handlerMap["offer42"] = question42.Run // 连续子数组的最大和
	Setting.handlerMap["offer85"] = question85.Run // 连续子数组的最大和(二)
	Setting.handlerMap["offer69"] = question69.Run // 跳台阶
	Setting.handlerMap["offer10"] = question10.Run // 斐波那契数列
	Setting.handlerMap["offer19"] = question19.Run // 正则表达式匹配
	Setting.handlerMap["offer71"] = question71.Run // 跳台阶扩展问题
	Setting.handlerMap["offer70"] = question70.Run // 矩形覆盖
	Setting.handlerMap["offer63"] = question63.Run // 买卖股票的最好时机(一)
	Setting.handlerMap["offer47"] = question47.Run // 礼物的最大价值
	Setting.handlerMap["offer48"] = question48.Run // 最长不含重复字符的子字符串
	Setting.handlerMap["offer46"] = question46.Run // 把数字翻译成字符串

	// 回溯
	Setting.handlerMap["offer12"] = question12.Run // 矩阵中的路径
	Setting.handlerMap["offer13"] = question13.Run // 机器人的运动范围

	// 排序
	Setting.handlerMap["offer3"] = question3.Run   // 数组中重复的数字
	Setting.handlerMap["offer51"] = question51.Run // 数组中的逆序对
	Setting.handlerMap["offer40"] = question40.Run // 最小的K个数
	Setting.handlerMap["offer41"] = question41.Run // 数据流中的中位数

	// 位运算
	Setting.handlerMap["offer65"] = question65.Run // 不用加减乘除做加法
	Setting.handlerMap["offer15"] = question15.Run // 二进制中1的个数
	Setting.handlerMap["offer16"] = question16.Run // 数值的整数次方
	Setting.handlerMap["offer56"] = question56.Run // 数组中只出现一次的两个数字
	Setting.handlerMap["offer64"] = question64.Run // 求1+2+3+...+n

	// 模拟
	Setting.handlerMap["offer29"] = question29.Run // 顺时针打印矩阵
	Setting.handlerMap["offer61"] = question61.Run // 扑克牌顺子
	Setting.handlerMap["offer67"] = question67.Run // 把字符串转换成整数(atoi)
	Setting.handlerMap["offer20"] = question20.Run // 表示数值的字符串

	// 其他算法
	Setting.handlerMap["offer66"] = question66.Run // 构建乘积数组
	Setting.handlerMap["offer50"] = question50.Run // 第一个只出现一次的字符
	Setting.handlerMap["offer5"] = question5.Run   // 替换空格
	Setting.handlerMap["offer21"] = question5.Run  // 调整数组顺序使奇数位于偶数前面(一)
	Setting.handlerMap["offer39"] = question5.Run  // 数组中出现次数超过一半的数字
	Setting.handlerMap["offer43"] = question43.Run // 整数中1出现的次数（从1到n整数中1出现的次数）
	Setting.handlerMap["offer45"] = question45.Run // 把数组排成最小的数
	Setting.handlerMap["offer49"] = question49.Run // 丑数
	Setting.handlerMap["offer74"] = question74.Run // 和为S的连续正数序列
	Setting.handlerMap["offer57"] = question57.Run // 和为S的两个数字
	Setting.handlerMap["offer58"] = question58.Run // 左旋转字符串
	Setting.handlerMap["offer62"] = question62.Run // 孩子们的游戏(圆圈中最后剩下的数)
	Setting.handlerMap["offer75"] = question75.Run // 字符流中第一个不重复的字符
	Setting.handlerMap["offer14"] = question14.Run // 剪绳子
	Setting.handlerMap["offer81"] = question81.Run // 调整数组顺序使奇数位于偶数前面(二)
	Setting.handlerMap["offer83"] = question83.Run // 剪绳子（进阶版）
	Setting.handlerMap["offer17"] = question17.Run // 打印从1到最大的n位数
}
