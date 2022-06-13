package question59

// 滑动窗口的最大值
func Run() string {
	num := []int{2, 3, 4, 2, 6, 2, 5, 1}
	size := 3
	MaxInWindows(num, size)
	return "done!"
}

func MaxInWindows(num []int, size int) []int {
	res := make([]int, 0)
	if len(num) == 0 || size == 0 {
		return res
	}
	dq := make([]int, 0)
	for i, v := range num {
		// 辅助队尾的值小于当前元素，就替换为当前元素
		for len(dq) != 0 && num[dq[len(dq)-1]] < v {
			/*
				说明以此下标为右边的窗口用不到队列的最后一个元素了
				而且以此下标 +n 为右边的后边窗口也用不到队列的最后一个元素了，因为当前元素都比这个元素大了
			*/
			dq = dq[:len(dq)-1]
		}
		// 压入的是当前元素的下标
		dq = append(dq, i)
		// 判断队列的头部的下标是否过期
		if dq[0] == i-size {
			dq = dq[1:]
		}
		// 判断是否形成了窗口
		if i+1 >= size {
			// 辅助队列最前即是窗口最大值
			res = append(res, num[dq[0]])
		}
	}
	return res
}
