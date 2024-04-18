# go-get-offer

> 用 Go 实现牛客网《剑指offer》

## Feature

* 可运行 { 可调试查看过程 }
* 可打印 { 面试前手写突击 }

## Run

```bash
go run main.go offer68
```

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

### Leetcode { 85 }

* [出门右转](./doc/leetcode/readme.md)

### Go 相关 { 3 }

* [起飞](./doc/language/readme.md)

## Get started

> 链表

* [从尾到头打印链表](./doc/offer/link/6.md) 
* [反转链表](./doc/offer/link/24.md) 
* [合并两个排序的链表](./doc/offer/link/25.md) 
* [两个链表的第一个公共结点](./doc/offer/link/52.md) 
* [链表中环的入口结点](./doc/offer/link/23.md) 
* [链表中倒数最后k个结点](./doc/offer/link/22.md) 
* [复杂链表的复制](./doc/offer/link/35.md) 
* [删除链表中重复的结点](./doc/offer/link/76.md) 
* [删除链表的节点](./doc/offer/link/18.md)

> 树

* [二叉树的深度](./doc/offer/tree/55.md) 
* [按之字形顺序打印二叉树](./doc/offer/tree/77.md) 
* [二叉搜索树的第k个结点](./doc/offer/tree/54.md) 
* [重建二叉树](./doc/offer/tree/7.md) 
* [树的子结构](./doc/offer/tree/26.md) 
* [二叉树的镜像](./doc/offer/tree/27.md) 
* [从上往下打印二叉树](./doc/offer/tree/32.md) 
* [二叉搜索树的后序遍历序列](./doc/offer/tree/33.md) 
* [二叉树中和为某一值的路径(一)](./doc/offer/tree/82.md) 
* [二叉树中和为某一值的路径(二)](./doc/offer/tree/34.md) 
* [二叉搜索树与双向链表](./doc/offer/tree/36.md) 
* [判断是不是平衡二叉树](./doc/offer/tree/79.md) 
* [二叉树的下一个结点](./doc/offer/tree/8.md) 
* [对称的二叉树](./doc/offer/tree/28.md) 
* [把二叉树打印成多行](./doc/offer/tree/78.md) 
* [序列化二叉树](./doc/offer/tree/37.md) 
* [二叉树中和为某一值的路径(三)](./doc/offer/tree/84.md)
* [在二叉树中找到两个节点的最近公共祖先](./doc/offer/tree/86.md)
* [二叉搜索树的最近公共祖先](./doc/offer/tree/68.md) 

> 队列 & 栈

* [用两个栈实现队列](./doc/offer/queueAndStack/9.md) 
* [包含min函数的栈](./doc/offer/queueAndStack/30.md) 
* [栈的压入、弹出序列](./doc/offer/queueAndStack/31.md) 
* [翻转单词序列](./doc/offer/queueAndStack/73.md) 
* [滑动窗口的最大值](./doc/offer/queueAndStack/59.md)

> 搜索算法

* [数字在升序数组中出现的次数](./doc/offer/search/53.md) 
* [二维数组中的查找](./doc/offer/search/4.md) 
* [旋转数组的最小数字](./doc/offer/search/11.md) 
* [字符串的排列](./doc/offer/search/38.md) 
* [数字序列中某一位的数字](./doc/offer/search/44.md)

> 动态规划

* [连续子数组的最大和](./doc/offer/dynamicProgramming/42.md) 
* [连续子数组的最大和(二)](./doc/offer/dynamicProgramming/85.md) 
* [跳台阶](./doc/offer/dynamicProgramming/69.md) 
* [斐波那契数列](./doc/offer/dynamicProgramming/10.md) 
* [正则表达式匹配](./doc/offer/dynamicProgramming/19.md) 
* [跳台阶扩展问题](./doc/offer/dynamicProgramming/71.md) 
* [矩形覆盖](./doc/offer/dynamicProgramming/70.md) 
* [买卖股票的最好时机(一)](./doc/offer/dynamicProgramming/63.md) 
* [礼物的最大价值](./doc/offer/dynamicProgramming/47.md) 
* [最长不含重复字符的子字符串](./doc/offer/dynamicProgramming/48.md) 
* [把数字翻译成字符串](./doc/offer/dynamicProgramming/46.md)

> 回溯

* [矩阵中的路径](./doc/offer/other/12.md) 
* [机器人的运动范围](./doc/offer/other/13.md) 

> 排序

* [数组中重复的数字](./doc/offer/sort/3.md) 
* [数组中的逆序对](./doc/offer/sort/51.md) 
* [最小的K个数](./doc/offer/sort/40.md) 
* [数据流中的中位数](./doc/offer/sort/41.md) 

> 位运算

* [不用加减乘除做加法](./doc/offer/other/65.md) 
* [二进制中1的个数](./doc/offer/other/15.md) 
* [数值的整数次方](./doc/offer/other/16.md) 
* [数组中只出现一次的两个数字](./doc/offer/other/56.md) 
* [求1+2+3+...+n](./doc/offer/other/64.md) 

> 模拟

* [顺时针打印矩阵](./doc/offer/other/29.md) 
* [扑克牌顺子](./doc/offer/other/61.md) 
* [把字符串转换成整数(atoi)](./doc/offer/other/67.md) 
* [表示数值的字符串](./doc/offer/other/20.md)

> 其他算法

* [构建乘积数组](./doc/offer/other/66.md) 
* [第一个只出现一次的字符](./doc/offer/other/50.md) 
* [替换空格](./doc/offer/other/5.md) 
* [调整数组顺序使奇数位于偶数前面(一)](./doc/offer/other/21.md) 
* [数组中出现次数超过一半的数字](./doc/offer/other/39.md) 
* [整数中1出现的次数（从1到n整数中1出现的次数）](./doc/offer/other/43.md) 
* [把数组排成最小的数](./doc/offer/other/45.md) 
* [丑数](./doc/offer/other/49.md) 
* [和为S的连续正数序列](./doc/offer/other/74.md) 
* [和为S的两个数字](./doc/offer/other/57.md) 
* [左旋转字符串](./doc/offer/other/58.md) 
* [孩子们的游戏(圆圈中最后剩下的数)](./doc/offer/other/62.md) 
* [字符流中第一个不重复的字符](./doc/offer/other/75.md) 
* [剪绳子](./doc/offer/other/14.md) 
* [调整数组顺序使奇数位于偶数前面(二)](./doc/offer/other/81.md) 
* [剪绳子（进阶版）](./doc/offer/other/83.md) 
* [打印从1到最大的n位数](./doc/offer/other/17.md)