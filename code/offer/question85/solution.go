package question85

import "fmt"

// 连续子数组的最大和(二)
func Run() string {
	data := []int{1, -2, 3, 10, -4, 7, 2, -5}
	res := FindGreatestSumOfSubArray(data)
	fmt.Println(res)
	return "done!"
}

func FindGreatestSumOfSubArray(array []int) []int {
	res := make([]int, 0)
	dp := make([]int, len(array))
	dp[0] = array[0]
	var maxSum = dp[0]
	var left, right int
	var leftRes, rightRes int
	for i := 1; i < len(array); i++ {
		right++
		dp[i] = max(dp[i-1]+array[i], array[i])
		if dp[i-1]+array[i] < array[i] {
			// 当前值比和前边最大加起来都大了，那这就是新边界
			left = right
		}
		if dp[i] > maxSum || dp[i] == maxSum && (right-left+1) > (rightRes-leftRes+1) {
			// 和更大或者和一样数组更长
			maxSum = dp[i]
			leftRes = left
			rightRes = right
		}
	}
	for i := leftRes; i <= rightRes; i++ {
		res = append(res, array[i])
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
