package question81

// 调整数组顺序使奇数位于偶数前面(二)
func Run() string {
	array := []int{1, 2, 2, 3}
	reOrderArrayTwo(array)
	return "done!"
}

func reOrderArrayTwo(array []int) []int {
	end := len(array) - 1
	index := end
	for i := 0; i < len(array)-1; {
		if index < i {
			break
		}
		if array[i]%2 == 0 {
			array[index], array[i] = array[i], array[index]
			index--
		} else {
			i++
		}
	}
	return array
}
