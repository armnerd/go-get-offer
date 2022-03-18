package question21

// 调整数组顺序使奇数位于偶数前面(一)
func Run() string {
	array := []int{1, 2, 2, 3}
	reOrderArray(array)
	return "done!"
}

func reOrderArray(array []int) []int {
	res := []int{}
	temp := []int{}
	for i := 0; i < len(array); i++ {
		if array[i]%2 != 0 {
			res = append(res, array[i])
		} else {
			temp = append(temp, array[i])
		}
	}
	res = append(res, temp...)
	return res
}
