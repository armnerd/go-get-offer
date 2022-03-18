package question42

// 连续子数组的最大和
func Run() string {
	data := []int{1, -2, 3, 10, -4, 7, 2, -5}
	FindGreatestSumOfSubArray(data)
	return "done!"
}

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
