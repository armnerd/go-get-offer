package leetcode51

import (
	"strings"
)

// N 皇后
func Run() string {
	SolveNQueens(4)
	return "done!"
}

func SolveNQueens(n int) [][]string {
	var res [][]string
	// 初始化棋盘
	var board = make([][]string, n)
	for i := 0; i < n; i++ {
		var tmp []string
		for j := 0; j < n; j++ {
			tmp = append(tmp, ".")
		}
		board[i] = tmp
	}
	var backtrack func(row int)
	backtrack = func(row int) {
		if row == n {
			//收敛结果
			var solution []string
			for _, row := range board {
				solution = append(solution, strings.Join(row, ""))
			}
			res = append(res, solution)
			return
		}
		for col := 0; col < n; col++ {
			// 过滤掉不满足条件的剩下就是可选择的列表了
			if !isValid(board, row, col) {
				continue
			}
			board[row][col] = "Q" // 先在这个位置把 Q 放下
			backtrack(row + 1)    // 去下一层做探索
			board[row][col] = "." // 恢复选择
		}
	}
	backtrack(0)
	return res
}

func isValid(board [][]string, row, col int) bool {
	n := len(board)
	// 检查本列是否有冲突
	for i := 0; i < row; i++ {
		if board[i][col] == "Q" {
			return false
		}
	}
	// 检查右上方
	for i, j := row-1, col+1; i >= 0 && j < n; {
		if board[i][j] == "Q" {
			return false
		}
		i--
		j++
	}
	// 检查左上方
	for i, j := row-1, col-1; i >= 0 && j >= 0; {
		if board[i][j] == "Q" {
			return false
		}
		i--
		j--
	}
	return true
}
