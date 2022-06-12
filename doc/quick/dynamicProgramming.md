# 动态规划

* 连续子数组的最大和
* 跳台阶
* 斐波那契数列
* 矩形覆盖
* 买卖股票的最好时机(一)
* 礼物的最大价值
* 最长不含重复字符的子字符串

# 42. 连续子数组的最大和

## 题目

输入一个长度为n的整型数组array，数组中的一个或连续多个整数组成一个子数组。求所有子数组的和的最大值。

要求:时间复杂度为 O(n)，空间复杂度为 O(n)
进阶:时间复杂度为 O(n)，空间复杂度为 O(1)

## 示例

输入：[1,-2,3,10,-4,7,2,-5]

返回值：18

说明：经分析可知，输入数组的子数组[3,10,-4,7,2]可以求得最大和为18 

## 思路

* 状态定义：dp[i]表示以i结尾的连续子数组的最大和。所以最终要求dp[n-1]
* 状态转移方程：dp[i] = max(array[i], dp[i-1]+array[i])

> 但是这个公式怎么来的呢？

* 假设中间两个紧挨着的位置 m 和 n， m 在 n 的前边
* 假如 m 已经和前边元素都遍历相加对比过，知道了以 m 结尾的连续子数组的最大和 为 x
* 那么以 n 结尾的连续子数组的最大和就是 n + x 和 n 中的较大值
* 当知道了以数组中每个元素结尾的子数组的连续子数组的最大和后，再取其中最大的

## 实现

```go
func FindGreatestSumOfSubArray(array []int) int {
	size := len(array)
    dp := 0
	ret := array[0]
	for i := 1; i <= size; i++ {
		dp = max(array[i-1], dp+array[i-1])
		ret = max(ret, dp)
	}
	return ret
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
```

------

# 69. 跳台阶

> 就是个斐波那契数列

## 题目

一只青蛙一次可以跳上 1 级台阶，也可以跳上 2级。求该青蛙跳上一个n级的台阶总共有多少种跳法（先后次序不同算不同的结果）。

## 思路

* 假设对于第 n 级台阶，总共有 f(n) 种跳法.
* 那么 f(n) = f(n-1) + f(n-2)，其中 f(1) = 1, f(2) = 2
* 其中肯定有重复运算的，搞个 map 备忘录存着

## 复杂度

* 如果单纯递归时间复杂度是 O(2^n)
* 加上记忆化搜索后时间复杂度是 O(n)，因为没有了重复计算
* 动态规划可以将空间复杂度降到 O(n) 甚至 O(1)

## 实现

```go
// 纯递归
func jumpFloor1(number int) int {
	if number <= 1 {
		return 1
	}
	return jumpFloor1(number-1) + jumpFloor1(number-2)
}

// 备忘录
func jumpFloor2(number int) int {
	note := make([]int, number)
	return fib(number, note)
}

func fib(number int, note []int) int {
	if number <= 1 {
		return 1
	}

	if note[number-1] != 0 {
		return note[number-1]
	}

	res := fib(number-1, note) + fib(number-2, note)
	note[number-1] = res
	return res
}

// 备忘录升级为动态规划, 不用递归的过程, 直接从子树求得答案, 过程是从下往上
func jumpFloor3(number int) int {
	if number <= 1 {
		return 1
	}
	note := make([]int, number)
	note[0] = 1
	note[1] = 2
	for i := 2; i < number; i++ {
		note[i] = note[i-1] + note[i-2]
	}
	return note[number-1]
}

// 究极动态规划
func jumpFloor4(number int) int {
	if number <= 1 {
		return 1
	}
	var a = 1
	var b = 1
	var c int
	for i := 2; i <= number; i++ {
		c = a + b
		a = b
		b = c
	}
	return c
}
```

------

# 10. 斐波那契数列

> 这个数列从第 3 项开始，每一项都等于前两项之和

## 题目

大家都知道斐波那契数列，现在要求输入一个整数 n，请你输出斐波那契数列的第 n 项（从 0 开始，第 0 项为 0，第 1 项是 1 ）。 n <= 39

## 思路

> 递归

F(0) = 0，F(1) = 1, F(n) = F(n - 1) + F(n - 2)

## 实现

```go
func Fibonacci(n int) int {
	f1 := 0
	f2 := 1
	if n == 0 {
		return f1
	} else if n == 1 {
		return f2
	}
	for i := 1; i < n; i++ {
		tmp := f2
		f2 = f1 + f2
		f1 = tmp
	}
	return f2
}
```

------

# 70. 矩形覆盖

## 题目

我们可以用 2*1 的小矩形横着或者竖着去覆盖更大的矩形。请问用 n 个 2*1 的小矩形无重叠地覆盖一个 2*n 的大矩形，从同一个方向看总共有多少种不同的方法？

进阶：空间复杂度 O(1)，时间复杂度 O(n)

注意：约定 n == 0 时，输出 0

## 示例

比如n=3时，2*3的矩形块有3种不同的覆盖方法(从同一个方向看)：

输入描述：2*1的小矩形的总个数n

返回值描述：覆盖一个2*n的大矩形总共有多少种不同的方法(从同一个方向看)

## 思路

> 先递推下

* n=1时，显然只有一种方法
* n=2时，有2种方法
* n=3，有3中方法
* n=4,有5种方法

> 从n=3到n=4，怎么来的呢？

这里有2种情况：

