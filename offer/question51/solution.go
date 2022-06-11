package question51

// 数组中的逆序对
func Run() string {
	data := []int{1, 2, 3, 4, 5, 6, 7, 0}
	InversePairs(data)
	return "done!"
}

func InversePairs(data []int) int {
	if len(data) < 2 {
		return 0
	}

	replica := make([]int, len(data))
	copy(replica, data)

	count := mergeSort(data, replica, 0, len(data)-1)
	return count % 1000000007
}

func mergeSort(data, copy []int, start, end int) int {
	if start == end {
		copy[start] = data[start]
		return 0
	}
	length := (end - start) / 2
	left := mergeSort(copy, data, start, start+length)
	right := mergeSort(copy, data, start+length+1, end)

	i := start + length
	j := end
	indexCopy := end
	count := 0

	for i >= start && j >= start+length+1 {
		if data[i] > data[j] {
			copy[indexCopy] = data[i]
			indexCopy--
			i--
			count += j - start - length
		} else {
			copy[indexCopy] = data[j]
			indexCopy--
			j--
		}
	}
	for i >= start {
		copy[indexCopy] = data[i]
		indexCopy--
		i--
	}
	for j >= start+length+1 {
		copy[indexCopy] = data[j]
		indexCopy--
		j--
	}
	return left + right + count
}
