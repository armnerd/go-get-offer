package question13

import "fmt"

// 机器人的运动范围
func Run() string {
	res := movingCount(1, 2, 3)
	fmt.Println(res)
	return "done!"
}

func movingCount(threshold, rows, cols int) int {
	// 初始化做标记的矩阵
	flags := make([][]int, rows)
	for i := 0; i < rows; i++ {
		flags[i] = make([]int, cols)
	}
	// 烘托气氛
	var handler func(int, int) int
	handler = func(i, j int) int {
		// 边界检查
		if i < 0 || i >= rows || j < 0 || j >= cols {
			return 0
		}
		// 是否访问过
		if flags[i][j] != 0 {
			return 0
		}
		// 检查当前坐标是否满足条件
		if check(threshold, i, j) {
			return 0
		}
		// 标记访问过
		flags[i][j] = 1
		return 1 + handler(i-1, j) + handler(i+1, j) + handler(i, j-1) + handler(i, j+1)
	}
	return handler(0, 0)
}

func check(threshold, i, j int) bool {
	sum := 0
	for i > 0 {
		sum += i % 10 // 取个位数
		i = i / 10    // 移除个位
	}
	for j > 0 {
		sum += j % 10 // 取个位数
		j = j / 10    // 移除个位
	}
	return sum > threshold
}
