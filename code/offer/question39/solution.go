package question39

// 数组中出现次数超过一半的数字
func Run() string {
	numbers := []int{1, 2, 3, 2, 2, 2, 5, 4, 2}
	MoreThanHalfNum_Solution(numbers)
	return "done!"
}

func MoreThanHalfNum_Solution(numbers []int) int {
	res := 0
	limit := len(numbers)/2 + 1
	record := make(map[int]int)
	for _, v := range numbers {
		record[v] += 1
		if record[v] == limit {
			res = v
		}
	}
	return res
}
