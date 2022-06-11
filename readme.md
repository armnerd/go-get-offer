# go-get-offer

> 用 Go 实现牛客网《剑指offer》

## Feature

* 可运行 { 可调试查看过程 }
* 可打印 { 面试前手写突击 }

## Test

```bash
cd test/leetcode

// 都拉出来溜溜
go test -v

// 测试一个函数
go test -run TestTree

// 可视化
goconvey
```

### 牛客 { 81 }

* https://www.nowcoder.com/ta/coding-interviews

### Leetcode { 84 }

* [出门右转](./doc/leetcode/readme.md)

### Go 相关 { 3 }

* [起飞](./doc/language/readme.md)

## Get started

> 链表

* [从尾到头打印链表](./doc/offer/6.md) 
* [反转链表](./doc/offer/24.md) 
* [合并两个排序的链表](./doc/offer/25.md) 
* [两个链表的第一个公共结点](./doc/offer/52.md) 
* [链表中环的入口结点](./doc/offer/23.md) 
* [链表中倒数最后k个结点](./doc/offer/22.md) 
* [复杂链表的复制](./doc/offer/35.md) 
* [删除链表中重复的结点](./doc/offer/76.md) 
* [删除链表的节点](./doc/offer/18.md)

> 树

* [二叉树的深度](./doc/offer/55.md) 
* [按之字形顺序打印二叉树](./doc/offer/77.md) 
* [二叉搜索树的第k个结点](./doc/offer/54.md) 
* [重建二叉树](./doc/offer/7.md) 
* [树的子结构](./doc/offer/26.md) 
* [二叉树的镜像](./doc/offer/27.md) 
* [从上往下打印二叉树](./doc/offer/32.md) 
* [二叉搜索树的后序遍历序列](./doc/offer/33.md) 
* [二叉树中和为某一值的路径(一)](./doc/offer/82.md) 
* [二叉树中和为某一值的路径(二)](./doc/offer/34.md) 
* [二叉搜索树与双向链表](./doc/offer/36.md) 
* [判断是不是平衡二叉树](./doc/offer/79.md) 
* [二叉树的下一个结点](./doc/offer/8.md) 
* [对称的二叉树](./doc/offer/28.md) 
* [把二叉树打印成多行](./doc/offer/78.md) 
* [序列化二叉树](./doc/offer/37.md) 
* [二叉树中和为某一值的路径(三)](./doc/offer/84.md)
* [在二叉树中找到两个节点的最近公共祖先](./doc/offer/86.md)
* [二叉搜索树的最近公共祖先](./doc/offer/68.md) 

> 队列 & 栈

* [用两个栈实现队列](./doc/offer/9.md) 
* [包含min函数的栈](./doc/offer/30.md) 
* [栈的压入、弹出序列](./doc/offer/31.md) 
* [翻转单词序列](./doc/offer/73.md) 
* [滑动窗口的最大值](./doc/offer/59.md)

> 搜索算法

* [数字在升序数组中出现的次数](./doc/offer/53.md) 
* [二维数组中的查找](./doc/offer/4.md) 
* [旋转数组的最小数字](./doc/offer/11.md) 
* [字符串的排列](./doc/offer/38.md) 
* [数字序列中某一位的数字](./doc/offer/44.md)

> 动态规划

* [连续子数组的最大和](./doc/offer/42.md) 
* [连续子数组的最大和(二)](./doc/offer/85.md) 
* [跳台阶](./doc/offer/69.md) 
* [斐波那契数列](./doc/offer/10.md) 
* [正则表达式匹配](./doc/offer/19.md) 
* [跳台阶扩展问题](./doc/offer/71.md) 
* [矩形覆盖](./doc/offer/70.md) 
* [买卖股票的最好时机(一)](./doc/offer/63.md) 
* [礼物的最大价值](./doc/offer/47.md) 
* [最长不含重复字符的子字符串](./doc/offer/48.md) 
* [把数字翻译成字符串](./doc/offer/46.md)

> 回溯

* [矩阵中的路径](./doc/offer/12.md) 
* [机器人的运动范围](./doc/offer/13.md) 

> 排序

* [数组中重复的数字](./doc/offer/3.md) 
* [数组中的逆序对](./doc/offer/51.md) 
* [最小的K个数](./doc/offer/40.md) 
* [数据流中的中位数](./doc/offer/41.md) 

> 位运算

* [不用加减乘除做加法](./doc/offer/65.md) 
* [二进制中1的个数](./doc/offer/15.md) 
* [数值的整数次方](./doc/offer/16.md) 
* [数组中只出现一次的两个数字](./doc/offer/56.md) 
* [求1+2+3+...+n](./doc/offer/64.md) 

> 模拟

* [顺时针打印矩阵](./doc/offer/29.md) 
* [扑克牌顺子](./doc/offer/61.md) 
* [把字符串转换成整数(atoi)](./doc/offer/67.md) 
* [表示数值的字符串](./doc/offer/20.md)

> 其他算法

* [构建乘积数组](./doc/offer/66.md) 
* [第一个只出现一次的字符](./doc/offer/50.md) 
* [替换空格](./doc/offer/5.md) 
* [调整数组顺序使奇数位于偶数前面(一)](./doc/offer/21.md) 
* [数组中出现次数超过一半的数字](./doc/offer/39.md) 
* [整数中1出现的次数（从1到n整数中1出现的次数）](./doc/offer/43.md) 
* [把数组排成最小的数](./doc/offer/45.md) 
* [丑数](./doc/offer/49.md) 
* [和为S的连续正数序列](./doc/offer/74.md) 
* [和为S的两个数字](./doc/offer/57.md) 
* [左旋转字符串](./doc/offer/58.md) 
* [孩子们的游戏(圆圈中最后剩下的数)](./doc/offer/62.md) 
* [字符流中第一个不重复的字符](./doc/offer/75.md) 
* [剪绳子](./doc/offer/14.md) 
* [调整数组顺序使奇数位于偶数前面(二)](./doc/offer/81.md) 
* [剪绳子（进阶版）](./doc/offer/83.md) 
* [打印从1到最大的n位数](./doc/offer/17.md)