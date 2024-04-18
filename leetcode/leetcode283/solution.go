package leetcode283

// 移动零
func Run() string {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	return "done!"
}

func moveZeroes(nums []int) {
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[j] = nums[i]
			j += 1
		}
	}
	for i := j; i < len(nums); i++ {
		nums[i] = 0
	}
}
