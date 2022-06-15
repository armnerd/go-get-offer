# 搜索

* 数字在升序数组中出现的次数
* 二维数组中的查找
* 旋转数组的最小数字
* 数字序列中某一位的数字

# 53. 数字在升序数组中出现的次数

## 题目

给定一个长度为 n 的非降序数组和一个非负数整数 k ，要求统计 k 在数组中出现的次数

要求：空间复杂度 O(1)，时间复杂度 O(logn)

## 示例

输入：[1,2,3,3,3,3,4,5],3

返回值：4

## 思路

* 暴力遍历是可以，但是面试官不开心
* 利用数组升序的特点使用二分查找
* 目标值如果有多个，肯定是连在一起
* 首先找到一个目标，再顺着两头数
* 或者可以先找上下边界
* 上界定义为：不管目标值存在与否，都指向大于目标值的第一个值
* 下界定义为：如果存在目标值，则指向第一个目标值，否则，如果不存在， 则指向大于目标值的第一个值
* 最后的结果就是：上界下标 - 上界下标

## 实现

```go
func GetNumberOfK(data []int, k int) int {
	cur := 0
	end := len(data)
	for {
		if cur >= end {
			break
		}
		mid := cur + (end-cur)/2
		if data[mid] < k {
			cur = mid + 1
		} else {
			end = mid
		}
	}
	lbound := cur
	cur = 0
	end = len(data)
	for {
		if cur >= end {
			break
		}
		mid := cur + (end-cur)/2
		if data[mid] <= k {
			cur = mid + 1
		} else {
			end = mid
		}
	}
	rbound := cur
	return rbound - lbound
}
```

------

# 4. 二维数组中的查找

## 题目

在一个二维数组中（每个一维数组的长度相同），每一行都按照从左到右递增的顺序排序，每一列都按照从上到下递增的顺序排序。请完成一个函数，输入这样的一个二维数组和一个整数，判断数组中是否含有该整数。

```
[
 [1,2,8,9],
 [2,4,9,12],
 [4,7,10,13],
 [6,8,11,15]
]
```

给定 target = 7，返回 true。
给定 target = 3，返回 false。

## 思路

> 二分查找 / 分治

* 从右上角(或者左下角)开始查找
* 如果 target 更大，指针下移
* 如果 target 更小，指针左移

## 复杂度

* 时间复杂度：O(m+n) ，其中 m 为行数，n 为列数，最坏情况下，需要遍历 m+n 次。
* 空间复杂度：O(1)

## 实现

```go
func Find(target int, array [][]int) bool {
	// 参数检查
	var m, n int
	if m = len(array); m == 0 {
		return false
	}
	if n = len(array[0]); n == 0 {
		return false
	}

	// 行与列
	r := 0
	c := n - 1
	for r < m && c >= 0 {
		// 二分查找
		if target == array[r][c] {
			return true
		} else if target > array[r][c] {
			r++
		} else {
			c--
		}
	}
	return false
}
```

------

# 11. 旋转数组的最小数字

## 题目

把一个数组最开始的若干个元素搬到数组的末尾，我们称之为数组的旋转。
输入一个非递减排序的数组的一个旋转，输出旋转数组的最小元素。
NOTE：给出的所有元素都大于0，若数组大小为0，请返回0。

## 示例

输入 [3,4,5,1,2]
返回值 1

## 思路

二分查找，关键在于旋转前的数组非递减排序，即递增但是会有相同的连续值

## 实现

```go
func MinNumberInRotateArray(rotateArray []int) int {
	if len(rotateArray) == 0 {
		return 0
	}
	low := 0
	high := len(rotateArray) - 1
	for high > low {
		mid := (high + low) / 2
		if rotateArray[mid] > rotateArray[high] {
			// 中位大于高位，低位是中位后一位
			low = mid + 1
		} else if rotateArray[mid] < rotateArray[high] {
			// 中位小于高位，高位置为当前中位
			high = mid
		} else {
			if rotateArray[high-1] > rotateArray[high] {
				// 高位前一位大于高位，最小值即使当前高位
				low = high
				break
			}
			// 连续相同值的情况，高位向前推进一步
			high -= 1
		}
	}
	return rotateArray[low]
}
```

------

# 44. 数字序列中某一位的数字

## 题目

数字以 0123456789101112131415... 的格式作为一个字符序列，在这个序列中第 2 位（从下标 0 开始计算）是 2 ，第 10 位是 1 ，第 13 位是 1 ，以此类题，请你输出第 n 位对应的数字。

## 思路

* 首先确定该数字是属于几位数的
* 确定该数字是哪个数
* 确定是该数中哪一位

## 实现

```go
import "strconv"

func findNthDigit(n int) int {
	start, digit, count := 1, 1, 9
	for n > count {
		n -= count
		start *= 10
		digit += 1
		count = start * digit * 9
	}
	num := start + (n-1)/digit
	res, _ := strconv.Atoi(string(strconv.Itoa(num)[(n-1)%digit]))
	return res
}
```