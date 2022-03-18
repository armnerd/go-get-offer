package question40

// 最小的K个数
func Run() string {
	examples := []struct {
		Input []int
		K     int
	}{
		{[]int{4, 5, 1, 6, 2, 7, 3, 8}, 4},
		{[]int{4, 5, 1, 6, 2, 7, 3, 8}, 1},
	}
	for _, example := range examples {
		GetLeastNumbers_Solution(example.Input, example.K)
	}
	return "done!"
}

func GetLeastNumbers_Solution(input []int, k int) []int {
	// 参数检查
	if len(input) == 0 || k <= 0 {
		return []int{}
	}
	if k >= len(input) {
		return input
	}
	var max func(int, int)
	max = func(i int, length int) {
		left := 2*i + 1
		right := 2*i + 2
		largest := 0
		if left < length && input[left] > input[i] {
			largest = left
		} else {
			largest = i
		}
		if right < length && input[right] > input[largest] {
			largest = right
		}
		if largest != i {
			input[i], input[largest] = input[largest], input[i]
			// 比较其父节点
			max(largest, length)
		}
	}
	// 建堆
	length := len(input)
	for i := length/2 - 1; i >= 0; i-- {
		max(i, length)
	}
	// 排序
	for i := length - 1; i > 0; i-- {
		// 将当前最大 { input[0] } 的放在队尾 { input[i] }
		input[i], input[0] = input[0], input[i]
		// 堆的大小减 1
		length--
		// 在小堆上重建堆
		max(0, length)
	}
	return input[:k]
}
