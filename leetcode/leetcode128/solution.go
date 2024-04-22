package leetcode128

// 最长连续序列
func Run() string {
	nums := []int{9, 1, 4, 7, 3, -1, 0, 5, 8, -1, 6}
	longestConsecutive(nums)
	return "done!"
}

func longestConsecutive(nums []int) int {
	numSet := make(map[int]bool)
	for _, v := range nums {
		numSet[v] = true
	}
	longestStreak := 0
	for num := range numSet {
		if !numSet[num-1] {
			currentNum := num
			currentStreak := 1
			for numSet[currentNum+1] {
				currentNum++
				currentStreak++
			}
			if longestStreak < currentStreak {
				longestStreak = currentStreak
			}
		}
	}
	return longestStreak
}