* 直接在n=3的情况下，再后面中添加一个竖着的。这个很显然成立，有3种情况
* 然后横着的显然能添加到n-2的情况上，也就是在n=2后面，添加2个横着的。有2种情况

所以总结：f[n]表示2*n大矩阵的方法数
可以得出：f[n] = f[n-1] + f[n-2]，初始条件f[1] = 1, f[2] =2

## 实现

```go
func rectCover(number int) int {
	if number == 0 || number == 1 || number == 2 {
		return number
	}
	first := 1
	second := 2
	res := 0
	for i := 3; i <= number; i++ {
		res = first + second
		first = second
		second = res
	}
	return res
}
```

------

# 63. 买卖股票的最好时机(一)

## 题目

假设你有一个数组prices，长度为n，其中prices[i]是股票在第i天的价格，请根据这个价格数组，返回买卖股票能获得的最大收益
1.你可以买入一次股票和卖出一次股票，并非每天都可以买入或卖出一次，总共只能买入和卖出一次，且买入必须在卖出的前面的某一天
2.如果不能获取到任何利润，请返回0
3.假设买入卖出均无手续费

要求：空间复杂度 O(1)，时间复杂度 O(n)

## 示例

输入：[8,9,2,5,4,7,1]

返回值：5

说明：在第3天(股票价格 = 2)的时候买入，在第6天(股票价格 = 7)的时候卖出，最大利润 = 7-2 = 5 ，不能选择在第2天买入，第3天卖出，这样就亏损7了；同时，你也不能在买入前卖出股票。

## 思路

* 你不能在买入股票前卖出股票
* 动态规划

## 实现

```go
func maxProfit(prices []int) int {
	earn := 0        // 最大收益
	min := prices[0] // 股票对低值
	for k, v := range prices {
		if v < min {
			min = v
		}
		if k == 0 {
			// 第一天没有收入
			earn = 0
		} else {
			// 当天股票值减去最低股票值就是今天的预期最大收益，再和历史的最大收益值比较，选择较大的
			if (v - min) > earn {
				earn = v - min
			}
		}
	}
	return earn
}
```

------

# 47. 礼物的最大价值

## 题目

在一个 m×n 的棋盘的每一格都放有一个礼物，每个礼物都有一定的价值（价值大于 0）。你可以从棋盘的左上角开始拿格子里的礼物，并每次向右或者向下移动一格、直到到达棋盘的右下角。给定一个棋盘及其上面的礼物的价值，请计算你最多能拿到多少价值的礼物？
如输入这样的一个二维数组，
[
[1,3,1],
[1,5,1],
[4,2,1]
]
那么路径 1→3→5→2→1 可以拿到最多价值的礼物，价值为12

## 思路

* 每次只能向下或者向上走一格，所以对于每一个单元格来说，只能通过其左边的单元格或者上边的单元格到达，因此dp数组保存走到当前grid单元格能够得到的最大价值
* 对于第i行第j列的格子而言，走到这个单元格所能得到的最大值为 dp[同列上一行] 与 dp[同行前一列] 中最大值加上本格的值
* 递推公式： max(dp[i-1][j], dp[i][j-1]) + grid[i][j]

## 实现

```go
func maxValue(grid [][]int) int {
	// 初始化最上一行
	for j := 1; j < len(grid[0]); j++ {
		grid[0][j] = grid[0][j-1] + grid[0][j]
	}
	// 初始化最左一列
	for i := 1; i < len(grid); i++ {
		grid[i][0] = grid[i-1][0] + grid[i][0]
	}
	// 逐列
	for i := 1; i < len(grid); i++ {
		// 逐行
		for j := 1; j < len(grid[0]); j++ {
			grid[i][j] = max(grid[i-1][j], grid[i][j-1]) + grid[i][j]
		}
	}
	return grid[len(grid)-1][len(grid[0])-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
```

------

# 48. 最长不含重复字符的子字符串

## 题目

请从字符串中找出一个最长的不包含重复字符的子字符串，计算该最长子字符串的长度

## 示例

> 1

输入："abcabcbb"

返回值：3

说明：因为无重复字符的最长子串是"abc"，所以其长度为 3

> 2

输入："bbbbb"

返回值：1

说明：因为无重复字符的最长子串是"b"，所以其长度为 1

> 3

输入："pwwkew"

返回值：3

说明：因为无重复字符的最长子串是 "wke"，所以其长度为 3。请注意，你的答案必须是子串的长度，"pwke" 是一个子序列，不是子串

## 思路

* 用dp[]记录状态，dp[i]表示以下标为i的字符结尾不包含重复字符的最长子字符串长度
* 再在dp中求maxdp
* 每次可以根据dp的前一个状态推导出后一个状态
* 因此可以省略dp数组，使用一个变量记录dp值，使用maxdp记录最大的dp值

## 实现

```go
func lengthOfLongestSubstring(s string) int {
	if s == "" {
		return 0
	}
	prevMaxLength, maxLength := 0, 0
	for k, v := range s {
		// 从上一位开始往前数，数到以上一位结尾的最长不重复字符串的第一个字符
		i := k - 1
		for i >= (k - prevMaxLength) {
			if string(v) == string(s[i]) {
				break
			}
			i--
		}
		prevMaxLength = k - i
		// 比较历史中最大
		if prevMaxLength > maxLength {
			maxLength = prevMaxLength
		}
	}
	return maxLength
}
```