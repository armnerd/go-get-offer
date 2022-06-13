package question53

// 数字在升序数组中出现的次数
func Run() string {
	num := []int{1, 2, 3, 3, 3, 3, 4, 5}
	size := 3
	GetNumberOfK(num, size)
	return "done!"
}

func GetNumberOfK(data []int, k int) int {
	cur := 0
	end := len(data)
	for {
		if cur >= end {
			break
		}
		mid := cur + (end-cur)/2
		if data[mid] < k {
			cur = mid + 1
		} else {
			end = mid
		}
	}
	lbound := cur
	cur = 0
	end = len(data)
	for {
		if cur >= end {
			break
		}
		mid := cur + (end-cur)/2
		if data[mid] <= k {
			cur = mid + 1
		} else {
			end = mid
		}
	}
	rbound := cur
	return rbound - lbound
}
