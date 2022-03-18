# 排序

* 数组中重复的数字
* 数组中的逆序对

# 3. 数组中重复的数字

## 题目

在一个长度为n的数组里的所有数字都在0到n-1的范围内。 数组中某些数字是重复的，但不知道有几个数字是重复的。也不知道每个数字重复几次。请找出数组中任一一个重复的数字。 例如，如果输入长度为7的数组[2,3,1,0,2,5,3]，那么对应的输出是2或者3。存在不合法的输入的话输出-1

进阶：空间复杂度 O(n)，时间复杂度 O(n)

## 示例

输入：[2,3,1,0,2,5,3]

返回值：2

说明：2或3都是对的

## 思路

使用map

## 实现

```go
func duplicate(numbers []int) int {
	res := -1
	replicateSet := make(map[int]int)
	for _, v := range numbers {
		_, ok := replicateSet[v]
		if ok {
			res = v
			break
		}
		replicateSet[v] = 1
	}
	return res
}
```

------

# 51. 数组中的逆序对

## 题目

在数组中的两个数字，如果前面一个数字大于后面的数字，则这两个数字组成一个逆序对。输入一个数组,求出这个数组中的逆序对的总数P。并将P对1000000007取模的结果输出。 即输出P mod 1000000007

要求：空间复杂度 O(n)，时间复杂度 O(nlogn)

输入描述：题目保证输入的数组中没有的相同的数字

## 示例

输入：[1,2,3,4,5,6,7,0]

返回值：7

## 思路

* 首先【前面一个数字大于后面的数字】这俩数不一定紧挨着
* 然后，逆序是说 a[i] > a[j]，i < j
* 那么在排序的过程中，会把 a[i] 和 a[j] 交换过来，每交换一次计数 +1
* 暴力冒泡遍历不满足题目要求，需要用到归并排序
* 归并排序的过程就是，递归划分整个区间为基本相等的左右两个区间，然后就合并两个有序区间
* 在合并数组的时候，当发现右边的小于左边的时候，逆序对的个数 +1

## 实现

```go
func InversePairs(data []int) int {
	if len(data) < 2 {
		return 0
	}

	replica := make([]int, len(data))
	copy(replica, data)

	count := mergeSort(data, replica, 0, len(data)-1)
	return count % 1000000007
}

func mergeSort(data, copy []int, start, end int) int {
	if start == end {
		copy[start] = data[start]
		return 0
	}
	length := (end - start) / 2
	left := mergeSort(copy, data, start, start+length)
	right := mergeSort(copy, data, start+length+1, end)

	i := start + length
	j := end
	indexCopy := end
	count := 0

	for i >= start && j >= start+length+1 {
		if data[i] > data[j] {
			copy[indexCopy] = data[i]
			indexCopy--
			i--
			count += j - start - length
		} else {
			copy[indexCopy] = data[j]
			indexCopy--
			j--
		}
	}
	for i >= start {
		copy[indexCopy] = data[i]
		indexCopy--
        i--
	}
	for j >= start+length+1 {
		copy[indexCopy] = data[j]
		indexCopy--
        j--
	}
	return left + right + count
}
```