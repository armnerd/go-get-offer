package question57

// 和为S的两个数字
func Run() string {
	array := []int{1, 2, 4, 7, 11, 15}
	FindNumbersWithSum(array, 15)
	return "done!"
}

func FindNumbersWithSum(nums []int, target int) []int {
	for i, j := 0, len(nums)-1; i < j; {
		if nums[i]+nums[j] > target {
			j--
		} else if nums[i]+nums[j] < target {
			i++
		} else {
			return []int{nums[i], nums[j]}
		}
	}
	return nil
}
