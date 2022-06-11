package question12

// 矩阵中的路径
func Run() string {
	matrix := [][]byte{
		{'a', 'b', 'c', 'e'},
		{'s', 'f', 'c', 's'},
		{'a', 'd', 'e', 'e'},
	}
	hasPath(matrix, "abcced")
	return "done!"
}

func hasPath(board [][]byte, word string) bool {
	row := len(board)
	col := len(board[0])
	// 烘托一下气氛
	var dfs func(int, int, int) bool
	dfs = func(row int, col int, index int) bool {
		// 边界条件的判断
		if row < 0 || row > len(board)-1 || col < 0 || col > len(board[0])-1 {
			return false
		}
		// 判断当前坐标是否符合
		if board[row][col] != word[index] {
			return false
		}
		// 好了，所有字符都找到路径了
		if index == len(word)-1 {
			return true
		}
		// 目标字符向后进一位，准备递归
		index++
		var temp byte
		/*
			把当前坐标的值保存下来用于后边复原
			为什么要修改当前坐标的值呢？
			题目说：如果一条路径经过了矩阵中的某一个格子，则该路径不能再进入该格子
		*/
		board[row][col], temp = '#', board[row][col]
		// 四个方向，只要有一个能查找到，就返回 true
		ret := dfs(row+1, col, index) || dfs(row-1, col, index) || dfs(row, col+1, index) || dfs(row, col-1, index)
		// 还原
		board[row][col] = temp
		return ret
	}
	// 翻箱倒柜地找
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			// 矩阵中每个节点都有可能是起点
			if dfs(i, j, 0) {
				return true
			}
		}
	}
	return false
}
