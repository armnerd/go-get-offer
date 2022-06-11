package question61

import "sort"

// 扑克牌顺子
func Run() string {
	numbers := []int{6, 0, 2, 0, 4}
	IsContinuous(numbers)
	return "done!"
}

func IsContinuous(nums []int) bool {
	sort.Ints(nums)
	start := 0
	// 统计数组中的0的个数
	for nums[start] == 0 {
		start++
	}
	for i := start + 1; i < len(nums); i++ {
		// 如果空缺的总数大于个数 或者 数组中的非0数字重复出现
		if nums[i]-nums[i-1]-1 > start || nums[i] == nums[i-1] {
			return false
		}
		start -= (nums[i] - nums[i-1] - 1)
	}
	return start >= 0
}
