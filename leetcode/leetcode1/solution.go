package leetcode1

// 两数之和
func Run() string {
	nums := []int{2, 7, 11, 15}
	TwoSum(nums, 9)
	return "done!"
}

func TwoSum(nums []int, target int) []int {
	var dict = make(map[int]int)
	for k, x := range nums {
		peer := target - x
		index, ok := dict[peer]
		if ok {
			return []int{index, k}
		}
		dict[x] = k
	}
	return nil
}
