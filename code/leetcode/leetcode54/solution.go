package leetcode54

// 螺旋矩阵
func Run() string {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	SpiralOrder(matrix)
	return "done!"
}

func SpiralOrder(matrix [][]int) []int {
	res := []int{}
	if len(matrix) == 0 {
		return res
	}
	// 赋值上下左右边界
	upLimit := 0
	downLimit := len(matrix) - 1
	leftLimit := 0
	rightLimit := len(matrix[0]) - 1
	for {
		// 【1】向右移动直到最右
		for i := leftLimit; i <= rightLimit; i++ {
			res = append(res, matrix[upLimit][i])
		}
		// 重新设定上边界，若上边界大于下边界，则遍历遍历完成，下同
		upLimit += 1
		if upLimit > downLimit {
			break
		}
		// 【2】向下
		for i := upLimit; i <= downLimit; i++ {
			res = append(res, matrix[i][rightLimit])
		}
		// 重新设定边界
		rightLimit -= 1
		if leftLimit > rightLimit {
			break
		}
		// 【3】向左
		for i := rightLimit; i >= leftLimit; i-- {
			res = append(res, matrix[downLimit][i])
		}
		// 重新设定边界
		downLimit -= 1
		if upLimit > downLimit {
			break
		}
		// 【4】向上
		for i := downLimit; i >= upLimit; i-- {
			res = append(res, matrix[i][leftLimit])
		}
		// 重新设定边界
		leftLimit += 1
		if leftLimit > rightLimit {
			break
		}
	}
	return res
}
