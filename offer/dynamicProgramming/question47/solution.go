package question47

import "fmt"

// 礼物的最大价值
func Run() string {
	var grid [][]int
	row := []int{1, 3, 1}
	grid = append(grid, row)
	row = []int{1, 5, 1}
	grid = append(grid, row)
	row = []int{4, 2, 1}
	grid = append(grid, row)
	res := maxValue(grid)
	fmt.Println(res)
	return "done!"
}

func maxValue(grid [][]int) int {
	// 初始化最上一行
	for j := 1; j < len(grid[0]); j++ {
		grid[0][j] = grid[0][j-1] + grid[0][j]
	}
	// 初始化最左一列
	for i := 1; i < len(grid); i++ {
		grid[i][0] = grid[i-1][0] + grid[i][0]
	}
	// 逐列
	for i := 1; i < len(grid); i++ {
		// 逐行
		for j := 1; j < len(grid[0]); j++ {
			grid[i][j] = max(grid[i-1][j], grid[i][j-1]) + grid[i][j]
		}
	}
	return grid[len(grid)-1][len(grid[0])-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
