package question3

// 数组中重复的数字
func Run() string {
	numbers := []int{2, 3, 1, 0, 2, 5, 3}
	duplicate(numbers)
	return "done!"
}

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
