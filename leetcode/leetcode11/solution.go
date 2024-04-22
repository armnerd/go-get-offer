package leetcode11

// 盛最多水的容器
func Run() string {
	nums := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	maxArea(nums)
	return "done!"
}

func maxArea(height []int) int {
	left := 0
	right := len(height) - 1
	res := 0
	for left < right {
		res = max(res, min(height[left], height[right])*(right-left))
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
